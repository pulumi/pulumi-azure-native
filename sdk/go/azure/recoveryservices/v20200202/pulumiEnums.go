


package v20200202

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateEnabled  = InfrastructureEncryptionState("Enabled")
	InfrastructureEncryptionStateDisabled = InfrastructureEncryptionState("Disabled")
)

func (InfrastructureEncryptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureEncryptionState)(nil)).Elem()
}

func (e InfrastructureEncryptionState) ToInfrastructureEncryptionStateOutput() InfrastructureEncryptionStateOutput {
	return pulumi.ToOutput(e).(InfrastructureEncryptionStateOutput)
}

func (e InfrastructureEncryptionState) ToInfrastructureEncryptionStateOutputWithContext(ctx context.Context) InfrastructureEncryptionStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InfrastructureEncryptionStateOutput)
}

func (e InfrastructureEncryptionState) ToInfrastructureEncryptionStatePtrOutput() InfrastructureEncryptionStatePtrOutput {
	return e.ToInfrastructureEncryptionStatePtrOutputWithContext(context.Background())
}

func (e InfrastructureEncryptionState) ToInfrastructureEncryptionStatePtrOutputWithContext(ctx context.Context) InfrastructureEncryptionStatePtrOutput {
	return InfrastructureEncryptionState(e).ToInfrastructureEncryptionStateOutputWithContext(ctx).ToInfrastructureEncryptionStatePtrOutputWithContext(ctx)
}

func (e InfrastructureEncryptionState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InfrastructureEncryptionState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InfrastructureEncryptionState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InfrastructureEncryptionState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InfrastructureEncryptionStateOutput struct{ *pulumi.OutputState }

func (InfrastructureEncryptionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureEncryptionState)(nil)).Elem()
}

func (o InfrastructureEncryptionStateOutput) ToInfrastructureEncryptionStateOutput() InfrastructureEncryptionStateOutput {
	return o
}

func (o InfrastructureEncryptionStateOutput) ToInfrastructureEncryptionStateOutputWithContext(ctx context.Context) InfrastructureEncryptionStateOutput {
	return o
}

func (o InfrastructureEncryptionStateOutput) ToInfrastructureEncryptionStatePtrOutput() InfrastructureEncryptionStatePtrOutput {
	return o.ToInfrastructureEncryptionStatePtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionStateOutput) ToInfrastructureEncryptionStatePtrOutputWithContext(ctx context.Context) InfrastructureEncryptionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InfrastructureEncryptionState) *InfrastructureEncryptionState {
		return &v
	}).(InfrastructureEncryptionStatePtrOutput)
}

func (o InfrastructureEncryptionStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InfrastructureEncryptionState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InfrastructureEncryptionStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InfrastructureEncryptionState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InfrastructureEncryptionStatePtrOutput struct{ *pulumi.OutputState }

func (InfrastructureEncryptionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InfrastructureEncryptionState)(nil)).Elem()
}

func (o InfrastructureEncryptionStatePtrOutput) ToInfrastructureEncryptionStatePtrOutput() InfrastructureEncryptionStatePtrOutput {
	return o
}

func (o InfrastructureEncryptionStatePtrOutput) ToInfrastructureEncryptionStatePtrOutputWithContext(ctx context.Context) InfrastructureEncryptionStatePtrOutput {
	return o
}

func (o InfrastructureEncryptionStatePtrOutput) Elem() InfrastructureEncryptionStateOutput {
	return o.ApplyT(func(v *InfrastructureEncryptionState) InfrastructureEncryptionState {
		if v != nil {
			return *v
		}
		var ret InfrastructureEncryptionState
		return ret
	}).(InfrastructureEncryptionStateOutput)
}

func (o InfrastructureEncryptionStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InfrastructureEncryptionState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InfrastructureEncryptionStateInput interface {
	pulumi.Input

	ToInfrastructureEncryptionStateOutput() InfrastructureEncryptionStateOutput
	ToInfrastructureEncryptionStateOutputWithContext(context.Context) InfrastructureEncryptionStateOutput
}

var infrastructureEncryptionStatePtrType = reflect.TypeOf((**InfrastructureEncryptionState)(nil)).Elem()

type InfrastructureEncryptionStatePtrInput interface {
	pulumi.Input

	ToInfrastructureEncryptionStatePtrOutput() InfrastructureEncryptionStatePtrOutput
	ToInfrastructureEncryptionStatePtrOutputWithContext(context.Context) InfrastructureEncryptionStatePtrOutput
}

type infrastructureEncryptionStatePtr string

func InfrastructureEncryptionStatePtr(v string) InfrastructureEncryptionStatePtrInput {
	return (*infrastructureEncryptionStatePtr)(&v)
}

func (*infrastructureEncryptionStatePtr) ElementType() reflect.Type {
	return infrastructureEncryptionStatePtrType
}

func (in *infrastructureEncryptionStatePtr) ToInfrastructureEncryptionStatePtrOutput() InfrastructureEncryptionStatePtrOutput {
	return pulumi.ToOutput(in).(InfrastructureEncryptionStatePtrOutput)
}

func (in *infrastructureEncryptionStatePtr) ToInfrastructureEncryptionStatePtrOutputWithContext(ctx context.Context) InfrastructureEncryptionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InfrastructureEncryptionStatePtrOutput)
}

type PrivateEndpointConnectionStatus string

const (
	PrivateEndpointConnectionStatusPending      = PrivateEndpointConnectionStatus("Pending")
	PrivateEndpointConnectionStatusApproved     = PrivateEndpointConnectionStatus("Approved")
	PrivateEndpointConnectionStatusRejected     = PrivateEndpointConnectionStatus("Rejected")
	PrivateEndpointConnectionStatusDisconnected = PrivateEndpointConnectionStatus("Disconnected")
)

func (PrivateEndpointConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionStatus)(nil)).Elem()
}

func (e PrivateEndpointConnectionStatus) ToPrivateEndpointConnectionStatusOutput() PrivateEndpointConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateEndpointConnectionStatusOutput)
}

func (e PrivateEndpointConnectionStatus) ToPrivateEndpointConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointConnectionStatusOutput)
}

func (e PrivateEndpointConnectionStatus) ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput {
	return e.ToPrivateEndpointConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointConnectionStatus) ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusPtrOutput {
	return PrivateEndpointConnectionStatus(e).ToPrivateEndpointConnectionStatusOutputWithContext(ctx).ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateEndpointConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointConnectionStatusOutput) ToPrivateEndpointConnectionStatusOutput() PrivateEndpointConnectionStatusOutput {
	return o
}

func (o PrivateEndpointConnectionStatusOutput) ToPrivateEndpointConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusOutput {
	return o
}

func (o PrivateEndpointConnectionStatusOutput) ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput {
	return o.ToPrivateEndpointConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionStatusOutput) ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionStatus) *PrivateEndpointConnectionStatus {
		return &v
	}).(PrivateEndpointConnectionStatusPtrOutput)
}

func (o PrivateEndpointConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointConnectionStatusPtrOutput) ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointConnectionStatusPtrOutput) ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointConnectionStatusPtrOutput) Elem() PrivateEndpointConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionStatus) PrivateEndpointConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionStatus
		return ret
	}).(PrivateEndpointConnectionStatusOutput)
}

func (o PrivateEndpointConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateEndpointConnectionStatusInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionStatusOutput() PrivateEndpointConnectionStatusOutput
	ToPrivateEndpointConnectionStatusOutputWithContext(context.Context) PrivateEndpointConnectionStatusOutput
}

var privateEndpointConnectionStatusPtrType = reflect.TypeOf((**PrivateEndpointConnectionStatus)(nil)).Elem()

type PrivateEndpointConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput
	ToPrivateEndpointConnectionStatusPtrOutputWithContext(context.Context) PrivateEndpointConnectionStatusPtrOutput
}

type privateEndpointConnectionStatusPtr string

func PrivateEndpointConnectionStatusPtr(v string) PrivateEndpointConnectionStatusPtrInput {
	return (*privateEndpointConnectionStatusPtr)(&v)
}

func (*privateEndpointConnectionStatusPtr) ElementType() reflect.Type {
	return privateEndpointConnectionStatusPtrType
}

func (in *privateEndpointConnectionStatusPtr) ToPrivateEndpointConnectionStatusPtrOutput() PrivateEndpointConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointConnectionStatusPtrOutput)
}

func (in *privateEndpointConnectionStatusPtr) ToPrivateEndpointConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointConnectionStatusPtrOutput)
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
	ProvisioningStateDeleting  = ProvisioningState("Deleting")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStatePending   = ProvisioningState("Pending")
)

func (ProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (e ProvisioningState) ToProvisioningStateOutput() ProvisioningStateOutput {
	return pulumi.ToOutput(e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return e.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return ProvisioningState(e).ToProvisioningStateOutputWithContext(ctx).ToProvisioningStatePtrOutputWithContext(ctx)
}

func (e ProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateOutput struct{ *pulumi.OutputState }

func (ProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStateOutput) ToProvisioningStateOutput() ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningState) *ProvisioningState {
		return &v
	}).(ProvisioningStatePtrOutput)
}

func (o ProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) Elem() ProvisioningStateOutput {
	return o.ApplyT(func(v *ProvisioningState) ProvisioningState {
		if v != nil {
			return *v
		}
		var ret ProvisioningState
		return ret
	}).(ProvisioningStateOutput)
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProvisioningStateInput interface {
	pulumi.Input

	ToProvisioningStateOutput() ProvisioningStateOutput
	ToProvisioningStateOutputWithContext(context.Context) ProvisioningStateOutput
}

var provisioningStatePtrType = reflect.TypeOf((**ProvisioningState)(nil)).Elem()

type ProvisioningStatePtrInput interface {
	pulumi.Input

	ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput
	ToProvisioningStatePtrOutputWithContext(context.Context) ProvisioningStatePtrOutput
}

type provisioningStatePtr string

func ProvisioningStatePtr(v string) ProvisioningStatePtrInput {
	return (*provisioningStatePtr)(&v)
}

func (*provisioningStatePtr) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStatePtrOutput)
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStatePtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
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
	SkuNameStandard = SkuName("Standard")
	SkuNameRS0      = SkuName("RS0")
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InfrastructureEncryptionStateInput)(nil)).Elem(), InfrastructureEncryptionState("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*InfrastructureEncryptionStatePtrInput)(nil)).Elem(), InfrastructureEncryptionState("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionStatusInput)(nil)).Elem(), PrivateEndpointConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionStatusPtrInput)(nil)).Elem(), PrivateEndpointConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningStateInput)(nil)).Elem(), ProvisioningState("Succeeded"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningStatePtrInput)(nil)).Elem(), ProvisioningState("Succeeded"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNameInput)(nil)).Elem(), SkuName("Standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNamePtrInput)(nil)).Elem(), SkuName("Standard"))
	pulumi.RegisterOutputType(InfrastructureEncryptionStateOutput{})
	pulumi.RegisterOutputType(InfrastructureEncryptionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateOutput{})
	pulumi.RegisterOutputType(ProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
