package main

import (
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func main() {
	logger := log.New()

	ds := newDataSource(logger)

	//nolint:staticcheck // let's ignore this for now until we rewrite the backend to plugin manage
	opts := datasource.ServeOpts{
		QueryDataHandler:   ds,
		CheckHealthHandler: ds,
	}

	//nolint:staticcheck // let's ignore this for now until we rewrite the backend to plugin manage
	if err := datasource.Serve(opts); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
