package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewRoute(ctx, "route", &network.RouteArgs{
			AddressPrefix:     pulumi.String("10.0.3.0/24"),
			NextHopType:       pulumi.String("VirtualNetworkGateway"),
			ResourceGroupName: pulumi.String("rg1"),
			RouteName:         pulumi.String("route1"),
			RouteTableName:    pulumi.String("testrt"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
