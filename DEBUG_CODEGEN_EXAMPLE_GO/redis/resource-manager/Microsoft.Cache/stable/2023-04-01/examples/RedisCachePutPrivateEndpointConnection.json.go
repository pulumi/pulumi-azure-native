package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cache.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &cache.PrivateEndpointConnectionArgs{
			CacheName:                     pulumi.String("cachetest01"),
			PrivateEndpointConnectionName: pulumi.String("pectest01"),
			PrivateLinkServiceConnectionState: &cache.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rgtest01"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
