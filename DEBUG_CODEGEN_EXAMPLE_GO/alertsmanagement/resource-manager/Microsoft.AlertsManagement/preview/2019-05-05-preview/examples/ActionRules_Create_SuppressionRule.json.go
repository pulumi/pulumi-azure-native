package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alertsmanagement.NewActionRuleByName(ctx, "actionRuleByName", &alertsmanagement.ActionRuleByNameArgs{
			ActionRuleName: pulumi.String("DailySuppression"),
			Location:       pulumi.String("Global"),
			Properties: alertsmanagement.Suppression{
				Conditions: alertsmanagement.Conditions{
					MonitorCondition: alertsmanagement.Condition{
						Operator: "Equals",
						Values: []string{
							"Fired",
						},
					},
					MonitorService: alertsmanagement.Condition{
						Operator: "Equals",
						Values: []string{
							"Platform",
							"Application Insights",
						},
					},
					Severity: alertsmanagement.Condition{
						Operator: "Equals",
						Values: []string{
							"Sev0",
							"Sev2",
						},
					},
					TargetResourceType: alertsmanagement.Condition{
						Operator: "NotEquals",
						Values: []string{
							"Microsoft.Compute/VirtualMachines",
						},
					},
				},
				Description: "Action rule on resource group for daily suppression",
				Scope: alertsmanagement.Scope{
					ScopeType: "ResourceGroup",
					Values: []string{
						"/subscriptions/1e3ff1c0-771a-4119-a03b-be82a51e232d/resourceGroups/alertscorrelationrg",
					},
				},
				Status: "Enabled",
				SuppressionConfig: alertsmanagement.SuppressionConfig{
					RecurrenceType: "Daily",
					Schedule: alertsmanagement.SuppressionSchedule{
						EndDate:   "12/18/2018",
						EndTime:   "14:00:00",
						StartDate: "12/09/2018",
						StartTime: "06:00:00",
					},
				},
				Type: "Suppression",
			},
			ResourceGroupName: pulumi.String("alertscorrelationrg"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
