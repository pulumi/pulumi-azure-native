package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExpressRouteCircuitPeering(ctx, "expressRouteCircuitPeering", &network.ExpressRouteCircuitPeeringArgs{
			CircuitName:                pulumi.String("circuitName"),
			PeerASN:                    pulumi.Float64(200),
			PeeringName:                pulumi.String("AzurePrivatePeering"),
			PrimaryPeerAddressPrefix:   pulumi.String("192.168.16.252/30"),
			ResourceGroupName:          pulumi.String("rg1"),
			SecondaryPeerAddressPrefix: pulumi.String("192.168.18.252/30"),
			VlanId:                     pulumi.Int(200),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
