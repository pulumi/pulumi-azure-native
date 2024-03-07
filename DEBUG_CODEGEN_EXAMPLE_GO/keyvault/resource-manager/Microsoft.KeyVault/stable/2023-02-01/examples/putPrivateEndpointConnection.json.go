package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/keyvault/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := keyvault.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &keyvault.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("sample-pec"),
			PrivateLinkServiceConnectionState: &keyvault.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("My name is Joe and I'm approving this."),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("sample-group"),
			VaultName:         pulumi.String("sample-vault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
