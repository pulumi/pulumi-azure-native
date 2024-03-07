package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kusto/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kusto.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &kusto.PrivateEndpointConnectionArgs{
			ClusterName:                   pulumi.String("kustoclusterrptest4"),
			PrivateEndpointConnectionName: pulumi.String("privateEndpointTest"),
			PrivateLinkServiceConnectionState: &kusto.PrivateLinkServiceConnectionStatePropertyArgs{
				Description: pulumi.String("Approved by johndoe@contoso.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("kustorptest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
