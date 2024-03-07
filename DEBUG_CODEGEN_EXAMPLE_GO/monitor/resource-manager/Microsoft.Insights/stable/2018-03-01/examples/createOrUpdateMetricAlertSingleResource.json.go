package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewMetricAlert(ctx, "metricAlert", &insights.MetricAlertArgs{
			Actions: []insights.MetricAlertActionArgs{
				{
					ActionGroupId: pulumi.String("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
					WebHookProperties: {
						"key11": pulumi.String("value11"),
						"key12": pulumi.String("value12"),
					},
				},
			},
			AutoMitigate: pulumi.Bool(true),
			Criteria: insights.MetricAlertSingleResourceMultipleMetricCriteria{
				AllOf: []insights.MetricCriteria{
					{
						CriterionType:   "StaticThresholdCriterion",
						Dimensions:      []insights.MetricDimension{},
						MetricName:      "\\Processor(_Total)\\% Processor Time",
						Name:            "High_CPU_80",
						Operator:        "GreaterThan",
						Threshold:       80.5,
						TimeAggregation: "Average",
					},
				},
				OdataType: "Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria",
			},
			Description:         pulumi.String("This is the description of the rule1"),
			Enabled:             pulumi.Bool(true),
			EvaluationFrequency: pulumi.String("Pt1m"),
			Location:            pulumi.String("global"),
			ResourceGroupName:   pulumi.String("gigtest"),
			RuleName:            pulumi.String("chiricutin"),
			Scopes: pulumi.StringArray{
				pulumi.String("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme"),
			},
			Severity:   pulumi.Int(3),
			Tags:       nil,
			WindowSize: pulumi.String("Pt15m"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
