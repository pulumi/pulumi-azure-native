// Copyright 2025, Pulumi Corporation.

package provider

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A struct to get the schema version without deserializing the whole schema.
type schemaWithVersion struct {
	Version string `json:"version"`
}

func TestParameterizeCreatesSchemaAndMetadata(t *testing.T) {
	t.Parallel()

	schemaBytes, err := os.ReadFile("../../../bin/schema-full.json")
	require.NoError(t, err)

	// Get the schema version from the schema. In production, it would be the same than version.GetVersion(), but in
	// CI it's not - 2.0.0-alpha.0+dev vs. 2.90.0-alpha.1741284698+8268d88.
	var v schemaWithVersion
	err = json.Unmarshal(schemaBytes, &v)
	require.NoError(t, err)
	providerVersion := v.Version

	provider, err := makeProviderInternal(nil, "azure-native", providerVersion, schemaBytes, &resources.PartialAzureAPIMetadata{
		Types:     resources.PartialMap[resources.AzureAPIType]{},
		Resources: resources.PartialMap[resources.AzureAPIResource]{},
		Invokes:   resources.PartialMap[resources.AzureAPIInvoke]{},
	})
	require.NoError(t, err)

	resp, err := provider.Parameterize(context.Background(), &rpc.ParameterizeRequest{
		Parameters: &rpc.ParameterizeRequest_Args{
			Args: &rpc.ParameterizeRequest_ParametersArgs{
				Args: []string{"aad", "v20221201"},
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, "azure-native_aad_v20221201", resp.Name)

	var schema pschema.PackageSpec
	err = json.Unmarshal(provider.schemaBytes, &schema)
	require.NoError(t, err)
	assert.NotEmpty(t, schema.Types)
	assert.NotEmpty(t, schema.Resources)
	assert.NotEmpty(t, schema.Functions)

	// metadata is serialized as a base64-encoded JSON string
	assert.NotEmpty(t, schema.Parameterization.Parameter)
	_, err = deserializeMetadata(schema.Parameterization.Parameter)
	require.NoError(t, err)
}

func TestFilterTokens(t *testing.T) {
	t.Parallel()

	types := map[string]struct{}{
		"azure-native:aad:Application":                  {},
		"azure-native:aad/v20221201:Application":        {},
		"azure-native:aad/v20331201:Application":        {},
		"azure-native:storage/v20221201:StorageAccount": {},
	}
	names, err := filterTokens(types, "aad", "v20221201")
	require.NoError(t, err)
	require.Len(t, names, 1)
	require.Equal(t, "Application", names["azure-native:aad/v20221201:Application"])
}

// equalInPartialMap asserts that a key exists in a PartialMap and its value is equal to an expected value.
func equalInPartialMap[T any](t *testing.T, pm resources.PartialMap[T], key string, expected T) {
	actual, ok, err := pm.Get(key)
	require.NoError(t, err)
	require.True(t, ok)
	require.Equal(t, expected, actual)
}

func TestRoundtripMetadata(t *testing.T) {
	t.Parallel()

	metadata := &resources.AzureAPIMetadata{
		Types: map[string]resources.AzureAPIType{
			"azure-native:aad:Thing": {
				Properties: map[string]resources.AzureAPIProperty{
					"prop": {
						Type:       "string",
						Containers: []string{"foo"},
						ForceNew:   true,
					},
				},
			},
		},
		Resources: map[string]resources.AzureAPIResource{
			"azure-native:aad:Application": {
				PutParameters: []resources.AzureAPIParameter{
					{
						Name:       "prop",
						Location:   "body",
						IsRequired: true,
						Body: &resources.AzureAPIType{
							Properties: map[string]resources.AzureAPIProperty{
								"prop": {
									Type: "object",
									Ref:  "#/azure-native:aad:Thing",
								},
							},
						},
					},
				},
			},
		},
		Invokes: map[string]resources.AzureAPIInvoke{},
	}

	metadataBytes, err := serializeMetadata(metadata)
	require.NoError(t, err)
	require.NotNil(t, metadataBytes)

	deserializedMetadata, err := deserializeMetadata(metadataBytes)
	require.NoError(t, err)

	for typeName, typeDef := range metadata.Types {
		equalInPartialMap(t, deserializedMetadata.Types, typeName, typeDef)
	}

	for resourceName, resourceDef := range metadata.Resources {
		equalInPartialMap(t, deserializedMetadata.Resources, resourceName, resourceDef)
	}

	for invokeName, invokeDef := range metadata.Invokes {
		equalInPartialMap(t, deserializedMetadata.Invokes, invokeName, invokeDef)
	}
}

// Ensure that we can run `pulumi package add` with a local provider binary and get a new SDK.
func TestParameterizePackageAdd(t *testing.T) {
	t.Parallel()

	pt := pulumitest.NewPulumiTest(t, filepath.Join("test-programs", "parameterize-storage"))
	pulumiPackageAdd(t, pt, "../../../bin/pulumi-resource-azure-native", "storage", "v20240101")
}

func pulumiPackageAdd(
	t *testing.T,
	pt *pulumitest.PulumiTest,
	localProviderBinPath string,
	args ...string,
) {
	if _, err := os.Stat(localProviderBinPath); os.IsNotExist(err) {
		t.Fatalf("Provider binary not found at path: %s", localProviderBinPath)
	}
	absLocalProviderBinPath, err := filepath.Abs(localProviderBinPath)
	require.NoError(t, err)

	ctx := context.Background()
	allArgs := append([]string{"package", "add", absLocalProviderBinPath}, args...)
	stdout, stderr, exitCode, err := pt.CurrentStack().Workspace().PulumiCommand().Run(
		ctx,
		pt.WorkingDir(),
		nil, /* reader */
		nil, /* additionalOutput */
		nil, /* additionalErrorOutput */
		nil, /* additionalEnv */
		allArgs...,
	)
	if err != nil || exitCode != 0 {
		t.Fatalf("Failed to run pulumi package add\nExit code: %d\nError: %v\n%s\n%s",
			exitCode, err, stdout, stderr)
	}

	sdkPath := filepath.Join(pt.WorkingDir(), "sdks", "azure-native_storage_v20240101")
	if _, err := os.Stat(sdkPath); os.IsNotExist(err) {
		t.Fatalf("generated SDK directory not found at path: %s", sdkPath)
	}
}
