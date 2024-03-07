package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/signalrservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := signalrservice.NewSignalRPrivateEndpointConnection(ctx, "signalRPrivateEndpointConnection", &signalrservice.SignalRPrivateEndpointConnectionArgs{
			PrivateEndpoint: &signalrservice.PrivateEndpointArgs{
				Id: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
			},
			PrivateEndpointConnectionName: pulumi.String("mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"),
			PrivateLinkServiceConnectionState: &signalrservice.PrivateLinkServiceConnectionStateArgs{
				ActionsRequired: pulumi.String("None"),
				Status:          pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourceName:      pulumi.String("mySignalRService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
