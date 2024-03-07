package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewServerInstance(ctx, "serverInstance", &workloads.ServerInstanceArgs{
			ResourceGroupName:    pulumi.String("test-rg"),
			SapDiscoverySiteName: pulumi.String("SampleSite"),
			SapInstanceName:      pulumi.String("MPP_MPP"),
			ServerInstanceName:   pulumi.String("APP_SapServer1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
