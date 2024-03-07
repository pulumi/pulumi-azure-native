package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appconfiguration.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &appconfiguration.PrivateEndpointConnectionArgs{
			ConfigStoreName:               pulumi.String("contoso"),
			PrivateEndpointConnectionName: pulumi.String("myConnection"),
			PrivateLinkServiceConnectionState: &appconfiguration.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
