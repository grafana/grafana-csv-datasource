package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

	storage, err := newHTTPStorage(instance, logger)
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
		fmt.Fprintln(w, csv)
	}))

	instance := &dataSourceInstance{
		httpClient: &http.Client{},
		settings: backend.DataSourceInstanceSettings{
			URL:      srv.URL,
			JSONData: json.RawMessage(`{}`),
		},
	}

	logger := log.New()

	storage, err := newHTTPStorage(instance, logger)
	if err != nil {
		t.Fatal(err)
	}

	r, err := storage.Open()
	if err != nil {
		t.Fatal(err)
	}

	got, err := ioutil.ReadAll(r)
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

	storage, err := newHTTPStorage(instance, logger)
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
