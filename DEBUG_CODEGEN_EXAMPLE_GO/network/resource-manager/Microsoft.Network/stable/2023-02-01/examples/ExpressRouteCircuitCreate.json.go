package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExpressRouteCircuit(ctx, "expressRouteCircuit", &network.ExpressRouteCircuitArgs{
			AllowClassicOperations: pulumi.Bool(false),
			Authorizations:         network.ExpressRouteCircuitAuthorizationTypeArray{},
			CircuitName:            pulumi.String("circuitName"),
			Location:               pulumi.String("Brazil South"),
			Peerings:               network.ExpressRouteCircuitPeeringTypeArray{},
			ResourceGroupName:      pulumi.String("rg1"),
			ServiceProviderProperties: &network.ExpressRouteCircuitServiceProviderPropertiesArgs{
				BandwidthInMbps:     pulumi.Int(200),
				PeeringLocation:     pulumi.String("Silicon Valley"),
				ServiceProviderName: pulumi.String("Equinix"),
			},
			Sku: &network.ExpressRouteCircuitSkuArgs{
				Family: pulumi.String("MeteredData"),
				Name:   pulumi.String("Standard_MeteredData"),
				Tier:   pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
