package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewAlertRule(ctx, "alertRule", &insights.AlertRuleArgs{
			Actions: pulumi.Array{},
			Condition: insights.ThresholdRuleCondition{
				DataSource: insights.RuleMetricDataSource{
					MetricName:  "Requests",
					OdataType:   "Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource",
					ResourceUri: "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest",
				},
				OdataType:       "Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition",
				Operator:        insights.ConditionOperatorGreaterThan,
				Threshold:       3,
				TimeAggregation: insights.TimeAggregationOperatorTotal,
				WindowSize:      "PT5M",
			},
			Description:       pulumi.String("Pura Vida"),
			IsEnabled:         pulumi.Bool(true),
			Location:          pulumi.String("West US"),
			Name:              pulumi.String("chiricutin"),
			ResourceGroupName: pulumi.String("Rac46PostSwapRG"),
			RuleName:          pulumi.String("chiricutin"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
