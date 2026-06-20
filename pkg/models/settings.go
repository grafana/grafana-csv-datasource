package models

// PluginSettings models the jsonData fields stored for a CSV datasource
// instance. It is the single source of truth for the datasource settings shape
// and is referenced by the schema conformance tests.
type PluginSettings struct {
	Storage     string `json:"storage"`
	QueryParams string `json:"queryParams"`
}
