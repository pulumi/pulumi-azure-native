// Copyright 2016-2026, Pulumi Corporation.

package provider

import (
	"encoding/json"
	"os"
	"testing"

	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure/cloud"
)

// The fixtures in testdata/az-cloud-show are `az cloud show --name <cloud> -o json` outputs: azurecloud.json is
// recorded from Azure CLI 2.90.0, azurestackuser.json is the same document with an Azure Stack Hub's endpoints.
func TestCloudConfigurationFromAzureCLI(t *testing.T) {
	t.Run("the public cloud reads as the built-in public cloud configuration", func(t *testing.T) {
		output, err := os.ReadFile("testdata/az-cloud-show/azurecloud.json")
		require.NoError(t, err)

		actual, err := cloudConfigurationFromAzureCLI(output)

		require.NoError(t, err)
		assert.Equal(t, cloud.AzurePublic, *actual)
	})

	t.Run("a custom cloud reads its endpoints and suffixes", func(t *testing.T) {
		output, err := os.ReadFile("testdata/az-cloud-show/azurestackuser.json")
		require.NoError(t, err)

		actual, err := cloudConfigurationFromAzureCLI(output)

		require.NoError(t, err)
		assert.Equal(t, cloud.Configuration{
			Name: "AzureStackUser",
			Configuration: azcloud.Configuration{
				ActiveDirectoryAuthorityHost: "https://login.microsoftonline.com/",
				Services: map[azcloud.ServiceName]azcloud.ServiceConfiguration{
					azcloud.ResourceManager: {
						Audience: "https://management.contoso.onmicrosoft.com/71fa64d0-6c18-4fef-9b3a-0a4c6e8f3c1d",
						Endpoint: "https://management.local.azurestack.external",
					},
				},
			},
			Endpoints: cloud.ConfigurationEndpoints{
				MicrosoftGraph: "https://graph.microsoft.com/",
			},
			Suffixes: cloud.ConfigurationSuffixes{
				StorageEndpoint: "local.azurestack.external",
				KeyVaultDNS:     ".vault.local.azurestack.external",
			},
		}, *actual)
	})

	t.Run("a cloud without the endpoints authentication needs is rejected", func(t *testing.T) {
		for _, missing := range []string{"activeDirectory", "activeDirectoryResourceId", "resourceManager"} {
			endpoints := map[string]string{
				"activeDirectory":           "https://login.example.com",
				"activeDirectoryResourceId": "https://management.example.com/",
				"resourceManager":           "https://management.example.com/",
			}
			delete(endpoints, missing)
			output, err := json.Marshal(map[string]any{"name": "Incomplete", "endpoints": endpoints})
			require.NoError(t, err)

			_, err = cloudConfigurationFromAzureCLI(output)

			// Quoted, so that "activeDirectory" does not match an error about "activeDirectoryResourceId".
			assert.ErrorContains(t, err, `"`+missing+`"`)
		}
	})

	t.Run("output that is not a cloud is rejected", func(t *testing.T) {
		_, err := cloudConfigurationFromAzureCLI([]byte("ERROR: The cloud 'Unknown' is not registered."))

		assert.Error(t, err)
	})
}
