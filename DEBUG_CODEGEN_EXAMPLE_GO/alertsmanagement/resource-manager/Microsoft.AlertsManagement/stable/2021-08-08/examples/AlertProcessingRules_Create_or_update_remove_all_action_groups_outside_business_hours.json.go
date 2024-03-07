package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/alertsmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := alertsmanagement.NewAlertProcessingRuleByName(ctx, "alertProcessingRuleByName", &alertsmanagement.AlertProcessingRuleByNameArgs{
AlertProcessingRuleName: pulumi.String("RemoveActionGroupsOutsideBusinessHours"),
Location: pulumi.String("Global"),
Properties: alertsmanagement.AlertProcessingRulePropertiesResponse{
Actions: pulumi.Array{
alertsmanagement.RemoveAllActionGroups{
ActionType: "RemoveAllActionGroups",
},
},
Description: pulumi.String("Remove all ActionGroups outside business hours"),
Enabled: pulumi.Bool(true),
Schedule: interface{}{
Recurrences: pulumi.Array{
alertsmanagement.DailyRecurrence{
EndTime: "09:00:00",
RecurrenceType: "Daily",
StartTime: "17:00:00",
},
alertsmanagement.WeeklyRecurrence{
DaysOfWeek: []alertsmanagement.DaysOfWeek{
"Saturday",
"Sunday",
},
RecurrenceType: "Weekly",
},
},
TimeZone: pulumi.String("Eastern Standard Time"),
},
Scopes: pulumi.StringArray{
pulumi.String("/subscriptions/subId1"),
},
},
ResourceGroupName: pulumi.String("alertscorrelationrg"),
Tags: nil,
})
if err != nil {
return err
}
return nil
})
}
