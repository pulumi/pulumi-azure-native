package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformysql.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &dbformysql.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("private-endpoint-connection-name"),
			PrivateLinkServiceConnectionState: &dbformysql.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Approved by johndoe@contoso.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("Default"),
			ServerName:        pulumi.String("test-svr"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
