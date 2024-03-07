package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/attestation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := attestation.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &attestation.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("{privateEndpointConnectionName}"),
			PrivateLinkServiceConnectionState: &attestation.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Auto-Approved"),
				Status:      pulumi.String("Approved"),
			},
			ProviderName:      pulumi.String("sto9699"),
			ResourceGroupName: pulumi.String("res7687"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
