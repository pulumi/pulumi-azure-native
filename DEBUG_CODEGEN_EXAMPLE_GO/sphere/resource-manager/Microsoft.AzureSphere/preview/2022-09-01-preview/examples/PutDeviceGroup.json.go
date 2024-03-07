package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuresphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuresphere.NewDeviceGroup(ctx, "deviceGroup", &azuresphere.DeviceGroupArgs{
			CatalogName:       pulumi.String("MyCatalog1"),
			Description:       pulumi.String("Description for MyDeviceGroup1"),
			DeviceGroupName:   pulumi.String("MyDeviceGroup1"),
			OsFeedType:        pulumi.String("Retail"),
			ProductName:       pulumi.String("MyProduct1"),
			ResourceGroupName: pulumi.String("MyResourceGroup1"),
			UpdatePolicy:      pulumi.String("UpdateAll"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
