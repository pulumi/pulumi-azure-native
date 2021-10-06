


package v20161010

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdditionalRegion struct {
	Location         string                       `pulumi:"location"`
	SkuType          SkuType                      `pulumi:"skuType"`
	SkuUnitCount     *int                         `pulumi:"skuUnitCount"`
	Vpnconfiguration *VirtualNetworkConfiguration `pulumi:"vpnconfiguration"`
}





type AdditionalRegionInput interface {
	pulumi.Input

	ToAdditionalRegionOutput() AdditionalRegionOutput
	ToAdditionalRegionOutputWithContext(context.Context) AdditionalRegionOutput
}

type AdditionalRegionArgs struct {
	Location         pulumi.StringInput                  `pulumi:"location"`
	SkuType          SkuTypeInput                        `pulumi:"skuType"`
	SkuUnitCount     pulumi.IntPtrInput                  `pulumi:"skuUnitCount"`
	Vpnconfiguration VirtualNetworkConfigurationPtrInput `pulumi:"vpnconfiguration"`
}

func (AdditionalRegionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalRegion)(nil)).Elem()
}

func (i AdditionalRegionArgs) ToAdditionalRegionOutput() AdditionalRegionOutput {
	return i.ToAdditionalRegionOutputWithContext(context.Background())
}

func (i AdditionalRegionArgs) ToAdditionalRegionOutputWithContext(ctx context.Context) AdditionalRegionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalRegionOutput)
}





type AdditionalRegionArrayInput interface {
	pulumi.Input

	ToAdditionalRegionArrayOutput() AdditionalRegionArrayOutput
	ToAdditionalRegionArrayOutputWithContext(context.Context) AdditionalRegionArrayOutput
}

type AdditionalRegionArray []AdditionalRegionInput

func (AdditionalRegionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalRegion)(nil)).Elem()
}

func (i AdditionalRegionArray) ToAdditionalRegionArrayOutput() AdditionalRegionArrayOutput {
	return i.ToAdditionalRegionArrayOutputWithContext(context.Background())
}

func (i AdditionalRegionArray) ToAdditionalRegionArrayOutputWithContext(ctx context.Context) AdditionalRegionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalRegionArrayOutput)
}

type AdditionalRegionOutput struct{ *pulumi.OutputState }

func (AdditionalRegionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalRegion)(nil)).Elem()
}

func (o AdditionalRegionOutput) ToAdditionalRegionOutput() AdditionalRegionOutput {
	return o
}

func (o AdditionalRegionOutput) ToAdditionalRegionOutputWithContext(ctx context.Context) AdditionalRegionOutput {
	return o
}

func (o AdditionalRegionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalRegion) string { return v.Location }).(pulumi.StringOutput)
}

func (o AdditionalRegionOutput) SkuType() SkuTypeOutput {
	return o.ApplyT(func(v AdditionalRegion) SkuType { return v.SkuType }).(SkuTypeOutput)
}

func (o AdditionalRegionOutput) SkuUnitCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AdditionalRegion) *int { return v.SkuUnitCount }).(pulumi.IntPtrOutput)
}

func (o AdditionalRegionOutput) Vpnconfiguration() VirtualNetworkConfigurationPtrOutput {
	return o.ApplyT(func(v AdditionalRegion) *VirtualNetworkConfiguration { return v.Vpnconfiguration }).(VirtualNetworkConfigurationPtrOutput)
}

type AdditionalRegionArrayOutput struct{ *pulumi.OutputState }

func (AdditionalRegionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalRegion)(nil)).Elem()
}

func (o AdditionalRegionArrayOutput) ToAdditionalRegionArrayOutput() AdditionalRegionArrayOutput {
	return o
}

func (o AdditionalRegionArrayOutput) ToAdditionalRegionArrayOutputWithContext(ctx context.Context) AdditionalRegionArrayOutput {
	return o
}

func (o AdditionalRegionArrayOutput) Index(i pulumi.IntInput) AdditionalRegionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalRegion {
		return vs[0].([]AdditionalRegion)[vs[1].(int)]
	}).(AdditionalRegionOutput)
}

type AdditionalRegionResponse struct {
	Location         string                               `pulumi:"location"`
	SkuType          string                               `pulumi:"skuType"`
	SkuUnitCount     *int                                 `pulumi:"skuUnitCount"`
	StaticIPs        []string                             `pulumi:"staticIPs"`
	Vpnconfiguration *VirtualNetworkConfigurationResponse `pulumi:"vpnconfiguration"`
}





type AdditionalRegionResponseInput interface {
	pulumi.Input

	ToAdditionalRegionResponseOutput() AdditionalRegionResponseOutput
	ToAdditionalRegionResponseOutputWithContext(context.Context) AdditionalRegionResponseOutput
}

type AdditionalRegionResponseArgs struct {
	Location         pulumi.StringInput                          `pulumi:"location"`
	SkuType          pulumi.StringInput                          `pulumi:"skuType"`
	SkuUnitCount     pulumi.IntPtrInput                          `pulumi:"skuUnitCount"`
	StaticIPs        pulumi.StringArrayInput                     `pulumi:"staticIPs"`
	Vpnconfiguration VirtualNetworkConfigurationResponsePtrInput `pulumi:"vpnconfiguration"`
}

func (AdditionalRegionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalRegionResponse)(nil)).Elem()
}

func (i AdditionalRegionResponseArgs) ToAdditionalRegionResponseOutput() AdditionalRegionResponseOutput {
	return i.ToAdditionalRegionResponseOutputWithContext(context.Background())
}

func (i AdditionalRegionResponseArgs) ToAdditionalRegionResponseOutputWithContext(ctx context.Context) AdditionalRegionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalRegionResponseOutput)
}





type AdditionalRegionResponseArrayInput interface {
	pulumi.Input

	ToAdditionalRegionResponseArrayOutput() AdditionalRegionResponseArrayOutput
	ToAdditionalRegionResponseArrayOutputWithContext(context.Context) AdditionalRegionResponseArrayOutput
}

type AdditionalRegionResponseArray []AdditionalRegionResponseInput

func (AdditionalRegionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalRegionResponse)(nil)).Elem()
}

func (i AdditionalRegionResponseArray) ToAdditionalRegionResponseArrayOutput() AdditionalRegionResponseArrayOutput {
	return i.ToAdditionalRegionResponseArrayOutputWithContext(context.Background())
}

func (i AdditionalRegionResponseArray) ToAdditionalRegionResponseArrayOutputWithContext(ctx context.Context) AdditionalRegionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalRegionResponseArrayOutput)
}

type AdditionalRegionResponseOutput struct{ *pulumi.OutputState }

func (AdditionalRegionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalRegionResponse)(nil)).Elem()
}

func (o AdditionalRegionResponseOutput) ToAdditionalRegionResponseOutput() AdditionalRegionResponseOutput {
	return o
}

func (o AdditionalRegionResponseOutput) ToAdditionalRegionResponseOutputWithContext(ctx context.Context) AdditionalRegionResponseOutput {
	return o
}

func (o AdditionalRegionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalRegionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o AdditionalRegionResponseOutput) SkuType() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalRegionResponse) string { return v.SkuType }).(pulumi.StringOutput)
}

func (o AdditionalRegionResponseOutput) SkuUnitCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AdditionalRegionResponse) *int { return v.SkuUnitCount }).(pulumi.IntPtrOutput)
}

func (o AdditionalRegionResponseOutput) StaticIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalRegionResponse) []string { return v.StaticIPs }).(pulumi.StringArrayOutput)
}

func (o AdditionalRegionResponseOutput) Vpnconfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AdditionalRegionResponse) *VirtualNetworkConfigurationResponse { return v.Vpnconfiguration }).(VirtualNetworkConfigurationResponsePtrOutput)
}

type AdditionalRegionResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalRegionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalRegionResponse)(nil)).Elem()
}

func (o AdditionalRegionResponseArrayOutput) ToAdditionalRegionResponseArrayOutput() AdditionalRegionResponseArrayOutput {
	return o
}

func (o AdditionalRegionResponseArrayOutput) ToAdditionalRegionResponseArrayOutputWithContext(ctx context.Context) AdditionalRegionResponseArrayOutput {
	return o
}

func (o AdditionalRegionResponseArrayOutput) Index(i pulumi.IntInput) AdditionalRegionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalRegionResponse {
		return vs[0].([]AdditionalRegionResponse)[vs[1].(int)]
	}).(AdditionalRegionResponseOutput)
}

type ApiManagementServiceSkuProperties struct {
	Capacity *int    `pulumi:"capacity"`
	Name     SkuType `pulumi:"name"`
}





type ApiManagementServiceSkuPropertiesInput interface {
	pulumi.Input

	ToApiManagementServiceSkuPropertiesOutput() ApiManagementServiceSkuPropertiesOutput
	ToApiManagementServiceSkuPropertiesOutputWithContext(context.Context) ApiManagementServiceSkuPropertiesOutput
}

type ApiManagementServiceSkuPropertiesArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     SkuTypeInput       `pulumi:"name"`
}

func (ApiManagementServiceSkuPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementServiceSkuProperties)(nil)).Elem()
}

func (i ApiManagementServiceSkuPropertiesArgs) ToApiManagementServiceSkuPropertiesOutput() ApiManagementServiceSkuPropertiesOutput {
	return i.ToApiManagementServiceSkuPropertiesOutputWithContext(context.Background())
}

func (i ApiManagementServiceSkuPropertiesArgs) ToApiManagementServiceSkuPropertiesOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceSkuPropertiesOutput)
}

func (i ApiManagementServiceSkuPropertiesArgs) ToApiManagementServiceSkuPropertiesPtrOutput() ApiManagementServiceSkuPropertiesPtrOutput {
	return i.ToApiManagementServiceSkuPropertiesPtrOutputWithContext(context.Background())
}

func (i ApiManagementServiceSkuPropertiesArgs) ToApiManagementServiceSkuPropertiesPtrOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceSkuPropertiesOutput).ToApiManagementServiceSkuPropertiesPtrOutputWithContext(ctx)
}









type ApiManagementServiceSkuPropertiesPtrInput interface {
	pulumi.Input

	ToApiManagementServiceSkuPropertiesPtrOutput() ApiManagementServiceSkuPropertiesPtrOutput
	ToApiManagementServiceSkuPropertiesPtrOutputWithContext(context.Context) ApiManagementServiceSkuPropertiesPtrOutput
}

type apiManagementServiceSkuPropertiesPtrType ApiManagementServiceSkuPropertiesArgs

func ApiManagementServiceSkuPropertiesPtr(v *ApiManagementServiceSkuPropertiesArgs) ApiManagementServiceSkuPropertiesPtrInput {
	return (*apiManagementServiceSkuPropertiesPtrType)(v)
}

func (*apiManagementServiceSkuPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementServiceSkuProperties)(nil)).Elem()
}

func (i *apiManagementServiceSkuPropertiesPtrType) ToApiManagementServiceSkuPropertiesPtrOutput() ApiManagementServiceSkuPropertiesPtrOutput {
	return i.ToApiManagementServiceSkuPropertiesPtrOutputWithContext(context.Background())
}

func (i *apiManagementServiceSkuPropertiesPtrType) ToApiManagementServiceSkuPropertiesPtrOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceSkuPropertiesPtrOutput)
}

type ApiManagementServiceSkuPropertiesOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceSkuPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementServiceSkuProperties)(nil)).Elem()
}

func (o ApiManagementServiceSkuPropertiesOutput) ToApiManagementServiceSkuPropertiesOutput() ApiManagementServiceSkuPropertiesOutput {
	return o
}

func (o ApiManagementServiceSkuPropertiesOutput) ToApiManagementServiceSkuPropertiesOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesOutput {
	return o
}

func (o ApiManagementServiceSkuPropertiesOutput) ToApiManagementServiceSkuPropertiesPtrOutput() ApiManagementServiceSkuPropertiesPtrOutput {
	return o.ToApiManagementServiceSkuPropertiesPtrOutputWithContext(context.Background())
}

func (o ApiManagementServiceSkuPropertiesOutput) ToApiManagementServiceSkuPropertiesPtrOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiManagementServiceSkuProperties) *ApiManagementServiceSkuProperties {
		return &v
	}).(ApiManagementServiceSkuPropertiesPtrOutput)
}

func (o ApiManagementServiceSkuPropertiesOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApiManagementServiceSkuProperties) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ApiManagementServiceSkuPropertiesOutput) Name() SkuTypeOutput {
	return o.ApplyT(func(v ApiManagementServiceSkuProperties) SkuType { return v.Name }).(SkuTypeOutput)
}

type ApiManagementServiceSkuPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceSkuPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementServiceSkuProperties)(nil)).Elem()
}

func (o ApiManagementServiceSkuPropertiesPtrOutput) ToApiManagementServiceSkuPropertiesPtrOutput() ApiManagementServiceSkuPropertiesPtrOutput {
	return o
}

func (o ApiManagementServiceSkuPropertiesPtrOutput) ToApiManagementServiceSkuPropertiesPtrOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesPtrOutput {
	return o
}

func (o ApiManagementServiceSkuPropertiesPtrOutput) Elem() ApiManagementServiceSkuPropertiesOutput {
	return o.ApplyT(func(v *ApiManagementServiceSkuProperties) ApiManagementServiceSkuProperties {
		if v != nil {
			return *v
		}
		var ret ApiManagementServiceSkuProperties
		return ret
	}).(ApiManagementServiceSkuPropertiesOutput)
}

func (o ApiManagementServiceSkuPropertiesPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApiManagementServiceSkuProperties) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ApiManagementServiceSkuPropertiesPtrOutput) Name() SkuTypePtrOutput {
	return o.ApplyT(func(v *ApiManagementServiceSkuProperties) *SkuType {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(SkuTypePtrOutput)
}

type ApiManagementServiceSkuPropertiesResponse struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}





type ApiManagementServiceSkuPropertiesResponseInput interface {
	pulumi.Input

	ToApiManagementServiceSkuPropertiesResponseOutput() ApiManagementServiceSkuPropertiesResponseOutput
	ToApiManagementServiceSkuPropertiesResponseOutputWithContext(context.Context) ApiManagementServiceSkuPropertiesResponseOutput
}

type ApiManagementServiceSkuPropertiesResponseArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (ApiManagementServiceSkuPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementServiceSkuPropertiesResponse)(nil)).Elem()
}

func (i ApiManagementServiceSkuPropertiesResponseArgs) ToApiManagementServiceSkuPropertiesResponseOutput() ApiManagementServiceSkuPropertiesResponseOutput {
	return i.ToApiManagementServiceSkuPropertiesResponseOutputWithContext(context.Background())
}

func (i ApiManagementServiceSkuPropertiesResponseArgs) ToApiManagementServiceSkuPropertiesResponseOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceSkuPropertiesResponseOutput)
}

func (i ApiManagementServiceSkuPropertiesResponseArgs) ToApiManagementServiceSkuPropertiesResponsePtrOutput() ApiManagementServiceSkuPropertiesResponsePtrOutput {
	return i.ToApiManagementServiceSkuPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ApiManagementServiceSkuPropertiesResponseArgs) ToApiManagementServiceSkuPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceSkuPropertiesResponseOutput).ToApiManagementServiceSkuPropertiesResponsePtrOutputWithContext(ctx)
}









type ApiManagementServiceSkuPropertiesResponsePtrInput interface {
	pulumi.Input

	ToApiManagementServiceSkuPropertiesResponsePtrOutput() ApiManagementServiceSkuPropertiesResponsePtrOutput
	ToApiManagementServiceSkuPropertiesResponsePtrOutputWithContext(context.Context) ApiManagementServiceSkuPropertiesResponsePtrOutput
}

type apiManagementServiceSkuPropertiesResponsePtrType ApiManagementServiceSkuPropertiesResponseArgs

func ApiManagementServiceSkuPropertiesResponsePtr(v *ApiManagementServiceSkuPropertiesResponseArgs) ApiManagementServiceSkuPropertiesResponsePtrInput {
	return (*apiManagementServiceSkuPropertiesResponsePtrType)(v)
}

func (*apiManagementServiceSkuPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementServiceSkuPropertiesResponse)(nil)).Elem()
}

func (i *apiManagementServiceSkuPropertiesResponsePtrType) ToApiManagementServiceSkuPropertiesResponsePtrOutput() ApiManagementServiceSkuPropertiesResponsePtrOutput {
	return i.ToApiManagementServiceSkuPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *apiManagementServiceSkuPropertiesResponsePtrType) ToApiManagementServiceSkuPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceSkuPropertiesResponsePtrOutput)
}

type ApiManagementServiceSkuPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceSkuPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementServiceSkuPropertiesResponse)(nil)).Elem()
}

func (o ApiManagementServiceSkuPropertiesResponseOutput) ToApiManagementServiceSkuPropertiesResponseOutput() ApiManagementServiceSkuPropertiesResponseOutput {
	return o
}

func (o ApiManagementServiceSkuPropertiesResponseOutput) ToApiManagementServiceSkuPropertiesResponseOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesResponseOutput {
	return o
}

func (o ApiManagementServiceSkuPropertiesResponseOutput) ToApiManagementServiceSkuPropertiesResponsePtrOutput() ApiManagementServiceSkuPropertiesResponsePtrOutput {
	return o.ToApiManagementServiceSkuPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ApiManagementServiceSkuPropertiesResponseOutput) ToApiManagementServiceSkuPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiManagementServiceSkuPropertiesResponse) *ApiManagementServiceSkuPropertiesResponse {
		return &v
	}).(ApiManagementServiceSkuPropertiesResponsePtrOutput)
}

func (o ApiManagementServiceSkuPropertiesResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApiManagementServiceSkuPropertiesResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ApiManagementServiceSkuPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApiManagementServiceSkuPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ApiManagementServiceSkuPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceSkuPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementServiceSkuPropertiesResponse)(nil)).Elem()
}

func (o ApiManagementServiceSkuPropertiesResponsePtrOutput) ToApiManagementServiceSkuPropertiesResponsePtrOutput() ApiManagementServiceSkuPropertiesResponsePtrOutput {
	return o
}

func (o ApiManagementServiceSkuPropertiesResponsePtrOutput) ToApiManagementServiceSkuPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiManagementServiceSkuPropertiesResponsePtrOutput {
	return o
}

func (o ApiManagementServiceSkuPropertiesResponsePtrOutput) Elem() ApiManagementServiceSkuPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiManagementServiceSkuPropertiesResponse) ApiManagementServiceSkuPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ApiManagementServiceSkuPropertiesResponse
		return ret
	}).(ApiManagementServiceSkuPropertiesResponseOutput)
}

func (o ApiManagementServiceSkuPropertiesResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApiManagementServiceSkuPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ApiManagementServiceSkuPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementServiceSkuPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type CertificateInformation struct {
	Expiry     string `pulumi:"expiry"`
	Subject    string `pulumi:"subject"`
	Thumbprint string `pulumi:"thumbprint"`
}





type CertificateInformationInput interface {
	pulumi.Input

	ToCertificateInformationOutput() CertificateInformationOutput
	ToCertificateInformationOutputWithContext(context.Context) CertificateInformationOutput
}

type CertificateInformationArgs struct {
	Expiry     pulumi.StringInput `pulumi:"expiry"`
	Subject    pulumi.StringInput `pulumi:"subject"`
	Thumbprint pulumi.StringInput `pulumi:"thumbprint"`
}

func (CertificateInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateInformation)(nil)).Elem()
}

func (i CertificateInformationArgs) ToCertificateInformationOutput() CertificateInformationOutput {
	return i.ToCertificateInformationOutputWithContext(context.Background())
}

func (i CertificateInformationArgs) ToCertificateInformationOutputWithContext(ctx context.Context) CertificateInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateInformationOutput)
}

type CertificateInformationOutput struct{ *pulumi.OutputState }

func (CertificateInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateInformation)(nil)).Elem()
}

func (o CertificateInformationOutput) ToCertificateInformationOutput() CertificateInformationOutput {
	return o
}

func (o CertificateInformationOutput) ToCertificateInformationOutputWithContext(ctx context.Context) CertificateInformationOutput {
	return o
}

func (o CertificateInformationOutput) Expiry() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateInformation) string { return v.Expiry }).(pulumi.StringOutput)
}

func (o CertificateInformationOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateInformation) string { return v.Subject }).(pulumi.StringOutput)
}

func (o CertificateInformationOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateInformation) string { return v.Thumbprint }).(pulumi.StringOutput)
}

type CertificateInformationResponse struct {
	Expiry     string `pulumi:"expiry"`
	Subject    string `pulumi:"subject"`
	Thumbprint string `pulumi:"thumbprint"`
}





type CertificateInformationResponseInput interface {
	pulumi.Input

	ToCertificateInformationResponseOutput() CertificateInformationResponseOutput
	ToCertificateInformationResponseOutputWithContext(context.Context) CertificateInformationResponseOutput
}

type CertificateInformationResponseArgs struct {
	Expiry     pulumi.StringInput `pulumi:"expiry"`
	Subject    pulumi.StringInput `pulumi:"subject"`
	Thumbprint pulumi.StringInput `pulumi:"thumbprint"`
}

func (CertificateInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateInformationResponse)(nil)).Elem()
}

func (i CertificateInformationResponseArgs) ToCertificateInformationResponseOutput() CertificateInformationResponseOutput {
	return i.ToCertificateInformationResponseOutputWithContext(context.Background())
}

func (i CertificateInformationResponseArgs) ToCertificateInformationResponseOutputWithContext(ctx context.Context) CertificateInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateInformationResponseOutput)
}

type CertificateInformationResponseOutput struct{ *pulumi.OutputState }

func (CertificateInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateInformationResponse)(nil)).Elem()
}

func (o CertificateInformationResponseOutput) ToCertificateInformationResponseOutput() CertificateInformationResponseOutput {
	return o
}

func (o CertificateInformationResponseOutput) ToCertificateInformationResponseOutputWithContext(ctx context.Context) CertificateInformationResponseOutput {
	return o
}

func (o CertificateInformationResponseOutput) Expiry() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateInformationResponse) string { return v.Expiry }).(pulumi.StringOutput)
}

func (o CertificateInformationResponseOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateInformationResponse) string { return v.Subject }).(pulumi.StringOutput)
}

func (o CertificateInformationResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateInformationResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

type HostnameConfiguration struct {
	Certificate CertificateInformation `pulumi:"certificate"`
	Hostname    string                 `pulumi:"hostname"`
	Type        HostnameType           `pulumi:"type"`
}





type HostnameConfigurationInput interface {
	pulumi.Input

	ToHostnameConfigurationOutput() HostnameConfigurationOutput
	ToHostnameConfigurationOutputWithContext(context.Context) HostnameConfigurationOutput
}

type HostnameConfigurationArgs struct {
	Certificate CertificateInformationInput `pulumi:"certificate"`
	Hostname    pulumi.StringInput          `pulumi:"hostname"`
	Type        HostnameTypeInput           `pulumi:"type"`
}

func (HostnameConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostnameConfiguration)(nil)).Elem()
}

func (i HostnameConfigurationArgs) ToHostnameConfigurationOutput() HostnameConfigurationOutput {
	return i.ToHostnameConfigurationOutputWithContext(context.Background())
}

func (i HostnameConfigurationArgs) ToHostnameConfigurationOutputWithContext(ctx context.Context) HostnameConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostnameConfigurationOutput)
}





type HostnameConfigurationArrayInput interface {
	pulumi.Input

	ToHostnameConfigurationArrayOutput() HostnameConfigurationArrayOutput
	ToHostnameConfigurationArrayOutputWithContext(context.Context) HostnameConfigurationArrayOutput
}

type HostnameConfigurationArray []HostnameConfigurationInput

func (HostnameConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostnameConfiguration)(nil)).Elem()
}

func (i HostnameConfigurationArray) ToHostnameConfigurationArrayOutput() HostnameConfigurationArrayOutput {
	return i.ToHostnameConfigurationArrayOutputWithContext(context.Background())
}

func (i HostnameConfigurationArray) ToHostnameConfigurationArrayOutputWithContext(ctx context.Context) HostnameConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostnameConfigurationArrayOutput)
}

type HostnameConfigurationOutput struct{ *pulumi.OutputState }

func (HostnameConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostnameConfiguration)(nil)).Elem()
}

func (o HostnameConfigurationOutput) ToHostnameConfigurationOutput() HostnameConfigurationOutput {
	return o
}

func (o HostnameConfigurationOutput) ToHostnameConfigurationOutputWithContext(ctx context.Context) HostnameConfigurationOutput {
	return o
}

func (o HostnameConfigurationOutput) Certificate() CertificateInformationOutput {
	return o.ApplyT(func(v HostnameConfiguration) CertificateInformation { return v.Certificate }).(CertificateInformationOutput)
}

func (o HostnameConfigurationOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v HostnameConfiguration) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o HostnameConfigurationOutput) Type() HostnameTypeOutput {
	return o.ApplyT(func(v HostnameConfiguration) HostnameType { return v.Type }).(HostnameTypeOutput)
}

type HostnameConfigurationArrayOutput struct{ *pulumi.OutputState }

func (HostnameConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostnameConfiguration)(nil)).Elem()
}

func (o HostnameConfigurationArrayOutput) ToHostnameConfigurationArrayOutput() HostnameConfigurationArrayOutput {
	return o
}

func (o HostnameConfigurationArrayOutput) ToHostnameConfigurationArrayOutputWithContext(ctx context.Context) HostnameConfigurationArrayOutput {
	return o
}

func (o HostnameConfigurationArrayOutput) Index(i pulumi.IntInput) HostnameConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostnameConfiguration {
		return vs[0].([]HostnameConfiguration)[vs[1].(int)]
	}).(HostnameConfigurationOutput)
}

type HostnameConfigurationResponse struct {
	Certificate CertificateInformationResponse `pulumi:"certificate"`
	Hostname    string                         `pulumi:"hostname"`
	Type        string                         `pulumi:"type"`
}





type HostnameConfigurationResponseInput interface {
	pulumi.Input

	ToHostnameConfigurationResponseOutput() HostnameConfigurationResponseOutput
	ToHostnameConfigurationResponseOutputWithContext(context.Context) HostnameConfigurationResponseOutput
}

type HostnameConfigurationResponseArgs struct {
	Certificate CertificateInformationResponseInput `pulumi:"certificate"`
	Hostname    pulumi.StringInput                  `pulumi:"hostname"`
	Type        pulumi.StringInput                  `pulumi:"type"`
}

func (HostnameConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostnameConfigurationResponse)(nil)).Elem()
}

func (i HostnameConfigurationResponseArgs) ToHostnameConfigurationResponseOutput() HostnameConfigurationResponseOutput {
	return i.ToHostnameConfigurationResponseOutputWithContext(context.Background())
}

func (i HostnameConfigurationResponseArgs) ToHostnameConfigurationResponseOutputWithContext(ctx context.Context) HostnameConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostnameConfigurationResponseOutput)
}





type HostnameConfigurationResponseArrayInput interface {
	pulumi.Input

	ToHostnameConfigurationResponseArrayOutput() HostnameConfigurationResponseArrayOutput
	ToHostnameConfigurationResponseArrayOutputWithContext(context.Context) HostnameConfigurationResponseArrayOutput
}

type HostnameConfigurationResponseArray []HostnameConfigurationResponseInput

func (HostnameConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostnameConfigurationResponse)(nil)).Elem()
}

func (i HostnameConfigurationResponseArray) ToHostnameConfigurationResponseArrayOutput() HostnameConfigurationResponseArrayOutput {
	return i.ToHostnameConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i HostnameConfigurationResponseArray) ToHostnameConfigurationResponseArrayOutputWithContext(ctx context.Context) HostnameConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostnameConfigurationResponseArrayOutput)
}

type HostnameConfigurationResponseOutput struct{ *pulumi.OutputState }

func (HostnameConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostnameConfigurationResponse)(nil)).Elem()
}

func (o HostnameConfigurationResponseOutput) ToHostnameConfigurationResponseOutput() HostnameConfigurationResponseOutput {
	return o
}

func (o HostnameConfigurationResponseOutput) ToHostnameConfigurationResponseOutputWithContext(ctx context.Context) HostnameConfigurationResponseOutput {
	return o
}

func (o HostnameConfigurationResponseOutput) Certificate() CertificateInformationResponseOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) CertificateInformationResponse { return v.Certificate }).(CertificateInformationResponseOutput)
}

func (o HostnameConfigurationResponseOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) string { return v.Hostname }).(pulumi.StringOutput)
}

func (o HostnameConfigurationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type HostnameConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (HostnameConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostnameConfigurationResponse)(nil)).Elem()
}

func (o HostnameConfigurationResponseArrayOutput) ToHostnameConfigurationResponseArrayOutput() HostnameConfigurationResponseArrayOutput {
	return o
}

func (o HostnameConfigurationResponseArrayOutput) ToHostnameConfigurationResponseArrayOutputWithContext(ctx context.Context) HostnameConfigurationResponseArrayOutput {
	return o
}

func (o HostnameConfigurationResponseArrayOutput) Index(i pulumi.IntInput) HostnameConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostnameConfigurationResponse {
		return vs[0].([]HostnameConfigurationResponse)[vs[1].(int)]
	}).(HostnameConfigurationResponseOutput)
}

type VirtualNetworkConfiguration struct {
	Location         *string `pulumi:"location"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
}





type VirtualNetworkConfigurationInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput
	ToVirtualNetworkConfigurationOutputWithContext(context.Context) VirtualNetworkConfigurationOutput
}

type VirtualNetworkConfigurationArgs struct {
	Location         pulumi.StringPtrInput `pulumi:"location"`
	SubnetResourceId pulumi.StringPtrInput `pulumi:"subnetResourceId"`
}

func (VirtualNetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfiguration)(nil)).Elem()
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput {
	return i.ToVirtualNetworkConfigurationOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationOutputWithContext(ctx context.Context) VirtualNetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationOutput)
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return i.ToVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigurationArgs) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationOutput).ToVirtualNetworkConfigurationPtrOutputWithContext(ctx)
}









type VirtualNetworkConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput
	ToVirtualNetworkConfigurationPtrOutputWithContext(context.Context) VirtualNetworkConfigurationPtrOutput
}

type virtualNetworkConfigurationPtrType VirtualNetworkConfigurationArgs

func VirtualNetworkConfigurationPtr(v *VirtualNetworkConfigurationArgs) VirtualNetworkConfigurationPtrInput {
	return (*virtualNetworkConfigurationPtrType)(v)
}

func (*virtualNetworkConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfiguration)(nil)).Elem()
}

func (i *virtualNetworkConfigurationPtrType) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return i.ToVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkConfigurationPtrType) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationPtrOutput)
}

type VirtualNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfiguration)(nil)).Elem()
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput {
	return o
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationOutputWithContext(ctx context.Context) VirtualNetworkConfigurationOutput {
	return o
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return o.ToVirtualNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkConfigurationOutput) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkConfiguration) *VirtualNetworkConfiguration {
		return &v
	}).(VirtualNetworkConfigurationPtrOutput)
}

func (o VirtualNetworkConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfiguration) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfiguration) *string { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfiguration)(nil)).Elem()
}

func (o VirtualNetworkConfigurationPtrOutput) ToVirtualNetworkConfigurationPtrOutput() VirtualNetworkConfigurationPtrOutput {
	return o
}

func (o VirtualNetworkConfigurationPtrOutput) ToVirtualNetworkConfigurationPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationPtrOutput {
	return o
}

func (o VirtualNetworkConfigurationPtrOutput) Elem() VirtualNetworkConfigurationOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) VirtualNetworkConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfiguration
		return ret
	}).(VirtualNetworkConfigurationOutput)
}

func (o VirtualNetworkConfigurationPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationPtrOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SubnetResourceId
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigurationResponse struct {
	Location         *string `pulumi:"location"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	Subnetname       string  `pulumi:"subnetname"`
	Vnetid           string  `pulumi:"vnetid"`
}





type VirtualNetworkConfigurationResponseInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationResponseOutput() VirtualNetworkConfigurationResponseOutput
	ToVirtualNetworkConfigurationResponseOutputWithContext(context.Context) VirtualNetworkConfigurationResponseOutput
}

type VirtualNetworkConfigurationResponseArgs struct {
	Location         pulumi.StringPtrInput `pulumi:"location"`
	SubnetResourceId pulumi.StringPtrInput `pulumi:"subnetResourceId"`
	Subnetname       pulumi.StringInput    `pulumi:"subnetname"`
	Vnetid           pulumi.StringInput    `pulumi:"vnetid"`
}

func (VirtualNetworkConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfigurationResponse)(nil)).Elem()
}

func (i VirtualNetworkConfigurationResponseArgs) ToVirtualNetworkConfigurationResponseOutput() VirtualNetworkConfigurationResponseOutput {
	return i.ToVirtualNetworkConfigurationResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigurationResponseArgs) ToVirtualNetworkConfigurationResponseOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationResponseOutput)
}

func (i VirtualNetworkConfigurationResponseArgs) ToVirtualNetworkConfigurationResponsePtrOutput() VirtualNetworkConfigurationResponsePtrOutput {
	return i.ToVirtualNetworkConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigurationResponseArgs) ToVirtualNetworkConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationResponseOutput).ToVirtualNetworkConfigurationResponsePtrOutputWithContext(ctx)
}









type VirtualNetworkConfigurationResponsePtrInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationResponsePtrOutput() VirtualNetworkConfigurationResponsePtrOutput
	ToVirtualNetworkConfigurationResponsePtrOutputWithContext(context.Context) VirtualNetworkConfigurationResponsePtrOutput
}

type virtualNetworkConfigurationResponsePtrType VirtualNetworkConfigurationResponseArgs

func VirtualNetworkConfigurationResponsePtr(v *VirtualNetworkConfigurationResponseArgs) VirtualNetworkConfigurationResponsePtrInput {
	return (*virtualNetworkConfigurationResponsePtrType)(v)
}

func (*virtualNetworkConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfigurationResponse)(nil)).Elem()
}

func (i *virtualNetworkConfigurationResponsePtrType) ToVirtualNetworkConfigurationResponsePtrOutput() VirtualNetworkConfigurationResponsePtrOutput {
	return i.ToVirtualNetworkConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkConfigurationResponsePtrType) ToVirtualNetworkConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigurationResponsePtrOutput)
}

type VirtualNetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigurationResponseOutput) ToVirtualNetworkConfigurationResponseOutput() VirtualNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkConfigurationResponseOutput) ToVirtualNetworkConfigurationResponseOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkConfigurationResponseOutput) ToVirtualNetworkConfigurationResponsePtrOutput() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ToVirtualNetworkConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkConfigurationResponseOutput) ToVirtualNetworkConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkConfigurationResponse) *VirtualNetworkConfigurationResponse {
		return &v
	}).(VirtualNetworkConfigurationResponsePtrOutput)
}

func (o VirtualNetworkConfigurationResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationResponseOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) *string { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationResponseOutput) Subnetname() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) string { return v.Subnetname }).(pulumi.StringOutput)
}

func (o VirtualNetworkConfigurationResponseOutput) Vnetid() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkConfigurationResponse) string { return v.Vnetid }).(pulumi.StringOutput)
}

type VirtualNetworkConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigurationResponsePtrOutput) ToVirtualNetworkConfigurationResponsePtrOutput() VirtualNetworkConfigurationResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigurationResponsePtrOutput) ToVirtualNetworkConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigurationResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigurationResponsePtrOutput) Elem() VirtualNetworkConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) VirtualNetworkConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfigurationResponse
		return ret
	}).(VirtualNetworkConfigurationResponseOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) Subnetname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Subnetname
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigurationResponsePtrOutput) Vnetid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Vnetid
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdditionalRegionOutput{})
	pulumi.RegisterOutputType(AdditionalRegionArrayOutput{})
	pulumi.RegisterOutputType(AdditionalRegionResponseOutput{})
	pulumi.RegisterOutputType(AdditionalRegionResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CertificateInformationOutput{})
	pulumi.RegisterOutputType(CertificateInformationResponseOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationArrayOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationResponseOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponsePtrOutput{})
}
