


package v20180915

import (
	"context"
	"reflect"

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

func LookupServiceFabricScheduleOutput(ctx *pulumi.Context, args LookupServiceFabricScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupServiceFabricScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceFabricScheduleResult, error) {
			args := v.(LookupServiceFabricScheduleArgs)
			r, err := LookupServiceFabricSchedule(ctx, &args, opts...)
			var s LookupServiceFabricScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceFabricScheduleResultOutput)
}

type LookupServiceFabricScheduleOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ServiceFabricName pulumi.StringInput    `pulumi:"serviceFabricName"`
	UserName          pulumi.StringInput    `pulumi:"userName"`
}

func (LookupServiceFabricScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceFabricScheduleArgs)(nil)).Elem()
}


type LookupServiceFabricScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupServiceFabricScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceFabricScheduleResult)(nil)).Elem()
}

func (o LookupServiceFabricScheduleResultOutput) ToLookupServiceFabricScheduleResultOutput() LookupServiceFabricScheduleResultOutput {
	return o
}

func (o LookupServiceFabricScheduleResultOutput) ToLookupServiceFabricScheduleResultOutputWithContext(ctx context.Context) LookupServiceFabricScheduleResultOutput {
	return o
}

func (o LookupServiceFabricScheduleResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupServiceFabricScheduleResultOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) *DayDetailsResponse { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o LookupServiceFabricScheduleResultOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) *HourDetailsResponse { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o LookupServiceFabricScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceFabricScheduleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupServiceFabricScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceFabricScheduleResultOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) *NotificationSettingsResponse { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o LookupServiceFabricScheduleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServiceFabricScheduleResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupServiceFabricScheduleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceFabricScheduleResultOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupServiceFabricScheduleResultOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) *string { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o LookupServiceFabricScheduleResultOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) *string { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o LookupServiceFabricScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupServiceFabricScheduleResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o LookupServiceFabricScheduleResultOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceFabricScheduleResult) *WeekDetailsResponse { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceFabricScheduleResultOutput{})
}
