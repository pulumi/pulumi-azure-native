package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualRouterPeering(ctx, "virtualRouterPeering", &network.VirtualRouterPeeringArgs{
			PeerAsn:           pulumi.Float64(20000),
			PeerIp:            pulumi.String("192.168.1.5"),
			PeeringName:       pulumi.String("peering1"),
			ResourceGroupName: pulumi.String("rg1"),
			VirtualRouterName: pulumi.String("virtualRouter"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
