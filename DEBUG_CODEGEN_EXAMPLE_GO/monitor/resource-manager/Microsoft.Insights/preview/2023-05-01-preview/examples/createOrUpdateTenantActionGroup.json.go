package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewTenantActionGroup(ctx, "tenantActionGroup", &insights.TenantActionGroupArgs{
			AzureAppPushReceivers: []insights.AzureAppPushReceiverArgs{
				{
					EmailAddress: pulumi.String("johndoe@email.com"),
					Name:         pulumi.String("Sample azureAppPush"),
				},
			},
			EmailReceivers: []insights.EmailReceiverArgs{
				{
					EmailAddress:         pulumi.String("johndoe@email.com"),
					Name:                 pulumi.String("John Doe's email"),
					UseCommonAlertSchema: pulumi.Bool(false),
				},
				{
					EmailAddress:         pulumi.String("janesmith@email.com"),
					Name:                 pulumi.String("Jane Smith's email"),
					UseCommonAlertSchema: pulumi.Bool(true),
				},
			},
			Enabled:           pulumi.Bool(true),
			GroupShortName:    pulumi.String("sample"),
			Location:          pulumi.String("Global"),
			ManagementGroupId: pulumi.String("72f988bf-86f1-41af-91ab-2d7cd011db47"),
			SmsReceivers: []insights.SmsReceiverArgs{
				{
					CountryCode: pulumi.String("1"),
					Name:        pulumi.String("John Doe's mobile"),
					PhoneNumber: pulumi.String("2062022299"),
				},
				{
					CountryCode: pulumi.String("1"),
					Name:        pulumi.String("Jane Smith's mobile"),
					PhoneNumber: pulumi.String("0987654321"),
				},
			},
			Tags:                  nil,
			TenantActionGroupName: pulumi.String("testTenantActionGroup"),
			VoiceReceivers: []insights.VoiceReceiverArgs{
				{
					CountryCode: pulumi.String("1"),
					Name:        pulumi.String("Sample voice"),
					PhoneNumber: pulumi.String("2062022299"),
				},
			},
			WebhookReceivers: []insights.WebhookReceiverArgs{
				{
					Name:                 pulumi.String("Sample webhook 1"),
					ServiceUri:           pulumi.String("http://www.example.com/webhook1"),
					UseCommonAlertSchema: pulumi.Bool(true),
				},
				{
					IdentifierUri:        pulumi.String("http://someidentifier/d7811ba3-7996-4a93-99b6-6b2f3f355f8a"),
					Name:                 pulumi.String("Sample webhook 2"),
					ObjectId:             pulumi.String("d3bb868c-fe44-452c-aa26-769a6538c808"),
					ServiceUri:           pulumi.String("http://www.example.com/webhook2"),
					TenantId:             pulumi.String("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
					UseAadAuth:           pulumi.Bool(true),
					UseCommonAlertSchema: pulumi.Bool(true),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
