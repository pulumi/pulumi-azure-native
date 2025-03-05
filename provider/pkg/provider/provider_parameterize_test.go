// Copyright 2025, Pulumi Corporation.

package provider

import (
	"context"
	"os"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/require"
)

func TestParameterize(t *testing.T) {
	t.Parallel()

	providerVersion := version.GetVersion().String()

	schemaBytes, err := os.ReadFile("../../../bin/schema-full.json")
	require.NoError(t, err)
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
	schema, ok := provider.parameterizedSchemas.get("azure-native_aad_v20221201", providerVersion)
	require.True(t, ok)
	require.NotEmpty(t, schema)
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
