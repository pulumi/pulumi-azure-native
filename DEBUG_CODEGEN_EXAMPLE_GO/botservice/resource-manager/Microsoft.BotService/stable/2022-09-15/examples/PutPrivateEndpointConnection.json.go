package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/botservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := botservice.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &botservice.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("{privateEndpointConnectionName}"),
			PrivateLinkServiceConnectionState: &botservice.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("res7687"),
			ResourceName:      pulumi.String("sto9699"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
