package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualHubBgpConnection(ctx, "virtualHubBgpConnection", &network.VirtualHubBgpConnectionArgs{
			ConnectionName: pulumi.String("conn1"),
			HubVirtualNetworkConnection: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubVirtualNetworkConnections/hubVnetConn1"),
			},
			PeerAsn:           pulumi.Float64(20000),
			PeerIp:            pulumi.String("192.168.1.5"),
			ResourceGroupName: pulumi.String("rg1"),
			VirtualHubName:    pulumi.String("hub1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
