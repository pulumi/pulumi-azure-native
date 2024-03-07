package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewPipelineRun(ctx, "pipelineRun", &containerregistry.PipelineRunArgs{
			PipelineRunName: pulumi.String("myPipelineRun"),
			RegistryName:    pulumi.String("myRegistry"),
			Request: &containerregistry.PipelineRunRequestArgs{
				Artifacts: pulumi.StringArray{
					pulumi.String("sourceRepository/hello-world"),
					pulumi.String("sourceRepository2@sha256:00000000000000000000000000000000000"),
				},
				PipelineResourceId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/exportPipelines/myExportPipeline"),
				Target: &containerregistry.PipelineRunTargetPropertiesArgs{
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
