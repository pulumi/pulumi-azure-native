package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &migrate.PrivateEndpointConnectionArgs{
			ETag:                          pulumi.String("\"00009300-0000-0300-0000-602b967b0000\""),
			PrivateEndpointConnectionName: pulumi.String("custestpece80project3980pe.7e35576b-3df4-478e-9759-f64351cf4f43"),
			ProjectName:                   pulumi.String("abgoyalWEselfhostb72bproject"),
			Properties: migrate.PrivateEndpointConnectionPropertiesResponse{
				PrivateLinkServiceConnectionState: &migrate.PrivateLinkServiceConnectionStateArgs{
					ActionsRequired: pulumi.String(""),
					Status:          pulumi.String("Approved"),
				},
			},
			ResourceGroupName: pulumi.String("abgoyal-westEurope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
