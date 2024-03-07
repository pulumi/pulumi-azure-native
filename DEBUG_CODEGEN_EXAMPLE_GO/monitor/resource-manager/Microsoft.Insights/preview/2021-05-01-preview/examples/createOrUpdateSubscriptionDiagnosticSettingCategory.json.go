package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewSubscriptionDiagnosticSetting(ctx, "subscriptionDiagnosticSetting", &insights.SubscriptionDiagnosticSettingArgs{
			EventHubAuthorizationRuleId: pulumi.String("/subscriptions/fb9f25f9-5785-4510-a38f-a62f188eb9f8/resourceGroups/montest/providers/microsoft.eventhub/namespaces/mynamespace/authorizationrules/myrule"),
			EventHubName:                pulumi.String("myeventhub"),
			Logs: insights.SubscriptionLogSettingsArray{
				&insights.SubscriptionLogSettingsArgs{
					Category: pulumi.String("Security"),
					Enabled:  pulumi.Bool(true),
				},
			},
			MarketplacePartnerId: pulumi.String("/subscriptions/abcdeabc-1234-1234-ab12-123a1234567a/resourceGroups/test-rg/providers/Microsoft.Datadog/monitors/dd1"),
			Name:                 pulumi.String("ds4"),
			StorageAccountId:     pulumi.String("/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/apptest/providers/Microsoft.Storage/storageAccounts/appteststorage1"),
			WorkspaceId:          pulumi.String(""),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
