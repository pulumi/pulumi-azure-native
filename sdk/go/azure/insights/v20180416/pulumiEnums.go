


package v20180416

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertSeverity string

const (
	AlertSeverityZero  = AlertSeverity("0")
	AlertSeverityOne   = AlertSeverity("1")
	AlertSeverityTwo   = AlertSeverity("2")
	AlertSeverityThree = AlertSeverity("3")
	AlertSeverityFour  = AlertSeverity("4")
)

func (AlertSeverity) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertSeverity)(nil)).Elem()
}

func (e AlertSeverity) ToAlertSeverityOutput() AlertSeverityOutput {
	return pulumi.ToOutput(e).(AlertSeverityOutput)
}

func (e AlertSeverity) ToAlertSeverityOutputWithContext(ctx context.Context) AlertSeverityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AlertSeverityOutput)
}

func (e AlertSeverity) ToAlertSeverityPtrOutput() AlertSeverityPtrOutput {
	return e.ToAlertSeverityPtrOutputWithContext(context.Background())
}

func (e AlertSeverity) ToAlertSeverityPtrOutputWithContext(ctx context.Context) AlertSeverityPtrOutput {
	return AlertSeverity(e).ToAlertSeverityOutputWithContext(ctx).ToAlertSeverityPtrOutputWithContext(ctx)
}

func (e AlertSeverity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertSeverity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertSeverity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AlertSeverity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AlertSeverityOutput struct{ *pulumi.OutputState }

func (AlertSeverityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertSeverity)(nil)).Elem()
}

func (o AlertSeverityOutput) ToAlertSeverityOutput() AlertSeverityOutput {
	return o
}

func (o AlertSeverityOutput) ToAlertSeverityOutputWithContext(ctx context.Context) AlertSeverityOutput {
	return o
}

func (o AlertSeverityOutput) ToAlertSeverityPtrOutput() AlertSeverityPtrOutput {
	return o.ToAlertSeverityPtrOutputWithContext(context.Background())
}

func (o AlertSeverityOutput) ToAlertSeverityPtrOutputWithContext(ctx context.Context) AlertSeverityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertSeverity) *AlertSeverity {
		return &v
	}).(AlertSeverityPtrOutput)
}

func (o AlertSeverityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AlertSeverityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlertSeverity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AlertSeverityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlertSeverityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlertSeverity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AlertSeverityPtrOutput struct{ *pulumi.OutputState }

func (AlertSeverityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertSeverity)(nil)).Elem()
}

func (o AlertSeverityPtrOutput) ToAlertSeverityPtrOutput() AlertSeverityPtrOutput {
	return o
}

func (o AlertSeverityPtrOutput) ToAlertSeverityPtrOutputWithContext(ctx context.Context) AlertSeverityPtrOutput {
	return o
}

func (o AlertSeverityPtrOutput) Elem() AlertSeverityOutput {
	return o.ApplyT(func(v *AlertSeverity) AlertSeverity {
		if v != nil {
			return *v
		}
		var ret AlertSeverity
		return ret
	}).(AlertSeverityOutput)
}

func (o AlertSeverityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlertSeverityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AlertSeverity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AlertSeverityInput interface {
	pulumi.Input

	ToAlertSeverityOutput() AlertSeverityOutput
	ToAlertSeverityOutputWithContext(context.Context) AlertSeverityOutput
}

var alertSeverityPtrType = reflect.TypeOf((**AlertSeverity)(nil)).Elem()

type AlertSeverityPtrInput interface {
	pulumi.Input

	ToAlertSeverityPtrOutput() AlertSeverityPtrOutput
	ToAlertSeverityPtrOutputWithContext(context.Context) AlertSeverityPtrOutput
}

type alertSeverityPtr string

func AlertSeverityPtr(v string) AlertSeverityPtrInput {
	return (*alertSeverityPtr)(&v)
}

func (*alertSeverityPtr) ElementType() reflect.Type {
	return alertSeverityPtrType
}

func (in *alertSeverityPtr) ToAlertSeverityPtrOutput() AlertSeverityPtrOutput {
	return pulumi.ToOutput(in).(AlertSeverityPtrOutput)
}

func (in *alertSeverityPtr) ToAlertSeverityPtrOutputWithContext(ctx context.Context) AlertSeverityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AlertSeverityPtrOutput)
}

type ConditionalOperator string

const (
	ConditionalOperatorGreaterThanOrEqual = ConditionalOperator("GreaterThanOrEqual")
	ConditionalOperatorLessThanOrEqual    = ConditionalOperator("LessThanOrEqual")
	ConditionalOperatorGreaterThan        = ConditionalOperator("GreaterThan")
	ConditionalOperatorLessThan           = ConditionalOperator("LessThan")
	ConditionalOperatorEqual              = ConditionalOperator("Equal")
)

func (ConditionalOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionalOperator)(nil)).Elem()
}

func (e ConditionalOperator) ToConditionalOperatorOutput() ConditionalOperatorOutput {
	return pulumi.ToOutput(e).(ConditionalOperatorOutput)
}

func (e ConditionalOperator) ToConditionalOperatorOutputWithContext(ctx context.Context) ConditionalOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConditionalOperatorOutput)
}

func (e ConditionalOperator) ToConditionalOperatorPtrOutput() ConditionalOperatorPtrOutput {
	return e.ToConditionalOperatorPtrOutputWithContext(context.Background())
}

func (e ConditionalOperator) ToConditionalOperatorPtrOutputWithContext(ctx context.Context) ConditionalOperatorPtrOutput {
	return ConditionalOperator(e).ToConditionalOperatorOutputWithContext(ctx).ToConditionalOperatorPtrOutputWithContext(ctx)
}

func (e ConditionalOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConditionalOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConditionalOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConditionalOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConditionalOperatorOutput struct{ *pulumi.OutputState }

func (ConditionalOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionalOperator)(nil)).Elem()
}

func (o ConditionalOperatorOutput) ToConditionalOperatorOutput() ConditionalOperatorOutput {
	return o
}

func (o ConditionalOperatorOutput) ToConditionalOperatorOutputWithContext(ctx context.Context) ConditionalOperatorOutput {
	return o
}

func (o ConditionalOperatorOutput) ToConditionalOperatorPtrOutput() ConditionalOperatorPtrOutput {
	return o.ToConditionalOperatorPtrOutputWithContext(context.Background())
}

func (o ConditionalOperatorOutput) ToConditionalOperatorPtrOutputWithContext(ctx context.Context) ConditionalOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConditionalOperator) *ConditionalOperator {
		return &v
	}).(ConditionalOperatorPtrOutput)
}

func (o ConditionalOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConditionalOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConditionalOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConditionalOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConditionalOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConditionalOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConditionalOperatorPtrOutput struct{ *pulumi.OutputState }

func (ConditionalOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionalOperator)(nil)).Elem()
}

func (o ConditionalOperatorPtrOutput) ToConditionalOperatorPtrOutput() ConditionalOperatorPtrOutput {
	return o
}

func (o ConditionalOperatorPtrOutput) ToConditionalOperatorPtrOutputWithContext(ctx context.Context) ConditionalOperatorPtrOutput {
	return o
}

func (o ConditionalOperatorPtrOutput) Elem() ConditionalOperatorOutput {
	return o.ApplyT(func(v *ConditionalOperator) ConditionalOperator {
		if v != nil {
			return *v
		}
		var ret ConditionalOperator
		return ret
	}).(ConditionalOperatorOutput)
}

func (o ConditionalOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConditionalOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConditionalOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConditionalOperatorInput interface {
	pulumi.Input

	ToConditionalOperatorOutput() ConditionalOperatorOutput
	ToConditionalOperatorOutputWithContext(context.Context) ConditionalOperatorOutput
}

var conditionalOperatorPtrType = reflect.TypeOf((**ConditionalOperator)(nil)).Elem()

type ConditionalOperatorPtrInput interface {
	pulumi.Input

	ToConditionalOperatorPtrOutput() ConditionalOperatorPtrOutput
	ToConditionalOperatorPtrOutputWithContext(context.Context) ConditionalOperatorPtrOutput
}

type conditionalOperatorPtr string

func ConditionalOperatorPtr(v string) ConditionalOperatorPtrInput {
	return (*conditionalOperatorPtr)(&v)
}

func (*conditionalOperatorPtr) ElementType() reflect.Type {
	return conditionalOperatorPtrType
}

func (in *conditionalOperatorPtr) ToConditionalOperatorPtrOutput() ConditionalOperatorPtrOutput {
	return pulumi.ToOutput(in).(ConditionalOperatorPtrOutput)
}

func (in *conditionalOperatorPtr) ToConditionalOperatorPtrOutputWithContext(ctx context.Context) ConditionalOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConditionalOperatorPtrOutput)
}

type Enabled string

const (
	EnabledTrue  = Enabled("true")
	EnabledFalse = Enabled("false")
)

func (Enabled) ElementType() reflect.Type {
	return reflect.TypeOf((*Enabled)(nil)).Elem()
}

func (e Enabled) ToEnabledOutput() EnabledOutput {
	return pulumi.ToOutput(e).(EnabledOutput)
}

func (e Enabled) ToEnabledOutputWithContext(ctx context.Context) EnabledOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnabledOutput)
}

func (e Enabled) ToEnabledPtrOutput() EnabledPtrOutput {
	return e.ToEnabledPtrOutputWithContext(context.Background())
}

func (e Enabled) ToEnabledPtrOutputWithContext(ctx context.Context) EnabledPtrOutput {
	return Enabled(e).ToEnabledOutputWithContext(ctx).ToEnabledPtrOutputWithContext(ctx)
}

func (e Enabled) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Enabled) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Enabled) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Enabled) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnabledOutput struct{ *pulumi.OutputState }

func (EnabledOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Enabled)(nil)).Elem()
}

func (o EnabledOutput) ToEnabledOutput() EnabledOutput {
	return o
}

func (o EnabledOutput) ToEnabledOutputWithContext(ctx context.Context) EnabledOutput {
	return o
}

func (o EnabledOutput) ToEnabledPtrOutput() EnabledPtrOutput {
	return o.ToEnabledPtrOutputWithContext(context.Background())
}

func (o EnabledOutput) ToEnabledPtrOutputWithContext(ctx context.Context) EnabledPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Enabled) *Enabled {
		return &v
	}).(EnabledPtrOutput)
}

func (o EnabledOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnabledOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Enabled) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnabledOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnabledOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Enabled) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnabledPtrOutput struct{ *pulumi.OutputState }

func (EnabledPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Enabled)(nil)).Elem()
}

func (o EnabledPtrOutput) ToEnabledPtrOutput() EnabledPtrOutput {
	return o
}

func (o EnabledPtrOutput) ToEnabledPtrOutputWithContext(ctx context.Context) EnabledPtrOutput {
	return o
}

func (o EnabledPtrOutput) Elem() EnabledOutput {
	return o.ApplyT(func(v *Enabled) Enabled {
		if v != nil {
			return *v
		}
		var ret Enabled
		return ret
	}).(EnabledOutput)
}

func (o EnabledPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnabledPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Enabled) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EnabledInput interface {
	pulumi.Input

	ToEnabledOutput() EnabledOutput
	ToEnabledOutputWithContext(context.Context) EnabledOutput
}

var enabledPtrType = reflect.TypeOf((**Enabled)(nil)).Elem()

type EnabledPtrInput interface {
	pulumi.Input

	ToEnabledPtrOutput() EnabledPtrOutput
	ToEnabledPtrOutputWithContext(context.Context) EnabledPtrOutput
}

type enabledPtr string

func EnabledPtr(v string) EnabledPtrInput {
	return (*enabledPtr)(&v)
}

func (*enabledPtr) ElementType() reflect.Type {
	return enabledPtrType
}

func (in *enabledPtr) ToEnabledPtrOutput() EnabledPtrOutput {
	return pulumi.ToOutput(in).(EnabledPtrOutput)
}

func (in *enabledPtr) ToEnabledPtrOutputWithContext(ctx context.Context) EnabledPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnabledPtrOutput)
}

type MetricTriggerType string

const (
	MetricTriggerTypeConsecutive = MetricTriggerType("Consecutive")
	MetricTriggerTypeTotal       = MetricTriggerType("Total")
)

func (MetricTriggerType) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricTriggerType)(nil)).Elem()
}

func (e MetricTriggerType) ToMetricTriggerTypeOutput() MetricTriggerTypeOutput {
	return pulumi.ToOutput(e).(MetricTriggerTypeOutput)
}

func (e MetricTriggerType) ToMetricTriggerTypeOutputWithContext(ctx context.Context) MetricTriggerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MetricTriggerTypeOutput)
}

func (e MetricTriggerType) ToMetricTriggerTypePtrOutput() MetricTriggerTypePtrOutput {
	return e.ToMetricTriggerTypePtrOutputWithContext(context.Background())
}

func (e MetricTriggerType) ToMetricTriggerTypePtrOutputWithContext(ctx context.Context) MetricTriggerTypePtrOutput {
	return MetricTriggerType(e).ToMetricTriggerTypeOutputWithContext(ctx).ToMetricTriggerTypePtrOutputWithContext(ctx)
}

func (e MetricTriggerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MetricTriggerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MetricTriggerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MetricTriggerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MetricTriggerTypeOutput struct{ *pulumi.OutputState }

func (MetricTriggerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricTriggerType)(nil)).Elem()
}

func (o MetricTriggerTypeOutput) ToMetricTriggerTypeOutput() MetricTriggerTypeOutput {
	return o
}

func (o MetricTriggerTypeOutput) ToMetricTriggerTypeOutputWithContext(ctx context.Context) MetricTriggerTypeOutput {
	return o
}

func (o MetricTriggerTypeOutput) ToMetricTriggerTypePtrOutput() MetricTriggerTypePtrOutput {
	return o.ToMetricTriggerTypePtrOutputWithContext(context.Background())
}

func (o MetricTriggerTypeOutput) ToMetricTriggerTypePtrOutputWithContext(ctx context.Context) MetricTriggerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetricTriggerType) *MetricTriggerType {
		return &v
	}).(MetricTriggerTypePtrOutput)
}

func (o MetricTriggerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MetricTriggerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MetricTriggerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MetricTriggerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MetricTriggerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MetricTriggerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MetricTriggerTypePtrOutput struct{ *pulumi.OutputState }

func (MetricTriggerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricTriggerType)(nil)).Elem()
}

func (o MetricTriggerTypePtrOutput) ToMetricTriggerTypePtrOutput() MetricTriggerTypePtrOutput {
	return o
}

func (o MetricTriggerTypePtrOutput) ToMetricTriggerTypePtrOutputWithContext(ctx context.Context) MetricTriggerTypePtrOutput {
	return o
}

func (o MetricTriggerTypePtrOutput) Elem() MetricTriggerTypeOutput {
	return o.ApplyT(func(v *MetricTriggerType) MetricTriggerType {
		if v != nil {
			return *v
		}
		var ret MetricTriggerType
		return ret
	}).(MetricTriggerTypeOutput)
}

func (o MetricTriggerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MetricTriggerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MetricTriggerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MetricTriggerTypeInput interface {
	pulumi.Input

	ToMetricTriggerTypeOutput() MetricTriggerTypeOutput
	ToMetricTriggerTypeOutputWithContext(context.Context) MetricTriggerTypeOutput
}

var metricTriggerTypePtrType = reflect.TypeOf((**MetricTriggerType)(nil)).Elem()

type MetricTriggerTypePtrInput interface {
	pulumi.Input

	ToMetricTriggerTypePtrOutput() MetricTriggerTypePtrOutput
	ToMetricTriggerTypePtrOutputWithContext(context.Context) MetricTriggerTypePtrOutput
}

type metricTriggerTypePtr string

func MetricTriggerTypePtr(v string) MetricTriggerTypePtrInput {
	return (*metricTriggerTypePtr)(&v)
}

func (*metricTriggerTypePtr) ElementType() reflect.Type {
	return metricTriggerTypePtrType
}

func (in *metricTriggerTypePtr) ToMetricTriggerTypePtrOutput() MetricTriggerTypePtrOutput {
	return pulumi.ToOutput(in).(MetricTriggerTypePtrOutput)
}

func (in *metricTriggerTypePtr) ToMetricTriggerTypePtrOutputWithContext(ctx context.Context) MetricTriggerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MetricTriggerTypePtrOutput)
}

type Operator string

const (
	OperatorInclude = Operator("Include")
)

func (Operator) ElementType() reflect.Type {
	return reflect.TypeOf((*Operator)(nil)).Elem()
}

func (e Operator) ToOperatorOutput() OperatorOutput {
	return pulumi.ToOutput(e).(OperatorOutput)
}

func (e Operator) ToOperatorOutputWithContext(ctx context.Context) OperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatorOutput)
}

func (e Operator) ToOperatorPtrOutput() OperatorPtrOutput {
	return e.ToOperatorPtrOutputWithContext(context.Background())
}

func (e Operator) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return Operator(e).ToOperatorOutputWithContext(ctx).ToOperatorPtrOutputWithContext(ctx)
}

func (e Operator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Operator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatorOutput struct{ *pulumi.OutputState }

func (OperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Operator)(nil)).Elem()
}

func (o OperatorOutput) ToOperatorOutput() OperatorOutput {
	return o
}

func (o OperatorOutput) ToOperatorOutputWithContext(ctx context.Context) OperatorOutput {
	return o
}

func (o OperatorOutput) ToOperatorPtrOutput() OperatorPtrOutput {
	return o.ToOperatorPtrOutputWithContext(context.Background())
}

func (o OperatorOutput) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Operator) *Operator {
		return &v
	}).(OperatorPtrOutput)
}

func (o OperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatorPtrOutput struct{ *pulumi.OutputState }

func (OperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Operator)(nil)).Elem()
}

func (o OperatorPtrOutput) ToOperatorPtrOutput() OperatorPtrOutput {
	return o
}

func (o OperatorPtrOutput) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return o
}

func (o OperatorPtrOutput) Elem() OperatorOutput {
	return o.ApplyT(func(v *Operator) Operator {
		if v != nil {
			return *v
		}
		var ret Operator
		return ret
	}).(OperatorOutput)
}

func (o OperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Operator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatorInput interface {
	pulumi.Input

	ToOperatorOutput() OperatorOutput
	ToOperatorOutputWithContext(context.Context) OperatorOutput
}

var operatorPtrType = reflect.TypeOf((**Operator)(nil)).Elem()

type OperatorPtrInput interface {
	pulumi.Input

	ToOperatorPtrOutput() OperatorPtrOutput
	ToOperatorPtrOutputWithContext(context.Context) OperatorPtrOutput
}

type operatorPtr string

func OperatorPtr(v string) OperatorPtrInput {
	return (*operatorPtr)(&v)
}

func (*operatorPtr) ElementType() reflect.Type {
	return operatorPtrType
}

func (in *operatorPtr) ToOperatorPtrOutput() OperatorPtrOutput {
	return pulumi.ToOutput(in).(OperatorPtrOutput)
}

func (in *operatorPtr) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatorPtrOutput)
}

type QueryType string

const (
	QueryTypeResultCount = QueryType("ResultCount")
)

func (QueryType) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryType)(nil)).Elem()
}

func (e QueryType) ToQueryTypeOutput() QueryTypeOutput {
	return pulumi.ToOutput(e).(QueryTypeOutput)
}

func (e QueryType) ToQueryTypeOutputWithContext(ctx context.Context) QueryTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(QueryTypeOutput)
}

func (e QueryType) ToQueryTypePtrOutput() QueryTypePtrOutput {
	return e.ToQueryTypePtrOutputWithContext(context.Background())
}

func (e QueryType) ToQueryTypePtrOutputWithContext(ctx context.Context) QueryTypePtrOutput {
	return QueryType(e).ToQueryTypeOutputWithContext(ctx).ToQueryTypePtrOutputWithContext(ctx)
}

func (e QueryType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e QueryType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e QueryType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e QueryType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type QueryTypeOutput struct{ *pulumi.OutputState }

func (QueryTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryType)(nil)).Elem()
}

func (o QueryTypeOutput) ToQueryTypeOutput() QueryTypeOutput {
	return o
}

func (o QueryTypeOutput) ToQueryTypeOutputWithContext(ctx context.Context) QueryTypeOutput {
	return o
}

func (o QueryTypeOutput) ToQueryTypePtrOutput() QueryTypePtrOutput {
	return o.ToQueryTypePtrOutputWithContext(context.Background())
}

func (o QueryTypeOutput) ToQueryTypePtrOutputWithContext(ctx context.Context) QueryTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryType) *QueryType {
		return &v
	}).(QueryTypePtrOutput)
}

func (o QueryTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o QueryTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e QueryType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o QueryTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o QueryTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e QueryType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type QueryTypePtrOutput struct{ *pulumi.OutputState }

func (QueryTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryType)(nil)).Elem()
}

func (o QueryTypePtrOutput) ToQueryTypePtrOutput() QueryTypePtrOutput {
	return o
}

func (o QueryTypePtrOutput) ToQueryTypePtrOutputWithContext(ctx context.Context) QueryTypePtrOutput {
	return o
}

func (o QueryTypePtrOutput) Elem() QueryTypeOutput {
	return o.ApplyT(func(v *QueryType) QueryType {
		if v != nil {
			return *v
		}
		var ret QueryType
		return ret
	}).(QueryTypeOutput)
}

func (o QueryTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o QueryTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *QueryType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type QueryTypeInput interface {
	pulumi.Input

	ToQueryTypeOutput() QueryTypeOutput
	ToQueryTypeOutputWithContext(context.Context) QueryTypeOutput
}

var queryTypePtrType = reflect.TypeOf((**QueryType)(nil)).Elem()

type QueryTypePtrInput interface {
	pulumi.Input

	ToQueryTypePtrOutput() QueryTypePtrOutput
	ToQueryTypePtrOutputWithContext(context.Context) QueryTypePtrOutput
}

type queryTypePtr string

func QueryTypePtr(v string) QueryTypePtrInput {
	return (*queryTypePtr)(&v)
}

func (*queryTypePtr) ElementType() reflect.Type {
	return queryTypePtrType
}

func (in *queryTypePtr) ToQueryTypePtrOutput() QueryTypePtrOutput {
	return pulumi.ToOutput(in).(QueryTypePtrOutput)
}

func (in *queryTypePtr) ToQueryTypePtrOutputWithContext(ctx context.Context) QueryTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(QueryTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertSeverityInput)(nil)).Elem(), AlertSeverity("0"))
	pulumi.RegisterInputType(reflect.TypeOf((*AlertSeverityPtrInput)(nil)).Elem(), AlertSeverity("0"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionalOperatorInput)(nil)).Elem(), ConditionalOperator("GreaterThanOrEqual"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionalOperatorPtrInput)(nil)).Elem(), ConditionalOperator("GreaterThanOrEqual"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledInput)(nil)).Elem(), Enabled("true"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledPtrInput)(nil)).Elem(), Enabled("true"))
	pulumi.RegisterInputType(reflect.TypeOf((*MetricTriggerTypeInput)(nil)).Elem(), MetricTriggerType("Consecutive"))
	pulumi.RegisterInputType(reflect.TypeOf((*MetricTriggerTypePtrInput)(nil)).Elem(), MetricTriggerType("Consecutive"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorInput)(nil)).Elem(), Operator("Include"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorPtrInput)(nil)).Elem(), Operator("Include"))
	pulumi.RegisterInputType(reflect.TypeOf((*QueryTypeInput)(nil)).Elem(), QueryType("ResultCount"))
	pulumi.RegisterInputType(reflect.TypeOf((*QueryTypePtrInput)(nil)).Elem(), QueryType("ResultCount"))
	pulumi.RegisterOutputType(AlertSeverityOutput{})
	pulumi.RegisterOutputType(AlertSeverityPtrOutput{})
	pulumi.RegisterOutputType(ConditionalOperatorOutput{})
	pulumi.RegisterOutputType(ConditionalOperatorPtrOutput{})
	pulumi.RegisterOutputType(EnabledOutput{})
	pulumi.RegisterOutputType(EnabledPtrOutput{})
	pulumi.RegisterOutputType(MetricTriggerTypeOutput{})
	pulumi.RegisterOutputType(MetricTriggerTypePtrOutput{})
	pulumi.RegisterOutputType(OperatorOutput{})
	pulumi.RegisterOutputType(OperatorPtrOutput{})
	pulumi.RegisterOutputType(QueryTypeOutput{})
	pulumi.RegisterOutputType(QueryTypePtrOutput{})
}
