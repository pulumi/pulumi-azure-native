


package powerbi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureSkuName string

const (
	AzureSkuNameS1 = AzureSkuName("S1")
)

func (AzureSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuName)(nil)).Elem()
}

func (e AzureSkuName) ToAzureSkuNameOutput() AzureSkuNameOutput {
	return pulumi.ToOutput(e).(AzureSkuNameOutput)
}

func (e AzureSkuName) ToAzureSkuNameOutputWithContext(ctx context.Context) AzureSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureSkuNameOutput)
}

func (e AzureSkuName) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return e.ToAzureSkuNamePtrOutputWithContext(context.Background())
}

func (e AzureSkuName) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return AzureSkuName(e).ToAzureSkuNameOutputWithContext(ctx).ToAzureSkuNamePtrOutputWithContext(ctx)
}

func (e AzureSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureSkuNameOutput struct{ *pulumi.OutputState }

func (AzureSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuName)(nil)).Elem()
}

func (o AzureSkuNameOutput) ToAzureSkuNameOutput() AzureSkuNameOutput {
	return o
}

func (o AzureSkuNameOutput) ToAzureSkuNameOutputWithContext(ctx context.Context) AzureSkuNameOutput {
	return o
}

func (o AzureSkuNameOutput) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return o.ToAzureSkuNamePtrOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSkuName) *AzureSkuName {
		return &v
	}).(AzureSkuNamePtrOutput)
}

func (o AzureSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureSkuNamePtrOutput struct{ *pulumi.OutputState }

func (AzureSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuName)(nil)).Elem()
}

func (o AzureSkuNamePtrOutput) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return o
}

func (o AzureSkuNamePtrOutput) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return o
}

func (o AzureSkuNamePtrOutput) Elem() AzureSkuNameOutput {
	return o.ApplyT(func(v *AzureSkuName) AzureSkuName {
		if v != nil {
			return *v
		}
		var ret AzureSkuName
		return ret
	}).(AzureSkuNameOutput)
}

func (o AzureSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureSkuNameInput interface {
	pulumi.Input

	ToAzureSkuNameOutput() AzureSkuNameOutput
	ToAzureSkuNameOutputWithContext(context.Context) AzureSkuNameOutput
}

var azureSkuNamePtrType = reflect.TypeOf((**AzureSkuName)(nil)).Elem()

type AzureSkuNamePtrInput interface {
	pulumi.Input

	ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput
	ToAzureSkuNamePtrOutputWithContext(context.Context) AzureSkuNamePtrOutput
}

type azureSkuNamePtr string

func AzureSkuNamePtr(v string) AzureSkuNamePtrInput {
	return (*azureSkuNamePtr)(&v)
}

func (*azureSkuNamePtr) ElementType() reflect.Type {
	return azureSkuNamePtrType
}

func (in *azureSkuNamePtr) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return pulumi.ToOutput(in).(AzureSkuNamePtrOutput)
}

func (in *azureSkuNamePtr) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureSkuNamePtrOutput)
}

type AzureSkuTier string

const (
	AzureSkuTierStandard = AzureSkuTier("Standard")
)

func (AzureSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuTier)(nil)).Elem()
}

func (e AzureSkuTier) ToAzureSkuTierOutput() AzureSkuTierOutput {
	return pulumi.ToOutput(e).(AzureSkuTierOutput)
}

func (e AzureSkuTier) ToAzureSkuTierOutputWithContext(ctx context.Context) AzureSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureSkuTierOutput)
}

func (e AzureSkuTier) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return e.ToAzureSkuTierPtrOutputWithContext(context.Background())
}

func (e AzureSkuTier) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return AzureSkuTier(e).ToAzureSkuTierOutputWithContext(ctx).ToAzureSkuTierPtrOutputWithContext(ctx)
}

func (e AzureSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureSkuTierOutput struct{ *pulumi.OutputState }

func (AzureSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuTier)(nil)).Elem()
}

func (o AzureSkuTierOutput) ToAzureSkuTierOutput() AzureSkuTierOutput {
	return o
}

func (o AzureSkuTierOutput) ToAzureSkuTierOutputWithContext(ctx context.Context) AzureSkuTierOutput {
	return o
}

func (o AzureSkuTierOutput) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return o.ToAzureSkuTierPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSkuTier) *AzureSkuTier {
		return &v
	}).(AzureSkuTierPtrOutput)
}

func (o AzureSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureSkuTierPtrOutput struct{ *pulumi.OutputState }

func (AzureSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuTier)(nil)).Elem()
}

func (o AzureSkuTierPtrOutput) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return o
}

func (o AzureSkuTierPtrOutput) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return o
}

func (o AzureSkuTierPtrOutput) Elem() AzureSkuTierOutput {
	return o.ApplyT(func(v *AzureSkuTier) AzureSkuTier {
		if v != nil {
			return *v
		}
		var ret AzureSkuTier
		return ret
	}).(AzureSkuTierOutput)
}

func (o AzureSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureSkuTierInput interface {
	pulumi.Input

	ToAzureSkuTierOutput() AzureSkuTierOutput
	ToAzureSkuTierOutputWithContext(context.Context) AzureSkuTierOutput
}

var azureSkuTierPtrType = reflect.TypeOf((**AzureSkuTier)(nil)).Elem()

type AzureSkuTierPtrInput interface {
	pulumi.Input

	ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput
	ToAzureSkuTierPtrOutputWithContext(context.Context) AzureSkuTierPtrOutput
}

type azureSkuTierPtr string

func AzureSkuTierPtr(v string) AzureSkuTierPtrInput {
	return (*azureSkuTierPtr)(&v)
}

func (*azureSkuTierPtr) ElementType() reflect.Type {
	return azureSkuTierPtrType
}

func (in *azureSkuTierPtr) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return pulumi.ToOutput(in).(AzureSkuTierPtrOutput)
}

func (in *azureSkuTierPtr) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureSkuTierPtrOutput)
}

type PersistedConnectionStatus string

const (
	PersistedConnectionStatusPending      = PersistedConnectionStatus("Pending")
	PersistedConnectionStatusApproved     = PersistedConnectionStatus("Approved")
	PersistedConnectionStatusRejected     = PersistedConnectionStatus("Rejected")
	PersistedConnectionStatusDisconnected = PersistedConnectionStatus("Disconnected")
)

func (PersistedConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistedConnectionStatus)(nil)).Elem()
}

func (e PersistedConnectionStatus) ToPersistedConnectionStatusOutput() PersistedConnectionStatusOutput {
	return pulumi.ToOutput(e).(PersistedConnectionStatusOutput)
}

func (e PersistedConnectionStatus) ToPersistedConnectionStatusOutputWithContext(ctx context.Context) PersistedConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PersistedConnectionStatusOutput)
}

func (e PersistedConnectionStatus) ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput {
	return e.ToPersistedConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PersistedConnectionStatus) ToPersistedConnectionStatusPtrOutputWithContext(ctx context.Context) PersistedConnectionStatusPtrOutput {
	return PersistedConnectionStatus(e).ToPersistedConnectionStatusOutputWithContext(ctx).ToPersistedConnectionStatusPtrOutputWithContext(ctx)
}

func (e PersistedConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PersistedConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PersistedConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PersistedConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PersistedConnectionStatusOutput struct{ *pulumi.OutputState }

func (PersistedConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistedConnectionStatus)(nil)).Elem()
}

func (o PersistedConnectionStatusOutput) ToPersistedConnectionStatusOutput() PersistedConnectionStatusOutput {
	return o
}

func (o PersistedConnectionStatusOutput) ToPersistedConnectionStatusOutputWithContext(ctx context.Context) PersistedConnectionStatusOutput {
	return o
}

func (o PersistedConnectionStatusOutput) ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput {
	return o.ToPersistedConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PersistedConnectionStatusOutput) ToPersistedConnectionStatusPtrOutputWithContext(ctx context.Context) PersistedConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PersistedConnectionStatus) *PersistedConnectionStatus {
		return &v
	}).(PersistedConnectionStatusPtrOutput)
}

func (o PersistedConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PersistedConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PersistedConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PersistedConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PersistedConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PersistedConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PersistedConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PersistedConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistedConnectionStatus)(nil)).Elem()
}

func (o PersistedConnectionStatusPtrOutput) ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput {
	return o
}

func (o PersistedConnectionStatusPtrOutput) ToPersistedConnectionStatusPtrOutputWithContext(ctx context.Context) PersistedConnectionStatusPtrOutput {
	return o
}

func (o PersistedConnectionStatusPtrOutput) Elem() PersistedConnectionStatusOutput {
	return o.ApplyT(func(v *PersistedConnectionStatus) PersistedConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PersistedConnectionStatus
		return ret
	}).(PersistedConnectionStatusOutput)
}

func (o PersistedConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PersistedConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PersistedConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PersistedConnectionStatusInput interface {
	pulumi.Input

	ToPersistedConnectionStatusOutput() PersistedConnectionStatusOutput
	ToPersistedConnectionStatusOutputWithContext(context.Context) PersistedConnectionStatusOutput
}

var persistedConnectionStatusPtrType = reflect.TypeOf((**PersistedConnectionStatus)(nil)).Elem()

type PersistedConnectionStatusPtrInput interface {
	pulumi.Input

	ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput
	ToPersistedConnectionStatusPtrOutputWithContext(context.Context) PersistedConnectionStatusPtrOutput
}

type persistedConnectionStatusPtr string

func PersistedConnectionStatusPtr(v string) PersistedConnectionStatusPtrInput {
	return (*persistedConnectionStatusPtr)(&v)
}

func (*persistedConnectionStatusPtr) ElementType() reflect.Type {
	return persistedConnectionStatusPtrType
}

func (in *persistedConnectionStatusPtr) ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PersistedConnectionStatusPtrOutput)
}

func (in *persistedConnectionStatusPtr) ToPersistedConnectionStatusPtrOutputWithContext(ctx context.Context) PersistedConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PersistedConnectionStatusPtrOutput)
}

type ResourceProvisioningState string

const (
	ResourceProvisioningStateCreating  = ResourceProvisioningState("Creating")
	ResourceProvisioningStateUpdating  = ResourceProvisioningState("Updating")
	ResourceProvisioningStateDeleting  = ResourceProvisioningState("Deleting")
	ResourceProvisioningStateSucceeded = ResourceProvisioningState("Succeeded")
	ResourceProvisioningStateCanceled  = ResourceProvisioningState("Canceled")
	ResourceProvisioningStateFailed    = ResourceProvisioningState("Failed")
)

func (ResourceProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceProvisioningState)(nil)).Elem()
}

func (e ResourceProvisioningState) ToResourceProvisioningStateOutput() ResourceProvisioningStateOutput {
	return pulumi.ToOutput(e).(ResourceProvisioningStateOutput)
}

func (e ResourceProvisioningState) ToResourceProvisioningStateOutputWithContext(ctx context.Context) ResourceProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceProvisioningStateOutput)
}

func (e ResourceProvisioningState) ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput {
	return e.ToResourceProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ResourceProvisioningState) ToResourceProvisioningStatePtrOutputWithContext(ctx context.Context) ResourceProvisioningStatePtrOutput {
	return ResourceProvisioningState(e).ToResourceProvisioningStateOutputWithContext(ctx).ToResourceProvisioningStatePtrOutputWithContext(ctx)
}

func (e ResourceProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceProvisioningStateOutput struct{ *pulumi.OutputState }

func (ResourceProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceProvisioningState)(nil)).Elem()
}

func (o ResourceProvisioningStateOutput) ToResourceProvisioningStateOutput() ResourceProvisioningStateOutput {
	return o
}

func (o ResourceProvisioningStateOutput) ToResourceProvisioningStateOutputWithContext(ctx context.Context) ResourceProvisioningStateOutput {
	return o
}

func (o ResourceProvisioningStateOutput) ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput {
	return o.ToResourceProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ResourceProvisioningStateOutput) ToResourceProvisioningStatePtrOutputWithContext(ctx context.Context) ResourceProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceProvisioningState) *ResourceProvisioningState {
		return &v
	}).(ResourceProvisioningStatePtrOutput)
}

func (o ResourceProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ResourceProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceProvisioningState)(nil)).Elem()
}

func (o ResourceProvisioningStatePtrOutput) ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput {
	return o
}

func (o ResourceProvisioningStatePtrOutput) ToResourceProvisioningStatePtrOutputWithContext(ctx context.Context) ResourceProvisioningStatePtrOutput {
	return o
}

func (o ResourceProvisioningStatePtrOutput) Elem() ResourceProvisioningStateOutput {
	return o.ApplyT(func(v *ResourceProvisioningState) ResourceProvisioningState {
		if v != nil {
			return *v
		}
		var ret ResourceProvisioningState
		return ret
	}).(ResourceProvisioningStateOutput)
}

func (o ResourceProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceProvisioningStateInput interface {
	pulumi.Input

	ToResourceProvisioningStateOutput() ResourceProvisioningStateOutput
	ToResourceProvisioningStateOutputWithContext(context.Context) ResourceProvisioningStateOutput
}

var resourceProvisioningStatePtrType = reflect.TypeOf((**ResourceProvisioningState)(nil)).Elem()

type ResourceProvisioningStatePtrInput interface {
	pulumi.Input

	ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput
	ToResourceProvisioningStatePtrOutputWithContext(context.Context) ResourceProvisioningStatePtrOutput
}

type resourceProvisioningStatePtr string

func ResourceProvisioningStatePtr(v string) ResourceProvisioningStatePtrInput {
	return (*resourceProvisioningStatePtr)(&v)
}

func (*resourceProvisioningStatePtr) ElementType() reflect.Type {
	return resourceProvisioningStatePtrType
}

func (in *resourceProvisioningStatePtr) ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ResourceProvisioningStatePtrOutput)
}

func (in *resourceProvisioningStatePtr) ToResourceProvisioningStatePtrOutputWithContext(ctx context.Context) ResourceProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceProvisioningStatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureSkuNameOutput{})
	pulumi.RegisterOutputType(AzureSkuNamePtrOutput{})
	pulumi.RegisterOutputType(AzureSkuTierOutput{})
	pulumi.RegisterOutputType(AzureSkuTierPtrOutput{})
	pulumi.RegisterOutputType(PersistedConnectionStatusOutput{})
	pulumi.RegisterOutputType(PersistedConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ResourceProvisioningStateOutput{})
	pulumi.RegisterOutputType(ResourceProvisioningStatePtrOutput{})
}
