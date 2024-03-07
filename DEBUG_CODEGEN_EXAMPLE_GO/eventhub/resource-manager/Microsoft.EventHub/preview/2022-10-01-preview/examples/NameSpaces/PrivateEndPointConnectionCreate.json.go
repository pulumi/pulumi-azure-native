package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &eventhub.PrivateEndpointConnectionArgs{
			NamespaceName: pulumi.String("sdk-Namespace-2924"),
			PrivateEndpoint: &eventhub.PrivateEndpointArgs{
				Id: pulumi.String("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-EventHub-8396/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-2847"),
			},
			PrivateEndpointConnectionName: pulumi.String("privateEndpointConnectionName"),
			PrivateLinkServiceConnectionState: &eventhub.ConnectionStateArgs{
				Description: pulumi.String("testing"),
				Status:      pulumi.String("Rejected"),
			},
			ProvisioningState: pulumi.String("Succeeded"),
			ResourceGroupName: pulumi.String("ArunMonocle"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
