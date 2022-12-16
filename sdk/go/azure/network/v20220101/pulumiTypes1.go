


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnLinkBgpSettingsResponse struct {
	Asn               *float64 `pulumi:"asn"`
	BgpPeeringAddress *string  `pulumi:"bgpPeeringAddress"`
}

type VpnLinkBgpSettingsResponseOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkBgpSettingsResponse)(nil)).Elem()
}

func (o VpnLinkBgpSettingsResponseOutput) ToVpnLinkBgpSettingsResponseOutput() VpnLinkBgpSettingsResponseOutput {
	return o
}

func (o VpnLinkBgpSettingsResponseOutput) ToVpnLinkBgpSettingsResponseOutputWithContext(ctx context.Context) VpnLinkBgpSettingsResponseOutput {
	return o
}

func (o VpnLinkBgpSettingsResponseOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettingsResponse) *float64 { return v.Asn }).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsResponseOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettingsResponse) *string { return v.BgpPeeringAddress }).(pulumi.StringPtrOutput)
}

type VpnLinkBgpSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkBgpSettingsResponse)(nil)).Elem()
}

func (o VpnLinkBgpSettingsResponsePtrOutput) ToVpnLinkBgpSettingsResponsePtrOutput() VpnLinkBgpSettingsResponsePtrOutput {
	return o
}

func (o VpnLinkBgpSettingsResponsePtrOutput) ToVpnLinkBgpSettingsResponsePtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsResponsePtrOutput {
	return o
}

func (o VpnLinkBgpSettingsResponsePtrOutput) Elem() VpnLinkBgpSettingsResponseOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettingsResponse) VpnLinkBgpSettingsResponse {
		if v != nil {
			return *v
		}
		var ret VpnLinkBgpSettingsResponse
		return ret
	}).(VpnLinkBgpSettingsResponseOutput)
}

func (o VpnLinkBgpSettingsResponsePtrOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Asn
	}).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsResponsePtrOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.BgpPeeringAddress
	}).(pulumi.StringPtrOutput)
}

type VpnLinkProviderProperties struct {
	LinkProviderName *string `pulumi:"linkProviderName"`
	LinkSpeedInMbps  *int    `pulumi:"linkSpeedInMbps"`
}





type VpnLinkProviderPropertiesInput interface {
	pulumi.Input

	ToVpnLinkProviderPropertiesOutput() VpnLinkProviderPropertiesOutput
	ToVpnLinkProviderPropertiesOutputWithContext(context.Context) VpnLinkProviderPropertiesOutput
}

type VpnLinkProviderPropertiesArgs struct {
	LinkProviderName pulumi.StringPtrInput `pulumi:"linkProviderName"`
	LinkSpeedInMbps  pulumi.IntPtrInput    `pulumi:"linkSpeedInMbps"`
}

func (VpnLinkProviderPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkProviderProperties)(nil)).Elem()
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesOutput() VpnLinkProviderPropertiesOutput {
	return i.ToVpnLinkProviderPropertiesOutputWithContext(context.Background())
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkProviderPropertiesOutput)
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return i.ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkProviderPropertiesOutput).ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx)
}









type VpnLinkProviderPropertiesPtrInput interface {
	pulumi.Input

	ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput
	ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Context) VpnLinkProviderPropertiesPtrOutput
}

type vpnLinkProviderPropertiesPtrType VpnLinkProviderPropertiesArgs

func VpnLinkProviderPropertiesPtr(v *VpnLinkProviderPropertiesArgs) VpnLinkProviderPropertiesPtrInput {
	return (*vpnLinkProviderPropertiesPtrType)(v)
}

func (*vpnLinkProviderPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkProviderProperties)(nil)).Elem()
}

func (i *vpnLinkProviderPropertiesPtrType) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return i.ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i *vpnLinkProviderPropertiesPtrType) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkProviderPropertiesPtrOutput)
}

type VpnLinkProviderPropertiesOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkProviderProperties)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesOutput() VpnLinkProviderPropertiesOutput {
	return o
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesOutput {
	return o
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return o.ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Background())
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnLinkProviderProperties) *VpnLinkProviderProperties {
		return &v
	}).(VpnLinkProviderPropertiesPtrOutput)
}

func (o VpnLinkProviderPropertiesOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderProperties) *string { return v.LinkProviderName }).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderProperties) *int { return v.LinkSpeedInMbps }).(pulumi.IntPtrOutput)
}

type VpnLinkProviderPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkProviderProperties)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesPtrOutput) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesPtrOutput) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesPtrOutput) Elem() VpnLinkProviderPropertiesOutput {
	return o.ApplyT(func(v *VpnLinkProviderProperties) VpnLinkProviderProperties {
		if v != nil {
			return *v
		}
		var ret VpnLinkProviderProperties
		return ret
	}).(VpnLinkProviderPropertiesOutput)
}

func (o VpnLinkProviderPropertiesPtrOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderProperties) *string {
		if v == nil {
			return nil
		}
		return v.LinkProviderName
	}).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesPtrOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderProperties) *int {
		if v == nil {
			return nil
		}
		return v.LinkSpeedInMbps
	}).(pulumi.IntPtrOutput)
}

type VpnLinkProviderPropertiesResponse struct {
	LinkProviderName *string `pulumi:"linkProviderName"`
	LinkSpeedInMbps  *int    `pulumi:"linkSpeedInMbps"`
}

type VpnLinkProviderPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkProviderPropertiesResponse)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesResponseOutput) ToVpnLinkProviderPropertiesResponseOutput() VpnLinkProviderPropertiesResponseOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponseOutput) ToVpnLinkProviderPropertiesResponseOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesResponseOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponseOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderPropertiesResponse) *string { return v.LinkProviderName }).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesResponseOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderPropertiesResponse) *int { return v.LinkSpeedInMbps }).(pulumi.IntPtrOutput)
}

type VpnLinkProviderPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkProviderPropertiesResponse)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) ToVpnLinkProviderPropertiesResponsePtrOutput() VpnLinkProviderPropertiesResponsePtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) ToVpnLinkProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesResponsePtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) Elem() VpnLinkProviderPropertiesResponseOutput {
	return o.ApplyT(func(v *VpnLinkProviderPropertiesResponse) VpnLinkProviderPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VpnLinkProviderPropertiesResponse
		return ret
	}).(VpnLinkProviderPropertiesResponseOutput)
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinkProviderName
	}).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.LinkSpeedInMbps
	}).(pulumi.IntPtrOutput)
}

type VpnNatRuleMapping struct {
	AddressSpace *string `pulumi:"addressSpace"`
	PortRange    *string `pulumi:"portRange"`
}





type VpnNatRuleMappingInput interface {
	pulumi.Input

	ToVpnNatRuleMappingOutput() VpnNatRuleMappingOutput
	ToVpnNatRuleMappingOutputWithContext(context.Context) VpnNatRuleMappingOutput
}

type VpnNatRuleMappingArgs struct {
	AddressSpace pulumi.StringPtrInput `pulumi:"addressSpace"`
	PortRange    pulumi.StringPtrInput `pulumi:"portRange"`
}

func (VpnNatRuleMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMapping)(nil)).Elem()
}

func (i VpnNatRuleMappingArgs) ToVpnNatRuleMappingOutput() VpnNatRuleMappingOutput {
	return i.ToVpnNatRuleMappingOutputWithContext(context.Background())
}

func (i VpnNatRuleMappingArgs) ToVpnNatRuleMappingOutputWithContext(ctx context.Context) VpnNatRuleMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnNatRuleMappingOutput)
}





type VpnNatRuleMappingArrayInput interface {
	pulumi.Input

	ToVpnNatRuleMappingArrayOutput() VpnNatRuleMappingArrayOutput
	ToVpnNatRuleMappingArrayOutputWithContext(context.Context) VpnNatRuleMappingArrayOutput
}

type VpnNatRuleMappingArray []VpnNatRuleMappingInput

func (VpnNatRuleMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnNatRuleMapping)(nil)).Elem()
}

func (i VpnNatRuleMappingArray) ToVpnNatRuleMappingArrayOutput() VpnNatRuleMappingArrayOutput {
	return i.ToVpnNatRuleMappingArrayOutputWithContext(context.Background())
}

func (i VpnNatRuleMappingArray) ToVpnNatRuleMappingArrayOutputWithContext(ctx context.Context) VpnNatRuleMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnNatRuleMappingArrayOutput)
}

type VpnNatRuleMappingOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMapping)(nil)).Elem()
}

func (o VpnNatRuleMappingOutput) ToVpnNatRuleMappingOutput() VpnNatRuleMappingOutput {
	return o
}

func (o VpnNatRuleMappingOutput) ToVpnNatRuleMappingOutputWithContext(ctx context.Context) VpnNatRuleMappingOutput {
	return o
}

func (o VpnNatRuleMappingOutput) AddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMapping) *string { return v.AddressSpace }).(pulumi.StringPtrOutput)
}

func (o VpnNatRuleMappingOutput) PortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMapping) *string { return v.PortRange }).(pulumi.StringPtrOutput)
}

type VpnNatRuleMappingArrayOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnNatRuleMapping)(nil)).Elem()
}

func (o VpnNatRuleMappingArrayOutput) ToVpnNatRuleMappingArrayOutput() VpnNatRuleMappingArrayOutput {
	return o
}

func (o VpnNatRuleMappingArrayOutput) ToVpnNatRuleMappingArrayOutputWithContext(ctx context.Context) VpnNatRuleMappingArrayOutput {
	return o
}

func (o VpnNatRuleMappingArrayOutput) Index(i pulumi.IntInput) VpnNatRuleMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnNatRuleMapping {
		return vs[0].([]VpnNatRuleMapping)[vs[1].(int)]
	}).(VpnNatRuleMappingOutput)
}

type VpnNatRuleMappingResponse struct {
	AddressSpace *string `pulumi:"addressSpace"`
	PortRange    *string `pulumi:"portRange"`
}

type VpnNatRuleMappingResponseOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMappingResponse)(nil)).Elem()
}

func (o VpnNatRuleMappingResponseOutput) ToVpnNatRuleMappingResponseOutput() VpnNatRuleMappingResponseOutput {
	return o
}

func (o VpnNatRuleMappingResponseOutput) ToVpnNatRuleMappingResponseOutputWithContext(ctx context.Context) VpnNatRuleMappingResponseOutput {
	return o
}

func (o VpnNatRuleMappingResponseOutput) AddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMappingResponse) *string { return v.AddressSpace }).(pulumi.StringPtrOutput)
}

func (o VpnNatRuleMappingResponseOutput) PortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMappingResponse) *string { return v.PortRange }).(pulumi.StringPtrOutput)
}

type VpnNatRuleMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnNatRuleMappingResponse)(nil)).Elem()
}

func (o VpnNatRuleMappingResponseArrayOutput) ToVpnNatRuleMappingResponseArrayOutput() VpnNatRuleMappingResponseArrayOutput {
	return o
}

func (o VpnNatRuleMappingResponseArrayOutput) ToVpnNatRuleMappingResponseArrayOutputWithContext(ctx context.Context) VpnNatRuleMappingResponseArrayOutput {
	return o
}

func (o VpnNatRuleMappingResponseArrayOutput) Index(i pulumi.IntInput) VpnNatRuleMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnNatRuleMappingResponse {
		return vs[0].([]VpnNatRuleMappingResponse)[vs[1].(int)]
	}).(VpnNatRuleMappingResponseOutput)
}

type VpnServerConfigRadiusClientRootCertificate struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnServerConfigRadiusClientRootCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusClientRootCertificateOutput() VpnServerConfigRadiusClientRootCertificateOutput
	ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(context.Context) VpnServerConfigRadiusClientRootCertificateOutput
}

type VpnServerConfigRadiusClientRootCertificateArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnServerConfigRadiusClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusClientRootCertificateArgs) ToVpnServerConfigRadiusClientRootCertificateOutput() VpnServerConfigRadiusClientRootCertificateOutput {
	return i.ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusClientRootCertificateArgs) ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusClientRootCertificateOutput)
}





type VpnServerConfigRadiusClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusClientRootCertificateArrayOutput() VpnServerConfigRadiusClientRootCertificateArrayOutput
	ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(context.Context) VpnServerConfigRadiusClientRootCertificateArrayOutput
}

type VpnServerConfigRadiusClientRootCertificateArray []VpnServerConfigRadiusClientRootCertificateInput

func (VpnServerConfigRadiusClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusClientRootCertificateArray) ToVpnServerConfigRadiusClientRootCertificateArrayOutput() VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return i.ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusClientRootCertificateArray) ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusClientRootCertificateArrayOutput)
}

type VpnServerConfigRadiusClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) ToVpnServerConfigRadiusClientRootCertificateOutput() VpnServerConfigRadiusClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateArrayOutput) ToVpnServerConfigRadiusClientRootCertificateArrayOutput() VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateArrayOutput) ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusClientRootCertificate {
		return vs[0].([]VpnServerConfigRadiusClientRootCertificate)[vs[1].(int)]
	}).(VpnServerConfigRadiusClientRootCertificateOutput)
}

type VpnServerConfigRadiusClientRootCertificateResponse struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}

type VpnServerConfigRadiusClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) ToVpnServerConfigRadiusClientRootCertificateResponseOutput() VpnServerConfigRadiusClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) ToVpnServerConfigRadiusClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusClientRootCertificateResponseArrayOutput() VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusClientRootCertificateResponse {
		return vs[0].([]VpnServerConfigRadiusClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigRadiusClientRootCertificateResponseOutput)
}

type VpnServerConfigRadiusServerRootCertificate struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}





type VpnServerConfigRadiusServerRootCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusServerRootCertificateOutput() VpnServerConfigRadiusServerRootCertificateOutput
	ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(context.Context) VpnServerConfigRadiusServerRootCertificateOutput
}

type VpnServerConfigRadiusServerRootCertificateArgs struct {
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringPtrInput `pulumi:"publicCertData"`
}

func (VpnServerConfigRadiusServerRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusServerRootCertificateArgs) ToVpnServerConfigRadiusServerRootCertificateOutput() VpnServerConfigRadiusServerRootCertificateOutput {
	return i.ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusServerRootCertificateArgs) ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusServerRootCertificateOutput)
}





type VpnServerConfigRadiusServerRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusServerRootCertificateArrayOutput() VpnServerConfigRadiusServerRootCertificateArrayOutput
	ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(context.Context) VpnServerConfigRadiusServerRootCertificateArrayOutput
}

type VpnServerConfigRadiusServerRootCertificateArray []VpnServerConfigRadiusServerRootCertificateInput

func (VpnServerConfigRadiusServerRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusServerRootCertificateArray) ToVpnServerConfigRadiusServerRootCertificateArrayOutput() VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return i.ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusServerRootCertificateArray) ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusServerRootCertificateArrayOutput)
}

type VpnServerConfigRadiusServerRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) ToVpnServerConfigRadiusServerRootCertificateOutput() VpnServerConfigRadiusServerRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificate) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusServerRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateArrayOutput) ToVpnServerConfigRadiusServerRootCertificateArrayOutput() VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateArrayOutput) ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusServerRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusServerRootCertificate {
		return vs[0].([]VpnServerConfigRadiusServerRootCertificate)[vs[1].(int)]
	}).(VpnServerConfigRadiusServerRootCertificateOutput)
}

type VpnServerConfigRadiusServerRootCertificateResponse struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}

type VpnServerConfigRadiusServerRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusServerRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) ToVpnServerConfigRadiusServerRootCertificateResponseOutput() VpnServerConfigRadiusServerRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) ToVpnServerConfigRadiusServerRootCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificateResponse) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusServerRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusServerRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusServerRootCertificateResponseArrayOutput() VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusServerRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusServerRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusServerRootCertificateResponse {
		return vs[0].([]VpnServerConfigRadiusServerRootCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigRadiusServerRootCertificateResponseOutput)
}

type VpnServerConfigVpnClientRevokedCertificate struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnServerConfigVpnClientRevokedCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRevokedCertificateOutput() VpnServerConfigVpnClientRevokedCertificateOutput
	ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(context.Context) VpnServerConfigVpnClientRevokedCertificateOutput
}

type VpnServerConfigVpnClientRevokedCertificateArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnServerConfigVpnClientRevokedCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRevokedCertificateArgs) ToVpnServerConfigVpnClientRevokedCertificateOutput() VpnServerConfigVpnClientRevokedCertificateOutput {
	return i.ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRevokedCertificateArgs) ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRevokedCertificateOutput)
}





type VpnServerConfigVpnClientRevokedCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRevokedCertificateArrayOutput() VpnServerConfigVpnClientRevokedCertificateArrayOutput
	ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(context.Context) VpnServerConfigVpnClientRevokedCertificateArrayOutput
}

type VpnServerConfigVpnClientRevokedCertificateArray []VpnServerConfigVpnClientRevokedCertificateInput

func (VpnServerConfigVpnClientRevokedCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRevokedCertificateArray) ToVpnServerConfigVpnClientRevokedCertificateArrayOutput() VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return i.ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRevokedCertificateArray) ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRevokedCertificateArrayOutput)
}

type VpnServerConfigVpnClientRevokedCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) ToVpnServerConfigVpnClientRevokedCertificateOutput() VpnServerConfigVpnClientRevokedCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRevokedCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateArrayOutput() VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRevokedCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRevokedCertificate {
		return vs[0].([]VpnServerConfigVpnClientRevokedCertificate)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRevokedCertificateOutput)
}

type VpnServerConfigVpnClientRevokedCertificateResponse struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}

type VpnServerConfigVpnClientRevokedCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseOutput() VpnServerConfigVpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput() VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRevokedCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRevokedCertificateResponse {
		return vs[0].([]VpnServerConfigVpnClientRevokedCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRevokedCertificateResponseOutput)
}

type VpnServerConfigVpnClientRootCertificate struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}





type VpnServerConfigVpnClientRootCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRootCertificateOutput() VpnServerConfigVpnClientRootCertificateOutput
	ToVpnServerConfigVpnClientRootCertificateOutputWithContext(context.Context) VpnServerConfigVpnClientRootCertificateOutput
}

type VpnServerConfigVpnClientRootCertificateArgs struct {
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringPtrInput `pulumi:"publicCertData"`
}

func (VpnServerConfigVpnClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRootCertificateArgs) ToVpnServerConfigVpnClientRootCertificateOutput() VpnServerConfigVpnClientRootCertificateOutput {
	return i.ToVpnServerConfigVpnClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRootCertificateArgs) ToVpnServerConfigVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRootCertificateOutput)
}





type VpnServerConfigVpnClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRootCertificateArrayOutput() VpnServerConfigVpnClientRootCertificateArrayOutput
	ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(context.Context) VpnServerConfigVpnClientRootCertificateArrayOutput
}

type VpnServerConfigVpnClientRootCertificateArray []VpnServerConfigVpnClientRootCertificateInput

func (VpnServerConfigVpnClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRootCertificateArray) ToVpnServerConfigVpnClientRootCertificateArrayOutput() VpnServerConfigVpnClientRootCertificateArrayOutput {
	return i.ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRootCertificateArray) ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRootCertificateArrayOutput)
}

type VpnServerConfigVpnClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateOutput) ToVpnServerConfigVpnClientRootCertificateOutput() VpnServerConfigVpnClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateOutput) ToVpnServerConfigVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRootCertificateOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificate) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateArrayOutput) ToVpnServerConfigVpnClientRootCertificateArrayOutput() VpnServerConfigVpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateArrayOutput) ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRootCertificate {
		return vs[0].([]VpnServerConfigVpnClientRootCertificate)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRootCertificateOutput)
}

type VpnServerConfigVpnClientRootCertificateResponse struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}

type VpnServerConfigVpnClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) ToVpnServerConfigVpnClientRootCertificateResponseOutput() VpnServerConfigVpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) ToVpnServerConfigVpnClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificateResponse) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRootCertificateResponseArrayOutput() VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRootCertificateResponse {
		return vs[0].([]VpnServerConfigVpnClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRootCertificateResponseOutput)
}

type VpnServerConfigurationPolicyGroup struct {
	Id            *string                                   `pulumi:"id"`
	IsDefault     *bool                                     `pulumi:"isDefault"`
	Name          *string                                   `pulumi:"name"`
	PolicyMembers []VpnServerConfigurationPolicyGroupMember `pulumi:"policyMembers"`
	Priority      *int                                      `pulumi:"priority"`
}





type VpnServerConfigurationPolicyGroupInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupOutput() VpnServerConfigurationPolicyGroupOutput
	ToVpnServerConfigurationPolicyGroupOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupOutput
}

type VpnServerConfigurationPolicyGroupArgs struct {
	Id            pulumi.StringPtrInput                             `pulumi:"id"`
	IsDefault     pulumi.BoolPtrInput                               `pulumi:"isDefault"`
	Name          pulumi.StringPtrInput                             `pulumi:"name"`
	PolicyMembers VpnServerConfigurationPolicyGroupMemberArrayInput `pulumi:"policyMembers"`
	Priority      pulumi.IntPtrInput                                `pulumi:"priority"`
}

func (VpnServerConfigurationPolicyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroup)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupArgs) ToVpnServerConfigurationPolicyGroupOutput() VpnServerConfigurationPolicyGroupOutput {
	return i.ToVpnServerConfigurationPolicyGroupOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupArgs) ToVpnServerConfigurationPolicyGroupOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupOutput)
}





type VpnServerConfigurationPolicyGroupArrayInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupArrayOutput() VpnServerConfigurationPolicyGroupArrayOutput
	ToVpnServerConfigurationPolicyGroupArrayOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupArrayOutput
}

type VpnServerConfigurationPolicyGroupArray []VpnServerConfigurationPolicyGroupInput

func (VpnServerConfigurationPolicyGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroup)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupArray) ToVpnServerConfigurationPolicyGroupArrayOutput() VpnServerConfigurationPolicyGroupArrayOutput {
	return i.ToVpnServerConfigurationPolicyGroupArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupArray) ToVpnServerConfigurationPolicyGroupArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupArrayOutput)
}

type VpnServerConfigurationPolicyGroupOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroup)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupOutput) ToVpnServerConfigurationPolicyGroupOutput() VpnServerConfigurationPolicyGroupOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupOutput) ToVpnServerConfigurationPolicyGroupOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupOutput) PolicyMembers() VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) []VpnServerConfigurationPolicyGroupMember {
		return v.PolicyMembers
	}).(VpnServerConfigurationPolicyGroupMemberArrayOutput)
}

func (o VpnServerConfigurationPolicyGroupOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

type VpnServerConfigurationPolicyGroupArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroup)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupArrayOutput) ToVpnServerConfigurationPolicyGroupArrayOutput() VpnServerConfigurationPolicyGroupArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupArrayOutput) ToVpnServerConfigurationPolicyGroupArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroup {
		return vs[0].([]VpnServerConfigurationPolicyGroup)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupOutput)
}

type VpnServerConfigurationPolicyGroupMember struct {
	AttributeType  *string `pulumi:"attributeType"`
	AttributeValue *string `pulumi:"attributeValue"`
	Name           *string `pulumi:"name"`
}





type VpnServerConfigurationPolicyGroupMemberInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupMemberOutput() VpnServerConfigurationPolicyGroupMemberOutput
	ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupMemberOutput
}

type VpnServerConfigurationPolicyGroupMemberArgs struct {
	AttributeType  pulumi.StringPtrInput `pulumi:"attributeType"`
	AttributeValue pulumi.StringPtrInput `pulumi:"attributeValue"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
}

func (VpnServerConfigurationPolicyGroupMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupMemberArgs) ToVpnServerConfigurationPolicyGroupMemberOutput() VpnServerConfigurationPolicyGroupMemberOutput {
	return i.ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupMemberArgs) ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupMemberOutput)
}





type VpnServerConfigurationPolicyGroupMemberArrayInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupMemberArrayOutput() VpnServerConfigurationPolicyGroupMemberArrayOutput
	ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupMemberArrayOutput
}

type VpnServerConfigurationPolicyGroupMemberArray []VpnServerConfigurationPolicyGroupMemberInput

func (VpnServerConfigurationPolicyGroupMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupMemberArray) ToVpnServerConfigurationPolicyGroupMemberArrayOutput() VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return i.ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupMemberArray) ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupMemberArrayOutput)
}

type VpnServerConfigurationPolicyGroupMemberOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) ToVpnServerConfigurationPolicyGroupMemberOutput() VpnServerConfigurationPolicyGroupMemberOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) AttributeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMember) *string { return v.AttributeType }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMember) *string { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMember) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VpnServerConfigurationPolicyGroupMemberArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberArrayOutput) ToVpnServerConfigurationPolicyGroupMemberArrayOutput() VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberArrayOutput) ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroupMember {
		return vs[0].([]VpnServerConfigurationPolicyGroupMember)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupMemberOutput)
}

type VpnServerConfigurationPolicyGroupMemberResponse struct {
	AttributeType  *string `pulumi:"attributeType"`
	AttributeValue *string `pulumi:"attributeValue"`
	Name           *string `pulumi:"name"`
}

type VpnServerConfigurationPolicyGroupMemberResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupMemberResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) ToVpnServerConfigurationPolicyGroupMemberResponseOutput() VpnServerConfigurationPolicyGroupMemberResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) ToVpnServerConfigurationPolicyGroupMemberResponseOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) AttributeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMemberResponse) *string { return v.AttributeType }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMemberResponse) *string { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMemberResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VpnServerConfigurationPolicyGroupMemberResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupMemberResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) ToVpnServerConfigurationPolicyGroupMemberResponseArrayOutput() VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) ToVpnServerConfigurationPolicyGroupMemberResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupMemberResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroupMemberResponse {
		return vs[0].([]VpnServerConfigurationPolicyGroupMemberResponse)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupMemberResponseOutput)
}

type VpnServerConfigurationPolicyGroupResponse struct {
	Etag                        string                                            `pulumi:"etag"`
	Id                          *string                                           `pulumi:"id"`
	IsDefault                   *bool                                             `pulumi:"isDefault"`
	Name                        *string                                           `pulumi:"name"`
	P2SConnectionConfigurations []SubResourceResponse                             `pulumi:"p2SConnectionConfigurations"`
	PolicyMembers               []VpnServerConfigurationPolicyGroupMemberResponse `pulumi:"policyMembers"`
	Priority                    *int                                              `pulumi:"priority"`
	ProvisioningState           string                                            `pulumi:"provisioningState"`
	Type                        string                                            `pulumi:"type"`
}

type VpnServerConfigurationPolicyGroupResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) ToVpnServerConfigurationPolicyGroupResponseOutput() VpnServerConfigurationPolicyGroupResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) ToVpnServerConfigurationPolicyGroupResponseOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) P2SConnectionConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) []SubResourceResponse {
		return v.P2SConnectionConfigurations
	}).(SubResourceResponseArrayOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) PolicyMembers() VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) []VpnServerConfigurationPolicyGroupMemberResponse {
		return v.PolicyMembers
	}).(VpnServerConfigurationPolicyGroupMemberResponseArrayOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VpnServerConfigurationPolicyGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupResponseArrayOutput) ToVpnServerConfigurationPolicyGroupResponseArrayOutput() VpnServerConfigurationPolicyGroupResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupResponseArrayOutput) ToVpnServerConfigurationPolicyGroupResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroupResponse {
		return vs[0].([]VpnServerConfigurationPolicyGroupResponse)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupResponseOutput)
}

type VpnSiteLink struct {
	BgpProperties  *VpnLinkBgpSettings        `pulumi:"bgpProperties"`
	Fqdn           *string                    `pulumi:"fqdn"`
	Id             *string                    `pulumi:"id"`
	IpAddress      *string                    `pulumi:"ipAddress"`
	LinkProperties *VpnLinkProviderProperties `pulumi:"linkProperties"`
	Name           *string                    `pulumi:"name"`
}





type VpnSiteLinkInput interface {
	pulumi.Input

	ToVpnSiteLinkOutput() VpnSiteLinkOutput
	ToVpnSiteLinkOutputWithContext(context.Context) VpnSiteLinkOutput
}

type VpnSiteLinkArgs struct {
	BgpProperties  VpnLinkBgpSettingsPtrInput        `pulumi:"bgpProperties"`
	Fqdn           pulumi.StringPtrInput             `pulumi:"fqdn"`
	Id             pulumi.StringPtrInput             `pulumi:"id"`
	IpAddress      pulumi.StringPtrInput             `pulumi:"ipAddress"`
	LinkProperties VpnLinkProviderPropertiesPtrInput `pulumi:"linkProperties"`
	Name           pulumi.StringPtrInput             `pulumi:"name"`
}

func (VpnSiteLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLink)(nil)).Elem()
}

func (i VpnSiteLinkArgs) ToVpnSiteLinkOutput() VpnSiteLinkOutput {
	return i.ToVpnSiteLinkOutputWithContext(context.Background())
}

func (i VpnSiteLinkArgs) ToVpnSiteLinkOutputWithContext(ctx context.Context) VpnSiteLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkOutput)
}





type VpnSiteLinkArrayInput interface {
	pulumi.Input

	ToVpnSiteLinkArrayOutput() VpnSiteLinkArrayOutput
	ToVpnSiteLinkArrayOutputWithContext(context.Context) VpnSiteLinkArrayOutput
}

type VpnSiteLinkArray []VpnSiteLinkInput

func (VpnSiteLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLink)(nil)).Elem()
}

func (i VpnSiteLinkArray) ToVpnSiteLinkArrayOutput() VpnSiteLinkArrayOutput {
	return i.ToVpnSiteLinkArrayOutputWithContext(context.Background())
}

func (i VpnSiteLinkArray) ToVpnSiteLinkArrayOutputWithContext(ctx context.Context) VpnSiteLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkArrayOutput)
}

type VpnSiteLinkOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLink)(nil)).Elem()
}

func (o VpnSiteLinkOutput) ToVpnSiteLinkOutput() VpnSiteLinkOutput {
	return o
}

func (o VpnSiteLinkOutput) ToVpnSiteLinkOutputWithContext(ctx context.Context) VpnSiteLinkOutput {
	return o
}

func (o VpnSiteLinkOutput) BgpProperties() VpnLinkBgpSettingsPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *VpnLinkBgpSettings { return v.BgpProperties }).(VpnLinkBgpSettingsPtrOutput)
}

func (o VpnSiteLinkOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkOutput) LinkProperties() VpnLinkProviderPropertiesPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *VpnLinkProviderProperties { return v.LinkProperties }).(VpnLinkProviderPropertiesPtrOutput)
}

func (o VpnSiteLinkOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VpnSiteLinkArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLink)(nil)).Elem()
}

func (o VpnSiteLinkArrayOutput) ToVpnSiteLinkArrayOutput() VpnSiteLinkArrayOutput {
	return o
}

func (o VpnSiteLinkArrayOutput) ToVpnSiteLinkArrayOutputWithContext(ctx context.Context) VpnSiteLinkArrayOutput {
	return o
}

func (o VpnSiteLinkArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLink {
		return vs[0].([]VpnSiteLink)[vs[1].(int)]
	}).(VpnSiteLinkOutput)
}

type VpnSiteLinkConnection struct {
	ConnectionBandwidth            *int                                       `pulumi:"connectionBandwidth"`
	EgressNatRules                 []SubResource                              `pulumi:"egressNatRules"`
	EnableBgp                      *bool                                      `pulumi:"enableBgp"`
	EnableRateLimiting             *bool                                      `pulumi:"enableRateLimiting"`
	Id                             *string                                    `pulumi:"id"`
	IngressNatRules                []SubResource                              `pulumi:"ingressNatRules"`
	IpsecPolicies                  []IpsecPolicy                              `pulumi:"ipsecPolicies"`
	Name                           *string                                    `pulumi:"name"`
	RoutingWeight                  *int                                       `pulumi:"routingWeight"`
	SharedKey                      *string                                    `pulumi:"sharedKey"`
	UseLocalAzureIpAddress         *bool                                      `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                                      `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string                                    `pulumi:"vpnConnectionProtocolType"`
	VpnGatewayCustomBgpAddresses   []GatewayCustomBgpIpAddressIpConfiguration `pulumi:"vpnGatewayCustomBgpAddresses"`
	VpnLinkConnectionMode          *string                                    `pulumi:"vpnLinkConnectionMode"`
	VpnSiteLink                    *SubResource                               `pulumi:"vpnSiteLink"`
}





type VpnSiteLinkConnectionInput interface {
	pulumi.Input

	ToVpnSiteLinkConnectionOutput() VpnSiteLinkConnectionOutput
	ToVpnSiteLinkConnectionOutputWithContext(context.Context) VpnSiteLinkConnectionOutput
}

type VpnSiteLinkConnectionArgs struct {
	ConnectionBandwidth            pulumi.IntPtrInput                                 `pulumi:"connectionBandwidth"`
	EgressNatRules                 SubResourceArrayInput                              `pulumi:"egressNatRules"`
	EnableBgp                      pulumi.BoolPtrInput                                `pulumi:"enableBgp"`
	EnableRateLimiting             pulumi.BoolPtrInput                                `pulumi:"enableRateLimiting"`
	Id                             pulumi.StringPtrInput                              `pulumi:"id"`
	IngressNatRules                SubResourceArrayInput                              `pulumi:"ingressNatRules"`
	IpsecPolicies                  IpsecPolicyArrayInput                              `pulumi:"ipsecPolicies"`
	Name                           pulumi.StringPtrInput                              `pulumi:"name"`
	RoutingWeight                  pulumi.IntPtrInput                                 `pulumi:"routingWeight"`
	SharedKey                      pulumi.StringPtrInput                              `pulumi:"sharedKey"`
	UseLocalAzureIpAddress         pulumi.BoolPtrInput                                `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrInput                                `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      pulumi.StringPtrInput                              `pulumi:"vpnConnectionProtocolType"`
	VpnGatewayCustomBgpAddresses   GatewayCustomBgpIpAddressIpConfigurationArrayInput `pulumi:"vpnGatewayCustomBgpAddresses"`
	VpnLinkConnectionMode          pulumi.StringPtrInput                              `pulumi:"vpnLinkConnectionMode"`
	VpnSiteLink                    SubResourcePtrInput                                `pulumi:"vpnSiteLink"`
}

func (VpnSiteLinkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkConnection)(nil)).Elem()
}

func (i VpnSiteLinkConnectionArgs) ToVpnSiteLinkConnectionOutput() VpnSiteLinkConnectionOutput {
	return i.ToVpnSiteLinkConnectionOutputWithContext(context.Background())
}

func (i VpnSiteLinkConnectionArgs) ToVpnSiteLinkConnectionOutputWithContext(ctx context.Context) VpnSiteLinkConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkConnectionOutput)
}





type VpnSiteLinkConnectionArrayInput interface {
	pulumi.Input

	ToVpnSiteLinkConnectionArrayOutput() VpnSiteLinkConnectionArrayOutput
	ToVpnSiteLinkConnectionArrayOutputWithContext(context.Context) VpnSiteLinkConnectionArrayOutput
}

type VpnSiteLinkConnectionArray []VpnSiteLinkConnectionInput

func (VpnSiteLinkConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkConnection)(nil)).Elem()
}

func (i VpnSiteLinkConnectionArray) ToVpnSiteLinkConnectionArrayOutput() VpnSiteLinkConnectionArrayOutput {
	return i.ToVpnSiteLinkConnectionArrayOutputWithContext(context.Background())
}

func (i VpnSiteLinkConnectionArray) ToVpnSiteLinkConnectionArrayOutputWithContext(ctx context.Context) VpnSiteLinkConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkConnectionArrayOutput)
}

type VpnSiteLinkConnectionOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkConnection)(nil)).Elem()
}

func (o VpnSiteLinkConnectionOutput) ToVpnSiteLinkConnectionOutput() VpnSiteLinkConnectionOutput {
	return o
}

func (o VpnSiteLinkConnectionOutput) ToVpnSiteLinkConnectionOutputWithContext(ctx context.Context) VpnSiteLinkConnectionOutput {
	return o
}

func (o VpnSiteLinkConnectionOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) EgressNatRules() SubResourceArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []SubResource { return v.EgressNatRules }).(SubResourceArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) IngressNatRules() SubResourceArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []SubResource { return v.IngressNatRules }).(SubResourceArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) IpsecPolicies() IpsecPolicyArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []IpsecPolicy { return v.IpsecPolicies }).(IpsecPolicyArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnGatewayCustomBgpAddresses() GatewayCustomBgpIpAddressIpConfigurationArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []GatewayCustomBgpIpAddressIpConfiguration {
		return v.VpnGatewayCustomBgpAddresses
	}).(GatewayCustomBgpIpAddressIpConfigurationArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnLinkConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.VpnLinkConnectionMode }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnSiteLink() SubResourcePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *SubResource { return v.VpnSiteLink }).(SubResourcePtrOutput)
}

type VpnSiteLinkConnectionArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkConnection)(nil)).Elem()
}

func (o VpnSiteLinkConnectionArrayOutput) ToVpnSiteLinkConnectionArrayOutput() VpnSiteLinkConnectionArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionArrayOutput) ToVpnSiteLinkConnectionArrayOutputWithContext(ctx context.Context) VpnSiteLinkConnectionArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLinkConnection {
		return vs[0].([]VpnSiteLinkConnection)[vs[1].(int)]
	}).(VpnSiteLinkConnectionOutput)
}

type VpnSiteLinkConnectionResponse struct {
	ConnectionBandwidth            *int                                               `pulumi:"connectionBandwidth"`
	ConnectionStatus               string                                             `pulumi:"connectionStatus"`
	EgressBytesTransferred         float64                                            `pulumi:"egressBytesTransferred"`
	EgressNatRules                 []SubResourceResponse                              `pulumi:"egressNatRules"`
	EnableBgp                      *bool                                              `pulumi:"enableBgp"`
	EnableRateLimiting             *bool                                              `pulumi:"enableRateLimiting"`
	Etag                           string                                             `pulumi:"etag"`
	Id                             *string                                            `pulumi:"id"`
	IngressBytesTransferred        float64                                            `pulumi:"ingressBytesTransferred"`
	IngressNatRules                []SubResourceResponse                              `pulumi:"ingressNatRules"`
	IpsecPolicies                  []IpsecPolicyResponse                              `pulumi:"ipsecPolicies"`
	Name                           *string                                            `pulumi:"name"`
	ProvisioningState              string                                             `pulumi:"provisioningState"`
	RoutingWeight                  *int                                               `pulumi:"routingWeight"`
	SharedKey                      *string                                            `pulumi:"sharedKey"`
	Type                           string                                             `pulumi:"type"`
	UseLocalAzureIpAddress         *bool                                              `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                                              `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string                                            `pulumi:"vpnConnectionProtocolType"`
	VpnGatewayCustomBgpAddresses   []GatewayCustomBgpIpAddressIpConfigurationResponse `pulumi:"vpnGatewayCustomBgpAddresses"`
	VpnLinkConnectionMode          *string                                            `pulumi:"vpnLinkConnectionMode"`
	VpnSiteLink                    *SubResourceResponse                               `pulumi:"vpnSiteLink"`
}

type VpnSiteLinkConnectionResponseOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkConnectionResponse)(nil)).Elem()
}

func (o VpnSiteLinkConnectionResponseOutput) ToVpnSiteLinkConnectionResponseOutput() VpnSiteLinkConnectionResponseOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseOutput) ToVpnSiteLinkConnectionResponseOutputWithContext(ctx context.Context) VpnSiteLinkConnectionResponseOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnSiteLinkConnectionResponseOutput) EgressNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []SubResourceResponse { return v.EgressNatRules }).(SubResourceResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnSiteLinkConnectionResponseOutput) IngressNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []SubResourceResponse { return v.IngressNatRules }).(SubResourceResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnGatewayCustomBgpAddresses() GatewayCustomBgpIpAddressIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []GatewayCustomBgpIpAddressIpConfigurationResponse {
		return v.VpnGatewayCustomBgpAddresses
	}).(GatewayCustomBgpIpAddressIpConfigurationResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnLinkConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.VpnLinkConnectionMode }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnSiteLink() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *SubResourceResponse { return v.VpnSiteLink }).(SubResourceResponsePtrOutput)
}

type VpnSiteLinkConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkConnectionResponse)(nil)).Elem()
}

func (o VpnSiteLinkConnectionResponseArrayOutput) ToVpnSiteLinkConnectionResponseArrayOutput() VpnSiteLinkConnectionResponseArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseArrayOutput) ToVpnSiteLinkConnectionResponseArrayOutputWithContext(ctx context.Context) VpnSiteLinkConnectionResponseArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLinkConnectionResponse {
		return vs[0].([]VpnSiteLinkConnectionResponse)[vs[1].(int)]
	}).(VpnSiteLinkConnectionResponseOutput)
}

type VpnSiteLinkResponse struct {
	BgpProperties     *VpnLinkBgpSettingsResponse        `pulumi:"bgpProperties"`
	Etag              string                             `pulumi:"etag"`
	Fqdn              *string                            `pulumi:"fqdn"`
	Id                *string                            `pulumi:"id"`
	IpAddress         *string                            `pulumi:"ipAddress"`
	LinkProperties    *VpnLinkProviderPropertiesResponse `pulumi:"linkProperties"`
	Name              *string                            `pulumi:"name"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	Type              string                             `pulumi:"type"`
}

type VpnSiteLinkResponseOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkResponse)(nil)).Elem()
}

func (o VpnSiteLinkResponseOutput) ToVpnSiteLinkResponseOutput() VpnSiteLinkResponseOutput {
	return o
}

func (o VpnSiteLinkResponseOutput) ToVpnSiteLinkResponseOutputWithContext(ctx context.Context) VpnSiteLinkResponseOutput {
	return o
}

func (o VpnSiteLinkResponseOutput) BgpProperties() VpnLinkBgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *VpnLinkBgpSettingsResponse { return v.BgpProperties }).(VpnLinkBgpSettingsResponsePtrOutput)
}

func (o VpnSiteLinkResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnSiteLinkResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) LinkProperties() VpnLinkProviderPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *VpnLinkProviderPropertiesResponse { return v.LinkProperties }).(VpnLinkProviderPropertiesResponsePtrOutput)
}

func (o VpnSiteLinkResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnSiteLinkResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VpnSiteLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkResponse)(nil)).Elem()
}

func (o VpnSiteLinkResponseArrayOutput) ToVpnSiteLinkResponseArrayOutput() VpnSiteLinkResponseArrayOutput {
	return o
}

func (o VpnSiteLinkResponseArrayOutput) ToVpnSiteLinkResponseArrayOutputWithContext(ctx context.Context) VpnSiteLinkResponseArrayOutput {
	return o
}

func (o VpnSiteLinkResponseArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLinkResponse {
		return vs[0].([]VpnSiteLinkResponse)[vs[1].(int)]
	}).(VpnSiteLinkResponseOutput)
}

type WebApplicationFirewallCustomRule struct {
	Action          string           `pulumi:"action"`
	MatchConditions []MatchCondition `pulumi:"matchConditions"`
	Name            *string          `pulumi:"name"`
	Priority        int              `pulumi:"priority"`
	RuleType        string           `pulumi:"ruleType"`
}





type WebApplicationFirewallCustomRuleInput interface {
	pulumi.Input

	ToWebApplicationFirewallCustomRuleOutput() WebApplicationFirewallCustomRuleOutput
	ToWebApplicationFirewallCustomRuleOutputWithContext(context.Context) WebApplicationFirewallCustomRuleOutput
}

type WebApplicationFirewallCustomRuleArgs struct {
	Action          pulumi.StringInput       `pulumi:"action"`
	MatchConditions MatchConditionArrayInput `pulumi:"matchConditions"`
	Name            pulumi.StringPtrInput    `pulumi:"name"`
	Priority        pulumi.IntInput          `pulumi:"priority"`
	RuleType        pulumi.StringInput       `pulumi:"ruleType"`
}

func (WebApplicationFirewallCustomRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (i WebApplicationFirewallCustomRuleArgs) ToWebApplicationFirewallCustomRuleOutput() WebApplicationFirewallCustomRuleOutput {
	return i.ToWebApplicationFirewallCustomRuleOutputWithContext(context.Background())
}

func (i WebApplicationFirewallCustomRuleArgs) ToWebApplicationFirewallCustomRuleOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebApplicationFirewallCustomRuleOutput)
}





type WebApplicationFirewallCustomRuleArrayInput interface {
	pulumi.Input

	ToWebApplicationFirewallCustomRuleArrayOutput() WebApplicationFirewallCustomRuleArrayOutput
	ToWebApplicationFirewallCustomRuleArrayOutputWithContext(context.Context) WebApplicationFirewallCustomRuleArrayOutput
}

type WebApplicationFirewallCustomRuleArray []WebApplicationFirewallCustomRuleInput

func (WebApplicationFirewallCustomRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (i WebApplicationFirewallCustomRuleArray) ToWebApplicationFirewallCustomRuleArrayOutput() WebApplicationFirewallCustomRuleArrayOutput {
	return i.ToWebApplicationFirewallCustomRuleArrayOutputWithContext(context.Background())
}

func (i WebApplicationFirewallCustomRuleArray) ToWebApplicationFirewallCustomRuleArrayOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebApplicationFirewallCustomRuleArrayOutput)
}

type WebApplicationFirewallCustomRuleOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleOutput) ToWebApplicationFirewallCustomRuleOutput() WebApplicationFirewallCustomRuleOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleOutput) ToWebApplicationFirewallCustomRuleOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) string { return v.Action }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) MatchConditions() MatchConditionArrayOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) []MatchCondition { return v.MatchConditions }).(MatchConditionArrayOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) int { return v.Priority }).(pulumi.IntOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type WebApplicationFirewallCustomRuleArrayOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleArrayOutput) ToWebApplicationFirewallCustomRuleArrayOutput() WebApplicationFirewallCustomRuleArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleArrayOutput) ToWebApplicationFirewallCustomRuleArrayOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleArrayOutput) Index(i pulumi.IntInput) WebApplicationFirewallCustomRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebApplicationFirewallCustomRule {
		return vs[0].([]WebApplicationFirewallCustomRule)[vs[1].(int)]
	}).(WebApplicationFirewallCustomRuleOutput)
}

type WebApplicationFirewallCustomRuleResponse struct {
	Action          string                   `pulumi:"action"`
	Etag            string                   `pulumi:"etag"`
	MatchConditions []MatchConditionResponse `pulumi:"matchConditions"`
	Name            *string                  `pulumi:"name"`
	Priority        int                      `pulumi:"priority"`
	RuleType        string                   `pulumi:"ruleType"`
}

type WebApplicationFirewallCustomRuleResponseOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallCustomRuleResponse)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleResponseOutput) ToWebApplicationFirewallCustomRuleResponseOutput() WebApplicationFirewallCustomRuleResponseOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseOutput) ToWebApplicationFirewallCustomRuleResponseOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleResponseOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) MatchConditions() MatchConditionResponseArrayOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) []MatchConditionResponse { return v.MatchConditions }).(MatchConditionResponseArrayOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

type WebApplicationFirewallCustomRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApplicationFirewallCustomRuleResponse)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleResponseArrayOutput) ToWebApplicationFirewallCustomRuleResponseArrayOutput() WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseArrayOutput) ToWebApplicationFirewallCustomRuleResponseArrayOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseArrayOutput) Index(i pulumi.IntInput) WebApplicationFirewallCustomRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebApplicationFirewallCustomRuleResponse {
		return vs[0].([]WebApplicationFirewallCustomRuleResponse)[vs[1].(int)]
	}).(WebApplicationFirewallCustomRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnLinkBgpSettingsResponseOutput{})
	pulumi.RegisterOutputType(VpnLinkBgpSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingArrayOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingResponseOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionResponseOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkResponseOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleArrayOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleResponseOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleResponseArrayOutput{})
}
