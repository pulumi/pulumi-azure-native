


package v20220909

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScalingPlanPooledSchedule(ctx *pulumi.Context, args *LookupScalingPlanPooledScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScalingPlanPooledScheduleResult, error) {
	var rv LookupScalingPlanPooledScheduleResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20220909:getScalingPlanPooledSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScalingPlanPooledScheduleArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ScalingPlanName         string `pulumi:"scalingPlanName"`
	ScalingPlanScheduleName string `pulumi:"scalingPlanScheduleName"`
}


type LookupScalingPlanPooledScheduleResult struct {
	DaysOfWeek                     []string           `pulumi:"daysOfWeek"`
	Id                             string             `pulumi:"id"`
	Name                           string             `pulumi:"name"`
	OffPeakLoadBalancingAlgorithm  *string            `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               *TimeResponse      `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     *string            `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  *TimeResponse      `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   *int               `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       *bool              `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm *string            `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        *int               `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    *string            `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              *TimeResponse      `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          *string            `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        *int               `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     *int               `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   *string            `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          *int               `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                *TimeResponse      `pulumi:"rampUpStartTime"`
	SystemData                     SystemDataResponse `pulumi:"systemData"`
	Type                           string             `pulumi:"type"`
}

func LookupScalingPlanPooledScheduleOutput(ctx *pulumi.Context, args LookupScalingPlanPooledScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupScalingPlanPooledScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScalingPlanPooledScheduleResult, error) {
			args := v.(LookupScalingPlanPooledScheduleArgs)
			r, err := LookupScalingPlanPooledSchedule(ctx, &args, opts...)
			var s LookupScalingPlanPooledScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScalingPlanPooledScheduleResultOutput)
}

type LookupScalingPlanPooledScheduleOutputArgs struct {
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	ScalingPlanName         pulumi.StringInput `pulumi:"scalingPlanName"`
	ScalingPlanScheduleName pulumi.StringInput `pulumi:"scalingPlanScheduleName"`
}

func (LookupScalingPlanPooledScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanPooledScheduleArgs)(nil)).Elem()
}


type LookupScalingPlanPooledScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupScalingPlanPooledScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanPooledScheduleResult)(nil)).Elem()
}

func (o LookupScalingPlanPooledScheduleResultOutput) ToLookupScalingPlanPooledScheduleResultOutput() LookupScalingPlanPooledScheduleResultOutput {
	return o
}

func (o LookupScalingPlanPooledScheduleResultOutput) ToLookupScalingPlanPooledScheduleResultOutputWithContext(ctx context.Context) LookupScalingPlanPooledScheduleResultOutput {
	return o
}

func (o LookupScalingPlanPooledScheduleResultOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) OffPeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.OffPeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) OffPeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *TimeResponse { return v.OffPeakStartTime }).(TimeResponsePtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) PeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.PeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) PeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *TimeResponse { return v.PeakStartTime }).(TimeResponsePtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampDownCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampDownCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampDownForceLogoffUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *bool { return v.RampDownForceLogoffUsers }).(pulumi.BoolPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampDownLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.RampDownLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampDownMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampDownMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampDownNotificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.RampDownNotificationMessage }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampDownStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *TimeResponse { return v.RampDownStartTime }).(TimeResponsePtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampDownStopHostsWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.RampDownStopHostsWhen }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampDownWaitTimeMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampDownWaitTimeMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampUpCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampUpCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampUpLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *string { return v.RampUpLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampUpMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *int { return v.RampUpMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) RampUpStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) *TimeResponse { return v.RampUpStartTime }).(TimeResponsePtrOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupScalingPlanPooledScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScalingPlanPooledScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScalingPlanPooledScheduleResultOutput{})
}
