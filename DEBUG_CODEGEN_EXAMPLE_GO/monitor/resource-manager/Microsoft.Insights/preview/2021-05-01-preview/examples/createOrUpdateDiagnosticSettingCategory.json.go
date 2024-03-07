package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewDiagnosticSetting(ctx, "diagnosticSetting", &insights.DiagnosticSettingArgs{
			EventHubAuthorizationRuleId: pulumi.String("/subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourceGroups/montest/providers/microsoft.eventhub/namespaces/mynamespace/authorizationrules/myrule"),
			EventHubName:                pulumi.String("myeventhub"),
			LogAnalyticsDestinationType: pulumi.String("Dedicated"),
			Logs: []insights.LogSettingsArgs{
				{
					Category: pulumi.String("WorkflowRuntime"),
					Enabled:  pulumi.Bool(true),
					RetentionPolicy: {
						Days:    pulumi.Int(0),
						Enabled: pulumi.Bool(false),
					},
				},
			},
			MarketplacePartnerId: pulumi.String("/subscriptions/abcdeabc-1234-1234-ab12-123a1234567a/resourceGroups/test-rg/providers/Microsoft.Datadog/monitors/dd1"),
			Metrics: []insights.MetricSettingsArgs{
				{
					Category: pulumi.String("WorkflowMetrics"),
					Enabled:  pulumi.Bool(true),
					RetentionPolicy: {
						Days:    pulumi.Int(0),
						Enabled: pulumi.Bool(false),
					},
				},
			},
			Name:             pulumi.String("mysetting"),
			ResourceUri:      pulumi.String("subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6"),
			StorageAccountId: pulumi.String("/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/apptest/providers/Microsoft.Storage/storageAccounts/appteststorage1"),
			WorkspaceId:      pulumi.String(""),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
