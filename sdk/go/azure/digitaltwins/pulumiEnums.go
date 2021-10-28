


package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthenticationType string

const (
	AuthenticationTypeKeyBased      = AuthenticationType("KeyBased")
	AuthenticationTypeIdentityBased = AuthenticationType("IdentityBased")
)

func (AuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationType)(nil)).Elem()
}

func (e AuthenticationType) ToAuthenticationTypeOutput() AuthenticationTypeOutput {
	return pulumi.ToOutput(e).(AuthenticationTypeOutput)
}

func (e AuthenticationType) ToAuthenticationTypeOutputWithContext(ctx context.Context) AuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthenticationTypeOutput)
}

func (e AuthenticationType) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return e.ToAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e AuthenticationType) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return AuthenticationType(e).ToAuthenticationTypeOutputWithContext(ctx).ToAuthenticationTypePtrOutputWithContext(ctx)
}

func (e AuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthenticationTypeOutput struct{ *pulumi.OutputState }

func (AuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationType)(nil)).Elem()
}

func (o AuthenticationTypeOutput) ToAuthenticationTypeOutput() AuthenticationTypeOutput {
	return o
}

func (o AuthenticationTypeOutput) ToAuthenticationTypeOutputWithContext(ctx context.Context) AuthenticationTypeOutput {
	return o
}

func (o AuthenticationTypeOutput) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return o.ToAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationType) *AuthenticationType {
		return &v
	}).(AuthenticationTypePtrOutput)
}

func (o AuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (AuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationType)(nil)).Elem()
}

func (o AuthenticationTypePtrOutput) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return o
}

func (o AuthenticationTypePtrOutput) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return o
}

func (o AuthenticationTypePtrOutput) Elem() AuthenticationTypeOutput {
	return o.ApplyT(func(v *AuthenticationType) AuthenticationType {
		if v != nil {
			return *v
		}
		var ret AuthenticationType
		return ret
	}).(AuthenticationTypeOutput)
}

func (o AuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AuthenticationTypeInput interface {
	pulumi.Input

	ToAuthenticationTypeOutput() AuthenticationTypeOutput
	ToAuthenticationTypeOutputWithContext(context.Context) AuthenticationTypeOutput
}

var authenticationTypePtrType = reflect.TypeOf((**AuthenticationType)(nil)).Elem()

type AuthenticationTypePtrInput interface {
	pulumi.Input

	ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput
	ToAuthenticationTypePtrOutputWithContext(context.Context) AuthenticationTypePtrOutput
}

type authenticationTypePtr string

func AuthenticationTypePtr(v string) AuthenticationTypePtrInput {
	return (*authenticationTypePtr)(&v)
}

func (*authenticationTypePtr) ElementType() reflect.Type {
	return authenticationTypePtrType
}

func (in *authenticationTypePtr) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(AuthenticationTypePtrOutput)
}

func (in *authenticationTypePtr) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthenticationTypePtrOutput)
}

type DigitalTwinsIdentityType string

const (
	DigitalTwinsIdentityTypeNone           = DigitalTwinsIdentityType("None")
	DigitalTwinsIdentityTypeSystemAssigned = DigitalTwinsIdentityType("SystemAssigned")
)

func (DigitalTwinsIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsIdentityType)(nil)).Elem()
}

func (e DigitalTwinsIdentityType) ToDigitalTwinsIdentityTypeOutput() DigitalTwinsIdentityTypeOutput {
	return pulumi.ToOutput(e).(DigitalTwinsIdentityTypeOutput)
}

func (e DigitalTwinsIdentityType) ToDigitalTwinsIdentityTypeOutputWithContext(ctx context.Context) DigitalTwinsIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DigitalTwinsIdentityTypeOutput)
}

func (e DigitalTwinsIdentityType) ToDigitalTwinsIdentityTypePtrOutput() DigitalTwinsIdentityTypePtrOutput {
	return e.ToDigitalTwinsIdentityTypePtrOutputWithContext(context.Background())
}

func (e DigitalTwinsIdentityType) ToDigitalTwinsIdentityTypePtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityTypePtrOutput {
	return DigitalTwinsIdentityType(e).ToDigitalTwinsIdentityTypeOutputWithContext(ctx).ToDigitalTwinsIdentityTypePtrOutputWithContext(ctx)
}

func (e DigitalTwinsIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DigitalTwinsIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DigitalTwinsIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DigitalTwinsIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DigitalTwinsIdentityTypeOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsIdentityType)(nil)).Elem()
}

func (o DigitalTwinsIdentityTypeOutput) ToDigitalTwinsIdentityTypeOutput() DigitalTwinsIdentityTypeOutput {
	return o
}

func (o DigitalTwinsIdentityTypeOutput) ToDigitalTwinsIdentityTypeOutputWithContext(ctx context.Context) DigitalTwinsIdentityTypeOutput {
	return o
}

func (o DigitalTwinsIdentityTypeOutput) ToDigitalTwinsIdentityTypePtrOutput() DigitalTwinsIdentityTypePtrOutput {
	return o.ToDigitalTwinsIdentityTypePtrOutputWithContext(context.Background())
}

func (o DigitalTwinsIdentityTypeOutput) ToDigitalTwinsIdentityTypePtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DigitalTwinsIdentityType) *DigitalTwinsIdentityType {
		return &v
	}).(DigitalTwinsIdentityTypePtrOutput)
}

func (o DigitalTwinsIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DigitalTwinsIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DigitalTwinsIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DigitalTwinsIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DigitalTwinsIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DigitalTwinsIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DigitalTwinsIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsIdentityType)(nil)).Elem()
}

func (o DigitalTwinsIdentityTypePtrOutput) ToDigitalTwinsIdentityTypePtrOutput() DigitalTwinsIdentityTypePtrOutput {
	return o
}

func (o DigitalTwinsIdentityTypePtrOutput) ToDigitalTwinsIdentityTypePtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityTypePtrOutput {
	return o
}

func (o DigitalTwinsIdentityTypePtrOutput) Elem() DigitalTwinsIdentityTypeOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityType) DigitalTwinsIdentityType {
		if v != nil {
			return *v
		}
		var ret DigitalTwinsIdentityType
		return ret
	}).(DigitalTwinsIdentityTypeOutput)
}

func (o DigitalTwinsIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DigitalTwinsIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DigitalTwinsIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DigitalTwinsIdentityTypeInput interface {
	pulumi.Input

	ToDigitalTwinsIdentityTypeOutput() DigitalTwinsIdentityTypeOutput
	ToDigitalTwinsIdentityTypeOutputWithContext(context.Context) DigitalTwinsIdentityTypeOutput
}

var digitalTwinsIdentityTypePtrType = reflect.TypeOf((**DigitalTwinsIdentityType)(nil)).Elem()

type DigitalTwinsIdentityTypePtrInput interface {
	pulumi.Input

	ToDigitalTwinsIdentityTypePtrOutput() DigitalTwinsIdentityTypePtrOutput
	ToDigitalTwinsIdentityTypePtrOutputWithContext(context.Context) DigitalTwinsIdentityTypePtrOutput
}

type digitalTwinsIdentityTypePtr string

func DigitalTwinsIdentityTypePtr(v string) DigitalTwinsIdentityTypePtrInput {
	return (*digitalTwinsIdentityTypePtr)(&v)
}

func (*digitalTwinsIdentityTypePtr) ElementType() reflect.Type {
	return digitalTwinsIdentityTypePtrType
}

func (in *digitalTwinsIdentityTypePtr) ToDigitalTwinsIdentityTypePtrOutput() DigitalTwinsIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(DigitalTwinsIdentityTypePtrOutput)
}

func (in *digitalTwinsIdentityTypePtr) ToDigitalTwinsIdentityTypePtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DigitalTwinsIdentityTypePtrOutput)
}

type EndpointType string

const (
	EndpointTypeEventHub   = EndpointType("EventHub")
	EndpointTypeEventGrid  = EndpointType("EventGrid")
	EndpointTypeServiceBus = EndpointType("ServiceBus")
)

func (EndpointType) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (e EndpointType) ToEndpointTypeOutput() EndpointTypeOutput {
	return pulumi.ToOutput(e).(EndpointTypeOutput)
}

func (e EndpointType) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointTypeOutput)
}

func (e EndpointType) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return e.ToEndpointTypePtrOutputWithContext(context.Background())
}

func (e EndpointType) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return EndpointType(e).ToEndpointTypeOutputWithContext(ctx).ToEndpointTypePtrOutputWithContext(ctx)
}

func (e EndpointType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointTypeOutput struct{ *pulumi.OutputState }

func (EndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (o EndpointTypeOutput) ToEndpointTypeOutput() EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return o.ToEndpointTypePtrOutputWithContext(context.Background())
}

func (o EndpointTypeOutput) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointType) *EndpointType {
		return &v
	}).(EndpointTypePtrOutput)
}

func (o EndpointTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointTypePtrOutput struct{ *pulumi.OutputState }

func (EndpointTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointType)(nil)).Elem()
}

func (o EndpointTypePtrOutput) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return o
}

func (o EndpointTypePtrOutput) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return o
}

func (o EndpointTypePtrOutput) Elem() EndpointTypeOutput {
	return o.ApplyT(func(v *EndpointType) EndpointType {
		if v != nil {
			return *v
		}
		var ret EndpointType
		return ret
	}).(EndpointTypeOutput)
}

func (o EndpointTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EndpointTypeInput interface {
	pulumi.Input

	ToEndpointTypeOutput() EndpointTypeOutput
	ToEndpointTypeOutputWithContext(context.Context) EndpointTypeOutput
}

var endpointTypePtrType = reflect.TypeOf((**EndpointType)(nil)).Elem()

type EndpointTypePtrInput interface {
	pulumi.Input

	ToEndpointTypePtrOutput() EndpointTypePtrOutput
	ToEndpointTypePtrOutputWithContext(context.Context) EndpointTypePtrOutput
}

type endpointTypePtr string

func EndpointTypePtr(v string) EndpointTypePtrInput {
	return (*endpointTypePtr)(&v)
}

func (*endpointTypePtr) ElementType() reflect.Type {
	return endpointTypePtrType
}

func (in *endpointTypePtr) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return pulumi.ToOutput(in).(EndpointTypePtrOutput)
}

func (in *endpointTypePtr) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointTypePtrOutput)
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationTypeInput)(nil)).Elem(), AuthenticationType("KeyBased"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationTypePtrInput)(nil)).Elem(), AuthenticationType("KeyBased"))
	pulumi.RegisterInputType(reflect.TypeOf((*DigitalTwinsIdentityTypeInput)(nil)).Elem(), DigitalTwinsIdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*DigitalTwinsIdentityTypePtrInput)(nil)).Elem(), DigitalTwinsIdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointTypeInput)(nil)).Elem(), EndpointType("EventHub"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointTypePtrInput)(nil)).Elem(), EndpointType("EventHub"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessPtrInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterOutputType(AuthenticationTypeOutput{})
	pulumi.RegisterOutputType(AuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityTypeOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
}
