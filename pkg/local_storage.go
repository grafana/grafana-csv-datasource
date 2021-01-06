package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type localStorage struct {
	settings       *backend.DataSourceInstanceSettings
	customSettings dataSourceSettings
}

func newLocalStorage(instance *dataSourceInstance, logger log.Logger) (*localStorage, error) {
	customSettings, err := instance.Settings()
	if err != nil {
		return nil, err
	}

	return &localStorage{
		settings:       &instance.settings,
		customSettings: customSettings,
	}, nil
}

func (c *localStorage) Open() (io.ReadCloser, error) {
	return os.Open(filepath.ToSlash(c.settings.URL))
}

func (c *localStorage) Stat() error {
	_, err := os.Stat(filepath.ToSlash(c.settings.URL))
	return err
}
