


package v20180915

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineSchedule(ctx *pulumi.Context, args *LookupVirtualMachineScheduleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScheduleResult, error) {
	var rv LookupVirtualMachineScheduleResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getVirtualMachineSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineScheduleArgs struct {
	Expand             *string `pulumi:"expand"`
	LabName            string  `pulumi:"labName"`
	Name               string  `pulumi:"name"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	VirtualMachineName string  `pulumi:"virtualMachineName"`
}


type LookupVirtualMachineScheduleResult struct {
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


func (val *LookupVirtualMachineScheduleResult) Defaults() *LookupVirtualMachineScheduleResult {
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
