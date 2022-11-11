


package v20180915

import (
	"context"
	"reflect"

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

func LookupVirtualMachineScheduleOutput(ctx *pulumi.Context, args LookupVirtualMachineScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineScheduleResult, error) {
			args := v.(LookupVirtualMachineScheduleArgs)
			r, err := LookupVirtualMachineSchedule(ctx, &args, opts...)
			var s LookupVirtualMachineScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineScheduleResultOutput)
}

type LookupVirtualMachineScheduleOutputArgs struct {
	Expand             pulumi.StringPtrInput `pulumi:"expand"`
	LabName            pulumi.StringInput    `pulumi:"labName"`
	Name               pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
	VirtualMachineName pulumi.StringInput    `pulumi:"virtualMachineName"`
}

func (LookupVirtualMachineScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScheduleArgs)(nil)).Elem()
}


type LookupVirtualMachineScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScheduleResult)(nil)).Elem()
}

func (o LookupVirtualMachineScheduleResultOutput) ToLookupVirtualMachineScheduleResultOutput() LookupVirtualMachineScheduleResultOutput {
	return o
}

func (o LookupVirtualMachineScheduleResultOutput) ToLookupVirtualMachineScheduleResultOutputWithContext(ctx context.Context) LookupVirtualMachineScheduleResultOutput {
	return o
}

func (o LookupVirtualMachineScheduleResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) *DayDetailsResponse { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) *HourDetailsResponse { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) *NotificationSettingsResponse {
		return v.NotificationSettings
	}).(NotificationSettingsResponsePtrOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) *string { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) *string { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScheduleResultOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScheduleResult) *WeekDetailsResponse { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineScheduleResultOutput{})
}
