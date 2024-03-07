package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &documentdb.PrivateEndpointConnectionArgs{
			AccountName:                   pulumi.String("ddb1"),
			PrivateEndpointConnectionName: pulumi.String("privateEndpointConnectionName"),
			PrivateLinkServiceConnectionState: &documentdb.PrivateLinkServiceConnectionStatePropertyArgs{
				Description: pulumi.String("Approved by johndoe@contoso.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
