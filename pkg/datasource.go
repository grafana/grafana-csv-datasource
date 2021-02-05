package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path/filepath"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// dataSource defines common operations for all instances of this data source.
type dataSource struct {
	im     instancemgmt.InstanceManager
	logger log.Logger
}

func newDataSource(logger log.Logger) *dataSource {
	im := datasource.NewInstanceManager(newDataSourceInstance)

	return &dataSource{
		im:     im,
		logger: logger,
	}
}

func (ds *dataSource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	instance, err := ds.getInstance(req.PluginContext)
	if err != nil {
		return nil, err
	}

	responses := make(backend.Responses)

	for _, q := range req.Queries {
		responses[q.RefID] = ds.query(ctx, q, instance)
	}

	return &backend.QueryDataResponse{
		Responses: responses,
	}, nil
}

// query is a helper that performs the actual data source query.
func (ds *dataSource) query(ctx context.Context, query backend.DataQuery, instance *dataSourceInstance) backend.DataResponse {
	var opts csvOptions
	if err := json.Unmarshal(query.JSON, &opts); err != nil {
		return backend.DataResponse{Error: err}
	}

	store, err := newStorage(instance, ds.logger)
	if err != nil {
		return backend.DataResponse{Error: err}
	}

	f, err := store.Open()
	if err != nil {
		return backend.DataResponse{Error: err}
	}
	defer f.Close()

	fields, err := parseCSV(opts, f)
	if err != nil {
		return backend.DataResponse{Error: err}
	}

	return backend.DataResponse{
		Frames: []*data.Frame{{
			Name:   filepath.Base(instance.settings.URL),
			Fields: fields,
		}},
	}
}

// CheckHealth returns the current health of the backend.
func (ds *dataSource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	instance, err := ds.getInstance(req.PluginContext)
	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: err.Error(),
		}, nil
	}

	store, err := newStorage(instance, ds.logger)
	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: err.Error(),
		}, nil
	}

	if err := store.Stat(); err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: err.Error(),
		}, nil
	}

	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Success",
	}, nil
}

func (ds *dataSource) getInstance(ctx backend.PluginContext) (*dataSourceInstance, error) {
	instance, err := ds.im.Get(ctx)
	if err != nil {
		return nil, err
	}
	return instance.(*dataSourceInstance), nil
}

// dataSourceInstance represents a single instance of this data source.
type dataSourceInstance struct {
	httpClient *http.Client
	settings   backend.DataSourceInstanceSettings
}

func newDataSourceInstance(settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	return &dataSourceInstance{
		httpClient: &http.Client{},
		settings:   settings,
	}, nil
}

func (s *dataSourceInstance) Dispose() {}

func (s *dataSourceInstance) Settings() (dataSourceSettings, error) {
	return newDataSourceSettings(s.settings)
}

type dataSourceSettings struct {
	QueryParams string `json:"queryParams"`
	Storage     string `json:"storage"`
}

func newDataSourceSettings(instanceSettings backend.DataSourceInstanceSettings) (dataSourceSettings, error) {
	var settings dataSourceSettings
	if err := json.Unmarshal(instanceSettings.JSONData, &settings); err != nil {
		return dataSourceSettings{}, err
	}
	return settings, nil
}

type storage interface {
	Open() (io.ReadCloser, error)
	Stat() error
}

func newStorage(instance *dataSourceInstance, logger log.Logger) (storage, error) {
	sett, err := instance.Settings()
	if err != nil {
		return nil, err
	}

	// Default to HTTP storage for backwards compatibility.
	if sett.Storage == "" {
		return newHTTPStorage(instance, logger)
	}

	switch sett.Storage {
	case "http":
		return newHTTPStorage(instance, logger)
	case "local":
		return newLocalStorage(instance, logger)
	default:
		return nil, errors.New("unsupported storage type")
	}
}
