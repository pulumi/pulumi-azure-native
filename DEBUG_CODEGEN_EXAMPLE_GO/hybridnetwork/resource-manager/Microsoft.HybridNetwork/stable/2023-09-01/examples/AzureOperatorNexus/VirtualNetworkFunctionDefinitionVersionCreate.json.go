package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridnetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridnetwork.NewNetworkFunctionDefinitionVersion(ctx, "networkFunctionDefinitionVersion", &hybridnetwork.NetworkFunctionDefinitionVersionArgs{
			Location:                             pulumi.String("eastus"),
			NetworkFunctionDefinitionGroupName:   pulumi.String("TestNetworkFunctionDefinitionGroupName"),
			NetworkFunctionDefinitionVersionName: pulumi.String("1.0.0"),
			Properties: hybridnetwork.VirtualNetworkFunctionNetworkFunctionDefinitionVersion{
				DeployParameters: "{\"virtualMachineName\":{\"type\":\"string\"},\"extendedLocationName\":{\"type\":\"string\"},\"cpuCores\":{\"type\":\"int\"},\"memorySizeGB\":{\"type\":\"int\"},\"cloudServicesNetworkAttachment\":{\"type\":\"object\",\"properties\":{\"networkAttachmentName\":{\"type\":\"string\"},\"attachedNetworkId\":{\"type\":\"string\"},\"ipAllocationMethod\":{\"type\":\"string\"},\"ipv4Address\":{\"type\":\"string\"},\"ipv6Address\":{\"type\":\"string\"},\"defaultGateway\":{\"type\":\"string\"}},\"required\":[\"attachedNetworkId\",\"ipAllocationMethod\"]},\"networkAttachments\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"networkAttachmentName\":{\"type\":\"string\"},\"attachedNetworkId\":{\"type\":\"string\"},\"ipAllocationMethod\":{\"type\":\"string\"},\"ipv4Address\":{\"type\":\"string\"},\"ipv6Address\":{\"type\":\"string\"},\"defaultGateway\":{\"type\":\"string\"}},\"required\":[\"attachedNetworkId\",\"ipAllocationMethod\"]}},\"storageProfile\":{\"type\":\"object\",\"properties\":{\"osDisk\":{\"type\":\"object\",\"properties\":{\"createOption\":{\"type\":\"string\"},\"deleteOption\":{\"type\":\"string\"},\"diskSizeGB\":{\"type\":\"integer\"}},\"required\":[\"diskSizeGB\"]}},\"required\":[\"osDisk\"]},\"sshPublicKeys\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"keyData\":{\"type\":\"string\"}},\"required\":[\"keyData\"]}},\"userData\":{\"type\":\"string\"},\"adminUsername\":{\"type\":\"string\"},\"bootMethod\":{\"type\":\"string\",\"default\":\"UEFI\",\"enum\":[\"UEFI\",\"BIOS\"]},\"isolateEmulatorThread\":{\"type\":\"string\"},\"virtioInterface\":{\"type\":\"string\"},\"placementHints\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"hintType\":{\"type\":\"string\",\"enum\":[\"Affinity\",\"AntiAffinity\"]},\"resourceId\":{\"type\":\"string\"},\"schedulingExecution\":{\"type\":\"string\",\"enum\":[\"Soft\",\"Hard\"]},\"scope\":{\"type\":\"string\"}},\"required\":[\"hintType\",\"schedulingExecution\",\"resourceId\",\"scope\"]}}}",
				Description:      "test NFDV for AzureOperatorNexus",
				NetworkFunctionTemplate: hybridnetwork.AzureOperatorNexusNetworkFunctionTemplate{
					NetworkFunctionApplications: []interface{}{
						hybridnetwork.AzureOperatorNexusNetworkFunctionImageApplication{
							ArtifactProfile: hybridnetwork.AzureOperatorNexusImageArtifactProfile{
								ArtifactStore: hybridnetwork.ReferencedResource{
									Id: "/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore",
								},
								ImageArtifactProfile: hybridnetwork.ImageArtifactProfile{
									ImageName:    "test-image",
									ImageVersion: "1.0.0",
								},
							},
							ArtifactType: "ImageFile",
							DependsOnProfile: hybridnetwork.DependsOnProfile{
								InstallDependsOn:   []interface{}{},
								UninstallDependsOn: []interface{}{},
								UpdateDependsOn:    []interface{}{},
							},
							DeployParametersMappingRuleProfile: hybridnetwork.AzureOperatorNexusImageDeployMappingRuleProfile{
								ApplicationEnablement: "Unknown",
								ImageMappingRuleProfile: hybridnetwork.ImageMappingRuleProfile{
									UserConfiguration: "",
								},
							},
							Name: "testImageRole",
						},
						hybridnetwork.AzureOperatorNexusNetworkFunctionArmTemplateApplication{
							ArtifactProfile: hybridnetwork.AzureOperatorNexusArmTemplateArtifactProfile{
								ArtifactStore: hybridnetwork.ReferencedResource{
									Id: "/subscriptions/subid/resourceGroups/rg/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore",
								},
								TemplateArtifactProfile: hybridnetwork.ArmTemplateArtifactProfile{
									TemplateName:    "test-template",
									TemplateVersion: "1.0.0",
								},
							},
							ArtifactType: "ArmTemplate",
							DependsOnProfile: hybridnetwork.DependsOnProfile{
								InstallDependsOn: []string{
									"testImageRole",
								},
								UninstallDependsOn: []string{
									"testImageRole",
								},
								UpdateDependsOn: []string{
									"testImageRole",
								},
							},
							DeployParametersMappingRuleProfile: hybridnetwork.AzureOperatorNexusArmTemplateDeployMappingRuleProfile{
								ApplicationEnablement: "Unknown",
								TemplateMappingRuleProfile: hybridnetwork.ArmTemplateMappingRuleProfile{
									TemplateParameters: "{\"virtualMachineName\":\"{deployParameters.virtualMachineName}\",\"extendedLocationName\":\"{deployParameters.extendedLocationName}\",\"cpuCores\":\"{deployParameters.cpuCores}\",\"memorySizeGB\":\"{deployParameters.memorySizeGB}\",\"cloudServicesNetworkAttachment\":\"{deployParameters.cloudServicesNetworkAttachment}\",\"networkAttachments\":\"{deployParameters.networkAttachments}\",\"sshPublicKeys\":\"{deployParameters.sshPublicKeys}\",\"storageProfile\":\"{deployParameters.storageProfile}\",\"isolateEmulatorThread\":\"{deployParameters.isolateEmulatorThread}\",\"virtioInterface\":\"{deployParameters.virtioInterface}\",\"userData\":\"{deployParameters.userData}\",\"adminUsername\":\"{deployParameters.adminUsername}\",\"bootMethod\":\"{deployParameters.bootMethod}\",\"placementHints\":\"{deployParameters.placementHints}\"}",
								},
							},
							Name: "testTemplateRole",
						},
					},
					NfviType: "AzureOperatorNexus",
				},
				NetworkFunctionType: "VirtualNetworkFunction",
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
