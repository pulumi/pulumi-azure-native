package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/powerbi/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := powerbi.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &powerbi.PrivateEndpointConnectionArgs{
			AzureResourceName: pulumi.String("azureResourceName"),
			PrivateEndpoint: &powerbi.PrivateEndpointArgs{
				Id: pulumi.String("/subscriptions/a0020869-4d28-422a-89f4-c2413130d73c/resourceGroups/resourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpointName"),
			},
			PrivateEndpointName: pulumi.String("myPrivateEndpointName"),
			PrivateLinkServiceConnectionState: &powerbi.ConnectionStateArgs{
				ActionsRequired: pulumi.String("None"),
				Description:     pulumi.String(""),
				Status:          pulumi.String("Approved "),
			},
			ResourceGroupName: pulumi.String("resourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
