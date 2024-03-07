package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/powerplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := powerplatform.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &powerplatform.PrivateEndpointConnectionArgs{
			EnterprisePolicyName:          pulumi.String("ddb1"),
			PrivateEndpointConnectionName: pulumi.String("privateEndpointConnectionName"),
			PrivateLinkServiceConnectionState: &powerplatform.PrivateLinkServiceConnectionStateArgs{
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
