package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewPrivateEndpointConnectionProxyController(ctx, "privateEndpointConnectionProxyController", &migrate.PrivateEndpointConnectionProxyControllerArgs{
			ETag:               pulumi.String("ftvkdifbymdoybmuhqocd"),
			MigrateProjectName: pulumi.String("1GQwlI-"),
			PecProxyName:       pulumi.String("R-0-fb4"),
			ResourceGroupName:  pulumi.String("rghubmigrate"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
