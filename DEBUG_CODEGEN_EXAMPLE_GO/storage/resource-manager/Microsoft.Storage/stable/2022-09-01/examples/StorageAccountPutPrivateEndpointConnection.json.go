package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storage.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &storage.PrivateEndpointConnectionArgs{
			AccountName:                   pulumi.String("sto9699"),
			PrivateEndpointConnectionName: pulumi.String("{privateEndpointConnectionName}"),
			PrivateLinkServiceConnectionState: &storage.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("res7687"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
