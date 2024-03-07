package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExpressRouteCircuitAuthorization(ctx, "expressRouteCircuitAuthorization", &network.ExpressRouteCircuitAuthorizationArgs{
			AuthorizationName: pulumi.String("authorizatinName"),
			CircuitName:       pulumi.String("circuitName"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
