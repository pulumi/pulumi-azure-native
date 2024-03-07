package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewPipelineRun(ctx, "pipelineRun", &containerregistry.PipelineRunArgs{
			ForceUpdateTag:  pulumi.String("2020-03-04T17:23:21.9261521+00:00"),
			PipelineRunName: pulumi.String("myPipelineRun"),
			RegistryName:    pulumi.String("myRegistry"),
			Request: &containerregistry.PipelineRunRequestArgs{
				CatalogDigest:      pulumi.String("sha256@"),
				PipelineResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
				Source: &containerregistry.PipelineRunSourcePropertiesArgs{
					Name: pulumi.String("myblob.tar.gz"),
					Type: pulumi.String("AzureStorageBlob"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
