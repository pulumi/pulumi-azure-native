


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type GalleryImageStatusResponse struct {
	ErrorCode    *string `pulumi:"errorCode"`
	ErrorMessage *string `pulumi:"errorMessage"`
}

type GalleryImageStatusResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageStatusResponse)(nil)).Elem()
}

func (o GalleryImageStatusResponseOutput) ToGalleryImageStatusResponseOutput() GalleryImageStatusResponseOutput {
	return o
}

func (o GalleryImageStatusResponseOutput) ToGalleryImageStatusResponseOutputWithContext(ctx context.Context) GalleryImageStatusResponseOutput {
	return o
}

func (o GalleryImageStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o GalleryImageStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

type IPPool struct {
	End        *string         `pulumi:"end"`
	IpPoolType *IPPoolTypeEnum `pulumi:"ipPoolType"`
	Start      *string         `pulumi:"start"`
}





type IPPoolInput interface {
	pulumi.Input

	ToIPPoolOutput() IPPoolOutput
	ToIPPoolOutputWithContext(context.Context) IPPoolOutput
}

type IPPoolArgs struct {
	End        pulumi.StringPtrInput  `pulumi:"end"`
	IpPoolType IPPoolTypeEnumPtrInput `pulumi:"ipPoolType"`
	Start      pulumi.StringPtrInput  `pulumi:"start"`
}

func (IPPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPool)(nil)).Elem()
}

func (i IPPoolArgs) ToIPPoolOutput() IPPoolOutput {
	return i.ToIPPoolOutputWithContext(context.Background())
}

func (i IPPoolArgs) ToIPPoolOutputWithContext(ctx context.Context) IPPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPPoolOutput)
}





type IPPoolArrayInput interface {
	pulumi.Input

	ToIPPoolArrayOutput() IPPoolArrayOutput
	ToIPPoolArrayOutputWithContext(context.Context) IPPoolArrayOutput
}

type IPPoolArray []IPPoolInput

func (IPPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPPool)(nil)).Elem()
}

func (i IPPoolArray) ToIPPoolArrayOutput() IPPoolArrayOutput {
	return i.ToIPPoolArrayOutputWithContext(context.Background())
}

func (i IPPoolArray) ToIPPoolArrayOutputWithContext(ctx context.Context) IPPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPPoolArrayOutput)
}

type IPPoolOutput struct{ *pulumi.OutputState }

func (IPPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPool)(nil)).Elem()
}

func (o IPPoolOutput) ToIPPoolOutput() IPPoolOutput {
	return o
}

func (o IPPoolOutput) ToIPPoolOutputWithContext(ctx context.Context) IPPoolOutput {
	return o
}

func (o IPPoolOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPool) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o IPPoolOutput) IpPoolType() IPPoolTypeEnumPtrOutput {
	return o.ApplyT(func(v IPPool) *IPPoolTypeEnum { return v.IpPoolType }).(IPPoolTypeEnumPtrOutput)
}

func (o IPPoolOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPool) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type IPPoolArrayOutput struct{ *pulumi.OutputState }

func (IPPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPPool)(nil)).Elem()
}

func (o IPPoolArrayOutput) ToIPPoolArrayOutput() IPPoolArrayOutput {
	return o
}

func (o IPPoolArrayOutput) ToIPPoolArrayOutputWithContext(ctx context.Context) IPPoolArrayOutput {
	return o
}

func (o IPPoolArrayOutput) Index(i pulumi.IntInput) IPPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPPool {
		return vs[0].([]IPPool)[vs[1].(int)]
	}).(IPPoolOutput)
}

type IPPoolInfoResponse struct {
	Available string `pulumi:"available"`
	Used      string `pulumi:"used"`
}

type IPPoolInfoResponseOutput struct{ *pulumi.OutputState }

func (IPPoolInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPoolInfoResponse)(nil)).Elem()
}

func (o IPPoolInfoResponseOutput) ToIPPoolInfoResponseOutput() IPPoolInfoResponseOutput {
	return o
}

func (o IPPoolInfoResponseOutput) ToIPPoolInfoResponseOutputWithContext(ctx context.Context) IPPoolInfoResponseOutput {
	return o
}

func (o IPPoolInfoResponseOutput) Available() pulumi.StringOutput {
	return o.ApplyT(func(v IPPoolInfoResponse) string { return v.Available }).(pulumi.StringOutput)
}

func (o IPPoolInfoResponseOutput) Used() pulumi.StringOutput {
	return o.ApplyT(func(v IPPoolInfoResponse) string { return v.Used }).(pulumi.StringOutput)
}

type IPPoolInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (IPPoolInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPPoolInfoResponse)(nil)).Elem()
}

func (o IPPoolInfoResponsePtrOutput) ToIPPoolInfoResponsePtrOutput() IPPoolInfoResponsePtrOutput {
	return o
}

func (o IPPoolInfoResponsePtrOutput) ToIPPoolInfoResponsePtrOutputWithContext(ctx context.Context) IPPoolInfoResponsePtrOutput {
	return o
}

func (o IPPoolInfoResponsePtrOutput) Elem() IPPoolInfoResponseOutput {
	return o.ApplyT(func(v *IPPoolInfoResponse) IPPoolInfoResponse {
		if v != nil {
			return *v
		}
		var ret IPPoolInfoResponse
		return ret
	}).(IPPoolInfoResponseOutput)
}

func (o IPPoolInfoResponsePtrOutput) Available() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPPoolInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Available
	}).(pulumi.StringPtrOutput)
}

func (o IPPoolInfoResponsePtrOutput) Used() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPPoolInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Used
	}).(pulumi.StringPtrOutput)
}

type IPPoolResponse struct {
	End        *string             `pulumi:"end"`
	Info       *IPPoolInfoResponse `pulumi:"info"`
	IpPoolType *string             `pulumi:"ipPoolType"`
	Start      *string             `pulumi:"start"`
}

type IPPoolResponseOutput struct{ *pulumi.OutputState }

func (IPPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPoolResponse)(nil)).Elem()
}

func (o IPPoolResponseOutput) ToIPPoolResponseOutput() IPPoolResponseOutput {
	return o
}

func (o IPPoolResponseOutput) ToIPPoolResponseOutputWithContext(ctx context.Context) IPPoolResponseOutput {
	return o
}

func (o IPPoolResponseOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPoolResponse) *string { return v.End }).(pulumi.StringPtrOutput)
}

func (o IPPoolResponseOutput) Info() IPPoolInfoResponsePtrOutput {
	return o.ApplyT(func(v IPPoolResponse) *IPPoolInfoResponse { return v.Info }).(IPPoolInfoResponsePtrOutput)
}

func (o IPPoolResponseOutput) IpPoolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPoolResponse) *string { return v.IpPoolType }).(pulumi.StringPtrOutput)
}

func (o IPPoolResponseOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPPoolResponse) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type IPPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (IPPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPPoolResponse)(nil)).Elem()
}

func (o IPPoolResponseArrayOutput) ToIPPoolResponseArrayOutput() IPPoolResponseArrayOutput {
	return o
}

func (o IPPoolResponseArrayOutput) ToIPPoolResponseArrayOutputWithContext(ctx context.Context) IPPoolResponseArrayOutput {
	return o
}

func (o IPPoolResponseArrayOutput) Index(i pulumi.IntInput) IPPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPPoolResponse {
		return vs[0].([]IPPoolResponse)[vs[1].(int)]
	}).(IPPoolResponseOutput)
}

type InterfaceDNSSettings struct {
	DnsServers []string `pulumi:"dnsServers"`
}





type InterfaceDNSSettingsInput interface {
	pulumi.Input

	ToInterfaceDNSSettingsOutput() InterfaceDNSSettingsOutput
	ToInterfaceDNSSettingsOutputWithContext(context.Context) InterfaceDNSSettingsOutput
}

type InterfaceDNSSettingsArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
}

func (InterfaceDNSSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InterfaceDNSSettings)(nil)).Elem()
}

func (i InterfaceDNSSettingsArgs) ToInterfaceDNSSettingsOutput() InterfaceDNSSettingsOutput {
	return i.ToInterfaceDNSSettingsOutputWithContext(context.Background())
}

func (i InterfaceDNSSettingsArgs) ToInterfaceDNSSettingsOutputWithContext(ctx context.Context) InterfaceDNSSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceDNSSettingsOutput)
}

func (i InterfaceDNSSettingsArgs) ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput {
	return i.ToInterfaceDNSSettingsPtrOutputWithContext(context.Background())
}

func (i InterfaceDNSSettingsArgs) ToInterfaceDNSSettingsPtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceDNSSettingsOutput).ToInterfaceDNSSettingsPtrOutputWithContext(ctx)
}









type InterfaceDNSSettingsPtrInput interface {
	pulumi.Input

	ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput
	ToInterfaceDNSSettingsPtrOutputWithContext(context.Context) InterfaceDNSSettingsPtrOutput
}

type interfaceDNSSettingsPtrType InterfaceDNSSettingsArgs

func InterfaceDNSSettingsPtr(v *InterfaceDNSSettingsArgs) InterfaceDNSSettingsPtrInput {
	return (*interfaceDNSSettingsPtrType)(v)
}

func (*interfaceDNSSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceDNSSettings)(nil)).Elem()
}

func (i *interfaceDNSSettingsPtrType) ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput {
	return i.ToInterfaceDNSSettingsPtrOutputWithContext(context.Background())
}

func (i *interfaceDNSSettingsPtrType) ToInterfaceDNSSettingsPtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceDNSSettingsPtrOutput)
}

type InterfaceDNSSettingsOutput struct{ *pulumi.OutputState }

func (InterfaceDNSSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InterfaceDNSSettings)(nil)).Elem()
}

func (o InterfaceDNSSettingsOutput) ToInterfaceDNSSettingsOutput() InterfaceDNSSettingsOutput {
	return o
}

func (o InterfaceDNSSettingsOutput) ToInterfaceDNSSettingsOutputWithContext(ctx context.Context) InterfaceDNSSettingsOutput {
	return o
}

func (o InterfaceDNSSettingsOutput) ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput {
	return o.ToInterfaceDNSSettingsPtrOutputWithContext(context.Background())
}

func (o InterfaceDNSSettingsOutput) ToInterfaceDNSSettingsPtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InterfaceDNSSettings) *InterfaceDNSSettings {
		return &v
	}).(InterfaceDNSSettingsPtrOutput)
}

func (o InterfaceDNSSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InterfaceDNSSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type InterfaceDNSSettingsPtrOutput struct{ *pulumi.OutputState }

func (InterfaceDNSSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceDNSSettings)(nil)).Elem()
}

func (o InterfaceDNSSettingsPtrOutput) ToInterfaceDNSSettingsPtrOutput() InterfaceDNSSettingsPtrOutput {
	return o
}

func (o InterfaceDNSSettingsPtrOutput) ToInterfaceDNSSettingsPtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsPtrOutput {
	return o
}

func (o InterfaceDNSSettingsPtrOutput) Elem() InterfaceDNSSettingsOutput {
	return o.ApplyT(func(v *InterfaceDNSSettings) InterfaceDNSSettings {
		if v != nil {
			return *v
		}
		var ret InterfaceDNSSettings
		return ret
	}).(InterfaceDNSSettingsOutput)
}

func (o InterfaceDNSSettingsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InterfaceDNSSettings) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type InterfaceDNSSettingsResponse struct {
	DnsServers []string `pulumi:"dnsServers"`
}

type InterfaceDNSSettingsResponseOutput struct{ *pulumi.OutputState }

func (InterfaceDNSSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InterfaceDNSSettingsResponse)(nil)).Elem()
}

func (o InterfaceDNSSettingsResponseOutput) ToInterfaceDNSSettingsResponseOutput() InterfaceDNSSettingsResponseOutput {
	return o
}

func (o InterfaceDNSSettingsResponseOutput) ToInterfaceDNSSettingsResponseOutputWithContext(ctx context.Context) InterfaceDNSSettingsResponseOutput {
	return o
}

func (o InterfaceDNSSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InterfaceDNSSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type InterfaceDNSSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (InterfaceDNSSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceDNSSettingsResponse)(nil)).Elem()
}

func (o InterfaceDNSSettingsResponsePtrOutput) ToInterfaceDNSSettingsResponsePtrOutput() InterfaceDNSSettingsResponsePtrOutput {
	return o
}

func (o InterfaceDNSSettingsResponsePtrOutput) ToInterfaceDNSSettingsResponsePtrOutputWithContext(ctx context.Context) InterfaceDNSSettingsResponsePtrOutput {
	return o
}

func (o InterfaceDNSSettingsResponsePtrOutput) Elem() InterfaceDNSSettingsResponseOutput {
	return o.ApplyT(func(v *InterfaceDNSSettingsResponse) InterfaceDNSSettingsResponse {
		if v != nil {
			return *v
		}
		var ret InterfaceDNSSettingsResponse
		return ret
	}).(InterfaceDNSSettingsResponseOutput)
}

func (o InterfaceDNSSettingsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InterfaceDNSSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type IpConfiguration struct {
	Name       *string                    `pulumi:"name"`
	Properties *IpConfigurationProperties `pulumi:"properties"`
}





type IpConfigurationInput interface {
	pulumi.Input

	ToIpConfigurationOutput() IpConfigurationOutput
	ToIpConfigurationOutputWithContext(context.Context) IpConfigurationOutput
}

type IpConfigurationArgs struct {
	Name       pulumi.StringPtrInput             `pulumi:"name"`
	Properties IpConfigurationPropertiesPtrInput `pulumi:"properties"`
}

func (IpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfiguration)(nil)).Elem()
}

func (i IpConfigurationArgs) ToIpConfigurationOutput() IpConfigurationOutput {
	return i.ToIpConfigurationOutputWithContext(context.Background())
}

func (i IpConfigurationArgs) ToIpConfigurationOutputWithContext(ctx context.Context) IpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationOutput)
}





type IpConfigurationArrayInput interface {
	pulumi.Input

	ToIpConfigurationArrayOutput() IpConfigurationArrayOutput
	ToIpConfigurationArrayOutputWithContext(context.Context) IpConfigurationArrayOutput
}

type IpConfigurationArray []IpConfigurationInput

func (IpConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpConfiguration)(nil)).Elem()
}

func (i IpConfigurationArray) ToIpConfigurationArrayOutput() IpConfigurationArrayOutput {
	return i.ToIpConfigurationArrayOutputWithContext(context.Background())
}

func (i IpConfigurationArray) ToIpConfigurationArrayOutputWithContext(ctx context.Context) IpConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationArrayOutput)
}

type IpConfigurationOutput struct{ *pulumi.OutputState }

func (IpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfiguration)(nil)).Elem()
}

func (o IpConfigurationOutput) ToIpConfigurationOutput() IpConfigurationOutput {
	return o
}

func (o IpConfigurationOutput) ToIpConfigurationOutputWithContext(ctx context.Context) IpConfigurationOutput {
	return o
}

func (o IpConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationOutput) Properties() IpConfigurationPropertiesPtrOutput {
	return o.ApplyT(func(v IpConfiguration) *IpConfigurationProperties { return v.Properties }).(IpConfigurationPropertiesPtrOutput)
}

type IpConfigurationArrayOutput struct{ *pulumi.OutputState }

func (IpConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpConfiguration)(nil)).Elem()
}

func (o IpConfigurationArrayOutput) ToIpConfigurationArrayOutput() IpConfigurationArrayOutput {
	return o
}

func (o IpConfigurationArrayOutput) ToIpConfigurationArrayOutputWithContext(ctx context.Context) IpConfigurationArrayOutput {
	return o
}

func (o IpConfigurationArrayOutput) Index(i pulumi.IntInput) IpConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpConfiguration {
		return vs[0].([]IpConfiguration)[vs[1].(int)]
	}).(IpConfigurationOutput)
}

type IpConfigurationProperties struct {
	PrefixLength              *string                `pulumi:"prefixLength"`
	PrivateIPAddress          *string                `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                `pulumi:"privateIPAllocationMethod"`
	Subnet                    *IpConfigurationSubnet `pulumi:"subnet"`
}





type IpConfigurationPropertiesInput interface {
	pulumi.Input

	ToIpConfigurationPropertiesOutput() IpConfigurationPropertiesOutput
	ToIpConfigurationPropertiesOutputWithContext(context.Context) IpConfigurationPropertiesOutput
}

type IpConfigurationPropertiesArgs struct {
	PrefixLength              pulumi.StringPtrInput         `pulumi:"prefixLength"`
	PrivateIPAddress          pulumi.StringPtrInput         `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput         `pulumi:"privateIPAllocationMethod"`
	Subnet                    IpConfigurationSubnetPtrInput `pulumi:"subnet"`
}

func (IpConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationProperties)(nil)).Elem()
}

func (i IpConfigurationPropertiesArgs) ToIpConfigurationPropertiesOutput() IpConfigurationPropertiesOutput {
	return i.ToIpConfigurationPropertiesOutputWithContext(context.Background())
}

func (i IpConfigurationPropertiesArgs) ToIpConfigurationPropertiesOutputWithContext(ctx context.Context) IpConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationPropertiesOutput)
}

func (i IpConfigurationPropertiesArgs) ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput {
	return i.ToIpConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i IpConfigurationPropertiesArgs) ToIpConfigurationPropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationPropertiesOutput).ToIpConfigurationPropertiesPtrOutputWithContext(ctx)
}









type IpConfigurationPropertiesPtrInput interface {
	pulumi.Input

	ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput
	ToIpConfigurationPropertiesPtrOutputWithContext(context.Context) IpConfigurationPropertiesPtrOutput
}

type ipConfigurationPropertiesPtrType IpConfigurationPropertiesArgs

func IpConfigurationPropertiesPtr(v *IpConfigurationPropertiesArgs) IpConfigurationPropertiesPtrInput {
	return (*ipConfigurationPropertiesPtrType)(v)
}

func (*ipConfigurationPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationProperties)(nil)).Elem()
}

func (i *ipConfigurationPropertiesPtrType) ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput {
	return i.ToIpConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (i *ipConfigurationPropertiesPtrType) ToIpConfigurationPropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationPropertiesPtrOutput)
}

type IpConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (IpConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationProperties)(nil)).Elem()
}

func (o IpConfigurationPropertiesOutput) ToIpConfigurationPropertiesOutput() IpConfigurationPropertiesOutput {
	return o
}

func (o IpConfigurationPropertiesOutput) ToIpConfigurationPropertiesOutputWithContext(ctx context.Context) IpConfigurationPropertiesOutput {
	return o
}

func (o IpConfigurationPropertiesOutput) ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput {
	return o.ToIpConfigurationPropertiesPtrOutputWithContext(context.Background())
}

func (o IpConfigurationPropertiesOutput) ToIpConfigurationPropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpConfigurationProperties) *IpConfigurationProperties {
		return &v
	}).(IpConfigurationPropertiesPtrOutput)
}

func (o IpConfigurationPropertiesOutput) PrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationProperties) *string { return v.PrefixLength }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationProperties) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationProperties) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesOutput) Subnet() IpConfigurationSubnetPtrOutput {
	return o.ApplyT(func(v IpConfigurationProperties) *IpConfigurationSubnet { return v.Subnet }).(IpConfigurationSubnetPtrOutput)
}

type IpConfigurationPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IpConfigurationPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationProperties)(nil)).Elem()
}

func (o IpConfigurationPropertiesPtrOutput) ToIpConfigurationPropertiesPtrOutput() IpConfigurationPropertiesPtrOutput {
	return o
}

func (o IpConfigurationPropertiesPtrOutput) ToIpConfigurationPropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationPropertiesPtrOutput {
	return o
}

func (o IpConfigurationPropertiesPtrOutput) Elem() IpConfigurationPropertiesOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) IpConfigurationProperties {
		if v != nil {
			return *v
		}
		var ret IpConfigurationProperties
		return ret
	}).(IpConfigurationPropertiesOutput)
}

func (o IpConfigurationPropertiesPtrOutput) PrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrefixLength
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesPtrOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAddress
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesPtrOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationPropertiesPtrOutput) Subnet() IpConfigurationSubnetPtrOutput {
	return o.ApplyT(func(v *IpConfigurationProperties) *IpConfigurationSubnet {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(IpConfigurationSubnetPtrOutput)
}

type IpConfigurationResponse struct {
	Name       *string                            `pulumi:"name"`
	Properties *IpConfigurationResponseProperties `pulumi:"properties"`
}

type IpConfigurationResponseOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationResponse)(nil)).Elem()
}

func (o IpConfigurationResponseOutput) ToIpConfigurationResponseOutput() IpConfigurationResponseOutput {
	return o
}

func (o IpConfigurationResponseOutput) ToIpConfigurationResponseOutputWithContext(ctx context.Context) IpConfigurationResponseOutput {
	return o
}

func (o IpConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponseOutput) Properties() IpConfigurationResponsePropertiesPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponse) *IpConfigurationResponseProperties { return v.Properties }).(IpConfigurationResponsePropertiesPtrOutput)
}

type IpConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpConfigurationResponse)(nil)).Elem()
}

func (o IpConfigurationResponseArrayOutput) ToIpConfigurationResponseArrayOutput() IpConfigurationResponseArrayOutput {
	return o
}

func (o IpConfigurationResponseArrayOutput) ToIpConfigurationResponseArrayOutputWithContext(ctx context.Context) IpConfigurationResponseArrayOutput {
	return o
}

func (o IpConfigurationResponseArrayOutput) Index(i pulumi.IntInput) IpConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpConfigurationResponse {
		return vs[0].([]IpConfigurationResponse)[vs[1].(int)]
	}).(IpConfigurationResponseOutput)
}

type IpConfigurationResponseProperties struct {
	PrefixLength              *string                        `pulumi:"prefixLength"`
	PrivateIPAddress          *string                        `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                        `pulumi:"privateIPAllocationMethod"`
	Subnet                    *IpConfigurationResponseSubnet `pulumi:"subnet"`
}

type IpConfigurationResponsePropertiesOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationResponseProperties)(nil)).Elem()
}

func (o IpConfigurationResponsePropertiesOutput) ToIpConfigurationResponsePropertiesOutput() IpConfigurationResponsePropertiesOutput {
	return o
}

func (o IpConfigurationResponsePropertiesOutput) ToIpConfigurationResponsePropertiesOutputWithContext(ctx context.Context) IpConfigurationResponsePropertiesOutput {
	return o
}

func (o IpConfigurationResponsePropertiesOutput) PrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseProperties) *string { return v.PrefixLength }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseProperties) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseProperties) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesOutput) Subnet() IpConfigurationResponseSubnetPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseProperties) *IpConfigurationResponseSubnet { return v.Subnet }).(IpConfigurationResponseSubnetPtrOutput)
}

type IpConfigurationResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationResponseProperties)(nil)).Elem()
}

func (o IpConfigurationResponsePropertiesPtrOutput) ToIpConfigurationResponsePropertiesPtrOutput() IpConfigurationResponsePropertiesPtrOutput {
	return o
}

func (o IpConfigurationResponsePropertiesPtrOutput) ToIpConfigurationResponsePropertiesPtrOutputWithContext(ctx context.Context) IpConfigurationResponsePropertiesPtrOutput {
	return o
}

func (o IpConfigurationResponsePropertiesPtrOutput) Elem() IpConfigurationResponsePropertiesOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) IpConfigurationResponseProperties {
		if v != nil {
			return *v
		}
		var ret IpConfigurationResponseProperties
		return ret
	}).(IpConfigurationResponsePropertiesOutput)
}

func (o IpConfigurationResponsePropertiesPtrOutput) PrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrefixLength
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesPtrOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAddress
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesPtrOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o IpConfigurationResponsePropertiesPtrOutput) Subnet() IpConfigurationResponseSubnetPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseProperties) *IpConfigurationResponseSubnet {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(IpConfigurationResponseSubnetPtrOutput)
}

type IpConfigurationResponseSubnet struct {
	Id *string `pulumi:"id"`
}

type IpConfigurationResponseSubnetOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponseSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationResponseSubnet)(nil)).Elem()
}

func (o IpConfigurationResponseSubnetOutput) ToIpConfigurationResponseSubnetOutput() IpConfigurationResponseSubnetOutput {
	return o
}

func (o IpConfigurationResponseSubnetOutput) ToIpConfigurationResponseSubnetOutputWithContext(ctx context.Context) IpConfigurationResponseSubnetOutput {
	return o
}

func (o IpConfigurationResponseSubnetOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationResponseSubnet) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type IpConfigurationResponseSubnetPtrOutput struct{ *pulumi.OutputState }

func (IpConfigurationResponseSubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationResponseSubnet)(nil)).Elem()
}

func (o IpConfigurationResponseSubnetPtrOutput) ToIpConfigurationResponseSubnetPtrOutput() IpConfigurationResponseSubnetPtrOutput {
	return o
}

func (o IpConfigurationResponseSubnetPtrOutput) ToIpConfigurationResponseSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationResponseSubnetPtrOutput {
	return o
}

func (o IpConfigurationResponseSubnetPtrOutput) Elem() IpConfigurationResponseSubnetOutput {
	return o.ApplyT(func(v *IpConfigurationResponseSubnet) IpConfigurationResponseSubnet {
		if v != nil {
			return *v
		}
		var ret IpConfigurationResponseSubnet
		return ret
	}).(IpConfigurationResponseSubnetOutput)
}

func (o IpConfigurationResponseSubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationResponseSubnet) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type IpConfigurationSubnet struct {
	Id *string `pulumi:"id"`
}





type IpConfigurationSubnetInput interface {
	pulumi.Input

	ToIpConfigurationSubnetOutput() IpConfigurationSubnetOutput
	ToIpConfigurationSubnetOutputWithContext(context.Context) IpConfigurationSubnetOutput
}

type IpConfigurationSubnetArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (IpConfigurationSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationSubnet)(nil)).Elem()
}

func (i IpConfigurationSubnetArgs) ToIpConfigurationSubnetOutput() IpConfigurationSubnetOutput {
	return i.ToIpConfigurationSubnetOutputWithContext(context.Background())
}

func (i IpConfigurationSubnetArgs) ToIpConfigurationSubnetOutputWithContext(ctx context.Context) IpConfigurationSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationSubnetOutput)
}

func (i IpConfigurationSubnetArgs) ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput {
	return i.ToIpConfigurationSubnetPtrOutputWithContext(context.Background())
}

func (i IpConfigurationSubnetArgs) ToIpConfigurationSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationSubnetOutput).ToIpConfigurationSubnetPtrOutputWithContext(ctx)
}









type IpConfigurationSubnetPtrInput interface {
	pulumi.Input

	ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput
	ToIpConfigurationSubnetPtrOutputWithContext(context.Context) IpConfigurationSubnetPtrOutput
}

type ipConfigurationSubnetPtrType IpConfigurationSubnetArgs

func IpConfigurationSubnetPtr(v *IpConfigurationSubnetArgs) IpConfigurationSubnetPtrInput {
	return (*ipConfigurationSubnetPtrType)(v)
}

func (*ipConfigurationSubnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationSubnet)(nil)).Elem()
}

func (i *ipConfigurationSubnetPtrType) ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput {
	return i.ToIpConfigurationSubnetPtrOutputWithContext(context.Background())
}

func (i *ipConfigurationSubnetPtrType) ToIpConfigurationSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpConfigurationSubnetPtrOutput)
}

type IpConfigurationSubnetOutput struct{ *pulumi.OutputState }

func (IpConfigurationSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpConfigurationSubnet)(nil)).Elem()
}

func (o IpConfigurationSubnetOutput) ToIpConfigurationSubnetOutput() IpConfigurationSubnetOutput {
	return o
}

func (o IpConfigurationSubnetOutput) ToIpConfigurationSubnetOutputWithContext(ctx context.Context) IpConfigurationSubnetOutput {
	return o
}

func (o IpConfigurationSubnetOutput) ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput {
	return o.ToIpConfigurationSubnetPtrOutputWithContext(context.Background())
}

func (o IpConfigurationSubnetOutput) ToIpConfigurationSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationSubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpConfigurationSubnet) *IpConfigurationSubnet {
		return &v
	}).(IpConfigurationSubnetPtrOutput)
}

func (o IpConfigurationSubnetOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpConfigurationSubnet) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type IpConfigurationSubnetPtrOutput struct{ *pulumi.OutputState }

func (IpConfigurationSubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpConfigurationSubnet)(nil)).Elem()
}

func (o IpConfigurationSubnetPtrOutput) ToIpConfigurationSubnetPtrOutput() IpConfigurationSubnetPtrOutput {
	return o
}

func (o IpConfigurationSubnetPtrOutput) ToIpConfigurationSubnetPtrOutputWithContext(ctx context.Context) IpConfigurationSubnetPtrOutput {
	return o
}

func (o IpConfigurationSubnetPtrOutput) Elem() IpConfigurationSubnetOutput {
	return o.ApplyT(func(v *IpConfigurationSubnet) IpConfigurationSubnet {
		if v != nil {
			return *v
		}
		var ret IpConfigurationSubnet
		return ret
	}).(IpConfigurationSubnetOutput)
}

func (o IpConfigurationSubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpConfigurationSubnet) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceStatusResponse struct {
	ErrorCode    *string `pulumi:"errorCode"`
	ErrorMessage *string `pulumi:"errorMessage"`
}

type NetworkInterfaceStatusResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceStatusResponse)(nil)).Elem()
}

func (o NetworkInterfaceStatusResponseOutput) ToNetworkInterfaceStatusResponseOutput() NetworkInterfaceStatusResponseOutput {
	return o
}

func (o NetworkInterfaceStatusResponseOutput) ToNetworkInterfaceStatusResponseOutputWithContext(ctx context.Context) NetworkInterfaceStatusResponseOutput {
	return o
}

func (o NetworkInterfaceStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type VirtualHardDiskStatusResponse struct {
	ErrorCode    *string `pulumi:"errorCode"`
	ErrorMessage *string `pulumi:"errorMessage"`
}

type VirtualHardDiskStatusResponseOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDiskStatusResponse)(nil)).Elem()
}

func (o VirtualHardDiskStatusResponseOutput) ToVirtualHardDiskStatusResponseOutput() VirtualHardDiskStatusResponseOutput {
	return o
}

func (o VirtualHardDiskStatusResponseOutput) ToVirtualHardDiskStatusResponseOutputWithContext(ctx context.Context) VirtualHardDiskStatusResponseOutput {
	return o
}

func (o VirtualHardDiskStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o VirtualHardDiskStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

type VirtualMachineStatusResponse struct {
	ErrorCode    *string `pulumi:"errorCode"`
	ErrorMessage *string `pulumi:"errorMessage"`
	PowerState   *string `pulumi:"powerState"`
}

type VirtualMachineStatusResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineStatusResponse)(nil)).Elem()
}

func (o VirtualMachineStatusResponseOutput) ToVirtualMachineStatusResponseOutput() VirtualMachineStatusResponseOutput {
	return o
}

func (o VirtualMachineStatusResponseOutput) ToVirtualMachineStatusResponseOutputWithContext(ctx context.Context) VirtualMachineStatusResponseOutput {
	return o
}

func (o VirtualMachineStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineStatusResponseOutput) PowerState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineStatusResponse) *string { return v.PowerState }).(pulumi.StringPtrOutput)
}

type VirtualNetworkStatusResponse struct {
	ErrorCode    *string `pulumi:"errorCode"`
	ErrorMessage *string `pulumi:"errorMessage"`
}

type VirtualNetworkStatusResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkStatusResponse)(nil)).Elem()
}

func (o VirtualNetworkStatusResponseOutput) ToVirtualNetworkStatusResponseOutput() VirtualNetworkStatusResponseOutput {
	return o
}

func (o VirtualNetworkStatusResponseOutput) ToVirtualNetworkStatusResponseOutputWithContext(ctx context.Context) VirtualNetworkStatusResponseOutput {
	return o
}

func (o VirtualNetworkStatusResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkStatusResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkStatusResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkStatusResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesDataDisks struct {
	Name *string `pulumi:"name"`
}





type VirtualmachinesPropertiesDataDisksInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesDataDisksOutput() VirtualmachinesPropertiesDataDisksOutput
	ToVirtualmachinesPropertiesDataDisksOutputWithContext(context.Context) VirtualmachinesPropertiesDataDisksOutput
}

type VirtualmachinesPropertiesDataDisksArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (VirtualmachinesPropertiesDataDisksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesDataDisks)(nil)).Elem()
}

func (i VirtualmachinesPropertiesDataDisksArgs) ToVirtualmachinesPropertiesDataDisksOutput() VirtualmachinesPropertiesDataDisksOutput {
	return i.ToVirtualmachinesPropertiesDataDisksOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesDataDisksArgs) ToVirtualmachinesPropertiesDataDisksOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDataDisksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDataDisksOutput)
}





type VirtualmachinesPropertiesDataDisksArrayInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesDataDisksArrayOutput() VirtualmachinesPropertiesDataDisksArrayOutput
	ToVirtualmachinesPropertiesDataDisksArrayOutputWithContext(context.Context) VirtualmachinesPropertiesDataDisksArrayOutput
}

type VirtualmachinesPropertiesDataDisksArray []VirtualmachinesPropertiesDataDisksInput

func (VirtualmachinesPropertiesDataDisksArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesDataDisks)(nil)).Elem()
}

func (i VirtualmachinesPropertiesDataDisksArray) ToVirtualmachinesPropertiesDataDisksArrayOutput() VirtualmachinesPropertiesDataDisksArrayOutput {
	return i.ToVirtualmachinesPropertiesDataDisksArrayOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesDataDisksArray) ToVirtualmachinesPropertiesDataDisksArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDataDisksArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDataDisksArrayOutput)
}

type VirtualmachinesPropertiesDataDisksOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesDataDisksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesDataDisks)(nil)).Elem()
}

func (o VirtualmachinesPropertiesDataDisksOutput) ToVirtualmachinesPropertiesDataDisksOutput() VirtualmachinesPropertiesDataDisksOutput {
	return o
}

func (o VirtualmachinesPropertiesDataDisksOutput) ToVirtualmachinesPropertiesDataDisksOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDataDisksOutput {
	return o
}

func (o VirtualmachinesPropertiesDataDisksOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesDataDisks) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesDataDisksArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesDataDisksArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesDataDisks)(nil)).Elem()
}

func (o VirtualmachinesPropertiesDataDisksArrayOutput) ToVirtualmachinesPropertiesDataDisksArrayOutput() VirtualmachinesPropertiesDataDisksArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesDataDisksArrayOutput) ToVirtualmachinesPropertiesDataDisksArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDataDisksArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesDataDisksArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesDataDisksOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesDataDisks {
		return vs[0].([]VirtualmachinesPropertiesDataDisks)[vs[1].(int)]
	}).(VirtualmachinesPropertiesDataDisksOutput)
}

type VirtualmachinesPropertiesDynamicMemoryConfig struct {
	MaximumMemoryGB    *float64 `pulumi:"maximumMemoryGB"`
	MinimumMemoryGB    *float64 `pulumi:"minimumMemoryGB"`
	TargetMemoryBuffer *int     `pulumi:"targetMemoryBuffer"`
}





type VirtualmachinesPropertiesDynamicMemoryConfigInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesDynamicMemoryConfigOutput() VirtualmachinesPropertiesDynamicMemoryConfigOutput
	ToVirtualmachinesPropertiesDynamicMemoryConfigOutputWithContext(context.Context) VirtualmachinesPropertiesDynamicMemoryConfigOutput
}

type VirtualmachinesPropertiesDynamicMemoryConfigArgs struct {
	MaximumMemoryGB    pulumi.Float64PtrInput `pulumi:"maximumMemoryGB"`
	MinimumMemoryGB    pulumi.Float64PtrInput `pulumi:"minimumMemoryGB"`
	TargetMemoryBuffer pulumi.IntPtrInput     `pulumi:"targetMemoryBuffer"`
}

func (VirtualmachinesPropertiesDynamicMemoryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesDynamicMemoryConfig)(nil)).Elem()
}

func (i VirtualmachinesPropertiesDynamicMemoryConfigArgs) ToVirtualmachinesPropertiesDynamicMemoryConfigOutput() VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return i.ToVirtualmachinesPropertiesDynamicMemoryConfigOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesDynamicMemoryConfigArgs) ToVirtualmachinesPropertiesDynamicMemoryConfigOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDynamicMemoryConfigOutput)
}

func (i VirtualmachinesPropertiesDynamicMemoryConfigArgs) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return i.ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesDynamicMemoryConfigArgs) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDynamicMemoryConfigOutput).ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesDynamicMemoryConfigPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput
	ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput
}

type virtualmachinesPropertiesDynamicMemoryConfigPtrType VirtualmachinesPropertiesDynamicMemoryConfigArgs

func VirtualmachinesPropertiesDynamicMemoryConfigPtr(v *VirtualmachinesPropertiesDynamicMemoryConfigArgs) VirtualmachinesPropertiesDynamicMemoryConfigPtrInput {
	return (*virtualmachinesPropertiesDynamicMemoryConfigPtrType)(v)
}

func (*virtualmachinesPropertiesDynamicMemoryConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesDynamicMemoryConfig)(nil)).Elem()
}

func (i *virtualmachinesPropertiesDynamicMemoryConfigPtrType) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return i.ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesDynamicMemoryConfigPtrType) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput)
}

type VirtualmachinesPropertiesDynamicMemoryConfigOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesDynamicMemoryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesDynamicMemoryConfig)(nil)).Elem()
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigOutput() VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return o
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return o
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o.ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesDynamicMemoryConfig) *VirtualmachinesPropertiesDynamicMemoryConfig {
		return &v
	}).(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) MaximumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesDynamicMemoryConfig) *float64 { return v.MaximumMemoryGB }).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) MinimumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesDynamicMemoryConfig) *float64 { return v.MinimumMemoryGB }).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigOutput) TargetMemoryBuffer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesDynamicMemoryConfig) *int { return v.TargetMemoryBuffer }).(pulumi.IntPtrOutput)
}

type VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesDynamicMemoryConfig)(nil)).Elem()
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) ToVirtualmachinesPropertiesDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) Elem() VirtualmachinesPropertiesDynamicMemoryConfigOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesDynamicMemoryConfig) VirtualmachinesPropertiesDynamicMemoryConfig {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesDynamicMemoryConfig
		return ret
	}).(VirtualmachinesPropertiesDynamicMemoryConfigOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) MaximumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesDynamicMemoryConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.MaximumMemoryGB
	}).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) MinimumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesDynamicMemoryConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumMemoryGB
	}).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput) TargetMemoryBuffer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesDynamicMemoryConfig) *int {
		if v == nil {
			return nil
		}
		return v.TargetMemoryBuffer
	}).(pulumi.IntPtrOutput)
}

type VirtualmachinesPropertiesHardwareProfile struct {
	DynamicMemoryConfig *VirtualmachinesPropertiesDynamicMemoryConfig `pulumi:"dynamicMemoryConfig"`
	MemoryGB            *int                                          `pulumi:"memoryGB"`
	Processors          *int                                          `pulumi:"processors"`
	VmSize              *string                                       `pulumi:"vmSize"`
}





type VirtualmachinesPropertiesHardwareProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesHardwareProfileOutput() VirtualmachinesPropertiesHardwareProfileOutput
	ToVirtualmachinesPropertiesHardwareProfileOutputWithContext(context.Context) VirtualmachinesPropertiesHardwareProfileOutput
}

type VirtualmachinesPropertiesHardwareProfileArgs struct {
	DynamicMemoryConfig VirtualmachinesPropertiesDynamicMemoryConfigPtrInput `pulumi:"dynamicMemoryConfig"`
	MemoryGB            pulumi.IntPtrInput                                   `pulumi:"memoryGB"`
	Processors          pulumi.IntPtrInput                                   `pulumi:"processors"`
	VmSize              pulumi.StringPtrInput                                `pulumi:"vmSize"`
}

func (VirtualmachinesPropertiesHardwareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesHardwareProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesHardwareProfileArgs) ToVirtualmachinesPropertiesHardwareProfileOutput() VirtualmachinesPropertiesHardwareProfileOutput {
	return i.ToVirtualmachinesPropertiesHardwareProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesHardwareProfileArgs) ToVirtualmachinesPropertiesHardwareProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesHardwareProfileOutput)
}

func (i VirtualmachinesPropertiesHardwareProfileArgs) ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesHardwareProfileArgs) ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesHardwareProfileOutput).ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesHardwareProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput
	ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput
}

type virtualmachinesPropertiesHardwareProfilePtrType VirtualmachinesPropertiesHardwareProfileArgs

func VirtualmachinesPropertiesHardwareProfilePtr(v *VirtualmachinesPropertiesHardwareProfileArgs) VirtualmachinesPropertiesHardwareProfilePtrInput {
	return (*virtualmachinesPropertiesHardwareProfilePtrType)(v)
}

func (*virtualmachinesPropertiesHardwareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesHardwareProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesHardwareProfilePtrType) ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesHardwareProfilePtrType) ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesHardwareProfilePtrOutput)
}

type VirtualmachinesPropertiesHardwareProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesHardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesHardwareProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) ToVirtualmachinesPropertiesHardwareProfileOutput() VirtualmachinesPropertiesHardwareProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) ToVirtualmachinesPropertiesHardwareProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesHardwareProfile) *VirtualmachinesPropertiesHardwareProfile {
		return &v
	}).(VirtualmachinesPropertiesHardwareProfilePtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) DynamicMemoryConfig() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesHardwareProfile) *VirtualmachinesPropertiesDynamicMemoryConfig {
		return v.DynamicMemoryConfig
	}).(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) MemoryGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesHardwareProfile) *int { return v.MemoryGB }).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) Processors() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesHardwareProfile) *int { return v.Processors }).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesHardwareProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesHardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesHardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesHardwareProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) ToVirtualmachinesPropertiesHardwareProfilePtrOutput() VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) ToVirtualmachinesPropertiesHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesHardwareProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) Elem() VirtualmachinesPropertiesHardwareProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) VirtualmachinesPropertiesHardwareProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesHardwareProfile
		return ret
	}).(VirtualmachinesPropertiesHardwareProfileOutput)
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) DynamicMemoryConfig() VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) *VirtualmachinesPropertiesDynamicMemoryConfig {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryConfig
	}).(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) MemoryGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.MemoryGB
	}).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) Processors() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.Processors
	}).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesHardwareProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesHardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesImageReference struct {
	Name *string `pulumi:"name"`
}





type VirtualmachinesPropertiesImageReferenceInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesImageReferenceOutput() VirtualmachinesPropertiesImageReferenceOutput
	ToVirtualmachinesPropertiesImageReferenceOutputWithContext(context.Context) VirtualmachinesPropertiesImageReferenceOutput
}

type VirtualmachinesPropertiesImageReferenceArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (VirtualmachinesPropertiesImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesImageReference)(nil)).Elem()
}

func (i VirtualmachinesPropertiesImageReferenceArgs) ToVirtualmachinesPropertiesImageReferenceOutput() VirtualmachinesPropertiesImageReferenceOutput {
	return i.ToVirtualmachinesPropertiesImageReferenceOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesImageReferenceArgs) ToVirtualmachinesPropertiesImageReferenceOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesImageReferenceOutput)
}

func (i VirtualmachinesPropertiesImageReferenceArgs) ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput {
	return i.ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesImageReferenceArgs) ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesImageReferenceOutput).ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesImageReferencePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput
	ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(context.Context) VirtualmachinesPropertiesImageReferencePtrOutput
}

type virtualmachinesPropertiesImageReferencePtrType VirtualmachinesPropertiesImageReferenceArgs

func VirtualmachinesPropertiesImageReferencePtr(v *VirtualmachinesPropertiesImageReferenceArgs) VirtualmachinesPropertiesImageReferencePtrInput {
	return (*virtualmachinesPropertiesImageReferencePtrType)(v)
}

func (*virtualmachinesPropertiesImageReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesImageReference)(nil)).Elem()
}

func (i *virtualmachinesPropertiesImageReferencePtrType) ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput {
	return i.ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesImageReferencePtrType) ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesImageReferencePtrOutput)
}

type VirtualmachinesPropertiesImageReferenceOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesImageReference)(nil)).Elem()
}

func (o VirtualmachinesPropertiesImageReferenceOutput) ToVirtualmachinesPropertiesImageReferenceOutput() VirtualmachinesPropertiesImageReferenceOutput {
	return o
}

func (o VirtualmachinesPropertiesImageReferenceOutput) ToVirtualmachinesPropertiesImageReferenceOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferenceOutput {
	return o
}

func (o VirtualmachinesPropertiesImageReferenceOutput) ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput {
	return o.ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesImageReferenceOutput) ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesImageReference) *VirtualmachinesPropertiesImageReference {
		return &v
	}).(VirtualmachinesPropertiesImageReferencePtrOutput)
}

func (o VirtualmachinesPropertiesImageReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesImageReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesImageReferencePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesImageReference)(nil)).Elem()
}

func (o VirtualmachinesPropertiesImageReferencePtrOutput) ToVirtualmachinesPropertiesImageReferencePtrOutput() VirtualmachinesPropertiesImageReferencePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesImageReferencePtrOutput) ToVirtualmachinesPropertiesImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesImageReferencePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesImageReferencePtrOutput) Elem() VirtualmachinesPropertiesImageReferenceOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesImageReference) VirtualmachinesPropertiesImageReference {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesImageReference
		return ret
	}).(VirtualmachinesPropertiesImageReferenceOutput)
}

func (o VirtualmachinesPropertiesImageReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesLinuxConfiguration struct {
	DisablePasswordAuthentication *bool                         `pulumi:"disablePasswordAuthentication"`
	Ssh                           *VirtualmachinesPropertiesSsh `pulumi:"ssh"`
}





type VirtualmachinesPropertiesLinuxConfigurationInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesLinuxConfigurationOutput() VirtualmachinesPropertiesLinuxConfigurationOutput
	ToVirtualmachinesPropertiesLinuxConfigurationOutputWithContext(context.Context) VirtualmachinesPropertiesLinuxConfigurationOutput
}

type VirtualmachinesPropertiesLinuxConfigurationArgs struct {
	DisablePasswordAuthentication pulumi.BoolPtrInput                  `pulumi:"disablePasswordAuthentication"`
	Ssh                           VirtualmachinesPropertiesSshPtrInput `pulumi:"ssh"`
}

func (VirtualmachinesPropertiesLinuxConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesLinuxConfiguration)(nil)).Elem()
}

func (i VirtualmachinesPropertiesLinuxConfigurationArgs) ToVirtualmachinesPropertiesLinuxConfigurationOutput() VirtualmachinesPropertiesLinuxConfigurationOutput {
	return i.ToVirtualmachinesPropertiesLinuxConfigurationOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesLinuxConfigurationArgs) ToVirtualmachinesPropertiesLinuxConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesLinuxConfigurationOutput)
}

func (i VirtualmachinesPropertiesLinuxConfigurationArgs) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return i.ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesLinuxConfigurationArgs) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesLinuxConfigurationOutput).ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesLinuxConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput
	ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput
}

type virtualmachinesPropertiesLinuxConfigurationPtrType VirtualmachinesPropertiesLinuxConfigurationArgs

func VirtualmachinesPropertiesLinuxConfigurationPtr(v *VirtualmachinesPropertiesLinuxConfigurationArgs) VirtualmachinesPropertiesLinuxConfigurationPtrInput {
	return (*virtualmachinesPropertiesLinuxConfigurationPtrType)(v)
}

func (*virtualmachinesPropertiesLinuxConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesLinuxConfiguration)(nil)).Elem()
}

func (i *virtualmachinesPropertiesLinuxConfigurationPtrType) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return i.ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesLinuxConfigurationPtrType) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesLinuxConfigurationPtrOutput)
}

type VirtualmachinesPropertiesLinuxConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesLinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesLinuxConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) ToVirtualmachinesPropertiesLinuxConfigurationOutput() VirtualmachinesPropertiesLinuxConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) ToVirtualmachinesPropertiesLinuxConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o.ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesLinuxConfiguration) *VirtualmachinesPropertiesLinuxConfiguration {
		return &v
	}).(VirtualmachinesPropertiesLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesLinuxConfiguration) *bool { return v.DisablePasswordAuthentication }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationOutput) Ssh() VirtualmachinesPropertiesSshPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesLinuxConfiguration) *VirtualmachinesPropertiesSsh { return v.Ssh }).(VirtualmachinesPropertiesSshPtrOutput)
}

type VirtualmachinesPropertiesLinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesLinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesLinuxConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutput() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) ToVirtualmachinesPropertiesLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) Elem() VirtualmachinesPropertiesLinuxConfigurationOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesLinuxConfiguration) VirtualmachinesPropertiesLinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesLinuxConfiguration
		return ret
	}).(VirtualmachinesPropertiesLinuxConfigurationOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesLinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesLinuxConfigurationPtrOutput) Ssh() VirtualmachinesPropertiesSshPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesLinuxConfiguration) *VirtualmachinesPropertiesSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(VirtualmachinesPropertiesSshPtrOutput)
}

type VirtualmachinesPropertiesNetworkInterfaces struct {
	Id *string `pulumi:"id"`
}





type VirtualmachinesPropertiesNetworkInterfacesInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesNetworkInterfacesOutput() VirtualmachinesPropertiesNetworkInterfacesOutput
	ToVirtualmachinesPropertiesNetworkInterfacesOutputWithContext(context.Context) VirtualmachinesPropertiesNetworkInterfacesOutput
}

type VirtualmachinesPropertiesNetworkInterfacesArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualmachinesPropertiesNetworkInterfacesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesNetworkInterfaces)(nil)).Elem()
}

func (i VirtualmachinesPropertiesNetworkInterfacesArgs) ToVirtualmachinesPropertiesNetworkInterfacesOutput() VirtualmachinesPropertiesNetworkInterfacesOutput {
	return i.ToVirtualmachinesPropertiesNetworkInterfacesOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesNetworkInterfacesArgs) ToVirtualmachinesPropertiesNetworkInterfacesOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkInterfacesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkInterfacesOutput)
}





type VirtualmachinesPropertiesNetworkInterfacesArrayInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesNetworkInterfacesArrayOutput() VirtualmachinesPropertiesNetworkInterfacesArrayOutput
	ToVirtualmachinesPropertiesNetworkInterfacesArrayOutputWithContext(context.Context) VirtualmachinesPropertiesNetworkInterfacesArrayOutput
}

type VirtualmachinesPropertiesNetworkInterfacesArray []VirtualmachinesPropertiesNetworkInterfacesInput

func (VirtualmachinesPropertiesNetworkInterfacesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesNetworkInterfaces)(nil)).Elem()
}

func (i VirtualmachinesPropertiesNetworkInterfacesArray) ToVirtualmachinesPropertiesNetworkInterfacesArrayOutput() VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return i.ToVirtualmachinesPropertiesNetworkInterfacesArrayOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesNetworkInterfacesArray) ToVirtualmachinesPropertiesNetworkInterfacesArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesNetworkInterfacesOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesNetworkInterfacesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesNetworkInterfaces)(nil)).Elem()
}

func (o VirtualmachinesPropertiesNetworkInterfacesOutput) ToVirtualmachinesPropertiesNetworkInterfacesOutput() VirtualmachinesPropertiesNetworkInterfacesOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkInterfacesOutput) ToVirtualmachinesPropertiesNetworkInterfacesOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkInterfacesOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkInterfacesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesNetworkInterfaces) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesNetworkInterfacesArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesNetworkInterfacesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesNetworkInterfaces)(nil)).Elem()
}

func (o VirtualmachinesPropertiesNetworkInterfacesArrayOutput) ToVirtualmachinesPropertiesNetworkInterfacesArrayOutput() VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkInterfacesArrayOutput) ToVirtualmachinesPropertiesNetworkInterfacesArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkInterfacesArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesNetworkInterfacesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesNetworkInterfaces {
		return vs[0].([]VirtualmachinesPropertiesNetworkInterfaces)[vs[1].(int)]
	}).(VirtualmachinesPropertiesNetworkInterfacesOutput)
}

type VirtualmachinesPropertiesNetworkProfile struct {
	NetworkInterfaces []VirtualmachinesPropertiesNetworkInterfaces `pulumi:"networkInterfaces"`
}





type VirtualmachinesPropertiesNetworkProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesNetworkProfileOutput() VirtualmachinesPropertiesNetworkProfileOutput
	ToVirtualmachinesPropertiesNetworkProfileOutputWithContext(context.Context) VirtualmachinesPropertiesNetworkProfileOutput
}

type VirtualmachinesPropertiesNetworkProfileArgs struct {
	NetworkInterfaces VirtualmachinesPropertiesNetworkInterfacesArrayInput `pulumi:"networkInterfaces"`
}

func (VirtualmachinesPropertiesNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesNetworkProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesNetworkProfileArgs) ToVirtualmachinesPropertiesNetworkProfileOutput() VirtualmachinesPropertiesNetworkProfileOutput {
	return i.ToVirtualmachinesPropertiesNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesNetworkProfileArgs) ToVirtualmachinesPropertiesNetworkProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkProfileOutput)
}

func (i VirtualmachinesPropertiesNetworkProfileArgs) ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesNetworkProfileArgs) ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkProfileOutput).ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesNetworkProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput
	ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput
}

type virtualmachinesPropertiesNetworkProfilePtrType VirtualmachinesPropertiesNetworkProfileArgs

func VirtualmachinesPropertiesNetworkProfilePtr(v *VirtualmachinesPropertiesNetworkProfileArgs) VirtualmachinesPropertiesNetworkProfilePtrInput {
	return (*virtualmachinesPropertiesNetworkProfilePtrType)(v)
}

func (*virtualmachinesPropertiesNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesNetworkProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesNetworkProfilePtrType) ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesNetworkProfilePtrType) ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesNetworkProfilePtrOutput)
}

type VirtualmachinesPropertiesNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesNetworkProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) ToVirtualmachinesPropertiesNetworkProfileOutput() VirtualmachinesPropertiesNetworkProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) ToVirtualmachinesPropertiesNetworkProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesNetworkProfile) *VirtualmachinesPropertiesNetworkProfile {
		return &v
	}).(VirtualmachinesPropertiesNetworkProfilePtrOutput)
}

func (o VirtualmachinesPropertiesNetworkProfileOutput) NetworkInterfaces() VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesNetworkProfile) []VirtualmachinesPropertiesNetworkInterfaces {
		return v.NetworkInterfaces
	}).(VirtualmachinesPropertiesNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesNetworkProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesNetworkProfilePtrOutput) ToVirtualmachinesPropertiesNetworkProfilePtrOutput() VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkProfilePtrOutput) ToVirtualmachinesPropertiesNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesNetworkProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesNetworkProfilePtrOutput) Elem() VirtualmachinesPropertiesNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesNetworkProfile) VirtualmachinesPropertiesNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesNetworkProfile
		return ret
	}).(VirtualmachinesPropertiesNetworkProfileOutput)
}

func (o VirtualmachinesPropertiesNetworkProfilePtrOutput) NetworkInterfaces() VirtualmachinesPropertiesNetworkInterfacesArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesNetworkProfile) []VirtualmachinesPropertiesNetworkInterfaces {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(VirtualmachinesPropertiesNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesOsProfile struct {
	AdminPassword        *string                                        `pulumi:"adminPassword"`
	AdminUsername        *string                                        `pulumi:"adminUsername"`
	ComputerName         *string                                        `pulumi:"computerName"`
	LinuxConfiguration   *VirtualmachinesPropertiesLinuxConfiguration   `pulumi:"linuxConfiguration"`
	OsType               *string                                        `pulumi:"osType"`
	WindowsConfiguration *VirtualmachinesPropertiesWindowsConfiguration `pulumi:"windowsConfiguration"`
}





type VirtualmachinesPropertiesOsProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesOsProfileOutput() VirtualmachinesPropertiesOsProfileOutput
	ToVirtualmachinesPropertiesOsProfileOutputWithContext(context.Context) VirtualmachinesPropertiesOsProfileOutput
}

type VirtualmachinesPropertiesOsProfileArgs struct {
	AdminPassword        pulumi.StringPtrInput                                 `pulumi:"adminPassword"`
	AdminUsername        pulumi.StringPtrInput                                 `pulumi:"adminUsername"`
	ComputerName         pulumi.StringPtrInput                                 `pulumi:"computerName"`
	LinuxConfiguration   VirtualmachinesPropertiesLinuxConfigurationPtrInput   `pulumi:"linuxConfiguration"`
	OsType               pulumi.StringPtrInput                                 `pulumi:"osType"`
	WindowsConfiguration VirtualmachinesPropertiesWindowsConfigurationPtrInput `pulumi:"windowsConfiguration"`
}

func (VirtualmachinesPropertiesOsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesOsProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesOsProfileArgs) ToVirtualmachinesPropertiesOsProfileOutput() VirtualmachinesPropertiesOsProfileOutput {
	return i.ToVirtualmachinesPropertiesOsProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesOsProfileArgs) ToVirtualmachinesPropertiesOsProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesOsProfileOutput)
}

func (i VirtualmachinesPropertiesOsProfileArgs) ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesOsProfileArgs) ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesOsProfileOutput).ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesOsProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput
	ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesOsProfilePtrOutput
}

type virtualmachinesPropertiesOsProfilePtrType VirtualmachinesPropertiesOsProfileArgs

func VirtualmachinesPropertiesOsProfilePtr(v *VirtualmachinesPropertiesOsProfileArgs) VirtualmachinesPropertiesOsProfilePtrInput {
	return (*virtualmachinesPropertiesOsProfilePtrType)(v)
}

func (*virtualmachinesPropertiesOsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesOsProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesOsProfilePtrType) ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesOsProfilePtrType) ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesOsProfilePtrOutput)
}

type VirtualmachinesPropertiesOsProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesOsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesOsProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesOsProfileOutput) ToVirtualmachinesPropertiesOsProfileOutput() VirtualmachinesPropertiesOsProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesOsProfileOutput) ToVirtualmachinesPropertiesOsProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesOsProfileOutput) ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesOsProfileOutput) ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesOsProfile {
		return &v
	}).(VirtualmachinesPropertiesOsProfilePtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) LinuxConfiguration() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesLinuxConfiguration {
		return v.LinuxConfiguration
	}).(VirtualmachinesPropertiesLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfileOutput) WindowsConfiguration() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesWindowsConfiguration {
		return v.WindowsConfiguration
	}).(VirtualmachinesPropertiesWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesOsProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesOsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesOsProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) ToVirtualmachinesPropertiesOsProfilePtrOutput() VirtualmachinesPropertiesOsProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) ToVirtualmachinesPropertiesOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesOsProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) Elem() VirtualmachinesPropertiesOsProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) VirtualmachinesPropertiesOsProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesOsProfile
		return ret
	}).(VirtualmachinesPropertiesOsProfileOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) LinuxConfiguration() VirtualmachinesPropertiesLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesLinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(VirtualmachinesPropertiesLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesOsProfilePtrOutput) WindowsConfiguration() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesOsProfile) *VirtualmachinesPropertiesWindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(VirtualmachinesPropertiesWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesPublicKeys struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}





type VirtualmachinesPropertiesPublicKeysInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesPublicKeysOutput() VirtualmachinesPropertiesPublicKeysOutput
	ToVirtualmachinesPropertiesPublicKeysOutputWithContext(context.Context) VirtualmachinesPropertiesPublicKeysOutput
}

type VirtualmachinesPropertiesPublicKeysArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (VirtualmachinesPropertiesPublicKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesPublicKeys)(nil)).Elem()
}

func (i VirtualmachinesPropertiesPublicKeysArgs) ToVirtualmachinesPropertiesPublicKeysOutput() VirtualmachinesPropertiesPublicKeysOutput {
	return i.ToVirtualmachinesPropertiesPublicKeysOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesPublicKeysArgs) ToVirtualmachinesPropertiesPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesPublicKeysOutput)
}





type VirtualmachinesPropertiesPublicKeysArrayInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysArrayOutput
	ToVirtualmachinesPropertiesPublicKeysArrayOutputWithContext(context.Context) VirtualmachinesPropertiesPublicKeysArrayOutput
}

type VirtualmachinesPropertiesPublicKeysArray []VirtualmachinesPropertiesPublicKeysInput

func (VirtualmachinesPropertiesPublicKeysArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesPublicKeys)(nil)).Elem()
}

func (i VirtualmachinesPropertiesPublicKeysArray) ToVirtualmachinesPropertiesPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysArrayOutput {
	return i.ToVirtualmachinesPropertiesPublicKeysArrayOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesPublicKeysArray) ToVirtualmachinesPropertiesPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesPublicKeysOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesPublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesPublicKeysOutput) ToVirtualmachinesPropertiesPublicKeysOutput() VirtualmachinesPropertiesPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysOutput) ToVirtualmachinesPropertiesPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesPublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesPublicKeysOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesPublicKeys) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesPublicKeysArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesPublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesPublicKeysArrayOutput) ToVirtualmachinesPropertiesPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysArrayOutput) ToVirtualmachinesPropertiesPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesPublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesPublicKeys {
		return vs[0].([]VirtualmachinesPropertiesPublicKeys)[vs[1].(int)]
	}).(VirtualmachinesPropertiesPublicKeysOutput)
}

type VirtualmachinesPropertiesPublicKeysPublicKeys struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}





type VirtualmachinesPropertiesPublicKeysPublicKeysInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesPublicKeysPublicKeysOutput() VirtualmachinesPropertiesPublicKeysPublicKeysOutput
	ToVirtualmachinesPropertiesPublicKeysPublicKeysOutputWithContext(context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysOutput
}

type VirtualmachinesPropertiesPublicKeysPublicKeysArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (VirtualmachinesPropertiesPublicKeysPublicKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesPublicKeysPublicKeys)(nil)).Elem()
}

func (i VirtualmachinesPropertiesPublicKeysPublicKeysArgs) ToVirtualmachinesPropertiesPublicKeysPublicKeysOutput() VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return i.ToVirtualmachinesPropertiesPublicKeysPublicKeysOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesPublicKeysPublicKeysArgs) ToVirtualmachinesPropertiesPublicKeysPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesPublicKeysPublicKeysOutput)
}





type VirtualmachinesPropertiesPublicKeysPublicKeysArrayInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput
	ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutputWithContext(context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput
}

type VirtualmachinesPropertiesPublicKeysPublicKeysArray []VirtualmachinesPropertiesPublicKeysPublicKeysInput

func (VirtualmachinesPropertiesPublicKeysPublicKeysArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesPublicKeysPublicKeys)(nil)).Elem()
}

func (i VirtualmachinesPropertiesPublicKeysPublicKeysArray) ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return i.ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesPublicKeysPublicKeysArray) ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesPublicKeysPublicKeysOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesPublicKeysPublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesPublicKeysPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysOutput) ToVirtualmachinesPropertiesPublicKeysPublicKeysOutput() VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysOutput) ToVirtualmachinesPropertiesPublicKeysPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesPublicKeysPublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesPublicKeysPublicKeys) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesPublicKeysPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput) ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput) ToVirtualmachinesPropertiesPublicKeysPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesPublicKeysPublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesPublicKeysPublicKeys {
		return vs[0].([]VirtualmachinesPropertiesPublicKeysPublicKeys)[vs[1].(int)]
	}).(VirtualmachinesPropertiesPublicKeysPublicKeysOutput)
}

type VirtualmachinesPropertiesResponseDataDisks struct {
	Name *string `pulumi:"name"`
}

type VirtualmachinesPropertiesResponseDataDisksOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseDataDisksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseDataDisks)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseDataDisksOutput) ToVirtualmachinesPropertiesResponseDataDisksOutput() VirtualmachinesPropertiesResponseDataDisksOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDataDisksOutput) ToVirtualmachinesPropertiesResponseDataDisksOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseDataDisksOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDataDisksOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseDataDisks) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseDataDisksArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseDataDisksArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesResponseDataDisks)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseDataDisksArrayOutput) ToVirtualmachinesPropertiesResponseDataDisksArrayOutput() VirtualmachinesPropertiesResponseDataDisksArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDataDisksArrayOutput) ToVirtualmachinesPropertiesResponseDataDisksArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseDataDisksArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDataDisksArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesResponseDataDisksOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesResponseDataDisks {
		return vs[0].([]VirtualmachinesPropertiesResponseDataDisks)[vs[1].(int)]
	}).(VirtualmachinesPropertiesResponseDataDisksOutput)
}

type VirtualmachinesPropertiesResponseDynamicMemoryConfig struct {
	MaximumMemoryGB    *float64 `pulumi:"maximumMemoryGB"`
	MinimumMemoryGB    *float64 `pulumi:"minimumMemoryGB"`
	TargetMemoryBuffer *int     `pulumi:"targetMemoryBuffer"`
}

type VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseDynamicMemoryConfig)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesResponseDynamicMemoryConfigOutput() VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) ToVirtualmachinesPropertiesResponseDynamicMemoryConfigOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) MaximumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseDynamicMemoryConfig) *float64 { return v.MaximumMemoryGB }).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) MinimumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseDynamicMemoryConfig) *float64 { return v.MinimumMemoryGB }).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput) TargetMemoryBuffer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseDynamicMemoryConfig) *int { return v.TargetMemoryBuffer }).(pulumi.IntPtrOutput)
}

type VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseDynamicMemoryConfig)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) ToVirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput() VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) ToVirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) Elem() VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseDynamicMemoryConfig) VirtualmachinesPropertiesResponseDynamicMemoryConfig {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseDynamicMemoryConfig
		return ret
	}).(VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) MaximumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseDynamicMemoryConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.MaximumMemoryGB
	}).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) MinimumMemoryGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseDynamicMemoryConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumMemoryGB
	}).(pulumi.Float64PtrOutput)
}

func (o VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput) TargetMemoryBuffer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseDynamicMemoryConfig) *int {
		if v == nil {
			return nil
		}
		return v.TargetMemoryBuffer
	}).(pulumi.IntPtrOutput)
}

type VirtualmachinesPropertiesResponseHardwareProfile struct {
	DynamicMemoryConfig *VirtualmachinesPropertiesResponseDynamicMemoryConfig `pulumi:"dynamicMemoryConfig"`
	MemoryGB            *int                                                  `pulumi:"memoryGB"`
	Processors          *int                                                  `pulumi:"processors"`
	VmSize              *string                                               `pulumi:"vmSize"`
}

type VirtualmachinesPropertiesResponseHardwareProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseHardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseHardwareProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) ToVirtualmachinesPropertiesResponseHardwareProfileOutput() VirtualmachinesPropertiesResponseHardwareProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) ToVirtualmachinesPropertiesResponseHardwareProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseHardwareProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) DynamicMemoryConfig() VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseHardwareProfile) *VirtualmachinesPropertiesResponseDynamicMemoryConfig {
		return v.DynamicMemoryConfig
	}).(VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) MemoryGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseHardwareProfile) *int { return v.MemoryGB }).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) Processors() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseHardwareProfile) *int { return v.Processors }).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseHardwareProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseHardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseHardwareProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) ToVirtualmachinesPropertiesResponseHardwareProfilePtrOutput() VirtualmachinesPropertiesResponseHardwareProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) ToVirtualmachinesPropertiesResponseHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseHardwareProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseHardwareProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) VirtualmachinesPropertiesResponseHardwareProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseHardwareProfile
		return ret
	}).(VirtualmachinesPropertiesResponseHardwareProfileOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) DynamicMemoryConfig() VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) *VirtualmachinesPropertiesResponseDynamicMemoryConfig {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryConfig
	}).(VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) MemoryGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.MemoryGB
	}).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) Processors() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.Processors
	}).(pulumi.IntPtrOutput)
}

func (o VirtualmachinesPropertiesResponseHardwareProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseHardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseImageReference struct {
	Name *string `pulumi:"name"`
}

type VirtualmachinesPropertiesResponseImageReferenceOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseImageReference)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseImageReferenceOutput) ToVirtualmachinesPropertiesResponseImageReferenceOutput() VirtualmachinesPropertiesResponseImageReferenceOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseImageReferenceOutput) ToVirtualmachinesPropertiesResponseImageReferenceOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseImageReferenceOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseImageReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseImageReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseImageReferencePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseImageReference)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseImageReferencePtrOutput) ToVirtualmachinesPropertiesResponseImageReferencePtrOutput() VirtualmachinesPropertiesResponseImageReferencePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseImageReferencePtrOutput) ToVirtualmachinesPropertiesResponseImageReferencePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseImageReferencePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseImageReferencePtrOutput) Elem() VirtualmachinesPropertiesResponseImageReferenceOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseImageReference) VirtualmachinesPropertiesResponseImageReference {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseImageReference
		return ret
	}).(VirtualmachinesPropertiesResponseImageReferenceOutput)
}

func (o VirtualmachinesPropertiesResponseImageReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseLinuxConfiguration struct {
	DisablePasswordAuthentication *bool                                 `pulumi:"disablePasswordAuthentication"`
	Ssh                           *VirtualmachinesPropertiesResponseSsh `pulumi:"ssh"`
}

type VirtualmachinesPropertiesResponseLinuxConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseLinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseLinuxConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationOutput) ToVirtualmachinesPropertiesResponseLinuxConfigurationOutput() VirtualmachinesPropertiesResponseLinuxConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationOutput) ToVirtualmachinesPropertiesResponseLinuxConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseLinuxConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseLinuxConfiguration) *bool {
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationOutput) Ssh() VirtualmachinesPropertiesResponseSshPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseLinuxConfiguration) *VirtualmachinesPropertiesResponseSsh {
		return v.Ssh
	}).(VirtualmachinesPropertiesResponseSshPtrOutput)
}

type VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseLinuxConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) ToVirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput() VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) ToVirtualmachinesPropertiesResponseLinuxConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) Elem() VirtualmachinesPropertiesResponseLinuxConfigurationOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseLinuxConfiguration) VirtualmachinesPropertiesResponseLinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseLinuxConfiguration
		return ret
	}).(VirtualmachinesPropertiesResponseLinuxConfigurationOutput)
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseLinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput) Ssh() VirtualmachinesPropertiesResponseSshPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseLinuxConfiguration) *VirtualmachinesPropertiesResponseSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(VirtualmachinesPropertiesResponseSshPtrOutput)
}

type VirtualmachinesPropertiesResponseNetworkInterfaces struct {
	Id *string `pulumi:"id"`
}

type VirtualmachinesPropertiesResponseNetworkInterfacesOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseNetworkInterfacesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseNetworkInterfaces)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesOutput) ToVirtualmachinesPropertiesResponseNetworkInterfacesOutput() VirtualmachinesPropertiesResponseNetworkInterfacesOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesOutput) ToVirtualmachinesPropertiesResponseNetworkInterfacesOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseNetworkInterfacesOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseNetworkInterfaces) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesResponseNetworkInterfaces)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput) ToVirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput() VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput) ToVirtualmachinesPropertiesResponseNetworkInterfacesArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesResponseNetworkInterfacesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesResponseNetworkInterfaces {
		return vs[0].([]VirtualmachinesPropertiesResponseNetworkInterfaces)[vs[1].(int)]
	}).(VirtualmachinesPropertiesResponseNetworkInterfacesOutput)
}

type VirtualmachinesPropertiesResponseNetworkProfile struct {
	NetworkInterfaces []VirtualmachinesPropertiesResponseNetworkInterfaces `pulumi:"networkInterfaces"`
}

type VirtualmachinesPropertiesResponseNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseNetworkProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseNetworkProfileOutput) ToVirtualmachinesPropertiesResponseNetworkProfileOutput() VirtualmachinesPropertiesResponseNetworkProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkProfileOutput) ToVirtualmachinesPropertiesResponseNetworkProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseNetworkProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkProfileOutput) NetworkInterfaces() VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseNetworkProfile) []VirtualmachinesPropertiesResponseNetworkInterfaces {
		return v.NetworkInterfaces
	}).(VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesResponseNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseNetworkProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) ToVirtualmachinesPropertiesResponseNetworkProfilePtrOutput() VirtualmachinesPropertiesResponseNetworkProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) ToVirtualmachinesPropertiesResponseNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseNetworkProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseNetworkProfile) VirtualmachinesPropertiesResponseNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseNetworkProfile
		return ret
	}).(VirtualmachinesPropertiesResponseNetworkProfileOutput)
}

func (o VirtualmachinesPropertiesResponseNetworkProfilePtrOutput) NetworkInterfaces() VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseNetworkProfile) []VirtualmachinesPropertiesResponseNetworkInterfaces {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput)
}

type VirtualmachinesPropertiesResponseOsProfile struct {
	AdminUsername        *string                                                `pulumi:"adminUsername"`
	ComputerName         *string                                                `pulumi:"computerName"`
	LinuxConfiguration   *VirtualmachinesPropertiesResponseLinuxConfiguration   `pulumi:"linuxConfiguration"`
	OsType               *string                                                `pulumi:"osType"`
	WindowsConfiguration *VirtualmachinesPropertiesResponseWindowsConfiguration `pulumi:"windowsConfiguration"`
}

type VirtualmachinesPropertiesResponseOsProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseOsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseOsProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) ToVirtualmachinesPropertiesResponseOsProfileOutput() VirtualmachinesPropertiesResponseOsProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) ToVirtualmachinesPropertiesResponseOsProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseOsProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) LinuxConfiguration() VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *VirtualmachinesPropertiesResponseLinuxConfiguration {
		return v.LinuxConfiguration
	}).(VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfileOutput) WindowsConfiguration() VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseOsProfile) *VirtualmachinesPropertiesResponseWindowsConfiguration {
		return v.WindowsConfiguration
	}).(VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesResponseOsProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseOsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseOsProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) ToVirtualmachinesPropertiesResponseOsProfilePtrOutput() VirtualmachinesPropertiesResponseOsProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) ToVirtualmachinesPropertiesResponseOsProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseOsProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseOsProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) VirtualmachinesPropertiesResponseOsProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseOsProfile
		return ret
	}).(VirtualmachinesPropertiesResponseOsProfileOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) LinuxConfiguration() VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *VirtualmachinesPropertiesResponseLinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponseOsProfilePtrOutput) WindowsConfiguration() VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseOsProfile) *VirtualmachinesPropertiesResponseWindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesResponsePublicKeys struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}

type VirtualmachinesPropertiesResponsePublicKeysOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponsePublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponsePublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponsePublicKeysOutput) ToVirtualmachinesPropertiesResponsePublicKeysOutput() VirtualmachinesPropertiesResponsePublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysOutput) ToVirtualmachinesPropertiesResponsePublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponsePublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponsePublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponsePublicKeysOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponsePublicKeys) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponsePublicKeysArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponsePublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesResponsePublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponsePublicKeysArrayOutput) ToVirtualmachinesPropertiesResponsePublicKeysArrayOutput() VirtualmachinesPropertiesResponsePublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysArrayOutput) ToVirtualmachinesPropertiesResponsePublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponsePublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesResponsePublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesResponsePublicKeys {
		return vs[0].([]VirtualmachinesPropertiesResponsePublicKeys)[vs[1].(int)]
	}).(VirtualmachinesPropertiesResponsePublicKeysOutput)
}

type VirtualmachinesPropertiesResponsePublicKeysPublicKeys struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}

type VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponsePublicKeysPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) ToVirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput() VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) ToVirtualmachinesPropertiesResponsePublicKeysPublicKeysOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponsePublicKeysPublicKeys) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponsePublicKeysPublicKeys) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualmachinesPropertiesResponsePublicKeysPublicKeys)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput) ToVirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput() VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput) ToVirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput {
	return o
}

func (o VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput) Index(i pulumi.IntInput) VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualmachinesPropertiesResponsePublicKeysPublicKeys {
		return vs[0].([]VirtualmachinesPropertiesResponsePublicKeysPublicKeys)[vs[1].(int)]
	}).(VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput)
}

type VirtualmachinesPropertiesResponseSecurityProfile struct {
	EnableTPM *bool `pulumi:"enableTPM"`
}

type VirtualmachinesPropertiesResponseSecurityProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseSecurityProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSecurityProfileOutput) ToVirtualmachinesPropertiesResponseSecurityProfileOutput() VirtualmachinesPropertiesResponseSecurityProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSecurityProfileOutput) ToVirtualmachinesPropertiesResponseSecurityProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSecurityProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSecurityProfileOutput) EnableTPM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseSecurityProfile) *bool { return v.EnableTPM }).(pulumi.BoolPtrOutput)
}

type VirtualmachinesPropertiesResponseSecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseSecurityProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) ToVirtualmachinesPropertiesResponseSecurityProfilePtrOutput() VirtualmachinesPropertiesResponseSecurityProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) ToVirtualmachinesPropertiesResponseSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSecurityProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseSecurityProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSecurityProfile) VirtualmachinesPropertiesResponseSecurityProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseSecurityProfile
		return ret
	}).(VirtualmachinesPropertiesResponseSecurityProfileOutput)
}

func (o VirtualmachinesPropertiesResponseSecurityProfilePtrOutput) EnableTPM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSecurityProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnableTPM
	}).(pulumi.BoolPtrOutput)
}

type VirtualmachinesPropertiesResponseSsh struct {
	PublicKeys []VirtualmachinesPropertiesResponsePublicKeys `pulumi:"publicKeys"`
}

type VirtualmachinesPropertiesResponseSshOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSshOutput) ToVirtualmachinesPropertiesResponseSshOutput() VirtualmachinesPropertiesResponseSshOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshOutput) ToVirtualmachinesPropertiesResponseSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSshOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshOutput) PublicKeys() VirtualmachinesPropertiesResponsePublicKeysArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseSsh) []VirtualmachinesPropertiesResponsePublicKeys {
		return v.PublicKeys
	}).(VirtualmachinesPropertiesResponsePublicKeysArrayOutput)
}

type VirtualmachinesPropertiesResponseSshPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSshPtrOutput) ToVirtualmachinesPropertiesResponseSshPtrOutput() VirtualmachinesPropertiesResponseSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshPtrOutput) ToVirtualmachinesPropertiesResponseSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshPtrOutput) Elem() VirtualmachinesPropertiesResponseSshOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSsh) VirtualmachinesPropertiesResponseSsh {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseSsh
		return ret
	}).(VirtualmachinesPropertiesResponseSshOutput)
}

func (o VirtualmachinesPropertiesResponseSshPtrOutput) PublicKeys() VirtualmachinesPropertiesResponsePublicKeysArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSsh) []VirtualmachinesPropertiesResponsePublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(VirtualmachinesPropertiesResponsePublicKeysArrayOutput)
}

type VirtualmachinesPropertiesResponseSshSsh struct {
	PublicKeys []VirtualmachinesPropertiesResponsePublicKeysPublicKeys `pulumi:"publicKeys"`
}

type VirtualmachinesPropertiesResponseSshSshOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSshSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseSshSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSshSshOutput) ToVirtualmachinesPropertiesResponseSshSshOutput() VirtualmachinesPropertiesResponseSshSshOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshSshOutput) ToVirtualmachinesPropertiesResponseSshSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSshSshOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshSshOutput) PublicKeys() VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseSshSsh) []VirtualmachinesPropertiesResponsePublicKeysPublicKeys {
		return v.PublicKeys
	}).(VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesResponseSshSshPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseSshSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseSshSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseSshSshPtrOutput) ToVirtualmachinesPropertiesResponseSshSshPtrOutput() VirtualmachinesPropertiesResponseSshSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshSshPtrOutput) ToVirtualmachinesPropertiesResponseSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseSshSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseSshSshPtrOutput) Elem() VirtualmachinesPropertiesResponseSshSshOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSshSsh) VirtualmachinesPropertiesResponseSshSsh {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseSshSsh
		return ret
	}).(VirtualmachinesPropertiesResponseSshSshOutput)
}

func (o VirtualmachinesPropertiesResponseSshSshPtrOutput) PublicKeys() VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseSshSsh) []VirtualmachinesPropertiesResponsePublicKeysPublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesResponseStorageProfile struct {
	DataDisks      []VirtualmachinesPropertiesResponseDataDisks     `pulumi:"dataDisks"`
	ImageReference *VirtualmachinesPropertiesResponseImageReference `pulumi:"imageReference"`
}

type VirtualmachinesPropertiesResponseStorageProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseStorageProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) ToVirtualmachinesPropertiesResponseStorageProfileOutput() VirtualmachinesPropertiesResponseStorageProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) ToVirtualmachinesPropertiesResponseStorageProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseStorageProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) DataDisks() VirtualmachinesPropertiesResponseDataDisksArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseStorageProfile) []VirtualmachinesPropertiesResponseDataDisks {
		return v.DataDisks
	}).(VirtualmachinesPropertiesResponseDataDisksArrayOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfileOutput) ImageReference() VirtualmachinesPropertiesResponseImageReferencePtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseStorageProfile) *VirtualmachinesPropertiesResponseImageReference {
		return v.ImageReference
	}).(VirtualmachinesPropertiesResponseImageReferencePtrOutput)
}

type VirtualmachinesPropertiesResponseStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseStorageProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) ToVirtualmachinesPropertiesResponseStorageProfilePtrOutput() VirtualmachinesPropertiesResponseStorageProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) ToVirtualmachinesPropertiesResponseStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseStorageProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) Elem() VirtualmachinesPropertiesResponseStorageProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseStorageProfile) VirtualmachinesPropertiesResponseStorageProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseStorageProfile
		return ret
	}).(VirtualmachinesPropertiesResponseStorageProfileOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) DataDisks() VirtualmachinesPropertiesResponseDataDisksArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseStorageProfile) []VirtualmachinesPropertiesResponseDataDisks {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(VirtualmachinesPropertiesResponseDataDisksArrayOutput)
}

func (o VirtualmachinesPropertiesResponseStorageProfilePtrOutput) ImageReference() VirtualmachinesPropertiesResponseImageReferencePtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseStorageProfile) *VirtualmachinesPropertiesResponseImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(VirtualmachinesPropertiesResponseImageReferencePtrOutput)
}

type VirtualmachinesPropertiesResponseWindowsConfiguration struct {
	EnableAutomaticUpdates *bool                                    `pulumi:"enableAutomaticUpdates"`
	Ssh                    *VirtualmachinesPropertiesResponseSshSsh `pulumi:"ssh"`
	TimeZone               *string                                  `pulumi:"timeZone"`
}

type VirtualmachinesPropertiesResponseWindowsConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseWindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesResponseWindowsConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) ToVirtualmachinesPropertiesResponseWindowsConfigurationOutput() VirtualmachinesPropertiesResponseWindowsConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) ToVirtualmachinesPropertiesResponseWindowsConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseWindowsConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseWindowsConfiguration) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) Ssh() VirtualmachinesPropertiesResponseSshSshPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseWindowsConfiguration) *VirtualmachinesPropertiesResponseSshSsh {
		return v.Ssh
	}).(VirtualmachinesPropertiesResponseSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesResponseWindowsConfiguration) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesResponseWindowsConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) ToVirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput() VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) ToVirtualmachinesPropertiesResponseWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) Elem() VirtualmachinesPropertiesResponseWindowsConfigurationOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseWindowsConfiguration) VirtualmachinesPropertiesResponseWindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesResponseWindowsConfiguration
		return ret
	}).(VirtualmachinesPropertiesResponseWindowsConfigurationOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseWindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) Ssh() VirtualmachinesPropertiesResponseSshSshPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseWindowsConfiguration) *VirtualmachinesPropertiesResponseSshSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(VirtualmachinesPropertiesResponseSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesResponseWindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesSecurityProfile struct {
	EnableTPM *bool `pulumi:"enableTPM"`
}





type VirtualmachinesPropertiesSecurityProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSecurityProfileOutput() VirtualmachinesPropertiesSecurityProfileOutput
	ToVirtualmachinesPropertiesSecurityProfileOutputWithContext(context.Context) VirtualmachinesPropertiesSecurityProfileOutput
}

type VirtualmachinesPropertiesSecurityProfileArgs struct {
	EnableTPM pulumi.BoolPtrInput `pulumi:"enableTPM"`
}

func (VirtualmachinesPropertiesSecurityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSecurityProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesSecurityProfileArgs) ToVirtualmachinesPropertiesSecurityProfileOutput() VirtualmachinesPropertiesSecurityProfileOutput {
	return i.ToVirtualmachinesPropertiesSecurityProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSecurityProfileArgs) ToVirtualmachinesPropertiesSecurityProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSecurityProfileOutput)
}

func (i VirtualmachinesPropertiesSecurityProfileArgs) ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSecurityProfileArgs) ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSecurityProfileOutput).ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesSecurityProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput
	ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput
}

type virtualmachinesPropertiesSecurityProfilePtrType VirtualmachinesPropertiesSecurityProfileArgs

func VirtualmachinesPropertiesSecurityProfilePtr(v *VirtualmachinesPropertiesSecurityProfileArgs) VirtualmachinesPropertiesSecurityProfilePtrInput {
	return (*virtualmachinesPropertiesSecurityProfilePtrType)(v)
}

func (*virtualmachinesPropertiesSecurityProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSecurityProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesSecurityProfilePtrType) ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesSecurityProfilePtrType) ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSecurityProfilePtrOutput)
}

type VirtualmachinesPropertiesSecurityProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSecurityProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) ToVirtualmachinesPropertiesSecurityProfileOutput() VirtualmachinesPropertiesSecurityProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) ToVirtualmachinesPropertiesSecurityProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesSecurityProfile) *VirtualmachinesPropertiesSecurityProfile {
		return &v
	}).(VirtualmachinesPropertiesSecurityProfilePtrOutput)
}

func (o VirtualmachinesPropertiesSecurityProfileOutput) EnableTPM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesSecurityProfile) *bool { return v.EnableTPM }).(pulumi.BoolPtrOutput)
}

type VirtualmachinesPropertiesSecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSecurityProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSecurityProfilePtrOutput) ToVirtualmachinesPropertiesSecurityProfilePtrOutput() VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSecurityProfilePtrOutput) ToVirtualmachinesPropertiesSecurityProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSecurityProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSecurityProfilePtrOutput) Elem() VirtualmachinesPropertiesSecurityProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSecurityProfile) VirtualmachinesPropertiesSecurityProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesSecurityProfile
		return ret
	}).(VirtualmachinesPropertiesSecurityProfileOutput)
}

func (o VirtualmachinesPropertiesSecurityProfilePtrOutput) EnableTPM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSecurityProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnableTPM
	}).(pulumi.BoolPtrOutput)
}

type VirtualmachinesPropertiesSsh struct {
	PublicKeys []VirtualmachinesPropertiesPublicKeys `pulumi:"publicKeys"`
}





type VirtualmachinesPropertiesSshInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSshOutput() VirtualmachinesPropertiesSshOutput
	ToVirtualmachinesPropertiesSshOutputWithContext(context.Context) VirtualmachinesPropertiesSshOutput
}

type VirtualmachinesPropertiesSshArgs struct {
	PublicKeys VirtualmachinesPropertiesPublicKeysArrayInput `pulumi:"publicKeys"`
}

func (VirtualmachinesPropertiesSshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSsh)(nil)).Elem()
}

func (i VirtualmachinesPropertiesSshArgs) ToVirtualmachinesPropertiesSshOutput() VirtualmachinesPropertiesSshOutput {
	return i.ToVirtualmachinesPropertiesSshOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSshArgs) ToVirtualmachinesPropertiesSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshOutput)
}

func (i VirtualmachinesPropertiesSshArgs) ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput {
	return i.ToVirtualmachinesPropertiesSshPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSshArgs) ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshOutput).ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesSshPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput
	ToVirtualmachinesPropertiesSshPtrOutputWithContext(context.Context) VirtualmachinesPropertiesSshPtrOutput
}

type virtualmachinesPropertiesSshPtrType VirtualmachinesPropertiesSshArgs

func VirtualmachinesPropertiesSshPtr(v *VirtualmachinesPropertiesSshArgs) VirtualmachinesPropertiesSshPtrInput {
	return (*virtualmachinesPropertiesSshPtrType)(v)
}

func (*virtualmachinesPropertiesSshPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSsh)(nil)).Elem()
}

func (i *virtualmachinesPropertiesSshPtrType) ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput {
	return i.ToVirtualmachinesPropertiesSshPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesSshPtrType) ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshPtrOutput)
}

type VirtualmachinesPropertiesSshOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSshOutput) ToVirtualmachinesPropertiesSshOutput() VirtualmachinesPropertiesSshOutput {
	return o
}

func (o VirtualmachinesPropertiesSshOutput) ToVirtualmachinesPropertiesSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshOutput {
	return o
}

func (o VirtualmachinesPropertiesSshOutput) ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput {
	return o.ToVirtualmachinesPropertiesSshPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesSshOutput) ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesSsh) *VirtualmachinesPropertiesSsh {
		return &v
	}).(VirtualmachinesPropertiesSshPtrOutput)
}

func (o VirtualmachinesPropertiesSshOutput) PublicKeys() VirtualmachinesPropertiesPublicKeysArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesSsh) []VirtualmachinesPropertiesPublicKeys { return v.PublicKeys }).(VirtualmachinesPropertiesPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesSshPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSshPtrOutput) ToVirtualmachinesPropertiesSshPtrOutput() VirtualmachinesPropertiesSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSshPtrOutput) ToVirtualmachinesPropertiesSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSshPtrOutput) Elem() VirtualmachinesPropertiesSshOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSsh) VirtualmachinesPropertiesSsh {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesSsh
		return ret
	}).(VirtualmachinesPropertiesSshOutput)
}

func (o VirtualmachinesPropertiesSshPtrOutput) PublicKeys() VirtualmachinesPropertiesPublicKeysArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSsh) []VirtualmachinesPropertiesPublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(VirtualmachinesPropertiesPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesSshSsh struct {
	PublicKeys []VirtualmachinesPropertiesPublicKeysPublicKeys `pulumi:"publicKeys"`
}





type VirtualmachinesPropertiesSshSshInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSshSshOutput() VirtualmachinesPropertiesSshSshOutput
	ToVirtualmachinesPropertiesSshSshOutputWithContext(context.Context) VirtualmachinesPropertiesSshSshOutput
}

type VirtualmachinesPropertiesSshSshArgs struct {
	PublicKeys VirtualmachinesPropertiesPublicKeysPublicKeysArrayInput `pulumi:"publicKeys"`
}

func (VirtualmachinesPropertiesSshSshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSshSsh)(nil)).Elem()
}

func (i VirtualmachinesPropertiesSshSshArgs) ToVirtualmachinesPropertiesSshSshOutput() VirtualmachinesPropertiesSshSshOutput {
	return i.ToVirtualmachinesPropertiesSshSshOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSshSshArgs) ToVirtualmachinesPropertiesSshSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshSshOutput)
}

func (i VirtualmachinesPropertiesSshSshArgs) ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput {
	return i.ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesSshSshArgs) ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshSshOutput).ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesSshSshPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput
	ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(context.Context) VirtualmachinesPropertiesSshSshPtrOutput
}

type virtualmachinesPropertiesSshSshPtrType VirtualmachinesPropertiesSshSshArgs

func VirtualmachinesPropertiesSshSshPtr(v *VirtualmachinesPropertiesSshSshArgs) VirtualmachinesPropertiesSshSshPtrInput {
	return (*virtualmachinesPropertiesSshSshPtrType)(v)
}

func (*virtualmachinesPropertiesSshSshPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSshSsh)(nil)).Elem()
}

func (i *virtualmachinesPropertiesSshSshPtrType) ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput {
	return i.ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesSshSshPtrType) ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesSshSshPtrOutput)
}

type VirtualmachinesPropertiesSshSshOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSshSshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesSshSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSshSshOutput) ToVirtualmachinesPropertiesSshSshOutput() VirtualmachinesPropertiesSshSshOutput {
	return o
}

func (o VirtualmachinesPropertiesSshSshOutput) ToVirtualmachinesPropertiesSshSshOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshOutput {
	return o
}

func (o VirtualmachinesPropertiesSshSshOutput) ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput {
	return o.ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesSshSshOutput) ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesSshSsh) *VirtualmachinesPropertiesSshSsh {
		return &v
	}).(VirtualmachinesPropertiesSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesSshSshOutput) PublicKeys() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesSshSsh) []VirtualmachinesPropertiesPublicKeysPublicKeys {
		return v.PublicKeys
	}).(VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesSshSshPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesSshSshPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesSshSsh)(nil)).Elem()
}

func (o VirtualmachinesPropertiesSshSshPtrOutput) ToVirtualmachinesPropertiesSshSshPtrOutput() VirtualmachinesPropertiesSshSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSshSshPtrOutput) ToVirtualmachinesPropertiesSshSshPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesSshSshPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesSshSshPtrOutput) Elem() VirtualmachinesPropertiesSshSshOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSshSsh) VirtualmachinesPropertiesSshSsh {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesSshSsh
		return ret
	}).(VirtualmachinesPropertiesSshSshOutput)
}

func (o VirtualmachinesPropertiesSshSshPtrOutput) PublicKeys() VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesSshSsh) []VirtualmachinesPropertiesPublicKeysPublicKeys {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput)
}

type VirtualmachinesPropertiesStorageProfile struct {
	DataDisks      []VirtualmachinesPropertiesDataDisks     `pulumi:"dataDisks"`
	ImageReference *VirtualmachinesPropertiesImageReference `pulumi:"imageReference"`
}





type VirtualmachinesPropertiesStorageProfileInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesStorageProfileOutput() VirtualmachinesPropertiesStorageProfileOutput
	ToVirtualmachinesPropertiesStorageProfileOutputWithContext(context.Context) VirtualmachinesPropertiesStorageProfileOutput
}

type VirtualmachinesPropertiesStorageProfileArgs struct {
	DataDisks      VirtualmachinesPropertiesDataDisksArrayInput    `pulumi:"dataDisks"`
	ImageReference VirtualmachinesPropertiesImageReferencePtrInput `pulumi:"imageReference"`
}

func (VirtualmachinesPropertiesStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesStorageProfile)(nil)).Elem()
}

func (i VirtualmachinesPropertiesStorageProfileArgs) ToVirtualmachinesPropertiesStorageProfileOutput() VirtualmachinesPropertiesStorageProfileOutput {
	return i.ToVirtualmachinesPropertiesStorageProfileOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesStorageProfileArgs) ToVirtualmachinesPropertiesStorageProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesStorageProfileOutput)
}

func (i VirtualmachinesPropertiesStorageProfileArgs) ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesStorageProfileArgs) ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesStorageProfileOutput).ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesStorageProfilePtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput
	ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput
}

type virtualmachinesPropertiesStorageProfilePtrType VirtualmachinesPropertiesStorageProfileArgs

func VirtualmachinesPropertiesStorageProfilePtr(v *VirtualmachinesPropertiesStorageProfileArgs) VirtualmachinesPropertiesStorageProfilePtrInput {
	return (*virtualmachinesPropertiesStorageProfilePtrType)(v)
}

func (*virtualmachinesPropertiesStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesStorageProfile)(nil)).Elem()
}

func (i *virtualmachinesPropertiesStorageProfilePtrType) ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput {
	return i.ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesStorageProfilePtrType) ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesStorageProfilePtrOutput)
}

type VirtualmachinesPropertiesStorageProfileOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesStorageProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ToVirtualmachinesPropertiesStorageProfileOutput() VirtualmachinesPropertiesStorageProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ToVirtualmachinesPropertiesStorageProfileOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfileOutput {
	return o
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput {
	return o.ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesStorageProfile) *VirtualmachinesPropertiesStorageProfile {
		return &v
	}).(VirtualmachinesPropertiesStorageProfilePtrOutput)
}

func (o VirtualmachinesPropertiesStorageProfileOutput) DataDisks() VirtualmachinesPropertiesDataDisksArrayOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesStorageProfile) []VirtualmachinesPropertiesDataDisks {
		return v.DataDisks
	}).(VirtualmachinesPropertiesDataDisksArrayOutput)
}

func (o VirtualmachinesPropertiesStorageProfileOutput) ImageReference() VirtualmachinesPropertiesImageReferencePtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesStorageProfile) *VirtualmachinesPropertiesImageReference {
		return v.ImageReference
	}).(VirtualmachinesPropertiesImageReferencePtrOutput)
}

type VirtualmachinesPropertiesStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesStorageProfile)(nil)).Elem()
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) ToVirtualmachinesPropertiesStorageProfilePtrOutput() VirtualmachinesPropertiesStorageProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) ToVirtualmachinesPropertiesStorageProfilePtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesStorageProfilePtrOutput {
	return o
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) Elem() VirtualmachinesPropertiesStorageProfileOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesStorageProfile) VirtualmachinesPropertiesStorageProfile {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesStorageProfile
		return ret
	}).(VirtualmachinesPropertiesStorageProfileOutput)
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) DataDisks() VirtualmachinesPropertiesDataDisksArrayOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesStorageProfile) []VirtualmachinesPropertiesDataDisks {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(VirtualmachinesPropertiesDataDisksArrayOutput)
}

func (o VirtualmachinesPropertiesStorageProfilePtrOutput) ImageReference() VirtualmachinesPropertiesImageReferencePtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesStorageProfile) *VirtualmachinesPropertiesImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(VirtualmachinesPropertiesImageReferencePtrOutput)
}

type VirtualmachinesPropertiesWindowsConfiguration struct {
	EnableAutomaticUpdates *bool                            `pulumi:"enableAutomaticUpdates"`
	Ssh                    *VirtualmachinesPropertiesSshSsh `pulumi:"ssh"`
	TimeZone               *string                          `pulumi:"timeZone"`
}





type VirtualmachinesPropertiesWindowsConfigurationInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesWindowsConfigurationOutput() VirtualmachinesPropertiesWindowsConfigurationOutput
	ToVirtualmachinesPropertiesWindowsConfigurationOutputWithContext(context.Context) VirtualmachinesPropertiesWindowsConfigurationOutput
}

type VirtualmachinesPropertiesWindowsConfigurationArgs struct {
	EnableAutomaticUpdates pulumi.BoolPtrInput                     `pulumi:"enableAutomaticUpdates"`
	Ssh                    VirtualmachinesPropertiesSshSshPtrInput `pulumi:"ssh"`
	TimeZone               pulumi.StringPtrInput                   `pulumi:"timeZone"`
}

func (VirtualmachinesPropertiesWindowsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesWindowsConfiguration)(nil)).Elem()
}

func (i VirtualmachinesPropertiesWindowsConfigurationArgs) ToVirtualmachinesPropertiesWindowsConfigurationOutput() VirtualmachinesPropertiesWindowsConfigurationOutput {
	return i.ToVirtualmachinesPropertiesWindowsConfigurationOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesWindowsConfigurationArgs) ToVirtualmachinesPropertiesWindowsConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesWindowsConfigurationOutput)
}

func (i VirtualmachinesPropertiesWindowsConfigurationArgs) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return i.ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualmachinesPropertiesWindowsConfigurationArgs) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesWindowsConfigurationOutput).ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx)
}









type VirtualmachinesPropertiesWindowsConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput
	ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput
}

type virtualmachinesPropertiesWindowsConfigurationPtrType VirtualmachinesPropertiesWindowsConfigurationArgs

func VirtualmachinesPropertiesWindowsConfigurationPtr(v *VirtualmachinesPropertiesWindowsConfigurationArgs) VirtualmachinesPropertiesWindowsConfigurationPtrInput {
	return (*virtualmachinesPropertiesWindowsConfigurationPtrType)(v)
}

func (*virtualmachinesPropertiesWindowsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesWindowsConfiguration)(nil)).Elem()
}

func (i *virtualmachinesPropertiesWindowsConfigurationPtrType) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return i.ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualmachinesPropertiesWindowsConfigurationPtrType) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachinesPropertiesWindowsConfigurationPtrOutput)
}

type VirtualmachinesPropertiesWindowsConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesWindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualmachinesPropertiesWindowsConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) ToVirtualmachinesPropertiesWindowsConfigurationOutput() VirtualmachinesPropertiesWindowsConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) ToVirtualmachinesPropertiesWindowsConfigurationOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationOutput {
	return o
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o.ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualmachinesPropertiesWindowsConfiguration) *VirtualmachinesPropertiesWindowsConfiguration {
		return &v
	}).(VirtualmachinesPropertiesWindowsConfigurationPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesWindowsConfiguration) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) Ssh() VirtualmachinesPropertiesSshSshPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesWindowsConfiguration) *VirtualmachinesPropertiesSshSsh { return v.Ssh }).(VirtualmachinesPropertiesSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualmachinesPropertiesWindowsConfiguration) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type VirtualmachinesPropertiesWindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualmachinesPropertiesWindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachinesPropertiesWindowsConfiguration)(nil)).Elem()
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutput() VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) ToVirtualmachinesPropertiesWindowsConfigurationPtrOutputWithContext(ctx context.Context) VirtualmachinesPropertiesWindowsConfigurationPtrOutput {
	return o
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) Elem() VirtualmachinesPropertiesWindowsConfigurationOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesWindowsConfiguration) VirtualmachinesPropertiesWindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualmachinesPropertiesWindowsConfiguration
		return ret
	}).(VirtualmachinesPropertiesWindowsConfigurationOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesWindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) Ssh() VirtualmachinesPropertiesSshSshPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesWindowsConfiguration) *VirtualmachinesPropertiesSshSsh {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(VirtualmachinesPropertiesSshSshPtrOutput)
}

func (o VirtualmachinesPropertiesWindowsConfigurationPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachinesPropertiesWindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesIpConfigurationReferences struct {
	Id *string `pulumi:"id"`
}





type VirtualnetworksPropertiesIpConfigurationReferencesInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesIpConfigurationReferencesOutput() VirtualnetworksPropertiesIpConfigurationReferencesOutput
	ToVirtualnetworksPropertiesIpConfigurationReferencesOutputWithContext(context.Context) VirtualnetworksPropertiesIpConfigurationReferencesOutput
}

type VirtualnetworksPropertiesIpConfigurationReferencesArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualnetworksPropertiesIpConfigurationReferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesIpConfigurationReferences)(nil)).Elem()
}

func (i VirtualnetworksPropertiesIpConfigurationReferencesArgs) ToVirtualnetworksPropertiesIpConfigurationReferencesOutput() VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return i.ToVirtualnetworksPropertiesIpConfigurationReferencesOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesIpConfigurationReferencesArgs) ToVirtualnetworksPropertiesIpConfigurationReferencesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesIpConfigurationReferencesOutput)
}





type VirtualnetworksPropertiesIpConfigurationReferencesArrayInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutput() VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput
	ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutputWithContext(context.Context) VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput
}

type VirtualnetworksPropertiesIpConfigurationReferencesArray []VirtualnetworksPropertiesIpConfigurationReferencesInput

func (VirtualnetworksPropertiesIpConfigurationReferencesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesIpConfigurationReferences)(nil)).Elem()
}

func (i VirtualnetworksPropertiesIpConfigurationReferencesArray) ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutput() VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return i.ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesIpConfigurationReferencesArray) ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput)
}

type VirtualnetworksPropertiesIpConfigurationReferencesOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesIpConfigurationReferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesIpConfigurationReferences)(nil)).Elem()
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesOutput) ToVirtualnetworksPropertiesIpConfigurationReferencesOutput() VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return o
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesOutput) ToVirtualnetworksPropertiesIpConfigurationReferencesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return o
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesIpConfigurationReferences) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesIpConfigurationReferences)(nil)).Elem()
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput) ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutput() VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput) ToVirtualnetworksPropertiesIpConfigurationReferencesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesIpConfigurationReferencesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesIpConfigurationReferences {
		return vs[0].([]VirtualnetworksPropertiesIpConfigurationReferences)[vs[1].(int)]
	}).(VirtualnetworksPropertiesIpConfigurationReferencesOutput)
}

type VirtualnetworksPropertiesResponseIpConfigurationReferences struct {
	Id *string `pulumi:"id"`
}

type VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesResponseIpConfigurationReferences)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput) ToVirtualnetworksPropertiesResponseIpConfigurationReferencesOutput() VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput) ToVirtualnetworksPropertiesResponseIpConfigurationReferencesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseIpConfigurationReferences) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesResponseIpConfigurationReferences)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput) ToVirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput() VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput) ToVirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesResponseIpConfigurationReferences {
		return vs[0].([]VirtualnetworksPropertiesResponseIpConfigurationReferences)[vs[1].(int)]
	}).(VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput)
}

type VirtualnetworksPropertiesResponseRouteTable struct {
	Id     *string                                   `pulumi:"id"`
	Name   *string                                   `pulumi:"name"`
	Routes []VirtualnetworksPropertiesResponseRoutes `pulumi:"routes"`
	Type   *string                                   `pulumi:"type"`
}

type VirtualnetworksPropertiesResponseRouteTableOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesResponseRouteTable)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) ToVirtualnetworksPropertiesResponseRouteTableOutput() VirtualnetworksPropertiesResponseRouteTableOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) ToVirtualnetworksPropertiesResponseRouteTableOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseRouteTableOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRouteTable) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRouteTable) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) Routes() VirtualnetworksPropertiesResponseRoutesArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRouteTable) []VirtualnetworksPropertiesResponseRoutes {
		return v.Routes
	}).(VirtualnetworksPropertiesResponseRoutesArrayOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTableOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRouteTable) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesResponseRouteTablePtrOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseRouteTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualnetworksPropertiesResponseRouteTable)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) ToVirtualnetworksPropertiesResponseRouteTablePtrOutput() VirtualnetworksPropertiesResponseRouteTablePtrOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) ToVirtualnetworksPropertiesResponseRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseRouteTablePtrOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Elem() VirtualnetworksPropertiesResponseRouteTableOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) VirtualnetworksPropertiesResponseRouteTable {
		if v != nil {
			return *v
		}
		var ret VirtualnetworksPropertiesResponseRouteTable
		return ret
	}).(VirtualnetworksPropertiesResponseRouteTableOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Routes() VirtualnetworksPropertiesResponseRoutesArrayOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) []VirtualnetworksPropertiesResponseRoutes {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(VirtualnetworksPropertiesResponseRoutesArrayOutput)
}

func (o VirtualnetworksPropertiesResponseRouteTablePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesResponseRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesResponseRoutes struct {
	AddressPrefix    *string `pulumi:"addressPrefix"`
	Name             *string `pulumi:"name"`
	NextHopIpAddress *string `pulumi:"nextHopIpAddress"`
}

type VirtualnetworksPropertiesResponseRoutesOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseRoutesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesResponseRoutes)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) ToVirtualnetworksPropertiesResponseRoutesOutput() VirtualnetworksPropertiesResponseRoutesOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) ToVirtualnetworksPropertiesResponseRoutesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseRoutesOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRoutes) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRoutes) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseRoutesOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseRoutes) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesResponseRoutesArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseRoutesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesResponseRoutes)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseRoutesArrayOutput) ToVirtualnetworksPropertiesResponseRoutesArrayOutput() VirtualnetworksPropertiesResponseRoutesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRoutesArrayOutput) ToVirtualnetworksPropertiesResponseRoutesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseRoutesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseRoutesArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesResponseRoutesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesResponseRoutes {
		return vs[0].([]VirtualnetworksPropertiesResponseRoutes)[vs[1].(int)]
	}).(VirtualnetworksPropertiesResponseRoutesOutput)
}

type VirtualnetworksPropertiesResponseSubnets struct {
	AddressPrefix             *string                                                      `pulumi:"addressPrefix"`
	AddressPrefixes           []string                                                     `pulumi:"addressPrefixes"`
	IpAllocationMethod        *string                                                      `pulumi:"ipAllocationMethod"`
	IpConfigurationReferences []VirtualnetworksPropertiesResponseIpConfigurationReferences `pulumi:"ipConfigurationReferences"`
	IpPools                   []IPPoolResponse                                             `pulumi:"ipPools"`
	Name                      *string                                                      `pulumi:"name"`
	RouteTable                *VirtualnetworksPropertiesResponseRouteTable                 `pulumi:"routeTable"`
	Vlan                      *int                                                         `pulumi:"vlan"`
}

type VirtualnetworksPropertiesResponseSubnetsOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseSubnetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesResponseSubnets)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) ToVirtualnetworksPropertiesResponseSubnetsOutput() VirtualnetworksPropertiesResponseSubnetsOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) ToVirtualnetworksPropertiesResponseSubnetsOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseSubnetsOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) IpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *string { return v.IpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) IpConfigurationReferences() VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) []VirtualnetworksPropertiesResponseIpConfigurationReferences {
		return v.IpConfigurationReferences
	}).(VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) IpPools() IPPoolResponseArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) []IPPoolResponse { return v.IpPools }).(IPPoolResponseArrayOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) RouteTable() VirtualnetworksPropertiesResponseRouteTablePtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *VirtualnetworksPropertiesResponseRouteTable {
		return v.RouteTable
	}).(VirtualnetworksPropertiesResponseRouteTablePtrOutput)
}

func (o VirtualnetworksPropertiesResponseSubnetsOutput) Vlan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesResponseSubnets) *int { return v.Vlan }).(pulumi.IntPtrOutput)
}

type VirtualnetworksPropertiesResponseSubnetsArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesResponseSubnetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesResponseSubnets)(nil)).Elem()
}

func (o VirtualnetworksPropertiesResponseSubnetsArrayOutput) ToVirtualnetworksPropertiesResponseSubnetsArrayOutput() VirtualnetworksPropertiesResponseSubnetsArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseSubnetsArrayOutput) ToVirtualnetworksPropertiesResponseSubnetsArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesResponseSubnetsArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesResponseSubnetsArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesResponseSubnetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesResponseSubnets {
		return vs[0].([]VirtualnetworksPropertiesResponseSubnets)[vs[1].(int)]
	}).(VirtualnetworksPropertiesResponseSubnetsOutput)
}

type VirtualnetworksPropertiesRouteTable struct {
	Id     *string                           `pulumi:"id"`
	Name   *string                           `pulumi:"name"`
	Routes []VirtualnetworksPropertiesRoutes `pulumi:"routes"`
	Type   *string                           `pulumi:"type"`
}





type VirtualnetworksPropertiesRouteTableInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesRouteTableOutput() VirtualnetworksPropertiesRouteTableOutput
	ToVirtualnetworksPropertiesRouteTableOutputWithContext(context.Context) VirtualnetworksPropertiesRouteTableOutput
}

type VirtualnetworksPropertiesRouteTableArgs struct {
	Id     pulumi.StringPtrInput                     `pulumi:"id"`
	Name   pulumi.StringPtrInput                     `pulumi:"name"`
	Routes VirtualnetworksPropertiesRoutesArrayInput `pulumi:"routes"`
	Type   pulumi.StringPtrInput                     `pulumi:"type"`
}

func (VirtualnetworksPropertiesRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesRouteTable)(nil)).Elem()
}

func (i VirtualnetworksPropertiesRouteTableArgs) ToVirtualnetworksPropertiesRouteTableOutput() VirtualnetworksPropertiesRouteTableOutput {
	return i.ToVirtualnetworksPropertiesRouteTableOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesRouteTableArgs) ToVirtualnetworksPropertiesRouteTableOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRouteTableOutput)
}

func (i VirtualnetworksPropertiesRouteTableArgs) ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput {
	return i.ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesRouteTableArgs) ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRouteTableOutput).ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx)
}









type VirtualnetworksPropertiesRouteTablePtrInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput
	ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(context.Context) VirtualnetworksPropertiesRouteTablePtrOutput
}

type virtualnetworksPropertiesRouteTablePtrType VirtualnetworksPropertiesRouteTableArgs

func VirtualnetworksPropertiesRouteTablePtr(v *VirtualnetworksPropertiesRouteTableArgs) VirtualnetworksPropertiesRouteTablePtrInput {
	return (*virtualnetworksPropertiesRouteTablePtrType)(v)
}

func (*virtualnetworksPropertiesRouteTablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualnetworksPropertiesRouteTable)(nil)).Elem()
}

func (i *virtualnetworksPropertiesRouteTablePtrType) ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput {
	return i.ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(context.Background())
}

func (i *virtualnetworksPropertiesRouteTablePtrType) ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRouteTablePtrOutput)
}

type VirtualnetworksPropertiesRouteTableOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesRouteTable)(nil)).Elem()
}

func (o VirtualnetworksPropertiesRouteTableOutput) ToVirtualnetworksPropertiesRouteTableOutput() VirtualnetworksPropertiesRouteTableOutput {
	return o
}

func (o VirtualnetworksPropertiesRouteTableOutput) ToVirtualnetworksPropertiesRouteTableOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTableOutput {
	return o
}

func (o VirtualnetworksPropertiesRouteTableOutput) ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput {
	return o.ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(context.Background())
}

func (o VirtualnetworksPropertiesRouteTableOutput) ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTablePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualnetworksPropertiesRouteTable) *VirtualnetworksPropertiesRouteTable {
		return &v
	}).(VirtualnetworksPropertiesRouteTablePtrOutput)
}

func (o VirtualnetworksPropertiesRouteTableOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRouteTable) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRouteTableOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRouteTable) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRouteTableOutput) Routes() VirtualnetworksPropertiesRoutesArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRouteTable) []VirtualnetworksPropertiesRoutes { return v.Routes }).(VirtualnetworksPropertiesRoutesArrayOutput)
}

func (o VirtualnetworksPropertiesRouteTableOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRouteTable) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesRouteTablePtrOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesRouteTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualnetworksPropertiesRouteTable)(nil)).Elem()
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) ToVirtualnetworksPropertiesRouteTablePtrOutput() VirtualnetworksPropertiesRouteTablePtrOutput {
	return o
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) ToVirtualnetworksPropertiesRouteTablePtrOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRouteTablePtrOutput {
	return o
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Elem() VirtualnetworksPropertiesRouteTableOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) VirtualnetworksPropertiesRouteTable {
		if v != nil {
			return *v
		}
		var ret VirtualnetworksPropertiesRouteTable
		return ret
	}).(VirtualnetworksPropertiesRouteTableOutput)
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Routes() VirtualnetworksPropertiesRoutesArrayOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) []VirtualnetworksPropertiesRoutes {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(VirtualnetworksPropertiesRoutesArrayOutput)
}

func (o VirtualnetworksPropertiesRouteTablePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualnetworksPropertiesRouteTable) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesRoutes struct {
	AddressPrefix    *string `pulumi:"addressPrefix"`
	Name             *string `pulumi:"name"`
	NextHopIpAddress *string `pulumi:"nextHopIpAddress"`
}





type VirtualnetworksPropertiesRoutesInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesRoutesOutput() VirtualnetworksPropertiesRoutesOutput
	ToVirtualnetworksPropertiesRoutesOutputWithContext(context.Context) VirtualnetworksPropertiesRoutesOutput
}

type VirtualnetworksPropertiesRoutesArgs struct {
	AddressPrefix    pulumi.StringPtrInput `pulumi:"addressPrefix"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	NextHopIpAddress pulumi.StringPtrInput `pulumi:"nextHopIpAddress"`
}

func (VirtualnetworksPropertiesRoutesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesRoutes)(nil)).Elem()
}

func (i VirtualnetworksPropertiesRoutesArgs) ToVirtualnetworksPropertiesRoutesOutput() VirtualnetworksPropertiesRoutesOutput {
	return i.ToVirtualnetworksPropertiesRoutesOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesRoutesArgs) ToVirtualnetworksPropertiesRoutesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRoutesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRoutesOutput)
}





type VirtualnetworksPropertiesRoutesArrayInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesRoutesArrayOutput() VirtualnetworksPropertiesRoutesArrayOutput
	ToVirtualnetworksPropertiesRoutesArrayOutputWithContext(context.Context) VirtualnetworksPropertiesRoutesArrayOutput
}

type VirtualnetworksPropertiesRoutesArray []VirtualnetworksPropertiesRoutesInput

func (VirtualnetworksPropertiesRoutesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesRoutes)(nil)).Elem()
}

func (i VirtualnetworksPropertiesRoutesArray) ToVirtualnetworksPropertiesRoutesArrayOutput() VirtualnetworksPropertiesRoutesArrayOutput {
	return i.ToVirtualnetworksPropertiesRoutesArrayOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesRoutesArray) ToVirtualnetworksPropertiesRoutesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRoutesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesRoutesArrayOutput)
}

type VirtualnetworksPropertiesRoutesOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesRoutesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesRoutes)(nil)).Elem()
}

func (o VirtualnetworksPropertiesRoutesOutput) ToVirtualnetworksPropertiesRoutesOutput() VirtualnetworksPropertiesRoutesOutput {
	return o
}

func (o VirtualnetworksPropertiesRoutesOutput) ToVirtualnetworksPropertiesRoutesOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRoutesOutput {
	return o
}

func (o VirtualnetworksPropertiesRoutesOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRoutes) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRoutesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRoutes) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesRoutesOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesRoutes) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

type VirtualnetworksPropertiesRoutesArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesRoutesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesRoutes)(nil)).Elem()
}

func (o VirtualnetworksPropertiesRoutesArrayOutput) ToVirtualnetworksPropertiesRoutesArrayOutput() VirtualnetworksPropertiesRoutesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesRoutesArrayOutput) ToVirtualnetworksPropertiesRoutesArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesRoutesArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesRoutesArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesRoutesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesRoutes {
		return vs[0].([]VirtualnetworksPropertiesRoutes)[vs[1].(int)]
	}).(VirtualnetworksPropertiesRoutesOutput)
}

type VirtualnetworksPropertiesSubnets struct {
	AddressPrefix             *string                                              `pulumi:"addressPrefix"`
	AddressPrefixes           []string                                             `pulumi:"addressPrefixes"`
	IpAllocationMethod        *string                                              `pulumi:"ipAllocationMethod"`
	IpConfigurationReferences []VirtualnetworksPropertiesIpConfigurationReferences `pulumi:"ipConfigurationReferences"`
	IpPools                   []IPPool                                             `pulumi:"ipPools"`
	Name                      *string                                              `pulumi:"name"`
	RouteTable                *VirtualnetworksPropertiesRouteTable                 `pulumi:"routeTable"`
	Vlan                      *int                                                 `pulumi:"vlan"`
}





type VirtualnetworksPropertiesSubnetsInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesSubnetsOutput() VirtualnetworksPropertiesSubnetsOutput
	ToVirtualnetworksPropertiesSubnetsOutputWithContext(context.Context) VirtualnetworksPropertiesSubnetsOutput
}

type VirtualnetworksPropertiesSubnetsArgs struct {
	AddressPrefix             pulumi.StringPtrInput                                        `pulumi:"addressPrefix"`
	AddressPrefixes           pulumi.StringArrayInput                                      `pulumi:"addressPrefixes"`
	IpAllocationMethod        pulumi.StringPtrInput                                        `pulumi:"ipAllocationMethod"`
	IpConfigurationReferences VirtualnetworksPropertiesIpConfigurationReferencesArrayInput `pulumi:"ipConfigurationReferences"`
	IpPools                   IPPoolArrayInput                                             `pulumi:"ipPools"`
	Name                      pulumi.StringPtrInput                                        `pulumi:"name"`
	RouteTable                VirtualnetworksPropertiesRouteTablePtrInput                  `pulumi:"routeTable"`
	Vlan                      pulumi.IntPtrInput                                           `pulumi:"vlan"`
}

func (VirtualnetworksPropertiesSubnetsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesSubnets)(nil)).Elem()
}

func (i VirtualnetworksPropertiesSubnetsArgs) ToVirtualnetworksPropertiesSubnetsOutput() VirtualnetworksPropertiesSubnetsOutput {
	return i.ToVirtualnetworksPropertiesSubnetsOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesSubnetsArgs) ToVirtualnetworksPropertiesSubnetsOutputWithContext(ctx context.Context) VirtualnetworksPropertiesSubnetsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesSubnetsOutput)
}





type VirtualnetworksPropertiesSubnetsArrayInput interface {
	pulumi.Input

	ToVirtualnetworksPropertiesSubnetsArrayOutput() VirtualnetworksPropertiesSubnetsArrayOutput
	ToVirtualnetworksPropertiesSubnetsArrayOutputWithContext(context.Context) VirtualnetworksPropertiesSubnetsArrayOutput
}

type VirtualnetworksPropertiesSubnetsArray []VirtualnetworksPropertiesSubnetsInput

func (VirtualnetworksPropertiesSubnetsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesSubnets)(nil)).Elem()
}

func (i VirtualnetworksPropertiesSubnetsArray) ToVirtualnetworksPropertiesSubnetsArrayOutput() VirtualnetworksPropertiesSubnetsArrayOutput {
	return i.ToVirtualnetworksPropertiesSubnetsArrayOutputWithContext(context.Background())
}

func (i VirtualnetworksPropertiesSubnetsArray) ToVirtualnetworksPropertiesSubnetsArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesSubnetsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualnetworksPropertiesSubnetsArrayOutput)
}

type VirtualnetworksPropertiesSubnetsOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesSubnetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualnetworksPropertiesSubnets)(nil)).Elem()
}

func (o VirtualnetworksPropertiesSubnetsOutput) ToVirtualnetworksPropertiesSubnetsOutput() VirtualnetworksPropertiesSubnetsOutput {
	return o
}

func (o VirtualnetworksPropertiesSubnetsOutput) ToVirtualnetworksPropertiesSubnetsOutputWithContext(ctx context.Context) VirtualnetworksPropertiesSubnetsOutput {
	return o
}

func (o VirtualnetworksPropertiesSubnetsOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) IpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *string { return v.IpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) IpConfigurationReferences() VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) []VirtualnetworksPropertiesIpConfigurationReferences {
		return v.IpConfigurationReferences
	}).(VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) IpPools() IPPoolArrayOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) []IPPool { return v.IpPools }).(IPPoolArrayOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) RouteTable() VirtualnetworksPropertiesRouteTablePtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *VirtualnetworksPropertiesRouteTable { return v.RouteTable }).(VirtualnetworksPropertiesRouteTablePtrOutput)
}

func (o VirtualnetworksPropertiesSubnetsOutput) Vlan() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualnetworksPropertiesSubnets) *int { return v.Vlan }).(pulumi.IntPtrOutput)
}

type VirtualnetworksPropertiesSubnetsArrayOutput struct{ *pulumi.OutputState }

func (VirtualnetworksPropertiesSubnetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualnetworksPropertiesSubnets)(nil)).Elem()
}

func (o VirtualnetworksPropertiesSubnetsArrayOutput) ToVirtualnetworksPropertiesSubnetsArrayOutput() VirtualnetworksPropertiesSubnetsArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesSubnetsArrayOutput) ToVirtualnetworksPropertiesSubnetsArrayOutputWithContext(ctx context.Context) VirtualnetworksPropertiesSubnetsArrayOutput {
	return o
}

func (o VirtualnetworksPropertiesSubnetsArrayOutput) Index(i pulumi.IntInput) VirtualnetworksPropertiesSubnetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualnetworksPropertiesSubnets {
		return vs[0].([]VirtualnetworksPropertiesSubnets)[vs[1].(int)]
	}).(VirtualnetworksPropertiesSubnetsOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageStatusResponseOutput{})
	pulumi.RegisterOutputType(IPPoolOutput{})
	pulumi.RegisterOutputType(IPPoolArrayOutput{})
	pulumi.RegisterOutputType(IPPoolInfoResponseOutput{})
	pulumi.RegisterOutputType(IPPoolInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(IPPoolResponseOutput{})
	pulumi.RegisterOutputType(IPPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(InterfaceDNSSettingsOutput{})
	pulumi.RegisterOutputType(InterfaceDNSSettingsPtrOutput{})
	pulumi.RegisterOutputType(InterfaceDNSSettingsResponseOutput{})
	pulumi.RegisterOutputType(InterfaceDNSSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(IpConfigurationOutput{})
	pulumi.RegisterOutputType(IpConfigurationArrayOutput{})
	pulumi.RegisterOutputType(IpConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(IpConfigurationPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponseOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponsePropertiesOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponseSubnetOutput{})
	pulumi.RegisterOutputType(IpConfigurationResponseSubnetPtrOutput{})
	pulumi.RegisterOutputType(IpConfigurationSubnetOutput{})
	pulumi.RegisterOutputType(IpConfigurationSubnetPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceStatusResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskStatusResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineStatusResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkStatusResponseOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesDataDisksOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesDataDisksArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesDynamicMemoryConfigOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesDynamicMemoryConfigPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesHardwareProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesHardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesImageReferenceOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesImageReferencePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesLinuxConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesLinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesNetworkInterfacesOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesNetworkInterfacesArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesOsProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesOsProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesPublicKeysOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesPublicKeysArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesPublicKeysPublicKeysOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesPublicKeysPublicKeysArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseDataDisksOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseDataDisksArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseDynamicMemoryConfigOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseDynamicMemoryConfigPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseHardwareProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseHardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseImageReferenceOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseImageReferencePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseLinuxConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseLinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseNetworkInterfacesOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseNetworkInterfacesArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseOsProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseOsProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponsePublicKeysOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponsePublicKeysArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponsePublicKeysPublicKeysOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponsePublicKeysPublicKeysArrayOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSecurityProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSshOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSshPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSshSshOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseSshSshPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseStorageProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseWindowsConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesResponseWindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSecurityProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSshOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSshPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSshSshOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesSshSshPtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesStorageProfileOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesWindowsConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualmachinesPropertiesWindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesIpConfigurationReferencesOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesIpConfigurationReferencesArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseIpConfigurationReferencesOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseIpConfigurationReferencesArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseRouteTableOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseRouteTablePtrOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseRoutesOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseRoutesArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseSubnetsOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesResponseSubnetsArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesRouteTableOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesRouteTablePtrOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesRoutesOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesRoutesArrayOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesSubnetsOutput{})
	pulumi.RegisterOutputType(VirtualnetworksPropertiesSubnetsArrayOutput{})
}
