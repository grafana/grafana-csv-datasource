package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"

	"github.com/marcusolsson/grafana-csv-datasource/pkg/httpclient"
)

type fieldSchema struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type queryModel struct {
	Delimiter     string        `json:"delimiter"`
	Header        bool          `json:"header"`
	IgnoreUnknown bool          `json:"ignoreUnknown"`
	Schema        []fieldSchema `json:"schema"`
	SkipRows      int           `json:"skipRows"`
}

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
		responses[q.RefID] = ds.query(ctx, q, &instance.settings)
	}

	return &backend.QueryDataResponse{
		Responses: responses,
	}, nil
}

// query is a helper that performs the actual data source query.
func (ds *dataSource) query(ctx context.Context, query backend.DataQuery, settings *backend.DataSourceInstanceSettings) backend.DataResponse {
	httpClient, err := httpclient.New(settings, 10*time.Second, ds.logger)
	if err != nil {
		return backend.DataResponse{
			Error: err,
		}
	}

	req, err := http.NewRequest("GET", settings.URL, nil)
	if err != nil {
		return backend.DataResponse{
			Error: err,
		}
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return backend.DataResponse{
			Error: err,
		}
	}
	defer resp.Body.Close()

	var q queryModel
	if err := json.Unmarshal(query.JSON, &q); err != nil {
		return backend.DataResponse{
			Error: err,
		}
	}

	fields, err := parseCSV(q, resp.Body)
	if err != nil {
		return backend.DataResponse{
			Error: err,
		}
	}

	return backend.DataResponse{
		Frames: []*data.Frame{{
			Name:   filepath.Base(settings.URL),
			Fields: fields,
		}},
	}
}

func parseCSV(query queryModel, r io.Reader) ([]*data.Field, error) {
	// Read one byte at a time until we've counted newlines equal to the number
	// of skipped rows.
	for i := 0; i < query.SkipRows; i++ {
		buf := make([]byte, 1)
		for {
			_, err := r.Read(buf)
			if err != nil || buf[0] == '\n' {
				break
			}
		}
	}

	rd := csv.NewReader(r)

	if len(query.Delimiter) == 1 {
		rd.Comma = rune(query.Delimiter[0])
	}

	records, err := rd.ReadAll()
	if err != nil {
		return nil, err
	}

	fieldType := func(str string) data.FieldType {
		switch str {
		case "number":
			return data.FieldTypeNullableFloat64
		case "boolean":
			return data.FieldTypeNullableBool
		case "time":
			return data.FieldTypeNullableTime
		default:
			return data.FieldTypeNullableString
		}
	}

	header := records[0]
	rows := records[1:]

	fields := make([]*data.Field, 0)

	// Create fields from schema.
	for _, name := range header {
		sch, ok := schemaContains(query.Schema, name)
		if !ok {
			if query.IgnoreUnknown {
				// Add a null field to maintain index.
				fields = append(fields, nil)
				continue
			} else {
				sch = fieldSchema{Name: name, Type: "string"}
			}
		}
		f := data.NewFieldFromFieldType(fieldType(sch.Type), len(rows))
		f.Name = name
		fields = append(fields, f)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(fields))

	for fieldIdx := range fields {
		go func(fieldIdx int) {
			defer wg.Done()

			f := fields[fieldIdx]
			if f == nil {
				return
			}

			var timeLayout string
			if f.Type() == data.FieldTypeNullableTime {
				layout, err := detectTimeLayoutNaive(rows[0][fieldIdx])
				if err == nil {
					timeLayout = layout
				}
			}

			for rowIdx := 0; rowIdx < f.Len(); rowIdx++ {
				value := rows[rowIdx][fieldIdx]

				switch f.Type() {
				case data.FieldTypeNullableFloat64:
					n, err := strconv.ParseFloat(value, 10)
					if err != nil {
						f.Set(rowIdx, nil)
						continue
					}
					f.Set(rowIdx, &n)
				case data.FieldTypeNullableBool:
					n, err := strconv.ParseBool(value)
					if err != nil {
						f.Set(rowIdx, nil)
						continue
					}
					f.Set(rowIdx, &n)
				case data.FieldTypeNullableTime:
					n, err := strconv.ParseInt(value, 10, 64)
					if err == nil {
						t := time.Unix(n, 0)
						f.Set(rowIdx, &t)
						continue
					}

					if timeLayout != "" {
						t, err := time.Parse(timeLayout, value)
						if err == nil {
							f.Set(rowIdx, &t)
							continue
						}
					}

					f.Set(rowIdx, nil)
				default:
					f.Set(rowIdx, &value)
				}
			}
		}(fieldIdx)
	}

	wg.Wait()

	// Remove ignored fields from result.
	var res []*data.Field
	for _, f := range fields {
		if f != nil {
			res = append(res, f)
		}
	}

	return res, nil
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

	resp, err := http.Get(instance.settings.URL)
	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: err.Error(),
		}, nil
	}

	if resp.StatusCode != 200 {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusError,
			Message: resp.Status,
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

func newDataSourceInstance(setting backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	return &dataSourceInstance{
		httpClient: &http.Client{},
		settings:   setting,
	}, nil
}

func (s *dataSourceInstance) Dispose() {}

func (s *dataSourceInstance) Settings() (dataSourceSettings, error) {
	return newDataSourceSettings(s.settings)
}

type dataSourceSettings struct {
	URL string `json:"url"`
}

func newDataSourceSettings(instanceSettings backend.DataSourceInstanceSettings) (dataSourceSettings, error) {
	var settings dataSourceSettings
	if err := json.Unmarshal(instanceSettings.JSONData, &settings); err != nil {
		return dataSourceSettings{}, err
	}
	return settings, nil
}

func schemaContains(fields []fieldSchema, name string) (fieldSchema, bool) {
	for _, sch := range fields {
		if sch.Name == name {
			return sch, true
		}
	}
	return fieldSchema{}, false
}

// detectTimeLayoutNaive attempts to parse the string from a set of layouts, and
// returns the first one that matched.
func detectTimeLayoutNaive(str string) (string, error) {
	layouts := []string{
		"2006-01-02",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05.999999",
		"2006-01-02 15:04:05.999999 -07:00",
		"2006-01-02 15:04:05.999999Z",
		"2006-01-02T15:04",
		"2006-01-02T15:04:05 MST",
		"2006-01-02T15:04:05.999999",
		"2006-01-02T15:04:05.999999 -07:00",
		"2006-01-02T15:04:05.999999Z",
		"2006/01/02",
		"2006/01/02 15:04",
		"2006/01/02 15:04:05 MST",
		"2006/01/02 15:04:05.999999",
		"2006/01/02 15:04:05.999999 -07:00",
		"2006/01/02 15:04:05.999999Z",
		"2006/01/02T15:04",
		"2006/01/02T15:04:05 MST",
		"2006/01/02T15:04:05.999999",
		"2006/01/02T15:04:05.999999 -07:00",
		"2006/01/02T15:04:05.999999Z",
		"01/02/2006",
		"01/02/2006 15:04",
		"01/02/2006 15:04:05 MST",
		"01/02/2006 15:04:05.999999",
		"01/02/2006 15:04:05.999999 -07:00",
		"01/02/2006 15:04:05.999999Z",
		"01/02/2006T15:04",
		"01/02/2006T15:04:05 MST",
		"01/02/2006T15:04:05.999999",
		"01/02/2006T15:04:05.999999 -07:00",
		"01/02/2006T15:04:05.999999Z",
	}

	for _, layout := range layouts {
		if _, err := time.Parse(layout, str); err == nil {
			return layout, nil
		}
	}

	return "", errors.New("unsupported time format")
}
