package httpclient

import (
	"net/http"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func New(dsInfo *backend.DataSourceInstanceSettings, timeout time.Duration, logger log.Logger) (*http.Client, error) {
	transport, err := newTransport(dsInfo)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}, nil
}
