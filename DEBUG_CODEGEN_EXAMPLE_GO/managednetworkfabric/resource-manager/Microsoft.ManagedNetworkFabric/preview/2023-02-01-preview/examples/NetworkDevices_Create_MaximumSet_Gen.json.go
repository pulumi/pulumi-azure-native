package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewNetworkDevice(ctx, "networkDevice", &managednetworkfabric.NetworkDeviceArgs{
			Annotation:        pulumi.String("null"),
			HostName:          pulumi.String("networkDeviceName"),
			Location:          pulumi.String("eastus"),
			NetworkDeviceName: pulumi.String("networkDeviceName"),
			NetworkDeviceRole: pulumi.String("CE"),
			NetworkDeviceSku:  pulumi.String("DefaultSku"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			SerialNumber:      pulumi.String("Arista;DCS-7280PR3-24;12.05;JPE21330382"),
			Tags: pulumi.StringMap{
				"keyID": pulumi.String("keyValue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
