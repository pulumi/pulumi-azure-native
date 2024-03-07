package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/aadiam/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aadiam.NewDiagnosticSetting(ctx, "diagnosticSetting", &aadiam.DiagnosticSettingArgs{
			EventHubAuthorizationRuleId: pulumi.String("/subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourceGroups/montest/providers/microsoft.eventhub/namespaces/mynamespace/eventhubs/myeventhub/authorizationrules/myrule"),
			EventHubName:                pulumi.String("myeventhub"),
			Logs: []aadiam.LogSettingsArgs{
				{
					Category: pulumi.String("AuditLogs"),
					Enabled:  pulumi.Bool(true),
					RetentionPolicy: {
						Days:    pulumi.Int(0),
						Enabled: pulumi.Bool(false),
					},
				},
			},
			Name:             pulumi.String("mysetting"),
			StorageAccountId: pulumi.String("/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/apptest/providers/Microsoft.Storage/storageAccounts/appteststorage1"),
			WorkspaceId:      pulumi.String(""),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
