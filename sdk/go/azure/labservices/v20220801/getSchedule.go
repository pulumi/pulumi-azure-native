


package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	var rv LookupScheduleResult
	err := ctx.Invoke("azure-native:labservices/v20220801:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduleArgs struct {
	LabName           string `pulumi:"labName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ScheduleName      string `pulumi:"scheduleName"`
}


type LookupScheduleResult struct {
	Id                string                     `pulumi:"id"`
	Name              string                     `pulumi:"name"`
	Notes             *string                    `pulumi:"notes"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	RecurrencePattern *RecurrencePatternResponse `pulumi:"recurrencePattern"`
	StartAt           *string                    `pulumi:"startAt"`
	StopAt            string                     `pulumi:"stopAt"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	TimeZoneId        string                     `pulumi:"timeZoneId"`
	Type              string                     `pulumi:"type"`
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
	LabName           pulumi.StringInput `pulumi:"labName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ScheduleName      pulumi.StringInput `pulumi:"scheduleName"`
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

func (o LookupScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) RecurrencePattern() RecurrencePatternResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *RecurrencePatternResponse { return v.RecurrencePattern }).(RecurrencePatternResponsePtrOutput)
}

func (o LookupScheduleResultOutput) StartAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.StartAt }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) StopAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.StopAt }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupScheduleResultOutput) TimeZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.TimeZoneId }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduleResultOutput{})
}
