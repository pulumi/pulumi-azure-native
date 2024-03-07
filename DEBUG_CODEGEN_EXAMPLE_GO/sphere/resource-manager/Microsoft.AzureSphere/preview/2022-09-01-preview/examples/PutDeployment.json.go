package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azuresphere/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azuresphere.NewDeployment(ctx, "deployment", &azuresphere.DeploymentArgs{
			CatalogName:       pulumi.String("MyCatalog1"),
			DeploymentName:    pulumi.String("MyDeployment1"),
			DeviceGroupName:   pulumi.String("myDeviceGroup1"),
			ProductName:       pulumi.String("MyProduct1"),
			ResourceGroupName: pulumi.String("MyResourceGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
