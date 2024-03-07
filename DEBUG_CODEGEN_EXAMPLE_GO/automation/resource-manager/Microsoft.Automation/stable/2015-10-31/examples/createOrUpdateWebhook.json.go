package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewWebhook(ctx, "webhook", &automation.WebhookArgs{
			AutomationAccountName: pulumi.String("myAutomationAccount33"),
			ExpiryTime:            pulumi.String("2018-03-29T22:18:13.7002872Z"),
			IsEnabled:             pulumi.Bool(true),
			Name:                  pulumi.String("TestWebhook"),
			ResourceGroupName:     pulumi.String("rg"),
			Runbook: &automation.RunbookAssociationPropertyArgs{
				Name: pulumi.String("TestRunbook"),
			},
			Uri:         pulumi.String("<uri>"),
			WebhookName: pulumi.String("TestWebhook"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
