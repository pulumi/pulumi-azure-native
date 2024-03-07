package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &media.PrivateEndpointConnectionArgs{
			AccountName: pulumi.String("contososports"),
			Name:        pulumi.String("connectionName1"),
			PrivateLinkServiceConnectionState: &media.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Test description."),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("contosorg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
