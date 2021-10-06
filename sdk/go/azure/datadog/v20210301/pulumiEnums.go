


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedIdentityTypes string

const (
	ManagedIdentityTypesSystemAssigned = ManagedIdentityTypes("SystemAssigned")
	ManagedIdentityTypesUserAssigned   = ManagedIdentityTypes("UserAssigned")
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

type MonitoringStatus string

const (
	MonitoringStatusEnabled  = MonitoringStatus("Enabled")
	MonitoringStatusDisabled = MonitoringStatus("Disabled")
)

func (MonitoringStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringStatus)(nil)).Elem()
}

func (e MonitoringStatus) ToMonitoringStatusOutput() MonitoringStatusOutput {
	return pulumi.ToOutput(e).(MonitoringStatusOutput)
}

func (e MonitoringStatus) ToMonitoringStatusOutputWithContext(ctx context.Context) MonitoringStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MonitoringStatusOutput)
}

func (e MonitoringStatus) ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput {
	return e.ToMonitoringStatusPtrOutputWithContext(context.Background())
}

func (e MonitoringStatus) ToMonitoringStatusPtrOutputWithContext(ctx context.Context) MonitoringStatusPtrOutput {
	return MonitoringStatus(e).ToMonitoringStatusOutputWithContext(ctx).ToMonitoringStatusPtrOutputWithContext(ctx)
}

func (e MonitoringStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitoringStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitoringStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonitoringStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MonitoringStatusOutput struct{ *pulumi.OutputState }

func (MonitoringStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringStatus)(nil)).Elem()
}

func (o MonitoringStatusOutput) ToMonitoringStatusOutput() MonitoringStatusOutput {
	return o
}

func (o MonitoringStatusOutput) ToMonitoringStatusOutputWithContext(ctx context.Context) MonitoringStatusOutput {
	return o
}

func (o MonitoringStatusOutput) ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput {
	return o.ToMonitoringStatusPtrOutputWithContext(context.Background())
}

func (o MonitoringStatusOutput) ToMonitoringStatusPtrOutputWithContext(ctx context.Context) MonitoringStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitoringStatus) *MonitoringStatus {
		return &v
	}).(MonitoringStatusPtrOutput)
}

func (o MonitoringStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonitoringStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitoringStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonitoringStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitoringStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitoringStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonitoringStatusPtrOutput struct{ *pulumi.OutputState }

func (MonitoringStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringStatus)(nil)).Elem()
}

func (o MonitoringStatusPtrOutput) ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput {
	return o
}

func (o MonitoringStatusPtrOutput) ToMonitoringStatusPtrOutputWithContext(ctx context.Context) MonitoringStatusPtrOutput {
	return o
}

func (o MonitoringStatusPtrOutput) Elem() MonitoringStatusOutput {
	return o.ApplyT(func(v *MonitoringStatus) MonitoringStatus {
		if v != nil {
			return *v
		}
		var ret MonitoringStatus
		return ret
	}).(MonitoringStatusOutput)
}

func (o MonitoringStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitoringStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonitoringStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MonitoringStatusInput interface {
	pulumi.Input

	ToMonitoringStatusOutput() MonitoringStatusOutput
	ToMonitoringStatusOutputWithContext(context.Context) MonitoringStatusOutput
}

var monitoringStatusPtrType = reflect.TypeOf((**MonitoringStatus)(nil)).Elem()

type MonitoringStatusPtrInput interface {
	pulumi.Input

	ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput
	ToMonitoringStatusPtrOutputWithContext(context.Context) MonitoringStatusPtrOutput
}

type monitoringStatusPtr string

func MonitoringStatusPtr(v string) MonitoringStatusPtrInput {
	return (*monitoringStatusPtr)(&v)
}

func (*monitoringStatusPtr) ElementType() reflect.Type {
	return monitoringStatusPtrType
}

func (in *monitoringStatusPtr) ToMonitoringStatusPtrOutput() MonitoringStatusPtrOutput {
	return pulumi.ToOutput(in).(MonitoringStatusPtrOutput)
}

func (in *monitoringStatusPtr) ToMonitoringStatusPtrOutputWithContext(ctx context.Context) MonitoringStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MonitoringStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedIdentityTypesOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypesPtrOutput{})
	pulumi.RegisterOutputType(MonitoringStatusOutput{})
	pulumi.RegisterOutputType(MonitoringStatusPtrOutput{})
}
