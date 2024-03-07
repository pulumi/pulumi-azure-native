package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewApplicationGatewayPrivateEndpointConnection(ctx, "applicationGatewayPrivateEndpointConnection", &network.ApplicationGatewayPrivateEndpointConnectionArgs{
			ApplicationGatewayName: pulumi.String("appgw"),
			ConnectionName:         pulumi.String("connection1"),
			Name:                   pulumi.String("connection1"),
			PrivateLinkServiceConnectionState: &network.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("approved it for some reason."),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
