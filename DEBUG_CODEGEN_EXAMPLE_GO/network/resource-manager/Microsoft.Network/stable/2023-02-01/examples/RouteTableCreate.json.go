package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewRouteTable(ctx, "routeTable", &network.RouteTableArgs{
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("rg1"),
			RouteTableName:    pulumi.String("testrt"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
