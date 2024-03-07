package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewWebAppPrivateEndpointConnectionSlot(ctx, "webAppPrivateEndpointConnectionSlot", &web.WebAppPrivateEndpointConnectionSlotArgs{
			Name:                          pulumi.String("testSite"),
			PrivateEndpointConnectionName: pulumi.String("connection"),
			PrivateLinkServiceConnectionState: &web.PrivateLinkConnectionStateArgs{
				ActionsRequired: pulumi.String(""),
				Description:     pulumi.String("Approved by admin."),
				Status:          pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rg"),
			Slot:              pulumi.String("stage"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
