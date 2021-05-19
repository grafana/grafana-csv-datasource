package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type localStorage struct {
	settings       *backend.DataSourceInstanceSettings
	customSettings dataSourceSettings
	query          dataSourceQuery
}

func newLocalStorage(instance *dataSourceInstance, query dataSourceQuery, logger log.Logger) (*localStorage, error) {
	customSettings, err := instance.Settings()
	if err != nil {
		return nil, err
	}

	return &localStorage{
		settings:       &instance.settings,
		customSettings: customSettings,
		query:          query,
	}, nil
}

func (c *localStorage) Open() (io.ReadCloser, error) {
	if c.query.Path == "" {
		return os.Open(filepath.ToSlash(c.settings.URL))
	}

	fullPath := filepath.Join(c.settings.URL, c.query.Path)

	// Ensure users can't slip out of the directory configured by the admin.
	if !strings.HasPrefix(fullPath, filepath.Clean(c.settings.URL)+string(os.PathSeparator)) {
		return nil, fmt.Errorf("illegal file path: %s", c.query.Path)
	}

	return os.Open(filepath.ToSlash(fullPath))
}

func (c *localStorage) Stat() error {
	_, err := os.Stat(filepath.ToSlash(c.settings.URL))
	return err
}
