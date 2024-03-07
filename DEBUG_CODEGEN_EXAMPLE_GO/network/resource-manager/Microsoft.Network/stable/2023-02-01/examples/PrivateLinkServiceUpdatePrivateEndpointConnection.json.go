package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPrivateLinkServicePrivateEndpointConnection(ctx, "privateLinkServicePrivateEndpointConnection", &network.PrivateLinkServicePrivateEndpointConnectionArgs{
			Name:             pulumi.String("testPlePeConnection"),
			PeConnectionName: pulumi.String("testPlePeConnection"),
			PrivateLinkServiceConnectionState: &network.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("approved it for some reason."),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("testPls"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
