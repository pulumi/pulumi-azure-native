package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewPrivateEndpointConnectionControllerPrivateEndpointConnection(ctx, "privateEndpointConnectionControllerPrivateEndpointConnection", &migrate.PrivateEndpointConnectionControllerPrivateEndpointConnectionArgs{
			MigrateProjectName: pulumi.String("proj567"),
			PeConnectionName:   pulumi.String("proj5675162pe.fdccace0-e303-4a79-80c8-3aa7c1f09cc6"),
			Properties: &migrate.ConnectionStateRequestBodyPropertiesArgs{
				PrivateLinkServiceConnectionState: &migrate.PrivateLinkServiceConnectionStateArgs{
					ActionsRequired: pulumi.String(""),
					Status:          pulumi.String("Approved"),
				},
			},
			ResourceGroupName: pulumi.String("pajindTest1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
