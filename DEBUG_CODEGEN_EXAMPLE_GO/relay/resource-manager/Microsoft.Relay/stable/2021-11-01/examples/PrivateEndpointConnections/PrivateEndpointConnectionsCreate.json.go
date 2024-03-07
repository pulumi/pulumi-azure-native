package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/relay/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := relay.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &relay.PrivateEndpointConnectionArgs{
			NamespaceName: pulumi.String("example-RelayNamespace-5849"),
			PrivateEndpoint: &relay.PrivateEndpointArgs{
				Id: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Network/privateEndpoints/ali-relay-pve-1"),
			},
			PrivateEndpointConnectionName: pulumi.String("{privateEndpointConnection name}"),
			PrivateLinkServiceConnectionState: &relay.ConnectionStateArgs{
				Description: pulumi.String("You may pass"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("resourcegroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
