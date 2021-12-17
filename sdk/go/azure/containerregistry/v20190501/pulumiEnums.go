


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Action string

const (
	ActionAllow = Action("Allow")
)

func (Action) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (e Action) ToActionOutput() ActionOutput {
	return pulumi.ToOutput(e).(ActionOutput)
}

func (e Action) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionOutput)
}

func (e Action) ToActionPtrOutput() ActionPtrOutput {
	return e.ToActionPtrOutputWithContext(context.Background())
}

func (e Action) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return Action(e).ToActionOutputWithContext(ctx).ToActionPtrOutputWithContext(ctx)
}

func (e Action) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Action) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Action) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Action) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

func (o ActionOutput) ToActionPtrOutput() ActionPtrOutput {
	return o.ToActionPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Action) *Action {
		return &v
	}).(ActionPtrOutput)
}

func (o ActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Action) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Action) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionPtrOutput struct{ *pulumi.OutputState }

func (ActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionPtrOutput) ToActionPtrOutput() ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) Elem() ActionOutput {
	return o.ApplyT(func(v *Action) Action {
		if v != nil {
			return *v
		}
		var ret Action
		return ret
	}).(ActionOutput)
}

func (o ActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Action) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

var actionPtrType = reflect.TypeOf((**Action)(nil)).Elem()

type ActionPtrInput interface {
	pulumi.Input

	ToActionPtrOutput() ActionPtrOutput
	ToActionPtrOutputWithContext(context.Context) ActionPtrOutput
}

type actionPtr string

func ActionPtr(v string) ActionPtrInput {
	return (*actionPtr)(&v)
}

func (*actionPtr) ElementType() reflect.Type {
	return actionPtrType
}

func (in *actionPtr) ToActionPtrOutput() ActionPtrOutput {
	return pulumi.ToOutput(in).(ActionPtrOutput)
}

func (in *actionPtr) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionPtrOutput)
}

type DefaultAction string

const (
	DefaultActionAllow = DefaultAction("Allow")
	DefaultActionDeny  = DefaultAction("Deny")
)

func (DefaultAction) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAction)(nil)).Elem()
}

func (e DefaultAction) ToDefaultActionOutput() DefaultActionOutput {
	return pulumi.ToOutput(e).(DefaultActionOutput)
}

func (e DefaultAction) ToDefaultActionOutputWithContext(ctx context.Context) DefaultActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefaultActionOutput)
}

func (e DefaultAction) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return e.ToDefaultActionPtrOutputWithContext(context.Background())
}

func (e DefaultAction) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return DefaultAction(e).ToDefaultActionOutputWithContext(ctx).ToDefaultActionPtrOutputWithContext(ctx)
}

func (e DefaultAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefaultActionOutput struct{ *pulumi.OutputState }

func (DefaultActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAction)(nil)).Elem()
}

func (o DefaultActionOutput) ToDefaultActionOutput() DefaultActionOutput {
	return o
}

func (o DefaultActionOutput) ToDefaultActionOutputWithContext(ctx context.Context) DefaultActionOutput {
	return o
}

func (o DefaultActionOutput) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return o.ToDefaultActionPtrOutputWithContext(context.Background())
}

func (o DefaultActionOutput) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultAction) *DefaultAction {
		return &v
	}).(DefaultActionPtrOutput)
}

func (o DefaultActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefaultActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefaultActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefaultActionPtrOutput struct{ *pulumi.OutputState }

func (DefaultActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAction)(nil)).Elem()
}

func (o DefaultActionPtrOutput) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return o
}

func (o DefaultActionPtrOutput) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return o
}

func (o DefaultActionPtrOutput) Elem() DefaultActionOutput {
	return o.ApplyT(func(v *DefaultAction) DefaultAction {
		if v != nil {
			return *v
		}
		var ret DefaultAction
		return ret
	}).(DefaultActionOutput)
}

func (o DefaultActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefaultAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DefaultActionInput interface {
	pulumi.Input

	ToDefaultActionOutput() DefaultActionOutput
	ToDefaultActionOutputWithContext(context.Context) DefaultActionOutput
}

var defaultActionPtrType = reflect.TypeOf((**DefaultAction)(nil)).Elem()

type DefaultActionPtrInput interface {
	pulumi.Input

	ToDefaultActionPtrOutput() DefaultActionPtrOutput
	ToDefaultActionPtrOutputWithContext(context.Context) DefaultActionPtrOutput
}

type defaultActionPtr string

func DefaultActionPtr(v string) DefaultActionPtrInput {
	return (*defaultActionPtr)(&v)
}

func (*defaultActionPtr) ElementType() reflect.Type {
	return defaultActionPtrType
}

func (in *defaultActionPtr) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return pulumi.ToOutput(in).(DefaultActionPtrOutput)
}

func (in *defaultActionPtr) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefaultActionPtrOutput)
}

type PolicyStatus string

const (
	PolicyStatusEnabled  = PolicyStatus("enabled")
	PolicyStatusDisabled = PolicyStatus("disabled")
)

func (PolicyStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyStatus)(nil)).Elem()
}

func (e PolicyStatus) ToPolicyStatusOutput() PolicyStatusOutput {
	return pulumi.ToOutput(e).(PolicyStatusOutput)
}

func (e PolicyStatus) ToPolicyStatusOutputWithContext(ctx context.Context) PolicyStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyStatusOutput)
}

func (e PolicyStatus) ToPolicyStatusPtrOutput() PolicyStatusPtrOutput {
	return e.ToPolicyStatusPtrOutputWithContext(context.Background())
}

func (e PolicyStatus) ToPolicyStatusPtrOutputWithContext(ctx context.Context) PolicyStatusPtrOutput {
	return PolicyStatus(e).ToPolicyStatusOutputWithContext(ctx).ToPolicyStatusPtrOutputWithContext(ctx)
}

func (e PolicyStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyStatusOutput struct{ *pulumi.OutputState }

func (PolicyStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyStatus)(nil)).Elem()
}

func (o PolicyStatusOutput) ToPolicyStatusOutput() PolicyStatusOutput {
	return o
}

func (o PolicyStatusOutput) ToPolicyStatusOutputWithContext(ctx context.Context) PolicyStatusOutput {
	return o
}

func (o PolicyStatusOutput) ToPolicyStatusPtrOutput() PolicyStatusPtrOutput {
	return o.ToPolicyStatusPtrOutputWithContext(context.Background())
}

func (o PolicyStatusOutput) ToPolicyStatusPtrOutputWithContext(ctx context.Context) PolicyStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyStatus) *PolicyStatus {
		return &v
	}).(PolicyStatusPtrOutput)
}

func (o PolicyStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyStatusPtrOutput struct{ *pulumi.OutputState }

func (PolicyStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyStatus)(nil)).Elem()
}

func (o PolicyStatusPtrOutput) ToPolicyStatusPtrOutput() PolicyStatusPtrOutput {
	return o
}

func (o PolicyStatusPtrOutput) ToPolicyStatusPtrOutputWithContext(ctx context.Context) PolicyStatusPtrOutput {
	return o
}

func (o PolicyStatusPtrOutput) Elem() PolicyStatusOutput {
	return o.ApplyT(func(v *PolicyStatus) PolicyStatus {
		if v != nil {
			return *v
		}
		var ret PolicyStatus
		return ret
	}).(PolicyStatusOutput)
}

func (o PolicyStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PolicyStatusInput interface {
	pulumi.Input

	ToPolicyStatusOutput() PolicyStatusOutput
	ToPolicyStatusOutputWithContext(context.Context) PolicyStatusOutput
}

var policyStatusPtrType = reflect.TypeOf((**PolicyStatus)(nil)).Elem()

type PolicyStatusPtrInput interface {
	pulumi.Input

	ToPolicyStatusPtrOutput() PolicyStatusPtrOutput
	ToPolicyStatusPtrOutputWithContext(context.Context) PolicyStatusPtrOutput
}

type policyStatusPtr string

func PolicyStatusPtr(v string) PolicyStatusPtrInput {
	return (*policyStatusPtr)(&v)
}

func (*policyStatusPtr) ElementType() reflect.Type {
	return policyStatusPtrType
}

func (in *policyStatusPtr) ToPolicyStatusPtrOutput() PolicyStatusPtrOutput {
	return pulumi.ToOutput(in).(PolicyStatusPtrOutput)
}

func (in *policyStatusPtr) ToPolicyStatusPtrOutputWithContext(ctx context.Context) PolicyStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyStatusPtrOutput)
}

type SkuName string

const (
	SkuNameClassic  = SkuName("Classic")
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
	SkuNamePremium  = SkuName("Premium")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

type TrustPolicyType string

const (
	TrustPolicyTypeNotary = TrustPolicyType("Notary")
)

func (TrustPolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicyType)(nil)).Elem()
}

func (e TrustPolicyType) ToTrustPolicyTypeOutput() TrustPolicyTypeOutput {
	return pulumi.ToOutput(e).(TrustPolicyTypeOutput)
}

func (e TrustPolicyType) ToTrustPolicyTypeOutputWithContext(ctx context.Context) TrustPolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TrustPolicyTypeOutput)
}

func (e TrustPolicyType) ToTrustPolicyTypePtrOutput() TrustPolicyTypePtrOutput {
	return e.ToTrustPolicyTypePtrOutputWithContext(context.Background())
}

func (e TrustPolicyType) ToTrustPolicyTypePtrOutputWithContext(ctx context.Context) TrustPolicyTypePtrOutput {
	return TrustPolicyType(e).ToTrustPolicyTypeOutputWithContext(ctx).ToTrustPolicyTypePtrOutputWithContext(ctx)
}

func (e TrustPolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrustPolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrustPolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TrustPolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TrustPolicyTypeOutput struct{ *pulumi.OutputState }

func (TrustPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicyType)(nil)).Elem()
}

func (o TrustPolicyTypeOutput) ToTrustPolicyTypeOutput() TrustPolicyTypeOutput {
	return o
}

func (o TrustPolicyTypeOutput) ToTrustPolicyTypeOutputWithContext(ctx context.Context) TrustPolicyTypeOutput {
	return o
}

func (o TrustPolicyTypeOutput) ToTrustPolicyTypePtrOutput() TrustPolicyTypePtrOutput {
	return o.ToTrustPolicyTypePtrOutputWithContext(context.Background())
}

func (o TrustPolicyTypeOutput) ToTrustPolicyTypePtrOutputWithContext(ctx context.Context) TrustPolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrustPolicyType) *TrustPolicyType {
		return &v
	}).(TrustPolicyTypePtrOutput)
}

func (o TrustPolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TrustPolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrustPolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TrustPolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrustPolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrustPolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TrustPolicyTypePtrOutput struct{ *pulumi.OutputState }

func (TrustPolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicyType)(nil)).Elem()
}

func (o TrustPolicyTypePtrOutput) ToTrustPolicyTypePtrOutput() TrustPolicyTypePtrOutput {
	return o
}

func (o TrustPolicyTypePtrOutput) ToTrustPolicyTypePtrOutputWithContext(ctx context.Context) TrustPolicyTypePtrOutput {
	return o
}

func (o TrustPolicyTypePtrOutput) Elem() TrustPolicyTypeOutput {
	return o.ApplyT(func(v *TrustPolicyType) TrustPolicyType {
		if v != nil {
			return *v
		}
		var ret TrustPolicyType
		return ret
	}).(TrustPolicyTypeOutput)
}

func (o TrustPolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrustPolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TrustPolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TrustPolicyTypeInput interface {
	pulumi.Input

	ToTrustPolicyTypeOutput() TrustPolicyTypeOutput
	ToTrustPolicyTypeOutputWithContext(context.Context) TrustPolicyTypeOutput
}

var trustPolicyTypePtrType = reflect.TypeOf((**TrustPolicyType)(nil)).Elem()

type TrustPolicyTypePtrInput interface {
	pulumi.Input

	ToTrustPolicyTypePtrOutput() TrustPolicyTypePtrOutput
	ToTrustPolicyTypePtrOutputWithContext(context.Context) TrustPolicyTypePtrOutput
}

type trustPolicyTypePtr string

func TrustPolicyTypePtr(v string) TrustPolicyTypePtrInput {
	return (*trustPolicyTypePtr)(&v)
}

func (*trustPolicyTypePtr) ElementType() reflect.Type {
	return trustPolicyTypePtrType
}

func (in *trustPolicyTypePtr) ToTrustPolicyTypePtrOutput() TrustPolicyTypePtrOutput {
	return pulumi.ToOutput(in).(TrustPolicyTypePtrOutput)
}

func (in *trustPolicyTypePtr) ToTrustPolicyTypePtrOutputWithContext(ctx context.Context) TrustPolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TrustPolicyTypePtrOutput)
}

type WebhookAction string

const (
	WebhookActionPush          = WebhookAction("push")
	WebhookActionDelete        = WebhookAction("delete")
	WebhookActionQuarantine    = WebhookAction("quarantine")
	WebhookAction_Chart_push   = WebhookAction("chart_push")
	WebhookAction_Chart_delete = WebhookAction("chart_delete")
)

func (WebhookAction) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookAction)(nil)).Elem()
}

func (e WebhookAction) ToWebhookActionOutput() WebhookActionOutput {
	return pulumi.ToOutput(e).(WebhookActionOutput)
}

func (e WebhookAction) ToWebhookActionOutputWithContext(ctx context.Context) WebhookActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebhookActionOutput)
}

func (e WebhookAction) ToWebhookActionPtrOutput() WebhookActionPtrOutput {
	return e.ToWebhookActionPtrOutputWithContext(context.Background())
}

func (e WebhookAction) ToWebhookActionPtrOutputWithContext(ctx context.Context) WebhookActionPtrOutput {
	return WebhookAction(e).ToWebhookActionOutputWithContext(ctx).ToWebhookActionPtrOutputWithContext(ctx)
}

func (e WebhookAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebhookAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebhookAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebhookAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebhookActionOutput struct{ *pulumi.OutputState }

func (WebhookActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookAction)(nil)).Elem()
}

func (o WebhookActionOutput) ToWebhookActionOutput() WebhookActionOutput {
	return o
}

func (o WebhookActionOutput) ToWebhookActionOutputWithContext(ctx context.Context) WebhookActionOutput {
	return o
}

func (o WebhookActionOutput) ToWebhookActionPtrOutput() WebhookActionPtrOutput {
	return o.ToWebhookActionPtrOutputWithContext(context.Background())
}

func (o WebhookActionOutput) ToWebhookActionPtrOutputWithContext(ctx context.Context) WebhookActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebhookAction) *WebhookAction {
		return &v
	}).(WebhookActionPtrOutput)
}

func (o WebhookActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebhookActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebhookAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebhookActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebhookActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebhookAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebhookActionPtrOutput struct{ *pulumi.OutputState }

func (WebhookActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookAction)(nil)).Elem()
}

func (o WebhookActionPtrOutput) ToWebhookActionPtrOutput() WebhookActionPtrOutput {
	return o
}

func (o WebhookActionPtrOutput) ToWebhookActionPtrOutputWithContext(ctx context.Context) WebhookActionPtrOutput {
	return o
}

func (o WebhookActionPtrOutput) Elem() WebhookActionOutput {
	return o.ApplyT(func(v *WebhookAction) WebhookAction {
		if v != nil {
			return *v
		}
		var ret WebhookAction
		return ret
	}).(WebhookActionOutput)
}

func (o WebhookActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebhookActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebhookAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebhookActionInput interface {
	pulumi.Input

	ToWebhookActionOutput() WebhookActionOutput
	ToWebhookActionOutputWithContext(context.Context) WebhookActionOutput
}

var webhookActionPtrType = reflect.TypeOf((**WebhookAction)(nil)).Elem()

type WebhookActionPtrInput interface {
	pulumi.Input

	ToWebhookActionPtrOutput() WebhookActionPtrOutput
	ToWebhookActionPtrOutputWithContext(context.Context) WebhookActionPtrOutput
}

type webhookActionPtr string

func WebhookActionPtr(v string) WebhookActionPtrInput {
	return (*webhookActionPtr)(&v)
}

func (*webhookActionPtr) ElementType() reflect.Type {
	return webhookActionPtrType
}

func (in *webhookActionPtr) ToWebhookActionPtrOutput() WebhookActionPtrOutput {
	return pulumi.ToOutput(in).(WebhookActionPtrOutput)
}

func (in *webhookActionPtr) ToWebhookActionPtrOutputWithContext(ctx context.Context) WebhookActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebhookActionPtrOutput)
}

type WebhookStatus string

const (
	WebhookStatusEnabled  = WebhookStatus("enabled")
	WebhookStatusDisabled = WebhookStatus("disabled")
)

func (WebhookStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookStatus)(nil)).Elem()
}

func (e WebhookStatus) ToWebhookStatusOutput() WebhookStatusOutput {
	return pulumi.ToOutput(e).(WebhookStatusOutput)
}

func (e WebhookStatus) ToWebhookStatusOutputWithContext(ctx context.Context) WebhookStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebhookStatusOutput)
}

func (e WebhookStatus) ToWebhookStatusPtrOutput() WebhookStatusPtrOutput {
	return e.ToWebhookStatusPtrOutputWithContext(context.Background())
}

func (e WebhookStatus) ToWebhookStatusPtrOutputWithContext(ctx context.Context) WebhookStatusPtrOutput {
	return WebhookStatus(e).ToWebhookStatusOutputWithContext(ctx).ToWebhookStatusPtrOutputWithContext(ctx)
}

func (e WebhookStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebhookStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebhookStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebhookStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebhookStatusOutput struct{ *pulumi.OutputState }

func (WebhookStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookStatus)(nil)).Elem()
}

func (o WebhookStatusOutput) ToWebhookStatusOutput() WebhookStatusOutput {
	return o
}

func (o WebhookStatusOutput) ToWebhookStatusOutputWithContext(ctx context.Context) WebhookStatusOutput {
	return o
}

func (o WebhookStatusOutput) ToWebhookStatusPtrOutput() WebhookStatusPtrOutput {
	return o.ToWebhookStatusPtrOutputWithContext(context.Background())
}

func (o WebhookStatusOutput) ToWebhookStatusPtrOutputWithContext(ctx context.Context) WebhookStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebhookStatus) *WebhookStatus {
		return &v
	}).(WebhookStatusPtrOutput)
}

func (o WebhookStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebhookStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebhookStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebhookStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebhookStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebhookStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebhookStatusPtrOutput struct{ *pulumi.OutputState }

func (WebhookStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookStatus)(nil)).Elem()
}

func (o WebhookStatusPtrOutput) ToWebhookStatusPtrOutput() WebhookStatusPtrOutput {
	return o
}

func (o WebhookStatusPtrOutput) ToWebhookStatusPtrOutputWithContext(ctx context.Context) WebhookStatusPtrOutput {
	return o
}

func (o WebhookStatusPtrOutput) Elem() WebhookStatusOutput {
	return o.ApplyT(func(v *WebhookStatus) WebhookStatus {
		if v != nil {
			return *v
		}
		var ret WebhookStatus
		return ret
	}).(WebhookStatusOutput)
}

func (o WebhookStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebhookStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebhookStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebhookStatusInput interface {
	pulumi.Input

	ToWebhookStatusOutput() WebhookStatusOutput
	ToWebhookStatusOutputWithContext(context.Context) WebhookStatusOutput
}

var webhookStatusPtrType = reflect.TypeOf((**WebhookStatus)(nil)).Elem()

type WebhookStatusPtrInput interface {
	pulumi.Input

	ToWebhookStatusPtrOutput() WebhookStatusPtrOutput
	ToWebhookStatusPtrOutputWithContext(context.Context) WebhookStatusPtrOutput
}

type webhookStatusPtr string

func WebhookStatusPtr(v string) WebhookStatusPtrInput {
	return (*webhookStatusPtr)(&v)
}

func (*webhookStatusPtr) ElementType() reflect.Type {
	return webhookStatusPtrType
}

func (in *webhookStatusPtr) ToWebhookStatusPtrOutput() WebhookStatusPtrOutput {
	return pulumi.ToOutput(in).(WebhookStatusPtrOutput)
}

func (in *webhookStatusPtr) ToWebhookStatusPtrOutputWithContext(ctx context.Context) WebhookStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebhookStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(DefaultActionOutput{})
	pulumi.RegisterOutputType(DefaultActionPtrOutput{})
	pulumi.RegisterOutputType(PolicyStatusOutput{})
	pulumi.RegisterOutputType(PolicyStatusPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(TrustPolicyTypeOutput{})
	pulumi.RegisterOutputType(TrustPolicyTypePtrOutput{})
	pulumi.RegisterOutputType(WebhookActionOutput{})
	pulumi.RegisterOutputType(WebhookActionPtrOutput{})
	pulumi.RegisterOutputType(WebhookStatusOutput{})
	pulumi.RegisterOutputType(WebhookStatusPtrOutput{})
}
