package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/notificationhubs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := notificationhubs.NewNotificationHub(ctx, "notificationHub", &notificationhubs.NotificationHubArgs{
			Location:            pulumi.String("eastus"),
			NamespaceName:       pulumi.String("nh-sdk-ns"),
			NotificationHubName: pulumi.String("nh-sdk-hub"),
			Properties:          nil,
			ResourceGroupName:   pulumi.String("5ktrial"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
