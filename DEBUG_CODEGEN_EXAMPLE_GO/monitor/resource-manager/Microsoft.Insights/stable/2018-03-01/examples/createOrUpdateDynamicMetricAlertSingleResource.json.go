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
					ActionGroupId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
					WebHookProperties: {
						"key11": pulumi.String("value11"),
						"key12": pulumi.String("value12"),
					},
				},
			},
			AutoMitigate: pulumi.Bool(true),
			Criteria: insights.MetricAlertMultipleResourceMultipleMetricCriteria{
				AllOf: []interface{}{
					insights.DynamicMetricCriteria{
						AlertSensitivity: "Medium",
						CriterionType:    "DynamicThresholdCriterion",
						Dimensions:       []insights.MetricDimension{},
						FailingPeriods: insights.DynamicThresholdFailingPeriods{
							MinFailingPeriodsToAlert:  4,
							NumberOfEvaluationPeriods: 4,
						},
						IgnoreDataBefore: "2019-04-04T21:00:00.000Z",
						MetricName:       "Percentage CPU",
						MetricNamespace:  "microsoft.compute/virtualmachines",
						Name:             "High_CPU_80",
						Operator:         "GreaterOrLessThan",
						TimeAggregation:  "Average",
					},
				},
				OdataType: "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
			},
			Description:         pulumi.String("This is the description of the rule1"),
			Enabled:             pulumi.Bool(true),
			EvaluationFrequency: pulumi.String("PT1M"),
			Location:            pulumi.String("global"),
			ResourceGroupName:   pulumi.String("gigtest"),
			RuleName:            pulumi.String("chiricutin"),
			Scopes: pulumi.StringArray{
				pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme"),
			},
			Severity:   pulumi.Int(3),
			Tags:       nil,
			WindowSize: pulumi.String("PT15M"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
