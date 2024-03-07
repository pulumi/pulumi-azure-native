package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewManagedInstancePrivateEndpointConnection(ctx, "managedInstancePrivateEndpointConnection", &sql.ManagedInstancePrivateEndpointConnectionArgs{
			ManagedInstanceName:           pulumi.String("test-cl"),
			PrivateEndpointConnectionName: pulumi.String("private-endpoint-connection-name"),
			PrivateLinkServiceConnectionState: &sql.ManagedInstancePrivateLinkServiceConnectionStatePropertyArgs{
				Description: pulumi.String("Approved by johndoe@contoso.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("Default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
