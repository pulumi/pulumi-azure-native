package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datafactory/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datafactory.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &datafactory.PrivateEndpointConnectionArgs{
			FactoryName:                   pulumi.String("exampleFactoryName"),
			PrivateEndpointConnectionName: pulumi.String("connection"),
			Properties: datafactory.RemotePrivateEndpointConnectionResponse{
				PrivateEndpoint: &datafactory.PrivateEndpointArgs{
					Id: pulumi.String("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/privateEndpoints/myPrivateEndpoint"),
				},
				PrivateLinkServiceConnectionState: &datafactory.PrivateLinkConnectionStateArgs{
					ActionsRequired: pulumi.String(""),
					Description:     pulumi.String("Approved by admin."),
					Status:          pulumi.String("Approved"),
				},
			},
			ResourceGroupName: pulumi.String("exampleResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
