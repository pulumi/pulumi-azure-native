package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewNetworkServiceDesignVersion(ctx, "networkServiceDesignVersion", &hybridnetwork.NetworkServiceDesignVersionArgs{
			Location:                        pulumi.String("eastus"),
			NetworkServiceDesignGroupName:   pulumi.String("TestNetworkServiceDesignGroupName"),
			NetworkServiceDesignVersionName: pulumi.String("1.0.0"),
			Properties: hybridnetwork.NetworkServiceDesignVersionPropertiesFormatResponse{
				ConfigurationGroupSchemaReferences: hybridnetwork.ReferencedResourceMap{
					"MyVM_Configuration": &hybridnetwork.ReferencedResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/publishers/contosoGroup/networkServiceDesignGroups/NSD_contoso/configurationGroupSchemas/MyVM_Configuration_Schema"),
					},
				},
				ResourceElementTemplates: pulumi.Array{
					hybridnetwork.ArmResourceDefinitionResourceElementTemplateDetails{
						Configuration: hybridnetwork.ArmResourceDefinitionResourceElementTemplate{
							ArtifactProfile: hybridnetwork.NSDArtifactProfile{
								ArtifactName: "MyVMArmTemplate",
								ArtifactStoreReference: hybridnetwork.ReferencedResource{
									Id: "/subscriptions/subid/providers/Microsoft.HybridNetwork/publishers/contosoGroup/artifactStoreReference/store1",
								},
								ArtifactVersion: "1.0.0",
							},
							ParameterValues: "{\"publisherName\":\"{configurationparameters('MyVM_Configuration').publisherName}\",\"skuGroupName\":\"{configurationparameters('MyVM_Configuration').skuGroupName}\",\"skuVersion\":\"{configurationparameters('MyVM_Configuration').skuVersion}\",\"skuOfferingLocation\":\"{configurationparameters('MyVM_Configuration').skuOfferingLocation}\",\"nfviType\":\"{nfvis().nfvisFromSitePerNfviType.AzureCore.nfviAlias1.nfviType}\",\"nfviId\":\"{nfvis().nfvisFromSitePerNfviType.AzureCore.nfviAlias1.nfviId}\",\"allowSoftwareUpdates\":\"{configurationparameters('MyVM_Configuration').allowSoftwareUpdates}\",\"virtualNetworkName\":\"{configurationparameters('MyVM_Configuration').vnetName}\",\"subnetName\":\"{configurationparameters('MyVM_Configuration').subnetName}\",\"subnetAddressPrefix\":\"{configurationparameters('MyVM_Configuration').subnetAddressPrefix}\",\"managedResourceGroup\":\"{configurationparameters('SNSSelf').managedResourceGroupName}\",\"adminPassword\":\"{secretparameters('MyVM_Configuration').adminPassword}\"}",
							TemplateType:    "ArmTemplate",
						},
						DependsOnProfile: hybridnetwork.DependsOnProfile{
							InstallDependsOn: []interface{}{},
						},
						Name:                "MyVM",
						ResourceElementType: "ArmResourceDefinition",
					},
				},
			},
			PublisherName:     pulumi.String("TestPublisher"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
