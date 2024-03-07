package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewAutoscaleSetting(ctx, "autoscaleSetting", &insights.AutoscaleSettingArgs{
			AutoscaleSettingName: pulumi.String("MySetting"),
			Enabled:              pulumi.Bool(true),
			Location:             pulumi.String("West US"),
			Notifications: insights.AutoscaleNotificationArray{
				&insights.AutoscaleNotificationArgs{
					Email: &insights.EmailNotificationArgs{
						CustomEmails: pulumi.StringArray{
							pulumi.String("gu@ms.com"),
							pulumi.String("ge@ns.net"),
						},
						SendToSubscriptionAdministrator:    pulumi.Bool(true),
						SendToSubscriptionCoAdministrators: pulumi.Bool(true),
					},
					Operation: insights.OperationTypeScale,
					Webhooks: insights.WebhookNotificationArray{
						&insights.WebhookNotificationArgs{
							Properties: nil,
							ServiceUri: pulumi.String("http://myservice.com"),
						},
					},
				},
			},
			PredictiveAutoscalePolicy: &insights.PredictiveAutoscalePolicyArgs{
				ScaleMode: insights.PredictiveAutoscalePolicyScaleModeEnabled,
			},
			Profiles: insights.AutoscaleProfileArray{
				&insights.AutoscaleProfileArgs{
					Capacity: &insights.ScaleCapacityArgs{
						Default: pulumi.String("1"),
						Maximum: pulumi.String("10"),
						Minimum: pulumi.String("1"),
					},
					FixedDate: &insights.TimeWindowArgs{
						End:      pulumi.String("2015-03-05T14:30:00Z"),
						Start:    pulumi.String("2015-03-05T14:00:00Z"),
						TimeZone: pulumi.String("UTC"),
					},
					Name: pulumi.String("adios"),
					Rules: insights.ScaleRuleArray{
						&insights.ScaleRuleArgs{
							MetricTrigger: &insights.MetricTriggerArgs{
								DividePerInstance: pulumi.Bool(false),
								MetricName:        pulumi.String("Percentage CPU"),
								MetricResourceUri: pulumi.String("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
								Operator:          insights.ComparisonOperationTypeGreaterThan,
								Statistic:         insights.MetricStatisticTypeAverage,
								Threshold:         pulumi.Float64(10),
								TimeAggregation:   insights.TimeAggregationTypeAverage,
								TimeGrain:         pulumi.String("PT1M"),
								TimeWindow:        pulumi.String("PT5M"),
							},
							ScaleAction: &insights.ScaleActionArgs{
								Cooldown:  pulumi.String("PT5M"),
								Direction: insights.ScaleDirectionIncrease,
								Type:      insights.ScaleTypeChangeCount,
								Value:     pulumi.String("1"),
							},
						},
						&insights.ScaleRuleArgs{
							MetricTrigger: &insights.MetricTriggerArgs{
								DividePerInstance: pulumi.Bool(false),
								MetricName:        pulumi.String("Percentage CPU"),
								MetricResourceUri: pulumi.String("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
								Operator:          insights.ComparisonOperationTypeGreaterThan,
								Statistic:         insights.MetricStatisticTypeAverage,
								Threshold:         pulumi.Float64(15),
								TimeAggregation:   insights.TimeAggregationTypeAverage,
								TimeGrain:         pulumi.String("PT2M"),
								TimeWindow:        pulumi.String("PT5M"),
							},
							ScaleAction: &insights.ScaleActionArgs{
								Cooldown:  pulumi.String("PT6M"),
								Direction: insights.ScaleDirectionDecrease,
								Type:      insights.ScaleTypeChangeCount,
								Value:     pulumi.String("2"),
							},
						},
					},
				},
				&insights.AutoscaleProfileArgs{
					Capacity: &insights.ScaleCapacityArgs{
						Default: pulumi.String("1"),
						Maximum: pulumi.String("10"),
						Minimum: pulumi.String("1"),
					},
					Name: pulumi.String("saludos"),
					Recurrence: &insights.RecurrenceArgs{
						Frequency: insights.RecurrenceFrequencyWeek,
						Schedule: &insights.RecurrentScheduleArgs{
							Days: pulumi.StringArray{
								pulumi.String("1"),
							},
							Hours: pulumi.IntArray{
								pulumi.Int(5),
							},
							Minutes: pulumi.IntArray{
								pulumi.Int(15),
							},
							TimeZone: pulumi.String("UTC"),
						},
					},
					Rules: insights.ScaleRuleArray{
						&insights.ScaleRuleArgs{
							MetricTrigger: &insights.MetricTriggerArgs{
								DividePerInstance: pulumi.Bool(false),
								MetricName:        pulumi.String("Percentage CPU"),
								MetricResourceUri: pulumi.String("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
								Operator:          insights.ComparisonOperationTypeGreaterThan,
								Statistic:         insights.MetricStatisticTypeAverage,
								Threshold:         pulumi.Float64(10),
								TimeAggregation:   insights.TimeAggregationTypeAverage,
								TimeGrain:         pulumi.String("PT1M"),
								TimeWindow:        pulumi.String("PT5M"),
							},
							ScaleAction: &insights.ScaleActionArgs{
								Cooldown:  pulumi.String("PT5M"),
								Direction: insights.ScaleDirectionIncrease,
								Type:      insights.ScaleTypeChangeCount,
								Value:     pulumi.String("1"),
							},
						},
						&insights.ScaleRuleArgs{
							MetricTrigger: &insights.MetricTriggerArgs{
								DividePerInstance: pulumi.Bool(false),
								MetricName:        pulumi.String("Percentage CPU"),
								MetricResourceUri: pulumi.String("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
								Operator:          insights.ComparisonOperationTypeGreaterThan,
								Statistic:         insights.MetricStatisticTypeAverage,
								Threshold:         pulumi.Float64(15),
								TimeAggregation:   insights.TimeAggregationTypeAverage,
								TimeGrain:         pulumi.String("PT2M"),
								TimeWindow:        pulumi.String("PT5M"),
							},
							ScaleAction: &insights.ScaleActionArgs{
								Cooldown:  pulumi.String("PT6M"),
								Direction: insights.ScaleDirectionDecrease,
								Type:      insights.ScaleTypeChangeCount,
								Value:     pulumi.String("2"),
							},
						},
					},
				},
			},
			ResourceGroupName: pulumi.String("TestingMetricsScaleSet"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
				"key2": pulumi.String("value2"),
			},
			TargetResourceUri: pulumi.String("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
