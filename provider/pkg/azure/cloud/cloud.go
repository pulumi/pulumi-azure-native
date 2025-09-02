// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloud

import (
	"context"
	"fmt"
	"strings"

	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure/cloud/metadata"
)

// Configuration extends azcore/cloud.Configuration with additional properties.
type Configuration struct {
	azcloud.Configuration
	Name      string
	Endpoints ConfigurationEndpoints
	Suffixes  ConfigurationSuffixes
}

type ConfigurationEndpoints struct {
	MicrosoftGraph string
}

type ConfigurationSuffixes struct {
	StorageEndpoint string
	KeyVaultDNS     string
}

var (
	// AzurePublic contains the settings for the public Azure cloud.
	AzurePublic Configuration
	// AzureChina contains the settings for the Azure China cloud.
	AzureChina Configuration
	// AzureGovernment contains the settings for the Azure Government cloud.
	AzureGovernment Configuration
)

func init() {
	// Azure Public Cloud
	AzurePublic = Configuration{
		Name:          "AzureCloud",
		Configuration: azcloud.AzurePublic,
		Endpoints: ConfigurationEndpoints{
			MicrosoftGraph: "https://graph.microsoft.com/",
		},
		Suffixes: ConfigurationSuffixes{
			StorageEndpoint: "core.windows.net",
			KeyVaultDNS:     ".vault.azure.net",
		},
	}

	// Azure China Cloud
	AzureChina = Configuration{
		Name:          "AzureChinaCloud",
		Configuration: azcloud.AzureChina,
		Endpoints: ConfigurationEndpoints{
			MicrosoftGraph: "https://microsoftgraph.chinacloudapi.cn",
		},
		Suffixes: ConfigurationSuffixes{
			StorageEndpoint: "core.chinacloudapi.cn",
			KeyVaultDNS:     ".vault.azure.cn",
		},
	}

	// Azure Government Cloud
	AzureGovernment = Configuration{
		Name:          "AzureUSGovernment",
		Configuration: azcloud.AzureGovernment,
		Endpoints: ConfigurationEndpoints{
			MicrosoftGraph: "https://graph.microsoft.us/",
		},
		Suffixes: ConfigurationSuffixes{
			StorageEndpoint: "core.usgovcloudapi.net",
			KeyVaultDNS:     ".vault.usgovcloudapi.net",
		},
	}
}

// FromName returns the azure-sdk-for-go/sdk/azcore/cloud configuration for the given cloud.
// Valid names are as documented in the provider's installation & configuration guide, currently
// public, china, usgovernment, or the empty value for public.
// If the cloud name is unknown, it falls back to AzurePublic and returns false.
func FromName(cloudName string) (Configuration, bool) {
	switch strings.ToLower(cloudName) {
	case "public", "azurecloud":
		return AzurePublic, true
	case "china", "azurechinacloud":
		return AzureChina, true
	case "usgov", "usgovernment", "azureusgovernment", "azureusgovernmentcloud", "usgovernmentl4":
		return AzureGovernment, true
	}
	return AzurePublic, false
}

// FromMetadataEndpoint returns the azure-sdk-for-go/sdk/azcore/cloud configuration from the given metadata endpoint.
func FromMetadataEndpoint(ctx context.Context, endpoint string) (Configuration, error) {
	if !strings.Contains(endpoint, "://") {
		endpoint = "https://" + endpoint
	}
	client := metadata.NewClientWithEndpoint(endpoint)

	metadata, err := client.GetMetaData(ctx)
	if err != nil {
		return Configuration{}, fmt.Errorf("failed to get Azure environment metadata from %s: %w", endpoint, err)
	}

	return Configuration{
		Name: metadata.Name, // note: if using CLI authentication, this name must match the CLI's active cloud.
		Configuration: azcloud.Configuration{
			ActiveDirectoryAuthorityHost: metadata.Authentication.LoginEndpoint,
			Services: map[azcloud.ServiceName]azcloud.ServiceConfiguration{
				azcloud.ResourceManager: {
					Audience: metadata.Authentication.Audiences[0],
					Endpoint: metadata.ResourceManagerEndpoint,
				},
			},
		},
		Endpoints: ConfigurationEndpoints{
			MicrosoftGraph: metadata.ResourceIdentifiers.MicrosoftGraph,
		},
		Suffixes: ConfigurationSuffixes{
			StorageEndpoint: metadata.DnsSuffixes.Storage,
			KeyVaultDNS:     metadata.DnsSuffixes.KeyVault,
		},
	}, nil
}
