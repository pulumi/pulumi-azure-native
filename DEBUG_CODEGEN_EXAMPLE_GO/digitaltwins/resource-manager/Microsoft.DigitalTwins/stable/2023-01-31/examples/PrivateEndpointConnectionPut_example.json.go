package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/digitaltwins/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := digitaltwins.NewPrivateEndpointConnection(ctx, "privateEndpointConnection", &digitaltwins.PrivateEndpointConnectionArgs{
			PrivateEndpointConnectionName: pulumi.String("myPrivateConnection"),
			Properties: &digitaltwins.ConnectionPropertiesArgs{
				PrivateLinkServiceConnectionState: &digitaltwins.ConnectionPropertiesPrivateLinkServiceConnectionStateArgs{
					Description: pulumi.String("Approved by johndoe@company.com."),
					Status:      pulumi.String("Approved"),
				},
			},
			ResourceGroupName: pulumi.String("resRg"),
			ResourceName:      pulumi.String("myDigitalTwinsService"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
