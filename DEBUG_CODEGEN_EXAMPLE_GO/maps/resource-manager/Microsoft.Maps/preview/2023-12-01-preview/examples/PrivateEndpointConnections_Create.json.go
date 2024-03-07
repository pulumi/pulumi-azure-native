package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/maps/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := maps.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &maps.PrivateEndpointConnectionArgs{
			AccountName:                   pulumi.String("myMapsAccount"),
			PrivateEndpointConnectionName: pulumi.String("privateEndpointConnectionName"),
			PrivateLinkServiceConnectionState: &maps.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
