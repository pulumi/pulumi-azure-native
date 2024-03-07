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
Criteria: insights.ScheduledQueryRuleCriteriaResponse{
AllOf: insights.ConditionArray{
interface{}{
Dimensions: insights.DimensionArray{
&insights.DimensionArgs{
Name: pulumi.String("ComputerIp"),
Operator: pulumi.String("Exclude"),
Values: pulumi.StringArray{
pulumi.String("192.168.1.1"),
},
},
&insights.DimensionArgs{
Name: pulumi.String("OSType"),
Operator: pulumi.String("Include"),
Values: pulumi.StringArray{
pulumi.String("*"),
},
},
},
FailingPeriods: &insights.ConditionFailingPeriodsArgs{
MinFailingPeriodsToAlert: pulumi.Float64(1),
NumberOfEvaluationPeriods: pulumi.Float64(1),
},
MetricMeasureColumn: pulumi.String("% Processor Time"),
Operator: pulumi.String("GreaterThan"),
Query: pulumi.String("Perf | where ObjectName == \"Processor\""),
ResourceIdColumn: pulumi.String("resourceId"),
Threshold: pulumi.Float64(70),
TimeAggregation: pulumi.String("Average"),
},
},
},
Description: pulumi.String("Performance rule"),
Enabled: pulumi.Bool(true),
EvaluationFrequency: pulumi.String("PT5M"),
Location: pulumi.String("eastus"),
MuteActionsDuration: pulumi.String("PT30M"),
ResourceGroupName: pulumi.String("QueryResourceGroupName"),
RuleName: pulumi.String("perf"),
RuleResolveConfiguration: &insights.RuleResolveConfigurationArgs{
AutoResolved: pulumi.Bool(true),
TimeToResolve: pulumi.String("PT10M"),
},
Scopes: pulumi.StringArray{
pulumi.String("/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147"),
},
Severity: pulumi.Float64(4),
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
