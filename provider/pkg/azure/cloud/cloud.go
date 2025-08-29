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
	"strings"

	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
)

type EndpointName string

const (
	// Blob is the service name for Azure Blob Storage.
	Blob EndpointName = "blob"
	// MSGraph is the service name for Microsoft Graph.
	MSGraph EndpointName = "msGraph"
	// KeyVault is the service name for Azure Key Vault.
	KeyVault EndpointName = "keyVault"
)

const (
	// StorageEndpointSuffix is the configuration key for the storage endpoint suffix.
	StorageEndpointSuffix = "storageEndpoint"
	// KeyVaultDNSSuffix is the configuration key for the Key Vault DNS suffix.
	KeyVaultDNSSuffix = "keyVaultDNSSuffix"
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

// ByName returns the azure-sdk-for-go/sdk/azcore/cloud configuration for the given cloud.
// Valid names are as documented in the provider's installation & configuration guide, currently
// public, china, usgovernment, or the empty value for public.
// If the cloud name is unknown, it falls back to AzurePublic and returns false.
func ByName(cloudName string) (Configuration, bool) {
	switch strings.ToLower(cloudName) {
	case "", "public", "azurecloud":
		return AzurePublic, true
	case "china", "azurechinacloud":
		return AzureChina, true
	case "usgov", "usgovernment", "azureusgovernment", "azureusgovernmentcloud":
		return AzureGovernment, true
	}
	return AzurePublic, false
}
