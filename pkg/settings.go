package main

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type PluginSettings struct {
	Storage     string `json:"storage"`
	QueryParams string `json:"queryParams"`
}

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
