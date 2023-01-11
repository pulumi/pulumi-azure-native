


package v20160515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	var rv LookupScheduleResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduleArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupScheduleResult struct {
	CreatedDate          string                        `pulumi:"createdDate"`
	DailyRecurrence      *DayDetailsResponse           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetailsResponse          `pulumi:"hourlyRecurrence"`
	Id                   string                        `pulumi:"id"`
	Location             *string                       `pulumi:"location"`
	Name                 string                        `pulumi:"name"`
	NotificationSettings *NotificationSettingsResponse `pulumi:"notificationSettings"`
	ProvisioningState    *string                       `pulumi:"provisioningState"`
	Status               *string                       `pulumi:"status"`
	Tags                 map[string]string             `pulumi:"tags"`
	TargetResourceId     *string                       `pulumi:"targetResourceId"`
	TaskType             *string                       `pulumi:"taskType"`
	TimeZoneId           *string                       `pulumi:"timeZoneId"`
	Type                 string                        `pulumi:"type"`
	UniqueIdentifier     *string                       `pulumi:"uniqueIdentifier"`
	WeeklyRecurrence     *WeekDetailsResponse          `pulumi:"weeklyRecurrence"`
}

func LookupScheduleOutput(ctx *pulumi.Context, args LookupScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduleResult, error) {
			args := v.(LookupScheduleArgs)
			r, err := LookupSchedule(ctx, &args, opts...)
			var s LookupScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduleResultOutput)
}

type LookupScheduleOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleArgs)(nil)).Elem()
}


type LookupScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleResult)(nil)).Elem()
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutput() LookupScheduleResultOutput {
	return o
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutputWithContext(ctx context.Context) LookupScheduleResultOutput {
	return o
}

func (o LookupScheduleResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *DayDetailsResponse { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o LookupScheduleResultOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *HourDetailsResponse { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o LookupScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *NotificationSettingsResponse { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o LookupScheduleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupScheduleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupScheduleResultOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *WeekDetailsResponse { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduleResultOutput{})
}
