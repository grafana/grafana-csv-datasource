package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type httpStorage struct {
	httpClient     *http.Client
	settings       *backend.DataSourceInstanceSettings
	customSettings dataSourceSettings
	query          dataSourceQuery
}

func newHTTPStorage(instance *dataSourceInstance, query dataSourceQuery, logger log.Logger) (*httpStorage, error) {
	customSettings, err := instance.Settings()
	if err != nil {
		return nil, err
	}

	httpOptions, err := instance.settings.HTTPClientOptions()
	if err != nil {
		return nil, err
	}

	httpClient, err := httpclient.New(httpOptions)
	if err != nil {
		return nil, err
	}

	return &httpStorage{
		httpClient:     httpClient,
		settings:       &instance.settings,
		customSettings: customSettings,
		query:          query,
	}, nil
}

func (c *httpStorage) do() (*http.Response, error) {
	req, err := newRequestFromQuery(c.settings, c.customSettings, c.query)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(req)
}

func (c *httpStorage) Open() (io.ReadCloser, error) {
	resp, err := c.do()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return resp.Body, nil
}

func (c *httpStorage) Stat() error {
	resp, err := c.do()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}

func newRequestFromQuery(settings *backend.DataSourceInstanceSettings, customSettings dataSourceSettings, query dataSourceQuery) (*http.Request, error) {
	u, err := url.Parse(settings.URL + query.Path)
	if err != nil {
		return nil, err
	}

	// we need to verify that `query.Path` did not modify the hostname by doing tricks
	settingsURL, err := url.Parse(settings.URL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL")
	}

	if settingsURL.Host != u.Host {
		// the host got changed by adding the path to it. this must not happen.
		return nil, fmt.Errorf("invalid URL + path combination")
	}

	params := make(url.Values)
	for _, p := range query.Params {
		params.Set(p[0], p[1])
	}

	// Query params set by admin overrides params set by query editor.
	values, err := url.ParseQuery(customSettings.QueryParams)
	if err != nil {
		return nil, err
	}
	for k, v := range values {
		params[k] = v
	}

	u.RawQuery = params.Encode()

	var method string
	if query.Method != "" {
		method = query.Method
	} else {
		method = "GET"
	}

	req, err := http.NewRequest(method, u.String(), strings.NewReader(query.Body))
	if err != nil {
		return nil, err
	}

	for _, p := range query.Headers {
		req.Header.Set(p[0], p[1])
	}

	return req, nil
}
