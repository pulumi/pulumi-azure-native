


package v20151031

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	var rv LookupScheduleResult
	err := ctx.Invoke("azure-native:automation/v20151031:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupScheduleArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	ScheduleName          string `pulumi:"scheduleName"`
}


type LookupScheduleResult struct {
	AdvancedSchedule        *AdvancedScheduleResponse `pulumi:"advancedSchedule"`
	CreationTime            *string                   `pulumi:"creationTime"`
	Description             *string                   `pulumi:"description"`
	ExpiryTime              *string                   `pulumi:"expiryTime"`
	ExpiryTimeOffsetMinutes *float64                  `pulumi:"expiryTimeOffsetMinutes"`
	Frequency               *string                   `pulumi:"frequency"`
	Id                      string                    `pulumi:"id"`
	Interval                interface{}               `pulumi:"interval"`
	IsEnabled               *bool                     `pulumi:"isEnabled"`
	LastModifiedTime        *string                   `pulumi:"lastModifiedTime"`
	Name                    string                    `pulumi:"name"`
	NextRun                 *string                   `pulumi:"nextRun"`
	NextRunOffsetMinutes    *float64                  `pulumi:"nextRunOffsetMinutes"`
	StartTime               *string                   `pulumi:"startTime"`
	StartTimeOffsetMinutes  float64                   `pulumi:"startTimeOffsetMinutes"`
	TimeZone                *string                   `pulumi:"timeZone"`
	Type                    string                    `pulumi:"type"`
}


func (val *LookupScheduleResult) Defaults() *LookupScheduleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsEnabled) {
		isEnabled_ := false
		tmp.IsEnabled = &isEnabled_
	}
	return &tmp
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
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	ScheduleName          pulumi.StringInput `pulumi:"scheduleName"`
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

func (o LookupScheduleResultOutput) AdvancedSchedule() AdvancedScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *AdvancedScheduleResponse { return v.AdvancedSchedule }).(AdvancedScheduleResponsePtrOutput)
}

func (o LookupScheduleResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) ExpiryTimeOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *float64 { return v.ExpiryTimeOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o LookupScheduleResultOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) Interval() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupScheduleResult) interface{} { return v.Interval }).(pulumi.AnyOutput)
}

func (o LookupScheduleResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupScheduleResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduleResultOutput) NextRun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.NextRun }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) NextRunOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *float64 { return v.NextRunOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o LookupScheduleResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) StartTimeOffsetMinutes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupScheduleResult) float64 { return v.StartTimeOffsetMinutes }).(pulumi.Float64Output)
}

func (o LookupScheduleResultOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o LookupScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduleResultOutput{})
}
