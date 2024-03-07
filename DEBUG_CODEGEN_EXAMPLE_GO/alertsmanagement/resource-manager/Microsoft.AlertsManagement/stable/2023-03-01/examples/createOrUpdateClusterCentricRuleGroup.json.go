package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := alertsmanagement.NewPrometheusRuleGroup(ctx, "prometheusRuleGroup", &alertsmanagement.PrometheusRuleGroupArgs{
			ClusterName:       pulumi.String("myClusterName"),
			Description:       pulumi.String("This is a rule group with culster centric configuration"),
			Interval:          pulumi.String("PT10M"),
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("promResourceGroup"),
			RuleGroupName:     pulumi.String("myPrometheusRuleGroup"),
			Rules: alertsmanagement.PrometheusRuleArray{
				&alertsmanagement.PrometheusRuleArgs{
					Actions: alertsmanagement.PrometheusRuleGroupActionArray{},
					Alert:   pulumi.String("Billing_Processing_Very_Slow"),
					Annotations: pulumi.StringMap{
						"annotationName1": pulumi.String("annotationValue1"),
					},
					Enabled:    pulumi.Bool(true),
					Expression: pulumi.String("job_type:billing_jobs_duration_seconds:99p5m > 30"),
					For:        pulumi.String("PT5M"),
					Labels: pulumi.StringMap{
						"team": pulumi.String("prod"),
					},
					ResolveConfiguration: &alertsmanagement.PrometheusRuleResolveConfigurationArgs{
						AutoResolved:  pulumi.Bool(true),
						TimeToResolve: pulumi.String("PT10M"),
					},
					Severity: pulumi.Int(2),
				},
			},
			Scopes: pulumi.StringArray{
				pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace"),
				pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myClusterName"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
