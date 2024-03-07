package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &dbforpostgresql.PrivateEndpointConnectionArgs{
			ClusterName:                   pulumi.String("testcluster"),
			PrivateEndpointConnectionName: pulumi.String("private-endpoint-connection-name"),
			PrivateLinkServiceConnectionState: &dbforpostgresql.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Approved by johndoe@contoso.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("TestGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
