


package changeanalysis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ChangeDetailsMode string

const (
	ChangeDetailsModeNone    = ChangeDetailsMode("None")
	ChangeDetailsModeInclude = ChangeDetailsMode("Include")
	ChangeDetailsModeExclude = ChangeDetailsMode("Exclude")
)

func (ChangeDetailsMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ChangeDetailsMode)(nil)).Elem()
}

func (e ChangeDetailsMode) ToChangeDetailsModeOutput() ChangeDetailsModeOutput {
	return pulumi.ToOutput(e).(ChangeDetailsModeOutput)
}

func (e ChangeDetailsMode) ToChangeDetailsModeOutputWithContext(ctx context.Context) ChangeDetailsModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ChangeDetailsModeOutput)
}

func (e ChangeDetailsMode) ToChangeDetailsModePtrOutput() ChangeDetailsModePtrOutput {
	return e.ToChangeDetailsModePtrOutputWithContext(context.Background())
}

func (e ChangeDetailsMode) ToChangeDetailsModePtrOutputWithContext(ctx context.Context) ChangeDetailsModePtrOutput {
	return ChangeDetailsMode(e).ToChangeDetailsModeOutputWithContext(ctx).ToChangeDetailsModePtrOutputWithContext(ctx)
}

func (e ChangeDetailsMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChangeDetailsMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChangeDetailsMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ChangeDetailsMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ChangeDetailsModeOutput struct{ *pulumi.OutputState }

func (ChangeDetailsModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChangeDetailsMode)(nil)).Elem()
}

func (o ChangeDetailsModeOutput) ToChangeDetailsModeOutput() ChangeDetailsModeOutput {
	return o
}

func (o ChangeDetailsModeOutput) ToChangeDetailsModeOutputWithContext(ctx context.Context) ChangeDetailsModeOutput {
	return o
}

func (o ChangeDetailsModeOutput) ToChangeDetailsModePtrOutput() ChangeDetailsModePtrOutput {
	return o.ToChangeDetailsModePtrOutputWithContext(context.Background())
}

func (o ChangeDetailsModeOutput) ToChangeDetailsModePtrOutputWithContext(ctx context.Context) ChangeDetailsModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChangeDetailsMode) *ChangeDetailsMode {
		return &v
	}).(ChangeDetailsModePtrOutput)
}

func (o ChangeDetailsModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ChangeDetailsModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChangeDetailsMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ChangeDetailsModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChangeDetailsModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChangeDetailsMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ChangeDetailsModePtrOutput struct{ *pulumi.OutputState }

func (ChangeDetailsModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChangeDetailsMode)(nil)).Elem()
}

func (o ChangeDetailsModePtrOutput) ToChangeDetailsModePtrOutput() ChangeDetailsModePtrOutput {
	return o
}

func (o ChangeDetailsModePtrOutput) ToChangeDetailsModePtrOutputWithContext(ctx context.Context) ChangeDetailsModePtrOutput {
	return o
}

func (o ChangeDetailsModePtrOutput) Elem() ChangeDetailsModeOutput {
	return o.ApplyT(func(v *ChangeDetailsMode) ChangeDetailsMode {
		if v != nil {
			return *v
		}
		var ret ChangeDetailsMode
		return ret
	}).(ChangeDetailsModeOutput)
}

func (o ChangeDetailsModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChangeDetailsModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ChangeDetailsMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ChangeDetailsModeInput interface {
	pulumi.Input

	ToChangeDetailsModeOutput() ChangeDetailsModeOutput
	ToChangeDetailsModeOutputWithContext(context.Context) ChangeDetailsModeOutput
}

var changeDetailsModePtrType = reflect.TypeOf((**ChangeDetailsMode)(nil)).Elem()

type ChangeDetailsModePtrInput interface {
	pulumi.Input

	ToChangeDetailsModePtrOutput() ChangeDetailsModePtrOutput
	ToChangeDetailsModePtrOutputWithContext(context.Context) ChangeDetailsModePtrOutput
}

type changeDetailsModePtr string

func ChangeDetailsModePtr(v string) ChangeDetailsModePtrInput {
	return (*changeDetailsModePtr)(&v)
}

func (*changeDetailsModePtr) ElementType() reflect.Type {
	return changeDetailsModePtrType
}

func (in *changeDetailsModePtr) ToChangeDetailsModePtrOutput() ChangeDetailsModePtrOutput {
	return pulumi.ToOutput(in).(ChangeDetailsModePtrOutput)
}

func (in *changeDetailsModePtr) ToChangeDetailsModePtrOutputWithContext(ctx context.Context) ChangeDetailsModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ChangeDetailsModePtrOutput)
}

type ManagedIdentityTypes string

const (
	ManagedIdentityTypesNone           = ManagedIdentityTypes("None")
	ManagedIdentityTypesSystemAssigned = ManagedIdentityTypes("SystemAssigned")
)

func (ManagedIdentityTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityTypes)(nil)).Elem()
}

func (e ManagedIdentityTypes) ToManagedIdentityTypesOutput() ManagedIdentityTypesOutput {
	return pulumi.ToOutput(e).(ManagedIdentityTypesOutput)
}

func (e ManagedIdentityTypes) ToManagedIdentityTypesOutputWithContext(ctx context.Context) ManagedIdentityTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedIdentityTypesOutput)
}

func (e ManagedIdentityTypes) ToManagedIdentityTypesPtrOutput() ManagedIdentityTypesPtrOutput {
	return e.ToManagedIdentityTypesPtrOutputWithContext(context.Background())
}

func (e ManagedIdentityTypes) ToManagedIdentityTypesPtrOutputWithContext(ctx context.Context) ManagedIdentityTypesPtrOutput {
	return ManagedIdentityTypes(e).ToManagedIdentityTypesOutputWithContext(ctx).ToManagedIdentityTypesPtrOutputWithContext(ctx)
}

func (e ManagedIdentityTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedIdentityTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedIdentityTypesOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityTypes)(nil)).Elem()
}

func (o ManagedIdentityTypesOutput) ToManagedIdentityTypesOutput() ManagedIdentityTypesOutput {
	return o
}

func (o ManagedIdentityTypesOutput) ToManagedIdentityTypesOutputWithContext(ctx context.Context) ManagedIdentityTypesOutput {
	return o
}

func (o ManagedIdentityTypesOutput) ToManagedIdentityTypesPtrOutput() ManagedIdentityTypesPtrOutput {
	return o.ToManagedIdentityTypesPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypesOutput) ToManagedIdentityTypesPtrOutputWithContext(ctx context.Context) ManagedIdentityTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityTypes) *ManagedIdentityTypes {
		return &v
	}).(ManagedIdentityTypesPtrOutput)
}

func (o ManagedIdentityTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedIdentityTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedIdentityTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityTypesPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityTypes)(nil)).Elem()
}

func (o ManagedIdentityTypesPtrOutput) ToManagedIdentityTypesPtrOutput() ManagedIdentityTypesPtrOutput {
	return o
}

func (o ManagedIdentityTypesPtrOutput) ToManagedIdentityTypesPtrOutputWithContext(ctx context.Context) ManagedIdentityTypesPtrOutput {
	return o
}

func (o ManagedIdentityTypesPtrOutput) Elem() ManagedIdentityTypesOutput {
	return o.ApplyT(func(v *ManagedIdentityTypes) ManagedIdentityTypes {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityTypes
		return ret
	}).(ManagedIdentityTypesOutput)
}

func (o ManagedIdentityTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedIdentityTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedIdentityTypesInput interface {
	pulumi.Input

	ToManagedIdentityTypesOutput() ManagedIdentityTypesOutput
	ToManagedIdentityTypesOutputWithContext(context.Context) ManagedIdentityTypesOutput
}

var managedIdentityTypesPtrType = reflect.TypeOf((**ManagedIdentityTypes)(nil)).Elem()

type ManagedIdentityTypesPtrInput interface {
	pulumi.Input

	ToManagedIdentityTypesPtrOutput() ManagedIdentityTypesPtrOutput
	ToManagedIdentityTypesPtrOutputWithContext(context.Context) ManagedIdentityTypesPtrOutput
}

type managedIdentityTypesPtr string

func ManagedIdentityTypesPtr(v string) ManagedIdentityTypesPtrInput {
	return (*managedIdentityTypesPtr)(&v)
}

func (*managedIdentityTypesPtr) ElementType() reflect.Type {
	return managedIdentityTypesPtrType
}

func (in *managedIdentityTypesPtr) ToManagedIdentityTypesPtrOutput() ManagedIdentityTypesPtrOutput {
	return pulumi.ToOutput(in).(ManagedIdentityTypesPtrOutput)
}

func (in *managedIdentityTypesPtr) ToManagedIdentityTypesPtrOutputWithContext(ctx context.Context) ManagedIdentityTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedIdentityTypesPtrOutput)
}

type NotificationsState string

const (
	NotificationsStateNone     = NotificationsState("None")
	NotificationsStateEnabled  = NotificationsState("Enabled")
	NotificationsStateDisabled = NotificationsState("Disabled")
)

func (NotificationsState) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationsState)(nil)).Elem()
}

func (e NotificationsState) ToNotificationsStateOutput() NotificationsStateOutput {
	return pulumi.ToOutput(e).(NotificationsStateOutput)
}

func (e NotificationsState) ToNotificationsStateOutputWithContext(ctx context.Context) NotificationsStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NotificationsStateOutput)
}

func (e NotificationsState) ToNotificationsStatePtrOutput() NotificationsStatePtrOutput {
	return e.ToNotificationsStatePtrOutputWithContext(context.Background())
}

func (e NotificationsState) ToNotificationsStatePtrOutputWithContext(ctx context.Context) NotificationsStatePtrOutput {
	return NotificationsState(e).ToNotificationsStateOutputWithContext(ctx).ToNotificationsStatePtrOutputWithContext(ctx)
}

func (e NotificationsState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationsState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationsState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NotificationsState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NotificationsStateOutput struct{ *pulumi.OutputState }

func (NotificationsStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationsState)(nil)).Elem()
}

func (o NotificationsStateOutput) ToNotificationsStateOutput() NotificationsStateOutput {
	return o
}

func (o NotificationsStateOutput) ToNotificationsStateOutputWithContext(ctx context.Context) NotificationsStateOutput {
	return o
}

func (o NotificationsStateOutput) ToNotificationsStatePtrOutput() NotificationsStatePtrOutput {
	return o.ToNotificationsStatePtrOutputWithContext(context.Background())
}

func (o NotificationsStateOutput) ToNotificationsStatePtrOutputWithContext(ctx context.Context) NotificationsStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationsState) *NotificationsState {
		return &v
	}).(NotificationsStatePtrOutput)
}

func (o NotificationsStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NotificationsStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotificationsState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NotificationsStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotificationsStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotificationsState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NotificationsStatePtrOutput struct{ *pulumi.OutputState }

func (NotificationsStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationsState)(nil)).Elem()
}

func (o NotificationsStatePtrOutput) ToNotificationsStatePtrOutput() NotificationsStatePtrOutput {
	return o
}

func (o NotificationsStatePtrOutput) ToNotificationsStatePtrOutputWithContext(ctx context.Context) NotificationsStatePtrOutput {
	return o
}

func (o NotificationsStatePtrOutput) Elem() NotificationsStateOutput {
	return o.ApplyT(func(v *NotificationsState) NotificationsState {
		if v != nil {
			return *v
		}
		var ret NotificationsState
		return ret
	}).(NotificationsStateOutput)
}

func (o NotificationsStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotificationsStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NotificationsState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NotificationsStateInput interface {
	pulumi.Input

	ToNotificationsStateOutput() NotificationsStateOutput
	ToNotificationsStateOutputWithContext(context.Context) NotificationsStateOutput
}

var notificationsStatePtrType = reflect.TypeOf((**NotificationsState)(nil)).Elem()

type NotificationsStatePtrInput interface {
	pulumi.Input

	ToNotificationsStatePtrOutput() NotificationsStatePtrOutput
	ToNotificationsStatePtrOutputWithContext(context.Context) NotificationsStatePtrOutput
}

type notificationsStatePtr string

func NotificationsStatePtr(v string) NotificationsStatePtrInput {
	return (*notificationsStatePtr)(&v)
}

func (*notificationsStatePtr) ElementType() reflect.Type {
	return notificationsStatePtrType
}

func (in *notificationsStatePtr) ToNotificationsStatePtrOutput() NotificationsStatePtrOutput {
	return pulumi.ToOutput(in).(NotificationsStatePtrOutput)
}

func (in *notificationsStatePtr) ToNotificationsStatePtrOutputWithContext(ctx context.Context) NotificationsStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NotificationsStatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ChangeDetailsModeOutput{})
	pulumi.RegisterOutputType(ChangeDetailsModePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypesOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypesPtrOutput{})
	pulumi.RegisterOutputType(NotificationsStateOutput{})
	pulumi.RegisterOutputType(NotificationsStatePtrOutput{})
}
