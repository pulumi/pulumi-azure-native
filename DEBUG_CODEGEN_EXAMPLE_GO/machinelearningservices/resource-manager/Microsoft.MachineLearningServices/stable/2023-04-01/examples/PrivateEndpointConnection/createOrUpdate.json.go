package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &machinelearningservices.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("{privateEndpointConnectionName}"),
			PrivateLinkServiceConnectionState: &machinelearningservices.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rg-1234"),
			WorkspaceName:     pulumi.String("testworkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
