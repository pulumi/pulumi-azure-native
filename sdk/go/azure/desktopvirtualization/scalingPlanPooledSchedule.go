


package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScalingPlanPooledSchedule struct {
	pulumi.CustomResourceState

	DaysOfWeek                     pulumi.StringArrayOutput `pulumi:"daysOfWeek"`
	Name                           pulumi.StringOutput      `pulumi:"name"`
	OffPeakLoadBalancingAlgorithm  pulumi.StringPtrOutput   `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               TimeResponsePtrOutput    `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     pulumi.StringPtrOutput   `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  TimeResponsePtrOutput    `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   pulumi.IntPtrOutput      `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       pulumi.BoolPtrOutput     `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm pulumi.StringPtrOutput   `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        pulumi.IntPtrOutput      `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    pulumi.StringPtrOutput   `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              TimeResponsePtrOutput    `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          pulumi.StringPtrOutput   `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        pulumi.IntPtrOutput      `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     pulumi.IntPtrOutput      `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   pulumi.StringPtrOutput   `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          pulumi.IntPtrOutput      `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                TimeResponsePtrOutput    `pulumi:"rampUpStartTime"`
	SystemData                     SystemDataResponseOutput `pulumi:"systemData"`
	Type                           pulumi.StringOutput      `pulumi:"type"`
}


func NewScalingPlanPooledSchedule(ctx *pulumi.Context,
	name string, args *ScalingPlanPooledScheduleArgs, opts ...pulumi.ResourceOption) (*ScalingPlanPooledSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScalingPlanName == nil {
		return nil, errors.New("invalid value for required argument 'ScalingPlanName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:ScalingPlanPooledSchedule"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220909:ScalingPlanPooledSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource ScalingPlanPooledSchedule
	err := ctx.RegisterResource("azure-native:desktopvirtualization:ScalingPlanPooledSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScalingPlanPooledSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPlanPooledScheduleState, opts ...pulumi.ResourceOption) (*ScalingPlanPooledSchedule, error) {
	var resource ScalingPlanPooledSchedule
	err := ctx.ReadResource("azure-native:desktopvirtualization:ScalingPlanPooledSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scalingPlanPooledScheduleState struct {
}

type ScalingPlanPooledScheduleState struct {
}

func (ScalingPlanPooledScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanPooledScheduleState)(nil)).Elem()
}

type scalingPlanPooledScheduleArgs struct {
	DaysOfWeek                     []string `pulumi:"daysOfWeek"`
	OffPeakLoadBalancingAlgorithm  *string  `pulumi:"offPeakLoadBalancingAlgorithm"`
	OffPeakStartTime               *Time    `pulumi:"offPeakStartTime"`
	PeakLoadBalancingAlgorithm     *string  `pulumi:"peakLoadBalancingAlgorithm"`
	PeakStartTime                  *Time    `pulumi:"peakStartTime"`
	RampDownCapacityThresholdPct   *int     `pulumi:"rampDownCapacityThresholdPct"`
	RampDownForceLogoffUsers       *bool    `pulumi:"rampDownForceLogoffUsers"`
	RampDownLoadBalancingAlgorithm *string  `pulumi:"rampDownLoadBalancingAlgorithm"`
	RampDownMinimumHostsPct        *int     `pulumi:"rampDownMinimumHostsPct"`
	RampDownNotificationMessage    *string  `pulumi:"rampDownNotificationMessage"`
	RampDownStartTime              *Time    `pulumi:"rampDownStartTime"`
	RampDownStopHostsWhen          *string  `pulumi:"rampDownStopHostsWhen"`
	RampDownWaitTimeMinutes        *int     `pulumi:"rampDownWaitTimeMinutes"`
	RampUpCapacityThresholdPct     *int     `pulumi:"rampUpCapacityThresholdPct"`
	RampUpLoadBalancingAlgorithm   *string  `pulumi:"rampUpLoadBalancingAlgorithm"`
	RampUpMinimumHostsPct          *int     `pulumi:"rampUpMinimumHostsPct"`
	RampUpStartTime                *Time    `pulumi:"rampUpStartTime"`
	ResourceGroupName              string   `pulumi:"resourceGroupName"`
	ScalingPlanName                string   `pulumi:"scalingPlanName"`
	ScalingPlanScheduleName        *string  `pulumi:"scalingPlanScheduleName"`
}


type ScalingPlanPooledScheduleArgs struct {
	DaysOfWeek                     pulumi.StringArrayInput
	OffPeakLoadBalancingAlgorithm  pulumi.StringPtrInput
	OffPeakStartTime               TimePtrInput
	PeakLoadBalancingAlgorithm     pulumi.StringPtrInput
	PeakStartTime                  TimePtrInput
	RampDownCapacityThresholdPct   pulumi.IntPtrInput
	RampDownForceLogoffUsers       pulumi.BoolPtrInput
	RampDownLoadBalancingAlgorithm pulumi.StringPtrInput
	RampDownMinimumHostsPct        pulumi.IntPtrInput
	RampDownNotificationMessage    pulumi.StringPtrInput
	RampDownStartTime              TimePtrInput
	RampDownStopHostsWhen          pulumi.StringPtrInput
	RampDownWaitTimeMinutes        pulumi.IntPtrInput
	RampUpCapacityThresholdPct     pulumi.IntPtrInput
	RampUpLoadBalancingAlgorithm   pulumi.StringPtrInput
	RampUpMinimumHostsPct          pulumi.IntPtrInput
	RampUpStartTime                TimePtrInput
	ResourceGroupName              pulumi.StringInput
	ScalingPlanName                pulumi.StringInput
	ScalingPlanScheduleName        pulumi.StringPtrInput
}

func (ScalingPlanPooledScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanPooledScheduleArgs)(nil)).Elem()
}

type ScalingPlanPooledScheduleInput interface {
	pulumi.Input

	ToScalingPlanPooledScheduleOutput() ScalingPlanPooledScheduleOutput
	ToScalingPlanPooledScheduleOutputWithContext(ctx context.Context) ScalingPlanPooledScheduleOutput
}

func (*ScalingPlanPooledSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlanPooledSchedule)(nil)).Elem()
}

func (i *ScalingPlanPooledSchedule) ToScalingPlanPooledScheduleOutput() ScalingPlanPooledScheduleOutput {
	return i.ToScalingPlanPooledScheduleOutputWithContext(context.Background())
}

func (i *ScalingPlanPooledSchedule) ToScalingPlanPooledScheduleOutputWithContext(ctx context.Context) ScalingPlanPooledScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanPooledScheduleOutput)
}

type ScalingPlanPooledScheduleOutput struct{ *pulumi.OutputState }

func (ScalingPlanPooledScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlanPooledSchedule)(nil)).Elem()
}

func (o ScalingPlanPooledScheduleOutput) ToScalingPlanPooledScheduleOutput() ScalingPlanPooledScheduleOutput {
	return o
}

func (o ScalingPlanPooledScheduleOutput) ToScalingPlanPooledScheduleOutputWithContext(ctx context.Context) ScalingPlanPooledScheduleOutput {
	return o
}

func (o ScalingPlanPooledScheduleOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringArrayOutput { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o ScalingPlanPooledScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScalingPlanPooledScheduleOutput) OffPeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.OffPeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) OffPeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) TimeResponsePtrOutput { return v.OffPeakStartTime }).(TimeResponsePtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) PeakLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.PeakLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) PeakStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) TimeResponsePtrOutput { return v.PeakStartTime }).(TimeResponsePtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampDownCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampDownCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampDownForceLogoffUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.BoolPtrOutput { return v.RampDownForceLogoffUsers }).(pulumi.BoolPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampDownLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.RampDownLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampDownMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampDownMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampDownNotificationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.RampDownNotificationMessage }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampDownStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) TimeResponsePtrOutput { return v.RampDownStartTime }).(TimeResponsePtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampDownStopHostsWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.RampDownStopHostsWhen }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampDownWaitTimeMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampDownWaitTimeMinutes }).(pulumi.IntPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampUpCapacityThresholdPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampUpCapacityThresholdPct }).(pulumi.IntPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampUpLoadBalancingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringPtrOutput { return v.RampUpLoadBalancingAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampUpMinimumHostsPct() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.IntPtrOutput { return v.RampUpMinimumHostsPct }).(pulumi.IntPtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) RampUpStartTime() TimeResponsePtrOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) TimeResponsePtrOutput { return v.RampUpStartTime }).(TimeResponsePtrOutput)
}

func (o ScalingPlanPooledScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScalingPlanPooledScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingPlanPooledSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalingPlanPooledScheduleOutput{})
}
