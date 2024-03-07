package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := machinelearningservices.NewFeaturesetVersion(ctx, "featuresetVersion", &machinelearningservices.FeaturesetVersionArgs{
FeaturesetVersionProperties: &machinelearningservices.FeaturesetVersionTypeArgs{
Description: pulumi.String("string"),
Entities: pulumi.StringArray{
pulumi.String("string"),
},
IsAnonymous: pulumi.Bool(false),
IsArchived: pulumi.Bool(false),
MaterializationSettings: &machinelearningservices.MaterializationSettingsArgs{
Notification: &machinelearningservices.NotificationSettingArgs{
EmailOn: pulumi.StringArray{
pulumi.String("JobFailed"),
},
Emails: pulumi.StringArray{
pulumi.String("string"),
},
},
Resource: &machinelearningservices.MaterializationComputeResourceArgs{
InstanceType: pulumi.String("string"),
},
Schedule: interface{}{
EndTime: pulumi.String("string"),
Frequency: pulumi.String("Day"),
Interval: pulumi.Int(1),
Schedule: &machinelearningservices.RecurrenceScheduleArgs{
Hours: pulumi.IntArray{
pulumi.Int(1),
},
Minutes: pulumi.IntArray{
pulumi.Int(1),
},
MonthDays: pulumi.IntArray{
pulumi.Int(1),
},
WeekDays: pulumi.StringArray{
pulumi.String("Monday"),
},
},
StartTime: pulumi.String("string"),
TimeZone: pulumi.String("string"),
TriggerType: pulumi.String("Recurrence"),
},
SparkConfiguration: pulumi.StringMap{
"string": pulumi.String("string"),
},
StoreType: pulumi.String("Online"),
},
Properties: pulumi.StringMap{
"string": pulumi.String("string"),
},
Specification: &machinelearningservices.FeaturesetSpecificationArgs{
Path: pulumi.String("string"),
},
Stage: pulumi.String("string"),
Tags: pulumi.StringMap{
"string": pulumi.String("string"),
},
},
Name: pulumi.String("string"),
ResourceGroupName: pulumi.String("test-rg"),
Version: pulumi.String("string"),
WorkspaceName: pulumi.String("my-aml-workspace"),
})
if err != nil {
return err
}
return nil
})
}
