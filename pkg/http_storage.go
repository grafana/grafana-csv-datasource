package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type httpStorage struct {
	httpClient     *http.Client
	settings       backend.DataSourceInstanceSettings
	query          queryModel
}

func newHTTPStorage(query queryModel, instanceSettings backend.DataSourceInstanceSettings, httpClient *http.Client) (*httpStorage, error) {
	return &httpStorage{
		httpClient:     httpClient,
		settings:       instanceSettings,
		query:          query,
	}, nil
}

func (c *httpStorage) do() (*http.Response, error) {
	req, err := newRequestFromQuery(c.settings, c.query)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, backend.DownstreamError(err)
	}

	return resp, nil
}

func (c *httpStorage) Open() (io.ReadCloser, error) {
	resp, err := c.do()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return nil, backend.DownstreamErrorf("unexpected response status: %s", resp.Status)
	}

	return resp.Body, nil
}

func (c *httpStorage) Stat(logger log.Logger) error {
	resp, err := c.do()
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.Warn("failed to close response body", "error", err)
		}
	}()

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}

func newRequestFromQuery(settings backend.DataSourceInstanceSettings, query queryModel) (*http.Request, error) {
	customSettings, err := LoadPluginSettings(settings)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(settings.URL + query.Path)
	if err != nil {
		return nil, backend.DownstreamError(err)
	}

	// we need to verify that `query.Path` did not modify the hostname by doing tricks
	settingsURL, err := url.Parse(settings.URL)
	if err != nil {
		return nil, backend.DownstreamErrorf("invalid URL")
	}

	if settingsURL.Host != u.Host {
		// the host got changed by adding the path to it. this must not happen.
		return nil, backend.DownstreamErrorf("invalid URL + path combination")
	}

	params := make(url.Values)
	for _, p := range query.Params {
		if len(p) >= 2 {
			params.Set(p[0], p[1])
		}
	}

	// Query params set by admin overrides params set by query editor.
	values, err := url.ParseQuery(customSettings.QueryParams)
	if err != nil {
		return nil, backend.DownstreamError(err)
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
		return nil, backend.DownstreamError(err)
	}

	for _, p := range query.Headers {
		if len(p) >= 2 {
			req.Header.Set(p[0], p[1])
		}
	}

	return req, nil
}
