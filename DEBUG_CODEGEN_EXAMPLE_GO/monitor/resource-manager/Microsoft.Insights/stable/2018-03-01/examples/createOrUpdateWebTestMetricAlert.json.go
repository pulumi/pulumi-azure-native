package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewMetricAlert(ctx, "metricAlert", &insights.MetricAlertArgs{
			Actions: insights.MetricAlertActionArray{},
			Criteria: insights.WebtestLocationAvailabilityCriteria{
				ComponentId:         "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example",
				FailedLocationCount: 2,
				OdataType:           "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria",
				WebTestId:           "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example",
			},
			Description:         pulumi.String("Automatically created alert rule for availability test \"component-example\" a"),
			Enabled:             pulumi.Bool(true),
			EvaluationFrequency: pulumi.String("PT1M"),
			Location:            pulumi.String("global"),
			ResourceGroupName:   pulumi.String("rg-example"),
			RuleName:            pulumi.String("webtest-name-example"),
			Scopes: pulumi.StringArray{
				pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example"),
				pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example"),
			},
			Severity: pulumi.Int(4),
			Tags: pulumi.StringMap{
				"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example": pulumi.String("Resource"),
				"hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example":      pulumi.String("Resource"),
			},
			WindowSize: pulumi.String("PT15M"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
