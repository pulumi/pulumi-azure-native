package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devices.NewIotDpsResource(ctx, "iotDpsResource", &devices.IotDpsResourceArgs{
			Location: pulumi.String("East US"),
			Properties: &devices.IotDpsPropertiesDescriptionArgs{
				EnableDataResidency: pulumi.Bool(false),
			},
			ProvisioningServiceName: pulumi.String("myFirstProvisioningService"),
			ResourceGroupName:       pulumi.String("myResourceGroup"),
			Sku: &devices.IotDpsSkuInfoArgs{
				Capacity: pulumi.Float64(1),
				Name:     pulumi.String("S1"),
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
