package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/marcusolsson/grafana-csv-datasource/pkg/httpclient"
)

type httpStorage struct {
	httpClient     *http.Client
	settings       *backend.DataSourceInstanceSettings
	customSettings dataSourceSettings
}

func newHTTPStorage(instance *dataSourceInstance, logger log.Logger) (*httpStorage, error) {
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
	}, nil
}

func (c *httpStorage) do() (*http.Response, error) {
	u, err := url.Parse(c.settings.URL)
	if err != nil {
		return nil, err
	}

	values, err := url.ParseQuery(c.customSettings.QueryParams)
	if err != nil {
		return nil, err
	}
	u.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(req)
}

func (c *httpStorage) open() (io.ReadCloser, error) {
	resp, err := c.do()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return resp.Body, nil
}

func (c *httpStorage) stat() error {
	resp, err := c.do()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}
