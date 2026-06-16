package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/experimental/concurrent"
)

var (
	_ backend.QueryDataHandler      = (*Datasource)(nil)
	_ backend.CheckHealthHandler    = (*Datasource)(nil)
	_ instancemgmt.InstanceDisposer = (*Datasource)(nil)
)

type Datasource struct {
	allowLocalMode bool
	HTTPClient     *http.Client
}

// Dispose implements instancemgmt.InstanceDisposer.
func (d *Datasource) Dispose() {
	backend.Logger.Info("Disposing datasource")
}

func NewDatasource(ctx context.Context, instanceSettings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	opts, err := instanceSettings.HTTPClientOptions(ctx)
	if err != nil {
		return nil, err
	}
	httpClient, err := httpclient.New(opts)
	if err != nil {
		return nil, err
	}
	return &Datasource{
		allowLocalMode: os.Getenv("GF_PLUGIN_ALLOW_LOCAL_MODE") == "true",
		HTTPClient:     httpClient,
	}, nil
}

func (d *Datasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	return concurrent.QueryData(ctx, req, d.query, 10)
}

type queryModel struct {
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

func (d *Datasource) query(ctx context.Context, q concurrent.Query) backend.DataResponse {
	var response backend.DataResponse
	logger := backend.Logger.FromContext(ctx)

	// Unmarshal the JSON into our queryModel.
	var qm queryModel
	err := json.Unmarshal(q.DataQuery.JSON, &qm)
	if err != nil {
		return backend.ErrorResponseWithErrorSource(err)
	}

	store, err := d.newStorage(*q.PluginContext.DataSourceInstanceSettings, qm)
	if err != nil {
		return backend.ErrorResponseWithErrorSource(err)
	}

	f, err := store.Open()
	if err != nil {
		return backend.ErrorResponseWithErrorSource(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			logger.Warn("failed to close file", "error", err)
		}
	}()

	frame := data.NewFrame(q.PluginContext.DataSourceInstanceSettings.URL)

	fields, err := parseCSV(qm.csvOptions, qm.Experimental.Regex, f, logger)
	if err != nil {
		return backend.ErrorResponseWithErrorSource(err)
	}

	frame.Fields = fields

	// add the frames to the response.
	response.Frames = append(response.Frames, frame)

	return response
}

// CheckHealth returns the current health of the backend.
func (d *Datasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	res := &backend.CheckHealthResult{}
	logger := backend.Logger.FromContext(ctx)

	settings, err := LoadPluginSettings(*req.PluginContext.DataSourceInstanceSettings)
	if err != nil {
		logger.Warn("failed to load plugin settings", "error", err)
		res.Status = backend.HealthStatusError
		res.Message = "Failed to load plugin settings"
		return res, nil
	}

	if settings.Storage == "http" && req.PluginContext.DataSourceInstanceSettings.URL == "" {
		res.Status = backend.HealthStatusError
		res.Message = "URL is required for HTTP storage"
		return res, nil
	}

	store, err := d.newStorage(*req.PluginContext.DataSourceInstanceSettings, queryModel{})
	if err != nil {
		logger.Warn("failed to create storage", "error", err)
		res.Status = backend.HealthStatusError
		res.Message = err.Error()
		return res, nil
	}

	if err := store.Stat(logger); err != nil {
		logger.Warn("failed to connect to storage", "error", err)
		res.Status = backend.HealthStatusError
		res.Message = err.Error()
		return res, nil
	}

	res.Status = backend.HealthStatusOk
	res.Message = "Success"
	return res, nil
}

type storage interface {
	Open() (io.ReadCloser, error)
	Stat(logger log.Logger) error
}

func (d *Datasource) newStorage(instanceSettings backend.DataSourceInstanceSettings, query queryModel) (storage, error) {
	settings, err := LoadPluginSettings(instanceSettings)
	if err != nil {
		return nil, err
	}

	if settings.Storage == "local" && !d.allowLocalMode {
		return nil, backend.DownstreamErrorf("local mode has been disabled by your administrator")
	}

	switch settings.Storage {
	case "local":
		return newLocalStorage(query, instanceSettings)
	case "http":
		fallthrough
	default:
		// Default to HTTP storage for backwards compatibility.
		return newHTTPStorage(query, instanceSettings, d.HTTPClient)
	}
}
