


package v20180915

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceFabricSchedule(ctx *pulumi.Context, args *LookupServiceFabricScheduleArgs, opts ...pulumi.InvokeOption) (*LookupServiceFabricScheduleResult, error) {
	var rv LookupServiceFabricScheduleResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getServiceFabricSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServiceFabricScheduleArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceFabricName string  `pulumi:"serviceFabricName"`
	UserName          string  `pulumi:"userName"`
}


type LookupServiceFabricScheduleResult struct {
	CreatedDate          string                        `pulumi:"createdDate"`
	DailyRecurrence      *DayDetailsResponse           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetailsResponse          `pulumi:"hourlyRecurrence"`
	Id                   string                        `pulumi:"id"`
	Location             *string                       `pulumi:"location"`
	Name                 string                        `pulumi:"name"`
	NotificationSettings *NotificationSettingsResponse `pulumi:"notificationSettings"`
	ProvisioningState    string                        `pulumi:"provisioningState"`
	Status               *string                       `pulumi:"status"`
	Tags                 map[string]string             `pulumi:"tags"`
	TargetResourceId     *string                       `pulumi:"targetResourceId"`
	TaskType             *string                       `pulumi:"taskType"`
	TimeZoneId           *string                       `pulumi:"timeZoneId"`
	Type                 string                        `pulumi:"type"`
	UniqueIdentifier     string                        `pulumi:"uniqueIdentifier"`
	WeeklyRecurrence     *WeekDetailsResponse          `pulumi:"weeklyRecurrence"`
}


func (val *LookupServiceFabricScheduleResult) Defaults() *LookupServiceFabricScheduleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NotificationSettings = tmp.NotificationSettings.Defaults()

	if isZero(tmp.Status) {
		status_ := "Disabled"
		tmp.Status = &status_
	}
	return &tmp
}
