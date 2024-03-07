package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuresphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuresphere.NewDevice(ctx, "device", &azuresphere.DeviceArgs{
			CatalogName:       pulumi.String("MyCatalog1"),
			DeviceGroupName:   pulumi.String("myDeviceGroup1"),
			DeviceName:        pulumi.String("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
			ProductName:       pulumi.String("MyProduct1"),
			ResourceGroupName: pulumi.String("MyResourceGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
