package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagesync/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagesync.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &storagesync.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("{privateEndpointConnectionName}"),
			PrivateLinkServiceConnectionState: &storagesync.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName:      pulumi.String("res7687"),
			StorageSyncServiceName: pulumi.String("sss2527"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
