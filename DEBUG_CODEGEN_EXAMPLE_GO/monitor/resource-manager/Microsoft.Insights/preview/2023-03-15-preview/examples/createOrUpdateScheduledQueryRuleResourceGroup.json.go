package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewScheduledQueryRule(ctx, "scheduledQueryRule", &insights.ScheduledQueryRuleArgs{
			Actions: &insights.ActionsArgs{
				ActionGroups: pulumi.StringArray{
					pulumi.String("/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup"),
				},
				CustomProperties: pulumi.StringMap{
					"key11": pulumi.String("value11"),
					"key12": pulumi.String("value12"),
				},
			},
			CheckWorkspaceAlertsStorageConfigured: pulumi.Bool(true),
			Criteria: &insights.ScheduledQueryRuleCriteriaArgs{
				AllOf: insights.ConditionArray{
					&insights.ConditionArgs{
						Dimensions: insights.DimensionArray{},
						FailingPeriods: &insights.ConditionFailingPeriodsArgs{
							MinFailingPeriodsToAlert:  pulumi.Float64(1),
							NumberOfEvaluationPeriods: pulumi.Float64(1),
						},
						Operator:        pulumi.String("GreaterThan"),
						Query:           pulumi.String("Heartbeat"),
						Threshold:       pulumi.Float64(360),
						TimeAggregation: pulumi.String("Count"),
					},
				},
			},
			Description:         pulumi.String("Health check rule"),
			Enabled:             pulumi.Bool(true),
			EvaluationFrequency: pulumi.String("PT5M"),
			Location:            pulumi.String("eastus"),
			MuteActionsDuration: pulumi.String("PT30M"),
			ResourceGroupName:   pulumi.String("QueryResourceGroupName"),
			RuleName:            pulumi.String("heartbeat"),
			RuleResolveConfiguration: &insights.RuleResolveConfigurationArgs{
				AutoResolved:  pulumi.Bool(true),
				TimeToResolve: pulumi.String("PT10M"),
			},
			Scopes: pulumi.StringArray{
				pulumi.String("/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147/resourceGroups/scopeResourceGroup1"),
			},
			Severity:            pulumi.Float64(4),
			SkipQueryValidation: pulumi.Bool(true),
			TargetResourceTypes: pulumi.StringArray{
				pulumi.String("Microsoft.Compute/virtualMachines"),
			},
			WindowSize: pulumi.String("PT10M"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
