package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExpressRouteCircuit(ctx, "expressRouteCircuit", &network.ExpressRouteCircuitArgs{
			AuthorizationKey: pulumi.String("b0be57f5-1fba-463b-adec-ffe767354cdd"),
			BandwidthInGbps:  pulumi.Float64(10),
			CircuitName:      pulumi.String("expressRouteCircuit1"),
			ExpressRoutePort: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName"),
			},
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("rg1"),
			Sku: &network.ExpressRouteCircuitSkuArgs{
				Family: pulumi.String("MeteredData"),
				Name:   pulumi.String("Premium_MeteredData"),
				Tier:   pulumi.String("Premium"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
