package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-csv-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func LoadPluginSettings(source backend.DataSourceInstanceSettings) (*models.PluginSettings, error) {
	settings := models.PluginSettings{}
	err := json.Unmarshal(source.JSONData, &settings)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal PluginSettings json: %w", err)
	}

	// Default to HTTP storage for backwards compatibility.
	if settings.Storage == "" {
		settings.Storage = "http"
	}

	return &settings, nil
}
