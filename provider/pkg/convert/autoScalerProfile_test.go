package convert

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
)

// TestRewriteAutoScalerProfileKeys verifies outbound mapping camelCase -> kebab-case.
func TestRewriteAutoScalerProfileKeys(t *testing.T) {
	in := map[string]interface{}{
		"scanInterval":            "30s",
		"scaleDownDelayAfterAdd":  "5m",
		"skipNodesWithSystemPods": true,
		"unknownProp":             "keep",
	}
	out := rewriteAutoScalerProfileKeys(in)
	assert.Equal(t, "30s", out["scan-interval"])
	assert.Equal(t, "5m", out["scale-down-delay-after-add"])
	assert.Equal(t, true, out["skip-nodes-with-system-pods"])
	assert.Equal(t, "keep", out["unknownProp"], "unknown properties should pass through unchanged")
	// Original camelCase keys should be gone for mapped ones
	assert.NotContains(t, out, "scanInterval")
	assert.NotContains(t, out, "scaleDownDelayAfterAdd")
	assert.NotContains(t, out, "skipNodesWithSystemPods")
}

// TestRestoreAutoScalerProfileKeys verifies inbound mapping kebab-case -> camelCase.
func TestRestoreAutoScalerProfileKeys(t *testing.T) {
	in := map[string]interface{}{
		"scan-interval":               "30s",
		"scale-down-delay-after-add":  "5m",
		"skip-nodes-with-system-pods": true,
		"unknown-prop":                "keep",
	}
	out := restoreAutoScalerProfileKeys(in)
	assert.Equal(t, "30s", out["scanInterval"])
	assert.Equal(t, "5m", out["scaleDownDelayAfterAdd"])
	assert.Equal(t, true, out["skipNodesWithSystemPods"])
	assert.Equal(t, "keep", out["unknown-prop"], "unknown properties should pass through unchanged")
	// Original kebab-case keys should be gone for mapped ones
	assert.NotContains(t, out, "scan-interval")
	assert.NotContains(t, out, "scale-down-delay-after-add")
	assert.NotContains(t, out, "skip-nodes-with-system-pods")
}

// TestFullConversionPathOutbound ensures the rewrite is applied during SdkInputsToRequestBody when type suffix matches.
func TestFullConversionPathOutbound(t *testing.T) {
	// Define a type whose name ends with ManagedClusterPropertiesAutoScalerProfile to trigger the hack.
	types := map[string]resources.AzureAPIType{
		"azure-native:testing:ManagedClusterPropertiesAutoScalerProfile": {
			Properties: map[string]resources.AzureAPIProperty{
				// These property names are currently camelCase in generated schema (regression scenario)
				"scanInterval": {Type: "string"},
			},
		},
	}
	converter := NewSdkShapeConverterFull(types)
	props := map[string]resources.AzureAPIProperty{
		"profile": {Ref: "#/types/azure-native:testing:ManagedClusterPropertiesAutoScalerProfile"},
	}
	inputs := map[string]interface{}{
		"profile": map[string]interface{}{
			"scanInterval": "20s",
		},
	}
	body, err := converter.SdkInputsToRequestBody(props, inputs, "/subscriptions/abc/resourceGroups/rg/providers/Microsoft.ContainerService/managedClusters/test")
	assert.NoError(t, err)
	profile := body["profile"].(map[string]interface{})
	assert.Equal(t, "20s", profile["scan-interval"], "expected rewritten kebab-case key present")
	assert.NotContains(t, profile, "scanInterval", "camelCase key should be removed")
}

// TestFullConversionPathInbound ensures the restore is applied during ResponseBodyToSdkOutputs.
func TestFullConversionPathInbound(t *testing.T) {
	types := map[string]resources.AzureAPIType{
		"azure-native:testing:ManagedClusterPropertiesResponseAutoScalerProfile": {
			Properties: map[string]resources.AzureAPIProperty{
				// In generated schema these would appear camelCase, but inbound wire format is kebab-case; we test restoration.
				"scanInterval": {Type: "string"},
			},
		},
	}
	converter := NewSdkShapeConverterFull(types)
	props := map[string]resources.AzureAPIProperty{
		"profile": {Ref: "#/types/azure-native:testing:ManagedClusterPropertiesResponseAutoScalerProfile"},
	}
	response := map[string]interface{}{
		"profile": map[string]interface{}{
			"scan-interval": "15s",
		},
	}
	outputs := converter.ResponseBodyToSdkOutputs(props, response)
	profile := outputs["profile"].(map[string]interface{})
	assert.Equal(t, "15s", profile["scanInterval"], "expected restored camelCase key present")
	assert.NotContains(t, profile, "scan-interval", "kebab-case key should be removed")
}
