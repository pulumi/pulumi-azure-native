package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcompute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcompute.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &hybridcompute.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("private-endpoint-connection-name"),
			Properties: hybridcompute.PrivateEndpointConnectionPropertiesResponse{
				PrivateLinkServiceConnectionState: &hybridcompute.PrivateLinkServiceConnectionStatePropertyArgs{
					Description: pulumi.String("Approved by johndoe@contoso.com"),
					Status:      pulumi.String("Approved"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ScopeName:         pulumi.String("myPrivateLinkScope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
