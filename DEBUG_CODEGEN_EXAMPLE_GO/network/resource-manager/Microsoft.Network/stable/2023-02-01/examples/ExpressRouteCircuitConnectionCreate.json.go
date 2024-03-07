package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExpressRouteCircuitConnection(ctx, "expressRouteCircuitConnection", &network.ExpressRouteCircuitConnectionArgs{
			AddressPrefix:    pulumi.String("10.0.0.0/29"),
			AuthorizationKey: pulumi.String("946a1918-b7a2-4917-b43c-8c4cdaee006a"),
			CircuitName:      pulumi.String("ExpressRouteARMCircuitA"),
			ConnectionName:   pulumi.String("circuitConnectionUSAUS"),
			ExpressRouteCircuitPeering: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid1/resourceGroups/dedharcktinit/providers/Microsoft.Network/expressRouteCircuits/dedharcktlocal/peerings/AzurePrivatePeering"),
			},
			Ipv6CircuitConnectionConfig: &network.Ipv6CircuitConnectionConfigArgs{
				AddressPrefix: pulumi.String("aa:bb::/125"),
			},
			PeerExpressRouteCircuitPeering: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid2/resourceGroups/dedharcktpeer/providers/Microsoft.Network/expressRouteCircuits/dedharcktremote/peerings/AzurePrivatePeering"),
			},
			PeeringName:       pulumi.String("AzurePrivatePeering"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
