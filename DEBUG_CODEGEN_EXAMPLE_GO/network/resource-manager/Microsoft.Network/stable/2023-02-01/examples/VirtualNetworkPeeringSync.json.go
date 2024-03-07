package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualNetworkPeering(ctx, "virtualNetworkPeering", &network.VirtualNetworkPeeringArgs{
			AllowForwardedTraffic:     pulumi.Bool(true),
			AllowGatewayTransit:       pulumi.Bool(false),
			AllowVirtualNetworkAccess: pulumi.Bool(true),
			RemoteVirtualNetwork: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
			},
			ResourceGroupName:         pulumi.String("peerTest"),
			SyncRemoteAddressSpace:    pulumi.String("true"),
			UseRemoteGateways:         pulumi.Bool(false),
			VirtualNetworkName:        pulumi.String("vnet1"),
			VirtualNetworkPeeringName: pulumi.String("peer"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
