package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewNetworkInterface(ctx, "networkInterface", &managednetworkfabric.NetworkInterfaceArgs{
			Annotation:           pulumi.String("null"),
			NetworkDeviceName:    pulumi.String("networkDeviceName"),
			NetworkInterfaceName: pulumi.String("networkInterfaceName"),
			ResourceGroupName:    pulumi.String("resourceGroupName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
