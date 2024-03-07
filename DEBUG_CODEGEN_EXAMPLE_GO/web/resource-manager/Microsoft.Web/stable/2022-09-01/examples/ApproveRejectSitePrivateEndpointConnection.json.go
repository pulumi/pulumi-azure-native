package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewWebAppPrivateEndpointConnection(ctx, "webAppPrivateEndpointConnection", &web.WebAppPrivateEndpointConnectionArgs{
			Name:                          pulumi.String("testSite"),
			PrivateEndpointConnectionName: pulumi.String("connection"),
			PrivateLinkServiceConnectionState: &web.PrivateLinkConnectionStateArgs{
				ActionsRequired: pulumi.String(""),
				Description:     pulumi.String("Approved by admin."),
				Status:          pulumi.String("Approved"),
			},
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
