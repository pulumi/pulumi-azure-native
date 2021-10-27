


package v20190301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertRuleState string

const (
	AlertRuleStateEnabled  = AlertRuleState("Enabled")
	AlertRuleStateDisabled = AlertRuleState("Disabled")
)

func (AlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleState)(nil)).Elem()
}

func (e AlertRuleState) ToAlertRuleStateOutput() AlertRuleStateOutput {
	return pulumi.ToOutput(e).(AlertRuleStateOutput)
}

func (e AlertRuleState) ToAlertRuleStateOutputWithContext(ctx context.Context) AlertRuleStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AlertRuleStateOutput)
}

func (e AlertRuleState) ToAlertRuleStatePtrOutput() AlertRuleStatePtrOutput {
	return e.ToAlertRuleStatePtrOutputWithContext(context.Background())
}

func (e AlertRuleState) ToAlertRuleStatePtrOutputWithContext(ctx context.Context) AlertRuleStatePtrOutput {
	return AlertRuleState(e).ToAlertRuleStateOutputWithContext(ctx).ToAlertRuleStatePtrOutputWithContext(ctx)
}

func (e AlertRuleState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertRuleState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertRuleState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AlertRuleState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AlertRuleStateOutput struct{ *pulumi.OutputState }

func (AlertRuleStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleState)(nil)).Elem()
}

func (o AlertRuleStateOutput) ToAlertRuleStateOutput() AlertRuleStateOutput {
	return o
}

func (o AlertRuleStateOutput) ToAlertRuleStateOutputWithContext(ctx context.Context) AlertRuleStateOutput {
	return o
}

func (o AlertRuleStateOutput) ToAlertRuleStatePtrOutput() AlertRuleStatePtrOutput {
	return o.ToAlertRuleStatePtrOutputWithContext(context.Background())
}

func (o AlertRuleStateOutput) ToAlertRuleStatePtrOutputWithContext(ctx context.Context) AlertRuleStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertRuleState) *AlertRuleState {
		return &v
	}).(AlertRuleStatePtrOutput)
}

func (o AlertRuleStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AlertRuleStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlertRuleState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AlertRuleStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlertRuleStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlertRuleState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AlertRuleStatePtrOutput struct{ *pulumi.OutputState }

func (AlertRuleStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleState)(nil)).Elem()
}

func (o AlertRuleStatePtrOutput) ToAlertRuleStatePtrOutput() AlertRuleStatePtrOutput {
	return o
}

func (o AlertRuleStatePtrOutput) ToAlertRuleStatePtrOutputWithContext(ctx context.Context) AlertRuleStatePtrOutput {
	return o
}

func (o AlertRuleStatePtrOutput) Elem() AlertRuleStateOutput {
	return o.ApplyT(func(v *AlertRuleState) AlertRuleState {
		if v != nil {
			return *v
		}
		var ret AlertRuleState
		return ret
	}).(AlertRuleStateOutput)
}

func (o AlertRuleStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlertRuleStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AlertRuleState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AlertRuleStateInput interface {
	pulumi.Input

	ToAlertRuleStateOutput() AlertRuleStateOutput
	ToAlertRuleStateOutputWithContext(context.Context) AlertRuleStateOutput
}

var alertRuleStatePtrType = reflect.TypeOf((**AlertRuleState)(nil)).Elem()

type AlertRuleStatePtrInput interface {
	pulumi.Input

	ToAlertRuleStatePtrOutput() AlertRuleStatePtrOutput
	ToAlertRuleStatePtrOutputWithContext(context.Context) AlertRuleStatePtrOutput
}

type alertRuleStatePtr string

func AlertRuleStatePtr(v string) AlertRuleStatePtrInput {
	return (*alertRuleStatePtr)(&v)
}

func (*alertRuleStatePtr) ElementType() reflect.Type {
	return alertRuleStatePtrType
}

func (in *alertRuleStatePtr) ToAlertRuleStatePtrOutput() AlertRuleStatePtrOutput {
	return pulumi.ToOutput(in).(AlertRuleStatePtrOutput)
}

func (in *alertRuleStatePtr) ToAlertRuleStatePtrOutputWithContext(ctx context.Context) AlertRuleStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AlertRuleStatePtrOutput)
}

type Severity string

const (
	SeveritySev0 = Severity("Sev0")
	SeveritySev1 = Severity("Sev1")
	SeveritySev2 = Severity("Sev2")
	SeveritySev3 = Severity("Sev3")
	SeveritySev4 = Severity("Sev4")
)

func (Severity) ElementType() reflect.Type {
	return reflect.TypeOf((*Severity)(nil)).Elem()
}

func (e Severity) ToSeverityOutput() SeverityOutput {
	return pulumi.ToOutput(e).(SeverityOutput)
}

func (e Severity) ToSeverityOutputWithContext(ctx context.Context) SeverityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SeverityOutput)
}

func (e Severity) ToSeverityPtrOutput() SeverityPtrOutput {
	return e.ToSeverityPtrOutputWithContext(context.Background())
}

func (e Severity) ToSeverityPtrOutputWithContext(ctx context.Context) SeverityPtrOutput {
	return Severity(e).ToSeverityOutputWithContext(ctx).ToSeverityPtrOutputWithContext(ctx)
}

func (e Severity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Severity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Severity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Severity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SeverityOutput struct{ *pulumi.OutputState }

func (SeverityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Severity)(nil)).Elem()
}

func (o SeverityOutput) ToSeverityOutput() SeverityOutput {
	return o
}

func (o SeverityOutput) ToSeverityOutputWithContext(ctx context.Context) SeverityOutput {
	return o
}

func (o SeverityOutput) ToSeverityPtrOutput() SeverityPtrOutput {
	return o.ToSeverityPtrOutputWithContext(context.Background())
}

func (o SeverityOutput) ToSeverityPtrOutputWithContext(ctx context.Context) SeverityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Severity) *Severity {
		return &v
	}).(SeverityPtrOutput)
}

func (o SeverityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SeverityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Severity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SeverityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SeverityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Severity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SeverityPtrOutput struct{ *pulumi.OutputState }

func (SeverityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Severity)(nil)).Elem()
}

func (o SeverityPtrOutput) ToSeverityPtrOutput() SeverityPtrOutput {
	return o
}

func (o SeverityPtrOutput) ToSeverityPtrOutputWithContext(ctx context.Context) SeverityPtrOutput {
	return o
}

func (o SeverityPtrOutput) Elem() SeverityOutput {
	return o.ApplyT(func(v *Severity) Severity {
		if v != nil {
			return *v
		}
		var ret Severity
		return ret
	}).(SeverityOutput)
}

func (o SeverityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SeverityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Severity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SeverityInput interface {
	pulumi.Input

	ToSeverityOutput() SeverityOutput
	ToSeverityOutputWithContext(context.Context) SeverityOutput
}

var severityPtrType = reflect.TypeOf((**Severity)(nil)).Elem()

type SeverityPtrInput interface {
	pulumi.Input

	ToSeverityPtrOutput() SeverityPtrOutput
	ToSeverityPtrOutputWithContext(context.Context) SeverityPtrOutput
}

type severityPtr string

func SeverityPtr(v string) SeverityPtrInput {
	return (*severityPtr)(&v)
}

func (*severityPtr) ElementType() reflect.Type {
	return severityPtrType
}

func (in *severityPtr) ToSeverityPtrOutput() SeverityPtrOutput {
	return pulumi.ToOutput(in).(SeverityPtrOutput)
}

func (in *severityPtr) ToSeverityPtrOutputWithContext(ctx context.Context) SeverityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SeverityPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertRuleStateInput)(nil)).Elem(), AlertRuleState("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*AlertRuleStatePtrInput)(nil)).Elem(), AlertRuleState("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*SeverityInput)(nil)).Elem(), Severity("Sev0"))
	pulumi.RegisterInputType(reflect.TypeOf((*SeverityPtrInput)(nil)).Elem(), Severity("Sev0"))
	pulumi.RegisterOutputType(AlertRuleStateOutput{})
	pulumi.RegisterOutputType(AlertRuleStatePtrOutput{})
	pulumi.RegisterOutputType(SeverityOutput{})
	pulumi.RegisterOutputType(SeverityPtrOutput{})
}
