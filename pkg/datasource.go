package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type dataSourceQuery struct {
	csvOptions

	Method  string      `json:"method"`
	Path    string      `json:"path"`
	Params  [][2]string `json:"params"`
	Headers [][2]string `json:"headers"`
	Body    string      `json:"body"`

	Experimental struct {
		Regex bool `json:"regex"`
	} `json:"experimental"`
}

// dataSource defines common operations for all instances of this data source.
type dataSource struct {
	instanceManager instancemgmt.InstanceManager
	logger          log.Logger
}

func newDataSource(logger log.Logger) *dataSource {
	instanceManager := datasource.NewInstanceManager(newDataSourceInstance)

	return &dataSource{
		instanceManager: instanceManager,
		logger:          logger,
	}
}

func (ds *dataSource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	instance, err := ds.getInstance(req.PluginContext)
	if err != nil {
		return nil, err
	}

	responses := make(backend.Responses)

	var wg sync.WaitGroup
	wg.Add(len(req.Queries))

	for _, q := range req.Queries {
		go func(q backend.DataQuery) {
			responses[q.RefID] = ds.query(ctx, q, instance)
			wg.Done()
		}(q)
	}

	// Wait for all queries to finish before returning the result.
	wg.Wait()

	return &backend.QueryDataResponse{
		Responses: responses,
	}, nil
}

// query is a helper that performs a single data source query.
func (ds *dataSource) query(ctx context.Context, query backend.DataQuery, instance *dataSourceInstance) backend.DataResponse {
	var dsQuery dataSourceQuery
	if err := json.Unmarshal(query.JSON, &dsQuery); err != nil {
		return backend.DataResponse{Error: err}
	}

	store, err := newStorage(instance, dsQuery, ds.logger)
	if err != nil {
		return backend.DataResponse{Error: err}
	}

	f, err := store.Open()
	if err != nil {
		return backend.DataResponse{Error: err}
	}
	defer f.Close()

	fields, err := parseCSV(dsQuery.csvOptions, dsQuery.Experimental.Regex, f, ds.logger)
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

	var query dataSourceQuery

	store, err := newStorage(instance, query, ds.logger)
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
	instance, err := ds.instanceManager.Get(ctx)
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

func newStorage(instance *dataSourceInstance, query dataSourceQuery, logger log.Logger) (storage, error) {
	settings, err := instance.Settings()
	if err != nil {
		return nil, err
	}

	// Default to HTTP storage for backwards compatibility.
	if settings.Storage == "" {
		return newHTTPStorage(instance, query, logger)
	}

	switch settings.Storage {
	case "http":
		return newHTTPStorage(instance, query, logger)
	case "local":
		return newLocalStorage(instance, query, logger)
	default:
		return nil, errors.New("unsupported storage type")
	}
}
