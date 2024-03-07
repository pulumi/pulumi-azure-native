package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alertsmanagement.NewAlertProcessingRuleByName(ctx, "alertProcessingRuleByName", &alertsmanagement.AlertProcessingRuleByNameArgs{
			AlertProcessingRuleName: pulumi.String("RemoveActionGroupsMaintenanceWindow"),
			Location:                pulumi.String("Global"),
			Properties: &alertsmanagement.AlertProcessingRulePropertiesArgs{
				Actions: pulumi.Array{
					alertsmanagement.RemoveAllActionGroups{
						ActionType: "RemoveAllActionGroups",
					},
				},
				Description: pulumi.String("Removes all ActionGroups from all Alerts on VMName during the maintenance window"),
				Enabled:     pulumi.Bool(true),
				Schedule: &alertsmanagement.ScheduleArgs{
					EffectiveFrom:  pulumi.String("2021-04-15T18:00:00"),
					EffectiveUntil: pulumi.String("2021-04-15T20:00:00"),
					TimeZone:       pulumi.String("Pacific Standard Time"),
				},
				Scopes: pulumi.StringArray{
					pulumi.String("/subscriptions/subId1/resourceGroups/RGId1/providers/Microsoft.Compute/virtualMachines/VMName"),
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
