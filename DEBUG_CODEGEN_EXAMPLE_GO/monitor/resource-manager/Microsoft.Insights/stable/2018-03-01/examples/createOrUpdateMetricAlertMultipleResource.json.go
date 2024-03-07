package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewMetricAlert(ctx, "metricAlert", &insights.MetricAlertArgs{
			Actions: insights.MetricAlertActionArray{
				&insights.MetricAlertActionArgs{
					ActionGroupId: pulumi.String("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
					WebHookProperties: pulumi.StringMap{
						"key11": pulumi.String("value11"),
						"key12": pulumi.String("value12"),
					},
				},
			},
			AutoMitigate: pulumi.Bool(true),
			Criteria: insights.MetricAlertMultipleResourceMultipleMetricCriteria{
				AllOf: []interface{}{
					insights.MetricCriteria{
						CriterionType:   "StaticThresholdCriterion",
						Dimensions:      []insights.MetricDimension{},
						MetricName:      "Percentage CPU",
						MetricNamespace: "microsoft.compute/virtualmachines",
						Name:            "High_CPU_80",
						Operator:        "GreaterThan",
						Threshold:       80.5,
						TimeAggregation: "Average",
					},
				},
				OdataType: "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
			},
			Description:         pulumi.String("This is the description of the rule1"),
			Enabled:             pulumi.Bool(true),
			EvaluationFrequency: pulumi.String("PT1M"),
			Location:            pulumi.String("global"),
			ResourceGroupName:   pulumi.String("gigtest"),
			RuleName:            pulumi.String("MetricAlertOnMultipleResources"),
			Scopes: pulumi.StringArray{
				pulumi.String("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1"),
				pulumi.String("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2"),
			},
			Severity:             pulumi.Int(3),
			Tags:                 nil,
			TargetResourceRegion: pulumi.String("southcentralus"),
			TargetResourceType:   pulumi.String("Microsoft.Compute/virtualMachines"),
			WindowSize:           pulumi.String("PT15M"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
