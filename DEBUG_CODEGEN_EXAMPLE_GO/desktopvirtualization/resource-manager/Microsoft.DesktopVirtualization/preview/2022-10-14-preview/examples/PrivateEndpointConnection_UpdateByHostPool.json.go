package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/desktopvirtualization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := desktopvirtualization.NewPrivateEndpointConnectionByHostPool(ctx, "privateEndpointConnectionByHostPool", &desktopvirtualization.PrivateEndpointConnectionByHostPoolArgs{
			HostPoolName:                  pulumi.String("hostPool1"),
			PrivateEndpointConnectionName: pulumi.String("hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b"),
			PrivateLinkServiceConnectionState: &desktopvirtualization.PrivateLinkServiceConnectionStateArgs{
				ActionsRequired: pulumi.String("None"),
				Description:     pulumi.String("Approved by admin@consoto.com"),
				Status:          pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("resourceGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
