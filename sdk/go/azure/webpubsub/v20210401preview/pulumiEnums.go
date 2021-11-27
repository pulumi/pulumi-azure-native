


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ACLAction string

const (
	ACLActionAllow = ACLAction("Allow")
	ACLActionDeny  = ACLAction("Deny")
)

func (ACLAction) ElementType() reflect.Type {
	return reflect.TypeOf((*ACLAction)(nil)).Elem()
}

func (e ACLAction) ToACLActionOutput() ACLActionOutput {
	return pulumi.ToOutput(e).(ACLActionOutput)
}

func (e ACLAction) ToACLActionOutputWithContext(ctx context.Context) ACLActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ACLActionOutput)
}

func (e ACLAction) ToACLActionPtrOutput() ACLActionPtrOutput {
	return e.ToACLActionPtrOutputWithContext(context.Background())
}

func (e ACLAction) ToACLActionPtrOutputWithContext(ctx context.Context) ACLActionPtrOutput {
	return ACLAction(e).ToACLActionOutputWithContext(ctx).ToACLActionPtrOutputWithContext(ctx)
}

func (e ACLAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ACLAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ACLAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ACLAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ACLActionOutput struct{ *pulumi.OutputState }

func (ACLActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACLAction)(nil)).Elem()
}

func (o ACLActionOutput) ToACLActionOutput() ACLActionOutput {
	return o
}

func (o ACLActionOutput) ToACLActionOutputWithContext(ctx context.Context) ACLActionOutput {
	return o
}

func (o ACLActionOutput) ToACLActionPtrOutput() ACLActionPtrOutput {
	return o.ToACLActionPtrOutputWithContext(context.Background())
}

func (o ACLActionOutput) ToACLActionPtrOutputWithContext(ctx context.Context) ACLActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ACLAction) *ACLAction {
		return &v
	}).(ACLActionPtrOutput)
}

func (o ACLActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ACLActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ACLAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ACLActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ACLActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ACLAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ACLActionPtrOutput struct{ *pulumi.OutputState }

func (ACLActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ACLAction)(nil)).Elem()
}

func (o ACLActionPtrOutput) ToACLActionPtrOutput() ACLActionPtrOutput {
	return o
}

func (o ACLActionPtrOutput) ToACLActionPtrOutputWithContext(ctx context.Context) ACLActionPtrOutput {
	return o
}

func (o ACLActionPtrOutput) Elem() ACLActionOutput {
	return o.ApplyT(func(v *ACLAction) ACLAction {
		if v != nil {
			return *v
		}
		var ret ACLAction
		return ret
	}).(ACLActionOutput)
}

func (o ACLActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ACLActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ACLAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ACLActionInput interface {
	pulumi.Input

	ToACLActionOutput() ACLActionOutput
	ToACLActionOutputWithContext(context.Context) ACLActionOutput
}

var aclactionPtrType = reflect.TypeOf((**ACLAction)(nil)).Elem()

type ACLActionPtrInput interface {
	pulumi.Input

	ToACLActionPtrOutput() ACLActionPtrOutput
	ToACLActionPtrOutputWithContext(context.Context) ACLActionPtrOutput
}

type aclactionPtr string

func ACLActionPtr(v string) ACLActionPtrInput {
	return (*aclactionPtr)(&v)
}

func (*aclactionPtr) ElementType() reflect.Type {
	return aclactionPtrType
}

func (in *aclactionPtr) ToACLActionPtrOutput() ACLActionPtrOutput {
	return pulumi.ToOutput(in).(ACLActionPtrOutput)
}

func (in *aclactionPtr) ToACLActionPtrOutputWithContext(ctx context.Context) ACLActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ACLActionPtrOutput)
}

type FeatureFlags string

const (
	FeatureFlagsEnableConnectivityLogs = FeatureFlags("EnableConnectivityLogs")
	FeatureFlagsEnableMessagingLogs    = FeatureFlags("EnableMessagingLogs")
	FeatureFlagsEnableLiveTrace        = FeatureFlags("EnableLiveTrace")
)

func (FeatureFlags) ElementType() reflect.Type {
	return reflect.TypeOf((*FeatureFlags)(nil)).Elem()
}

func (e FeatureFlags) ToFeatureFlagsOutput() FeatureFlagsOutput {
	return pulumi.ToOutput(e).(FeatureFlagsOutput)
}

func (e FeatureFlags) ToFeatureFlagsOutputWithContext(ctx context.Context) FeatureFlagsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FeatureFlagsOutput)
}

func (e FeatureFlags) ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput {
	return e.ToFeatureFlagsPtrOutputWithContext(context.Background())
}

func (e FeatureFlags) ToFeatureFlagsPtrOutputWithContext(ctx context.Context) FeatureFlagsPtrOutput {
	return FeatureFlags(e).ToFeatureFlagsOutputWithContext(ctx).ToFeatureFlagsPtrOutputWithContext(ctx)
}

func (e FeatureFlags) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FeatureFlags) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FeatureFlags) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FeatureFlags) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FeatureFlagsOutput struct{ *pulumi.OutputState }

func (FeatureFlagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FeatureFlags)(nil)).Elem()
}

func (o FeatureFlagsOutput) ToFeatureFlagsOutput() FeatureFlagsOutput {
	return o
}

func (o FeatureFlagsOutput) ToFeatureFlagsOutputWithContext(ctx context.Context) FeatureFlagsOutput {
	return o
}

func (o FeatureFlagsOutput) ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput {
	return o.ToFeatureFlagsPtrOutputWithContext(context.Background())
}

func (o FeatureFlagsOutput) ToFeatureFlagsPtrOutputWithContext(ctx context.Context) FeatureFlagsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FeatureFlags) *FeatureFlags {
		return &v
	}).(FeatureFlagsPtrOutput)
}

func (o FeatureFlagsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FeatureFlagsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FeatureFlags) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FeatureFlagsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FeatureFlagsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FeatureFlags) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FeatureFlagsPtrOutput struct{ *pulumi.OutputState }

func (FeatureFlagsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureFlags)(nil)).Elem()
}

func (o FeatureFlagsPtrOutput) ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput {
	return o
}

func (o FeatureFlagsPtrOutput) ToFeatureFlagsPtrOutputWithContext(ctx context.Context) FeatureFlagsPtrOutput {
	return o
}

func (o FeatureFlagsPtrOutput) Elem() FeatureFlagsOutput {
	return o.ApplyT(func(v *FeatureFlags) FeatureFlags {
		if v != nil {
			return *v
		}
		var ret FeatureFlags
		return ret
	}).(FeatureFlagsOutput)
}

func (o FeatureFlagsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FeatureFlagsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FeatureFlags) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FeatureFlagsInput interface {
	pulumi.Input

	ToFeatureFlagsOutput() FeatureFlagsOutput
	ToFeatureFlagsOutputWithContext(context.Context) FeatureFlagsOutput
}

var featureFlagsPtrType = reflect.TypeOf((**FeatureFlags)(nil)).Elem()

type FeatureFlagsPtrInput interface {
	pulumi.Input

	ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput
	ToFeatureFlagsPtrOutputWithContext(context.Context) FeatureFlagsPtrOutput
}

type featureFlagsPtr string

func FeatureFlagsPtr(v string) FeatureFlagsPtrInput {
	return (*featureFlagsPtr)(&v)
}

func (*featureFlagsPtr) ElementType() reflect.Type {
	return featureFlagsPtrType
}

func (in *featureFlagsPtr) ToFeatureFlagsPtrOutput() FeatureFlagsPtrOutput {
	return pulumi.ToOutput(in).(FeatureFlagsPtrOutput)
}

func (in *featureFlagsPtr) ToFeatureFlagsPtrOutputWithContext(ctx context.Context) FeatureFlagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FeatureFlagsPtrOutput)
}

type ManagedIdentityType string

const (
	ManagedIdentityTypeNone           = ManagedIdentityType("None")
	ManagedIdentityTypeSystemAssigned = ManagedIdentityType("SystemAssigned")
	ManagedIdentityTypeUserAssigned   = ManagedIdentityType("UserAssigned")
)

func (ManagedIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityType)(nil)).Elem()
}

func (e ManagedIdentityType) ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedIdentityTypeOutput)
}

func (e ManagedIdentityType) ToManagedIdentityTypeOutputWithContext(ctx context.Context) ManagedIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedIdentityTypeOutput)
}

func (e ManagedIdentityType) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return e.ToManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedIdentityType) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return ManagedIdentityType(e).ToManagedIdentityTypeOutputWithContext(ctx).ToManagedIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityType)(nil)).Elem()
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput {
	return o
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypeOutputWithContext(ctx context.Context) ManagedIdentityTypeOutput {
	return o
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return o.ToManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityType) *ManagedIdentityType {
		return &v
	}).(ManagedIdentityTypePtrOutput)
}

func (o ManagedIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityType)(nil)).Elem()
}

func (o ManagedIdentityTypePtrOutput) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return o
}

func (o ManagedIdentityTypePtrOutput) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return o
}

func (o ManagedIdentityTypePtrOutput) Elem() ManagedIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedIdentityType) ManagedIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityType
		return ret
	}).(ManagedIdentityTypeOutput)
}

func (o ManagedIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedIdentityTypeInput interface {
	pulumi.Input

	ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput
	ToManagedIdentityTypeOutputWithContext(context.Context) ManagedIdentityTypeOutput
}

var managedIdentityTypePtrType = reflect.TypeOf((**ManagedIdentityType)(nil)).Elem()

type ManagedIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput
	ToManagedIdentityTypePtrOutputWithContext(context.Context) ManagedIdentityTypePtrOutput
}

type managedIdentityTypePtr string

func ManagedIdentityTypePtr(v string) ManagedIdentityTypePtrInput {
	return (*managedIdentityTypePtr)(&v)
}

func (*managedIdentityTypePtr) ElementType() reflect.Type {
	return managedIdentityTypePtrType
}

func (in *managedIdentityTypePtr) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedIdentityTypePtrOutput)
}

func (in *managedIdentityTypePtr) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedIdentityTypePtrOutput)
}

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

func (PrivateLinkServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateLinkServiceConnectionStatusOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateLinkServiceConnectionStatusOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return e.ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return PrivateLinkServiceConnectionStatus(e).ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx).ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateLinkServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatus) *PrivateLinkServiceConnectionStatus {
		return &v
	}).(PrivateLinkServiceConnectionStatusPtrOutput)
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) Elem() PrivateLinkServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatus) PrivateLinkServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatus
		return ret
	}).(PrivateLinkServiceConnectionStatusOutput)
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateLinkServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateLinkServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput
	ToPrivateLinkServiceConnectionStatusOutputWithContext(context.Context) PrivateLinkServiceConnectionStatusOutput
}

var privateLinkServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateLinkServiceConnectionStatus)(nil)).Elem()

type PrivateLinkServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput
	ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatusPtrOutput
}

type privateLinkServiceConnectionStatusPtr string

func PrivateLinkServiceConnectionStatusPtr(v string) PrivateLinkServiceConnectionStatusPtrInput {
	return (*privateLinkServiceConnectionStatusPtr)(&v)
}

func (*privateLinkServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateLinkServiceConnectionStatusPtrType
}

func (in *privateLinkServiceConnectionStatusPtr) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateLinkServiceConnectionStatusPtrOutput)
}

func (in *privateLinkServiceConnectionStatusPtr) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateLinkServiceConnectionStatusPtrOutput)
}

type UpstreamAuthType string

const (
	UpstreamAuthTypeNone            = UpstreamAuthType("None")
	UpstreamAuthTypeManagedIdentity = UpstreamAuthType("ManagedIdentity")
)

func (UpstreamAuthType) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamAuthType)(nil)).Elem()
}

func (e UpstreamAuthType) ToUpstreamAuthTypeOutput() UpstreamAuthTypeOutput {
	return pulumi.ToOutput(e).(UpstreamAuthTypeOutput)
}

func (e UpstreamAuthType) ToUpstreamAuthTypeOutputWithContext(ctx context.Context) UpstreamAuthTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UpstreamAuthTypeOutput)
}

func (e UpstreamAuthType) ToUpstreamAuthTypePtrOutput() UpstreamAuthTypePtrOutput {
	return e.ToUpstreamAuthTypePtrOutputWithContext(context.Background())
}

func (e UpstreamAuthType) ToUpstreamAuthTypePtrOutputWithContext(ctx context.Context) UpstreamAuthTypePtrOutput {
	return UpstreamAuthType(e).ToUpstreamAuthTypeOutputWithContext(ctx).ToUpstreamAuthTypePtrOutputWithContext(ctx)
}

func (e UpstreamAuthType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpstreamAuthType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpstreamAuthType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UpstreamAuthType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UpstreamAuthTypeOutput struct{ *pulumi.OutputState }

func (UpstreamAuthTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamAuthType)(nil)).Elem()
}

func (o UpstreamAuthTypeOutput) ToUpstreamAuthTypeOutput() UpstreamAuthTypeOutput {
	return o
}

func (o UpstreamAuthTypeOutput) ToUpstreamAuthTypeOutputWithContext(ctx context.Context) UpstreamAuthTypeOutput {
	return o
}

func (o UpstreamAuthTypeOutput) ToUpstreamAuthTypePtrOutput() UpstreamAuthTypePtrOutput {
	return o.ToUpstreamAuthTypePtrOutputWithContext(context.Background())
}

func (o UpstreamAuthTypeOutput) ToUpstreamAuthTypePtrOutputWithContext(ctx context.Context) UpstreamAuthTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpstreamAuthType) *UpstreamAuthType {
		return &v
	}).(UpstreamAuthTypePtrOutput)
}

func (o UpstreamAuthTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UpstreamAuthTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpstreamAuthType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UpstreamAuthTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpstreamAuthTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpstreamAuthType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UpstreamAuthTypePtrOutput struct{ *pulumi.OutputState }

func (UpstreamAuthTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpstreamAuthType)(nil)).Elem()
}

func (o UpstreamAuthTypePtrOutput) ToUpstreamAuthTypePtrOutput() UpstreamAuthTypePtrOutput {
	return o
}

func (o UpstreamAuthTypePtrOutput) ToUpstreamAuthTypePtrOutputWithContext(ctx context.Context) UpstreamAuthTypePtrOutput {
	return o
}

func (o UpstreamAuthTypePtrOutput) Elem() UpstreamAuthTypeOutput {
	return o.ApplyT(func(v *UpstreamAuthType) UpstreamAuthType {
		if v != nil {
			return *v
		}
		var ret UpstreamAuthType
		return ret
	}).(UpstreamAuthTypeOutput)
}

func (o UpstreamAuthTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpstreamAuthTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UpstreamAuthType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UpstreamAuthTypeInput interface {
	pulumi.Input

	ToUpstreamAuthTypeOutput() UpstreamAuthTypeOutput
	ToUpstreamAuthTypeOutputWithContext(context.Context) UpstreamAuthTypeOutput
}

var upstreamAuthTypePtrType = reflect.TypeOf((**UpstreamAuthType)(nil)).Elem()

type UpstreamAuthTypePtrInput interface {
	pulumi.Input

	ToUpstreamAuthTypePtrOutput() UpstreamAuthTypePtrOutput
	ToUpstreamAuthTypePtrOutputWithContext(context.Context) UpstreamAuthTypePtrOutput
}

type upstreamAuthTypePtr string

func UpstreamAuthTypePtr(v string) UpstreamAuthTypePtrInput {
	return (*upstreamAuthTypePtr)(&v)
}

func (*upstreamAuthTypePtr) ElementType() reflect.Type {
	return upstreamAuthTypePtrType
}

func (in *upstreamAuthTypePtr) ToUpstreamAuthTypePtrOutput() UpstreamAuthTypePtrOutput {
	return pulumi.ToOutput(in).(UpstreamAuthTypePtrOutput)
}

func (in *upstreamAuthTypePtr) ToUpstreamAuthTypePtrOutputWithContext(ctx context.Context) UpstreamAuthTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UpstreamAuthTypePtrOutput)
}

type WebPubSubRequestType string

const (
	WebPubSubRequestTypeClientConnection = WebPubSubRequestType("ClientConnection")
	WebPubSubRequestTypeServerConnection = WebPubSubRequestType("ServerConnection")
	WebPubSubRequestTypeRESTAPI          = WebPubSubRequestType("RESTAPI")
	WebPubSubRequestTypeTrace            = WebPubSubRequestType("Trace")
)

func (WebPubSubRequestType) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubRequestType)(nil)).Elem()
}

func (e WebPubSubRequestType) ToWebPubSubRequestTypeOutput() WebPubSubRequestTypeOutput {
	return pulumi.ToOutput(e).(WebPubSubRequestTypeOutput)
}

func (e WebPubSubRequestType) ToWebPubSubRequestTypeOutputWithContext(ctx context.Context) WebPubSubRequestTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebPubSubRequestTypeOutput)
}

func (e WebPubSubRequestType) ToWebPubSubRequestTypePtrOutput() WebPubSubRequestTypePtrOutput {
	return e.ToWebPubSubRequestTypePtrOutputWithContext(context.Background())
}

func (e WebPubSubRequestType) ToWebPubSubRequestTypePtrOutputWithContext(ctx context.Context) WebPubSubRequestTypePtrOutput {
	return WebPubSubRequestType(e).ToWebPubSubRequestTypeOutputWithContext(ctx).ToWebPubSubRequestTypePtrOutputWithContext(ctx)
}

func (e WebPubSubRequestType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebPubSubRequestType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebPubSubRequestType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebPubSubRequestType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebPubSubRequestTypeOutput struct{ *pulumi.OutputState }

func (WebPubSubRequestTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubRequestType)(nil)).Elem()
}

func (o WebPubSubRequestTypeOutput) ToWebPubSubRequestTypeOutput() WebPubSubRequestTypeOutput {
	return o
}

func (o WebPubSubRequestTypeOutput) ToWebPubSubRequestTypeOutputWithContext(ctx context.Context) WebPubSubRequestTypeOutput {
	return o
}

func (o WebPubSubRequestTypeOutput) ToWebPubSubRequestTypePtrOutput() WebPubSubRequestTypePtrOutput {
	return o.ToWebPubSubRequestTypePtrOutputWithContext(context.Background())
}

func (o WebPubSubRequestTypeOutput) ToWebPubSubRequestTypePtrOutputWithContext(ctx context.Context) WebPubSubRequestTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebPubSubRequestType) *WebPubSubRequestType {
		return &v
	}).(WebPubSubRequestTypePtrOutput)
}

func (o WebPubSubRequestTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebPubSubRequestTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebPubSubRequestType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebPubSubRequestTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebPubSubRequestTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebPubSubRequestType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebPubSubRequestTypePtrOutput struct{ *pulumi.OutputState }

func (WebPubSubRequestTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubRequestType)(nil)).Elem()
}

func (o WebPubSubRequestTypePtrOutput) ToWebPubSubRequestTypePtrOutput() WebPubSubRequestTypePtrOutput {
	return o
}

func (o WebPubSubRequestTypePtrOutput) ToWebPubSubRequestTypePtrOutputWithContext(ctx context.Context) WebPubSubRequestTypePtrOutput {
	return o
}

func (o WebPubSubRequestTypePtrOutput) Elem() WebPubSubRequestTypeOutput {
	return o.ApplyT(func(v *WebPubSubRequestType) WebPubSubRequestType {
		if v != nil {
			return *v
		}
		var ret WebPubSubRequestType
		return ret
	}).(WebPubSubRequestTypeOutput)
}

func (o WebPubSubRequestTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebPubSubRequestTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebPubSubRequestType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebPubSubRequestTypeInput interface {
	pulumi.Input

	ToWebPubSubRequestTypeOutput() WebPubSubRequestTypeOutput
	ToWebPubSubRequestTypeOutputWithContext(context.Context) WebPubSubRequestTypeOutput
}

var webPubSubRequestTypePtrType = reflect.TypeOf((**WebPubSubRequestType)(nil)).Elem()

type WebPubSubRequestTypePtrInput interface {
	pulumi.Input

	ToWebPubSubRequestTypePtrOutput() WebPubSubRequestTypePtrOutput
	ToWebPubSubRequestTypePtrOutputWithContext(context.Context) WebPubSubRequestTypePtrOutput
}

type webPubSubRequestTypePtr string

func WebPubSubRequestTypePtr(v string) WebPubSubRequestTypePtrInput {
	return (*webPubSubRequestTypePtr)(&v)
}

func (*webPubSubRequestTypePtr) ElementType() reflect.Type {
	return webPubSubRequestTypePtrType
}

func (in *webPubSubRequestTypePtr) ToWebPubSubRequestTypePtrOutput() WebPubSubRequestTypePtrOutput {
	return pulumi.ToOutput(in).(WebPubSubRequestTypePtrOutput)
}

func (in *webPubSubRequestTypePtr) ToWebPubSubRequestTypePtrOutputWithContext(ctx context.Context) WebPubSubRequestTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebPubSubRequestTypePtrOutput)
}

type WebPubSubSkuTier string

const (
	WebPubSubSkuTierFree     = WebPubSubSkuTier("Free")
	WebPubSubSkuTierBasic    = WebPubSubSkuTier("Basic")
	WebPubSubSkuTierStandard = WebPubSubSkuTier("Standard")
	WebPubSubSkuTierPremium  = WebPubSubSkuTier("Premium")
)

func (WebPubSubSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubSkuTier)(nil)).Elem()
}

func (e WebPubSubSkuTier) ToWebPubSubSkuTierOutput() WebPubSubSkuTierOutput {
	return pulumi.ToOutput(e).(WebPubSubSkuTierOutput)
}

func (e WebPubSubSkuTier) ToWebPubSubSkuTierOutputWithContext(ctx context.Context) WebPubSubSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebPubSubSkuTierOutput)
}

func (e WebPubSubSkuTier) ToWebPubSubSkuTierPtrOutput() WebPubSubSkuTierPtrOutput {
	return e.ToWebPubSubSkuTierPtrOutputWithContext(context.Background())
}

func (e WebPubSubSkuTier) ToWebPubSubSkuTierPtrOutputWithContext(ctx context.Context) WebPubSubSkuTierPtrOutput {
	return WebPubSubSkuTier(e).ToWebPubSubSkuTierOutputWithContext(ctx).ToWebPubSubSkuTierPtrOutputWithContext(ctx)
}

func (e WebPubSubSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebPubSubSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebPubSubSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebPubSubSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebPubSubSkuTierOutput struct{ *pulumi.OutputState }

func (WebPubSubSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubSkuTier)(nil)).Elem()
}

func (o WebPubSubSkuTierOutput) ToWebPubSubSkuTierOutput() WebPubSubSkuTierOutput {
	return o
}

func (o WebPubSubSkuTierOutput) ToWebPubSubSkuTierOutputWithContext(ctx context.Context) WebPubSubSkuTierOutput {
	return o
}

func (o WebPubSubSkuTierOutput) ToWebPubSubSkuTierPtrOutput() WebPubSubSkuTierPtrOutput {
	return o.ToWebPubSubSkuTierPtrOutputWithContext(context.Background())
}

func (o WebPubSubSkuTierOutput) ToWebPubSubSkuTierPtrOutputWithContext(ctx context.Context) WebPubSubSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebPubSubSkuTier) *WebPubSubSkuTier {
		return &v
	}).(WebPubSubSkuTierPtrOutput)
}

func (o WebPubSubSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebPubSubSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebPubSubSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebPubSubSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebPubSubSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebPubSubSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebPubSubSkuTierPtrOutput struct{ *pulumi.OutputState }

func (WebPubSubSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubSkuTier)(nil)).Elem()
}

func (o WebPubSubSkuTierPtrOutput) ToWebPubSubSkuTierPtrOutput() WebPubSubSkuTierPtrOutput {
	return o
}

func (o WebPubSubSkuTierPtrOutput) ToWebPubSubSkuTierPtrOutputWithContext(ctx context.Context) WebPubSubSkuTierPtrOutput {
	return o
}

func (o WebPubSubSkuTierPtrOutput) Elem() WebPubSubSkuTierOutput {
	return o.ApplyT(func(v *WebPubSubSkuTier) WebPubSubSkuTier {
		if v != nil {
			return *v
		}
		var ret WebPubSubSkuTier
		return ret
	}).(WebPubSubSkuTierOutput)
}

func (o WebPubSubSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebPubSubSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebPubSubSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebPubSubSkuTierInput interface {
	pulumi.Input

	ToWebPubSubSkuTierOutput() WebPubSubSkuTierOutput
	ToWebPubSubSkuTierOutputWithContext(context.Context) WebPubSubSkuTierOutput
}

var webPubSubSkuTierPtrType = reflect.TypeOf((**WebPubSubSkuTier)(nil)).Elem()

type WebPubSubSkuTierPtrInput interface {
	pulumi.Input

	ToWebPubSubSkuTierPtrOutput() WebPubSubSkuTierPtrOutput
	ToWebPubSubSkuTierPtrOutputWithContext(context.Context) WebPubSubSkuTierPtrOutput
}

type webPubSubSkuTierPtr string

func WebPubSubSkuTierPtr(v string) WebPubSubSkuTierPtrInput {
	return (*webPubSubSkuTierPtr)(&v)
}

func (*webPubSubSkuTierPtr) ElementType() reflect.Type {
	return webPubSubSkuTierPtrType
}

func (in *webPubSubSkuTierPtr) ToWebPubSubSkuTierPtrOutput() WebPubSubSkuTierPtrOutput {
	return pulumi.ToOutput(in).(WebPubSubSkuTierPtrOutput)
}

func (in *webPubSubSkuTierPtr) ToWebPubSubSkuTierPtrOutputWithContext(ctx context.Context) WebPubSubSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebPubSubSkuTierPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ACLActionOutput{})
	pulumi.RegisterOutputType(ACLActionPtrOutput{})
	pulumi.RegisterOutputType(FeatureFlagsOutput{})
	pulumi.RegisterOutputType(FeatureFlagsPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(UpstreamAuthTypeOutput{})
	pulumi.RegisterOutputType(UpstreamAuthTypePtrOutput{})
	pulumi.RegisterOutputType(WebPubSubRequestTypeOutput{})
	pulumi.RegisterOutputType(WebPubSubRequestTypePtrOutput{})
	pulumi.RegisterOutputType(WebPubSubSkuTierOutput{})
	pulumi.RegisterOutputType(WebPubSubSkuTierPtrOutput{})
}
