


package v20160515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupGlobalSchedule(ctx *pulumi.Context, args *LookupGlobalScheduleArgs, opts ...pulumi.InvokeOption) (*LookupGlobalScheduleResult, error) {
	var rv LookupGlobalScheduleResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getGlobalSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalScheduleArgs struct {
	Expand            *string `pulumi:"expand"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupGlobalScheduleResult struct {
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

func LookupGlobalScheduleOutput(ctx *pulumi.Context, args LookupGlobalScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalScheduleResult, error) {
			args := v.(LookupGlobalScheduleArgs)
			r, err := LookupGlobalSchedule(ctx, &args, opts...)
			var s LookupGlobalScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalScheduleResultOutput)
}

type LookupGlobalScheduleOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupGlobalScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalScheduleArgs)(nil)).Elem()
}


type LookupGlobalScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalScheduleResult)(nil)).Elem()
}

func (o LookupGlobalScheduleResultOutput) ToLookupGlobalScheduleResultOutput() LookupGlobalScheduleResultOutput {
	return o
}

func (o LookupGlobalScheduleResultOutput) ToLookupGlobalScheduleResultOutputWithContext(ctx context.Context) LookupGlobalScheduleResultOutput {
	return o
}

func (o LookupGlobalScheduleResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupGlobalScheduleResultOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *DayDetailsResponse { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o LookupGlobalScheduleResultOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *HourDetailsResponse { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o LookupGlobalScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGlobalScheduleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGlobalScheduleResultOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *NotificationSettingsResponse { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o LookupGlobalScheduleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalScheduleResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalScheduleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGlobalScheduleResultOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalScheduleResultOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *string { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalScheduleResultOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *string { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupGlobalScheduleResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupGlobalScheduleResultOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupGlobalScheduleResult) *WeekDetailsResponse { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalScheduleResultOutput{})
}
