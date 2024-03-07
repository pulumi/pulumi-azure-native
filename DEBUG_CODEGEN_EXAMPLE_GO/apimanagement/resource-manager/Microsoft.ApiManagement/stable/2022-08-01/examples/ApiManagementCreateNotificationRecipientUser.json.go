package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewNotificationRecipientUser(ctx, "notificationRecipientUser", &apimanagement.NotificationRecipientUserArgs{
			NotificationName:  pulumi.String("RequestPublisherNotificationMessage"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			UserId:            pulumi.String("576823d0a40f7e74ec07d642"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
