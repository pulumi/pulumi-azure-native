


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConditionOperator string

const (
	ConditionOperatorEquals             = ConditionOperator("Equals")
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

type DimensionOperator string

const (
	DimensionOperatorInclude = DimensionOperator("Include")
	DimensionOperatorExclude = DimensionOperator("Exclude")
)

func (DimensionOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*DimensionOperator)(nil)).Elem()
}

func (e DimensionOperator) ToDimensionOperatorOutput() DimensionOperatorOutput {
	return pulumi.ToOutput(e).(DimensionOperatorOutput)
}

func (e DimensionOperator) ToDimensionOperatorOutputWithContext(ctx context.Context) DimensionOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DimensionOperatorOutput)
}

func (e DimensionOperator) ToDimensionOperatorPtrOutput() DimensionOperatorPtrOutput {
	return e.ToDimensionOperatorPtrOutputWithContext(context.Background())
}

func (e DimensionOperator) ToDimensionOperatorPtrOutputWithContext(ctx context.Context) DimensionOperatorPtrOutput {
	return DimensionOperator(e).ToDimensionOperatorOutputWithContext(ctx).ToDimensionOperatorPtrOutputWithContext(ctx)
}

func (e DimensionOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DimensionOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DimensionOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DimensionOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DimensionOperatorOutput struct{ *pulumi.OutputState }

func (DimensionOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DimensionOperator)(nil)).Elem()
}

func (o DimensionOperatorOutput) ToDimensionOperatorOutput() DimensionOperatorOutput {
	return o
}

func (o DimensionOperatorOutput) ToDimensionOperatorOutputWithContext(ctx context.Context) DimensionOperatorOutput {
	return o
}

func (o DimensionOperatorOutput) ToDimensionOperatorPtrOutput() DimensionOperatorPtrOutput {
	return o.ToDimensionOperatorPtrOutputWithContext(context.Background())
}

func (o DimensionOperatorOutput) ToDimensionOperatorPtrOutputWithContext(ctx context.Context) DimensionOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DimensionOperator) *DimensionOperator {
		return &v
	}).(DimensionOperatorPtrOutput)
}

func (o DimensionOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DimensionOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DimensionOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DimensionOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DimensionOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DimensionOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DimensionOperatorPtrOutput struct{ *pulumi.OutputState }

func (DimensionOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DimensionOperator)(nil)).Elem()
}

func (o DimensionOperatorPtrOutput) ToDimensionOperatorPtrOutput() DimensionOperatorPtrOutput {
	return o
}

func (o DimensionOperatorPtrOutput) ToDimensionOperatorPtrOutputWithContext(ctx context.Context) DimensionOperatorPtrOutput {
	return o
}

func (o DimensionOperatorPtrOutput) Elem() DimensionOperatorOutput {
	return o.ApplyT(func(v *DimensionOperator) DimensionOperator {
		if v != nil {
			return *v
		}
		var ret DimensionOperator
		return ret
	}).(DimensionOperatorOutput)
}

func (o DimensionOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DimensionOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DimensionOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DimensionOperatorInput interface {
	pulumi.Input

	ToDimensionOperatorOutput() DimensionOperatorOutput
	ToDimensionOperatorOutputWithContext(context.Context) DimensionOperatorOutput
}

var dimensionOperatorPtrType = reflect.TypeOf((**DimensionOperator)(nil)).Elem()

type DimensionOperatorPtrInput interface {
	pulumi.Input

	ToDimensionOperatorPtrOutput() DimensionOperatorPtrOutput
	ToDimensionOperatorPtrOutputWithContext(context.Context) DimensionOperatorPtrOutput
}

type dimensionOperatorPtr string

func DimensionOperatorPtr(v string) DimensionOperatorPtrInput {
	return (*dimensionOperatorPtr)(&v)
}

func (*dimensionOperatorPtr) ElementType() reflect.Type {
	return dimensionOperatorPtrType
}

func (in *dimensionOperatorPtr) ToDimensionOperatorPtrOutput() DimensionOperatorPtrOutput {
	return pulumi.ToOutput(in).(DimensionOperatorPtrOutput)
}

func (in *dimensionOperatorPtr) ToDimensionOperatorPtrOutputWithContext(ctx context.Context) DimensionOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DimensionOperatorPtrOutput)
}

type Kind string

const (
	KindLogAlert    = Kind("LogAlert")
	KindLogToMetric = Kind("LogToMetric")
)

func (Kind) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (e Kind) ToKindOutput() KindOutput {
	return pulumi.ToOutput(e).(KindOutput)
}

func (e Kind) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KindOutput)
}

func (e Kind) ToKindPtrOutput() KindPtrOutput {
	return e.ToKindPtrOutputWithContext(context.Background())
}

func (e Kind) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return Kind(e).ToKindOutputWithContext(ctx).ToKindPtrOutputWithContext(ctx)
}

func (e Kind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Kind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KindOutput struct{ *pulumi.OutputState }

func (KindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (o KindOutput) ToKindOutput() KindOutput {
	return o
}

func (o KindOutput) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return o
}

func (o KindOutput) ToKindPtrOutput() KindPtrOutput {
	return o.ToKindPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Kind) *Kind {
		return &v
	}).(KindPtrOutput)
}

func (o KindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KindPtrOutput struct{ *pulumi.OutputState }

func (KindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kind)(nil)).Elem()
}

func (o KindPtrOutput) ToKindPtrOutput() KindPtrOutput {
	return o
}

func (o KindPtrOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o
}

func (o KindPtrOutput) Elem() KindOutput {
	return o.ApplyT(func(v *Kind) Kind {
		if v != nil {
			return *v
		}
		var ret Kind
		return ret
	}).(KindOutput)
}

func (o KindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Kind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KindInput interface {
	pulumi.Input

	ToKindOutput() KindOutput
	ToKindOutputWithContext(context.Context) KindOutput
}

var kindPtrType = reflect.TypeOf((**Kind)(nil)).Elem()

type KindPtrInput interface {
	pulumi.Input

	ToKindPtrOutput() KindPtrOutput
	ToKindPtrOutputWithContext(context.Context) KindPtrOutput
}

type kindPtr string

func KindPtr(v string) KindPtrInput {
	return (*kindPtr)(&v)
}

func (*kindPtr) ElementType() reflect.Type {
	return kindPtrType
}

func (in *kindPtr) ToKindPtrOutput() KindPtrOutput {
	return pulumi.ToOutput(in).(KindPtrOutput)
}

func (in *kindPtr) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KindPtrOutput)
}

type TimeAggregation string

const (
	TimeAggregationCount   = TimeAggregation("Count")
	TimeAggregationAverage = TimeAggregation("Average")
	TimeAggregationMinimum = TimeAggregation("Minimum")
	TimeAggregationMaximum = TimeAggregation("Maximum")
	TimeAggregationTotal   = TimeAggregation("Total")
)

func (TimeAggregation) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeAggregation)(nil)).Elem()
}

func (e TimeAggregation) ToTimeAggregationOutput() TimeAggregationOutput {
	return pulumi.ToOutput(e).(TimeAggregationOutput)
}

func (e TimeAggregation) ToTimeAggregationOutputWithContext(ctx context.Context) TimeAggregationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TimeAggregationOutput)
}

func (e TimeAggregation) ToTimeAggregationPtrOutput() TimeAggregationPtrOutput {
	return e.ToTimeAggregationPtrOutputWithContext(context.Background())
}

func (e TimeAggregation) ToTimeAggregationPtrOutputWithContext(ctx context.Context) TimeAggregationPtrOutput {
	return TimeAggregation(e).ToTimeAggregationOutputWithContext(ctx).ToTimeAggregationPtrOutputWithContext(ctx)
}

func (e TimeAggregation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeAggregation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeAggregation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TimeAggregation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TimeAggregationOutput struct{ *pulumi.OutputState }

func (TimeAggregationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeAggregation)(nil)).Elem()
}

func (o TimeAggregationOutput) ToTimeAggregationOutput() TimeAggregationOutput {
	return o
}

func (o TimeAggregationOutput) ToTimeAggregationOutputWithContext(ctx context.Context) TimeAggregationOutput {
	return o
}

func (o TimeAggregationOutput) ToTimeAggregationPtrOutput() TimeAggregationPtrOutput {
	return o.ToTimeAggregationPtrOutputWithContext(context.Background())
}

func (o TimeAggregationOutput) ToTimeAggregationPtrOutputWithContext(ctx context.Context) TimeAggregationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeAggregation) *TimeAggregation {
		return &v
	}).(TimeAggregationPtrOutput)
}

func (o TimeAggregationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TimeAggregationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeAggregation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TimeAggregationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeAggregationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeAggregation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TimeAggregationPtrOutput struct{ *pulumi.OutputState }

func (TimeAggregationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeAggregation)(nil)).Elem()
}

func (o TimeAggregationPtrOutput) ToTimeAggregationPtrOutput() TimeAggregationPtrOutput {
	return o
}

func (o TimeAggregationPtrOutput) ToTimeAggregationPtrOutputWithContext(ctx context.Context) TimeAggregationPtrOutput {
	return o
}

func (o TimeAggregationPtrOutput) Elem() TimeAggregationOutput {
	return o.ApplyT(func(v *TimeAggregation) TimeAggregation {
		if v != nil {
			return *v
		}
		var ret TimeAggregation
		return ret
	}).(TimeAggregationOutput)
}

func (o TimeAggregationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeAggregationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TimeAggregation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TimeAggregationInput interface {
	pulumi.Input

	ToTimeAggregationOutput() TimeAggregationOutput
	ToTimeAggregationOutputWithContext(context.Context) TimeAggregationOutput
}

var timeAggregationPtrType = reflect.TypeOf((**TimeAggregation)(nil)).Elem()

type TimeAggregationPtrInput interface {
	pulumi.Input

	ToTimeAggregationPtrOutput() TimeAggregationPtrOutput
	ToTimeAggregationPtrOutputWithContext(context.Context) TimeAggregationPtrOutput
}

type timeAggregationPtr string

func TimeAggregationPtr(v string) TimeAggregationPtrInput {
	return (*timeAggregationPtr)(&v)
}

func (*timeAggregationPtr) ElementType() reflect.Type {
	return timeAggregationPtrType
}

func (in *timeAggregationPtr) ToTimeAggregationPtrOutput() TimeAggregationPtrOutput {
	return pulumi.ToOutput(in).(TimeAggregationPtrOutput)
}

func (in *timeAggregationPtr) ToTimeAggregationPtrOutputWithContext(ctx context.Context) TimeAggregationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TimeAggregationPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionOperatorInput)(nil)).Elem(), ConditionOperator("Equals"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionOperatorPtrInput)(nil)).Elem(), ConditionOperator("Equals"))
	pulumi.RegisterInputType(reflect.TypeOf((*DimensionOperatorInput)(nil)).Elem(), DimensionOperator("Include"))
	pulumi.RegisterInputType(reflect.TypeOf((*DimensionOperatorPtrInput)(nil)).Elem(), DimensionOperator("Include"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindInput)(nil)).Elem(), Kind("LogAlert"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindPtrInput)(nil)).Elem(), Kind("LogAlert"))
	pulumi.RegisterInputType(reflect.TypeOf((*TimeAggregationInput)(nil)).Elem(), TimeAggregation("Count"))
	pulumi.RegisterInputType(reflect.TypeOf((*TimeAggregationPtrInput)(nil)).Elem(), TimeAggregation("Count"))
	pulumi.RegisterOutputType(ConditionOperatorOutput{})
	pulumi.RegisterOutputType(ConditionOperatorPtrOutput{})
	pulumi.RegisterOutputType(DimensionOperatorOutput{})
	pulumi.RegisterOutputType(DimensionOperatorPtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(TimeAggregationOutput{})
	pulumi.RegisterOutputType(TimeAggregationPtrOutput{})
}
