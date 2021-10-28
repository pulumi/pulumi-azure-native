


package v20210801preview

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

type ConnectedRegistryMode string

const (
	ConnectedRegistryModeReadWrite = ConnectedRegistryMode("ReadWrite")
	ConnectedRegistryModeReadOnly  = ConnectedRegistryMode("ReadOnly")
	ConnectedRegistryModeRegistry  = ConnectedRegistryMode("Registry")
	ConnectedRegistryModeMirror    = ConnectedRegistryMode("Mirror")
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

type EncryptionStatus string

const (
	EncryptionStatusEnabled  = EncryptionStatus("enabled")
	EncryptionStatusDisabled = EncryptionStatus("disabled")
)

func (EncryptionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionStatus)(nil)).Elem()
}

func (e EncryptionStatus) ToEncryptionStatusOutput() EncryptionStatusOutput {
	return pulumi.ToOutput(e).(EncryptionStatusOutput)
}

func (e EncryptionStatus) ToEncryptionStatusOutputWithContext(ctx context.Context) EncryptionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionStatusOutput)
}

func (e EncryptionStatus) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return e.ToEncryptionStatusPtrOutputWithContext(context.Background())
}

func (e EncryptionStatus) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return EncryptionStatus(e).ToEncryptionStatusOutputWithContext(ctx).ToEncryptionStatusPtrOutputWithContext(ctx)
}

func (e EncryptionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionStatusOutput struct{ *pulumi.OutputState }

func (EncryptionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionStatus)(nil)).Elem()
}

func (o EncryptionStatusOutput) ToEncryptionStatusOutput() EncryptionStatusOutput {
	return o
}

func (o EncryptionStatusOutput) ToEncryptionStatusOutputWithContext(ctx context.Context) EncryptionStatusOutput {
	return o
}

func (o EncryptionStatusOutput) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return o.ToEncryptionStatusPtrOutputWithContext(context.Background())
}

func (o EncryptionStatusOutput) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionStatus) *EncryptionStatus {
		return &v
	}).(EncryptionStatusPtrOutput)
}

func (o EncryptionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionStatusPtrOutput struct{ *pulumi.OutputState }

func (EncryptionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionStatus)(nil)).Elem()
}

func (o EncryptionStatusPtrOutput) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return o
}

func (o EncryptionStatusPtrOutput) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return o
}

func (o EncryptionStatusPtrOutput) Elem() EncryptionStatusOutput {
	return o.ApplyT(func(v *EncryptionStatus) EncryptionStatus {
		if v != nil {
			return *v
		}
		var ret EncryptionStatus
		return ret
	}).(EncryptionStatusOutput)
}

func (o EncryptionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionStatusInput interface {
	pulumi.Input

	ToEncryptionStatusOutput() EncryptionStatusOutput
	ToEncryptionStatusOutputWithContext(context.Context) EncryptionStatusOutput
}

var encryptionStatusPtrType = reflect.TypeOf((**EncryptionStatus)(nil)).Elem()

type EncryptionStatusPtrInput interface {
	pulumi.Input

	ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput
	ToEncryptionStatusPtrOutputWithContext(context.Context) EncryptionStatusPtrOutput
}

type encryptionStatusPtr string

func EncryptionStatusPtr(v string) EncryptionStatusPtrInput {
	return (*encryptionStatusPtr)(&v)
}

func (*encryptionStatusPtr) ElementType() reflect.Type {
	return encryptionStatusPtrType
}

func (in *encryptionStatusPtr) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return pulumi.ToOutput(in).(EncryptionStatusPtrOutput)
}

func (in *encryptionStatusPtr) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionStatusPtrOutput)
}

type ExportPolicyStatus string

const (
	ExportPolicyStatusEnabled  = ExportPolicyStatus("enabled")
	ExportPolicyStatusDisabled = ExportPolicyStatus("disabled")
)

func (ExportPolicyStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyStatus)(nil)).Elem()
}

func (e ExportPolicyStatus) ToExportPolicyStatusOutput() ExportPolicyStatusOutput {
	return pulumi.ToOutput(e).(ExportPolicyStatusOutput)
}

func (e ExportPolicyStatus) ToExportPolicyStatusOutputWithContext(ctx context.Context) ExportPolicyStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExportPolicyStatusOutput)
}

func (e ExportPolicyStatus) ToExportPolicyStatusPtrOutput() ExportPolicyStatusPtrOutput {
	return e.ToExportPolicyStatusPtrOutputWithContext(context.Background())
}

func (e ExportPolicyStatus) ToExportPolicyStatusPtrOutputWithContext(ctx context.Context) ExportPolicyStatusPtrOutput {
	return ExportPolicyStatus(e).ToExportPolicyStatusOutputWithContext(ctx).ToExportPolicyStatusPtrOutputWithContext(ctx)
}

func (e ExportPolicyStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExportPolicyStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExportPolicyStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExportPolicyStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExportPolicyStatusOutput struct{ *pulumi.OutputState }

func (ExportPolicyStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPolicyStatus)(nil)).Elem()
}

func (o ExportPolicyStatusOutput) ToExportPolicyStatusOutput() ExportPolicyStatusOutput {
	return o
}

func (o ExportPolicyStatusOutput) ToExportPolicyStatusOutputWithContext(ctx context.Context) ExportPolicyStatusOutput {
	return o
}

func (o ExportPolicyStatusOutput) ToExportPolicyStatusPtrOutput() ExportPolicyStatusPtrOutput {
	return o.ToExportPolicyStatusPtrOutputWithContext(context.Background())
}

func (o ExportPolicyStatusOutput) ToExportPolicyStatusPtrOutputWithContext(ctx context.Context) ExportPolicyStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportPolicyStatus) *ExportPolicyStatus {
		return &v
	}).(ExportPolicyStatusPtrOutput)
}

func (o ExportPolicyStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExportPolicyStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExportPolicyStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExportPolicyStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExportPolicyStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExportPolicyStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExportPolicyStatusPtrOutput struct{ *pulumi.OutputState }

func (ExportPolicyStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportPolicyStatus)(nil)).Elem()
}

func (o ExportPolicyStatusPtrOutput) ToExportPolicyStatusPtrOutput() ExportPolicyStatusPtrOutput {
	return o
}

func (o ExportPolicyStatusPtrOutput) ToExportPolicyStatusPtrOutputWithContext(ctx context.Context) ExportPolicyStatusPtrOutput {
	return o
}

func (o ExportPolicyStatusPtrOutput) Elem() ExportPolicyStatusOutput {
	return o.ApplyT(func(v *ExportPolicyStatus) ExportPolicyStatus {
		if v != nil {
			return *v
		}
		var ret ExportPolicyStatus
		return ret
	}).(ExportPolicyStatusOutput)
}

func (o ExportPolicyStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExportPolicyStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExportPolicyStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExportPolicyStatusInput interface {
	pulumi.Input

	ToExportPolicyStatusOutput() ExportPolicyStatusOutput
	ToExportPolicyStatusOutputWithContext(context.Context) ExportPolicyStatusOutput
}

var exportPolicyStatusPtrType = reflect.TypeOf((**ExportPolicyStatus)(nil)).Elem()

type ExportPolicyStatusPtrInput interface {
	pulumi.Input

	ToExportPolicyStatusPtrOutput() ExportPolicyStatusPtrOutput
	ToExportPolicyStatusPtrOutputWithContext(context.Context) ExportPolicyStatusPtrOutput
}

type exportPolicyStatusPtr string

func ExportPolicyStatusPtr(v string) ExportPolicyStatusPtrInput {
	return (*exportPolicyStatusPtr)(&v)
}

func (*exportPolicyStatusPtr) ElementType() reflect.Type {
	return exportPolicyStatusPtrType
}

func (in *exportPolicyStatusPtr) ToExportPolicyStatusPtrOutput() ExportPolicyStatusPtrOutput {
	return pulumi.ToOutput(in).(ExportPolicyStatusPtrOutput)
}

func (in *exportPolicyStatusPtr) ToExportPolicyStatusPtrOutputWithContext(ctx context.Context) ExportPolicyStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExportPolicyStatusPtrOutput)
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

type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices = NetworkRuleBypassOptions("AzureServices")
	NetworkRuleBypassOptionsNone          = NetworkRuleBypassOptions("None")
)

func (NetworkRuleBypassOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleBypassOptions)(nil)).Elem()
}

func (e NetworkRuleBypassOptions) ToNetworkRuleBypassOptionsOutput() NetworkRuleBypassOptionsOutput {
	return pulumi.ToOutput(e).(NetworkRuleBypassOptionsOutput)
}

func (e NetworkRuleBypassOptions) ToNetworkRuleBypassOptionsOutputWithContext(ctx context.Context) NetworkRuleBypassOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkRuleBypassOptionsOutput)
}

func (e NetworkRuleBypassOptions) ToNetworkRuleBypassOptionsPtrOutput() NetworkRuleBypassOptionsPtrOutput {
	return e.ToNetworkRuleBypassOptionsPtrOutputWithContext(context.Background())
}

func (e NetworkRuleBypassOptions) ToNetworkRuleBypassOptionsPtrOutputWithContext(ctx context.Context) NetworkRuleBypassOptionsPtrOutput {
	return NetworkRuleBypassOptions(e).ToNetworkRuleBypassOptionsOutputWithContext(ctx).ToNetworkRuleBypassOptionsPtrOutputWithContext(ctx)
}

func (e NetworkRuleBypassOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkRuleBypassOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkRuleBypassOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkRuleBypassOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkRuleBypassOptionsOutput struct{ *pulumi.OutputState }

func (NetworkRuleBypassOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleBypassOptions)(nil)).Elem()
}

func (o NetworkRuleBypassOptionsOutput) ToNetworkRuleBypassOptionsOutput() NetworkRuleBypassOptionsOutput {
	return o
}

func (o NetworkRuleBypassOptionsOutput) ToNetworkRuleBypassOptionsOutputWithContext(ctx context.Context) NetworkRuleBypassOptionsOutput {
	return o
}

func (o NetworkRuleBypassOptionsOutput) ToNetworkRuleBypassOptionsPtrOutput() NetworkRuleBypassOptionsPtrOutput {
	return o.ToNetworkRuleBypassOptionsPtrOutputWithContext(context.Background())
}

func (o NetworkRuleBypassOptionsOutput) ToNetworkRuleBypassOptionsPtrOutputWithContext(ctx context.Context) NetworkRuleBypassOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleBypassOptions) *NetworkRuleBypassOptions {
		return &v
	}).(NetworkRuleBypassOptionsPtrOutput)
}

func (o NetworkRuleBypassOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkRuleBypassOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkRuleBypassOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkRuleBypassOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkRuleBypassOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkRuleBypassOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkRuleBypassOptionsPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleBypassOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleBypassOptions)(nil)).Elem()
}

func (o NetworkRuleBypassOptionsPtrOutput) ToNetworkRuleBypassOptionsPtrOutput() NetworkRuleBypassOptionsPtrOutput {
	return o
}

func (o NetworkRuleBypassOptionsPtrOutput) ToNetworkRuleBypassOptionsPtrOutputWithContext(ctx context.Context) NetworkRuleBypassOptionsPtrOutput {
	return o
}

func (o NetworkRuleBypassOptionsPtrOutput) Elem() NetworkRuleBypassOptionsOutput {
	return o.ApplyT(func(v *NetworkRuleBypassOptions) NetworkRuleBypassOptions {
		if v != nil {
			return *v
		}
		var ret NetworkRuleBypassOptions
		return ret
	}).(NetworkRuleBypassOptionsOutput)
}

func (o NetworkRuleBypassOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkRuleBypassOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkRuleBypassOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkRuleBypassOptionsInput interface {
	pulumi.Input

	ToNetworkRuleBypassOptionsOutput() NetworkRuleBypassOptionsOutput
	ToNetworkRuleBypassOptionsOutputWithContext(context.Context) NetworkRuleBypassOptionsOutput
}

var networkRuleBypassOptionsPtrType = reflect.TypeOf((**NetworkRuleBypassOptions)(nil)).Elem()

type NetworkRuleBypassOptionsPtrInput interface {
	pulumi.Input

	ToNetworkRuleBypassOptionsPtrOutput() NetworkRuleBypassOptionsPtrOutput
	ToNetworkRuleBypassOptionsPtrOutputWithContext(context.Context) NetworkRuleBypassOptionsPtrOutput
}

type networkRuleBypassOptionsPtr string

func NetworkRuleBypassOptionsPtr(v string) NetworkRuleBypassOptionsPtrInput {
	return (*networkRuleBypassOptionsPtr)(&v)
}

func (*networkRuleBypassOptionsPtr) ElementType() reflect.Type {
	return networkRuleBypassOptionsPtrType
}

func (in *networkRuleBypassOptionsPtr) ToNetworkRuleBypassOptionsPtrOutput() NetworkRuleBypassOptionsPtrOutput {
	return pulumi.ToOutput(in).(NetworkRuleBypassOptionsPtrOutput)
}

func (in *networkRuleBypassOptionsPtr) ToNetworkRuleBypassOptionsPtrOutputWithContext(ctx context.Context) NetworkRuleBypassOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkRuleBypassOptionsPtrOutput)
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

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func (PublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return e.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return PublicNetworkAccess(e).ToPublicNetworkAccessOutputWithContext(ctx).ToPublicNetworkAccessPtrOutputWithContext(ctx)
}

func (e PublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccess) *PublicNetworkAccess {
		return &v
	}).(PublicNetworkAccessPtrOutput)
}

func (o PublicNetworkAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) Elem() PublicNetworkAccessOutput {
	return o.ApplyT(func(v *PublicNetworkAccess) PublicNetworkAccess {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccess
		return ret
	}).(PublicNetworkAccessOutput)
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessInput interface {
	pulumi.Input

	ToPublicNetworkAccessOutput() PublicNetworkAccessOutput
	ToPublicNetworkAccessOutputWithContext(context.Context) PublicNetworkAccessOutput
}

var publicNetworkAccessPtrType = reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()

type PublicNetworkAccessPtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput
	ToPublicNetworkAccessPtrOutputWithContext(context.Context) PublicNetworkAccessPtrOutput
}

type publicNetworkAccessPtr string

func PublicNetworkAccessPtr(v string) PublicNetworkAccessPtrInput {
	return (*publicNetworkAccessPtr)(&v)
}

func (*publicNetworkAccessPtr) ElementType() reflect.Type {
	return publicNetworkAccessPtrType
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessPtrOutput)
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessPtrOutput)
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

type TriggerStatus string

const (
	TriggerStatusEnabled  = TriggerStatus("Enabled")
	TriggerStatusDisabled = TriggerStatus("Disabled")
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

type ZoneRedundancy string

const (
	ZoneRedundancyEnabled  = ZoneRedundancy("Enabled")
	ZoneRedundancyDisabled = ZoneRedundancy("Disabled")
)

func (ZoneRedundancy) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneRedundancy)(nil)).Elem()
}

func (e ZoneRedundancy) ToZoneRedundancyOutput() ZoneRedundancyOutput {
	return pulumi.ToOutput(e).(ZoneRedundancyOutput)
}

func (e ZoneRedundancy) ToZoneRedundancyOutputWithContext(ctx context.Context) ZoneRedundancyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ZoneRedundancyOutput)
}

func (e ZoneRedundancy) ToZoneRedundancyPtrOutput() ZoneRedundancyPtrOutput {
	return e.ToZoneRedundancyPtrOutputWithContext(context.Background())
}

func (e ZoneRedundancy) ToZoneRedundancyPtrOutputWithContext(ctx context.Context) ZoneRedundancyPtrOutput {
	return ZoneRedundancy(e).ToZoneRedundancyOutputWithContext(ctx).ToZoneRedundancyPtrOutputWithContext(ctx)
}

func (e ZoneRedundancy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZoneRedundancy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZoneRedundancy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ZoneRedundancy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ZoneRedundancyOutput struct{ *pulumi.OutputState }

func (ZoneRedundancyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneRedundancy)(nil)).Elem()
}

func (o ZoneRedundancyOutput) ToZoneRedundancyOutput() ZoneRedundancyOutput {
	return o
}

func (o ZoneRedundancyOutput) ToZoneRedundancyOutputWithContext(ctx context.Context) ZoneRedundancyOutput {
	return o
}

func (o ZoneRedundancyOutput) ToZoneRedundancyPtrOutput() ZoneRedundancyPtrOutput {
	return o.ToZoneRedundancyPtrOutputWithContext(context.Background())
}

func (o ZoneRedundancyOutput) ToZoneRedundancyPtrOutputWithContext(ctx context.Context) ZoneRedundancyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ZoneRedundancy) *ZoneRedundancy {
		return &v
	}).(ZoneRedundancyPtrOutput)
}

func (o ZoneRedundancyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ZoneRedundancyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZoneRedundancy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ZoneRedundancyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZoneRedundancyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZoneRedundancy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ZoneRedundancyPtrOutput struct{ *pulumi.OutputState }

func (ZoneRedundancyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneRedundancy)(nil)).Elem()
}

func (o ZoneRedundancyPtrOutput) ToZoneRedundancyPtrOutput() ZoneRedundancyPtrOutput {
	return o
}

func (o ZoneRedundancyPtrOutput) ToZoneRedundancyPtrOutputWithContext(ctx context.Context) ZoneRedundancyPtrOutput {
	return o
}

func (o ZoneRedundancyPtrOutput) Elem() ZoneRedundancyOutput {
	return o.ApplyT(func(v *ZoneRedundancy) ZoneRedundancy {
		if v != nil {
			return *v
		}
		var ret ZoneRedundancy
		return ret
	}).(ZoneRedundancyOutput)
}

func (o ZoneRedundancyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZoneRedundancyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ZoneRedundancy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ZoneRedundancyInput interface {
	pulumi.Input

	ToZoneRedundancyOutput() ZoneRedundancyOutput
	ToZoneRedundancyOutputWithContext(context.Context) ZoneRedundancyOutput
}

var zoneRedundancyPtrType = reflect.TypeOf((**ZoneRedundancy)(nil)).Elem()

type ZoneRedundancyPtrInput interface {
	pulumi.Input

	ToZoneRedundancyPtrOutput() ZoneRedundancyPtrOutput
	ToZoneRedundancyPtrOutputWithContext(context.Context) ZoneRedundancyPtrOutput
}

type zoneRedundancyPtr string

func ZoneRedundancyPtr(v string) ZoneRedundancyPtrInput {
	return (*zoneRedundancyPtr)(&v)
}

func (*zoneRedundancyPtr) ElementType() reflect.Type {
	return zoneRedundancyPtrType
}

func (in *zoneRedundancyPtr) ToZoneRedundancyPtrOutput() ZoneRedundancyPtrOutput {
	return pulumi.ToOutput(in).(ZoneRedundancyPtrOutput)
}

func (in *zoneRedundancyPtr) ToZoneRedundancyPtrOutputWithContext(ctx context.Context) ZoneRedundancyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ZoneRedundancyPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(ActionsRequiredOutput{})
	pulumi.RegisterOutputType(ActionsRequiredPtrOutput{})
	pulumi.RegisterOutputType(AuditLogStatusOutput{})
	pulumi.RegisterOutputType(AuditLogStatusPtrOutput{})
	pulumi.RegisterOutputType(ConnectedRegistryModeOutput{})
	pulumi.RegisterOutputType(ConnectedRegistryModePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStatusOutput{})
	pulumi.RegisterOutputType(ConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(DefaultActionOutput{})
	pulumi.RegisterOutputType(DefaultActionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionStatusOutput{})
	pulumi.RegisterOutputType(EncryptionStatusPtrOutput{})
	pulumi.RegisterOutputType(ExportPolicyStatusOutput{})
	pulumi.RegisterOutputType(ExportPolicyStatusPtrOutput{})
	pulumi.RegisterOutputType(LogLevelOutput{})
	pulumi.RegisterOutputType(LogLevelPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleBypassOptionsOutput{})
	pulumi.RegisterOutputType(NetworkRuleBypassOptionsPtrOutput{})
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
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(TokenCertificateNameOutput{})
	pulumi.RegisterOutputType(TokenCertificateNamePtrOutput{})
	pulumi.RegisterOutputType(TokenPasswordNameOutput{})
	pulumi.RegisterOutputType(TokenPasswordNamePtrOutput{})
	pulumi.RegisterOutputType(TokenStatusOutput{})
	pulumi.RegisterOutputType(TokenStatusPtrOutput{})
	pulumi.RegisterOutputType(TriggerStatusOutput{})
	pulumi.RegisterOutputType(TriggerStatusPtrOutput{})
	pulumi.RegisterOutputType(TrustPolicyTypeOutput{})
	pulumi.RegisterOutputType(TrustPolicyTypePtrOutput{})
	pulumi.RegisterOutputType(WebhookActionOutput{})
	pulumi.RegisterOutputType(WebhookActionPtrOutput{})
	pulumi.RegisterOutputType(WebhookStatusOutput{})
	pulumi.RegisterOutputType(WebhookStatusPtrOutput{})
	pulumi.RegisterOutputType(ZoneRedundancyOutput{})
	pulumi.RegisterOutputType(ZoneRedundancyPtrOutput{})
}
