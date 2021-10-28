


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AadAuthFailureMode string

const (
	// Indicates that requests that failed authentication should be presented with an HTTP status code of 403 (Forbidden).
	AadAuthFailureModeHttp403 = AadAuthFailureMode("http403")
	// Indicates that requests that failed authentication should be presented with an HTTP status code of 401 (Unauthorized) and present a Bearer Challenge.
	AadAuthFailureModeHttp401WithBearerChallenge = AadAuthFailureMode("http401WithBearerChallenge")
)

func (AadAuthFailureMode) ElementType() reflect.Type {
	return reflect.TypeOf((*AadAuthFailureMode)(nil)).Elem()
}

func (e AadAuthFailureMode) ToAadAuthFailureModeOutput() AadAuthFailureModeOutput {
	return pulumi.ToOutput(e).(AadAuthFailureModeOutput)
}

func (e AadAuthFailureMode) ToAadAuthFailureModeOutputWithContext(ctx context.Context) AadAuthFailureModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AadAuthFailureModeOutput)
}

func (e AadAuthFailureMode) ToAadAuthFailureModePtrOutput() AadAuthFailureModePtrOutput {
	return e.ToAadAuthFailureModePtrOutputWithContext(context.Background())
}

func (e AadAuthFailureMode) ToAadAuthFailureModePtrOutputWithContext(ctx context.Context) AadAuthFailureModePtrOutput {
	return AadAuthFailureMode(e).ToAadAuthFailureModeOutputWithContext(ctx).ToAadAuthFailureModePtrOutputWithContext(ctx)
}

func (e AadAuthFailureMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AadAuthFailureMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AadAuthFailureMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AadAuthFailureMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AadAuthFailureModeOutput struct{ *pulumi.OutputState }

func (AadAuthFailureModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AadAuthFailureMode)(nil)).Elem()
}

func (o AadAuthFailureModeOutput) ToAadAuthFailureModeOutput() AadAuthFailureModeOutput {
	return o
}

func (o AadAuthFailureModeOutput) ToAadAuthFailureModeOutputWithContext(ctx context.Context) AadAuthFailureModeOutput {
	return o
}

func (o AadAuthFailureModeOutput) ToAadAuthFailureModePtrOutput() AadAuthFailureModePtrOutput {
	return o.ToAadAuthFailureModePtrOutputWithContext(context.Background())
}

func (o AadAuthFailureModeOutput) ToAadAuthFailureModePtrOutputWithContext(ctx context.Context) AadAuthFailureModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AadAuthFailureMode) *AadAuthFailureMode {
		return &v
	}).(AadAuthFailureModePtrOutput)
}

func (o AadAuthFailureModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AadAuthFailureModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AadAuthFailureMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AadAuthFailureModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AadAuthFailureModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AadAuthFailureMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AadAuthFailureModePtrOutput struct{ *pulumi.OutputState }

func (AadAuthFailureModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AadAuthFailureMode)(nil)).Elem()
}

func (o AadAuthFailureModePtrOutput) ToAadAuthFailureModePtrOutput() AadAuthFailureModePtrOutput {
	return o
}

func (o AadAuthFailureModePtrOutput) ToAadAuthFailureModePtrOutputWithContext(ctx context.Context) AadAuthFailureModePtrOutput {
	return o
}

func (o AadAuthFailureModePtrOutput) Elem() AadAuthFailureModeOutput {
	return o.ApplyT(func(v *AadAuthFailureMode) AadAuthFailureMode {
		if v != nil {
			return *v
		}
		var ret AadAuthFailureMode
		return ret
	}).(AadAuthFailureModeOutput)
}

func (o AadAuthFailureModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AadAuthFailureModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AadAuthFailureMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AadAuthFailureModeInput interface {
	pulumi.Input

	ToAadAuthFailureModeOutput() AadAuthFailureModeOutput
	ToAadAuthFailureModeOutputWithContext(context.Context) AadAuthFailureModeOutput
}

var aadAuthFailureModePtrType = reflect.TypeOf((**AadAuthFailureMode)(nil)).Elem()

type AadAuthFailureModePtrInput interface {
	pulumi.Input

	ToAadAuthFailureModePtrOutput() AadAuthFailureModePtrOutput
	ToAadAuthFailureModePtrOutputWithContext(context.Context) AadAuthFailureModePtrOutput
}

type aadAuthFailureModePtr string

func AadAuthFailureModePtr(v string) AadAuthFailureModePtrInput {
	return (*aadAuthFailureModePtr)(&v)
}

func (*aadAuthFailureModePtr) ElementType() reflect.Type {
	return aadAuthFailureModePtrType
}

func (in *aadAuthFailureModePtr) ToAadAuthFailureModePtrOutput() AadAuthFailureModePtrOutput {
	return pulumi.ToOutput(in).(AadAuthFailureModePtrOutput)
}

func (in *aadAuthFailureModePtr) ToAadAuthFailureModePtrOutputWithContext(ctx context.Context) AadAuthFailureModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AadAuthFailureModePtrOutput)
}

type HostingMode string

const (
	// The limit on number of indexes is determined by the default limits for the SKU.
	HostingModeDefault = HostingMode("default")
	// Only application for standard3 SKU, where the search service can have up to 1000 indexes.
	HostingModeHighDensity = HostingMode("highDensity")
)

func (HostingMode) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingMode)(nil)).Elem()
}

func (e HostingMode) ToHostingModeOutput() HostingModeOutput {
	return pulumi.ToOutput(e).(HostingModeOutput)
}

func (e HostingMode) ToHostingModeOutputWithContext(ctx context.Context) HostingModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostingModeOutput)
}

func (e HostingMode) ToHostingModePtrOutput() HostingModePtrOutput {
	return e.ToHostingModePtrOutputWithContext(context.Background())
}

func (e HostingMode) ToHostingModePtrOutputWithContext(ctx context.Context) HostingModePtrOutput {
	return HostingMode(e).ToHostingModeOutputWithContext(ctx).ToHostingModePtrOutputWithContext(ctx)
}

func (e HostingMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostingMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostingModeOutput struct{ *pulumi.OutputState }

func (HostingModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingMode)(nil)).Elem()
}

func (o HostingModeOutput) ToHostingModeOutput() HostingModeOutput {
	return o
}

func (o HostingModeOutput) ToHostingModeOutputWithContext(ctx context.Context) HostingModeOutput {
	return o
}

func (o HostingModeOutput) ToHostingModePtrOutput() HostingModePtrOutput {
	return o.ToHostingModePtrOutputWithContext(context.Background())
}

func (o HostingModeOutput) ToHostingModePtrOutputWithContext(ctx context.Context) HostingModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingMode) *HostingMode {
		return &v
	}).(HostingModePtrOutput)
}

func (o HostingModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostingModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostingMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostingModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostingModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostingMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostingModePtrOutput struct{ *pulumi.OutputState }

func (HostingModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingMode)(nil)).Elem()
}

func (o HostingModePtrOutput) ToHostingModePtrOutput() HostingModePtrOutput {
	return o
}

func (o HostingModePtrOutput) ToHostingModePtrOutputWithContext(ctx context.Context) HostingModePtrOutput {
	return o
}

func (o HostingModePtrOutput) Elem() HostingModeOutput {
	return o.ApplyT(func(v *HostingMode) HostingMode {
		if v != nil {
			return *v
		}
		var ret HostingMode
		return ret
	}).(HostingModeOutput)
}

func (o HostingModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostingModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostingMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostingModeInput interface {
	pulumi.Input

	ToHostingModeOutput() HostingModeOutput
	ToHostingModeOutputWithContext(context.Context) HostingModeOutput
}

var hostingModePtrType = reflect.TypeOf((**HostingMode)(nil)).Elem()

type HostingModePtrInput interface {
	pulumi.Input

	ToHostingModePtrOutput() HostingModePtrOutput
	ToHostingModePtrOutputWithContext(context.Context) HostingModePtrOutput
}

type hostingModePtr string

func HostingModePtr(v string) HostingModePtrInput {
	return (*hostingModePtr)(&v)
}

func (*hostingModePtr) ElementType() reflect.Type {
	return hostingModePtrType
}

func (in *hostingModePtr) ToHostingModePtrOutput() HostingModePtrOutput {
	return pulumi.ToOutput(in).(HostingModePtrOutput)
}

func (in *hostingModePtr) ToHostingModePtrOutputWithContext(ctx context.Context) HostingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostingModePtrOutput)
}

type IdentityType string

const (
	// Indicates that any identity associated with the search service needs to be removed.
	IdentityTypeNone = IdentityType("None")
	// Indicates that system-assigned identity for the search service will be enabled.
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
	// Indicates that one or more user assigned identities will be assigned to the search service.
	IdentityTypeUserAssigned = IdentityType("UserAssigned")
	// Indicates that system-assigned identity for the search service will be enabled along with the assignment of one or more user assigned identities.
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned, UserAssigned")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

type PrivateLinkServiceConnectionStatus string

const (
	// The private endpoint connection has been created and is pending approval.
	PrivateLinkServiceConnectionStatusPending = PrivateLinkServiceConnectionStatus("Pending")
	// The private endpoint connection is approved and is ready for use.
	PrivateLinkServiceConnectionStatusApproved = PrivateLinkServiceConnectionStatus("Approved")
	// The private endpoint connection has been rejected and cannot be used.
	PrivateLinkServiceConnectionStatusRejected = PrivateLinkServiceConnectionStatus("Rejected")
	// The private endpoint connection has been removed from the service.
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

type PublicNetworkAccess string

const (
	// The search service is accessible from traffic originating from the public internet.
	PublicNetworkAccessEnabled = PublicNetworkAccess("enabled")
	// The search service is not accessible from traffic originating from the public internet. Access is only permitted over approved private endpoint connections.
	PublicNetworkAccessDisabled = PublicNetworkAccess("disabled")
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

type SearchBypass string

const (
	// Indicates that no origin can bypass the rules defined in the 'ipRules' section. This is the default.
	SearchBypassNone = SearchBypass("None")
)

func (SearchBypass) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchBypass)(nil)).Elem()
}

func (e SearchBypass) ToSearchBypassOutput() SearchBypassOutput {
	return pulumi.ToOutput(e).(SearchBypassOutput)
}

func (e SearchBypass) ToSearchBypassOutputWithContext(ctx context.Context) SearchBypassOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SearchBypassOutput)
}

func (e SearchBypass) ToSearchBypassPtrOutput() SearchBypassPtrOutput {
	return e.ToSearchBypassPtrOutputWithContext(context.Background())
}

func (e SearchBypass) ToSearchBypassPtrOutputWithContext(ctx context.Context) SearchBypassPtrOutput {
	return SearchBypass(e).ToSearchBypassOutputWithContext(ctx).ToSearchBypassPtrOutputWithContext(ctx)
}

func (e SearchBypass) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SearchBypass) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SearchBypass) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SearchBypass) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SearchBypassOutput struct{ *pulumi.OutputState }

func (SearchBypassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchBypass)(nil)).Elem()
}

func (o SearchBypassOutput) ToSearchBypassOutput() SearchBypassOutput {
	return o
}

func (o SearchBypassOutput) ToSearchBypassOutputWithContext(ctx context.Context) SearchBypassOutput {
	return o
}

func (o SearchBypassOutput) ToSearchBypassPtrOutput() SearchBypassPtrOutput {
	return o.ToSearchBypassPtrOutputWithContext(context.Background())
}

func (o SearchBypassOutput) ToSearchBypassPtrOutputWithContext(ctx context.Context) SearchBypassPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SearchBypass) *SearchBypass {
		return &v
	}).(SearchBypassPtrOutput)
}

func (o SearchBypassOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SearchBypassOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SearchBypass) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SearchBypassOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SearchBypassOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SearchBypass) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SearchBypassPtrOutput struct{ *pulumi.OutputState }

func (SearchBypassPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchBypass)(nil)).Elem()
}

func (o SearchBypassPtrOutput) ToSearchBypassPtrOutput() SearchBypassPtrOutput {
	return o
}

func (o SearchBypassPtrOutput) ToSearchBypassPtrOutputWithContext(ctx context.Context) SearchBypassPtrOutput {
	return o
}

func (o SearchBypassPtrOutput) Elem() SearchBypassOutput {
	return o.ApplyT(func(v *SearchBypass) SearchBypass {
		if v != nil {
			return *v
		}
		var ret SearchBypass
		return ret
	}).(SearchBypassOutput)
}

func (o SearchBypassPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SearchBypassPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SearchBypass) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SearchBypassInput interface {
	pulumi.Input

	ToSearchBypassOutput() SearchBypassOutput
	ToSearchBypassOutputWithContext(context.Context) SearchBypassOutput
}

var searchBypassPtrType = reflect.TypeOf((**SearchBypass)(nil)).Elem()

type SearchBypassPtrInput interface {
	pulumi.Input

	ToSearchBypassPtrOutput() SearchBypassPtrOutput
	ToSearchBypassPtrOutputWithContext(context.Context) SearchBypassPtrOutput
}

type searchBypassPtr string

func SearchBypassPtr(v string) SearchBypassPtrInput {
	return (*searchBypassPtr)(&v)
}

func (*searchBypassPtr) ElementType() reflect.Type {
	return searchBypassPtrType
}

func (in *searchBypassPtr) ToSearchBypassPtrOutput() SearchBypassPtrOutput {
	return pulumi.ToOutput(in).(SearchBypassPtrOutput)
}

func (in *searchBypassPtr) ToSearchBypassPtrOutputWithContext(ctx context.Context) SearchBypassPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SearchBypassPtrOutput)
}

type SearchDisabledDataExfiltrationOption string

const (
	// Indicates that all data exfiltration scenarios are disabled.
	SearchDisabledDataExfiltrationOptionAll = SearchDisabledDataExfiltrationOption("All")
)

func (SearchDisabledDataExfiltrationOption) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchDisabledDataExfiltrationOption)(nil)).Elem()
}

func (e SearchDisabledDataExfiltrationOption) ToSearchDisabledDataExfiltrationOptionOutput() SearchDisabledDataExfiltrationOptionOutput {
	return pulumi.ToOutput(e).(SearchDisabledDataExfiltrationOptionOutput)
}

func (e SearchDisabledDataExfiltrationOption) ToSearchDisabledDataExfiltrationOptionOutputWithContext(ctx context.Context) SearchDisabledDataExfiltrationOptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SearchDisabledDataExfiltrationOptionOutput)
}

func (e SearchDisabledDataExfiltrationOption) ToSearchDisabledDataExfiltrationOptionPtrOutput() SearchDisabledDataExfiltrationOptionPtrOutput {
	return e.ToSearchDisabledDataExfiltrationOptionPtrOutputWithContext(context.Background())
}

func (e SearchDisabledDataExfiltrationOption) ToSearchDisabledDataExfiltrationOptionPtrOutputWithContext(ctx context.Context) SearchDisabledDataExfiltrationOptionPtrOutput {
	return SearchDisabledDataExfiltrationOption(e).ToSearchDisabledDataExfiltrationOptionOutputWithContext(ctx).ToSearchDisabledDataExfiltrationOptionPtrOutputWithContext(ctx)
}

func (e SearchDisabledDataExfiltrationOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SearchDisabledDataExfiltrationOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SearchDisabledDataExfiltrationOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SearchDisabledDataExfiltrationOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SearchDisabledDataExfiltrationOptionOutput struct{ *pulumi.OutputState }

func (SearchDisabledDataExfiltrationOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchDisabledDataExfiltrationOption)(nil)).Elem()
}

func (o SearchDisabledDataExfiltrationOptionOutput) ToSearchDisabledDataExfiltrationOptionOutput() SearchDisabledDataExfiltrationOptionOutput {
	return o
}

func (o SearchDisabledDataExfiltrationOptionOutput) ToSearchDisabledDataExfiltrationOptionOutputWithContext(ctx context.Context) SearchDisabledDataExfiltrationOptionOutput {
	return o
}

func (o SearchDisabledDataExfiltrationOptionOutput) ToSearchDisabledDataExfiltrationOptionPtrOutput() SearchDisabledDataExfiltrationOptionPtrOutput {
	return o.ToSearchDisabledDataExfiltrationOptionPtrOutputWithContext(context.Background())
}

func (o SearchDisabledDataExfiltrationOptionOutput) ToSearchDisabledDataExfiltrationOptionPtrOutputWithContext(ctx context.Context) SearchDisabledDataExfiltrationOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SearchDisabledDataExfiltrationOption) *SearchDisabledDataExfiltrationOption {
		return &v
	}).(SearchDisabledDataExfiltrationOptionPtrOutput)
}

func (o SearchDisabledDataExfiltrationOptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SearchDisabledDataExfiltrationOptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SearchDisabledDataExfiltrationOption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SearchDisabledDataExfiltrationOptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SearchDisabledDataExfiltrationOptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SearchDisabledDataExfiltrationOption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SearchDisabledDataExfiltrationOptionPtrOutput struct{ *pulumi.OutputState }

func (SearchDisabledDataExfiltrationOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchDisabledDataExfiltrationOption)(nil)).Elem()
}

func (o SearchDisabledDataExfiltrationOptionPtrOutput) ToSearchDisabledDataExfiltrationOptionPtrOutput() SearchDisabledDataExfiltrationOptionPtrOutput {
	return o
}

func (o SearchDisabledDataExfiltrationOptionPtrOutput) ToSearchDisabledDataExfiltrationOptionPtrOutputWithContext(ctx context.Context) SearchDisabledDataExfiltrationOptionPtrOutput {
	return o
}

func (o SearchDisabledDataExfiltrationOptionPtrOutput) Elem() SearchDisabledDataExfiltrationOptionOutput {
	return o.ApplyT(func(v *SearchDisabledDataExfiltrationOption) SearchDisabledDataExfiltrationOption {
		if v != nil {
			return *v
		}
		var ret SearchDisabledDataExfiltrationOption
		return ret
	}).(SearchDisabledDataExfiltrationOptionOutput)
}

func (o SearchDisabledDataExfiltrationOptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SearchDisabledDataExfiltrationOptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SearchDisabledDataExfiltrationOption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SearchDisabledDataExfiltrationOptionInput interface {
	pulumi.Input

	ToSearchDisabledDataExfiltrationOptionOutput() SearchDisabledDataExfiltrationOptionOutput
	ToSearchDisabledDataExfiltrationOptionOutputWithContext(context.Context) SearchDisabledDataExfiltrationOptionOutput
}

var searchDisabledDataExfiltrationOptionPtrType = reflect.TypeOf((**SearchDisabledDataExfiltrationOption)(nil)).Elem()

type SearchDisabledDataExfiltrationOptionPtrInput interface {
	pulumi.Input

	ToSearchDisabledDataExfiltrationOptionPtrOutput() SearchDisabledDataExfiltrationOptionPtrOutput
	ToSearchDisabledDataExfiltrationOptionPtrOutputWithContext(context.Context) SearchDisabledDataExfiltrationOptionPtrOutput
}

type searchDisabledDataExfiltrationOptionPtr string

func SearchDisabledDataExfiltrationOptionPtr(v string) SearchDisabledDataExfiltrationOptionPtrInput {
	return (*searchDisabledDataExfiltrationOptionPtr)(&v)
}

func (*searchDisabledDataExfiltrationOptionPtr) ElementType() reflect.Type {
	return searchDisabledDataExfiltrationOptionPtrType
}

func (in *searchDisabledDataExfiltrationOptionPtr) ToSearchDisabledDataExfiltrationOptionPtrOutput() SearchDisabledDataExfiltrationOptionPtrOutput {
	return pulumi.ToOutput(in).(SearchDisabledDataExfiltrationOptionPtrOutput)
}

func (in *searchDisabledDataExfiltrationOptionPtr) ToSearchDisabledDataExfiltrationOptionPtrOutputWithContext(ctx context.Context) SearchDisabledDataExfiltrationOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SearchDisabledDataExfiltrationOptionPtrOutput)
}

type SearchEncryptionWithCmk string

const (
	// No enforcement will be made and the search service can have non customer encrypted resources.
	SearchEncryptionWithCmkDisabled = SearchEncryptionWithCmk("Disabled")
	// Search service will be marked as non-compliant if there are one or more non customer encrypted resources.
	SearchEncryptionWithCmkEnabled = SearchEncryptionWithCmk("Enabled")
	// Enforcement policy is not explicitly specified, with the behavior being the same as if it were set to 'Disabled'.
	SearchEncryptionWithCmkUnspecified = SearchEncryptionWithCmk("Unspecified")
)

func (SearchEncryptionWithCmk) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchEncryptionWithCmk)(nil)).Elem()
}

func (e SearchEncryptionWithCmk) ToSearchEncryptionWithCmkOutput() SearchEncryptionWithCmkOutput {
	return pulumi.ToOutput(e).(SearchEncryptionWithCmkOutput)
}

func (e SearchEncryptionWithCmk) ToSearchEncryptionWithCmkOutputWithContext(ctx context.Context) SearchEncryptionWithCmkOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SearchEncryptionWithCmkOutput)
}

func (e SearchEncryptionWithCmk) ToSearchEncryptionWithCmkPtrOutput() SearchEncryptionWithCmkPtrOutput {
	return e.ToSearchEncryptionWithCmkPtrOutputWithContext(context.Background())
}

func (e SearchEncryptionWithCmk) ToSearchEncryptionWithCmkPtrOutputWithContext(ctx context.Context) SearchEncryptionWithCmkPtrOutput {
	return SearchEncryptionWithCmk(e).ToSearchEncryptionWithCmkOutputWithContext(ctx).ToSearchEncryptionWithCmkPtrOutputWithContext(ctx)
}

func (e SearchEncryptionWithCmk) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SearchEncryptionWithCmk) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SearchEncryptionWithCmk) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SearchEncryptionWithCmk) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SearchEncryptionWithCmkOutput struct{ *pulumi.OutputState }

func (SearchEncryptionWithCmkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchEncryptionWithCmk)(nil)).Elem()
}

func (o SearchEncryptionWithCmkOutput) ToSearchEncryptionWithCmkOutput() SearchEncryptionWithCmkOutput {
	return o
}

func (o SearchEncryptionWithCmkOutput) ToSearchEncryptionWithCmkOutputWithContext(ctx context.Context) SearchEncryptionWithCmkOutput {
	return o
}

func (o SearchEncryptionWithCmkOutput) ToSearchEncryptionWithCmkPtrOutput() SearchEncryptionWithCmkPtrOutput {
	return o.ToSearchEncryptionWithCmkPtrOutputWithContext(context.Background())
}

func (o SearchEncryptionWithCmkOutput) ToSearchEncryptionWithCmkPtrOutputWithContext(ctx context.Context) SearchEncryptionWithCmkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SearchEncryptionWithCmk) *SearchEncryptionWithCmk {
		return &v
	}).(SearchEncryptionWithCmkPtrOutput)
}

func (o SearchEncryptionWithCmkOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SearchEncryptionWithCmkOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SearchEncryptionWithCmk) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SearchEncryptionWithCmkOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SearchEncryptionWithCmkOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SearchEncryptionWithCmk) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SearchEncryptionWithCmkPtrOutput struct{ *pulumi.OutputState }

func (SearchEncryptionWithCmkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchEncryptionWithCmk)(nil)).Elem()
}

func (o SearchEncryptionWithCmkPtrOutput) ToSearchEncryptionWithCmkPtrOutput() SearchEncryptionWithCmkPtrOutput {
	return o
}

func (o SearchEncryptionWithCmkPtrOutput) ToSearchEncryptionWithCmkPtrOutputWithContext(ctx context.Context) SearchEncryptionWithCmkPtrOutput {
	return o
}

func (o SearchEncryptionWithCmkPtrOutput) Elem() SearchEncryptionWithCmkOutput {
	return o.ApplyT(func(v *SearchEncryptionWithCmk) SearchEncryptionWithCmk {
		if v != nil {
			return *v
		}
		var ret SearchEncryptionWithCmk
		return ret
	}).(SearchEncryptionWithCmkOutput)
}

func (o SearchEncryptionWithCmkPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SearchEncryptionWithCmkPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SearchEncryptionWithCmk) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SearchEncryptionWithCmkInput interface {
	pulumi.Input

	ToSearchEncryptionWithCmkOutput() SearchEncryptionWithCmkOutput
	ToSearchEncryptionWithCmkOutputWithContext(context.Context) SearchEncryptionWithCmkOutput
}

var searchEncryptionWithCmkPtrType = reflect.TypeOf((**SearchEncryptionWithCmk)(nil)).Elem()

type SearchEncryptionWithCmkPtrInput interface {
	pulumi.Input

	ToSearchEncryptionWithCmkPtrOutput() SearchEncryptionWithCmkPtrOutput
	ToSearchEncryptionWithCmkPtrOutputWithContext(context.Context) SearchEncryptionWithCmkPtrOutput
}

type searchEncryptionWithCmkPtr string

func SearchEncryptionWithCmkPtr(v string) SearchEncryptionWithCmkPtrInput {
	return (*searchEncryptionWithCmkPtr)(&v)
}

func (*searchEncryptionWithCmkPtr) ElementType() reflect.Type {
	return searchEncryptionWithCmkPtrType
}

func (in *searchEncryptionWithCmkPtr) ToSearchEncryptionWithCmkPtrOutput() SearchEncryptionWithCmkPtrOutput {
	return pulumi.ToOutput(in).(SearchEncryptionWithCmkPtrOutput)
}

func (in *searchEncryptionWithCmkPtr) ToSearchEncryptionWithCmkPtrOutputWithContext(ctx context.Context) SearchEncryptionWithCmkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SearchEncryptionWithCmkPtrOutput)
}

type SearchSemanticSearch string

const (
	// Indicates that semantic search is disabled for the search service. This is the default.
	SearchSemanticSearchDisabled = SearchSemanticSearch("disabled")
	// Enables semantic search on a search service and indicates that it is to be used within the limits of the free tier. This would cap the volume of semantic search requests and is offered at no extra charge.
	SearchSemanticSearchFree = SearchSemanticSearch("free")
	// Enables semantic search on a search service as a billable feature, with higher throughput and volume of semantic search queries.
	SearchSemanticSearchStandard = SearchSemanticSearch("standard")
)

func (SearchSemanticSearch) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchSemanticSearch)(nil)).Elem()
}

func (e SearchSemanticSearch) ToSearchSemanticSearchOutput() SearchSemanticSearchOutput {
	return pulumi.ToOutput(e).(SearchSemanticSearchOutput)
}

func (e SearchSemanticSearch) ToSearchSemanticSearchOutputWithContext(ctx context.Context) SearchSemanticSearchOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SearchSemanticSearchOutput)
}

func (e SearchSemanticSearch) ToSearchSemanticSearchPtrOutput() SearchSemanticSearchPtrOutput {
	return e.ToSearchSemanticSearchPtrOutputWithContext(context.Background())
}

func (e SearchSemanticSearch) ToSearchSemanticSearchPtrOutputWithContext(ctx context.Context) SearchSemanticSearchPtrOutput {
	return SearchSemanticSearch(e).ToSearchSemanticSearchOutputWithContext(ctx).ToSearchSemanticSearchPtrOutputWithContext(ctx)
}

func (e SearchSemanticSearch) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SearchSemanticSearch) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SearchSemanticSearch) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SearchSemanticSearch) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SearchSemanticSearchOutput struct{ *pulumi.OutputState }

func (SearchSemanticSearchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchSemanticSearch)(nil)).Elem()
}

func (o SearchSemanticSearchOutput) ToSearchSemanticSearchOutput() SearchSemanticSearchOutput {
	return o
}

func (o SearchSemanticSearchOutput) ToSearchSemanticSearchOutputWithContext(ctx context.Context) SearchSemanticSearchOutput {
	return o
}

func (o SearchSemanticSearchOutput) ToSearchSemanticSearchPtrOutput() SearchSemanticSearchPtrOutput {
	return o.ToSearchSemanticSearchPtrOutputWithContext(context.Background())
}

func (o SearchSemanticSearchOutput) ToSearchSemanticSearchPtrOutputWithContext(ctx context.Context) SearchSemanticSearchPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SearchSemanticSearch) *SearchSemanticSearch {
		return &v
	}).(SearchSemanticSearchPtrOutput)
}

func (o SearchSemanticSearchOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SearchSemanticSearchOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SearchSemanticSearch) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SearchSemanticSearchOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SearchSemanticSearchOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SearchSemanticSearch) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SearchSemanticSearchPtrOutput struct{ *pulumi.OutputState }

func (SearchSemanticSearchPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchSemanticSearch)(nil)).Elem()
}

func (o SearchSemanticSearchPtrOutput) ToSearchSemanticSearchPtrOutput() SearchSemanticSearchPtrOutput {
	return o
}

func (o SearchSemanticSearchPtrOutput) ToSearchSemanticSearchPtrOutputWithContext(ctx context.Context) SearchSemanticSearchPtrOutput {
	return o
}

func (o SearchSemanticSearchPtrOutput) Elem() SearchSemanticSearchOutput {
	return o.ApplyT(func(v *SearchSemanticSearch) SearchSemanticSearch {
		if v != nil {
			return *v
		}
		var ret SearchSemanticSearch
		return ret
	}).(SearchSemanticSearchOutput)
}

func (o SearchSemanticSearchPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SearchSemanticSearchPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SearchSemanticSearch) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SearchSemanticSearchInput interface {
	pulumi.Input

	ToSearchSemanticSearchOutput() SearchSemanticSearchOutput
	ToSearchSemanticSearchOutputWithContext(context.Context) SearchSemanticSearchOutput
}

var searchSemanticSearchPtrType = reflect.TypeOf((**SearchSemanticSearch)(nil)).Elem()

type SearchSemanticSearchPtrInput interface {
	pulumi.Input

	ToSearchSemanticSearchPtrOutput() SearchSemanticSearchPtrOutput
	ToSearchSemanticSearchPtrOutputWithContext(context.Context) SearchSemanticSearchPtrOutput
}

type searchSemanticSearchPtr string

func SearchSemanticSearchPtr(v string) SearchSemanticSearchPtrInput {
	return (*searchSemanticSearchPtr)(&v)
}

func (*searchSemanticSearchPtr) ElementType() reflect.Type {
	return searchSemanticSearchPtrType
}

func (in *searchSemanticSearchPtr) ToSearchSemanticSearchPtrOutput() SearchSemanticSearchPtrOutput {
	return pulumi.ToOutput(in).(SearchSemanticSearchPtrOutput)
}

func (in *searchSemanticSearchPtr) ToSearchSemanticSearchPtrOutputWithContext(ctx context.Context) SearchSemanticSearchPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SearchSemanticSearchPtrOutput)
}

type SharedPrivateLinkResourceProvisioningState string

const (
	// The shared private link resource is in the process of being created along with other resources for it to be fully functional.
	SharedPrivateLinkResourceProvisioningStateUpdating = SharedPrivateLinkResourceProvisioningState("Updating")
	// The shared private link resource is in the process of being deleted.
	SharedPrivateLinkResourceProvisioningStateDeleting = SharedPrivateLinkResourceProvisioningState("Deleting")
	// The shared private link resource has failed to be provisioned or deleted.
	SharedPrivateLinkResourceProvisioningStateFailed = SharedPrivateLinkResourceProvisioningState("Failed")
	// The shared private link resource has finished provisioning and is ready for approval.
	SharedPrivateLinkResourceProvisioningStateSucceeded = SharedPrivateLinkResourceProvisioningState("Succeeded")
	// Provisioning request for the shared private link resource has been accepted but the process of creation has not commenced yet.
	SharedPrivateLinkResourceProvisioningStateIncomplete = SharedPrivateLinkResourceProvisioningState("Incomplete")
)

func (SharedPrivateLinkResourceProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceProvisioningState)(nil)).Elem()
}

func (e SharedPrivateLinkResourceProvisioningState) ToSharedPrivateLinkResourceProvisioningStateOutput() SharedPrivateLinkResourceProvisioningStateOutput {
	return pulumi.ToOutput(e).(SharedPrivateLinkResourceProvisioningStateOutput)
}

func (e SharedPrivateLinkResourceProvisioningState) ToSharedPrivateLinkResourceProvisioningStateOutputWithContext(ctx context.Context) SharedPrivateLinkResourceProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SharedPrivateLinkResourceProvisioningStateOutput)
}

func (e SharedPrivateLinkResourceProvisioningState) ToSharedPrivateLinkResourceProvisioningStatePtrOutput() SharedPrivateLinkResourceProvisioningStatePtrOutput {
	return e.ToSharedPrivateLinkResourceProvisioningStatePtrOutputWithContext(context.Background())
}

func (e SharedPrivateLinkResourceProvisioningState) ToSharedPrivateLinkResourceProvisioningStatePtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourceProvisioningStatePtrOutput {
	return SharedPrivateLinkResourceProvisioningState(e).ToSharedPrivateLinkResourceProvisioningStateOutputWithContext(ctx).ToSharedPrivateLinkResourceProvisioningStatePtrOutputWithContext(ctx)
}

func (e SharedPrivateLinkResourceProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedPrivateLinkResourceProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedPrivateLinkResourceProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SharedPrivateLinkResourceProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SharedPrivateLinkResourceProvisioningStateOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceProvisioningState)(nil)).Elem()
}

func (o SharedPrivateLinkResourceProvisioningStateOutput) ToSharedPrivateLinkResourceProvisioningStateOutput() SharedPrivateLinkResourceProvisioningStateOutput {
	return o
}

func (o SharedPrivateLinkResourceProvisioningStateOutput) ToSharedPrivateLinkResourceProvisioningStateOutputWithContext(ctx context.Context) SharedPrivateLinkResourceProvisioningStateOutput {
	return o
}

func (o SharedPrivateLinkResourceProvisioningStateOutput) ToSharedPrivateLinkResourceProvisioningStatePtrOutput() SharedPrivateLinkResourceProvisioningStatePtrOutput {
	return o.ToSharedPrivateLinkResourceProvisioningStatePtrOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourceProvisioningStateOutput) ToSharedPrivateLinkResourceProvisioningStatePtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourceProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedPrivateLinkResourceProvisioningState) *SharedPrivateLinkResourceProvisioningState {
		return &v
	}).(SharedPrivateLinkResourceProvisioningStatePtrOutput)
}

func (o SharedPrivateLinkResourceProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourceProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SharedPrivateLinkResourceProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourceProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SharedPrivateLinkResourceProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourceProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPrivateLinkResourceProvisioningState)(nil)).Elem()
}

func (o SharedPrivateLinkResourceProvisioningStatePtrOutput) ToSharedPrivateLinkResourceProvisioningStatePtrOutput() SharedPrivateLinkResourceProvisioningStatePtrOutput {
	return o
}

func (o SharedPrivateLinkResourceProvisioningStatePtrOutput) ToSharedPrivateLinkResourceProvisioningStatePtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourceProvisioningStatePtrOutput {
	return o
}

func (o SharedPrivateLinkResourceProvisioningStatePtrOutput) Elem() SharedPrivateLinkResourceProvisioningStateOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceProvisioningState) SharedPrivateLinkResourceProvisioningState {
		if v != nil {
			return *v
		}
		var ret SharedPrivateLinkResourceProvisioningState
		return ret
	}).(SharedPrivateLinkResourceProvisioningStateOutput)
}

func (o SharedPrivateLinkResourceProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourceProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SharedPrivateLinkResourceProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SharedPrivateLinkResourceProvisioningStateInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceProvisioningStateOutput() SharedPrivateLinkResourceProvisioningStateOutput
	ToSharedPrivateLinkResourceProvisioningStateOutputWithContext(context.Context) SharedPrivateLinkResourceProvisioningStateOutput
}

var sharedPrivateLinkResourceProvisioningStatePtrType = reflect.TypeOf((**SharedPrivateLinkResourceProvisioningState)(nil)).Elem()

type SharedPrivateLinkResourceProvisioningStatePtrInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceProvisioningStatePtrOutput() SharedPrivateLinkResourceProvisioningStatePtrOutput
	ToSharedPrivateLinkResourceProvisioningStatePtrOutputWithContext(context.Context) SharedPrivateLinkResourceProvisioningStatePtrOutput
}

type sharedPrivateLinkResourceProvisioningStatePtr string

func SharedPrivateLinkResourceProvisioningStatePtr(v string) SharedPrivateLinkResourceProvisioningStatePtrInput {
	return (*sharedPrivateLinkResourceProvisioningStatePtr)(&v)
}

func (*sharedPrivateLinkResourceProvisioningStatePtr) ElementType() reflect.Type {
	return sharedPrivateLinkResourceProvisioningStatePtrType
}

func (in *sharedPrivateLinkResourceProvisioningStatePtr) ToSharedPrivateLinkResourceProvisioningStatePtrOutput() SharedPrivateLinkResourceProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(SharedPrivateLinkResourceProvisioningStatePtrOutput)
}

func (in *sharedPrivateLinkResourceProvisioningStatePtr) ToSharedPrivateLinkResourceProvisioningStatePtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourceProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SharedPrivateLinkResourceProvisioningStatePtrOutput)
}

type SharedPrivateLinkResourceStatus string

const (
	// The shared private link resource has been created and is pending approval.
	SharedPrivateLinkResourceStatusPending = SharedPrivateLinkResourceStatus("Pending")
	// The shared private link resource is approved and is ready for use.
	SharedPrivateLinkResourceStatusApproved = SharedPrivateLinkResourceStatus("Approved")
	// The shared private link resource has been rejected and cannot be used.
	SharedPrivateLinkResourceStatusRejected = SharedPrivateLinkResourceStatus("Rejected")
	// The shared private link resource has been removed from the service.
	SharedPrivateLinkResourceStatusDisconnected = SharedPrivateLinkResourceStatus("Disconnected")
)

func (SharedPrivateLinkResourceStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceStatus)(nil)).Elem()
}

func (e SharedPrivateLinkResourceStatus) ToSharedPrivateLinkResourceStatusOutput() SharedPrivateLinkResourceStatusOutput {
	return pulumi.ToOutput(e).(SharedPrivateLinkResourceStatusOutput)
}

func (e SharedPrivateLinkResourceStatus) ToSharedPrivateLinkResourceStatusOutputWithContext(ctx context.Context) SharedPrivateLinkResourceStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SharedPrivateLinkResourceStatusOutput)
}

func (e SharedPrivateLinkResourceStatus) ToSharedPrivateLinkResourceStatusPtrOutput() SharedPrivateLinkResourceStatusPtrOutput {
	return e.ToSharedPrivateLinkResourceStatusPtrOutputWithContext(context.Background())
}

func (e SharedPrivateLinkResourceStatus) ToSharedPrivateLinkResourceStatusPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourceStatusPtrOutput {
	return SharedPrivateLinkResourceStatus(e).ToSharedPrivateLinkResourceStatusOutputWithContext(ctx).ToSharedPrivateLinkResourceStatusPtrOutputWithContext(ctx)
}

func (e SharedPrivateLinkResourceStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedPrivateLinkResourceStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedPrivateLinkResourceStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SharedPrivateLinkResourceStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SharedPrivateLinkResourceStatusOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceStatus)(nil)).Elem()
}

func (o SharedPrivateLinkResourceStatusOutput) ToSharedPrivateLinkResourceStatusOutput() SharedPrivateLinkResourceStatusOutput {
	return o
}

func (o SharedPrivateLinkResourceStatusOutput) ToSharedPrivateLinkResourceStatusOutputWithContext(ctx context.Context) SharedPrivateLinkResourceStatusOutput {
	return o
}

func (o SharedPrivateLinkResourceStatusOutput) ToSharedPrivateLinkResourceStatusPtrOutput() SharedPrivateLinkResourceStatusPtrOutput {
	return o.ToSharedPrivateLinkResourceStatusPtrOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourceStatusOutput) ToSharedPrivateLinkResourceStatusPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourceStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedPrivateLinkResourceStatus) *SharedPrivateLinkResourceStatus {
		return &v
	}).(SharedPrivateLinkResourceStatusPtrOutput)
}

func (o SharedPrivateLinkResourceStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourceStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SharedPrivateLinkResourceStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourceStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SharedPrivateLinkResourceStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourceStatusPtrOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPrivateLinkResourceStatus)(nil)).Elem()
}

func (o SharedPrivateLinkResourceStatusPtrOutput) ToSharedPrivateLinkResourceStatusPtrOutput() SharedPrivateLinkResourceStatusPtrOutput {
	return o
}

func (o SharedPrivateLinkResourceStatusPtrOutput) ToSharedPrivateLinkResourceStatusPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourceStatusPtrOutput {
	return o
}

func (o SharedPrivateLinkResourceStatusPtrOutput) Elem() SharedPrivateLinkResourceStatusOutput {
	return o.ApplyT(func(v *SharedPrivateLinkResourceStatus) SharedPrivateLinkResourceStatus {
		if v != nil {
			return *v
		}
		var ret SharedPrivateLinkResourceStatus
		return ret
	}).(SharedPrivateLinkResourceStatusOutput)
}

func (o SharedPrivateLinkResourceStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SharedPrivateLinkResourceStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SharedPrivateLinkResourceStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SharedPrivateLinkResourceStatusInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceStatusOutput() SharedPrivateLinkResourceStatusOutput
	ToSharedPrivateLinkResourceStatusOutputWithContext(context.Context) SharedPrivateLinkResourceStatusOutput
}

var sharedPrivateLinkResourceStatusPtrType = reflect.TypeOf((**SharedPrivateLinkResourceStatus)(nil)).Elem()

type SharedPrivateLinkResourceStatusPtrInput interface {
	pulumi.Input

	ToSharedPrivateLinkResourceStatusPtrOutput() SharedPrivateLinkResourceStatusPtrOutput
	ToSharedPrivateLinkResourceStatusPtrOutputWithContext(context.Context) SharedPrivateLinkResourceStatusPtrOutput
}

type sharedPrivateLinkResourceStatusPtr string

func SharedPrivateLinkResourceStatusPtr(v string) SharedPrivateLinkResourceStatusPtrInput {
	return (*sharedPrivateLinkResourceStatusPtr)(&v)
}

func (*sharedPrivateLinkResourceStatusPtr) ElementType() reflect.Type {
	return sharedPrivateLinkResourceStatusPtrType
}

func (in *sharedPrivateLinkResourceStatusPtr) ToSharedPrivateLinkResourceStatusPtrOutput() SharedPrivateLinkResourceStatusPtrOutput {
	return pulumi.ToOutput(in).(SharedPrivateLinkResourceStatusPtrOutput)
}

func (in *sharedPrivateLinkResourceStatusPtr) ToSharedPrivateLinkResourceStatusPtrOutputWithContext(ctx context.Context) SharedPrivateLinkResourceStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SharedPrivateLinkResourceStatusPtrOutput)
}

type SkuName string

const (
	// Free tier, with no SLA guarantees and a subset of features offered to paid tiers.
	SkuNameFree = SkuName("free")
	// Paid tier dedicated service with up to 3 replicas.
	SkuNameBasic = SkuName("basic")
	// Paid tier dedicated service with up to 12 partitions and 12 replicas.
	SkuNameStandard = SkuName("standard")
	// Similar to 'standard', but with more capacity per search unit.
	SkuNameStandard2 = SkuName("standard2")
	//  The largest Standard offering with up to 12 partitions and 12 replicas (or up to 3 partitions with more indexes if you also set the hostingMode property to 'highDensity').
	SkuNameStandard3 = SkuName("standard3")
	// Paid tier dedicated service that supports 1TB per partition, up to 12 partitions.
	SkuNameStorageOptimizedL1 = SkuName("storage_optimized_l1")
	// Paid tier dedicated service that supports 2TB per partition, up to 12 partitions.
	SkuNameStorageOptimizedL2 = SkuName("storage_optimized_l2")
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
	pulumi.RegisterOutputType(AadAuthFailureModeOutput{})
	pulumi.RegisterOutputType(AadAuthFailureModePtrOutput{})
	pulumi.RegisterOutputType(HostingModeOutput{})
	pulumi.RegisterOutputType(HostingModePtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(SearchBypassOutput{})
	pulumi.RegisterOutputType(SearchBypassPtrOutput{})
	pulumi.RegisterOutputType(SearchDisabledDataExfiltrationOptionOutput{})
	pulumi.RegisterOutputType(SearchDisabledDataExfiltrationOptionPtrOutput{})
	pulumi.RegisterOutputType(SearchEncryptionWithCmkOutput{})
	pulumi.RegisterOutputType(SearchEncryptionWithCmkPtrOutput{})
	pulumi.RegisterOutputType(SearchSemanticSearchOutput{})
	pulumi.RegisterOutputType(SearchSemanticSearchPtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceProvisioningStateOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceStatusOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceStatusPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
