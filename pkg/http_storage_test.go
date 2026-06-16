package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func TestHTTPStorage_Stat(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "time,value\n1609754451933,12")
		if err != nil {
			t.Fatal(err)
		}
	}))

	settings := backend.DataSourceInstanceSettings{
		URL:      srv.URL,
		JSONData: json.RawMessage(`{}`),
	}

	var qm queryModel
	storage, err := newHTTPStorage(qm, settings, &http.Client{})
	if err != nil {
		t.Fatal(err)
	}

	if err := storage.Stat(log.New()); err != nil {
		t.Fatal(err)
	}
}

func TestHTTPStorage_Open(t *testing.T) {
	csv := "time,value\n1609754451933,12"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, csv)
		if err != nil {
			t.Fatal(err)
		}
	}))

	settings := backend.DataSourceInstanceSettings{
		URL:      srv.URL,
		JSONData: json.RawMessage(`{}`),
	}

	var qm queryModel

	storage, err := newHTTPStorage(qm, settings, &http.Client{})
	if err != nil {
		t.Fatal(err)
	}

	r, err := storage.Open()
	if err != nil {
		t.Fatal(err)
	}

	got, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(got, []byte(csv)) {
		t.Fatalf("unexpected response: %s", got)
	}
}

func TestHTTPStorage_Settings(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") != "XYZ" {
			t.Error("missing custom header")
		}

		if r.URL.RawQuery != "limit=100&page=1" {
			t.Errorf("unexpected query string: %q", r.URL.RawQuery)
		}

		_, err := fmt.Fprintf(w, "time,value\n1609754451933,12")
		if err != nil {
			t.Fatal(err)
		}
	}))

	settings := backend.DataSourceInstanceSettings{
		URL: srv.URL,
		JSONData: json.RawMessage(`{
				"queryParams": "limit=100&page=1",
				"httpHeaderName1": "X-Api-Key"
			}`),
		DecryptedSecureJSONData: map[string]string{
			"httpHeaderValue1": "XYZ",
		},
	}
	var qm queryModel
	opts, err := settings.HTTPClientOptions(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	httpClient, err := httpclient.New(opts)
	if err != nil {
		t.Fatal(err)
	}
	storage, err := newHTTPStorage(qm, settings, httpClient)
	if err != nil {
		t.Fatal(err)
	}

	_, err = storage.Open()
	if err != nil {
		t.Error(err)
	}

	if err := storage.Stat(log.New()); err != nil {
		t.Error(err)
	}
}

func TestHTTPStorage_Options(t *testing.T) {
	opts := queryModel{
		Method:  "POST",
		Path:    "/orders",
		Params:  [][2]string{{"foo", "bar"}},
		Headers: [][2]string{{"baz", "test"}},
		Body:    `{"something": "anything"}`,
	}

	var invoked bool

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != opts.Method {
			t.Errorf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != opts.Path {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.RawQuery != "foo=bar" {
			t.Errorf("unexpected query: %s", r.URL.RawQuery)
		}

		for _, p := range opts.Headers {
			val := r.Header.Get(p[0])
			if val != "" {
				if !reflect.DeepEqual(val, p[1]) {
					t.Errorf("unexpected header value: %s", p[1])
				}
			} else {
				t.Errorf("missing header key: %s", p[0])
			}
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		if opts.Body != string(b) {
			t.Errorf("unexpected body: %s", b)
		}

		invoked = true
	}))

	settings := backend.DataSourceInstanceSettings{
		URL:      srv.URL,
		JSONData: json.RawMessage(`{"storage": "http"}`),
	}

	storage, err := newHTTPStorage(opts, settings, &http.Client{})
	if err != nil {
		t.Fatal(err)
	}

	if err := storage.Stat(log.New()); err != nil {
		t.Fatal(err)
	}

	if !invoked {
		t.Fatal("server didn't receive any requests")
	}
}

func TestHTTPStorage_MalformedParamsHeaders(t *testing.T) {
	tests := []struct {
		name    string
		params  [][2]string
		headers [][2]string
	}{
		{
			name:    "empty params array",
			params:  [][2]string{},
			headers: [][2]string{{"valid-header", "value"}},
		},
		{
			name:    "params with single element arrays",
			params:  [][2]string{{"incomplete"}},
			headers: [][2]string{{"valid-header", "value"}},
		},
		{
			name:    "headers with single element arrays",
			params:  [][2]string{{"valid-param", "value"}},
			headers: [][2]string{{"incomplete"}},
		},
		{
			name:    "mix of valid and invalid params/headers",
			params:  [][2]string{{"valid-param", "value"}, {"incomplete"}},
			headers: [][2]string{{"valid-header", "value"}, {"incomplete"}},
		},
		{
			name:    "completely empty arrays",
			params:  [][2]string{},
			headers: [][2]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify that only valid params/headers are set
				for _, p := range tt.params {
					if len(p) >= 2 {
						if r.URL.Query().Get(p[0]) != p[1] {
							t.Errorf("expected param %s=%s", p[0], p[1])
						}
					}
				}
				
				for _, h := range tt.headers {
					if len(h) >= 2 {
						if r.Header.Get(h[0]) != h[1] {
							t.Errorf("expected header %s=%s", h[0], h[1])
						}
					}
				}
				
				_, err := fmt.Fprintf(w, "time,value\n1609754451933,12")
				if err != nil {
					t.Fatal(err)
				}
			}))
			defer srv.Close()

			settings := backend.DataSourceInstanceSettings{
				URL:      srv.URL,
				JSONData: json.RawMessage(`{}`),
			}

			qm := queryModel{
				Params:  tt.params,
				Headers: tt.headers,
			}

			storage, err := newHTTPStorage(qm, settings, &http.Client{})
			if err != nil {
				t.Fatal(err)
			}

			// This should not panic even with malformed params/headers
			_, err = storage.Open()
			if err != nil {
				t.Fatal(err)
			}

			err = storage.Stat(log.New())
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestNewRequestFromQuery_EdgeCases(t *testing.T) {
	settings := backend.DataSourceInstanceSettings{
		URL:      "http://localhost:8000",
		JSONData: json.RawMessage(`{}`),
	}

	tests := []struct {
		name        string
		query       queryModel
		expectError bool
	}{
		{
			name: "empty params and headers",
			query: queryModel{
				Method:  "GET",
				Params:  [][2]string{},
				Headers: [][2]string{},
			},
			expectError: false,
		},
		{
			name: "malformed params",
			query: queryModel{
				Method:  "GET",
				Params:  [][2]string{{"incomplete"}},
				Headers: [][2]string{{"valid", "header"}},
			},
			expectError: false, // Should not error, just skip incomplete params
		},
		{
			name: "malformed headers",
			query: queryModel{
				Method:  "GET",
				Params:  [][2]string{{"valid", "param"}},
				Headers: [][2]string{{"incomplete"}},
			},
			expectError: false, // Should not error, just skip incomplete headers
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := newRequestFromQuery(settings, tt.query)
			
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if req == nil {
					t.Fatalf("expected request but got nil")
				}
			}
		})
	}
}

func TestHTTPStorage_UrlHandling(t *testing.T) {
	instanceSettings := backend.DataSourceInstanceSettings{
		URL:      "http://localhost:8000",
		JSONData: json.RawMessage(`{"storage": "http"}`),
	}
	badPath := "@example.com/test/1"
	goodPath := "/test/1"

	parsedSettingsURL, err := url.Parse(instanceSettings.URL)
	if err != nil {
		t.Fatal(err)
	}

	req1, err := newRequestFromQuery(instanceSettings, queryModel{
		Path:   goodPath,
		Method: "GET",
	})

	if err != nil {
		t.Fatal(err)
	}

	if req1.URL.Host != parsedSettingsURL.Host {
		t.Fatal("hostname changed because of the path")
	}

	_, err = newRequestFromQuery(instanceSettings, queryModel{
		Path:   badPath,
		Method: "GET",
	})

	// this must have failed
	if err == nil {
		t.Fatal("host-modifying path was accepted")
	}
}
