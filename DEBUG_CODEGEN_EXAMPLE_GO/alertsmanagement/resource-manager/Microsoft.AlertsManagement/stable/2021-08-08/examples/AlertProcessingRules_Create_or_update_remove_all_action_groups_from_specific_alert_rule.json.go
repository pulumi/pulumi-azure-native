package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alertsmanagement.NewAlertProcessingRuleByName(ctx, "alertProcessingRuleByName", &alertsmanagement.AlertProcessingRuleByNameArgs{
			AlertProcessingRuleName: pulumi.String("RemoveActionGroupsSpecificAlertRule"),
			Location:                pulumi.String("Global"),
			Properties: alertsmanagement.AlertProcessingRulePropertiesResponse{
				Actions: pulumi.Array{
					alertsmanagement.RemoveAllActionGroups{
						ActionType: "RemoveAllActionGroups",
					},
				},
				Conditions: alertsmanagement.ConditionArray{
					&alertsmanagement.ConditionArgs{
						Field:    pulumi.String("AlertRuleId"),
						Operator: pulumi.String("Equals"),
						Values: pulumi.StringArray{
							pulumi.String("/subscriptions/suubId1/resourceGroups/Rgid2/providers/microsoft.insights/activityLogAlerts/RuleName"),
						},
					},
				},
				Description: pulumi.String("Removes all ActionGroups from all Alerts that fire on above AlertRule"),
				Enabled:     pulumi.Bool(true),
				Scopes: pulumi.StringArray{
					pulumi.String("/subscriptions/subId1"),
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
