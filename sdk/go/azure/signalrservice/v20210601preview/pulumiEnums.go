


package v20210601preview

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
	FeatureFlagsServiceMode            = FeatureFlags("ServiceMode")
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

type ServiceKind string

const (
	ServiceKindSignalR       = ServiceKind("SignalR")
	ServiceKindRawWebSockets = ServiceKind("RawWebSockets")
)

func (ServiceKind) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceKind)(nil)).Elem()
}

func (e ServiceKind) ToServiceKindOutput() ServiceKindOutput {
	return pulumi.ToOutput(e).(ServiceKindOutput)
}

func (e ServiceKind) ToServiceKindOutputWithContext(ctx context.Context) ServiceKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceKindOutput)
}

func (e ServiceKind) ToServiceKindPtrOutput() ServiceKindPtrOutput {
	return e.ToServiceKindPtrOutputWithContext(context.Background())
}

func (e ServiceKind) ToServiceKindPtrOutputWithContext(ctx context.Context) ServiceKindPtrOutput {
	return ServiceKind(e).ToServiceKindOutputWithContext(ctx).ToServiceKindPtrOutputWithContext(ctx)
}

func (e ServiceKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceKindOutput struct{ *pulumi.OutputState }

func (ServiceKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceKind)(nil)).Elem()
}

func (o ServiceKindOutput) ToServiceKindOutput() ServiceKindOutput {
	return o
}

func (o ServiceKindOutput) ToServiceKindOutputWithContext(ctx context.Context) ServiceKindOutput {
	return o
}

func (o ServiceKindOutput) ToServiceKindPtrOutput() ServiceKindPtrOutput {
	return o.ToServiceKindPtrOutputWithContext(context.Background())
}

func (o ServiceKindOutput) ToServiceKindPtrOutputWithContext(ctx context.Context) ServiceKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceKind) *ServiceKind {
		return &v
	}).(ServiceKindPtrOutput)
}

func (o ServiceKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceKindPtrOutput struct{ *pulumi.OutputState }

func (ServiceKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceKind)(nil)).Elem()
}

func (o ServiceKindPtrOutput) ToServiceKindPtrOutput() ServiceKindPtrOutput {
	return o
}

func (o ServiceKindPtrOutput) ToServiceKindPtrOutputWithContext(ctx context.Context) ServiceKindPtrOutput {
	return o
}

func (o ServiceKindPtrOutput) Elem() ServiceKindOutput {
	return o.ApplyT(func(v *ServiceKind) ServiceKind {
		if v != nil {
			return *v
		}
		var ret ServiceKind
		return ret
	}).(ServiceKindOutput)
}

func (o ServiceKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceKindInput interface {
	pulumi.Input

	ToServiceKindOutput() ServiceKindOutput
	ToServiceKindOutputWithContext(context.Context) ServiceKindOutput
}

var serviceKindPtrType = reflect.TypeOf((**ServiceKind)(nil)).Elem()

type ServiceKindPtrInput interface {
	pulumi.Input

	ToServiceKindPtrOutput() ServiceKindPtrOutput
	ToServiceKindPtrOutputWithContext(context.Context) ServiceKindPtrOutput
}

type serviceKindPtr string

func ServiceKindPtr(v string) ServiceKindPtrInput {
	return (*serviceKindPtr)(&v)
}

func (*serviceKindPtr) ElementType() reflect.Type {
	return serviceKindPtrType
}

func (in *serviceKindPtr) ToServiceKindPtrOutput() ServiceKindPtrOutput {
	return pulumi.ToOutput(in).(ServiceKindPtrOutput)
}

func (in *serviceKindPtr) ToServiceKindPtrOutputWithContext(ctx context.Context) ServiceKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceKindPtrOutput)
}

type SignalRRequestType string

const (
	SignalRRequestTypeClientConnection = SignalRRequestType("ClientConnection")
	SignalRRequestTypeServerConnection = SignalRRequestType("ServerConnection")
	SignalRRequestTypeRESTAPI          = SignalRRequestType("RESTAPI")
	SignalRRequestTypeTrace            = SignalRRequestType("Trace")
)

func (SignalRRequestType) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRRequestType)(nil)).Elem()
}

func (e SignalRRequestType) ToSignalRRequestTypeOutput() SignalRRequestTypeOutput {
	return pulumi.ToOutput(e).(SignalRRequestTypeOutput)
}

func (e SignalRRequestType) ToSignalRRequestTypeOutputWithContext(ctx context.Context) SignalRRequestTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SignalRRequestTypeOutput)
}

func (e SignalRRequestType) ToSignalRRequestTypePtrOutput() SignalRRequestTypePtrOutput {
	return e.ToSignalRRequestTypePtrOutputWithContext(context.Background())
}

func (e SignalRRequestType) ToSignalRRequestTypePtrOutputWithContext(ctx context.Context) SignalRRequestTypePtrOutput {
	return SignalRRequestType(e).ToSignalRRequestTypeOutputWithContext(ctx).ToSignalRRequestTypePtrOutputWithContext(ctx)
}

func (e SignalRRequestType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignalRRequestType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignalRRequestType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SignalRRequestType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SignalRRequestTypeOutput struct{ *pulumi.OutputState }

func (SignalRRequestTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRRequestType)(nil)).Elem()
}

func (o SignalRRequestTypeOutput) ToSignalRRequestTypeOutput() SignalRRequestTypeOutput {
	return o
}

func (o SignalRRequestTypeOutput) ToSignalRRequestTypeOutputWithContext(ctx context.Context) SignalRRequestTypeOutput {
	return o
}

func (o SignalRRequestTypeOutput) ToSignalRRequestTypePtrOutput() SignalRRequestTypePtrOutput {
	return o.ToSignalRRequestTypePtrOutputWithContext(context.Background())
}

func (o SignalRRequestTypeOutput) ToSignalRRequestTypePtrOutputWithContext(ctx context.Context) SignalRRequestTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignalRRequestType) *SignalRRequestType {
		return &v
	}).(SignalRRequestTypePtrOutput)
}

func (o SignalRRequestTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SignalRRequestTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignalRRequestType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SignalRRequestTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignalRRequestTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignalRRequestType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SignalRRequestTypePtrOutput struct{ *pulumi.OutputState }

func (SignalRRequestTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRRequestType)(nil)).Elem()
}

func (o SignalRRequestTypePtrOutput) ToSignalRRequestTypePtrOutput() SignalRRequestTypePtrOutput {
	return o
}

func (o SignalRRequestTypePtrOutput) ToSignalRRequestTypePtrOutputWithContext(ctx context.Context) SignalRRequestTypePtrOutput {
	return o
}

func (o SignalRRequestTypePtrOutput) Elem() SignalRRequestTypeOutput {
	return o.ApplyT(func(v *SignalRRequestType) SignalRRequestType {
		if v != nil {
			return *v
		}
		var ret SignalRRequestType
		return ret
	}).(SignalRRequestTypeOutput)
}

func (o SignalRRequestTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignalRRequestTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SignalRRequestType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SignalRRequestTypeInput interface {
	pulumi.Input

	ToSignalRRequestTypeOutput() SignalRRequestTypeOutput
	ToSignalRRequestTypeOutputWithContext(context.Context) SignalRRequestTypeOutput
}

var signalRRequestTypePtrType = reflect.TypeOf((**SignalRRequestType)(nil)).Elem()

type SignalRRequestTypePtrInput interface {
	pulumi.Input

	ToSignalRRequestTypePtrOutput() SignalRRequestTypePtrOutput
	ToSignalRRequestTypePtrOutputWithContext(context.Context) SignalRRequestTypePtrOutput
}

type signalRRequestTypePtr string

func SignalRRequestTypePtr(v string) SignalRRequestTypePtrInput {
	return (*signalRRequestTypePtr)(&v)
}

func (*signalRRequestTypePtr) ElementType() reflect.Type {
	return signalRRequestTypePtrType
}

func (in *signalRRequestTypePtr) ToSignalRRequestTypePtrOutput() SignalRRequestTypePtrOutput {
	return pulumi.ToOutput(in).(SignalRRequestTypePtrOutput)
}

func (in *signalRRequestTypePtr) ToSignalRRequestTypePtrOutputWithContext(ctx context.Context) SignalRRequestTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SignalRRequestTypePtrOutput)
}

type SignalRSkuTier string

const (
	SignalRSkuTierFree     = SignalRSkuTier("Free")
	SignalRSkuTierBasic    = SignalRSkuTier("Basic")
	SignalRSkuTierStandard = SignalRSkuTier("Standard")
	SignalRSkuTierPremium  = SignalRSkuTier("Premium")
)

func (SignalRSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRSkuTier)(nil)).Elem()
}

func (e SignalRSkuTier) ToSignalRSkuTierOutput() SignalRSkuTierOutput {
	return pulumi.ToOutput(e).(SignalRSkuTierOutput)
}

func (e SignalRSkuTier) ToSignalRSkuTierOutputWithContext(ctx context.Context) SignalRSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SignalRSkuTierOutput)
}

func (e SignalRSkuTier) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return e.ToSignalRSkuTierPtrOutputWithContext(context.Background())
}

func (e SignalRSkuTier) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return SignalRSkuTier(e).ToSignalRSkuTierOutputWithContext(ctx).ToSignalRSkuTierPtrOutputWithContext(ctx)
}

func (e SignalRSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignalRSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignalRSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SignalRSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SignalRSkuTierOutput struct{ *pulumi.OutputState }

func (SignalRSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRSkuTier)(nil)).Elem()
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierOutput() SignalRSkuTierOutput {
	return o
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierOutputWithContext(ctx context.Context) SignalRSkuTierOutput {
	return o
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return o.ToSignalRSkuTierPtrOutputWithContext(context.Background())
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignalRSkuTier) *SignalRSkuTier {
		return &v
	}).(SignalRSkuTierPtrOutput)
}

func (o SignalRSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SignalRSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignalRSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SignalRSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignalRSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignalRSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SignalRSkuTierPtrOutput struct{ *pulumi.OutputState }

func (SignalRSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRSkuTier)(nil)).Elem()
}

func (o SignalRSkuTierPtrOutput) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return o
}

func (o SignalRSkuTierPtrOutput) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return o
}

func (o SignalRSkuTierPtrOutput) Elem() SignalRSkuTierOutput {
	return o.ApplyT(func(v *SignalRSkuTier) SignalRSkuTier {
		if v != nil {
			return *v
		}
		var ret SignalRSkuTier
		return ret
	}).(SignalRSkuTierOutput)
}

func (o SignalRSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignalRSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SignalRSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SignalRSkuTierInput interface {
	pulumi.Input

	ToSignalRSkuTierOutput() SignalRSkuTierOutput
	ToSignalRSkuTierOutputWithContext(context.Context) SignalRSkuTierOutput
}

var signalRSkuTierPtrType = reflect.TypeOf((**SignalRSkuTier)(nil)).Elem()

type SignalRSkuTierPtrInput interface {
	pulumi.Input

	ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput
	ToSignalRSkuTierPtrOutputWithContext(context.Context) SignalRSkuTierPtrOutput
}

type signalRSkuTierPtr string

func SignalRSkuTierPtr(v string) SignalRSkuTierPtrInput {
	return (*signalRSkuTierPtr)(&v)
}

func (*signalRSkuTierPtr) ElementType() reflect.Type {
	return signalRSkuTierPtrType
}

func (in *signalRSkuTierPtr) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return pulumi.ToOutput(in).(SignalRSkuTierPtrOutput)
}

func (in *signalRSkuTierPtr) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SignalRSkuTierPtrOutput)
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ACLActionInput)(nil)).Elem(), ACLAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*ACLActionPtrInput)(nil)).Elem(), ACLAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureFlagsInput)(nil)).Elem(), FeatureFlags("ServiceMode"))
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureFlagsPtrInput)(nil)).Elem(), FeatureFlags("ServiceMode"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityTypeInput)(nil)).Elem(), ManagedIdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityTypePtrInput)(nil)).Elem(), ManagedIdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceKindInput)(nil)).Elem(), ServiceKind("SignalR"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceKindPtrInput)(nil)).Elem(), ServiceKind("SignalR"))
	pulumi.RegisterInputType(reflect.TypeOf((*SignalRRequestTypeInput)(nil)).Elem(), SignalRRequestType("ClientConnection"))
	pulumi.RegisterInputType(reflect.TypeOf((*SignalRRequestTypePtrInput)(nil)).Elem(), SignalRRequestType("ClientConnection"))
	pulumi.RegisterInputType(reflect.TypeOf((*SignalRSkuTierInput)(nil)).Elem(), SignalRSkuTier("Free"))
	pulumi.RegisterInputType(reflect.TypeOf((*SignalRSkuTierPtrInput)(nil)).Elem(), SignalRSkuTier("Free"))
	pulumi.RegisterInputType(reflect.TypeOf((*UpstreamAuthTypeInput)(nil)).Elem(), UpstreamAuthType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*UpstreamAuthTypePtrInput)(nil)).Elem(), UpstreamAuthType("None"))
	pulumi.RegisterOutputType(ACLActionOutput{})
	pulumi.RegisterOutputType(ACLActionPtrOutput{})
	pulumi.RegisterOutputType(FeatureFlagsOutput{})
	pulumi.RegisterOutputType(FeatureFlagsPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ServiceKindOutput{})
	pulumi.RegisterOutputType(ServiceKindPtrOutput{})
	pulumi.RegisterOutputType(SignalRRequestTypeOutput{})
	pulumi.RegisterOutputType(SignalRRequestTypePtrOutput{})
	pulumi.RegisterOutputType(SignalRSkuTierOutput{})
	pulumi.RegisterOutputType(SignalRSkuTierPtrOutput{})
	pulumi.RegisterOutputType(UpstreamAuthTypeOutput{})
	pulumi.RegisterOutputType(UpstreamAuthTypePtrOutput{})
}
