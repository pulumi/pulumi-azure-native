package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewImportPipeline(ctx, "importPipeline", &containerregistry.ImportPipelineArgs{
			Identity: &containerregistry.IdentityPropertiesArgs{
				Type: containerregistry.ResourceIdentityTypeUserAssigned,
				UserAssignedIdentities: containerregistry.UserIdentityPropertiesMap{
					"/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2": nil,
				},
			},
			ImportPipelineName: pulumi.String("myImportPipeline"),
			Location:           pulumi.String("westus"),
			Options: pulumi.StringArray{
				pulumi.String("OverwriteTags"),
				pulumi.String("DeleteSourceBlobOnSuccess"),
				pulumi.String("ContinueOnErrors"),
			},
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Source: &containerregistry.ImportPipelineSourcePropertiesArgs{
				KeyVaultUri: pulumi.String("https://myvault.vault.azure.net/secrets/acrimportsas"),
				Type:        pulumi.String("AzureStorageBlobContainer"),
				Uri:         pulumi.String("https://accountname.blob.core.windows.net/containername"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
