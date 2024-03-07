package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewDevice(ctx, "device", &databoxedge.DeviceArgs{
			DeviceName:        pulumi.String("testedgedevice"),
			Location:          pulumi.String("WUS"),
			ResourceGroupName: pulumi.String("GroupForEdgeAutomation"),
			Sku: &databoxedge.SkuArgs{
				Name: pulumi.String("Edge"),
				Tier: pulumi.String("Standard"),
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
