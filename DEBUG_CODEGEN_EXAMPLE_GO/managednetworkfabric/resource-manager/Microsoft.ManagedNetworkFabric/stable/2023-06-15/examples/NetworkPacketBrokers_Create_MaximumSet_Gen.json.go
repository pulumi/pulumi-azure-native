package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewNetworkPacketBroker(ctx, "networkPacketBroker", &managednetworkfabric.NetworkPacketBrokerArgs{
			Location:                pulumi.String("eastuseuap"),
			NetworkFabricId:         pulumi.String("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
			NetworkPacketBrokerName: pulumi.String("example-networkPacketBroker"),
			ResourceGroupName:       pulumi.String("example-rg"),
			Tags: pulumi.StringMap{
				"key2806": pulumi.String("key"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
