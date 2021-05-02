package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/marcusolsson/grafana-csv-datasource/pkg/httpclient"
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

	httpClient, err := httpclient.New(&instance.settings, 10*time.Second, logger)
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

	req.Header.Set("Accept", "text/csv")

	for _, p := range query.Headers {
		req.Header.Set(p[0], p[1])
	}

	return req, nil
}
