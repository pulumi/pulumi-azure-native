package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewPrivateEndpointConnectionProxyController(ctx, "privateEndpointConnectionProxyController", &migrate.PrivateEndpointConnectionProxyControllerArgs{
			MigrateProjectName: pulumi.String("1GQwlI-"),
			PecProxyName:       pulumi.String("z1LfRIz4-M2-1-V7"),
			ResourceGroupName:  pulumi.String("rghubmigrate"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
