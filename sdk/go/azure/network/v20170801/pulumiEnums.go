


package v20170801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Access string

const (
	AccessAllow = Access("Allow")
	AccessDeny  = Access("Deny")
)

func (Access) ElementType() reflect.Type {
	return reflect.TypeOf((*Access)(nil)).Elem()
}

func (e Access) ToAccessOutput() AccessOutput {
	return pulumi.ToOutput(e).(AccessOutput)
}

func (e Access) ToAccessOutputWithContext(ctx context.Context) AccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessOutput)
}

func (e Access) ToAccessPtrOutput() AccessPtrOutput {
	return e.ToAccessPtrOutputWithContext(context.Background())
}

func (e Access) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return Access(e).ToAccessOutputWithContext(ctx).ToAccessPtrOutputWithContext(ctx)
}

func (e Access) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Access) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Access) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Access) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessOutput struct{ *pulumi.OutputState }

func (AccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Access)(nil)).Elem()
}

func (o AccessOutput) ToAccessOutput() AccessOutput {
	return o
}

func (o AccessOutput) ToAccessOutputWithContext(ctx context.Context) AccessOutput {
	return o
}

func (o AccessOutput) ToAccessPtrOutput() AccessPtrOutput {
	return o.ToAccessPtrOutputWithContext(context.Background())
}

func (o AccessOutput) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Access) *Access {
		return &v
	}).(AccessPtrOutput)
}

func (o AccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Access) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Access) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessPtrOutput struct{ *pulumi.OutputState }

func (AccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Access)(nil)).Elem()
}

func (o AccessPtrOutput) ToAccessPtrOutput() AccessPtrOutput {
	return o
}

func (o AccessPtrOutput) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return o
}

func (o AccessPtrOutput) Elem() AccessOutput {
	return o.ApplyT(func(v *Access) Access {
		if v != nil {
			return *v
		}
		var ret Access
		return ret
	}).(AccessOutput)
}

func (o AccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Access) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessInput interface {
	pulumi.Input

	ToAccessOutput() AccessOutput
	ToAccessOutputWithContext(context.Context) AccessOutput
}

var accessPtrType = reflect.TypeOf((**Access)(nil)).Elem()

type AccessPtrInput interface {
	pulumi.Input

	ToAccessPtrOutput() AccessPtrOutput
	ToAccessPtrOutputWithContext(context.Context) AccessPtrOutput
}

type accessPtr string

func AccessPtr(v string) AccessPtrInput {
	return (*accessPtr)(&v)
}

func (*accessPtr) ElementType() reflect.Type {
	return accessPtrType
}

func (in *accessPtr) ToAccessPtrOutput() AccessPtrOutput {
	return pulumi.ToOutput(in).(AccessPtrOutput)
}

func (in *accessPtr) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessPtrOutput)
}

type ApplicationGatewayCookieBasedAffinity string

const (
	ApplicationGatewayCookieBasedAffinityEnabled  = ApplicationGatewayCookieBasedAffinity("Enabled")
	ApplicationGatewayCookieBasedAffinityDisabled = ApplicationGatewayCookieBasedAffinity("Disabled")
)

func (ApplicationGatewayCookieBasedAffinity) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayCookieBasedAffinity)(nil)).Elem()
}

func (e ApplicationGatewayCookieBasedAffinity) ToApplicationGatewayCookieBasedAffinityOutput() ApplicationGatewayCookieBasedAffinityOutput {
	return pulumi.ToOutput(e).(ApplicationGatewayCookieBasedAffinityOutput)
}

func (e ApplicationGatewayCookieBasedAffinity) ToApplicationGatewayCookieBasedAffinityOutputWithContext(ctx context.Context) ApplicationGatewayCookieBasedAffinityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewayCookieBasedAffinityOutput)
}

func (e ApplicationGatewayCookieBasedAffinity) ToApplicationGatewayCookieBasedAffinityPtrOutput() ApplicationGatewayCookieBasedAffinityPtrOutput {
	return e.ToApplicationGatewayCookieBasedAffinityPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayCookieBasedAffinity) ToApplicationGatewayCookieBasedAffinityPtrOutputWithContext(ctx context.Context) ApplicationGatewayCookieBasedAffinityPtrOutput {
	return ApplicationGatewayCookieBasedAffinity(e).ToApplicationGatewayCookieBasedAffinityOutputWithContext(ctx).ToApplicationGatewayCookieBasedAffinityPtrOutputWithContext(ctx)
}

func (e ApplicationGatewayCookieBasedAffinity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayCookieBasedAffinity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayCookieBasedAffinity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayCookieBasedAffinity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewayCookieBasedAffinityOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayCookieBasedAffinityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayCookieBasedAffinity)(nil)).Elem()
}

func (o ApplicationGatewayCookieBasedAffinityOutput) ToApplicationGatewayCookieBasedAffinityOutput() ApplicationGatewayCookieBasedAffinityOutput {
	return o
}

func (o ApplicationGatewayCookieBasedAffinityOutput) ToApplicationGatewayCookieBasedAffinityOutputWithContext(ctx context.Context) ApplicationGatewayCookieBasedAffinityOutput {
	return o
}

func (o ApplicationGatewayCookieBasedAffinityOutput) ToApplicationGatewayCookieBasedAffinityPtrOutput() ApplicationGatewayCookieBasedAffinityPtrOutput {
	return o.ToApplicationGatewayCookieBasedAffinityPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayCookieBasedAffinityOutput) ToApplicationGatewayCookieBasedAffinityPtrOutputWithContext(ctx context.Context) ApplicationGatewayCookieBasedAffinityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayCookieBasedAffinity) *ApplicationGatewayCookieBasedAffinity {
		return &v
	}).(ApplicationGatewayCookieBasedAffinityPtrOutput)
}

func (o ApplicationGatewayCookieBasedAffinityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewayCookieBasedAffinityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayCookieBasedAffinity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewayCookieBasedAffinityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayCookieBasedAffinityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayCookieBasedAffinity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayCookieBasedAffinityPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayCookieBasedAffinityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayCookieBasedAffinity)(nil)).Elem()
}

func (o ApplicationGatewayCookieBasedAffinityPtrOutput) ToApplicationGatewayCookieBasedAffinityPtrOutput() ApplicationGatewayCookieBasedAffinityPtrOutput {
	return o
}

func (o ApplicationGatewayCookieBasedAffinityPtrOutput) ToApplicationGatewayCookieBasedAffinityPtrOutputWithContext(ctx context.Context) ApplicationGatewayCookieBasedAffinityPtrOutput {
	return o
}

func (o ApplicationGatewayCookieBasedAffinityPtrOutput) Elem() ApplicationGatewayCookieBasedAffinityOutput {
	return o.ApplyT(func(v *ApplicationGatewayCookieBasedAffinity) ApplicationGatewayCookieBasedAffinity {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayCookieBasedAffinity
		return ret
	}).(ApplicationGatewayCookieBasedAffinityOutput)
}

func (o ApplicationGatewayCookieBasedAffinityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayCookieBasedAffinityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewayCookieBasedAffinity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewayCookieBasedAffinityInput interface {
	pulumi.Input

	ToApplicationGatewayCookieBasedAffinityOutput() ApplicationGatewayCookieBasedAffinityOutput
	ToApplicationGatewayCookieBasedAffinityOutputWithContext(context.Context) ApplicationGatewayCookieBasedAffinityOutput
}

var applicationGatewayCookieBasedAffinityPtrType = reflect.TypeOf((**ApplicationGatewayCookieBasedAffinity)(nil)).Elem()

type ApplicationGatewayCookieBasedAffinityPtrInput interface {
	pulumi.Input

	ToApplicationGatewayCookieBasedAffinityPtrOutput() ApplicationGatewayCookieBasedAffinityPtrOutput
	ToApplicationGatewayCookieBasedAffinityPtrOutputWithContext(context.Context) ApplicationGatewayCookieBasedAffinityPtrOutput
}

type applicationGatewayCookieBasedAffinityPtr string

func ApplicationGatewayCookieBasedAffinityPtr(v string) ApplicationGatewayCookieBasedAffinityPtrInput {
	return (*applicationGatewayCookieBasedAffinityPtr)(&v)
}

func (*applicationGatewayCookieBasedAffinityPtr) ElementType() reflect.Type {
	return applicationGatewayCookieBasedAffinityPtrType
}

func (in *applicationGatewayCookieBasedAffinityPtr) ToApplicationGatewayCookieBasedAffinityPtrOutput() ApplicationGatewayCookieBasedAffinityPtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewayCookieBasedAffinityPtrOutput)
}

func (in *applicationGatewayCookieBasedAffinityPtr) ToApplicationGatewayCookieBasedAffinityPtrOutputWithContext(ctx context.Context) ApplicationGatewayCookieBasedAffinityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewayCookieBasedAffinityPtrOutput)
}

type ApplicationGatewayFirewallMode string

const (
	ApplicationGatewayFirewallModeDetection  = ApplicationGatewayFirewallMode("Detection")
	ApplicationGatewayFirewallModePrevention = ApplicationGatewayFirewallMode("Prevention")
)

func (ApplicationGatewayFirewallMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFirewallMode)(nil)).Elem()
}

func (e ApplicationGatewayFirewallMode) ToApplicationGatewayFirewallModeOutput() ApplicationGatewayFirewallModeOutput {
	return pulumi.ToOutput(e).(ApplicationGatewayFirewallModeOutput)
}

func (e ApplicationGatewayFirewallMode) ToApplicationGatewayFirewallModeOutputWithContext(ctx context.Context) ApplicationGatewayFirewallModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewayFirewallModeOutput)
}

func (e ApplicationGatewayFirewallMode) ToApplicationGatewayFirewallModePtrOutput() ApplicationGatewayFirewallModePtrOutput {
	return e.ToApplicationGatewayFirewallModePtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayFirewallMode) ToApplicationGatewayFirewallModePtrOutputWithContext(ctx context.Context) ApplicationGatewayFirewallModePtrOutput {
	return ApplicationGatewayFirewallMode(e).ToApplicationGatewayFirewallModeOutputWithContext(ctx).ToApplicationGatewayFirewallModePtrOutputWithContext(ctx)
}

func (e ApplicationGatewayFirewallMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayFirewallMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayFirewallMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayFirewallMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewayFirewallModeOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFirewallMode)(nil)).Elem()
}

func (o ApplicationGatewayFirewallModeOutput) ToApplicationGatewayFirewallModeOutput() ApplicationGatewayFirewallModeOutput {
	return o
}

func (o ApplicationGatewayFirewallModeOutput) ToApplicationGatewayFirewallModeOutputWithContext(ctx context.Context) ApplicationGatewayFirewallModeOutput {
	return o
}

func (o ApplicationGatewayFirewallModeOutput) ToApplicationGatewayFirewallModePtrOutput() ApplicationGatewayFirewallModePtrOutput {
	return o.ToApplicationGatewayFirewallModePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayFirewallModeOutput) ToApplicationGatewayFirewallModePtrOutputWithContext(ctx context.Context) ApplicationGatewayFirewallModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayFirewallMode) *ApplicationGatewayFirewallMode {
		return &v
	}).(ApplicationGatewayFirewallModePtrOutput)
}

func (o ApplicationGatewayFirewallModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewayFirewallModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayFirewallMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewayFirewallModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayFirewallModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayFirewallMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayFirewallModePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayFirewallMode)(nil)).Elem()
}

func (o ApplicationGatewayFirewallModePtrOutput) ToApplicationGatewayFirewallModePtrOutput() ApplicationGatewayFirewallModePtrOutput {
	return o
}

func (o ApplicationGatewayFirewallModePtrOutput) ToApplicationGatewayFirewallModePtrOutputWithContext(ctx context.Context) ApplicationGatewayFirewallModePtrOutput {
	return o
}

func (o ApplicationGatewayFirewallModePtrOutput) Elem() ApplicationGatewayFirewallModeOutput {
	return o.ApplyT(func(v *ApplicationGatewayFirewallMode) ApplicationGatewayFirewallMode {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayFirewallMode
		return ret
	}).(ApplicationGatewayFirewallModeOutput)
}

func (o ApplicationGatewayFirewallModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayFirewallModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewayFirewallMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewayFirewallModeInput interface {
	pulumi.Input

	ToApplicationGatewayFirewallModeOutput() ApplicationGatewayFirewallModeOutput
	ToApplicationGatewayFirewallModeOutputWithContext(context.Context) ApplicationGatewayFirewallModeOutput
}

var applicationGatewayFirewallModePtrType = reflect.TypeOf((**ApplicationGatewayFirewallMode)(nil)).Elem()

type ApplicationGatewayFirewallModePtrInput interface {
	pulumi.Input

	ToApplicationGatewayFirewallModePtrOutput() ApplicationGatewayFirewallModePtrOutput
	ToApplicationGatewayFirewallModePtrOutputWithContext(context.Context) ApplicationGatewayFirewallModePtrOutput
}

type applicationGatewayFirewallModePtr string

func ApplicationGatewayFirewallModePtr(v string) ApplicationGatewayFirewallModePtrInput {
	return (*applicationGatewayFirewallModePtr)(&v)
}

func (*applicationGatewayFirewallModePtr) ElementType() reflect.Type {
	return applicationGatewayFirewallModePtrType
}

func (in *applicationGatewayFirewallModePtr) ToApplicationGatewayFirewallModePtrOutput() ApplicationGatewayFirewallModePtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewayFirewallModePtrOutput)
}

func (in *applicationGatewayFirewallModePtr) ToApplicationGatewayFirewallModePtrOutputWithContext(ctx context.Context) ApplicationGatewayFirewallModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewayFirewallModePtrOutput)
}

type ApplicationGatewayProtocol string

const (
	ApplicationGatewayProtocolHttp  = ApplicationGatewayProtocol("Http")
	ApplicationGatewayProtocolHttps = ApplicationGatewayProtocol("Https")
)

func (ApplicationGatewayProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProtocol)(nil)).Elem()
}

func (e ApplicationGatewayProtocol) ToApplicationGatewayProtocolOutput() ApplicationGatewayProtocolOutput {
	return pulumi.ToOutput(e).(ApplicationGatewayProtocolOutput)
}

func (e ApplicationGatewayProtocol) ToApplicationGatewayProtocolOutputWithContext(ctx context.Context) ApplicationGatewayProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewayProtocolOutput)
}

func (e ApplicationGatewayProtocol) ToApplicationGatewayProtocolPtrOutput() ApplicationGatewayProtocolPtrOutput {
	return e.ToApplicationGatewayProtocolPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayProtocol) ToApplicationGatewayProtocolPtrOutputWithContext(ctx context.Context) ApplicationGatewayProtocolPtrOutput {
	return ApplicationGatewayProtocol(e).ToApplicationGatewayProtocolOutputWithContext(ctx).ToApplicationGatewayProtocolPtrOutputWithContext(ctx)
}

func (e ApplicationGatewayProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewayProtocolOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProtocol)(nil)).Elem()
}

func (o ApplicationGatewayProtocolOutput) ToApplicationGatewayProtocolOutput() ApplicationGatewayProtocolOutput {
	return o
}

func (o ApplicationGatewayProtocolOutput) ToApplicationGatewayProtocolOutputWithContext(ctx context.Context) ApplicationGatewayProtocolOutput {
	return o
}

func (o ApplicationGatewayProtocolOutput) ToApplicationGatewayProtocolPtrOutput() ApplicationGatewayProtocolPtrOutput {
	return o.ToApplicationGatewayProtocolPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayProtocolOutput) ToApplicationGatewayProtocolPtrOutputWithContext(ctx context.Context) ApplicationGatewayProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayProtocol) *ApplicationGatewayProtocol {
		return &v
	}).(ApplicationGatewayProtocolPtrOutput)
}

func (o ApplicationGatewayProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewayProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewayProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayProtocolPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayProtocol)(nil)).Elem()
}

func (o ApplicationGatewayProtocolPtrOutput) ToApplicationGatewayProtocolPtrOutput() ApplicationGatewayProtocolPtrOutput {
	return o
}

func (o ApplicationGatewayProtocolPtrOutput) ToApplicationGatewayProtocolPtrOutputWithContext(ctx context.Context) ApplicationGatewayProtocolPtrOutput {
	return o
}

func (o ApplicationGatewayProtocolPtrOutput) Elem() ApplicationGatewayProtocolOutput {
	return o.ApplyT(func(v *ApplicationGatewayProtocol) ApplicationGatewayProtocol {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayProtocol
		return ret
	}).(ApplicationGatewayProtocolOutput)
}

func (o ApplicationGatewayProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewayProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewayProtocolInput interface {
	pulumi.Input

	ToApplicationGatewayProtocolOutput() ApplicationGatewayProtocolOutput
	ToApplicationGatewayProtocolOutputWithContext(context.Context) ApplicationGatewayProtocolOutput
}

var applicationGatewayProtocolPtrType = reflect.TypeOf((**ApplicationGatewayProtocol)(nil)).Elem()

type ApplicationGatewayProtocolPtrInput interface {
	pulumi.Input

	ToApplicationGatewayProtocolPtrOutput() ApplicationGatewayProtocolPtrOutput
	ToApplicationGatewayProtocolPtrOutputWithContext(context.Context) ApplicationGatewayProtocolPtrOutput
}

type applicationGatewayProtocolPtr string

func ApplicationGatewayProtocolPtr(v string) ApplicationGatewayProtocolPtrInput {
	return (*applicationGatewayProtocolPtr)(&v)
}

func (*applicationGatewayProtocolPtr) ElementType() reflect.Type {
	return applicationGatewayProtocolPtrType
}

func (in *applicationGatewayProtocolPtr) ToApplicationGatewayProtocolPtrOutput() ApplicationGatewayProtocolPtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewayProtocolPtrOutput)
}

func (in *applicationGatewayProtocolPtr) ToApplicationGatewayProtocolPtrOutputWithContext(ctx context.Context) ApplicationGatewayProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewayProtocolPtrOutput)
}

type ApplicationGatewayRedirectType string

const (
	ApplicationGatewayRedirectTypePermanent = ApplicationGatewayRedirectType("Permanent")
	ApplicationGatewayRedirectTypeFound     = ApplicationGatewayRedirectType("Found")
	ApplicationGatewayRedirectTypeSeeOther  = ApplicationGatewayRedirectType("SeeOther")
	ApplicationGatewayRedirectTypeTemporary = ApplicationGatewayRedirectType("Temporary")
)

func (ApplicationGatewayRedirectType) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRedirectType)(nil)).Elem()
}

func (e ApplicationGatewayRedirectType) ToApplicationGatewayRedirectTypeOutput() ApplicationGatewayRedirectTypeOutput {
	return pulumi.ToOutput(e).(ApplicationGatewayRedirectTypeOutput)
}

func (e ApplicationGatewayRedirectType) ToApplicationGatewayRedirectTypeOutputWithContext(ctx context.Context) ApplicationGatewayRedirectTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewayRedirectTypeOutput)
}

func (e ApplicationGatewayRedirectType) ToApplicationGatewayRedirectTypePtrOutput() ApplicationGatewayRedirectTypePtrOutput {
	return e.ToApplicationGatewayRedirectTypePtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayRedirectType) ToApplicationGatewayRedirectTypePtrOutputWithContext(ctx context.Context) ApplicationGatewayRedirectTypePtrOutput {
	return ApplicationGatewayRedirectType(e).ToApplicationGatewayRedirectTypeOutputWithContext(ctx).ToApplicationGatewayRedirectTypePtrOutputWithContext(ctx)
}

func (e ApplicationGatewayRedirectType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayRedirectType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayRedirectType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayRedirectType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewayRedirectTypeOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRedirectTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRedirectType)(nil)).Elem()
}

func (o ApplicationGatewayRedirectTypeOutput) ToApplicationGatewayRedirectTypeOutput() ApplicationGatewayRedirectTypeOutput {
	return o
}

func (o ApplicationGatewayRedirectTypeOutput) ToApplicationGatewayRedirectTypeOutputWithContext(ctx context.Context) ApplicationGatewayRedirectTypeOutput {
	return o
}

func (o ApplicationGatewayRedirectTypeOutput) ToApplicationGatewayRedirectTypePtrOutput() ApplicationGatewayRedirectTypePtrOutput {
	return o.ToApplicationGatewayRedirectTypePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayRedirectTypeOutput) ToApplicationGatewayRedirectTypePtrOutputWithContext(ctx context.Context) ApplicationGatewayRedirectTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayRedirectType) *ApplicationGatewayRedirectType {
		return &v
	}).(ApplicationGatewayRedirectTypePtrOutput)
}

func (o ApplicationGatewayRedirectTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewayRedirectTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayRedirectType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewayRedirectTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayRedirectTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayRedirectType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayRedirectTypePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRedirectTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayRedirectType)(nil)).Elem()
}

func (o ApplicationGatewayRedirectTypePtrOutput) ToApplicationGatewayRedirectTypePtrOutput() ApplicationGatewayRedirectTypePtrOutput {
	return o
}

func (o ApplicationGatewayRedirectTypePtrOutput) ToApplicationGatewayRedirectTypePtrOutputWithContext(ctx context.Context) ApplicationGatewayRedirectTypePtrOutput {
	return o
}

func (o ApplicationGatewayRedirectTypePtrOutput) Elem() ApplicationGatewayRedirectTypeOutput {
	return o.ApplyT(func(v *ApplicationGatewayRedirectType) ApplicationGatewayRedirectType {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayRedirectType
		return ret
	}).(ApplicationGatewayRedirectTypeOutput)
}

func (o ApplicationGatewayRedirectTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayRedirectTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewayRedirectType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewayRedirectTypeInput interface {
	pulumi.Input

	ToApplicationGatewayRedirectTypeOutput() ApplicationGatewayRedirectTypeOutput
	ToApplicationGatewayRedirectTypeOutputWithContext(context.Context) ApplicationGatewayRedirectTypeOutput
}

var applicationGatewayRedirectTypePtrType = reflect.TypeOf((**ApplicationGatewayRedirectType)(nil)).Elem()

type ApplicationGatewayRedirectTypePtrInput interface {
	pulumi.Input

	ToApplicationGatewayRedirectTypePtrOutput() ApplicationGatewayRedirectTypePtrOutput
	ToApplicationGatewayRedirectTypePtrOutputWithContext(context.Context) ApplicationGatewayRedirectTypePtrOutput
}

type applicationGatewayRedirectTypePtr string

func ApplicationGatewayRedirectTypePtr(v string) ApplicationGatewayRedirectTypePtrInput {
	return (*applicationGatewayRedirectTypePtr)(&v)
}

func (*applicationGatewayRedirectTypePtr) ElementType() reflect.Type {
	return applicationGatewayRedirectTypePtrType
}

func (in *applicationGatewayRedirectTypePtr) ToApplicationGatewayRedirectTypePtrOutput() ApplicationGatewayRedirectTypePtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewayRedirectTypePtrOutput)
}

func (in *applicationGatewayRedirectTypePtr) ToApplicationGatewayRedirectTypePtrOutputWithContext(ctx context.Context) ApplicationGatewayRedirectTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewayRedirectTypePtrOutput)
}

type ApplicationGatewayRequestRoutingRuleType string

const (
	ApplicationGatewayRequestRoutingRuleTypeBasic            = ApplicationGatewayRequestRoutingRuleType("Basic")
	ApplicationGatewayRequestRoutingRuleTypePathBasedRouting = ApplicationGatewayRequestRoutingRuleType("PathBasedRouting")
)

func (ApplicationGatewayRequestRoutingRuleType) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRequestRoutingRuleType)(nil)).Elem()
}

func (e ApplicationGatewayRequestRoutingRuleType) ToApplicationGatewayRequestRoutingRuleTypeOutput() ApplicationGatewayRequestRoutingRuleTypeOutput {
	return pulumi.ToOutput(e).(ApplicationGatewayRequestRoutingRuleTypeOutput)
}

func (e ApplicationGatewayRequestRoutingRuleType) ToApplicationGatewayRequestRoutingRuleTypeOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewayRequestRoutingRuleTypeOutput)
}

func (e ApplicationGatewayRequestRoutingRuleType) ToApplicationGatewayRequestRoutingRuleTypePtrOutput() ApplicationGatewayRequestRoutingRuleTypePtrOutput {
	return e.ToApplicationGatewayRequestRoutingRuleTypePtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayRequestRoutingRuleType) ToApplicationGatewayRequestRoutingRuleTypePtrOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleTypePtrOutput {
	return ApplicationGatewayRequestRoutingRuleType(e).ToApplicationGatewayRequestRoutingRuleTypeOutputWithContext(ctx).ToApplicationGatewayRequestRoutingRuleTypePtrOutputWithContext(ctx)
}

func (e ApplicationGatewayRequestRoutingRuleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayRequestRoutingRuleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayRequestRoutingRuleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayRequestRoutingRuleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewayRequestRoutingRuleTypeOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRequestRoutingRuleType)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleTypeOutput) ToApplicationGatewayRequestRoutingRuleTypeOutput() ApplicationGatewayRequestRoutingRuleTypeOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleTypeOutput) ToApplicationGatewayRequestRoutingRuleTypeOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleTypeOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleTypeOutput) ToApplicationGatewayRequestRoutingRuleTypePtrOutput() ApplicationGatewayRequestRoutingRuleTypePtrOutput {
	return o.ToApplicationGatewayRequestRoutingRuleTypePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayRequestRoutingRuleTypeOutput) ToApplicationGatewayRequestRoutingRuleTypePtrOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayRequestRoutingRuleType) *ApplicationGatewayRequestRoutingRuleType {
		return &v
	}).(ApplicationGatewayRequestRoutingRuleTypePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewayRequestRoutingRuleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayRequestRoutingRuleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewayRequestRoutingRuleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayRequestRoutingRuleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayRequestRoutingRuleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayRequestRoutingRuleTypePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayRequestRoutingRuleType)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleTypePtrOutput) ToApplicationGatewayRequestRoutingRuleTypePtrOutput() ApplicationGatewayRequestRoutingRuleTypePtrOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleTypePtrOutput) ToApplicationGatewayRequestRoutingRuleTypePtrOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleTypePtrOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleTypePtrOutput) Elem() ApplicationGatewayRequestRoutingRuleTypeOutput {
	return o.ApplyT(func(v *ApplicationGatewayRequestRoutingRuleType) ApplicationGatewayRequestRoutingRuleType {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayRequestRoutingRuleType
		return ret
	}).(ApplicationGatewayRequestRoutingRuleTypeOutput)
}

func (o ApplicationGatewayRequestRoutingRuleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayRequestRoutingRuleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewayRequestRoutingRuleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewayRequestRoutingRuleTypeInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleTypeOutput() ApplicationGatewayRequestRoutingRuleTypeOutput
	ToApplicationGatewayRequestRoutingRuleTypeOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleTypeOutput
}

var applicationGatewayRequestRoutingRuleTypePtrType = reflect.TypeOf((**ApplicationGatewayRequestRoutingRuleType)(nil)).Elem()

type ApplicationGatewayRequestRoutingRuleTypePtrInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleTypePtrOutput() ApplicationGatewayRequestRoutingRuleTypePtrOutput
	ToApplicationGatewayRequestRoutingRuleTypePtrOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleTypePtrOutput
}

type applicationGatewayRequestRoutingRuleTypePtr string

func ApplicationGatewayRequestRoutingRuleTypePtr(v string) ApplicationGatewayRequestRoutingRuleTypePtrInput {
	return (*applicationGatewayRequestRoutingRuleTypePtr)(&v)
}

func (*applicationGatewayRequestRoutingRuleTypePtr) ElementType() reflect.Type {
	return applicationGatewayRequestRoutingRuleTypePtrType
}

func (in *applicationGatewayRequestRoutingRuleTypePtr) ToApplicationGatewayRequestRoutingRuleTypePtrOutput() ApplicationGatewayRequestRoutingRuleTypePtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewayRequestRoutingRuleTypePtrOutput)
}

func (in *applicationGatewayRequestRoutingRuleTypePtr) ToApplicationGatewayRequestRoutingRuleTypePtrOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewayRequestRoutingRuleTypePtrOutput)
}

type ApplicationGatewaySkuName string

const (
	ApplicationGatewaySkuName_Standard_Small  = ApplicationGatewaySkuName("Standard_Small")
	ApplicationGatewaySkuName_Standard_Medium = ApplicationGatewaySkuName("Standard_Medium")
	ApplicationGatewaySkuName_Standard_Large  = ApplicationGatewaySkuName("Standard_Large")
	ApplicationGatewaySkuName_WAF_Medium      = ApplicationGatewaySkuName("WAF_Medium")
	ApplicationGatewaySkuName_WAF_Large       = ApplicationGatewaySkuName("WAF_Large")
)

func (ApplicationGatewaySkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySkuName)(nil)).Elem()
}

func (e ApplicationGatewaySkuName) ToApplicationGatewaySkuNameOutput() ApplicationGatewaySkuNameOutput {
	return pulumi.ToOutput(e).(ApplicationGatewaySkuNameOutput)
}

func (e ApplicationGatewaySkuName) ToApplicationGatewaySkuNameOutputWithContext(ctx context.Context) ApplicationGatewaySkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewaySkuNameOutput)
}

func (e ApplicationGatewaySkuName) ToApplicationGatewaySkuNamePtrOutput() ApplicationGatewaySkuNamePtrOutput {
	return e.ToApplicationGatewaySkuNamePtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySkuName) ToApplicationGatewaySkuNamePtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuNamePtrOutput {
	return ApplicationGatewaySkuName(e).ToApplicationGatewaySkuNameOutputWithContext(ctx).ToApplicationGatewaySkuNamePtrOutputWithContext(ctx)
}

func (e ApplicationGatewaySkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewaySkuNameOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySkuName)(nil)).Elem()
}

func (o ApplicationGatewaySkuNameOutput) ToApplicationGatewaySkuNameOutput() ApplicationGatewaySkuNameOutput {
	return o
}

func (o ApplicationGatewaySkuNameOutput) ToApplicationGatewaySkuNameOutputWithContext(ctx context.Context) ApplicationGatewaySkuNameOutput {
	return o
}

func (o ApplicationGatewaySkuNameOutput) ToApplicationGatewaySkuNamePtrOutput() ApplicationGatewaySkuNamePtrOutput {
	return o.ToApplicationGatewaySkuNamePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySkuNameOutput) ToApplicationGatewaySkuNamePtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySkuName) *ApplicationGatewaySkuName {
		return &v
	}).(ApplicationGatewaySkuNamePtrOutput)
}

func (o ApplicationGatewaySkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewaySkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewaySkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySkuNamePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySkuName)(nil)).Elem()
}

func (o ApplicationGatewaySkuNamePtrOutput) ToApplicationGatewaySkuNamePtrOutput() ApplicationGatewaySkuNamePtrOutput {
	return o
}

func (o ApplicationGatewaySkuNamePtrOutput) ToApplicationGatewaySkuNamePtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuNamePtrOutput {
	return o
}

func (o ApplicationGatewaySkuNamePtrOutput) Elem() ApplicationGatewaySkuNameOutput {
	return o.ApplyT(func(v *ApplicationGatewaySkuName) ApplicationGatewaySkuName {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySkuName
		return ret
	}).(ApplicationGatewaySkuNameOutput)
}

func (o ApplicationGatewaySkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewaySkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewaySkuNameInput interface {
	pulumi.Input

	ToApplicationGatewaySkuNameOutput() ApplicationGatewaySkuNameOutput
	ToApplicationGatewaySkuNameOutputWithContext(context.Context) ApplicationGatewaySkuNameOutput
}

var applicationGatewaySkuNamePtrType = reflect.TypeOf((**ApplicationGatewaySkuName)(nil)).Elem()

type ApplicationGatewaySkuNamePtrInput interface {
	pulumi.Input

	ToApplicationGatewaySkuNamePtrOutput() ApplicationGatewaySkuNamePtrOutput
	ToApplicationGatewaySkuNamePtrOutputWithContext(context.Context) ApplicationGatewaySkuNamePtrOutput
}

type applicationGatewaySkuNamePtr string

func ApplicationGatewaySkuNamePtr(v string) ApplicationGatewaySkuNamePtrInput {
	return (*applicationGatewaySkuNamePtr)(&v)
}

func (*applicationGatewaySkuNamePtr) ElementType() reflect.Type {
	return applicationGatewaySkuNamePtrType
}

func (in *applicationGatewaySkuNamePtr) ToApplicationGatewaySkuNamePtrOutput() ApplicationGatewaySkuNamePtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewaySkuNamePtrOutput)
}

func (in *applicationGatewaySkuNamePtr) ToApplicationGatewaySkuNamePtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewaySkuNamePtrOutput)
}

type ApplicationGatewaySslCipherSuite string

const (
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384   = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256   = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA      = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA      = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_RSA_WITH_AES_256_GCM_SHA384     = ApplicationGatewaySslCipherSuite("TLS_DHE_RSA_WITH_AES_256_GCM_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_DHE_RSA_WITH_AES_128_GCM_SHA256     = ApplicationGatewaySslCipherSuite("TLS_DHE_RSA_WITH_AES_128_GCM_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_DHE_RSA_WITH_AES_256_CBC_SHA        = ApplicationGatewaySslCipherSuite("TLS_DHE_RSA_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_RSA_WITH_AES_128_CBC_SHA        = ApplicationGatewaySslCipherSuite("TLS_DHE_RSA_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_256_GCM_SHA384         = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_256_GCM_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_128_GCM_SHA256         = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_128_GCM_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_256_CBC_SHA256         = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_256_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_128_CBC_SHA256         = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_128_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_256_CBC_SHA            = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_AES_128_CBC_SHA            = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384 = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256 = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA    = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA    = ApplicationGatewaySslCipherSuite("TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_AES_256_CBC_SHA256     = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_AES_256_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_AES_128_CBC_SHA256     = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_AES_128_CBC_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_AES_256_CBC_SHA        = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_AES_256_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_AES_128_CBC_SHA        = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_AES_128_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_RSA_WITH_3DES_EDE_CBC_SHA           = ApplicationGatewaySslCipherSuite("TLS_RSA_WITH_3DES_EDE_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA       = ApplicationGatewaySslCipherSuite("TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256   = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256")
	ApplicationGatewaySslCipherSuite_TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384   = ApplicationGatewaySslCipherSuite("TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384")
)

func (ApplicationGatewaySslCipherSuite) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslCipherSuite)(nil)).Elem()
}

func (e ApplicationGatewaySslCipherSuite) ToApplicationGatewaySslCipherSuiteOutput() ApplicationGatewaySslCipherSuiteOutput {
	return pulumi.ToOutput(e).(ApplicationGatewaySslCipherSuiteOutput)
}

func (e ApplicationGatewaySslCipherSuite) ToApplicationGatewaySslCipherSuiteOutputWithContext(ctx context.Context) ApplicationGatewaySslCipherSuiteOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewaySslCipherSuiteOutput)
}

func (e ApplicationGatewaySslCipherSuite) ToApplicationGatewaySslCipherSuitePtrOutput() ApplicationGatewaySslCipherSuitePtrOutput {
	return e.ToApplicationGatewaySslCipherSuitePtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySslCipherSuite) ToApplicationGatewaySslCipherSuitePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslCipherSuitePtrOutput {
	return ApplicationGatewaySslCipherSuite(e).ToApplicationGatewaySslCipherSuiteOutputWithContext(ctx).ToApplicationGatewaySslCipherSuitePtrOutputWithContext(ctx)
}

func (e ApplicationGatewaySslCipherSuite) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySslCipherSuite) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySslCipherSuite) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySslCipherSuite) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewaySslCipherSuiteOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCipherSuiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslCipherSuite)(nil)).Elem()
}

func (o ApplicationGatewaySslCipherSuiteOutput) ToApplicationGatewaySslCipherSuiteOutput() ApplicationGatewaySslCipherSuiteOutput {
	return o
}

func (o ApplicationGatewaySslCipherSuiteOutput) ToApplicationGatewaySslCipherSuiteOutputWithContext(ctx context.Context) ApplicationGatewaySslCipherSuiteOutput {
	return o
}

func (o ApplicationGatewaySslCipherSuiteOutput) ToApplicationGatewaySslCipherSuitePtrOutput() ApplicationGatewaySslCipherSuitePtrOutput {
	return o.ToApplicationGatewaySslCipherSuitePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslCipherSuiteOutput) ToApplicationGatewaySslCipherSuitePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslCipherSuitePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySslCipherSuite) *ApplicationGatewaySslCipherSuite {
		return &v
	}).(ApplicationGatewaySslCipherSuitePtrOutput)
}

func (o ApplicationGatewaySslCipherSuiteOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslCipherSuiteOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySslCipherSuite) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewaySslCipherSuiteOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslCipherSuiteOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySslCipherSuite) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslCipherSuitePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCipherSuitePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslCipherSuite)(nil)).Elem()
}

func (o ApplicationGatewaySslCipherSuitePtrOutput) ToApplicationGatewaySslCipherSuitePtrOutput() ApplicationGatewaySslCipherSuitePtrOutput {
	return o
}

func (o ApplicationGatewaySslCipherSuitePtrOutput) ToApplicationGatewaySslCipherSuitePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslCipherSuitePtrOutput {
	return o
}

func (o ApplicationGatewaySslCipherSuitePtrOutput) Elem() ApplicationGatewaySslCipherSuiteOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslCipherSuite) ApplicationGatewaySslCipherSuite {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySslCipherSuite
		return ret
	}).(ApplicationGatewaySslCipherSuiteOutput)
}

func (o ApplicationGatewaySslCipherSuitePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslCipherSuitePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewaySslCipherSuite) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewaySslCipherSuiteInput interface {
	pulumi.Input

	ToApplicationGatewaySslCipherSuiteOutput() ApplicationGatewaySslCipherSuiteOutput
	ToApplicationGatewaySslCipherSuiteOutputWithContext(context.Context) ApplicationGatewaySslCipherSuiteOutput
}

var applicationGatewaySslCipherSuitePtrType = reflect.TypeOf((**ApplicationGatewaySslCipherSuite)(nil)).Elem()

type ApplicationGatewaySslCipherSuitePtrInput interface {
	pulumi.Input

	ToApplicationGatewaySslCipherSuitePtrOutput() ApplicationGatewaySslCipherSuitePtrOutput
	ToApplicationGatewaySslCipherSuitePtrOutputWithContext(context.Context) ApplicationGatewaySslCipherSuitePtrOutput
}

type applicationGatewaySslCipherSuitePtr string

func ApplicationGatewaySslCipherSuitePtr(v string) ApplicationGatewaySslCipherSuitePtrInput {
	return (*applicationGatewaySslCipherSuitePtr)(&v)
}

func (*applicationGatewaySslCipherSuitePtr) ElementType() reflect.Type {
	return applicationGatewaySslCipherSuitePtrType
}

func (in *applicationGatewaySslCipherSuitePtr) ToApplicationGatewaySslCipherSuitePtrOutput() ApplicationGatewaySslCipherSuitePtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewaySslCipherSuitePtrOutput)
}

func (in *applicationGatewaySslCipherSuitePtr) ToApplicationGatewaySslCipherSuitePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslCipherSuitePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewaySslCipherSuitePtrOutput)
}

type ApplicationGatewaySslPolicyName string

const (
	ApplicationGatewaySslPolicyNameAppGwSslPolicy20150501  = ApplicationGatewaySslPolicyName("AppGwSslPolicy20150501")
	ApplicationGatewaySslPolicyNameAppGwSslPolicy20170401  = ApplicationGatewaySslPolicyName("AppGwSslPolicy20170401")
	ApplicationGatewaySslPolicyNameAppGwSslPolicy20170401S = ApplicationGatewaySslPolicyName("AppGwSslPolicy20170401S")
)

func (ApplicationGatewaySslPolicyName) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicyName)(nil)).Elem()
}

func (e ApplicationGatewaySslPolicyName) ToApplicationGatewaySslPolicyNameOutput() ApplicationGatewaySslPolicyNameOutput {
	return pulumi.ToOutput(e).(ApplicationGatewaySslPolicyNameOutput)
}

func (e ApplicationGatewaySslPolicyName) ToApplicationGatewaySslPolicyNameOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewaySslPolicyNameOutput)
}

func (e ApplicationGatewaySslPolicyName) ToApplicationGatewaySslPolicyNamePtrOutput() ApplicationGatewaySslPolicyNamePtrOutput {
	return e.ToApplicationGatewaySslPolicyNamePtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySslPolicyName) ToApplicationGatewaySslPolicyNamePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyNamePtrOutput {
	return ApplicationGatewaySslPolicyName(e).ToApplicationGatewaySslPolicyNameOutputWithContext(ctx).ToApplicationGatewaySslPolicyNamePtrOutputWithContext(ctx)
}

func (e ApplicationGatewaySslPolicyName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySslPolicyName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySslPolicyName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySslPolicyName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewaySslPolicyNameOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicyName)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyNameOutput) ToApplicationGatewaySslPolicyNameOutput() ApplicationGatewaySslPolicyNameOutput {
	return o
}

func (o ApplicationGatewaySslPolicyNameOutput) ToApplicationGatewaySslPolicyNameOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyNameOutput {
	return o
}

func (o ApplicationGatewaySslPolicyNameOutput) ToApplicationGatewaySslPolicyNamePtrOutput() ApplicationGatewaySslPolicyNamePtrOutput {
	return o.ToApplicationGatewaySslPolicyNamePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyNameOutput) ToApplicationGatewaySslPolicyNamePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySslPolicyName) *ApplicationGatewaySslPolicyName {
		return &v
	}).(ApplicationGatewaySslPolicyNamePtrOutput)
}

func (o ApplicationGatewaySslPolicyNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySslPolicyName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewaySslPolicyNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySslPolicyName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslPolicyNamePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslPolicyName)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyNamePtrOutput) ToApplicationGatewaySslPolicyNamePtrOutput() ApplicationGatewaySslPolicyNamePtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyNamePtrOutput) ToApplicationGatewaySslPolicyNamePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyNamePtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyNamePtrOutput) Elem() ApplicationGatewaySslPolicyNameOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyName) ApplicationGatewaySslPolicyName {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySslPolicyName
		return ret
	}).(ApplicationGatewaySslPolicyNameOutput)
}

func (o ApplicationGatewaySslPolicyNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewaySslPolicyName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewaySslPolicyNameInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyNameOutput() ApplicationGatewaySslPolicyNameOutput
	ToApplicationGatewaySslPolicyNameOutputWithContext(context.Context) ApplicationGatewaySslPolicyNameOutput
}

var applicationGatewaySslPolicyNamePtrType = reflect.TypeOf((**ApplicationGatewaySslPolicyName)(nil)).Elem()

type ApplicationGatewaySslPolicyNamePtrInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyNamePtrOutput() ApplicationGatewaySslPolicyNamePtrOutput
	ToApplicationGatewaySslPolicyNamePtrOutputWithContext(context.Context) ApplicationGatewaySslPolicyNamePtrOutput
}

type applicationGatewaySslPolicyNamePtr string

func ApplicationGatewaySslPolicyNamePtr(v string) ApplicationGatewaySslPolicyNamePtrInput {
	return (*applicationGatewaySslPolicyNamePtr)(&v)
}

func (*applicationGatewaySslPolicyNamePtr) ElementType() reflect.Type {
	return applicationGatewaySslPolicyNamePtrType
}

func (in *applicationGatewaySslPolicyNamePtr) ToApplicationGatewaySslPolicyNamePtrOutput() ApplicationGatewaySslPolicyNamePtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewaySslPolicyNamePtrOutput)
}

func (in *applicationGatewaySslPolicyNamePtr) ToApplicationGatewaySslPolicyNamePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewaySslPolicyNamePtrOutput)
}

type ApplicationGatewaySslPolicyType string

const (
	ApplicationGatewaySslPolicyTypePredefined = ApplicationGatewaySslPolicyType("Predefined")
	ApplicationGatewaySslPolicyTypeCustom     = ApplicationGatewaySslPolicyType("Custom")
)

func (ApplicationGatewaySslPolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicyType)(nil)).Elem()
}

func (e ApplicationGatewaySslPolicyType) ToApplicationGatewaySslPolicyTypeOutput() ApplicationGatewaySslPolicyTypeOutput {
	return pulumi.ToOutput(e).(ApplicationGatewaySslPolicyTypeOutput)
}

func (e ApplicationGatewaySslPolicyType) ToApplicationGatewaySslPolicyTypeOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewaySslPolicyTypeOutput)
}

func (e ApplicationGatewaySslPolicyType) ToApplicationGatewaySslPolicyTypePtrOutput() ApplicationGatewaySslPolicyTypePtrOutput {
	return e.ToApplicationGatewaySslPolicyTypePtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySslPolicyType) ToApplicationGatewaySslPolicyTypePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyTypePtrOutput {
	return ApplicationGatewaySslPolicyType(e).ToApplicationGatewaySslPolicyTypeOutputWithContext(ctx).ToApplicationGatewaySslPolicyTypePtrOutputWithContext(ctx)
}

func (e ApplicationGatewaySslPolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySslPolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySslPolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySslPolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewaySslPolicyTypeOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicyType)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyTypeOutput) ToApplicationGatewaySslPolicyTypeOutput() ApplicationGatewaySslPolicyTypeOutput {
	return o
}

func (o ApplicationGatewaySslPolicyTypeOutput) ToApplicationGatewaySslPolicyTypeOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyTypeOutput {
	return o
}

func (o ApplicationGatewaySslPolicyTypeOutput) ToApplicationGatewaySslPolicyTypePtrOutput() ApplicationGatewaySslPolicyTypePtrOutput {
	return o.ToApplicationGatewaySslPolicyTypePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyTypeOutput) ToApplicationGatewaySslPolicyTypePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySslPolicyType) *ApplicationGatewaySslPolicyType {
		return &v
	}).(ApplicationGatewaySslPolicyTypePtrOutput)
}

func (o ApplicationGatewaySslPolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySslPolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewaySslPolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySslPolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslPolicyTypePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslPolicyType)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyTypePtrOutput) ToApplicationGatewaySslPolicyTypePtrOutput() ApplicationGatewaySslPolicyTypePtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyTypePtrOutput) ToApplicationGatewaySslPolicyTypePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyTypePtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyTypePtrOutput) Elem() ApplicationGatewaySslPolicyTypeOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyType) ApplicationGatewaySslPolicyType {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySslPolicyType
		return ret
	}).(ApplicationGatewaySslPolicyTypeOutput)
}

func (o ApplicationGatewaySslPolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewaySslPolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewaySslPolicyTypeInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyTypeOutput() ApplicationGatewaySslPolicyTypeOutput
	ToApplicationGatewaySslPolicyTypeOutputWithContext(context.Context) ApplicationGatewaySslPolicyTypeOutput
}

var applicationGatewaySslPolicyTypePtrType = reflect.TypeOf((**ApplicationGatewaySslPolicyType)(nil)).Elem()

type ApplicationGatewaySslPolicyTypePtrInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyTypePtrOutput() ApplicationGatewaySslPolicyTypePtrOutput
	ToApplicationGatewaySslPolicyTypePtrOutputWithContext(context.Context) ApplicationGatewaySslPolicyTypePtrOutput
}

type applicationGatewaySslPolicyTypePtr string

func ApplicationGatewaySslPolicyTypePtr(v string) ApplicationGatewaySslPolicyTypePtrInput {
	return (*applicationGatewaySslPolicyTypePtr)(&v)
}

func (*applicationGatewaySslPolicyTypePtr) ElementType() reflect.Type {
	return applicationGatewaySslPolicyTypePtrType
}

func (in *applicationGatewaySslPolicyTypePtr) ToApplicationGatewaySslPolicyTypePtrOutput() ApplicationGatewaySslPolicyTypePtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewaySslPolicyTypePtrOutput)
}

func (in *applicationGatewaySslPolicyTypePtr) ToApplicationGatewaySslPolicyTypePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewaySslPolicyTypePtrOutput)
}

type ApplicationGatewaySslProtocol string

const (
	ApplicationGatewaySslProtocol_TLSv1_0 = ApplicationGatewaySslProtocol("TLSv1_0")
	ApplicationGatewaySslProtocol_TLSv1_1 = ApplicationGatewaySslProtocol("TLSv1_1")
	ApplicationGatewaySslProtocol_TLSv1_2 = ApplicationGatewaySslProtocol("TLSv1_2")
)

func (ApplicationGatewaySslProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslProtocol)(nil)).Elem()
}

func (e ApplicationGatewaySslProtocol) ToApplicationGatewaySslProtocolOutput() ApplicationGatewaySslProtocolOutput {
	return pulumi.ToOutput(e).(ApplicationGatewaySslProtocolOutput)
}

func (e ApplicationGatewaySslProtocol) ToApplicationGatewaySslProtocolOutputWithContext(ctx context.Context) ApplicationGatewaySslProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewaySslProtocolOutput)
}

func (e ApplicationGatewaySslProtocol) ToApplicationGatewaySslProtocolPtrOutput() ApplicationGatewaySslProtocolPtrOutput {
	return e.ToApplicationGatewaySslProtocolPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySslProtocol) ToApplicationGatewaySslProtocolPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslProtocolPtrOutput {
	return ApplicationGatewaySslProtocol(e).ToApplicationGatewaySslProtocolOutputWithContext(ctx).ToApplicationGatewaySslProtocolPtrOutputWithContext(ctx)
}

func (e ApplicationGatewaySslProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySslProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewaySslProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewaySslProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewaySslProtocolOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslProtocol)(nil)).Elem()
}

func (o ApplicationGatewaySslProtocolOutput) ToApplicationGatewaySslProtocolOutput() ApplicationGatewaySslProtocolOutput {
	return o
}

func (o ApplicationGatewaySslProtocolOutput) ToApplicationGatewaySslProtocolOutputWithContext(ctx context.Context) ApplicationGatewaySslProtocolOutput {
	return o
}

func (o ApplicationGatewaySslProtocolOutput) ToApplicationGatewaySslProtocolPtrOutput() ApplicationGatewaySslProtocolPtrOutput {
	return o.ToApplicationGatewaySslProtocolPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslProtocolOutput) ToApplicationGatewaySslProtocolPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySslProtocol) *ApplicationGatewaySslProtocol {
		return &v
	}).(ApplicationGatewaySslProtocolPtrOutput)
}

func (o ApplicationGatewaySslProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySslProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewaySslProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewaySslProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslProtocolPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslProtocol)(nil)).Elem()
}

func (o ApplicationGatewaySslProtocolPtrOutput) ToApplicationGatewaySslProtocolPtrOutput() ApplicationGatewaySslProtocolPtrOutput {
	return o
}

func (o ApplicationGatewaySslProtocolPtrOutput) ToApplicationGatewaySslProtocolPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslProtocolPtrOutput {
	return o
}

func (o ApplicationGatewaySslProtocolPtrOutput) Elem() ApplicationGatewaySslProtocolOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslProtocol) ApplicationGatewaySslProtocol {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySslProtocol
		return ret
	}).(ApplicationGatewaySslProtocolOutput)
}

func (o ApplicationGatewaySslProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewaySslProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewaySslProtocolInput interface {
	pulumi.Input

	ToApplicationGatewaySslProtocolOutput() ApplicationGatewaySslProtocolOutput
	ToApplicationGatewaySslProtocolOutputWithContext(context.Context) ApplicationGatewaySslProtocolOutput
}

var applicationGatewaySslProtocolPtrType = reflect.TypeOf((**ApplicationGatewaySslProtocol)(nil)).Elem()

type ApplicationGatewaySslProtocolPtrInput interface {
	pulumi.Input

	ToApplicationGatewaySslProtocolPtrOutput() ApplicationGatewaySslProtocolPtrOutput
	ToApplicationGatewaySslProtocolPtrOutputWithContext(context.Context) ApplicationGatewaySslProtocolPtrOutput
}

type applicationGatewaySslProtocolPtr string

func ApplicationGatewaySslProtocolPtr(v string) ApplicationGatewaySslProtocolPtrInput {
	return (*applicationGatewaySslProtocolPtr)(&v)
}

func (*applicationGatewaySslProtocolPtr) ElementType() reflect.Type {
	return applicationGatewaySslProtocolPtrType
}

func (in *applicationGatewaySslProtocolPtr) ToApplicationGatewaySslProtocolPtrOutput() ApplicationGatewaySslProtocolPtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewaySslProtocolPtrOutput)
}

func (in *applicationGatewaySslProtocolPtr) ToApplicationGatewaySslProtocolPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewaySslProtocolPtrOutput)
}

type ApplicationGatewayTier string

const (
	ApplicationGatewayTierStandard = ApplicationGatewayTier("Standard")
	ApplicationGatewayTierWAF      = ApplicationGatewayTier("WAF")
)

func (ApplicationGatewayTier) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayTier)(nil)).Elem()
}

func (e ApplicationGatewayTier) ToApplicationGatewayTierOutput() ApplicationGatewayTierOutput {
	return pulumi.ToOutput(e).(ApplicationGatewayTierOutput)
}

func (e ApplicationGatewayTier) ToApplicationGatewayTierOutputWithContext(ctx context.Context) ApplicationGatewayTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewayTierOutput)
}

func (e ApplicationGatewayTier) ToApplicationGatewayTierPtrOutput() ApplicationGatewayTierPtrOutput {
	return e.ToApplicationGatewayTierPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayTier) ToApplicationGatewayTierPtrOutputWithContext(ctx context.Context) ApplicationGatewayTierPtrOutput {
	return ApplicationGatewayTier(e).ToApplicationGatewayTierOutputWithContext(ctx).ToApplicationGatewayTierPtrOutputWithContext(ctx)
}

func (e ApplicationGatewayTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewayTierOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayTier)(nil)).Elem()
}

func (o ApplicationGatewayTierOutput) ToApplicationGatewayTierOutput() ApplicationGatewayTierOutput {
	return o
}

func (o ApplicationGatewayTierOutput) ToApplicationGatewayTierOutputWithContext(ctx context.Context) ApplicationGatewayTierOutput {
	return o
}

func (o ApplicationGatewayTierOutput) ToApplicationGatewayTierPtrOutput() ApplicationGatewayTierPtrOutput {
	return o.ToApplicationGatewayTierPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayTierOutput) ToApplicationGatewayTierPtrOutputWithContext(ctx context.Context) ApplicationGatewayTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayTier) *ApplicationGatewayTier {
		return &v
	}).(ApplicationGatewayTierPtrOutput)
}

func (o ApplicationGatewayTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewayTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewayTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayTierPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayTier)(nil)).Elem()
}

func (o ApplicationGatewayTierPtrOutput) ToApplicationGatewayTierPtrOutput() ApplicationGatewayTierPtrOutput {
	return o
}

func (o ApplicationGatewayTierPtrOutput) ToApplicationGatewayTierPtrOutputWithContext(ctx context.Context) ApplicationGatewayTierPtrOutput {
	return o
}

func (o ApplicationGatewayTierPtrOutput) Elem() ApplicationGatewayTierOutput {
	return o.ApplyT(func(v *ApplicationGatewayTier) ApplicationGatewayTier {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayTier
		return ret
	}).(ApplicationGatewayTierOutput)
}

func (o ApplicationGatewayTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewayTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewayTierInput interface {
	pulumi.Input

	ToApplicationGatewayTierOutput() ApplicationGatewayTierOutput
	ToApplicationGatewayTierOutputWithContext(context.Context) ApplicationGatewayTierOutput
}

var applicationGatewayTierPtrType = reflect.TypeOf((**ApplicationGatewayTier)(nil)).Elem()

type ApplicationGatewayTierPtrInput interface {
	pulumi.Input

	ToApplicationGatewayTierPtrOutput() ApplicationGatewayTierPtrOutput
	ToApplicationGatewayTierPtrOutputWithContext(context.Context) ApplicationGatewayTierPtrOutput
}

type applicationGatewayTierPtr string

func ApplicationGatewayTierPtr(v string) ApplicationGatewayTierPtrInput {
	return (*applicationGatewayTierPtr)(&v)
}

func (*applicationGatewayTierPtr) ElementType() reflect.Type {
	return applicationGatewayTierPtrType
}

func (in *applicationGatewayTierPtr) ToApplicationGatewayTierPtrOutput() ApplicationGatewayTierPtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewayTierPtrOutput)
}

func (in *applicationGatewayTierPtr) ToApplicationGatewayTierPtrOutputWithContext(ctx context.Context) ApplicationGatewayTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewayTierPtrOutput)
}

type AuthorizationUseStatus string

const (
	AuthorizationUseStatusAvailable = AuthorizationUseStatus("Available")
	AuthorizationUseStatusInUse     = AuthorizationUseStatus("InUse")
)

func (AuthorizationUseStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationUseStatus)(nil)).Elem()
}

func (e AuthorizationUseStatus) ToAuthorizationUseStatusOutput() AuthorizationUseStatusOutput {
	return pulumi.ToOutput(e).(AuthorizationUseStatusOutput)
}

func (e AuthorizationUseStatus) ToAuthorizationUseStatusOutputWithContext(ctx context.Context) AuthorizationUseStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthorizationUseStatusOutput)
}

func (e AuthorizationUseStatus) ToAuthorizationUseStatusPtrOutput() AuthorizationUseStatusPtrOutput {
	return e.ToAuthorizationUseStatusPtrOutputWithContext(context.Background())
}

func (e AuthorizationUseStatus) ToAuthorizationUseStatusPtrOutputWithContext(ctx context.Context) AuthorizationUseStatusPtrOutput {
	return AuthorizationUseStatus(e).ToAuthorizationUseStatusOutputWithContext(ctx).ToAuthorizationUseStatusPtrOutputWithContext(ctx)
}

func (e AuthorizationUseStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthorizationUseStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthorizationUseStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthorizationUseStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthorizationUseStatusOutput struct{ *pulumi.OutputState }

func (AuthorizationUseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationUseStatus)(nil)).Elem()
}

func (o AuthorizationUseStatusOutput) ToAuthorizationUseStatusOutput() AuthorizationUseStatusOutput {
	return o
}

func (o AuthorizationUseStatusOutput) ToAuthorizationUseStatusOutputWithContext(ctx context.Context) AuthorizationUseStatusOutput {
	return o
}

func (o AuthorizationUseStatusOutput) ToAuthorizationUseStatusPtrOutput() AuthorizationUseStatusPtrOutput {
	return o.ToAuthorizationUseStatusPtrOutputWithContext(context.Background())
}

func (o AuthorizationUseStatusOutput) ToAuthorizationUseStatusPtrOutputWithContext(ctx context.Context) AuthorizationUseStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthorizationUseStatus) *AuthorizationUseStatus {
		return &v
	}).(AuthorizationUseStatusPtrOutput)
}

func (o AuthorizationUseStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthorizationUseStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthorizationUseStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthorizationUseStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthorizationUseStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthorizationUseStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthorizationUseStatusPtrOutput struct{ *pulumi.OutputState }

func (AuthorizationUseStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationUseStatus)(nil)).Elem()
}

func (o AuthorizationUseStatusPtrOutput) ToAuthorizationUseStatusPtrOutput() AuthorizationUseStatusPtrOutput {
	return o
}

func (o AuthorizationUseStatusPtrOutput) ToAuthorizationUseStatusPtrOutputWithContext(ctx context.Context) AuthorizationUseStatusPtrOutput {
	return o
}

func (o AuthorizationUseStatusPtrOutput) Elem() AuthorizationUseStatusOutput {
	return o.ApplyT(func(v *AuthorizationUseStatus) AuthorizationUseStatus {
		if v != nil {
			return *v
		}
		var ret AuthorizationUseStatus
		return ret
	}).(AuthorizationUseStatusOutput)
}

func (o AuthorizationUseStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthorizationUseStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthorizationUseStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AuthorizationUseStatusInput interface {
	pulumi.Input

	ToAuthorizationUseStatusOutput() AuthorizationUseStatusOutput
	ToAuthorizationUseStatusOutputWithContext(context.Context) AuthorizationUseStatusOutput
}

var authorizationUseStatusPtrType = reflect.TypeOf((**AuthorizationUseStatus)(nil)).Elem()

type AuthorizationUseStatusPtrInput interface {
	pulumi.Input

	ToAuthorizationUseStatusPtrOutput() AuthorizationUseStatusPtrOutput
	ToAuthorizationUseStatusPtrOutputWithContext(context.Context) AuthorizationUseStatusPtrOutput
}

type authorizationUseStatusPtr string

func AuthorizationUseStatusPtr(v string) AuthorizationUseStatusPtrInput {
	return (*authorizationUseStatusPtr)(&v)
}

func (*authorizationUseStatusPtr) ElementType() reflect.Type {
	return authorizationUseStatusPtrType
}

func (in *authorizationUseStatusPtr) ToAuthorizationUseStatusPtrOutput() AuthorizationUseStatusPtrOutput {
	return pulumi.ToOutput(in).(AuthorizationUseStatusPtrOutput)
}

func (in *authorizationUseStatusPtr) ToAuthorizationUseStatusPtrOutputWithContext(ctx context.Context) AuthorizationUseStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthorizationUseStatusPtrOutput)
}

type DhGroup string

const (
	DhGroupNone        = DhGroup("None")
	DhGroupDHGroup1    = DhGroup("DHGroup1")
	DhGroupDHGroup2    = DhGroup("DHGroup2")
	DhGroupDHGroup14   = DhGroup("DHGroup14")
	DhGroupDHGroup2048 = DhGroup("DHGroup2048")
	DhGroupECP256      = DhGroup("ECP256")
	DhGroupECP384      = DhGroup("ECP384")
	DhGroupDHGroup24   = DhGroup("DHGroup24")
)

func (DhGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DhGroup)(nil)).Elem()
}

func (e DhGroup) ToDhGroupOutput() DhGroupOutput {
	return pulumi.ToOutput(e).(DhGroupOutput)
}

func (e DhGroup) ToDhGroupOutputWithContext(ctx context.Context) DhGroupOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DhGroupOutput)
}

func (e DhGroup) ToDhGroupPtrOutput() DhGroupPtrOutput {
	return e.ToDhGroupPtrOutputWithContext(context.Background())
}

func (e DhGroup) ToDhGroupPtrOutputWithContext(ctx context.Context) DhGroupPtrOutput {
	return DhGroup(e).ToDhGroupOutputWithContext(ctx).ToDhGroupPtrOutputWithContext(ctx)
}

func (e DhGroup) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DhGroup) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DhGroup) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DhGroup) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DhGroupOutput struct{ *pulumi.OutputState }

func (DhGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DhGroup)(nil)).Elem()
}

func (o DhGroupOutput) ToDhGroupOutput() DhGroupOutput {
	return o
}

func (o DhGroupOutput) ToDhGroupOutputWithContext(ctx context.Context) DhGroupOutput {
	return o
}

func (o DhGroupOutput) ToDhGroupPtrOutput() DhGroupPtrOutput {
	return o.ToDhGroupPtrOutputWithContext(context.Background())
}

func (o DhGroupOutput) ToDhGroupPtrOutputWithContext(ctx context.Context) DhGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DhGroup) *DhGroup {
		return &v
	}).(DhGroupPtrOutput)
}

func (o DhGroupOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DhGroupOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DhGroup) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DhGroupOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DhGroupOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DhGroup) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DhGroupPtrOutput struct{ *pulumi.OutputState }

func (DhGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhGroup)(nil)).Elem()
}

func (o DhGroupPtrOutput) ToDhGroupPtrOutput() DhGroupPtrOutput {
	return o
}

func (o DhGroupPtrOutput) ToDhGroupPtrOutputWithContext(ctx context.Context) DhGroupPtrOutput {
	return o
}

func (o DhGroupPtrOutput) Elem() DhGroupOutput {
	return o.ApplyT(func(v *DhGroup) DhGroup {
		if v != nil {
			return *v
		}
		var ret DhGroup
		return ret
	}).(DhGroupOutput)
}

func (o DhGroupPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DhGroupPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DhGroup) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DhGroupInput interface {
	pulumi.Input

	ToDhGroupOutput() DhGroupOutput
	ToDhGroupOutputWithContext(context.Context) DhGroupOutput
}

var dhGroupPtrType = reflect.TypeOf((**DhGroup)(nil)).Elem()

type DhGroupPtrInput interface {
	pulumi.Input

	ToDhGroupPtrOutput() DhGroupPtrOutput
	ToDhGroupPtrOutputWithContext(context.Context) DhGroupPtrOutput
}

type dhGroupPtr string

func DhGroupPtr(v string) DhGroupPtrInput {
	return (*dhGroupPtr)(&v)
}

func (*dhGroupPtr) ElementType() reflect.Type {
	return dhGroupPtrType
}

func (in *dhGroupPtr) ToDhGroupPtrOutput() DhGroupPtrOutput {
	return pulumi.ToOutput(in).(DhGroupPtrOutput)
}

func (in *dhGroupPtr) ToDhGroupPtrOutputWithContext(ctx context.Context) DhGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DhGroupPtrOutput)
}

type ExpressRouteCircuitPeeringAdvertisedPublicPrefixState string

const (
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateNotConfigured    = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("NotConfigured")
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateConfiguring      = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("Configuring")
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateConfigured       = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("Configured")
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateValidationNeeded = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("ValidationNeeded")
)

func (ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringAdvertisedPublicPrefixState)(nil)).Elem()
}

func (e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput() ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput {
	return pulumi.ToOutput(e).(ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput)
}

func (e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput)
}

func (e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput() ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput {
	return e.ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput {
	return ExpressRouteCircuitPeeringAdvertisedPublicPrefixState(e).ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutputWithContext(ctx).ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutputWithContext(ctx)
}

func (e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringAdvertisedPublicPrefixState)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput() ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput {
	return o
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput {
	return o
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput() ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput {
	return o.ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) *ExpressRouteCircuitPeeringAdvertisedPublicPrefixState {
		return &v
	}).(ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput)
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringAdvertisedPublicPrefixState)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput() ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput) Elem() ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) ExpressRouteCircuitPeeringAdvertisedPublicPrefixState {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitPeeringAdvertisedPublicPrefixState
		return ret
	}).(ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput)
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRouteCircuitPeeringAdvertisedPublicPrefixState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput() ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput
	ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutputWithContext(context.Context) ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput
}

var expressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrType = reflect.TypeOf((**ExpressRouteCircuitPeeringAdvertisedPublicPrefixState)(nil)).Elem()

type ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput() ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput
	ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutputWithContext(context.Context) ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput
}

type expressRouteCircuitPeeringAdvertisedPublicPrefixStatePtr string

func ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtr(v string) ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrInput {
	return (*expressRouteCircuitPeeringAdvertisedPublicPrefixStatePtr)(&v)
}

func (*expressRouteCircuitPeeringAdvertisedPublicPrefixStatePtr) ElementType() reflect.Type {
	return expressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrType
}

func (in *expressRouteCircuitPeeringAdvertisedPublicPrefixStatePtr) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput() ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput {
	return pulumi.ToOutput(in).(ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput)
}

func (in *expressRouteCircuitPeeringAdvertisedPublicPrefixStatePtr) ToExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput)
}

type ExpressRouteCircuitPeeringStateEnum string

const (
	ExpressRouteCircuitPeeringStateEnumDisabled = ExpressRouteCircuitPeeringStateEnum("Disabled")
	ExpressRouteCircuitPeeringStateEnumEnabled  = ExpressRouteCircuitPeeringStateEnum("Enabled")
)

func (ExpressRouteCircuitPeeringStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringStateEnum)(nil)).Elem()
}

func (e ExpressRouteCircuitPeeringStateEnum) ToExpressRouteCircuitPeeringStateEnumOutput() ExpressRouteCircuitPeeringStateEnumOutput {
	return pulumi.ToOutput(e).(ExpressRouteCircuitPeeringStateEnumOutput)
}

func (e ExpressRouteCircuitPeeringStateEnum) ToExpressRouteCircuitPeeringStateEnumOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRouteCircuitPeeringStateEnumOutput)
}

func (e ExpressRouteCircuitPeeringStateEnum) ToExpressRouteCircuitPeeringStateEnumPtrOutput() ExpressRouteCircuitPeeringStateEnumPtrOutput {
	return e.ToExpressRouteCircuitPeeringStateEnumPtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitPeeringStateEnum) ToExpressRouteCircuitPeeringStateEnumPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringStateEnumPtrOutput {
	return ExpressRouteCircuitPeeringStateEnum(e).ToExpressRouteCircuitPeeringStateEnumOutputWithContext(ctx).ToExpressRouteCircuitPeeringStateEnumPtrOutputWithContext(ctx)
}

func (e ExpressRouteCircuitPeeringStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitPeeringStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitPeeringStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitPeeringStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRouteCircuitPeeringStateEnumOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringStateEnum)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringStateEnumOutput) ToExpressRouteCircuitPeeringStateEnumOutput() ExpressRouteCircuitPeeringStateEnumOutput {
	return o
}

func (o ExpressRouteCircuitPeeringStateEnumOutput) ToExpressRouteCircuitPeeringStateEnumOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringStateEnumOutput {
	return o
}

func (o ExpressRouteCircuitPeeringStateEnumOutput) ToExpressRouteCircuitPeeringStateEnumPtrOutput() ExpressRouteCircuitPeeringStateEnumPtrOutput {
	return o.ToExpressRouteCircuitPeeringStateEnumPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringStateEnumOutput) ToExpressRouteCircuitPeeringStateEnumPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitPeeringStateEnum) *ExpressRouteCircuitPeeringStateEnum {
		return &v
	}).(ExpressRouteCircuitPeeringStateEnumPtrOutput)
}

func (o ExpressRouteCircuitPeeringStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitPeeringStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitPeeringStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitPeeringStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringStateEnumPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringStateEnum)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringStateEnumPtrOutput) ToExpressRouteCircuitPeeringStateEnumPtrOutput() ExpressRouteCircuitPeeringStateEnumPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringStateEnumPtrOutput) ToExpressRouteCircuitPeeringStateEnumPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringStateEnumPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringStateEnumPtrOutput) Elem() ExpressRouteCircuitPeeringStateEnumOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringStateEnum) ExpressRouteCircuitPeeringStateEnum {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitPeeringStateEnum
		return ret
	}).(ExpressRouteCircuitPeeringStateEnumOutput)
}

func (o ExpressRouteCircuitPeeringStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRouteCircuitPeeringStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRouteCircuitPeeringStateEnumInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringStateEnumOutput() ExpressRouteCircuitPeeringStateEnumOutput
	ToExpressRouteCircuitPeeringStateEnumOutputWithContext(context.Context) ExpressRouteCircuitPeeringStateEnumOutput
}

var expressRouteCircuitPeeringStateEnumPtrType = reflect.TypeOf((**ExpressRouteCircuitPeeringStateEnum)(nil)).Elem()

type ExpressRouteCircuitPeeringStateEnumPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringStateEnumPtrOutput() ExpressRouteCircuitPeeringStateEnumPtrOutput
	ToExpressRouteCircuitPeeringStateEnumPtrOutputWithContext(context.Context) ExpressRouteCircuitPeeringStateEnumPtrOutput
}

type expressRouteCircuitPeeringStateEnumPtr string

func ExpressRouteCircuitPeeringStateEnumPtr(v string) ExpressRouteCircuitPeeringStateEnumPtrInput {
	return (*expressRouteCircuitPeeringStateEnumPtr)(&v)
}

func (*expressRouteCircuitPeeringStateEnumPtr) ElementType() reflect.Type {
	return expressRouteCircuitPeeringStateEnumPtrType
}

func (in *expressRouteCircuitPeeringStateEnumPtr) ToExpressRouteCircuitPeeringStateEnumPtrOutput() ExpressRouteCircuitPeeringStateEnumPtrOutput {
	return pulumi.ToOutput(in).(ExpressRouteCircuitPeeringStateEnumPtrOutput)
}

func (in *expressRouteCircuitPeeringStateEnumPtr) ToExpressRouteCircuitPeeringStateEnumPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRouteCircuitPeeringStateEnumPtrOutput)
}

type ExpressRouteCircuitPeeringTypeEnum string

const (
	ExpressRouteCircuitPeeringTypeEnumAzurePublicPeering  = ExpressRouteCircuitPeeringTypeEnum("AzurePublicPeering")
	ExpressRouteCircuitPeeringTypeEnumAzurePrivatePeering = ExpressRouteCircuitPeeringTypeEnum("AzurePrivatePeering")
	ExpressRouteCircuitPeeringTypeEnumMicrosoftPeering    = ExpressRouteCircuitPeeringTypeEnum("MicrosoftPeering")
)

func (ExpressRouteCircuitPeeringTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringTypeEnum)(nil)).Elem()
}

func (e ExpressRouteCircuitPeeringTypeEnum) ToExpressRouteCircuitPeeringTypeEnumOutput() ExpressRouteCircuitPeeringTypeEnumOutput {
	return pulumi.ToOutput(e).(ExpressRouteCircuitPeeringTypeEnumOutput)
}

func (e ExpressRouteCircuitPeeringTypeEnum) ToExpressRouteCircuitPeeringTypeEnumOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRouteCircuitPeeringTypeEnumOutput)
}

func (e ExpressRouteCircuitPeeringTypeEnum) ToExpressRouteCircuitPeeringTypeEnumPtrOutput() ExpressRouteCircuitPeeringTypeEnumPtrOutput {
	return e.ToExpressRouteCircuitPeeringTypeEnumPtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitPeeringTypeEnum) ToExpressRouteCircuitPeeringTypeEnumPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeEnumPtrOutput {
	return ExpressRouteCircuitPeeringTypeEnum(e).ToExpressRouteCircuitPeeringTypeEnumOutputWithContext(ctx).ToExpressRouteCircuitPeeringTypeEnumPtrOutputWithContext(ctx)
}

func (e ExpressRouteCircuitPeeringTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitPeeringTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitPeeringTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitPeeringTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRouteCircuitPeeringTypeEnumOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringTypeEnum)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringTypeEnumOutput) ToExpressRouteCircuitPeeringTypeEnumOutput() ExpressRouteCircuitPeeringTypeEnumOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeEnumOutput) ToExpressRouteCircuitPeeringTypeEnumOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeEnumOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeEnumOutput) ToExpressRouteCircuitPeeringTypeEnumPtrOutput() ExpressRouteCircuitPeeringTypeEnumPtrOutput {
	return o.ToExpressRouteCircuitPeeringTypeEnumPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringTypeEnumOutput) ToExpressRouteCircuitPeeringTypeEnumPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitPeeringTypeEnum) *ExpressRouteCircuitPeeringTypeEnum {
		return &v
	}).(ExpressRouteCircuitPeeringTypeEnumPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitPeeringTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitPeeringTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitPeeringTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringTypeEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringTypeEnum)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringTypeEnumPtrOutput) ToExpressRouteCircuitPeeringTypeEnumPtrOutput() ExpressRouteCircuitPeeringTypeEnumPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeEnumPtrOutput) ToExpressRouteCircuitPeeringTypeEnumPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeEnumPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeEnumPtrOutput) Elem() ExpressRouteCircuitPeeringTypeEnumOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringTypeEnum) ExpressRouteCircuitPeeringTypeEnum {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitPeeringTypeEnum
		return ret
	}).(ExpressRouteCircuitPeeringTypeEnumOutput)
}

func (o ExpressRouteCircuitPeeringTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRouteCircuitPeeringTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRouteCircuitPeeringTypeEnumInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringTypeEnumOutput() ExpressRouteCircuitPeeringTypeEnumOutput
	ToExpressRouteCircuitPeeringTypeEnumOutputWithContext(context.Context) ExpressRouteCircuitPeeringTypeEnumOutput
}

var expressRouteCircuitPeeringTypeEnumPtrType = reflect.TypeOf((**ExpressRouteCircuitPeeringTypeEnum)(nil)).Elem()

type ExpressRouteCircuitPeeringTypeEnumPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringTypeEnumPtrOutput() ExpressRouteCircuitPeeringTypeEnumPtrOutput
	ToExpressRouteCircuitPeeringTypeEnumPtrOutputWithContext(context.Context) ExpressRouteCircuitPeeringTypeEnumPtrOutput
}

type expressRouteCircuitPeeringTypeEnumPtr string

func ExpressRouteCircuitPeeringTypeEnumPtr(v string) ExpressRouteCircuitPeeringTypeEnumPtrInput {
	return (*expressRouteCircuitPeeringTypeEnumPtr)(&v)
}

func (*expressRouteCircuitPeeringTypeEnumPtr) ElementType() reflect.Type {
	return expressRouteCircuitPeeringTypeEnumPtrType
}

func (in *expressRouteCircuitPeeringTypeEnumPtr) ToExpressRouteCircuitPeeringTypeEnumPtrOutput() ExpressRouteCircuitPeeringTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(ExpressRouteCircuitPeeringTypeEnumPtrOutput)
}

func (in *expressRouteCircuitPeeringTypeEnumPtr) ToExpressRouteCircuitPeeringTypeEnumPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRouteCircuitPeeringTypeEnumPtrOutput)
}

type ExpressRouteCircuitSkuFamily string

const (
	ExpressRouteCircuitSkuFamilyUnlimitedData = ExpressRouteCircuitSkuFamily("UnlimitedData")
	ExpressRouteCircuitSkuFamilyMeteredData   = ExpressRouteCircuitSkuFamily("MeteredData")
)

func (ExpressRouteCircuitSkuFamily) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSkuFamily)(nil)).Elem()
}

func (e ExpressRouteCircuitSkuFamily) ToExpressRouteCircuitSkuFamilyOutput() ExpressRouteCircuitSkuFamilyOutput {
	return pulumi.ToOutput(e).(ExpressRouteCircuitSkuFamilyOutput)
}

func (e ExpressRouteCircuitSkuFamily) ToExpressRouteCircuitSkuFamilyOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuFamilyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRouteCircuitSkuFamilyOutput)
}

func (e ExpressRouteCircuitSkuFamily) ToExpressRouteCircuitSkuFamilyPtrOutput() ExpressRouteCircuitSkuFamilyPtrOutput {
	return e.ToExpressRouteCircuitSkuFamilyPtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitSkuFamily) ToExpressRouteCircuitSkuFamilyPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuFamilyPtrOutput {
	return ExpressRouteCircuitSkuFamily(e).ToExpressRouteCircuitSkuFamilyOutputWithContext(ctx).ToExpressRouteCircuitSkuFamilyPtrOutputWithContext(ctx)
}

func (e ExpressRouteCircuitSkuFamily) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitSkuFamily) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitSkuFamily) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitSkuFamily) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRouteCircuitSkuFamilyOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuFamilyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSkuFamily)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuFamilyOutput) ToExpressRouteCircuitSkuFamilyOutput() ExpressRouteCircuitSkuFamilyOutput {
	return o
}

func (o ExpressRouteCircuitSkuFamilyOutput) ToExpressRouteCircuitSkuFamilyOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuFamilyOutput {
	return o
}

func (o ExpressRouteCircuitSkuFamilyOutput) ToExpressRouteCircuitSkuFamilyPtrOutput() ExpressRouteCircuitSkuFamilyPtrOutput {
	return o.ToExpressRouteCircuitSkuFamilyPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuFamilyOutput) ToExpressRouteCircuitSkuFamilyPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuFamilyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitSkuFamily) *ExpressRouteCircuitSkuFamily {
		return &v
	}).(ExpressRouteCircuitSkuFamilyPtrOutput)
}

func (o ExpressRouteCircuitSkuFamilyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuFamilyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitSkuFamily) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitSkuFamilyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuFamilyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitSkuFamily) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSkuFamilyPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuFamilyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitSkuFamily)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuFamilyPtrOutput) ToExpressRouteCircuitSkuFamilyPtrOutput() ExpressRouteCircuitSkuFamilyPtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuFamilyPtrOutput) ToExpressRouteCircuitSkuFamilyPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuFamilyPtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuFamilyPtrOutput) Elem() ExpressRouteCircuitSkuFamilyOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuFamily) ExpressRouteCircuitSkuFamily {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitSkuFamily
		return ret
	}).(ExpressRouteCircuitSkuFamilyOutput)
}

func (o ExpressRouteCircuitSkuFamilyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuFamilyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRouteCircuitSkuFamily) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRouteCircuitSkuFamilyInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuFamilyOutput() ExpressRouteCircuitSkuFamilyOutput
	ToExpressRouteCircuitSkuFamilyOutputWithContext(context.Context) ExpressRouteCircuitSkuFamilyOutput
}

var expressRouteCircuitSkuFamilyPtrType = reflect.TypeOf((**ExpressRouteCircuitSkuFamily)(nil)).Elem()

type ExpressRouteCircuitSkuFamilyPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuFamilyPtrOutput() ExpressRouteCircuitSkuFamilyPtrOutput
	ToExpressRouteCircuitSkuFamilyPtrOutputWithContext(context.Context) ExpressRouteCircuitSkuFamilyPtrOutput
}

type expressRouteCircuitSkuFamilyPtr string

func ExpressRouteCircuitSkuFamilyPtr(v string) ExpressRouteCircuitSkuFamilyPtrInput {
	return (*expressRouteCircuitSkuFamilyPtr)(&v)
}

func (*expressRouteCircuitSkuFamilyPtr) ElementType() reflect.Type {
	return expressRouteCircuitSkuFamilyPtrType
}

func (in *expressRouteCircuitSkuFamilyPtr) ToExpressRouteCircuitSkuFamilyPtrOutput() ExpressRouteCircuitSkuFamilyPtrOutput {
	return pulumi.ToOutput(in).(ExpressRouteCircuitSkuFamilyPtrOutput)
}

func (in *expressRouteCircuitSkuFamilyPtr) ToExpressRouteCircuitSkuFamilyPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuFamilyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRouteCircuitSkuFamilyPtrOutput)
}

type ExpressRouteCircuitSkuTier string

const (
	ExpressRouteCircuitSkuTierStandard = ExpressRouteCircuitSkuTier("Standard")
	ExpressRouteCircuitSkuTierPremium  = ExpressRouteCircuitSkuTier("Premium")
)

func (ExpressRouteCircuitSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSkuTier)(nil)).Elem()
}

func (e ExpressRouteCircuitSkuTier) ToExpressRouteCircuitSkuTierOutput() ExpressRouteCircuitSkuTierOutput {
	return pulumi.ToOutput(e).(ExpressRouteCircuitSkuTierOutput)
}

func (e ExpressRouteCircuitSkuTier) ToExpressRouteCircuitSkuTierOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRouteCircuitSkuTierOutput)
}

func (e ExpressRouteCircuitSkuTier) ToExpressRouteCircuitSkuTierPtrOutput() ExpressRouteCircuitSkuTierPtrOutput {
	return e.ToExpressRouteCircuitSkuTierPtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitSkuTier) ToExpressRouteCircuitSkuTierPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuTierPtrOutput {
	return ExpressRouteCircuitSkuTier(e).ToExpressRouteCircuitSkuTierOutputWithContext(ctx).ToExpressRouteCircuitSkuTierPtrOutputWithContext(ctx)
}

func (e ExpressRouteCircuitSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteCircuitSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRouteCircuitSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRouteCircuitSkuTierOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSkuTier)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuTierOutput) ToExpressRouteCircuitSkuTierOutput() ExpressRouteCircuitSkuTierOutput {
	return o
}

func (o ExpressRouteCircuitSkuTierOutput) ToExpressRouteCircuitSkuTierOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuTierOutput {
	return o
}

func (o ExpressRouteCircuitSkuTierOutput) ToExpressRouteCircuitSkuTierPtrOutput() ExpressRouteCircuitSkuTierPtrOutput {
	return o.ToExpressRouteCircuitSkuTierPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuTierOutput) ToExpressRouteCircuitSkuTierPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitSkuTier) *ExpressRouteCircuitSkuTier {
		return &v
	}).(ExpressRouteCircuitSkuTierPtrOutput)
}

func (o ExpressRouteCircuitSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteCircuitSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSkuTierPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitSkuTier)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuTierPtrOutput) ToExpressRouteCircuitSkuTierPtrOutput() ExpressRouteCircuitSkuTierPtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuTierPtrOutput) ToExpressRouteCircuitSkuTierPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuTierPtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuTierPtrOutput) Elem() ExpressRouteCircuitSkuTierOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuTier) ExpressRouteCircuitSkuTier {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitSkuTier
		return ret
	}).(ExpressRouteCircuitSkuTierOutput)
}

func (o ExpressRouteCircuitSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRouteCircuitSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRouteCircuitSkuTierInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuTierOutput() ExpressRouteCircuitSkuTierOutput
	ToExpressRouteCircuitSkuTierOutputWithContext(context.Context) ExpressRouteCircuitSkuTierOutput
}

var expressRouteCircuitSkuTierPtrType = reflect.TypeOf((**ExpressRouteCircuitSkuTier)(nil)).Elem()

type ExpressRouteCircuitSkuTierPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuTierPtrOutput() ExpressRouteCircuitSkuTierPtrOutput
	ToExpressRouteCircuitSkuTierPtrOutputWithContext(context.Context) ExpressRouteCircuitSkuTierPtrOutput
}

type expressRouteCircuitSkuTierPtr string

func ExpressRouteCircuitSkuTierPtr(v string) ExpressRouteCircuitSkuTierPtrInput {
	return (*expressRouteCircuitSkuTierPtr)(&v)
}

func (*expressRouteCircuitSkuTierPtr) ElementType() reflect.Type {
	return expressRouteCircuitSkuTierPtrType
}

func (in *expressRouteCircuitSkuTierPtr) ToExpressRouteCircuitSkuTierPtrOutput() ExpressRouteCircuitSkuTierPtrOutput {
	return pulumi.ToOutput(in).(ExpressRouteCircuitSkuTierPtrOutput)
}

func (in *expressRouteCircuitSkuTierPtr) ToExpressRouteCircuitSkuTierPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRouteCircuitSkuTierPtrOutput)
}

type IPAllocationMethod string

const (
	IPAllocationMethodStatic  = IPAllocationMethod("Static")
	IPAllocationMethodDynamic = IPAllocationMethod("Dynamic")
)

func (IPAllocationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAllocationMethod)(nil)).Elem()
}

func (e IPAllocationMethod) ToIPAllocationMethodOutput() IPAllocationMethodOutput {
	return pulumi.ToOutput(e).(IPAllocationMethodOutput)
}

func (e IPAllocationMethod) ToIPAllocationMethodOutputWithContext(ctx context.Context) IPAllocationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IPAllocationMethodOutput)
}

func (e IPAllocationMethod) ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput {
	return e.ToIPAllocationMethodPtrOutputWithContext(context.Background())
}

func (e IPAllocationMethod) ToIPAllocationMethodPtrOutputWithContext(ctx context.Context) IPAllocationMethodPtrOutput {
	return IPAllocationMethod(e).ToIPAllocationMethodOutputWithContext(ctx).ToIPAllocationMethodPtrOutputWithContext(ctx)
}

func (e IPAllocationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAllocationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAllocationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IPAllocationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IPAllocationMethodOutput struct{ *pulumi.OutputState }

func (IPAllocationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAllocationMethod)(nil)).Elem()
}

func (o IPAllocationMethodOutput) ToIPAllocationMethodOutput() IPAllocationMethodOutput {
	return o
}

func (o IPAllocationMethodOutput) ToIPAllocationMethodOutputWithContext(ctx context.Context) IPAllocationMethodOutput {
	return o
}

func (o IPAllocationMethodOutput) ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput {
	return o.ToIPAllocationMethodPtrOutputWithContext(context.Background())
}

func (o IPAllocationMethodOutput) ToIPAllocationMethodPtrOutputWithContext(ctx context.Context) IPAllocationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAllocationMethod) *IPAllocationMethod {
		return &v
	}).(IPAllocationMethodPtrOutput)
}

func (o IPAllocationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IPAllocationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAllocationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IPAllocationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPAllocationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAllocationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IPAllocationMethodPtrOutput struct{ *pulumi.OutputState }

func (IPAllocationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAllocationMethod)(nil)).Elem()
}

func (o IPAllocationMethodPtrOutput) ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput {
	return o
}

func (o IPAllocationMethodPtrOutput) ToIPAllocationMethodPtrOutputWithContext(ctx context.Context) IPAllocationMethodPtrOutput {
	return o
}

func (o IPAllocationMethodPtrOutput) Elem() IPAllocationMethodOutput {
	return o.ApplyT(func(v *IPAllocationMethod) IPAllocationMethod {
		if v != nil {
			return *v
		}
		var ret IPAllocationMethod
		return ret
	}).(IPAllocationMethodOutput)
}

func (o IPAllocationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPAllocationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IPAllocationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IPAllocationMethodInput interface {
	pulumi.Input

	ToIPAllocationMethodOutput() IPAllocationMethodOutput
	ToIPAllocationMethodOutputWithContext(context.Context) IPAllocationMethodOutput
}

var ipallocationMethodPtrType = reflect.TypeOf((**IPAllocationMethod)(nil)).Elem()

type IPAllocationMethodPtrInput interface {
	pulumi.Input

	ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput
	ToIPAllocationMethodPtrOutputWithContext(context.Context) IPAllocationMethodPtrOutput
}

type ipallocationMethodPtr string

func IPAllocationMethodPtr(v string) IPAllocationMethodPtrInput {
	return (*ipallocationMethodPtr)(&v)
}

func (*ipallocationMethodPtr) ElementType() reflect.Type {
	return ipallocationMethodPtrType
}

func (in *ipallocationMethodPtr) ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput {
	return pulumi.ToOutput(in).(IPAllocationMethodPtrOutput)
}

func (in *ipallocationMethodPtr) ToIPAllocationMethodPtrOutputWithContext(ctx context.Context) IPAllocationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IPAllocationMethodPtrOutput)
}

type IPVersion string

const (
	IPVersionIPv4 = IPVersion("IPv4")
	IPVersionIPv6 = IPVersion("IPv6")
)

func (IPVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*IPVersion)(nil)).Elem()
}

func (e IPVersion) ToIPVersionOutput() IPVersionOutput {
	return pulumi.ToOutput(e).(IPVersionOutput)
}

func (e IPVersion) ToIPVersionOutputWithContext(ctx context.Context) IPVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IPVersionOutput)
}

func (e IPVersion) ToIPVersionPtrOutput() IPVersionPtrOutput {
	return e.ToIPVersionPtrOutputWithContext(context.Background())
}

func (e IPVersion) ToIPVersionPtrOutputWithContext(ctx context.Context) IPVersionPtrOutput {
	return IPVersion(e).ToIPVersionOutputWithContext(ctx).ToIPVersionPtrOutputWithContext(ctx)
}

func (e IPVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IPVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IPVersionOutput struct{ *pulumi.OutputState }

func (IPVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPVersion)(nil)).Elem()
}

func (o IPVersionOutput) ToIPVersionOutput() IPVersionOutput {
	return o
}

func (o IPVersionOutput) ToIPVersionOutputWithContext(ctx context.Context) IPVersionOutput {
	return o
}

func (o IPVersionOutput) ToIPVersionPtrOutput() IPVersionPtrOutput {
	return o.ToIPVersionPtrOutputWithContext(context.Background())
}

func (o IPVersionOutput) ToIPVersionPtrOutputWithContext(ctx context.Context) IPVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPVersion) *IPVersion {
		return &v
	}).(IPVersionPtrOutput)
}

func (o IPVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IPVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IPVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IPVersionPtrOutput struct{ *pulumi.OutputState }

func (IPVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPVersion)(nil)).Elem()
}

func (o IPVersionPtrOutput) ToIPVersionPtrOutput() IPVersionPtrOutput {
	return o
}

func (o IPVersionPtrOutput) ToIPVersionPtrOutputWithContext(ctx context.Context) IPVersionPtrOutput {
	return o
}

func (o IPVersionPtrOutput) Elem() IPVersionOutput {
	return o.ApplyT(func(v *IPVersion) IPVersion {
		if v != nil {
			return *v
		}
		var ret IPVersion
		return ret
	}).(IPVersionOutput)
}

func (o IPVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IPVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IPVersionInput interface {
	pulumi.Input

	ToIPVersionOutput() IPVersionOutput
	ToIPVersionOutputWithContext(context.Context) IPVersionOutput
}

var ipversionPtrType = reflect.TypeOf((**IPVersion)(nil)).Elem()

type IPVersionPtrInput interface {
	pulumi.Input

	ToIPVersionPtrOutput() IPVersionPtrOutput
	ToIPVersionPtrOutputWithContext(context.Context) IPVersionPtrOutput
}

type ipversionPtr string

func IPVersionPtr(v string) IPVersionPtrInput {
	return (*ipversionPtr)(&v)
}

func (*ipversionPtr) ElementType() reflect.Type {
	return ipversionPtrType
}

func (in *ipversionPtr) ToIPVersionPtrOutput() IPVersionPtrOutput {
	return pulumi.ToOutput(in).(IPVersionPtrOutput)
}

func (in *ipversionPtr) ToIPVersionPtrOutputWithContext(ctx context.Context) IPVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IPVersionPtrOutput)
}

type IkeEncryption string

const (
	IkeEncryptionDES    = IkeEncryption("DES")
	IkeEncryptionDES3   = IkeEncryption("DES3")
	IkeEncryptionAES128 = IkeEncryption("AES128")
	IkeEncryptionAES192 = IkeEncryption("AES192")
	IkeEncryptionAES256 = IkeEncryption("AES256")
)

func (IkeEncryption) ElementType() reflect.Type {
	return reflect.TypeOf((*IkeEncryption)(nil)).Elem()
}

func (e IkeEncryption) ToIkeEncryptionOutput() IkeEncryptionOutput {
	return pulumi.ToOutput(e).(IkeEncryptionOutput)
}

func (e IkeEncryption) ToIkeEncryptionOutputWithContext(ctx context.Context) IkeEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IkeEncryptionOutput)
}

func (e IkeEncryption) ToIkeEncryptionPtrOutput() IkeEncryptionPtrOutput {
	return e.ToIkeEncryptionPtrOutputWithContext(context.Background())
}

func (e IkeEncryption) ToIkeEncryptionPtrOutputWithContext(ctx context.Context) IkeEncryptionPtrOutput {
	return IkeEncryption(e).ToIkeEncryptionOutputWithContext(ctx).ToIkeEncryptionPtrOutputWithContext(ctx)
}

func (e IkeEncryption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IkeEncryption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IkeEncryption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IkeEncryption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IkeEncryptionOutput struct{ *pulumi.OutputState }

func (IkeEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IkeEncryption)(nil)).Elem()
}

func (o IkeEncryptionOutput) ToIkeEncryptionOutput() IkeEncryptionOutput {
	return o
}

func (o IkeEncryptionOutput) ToIkeEncryptionOutputWithContext(ctx context.Context) IkeEncryptionOutput {
	return o
}

func (o IkeEncryptionOutput) ToIkeEncryptionPtrOutput() IkeEncryptionPtrOutput {
	return o.ToIkeEncryptionPtrOutputWithContext(context.Background())
}

func (o IkeEncryptionOutput) ToIkeEncryptionPtrOutputWithContext(ctx context.Context) IkeEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IkeEncryption) *IkeEncryption {
		return &v
	}).(IkeEncryptionPtrOutput)
}

func (o IkeEncryptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IkeEncryptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IkeEncryption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IkeEncryptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IkeEncryptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IkeEncryption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IkeEncryptionPtrOutput struct{ *pulumi.OutputState }

func (IkeEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IkeEncryption)(nil)).Elem()
}

func (o IkeEncryptionPtrOutput) ToIkeEncryptionPtrOutput() IkeEncryptionPtrOutput {
	return o
}

func (o IkeEncryptionPtrOutput) ToIkeEncryptionPtrOutputWithContext(ctx context.Context) IkeEncryptionPtrOutput {
	return o
}

func (o IkeEncryptionPtrOutput) Elem() IkeEncryptionOutput {
	return o.ApplyT(func(v *IkeEncryption) IkeEncryption {
		if v != nil {
			return *v
		}
		var ret IkeEncryption
		return ret
	}).(IkeEncryptionOutput)
}

func (o IkeEncryptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IkeEncryptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IkeEncryption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IkeEncryptionInput interface {
	pulumi.Input

	ToIkeEncryptionOutput() IkeEncryptionOutput
	ToIkeEncryptionOutputWithContext(context.Context) IkeEncryptionOutput
}

var ikeEncryptionPtrType = reflect.TypeOf((**IkeEncryption)(nil)).Elem()

type IkeEncryptionPtrInput interface {
	pulumi.Input

	ToIkeEncryptionPtrOutput() IkeEncryptionPtrOutput
	ToIkeEncryptionPtrOutputWithContext(context.Context) IkeEncryptionPtrOutput
}

type ikeEncryptionPtr string

func IkeEncryptionPtr(v string) IkeEncryptionPtrInput {
	return (*ikeEncryptionPtr)(&v)
}

func (*ikeEncryptionPtr) ElementType() reflect.Type {
	return ikeEncryptionPtrType
}

func (in *ikeEncryptionPtr) ToIkeEncryptionPtrOutput() IkeEncryptionPtrOutput {
	return pulumi.ToOutput(in).(IkeEncryptionPtrOutput)
}

func (in *ikeEncryptionPtr) ToIkeEncryptionPtrOutputWithContext(ctx context.Context) IkeEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IkeEncryptionPtrOutput)
}

type IkeIntegrity string

const (
	IkeIntegrityMD5    = IkeIntegrity("MD5")
	IkeIntegritySHA1   = IkeIntegrity("SHA1")
	IkeIntegritySHA256 = IkeIntegrity("SHA256")
	IkeIntegritySHA384 = IkeIntegrity("SHA384")
)

func (IkeIntegrity) ElementType() reflect.Type {
	return reflect.TypeOf((*IkeIntegrity)(nil)).Elem()
}

func (e IkeIntegrity) ToIkeIntegrityOutput() IkeIntegrityOutput {
	return pulumi.ToOutput(e).(IkeIntegrityOutput)
}

func (e IkeIntegrity) ToIkeIntegrityOutputWithContext(ctx context.Context) IkeIntegrityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IkeIntegrityOutput)
}

func (e IkeIntegrity) ToIkeIntegrityPtrOutput() IkeIntegrityPtrOutput {
	return e.ToIkeIntegrityPtrOutputWithContext(context.Background())
}

func (e IkeIntegrity) ToIkeIntegrityPtrOutputWithContext(ctx context.Context) IkeIntegrityPtrOutput {
	return IkeIntegrity(e).ToIkeIntegrityOutputWithContext(ctx).ToIkeIntegrityPtrOutputWithContext(ctx)
}

func (e IkeIntegrity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IkeIntegrity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IkeIntegrity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IkeIntegrity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IkeIntegrityOutput struct{ *pulumi.OutputState }

func (IkeIntegrityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IkeIntegrity)(nil)).Elem()
}

func (o IkeIntegrityOutput) ToIkeIntegrityOutput() IkeIntegrityOutput {
	return o
}

func (o IkeIntegrityOutput) ToIkeIntegrityOutputWithContext(ctx context.Context) IkeIntegrityOutput {
	return o
}

func (o IkeIntegrityOutput) ToIkeIntegrityPtrOutput() IkeIntegrityPtrOutput {
	return o.ToIkeIntegrityPtrOutputWithContext(context.Background())
}

func (o IkeIntegrityOutput) ToIkeIntegrityPtrOutputWithContext(ctx context.Context) IkeIntegrityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IkeIntegrity) *IkeIntegrity {
		return &v
	}).(IkeIntegrityPtrOutput)
}

func (o IkeIntegrityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IkeIntegrityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IkeIntegrity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IkeIntegrityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IkeIntegrityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IkeIntegrity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IkeIntegrityPtrOutput struct{ *pulumi.OutputState }

func (IkeIntegrityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IkeIntegrity)(nil)).Elem()
}

func (o IkeIntegrityPtrOutput) ToIkeIntegrityPtrOutput() IkeIntegrityPtrOutput {
	return o
}

func (o IkeIntegrityPtrOutput) ToIkeIntegrityPtrOutputWithContext(ctx context.Context) IkeIntegrityPtrOutput {
	return o
}

func (o IkeIntegrityPtrOutput) Elem() IkeIntegrityOutput {
	return o.ApplyT(func(v *IkeIntegrity) IkeIntegrity {
		if v != nil {
			return *v
		}
		var ret IkeIntegrity
		return ret
	}).(IkeIntegrityOutput)
}

func (o IkeIntegrityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IkeIntegrityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IkeIntegrity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IkeIntegrityInput interface {
	pulumi.Input

	ToIkeIntegrityOutput() IkeIntegrityOutput
	ToIkeIntegrityOutputWithContext(context.Context) IkeIntegrityOutput
}

var ikeIntegrityPtrType = reflect.TypeOf((**IkeIntegrity)(nil)).Elem()

type IkeIntegrityPtrInput interface {
	pulumi.Input

	ToIkeIntegrityPtrOutput() IkeIntegrityPtrOutput
	ToIkeIntegrityPtrOutputWithContext(context.Context) IkeIntegrityPtrOutput
}

type ikeIntegrityPtr string

func IkeIntegrityPtr(v string) IkeIntegrityPtrInput {
	return (*ikeIntegrityPtr)(&v)
}

func (*ikeIntegrityPtr) ElementType() reflect.Type {
	return ikeIntegrityPtrType
}

func (in *ikeIntegrityPtr) ToIkeIntegrityPtrOutput() IkeIntegrityPtrOutput {
	return pulumi.ToOutput(in).(IkeIntegrityPtrOutput)
}

func (in *ikeIntegrityPtr) ToIkeIntegrityPtrOutputWithContext(ctx context.Context) IkeIntegrityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IkeIntegrityPtrOutput)
}

type IpsecEncryption string

const (
	IpsecEncryptionNone      = IpsecEncryption("None")
	IpsecEncryptionDES       = IpsecEncryption("DES")
	IpsecEncryptionDES3      = IpsecEncryption("DES3")
	IpsecEncryptionAES128    = IpsecEncryption("AES128")
	IpsecEncryptionAES192    = IpsecEncryption("AES192")
	IpsecEncryptionAES256    = IpsecEncryption("AES256")
	IpsecEncryptionGCMAES128 = IpsecEncryption("GCMAES128")
	IpsecEncryptionGCMAES192 = IpsecEncryption("GCMAES192")
	IpsecEncryptionGCMAES256 = IpsecEncryption("GCMAES256")
)

func (IpsecEncryption) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecEncryption)(nil)).Elem()
}

func (e IpsecEncryption) ToIpsecEncryptionOutput() IpsecEncryptionOutput {
	return pulumi.ToOutput(e).(IpsecEncryptionOutput)
}

func (e IpsecEncryption) ToIpsecEncryptionOutputWithContext(ctx context.Context) IpsecEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpsecEncryptionOutput)
}

func (e IpsecEncryption) ToIpsecEncryptionPtrOutput() IpsecEncryptionPtrOutput {
	return e.ToIpsecEncryptionPtrOutputWithContext(context.Background())
}

func (e IpsecEncryption) ToIpsecEncryptionPtrOutputWithContext(ctx context.Context) IpsecEncryptionPtrOutput {
	return IpsecEncryption(e).ToIpsecEncryptionOutputWithContext(ctx).ToIpsecEncryptionPtrOutputWithContext(ctx)
}

func (e IpsecEncryption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpsecEncryption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpsecEncryption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpsecEncryption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpsecEncryptionOutput struct{ *pulumi.OutputState }

func (IpsecEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecEncryption)(nil)).Elem()
}

func (o IpsecEncryptionOutput) ToIpsecEncryptionOutput() IpsecEncryptionOutput {
	return o
}

func (o IpsecEncryptionOutput) ToIpsecEncryptionOutputWithContext(ctx context.Context) IpsecEncryptionOutput {
	return o
}

func (o IpsecEncryptionOutput) ToIpsecEncryptionPtrOutput() IpsecEncryptionPtrOutput {
	return o.ToIpsecEncryptionPtrOutputWithContext(context.Background())
}

func (o IpsecEncryptionOutput) ToIpsecEncryptionPtrOutputWithContext(ctx context.Context) IpsecEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpsecEncryption) *IpsecEncryption {
		return &v
	}).(IpsecEncryptionPtrOutput)
}

func (o IpsecEncryptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpsecEncryptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpsecEncryption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpsecEncryptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpsecEncryptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpsecEncryption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpsecEncryptionPtrOutput struct{ *pulumi.OutputState }

func (IpsecEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsecEncryption)(nil)).Elem()
}

func (o IpsecEncryptionPtrOutput) ToIpsecEncryptionPtrOutput() IpsecEncryptionPtrOutput {
	return o
}

func (o IpsecEncryptionPtrOutput) ToIpsecEncryptionPtrOutputWithContext(ctx context.Context) IpsecEncryptionPtrOutput {
	return o
}

func (o IpsecEncryptionPtrOutput) Elem() IpsecEncryptionOutput {
	return o.ApplyT(func(v *IpsecEncryption) IpsecEncryption {
		if v != nil {
			return *v
		}
		var ret IpsecEncryption
		return ret
	}).(IpsecEncryptionOutput)
}

func (o IpsecEncryptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpsecEncryptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpsecEncryption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IpsecEncryptionInput interface {
	pulumi.Input

	ToIpsecEncryptionOutput() IpsecEncryptionOutput
	ToIpsecEncryptionOutputWithContext(context.Context) IpsecEncryptionOutput
}

var ipsecEncryptionPtrType = reflect.TypeOf((**IpsecEncryption)(nil)).Elem()

type IpsecEncryptionPtrInput interface {
	pulumi.Input

	ToIpsecEncryptionPtrOutput() IpsecEncryptionPtrOutput
	ToIpsecEncryptionPtrOutputWithContext(context.Context) IpsecEncryptionPtrOutput
}

type ipsecEncryptionPtr string

func IpsecEncryptionPtr(v string) IpsecEncryptionPtrInput {
	return (*ipsecEncryptionPtr)(&v)
}

func (*ipsecEncryptionPtr) ElementType() reflect.Type {
	return ipsecEncryptionPtrType
}

func (in *ipsecEncryptionPtr) ToIpsecEncryptionPtrOutput() IpsecEncryptionPtrOutput {
	return pulumi.ToOutput(in).(IpsecEncryptionPtrOutput)
}

func (in *ipsecEncryptionPtr) ToIpsecEncryptionPtrOutputWithContext(ctx context.Context) IpsecEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpsecEncryptionPtrOutput)
}

type IpsecIntegrity string

const (
	IpsecIntegrityMD5       = IpsecIntegrity("MD5")
	IpsecIntegritySHA1      = IpsecIntegrity("SHA1")
	IpsecIntegritySHA256    = IpsecIntegrity("SHA256")
	IpsecIntegrityGCMAES128 = IpsecIntegrity("GCMAES128")
	IpsecIntegrityGCMAES192 = IpsecIntegrity("GCMAES192")
	IpsecIntegrityGCMAES256 = IpsecIntegrity("GCMAES256")
)

func (IpsecIntegrity) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecIntegrity)(nil)).Elem()
}

func (e IpsecIntegrity) ToIpsecIntegrityOutput() IpsecIntegrityOutput {
	return pulumi.ToOutput(e).(IpsecIntegrityOutput)
}

func (e IpsecIntegrity) ToIpsecIntegrityOutputWithContext(ctx context.Context) IpsecIntegrityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpsecIntegrityOutput)
}

func (e IpsecIntegrity) ToIpsecIntegrityPtrOutput() IpsecIntegrityPtrOutput {
	return e.ToIpsecIntegrityPtrOutputWithContext(context.Background())
}

func (e IpsecIntegrity) ToIpsecIntegrityPtrOutputWithContext(ctx context.Context) IpsecIntegrityPtrOutput {
	return IpsecIntegrity(e).ToIpsecIntegrityOutputWithContext(ctx).ToIpsecIntegrityPtrOutputWithContext(ctx)
}

func (e IpsecIntegrity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpsecIntegrity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpsecIntegrity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpsecIntegrity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpsecIntegrityOutput struct{ *pulumi.OutputState }

func (IpsecIntegrityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecIntegrity)(nil)).Elem()
}

func (o IpsecIntegrityOutput) ToIpsecIntegrityOutput() IpsecIntegrityOutput {
	return o
}

func (o IpsecIntegrityOutput) ToIpsecIntegrityOutputWithContext(ctx context.Context) IpsecIntegrityOutput {
	return o
}

func (o IpsecIntegrityOutput) ToIpsecIntegrityPtrOutput() IpsecIntegrityPtrOutput {
	return o.ToIpsecIntegrityPtrOutputWithContext(context.Background())
}

func (o IpsecIntegrityOutput) ToIpsecIntegrityPtrOutputWithContext(ctx context.Context) IpsecIntegrityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpsecIntegrity) *IpsecIntegrity {
		return &v
	}).(IpsecIntegrityPtrOutput)
}

func (o IpsecIntegrityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpsecIntegrityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpsecIntegrity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpsecIntegrityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpsecIntegrityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpsecIntegrity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpsecIntegrityPtrOutput struct{ *pulumi.OutputState }

func (IpsecIntegrityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsecIntegrity)(nil)).Elem()
}

func (o IpsecIntegrityPtrOutput) ToIpsecIntegrityPtrOutput() IpsecIntegrityPtrOutput {
	return o
}

func (o IpsecIntegrityPtrOutput) ToIpsecIntegrityPtrOutputWithContext(ctx context.Context) IpsecIntegrityPtrOutput {
	return o
}

func (o IpsecIntegrityPtrOutput) Elem() IpsecIntegrityOutput {
	return o.ApplyT(func(v *IpsecIntegrity) IpsecIntegrity {
		if v != nil {
			return *v
		}
		var ret IpsecIntegrity
		return ret
	}).(IpsecIntegrityOutput)
}

func (o IpsecIntegrityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpsecIntegrityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpsecIntegrity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IpsecIntegrityInput interface {
	pulumi.Input

	ToIpsecIntegrityOutput() IpsecIntegrityOutput
	ToIpsecIntegrityOutputWithContext(context.Context) IpsecIntegrityOutput
}

var ipsecIntegrityPtrType = reflect.TypeOf((**IpsecIntegrity)(nil)).Elem()

type IpsecIntegrityPtrInput interface {
	pulumi.Input

	ToIpsecIntegrityPtrOutput() IpsecIntegrityPtrOutput
	ToIpsecIntegrityPtrOutputWithContext(context.Context) IpsecIntegrityPtrOutput
}

type ipsecIntegrityPtr string

func IpsecIntegrityPtr(v string) IpsecIntegrityPtrInput {
	return (*ipsecIntegrityPtr)(&v)
}

func (*ipsecIntegrityPtr) ElementType() reflect.Type {
	return ipsecIntegrityPtrType
}

func (in *ipsecIntegrityPtr) ToIpsecIntegrityPtrOutput() IpsecIntegrityPtrOutput {
	return pulumi.ToOutput(in).(IpsecIntegrityPtrOutput)
}

func (in *ipsecIntegrityPtr) ToIpsecIntegrityPtrOutputWithContext(ctx context.Context) IpsecIntegrityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpsecIntegrityPtrOutput)
}

type LoadBalancerSkuName string

const (
	LoadBalancerSkuNameBasic    = LoadBalancerSkuName("Basic")
	LoadBalancerSkuNameStandard = LoadBalancerSkuName("Standard")
)

func (LoadBalancerSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSkuName)(nil)).Elem()
}

func (e LoadBalancerSkuName) ToLoadBalancerSkuNameOutput() LoadBalancerSkuNameOutput {
	return pulumi.ToOutput(e).(LoadBalancerSkuNameOutput)
}

func (e LoadBalancerSkuName) ToLoadBalancerSkuNameOutputWithContext(ctx context.Context) LoadBalancerSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoadBalancerSkuNameOutput)
}

func (e LoadBalancerSkuName) ToLoadBalancerSkuNamePtrOutput() LoadBalancerSkuNamePtrOutput {
	return e.ToLoadBalancerSkuNamePtrOutputWithContext(context.Background())
}

func (e LoadBalancerSkuName) ToLoadBalancerSkuNamePtrOutputWithContext(ctx context.Context) LoadBalancerSkuNamePtrOutput {
	return LoadBalancerSkuName(e).ToLoadBalancerSkuNameOutputWithContext(ctx).ToLoadBalancerSkuNamePtrOutputWithContext(ctx)
}

func (e LoadBalancerSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoadBalancerSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoadBalancerSkuNameOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSkuName)(nil)).Elem()
}

func (o LoadBalancerSkuNameOutput) ToLoadBalancerSkuNameOutput() LoadBalancerSkuNameOutput {
	return o
}

func (o LoadBalancerSkuNameOutput) ToLoadBalancerSkuNameOutputWithContext(ctx context.Context) LoadBalancerSkuNameOutput {
	return o
}

func (o LoadBalancerSkuNameOutput) ToLoadBalancerSkuNamePtrOutput() LoadBalancerSkuNamePtrOutput {
	return o.ToLoadBalancerSkuNamePtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuNameOutput) ToLoadBalancerSkuNamePtrOutputWithContext(ctx context.Context) LoadBalancerSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerSkuName) *LoadBalancerSkuName {
		return &v
	}).(LoadBalancerSkuNamePtrOutput)
}

func (o LoadBalancerSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoadBalancerSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoadBalancerSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerSkuNamePtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerSkuName)(nil)).Elem()
}

func (o LoadBalancerSkuNamePtrOutput) ToLoadBalancerSkuNamePtrOutput() LoadBalancerSkuNamePtrOutput {
	return o
}

func (o LoadBalancerSkuNamePtrOutput) ToLoadBalancerSkuNamePtrOutputWithContext(ctx context.Context) LoadBalancerSkuNamePtrOutput {
	return o
}

func (o LoadBalancerSkuNamePtrOutput) Elem() LoadBalancerSkuNameOutput {
	return o.ApplyT(func(v *LoadBalancerSkuName) LoadBalancerSkuName {
		if v != nil {
			return *v
		}
		var ret LoadBalancerSkuName
		return ret
	}).(LoadBalancerSkuNameOutput)
}

func (o LoadBalancerSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoadBalancerSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LoadBalancerSkuNameInput interface {
	pulumi.Input

	ToLoadBalancerSkuNameOutput() LoadBalancerSkuNameOutput
	ToLoadBalancerSkuNameOutputWithContext(context.Context) LoadBalancerSkuNameOutput
}

var loadBalancerSkuNamePtrType = reflect.TypeOf((**LoadBalancerSkuName)(nil)).Elem()

type LoadBalancerSkuNamePtrInput interface {
	pulumi.Input

	ToLoadBalancerSkuNamePtrOutput() LoadBalancerSkuNamePtrOutput
	ToLoadBalancerSkuNamePtrOutputWithContext(context.Context) LoadBalancerSkuNamePtrOutput
}

type loadBalancerSkuNamePtr string

func LoadBalancerSkuNamePtr(v string) LoadBalancerSkuNamePtrInput {
	return (*loadBalancerSkuNamePtr)(&v)
}

func (*loadBalancerSkuNamePtr) ElementType() reflect.Type {
	return loadBalancerSkuNamePtrType
}

func (in *loadBalancerSkuNamePtr) ToLoadBalancerSkuNamePtrOutput() LoadBalancerSkuNamePtrOutput {
	return pulumi.ToOutput(in).(LoadBalancerSkuNamePtrOutput)
}

func (in *loadBalancerSkuNamePtr) ToLoadBalancerSkuNamePtrOutputWithContext(ctx context.Context) LoadBalancerSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoadBalancerSkuNamePtrOutput)
}

type LoadDistribution string

const (
	LoadDistributionDefault          = LoadDistribution("Default")
	LoadDistributionSourceIP         = LoadDistribution("SourceIP")
	LoadDistributionSourceIPProtocol = LoadDistribution("SourceIPProtocol")
)

func (LoadDistribution) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadDistribution)(nil)).Elem()
}

func (e LoadDistribution) ToLoadDistributionOutput() LoadDistributionOutput {
	return pulumi.ToOutput(e).(LoadDistributionOutput)
}

func (e LoadDistribution) ToLoadDistributionOutputWithContext(ctx context.Context) LoadDistributionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoadDistributionOutput)
}

func (e LoadDistribution) ToLoadDistributionPtrOutput() LoadDistributionPtrOutput {
	return e.ToLoadDistributionPtrOutputWithContext(context.Background())
}

func (e LoadDistribution) ToLoadDistributionPtrOutputWithContext(ctx context.Context) LoadDistributionPtrOutput {
	return LoadDistribution(e).ToLoadDistributionOutputWithContext(ctx).ToLoadDistributionPtrOutputWithContext(ctx)
}

func (e LoadDistribution) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadDistribution) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadDistribution) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoadDistribution) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoadDistributionOutput struct{ *pulumi.OutputState }

func (LoadDistributionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadDistribution)(nil)).Elem()
}

func (o LoadDistributionOutput) ToLoadDistributionOutput() LoadDistributionOutput {
	return o
}

func (o LoadDistributionOutput) ToLoadDistributionOutputWithContext(ctx context.Context) LoadDistributionOutput {
	return o
}

func (o LoadDistributionOutput) ToLoadDistributionPtrOutput() LoadDistributionPtrOutput {
	return o.ToLoadDistributionPtrOutputWithContext(context.Background())
}

func (o LoadDistributionOutput) ToLoadDistributionPtrOutputWithContext(ctx context.Context) LoadDistributionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadDistribution) *LoadDistribution {
		return &v
	}).(LoadDistributionPtrOutput)
}

func (o LoadDistributionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoadDistributionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadDistribution) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoadDistributionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadDistributionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadDistribution) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoadDistributionPtrOutput struct{ *pulumi.OutputState }

func (LoadDistributionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadDistribution)(nil)).Elem()
}

func (o LoadDistributionPtrOutput) ToLoadDistributionPtrOutput() LoadDistributionPtrOutput {
	return o
}

func (o LoadDistributionPtrOutput) ToLoadDistributionPtrOutputWithContext(ctx context.Context) LoadDistributionPtrOutput {
	return o
}

func (o LoadDistributionPtrOutput) Elem() LoadDistributionOutput {
	return o.ApplyT(func(v *LoadDistribution) LoadDistribution {
		if v != nil {
			return *v
		}
		var ret LoadDistribution
		return ret
	}).(LoadDistributionOutput)
}

func (o LoadDistributionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadDistributionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoadDistribution) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LoadDistributionInput interface {
	pulumi.Input

	ToLoadDistributionOutput() LoadDistributionOutput
	ToLoadDistributionOutputWithContext(context.Context) LoadDistributionOutput
}

var loadDistributionPtrType = reflect.TypeOf((**LoadDistribution)(nil)).Elem()

type LoadDistributionPtrInput interface {
	pulumi.Input

	ToLoadDistributionPtrOutput() LoadDistributionPtrOutput
	ToLoadDistributionPtrOutputWithContext(context.Context) LoadDistributionPtrOutput
}

type loadDistributionPtr string

func LoadDistributionPtr(v string) LoadDistributionPtrInput {
	return (*loadDistributionPtr)(&v)
}

func (*loadDistributionPtr) ElementType() reflect.Type {
	return loadDistributionPtrType
}

func (in *loadDistributionPtr) ToLoadDistributionPtrOutput() LoadDistributionPtrOutput {
	return pulumi.ToOutput(in).(LoadDistributionPtrOutput)
}

func (in *loadDistributionPtr) ToLoadDistributionPtrOutputWithContext(ctx context.Context) LoadDistributionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoadDistributionPtrOutput)
}

type PcProtocol string

const (
	PcProtocolTCP = PcProtocol("TCP")
	PcProtocolUDP = PcProtocol("UDP")
	PcProtocolAny = PcProtocol("Any")
)

func (PcProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*PcProtocol)(nil)).Elem()
}

func (e PcProtocol) ToPcProtocolOutput() PcProtocolOutput {
	return pulumi.ToOutput(e).(PcProtocolOutput)
}

func (e PcProtocol) ToPcProtocolOutputWithContext(ctx context.Context) PcProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PcProtocolOutput)
}

func (e PcProtocol) ToPcProtocolPtrOutput() PcProtocolPtrOutput {
	return e.ToPcProtocolPtrOutputWithContext(context.Background())
}

func (e PcProtocol) ToPcProtocolPtrOutputWithContext(ctx context.Context) PcProtocolPtrOutput {
	return PcProtocol(e).ToPcProtocolOutputWithContext(ctx).ToPcProtocolPtrOutputWithContext(ctx)
}

func (e PcProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PcProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PcProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PcProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PcProtocolOutput struct{ *pulumi.OutputState }

func (PcProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PcProtocol)(nil)).Elem()
}

func (o PcProtocolOutput) ToPcProtocolOutput() PcProtocolOutput {
	return o
}

func (o PcProtocolOutput) ToPcProtocolOutputWithContext(ctx context.Context) PcProtocolOutput {
	return o
}

func (o PcProtocolOutput) ToPcProtocolPtrOutput() PcProtocolPtrOutput {
	return o.ToPcProtocolPtrOutputWithContext(context.Background())
}

func (o PcProtocolOutput) ToPcProtocolPtrOutputWithContext(ctx context.Context) PcProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PcProtocol) *PcProtocol {
		return &v
	}).(PcProtocolPtrOutput)
}

func (o PcProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PcProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PcProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PcProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PcProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PcProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PcProtocolPtrOutput struct{ *pulumi.OutputState }

func (PcProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PcProtocol)(nil)).Elem()
}

func (o PcProtocolPtrOutput) ToPcProtocolPtrOutput() PcProtocolPtrOutput {
	return o
}

func (o PcProtocolPtrOutput) ToPcProtocolPtrOutputWithContext(ctx context.Context) PcProtocolPtrOutput {
	return o
}

func (o PcProtocolPtrOutput) Elem() PcProtocolOutput {
	return o.ApplyT(func(v *PcProtocol) PcProtocol {
		if v != nil {
			return *v
		}
		var ret PcProtocol
		return ret
	}).(PcProtocolOutput)
}

func (o PcProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PcProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PcProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PcProtocolInput interface {
	pulumi.Input

	ToPcProtocolOutput() PcProtocolOutput
	ToPcProtocolOutputWithContext(context.Context) PcProtocolOutput
}

var pcProtocolPtrType = reflect.TypeOf((**PcProtocol)(nil)).Elem()

type PcProtocolPtrInput interface {
	pulumi.Input

	ToPcProtocolPtrOutput() PcProtocolPtrOutput
	ToPcProtocolPtrOutputWithContext(context.Context) PcProtocolPtrOutput
}

type pcProtocolPtr string

func PcProtocolPtr(v string) PcProtocolPtrInput {
	return (*pcProtocolPtr)(&v)
}

func (*pcProtocolPtr) ElementType() reflect.Type {
	return pcProtocolPtrType
}

func (in *pcProtocolPtr) ToPcProtocolPtrOutput() PcProtocolPtrOutput {
	return pulumi.ToOutput(in).(PcProtocolPtrOutput)
}

func (in *pcProtocolPtr) ToPcProtocolPtrOutputWithContext(ctx context.Context) PcProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PcProtocolPtrOutput)
}

type PfsGroup string

const (
	PfsGroupNone    = PfsGroup("None")
	PfsGroupPFS1    = PfsGroup("PFS1")
	PfsGroupPFS2    = PfsGroup("PFS2")
	PfsGroupPFS2048 = PfsGroup("PFS2048")
	PfsGroupECP256  = PfsGroup("ECP256")
	PfsGroupECP384  = PfsGroup("ECP384")
	PfsGroupPFS24   = PfsGroup("PFS24")
)

func (PfsGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*PfsGroup)(nil)).Elem()
}

func (e PfsGroup) ToPfsGroupOutput() PfsGroupOutput {
	return pulumi.ToOutput(e).(PfsGroupOutput)
}

func (e PfsGroup) ToPfsGroupOutputWithContext(ctx context.Context) PfsGroupOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PfsGroupOutput)
}

func (e PfsGroup) ToPfsGroupPtrOutput() PfsGroupPtrOutput {
	return e.ToPfsGroupPtrOutputWithContext(context.Background())
}

func (e PfsGroup) ToPfsGroupPtrOutputWithContext(ctx context.Context) PfsGroupPtrOutput {
	return PfsGroup(e).ToPfsGroupOutputWithContext(ctx).ToPfsGroupPtrOutputWithContext(ctx)
}

func (e PfsGroup) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PfsGroup) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PfsGroup) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PfsGroup) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PfsGroupOutput struct{ *pulumi.OutputState }

func (PfsGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PfsGroup)(nil)).Elem()
}

func (o PfsGroupOutput) ToPfsGroupOutput() PfsGroupOutput {
	return o
}

func (o PfsGroupOutput) ToPfsGroupOutputWithContext(ctx context.Context) PfsGroupOutput {
	return o
}

func (o PfsGroupOutput) ToPfsGroupPtrOutput() PfsGroupPtrOutput {
	return o.ToPfsGroupPtrOutputWithContext(context.Background())
}

func (o PfsGroupOutput) ToPfsGroupPtrOutputWithContext(ctx context.Context) PfsGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PfsGroup) *PfsGroup {
		return &v
	}).(PfsGroupPtrOutput)
}

func (o PfsGroupOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PfsGroupOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PfsGroup) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PfsGroupOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PfsGroupOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PfsGroup) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PfsGroupPtrOutput struct{ *pulumi.OutputState }

func (PfsGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PfsGroup)(nil)).Elem()
}

func (o PfsGroupPtrOutput) ToPfsGroupPtrOutput() PfsGroupPtrOutput {
	return o
}

func (o PfsGroupPtrOutput) ToPfsGroupPtrOutputWithContext(ctx context.Context) PfsGroupPtrOutput {
	return o
}

func (o PfsGroupPtrOutput) Elem() PfsGroupOutput {
	return o.ApplyT(func(v *PfsGroup) PfsGroup {
		if v != nil {
			return *v
		}
		var ret PfsGroup
		return ret
	}).(PfsGroupOutput)
}

func (o PfsGroupPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PfsGroupPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PfsGroup) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PfsGroupInput interface {
	pulumi.Input

	ToPfsGroupOutput() PfsGroupOutput
	ToPfsGroupOutputWithContext(context.Context) PfsGroupOutput
}

var pfsGroupPtrType = reflect.TypeOf((**PfsGroup)(nil)).Elem()

type PfsGroupPtrInput interface {
	pulumi.Input

	ToPfsGroupPtrOutput() PfsGroupPtrOutput
	ToPfsGroupPtrOutputWithContext(context.Context) PfsGroupPtrOutput
}

type pfsGroupPtr string

func PfsGroupPtr(v string) PfsGroupPtrInput {
	return (*pfsGroupPtr)(&v)
}

func (*pfsGroupPtr) ElementType() reflect.Type {
	return pfsGroupPtrType
}

func (in *pfsGroupPtr) ToPfsGroupPtrOutput() PfsGroupPtrOutput {
	return pulumi.ToOutput(in).(PfsGroupPtrOutput)
}

func (in *pfsGroupPtr) ToPfsGroupPtrOutputWithContext(ctx context.Context) PfsGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PfsGroupPtrOutput)
}

type ProbeProtocol string

const (
	ProbeProtocolHttp = ProbeProtocol("Http")
	ProbeProtocolTcp  = ProbeProtocol("Tcp")
)

func (ProbeProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeProtocol)(nil)).Elem()
}

func (e ProbeProtocol) ToProbeProtocolOutput() ProbeProtocolOutput {
	return pulumi.ToOutput(e).(ProbeProtocolOutput)
}

func (e ProbeProtocol) ToProbeProtocolOutputWithContext(ctx context.Context) ProbeProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProbeProtocolOutput)
}

func (e ProbeProtocol) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return e.ToProbeProtocolPtrOutputWithContext(context.Background())
}

func (e ProbeProtocol) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return ProbeProtocol(e).ToProbeProtocolOutputWithContext(ctx).ToProbeProtocolPtrOutputWithContext(ctx)
}

func (e ProbeProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProbeProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProbeProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProbeProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProbeProtocolOutput struct{ *pulumi.OutputState }

func (ProbeProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeProtocol)(nil)).Elem()
}

func (o ProbeProtocolOutput) ToProbeProtocolOutput() ProbeProtocolOutput {
	return o
}

func (o ProbeProtocolOutput) ToProbeProtocolOutputWithContext(ctx context.Context) ProbeProtocolOutput {
	return o
}

func (o ProbeProtocolOutput) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return o.ToProbeProtocolPtrOutputWithContext(context.Background())
}

func (o ProbeProtocolOutput) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProbeProtocol) *ProbeProtocol {
		return &v
	}).(ProbeProtocolPtrOutput)
}

func (o ProbeProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProbeProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProbeProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProbeProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProbeProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProbeProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProbeProtocolPtrOutput struct{ *pulumi.OutputState }

func (ProbeProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProbeProtocol)(nil)).Elem()
}

func (o ProbeProtocolPtrOutput) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return o
}

func (o ProbeProtocolPtrOutput) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return o
}

func (o ProbeProtocolPtrOutput) Elem() ProbeProtocolOutput {
	return o.ApplyT(func(v *ProbeProtocol) ProbeProtocol {
		if v != nil {
			return *v
		}
		var ret ProbeProtocol
		return ret
	}).(ProbeProtocolOutput)
}

func (o ProbeProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProbeProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProbeProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProbeProtocolInput interface {
	pulumi.Input

	ToProbeProtocolOutput() ProbeProtocolOutput
	ToProbeProtocolOutputWithContext(context.Context) ProbeProtocolOutput
}

var probeProtocolPtrType = reflect.TypeOf((**ProbeProtocol)(nil)).Elem()

type ProbeProtocolPtrInput interface {
	pulumi.Input

	ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput
	ToProbeProtocolPtrOutputWithContext(context.Context) ProbeProtocolPtrOutput
}

type probeProtocolPtr string

func ProbeProtocolPtr(v string) ProbeProtocolPtrInput {
	return (*probeProtocolPtr)(&v)
}

func (*probeProtocolPtr) ElementType() reflect.Type {
	return probeProtocolPtrType
}

func (in *probeProtocolPtr) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return pulumi.ToOutput(in).(ProbeProtocolPtrOutput)
}

func (in *probeProtocolPtr) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProbeProtocolPtrOutput)
}

type PublicIPAddressSkuName string

const (
	PublicIPAddressSkuNameBasic    = PublicIPAddressSkuName("Basic")
	PublicIPAddressSkuNameStandard = PublicIPAddressSkuName("Standard")
)

func (PublicIPAddressSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSkuName)(nil)).Elem()
}

func (e PublicIPAddressSkuName) ToPublicIPAddressSkuNameOutput() PublicIPAddressSkuNameOutput {
	return pulumi.ToOutput(e).(PublicIPAddressSkuNameOutput)
}

func (e PublicIPAddressSkuName) ToPublicIPAddressSkuNameOutputWithContext(ctx context.Context) PublicIPAddressSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicIPAddressSkuNameOutput)
}

func (e PublicIPAddressSkuName) ToPublicIPAddressSkuNamePtrOutput() PublicIPAddressSkuNamePtrOutput {
	return e.ToPublicIPAddressSkuNamePtrOutputWithContext(context.Background())
}

func (e PublicIPAddressSkuName) ToPublicIPAddressSkuNamePtrOutputWithContext(ctx context.Context) PublicIPAddressSkuNamePtrOutput {
	return PublicIPAddressSkuName(e).ToPublicIPAddressSkuNameOutputWithContext(ctx).ToPublicIPAddressSkuNamePtrOutputWithContext(ctx)
}

func (e PublicIPAddressSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPAddressSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPAddressSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicIPAddressSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicIPAddressSkuNameOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSkuName)(nil)).Elem()
}

func (o PublicIPAddressSkuNameOutput) ToPublicIPAddressSkuNameOutput() PublicIPAddressSkuNameOutput {
	return o
}

func (o PublicIPAddressSkuNameOutput) ToPublicIPAddressSkuNameOutputWithContext(ctx context.Context) PublicIPAddressSkuNameOutput {
	return o
}

func (o PublicIPAddressSkuNameOutput) ToPublicIPAddressSkuNamePtrOutput() PublicIPAddressSkuNamePtrOutput {
	return o.ToPublicIPAddressSkuNamePtrOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuNameOutput) ToPublicIPAddressSkuNamePtrOutputWithContext(ctx context.Context) PublicIPAddressSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressSkuName) *PublicIPAddressSkuName {
		return &v
	}).(PublicIPAddressSkuNamePtrOutput)
}

func (o PublicIPAddressSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPAddressSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicIPAddressSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPAddressSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicIPAddressSkuNamePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressSkuName)(nil)).Elem()
}

func (o PublicIPAddressSkuNamePtrOutput) ToPublicIPAddressSkuNamePtrOutput() PublicIPAddressSkuNamePtrOutput {
	return o
}

func (o PublicIPAddressSkuNamePtrOutput) ToPublicIPAddressSkuNamePtrOutputWithContext(ctx context.Context) PublicIPAddressSkuNamePtrOutput {
	return o
}

func (o PublicIPAddressSkuNamePtrOutput) Elem() PublicIPAddressSkuNameOutput {
	return o.ApplyT(func(v *PublicIPAddressSkuName) PublicIPAddressSkuName {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressSkuName
		return ret
	}).(PublicIPAddressSkuNameOutput)
}

func (o PublicIPAddressSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicIPAddressSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicIPAddressSkuNameInput interface {
	pulumi.Input

	ToPublicIPAddressSkuNameOutput() PublicIPAddressSkuNameOutput
	ToPublicIPAddressSkuNameOutputWithContext(context.Context) PublicIPAddressSkuNameOutput
}

var publicIPAddressSkuNamePtrType = reflect.TypeOf((**PublicIPAddressSkuName)(nil)).Elem()

type PublicIPAddressSkuNamePtrInput interface {
	pulumi.Input

	ToPublicIPAddressSkuNamePtrOutput() PublicIPAddressSkuNamePtrOutput
	ToPublicIPAddressSkuNamePtrOutputWithContext(context.Context) PublicIPAddressSkuNamePtrOutput
}

type publicIPAddressSkuNamePtr string

func PublicIPAddressSkuNamePtr(v string) PublicIPAddressSkuNamePtrInput {
	return (*publicIPAddressSkuNamePtr)(&v)
}

func (*publicIPAddressSkuNamePtr) ElementType() reflect.Type {
	return publicIPAddressSkuNamePtrType
}

func (in *publicIPAddressSkuNamePtr) ToPublicIPAddressSkuNamePtrOutput() PublicIPAddressSkuNamePtrOutput {
	return pulumi.ToOutput(in).(PublicIPAddressSkuNamePtrOutput)
}

func (in *publicIPAddressSkuNamePtr) ToPublicIPAddressSkuNamePtrOutputWithContext(ctx context.Context) PublicIPAddressSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicIPAddressSkuNamePtrOutput)
}

type RouteFilterRuleTypeEnum string

const (
	RouteFilterRuleTypeEnumCommunity = RouteFilterRuleTypeEnum("Community")
)

func (RouteFilterRuleTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterRuleTypeEnum)(nil)).Elem()
}

func (e RouteFilterRuleTypeEnum) ToRouteFilterRuleTypeEnumOutput() RouteFilterRuleTypeEnumOutput {
	return pulumi.ToOutput(e).(RouteFilterRuleTypeEnumOutput)
}

func (e RouteFilterRuleTypeEnum) ToRouteFilterRuleTypeEnumOutputWithContext(ctx context.Context) RouteFilterRuleTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RouteFilterRuleTypeEnumOutput)
}

func (e RouteFilterRuleTypeEnum) ToRouteFilterRuleTypeEnumPtrOutput() RouteFilterRuleTypeEnumPtrOutput {
	return e.ToRouteFilterRuleTypeEnumPtrOutputWithContext(context.Background())
}

func (e RouteFilterRuleTypeEnum) ToRouteFilterRuleTypeEnumPtrOutputWithContext(ctx context.Context) RouteFilterRuleTypeEnumPtrOutput {
	return RouteFilterRuleTypeEnum(e).ToRouteFilterRuleTypeEnumOutputWithContext(ctx).ToRouteFilterRuleTypeEnumPtrOutputWithContext(ctx)
}

func (e RouteFilterRuleTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RouteFilterRuleTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RouteFilterRuleTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RouteFilterRuleTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RouteFilterRuleTypeEnumOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterRuleTypeEnum)(nil)).Elem()
}

func (o RouteFilterRuleTypeEnumOutput) ToRouteFilterRuleTypeEnumOutput() RouteFilterRuleTypeEnumOutput {
	return o
}

func (o RouteFilterRuleTypeEnumOutput) ToRouteFilterRuleTypeEnumOutputWithContext(ctx context.Context) RouteFilterRuleTypeEnumOutput {
	return o
}

func (o RouteFilterRuleTypeEnumOutput) ToRouteFilterRuleTypeEnumPtrOutput() RouteFilterRuleTypeEnumPtrOutput {
	return o.ToRouteFilterRuleTypeEnumPtrOutputWithContext(context.Background())
}

func (o RouteFilterRuleTypeEnumOutput) ToRouteFilterRuleTypeEnumPtrOutputWithContext(ctx context.Context) RouteFilterRuleTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RouteFilterRuleTypeEnum) *RouteFilterRuleTypeEnum {
		return &v
	}).(RouteFilterRuleTypeEnumPtrOutput)
}

func (o RouteFilterRuleTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RouteFilterRuleTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RouteFilterRuleTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RouteFilterRuleTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RouteFilterRuleTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RouteFilterRuleTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RouteFilterRuleTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleTypeEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterRuleTypeEnum)(nil)).Elem()
}

func (o RouteFilterRuleTypeEnumPtrOutput) ToRouteFilterRuleTypeEnumPtrOutput() RouteFilterRuleTypeEnumPtrOutput {
	return o
}

func (o RouteFilterRuleTypeEnumPtrOutput) ToRouteFilterRuleTypeEnumPtrOutputWithContext(ctx context.Context) RouteFilterRuleTypeEnumPtrOutput {
	return o
}

func (o RouteFilterRuleTypeEnumPtrOutput) Elem() RouteFilterRuleTypeEnumOutput {
	return o.ApplyT(func(v *RouteFilterRuleTypeEnum) RouteFilterRuleTypeEnum {
		if v != nil {
			return *v
		}
		var ret RouteFilterRuleTypeEnum
		return ret
	}).(RouteFilterRuleTypeEnumOutput)
}

func (o RouteFilterRuleTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RouteFilterRuleTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RouteFilterRuleTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RouteFilterRuleTypeEnumInput interface {
	pulumi.Input

	ToRouteFilterRuleTypeEnumOutput() RouteFilterRuleTypeEnumOutput
	ToRouteFilterRuleTypeEnumOutputWithContext(context.Context) RouteFilterRuleTypeEnumOutput
}

var routeFilterRuleTypeEnumPtrType = reflect.TypeOf((**RouteFilterRuleTypeEnum)(nil)).Elem()

type RouteFilterRuleTypeEnumPtrInput interface {
	pulumi.Input

	ToRouteFilterRuleTypeEnumPtrOutput() RouteFilterRuleTypeEnumPtrOutput
	ToRouteFilterRuleTypeEnumPtrOutputWithContext(context.Context) RouteFilterRuleTypeEnumPtrOutput
}

type routeFilterRuleTypeEnumPtr string

func RouteFilterRuleTypeEnumPtr(v string) RouteFilterRuleTypeEnumPtrInput {
	return (*routeFilterRuleTypeEnumPtr)(&v)
}

func (*routeFilterRuleTypeEnumPtr) ElementType() reflect.Type {
	return routeFilterRuleTypeEnumPtrType
}

func (in *routeFilterRuleTypeEnumPtr) ToRouteFilterRuleTypeEnumPtrOutput() RouteFilterRuleTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(RouteFilterRuleTypeEnumPtrOutput)
}

func (in *routeFilterRuleTypeEnumPtr) ToRouteFilterRuleTypeEnumPtrOutputWithContext(ctx context.Context) RouteFilterRuleTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RouteFilterRuleTypeEnumPtrOutput)
}

type RouteNextHopType string

const (
	RouteNextHopTypeVirtualNetworkGateway = RouteNextHopType("VirtualNetworkGateway")
	RouteNextHopTypeVnetLocal             = RouteNextHopType("VnetLocal")
	RouteNextHopTypeInternet              = RouteNextHopType("Internet")
	RouteNextHopTypeVirtualAppliance      = RouteNextHopType("VirtualAppliance")
	RouteNextHopTypeNone                  = RouteNextHopType("None")
)

func (RouteNextHopType) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteNextHopType)(nil)).Elem()
}

func (e RouteNextHopType) ToRouteNextHopTypeOutput() RouteNextHopTypeOutput {
	return pulumi.ToOutput(e).(RouteNextHopTypeOutput)
}

func (e RouteNextHopType) ToRouteNextHopTypeOutputWithContext(ctx context.Context) RouteNextHopTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RouteNextHopTypeOutput)
}

func (e RouteNextHopType) ToRouteNextHopTypePtrOutput() RouteNextHopTypePtrOutput {
	return e.ToRouteNextHopTypePtrOutputWithContext(context.Background())
}

func (e RouteNextHopType) ToRouteNextHopTypePtrOutputWithContext(ctx context.Context) RouteNextHopTypePtrOutput {
	return RouteNextHopType(e).ToRouteNextHopTypeOutputWithContext(ctx).ToRouteNextHopTypePtrOutputWithContext(ctx)
}

func (e RouteNextHopType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RouteNextHopType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RouteNextHopType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RouteNextHopType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RouteNextHopTypeOutput struct{ *pulumi.OutputState }

func (RouteNextHopTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteNextHopType)(nil)).Elem()
}

func (o RouteNextHopTypeOutput) ToRouteNextHopTypeOutput() RouteNextHopTypeOutput {
	return o
}

func (o RouteNextHopTypeOutput) ToRouteNextHopTypeOutputWithContext(ctx context.Context) RouteNextHopTypeOutput {
	return o
}

func (o RouteNextHopTypeOutput) ToRouteNextHopTypePtrOutput() RouteNextHopTypePtrOutput {
	return o.ToRouteNextHopTypePtrOutputWithContext(context.Background())
}

func (o RouteNextHopTypeOutput) ToRouteNextHopTypePtrOutputWithContext(ctx context.Context) RouteNextHopTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RouteNextHopType) *RouteNextHopType {
		return &v
	}).(RouteNextHopTypePtrOutput)
}

func (o RouteNextHopTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RouteNextHopTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RouteNextHopType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RouteNextHopTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RouteNextHopTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RouteNextHopType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RouteNextHopTypePtrOutput struct{ *pulumi.OutputState }

func (RouteNextHopTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteNextHopType)(nil)).Elem()
}

func (o RouteNextHopTypePtrOutput) ToRouteNextHopTypePtrOutput() RouteNextHopTypePtrOutput {
	return o
}

func (o RouteNextHopTypePtrOutput) ToRouteNextHopTypePtrOutputWithContext(ctx context.Context) RouteNextHopTypePtrOutput {
	return o
}

func (o RouteNextHopTypePtrOutput) Elem() RouteNextHopTypeOutput {
	return o.ApplyT(func(v *RouteNextHopType) RouteNextHopType {
		if v != nil {
			return *v
		}
		var ret RouteNextHopType
		return ret
	}).(RouteNextHopTypeOutput)
}

func (o RouteNextHopTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RouteNextHopTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RouteNextHopType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RouteNextHopTypeInput interface {
	pulumi.Input

	ToRouteNextHopTypeOutput() RouteNextHopTypeOutput
	ToRouteNextHopTypeOutputWithContext(context.Context) RouteNextHopTypeOutput
}

var routeNextHopTypePtrType = reflect.TypeOf((**RouteNextHopType)(nil)).Elem()

type RouteNextHopTypePtrInput interface {
	pulumi.Input

	ToRouteNextHopTypePtrOutput() RouteNextHopTypePtrOutput
	ToRouteNextHopTypePtrOutputWithContext(context.Context) RouteNextHopTypePtrOutput
}

type routeNextHopTypePtr string

func RouteNextHopTypePtr(v string) RouteNextHopTypePtrInput {
	return (*routeNextHopTypePtr)(&v)
}

func (*routeNextHopTypePtr) ElementType() reflect.Type {
	return routeNextHopTypePtrType
}

func (in *routeNextHopTypePtr) ToRouteNextHopTypePtrOutput() RouteNextHopTypePtrOutput {
	return pulumi.ToOutput(in).(RouteNextHopTypePtrOutput)
}

func (in *routeNextHopTypePtr) ToRouteNextHopTypePtrOutputWithContext(ctx context.Context) RouteNextHopTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RouteNextHopTypePtrOutput)
}

type SecurityRuleAccess string

const (
	SecurityRuleAccessAllow = SecurityRuleAccess("Allow")
	SecurityRuleAccessDeny  = SecurityRuleAccess("Deny")
)

func (SecurityRuleAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleAccess)(nil)).Elem()
}

func (e SecurityRuleAccess) ToSecurityRuleAccessOutput() SecurityRuleAccessOutput {
	return pulumi.ToOutput(e).(SecurityRuleAccessOutput)
}

func (e SecurityRuleAccess) ToSecurityRuleAccessOutputWithContext(ctx context.Context) SecurityRuleAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityRuleAccessOutput)
}

func (e SecurityRuleAccess) ToSecurityRuleAccessPtrOutput() SecurityRuleAccessPtrOutput {
	return e.ToSecurityRuleAccessPtrOutputWithContext(context.Background())
}

func (e SecurityRuleAccess) ToSecurityRuleAccessPtrOutputWithContext(ctx context.Context) SecurityRuleAccessPtrOutput {
	return SecurityRuleAccess(e).ToSecurityRuleAccessOutputWithContext(ctx).ToSecurityRuleAccessPtrOutputWithContext(ctx)
}

func (e SecurityRuleAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityRuleAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityRuleAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityRuleAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityRuleAccessOutput struct{ *pulumi.OutputState }

func (SecurityRuleAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleAccess)(nil)).Elem()
}

func (o SecurityRuleAccessOutput) ToSecurityRuleAccessOutput() SecurityRuleAccessOutput {
	return o
}

func (o SecurityRuleAccessOutput) ToSecurityRuleAccessOutputWithContext(ctx context.Context) SecurityRuleAccessOutput {
	return o
}

func (o SecurityRuleAccessOutput) ToSecurityRuleAccessPtrOutput() SecurityRuleAccessPtrOutput {
	return o.ToSecurityRuleAccessPtrOutputWithContext(context.Background())
}

func (o SecurityRuleAccessOutput) ToSecurityRuleAccessPtrOutputWithContext(ctx context.Context) SecurityRuleAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityRuleAccess) *SecurityRuleAccess {
		return &v
	}).(SecurityRuleAccessPtrOutput)
}

func (o SecurityRuleAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityRuleAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityRuleAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityRuleAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityRuleAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityRuleAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityRuleAccessPtrOutput struct{ *pulumi.OutputState }

func (SecurityRuleAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityRuleAccess)(nil)).Elem()
}

func (o SecurityRuleAccessPtrOutput) ToSecurityRuleAccessPtrOutput() SecurityRuleAccessPtrOutput {
	return o
}

func (o SecurityRuleAccessPtrOutput) ToSecurityRuleAccessPtrOutputWithContext(ctx context.Context) SecurityRuleAccessPtrOutput {
	return o
}

func (o SecurityRuleAccessPtrOutput) Elem() SecurityRuleAccessOutput {
	return o.ApplyT(func(v *SecurityRuleAccess) SecurityRuleAccess {
		if v != nil {
			return *v
		}
		var ret SecurityRuleAccess
		return ret
	}).(SecurityRuleAccessOutput)
}

func (o SecurityRuleAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityRuleAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityRuleAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityRuleAccessInput interface {
	pulumi.Input

	ToSecurityRuleAccessOutput() SecurityRuleAccessOutput
	ToSecurityRuleAccessOutputWithContext(context.Context) SecurityRuleAccessOutput
}

var securityRuleAccessPtrType = reflect.TypeOf((**SecurityRuleAccess)(nil)).Elem()

type SecurityRuleAccessPtrInput interface {
	pulumi.Input

	ToSecurityRuleAccessPtrOutput() SecurityRuleAccessPtrOutput
	ToSecurityRuleAccessPtrOutputWithContext(context.Context) SecurityRuleAccessPtrOutput
}

type securityRuleAccessPtr string

func SecurityRuleAccessPtr(v string) SecurityRuleAccessPtrInput {
	return (*securityRuleAccessPtr)(&v)
}

func (*securityRuleAccessPtr) ElementType() reflect.Type {
	return securityRuleAccessPtrType
}

func (in *securityRuleAccessPtr) ToSecurityRuleAccessPtrOutput() SecurityRuleAccessPtrOutput {
	return pulumi.ToOutput(in).(SecurityRuleAccessPtrOutput)
}

func (in *securityRuleAccessPtr) ToSecurityRuleAccessPtrOutputWithContext(ctx context.Context) SecurityRuleAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityRuleAccessPtrOutput)
}

type SecurityRuleDirection string

const (
	SecurityRuleDirectionInbound  = SecurityRuleDirection("Inbound")
	SecurityRuleDirectionOutbound = SecurityRuleDirection("Outbound")
)

func (SecurityRuleDirection) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleDirection)(nil)).Elem()
}

func (e SecurityRuleDirection) ToSecurityRuleDirectionOutput() SecurityRuleDirectionOutput {
	return pulumi.ToOutput(e).(SecurityRuleDirectionOutput)
}

func (e SecurityRuleDirection) ToSecurityRuleDirectionOutputWithContext(ctx context.Context) SecurityRuleDirectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityRuleDirectionOutput)
}

func (e SecurityRuleDirection) ToSecurityRuleDirectionPtrOutput() SecurityRuleDirectionPtrOutput {
	return e.ToSecurityRuleDirectionPtrOutputWithContext(context.Background())
}

func (e SecurityRuleDirection) ToSecurityRuleDirectionPtrOutputWithContext(ctx context.Context) SecurityRuleDirectionPtrOutput {
	return SecurityRuleDirection(e).ToSecurityRuleDirectionOutputWithContext(ctx).ToSecurityRuleDirectionPtrOutputWithContext(ctx)
}

func (e SecurityRuleDirection) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityRuleDirection) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityRuleDirection) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityRuleDirection) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityRuleDirectionOutput struct{ *pulumi.OutputState }

func (SecurityRuleDirectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleDirection)(nil)).Elem()
}

func (o SecurityRuleDirectionOutput) ToSecurityRuleDirectionOutput() SecurityRuleDirectionOutput {
	return o
}

func (o SecurityRuleDirectionOutput) ToSecurityRuleDirectionOutputWithContext(ctx context.Context) SecurityRuleDirectionOutput {
	return o
}

func (o SecurityRuleDirectionOutput) ToSecurityRuleDirectionPtrOutput() SecurityRuleDirectionPtrOutput {
	return o.ToSecurityRuleDirectionPtrOutputWithContext(context.Background())
}

func (o SecurityRuleDirectionOutput) ToSecurityRuleDirectionPtrOutputWithContext(ctx context.Context) SecurityRuleDirectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityRuleDirection) *SecurityRuleDirection {
		return &v
	}).(SecurityRuleDirectionPtrOutput)
}

func (o SecurityRuleDirectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityRuleDirectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityRuleDirection) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityRuleDirectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityRuleDirectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityRuleDirection) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityRuleDirectionPtrOutput struct{ *pulumi.OutputState }

func (SecurityRuleDirectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityRuleDirection)(nil)).Elem()
}

func (o SecurityRuleDirectionPtrOutput) ToSecurityRuleDirectionPtrOutput() SecurityRuleDirectionPtrOutput {
	return o
}

func (o SecurityRuleDirectionPtrOutput) ToSecurityRuleDirectionPtrOutputWithContext(ctx context.Context) SecurityRuleDirectionPtrOutput {
	return o
}

func (o SecurityRuleDirectionPtrOutput) Elem() SecurityRuleDirectionOutput {
	return o.ApplyT(func(v *SecurityRuleDirection) SecurityRuleDirection {
		if v != nil {
			return *v
		}
		var ret SecurityRuleDirection
		return ret
	}).(SecurityRuleDirectionOutput)
}

func (o SecurityRuleDirectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityRuleDirectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityRuleDirection) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityRuleDirectionInput interface {
	pulumi.Input

	ToSecurityRuleDirectionOutput() SecurityRuleDirectionOutput
	ToSecurityRuleDirectionOutputWithContext(context.Context) SecurityRuleDirectionOutput
}

var securityRuleDirectionPtrType = reflect.TypeOf((**SecurityRuleDirection)(nil)).Elem()

type SecurityRuleDirectionPtrInput interface {
	pulumi.Input

	ToSecurityRuleDirectionPtrOutput() SecurityRuleDirectionPtrOutput
	ToSecurityRuleDirectionPtrOutputWithContext(context.Context) SecurityRuleDirectionPtrOutput
}

type securityRuleDirectionPtr string

func SecurityRuleDirectionPtr(v string) SecurityRuleDirectionPtrInput {
	return (*securityRuleDirectionPtr)(&v)
}

func (*securityRuleDirectionPtr) ElementType() reflect.Type {
	return securityRuleDirectionPtrType
}

func (in *securityRuleDirectionPtr) ToSecurityRuleDirectionPtrOutput() SecurityRuleDirectionPtrOutput {
	return pulumi.ToOutput(in).(SecurityRuleDirectionPtrOutput)
}

func (in *securityRuleDirectionPtr) ToSecurityRuleDirectionPtrOutputWithContext(ctx context.Context) SecurityRuleDirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityRuleDirectionPtrOutput)
}

type SecurityRuleProtocol string

const (
	SecurityRuleProtocolTcp      = SecurityRuleProtocol("Tcp")
	SecurityRuleProtocolUdp      = SecurityRuleProtocol("Udp")
	SecurityRuleProtocolAsterisk = SecurityRuleProtocol("*")
)

func (SecurityRuleProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleProtocol)(nil)).Elem()
}

func (e SecurityRuleProtocol) ToSecurityRuleProtocolOutput() SecurityRuleProtocolOutput {
	return pulumi.ToOutput(e).(SecurityRuleProtocolOutput)
}

func (e SecurityRuleProtocol) ToSecurityRuleProtocolOutputWithContext(ctx context.Context) SecurityRuleProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityRuleProtocolOutput)
}

func (e SecurityRuleProtocol) ToSecurityRuleProtocolPtrOutput() SecurityRuleProtocolPtrOutput {
	return e.ToSecurityRuleProtocolPtrOutputWithContext(context.Background())
}

func (e SecurityRuleProtocol) ToSecurityRuleProtocolPtrOutputWithContext(ctx context.Context) SecurityRuleProtocolPtrOutput {
	return SecurityRuleProtocol(e).ToSecurityRuleProtocolOutputWithContext(ctx).ToSecurityRuleProtocolPtrOutputWithContext(ctx)
}

func (e SecurityRuleProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityRuleProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityRuleProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityRuleProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityRuleProtocolOutput struct{ *pulumi.OutputState }

func (SecurityRuleProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleProtocol)(nil)).Elem()
}

func (o SecurityRuleProtocolOutput) ToSecurityRuleProtocolOutput() SecurityRuleProtocolOutput {
	return o
}

func (o SecurityRuleProtocolOutput) ToSecurityRuleProtocolOutputWithContext(ctx context.Context) SecurityRuleProtocolOutput {
	return o
}

func (o SecurityRuleProtocolOutput) ToSecurityRuleProtocolPtrOutput() SecurityRuleProtocolPtrOutput {
	return o.ToSecurityRuleProtocolPtrOutputWithContext(context.Background())
}

func (o SecurityRuleProtocolOutput) ToSecurityRuleProtocolPtrOutputWithContext(ctx context.Context) SecurityRuleProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityRuleProtocol) *SecurityRuleProtocol {
		return &v
	}).(SecurityRuleProtocolPtrOutput)
}

func (o SecurityRuleProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityRuleProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityRuleProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityRuleProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityRuleProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityRuleProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityRuleProtocolPtrOutput struct{ *pulumi.OutputState }

func (SecurityRuleProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityRuleProtocol)(nil)).Elem()
}

func (o SecurityRuleProtocolPtrOutput) ToSecurityRuleProtocolPtrOutput() SecurityRuleProtocolPtrOutput {
	return o
}

func (o SecurityRuleProtocolPtrOutput) ToSecurityRuleProtocolPtrOutputWithContext(ctx context.Context) SecurityRuleProtocolPtrOutput {
	return o
}

func (o SecurityRuleProtocolPtrOutput) Elem() SecurityRuleProtocolOutput {
	return o.ApplyT(func(v *SecurityRuleProtocol) SecurityRuleProtocol {
		if v != nil {
			return *v
		}
		var ret SecurityRuleProtocol
		return ret
	}).(SecurityRuleProtocolOutput)
}

func (o SecurityRuleProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityRuleProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityRuleProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityRuleProtocolInput interface {
	pulumi.Input

	ToSecurityRuleProtocolOutput() SecurityRuleProtocolOutput
	ToSecurityRuleProtocolOutputWithContext(context.Context) SecurityRuleProtocolOutput
}

var securityRuleProtocolPtrType = reflect.TypeOf((**SecurityRuleProtocol)(nil)).Elem()

type SecurityRuleProtocolPtrInput interface {
	pulumi.Input

	ToSecurityRuleProtocolPtrOutput() SecurityRuleProtocolPtrOutput
	ToSecurityRuleProtocolPtrOutputWithContext(context.Context) SecurityRuleProtocolPtrOutput
}

type securityRuleProtocolPtr string

func SecurityRuleProtocolPtr(v string) SecurityRuleProtocolPtrInput {
	return (*securityRuleProtocolPtr)(&v)
}

func (*securityRuleProtocolPtr) ElementType() reflect.Type {
	return securityRuleProtocolPtrType
}

func (in *securityRuleProtocolPtr) ToSecurityRuleProtocolPtrOutput() SecurityRuleProtocolPtrOutput {
	return pulumi.ToOutput(in).(SecurityRuleProtocolPtrOutput)
}

func (in *securityRuleProtocolPtr) ToSecurityRuleProtocolPtrOutputWithContext(ctx context.Context) SecurityRuleProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityRuleProtocolPtrOutput)
}

type ServiceProviderProvisioningState string

const (
	ServiceProviderProvisioningStateNotProvisioned = ServiceProviderProvisioningState("NotProvisioned")
	ServiceProviderProvisioningStateProvisioning   = ServiceProviderProvisioningState("Provisioning")
	ServiceProviderProvisioningStateProvisioned    = ServiceProviderProvisioningState("Provisioned")
	ServiceProviderProvisioningStateDeprovisioning = ServiceProviderProvisioningState("Deprovisioning")
)

func (ServiceProviderProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderProvisioningState)(nil)).Elem()
}

func (e ServiceProviderProvisioningState) ToServiceProviderProvisioningStateOutput() ServiceProviderProvisioningStateOutput {
	return pulumi.ToOutput(e).(ServiceProviderProvisioningStateOutput)
}

func (e ServiceProviderProvisioningState) ToServiceProviderProvisioningStateOutputWithContext(ctx context.Context) ServiceProviderProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceProviderProvisioningStateOutput)
}

func (e ServiceProviderProvisioningState) ToServiceProviderProvisioningStatePtrOutput() ServiceProviderProvisioningStatePtrOutput {
	return e.ToServiceProviderProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ServiceProviderProvisioningState) ToServiceProviderProvisioningStatePtrOutputWithContext(ctx context.Context) ServiceProviderProvisioningStatePtrOutput {
	return ServiceProviderProvisioningState(e).ToServiceProviderProvisioningStateOutputWithContext(ctx).ToServiceProviderProvisioningStatePtrOutputWithContext(ctx)
}

func (e ServiceProviderProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceProviderProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceProviderProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceProviderProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceProviderProvisioningStateOutput struct{ *pulumi.OutputState }

func (ServiceProviderProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderProvisioningState)(nil)).Elem()
}

func (o ServiceProviderProvisioningStateOutput) ToServiceProviderProvisioningStateOutput() ServiceProviderProvisioningStateOutput {
	return o
}

func (o ServiceProviderProvisioningStateOutput) ToServiceProviderProvisioningStateOutputWithContext(ctx context.Context) ServiceProviderProvisioningStateOutput {
	return o
}

func (o ServiceProviderProvisioningStateOutput) ToServiceProviderProvisioningStatePtrOutput() ServiceProviderProvisioningStatePtrOutput {
	return o.ToServiceProviderProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ServiceProviderProvisioningStateOutput) ToServiceProviderProvisioningStatePtrOutputWithContext(ctx context.Context) ServiceProviderProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceProviderProvisioningState) *ServiceProviderProvisioningState {
		return &v
	}).(ServiceProviderProvisioningStatePtrOutput)
}

func (o ServiceProviderProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceProviderProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceProviderProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceProviderProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceProviderProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceProviderProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceProviderProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ServiceProviderProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceProviderProvisioningState)(nil)).Elem()
}

func (o ServiceProviderProvisioningStatePtrOutput) ToServiceProviderProvisioningStatePtrOutput() ServiceProviderProvisioningStatePtrOutput {
	return o
}

func (o ServiceProviderProvisioningStatePtrOutput) ToServiceProviderProvisioningStatePtrOutputWithContext(ctx context.Context) ServiceProviderProvisioningStatePtrOutput {
	return o
}

func (o ServiceProviderProvisioningStatePtrOutput) Elem() ServiceProviderProvisioningStateOutput {
	return o.ApplyT(func(v *ServiceProviderProvisioningState) ServiceProviderProvisioningState {
		if v != nil {
			return *v
		}
		var ret ServiceProviderProvisioningState
		return ret
	}).(ServiceProviderProvisioningStateOutput)
}

func (o ServiceProviderProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceProviderProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceProviderProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceProviderProvisioningStateInput interface {
	pulumi.Input

	ToServiceProviderProvisioningStateOutput() ServiceProviderProvisioningStateOutput
	ToServiceProviderProvisioningStateOutputWithContext(context.Context) ServiceProviderProvisioningStateOutput
}

var serviceProviderProvisioningStatePtrType = reflect.TypeOf((**ServiceProviderProvisioningState)(nil)).Elem()

type ServiceProviderProvisioningStatePtrInput interface {
	pulumi.Input

	ToServiceProviderProvisioningStatePtrOutput() ServiceProviderProvisioningStatePtrOutput
	ToServiceProviderProvisioningStatePtrOutputWithContext(context.Context) ServiceProviderProvisioningStatePtrOutput
}

type serviceProviderProvisioningStatePtr string

func ServiceProviderProvisioningStatePtr(v string) ServiceProviderProvisioningStatePtrInput {
	return (*serviceProviderProvisioningStatePtr)(&v)
}

func (*serviceProviderProvisioningStatePtr) ElementType() reflect.Type {
	return serviceProviderProvisioningStatePtrType
}

func (in *serviceProviderProvisioningStatePtr) ToServiceProviderProvisioningStatePtrOutput() ServiceProviderProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ServiceProviderProvisioningStatePtrOutput)
}

func (in *serviceProviderProvisioningStatePtr) ToServiceProviderProvisioningStatePtrOutputWithContext(ctx context.Context) ServiceProviderProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceProviderProvisioningStatePtrOutput)
}

type TransportProtocol string

const (
	TransportProtocolUdp = TransportProtocol("Udp")
	TransportProtocolTcp = TransportProtocol("Tcp")
)

func (TransportProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportProtocol)(nil)).Elem()
}

func (e TransportProtocol) ToTransportProtocolOutput() TransportProtocolOutput {
	return pulumi.ToOutput(e).(TransportProtocolOutput)
}

func (e TransportProtocol) ToTransportProtocolOutputWithContext(ctx context.Context) TransportProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TransportProtocolOutput)
}

func (e TransportProtocol) ToTransportProtocolPtrOutput() TransportProtocolPtrOutput {
	return e.ToTransportProtocolPtrOutputWithContext(context.Background())
}

func (e TransportProtocol) ToTransportProtocolPtrOutputWithContext(ctx context.Context) TransportProtocolPtrOutput {
	return TransportProtocol(e).ToTransportProtocolOutputWithContext(ctx).ToTransportProtocolPtrOutputWithContext(ctx)
}

func (e TransportProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransportProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TransportProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TransportProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TransportProtocolOutput struct{ *pulumi.OutputState }

func (TransportProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransportProtocol)(nil)).Elem()
}

func (o TransportProtocolOutput) ToTransportProtocolOutput() TransportProtocolOutput {
	return o
}

func (o TransportProtocolOutput) ToTransportProtocolOutputWithContext(ctx context.Context) TransportProtocolOutput {
	return o
}

func (o TransportProtocolOutput) ToTransportProtocolPtrOutput() TransportProtocolPtrOutput {
	return o.ToTransportProtocolPtrOutputWithContext(context.Background())
}

func (o TransportProtocolOutput) ToTransportProtocolPtrOutputWithContext(ctx context.Context) TransportProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransportProtocol) *TransportProtocol {
		return &v
	}).(TransportProtocolPtrOutput)
}

func (o TransportProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TransportProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TransportProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TransportProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TransportProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TransportProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TransportProtocolPtrOutput struct{ *pulumi.OutputState }

func (TransportProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransportProtocol)(nil)).Elem()
}

func (o TransportProtocolPtrOutput) ToTransportProtocolPtrOutput() TransportProtocolPtrOutput {
	return o
}

func (o TransportProtocolPtrOutput) ToTransportProtocolPtrOutputWithContext(ctx context.Context) TransportProtocolPtrOutput {
	return o
}

func (o TransportProtocolPtrOutput) Elem() TransportProtocolOutput {
	return o.ApplyT(func(v *TransportProtocol) TransportProtocol {
		if v != nil {
			return *v
		}
		var ret TransportProtocol
		return ret
	}).(TransportProtocolOutput)
}

func (o TransportProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TransportProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TransportProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TransportProtocolInput interface {
	pulumi.Input

	ToTransportProtocolOutput() TransportProtocolOutput
	ToTransportProtocolOutputWithContext(context.Context) TransportProtocolOutput
}

var transportProtocolPtrType = reflect.TypeOf((**TransportProtocol)(nil)).Elem()

type TransportProtocolPtrInput interface {
	pulumi.Input

	ToTransportProtocolPtrOutput() TransportProtocolPtrOutput
	ToTransportProtocolPtrOutputWithContext(context.Context) TransportProtocolPtrOutput
}

type transportProtocolPtr string

func TransportProtocolPtr(v string) TransportProtocolPtrInput {
	return (*transportProtocolPtr)(&v)
}

func (*transportProtocolPtr) ElementType() reflect.Type {
	return transportProtocolPtrType
}

func (in *transportProtocolPtr) ToTransportProtocolPtrOutput() TransportProtocolPtrOutput {
	return pulumi.ToOutput(in).(TransportProtocolPtrOutput)
}

func (in *transportProtocolPtr) ToTransportProtocolPtrOutputWithContext(ctx context.Context) TransportProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TransportProtocolPtrOutput)
}

type VirtualNetworkGatewayConnectionType string

const (
	VirtualNetworkGatewayConnectionTypeIPsec        = VirtualNetworkGatewayConnectionType("IPsec")
	VirtualNetworkGatewayConnectionTypeVnet2Vnet    = VirtualNetworkGatewayConnectionType("Vnet2Vnet")
	VirtualNetworkGatewayConnectionTypeExpressRoute = VirtualNetworkGatewayConnectionType("ExpressRoute")
	VirtualNetworkGatewayConnectionTypeVPNClient    = VirtualNetworkGatewayConnectionType("VPNClient")
)

func (VirtualNetworkGatewayConnectionType) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayConnectionType)(nil)).Elem()
}

func (e VirtualNetworkGatewayConnectionType) ToVirtualNetworkGatewayConnectionTypeOutput() VirtualNetworkGatewayConnectionTypeOutput {
	return pulumi.ToOutput(e).(VirtualNetworkGatewayConnectionTypeOutput)
}

func (e VirtualNetworkGatewayConnectionType) ToVirtualNetworkGatewayConnectionTypeOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkGatewayConnectionTypeOutput)
}

func (e VirtualNetworkGatewayConnectionType) ToVirtualNetworkGatewayConnectionTypePtrOutput() VirtualNetworkGatewayConnectionTypePtrOutput {
	return e.ToVirtualNetworkGatewayConnectionTypePtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewayConnectionType) ToVirtualNetworkGatewayConnectionTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionTypePtrOutput {
	return VirtualNetworkGatewayConnectionType(e).ToVirtualNetworkGatewayConnectionTypeOutputWithContext(ctx).ToVirtualNetworkGatewayConnectionTypePtrOutputWithContext(ctx)
}

func (e VirtualNetworkGatewayConnectionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewayConnectionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewayConnectionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewayConnectionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkGatewayConnectionTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayConnectionType)(nil)).Elem()
}

func (o VirtualNetworkGatewayConnectionTypeOutput) ToVirtualNetworkGatewayConnectionTypeOutput() VirtualNetworkGatewayConnectionTypeOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionTypeOutput) ToVirtualNetworkGatewayConnectionTypeOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionTypeOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionTypeOutput) ToVirtualNetworkGatewayConnectionTypePtrOutput() VirtualNetworkGatewayConnectionTypePtrOutput {
	return o.ToVirtualNetworkGatewayConnectionTypePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionTypeOutput) ToVirtualNetworkGatewayConnectionTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewayConnectionType) *VirtualNetworkGatewayConnectionType {
		return &v
	}).(VirtualNetworkGatewayConnectionTypePtrOutput)
}

func (o VirtualNetworkGatewayConnectionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewayConnectionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewayConnectionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayConnectionTypePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayConnectionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayConnectionType)(nil)).Elem()
}

func (o VirtualNetworkGatewayConnectionTypePtrOutput) ToVirtualNetworkGatewayConnectionTypePtrOutput() VirtualNetworkGatewayConnectionTypePtrOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionTypePtrOutput) ToVirtualNetworkGatewayConnectionTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionTypePtrOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionTypePtrOutput) Elem() VirtualNetworkGatewayConnectionTypeOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnectionType) VirtualNetworkGatewayConnectionType {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayConnectionType
		return ret
	}).(VirtualNetworkGatewayConnectionTypeOutput)
}

func (o VirtualNetworkGatewayConnectionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkGatewayConnectionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkGatewayConnectionTypeInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayConnectionTypeOutput() VirtualNetworkGatewayConnectionTypeOutput
	ToVirtualNetworkGatewayConnectionTypeOutputWithContext(context.Context) VirtualNetworkGatewayConnectionTypeOutput
}

var virtualNetworkGatewayConnectionTypePtrType = reflect.TypeOf((**VirtualNetworkGatewayConnectionType)(nil)).Elem()

type VirtualNetworkGatewayConnectionTypePtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayConnectionTypePtrOutput() VirtualNetworkGatewayConnectionTypePtrOutput
	ToVirtualNetworkGatewayConnectionTypePtrOutputWithContext(context.Context) VirtualNetworkGatewayConnectionTypePtrOutput
}

type virtualNetworkGatewayConnectionTypePtr string

func VirtualNetworkGatewayConnectionTypePtr(v string) VirtualNetworkGatewayConnectionTypePtrInput {
	return (*virtualNetworkGatewayConnectionTypePtr)(&v)
}

func (*virtualNetworkGatewayConnectionTypePtr) ElementType() reflect.Type {
	return virtualNetworkGatewayConnectionTypePtrType
}

func (in *virtualNetworkGatewayConnectionTypePtr) ToVirtualNetworkGatewayConnectionTypePtrOutput() VirtualNetworkGatewayConnectionTypePtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkGatewayConnectionTypePtrOutput)
}

func (in *virtualNetworkGatewayConnectionTypePtr) ToVirtualNetworkGatewayConnectionTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkGatewayConnectionTypePtrOutput)
}

type VirtualNetworkGatewaySkuName string

const (
	VirtualNetworkGatewaySkuNameBasic            = VirtualNetworkGatewaySkuName("Basic")
	VirtualNetworkGatewaySkuNameHighPerformance  = VirtualNetworkGatewaySkuName("HighPerformance")
	VirtualNetworkGatewaySkuNameStandard         = VirtualNetworkGatewaySkuName("Standard")
	VirtualNetworkGatewaySkuNameUltraPerformance = VirtualNetworkGatewaySkuName("UltraPerformance")
	VirtualNetworkGatewaySkuNameVpnGw1           = VirtualNetworkGatewaySkuName("VpnGw1")
	VirtualNetworkGatewaySkuNameVpnGw2           = VirtualNetworkGatewaySkuName("VpnGw2")
	VirtualNetworkGatewaySkuNameVpnGw3           = VirtualNetworkGatewaySkuName("VpnGw3")
)

func (VirtualNetworkGatewaySkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySkuName)(nil)).Elem()
}

func (e VirtualNetworkGatewaySkuName) ToVirtualNetworkGatewaySkuNameOutput() VirtualNetworkGatewaySkuNameOutput {
	return pulumi.ToOutput(e).(VirtualNetworkGatewaySkuNameOutput)
}

func (e VirtualNetworkGatewaySkuName) ToVirtualNetworkGatewaySkuNameOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkGatewaySkuNameOutput)
}

func (e VirtualNetworkGatewaySkuName) ToVirtualNetworkGatewaySkuNamePtrOutput() VirtualNetworkGatewaySkuNamePtrOutput {
	return e.ToVirtualNetworkGatewaySkuNamePtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewaySkuName) ToVirtualNetworkGatewaySkuNamePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuNamePtrOutput {
	return VirtualNetworkGatewaySkuName(e).ToVirtualNetworkGatewaySkuNameOutputWithContext(ctx).ToVirtualNetworkGatewaySkuNamePtrOutputWithContext(ctx)
}

func (e VirtualNetworkGatewaySkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewaySkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewaySkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewaySkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkGatewaySkuNameOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySkuName)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuNameOutput) ToVirtualNetworkGatewaySkuNameOutput() VirtualNetworkGatewaySkuNameOutput {
	return o
}

func (o VirtualNetworkGatewaySkuNameOutput) ToVirtualNetworkGatewaySkuNameOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuNameOutput {
	return o
}

func (o VirtualNetworkGatewaySkuNameOutput) ToVirtualNetworkGatewaySkuNamePtrOutput() VirtualNetworkGatewaySkuNamePtrOutput {
	return o.ToVirtualNetworkGatewaySkuNamePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuNameOutput) ToVirtualNetworkGatewaySkuNamePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewaySkuName) *VirtualNetworkGatewaySkuName {
		return &v
	}).(VirtualNetworkGatewaySkuNamePtrOutput)
}

func (o VirtualNetworkGatewaySkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewaySkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewaySkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewaySkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuNamePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySkuName)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuNamePtrOutput) ToVirtualNetworkGatewaySkuNamePtrOutput() VirtualNetworkGatewaySkuNamePtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuNamePtrOutput) ToVirtualNetworkGatewaySkuNamePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuNamePtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuNamePtrOutput) Elem() VirtualNetworkGatewaySkuNameOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuName) VirtualNetworkGatewaySkuName {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewaySkuName
		return ret
	}).(VirtualNetworkGatewaySkuNameOutput)
}

func (o VirtualNetworkGatewaySkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkGatewaySkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkGatewaySkuNameInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuNameOutput() VirtualNetworkGatewaySkuNameOutput
	ToVirtualNetworkGatewaySkuNameOutputWithContext(context.Context) VirtualNetworkGatewaySkuNameOutput
}

var virtualNetworkGatewaySkuNamePtrType = reflect.TypeOf((**VirtualNetworkGatewaySkuName)(nil)).Elem()

type VirtualNetworkGatewaySkuNamePtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuNamePtrOutput() VirtualNetworkGatewaySkuNamePtrOutput
	ToVirtualNetworkGatewaySkuNamePtrOutputWithContext(context.Context) VirtualNetworkGatewaySkuNamePtrOutput
}

type virtualNetworkGatewaySkuNamePtr string

func VirtualNetworkGatewaySkuNamePtr(v string) VirtualNetworkGatewaySkuNamePtrInput {
	return (*virtualNetworkGatewaySkuNamePtr)(&v)
}

func (*virtualNetworkGatewaySkuNamePtr) ElementType() reflect.Type {
	return virtualNetworkGatewaySkuNamePtrType
}

func (in *virtualNetworkGatewaySkuNamePtr) ToVirtualNetworkGatewaySkuNamePtrOutput() VirtualNetworkGatewaySkuNamePtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkGatewaySkuNamePtrOutput)
}

func (in *virtualNetworkGatewaySkuNamePtr) ToVirtualNetworkGatewaySkuNamePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkGatewaySkuNamePtrOutput)
}

type VirtualNetworkGatewaySkuTier string

const (
	VirtualNetworkGatewaySkuTierBasic            = VirtualNetworkGatewaySkuTier("Basic")
	VirtualNetworkGatewaySkuTierHighPerformance  = VirtualNetworkGatewaySkuTier("HighPerformance")
	VirtualNetworkGatewaySkuTierStandard         = VirtualNetworkGatewaySkuTier("Standard")
	VirtualNetworkGatewaySkuTierUltraPerformance = VirtualNetworkGatewaySkuTier("UltraPerformance")
	VirtualNetworkGatewaySkuTierVpnGw1           = VirtualNetworkGatewaySkuTier("VpnGw1")
	VirtualNetworkGatewaySkuTierVpnGw2           = VirtualNetworkGatewaySkuTier("VpnGw2")
	VirtualNetworkGatewaySkuTierVpnGw3           = VirtualNetworkGatewaySkuTier("VpnGw3")
)

func (VirtualNetworkGatewaySkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySkuTier)(nil)).Elem()
}

func (e VirtualNetworkGatewaySkuTier) ToVirtualNetworkGatewaySkuTierOutput() VirtualNetworkGatewaySkuTierOutput {
	return pulumi.ToOutput(e).(VirtualNetworkGatewaySkuTierOutput)
}

func (e VirtualNetworkGatewaySkuTier) ToVirtualNetworkGatewaySkuTierOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkGatewaySkuTierOutput)
}

func (e VirtualNetworkGatewaySkuTier) ToVirtualNetworkGatewaySkuTierPtrOutput() VirtualNetworkGatewaySkuTierPtrOutput {
	return e.ToVirtualNetworkGatewaySkuTierPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewaySkuTier) ToVirtualNetworkGatewaySkuTierPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuTierPtrOutput {
	return VirtualNetworkGatewaySkuTier(e).ToVirtualNetworkGatewaySkuTierOutputWithContext(ctx).ToVirtualNetworkGatewaySkuTierPtrOutputWithContext(ctx)
}

func (e VirtualNetworkGatewaySkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewaySkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewaySkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewaySkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkGatewaySkuTierOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySkuTier)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuTierOutput) ToVirtualNetworkGatewaySkuTierOutput() VirtualNetworkGatewaySkuTierOutput {
	return o
}

func (o VirtualNetworkGatewaySkuTierOutput) ToVirtualNetworkGatewaySkuTierOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuTierOutput {
	return o
}

func (o VirtualNetworkGatewaySkuTierOutput) ToVirtualNetworkGatewaySkuTierPtrOutput() VirtualNetworkGatewaySkuTierPtrOutput {
	return o.ToVirtualNetworkGatewaySkuTierPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuTierOutput) ToVirtualNetworkGatewaySkuTierPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewaySkuTier) *VirtualNetworkGatewaySkuTier {
		return &v
	}).(VirtualNetworkGatewaySkuTierPtrOutput)
}

func (o VirtualNetworkGatewaySkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewaySkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewaySkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewaySkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuTierPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySkuTier)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuTierPtrOutput) ToVirtualNetworkGatewaySkuTierPtrOutput() VirtualNetworkGatewaySkuTierPtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuTierPtrOutput) ToVirtualNetworkGatewaySkuTierPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuTierPtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuTierPtrOutput) Elem() VirtualNetworkGatewaySkuTierOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuTier) VirtualNetworkGatewaySkuTier {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewaySkuTier
		return ret
	}).(VirtualNetworkGatewaySkuTierOutput)
}

func (o VirtualNetworkGatewaySkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkGatewaySkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkGatewaySkuTierInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuTierOutput() VirtualNetworkGatewaySkuTierOutput
	ToVirtualNetworkGatewaySkuTierOutputWithContext(context.Context) VirtualNetworkGatewaySkuTierOutput
}

var virtualNetworkGatewaySkuTierPtrType = reflect.TypeOf((**VirtualNetworkGatewaySkuTier)(nil)).Elem()

type VirtualNetworkGatewaySkuTierPtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuTierPtrOutput() VirtualNetworkGatewaySkuTierPtrOutput
	ToVirtualNetworkGatewaySkuTierPtrOutputWithContext(context.Context) VirtualNetworkGatewaySkuTierPtrOutput
}

type virtualNetworkGatewaySkuTierPtr string

func VirtualNetworkGatewaySkuTierPtr(v string) VirtualNetworkGatewaySkuTierPtrInput {
	return (*virtualNetworkGatewaySkuTierPtr)(&v)
}

func (*virtualNetworkGatewaySkuTierPtr) ElementType() reflect.Type {
	return virtualNetworkGatewaySkuTierPtrType
}

func (in *virtualNetworkGatewaySkuTierPtr) ToVirtualNetworkGatewaySkuTierPtrOutput() VirtualNetworkGatewaySkuTierPtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkGatewaySkuTierPtrOutput)
}

func (in *virtualNetworkGatewaySkuTierPtr) ToVirtualNetworkGatewaySkuTierPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkGatewaySkuTierPtrOutput)
}

type VirtualNetworkGatewayTypeEnum string

const (
	VirtualNetworkGatewayTypeEnumVpn          = VirtualNetworkGatewayTypeEnum("Vpn")
	VirtualNetworkGatewayTypeEnumExpressRoute = VirtualNetworkGatewayTypeEnum("ExpressRoute")
)

func (VirtualNetworkGatewayTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayTypeEnum)(nil)).Elem()
}

func (e VirtualNetworkGatewayTypeEnum) ToVirtualNetworkGatewayTypeEnumOutput() VirtualNetworkGatewayTypeEnumOutput {
	return pulumi.ToOutput(e).(VirtualNetworkGatewayTypeEnumOutput)
}

func (e VirtualNetworkGatewayTypeEnum) ToVirtualNetworkGatewayTypeEnumOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkGatewayTypeEnumOutput)
}

func (e VirtualNetworkGatewayTypeEnum) ToVirtualNetworkGatewayTypeEnumPtrOutput() VirtualNetworkGatewayTypeEnumPtrOutput {
	return e.ToVirtualNetworkGatewayTypeEnumPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewayTypeEnum) ToVirtualNetworkGatewayTypeEnumPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeEnumPtrOutput {
	return VirtualNetworkGatewayTypeEnum(e).ToVirtualNetworkGatewayTypeEnumOutputWithContext(ctx).ToVirtualNetworkGatewayTypeEnumPtrOutputWithContext(ctx)
}

func (e VirtualNetworkGatewayTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewayTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewayTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewayTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkGatewayTypeEnumOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayTypeEnum)(nil)).Elem()
}

func (o VirtualNetworkGatewayTypeEnumOutput) ToVirtualNetworkGatewayTypeEnumOutput() VirtualNetworkGatewayTypeEnumOutput {
	return o
}

func (o VirtualNetworkGatewayTypeEnumOutput) ToVirtualNetworkGatewayTypeEnumOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeEnumOutput {
	return o
}

func (o VirtualNetworkGatewayTypeEnumOutput) ToVirtualNetworkGatewayTypeEnumPtrOutput() VirtualNetworkGatewayTypeEnumPtrOutput {
	return o.ToVirtualNetworkGatewayTypeEnumPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayTypeEnumOutput) ToVirtualNetworkGatewayTypeEnumPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewayTypeEnum) *VirtualNetworkGatewayTypeEnum {
		return &v
	}).(VirtualNetworkGatewayTypeEnumPtrOutput)
}

func (o VirtualNetworkGatewayTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewayTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewayTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayTypeEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayTypeEnum)(nil)).Elem()
}

func (o VirtualNetworkGatewayTypeEnumPtrOutput) ToVirtualNetworkGatewayTypeEnumPtrOutput() VirtualNetworkGatewayTypeEnumPtrOutput {
	return o
}

func (o VirtualNetworkGatewayTypeEnumPtrOutput) ToVirtualNetworkGatewayTypeEnumPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeEnumPtrOutput {
	return o
}

func (o VirtualNetworkGatewayTypeEnumPtrOutput) Elem() VirtualNetworkGatewayTypeEnumOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayTypeEnum) VirtualNetworkGatewayTypeEnum {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayTypeEnum
		return ret
	}).(VirtualNetworkGatewayTypeEnumOutput)
}

func (o VirtualNetworkGatewayTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkGatewayTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkGatewayTypeEnumInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayTypeEnumOutput() VirtualNetworkGatewayTypeEnumOutput
	ToVirtualNetworkGatewayTypeEnumOutputWithContext(context.Context) VirtualNetworkGatewayTypeEnumOutput
}

var virtualNetworkGatewayTypeEnumPtrType = reflect.TypeOf((**VirtualNetworkGatewayTypeEnum)(nil)).Elem()

type VirtualNetworkGatewayTypeEnumPtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayTypeEnumPtrOutput() VirtualNetworkGatewayTypeEnumPtrOutput
	ToVirtualNetworkGatewayTypeEnumPtrOutputWithContext(context.Context) VirtualNetworkGatewayTypeEnumPtrOutput
}

type virtualNetworkGatewayTypeEnumPtr string

func VirtualNetworkGatewayTypeEnumPtr(v string) VirtualNetworkGatewayTypeEnumPtrInput {
	return (*virtualNetworkGatewayTypeEnumPtr)(&v)
}

func (*virtualNetworkGatewayTypeEnumPtr) ElementType() reflect.Type {
	return virtualNetworkGatewayTypeEnumPtrType
}

func (in *virtualNetworkGatewayTypeEnumPtr) ToVirtualNetworkGatewayTypeEnumPtrOutput() VirtualNetworkGatewayTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkGatewayTypeEnumPtrOutput)
}

func (in *virtualNetworkGatewayTypeEnumPtr) ToVirtualNetworkGatewayTypeEnumPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkGatewayTypeEnumPtrOutput)
}

type VirtualNetworkPeeringStateEnum string

const (
	VirtualNetworkPeeringStateEnumInitiated    = VirtualNetworkPeeringStateEnum("Initiated")
	VirtualNetworkPeeringStateEnumConnected    = VirtualNetworkPeeringStateEnum("Connected")
	VirtualNetworkPeeringStateEnumDisconnected = VirtualNetworkPeeringStateEnum("Disconnected")
)

func (VirtualNetworkPeeringStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringStateEnum)(nil)).Elem()
}

func (e VirtualNetworkPeeringStateEnum) ToVirtualNetworkPeeringStateEnumOutput() VirtualNetworkPeeringStateEnumOutput {
	return pulumi.ToOutput(e).(VirtualNetworkPeeringStateEnumOutput)
}

func (e VirtualNetworkPeeringStateEnum) ToVirtualNetworkPeeringStateEnumOutputWithContext(ctx context.Context) VirtualNetworkPeeringStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkPeeringStateEnumOutput)
}

func (e VirtualNetworkPeeringStateEnum) ToVirtualNetworkPeeringStateEnumPtrOutput() VirtualNetworkPeeringStateEnumPtrOutput {
	return e.ToVirtualNetworkPeeringStateEnumPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkPeeringStateEnum) ToVirtualNetworkPeeringStateEnumPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringStateEnumPtrOutput {
	return VirtualNetworkPeeringStateEnum(e).ToVirtualNetworkPeeringStateEnumOutputWithContext(ctx).ToVirtualNetworkPeeringStateEnumPtrOutputWithContext(ctx)
}

func (e VirtualNetworkPeeringStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkPeeringStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkPeeringStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkPeeringStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkPeeringStateEnumOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringStateEnum)(nil)).Elem()
}

func (o VirtualNetworkPeeringStateEnumOutput) ToVirtualNetworkPeeringStateEnumOutput() VirtualNetworkPeeringStateEnumOutput {
	return o
}

func (o VirtualNetworkPeeringStateEnumOutput) ToVirtualNetworkPeeringStateEnumOutputWithContext(ctx context.Context) VirtualNetworkPeeringStateEnumOutput {
	return o
}

func (o VirtualNetworkPeeringStateEnumOutput) ToVirtualNetworkPeeringStateEnumPtrOutput() VirtualNetworkPeeringStateEnumPtrOutput {
	return o.ToVirtualNetworkPeeringStateEnumPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringStateEnumOutput) ToVirtualNetworkPeeringStateEnumPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPeeringStateEnum) *VirtualNetworkPeeringStateEnum {
		return &v
	}).(VirtualNetworkPeeringStateEnumPtrOutput)
}

func (o VirtualNetworkPeeringStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkPeeringStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkPeeringStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkPeeringStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringStateEnumPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringStateEnum)(nil)).Elem()
}

func (o VirtualNetworkPeeringStateEnumPtrOutput) ToVirtualNetworkPeeringStateEnumPtrOutput() VirtualNetworkPeeringStateEnumPtrOutput {
	return o
}

func (o VirtualNetworkPeeringStateEnumPtrOutput) ToVirtualNetworkPeeringStateEnumPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringStateEnumPtrOutput {
	return o
}

func (o VirtualNetworkPeeringStateEnumPtrOutput) Elem() VirtualNetworkPeeringStateEnumOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringStateEnum) VirtualNetworkPeeringStateEnum {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPeeringStateEnum
		return ret
	}).(VirtualNetworkPeeringStateEnumOutput)
}

func (o VirtualNetworkPeeringStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkPeeringStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkPeeringStateEnumInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringStateEnumOutput() VirtualNetworkPeeringStateEnumOutput
	ToVirtualNetworkPeeringStateEnumOutputWithContext(context.Context) VirtualNetworkPeeringStateEnumOutput
}

var virtualNetworkPeeringStateEnumPtrType = reflect.TypeOf((**VirtualNetworkPeeringStateEnum)(nil)).Elem()

type VirtualNetworkPeeringStateEnumPtrInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringStateEnumPtrOutput() VirtualNetworkPeeringStateEnumPtrOutput
	ToVirtualNetworkPeeringStateEnumPtrOutputWithContext(context.Context) VirtualNetworkPeeringStateEnumPtrOutput
}

type virtualNetworkPeeringStateEnumPtr string

func VirtualNetworkPeeringStateEnumPtr(v string) VirtualNetworkPeeringStateEnumPtrInput {
	return (*virtualNetworkPeeringStateEnumPtr)(&v)
}

func (*virtualNetworkPeeringStateEnumPtr) ElementType() reflect.Type {
	return virtualNetworkPeeringStateEnumPtrType
}

func (in *virtualNetworkPeeringStateEnumPtr) ToVirtualNetworkPeeringStateEnumPtrOutput() VirtualNetworkPeeringStateEnumPtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkPeeringStateEnumPtrOutput)
}

func (in *virtualNetworkPeeringStateEnumPtr) ToVirtualNetworkPeeringStateEnumPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkPeeringStateEnumPtrOutput)
}

type VpnClientProtocol string

const (
	VpnClientProtocolIkeV2 = VpnClientProtocol("IkeV2")
	VpnClientProtocolSSTP  = VpnClientProtocol("SSTP")
)

func (VpnClientProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientProtocol)(nil)).Elem()
}

func (e VpnClientProtocol) ToVpnClientProtocolOutput() VpnClientProtocolOutput {
	return pulumi.ToOutput(e).(VpnClientProtocolOutput)
}

func (e VpnClientProtocol) ToVpnClientProtocolOutputWithContext(ctx context.Context) VpnClientProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VpnClientProtocolOutput)
}

func (e VpnClientProtocol) ToVpnClientProtocolPtrOutput() VpnClientProtocolPtrOutput {
	return e.ToVpnClientProtocolPtrOutputWithContext(context.Background())
}

func (e VpnClientProtocol) ToVpnClientProtocolPtrOutputWithContext(ctx context.Context) VpnClientProtocolPtrOutput {
	return VpnClientProtocol(e).ToVpnClientProtocolOutputWithContext(ctx).ToVpnClientProtocolPtrOutputWithContext(ctx)
}

func (e VpnClientProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnClientProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnClientProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VpnClientProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VpnClientProtocolOutput struct{ *pulumi.OutputState }

func (VpnClientProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientProtocol)(nil)).Elem()
}

func (o VpnClientProtocolOutput) ToVpnClientProtocolOutput() VpnClientProtocolOutput {
	return o
}

func (o VpnClientProtocolOutput) ToVpnClientProtocolOutputWithContext(ctx context.Context) VpnClientProtocolOutput {
	return o
}

func (o VpnClientProtocolOutput) ToVpnClientProtocolPtrOutput() VpnClientProtocolPtrOutput {
	return o.ToVpnClientProtocolPtrOutputWithContext(context.Background())
}

func (o VpnClientProtocolOutput) ToVpnClientProtocolPtrOutputWithContext(ctx context.Context) VpnClientProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnClientProtocol) *VpnClientProtocol {
		return &v
	}).(VpnClientProtocolPtrOutput)
}

func (o VpnClientProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VpnClientProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnClientProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VpnClientProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnClientProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnClientProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VpnClientProtocolPtrOutput struct{ *pulumi.OutputState }

func (VpnClientProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientProtocol)(nil)).Elem()
}

func (o VpnClientProtocolPtrOutput) ToVpnClientProtocolPtrOutput() VpnClientProtocolPtrOutput {
	return o
}

func (o VpnClientProtocolPtrOutput) ToVpnClientProtocolPtrOutputWithContext(ctx context.Context) VpnClientProtocolPtrOutput {
	return o
}

func (o VpnClientProtocolPtrOutput) Elem() VpnClientProtocolOutput {
	return o.ApplyT(func(v *VpnClientProtocol) VpnClientProtocol {
		if v != nil {
			return *v
		}
		var ret VpnClientProtocol
		return ret
	}).(VpnClientProtocolOutput)
}

func (o VpnClientProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnClientProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VpnClientProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VpnClientProtocolInput interface {
	pulumi.Input

	ToVpnClientProtocolOutput() VpnClientProtocolOutput
	ToVpnClientProtocolOutputWithContext(context.Context) VpnClientProtocolOutput
}

var vpnClientProtocolPtrType = reflect.TypeOf((**VpnClientProtocol)(nil)).Elem()

type VpnClientProtocolPtrInput interface {
	pulumi.Input

	ToVpnClientProtocolPtrOutput() VpnClientProtocolPtrOutput
	ToVpnClientProtocolPtrOutputWithContext(context.Context) VpnClientProtocolPtrOutput
}

type vpnClientProtocolPtr string

func VpnClientProtocolPtr(v string) VpnClientProtocolPtrInput {
	return (*vpnClientProtocolPtr)(&v)
}

func (*vpnClientProtocolPtr) ElementType() reflect.Type {
	return vpnClientProtocolPtrType
}

func (in *vpnClientProtocolPtr) ToVpnClientProtocolPtrOutput() VpnClientProtocolPtrOutput {
	return pulumi.ToOutput(in).(VpnClientProtocolPtrOutput)
}

func (in *vpnClientProtocolPtr) ToVpnClientProtocolPtrOutputWithContext(ctx context.Context) VpnClientProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VpnClientProtocolPtrOutput)
}

type VpnType string

const (
	VpnTypePolicyBased = VpnType("PolicyBased")
	VpnTypeRouteBased  = VpnType("RouteBased")
)

func (VpnType) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnType)(nil)).Elem()
}

func (e VpnType) ToVpnTypeOutput() VpnTypeOutput {
	return pulumi.ToOutput(e).(VpnTypeOutput)
}

func (e VpnType) ToVpnTypeOutputWithContext(ctx context.Context) VpnTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VpnTypeOutput)
}

func (e VpnType) ToVpnTypePtrOutput() VpnTypePtrOutput {
	return e.ToVpnTypePtrOutputWithContext(context.Background())
}

func (e VpnType) ToVpnTypePtrOutputWithContext(ctx context.Context) VpnTypePtrOutput {
	return VpnType(e).ToVpnTypeOutputWithContext(ctx).ToVpnTypePtrOutputWithContext(ctx)
}

func (e VpnType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VpnType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VpnTypeOutput struct{ *pulumi.OutputState }

func (VpnTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnType)(nil)).Elem()
}

func (o VpnTypeOutput) ToVpnTypeOutput() VpnTypeOutput {
	return o
}

func (o VpnTypeOutput) ToVpnTypeOutputWithContext(ctx context.Context) VpnTypeOutput {
	return o
}

func (o VpnTypeOutput) ToVpnTypePtrOutput() VpnTypePtrOutput {
	return o.ToVpnTypePtrOutputWithContext(context.Background())
}

func (o VpnTypeOutput) ToVpnTypePtrOutputWithContext(ctx context.Context) VpnTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnType) *VpnType {
		return &v
	}).(VpnTypePtrOutput)
}

func (o VpnTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VpnTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VpnTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VpnTypePtrOutput struct{ *pulumi.OutputState }

func (VpnTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnType)(nil)).Elem()
}

func (o VpnTypePtrOutput) ToVpnTypePtrOutput() VpnTypePtrOutput {
	return o
}

func (o VpnTypePtrOutput) ToVpnTypePtrOutputWithContext(ctx context.Context) VpnTypePtrOutput {
	return o
}

func (o VpnTypePtrOutput) Elem() VpnTypeOutput {
	return o.ApplyT(func(v *VpnType) VpnType {
		if v != nil {
			return *v
		}
		var ret VpnType
		return ret
	}).(VpnTypeOutput)
}

func (o VpnTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VpnType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VpnTypeInput interface {
	pulumi.Input

	ToVpnTypeOutput() VpnTypeOutput
	ToVpnTypeOutputWithContext(context.Context) VpnTypeOutput
}

var vpnTypePtrType = reflect.TypeOf((**VpnType)(nil)).Elem()

type VpnTypePtrInput interface {
	pulumi.Input

	ToVpnTypePtrOutput() VpnTypePtrOutput
	ToVpnTypePtrOutputWithContext(context.Context) VpnTypePtrOutput
}

type vpnTypePtr string

func VpnTypePtr(v string) VpnTypePtrInput {
	return (*vpnTypePtr)(&v)
}

func (*vpnTypePtr) ElementType() reflect.Type {
	return vpnTypePtrType
}

func (in *vpnTypePtr) ToVpnTypePtrOutput() VpnTypePtrOutput {
	return pulumi.ToOutput(in).(VpnTypePtrOutput)
}

func (in *vpnTypePtr) ToVpnTypePtrOutputWithContext(ctx context.Context) VpnTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VpnTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessOutput{})
	pulumi.RegisterOutputType(AccessPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayCookieBasedAffinityOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayCookieBasedAffinityPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallModeOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallModePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProtocolOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProtocolPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRedirectTypeOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRedirectTypePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleTypeOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleTypePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuNameOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuNamePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCipherSuiteOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCipherSuitePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyNameOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyNamePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyTypeOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyTypePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslProtocolOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslProtocolPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayTierOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayTierPtrOutput{})
	pulumi.RegisterOutputType(AuthorizationUseStatusOutput{})
	pulumi.RegisterOutputType(AuthorizationUseStatusPtrOutput{})
	pulumi.RegisterOutputType(DhGroupOutput{})
	pulumi.RegisterOutputType(DhGroupPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringAdvertisedPublicPrefixStatePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringStateEnumOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringStateEnumPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringTypeEnumOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuFamilyOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuFamilyPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuTierOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuTierPtrOutput{})
	pulumi.RegisterOutputType(IPAllocationMethodOutput{})
	pulumi.RegisterOutputType(IPAllocationMethodPtrOutput{})
	pulumi.RegisterOutputType(IPVersionOutput{})
	pulumi.RegisterOutputType(IPVersionPtrOutput{})
	pulumi.RegisterOutputType(IkeEncryptionOutput{})
	pulumi.RegisterOutputType(IkeEncryptionPtrOutput{})
	pulumi.RegisterOutputType(IkeIntegrityOutput{})
	pulumi.RegisterOutputType(IkeIntegrityPtrOutput{})
	pulumi.RegisterOutputType(IpsecEncryptionOutput{})
	pulumi.RegisterOutputType(IpsecEncryptionPtrOutput{})
	pulumi.RegisterOutputType(IpsecIntegrityOutput{})
	pulumi.RegisterOutputType(IpsecIntegrityPtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuNameOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuNamePtrOutput{})
	pulumi.RegisterOutputType(LoadDistributionOutput{})
	pulumi.RegisterOutputType(LoadDistributionPtrOutput{})
	pulumi.RegisterOutputType(PcProtocolOutput{})
	pulumi.RegisterOutputType(PcProtocolPtrOutput{})
	pulumi.RegisterOutputType(PfsGroupOutput{})
	pulumi.RegisterOutputType(PfsGroupPtrOutput{})
	pulumi.RegisterOutputType(ProbeProtocolOutput{})
	pulumi.RegisterOutputType(ProbeProtocolPtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuNameOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuNamePtrOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleTypeEnumOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(RouteNextHopTypeOutput{})
	pulumi.RegisterOutputType(RouteNextHopTypePtrOutput{})
	pulumi.RegisterOutputType(SecurityRuleAccessOutput{})
	pulumi.RegisterOutputType(SecurityRuleAccessPtrOutput{})
	pulumi.RegisterOutputType(SecurityRuleDirectionOutput{})
	pulumi.RegisterOutputType(SecurityRuleDirectionPtrOutput{})
	pulumi.RegisterOutputType(SecurityRuleProtocolOutput{})
	pulumi.RegisterOutputType(SecurityRuleProtocolPtrOutput{})
	pulumi.RegisterOutputType(ServiceProviderProvisioningStateOutput{})
	pulumi.RegisterOutputType(ServiceProviderProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(TransportProtocolOutput{})
	pulumi.RegisterOutputType(TransportProtocolPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionTypePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuNameOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuNamePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuTierOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuTierPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypeEnumOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringStateEnumOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringStateEnumPtrOutput{})
	pulumi.RegisterOutputType(VpnClientProtocolOutput{})
	pulumi.RegisterOutputType(VpnClientProtocolPtrOutput{})
	pulumi.RegisterOutputType(VpnTypeOutput{})
	pulumi.RegisterOutputType(VpnTypePtrOutput{})
}
