package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/notificationhubs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := notificationhubs.NewNotificationHubAuthorizationRule(ctx, "notificationHubAuthorizationRule", &notificationhubs.NotificationHubAuthorizationRuleArgs{
			AuthorizationRuleName: pulumi.String("MyManageSharedAccessKey"),
			NamespaceName:         pulumi.String("nh-sdk-ns"),
			NotificationHubName:   pulumi.String("nh-sdk-hub"),
			Properties: &notificationhubs.SharedAccessAuthorizationRulePropertiesArgs{
				Rights: pulumi.StringArray{
					pulumi.String("Listen"),
					pulumi.String("Send"),
				},
			},
			ResourceGroupName: pulumi.String("5ktrial"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
