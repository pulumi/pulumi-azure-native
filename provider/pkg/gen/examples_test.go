// Disabling this

package gen

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/pulumi/pulumi-azurerm/provider/pkg/provider"
	"github.com/stretchr/testify/require"
)

func TestFlattenInput(t *testing.T) {
	var metadata provider.AzureAPIMetadata
	// TODO - Requires `make generate_schema` to be run first
	// turn this into a proper unit test instead
	f, err := os.Open("../../cmd/pulumi-resource-azurerm/metadata.json")
	require.NoError(t, err)
	require.NoError(t, json.NewDecoder(f).Decode(&metadata))
	f.Close()

	for _, test := range []struct {
		name         string
		input        map[string]interface{}
		resourceName string
		expected     map[string]interface{}
		err          error
	}{
		{
			name:         "ContainersInBody",
			resourceName: "azurerm:compute/v20200601:VirtualMachine",
			input: map[string]interface{}{
				"parameters": map[string]interface{}{
					"resourceGroupName": "myResourceGroup",
					"vmName":            "myVM",
					"parameters": map[string]interface{}{
						"location": "westus",
						"properties": map[string]interface{}{
							"hardwareProfile": map[string]interface{}{
								"vmSize": "Standard_D1_v2",
							},
							"storageProfile": map[string]interface{}{
								"imageReference": map[string]interface{}{
									"sku":       "2016-Datacenter",
									"publisher": "MicrosoftWindowsServer",
									"version":   "latest",
									"offer":     "WindowsServer",
								},
								"osDisk": map[string]interface{}{
									"caching": "ReadWrite",
									"managedDisk": map[string]interface{}{
										"storageAccountType": "Standard_LRS",
									},
									"name":         "myVMosdisk",
									"createOption": "FromImage",
								},
							},
							"networkProfile": map[string]interface{}{
								"networkInterfaces": []map[string]interface{}{{
									"id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
									"properties": map[string]interface{}{
										"primary": true,
									},
								}},
							},
							"osProfile": map[string]interface{}{
								"adminUsername": "{your-username}",
								"computerName":  "myVM",
								"adminPassword": "{your-password}",
							},
							"availabilitySet": map[string]interface{}{
								"id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/{existing-availability-set-name}",
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"vmName":            "myVM",
				"resourceGroupName": "myResourceGroup",
				"availabilitySet": map[string]interface{}{
					"id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/availabilitySets/{existing-availability-set-name}",
				},
				"hardwareProfile": map[string]interface{}{
					"vmSize": "Standard_D1_v2",
				},
				"networkProfile": map[string]interface{}{
					"networkInterfaces": []map[string]interface{}{
						{
							"id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
							//"primary": true,
						},
					},
				},
				"location": "westus",
				"osProfile": map[string]interface{}{
					"adminPassword": "{your-password}",
					"adminUsername": "{your-username}",
					"computerName":  "myVM",
				},
				"storageProfile": map[string]interface{}{
					"imageReference": map[string]interface{}{
						"offer":     "WindowsServer",
						"publisher": "MicrosoftWindowsServer",
						"sku":       "2016-Datacenter",
						"version":   "latest",
					},
					"osDisk": map[string]interface{}{
						"caching":      "ReadWrite",
						"createOption": "FromImage",
						"managedDisk": map[string]interface{}{
							"storageAccountType": "Standard_LRS",
						},
						"name": "myVMosdisk",
					},
				},
			},
		},
		{
			name: "Props",
			input: map[string]interface{}{
				"parameters": map[string]interface{}{
					"subscriptionId":    "subscription-id",
					"resourceGroupName": "OneResourceGroupName",
					"api-version":       "2020-06-02",
					"resourceName":      "samplebotname",
					"connectionName":    "sampleConnection",
					"parameters": map[string]interface{}{
						"location": "West US",
						"etag":     "etag1",
						"properties": map[string]interface{}{
							"properties": map[string]interface{}{
								"clientId":          "sampleclientid",
								"clientSecret":      "samplesecret",
								"scopes":            "samplescope",
								"serviceProviderId": "serviceproviderid",
								"parameters": []map[string]interface{}{
									{
										"key":   "key1",
										"value": "value1",
									},
									{
										"key":   "key2",
										"value": "value2",
									},
								},
							},
						},
					},
				},
			},
			resourceName: "azurerm:botservice/v20200602:BotConnection",
			expected: map[string]interface{}{
				"resourceGroupName": "OneResourceGroupName",
				"resourceName":      "samplebotname",
				"connectionName":    "sampleConnection",
				"location":          "West US",
				"etag":              "etag1",
				"properties": map[string]interface{}{
					"clientId":          "sampleclientid",
					"clientSecret":      "samplesecret",
					"scopes":            "samplescope",
					"serviceProviderId": "serviceproviderid",
					"parameters": []map[string]interface{}{
						{
							"key":   "key1",
							"value": "value1",
						},
						{
							"key":   "key2",
							"value": "value2",
						},
					},
				},
			},
		},
		{
			name: "MoreProps",
			input: map[string]interface{}{
				"parameters": map[string]interface{}{
					"serviceName":       "apimService1",
					"resourceGroupName": "rg1",
					"api-version":       "2017-03-01",
					"subscriptionId":    "subid",
					"backendid":         "proxybackend",
					"parameters": map[string]interface{}{
						"properties": map[string]interface{}{
							"description": "description5308",
							"url":         "https://backendname2644/",
							"protocol":    "http",
							"tls": map[string]interface{}{
								"validateCertificateChain": true,
								"validateCertificateName":  true,
							},
							"proxy": map[string]interface{}{
								"url":      "http://192.168.1.1:8080",
								"username": "Contoso\\admin",
								"password": "opensesame",
							},
							"credentials": map[string]interface{}{
								"query": map[string]interface{}{
									"sv": []string{
										"xx",
										"bb",
										"cc",
									},
								},
								"header": map[string]interface{}{
									"x-my-1": []string{
										"val1",
										"val2",
									},
								},
								"authorization": map[string]interface{}{
									"scheme":    "Basic",
									"parameter": "opensesma",
								},
							},
						},
					},
				},
			},
			resourceName: "azurerm:apimanagement/v20170301:Backend",
			expected: map[string]interface{}{
				"serviceName":       "apimService1",
				"resourceGroupName": "rg1",
				"backendid":         "proxybackend",
				"description":       "description5308",
				"url":               "https://backendname2644/",
				"protocol":          "http",
				"tls": map[string]interface{}{
					"validateCertificateChain": true,
					"validateCertificateName":  true,
				},
				"proxy": map[string]interface{}{
					"url":      "http://192.168.1.1:8080",
					"username": "Contoso\\admin",
					"password": "opensesame",
				},
				"credentials": map[string]interface{}{
					"query": map[string]interface{}{
						"sv": []string{
							"xx",
							"bb",
							"cc",
						},
					},
					"header": map[string]interface{}{
						"x-my-1": []string{
							"val1",
							"val2",
						},
					},
					"authorization": map[string]interface{}{
						"scheme":    "Basic",
						"parameter": "opensesma",
					},
				},
			},
		},
		{
			name: "ServiceFabric",
			input: map[string]interface{}{
				"parameters": map[string]interface{}{
					"serviceName":       "apimService1",
					"resourceGroupName": "rg1",
					"api-version":       "2017-03-01",
					"subscriptionId":    "subid",
					"backendid":         "sfbackend",
					"parameters": map[string]interface{}{
						"properties": map[string]interface{}{
							"description": "Service Fabric Test App 1",
							"protocol":    "http",
							"url":         "fabric:/mytestapp/mytestservice",
							"properties": map[string]interface{}{
								"serviceFabricCluster": map[string]interface{}{
									"managementEndpoints": []interface{}{
										"https://somecluster.com",
									},
									"clientCertificatethumbprint": "EBA029198AA3E76EF0D70482626E5BCF148594A6",
									"serverX509Names": []map[string]interface{}{
										{
											"name":                        "ServerCommonName1",
											"issuerCertificateThumbprint": "IssuerCertificateThumbprint1",
										},
									},
									"maxPartitionResolutionRetries": 5,
								},
							},
						},
					},
				},
			},
			resourceName: "azurerm:apimanagement/v20170301:Backend",
			expected: map[string]interface{}{
				"serviceName":       "apimService1",
				"resourceGroupName": "rg1",
				"backendid":         "sfbackend",
				"description":       "Service Fabric Test App 1",
				"protocol":          "http",
				"url":               "fabric:/mytestapp/mytestservice",
				"properties": map[string]interface{}{
					"serviceFabricCluster": map[string]interface{}{
						"managementEndpoints": []interface{}{
							"https://somecluster.com",
						},
						"clientCertificatethumbprint": "EBA029198AA3E76EF0D70482626E5BCF148594A6",
						"serverX509Names": []map[string]interface{}{
							{
								"name":                        "ServerCommonName1",
								"issuerCertificateThumbprint": "IssuerCertificateThumbprint1",
							},
						},
						"maxPartitionResolutionRetries": 5,
					},
				},
			},
		},
	}[3:] {
		t.Run(test.name, func(t *testing.T) {
			in := test.input["parameters"].(map[string]interface{})
			params := map[string]provider.AzureAPIParameter{}
			for _, param := range metadata.Resources[test.resourceName].PutParameters {
				params[param.Name] = param
			}
			out, err := flattenInput(in, params, metadata.Types)
			if test.err != nil {
				require.Error(t, err)
				require.Empty(t, out)
			} else {
				require.NoError(t, err)
				expected, actual := &strings.Builder{}, &strings.Builder{}
				require.NoError(t, json.NewEncoder(expected).Encode(test.expected))
				require.NoError(t, json.NewEncoder(actual).Encode(out))
				require.JSONEq(t, expected.String(), actual.String())
			}
		})
	}
}
