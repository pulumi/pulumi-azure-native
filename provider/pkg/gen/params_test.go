// Copyright 2021, Pulumi Corporation.  All rights reserved.

package gen

import (
	"fmt"
	"github.com/segmentio/encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/resources"
	"github.com/stretchr/testify/require"
)

func TestFlattenParams(t *testing.T) {
	var metadata resources.AzureAPIMetadata
	// TODO - Requires `make generate_schema` to be run first
	// turn this into a proper unit test instead
	f, err := os.Open("../../cmd/pulumi-resource-azure-native/metadata.json")
	require.NoError(t, err)
	require.NoError(t, json.NewDecoder(f).Decode(&metadata))
	f.Close()

	resourceID := func(s string) string {
		return fmt.Sprintf("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/%s", s)
	}

	for _, test := range []struct {
		name         string
		inputFunc    func(t *testing.T) map[string]interface{}
		input        map[string]interface{}
		resourceName string
		expected     map[string]interface{}
		err          error
	}{
		{
			name:         "ContainersInBody",
			resourceName: "azure-native:compute/v20200601:VirtualMachine",
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
									"id": resourceID("Microsoft.Network/networkInterfaces/{existing-nic-name}"),
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
								"id": resourceID("Microsoft.Compute/availabilitySets/{existing-availability-set-name}"),
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"vmName":            "myVM",
				"resourceGroupName": "myResourceGroup",
				"availabilitySet": map[string]interface{}{
					"id": resourceID("Microsoft.Compute/availabilitySets/{existing-availability-set-name}"),
				},
				"hardwareProfile": map[string]interface{}{
					"vmSize": "Standard_D1_v2",
				},
				"networkProfile": map[string]interface{}{
					"networkInterfaces": []map[string]interface{}{
						{
							"id":      resourceID("Microsoft.Network/networkInterfaces/{existing-nic-name}"),
							"primary": true,
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
			resourceName: "azure-native:botservice/v20200602:BotConnection",
			expected: map[string]interface{}{
				"resourceGroupName": "OneResourceGroupName",
				"resourceName":      "samplebotname",
				"connectionName":    "sampleConnection",
				"location":          "West US",
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
			resourceName: "azure-native:apimanagement/v20170301:Backend",
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
			resourceName: "azure-native:apimanagement/v20170301:Backend",
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
		{
			name: "DeepNesting",
			input: map[string]interface{}{
				"parameters": map[string]interface{}{
					"networkSecurityGroupName": "rancher-security-group",
					"location":                 "westus2",
					"parameters": map[string]interface{}{
						"properties": map[string]interface{}{
							"securityRules": []map[string]interface{}{
								{
									"name": "SSH",
									"properties": map[string]interface{}{
										"description":              "SSH",
										"protocol":                 "*",
										"sourcePortRange":          "*",
										"destinationPortRange":     "22",
										"sourceAddressPrefix":      "*",
										"destinationAddressPrefix": "*",
										"access":                   "Allow",
										"priority":                 100,
										"direction":                "Inbound",
										"sourcePortRanges": []interface{}{
											"22",
											"80",
										},
										"destinationPortRanges": []interface{}{
											"22",
											"80",
										},
										"sourceAddressPrefixes": []interface{}{
											"192.168.0.1",
										},
										"destinationAddressPrefixes": []interface{}{
											"192.168.0.1",
										},
									},
								},
							},
						},
					},
				},
			},
			resourceName: "azure-native:network/v20200501:NetworkSecurityGroup",
			expected: map[string]interface{}{
				"networkSecurityGroupName": "rancher-security-group",
				"securityRules": []interface{}{
					map[string]interface{}{
						"access":                     "Allow",
						"description":                "SSH",
						"destinationAddressPrefix":   "*",
						"destinationAddressPrefixes": []interface{}{"192.168.0.1"},
						"destinationPortRange":       "22",
						"destinationPortRanges":      []interface{}{"22", "80"},
						"direction":                  "Inbound",
						"name":                       "SSH",
						"priority":                   100,
						"protocol":                   "*",
						"sourceAddressPrefix":        "*",
						"sourceAddressPrefixes":      []interface{}{"192.168.0.1"},
						"sourcePortRange":            "*",
						"sourcePortRanges":           []interface{}{"22", "80"},
					},
				},
			},
		},
		{
			name:         "NestedObject",
			inputFunc:    serialize(npe),
			resourceName: "azure-native:automation/v20170515preview:SoftwareUpdateConfigurationByName",
			expected: map[string]interface{}{
				"automationAccountName": "myaccount",
				"resourceGroupName":     "mygroup",
				"scheduleInfo": map[string]interface{}{
					"advancedSchedule": map[string]interface{}{
						"weekDays": []interface{}{
							"Monday",
							"Thursday",
						},
					},
					"expiryTime": "2018-11-09T11:22:57+00:00",
					"frequency":  "Hour",
					"interval":   1,
					"startTime":  "2017-10-19T12:22:57+00:00",
					"timeZone":   "America/Los_Angeles",
				},
				"softwareUpdateConfigurationName": "testpatch",
				"tasks": map[string]interface{}{
					"postTask": map[string]interface{}{
						"source": "GetCache",
					},
					"preTask": map[string]interface{}{
						"parameters": map[string]interface{}{
							"COMPUTERNAME": "Computer1",
						},
						"source": "HelloWorld",
					},
				},
				"updateConfiguration": map[string]interface{}{
					"azureVirtualMachines": []interface{}{
						"/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01",
						"/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02",
						"/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03",
					},
					"duration": "PT2H0M",
					"nonAzureComputerNames": []interface{}{
						"box1.contoso.com",
						"box2.contoso.com",
					},
					"operatingSystem": "Windows",
					"targets": map[string]interface{}{
						"azureQueries": []interface{}{
							map[string]interface{}{
								"locations": []interface{}{
									"Japan East",
									"UK South",
								},
								"scope": []interface{}{
									"/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources",
									"/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067",
								}, "tagSettings": map[string]interface{}{
									"filterOperator": "All",
									"tags": []interface{}{
										map[string]interface{}{
											"tag1": []interface{}{
												"tag1Value1",
												"tag1Value2",
												"tag1Value3",
											},
										},
										map[string]interface{}{
											"tag2": []interface{}{
												"tag2Value1",
												"tag2Value2",
												"tag2Value3",
											},
										},
									},
								},
							},
						},
						"nonAzureQueries": []interface{}{
							map[string]interface{}{
								"functionAlias": "SavedSearch1",
								"workspaceId":   "WorkspaceId1",
							},
							map[string]interface{}{
								"functionAlias": "SavedSearch2",
								"workspaceId":   "WorkspaceId2",
							},
						},
					}, "windows": map[string]interface{}{
						"excludedKbNumbers": []interface{}{
							"168934",
							"168973",
						},
						"includedUpdateClassifications": "Critical",
						"rebootSetting":                 "IfRequired",
					},
				},
			},
		},
		{
			name:         "VMScaleSet",
			inputFunc:    serialize(vmScaleSet),
			resourceName: "azure-native:compute:VirtualMachineScaleSet",
			expected: map[string]interface{}{
				"location":          "westus",
				"overprovision":     true,
				"resourceGroupName": "myResourceGroup",
				"sku": map[string]interface{}{
					"capacity": 3,
					"name":     "Standard_D1_v2",
					"tier":     "Standard",
				},
				"upgradePolicy": map[string]interface{}{"mode": "Manual"},
				"virtualMachineProfile": map[string]interface{}{
					"networkProfile": map[string]interface{}{
						"networkInterfaceConfigurations": []interface{}{
							map[string]interface{}{
								"enableIPForwarding": true,
								"ipConfigurations": []interface{}{
									map[string]interface{}{
										"name": "{vmss-name}",
										"subnet": map[string]interface{}{
											"id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}",
										},
									},
								},
								"name":    "{vmss-name}",
								"primary": true,
							},
						},
					},
					"osProfile": map[string]interface{}{
						"adminPassword":      "{your-password}",
						"adminUsername":      "{your-username}",
						"computerNamePrefix": "{vmss-name}",
					},
					"storageProfile": map[string]interface{}{
						"imageReference": map[string]interface{}{
							"id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}",
						},
						"osDisk": map[string]interface{}{
							"caching":      "ReadWrite",
							"createOption": "FromImage",
							"managedDisk": map[string]interface{}{
								"storageAccountType": "Standard_LRS",
							},
						},
					},
				},
				"vmScaleSetName": "{vmss-name}",
			},
		},
		{
			name: "NoMissingFields",
			input: map[string]interface{}{
				"parameters": map[string]interface{}{
					"appResource": map[string]interface{}{
						"properties": map[string]interface{}{
							"public":               true,
							"activeDeploymentName": "mydeployment1",
							"fqdn":                 "myapp.mydomain.com",
							"httpsOnly":            false,
							"temporaryDisk": map[string]interface{}{
								"sizeInGB":  2,
								"mountPath": "mytemporarydisk",
							},
							"persistentDisk": map[string]interface{}{
								"sizeInGB":  2,
								"mountPath": "mypersistentdisk",
							},
						},
						"identity": nil,
						"location": "eastus",
					},
					"api-version":       "2020-07-01",
					"subscriptionId":    "00000000-0000-0000-0000-000000000000",
					"resourceGroupName": "myResourceGroup",
					"serviceName":       "myservice",
					"appName":           "myapp",
				},
			},
			resourceName: "azure-native:appplatform:App",
			expected: map[string]interface{}{
				"appName":  "myapp",
				"location": "eastus",
				"properties": map[string]interface{}{
					"activeDeploymentName": "mydeployment1",
					"fqdn":                 "myapp.mydomain.com",
					"httpsOnly":            false,
					"persistentDisk": map[string]interface{}{
						"mountPath": "mypersistentdisk",
						"sizeInGB":  2,
					},
					"public": true,
					"temporaryDisk": map[string]interface{}{
						"mountPath": "mytemporarydisk",
						"sizeInGB":  2,
					},
				},
				"resourceGroupName": "myResourceGroup",
				"serviceName":       "myservice",
			},
		},
		{
			name: "ACI",
			inputFunc: serialize(`
{
  "parameters":
    {"containerGroup":
      { "properties": {
        "containers": [
          {
            "name": "[parameters('name')]",
            "properties": {
              "image": "[parameters('image')]",
              "ports": [
                {
                  "port": "[parameters('port')]"
                }
              ],
              "resources": {
                "requests": {
                  "cpu": "[parameters('cpuCores')]",
                  "memoryInGB": "[parameters('memoryInGb')]"
                }
              }
            }
          }
        ],
        "osType": "Linux",
        "restartPolicy": "[parameters('restartPolicy')]",
        "ipAddress": {
          "type": "Public",
          "ports": [
            {
              "protocol": "Tcp",
              "port": "[parameters('port')]"
            }
          ]
        }
      }
    }
  }
}`),
			resourceName: "azure-native:containerinstance:ContainerGroup",
			expected: map[string]interface{}{
				"containers": []interface{}{
					map[string]interface{}{
						"image": "[parameters('image')]",
						"name":  "[parameters('name')]",
						"ports": []interface{}{
							map[string]interface{}{
								"port": "[parameters('port')]",
							},
						},
						"resources": map[string]interface{}{
							"requests": map[string]interface{}{
								"cpu": "[parameters('cpuCores')]", "memoryInGB": "[parameters('memoryInGb')]",
							},
						},
					},
				},
				"ipAddress": map[string]interface{}{
					"ports": []interface{}{
						map[string]interface{}{
							"port":     "[parameters('port')]",
							"protocol": "Tcp",
						},
					},
					"type": "Public",
				},
				"osType":        "Linux",
				"restartPolicy": "[parameters('restartPolicy')]",
			},
		},
		{
			name: "Site",
			inputFunc: serialize(`
{
  "parameters": {
    "siteEnvelope": {
      "properties": {
        "siteConfig": "[variables('configReference')[parameters('language')]]",
        "serverFarmId": "[resourceId('Microsoft.Web/serverfarms', variables('appServicePlanPortalName'))]"
      }
    }
  }
}`),
			resourceName: "azure-native:web/v20150801:Site",
			expected: map[string]interface{}{
				"serverFarmId": "[resourceId('Microsoft.Web/serverfarms', variables('appServicePlanPortalName'))]",
				"siteConfig":   "[variables('configReference')[parameters('language')]]",
			},
		},
		{
			name: "DatabaseAccount",
			inputFunc: serialize(`
{
  "parameters": {
    "createUpdateParameters": {
        "properties": {
          "databaseAccountOfferType": "Standard",
          "locations": [
            {
                "id": "[concat(parameters('name'), '-', parameters('location'))]",
                "failoverPriority": 0,
                "locationName": "[parameters('location')]"
            }
          ],
          "backupPolicy": {
            "type": "Periodic",
            "periodicModeProperties": {
                "backupIntervalInMinutes": 240,
                "backupRetentionIntervalInHours": 8
            }
          },
          "capabilities": [
            {
                "name": "EnableServerless"
            }
          ],
          "enableFreeTier": false
        }
    }
  }
}`),
			resourceName: "azure-native:documentdb/v20200601preview:DatabaseAccount",
			expected: map[string]interface{}{
				"properties": map[string]interface{}{
					"backupPolicy": map[string]interface{}{
						"periodicModeProperties": map[string]interface{}{
							"backupIntervalInMinutes":        240,
							"backupRetentionIntervalInHours": 8,
						},
						"type": "Periodic",
					},
					"capabilities": []interface{}{
						map[string]interface{}{
							"name": "EnableServerless",
						},
					},
					"databaseAccountOfferType": "Standard",
					"enableFreeTier":           false,
					"locations": []interface{}{
						map[string]interface{}{
							"failoverPriority": 0,
							"locationName":     "[parameters('location')]",
						},
					},
				},
			},
		},
		{
			name: "JsonAny",
			inputFunc: serialize(`{
  "parameters": {
    "extensionParameters": {
      "properties": {
        "autoUpgradeMinorVersion": true,
        "publisher": "Microsoft.OSTCExtensions",
        "type": "VMAccessForLinux",
        "typeHandlerVersion": "1.4",
        "settings": {},
        "protectedSettings": {
            "username": "[parameters('extensions_enablevmaccess_username')]",
            "password": "[parameters('extensions_enablevmaccess_password')]",
            "ssh_key": "[parameters('extensions_enablevmaccess_ssh_key')]",
            "reset_ssh": "[parameters('extensions_enablevmaccess_reset_ssh')]",
            "remove_user": "[parameters('extensions_enablevmaccess_remove_user')]",
            "expiration": "[parameters('extensions_enablevmaccess_expiration')]"
        }
      }
    }
  }
}`),
			resourceName: "azure-native:compute/v20190701:VirtualMachineExtension",
			expected: map[string]interface{}{
				"autoUpgradeMinorVersion": true,
				"publisher":               "Microsoft.OSTCExtensions",
				"type":                    "VMAccessForLinux",
				"typeHandlerVersion":      "1.4",
				"settings":                map[string]interface{}{},
				"protectedSettings": map[string]interface{}{
					"username":    "[parameters('extensions_enablevmaccess_username')]",
					"password":    "[parameters('extensions_enablevmaccess_password')]",
					"ssh_key":     "[parameters('extensions_enablevmaccess_ssh_key')]",
					"reset_ssh":   "[parameters('extensions_enablevmaccess_reset_ssh')]",
					"remove_user": "[parameters('extensions_enablevmaccess_remove_user')]",
					"expiration":  "[parameters('extensions_enablevmaccess_expiration')]",
				},
			},
		},
		{
			name: "OneOf",
			input: map[string]interface{}{
				"parameters": map[string]interface{}{
					"subscriptionId":    "subscription-id",
					"resourceGroupName": "OneResourceGroupName",
					"api-version":       "2020-06-02",
					"resourceName":      "samplebotname",
					"channelName":       "AlexaChannel",
					"parameters": map[string]interface{}{
						"location": "global",
						"properties": map[string]interface{}{
							"channelName": "AlexaChannel",
							"properties": map[string]interface{}{
								"alexaSkillId": "XAlexaSkillIdX",
								"isEnabled":    true,
							},
						},
					},
				},
			},
			resourceName: "azure-native:botservice/v20200602:Channel",
			expected: map[string]interface{}{
				"channelName": "AlexaChannel",
				"location":    "global",
				"properties": map[string]interface{}{
					"channelName": "AlexaChannel",
					"properties": map[string]interface{}{
						"alexaSkillId": "XAlexaSkillIdX",
						"isEnabled":    true,
					},
				},
				"resourceGroupName": "OneResourceGroupName",
				"resourceName":      "samplebotname",
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.input == nil {
				test.input = test.inputFunc(t)
			}
			in := test.input["parameters"].(map[string]interface{})
			params := map[string]resources.AzureAPIParameter{}
			for _, param := range metadata.Resources[test.resourceName].PutParameters {
				params[param.Name] = param
			}
			out, err := FlattenParams(in, params, metadata.Types)
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

func serialize(input string) func(*testing.T) map[string]interface{} {
	return func(t *testing.T) map[string]interface{} {
		serialized := make(map[string]interface{})
		require.NoError(t, json.Unmarshal([]byte(input), &serialized))
		return serialized
	}
}

const (
	npe = `{
  "parameters": {
    "subscriptionId": "51766542-3ed7-4a72-a187-0c8ab644ddab",
    "resourceGroupName": "mygroup",
    "automationAccountName": "myaccount",
    "softwareUpdateConfigurationName": "testpatch",
    "api-version": "2017-05-15-preview",
    "parameters": {
      "properties": {
        "updateConfiguration": {
          "operatingSystem": "Windows",
          "duration": "PT2H0M",
          "windows": {
            "excludedKbNumbers": [
              "168934",
              "168973"
            ],
            "includedUpdateClassifications": "Critical",
            "rebootSetting": "IfRequired"
          },
          "azureVirtualMachines": [
            "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01",
            "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02",
            "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03"
          ],
          "nonAzureComputerNames": [
            "box1.contoso.com",
            "box2.contoso.com"
          ],
          "targets": {
            "azureQueries": [
              {
                "scope": [
                  "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources",
                  "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067"
                ],
                "tagSettings": {
                  "tags": [
                    {
                      "tag1": [
                        "tag1Value1",
                        "tag1Value2",
                        "tag1Value3"
                      ]
                    },
                    {
                      "tag2": [
                        "tag2Value1",
                        "tag2Value2",
                        "tag2Value3"
                      ]
                    }
                  ],
                  "filterOperator": "All"
                },
                "locations": [
                  "Japan East",
                  "UK South"
                ]
              }
            ],
            "nonAzureQueries": [
              {
                "functionAlias": "SavedSearch1",
                "workspaceId": "WorkspaceId1"
              },
              {
                "functionAlias": "SavedSearch2",
                "workspaceId": "WorkspaceId2"
              }
            ]
          }
        },
        "scheduleInfo": {
          "frequency": "Hour",
          "startTime": "2017-10-19T12:22:57+00:00",
          "timeZone": "America/Los_Angeles",
          "interval": 1,
          "expiryTime": "2018-11-09T11:22:57+00:00",
          "advancedSchedule": {
            "weekDays": [
              "Monday",
              "Thursday"
            ]
          }
        },
        "tasks": {
          "preTask": {
            "source": "HelloWorld",
            "parameters": {
              "COMPUTERNAME": "Computer1"
            }
          },
          "postTask": {
            "source": "GetCache",
            "parameters": null
          }
        }
      }
    }
  }
}`

	vmScaleSet = `{
  "parameters": {
    "subscriptionId": "{subscription-id}",
    "resourceGroupName": "myResourceGroup",
    "vmScaleSetName": "{vmss-name}",
    "api-version": "2020-06-01",
    "parameters": {
      "sku": {
        "tier": "Standard",
        "capacity": 3,
        "name": "Standard_D1_v2"
      },
      "location": "westus",
      "properties": {
        "overprovision": true,
        "virtualMachineProfile": {
          "storageProfile": {
            "imageReference": {
              "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"
            },
            "osDisk": {
              "caching": "ReadWrite",
              "managedDisk": {
                "storageAccountType": "Standard_LRS"
              },
              "createOption": "FromImage"
            }
          },
          "osProfile": {
            "computerNamePrefix": "{vmss-name}",
            "adminUsername": "{your-username}",
            "adminPassword": "{your-password}"
          },
          "networkProfile": {
            "networkInterfaceConfigurations": [
              {
                "name": "{vmss-name}",
                "properties": {
                  "primary": true,
                  "enableIPForwarding": true,
                  "ipConfigurations": [
                    {
                      "name": "{vmss-name}",
                      "properties": {
                        "subnet": {
                          "id": "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"
                        }
                      }
                    }
                  ]
                }
              }
            ]
          }
        },
        "upgradePolicy": {
          "mode": "Manual"
        }
      }
    }
  }
}`
)
