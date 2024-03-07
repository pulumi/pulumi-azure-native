package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExpressRouteCrossConnectionPeering(ctx, "expressRouteCrossConnectionPeering", &network.ExpressRouteCrossConnectionPeeringArgs{
			CrossConnectionName: pulumi.String("<circuitServiceKey>"),
			Ipv6PeeringConfig: &network.Ipv6ExpressRouteCircuitPeeringConfigArgs{
				PrimaryPeerAddressPrefix:   pulumi.String("3FFE:FFFF:0:CD30::/126"),
				SecondaryPeerAddressPrefix: pulumi.String("3FFE:FFFF:0:CD30::4/126"),
			},
			PeerASN:                    pulumi.Float64(200),
			PeeringName:                pulumi.String("AzurePrivatePeering"),
			PrimaryPeerAddressPrefix:   pulumi.String("192.168.16.252/30"),
			ResourceGroupName:          pulumi.String("CrossConnection-SiliconValley"),
			SecondaryPeerAddressPrefix: pulumi.String("192.168.18.252/30"),
			VlanId:                     pulumi.Int(200),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
