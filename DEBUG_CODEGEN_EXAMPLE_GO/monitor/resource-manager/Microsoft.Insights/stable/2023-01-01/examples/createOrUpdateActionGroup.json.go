package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewActionGroup(ctx, "actionGroup", &insights.ActionGroupArgs{
			ActionGroupName: pulumi.String("SampleActionGroup"),
			ArmRoleReceivers: insights.ArmRoleReceiverArray{
				&insights.ArmRoleReceiverArgs{
					Name:                 pulumi.String("Sample armRole"),
					RoleId:               pulumi.String("8e3af657-a8ff-443c-a75c-2fe8c4bcb635"),
					UseCommonAlertSchema: pulumi.Bool(true),
				},
			},
			AutomationRunbookReceivers: insights.AutomationRunbookReceiverArray{
				&insights.AutomationRunbookReceiverArgs{
					AutomationAccountId:  pulumi.String("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest"),
					IsGlobalRunbook:      pulumi.Bool(false),
					Name:                 pulumi.String("testRunbook"),
					RunbookName:          pulumi.String("Sample runbook"),
					ServiceUri:           pulumi.String("<serviceUri>"),
					UseCommonAlertSchema: pulumi.Bool(true),
					WebhookResourceId:    pulumi.String("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest/webhooks/Alert1510184037084"),
				},
			},
			AzureAppPushReceivers: insights.AzureAppPushReceiverArray{
				&insights.AzureAppPushReceiverArgs{
					EmailAddress: pulumi.String("johndoe@email.com"),
					Name:         pulumi.String("Sample azureAppPush"),
				},
			},
			AzureFunctionReceivers: insights.AzureFunctionReceiverArray{
				&insights.AzureFunctionReceiverArgs{
					FunctionAppResourceId: pulumi.String("/subscriptions/5def922a-3ed4-49c1-b9fd-05ec533819a3/resourceGroups/aznsTest/providers/Microsoft.Web/sites/testFunctionApp"),
					FunctionName:          pulumi.String("HttpTriggerCSharp1"),
					HttpTriggerUrl:        pulumi.String("http://test.me"),
					Name:                  pulumi.String("Sample azureFunction"),
					UseCommonAlertSchema:  pulumi.Bool(true),
				},
			},
			EmailReceivers: insights.EmailReceiverArray{
				&insights.EmailReceiverArgs{
					EmailAddress:         pulumi.String("johndoe@email.com"),
					Name:                 pulumi.String("John Doe's email"),
					UseCommonAlertSchema: pulumi.Bool(false),
				},
				&insights.EmailReceiverArgs{
					EmailAddress:         pulumi.String("janesmith@email.com"),
					Name:                 pulumi.String("Jane Smith's email"),
					UseCommonAlertSchema: pulumi.Bool(true),
				},
			},
			Enabled: pulumi.Bool(true),
			EventHubReceivers: insights.EventHubReceiverArray{
				&insights.EventHubReceiverArgs{
					EventHubName:      pulumi.String("testEventHub"),
					EventHubNameSpace: pulumi.String("testEventHubNameSpace"),
					Name:              pulumi.String("Sample eventHub"),
					SubscriptionId:    pulumi.String("187f412d-1758-44d9-b052-169e2564721d"),
					TenantId:          pulumi.String("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
				},
			},
			GroupShortName: pulumi.String("sample"),
			ItsmReceivers: insights.ItsmReceiverArray{
				&insights.ItsmReceiverArgs{
					ConnectionId:        pulumi.String("a3b9076c-ce8e-434e-85b4-aff10cb3c8f1"),
					Name:                pulumi.String("Sample itsm"),
					Region:              pulumi.String("westcentralus"),
					TicketConfiguration: pulumi.String("{\"PayloadRevision\":0,\"WorkItemType\":\"Incident\",\"UseTemplate\":false,\"WorkItemData\":\"{}\",\"CreateOneWIPerCI\":false}"),
					WorkspaceId:         pulumi.String("5def922a-3ed4-49c1-b9fd-05ec533819a3|55dfd1f8-7e59-4f89-bf56-4c82f5ace23c"),
				},
			},
			Location: pulumi.String("Global"),
			LogicAppReceivers: insights.LogicAppReceiverArray{
				&insights.LogicAppReceiverArgs{
					CallbackUrl:          pulumi.String("https://prod-27.northcentralus.logic.azure.com/workflows/68e572e818e5457ba898763b7db90877/triggers/manual/paths/invoke/azns/test?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=Abpsb72UYJxPPvmDo937uzofupO5r_vIeWEx7KVHo7w"),
					Name:                 pulumi.String("Sample logicApp"),
					ResourceId:           pulumi.String("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/LogicApp/providers/Microsoft.Logic/workflows/testLogicApp"),
					UseCommonAlertSchema: pulumi.Bool(false),
				},
			},
			ResourceGroupName: pulumi.String("Default-NotificationRules"),
			SmsReceivers: insights.SmsReceiverArray{
				&insights.SmsReceiverArgs{
					CountryCode: pulumi.String("1"),
					Name:        pulumi.String("John Doe's mobile"),
					PhoneNumber: pulumi.String("1234567890"),
				},
				&insights.SmsReceiverArgs{
					CountryCode: pulumi.String("1"),
					Name:        pulumi.String("Jane Smith's mobile"),
					PhoneNumber: pulumi.String("0987654321"),
				},
			},
			Tags: nil,
			VoiceReceivers: insights.VoiceReceiverArray{
				&insights.VoiceReceiverArgs{
					CountryCode: pulumi.String("1"),
					Name:        pulumi.String("Sample voice"),
					PhoneNumber: pulumi.String("1234567890"),
				},
			},
			WebhookReceivers: insights.WebhookReceiverArray{
				&insights.WebhookReceiverArgs{
					Name:                 pulumi.String("Sample webhook 1"),
					ServiceUri:           pulumi.String("http://www.example.com/webhook1"),
					UseCommonAlertSchema: pulumi.Bool(true),
				},
				&insights.WebhookReceiverArgs{
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
