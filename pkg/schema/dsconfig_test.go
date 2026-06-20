package schema_test

import (
	"flag"
	"testing"

	dsconfigSchema "github.com/grafana/dsconfig/schema"
	"github.com/grafana/grafana-csv-datasource/pkg/models"
	pluginSchema "github.com/grafana/grafana-csv-datasource/pkg/schema"
	"github.com/stretchr/testify/require"
)

// generateArtifacts is set by `go generate ./pkg/schema`, which runs this test
// package with the -generateArtifacts flag to (re)write the committed schema
// artifacts. When the flag is not set, TestGenerateArtifacts is skipped.
var generateArtifacts = flag.Bool("generateArtifacts", false, "write the schema artifacts to disk instead of running tests")

//go:generate go test -run TestGenerateArtifacts -generateArtifacts
func TestGenerateArtifacts(t *testing.T) {
	if !*generateArtifacts {
		t.Skip("run via `go generate ./...` to write schema artifacts")
	}
	err := dsconfigSchema.WriteArtifacts(pluginSchema.NewSDKSchema())
	require.NoError(t, err)
	t.Log("schema artifacts generated")
}

// TestSchemaConformance runs the plugin-agnostic schema guard rails defined in
// the dsconfig SDK against this plugin's schema.
func TestSchemaConformance(t *testing.T) {
	cfg, err := pluginSchema.DSConfigSchema()
	require.NoError(t, err)
	dsconfigSchema.RunConformanceTests(t, dsconfigSchema.Params{
		PluginID:          "marcusolsson-csv-datasource",
		DSConfigSchema:    cfg,
		PluginSchema:      pluginSchema.NewSDKSchema(),
		SettingsJSONModel: models.PluginSettings{},
		SecureKeys:        []string{"basicAuthPassword"},
	})
}
