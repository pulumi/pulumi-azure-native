package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagemover/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagemover.NewJobDefinition(ctx, "jobDefinition", &storagemover.JobDefinitionArgs{
			AgentName:         pulumi.String("migration-agent"),
			CopyMode:          pulumi.String("Additive"),
			Description:       pulumi.String("Example Job Definition Description"),
			JobDefinitionName: pulumi.String("examples-jobDefinitionName"),
			ProjectName:       pulumi.String("examples-projectName"),
			ResourceGroupName: pulumi.String("examples-rg"),
			SourceName:        pulumi.String("examples-sourceEndpointName"),
			SourceSubpath:     pulumi.String("/"),
			StorageMoverName:  pulumi.String("examples-storageMoverName"),
			TargetName:        pulumi.String("examples-targetEndpointName"),
			TargetSubpath:     pulumi.String("/"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
