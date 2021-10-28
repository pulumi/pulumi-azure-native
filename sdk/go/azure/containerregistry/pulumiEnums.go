


package containerregistry

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

type ActionsRequired string

const (
	ActionsRequiredNone     = ActionsRequired("None")
	ActionsRequiredRecreate = ActionsRequired("Recreate")
)

func (ActionsRequired) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionsRequired)(nil)).Elem()
}

func (e ActionsRequired) ToActionsRequiredOutput() ActionsRequiredOutput {
	return pulumi.ToOutput(e).(ActionsRequiredOutput)
}

func (e ActionsRequired) ToActionsRequiredOutputWithContext(ctx context.Context) ActionsRequiredOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionsRequiredOutput)
}

func (e ActionsRequired) ToActionsRequiredPtrOutput() ActionsRequiredPtrOutput {
	return e.ToActionsRequiredPtrOutputWithContext(context.Background())
}

func (e ActionsRequired) ToActionsRequiredPtrOutputWithContext(ctx context.Context) ActionsRequiredPtrOutput {
	return ActionsRequired(e).ToActionsRequiredOutputWithContext(ctx).ToActionsRequiredPtrOutputWithContext(ctx)
}

func (e ActionsRequired) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionsRequired) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionsRequired) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ActionsRequired) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionsRequiredOutput struct{ *pulumi.OutputState }

func (ActionsRequiredOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionsRequired)(nil)).Elem()
}

func (o ActionsRequiredOutput) ToActionsRequiredOutput() ActionsRequiredOutput {
	return o
}

func (o ActionsRequiredOutput) ToActionsRequiredOutputWithContext(ctx context.Context) ActionsRequiredOutput {
	return o
}

func (o ActionsRequiredOutput) ToActionsRequiredPtrOutput() ActionsRequiredPtrOutput {
	return o.ToActionsRequiredPtrOutputWithContext(context.Background())
}

func (o ActionsRequiredOutput) ToActionsRequiredPtrOutputWithContext(ctx context.Context) ActionsRequiredPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionsRequired) *ActionsRequired {
		return &v
	}).(ActionsRequiredPtrOutput)
}

func (o ActionsRequiredOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionsRequiredOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionsRequired) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionsRequiredOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionsRequiredOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionsRequired) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionsRequiredPtrOutput struct{ *pulumi.OutputState }

func (ActionsRequiredPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsRequired)(nil)).Elem()
}

func (o ActionsRequiredPtrOutput) ToActionsRequiredPtrOutput() ActionsRequiredPtrOutput {
	return o
}

func (o ActionsRequiredPtrOutput) ToActionsRequiredPtrOutputWithContext(ctx context.Context) ActionsRequiredPtrOutput {
	return o
}

func (o ActionsRequiredPtrOutput) Elem() ActionsRequiredOutput {
	return o.ApplyT(func(v *ActionsRequired) ActionsRequired {
		if v != nil {
			return *v
		}
		var ret ActionsRequired
		return ret
	}).(ActionsRequiredOutput)
}

func (o ActionsRequiredPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionsRequiredPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ActionsRequired) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ActionsRequiredInput interface {
	pulumi.Input

	ToActionsRequiredOutput() ActionsRequiredOutput
	ToActionsRequiredOutputWithContext(context.Context) ActionsRequiredOutput
}

var actionsRequiredPtrType = reflect.TypeOf((**ActionsRequired)(nil)).Elem()

type ActionsRequiredPtrInput interface {
	pulumi.Input

	ToActionsRequiredPtrOutput() ActionsRequiredPtrOutput
	ToActionsRequiredPtrOutputWithContext(context.Context) ActionsRequiredPtrOutput
}

type actionsRequiredPtr string

func ActionsRequiredPtr(v string) ActionsRequiredPtrInput {
	return (*actionsRequiredPtr)(&v)
}

func (*actionsRequiredPtr) ElementType() reflect.Type {
	return actionsRequiredPtrType
}

func (in *actionsRequiredPtr) ToActionsRequiredPtrOutput() ActionsRequiredPtrOutput {
	return pulumi.ToOutput(in).(ActionsRequiredPtrOutput)
}

func (in *actionsRequiredPtr) ToActionsRequiredPtrOutputWithContext(ctx context.Context) ActionsRequiredPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionsRequiredPtrOutput)
}

type Architecture string

const (
	ArchitectureAmd64 = Architecture("amd64")
	ArchitectureX86   = Architecture("x86")
	Architecture_386  = Architecture("386")
	ArchitectureArm   = Architecture("arm")
	ArchitectureArm64 = Architecture("arm64")
)

func (Architecture) ElementType() reflect.Type {
	return reflect.TypeOf((*Architecture)(nil)).Elem()
}

func (e Architecture) ToArchitectureOutput() ArchitectureOutput {
	return pulumi.ToOutput(e).(ArchitectureOutput)
}

func (e Architecture) ToArchitectureOutputWithContext(ctx context.Context) ArchitectureOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ArchitectureOutput)
}

func (e Architecture) ToArchitecturePtrOutput() ArchitecturePtrOutput {
	return e.ToArchitecturePtrOutputWithContext(context.Background())
}

func (e Architecture) ToArchitecturePtrOutputWithContext(ctx context.Context) ArchitecturePtrOutput {
	return Architecture(e).ToArchitectureOutputWithContext(ctx).ToArchitecturePtrOutputWithContext(ctx)
}

func (e Architecture) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Architecture) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Architecture) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Architecture) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ArchitectureOutput struct{ *pulumi.OutputState }

func (ArchitectureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Architecture)(nil)).Elem()
}

func (o ArchitectureOutput) ToArchitectureOutput() ArchitectureOutput {
	return o
}

func (o ArchitectureOutput) ToArchitectureOutputWithContext(ctx context.Context) ArchitectureOutput {
	return o
}

func (o ArchitectureOutput) ToArchitecturePtrOutput() ArchitecturePtrOutput {
	return o.ToArchitecturePtrOutputWithContext(context.Background())
}

func (o ArchitectureOutput) ToArchitecturePtrOutputWithContext(ctx context.Context) ArchitecturePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Architecture) *Architecture {
		return &v
	}).(ArchitecturePtrOutput)
}

func (o ArchitectureOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ArchitectureOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Architecture) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ArchitectureOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ArchitectureOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Architecture) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ArchitecturePtrOutput struct{ *pulumi.OutputState }

func (ArchitecturePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Architecture)(nil)).Elem()
}

func (o ArchitecturePtrOutput) ToArchitecturePtrOutput() ArchitecturePtrOutput {
	return o
}

func (o ArchitecturePtrOutput) ToArchitecturePtrOutputWithContext(ctx context.Context) ArchitecturePtrOutput {
	return o
}

func (o ArchitecturePtrOutput) Elem() ArchitectureOutput {
	return o.ApplyT(func(v *Architecture) Architecture {
		if v != nil {
			return *v
		}
		var ret Architecture
		return ret
	}).(ArchitectureOutput)
}

func (o ArchitecturePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ArchitecturePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Architecture) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ArchitectureInput interface {
	pulumi.Input

	ToArchitectureOutput() ArchitectureOutput
	ToArchitectureOutputWithContext(context.Context) ArchitectureOutput
}

var architecturePtrType = reflect.TypeOf((**Architecture)(nil)).Elem()

type ArchitecturePtrInput interface {
	pulumi.Input

	ToArchitecturePtrOutput() ArchitecturePtrOutput
	ToArchitecturePtrOutputWithContext(context.Context) ArchitecturePtrOutput
}

type architecturePtr string

func ArchitecturePtr(v string) ArchitecturePtrInput {
	return (*architecturePtr)(&v)
}

func (*architecturePtr) ElementType() reflect.Type {
	return architecturePtrType
}

func (in *architecturePtr) ToArchitecturePtrOutput() ArchitecturePtrOutput {
	return pulumi.ToOutput(in).(ArchitecturePtrOutput)
}

func (in *architecturePtr) ToArchitecturePtrOutputWithContext(ctx context.Context) ArchitecturePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ArchitecturePtrOutput)
}

type AuditLogStatus string

const (
	AuditLogStatusEnabled  = AuditLogStatus("Enabled")
	AuditLogStatusDisabled = AuditLogStatus("Disabled")
)

func (AuditLogStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogStatus)(nil)).Elem()
}

func (e AuditLogStatus) ToAuditLogStatusOutput() AuditLogStatusOutput {
	return pulumi.ToOutput(e).(AuditLogStatusOutput)
}

func (e AuditLogStatus) ToAuditLogStatusOutputWithContext(ctx context.Context) AuditLogStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuditLogStatusOutput)
}

func (e AuditLogStatus) ToAuditLogStatusPtrOutput() AuditLogStatusPtrOutput {
	return e.ToAuditLogStatusPtrOutputWithContext(context.Background())
}

func (e AuditLogStatus) ToAuditLogStatusPtrOutputWithContext(ctx context.Context) AuditLogStatusPtrOutput {
	return AuditLogStatus(e).ToAuditLogStatusOutputWithContext(ctx).ToAuditLogStatusPtrOutputWithContext(ctx)
}

func (e AuditLogStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuditLogStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuditLogStatusOutput struct{ *pulumi.OutputState }

func (AuditLogStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogStatus)(nil)).Elem()
}

func (o AuditLogStatusOutput) ToAuditLogStatusOutput() AuditLogStatusOutput {
	return o
}

func (o AuditLogStatusOutput) ToAuditLogStatusOutputWithContext(ctx context.Context) AuditLogStatusOutput {
	return o
}

func (o AuditLogStatusOutput) ToAuditLogStatusPtrOutput() AuditLogStatusPtrOutput {
	return o.ToAuditLogStatusPtrOutputWithContext(context.Background())
}

func (o AuditLogStatusOutput) ToAuditLogStatusPtrOutputWithContext(ctx context.Context) AuditLogStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuditLogStatus) *AuditLogStatus {
		return &v
	}).(AuditLogStatusPtrOutput)
}

func (o AuditLogStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuditLogStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuditLogStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuditLogStatusPtrOutput struct{ *pulumi.OutputState }

func (AuditLogStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogStatus)(nil)).Elem()
}

func (o AuditLogStatusPtrOutput) ToAuditLogStatusPtrOutput() AuditLogStatusPtrOutput {
	return o
}

func (o AuditLogStatusPtrOutput) ToAuditLogStatusPtrOutputWithContext(ctx context.Context) AuditLogStatusPtrOutput {
	return o
}

func (o AuditLogStatusPtrOutput) Elem() AuditLogStatusOutput {
	return o.ApplyT(func(v *AuditLogStatus) AuditLogStatus {
		if v != nil {
			return *v
		}
		var ret AuditLogStatus
		return ret
	}).(AuditLogStatusOutput)
}

func (o AuditLogStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuditLogStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AuditLogStatusInput interface {
	pulumi.Input

	ToAuditLogStatusOutput() AuditLogStatusOutput
	ToAuditLogStatusOutputWithContext(context.Context) AuditLogStatusOutput
}

var auditLogStatusPtrType = reflect.TypeOf((**AuditLogStatus)(nil)).Elem()

type AuditLogStatusPtrInput interface {
	pulumi.Input

	ToAuditLogStatusPtrOutput() AuditLogStatusPtrOutput
	ToAuditLogStatusPtrOutputWithContext(context.Context) AuditLogStatusPtrOutput
}

type auditLogStatusPtr string

func AuditLogStatusPtr(v string) AuditLogStatusPtrInput {
	return (*auditLogStatusPtr)(&v)
}

func (*auditLogStatusPtr) ElementType() reflect.Type {
	return auditLogStatusPtrType
}

func (in *auditLogStatusPtr) ToAuditLogStatusPtrOutput() AuditLogStatusPtrOutput {
	return pulumi.ToOutput(in).(AuditLogStatusPtrOutput)
}

func (in *auditLogStatusPtr) ToAuditLogStatusPtrOutputWithContext(ctx context.Context) AuditLogStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuditLogStatusPtrOutput)
}

type BaseImageTriggerType string

const (
	BaseImageTriggerTypeAll     = BaseImageTriggerType("All")
	BaseImageTriggerTypeRuntime = BaseImageTriggerType("Runtime")
)

func (BaseImageTriggerType) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageTriggerType)(nil)).Elem()
}

func (e BaseImageTriggerType) ToBaseImageTriggerTypeOutput() BaseImageTriggerTypeOutput {
	return pulumi.ToOutput(e).(BaseImageTriggerTypeOutput)
}

func (e BaseImageTriggerType) ToBaseImageTriggerTypeOutputWithContext(ctx context.Context) BaseImageTriggerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BaseImageTriggerTypeOutput)
}

func (e BaseImageTriggerType) ToBaseImageTriggerTypePtrOutput() BaseImageTriggerTypePtrOutput {
	return e.ToBaseImageTriggerTypePtrOutputWithContext(context.Background())
}

func (e BaseImageTriggerType) ToBaseImageTriggerTypePtrOutputWithContext(ctx context.Context) BaseImageTriggerTypePtrOutput {
	return BaseImageTriggerType(e).ToBaseImageTriggerTypeOutputWithContext(ctx).ToBaseImageTriggerTypePtrOutputWithContext(ctx)
}

func (e BaseImageTriggerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BaseImageTriggerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BaseImageTriggerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BaseImageTriggerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BaseImageTriggerTypeOutput struct{ *pulumi.OutputState }

func (BaseImageTriggerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaseImageTriggerType)(nil)).Elem()
}

func (o BaseImageTriggerTypeOutput) ToBaseImageTriggerTypeOutput() BaseImageTriggerTypeOutput {
	return o
}

func (o BaseImageTriggerTypeOutput) ToBaseImageTriggerTypeOutputWithContext(ctx context.Context) BaseImageTriggerTypeOutput {
	return o
}

func (o BaseImageTriggerTypeOutput) ToBaseImageTriggerTypePtrOutput() BaseImageTriggerTypePtrOutput {
	return o.ToBaseImageTriggerTypePtrOutputWithContext(context.Background())
}

func (o BaseImageTriggerTypeOutput) ToBaseImageTriggerTypePtrOutputWithContext(ctx context.Context) BaseImageTriggerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BaseImageTriggerType) *BaseImageTriggerType {
		return &v
	}).(BaseImageTriggerTypePtrOutput)
}

func (o BaseImageTriggerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BaseImageTriggerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BaseImageTriggerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BaseImageTriggerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BaseImageTriggerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BaseImageTriggerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BaseImageTriggerTypePtrOutput struct{ *pulumi.OutputState }

func (BaseImageTriggerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BaseImageTriggerType)(nil)).Elem()
}

func (o BaseImageTriggerTypePtrOutput) ToBaseImageTriggerTypePtrOutput() BaseImageTriggerTypePtrOutput {
	return o
}

func (o BaseImageTriggerTypePtrOutput) ToBaseImageTriggerTypePtrOutputWithContext(ctx context.Context) BaseImageTriggerTypePtrOutput {
	return o
}

func (o BaseImageTriggerTypePtrOutput) Elem() BaseImageTriggerTypeOutput {
	return o.ApplyT(func(v *BaseImageTriggerType) BaseImageTriggerType {
		if v != nil {
			return *v
		}
		var ret BaseImageTriggerType
		return ret
	}).(BaseImageTriggerTypeOutput)
}

func (o BaseImageTriggerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BaseImageTriggerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BaseImageTriggerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BaseImageTriggerTypeInput interface {
	pulumi.Input

	ToBaseImageTriggerTypeOutput() BaseImageTriggerTypeOutput
	ToBaseImageTriggerTypeOutputWithContext(context.Context) BaseImageTriggerTypeOutput
}

var baseImageTriggerTypePtrType = reflect.TypeOf((**BaseImageTriggerType)(nil)).Elem()

type BaseImageTriggerTypePtrInput interface {
	pulumi.Input

	ToBaseImageTriggerTypePtrOutput() BaseImageTriggerTypePtrOutput
	ToBaseImageTriggerTypePtrOutputWithContext(context.Context) BaseImageTriggerTypePtrOutput
}

type baseImageTriggerTypePtr string

func BaseImageTriggerTypePtr(v string) BaseImageTriggerTypePtrInput {
	return (*baseImageTriggerTypePtr)(&v)
}

func (*baseImageTriggerTypePtr) ElementType() reflect.Type {
	return baseImageTriggerTypePtrType
}

func (in *baseImageTriggerTypePtr) ToBaseImageTriggerTypePtrOutput() BaseImageTriggerTypePtrOutput {
	return pulumi.ToOutput(in).(BaseImageTriggerTypePtrOutput)
}

func (in *baseImageTriggerTypePtr) ToBaseImageTriggerTypePtrOutputWithContext(ctx context.Context) BaseImageTriggerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BaseImageTriggerTypePtrOutput)
}

type ConnectedRegistryMode string

const (
	ConnectedRegistryModeRegistry = ConnectedRegistryMode("Registry")
	ConnectedRegistryModeMirror   = ConnectedRegistryMode("Mirror")
)

func (ConnectedRegistryMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedRegistryMode)(nil)).Elem()
}

func (e ConnectedRegistryMode) ToConnectedRegistryModeOutput() ConnectedRegistryModeOutput {
	return pulumi.ToOutput(e).(ConnectedRegistryModeOutput)
}

func (e ConnectedRegistryMode) ToConnectedRegistryModeOutputWithContext(ctx context.Context) ConnectedRegistryModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectedRegistryModeOutput)
}

func (e ConnectedRegistryMode) ToConnectedRegistryModePtrOutput() ConnectedRegistryModePtrOutput {
	return e.ToConnectedRegistryModePtrOutputWithContext(context.Background())
}

func (e ConnectedRegistryMode) ToConnectedRegistryModePtrOutputWithContext(ctx context.Context) ConnectedRegistryModePtrOutput {
	return ConnectedRegistryMode(e).ToConnectedRegistryModeOutputWithContext(ctx).ToConnectedRegistryModePtrOutputWithContext(ctx)
}

func (e ConnectedRegistryMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectedRegistryMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectedRegistryMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectedRegistryMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectedRegistryModeOutput struct{ *pulumi.OutputState }

func (ConnectedRegistryModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectedRegistryMode)(nil)).Elem()
}

func (o ConnectedRegistryModeOutput) ToConnectedRegistryModeOutput() ConnectedRegistryModeOutput {
	return o
}

func (o ConnectedRegistryModeOutput) ToConnectedRegistryModeOutputWithContext(ctx context.Context) ConnectedRegistryModeOutput {
	return o
}

func (o ConnectedRegistryModeOutput) ToConnectedRegistryModePtrOutput() ConnectedRegistryModePtrOutput {
	return o.ToConnectedRegistryModePtrOutputWithContext(context.Background())
}

func (o ConnectedRegistryModeOutput) ToConnectedRegistryModePtrOutputWithContext(ctx context.Context) ConnectedRegistryModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectedRegistryMode) *ConnectedRegistryMode {
		return &v
	}).(ConnectedRegistryModePtrOutput)
}

func (o ConnectedRegistryModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectedRegistryModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectedRegistryMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectedRegistryModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectedRegistryModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectedRegistryMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectedRegistryModePtrOutput struct{ *pulumi.OutputState }

func (ConnectedRegistryModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedRegistryMode)(nil)).Elem()
}

func (o ConnectedRegistryModePtrOutput) ToConnectedRegistryModePtrOutput() ConnectedRegistryModePtrOutput {
	return o
}

func (o ConnectedRegistryModePtrOutput) ToConnectedRegistryModePtrOutputWithContext(ctx context.Context) ConnectedRegistryModePtrOutput {
	return o
}

func (o ConnectedRegistryModePtrOutput) Elem() ConnectedRegistryModeOutput {
	return o.ApplyT(func(v *ConnectedRegistryMode) ConnectedRegistryMode {
		if v != nil {
			return *v
		}
		var ret ConnectedRegistryMode
		return ret
	}).(ConnectedRegistryModeOutput)
}

func (o ConnectedRegistryModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectedRegistryModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectedRegistryMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectedRegistryModeInput interface {
	pulumi.Input

	ToConnectedRegistryModeOutput() ConnectedRegistryModeOutput
	ToConnectedRegistryModeOutputWithContext(context.Context) ConnectedRegistryModeOutput
}

var connectedRegistryModePtrType = reflect.TypeOf((**ConnectedRegistryMode)(nil)).Elem()

type ConnectedRegistryModePtrInput interface {
	pulumi.Input

	ToConnectedRegistryModePtrOutput() ConnectedRegistryModePtrOutput
	ToConnectedRegistryModePtrOutputWithContext(context.Context) ConnectedRegistryModePtrOutput
}

type connectedRegistryModePtr string

func ConnectedRegistryModePtr(v string) ConnectedRegistryModePtrInput {
	return (*connectedRegistryModePtr)(&v)
}

func (*connectedRegistryModePtr) ElementType() reflect.Type {
	return connectedRegistryModePtrType
}

func (in *connectedRegistryModePtr) ToConnectedRegistryModePtrOutput() ConnectedRegistryModePtrOutput {
	return pulumi.ToOutput(in).(ConnectedRegistryModePtrOutput)
}

func (in *connectedRegistryModePtr) ToConnectedRegistryModePtrOutputWithContext(ctx context.Context) ConnectedRegistryModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectedRegistryModePtrOutput)
}

type ConnectionStatus string

const (
	ConnectionStatusApproved     = ConnectionStatus("Approved")
	ConnectionStatusPending      = ConnectionStatus("Pending")
	ConnectionStatusRejected     = ConnectionStatus("Rejected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
)

func (ConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatus)(nil)).Elem()
}

func (e ConnectionStatus) ToConnectionStatusOutput() ConnectionStatusOutput {
	return pulumi.ToOutput(e).(ConnectionStatusOutput)
}

func (e ConnectionStatus) ToConnectionStatusOutputWithContext(ctx context.Context) ConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionStatusOutput)
}

func (e ConnectionStatus) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return e.ToConnectionStatusPtrOutputWithContext(context.Background())
}

func (e ConnectionStatus) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return ConnectionStatus(e).ToConnectionStatusOutputWithContext(ctx).ToConnectionStatusPtrOutputWithContext(ctx)
}

func (e ConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionStatusOutput struct{ *pulumi.OutputState }

func (ConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStatus)(nil)).Elem()
}

func (o ConnectionStatusOutput) ToConnectionStatusOutput() ConnectionStatusOutput {
	return o
}

func (o ConnectionStatusOutput) ToConnectionStatusOutputWithContext(ctx context.Context) ConnectionStatusOutput {
	return o
}

func (o ConnectionStatusOutput) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return o.ToConnectionStatusPtrOutputWithContext(context.Background())
}

func (o ConnectionStatusOutput) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionStatus) *ConnectionStatus {
		return &v
	}).(ConnectionStatusPtrOutput)
}

func (o ConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStatus)(nil)).Elem()
}

func (o ConnectionStatusPtrOutput) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return o
}

func (o ConnectionStatusPtrOutput) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return o
}

func (o ConnectionStatusPtrOutput) Elem() ConnectionStatusOutput {
	return o.ApplyT(func(v *ConnectionStatus) ConnectionStatus {
		if v != nil {
			return *v
		}
		var ret ConnectionStatus
		return ret
	}).(ConnectionStatusOutput)
}

func (o ConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionStatusInput interface {
	pulumi.Input

	ToConnectionStatusOutput() ConnectionStatusOutput
	ToConnectionStatusOutputWithContext(context.Context) ConnectionStatusOutput
}

var connectionStatusPtrType = reflect.TypeOf((**ConnectionStatus)(nil)).Elem()

type ConnectionStatusPtrInput interface {
	pulumi.Input

	ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput
	ToConnectionStatusPtrOutputWithContext(context.Context) ConnectionStatusPtrOutput
}

type connectionStatusPtr string

func ConnectionStatusPtr(v string) ConnectionStatusPtrInput {
	return (*connectionStatusPtr)(&v)
}

func (*connectionStatusPtr) ElementType() reflect.Type {
	return connectionStatusPtrType
}

func (in *connectionStatusPtr) ToConnectionStatusPtrOutput() ConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(ConnectionStatusPtrOutput)
}

func (in *connectionStatusPtr) ToConnectionStatusPtrOutputWithContext(ctx context.Context) ConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionStatusPtrOutput)
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

type LogLevel string

const (
	LogLevelDebug       = LogLevel("Debug")
	LogLevelInformation = LogLevel("Information")
	LogLevelWarning     = LogLevel("Warning")
	LogLevelError       = LogLevel("Error")
	LogLevelNone        = LogLevel("None")
)

func (LogLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*LogLevel)(nil)).Elem()
}

func (e LogLevel) ToLogLevelOutput() LogLevelOutput {
	return pulumi.ToOutput(e).(LogLevelOutput)
}

func (e LogLevel) ToLogLevelOutputWithContext(ctx context.Context) LogLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LogLevelOutput)
}

func (e LogLevel) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return e.ToLogLevelPtrOutputWithContext(context.Background())
}

func (e LogLevel) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return LogLevel(e).ToLogLevelOutputWithContext(ctx).ToLogLevelPtrOutputWithContext(ctx)
}

func (e LogLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LogLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LogLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LogLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LogLevelOutput struct{ *pulumi.OutputState }

func (LogLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogLevel)(nil)).Elem()
}

func (o LogLevelOutput) ToLogLevelOutput() LogLevelOutput {
	return o
}

func (o LogLevelOutput) ToLogLevelOutputWithContext(ctx context.Context) LogLevelOutput {
	return o
}

func (o LogLevelOutput) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return o.ToLogLevelPtrOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogLevel) *LogLevel {
		return &v
	}).(LogLevelPtrOutput)
}

func (o LogLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LogLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LogLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LogLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LogLevelPtrOutput struct{ *pulumi.OutputState }

func (LogLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogLevel)(nil)).Elem()
}

func (o LogLevelPtrOutput) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return o
}

func (o LogLevelPtrOutput) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return o
}

func (o LogLevelPtrOutput) Elem() LogLevelOutput {
	return o.ApplyT(func(v *LogLevel) LogLevel {
		if v != nil {
			return *v
		}
		var ret LogLevel
		return ret
	}).(LogLevelOutput)
}

func (o LogLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LogLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LogLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LogLevelInput interface {
	pulumi.Input

	ToLogLevelOutput() LogLevelOutput
	ToLogLevelOutputWithContext(context.Context) LogLevelOutput
}

var logLevelPtrType = reflect.TypeOf((**LogLevel)(nil)).Elem()

type LogLevelPtrInput interface {
	pulumi.Input

	ToLogLevelPtrOutput() LogLevelPtrOutput
	ToLogLevelPtrOutputWithContext(context.Context) LogLevelPtrOutput
}

type logLevelPtr string

func LogLevelPtr(v string) LogLevelPtrInput {
	return (*logLevelPtr)(&v)
}

func (*logLevelPtr) ElementType() reflect.Type {
	return logLevelPtrType
}

func (in *logLevelPtr) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return pulumi.ToOutput(in).(LogLevelPtrOutput)
}

func (in *logLevelPtr) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LogLevelPtrOutput)
}

type OS string

const (
	OSWindows = OS("Windows")
	OSLinux   = OS("Linux")
)

func (OS) ElementType() reflect.Type {
	return reflect.TypeOf((*OS)(nil)).Elem()
}

func (e OS) ToOSOutput() OSOutput {
	return pulumi.ToOutput(e).(OSOutput)
}

func (e OS) ToOSOutputWithContext(ctx context.Context) OSOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OSOutput)
}

func (e OS) ToOSPtrOutput() OSPtrOutput {
	return e.ToOSPtrOutputWithContext(context.Background())
}

func (e OS) ToOSPtrOutputWithContext(ctx context.Context) OSPtrOutput {
	return OS(e).ToOSOutputWithContext(ctx).ToOSPtrOutputWithContext(ctx)
}

func (e OS) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OS) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OS) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OS) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OSOutput struct{ *pulumi.OutputState }

func (OSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OS)(nil)).Elem()
}

func (o OSOutput) ToOSOutput() OSOutput {
	return o
}

func (o OSOutput) ToOSOutputWithContext(ctx context.Context) OSOutput {
	return o
}

func (o OSOutput) ToOSPtrOutput() OSPtrOutput {
	return o.ToOSPtrOutputWithContext(context.Background())
}

func (o OSOutput) ToOSPtrOutputWithContext(ctx context.Context) OSPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OS) *OS {
		return &v
	}).(OSPtrOutput)
}

func (o OSOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OSOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OS) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OSOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OS) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OSPtrOutput struct{ *pulumi.OutputState }

func (OSPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OS)(nil)).Elem()
}

func (o OSPtrOutput) ToOSPtrOutput() OSPtrOutput {
	return o
}

func (o OSPtrOutput) ToOSPtrOutputWithContext(ctx context.Context) OSPtrOutput {
	return o
}

func (o OSPtrOutput) Elem() OSOutput {
	return o.ApplyT(func(v *OS) OS {
		if v != nil {
			return *v
		}
		var ret OS
		return ret
	}).(OSOutput)
}

func (o OSPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OSPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OS) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OSInput interface {
	pulumi.Input

	ToOSOutput() OSOutput
	ToOSOutputWithContext(context.Context) OSOutput
}

var osPtrType = reflect.TypeOf((**OS)(nil)).Elem()

type OSPtrInput interface {
	pulumi.Input

	ToOSPtrOutput() OSPtrOutput
	ToOSPtrOutputWithContext(context.Context) OSPtrOutput
}

type osPtr string

func OSPtr(v string) OSPtrInput {
	return (*osPtr)(&v)
}

func (*osPtr) ElementType() reflect.Type {
	return osPtrType
}

func (in *osPtr) ToOSPtrOutput() OSPtrOutput {
	return pulumi.ToOutput(in).(OSPtrOutput)
}

func (in *osPtr) ToOSPtrOutputWithContext(ctx context.Context) OSPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OSPtrOutput)
}

type PipelineOptions string

const (
	PipelineOptionsOverwriteTags             = PipelineOptions("OverwriteTags")
	PipelineOptionsOverwriteBlobs            = PipelineOptions("OverwriteBlobs")
	PipelineOptionsDeleteSourceBlobOnSuccess = PipelineOptions("DeleteSourceBlobOnSuccess")
	PipelineOptionsContinueOnErrors          = PipelineOptions("ContinueOnErrors")
)

func (PipelineOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineOptions)(nil)).Elem()
}

func (e PipelineOptions) ToPipelineOptionsOutput() PipelineOptionsOutput {
	return pulumi.ToOutput(e).(PipelineOptionsOutput)
}

func (e PipelineOptions) ToPipelineOptionsOutputWithContext(ctx context.Context) PipelineOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PipelineOptionsOutput)
}

func (e PipelineOptions) ToPipelineOptionsPtrOutput() PipelineOptionsPtrOutput {
	return e.ToPipelineOptionsPtrOutputWithContext(context.Background())
}

func (e PipelineOptions) ToPipelineOptionsPtrOutputWithContext(ctx context.Context) PipelineOptionsPtrOutput {
	return PipelineOptions(e).ToPipelineOptionsOutputWithContext(ctx).ToPipelineOptionsPtrOutputWithContext(ctx)
}

func (e PipelineOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PipelineOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PipelineOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PipelineOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PipelineOptionsOutput struct{ *pulumi.OutputState }

func (PipelineOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineOptions)(nil)).Elem()
}

func (o PipelineOptionsOutput) ToPipelineOptionsOutput() PipelineOptionsOutput {
	return o
}

func (o PipelineOptionsOutput) ToPipelineOptionsOutputWithContext(ctx context.Context) PipelineOptionsOutput {
	return o
}

func (o PipelineOptionsOutput) ToPipelineOptionsPtrOutput() PipelineOptionsPtrOutput {
	return o.ToPipelineOptionsPtrOutputWithContext(context.Background())
}

func (o PipelineOptionsOutput) ToPipelineOptionsPtrOutputWithContext(ctx context.Context) PipelineOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineOptions) *PipelineOptions {
		return &v
	}).(PipelineOptionsPtrOutput)
}

func (o PipelineOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PipelineOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PipelineOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PipelineOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PipelineOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PipelineOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PipelineOptionsPtrOutput struct{ *pulumi.OutputState }

func (PipelineOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineOptions)(nil)).Elem()
}

func (o PipelineOptionsPtrOutput) ToPipelineOptionsPtrOutput() PipelineOptionsPtrOutput {
	return o
}

func (o PipelineOptionsPtrOutput) ToPipelineOptionsPtrOutputWithContext(ctx context.Context) PipelineOptionsPtrOutput {
	return o
}

func (o PipelineOptionsPtrOutput) Elem() PipelineOptionsOutput {
	return o.ApplyT(func(v *PipelineOptions) PipelineOptions {
		if v != nil {
			return *v
		}
		var ret PipelineOptions
		return ret
	}).(PipelineOptionsOutput)
}

func (o PipelineOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PipelineOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PipelineOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PipelineOptionsInput interface {
	pulumi.Input

	ToPipelineOptionsOutput() PipelineOptionsOutput
	ToPipelineOptionsOutputWithContext(context.Context) PipelineOptionsOutput
}

var pipelineOptionsPtrType = reflect.TypeOf((**PipelineOptions)(nil)).Elem()

type PipelineOptionsPtrInput interface {
	pulumi.Input

	ToPipelineOptionsPtrOutput() PipelineOptionsPtrOutput
	ToPipelineOptionsPtrOutputWithContext(context.Context) PipelineOptionsPtrOutput
}

type pipelineOptionsPtr string

func PipelineOptionsPtr(v string) PipelineOptionsPtrInput {
	return (*pipelineOptionsPtr)(&v)
}

func (*pipelineOptionsPtr) ElementType() reflect.Type {
	return pipelineOptionsPtrType
}

func (in *pipelineOptionsPtr) ToPipelineOptionsPtrOutput() PipelineOptionsPtrOutput {
	return pulumi.ToOutput(in).(PipelineOptionsPtrOutput)
}

func (in *pipelineOptionsPtr) ToPipelineOptionsPtrOutputWithContext(ctx context.Context) PipelineOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PipelineOptionsPtrOutput)
}

type PipelineRunSourceType string

const (
	PipelineRunSourceTypeAzureStorageBlob = PipelineRunSourceType("AzureStorageBlob")
)

func (PipelineRunSourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunSourceType)(nil)).Elem()
}

func (e PipelineRunSourceType) ToPipelineRunSourceTypeOutput() PipelineRunSourceTypeOutput {
	return pulumi.ToOutput(e).(PipelineRunSourceTypeOutput)
}

func (e PipelineRunSourceType) ToPipelineRunSourceTypeOutputWithContext(ctx context.Context) PipelineRunSourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PipelineRunSourceTypeOutput)
}

func (e PipelineRunSourceType) ToPipelineRunSourceTypePtrOutput() PipelineRunSourceTypePtrOutput {
	return e.ToPipelineRunSourceTypePtrOutputWithContext(context.Background())
}

func (e PipelineRunSourceType) ToPipelineRunSourceTypePtrOutputWithContext(ctx context.Context) PipelineRunSourceTypePtrOutput {
	return PipelineRunSourceType(e).ToPipelineRunSourceTypeOutputWithContext(ctx).ToPipelineRunSourceTypePtrOutputWithContext(ctx)
}

func (e PipelineRunSourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PipelineRunSourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PipelineRunSourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PipelineRunSourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PipelineRunSourceTypeOutput struct{ *pulumi.OutputState }

func (PipelineRunSourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunSourceType)(nil)).Elem()
}

func (o PipelineRunSourceTypeOutput) ToPipelineRunSourceTypeOutput() PipelineRunSourceTypeOutput {
	return o
}

func (o PipelineRunSourceTypeOutput) ToPipelineRunSourceTypeOutputWithContext(ctx context.Context) PipelineRunSourceTypeOutput {
	return o
}

func (o PipelineRunSourceTypeOutput) ToPipelineRunSourceTypePtrOutput() PipelineRunSourceTypePtrOutput {
	return o.ToPipelineRunSourceTypePtrOutputWithContext(context.Background())
}

func (o PipelineRunSourceTypeOutput) ToPipelineRunSourceTypePtrOutputWithContext(ctx context.Context) PipelineRunSourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunSourceType) *PipelineRunSourceType {
		return &v
	}).(PipelineRunSourceTypePtrOutput)
}

func (o PipelineRunSourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PipelineRunSourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PipelineRunSourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PipelineRunSourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PipelineRunSourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PipelineRunSourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PipelineRunSourceTypePtrOutput struct{ *pulumi.OutputState }

func (PipelineRunSourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunSourceType)(nil)).Elem()
}

func (o PipelineRunSourceTypePtrOutput) ToPipelineRunSourceTypePtrOutput() PipelineRunSourceTypePtrOutput {
	return o
}

func (o PipelineRunSourceTypePtrOutput) ToPipelineRunSourceTypePtrOutputWithContext(ctx context.Context) PipelineRunSourceTypePtrOutput {
	return o
}

func (o PipelineRunSourceTypePtrOutput) Elem() PipelineRunSourceTypeOutput {
	return o.ApplyT(func(v *PipelineRunSourceType) PipelineRunSourceType {
		if v != nil {
			return *v
		}
		var ret PipelineRunSourceType
		return ret
	}).(PipelineRunSourceTypeOutput)
}

func (o PipelineRunSourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PipelineRunSourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PipelineRunSourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PipelineRunSourceTypeInput interface {
	pulumi.Input

	ToPipelineRunSourceTypeOutput() PipelineRunSourceTypeOutput
	ToPipelineRunSourceTypeOutputWithContext(context.Context) PipelineRunSourceTypeOutput
}

var pipelineRunSourceTypePtrType = reflect.TypeOf((**PipelineRunSourceType)(nil)).Elem()

type PipelineRunSourceTypePtrInput interface {
	pulumi.Input

	ToPipelineRunSourceTypePtrOutput() PipelineRunSourceTypePtrOutput
	ToPipelineRunSourceTypePtrOutputWithContext(context.Context) PipelineRunSourceTypePtrOutput
}

type pipelineRunSourceTypePtr string

func PipelineRunSourceTypePtr(v string) PipelineRunSourceTypePtrInput {
	return (*pipelineRunSourceTypePtr)(&v)
}

func (*pipelineRunSourceTypePtr) ElementType() reflect.Type {
	return pipelineRunSourceTypePtrType
}

func (in *pipelineRunSourceTypePtr) ToPipelineRunSourceTypePtrOutput() PipelineRunSourceTypePtrOutput {
	return pulumi.ToOutput(in).(PipelineRunSourceTypePtrOutput)
}

func (in *pipelineRunSourceTypePtr) ToPipelineRunSourceTypePtrOutputWithContext(ctx context.Context) PipelineRunSourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PipelineRunSourceTypePtrOutput)
}

type PipelineRunTargetType string

const (
	PipelineRunTargetTypeAzureStorageBlob = PipelineRunTargetType("AzureStorageBlob")
)

func (PipelineRunTargetType) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunTargetType)(nil)).Elem()
}

func (e PipelineRunTargetType) ToPipelineRunTargetTypeOutput() PipelineRunTargetTypeOutput {
	return pulumi.ToOutput(e).(PipelineRunTargetTypeOutput)
}

func (e PipelineRunTargetType) ToPipelineRunTargetTypeOutputWithContext(ctx context.Context) PipelineRunTargetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PipelineRunTargetTypeOutput)
}

func (e PipelineRunTargetType) ToPipelineRunTargetTypePtrOutput() PipelineRunTargetTypePtrOutput {
	return e.ToPipelineRunTargetTypePtrOutputWithContext(context.Background())
}

func (e PipelineRunTargetType) ToPipelineRunTargetTypePtrOutputWithContext(ctx context.Context) PipelineRunTargetTypePtrOutput {
	return PipelineRunTargetType(e).ToPipelineRunTargetTypeOutputWithContext(ctx).ToPipelineRunTargetTypePtrOutputWithContext(ctx)
}

func (e PipelineRunTargetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PipelineRunTargetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PipelineRunTargetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PipelineRunTargetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PipelineRunTargetTypeOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunTargetType)(nil)).Elem()
}

func (o PipelineRunTargetTypeOutput) ToPipelineRunTargetTypeOutput() PipelineRunTargetTypeOutput {
	return o
}

func (o PipelineRunTargetTypeOutput) ToPipelineRunTargetTypeOutputWithContext(ctx context.Context) PipelineRunTargetTypeOutput {
	return o
}

func (o PipelineRunTargetTypeOutput) ToPipelineRunTargetTypePtrOutput() PipelineRunTargetTypePtrOutput {
	return o.ToPipelineRunTargetTypePtrOutputWithContext(context.Background())
}

func (o PipelineRunTargetTypeOutput) ToPipelineRunTargetTypePtrOutputWithContext(ctx context.Context) PipelineRunTargetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunTargetType) *PipelineRunTargetType {
		return &v
	}).(PipelineRunTargetTypePtrOutput)
}

func (o PipelineRunTargetTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PipelineRunTargetTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PipelineRunTargetType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PipelineRunTargetTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PipelineRunTargetTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PipelineRunTargetType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PipelineRunTargetTypePtrOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunTargetType)(nil)).Elem()
}

func (o PipelineRunTargetTypePtrOutput) ToPipelineRunTargetTypePtrOutput() PipelineRunTargetTypePtrOutput {
	return o
}

func (o PipelineRunTargetTypePtrOutput) ToPipelineRunTargetTypePtrOutputWithContext(ctx context.Context) PipelineRunTargetTypePtrOutput {
	return o
}

func (o PipelineRunTargetTypePtrOutput) Elem() PipelineRunTargetTypeOutput {
	return o.ApplyT(func(v *PipelineRunTargetType) PipelineRunTargetType {
		if v != nil {
			return *v
		}
		var ret PipelineRunTargetType
		return ret
	}).(PipelineRunTargetTypeOutput)
}

func (o PipelineRunTargetTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PipelineRunTargetTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PipelineRunTargetType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PipelineRunTargetTypeInput interface {
	pulumi.Input

	ToPipelineRunTargetTypeOutput() PipelineRunTargetTypeOutput
	ToPipelineRunTargetTypeOutputWithContext(context.Context) PipelineRunTargetTypeOutput
}

var pipelineRunTargetTypePtrType = reflect.TypeOf((**PipelineRunTargetType)(nil)).Elem()

type PipelineRunTargetTypePtrInput interface {
	pulumi.Input

	ToPipelineRunTargetTypePtrOutput() PipelineRunTargetTypePtrOutput
	ToPipelineRunTargetTypePtrOutputWithContext(context.Context) PipelineRunTargetTypePtrOutput
}

type pipelineRunTargetTypePtr string

func PipelineRunTargetTypePtr(v string) PipelineRunTargetTypePtrInput {
	return (*pipelineRunTargetTypePtr)(&v)
}

func (*pipelineRunTargetTypePtr) ElementType() reflect.Type {
	return pipelineRunTargetTypePtrType
}

func (in *pipelineRunTargetTypePtr) ToPipelineRunTargetTypePtrOutput() PipelineRunTargetTypePtrOutput {
	return pulumi.ToOutput(in).(PipelineRunTargetTypePtrOutput)
}

func (in *pipelineRunTargetTypePtr) ToPipelineRunTargetTypePtrOutputWithContext(ctx context.Context) PipelineRunTargetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PipelineRunTargetTypePtrOutput)
}

type PipelineSourceType string

const (
	PipelineSourceTypeAzureStorageBlobContainer = PipelineSourceType("AzureStorageBlobContainer")
)

func (PipelineSourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceType)(nil)).Elem()
}

func (e PipelineSourceType) ToPipelineSourceTypeOutput() PipelineSourceTypeOutput {
	return pulumi.ToOutput(e).(PipelineSourceTypeOutput)
}

func (e PipelineSourceType) ToPipelineSourceTypeOutputWithContext(ctx context.Context) PipelineSourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PipelineSourceTypeOutput)
}

func (e PipelineSourceType) ToPipelineSourceTypePtrOutput() PipelineSourceTypePtrOutput {
	return e.ToPipelineSourceTypePtrOutputWithContext(context.Background())
}

func (e PipelineSourceType) ToPipelineSourceTypePtrOutputWithContext(ctx context.Context) PipelineSourceTypePtrOutput {
	return PipelineSourceType(e).ToPipelineSourceTypeOutputWithContext(ctx).ToPipelineSourceTypePtrOutputWithContext(ctx)
}

func (e PipelineSourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PipelineSourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PipelineSourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PipelineSourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PipelineSourceTypeOutput struct{ *pulumi.OutputState }

func (PipelineSourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceType)(nil)).Elem()
}

func (o PipelineSourceTypeOutput) ToPipelineSourceTypeOutput() PipelineSourceTypeOutput {
	return o
}

func (o PipelineSourceTypeOutput) ToPipelineSourceTypeOutputWithContext(ctx context.Context) PipelineSourceTypeOutput {
	return o
}

func (o PipelineSourceTypeOutput) ToPipelineSourceTypePtrOutput() PipelineSourceTypePtrOutput {
	return o.ToPipelineSourceTypePtrOutputWithContext(context.Background())
}

func (o PipelineSourceTypeOutput) ToPipelineSourceTypePtrOutputWithContext(ctx context.Context) PipelineSourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineSourceType) *PipelineSourceType {
		return &v
	}).(PipelineSourceTypePtrOutput)
}

func (o PipelineSourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PipelineSourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PipelineSourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PipelineSourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PipelineSourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PipelineSourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PipelineSourceTypePtrOutput struct{ *pulumi.OutputState }

func (PipelineSourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceType)(nil)).Elem()
}

func (o PipelineSourceTypePtrOutput) ToPipelineSourceTypePtrOutput() PipelineSourceTypePtrOutput {
	return o
}

func (o PipelineSourceTypePtrOutput) ToPipelineSourceTypePtrOutputWithContext(ctx context.Context) PipelineSourceTypePtrOutput {
	return o
}

func (o PipelineSourceTypePtrOutput) Elem() PipelineSourceTypeOutput {
	return o.ApplyT(func(v *PipelineSourceType) PipelineSourceType {
		if v != nil {
			return *v
		}
		var ret PipelineSourceType
		return ret
	}).(PipelineSourceTypeOutput)
}

func (o PipelineSourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PipelineSourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PipelineSourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PipelineSourceTypeInput interface {
	pulumi.Input

	ToPipelineSourceTypeOutput() PipelineSourceTypeOutput
	ToPipelineSourceTypeOutputWithContext(context.Context) PipelineSourceTypeOutput
}

var pipelineSourceTypePtrType = reflect.TypeOf((**PipelineSourceType)(nil)).Elem()

type PipelineSourceTypePtrInput interface {
	pulumi.Input

	ToPipelineSourceTypePtrOutput() PipelineSourceTypePtrOutput
	ToPipelineSourceTypePtrOutputWithContext(context.Context) PipelineSourceTypePtrOutput
}

type pipelineSourceTypePtr string

func PipelineSourceTypePtr(v string) PipelineSourceTypePtrInput {
	return (*pipelineSourceTypePtr)(&v)
}

func (*pipelineSourceTypePtr) ElementType() reflect.Type {
	return pipelineSourceTypePtrType
}

func (in *pipelineSourceTypePtr) ToPipelineSourceTypePtrOutput() PipelineSourceTypePtrOutput {
	return pulumi.ToOutput(in).(PipelineSourceTypePtrOutput)
}

func (in *pipelineSourceTypePtr) ToPipelineSourceTypePtrOutputWithContext(ctx context.Context) PipelineSourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PipelineSourceTypePtrOutput)
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

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type SecretObjectType string

const (
	SecretObjectTypeOpaque      = SecretObjectType("Opaque")
	SecretObjectTypeVaultsecret = SecretObjectType("Vaultsecret")
)

func (SecretObjectType) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretObjectType)(nil)).Elem()
}

func (e SecretObjectType) ToSecretObjectTypeOutput() SecretObjectTypeOutput {
	return pulumi.ToOutput(e).(SecretObjectTypeOutput)
}

func (e SecretObjectType) ToSecretObjectTypeOutputWithContext(ctx context.Context) SecretObjectTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecretObjectTypeOutput)
}

func (e SecretObjectType) ToSecretObjectTypePtrOutput() SecretObjectTypePtrOutput {
	return e.ToSecretObjectTypePtrOutputWithContext(context.Background())
}

func (e SecretObjectType) ToSecretObjectTypePtrOutputWithContext(ctx context.Context) SecretObjectTypePtrOutput {
	return SecretObjectType(e).ToSecretObjectTypeOutputWithContext(ctx).ToSecretObjectTypePtrOutputWithContext(ctx)
}

func (e SecretObjectType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecretObjectType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecretObjectType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecretObjectType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecretObjectTypeOutput struct{ *pulumi.OutputState }

func (SecretObjectTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretObjectType)(nil)).Elem()
}

func (o SecretObjectTypeOutput) ToSecretObjectTypeOutput() SecretObjectTypeOutput {
	return o
}

func (o SecretObjectTypeOutput) ToSecretObjectTypeOutputWithContext(ctx context.Context) SecretObjectTypeOutput {
	return o
}

func (o SecretObjectTypeOutput) ToSecretObjectTypePtrOutput() SecretObjectTypePtrOutput {
	return o.ToSecretObjectTypePtrOutputWithContext(context.Background())
}

func (o SecretObjectTypeOutput) ToSecretObjectTypePtrOutputWithContext(ctx context.Context) SecretObjectTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretObjectType) *SecretObjectType {
		return &v
	}).(SecretObjectTypePtrOutput)
}

func (o SecretObjectTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecretObjectTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecretObjectType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecretObjectTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecretObjectTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecretObjectType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecretObjectTypePtrOutput struct{ *pulumi.OutputState }

func (SecretObjectTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretObjectType)(nil)).Elem()
}

func (o SecretObjectTypePtrOutput) ToSecretObjectTypePtrOutput() SecretObjectTypePtrOutput {
	return o
}

func (o SecretObjectTypePtrOutput) ToSecretObjectTypePtrOutputWithContext(ctx context.Context) SecretObjectTypePtrOutput {
	return o
}

func (o SecretObjectTypePtrOutput) Elem() SecretObjectTypeOutput {
	return o.ApplyT(func(v *SecretObjectType) SecretObjectType {
		if v != nil {
			return *v
		}
		var ret SecretObjectType
		return ret
	}).(SecretObjectTypeOutput)
}

func (o SecretObjectTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecretObjectTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecretObjectType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecretObjectTypeInput interface {
	pulumi.Input

	ToSecretObjectTypeOutput() SecretObjectTypeOutput
	ToSecretObjectTypeOutputWithContext(context.Context) SecretObjectTypeOutput
}

var secretObjectTypePtrType = reflect.TypeOf((**SecretObjectType)(nil)).Elem()

type SecretObjectTypePtrInput interface {
	pulumi.Input

	ToSecretObjectTypePtrOutput() SecretObjectTypePtrOutput
	ToSecretObjectTypePtrOutputWithContext(context.Context) SecretObjectTypePtrOutput
}

type secretObjectTypePtr string

func SecretObjectTypePtr(v string) SecretObjectTypePtrInput {
	return (*secretObjectTypePtr)(&v)
}

func (*secretObjectTypePtr) ElementType() reflect.Type {
	return secretObjectTypePtrType
}

func (in *secretObjectTypePtr) ToSecretObjectTypePtrOutput() SecretObjectTypePtrOutput {
	return pulumi.ToOutput(in).(SecretObjectTypePtrOutput)
}

func (in *secretObjectTypePtr) ToSecretObjectTypePtrOutputWithContext(ctx context.Context) SecretObjectTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecretObjectTypePtrOutput)
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

type SourceControlType string

const (
	SourceControlTypeGithub                  = SourceControlType("Github")
	SourceControlTypeVisualStudioTeamService = SourceControlType("VisualStudioTeamService")
)

func (SourceControlType) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlType)(nil)).Elem()
}

func (e SourceControlType) ToSourceControlTypeOutput() SourceControlTypeOutput {
	return pulumi.ToOutput(e).(SourceControlTypeOutput)
}

func (e SourceControlType) ToSourceControlTypeOutputWithContext(ctx context.Context) SourceControlTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceControlTypeOutput)
}

func (e SourceControlType) ToSourceControlTypePtrOutput() SourceControlTypePtrOutput {
	return e.ToSourceControlTypePtrOutputWithContext(context.Background())
}

func (e SourceControlType) ToSourceControlTypePtrOutputWithContext(ctx context.Context) SourceControlTypePtrOutput {
	return SourceControlType(e).ToSourceControlTypeOutputWithContext(ctx).ToSourceControlTypePtrOutputWithContext(ctx)
}

func (e SourceControlType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceControlType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceControlType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceControlType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceControlTypeOutput struct{ *pulumi.OutputState }

func (SourceControlTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlType)(nil)).Elem()
}

func (o SourceControlTypeOutput) ToSourceControlTypeOutput() SourceControlTypeOutput {
	return o
}

func (o SourceControlTypeOutput) ToSourceControlTypeOutputWithContext(ctx context.Context) SourceControlTypeOutput {
	return o
}

func (o SourceControlTypeOutput) ToSourceControlTypePtrOutput() SourceControlTypePtrOutput {
	return o.ToSourceControlTypePtrOutputWithContext(context.Background())
}

func (o SourceControlTypeOutput) ToSourceControlTypePtrOutputWithContext(ctx context.Context) SourceControlTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceControlType) *SourceControlType {
		return &v
	}).(SourceControlTypePtrOutput)
}

func (o SourceControlTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceControlTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceControlType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceControlTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceControlTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceControlType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceControlTypePtrOutput struct{ *pulumi.OutputState }

func (SourceControlTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlType)(nil)).Elem()
}

func (o SourceControlTypePtrOutput) ToSourceControlTypePtrOutput() SourceControlTypePtrOutput {
	return o
}

func (o SourceControlTypePtrOutput) ToSourceControlTypePtrOutputWithContext(ctx context.Context) SourceControlTypePtrOutput {
	return o
}

func (o SourceControlTypePtrOutput) Elem() SourceControlTypeOutput {
	return o.ApplyT(func(v *SourceControlType) SourceControlType {
		if v != nil {
			return *v
		}
		var ret SourceControlType
		return ret
	}).(SourceControlTypeOutput)
}

func (o SourceControlTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceControlTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceControlType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceControlTypeInput interface {
	pulumi.Input

	ToSourceControlTypeOutput() SourceControlTypeOutput
	ToSourceControlTypeOutputWithContext(context.Context) SourceControlTypeOutput
}

var sourceControlTypePtrType = reflect.TypeOf((**SourceControlType)(nil)).Elem()

type SourceControlTypePtrInput interface {
	pulumi.Input

	ToSourceControlTypePtrOutput() SourceControlTypePtrOutput
	ToSourceControlTypePtrOutputWithContext(context.Context) SourceControlTypePtrOutput
}

type sourceControlTypePtr string

func SourceControlTypePtr(v string) SourceControlTypePtrInput {
	return (*sourceControlTypePtr)(&v)
}

func (*sourceControlTypePtr) ElementType() reflect.Type {
	return sourceControlTypePtrType
}

func (in *sourceControlTypePtr) ToSourceControlTypePtrOutput() SourceControlTypePtrOutput {
	return pulumi.ToOutput(in).(SourceControlTypePtrOutput)
}

func (in *sourceControlTypePtr) ToSourceControlTypePtrOutputWithContext(ctx context.Context) SourceControlTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceControlTypePtrOutput)
}

type SourceRegistryLoginMode string

const (
	SourceRegistryLoginModeNone    = SourceRegistryLoginMode("None")
	SourceRegistryLoginModeDefault = SourceRegistryLoginMode("Default")
)

func (SourceRegistryLoginMode) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRegistryLoginMode)(nil)).Elem()
}

func (e SourceRegistryLoginMode) ToSourceRegistryLoginModeOutput() SourceRegistryLoginModeOutput {
	return pulumi.ToOutput(e).(SourceRegistryLoginModeOutput)
}

func (e SourceRegistryLoginMode) ToSourceRegistryLoginModeOutputWithContext(ctx context.Context) SourceRegistryLoginModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceRegistryLoginModeOutput)
}

func (e SourceRegistryLoginMode) ToSourceRegistryLoginModePtrOutput() SourceRegistryLoginModePtrOutput {
	return e.ToSourceRegistryLoginModePtrOutputWithContext(context.Background())
}

func (e SourceRegistryLoginMode) ToSourceRegistryLoginModePtrOutputWithContext(ctx context.Context) SourceRegistryLoginModePtrOutput {
	return SourceRegistryLoginMode(e).ToSourceRegistryLoginModeOutputWithContext(ctx).ToSourceRegistryLoginModePtrOutputWithContext(ctx)
}

func (e SourceRegistryLoginMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceRegistryLoginMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceRegistryLoginMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceRegistryLoginMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceRegistryLoginModeOutput struct{ *pulumi.OutputState }

func (SourceRegistryLoginModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceRegistryLoginMode)(nil)).Elem()
}

func (o SourceRegistryLoginModeOutput) ToSourceRegistryLoginModeOutput() SourceRegistryLoginModeOutput {
	return o
}

func (o SourceRegistryLoginModeOutput) ToSourceRegistryLoginModeOutputWithContext(ctx context.Context) SourceRegistryLoginModeOutput {
	return o
}

func (o SourceRegistryLoginModeOutput) ToSourceRegistryLoginModePtrOutput() SourceRegistryLoginModePtrOutput {
	return o.ToSourceRegistryLoginModePtrOutputWithContext(context.Background())
}

func (o SourceRegistryLoginModeOutput) ToSourceRegistryLoginModePtrOutputWithContext(ctx context.Context) SourceRegistryLoginModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceRegistryLoginMode) *SourceRegistryLoginMode {
		return &v
	}).(SourceRegistryLoginModePtrOutput)
}

func (o SourceRegistryLoginModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceRegistryLoginModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceRegistryLoginMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceRegistryLoginModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceRegistryLoginModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceRegistryLoginMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceRegistryLoginModePtrOutput struct{ *pulumi.OutputState }

func (SourceRegistryLoginModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceRegistryLoginMode)(nil)).Elem()
}

func (o SourceRegistryLoginModePtrOutput) ToSourceRegistryLoginModePtrOutput() SourceRegistryLoginModePtrOutput {
	return o
}

func (o SourceRegistryLoginModePtrOutput) ToSourceRegistryLoginModePtrOutputWithContext(ctx context.Context) SourceRegistryLoginModePtrOutput {
	return o
}

func (o SourceRegistryLoginModePtrOutput) Elem() SourceRegistryLoginModeOutput {
	return o.ApplyT(func(v *SourceRegistryLoginMode) SourceRegistryLoginMode {
		if v != nil {
			return *v
		}
		var ret SourceRegistryLoginMode
		return ret
	}).(SourceRegistryLoginModeOutput)
}

func (o SourceRegistryLoginModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceRegistryLoginModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceRegistryLoginMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceRegistryLoginModeInput interface {
	pulumi.Input

	ToSourceRegistryLoginModeOutput() SourceRegistryLoginModeOutput
	ToSourceRegistryLoginModeOutputWithContext(context.Context) SourceRegistryLoginModeOutput
}

var sourceRegistryLoginModePtrType = reflect.TypeOf((**SourceRegistryLoginMode)(nil)).Elem()

type SourceRegistryLoginModePtrInput interface {
	pulumi.Input

	ToSourceRegistryLoginModePtrOutput() SourceRegistryLoginModePtrOutput
	ToSourceRegistryLoginModePtrOutputWithContext(context.Context) SourceRegistryLoginModePtrOutput
}

type sourceRegistryLoginModePtr string

func SourceRegistryLoginModePtr(v string) SourceRegistryLoginModePtrInput {
	return (*sourceRegistryLoginModePtr)(&v)
}

func (*sourceRegistryLoginModePtr) ElementType() reflect.Type {
	return sourceRegistryLoginModePtrType
}

func (in *sourceRegistryLoginModePtr) ToSourceRegistryLoginModePtrOutput() SourceRegistryLoginModePtrOutput {
	return pulumi.ToOutput(in).(SourceRegistryLoginModePtrOutput)
}

func (in *sourceRegistryLoginModePtr) ToSourceRegistryLoginModePtrOutputWithContext(ctx context.Context) SourceRegistryLoginModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceRegistryLoginModePtrOutput)
}

type SourceTriggerEvent string

const (
	SourceTriggerEventCommit      = SourceTriggerEvent("commit")
	SourceTriggerEventPullrequest = SourceTriggerEvent("pullrequest")
)

func (SourceTriggerEvent) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceTriggerEvent)(nil)).Elem()
}

func (e SourceTriggerEvent) ToSourceTriggerEventOutput() SourceTriggerEventOutput {
	return pulumi.ToOutput(e).(SourceTriggerEventOutput)
}

func (e SourceTriggerEvent) ToSourceTriggerEventOutputWithContext(ctx context.Context) SourceTriggerEventOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceTriggerEventOutput)
}

func (e SourceTriggerEvent) ToSourceTriggerEventPtrOutput() SourceTriggerEventPtrOutput {
	return e.ToSourceTriggerEventPtrOutputWithContext(context.Background())
}

func (e SourceTriggerEvent) ToSourceTriggerEventPtrOutputWithContext(ctx context.Context) SourceTriggerEventPtrOutput {
	return SourceTriggerEvent(e).ToSourceTriggerEventOutputWithContext(ctx).ToSourceTriggerEventPtrOutputWithContext(ctx)
}

func (e SourceTriggerEvent) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceTriggerEvent) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceTriggerEvent) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceTriggerEvent) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceTriggerEventOutput struct{ *pulumi.OutputState }

func (SourceTriggerEventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceTriggerEvent)(nil)).Elem()
}

func (o SourceTriggerEventOutput) ToSourceTriggerEventOutput() SourceTriggerEventOutput {
	return o
}

func (o SourceTriggerEventOutput) ToSourceTriggerEventOutputWithContext(ctx context.Context) SourceTriggerEventOutput {
	return o
}

func (o SourceTriggerEventOutput) ToSourceTriggerEventPtrOutput() SourceTriggerEventPtrOutput {
	return o.ToSourceTriggerEventPtrOutputWithContext(context.Background())
}

func (o SourceTriggerEventOutput) ToSourceTriggerEventPtrOutputWithContext(ctx context.Context) SourceTriggerEventPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceTriggerEvent) *SourceTriggerEvent {
		return &v
	}).(SourceTriggerEventPtrOutput)
}

func (o SourceTriggerEventOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceTriggerEventOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceTriggerEvent) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceTriggerEventOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTriggerEventOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceTriggerEvent) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceTriggerEventPtrOutput struct{ *pulumi.OutputState }

func (SourceTriggerEventPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceTriggerEvent)(nil)).Elem()
}

func (o SourceTriggerEventPtrOutput) ToSourceTriggerEventPtrOutput() SourceTriggerEventPtrOutput {
	return o
}

func (o SourceTriggerEventPtrOutput) ToSourceTriggerEventPtrOutputWithContext(ctx context.Context) SourceTriggerEventPtrOutput {
	return o
}

func (o SourceTriggerEventPtrOutput) Elem() SourceTriggerEventOutput {
	return o.ApplyT(func(v *SourceTriggerEvent) SourceTriggerEvent {
		if v != nil {
			return *v
		}
		var ret SourceTriggerEvent
		return ret
	}).(SourceTriggerEventOutput)
}

func (o SourceTriggerEventPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTriggerEventPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceTriggerEvent) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceTriggerEventInput interface {
	pulumi.Input

	ToSourceTriggerEventOutput() SourceTriggerEventOutput
	ToSourceTriggerEventOutputWithContext(context.Context) SourceTriggerEventOutput
}

var sourceTriggerEventPtrType = reflect.TypeOf((**SourceTriggerEvent)(nil)).Elem()

type SourceTriggerEventPtrInput interface {
	pulumi.Input

	ToSourceTriggerEventPtrOutput() SourceTriggerEventPtrOutput
	ToSourceTriggerEventPtrOutputWithContext(context.Context) SourceTriggerEventPtrOutput
}

type sourceTriggerEventPtr string

func SourceTriggerEventPtr(v string) SourceTriggerEventPtrInput {
	return (*sourceTriggerEventPtr)(&v)
}

func (*sourceTriggerEventPtr) ElementType() reflect.Type {
	return sourceTriggerEventPtrType
}

func (in *sourceTriggerEventPtr) ToSourceTriggerEventPtrOutput() SourceTriggerEventPtrOutput {
	return pulumi.ToOutput(in).(SourceTriggerEventPtrOutput)
}

func (in *sourceTriggerEventPtr) ToSourceTriggerEventPtrOutputWithContext(ctx context.Context) SourceTriggerEventPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceTriggerEventPtrOutput)
}

type StepType string

const (
	StepTypeDocker      = StepType("Docker")
	StepTypeFileTask    = StepType("FileTask")
	StepTypeEncodedTask = StepType("EncodedTask")
)

func (StepType) ElementType() reflect.Type {
	return reflect.TypeOf((*StepType)(nil)).Elem()
}

func (e StepType) ToStepTypeOutput() StepTypeOutput {
	return pulumi.ToOutput(e).(StepTypeOutput)
}

func (e StepType) ToStepTypeOutputWithContext(ctx context.Context) StepTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StepTypeOutput)
}

func (e StepType) ToStepTypePtrOutput() StepTypePtrOutput {
	return e.ToStepTypePtrOutputWithContext(context.Background())
}

func (e StepType) ToStepTypePtrOutputWithContext(ctx context.Context) StepTypePtrOutput {
	return StepType(e).ToStepTypeOutputWithContext(ctx).ToStepTypePtrOutputWithContext(ctx)
}

func (e StepType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StepType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StepType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StepType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StepTypeOutput struct{ *pulumi.OutputState }

func (StepTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepType)(nil)).Elem()
}

func (o StepTypeOutput) ToStepTypeOutput() StepTypeOutput {
	return o
}

func (o StepTypeOutput) ToStepTypeOutputWithContext(ctx context.Context) StepTypeOutput {
	return o
}

func (o StepTypeOutput) ToStepTypePtrOutput() StepTypePtrOutput {
	return o.ToStepTypePtrOutputWithContext(context.Background())
}

func (o StepTypeOutput) ToStepTypePtrOutputWithContext(ctx context.Context) StepTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StepType) *StepType {
		return &v
	}).(StepTypePtrOutput)
}

func (o StepTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StepTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StepType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StepTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StepTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StepType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StepTypePtrOutput struct{ *pulumi.OutputState }

func (StepTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StepType)(nil)).Elem()
}

func (o StepTypePtrOutput) ToStepTypePtrOutput() StepTypePtrOutput {
	return o
}

func (o StepTypePtrOutput) ToStepTypePtrOutputWithContext(ctx context.Context) StepTypePtrOutput {
	return o
}

func (o StepTypePtrOutput) Elem() StepTypeOutput {
	return o.ApplyT(func(v *StepType) StepType {
		if v != nil {
			return *v
		}
		var ret StepType
		return ret
	}).(StepTypeOutput)
}

func (o StepTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StepTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StepType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StepTypeInput interface {
	pulumi.Input

	ToStepTypeOutput() StepTypeOutput
	ToStepTypeOutputWithContext(context.Context) StepTypeOutput
}

var stepTypePtrType = reflect.TypeOf((**StepType)(nil)).Elem()

type StepTypePtrInput interface {
	pulumi.Input

	ToStepTypePtrOutput() StepTypePtrOutput
	ToStepTypePtrOutputWithContext(context.Context) StepTypePtrOutput
}

type stepTypePtr string

func StepTypePtr(v string) StepTypePtrInput {
	return (*stepTypePtr)(&v)
}

func (*stepTypePtr) ElementType() reflect.Type {
	return stepTypePtrType
}

func (in *stepTypePtr) ToStepTypePtrOutput() StepTypePtrOutput {
	return pulumi.ToOutput(in).(StepTypePtrOutput)
}

func (in *stepTypePtr) ToStepTypePtrOutputWithContext(ctx context.Context) StepTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StepTypePtrOutput)
}

type TaskStatus string

const (
	TaskStatusDisabled = TaskStatus("Disabled")
	TaskStatusEnabled  = TaskStatus("Enabled")
)

func (TaskStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskStatus)(nil)).Elem()
}

func (e TaskStatus) ToTaskStatusOutput() TaskStatusOutput {
	return pulumi.ToOutput(e).(TaskStatusOutput)
}

func (e TaskStatus) ToTaskStatusOutputWithContext(ctx context.Context) TaskStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TaskStatusOutput)
}

func (e TaskStatus) ToTaskStatusPtrOutput() TaskStatusPtrOutput {
	return e.ToTaskStatusPtrOutputWithContext(context.Background())
}

func (e TaskStatus) ToTaskStatusPtrOutputWithContext(ctx context.Context) TaskStatusPtrOutput {
	return TaskStatus(e).ToTaskStatusOutputWithContext(ctx).ToTaskStatusPtrOutputWithContext(ctx)
}

func (e TaskStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TaskStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TaskStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TaskStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TaskStatusOutput struct{ *pulumi.OutputState }

func (TaskStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskStatus)(nil)).Elem()
}

func (o TaskStatusOutput) ToTaskStatusOutput() TaskStatusOutput {
	return o
}

func (o TaskStatusOutput) ToTaskStatusOutputWithContext(ctx context.Context) TaskStatusOutput {
	return o
}

func (o TaskStatusOutput) ToTaskStatusPtrOutput() TaskStatusPtrOutput {
	return o.ToTaskStatusPtrOutputWithContext(context.Background())
}

func (o TaskStatusOutput) ToTaskStatusPtrOutputWithContext(ctx context.Context) TaskStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TaskStatus) *TaskStatus {
		return &v
	}).(TaskStatusPtrOutput)
}

func (o TaskStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TaskStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TaskStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TaskStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TaskStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TaskStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TaskStatusPtrOutput struct{ *pulumi.OutputState }

func (TaskStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskStatus)(nil)).Elem()
}

func (o TaskStatusPtrOutput) ToTaskStatusPtrOutput() TaskStatusPtrOutput {
	return o
}

func (o TaskStatusPtrOutput) ToTaskStatusPtrOutputWithContext(ctx context.Context) TaskStatusPtrOutput {
	return o
}

func (o TaskStatusPtrOutput) Elem() TaskStatusOutput {
	return o.ApplyT(func(v *TaskStatus) TaskStatus {
		if v != nil {
			return *v
		}
		var ret TaskStatus
		return ret
	}).(TaskStatusOutput)
}

func (o TaskStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TaskStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TaskStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TaskStatusInput interface {
	pulumi.Input

	ToTaskStatusOutput() TaskStatusOutput
	ToTaskStatusOutputWithContext(context.Context) TaskStatusOutput
}

var taskStatusPtrType = reflect.TypeOf((**TaskStatus)(nil)).Elem()

type TaskStatusPtrInput interface {
	pulumi.Input

	ToTaskStatusPtrOutput() TaskStatusPtrOutput
	ToTaskStatusPtrOutputWithContext(context.Context) TaskStatusPtrOutput
}

type taskStatusPtr string

func TaskStatusPtr(v string) TaskStatusPtrInput {
	return (*taskStatusPtr)(&v)
}

func (*taskStatusPtr) ElementType() reflect.Type {
	return taskStatusPtrType
}

func (in *taskStatusPtr) ToTaskStatusPtrOutput() TaskStatusPtrOutput {
	return pulumi.ToOutput(in).(TaskStatusPtrOutput)
}

func (in *taskStatusPtr) ToTaskStatusPtrOutputWithContext(ctx context.Context) TaskStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TaskStatusPtrOutput)
}

type TokenCertificateName string

const (
	TokenCertificateNameCertificate1 = TokenCertificateName("certificate1")
	TokenCertificateNameCertificate2 = TokenCertificateName("certificate2")
)

func (TokenCertificateName) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificateName)(nil)).Elem()
}

func (e TokenCertificateName) ToTokenCertificateNameOutput() TokenCertificateNameOutput {
	return pulumi.ToOutput(e).(TokenCertificateNameOutput)
}

func (e TokenCertificateName) ToTokenCertificateNameOutputWithContext(ctx context.Context) TokenCertificateNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TokenCertificateNameOutput)
}

func (e TokenCertificateName) ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput {
	return e.ToTokenCertificateNamePtrOutputWithContext(context.Background())
}

func (e TokenCertificateName) ToTokenCertificateNamePtrOutputWithContext(ctx context.Context) TokenCertificateNamePtrOutput {
	return TokenCertificateName(e).ToTokenCertificateNameOutputWithContext(ctx).ToTokenCertificateNamePtrOutputWithContext(ctx)
}

func (e TokenCertificateName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenCertificateName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenCertificateName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TokenCertificateName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TokenCertificateNameOutput struct{ *pulumi.OutputState }

func (TokenCertificateNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificateName)(nil)).Elem()
}

func (o TokenCertificateNameOutput) ToTokenCertificateNameOutput() TokenCertificateNameOutput {
	return o
}

func (o TokenCertificateNameOutput) ToTokenCertificateNameOutputWithContext(ctx context.Context) TokenCertificateNameOutput {
	return o
}

func (o TokenCertificateNameOutput) ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput {
	return o.ToTokenCertificateNamePtrOutputWithContext(context.Background())
}

func (o TokenCertificateNameOutput) ToTokenCertificateNamePtrOutputWithContext(ctx context.Context) TokenCertificateNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenCertificateName) *TokenCertificateName {
		return &v
	}).(TokenCertificateNamePtrOutput)
}

func (o TokenCertificateNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TokenCertificateNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenCertificateName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TokenCertificateNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenCertificateNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenCertificateName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TokenCertificateNamePtrOutput struct{ *pulumi.OutputState }

func (TokenCertificateNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCertificateName)(nil)).Elem()
}

func (o TokenCertificateNamePtrOutput) ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput {
	return o
}

func (o TokenCertificateNamePtrOutput) ToTokenCertificateNamePtrOutputWithContext(ctx context.Context) TokenCertificateNamePtrOutput {
	return o
}

func (o TokenCertificateNamePtrOutput) Elem() TokenCertificateNameOutput {
	return o.ApplyT(func(v *TokenCertificateName) TokenCertificateName {
		if v != nil {
			return *v
		}
		var ret TokenCertificateName
		return ret
	}).(TokenCertificateNameOutput)
}

func (o TokenCertificateNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenCertificateNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TokenCertificateName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TokenCertificateNameInput interface {
	pulumi.Input

	ToTokenCertificateNameOutput() TokenCertificateNameOutput
	ToTokenCertificateNameOutputWithContext(context.Context) TokenCertificateNameOutput
}

var tokenCertificateNamePtrType = reflect.TypeOf((**TokenCertificateName)(nil)).Elem()

type TokenCertificateNamePtrInput interface {
	pulumi.Input

	ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput
	ToTokenCertificateNamePtrOutputWithContext(context.Context) TokenCertificateNamePtrOutput
}

type tokenCertificateNamePtr string

func TokenCertificateNamePtr(v string) TokenCertificateNamePtrInput {
	return (*tokenCertificateNamePtr)(&v)
}

func (*tokenCertificateNamePtr) ElementType() reflect.Type {
	return tokenCertificateNamePtrType
}

func (in *tokenCertificateNamePtr) ToTokenCertificateNamePtrOutput() TokenCertificateNamePtrOutput {
	return pulumi.ToOutput(in).(TokenCertificateNamePtrOutput)
}

func (in *tokenCertificateNamePtr) ToTokenCertificateNamePtrOutputWithContext(ctx context.Context) TokenCertificateNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TokenCertificateNamePtrOutput)
}

type TokenPasswordName string

const (
	TokenPasswordNamePassword1 = TokenPasswordName("password1")
	TokenPasswordNamePassword2 = TokenPasswordName("password2")
)

func (TokenPasswordName) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPasswordName)(nil)).Elem()
}

func (e TokenPasswordName) ToTokenPasswordNameOutput() TokenPasswordNameOutput {
	return pulumi.ToOutput(e).(TokenPasswordNameOutput)
}

func (e TokenPasswordName) ToTokenPasswordNameOutputWithContext(ctx context.Context) TokenPasswordNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TokenPasswordNameOutput)
}

func (e TokenPasswordName) ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput {
	return e.ToTokenPasswordNamePtrOutputWithContext(context.Background())
}

func (e TokenPasswordName) ToTokenPasswordNamePtrOutputWithContext(ctx context.Context) TokenPasswordNamePtrOutput {
	return TokenPasswordName(e).ToTokenPasswordNameOutputWithContext(ctx).ToTokenPasswordNamePtrOutputWithContext(ctx)
}

func (e TokenPasswordName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenPasswordName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenPasswordName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TokenPasswordName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TokenPasswordNameOutput struct{ *pulumi.OutputState }

func (TokenPasswordNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPasswordName)(nil)).Elem()
}

func (o TokenPasswordNameOutput) ToTokenPasswordNameOutput() TokenPasswordNameOutput {
	return o
}

func (o TokenPasswordNameOutput) ToTokenPasswordNameOutputWithContext(ctx context.Context) TokenPasswordNameOutput {
	return o
}

func (o TokenPasswordNameOutput) ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput {
	return o.ToTokenPasswordNamePtrOutputWithContext(context.Background())
}

func (o TokenPasswordNameOutput) ToTokenPasswordNamePtrOutputWithContext(ctx context.Context) TokenPasswordNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenPasswordName) *TokenPasswordName {
		return &v
	}).(TokenPasswordNamePtrOutput)
}

func (o TokenPasswordNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TokenPasswordNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenPasswordName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TokenPasswordNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenPasswordNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenPasswordName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TokenPasswordNamePtrOutput struct{ *pulumi.OutputState }

func (TokenPasswordNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenPasswordName)(nil)).Elem()
}

func (o TokenPasswordNamePtrOutput) ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput {
	return o
}

func (o TokenPasswordNamePtrOutput) ToTokenPasswordNamePtrOutputWithContext(ctx context.Context) TokenPasswordNamePtrOutput {
	return o
}

func (o TokenPasswordNamePtrOutput) Elem() TokenPasswordNameOutput {
	return o.ApplyT(func(v *TokenPasswordName) TokenPasswordName {
		if v != nil {
			return *v
		}
		var ret TokenPasswordName
		return ret
	}).(TokenPasswordNameOutput)
}

func (o TokenPasswordNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenPasswordNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TokenPasswordName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TokenPasswordNameInput interface {
	pulumi.Input

	ToTokenPasswordNameOutput() TokenPasswordNameOutput
	ToTokenPasswordNameOutputWithContext(context.Context) TokenPasswordNameOutput
}

var tokenPasswordNamePtrType = reflect.TypeOf((**TokenPasswordName)(nil)).Elem()

type TokenPasswordNamePtrInput interface {
	pulumi.Input

	ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput
	ToTokenPasswordNamePtrOutputWithContext(context.Context) TokenPasswordNamePtrOutput
}

type tokenPasswordNamePtr string

func TokenPasswordNamePtr(v string) TokenPasswordNamePtrInput {
	return (*tokenPasswordNamePtr)(&v)
}

func (*tokenPasswordNamePtr) ElementType() reflect.Type {
	return tokenPasswordNamePtrType
}

func (in *tokenPasswordNamePtr) ToTokenPasswordNamePtrOutput() TokenPasswordNamePtrOutput {
	return pulumi.ToOutput(in).(TokenPasswordNamePtrOutput)
}

func (in *tokenPasswordNamePtr) ToTokenPasswordNamePtrOutputWithContext(ctx context.Context) TokenPasswordNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TokenPasswordNamePtrOutput)
}

type TokenStatus string

const (
	TokenStatusEnabled  = TokenStatus("enabled")
	TokenStatusDisabled = TokenStatus("disabled")
)

func (TokenStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenStatus)(nil)).Elem()
}

func (e TokenStatus) ToTokenStatusOutput() TokenStatusOutput {
	return pulumi.ToOutput(e).(TokenStatusOutput)
}

func (e TokenStatus) ToTokenStatusOutputWithContext(ctx context.Context) TokenStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TokenStatusOutput)
}

func (e TokenStatus) ToTokenStatusPtrOutput() TokenStatusPtrOutput {
	return e.ToTokenStatusPtrOutputWithContext(context.Background())
}

func (e TokenStatus) ToTokenStatusPtrOutputWithContext(ctx context.Context) TokenStatusPtrOutput {
	return TokenStatus(e).ToTokenStatusOutputWithContext(ctx).ToTokenStatusPtrOutputWithContext(ctx)
}

func (e TokenStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TokenStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TokenStatusOutput struct{ *pulumi.OutputState }

func (TokenStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenStatus)(nil)).Elem()
}

func (o TokenStatusOutput) ToTokenStatusOutput() TokenStatusOutput {
	return o
}

func (o TokenStatusOutput) ToTokenStatusOutputWithContext(ctx context.Context) TokenStatusOutput {
	return o
}

func (o TokenStatusOutput) ToTokenStatusPtrOutput() TokenStatusPtrOutput {
	return o.ToTokenStatusPtrOutputWithContext(context.Background())
}

func (o TokenStatusOutput) ToTokenStatusPtrOutputWithContext(ctx context.Context) TokenStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenStatus) *TokenStatus {
		return &v
	}).(TokenStatusPtrOutput)
}

func (o TokenStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TokenStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TokenStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TokenStatusPtrOutput struct{ *pulumi.OutputState }

func (TokenStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenStatus)(nil)).Elem()
}

func (o TokenStatusPtrOutput) ToTokenStatusPtrOutput() TokenStatusPtrOutput {
	return o
}

func (o TokenStatusPtrOutput) ToTokenStatusPtrOutputWithContext(ctx context.Context) TokenStatusPtrOutput {
	return o
}

func (o TokenStatusPtrOutput) Elem() TokenStatusOutput {
	return o.ApplyT(func(v *TokenStatus) TokenStatus {
		if v != nil {
			return *v
		}
		var ret TokenStatus
		return ret
	}).(TokenStatusOutput)
}

func (o TokenStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TokenStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TokenStatusInput interface {
	pulumi.Input

	ToTokenStatusOutput() TokenStatusOutput
	ToTokenStatusOutputWithContext(context.Context) TokenStatusOutput
}

var tokenStatusPtrType = reflect.TypeOf((**TokenStatus)(nil)).Elem()

type TokenStatusPtrInput interface {
	pulumi.Input

	ToTokenStatusPtrOutput() TokenStatusPtrOutput
	ToTokenStatusPtrOutputWithContext(context.Context) TokenStatusPtrOutput
}

type tokenStatusPtr string

func TokenStatusPtr(v string) TokenStatusPtrInput {
	return (*tokenStatusPtr)(&v)
}

func (*tokenStatusPtr) ElementType() reflect.Type {
	return tokenStatusPtrType
}

func (in *tokenStatusPtr) ToTokenStatusPtrOutput() TokenStatusPtrOutput {
	return pulumi.ToOutput(in).(TokenStatusPtrOutput)
}

func (in *tokenStatusPtr) ToTokenStatusPtrOutputWithContext(ctx context.Context) TokenStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TokenStatusPtrOutput)
}

type TokenType string

const (
	TokenTypePAT   = TokenType("PAT")
	TokenTypeOAuth = TokenType("OAuth")
)

func (TokenType) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenType)(nil)).Elem()
}

func (e TokenType) ToTokenTypeOutput() TokenTypeOutput {
	return pulumi.ToOutput(e).(TokenTypeOutput)
}

func (e TokenType) ToTokenTypeOutputWithContext(ctx context.Context) TokenTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TokenTypeOutput)
}

func (e TokenType) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return e.ToTokenTypePtrOutputWithContext(context.Background())
}

func (e TokenType) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return TokenType(e).ToTokenTypeOutputWithContext(ctx).ToTokenTypePtrOutputWithContext(ctx)
}

func (e TokenType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TokenType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TokenTypeOutput struct{ *pulumi.OutputState }

func (TokenTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenType)(nil)).Elem()
}

func (o TokenTypeOutput) ToTokenTypeOutput() TokenTypeOutput {
	return o
}

func (o TokenTypeOutput) ToTokenTypeOutputWithContext(ctx context.Context) TokenTypeOutput {
	return o
}

func (o TokenTypeOutput) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return o.ToTokenTypePtrOutputWithContext(context.Background())
}

func (o TokenTypeOutput) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenType) *TokenType {
		return &v
	}).(TokenTypePtrOutput)
}

func (o TokenTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TokenTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TokenTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TokenTypePtrOutput struct{ *pulumi.OutputState }

func (TokenTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenType)(nil)).Elem()
}

func (o TokenTypePtrOutput) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return o
}

func (o TokenTypePtrOutput) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return o
}

func (o TokenTypePtrOutput) Elem() TokenTypeOutput {
	return o.ApplyT(func(v *TokenType) TokenType {
		if v != nil {
			return *v
		}
		var ret TokenType
		return ret
	}).(TokenTypeOutput)
}

func (o TokenTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TokenType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TokenTypeInput interface {
	pulumi.Input

	ToTokenTypeOutput() TokenTypeOutput
	ToTokenTypeOutputWithContext(context.Context) TokenTypeOutput
}

var tokenTypePtrType = reflect.TypeOf((**TokenType)(nil)).Elem()

type TokenTypePtrInput interface {
	pulumi.Input

	ToTokenTypePtrOutput() TokenTypePtrOutput
	ToTokenTypePtrOutputWithContext(context.Context) TokenTypePtrOutput
}

type tokenTypePtr string

func TokenTypePtr(v string) TokenTypePtrInput {
	return (*tokenTypePtr)(&v)
}

func (*tokenTypePtr) ElementType() reflect.Type {
	return tokenTypePtrType
}

func (in *tokenTypePtr) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return pulumi.ToOutput(in).(TokenTypePtrOutput)
}

func (in *tokenTypePtr) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TokenTypePtrOutput)
}

type TriggerStatus string

const (
	TriggerStatusDisabled = TriggerStatus("Disabled")
	TriggerStatusEnabled  = TriggerStatus("Enabled")
)

func (TriggerStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerStatus)(nil)).Elem()
}

func (e TriggerStatus) ToTriggerStatusOutput() TriggerStatusOutput {
	return pulumi.ToOutput(e).(TriggerStatusOutput)
}

func (e TriggerStatus) ToTriggerStatusOutputWithContext(ctx context.Context) TriggerStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TriggerStatusOutput)
}

func (e TriggerStatus) ToTriggerStatusPtrOutput() TriggerStatusPtrOutput {
	return e.ToTriggerStatusPtrOutputWithContext(context.Background())
}

func (e TriggerStatus) ToTriggerStatusPtrOutputWithContext(ctx context.Context) TriggerStatusPtrOutput {
	return TriggerStatus(e).ToTriggerStatusOutputWithContext(ctx).ToTriggerStatusPtrOutputWithContext(ctx)
}

func (e TriggerStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TriggerStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TriggerStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TriggerStatusOutput struct{ *pulumi.OutputState }

func (TriggerStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerStatus)(nil)).Elem()
}

func (o TriggerStatusOutput) ToTriggerStatusOutput() TriggerStatusOutput {
	return o
}

func (o TriggerStatusOutput) ToTriggerStatusOutputWithContext(ctx context.Context) TriggerStatusOutput {
	return o
}

func (o TriggerStatusOutput) ToTriggerStatusPtrOutput() TriggerStatusPtrOutput {
	return o.ToTriggerStatusPtrOutputWithContext(context.Background())
}

func (o TriggerStatusOutput) ToTriggerStatusPtrOutputWithContext(ctx context.Context) TriggerStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerStatus) *TriggerStatus {
		return &v
	}).(TriggerStatusPtrOutput)
}

func (o TriggerStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TriggerStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TriggerStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TriggerStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TriggerStatusPtrOutput struct{ *pulumi.OutputState }

func (TriggerStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerStatus)(nil)).Elem()
}

func (o TriggerStatusPtrOutput) ToTriggerStatusPtrOutput() TriggerStatusPtrOutput {
	return o
}

func (o TriggerStatusPtrOutput) ToTriggerStatusPtrOutputWithContext(ctx context.Context) TriggerStatusPtrOutput {
	return o
}

func (o TriggerStatusPtrOutput) Elem() TriggerStatusOutput {
	return o.ApplyT(func(v *TriggerStatus) TriggerStatus {
		if v != nil {
			return *v
		}
		var ret TriggerStatus
		return ret
	}).(TriggerStatusOutput)
}

func (o TriggerStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TriggerStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TriggerStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TriggerStatusInput interface {
	pulumi.Input

	ToTriggerStatusOutput() TriggerStatusOutput
	ToTriggerStatusOutputWithContext(context.Context) TriggerStatusOutput
}

var triggerStatusPtrType = reflect.TypeOf((**TriggerStatus)(nil)).Elem()

type TriggerStatusPtrInput interface {
	pulumi.Input

	ToTriggerStatusPtrOutput() TriggerStatusPtrOutput
	ToTriggerStatusPtrOutputWithContext(context.Context) TriggerStatusPtrOutput
}

type triggerStatusPtr string

func TriggerStatusPtr(v string) TriggerStatusPtrInput {
	return (*triggerStatusPtr)(&v)
}

func (*triggerStatusPtr) ElementType() reflect.Type {
	return triggerStatusPtrType
}

func (in *triggerStatusPtr) ToTriggerStatusPtrOutput() TriggerStatusPtrOutput {
	return pulumi.ToOutput(in).(TriggerStatusPtrOutput)
}

func (in *triggerStatusPtr) ToTriggerStatusPtrOutputWithContext(ctx context.Context) TriggerStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TriggerStatusPtrOutput)
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

type UpdateTriggerPayloadType string

const (
	UpdateTriggerPayloadTypeDefault = UpdateTriggerPayloadType("Default")
	UpdateTriggerPayloadTypeToken   = UpdateTriggerPayloadType("Token")
)

func (UpdateTriggerPayloadType) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateTriggerPayloadType)(nil)).Elem()
}

func (e UpdateTriggerPayloadType) ToUpdateTriggerPayloadTypeOutput() UpdateTriggerPayloadTypeOutput {
	return pulumi.ToOutput(e).(UpdateTriggerPayloadTypeOutput)
}

func (e UpdateTriggerPayloadType) ToUpdateTriggerPayloadTypeOutputWithContext(ctx context.Context) UpdateTriggerPayloadTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UpdateTriggerPayloadTypeOutput)
}

func (e UpdateTriggerPayloadType) ToUpdateTriggerPayloadTypePtrOutput() UpdateTriggerPayloadTypePtrOutput {
	return e.ToUpdateTriggerPayloadTypePtrOutputWithContext(context.Background())
}

func (e UpdateTriggerPayloadType) ToUpdateTriggerPayloadTypePtrOutputWithContext(ctx context.Context) UpdateTriggerPayloadTypePtrOutput {
	return UpdateTriggerPayloadType(e).ToUpdateTriggerPayloadTypeOutputWithContext(ctx).ToUpdateTriggerPayloadTypePtrOutputWithContext(ctx)
}

func (e UpdateTriggerPayloadType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpdateTriggerPayloadType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpdateTriggerPayloadType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UpdateTriggerPayloadType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UpdateTriggerPayloadTypeOutput struct{ *pulumi.OutputState }

func (UpdateTriggerPayloadTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateTriggerPayloadType)(nil)).Elem()
}

func (o UpdateTriggerPayloadTypeOutput) ToUpdateTriggerPayloadTypeOutput() UpdateTriggerPayloadTypeOutput {
	return o
}

func (o UpdateTriggerPayloadTypeOutput) ToUpdateTriggerPayloadTypeOutputWithContext(ctx context.Context) UpdateTriggerPayloadTypeOutput {
	return o
}

func (o UpdateTriggerPayloadTypeOutput) ToUpdateTriggerPayloadTypePtrOutput() UpdateTriggerPayloadTypePtrOutput {
	return o.ToUpdateTriggerPayloadTypePtrOutputWithContext(context.Background())
}

func (o UpdateTriggerPayloadTypeOutput) ToUpdateTriggerPayloadTypePtrOutputWithContext(ctx context.Context) UpdateTriggerPayloadTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpdateTriggerPayloadType) *UpdateTriggerPayloadType {
		return &v
	}).(UpdateTriggerPayloadTypePtrOutput)
}

func (o UpdateTriggerPayloadTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UpdateTriggerPayloadTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpdateTriggerPayloadType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UpdateTriggerPayloadTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpdateTriggerPayloadTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpdateTriggerPayloadType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UpdateTriggerPayloadTypePtrOutput struct{ *pulumi.OutputState }

func (UpdateTriggerPayloadTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateTriggerPayloadType)(nil)).Elem()
}

func (o UpdateTriggerPayloadTypePtrOutput) ToUpdateTriggerPayloadTypePtrOutput() UpdateTriggerPayloadTypePtrOutput {
	return o
}

func (o UpdateTriggerPayloadTypePtrOutput) ToUpdateTriggerPayloadTypePtrOutputWithContext(ctx context.Context) UpdateTriggerPayloadTypePtrOutput {
	return o
}

func (o UpdateTriggerPayloadTypePtrOutput) Elem() UpdateTriggerPayloadTypeOutput {
	return o.ApplyT(func(v *UpdateTriggerPayloadType) UpdateTriggerPayloadType {
		if v != nil {
			return *v
		}
		var ret UpdateTriggerPayloadType
		return ret
	}).(UpdateTriggerPayloadTypeOutput)
}

func (o UpdateTriggerPayloadTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpdateTriggerPayloadTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UpdateTriggerPayloadType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UpdateTriggerPayloadTypeInput interface {
	pulumi.Input

	ToUpdateTriggerPayloadTypeOutput() UpdateTriggerPayloadTypeOutput
	ToUpdateTriggerPayloadTypeOutputWithContext(context.Context) UpdateTriggerPayloadTypeOutput
}

var updateTriggerPayloadTypePtrType = reflect.TypeOf((**UpdateTriggerPayloadType)(nil)).Elem()

type UpdateTriggerPayloadTypePtrInput interface {
	pulumi.Input

	ToUpdateTriggerPayloadTypePtrOutput() UpdateTriggerPayloadTypePtrOutput
	ToUpdateTriggerPayloadTypePtrOutputWithContext(context.Context) UpdateTriggerPayloadTypePtrOutput
}

type updateTriggerPayloadTypePtr string

func UpdateTriggerPayloadTypePtr(v string) UpdateTriggerPayloadTypePtrInput {
	return (*updateTriggerPayloadTypePtr)(&v)
}

func (*updateTriggerPayloadTypePtr) ElementType() reflect.Type {
	return updateTriggerPayloadTypePtrType
}

func (in *updateTriggerPayloadTypePtr) ToUpdateTriggerPayloadTypePtrOutput() UpdateTriggerPayloadTypePtrOutput {
	return pulumi.ToOutput(in).(UpdateTriggerPayloadTypePtrOutput)
}

func (in *updateTriggerPayloadTypePtr) ToUpdateTriggerPayloadTypePtrOutputWithContext(ctx context.Context) UpdateTriggerPayloadTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UpdateTriggerPayloadTypePtrOutput)
}

type Variant string

const (
	VariantV6 = Variant("v6")
	VariantV7 = Variant("v7")
	VariantV8 = Variant("v8")
)

func (Variant) ElementType() reflect.Type {
	return reflect.TypeOf((*Variant)(nil)).Elem()
}

func (e Variant) ToVariantOutput() VariantOutput {
	return pulumi.ToOutput(e).(VariantOutput)
}

func (e Variant) ToVariantOutputWithContext(ctx context.Context) VariantOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VariantOutput)
}

func (e Variant) ToVariantPtrOutput() VariantPtrOutput {
	return e.ToVariantPtrOutputWithContext(context.Background())
}

func (e Variant) ToVariantPtrOutputWithContext(ctx context.Context) VariantPtrOutput {
	return Variant(e).ToVariantOutputWithContext(ctx).ToVariantPtrOutputWithContext(ctx)
}

func (e Variant) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Variant) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Variant) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Variant) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VariantOutput struct{ *pulumi.OutputState }

func (VariantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Variant)(nil)).Elem()
}

func (o VariantOutput) ToVariantOutput() VariantOutput {
	return o
}

func (o VariantOutput) ToVariantOutputWithContext(ctx context.Context) VariantOutput {
	return o
}

func (o VariantOutput) ToVariantPtrOutput() VariantPtrOutput {
	return o.ToVariantPtrOutputWithContext(context.Background())
}

func (o VariantOutput) ToVariantPtrOutputWithContext(ctx context.Context) VariantPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Variant) *Variant {
		return &v
	}).(VariantPtrOutput)
}

func (o VariantOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VariantOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Variant) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VariantOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VariantOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Variant) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VariantPtrOutput struct{ *pulumi.OutputState }

func (VariantPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Variant)(nil)).Elem()
}

func (o VariantPtrOutput) ToVariantPtrOutput() VariantPtrOutput {
	return o
}

func (o VariantPtrOutput) ToVariantPtrOutputWithContext(ctx context.Context) VariantPtrOutput {
	return o
}

func (o VariantPtrOutput) Elem() VariantOutput {
	return o.ApplyT(func(v *Variant) Variant {
		if v != nil {
			return *v
		}
		var ret Variant
		return ret
	}).(VariantOutput)
}

func (o VariantPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VariantPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Variant) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VariantInput interface {
	pulumi.Input

	ToVariantOutput() VariantOutput
	ToVariantOutputWithContext(context.Context) VariantOutput
}

var variantPtrType = reflect.TypeOf((**Variant)(nil)).Elem()

type VariantPtrInput interface {
	pulumi.Input

	ToVariantPtrOutput() VariantPtrOutput
	ToVariantPtrOutputWithContext(context.Context) VariantPtrOutput
}

type variantPtr string

func VariantPtr(v string) VariantPtrInput {
	return (*variantPtr)(&v)
}

func (*variantPtr) ElementType() reflect.Type {
	return variantPtrType
}

func (in *variantPtr) ToVariantPtrOutput() VariantPtrOutput {
	return pulumi.ToOutput(in).(VariantPtrOutput)
}

func (in *variantPtr) ToVariantPtrOutputWithContext(ctx context.Context) VariantPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VariantPtrOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*ActionInput)(nil)).Elem(), Action("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*ActionPtrInput)(nil)).Elem(), Action("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsRequiredInput)(nil)).Elem(), ActionsRequired("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*ActionsRequiredPtrInput)(nil)).Elem(), ActionsRequired("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*ArchitectureInput)(nil)).Elem(), Architecture("amd64"))
	pulumi.RegisterInputType(reflect.TypeOf((*ArchitecturePtrInput)(nil)).Elem(), Architecture("amd64"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogStatusInput)(nil)).Elem(), AuditLogStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogStatusPtrInput)(nil)).Elem(), AuditLogStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*BaseImageTriggerTypeInput)(nil)).Elem(), BaseImageTriggerType("All"))
	pulumi.RegisterInputType(reflect.TypeOf((*BaseImageTriggerTypePtrInput)(nil)).Elem(), BaseImageTriggerType("All"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectedRegistryModeInput)(nil)).Elem(), ConnectedRegistryMode("Registry"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectedRegistryModePtrInput)(nil)).Elem(), ConnectedRegistryMode("Registry"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStatusInput)(nil)).Elem(), ConnectionStatus("Approved"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionStatusPtrInput)(nil)).Elem(), ConnectionStatus("Approved"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultActionInput)(nil)).Elem(), DefaultAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultActionPtrInput)(nil)).Elem(), DefaultAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*LogLevelInput)(nil)).Elem(), LogLevel("Debug"))
	pulumi.RegisterInputType(reflect.TypeOf((*LogLevelPtrInput)(nil)).Elem(), LogLevel("Debug"))
	pulumi.RegisterInputType(reflect.TypeOf((*OSInput)(nil)).Elem(), OS("Windows"))
	pulumi.RegisterInputType(reflect.TypeOf((*OSPtrInput)(nil)).Elem(), OS("Windows"))
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineOptionsInput)(nil)).Elem(), PipelineOptions("OverwriteTags"))
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineOptionsPtrInput)(nil)).Elem(), PipelineOptions("OverwriteTags"))
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineRunSourceTypeInput)(nil)).Elem(), PipelineRunSourceType("AzureStorageBlob"))
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineRunSourceTypePtrInput)(nil)).Elem(), PipelineRunSourceType("AzureStorageBlob"))
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineRunTargetTypeInput)(nil)).Elem(), PipelineRunTargetType("AzureStorageBlob"))
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineRunTargetTypePtrInput)(nil)).Elem(), PipelineRunTargetType("AzureStorageBlob"))
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineSourceTypeInput)(nil)).Elem(), PipelineSourceType("AzureStorageBlobContainer"))
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineSourceTypePtrInput)(nil)).Elem(), PipelineSourceType("AzureStorageBlobContainer"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyStatusInput)(nil)).Elem(), PolicyStatus("enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyStatusPtrInput)(nil)).Elem(), PolicyStatus("enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecretObjectTypeInput)(nil)).Elem(), SecretObjectType("Opaque"))
	pulumi.RegisterInputType(reflect.TypeOf((*SecretObjectTypePtrInput)(nil)).Elem(), SecretObjectType("Opaque"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNameInput)(nil)).Elem(), SkuName("Classic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNamePtrInput)(nil)).Elem(), SkuName("Classic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceControlTypeInput)(nil)).Elem(), SourceControlType("Github"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceControlTypePtrInput)(nil)).Elem(), SourceControlType("Github"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRegistryLoginModeInput)(nil)).Elem(), SourceRegistryLoginMode("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceRegistryLoginModePtrInput)(nil)).Elem(), SourceRegistryLoginMode("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTriggerEventInput)(nil)).Elem(), SourceTriggerEvent("commit"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTriggerEventPtrInput)(nil)).Elem(), SourceTriggerEvent("commit"))
	pulumi.RegisterInputType(reflect.TypeOf((*StepTypeInput)(nil)).Elem(), StepType("Docker"))
	pulumi.RegisterInputType(reflect.TypeOf((*StepTypePtrInput)(nil)).Elem(), StepType("Docker"))
	pulumi.RegisterInputType(reflect.TypeOf((*TaskStatusInput)(nil)).Elem(), TaskStatus("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TaskStatusPtrInput)(nil)).Elem(), TaskStatus("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenCertificateNameInput)(nil)).Elem(), TokenCertificateName("certificate1"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenCertificateNamePtrInput)(nil)).Elem(), TokenCertificateName("certificate1"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenPasswordNameInput)(nil)).Elem(), TokenPasswordName("password1"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenPasswordNamePtrInput)(nil)).Elem(), TokenPasswordName("password1"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenStatusInput)(nil)).Elem(), TokenStatus("enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenStatusPtrInput)(nil)).Elem(), TokenStatus("enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenTypeInput)(nil)).Elem(), TokenType("PAT"))
	pulumi.RegisterInputType(reflect.TypeOf((*TokenTypePtrInput)(nil)).Elem(), TokenType("PAT"))
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerStatusInput)(nil)).Elem(), TriggerStatus("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerStatusPtrInput)(nil)).Elem(), TriggerStatus("Disabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrustPolicyTypeInput)(nil)).Elem(), TrustPolicyType("Notary"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrustPolicyTypePtrInput)(nil)).Elem(), TrustPolicyType("Notary"))
	pulumi.RegisterInputType(reflect.TypeOf((*UpdateTriggerPayloadTypeInput)(nil)).Elem(), UpdateTriggerPayloadType("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*UpdateTriggerPayloadTypePtrInput)(nil)).Elem(), UpdateTriggerPayloadType("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*VariantInput)(nil)).Elem(), Variant("v6"))
	pulumi.RegisterInputType(reflect.TypeOf((*VariantPtrInput)(nil)).Elem(), Variant("v6"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookActionInput)(nil)).Elem(), WebhookAction("push"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookActionPtrInput)(nil)).Elem(), WebhookAction("push"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookStatusInput)(nil)).Elem(), WebhookStatus("enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookStatusPtrInput)(nil)).Elem(), WebhookStatus("enabled"))
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(ActionsRequiredOutput{})
	pulumi.RegisterOutputType(ActionsRequiredPtrOutput{})
	pulumi.RegisterOutputType(ArchitectureOutput{})
	pulumi.RegisterOutputType(ArchitecturePtrOutput{})
	pulumi.RegisterOutputType(AuditLogStatusOutput{})
	pulumi.RegisterOutputType(AuditLogStatusPtrOutput{})
	pulumi.RegisterOutputType(BaseImageTriggerTypeOutput{})
	pulumi.RegisterOutputType(BaseImageTriggerTypePtrOutput{})
	pulumi.RegisterOutputType(ConnectedRegistryModeOutput{})
	pulumi.RegisterOutputType(ConnectedRegistryModePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStatusOutput{})
	pulumi.RegisterOutputType(ConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(DefaultActionOutput{})
	pulumi.RegisterOutputType(DefaultActionPtrOutput{})
	pulumi.RegisterOutputType(LogLevelOutput{})
	pulumi.RegisterOutputType(LogLevelPtrOutput{})
	pulumi.RegisterOutputType(OSOutput{})
	pulumi.RegisterOutputType(OSPtrOutput{})
	pulumi.RegisterOutputType(PipelineOptionsOutput{})
	pulumi.RegisterOutputType(PipelineOptionsPtrOutput{})
	pulumi.RegisterOutputType(PipelineRunSourceTypeOutput{})
	pulumi.RegisterOutputType(PipelineRunSourceTypePtrOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetTypeOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetTypePtrOutput{})
	pulumi.RegisterOutputType(PipelineSourceTypeOutput{})
	pulumi.RegisterOutputType(PipelineSourceTypePtrOutput{})
	pulumi.RegisterOutputType(PolicyStatusOutput{})
	pulumi.RegisterOutputType(PolicyStatusPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SecretObjectTypeOutput{})
	pulumi.RegisterOutputType(SecretObjectTypePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(SourceControlTypeOutput{})
	pulumi.RegisterOutputType(SourceControlTypePtrOutput{})
	pulumi.RegisterOutputType(SourceRegistryLoginModeOutput{})
	pulumi.RegisterOutputType(SourceRegistryLoginModePtrOutput{})
	pulumi.RegisterOutputType(SourceTriggerEventOutput{})
	pulumi.RegisterOutputType(SourceTriggerEventPtrOutput{})
	pulumi.RegisterOutputType(StepTypeOutput{})
	pulumi.RegisterOutputType(StepTypePtrOutput{})
	pulumi.RegisterOutputType(TaskStatusOutput{})
	pulumi.RegisterOutputType(TaskStatusPtrOutput{})
	pulumi.RegisterOutputType(TokenCertificateNameOutput{})
	pulumi.RegisterOutputType(TokenCertificateNamePtrOutput{})
	pulumi.RegisterOutputType(TokenPasswordNameOutput{})
	pulumi.RegisterOutputType(TokenPasswordNamePtrOutput{})
	pulumi.RegisterOutputType(TokenStatusOutput{})
	pulumi.RegisterOutputType(TokenStatusPtrOutput{})
	pulumi.RegisterOutputType(TokenTypeOutput{})
	pulumi.RegisterOutputType(TokenTypePtrOutput{})
	pulumi.RegisterOutputType(TriggerStatusOutput{})
	pulumi.RegisterOutputType(TriggerStatusPtrOutput{})
	pulumi.RegisterOutputType(TrustPolicyTypeOutput{})
	pulumi.RegisterOutputType(TrustPolicyTypePtrOutput{})
	pulumi.RegisterOutputType(UpdateTriggerPayloadTypeOutput{})
	pulumi.RegisterOutputType(UpdateTriggerPayloadTypePtrOutput{})
	pulumi.RegisterOutputType(VariantOutput{})
	pulumi.RegisterOutputType(VariantPtrOutput{})
	pulumi.RegisterOutputType(WebhookActionOutput{})
	pulumi.RegisterOutputType(WebhookActionPtrOutput{})
	pulumi.RegisterOutputType(WebhookStatusOutput{})
	pulumi.RegisterOutputType(WebhookStatusPtrOutput{})
}
