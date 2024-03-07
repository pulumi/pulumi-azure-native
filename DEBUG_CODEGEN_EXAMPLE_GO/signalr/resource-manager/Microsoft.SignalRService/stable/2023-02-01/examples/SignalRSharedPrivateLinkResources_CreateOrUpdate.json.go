package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/signalrservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := signalrservice.NewSignalRSharedPrivateLinkResource(ctx, "signalRSharedPrivateLinkResource", &signalrservice.SignalRSharedPrivateLinkResourceArgs{
			GroupId:                       pulumi.String("sites"),
			PrivateLinkResourceId:         pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Web/sites/myWebApp"),
			RequestMessage:                pulumi.String("Please approve"),
			ResourceGroupName:             pulumi.String("myResourceGroup"),
			ResourceName:                  pulumi.String("mySignalRService"),
			SharedPrivateLinkResourceName: pulumi.String("upstream"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
