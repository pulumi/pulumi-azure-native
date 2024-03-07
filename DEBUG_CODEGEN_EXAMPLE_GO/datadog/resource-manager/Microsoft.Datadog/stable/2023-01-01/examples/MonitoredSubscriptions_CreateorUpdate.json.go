package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datadog/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := datadog.NewMonitoredSubscription(ctx, "monitoredSubscription", &datadog.MonitoredSubscriptionArgs{
ConfigurationName: pulumi.String("default"),
MonitorName: pulumi.String("myMonitor"),
Properties: datadog.SubscriptionListResponse{
MonitoredSubscriptionList: datadog.MonitoredSubscriptionTypeArray{
interface{}{
Status: pulumi.String("Active"),
SubscriptionId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000"),
TagRules: interface{}{
Automuting: pulumi.Bool(true),
LogRules: interface{}{
FilteringTags: datadog.FilteringTagArray{
&datadog.FilteringTagArgs{
Action: pulumi.String("Include"),
Name: pulumi.String("Environment"),
Value: pulumi.String("Prod"),
},
&datadog.FilteringTagArgs{
Action: pulumi.String("Exclude"),
Name: pulumi.String("Environment"),
Value: pulumi.String("Dev"),
},
},
SendAadLogs: pulumi.Bool(false),
SendResourceLogs: pulumi.Bool(true),
SendSubscriptionLogs: pulumi.Bool(true),
},
MetricRules: interface{}{
FilteringTags: datadog.FilteringTagArray{
},
},
},
},
interface{}{
Status: pulumi.String("Failed"),
SubscriptionId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000001"),
TagRules: interface{}{
Automuting: pulumi.Bool(true),
LogRules: interface{}{
FilteringTags: datadog.FilteringTagArray{
&datadog.FilteringTagArgs{
Action: pulumi.String("Include"),
Name: pulumi.String("Environment"),
Value: pulumi.String("Prod"),
},
&datadog.FilteringTagArgs{
Action: pulumi.String("Exclude"),
Name: pulumi.String("Environment"),
Value: pulumi.String("Dev"),
},
},
SendAadLogs: pulumi.Bool(false),
SendResourceLogs: pulumi.Bool(true),
SendSubscriptionLogs: pulumi.Bool(true),
},
MetricRules: interface{}{
FilteringTags: datadog.FilteringTagArray{
},
},
},
},
},
Operation: pulumi.String("AddBegin"),
},
ResourceGroupName: pulumi.String("myResourceGroup"),
})
if err != nil {
return err
}
return nil
})
}
