package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewManagementGroupDiagnosticSetting(ctx, "managementGroupDiagnosticSetting", &insights.ManagementGroupDiagnosticSettingArgs{
			EventHubAuthorizationRuleId: pulumi.String("/subscriptions/fb9f25f9-5785-4510-a38f-a62f188eb9f8/resourceGroups/montest/providers/microsoft.eventhub/namespaces/mynamespace/authorizationrules/myrule"),
			EventHubName:                pulumi.String("myeventhub"),
			Logs: insights.ManagementGroupLogSettingsArray{
				&insights.ManagementGroupLogSettingsArgs{
					Category: pulumi.String("Administrative"),
					Enabled:  pulumi.Bool(true),
				},
				&insights.ManagementGroupLogSettingsArgs{
					Category: pulumi.String("Policy"),
					Enabled:  pulumi.Bool(true),
				},
			},
			ManagementGroupId:    pulumi.String("testChildMG7"),
			MarketplacePartnerId: pulumi.String("/subscriptions/abcdeabc-1234-1234-ab12-123a1234567a/resourceGroups/test-rg/providers/Microsoft.Datadog/monitors/dd1"),
			Name:                 pulumi.String("setting1"),
			StorageAccountId:     pulumi.String("/subscriptions/bfaef57f-297e-4210-bfe5-27c18cc671f7/resourceGroups/FuncAppRunners/providers/Microsoft.Storage/storageAccounts/testpersonalb6a5"),
			WorkspaceId:          pulumi.String("/subscriptions/9cf7cc0a-0ba1-4624-bc82-97e1ee25dc45/resourceGroups/mgTest/providers/Microsoft.OperationalInsights/workspaces/mgTestWorkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
