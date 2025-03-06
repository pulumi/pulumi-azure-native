// Copyright 2025, Pulumi Corporation.

package provider

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/require"
)

// A struct to get the schema version without deserializing the whole schema.
type schemaWithVersion struct {
	Version string `json:"version"`
}

func TestParameterize(t *testing.T) {
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

	keys := []string{}
	for k := range provider.parameterizedSchemas.schemas {
		keys = append(keys, k)
	}
	require.Len(t, keys, 1, strings.Join(keys, ", "))
	schema, ok := provider.parameterizedSchemas.get("azure-native_aad_v20221201", providerVersion)
	require.Truef(t, ok, "providerVersion=%s, keys=%s", providerVersion, strings.Join(keys, ", "))
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
