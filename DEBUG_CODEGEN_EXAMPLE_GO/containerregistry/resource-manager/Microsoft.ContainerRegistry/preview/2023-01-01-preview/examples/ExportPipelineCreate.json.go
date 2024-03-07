package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewExportPipeline(ctx, "exportPipeline", &containerregistry.ExportPipelineArgs{
			ExportPipelineName: pulumi.String("myExportPipeline"),
			Identity: &containerregistry.IdentityPropertiesArgs{
				Type: containerregistry.ResourceIdentityTypeSystemAssigned,
			},
			Location: pulumi.String("westus"),
			Options: pulumi.StringArray{
				pulumi.String("OverwriteBlobs"),
			},
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Target: &containerregistry.ExportPipelineTargetPropertiesArgs{
				KeyVaultUri: pulumi.String("https://myvault.vault.azure.net/secrets/acrexportsas"),
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
