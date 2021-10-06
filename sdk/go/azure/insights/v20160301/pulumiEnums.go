


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConditionOperator string

const (
	ConditionOperatorGreaterThan        = ConditionOperator("GreaterThan")
	ConditionOperatorGreaterThanOrEqual = ConditionOperator("GreaterThanOrEqual")
	ConditionOperatorLessThan           = ConditionOperator("LessThan")
	ConditionOperatorLessThanOrEqual    = ConditionOperator("LessThanOrEqual")
)

func (ConditionOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionOperator)(nil)).Elem()
}

func (e ConditionOperator) ToConditionOperatorOutput() ConditionOperatorOutput {
	return pulumi.ToOutput(e).(ConditionOperatorOutput)
}

func (e ConditionOperator) ToConditionOperatorOutputWithContext(ctx context.Context) ConditionOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConditionOperatorOutput)
}

func (e ConditionOperator) ToConditionOperatorPtrOutput() ConditionOperatorPtrOutput {
	return e.ToConditionOperatorPtrOutputWithContext(context.Background())
}

func (e ConditionOperator) ToConditionOperatorPtrOutputWithContext(ctx context.Context) ConditionOperatorPtrOutput {
	return ConditionOperator(e).ToConditionOperatorOutputWithContext(ctx).ToConditionOperatorPtrOutputWithContext(ctx)
}

func (e ConditionOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConditionOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConditionOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConditionOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConditionOperatorOutput struct{ *pulumi.OutputState }

func (ConditionOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionOperator)(nil)).Elem()
}

func (o ConditionOperatorOutput) ToConditionOperatorOutput() ConditionOperatorOutput {
	return o
}

func (o ConditionOperatorOutput) ToConditionOperatorOutputWithContext(ctx context.Context) ConditionOperatorOutput {
	return o
}

func (o ConditionOperatorOutput) ToConditionOperatorPtrOutput() ConditionOperatorPtrOutput {
	return o.ToConditionOperatorPtrOutputWithContext(context.Background())
}

func (o ConditionOperatorOutput) ToConditionOperatorPtrOutputWithContext(ctx context.Context) ConditionOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConditionOperator) *ConditionOperator {
		return &v
	}).(ConditionOperatorPtrOutput)
}

func (o ConditionOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConditionOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConditionOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConditionOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConditionOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConditionOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConditionOperatorPtrOutput struct{ *pulumi.OutputState }

func (ConditionOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionOperator)(nil)).Elem()
}

func (o ConditionOperatorPtrOutput) ToConditionOperatorPtrOutput() ConditionOperatorPtrOutput {
	return o
}

func (o ConditionOperatorPtrOutput) ToConditionOperatorPtrOutputWithContext(ctx context.Context) ConditionOperatorPtrOutput {
	return o
}

func (o ConditionOperatorPtrOutput) Elem() ConditionOperatorOutput {
	return o.ApplyT(func(v *ConditionOperator) ConditionOperator {
		if v != nil {
			return *v
		}
		var ret ConditionOperator
		return ret
	}).(ConditionOperatorOutput)
}

func (o ConditionOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConditionOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConditionOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConditionOperatorInput interface {
	pulumi.Input

	ToConditionOperatorOutput() ConditionOperatorOutput
	ToConditionOperatorOutputWithContext(context.Context) ConditionOperatorOutput
}

var conditionOperatorPtrType = reflect.TypeOf((**ConditionOperator)(nil)).Elem()

type ConditionOperatorPtrInput interface {
	pulumi.Input

	ToConditionOperatorPtrOutput() ConditionOperatorPtrOutput
	ToConditionOperatorPtrOutputWithContext(context.Context) ConditionOperatorPtrOutput
}

type conditionOperatorPtr string

func ConditionOperatorPtr(v string) ConditionOperatorPtrInput {
	return (*conditionOperatorPtr)(&v)
}

func (*conditionOperatorPtr) ElementType() reflect.Type {
	return conditionOperatorPtrType
}

func (in *conditionOperatorPtr) ToConditionOperatorPtrOutput() ConditionOperatorPtrOutput {
	return pulumi.ToOutput(in).(ConditionOperatorPtrOutput)
}

func (in *conditionOperatorPtr) ToConditionOperatorPtrOutputWithContext(ctx context.Context) ConditionOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConditionOperatorPtrOutput)
}

type TimeAggregationOperator string

const (
	TimeAggregationOperatorAverage = TimeAggregationOperator("Average")
	TimeAggregationOperatorMinimum = TimeAggregationOperator("Minimum")
	TimeAggregationOperatorMaximum = TimeAggregationOperator("Maximum")
	TimeAggregationOperatorTotal   = TimeAggregationOperator("Total")
	TimeAggregationOperatorLast    = TimeAggregationOperator("Last")
)

func (TimeAggregationOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeAggregationOperator)(nil)).Elem()
}

func (e TimeAggregationOperator) ToTimeAggregationOperatorOutput() TimeAggregationOperatorOutput {
	return pulumi.ToOutput(e).(TimeAggregationOperatorOutput)
}

func (e TimeAggregationOperator) ToTimeAggregationOperatorOutputWithContext(ctx context.Context) TimeAggregationOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TimeAggregationOperatorOutput)
}

func (e TimeAggregationOperator) ToTimeAggregationOperatorPtrOutput() TimeAggregationOperatorPtrOutput {
	return e.ToTimeAggregationOperatorPtrOutputWithContext(context.Background())
}

func (e TimeAggregationOperator) ToTimeAggregationOperatorPtrOutputWithContext(ctx context.Context) TimeAggregationOperatorPtrOutput {
	return TimeAggregationOperator(e).ToTimeAggregationOperatorOutputWithContext(ctx).ToTimeAggregationOperatorPtrOutputWithContext(ctx)
}

func (e TimeAggregationOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeAggregationOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeAggregationOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TimeAggregationOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TimeAggregationOperatorOutput struct{ *pulumi.OutputState }

func (TimeAggregationOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeAggregationOperator)(nil)).Elem()
}

func (o TimeAggregationOperatorOutput) ToTimeAggregationOperatorOutput() TimeAggregationOperatorOutput {
	return o
}

func (o TimeAggregationOperatorOutput) ToTimeAggregationOperatorOutputWithContext(ctx context.Context) TimeAggregationOperatorOutput {
	return o
}

func (o TimeAggregationOperatorOutput) ToTimeAggregationOperatorPtrOutput() TimeAggregationOperatorPtrOutput {
	return o.ToTimeAggregationOperatorPtrOutputWithContext(context.Background())
}

func (o TimeAggregationOperatorOutput) ToTimeAggregationOperatorPtrOutputWithContext(ctx context.Context) TimeAggregationOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeAggregationOperator) *TimeAggregationOperator {
		return &v
	}).(TimeAggregationOperatorPtrOutput)
}

func (o TimeAggregationOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TimeAggregationOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeAggregationOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TimeAggregationOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeAggregationOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeAggregationOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TimeAggregationOperatorPtrOutput struct{ *pulumi.OutputState }

func (TimeAggregationOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeAggregationOperator)(nil)).Elem()
}

func (o TimeAggregationOperatorPtrOutput) ToTimeAggregationOperatorPtrOutput() TimeAggregationOperatorPtrOutput {
	return o
}

func (o TimeAggregationOperatorPtrOutput) ToTimeAggregationOperatorPtrOutputWithContext(ctx context.Context) TimeAggregationOperatorPtrOutput {
	return o
}

func (o TimeAggregationOperatorPtrOutput) Elem() TimeAggregationOperatorOutput {
	return o.ApplyT(func(v *TimeAggregationOperator) TimeAggregationOperator {
		if v != nil {
			return *v
		}
		var ret TimeAggregationOperator
		return ret
	}).(TimeAggregationOperatorOutput)
}

func (o TimeAggregationOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeAggregationOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TimeAggregationOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TimeAggregationOperatorInput interface {
	pulumi.Input

	ToTimeAggregationOperatorOutput() TimeAggregationOperatorOutput
	ToTimeAggregationOperatorOutputWithContext(context.Context) TimeAggregationOperatorOutput
}

var timeAggregationOperatorPtrType = reflect.TypeOf((**TimeAggregationOperator)(nil)).Elem()

type TimeAggregationOperatorPtrInput interface {
	pulumi.Input

	ToTimeAggregationOperatorPtrOutput() TimeAggregationOperatorPtrOutput
	ToTimeAggregationOperatorPtrOutputWithContext(context.Context) TimeAggregationOperatorPtrOutput
}

type timeAggregationOperatorPtr string

func TimeAggregationOperatorPtr(v string) TimeAggregationOperatorPtrInput {
	return (*timeAggregationOperatorPtr)(&v)
}

func (*timeAggregationOperatorPtr) ElementType() reflect.Type {
	return timeAggregationOperatorPtrType
}

func (in *timeAggregationOperatorPtr) ToTimeAggregationOperatorPtrOutput() TimeAggregationOperatorPtrOutput {
	return pulumi.ToOutput(in).(TimeAggregationOperatorPtrOutput)
}

func (in *timeAggregationOperatorPtr) ToTimeAggregationOperatorPtrOutputWithContext(ctx context.Context) TimeAggregationOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TimeAggregationOperatorPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConditionOperatorOutput{})
	pulumi.RegisterOutputType(ConditionOperatorPtrOutput{})
	pulumi.RegisterOutputType(TimeAggregationOperatorOutput{})
	pulumi.RegisterOutputType(TimeAggregationOperatorPtrOutput{})
}
