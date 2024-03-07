package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewStaticMember(ctx, "staticMember", &network.StaticMemberArgs{
			NetworkGroupName:   pulumi.String("testNetworkGroup"),
			NetworkManagerName: pulumi.String("testNetworkManager"),
			ResourceGroupName:  pulumi.String("rg1"),
			ResourceId:         pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroup/rg1/providers/Microsoft.Network/virtualnetworks/vnet1"),
			StaticMemberName:   pulumi.String("testStaticMember"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
