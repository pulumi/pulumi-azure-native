package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/security/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := security.NewSecurityContact(ctx, "securityContact", &security.SecurityContactArgs{
			AlertNotifications: &security.SecurityContactPropertiesAlertNotificationsArgs{
				MinimalSeverity: pulumi.String("Low"),
				State:           pulumi.String("On"),
			},
			Emails: pulumi.String("john@contoso.com;jane@contoso.com"),
			NotificationsByRole: &security.SecurityContactPropertiesNotificationsByRoleArgs{
				Roles: pulumi.StringArray{
					pulumi.String("Owner"),
				},
				State: pulumi.String("On"),
			},
			Phone:               pulumi.String("(214)275-4038"),
			SecurityContactName: pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
