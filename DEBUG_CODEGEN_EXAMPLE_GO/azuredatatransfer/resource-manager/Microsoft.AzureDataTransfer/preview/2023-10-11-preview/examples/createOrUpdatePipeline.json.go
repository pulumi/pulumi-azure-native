package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuredatatransfer/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuredatatransfer.NewPipeline(ctx, "pipeline", &azuredatatransfer.PipelineArgs{
			Location:     pulumi.String("East US"),
			PipelineName: pulumi.String("testPipeline"),
			Properties: &azuredatatransfer.PipelinePropertiesArgs{
				RemoteCloud: pulumi.String("testdc"),
			},
			ResourceGroupName: pulumi.String("testRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
