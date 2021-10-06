


package v20181102privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionRuleStatus string

const (
	ActionRuleStatusEnabled  = ActionRuleStatus("enabled")
	ActionRuleStatusDisabled = ActionRuleStatus("disabled")
)

func (ActionRuleStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRuleStatus)(nil)).Elem()
}

func (e ActionRuleStatus) ToActionRuleStatusOutput() ActionRuleStatusOutput {
	return pulumi.ToOutput(e).(ActionRuleStatusOutput)
}

func (e ActionRuleStatus) ToActionRuleStatusOutputWithContext(ctx context.Context) ActionRuleStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionRuleStatusOutput)
}

func (e ActionRuleStatus) ToActionRuleStatusPtrOutput() ActionRuleStatusPtrOutput {
	return e.ToActionRuleStatusPtrOutputWithContext(context.Background())
}

func (e ActionRuleStatus) ToActionRuleStatusPtrOutputWithContext(ctx context.Context) ActionRuleStatusPtrOutput {
	return ActionRuleStatus(e).ToActionRuleStatusOutputWithContext(ctx).ToActionRuleStatusPtrOutputWithContext(ctx)
}

func (e ActionRuleStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionRuleStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionRuleStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ActionRuleStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionRuleStatusOutput struct{ *pulumi.OutputState }

func (ActionRuleStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRuleStatus)(nil)).Elem()
}

func (o ActionRuleStatusOutput) ToActionRuleStatusOutput() ActionRuleStatusOutput {
	return o
}

func (o ActionRuleStatusOutput) ToActionRuleStatusOutputWithContext(ctx context.Context) ActionRuleStatusOutput {
	return o
}

func (o ActionRuleStatusOutput) ToActionRuleStatusPtrOutput() ActionRuleStatusPtrOutput {
	return o.ToActionRuleStatusPtrOutputWithContext(context.Background())
}

func (o ActionRuleStatusOutput) ToActionRuleStatusPtrOutputWithContext(ctx context.Context) ActionRuleStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionRuleStatus) *ActionRuleStatus {
		return &v
	}).(ActionRuleStatusPtrOutput)
}

func (o ActionRuleStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionRuleStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionRuleStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionRuleStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionRuleStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionRuleStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionRuleStatusPtrOutput struct{ *pulumi.OutputState }

func (ActionRuleStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleStatus)(nil)).Elem()
}

func (o ActionRuleStatusPtrOutput) ToActionRuleStatusPtrOutput() ActionRuleStatusPtrOutput {
	return o
}

func (o ActionRuleStatusPtrOutput) ToActionRuleStatusPtrOutputWithContext(ctx context.Context) ActionRuleStatusPtrOutput {
	return o
}

func (o ActionRuleStatusPtrOutput) Elem() ActionRuleStatusOutput {
	return o.ApplyT(func(v *ActionRuleStatus) ActionRuleStatus {
		if v != nil {
			return *v
		}
		var ret ActionRuleStatus
		return ret
	}).(ActionRuleStatusOutput)
}

func (o ActionRuleStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionRuleStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ActionRuleStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ActionRuleStatusInput interface {
	pulumi.Input

	ToActionRuleStatusOutput() ActionRuleStatusOutput
	ToActionRuleStatusOutputWithContext(context.Context) ActionRuleStatusOutput
}

var actionRuleStatusPtrType = reflect.TypeOf((**ActionRuleStatus)(nil)).Elem()

type ActionRuleStatusPtrInput interface {
	pulumi.Input

	ToActionRuleStatusPtrOutput() ActionRuleStatusPtrOutput
	ToActionRuleStatusPtrOutputWithContext(context.Context) ActionRuleStatusPtrOutput
}

type actionRuleStatusPtr string

func ActionRuleStatusPtr(v string) ActionRuleStatusPtrInput {
	return (*actionRuleStatusPtr)(&v)
}

func (*actionRuleStatusPtr) ElementType() reflect.Type {
	return actionRuleStatusPtrType
}

func (in *actionRuleStatusPtr) ToActionRuleStatusPtrOutput() ActionRuleStatusPtrOutput {
	return pulumi.ToOutput(in).(ActionRuleStatusPtrOutput)
}

func (in *actionRuleStatusPtr) ToActionRuleStatusPtrOutputWithContext(ctx context.Context) ActionRuleStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionRuleStatusPtrOutput)
}

type ScopeType string

const (
	ScopeTypeResourceGroup = ScopeType("ResourceGroup")
	ScopeTypeResource      = ScopeType("Resource")
)

func (ScopeType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeType)(nil)).Elem()
}

func (e ScopeType) ToScopeTypeOutput() ScopeTypeOutput {
	return pulumi.ToOutput(e).(ScopeTypeOutput)
}

func (e ScopeType) ToScopeTypeOutputWithContext(ctx context.Context) ScopeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScopeTypeOutput)
}

func (e ScopeType) ToScopeTypePtrOutput() ScopeTypePtrOutput {
	return e.ToScopeTypePtrOutputWithContext(context.Background())
}

func (e ScopeType) ToScopeTypePtrOutputWithContext(ctx context.Context) ScopeTypePtrOutput {
	return ScopeType(e).ToScopeTypeOutputWithContext(ctx).ToScopeTypePtrOutputWithContext(ctx)
}

func (e ScopeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScopeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScopeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScopeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScopeTypeOutput struct{ *pulumi.OutputState }

func (ScopeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeType)(nil)).Elem()
}

func (o ScopeTypeOutput) ToScopeTypeOutput() ScopeTypeOutput {
	return o
}

func (o ScopeTypeOutput) ToScopeTypeOutputWithContext(ctx context.Context) ScopeTypeOutput {
	return o
}

func (o ScopeTypeOutput) ToScopeTypePtrOutput() ScopeTypePtrOutput {
	return o.ToScopeTypePtrOutputWithContext(context.Background())
}

func (o ScopeTypeOutput) ToScopeTypePtrOutputWithContext(ctx context.Context) ScopeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeType) *ScopeType {
		return &v
	}).(ScopeTypePtrOutput)
}

func (o ScopeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScopeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScopeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScopeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScopeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScopeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScopeTypePtrOutput struct{ *pulumi.OutputState }

func (ScopeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeType)(nil)).Elem()
}

func (o ScopeTypePtrOutput) ToScopeTypePtrOutput() ScopeTypePtrOutput {
	return o
}

func (o ScopeTypePtrOutput) ToScopeTypePtrOutputWithContext(ctx context.Context) ScopeTypePtrOutput {
	return o
}

func (o ScopeTypePtrOutput) Elem() ScopeTypeOutput {
	return o.ApplyT(func(v *ScopeType) ScopeType {
		if v != nil {
			return *v
		}
		var ret ScopeType
		return ret
	}).(ScopeTypeOutput)
}

func (o ScopeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScopeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScopeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScopeTypeInput interface {
	pulumi.Input

	ToScopeTypeOutput() ScopeTypeOutput
	ToScopeTypeOutputWithContext(context.Context) ScopeTypeOutput
}

var scopeTypePtrType = reflect.TypeOf((**ScopeType)(nil)).Elem()

type ScopeTypePtrInput interface {
	pulumi.Input

	ToScopeTypePtrOutput() ScopeTypePtrOutput
	ToScopeTypePtrOutputWithContext(context.Context) ScopeTypePtrOutput
}

type scopeTypePtr string

func ScopeTypePtr(v string) ScopeTypePtrInput {
	return (*scopeTypePtr)(&v)
}

func (*scopeTypePtr) ElementType() reflect.Type {
	return scopeTypePtrType
}

func (in *scopeTypePtr) ToScopeTypePtrOutput() ScopeTypePtrOutput {
	return pulumi.ToOutput(in).(ScopeTypePtrOutput)
}

func (in *scopeTypePtr) ToScopeTypePtrOutputWithContext(ctx context.Context) ScopeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScopeTypePtrOutput)
}

type SuppressionType string

const (
	SuppressionTypeAlways  = SuppressionType("Always")
	SuppressionTypeOnce    = SuppressionType("Once")
	SuppressionTypeDaily   = SuppressionType("Daily")
	SuppressionTypeWeekly  = SuppressionType("Weekly")
	SuppressionTypeMonthly = SuppressionType("Monthly")
)

func (SuppressionType) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionType)(nil)).Elem()
}

func (e SuppressionType) ToSuppressionTypeOutput() SuppressionTypeOutput {
	return pulumi.ToOutput(e).(SuppressionTypeOutput)
}

func (e SuppressionType) ToSuppressionTypeOutputWithContext(ctx context.Context) SuppressionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SuppressionTypeOutput)
}

func (e SuppressionType) ToSuppressionTypePtrOutput() SuppressionTypePtrOutput {
	return e.ToSuppressionTypePtrOutputWithContext(context.Background())
}

func (e SuppressionType) ToSuppressionTypePtrOutputWithContext(ctx context.Context) SuppressionTypePtrOutput {
	return SuppressionType(e).ToSuppressionTypeOutputWithContext(ctx).ToSuppressionTypePtrOutputWithContext(ctx)
}

func (e SuppressionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SuppressionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SuppressionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SuppressionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SuppressionTypeOutput struct{ *pulumi.OutputState }

func (SuppressionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionType)(nil)).Elem()
}

func (o SuppressionTypeOutput) ToSuppressionTypeOutput() SuppressionTypeOutput {
	return o
}

func (o SuppressionTypeOutput) ToSuppressionTypeOutputWithContext(ctx context.Context) SuppressionTypeOutput {
	return o
}

func (o SuppressionTypeOutput) ToSuppressionTypePtrOutput() SuppressionTypePtrOutput {
	return o.ToSuppressionTypePtrOutputWithContext(context.Background())
}

func (o SuppressionTypeOutput) ToSuppressionTypePtrOutputWithContext(ctx context.Context) SuppressionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SuppressionType) *SuppressionType {
		return &v
	}).(SuppressionTypePtrOutput)
}

func (o SuppressionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SuppressionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SuppressionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SuppressionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SuppressionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SuppressionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SuppressionTypePtrOutput struct{ *pulumi.OutputState }

func (SuppressionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionType)(nil)).Elem()
}

func (o SuppressionTypePtrOutput) ToSuppressionTypePtrOutput() SuppressionTypePtrOutput {
	return o
}

func (o SuppressionTypePtrOutput) ToSuppressionTypePtrOutputWithContext(ctx context.Context) SuppressionTypePtrOutput {
	return o
}

func (o SuppressionTypePtrOutput) Elem() SuppressionTypeOutput {
	return o.ApplyT(func(v *SuppressionType) SuppressionType {
		if v != nil {
			return *v
		}
		var ret SuppressionType
		return ret
	}).(SuppressionTypeOutput)
}

func (o SuppressionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SuppressionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SuppressionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SuppressionTypeInput interface {
	pulumi.Input

	ToSuppressionTypeOutput() SuppressionTypeOutput
	ToSuppressionTypeOutputWithContext(context.Context) SuppressionTypeOutput
}

var suppressionTypePtrType = reflect.TypeOf((**SuppressionType)(nil)).Elem()

type SuppressionTypePtrInput interface {
	pulumi.Input

	ToSuppressionTypePtrOutput() SuppressionTypePtrOutput
	ToSuppressionTypePtrOutputWithContext(context.Context) SuppressionTypePtrOutput
}

type suppressionTypePtr string

func SuppressionTypePtr(v string) SuppressionTypePtrInput {
	return (*suppressionTypePtr)(&v)
}

func (*suppressionTypePtr) ElementType() reflect.Type {
	return suppressionTypePtrType
}

func (in *suppressionTypePtr) ToSuppressionTypePtrOutput() SuppressionTypePtrOutput {
	return pulumi.ToOutput(in).(SuppressionTypePtrOutput)
}

func (in *suppressionTypePtr) ToSuppressionTypePtrOutputWithContext(ctx context.Context) SuppressionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SuppressionTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionRuleStatusOutput{})
	pulumi.RegisterOutputType(ActionRuleStatusPtrOutput{})
	pulumi.RegisterOutputType(ScopeTypeOutput{})
	pulumi.RegisterOutputType(ScopeTypePtrOutput{})
	pulumi.RegisterOutputType(SuppressionTypeOutput{})
	pulumi.RegisterOutputType(SuppressionTypePtrOutput{})
}
