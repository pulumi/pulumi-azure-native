


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupScheduleResource(ctx *pulumi.Context, args *LookupScheduleResourceArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResourceResult, error) {
	var rv LookupScheduleResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getScheduleResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduleResourceArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupScheduleResourceResult struct {
	DailyRecurrence   *DayDetailsResponse  `pulumi:"dailyRecurrence"`
	HourlyRecurrence  *HourDetailsResponse `pulumi:"hourlyRecurrence"`
	Id                *string              `pulumi:"id"`
	Location          *string              `pulumi:"location"`
	Name              *string              `pulumi:"name"`
	ProvisioningState *string              `pulumi:"provisioningState"`
	Status            *string              `pulumi:"status"`
	Tags              map[string]string    `pulumi:"tags"`
	TaskType          *string              `pulumi:"taskType"`
	TimeZoneId        *string              `pulumi:"timeZoneId"`
	Type              *string              `pulumi:"type"`
	WeeklyRecurrence  *WeekDetailsResponse `pulumi:"weeklyRecurrence"`
}

func LookupScheduleResourceOutput(ctx *pulumi.Context, args LookupScheduleResourceOutputArgs, opts ...pulumi.InvokeOption) LookupScheduleResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduleResourceResult, error) {
			args := v.(LookupScheduleResourceArgs)
			r, err := LookupScheduleResource(ctx, &args, opts...)
			var s LookupScheduleResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduleResourceResultOutput)
}

type LookupScheduleResourceOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupScheduleResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleResourceArgs)(nil)).Elem()
}


type LookupScheduleResourceResultOutput struct{ *pulumi.OutputState }

func (LookupScheduleResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleResourceResult)(nil)).Elem()
}

func (o LookupScheduleResourceResultOutput) ToLookupScheduleResourceResultOutput() LookupScheduleResourceResultOutput {
	return o
}

func (o LookupScheduleResourceResultOutput) ToLookupScheduleResourceResultOutputWithContext(ctx context.Context) LookupScheduleResourceResultOutput {
	return o
}

func (o LookupScheduleResourceResultOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *DayDetailsResponse { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o LookupScheduleResourceResultOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *HourDetailsResponse { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o LookupScheduleResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResourceResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupScheduleResourceResultOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *string { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResourceResultOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *string { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResourceResultOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResourceResult) *WeekDetailsResponse { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduleResourceResultOutput{})
}
