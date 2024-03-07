package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/migrate/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := migrate.NewPrivateEndpointConnectionOperation(ctx, "privateEndpointConnectionOperation", &migrate.PrivateEndpointConnectionOperationArgs{
			PrivateEndpointConnectionName: pulumi.String("sakanwar1204project1634pe.bf42f8a1-09f5-4ee4-aea6-a019cc60f9d7"),
			PrivateLinkServiceConnectionState: &migrate.PrivateLinkServiceConnectionStateArgs{
				ActionsRequired: pulumi.String(""),
				Status:          pulumi.String("Approved"),
			},
			ProjectName:       pulumi.String("sakanwar1204project"),
			ResourceGroupName: pulumi.String("sakanwar"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
