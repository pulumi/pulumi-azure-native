package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alertsmanagement.NewAlertProcessingRuleByName(ctx, "alertProcessingRuleByName", &alertsmanagement.AlertProcessingRuleByNameArgs{
			AlertProcessingRuleName: pulumi.String("RemoveActionGroupsRecurringMaintenance"),
			Location:                pulumi.String("Global"),
			Properties: &alertsmanagement.AlertProcessingRulePropertiesArgs{
				Actions: pulumi.Array{
					alertsmanagement.RemoveAllActionGroups{
						ActionType: "RemoveAllActionGroups",
					},
				},
				Conditions: alertsmanagement.ConditionArray{
					&alertsmanagement.ConditionArgs{
						Field:    pulumi.String("TargetResourceType"),
						Operator: pulumi.String("Equals"),
						Values: pulumi.StringArray{
							pulumi.String("microsoft.compute/virtualmachines"),
						},
					},
				},
				Description: pulumi.String("Remove all ActionGroups from all Vitual machine Alerts during the recurring maintenance"),
				Enabled:     pulumi.Bool(true),
				Schedule: &alertsmanagement.ScheduleArgs{
					Recurrences: pulumi.Array{
						alertsmanagement.WeeklyRecurrence{
							DaysOfWeek: []alertsmanagement.DaysOfWeek{
								"Saturday",
								"Sunday",
							},
							EndTime:        "04:00:00",
							RecurrenceType: "Weekly",
							StartTime:      "22:00:00",
						},
					},
					TimeZone: pulumi.String("India Standard Time"),
				},
				Scopes: pulumi.StringArray{
					pulumi.String("/subscriptions/subId1/resourceGroups/RGId1"),
					pulumi.String("/subscriptions/subId1/resourceGroups/RGId2"),
				},
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
