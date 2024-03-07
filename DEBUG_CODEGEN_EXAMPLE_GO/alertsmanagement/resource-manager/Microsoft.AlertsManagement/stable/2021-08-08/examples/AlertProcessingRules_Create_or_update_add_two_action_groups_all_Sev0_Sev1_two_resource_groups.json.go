package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alertsmanagement.NewAlertProcessingRuleByName(ctx, "alertProcessingRuleByName", &alertsmanagement.AlertProcessingRuleByNameArgs{
			AlertProcessingRuleName: pulumi.String("AddActionGroupsBySeverity"),
			Location:                pulumi.String("Global"),
			Properties: &alertsmanagement.AlertProcessingRulePropertiesArgs{
				Actions: pulumi.Array{
					alertsmanagement.AddActionGroups{
						ActionGroupIds: []string{
							"/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId1",
							"/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/AGId2",
						},
						ActionType: "AddActionGroups",
					},
				},
				Conditions: alertsmanagement.ConditionArray{
					&alertsmanagement.ConditionArgs{
						Field:    pulumi.String("Severity"),
						Operator: pulumi.String("Equals"),
						Values: pulumi.StringArray{
							pulumi.String("sev0"),
							pulumi.String("sev1"),
						},
					},
				},
				Description: pulumi.String("Add AGId1 and AGId2 to all Sev0 and Sev1 alerts in these resourceGroups"),
				Enabled:     pulumi.Bool(true),
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
