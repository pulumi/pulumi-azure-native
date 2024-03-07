package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualHub(ctx, "virtualHub", &network.VirtualHubArgs{
			AddressPrefix:     pulumi.String("10.168.0.0/24"),
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("rg1"),
			Sku:               pulumi.String("Basic"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			VirtualHubName: pulumi.String("virtualHub2"),
			VirtualWan: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
