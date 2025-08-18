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
	settings       backend.DataSourceInstanceSettings
	customSettings *PluginSettings
	query          queryModel
}

func newLocalStorage(query queryModel, instanceSettings backend.DataSourceInstanceSettings) (*localStorage, error) {
	customSettings, err := LoadPluginSettings(instanceSettings)
	if err != nil {
		return nil, err
	}

	return &localStorage{
		settings:       instanceSettings,
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

func (c *localStorage) Stat(_ log.Logger) error {
	_, err := os.Stat(filepath.ToSlash(c.settings.URL))
	return err
}
