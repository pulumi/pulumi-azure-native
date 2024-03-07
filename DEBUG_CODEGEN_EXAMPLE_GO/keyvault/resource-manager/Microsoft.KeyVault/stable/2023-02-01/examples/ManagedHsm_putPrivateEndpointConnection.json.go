package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/keyvault/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := keyvault.NewMHSMPrivateEndpointConnection(ctx, "mhsmPrivateEndpointConnection", &keyvault.MHSMPrivateEndpointConnectionArgs{
			Name:                          pulumi.String("sample-mhsm"),
			PrivateEndpointConnectionName: pulumi.String("sample-pec"),
			PrivateLinkServiceConnectionState: &keyvault.MHSMPrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("My name is Joe and I'm approving this."),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("sample-group"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
