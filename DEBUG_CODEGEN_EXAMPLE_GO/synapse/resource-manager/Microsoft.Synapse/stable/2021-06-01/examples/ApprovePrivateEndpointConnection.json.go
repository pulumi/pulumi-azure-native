package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/synapse/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synapse.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &synapse.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("ExamplePrivateEndpointConnection"),
			PrivateLinkServiceConnectionState: &synapse.PrivateLinkServiceConnectionStateArgs{
				Description: pulumi.String("Approved by abc@example.com"),
				Status:      pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("ExampleResourceGroup"),
			WorkspaceName:     pulumi.String("ExampleWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
