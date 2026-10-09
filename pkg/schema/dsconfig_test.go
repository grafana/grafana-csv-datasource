package schema_test

import (
	_ "embed"
	"testing"

	"github.com/grafana/dsconfig/schema"
	"github.com/grafana/grafana-csv-datasource/pkg/models"
)

//go:embed dsconfig.json
var configSchemaJSON []byte

// settingsJSONModel adds the keys the plugin SDK reads from jsonData.
type settingsJSONModel struct {
	models.PluginSettings
	KeepCookies       []string `json:"keepCookies"`
	Timeout           float64  `json:"timeout"`
	OAuthPassThru     bool     `json:"oauthPassThru"`
	TLSAuth           bool     `json:"tlsAuth"`
	TLSAuthWithCACert bool     `json:"tlsAuthWithCACert"`
	TLSSkipVerify     bool     `json:"tlsSkipVerify"`
	ServerName        string   `json:"serverName"`
}

//go:generate go test -run TestPlugin -generateArtifacts
func TestPlugin(t *testing.T) {
	schema.RunPluginTests(t, schema.PluginUnderTest{
		ID:                "marcusolsson-csv-datasource",
		ConfigSchemaJSON:  configSchemaJSON,
		SettingsJSONModel: settingsJSONModel{},
		SecureKeys:        []string{"basicAuthPassword", "tlsCACert", "tlsClientCert", "tlsClientKey"},
	})
}
