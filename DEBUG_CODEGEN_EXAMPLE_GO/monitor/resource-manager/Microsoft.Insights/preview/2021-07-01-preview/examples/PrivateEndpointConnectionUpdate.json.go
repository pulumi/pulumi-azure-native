package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &insights.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("private-endpoint-connection-name"),
			PrivateLinkServiceConnectionState: &insights.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Approved by johndoe@contoso.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("MyResourceGroup"),
			ScopeName:         pulumi.String("MyPrivateLinkScope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
