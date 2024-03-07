package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewActivityLogAlert(ctx, "activityLogAlert", &insights.ActivityLogAlertArgs{
			Actions: &insights.ActionListArgs{
				ActionGroups: insights.ActionGroupTypeArray{
					&insights.ActionGroupTypeArgs{
						ActionGroupId: pulumi.String("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
						WebhookProperties: pulumi.StringMap{
							"sampleWebhookProperty": pulumi.String("SamplePropertyValue"),
						},
					},
				},
			},
			ActivityLogAlertName: pulumi.String("SampleActivityLogAlertRuleWithContainsAny"),
			Condition: &insights.AlertRuleAllOfConditionArgs{
				AllOf: insights.AlertRuleAnyOfOrLeafConditionArray{
					&insights.AlertRuleAnyOfOrLeafConditionArgs{
						Equals: pulumi.String("ServiceHealth"),
						Field:  pulumi.String("category"),
					},
					&insights.AlertRuleAnyOfOrLeafConditionArgs{
						ContainsAny: pulumi.StringArray{
							pulumi.String("North Europe"),
							pulumi.String("West Europe"),
						},
						Field: pulumi.String("properties.impactedServices[*].ImpactedRegions[*].RegionName"),
					},
				},
			},
			Description:       pulumi.String("Description of sample Activity Log Alert rule with 'containsAny'."),
			Enabled:           pulumi.Bool(true),
			Location:          pulumi.String("Global"),
			ResourceGroupName: pulumi.String("MyResourceGroup"),
			Scopes: pulumi.StringArray{
				pulumi.String("subscriptions/187f412d-1758-44d9-b052-169e2564721d"),
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
