package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewBigDataPool(ctx, "bigDataPool", &synapse.BigDataPoolArgs{
			AutoPause: &synapse.AutoPausePropertiesArgs{
				DelayInMinutes: pulumi.Int(15),
				Enabled:        pulumi.Bool(true),
			},
			AutoScale: &synapse.AutoScalePropertiesArgs{
				Enabled:      pulumi.Bool(true),
				MaxNodeCount: pulumi.Int(50),
				MinNodeCount: pulumi.Int(3),
			},
			BigDataPoolName:       pulumi.String("ExamplePool"),
			DefaultSparkLogFolder: pulumi.String("/logs"),
			IsAutotuneEnabled:     pulumi.Bool(false),
			LibraryRequirements: &synapse.LibraryRequirementsArgs{
				Content:  pulumi.String(""),
				Filename: pulumi.String("requirements.txt"),
			},
			Location:          pulumi.String("West US 2"),
			NodeCount:         pulumi.Int(4),
			NodeSize:          pulumi.String("Medium"),
			NodeSizeFamily:    pulumi.String("MemoryOptimized"),
			ResourceGroupName: pulumi.String("ExampleResourceGroup"),
			SparkEventsFolder: pulumi.String("/events"),
			SparkVersion:      pulumi.String("3.3"),
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
			WorkspaceName: pulumi.String("ExampleWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
