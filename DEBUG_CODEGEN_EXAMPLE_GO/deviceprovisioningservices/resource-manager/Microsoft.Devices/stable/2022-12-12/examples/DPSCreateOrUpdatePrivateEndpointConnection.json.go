package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devices.NewIotDpsResourcePrivateEndpointConnection(ctx, "iotDpsResourcePrivateEndpointConnection", &devices.IotDpsResourcePrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("myPrivateEndpointConnection"),
			Properties: &devices.PrivateEndpointConnectionPropertiesArgs{
				PrivateLinkServiceConnectionState: &devices.PrivateLinkServiceConnectionStateArgs{
					Description: pulumi.String("Approved by johndoe@contoso.com"),
					Status:      pulumi.String("Approved"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ResourceName:      pulumi.String("myFirstProvisioningService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
