package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewExpressRoutePortAuthorization(ctx, "expressRoutePortAuthorization", &network.ExpressRoutePortAuthorizationArgs{
			AuthorizationName:    pulumi.String("authorizatinName"),
			ExpressRoutePortName: pulumi.String("expressRoutePortName"),
			ResourceGroupName:    pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
