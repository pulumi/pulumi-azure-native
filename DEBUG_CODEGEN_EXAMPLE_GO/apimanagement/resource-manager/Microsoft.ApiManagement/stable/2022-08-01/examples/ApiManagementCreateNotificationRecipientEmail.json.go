package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewNotificationRecipientEmail(ctx, "notificationRecipientEmail", &apimanagement.NotificationRecipientEmailArgs{
			Email:             pulumi.String("foobar@live.com"),
			NotificationName:  pulumi.String("RequestPublisherNotificationMessage"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
