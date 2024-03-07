package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alertsmanagement.NewAlertProcessingRuleByName(ctx, "alertProcessingRuleByName", &alertsmanagement.AlertProcessingRuleByNameArgs{
			AlertProcessingRuleName: pulumi.String("AddActionGroupToSubscription"),
			Location:                pulumi.String("Global"),
			Properties: &alertsmanagement.AlertProcessingRulePropertiesArgs{
				Actions: pulumi.Array{
					alertsmanagement.AddActionGroups{
						ActionGroupIds: []string{
							"/subscriptions/subId1/resourcegroups/RGId1/providers/microsoft.insights/actiongroups/ActionGroup1",
						},
						ActionType: "AddActionGroups",
					},
				},
				Description: pulumi.String("Add ActionGroup1 to all alerts in the subscription"),
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
