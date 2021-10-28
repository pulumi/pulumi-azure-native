


package v20151031

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContentSourceType string

const (
	ContentSourceTypeEmbeddedContent = ContentSourceType("embeddedContent")
	ContentSourceTypeUri             = ContentSourceType("uri")
)

func (ContentSourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSourceType)(nil)).Elem()
}

func (e ContentSourceType) ToContentSourceTypeOutput() ContentSourceTypeOutput {
	return pulumi.ToOutput(e).(ContentSourceTypeOutput)
}

func (e ContentSourceType) ToContentSourceTypeOutputWithContext(ctx context.Context) ContentSourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContentSourceTypeOutput)
}

func (e ContentSourceType) ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput {
	return e.ToContentSourceTypePtrOutputWithContext(context.Background())
}

func (e ContentSourceType) ToContentSourceTypePtrOutputWithContext(ctx context.Context) ContentSourceTypePtrOutput {
	return ContentSourceType(e).ToContentSourceTypeOutputWithContext(ctx).ToContentSourceTypePtrOutputWithContext(ctx)
}

func (e ContentSourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentSourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentSourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContentSourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContentSourceTypeOutput struct{ *pulumi.OutputState }

func (ContentSourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSourceType)(nil)).Elem()
}

func (o ContentSourceTypeOutput) ToContentSourceTypeOutput() ContentSourceTypeOutput {
	return o
}

func (o ContentSourceTypeOutput) ToContentSourceTypeOutputWithContext(ctx context.Context) ContentSourceTypeOutput {
	return o
}

func (o ContentSourceTypeOutput) ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput {
	return o.ToContentSourceTypePtrOutputWithContext(context.Background())
}

func (o ContentSourceTypeOutput) ToContentSourceTypePtrOutputWithContext(ctx context.Context) ContentSourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentSourceType) *ContentSourceType {
		return &v
	}).(ContentSourceTypePtrOutput)
}

func (o ContentSourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContentSourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentSourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContentSourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentSourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentSourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContentSourceTypePtrOutput struct{ *pulumi.OutputState }

func (ContentSourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentSourceType)(nil)).Elem()
}

func (o ContentSourceTypePtrOutput) ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput {
	return o
}

func (o ContentSourceTypePtrOutput) ToContentSourceTypePtrOutputWithContext(ctx context.Context) ContentSourceTypePtrOutput {
	return o
}

func (o ContentSourceTypePtrOutput) Elem() ContentSourceTypeOutput {
	return o.ApplyT(func(v *ContentSourceType) ContentSourceType {
		if v != nil {
			return *v
		}
		var ret ContentSourceType
		return ret
	}).(ContentSourceTypeOutput)
}

func (o ContentSourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentSourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContentSourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContentSourceTypeInput interface {
	pulumi.Input

	ToContentSourceTypeOutput() ContentSourceTypeOutput
	ToContentSourceTypeOutputWithContext(context.Context) ContentSourceTypeOutput
}

var contentSourceTypePtrType = reflect.TypeOf((**ContentSourceType)(nil)).Elem()

type ContentSourceTypePtrInput interface {
	pulumi.Input

	ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput
	ToContentSourceTypePtrOutputWithContext(context.Context) ContentSourceTypePtrOutput
}

type contentSourceTypePtr string

func ContentSourceTypePtr(v string) ContentSourceTypePtrInput {
	return (*contentSourceTypePtr)(&v)
}

func (*contentSourceTypePtr) ElementType() reflect.Type {
	return contentSourceTypePtrType
}

func (in *contentSourceTypePtr) ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput {
	return pulumi.ToOutput(in).(ContentSourceTypePtrOutput)
}

func (in *contentSourceTypePtr) ToContentSourceTypePtrOutputWithContext(ctx context.Context) ContentSourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContentSourceTypePtrOutput)
}

type RunbookTypeEnum string

const (
	RunbookTypeEnumScript                  = RunbookTypeEnum("Script")
	RunbookTypeEnumGraph                   = RunbookTypeEnum("Graph")
	RunbookTypeEnumPowerShellWorkflow      = RunbookTypeEnum("PowerShellWorkflow")
	RunbookTypeEnumPowerShell              = RunbookTypeEnum("PowerShell")
	RunbookTypeEnumGraphPowerShellWorkflow = RunbookTypeEnum("GraphPowerShellWorkflow")
	RunbookTypeEnumGraphPowerShell         = RunbookTypeEnum("GraphPowerShell")
)

func (RunbookTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookTypeEnum)(nil)).Elem()
}

func (e RunbookTypeEnum) ToRunbookTypeEnumOutput() RunbookTypeEnumOutput {
	return pulumi.ToOutput(e).(RunbookTypeEnumOutput)
}

func (e RunbookTypeEnum) ToRunbookTypeEnumOutputWithContext(ctx context.Context) RunbookTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RunbookTypeEnumOutput)
}

func (e RunbookTypeEnum) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return e.ToRunbookTypeEnumPtrOutputWithContext(context.Background())
}

func (e RunbookTypeEnum) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return RunbookTypeEnum(e).ToRunbookTypeEnumOutputWithContext(ctx).ToRunbookTypeEnumPtrOutputWithContext(ctx)
}

func (e RunbookTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RunbookTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RunbookTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RunbookTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RunbookTypeEnumOutput struct{ *pulumi.OutputState }

func (RunbookTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookTypeEnum)(nil)).Elem()
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumOutput() RunbookTypeEnumOutput {
	return o
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumOutputWithContext(ctx context.Context) RunbookTypeEnumOutput {
	return o
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return o.ToRunbookTypeEnumPtrOutputWithContext(context.Background())
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookTypeEnum) *RunbookTypeEnum {
		return &v
	}).(RunbookTypeEnumPtrOutput)
}

func (o RunbookTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RunbookTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RunbookTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RunbookTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RunbookTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RunbookTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RunbookTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (RunbookTypeEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookTypeEnum)(nil)).Elem()
}

func (o RunbookTypeEnumPtrOutput) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return o
}

func (o RunbookTypeEnumPtrOutput) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return o
}

func (o RunbookTypeEnumPtrOutput) Elem() RunbookTypeEnumOutput {
	return o.ApplyT(func(v *RunbookTypeEnum) RunbookTypeEnum {
		if v != nil {
			return *v
		}
		var ret RunbookTypeEnum
		return ret
	}).(RunbookTypeEnumOutput)
}

func (o RunbookTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RunbookTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RunbookTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RunbookTypeEnumInput interface {
	pulumi.Input

	ToRunbookTypeEnumOutput() RunbookTypeEnumOutput
	ToRunbookTypeEnumOutputWithContext(context.Context) RunbookTypeEnumOutput
}

var runbookTypeEnumPtrType = reflect.TypeOf((**RunbookTypeEnum)(nil)).Elem()

type RunbookTypeEnumPtrInput interface {
	pulumi.Input

	ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput
	ToRunbookTypeEnumPtrOutputWithContext(context.Context) RunbookTypeEnumPtrOutput
}

type runbookTypeEnumPtr string

func RunbookTypeEnumPtr(v string) RunbookTypeEnumPtrInput {
	return (*runbookTypeEnumPtr)(&v)
}

func (*runbookTypeEnumPtr) ElementType() reflect.Type {
	return runbookTypeEnumPtrType
}

func (in *runbookTypeEnumPtr) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(RunbookTypeEnumPtrOutput)
}

func (in *runbookTypeEnumPtr) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RunbookTypeEnumPtrOutput)
}

type ScheduleDay string

const (
	ScheduleDayMonday    = ScheduleDay("Monday")
	ScheduleDayTuesday   = ScheduleDay("Tuesday")
	ScheduleDayWednesday = ScheduleDay("Wednesday")
	ScheduleDayThursday  = ScheduleDay("Thursday")
	ScheduleDayFriday    = ScheduleDay("Friday")
	ScheduleDaySaturday  = ScheduleDay("Saturday")
	ScheduleDaySunday    = ScheduleDay("Sunday")
)

func (ScheduleDay) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleDay)(nil)).Elem()
}

func (e ScheduleDay) ToScheduleDayOutput() ScheduleDayOutput {
	return pulumi.ToOutput(e).(ScheduleDayOutput)
}

func (e ScheduleDay) ToScheduleDayOutputWithContext(ctx context.Context) ScheduleDayOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduleDayOutput)
}

func (e ScheduleDay) ToScheduleDayPtrOutput() ScheduleDayPtrOutput {
	return e.ToScheduleDayPtrOutputWithContext(context.Background())
}

func (e ScheduleDay) ToScheduleDayPtrOutputWithContext(ctx context.Context) ScheduleDayPtrOutput {
	return ScheduleDay(e).ToScheduleDayOutputWithContext(ctx).ToScheduleDayPtrOutputWithContext(ctx)
}

func (e ScheduleDay) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleDay) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleDay) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduleDay) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduleDayOutput struct{ *pulumi.OutputState }

func (ScheduleDayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleDay)(nil)).Elem()
}

func (o ScheduleDayOutput) ToScheduleDayOutput() ScheduleDayOutput {
	return o
}

func (o ScheduleDayOutput) ToScheduleDayOutputWithContext(ctx context.Context) ScheduleDayOutput {
	return o
}

func (o ScheduleDayOutput) ToScheduleDayPtrOutput() ScheduleDayPtrOutput {
	return o.ToScheduleDayPtrOutputWithContext(context.Background())
}

func (o ScheduleDayOutput) ToScheduleDayPtrOutputWithContext(ctx context.Context) ScheduleDayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleDay) *ScheduleDay {
		return &v
	}).(ScheduleDayPtrOutput)
}

func (o ScheduleDayOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduleDayOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleDay) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduleDayOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleDayOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleDay) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduleDayPtrOutput struct{ *pulumi.OutputState }

func (ScheduleDayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleDay)(nil)).Elem()
}

func (o ScheduleDayPtrOutput) ToScheduleDayPtrOutput() ScheduleDayPtrOutput {
	return o
}

func (o ScheduleDayPtrOutput) ToScheduleDayPtrOutputWithContext(ctx context.Context) ScheduleDayPtrOutput {
	return o
}

func (o ScheduleDayPtrOutput) Elem() ScheduleDayOutput {
	return o.ApplyT(func(v *ScheduleDay) ScheduleDay {
		if v != nil {
			return *v
		}
		var ret ScheduleDay
		return ret
	}).(ScheduleDayOutput)
}

func (o ScheduleDayPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleDayPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduleDay) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScheduleDayInput interface {
	pulumi.Input

	ToScheduleDayOutput() ScheduleDayOutput
	ToScheduleDayOutputWithContext(context.Context) ScheduleDayOutput
}

var scheduleDayPtrType = reflect.TypeOf((**ScheduleDay)(nil)).Elem()

type ScheduleDayPtrInput interface {
	pulumi.Input

	ToScheduleDayPtrOutput() ScheduleDayPtrOutput
	ToScheduleDayPtrOutputWithContext(context.Context) ScheduleDayPtrOutput
}

type scheduleDayPtr string

func ScheduleDayPtr(v string) ScheduleDayPtrInput {
	return (*scheduleDayPtr)(&v)
}

func (*scheduleDayPtr) ElementType() reflect.Type {
	return scheduleDayPtrType
}

func (in *scheduleDayPtr) ToScheduleDayPtrOutput() ScheduleDayPtrOutput {
	return pulumi.ToOutput(in).(ScheduleDayPtrOutput)
}

func (in *scheduleDayPtr) ToScheduleDayPtrOutputWithContext(ctx context.Context) ScheduleDayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduleDayPtrOutput)
}

type ScheduleFrequency string

const (
	ScheduleFrequencyOneTime = ScheduleFrequency("OneTime")
	ScheduleFrequencyDay     = ScheduleFrequency("Day")
	ScheduleFrequencyHour    = ScheduleFrequency("Hour")
	ScheduleFrequencyWeek    = ScheduleFrequency("Week")
	ScheduleFrequencyMonth   = ScheduleFrequency("Month")
	// The minimum allowed interval for Minute schedules is 15 minutes.
	ScheduleFrequencyMinute = ScheduleFrequency("Minute")
)

func (ScheduleFrequency) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleFrequency)(nil)).Elem()
}

func (e ScheduleFrequency) ToScheduleFrequencyOutput() ScheduleFrequencyOutput {
	return pulumi.ToOutput(e).(ScheduleFrequencyOutput)
}

func (e ScheduleFrequency) ToScheduleFrequencyOutputWithContext(ctx context.Context) ScheduleFrequencyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduleFrequencyOutput)
}

func (e ScheduleFrequency) ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput {
	return e.ToScheduleFrequencyPtrOutputWithContext(context.Background())
}

func (e ScheduleFrequency) ToScheduleFrequencyPtrOutputWithContext(ctx context.Context) ScheduleFrequencyPtrOutput {
	return ScheduleFrequency(e).ToScheduleFrequencyOutputWithContext(ctx).ToScheduleFrequencyPtrOutputWithContext(ctx)
}

func (e ScheduleFrequency) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleFrequency) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleFrequency) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduleFrequency) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduleFrequencyOutput struct{ *pulumi.OutputState }

func (ScheduleFrequencyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleFrequency)(nil)).Elem()
}

func (o ScheduleFrequencyOutput) ToScheduleFrequencyOutput() ScheduleFrequencyOutput {
	return o
}

func (o ScheduleFrequencyOutput) ToScheduleFrequencyOutputWithContext(ctx context.Context) ScheduleFrequencyOutput {
	return o
}

func (o ScheduleFrequencyOutput) ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput {
	return o.ToScheduleFrequencyPtrOutputWithContext(context.Background())
}

func (o ScheduleFrequencyOutput) ToScheduleFrequencyPtrOutputWithContext(ctx context.Context) ScheduleFrequencyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleFrequency) *ScheduleFrequency {
		return &v
	}).(ScheduleFrequencyPtrOutput)
}

func (o ScheduleFrequencyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduleFrequencyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleFrequency) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduleFrequencyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleFrequencyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleFrequency) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduleFrequencyPtrOutput struct{ *pulumi.OutputState }

func (ScheduleFrequencyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleFrequency)(nil)).Elem()
}

func (o ScheduleFrequencyPtrOutput) ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput {
	return o
}

func (o ScheduleFrequencyPtrOutput) ToScheduleFrequencyPtrOutputWithContext(ctx context.Context) ScheduleFrequencyPtrOutput {
	return o
}

func (o ScheduleFrequencyPtrOutput) Elem() ScheduleFrequencyOutput {
	return o.ApplyT(func(v *ScheduleFrequency) ScheduleFrequency {
		if v != nil {
			return *v
		}
		var ret ScheduleFrequency
		return ret
	}).(ScheduleFrequencyOutput)
}

func (o ScheduleFrequencyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleFrequencyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduleFrequency) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScheduleFrequencyInput interface {
	pulumi.Input

	ToScheduleFrequencyOutput() ScheduleFrequencyOutput
	ToScheduleFrequencyOutputWithContext(context.Context) ScheduleFrequencyOutput
}

var scheduleFrequencyPtrType = reflect.TypeOf((**ScheduleFrequency)(nil)).Elem()

type ScheduleFrequencyPtrInput interface {
	pulumi.Input

	ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput
	ToScheduleFrequencyPtrOutputWithContext(context.Context) ScheduleFrequencyPtrOutput
}

type scheduleFrequencyPtr string

func ScheduleFrequencyPtr(v string) ScheduleFrequencyPtrInput {
	return (*scheduleFrequencyPtr)(&v)
}

func (*scheduleFrequencyPtr) ElementType() reflect.Type {
	return scheduleFrequencyPtrType
}

func (in *scheduleFrequencyPtr) ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput {
	return pulumi.ToOutput(in).(ScheduleFrequencyPtrOutput)
}

func (in *scheduleFrequencyPtr) ToScheduleFrequencyPtrOutputWithContext(ctx context.Context) ScheduleFrequencyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduleFrequencyPtrOutput)
}

type SkuNameEnum string

const (
	SkuNameEnumFree  = SkuNameEnum("Free")
	SkuNameEnumBasic = SkuNameEnum("Basic")
)

func (SkuNameEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuNameEnum)(nil)).Elem()
}

func (e SkuNameEnum) ToSkuNameEnumOutput() SkuNameEnumOutput {
	return pulumi.ToOutput(e).(SkuNameEnumOutput)
}

func (e SkuNameEnum) ToSkuNameEnumOutputWithContext(ctx context.Context) SkuNameEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameEnumOutput)
}

func (e SkuNameEnum) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return e.ToSkuNameEnumPtrOutputWithContext(context.Background())
}

func (e SkuNameEnum) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return SkuNameEnum(e).ToSkuNameEnumOutputWithContext(ctx).ToSkuNameEnumPtrOutputWithContext(ctx)
}

func (e SkuNameEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuNameEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuNameEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuNameEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameEnumOutput struct{ *pulumi.OutputState }

func (SkuNameEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuNameEnum)(nil)).Elem()
}

func (o SkuNameEnumOutput) ToSkuNameEnumOutput() SkuNameEnumOutput {
	return o
}

func (o SkuNameEnumOutput) ToSkuNameEnumOutputWithContext(ctx context.Context) SkuNameEnumOutput {
	return o
}

func (o SkuNameEnumOutput) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return o.ToSkuNameEnumPtrOutputWithContext(context.Background())
}

func (o SkuNameEnumOutput) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuNameEnum) *SkuNameEnum {
		return &v
	}).(SkuNameEnumPtrOutput)
}

func (o SkuNameEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuNameEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuNameEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNameEnumPtrOutput struct{ *pulumi.OutputState }

func (SkuNameEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuNameEnum)(nil)).Elem()
}

func (o SkuNameEnumPtrOutput) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return o
}

func (o SkuNameEnumPtrOutput) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return o
}

func (o SkuNameEnumPtrOutput) Elem() SkuNameEnumOutput {
	return o.ApplyT(func(v *SkuNameEnum) SkuNameEnum {
		if v != nil {
			return *v
		}
		var ret SkuNameEnum
		return ret
	}).(SkuNameEnumOutput)
}

func (o SkuNameEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuNameEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameEnumInput interface {
	pulumi.Input

	ToSkuNameEnumOutput() SkuNameEnumOutput
	ToSkuNameEnumOutputWithContext(context.Context) SkuNameEnumOutput
}

var skuNameEnumPtrType = reflect.TypeOf((**SkuNameEnum)(nil)).Elem()

type SkuNameEnumPtrInput interface {
	pulumi.Input

	ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput
	ToSkuNameEnumPtrOutputWithContext(context.Context) SkuNameEnumPtrOutput
}

type skuNameEnumPtr string

func SkuNameEnumPtr(v string) SkuNameEnumPtrInput {
	return (*skuNameEnumPtr)(&v)
}

func (*skuNameEnumPtr) ElementType() reflect.Type {
	return skuNameEnumPtrType
}

func (in *skuNameEnumPtr) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return pulumi.ToOutput(in).(SkuNameEnumPtrOutput)
}

func (in *skuNameEnumPtr) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNameEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentSourceTypeOutput{})
	pulumi.RegisterOutputType(ContentSourceTypePtrOutput{})
	pulumi.RegisterOutputType(RunbookTypeEnumOutput{})
	pulumi.RegisterOutputType(RunbookTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(ScheduleDayOutput{})
	pulumi.RegisterOutputType(ScheduleDayPtrOutput{})
	pulumi.RegisterOutputType(ScheduleFrequencyOutput{})
	pulumi.RegisterOutputType(ScheduleFrequencyPtrOutput{})
	pulumi.RegisterOutputType(SkuNameEnumOutput{})
	pulumi.RegisterOutputType(SkuNameEnumPtrOutput{})
}
