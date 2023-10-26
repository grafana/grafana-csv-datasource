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
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func TestHTTPStorage_Stat(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "time,value\n1609754451933,12")
	}))

	instance := &dataSourceInstance{
		httpClient: &http.Client{},
		settings: backend.DataSourceInstanceSettings{
			URL:      srv.URL,
			JSONData: json.RawMessage(`{}`),
		},
	}

	logger := log.New()

	var opts dataSourceQuery
	storage, err := newHTTPStorage(instance, opts, logger)
	if err != nil {
		t.Fatal(err)
	}

	if err := storage.Stat(); err != nil {
		t.Fatal(err)
	}
}

func TestHTTPStorage_Open(t *testing.T) {
	csv := "time,value\n1609754451933,12"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, csv)
	}))

	instance := &dataSourceInstance{
		httpClient: &http.Client{},
		settings: backend.DataSourceInstanceSettings{
			URL:      srv.URL,
			JSONData: json.RawMessage(`{}`),
		},
	}

	logger := log.New()

	var opts dataSourceQuery
	storage, err := newHTTPStorage(instance, opts, logger)
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

		fmt.Fprintf(w, "time,value\n1609754451933,12")
	}))

	instance := &dataSourceInstance{
		httpClient: &http.Client{},
		settings: backend.DataSourceInstanceSettings{
			URL: srv.URL,
			JSONData: json.RawMessage(`{
				"queryParams": "limit=100&page=1",
				"httpHeaderName1": "X-Api-Key"
			}`),
			DecryptedSecureJSONData: map[string]string{
				"httpHeaderValue1": "XYZ",
			},
		},
	}

	logger := log.New()

	var opts dataSourceQuery
	storage, err := newHTTPStorage(instance, opts, logger)
	if err != nil {
		t.Fatal(err)
	}

	_, err = storage.Open()
	if err != nil {
		t.Error(err)
	}

	if err := storage.Stat(); err != nil {
		t.Error(err)
	}
}

func TestHTTPStorage_Options(t *testing.T) {
	opts := dataSourceQuery{
		Method:  "POST",
		Path:    "/orders",
		Params:  [][2]string{[2]string{"foo", "bar"}},
		Headers: [][2]string{[2]string{"baz", "test"}},
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

	instance := &dataSourceInstance{
		httpClient: &http.Client{},
		settings: backend.DataSourceInstanceSettings{
			URL:      srv.URL,
			JSONData: json.RawMessage(`{}`),
		},
	}

	logger := log.New()

	storage, err := newHTTPStorage(instance, opts, logger)
	if err != nil {
		t.Fatal(err)
	}

	if err := storage.Stat(); err != nil {
		t.Fatal(err)
	}

	if !invoked {
		t.Fatal("server didn't receive any requests")
	}
}

func TestHTTPStorage_UrlHandling(t *testing.T) {
	instanceSettings := backend.DataSourceInstanceSettings{
		URL: "http://localhost:8000",
	}
	badPath := "@example.com/test/1"
	goodPath := "/test/1"

	parsedSettingsURL, err := url.Parse(instanceSettings.URL)
	if err != nil {
		t.Fatal(err)
	}

	req1, err := newRequestFromQuery(&instanceSettings, dataSourceSettings{}, dataSourceQuery{
		Path:   goodPath,
		Method: "GET",
	})

	if err != nil {
		t.Fatal(err)
	}

	if req1.URL.Host != parsedSettingsURL.Host {
		t.Fatal("hostname changed because of the path")
	}

	_, err = newRequestFromQuery(&instanceSettings, dataSourceSettings{}, dataSourceQuery{
		Path:   badPath,
		Method: "GET",
	})

	// this must have failed
	if err == nil {
		t.Fatal("host-modifying path was accepted")
	}
}
