package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewSapInstance(ctx, "sapInstance", &workloads.SapInstanceArgs{
			Location:             pulumi.String("eastus"),
			ResourceGroupName:    pulumi.String("test-rg"),
			SapDiscoverySiteName: pulumi.String("SampleSite"),
			SapInstanceName:      pulumi.String("MPP_MPP"),
			Tags: pulumi.StringMap{
				"property1": pulumi.String("value1"),
				"property2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
