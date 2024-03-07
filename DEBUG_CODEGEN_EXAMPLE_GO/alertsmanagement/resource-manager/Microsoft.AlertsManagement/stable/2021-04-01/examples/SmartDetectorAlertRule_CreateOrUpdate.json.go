package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alertsmanagement.NewSmartDetectorAlertRule(ctx, "smartDetectorAlertRule", &alertsmanagement.SmartDetectorAlertRuleArgs{
			ActionGroups: &alertsmanagement.ActionGroupsInformationArgs{
				CustomEmailSubject:   pulumi.String("My custom email subject"),
				CustomWebhookPayload: pulumi.String("{\"AlertRuleName\":\"#alertrulename\"}"),
				GroupIds: pulumi.StringArray{
					pulumi.String("/subscriptions/b368ca2f-e298-46b7-b0ab-012281956afa/resourcegroups/actionGroups/providers/microsoft.insights/actiongroups/MyActionGroup"),
				},
			},
			AlertRuleName: pulumi.String("MyAlertRule"),
			Description:   pulumi.String("Sample smart detector alert rule description"),
			Detector: &alertsmanagement.DetectorArgs{
				Id: pulumi.String("VMMemoryLeak"),
			},
			Frequency:         pulumi.String("PT5M"),
			ResourceGroupName: pulumi.String("MyAlertRules"),
			Scope: pulumi.StringArray{
				pulumi.String("/subscriptions/b368ca2f-e298-46b7-b0ab-012281956afa/resourceGroups/MyVms/providers/Microsoft.Compute/virtualMachines/vm1"),
			},
			Severity: pulumi.String("Sev3"),
			State:    pulumi.String("Enabled"),
			Throttling: &alertsmanagement.ThrottlingInformationArgs{
				Duration: pulumi.String("PT20M"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
