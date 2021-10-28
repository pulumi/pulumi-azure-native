


package v20210301

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

type ApplicationGatewayCustomErrorStatusCode string

const (
	ApplicationGatewayCustomErrorStatusCodeHttpStatus403 = ApplicationGatewayCustomErrorStatusCode("HttpStatus403")
	ApplicationGatewayCustomErrorStatusCodeHttpStatus502 = ApplicationGatewayCustomErrorStatusCode("HttpStatus502")
)

func (ApplicationGatewayCustomErrorStatusCode) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayCustomErrorStatusCode)(nil)).Elem()
}

func (e ApplicationGatewayCustomErrorStatusCode) ToApplicationGatewayCustomErrorStatusCodeOutput() ApplicationGatewayCustomErrorStatusCodeOutput {
	return pulumi.ToOutput(e).(ApplicationGatewayCustomErrorStatusCodeOutput)
}

func (e ApplicationGatewayCustomErrorStatusCode) ToApplicationGatewayCustomErrorStatusCodeOutputWithContext(ctx context.Context) ApplicationGatewayCustomErrorStatusCodeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewayCustomErrorStatusCodeOutput)
}

func (e ApplicationGatewayCustomErrorStatusCode) ToApplicationGatewayCustomErrorStatusCodePtrOutput() ApplicationGatewayCustomErrorStatusCodePtrOutput {
	return e.ToApplicationGatewayCustomErrorStatusCodePtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayCustomErrorStatusCode) ToApplicationGatewayCustomErrorStatusCodePtrOutputWithContext(ctx context.Context) ApplicationGatewayCustomErrorStatusCodePtrOutput {
	return ApplicationGatewayCustomErrorStatusCode(e).ToApplicationGatewayCustomErrorStatusCodeOutputWithContext(ctx).ToApplicationGatewayCustomErrorStatusCodePtrOutputWithContext(ctx)
}

func (e ApplicationGatewayCustomErrorStatusCode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayCustomErrorStatusCode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayCustomErrorStatusCode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayCustomErrorStatusCode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewayCustomErrorStatusCodeOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayCustomErrorStatusCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayCustomErrorStatusCode)(nil)).Elem()
}

func (o ApplicationGatewayCustomErrorStatusCodeOutput) ToApplicationGatewayCustomErrorStatusCodeOutput() ApplicationGatewayCustomErrorStatusCodeOutput {
	return o
}

func (o ApplicationGatewayCustomErrorStatusCodeOutput) ToApplicationGatewayCustomErrorStatusCodeOutputWithContext(ctx context.Context) ApplicationGatewayCustomErrorStatusCodeOutput {
	return o
}

func (o ApplicationGatewayCustomErrorStatusCodeOutput) ToApplicationGatewayCustomErrorStatusCodePtrOutput() ApplicationGatewayCustomErrorStatusCodePtrOutput {
	return o.ToApplicationGatewayCustomErrorStatusCodePtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayCustomErrorStatusCodeOutput) ToApplicationGatewayCustomErrorStatusCodePtrOutputWithContext(ctx context.Context) ApplicationGatewayCustomErrorStatusCodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayCustomErrorStatusCode) *ApplicationGatewayCustomErrorStatusCode {
		return &v
	}).(ApplicationGatewayCustomErrorStatusCodePtrOutput)
}

func (o ApplicationGatewayCustomErrorStatusCodeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewayCustomErrorStatusCodeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayCustomErrorStatusCode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewayCustomErrorStatusCodeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayCustomErrorStatusCodeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayCustomErrorStatusCode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayCustomErrorStatusCodePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayCustomErrorStatusCodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayCustomErrorStatusCode)(nil)).Elem()
}

func (o ApplicationGatewayCustomErrorStatusCodePtrOutput) ToApplicationGatewayCustomErrorStatusCodePtrOutput() ApplicationGatewayCustomErrorStatusCodePtrOutput {
	return o
}

func (o ApplicationGatewayCustomErrorStatusCodePtrOutput) ToApplicationGatewayCustomErrorStatusCodePtrOutputWithContext(ctx context.Context) ApplicationGatewayCustomErrorStatusCodePtrOutput {
	return o
}

func (o ApplicationGatewayCustomErrorStatusCodePtrOutput) Elem() ApplicationGatewayCustomErrorStatusCodeOutput {
	return o.ApplyT(func(v *ApplicationGatewayCustomErrorStatusCode) ApplicationGatewayCustomErrorStatusCode {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayCustomErrorStatusCode
		return ret
	}).(ApplicationGatewayCustomErrorStatusCodeOutput)
}

func (o ApplicationGatewayCustomErrorStatusCodePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayCustomErrorStatusCodePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewayCustomErrorStatusCode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewayCustomErrorStatusCodeInput interface {
	pulumi.Input

	ToApplicationGatewayCustomErrorStatusCodeOutput() ApplicationGatewayCustomErrorStatusCodeOutput
	ToApplicationGatewayCustomErrorStatusCodeOutputWithContext(context.Context) ApplicationGatewayCustomErrorStatusCodeOutput
}

var applicationGatewayCustomErrorStatusCodePtrType = reflect.TypeOf((**ApplicationGatewayCustomErrorStatusCode)(nil)).Elem()

type ApplicationGatewayCustomErrorStatusCodePtrInput interface {
	pulumi.Input

	ToApplicationGatewayCustomErrorStatusCodePtrOutput() ApplicationGatewayCustomErrorStatusCodePtrOutput
	ToApplicationGatewayCustomErrorStatusCodePtrOutputWithContext(context.Context) ApplicationGatewayCustomErrorStatusCodePtrOutput
}

type applicationGatewayCustomErrorStatusCodePtr string

func ApplicationGatewayCustomErrorStatusCodePtr(v string) ApplicationGatewayCustomErrorStatusCodePtrInput {
	return (*applicationGatewayCustomErrorStatusCodePtr)(&v)
}

func (*applicationGatewayCustomErrorStatusCodePtr) ElementType() reflect.Type {
	return applicationGatewayCustomErrorStatusCodePtrType
}

func (in *applicationGatewayCustomErrorStatusCodePtr) ToApplicationGatewayCustomErrorStatusCodePtrOutput() ApplicationGatewayCustomErrorStatusCodePtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewayCustomErrorStatusCodePtrOutput)
}

func (in *applicationGatewayCustomErrorStatusCodePtr) ToApplicationGatewayCustomErrorStatusCodePtrOutputWithContext(ctx context.Context) ApplicationGatewayCustomErrorStatusCodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewayCustomErrorStatusCodePtrOutput)
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

type ApplicationGatewayLoadDistributionAlgorithm string

const (
	ApplicationGatewayLoadDistributionAlgorithmRoundRobin       = ApplicationGatewayLoadDistributionAlgorithm("RoundRobin")
	ApplicationGatewayLoadDistributionAlgorithmLeastConnections = ApplicationGatewayLoadDistributionAlgorithm("LeastConnections")
	ApplicationGatewayLoadDistributionAlgorithmIpHash           = ApplicationGatewayLoadDistributionAlgorithm("IpHash")
)

func (ApplicationGatewayLoadDistributionAlgorithm) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayLoadDistributionAlgorithm)(nil)).Elem()
}

func (e ApplicationGatewayLoadDistributionAlgorithm) ToApplicationGatewayLoadDistributionAlgorithmOutput() ApplicationGatewayLoadDistributionAlgorithmOutput {
	return pulumi.ToOutput(e).(ApplicationGatewayLoadDistributionAlgorithmOutput)
}

func (e ApplicationGatewayLoadDistributionAlgorithm) ToApplicationGatewayLoadDistributionAlgorithmOutputWithContext(ctx context.Context) ApplicationGatewayLoadDistributionAlgorithmOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGatewayLoadDistributionAlgorithmOutput)
}

func (e ApplicationGatewayLoadDistributionAlgorithm) ToApplicationGatewayLoadDistributionAlgorithmPtrOutput() ApplicationGatewayLoadDistributionAlgorithmPtrOutput {
	return e.ToApplicationGatewayLoadDistributionAlgorithmPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayLoadDistributionAlgorithm) ToApplicationGatewayLoadDistributionAlgorithmPtrOutputWithContext(ctx context.Context) ApplicationGatewayLoadDistributionAlgorithmPtrOutput {
	return ApplicationGatewayLoadDistributionAlgorithm(e).ToApplicationGatewayLoadDistributionAlgorithmOutputWithContext(ctx).ToApplicationGatewayLoadDistributionAlgorithmPtrOutputWithContext(ctx)
}

func (e ApplicationGatewayLoadDistributionAlgorithm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayLoadDistributionAlgorithm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGatewayLoadDistributionAlgorithm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGatewayLoadDistributionAlgorithm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGatewayLoadDistributionAlgorithmOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayLoadDistributionAlgorithmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayLoadDistributionAlgorithm)(nil)).Elem()
}

func (o ApplicationGatewayLoadDistributionAlgorithmOutput) ToApplicationGatewayLoadDistributionAlgorithmOutput() ApplicationGatewayLoadDistributionAlgorithmOutput {
	return o
}

func (o ApplicationGatewayLoadDistributionAlgorithmOutput) ToApplicationGatewayLoadDistributionAlgorithmOutputWithContext(ctx context.Context) ApplicationGatewayLoadDistributionAlgorithmOutput {
	return o
}

func (o ApplicationGatewayLoadDistributionAlgorithmOutput) ToApplicationGatewayLoadDistributionAlgorithmPtrOutput() ApplicationGatewayLoadDistributionAlgorithmPtrOutput {
	return o.ToApplicationGatewayLoadDistributionAlgorithmPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayLoadDistributionAlgorithmOutput) ToApplicationGatewayLoadDistributionAlgorithmPtrOutputWithContext(ctx context.Context) ApplicationGatewayLoadDistributionAlgorithmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayLoadDistributionAlgorithm) *ApplicationGatewayLoadDistributionAlgorithm {
		return &v
	}).(ApplicationGatewayLoadDistributionAlgorithmPtrOutput)
}

func (o ApplicationGatewayLoadDistributionAlgorithmOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGatewayLoadDistributionAlgorithmOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayLoadDistributionAlgorithm) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGatewayLoadDistributionAlgorithmOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayLoadDistributionAlgorithmOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGatewayLoadDistributionAlgorithm) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayLoadDistributionAlgorithmPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayLoadDistributionAlgorithmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayLoadDistributionAlgorithm)(nil)).Elem()
}

func (o ApplicationGatewayLoadDistributionAlgorithmPtrOutput) ToApplicationGatewayLoadDistributionAlgorithmPtrOutput() ApplicationGatewayLoadDistributionAlgorithmPtrOutput {
	return o
}

func (o ApplicationGatewayLoadDistributionAlgorithmPtrOutput) ToApplicationGatewayLoadDistributionAlgorithmPtrOutputWithContext(ctx context.Context) ApplicationGatewayLoadDistributionAlgorithmPtrOutput {
	return o
}

func (o ApplicationGatewayLoadDistributionAlgorithmPtrOutput) Elem() ApplicationGatewayLoadDistributionAlgorithmOutput {
	return o.ApplyT(func(v *ApplicationGatewayLoadDistributionAlgorithm) ApplicationGatewayLoadDistributionAlgorithm {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayLoadDistributionAlgorithm
		return ret
	}).(ApplicationGatewayLoadDistributionAlgorithmOutput)
}

func (o ApplicationGatewayLoadDistributionAlgorithmPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayLoadDistributionAlgorithmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGatewayLoadDistributionAlgorithm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGatewayLoadDistributionAlgorithmInput interface {
	pulumi.Input

	ToApplicationGatewayLoadDistributionAlgorithmOutput() ApplicationGatewayLoadDistributionAlgorithmOutput
	ToApplicationGatewayLoadDistributionAlgorithmOutputWithContext(context.Context) ApplicationGatewayLoadDistributionAlgorithmOutput
}

var applicationGatewayLoadDistributionAlgorithmPtrType = reflect.TypeOf((**ApplicationGatewayLoadDistributionAlgorithm)(nil)).Elem()

type ApplicationGatewayLoadDistributionAlgorithmPtrInput interface {
	pulumi.Input

	ToApplicationGatewayLoadDistributionAlgorithmPtrOutput() ApplicationGatewayLoadDistributionAlgorithmPtrOutput
	ToApplicationGatewayLoadDistributionAlgorithmPtrOutputWithContext(context.Context) ApplicationGatewayLoadDistributionAlgorithmPtrOutput
}

type applicationGatewayLoadDistributionAlgorithmPtr string

func ApplicationGatewayLoadDistributionAlgorithmPtr(v string) ApplicationGatewayLoadDistributionAlgorithmPtrInput {
	return (*applicationGatewayLoadDistributionAlgorithmPtr)(&v)
}

func (*applicationGatewayLoadDistributionAlgorithmPtr) ElementType() reflect.Type {
	return applicationGatewayLoadDistributionAlgorithmPtrType
}

func (in *applicationGatewayLoadDistributionAlgorithmPtr) ToApplicationGatewayLoadDistributionAlgorithmPtrOutput() ApplicationGatewayLoadDistributionAlgorithmPtrOutput {
	return pulumi.ToOutput(in).(ApplicationGatewayLoadDistributionAlgorithmPtrOutput)
}

func (in *applicationGatewayLoadDistributionAlgorithmPtr) ToApplicationGatewayLoadDistributionAlgorithmPtrOutputWithContext(ctx context.Context) ApplicationGatewayLoadDistributionAlgorithmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGatewayLoadDistributionAlgorithmPtrOutput)
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
	ApplicationGatewaySkuName_Standard_v2     = ApplicationGatewaySkuName("Standard_v2")
	ApplicationGatewaySkuName_WAF_v2          = ApplicationGatewaySkuName("WAF_v2")
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
	ApplicationGatewayTierStandard     = ApplicationGatewayTier("Standard")
	ApplicationGatewayTierWAF          = ApplicationGatewayTier("WAF")
	ApplicationGatewayTier_Standard_v2 = ApplicationGatewayTier("Standard_v2")
	ApplicationGatewayTier_WAF_v2      = ApplicationGatewayTier("WAF_v2")
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

type AzureFirewallApplicationRuleProtocolType string

const (
	AzureFirewallApplicationRuleProtocolTypeHttp  = AzureFirewallApplicationRuleProtocolType("Http")
	AzureFirewallApplicationRuleProtocolTypeHttps = AzureFirewallApplicationRuleProtocolType("Https")
	AzureFirewallApplicationRuleProtocolTypeMssql = AzureFirewallApplicationRuleProtocolType("Mssql")
)

func (AzureFirewallApplicationRuleProtocolType) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallApplicationRuleProtocolType)(nil)).Elem()
}

func (e AzureFirewallApplicationRuleProtocolType) ToAzureFirewallApplicationRuleProtocolTypeOutput() AzureFirewallApplicationRuleProtocolTypeOutput {
	return pulumi.ToOutput(e).(AzureFirewallApplicationRuleProtocolTypeOutput)
}

func (e AzureFirewallApplicationRuleProtocolType) ToAzureFirewallApplicationRuleProtocolTypeOutputWithContext(ctx context.Context) AzureFirewallApplicationRuleProtocolTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureFirewallApplicationRuleProtocolTypeOutput)
}

func (e AzureFirewallApplicationRuleProtocolType) ToAzureFirewallApplicationRuleProtocolTypePtrOutput() AzureFirewallApplicationRuleProtocolTypePtrOutput {
	return e.ToAzureFirewallApplicationRuleProtocolTypePtrOutputWithContext(context.Background())
}

func (e AzureFirewallApplicationRuleProtocolType) ToAzureFirewallApplicationRuleProtocolTypePtrOutputWithContext(ctx context.Context) AzureFirewallApplicationRuleProtocolTypePtrOutput {
	return AzureFirewallApplicationRuleProtocolType(e).ToAzureFirewallApplicationRuleProtocolTypeOutputWithContext(ctx).ToAzureFirewallApplicationRuleProtocolTypePtrOutputWithContext(ctx)
}

func (e AzureFirewallApplicationRuleProtocolType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallApplicationRuleProtocolType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallApplicationRuleProtocolType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureFirewallApplicationRuleProtocolType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureFirewallApplicationRuleProtocolTypeOutput struct{ *pulumi.OutputState }

func (AzureFirewallApplicationRuleProtocolTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallApplicationRuleProtocolType)(nil)).Elem()
}

func (o AzureFirewallApplicationRuleProtocolTypeOutput) ToAzureFirewallApplicationRuleProtocolTypeOutput() AzureFirewallApplicationRuleProtocolTypeOutput {
	return o
}

func (o AzureFirewallApplicationRuleProtocolTypeOutput) ToAzureFirewallApplicationRuleProtocolTypeOutputWithContext(ctx context.Context) AzureFirewallApplicationRuleProtocolTypeOutput {
	return o
}

func (o AzureFirewallApplicationRuleProtocolTypeOutput) ToAzureFirewallApplicationRuleProtocolTypePtrOutput() AzureFirewallApplicationRuleProtocolTypePtrOutput {
	return o.ToAzureFirewallApplicationRuleProtocolTypePtrOutputWithContext(context.Background())
}

func (o AzureFirewallApplicationRuleProtocolTypeOutput) ToAzureFirewallApplicationRuleProtocolTypePtrOutputWithContext(ctx context.Context) AzureFirewallApplicationRuleProtocolTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFirewallApplicationRuleProtocolType) *AzureFirewallApplicationRuleProtocolType {
		return &v
	}).(AzureFirewallApplicationRuleProtocolTypePtrOutput)
}

func (o AzureFirewallApplicationRuleProtocolTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureFirewallApplicationRuleProtocolTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallApplicationRuleProtocolType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureFirewallApplicationRuleProtocolTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallApplicationRuleProtocolTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallApplicationRuleProtocolType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureFirewallApplicationRuleProtocolTypePtrOutput struct{ *pulumi.OutputState }

func (AzureFirewallApplicationRuleProtocolTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewallApplicationRuleProtocolType)(nil)).Elem()
}

func (o AzureFirewallApplicationRuleProtocolTypePtrOutput) ToAzureFirewallApplicationRuleProtocolTypePtrOutput() AzureFirewallApplicationRuleProtocolTypePtrOutput {
	return o
}

func (o AzureFirewallApplicationRuleProtocolTypePtrOutput) ToAzureFirewallApplicationRuleProtocolTypePtrOutputWithContext(ctx context.Context) AzureFirewallApplicationRuleProtocolTypePtrOutput {
	return o
}

func (o AzureFirewallApplicationRuleProtocolTypePtrOutput) Elem() AzureFirewallApplicationRuleProtocolTypeOutput {
	return o.ApplyT(func(v *AzureFirewallApplicationRuleProtocolType) AzureFirewallApplicationRuleProtocolType {
		if v != nil {
			return *v
		}
		var ret AzureFirewallApplicationRuleProtocolType
		return ret
	}).(AzureFirewallApplicationRuleProtocolTypeOutput)
}

func (o AzureFirewallApplicationRuleProtocolTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallApplicationRuleProtocolTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureFirewallApplicationRuleProtocolType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureFirewallApplicationRuleProtocolTypeInput interface {
	pulumi.Input

	ToAzureFirewallApplicationRuleProtocolTypeOutput() AzureFirewallApplicationRuleProtocolTypeOutput
	ToAzureFirewallApplicationRuleProtocolTypeOutputWithContext(context.Context) AzureFirewallApplicationRuleProtocolTypeOutput
}

var azureFirewallApplicationRuleProtocolTypePtrType = reflect.TypeOf((**AzureFirewallApplicationRuleProtocolType)(nil)).Elem()

type AzureFirewallApplicationRuleProtocolTypePtrInput interface {
	pulumi.Input

	ToAzureFirewallApplicationRuleProtocolTypePtrOutput() AzureFirewallApplicationRuleProtocolTypePtrOutput
	ToAzureFirewallApplicationRuleProtocolTypePtrOutputWithContext(context.Context) AzureFirewallApplicationRuleProtocolTypePtrOutput
}

type azureFirewallApplicationRuleProtocolTypePtr string

func AzureFirewallApplicationRuleProtocolTypePtr(v string) AzureFirewallApplicationRuleProtocolTypePtrInput {
	return (*azureFirewallApplicationRuleProtocolTypePtr)(&v)
}

func (*azureFirewallApplicationRuleProtocolTypePtr) ElementType() reflect.Type {
	return azureFirewallApplicationRuleProtocolTypePtrType
}

func (in *azureFirewallApplicationRuleProtocolTypePtr) ToAzureFirewallApplicationRuleProtocolTypePtrOutput() AzureFirewallApplicationRuleProtocolTypePtrOutput {
	return pulumi.ToOutput(in).(AzureFirewallApplicationRuleProtocolTypePtrOutput)
}

func (in *azureFirewallApplicationRuleProtocolTypePtr) ToAzureFirewallApplicationRuleProtocolTypePtrOutputWithContext(ctx context.Context) AzureFirewallApplicationRuleProtocolTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureFirewallApplicationRuleProtocolTypePtrOutput)
}

type AzureFirewallNatRCActionType string

const (
	AzureFirewallNatRCActionTypeSnat = AzureFirewallNatRCActionType("Snat")
	AzureFirewallNatRCActionTypeDnat = AzureFirewallNatRCActionType("Dnat")
)

func (AzureFirewallNatRCActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallNatRCActionType)(nil)).Elem()
}

func (e AzureFirewallNatRCActionType) ToAzureFirewallNatRCActionTypeOutput() AzureFirewallNatRCActionTypeOutput {
	return pulumi.ToOutput(e).(AzureFirewallNatRCActionTypeOutput)
}

func (e AzureFirewallNatRCActionType) ToAzureFirewallNatRCActionTypeOutputWithContext(ctx context.Context) AzureFirewallNatRCActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureFirewallNatRCActionTypeOutput)
}

func (e AzureFirewallNatRCActionType) ToAzureFirewallNatRCActionTypePtrOutput() AzureFirewallNatRCActionTypePtrOutput {
	return e.ToAzureFirewallNatRCActionTypePtrOutputWithContext(context.Background())
}

func (e AzureFirewallNatRCActionType) ToAzureFirewallNatRCActionTypePtrOutputWithContext(ctx context.Context) AzureFirewallNatRCActionTypePtrOutput {
	return AzureFirewallNatRCActionType(e).ToAzureFirewallNatRCActionTypeOutputWithContext(ctx).ToAzureFirewallNatRCActionTypePtrOutputWithContext(ctx)
}

func (e AzureFirewallNatRCActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallNatRCActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallNatRCActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureFirewallNatRCActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureFirewallNatRCActionTypeOutput struct{ *pulumi.OutputState }

func (AzureFirewallNatRCActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallNatRCActionType)(nil)).Elem()
}

func (o AzureFirewallNatRCActionTypeOutput) ToAzureFirewallNatRCActionTypeOutput() AzureFirewallNatRCActionTypeOutput {
	return o
}

func (o AzureFirewallNatRCActionTypeOutput) ToAzureFirewallNatRCActionTypeOutputWithContext(ctx context.Context) AzureFirewallNatRCActionTypeOutput {
	return o
}

func (o AzureFirewallNatRCActionTypeOutput) ToAzureFirewallNatRCActionTypePtrOutput() AzureFirewallNatRCActionTypePtrOutput {
	return o.ToAzureFirewallNatRCActionTypePtrOutputWithContext(context.Background())
}

func (o AzureFirewallNatRCActionTypeOutput) ToAzureFirewallNatRCActionTypePtrOutputWithContext(ctx context.Context) AzureFirewallNatRCActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFirewallNatRCActionType) *AzureFirewallNatRCActionType {
		return &v
	}).(AzureFirewallNatRCActionTypePtrOutput)
}

func (o AzureFirewallNatRCActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureFirewallNatRCActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallNatRCActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureFirewallNatRCActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallNatRCActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallNatRCActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureFirewallNatRCActionTypePtrOutput struct{ *pulumi.OutputState }

func (AzureFirewallNatRCActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewallNatRCActionType)(nil)).Elem()
}

func (o AzureFirewallNatRCActionTypePtrOutput) ToAzureFirewallNatRCActionTypePtrOutput() AzureFirewallNatRCActionTypePtrOutput {
	return o
}

func (o AzureFirewallNatRCActionTypePtrOutput) ToAzureFirewallNatRCActionTypePtrOutputWithContext(ctx context.Context) AzureFirewallNatRCActionTypePtrOutput {
	return o
}

func (o AzureFirewallNatRCActionTypePtrOutput) Elem() AzureFirewallNatRCActionTypeOutput {
	return o.ApplyT(func(v *AzureFirewallNatRCActionType) AzureFirewallNatRCActionType {
		if v != nil {
			return *v
		}
		var ret AzureFirewallNatRCActionType
		return ret
	}).(AzureFirewallNatRCActionTypeOutput)
}

func (o AzureFirewallNatRCActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallNatRCActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureFirewallNatRCActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureFirewallNatRCActionTypeInput interface {
	pulumi.Input

	ToAzureFirewallNatRCActionTypeOutput() AzureFirewallNatRCActionTypeOutput
	ToAzureFirewallNatRCActionTypeOutputWithContext(context.Context) AzureFirewallNatRCActionTypeOutput
}

var azureFirewallNatRCActionTypePtrType = reflect.TypeOf((**AzureFirewallNatRCActionType)(nil)).Elem()

type AzureFirewallNatRCActionTypePtrInput interface {
	pulumi.Input

	ToAzureFirewallNatRCActionTypePtrOutput() AzureFirewallNatRCActionTypePtrOutput
	ToAzureFirewallNatRCActionTypePtrOutputWithContext(context.Context) AzureFirewallNatRCActionTypePtrOutput
}

type azureFirewallNatRCActionTypePtr string

func AzureFirewallNatRCActionTypePtr(v string) AzureFirewallNatRCActionTypePtrInput {
	return (*azureFirewallNatRCActionTypePtr)(&v)
}

func (*azureFirewallNatRCActionTypePtr) ElementType() reflect.Type {
	return azureFirewallNatRCActionTypePtrType
}

func (in *azureFirewallNatRCActionTypePtr) ToAzureFirewallNatRCActionTypePtrOutput() AzureFirewallNatRCActionTypePtrOutput {
	return pulumi.ToOutput(in).(AzureFirewallNatRCActionTypePtrOutput)
}

func (in *azureFirewallNatRCActionTypePtr) ToAzureFirewallNatRCActionTypePtrOutputWithContext(ctx context.Context) AzureFirewallNatRCActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureFirewallNatRCActionTypePtrOutput)
}

type AzureFirewallNetworkRuleProtocol string

const (
	AzureFirewallNetworkRuleProtocolTCP  = AzureFirewallNetworkRuleProtocol("TCP")
	AzureFirewallNetworkRuleProtocolUDP  = AzureFirewallNetworkRuleProtocol("UDP")
	AzureFirewallNetworkRuleProtocolAny  = AzureFirewallNetworkRuleProtocol("Any")
	AzureFirewallNetworkRuleProtocolICMP = AzureFirewallNetworkRuleProtocol("ICMP")
)

func (AzureFirewallNetworkRuleProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallNetworkRuleProtocol)(nil)).Elem()
}

func (e AzureFirewallNetworkRuleProtocol) ToAzureFirewallNetworkRuleProtocolOutput() AzureFirewallNetworkRuleProtocolOutput {
	return pulumi.ToOutput(e).(AzureFirewallNetworkRuleProtocolOutput)
}

func (e AzureFirewallNetworkRuleProtocol) ToAzureFirewallNetworkRuleProtocolOutputWithContext(ctx context.Context) AzureFirewallNetworkRuleProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureFirewallNetworkRuleProtocolOutput)
}

func (e AzureFirewallNetworkRuleProtocol) ToAzureFirewallNetworkRuleProtocolPtrOutput() AzureFirewallNetworkRuleProtocolPtrOutput {
	return e.ToAzureFirewallNetworkRuleProtocolPtrOutputWithContext(context.Background())
}

func (e AzureFirewallNetworkRuleProtocol) ToAzureFirewallNetworkRuleProtocolPtrOutputWithContext(ctx context.Context) AzureFirewallNetworkRuleProtocolPtrOutput {
	return AzureFirewallNetworkRuleProtocol(e).ToAzureFirewallNetworkRuleProtocolOutputWithContext(ctx).ToAzureFirewallNetworkRuleProtocolPtrOutputWithContext(ctx)
}

func (e AzureFirewallNetworkRuleProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallNetworkRuleProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallNetworkRuleProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureFirewallNetworkRuleProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureFirewallNetworkRuleProtocolOutput struct{ *pulumi.OutputState }

func (AzureFirewallNetworkRuleProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallNetworkRuleProtocol)(nil)).Elem()
}

func (o AzureFirewallNetworkRuleProtocolOutput) ToAzureFirewallNetworkRuleProtocolOutput() AzureFirewallNetworkRuleProtocolOutput {
	return o
}

func (o AzureFirewallNetworkRuleProtocolOutput) ToAzureFirewallNetworkRuleProtocolOutputWithContext(ctx context.Context) AzureFirewallNetworkRuleProtocolOutput {
	return o
}

func (o AzureFirewallNetworkRuleProtocolOutput) ToAzureFirewallNetworkRuleProtocolPtrOutput() AzureFirewallNetworkRuleProtocolPtrOutput {
	return o.ToAzureFirewallNetworkRuleProtocolPtrOutputWithContext(context.Background())
}

func (o AzureFirewallNetworkRuleProtocolOutput) ToAzureFirewallNetworkRuleProtocolPtrOutputWithContext(ctx context.Context) AzureFirewallNetworkRuleProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFirewallNetworkRuleProtocol) *AzureFirewallNetworkRuleProtocol {
		return &v
	}).(AzureFirewallNetworkRuleProtocolPtrOutput)
}

func (o AzureFirewallNetworkRuleProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureFirewallNetworkRuleProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallNetworkRuleProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureFirewallNetworkRuleProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallNetworkRuleProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallNetworkRuleProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureFirewallNetworkRuleProtocolPtrOutput struct{ *pulumi.OutputState }

func (AzureFirewallNetworkRuleProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewallNetworkRuleProtocol)(nil)).Elem()
}

func (o AzureFirewallNetworkRuleProtocolPtrOutput) ToAzureFirewallNetworkRuleProtocolPtrOutput() AzureFirewallNetworkRuleProtocolPtrOutput {
	return o
}

func (o AzureFirewallNetworkRuleProtocolPtrOutput) ToAzureFirewallNetworkRuleProtocolPtrOutputWithContext(ctx context.Context) AzureFirewallNetworkRuleProtocolPtrOutput {
	return o
}

func (o AzureFirewallNetworkRuleProtocolPtrOutput) Elem() AzureFirewallNetworkRuleProtocolOutput {
	return o.ApplyT(func(v *AzureFirewallNetworkRuleProtocol) AzureFirewallNetworkRuleProtocol {
		if v != nil {
			return *v
		}
		var ret AzureFirewallNetworkRuleProtocol
		return ret
	}).(AzureFirewallNetworkRuleProtocolOutput)
}

func (o AzureFirewallNetworkRuleProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallNetworkRuleProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureFirewallNetworkRuleProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureFirewallNetworkRuleProtocolInput interface {
	pulumi.Input

	ToAzureFirewallNetworkRuleProtocolOutput() AzureFirewallNetworkRuleProtocolOutput
	ToAzureFirewallNetworkRuleProtocolOutputWithContext(context.Context) AzureFirewallNetworkRuleProtocolOutput
}

var azureFirewallNetworkRuleProtocolPtrType = reflect.TypeOf((**AzureFirewallNetworkRuleProtocol)(nil)).Elem()

type AzureFirewallNetworkRuleProtocolPtrInput interface {
	pulumi.Input

	ToAzureFirewallNetworkRuleProtocolPtrOutput() AzureFirewallNetworkRuleProtocolPtrOutput
	ToAzureFirewallNetworkRuleProtocolPtrOutputWithContext(context.Context) AzureFirewallNetworkRuleProtocolPtrOutput
}

type azureFirewallNetworkRuleProtocolPtr string

func AzureFirewallNetworkRuleProtocolPtr(v string) AzureFirewallNetworkRuleProtocolPtrInput {
	return (*azureFirewallNetworkRuleProtocolPtr)(&v)
}

func (*azureFirewallNetworkRuleProtocolPtr) ElementType() reflect.Type {
	return azureFirewallNetworkRuleProtocolPtrType
}

func (in *azureFirewallNetworkRuleProtocolPtr) ToAzureFirewallNetworkRuleProtocolPtrOutput() AzureFirewallNetworkRuleProtocolPtrOutput {
	return pulumi.ToOutput(in).(AzureFirewallNetworkRuleProtocolPtrOutput)
}

func (in *azureFirewallNetworkRuleProtocolPtr) ToAzureFirewallNetworkRuleProtocolPtrOutputWithContext(ctx context.Context) AzureFirewallNetworkRuleProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureFirewallNetworkRuleProtocolPtrOutput)
}

type AzureFirewallRCActionType string

const (
	AzureFirewallRCActionTypeAllow = AzureFirewallRCActionType("Allow")
	AzureFirewallRCActionTypeDeny  = AzureFirewallRCActionType("Deny")
)

func (AzureFirewallRCActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallRCActionType)(nil)).Elem()
}

func (e AzureFirewallRCActionType) ToAzureFirewallRCActionTypeOutput() AzureFirewallRCActionTypeOutput {
	return pulumi.ToOutput(e).(AzureFirewallRCActionTypeOutput)
}

func (e AzureFirewallRCActionType) ToAzureFirewallRCActionTypeOutputWithContext(ctx context.Context) AzureFirewallRCActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureFirewallRCActionTypeOutput)
}

func (e AzureFirewallRCActionType) ToAzureFirewallRCActionTypePtrOutput() AzureFirewallRCActionTypePtrOutput {
	return e.ToAzureFirewallRCActionTypePtrOutputWithContext(context.Background())
}

func (e AzureFirewallRCActionType) ToAzureFirewallRCActionTypePtrOutputWithContext(ctx context.Context) AzureFirewallRCActionTypePtrOutput {
	return AzureFirewallRCActionType(e).ToAzureFirewallRCActionTypeOutputWithContext(ctx).ToAzureFirewallRCActionTypePtrOutputWithContext(ctx)
}

func (e AzureFirewallRCActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallRCActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallRCActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureFirewallRCActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureFirewallRCActionTypeOutput struct{ *pulumi.OutputState }

func (AzureFirewallRCActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallRCActionType)(nil)).Elem()
}

func (o AzureFirewallRCActionTypeOutput) ToAzureFirewallRCActionTypeOutput() AzureFirewallRCActionTypeOutput {
	return o
}

func (o AzureFirewallRCActionTypeOutput) ToAzureFirewallRCActionTypeOutputWithContext(ctx context.Context) AzureFirewallRCActionTypeOutput {
	return o
}

func (o AzureFirewallRCActionTypeOutput) ToAzureFirewallRCActionTypePtrOutput() AzureFirewallRCActionTypePtrOutput {
	return o.ToAzureFirewallRCActionTypePtrOutputWithContext(context.Background())
}

func (o AzureFirewallRCActionTypeOutput) ToAzureFirewallRCActionTypePtrOutputWithContext(ctx context.Context) AzureFirewallRCActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFirewallRCActionType) *AzureFirewallRCActionType {
		return &v
	}).(AzureFirewallRCActionTypePtrOutput)
}

func (o AzureFirewallRCActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureFirewallRCActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallRCActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureFirewallRCActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallRCActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallRCActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureFirewallRCActionTypePtrOutput struct{ *pulumi.OutputState }

func (AzureFirewallRCActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewallRCActionType)(nil)).Elem()
}

func (o AzureFirewallRCActionTypePtrOutput) ToAzureFirewallRCActionTypePtrOutput() AzureFirewallRCActionTypePtrOutput {
	return o
}

func (o AzureFirewallRCActionTypePtrOutput) ToAzureFirewallRCActionTypePtrOutputWithContext(ctx context.Context) AzureFirewallRCActionTypePtrOutput {
	return o
}

func (o AzureFirewallRCActionTypePtrOutput) Elem() AzureFirewallRCActionTypeOutput {
	return o.ApplyT(func(v *AzureFirewallRCActionType) AzureFirewallRCActionType {
		if v != nil {
			return *v
		}
		var ret AzureFirewallRCActionType
		return ret
	}).(AzureFirewallRCActionTypeOutput)
}

func (o AzureFirewallRCActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallRCActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureFirewallRCActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureFirewallRCActionTypeInput interface {
	pulumi.Input

	ToAzureFirewallRCActionTypeOutput() AzureFirewallRCActionTypeOutput
	ToAzureFirewallRCActionTypeOutputWithContext(context.Context) AzureFirewallRCActionTypeOutput
}

var azureFirewallRCActionTypePtrType = reflect.TypeOf((**AzureFirewallRCActionType)(nil)).Elem()

type AzureFirewallRCActionTypePtrInput interface {
	pulumi.Input

	ToAzureFirewallRCActionTypePtrOutput() AzureFirewallRCActionTypePtrOutput
	ToAzureFirewallRCActionTypePtrOutputWithContext(context.Context) AzureFirewallRCActionTypePtrOutput
}

type azureFirewallRCActionTypePtr string

func AzureFirewallRCActionTypePtr(v string) AzureFirewallRCActionTypePtrInput {
	return (*azureFirewallRCActionTypePtr)(&v)
}

func (*azureFirewallRCActionTypePtr) ElementType() reflect.Type {
	return azureFirewallRCActionTypePtrType
}

func (in *azureFirewallRCActionTypePtr) ToAzureFirewallRCActionTypePtrOutput() AzureFirewallRCActionTypePtrOutput {
	return pulumi.ToOutput(in).(AzureFirewallRCActionTypePtrOutput)
}

func (in *azureFirewallRCActionTypePtr) ToAzureFirewallRCActionTypePtrOutputWithContext(ctx context.Context) AzureFirewallRCActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureFirewallRCActionTypePtrOutput)
}

type AzureFirewallSkuName string

const (
	AzureFirewallSkuName_AZFW_VNet = AzureFirewallSkuName("AZFW_VNet")
	AzureFirewallSkuName_AZFW_Hub  = AzureFirewallSkuName("AZFW_Hub")
)

func (AzureFirewallSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallSkuName)(nil)).Elem()
}

func (e AzureFirewallSkuName) ToAzureFirewallSkuNameOutput() AzureFirewallSkuNameOutput {
	return pulumi.ToOutput(e).(AzureFirewallSkuNameOutput)
}

func (e AzureFirewallSkuName) ToAzureFirewallSkuNameOutputWithContext(ctx context.Context) AzureFirewallSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureFirewallSkuNameOutput)
}

func (e AzureFirewallSkuName) ToAzureFirewallSkuNamePtrOutput() AzureFirewallSkuNamePtrOutput {
	return e.ToAzureFirewallSkuNamePtrOutputWithContext(context.Background())
}

func (e AzureFirewallSkuName) ToAzureFirewallSkuNamePtrOutputWithContext(ctx context.Context) AzureFirewallSkuNamePtrOutput {
	return AzureFirewallSkuName(e).ToAzureFirewallSkuNameOutputWithContext(ctx).ToAzureFirewallSkuNamePtrOutputWithContext(ctx)
}

func (e AzureFirewallSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureFirewallSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureFirewallSkuNameOutput struct{ *pulumi.OutputState }

func (AzureFirewallSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallSkuName)(nil)).Elem()
}

func (o AzureFirewallSkuNameOutput) ToAzureFirewallSkuNameOutput() AzureFirewallSkuNameOutput {
	return o
}

func (o AzureFirewallSkuNameOutput) ToAzureFirewallSkuNameOutputWithContext(ctx context.Context) AzureFirewallSkuNameOutput {
	return o
}

func (o AzureFirewallSkuNameOutput) ToAzureFirewallSkuNamePtrOutput() AzureFirewallSkuNamePtrOutput {
	return o.ToAzureFirewallSkuNamePtrOutputWithContext(context.Background())
}

func (o AzureFirewallSkuNameOutput) ToAzureFirewallSkuNamePtrOutputWithContext(ctx context.Context) AzureFirewallSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFirewallSkuName) *AzureFirewallSkuName {
		return &v
	}).(AzureFirewallSkuNamePtrOutput)
}

func (o AzureFirewallSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureFirewallSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureFirewallSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureFirewallSkuNamePtrOutput struct{ *pulumi.OutputState }

func (AzureFirewallSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewallSkuName)(nil)).Elem()
}

func (o AzureFirewallSkuNamePtrOutput) ToAzureFirewallSkuNamePtrOutput() AzureFirewallSkuNamePtrOutput {
	return o
}

func (o AzureFirewallSkuNamePtrOutput) ToAzureFirewallSkuNamePtrOutputWithContext(ctx context.Context) AzureFirewallSkuNamePtrOutput {
	return o
}

func (o AzureFirewallSkuNamePtrOutput) Elem() AzureFirewallSkuNameOutput {
	return o.ApplyT(func(v *AzureFirewallSkuName) AzureFirewallSkuName {
		if v != nil {
			return *v
		}
		var ret AzureFirewallSkuName
		return ret
	}).(AzureFirewallSkuNameOutput)
}

func (o AzureFirewallSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureFirewallSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureFirewallSkuNameInput interface {
	pulumi.Input

	ToAzureFirewallSkuNameOutput() AzureFirewallSkuNameOutput
	ToAzureFirewallSkuNameOutputWithContext(context.Context) AzureFirewallSkuNameOutput
}

var azureFirewallSkuNamePtrType = reflect.TypeOf((**AzureFirewallSkuName)(nil)).Elem()

type AzureFirewallSkuNamePtrInput interface {
	pulumi.Input

	ToAzureFirewallSkuNamePtrOutput() AzureFirewallSkuNamePtrOutput
	ToAzureFirewallSkuNamePtrOutputWithContext(context.Context) AzureFirewallSkuNamePtrOutput
}

type azureFirewallSkuNamePtr string

func AzureFirewallSkuNamePtr(v string) AzureFirewallSkuNamePtrInput {
	return (*azureFirewallSkuNamePtr)(&v)
}

func (*azureFirewallSkuNamePtr) ElementType() reflect.Type {
	return azureFirewallSkuNamePtrType
}

func (in *azureFirewallSkuNamePtr) ToAzureFirewallSkuNamePtrOutput() AzureFirewallSkuNamePtrOutput {
	return pulumi.ToOutput(in).(AzureFirewallSkuNamePtrOutput)
}

func (in *azureFirewallSkuNamePtr) ToAzureFirewallSkuNamePtrOutputWithContext(ctx context.Context) AzureFirewallSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureFirewallSkuNamePtrOutput)
}

type AzureFirewallSkuTier string

const (
	AzureFirewallSkuTierStandard = AzureFirewallSkuTier("Standard")
	AzureFirewallSkuTierPremium  = AzureFirewallSkuTier("Premium")
)

func (AzureFirewallSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallSkuTier)(nil)).Elem()
}

func (e AzureFirewallSkuTier) ToAzureFirewallSkuTierOutput() AzureFirewallSkuTierOutput {
	return pulumi.ToOutput(e).(AzureFirewallSkuTierOutput)
}

func (e AzureFirewallSkuTier) ToAzureFirewallSkuTierOutputWithContext(ctx context.Context) AzureFirewallSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureFirewallSkuTierOutput)
}

func (e AzureFirewallSkuTier) ToAzureFirewallSkuTierPtrOutput() AzureFirewallSkuTierPtrOutput {
	return e.ToAzureFirewallSkuTierPtrOutputWithContext(context.Background())
}

func (e AzureFirewallSkuTier) ToAzureFirewallSkuTierPtrOutputWithContext(ctx context.Context) AzureFirewallSkuTierPtrOutput {
	return AzureFirewallSkuTier(e).ToAzureFirewallSkuTierOutputWithContext(ctx).ToAzureFirewallSkuTierPtrOutputWithContext(ctx)
}

func (e AzureFirewallSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureFirewallSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureFirewallSkuTierOutput struct{ *pulumi.OutputState }

func (AzureFirewallSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallSkuTier)(nil)).Elem()
}

func (o AzureFirewallSkuTierOutput) ToAzureFirewallSkuTierOutput() AzureFirewallSkuTierOutput {
	return o
}

func (o AzureFirewallSkuTierOutput) ToAzureFirewallSkuTierOutputWithContext(ctx context.Context) AzureFirewallSkuTierOutput {
	return o
}

func (o AzureFirewallSkuTierOutput) ToAzureFirewallSkuTierPtrOutput() AzureFirewallSkuTierPtrOutput {
	return o.ToAzureFirewallSkuTierPtrOutputWithContext(context.Background())
}

func (o AzureFirewallSkuTierOutput) ToAzureFirewallSkuTierPtrOutputWithContext(ctx context.Context) AzureFirewallSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFirewallSkuTier) *AzureFirewallSkuTier {
		return &v
	}).(AzureFirewallSkuTierPtrOutput)
}

func (o AzureFirewallSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureFirewallSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureFirewallSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureFirewallSkuTierPtrOutput struct{ *pulumi.OutputState }

func (AzureFirewallSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewallSkuTier)(nil)).Elem()
}

func (o AzureFirewallSkuTierPtrOutput) ToAzureFirewallSkuTierPtrOutput() AzureFirewallSkuTierPtrOutput {
	return o
}

func (o AzureFirewallSkuTierPtrOutput) ToAzureFirewallSkuTierPtrOutputWithContext(ctx context.Context) AzureFirewallSkuTierPtrOutput {
	return o
}

func (o AzureFirewallSkuTierPtrOutput) Elem() AzureFirewallSkuTierOutput {
	return o.ApplyT(func(v *AzureFirewallSkuTier) AzureFirewallSkuTier {
		if v != nil {
			return *v
		}
		var ret AzureFirewallSkuTier
		return ret
	}).(AzureFirewallSkuTierOutput)
}

func (o AzureFirewallSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureFirewallSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureFirewallSkuTierInput interface {
	pulumi.Input

	ToAzureFirewallSkuTierOutput() AzureFirewallSkuTierOutput
	ToAzureFirewallSkuTierOutputWithContext(context.Context) AzureFirewallSkuTierOutput
}

var azureFirewallSkuTierPtrType = reflect.TypeOf((**AzureFirewallSkuTier)(nil)).Elem()

type AzureFirewallSkuTierPtrInput interface {
	pulumi.Input

	ToAzureFirewallSkuTierPtrOutput() AzureFirewallSkuTierPtrOutput
	ToAzureFirewallSkuTierPtrOutputWithContext(context.Context) AzureFirewallSkuTierPtrOutput
}

type azureFirewallSkuTierPtr string

func AzureFirewallSkuTierPtr(v string) AzureFirewallSkuTierPtrInput {
	return (*azureFirewallSkuTierPtr)(&v)
}

func (*azureFirewallSkuTierPtr) ElementType() reflect.Type {
	return azureFirewallSkuTierPtrType
}

func (in *azureFirewallSkuTierPtr) ToAzureFirewallSkuTierPtrOutput() AzureFirewallSkuTierPtrOutput {
	return pulumi.ToOutput(in).(AzureFirewallSkuTierPtrOutput)
}

func (in *azureFirewallSkuTierPtr) ToAzureFirewallSkuTierPtrOutputWithContext(ctx context.Context) AzureFirewallSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureFirewallSkuTierPtrOutput)
}

type AzureFirewallThreatIntelMode string

const (
	AzureFirewallThreatIntelModeAlert = AzureFirewallThreatIntelMode("Alert")
	AzureFirewallThreatIntelModeDeny  = AzureFirewallThreatIntelMode("Deny")
	AzureFirewallThreatIntelModeOff   = AzureFirewallThreatIntelMode("Off")
)

func (AzureFirewallThreatIntelMode) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallThreatIntelMode)(nil)).Elem()
}

func (e AzureFirewallThreatIntelMode) ToAzureFirewallThreatIntelModeOutput() AzureFirewallThreatIntelModeOutput {
	return pulumi.ToOutput(e).(AzureFirewallThreatIntelModeOutput)
}

func (e AzureFirewallThreatIntelMode) ToAzureFirewallThreatIntelModeOutputWithContext(ctx context.Context) AzureFirewallThreatIntelModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureFirewallThreatIntelModeOutput)
}

func (e AzureFirewallThreatIntelMode) ToAzureFirewallThreatIntelModePtrOutput() AzureFirewallThreatIntelModePtrOutput {
	return e.ToAzureFirewallThreatIntelModePtrOutputWithContext(context.Background())
}

func (e AzureFirewallThreatIntelMode) ToAzureFirewallThreatIntelModePtrOutputWithContext(ctx context.Context) AzureFirewallThreatIntelModePtrOutput {
	return AzureFirewallThreatIntelMode(e).ToAzureFirewallThreatIntelModeOutputWithContext(ctx).ToAzureFirewallThreatIntelModePtrOutputWithContext(ctx)
}

func (e AzureFirewallThreatIntelMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallThreatIntelMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureFirewallThreatIntelMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureFirewallThreatIntelMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureFirewallThreatIntelModeOutput struct{ *pulumi.OutputState }

func (AzureFirewallThreatIntelModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFirewallThreatIntelMode)(nil)).Elem()
}

func (o AzureFirewallThreatIntelModeOutput) ToAzureFirewallThreatIntelModeOutput() AzureFirewallThreatIntelModeOutput {
	return o
}

func (o AzureFirewallThreatIntelModeOutput) ToAzureFirewallThreatIntelModeOutputWithContext(ctx context.Context) AzureFirewallThreatIntelModeOutput {
	return o
}

func (o AzureFirewallThreatIntelModeOutput) ToAzureFirewallThreatIntelModePtrOutput() AzureFirewallThreatIntelModePtrOutput {
	return o.ToAzureFirewallThreatIntelModePtrOutputWithContext(context.Background())
}

func (o AzureFirewallThreatIntelModeOutput) ToAzureFirewallThreatIntelModePtrOutputWithContext(ctx context.Context) AzureFirewallThreatIntelModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureFirewallThreatIntelMode) *AzureFirewallThreatIntelMode {
		return &v
	}).(AzureFirewallThreatIntelModePtrOutput)
}

func (o AzureFirewallThreatIntelModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureFirewallThreatIntelModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallThreatIntelMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureFirewallThreatIntelModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallThreatIntelModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureFirewallThreatIntelMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureFirewallThreatIntelModePtrOutput struct{ *pulumi.OutputState }

func (AzureFirewallThreatIntelModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureFirewallThreatIntelMode)(nil)).Elem()
}

func (o AzureFirewallThreatIntelModePtrOutput) ToAzureFirewallThreatIntelModePtrOutput() AzureFirewallThreatIntelModePtrOutput {
	return o
}

func (o AzureFirewallThreatIntelModePtrOutput) ToAzureFirewallThreatIntelModePtrOutputWithContext(ctx context.Context) AzureFirewallThreatIntelModePtrOutput {
	return o
}

func (o AzureFirewallThreatIntelModePtrOutput) Elem() AzureFirewallThreatIntelModeOutput {
	return o.ApplyT(func(v *AzureFirewallThreatIntelMode) AzureFirewallThreatIntelMode {
		if v != nil {
			return *v
		}
		var ret AzureFirewallThreatIntelMode
		return ret
	}).(AzureFirewallThreatIntelModeOutput)
}

func (o AzureFirewallThreatIntelModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureFirewallThreatIntelModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureFirewallThreatIntelMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureFirewallThreatIntelModeInput interface {
	pulumi.Input

	ToAzureFirewallThreatIntelModeOutput() AzureFirewallThreatIntelModeOutput
	ToAzureFirewallThreatIntelModeOutputWithContext(context.Context) AzureFirewallThreatIntelModeOutput
}

var azureFirewallThreatIntelModePtrType = reflect.TypeOf((**AzureFirewallThreatIntelMode)(nil)).Elem()

type AzureFirewallThreatIntelModePtrInput interface {
	pulumi.Input

	ToAzureFirewallThreatIntelModePtrOutput() AzureFirewallThreatIntelModePtrOutput
	ToAzureFirewallThreatIntelModePtrOutputWithContext(context.Context) AzureFirewallThreatIntelModePtrOutput
}

type azureFirewallThreatIntelModePtr string

func AzureFirewallThreatIntelModePtr(v string) AzureFirewallThreatIntelModePtrInput {
	return (*azureFirewallThreatIntelModePtr)(&v)
}

func (*azureFirewallThreatIntelModePtr) ElementType() reflect.Type {
	return azureFirewallThreatIntelModePtrType
}

func (in *azureFirewallThreatIntelModePtr) ToAzureFirewallThreatIntelModePtrOutput() AzureFirewallThreatIntelModePtrOutput {
	return pulumi.ToOutput(in).(AzureFirewallThreatIntelModePtrOutput)
}

func (in *azureFirewallThreatIntelModePtr) ToAzureFirewallThreatIntelModePtrOutputWithContext(ctx context.Context) AzureFirewallThreatIntelModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureFirewallThreatIntelModePtrOutput)
}

type BastionHostSkuName string

const (
	BastionHostSkuNameBasic    = BastionHostSkuName("Basic")
	BastionHostSkuNameStandard = BastionHostSkuName("Standard")
)

func (BastionHostSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*BastionHostSkuName)(nil)).Elem()
}

func (e BastionHostSkuName) ToBastionHostSkuNameOutput() BastionHostSkuNameOutput {
	return pulumi.ToOutput(e).(BastionHostSkuNameOutput)
}

func (e BastionHostSkuName) ToBastionHostSkuNameOutputWithContext(ctx context.Context) BastionHostSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BastionHostSkuNameOutput)
}

func (e BastionHostSkuName) ToBastionHostSkuNamePtrOutput() BastionHostSkuNamePtrOutput {
	return e.ToBastionHostSkuNamePtrOutputWithContext(context.Background())
}

func (e BastionHostSkuName) ToBastionHostSkuNamePtrOutputWithContext(ctx context.Context) BastionHostSkuNamePtrOutput {
	return BastionHostSkuName(e).ToBastionHostSkuNameOutputWithContext(ctx).ToBastionHostSkuNamePtrOutputWithContext(ctx)
}

func (e BastionHostSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BastionHostSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BastionHostSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BastionHostSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BastionHostSkuNameOutput struct{ *pulumi.OutputState }

func (BastionHostSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BastionHostSkuName)(nil)).Elem()
}

func (o BastionHostSkuNameOutput) ToBastionHostSkuNameOutput() BastionHostSkuNameOutput {
	return o
}

func (o BastionHostSkuNameOutput) ToBastionHostSkuNameOutputWithContext(ctx context.Context) BastionHostSkuNameOutput {
	return o
}

func (o BastionHostSkuNameOutput) ToBastionHostSkuNamePtrOutput() BastionHostSkuNamePtrOutput {
	return o.ToBastionHostSkuNamePtrOutputWithContext(context.Background())
}

func (o BastionHostSkuNameOutput) ToBastionHostSkuNamePtrOutputWithContext(ctx context.Context) BastionHostSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BastionHostSkuName) *BastionHostSkuName {
		return &v
	}).(BastionHostSkuNamePtrOutput)
}

func (o BastionHostSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BastionHostSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BastionHostSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BastionHostSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BastionHostSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BastionHostSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BastionHostSkuNamePtrOutput struct{ *pulumi.OutputState }

func (BastionHostSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHostSkuName)(nil)).Elem()
}

func (o BastionHostSkuNamePtrOutput) ToBastionHostSkuNamePtrOutput() BastionHostSkuNamePtrOutput {
	return o
}

func (o BastionHostSkuNamePtrOutput) ToBastionHostSkuNamePtrOutputWithContext(ctx context.Context) BastionHostSkuNamePtrOutput {
	return o
}

func (o BastionHostSkuNamePtrOutput) Elem() BastionHostSkuNameOutput {
	return o.ApplyT(func(v *BastionHostSkuName) BastionHostSkuName {
		if v != nil {
			return *v
		}
		var ret BastionHostSkuName
		return ret
	}).(BastionHostSkuNameOutput)
}

func (o BastionHostSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BastionHostSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BastionHostSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BastionHostSkuNameInput interface {
	pulumi.Input

	ToBastionHostSkuNameOutput() BastionHostSkuNameOutput
	ToBastionHostSkuNameOutputWithContext(context.Context) BastionHostSkuNameOutput
}

var bastionHostSkuNamePtrType = reflect.TypeOf((**BastionHostSkuName)(nil)).Elem()

type BastionHostSkuNamePtrInput interface {
	pulumi.Input

	ToBastionHostSkuNamePtrOutput() BastionHostSkuNamePtrOutput
	ToBastionHostSkuNamePtrOutputWithContext(context.Context) BastionHostSkuNamePtrOutput
}

type bastionHostSkuNamePtr string

func BastionHostSkuNamePtr(v string) BastionHostSkuNamePtrInput {
	return (*bastionHostSkuNamePtr)(&v)
}

func (*bastionHostSkuNamePtr) ElementType() reflect.Type {
	return bastionHostSkuNamePtrType
}

func (in *bastionHostSkuNamePtr) ToBastionHostSkuNamePtrOutput() BastionHostSkuNamePtrOutput {
	return pulumi.ToOutput(in).(BastionHostSkuNamePtrOutput)
}

func (in *bastionHostSkuNamePtr) ToBastionHostSkuNamePtrOutputWithContext(ctx context.Context) BastionHostSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BastionHostSkuNamePtrOutput)
}

type CommissionedState string

const (
	CommissionedStateProvisioning    = CommissionedState("Provisioning")
	CommissionedStateProvisioned     = CommissionedState("Provisioned")
	CommissionedStateCommissioning   = CommissionedState("Commissioning")
	CommissionedStateCommissioned    = CommissionedState("Commissioned")
	CommissionedStateDecommissioning = CommissionedState("Decommissioning")
	CommissionedStateDeprovisioning  = CommissionedState("Deprovisioning")
)

func (CommissionedState) ElementType() reflect.Type {
	return reflect.TypeOf((*CommissionedState)(nil)).Elem()
}

func (e CommissionedState) ToCommissionedStateOutput() CommissionedStateOutput {
	return pulumi.ToOutput(e).(CommissionedStateOutput)
}

func (e CommissionedState) ToCommissionedStateOutputWithContext(ctx context.Context) CommissionedStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CommissionedStateOutput)
}

func (e CommissionedState) ToCommissionedStatePtrOutput() CommissionedStatePtrOutput {
	return e.ToCommissionedStatePtrOutputWithContext(context.Background())
}

func (e CommissionedState) ToCommissionedStatePtrOutputWithContext(ctx context.Context) CommissionedStatePtrOutput {
	return CommissionedState(e).ToCommissionedStateOutputWithContext(ctx).ToCommissionedStatePtrOutputWithContext(ctx)
}

func (e CommissionedState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CommissionedState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CommissionedState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CommissionedState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CommissionedStateOutput struct{ *pulumi.OutputState }

func (CommissionedStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommissionedState)(nil)).Elem()
}

func (o CommissionedStateOutput) ToCommissionedStateOutput() CommissionedStateOutput {
	return o
}

func (o CommissionedStateOutput) ToCommissionedStateOutputWithContext(ctx context.Context) CommissionedStateOutput {
	return o
}

func (o CommissionedStateOutput) ToCommissionedStatePtrOutput() CommissionedStatePtrOutput {
	return o.ToCommissionedStatePtrOutputWithContext(context.Background())
}

func (o CommissionedStateOutput) ToCommissionedStatePtrOutputWithContext(ctx context.Context) CommissionedStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommissionedState) *CommissionedState {
		return &v
	}).(CommissionedStatePtrOutput)
}

func (o CommissionedStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CommissionedStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CommissionedState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CommissionedStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CommissionedStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CommissionedState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CommissionedStatePtrOutput struct{ *pulumi.OutputState }

func (CommissionedStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommissionedState)(nil)).Elem()
}

func (o CommissionedStatePtrOutput) ToCommissionedStatePtrOutput() CommissionedStatePtrOutput {
	return o
}

func (o CommissionedStatePtrOutput) ToCommissionedStatePtrOutputWithContext(ctx context.Context) CommissionedStatePtrOutput {
	return o
}

func (o CommissionedStatePtrOutput) Elem() CommissionedStateOutput {
	return o.ApplyT(func(v *CommissionedState) CommissionedState {
		if v != nil {
			return *v
		}
		var ret CommissionedState
		return ret
	}).(CommissionedStateOutput)
}

func (o CommissionedStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CommissionedStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CommissionedState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CommissionedStateInput interface {
	pulumi.Input

	ToCommissionedStateOutput() CommissionedStateOutput
	ToCommissionedStateOutputWithContext(context.Context) CommissionedStateOutput
}

var commissionedStatePtrType = reflect.TypeOf((**CommissionedState)(nil)).Elem()

type CommissionedStatePtrInput interface {
	pulumi.Input

	ToCommissionedStatePtrOutput() CommissionedStatePtrOutput
	ToCommissionedStatePtrOutputWithContext(context.Context) CommissionedStatePtrOutput
}

type commissionedStatePtr string

func CommissionedStatePtr(v string) CommissionedStatePtrInput {
	return (*commissionedStatePtr)(&v)
}

func (*commissionedStatePtr) ElementType() reflect.Type {
	return commissionedStatePtrType
}

func (in *commissionedStatePtr) ToCommissionedStatePtrOutput() CommissionedStatePtrOutput {
	return pulumi.ToOutput(in).(CommissionedStatePtrOutput)
}

func (in *commissionedStatePtr) ToCommissionedStatePtrOutputWithContext(ctx context.Context) CommissionedStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CommissionedStatePtrOutput)
}

type ConnectionMonitorEndpointFilterItemType string

const (
	ConnectionMonitorEndpointFilterItemTypeAgentAddress = ConnectionMonitorEndpointFilterItemType("AgentAddress")
)

func (ConnectionMonitorEndpointFilterItemType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorEndpointFilterItemType)(nil)).Elem()
}

func (e ConnectionMonitorEndpointFilterItemType) ToConnectionMonitorEndpointFilterItemTypeOutput() ConnectionMonitorEndpointFilterItemTypeOutput {
	return pulumi.ToOutput(e).(ConnectionMonitorEndpointFilterItemTypeOutput)
}

func (e ConnectionMonitorEndpointFilterItemType) ToConnectionMonitorEndpointFilterItemTypeOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterItemTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionMonitorEndpointFilterItemTypeOutput)
}

func (e ConnectionMonitorEndpointFilterItemType) ToConnectionMonitorEndpointFilterItemTypePtrOutput() ConnectionMonitorEndpointFilterItemTypePtrOutput {
	return e.ToConnectionMonitorEndpointFilterItemTypePtrOutputWithContext(context.Background())
}

func (e ConnectionMonitorEndpointFilterItemType) ToConnectionMonitorEndpointFilterItemTypePtrOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterItemTypePtrOutput {
	return ConnectionMonitorEndpointFilterItemType(e).ToConnectionMonitorEndpointFilterItemTypeOutputWithContext(ctx).ToConnectionMonitorEndpointFilterItemTypePtrOutputWithContext(ctx)
}

func (e ConnectionMonitorEndpointFilterItemType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionMonitorEndpointFilterItemType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionMonitorEndpointFilterItemType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionMonitorEndpointFilterItemType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionMonitorEndpointFilterItemTypeOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorEndpointFilterItemTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorEndpointFilterItemType)(nil)).Elem()
}

func (o ConnectionMonitorEndpointFilterItemTypeOutput) ToConnectionMonitorEndpointFilterItemTypeOutput() ConnectionMonitorEndpointFilterItemTypeOutput {
	return o
}

func (o ConnectionMonitorEndpointFilterItemTypeOutput) ToConnectionMonitorEndpointFilterItemTypeOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterItemTypeOutput {
	return o
}

func (o ConnectionMonitorEndpointFilterItemTypeOutput) ToConnectionMonitorEndpointFilterItemTypePtrOutput() ConnectionMonitorEndpointFilterItemTypePtrOutput {
	return o.ToConnectionMonitorEndpointFilterItemTypePtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorEndpointFilterItemTypeOutput) ToConnectionMonitorEndpointFilterItemTypePtrOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterItemTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionMonitorEndpointFilterItemType) *ConnectionMonitorEndpointFilterItemType {
		return &v
	}).(ConnectionMonitorEndpointFilterItemTypePtrOutput)
}

func (o ConnectionMonitorEndpointFilterItemTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionMonitorEndpointFilterItemTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionMonitorEndpointFilterItemType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionMonitorEndpointFilterItemTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorEndpointFilterItemTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionMonitorEndpointFilterItemType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionMonitorEndpointFilterItemTypePtrOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorEndpointFilterItemTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorEndpointFilterItemType)(nil)).Elem()
}

func (o ConnectionMonitorEndpointFilterItemTypePtrOutput) ToConnectionMonitorEndpointFilterItemTypePtrOutput() ConnectionMonitorEndpointFilterItemTypePtrOutput {
	return o
}

func (o ConnectionMonitorEndpointFilterItemTypePtrOutput) ToConnectionMonitorEndpointFilterItemTypePtrOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterItemTypePtrOutput {
	return o
}

func (o ConnectionMonitorEndpointFilterItemTypePtrOutput) Elem() ConnectionMonitorEndpointFilterItemTypeOutput {
	return o.ApplyT(func(v *ConnectionMonitorEndpointFilterItemType) ConnectionMonitorEndpointFilterItemType {
		if v != nil {
			return *v
		}
		var ret ConnectionMonitorEndpointFilterItemType
		return ret
	}).(ConnectionMonitorEndpointFilterItemTypeOutput)
}

func (o ConnectionMonitorEndpointFilterItemTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorEndpointFilterItemTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionMonitorEndpointFilterItemType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionMonitorEndpointFilterItemTypeInput interface {
	pulumi.Input

	ToConnectionMonitorEndpointFilterItemTypeOutput() ConnectionMonitorEndpointFilterItemTypeOutput
	ToConnectionMonitorEndpointFilterItemTypeOutputWithContext(context.Context) ConnectionMonitorEndpointFilterItemTypeOutput
}

var connectionMonitorEndpointFilterItemTypePtrType = reflect.TypeOf((**ConnectionMonitorEndpointFilterItemType)(nil)).Elem()

type ConnectionMonitorEndpointFilterItemTypePtrInput interface {
	pulumi.Input

	ToConnectionMonitorEndpointFilterItemTypePtrOutput() ConnectionMonitorEndpointFilterItemTypePtrOutput
	ToConnectionMonitorEndpointFilterItemTypePtrOutputWithContext(context.Context) ConnectionMonitorEndpointFilterItemTypePtrOutput
}

type connectionMonitorEndpointFilterItemTypePtr string

func ConnectionMonitorEndpointFilterItemTypePtr(v string) ConnectionMonitorEndpointFilterItemTypePtrInput {
	return (*connectionMonitorEndpointFilterItemTypePtr)(&v)
}

func (*connectionMonitorEndpointFilterItemTypePtr) ElementType() reflect.Type {
	return connectionMonitorEndpointFilterItemTypePtrType
}

func (in *connectionMonitorEndpointFilterItemTypePtr) ToConnectionMonitorEndpointFilterItemTypePtrOutput() ConnectionMonitorEndpointFilterItemTypePtrOutput {
	return pulumi.ToOutput(in).(ConnectionMonitorEndpointFilterItemTypePtrOutput)
}

func (in *connectionMonitorEndpointFilterItemTypePtr) ToConnectionMonitorEndpointFilterItemTypePtrOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterItemTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionMonitorEndpointFilterItemTypePtrOutput)
}

type ConnectionMonitorEndpointFilterType string

const (
	ConnectionMonitorEndpointFilterTypeInclude = ConnectionMonitorEndpointFilterType("Include")
)

func (ConnectionMonitorEndpointFilterType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorEndpointFilterType)(nil)).Elem()
}

func (e ConnectionMonitorEndpointFilterType) ToConnectionMonitorEndpointFilterTypeOutput() ConnectionMonitorEndpointFilterTypeOutput {
	return pulumi.ToOutput(e).(ConnectionMonitorEndpointFilterTypeOutput)
}

func (e ConnectionMonitorEndpointFilterType) ToConnectionMonitorEndpointFilterTypeOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionMonitorEndpointFilterTypeOutput)
}

func (e ConnectionMonitorEndpointFilterType) ToConnectionMonitorEndpointFilterTypePtrOutput() ConnectionMonitorEndpointFilterTypePtrOutput {
	return e.ToConnectionMonitorEndpointFilterTypePtrOutputWithContext(context.Background())
}

func (e ConnectionMonitorEndpointFilterType) ToConnectionMonitorEndpointFilterTypePtrOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterTypePtrOutput {
	return ConnectionMonitorEndpointFilterType(e).ToConnectionMonitorEndpointFilterTypeOutputWithContext(ctx).ToConnectionMonitorEndpointFilterTypePtrOutputWithContext(ctx)
}

func (e ConnectionMonitorEndpointFilterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionMonitorEndpointFilterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionMonitorEndpointFilterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionMonitorEndpointFilterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionMonitorEndpointFilterTypeOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorEndpointFilterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorEndpointFilterType)(nil)).Elem()
}

func (o ConnectionMonitorEndpointFilterTypeOutput) ToConnectionMonitorEndpointFilterTypeOutput() ConnectionMonitorEndpointFilterTypeOutput {
	return o
}

func (o ConnectionMonitorEndpointFilterTypeOutput) ToConnectionMonitorEndpointFilterTypeOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterTypeOutput {
	return o
}

func (o ConnectionMonitorEndpointFilterTypeOutput) ToConnectionMonitorEndpointFilterTypePtrOutput() ConnectionMonitorEndpointFilterTypePtrOutput {
	return o.ToConnectionMonitorEndpointFilterTypePtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorEndpointFilterTypeOutput) ToConnectionMonitorEndpointFilterTypePtrOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionMonitorEndpointFilterType) *ConnectionMonitorEndpointFilterType {
		return &v
	}).(ConnectionMonitorEndpointFilterTypePtrOutput)
}

func (o ConnectionMonitorEndpointFilterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionMonitorEndpointFilterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionMonitorEndpointFilterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionMonitorEndpointFilterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorEndpointFilterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionMonitorEndpointFilterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionMonitorEndpointFilterTypePtrOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorEndpointFilterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorEndpointFilterType)(nil)).Elem()
}

func (o ConnectionMonitorEndpointFilterTypePtrOutput) ToConnectionMonitorEndpointFilterTypePtrOutput() ConnectionMonitorEndpointFilterTypePtrOutput {
	return o
}

func (o ConnectionMonitorEndpointFilterTypePtrOutput) ToConnectionMonitorEndpointFilterTypePtrOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterTypePtrOutput {
	return o
}

func (o ConnectionMonitorEndpointFilterTypePtrOutput) Elem() ConnectionMonitorEndpointFilterTypeOutput {
	return o.ApplyT(func(v *ConnectionMonitorEndpointFilterType) ConnectionMonitorEndpointFilterType {
		if v != nil {
			return *v
		}
		var ret ConnectionMonitorEndpointFilterType
		return ret
	}).(ConnectionMonitorEndpointFilterTypeOutput)
}

func (o ConnectionMonitorEndpointFilterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorEndpointFilterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionMonitorEndpointFilterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionMonitorEndpointFilterTypeInput interface {
	pulumi.Input

	ToConnectionMonitorEndpointFilterTypeOutput() ConnectionMonitorEndpointFilterTypeOutput
	ToConnectionMonitorEndpointFilterTypeOutputWithContext(context.Context) ConnectionMonitorEndpointFilterTypeOutput
}

var connectionMonitorEndpointFilterTypePtrType = reflect.TypeOf((**ConnectionMonitorEndpointFilterType)(nil)).Elem()

type ConnectionMonitorEndpointFilterTypePtrInput interface {
	pulumi.Input

	ToConnectionMonitorEndpointFilterTypePtrOutput() ConnectionMonitorEndpointFilterTypePtrOutput
	ToConnectionMonitorEndpointFilterTypePtrOutputWithContext(context.Context) ConnectionMonitorEndpointFilterTypePtrOutput
}

type connectionMonitorEndpointFilterTypePtr string

func ConnectionMonitorEndpointFilterTypePtr(v string) ConnectionMonitorEndpointFilterTypePtrInput {
	return (*connectionMonitorEndpointFilterTypePtr)(&v)
}

func (*connectionMonitorEndpointFilterTypePtr) ElementType() reflect.Type {
	return connectionMonitorEndpointFilterTypePtrType
}

func (in *connectionMonitorEndpointFilterTypePtr) ToConnectionMonitorEndpointFilterTypePtrOutput() ConnectionMonitorEndpointFilterTypePtrOutput {
	return pulumi.ToOutput(in).(ConnectionMonitorEndpointFilterTypePtrOutput)
}

func (in *connectionMonitorEndpointFilterTypePtr) ToConnectionMonitorEndpointFilterTypePtrOutputWithContext(ctx context.Context) ConnectionMonitorEndpointFilterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionMonitorEndpointFilterTypePtrOutput)
}

type ConnectionMonitorTestConfigurationProtocol string

const (
	ConnectionMonitorTestConfigurationProtocolTcp  = ConnectionMonitorTestConfigurationProtocol("Tcp")
	ConnectionMonitorTestConfigurationProtocolHttp = ConnectionMonitorTestConfigurationProtocol("Http")
	ConnectionMonitorTestConfigurationProtocolIcmp = ConnectionMonitorTestConfigurationProtocol("Icmp")
)

func (ConnectionMonitorTestConfigurationProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorTestConfigurationProtocol)(nil)).Elem()
}

func (e ConnectionMonitorTestConfigurationProtocol) ToConnectionMonitorTestConfigurationProtocolOutput() ConnectionMonitorTestConfigurationProtocolOutput {
	return pulumi.ToOutput(e).(ConnectionMonitorTestConfigurationProtocolOutput)
}

func (e ConnectionMonitorTestConfigurationProtocol) ToConnectionMonitorTestConfigurationProtocolOutputWithContext(ctx context.Context) ConnectionMonitorTestConfigurationProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionMonitorTestConfigurationProtocolOutput)
}

func (e ConnectionMonitorTestConfigurationProtocol) ToConnectionMonitorTestConfigurationProtocolPtrOutput() ConnectionMonitorTestConfigurationProtocolPtrOutput {
	return e.ToConnectionMonitorTestConfigurationProtocolPtrOutputWithContext(context.Background())
}

func (e ConnectionMonitorTestConfigurationProtocol) ToConnectionMonitorTestConfigurationProtocolPtrOutputWithContext(ctx context.Context) ConnectionMonitorTestConfigurationProtocolPtrOutput {
	return ConnectionMonitorTestConfigurationProtocol(e).ToConnectionMonitorTestConfigurationProtocolOutputWithContext(ctx).ToConnectionMonitorTestConfigurationProtocolPtrOutputWithContext(ctx)
}

func (e ConnectionMonitorTestConfigurationProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionMonitorTestConfigurationProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionMonitorTestConfigurationProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionMonitorTestConfigurationProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionMonitorTestConfigurationProtocolOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorTestConfigurationProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionMonitorTestConfigurationProtocol)(nil)).Elem()
}

func (o ConnectionMonitorTestConfigurationProtocolOutput) ToConnectionMonitorTestConfigurationProtocolOutput() ConnectionMonitorTestConfigurationProtocolOutput {
	return o
}

func (o ConnectionMonitorTestConfigurationProtocolOutput) ToConnectionMonitorTestConfigurationProtocolOutputWithContext(ctx context.Context) ConnectionMonitorTestConfigurationProtocolOutput {
	return o
}

func (o ConnectionMonitorTestConfigurationProtocolOutput) ToConnectionMonitorTestConfigurationProtocolPtrOutput() ConnectionMonitorTestConfigurationProtocolPtrOutput {
	return o.ToConnectionMonitorTestConfigurationProtocolPtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorTestConfigurationProtocolOutput) ToConnectionMonitorTestConfigurationProtocolPtrOutputWithContext(ctx context.Context) ConnectionMonitorTestConfigurationProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionMonitorTestConfigurationProtocol) *ConnectionMonitorTestConfigurationProtocol {
		return &v
	}).(ConnectionMonitorTestConfigurationProtocolPtrOutput)
}

func (o ConnectionMonitorTestConfigurationProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionMonitorTestConfigurationProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionMonitorTestConfigurationProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionMonitorTestConfigurationProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorTestConfigurationProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionMonitorTestConfigurationProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionMonitorTestConfigurationProtocolPtrOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorTestConfigurationProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitorTestConfigurationProtocol)(nil)).Elem()
}

func (o ConnectionMonitorTestConfigurationProtocolPtrOutput) ToConnectionMonitorTestConfigurationProtocolPtrOutput() ConnectionMonitorTestConfigurationProtocolPtrOutput {
	return o
}

func (o ConnectionMonitorTestConfigurationProtocolPtrOutput) ToConnectionMonitorTestConfigurationProtocolPtrOutputWithContext(ctx context.Context) ConnectionMonitorTestConfigurationProtocolPtrOutput {
	return o
}

func (o ConnectionMonitorTestConfigurationProtocolPtrOutput) Elem() ConnectionMonitorTestConfigurationProtocolOutput {
	return o.ApplyT(func(v *ConnectionMonitorTestConfigurationProtocol) ConnectionMonitorTestConfigurationProtocol {
		if v != nil {
			return *v
		}
		var ret ConnectionMonitorTestConfigurationProtocol
		return ret
	}).(ConnectionMonitorTestConfigurationProtocolOutput)
}

func (o ConnectionMonitorTestConfigurationProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionMonitorTestConfigurationProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionMonitorTestConfigurationProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionMonitorTestConfigurationProtocolInput interface {
	pulumi.Input

	ToConnectionMonitorTestConfigurationProtocolOutput() ConnectionMonitorTestConfigurationProtocolOutput
	ToConnectionMonitorTestConfigurationProtocolOutputWithContext(context.Context) ConnectionMonitorTestConfigurationProtocolOutput
}

var connectionMonitorTestConfigurationProtocolPtrType = reflect.TypeOf((**ConnectionMonitorTestConfigurationProtocol)(nil)).Elem()

type ConnectionMonitorTestConfigurationProtocolPtrInput interface {
	pulumi.Input

	ToConnectionMonitorTestConfigurationProtocolPtrOutput() ConnectionMonitorTestConfigurationProtocolPtrOutput
	ToConnectionMonitorTestConfigurationProtocolPtrOutputWithContext(context.Context) ConnectionMonitorTestConfigurationProtocolPtrOutput
}

type connectionMonitorTestConfigurationProtocolPtr string

func ConnectionMonitorTestConfigurationProtocolPtr(v string) ConnectionMonitorTestConfigurationProtocolPtrInput {
	return (*connectionMonitorTestConfigurationProtocolPtr)(&v)
}

func (*connectionMonitorTestConfigurationProtocolPtr) ElementType() reflect.Type {
	return connectionMonitorTestConfigurationProtocolPtrType
}

func (in *connectionMonitorTestConfigurationProtocolPtr) ToConnectionMonitorTestConfigurationProtocolPtrOutput() ConnectionMonitorTestConfigurationProtocolPtrOutput {
	return pulumi.ToOutput(in).(ConnectionMonitorTestConfigurationProtocolPtrOutput)
}

func (in *connectionMonitorTestConfigurationProtocolPtr) ToConnectionMonitorTestConfigurationProtocolPtrOutputWithContext(ctx context.Context) ConnectionMonitorTestConfigurationProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionMonitorTestConfigurationProtocolPtrOutput)
}

type CoverageLevel string

const (
	CoverageLevelDefault      = CoverageLevel("Default")
	CoverageLevelLow          = CoverageLevel("Low")
	CoverageLevelBelowAverage = CoverageLevel("BelowAverage")
	CoverageLevelAverage      = CoverageLevel("Average")
	CoverageLevelAboveAverage = CoverageLevel("AboveAverage")
	CoverageLevelFull         = CoverageLevel("Full")
)

func (CoverageLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*CoverageLevel)(nil)).Elem()
}

func (e CoverageLevel) ToCoverageLevelOutput() CoverageLevelOutput {
	return pulumi.ToOutput(e).(CoverageLevelOutput)
}

func (e CoverageLevel) ToCoverageLevelOutputWithContext(ctx context.Context) CoverageLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CoverageLevelOutput)
}

func (e CoverageLevel) ToCoverageLevelPtrOutput() CoverageLevelPtrOutput {
	return e.ToCoverageLevelPtrOutputWithContext(context.Background())
}

func (e CoverageLevel) ToCoverageLevelPtrOutputWithContext(ctx context.Context) CoverageLevelPtrOutput {
	return CoverageLevel(e).ToCoverageLevelOutputWithContext(ctx).ToCoverageLevelPtrOutputWithContext(ctx)
}

func (e CoverageLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CoverageLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CoverageLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CoverageLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CoverageLevelOutput struct{ *pulumi.OutputState }

func (CoverageLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CoverageLevel)(nil)).Elem()
}

func (o CoverageLevelOutput) ToCoverageLevelOutput() CoverageLevelOutput {
	return o
}

func (o CoverageLevelOutput) ToCoverageLevelOutputWithContext(ctx context.Context) CoverageLevelOutput {
	return o
}

func (o CoverageLevelOutput) ToCoverageLevelPtrOutput() CoverageLevelPtrOutput {
	return o.ToCoverageLevelPtrOutputWithContext(context.Background())
}

func (o CoverageLevelOutput) ToCoverageLevelPtrOutputWithContext(ctx context.Context) CoverageLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CoverageLevel) *CoverageLevel {
		return &v
	}).(CoverageLevelPtrOutput)
}

func (o CoverageLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CoverageLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CoverageLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CoverageLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CoverageLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CoverageLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CoverageLevelPtrOutput struct{ *pulumi.OutputState }

func (CoverageLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CoverageLevel)(nil)).Elem()
}

func (o CoverageLevelPtrOutput) ToCoverageLevelPtrOutput() CoverageLevelPtrOutput {
	return o
}

func (o CoverageLevelPtrOutput) ToCoverageLevelPtrOutputWithContext(ctx context.Context) CoverageLevelPtrOutput {
	return o
}

func (o CoverageLevelPtrOutput) Elem() CoverageLevelOutput {
	return o.ApplyT(func(v *CoverageLevel) CoverageLevel {
		if v != nil {
			return *v
		}
		var ret CoverageLevel
		return ret
	}).(CoverageLevelOutput)
}

func (o CoverageLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CoverageLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CoverageLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CoverageLevelInput interface {
	pulumi.Input

	ToCoverageLevelOutput() CoverageLevelOutput
	ToCoverageLevelOutputWithContext(context.Context) CoverageLevelOutput
}

var coverageLevelPtrType = reflect.TypeOf((**CoverageLevel)(nil)).Elem()

type CoverageLevelPtrInput interface {
	pulumi.Input

	ToCoverageLevelPtrOutput() CoverageLevelPtrOutput
	ToCoverageLevelPtrOutputWithContext(context.Context) CoverageLevelPtrOutput
}

type coverageLevelPtr string

func CoverageLevelPtr(v string) CoverageLevelPtrInput {
	return (*coverageLevelPtr)(&v)
}

func (*coverageLevelPtr) ElementType() reflect.Type {
	return coverageLevelPtrType
}

func (in *coverageLevelPtr) ToCoverageLevelPtrOutput() CoverageLevelPtrOutput {
	return pulumi.ToOutput(in).(CoverageLevelPtrOutput)
}

func (in *coverageLevelPtr) ToCoverageLevelPtrOutputWithContext(ctx context.Context) CoverageLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CoverageLevelPtrOutput)
}

type DdosCustomPolicyProtocol string

const (
	DdosCustomPolicyProtocolTcp = DdosCustomPolicyProtocol("Tcp")
	DdosCustomPolicyProtocolUdp = DdosCustomPolicyProtocol("Udp")
	DdosCustomPolicyProtocolSyn = DdosCustomPolicyProtocol("Syn")
)

func (DdosCustomPolicyProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosCustomPolicyProtocol)(nil)).Elem()
}

func (e DdosCustomPolicyProtocol) ToDdosCustomPolicyProtocolOutput() DdosCustomPolicyProtocolOutput {
	return pulumi.ToOutput(e).(DdosCustomPolicyProtocolOutput)
}

func (e DdosCustomPolicyProtocol) ToDdosCustomPolicyProtocolOutputWithContext(ctx context.Context) DdosCustomPolicyProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DdosCustomPolicyProtocolOutput)
}

func (e DdosCustomPolicyProtocol) ToDdosCustomPolicyProtocolPtrOutput() DdosCustomPolicyProtocolPtrOutput {
	return e.ToDdosCustomPolicyProtocolPtrOutputWithContext(context.Background())
}

func (e DdosCustomPolicyProtocol) ToDdosCustomPolicyProtocolPtrOutputWithContext(ctx context.Context) DdosCustomPolicyProtocolPtrOutput {
	return DdosCustomPolicyProtocol(e).ToDdosCustomPolicyProtocolOutputWithContext(ctx).ToDdosCustomPolicyProtocolPtrOutputWithContext(ctx)
}

func (e DdosCustomPolicyProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DdosCustomPolicyProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DdosCustomPolicyProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DdosCustomPolicyProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DdosCustomPolicyProtocolOutput struct{ *pulumi.OutputState }

func (DdosCustomPolicyProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosCustomPolicyProtocol)(nil)).Elem()
}

func (o DdosCustomPolicyProtocolOutput) ToDdosCustomPolicyProtocolOutput() DdosCustomPolicyProtocolOutput {
	return o
}

func (o DdosCustomPolicyProtocolOutput) ToDdosCustomPolicyProtocolOutputWithContext(ctx context.Context) DdosCustomPolicyProtocolOutput {
	return o
}

func (o DdosCustomPolicyProtocolOutput) ToDdosCustomPolicyProtocolPtrOutput() DdosCustomPolicyProtocolPtrOutput {
	return o.ToDdosCustomPolicyProtocolPtrOutputWithContext(context.Background())
}

func (o DdosCustomPolicyProtocolOutput) ToDdosCustomPolicyProtocolPtrOutputWithContext(ctx context.Context) DdosCustomPolicyProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DdosCustomPolicyProtocol) *DdosCustomPolicyProtocol {
		return &v
	}).(DdosCustomPolicyProtocolPtrOutput)
}

func (o DdosCustomPolicyProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DdosCustomPolicyProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DdosCustomPolicyProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DdosCustomPolicyProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DdosCustomPolicyProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DdosCustomPolicyProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DdosCustomPolicyProtocolPtrOutput struct{ *pulumi.OutputState }

func (DdosCustomPolicyProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosCustomPolicyProtocol)(nil)).Elem()
}

func (o DdosCustomPolicyProtocolPtrOutput) ToDdosCustomPolicyProtocolPtrOutput() DdosCustomPolicyProtocolPtrOutput {
	return o
}

func (o DdosCustomPolicyProtocolPtrOutput) ToDdosCustomPolicyProtocolPtrOutputWithContext(ctx context.Context) DdosCustomPolicyProtocolPtrOutput {
	return o
}

func (o DdosCustomPolicyProtocolPtrOutput) Elem() DdosCustomPolicyProtocolOutput {
	return o.ApplyT(func(v *DdosCustomPolicyProtocol) DdosCustomPolicyProtocol {
		if v != nil {
			return *v
		}
		var ret DdosCustomPolicyProtocol
		return ret
	}).(DdosCustomPolicyProtocolOutput)
}

func (o DdosCustomPolicyProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DdosCustomPolicyProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DdosCustomPolicyProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DdosCustomPolicyProtocolInput interface {
	pulumi.Input

	ToDdosCustomPolicyProtocolOutput() DdosCustomPolicyProtocolOutput
	ToDdosCustomPolicyProtocolOutputWithContext(context.Context) DdosCustomPolicyProtocolOutput
}

var ddosCustomPolicyProtocolPtrType = reflect.TypeOf((**DdosCustomPolicyProtocol)(nil)).Elem()

type DdosCustomPolicyProtocolPtrInput interface {
	pulumi.Input

	ToDdosCustomPolicyProtocolPtrOutput() DdosCustomPolicyProtocolPtrOutput
	ToDdosCustomPolicyProtocolPtrOutputWithContext(context.Context) DdosCustomPolicyProtocolPtrOutput
}

type ddosCustomPolicyProtocolPtr string

func DdosCustomPolicyProtocolPtr(v string) DdosCustomPolicyProtocolPtrInput {
	return (*ddosCustomPolicyProtocolPtr)(&v)
}

func (*ddosCustomPolicyProtocolPtr) ElementType() reflect.Type {
	return ddosCustomPolicyProtocolPtrType
}

func (in *ddosCustomPolicyProtocolPtr) ToDdosCustomPolicyProtocolPtrOutput() DdosCustomPolicyProtocolPtrOutput {
	return pulumi.ToOutput(in).(DdosCustomPolicyProtocolPtrOutput)
}

func (in *ddosCustomPolicyProtocolPtr) ToDdosCustomPolicyProtocolPtrOutputWithContext(ctx context.Context) DdosCustomPolicyProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DdosCustomPolicyProtocolPtrOutput)
}

type DdosCustomPolicyTriggerSensitivityOverride string

const (
	DdosCustomPolicyTriggerSensitivityOverrideRelaxed = DdosCustomPolicyTriggerSensitivityOverride("Relaxed")
	DdosCustomPolicyTriggerSensitivityOverrideLow     = DdosCustomPolicyTriggerSensitivityOverride("Low")
	DdosCustomPolicyTriggerSensitivityOverrideDefault = DdosCustomPolicyTriggerSensitivityOverride("Default")
	DdosCustomPolicyTriggerSensitivityOverrideHigh    = DdosCustomPolicyTriggerSensitivityOverride("High")
)

func (DdosCustomPolicyTriggerSensitivityOverride) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosCustomPolicyTriggerSensitivityOverride)(nil)).Elem()
}

func (e DdosCustomPolicyTriggerSensitivityOverride) ToDdosCustomPolicyTriggerSensitivityOverrideOutput() DdosCustomPolicyTriggerSensitivityOverrideOutput {
	return pulumi.ToOutput(e).(DdosCustomPolicyTriggerSensitivityOverrideOutput)
}

func (e DdosCustomPolicyTriggerSensitivityOverride) ToDdosCustomPolicyTriggerSensitivityOverrideOutputWithContext(ctx context.Context) DdosCustomPolicyTriggerSensitivityOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DdosCustomPolicyTriggerSensitivityOverrideOutput)
}

func (e DdosCustomPolicyTriggerSensitivityOverride) ToDdosCustomPolicyTriggerSensitivityOverridePtrOutput() DdosCustomPolicyTriggerSensitivityOverridePtrOutput {
	return e.ToDdosCustomPolicyTriggerSensitivityOverridePtrOutputWithContext(context.Background())
}

func (e DdosCustomPolicyTriggerSensitivityOverride) ToDdosCustomPolicyTriggerSensitivityOverridePtrOutputWithContext(ctx context.Context) DdosCustomPolicyTriggerSensitivityOverridePtrOutput {
	return DdosCustomPolicyTriggerSensitivityOverride(e).ToDdosCustomPolicyTriggerSensitivityOverrideOutputWithContext(ctx).ToDdosCustomPolicyTriggerSensitivityOverridePtrOutputWithContext(ctx)
}

func (e DdosCustomPolicyTriggerSensitivityOverride) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DdosCustomPolicyTriggerSensitivityOverride) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DdosCustomPolicyTriggerSensitivityOverride) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DdosCustomPolicyTriggerSensitivityOverride) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DdosCustomPolicyTriggerSensitivityOverrideOutput struct{ *pulumi.OutputState }

func (DdosCustomPolicyTriggerSensitivityOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosCustomPolicyTriggerSensitivityOverride)(nil)).Elem()
}

func (o DdosCustomPolicyTriggerSensitivityOverrideOutput) ToDdosCustomPolicyTriggerSensitivityOverrideOutput() DdosCustomPolicyTriggerSensitivityOverrideOutput {
	return o
}

func (o DdosCustomPolicyTriggerSensitivityOverrideOutput) ToDdosCustomPolicyTriggerSensitivityOverrideOutputWithContext(ctx context.Context) DdosCustomPolicyTriggerSensitivityOverrideOutput {
	return o
}

func (o DdosCustomPolicyTriggerSensitivityOverrideOutput) ToDdosCustomPolicyTriggerSensitivityOverridePtrOutput() DdosCustomPolicyTriggerSensitivityOverridePtrOutput {
	return o.ToDdosCustomPolicyTriggerSensitivityOverridePtrOutputWithContext(context.Background())
}

func (o DdosCustomPolicyTriggerSensitivityOverrideOutput) ToDdosCustomPolicyTriggerSensitivityOverridePtrOutputWithContext(ctx context.Context) DdosCustomPolicyTriggerSensitivityOverridePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DdosCustomPolicyTriggerSensitivityOverride) *DdosCustomPolicyTriggerSensitivityOverride {
		return &v
	}).(DdosCustomPolicyTriggerSensitivityOverridePtrOutput)
}

func (o DdosCustomPolicyTriggerSensitivityOverrideOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DdosCustomPolicyTriggerSensitivityOverrideOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DdosCustomPolicyTriggerSensitivityOverride) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DdosCustomPolicyTriggerSensitivityOverrideOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DdosCustomPolicyTriggerSensitivityOverrideOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DdosCustomPolicyTriggerSensitivityOverride) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DdosCustomPolicyTriggerSensitivityOverridePtrOutput struct{ *pulumi.OutputState }

func (DdosCustomPolicyTriggerSensitivityOverridePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosCustomPolicyTriggerSensitivityOverride)(nil)).Elem()
}

func (o DdosCustomPolicyTriggerSensitivityOverridePtrOutput) ToDdosCustomPolicyTriggerSensitivityOverridePtrOutput() DdosCustomPolicyTriggerSensitivityOverridePtrOutput {
	return o
}

func (o DdosCustomPolicyTriggerSensitivityOverridePtrOutput) ToDdosCustomPolicyTriggerSensitivityOverridePtrOutputWithContext(ctx context.Context) DdosCustomPolicyTriggerSensitivityOverridePtrOutput {
	return o
}

func (o DdosCustomPolicyTriggerSensitivityOverridePtrOutput) Elem() DdosCustomPolicyTriggerSensitivityOverrideOutput {
	return o.ApplyT(func(v *DdosCustomPolicyTriggerSensitivityOverride) DdosCustomPolicyTriggerSensitivityOverride {
		if v != nil {
			return *v
		}
		var ret DdosCustomPolicyTriggerSensitivityOverride
		return ret
	}).(DdosCustomPolicyTriggerSensitivityOverrideOutput)
}

func (o DdosCustomPolicyTriggerSensitivityOverridePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DdosCustomPolicyTriggerSensitivityOverridePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DdosCustomPolicyTriggerSensitivityOverride) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DdosCustomPolicyTriggerSensitivityOverrideInput interface {
	pulumi.Input

	ToDdosCustomPolicyTriggerSensitivityOverrideOutput() DdosCustomPolicyTriggerSensitivityOverrideOutput
	ToDdosCustomPolicyTriggerSensitivityOverrideOutputWithContext(context.Context) DdosCustomPolicyTriggerSensitivityOverrideOutput
}

var ddosCustomPolicyTriggerSensitivityOverridePtrType = reflect.TypeOf((**DdosCustomPolicyTriggerSensitivityOverride)(nil)).Elem()

type DdosCustomPolicyTriggerSensitivityOverridePtrInput interface {
	pulumi.Input

	ToDdosCustomPolicyTriggerSensitivityOverridePtrOutput() DdosCustomPolicyTriggerSensitivityOverridePtrOutput
	ToDdosCustomPolicyTriggerSensitivityOverridePtrOutputWithContext(context.Context) DdosCustomPolicyTriggerSensitivityOverridePtrOutput
}

type ddosCustomPolicyTriggerSensitivityOverridePtr string

func DdosCustomPolicyTriggerSensitivityOverridePtr(v string) DdosCustomPolicyTriggerSensitivityOverridePtrInput {
	return (*ddosCustomPolicyTriggerSensitivityOverridePtr)(&v)
}

func (*ddosCustomPolicyTriggerSensitivityOverridePtr) ElementType() reflect.Type {
	return ddosCustomPolicyTriggerSensitivityOverridePtrType
}

func (in *ddosCustomPolicyTriggerSensitivityOverridePtr) ToDdosCustomPolicyTriggerSensitivityOverridePtrOutput() DdosCustomPolicyTriggerSensitivityOverridePtrOutput {
	return pulumi.ToOutput(in).(DdosCustomPolicyTriggerSensitivityOverridePtrOutput)
}

func (in *ddosCustomPolicyTriggerSensitivityOverridePtr) ToDdosCustomPolicyTriggerSensitivityOverridePtrOutputWithContext(ctx context.Context) DdosCustomPolicyTriggerSensitivityOverridePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DdosCustomPolicyTriggerSensitivityOverridePtrOutput)
}

type DdosSettingsProtectionCoverage string

const (
	DdosSettingsProtectionCoverageBasic    = DdosSettingsProtectionCoverage("Basic")
	DdosSettingsProtectionCoverageStandard = DdosSettingsProtectionCoverage("Standard")
)

func (DdosSettingsProtectionCoverage) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosSettingsProtectionCoverage)(nil)).Elem()
}

func (e DdosSettingsProtectionCoverage) ToDdosSettingsProtectionCoverageOutput() DdosSettingsProtectionCoverageOutput {
	return pulumi.ToOutput(e).(DdosSettingsProtectionCoverageOutput)
}

func (e DdosSettingsProtectionCoverage) ToDdosSettingsProtectionCoverageOutputWithContext(ctx context.Context) DdosSettingsProtectionCoverageOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DdosSettingsProtectionCoverageOutput)
}

func (e DdosSettingsProtectionCoverage) ToDdosSettingsProtectionCoveragePtrOutput() DdosSettingsProtectionCoveragePtrOutput {
	return e.ToDdosSettingsProtectionCoveragePtrOutputWithContext(context.Background())
}

func (e DdosSettingsProtectionCoverage) ToDdosSettingsProtectionCoveragePtrOutputWithContext(ctx context.Context) DdosSettingsProtectionCoveragePtrOutput {
	return DdosSettingsProtectionCoverage(e).ToDdosSettingsProtectionCoverageOutputWithContext(ctx).ToDdosSettingsProtectionCoveragePtrOutputWithContext(ctx)
}

func (e DdosSettingsProtectionCoverage) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DdosSettingsProtectionCoverage) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DdosSettingsProtectionCoverage) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DdosSettingsProtectionCoverage) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DdosSettingsProtectionCoverageOutput struct{ *pulumi.OutputState }

func (DdosSettingsProtectionCoverageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosSettingsProtectionCoverage)(nil)).Elem()
}

func (o DdosSettingsProtectionCoverageOutput) ToDdosSettingsProtectionCoverageOutput() DdosSettingsProtectionCoverageOutput {
	return o
}

func (o DdosSettingsProtectionCoverageOutput) ToDdosSettingsProtectionCoverageOutputWithContext(ctx context.Context) DdosSettingsProtectionCoverageOutput {
	return o
}

func (o DdosSettingsProtectionCoverageOutput) ToDdosSettingsProtectionCoveragePtrOutput() DdosSettingsProtectionCoveragePtrOutput {
	return o.ToDdosSettingsProtectionCoveragePtrOutputWithContext(context.Background())
}

func (o DdosSettingsProtectionCoverageOutput) ToDdosSettingsProtectionCoveragePtrOutputWithContext(ctx context.Context) DdosSettingsProtectionCoveragePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DdosSettingsProtectionCoverage) *DdosSettingsProtectionCoverage {
		return &v
	}).(DdosSettingsProtectionCoveragePtrOutput)
}

func (o DdosSettingsProtectionCoverageOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DdosSettingsProtectionCoverageOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DdosSettingsProtectionCoverage) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DdosSettingsProtectionCoverageOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DdosSettingsProtectionCoverageOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DdosSettingsProtectionCoverage) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DdosSettingsProtectionCoveragePtrOutput struct{ *pulumi.OutputState }

func (DdosSettingsProtectionCoveragePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosSettingsProtectionCoverage)(nil)).Elem()
}

func (o DdosSettingsProtectionCoveragePtrOutput) ToDdosSettingsProtectionCoveragePtrOutput() DdosSettingsProtectionCoveragePtrOutput {
	return o
}

func (o DdosSettingsProtectionCoveragePtrOutput) ToDdosSettingsProtectionCoveragePtrOutputWithContext(ctx context.Context) DdosSettingsProtectionCoveragePtrOutput {
	return o
}

func (o DdosSettingsProtectionCoveragePtrOutput) Elem() DdosSettingsProtectionCoverageOutput {
	return o.ApplyT(func(v *DdosSettingsProtectionCoverage) DdosSettingsProtectionCoverage {
		if v != nil {
			return *v
		}
		var ret DdosSettingsProtectionCoverage
		return ret
	}).(DdosSettingsProtectionCoverageOutput)
}

func (o DdosSettingsProtectionCoveragePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DdosSettingsProtectionCoveragePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DdosSettingsProtectionCoverage) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DdosSettingsProtectionCoverageInput interface {
	pulumi.Input

	ToDdosSettingsProtectionCoverageOutput() DdosSettingsProtectionCoverageOutput
	ToDdosSettingsProtectionCoverageOutputWithContext(context.Context) DdosSettingsProtectionCoverageOutput
}

var ddosSettingsProtectionCoveragePtrType = reflect.TypeOf((**DdosSettingsProtectionCoverage)(nil)).Elem()

type DdosSettingsProtectionCoveragePtrInput interface {
	pulumi.Input

	ToDdosSettingsProtectionCoveragePtrOutput() DdosSettingsProtectionCoveragePtrOutput
	ToDdosSettingsProtectionCoveragePtrOutputWithContext(context.Context) DdosSettingsProtectionCoveragePtrOutput
}

type ddosSettingsProtectionCoveragePtr string

func DdosSettingsProtectionCoveragePtr(v string) DdosSettingsProtectionCoveragePtrInput {
	return (*ddosSettingsProtectionCoveragePtr)(&v)
}

func (*ddosSettingsProtectionCoveragePtr) ElementType() reflect.Type {
	return ddosSettingsProtectionCoveragePtrType
}

func (in *ddosSettingsProtectionCoveragePtr) ToDdosSettingsProtectionCoveragePtrOutput() DdosSettingsProtectionCoveragePtrOutput {
	return pulumi.ToOutput(in).(DdosSettingsProtectionCoveragePtrOutput)
}

func (in *ddosSettingsProtectionCoveragePtr) ToDdosSettingsProtectionCoveragePtrOutputWithContext(ctx context.Context) DdosSettingsProtectionCoveragePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DdosSettingsProtectionCoveragePtrOutput)
}

type DeleteOptions string

const (
	DeleteOptionsDelete = DeleteOptions("Delete")
	DeleteOptionsDetach = DeleteOptions("Detach")
)

func (DeleteOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteOptions)(nil)).Elem()
}

func (e DeleteOptions) ToDeleteOptionsOutput() DeleteOptionsOutput {
	return pulumi.ToOutput(e).(DeleteOptionsOutput)
}

func (e DeleteOptions) ToDeleteOptionsOutputWithContext(ctx context.Context) DeleteOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeleteOptionsOutput)
}

func (e DeleteOptions) ToDeleteOptionsPtrOutput() DeleteOptionsPtrOutput {
	return e.ToDeleteOptionsPtrOutputWithContext(context.Background())
}

func (e DeleteOptions) ToDeleteOptionsPtrOutputWithContext(ctx context.Context) DeleteOptionsPtrOutput {
	return DeleteOptions(e).ToDeleteOptionsOutputWithContext(ctx).ToDeleteOptionsPtrOutputWithContext(ctx)
}

func (e DeleteOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeleteOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeleteOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeleteOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeleteOptionsOutput struct{ *pulumi.OutputState }

func (DeleteOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteOptions)(nil)).Elem()
}

func (o DeleteOptionsOutput) ToDeleteOptionsOutput() DeleteOptionsOutput {
	return o
}

func (o DeleteOptionsOutput) ToDeleteOptionsOutputWithContext(ctx context.Context) DeleteOptionsOutput {
	return o
}

func (o DeleteOptionsOutput) ToDeleteOptionsPtrOutput() DeleteOptionsPtrOutput {
	return o.ToDeleteOptionsPtrOutputWithContext(context.Background())
}

func (o DeleteOptionsOutput) ToDeleteOptionsPtrOutputWithContext(ctx context.Context) DeleteOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeleteOptions) *DeleteOptions {
		return &v
	}).(DeleteOptionsPtrOutput)
}

func (o DeleteOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeleteOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeleteOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeleteOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeleteOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeleteOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeleteOptionsPtrOutput struct{ *pulumi.OutputState }

func (DeleteOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteOptions)(nil)).Elem()
}

func (o DeleteOptionsPtrOutput) ToDeleteOptionsPtrOutput() DeleteOptionsPtrOutput {
	return o
}

func (o DeleteOptionsPtrOutput) ToDeleteOptionsPtrOutputWithContext(ctx context.Context) DeleteOptionsPtrOutput {
	return o
}

func (o DeleteOptionsPtrOutput) Elem() DeleteOptionsOutput {
	return o.ApplyT(func(v *DeleteOptions) DeleteOptions {
		if v != nil {
			return *v
		}
		var ret DeleteOptions
		return ret
	}).(DeleteOptionsOutput)
}

func (o DeleteOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeleteOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeleteOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeleteOptionsInput interface {
	pulumi.Input

	ToDeleteOptionsOutput() DeleteOptionsOutput
	ToDeleteOptionsOutputWithContext(context.Context) DeleteOptionsOutput
}

var deleteOptionsPtrType = reflect.TypeOf((**DeleteOptions)(nil)).Elem()

type DeleteOptionsPtrInput interface {
	pulumi.Input

	ToDeleteOptionsPtrOutput() DeleteOptionsPtrOutput
	ToDeleteOptionsPtrOutputWithContext(context.Context) DeleteOptionsPtrOutput
}

type deleteOptionsPtr string

func DeleteOptionsPtr(v string) DeleteOptionsPtrInput {
	return (*deleteOptionsPtr)(&v)
}

func (*deleteOptionsPtr) ElementType() reflect.Type {
	return deleteOptionsPtrType
}

func (in *deleteOptionsPtr) ToDeleteOptionsPtrOutput() DeleteOptionsPtrOutput {
	return pulumi.ToOutput(in).(DeleteOptionsPtrOutput)
}

func (in *deleteOptionsPtr) ToDeleteOptionsPtrOutputWithContext(ctx context.Context) DeleteOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeleteOptionsPtrOutput)
}

type DestinationPortBehavior string

const (
	DestinationPortBehaviorNone              = DestinationPortBehavior("None")
	DestinationPortBehaviorListenIfAvailable = DestinationPortBehavior("ListenIfAvailable")
)

func (DestinationPortBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationPortBehavior)(nil)).Elem()
}

func (e DestinationPortBehavior) ToDestinationPortBehaviorOutput() DestinationPortBehaviorOutput {
	return pulumi.ToOutput(e).(DestinationPortBehaviorOutput)
}

func (e DestinationPortBehavior) ToDestinationPortBehaviorOutputWithContext(ctx context.Context) DestinationPortBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DestinationPortBehaviorOutput)
}

func (e DestinationPortBehavior) ToDestinationPortBehaviorPtrOutput() DestinationPortBehaviorPtrOutput {
	return e.ToDestinationPortBehaviorPtrOutputWithContext(context.Background())
}

func (e DestinationPortBehavior) ToDestinationPortBehaviorPtrOutputWithContext(ctx context.Context) DestinationPortBehaviorPtrOutput {
	return DestinationPortBehavior(e).ToDestinationPortBehaviorOutputWithContext(ctx).ToDestinationPortBehaviorPtrOutputWithContext(ctx)
}

func (e DestinationPortBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DestinationPortBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DestinationPortBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DestinationPortBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DestinationPortBehaviorOutput struct{ *pulumi.OutputState }

func (DestinationPortBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationPortBehavior)(nil)).Elem()
}

func (o DestinationPortBehaviorOutput) ToDestinationPortBehaviorOutput() DestinationPortBehaviorOutput {
	return o
}

func (o DestinationPortBehaviorOutput) ToDestinationPortBehaviorOutputWithContext(ctx context.Context) DestinationPortBehaviorOutput {
	return o
}

func (o DestinationPortBehaviorOutput) ToDestinationPortBehaviorPtrOutput() DestinationPortBehaviorPtrOutput {
	return o.ToDestinationPortBehaviorPtrOutputWithContext(context.Background())
}

func (o DestinationPortBehaviorOutput) ToDestinationPortBehaviorPtrOutputWithContext(ctx context.Context) DestinationPortBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DestinationPortBehavior) *DestinationPortBehavior {
		return &v
	}).(DestinationPortBehaviorPtrOutput)
}

func (o DestinationPortBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DestinationPortBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DestinationPortBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DestinationPortBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DestinationPortBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DestinationPortBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DestinationPortBehaviorPtrOutput struct{ *pulumi.OutputState }

func (DestinationPortBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationPortBehavior)(nil)).Elem()
}

func (o DestinationPortBehaviorPtrOutput) ToDestinationPortBehaviorPtrOutput() DestinationPortBehaviorPtrOutput {
	return o
}

func (o DestinationPortBehaviorPtrOutput) ToDestinationPortBehaviorPtrOutputWithContext(ctx context.Context) DestinationPortBehaviorPtrOutput {
	return o
}

func (o DestinationPortBehaviorPtrOutput) Elem() DestinationPortBehaviorOutput {
	return o.ApplyT(func(v *DestinationPortBehavior) DestinationPortBehavior {
		if v != nil {
			return *v
		}
		var ret DestinationPortBehavior
		return ret
	}).(DestinationPortBehaviorOutput)
}

func (o DestinationPortBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DestinationPortBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DestinationPortBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DestinationPortBehaviorInput interface {
	pulumi.Input

	ToDestinationPortBehaviorOutput() DestinationPortBehaviorOutput
	ToDestinationPortBehaviorOutputWithContext(context.Context) DestinationPortBehaviorOutput
}

var destinationPortBehaviorPtrType = reflect.TypeOf((**DestinationPortBehavior)(nil)).Elem()

type DestinationPortBehaviorPtrInput interface {
	pulumi.Input

	ToDestinationPortBehaviorPtrOutput() DestinationPortBehaviorPtrOutput
	ToDestinationPortBehaviorPtrOutputWithContext(context.Context) DestinationPortBehaviorPtrOutput
}

type destinationPortBehaviorPtr string

func DestinationPortBehaviorPtr(v string) DestinationPortBehaviorPtrInput {
	return (*destinationPortBehaviorPtr)(&v)
}

func (*destinationPortBehaviorPtr) ElementType() reflect.Type {
	return destinationPortBehaviorPtrType
}

func (in *destinationPortBehaviorPtr) ToDestinationPortBehaviorPtrOutput() DestinationPortBehaviorPtrOutput {
	return pulumi.ToOutput(in).(DestinationPortBehaviorPtrOutput)
}

func (in *destinationPortBehaviorPtr) ToDestinationPortBehaviorPtrOutputWithContext(ctx context.Context) DestinationPortBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DestinationPortBehaviorPtrOutput)
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

type EndpointType string

const (
	EndpointTypeAzureVM             = EndpointType("AzureVM")
	EndpointTypeAzureVNet           = EndpointType("AzureVNet")
	EndpointTypeAzureSubnet         = EndpointType("AzureSubnet")
	EndpointTypeExternalAddress     = EndpointType("ExternalAddress")
	EndpointTypeMMAWorkspaceMachine = EndpointType("MMAWorkspaceMachine")
	EndpointTypeMMAWorkspaceNetwork = EndpointType("MMAWorkspaceNetwork")
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
	ExpressRouteCircuitSkuTierBasic    = ExpressRouteCircuitSkuTier("Basic")
	ExpressRouteCircuitSkuTierLocal    = ExpressRouteCircuitSkuTier("Local")
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

type ExpressRouteLinkAdminState string

const (
	ExpressRouteLinkAdminStateEnabled  = ExpressRouteLinkAdminState("Enabled")
	ExpressRouteLinkAdminStateDisabled = ExpressRouteLinkAdminState("Disabled")
)

func (ExpressRouteLinkAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteLinkAdminState)(nil)).Elem()
}

func (e ExpressRouteLinkAdminState) ToExpressRouteLinkAdminStateOutput() ExpressRouteLinkAdminStateOutput {
	return pulumi.ToOutput(e).(ExpressRouteLinkAdminStateOutput)
}

func (e ExpressRouteLinkAdminState) ToExpressRouteLinkAdminStateOutputWithContext(ctx context.Context) ExpressRouteLinkAdminStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRouteLinkAdminStateOutput)
}

func (e ExpressRouteLinkAdminState) ToExpressRouteLinkAdminStatePtrOutput() ExpressRouteLinkAdminStatePtrOutput {
	return e.ToExpressRouteLinkAdminStatePtrOutputWithContext(context.Background())
}

func (e ExpressRouteLinkAdminState) ToExpressRouteLinkAdminStatePtrOutputWithContext(ctx context.Context) ExpressRouteLinkAdminStatePtrOutput {
	return ExpressRouteLinkAdminState(e).ToExpressRouteLinkAdminStateOutputWithContext(ctx).ToExpressRouteLinkAdminStatePtrOutputWithContext(ctx)
}

func (e ExpressRouteLinkAdminState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteLinkAdminState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteLinkAdminState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRouteLinkAdminState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRouteLinkAdminStateOutput struct{ *pulumi.OutputState }

func (ExpressRouteLinkAdminStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteLinkAdminState)(nil)).Elem()
}

func (o ExpressRouteLinkAdminStateOutput) ToExpressRouteLinkAdminStateOutput() ExpressRouteLinkAdminStateOutput {
	return o
}

func (o ExpressRouteLinkAdminStateOutput) ToExpressRouteLinkAdminStateOutputWithContext(ctx context.Context) ExpressRouteLinkAdminStateOutput {
	return o
}

func (o ExpressRouteLinkAdminStateOutput) ToExpressRouteLinkAdminStatePtrOutput() ExpressRouteLinkAdminStatePtrOutput {
	return o.ToExpressRouteLinkAdminStatePtrOutputWithContext(context.Background())
}

func (o ExpressRouteLinkAdminStateOutput) ToExpressRouteLinkAdminStatePtrOutputWithContext(ctx context.Context) ExpressRouteLinkAdminStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteLinkAdminState) *ExpressRouteLinkAdminState {
		return &v
	}).(ExpressRouteLinkAdminStatePtrOutput)
}

func (o ExpressRouteLinkAdminStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRouteLinkAdminStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteLinkAdminState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRouteLinkAdminStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteLinkAdminStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteLinkAdminState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteLinkAdminStatePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteLinkAdminStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteLinkAdminState)(nil)).Elem()
}

func (o ExpressRouteLinkAdminStatePtrOutput) ToExpressRouteLinkAdminStatePtrOutput() ExpressRouteLinkAdminStatePtrOutput {
	return o
}

func (o ExpressRouteLinkAdminStatePtrOutput) ToExpressRouteLinkAdminStatePtrOutputWithContext(ctx context.Context) ExpressRouteLinkAdminStatePtrOutput {
	return o
}

func (o ExpressRouteLinkAdminStatePtrOutput) Elem() ExpressRouteLinkAdminStateOutput {
	return o.ApplyT(func(v *ExpressRouteLinkAdminState) ExpressRouteLinkAdminState {
		if v != nil {
			return *v
		}
		var ret ExpressRouteLinkAdminState
		return ret
	}).(ExpressRouteLinkAdminStateOutput)
}

func (o ExpressRouteLinkAdminStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteLinkAdminStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRouteLinkAdminState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRouteLinkAdminStateInput interface {
	pulumi.Input

	ToExpressRouteLinkAdminStateOutput() ExpressRouteLinkAdminStateOutput
	ToExpressRouteLinkAdminStateOutputWithContext(context.Context) ExpressRouteLinkAdminStateOutput
}

var expressRouteLinkAdminStatePtrType = reflect.TypeOf((**ExpressRouteLinkAdminState)(nil)).Elem()

type ExpressRouteLinkAdminStatePtrInput interface {
	pulumi.Input

	ToExpressRouteLinkAdminStatePtrOutput() ExpressRouteLinkAdminStatePtrOutput
	ToExpressRouteLinkAdminStatePtrOutputWithContext(context.Context) ExpressRouteLinkAdminStatePtrOutput
}

type expressRouteLinkAdminStatePtr string

func ExpressRouteLinkAdminStatePtr(v string) ExpressRouteLinkAdminStatePtrInput {
	return (*expressRouteLinkAdminStatePtr)(&v)
}

func (*expressRouteLinkAdminStatePtr) ElementType() reflect.Type {
	return expressRouteLinkAdminStatePtrType
}

func (in *expressRouteLinkAdminStatePtr) ToExpressRouteLinkAdminStatePtrOutput() ExpressRouteLinkAdminStatePtrOutput {
	return pulumi.ToOutput(in).(ExpressRouteLinkAdminStatePtrOutput)
}

func (in *expressRouteLinkAdminStatePtr) ToExpressRouteLinkAdminStatePtrOutputWithContext(ctx context.Context) ExpressRouteLinkAdminStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRouteLinkAdminStatePtrOutput)
}

type ExpressRouteLinkMacSecCipher string

const (
	ExpressRouteLinkMacSecCipherGcmAes256    = ExpressRouteLinkMacSecCipher("GcmAes256")
	ExpressRouteLinkMacSecCipherGcmAes128    = ExpressRouteLinkMacSecCipher("GcmAes128")
	ExpressRouteLinkMacSecCipherGcmAesXpn128 = ExpressRouteLinkMacSecCipher("GcmAesXpn128")
	ExpressRouteLinkMacSecCipherGcmAesXpn256 = ExpressRouteLinkMacSecCipher("GcmAesXpn256")
)

func (ExpressRouteLinkMacSecCipher) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteLinkMacSecCipher)(nil)).Elem()
}

func (e ExpressRouteLinkMacSecCipher) ToExpressRouteLinkMacSecCipherOutput() ExpressRouteLinkMacSecCipherOutput {
	return pulumi.ToOutput(e).(ExpressRouteLinkMacSecCipherOutput)
}

func (e ExpressRouteLinkMacSecCipher) ToExpressRouteLinkMacSecCipherOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecCipherOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRouteLinkMacSecCipherOutput)
}

func (e ExpressRouteLinkMacSecCipher) ToExpressRouteLinkMacSecCipherPtrOutput() ExpressRouteLinkMacSecCipherPtrOutput {
	return e.ToExpressRouteLinkMacSecCipherPtrOutputWithContext(context.Background())
}

func (e ExpressRouteLinkMacSecCipher) ToExpressRouteLinkMacSecCipherPtrOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecCipherPtrOutput {
	return ExpressRouteLinkMacSecCipher(e).ToExpressRouteLinkMacSecCipherOutputWithContext(ctx).ToExpressRouteLinkMacSecCipherPtrOutputWithContext(ctx)
}

func (e ExpressRouteLinkMacSecCipher) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteLinkMacSecCipher) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteLinkMacSecCipher) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRouteLinkMacSecCipher) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRouteLinkMacSecCipherOutput struct{ *pulumi.OutputState }

func (ExpressRouteLinkMacSecCipherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteLinkMacSecCipher)(nil)).Elem()
}

func (o ExpressRouteLinkMacSecCipherOutput) ToExpressRouteLinkMacSecCipherOutput() ExpressRouteLinkMacSecCipherOutput {
	return o
}

func (o ExpressRouteLinkMacSecCipherOutput) ToExpressRouteLinkMacSecCipherOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecCipherOutput {
	return o
}

func (o ExpressRouteLinkMacSecCipherOutput) ToExpressRouteLinkMacSecCipherPtrOutput() ExpressRouteLinkMacSecCipherPtrOutput {
	return o.ToExpressRouteLinkMacSecCipherPtrOutputWithContext(context.Background())
}

func (o ExpressRouteLinkMacSecCipherOutput) ToExpressRouteLinkMacSecCipherPtrOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecCipherPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteLinkMacSecCipher) *ExpressRouteLinkMacSecCipher {
		return &v
	}).(ExpressRouteLinkMacSecCipherPtrOutput)
}

func (o ExpressRouteLinkMacSecCipherOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRouteLinkMacSecCipherOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteLinkMacSecCipher) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRouteLinkMacSecCipherOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteLinkMacSecCipherOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteLinkMacSecCipher) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteLinkMacSecCipherPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteLinkMacSecCipherPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteLinkMacSecCipher)(nil)).Elem()
}

func (o ExpressRouteLinkMacSecCipherPtrOutput) ToExpressRouteLinkMacSecCipherPtrOutput() ExpressRouteLinkMacSecCipherPtrOutput {
	return o
}

func (o ExpressRouteLinkMacSecCipherPtrOutput) ToExpressRouteLinkMacSecCipherPtrOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecCipherPtrOutput {
	return o
}

func (o ExpressRouteLinkMacSecCipherPtrOutput) Elem() ExpressRouteLinkMacSecCipherOutput {
	return o.ApplyT(func(v *ExpressRouteLinkMacSecCipher) ExpressRouteLinkMacSecCipher {
		if v != nil {
			return *v
		}
		var ret ExpressRouteLinkMacSecCipher
		return ret
	}).(ExpressRouteLinkMacSecCipherOutput)
}

func (o ExpressRouteLinkMacSecCipherPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteLinkMacSecCipherPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRouteLinkMacSecCipher) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRouteLinkMacSecCipherInput interface {
	pulumi.Input

	ToExpressRouteLinkMacSecCipherOutput() ExpressRouteLinkMacSecCipherOutput
	ToExpressRouteLinkMacSecCipherOutputWithContext(context.Context) ExpressRouteLinkMacSecCipherOutput
}

var expressRouteLinkMacSecCipherPtrType = reflect.TypeOf((**ExpressRouteLinkMacSecCipher)(nil)).Elem()

type ExpressRouteLinkMacSecCipherPtrInput interface {
	pulumi.Input

	ToExpressRouteLinkMacSecCipherPtrOutput() ExpressRouteLinkMacSecCipherPtrOutput
	ToExpressRouteLinkMacSecCipherPtrOutputWithContext(context.Context) ExpressRouteLinkMacSecCipherPtrOutput
}

type expressRouteLinkMacSecCipherPtr string

func ExpressRouteLinkMacSecCipherPtr(v string) ExpressRouteLinkMacSecCipherPtrInput {
	return (*expressRouteLinkMacSecCipherPtr)(&v)
}

func (*expressRouteLinkMacSecCipherPtr) ElementType() reflect.Type {
	return expressRouteLinkMacSecCipherPtrType
}

func (in *expressRouteLinkMacSecCipherPtr) ToExpressRouteLinkMacSecCipherPtrOutput() ExpressRouteLinkMacSecCipherPtrOutput {
	return pulumi.ToOutput(in).(ExpressRouteLinkMacSecCipherPtrOutput)
}

func (in *expressRouteLinkMacSecCipherPtr) ToExpressRouteLinkMacSecCipherPtrOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecCipherPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRouteLinkMacSecCipherPtrOutput)
}

type ExpressRouteLinkMacSecSciState string

const (
	ExpressRouteLinkMacSecSciStateDisabled = ExpressRouteLinkMacSecSciState("Disabled")
	ExpressRouteLinkMacSecSciStateEnabled  = ExpressRouteLinkMacSecSciState("Enabled")
)

func (ExpressRouteLinkMacSecSciState) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteLinkMacSecSciState)(nil)).Elem()
}

func (e ExpressRouteLinkMacSecSciState) ToExpressRouteLinkMacSecSciStateOutput() ExpressRouteLinkMacSecSciStateOutput {
	return pulumi.ToOutput(e).(ExpressRouteLinkMacSecSciStateOutput)
}

func (e ExpressRouteLinkMacSecSciState) ToExpressRouteLinkMacSecSciStateOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecSciStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRouteLinkMacSecSciStateOutput)
}

func (e ExpressRouteLinkMacSecSciState) ToExpressRouteLinkMacSecSciStatePtrOutput() ExpressRouteLinkMacSecSciStatePtrOutput {
	return e.ToExpressRouteLinkMacSecSciStatePtrOutputWithContext(context.Background())
}

func (e ExpressRouteLinkMacSecSciState) ToExpressRouteLinkMacSecSciStatePtrOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecSciStatePtrOutput {
	return ExpressRouteLinkMacSecSciState(e).ToExpressRouteLinkMacSecSciStateOutputWithContext(ctx).ToExpressRouteLinkMacSecSciStatePtrOutputWithContext(ctx)
}

func (e ExpressRouteLinkMacSecSciState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteLinkMacSecSciState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRouteLinkMacSecSciState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRouteLinkMacSecSciState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRouteLinkMacSecSciStateOutput struct{ *pulumi.OutputState }

func (ExpressRouteLinkMacSecSciStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteLinkMacSecSciState)(nil)).Elem()
}

func (o ExpressRouteLinkMacSecSciStateOutput) ToExpressRouteLinkMacSecSciStateOutput() ExpressRouteLinkMacSecSciStateOutput {
	return o
}

func (o ExpressRouteLinkMacSecSciStateOutput) ToExpressRouteLinkMacSecSciStateOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecSciStateOutput {
	return o
}

func (o ExpressRouteLinkMacSecSciStateOutput) ToExpressRouteLinkMacSecSciStatePtrOutput() ExpressRouteLinkMacSecSciStatePtrOutput {
	return o.ToExpressRouteLinkMacSecSciStatePtrOutputWithContext(context.Background())
}

func (o ExpressRouteLinkMacSecSciStateOutput) ToExpressRouteLinkMacSecSciStatePtrOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecSciStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteLinkMacSecSciState) *ExpressRouteLinkMacSecSciState {
		return &v
	}).(ExpressRouteLinkMacSecSciStatePtrOutput)
}

func (o ExpressRouteLinkMacSecSciStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRouteLinkMacSecSciStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteLinkMacSecSciState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRouteLinkMacSecSciStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteLinkMacSecSciStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRouteLinkMacSecSciState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteLinkMacSecSciStatePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteLinkMacSecSciStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteLinkMacSecSciState)(nil)).Elem()
}

func (o ExpressRouteLinkMacSecSciStatePtrOutput) ToExpressRouteLinkMacSecSciStatePtrOutput() ExpressRouteLinkMacSecSciStatePtrOutput {
	return o
}

func (o ExpressRouteLinkMacSecSciStatePtrOutput) ToExpressRouteLinkMacSecSciStatePtrOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecSciStatePtrOutput {
	return o
}

func (o ExpressRouteLinkMacSecSciStatePtrOutput) Elem() ExpressRouteLinkMacSecSciStateOutput {
	return o.ApplyT(func(v *ExpressRouteLinkMacSecSciState) ExpressRouteLinkMacSecSciState {
		if v != nil {
			return *v
		}
		var ret ExpressRouteLinkMacSecSciState
		return ret
	}).(ExpressRouteLinkMacSecSciStateOutput)
}

func (o ExpressRouteLinkMacSecSciStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRouteLinkMacSecSciStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRouteLinkMacSecSciState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRouteLinkMacSecSciStateInput interface {
	pulumi.Input

	ToExpressRouteLinkMacSecSciStateOutput() ExpressRouteLinkMacSecSciStateOutput
	ToExpressRouteLinkMacSecSciStateOutputWithContext(context.Context) ExpressRouteLinkMacSecSciStateOutput
}

var expressRouteLinkMacSecSciStatePtrType = reflect.TypeOf((**ExpressRouteLinkMacSecSciState)(nil)).Elem()

type ExpressRouteLinkMacSecSciStatePtrInput interface {
	pulumi.Input

	ToExpressRouteLinkMacSecSciStatePtrOutput() ExpressRouteLinkMacSecSciStatePtrOutput
	ToExpressRouteLinkMacSecSciStatePtrOutputWithContext(context.Context) ExpressRouteLinkMacSecSciStatePtrOutput
}

type expressRouteLinkMacSecSciStatePtr string

func ExpressRouteLinkMacSecSciStatePtr(v string) ExpressRouteLinkMacSecSciStatePtrInput {
	return (*expressRouteLinkMacSecSciStatePtr)(&v)
}

func (*expressRouteLinkMacSecSciStatePtr) ElementType() reflect.Type {
	return expressRouteLinkMacSecSciStatePtrType
}

func (in *expressRouteLinkMacSecSciStatePtr) ToExpressRouteLinkMacSecSciStatePtrOutput() ExpressRouteLinkMacSecSciStatePtrOutput {
	return pulumi.ToOutput(in).(ExpressRouteLinkMacSecSciStatePtrOutput)
}

func (in *expressRouteLinkMacSecSciStatePtr) ToExpressRouteLinkMacSecSciStatePtrOutputWithContext(ctx context.Context) ExpressRouteLinkMacSecSciStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRouteLinkMacSecSciStatePtrOutput)
}

type ExpressRoutePeeringState string

const (
	ExpressRoutePeeringStateDisabled = ExpressRoutePeeringState("Disabled")
	ExpressRoutePeeringStateEnabled  = ExpressRoutePeeringState("Enabled")
)

func (ExpressRoutePeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRoutePeeringState)(nil)).Elem()
}

func (e ExpressRoutePeeringState) ToExpressRoutePeeringStateOutput() ExpressRoutePeeringStateOutput {
	return pulumi.ToOutput(e).(ExpressRoutePeeringStateOutput)
}

func (e ExpressRoutePeeringState) ToExpressRoutePeeringStateOutputWithContext(ctx context.Context) ExpressRoutePeeringStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRoutePeeringStateOutput)
}

func (e ExpressRoutePeeringState) ToExpressRoutePeeringStatePtrOutput() ExpressRoutePeeringStatePtrOutput {
	return e.ToExpressRoutePeeringStatePtrOutputWithContext(context.Background())
}

func (e ExpressRoutePeeringState) ToExpressRoutePeeringStatePtrOutputWithContext(ctx context.Context) ExpressRoutePeeringStatePtrOutput {
	return ExpressRoutePeeringState(e).ToExpressRoutePeeringStateOutputWithContext(ctx).ToExpressRoutePeeringStatePtrOutputWithContext(ctx)
}

func (e ExpressRoutePeeringState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRoutePeeringState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRoutePeeringState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRoutePeeringState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRoutePeeringStateOutput struct{ *pulumi.OutputState }

func (ExpressRoutePeeringStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRoutePeeringState)(nil)).Elem()
}

func (o ExpressRoutePeeringStateOutput) ToExpressRoutePeeringStateOutput() ExpressRoutePeeringStateOutput {
	return o
}

func (o ExpressRoutePeeringStateOutput) ToExpressRoutePeeringStateOutputWithContext(ctx context.Context) ExpressRoutePeeringStateOutput {
	return o
}

func (o ExpressRoutePeeringStateOutput) ToExpressRoutePeeringStatePtrOutput() ExpressRoutePeeringStatePtrOutput {
	return o.ToExpressRoutePeeringStatePtrOutputWithContext(context.Background())
}

func (o ExpressRoutePeeringStateOutput) ToExpressRoutePeeringStatePtrOutputWithContext(ctx context.Context) ExpressRoutePeeringStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRoutePeeringState) *ExpressRoutePeeringState {
		return &v
	}).(ExpressRoutePeeringStatePtrOutput)
}

func (o ExpressRoutePeeringStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRoutePeeringStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRoutePeeringState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRoutePeeringStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRoutePeeringStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRoutePeeringState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRoutePeeringStatePtrOutput struct{ *pulumi.OutputState }

func (ExpressRoutePeeringStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePeeringState)(nil)).Elem()
}

func (o ExpressRoutePeeringStatePtrOutput) ToExpressRoutePeeringStatePtrOutput() ExpressRoutePeeringStatePtrOutput {
	return o
}

func (o ExpressRoutePeeringStatePtrOutput) ToExpressRoutePeeringStatePtrOutputWithContext(ctx context.Context) ExpressRoutePeeringStatePtrOutput {
	return o
}

func (o ExpressRoutePeeringStatePtrOutput) Elem() ExpressRoutePeeringStateOutput {
	return o.ApplyT(func(v *ExpressRoutePeeringState) ExpressRoutePeeringState {
		if v != nil {
			return *v
		}
		var ret ExpressRoutePeeringState
		return ret
	}).(ExpressRoutePeeringStateOutput)
}

func (o ExpressRoutePeeringStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRoutePeeringStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRoutePeeringState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRoutePeeringStateInput interface {
	pulumi.Input

	ToExpressRoutePeeringStateOutput() ExpressRoutePeeringStateOutput
	ToExpressRoutePeeringStateOutputWithContext(context.Context) ExpressRoutePeeringStateOutput
}

var expressRoutePeeringStatePtrType = reflect.TypeOf((**ExpressRoutePeeringState)(nil)).Elem()

type ExpressRoutePeeringStatePtrInput interface {
	pulumi.Input

	ToExpressRoutePeeringStatePtrOutput() ExpressRoutePeeringStatePtrOutput
	ToExpressRoutePeeringStatePtrOutputWithContext(context.Context) ExpressRoutePeeringStatePtrOutput
}

type expressRoutePeeringStatePtr string

func ExpressRoutePeeringStatePtr(v string) ExpressRoutePeeringStatePtrInput {
	return (*expressRoutePeeringStatePtr)(&v)
}

func (*expressRoutePeeringStatePtr) ElementType() reflect.Type {
	return expressRoutePeeringStatePtrType
}

func (in *expressRoutePeeringStatePtr) ToExpressRoutePeeringStatePtrOutput() ExpressRoutePeeringStatePtrOutput {
	return pulumi.ToOutput(in).(ExpressRoutePeeringStatePtrOutput)
}

func (in *expressRoutePeeringStatePtr) ToExpressRoutePeeringStatePtrOutputWithContext(ctx context.Context) ExpressRoutePeeringStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRoutePeeringStatePtrOutput)
}

type ExpressRoutePeeringType string

const (
	ExpressRoutePeeringTypeAzurePublicPeering  = ExpressRoutePeeringType("AzurePublicPeering")
	ExpressRoutePeeringTypeAzurePrivatePeering = ExpressRoutePeeringType("AzurePrivatePeering")
	ExpressRoutePeeringTypeMicrosoftPeering    = ExpressRoutePeeringType("MicrosoftPeering")
)

func (ExpressRoutePeeringType) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRoutePeeringType)(nil)).Elem()
}

func (e ExpressRoutePeeringType) ToExpressRoutePeeringTypeOutput() ExpressRoutePeeringTypeOutput {
	return pulumi.ToOutput(e).(ExpressRoutePeeringTypeOutput)
}

func (e ExpressRoutePeeringType) ToExpressRoutePeeringTypeOutputWithContext(ctx context.Context) ExpressRoutePeeringTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRoutePeeringTypeOutput)
}

func (e ExpressRoutePeeringType) ToExpressRoutePeeringTypePtrOutput() ExpressRoutePeeringTypePtrOutput {
	return e.ToExpressRoutePeeringTypePtrOutputWithContext(context.Background())
}

func (e ExpressRoutePeeringType) ToExpressRoutePeeringTypePtrOutputWithContext(ctx context.Context) ExpressRoutePeeringTypePtrOutput {
	return ExpressRoutePeeringType(e).ToExpressRoutePeeringTypeOutputWithContext(ctx).ToExpressRoutePeeringTypePtrOutputWithContext(ctx)
}

func (e ExpressRoutePeeringType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRoutePeeringType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRoutePeeringType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRoutePeeringType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRoutePeeringTypeOutput struct{ *pulumi.OutputState }

func (ExpressRoutePeeringTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRoutePeeringType)(nil)).Elem()
}

func (o ExpressRoutePeeringTypeOutput) ToExpressRoutePeeringTypeOutput() ExpressRoutePeeringTypeOutput {
	return o
}

func (o ExpressRoutePeeringTypeOutput) ToExpressRoutePeeringTypeOutputWithContext(ctx context.Context) ExpressRoutePeeringTypeOutput {
	return o
}

func (o ExpressRoutePeeringTypeOutput) ToExpressRoutePeeringTypePtrOutput() ExpressRoutePeeringTypePtrOutput {
	return o.ToExpressRoutePeeringTypePtrOutputWithContext(context.Background())
}

func (o ExpressRoutePeeringTypeOutput) ToExpressRoutePeeringTypePtrOutputWithContext(ctx context.Context) ExpressRoutePeeringTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRoutePeeringType) *ExpressRoutePeeringType {
		return &v
	}).(ExpressRoutePeeringTypePtrOutput)
}

func (o ExpressRoutePeeringTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRoutePeeringTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRoutePeeringType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRoutePeeringTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRoutePeeringTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRoutePeeringType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRoutePeeringTypePtrOutput struct{ *pulumi.OutputState }

func (ExpressRoutePeeringTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePeeringType)(nil)).Elem()
}

func (o ExpressRoutePeeringTypePtrOutput) ToExpressRoutePeeringTypePtrOutput() ExpressRoutePeeringTypePtrOutput {
	return o
}

func (o ExpressRoutePeeringTypePtrOutput) ToExpressRoutePeeringTypePtrOutputWithContext(ctx context.Context) ExpressRoutePeeringTypePtrOutput {
	return o
}

func (o ExpressRoutePeeringTypePtrOutput) Elem() ExpressRoutePeeringTypeOutput {
	return o.ApplyT(func(v *ExpressRoutePeeringType) ExpressRoutePeeringType {
		if v != nil {
			return *v
		}
		var ret ExpressRoutePeeringType
		return ret
	}).(ExpressRoutePeeringTypeOutput)
}

func (o ExpressRoutePeeringTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRoutePeeringTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRoutePeeringType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRoutePeeringTypeInput interface {
	pulumi.Input

	ToExpressRoutePeeringTypeOutput() ExpressRoutePeeringTypeOutput
	ToExpressRoutePeeringTypeOutputWithContext(context.Context) ExpressRoutePeeringTypeOutput
}

var expressRoutePeeringTypePtrType = reflect.TypeOf((**ExpressRoutePeeringType)(nil)).Elem()

type ExpressRoutePeeringTypePtrInput interface {
	pulumi.Input

	ToExpressRoutePeeringTypePtrOutput() ExpressRoutePeeringTypePtrOutput
	ToExpressRoutePeeringTypePtrOutputWithContext(context.Context) ExpressRoutePeeringTypePtrOutput
}

type expressRoutePeeringTypePtr string

func ExpressRoutePeeringTypePtr(v string) ExpressRoutePeeringTypePtrInput {
	return (*expressRoutePeeringTypePtr)(&v)
}

func (*expressRoutePeeringTypePtr) ElementType() reflect.Type {
	return expressRoutePeeringTypePtrType
}

func (in *expressRoutePeeringTypePtr) ToExpressRoutePeeringTypePtrOutput() ExpressRoutePeeringTypePtrOutput {
	return pulumi.ToOutput(in).(ExpressRoutePeeringTypePtrOutput)
}

func (in *expressRoutePeeringTypePtr) ToExpressRoutePeeringTypePtrOutputWithContext(ctx context.Context) ExpressRoutePeeringTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRoutePeeringTypePtrOutput)
}

type ExpressRoutePortsEncapsulation string

const (
	ExpressRoutePortsEncapsulationDot1Q = ExpressRoutePortsEncapsulation("Dot1Q")
	ExpressRoutePortsEncapsulationQinQ  = ExpressRoutePortsEncapsulation("QinQ")
)

func (ExpressRoutePortsEncapsulation) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRoutePortsEncapsulation)(nil)).Elem()
}

func (e ExpressRoutePortsEncapsulation) ToExpressRoutePortsEncapsulationOutput() ExpressRoutePortsEncapsulationOutput {
	return pulumi.ToOutput(e).(ExpressRoutePortsEncapsulationOutput)
}

func (e ExpressRoutePortsEncapsulation) ToExpressRoutePortsEncapsulationOutputWithContext(ctx context.Context) ExpressRoutePortsEncapsulationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressRoutePortsEncapsulationOutput)
}

func (e ExpressRoutePortsEncapsulation) ToExpressRoutePortsEncapsulationPtrOutput() ExpressRoutePortsEncapsulationPtrOutput {
	return e.ToExpressRoutePortsEncapsulationPtrOutputWithContext(context.Background())
}

func (e ExpressRoutePortsEncapsulation) ToExpressRoutePortsEncapsulationPtrOutputWithContext(ctx context.Context) ExpressRoutePortsEncapsulationPtrOutput {
	return ExpressRoutePortsEncapsulation(e).ToExpressRoutePortsEncapsulationOutputWithContext(ctx).ToExpressRoutePortsEncapsulationPtrOutputWithContext(ctx)
}

func (e ExpressRoutePortsEncapsulation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRoutePortsEncapsulation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressRoutePortsEncapsulation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressRoutePortsEncapsulation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressRoutePortsEncapsulationOutput struct{ *pulumi.OutputState }

func (ExpressRoutePortsEncapsulationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRoutePortsEncapsulation)(nil)).Elem()
}

func (o ExpressRoutePortsEncapsulationOutput) ToExpressRoutePortsEncapsulationOutput() ExpressRoutePortsEncapsulationOutput {
	return o
}

func (o ExpressRoutePortsEncapsulationOutput) ToExpressRoutePortsEncapsulationOutputWithContext(ctx context.Context) ExpressRoutePortsEncapsulationOutput {
	return o
}

func (o ExpressRoutePortsEncapsulationOutput) ToExpressRoutePortsEncapsulationPtrOutput() ExpressRoutePortsEncapsulationPtrOutput {
	return o.ToExpressRoutePortsEncapsulationPtrOutputWithContext(context.Background())
}

func (o ExpressRoutePortsEncapsulationOutput) ToExpressRoutePortsEncapsulationPtrOutputWithContext(ctx context.Context) ExpressRoutePortsEncapsulationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRoutePortsEncapsulation) *ExpressRoutePortsEncapsulation {
		return &v
	}).(ExpressRoutePortsEncapsulationPtrOutput)
}

func (o ExpressRoutePortsEncapsulationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressRoutePortsEncapsulationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRoutePortsEncapsulation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressRoutePortsEncapsulationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRoutePortsEncapsulationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressRoutePortsEncapsulation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressRoutePortsEncapsulationPtrOutput struct{ *pulumi.OutputState }

func (ExpressRoutePortsEncapsulationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePortsEncapsulation)(nil)).Elem()
}

func (o ExpressRoutePortsEncapsulationPtrOutput) ToExpressRoutePortsEncapsulationPtrOutput() ExpressRoutePortsEncapsulationPtrOutput {
	return o
}

func (o ExpressRoutePortsEncapsulationPtrOutput) ToExpressRoutePortsEncapsulationPtrOutputWithContext(ctx context.Context) ExpressRoutePortsEncapsulationPtrOutput {
	return o
}

func (o ExpressRoutePortsEncapsulationPtrOutput) Elem() ExpressRoutePortsEncapsulationOutput {
	return o.ApplyT(func(v *ExpressRoutePortsEncapsulation) ExpressRoutePortsEncapsulation {
		if v != nil {
			return *v
		}
		var ret ExpressRoutePortsEncapsulation
		return ret
	}).(ExpressRoutePortsEncapsulationOutput)
}

func (o ExpressRoutePortsEncapsulationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressRoutePortsEncapsulationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressRoutePortsEncapsulation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressRoutePortsEncapsulationInput interface {
	pulumi.Input

	ToExpressRoutePortsEncapsulationOutput() ExpressRoutePortsEncapsulationOutput
	ToExpressRoutePortsEncapsulationOutputWithContext(context.Context) ExpressRoutePortsEncapsulationOutput
}

var expressRoutePortsEncapsulationPtrType = reflect.TypeOf((**ExpressRoutePortsEncapsulation)(nil)).Elem()

type ExpressRoutePortsEncapsulationPtrInput interface {
	pulumi.Input

	ToExpressRoutePortsEncapsulationPtrOutput() ExpressRoutePortsEncapsulationPtrOutput
	ToExpressRoutePortsEncapsulationPtrOutputWithContext(context.Context) ExpressRoutePortsEncapsulationPtrOutput
}

type expressRoutePortsEncapsulationPtr string

func ExpressRoutePortsEncapsulationPtr(v string) ExpressRoutePortsEncapsulationPtrInput {
	return (*expressRoutePortsEncapsulationPtr)(&v)
}

func (*expressRoutePortsEncapsulationPtr) ElementType() reflect.Type {
	return expressRoutePortsEncapsulationPtrType
}

func (in *expressRoutePortsEncapsulationPtr) ToExpressRoutePortsEncapsulationPtrOutput() ExpressRoutePortsEncapsulationPtrOutput {
	return pulumi.ToOutput(in).(ExpressRoutePortsEncapsulationPtrOutput)
}

func (in *expressRoutePortsEncapsulationPtr) ToExpressRoutePortsEncapsulationPtrOutputWithContext(ctx context.Context) ExpressRoutePortsEncapsulationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressRoutePortsEncapsulationPtrOutput)
}

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesEdgeZone = ExtendedLocationTypes("EdgeZone")
)

func (ExtendedLocationTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationTypes)(nil)).Elem()
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput {
	return pulumi.ToOutput(e).(ExtendedLocationTypesOutput)
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesOutputWithContext(ctx context.Context) ExtendedLocationTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExtendedLocationTypesOutput)
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return e.ToExtendedLocationTypesPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationTypes) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return ExtendedLocationTypes(e).ToExtendedLocationTypesOutputWithContext(ctx).ToExtendedLocationTypesPtrOutputWithContext(ctx)
}

func (e ExtendedLocationTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExtendedLocationTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExtendedLocationTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExtendedLocationTypesOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationTypes)(nil)).Elem()
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput {
	return o
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesOutputWithContext(ctx context.Context) ExtendedLocationTypesOutput {
	return o
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return o.ToExtendedLocationTypesPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocationTypes) *ExtendedLocationTypes {
		return &v
	}).(ExtendedLocationTypesPtrOutput)
}

func (o ExtendedLocationTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExtendedLocationTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExtendedLocationTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationTypesPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationTypes)(nil)).Elem()
}

func (o ExtendedLocationTypesPtrOutput) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return o
}

func (o ExtendedLocationTypesPtrOutput) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return o
}

func (o ExtendedLocationTypesPtrOutput) Elem() ExtendedLocationTypesOutput {
	return o.ApplyT(func(v *ExtendedLocationTypes) ExtendedLocationTypes {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationTypes
		return ret
	}).(ExtendedLocationTypesOutput)
}

func (o ExtendedLocationTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExtendedLocationTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExtendedLocationTypesInput interface {
	pulumi.Input

	ToExtendedLocationTypesOutput() ExtendedLocationTypesOutput
	ToExtendedLocationTypesOutputWithContext(context.Context) ExtendedLocationTypesOutput
}

var extendedLocationTypesPtrType = reflect.TypeOf((**ExtendedLocationTypes)(nil)).Elem()

type ExtendedLocationTypesPtrInput interface {
	pulumi.Input

	ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput
	ToExtendedLocationTypesPtrOutputWithContext(context.Context) ExtendedLocationTypesPtrOutput
}

type extendedLocationTypesPtr string

func ExtendedLocationTypesPtr(v string) ExtendedLocationTypesPtrInput {
	return (*extendedLocationTypesPtr)(&v)
}

func (*extendedLocationTypesPtr) ElementType() reflect.Type {
	return extendedLocationTypesPtrType
}

func (in *extendedLocationTypesPtr) ToExtendedLocationTypesPtrOutput() ExtendedLocationTypesPtrOutput {
	return pulumi.ToOutput(in).(ExtendedLocationTypesPtrOutput)
}

func (in *extendedLocationTypesPtr) ToExtendedLocationTypesPtrOutputWithContext(ctx context.Context) ExtendedLocationTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExtendedLocationTypesPtrOutput)
}

type FirewallPolicyFilterRuleCollectionActionType string

const (
	FirewallPolicyFilterRuleCollectionActionTypeAllow = FirewallPolicyFilterRuleCollectionActionType("Allow")
	FirewallPolicyFilterRuleCollectionActionTypeDeny  = FirewallPolicyFilterRuleCollectionActionType("Deny")
)

func (FirewallPolicyFilterRuleCollectionActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyFilterRuleCollectionActionType)(nil)).Elem()
}

func (e FirewallPolicyFilterRuleCollectionActionType) ToFirewallPolicyFilterRuleCollectionActionTypeOutput() FirewallPolicyFilterRuleCollectionActionTypeOutput {
	return pulumi.ToOutput(e).(FirewallPolicyFilterRuleCollectionActionTypeOutput)
}

func (e FirewallPolicyFilterRuleCollectionActionType) ToFirewallPolicyFilterRuleCollectionActionTypeOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallPolicyFilterRuleCollectionActionTypeOutput)
}

func (e FirewallPolicyFilterRuleCollectionActionType) ToFirewallPolicyFilterRuleCollectionActionTypePtrOutput() FirewallPolicyFilterRuleCollectionActionTypePtrOutput {
	return e.ToFirewallPolicyFilterRuleCollectionActionTypePtrOutputWithContext(context.Background())
}

func (e FirewallPolicyFilterRuleCollectionActionType) ToFirewallPolicyFilterRuleCollectionActionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionTypePtrOutput {
	return FirewallPolicyFilterRuleCollectionActionType(e).ToFirewallPolicyFilterRuleCollectionActionTypeOutputWithContext(ctx).ToFirewallPolicyFilterRuleCollectionActionTypePtrOutputWithContext(ctx)
}

func (e FirewallPolicyFilterRuleCollectionActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyFilterRuleCollectionActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyFilterRuleCollectionActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyFilterRuleCollectionActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallPolicyFilterRuleCollectionActionTypeOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleCollectionActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyFilterRuleCollectionActionType)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleCollectionActionTypeOutput) ToFirewallPolicyFilterRuleCollectionActionTypeOutput() FirewallPolicyFilterRuleCollectionActionTypeOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionActionTypeOutput) ToFirewallPolicyFilterRuleCollectionActionTypeOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionTypeOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionActionTypeOutput) ToFirewallPolicyFilterRuleCollectionActionTypePtrOutput() FirewallPolicyFilterRuleCollectionActionTypePtrOutput {
	return o.ToFirewallPolicyFilterRuleCollectionActionTypePtrOutputWithContext(context.Background())
}

func (o FirewallPolicyFilterRuleCollectionActionTypeOutput) ToFirewallPolicyFilterRuleCollectionActionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicyFilterRuleCollectionActionType) *FirewallPolicyFilterRuleCollectionActionType {
		return &v
	}).(FirewallPolicyFilterRuleCollectionActionTypePtrOutput)
}

func (o FirewallPolicyFilterRuleCollectionActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallPolicyFilterRuleCollectionActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyFilterRuleCollectionActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallPolicyFilterRuleCollectionActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyFilterRuleCollectionActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyFilterRuleCollectionActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallPolicyFilterRuleCollectionActionTypePtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyFilterRuleCollectionActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyFilterRuleCollectionActionType)(nil)).Elem()
}

func (o FirewallPolicyFilterRuleCollectionActionTypePtrOutput) ToFirewallPolicyFilterRuleCollectionActionTypePtrOutput() FirewallPolicyFilterRuleCollectionActionTypePtrOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionActionTypePtrOutput) ToFirewallPolicyFilterRuleCollectionActionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionTypePtrOutput {
	return o
}

func (o FirewallPolicyFilterRuleCollectionActionTypePtrOutput) Elem() FirewallPolicyFilterRuleCollectionActionTypeOutput {
	return o.ApplyT(func(v *FirewallPolicyFilterRuleCollectionActionType) FirewallPolicyFilterRuleCollectionActionType {
		if v != nil {
			return *v
		}
		var ret FirewallPolicyFilterRuleCollectionActionType
		return ret
	}).(FirewallPolicyFilterRuleCollectionActionTypeOutput)
}

func (o FirewallPolicyFilterRuleCollectionActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyFilterRuleCollectionActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallPolicyFilterRuleCollectionActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirewallPolicyFilterRuleCollectionActionTypeInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleCollectionActionTypeOutput() FirewallPolicyFilterRuleCollectionActionTypeOutput
	ToFirewallPolicyFilterRuleCollectionActionTypeOutputWithContext(context.Context) FirewallPolicyFilterRuleCollectionActionTypeOutput
}

var firewallPolicyFilterRuleCollectionActionTypePtrType = reflect.TypeOf((**FirewallPolicyFilterRuleCollectionActionType)(nil)).Elem()

type FirewallPolicyFilterRuleCollectionActionTypePtrInput interface {
	pulumi.Input

	ToFirewallPolicyFilterRuleCollectionActionTypePtrOutput() FirewallPolicyFilterRuleCollectionActionTypePtrOutput
	ToFirewallPolicyFilterRuleCollectionActionTypePtrOutputWithContext(context.Context) FirewallPolicyFilterRuleCollectionActionTypePtrOutput
}

type firewallPolicyFilterRuleCollectionActionTypePtr string

func FirewallPolicyFilterRuleCollectionActionTypePtr(v string) FirewallPolicyFilterRuleCollectionActionTypePtrInput {
	return (*firewallPolicyFilterRuleCollectionActionTypePtr)(&v)
}

func (*firewallPolicyFilterRuleCollectionActionTypePtr) ElementType() reflect.Type {
	return firewallPolicyFilterRuleCollectionActionTypePtrType
}

func (in *firewallPolicyFilterRuleCollectionActionTypePtr) ToFirewallPolicyFilterRuleCollectionActionTypePtrOutput() FirewallPolicyFilterRuleCollectionActionTypePtrOutput {
	return pulumi.ToOutput(in).(FirewallPolicyFilterRuleCollectionActionTypePtrOutput)
}

func (in *firewallPolicyFilterRuleCollectionActionTypePtr) ToFirewallPolicyFilterRuleCollectionActionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyFilterRuleCollectionActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallPolicyFilterRuleCollectionActionTypePtrOutput)
}

type FirewallPolicyIntrusionDetectionProtocol string

const (
	FirewallPolicyIntrusionDetectionProtocolTCP  = FirewallPolicyIntrusionDetectionProtocol("TCP")
	FirewallPolicyIntrusionDetectionProtocolUDP  = FirewallPolicyIntrusionDetectionProtocol("UDP")
	FirewallPolicyIntrusionDetectionProtocolICMP = FirewallPolicyIntrusionDetectionProtocol("ICMP")
	FirewallPolicyIntrusionDetectionProtocolANY  = FirewallPolicyIntrusionDetectionProtocol("ANY")
)

func (FirewallPolicyIntrusionDetectionProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyIntrusionDetectionProtocol)(nil)).Elem()
}

func (e FirewallPolicyIntrusionDetectionProtocol) ToFirewallPolicyIntrusionDetectionProtocolOutput() FirewallPolicyIntrusionDetectionProtocolOutput {
	return pulumi.ToOutput(e).(FirewallPolicyIntrusionDetectionProtocolOutput)
}

func (e FirewallPolicyIntrusionDetectionProtocol) ToFirewallPolicyIntrusionDetectionProtocolOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallPolicyIntrusionDetectionProtocolOutput)
}

func (e FirewallPolicyIntrusionDetectionProtocol) ToFirewallPolicyIntrusionDetectionProtocolPtrOutput() FirewallPolicyIntrusionDetectionProtocolPtrOutput {
	return e.ToFirewallPolicyIntrusionDetectionProtocolPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyIntrusionDetectionProtocol) ToFirewallPolicyIntrusionDetectionProtocolPtrOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionProtocolPtrOutput {
	return FirewallPolicyIntrusionDetectionProtocol(e).ToFirewallPolicyIntrusionDetectionProtocolOutputWithContext(ctx).ToFirewallPolicyIntrusionDetectionProtocolPtrOutputWithContext(ctx)
}

func (e FirewallPolicyIntrusionDetectionProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyIntrusionDetectionProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyIntrusionDetectionProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyIntrusionDetectionProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallPolicyIntrusionDetectionProtocolOutput struct{ *pulumi.OutputState }

func (FirewallPolicyIntrusionDetectionProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyIntrusionDetectionProtocol)(nil)).Elem()
}

func (o FirewallPolicyIntrusionDetectionProtocolOutput) ToFirewallPolicyIntrusionDetectionProtocolOutput() FirewallPolicyIntrusionDetectionProtocolOutput {
	return o
}

func (o FirewallPolicyIntrusionDetectionProtocolOutput) ToFirewallPolicyIntrusionDetectionProtocolOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionProtocolOutput {
	return o
}

func (o FirewallPolicyIntrusionDetectionProtocolOutput) ToFirewallPolicyIntrusionDetectionProtocolPtrOutput() FirewallPolicyIntrusionDetectionProtocolPtrOutput {
	return o.ToFirewallPolicyIntrusionDetectionProtocolPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyIntrusionDetectionProtocolOutput) ToFirewallPolicyIntrusionDetectionProtocolPtrOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicyIntrusionDetectionProtocol) *FirewallPolicyIntrusionDetectionProtocol {
		return &v
	}).(FirewallPolicyIntrusionDetectionProtocolPtrOutput)
}

func (o FirewallPolicyIntrusionDetectionProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallPolicyIntrusionDetectionProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyIntrusionDetectionProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallPolicyIntrusionDetectionProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyIntrusionDetectionProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyIntrusionDetectionProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallPolicyIntrusionDetectionProtocolPtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyIntrusionDetectionProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyIntrusionDetectionProtocol)(nil)).Elem()
}

func (o FirewallPolicyIntrusionDetectionProtocolPtrOutput) ToFirewallPolicyIntrusionDetectionProtocolPtrOutput() FirewallPolicyIntrusionDetectionProtocolPtrOutput {
	return o
}

func (o FirewallPolicyIntrusionDetectionProtocolPtrOutput) ToFirewallPolicyIntrusionDetectionProtocolPtrOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionProtocolPtrOutput {
	return o
}

func (o FirewallPolicyIntrusionDetectionProtocolPtrOutput) Elem() FirewallPolicyIntrusionDetectionProtocolOutput {
	return o.ApplyT(func(v *FirewallPolicyIntrusionDetectionProtocol) FirewallPolicyIntrusionDetectionProtocol {
		if v != nil {
			return *v
		}
		var ret FirewallPolicyIntrusionDetectionProtocol
		return ret
	}).(FirewallPolicyIntrusionDetectionProtocolOutput)
}

func (o FirewallPolicyIntrusionDetectionProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyIntrusionDetectionProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallPolicyIntrusionDetectionProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirewallPolicyIntrusionDetectionProtocolInput interface {
	pulumi.Input

	ToFirewallPolicyIntrusionDetectionProtocolOutput() FirewallPolicyIntrusionDetectionProtocolOutput
	ToFirewallPolicyIntrusionDetectionProtocolOutputWithContext(context.Context) FirewallPolicyIntrusionDetectionProtocolOutput
}

var firewallPolicyIntrusionDetectionProtocolPtrType = reflect.TypeOf((**FirewallPolicyIntrusionDetectionProtocol)(nil)).Elem()

type FirewallPolicyIntrusionDetectionProtocolPtrInput interface {
	pulumi.Input

	ToFirewallPolicyIntrusionDetectionProtocolPtrOutput() FirewallPolicyIntrusionDetectionProtocolPtrOutput
	ToFirewallPolicyIntrusionDetectionProtocolPtrOutputWithContext(context.Context) FirewallPolicyIntrusionDetectionProtocolPtrOutput
}

type firewallPolicyIntrusionDetectionProtocolPtr string

func FirewallPolicyIntrusionDetectionProtocolPtr(v string) FirewallPolicyIntrusionDetectionProtocolPtrInput {
	return (*firewallPolicyIntrusionDetectionProtocolPtr)(&v)
}

func (*firewallPolicyIntrusionDetectionProtocolPtr) ElementType() reflect.Type {
	return firewallPolicyIntrusionDetectionProtocolPtrType
}

func (in *firewallPolicyIntrusionDetectionProtocolPtr) ToFirewallPolicyIntrusionDetectionProtocolPtrOutput() FirewallPolicyIntrusionDetectionProtocolPtrOutput {
	return pulumi.ToOutput(in).(FirewallPolicyIntrusionDetectionProtocolPtrOutput)
}

func (in *firewallPolicyIntrusionDetectionProtocolPtr) ToFirewallPolicyIntrusionDetectionProtocolPtrOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallPolicyIntrusionDetectionProtocolPtrOutput)
}

type FirewallPolicyIntrusionDetectionStateType string

const (
	FirewallPolicyIntrusionDetectionStateTypeOff   = FirewallPolicyIntrusionDetectionStateType("Off")
	FirewallPolicyIntrusionDetectionStateTypeAlert = FirewallPolicyIntrusionDetectionStateType("Alert")
	FirewallPolicyIntrusionDetectionStateTypeDeny  = FirewallPolicyIntrusionDetectionStateType("Deny")
)

func (FirewallPolicyIntrusionDetectionStateType) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyIntrusionDetectionStateType)(nil)).Elem()
}

func (e FirewallPolicyIntrusionDetectionStateType) ToFirewallPolicyIntrusionDetectionStateTypeOutput() FirewallPolicyIntrusionDetectionStateTypeOutput {
	return pulumi.ToOutput(e).(FirewallPolicyIntrusionDetectionStateTypeOutput)
}

func (e FirewallPolicyIntrusionDetectionStateType) ToFirewallPolicyIntrusionDetectionStateTypeOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionStateTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallPolicyIntrusionDetectionStateTypeOutput)
}

func (e FirewallPolicyIntrusionDetectionStateType) ToFirewallPolicyIntrusionDetectionStateTypePtrOutput() FirewallPolicyIntrusionDetectionStateTypePtrOutput {
	return e.ToFirewallPolicyIntrusionDetectionStateTypePtrOutputWithContext(context.Background())
}

func (e FirewallPolicyIntrusionDetectionStateType) ToFirewallPolicyIntrusionDetectionStateTypePtrOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionStateTypePtrOutput {
	return FirewallPolicyIntrusionDetectionStateType(e).ToFirewallPolicyIntrusionDetectionStateTypeOutputWithContext(ctx).ToFirewallPolicyIntrusionDetectionStateTypePtrOutputWithContext(ctx)
}

func (e FirewallPolicyIntrusionDetectionStateType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyIntrusionDetectionStateType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyIntrusionDetectionStateType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyIntrusionDetectionStateType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallPolicyIntrusionDetectionStateTypeOutput struct{ *pulumi.OutputState }

func (FirewallPolicyIntrusionDetectionStateTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyIntrusionDetectionStateType)(nil)).Elem()
}

func (o FirewallPolicyIntrusionDetectionStateTypeOutput) ToFirewallPolicyIntrusionDetectionStateTypeOutput() FirewallPolicyIntrusionDetectionStateTypeOutput {
	return o
}

func (o FirewallPolicyIntrusionDetectionStateTypeOutput) ToFirewallPolicyIntrusionDetectionStateTypeOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionStateTypeOutput {
	return o
}

func (o FirewallPolicyIntrusionDetectionStateTypeOutput) ToFirewallPolicyIntrusionDetectionStateTypePtrOutput() FirewallPolicyIntrusionDetectionStateTypePtrOutput {
	return o.ToFirewallPolicyIntrusionDetectionStateTypePtrOutputWithContext(context.Background())
}

func (o FirewallPolicyIntrusionDetectionStateTypeOutput) ToFirewallPolicyIntrusionDetectionStateTypePtrOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionStateTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicyIntrusionDetectionStateType) *FirewallPolicyIntrusionDetectionStateType {
		return &v
	}).(FirewallPolicyIntrusionDetectionStateTypePtrOutput)
}

func (o FirewallPolicyIntrusionDetectionStateTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallPolicyIntrusionDetectionStateTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyIntrusionDetectionStateType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallPolicyIntrusionDetectionStateTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyIntrusionDetectionStateTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyIntrusionDetectionStateType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallPolicyIntrusionDetectionStateTypePtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyIntrusionDetectionStateTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyIntrusionDetectionStateType)(nil)).Elem()
}

func (o FirewallPolicyIntrusionDetectionStateTypePtrOutput) ToFirewallPolicyIntrusionDetectionStateTypePtrOutput() FirewallPolicyIntrusionDetectionStateTypePtrOutput {
	return o
}

func (o FirewallPolicyIntrusionDetectionStateTypePtrOutput) ToFirewallPolicyIntrusionDetectionStateTypePtrOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionStateTypePtrOutput {
	return o
}

func (o FirewallPolicyIntrusionDetectionStateTypePtrOutput) Elem() FirewallPolicyIntrusionDetectionStateTypeOutput {
	return o.ApplyT(func(v *FirewallPolicyIntrusionDetectionStateType) FirewallPolicyIntrusionDetectionStateType {
		if v != nil {
			return *v
		}
		var ret FirewallPolicyIntrusionDetectionStateType
		return ret
	}).(FirewallPolicyIntrusionDetectionStateTypeOutput)
}

func (o FirewallPolicyIntrusionDetectionStateTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyIntrusionDetectionStateTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallPolicyIntrusionDetectionStateType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirewallPolicyIntrusionDetectionStateTypeInput interface {
	pulumi.Input

	ToFirewallPolicyIntrusionDetectionStateTypeOutput() FirewallPolicyIntrusionDetectionStateTypeOutput
	ToFirewallPolicyIntrusionDetectionStateTypeOutputWithContext(context.Context) FirewallPolicyIntrusionDetectionStateTypeOutput
}

var firewallPolicyIntrusionDetectionStateTypePtrType = reflect.TypeOf((**FirewallPolicyIntrusionDetectionStateType)(nil)).Elem()

type FirewallPolicyIntrusionDetectionStateTypePtrInput interface {
	pulumi.Input

	ToFirewallPolicyIntrusionDetectionStateTypePtrOutput() FirewallPolicyIntrusionDetectionStateTypePtrOutput
	ToFirewallPolicyIntrusionDetectionStateTypePtrOutputWithContext(context.Context) FirewallPolicyIntrusionDetectionStateTypePtrOutput
}

type firewallPolicyIntrusionDetectionStateTypePtr string

func FirewallPolicyIntrusionDetectionStateTypePtr(v string) FirewallPolicyIntrusionDetectionStateTypePtrInput {
	return (*firewallPolicyIntrusionDetectionStateTypePtr)(&v)
}

func (*firewallPolicyIntrusionDetectionStateTypePtr) ElementType() reflect.Type {
	return firewallPolicyIntrusionDetectionStateTypePtrType
}

func (in *firewallPolicyIntrusionDetectionStateTypePtr) ToFirewallPolicyIntrusionDetectionStateTypePtrOutput() FirewallPolicyIntrusionDetectionStateTypePtrOutput {
	return pulumi.ToOutput(in).(FirewallPolicyIntrusionDetectionStateTypePtrOutput)
}

func (in *firewallPolicyIntrusionDetectionStateTypePtr) ToFirewallPolicyIntrusionDetectionStateTypePtrOutputWithContext(ctx context.Context) FirewallPolicyIntrusionDetectionStateTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallPolicyIntrusionDetectionStateTypePtrOutput)
}

type FirewallPolicyNatRuleCollectionActionType string

const (
	FirewallPolicyNatRuleCollectionActionTypeDNAT = FirewallPolicyNatRuleCollectionActionType("DNAT")
)

func (FirewallPolicyNatRuleCollectionActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyNatRuleCollectionActionType)(nil)).Elem()
}

func (e FirewallPolicyNatRuleCollectionActionType) ToFirewallPolicyNatRuleCollectionActionTypeOutput() FirewallPolicyNatRuleCollectionActionTypeOutput {
	return pulumi.ToOutput(e).(FirewallPolicyNatRuleCollectionActionTypeOutput)
}

func (e FirewallPolicyNatRuleCollectionActionType) ToFirewallPolicyNatRuleCollectionActionTypeOutputWithContext(ctx context.Context) FirewallPolicyNatRuleCollectionActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallPolicyNatRuleCollectionActionTypeOutput)
}

func (e FirewallPolicyNatRuleCollectionActionType) ToFirewallPolicyNatRuleCollectionActionTypePtrOutput() FirewallPolicyNatRuleCollectionActionTypePtrOutput {
	return e.ToFirewallPolicyNatRuleCollectionActionTypePtrOutputWithContext(context.Background())
}

func (e FirewallPolicyNatRuleCollectionActionType) ToFirewallPolicyNatRuleCollectionActionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyNatRuleCollectionActionTypePtrOutput {
	return FirewallPolicyNatRuleCollectionActionType(e).ToFirewallPolicyNatRuleCollectionActionTypeOutputWithContext(ctx).ToFirewallPolicyNatRuleCollectionActionTypePtrOutputWithContext(ctx)
}

func (e FirewallPolicyNatRuleCollectionActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyNatRuleCollectionActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyNatRuleCollectionActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyNatRuleCollectionActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallPolicyNatRuleCollectionActionTypeOutput struct{ *pulumi.OutputState }

func (FirewallPolicyNatRuleCollectionActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyNatRuleCollectionActionType)(nil)).Elem()
}

func (o FirewallPolicyNatRuleCollectionActionTypeOutput) ToFirewallPolicyNatRuleCollectionActionTypeOutput() FirewallPolicyNatRuleCollectionActionTypeOutput {
	return o
}

func (o FirewallPolicyNatRuleCollectionActionTypeOutput) ToFirewallPolicyNatRuleCollectionActionTypeOutputWithContext(ctx context.Context) FirewallPolicyNatRuleCollectionActionTypeOutput {
	return o
}

func (o FirewallPolicyNatRuleCollectionActionTypeOutput) ToFirewallPolicyNatRuleCollectionActionTypePtrOutput() FirewallPolicyNatRuleCollectionActionTypePtrOutput {
	return o.ToFirewallPolicyNatRuleCollectionActionTypePtrOutputWithContext(context.Background())
}

func (o FirewallPolicyNatRuleCollectionActionTypeOutput) ToFirewallPolicyNatRuleCollectionActionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyNatRuleCollectionActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicyNatRuleCollectionActionType) *FirewallPolicyNatRuleCollectionActionType {
		return &v
	}).(FirewallPolicyNatRuleCollectionActionTypePtrOutput)
}

func (o FirewallPolicyNatRuleCollectionActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallPolicyNatRuleCollectionActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyNatRuleCollectionActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallPolicyNatRuleCollectionActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyNatRuleCollectionActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyNatRuleCollectionActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallPolicyNatRuleCollectionActionTypePtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyNatRuleCollectionActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyNatRuleCollectionActionType)(nil)).Elem()
}

func (o FirewallPolicyNatRuleCollectionActionTypePtrOutput) ToFirewallPolicyNatRuleCollectionActionTypePtrOutput() FirewallPolicyNatRuleCollectionActionTypePtrOutput {
	return o
}

func (o FirewallPolicyNatRuleCollectionActionTypePtrOutput) ToFirewallPolicyNatRuleCollectionActionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyNatRuleCollectionActionTypePtrOutput {
	return o
}

func (o FirewallPolicyNatRuleCollectionActionTypePtrOutput) Elem() FirewallPolicyNatRuleCollectionActionTypeOutput {
	return o.ApplyT(func(v *FirewallPolicyNatRuleCollectionActionType) FirewallPolicyNatRuleCollectionActionType {
		if v != nil {
			return *v
		}
		var ret FirewallPolicyNatRuleCollectionActionType
		return ret
	}).(FirewallPolicyNatRuleCollectionActionTypeOutput)
}

func (o FirewallPolicyNatRuleCollectionActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyNatRuleCollectionActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallPolicyNatRuleCollectionActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirewallPolicyNatRuleCollectionActionTypeInput interface {
	pulumi.Input

	ToFirewallPolicyNatRuleCollectionActionTypeOutput() FirewallPolicyNatRuleCollectionActionTypeOutput
	ToFirewallPolicyNatRuleCollectionActionTypeOutputWithContext(context.Context) FirewallPolicyNatRuleCollectionActionTypeOutput
}

var firewallPolicyNatRuleCollectionActionTypePtrType = reflect.TypeOf((**FirewallPolicyNatRuleCollectionActionType)(nil)).Elem()

type FirewallPolicyNatRuleCollectionActionTypePtrInput interface {
	pulumi.Input

	ToFirewallPolicyNatRuleCollectionActionTypePtrOutput() FirewallPolicyNatRuleCollectionActionTypePtrOutput
	ToFirewallPolicyNatRuleCollectionActionTypePtrOutputWithContext(context.Context) FirewallPolicyNatRuleCollectionActionTypePtrOutput
}

type firewallPolicyNatRuleCollectionActionTypePtr string

func FirewallPolicyNatRuleCollectionActionTypePtr(v string) FirewallPolicyNatRuleCollectionActionTypePtrInput {
	return (*firewallPolicyNatRuleCollectionActionTypePtr)(&v)
}

func (*firewallPolicyNatRuleCollectionActionTypePtr) ElementType() reflect.Type {
	return firewallPolicyNatRuleCollectionActionTypePtrType
}

func (in *firewallPolicyNatRuleCollectionActionTypePtr) ToFirewallPolicyNatRuleCollectionActionTypePtrOutput() FirewallPolicyNatRuleCollectionActionTypePtrOutput {
	return pulumi.ToOutput(in).(FirewallPolicyNatRuleCollectionActionTypePtrOutput)
}

func (in *firewallPolicyNatRuleCollectionActionTypePtr) ToFirewallPolicyNatRuleCollectionActionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyNatRuleCollectionActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallPolicyNatRuleCollectionActionTypePtrOutput)
}

type FirewallPolicyRuleApplicationProtocolType string

const (
	FirewallPolicyRuleApplicationProtocolTypeHttp  = FirewallPolicyRuleApplicationProtocolType("Http")
	FirewallPolicyRuleApplicationProtocolTypeHttps = FirewallPolicyRuleApplicationProtocolType("Https")
)

func (FirewallPolicyRuleApplicationProtocolType) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleApplicationProtocolType)(nil)).Elem()
}

func (e FirewallPolicyRuleApplicationProtocolType) ToFirewallPolicyRuleApplicationProtocolTypeOutput() FirewallPolicyRuleApplicationProtocolTypeOutput {
	return pulumi.ToOutput(e).(FirewallPolicyRuleApplicationProtocolTypeOutput)
}

func (e FirewallPolicyRuleApplicationProtocolType) ToFirewallPolicyRuleApplicationProtocolTypeOutputWithContext(ctx context.Context) FirewallPolicyRuleApplicationProtocolTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallPolicyRuleApplicationProtocolTypeOutput)
}

func (e FirewallPolicyRuleApplicationProtocolType) ToFirewallPolicyRuleApplicationProtocolTypePtrOutput() FirewallPolicyRuleApplicationProtocolTypePtrOutput {
	return e.ToFirewallPolicyRuleApplicationProtocolTypePtrOutputWithContext(context.Background())
}

func (e FirewallPolicyRuleApplicationProtocolType) ToFirewallPolicyRuleApplicationProtocolTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleApplicationProtocolTypePtrOutput {
	return FirewallPolicyRuleApplicationProtocolType(e).ToFirewallPolicyRuleApplicationProtocolTypeOutputWithContext(ctx).ToFirewallPolicyRuleApplicationProtocolTypePtrOutputWithContext(ctx)
}

func (e FirewallPolicyRuleApplicationProtocolType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyRuleApplicationProtocolType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyRuleApplicationProtocolType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyRuleApplicationProtocolType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallPolicyRuleApplicationProtocolTypeOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleApplicationProtocolTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleApplicationProtocolType)(nil)).Elem()
}

func (o FirewallPolicyRuleApplicationProtocolTypeOutput) ToFirewallPolicyRuleApplicationProtocolTypeOutput() FirewallPolicyRuleApplicationProtocolTypeOutput {
	return o
}

func (o FirewallPolicyRuleApplicationProtocolTypeOutput) ToFirewallPolicyRuleApplicationProtocolTypeOutputWithContext(ctx context.Context) FirewallPolicyRuleApplicationProtocolTypeOutput {
	return o
}

func (o FirewallPolicyRuleApplicationProtocolTypeOutput) ToFirewallPolicyRuleApplicationProtocolTypePtrOutput() FirewallPolicyRuleApplicationProtocolTypePtrOutput {
	return o.ToFirewallPolicyRuleApplicationProtocolTypePtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleApplicationProtocolTypeOutput) ToFirewallPolicyRuleApplicationProtocolTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleApplicationProtocolTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicyRuleApplicationProtocolType) *FirewallPolicyRuleApplicationProtocolType {
		return &v
	}).(FirewallPolicyRuleApplicationProtocolTypePtrOutput)
}

func (o FirewallPolicyRuleApplicationProtocolTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleApplicationProtocolTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyRuleApplicationProtocolType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallPolicyRuleApplicationProtocolTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleApplicationProtocolTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyRuleApplicationProtocolType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallPolicyRuleApplicationProtocolTypePtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleApplicationProtocolTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleApplicationProtocolType)(nil)).Elem()
}

func (o FirewallPolicyRuleApplicationProtocolTypePtrOutput) ToFirewallPolicyRuleApplicationProtocolTypePtrOutput() FirewallPolicyRuleApplicationProtocolTypePtrOutput {
	return o
}

func (o FirewallPolicyRuleApplicationProtocolTypePtrOutput) ToFirewallPolicyRuleApplicationProtocolTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleApplicationProtocolTypePtrOutput {
	return o
}

func (o FirewallPolicyRuleApplicationProtocolTypePtrOutput) Elem() FirewallPolicyRuleApplicationProtocolTypeOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleApplicationProtocolType) FirewallPolicyRuleApplicationProtocolType {
		if v != nil {
			return *v
		}
		var ret FirewallPolicyRuleApplicationProtocolType
		return ret
	}).(FirewallPolicyRuleApplicationProtocolTypeOutput)
}

func (o FirewallPolicyRuleApplicationProtocolTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleApplicationProtocolTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallPolicyRuleApplicationProtocolType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirewallPolicyRuleApplicationProtocolTypeInput interface {
	pulumi.Input

	ToFirewallPolicyRuleApplicationProtocolTypeOutput() FirewallPolicyRuleApplicationProtocolTypeOutput
	ToFirewallPolicyRuleApplicationProtocolTypeOutputWithContext(context.Context) FirewallPolicyRuleApplicationProtocolTypeOutput
}

var firewallPolicyRuleApplicationProtocolTypePtrType = reflect.TypeOf((**FirewallPolicyRuleApplicationProtocolType)(nil)).Elem()

type FirewallPolicyRuleApplicationProtocolTypePtrInput interface {
	pulumi.Input

	ToFirewallPolicyRuleApplicationProtocolTypePtrOutput() FirewallPolicyRuleApplicationProtocolTypePtrOutput
	ToFirewallPolicyRuleApplicationProtocolTypePtrOutputWithContext(context.Context) FirewallPolicyRuleApplicationProtocolTypePtrOutput
}

type firewallPolicyRuleApplicationProtocolTypePtr string

func FirewallPolicyRuleApplicationProtocolTypePtr(v string) FirewallPolicyRuleApplicationProtocolTypePtrInput {
	return (*firewallPolicyRuleApplicationProtocolTypePtr)(&v)
}

func (*firewallPolicyRuleApplicationProtocolTypePtr) ElementType() reflect.Type {
	return firewallPolicyRuleApplicationProtocolTypePtrType
}

func (in *firewallPolicyRuleApplicationProtocolTypePtr) ToFirewallPolicyRuleApplicationProtocolTypePtrOutput() FirewallPolicyRuleApplicationProtocolTypePtrOutput {
	return pulumi.ToOutput(in).(FirewallPolicyRuleApplicationProtocolTypePtrOutput)
}

func (in *firewallPolicyRuleApplicationProtocolTypePtr) ToFirewallPolicyRuleApplicationProtocolTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleApplicationProtocolTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallPolicyRuleApplicationProtocolTypePtrOutput)
}

type FirewallPolicyRuleCollectionType string

const (
	FirewallPolicyRuleCollectionTypeFirewallPolicyNatRuleCollection    = FirewallPolicyRuleCollectionType("FirewallPolicyNatRuleCollection")
	FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection = FirewallPolicyRuleCollectionType("FirewallPolicyFilterRuleCollection")
)

func (FirewallPolicyRuleCollectionType) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleCollectionType)(nil)).Elem()
}

func (e FirewallPolicyRuleCollectionType) ToFirewallPolicyRuleCollectionTypeOutput() FirewallPolicyRuleCollectionTypeOutput {
	return pulumi.ToOutput(e).(FirewallPolicyRuleCollectionTypeOutput)
}

func (e FirewallPolicyRuleCollectionType) ToFirewallPolicyRuleCollectionTypeOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallPolicyRuleCollectionTypeOutput)
}

func (e FirewallPolicyRuleCollectionType) ToFirewallPolicyRuleCollectionTypePtrOutput() FirewallPolicyRuleCollectionTypePtrOutput {
	return e.ToFirewallPolicyRuleCollectionTypePtrOutputWithContext(context.Background())
}

func (e FirewallPolicyRuleCollectionType) ToFirewallPolicyRuleCollectionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionTypePtrOutput {
	return FirewallPolicyRuleCollectionType(e).ToFirewallPolicyRuleCollectionTypeOutputWithContext(ctx).ToFirewallPolicyRuleCollectionTypePtrOutputWithContext(ctx)
}

func (e FirewallPolicyRuleCollectionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyRuleCollectionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyRuleCollectionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyRuleCollectionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallPolicyRuleCollectionTypeOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleCollectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleCollectionType)(nil)).Elem()
}

func (o FirewallPolicyRuleCollectionTypeOutput) ToFirewallPolicyRuleCollectionTypeOutput() FirewallPolicyRuleCollectionTypeOutput {
	return o
}

func (o FirewallPolicyRuleCollectionTypeOutput) ToFirewallPolicyRuleCollectionTypeOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionTypeOutput {
	return o
}

func (o FirewallPolicyRuleCollectionTypeOutput) ToFirewallPolicyRuleCollectionTypePtrOutput() FirewallPolicyRuleCollectionTypePtrOutput {
	return o.ToFirewallPolicyRuleCollectionTypePtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleCollectionTypeOutput) ToFirewallPolicyRuleCollectionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicyRuleCollectionType) *FirewallPolicyRuleCollectionType {
		return &v
	}).(FirewallPolicyRuleCollectionTypePtrOutput)
}

func (o FirewallPolicyRuleCollectionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleCollectionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyRuleCollectionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallPolicyRuleCollectionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleCollectionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyRuleCollectionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallPolicyRuleCollectionTypePtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleCollectionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleCollectionType)(nil)).Elem()
}

func (o FirewallPolicyRuleCollectionTypePtrOutput) ToFirewallPolicyRuleCollectionTypePtrOutput() FirewallPolicyRuleCollectionTypePtrOutput {
	return o
}

func (o FirewallPolicyRuleCollectionTypePtrOutput) ToFirewallPolicyRuleCollectionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionTypePtrOutput {
	return o
}

func (o FirewallPolicyRuleCollectionTypePtrOutput) Elem() FirewallPolicyRuleCollectionTypeOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleCollectionType) FirewallPolicyRuleCollectionType {
		if v != nil {
			return *v
		}
		var ret FirewallPolicyRuleCollectionType
		return ret
	}).(FirewallPolicyRuleCollectionTypeOutput)
}

func (o FirewallPolicyRuleCollectionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleCollectionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallPolicyRuleCollectionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirewallPolicyRuleCollectionTypeInput interface {
	pulumi.Input

	ToFirewallPolicyRuleCollectionTypeOutput() FirewallPolicyRuleCollectionTypeOutput
	ToFirewallPolicyRuleCollectionTypeOutputWithContext(context.Context) FirewallPolicyRuleCollectionTypeOutput
}

var firewallPolicyRuleCollectionTypePtrType = reflect.TypeOf((**FirewallPolicyRuleCollectionType)(nil)).Elem()

type FirewallPolicyRuleCollectionTypePtrInput interface {
	pulumi.Input

	ToFirewallPolicyRuleCollectionTypePtrOutput() FirewallPolicyRuleCollectionTypePtrOutput
	ToFirewallPolicyRuleCollectionTypePtrOutputWithContext(context.Context) FirewallPolicyRuleCollectionTypePtrOutput
}

type firewallPolicyRuleCollectionTypePtr string

func FirewallPolicyRuleCollectionTypePtr(v string) FirewallPolicyRuleCollectionTypePtrInput {
	return (*firewallPolicyRuleCollectionTypePtr)(&v)
}

func (*firewallPolicyRuleCollectionTypePtr) ElementType() reflect.Type {
	return firewallPolicyRuleCollectionTypePtrType
}

func (in *firewallPolicyRuleCollectionTypePtr) ToFirewallPolicyRuleCollectionTypePtrOutput() FirewallPolicyRuleCollectionTypePtrOutput {
	return pulumi.ToOutput(in).(FirewallPolicyRuleCollectionTypePtrOutput)
}

func (in *firewallPolicyRuleCollectionTypePtr) ToFirewallPolicyRuleCollectionTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleCollectionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallPolicyRuleCollectionTypePtrOutput)
}

type FirewallPolicyRuleNetworkProtocol string

const (
	FirewallPolicyRuleNetworkProtocolTCP  = FirewallPolicyRuleNetworkProtocol("TCP")
	FirewallPolicyRuleNetworkProtocolUDP  = FirewallPolicyRuleNetworkProtocol("UDP")
	FirewallPolicyRuleNetworkProtocolAny  = FirewallPolicyRuleNetworkProtocol("Any")
	FirewallPolicyRuleNetworkProtocolICMP = FirewallPolicyRuleNetworkProtocol("ICMP")
)

func (FirewallPolicyRuleNetworkProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleNetworkProtocol)(nil)).Elem()
}

func (e FirewallPolicyRuleNetworkProtocol) ToFirewallPolicyRuleNetworkProtocolOutput() FirewallPolicyRuleNetworkProtocolOutput {
	return pulumi.ToOutput(e).(FirewallPolicyRuleNetworkProtocolOutput)
}

func (e FirewallPolicyRuleNetworkProtocol) ToFirewallPolicyRuleNetworkProtocolOutputWithContext(ctx context.Context) FirewallPolicyRuleNetworkProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallPolicyRuleNetworkProtocolOutput)
}

func (e FirewallPolicyRuleNetworkProtocol) ToFirewallPolicyRuleNetworkProtocolPtrOutput() FirewallPolicyRuleNetworkProtocolPtrOutput {
	return e.ToFirewallPolicyRuleNetworkProtocolPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyRuleNetworkProtocol) ToFirewallPolicyRuleNetworkProtocolPtrOutputWithContext(ctx context.Context) FirewallPolicyRuleNetworkProtocolPtrOutput {
	return FirewallPolicyRuleNetworkProtocol(e).ToFirewallPolicyRuleNetworkProtocolOutputWithContext(ctx).ToFirewallPolicyRuleNetworkProtocolPtrOutputWithContext(ctx)
}

func (e FirewallPolicyRuleNetworkProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyRuleNetworkProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyRuleNetworkProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyRuleNetworkProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallPolicyRuleNetworkProtocolOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleNetworkProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleNetworkProtocol)(nil)).Elem()
}

func (o FirewallPolicyRuleNetworkProtocolOutput) ToFirewallPolicyRuleNetworkProtocolOutput() FirewallPolicyRuleNetworkProtocolOutput {
	return o
}

func (o FirewallPolicyRuleNetworkProtocolOutput) ToFirewallPolicyRuleNetworkProtocolOutputWithContext(ctx context.Context) FirewallPolicyRuleNetworkProtocolOutput {
	return o
}

func (o FirewallPolicyRuleNetworkProtocolOutput) ToFirewallPolicyRuleNetworkProtocolPtrOutput() FirewallPolicyRuleNetworkProtocolPtrOutput {
	return o.ToFirewallPolicyRuleNetworkProtocolPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleNetworkProtocolOutput) ToFirewallPolicyRuleNetworkProtocolPtrOutputWithContext(ctx context.Context) FirewallPolicyRuleNetworkProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicyRuleNetworkProtocol) *FirewallPolicyRuleNetworkProtocol {
		return &v
	}).(FirewallPolicyRuleNetworkProtocolPtrOutput)
}

func (o FirewallPolicyRuleNetworkProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleNetworkProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyRuleNetworkProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallPolicyRuleNetworkProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleNetworkProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyRuleNetworkProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallPolicyRuleNetworkProtocolPtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleNetworkProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleNetworkProtocol)(nil)).Elem()
}

func (o FirewallPolicyRuleNetworkProtocolPtrOutput) ToFirewallPolicyRuleNetworkProtocolPtrOutput() FirewallPolicyRuleNetworkProtocolPtrOutput {
	return o
}

func (o FirewallPolicyRuleNetworkProtocolPtrOutput) ToFirewallPolicyRuleNetworkProtocolPtrOutputWithContext(ctx context.Context) FirewallPolicyRuleNetworkProtocolPtrOutput {
	return o
}

func (o FirewallPolicyRuleNetworkProtocolPtrOutput) Elem() FirewallPolicyRuleNetworkProtocolOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleNetworkProtocol) FirewallPolicyRuleNetworkProtocol {
		if v != nil {
			return *v
		}
		var ret FirewallPolicyRuleNetworkProtocol
		return ret
	}).(FirewallPolicyRuleNetworkProtocolOutput)
}

func (o FirewallPolicyRuleNetworkProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleNetworkProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallPolicyRuleNetworkProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirewallPolicyRuleNetworkProtocolInput interface {
	pulumi.Input

	ToFirewallPolicyRuleNetworkProtocolOutput() FirewallPolicyRuleNetworkProtocolOutput
	ToFirewallPolicyRuleNetworkProtocolOutputWithContext(context.Context) FirewallPolicyRuleNetworkProtocolOutput
}

var firewallPolicyRuleNetworkProtocolPtrType = reflect.TypeOf((**FirewallPolicyRuleNetworkProtocol)(nil)).Elem()

type FirewallPolicyRuleNetworkProtocolPtrInput interface {
	pulumi.Input

	ToFirewallPolicyRuleNetworkProtocolPtrOutput() FirewallPolicyRuleNetworkProtocolPtrOutput
	ToFirewallPolicyRuleNetworkProtocolPtrOutputWithContext(context.Context) FirewallPolicyRuleNetworkProtocolPtrOutput
}

type firewallPolicyRuleNetworkProtocolPtr string

func FirewallPolicyRuleNetworkProtocolPtr(v string) FirewallPolicyRuleNetworkProtocolPtrInput {
	return (*firewallPolicyRuleNetworkProtocolPtr)(&v)
}

func (*firewallPolicyRuleNetworkProtocolPtr) ElementType() reflect.Type {
	return firewallPolicyRuleNetworkProtocolPtrType
}

func (in *firewallPolicyRuleNetworkProtocolPtr) ToFirewallPolicyRuleNetworkProtocolPtrOutput() FirewallPolicyRuleNetworkProtocolPtrOutput {
	return pulumi.ToOutput(in).(FirewallPolicyRuleNetworkProtocolPtrOutput)
}

func (in *firewallPolicyRuleNetworkProtocolPtr) ToFirewallPolicyRuleNetworkProtocolPtrOutputWithContext(ctx context.Context) FirewallPolicyRuleNetworkProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallPolicyRuleNetworkProtocolPtrOutput)
}

type FirewallPolicyRuleType string

const (
	FirewallPolicyRuleTypeApplicationRule = FirewallPolicyRuleType("ApplicationRule")
	FirewallPolicyRuleTypeNetworkRule     = FirewallPolicyRuleType("NetworkRule")
	FirewallPolicyRuleTypeNatRule         = FirewallPolicyRuleType("NatRule")
)

func (FirewallPolicyRuleType) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleType)(nil)).Elem()
}

func (e FirewallPolicyRuleType) ToFirewallPolicyRuleTypeOutput() FirewallPolicyRuleTypeOutput {
	return pulumi.ToOutput(e).(FirewallPolicyRuleTypeOutput)
}

func (e FirewallPolicyRuleType) ToFirewallPolicyRuleTypeOutputWithContext(ctx context.Context) FirewallPolicyRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallPolicyRuleTypeOutput)
}

func (e FirewallPolicyRuleType) ToFirewallPolicyRuleTypePtrOutput() FirewallPolicyRuleTypePtrOutput {
	return e.ToFirewallPolicyRuleTypePtrOutputWithContext(context.Background())
}

func (e FirewallPolicyRuleType) ToFirewallPolicyRuleTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleTypePtrOutput {
	return FirewallPolicyRuleType(e).ToFirewallPolicyRuleTypeOutputWithContext(ctx).ToFirewallPolicyRuleTypePtrOutputWithContext(ctx)
}

func (e FirewallPolicyRuleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyRuleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicyRuleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallPolicyRuleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallPolicyRuleTypeOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicyRuleType)(nil)).Elem()
}

func (o FirewallPolicyRuleTypeOutput) ToFirewallPolicyRuleTypeOutput() FirewallPolicyRuleTypeOutput {
	return o
}

func (o FirewallPolicyRuleTypeOutput) ToFirewallPolicyRuleTypeOutputWithContext(ctx context.Context) FirewallPolicyRuleTypeOutput {
	return o
}

func (o FirewallPolicyRuleTypeOutput) ToFirewallPolicyRuleTypePtrOutput() FirewallPolicyRuleTypePtrOutput {
	return o.ToFirewallPolicyRuleTypePtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleTypeOutput) ToFirewallPolicyRuleTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicyRuleType) *FirewallPolicyRuleType {
		return &v
	}).(FirewallPolicyRuleTypePtrOutput)
}

func (o FirewallPolicyRuleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyRuleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallPolicyRuleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicyRuleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallPolicyRuleTypePtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleType)(nil)).Elem()
}

func (o FirewallPolicyRuleTypePtrOutput) ToFirewallPolicyRuleTypePtrOutput() FirewallPolicyRuleTypePtrOutput {
	return o
}

func (o FirewallPolicyRuleTypePtrOutput) ToFirewallPolicyRuleTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleTypePtrOutput {
	return o
}

func (o FirewallPolicyRuleTypePtrOutput) Elem() FirewallPolicyRuleTypeOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleType) FirewallPolicyRuleType {
		if v != nil {
			return *v
		}
		var ret FirewallPolicyRuleType
		return ret
	}).(FirewallPolicyRuleTypeOutput)
}

func (o FirewallPolicyRuleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyRuleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallPolicyRuleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirewallPolicyRuleTypeInput interface {
	pulumi.Input

	ToFirewallPolicyRuleTypeOutput() FirewallPolicyRuleTypeOutput
	ToFirewallPolicyRuleTypeOutputWithContext(context.Context) FirewallPolicyRuleTypeOutput
}

var firewallPolicyRuleTypePtrType = reflect.TypeOf((**FirewallPolicyRuleType)(nil)).Elem()

type FirewallPolicyRuleTypePtrInput interface {
	pulumi.Input

	ToFirewallPolicyRuleTypePtrOutput() FirewallPolicyRuleTypePtrOutput
	ToFirewallPolicyRuleTypePtrOutputWithContext(context.Context) FirewallPolicyRuleTypePtrOutput
}

type firewallPolicyRuleTypePtr string

func FirewallPolicyRuleTypePtr(v string) FirewallPolicyRuleTypePtrInput {
	return (*firewallPolicyRuleTypePtr)(&v)
}

func (*firewallPolicyRuleTypePtr) ElementType() reflect.Type {
	return firewallPolicyRuleTypePtrType
}

func (in *firewallPolicyRuleTypePtr) ToFirewallPolicyRuleTypePtrOutput() FirewallPolicyRuleTypePtrOutput {
	return pulumi.ToOutput(in).(FirewallPolicyRuleTypePtrOutput)
}

func (in *firewallPolicyRuleTypePtr) ToFirewallPolicyRuleTypePtrOutputWithContext(ctx context.Context) FirewallPolicyRuleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallPolicyRuleTypePtrOutput)
}

type FirewallPolicySkuTier string

const (
	FirewallPolicySkuTierStandard = FirewallPolicySkuTier("Standard")
	FirewallPolicySkuTierPremium  = FirewallPolicySkuTier("Premium")
)

func (FirewallPolicySkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicySkuTier)(nil)).Elem()
}

func (e FirewallPolicySkuTier) ToFirewallPolicySkuTierOutput() FirewallPolicySkuTierOutput {
	return pulumi.ToOutput(e).(FirewallPolicySkuTierOutput)
}

func (e FirewallPolicySkuTier) ToFirewallPolicySkuTierOutputWithContext(ctx context.Context) FirewallPolicySkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirewallPolicySkuTierOutput)
}

func (e FirewallPolicySkuTier) ToFirewallPolicySkuTierPtrOutput() FirewallPolicySkuTierPtrOutput {
	return e.ToFirewallPolicySkuTierPtrOutputWithContext(context.Background())
}

func (e FirewallPolicySkuTier) ToFirewallPolicySkuTierPtrOutputWithContext(ctx context.Context) FirewallPolicySkuTierPtrOutput {
	return FirewallPolicySkuTier(e).ToFirewallPolicySkuTierOutputWithContext(ctx).ToFirewallPolicySkuTierPtrOutputWithContext(ctx)
}

func (e FirewallPolicySkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicySkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirewallPolicySkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirewallPolicySkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirewallPolicySkuTierOutput struct{ *pulumi.OutputState }

func (FirewallPolicySkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicySkuTier)(nil)).Elem()
}

func (o FirewallPolicySkuTierOutput) ToFirewallPolicySkuTierOutput() FirewallPolicySkuTierOutput {
	return o
}

func (o FirewallPolicySkuTierOutput) ToFirewallPolicySkuTierOutputWithContext(ctx context.Context) FirewallPolicySkuTierOutput {
	return o
}

func (o FirewallPolicySkuTierOutput) ToFirewallPolicySkuTierPtrOutput() FirewallPolicySkuTierPtrOutput {
	return o.ToFirewallPolicySkuTierPtrOutputWithContext(context.Background())
}

func (o FirewallPolicySkuTierOutput) ToFirewallPolicySkuTierPtrOutputWithContext(ctx context.Context) FirewallPolicySkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirewallPolicySkuTier) *FirewallPolicySkuTier {
		return &v
	}).(FirewallPolicySkuTierPtrOutput)
}

func (o FirewallPolicySkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirewallPolicySkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicySkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirewallPolicySkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicySkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirewallPolicySkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirewallPolicySkuTierPtrOutput struct{ *pulumi.OutputState }

func (FirewallPolicySkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicySkuTier)(nil)).Elem()
}

func (o FirewallPolicySkuTierPtrOutput) ToFirewallPolicySkuTierPtrOutput() FirewallPolicySkuTierPtrOutput {
	return o
}

func (o FirewallPolicySkuTierPtrOutput) ToFirewallPolicySkuTierPtrOutputWithContext(ctx context.Context) FirewallPolicySkuTierPtrOutput {
	return o
}

func (o FirewallPolicySkuTierPtrOutput) Elem() FirewallPolicySkuTierOutput {
	return o.ApplyT(func(v *FirewallPolicySkuTier) FirewallPolicySkuTier {
		if v != nil {
			return *v
		}
		var ret FirewallPolicySkuTier
		return ret
	}).(FirewallPolicySkuTierOutput)
}

func (o FirewallPolicySkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirewallPolicySkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirewallPolicySkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirewallPolicySkuTierInput interface {
	pulumi.Input

	ToFirewallPolicySkuTierOutput() FirewallPolicySkuTierOutput
	ToFirewallPolicySkuTierOutputWithContext(context.Context) FirewallPolicySkuTierOutput
}

var firewallPolicySkuTierPtrType = reflect.TypeOf((**FirewallPolicySkuTier)(nil)).Elem()

type FirewallPolicySkuTierPtrInput interface {
	pulumi.Input

	ToFirewallPolicySkuTierPtrOutput() FirewallPolicySkuTierPtrOutput
	ToFirewallPolicySkuTierPtrOutputWithContext(context.Context) FirewallPolicySkuTierPtrOutput
}

type firewallPolicySkuTierPtr string

func FirewallPolicySkuTierPtr(v string) FirewallPolicySkuTierPtrInput {
	return (*firewallPolicySkuTierPtr)(&v)
}

func (*firewallPolicySkuTierPtr) ElementType() reflect.Type {
	return firewallPolicySkuTierPtrType
}

func (in *firewallPolicySkuTierPtr) ToFirewallPolicySkuTierPtrOutput() FirewallPolicySkuTierPtrOutput {
	return pulumi.ToOutput(in).(FirewallPolicySkuTierPtrOutput)
}

func (in *firewallPolicySkuTierPtr) ToFirewallPolicySkuTierPtrOutputWithContext(ctx context.Context) FirewallPolicySkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirewallPolicySkuTierPtrOutput)
}

type FlowLogFormatType string

const (
	FlowLogFormatTypeJSON = FlowLogFormatType("JSON")
)

func (FlowLogFormatType) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowLogFormatType)(nil)).Elem()
}

func (e FlowLogFormatType) ToFlowLogFormatTypeOutput() FlowLogFormatTypeOutput {
	return pulumi.ToOutput(e).(FlowLogFormatTypeOutput)
}

func (e FlowLogFormatType) ToFlowLogFormatTypeOutputWithContext(ctx context.Context) FlowLogFormatTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FlowLogFormatTypeOutput)
}

func (e FlowLogFormatType) ToFlowLogFormatTypePtrOutput() FlowLogFormatTypePtrOutput {
	return e.ToFlowLogFormatTypePtrOutputWithContext(context.Background())
}

func (e FlowLogFormatType) ToFlowLogFormatTypePtrOutputWithContext(ctx context.Context) FlowLogFormatTypePtrOutput {
	return FlowLogFormatType(e).ToFlowLogFormatTypeOutputWithContext(ctx).ToFlowLogFormatTypePtrOutputWithContext(ctx)
}

func (e FlowLogFormatType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FlowLogFormatType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FlowLogFormatType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FlowLogFormatType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FlowLogFormatTypeOutput struct{ *pulumi.OutputState }

func (FlowLogFormatTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowLogFormatType)(nil)).Elem()
}

func (o FlowLogFormatTypeOutput) ToFlowLogFormatTypeOutput() FlowLogFormatTypeOutput {
	return o
}

func (o FlowLogFormatTypeOutput) ToFlowLogFormatTypeOutputWithContext(ctx context.Context) FlowLogFormatTypeOutput {
	return o
}

func (o FlowLogFormatTypeOutput) ToFlowLogFormatTypePtrOutput() FlowLogFormatTypePtrOutput {
	return o.ToFlowLogFormatTypePtrOutputWithContext(context.Background())
}

func (o FlowLogFormatTypeOutput) ToFlowLogFormatTypePtrOutputWithContext(ctx context.Context) FlowLogFormatTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FlowLogFormatType) *FlowLogFormatType {
		return &v
	}).(FlowLogFormatTypePtrOutput)
}

func (o FlowLogFormatTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FlowLogFormatTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FlowLogFormatType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FlowLogFormatTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FlowLogFormatTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FlowLogFormatType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FlowLogFormatTypePtrOutput struct{ *pulumi.OutputState }

func (FlowLogFormatTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowLogFormatType)(nil)).Elem()
}

func (o FlowLogFormatTypePtrOutput) ToFlowLogFormatTypePtrOutput() FlowLogFormatTypePtrOutput {
	return o
}

func (o FlowLogFormatTypePtrOutput) ToFlowLogFormatTypePtrOutputWithContext(ctx context.Context) FlowLogFormatTypePtrOutput {
	return o
}

func (o FlowLogFormatTypePtrOutput) Elem() FlowLogFormatTypeOutput {
	return o.ApplyT(func(v *FlowLogFormatType) FlowLogFormatType {
		if v != nil {
			return *v
		}
		var ret FlowLogFormatType
		return ret
	}).(FlowLogFormatTypeOutput)
}

func (o FlowLogFormatTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FlowLogFormatTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FlowLogFormatType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FlowLogFormatTypeInput interface {
	pulumi.Input

	ToFlowLogFormatTypeOutput() FlowLogFormatTypeOutput
	ToFlowLogFormatTypeOutputWithContext(context.Context) FlowLogFormatTypeOutput
}

var flowLogFormatTypePtrType = reflect.TypeOf((**FlowLogFormatType)(nil)).Elem()

type FlowLogFormatTypePtrInput interface {
	pulumi.Input

	ToFlowLogFormatTypePtrOutput() FlowLogFormatTypePtrOutput
	ToFlowLogFormatTypePtrOutputWithContext(context.Context) FlowLogFormatTypePtrOutput
}

type flowLogFormatTypePtr string

func FlowLogFormatTypePtr(v string) FlowLogFormatTypePtrInput {
	return (*flowLogFormatTypePtr)(&v)
}

func (*flowLogFormatTypePtr) ElementType() reflect.Type {
	return flowLogFormatTypePtrType
}

func (in *flowLogFormatTypePtr) ToFlowLogFormatTypePtrOutput() FlowLogFormatTypePtrOutput {
	return pulumi.ToOutput(in).(FlowLogFormatTypePtrOutput)
}

func (in *flowLogFormatTypePtr) ToFlowLogFormatTypePtrOutputWithContext(ctx context.Context) FlowLogFormatTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FlowLogFormatTypePtrOutput)
}

type GatewayLoadBalancerTunnelInterfaceType string

const (
	GatewayLoadBalancerTunnelInterfaceTypeNone     = GatewayLoadBalancerTunnelInterfaceType("None")
	GatewayLoadBalancerTunnelInterfaceTypeInternal = GatewayLoadBalancerTunnelInterfaceType("Internal")
	GatewayLoadBalancerTunnelInterfaceTypeExternal = GatewayLoadBalancerTunnelInterfaceType("External")
)

func (GatewayLoadBalancerTunnelInterfaceType) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayLoadBalancerTunnelInterfaceType)(nil)).Elem()
}

func (e GatewayLoadBalancerTunnelInterfaceType) ToGatewayLoadBalancerTunnelInterfaceTypeOutput() GatewayLoadBalancerTunnelInterfaceTypeOutput {
	return pulumi.ToOutput(e).(GatewayLoadBalancerTunnelInterfaceTypeOutput)
}

func (e GatewayLoadBalancerTunnelInterfaceType) ToGatewayLoadBalancerTunnelInterfaceTypeOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelInterfaceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GatewayLoadBalancerTunnelInterfaceTypeOutput)
}

func (e GatewayLoadBalancerTunnelInterfaceType) ToGatewayLoadBalancerTunnelInterfaceTypePtrOutput() GatewayLoadBalancerTunnelInterfaceTypePtrOutput {
	return e.ToGatewayLoadBalancerTunnelInterfaceTypePtrOutputWithContext(context.Background())
}

func (e GatewayLoadBalancerTunnelInterfaceType) ToGatewayLoadBalancerTunnelInterfaceTypePtrOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelInterfaceTypePtrOutput {
	return GatewayLoadBalancerTunnelInterfaceType(e).ToGatewayLoadBalancerTunnelInterfaceTypeOutputWithContext(ctx).ToGatewayLoadBalancerTunnelInterfaceTypePtrOutputWithContext(ctx)
}

func (e GatewayLoadBalancerTunnelInterfaceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GatewayLoadBalancerTunnelInterfaceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GatewayLoadBalancerTunnelInterfaceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GatewayLoadBalancerTunnelInterfaceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GatewayLoadBalancerTunnelInterfaceTypeOutput struct{ *pulumi.OutputState }

func (GatewayLoadBalancerTunnelInterfaceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayLoadBalancerTunnelInterfaceType)(nil)).Elem()
}

func (o GatewayLoadBalancerTunnelInterfaceTypeOutput) ToGatewayLoadBalancerTunnelInterfaceTypeOutput() GatewayLoadBalancerTunnelInterfaceTypeOutput {
	return o
}

func (o GatewayLoadBalancerTunnelInterfaceTypeOutput) ToGatewayLoadBalancerTunnelInterfaceTypeOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelInterfaceTypeOutput {
	return o
}

func (o GatewayLoadBalancerTunnelInterfaceTypeOutput) ToGatewayLoadBalancerTunnelInterfaceTypePtrOutput() GatewayLoadBalancerTunnelInterfaceTypePtrOutput {
	return o.ToGatewayLoadBalancerTunnelInterfaceTypePtrOutputWithContext(context.Background())
}

func (o GatewayLoadBalancerTunnelInterfaceTypeOutput) ToGatewayLoadBalancerTunnelInterfaceTypePtrOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelInterfaceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayLoadBalancerTunnelInterfaceType) *GatewayLoadBalancerTunnelInterfaceType {
		return &v
	}).(GatewayLoadBalancerTunnelInterfaceTypePtrOutput)
}

func (o GatewayLoadBalancerTunnelInterfaceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GatewayLoadBalancerTunnelInterfaceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GatewayLoadBalancerTunnelInterfaceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GatewayLoadBalancerTunnelInterfaceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GatewayLoadBalancerTunnelInterfaceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GatewayLoadBalancerTunnelInterfaceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GatewayLoadBalancerTunnelInterfaceTypePtrOutput struct{ *pulumi.OutputState }

func (GatewayLoadBalancerTunnelInterfaceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayLoadBalancerTunnelInterfaceType)(nil)).Elem()
}

func (o GatewayLoadBalancerTunnelInterfaceTypePtrOutput) ToGatewayLoadBalancerTunnelInterfaceTypePtrOutput() GatewayLoadBalancerTunnelInterfaceTypePtrOutput {
	return o
}

func (o GatewayLoadBalancerTunnelInterfaceTypePtrOutput) ToGatewayLoadBalancerTunnelInterfaceTypePtrOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelInterfaceTypePtrOutput {
	return o
}

func (o GatewayLoadBalancerTunnelInterfaceTypePtrOutput) Elem() GatewayLoadBalancerTunnelInterfaceTypeOutput {
	return o.ApplyT(func(v *GatewayLoadBalancerTunnelInterfaceType) GatewayLoadBalancerTunnelInterfaceType {
		if v != nil {
			return *v
		}
		var ret GatewayLoadBalancerTunnelInterfaceType
		return ret
	}).(GatewayLoadBalancerTunnelInterfaceTypeOutput)
}

func (o GatewayLoadBalancerTunnelInterfaceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GatewayLoadBalancerTunnelInterfaceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GatewayLoadBalancerTunnelInterfaceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GatewayLoadBalancerTunnelInterfaceTypeInput interface {
	pulumi.Input

	ToGatewayLoadBalancerTunnelInterfaceTypeOutput() GatewayLoadBalancerTunnelInterfaceTypeOutput
	ToGatewayLoadBalancerTunnelInterfaceTypeOutputWithContext(context.Context) GatewayLoadBalancerTunnelInterfaceTypeOutput
}

var gatewayLoadBalancerTunnelInterfaceTypePtrType = reflect.TypeOf((**GatewayLoadBalancerTunnelInterfaceType)(nil)).Elem()

type GatewayLoadBalancerTunnelInterfaceTypePtrInput interface {
	pulumi.Input

	ToGatewayLoadBalancerTunnelInterfaceTypePtrOutput() GatewayLoadBalancerTunnelInterfaceTypePtrOutput
	ToGatewayLoadBalancerTunnelInterfaceTypePtrOutputWithContext(context.Context) GatewayLoadBalancerTunnelInterfaceTypePtrOutput
}

type gatewayLoadBalancerTunnelInterfaceTypePtr string

func GatewayLoadBalancerTunnelInterfaceTypePtr(v string) GatewayLoadBalancerTunnelInterfaceTypePtrInput {
	return (*gatewayLoadBalancerTunnelInterfaceTypePtr)(&v)
}

func (*gatewayLoadBalancerTunnelInterfaceTypePtr) ElementType() reflect.Type {
	return gatewayLoadBalancerTunnelInterfaceTypePtrType
}

func (in *gatewayLoadBalancerTunnelInterfaceTypePtr) ToGatewayLoadBalancerTunnelInterfaceTypePtrOutput() GatewayLoadBalancerTunnelInterfaceTypePtrOutput {
	return pulumi.ToOutput(in).(GatewayLoadBalancerTunnelInterfaceTypePtrOutput)
}

func (in *gatewayLoadBalancerTunnelInterfaceTypePtr) ToGatewayLoadBalancerTunnelInterfaceTypePtrOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelInterfaceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GatewayLoadBalancerTunnelInterfaceTypePtrOutput)
}

type GatewayLoadBalancerTunnelProtocol string

const (
	GatewayLoadBalancerTunnelProtocolNone   = GatewayLoadBalancerTunnelProtocol("None")
	GatewayLoadBalancerTunnelProtocolNative = GatewayLoadBalancerTunnelProtocol("Native")
	GatewayLoadBalancerTunnelProtocolVXLAN  = GatewayLoadBalancerTunnelProtocol("VXLAN")
)

func (GatewayLoadBalancerTunnelProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayLoadBalancerTunnelProtocol)(nil)).Elem()
}

func (e GatewayLoadBalancerTunnelProtocol) ToGatewayLoadBalancerTunnelProtocolOutput() GatewayLoadBalancerTunnelProtocolOutput {
	return pulumi.ToOutput(e).(GatewayLoadBalancerTunnelProtocolOutput)
}

func (e GatewayLoadBalancerTunnelProtocol) ToGatewayLoadBalancerTunnelProtocolOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GatewayLoadBalancerTunnelProtocolOutput)
}

func (e GatewayLoadBalancerTunnelProtocol) ToGatewayLoadBalancerTunnelProtocolPtrOutput() GatewayLoadBalancerTunnelProtocolPtrOutput {
	return e.ToGatewayLoadBalancerTunnelProtocolPtrOutputWithContext(context.Background())
}

func (e GatewayLoadBalancerTunnelProtocol) ToGatewayLoadBalancerTunnelProtocolPtrOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelProtocolPtrOutput {
	return GatewayLoadBalancerTunnelProtocol(e).ToGatewayLoadBalancerTunnelProtocolOutputWithContext(ctx).ToGatewayLoadBalancerTunnelProtocolPtrOutputWithContext(ctx)
}

func (e GatewayLoadBalancerTunnelProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GatewayLoadBalancerTunnelProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GatewayLoadBalancerTunnelProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GatewayLoadBalancerTunnelProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GatewayLoadBalancerTunnelProtocolOutput struct{ *pulumi.OutputState }

func (GatewayLoadBalancerTunnelProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayLoadBalancerTunnelProtocol)(nil)).Elem()
}

func (o GatewayLoadBalancerTunnelProtocolOutput) ToGatewayLoadBalancerTunnelProtocolOutput() GatewayLoadBalancerTunnelProtocolOutput {
	return o
}

func (o GatewayLoadBalancerTunnelProtocolOutput) ToGatewayLoadBalancerTunnelProtocolOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelProtocolOutput {
	return o
}

func (o GatewayLoadBalancerTunnelProtocolOutput) ToGatewayLoadBalancerTunnelProtocolPtrOutput() GatewayLoadBalancerTunnelProtocolPtrOutput {
	return o.ToGatewayLoadBalancerTunnelProtocolPtrOutputWithContext(context.Background())
}

func (o GatewayLoadBalancerTunnelProtocolOutput) ToGatewayLoadBalancerTunnelProtocolPtrOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayLoadBalancerTunnelProtocol) *GatewayLoadBalancerTunnelProtocol {
		return &v
	}).(GatewayLoadBalancerTunnelProtocolPtrOutput)
}

func (o GatewayLoadBalancerTunnelProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GatewayLoadBalancerTunnelProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GatewayLoadBalancerTunnelProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GatewayLoadBalancerTunnelProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GatewayLoadBalancerTunnelProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GatewayLoadBalancerTunnelProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GatewayLoadBalancerTunnelProtocolPtrOutput struct{ *pulumi.OutputState }

func (GatewayLoadBalancerTunnelProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayLoadBalancerTunnelProtocol)(nil)).Elem()
}

func (o GatewayLoadBalancerTunnelProtocolPtrOutput) ToGatewayLoadBalancerTunnelProtocolPtrOutput() GatewayLoadBalancerTunnelProtocolPtrOutput {
	return o
}

func (o GatewayLoadBalancerTunnelProtocolPtrOutput) ToGatewayLoadBalancerTunnelProtocolPtrOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelProtocolPtrOutput {
	return o
}

func (o GatewayLoadBalancerTunnelProtocolPtrOutput) Elem() GatewayLoadBalancerTunnelProtocolOutput {
	return o.ApplyT(func(v *GatewayLoadBalancerTunnelProtocol) GatewayLoadBalancerTunnelProtocol {
		if v != nil {
			return *v
		}
		var ret GatewayLoadBalancerTunnelProtocol
		return ret
	}).(GatewayLoadBalancerTunnelProtocolOutput)
}

func (o GatewayLoadBalancerTunnelProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GatewayLoadBalancerTunnelProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GatewayLoadBalancerTunnelProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GatewayLoadBalancerTunnelProtocolInput interface {
	pulumi.Input

	ToGatewayLoadBalancerTunnelProtocolOutput() GatewayLoadBalancerTunnelProtocolOutput
	ToGatewayLoadBalancerTunnelProtocolOutputWithContext(context.Context) GatewayLoadBalancerTunnelProtocolOutput
}

var gatewayLoadBalancerTunnelProtocolPtrType = reflect.TypeOf((**GatewayLoadBalancerTunnelProtocol)(nil)).Elem()

type GatewayLoadBalancerTunnelProtocolPtrInput interface {
	pulumi.Input

	ToGatewayLoadBalancerTunnelProtocolPtrOutput() GatewayLoadBalancerTunnelProtocolPtrOutput
	ToGatewayLoadBalancerTunnelProtocolPtrOutputWithContext(context.Context) GatewayLoadBalancerTunnelProtocolPtrOutput
}

type gatewayLoadBalancerTunnelProtocolPtr string

func GatewayLoadBalancerTunnelProtocolPtr(v string) GatewayLoadBalancerTunnelProtocolPtrInput {
	return (*gatewayLoadBalancerTunnelProtocolPtr)(&v)
}

func (*gatewayLoadBalancerTunnelProtocolPtr) ElementType() reflect.Type {
	return gatewayLoadBalancerTunnelProtocolPtrType
}

func (in *gatewayLoadBalancerTunnelProtocolPtr) ToGatewayLoadBalancerTunnelProtocolPtrOutput() GatewayLoadBalancerTunnelProtocolPtrOutput {
	return pulumi.ToOutput(in).(GatewayLoadBalancerTunnelProtocolPtrOutput)
}

func (in *gatewayLoadBalancerTunnelProtocolPtr) ToGatewayLoadBalancerTunnelProtocolPtrOutputWithContext(ctx context.Context) GatewayLoadBalancerTunnelProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GatewayLoadBalancerTunnelProtocolPtrOutput)
}

type HTTPConfigurationMethod string

const (
	HTTPConfigurationMethodGet  = HTTPConfigurationMethod("Get")
	HTTPConfigurationMethodPost = HTTPConfigurationMethod("Post")
)

func (HTTPConfigurationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*HTTPConfigurationMethod)(nil)).Elem()
}

func (e HTTPConfigurationMethod) ToHTTPConfigurationMethodOutput() HTTPConfigurationMethodOutput {
	return pulumi.ToOutput(e).(HTTPConfigurationMethodOutput)
}

func (e HTTPConfigurationMethod) ToHTTPConfigurationMethodOutputWithContext(ctx context.Context) HTTPConfigurationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HTTPConfigurationMethodOutput)
}

func (e HTTPConfigurationMethod) ToHTTPConfigurationMethodPtrOutput() HTTPConfigurationMethodPtrOutput {
	return e.ToHTTPConfigurationMethodPtrOutputWithContext(context.Background())
}

func (e HTTPConfigurationMethod) ToHTTPConfigurationMethodPtrOutputWithContext(ctx context.Context) HTTPConfigurationMethodPtrOutput {
	return HTTPConfigurationMethod(e).ToHTTPConfigurationMethodOutputWithContext(ctx).ToHTTPConfigurationMethodPtrOutputWithContext(ctx)
}

func (e HTTPConfigurationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HTTPConfigurationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HTTPConfigurationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HTTPConfigurationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HTTPConfigurationMethodOutput struct{ *pulumi.OutputState }

func (HTTPConfigurationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HTTPConfigurationMethod)(nil)).Elem()
}

func (o HTTPConfigurationMethodOutput) ToHTTPConfigurationMethodOutput() HTTPConfigurationMethodOutput {
	return o
}

func (o HTTPConfigurationMethodOutput) ToHTTPConfigurationMethodOutputWithContext(ctx context.Context) HTTPConfigurationMethodOutput {
	return o
}

func (o HTTPConfigurationMethodOutput) ToHTTPConfigurationMethodPtrOutput() HTTPConfigurationMethodPtrOutput {
	return o.ToHTTPConfigurationMethodPtrOutputWithContext(context.Background())
}

func (o HTTPConfigurationMethodOutput) ToHTTPConfigurationMethodPtrOutputWithContext(ctx context.Context) HTTPConfigurationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HTTPConfigurationMethod) *HTTPConfigurationMethod {
		return &v
	}).(HTTPConfigurationMethodPtrOutput)
}

func (o HTTPConfigurationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HTTPConfigurationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HTTPConfigurationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HTTPConfigurationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HTTPConfigurationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HTTPConfigurationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HTTPConfigurationMethodPtrOutput struct{ *pulumi.OutputState }

func (HTTPConfigurationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HTTPConfigurationMethod)(nil)).Elem()
}

func (o HTTPConfigurationMethodPtrOutput) ToHTTPConfigurationMethodPtrOutput() HTTPConfigurationMethodPtrOutput {
	return o
}

func (o HTTPConfigurationMethodPtrOutput) ToHTTPConfigurationMethodPtrOutputWithContext(ctx context.Context) HTTPConfigurationMethodPtrOutput {
	return o
}

func (o HTTPConfigurationMethodPtrOutput) Elem() HTTPConfigurationMethodOutput {
	return o.ApplyT(func(v *HTTPConfigurationMethod) HTTPConfigurationMethod {
		if v != nil {
			return *v
		}
		var ret HTTPConfigurationMethod
		return ret
	}).(HTTPConfigurationMethodOutput)
}

func (o HTTPConfigurationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HTTPConfigurationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HTTPConfigurationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HTTPConfigurationMethodInput interface {
	pulumi.Input

	ToHTTPConfigurationMethodOutput() HTTPConfigurationMethodOutput
	ToHTTPConfigurationMethodOutputWithContext(context.Context) HTTPConfigurationMethodOutput
}

var httpconfigurationMethodPtrType = reflect.TypeOf((**HTTPConfigurationMethod)(nil)).Elem()

type HTTPConfigurationMethodPtrInput interface {
	pulumi.Input

	ToHTTPConfigurationMethodPtrOutput() HTTPConfigurationMethodPtrOutput
	ToHTTPConfigurationMethodPtrOutputWithContext(context.Context) HTTPConfigurationMethodPtrOutput
}

type httpconfigurationMethodPtr string

func HTTPConfigurationMethodPtr(v string) HTTPConfigurationMethodPtrInput {
	return (*httpconfigurationMethodPtr)(&v)
}

func (*httpconfigurationMethodPtr) ElementType() reflect.Type {
	return httpconfigurationMethodPtrType
}

func (in *httpconfigurationMethodPtr) ToHTTPConfigurationMethodPtrOutput() HTTPConfigurationMethodPtrOutput {
	return pulumi.ToOutput(in).(HTTPConfigurationMethodPtrOutput)
}

func (in *httpconfigurationMethodPtr) ToHTTPConfigurationMethodPtrOutputWithContext(ctx context.Context) HTTPConfigurationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HTTPConfigurationMethodPtrOutput)
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
	IkeEncryptionDES       = IkeEncryption("DES")
	IkeEncryptionDES3      = IkeEncryption("DES3")
	IkeEncryptionAES128    = IkeEncryption("AES128")
	IkeEncryptionAES192    = IkeEncryption("AES192")
	IkeEncryptionAES256    = IkeEncryption("AES256")
	IkeEncryptionGCMAES256 = IkeEncryption("GCMAES256")
	IkeEncryptionGCMAES128 = IkeEncryption("GCMAES128")
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
	IkeIntegrityMD5       = IkeIntegrity("MD5")
	IkeIntegritySHA1      = IkeIntegrity("SHA1")
	IkeIntegritySHA256    = IkeIntegrity("SHA256")
	IkeIntegritySHA384    = IkeIntegrity("SHA384")
	IkeIntegrityGCMAES256 = IkeIntegrity("GCMAES256")
	IkeIntegrityGCMAES128 = IkeIntegrity("GCMAES128")
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

type IpAllocationType string

const (
	IpAllocationTypeUndefined = IpAllocationType("Undefined")
	IpAllocationTypeHypernet  = IpAllocationType("Hypernet")
)

func (IpAllocationType) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAllocationType)(nil)).Elem()
}

func (e IpAllocationType) ToIpAllocationTypeOutput() IpAllocationTypeOutput {
	return pulumi.ToOutput(e).(IpAllocationTypeOutput)
}

func (e IpAllocationType) ToIpAllocationTypeOutputWithContext(ctx context.Context) IpAllocationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpAllocationTypeOutput)
}

func (e IpAllocationType) ToIpAllocationTypePtrOutput() IpAllocationTypePtrOutput {
	return e.ToIpAllocationTypePtrOutputWithContext(context.Background())
}

func (e IpAllocationType) ToIpAllocationTypePtrOutputWithContext(ctx context.Context) IpAllocationTypePtrOutput {
	return IpAllocationType(e).ToIpAllocationTypeOutputWithContext(ctx).ToIpAllocationTypePtrOutputWithContext(ctx)
}

func (e IpAllocationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAllocationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAllocationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpAllocationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpAllocationTypeOutput struct{ *pulumi.OutputState }

func (IpAllocationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAllocationType)(nil)).Elem()
}

func (o IpAllocationTypeOutput) ToIpAllocationTypeOutput() IpAllocationTypeOutput {
	return o
}

func (o IpAllocationTypeOutput) ToIpAllocationTypeOutputWithContext(ctx context.Context) IpAllocationTypeOutput {
	return o
}

func (o IpAllocationTypeOutput) ToIpAllocationTypePtrOutput() IpAllocationTypePtrOutput {
	return o.ToIpAllocationTypePtrOutputWithContext(context.Background())
}

func (o IpAllocationTypeOutput) ToIpAllocationTypePtrOutputWithContext(ctx context.Context) IpAllocationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpAllocationType) *IpAllocationType {
		return &v
	}).(IpAllocationTypePtrOutput)
}

func (o IpAllocationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpAllocationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAllocationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpAllocationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAllocationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAllocationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpAllocationTypePtrOutput struct{ *pulumi.OutputState }

func (IpAllocationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAllocationType)(nil)).Elem()
}

func (o IpAllocationTypePtrOutput) ToIpAllocationTypePtrOutput() IpAllocationTypePtrOutput {
	return o
}

func (o IpAllocationTypePtrOutput) ToIpAllocationTypePtrOutputWithContext(ctx context.Context) IpAllocationTypePtrOutput {
	return o
}

func (o IpAllocationTypePtrOutput) Elem() IpAllocationTypeOutput {
	return o.ApplyT(func(v *IpAllocationType) IpAllocationType {
		if v != nil {
			return *v
		}
		var ret IpAllocationType
		return ret
	}).(IpAllocationTypeOutput)
}

func (o IpAllocationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAllocationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpAllocationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IpAllocationTypeInput interface {
	pulumi.Input

	ToIpAllocationTypeOutput() IpAllocationTypeOutput
	ToIpAllocationTypeOutputWithContext(context.Context) IpAllocationTypeOutput
}

var ipAllocationTypePtrType = reflect.TypeOf((**IpAllocationType)(nil)).Elem()

type IpAllocationTypePtrInput interface {
	pulumi.Input

	ToIpAllocationTypePtrOutput() IpAllocationTypePtrOutput
	ToIpAllocationTypePtrOutputWithContext(context.Context) IpAllocationTypePtrOutput
}

type ipAllocationTypePtr string

func IpAllocationTypePtr(v string) IpAllocationTypePtrInput {
	return (*ipAllocationTypePtr)(&v)
}

func (*ipAllocationTypePtr) ElementType() reflect.Type {
	return ipAllocationTypePtrType
}

func (in *ipAllocationTypePtr) ToIpAllocationTypePtrOutput() IpAllocationTypePtrOutput {
	return pulumi.ToOutput(in).(IpAllocationTypePtrOutput)
}

func (in *ipAllocationTypePtr) ToIpAllocationTypePtrOutputWithContext(ctx context.Context) IpAllocationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpAllocationTypePtrOutput)
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

type LoadBalancerOutboundRuleProtocol string

const (
	LoadBalancerOutboundRuleProtocolTcp = LoadBalancerOutboundRuleProtocol("Tcp")
	LoadBalancerOutboundRuleProtocolUdp = LoadBalancerOutboundRuleProtocol("Udp")
	LoadBalancerOutboundRuleProtocolAll = LoadBalancerOutboundRuleProtocol("All")
)

func (LoadBalancerOutboundRuleProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerOutboundRuleProtocol)(nil)).Elem()
}

func (e LoadBalancerOutboundRuleProtocol) ToLoadBalancerOutboundRuleProtocolOutput() LoadBalancerOutboundRuleProtocolOutput {
	return pulumi.ToOutput(e).(LoadBalancerOutboundRuleProtocolOutput)
}

func (e LoadBalancerOutboundRuleProtocol) ToLoadBalancerOutboundRuleProtocolOutputWithContext(ctx context.Context) LoadBalancerOutboundRuleProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoadBalancerOutboundRuleProtocolOutput)
}

func (e LoadBalancerOutboundRuleProtocol) ToLoadBalancerOutboundRuleProtocolPtrOutput() LoadBalancerOutboundRuleProtocolPtrOutput {
	return e.ToLoadBalancerOutboundRuleProtocolPtrOutputWithContext(context.Background())
}

func (e LoadBalancerOutboundRuleProtocol) ToLoadBalancerOutboundRuleProtocolPtrOutputWithContext(ctx context.Context) LoadBalancerOutboundRuleProtocolPtrOutput {
	return LoadBalancerOutboundRuleProtocol(e).ToLoadBalancerOutboundRuleProtocolOutputWithContext(ctx).ToLoadBalancerOutboundRuleProtocolPtrOutputWithContext(ctx)
}

func (e LoadBalancerOutboundRuleProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerOutboundRuleProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerOutboundRuleProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoadBalancerOutboundRuleProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoadBalancerOutboundRuleProtocolOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutboundRuleProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerOutboundRuleProtocol)(nil)).Elem()
}

func (o LoadBalancerOutboundRuleProtocolOutput) ToLoadBalancerOutboundRuleProtocolOutput() LoadBalancerOutboundRuleProtocolOutput {
	return o
}

func (o LoadBalancerOutboundRuleProtocolOutput) ToLoadBalancerOutboundRuleProtocolOutputWithContext(ctx context.Context) LoadBalancerOutboundRuleProtocolOutput {
	return o
}

func (o LoadBalancerOutboundRuleProtocolOutput) ToLoadBalancerOutboundRuleProtocolPtrOutput() LoadBalancerOutboundRuleProtocolPtrOutput {
	return o.ToLoadBalancerOutboundRuleProtocolPtrOutputWithContext(context.Background())
}

func (o LoadBalancerOutboundRuleProtocolOutput) ToLoadBalancerOutboundRuleProtocolPtrOutputWithContext(ctx context.Context) LoadBalancerOutboundRuleProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerOutboundRuleProtocol) *LoadBalancerOutboundRuleProtocol {
		return &v
	}).(LoadBalancerOutboundRuleProtocolPtrOutput)
}

func (o LoadBalancerOutboundRuleProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoadBalancerOutboundRuleProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerOutboundRuleProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoadBalancerOutboundRuleProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerOutboundRuleProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerOutboundRuleProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerOutboundRuleProtocolPtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutboundRuleProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerOutboundRuleProtocol)(nil)).Elem()
}

func (o LoadBalancerOutboundRuleProtocolPtrOutput) ToLoadBalancerOutboundRuleProtocolPtrOutput() LoadBalancerOutboundRuleProtocolPtrOutput {
	return o
}

func (o LoadBalancerOutboundRuleProtocolPtrOutput) ToLoadBalancerOutboundRuleProtocolPtrOutputWithContext(ctx context.Context) LoadBalancerOutboundRuleProtocolPtrOutput {
	return o
}

func (o LoadBalancerOutboundRuleProtocolPtrOutput) Elem() LoadBalancerOutboundRuleProtocolOutput {
	return o.ApplyT(func(v *LoadBalancerOutboundRuleProtocol) LoadBalancerOutboundRuleProtocol {
		if v != nil {
			return *v
		}
		var ret LoadBalancerOutboundRuleProtocol
		return ret
	}).(LoadBalancerOutboundRuleProtocolOutput)
}

func (o LoadBalancerOutboundRuleProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerOutboundRuleProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoadBalancerOutboundRuleProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LoadBalancerOutboundRuleProtocolInput interface {
	pulumi.Input

	ToLoadBalancerOutboundRuleProtocolOutput() LoadBalancerOutboundRuleProtocolOutput
	ToLoadBalancerOutboundRuleProtocolOutputWithContext(context.Context) LoadBalancerOutboundRuleProtocolOutput
}

var loadBalancerOutboundRuleProtocolPtrType = reflect.TypeOf((**LoadBalancerOutboundRuleProtocol)(nil)).Elem()

type LoadBalancerOutboundRuleProtocolPtrInput interface {
	pulumi.Input

	ToLoadBalancerOutboundRuleProtocolPtrOutput() LoadBalancerOutboundRuleProtocolPtrOutput
	ToLoadBalancerOutboundRuleProtocolPtrOutputWithContext(context.Context) LoadBalancerOutboundRuleProtocolPtrOutput
}

type loadBalancerOutboundRuleProtocolPtr string

func LoadBalancerOutboundRuleProtocolPtr(v string) LoadBalancerOutboundRuleProtocolPtrInput {
	return (*loadBalancerOutboundRuleProtocolPtr)(&v)
}

func (*loadBalancerOutboundRuleProtocolPtr) ElementType() reflect.Type {
	return loadBalancerOutboundRuleProtocolPtrType
}

func (in *loadBalancerOutboundRuleProtocolPtr) ToLoadBalancerOutboundRuleProtocolPtrOutput() LoadBalancerOutboundRuleProtocolPtrOutput {
	return pulumi.ToOutput(in).(LoadBalancerOutboundRuleProtocolPtrOutput)
}

func (in *loadBalancerOutboundRuleProtocolPtr) ToLoadBalancerOutboundRuleProtocolPtrOutputWithContext(ctx context.Context) LoadBalancerOutboundRuleProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoadBalancerOutboundRuleProtocolPtrOutput)
}

type LoadBalancerSkuName string

const (
	LoadBalancerSkuNameBasic    = LoadBalancerSkuName("Basic")
	LoadBalancerSkuNameStandard = LoadBalancerSkuName("Standard")
	LoadBalancerSkuNameGateway  = LoadBalancerSkuName("Gateway")
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

type LoadBalancerSkuTier string

const (
	LoadBalancerSkuTierRegional = LoadBalancerSkuTier("Regional")
	LoadBalancerSkuTierGlobal   = LoadBalancerSkuTier("Global")
)

func (LoadBalancerSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSkuTier)(nil)).Elem()
}

func (e LoadBalancerSkuTier) ToLoadBalancerSkuTierOutput() LoadBalancerSkuTierOutput {
	return pulumi.ToOutput(e).(LoadBalancerSkuTierOutput)
}

func (e LoadBalancerSkuTier) ToLoadBalancerSkuTierOutputWithContext(ctx context.Context) LoadBalancerSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoadBalancerSkuTierOutput)
}

func (e LoadBalancerSkuTier) ToLoadBalancerSkuTierPtrOutput() LoadBalancerSkuTierPtrOutput {
	return e.ToLoadBalancerSkuTierPtrOutputWithContext(context.Background())
}

func (e LoadBalancerSkuTier) ToLoadBalancerSkuTierPtrOutputWithContext(ctx context.Context) LoadBalancerSkuTierPtrOutput {
	return LoadBalancerSkuTier(e).ToLoadBalancerSkuTierOutputWithContext(ctx).ToLoadBalancerSkuTierPtrOutputWithContext(ctx)
}

func (e LoadBalancerSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoadBalancerSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoadBalancerSkuTierOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerSkuTier)(nil)).Elem()
}

func (o LoadBalancerSkuTierOutput) ToLoadBalancerSkuTierOutput() LoadBalancerSkuTierOutput {
	return o
}

func (o LoadBalancerSkuTierOutput) ToLoadBalancerSkuTierOutputWithContext(ctx context.Context) LoadBalancerSkuTierOutput {
	return o
}

func (o LoadBalancerSkuTierOutput) ToLoadBalancerSkuTierPtrOutput() LoadBalancerSkuTierPtrOutput {
	return o.ToLoadBalancerSkuTierPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuTierOutput) ToLoadBalancerSkuTierPtrOutputWithContext(ctx context.Context) LoadBalancerSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerSkuTier) *LoadBalancerSkuTier {
		return &v
	}).(LoadBalancerSkuTierPtrOutput)
}

func (o LoadBalancerSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoadBalancerSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoadBalancerSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerSkuTierPtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerSkuTier)(nil)).Elem()
}

func (o LoadBalancerSkuTierPtrOutput) ToLoadBalancerSkuTierPtrOutput() LoadBalancerSkuTierPtrOutput {
	return o
}

func (o LoadBalancerSkuTierPtrOutput) ToLoadBalancerSkuTierPtrOutputWithContext(ctx context.Context) LoadBalancerSkuTierPtrOutput {
	return o
}

func (o LoadBalancerSkuTierPtrOutput) Elem() LoadBalancerSkuTierOutput {
	return o.ApplyT(func(v *LoadBalancerSkuTier) LoadBalancerSkuTier {
		if v != nil {
			return *v
		}
		var ret LoadBalancerSkuTier
		return ret
	}).(LoadBalancerSkuTierOutput)
}

func (o LoadBalancerSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoadBalancerSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LoadBalancerSkuTierInput interface {
	pulumi.Input

	ToLoadBalancerSkuTierOutput() LoadBalancerSkuTierOutput
	ToLoadBalancerSkuTierOutputWithContext(context.Context) LoadBalancerSkuTierOutput
}

var loadBalancerSkuTierPtrType = reflect.TypeOf((**LoadBalancerSkuTier)(nil)).Elem()

type LoadBalancerSkuTierPtrInput interface {
	pulumi.Input

	ToLoadBalancerSkuTierPtrOutput() LoadBalancerSkuTierPtrOutput
	ToLoadBalancerSkuTierPtrOutputWithContext(context.Context) LoadBalancerSkuTierPtrOutput
}

type loadBalancerSkuTierPtr string

func LoadBalancerSkuTierPtr(v string) LoadBalancerSkuTierPtrInput {
	return (*loadBalancerSkuTierPtr)(&v)
}

func (*loadBalancerSkuTierPtr) ElementType() reflect.Type {
	return loadBalancerSkuTierPtrType
}

func (in *loadBalancerSkuTierPtr) ToLoadBalancerSkuTierPtrOutput() LoadBalancerSkuTierPtrOutput {
	return pulumi.ToOutput(in).(LoadBalancerSkuTierPtrOutput)
}

func (in *loadBalancerSkuTierPtr) ToLoadBalancerSkuTierPtrOutputWithContext(ctx context.Context) LoadBalancerSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoadBalancerSkuTierPtrOutput)
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

type ManagedRuleEnabledState string

const (
	ManagedRuleEnabledStateDisabled = ManagedRuleEnabledState("Disabled")
)

func (ManagedRuleEnabledState) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleEnabledState)(nil)).Elem()
}

func (e ManagedRuleEnabledState) ToManagedRuleEnabledStateOutput() ManagedRuleEnabledStateOutput {
	return pulumi.ToOutput(e).(ManagedRuleEnabledStateOutput)
}

func (e ManagedRuleEnabledState) ToManagedRuleEnabledStateOutputWithContext(ctx context.Context) ManagedRuleEnabledStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedRuleEnabledStateOutput)
}

func (e ManagedRuleEnabledState) ToManagedRuleEnabledStatePtrOutput() ManagedRuleEnabledStatePtrOutput {
	return e.ToManagedRuleEnabledStatePtrOutputWithContext(context.Background())
}

func (e ManagedRuleEnabledState) ToManagedRuleEnabledStatePtrOutputWithContext(ctx context.Context) ManagedRuleEnabledStatePtrOutput {
	return ManagedRuleEnabledState(e).ToManagedRuleEnabledStateOutputWithContext(ctx).ToManagedRuleEnabledStatePtrOutputWithContext(ctx)
}

func (e ManagedRuleEnabledState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedRuleEnabledState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedRuleEnabledState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedRuleEnabledState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedRuleEnabledStateOutput struct{ *pulumi.OutputState }

func (ManagedRuleEnabledStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleEnabledState)(nil)).Elem()
}

func (o ManagedRuleEnabledStateOutput) ToManagedRuleEnabledStateOutput() ManagedRuleEnabledStateOutput {
	return o
}

func (o ManagedRuleEnabledStateOutput) ToManagedRuleEnabledStateOutputWithContext(ctx context.Context) ManagedRuleEnabledStateOutput {
	return o
}

func (o ManagedRuleEnabledStateOutput) ToManagedRuleEnabledStatePtrOutput() ManagedRuleEnabledStatePtrOutput {
	return o.ToManagedRuleEnabledStatePtrOutputWithContext(context.Background())
}

func (o ManagedRuleEnabledStateOutput) ToManagedRuleEnabledStatePtrOutputWithContext(ctx context.Context) ManagedRuleEnabledStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedRuleEnabledState) *ManagedRuleEnabledState {
		return &v
	}).(ManagedRuleEnabledStatePtrOutput)
}

func (o ManagedRuleEnabledStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedRuleEnabledStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedRuleEnabledState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedRuleEnabledStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedRuleEnabledStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedRuleEnabledState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedRuleEnabledStatePtrOutput struct{ *pulumi.OutputState }

func (ManagedRuleEnabledStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleEnabledState)(nil)).Elem()
}

func (o ManagedRuleEnabledStatePtrOutput) ToManagedRuleEnabledStatePtrOutput() ManagedRuleEnabledStatePtrOutput {
	return o
}

func (o ManagedRuleEnabledStatePtrOutput) ToManagedRuleEnabledStatePtrOutputWithContext(ctx context.Context) ManagedRuleEnabledStatePtrOutput {
	return o
}

func (o ManagedRuleEnabledStatePtrOutput) Elem() ManagedRuleEnabledStateOutput {
	return o.ApplyT(func(v *ManagedRuleEnabledState) ManagedRuleEnabledState {
		if v != nil {
			return *v
		}
		var ret ManagedRuleEnabledState
		return ret
	}).(ManagedRuleEnabledStateOutput)
}

func (o ManagedRuleEnabledStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedRuleEnabledStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedRuleEnabledState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedRuleEnabledStateInput interface {
	pulumi.Input

	ToManagedRuleEnabledStateOutput() ManagedRuleEnabledStateOutput
	ToManagedRuleEnabledStateOutputWithContext(context.Context) ManagedRuleEnabledStateOutput
}

var managedRuleEnabledStatePtrType = reflect.TypeOf((**ManagedRuleEnabledState)(nil)).Elem()

type ManagedRuleEnabledStatePtrInput interface {
	pulumi.Input

	ToManagedRuleEnabledStatePtrOutput() ManagedRuleEnabledStatePtrOutput
	ToManagedRuleEnabledStatePtrOutputWithContext(context.Context) ManagedRuleEnabledStatePtrOutput
}

type managedRuleEnabledStatePtr string

func ManagedRuleEnabledStatePtr(v string) ManagedRuleEnabledStatePtrInput {
	return (*managedRuleEnabledStatePtr)(&v)
}

func (*managedRuleEnabledStatePtr) ElementType() reflect.Type {
	return managedRuleEnabledStatePtrType
}

func (in *managedRuleEnabledStatePtr) ToManagedRuleEnabledStatePtrOutput() ManagedRuleEnabledStatePtrOutput {
	return pulumi.ToOutput(in).(ManagedRuleEnabledStatePtrOutput)
}

func (in *managedRuleEnabledStatePtr) ToManagedRuleEnabledStatePtrOutputWithContext(ctx context.Context) ManagedRuleEnabledStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedRuleEnabledStatePtrOutput)
}

type NatGatewaySkuName string

const (
	NatGatewaySkuNameStandard = NatGatewaySkuName("Standard")
)

func (NatGatewaySkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*NatGatewaySkuName)(nil)).Elem()
}

func (e NatGatewaySkuName) ToNatGatewaySkuNameOutput() NatGatewaySkuNameOutput {
	return pulumi.ToOutput(e).(NatGatewaySkuNameOutput)
}

func (e NatGatewaySkuName) ToNatGatewaySkuNameOutputWithContext(ctx context.Context) NatGatewaySkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NatGatewaySkuNameOutput)
}

func (e NatGatewaySkuName) ToNatGatewaySkuNamePtrOutput() NatGatewaySkuNamePtrOutput {
	return e.ToNatGatewaySkuNamePtrOutputWithContext(context.Background())
}

func (e NatGatewaySkuName) ToNatGatewaySkuNamePtrOutputWithContext(ctx context.Context) NatGatewaySkuNamePtrOutput {
	return NatGatewaySkuName(e).ToNatGatewaySkuNameOutputWithContext(ctx).ToNatGatewaySkuNamePtrOutputWithContext(ctx)
}

func (e NatGatewaySkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NatGatewaySkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NatGatewaySkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NatGatewaySkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NatGatewaySkuNameOutput struct{ *pulumi.OutputState }

func (NatGatewaySkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NatGatewaySkuName)(nil)).Elem()
}

func (o NatGatewaySkuNameOutput) ToNatGatewaySkuNameOutput() NatGatewaySkuNameOutput {
	return o
}

func (o NatGatewaySkuNameOutput) ToNatGatewaySkuNameOutputWithContext(ctx context.Context) NatGatewaySkuNameOutput {
	return o
}

func (o NatGatewaySkuNameOutput) ToNatGatewaySkuNamePtrOutput() NatGatewaySkuNamePtrOutput {
	return o.ToNatGatewaySkuNamePtrOutputWithContext(context.Background())
}

func (o NatGatewaySkuNameOutput) ToNatGatewaySkuNamePtrOutputWithContext(ctx context.Context) NatGatewaySkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NatGatewaySkuName) *NatGatewaySkuName {
		return &v
	}).(NatGatewaySkuNamePtrOutput)
}

func (o NatGatewaySkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NatGatewaySkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NatGatewaySkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NatGatewaySkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NatGatewaySkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NatGatewaySkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NatGatewaySkuNamePtrOutput struct{ *pulumi.OutputState }

func (NatGatewaySkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGatewaySkuName)(nil)).Elem()
}

func (o NatGatewaySkuNamePtrOutput) ToNatGatewaySkuNamePtrOutput() NatGatewaySkuNamePtrOutput {
	return o
}

func (o NatGatewaySkuNamePtrOutput) ToNatGatewaySkuNamePtrOutputWithContext(ctx context.Context) NatGatewaySkuNamePtrOutput {
	return o
}

func (o NatGatewaySkuNamePtrOutput) Elem() NatGatewaySkuNameOutput {
	return o.ApplyT(func(v *NatGatewaySkuName) NatGatewaySkuName {
		if v != nil {
			return *v
		}
		var ret NatGatewaySkuName
		return ret
	}).(NatGatewaySkuNameOutput)
}

func (o NatGatewaySkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NatGatewaySkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NatGatewaySkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NatGatewaySkuNameInput interface {
	pulumi.Input

	ToNatGatewaySkuNameOutput() NatGatewaySkuNameOutput
	ToNatGatewaySkuNameOutputWithContext(context.Context) NatGatewaySkuNameOutput
}

var natGatewaySkuNamePtrType = reflect.TypeOf((**NatGatewaySkuName)(nil)).Elem()

type NatGatewaySkuNamePtrInput interface {
	pulumi.Input

	ToNatGatewaySkuNamePtrOutput() NatGatewaySkuNamePtrOutput
	ToNatGatewaySkuNamePtrOutputWithContext(context.Context) NatGatewaySkuNamePtrOutput
}

type natGatewaySkuNamePtr string

func NatGatewaySkuNamePtr(v string) NatGatewaySkuNamePtrInput {
	return (*natGatewaySkuNamePtr)(&v)
}

func (*natGatewaySkuNamePtr) ElementType() reflect.Type {
	return natGatewaySkuNamePtrType
}

func (in *natGatewaySkuNamePtr) ToNatGatewaySkuNamePtrOutput() NatGatewaySkuNamePtrOutput {
	return pulumi.ToOutput(in).(NatGatewaySkuNamePtrOutput)
}

func (in *natGatewaySkuNamePtr) ToNatGatewaySkuNamePtrOutputWithContext(ctx context.Context) NatGatewaySkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NatGatewaySkuNamePtrOutput)
}

type NetworkInterfaceMigrationPhase string

const (
	NetworkInterfaceMigrationPhaseNone      = NetworkInterfaceMigrationPhase("None")
	NetworkInterfaceMigrationPhasePrepare   = NetworkInterfaceMigrationPhase("Prepare")
	NetworkInterfaceMigrationPhaseCommit    = NetworkInterfaceMigrationPhase("Commit")
	NetworkInterfaceMigrationPhaseAbort     = NetworkInterfaceMigrationPhase("Abort")
	NetworkInterfaceMigrationPhaseCommitted = NetworkInterfaceMigrationPhase("Committed")
)

func (NetworkInterfaceMigrationPhase) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceMigrationPhase)(nil)).Elem()
}

func (e NetworkInterfaceMigrationPhase) ToNetworkInterfaceMigrationPhaseOutput() NetworkInterfaceMigrationPhaseOutput {
	return pulumi.ToOutput(e).(NetworkInterfaceMigrationPhaseOutput)
}

func (e NetworkInterfaceMigrationPhase) ToNetworkInterfaceMigrationPhaseOutputWithContext(ctx context.Context) NetworkInterfaceMigrationPhaseOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkInterfaceMigrationPhaseOutput)
}

func (e NetworkInterfaceMigrationPhase) ToNetworkInterfaceMigrationPhasePtrOutput() NetworkInterfaceMigrationPhasePtrOutput {
	return e.ToNetworkInterfaceMigrationPhasePtrOutputWithContext(context.Background())
}

func (e NetworkInterfaceMigrationPhase) ToNetworkInterfaceMigrationPhasePtrOutputWithContext(ctx context.Context) NetworkInterfaceMigrationPhasePtrOutput {
	return NetworkInterfaceMigrationPhase(e).ToNetworkInterfaceMigrationPhaseOutputWithContext(ctx).ToNetworkInterfaceMigrationPhasePtrOutputWithContext(ctx)
}

func (e NetworkInterfaceMigrationPhase) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkInterfaceMigrationPhase) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkInterfaceMigrationPhase) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkInterfaceMigrationPhase) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkInterfaceMigrationPhaseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceMigrationPhaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceMigrationPhase)(nil)).Elem()
}

func (o NetworkInterfaceMigrationPhaseOutput) ToNetworkInterfaceMigrationPhaseOutput() NetworkInterfaceMigrationPhaseOutput {
	return o
}

func (o NetworkInterfaceMigrationPhaseOutput) ToNetworkInterfaceMigrationPhaseOutputWithContext(ctx context.Context) NetworkInterfaceMigrationPhaseOutput {
	return o
}

func (o NetworkInterfaceMigrationPhaseOutput) ToNetworkInterfaceMigrationPhasePtrOutput() NetworkInterfaceMigrationPhasePtrOutput {
	return o.ToNetworkInterfaceMigrationPhasePtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceMigrationPhaseOutput) ToNetworkInterfaceMigrationPhasePtrOutputWithContext(ctx context.Context) NetworkInterfaceMigrationPhasePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceMigrationPhase) *NetworkInterfaceMigrationPhase {
		return &v
	}).(NetworkInterfaceMigrationPhasePtrOutput)
}

func (o NetworkInterfaceMigrationPhaseOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkInterfaceMigrationPhaseOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkInterfaceMigrationPhase) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkInterfaceMigrationPhaseOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceMigrationPhaseOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkInterfaceMigrationPhase) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceMigrationPhasePtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceMigrationPhasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceMigrationPhase)(nil)).Elem()
}

func (o NetworkInterfaceMigrationPhasePtrOutput) ToNetworkInterfaceMigrationPhasePtrOutput() NetworkInterfaceMigrationPhasePtrOutput {
	return o
}

func (o NetworkInterfaceMigrationPhasePtrOutput) ToNetworkInterfaceMigrationPhasePtrOutputWithContext(ctx context.Context) NetworkInterfaceMigrationPhasePtrOutput {
	return o
}

func (o NetworkInterfaceMigrationPhasePtrOutput) Elem() NetworkInterfaceMigrationPhaseOutput {
	return o.ApplyT(func(v *NetworkInterfaceMigrationPhase) NetworkInterfaceMigrationPhase {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceMigrationPhase
		return ret
	}).(NetworkInterfaceMigrationPhaseOutput)
}

func (o NetworkInterfaceMigrationPhasePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceMigrationPhasePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkInterfaceMigrationPhase) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkInterfaceMigrationPhaseInput interface {
	pulumi.Input

	ToNetworkInterfaceMigrationPhaseOutput() NetworkInterfaceMigrationPhaseOutput
	ToNetworkInterfaceMigrationPhaseOutputWithContext(context.Context) NetworkInterfaceMigrationPhaseOutput
}

var networkInterfaceMigrationPhasePtrType = reflect.TypeOf((**NetworkInterfaceMigrationPhase)(nil)).Elem()

type NetworkInterfaceMigrationPhasePtrInput interface {
	pulumi.Input

	ToNetworkInterfaceMigrationPhasePtrOutput() NetworkInterfaceMigrationPhasePtrOutput
	ToNetworkInterfaceMigrationPhasePtrOutputWithContext(context.Context) NetworkInterfaceMigrationPhasePtrOutput
}

type networkInterfaceMigrationPhasePtr string

func NetworkInterfaceMigrationPhasePtr(v string) NetworkInterfaceMigrationPhasePtrInput {
	return (*networkInterfaceMigrationPhasePtr)(&v)
}

func (*networkInterfaceMigrationPhasePtr) ElementType() reflect.Type {
	return networkInterfaceMigrationPhasePtrType
}

func (in *networkInterfaceMigrationPhasePtr) ToNetworkInterfaceMigrationPhasePtrOutput() NetworkInterfaceMigrationPhasePtrOutput {
	return pulumi.ToOutput(in).(NetworkInterfaceMigrationPhasePtrOutput)
}

func (in *networkInterfaceMigrationPhasePtr) ToNetworkInterfaceMigrationPhasePtrOutputWithContext(ctx context.Context) NetworkInterfaceMigrationPhasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkInterfaceMigrationPhasePtrOutput)
}

type NetworkInterfaceNicType string

const (
	NetworkInterfaceNicTypeStandard = NetworkInterfaceNicType("Standard")
	NetworkInterfaceNicTypeElastic  = NetworkInterfaceNicType("Elastic")
)

func (NetworkInterfaceNicType) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceNicType)(nil)).Elem()
}

func (e NetworkInterfaceNicType) ToNetworkInterfaceNicTypeOutput() NetworkInterfaceNicTypeOutput {
	return pulumi.ToOutput(e).(NetworkInterfaceNicTypeOutput)
}

func (e NetworkInterfaceNicType) ToNetworkInterfaceNicTypeOutputWithContext(ctx context.Context) NetworkInterfaceNicTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkInterfaceNicTypeOutput)
}

func (e NetworkInterfaceNicType) ToNetworkInterfaceNicTypePtrOutput() NetworkInterfaceNicTypePtrOutput {
	return e.ToNetworkInterfaceNicTypePtrOutputWithContext(context.Background())
}

func (e NetworkInterfaceNicType) ToNetworkInterfaceNicTypePtrOutputWithContext(ctx context.Context) NetworkInterfaceNicTypePtrOutput {
	return NetworkInterfaceNicType(e).ToNetworkInterfaceNicTypeOutputWithContext(ctx).ToNetworkInterfaceNicTypePtrOutputWithContext(ctx)
}

func (e NetworkInterfaceNicType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkInterfaceNicType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkInterfaceNicType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkInterfaceNicType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkInterfaceNicTypeOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceNicTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceNicType)(nil)).Elem()
}

func (o NetworkInterfaceNicTypeOutput) ToNetworkInterfaceNicTypeOutput() NetworkInterfaceNicTypeOutput {
	return o
}

func (o NetworkInterfaceNicTypeOutput) ToNetworkInterfaceNicTypeOutputWithContext(ctx context.Context) NetworkInterfaceNicTypeOutput {
	return o
}

func (o NetworkInterfaceNicTypeOutput) ToNetworkInterfaceNicTypePtrOutput() NetworkInterfaceNicTypePtrOutput {
	return o.ToNetworkInterfaceNicTypePtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceNicTypeOutput) ToNetworkInterfaceNicTypePtrOutputWithContext(ctx context.Context) NetworkInterfaceNicTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceNicType) *NetworkInterfaceNicType {
		return &v
	}).(NetworkInterfaceNicTypePtrOutput)
}

func (o NetworkInterfaceNicTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkInterfaceNicTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkInterfaceNicType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkInterfaceNicTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceNicTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkInterfaceNicType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceNicTypePtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceNicTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceNicType)(nil)).Elem()
}

func (o NetworkInterfaceNicTypePtrOutput) ToNetworkInterfaceNicTypePtrOutput() NetworkInterfaceNicTypePtrOutput {
	return o
}

func (o NetworkInterfaceNicTypePtrOutput) ToNetworkInterfaceNicTypePtrOutputWithContext(ctx context.Context) NetworkInterfaceNicTypePtrOutput {
	return o
}

func (o NetworkInterfaceNicTypePtrOutput) Elem() NetworkInterfaceNicTypeOutput {
	return o.ApplyT(func(v *NetworkInterfaceNicType) NetworkInterfaceNicType {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceNicType
		return ret
	}).(NetworkInterfaceNicTypeOutput)
}

func (o NetworkInterfaceNicTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceNicTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkInterfaceNicType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkInterfaceNicTypeInput interface {
	pulumi.Input

	ToNetworkInterfaceNicTypeOutput() NetworkInterfaceNicTypeOutput
	ToNetworkInterfaceNicTypeOutputWithContext(context.Context) NetworkInterfaceNicTypeOutput
}

var networkInterfaceNicTypePtrType = reflect.TypeOf((**NetworkInterfaceNicType)(nil)).Elem()

type NetworkInterfaceNicTypePtrInput interface {
	pulumi.Input

	ToNetworkInterfaceNicTypePtrOutput() NetworkInterfaceNicTypePtrOutput
	ToNetworkInterfaceNicTypePtrOutputWithContext(context.Context) NetworkInterfaceNicTypePtrOutput
}

type networkInterfaceNicTypePtr string

func NetworkInterfaceNicTypePtr(v string) NetworkInterfaceNicTypePtrInput {
	return (*networkInterfaceNicTypePtr)(&v)
}

func (*networkInterfaceNicTypePtr) ElementType() reflect.Type {
	return networkInterfaceNicTypePtrType
}

func (in *networkInterfaceNicTypePtr) ToNetworkInterfaceNicTypePtrOutput() NetworkInterfaceNicTypePtrOutput {
	return pulumi.ToOutput(in).(NetworkInterfaceNicTypePtrOutput)
}

func (in *networkInterfaceNicTypePtr) ToNetworkInterfaceNicTypePtrOutputWithContext(ctx context.Context) NetworkInterfaceNicTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkInterfaceNicTypePtrOutput)
}

type OutputType string

const (
	OutputTypeWorkspace = OutputType("Workspace")
)

func (OutputType) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (e OutputType) ToOutputTypeOutput() OutputTypeOutput {
	return pulumi.ToOutput(e).(OutputTypeOutput)
}

func (e OutputType) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OutputTypeOutput)
}

func (e OutputType) ToOutputTypePtrOutput() OutputTypePtrOutput {
	return e.ToOutputTypePtrOutputWithContext(context.Background())
}

func (e OutputType) ToOutputTypePtrOutputWithContext(ctx context.Context) OutputTypePtrOutput {
	return OutputType(e).ToOutputTypeOutputWithContext(ctx).ToOutputTypePtrOutputWithContext(ctx)
}

func (e OutputType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutputType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutputType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OutputType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OutputTypeOutput struct{ *pulumi.OutputState }

func (OutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (o OutputTypeOutput) ToOutputTypeOutput() OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) ToOutputTypePtrOutput() OutputTypePtrOutput {
	return o.ToOutputTypePtrOutputWithContext(context.Background())
}

func (o OutputTypeOutput) ToOutputTypePtrOutputWithContext(ctx context.Context) OutputTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutputType) *OutputType {
		return &v
	}).(OutputTypePtrOutput)
}

func (o OutputTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OutputTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OutputTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OutputTypePtrOutput struct{ *pulumi.OutputState }

func (OutputTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputType)(nil)).Elem()
}

func (o OutputTypePtrOutput) ToOutputTypePtrOutput() OutputTypePtrOutput {
	return o
}

func (o OutputTypePtrOutput) ToOutputTypePtrOutputWithContext(ctx context.Context) OutputTypePtrOutput {
	return o
}

func (o OutputTypePtrOutput) Elem() OutputTypeOutput {
	return o.ApplyT(func(v *OutputType) OutputType {
		if v != nil {
			return *v
		}
		var ret OutputType
		return ret
	}).(OutputTypeOutput)
}

func (o OutputTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OutputType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OutputTypeInput interface {
	pulumi.Input

	ToOutputTypeOutput() OutputTypeOutput
	ToOutputTypeOutputWithContext(context.Context) OutputTypeOutput
}

var outputTypePtrType = reflect.TypeOf((**OutputType)(nil)).Elem()

type OutputTypePtrInput interface {
	pulumi.Input

	ToOutputTypePtrOutput() OutputTypePtrOutput
	ToOutputTypePtrOutputWithContext(context.Context) OutputTypePtrOutput
}

type outputTypePtr string

func OutputTypePtr(v string) OutputTypePtrInput {
	return (*outputTypePtr)(&v)
}

func (*outputTypePtr) ElementType() reflect.Type {
	return outputTypePtrType
}

func (in *outputTypePtr) ToOutputTypePtrOutput() OutputTypePtrOutput {
	return pulumi.ToOutput(in).(OutputTypePtrOutput)
}

func (in *outputTypePtr) ToOutputTypePtrOutputWithContext(ctx context.Context) OutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OutputTypePtrOutput)
}

type OwaspCrsExclusionEntryMatchVariable string

const (
	OwaspCrsExclusionEntryMatchVariableRequestHeaderNames = OwaspCrsExclusionEntryMatchVariable("RequestHeaderNames")
	OwaspCrsExclusionEntryMatchVariableRequestCookieNames = OwaspCrsExclusionEntryMatchVariable("RequestCookieNames")
	OwaspCrsExclusionEntryMatchVariableRequestArgNames    = OwaspCrsExclusionEntryMatchVariable("RequestArgNames")
)

func (OwaspCrsExclusionEntryMatchVariable) ElementType() reflect.Type {
	return reflect.TypeOf((*OwaspCrsExclusionEntryMatchVariable)(nil)).Elem()
}

func (e OwaspCrsExclusionEntryMatchVariable) ToOwaspCrsExclusionEntryMatchVariableOutput() OwaspCrsExclusionEntryMatchVariableOutput {
	return pulumi.ToOutput(e).(OwaspCrsExclusionEntryMatchVariableOutput)
}

func (e OwaspCrsExclusionEntryMatchVariable) ToOwaspCrsExclusionEntryMatchVariableOutputWithContext(ctx context.Context) OwaspCrsExclusionEntryMatchVariableOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OwaspCrsExclusionEntryMatchVariableOutput)
}

func (e OwaspCrsExclusionEntryMatchVariable) ToOwaspCrsExclusionEntryMatchVariablePtrOutput() OwaspCrsExclusionEntryMatchVariablePtrOutput {
	return e.ToOwaspCrsExclusionEntryMatchVariablePtrOutputWithContext(context.Background())
}

func (e OwaspCrsExclusionEntryMatchVariable) ToOwaspCrsExclusionEntryMatchVariablePtrOutputWithContext(ctx context.Context) OwaspCrsExclusionEntryMatchVariablePtrOutput {
	return OwaspCrsExclusionEntryMatchVariable(e).ToOwaspCrsExclusionEntryMatchVariableOutputWithContext(ctx).ToOwaspCrsExclusionEntryMatchVariablePtrOutputWithContext(ctx)
}

func (e OwaspCrsExclusionEntryMatchVariable) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OwaspCrsExclusionEntryMatchVariable) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OwaspCrsExclusionEntryMatchVariable) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OwaspCrsExclusionEntryMatchVariable) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OwaspCrsExclusionEntryMatchVariableOutput struct{ *pulumi.OutputState }

func (OwaspCrsExclusionEntryMatchVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OwaspCrsExclusionEntryMatchVariable)(nil)).Elem()
}

func (o OwaspCrsExclusionEntryMatchVariableOutput) ToOwaspCrsExclusionEntryMatchVariableOutput() OwaspCrsExclusionEntryMatchVariableOutput {
	return o
}

func (o OwaspCrsExclusionEntryMatchVariableOutput) ToOwaspCrsExclusionEntryMatchVariableOutputWithContext(ctx context.Context) OwaspCrsExclusionEntryMatchVariableOutput {
	return o
}

func (o OwaspCrsExclusionEntryMatchVariableOutput) ToOwaspCrsExclusionEntryMatchVariablePtrOutput() OwaspCrsExclusionEntryMatchVariablePtrOutput {
	return o.ToOwaspCrsExclusionEntryMatchVariablePtrOutputWithContext(context.Background())
}

func (o OwaspCrsExclusionEntryMatchVariableOutput) ToOwaspCrsExclusionEntryMatchVariablePtrOutputWithContext(ctx context.Context) OwaspCrsExclusionEntryMatchVariablePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OwaspCrsExclusionEntryMatchVariable) *OwaspCrsExclusionEntryMatchVariable {
		return &v
	}).(OwaspCrsExclusionEntryMatchVariablePtrOutput)
}

func (o OwaspCrsExclusionEntryMatchVariableOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OwaspCrsExclusionEntryMatchVariableOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OwaspCrsExclusionEntryMatchVariable) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OwaspCrsExclusionEntryMatchVariableOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OwaspCrsExclusionEntryMatchVariableOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OwaspCrsExclusionEntryMatchVariable) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OwaspCrsExclusionEntryMatchVariablePtrOutput struct{ *pulumi.OutputState }

func (OwaspCrsExclusionEntryMatchVariablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OwaspCrsExclusionEntryMatchVariable)(nil)).Elem()
}

func (o OwaspCrsExclusionEntryMatchVariablePtrOutput) ToOwaspCrsExclusionEntryMatchVariablePtrOutput() OwaspCrsExclusionEntryMatchVariablePtrOutput {
	return o
}

func (o OwaspCrsExclusionEntryMatchVariablePtrOutput) ToOwaspCrsExclusionEntryMatchVariablePtrOutputWithContext(ctx context.Context) OwaspCrsExclusionEntryMatchVariablePtrOutput {
	return o
}

func (o OwaspCrsExclusionEntryMatchVariablePtrOutput) Elem() OwaspCrsExclusionEntryMatchVariableOutput {
	return o.ApplyT(func(v *OwaspCrsExclusionEntryMatchVariable) OwaspCrsExclusionEntryMatchVariable {
		if v != nil {
			return *v
		}
		var ret OwaspCrsExclusionEntryMatchVariable
		return ret
	}).(OwaspCrsExclusionEntryMatchVariableOutput)
}

func (o OwaspCrsExclusionEntryMatchVariablePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OwaspCrsExclusionEntryMatchVariablePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OwaspCrsExclusionEntryMatchVariable) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OwaspCrsExclusionEntryMatchVariableInput interface {
	pulumi.Input

	ToOwaspCrsExclusionEntryMatchVariableOutput() OwaspCrsExclusionEntryMatchVariableOutput
	ToOwaspCrsExclusionEntryMatchVariableOutputWithContext(context.Context) OwaspCrsExclusionEntryMatchVariableOutput
}

var owaspCrsExclusionEntryMatchVariablePtrType = reflect.TypeOf((**OwaspCrsExclusionEntryMatchVariable)(nil)).Elem()

type OwaspCrsExclusionEntryMatchVariablePtrInput interface {
	pulumi.Input

	ToOwaspCrsExclusionEntryMatchVariablePtrOutput() OwaspCrsExclusionEntryMatchVariablePtrOutput
	ToOwaspCrsExclusionEntryMatchVariablePtrOutputWithContext(context.Context) OwaspCrsExclusionEntryMatchVariablePtrOutput
}

type owaspCrsExclusionEntryMatchVariablePtr string

func OwaspCrsExclusionEntryMatchVariablePtr(v string) OwaspCrsExclusionEntryMatchVariablePtrInput {
	return (*owaspCrsExclusionEntryMatchVariablePtr)(&v)
}

func (*owaspCrsExclusionEntryMatchVariablePtr) ElementType() reflect.Type {
	return owaspCrsExclusionEntryMatchVariablePtrType
}

func (in *owaspCrsExclusionEntryMatchVariablePtr) ToOwaspCrsExclusionEntryMatchVariablePtrOutput() OwaspCrsExclusionEntryMatchVariablePtrOutput {
	return pulumi.ToOutput(in).(OwaspCrsExclusionEntryMatchVariablePtrOutput)
}

func (in *owaspCrsExclusionEntryMatchVariablePtr) ToOwaspCrsExclusionEntryMatchVariablePtrOutputWithContext(ctx context.Context) OwaspCrsExclusionEntryMatchVariablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OwaspCrsExclusionEntryMatchVariablePtrOutput)
}

type OwaspCrsExclusionEntrySelectorMatchOperator string

const (
	OwaspCrsExclusionEntrySelectorMatchOperatorEquals     = OwaspCrsExclusionEntrySelectorMatchOperator("Equals")
	OwaspCrsExclusionEntrySelectorMatchOperatorContains   = OwaspCrsExclusionEntrySelectorMatchOperator("Contains")
	OwaspCrsExclusionEntrySelectorMatchOperatorStartsWith = OwaspCrsExclusionEntrySelectorMatchOperator("StartsWith")
	OwaspCrsExclusionEntrySelectorMatchOperatorEndsWith   = OwaspCrsExclusionEntrySelectorMatchOperator("EndsWith")
	OwaspCrsExclusionEntrySelectorMatchOperatorEqualsAny  = OwaspCrsExclusionEntrySelectorMatchOperator("EqualsAny")
)

func (OwaspCrsExclusionEntrySelectorMatchOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*OwaspCrsExclusionEntrySelectorMatchOperator)(nil)).Elem()
}

func (e OwaspCrsExclusionEntrySelectorMatchOperator) ToOwaspCrsExclusionEntrySelectorMatchOperatorOutput() OwaspCrsExclusionEntrySelectorMatchOperatorOutput {
	return pulumi.ToOutput(e).(OwaspCrsExclusionEntrySelectorMatchOperatorOutput)
}

func (e OwaspCrsExclusionEntrySelectorMatchOperator) ToOwaspCrsExclusionEntrySelectorMatchOperatorOutputWithContext(ctx context.Context) OwaspCrsExclusionEntrySelectorMatchOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OwaspCrsExclusionEntrySelectorMatchOperatorOutput)
}

func (e OwaspCrsExclusionEntrySelectorMatchOperator) ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput() OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput {
	return e.ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutputWithContext(context.Background())
}

func (e OwaspCrsExclusionEntrySelectorMatchOperator) ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutputWithContext(ctx context.Context) OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput {
	return OwaspCrsExclusionEntrySelectorMatchOperator(e).ToOwaspCrsExclusionEntrySelectorMatchOperatorOutputWithContext(ctx).ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutputWithContext(ctx)
}

func (e OwaspCrsExclusionEntrySelectorMatchOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OwaspCrsExclusionEntrySelectorMatchOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OwaspCrsExclusionEntrySelectorMatchOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OwaspCrsExclusionEntrySelectorMatchOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OwaspCrsExclusionEntrySelectorMatchOperatorOutput struct{ *pulumi.OutputState }

func (OwaspCrsExclusionEntrySelectorMatchOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OwaspCrsExclusionEntrySelectorMatchOperator)(nil)).Elem()
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorOutput) ToOwaspCrsExclusionEntrySelectorMatchOperatorOutput() OwaspCrsExclusionEntrySelectorMatchOperatorOutput {
	return o
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorOutput) ToOwaspCrsExclusionEntrySelectorMatchOperatorOutputWithContext(ctx context.Context) OwaspCrsExclusionEntrySelectorMatchOperatorOutput {
	return o
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorOutput) ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput() OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput {
	return o.ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutputWithContext(context.Background())
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorOutput) ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutputWithContext(ctx context.Context) OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OwaspCrsExclusionEntrySelectorMatchOperator) *OwaspCrsExclusionEntrySelectorMatchOperator {
		return &v
	}).(OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput)
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OwaspCrsExclusionEntrySelectorMatchOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OwaspCrsExclusionEntrySelectorMatchOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput struct{ *pulumi.OutputState }

func (OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OwaspCrsExclusionEntrySelectorMatchOperator)(nil)).Elem()
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput) ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput() OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput {
	return o
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput) ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutputWithContext(ctx context.Context) OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput {
	return o
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput) Elem() OwaspCrsExclusionEntrySelectorMatchOperatorOutput {
	return o.ApplyT(func(v *OwaspCrsExclusionEntrySelectorMatchOperator) OwaspCrsExclusionEntrySelectorMatchOperator {
		if v != nil {
			return *v
		}
		var ret OwaspCrsExclusionEntrySelectorMatchOperator
		return ret
	}).(OwaspCrsExclusionEntrySelectorMatchOperatorOutput)
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OwaspCrsExclusionEntrySelectorMatchOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OwaspCrsExclusionEntrySelectorMatchOperatorInput interface {
	pulumi.Input

	ToOwaspCrsExclusionEntrySelectorMatchOperatorOutput() OwaspCrsExclusionEntrySelectorMatchOperatorOutput
	ToOwaspCrsExclusionEntrySelectorMatchOperatorOutputWithContext(context.Context) OwaspCrsExclusionEntrySelectorMatchOperatorOutput
}

var owaspCrsExclusionEntrySelectorMatchOperatorPtrType = reflect.TypeOf((**OwaspCrsExclusionEntrySelectorMatchOperator)(nil)).Elem()

type OwaspCrsExclusionEntrySelectorMatchOperatorPtrInput interface {
	pulumi.Input

	ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput() OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput
	ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutputWithContext(context.Context) OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput
}

type owaspCrsExclusionEntrySelectorMatchOperatorPtr string

func OwaspCrsExclusionEntrySelectorMatchOperatorPtr(v string) OwaspCrsExclusionEntrySelectorMatchOperatorPtrInput {
	return (*owaspCrsExclusionEntrySelectorMatchOperatorPtr)(&v)
}

func (*owaspCrsExclusionEntrySelectorMatchOperatorPtr) ElementType() reflect.Type {
	return owaspCrsExclusionEntrySelectorMatchOperatorPtrType
}

func (in *owaspCrsExclusionEntrySelectorMatchOperatorPtr) ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput() OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput {
	return pulumi.ToOutput(in).(OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput)
}

func (in *owaspCrsExclusionEntrySelectorMatchOperatorPtr) ToOwaspCrsExclusionEntrySelectorMatchOperatorPtrOutputWithContext(ctx context.Context) OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput)
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
	PfsGroupPFS14   = PfsGroup("PFS14")
	PfsGroupPFSMM   = PfsGroup("PFSMM")
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

type PreferredIPVersion string

const (
	PreferredIPVersionIPv4 = PreferredIPVersion("IPv4")
	PreferredIPVersionIPv6 = PreferredIPVersion("IPv6")
)

func (PreferredIPVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferredIPVersion)(nil)).Elem()
}

func (e PreferredIPVersion) ToPreferredIPVersionOutput() PreferredIPVersionOutput {
	return pulumi.ToOutput(e).(PreferredIPVersionOutput)
}

func (e PreferredIPVersion) ToPreferredIPVersionOutputWithContext(ctx context.Context) PreferredIPVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PreferredIPVersionOutput)
}

func (e PreferredIPVersion) ToPreferredIPVersionPtrOutput() PreferredIPVersionPtrOutput {
	return e.ToPreferredIPVersionPtrOutputWithContext(context.Background())
}

func (e PreferredIPVersion) ToPreferredIPVersionPtrOutputWithContext(ctx context.Context) PreferredIPVersionPtrOutput {
	return PreferredIPVersion(e).ToPreferredIPVersionOutputWithContext(ctx).ToPreferredIPVersionPtrOutputWithContext(ctx)
}

func (e PreferredIPVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PreferredIPVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PreferredIPVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PreferredIPVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PreferredIPVersionOutput struct{ *pulumi.OutputState }

func (PreferredIPVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferredIPVersion)(nil)).Elem()
}

func (o PreferredIPVersionOutput) ToPreferredIPVersionOutput() PreferredIPVersionOutput {
	return o
}

func (o PreferredIPVersionOutput) ToPreferredIPVersionOutputWithContext(ctx context.Context) PreferredIPVersionOutput {
	return o
}

func (o PreferredIPVersionOutput) ToPreferredIPVersionPtrOutput() PreferredIPVersionPtrOutput {
	return o.ToPreferredIPVersionPtrOutputWithContext(context.Background())
}

func (o PreferredIPVersionOutput) ToPreferredIPVersionPtrOutputWithContext(ctx context.Context) PreferredIPVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PreferredIPVersion) *PreferredIPVersion {
		return &v
	}).(PreferredIPVersionPtrOutput)
}

func (o PreferredIPVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PreferredIPVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PreferredIPVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PreferredIPVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PreferredIPVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PreferredIPVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PreferredIPVersionPtrOutput struct{ *pulumi.OutputState }

func (PreferredIPVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreferredIPVersion)(nil)).Elem()
}

func (o PreferredIPVersionPtrOutput) ToPreferredIPVersionPtrOutput() PreferredIPVersionPtrOutput {
	return o
}

func (o PreferredIPVersionPtrOutput) ToPreferredIPVersionPtrOutputWithContext(ctx context.Context) PreferredIPVersionPtrOutput {
	return o
}

func (o PreferredIPVersionPtrOutput) Elem() PreferredIPVersionOutput {
	return o.ApplyT(func(v *PreferredIPVersion) PreferredIPVersion {
		if v != nil {
			return *v
		}
		var ret PreferredIPVersion
		return ret
	}).(PreferredIPVersionOutput)
}

func (o PreferredIPVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PreferredIPVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PreferredIPVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PreferredIPVersionInput interface {
	pulumi.Input

	ToPreferredIPVersionOutput() PreferredIPVersionOutput
	ToPreferredIPVersionOutputWithContext(context.Context) PreferredIPVersionOutput
}

var preferredIPVersionPtrType = reflect.TypeOf((**PreferredIPVersion)(nil)).Elem()

type PreferredIPVersionPtrInput interface {
	pulumi.Input

	ToPreferredIPVersionPtrOutput() PreferredIPVersionPtrOutput
	ToPreferredIPVersionPtrOutputWithContext(context.Context) PreferredIPVersionPtrOutput
}

type preferredIPVersionPtr string

func PreferredIPVersionPtr(v string) PreferredIPVersionPtrInput {
	return (*preferredIPVersionPtr)(&v)
}

func (*preferredIPVersionPtr) ElementType() reflect.Type {
	return preferredIPVersionPtrType
}

func (in *preferredIPVersionPtr) ToPreferredIPVersionPtrOutput() PreferredIPVersionPtrOutput {
	return pulumi.ToOutput(in).(PreferredIPVersionPtrOutput)
}

func (in *preferredIPVersionPtr) ToPreferredIPVersionPtrOutputWithContext(ctx context.Context) PreferredIPVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PreferredIPVersionPtrOutput)
}

type PreferredRoutingGateway string

const (
	PreferredRoutingGatewayExpressRoute = PreferredRoutingGateway("ExpressRoute")
	PreferredRoutingGatewayVpnGateway   = PreferredRoutingGateway("VpnGateway")
	PreferredRoutingGatewayNone         = PreferredRoutingGateway("None")
)

func (PreferredRoutingGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferredRoutingGateway)(nil)).Elem()
}

func (e PreferredRoutingGateway) ToPreferredRoutingGatewayOutput() PreferredRoutingGatewayOutput {
	return pulumi.ToOutput(e).(PreferredRoutingGatewayOutput)
}

func (e PreferredRoutingGateway) ToPreferredRoutingGatewayOutputWithContext(ctx context.Context) PreferredRoutingGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PreferredRoutingGatewayOutput)
}

func (e PreferredRoutingGateway) ToPreferredRoutingGatewayPtrOutput() PreferredRoutingGatewayPtrOutput {
	return e.ToPreferredRoutingGatewayPtrOutputWithContext(context.Background())
}

func (e PreferredRoutingGateway) ToPreferredRoutingGatewayPtrOutputWithContext(ctx context.Context) PreferredRoutingGatewayPtrOutput {
	return PreferredRoutingGateway(e).ToPreferredRoutingGatewayOutputWithContext(ctx).ToPreferredRoutingGatewayPtrOutputWithContext(ctx)
}

func (e PreferredRoutingGateway) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PreferredRoutingGateway) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PreferredRoutingGateway) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PreferredRoutingGateway) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PreferredRoutingGatewayOutput struct{ *pulumi.OutputState }

func (PreferredRoutingGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferredRoutingGateway)(nil)).Elem()
}

func (o PreferredRoutingGatewayOutput) ToPreferredRoutingGatewayOutput() PreferredRoutingGatewayOutput {
	return o
}

func (o PreferredRoutingGatewayOutput) ToPreferredRoutingGatewayOutputWithContext(ctx context.Context) PreferredRoutingGatewayOutput {
	return o
}

func (o PreferredRoutingGatewayOutput) ToPreferredRoutingGatewayPtrOutput() PreferredRoutingGatewayPtrOutput {
	return o.ToPreferredRoutingGatewayPtrOutputWithContext(context.Background())
}

func (o PreferredRoutingGatewayOutput) ToPreferredRoutingGatewayPtrOutputWithContext(ctx context.Context) PreferredRoutingGatewayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PreferredRoutingGateway) *PreferredRoutingGateway {
		return &v
	}).(PreferredRoutingGatewayPtrOutput)
}

func (o PreferredRoutingGatewayOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PreferredRoutingGatewayOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PreferredRoutingGateway) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PreferredRoutingGatewayOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PreferredRoutingGatewayOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PreferredRoutingGateway) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PreferredRoutingGatewayPtrOutput struct{ *pulumi.OutputState }

func (PreferredRoutingGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreferredRoutingGateway)(nil)).Elem()
}

func (o PreferredRoutingGatewayPtrOutput) ToPreferredRoutingGatewayPtrOutput() PreferredRoutingGatewayPtrOutput {
	return o
}

func (o PreferredRoutingGatewayPtrOutput) ToPreferredRoutingGatewayPtrOutputWithContext(ctx context.Context) PreferredRoutingGatewayPtrOutput {
	return o
}

func (o PreferredRoutingGatewayPtrOutput) Elem() PreferredRoutingGatewayOutput {
	return o.ApplyT(func(v *PreferredRoutingGateway) PreferredRoutingGateway {
		if v != nil {
			return *v
		}
		var ret PreferredRoutingGateway
		return ret
	}).(PreferredRoutingGatewayOutput)
}

func (o PreferredRoutingGatewayPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PreferredRoutingGatewayPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PreferredRoutingGateway) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PreferredRoutingGatewayInput interface {
	pulumi.Input

	ToPreferredRoutingGatewayOutput() PreferredRoutingGatewayOutput
	ToPreferredRoutingGatewayOutputWithContext(context.Context) PreferredRoutingGatewayOutput
}

var preferredRoutingGatewayPtrType = reflect.TypeOf((**PreferredRoutingGateway)(nil)).Elem()

type PreferredRoutingGatewayPtrInput interface {
	pulumi.Input

	ToPreferredRoutingGatewayPtrOutput() PreferredRoutingGatewayPtrOutput
	ToPreferredRoutingGatewayPtrOutputWithContext(context.Context) PreferredRoutingGatewayPtrOutput
}

type preferredRoutingGatewayPtr string

func PreferredRoutingGatewayPtr(v string) PreferredRoutingGatewayPtrInput {
	return (*preferredRoutingGatewayPtr)(&v)
}

func (*preferredRoutingGatewayPtr) ElementType() reflect.Type {
	return preferredRoutingGatewayPtrType
}

func (in *preferredRoutingGatewayPtr) ToPreferredRoutingGatewayPtrOutput() PreferredRoutingGatewayPtrOutput {
	return pulumi.ToOutput(in).(PreferredRoutingGatewayPtrOutput)
}

func (in *preferredRoutingGatewayPtr) ToPreferredRoutingGatewayPtrOutputWithContext(ctx context.Context) PreferredRoutingGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PreferredRoutingGatewayPtrOutput)
}

type ProbeProtocol string

const (
	ProbeProtocolHttp  = ProbeProtocol("Http")
	ProbeProtocolTcp   = ProbeProtocol("Tcp")
	ProbeProtocolHttps = ProbeProtocol("Https")
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

type ProtocolType string

const (
	ProtocolTypeDoNotUse = ProtocolType("DoNotUse")
	ProtocolTypeIcmp     = ProtocolType("Icmp")
	ProtocolTypeTcp      = ProtocolType("Tcp")
	ProtocolTypeUdp      = ProtocolType("Udp")
	ProtocolTypeGre      = ProtocolType("Gre")
	ProtocolTypeEsp      = ProtocolType("Esp")
	ProtocolTypeAh       = ProtocolType("Ah")
	ProtocolTypeVxlan    = ProtocolType("Vxlan")
	ProtocolTypeAll      = ProtocolType("All")
)

func (ProtocolType) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtocolType)(nil)).Elem()
}

func (e ProtocolType) ToProtocolTypeOutput() ProtocolTypeOutput {
	return pulumi.ToOutput(e).(ProtocolTypeOutput)
}

func (e ProtocolType) ToProtocolTypeOutputWithContext(ctx context.Context) ProtocolTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtocolTypeOutput)
}

func (e ProtocolType) ToProtocolTypePtrOutput() ProtocolTypePtrOutput {
	return e.ToProtocolTypePtrOutputWithContext(context.Background())
}

func (e ProtocolType) ToProtocolTypePtrOutputWithContext(ctx context.Context) ProtocolTypePtrOutput {
	return ProtocolType(e).ToProtocolTypeOutputWithContext(ctx).ToProtocolTypePtrOutputWithContext(ctx)
}

func (e ProtocolType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtocolType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtocolType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProtocolType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtocolTypeOutput struct{ *pulumi.OutputState }

func (ProtocolTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtocolType)(nil)).Elem()
}

func (o ProtocolTypeOutput) ToProtocolTypeOutput() ProtocolTypeOutput {
	return o
}

func (o ProtocolTypeOutput) ToProtocolTypeOutputWithContext(ctx context.Context) ProtocolTypeOutput {
	return o
}

func (o ProtocolTypeOutput) ToProtocolTypePtrOutput() ProtocolTypePtrOutput {
	return o.ToProtocolTypePtrOutputWithContext(context.Background())
}

func (o ProtocolTypeOutput) ToProtocolTypePtrOutputWithContext(ctx context.Context) ProtocolTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtocolType) *ProtocolType {
		return &v
	}).(ProtocolTypePtrOutput)
}

func (o ProtocolTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtocolTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtocolType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtocolTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtocolType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtocolTypePtrOutput struct{ *pulumi.OutputState }

func (ProtocolTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtocolType)(nil)).Elem()
}

func (o ProtocolTypePtrOutput) ToProtocolTypePtrOutput() ProtocolTypePtrOutput {
	return o
}

func (o ProtocolTypePtrOutput) ToProtocolTypePtrOutputWithContext(ctx context.Context) ProtocolTypePtrOutput {
	return o
}

func (o ProtocolTypePtrOutput) Elem() ProtocolTypeOutput {
	return o.ApplyT(func(v *ProtocolType) ProtocolType {
		if v != nil {
			return *v
		}
		var ret ProtocolType
		return ret
	}).(ProtocolTypeOutput)
}

func (o ProtocolTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProtocolType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtocolTypeInput interface {
	pulumi.Input

	ToProtocolTypeOutput() ProtocolTypeOutput
	ToProtocolTypeOutputWithContext(context.Context) ProtocolTypeOutput
}

var protocolTypePtrType = reflect.TypeOf((**ProtocolType)(nil)).Elem()

type ProtocolTypePtrInput interface {
	pulumi.Input

	ToProtocolTypePtrOutput() ProtocolTypePtrOutput
	ToProtocolTypePtrOutputWithContext(context.Context) ProtocolTypePtrOutput
}

type protocolTypePtr string

func ProtocolTypePtr(v string) ProtocolTypePtrInput {
	return (*protocolTypePtr)(&v)
}

func (*protocolTypePtr) ElementType() reflect.Type {
	return protocolTypePtrType
}

func (in *protocolTypePtr) ToProtocolTypePtrOutput() ProtocolTypePtrOutput {
	return pulumi.ToOutput(in).(ProtocolTypePtrOutput)
}

func (in *protocolTypePtr) ToProtocolTypePtrOutputWithContext(ctx context.Context) ProtocolTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtocolTypePtrOutput)
}

type PublicIPAddressMigrationPhase string

const (
	PublicIPAddressMigrationPhaseNone      = PublicIPAddressMigrationPhase("None")
	PublicIPAddressMigrationPhasePrepare   = PublicIPAddressMigrationPhase("Prepare")
	PublicIPAddressMigrationPhaseCommit    = PublicIPAddressMigrationPhase("Commit")
	PublicIPAddressMigrationPhaseAbort     = PublicIPAddressMigrationPhase("Abort")
	PublicIPAddressMigrationPhaseCommitted = PublicIPAddressMigrationPhase("Committed")
)

func (PublicIPAddressMigrationPhase) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressMigrationPhase)(nil)).Elem()
}

func (e PublicIPAddressMigrationPhase) ToPublicIPAddressMigrationPhaseOutput() PublicIPAddressMigrationPhaseOutput {
	return pulumi.ToOutput(e).(PublicIPAddressMigrationPhaseOutput)
}

func (e PublicIPAddressMigrationPhase) ToPublicIPAddressMigrationPhaseOutputWithContext(ctx context.Context) PublicIPAddressMigrationPhaseOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicIPAddressMigrationPhaseOutput)
}

func (e PublicIPAddressMigrationPhase) ToPublicIPAddressMigrationPhasePtrOutput() PublicIPAddressMigrationPhasePtrOutput {
	return e.ToPublicIPAddressMigrationPhasePtrOutputWithContext(context.Background())
}

func (e PublicIPAddressMigrationPhase) ToPublicIPAddressMigrationPhasePtrOutputWithContext(ctx context.Context) PublicIPAddressMigrationPhasePtrOutput {
	return PublicIPAddressMigrationPhase(e).ToPublicIPAddressMigrationPhaseOutputWithContext(ctx).ToPublicIPAddressMigrationPhasePtrOutputWithContext(ctx)
}

func (e PublicIPAddressMigrationPhase) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPAddressMigrationPhase) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPAddressMigrationPhase) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicIPAddressMigrationPhase) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicIPAddressMigrationPhaseOutput struct{ *pulumi.OutputState }

func (PublicIPAddressMigrationPhaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressMigrationPhase)(nil)).Elem()
}

func (o PublicIPAddressMigrationPhaseOutput) ToPublicIPAddressMigrationPhaseOutput() PublicIPAddressMigrationPhaseOutput {
	return o
}

func (o PublicIPAddressMigrationPhaseOutput) ToPublicIPAddressMigrationPhaseOutputWithContext(ctx context.Context) PublicIPAddressMigrationPhaseOutput {
	return o
}

func (o PublicIPAddressMigrationPhaseOutput) ToPublicIPAddressMigrationPhasePtrOutput() PublicIPAddressMigrationPhasePtrOutput {
	return o.ToPublicIPAddressMigrationPhasePtrOutputWithContext(context.Background())
}

func (o PublicIPAddressMigrationPhaseOutput) ToPublicIPAddressMigrationPhasePtrOutputWithContext(ctx context.Context) PublicIPAddressMigrationPhasePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressMigrationPhase) *PublicIPAddressMigrationPhase {
		return &v
	}).(PublicIPAddressMigrationPhasePtrOutput)
}

func (o PublicIPAddressMigrationPhaseOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicIPAddressMigrationPhaseOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPAddressMigrationPhase) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicIPAddressMigrationPhaseOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressMigrationPhaseOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPAddressMigrationPhase) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicIPAddressMigrationPhasePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressMigrationPhasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressMigrationPhase)(nil)).Elem()
}

func (o PublicIPAddressMigrationPhasePtrOutput) ToPublicIPAddressMigrationPhasePtrOutput() PublicIPAddressMigrationPhasePtrOutput {
	return o
}

func (o PublicIPAddressMigrationPhasePtrOutput) ToPublicIPAddressMigrationPhasePtrOutputWithContext(ctx context.Context) PublicIPAddressMigrationPhasePtrOutput {
	return o
}

func (o PublicIPAddressMigrationPhasePtrOutput) Elem() PublicIPAddressMigrationPhaseOutput {
	return o.ApplyT(func(v *PublicIPAddressMigrationPhase) PublicIPAddressMigrationPhase {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressMigrationPhase
		return ret
	}).(PublicIPAddressMigrationPhaseOutput)
}

func (o PublicIPAddressMigrationPhasePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressMigrationPhasePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicIPAddressMigrationPhase) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicIPAddressMigrationPhaseInput interface {
	pulumi.Input

	ToPublicIPAddressMigrationPhaseOutput() PublicIPAddressMigrationPhaseOutput
	ToPublicIPAddressMigrationPhaseOutputWithContext(context.Context) PublicIPAddressMigrationPhaseOutput
}

var publicIPAddressMigrationPhasePtrType = reflect.TypeOf((**PublicIPAddressMigrationPhase)(nil)).Elem()

type PublicIPAddressMigrationPhasePtrInput interface {
	pulumi.Input

	ToPublicIPAddressMigrationPhasePtrOutput() PublicIPAddressMigrationPhasePtrOutput
	ToPublicIPAddressMigrationPhasePtrOutputWithContext(context.Context) PublicIPAddressMigrationPhasePtrOutput
}

type publicIPAddressMigrationPhasePtr string

func PublicIPAddressMigrationPhasePtr(v string) PublicIPAddressMigrationPhasePtrInput {
	return (*publicIPAddressMigrationPhasePtr)(&v)
}

func (*publicIPAddressMigrationPhasePtr) ElementType() reflect.Type {
	return publicIPAddressMigrationPhasePtrType
}

func (in *publicIPAddressMigrationPhasePtr) ToPublicIPAddressMigrationPhasePtrOutput() PublicIPAddressMigrationPhasePtrOutput {
	return pulumi.ToOutput(in).(PublicIPAddressMigrationPhasePtrOutput)
}

func (in *publicIPAddressMigrationPhasePtr) ToPublicIPAddressMigrationPhasePtrOutputWithContext(ctx context.Context) PublicIPAddressMigrationPhasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicIPAddressMigrationPhasePtrOutput)
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

type PublicIPAddressSkuTier string

const (
	PublicIPAddressSkuTierRegional = PublicIPAddressSkuTier("Regional")
	PublicIPAddressSkuTierGlobal   = PublicIPAddressSkuTier("Global")
)

func (PublicIPAddressSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSkuTier)(nil)).Elem()
}

func (e PublicIPAddressSkuTier) ToPublicIPAddressSkuTierOutput() PublicIPAddressSkuTierOutput {
	return pulumi.ToOutput(e).(PublicIPAddressSkuTierOutput)
}

func (e PublicIPAddressSkuTier) ToPublicIPAddressSkuTierOutputWithContext(ctx context.Context) PublicIPAddressSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicIPAddressSkuTierOutput)
}

func (e PublicIPAddressSkuTier) ToPublicIPAddressSkuTierPtrOutput() PublicIPAddressSkuTierPtrOutput {
	return e.ToPublicIPAddressSkuTierPtrOutputWithContext(context.Background())
}

func (e PublicIPAddressSkuTier) ToPublicIPAddressSkuTierPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuTierPtrOutput {
	return PublicIPAddressSkuTier(e).ToPublicIPAddressSkuTierOutputWithContext(ctx).ToPublicIPAddressSkuTierPtrOutputWithContext(ctx)
}

func (e PublicIPAddressSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPAddressSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPAddressSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicIPAddressSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicIPAddressSkuTierOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSkuTier)(nil)).Elem()
}

func (o PublicIPAddressSkuTierOutput) ToPublicIPAddressSkuTierOutput() PublicIPAddressSkuTierOutput {
	return o
}

func (o PublicIPAddressSkuTierOutput) ToPublicIPAddressSkuTierOutputWithContext(ctx context.Context) PublicIPAddressSkuTierOutput {
	return o
}

func (o PublicIPAddressSkuTierOutput) ToPublicIPAddressSkuTierPtrOutput() PublicIPAddressSkuTierPtrOutput {
	return o.ToPublicIPAddressSkuTierPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuTierOutput) ToPublicIPAddressSkuTierPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressSkuTier) *PublicIPAddressSkuTier {
		return &v
	}).(PublicIPAddressSkuTierPtrOutput)
}

func (o PublicIPAddressSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPAddressSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicIPAddressSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPAddressSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicIPAddressSkuTierPtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressSkuTier)(nil)).Elem()
}

func (o PublicIPAddressSkuTierPtrOutput) ToPublicIPAddressSkuTierPtrOutput() PublicIPAddressSkuTierPtrOutput {
	return o
}

func (o PublicIPAddressSkuTierPtrOutput) ToPublicIPAddressSkuTierPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuTierPtrOutput {
	return o
}

func (o PublicIPAddressSkuTierPtrOutput) Elem() PublicIPAddressSkuTierOutput {
	return o.ApplyT(func(v *PublicIPAddressSkuTier) PublicIPAddressSkuTier {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressSkuTier
		return ret
	}).(PublicIPAddressSkuTierOutput)
}

func (o PublicIPAddressSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicIPAddressSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicIPAddressSkuTierInput interface {
	pulumi.Input

	ToPublicIPAddressSkuTierOutput() PublicIPAddressSkuTierOutput
	ToPublicIPAddressSkuTierOutputWithContext(context.Context) PublicIPAddressSkuTierOutput
}

var publicIPAddressSkuTierPtrType = reflect.TypeOf((**PublicIPAddressSkuTier)(nil)).Elem()

type PublicIPAddressSkuTierPtrInput interface {
	pulumi.Input

	ToPublicIPAddressSkuTierPtrOutput() PublicIPAddressSkuTierPtrOutput
	ToPublicIPAddressSkuTierPtrOutputWithContext(context.Context) PublicIPAddressSkuTierPtrOutput
}

type publicIPAddressSkuTierPtr string

func PublicIPAddressSkuTierPtr(v string) PublicIPAddressSkuTierPtrInput {
	return (*publicIPAddressSkuTierPtr)(&v)
}

func (*publicIPAddressSkuTierPtr) ElementType() reflect.Type {
	return publicIPAddressSkuTierPtrType
}

func (in *publicIPAddressSkuTierPtr) ToPublicIPAddressSkuTierPtrOutput() PublicIPAddressSkuTierPtrOutput {
	return pulumi.ToOutput(in).(PublicIPAddressSkuTierPtrOutput)
}

func (in *publicIPAddressSkuTierPtr) ToPublicIPAddressSkuTierPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicIPAddressSkuTierPtrOutput)
}

type PublicIPPrefixSkuName string

const (
	PublicIPPrefixSkuNameStandard = PublicIPPrefixSkuName("Standard")
)

func (PublicIPPrefixSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPPrefixSkuName)(nil)).Elem()
}

func (e PublicIPPrefixSkuName) ToPublicIPPrefixSkuNameOutput() PublicIPPrefixSkuNameOutput {
	return pulumi.ToOutput(e).(PublicIPPrefixSkuNameOutput)
}

func (e PublicIPPrefixSkuName) ToPublicIPPrefixSkuNameOutputWithContext(ctx context.Context) PublicIPPrefixSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicIPPrefixSkuNameOutput)
}

func (e PublicIPPrefixSkuName) ToPublicIPPrefixSkuNamePtrOutput() PublicIPPrefixSkuNamePtrOutput {
	return e.ToPublicIPPrefixSkuNamePtrOutputWithContext(context.Background())
}

func (e PublicIPPrefixSkuName) ToPublicIPPrefixSkuNamePtrOutputWithContext(ctx context.Context) PublicIPPrefixSkuNamePtrOutput {
	return PublicIPPrefixSkuName(e).ToPublicIPPrefixSkuNameOutputWithContext(ctx).ToPublicIPPrefixSkuNamePtrOutputWithContext(ctx)
}

func (e PublicIPPrefixSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPPrefixSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPPrefixSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicIPPrefixSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicIPPrefixSkuNameOutput struct{ *pulumi.OutputState }

func (PublicIPPrefixSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPPrefixSkuName)(nil)).Elem()
}

func (o PublicIPPrefixSkuNameOutput) ToPublicIPPrefixSkuNameOutput() PublicIPPrefixSkuNameOutput {
	return o
}

func (o PublicIPPrefixSkuNameOutput) ToPublicIPPrefixSkuNameOutputWithContext(ctx context.Context) PublicIPPrefixSkuNameOutput {
	return o
}

func (o PublicIPPrefixSkuNameOutput) ToPublicIPPrefixSkuNamePtrOutput() PublicIPPrefixSkuNamePtrOutput {
	return o.ToPublicIPPrefixSkuNamePtrOutputWithContext(context.Background())
}

func (o PublicIPPrefixSkuNameOutput) ToPublicIPPrefixSkuNamePtrOutputWithContext(ctx context.Context) PublicIPPrefixSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPPrefixSkuName) *PublicIPPrefixSkuName {
		return &v
	}).(PublicIPPrefixSkuNamePtrOutput)
}

func (o PublicIPPrefixSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicIPPrefixSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPPrefixSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicIPPrefixSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPPrefixSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPPrefixSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicIPPrefixSkuNamePtrOutput struct{ *pulumi.OutputState }

func (PublicIPPrefixSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPPrefixSkuName)(nil)).Elem()
}

func (o PublicIPPrefixSkuNamePtrOutput) ToPublicIPPrefixSkuNamePtrOutput() PublicIPPrefixSkuNamePtrOutput {
	return o
}

func (o PublicIPPrefixSkuNamePtrOutput) ToPublicIPPrefixSkuNamePtrOutputWithContext(ctx context.Context) PublicIPPrefixSkuNamePtrOutput {
	return o
}

func (o PublicIPPrefixSkuNamePtrOutput) Elem() PublicIPPrefixSkuNameOutput {
	return o.ApplyT(func(v *PublicIPPrefixSkuName) PublicIPPrefixSkuName {
		if v != nil {
			return *v
		}
		var ret PublicIPPrefixSkuName
		return ret
	}).(PublicIPPrefixSkuNameOutput)
}

func (o PublicIPPrefixSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPPrefixSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicIPPrefixSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicIPPrefixSkuNameInput interface {
	pulumi.Input

	ToPublicIPPrefixSkuNameOutput() PublicIPPrefixSkuNameOutput
	ToPublicIPPrefixSkuNameOutputWithContext(context.Context) PublicIPPrefixSkuNameOutput
}

var publicIPPrefixSkuNamePtrType = reflect.TypeOf((**PublicIPPrefixSkuName)(nil)).Elem()

type PublicIPPrefixSkuNamePtrInput interface {
	pulumi.Input

	ToPublicIPPrefixSkuNamePtrOutput() PublicIPPrefixSkuNamePtrOutput
	ToPublicIPPrefixSkuNamePtrOutputWithContext(context.Context) PublicIPPrefixSkuNamePtrOutput
}

type publicIPPrefixSkuNamePtr string

func PublicIPPrefixSkuNamePtr(v string) PublicIPPrefixSkuNamePtrInput {
	return (*publicIPPrefixSkuNamePtr)(&v)
}

func (*publicIPPrefixSkuNamePtr) ElementType() reflect.Type {
	return publicIPPrefixSkuNamePtrType
}

func (in *publicIPPrefixSkuNamePtr) ToPublicIPPrefixSkuNamePtrOutput() PublicIPPrefixSkuNamePtrOutput {
	return pulumi.ToOutput(in).(PublicIPPrefixSkuNamePtrOutput)
}

func (in *publicIPPrefixSkuNamePtr) ToPublicIPPrefixSkuNamePtrOutputWithContext(ctx context.Context) PublicIPPrefixSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicIPPrefixSkuNamePtrOutput)
}

type PublicIPPrefixSkuTier string

const (
	PublicIPPrefixSkuTierRegional = PublicIPPrefixSkuTier("Regional")
	PublicIPPrefixSkuTierGlobal   = PublicIPPrefixSkuTier("Global")
)

func (PublicIPPrefixSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPPrefixSkuTier)(nil)).Elem()
}

func (e PublicIPPrefixSkuTier) ToPublicIPPrefixSkuTierOutput() PublicIPPrefixSkuTierOutput {
	return pulumi.ToOutput(e).(PublicIPPrefixSkuTierOutput)
}

func (e PublicIPPrefixSkuTier) ToPublicIPPrefixSkuTierOutputWithContext(ctx context.Context) PublicIPPrefixSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicIPPrefixSkuTierOutput)
}

func (e PublicIPPrefixSkuTier) ToPublicIPPrefixSkuTierPtrOutput() PublicIPPrefixSkuTierPtrOutput {
	return e.ToPublicIPPrefixSkuTierPtrOutputWithContext(context.Background())
}

func (e PublicIPPrefixSkuTier) ToPublicIPPrefixSkuTierPtrOutputWithContext(ctx context.Context) PublicIPPrefixSkuTierPtrOutput {
	return PublicIPPrefixSkuTier(e).ToPublicIPPrefixSkuTierOutputWithContext(ctx).ToPublicIPPrefixSkuTierPtrOutputWithContext(ctx)
}

func (e PublicIPPrefixSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPPrefixSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicIPPrefixSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicIPPrefixSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicIPPrefixSkuTierOutput struct{ *pulumi.OutputState }

func (PublicIPPrefixSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPPrefixSkuTier)(nil)).Elem()
}

func (o PublicIPPrefixSkuTierOutput) ToPublicIPPrefixSkuTierOutput() PublicIPPrefixSkuTierOutput {
	return o
}

func (o PublicIPPrefixSkuTierOutput) ToPublicIPPrefixSkuTierOutputWithContext(ctx context.Context) PublicIPPrefixSkuTierOutput {
	return o
}

func (o PublicIPPrefixSkuTierOutput) ToPublicIPPrefixSkuTierPtrOutput() PublicIPPrefixSkuTierPtrOutput {
	return o.ToPublicIPPrefixSkuTierPtrOutputWithContext(context.Background())
}

func (o PublicIPPrefixSkuTierOutput) ToPublicIPPrefixSkuTierPtrOutputWithContext(ctx context.Context) PublicIPPrefixSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPPrefixSkuTier) *PublicIPPrefixSkuTier {
		return &v
	}).(PublicIPPrefixSkuTierPtrOutput)
}

func (o PublicIPPrefixSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicIPPrefixSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPPrefixSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicIPPrefixSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPPrefixSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicIPPrefixSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicIPPrefixSkuTierPtrOutput struct{ *pulumi.OutputState }

func (PublicIPPrefixSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPPrefixSkuTier)(nil)).Elem()
}

func (o PublicIPPrefixSkuTierPtrOutput) ToPublicIPPrefixSkuTierPtrOutput() PublicIPPrefixSkuTierPtrOutput {
	return o
}

func (o PublicIPPrefixSkuTierPtrOutput) ToPublicIPPrefixSkuTierPtrOutputWithContext(ctx context.Context) PublicIPPrefixSkuTierPtrOutput {
	return o
}

func (o PublicIPPrefixSkuTierPtrOutput) Elem() PublicIPPrefixSkuTierOutput {
	return o.ApplyT(func(v *PublicIPPrefixSkuTier) PublicIPPrefixSkuTier {
		if v != nil {
			return *v
		}
		var ret PublicIPPrefixSkuTier
		return ret
	}).(PublicIPPrefixSkuTierOutput)
}

func (o PublicIPPrefixSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicIPPrefixSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicIPPrefixSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicIPPrefixSkuTierInput interface {
	pulumi.Input

	ToPublicIPPrefixSkuTierOutput() PublicIPPrefixSkuTierOutput
	ToPublicIPPrefixSkuTierOutputWithContext(context.Context) PublicIPPrefixSkuTierOutput
}

var publicIPPrefixSkuTierPtrType = reflect.TypeOf((**PublicIPPrefixSkuTier)(nil)).Elem()

type PublicIPPrefixSkuTierPtrInput interface {
	pulumi.Input

	ToPublicIPPrefixSkuTierPtrOutput() PublicIPPrefixSkuTierPtrOutput
	ToPublicIPPrefixSkuTierPtrOutputWithContext(context.Context) PublicIPPrefixSkuTierPtrOutput
}

type publicIPPrefixSkuTierPtr string

func PublicIPPrefixSkuTierPtr(v string) PublicIPPrefixSkuTierPtrInput {
	return (*publicIPPrefixSkuTierPtr)(&v)
}

func (*publicIPPrefixSkuTierPtr) ElementType() reflect.Type {
	return publicIPPrefixSkuTierPtrType
}

func (in *publicIPPrefixSkuTierPtr) ToPublicIPPrefixSkuTierPtrOutput() PublicIPPrefixSkuTierPtrOutput {
	return pulumi.ToOutput(in).(PublicIPPrefixSkuTierPtrOutput)
}

func (in *publicIPPrefixSkuTierPtr) ToPublicIPPrefixSkuTierPtrOutputWithContext(ctx context.Context) PublicIPPrefixSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicIPPrefixSkuTierPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
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

type SecurityProviderName string

const (
	SecurityProviderNameZScaler    = SecurityProviderName("ZScaler")
	SecurityProviderNameIBoss      = SecurityProviderName("IBoss")
	SecurityProviderNameCheckpoint = SecurityProviderName("Checkpoint")
)

func (SecurityProviderName) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProviderName)(nil)).Elem()
}

func (e SecurityProviderName) ToSecurityProviderNameOutput() SecurityProviderNameOutput {
	return pulumi.ToOutput(e).(SecurityProviderNameOutput)
}

func (e SecurityProviderName) ToSecurityProviderNameOutputWithContext(ctx context.Context) SecurityProviderNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityProviderNameOutput)
}

func (e SecurityProviderName) ToSecurityProviderNamePtrOutput() SecurityProviderNamePtrOutput {
	return e.ToSecurityProviderNamePtrOutputWithContext(context.Background())
}

func (e SecurityProviderName) ToSecurityProviderNamePtrOutputWithContext(ctx context.Context) SecurityProviderNamePtrOutput {
	return SecurityProviderName(e).ToSecurityProviderNameOutputWithContext(ctx).ToSecurityProviderNamePtrOutputWithContext(ctx)
}

func (e SecurityProviderName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityProviderName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityProviderName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityProviderName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityProviderNameOutput struct{ *pulumi.OutputState }

func (SecurityProviderNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProviderName)(nil)).Elem()
}

func (o SecurityProviderNameOutput) ToSecurityProviderNameOutput() SecurityProviderNameOutput {
	return o
}

func (o SecurityProviderNameOutput) ToSecurityProviderNameOutputWithContext(ctx context.Context) SecurityProviderNameOutput {
	return o
}

func (o SecurityProviderNameOutput) ToSecurityProviderNamePtrOutput() SecurityProviderNamePtrOutput {
	return o.ToSecurityProviderNamePtrOutputWithContext(context.Background())
}

func (o SecurityProviderNameOutput) ToSecurityProviderNamePtrOutputWithContext(ctx context.Context) SecurityProviderNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityProviderName) *SecurityProviderName {
		return &v
	}).(SecurityProviderNamePtrOutput)
}

func (o SecurityProviderNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityProviderNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityProviderName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityProviderNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityProviderNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityProviderName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityProviderNamePtrOutput struct{ *pulumi.OutputState }

func (SecurityProviderNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProviderName)(nil)).Elem()
}

func (o SecurityProviderNamePtrOutput) ToSecurityProviderNamePtrOutput() SecurityProviderNamePtrOutput {
	return o
}

func (o SecurityProviderNamePtrOutput) ToSecurityProviderNamePtrOutputWithContext(ctx context.Context) SecurityProviderNamePtrOutput {
	return o
}

func (o SecurityProviderNamePtrOutput) Elem() SecurityProviderNameOutput {
	return o.ApplyT(func(v *SecurityProviderName) SecurityProviderName {
		if v != nil {
			return *v
		}
		var ret SecurityProviderName
		return ret
	}).(SecurityProviderNameOutput)
}

func (o SecurityProviderNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityProviderNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityProviderName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityProviderNameInput interface {
	pulumi.Input

	ToSecurityProviderNameOutput() SecurityProviderNameOutput
	ToSecurityProviderNameOutputWithContext(context.Context) SecurityProviderNameOutput
}

var securityProviderNamePtrType = reflect.TypeOf((**SecurityProviderName)(nil)).Elem()

type SecurityProviderNamePtrInput interface {
	pulumi.Input

	ToSecurityProviderNamePtrOutput() SecurityProviderNamePtrOutput
	ToSecurityProviderNamePtrOutputWithContext(context.Context) SecurityProviderNamePtrOutput
}

type securityProviderNamePtr string

func SecurityProviderNamePtr(v string) SecurityProviderNamePtrInput {
	return (*securityProviderNamePtr)(&v)
}

func (*securityProviderNamePtr) ElementType() reflect.Type {
	return securityProviderNamePtrType
}

func (in *securityProviderNamePtr) ToSecurityProviderNamePtrOutput() SecurityProviderNamePtrOutput {
	return pulumi.ToOutput(in).(SecurityProviderNamePtrOutput)
}

func (in *securityProviderNamePtr) ToSecurityProviderNamePtrOutputWithContext(ctx context.Context) SecurityProviderNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityProviderNamePtrOutput)
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
	SecurityRuleProtocolIcmp     = SecurityRuleProtocol("Icmp")
	SecurityRuleProtocolEsp      = SecurityRuleProtocol("Esp")
	SecurityRuleProtocolAsterisk = SecurityRuleProtocol("*")
	SecurityRuleProtocolAh       = SecurityRuleProtocol("Ah")
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
	TransportProtocolAll = TransportProtocol("All")
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

type VirtualNetworkGatewayConnectionMode string

const (
	VirtualNetworkGatewayConnectionModeDefault       = VirtualNetworkGatewayConnectionMode("Default")
	VirtualNetworkGatewayConnectionModeResponderOnly = VirtualNetworkGatewayConnectionMode("ResponderOnly")
	VirtualNetworkGatewayConnectionModeInitiatorOnly = VirtualNetworkGatewayConnectionMode("InitiatorOnly")
)

func (VirtualNetworkGatewayConnectionMode) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayConnectionMode)(nil)).Elem()
}

func (e VirtualNetworkGatewayConnectionMode) ToVirtualNetworkGatewayConnectionModeOutput() VirtualNetworkGatewayConnectionModeOutput {
	return pulumi.ToOutput(e).(VirtualNetworkGatewayConnectionModeOutput)
}

func (e VirtualNetworkGatewayConnectionMode) ToVirtualNetworkGatewayConnectionModeOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkGatewayConnectionModeOutput)
}

func (e VirtualNetworkGatewayConnectionMode) ToVirtualNetworkGatewayConnectionModePtrOutput() VirtualNetworkGatewayConnectionModePtrOutput {
	return e.ToVirtualNetworkGatewayConnectionModePtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewayConnectionMode) ToVirtualNetworkGatewayConnectionModePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionModePtrOutput {
	return VirtualNetworkGatewayConnectionMode(e).ToVirtualNetworkGatewayConnectionModeOutputWithContext(ctx).ToVirtualNetworkGatewayConnectionModePtrOutputWithContext(ctx)
}

func (e VirtualNetworkGatewayConnectionMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewayConnectionMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewayConnectionMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewayConnectionMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkGatewayConnectionModeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayConnectionModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayConnectionMode)(nil)).Elem()
}

func (o VirtualNetworkGatewayConnectionModeOutput) ToVirtualNetworkGatewayConnectionModeOutput() VirtualNetworkGatewayConnectionModeOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionModeOutput) ToVirtualNetworkGatewayConnectionModeOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionModeOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionModeOutput) ToVirtualNetworkGatewayConnectionModePtrOutput() VirtualNetworkGatewayConnectionModePtrOutput {
	return o.ToVirtualNetworkGatewayConnectionModePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionModeOutput) ToVirtualNetworkGatewayConnectionModePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewayConnectionMode) *VirtualNetworkGatewayConnectionMode {
		return &v
	}).(VirtualNetworkGatewayConnectionModePtrOutput)
}

func (o VirtualNetworkGatewayConnectionModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewayConnectionMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewayConnectionMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayConnectionModePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayConnectionModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayConnectionMode)(nil)).Elem()
}

func (o VirtualNetworkGatewayConnectionModePtrOutput) ToVirtualNetworkGatewayConnectionModePtrOutput() VirtualNetworkGatewayConnectionModePtrOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionModePtrOutput) ToVirtualNetworkGatewayConnectionModePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionModePtrOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionModePtrOutput) Elem() VirtualNetworkGatewayConnectionModeOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnectionMode) VirtualNetworkGatewayConnectionMode {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayConnectionMode
		return ret
	}).(VirtualNetworkGatewayConnectionModeOutput)
}

func (o VirtualNetworkGatewayConnectionModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkGatewayConnectionMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkGatewayConnectionModeInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayConnectionModeOutput() VirtualNetworkGatewayConnectionModeOutput
	ToVirtualNetworkGatewayConnectionModeOutputWithContext(context.Context) VirtualNetworkGatewayConnectionModeOutput
}

var virtualNetworkGatewayConnectionModePtrType = reflect.TypeOf((**VirtualNetworkGatewayConnectionMode)(nil)).Elem()

type VirtualNetworkGatewayConnectionModePtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayConnectionModePtrOutput() VirtualNetworkGatewayConnectionModePtrOutput
	ToVirtualNetworkGatewayConnectionModePtrOutputWithContext(context.Context) VirtualNetworkGatewayConnectionModePtrOutput
}

type virtualNetworkGatewayConnectionModePtr string

func VirtualNetworkGatewayConnectionModePtr(v string) VirtualNetworkGatewayConnectionModePtrInput {
	return (*virtualNetworkGatewayConnectionModePtr)(&v)
}

func (*virtualNetworkGatewayConnectionModePtr) ElementType() reflect.Type {
	return virtualNetworkGatewayConnectionModePtrType
}

func (in *virtualNetworkGatewayConnectionModePtr) ToVirtualNetworkGatewayConnectionModePtrOutput() VirtualNetworkGatewayConnectionModePtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkGatewayConnectionModePtrOutput)
}

func (in *virtualNetworkGatewayConnectionModePtr) ToVirtualNetworkGatewayConnectionModePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkGatewayConnectionModePtrOutput)
}

type VirtualNetworkGatewayConnectionProtocol string

const (
	VirtualNetworkGatewayConnectionProtocolIKEv2 = VirtualNetworkGatewayConnectionProtocol("IKEv2")
	VirtualNetworkGatewayConnectionProtocolIKEv1 = VirtualNetworkGatewayConnectionProtocol("IKEv1")
)

func (VirtualNetworkGatewayConnectionProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayConnectionProtocol)(nil)).Elem()
}

func (e VirtualNetworkGatewayConnectionProtocol) ToVirtualNetworkGatewayConnectionProtocolOutput() VirtualNetworkGatewayConnectionProtocolOutput {
	return pulumi.ToOutput(e).(VirtualNetworkGatewayConnectionProtocolOutput)
}

func (e VirtualNetworkGatewayConnectionProtocol) ToVirtualNetworkGatewayConnectionProtocolOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkGatewayConnectionProtocolOutput)
}

func (e VirtualNetworkGatewayConnectionProtocol) ToVirtualNetworkGatewayConnectionProtocolPtrOutput() VirtualNetworkGatewayConnectionProtocolPtrOutput {
	return e.ToVirtualNetworkGatewayConnectionProtocolPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewayConnectionProtocol) ToVirtualNetworkGatewayConnectionProtocolPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionProtocolPtrOutput {
	return VirtualNetworkGatewayConnectionProtocol(e).ToVirtualNetworkGatewayConnectionProtocolOutputWithContext(ctx).ToVirtualNetworkGatewayConnectionProtocolPtrOutputWithContext(ctx)
}

func (e VirtualNetworkGatewayConnectionProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewayConnectionProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkGatewayConnectionProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkGatewayConnectionProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkGatewayConnectionProtocolOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayConnectionProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayConnectionProtocol)(nil)).Elem()
}

func (o VirtualNetworkGatewayConnectionProtocolOutput) ToVirtualNetworkGatewayConnectionProtocolOutput() VirtualNetworkGatewayConnectionProtocolOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionProtocolOutput) ToVirtualNetworkGatewayConnectionProtocolOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionProtocolOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionProtocolOutput) ToVirtualNetworkGatewayConnectionProtocolPtrOutput() VirtualNetworkGatewayConnectionProtocolPtrOutput {
	return o.ToVirtualNetworkGatewayConnectionProtocolPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionProtocolOutput) ToVirtualNetworkGatewayConnectionProtocolPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewayConnectionProtocol) *VirtualNetworkGatewayConnectionProtocol {
		return &v
	}).(VirtualNetworkGatewayConnectionProtocolPtrOutput)
}

func (o VirtualNetworkGatewayConnectionProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewayConnectionProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkGatewayConnectionProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayConnectionProtocolPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayConnectionProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayConnectionProtocol)(nil)).Elem()
}

func (o VirtualNetworkGatewayConnectionProtocolPtrOutput) ToVirtualNetworkGatewayConnectionProtocolPtrOutput() VirtualNetworkGatewayConnectionProtocolPtrOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionProtocolPtrOutput) ToVirtualNetworkGatewayConnectionProtocolPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionProtocolPtrOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionProtocolPtrOutput) Elem() VirtualNetworkGatewayConnectionProtocolOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnectionProtocol) VirtualNetworkGatewayConnectionProtocol {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayConnectionProtocol
		return ret
	}).(VirtualNetworkGatewayConnectionProtocolOutput)
}

func (o VirtualNetworkGatewayConnectionProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayConnectionProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkGatewayConnectionProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkGatewayConnectionProtocolInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayConnectionProtocolOutput() VirtualNetworkGatewayConnectionProtocolOutput
	ToVirtualNetworkGatewayConnectionProtocolOutputWithContext(context.Context) VirtualNetworkGatewayConnectionProtocolOutput
}

var virtualNetworkGatewayConnectionProtocolPtrType = reflect.TypeOf((**VirtualNetworkGatewayConnectionProtocol)(nil)).Elem()

type VirtualNetworkGatewayConnectionProtocolPtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayConnectionProtocolPtrOutput() VirtualNetworkGatewayConnectionProtocolPtrOutput
	ToVirtualNetworkGatewayConnectionProtocolPtrOutputWithContext(context.Context) VirtualNetworkGatewayConnectionProtocolPtrOutput
}

type virtualNetworkGatewayConnectionProtocolPtr string

func VirtualNetworkGatewayConnectionProtocolPtr(v string) VirtualNetworkGatewayConnectionProtocolPtrInput {
	return (*virtualNetworkGatewayConnectionProtocolPtr)(&v)
}

func (*virtualNetworkGatewayConnectionProtocolPtr) ElementType() reflect.Type {
	return virtualNetworkGatewayConnectionProtocolPtrType
}

func (in *virtualNetworkGatewayConnectionProtocolPtr) ToVirtualNetworkGatewayConnectionProtocolPtrOutput() VirtualNetworkGatewayConnectionProtocolPtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkGatewayConnectionProtocolPtrOutput)
}

func (in *virtualNetworkGatewayConnectionProtocolPtr) ToVirtualNetworkGatewayConnectionProtocolPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkGatewayConnectionProtocolPtrOutput)
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
	VirtualNetworkGatewaySkuNameVpnGw4           = VirtualNetworkGatewaySkuName("VpnGw4")
	VirtualNetworkGatewaySkuNameVpnGw5           = VirtualNetworkGatewaySkuName("VpnGw5")
	VirtualNetworkGatewaySkuNameVpnGw1AZ         = VirtualNetworkGatewaySkuName("VpnGw1AZ")
	VirtualNetworkGatewaySkuNameVpnGw2AZ         = VirtualNetworkGatewaySkuName("VpnGw2AZ")
	VirtualNetworkGatewaySkuNameVpnGw3AZ         = VirtualNetworkGatewaySkuName("VpnGw3AZ")
	VirtualNetworkGatewaySkuNameVpnGw4AZ         = VirtualNetworkGatewaySkuName("VpnGw4AZ")
	VirtualNetworkGatewaySkuNameVpnGw5AZ         = VirtualNetworkGatewaySkuName("VpnGw5AZ")
	VirtualNetworkGatewaySkuNameErGw1AZ          = VirtualNetworkGatewaySkuName("ErGw1AZ")
	VirtualNetworkGatewaySkuNameErGw2AZ          = VirtualNetworkGatewaySkuName("ErGw2AZ")
	VirtualNetworkGatewaySkuNameErGw3AZ          = VirtualNetworkGatewaySkuName("ErGw3AZ")
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
	VirtualNetworkGatewaySkuTierVpnGw4           = VirtualNetworkGatewaySkuTier("VpnGw4")
	VirtualNetworkGatewaySkuTierVpnGw5           = VirtualNetworkGatewaySkuTier("VpnGw5")
	VirtualNetworkGatewaySkuTierVpnGw1AZ         = VirtualNetworkGatewaySkuTier("VpnGw1AZ")
	VirtualNetworkGatewaySkuTierVpnGw2AZ         = VirtualNetworkGatewaySkuTier("VpnGw2AZ")
	VirtualNetworkGatewaySkuTierVpnGw3AZ         = VirtualNetworkGatewaySkuTier("VpnGw3AZ")
	VirtualNetworkGatewaySkuTierVpnGw4AZ         = VirtualNetworkGatewaySkuTier("VpnGw4AZ")
	VirtualNetworkGatewaySkuTierVpnGw5AZ         = VirtualNetworkGatewaySkuTier("VpnGw5AZ")
	VirtualNetworkGatewaySkuTierErGw1AZ          = VirtualNetworkGatewaySkuTier("ErGw1AZ")
	VirtualNetworkGatewaySkuTierErGw2AZ          = VirtualNetworkGatewaySkuTier("ErGw2AZ")
	VirtualNetworkGatewaySkuTierErGw3AZ          = VirtualNetworkGatewaySkuTier("ErGw3AZ")
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
	VirtualNetworkGatewayTypeEnumLocalGateway = VirtualNetworkGatewayTypeEnum("LocalGateway")
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

type VirtualNetworkPeeringLevel string

const (
	VirtualNetworkPeeringLevelFullyInSync             = VirtualNetworkPeeringLevel("FullyInSync")
	VirtualNetworkPeeringLevelRemoteNotInSync         = VirtualNetworkPeeringLevel("RemoteNotInSync")
	VirtualNetworkPeeringLevelLocalNotInSync          = VirtualNetworkPeeringLevel("LocalNotInSync")
	VirtualNetworkPeeringLevelLocalAndRemoteNotInSync = VirtualNetworkPeeringLevel("LocalAndRemoteNotInSync")
)

func (VirtualNetworkPeeringLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringLevel)(nil)).Elem()
}

func (e VirtualNetworkPeeringLevel) ToVirtualNetworkPeeringLevelOutput() VirtualNetworkPeeringLevelOutput {
	return pulumi.ToOutput(e).(VirtualNetworkPeeringLevelOutput)
}

func (e VirtualNetworkPeeringLevel) ToVirtualNetworkPeeringLevelOutputWithContext(ctx context.Context) VirtualNetworkPeeringLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkPeeringLevelOutput)
}

func (e VirtualNetworkPeeringLevel) ToVirtualNetworkPeeringLevelPtrOutput() VirtualNetworkPeeringLevelPtrOutput {
	return e.ToVirtualNetworkPeeringLevelPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkPeeringLevel) ToVirtualNetworkPeeringLevelPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringLevelPtrOutput {
	return VirtualNetworkPeeringLevel(e).ToVirtualNetworkPeeringLevelOutputWithContext(ctx).ToVirtualNetworkPeeringLevelPtrOutputWithContext(ctx)
}

func (e VirtualNetworkPeeringLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkPeeringLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkPeeringLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkPeeringLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkPeeringLevelOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringLevel)(nil)).Elem()
}

func (o VirtualNetworkPeeringLevelOutput) ToVirtualNetworkPeeringLevelOutput() VirtualNetworkPeeringLevelOutput {
	return o
}

func (o VirtualNetworkPeeringLevelOutput) ToVirtualNetworkPeeringLevelOutputWithContext(ctx context.Context) VirtualNetworkPeeringLevelOutput {
	return o
}

func (o VirtualNetworkPeeringLevelOutput) ToVirtualNetworkPeeringLevelPtrOutput() VirtualNetworkPeeringLevelPtrOutput {
	return o.ToVirtualNetworkPeeringLevelPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringLevelOutput) ToVirtualNetworkPeeringLevelPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPeeringLevel) *VirtualNetworkPeeringLevel {
		return &v
	}).(VirtualNetworkPeeringLevelPtrOutput)
}

func (o VirtualNetworkPeeringLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkPeeringLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkPeeringLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkPeeringLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringLevelPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringLevel)(nil)).Elem()
}

func (o VirtualNetworkPeeringLevelPtrOutput) ToVirtualNetworkPeeringLevelPtrOutput() VirtualNetworkPeeringLevelPtrOutput {
	return o
}

func (o VirtualNetworkPeeringLevelPtrOutput) ToVirtualNetworkPeeringLevelPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringLevelPtrOutput {
	return o
}

func (o VirtualNetworkPeeringLevelPtrOutput) Elem() VirtualNetworkPeeringLevelOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringLevel) VirtualNetworkPeeringLevel {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPeeringLevel
		return ret
	}).(VirtualNetworkPeeringLevelOutput)
}

func (o VirtualNetworkPeeringLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkPeeringLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkPeeringLevelInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringLevelOutput() VirtualNetworkPeeringLevelOutput
	ToVirtualNetworkPeeringLevelOutputWithContext(context.Context) VirtualNetworkPeeringLevelOutput
}

var virtualNetworkPeeringLevelPtrType = reflect.TypeOf((**VirtualNetworkPeeringLevel)(nil)).Elem()

type VirtualNetworkPeeringLevelPtrInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringLevelPtrOutput() VirtualNetworkPeeringLevelPtrOutput
	ToVirtualNetworkPeeringLevelPtrOutputWithContext(context.Context) VirtualNetworkPeeringLevelPtrOutput
}

type virtualNetworkPeeringLevelPtr string

func VirtualNetworkPeeringLevelPtr(v string) VirtualNetworkPeeringLevelPtrInput {
	return (*virtualNetworkPeeringLevelPtr)(&v)
}

func (*virtualNetworkPeeringLevelPtr) ElementType() reflect.Type {
	return virtualNetworkPeeringLevelPtrType
}

func (in *virtualNetworkPeeringLevelPtr) ToVirtualNetworkPeeringLevelPtrOutput() VirtualNetworkPeeringLevelPtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkPeeringLevelPtrOutput)
}

func (in *virtualNetworkPeeringLevelPtr) ToVirtualNetworkPeeringLevelPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkPeeringLevelPtrOutput)
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

type VirtualNetworkPrivateEndpointNetworkPolicies string

const (
	VirtualNetworkPrivateEndpointNetworkPoliciesEnabled  = VirtualNetworkPrivateEndpointNetworkPolicies("Enabled")
	VirtualNetworkPrivateEndpointNetworkPoliciesDisabled = VirtualNetworkPrivateEndpointNetworkPolicies("Disabled")
)

func (VirtualNetworkPrivateEndpointNetworkPolicies) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPrivateEndpointNetworkPolicies)(nil)).Elem()
}

func (e VirtualNetworkPrivateEndpointNetworkPolicies) ToVirtualNetworkPrivateEndpointNetworkPoliciesOutput() VirtualNetworkPrivateEndpointNetworkPoliciesOutput {
	return pulumi.ToOutput(e).(VirtualNetworkPrivateEndpointNetworkPoliciesOutput)
}

func (e VirtualNetworkPrivateEndpointNetworkPolicies) ToVirtualNetworkPrivateEndpointNetworkPoliciesOutputWithContext(ctx context.Context) VirtualNetworkPrivateEndpointNetworkPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkPrivateEndpointNetworkPoliciesOutput)
}

func (e VirtualNetworkPrivateEndpointNetworkPolicies) ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput() VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput {
	return e.ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkPrivateEndpointNetworkPolicies) ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx context.Context) VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput {
	return VirtualNetworkPrivateEndpointNetworkPolicies(e).ToVirtualNetworkPrivateEndpointNetworkPoliciesOutputWithContext(ctx).ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx)
}

func (e VirtualNetworkPrivateEndpointNetworkPolicies) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkPrivateEndpointNetworkPolicies) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkPrivateEndpointNetworkPolicies) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkPrivateEndpointNetworkPolicies) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkPrivateEndpointNetworkPoliciesOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPrivateEndpointNetworkPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPrivateEndpointNetworkPolicies)(nil)).Elem()
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesOutput) ToVirtualNetworkPrivateEndpointNetworkPoliciesOutput() VirtualNetworkPrivateEndpointNetworkPoliciesOutput {
	return o
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesOutput) ToVirtualNetworkPrivateEndpointNetworkPoliciesOutputWithContext(ctx context.Context) VirtualNetworkPrivateEndpointNetworkPoliciesOutput {
	return o
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesOutput) ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput() VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput {
	return o.ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesOutput) ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx context.Context) VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPrivateEndpointNetworkPolicies) *VirtualNetworkPrivateEndpointNetworkPolicies {
		return &v
	}).(VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput)
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkPrivateEndpointNetworkPolicies) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkPrivateEndpointNetworkPolicies) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPrivateEndpointNetworkPolicies)(nil)).Elem()
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput) ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput() VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput {
	return o
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput) ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx context.Context) VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput {
	return o
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput) Elem() VirtualNetworkPrivateEndpointNetworkPoliciesOutput {
	return o.ApplyT(func(v *VirtualNetworkPrivateEndpointNetworkPolicies) VirtualNetworkPrivateEndpointNetworkPolicies {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPrivateEndpointNetworkPolicies
		return ret
	}).(VirtualNetworkPrivateEndpointNetworkPoliciesOutput)
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkPrivateEndpointNetworkPolicies) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkPrivateEndpointNetworkPoliciesInput interface {
	pulumi.Input

	ToVirtualNetworkPrivateEndpointNetworkPoliciesOutput() VirtualNetworkPrivateEndpointNetworkPoliciesOutput
	ToVirtualNetworkPrivateEndpointNetworkPoliciesOutputWithContext(context.Context) VirtualNetworkPrivateEndpointNetworkPoliciesOutput
}

var virtualNetworkPrivateEndpointNetworkPoliciesPtrType = reflect.TypeOf((**VirtualNetworkPrivateEndpointNetworkPolicies)(nil)).Elem()

type VirtualNetworkPrivateEndpointNetworkPoliciesPtrInput interface {
	pulumi.Input

	ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput() VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput
	ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutputWithContext(context.Context) VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput
}

type virtualNetworkPrivateEndpointNetworkPoliciesPtr string

func VirtualNetworkPrivateEndpointNetworkPoliciesPtr(v string) VirtualNetworkPrivateEndpointNetworkPoliciesPtrInput {
	return (*virtualNetworkPrivateEndpointNetworkPoliciesPtr)(&v)
}

func (*virtualNetworkPrivateEndpointNetworkPoliciesPtr) ElementType() reflect.Type {
	return virtualNetworkPrivateEndpointNetworkPoliciesPtrType
}

func (in *virtualNetworkPrivateEndpointNetworkPoliciesPtr) ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput() VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput)
}

func (in *virtualNetworkPrivateEndpointNetworkPoliciesPtr) ToVirtualNetworkPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx context.Context) VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput)
}

type VirtualNetworkPrivateLinkServiceNetworkPolicies string

const (
	VirtualNetworkPrivateLinkServiceNetworkPoliciesEnabled  = VirtualNetworkPrivateLinkServiceNetworkPolicies("Enabled")
	VirtualNetworkPrivateLinkServiceNetworkPoliciesDisabled = VirtualNetworkPrivateLinkServiceNetworkPolicies("Disabled")
)

func (VirtualNetworkPrivateLinkServiceNetworkPolicies) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPrivateLinkServiceNetworkPolicies)(nil)).Elem()
}

func (e VirtualNetworkPrivateLinkServiceNetworkPolicies) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesOutput() VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput {
	return pulumi.ToOutput(e).(VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput)
}

func (e VirtualNetworkPrivateLinkServiceNetworkPolicies) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesOutputWithContext(ctx context.Context) VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput)
}

func (e VirtualNetworkPrivateLinkServiceNetworkPolicies) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput() VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput {
	return e.ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkPrivateLinkServiceNetworkPolicies) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx context.Context) VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput {
	return VirtualNetworkPrivateLinkServiceNetworkPolicies(e).ToVirtualNetworkPrivateLinkServiceNetworkPoliciesOutputWithContext(ctx).ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx)
}

func (e VirtualNetworkPrivateLinkServiceNetworkPolicies) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkPrivateLinkServiceNetworkPolicies) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualNetworkPrivateLinkServiceNetworkPolicies) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualNetworkPrivateLinkServiceNetworkPolicies) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPrivateLinkServiceNetworkPolicies)(nil)).Elem()
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesOutput() VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput {
	return o
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesOutputWithContext(ctx context.Context) VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput {
	return o
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput() VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput {
	return o.ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx context.Context) VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPrivateLinkServiceNetworkPolicies) *VirtualNetworkPrivateLinkServiceNetworkPolicies {
		return &v
	}).(VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput)
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkPrivateLinkServiceNetworkPolicies) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualNetworkPrivateLinkServiceNetworkPolicies) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPrivateLinkServiceNetworkPolicies)(nil)).Elem()
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput() VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput {
	return o
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx context.Context) VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput {
	return o
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput) Elem() VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput {
	return o.ApplyT(func(v *VirtualNetworkPrivateLinkServiceNetworkPolicies) VirtualNetworkPrivateLinkServiceNetworkPolicies {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPrivateLinkServiceNetworkPolicies
		return ret
	}).(VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput)
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualNetworkPrivateLinkServiceNetworkPolicies) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualNetworkPrivateLinkServiceNetworkPoliciesInput interface {
	pulumi.Input

	ToVirtualNetworkPrivateLinkServiceNetworkPoliciesOutput() VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput
	ToVirtualNetworkPrivateLinkServiceNetworkPoliciesOutputWithContext(context.Context) VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput
}

var virtualNetworkPrivateLinkServiceNetworkPoliciesPtrType = reflect.TypeOf((**VirtualNetworkPrivateLinkServiceNetworkPolicies)(nil)).Elem()

type VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrInput interface {
	pulumi.Input

	ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput() VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput
	ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(context.Context) VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput
}

type virtualNetworkPrivateLinkServiceNetworkPoliciesPtr string

func VirtualNetworkPrivateLinkServiceNetworkPoliciesPtr(v string) VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrInput {
	return (*virtualNetworkPrivateLinkServiceNetworkPoliciesPtr)(&v)
}

func (*virtualNetworkPrivateLinkServiceNetworkPoliciesPtr) ElementType() reflect.Type {
	return virtualNetworkPrivateLinkServiceNetworkPoliciesPtrType
}

func (in *virtualNetworkPrivateLinkServiceNetworkPoliciesPtr) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput() VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput {
	return pulumi.ToOutput(in).(VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput)
}

func (in *virtualNetworkPrivateLinkServiceNetworkPoliciesPtr) ToVirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx context.Context) VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput)
}

type VpnAuthenticationType string

const (
	VpnAuthenticationTypeCertificate = VpnAuthenticationType("Certificate")
	VpnAuthenticationTypeRadius      = VpnAuthenticationType("Radius")
	VpnAuthenticationTypeAAD         = VpnAuthenticationType("AAD")
)

func (VpnAuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnAuthenticationType)(nil)).Elem()
}

func (e VpnAuthenticationType) ToVpnAuthenticationTypeOutput() VpnAuthenticationTypeOutput {
	return pulumi.ToOutput(e).(VpnAuthenticationTypeOutput)
}

func (e VpnAuthenticationType) ToVpnAuthenticationTypeOutputWithContext(ctx context.Context) VpnAuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VpnAuthenticationTypeOutput)
}

func (e VpnAuthenticationType) ToVpnAuthenticationTypePtrOutput() VpnAuthenticationTypePtrOutput {
	return e.ToVpnAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e VpnAuthenticationType) ToVpnAuthenticationTypePtrOutputWithContext(ctx context.Context) VpnAuthenticationTypePtrOutput {
	return VpnAuthenticationType(e).ToVpnAuthenticationTypeOutputWithContext(ctx).ToVpnAuthenticationTypePtrOutputWithContext(ctx)
}

func (e VpnAuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnAuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnAuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VpnAuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VpnAuthenticationTypeOutput struct{ *pulumi.OutputState }

func (VpnAuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnAuthenticationType)(nil)).Elem()
}

func (o VpnAuthenticationTypeOutput) ToVpnAuthenticationTypeOutput() VpnAuthenticationTypeOutput {
	return o
}

func (o VpnAuthenticationTypeOutput) ToVpnAuthenticationTypeOutputWithContext(ctx context.Context) VpnAuthenticationTypeOutput {
	return o
}

func (o VpnAuthenticationTypeOutput) ToVpnAuthenticationTypePtrOutput() VpnAuthenticationTypePtrOutput {
	return o.ToVpnAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o VpnAuthenticationTypeOutput) ToVpnAuthenticationTypePtrOutputWithContext(ctx context.Context) VpnAuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnAuthenticationType) *VpnAuthenticationType {
		return &v
	}).(VpnAuthenticationTypePtrOutput)
}

func (o VpnAuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VpnAuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnAuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VpnAuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnAuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnAuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VpnAuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (VpnAuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnAuthenticationType)(nil)).Elem()
}

func (o VpnAuthenticationTypePtrOutput) ToVpnAuthenticationTypePtrOutput() VpnAuthenticationTypePtrOutput {
	return o
}

func (o VpnAuthenticationTypePtrOutput) ToVpnAuthenticationTypePtrOutputWithContext(ctx context.Context) VpnAuthenticationTypePtrOutput {
	return o
}

func (o VpnAuthenticationTypePtrOutput) Elem() VpnAuthenticationTypeOutput {
	return o.ApplyT(func(v *VpnAuthenticationType) VpnAuthenticationType {
		if v != nil {
			return *v
		}
		var ret VpnAuthenticationType
		return ret
	}).(VpnAuthenticationTypeOutput)
}

func (o VpnAuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnAuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VpnAuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VpnAuthenticationTypeInput interface {
	pulumi.Input

	ToVpnAuthenticationTypeOutput() VpnAuthenticationTypeOutput
	ToVpnAuthenticationTypeOutputWithContext(context.Context) VpnAuthenticationTypeOutput
}

var vpnAuthenticationTypePtrType = reflect.TypeOf((**VpnAuthenticationType)(nil)).Elem()

type VpnAuthenticationTypePtrInput interface {
	pulumi.Input

	ToVpnAuthenticationTypePtrOutput() VpnAuthenticationTypePtrOutput
	ToVpnAuthenticationTypePtrOutputWithContext(context.Context) VpnAuthenticationTypePtrOutput
}

type vpnAuthenticationTypePtr string

func VpnAuthenticationTypePtr(v string) VpnAuthenticationTypePtrInput {
	return (*vpnAuthenticationTypePtr)(&v)
}

func (*vpnAuthenticationTypePtr) ElementType() reflect.Type {
	return vpnAuthenticationTypePtrType
}

func (in *vpnAuthenticationTypePtr) ToVpnAuthenticationTypePtrOutput() VpnAuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(VpnAuthenticationTypePtrOutput)
}

func (in *vpnAuthenticationTypePtr) ToVpnAuthenticationTypePtrOutputWithContext(ctx context.Context) VpnAuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VpnAuthenticationTypePtrOutput)
}

type VpnClientProtocol string

const (
	VpnClientProtocolIkeV2   = VpnClientProtocol("IkeV2")
	VpnClientProtocolSSTP    = VpnClientProtocol("SSTP")
	VpnClientProtocolOpenVPN = VpnClientProtocol("OpenVPN")
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

type VpnGatewayGeneration string

const (
	VpnGatewayGenerationNone        = VpnGatewayGeneration("None")
	VpnGatewayGenerationGeneration1 = VpnGatewayGeneration("Generation1")
	VpnGatewayGenerationGeneration2 = VpnGatewayGeneration("Generation2")
)

func (VpnGatewayGeneration) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayGeneration)(nil)).Elem()
}

func (e VpnGatewayGeneration) ToVpnGatewayGenerationOutput() VpnGatewayGenerationOutput {
	return pulumi.ToOutput(e).(VpnGatewayGenerationOutput)
}

func (e VpnGatewayGeneration) ToVpnGatewayGenerationOutputWithContext(ctx context.Context) VpnGatewayGenerationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VpnGatewayGenerationOutput)
}

func (e VpnGatewayGeneration) ToVpnGatewayGenerationPtrOutput() VpnGatewayGenerationPtrOutput {
	return e.ToVpnGatewayGenerationPtrOutputWithContext(context.Background())
}

func (e VpnGatewayGeneration) ToVpnGatewayGenerationPtrOutputWithContext(ctx context.Context) VpnGatewayGenerationPtrOutput {
	return VpnGatewayGeneration(e).ToVpnGatewayGenerationOutputWithContext(ctx).ToVpnGatewayGenerationPtrOutputWithContext(ctx)
}

func (e VpnGatewayGeneration) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnGatewayGeneration) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnGatewayGeneration) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VpnGatewayGeneration) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VpnGatewayGenerationOutput struct{ *pulumi.OutputState }

func (VpnGatewayGenerationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayGeneration)(nil)).Elem()
}

func (o VpnGatewayGenerationOutput) ToVpnGatewayGenerationOutput() VpnGatewayGenerationOutput {
	return o
}

func (o VpnGatewayGenerationOutput) ToVpnGatewayGenerationOutputWithContext(ctx context.Context) VpnGatewayGenerationOutput {
	return o
}

func (o VpnGatewayGenerationOutput) ToVpnGatewayGenerationPtrOutput() VpnGatewayGenerationPtrOutput {
	return o.ToVpnGatewayGenerationPtrOutputWithContext(context.Background())
}

func (o VpnGatewayGenerationOutput) ToVpnGatewayGenerationPtrOutputWithContext(ctx context.Context) VpnGatewayGenerationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnGatewayGeneration) *VpnGatewayGeneration {
		return &v
	}).(VpnGatewayGenerationPtrOutput)
}

func (o VpnGatewayGenerationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VpnGatewayGenerationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnGatewayGeneration) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VpnGatewayGenerationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnGatewayGenerationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnGatewayGeneration) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VpnGatewayGenerationPtrOutput struct{ *pulumi.OutputState }

func (VpnGatewayGenerationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnGatewayGeneration)(nil)).Elem()
}

func (o VpnGatewayGenerationPtrOutput) ToVpnGatewayGenerationPtrOutput() VpnGatewayGenerationPtrOutput {
	return o
}

func (o VpnGatewayGenerationPtrOutput) ToVpnGatewayGenerationPtrOutputWithContext(ctx context.Context) VpnGatewayGenerationPtrOutput {
	return o
}

func (o VpnGatewayGenerationPtrOutput) Elem() VpnGatewayGenerationOutput {
	return o.ApplyT(func(v *VpnGatewayGeneration) VpnGatewayGeneration {
		if v != nil {
			return *v
		}
		var ret VpnGatewayGeneration
		return ret
	}).(VpnGatewayGenerationOutput)
}

func (o VpnGatewayGenerationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnGatewayGenerationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VpnGatewayGeneration) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VpnGatewayGenerationInput interface {
	pulumi.Input

	ToVpnGatewayGenerationOutput() VpnGatewayGenerationOutput
	ToVpnGatewayGenerationOutputWithContext(context.Context) VpnGatewayGenerationOutput
}

var vpnGatewayGenerationPtrType = reflect.TypeOf((**VpnGatewayGeneration)(nil)).Elem()

type VpnGatewayGenerationPtrInput interface {
	pulumi.Input

	ToVpnGatewayGenerationPtrOutput() VpnGatewayGenerationPtrOutput
	ToVpnGatewayGenerationPtrOutputWithContext(context.Context) VpnGatewayGenerationPtrOutput
}

type vpnGatewayGenerationPtr string

func VpnGatewayGenerationPtr(v string) VpnGatewayGenerationPtrInput {
	return (*vpnGatewayGenerationPtr)(&v)
}

func (*vpnGatewayGenerationPtr) ElementType() reflect.Type {
	return vpnGatewayGenerationPtrType
}

func (in *vpnGatewayGenerationPtr) ToVpnGatewayGenerationPtrOutput() VpnGatewayGenerationPtrOutput {
	return pulumi.ToOutput(in).(VpnGatewayGenerationPtrOutput)
}

func (in *vpnGatewayGenerationPtr) ToVpnGatewayGenerationPtrOutputWithContext(ctx context.Context) VpnGatewayGenerationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VpnGatewayGenerationPtrOutput)
}

type VpnGatewayTunnelingProtocol string

const (
	VpnGatewayTunnelingProtocolIkeV2   = VpnGatewayTunnelingProtocol("IkeV2")
	VpnGatewayTunnelingProtocolOpenVPN = VpnGatewayTunnelingProtocol("OpenVPN")
)

func (VpnGatewayTunnelingProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayTunnelingProtocol)(nil)).Elem()
}

func (e VpnGatewayTunnelingProtocol) ToVpnGatewayTunnelingProtocolOutput() VpnGatewayTunnelingProtocolOutput {
	return pulumi.ToOutput(e).(VpnGatewayTunnelingProtocolOutput)
}

func (e VpnGatewayTunnelingProtocol) ToVpnGatewayTunnelingProtocolOutputWithContext(ctx context.Context) VpnGatewayTunnelingProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VpnGatewayTunnelingProtocolOutput)
}

func (e VpnGatewayTunnelingProtocol) ToVpnGatewayTunnelingProtocolPtrOutput() VpnGatewayTunnelingProtocolPtrOutput {
	return e.ToVpnGatewayTunnelingProtocolPtrOutputWithContext(context.Background())
}

func (e VpnGatewayTunnelingProtocol) ToVpnGatewayTunnelingProtocolPtrOutputWithContext(ctx context.Context) VpnGatewayTunnelingProtocolPtrOutput {
	return VpnGatewayTunnelingProtocol(e).ToVpnGatewayTunnelingProtocolOutputWithContext(ctx).ToVpnGatewayTunnelingProtocolPtrOutputWithContext(ctx)
}

func (e VpnGatewayTunnelingProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnGatewayTunnelingProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnGatewayTunnelingProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VpnGatewayTunnelingProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VpnGatewayTunnelingProtocolOutput struct{ *pulumi.OutputState }

func (VpnGatewayTunnelingProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayTunnelingProtocol)(nil)).Elem()
}

func (o VpnGatewayTunnelingProtocolOutput) ToVpnGatewayTunnelingProtocolOutput() VpnGatewayTunnelingProtocolOutput {
	return o
}

func (o VpnGatewayTunnelingProtocolOutput) ToVpnGatewayTunnelingProtocolOutputWithContext(ctx context.Context) VpnGatewayTunnelingProtocolOutput {
	return o
}

func (o VpnGatewayTunnelingProtocolOutput) ToVpnGatewayTunnelingProtocolPtrOutput() VpnGatewayTunnelingProtocolPtrOutput {
	return o.ToVpnGatewayTunnelingProtocolPtrOutputWithContext(context.Background())
}

func (o VpnGatewayTunnelingProtocolOutput) ToVpnGatewayTunnelingProtocolPtrOutputWithContext(ctx context.Context) VpnGatewayTunnelingProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnGatewayTunnelingProtocol) *VpnGatewayTunnelingProtocol {
		return &v
	}).(VpnGatewayTunnelingProtocolPtrOutput)
}

func (o VpnGatewayTunnelingProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VpnGatewayTunnelingProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnGatewayTunnelingProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VpnGatewayTunnelingProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnGatewayTunnelingProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnGatewayTunnelingProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VpnGatewayTunnelingProtocolPtrOutput struct{ *pulumi.OutputState }

func (VpnGatewayTunnelingProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnGatewayTunnelingProtocol)(nil)).Elem()
}

func (o VpnGatewayTunnelingProtocolPtrOutput) ToVpnGatewayTunnelingProtocolPtrOutput() VpnGatewayTunnelingProtocolPtrOutput {
	return o
}

func (o VpnGatewayTunnelingProtocolPtrOutput) ToVpnGatewayTunnelingProtocolPtrOutputWithContext(ctx context.Context) VpnGatewayTunnelingProtocolPtrOutput {
	return o
}

func (o VpnGatewayTunnelingProtocolPtrOutput) Elem() VpnGatewayTunnelingProtocolOutput {
	return o.ApplyT(func(v *VpnGatewayTunnelingProtocol) VpnGatewayTunnelingProtocol {
		if v != nil {
			return *v
		}
		var ret VpnGatewayTunnelingProtocol
		return ret
	}).(VpnGatewayTunnelingProtocolOutput)
}

func (o VpnGatewayTunnelingProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnGatewayTunnelingProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VpnGatewayTunnelingProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VpnGatewayTunnelingProtocolInput interface {
	pulumi.Input

	ToVpnGatewayTunnelingProtocolOutput() VpnGatewayTunnelingProtocolOutput
	ToVpnGatewayTunnelingProtocolOutputWithContext(context.Context) VpnGatewayTunnelingProtocolOutput
}

var vpnGatewayTunnelingProtocolPtrType = reflect.TypeOf((**VpnGatewayTunnelingProtocol)(nil)).Elem()

type VpnGatewayTunnelingProtocolPtrInput interface {
	pulumi.Input

	ToVpnGatewayTunnelingProtocolPtrOutput() VpnGatewayTunnelingProtocolPtrOutput
	ToVpnGatewayTunnelingProtocolPtrOutputWithContext(context.Context) VpnGatewayTunnelingProtocolPtrOutput
}

type vpnGatewayTunnelingProtocolPtr string

func VpnGatewayTunnelingProtocolPtr(v string) VpnGatewayTunnelingProtocolPtrInput {
	return (*vpnGatewayTunnelingProtocolPtr)(&v)
}

func (*vpnGatewayTunnelingProtocolPtr) ElementType() reflect.Type {
	return vpnGatewayTunnelingProtocolPtrType
}

func (in *vpnGatewayTunnelingProtocolPtr) ToVpnGatewayTunnelingProtocolPtrOutput() VpnGatewayTunnelingProtocolPtrOutput {
	return pulumi.ToOutput(in).(VpnGatewayTunnelingProtocolPtrOutput)
}

func (in *vpnGatewayTunnelingProtocolPtr) ToVpnGatewayTunnelingProtocolPtrOutputWithContext(ctx context.Context) VpnGatewayTunnelingProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VpnGatewayTunnelingProtocolPtrOutput)
}

type VpnLinkConnectionMode string

const (
	VpnLinkConnectionModeDefault       = VpnLinkConnectionMode("Default")
	VpnLinkConnectionModeResponderOnly = VpnLinkConnectionMode("ResponderOnly")
	VpnLinkConnectionModeInitiatorOnly = VpnLinkConnectionMode("InitiatorOnly")
)

func (VpnLinkConnectionMode) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkConnectionMode)(nil)).Elem()
}

func (e VpnLinkConnectionMode) ToVpnLinkConnectionModeOutput() VpnLinkConnectionModeOutput {
	return pulumi.ToOutput(e).(VpnLinkConnectionModeOutput)
}

func (e VpnLinkConnectionMode) ToVpnLinkConnectionModeOutputWithContext(ctx context.Context) VpnLinkConnectionModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VpnLinkConnectionModeOutput)
}

func (e VpnLinkConnectionMode) ToVpnLinkConnectionModePtrOutput() VpnLinkConnectionModePtrOutput {
	return e.ToVpnLinkConnectionModePtrOutputWithContext(context.Background())
}

func (e VpnLinkConnectionMode) ToVpnLinkConnectionModePtrOutputWithContext(ctx context.Context) VpnLinkConnectionModePtrOutput {
	return VpnLinkConnectionMode(e).ToVpnLinkConnectionModeOutputWithContext(ctx).ToVpnLinkConnectionModePtrOutputWithContext(ctx)
}

func (e VpnLinkConnectionMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnLinkConnectionMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnLinkConnectionMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VpnLinkConnectionMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VpnLinkConnectionModeOutput struct{ *pulumi.OutputState }

func (VpnLinkConnectionModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkConnectionMode)(nil)).Elem()
}

func (o VpnLinkConnectionModeOutput) ToVpnLinkConnectionModeOutput() VpnLinkConnectionModeOutput {
	return o
}

func (o VpnLinkConnectionModeOutput) ToVpnLinkConnectionModeOutputWithContext(ctx context.Context) VpnLinkConnectionModeOutput {
	return o
}

func (o VpnLinkConnectionModeOutput) ToVpnLinkConnectionModePtrOutput() VpnLinkConnectionModePtrOutput {
	return o.ToVpnLinkConnectionModePtrOutputWithContext(context.Background())
}

func (o VpnLinkConnectionModeOutput) ToVpnLinkConnectionModePtrOutputWithContext(ctx context.Context) VpnLinkConnectionModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnLinkConnectionMode) *VpnLinkConnectionMode {
		return &v
	}).(VpnLinkConnectionModePtrOutput)
}

func (o VpnLinkConnectionModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VpnLinkConnectionModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnLinkConnectionMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VpnLinkConnectionModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnLinkConnectionModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnLinkConnectionMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VpnLinkConnectionModePtrOutput struct{ *pulumi.OutputState }

func (VpnLinkConnectionModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkConnectionMode)(nil)).Elem()
}

func (o VpnLinkConnectionModePtrOutput) ToVpnLinkConnectionModePtrOutput() VpnLinkConnectionModePtrOutput {
	return o
}

func (o VpnLinkConnectionModePtrOutput) ToVpnLinkConnectionModePtrOutputWithContext(ctx context.Context) VpnLinkConnectionModePtrOutput {
	return o
}

func (o VpnLinkConnectionModePtrOutput) Elem() VpnLinkConnectionModeOutput {
	return o.ApplyT(func(v *VpnLinkConnectionMode) VpnLinkConnectionMode {
		if v != nil {
			return *v
		}
		var ret VpnLinkConnectionMode
		return ret
	}).(VpnLinkConnectionModeOutput)
}

func (o VpnLinkConnectionModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnLinkConnectionModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VpnLinkConnectionMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VpnLinkConnectionModeInput interface {
	pulumi.Input

	ToVpnLinkConnectionModeOutput() VpnLinkConnectionModeOutput
	ToVpnLinkConnectionModeOutputWithContext(context.Context) VpnLinkConnectionModeOutput
}

var vpnLinkConnectionModePtrType = reflect.TypeOf((**VpnLinkConnectionMode)(nil)).Elem()

type VpnLinkConnectionModePtrInput interface {
	pulumi.Input

	ToVpnLinkConnectionModePtrOutput() VpnLinkConnectionModePtrOutput
	ToVpnLinkConnectionModePtrOutputWithContext(context.Context) VpnLinkConnectionModePtrOutput
}

type vpnLinkConnectionModePtr string

func VpnLinkConnectionModePtr(v string) VpnLinkConnectionModePtrInput {
	return (*vpnLinkConnectionModePtr)(&v)
}

func (*vpnLinkConnectionModePtr) ElementType() reflect.Type {
	return vpnLinkConnectionModePtrType
}

func (in *vpnLinkConnectionModePtr) ToVpnLinkConnectionModePtrOutput() VpnLinkConnectionModePtrOutput {
	return pulumi.ToOutput(in).(VpnLinkConnectionModePtrOutput)
}

func (in *vpnLinkConnectionModePtr) ToVpnLinkConnectionModePtrOutputWithContext(ctx context.Context) VpnLinkConnectionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VpnLinkConnectionModePtrOutput)
}

type VpnNatRuleMode string

const (
	VpnNatRuleModeEgressSnat  = VpnNatRuleMode("EgressSnat")
	VpnNatRuleModeIngressSnat = VpnNatRuleMode("IngressSnat")
)

func (VpnNatRuleMode) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMode)(nil)).Elem()
}

func (e VpnNatRuleMode) ToVpnNatRuleModeOutput() VpnNatRuleModeOutput {
	return pulumi.ToOutput(e).(VpnNatRuleModeOutput)
}

func (e VpnNatRuleMode) ToVpnNatRuleModeOutputWithContext(ctx context.Context) VpnNatRuleModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VpnNatRuleModeOutput)
}

func (e VpnNatRuleMode) ToVpnNatRuleModePtrOutput() VpnNatRuleModePtrOutput {
	return e.ToVpnNatRuleModePtrOutputWithContext(context.Background())
}

func (e VpnNatRuleMode) ToVpnNatRuleModePtrOutputWithContext(ctx context.Context) VpnNatRuleModePtrOutput {
	return VpnNatRuleMode(e).ToVpnNatRuleModeOutputWithContext(ctx).ToVpnNatRuleModePtrOutputWithContext(ctx)
}

func (e VpnNatRuleMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnNatRuleMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnNatRuleMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VpnNatRuleMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VpnNatRuleModeOutput struct{ *pulumi.OutputState }

func (VpnNatRuleModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMode)(nil)).Elem()
}

func (o VpnNatRuleModeOutput) ToVpnNatRuleModeOutput() VpnNatRuleModeOutput {
	return o
}

func (o VpnNatRuleModeOutput) ToVpnNatRuleModeOutputWithContext(ctx context.Context) VpnNatRuleModeOutput {
	return o
}

func (o VpnNatRuleModeOutput) ToVpnNatRuleModePtrOutput() VpnNatRuleModePtrOutput {
	return o.ToVpnNatRuleModePtrOutputWithContext(context.Background())
}

func (o VpnNatRuleModeOutput) ToVpnNatRuleModePtrOutputWithContext(ctx context.Context) VpnNatRuleModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnNatRuleMode) *VpnNatRuleMode {
		return &v
	}).(VpnNatRuleModePtrOutput)
}

func (o VpnNatRuleModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VpnNatRuleModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnNatRuleMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VpnNatRuleModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnNatRuleModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnNatRuleMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VpnNatRuleModePtrOutput struct{ *pulumi.OutputState }

func (VpnNatRuleModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnNatRuleMode)(nil)).Elem()
}

func (o VpnNatRuleModePtrOutput) ToVpnNatRuleModePtrOutput() VpnNatRuleModePtrOutput {
	return o
}

func (o VpnNatRuleModePtrOutput) ToVpnNatRuleModePtrOutputWithContext(ctx context.Context) VpnNatRuleModePtrOutput {
	return o
}

func (o VpnNatRuleModePtrOutput) Elem() VpnNatRuleModeOutput {
	return o.ApplyT(func(v *VpnNatRuleMode) VpnNatRuleMode {
		if v != nil {
			return *v
		}
		var ret VpnNatRuleMode
		return ret
	}).(VpnNatRuleModeOutput)
}

func (o VpnNatRuleModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnNatRuleModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VpnNatRuleMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VpnNatRuleModeInput interface {
	pulumi.Input

	ToVpnNatRuleModeOutput() VpnNatRuleModeOutput
	ToVpnNatRuleModeOutputWithContext(context.Context) VpnNatRuleModeOutput
}

var vpnNatRuleModePtrType = reflect.TypeOf((**VpnNatRuleMode)(nil)).Elem()

type VpnNatRuleModePtrInput interface {
	pulumi.Input

	ToVpnNatRuleModePtrOutput() VpnNatRuleModePtrOutput
	ToVpnNatRuleModePtrOutputWithContext(context.Context) VpnNatRuleModePtrOutput
}

type vpnNatRuleModePtr string

func VpnNatRuleModePtr(v string) VpnNatRuleModePtrInput {
	return (*vpnNatRuleModePtr)(&v)
}

func (*vpnNatRuleModePtr) ElementType() reflect.Type {
	return vpnNatRuleModePtrType
}

func (in *vpnNatRuleModePtr) ToVpnNatRuleModePtrOutput() VpnNatRuleModePtrOutput {
	return pulumi.ToOutput(in).(VpnNatRuleModePtrOutput)
}

func (in *vpnNatRuleModePtr) ToVpnNatRuleModePtrOutputWithContext(ctx context.Context) VpnNatRuleModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VpnNatRuleModePtrOutput)
}

type VpnNatRuleType string

const (
	VpnNatRuleTypeStatic  = VpnNatRuleType("Static")
	VpnNatRuleTypeDynamic = VpnNatRuleType("Dynamic")
)

func (VpnNatRuleType) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleType)(nil)).Elem()
}

func (e VpnNatRuleType) ToVpnNatRuleTypeOutput() VpnNatRuleTypeOutput {
	return pulumi.ToOutput(e).(VpnNatRuleTypeOutput)
}

func (e VpnNatRuleType) ToVpnNatRuleTypeOutputWithContext(ctx context.Context) VpnNatRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VpnNatRuleTypeOutput)
}

func (e VpnNatRuleType) ToVpnNatRuleTypePtrOutput() VpnNatRuleTypePtrOutput {
	return e.ToVpnNatRuleTypePtrOutputWithContext(context.Background())
}

func (e VpnNatRuleType) ToVpnNatRuleTypePtrOutputWithContext(ctx context.Context) VpnNatRuleTypePtrOutput {
	return VpnNatRuleType(e).ToVpnNatRuleTypeOutputWithContext(ctx).ToVpnNatRuleTypePtrOutputWithContext(ctx)
}

func (e VpnNatRuleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnNatRuleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VpnNatRuleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VpnNatRuleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VpnNatRuleTypeOutput struct{ *pulumi.OutputState }

func (VpnNatRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleType)(nil)).Elem()
}

func (o VpnNatRuleTypeOutput) ToVpnNatRuleTypeOutput() VpnNatRuleTypeOutput {
	return o
}

func (o VpnNatRuleTypeOutput) ToVpnNatRuleTypeOutputWithContext(ctx context.Context) VpnNatRuleTypeOutput {
	return o
}

func (o VpnNatRuleTypeOutput) ToVpnNatRuleTypePtrOutput() VpnNatRuleTypePtrOutput {
	return o.ToVpnNatRuleTypePtrOutputWithContext(context.Background())
}

func (o VpnNatRuleTypeOutput) ToVpnNatRuleTypePtrOutputWithContext(ctx context.Context) VpnNatRuleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnNatRuleType) *VpnNatRuleType {
		return &v
	}).(VpnNatRuleTypePtrOutput)
}

func (o VpnNatRuleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VpnNatRuleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnNatRuleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VpnNatRuleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnNatRuleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VpnNatRuleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VpnNatRuleTypePtrOutput struct{ *pulumi.OutputState }

func (VpnNatRuleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnNatRuleType)(nil)).Elem()
}

func (o VpnNatRuleTypePtrOutput) ToVpnNatRuleTypePtrOutput() VpnNatRuleTypePtrOutput {
	return o
}

func (o VpnNatRuleTypePtrOutput) ToVpnNatRuleTypePtrOutputWithContext(ctx context.Context) VpnNatRuleTypePtrOutput {
	return o
}

func (o VpnNatRuleTypePtrOutput) Elem() VpnNatRuleTypeOutput {
	return o.ApplyT(func(v *VpnNatRuleType) VpnNatRuleType {
		if v != nil {
			return *v
		}
		var ret VpnNatRuleType
		return ret
	}).(VpnNatRuleTypeOutput)
}

func (o VpnNatRuleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VpnNatRuleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VpnNatRuleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VpnNatRuleTypeInput interface {
	pulumi.Input

	ToVpnNatRuleTypeOutput() VpnNatRuleTypeOutput
	ToVpnNatRuleTypeOutputWithContext(context.Context) VpnNatRuleTypeOutput
}

var vpnNatRuleTypePtrType = reflect.TypeOf((**VpnNatRuleType)(nil)).Elem()

type VpnNatRuleTypePtrInput interface {
	pulumi.Input

	ToVpnNatRuleTypePtrOutput() VpnNatRuleTypePtrOutput
	ToVpnNatRuleTypePtrOutputWithContext(context.Context) VpnNatRuleTypePtrOutput
}

type vpnNatRuleTypePtr string

func VpnNatRuleTypePtr(v string) VpnNatRuleTypePtrInput {
	return (*vpnNatRuleTypePtr)(&v)
}

func (*vpnNatRuleTypePtr) ElementType() reflect.Type {
	return vpnNatRuleTypePtrType
}

func (in *vpnNatRuleTypePtr) ToVpnNatRuleTypePtrOutput() VpnNatRuleTypePtrOutput {
	return pulumi.ToOutput(in).(VpnNatRuleTypePtrOutput)
}

func (in *vpnNatRuleTypePtr) ToVpnNatRuleTypePtrOutputWithContext(ctx context.Context) VpnNatRuleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VpnNatRuleTypePtrOutput)
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

type WebApplicationFirewallAction string

const (
	WebApplicationFirewallActionAllow = WebApplicationFirewallAction("Allow")
	WebApplicationFirewallActionBlock = WebApplicationFirewallAction("Block")
	WebApplicationFirewallActionLog   = WebApplicationFirewallAction("Log")
)

func (WebApplicationFirewallAction) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallAction)(nil)).Elem()
}

func (e WebApplicationFirewallAction) ToWebApplicationFirewallActionOutput() WebApplicationFirewallActionOutput {
	return pulumi.ToOutput(e).(WebApplicationFirewallActionOutput)
}

func (e WebApplicationFirewallAction) ToWebApplicationFirewallActionOutputWithContext(ctx context.Context) WebApplicationFirewallActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebApplicationFirewallActionOutput)
}

func (e WebApplicationFirewallAction) ToWebApplicationFirewallActionPtrOutput() WebApplicationFirewallActionPtrOutput {
	return e.ToWebApplicationFirewallActionPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallAction) ToWebApplicationFirewallActionPtrOutputWithContext(ctx context.Context) WebApplicationFirewallActionPtrOutput {
	return WebApplicationFirewallAction(e).ToWebApplicationFirewallActionOutputWithContext(ctx).ToWebApplicationFirewallActionPtrOutputWithContext(ctx)
}

func (e WebApplicationFirewallAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebApplicationFirewallActionOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallAction)(nil)).Elem()
}

func (o WebApplicationFirewallActionOutput) ToWebApplicationFirewallActionOutput() WebApplicationFirewallActionOutput {
	return o
}

func (o WebApplicationFirewallActionOutput) ToWebApplicationFirewallActionOutputWithContext(ctx context.Context) WebApplicationFirewallActionOutput {
	return o
}

func (o WebApplicationFirewallActionOutput) ToWebApplicationFirewallActionPtrOutput() WebApplicationFirewallActionPtrOutput {
	return o.ToWebApplicationFirewallActionPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallActionOutput) ToWebApplicationFirewallActionPtrOutputWithContext(ctx context.Context) WebApplicationFirewallActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebApplicationFirewallAction) *WebApplicationFirewallAction {
		return &v
	}).(WebApplicationFirewallActionPtrOutput)
}

func (o WebApplicationFirewallActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebApplicationFirewallActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebApplicationFirewallActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebApplicationFirewallActionPtrOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallAction)(nil)).Elem()
}

func (o WebApplicationFirewallActionPtrOutput) ToWebApplicationFirewallActionPtrOutput() WebApplicationFirewallActionPtrOutput {
	return o
}

func (o WebApplicationFirewallActionPtrOutput) ToWebApplicationFirewallActionPtrOutputWithContext(ctx context.Context) WebApplicationFirewallActionPtrOutput {
	return o
}

func (o WebApplicationFirewallActionPtrOutput) Elem() WebApplicationFirewallActionOutput {
	return o.ApplyT(func(v *WebApplicationFirewallAction) WebApplicationFirewallAction {
		if v != nil {
			return *v
		}
		var ret WebApplicationFirewallAction
		return ret
	}).(WebApplicationFirewallActionOutput)
}

func (o WebApplicationFirewallActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebApplicationFirewallAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebApplicationFirewallActionInput interface {
	pulumi.Input

	ToWebApplicationFirewallActionOutput() WebApplicationFirewallActionOutput
	ToWebApplicationFirewallActionOutputWithContext(context.Context) WebApplicationFirewallActionOutput
}

var webApplicationFirewallActionPtrType = reflect.TypeOf((**WebApplicationFirewallAction)(nil)).Elem()

type WebApplicationFirewallActionPtrInput interface {
	pulumi.Input

	ToWebApplicationFirewallActionPtrOutput() WebApplicationFirewallActionPtrOutput
	ToWebApplicationFirewallActionPtrOutputWithContext(context.Context) WebApplicationFirewallActionPtrOutput
}

type webApplicationFirewallActionPtr string

func WebApplicationFirewallActionPtr(v string) WebApplicationFirewallActionPtrInput {
	return (*webApplicationFirewallActionPtr)(&v)
}

func (*webApplicationFirewallActionPtr) ElementType() reflect.Type {
	return webApplicationFirewallActionPtrType
}

func (in *webApplicationFirewallActionPtr) ToWebApplicationFirewallActionPtrOutput() WebApplicationFirewallActionPtrOutput {
	return pulumi.ToOutput(in).(WebApplicationFirewallActionPtrOutput)
}

func (in *webApplicationFirewallActionPtr) ToWebApplicationFirewallActionPtrOutputWithContext(ctx context.Context) WebApplicationFirewallActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebApplicationFirewallActionPtrOutput)
}

type WebApplicationFirewallEnabledState string

const (
	WebApplicationFirewallEnabledStateDisabled = WebApplicationFirewallEnabledState("Disabled")
	WebApplicationFirewallEnabledStateEnabled  = WebApplicationFirewallEnabledState("Enabled")
)

func (WebApplicationFirewallEnabledState) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallEnabledState)(nil)).Elem()
}

func (e WebApplicationFirewallEnabledState) ToWebApplicationFirewallEnabledStateOutput() WebApplicationFirewallEnabledStateOutput {
	return pulumi.ToOutput(e).(WebApplicationFirewallEnabledStateOutput)
}

func (e WebApplicationFirewallEnabledState) ToWebApplicationFirewallEnabledStateOutputWithContext(ctx context.Context) WebApplicationFirewallEnabledStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebApplicationFirewallEnabledStateOutput)
}

func (e WebApplicationFirewallEnabledState) ToWebApplicationFirewallEnabledStatePtrOutput() WebApplicationFirewallEnabledStatePtrOutput {
	return e.ToWebApplicationFirewallEnabledStatePtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallEnabledState) ToWebApplicationFirewallEnabledStatePtrOutputWithContext(ctx context.Context) WebApplicationFirewallEnabledStatePtrOutput {
	return WebApplicationFirewallEnabledState(e).ToWebApplicationFirewallEnabledStateOutputWithContext(ctx).ToWebApplicationFirewallEnabledStatePtrOutputWithContext(ctx)
}

func (e WebApplicationFirewallEnabledState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallEnabledState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallEnabledState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallEnabledState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebApplicationFirewallEnabledStateOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallEnabledStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallEnabledState)(nil)).Elem()
}

func (o WebApplicationFirewallEnabledStateOutput) ToWebApplicationFirewallEnabledStateOutput() WebApplicationFirewallEnabledStateOutput {
	return o
}

func (o WebApplicationFirewallEnabledStateOutput) ToWebApplicationFirewallEnabledStateOutputWithContext(ctx context.Context) WebApplicationFirewallEnabledStateOutput {
	return o
}

func (o WebApplicationFirewallEnabledStateOutput) ToWebApplicationFirewallEnabledStatePtrOutput() WebApplicationFirewallEnabledStatePtrOutput {
	return o.ToWebApplicationFirewallEnabledStatePtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallEnabledStateOutput) ToWebApplicationFirewallEnabledStatePtrOutputWithContext(ctx context.Context) WebApplicationFirewallEnabledStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebApplicationFirewallEnabledState) *WebApplicationFirewallEnabledState {
		return &v
	}).(WebApplicationFirewallEnabledStatePtrOutput)
}

func (o WebApplicationFirewallEnabledStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebApplicationFirewallEnabledStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallEnabledState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebApplicationFirewallEnabledStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallEnabledStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallEnabledState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebApplicationFirewallEnabledStatePtrOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallEnabledStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallEnabledState)(nil)).Elem()
}

func (o WebApplicationFirewallEnabledStatePtrOutput) ToWebApplicationFirewallEnabledStatePtrOutput() WebApplicationFirewallEnabledStatePtrOutput {
	return o
}

func (o WebApplicationFirewallEnabledStatePtrOutput) ToWebApplicationFirewallEnabledStatePtrOutputWithContext(ctx context.Context) WebApplicationFirewallEnabledStatePtrOutput {
	return o
}

func (o WebApplicationFirewallEnabledStatePtrOutput) Elem() WebApplicationFirewallEnabledStateOutput {
	return o.ApplyT(func(v *WebApplicationFirewallEnabledState) WebApplicationFirewallEnabledState {
		if v != nil {
			return *v
		}
		var ret WebApplicationFirewallEnabledState
		return ret
	}).(WebApplicationFirewallEnabledStateOutput)
}

func (o WebApplicationFirewallEnabledStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallEnabledStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebApplicationFirewallEnabledState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebApplicationFirewallEnabledStateInput interface {
	pulumi.Input

	ToWebApplicationFirewallEnabledStateOutput() WebApplicationFirewallEnabledStateOutput
	ToWebApplicationFirewallEnabledStateOutputWithContext(context.Context) WebApplicationFirewallEnabledStateOutput
}

var webApplicationFirewallEnabledStatePtrType = reflect.TypeOf((**WebApplicationFirewallEnabledState)(nil)).Elem()

type WebApplicationFirewallEnabledStatePtrInput interface {
	pulumi.Input

	ToWebApplicationFirewallEnabledStatePtrOutput() WebApplicationFirewallEnabledStatePtrOutput
	ToWebApplicationFirewallEnabledStatePtrOutputWithContext(context.Context) WebApplicationFirewallEnabledStatePtrOutput
}

type webApplicationFirewallEnabledStatePtr string

func WebApplicationFirewallEnabledStatePtr(v string) WebApplicationFirewallEnabledStatePtrInput {
	return (*webApplicationFirewallEnabledStatePtr)(&v)
}

func (*webApplicationFirewallEnabledStatePtr) ElementType() reflect.Type {
	return webApplicationFirewallEnabledStatePtrType
}

func (in *webApplicationFirewallEnabledStatePtr) ToWebApplicationFirewallEnabledStatePtrOutput() WebApplicationFirewallEnabledStatePtrOutput {
	return pulumi.ToOutput(in).(WebApplicationFirewallEnabledStatePtrOutput)
}

func (in *webApplicationFirewallEnabledStatePtr) ToWebApplicationFirewallEnabledStatePtrOutputWithContext(ctx context.Context) WebApplicationFirewallEnabledStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebApplicationFirewallEnabledStatePtrOutput)
}

type WebApplicationFirewallMatchVariable string

const (
	WebApplicationFirewallMatchVariableRemoteAddr     = WebApplicationFirewallMatchVariable("RemoteAddr")
	WebApplicationFirewallMatchVariableRequestMethod  = WebApplicationFirewallMatchVariable("RequestMethod")
	WebApplicationFirewallMatchVariableQueryString    = WebApplicationFirewallMatchVariable("QueryString")
	WebApplicationFirewallMatchVariablePostArgs       = WebApplicationFirewallMatchVariable("PostArgs")
	WebApplicationFirewallMatchVariableRequestUri     = WebApplicationFirewallMatchVariable("RequestUri")
	WebApplicationFirewallMatchVariableRequestHeaders = WebApplicationFirewallMatchVariable("RequestHeaders")
	WebApplicationFirewallMatchVariableRequestBody    = WebApplicationFirewallMatchVariable("RequestBody")
	WebApplicationFirewallMatchVariableRequestCookies = WebApplicationFirewallMatchVariable("RequestCookies")
)

func (WebApplicationFirewallMatchVariable) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallMatchVariable)(nil)).Elem()
}

func (e WebApplicationFirewallMatchVariable) ToWebApplicationFirewallMatchVariableOutput() WebApplicationFirewallMatchVariableOutput {
	return pulumi.ToOutput(e).(WebApplicationFirewallMatchVariableOutput)
}

func (e WebApplicationFirewallMatchVariable) ToWebApplicationFirewallMatchVariableOutputWithContext(ctx context.Context) WebApplicationFirewallMatchVariableOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebApplicationFirewallMatchVariableOutput)
}

func (e WebApplicationFirewallMatchVariable) ToWebApplicationFirewallMatchVariablePtrOutput() WebApplicationFirewallMatchVariablePtrOutput {
	return e.ToWebApplicationFirewallMatchVariablePtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallMatchVariable) ToWebApplicationFirewallMatchVariablePtrOutputWithContext(ctx context.Context) WebApplicationFirewallMatchVariablePtrOutput {
	return WebApplicationFirewallMatchVariable(e).ToWebApplicationFirewallMatchVariableOutputWithContext(ctx).ToWebApplicationFirewallMatchVariablePtrOutputWithContext(ctx)
}

func (e WebApplicationFirewallMatchVariable) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallMatchVariable) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallMatchVariable) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallMatchVariable) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebApplicationFirewallMatchVariableOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallMatchVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallMatchVariable)(nil)).Elem()
}

func (o WebApplicationFirewallMatchVariableOutput) ToWebApplicationFirewallMatchVariableOutput() WebApplicationFirewallMatchVariableOutput {
	return o
}

func (o WebApplicationFirewallMatchVariableOutput) ToWebApplicationFirewallMatchVariableOutputWithContext(ctx context.Context) WebApplicationFirewallMatchVariableOutput {
	return o
}

func (o WebApplicationFirewallMatchVariableOutput) ToWebApplicationFirewallMatchVariablePtrOutput() WebApplicationFirewallMatchVariablePtrOutput {
	return o.ToWebApplicationFirewallMatchVariablePtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallMatchVariableOutput) ToWebApplicationFirewallMatchVariablePtrOutputWithContext(ctx context.Context) WebApplicationFirewallMatchVariablePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebApplicationFirewallMatchVariable) *WebApplicationFirewallMatchVariable {
		return &v
	}).(WebApplicationFirewallMatchVariablePtrOutput)
}

func (o WebApplicationFirewallMatchVariableOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebApplicationFirewallMatchVariableOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallMatchVariable) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebApplicationFirewallMatchVariableOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallMatchVariableOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallMatchVariable) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebApplicationFirewallMatchVariablePtrOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallMatchVariablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallMatchVariable)(nil)).Elem()
}

func (o WebApplicationFirewallMatchVariablePtrOutput) ToWebApplicationFirewallMatchVariablePtrOutput() WebApplicationFirewallMatchVariablePtrOutput {
	return o
}

func (o WebApplicationFirewallMatchVariablePtrOutput) ToWebApplicationFirewallMatchVariablePtrOutputWithContext(ctx context.Context) WebApplicationFirewallMatchVariablePtrOutput {
	return o
}

func (o WebApplicationFirewallMatchVariablePtrOutput) Elem() WebApplicationFirewallMatchVariableOutput {
	return o.ApplyT(func(v *WebApplicationFirewallMatchVariable) WebApplicationFirewallMatchVariable {
		if v != nil {
			return *v
		}
		var ret WebApplicationFirewallMatchVariable
		return ret
	}).(WebApplicationFirewallMatchVariableOutput)
}

func (o WebApplicationFirewallMatchVariablePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallMatchVariablePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebApplicationFirewallMatchVariable) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebApplicationFirewallMatchVariableInput interface {
	pulumi.Input

	ToWebApplicationFirewallMatchVariableOutput() WebApplicationFirewallMatchVariableOutput
	ToWebApplicationFirewallMatchVariableOutputWithContext(context.Context) WebApplicationFirewallMatchVariableOutput
}

var webApplicationFirewallMatchVariablePtrType = reflect.TypeOf((**WebApplicationFirewallMatchVariable)(nil)).Elem()

type WebApplicationFirewallMatchVariablePtrInput interface {
	pulumi.Input

	ToWebApplicationFirewallMatchVariablePtrOutput() WebApplicationFirewallMatchVariablePtrOutput
	ToWebApplicationFirewallMatchVariablePtrOutputWithContext(context.Context) WebApplicationFirewallMatchVariablePtrOutput
}

type webApplicationFirewallMatchVariablePtr string

func WebApplicationFirewallMatchVariablePtr(v string) WebApplicationFirewallMatchVariablePtrInput {
	return (*webApplicationFirewallMatchVariablePtr)(&v)
}

func (*webApplicationFirewallMatchVariablePtr) ElementType() reflect.Type {
	return webApplicationFirewallMatchVariablePtrType
}

func (in *webApplicationFirewallMatchVariablePtr) ToWebApplicationFirewallMatchVariablePtrOutput() WebApplicationFirewallMatchVariablePtrOutput {
	return pulumi.ToOutput(in).(WebApplicationFirewallMatchVariablePtrOutput)
}

func (in *webApplicationFirewallMatchVariablePtr) ToWebApplicationFirewallMatchVariablePtrOutputWithContext(ctx context.Context) WebApplicationFirewallMatchVariablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebApplicationFirewallMatchVariablePtrOutput)
}

type WebApplicationFirewallMode string

const (
	WebApplicationFirewallModePrevention = WebApplicationFirewallMode("Prevention")
	WebApplicationFirewallModeDetection  = WebApplicationFirewallMode("Detection")
)

func (WebApplicationFirewallMode) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallMode)(nil)).Elem()
}

func (e WebApplicationFirewallMode) ToWebApplicationFirewallModeOutput() WebApplicationFirewallModeOutput {
	return pulumi.ToOutput(e).(WebApplicationFirewallModeOutput)
}

func (e WebApplicationFirewallMode) ToWebApplicationFirewallModeOutputWithContext(ctx context.Context) WebApplicationFirewallModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebApplicationFirewallModeOutput)
}

func (e WebApplicationFirewallMode) ToWebApplicationFirewallModePtrOutput() WebApplicationFirewallModePtrOutput {
	return e.ToWebApplicationFirewallModePtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallMode) ToWebApplicationFirewallModePtrOutputWithContext(ctx context.Context) WebApplicationFirewallModePtrOutput {
	return WebApplicationFirewallMode(e).ToWebApplicationFirewallModeOutputWithContext(ctx).ToWebApplicationFirewallModePtrOutputWithContext(ctx)
}

func (e WebApplicationFirewallMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebApplicationFirewallModeOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallMode)(nil)).Elem()
}

func (o WebApplicationFirewallModeOutput) ToWebApplicationFirewallModeOutput() WebApplicationFirewallModeOutput {
	return o
}

func (o WebApplicationFirewallModeOutput) ToWebApplicationFirewallModeOutputWithContext(ctx context.Context) WebApplicationFirewallModeOutput {
	return o
}

func (o WebApplicationFirewallModeOutput) ToWebApplicationFirewallModePtrOutput() WebApplicationFirewallModePtrOutput {
	return o.ToWebApplicationFirewallModePtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallModeOutput) ToWebApplicationFirewallModePtrOutputWithContext(ctx context.Context) WebApplicationFirewallModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebApplicationFirewallMode) *WebApplicationFirewallMode {
		return &v
	}).(WebApplicationFirewallModePtrOutput)
}

func (o WebApplicationFirewallModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebApplicationFirewallModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebApplicationFirewallModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebApplicationFirewallModePtrOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallMode)(nil)).Elem()
}

func (o WebApplicationFirewallModePtrOutput) ToWebApplicationFirewallModePtrOutput() WebApplicationFirewallModePtrOutput {
	return o
}

func (o WebApplicationFirewallModePtrOutput) ToWebApplicationFirewallModePtrOutputWithContext(ctx context.Context) WebApplicationFirewallModePtrOutput {
	return o
}

func (o WebApplicationFirewallModePtrOutput) Elem() WebApplicationFirewallModeOutput {
	return o.ApplyT(func(v *WebApplicationFirewallMode) WebApplicationFirewallMode {
		if v != nil {
			return *v
		}
		var ret WebApplicationFirewallMode
		return ret
	}).(WebApplicationFirewallModeOutput)
}

func (o WebApplicationFirewallModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebApplicationFirewallMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebApplicationFirewallModeInput interface {
	pulumi.Input

	ToWebApplicationFirewallModeOutput() WebApplicationFirewallModeOutput
	ToWebApplicationFirewallModeOutputWithContext(context.Context) WebApplicationFirewallModeOutput
}

var webApplicationFirewallModePtrType = reflect.TypeOf((**WebApplicationFirewallMode)(nil)).Elem()

type WebApplicationFirewallModePtrInput interface {
	pulumi.Input

	ToWebApplicationFirewallModePtrOutput() WebApplicationFirewallModePtrOutput
	ToWebApplicationFirewallModePtrOutputWithContext(context.Context) WebApplicationFirewallModePtrOutput
}

type webApplicationFirewallModePtr string

func WebApplicationFirewallModePtr(v string) WebApplicationFirewallModePtrInput {
	return (*webApplicationFirewallModePtr)(&v)
}

func (*webApplicationFirewallModePtr) ElementType() reflect.Type {
	return webApplicationFirewallModePtrType
}

func (in *webApplicationFirewallModePtr) ToWebApplicationFirewallModePtrOutput() WebApplicationFirewallModePtrOutput {
	return pulumi.ToOutput(in).(WebApplicationFirewallModePtrOutput)
}

func (in *webApplicationFirewallModePtr) ToWebApplicationFirewallModePtrOutputWithContext(ctx context.Context) WebApplicationFirewallModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebApplicationFirewallModePtrOutput)
}

type WebApplicationFirewallOperator string

const (
	WebApplicationFirewallOperatorIPMatch            = WebApplicationFirewallOperator("IPMatch")
	WebApplicationFirewallOperatorEqual              = WebApplicationFirewallOperator("Equal")
	WebApplicationFirewallOperatorContains           = WebApplicationFirewallOperator("Contains")
	WebApplicationFirewallOperatorLessThan           = WebApplicationFirewallOperator("LessThan")
	WebApplicationFirewallOperatorGreaterThan        = WebApplicationFirewallOperator("GreaterThan")
	WebApplicationFirewallOperatorLessThanOrEqual    = WebApplicationFirewallOperator("LessThanOrEqual")
	WebApplicationFirewallOperatorGreaterThanOrEqual = WebApplicationFirewallOperator("GreaterThanOrEqual")
	WebApplicationFirewallOperatorBeginsWith         = WebApplicationFirewallOperator("BeginsWith")
	WebApplicationFirewallOperatorEndsWith           = WebApplicationFirewallOperator("EndsWith")
	WebApplicationFirewallOperatorRegex              = WebApplicationFirewallOperator("Regex")
	WebApplicationFirewallOperatorGeoMatch           = WebApplicationFirewallOperator("GeoMatch")
)

func (WebApplicationFirewallOperator) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallOperator)(nil)).Elem()
}

func (e WebApplicationFirewallOperator) ToWebApplicationFirewallOperatorOutput() WebApplicationFirewallOperatorOutput {
	return pulumi.ToOutput(e).(WebApplicationFirewallOperatorOutput)
}

func (e WebApplicationFirewallOperator) ToWebApplicationFirewallOperatorOutputWithContext(ctx context.Context) WebApplicationFirewallOperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebApplicationFirewallOperatorOutput)
}

func (e WebApplicationFirewallOperator) ToWebApplicationFirewallOperatorPtrOutput() WebApplicationFirewallOperatorPtrOutput {
	return e.ToWebApplicationFirewallOperatorPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallOperator) ToWebApplicationFirewallOperatorPtrOutputWithContext(ctx context.Context) WebApplicationFirewallOperatorPtrOutput {
	return WebApplicationFirewallOperator(e).ToWebApplicationFirewallOperatorOutputWithContext(ctx).ToWebApplicationFirewallOperatorPtrOutputWithContext(ctx)
}

func (e WebApplicationFirewallOperator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallOperator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallOperator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallOperator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebApplicationFirewallOperatorOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallOperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallOperator)(nil)).Elem()
}

func (o WebApplicationFirewallOperatorOutput) ToWebApplicationFirewallOperatorOutput() WebApplicationFirewallOperatorOutput {
	return o
}

func (o WebApplicationFirewallOperatorOutput) ToWebApplicationFirewallOperatorOutputWithContext(ctx context.Context) WebApplicationFirewallOperatorOutput {
	return o
}

func (o WebApplicationFirewallOperatorOutput) ToWebApplicationFirewallOperatorPtrOutput() WebApplicationFirewallOperatorPtrOutput {
	return o.ToWebApplicationFirewallOperatorPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallOperatorOutput) ToWebApplicationFirewallOperatorPtrOutputWithContext(ctx context.Context) WebApplicationFirewallOperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebApplicationFirewallOperator) *WebApplicationFirewallOperator {
		return &v
	}).(WebApplicationFirewallOperatorPtrOutput)
}

func (o WebApplicationFirewallOperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebApplicationFirewallOperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallOperator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebApplicationFirewallOperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallOperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallOperator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebApplicationFirewallOperatorPtrOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallOperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallOperator)(nil)).Elem()
}

func (o WebApplicationFirewallOperatorPtrOutput) ToWebApplicationFirewallOperatorPtrOutput() WebApplicationFirewallOperatorPtrOutput {
	return o
}

func (o WebApplicationFirewallOperatorPtrOutput) ToWebApplicationFirewallOperatorPtrOutputWithContext(ctx context.Context) WebApplicationFirewallOperatorPtrOutput {
	return o
}

func (o WebApplicationFirewallOperatorPtrOutput) Elem() WebApplicationFirewallOperatorOutput {
	return o.ApplyT(func(v *WebApplicationFirewallOperator) WebApplicationFirewallOperator {
		if v != nil {
			return *v
		}
		var ret WebApplicationFirewallOperator
		return ret
	}).(WebApplicationFirewallOperatorOutput)
}

func (o WebApplicationFirewallOperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallOperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebApplicationFirewallOperator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebApplicationFirewallOperatorInput interface {
	pulumi.Input

	ToWebApplicationFirewallOperatorOutput() WebApplicationFirewallOperatorOutput
	ToWebApplicationFirewallOperatorOutputWithContext(context.Context) WebApplicationFirewallOperatorOutput
}

var webApplicationFirewallOperatorPtrType = reflect.TypeOf((**WebApplicationFirewallOperator)(nil)).Elem()

type WebApplicationFirewallOperatorPtrInput interface {
	pulumi.Input

	ToWebApplicationFirewallOperatorPtrOutput() WebApplicationFirewallOperatorPtrOutput
	ToWebApplicationFirewallOperatorPtrOutputWithContext(context.Context) WebApplicationFirewallOperatorPtrOutput
}

type webApplicationFirewallOperatorPtr string

func WebApplicationFirewallOperatorPtr(v string) WebApplicationFirewallOperatorPtrInput {
	return (*webApplicationFirewallOperatorPtr)(&v)
}

func (*webApplicationFirewallOperatorPtr) ElementType() reflect.Type {
	return webApplicationFirewallOperatorPtrType
}

func (in *webApplicationFirewallOperatorPtr) ToWebApplicationFirewallOperatorPtrOutput() WebApplicationFirewallOperatorPtrOutput {
	return pulumi.ToOutput(in).(WebApplicationFirewallOperatorPtrOutput)
}

func (in *webApplicationFirewallOperatorPtr) ToWebApplicationFirewallOperatorPtrOutputWithContext(ctx context.Context) WebApplicationFirewallOperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebApplicationFirewallOperatorPtrOutput)
}

type WebApplicationFirewallRuleType string

const (
	WebApplicationFirewallRuleTypeMatchRule = WebApplicationFirewallRuleType("MatchRule")
	WebApplicationFirewallRuleTypeInvalid   = WebApplicationFirewallRuleType("Invalid")
)

func (WebApplicationFirewallRuleType) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallRuleType)(nil)).Elem()
}

func (e WebApplicationFirewallRuleType) ToWebApplicationFirewallRuleTypeOutput() WebApplicationFirewallRuleTypeOutput {
	return pulumi.ToOutput(e).(WebApplicationFirewallRuleTypeOutput)
}

func (e WebApplicationFirewallRuleType) ToWebApplicationFirewallRuleTypeOutputWithContext(ctx context.Context) WebApplicationFirewallRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebApplicationFirewallRuleTypeOutput)
}

func (e WebApplicationFirewallRuleType) ToWebApplicationFirewallRuleTypePtrOutput() WebApplicationFirewallRuleTypePtrOutput {
	return e.ToWebApplicationFirewallRuleTypePtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallRuleType) ToWebApplicationFirewallRuleTypePtrOutputWithContext(ctx context.Context) WebApplicationFirewallRuleTypePtrOutput {
	return WebApplicationFirewallRuleType(e).ToWebApplicationFirewallRuleTypeOutputWithContext(ctx).ToWebApplicationFirewallRuleTypePtrOutputWithContext(ctx)
}

func (e WebApplicationFirewallRuleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallRuleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallRuleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallRuleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebApplicationFirewallRuleTypeOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallRuleType)(nil)).Elem()
}

func (o WebApplicationFirewallRuleTypeOutput) ToWebApplicationFirewallRuleTypeOutput() WebApplicationFirewallRuleTypeOutput {
	return o
}

func (o WebApplicationFirewallRuleTypeOutput) ToWebApplicationFirewallRuleTypeOutputWithContext(ctx context.Context) WebApplicationFirewallRuleTypeOutput {
	return o
}

func (o WebApplicationFirewallRuleTypeOutput) ToWebApplicationFirewallRuleTypePtrOutput() WebApplicationFirewallRuleTypePtrOutput {
	return o.ToWebApplicationFirewallRuleTypePtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallRuleTypeOutput) ToWebApplicationFirewallRuleTypePtrOutputWithContext(ctx context.Context) WebApplicationFirewallRuleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebApplicationFirewallRuleType) *WebApplicationFirewallRuleType {
		return &v
	}).(WebApplicationFirewallRuleTypePtrOutput)
}

func (o WebApplicationFirewallRuleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebApplicationFirewallRuleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallRuleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebApplicationFirewallRuleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallRuleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallRuleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebApplicationFirewallRuleTypePtrOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallRuleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallRuleType)(nil)).Elem()
}

func (o WebApplicationFirewallRuleTypePtrOutput) ToWebApplicationFirewallRuleTypePtrOutput() WebApplicationFirewallRuleTypePtrOutput {
	return o
}

func (o WebApplicationFirewallRuleTypePtrOutput) ToWebApplicationFirewallRuleTypePtrOutputWithContext(ctx context.Context) WebApplicationFirewallRuleTypePtrOutput {
	return o
}

func (o WebApplicationFirewallRuleTypePtrOutput) Elem() WebApplicationFirewallRuleTypeOutput {
	return o.ApplyT(func(v *WebApplicationFirewallRuleType) WebApplicationFirewallRuleType {
		if v != nil {
			return *v
		}
		var ret WebApplicationFirewallRuleType
		return ret
	}).(WebApplicationFirewallRuleTypeOutput)
}

func (o WebApplicationFirewallRuleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallRuleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebApplicationFirewallRuleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebApplicationFirewallRuleTypeInput interface {
	pulumi.Input

	ToWebApplicationFirewallRuleTypeOutput() WebApplicationFirewallRuleTypeOutput
	ToWebApplicationFirewallRuleTypeOutputWithContext(context.Context) WebApplicationFirewallRuleTypeOutput
}

var webApplicationFirewallRuleTypePtrType = reflect.TypeOf((**WebApplicationFirewallRuleType)(nil)).Elem()

type WebApplicationFirewallRuleTypePtrInput interface {
	pulumi.Input

	ToWebApplicationFirewallRuleTypePtrOutput() WebApplicationFirewallRuleTypePtrOutput
	ToWebApplicationFirewallRuleTypePtrOutputWithContext(context.Context) WebApplicationFirewallRuleTypePtrOutput
}

type webApplicationFirewallRuleTypePtr string

func WebApplicationFirewallRuleTypePtr(v string) WebApplicationFirewallRuleTypePtrInput {
	return (*webApplicationFirewallRuleTypePtr)(&v)
}

func (*webApplicationFirewallRuleTypePtr) ElementType() reflect.Type {
	return webApplicationFirewallRuleTypePtrType
}

func (in *webApplicationFirewallRuleTypePtr) ToWebApplicationFirewallRuleTypePtrOutput() WebApplicationFirewallRuleTypePtrOutput {
	return pulumi.ToOutput(in).(WebApplicationFirewallRuleTypePtrOutput)
}

func (in *webApplicationFirewallRuleTypePtr) ToWebApplicationFirewallRuleTypePtrOutputWithContext(ctx context.Context) WebApplicationFirewallRuleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebApplicationFirewallRuleTypePtrOutput)
}

type WebApplicationFirewallTransform string

const (
	WebApplicationFirewallTransformLowercase        = WebApplicationFirewallTransform("Lowercase")
	WebApplicationFirewallTransformTrim             = WebApplicationFirewallTransform("Trim")
	WebApplicationFirewallTransformUrlDecode        = WebApplicationFirewallTransform("UrlDecode")
	WebApplicationFirewallTransformUrlEncode        = WebApplicationFirewallTransform("UrlEncode")
	WebApplicationFirewallTransformRemoveNulls      = WebApplicationFirewallTransform("RemoveNulls")
	WebApplicationFirewallTransformHtmlEntityDecode = WebApplicationFirewallTransform("HtmlEntityDecode")
)

func (WebApplicationFirewallTransform) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallTransform)(nil)).Elem()
}

func (e WebApplicationFirewallTransform) ToWebApplicationFirewallTransformOutput() WebApplicationFirewallTransformOutput {
	return pulumi.ToOutput(e).(WebApplicationFirewallTransformOutput)
}

func (e WebApplicationFirewallTransform) ToWebApplicationFirewallTransformOutputWithContext(ctx context.Context) WebApplicationFirewallTransformOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebApplicationFirewallTransformOutput)
}

func (e WebApplicationFirewallTransform) ToWebApplicationFirewallTransformPtrOutput() WebApplicationFirewallTransformPtrOutput {
	return e.ToWebApplicationFirewallTransformPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallTransform) ToWebApplicationFirewallTransformPtrOutputWithContext(ctx context.Context) WebApplicationFirewallTransformPtrOutput {
	return WebApplicationFirewallTransform(e).ToWebApplicationFirewallTransformOutputWithContext(ctx).ToWebApplicationFirewallTransformPtrOutputWithContext(ctx)
}

func (e WebApplicationFirewallTransform) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallTransform) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebApplicationFirewallTransform) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebApplicationFirewallTransform) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebApplicationFirewallTransformOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallTransformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallTransform)(nil)).Elem()
}

func (o WebApplicationFirewallTransformOutput) ToWebApplicationFirewallTransformOutput() WebApplicationFirewallTransformOutput {
	return o
}

func (o WebApplicationFirewallTransformOutput) ToWebApplicationFirewallTransformOutputWithContext(ctx context.Context) WebApplicationFirewallTransformOutput {
	return o
}

func (o WebApplicationFirewallTransformOutput) ToWebApplicationFirewallTransformPtrOutput() WebApplicationFirewallTransformPtrOutput {
	return o.ToWebApplicationFirewallTransformPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallTransformOutput) ToWebApplicationFirewallTransformPtrOutputWithContext(ctx context.Context) WebApplicationFirewallTransformPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebApplicationFirewallTransform) *WebApplicationFirewallTransform {
		return &v
	}).(WebApplicationFirewallTransformPtrOutput)
}

func (o WebApplicationFirewallTransformOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebApplicationFirewallTransformOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallTransform) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebApplicationFirewallTransformOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallTransformOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebApplicationFirewallTransform) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebApplicationFirewallTransformPtrOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallTransformPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallTransform)(nil)).Elem()
}

func (o WebApplicationFirewallTransformPtrOutput) ToWebApplicationFirewallTransformPtrOutput() WebApplicationFirewallTransformPtrOutput {
	return o
}

func (o WebApplicationFirewallTransformPtrOutput) ToWebApplicationFirewallTransformPtrOutputWithContext(ctx context.Context) WebApplicationFirewallTransformPtrOutput {
	return o
}

func (o WebApplicationFirewallTransformPtrOutput) Elem() WebApplicationFirewallTransformOutput {
	return o.ApplyT(func(v *WebApplicationFirewallTransform) WebApplicationFirewallTransform {
		if v != nil {
			return *v
		}
		var ret WebApplicationFirewallTransform
		return ret
	}).(WebApplicationFirewallTransformOutput)
}

func (o WebApplicationFirewallTransformPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebApplicationFirewallTransformPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebApplicationFirewallTransform) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebApplicationFirewallTransformInput interface {
	pulumi.Input

	ToWebApplicationFirewallTransformOutput() WebApplicationFirewallTransformOutput
	ToWebApplicationFirewallTransformOutputWithContext(context.Context) WebApplicationFirewallTransformOutput
}

var webApplicationFirewallTransformPtrType = reflect.TypeOf((**WebApplicationFirewallTransform)(nil)).Elem()

type WebApplicationFirewallTransformPtrInput interface {
	pulumi.Input

	ToWebApplicationFirewallTransformPtrOutput() WebApplicationFirewallTransformPtrOutput
	ToWebApplicationFirewallTransformPtrOutputWithContext(context.Context) WebApplicationFirewallTransformPtrOutput
}

type webApplicationFirewallTransformPtr string

func WebApplicationFirewallTransformPtr(v string) WebApplicationFirewallTransformPtrInput {
	return (*webApplicationFirewallTransformPtr)(&v)
}

func (*webApplicationFirewallTransformPtr) ElementType() reflect.Type {
	return webApplicationFirewallTransformPtrType
}

func (in *webApplicationFirewallTransformPtr) ToWebApplicationFirewallTransformPtrOutput() WebApplicationFirewallTransformPtrOutput {
	return pulumi.ToOutput(in).(WebApplicationFirewallTransformPtrOutput)
}

func (in *webApplicationFirewallTransformPtr) ToWebApplicationFirewallTransformPtrOutputWithContext(ctx context.Context) WebApplicationFirewallTransformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebApplicationFirewallTransformPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessOutput{})
	pulumi.RegisterOutputType(AccessPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayCookieBasedAffinityOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayCookieBasedAffinityPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayCustomErrorStatusCodeOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayCustomErrorStatusCodePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallModeOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallModePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayLoadDistributionAlgorithmOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayLoadDistributionAlgorithmPtrOutput{})
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
	pulumi.RegisterOutputType(AzureFirewallApplicationRuleProtocolTypeOutput{})
	pulumi.RegisterOutputType(AzureFirewallApplicationRuleProtocolTypePtrOutput{})
	pulumi.RegisterOutputType(AzureFirewallNatRCActionTypeOutput{})
	pulumi.RegisterOutputType(AzureFirewallNatRCActionTypePtrOutput{})
	pulumi.RegisterOutputType(AzureFirewallNetworkRuleProtocolOutput{})
	pulumi.RegisterOutputType(AzureFirewallNetworkRuleProtocolPtrOutput{})
	pulumi.RegisterOutputType(AzureFirewallRCActionTypeOutput{})
	pulumi.RegisterOutputType(AzureFirewallRCActionTypePtrOutput{})
	pulumi.RegisterOutputType(AzureFirewallSkuNameOutput{})
	pulumi.RegisterOutputType(AzureFirewallSkuNamePtrOutput{})
	pulumi.RegisterOutputType(AzureFirewallSkuTierOutput{})
	pulumi.RegisterOutputType(AzureFirewallSkuTierPtrOutput{})
	pulumi.RegisterOutputType(AzureFirewallThreatIntelModeOutput{})
	pulumi.RegisterOutputType(AzureFirewallThreatIntelModePtrOutput{})
	pulumi.RegisterOutputType(BastionHostSkuNameOutput{})
	pulumi.RegisterOutputType(BastionHostSkuNamePtrOutput{})
	pulumi.RegisterOutputType(CommissionedStateOutput{})
	pulumi.RegisterOutputType(CommissionedStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorEndpointFilterItemTypeOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorEndpointFilterItemTypePtrOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorEndpointFilterTypeOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorEndpointFilterTypePtrOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorTestConfigurationProtocolOutput{})
	pulumi.RegisterOutputType(ConnectionMonitorTestConfigurationProtocolPtrOutput{})
	pulumi.RegisterOutputType(CoverageLevelOutput{})
	pulumi.RegisterOutputType(CoverageLevelPtrOutput{})
	pulumi.RegisterOutputType(DdosCustomPolicyProtocolOutput{})
	pulumi.RegisterOutputType(DdosCustomPolicyProtocolPtrOutput{})
	pulumi.RegisterOutputType(DdosCustomPolicyTriggerSensitivityOverrideOutput{})
	pulumi.RegisterOutputType(DdosCustomPolicyTriggerSensitivityOverridePtrOutput{})
	pulumi.RegisterOutputType(DdosSettingsProtectionCoverageOutput{})
	pulumi.RegisterOutputType(DdosSettingsProtectionCoveragePtrOutput{})
	pulumi.RegisterOutputType(DeleteOptionsOutput{})
	pulumi.RegisterOutputType(DeleteOptionsPtrOutput{})
	pulumi.RegisterOutputType(DestinationPortBehaviorOutput{})
	pulumi.RegisterOutputType(DestinationPortBehaviorPtrOutput{})
	pulumi.RegisterOutputType(DhGroupOutput{})
	pulumi.RegisterOutputType(DhGroupPtrOutput{})
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringStateEnumOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringStateEnumPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuFamilyOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuFamilyPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuTierOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuTierPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteLinkAdminStateOutput{})
	pulumi.RegisterOutputType(ExpressRouteLinkAdminStatePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteLinkMacSecCipherOutput{})
	pulumi.RegisterOutputType(ExpressRouteLinkMacSecCipherPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteLinkMacSecSciStateOutput{})
	pulumi.RegisterOutputType(ExpressRouteLinkMacSecSciStatePtrOutput{})
	pulumi.RegisterOutputType(ExpressRoutePeeringStateOutput{})
	pulumi.RegisterOutputType(ExpressRoutePeeringStatePtrOutput{})
	pulumi.RegisterOutputType(ExpressRoutePeeringTypeOutput{})
	pulumi.RegisterOutputType(ExpressRoutePeeringTypePtrOutput{})
	pulumi.RegisterOutputType(ExpressRoutePortsEncapsulationOutput{})
	pulumi.RegisterOutputType(ExpressRoutePortsEncapsulationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesOutput{})
	pulumi.RegisterOutputType(ExtendedLocationTypesPtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyFilterRuleCollectionActionTypeOutput{})
	pulumi.RegisterOutputType(FirewallPolicyFilterRuleCollectionActionTypePtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyIntrusionDetectionProtocolOutput{})
	pulumi.RegisterOutputType(FirewallPolicyIntrusionDetectionProtocolPtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyIntrusionDetectionStateTypeOutput{})
	pulumi.RegisterOutputType(FirewallPolicyIntrusionDetectionStateTypePtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyNatRuleCollectionActionTypeOutput{})
	pulumi.RegisterOutputType(FirewallPolicyNatRuleCollectionActionTypePtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleApplicationProtocolTypeOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleApplicationProtocolTypePtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionTypeOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleCollectionTypePtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleNetworkProtocolOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleNetworkProtocolPtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleTypeOutput{})
	pulumi.RegisterOutputType(FirewallPolicyRuleTypePtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicySkuTierOutput{})
	pulumi.RegisterOutputType(FirewallPolicySkuTierPtrOutput{})
	pulumi.RegisterOutputType(FlowLogFormatTypeOutput{})
	pulumi.RegisterOutputType(FlowLogFormatTypePtrOutput{})
	pulumi.RegisterOutputType(GatewayLoadBalancerTunnelInterfaceTypeOutput{})
	pulumi.RegisterOutputType(GatewayLoadBalancerTunnelInterfaceTypePtrOutput{})
	pulumi.RegisterOutputType(GatewayLoadBalancerTunnelProtocolOutput{})
	pulumi.RegisterOutputType(GatewayLoadBalancerTunnelProtocolPtrOutput{})
	pulumi.RegisterOutputType(HTTPConfigurationMethodOutput{})
	pulumi.RegisterOutputType(HTTPConfigurationMethodPtrOutput{})
	pulumi.RegisterOutputType(IPAllocationMethodOutput{})
	pulumi.RegisterOutputType(IPAllocationMethodPtrOutput{})
	pulumi.RegisterOutputType(IPVersionOutput{})
	pulumi.RegisterOutputType(IPVersionPtrOutput{})
	pulumi.RegisterOutputType(IkeEncryptionOutput{})
	pulumi.RegisterOutputType(IkeEncryptionPtrOutput{})
	pulumi.RegisterOutputType(IkeIntegrityOutput{})
	pulumi.RegisterOutputType(IkeIntegrityPtrOutput{})
	pulumi.RegisterOutputType(IpAllocationTypeOutput{})
	pulumi.RegisterOutputType(IpAllocationTypePtrOutput{})
	pulumi.RegisterOutputType(IpsecEncryptionOutput{})
	pulumi.RegisterOutputType(IpsecEncryptionPtrOutput{})
	pulumi.RegisterOutputType(IpsecIntegrityOutput{})
	pulumi.RegisterOutputType(IpsecIntegrityPtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerOutboundRuleProtocolOutput{})
	pulumi.RegisterOutputType(LoadBalancerOutboundRuleProtocolPtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuNameOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuNamePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuTierOutput{})
	pulumi.RegisterOutputType(LoadBalancerSkuTierPtrOutput{})
	pulumi.RegisterOutputType(LoadDistributionOutput{})
	pulumi.RegisterOutputType(LoadDistributionPtrOutput{})
	pulumi.RegisterOutputType(ManagedRuleEnabledStateOutput{})
	pulumi.RegisterOutputType(ManagedRuleEnabledStatePtrOutput{})
	pulumi.RegisterOutputType(NatGatewaySkuNameOutput{})
	pulumi.RegisterOutputType(NatGatewaySkuNamePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceMigrationPhaseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceMigrationPhasePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceNicTypeOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceNicTypePtrOutput{})
	pulumi.RegisterOutputType(OutputTypeOutput{})
	pulumi.RegisterOutputType(OutputTypePtrOutput{})
	pulumi.RegisterOutputType(OwaspCrsExclusionEntryMatchVariableOutput{})
	pulumi.RegisterOutputType(OwaspCrsExclusionEntryMatchVariablePtrOutput{})
	pulumi.RegisterOutputType(OwaspCrsExclusionEntrySelectorMatchOperatorOutput{})
	pulumi.RegisterOutputType(OwaspCrsExclusionEntrySelectorMatchOperatorPtrOutput{})
	pulumi.RegisterOutputType(PcProtocolOutput{})
	pulumi.RegisterOutputType(PcProtocolPtrOutput{})
	pulumi.RegisterOutputType(PfsGroupOutput{})
	pulumi.RegisterOutputType(PfsGroupPtrOutput{})
	pulumi.RegisterOutputType(PreferredIPVersionOutput{})
	pulumi.RegisterOutputType(PreferredIPVersionPtrOutput{})
	pulumi.RegisterOutputType(PreferredRoutingGatewayOutput{})
	pulumi.RegisterOutputType(PreferredRoutingGatewayPtrOutput{})
	pulumi.RegisterOutputType(ProbeProtocolOutput{})
	pulumi.RegisterOutputType(ProbeProtocolPtrOutput{})
	pulumi.RegisterOutputType(ProtocolTypeOutput{})
	pulumi.RegisterOutputType(ProtocolTypePtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressMigrationPhaseOutput{})
	pulumi.RegisterOutputType(PublicIPAddressMigrationPhasePtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuNameOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuNamePtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuTierOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuTierPtrOutput{})
	pulumi.RegisterOutputType(PublicIPPrefixSkuNameOutput{})
	pulumi.RegisterOutputType(PublicIPPrefixSkuNamePtrOutput{})
	pulumi.RegisterOutputType(PublicIPPrefixSkuTierOutput{})
	pulumi.RegisterOutputType(PublicIPPrefixSkuTierPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleTypeEnumOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(RouteNextHopTypeOutput{})
	pulumi.RegisterOutputType(RouteNextHopTypePtrOutput{})
	pulumi.RegisterOutputType(SecurityProviderNameOutput{})
	pulumi.RegisterOutputType(SecurityProviderNamePtrOutput{})
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
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionModeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionModePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionProtocolOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionProtocolPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionTypePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuNameOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuNamePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuTierOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuTierPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypeEnumOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringLevelOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringLevelPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringStateEnumOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringStateEnumPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPrivateEndpointNetworkPoliciesOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPrivateEndpointNetworkPoliciesPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPrivateLinkServiceNetworkPoliciesOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPrivateLinkServiceNetworkPoliciesPtrOutput{})
	pulumi.RegisterOutputType(VpnAuthenticationTypeOutput{})
	pulumi.RegisterOutputType(VpnAuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(VpnClientProtocolOutput{})
	pulumi.RegisterOutputType(VpnClientProtocolPtrOutput{})
	pulumi.RegisterOutputType(VpnGatewayGenerationOutput{})
	pulumi.RegisterOutputType(VpnGatewayGenerationPtrOutput{})
	pulumi.RegisterOutputType(VpnGatewayTunnelingProtocolOutput{})
	pulumi.RegisterOutputType(VpnGatewayTunnelingProtocolPtrOutput{})
	pulumi.RegisterOutputType(VpnLinkConnectionModeOutput{})
	pulumi.RegisterOutputType(VpnLinkConnectionModePtrOutput{})
	pulumi.RegisterOutputType(VpnNatRuleModeOutput{})
	pulumi.RegisterOutputType(VpnNatRuleModePtrOutput{})
	pulumi.RegisterOutputType(VpnNatRuleTypeOutput{})
	pulumi.RegisterOutputType(VpnNatRuleTypePtrOutput{})
	pulumi.RegisterOutputType(VpnTypeOutput{})
	pulumi.RegisterOutputType(VpnTypePtrOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallActionOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallActionPtrOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallEnabledStateOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallEnabledStatePtrOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallMatchVariableOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallMatchVariablePtrOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallModeOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallModePtrOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallOperatorOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallOperatorPtrOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallRuleTypeOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallRuleTypePtrOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallTransformOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallTransformPtrOutput{})
}
