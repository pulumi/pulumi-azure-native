package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/iotcentral/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iotcentral.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &iotcentral.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("myIoTCentralAppEndpoint"),
			PrivateLinkServiceConnectionState: &iotcentral.PrivateLinkServiceConnectionStateArgs{
				ActionsRequired: pulumi.String("None"),
				Description:     pulumi.String("Auto-approved"),
				Status:          pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("resRg"),
			ResourceName:      pulumi.String("myIoTCentralApp"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
