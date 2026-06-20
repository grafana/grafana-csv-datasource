package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-csv-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

// PluginSettings is the datasource settings model. The canonical definition
// lives in pkg/models so it can be shared with the schema conformance tests.
type PluginSettings = models.PluginSettings

func LoadPluginSettings(source backend.DataSourceInstanceSettings) (*PluginSettings, error) {
	settings := PluginSettings{}
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
