package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewRouteTable(ctx, "routeTable", &network.RouteTableArgs{
			DisableBgpRoutePropagation: pulumi.Bool(true),
			Location:                   pulumi.String("westus"),
			ResourceGroupName:          pulumi.String("rg1"),
			RouteTableName:             pulumi.String("testrt"),
			Routes: network.RouteTypeArray{
				&network.RouteTypeArgs{
					AddressPrefix: pulumi.String("10.0.3.0/24"),
					Name:          pulumi.String("route1"),
					NextHopType:   pulumi.String("VirtualNetworkGateway"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
