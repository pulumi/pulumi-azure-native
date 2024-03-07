package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alertsmanagement.NewPrometheusRuleGroup(ctx, "prometheusRuleGroup", &alertsmanagement.PrometheusRuleGroupArgs{
			ClusterName:       pulumi.String("myClusterName"),
			Description:       pulumi.String("This is the description of the following rule group"),
			Enabled:           pulumi.Bool(true),
			Interval:          pulumi.String("PT10M"),
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("promResourceGroup"),
			RuleGroupName:     pulumi.String("myPrometheusRuleGroup"),
			Rules: []alertsmanagement.PrometheusRuleArgs{
				{
					Expression: pulumi.String("histogram_quantile(0.99, sum(rate(jobs_duration_seconds_bucket{service=\"billing-processing\"}[5m])) by (job_type))"),
					Labels: {
						"team": pulumi.String("prod"),
					},
					Record: pulumi.String("job_type:billing_jobs_duration_seconds:99p5m"),
				},
				{
					Actions: alertsmanagement.PrometheusRuleGroupActionArray{
						{
							ActionGroupId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/myrg/providers/microsoft.insights/actiongroups/myactiongroup"),
							ActionProperties: {
								"key11": pulumi.String("value11"),
								"key12": pulumi.String("value12"),
							},
						},
						{
							ActionGroupId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourcegroups/myrg/providers/microsoft.insights/actiongroups/myotheractiongroup"),
							ActionProperties: {
								"key21": pulumi.String("value21"),
								"key22": pulumi.String("value22"),
							},
						},
					},
					Alert: pulumi.String("Billing_Processing_Very_Slow"),
					Annotations: {
						"annotationName1": pulumi.String("annotationValue1"),
					},
					Enabled:    pulumi.Bool(true),
					Expression: pulumi.String("job_type:billing_jobs_duration_seconds:99p5m > 30"),
					For:        pulumi.String("PT5M"),
					Labels: {
						"team": pulumi.String("prod"),
					},
					ResolveConfiguration: {
						AutoResolved:  pulumi.Bool(true),
						TimeToResolve: pulumi.String("PT10M"),
					},
					Severity: pulumi.Int(2),
				},
			},
			Scopes: pulumi.StringArray{
				pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
