package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/agfoodplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := agfoodplatform.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &agfoodplatform.PrivateEndpointConnectionArgs{
			DataManagerForAgricultureResourceName: pulumi.String("examples-farmbeatsResourceName"),
			PrivateEndpointConnectionName:         pulumi.String("privateEndpointConnectionName"),
			PrivateLinkServiceConnectionState: &agfoodplatform.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Approved by johndoe@contoso.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("examples-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
