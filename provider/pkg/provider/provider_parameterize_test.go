// Copyright 2025, Pulumi Corporation.

package provider

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetParameterizeArgs(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		args := parameterizeArgs{
			Module:  "aad",
			Version: "v20221201",
		}
		serialized, err := serializeParameterizeArgs(&args)
		require.NoError(t, err)

		req := &rpc.ParameterizeRequest{
			Parameters: &rpc.ParameterizeRequest_Value{
				Value: &rpc.ParameterizeRequest_ParametersValue{
					Value: serialized,
				},
			},
		}
		got, err := getParameterizeArgs(req)
		require.NoError(t, err)
		require.Equal(t, args, *got)
	})

	t.Run("args", func(t *testing.T) {
		t.Parallel()

		req := &rpc.ParameterizeRequest{
			Parameters: &rpc.ParameterizeRequest_Args{
				Args: &rpc.ParameterizeRequest_ParametersArgs{
					Args: []string{"aad", "v20221201"},
				},
			},
		}
		got, err := getParameterizeArgs(req)
		require.NoError(t, err)
		require.Equal(t, "aad", got.Module)
		require.Equal(t, "v20221201", got.Version)
	})

	t.Run("unexpected args", func(t *testing.T) {
		t.Parallel()

		req := &rpc.ParameterizeRequest{
			Parameters: &rpc.ParameterizeRequest_Args{
				Args: &rpc.ParameterizeRequest_ParametersArgs{
					Args: []string{"aad"},
				},
			},
		}
		got, err := getParameterizeArgs(req)
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		req := &rpc.ParameterizeRequest{}
		_, err := getParameterizeArgs(req)
		require.Error(t, err)
	})
}

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

	provider, err := makeProviderInternal(nil, "azure-native", providerVersion, schemaBytes, &resources.APIMetadata{
		Types: resources.GoMap[resources.AzureAPIType]{},
		Resources: resources.GoMap[resources.AzureAPIResource]{
			"azure-native:aad/v20221201:DomainService": {},
		},
		Invokes: resources.GoMap[resources.AzureAPIInvoke]{},
	})
	require.NoError(t, err)

	parameterizationArgs := []string{"aad", "v20221201"}
	expectedProviderName := "azure-native_aad_v20221201"

	require.NoError(t, err)

	resp, err := provider.Parameterize(context.Background(), &rpc.ParameterizeRequest{
		Parameters: &rpc.ParameterizeRequest_Args{
			Args: &rpc.ParameterizeRequest_ParametersArgs{
				Args: parameterizationArgs,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, expectedProviderName, resp.Name)

	// check that schema looks ok
	var schema pschema.PackageSpec
	err = json.Unmarshal(provider.schemaBytes, &schema)
	require.NoError(t, err)
	assert.NotEmpty(t, schema.Types)
	for typeName := range schema.Types {
		assert.True(t, strings.HasPrefix(typeName, expectedProviderName+":"))
	}
	assert.NotEmpty(t, schema.Resources)
	for resourceName := range schema.Resources {
		assert.True(t, strings.HasPrefix(resourceName, expectedProviderName+":"))
	}
	assert.NotEmpty(t, schema.Functions)

	// check that metadata looks ok
	metadata := provider.resourceMap
	assert.NotNil(t, metadata)
	assert.NotNil(t, metadata.Resources)
	assert.NotNil(t, metadata.Types)
	assert.NotNil(t, metadata.Invokes)
	resource, ok, err := metadata.Resources.Get(expectedProviderName + ":aad:DomainService")
	require.NoError(t, err)
	assert.True(t, ok)
	assert.NotNil(t, resource)
	_, ok, err = metadata.Resources.Get("azure-native:aad/v20221201:DomainService")
	require.NoError(t, err)
	assert.False(t, ok)

	// parameterization arguments a serialized as a base64-encoded JSON string
	assert.NotEmpty(t, schema.Parameterization.Parameter)
	args, err := deserializeParameterizeArgs(schema.Parameterization.Parameter)
	require.NoError(t, err)
	assert.Equal(t, "aad", args.Module)
	assert.Equal(t, "v20221201", args.Version)
}

func TestRoundtripParameterizeArgs(t *testing.T) {
	t.Parallel()

	args := parameterizeArgs{
		Module:  "aad",
		Version: "v20221201",
	}
	serialized, err := serializeParameterizeArgs(&args)
	require.NoError(t, err)

	deserialized, err := deserializeParameterizeArgs(serialized)
	require.NoError(t, err)

	require.Equal(t, args, *deserialized)
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

func TestParseApiVersion(t *testing.T) {
	t.Parallel()

	t.Run("Valid", func(t *testing.T) {
		t.Parallel()

		for _, v := range [][]string{
			{"2020-01-01", "v20200101"},
			{"2027-01-01", "v20270101"},
			{"2020-12-31", "v20201231"},
			{"v2020-12-31", "v20201231"},
			{"V2020-12-31", "v20201231"},
			{"v20230301", "v20230301"},
			{"V20230301", "v20230301"},
			{"20230301", "v20230301"},
			{" 20230301 ", "v20230301"},
			{"2023-03-01-preview", "v20230301preview"},
			{"2023-03-01-privatepreview ", "v20230301privatepreview"},
			{"v20230301preview", "v20230301preview"},
			{"v20230301privatepreview ", "v20230301privatepreview"},
		} {
			parsed, err := parseApiVersion(v[0])
			require.NoError(t, err)
			assert.Equal(t, v[1], parsed)
		}
	})

	t.Run("Invalid", func(t *testing.T) {
		t.Parallel()

		for _, v := range []string{
			"1999-01-01",
			"2027-011-01",
			"2020-12-311",
			"2020-12-41",
			"2020-22-31",
			"v2023030",
			"V202303011",
			"19990301",
			"2023-03-01-alpha",
		} {
			_, err := parseApiVersion(v)
			assert.Error(t, err, v)
		}
	})
}

func TestUpdateMetadataRefs(t *testing.T) {
	t.Parallel()

	t.Run("Empty metadata", func(t *testing.T) {
		t.Parallel()
		metadata := &resources.APIMetadata{}
		updated, err := updateMetadataRefs(metadata, "azure-native_storage_v20240101", "storage", "v20240101")
		require.NoError(t, err)
		require.Empty(t, updated.Resources)
		require.Empty(t, updated.Types)
		require.Empty(t, updated.Invokes)
	})

	t.Run("Updates refs in types", func(t *testing.T) {
		t.Parallel()
		metadata := &resources.APIMetadata{
			Types: resources.GoMap[resources.AzureAPIType]{
				"type1": {
					Properties: map[string]resources.AzureAPIProperty{
						"prop1": {
							Ref: "#/types/azure-native:storage/v20240101:StorageAccount",
						},
					},
				},
			},
		}

		updated, err := updateMetadataRefs(metadata, "azure-native_storage_v20240101", "storage", "v20240101")
		require.NoError(t, err)

		prop, ok, err := updated.Types.Get("type1")
		require.NoError(t, err)
		require.True(t, ok)
		assert.Equal(t, "#/types/azure-native_storage_v20240101:storage:StorageAccount", prop.Properties["prop1"].Ref)
	})

	t.Run("Updates refs in resources", func(t *testing.T) {
		t.Parallel()
		metadata := &resources.APIMetadata{
			Resources: resources.GoMap[resources.AzureAPIResource]{
				"resource1": {
					PutParameters: []resources.AzureAPIParameter{
						{
							Body: &resources.AzureAPIType{
								Properties: map[string]resources.AzureAPIProperty{
									"prop1": {
										Ref: "#/types/azure-native:storage/v20240101:StorageAccount",
									},
								},
							},
						},
					},
				},
			},
		}

		updated, err := updateMetadataRefs(metadata, "azure-native_storage_v20240101", "storage", "v20240101")
		require.NoError(t, err)

		resource, ok, err := updated.Resources.Get("resource1")
		require.NoError(t, err)
		require.True(t, ok)
		require.Len(t, resource.PutParameters, 1)
		assert.Equal(t, "#/types/azure-native_storage_v20240101:storage:StorageAccount", resource.PutParameters[0].Body.Properties["prop1"].Ref)
	})

	t.Run("Updates refs in invokes", func(t *testing.T) {
		t.Parallel()
		metadata := &resources.APIMetadata{
			Invokes: resources.GoMap[resources.AzureAPIInvoke]{
				"invoke1": {
					Response: map[string]resources.AzureAPIProperty{
						"prop1": {
							Ref: "#/types/azure-native:storage/v20240101:StorageAccount",
						},
					},
				},
			},
		}

		updated, err := updateMetadataRefs(metadata, "azure-native_storage_v20240101", "storage", "v20240101")
		require.NoError(t, err)

		invoke, ok, err := updated.Invokes.Get("invoke1")
		require.NoError(t, err)
		require.True(t, ok)
		assert.Equal(t, "#/types/azure-native_storage_v20240101:storage:StorageAccount", invoke.Response["prop1"].Ref)
	})

	t.Run("Does not update refs pointing to a different module", func(t *testing.T) {
		t.Parallel()
		metadata := &resources.APIMetadata{
			Resources: resources.GoMap[resources.AzureAPIResource]{
				"resource1": {
					PutParameters: []resources.AzureAPIParameter{
						{
							Body: &resources.AzureAPIType{
								Properties: map[string]resources.AzureAPIProperty{
									"prop1": {
										Ref: "#/types/azure-native:common:ManagedIdentity",
									},
								},
							},
						},
					},
				},
			},
		}

		updated, err := updateMetadataRefs(metadata, "azure-native_storage_v20240101", "storage", "v20240101")
		require.NoError(t, err)

		resource, ok, err := updated.Resources.Get("resource1")
		require.NoError(t, err)
		require.True(t, ok)
		require.Len(t, resource.PutParameters, 1)
		assert.Equal(t, "#/types/azure-native:common:ManagedIdentity", resource.PutParameters[0].Body.Properties["prop1"].Ref)
	})
}

// Ensure that we can run `pulumi package add` with a local provider binary and get a new SDK.
func TestParameterizePackageAdd(t *testing.T) {
	t.Parallel()

	pt := pulumitest.NewPulumiTest(t, filepath.Join("test-programs", "parameterize-storage"))
	pulumiPackageAdd(t, pt, "../../../bin/pulumi-resource-azure-native", "storage", "v20240101")

	sdkPath := filepath.Join(pt.WorkingDir(), "sdks", "azure-native_storage_v20240101")
	if _, err := os.Stat(sdkPath); os.IsNotExist(err) {
		t.Fatalf("generated SDK directory not found at path: %s", sdkPath)
	}
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
}
