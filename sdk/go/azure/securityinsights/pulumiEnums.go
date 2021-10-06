


package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertRuleKind string

const (
	AlertRuleKindScheduled                         = AlertRuleKind("Scheduled")
	AlertRuleKindMicrosoftSecurityIncidentCreation = AlertRuleKind("MicrosoftSecurityIncidentCreation")
	AlertRuleKindFusion                            = AlertRuleKind("Fusion")
)

func (AlertRuleKind) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleKind)(nil)).Elem()
}

func (e AlertRuleKind) ToAlertRuleKindOutput() AlertRuleKindOutput {
	return pulumi.ToOutput(e).(AlertRuleKindOutput)
}

func (e AlertRuleKind) ToAlertRuleKindOutputWithContext(ctx context.Context) AlertRuleKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AlertRuleKindOutput)
}

func (e AlertRuleKind) ToAlertRuleKindPtrOutput() AlertRuleKindPtrOutput {
	return e.ToAlertRuleKindPtrOutputWithContext(context.Background())
}

func (e AlertRuleKind) ToAlertRuleKindPtrOutputWithContext(ctx context.Context) AlertRuleKindPtrOutput {
	return AlertRuleKind(e).ToAlertRuleKindOutputWithContext(ctx).ToAlertRuleKindPtrOutputWithContext(ctx)
}

func (e AlertRuleKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertRuleKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AlertRuleKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AlertRuleKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AlertRuleKindOutput struct{ *pulumi.OutputState }

func (AlertRuleKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleKind)(nil)).Elem()
}

func (o AlertRuleKindOutput) ToAlertRuleKindOutput() AlertRuleKindOutput {
	return o
}

func (o AlertRuleKindOutput) ToAlertRuleKindOutputWithContext(ctx context.Context) AlertRuleKindOutput {
	return o
}

func (o AlertRuleKindOutput) ToAlertRuleKindPtrOutput() AlertRuleKindPtrOutput {
	return o.ToAlertRuleKindPtrOutputWithContext(context.Background())
}

func (o AlertRuleKindOutput) ToAlertRuleKindPtrOutputWithContext(ctx context.Context) AlertRuleKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertRuleKind) *AlertRuleKind {
		return &v
	}).(AlertRuleKindPtrOutput)
}

func (o AlertRuleKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AlertRuleKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlertRuleKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AlertRuleKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlertRuleKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AlertRuleKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AlertRuleKindPtrOutput struct{ *pulumi.OutputState }

func (AlertRuleKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleKind)(nil)).Elem()
}

func (o AlertRuleKindPtrOutput) ToAlertRuleKindPtrOutput() AlertRuleKindPtrOutput {
	return o
}

func (o AlertRuleKindPtrOutput) ToAlertRuleKindPtrOutputWithContext(ctx context.Context) AlertRuleKindPtrOutput {
	return o
}

func (o AlertRuleKindPtrOutput) Elem() AlertRuleKindOutput {
	return o.ApplyT(func(v *AlertRuleKind) AlertRuleKind {
		if v != nil {
			return *v
		}
		var ret AlertRuleKind
		return ret
	}).(AlertRuleKindOutput)
}

func (o AlertRuleKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AlertRuleKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AlertRuleKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AlertRuleKindInput interface {
	pulumi.Input

	ToAlertRuleKindOutput() AlertRuleKindOutput
	ToAlertRuleKindOutputWithContext(context.Context) AlertRuleKindOutput
}

var alertRuleKindPtrType = reflect.TypeOf((**AlertRuleKind)(nil)).Elem()

type AlertRuleKindPtrInput interface {
	pulumi.Input

	ToAlertRuleKindPtrOutput() AlertRuleKindPtrOutput
	ToAlertRuleKindPtrOutputWithContext(context.Context) AlertRuleKindPtrOutput
}

type alertRuleKindPtr string

func AlertRuleKindPtr(v string) AlertRuleKindPtrInput {
	return (*alertRuleKindPtr)(&v)
}

func (*alertRuleKindPtr) ElementType() reflect.Type {
	return alertRuleKindPtrType
}

func (in *alertRuleKindPtr) ToAlertRuleKindPtrOutput() AlertRuleKindPtrOutput {
	return pulumi.ToOutput(in).(AlertRuleKindPtrOutput)
}

func (in *alertRuleKindPtr) ToAlertRuleKindPtrOutputWithContext(ctx context.Context) AlertRuleKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AlertRuleKindPtrOutput)
}

type AlertSeverity string

const (
	// High severity
	AlertSeverityHigh = AlertSeverity("High")
	// Medium severity
	AlertSeverityMedium = AlertSeverity("Medium")
	// Low severity
	AlertSeverityLow = AlertSeverity("Low")
	// Informational severity
	AlertSeverityInformational = AlertSeverity("Informational")
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

type AttackTactic string

const (
	AttackTacticInitialAccess       = AttackTactic("InitialAccess")
	AttackTacticExecution           = AttackTactic("Execution")
	AttackTacticPersistence         = AttackTactic("Persistence")
	AttackTacticPrivilegeEscalation = AttackTactic("PrivilegeEscalation")
	AttackTacticDefenseEvasion      = AttackTactic("DefenseEvasion")
	AttackTacticCredentialAccess    = AttackTactic("CredentialAccess")
	AttackTacticDiscovery           = AttackTactic("Discovery")
	AttackTacticLateralMovement     = AttackTactic("LateralMovement")
	AttackTacticCollection          = AttackTactic("Collection")
	AttackTacticExfiltration        = AttackTactic("Exfiltration")
	AttackTacticCommandAndControl   = AttackTactic("CommandAndControl")
	AttackTacticImpact              = AttackTactic("Impact")
)

func (AttackTactic) ElementType() reflect.Type {
	return reflect.TypeOf((*AttackTactic)(nil)).Elem()
}

func (e AttackTactic) ToAttackTacticOutput() AttackTacticOutput {
	return pulumi.ToOutput(e).(AttackTacticOutput)
}

func (e AttackTactic) ToAttackTacticOutputWithContext(ctx context.Context) AttackTacticOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AttackTacticOutput)
}

func (e AttackTactic) ToAttackTacticPtrOutput() AttackTacticPtrOutput {
	return e.ToAttackTacticPtrOutputWithContext(context.Background())
}

func (e AttackTactic) ToAttackTacticPtrOutputWithContext(ctx context.Context) AttackTacticPtrOutput {
	return AttackTactic(e).ToAttackTacticOutputWithContext(ctx).ToAttackTacticPtrOutputWithContext(ctx)
}

func (e AttackTactic) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AttackTactic) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AttackTactic) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AttackTactic) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AttackTacticOutput struct{ *pulumi.OutputState }

func (AttackTacticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttackTactic)(nil)).Elem()
}

func (o AttackTacticOutput) ToAttackTacticOutput() AttackTacticOutput {
	return o
}

func (o AttackTacticOutput) ToAttackTacticOutputWithContext(ctx context.Context) AttackTacticOutput {
	return o
}

func (o AttackTacticOutput) ToAttackTacticPtrOutput() AttackTacticPtrOutput {
	return o.ToAttackTacticPtrOutputWithContext(context.Background())
}

func (o AttackTacticOutput) ToAttackTacticPtrOutputWithContext(ctx context.Context) AttackTacticPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AttackTactic) *AttackTactic {
		return &v
	}).(AttackTacticPtrOutput)
}

func (o AttackTacticOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AttackTacticOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AttackTactic) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AttackTacticOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AttackTacticOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AttackTactic) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AttackTacticPtrOutput struct{ *pulumi.OutputState }

func (AttackTacticPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttackTactic)(nil)).Elem()
}

func (o AttackTacticPtrOutput) ToAttackTacticPtrOutput() AttackTacticPtrOutput {
	return o
}

func (o AttackTacticPtrOutput) ToAttackTacticPtrOutputWithContext(ctx context.Context) AttackTacticPtrOutput {
	return o
}

func (o AttackTacticPtrOutput) Elem() AttackTacticOutput {
	return o.ApplyT(func(v *AttackTactic) AttackTactic {
		if v != nil {
			return *v
		}
		var ret AttackTactic
		return ret
	}).(AttackTacticOutput)
}

func (o AttackTacticPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AttackTacticPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AttackTactic) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AttackTacticInput interface {
	pulumi.Input

	ToAttackTacticOutput() AttackTacticOutput
	ToAttackTacticOutputWithContext(context.Context) AttackTacticOutput
}

var attackTacticPtrType = reflect.TypeOf((**AttackTactic)(nil)).Elem()

type AttackTacticPtrInput interface {
	pulumi.Input

	ToAttackTacticPtrOutput() AttackTacticPtrOutput
	ToAttackTacticPtrOutputWithContext(context.Context) AttackTacticPtrOutput
}

type attackTacticPtr string

func AttackTacticPtr(v string) AttackTacticPtrInput {
	return (*attackTacticPtr)(&v)
}

func (*attackTacticPtr) ElementType() reflect.Type {
	return attackTacticPtrType
}

func (in *attackTacticPtr) ToAttackTacticPtrOutput() AttackTacticPtrOutput {
	return pulumi.ToOutput(in).(AttackTacticPtrOutput)
}

func (in *attackTacticPtr) ToAttackTacticPtrOutputWithContext(ctx context.Context) AttackTacticPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AttackTacticPtrOutput)
}

type AutomationRuleActionType string

const (
	// Modify an object's properties
	AutomationRuleActionTypeModifyProperties = AutomationRuleActionType("ModifyProperties")
	// Run a playbook on an object
	AutomationRuleActionTypeRunPlaybook = AutomationRuleActionType("RunPlaybook")
)

func (AutomationRuleActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleActionType)(nil)).Elem()
}

func (e AutomationRuleActionType) ToAutomationRuleActionTypeOutput() AutomationRuleActionTypeOutput {
	return pulumi.ToOutput(e).(AutomationRuleActionTypeOutput)
}

func (e AutomationRuleActionType) ToAutomationRuleActionTypeOutputWithContext(ctx context.Context) AutomationRuleActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutomationRuleActionTypeOutput)
}

func (e AutomationRuleActionType) ToAutomationRuleActionTypePtrOutput() AutomationRuleActionTypePtrOutput {
	return e.ToAutomationRuleActionTypePtrOutputWithContext(context.Background())
}

func (e AutomationRuleActionType) ToAutomationRuleActionTypePtrOutputWithContext(ctx context.Context) AutomationRuleActionTypePtrOutput {
	return AutomationRuleActionType(e).ToAutomationRuleActionTypeOutputWithContext(ctx).ToAutomationRuleActionTypePtrOutputWithContext(ctx)
}

func (e AutomationRuleActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRuleActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRuleActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutomationRuleActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutomationRuleActionTypeOutput struct{ *pulumi.OutputState }

func (AutomationRuleActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleActionType)(nil)).Elem()
}

func (o AutomationRuleActionTypeOutput) ToAutomationRuleActionTypeOutput() AutomationRuleActionTypeOutput {
	return o
}

func (o AutomationRuleActionTypeOutput) ToAutomationRuleActionTypeOutputWithContext(ctx context.Context) AutomationRuleActionTypeOutput {
	return o
}

func (o AutomationRuleActionTypeOutput) ToAutomationRuleActionTypePtrOutput() AutomationRuleActionTypePtrOutput {
	return o.ToAutomationRuleActionTypePtrOutputWithContext(context.Background())
}

func (o AutomationRuleActionTypeOutput) ToAutomationRuleActionTypePtrOutputWithContext(ctx context.Context) AutomationRuleActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomationRuleActionType) *AutomationRuleActionType {
		return &v
	}).(AutomationRuleActionTypePtrOutput)
}

func (o AutomationRuleActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutomationRuleActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationRuleActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutomationRuleActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationRuleActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationRuleActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutomationRuleActionTypePtrOutput struct{ *pulumi.OutputState }

func (AutomationRuleActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRuleActionType)(nil)).Elem()
}

func (o AutomationRuleActionTypePtrOutput) ToAutomationRuleActionTypePtrOutput() AutomationRuleActionTypePtrOutput {
	return o
}

func (o AutomationRuleActionTypePtrOutput) ToAutomationRuleActionTypePtrOutputWithContext(ctx context.Context) AutomationRuleActionTypePtrOutput {
	return o
}

func (o AutomationRuleActionTypePtrOutput) Elem() AutomationRuleActionTypeOutput {
	return o.ApplyT(func(v *AutomationRuleActionType) AutomationRuleActionType {
		if v != nil {
			return *v
		}
		var ret AutomationRuleActionType
		return ret
	}).(AutomationRuleActionTypeOutput)
}

func (o AutomationRuleActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationRuleActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutomationRuleActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutomationRuleActionTypeInput interface {
	pulumi.Input

	ToAutomationRuleActionTypeOutput() AutomationRuleActionTypeOutput
	ToAutomationRuleActionTypeOutputWithContext(context.Context) AutomationRuleActionTypeOutput
}

var automationRuleActionTypePtrType = reflect.TypeOf((**AutomationRuleActionType)(nil)).Elem()

type AutomationRuleActionTypePtrInput interface {
	pulumi.Input

	ToAutomationRuleActionTypePtrOutput() AutomationRuleActionTypePtrOutput
	ToAutomationRuleActionTypePtrOutputWithContext(context.Context) AutomationRuleActionTypePtrOutput
}

type automationRuleActionTypePtr string

func AutomationRuleActionTypePtr(v string) AutomationRuleActionTypePtrInput {
	return (*automationRuleActionTypePtr)(&v)
}

func (*automationRuleActionTypePtr) ElementType() reflect.Type {
	return automationRuleActionTypePtrType
}

func (in *automationRuleActionTypePtr) ToAutomationRuleActionTypePtrOutput() AutomationRuleActionTypePtrOutput {
	return pulumi.ToOutput(in).(AutomationRuleActionTypePtrOutput)
}

func (in *automationRuleActionTypePtr) ToAutomationRuleActionTypePtrOutputWithContext(ctx context.Context) AutomationRuleActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutomationRuleActionTypePtrOutput)
}

type AutomationRuleConditionType string

const (
	// Evaluate an object property value
	AutomationRuleConditionTypeProperty = AutomationRuleConditionType("Property")
)

func (AutomationRuleConditionType) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleConditionType)(nil)).Elem()
}

func (e AutomationRuleConditionType) ToAutomationRuleConditionTypeOutput() AutomationRuleConditionTypeOutput {
	return pulumi.ToOutput(e).(AutomationRuleConditionTypeOutput)
}

func (e AutomationRuleConditionType) ToAutomationRuleConditionTypeOutputWithContext(ctx context.Context) AutomationRuleConditionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutomationRuleConditionTypeOutput)
}

func (e AutomationRuleConditionType) ToAutomationRuleConditionTypePtrOutput() AutomationRuleConditionTypePtrOutput {
	return e.ToAutomationRuleConditionTypePtrOutputWithContext(context.Background())
}

func (e AutomationRuleConditionType) ToAutomationRuleConditionTypePtrOutputWithContext(ctx context.Context) AutomationRuleConditionTypePtrOutput {
	return AutomationRuleConditionType(e).ToAutomationRuleConditionTypeOutputWithContext(ctx).ToAutomationRuleConditionTypePtrOutputWithContext(ctx)
}

func (e AutomationRuleConditionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRuleConditionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRuleConditionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutomationRuleConditionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutomationRuleConditionTypeOutput struct{ *pulumi.OutputState }

func (AutomationRuleConditionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleConditionType)(nil)).Elem()
}

func (o AutomationRuleConditionTypeOutput) ToAutomationRuleConditionTypeOutput() AutomationRuleConditionTypeOutput {
	return o
}

func (o AutomationRuleConditionTypeOutput) ToAutomationRuleConditionTypeOutputWithContext(ctx context.Context) AutomationRuleConditionTypeOutput {
	return o
}

func (o AutomationRuleConditionTypeOutput) ToAutomationRuleConditionTypePtrOutput() AutomationRuleConditionTypePtrOutput {
	return o.ToAutomationRuleConditionTypePtrOutputWithContext(context.Background())
}

func (o AutomationRuleConditionTypeOutput) ToAutomationRuleConditionTypePtrOutputWithContext(ctx context.Context) AutomationRuleConditionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomationRuleConditionType) *AutomationRuleConditionType {
		return &v
	}).(AutomationRuleConditionTypePtrOutput)
}

func (o AutomationRuleConditionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutomationRuleConditionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationRuleConditionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutomationRuleConditionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationRuleConditionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationRuleConditionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutomationRuleConditionTypePtrOutput struct{ *pulumi.OutputState }

func (AutomationRuleConditionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRuleConditionType)(nil)).Elem()
}

func (o AutomationRuleConditionTypePtrOutput) ToAutomationRuleConditionTypePtrOutput() AutomationRuleConditionTypePtrOutput {
	return o
}

func (o AutomationRuleConditionTypePtrOutput) ToAutomationRuleConditionTypePtrOutputWithContext(ctx context.Context) AutomationRuleConditionTypePtrOutput {
	return o
}

func (o AutomationRuleConditionTypePtrOutput) Elem() AutomationRuleConditionTypeOutput {
	return o.ApplyT(func(v *AutomationRuleConditionType) AutomationRuleConditionType {
		if v != nil {
			return *v
		}
		var ret AutomationRuleConditionType
		return ret
	}).(AutomationRuleConditionTypeOutput)
}

func (o AutomationRuleConditionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationRuleConditionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutomationRuleConditionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutomationRuleConditionTypeInput interface {
	pulumi.Input

	ToAutomationRuleConditionTypeOutput() AutomationRuleConditionTypeOutput
	ToAutomationRuleConditionTypeOutputWithContext(context.Context) AutomationRuleConditionTypeOutput
}

var automationRuleConditionTypePtrType = reflect.TypeOf((**AutomationRuleConditionType)(nil)).Elem()

type AutomationRuleConditionTypePtrInput interface {
	pulumi.Input

	ToAutomationRuleConditionTypePtrOutput() AutomationRuleConditionTypePtrOutput
	ToAutomationRuleConditionTypePtrOutputWithContext(context.Context) AutomationRuleConditionTypePtrOutput
}

type automationRuleConditionTypePtr string

func AutomationRuleConditionTypePtr(v string) AutomationRuleConditionTypePtrInput {
	return (*automationRuleConditionTypePtr)(&v)
}

func (*automationRuleConditionTypePtr) ElementType() reflect.Type {
	return automationRuleConditionTypePtrType
}

func (in *automationRuleConditionTypePtr) ToAutomationRuleConditionTypePtrOutput() AutomationRuleConditionTypePtrOutput {
	return pulumi.ToOutput(in).(AutomationRuleConditionTypePtrOutput)
}

func (in *automationRuleConditionTypePtr) ToAutomationRuleConditionTypePtrOutputWithContext(ctx context.Context) AutomationRuleConditionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutomationRuleConditionTypePtrOutput)
}

type AutomationRulePropertyConditionSupportedOperator string

const (
	// Evaluates if the property equals at least one of the condition values
	AutomationRulePropertyConditionSupportedOperatorEquals = AutomationRulePropertyConditionSupportedOperator("Equals")
	// Evaluates if the property does not equal any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotEquals = AutomationRulePropertyConditionSupportedOperator("NotEquals")
	// Evaluates if the property contains at least one of the condition values
	AutomationRulePropertyConditionSupportedOperatorContains = AutomationRulePropertyConditionSupportedOperator("Contains")
	// Evaluates if the property does not contain any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotContains = AutomationRulePropertyConditionSupportedOperator("NotContains")
	// Evaluates if the property starts with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorStartsWith = AutomationRulePropertyConditionSupportedOperator("StartsWith")
	// Evaluates if the property does not start with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotStartsWith = AutomationRulePropertyConditionSupportedOperator("NotStartsWith")
	// Evaluates if the property ends with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorEndsWith = AutomationRulePropertyConditionSupportedOperator("EndsWith")
	// Evaluates if the property does not end with any of the condition values
	AutomationRulePropertyConditionSupportedOperatorNotEndsWith = AutomationRulePropertyConditionSupportedOperator("NotEndsWith")
)

func (AutomationRulePropertyConditionSupportedOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyConditionSupportedOperator)(nil)).Elem()
}

func (e AutomationRulePropertyConditionSupportedOperator) ToAutomationRulePropertyConditionSupportedOperatorOutput() AutomationRulePropertyConditionSupportedOperatorOutput {
	return pulumi.ToOutput(e).(AutomationRulePropertyConditionSupportedOperatorOutput)
}

func (e AutomationRulePropertyConditionSupportedOperator) ToAutomationRulePropertyConditionSupportedOperatorOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutomationRulePropertyConditionSupportedOperatorOutput)
}

func (e AutomationRulePropertyConditionSupportedOperator) ToAutomationRulePropertyConditionSupportedOperatorPtrOutput() AutomationRulePropertyConditionSupportedOperatorPtrOutput {
	return e.ToAutomationRulePropertyConditionSupportedOperatorPtrOutputWithContext(context.Background())
}

func (e AutomationRulePropertyConditionSupportedOperator) ToAutomationRulePropertyConditionSupportedOperatorPtrOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedOperatorPtrOutput {
	return AutomationRulePropertyConditionSupportedOperator(e).ToAutomationRulePropertyConditionSupportedOperatorOutputWithContext(ctx).ToAutomationRulePropertyConditionSupportedOperatorPtrOutputWithContext(ctx)
}

func (e AutomationRulePropertyConditionSupportedOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRulePropertyConditionSupportedOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRulePropertyConditionSupportedOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutomationRulePropertyConditionSupportedOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutomationRulePropertyConditionSupportedOperatorOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyConditionSupportedOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyConditionSupportedOperator)(nil)).Elem()
}

func (o AutomationRulePropertyConditionSupportedOperatorOutput) ToAutomationRulePropertyConditionSupportedOperatorOutput() AutomationRulePropertyConditionSupportedOperatorOutput {
	return o
}

func (o AutomationRulePropertyConditionSupportedOperatorOutput) ToAutomationRulePropertyConditionSupportedOperatorOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedOperatorOutput {
	return o
}

func (o AutomationRulePropertyConditionSupportedOperatorOutput) ToAutomationRulePropertyConditionSupportedOperatorPtrOutput() AutomationRulePropertyConditionSupportedOperatorPtrOutput {
	return o.ToAutomationRulePropertyConditionSupportedOperatorPtrOutputWithContext(context.Background())
}

func (o AutomationRulePropertyConditionSupportedOperatorOutput) ToAutomationRulePropertyConditionSupportedOperatorPtrOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomationRulePropertyConditionSupportedOperator) *AutomationRulePropertyConditionSupportedOperator {
		return &v
	}).(AutomationRulePropertyConditionSupportedOperatorPtrOutput)
}

func (o AutomationRulePropertyConditionSupportedOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutomationRulePropertyConditionSupportedOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationRulePropertyConditionSupportedOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutomationRulePropertyConditionSupportedOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationRulePropertyConditionSupportedOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationRulePropertyConditionSupportedOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutomationRulePropertyConditionSupportedOperatorPtrOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyConditionSupportedOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRulePropertyConditionSupportedOperator)(nil)).Elem()
}

func (o AutomationRulePropertyConditionSupportedOperatorPtrOutput) ToAutomationRulePropertyConditionSupportedOperatorPtrOutput() AutomationRulePropertyConditionSupportedOperatorPtrOutput {
	return o
}

func (o AutomationRulePropertyConditionSupportedOperatorPtrOutput) ToAutomationRulePropertyConditionSupportedOperatorPtrOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedOperatorPtrOutput {
	return o
}

func (o AutomationRulePropertyConditionSupportedOperatorPtrOutput) Elem() AutomationRulePropertyConditionSupportedOperatorOutput {
	return o.ApplyT(func(v *AutomationRulePropertyConditionSupportedOperator) AutomationRulePropertyConditionSupportedOperator {
		if v != nil {
			return *v
		}
		var ret AutomationRulePropertyConditionSupportedOperator
		return ret
	}).(AutomationRulePropertyConditionSupportedOperatorOutput)
}

func (o AutomationRulePropertyConditionSupportedOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationRulePropertyConditionSupportedOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutomationRulePropertyConditionSupportedOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutomationRulePropertyConditionSupportedOperatorInput interface {
	pulumi.Input

	ToAutomationRulePropertyConditionSupportedOperatorOutput() AutomationRulePropertyConditionSupportedOperatorOutput
	ToAutomationRulePropertyConditionSupportedOperatorOutputWithContext(context.Context) AutomationRulePropertyConditionSupportedOperatorOutput
}

var automationRulePropertyConditionSupportedOperatorPtrType = reflect.TypeOf((**AutomationRulePropertyConditionSupportedOperator)(nil)).Elem()

type AutomationRulePropertyConditionSupportedOperatorPtrInput interface {
	pulumi.Input

	ToAutomationRulePropertyConditionSupportedOperatorPtrOutput() AutomationRulePropertyConditionSupportedOperatorPtrOutput
	ToAutomationRulePropertyConditionSupportedOperatorPtrOutputWithContext(context.Context) AutomationRulePropertyConditionSupportedOperatorPtrOutput
}

type automationRulePropertyConditionSupportedOperatorPtr string

func AutomationRulePropertyConditionSupportedOperatorPtr(v string) AutomationRulePropertyConditionSupportedOperatorPtrInput {
	return (*automationRulePropertyConditionSupportedOperatorPtr)(&v)
}

func (*automationRulePropertyConditionSupportedOperatorPtr) ElementType() reflect.Type {
	return automationRulePropertyConditionSupportedOperatorPtrType
}

func (in *automationRulePropertyConditionSupportedOperatorPtr) ToAutomationRulePropertyConditionSupportedOperatorPtrOutput() AutomationRulePropertyConditionSupportedOperatorPtrOutput {
	return pulumi.ToOutput(in).(AutomationRulePropertyConditionSupportedOperatorPtrOutput)
}

func (in *automationRulePropertyConditionSupportedOperatorPtr) ToAutomationRulePropertyConditionSupportedOperatorPtrOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutomationRulePropertyConditionSupportedOperatorPtrOutput)
}

type AutomationRulePropertyConditionSupportedProperty string

const (
	// The title of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentTitle = AutomationRulePropertyConditionSupportedProperty("IncidentTitle")
	// The description of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentDescription = AutomationRulePropertyConditionSupportedProperty("IncidentDescription")
	// The severity of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentSeverity = AutomationRulePropertyConditionSupportedProperty("IncidentSeverity")
	// The status of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentStatus = AutomationRulePropertyConditionSupportedProperty("IncidentStatus")
	// The tactics of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentTactics = AutomationRulePropertyConditionSupportedProperty("IncidentTactics")
	// The related Analytic rule ids of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentRelatedAnalyticRuleIds = AutomationRulePropertyConditionSupportedProperty("IncidentRelatedAnalyticRuleIds")
	// The provider name of the incident
	AutomationRulePropertyConditionSupportedPropertyIncidentProviderName = AutomationRulePropertyConditionSupportedProperty("IncidentProviderName")
	// The account Azure Active Directory tenant id
	AutomationRulePropertyConditionSupportedPropertyAccountAadTenantId = AutomationRulePropertyConditionSupportedProperty("AccountAadTenantId")
	// The account Azure Active Directory user id.
	AutomationRulePropertyConditionSupportedPropertyAccountAadUserId = AutomationRulePropertyConditionSupportedProperty("AccountAadUserId")
	// The account name
	AutomationRulePropertyConditionSupportedPropertyAccountName = AutomationRulePropertyConditionSupportedProperty("AccountName")
	// The account NetBIOS domain name
	AutomationRulePropertyConditionSupportedPropertyAccountNTDomain = AutomationRulePropertyConditionSupportedProperty("AccountNTDomain")
	// The account Azure Active Directory Passport User ID
	AutomationRulePropertyConditionSupportedPropertyAccountPUID = AutomationRulePropertyConditionSupportedProperty("AccountPUID")
	// The account security identifier
	AutomationRulePropertyConditionSupportedPropertyAccountSid = AutomationRulePropertyConditionSupportedProperty("AccountSid")
	// The account unique identifier
	AutomationRulePropertyConditionSupportedPropertyAccountObjectGuid = AutomationRulePropertyConditionSupportedProperty("AccountObjectGuid")
	// The account user principal name suffix
	AutomationRulePropertyConditionSupportedPropertyAccountUPNSuffix = AutomationRulePropertyConditionSupportedProperty("AccountUPNSuffix")
	// The Azure resource id
	AutomationRulePropertyConditionSupportedPropertyAzureResourceResourceId = AutomationRulePropertyConditionSupportedProperty("AzureResourceResourceId")
	// The Azure resource subscription id
	AutomationRulePropertyConditionSupportedPropertyAzureResourceSubscriptionId = AutomationRulePropertyConditionSupportedProperty("AzureResourceSubscriptionId")
	// The cloud application identifier
	AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppId = AutomationRulePropertyConditionSupportedProperty("CloudApplicationAppId")
	// The cloud application name
	AutomationRulePropertyConditionSupportedPropertyCloudApplicationAppName = AutomationRulePropertyConditionSupportedProperty("CloudApplicationAppName")
	// The dns record domain name
	AutomationRulePropertyConditionSupportedPropertyDNSDomainName = AutomationRulePropertyConditionSupportedProperty("DNSDomainName")
	// The file directory full path
	AutomationRulePropertyConditionSupportedPropertyFileDirectory = AutomationRulePropertyConditionSupportedProperty("FileDirectory")
	// The file name without path
	AutomationRulePropertyConditionSupportedPropertyFileName = AutomationRulePropertyConditionSupportedProperty("FileName")
	// The file hash value
	AutomationRulePropertyConditionSupportedPropertyFileHashValue = AutomationRulePropertyConditionSupportedProperty("FileHashValue")
	// The host Azure resource id
	AutomationRulePropertyConditionSupportedPropertyHostAzureID = AutomationRulePropertyConditionSupportedProperty("HostAzureID")
	// The host name without domain
	AutomationRulePropertyConditionSupportedPropertyHostName = AutomationRulePropertyConditionSupportedProperty("HostName")
	// The host NetBIOS name
	AutomationRulePropertyConditionSupportedPropertyHostNetBiosName = AutomationRulePropertyConditionSupportedProperty("HostNetBiosName")
	// The host NT domain
	AutomationRulePropertyConditionSupportedPropertyHostNTDomain = AutomationRulePropertyConditionSupportedProperty("HostNTDomain")
	// The host operating system
	AutomationRulePropertyConditionSupportedPropertyHostOSVersion = AutomationRulePropertyConditionSupportedProperty("HostOSVersion")
	// The IoT device id
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceId = AutomationRulePropertyConditionSupportedProperty("IoTDeviceId")
	// The IoT device name
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceName = AutomationRulePropertyConditionSupportedProperty("IoTDeviceName")
	// The IoT device type
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceType = AutomationRulePropertyConditionSupportedProperty("IoTDeviceType")
	// The IoT device vendor
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceVendor = AutomationRulePropertyConditionSupportedProperty("IoTDeviceVendor")
	// The IoT device model
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceModel = AutomationRulePropertyConditionSupportedProperty("IoTDeviceModel")
	// The IoT device operating system
	AutomationRulePropertyConditionSupportedPropertyIoTDeviceOperatingSystem = AutomationRulePropertyConditionSupportedProperty("IoTDeviceOperatingSystem")
	// The IP address
	AutomationRulePropertyConditionSupportedPropertyIPAddress = AutomationRulePropertyConditionSupportedProperty("IPAddress")
	// The mailbox display name
	AutomationRulePropertyConditionSupportedPropertyMailboxDisplayName = AutomationRulePropertyConditionSupportedProperty("MailboxDisplayName")
	// The mailbox primary address
	AutomationRulePropertyConditionSupportedPropertyMailboxPrimaryAddress = AutomationRulePropertyConditionSupportedProperty("MailboxPrimaryAddress")
	// The mailbox user principal name
	AutomationRulePropertyConditionSupportedPropertyMailboxUPN = AutomationRulePropertyConditionSupportedProperty("MailboxUPN")
	// The mail message delivery action
	AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryAction = AutomationRulePropertyConditionSupportedProperty("MailMessageDeliveryAction")
	// The mail message delivery location
	AutomationRulePropertyConditionSupportedPropertyMailMessageDeliveryLocation = AutomationRulePropertyConditionSupportedProperty("MailMessageDeliveryLocation")
	// The mail message recipient
	AutomationRulePropertyConditionSupportedPropertyMailMessageRecipient = AutomationRulePropertyConditionSupportedProperty("MailMessageRecipient")
	// The mail message sender IP address
	AutomationRulePropertyConditionSupportedPropertyMailMessageSenderIP = AutomationRulePropertyConditionSupportedProperty("MailMessageSenderIP")
	// The mail message subject
	AutomationRulePropertyConditionSupportedPropertyMailMessageSubject = AutomationRulePropertyConditionSupportedProperty("MailMessageSubject")
	// The mail message P1 sender
	AutomationRulePropertyConditionSupportedPropertyMailMessageP1Sender = AutomationRulePropertyConditionSupportedProperty("MailMessageP1Sender")
	// The mail message P2 sender
	AutomationRulePropertyConditionSupportedPropertyMailMessageP2Sender = AutomationRulePropertyConditionSupportedProperty("MailMessageP2Sender")
	// The malware category
	AutomationRulePropertyConditionSupportedPropertyMalwareCategory = AutomationRulePropertyConditionSupportedProperty("MalwareCategory")
	// The malware name
	AutomationRulePropertyConditionSupportedPropertyMalwareName = AutomationRulePropertyConditionSupportedProperty("MalwareName")
	// The process execution command line
	AutomationRulePropertyConditionSupportedPropertyProcessCommandLine = AutomationRulePropertyConditionSupportedProperty("ProcessCommandLine")
	// The process id
	AutomationRulePropertyConditionSupportedPropertyProcessId = AutomationRulePropertyConditionSupportedProperty("ProcessId")
	// The registry key path
	AutomationRulePropertyConditionSupportedPropertyRegistryKey = AutomationRulePropertyConditionSupportedProperty("RegistryKey")
	// The registry key value in string formatted representation
	AutomationRulePropertyConditionSupportedPropertyRegistryValueData = AutomationRulePropertyConditionSupportedProperty("RegistryValueData")
	// The url
	AutomationRulePropertyConditionSupportedPropertyUrl = AutomationRulePropertyConditionSupportedProperty("Url")
)

func (AutomationRulePropertyConditionSupportedProperty) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyConditionSupportedProperty)(nil)).Elem()
}

func (e AutomationRulePropertyConditionSupportedProperty) ToAutomationRulePropertyConditionSupportedPropertyOutput() AutomationRulePropertyConditionSupportedPropertyOutput {
	return pulumi.ToOutput(e).(AutomationRulePropertyConditionSupportedPropertyOutput)
}

func (e AutomationRulePropertyConditionSupportedProperty) ToAutomationRulePropertyConditionSupportedPropertyOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutomationRulePropertyConditionSupportedPropertyOutput)
}

func (e AutomationRulePropertyConditionSupportedProperty) ToAutomationRulePropertyConditionSupportedPropertyPtrOutput() AutomationRulePropertyConditionSupportedPropertyPtrOutput {
	return e.ToAutomationRulePropertyConditionSupportedPropertyPtrOutputWithContext(context.Background())
}

func (e AutomationRulePropertyConditionSupportedProperty) ToAutomationRulePropertyConditionSupportedPropertyPtrOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedPropertyPtrOutput {
	return AutomationRulePropertyConditionSupportedProperty(e).ToAutomationRulePropertyConditionSupportedPropertyOutputWithContext(ctx).ToAutomationRulePropertyConditionSupportedPropertyPtrOutputWithContext(ctx)
}

func (e AutomationRulePropertyConditionSupportedProperty) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRulePropertyConditionSupportedProperty) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutomationRulePropertyConditionSupportedProperty) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutomationRulePropertyConditionSupportedProperty) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutomationRulePropertyConditionSupportedPropertyOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyConditionSupportedPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyConditionSupportedProperty)(nil)).Elem()
}

func (o AutomationRulePropertyConditionSupportedPropertyOutput) ToAutomationRulePropertyConditionSupportedPropertyOutput() AutomationRulePropertyConditionSupportedPropertyOutput {
	return o
}

func (o AutomationRulePropertyConditionSupportedPropertyOutput) ToAutomationRulePropertyConditionSupportedPropertyOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedPropertyOutput {
	return o
}

func (o AutomationRulePropertyConditionSupportedPropertyOutput) ToAutomationRulePropertyConditionSupportedPropertyPtrOutput() AutomationRulePropertyConditionSupportedPropertyPtrOutput {
	return o.ToAutomationRulePropertyConditionSupportedPropertyPtrOutputWithContext(context.Background())
}

func (o AutomationRulePropertyConditionSupportedPropertyOutput) ToAutomationRulePropertyConditionSupportedPropertyPtrOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomationRulePropertyConditionSupportedProperty) *AutomationRulePropertyConditionSupportedProperty {
		return &v
	}).(AutomationRulePropertyConditionSupportedPropertyPtrOutput)
}

func (o AutomationRulePropertyConditionSupportedPropertyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutomationRulePropertyConditionSupportedPropertyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationRulePropertyConditionSupportedProperty) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutomationRulePropertyConditionSupportedPropertyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationRulePropertyConditionSupportedPropertyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutomationRulePropertyConditionSupportedProperty) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutomationRulePropertyConditionSupportedPropertyPtrOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyConditionSupportedPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRulePropertyConditionSupportedProperty)(nil)).Elem()
}

func (o AutomationRulePropertyConditionSupportedPropertyPtrOutput) ToAutomationRulePropertyConditionSupportedPropertyPtrOutput() AutomationRulePropertyConditionSupportedPropertyPtrOutput {
	return o
}

func (o AutomationRulePropertyConditionSupportedPropertyPtrOutput) ToAutomationRulePropertyConditionSupportedPropertyPtrOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedPropertyPtrOutput {
	return o
}

func (o AutomationRulePropertyConditionSupportedPropertyPtrOutput) Elem() AutomationRulePropertyConditionSupportedPropertyOutput {
	return o.ApplyT(func(v *AutomationRulePropertyConditionSupportedProperty) AutomationRulePropertyConditionSupportedProperty {
		if v != nil {
			return *v
		}
		var ret AutomationRulePropertyConditionSupportedProperty
		return ret
	}).(AutomationRulePropertyConditionSupportedPropertyOutput)
}

func (o AutomationRulePropertyConditionSupportedPropertyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutomationRulePropertyConditionSupportedPropertyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutomationRulePropertyConditionSupportedProperty) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutomationRulePropertyConditionSupportedPropertyInput interface {
	pulumi.Input

	ToAutomationRulePropertyConditionSupportedPropertyOutput() AutomationRulePropertyConditionSupportedPropertyOutput
	ToAutomationRulePropertyConditionSupportedPropertyOutputWithContext(context.Context) AutomationRulePropertyConditionSupportedPropertyOutput
}

var automationRulePropertyConditionSupportedPropertyPtrType = reflect.TypeOf((**AutomationRulePropertyConditionSupportedProperty)(nil)).Elem()

type AutomationRulePropertyConditionSupportedPropertyPtrInput interface {
	pulumi.Input

	ToAutomationRulePropertyConditionSupportedPropertyPtrOutput() AutomationRulePropertyConditionSupportedPropertyPtrOutput
	ToAutomationRulePropertyConditionSupportedPropertyPtrOutputWithContext(context.Context) AutomationRulePropertyConditionSupportedPropertyPtrOutput
}

type automationRulePropertyConditionSupportedPropertyPtr string

func AutomationRulePropertyConditionSupportedPropertyPtr(v string) AutomationRulePropertyConditionSupportedPropertyPtrInput {
	return (*automationRulePropertyConditionSupportedPropertyPtr)(&v)
}

func (*automationRulePropertyConditionSupportedPropertyPtr) ElementType() reflect.Type {
	return automationRulePropertyConditionSupportedPropertyPtrType
}

func (in *automationRulePropertyConditionSupportedPropertyPtr) ToAutomationRulePropertyConditionSupportedPropertyPtrOutput() AutomationRulePropertyConditionSupportedPropertyPtrOutput {
	return pulumi.ToOutput(in).(AutomationRulePropertyConditionSupportedPropertyPtrOutput)
}

func (in *automationRulePropertyConditionSupportedPropertyPtr) ToAutomationRulePropertyConditionSupportedPropertyPtrOutputWithContext(ctx context.Context) AutomationRulePropertyConditionSupportedPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutomationRulePropertyConditionSupportedPropertyPtrOutput)
}

type ContentType string

const (
	ContentTypeAnalyticRule = ContentType("AnalyticRule")
	ContentTypeWorkbook     = ContentType("Workbook")
)

func (ContentType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentType)(nil)).Elem()
}

func (e ContentType) ToContentTypeOutput() ContentTypeOutput {
	return pulumi.ToOutput(e).(ContentTypeOutput)
}

func (e ContentType) ToContentTypeOutputWithContext(ctx context.Context) ContentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContentTypeOutput)
}

func (e ContentType) ToContentTypePtrOutput() ContentTypePtrOutput {
	return e.ToContentTypePtrOutputWithContext(context.Background())
}

func (e ContentType) ToContentTypePtrOutputWithContext(ctx context.Context) ContentTypePtrOutput {
	return ContentType(e).ToContentTypeOutputWithContext(ctx).ToContentTypePtrOutputWithContext(ctx)
}

func (e ContentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContentTypeOutput struct{ *pulumi.OutputState }

func (ContentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentType)(nil)).Elem()
}

func (o ContentTypeOutput) ToContentTypeOutput() ContentTypeOutput {
	return o
}

func (o ContentTypeOutput) ToContentTypeOutputWithContext(ctx context.Context) ContentTypeOutput {
	return o
}

func (o ContentTypeOutput) ToContentTypePtrOutput() ContentTypePtrOutput {
	return o.ToContentTypePtrOutputWithContext(context.Background())
}

func (o ContentTypeOutput) ToContentTypePtrOutputWithContext(ctx context.Context) ContentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentType) *ContentType {
		return &v
	}).(ContentTypePtrOutput)
}

func (o ContentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContentTypePtrOutput struct{ *pulumi.OutputState }

func (ContentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentType)(nil)).Elem()
}

func (o ContentTypePtrOutput) ToContentTypePtrOutput() ContentTypePtrOutput {
	return o
}

func (o ContentTypePtrOutput) ToContentTypePtrOutputWithContext(ctx context.Context) ContentTypePtrOutput {
	return o
}

func (o ContentTypePtrOutput) Elem() ContentTypeOutput {
	return o.ApplyT(func(v *ContentType) ContentType {
		if v != nil {
			return *v
		}
		var ret ContentType
		return ret
	}).(ContentTypeOutput)
}

func (o ContentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContentTypeInput interface {
	pulumi.Input

	ToContentTypeOutput() ContentTypeOutput
	ToContentTypeOutputWithContext(context.Context) ContentTypeOutput
}

var contentTypePtrType = reflect.TypeOf((**ContentType)(nil)).Elem()

type ContentTypePtrInput interface {
	pulumi.Input

	ToContentTypePtrOutput() ContentTypePtrOutput
	ToContentTypePtrOutputWithContext(context.Context) ContentTypePtrOutput
}

type contentTypePtr string

func ContentTypePtr(v string) ContentTypePtrInput {
	return (*contentTypePtr)(&v)
}

func (*contentTypePtr) ElementType() reflect.Type {
	return contentTypePtrType
}

func (in *contentTypePtr) ToContentTypePtrOutput() ContentTypePtrOutput {
	return pulumi.ToOutput(in).(ContentTypePtrOutput)
}

func (in *contentTypePtr) ToContentTypePtrOutputWithContext(ctx context.Context) ContentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContentTypePtrOutput)
}

type CustomEntityQueryKind string

const (
	CustomEntityQueryKindActivity = CustomEntityQueryKind("Activity")
)

func (CustomEntityQueryKind) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomEntityQueryKind)(nil)).Elem()
}

func (e CustomEntityQueryKind) ToCustomEntityQueryKindOutput() CustomEntityQueryKindOutput {
	return pulumi.ToOutput(e).(CustomEntityQueryKindOutput)
}

func (e CustomEntityQueryKind) ToCustomEntityQueryKindOutputWithContext(ctx context.Context) CustomEntityQueryKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CustomEntityQueryKindOutput)
}

func (e CustomEntityQueryKind) ToCustomEntityQueryKindPtrOutput() CustomEntityQueryKindPtrOutput {
	return e.ToCustomEntityQueryKindPtrOutputWithContext(context.Background())
}

func (e CustomEntityQueryKind) ToCustomEntityQueryKindPtrOutputWithContext(ctx context.Context) CustomEntityQueryKindPtrOutput {
	return CustomEntityQueryKind(e).ToCustomEntityQueryKindOutputWithContext(ctx).ToCustomEntityQueryKindPtrOutputWithContext(ctx)
}

func (e CustomEntityQueryKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomEntityQueryKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomEntityQueryKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CustomEntityQueryKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CustomEntityQueryKindOutput struct{ *pulumi.OutputState }

func (CustomEntityQueryKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomEntityQueryKind)(nil)).Elem()
}

func (o CustomEntityQueryKindOutput) ToCustomEntityQueryKindOutput() CustomEntityQueryKindOutput {
	return o
}

func (o CustomEntityQueryKindOutput) ToCustomEntityQueryKindOutputWithContext(ctx context.Context) CustomEntityQueryKindOutput {
	return o
}

func (o CustomEntityQueryKindOutput) ToCustomEntityQueryKindPtrOutput() CustomEntityQueryKindPtrOutput {
	return o.ToCustomEntityQueryKindPtrOutputWithContext(context.Background())
}

func (o CustomEntityQueryKindOutput) ToCustomEntityQueryKindPtrOutputWithContext(ctx context.Context) CustomEntityQueryKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomEntityQueryKind) *CustomEntityQueryKind {
		return &v
	}).(CustomEntityQueryKindPtrOutput)
}

func (o CustomEntityQueryKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CustomEntityQueryKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomEntityQueryKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CustomEntityQueryKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomEntityQueryKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomEntityQueryKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CustomEntityQueryKindPtrOutput struct{ *pulumi.OutputState }

func (CustomEntityQueryKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomEntityQueryKind)(nil)).Elem()
}

func (o CustomEntityQueryKindPtrOutput) ToCustomEntityQueryKindPtrOutput() CustomEntityQueryKindPtrOutput {
	return o
}

func (o CustomEntityQueryKindPtrOutput) ToCustomEntityQueryKindPtrOutputWithContext(ctx context.Context) CustomEntityQueryKindPtrOutput {
	return o
}

func (o CustomEntityQueryKindPtrOutput) Elem() CustomEntityQueryKindOutput {
	return o.ApplyT(func(v *CustomEntityQueryKind) CustomEntityQueryKind {
		if v != nil {
			return *v
		}
		var ret CustomEntityQueryKind
		return ret
	}).(CustomEntityQueryKindOutput)
}

func (o CustomEntityQueryKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomEntityQueryKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CustomEntityQueryKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CustomEntityQueryKindInput interface {
	pulumi.Input

	ToCustomEntityQueryKindOutput() CustomEntityQueryKindOutput
	ToCustomEntityQueryKindOutputWithContext(context.Context) CustomEntityQueryKindOutput
}

var customEntityQueryKindPtrType = reflect.TypeOf((**CustomEntityQueryKind)(nil)).Elem()

type CustomEntityQueryKindPtrInput interface {
	pulumi.Input

	ToCustomEntityQueryKindPtrOutput() CustomEntityQueryKindPtrOutput
	ToCustomEntityQueryKindPtrOutputWithContext(context.Context) CustomEntityQueryKindPtrOutput
}

type customEntityQueryKindPtr string

func CustomEntityQueryKindPtr(v string) CustomEntityQueryKindPtrInput {
	return (*customEntityQueryKindPtr)(&v)
}

func (*customEntityQueryKindPtr) ElementType() reflect.Type {
	return customEntityQueryKindPtrType
}

func (in *customEntityQueryKindPtr) ToCustomEntityQueryKindPtrOutput() CustomEntityQueryKindPtrOutput {
	return pulumi.ToOutput(in).(CustomEntityQueryKindPtrOutput)
}

func (in *customEntityQueryKindPtr) ToCustomEntityQueryKindPtrOutputWithContext(ctx context.Context) CustomEntityQueryKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CustomEntityQueryKindPtrOutput)
}

type DataConnectorKind string

const (
	DataConnectorKindAzureActiveDirectory                      = DataConnectorKind("AzureActiveDirectory")
	DataConnectorKindAzureSecurityCenter                       = DataConnectorKind("AzureSecurityCenter")
	DataConnectorKindMicrosoftCloudAppSecurity                 = DataConnectorKind("MicrosoftCloudAppSecurity")
	DataConnectorKindThreatIntelligence                        = DataConnectorKind("ThreatIntelligence")
	DataConnectorKindOffice365                                 = DataConnectorKind("Office365")
	DataConnectorKindAmazonWebServicesCloudTrail               = DataConnectorKind("AmazonWebServicesCloudTrail")
	DataConnectorKindAzureAdvancedThreatProtection             = DataConnectorKind("AzureAdvancedThreatProtection")
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection = DataConnectorKind("MicrosoftDefenderAdvancedThreatProtection")
)

func (DataConnectorKind) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorKind)(nil)).Elem()
}

func (e DataConnectorKind) ToDataConnectorKindOutput() DataConnectorKindOutput {
	return pulumi.ToOutput(e).(DataConnectorKindOutput)
}

func (e DataConnectorKind) ToDataConnectorKindOutputWithContext(ctx context.Context) DataConnectorKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataConnectorKindOutput)
}

func (e DataConnectorKind) ToDataConnectorKindPtrOutput() DataConnectorKindPtrOutput {
	return e.ToDataConnectorKindPtrOutputWithContext(context.Background())
}

func (e DataConnectorKind) ToDataConnectorKindPtrOutputWithContext(ctx context.Context) DataConnectorKindPtrOutput {
	return DataConnectorKind(e).ToDataConnectorKindOutputWithContext(ctx).ToDataConnectorKindPtrOutputWithContext(ctx)
}

func (e DataConnectorKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataConnectorKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataConnectorKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataConnectorKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataConnectorKindOutput struct{ *pulumi.OutputState }

func (DataConnectorKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorKind)(nil)).Elem()
}

func (o DataConnectorKindOutput) ToDataConnectorKindOutput() DataConnectorKindOutput {
	return o
}

func (o DataConnectorKindOutput) ToDataConnectorKindOutputWithContext(ctx context.Context) DataConnectorKindOutput {
	return o
}

func (o DataConnectorKindOutput) ToDataConnectorKindPtrOutput() DataConnectorKindPtrOutput {
	return o.ToDataConnectorKindPtrOutputWithContext(context.Background())
}

func (o DataConnectorKindOutput) ToDataConnectorKindPtrOutputWithContext(ctx context.Context) DataConnectorKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataConnectorKind) *DataConnectorKind {
		return &v
	}).(DataConnectorKindPtrOutput)
}

func (o DataConnectorKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataConnectorKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataConnectorKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataConnectorKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataConnectorKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataConnectorKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataConnectorKindPtrOutput struct{ *pulumi.OutputState }

func (DataConnectorKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorKind)(nil)).Elem()
}

func (o DataConnectorKindPtrOutput) ToDataConnectorKindPtrOutput() DataConnectorKindPtrOutput {
	return o
}

func (o DataConnectorKindPtrOutput) ToDataConnectorKindPtrOutputWithContext(ctx context.Context) DataConnectorKindPtrOutput {
	return o
}

func (o DataConnectorKindPtrOutput) Elem() DataConnectorKindOutput {
	return o.ApplyT(func(v *DataConnectorKind) DataConnectorKind {
		if v != nil {
			return *v
		}
		var ret DataConnectorKind
		return ret
	}).(DataConnectorKindOutput)
}

func (o DataConnectorKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataConnectorKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataConnectorKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataConnectorKindInput interface {
	pulumi.Input

	ToDataConnectorKindOutput() DataConnectorKindOutput
	ToDataConnectorKindOutputWithContext(context.Context) DataConnectorKindOutput
}

var dataConnectorKindPtrType = reflect.TypeOf((**DataConnectorKind)(nil)).Elem()

type DataConnectorKindPtrInput interface {
	pulumi.Input

	ToDataConnectorKindPtrOutput() DataConnectorKindPtrOutput
	ToDataConnectorKindPtrOutputWithContext(context.Context) DataConnectorKindPtrOutput
}

type dataConnectorKindPtr string

func DataConnectorKindPtr(v string) DataConnectorKindPtrInput {
	return (*dataConnectorKindPtr)(&v)
}

func (*dataConnectorKindPtr) ElementType() reflect.Type {
	return dataConnectorKindPtrType
}

func (in *dataConnectorKindPtr) ToDataConnectorKindPtrOutput() DataConnectorKindPtrOutput {
	return pulumi.ToOutput(in).(DataConnectorKindPtrOutput)
}

func (in *dataConnectorKindPtr) ToDataConnectorKindPtrOutputWithContext(ctx context.Context) DataConnectorKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataConnectorKindPtrOutput)
}

type DataTypeState string

const (
	DataTypeStateEnabled  = DataTypeState("Enabled")
	DataTypeStateDisabled = DataTypeState("Disabled")
)

func (DataTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*DataTypeState)(nil)).Elem()
}

func (e DataTypeState) ToDataTypeStateOutput() DataTypeStateOutput {
	return pulumi.ToOutput(e).(DataTypeStateOutput)
}

func (e DataTypeState) ToDataTypeStateOutputWithContext(ctx context.Context) DataTypeStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataTypeStateOutput)
}

func (e DataTypeState) ToDataTypeStatePtrOutput() DataTypeStatePtrOutput {
	return e.ToDataTypeStatePtrOutputWithContext(context.Background())
}

func (e DataTypeState) ToDataTypeStatePtrOutputWithContext(ctx context.Context) DataTypeStatePtrOutput {
	return DataTypeState(e).ToDataTypeStateOutputWithContext(ctx).ToDataTypeStatePtrOutputWithContext(ctx)
}

func (e DataTypeState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataTypeState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataTypeState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataTypeState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataTypeStateOutput struct{ *pulumi.OutputState }

func (DataTypeStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataTypeState)(nil)).Elem()
}

func (o DataTypeStateOutput) ToDataTypeStateOutput() DataTypeStateOutput {
	return o
}

func (o DataTypeStateOutput) ToDataTypeStateOutputWithContext(ctx context.Context) DataTypeStateOutput {
	return o
}

func (o DataTypeStateOutput) ToDataTypeStatePtrOutput() DataTypeStatePtrOutput {
	return o.ToDataTypeStatePtrOutputWithContext(context.Background())
}

func (o DataTypeStateOutput) ToDataTypeStatePtrOutputWithContext(ctx context.Context) DataTypeStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataTypeState) *DataTypeState {
		return &v
	}).(DataTypeStatePtrOutput)
}

func (o DataTypeStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataTypeStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataTypeState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataTypeStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataTypeStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataTypeState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataTypeStatePtrOutput struct{ *pulumi.OutputState }

func (DataTypeStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataTypeState)(nil)).Elem()
}

func (o DataTypeStatePtrOutput) ToDataTypeStatePtrOutput() DataTypeStatePtrOutput {
	return o
}

func (o DataTypeStatePtrOutput) ToDataTypeStatePtrOutputWithContext(ctx context.Context) DataTypeStatePtrOutput {
	return o
}

func (o DataTypeStatePtrOutput) Elem() DataTypeStateOutput {
	return o.ApplyT(func(v *DataTypeState) DataTypeState {
		if v != nil {
			return *v
		}
		var ret DataTypeState
		return ret
	}).(DataTypeStateOutput)
}

func (o DataTypeStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataTypeStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataTypeState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataTypeStateInput interface {
	pulumi.Input

	ToDataTypeStateOutput() DataTypeStateOutput
	ToDataTypeStateOutputWithContext(context.Context) DataTypeStateOutput
}

var dataTypeStatePtrType = reflect.TypeOf((**DataTypeState)(nil)).Elem()

type DataTypeStatePtrInput interface {
	pulumi.Input

	ToDataTypeStatePtrOutput() DataTypeStatePtrOutput
	ToDataTypeStatePtrOutputWithContext(context.Context) DataTypeStatePtrOutput
}

type dataTypeStatePtr string

func DataTypeStatePtr(v string) DataTypeStatePtrInput {
	return (*dataTypeStatePtr)(&v)
}

func (*dataTypeStatePtr) ElementType() reflect.Type {
	return dataTypeStatePtrType
}

func (in *dataTypeStatePtr) ToDataTypeStatePtrOutput() DataTypeStatePtrOutput {
	return pulumi.ToOutput(in).(DataTypeStatePtrOutput)
}

func (in *dataTypeStatePtr) ToDataTypeStatePtrOutputWithContext(ctx context.Context) DataTypeStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataTypeStatePtrOutput)
}

type EntityTimelineKind string

const (
	// activity
	EntityTimelineKindActivity = EntityTimelineKind("Activity")
	// bookmarks
	EntityTimelineKindBookmark = EntityTimelineKind("Bookmark")
	// security alerts
	EntityTimelineKindSecurityAlert = EntityTimelineKind("SecurityAlert")
)

func (EntityTimelineKind) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityTimelineKind)(nil)).Elem()
}

func (e EntityTimelineKind) ToEntityTimelineKindOutput() EntityTimelineKindOutput {
	return pulumi.ToOutput(e).(EntityTimelineKindOutput)
}

func (e EntityTimelineKind) ToEntityTimelineKindOutputWithContext(ctx context.Context) EntityTimelineKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EntityTimelineKindOutput)
}

func (e EntityTimelineKind) ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput {
	return e.ToEntityTimelineKindPtrOutputWithContext(context.Background())
}

func (e EntityTimelineKind) ToEntityTimelineKindPtrOutputWithContext(ctx context.Context) EntityTimelineKindPtrOutput {
	return EntityTimelineKind(e).ToEntityTimelineKindOutputWithContext(ctx).ToEntityTimelineKindPtrOutputWithContext(ctx)
}

func (e EntityTimelineKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityTimelineKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityTimelineKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EntityTimelineKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EntityTimelineKindOutput struct{ *pulumi.OutputState }

func (EntityTimelineKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityTimelineKind)(nil)).Elem()
}

func (o EntityTimelineKindOutput) ToEntityTimelineKindOutput() EntityTimelineKindOutput {
	return o
}

func (o EntityTimelineKindOutput) ToEntityTimelineKindOutputWithContext(ctx context.Context) EntityTimelineKindOutput {
	return o
}

func (o EntityTimelineKindOutput) ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput {
	return o.ToEntityTimelineKindPtrOutputWithContext(context.Background())
}

func (o EntityTimelineKindOutput) ToEntityTimelineKindPtrOutputWithContext(ctx context.Context) EntityTimelineKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EntityTimelineKind) *EntityTimelineKind {
		return &v
	}).(EntityTimelineKindPtrOutput)
}

func (o EntityTimelineKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EntityTimelineKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityTimelineKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EntityTimelineKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityTimelineKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityTimelineKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EntityTimelineKindPtrOutput struct{ *pulumi.OutputState }

func (EntityTimelineKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityTimelineKind)(nil)).Elem()
}

func (o EntityTimelineKindPtrOutput) ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput {
	return o
}

func (o EntityTimelineKindPtrOutput) ToEntityTimelineKindPtrOutputWithContext(ctx context.Context) EntityTimelineKindPtrOutput {
	return o
}

func (o EntityTimelineKindPtrOutput) Elem() EntityTimelineKindOutput {
	return o.ApplyT(func(v *EntityTimelineKind) EntityTimelineKind {
		if v != nil {
			return *v
		}
		var ret EntityTimelineKind
		return ret
	}).(EntityTimelineKindOutput)
}

func (o EntityTimelineKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityTimelineKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EntityTimelineKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EntityTimelineKindInput interface {
	pulumi.Input

	ToEntityTimelineKindOutput() EntityTimelineKindOutput
	ToEntityTimelineKindOutputWithContext(context.Context) EntityTimelineKindOutput
}

var entityTimelineKindPtrType = reflect.TypeOf((**EntityTimelineKind)(nil)).Elem()

type EntityTimelineKindPtrInput interface {
	pulumi.Input

	ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput
	ToEntityTimelineKindPtrOutputWithContext(context.Context) EntityTimelineKindPtrOutput
}

type entityTimelineKindPtr string

func EntityTimelineKindPtr(v string) EntityTimelineKindPtrInput {
	return (*entityTimelineKindPtr)(&v)
}

func (*entityTimelineKindPtr) ElementType() reflect.Type {
	return entityTimelineKindPtrType
}

func (in *entityTimelineKindPtr) ToEntityTimelineKindPtrOutput() EntityTimelineKindPtrOutput {
	return pulumi.ToOutput(in).(EntityTimelineKindPtrOutput)
}

func (in *entityTimelineKindPtr) ToEntityTimelineKindPtrOutputWithContext(ctx context.Context) EntityTimelineKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EntityTimelineKindPtrOutput)
}

type EntityType string

const (
	// Entity represents account in the system.
	EntityTypeAccount = EntityType("Account")
	// Entity represents host in the system.
	EntityTypeHost = EntityType("Host")
	// Entity represents file in the system.
	EntityTypeFile = EntityType("File")
	// Entity represents azure resource in the system.
	EntityTypeAzureResource = EntityType("AzureResource")
	// Entity represents cloud application in the system.
	EntityTypeCloudApplication = EntityType("CloudApplication")
	// Entity represents dns in the system.
	EntityTypeDNS = EntityType("DNS")
	// Entity represents file hash in the system.
	EntityTypeFileHash = EntityType("FileHash")
	// Entity represents ip in the system.
	EntityTypeIP = EntityType("IP")
	// Entity represents malware in the system.
	EntityTypeMalware = EntityType("Malware")
	// Entity represents process in the system.
	EntityTypeProcess = EntityType("Process")
	// Entity represents registry key in the system.
	EntityTypeRegistryKey = EntityType("RegistryKey")
	// Entity represents registry value in the system.
	EntityTypeRegistryValue = EntityType("RegistryValue")
	// Entity represents security group in the system.
	EntityTypeSecurityGroup = EntityType("SecurityGroup")
	// Entity represents url in the system.
	EntityTypeURL = EntityType("URL")
	// Entity represents IoT device in the system.
	EntityTypeIoTDevice = EntityType("IoTDevice")
	// Entity represents security alert in the system.
	EntityTypeSecurityAlert = EntityType("SecurityAlert")
	// Entity represents HuntingBookmark in the system.
	EntityTypeHuntingBookmark = EntityType("HuntingBookmark")
	// Entity represents mail cluster in the system.
	EntityTypeMailCluster = EntityType("MailCluster")
	// Entity represents mail message in the system.
	EntityTypeMailMessage = EntityType("MailMessage")
	// Entity represents mailbox in the system.
	EntityTypeMailbox = EntityType("Mailbox")
	// Entity represents submission mail in the system.
	EntityTypeSubmissionMail = EntityType("SubmissionMail")
)

func (EntityType) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityType)(nil)).Elem()
}

func (e EntityType) ToEntityTypeOutput() EntityTypeOutput {
	return pulumi.ToOutput(e).(EntityTypeOutput)
}

func (e EntityType) ToEntityTypeOutputWithContext(ctx context.Context) EntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EntityTypeOutput)
}

func (e EntityType) ToEntityTypePtrOutput() EntityTypePtrOutput {
	return e.ToEntityTypePtrOutputWithContext(context.Background())
}

func (e EntityType) ToEntityTypePtrOutputWithContext(ctx context.Context) EntityTypePtrOutput {
	return EntityType(e).ToEntityTypeOutputWithContext(ctx).ToEntityTypePtrOutputWithContext(ctx)
}

func (e EntityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EntityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EntityTypeOutput struct{ *pulumi.OutputState }

func (EntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityType)(nil)).Elem()
}

func (o EntityTypeOutput) ToEntityTypeOutput() EntityTypeOutput {
	return o
}

func (o EntityTypeOutput) ToEntityTypeOutputWithContext(ctx context.Context) EntityTypeOutput {
	return o
}

func (o EntityTypeOutput) ToEntityTypePtrOutput() EntityTypePtrOutput {
	return o.ToEntityTypePtrOutputWithContext(context.Background())
}

func (o EntityTypeOutput) ToEntityTypePtrOutputWithContext(ctx context.Context) EntityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EntityType) *EntityType {
		return &v
	}).(EntityTypePtrOutput)
}

func (o EntityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EntityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EntityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EntityTypePtrOutput struct{ *pulumi.OutputState }

func (EntityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityType)(nil)).Elem()
}

func (o EntityTypePtrOutput) ToEntityTypePtrOutput() EntityTypePtrOutput {
	return o
}

func (o EntityTypePtrOutput) ToEntityTypePtrOutputWithContext(ctx context.Context) EntityTypePtrOutput {
	return o
}

func (o EntityTypePtrOutput) Elem() EntityTypeOutput {
	return o.ApplyT(func(v *EntityType) EntityType {
		if v != nil {
			return *v
		}
		var ret EntityType
		return ret
	}).(EntityTypeOutput)
}

func (o EntityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EntityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EntityTypeInput interface {
	pulumi.Input

	ToEntityTypeOutput() EntityTypeOutput
	ToEntityTypeOutputWithContext(context.Context) EntityTypeOutput
}

var entityTypePtrType = reflect.TypeOf((**EntityType)(nil)).Elem()

type EntityTypePtrInput interface {
	pulumi.Input

	ToEntityTypePtrOutput() EntityTypePtrOutput
	ToEntityTypePtrOutputWithContext(context.Context) EntityTypePtrOutput
}

type entityTypePtr string

func EntityTypePtr(v string) EntityTypePtrInput {
	return (*entityTypePtr)(&v)
}

func (*entityTypePtr) ElementType() reflect.Type {
	return entityTypePtrType
}

func (in *entityTypePtr) ToEntityTypePtrOutput() EntityTypePtrOutput {
	return pulumi.ToOutput(in).(EntityTypePtrOutput)
}

func (in *entityTypePtr) ToEntityTypePtrOutputWithContext(ctx context.Context) EntityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EntityTypePtrOutput)
}

type IncidentClassification string

const (
	// Incident classification was undetermined
	IncidentClassificationUndetermined = IncidentClassification("Undetermined")
	// Incident was true positive
	IncidentClassificationTruePositive = IncidentClassification("TruePositive")
	// Incident was benign positive
	IncidentClassificationBenignPositive = IncidentClassification("BenignPositive")
	// Incident was false positive
	IncidentClassificationFalsePositive = IncidentClassification("FalsePositive")
)

func (IncidentClassification) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentClassification)(nil)).Elem()
}

func (e IncidentClassification) ToIncidentClassificationOutput() IncidentClassificationOutput {
	return pulumi.ToOutput(e).(IncidentClassificationOutput)
}

func (e IncidentClassification) ToIncidentClassificationOutputWithContext(ctx context.Context) IncidentClassificationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncidentClassificationOutput)
}

func (e IncidentClassification) ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput {
	return e.ToIncidentClassificationPtrOutputWithContext(context.Background())
}

func (e IncidentClassification) ToIncidentClassificationPtrOutputWithContext(ctx context.Context) IncidentClassificationPtrOutput {
	return IncidentClassification(e).ToIncidentClassificationOutputWithContext(ctx).ToIncidentClassificationPtrOutputWithContext(ctx)
}

func (e IncidentClassification) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassification) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassification) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentClassification) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncidentClassificationOutput struct{ *pulumi.OutputState }

func (IncidentClassificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentClassification)(nil)).Elem()
}

func (o IncidentClassificationOutput) ToIncidentClassificationOutput() IncidentClassificationOutput {
	return o
}

func (o IncidentClassificationOutput) ToIncidentClassificationOutputWithContext(ctx context.Context) IncidentClassificationOutput {
	return o
}

func (o IncidentClassificationOutput) ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput {
	return o.ToIncidentClassificationPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationOutput) ToIncidentClassificationPtrOutputWithContext(ctx context.Context) IncidentClassificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentClassification) *IncidentClassification {
		return &v
	}).(IncidentClassificationPtrOutput)
}

func (o IncidentClassificationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncidentClassificationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentClassification) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncidentClassificationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentClassification) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncidentClassificationPtrOutput struct{ *pulumi.OutputState }

func (IncidentClassificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentClassification)(nil)).Elem()
}

func (o IncidentClassificationPtrOutput) ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput {
	return o
}

func (o IncidentClassificationPtrOutput) ToIncidentClassificationPtrOutputWithContext(ctx context.Context) IncidentClassificationPtrOutput {
	return o
}

func (o IncidentClassificationPtrOutput) Elem() IncidentClassificationOutput {
	return o.ApplyT(func(v *IncidentClassification) IncidentClassification {
		if v != nil {
			return *v
		}
		var ret IncidentClassification
		return ret
	}).(IncidentClassificationOutput)
}

func (o IncidentClassificationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncidentClassification) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IncidentClassificationInput interface {
	pulumi.Input

	ToIncidentClassificationOutput() IncidentClassificationOutput
	ToIncidentClassificationOutputWithContext(context.Context) IncidentClassificationOutput
}

var incidentClassificationPtrType = reflect.TypeOf((**IncidentClassification)(nil)).Elem()

type IncidentClassificationPtrInput interface {
	pulumi.Input

	ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput
	ToIncidentClassificationPtrOutputWithContext(context.Context) IncidentClassificationPtrOutput
}

type incidentClassificationPtr string

func IncidentClassificationPtr(v string) IncidentClassificationPtrInput {
	return (*incidentClassificationPtr)(&v)
}

func (*incidentClassificationPtr) ElementType() reflect.Type {
	return incidentClassificationPtrType
}

func (in *incidentClassificationPtr) ToIncidentClassificationPtrOutput() IncidentClassificationPtrOutput {
	return pulumi.ToOutput(in).(IncidentClassificationPtrOutput)
}

func (in *incidentClassificationPtr) ToIncidentClassificationPtrOutputWithContext(ctx context.Context) IncidentClassificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncidentClassificationPtrOutput)
}

type IncidentClassificationReason string

const (
	// Classification reason was suspicious activity
	IncidentClassificationReasonSuspiciousActivity = IncidentClassificationReason("SuspiciousActivity")
	// Classification reason was suspicious but expected
	IncidentClassificationReasonSuspiciousButExpected = IncidentClassificationReason("SuspiciousButExpected")
	// Classification reason was incorrect alert logic
	IncidentClassificationReasonIncorrectAlertLogic = IncidentClassificationReason("IncorrectAlertLogic")
	// Classification reason was inaccurate data
	IncidentClassificationReasonInaccurateData = IncidentClassificationReason("InaccurateData")
)

func (IncidentClassificationReason) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentClassificationReason)(nil)).Elem()
}

func (e IncidentClassificationReason) ToIncidentClassificationReasonOutput() IncidentClassificationReasonOutput {
	return pulumi.ToOutput(e).(IncidentClassificationReasonOutput)
}

func (e IncidentClassificationReason) ToIncidentClassificationReasonOutputWithContext(ctx context.Context) IncidentClassificationReasonOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncidentClassificationReasonOutput)
}

func (e IncidentClassificationReason) ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput {
	return e.ToIncidentClassificationReasonPtrOutputWithContext(context.Background())
}

func (e IncidentClassificationReason) ToIncidentClassificationReasonPtrOutputWithContext(ctx context.Context) IncidentClassificationReasonPtrOutput {
	return IncidentClassificationReason(e).ToIncidentClassificationReasonOutputWithContext(ctx).ToIncidentClassificationReasonPtrOutputWithContext(ctx)
}

func (e IncidentClassificationReason) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassificationReason) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentClassificationReason) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentClassificationReason) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncidentClassificationReasonOutput struct{ *pulumi.OutputState }

func (IncidentClassificationReasonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentClassificationReason)(nil)).Elem()
}

func (o IncidentClassificationReasonOutput) ToIncidentClassificationReasonOutput() IncidentClassificationReasonOutput {
	return o
}

func (o IncidentClassificationReasonOutput) ToIncidentClassificationReasonOutputWithContext(ctx context.Context) IncidentClassificationReasonOutput {
	return o
}

func (o IncidentClassificationReasonOutput) ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput {
	return o.ToIncidentClassificationReasonPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationReasonOutput) ToIncidentClassificationReasonPtrOutputWithContext(ctx context.Context) IncidentClassificationReasonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentClassificationReason) *IncidentClassificationReason {
		return &v
	}).(IncidentClassificationReasonPtrOutput)
}

func (o IncidentClassificationReasonOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncidentClassificationReasonOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentClassificationReason) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncidentClassificationReasonOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationReasonOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentClassificationReason) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncidentClassificationReasonPtrOutput struct{ *pulumi.OutputState }

func (IncidentClassificationReasonPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentClassificationReason)(nil)).Elem()
}

func (o IncidentClassificationReasonPtrOutput) ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput {
	return o
}

func (o IncidentClassificationReasonPtrOutput) ToIncidentClassificationReasonPtrOutputWithContext(ctx context.Context) IncidentClassificationReasonPtrOutput {
	return o
}

func (o IncidentClassificationReasonPtrOutput) Elem() IncidentClassificationReasonOutput {
	return o.ApplyT(func(v *IncidentClassificationReason) IncidentClassificationReason {
		if v != nil {
			return *v
		}
		var ret IncidentClassificationReason
		return ret
	}).(IncidentClassificationReasonOutput)
}

func (o IncidentClassificationReasonPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentClassificationReasonPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncidentClassificationReason) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IncidentClassificationReasonInput interface {
	pulumi.Input

	ToIncidentClassificationReasonOutput() IncidentClassificationReasonOutput
	ToIncidentClassificationReasonOutputWithContext(context.Context) IncidentClassificationReasonOutput
}

var incidentClassificationReasonPtrType = reflect.TypeOf((**IncidentClassificationReason)(nil)).Elem()

type IncidentClassificationReasonPtrInput interface {
	pulumi.Input

	ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput
	ToIncidentClassificationReasonPtrOutputWithContext(context.Context) IncidentClassificationReasonPtrOutput
}

type incidentClassificationReasonPtr string

func IncidentClassificationReasonPtr(v string) IncidentClassificationReasonPtrInput {
	return (*incidentClassificationReasonPtr)(&v)
}

func (*incidentClassificationReasonPtr) ElementType() reflect.Type {
	return incidentClassificationReasonPtrType
}

func (in *incidentClassificationReasonPtr) ToIncidentClassificationReasonPtrOutput() IncidentClassificationReasonPtrOutput {
	return pulumi.ToOutput(in).(IncidentClassificationReasonPtrOutput)
}

func (in *incidentClassificationReasonPtr) ToIncidentClassificationReasonPtrOutputWithContext(ctx context.Context) IncidentClassificationReasonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncidentClassificationReasonPtrOutput)
}

type IncidentSeverity string

const (
	// High severity
	IncidentSeverityHigh = IncidentSeverity("High")
	// Medium severity
	IncidentSeverityMedium = IncidentSeverity("Medium")
	// Low severity
	IncidentSeverityLow = IncidentSeverity("Low")
	// Informational severity
	IncidentSeverityInformational = IncidentSeverity("Informational")
)

func (IncidentSeverity) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentSeverity)(nil)).Elem()
}

func (e IncidentSeverity) ToIncidentSeverityOutput() IncidentSeverityOutput {
	return pulumi.ToOutput(e).(IncidentSeverityOutput)
}

func (e IncidentSeverity) ToIncidentSeverityOutputWithContext(ctx context.Context) IncidentSeverityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncidentSeverityOutput)
}

func (e IncidentSeverity) ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput {
	return e.ToIncidentSeverityPtrOutputWithContext(context.Background())
}

func (e IncidentSeverity) ToIncidentSeverityPtrOutputWithContext(ctx context.Context) IncidentSeverityPtrOutput {
	return IncidentSeverity(e).ToIncidentSeverityOutputWithContext(ctx).ToIncidentSeverityPtrOutputWithContext(ctx)
}

func (e IncidentSeverity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentSeverity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentSeverity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentSeverity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncidentSeverityOutput struct{ *pulumi.OutputState }

func (IncidentSeverityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentSeverity)(nil)).Elem()
}

func (o IncidentSeverityOutput) ToIncidentSeverityOutput() IncidentSeverityOutput {
	return o
}

func (o IncidentSeverityOutput) ToIncidentSeverityOutputWithContext(ctx context.Context) IncidentSeverityOutput {
	return o
}

func (o IncidentSeverityOutput) ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput {
	return o.ToIncidentSeverityPtrOutputWithContext(context.Background())
}

func (o IncidentSeverityOutput) ToIncidentSeverityPtrOutputWithContext(ctx context.Context) IncidentSeverityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentSeverity) *IncidentSeverity {
		return &v
	}).(IncidentSeverityPtrOutput)
}

func (o IncidentSeverityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncidentSeverityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentSeverity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncidentSeverityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentSeverityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentSeverity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncidentSeverityPtrOutput struct{ *pulumi.OutputState }

func (IncidentSeverityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentSeverity)(nil)).Elem()
}

func (o IncidentSeverityPtrOutput) ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput {
	return o
}

func (o IncidentSeverityPtrOutput) ToIncidentSeverityPtrOutputWithContext(ctx context.Context) IncidentSeverityPtrOutput {
	return o
}

func (o IncidentSeverityPtrOutput) Elem() IncidentSeverityOutput {
	return o.ApplyT(func(v *IncidentSeverity) IncidentSeverity {
		if v != nil {
			return *v
		}
		var ret IncidentSeverity
		return ret
	}).(IncidentSeverityOutput)
}

func (o IncidentSeverityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentSeverityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncidentSeverity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IncidentSeverityInput interface {
	pulumi.Input

	ToIncidentSeverityOutput() IncidentSeverityOutput
	ToIncidentSeverityOutputWithContext(context.Context) IncidentSeverityOutput
}

var incidentSeverityPtrType = reflect.TypeOf((**IncidentSeverity)(nil)).Elem()

type IncidentSeverityPtrInput interface {
	pulumi.Input

	ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput
	ToIncidentSeverityPtrOutputWithContext(context.Context) IncidentSeverityPtrOutput
}

type incidentSeverityPtr string

func IncidentSeverityPtr(v string) IncidentSeverityPtrInput {
	return (*incidentSeverityPtr)(&v)
}

func (*incidentSeverityPtr) ElementType() reflect.Type {
	return incidentSeverityPtrType
}

func (in *incidentSeverityPtr) ToIncidentSeverityPtrOutput() IncidentSeverityPtrOutput {
	return pulumi.ToOutput(in).(IncidentSeverityPtrOutput)
}

func (in *incidentSeverityPtr) ToIncidentSeverityPtrOutputWithContext(ctx context.Context) IncidentSeverityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncidentSeverityPtrOutput)
}

type IncidentStatus string

const (
	// An active incident which isn't being handled currently
	IncidentStatusNew = IncidentStatus("New")
	// An active incident which is being handled
	IncidentStatusActive = IncidentStatus("Active")
	// A non-active incident
	IncidentStatusClosed = IncidentStatus("Closed")
)

func (IncidentStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentStatus)(nil)).Elem()
}

func (e IncidentStatus) ToIncidentStatusOutput() IncidentStatusOutput {
	return pulumi.ToOutput(e).(IncidentStatusOutput)
}

func (e IncidentStatus) ToIncidentStatusOutputWithContext(ctx context.Context) IncidentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IncidentStatusOutput)
}

func (e IncidentStatus) ToIncidentStatusPtrOutput() IncidentStatusPtrOutput {
	return e.ToIncidentStatusPtrOutputWithContext(context.Background())
}

func (e IncidentStatus) ToIncidentStatusPtrOutputWithContext(ctx context.Context) IncidentStatusPtrOutput {
	return IncidentStatus(e).ToIncidentStatusOutputWithContext(ctx).ToIncidentStatusPtrOutputWithContext(ctx)
}

func (e IncidentStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IncidentStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IncidentStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IncidentStatusOutput struct{ *pulumi.OutputState }

func (IncidentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentStatus)(nil)).Elem()
}

func (o IncidentStatusOutput) ToIncidentStatusOutput() IncidentStatusOutput {
	return o
}

func (o IncidentStatusOutput) ToIncidentStatusOutputWithContext(ctx context.Context) IncidentStatusOutput {
	return o
}

func (o IncidentStatusOutput) ToIncidentStatusPtrOutput() IncidentStatusPtrOutput {
	return o.ToIncidentStatusPtrOutputWithContext(context.Background())
}

func (o IncidentStatusOutput) ToIncidentStatusPtrOutputWithContext(ctx context.Context) IncidentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentStatus) *IncidentStatus {
		return &v
	}).(IncidentStatusPtrOutput)
}

func (o IncidentStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IncidentStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IncidentStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IncidentStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IncidentStatusPtrOutput struct{ *pulumi.OutputState }

func (IncidentStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentStatus)(nil)).Elem()
}

func (o IncidentStatusPtrOutput) ToIncidentStatusPtrOutput() IncidentStatusPtrOutput {
	return o
}

func (o IncidentStatusPtrOutput) ToIncidentStatusPtrOutputWithContext(ctx context.Context) IncidentStatusPtrOutput {
	return o
}

func (o IncidentStatusPtrOutput) Elem() IncidentStatusOutput {
	return o.ApplyT(func(v *IncidentStatus) IncidentStatus {
		if v != nil {
			return *v
		}
		var ret IncidentStatus
		return ret
	}).(IncidentStatusOutput)
}

func (o IncidentStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IncidentStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IncidentStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IncidentStatusInput interface {
	pulumi.Input

	ToIncidentStatusOutput() IncidentStatusOutput
	ToIncidentStatusOutputWithContext(context.Context) IncidentStatusOutput
}

var incidentStatusPtrType = reflect.TypeOf((**IncidentStatus)(nil)).Elem()

type IncidentStatusPtrInput interface {
	pulumi.Input

	ToIncidentStatusPtrOutput() IncidentStatusPtrOutput
	ToIncidentStatusPtrOutputWithContext(context.Context) IncidentStatusPtrOutput
}

type incidentStatusPtr string

func IncidentStatusPtr(v string) IncidentStatusPtrInput {
	return (*incidentStatusPtr)(&v)
}

func (*incidentStatusPtr) ElementType() reflect.Type {
	return incidentStatusPtrType
}

func (in *incidentStatusPtr) ToIncidentStatusPtrOutput() IncidentStatusPtrOutput {
	return pulumi.ToOutput(in).(IncidentStatusPtrOutput)
}

func (in *incidentStatusPtr) ToIncidentStatusPtrOutputWithContext(ctx context.Context) IncidentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IncidentStatusPtrOutput)
}

type Kind string

const (
	KindDataConnector         = Kind("DataConnector")
	KindDataType              = Kind("DataType")
	KindWorkbook              = Kind("Workbook")
	KindWorkbookTemplate      = Kind("WorkbookTemplate")
	KindPlaybook              = Kind("Playbook")
	KindPlaybookTemplate      = Kind("PlaybookTemplate")
	KindAnalyticsRuleTemplate = Kind("AnalyticsRuleTemplate")
	KindAnalyticsRule         = Kind("AnalyticsRule")
	KindHuntingQuery          = Kind("HuntingQuery")
	KindInvestigationQuery    = Kind("InvestigationQuery")
	KindParser                = Kind("Parser")
	KindWatchlist             = Kind("Watchlist")
	KindWatchlistTemplate     = Kind("WatchlistTemplate")
	KindSolution              = Kind("Solution")
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

type MicrosoftSecurityProductName string

const (
	MicrosoftSecurityProductName_Microsoft_Cloud_App_Security               = MicrosoftSecurityProductName("Microsoft Cloud App Security")
	MicrosoftSecurityProductName_Azure_Security_Center                      = MicrosoftSecurityProductName("Azure Security Center")
	MicrosoftSecurityProductName_Azure_Advanced_Threat_Protection           = MicrosoftSecurityProductName("Azure Advanced Threat Protection")
	MicrosoftSecurityProductName_Azure_Active_Directory_Identity_Protection = MicrosoftSecurityProductName("Azure Active Directory Identity Protection")
	MicrosoftSecurityProductName_Azure_Security_Center_for_IoT              = MicrosoftSecurityProductName("Azure Security Center for IoT")
)

func (MicrosoftSecurityProductName) ElementType() reflect.Type {
	return reflect.TypeOf((*MicrosoftSecurityProductName)(nil)).Elem()
}

func (e MicrosoftSecurityProductName) ToMicrosoftSecurityProductNameOutput() MicrosoftSecurityProductNameOutput {
	return pulumi.ToOutput(e).(MicrosoftSecurityProductNameOutput)
}

func (e MicrosoftSecurityProductName) ToMicrosoftSecurityProductNameOutputWithContext(ctx context.Context) MicrosoftSecurityProductNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MicrosoftSecurityProductNameOutput)
}

func (e MicrosoftSecurityProductName) ToMicrosoftSecurityProductNamePtrOutput() MicrosoftSecurityProductNamePtrOutput {
	return e.ToMicrosoftSecurityProductNamePtrOutputWithContext(context.Background())
}

func (e MicrosoftSecurityProductName) ToMicrosoftSecurityProductNamePtrOutputWithContext(ctx context.Context) MicrosoftSecurityProductNamePtrOutput {
	return MicrosoftSecurityProductName(e).ToMicrosoftSecurityProductNameOutputWithContext(ctx).ToMicrosoftSecurityProductNamePtrOutputWithContext(ctx)
}

func (e MicrosoftSecurityProductName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MicrosoftSecurityProductName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MicrosoftSecurityProductName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MicrosoftSecurityProductName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MicrosoftSecurityProductNameOutput struct{ *pulumi.OutputState }

func (MicrosoftSecurityProductNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MicrosoftSecurityProductName)(nil)).Elem()
}

func (o MicrosoftSecurityProductNameOutput) ToMicrosoftSecurityProductNameOutput() MicrosoftSecurityProductNameOutput {
	return o
}

func (o MicrosoftSecurityProductNameOutput) ToMicrosoftSecurityProductNameOutputWithContext(ctx context.Context) MicrosoftSecurityProductNameOutput {
	return o
}

func (o MicrosoftSecurityProductNameOutput) ToMicrosoftSecurityProductNamePtrOutput() MicrosoftSecurityProductNamePtrOutput {
	return o.ToMicrosoftSecurityProductNamePtrOutputWithContext(context.Background())
}

func (o MicrosoftSecurityProductNameOutput) ToMicrosoftSecurityProductNamePtrOutputWithContext(ctx context.Context) MicrosoftSecurityProductNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MicrosoftSecurityProductName) *MicrosoftSecurityProductName {
		return &v
	}).(MicrosoftSecurityProductNamePtrOutput)
}

func (o MicrosoftSecurityProductNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MicrosoftSecurityProductNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MicrosoftSecurityProductName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MicrosoftSecurityProductNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MicrosoftSecurityProductNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MicrosoftSecurityProductName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MicrosoftSecurityProductNamePtrOutput struct{ *pulumi.OutputState }

func (MicrosoftSecurityProductNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftSecurityProductName)(nil)).Elem()
}

func (o MicrosoftSecurityProductNamePtrOutput) ToMicrosoftSecurityProductNamePtrOutput() MicrosoftSecurityProductNamePtrOutput {
	return o
}

func (o MicrosoftSecurityProductNamePtrOutput) ToMicrosoftSecurityProductNamePtrOutputWithContext(ctx context.Context) MicrosoftSecurityProductNamePtrOutput {
	return o
}

func (o MicrosoftSecurityProductNamePtrOutput) Elem() MicrosoftSecurityProductNameOutput {
	return o.ApplyT(func(v *MicrosoftSecurityProductName) MicrosoftSecurityProductName {
		if v != nil {
			return *v
		}
		var ret MicrosoftSecurityProductName
		return ret
	}).(MicrosoftSecurityProductNameOutput)
}

func (o MicrosoftSecurityProductNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MicrosoftSecurityProductNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MicrosoftSecurityProductName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MicrosoftSecurityProductNameInput interface {
	pulumi.Input

	ToMicrosoftSecurityProductNameOutput() MicrosoftSecurityProductNameOutput
	ToMicrosoftSecurityProductNameOutputWithContext(context.Context) MicrosoftSecurityProductNameOutput
}

var microsoftSecurityProductNamePtrType = reflect.TypeOf((**MicrosoftSecurityProductName)(nil)).Elem()

type MicrosoftSecurityProductNamePtrInput interface {
	pulumi.Input

	ToMicrosoftSecurityProductNamePtrOutput() MicrosoftSecurityProductNamePtrOutput
	ToMicrosoftSecurityProductNamePtrOutputWithContext(context.Context) MicrosoftSecurityProductNamePtrOutput
}

type microsoftSecurityProductNamePtr string

func MicrosoftSecurityProductNamePtr(v string) MicrosoftSecurityProductNamePtrInput {
	return (*microsoftSecurityProductNamePtr)(&v)
}

func (*microsoftSecurityProductNamePtr) ElementType() reflect.Type {
	return microsoftSecurityProductNamePtrType
}

func (in *microsoftSecurityProductNamePtr) ToMicrosoftSecurityProductNamePtrOutput() MicrosoftSecurityProductNamePtrOutput {
	return pulumi.ToOutput(in).(MicrosoftSecurityProductNamePtrOutput)
}

func (in *microsoftSecurityProductNamePtr) ToMicrosoftSecurityProductNamePtrOutputWithContext(ctx context.Context) MicrosoftSecurityProductNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MicrosoftSecurityProductNamePtrOutput)
}

type Operator string

const (
	OperatorAND = Operator("AND")
	OperatorOR  = Operator("OR")
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

type RepoType string

const (
	RepoTypeGithub = RepoType("Github")
	RepoTypeDevOps = RepoType("DevOps")
)

func (RepoType) ElementType() reflect.Type {
	return reflect.TypeOf((*RepoType)(nil)).Elem()
}

func (e RepoType) ToRepoTypeOutput() RepoTypeOutput {
	return pulumi.ToOutput(e).(RepoTypeOutput)
}

func (e RepoType) ToRepoTypeOutputWithContext(ctx context.Context) RepoTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RepoTypeOutput)
}

func (e RepoType) ToRepoTypePtrOutput() RepoTypePtrOutput {
	return e.ToRepoTypePtrOutputWithContext(context.Background())
}

func (e RepoType) ToRepoTypePtrOutputWithContext(ctx context.Context) RepoTypePtrOutput {
	return RepoType(e).ToRepoTypeOutputWithContext(ctx).ToRepoTypePtrOutputWithContext(ctx)
}

func (e RepoType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RepoType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RepoType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RepoType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RepoTypeOutput struct{ *pulumi.OutputState }

func (RepoTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepoType)(nil)).Elem()
}

func (o RepoTypeOutput) ToRepoTypeOutput() RepoTypeOutput {
	return o
}

func (o RepoTypeOutput) ToRepoTypeOutputWithContext(ctx context.Context) RepoTypeOutput {
	return o
}

func (o RepoTypeOutput) ToRepoTypePtrOutput() RepoTypePtrOutput {
	return o.ToRepoTypePtrOutputWithContext(context.Background())
}

func (o RepoTypeOutput) ToRepoTypePtrOutputWithContext(ctx context.Context) RepoTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RepoType) *RepoType {
		return &v
	}).(RepoTypePtrOutput)
}

func (o RepoTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RepoTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RepoType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RepoTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RepoTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RepoType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RepoTypePtrOutput struct{ *pulumi.OutputState }

func (RepoTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepoType)(nil)).Elem()
}

func (o RepoTypePtrOutput) ToRepoTypePtrOutput() RepoTypePtrOutput {
	return o
}

func (o RepoTypePtrOutput) ToRepoTypePtrOutputWithContext(ctx context.Context) RepoTypePtrOutput {
	return o
}

func (o RepoTypePtrOutput) Elem() RepoTypeOutput {
	return o.ApplyT(func(v *RepoType) RepoType {
		if v != nil {
			return *v
		}
		var ret RepoType
		return ret
	}).(RepoTypeOutput)
}

func (o RepoTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RepoTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RepoType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RepoTypeInput interface {
	pulumi.Input

	ToRepoTypeOutput() RepoTypeOutput
	ToRepoTypeOutputWithContext(context.Context) RepoTypeOutput
}

var repoTypePtrType = reflect.TypeOf((**RepoType)(nil)).Elem()

type RepoTypePtrInput interface {
	pulumi.Input

	ToRepoTypePtrOutput() RepoTypePtrOutput
	ToRepoTypePtrOutputWithContext(context.Context) RepoTypePtrOutput
}

type repoTypePtr string

func RepoTypePtr(v string) RepoTypePtrInput {
	return (*repoTypePtr)(&v)
}

func (*repoTypePtr) ElementType() reflect.Type {
	return repoTypePtrType
}

func (in *repoTypePtr) ToRepoTypePtrOutput() RepoTypePtrOutput {
	return pulumi.ToOutput(in).(RepoTypePtrOutput)
}

func (in *repoTypePtr) ToRepoTypePtrOutputWithContext(ctx context.Context) RepoTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RepoTypePtrOutput)
}

type SettingKind string

const (
	SettingKindAnomalies       = SettingKind("Anomalies")
	SettingKindEyesOn          = SettingKind("EyesOn")
	SettingKindEntityAnalytics = SettingKind("EntityAnalytics")
	SettingKindUeba            = SettingKind("Ueba")
)

func (SettingKind) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingKind)(nil)).Elem()
}

func (e SettingKind) ToSettingKindOutput() SettingKindOutput {
	return pulumi.ToOutput(e).(SettingKindOutput)
}

func (e SettingKind) ToSettingKindOutputWithContext(ctx context.Context) SettingKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SettingKindOutput)
}

func (e SettingKind) ToSettingKindPtrOutput() SettingKindPtrOutput {
	return e.ToSettingKindPtrOutputWithContext(context.Background())
}

func (e SettingKind) ToSettingKindPtrOutputWithContext(ctx context.Context) SettingKindPtrOutput {
	return SettingKind(e).ToSettingKindOutputWithContext(ctx).ToSettingKindPtrOutputWithContext(ctx)
}

func (e SettingKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SettingKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SettingKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SettingKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SettingKindOutput struct{ *pulumi.OutputState }

func (SettingKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingKind)(nil)).Elem()
}

func (o SettingKindOutput) ToSettingKindOutput() SettingKindOutput {
	return o
}

func (o SettingKindOutput) ToSettingKindOutputWithContext(ctx context.Context) SettingKindOutput {
	return o
}

func (o SettingKindOutput) ToSettingKindPtrOutput() SettingKindPtrOutput {
	return o.ToSettingKindPtrOutputWithContext(context.Background())
}

func (o SettingKindOutput) ToSettingKindPtrOutputWithContext(ctx context.Context) SettingKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SettingKind) *SettingKind {
		return &v
	}).(SettingKindPtrOutput)
}

func (o SettingKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SettingKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SettingKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SettingKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SettingKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SettingKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SettingKindPtrOutput struct{ *pulumi.OutputState }

func (SettingKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingKind)(nil)).Elem()
}

func (o SettingKindPtrOutput) ToSettingKindPtrOutput() SettingKindPtrOutput {
	return o
}

func (o SettingKindPtrOutput) ToSettingKindPtrOutputWithContext(ctx context.Context) SettingKindPtrOutput {
	return o
}

func (o SettingKindPtrOutput) Elem() SettingKindOutput {
	return o.ApplyT(func(v *SettingKind) SettingKind {
		if v != nil {
			return *v
		}
		var ret SettingKind
		return ret
	}).(SettingKindOutput)
}

func (o SettingKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SettingKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SettingKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SettingKindInput interface {
	pulumi.Input

	ToSettingKindOutput() SettingKindOutput
	ToSettingKindOutputWithContext(context.Context) SettingKindOutput
}

var settingKindPtrType = reflect.TypeOf((**SettingKind)(nil)).Elem()

type SettingKindPtrInput interface {
	pulumi.Input

	ToSettingKindPtrOutput() SettingKindPtrOutput
	ToSettingKindPtrOutputWithContext(context.Context) SettingKindPtrOutput
}

type settingKindPtr string

func SettingKindPtr(v string) SettingKindPtrInput {
	return (*settingKindPtr)(&v)
}

func (*settingKindPtr) ElementType() reflect.Type {
	return settingKindPtrType
}

func (in *settingKindPtr) ToSettingKindPtrOutput() SettingKindPtrOutput {
	return pulumi.ToOutput(in).(SettingKindPtrOutput)
}

func (in *settingKindPtr) ToSettingKindPtrOutputWithContext(ctx context.Context) SettingKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SettingKindPtrOutput)
}

type Source string

const (
	Source_Local_file     = Source("Local file")
	Source_Remote_storage = Source("Remote storage")
)

func (Source) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (e Source) ToSourceOutput() SourceOutput {
	return pulumi.ToOutput(e).(SourceOutput)
}

func (e Source) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceOutput)
}

func (e Source) ToSourcePtrOutput() SourcePtrOutput {
	return e.ToSourcePtrOutputWithContext(context.Background())
}

func (e Source) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return Source(e).ToSourceOutputWithContext(ctx).ToSourcePtrOutputWithContext(ctx)
}

func (e Source) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Source) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Source) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Source) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceOutput struct{ *pulumi.OutputState }

func (SourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (o SourceOutput) ToSourceOutput() SourceOutput {
	return o
}

func (o SourceOutput) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return o
}

func (o SourceOutput) ToSourcePtrOutput() SourcePtrOutput {
	return o.ToSourcePtrOutputWithContext(context.Background())
}

func (o SourceOutput) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Source) *Source {
		return &v
	}).(SourcePtrOutput)
}

func (o SourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Source) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Source) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourcePtrOutput struct{ *pulumi.OutputState }

func (SourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Source)(nil)).Elem()
}

func (o SourcePtrOutput) ToSourcePtrOutput() SourcePtrOutput {
	return o
}

func (o SourcePtrOutput) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return o
}

func (o SourcePtrOutput) Elem() SourceOutput {
	return o.ApplyT(func(v *Source) Source {
		if v != nil {
			return *v
		}
		var ret Source
		return ret
	}).(SourceOutput)
}

func (o SourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Source) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceInput interface {
	pulumi.Input

	ToSourceOutput() SourceOutput
	ToSourceOutputWithContext(context.Context) SourceOutput
}

var sourcePtrType = reflect.TypeOf((**Source)(nil)).Elem()

type SourcePtrInput interface {
	pulumi.Input

	ToSourcePtrOutput() SourcePtrOutput
	ToSourcePtrOutputWithContext(context.Context) SourcePtrOutput
}

type sourcePtr string

func SourcePtr(v string) SourcePtrInput {
	return (*sourcePtr)(&v)
}

func (*sourcePtr) ElementType() reflect.Type {
	return sourcePtrType
}

func (in *sourcePtr) ToSourcePtrOutput() SourcePtrOutput {
	return pulumi.ToOutput(in).(SourcePtrOutput)
}

func (in *sourcePtr) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourcePtrOutput)
}

type SourceKind string

const (
	SourceKindLocalWorkspace   = SourceKind("LocalWorkspace")
	SourceKindCommunity        = SourceKind("Community")
	SourceKindSolution         = SourceKind("Solution")
	SourceKindSourceRepository = SourceKind("SourceRepository")
)

func (SourceKind) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceKind)(nil)).Elem()
}

func (e SourceKind) ToSourceKindOutput() SourceKindOutput {
	return pulumi.ToOutput(e).(SourceKindOutput)
}

func (e SourceKind) ToSourceKindOutputWithContext(ctx context.Context) SourceKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceKindOutput)
}

func (e SourceKind) ToSourceKindPtrOutput() SourceKindPtrOutput {
	return e.ToSourceKindPtrOutputWithContext(context.Background())
}

func (e SourceKind) ToSourceKindPtrOutputWithContext(ctx context.Context) SourceKindPtrOutput {
	return SourceKind(e).ToSourceKindOutputWithContext(ctx).ToSourceKindPtrOutputWithContext(ctx)
}

func (e SourceKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceKindOutput struct{ *pulumi.OutputState }

func (SourceKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceKind)(nil)).Elem()
}

func (o SourceKindOutput) ToSourceKindOutput() SourceKindOutput {
	return o
}

func (o SourceKindOutput) ToSourceKindOutputWithContext(ctx context.Context) SourceKindOutput {
	return o
}

func (o SourceKindOutput) ToSourceKindPtrOutput() SourceKindPtrOutput {
	return o.ToSourceKindPtrOutputWithContext(context.Background())
}

func (o SourceKindOutput) ToSourceKindPtrOutputWithContext(ctx context.Context) SourceKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceKind) *SourceKind {
		return &v
	}).(SourceKindPtrOutput)
}

func (o SourceKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceKindPtrOutput struct{ *pulumi.OutputState }

func (SourceKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceKind)(nil)).Elem()
}

func (o SourceKindPtrOutput) ToSourceKindPtrOutput() SourceKindPtrOutput {
	return o
}

func (o SourceKindPtrOutput) ToSourceKindPtrOutputWithContext(ctx context.Context) SourceKindPtrOutput {
	return o
}

func (o SourceKindPtrOutput) Elem() SourceKindOutput {
	return o.ApplyT(func(v *SourceKind) SourceKind {
		if v != nil {
			return *v
		}
		var ret SourceKind
		return ret
	}).(SourceKindOutput)
}

func (o SourceKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceKindInput interface {
	pulumi.Input

	ToSourceKindOutput() SourceKindOutput
	ToSourceKindOutputWithContext(context.Context) SourceKindOutput
}

var sourceKindPtrType = reflect.TypeOf((**SourceKind)(nil)).Elem()

type SourceKindPtrInput interface {
	pulumi.Input

	ToSourceKindPtrOutput() SourceKindPtrOutput
	ToSourceKindPtrOutputWithContext(context.Context) SourceKindPtrOutput
}

type sourceKindPtr string

func SourceKindPtr(v string) SourceKindPtrInput {
	return (*sourceKindPtr)(&v)
}

func (*sourceKindPtr) ElementType() reflect.Type {
	return sourceKindPtrType
}

func (in *sourceKindPtr) ToSourceKindPtrOutput() SourceKindPtrOutput {
	return pulumi.ToOutput(in).(SourceKindPtrOutput)
}

func (in *sourceKindPtr) ToSourceKindPtrOutputWithContext(ctx context.Context) SourceKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceKindPtrOutput)
}

type SupportTier string

const (
	SupportTierMicrosoft = SupportTier("Microsoft")
	SupportTierPartner   = SupportTier("Partner")
	SupportTierCommunity = SupportTier("Community")
)

func (SupportTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportTier)(nil)).Elem()
}

func (e SupportTier) ToSupportTierOutput() SupportTierOutput {
	return pulumi.ToOutput(e).(SupportTierOutput)
}

func (e SupportTier) ToSupportTierOutputWithContext(ctx context.Context) SupportTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SupportTierOutput)
}

func (e SupportTier) ToSupportTierPtrOutput() SupportTierPtrOutput {
	return e.ToSupportTierPtrOutputWithContext(context.Background())
}

func (e SupportTier) ToSupportTierPtrOutputWithContext(ctx context.Context) SupportTierPtrOutput {
	return SupportTier(e).ToSupportTierOutputWithContext(ctx).ToSupportTierPtrOutputWithContext(ctx)
}

func (e SupportTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SupportTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SupportTierOutput struct{ *pulumi.OutputState }

func (SupportTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportTier)(nil)).Elem()
}

func (o SupportTierOutput) ToSupportTierOutput() SupportTierOutput {
	return o
}

func (o SupportTierOutput) ToSupportTierOutputWithContext(ctx context.Context) SupportTierOutput {
	return o
}

func (o SupportTierOutput) ToSupportTierPtrOutput() SupportTierPtrOutput {
	return o.ToSupportTierPtrOutputWithContext(context.Background())
}

func (o SupportTierOutput) ToSupportTierPtrOutputWithContext(ctx context.Context) SupportTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportTier) *SupportTier {
		return &v
	}).(SupportTierPtrOutput)
}

func (o SupportTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SupportTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SupportTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SupportTierPtrOutput struct{ *pulumi.OutputState }

func (SupportTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportTier)(nil)).Elem()
}

func (o SupportTierPtrOutput) ToSupportTierPtrOutput() SupportTierPtrOutput {
	return o
}

func (o SupportTierPtrOutput) ToSupportTierPtrOutputWithContext(ctx context.Context) SupportTierPtrOutput {
	return o
}

func (o SupportTierPtrOutput) Elem() SupportTierOutput {
	return o.ApplyT(func(v *SupportTier) SupportTier {
		if v != nil {
			return *v
		}
		var ret SupportTier
		return ret
	}).(SupportTierOutput)
}

func (o SupportTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SupportTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SupportTierInput interface {
	pulumi.Input

	ToSupportTierOutput() SupportTierOutput
	ToSupportTierOutputWithContext(context.Context) SupportTierOutput
}

var supportTierPtrType = reflect.TypeOf((**SupportTier)(nil)).Elem()

type SupportTierPtrInput interface {
	pulumi.Input

	ToSupportTierPtrOutput() SupportTierPtrOutput
	ToSupportTierPtrOutputWithContext(context.Context) SupportTierPtrOutput
}

type supportTierPtr string

func SupportTierPtr(v string) SupportTierPtrInput {
	return (*supportTierPtr)(&v)
}

func (*supportTierPtr) ElementType() reflect.Type {
	return supportTierPtrType
}

func (in *supportTierPtr) ToSupportTierPtrOutput() SupportTierPtrOutput {
	return pulumi.ToOutput(in).(SupportTierPtrOutput)
}

func (in *supportTierPtr) ToSupportTierPtrOutputWithContext(ctx context.Context) SupportTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SupportTierPtrOutput)
}

type ThreatIntelligenceResourceKind string

const (
	// Entity represents threat intelligence indicator in the system.
	ThreatIntelligenceResourceKindIndicator = ThreatIntelligenceResourceKind("indicator")
)

func (ThreatIntelligenceResourceKind) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceResourceKind)(nil)).Elem()
}

func (e ThreatIntelligenceResourceKind) ToThreatIntelligenceResourceKindOutput() ThreatIntelligenceResourceKindOutput {
	return pulumi.ToOutput(e).(ThreatIntelligenceResourceKindOutput)
}

func (e ThreatIntelligenceResourceKind) ToThreatIntelligenceResourceKindOutputWithContext(ctx context.Context) ThreatIntelligenceResourceKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ThreatIntelligenceResourceKindOutput)
}

func (e ThreatIntelligenceResourceKind) ToThreatIntelligenceResourceKindPtrOutput() ThreatIntelligenceResourceKindPtrOutput {
	return e.ToThreatIntelligenceResourceKindPtrOutputWithContext(context.Background())
}

func (e ThreatIntelligenceResourceKind) ToThreatIntelligenceResourceKindPtrOutputWithContext(ctx context.Context) ThreatIntelligenceResourceKindPtrOutput {
	return ThreatIntelligenceResourceKind(e).ToThreatIntelligenceResourceKindOutputWithContext(ctx).ToThreatIntelligenceResourceKindPtrOutputWithContext(ctx)
}

func (e ThreatIntelligenceResourceKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ThreatIntelligenceResourceKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ThreatIntelligenceResourceKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ThreatIntelligenceResourceKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ThreatIntelligenceResourceKindOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceResourceKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceResourceKind)(nil)).Elem()
}

func (o ThreatIntelligenceResourceKindOutput) ToThreatIntelligenceResourceKindOutput() ThreatIntelligenceResourceKindOutput {
	return o
}

func (o ThreatIntelligenceResourceKindOutput) ToThreatIntelligenceResourceKindOutputWithContext(ctx context.Context) ThreatIntelligenceResourceKindOutput {
	return o
}

func (o ThreatIntelligenceResourceKindOutput) ToThreatIntelligenceResourceKindPtrOutput() ThreatIntelligenceResourceKindPtrOutput {
	return o.ToThreatIntelligenceResourceKindPtrOutputWithContext(context.Background())
}

func (o ThreatIntelligenceResourceKindOutput) ToThreatIntelligenceResourceKindPtrOutputWithContext(ctx context.Context) ThreatIntelligenceResourceKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ThreatIntelligenceResourceKind) *ThreatIntelligenceResourceKind {
		return &v
	}).(ThreatIntelligenceResourceKindPtrOutput)
}

func (o ThreatIntelligenceResourceKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ThreatIntelligenceResourceKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ThreatIntelligenceResourceKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ThreatIntelligenceResourceKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ThreatIntelligenceResourceKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ThreatIntelligenceResourceKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ThreatIntelligenceResourceKindPtrOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceResourceKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelligenceResourceKind)(nil)).Elem()
}

func (o ThreatIntelligenceResourceKindPtrOutput) ToThreatIntelligenceResourceKindPtrOutput() ThreatIntelligenceResourceKindPtrOutput {
	return o
}

func (o ThreatIntelligenceResourceKindPtrOutput) ToThreatIntelligenceResourceKindPtrOutputWithContext(ctx context.Context) ThreatIntelligenceResourceKindPtrOutput {
	return o
}

func (o ThreatIntelligenceResourceKindPtrOutput) Elem() ThreatIntelligenceResourceKindOutput {
	return o.ApplyT(func(v *ThreatIntelligenceResourceKind) ThreatIntelligenceResourceKind {
		if v != nil {
			return *v
		}
		var ret ThreatIntelligenceResourceKind
		return ret
	}).(ThreatIntelligenceResourceKindOutput)
}

func (o ThreatIntelligenceResourceKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ThreatIntelligenceResourceKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ThreatIntelligenceResourceKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ThreatIntelligenceResourceKindInput interface {
	pulumi.Input

	ToThreatIntelligenceResourceKindOutput() ThreatIntelligenceResourceKindOutput
	ToThreatIntelligenceResourceKindOutputWithContext(context.Context) ThreatIntelligenceResourceKindOutput
}

var threatIntelligenceResourceKindPtrType = reflect.TypeOf((**ThreatIntelligenceResourceKind)(nil)).Elem()

type ThreatIntelligenceResourceKindPtrInput interface {
	pulumi.Input

	ToThreatIntelligenceResourceKindPtrOutput() ThreatIntelligenceResourceKindPtrOutput
	ToThreatIntelligenceResourceKindPtrOutputWithContext(context.Context) ThreatIntelligenceResourceKindPtrOutput
}

type threatIntelligenceResourceKindPtr string

func ThreatIntelligenceResourceKindPtr(v string) ThreatIntelligenceResourceKindPtrInput {
	return (*threatIntelligenceResourceKindPtr)(&v)
}

func (*threatIntelligenceResourceKindPtr) ElementType() reflect.Type {
	return threatIntelligenceResourceKindPtrType
}

func (in *threatIntelligenceResourceKindPtr) ToThreatIntelligenceResourceKindPtrOutput() ThreatIntelligenceResourceKindPtrOutput {
	return pulumi.ToOutput(in).(ThreatIntelligenceResourceKindPtrOutput)
}

func (in *threatIntelligenceResourceKindPtr) ToThreatIntelligenceResourceKindPtrOutputWithContext(ctx context.Context) ThreatIntelligenceResourceKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ThreatIntelligenceResourceKindPtrOutput)
}

type TriggerOperator string

const (
	TriggerOperatorGreaterThan = TriggerOperator("GreaterThan")
	TriggerOperatorLessThan    = TriggerOperator("LessThan")
	TriggerOperatorEqual       = TriggerOperator("Equal")
	TriggerOperatorNotEqual    = TriggerOperator("NotEqual")
)

func (TriggerOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerOperator)(nil)).Elem()
}

func (e TriggerOperator) ToTriggerOperatorOutput() TriggerOperatorOutput {
	return pulumi.ToOutput(e).(TriggerOperatorOutput)
}

func (e TriggerOperator) ToTriggerOperatorOutputWithContext(ctx context.Context) TriggerOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TriggerOperatorOutput)
}

func (e TriggerOperator) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return e.ToTriggerOperatorPtrOutputWithContext(context.Background())
}

func (e TriggerOperator) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return TriggerOperator(e).ToTriggerOperatorOutputWithContext(ctx).ToTriggerOperatorPtrOutputWithContext(ctx)
}

func (e TriggerOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggerOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TriggerOperatorOutput struct{ *pulumi.OutputState }

func (TriggerOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerOperator)(nil)).Elem()
}

func (o TriggerOperatorOutput) ToTriggerOperatorOutput() TriggerOperatorOutput {
	return o
}

func (o TriggerOperatorOutput) ToTriggerOperatorOutputWithContext(ctx context.Context) TriggerOperatorOutput {
	return o
}

func (o TriggerOperatorOutput) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return o.ToTriggerOperatorPtrOutputWithContext(context.Background())
}

func (o TriggerOperatorOutput) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerOperator) *TriggerOperator {
		return &v
	}).(TriggerOperatorPtrOutput)
}

func (o TriggerOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TriggerOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TriggerOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TriggerOperatorPtrOutput struct{ *pulumi.OutputState }

func (TriggerOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerOperator)(nil)).Elem()
}

func (o TriggerOperatorPtrOutput) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return o
}

func (o TriggerOperatorPtrOutput) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return o
}

func (o TriggerOperatorPtrOutput) Elem() TriggerOperatorOutput {
	return o.ApplyT(func(v *TriggerOperator) TriggerOperator {
		if v != nil {
			return *v
		}
		var ret TriggerOperator
		return ret
	}).(TriggerOperatorOutput)
}

func (o TriggerOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TriggerOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TriggerOperatorInput interface {
	pulumi.Input

	ToTriggerOperatorOutput() TriggerOperatorOutput
	ToTriggerOperatorOutputWithContext(context.Context) TriggerOperatorOutput
}

var triggerOperatorPtrType = reflect.TypeOf((**TriggerOperator)(nil)).Elem()

type TriggerOperatorPtrInput interface {
	pulumi.Input

	ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput
	ToTriggerOperatorPtrOutputWithContext(context.Context) TriggerOperatorPtrOutput
}

type triggerOperatorPtr string

func TriggerOperatorPtr(v string) TriggerOperatorPtrInput {
	return (*triggerOperatorPtr)(&v)
}

func (*triggerOperatorPtr) ElementType() reflect.Type {
	return triggerOperatorPtrType
}

func (in *triggerOperatorPtr) ToTriggerOperatorPtrOutput() TriggerOperatorPtrOutput {
	return pulumi.ToOutput(in).(TriggerOperatorPtrOutput)
}

func (in *triggerOperatorPtr) ToTriggerOperatorPtrOutputWithContext(ctx context.Context) TriggerOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TriggerOperatorPtrOutput)
}

type TriggersOn string

const (
	// Trigger on Incidents
	TriggersOnIncidents = TriggersOn("Incidents")
)

func (TriggersOn) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggersOn)(nil)).Elem()
}

func (e TriggersOn) ToTriggersOnOutput() TriggersOnOutput {
	return pulumi.ToOutput(e).(TriggersOnOutput)
}

func (e TriggersOn) ToTriggersOnOutputWithContext(ctx context.Context) TriggersOnOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TriggersOnOutput)
}

func (e TriggersOn) ToTriggersOnPtrOutput() TriggersOnPtrOutput {
	return e.ToTriggersOnPtrOutputWithContext(context.Background())
}

func (e TriggersOn) ToTriggersOnPtrOutputWithContext(ctx context.Context) TriggersOnPtrOutput {
	return TriggersOn(e).ToTriggersOnOutputWithContext(ctx).ToTriggersOnPtrOutputWithContext(ctx)
}

func (e TriggersOn) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggersOn) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggersOn) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggersOn) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TriggersOnOutput struct{ *pulumi.OutputState }

func (TriggersOnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggersOn)(nil)).Elem()
}

func (o TriggersOnOutput) ToTriggersOnOutput() TriggersOnOutput {
	return o
}

func (o TriggersOnOutput) ToTriggersOnOutputWithContext(ctx context.Context) TriggersOnOutput {
	return o
}

func (o TriggersOnOutput) ToTriggersOnPtrOutput() TriggersOnPtrOutput {
	return o.ToTriggersOnPtrOutputWithContext(context.Background())
}

func (o TriggersOnOutput) ToTriggersOnPtrOutputWithContext(ctx context.Context) TriggersOnPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggersOn) *TriggersOn {
		return &v
	}).(TriggersOnPtrOutput)
}

func (o TriggersOnOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TriggersOnOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggersOn) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TriggersOnOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggersOnOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggersOn) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TriggersOnPtrOutput struct{ *pulumi.OutputState }

func (TriggersOnPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggersOn)(nil)).Elem()
}

func (o TriggersOnPtrOutput) ToTriggersOnPtrOutput() TriggersOnPtrOutput {
	return o
}

func (o TriggersOnPtrOutput) ToTriggersOnPtrOutputWithContext(ctx context.Context) TriggersOnPtrOutput {
	return o
}

func (o TriggersOnPtrOutput) Elem() TriggersOnOutput {
	return o.ApplyT(func(v *TriggersOn) TriggersOn {
		if v != nil {
			return *v
		}
		var ret TriggersOn
		return ret
	}).(TriggersOnOutput)
}

func (o TriggersOnPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggersOnPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TriggersOn) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TriggersOnInput interface {
	pulumi.Input

	ToTriggersOnOutput() TriggersOnOutput
	ToTriggersOnOutputWithContext(context.Context) TriggersOnOutput
}

var triggersOnPtrType = reflect.TypeOf((**TriggersOn)(nil)).Elem()

type TriggersOnPtrInput interface {
	pulumi.Input

	ToTriggersOnPtrOutput() TriggersOnPtrOutput
	ToTriggersOnPtrOutputWithContext(context.Context) TriggersOnPtrOutput
}

type triggersOnPtr string

func TriggersOnPtr(v string) TriggersOnPtrInput {
	return (*triggersOnPtr)(&v)
}

func (*triggersOnPtr) ElementType() reflect.Type {
	return triggersOnPtrType
}

func (in *triggersOnPtr) ToTriggersOnPtrOutput() TriggersOnPtrOutput {
	return pulumi.ToOutput(in).(TriggersOnPtrOutput)
}

func (in *triggersOnPtr) ToTriggersOnPtrOutputWithContext(ctx context.Context) TriggersOnPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TriggersOnPtrOutput)
}

type TriggersWhen string

const (
	// Trigger on created objects
	TriggersWhenCreated = TriggersWhen("Created")
)

func (TriggersWhen) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggersWhen)(nil)).Elem()
}

func (e TriggersWhen) ToTriggersWhenOutput() TriggersWhenOutput {
	return pulumi.ToOutput(e).(TriggersWhenOutput)
}

func (e TriggersWhen) ToTriggersWhenOutputWithContext(ctx context.Context) TriggersWhenOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TriggersWhenOutput)
}

func (e TriggersWhen) ToTriggersWhenPtrOutput() TriggersWhenPtrOutput {
	return e.ToTriggersWhenPtrOutputWithContext(context.Background())
}

func (e TriggersWhen) ToTriggersWhenPtrOutputWithContext(ctx context.Context) TriggersWhenPtrOutput {
	return TriggersWhen(e).ToTriggersWhenOutputWithContext(ctx).ToTriggersWhenPtrOutputWithContext(ctx)
}

func (e TriggersWhen) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggersWhen) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggersWhen) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggersWhen) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TriggersWhenOutput struct{ *pulumi.OutputState }

func (TriggersWhenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggersWhen)(nil)).Elem()
}

func (o TriggersWhenOutput) ToTriggersWhenOutput() TriggersWhenOutput {
	return o
}

func (o TriggersWhenOutput) ToTriggersWhenOutputWithContext(ctx context.Context) TriggersWhenOutput {
	return o
}

func (o TriggersWhenOutput) ToTriggersWhenPtrOutput() TriggersWhenPtrOutput {
	return o.ToTriggersWhenPtrOutputWithContext(context.Background())
}

func (o TriggersWhenOutput) ToTriggersWhenPtrOutputWithContext(ctx context.Context) TriggersWhenPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggersWhen) *TriggersWhen {
		return &v
	}).(TriggersWhenPtrOutput)
}

func (o TriggersWhenOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TriggersWhenOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggersWhen) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TriggersWhenOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggersWhenOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggersWhen) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TriggersWhenPtrOutput struct{ *pulumi.OutputState }

func (TriggersWhenPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggersWhen)(nil)).Elem()
}

func (o TriggersWhenPtrOutput) ToTriggersWhenPtrOutput() TriggersWhenPtrOutput {
	return o
}

func (o TriggersWhenPtrOutput) ToTriggersWhenPtrOutputWithContext(ctx context.Context) TriggersWhenPtrOutput {
	return o
}

func (o TriggersWhenPtrOutput) Elem() TriggersWhenOutput {
	return o.ApplyT(func(v *TriggersWhen) TriggersWhen {
		if v != nil {
			return *v
		}
		var ret TriggersWhen
		return ret
	}).(TriggersWhenOutput)
}

func (o TriggersWhenPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggersWhenPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TriggersWhen) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TriggersWhenInput interface {
	pulumi.Input

	ToTriggersWhenOutput() TriggersWhenOutput
	ToTriggersWhenOutputWithContext(context.Context) TriggersWhenOutput
}

var triggersWhenPtrType = reflect.TypeOf((**TriggersWhen)(nil)).Elem()

type TriggersWhenPtrInput interface {
	pulumi.Input

	ToTriggersWhenPtrOutput() TriggersWhenPtrOutput
	ToTriggersWhenPtrOutputWithContext(context.Context) TriggersWhenPtrOutput
}

type triggersWhenPtr string

func TriggersWhenPtr(v string) TriggersWhenPtrInput {
	return (*triggersWhenPtr)(&v)
}

func (*triggersWhenPtr) ElementType() reflect.Type {
	return triggersWhenPtrType
}

func (in *triggersWhenPtr) ToTriggersWhenPtrOutput() TriggersWhenPtrOutput {
	return pulumi.ToOutput(in).(TriggersWhenPtrOutput)
}

func (in *triggersWhenPtr) ToTriggersWhenPtrOutputWithContext(ctx context.Context) TriggersWhenPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TriggersWhenPtrOutput)
}

type UebaDataSources string

const (
	UebaDataSourcesAuditLogs     = UebaDataSources("AuditLogs")
	UebaDataSourcesAzureActivity = UebaDataSources("AzureActivity")
	UebaDataSourcesSecurityEvent = UebaDataSources("SecurityEvent")
	UebaDataSourcesSigninLogs    = UebaDataSources("SigninLogs")
)

func (UebaDataSources) ElementType() reflect.Type {
	return reflect.TypeOf((*UebaDataSources)(nil)).Elem()
}

func (e UebaDataSources) ToUebaDataSourcesOutput() UebaDataSourcesOutput {
	return pulumi.ToOutput(e).(UebaDataSourcesOutput)
}

func (e UebaDataSources) ToUebaDataSourcesOutputWithContext(ctx context.Context) UebaDataSourcesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UebaDataSourcesOutput)
}

func (e UebaDataSources) ToUebaDataSourcesPtrOutput() UebaDataSourcesPtrOutput {
	return e.ToUebaDataSourcesPtrOutputWithContext(context.Background())
}

func (e UebaDataSources) ToUebaDataSourcesPtrOutputWithContext(ctx context.Context) UebaDataSourcesPtrOutput {
	return UebaDataSources(e).ToUebaDataSourcesOutputWithContext(ctx).ToUebaDataSourcesPtrOutputWithContext(ctx)
}

func (e UebaDataSources) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UebaDataSources) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UebaDataSources) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UebaDataSources) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UebaDataSourcesOutput struct{ *pulumi.OutputState }

func (UebaDataSourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UebaDataSources)(nil)).Elem()
}

func (o UebaDataSourcesOutput) ToUebaDataSourcesOutput() UebaDataSourcesOutput {
	return o
}

func (o UebaDataSourcesOutput) ToUebaDataSourcesOutputWithContext(ctx context.Context) UebaDataSourcesOutput {
	return o
}

func (o UebaDataSourcesOutput) ToUebaDataSourcesPtrOutput() UebaDataSourcesPtrOutput {
	return o.ToUebaDataSourcesPtrOutputWithContext(context.Background())
}

func (o UebaDataSourcesOutput) ToUebaDataSourcesPtrOutputWithContext(ctx context.Context) UebaDataSourcesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UebaDataSources) *UebaDataSources {
		return &v
	}).(UebaDataSourcesPtrOutput)
}

func (o UebaDataSourcesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UebaDataSourcesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UebaDataSources) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UebaDataSourcesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UebaDataSourcesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UebaDataSources) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UebaDataSourcesPtrOutput struct{ *pulumi.OutputState }

func (UebaDataSourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UebaDataSources)(nil)).Elem()
}

func (o UebaDataSourcesPtrOutput) ToUebaDataSourcesPtrOutput() UebaDataSourcesPtrOutput {
	return o
}

func (o UebaDataSourcesPtrOutput) ToUebaDataSourcesPtrOutputWithContext(ctx context.Context) UebaDataSourcesPtrOutput {
	return o
}

func (o UebaDataSourcesPtrOutput) Elem() UebaDataSourcesOutput {
	return o.ApplyT(func(v *UebaDataSources) UebaDataSources {
		if v != nil {
			return *v
		}
		var ret UebaDataSources
		return ret
	}).(UebaDataSourcesOutput)
}

func (o UebaDataSourcesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UebaDataSourcesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UebaDataSources) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UebaDataSourcesInput interface {
	pulumi.Input

	ToUebaDataSourcesOutput() UebaDataSourcesOutput
	ToUebaDataSourcesOutputWithContext(context.Context) UebaDataSourcesOutput
}

var uebaDataSourcesPtrType = reflect.TypeOf((**UebaDataSources)(nil)).Elem()

type UebaDataSourcesPtrInput interface {
	pulumi.Input

	ToUebaDataSourcesPtrOutput() UebaDataSourcesPtrOutput
	ToUebaDataSourcesPtrOutputWithContext(context.Context) UebaDataSourcesPtrOutput
}

type uebaDataSourcesPtr string

func UebaDataSourcesPtr(v string) UebaDataSourcesPtrInput {
	return (*uebaDataSourcesPtr)(&v)
}

func (*uebaDataSourcesPtr) ElementType() reflect.Type {
	return uebaDataSourcesPtrType
}

func (in *uebaDataSourcesPtr) ToUebaDataSourcesPtrOutput() UebaDataSourcesPtrOutput {
	return pulumi.ToOutput(in).(UebaDataSourcesPtrOutput)
}

func (in *uebaDataSourcesPtr) ToUebaDataSourcesPtrOutputWithContext(ctx context.Context) UebaDataSourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UebaDataSourcesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertRuleKindOutput{})
	pulumi.RegisterOutputType(AlertRuleKindPtrOutput{})
	pulumi.RegisterOutputType(AlertSeverityOutput{})
	pulumi.RegisterOutputType(AlertSeverityPtrOutput{})
	pulumi.RegisterOutputType(AttackTacticOutput{})
	pulumi.RegisterOutputType(AttackTacticPtrOutput{})
	pulumi.RegisterOutputType(AutomationRuleActionTypeOutput{})
	pulumi.RegisterOutputType(AutomationRuleActionTypePtrOutput{})
	pulumi.RegisterOutputType(AutomationRuleConditionTypeOutput{})
	pulumi.RegisterOutputType(AutomationRuleConditionTypePtrOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyConditionSupportedOperatorOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyConditionSupportedOperatorPtrOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyConditionSupportedPropertyOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyConditionSupportedPropertyPtrOutput{})
	pulumi.RegisterOutputType(ContentTypeOutput{})
	pulumi.RegisterOutputType(ContentTypePtrOutput{})
	pulumi.RegisterOutputType(CustomEntityQueryKindOutput{})
	pulumi.RegisterOutputType(CustomEntityQueryKindPtrOutput{})
	pulumi.RegisterOutputType(DataConnectorKindOutput{})
	pulumi.RegisterOutputType(DataConnectorKindPtrOutput{})
	pulumi.RegisterOutputType(DataTypeStateOutput{})
	pulumi.RegisterOutputType(DataTypeStatePtrOutput{})
	pulumi.RegisterOutputType(EntityTimelineKindOutput{})
	pulumi.RegisterOutputType(EntityTimelineKindPtrOutput{})
	pulumi.RegisterOutputType(EntityTypeOutput{})
	pulumi.RegisterOutputType(EntityTypePtrOutput{})
	pulumi.RegisterOutputType(IncidentClassificationOutput{})
	pulumi.RegisterOutputType(IncidentClassificationPtrOutput{})
	pulumi.RegisterOutputType(IncidentClassificationReasonOutput{})
	pulumi.RegisterOutputType(IncidentClassificationReasonPtrOutput{})
	pulumi.RegisterOutputType(IncidentSeverityOutput{})
	pulumi.RegisterOutputType(IncidentSeverityPtrOutput{})
	pulumi.RegisterOutputType(IncidentStatusOutput{})
	pulumi.RegisterOutputType(IncidentStatusPtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(MicrosoftSecurityProductNameOutput{})
	pulumi.RegisterOutputType(MicrosoftSecurityProductNamePtrOutput{})
	pulumi.RegisterOutputType(OperatorOutput{})
	pulumi.RegisterOutputType(OperatorPtrOutput{})
	pulumi.RegisterOutputType(RepoTypeOutput{})
	pulumi.RegisterOutputType(RepoTypePtrOutput{})
	pulumi.RegisterOutputType(SettingKindOutput{})
	pulumi.RegisterOutputType(SettingKindPtrOutput{})
	pulumi.RegisterOutputType(SourceOutput{})
	pulumi.RegisterOutputType(SourcePtrOutput{})
	pulumi.RegisterOutputType(SourceKindOutput{})
	pulumi.RegisterOutputType(SourceKindPtrOutput{})
	pulumi.RegisterOutputType(SupportTierOutput{})
	pulumi.RegisterOutputType(SupportTierPtrOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceResourceKindOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceResourceKindPtrOutput{})
	pulumi.RegisterOutputType(TriggerOperatorOutput{})
	pulumi.RegisterOutputType(TriggerOperatorPtrOutput{})
	pulumi.RegisterOutputType(TriggersOnOutput{})
	pulumi.RegisterOutputType(TriggersOnPtrOutput{})
	pulumi.RegisterOutputType(TriggersWhenOutput{})
	pulumi.RegisterOutputType(TriggersWhenPtrOutput{})
	pulumi.RegisterOutputType(UebaDataSourcesOutput{})
	pulumi.RegisterOutputType(UebaDataSourcesPtrOutput{})
}
