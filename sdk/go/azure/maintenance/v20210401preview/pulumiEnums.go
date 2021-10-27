


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MaintenanceScope string

const (
	// This maintenance scope controls installation of azure platform updates i.e. services on physical nodes hosting customer VMs.
	MaintenanceScopeHost = MaintenanceScope("Host")
	// This maintenance scope controls os image installation on VM/VMSS
	MaintenanceScopeOSImage = MaintenanceScope("OSImage")
	// This maintenance scope controls extension installation on VM/VMSS
	MaintenanceScopeExtension = MaintenanceScope("Extension")
	// This maintenance scope controls installation of windows and linux packages on VM/VMSS
	MaintenanceScopeInGuestPatch = MaintenanceScope("InGuestPatch")
	// This maintenance scope controls installation of SQL server platform updates.
	MaintenanceScopeSQLDB = MaintenanceScope("SQLDB")
	// This maintenance scope controls installation of SQL managed instance platform update.
	MaintenanceScopeSQLManagedInstance = MaintenanceScope("SQLManagedInstance")
)

func (MaintenanceScope) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceScope)(nil)).Elem()
}

func (e MaintenanceScope) ToMaintenanceScopeOutput() MaintenanceScopeOutput {
	return pulumi.ToOutput(e).(MaintenanceScopeOutput)
}

func (e MaintenanceScope) ToMaintenanceScopeOutputWithContext(ctx context.Context) MaintenanceScopeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MaintenanceScopeOutput)
}

func (e MaintenanceScope) ToMaintenanceScopePtrOutput() MaintenanceScopePtrOutput {
	return e.ToMaintenanceScopePtrOutputWithContext(context.Background())
}

func (e MaintenanceScope) ToMaintenanceScopePtrOutputWithContext(ctx context.Context) MaintenanceScopePtrOutput {
	return MaintenanceScope(e).ToMaintenanceScopeOutputWithContext(ctx).ToMaintenanceScopePtrOutputWithContext(ctx)
}

func (e MaintenanceScope) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MaintenanceScope) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MaintenanceScope) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MaintenanceScope) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MaintenanceScopeOutput struct{ *pulumi.OutputState }

func (MaintenanceScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceScope)(nil)).Elem()
}

func (o MaintenanceScopeOutput) ToMaintenanceScopeOutput() MaintenanceScopeOutput {
	return o
}

func (o MaintenanceScopeOutput) ToMaintenanceScopeOutputWithContext(ctx context.Context) MaintenanceScopeOutput {
	return o
}

func (o MaintenanceScopeOutput) ToMaintenanceScopePtrOutput() MaintenanceScopePtrOutput {
	return o.ToMaintenanceScopePtrOutputWithContext(context.Background())
}

func (o MaintenanceScopeOutput) ToMaintenanceScopePtrOutputWithContext(ctx context.Context) MaintenanceScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MaintenanceScope) *MaintenanceScope {
		return &v
	}).(MaintenanceScopePtrOutput)
}

func (o MaintenanceScopeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MaintenanceScopeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MaintenanceScope) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MaintenanceScopeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MaintenanceScopeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MaintenanceScope) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MaintenanceScopePtrOutput struct{ *pulumi.OutputState }

func (MaintenanceScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceScope)(nil)).Elem()
}

func (o MaintenanceScopePtrOutput) ToMaintenanceScopePtrOutput() MaintenanceScopePtrOutput {
	return o
}

func (o MaintenanceScopePtrOutput) ToMaintenanceScopePtrOutputWithContext(ctx context.Context) MaintenanceScopePtrOutput {
	return o
}

func (o MaintenanceScopePtrOutput) Elem() MaintenanceScopeOutput {
	return o.ApplyT(func(v *MaintenanceScope) MaintenanceScope {
		if v != nil {
			return *v
		}
		var ret MaintenanceScope
		return ret
	}).(MaintenanceScopeOutput)
}

func (o MaintenanceScopePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MaintenanceScopePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MaintenanceScope) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MaintenanceScopeInput interface {
	pulumi.Input

	ToMaintenanceScopeOutput() MaintenanceScopeOutput
	ToMaintenanceScopeOutputWithContext(context.Context) MaintenanceScopeOutput
}

var maintenanceScopePtrType = reflect.TypeOf((**MaintenanceScope)(nil)).Elem()

type MaintenanceScopePtrInput interface {
	pulumi.Input

	ToMaintenanceScopePtrOutput() MaintenanceScopePtrOutput
	ToMaintenanceScopePtrOutputWithContext(context.Context) MaintenanceScopePtrOutput
}

type maintenanceScopePtr string

func MaintenanceScopePtr(v string) MaintenanceScopePtrInput {
	return (*maintenanceScopePtr)(&v)
}

func (*maintenanceScopePtr) ElementType() reflect.Type {
	return maintenanceScopePtrType
}

func (in *maintenanceScopePtr) ToMaintenanceScopePtrOutput() MaintenanceScopePtrOutput {
	return pulumi.ToOutput(in).(MaintenanceScopePtrOutput)
}

func (in *maintenanceScopePtr) ToMaintenanceScopePtrOutputWithContext(ctx context.Context) MaintenanceScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MaintenanceScopePtrOutput)
}

type RebootOptions string

const (
	RebootOptionsNeverReboot      = RebootOptions("NeverReboot")
	RebootOptionsRebootIfRequired = RebootOptions("RebootIfRequired")
	RebootOptionsAlwaysReboot     = RebootOptions("AlwaysReboot")
)

func (RebootOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*RebootOptions)(nil)).Elem()
}

func (e RebootOptions) ToRebootOptionsOutput() RebootOptionsOutput {
	return pulumi.ToOutput(e).(RebootOptionsOutput)
}

func (e RebootOptions) ToRebootOptionsOutputWithContext(ctx context.Context) RebootOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RebootOptionsOutput)
}

func (e RebootOptions) ToRebootOptionsPtrOutput() RebootOptionsPtrOutput {
	return e.ToRebootOptionsPtrOutputWithContext(context.Background())
}

func (e RebootOptions) ToRebootOptionsPtrOutputWithContext(ctx context.Context) RebootOptionsPtrOutput {
	return RebootOptions(e).ToRebootOptionsOutputWithContext(ctx).ToRebootOptionsPtrOutputWithContext(ctx)
}

func (e RebootOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RebootOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RebootOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RebootOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RebootOptionsOutput struct{ *pulumi.OutputState }

func (RebootOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RebootOptions)(nil)).Elem()
}

func (o RebootOptionsOutput) ToRebootOptionsOutput() RebootOptionsOutput {
	return o
}

func (o RebootOptionsOutput) ToRebootOptionsOutputWithContext(ctx context.Context) RebootOptionsOutput {
	return o
}

func (o RebootOptionsOutput) ToRebootOptionsPtrOutput() RebootOptionsPtrOutput {
	return o.ToRebootOptionsPtrOutputWithContext(context.Background())
}

func (o RebootOptionsOutput) ToRebootOptionsPtrOutputWithContext(ctx context.Context) RebootOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RebootOptions) *RebootOptions {
		return &v
	}).(RebootOptionsPtrOutput)
}

func (o RebootOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RebootOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RebootOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RebootOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RebootOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RebootOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RebootOptionsPtrOutput struct{ *pulumi.OutputState }

func (RebootOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RebootOptions)(nil)).Elem()
}

func (o RebootOptionsPtrOutput) ToRebootOptionsPtrOutput() RebootOptionsPtrOutput {
	return o
}

func (o RebootOptionsPtrOutput) ToRebootOptionsPtrOutputWithContext(ctx context.Context) RebootOptionsPtrOutput {
	return o
}

func (o RebootOptionsPtrOutput) Elem() RebootOptionsOutput {
	return o.ApplyT(func(v *RebootOptions) RebootOptions {
		if v != nil {
			return *v
		}
		var ret RebootOptions
		return ret
	}).(RebootOptionsOutput)
}

func (o RebootOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RebootOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RebootOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RebootOptionsInput interface {
	pulumi.Input

	ToRebootOptionsOutput() RebootOptionsOutput
	ToRebootOptionsOutputWithContext(context.Context) RebootOptionsOutput
}

var rebootOptionsPtrType = reflect.TypeOf((**RebootOptions)(nil)).Elem()

type RebootOptionsPtrInput interface {
	pulumi.Input

	ToRebootOptionsPtrOutput() RebootOptionsPtrOutput
	ToRebootOptionsPtrOutputWithContext(context.Context) RebootOptionsPtrOutput
}

type rebootOptionsPtr string

func RebootOptionsPtr(v string) RebootOptionsPtrInput {
	return (*rebootOptionsPtr)(&v)
}

func (*rebootOptionsPtr) ElementType() reflect.Type {
	return rebootOptionsPtrType
}

func (in *rebootOptionsPtr) ToRebootOptionsPtrOutput() RebootOptionsPtrOutput {
	return pulumi.ToOutput(in).(RebootOptionsPtrOutput)
}

func (in *rebootOptionsPtr) ToRebootOptionsPtrOutputWithContext(ctx context.Context) RebootOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RebootOptionsPtrOutput)
}

type TaskScope string

const (
	TaskScopeGlobal   = TaskScope("Global")
	TaskScopeResource = TaskScope("Resource")
)

func (TaskScope) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskScope)(nil)).Elem()
}

func (e TaskScope) ToTaskScopeOutput() TaskScopeOutput {
	return pulumi.ToOutput(e).(TaskScopeOutput)
}

func (e TaskScope) ToTaskScopeOutputWithContext(ctx context.Context) TaskScopeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TaskScopeOutput)
}

func (e TaskScope) ToTaskScopePtrOutput() TaskScopePtrOutput {
	return e.ToTaskScopePtrOutputWithContext(context.Background())
}

func (e TaskScope) ToTaskScopePtrOutputWithContext(ctx context.Context) TaskScopePtrOutput {
	return TaskScope(e).ToTaskScopeOutputWithContext(ctx).ToTaskScopePtrOutputWithContext(ctx)
}

func (e TaskScope) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TaskScope) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TaskScope) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TaskScope) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TaskScopeOutput struct{ *pulumi.OutputState }

func (TaskScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskScope)(nil)).Elem()
}

func (o TaskScopeOutput) ToTaskScopeOutput() TaskScopeOutput {
	return o
}

func (o TaskScopeOutput) ToTaskScopeOutputWithContext(ctx context.Context) TaskScopeOutput {
	return o
}

func (o TaskScopeOutput) ToTaskScopePtrOutput() TaskScopePtrOutput {
	return o.ToTaskScopePtrOutputWithContext(context.Background())
}

func (o TaskScopeOutput) ToTaskScopePtrOutputWithContext(ctx context.Context) TaskScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TaskScope) *TaskScope {
		return &v
	}).(TaskScopePtrOutput)
}

func (o TaskScopeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TaskScopeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TaskScope) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TaskScopeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TaskScopeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TaskScope) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TaskScopePtrOutput struct{ *pulumi.OutputState }

func (TaskScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskScope)(nil)).Elem()
}

func (o TaskScopePtrOutput) ToTaskScopePtrOutput() TaskScopePtrOutput {
	return o
}

func (o TaskScopePtrOutput) ToTaskScopePtrOutputWithContext(ctx context.Context) TaskScopePtrOutput {
	return o
}

func (o TaskScopePtrOutput) Elem() TaskScopeOutput {
	return o.ApplyT(func(v *TaskScope) TaskScope {
		if v != nil {
			return *v
		}
		var ret TaskScope
		return ret
	}).(TaskScopeOutput)
}

func (o TaskScopePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TaskScopePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TaskScope) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TaskScopeInput interface {
	pulumi.Input

	ToTaskScopeOutput() TaskScopeOutput
	ToTaskScopeOutputWithContext(context.Context) TaskScopeOutput
}

var taskScopePtrType = reflect.TypeOf((**TaskScope)(nil)).Elem()

type TaskScopePtrInput interface {
	pulumi.Input

	ToTaskScopePtrOutput() TaskScopePtrOutput
	ToTaskScopePtrOutputWithContext(context.Context) TaskScopePtrOutput
}

type taskScopePtr string

func TaskScopePtr(v string) TaskScopePtrInput {
	return (*taskScopePtr)(&v)
}

func (*taskScopePtr) ElementType() reflect.Type {
	return taskScopePtrType
}

func (in *taskScopePtr) ToTaskScopePtrOutput() TaskScopePtrOutput {
	return pulumi.ToOutput(in).(TaskScopePtrOutput)
}

func (in *taskScopePtr) ToTaskScopePtrOutputWithContext(ctx context.Context) TaskScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TaskScopePtrOutput)
}

type Visibility string

const (
	// Only visible to users with permissions.
	VisibilityCustom = Visibility("Custom")
	// Visible to all users.
	VisibilityPublic = Visibility("Public")
)

func (Visibility) ElementType() reflect.Type {
	return reflect.TypeOf((*Visibility)(nil)).Elem()
}

func (e Visibility) ToVisibilityOutput() VisibilityOutput {
	return pulumi.ToOutput(e).(VisibilityOutput)
}

func (e Visibility) ToVisibilityOutputWithContext(ctx context.Context) VisibilityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VisibilityOutput)
}

func (e Visibility) ToVisibilityPtrOutput() VisibilityPtrOutput {
	return e.ToVisibilityPtrOutputWithContext(context.Background())
}

func (e Visibility) ToVisibilityPtrOutputWithContext(ctx context.Context) VisibilityPtrOutput {
	return Visibility(e).ToVisibilityOutputWithContext(ctx).ToVisibilityPtrOutputWithContext(ctx)
}

func (e Visibility) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Visibility) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Visibility) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Visibility) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VisibilityOutput struct{ *pulumi.OutputState }

func (VisibilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Visibility)(nil)).Elem()
}

func (o VisibilityOutput) ToVisibilityOutput() VisibilityOutput {
	return o
}

func (o VisibilityOutput) ToVisibilityOutputWithContext(ctx context.Context) VisibilityOutput {
	return o
}

func (o VisibilityOutput) ToVisibilityPtrOutput() VisibilityPtrOutput {
	return o.ToVisibilityPtrOutputWithContext(context.Background())
}

func (o VisibilityOutput) ToVisibilityPtrOutputWithContext(ctx context.Context) VisibilityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Visibility) *Visibility {
		return &v
	}).(VisibilityPtrOutput)
}

func (o VisibilityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VisibilityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Visibility) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VisibilityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VisibilityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Visibility) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VisibilityPtrOutput struct{ *pulumi.OutputState }

func (VisibilityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Visibility)(nil)).Elem()
}

func (o VisibilityPtrOutput) ToVisibilityPtrOutput() VisibilityPtrOutput {
	return o
}

func (o VisibilityPtrOutput) ToVisibilityPtrOutputWithContext(ctx context.Context) VisibilityPtrOutput {
	return o
}

func (o VisibilityPtrOutput) Elem() VisibilityOutput {
	return o.ApplyT(func(v *Visibility) Visibility {
		if v != nil {
			return *v
		}
		var ret Visibility
		return ret
	}).(VisibilityOutput)
}

func (o VisibilityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VisibilityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Visibility) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VisibilityInput interface {
	pulumi.Input

	ToVisibilityOutput() VisibilityOutput
	ToVisibilityOutputWithContext(context.Context) VisibilityOutput
}

var visibilityPtrType = reflect.TypeOf((**Visibility)(nil)).Elem()

type VisibilityPtrInput interface {
	pulumi.Input

	ToVisibilityPtrOutput() VisibilityPtrOutput
	ToVisibilityPtrOutputWithContext(context.Context) VisibilityPtrOutput
}

type visibilityPtr string

func VisibilityPtr(v string) VisibilityPtrInput {
	return (*visibilityPtr)(&v)
}

func (*visibilityPtr) ElementType() reflect.Type {
	return visibilityPtrType
}

func (in *visibilityPtr) ToVisibilityPtrOutput() VisibilityPtrOutput {
	return pulumi.ToOutput(in).(VisibilityPtrOutput)
}

func (in *visibilityPtr) ToVisibilityPtrOutputWithContext(ctx context.Context) VisibilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VisibilityPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceScopeInput)(nil)).Elem(), MaintenanceScope("Host"))
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceScopePtrInput)(nil)).Elem(), MaintenanceScope("Host"))
	pulumi.RegisterInputType(reflect.TypeOf((*RebootOptionsInput)(nil)).Elem(), RebootOptions("NeverReboot"))
	pulumi.RegisterInputType(reflect.TypeOf((*RebootOptionsPtrInput)(nil)).Elem(), RebootOptions("NeverReboot"))
	pulumi.RegisterInputType(reflect.TypeOf((*TaskScopeInput)(nil)).Elem(), TaskScope("Global"))
	pulumi.RegisterInputType(reflect.TypeOf((*TaskScopePtrInput)(nil)).Elem(), TaskScope("Global"))
	pulumi.RegisterInputType(reflect.TypeOf((*VisibilityInput)(nil)).Elem(), Visibility("Custom"))
	pulumi.RegisterInputType(reflect.TypeOf((*VisibilityPtrInput)(nil)).Elem(), Visibility("Custom"))
	pulumi.RegisterOutputType(MaintenanceScopeOutput{})
	pulumi.RegisterOutputType(MaintenanceScopePtrOutput{})
	pulumi.RegisterOutputType(RebootOptionsOutput{})
	pulumi.RegisterOutputType(RebootOptionsPtrOutput{})
	pulumi.RegisterOutputType(TaskScopeOutput{})
	pulumi.RegisterOutputType(TaskScopePtrOutput{})
	pulumi.RegisterOutputType(VisibilityOutput{})
	pulumi.RegisterOutputType(VisibilityPtrOutput{})
}
