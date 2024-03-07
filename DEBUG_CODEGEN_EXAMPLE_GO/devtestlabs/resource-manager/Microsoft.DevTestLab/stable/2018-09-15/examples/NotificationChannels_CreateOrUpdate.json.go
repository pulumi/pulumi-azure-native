package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewNotificationChannel(ctx, "notificationChannel", &devtestlab.NotificationChannelArgs{
			Description:    pulumi.String("Integration configured for auto-shutdown"),
			EmailRecipient: pulumi.String("{email}"),
			Events: devtestlab.EventArray{
				&devtestlab.EventArgs{
					EventName: pulumi.String("AutoShutdown"),
				},
			},
			LabName:            pulumi.String("{labName}"),
			Name:               pulumi.String("{notificationChannelName}"),
			NotificationLocale: pulumi.String("en"),
			ResourceGroupName:  pulumi.String("resourceGroupName"),
			WebHookUrl:         pulumi.String("{webhookUrl}"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
