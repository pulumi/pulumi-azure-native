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
					ActionGroupId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
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
						CriterionType: "StaticThresholdCriterion",
						Dimensions: []insights.MetricDimension{
							{
								Name:     "ActivityName",
								Operator: "Include",
								Values: []string{
									"*",
								},
							},
							{
								Name:     "StatusCode",
								Operator: "Include",
								Values: []string{
									"200",
								},
							},
						},
						MetricName:      "Availability",
						MetricNamespace: "Microsoft.KeyVault/vaults",
						Name:            "Metric1",
						Operator:        "GreaterThan",
						Threshold:       55,
						TimeAggregation: "Average",
					},
				},
				OdataType: "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
			},
			Description:         pulumi.String("This is the description of the rule1"),
			Enabled:             pulumi.Bool(true),
			EvaluationFrequency: pulumi.String("PT1H"),
			Location:            pulumi.String("global"),
			ResourceGroupName:   pulumi.String("gigtest"),
			RuleName:            pulumi.String("MetricAlertOnMultipleDimensions"),
			Scopes: pulumi.StringArray{
				pulumi.String("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.KeyVault/vaults/keyVaultResource"),
			},
			Severity:   pulumi.Int(3),
			Tags:       nil,
			WindowSize: pulumi.String("P1D"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
