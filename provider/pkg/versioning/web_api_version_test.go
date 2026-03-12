package versioning

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWebAppTrackingVersion validates that the Web module's tracking version
// is at least 2024-11-01 to include the clientAffinityProxyEnabled property.
// See: https://github.com/pulumi/pulumi-azure-native/issues/4403
func TestWebAppTrackingVersion(t *testing.T) {
	repoRoot := findRepoRoot(t)
	specPath := filepath.Join(repoRoot, "versions", "v3-spec.yaml")

	spec, err := ReadSpec(specPath)
	require.NoError(t, err)

	webSpec, ok := spec["Web"]
	require.True(t, ok, "Web module should exist in v3-spec.yaml")
	require.NotNil(t, webSpec.Tracking, "Web module should have a tracking version")

	trackingVersion := string(*webSpec.Tracking)
	assert.GreaterOrEqual(t, trackingVersion, "2024-11-01",
		"Web module tracking version should be >= 2024-11-01 to include clientAffinityProxyEnabled (issue #4403); got %s", trackingVersion)
}

// TestWebAppSchemaHasClientAffinityProxyEnabled verifies the generated schema
// includes the clientAffinityProxyEnabled property on the WebApp resource.
// See: https://github.com/pulumi/pulumi-azure-native/issues/4403
func TestWebAppSchemaHasClientAffinityProxyEnabled(t *testing.T) {
	repoRoot := findRepoRoot(t)
	schemaPath := filepath.Join(repoRoot, "provider", "cmd", "pulumi-resource-azure-native", "schema.json")

	data, err := os.ReadFile(schemaPath)
	require.NoError(t, err, "should be able to read schema.json")

	var schema map[string]interface{}
	err = json.Unmarshal(data, &schema)
	require.NoError(t, err, "should be able to parse schema.json")

	resources, ok := schema["resources"].(map[string]interface{})
	require.True(t, ok, "schema should have resources")

	// Check the default (non-versioned) WebApp resource
	webApp, ok := resources["azure-native:web:WebApp"].(map[string]interface{})
	require.True(t, ok, "schema should have azure-native:web:WebApp resource")

	inputProperties, ok := webApp["inputProperties"].(map[string]interface{})
	require.True(t, ok, "WebApp should have inputProperties")

	_, hasClientAffinityProxyEnabled := inputProperties["clientAffinityProxyEnabled"]
	assert.True(t, hasClientAffinityProxyEnabled,
		"WebApp resource should have clientAffinityProxyEnabled input property (issue #4403)")

	properties, ok := webApp["properties"].(map[string]interface{})
	require.True(t, ok, "WebApp should have properties")

	_, hasClientAffinityProxyEnabledOutput := properties["clientAffinityProxyEnabled"]
	assert.True(t, hasClientAffinityProxyEnabledOutput,
		"WebApp resource should have clientAffinityProxyEnabled output property (issue #4403)")
}

func findRepoRoot(t *testing.T) string {
	t.Helper()
	// Walk up from current directory to find the repo root (has versions/ directory)
	dir, err := os.Getwd()
	require.NoError(t, err)

	for {
		if _, err := os.Stat(filepath.Join(dir, "versions", "v3-spec.yaml")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not find repository root")
		}
		dir = parent
	}
}
