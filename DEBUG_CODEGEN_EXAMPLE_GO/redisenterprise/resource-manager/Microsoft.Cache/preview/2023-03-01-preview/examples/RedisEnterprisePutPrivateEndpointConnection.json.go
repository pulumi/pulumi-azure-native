package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cache.NewEnterprisePrivateEndpointConnection(ctx, "enterprisePrivateEndpointConnection", &cache.EnterprisePrivateEndpointConnectionArgs{
			ClusterName:                   pulumi.String("cache1"),
			PrivateEndpointConnectionName: pulumi.String("pectest01"),
			PrivateLinkServiceConnectionState: &cache.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
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
