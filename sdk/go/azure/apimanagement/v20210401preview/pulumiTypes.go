


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdditionalLocation struct {
	DisableGateway              *bool                             `pulumi:"disableGateway"`
	Location                    string                            `pulumi:"location"`
	PublicIpAddressId           *string                           `pulumi:"publicIpAddressId"`
	Sku                         ApiManagementServiceSkuProperties `pulumi:"sku"`
	VirtualNetworkConfiguration *VirtualNetworkConfiguration      `pulumi:"virtualNetworkConfiguration"`
	Zones                       []string                          `pulumi:"zones"`
}


func (val *AdditionalLocation) Defaults() *AdditionalLocation {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DisableGateway) {
		disableGateway_ := false
		tmp.DisableGateway = &disableGateway_
	}
	return &tmp
}





type AdditionalLocationInput interface {
	pulumi.Input

	ToAdditionalLocationOutput() AdditionalLocationOutput
	ToAdditionalLocationOutputWithContext(context.Context) AdditionalLocationOutput
}

type AdditionalLocationArgs struct {
	DisableGateway              pulumi.BoolPtrInput                    `pulumi:"disableGateway"`
	Location                    pulumi.StringInput                     `pulumi:"location"`
	PublicIpAddressId           pulumi.StringPtrInput                  `pulumi:"publicIpAddressId"`
	Sku                         ApiManagementServiceSkuPropertiesInput `pulumi:"sku"`
	VirtualNetworkConfiguration VirtualNetworkConfigurationPtrInput    `pulumi:"virtualNetworkConfiguration"`
	Zones                       pulumi.StringArrayInput                `pulumi:"zones"`
}

func (AdditionalLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalLocation)(nil)).Elem()
}

func (i AdditionalLocationArgs) ToAdditionalLocationOutput() AdditionalLocationOutput {
	return i.ToAdditionalLocationOutputWithContext(context.Background())
}

func (i AdditionalLocationArgs) ToAdditionalLocationOutputWithContext(ctx context.Context) AdditionalLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalLocationOutput)
}





type AdditionalLocationArrayInput interface {
	pulumi.Input

	ToAdditionalLocationArrayOutput() AdditionalLocationArrayOutput
	ToAdditionalLocationArrayOutputWithContext(context.Context) AdditionalLocationArrayOutput
}

type AdditionalLocationArray []AdditionalLocationInput

func (AdditionalLocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalLocation)(nil)).Elem()
}

func (i AdditionalLocationArray) ToAdditionalLocationArrayOutput() AdditionalLocationArrayOutput {
	return i.ToAdditionalLocationArrayOutputWithContext(context.Background())
}

func (i AdditionalLocationArray) ToAdditionalLocationArrayOutputWithContext(ctx context.Context) AdditionalLocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalLocationArrayOutput)
}

type AdditionalLocationOutput struct{ *pulumi.OutputState }

func (AdditionalLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalLocation)(nil)).Elem()
}

func (o AdditionalLocationOutput) ToAdditionalLocationOutput() AdditionalLocationOutput {
	return o
}

func (o AdditionalLocationOutput) ToAdditionalLocationOutputWithContext(ctx context.Context) AdditionalLocationOutput {
	return o
}

func (o AdditionalLocationOutput) DisableGateway() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AdditionalLocation) *bool { return v.DisableGateway }).(pulumi.BoolPtrOutput)
}

func (o AdditionalLocationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalLocation) string { return v.Location }).(pulumi.StringOutput)
}

func (o AdditionalLocationOutput) PublicIpAddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalLocation) *string { return v.PublicIpAddressId }).(pulumi.StringPtrOutput)
}

func (o AdditionalLocationOutput) Sku() ApiManagementServiceSkuPropertiesOutput {
	return o.ApplyT(func(v AdditionalLocation) ApiManagementServiceSkuProperties { return v.Sku }).(ApiManagementServiceSkuPropertiesOutput)
}

func (o AdditionalLocationOutput) VirtualNetworkConfiguration() VirtualNetworkConfigurationPtrOutput {
	return o.ApplyT(func(v AdditionalLocation) *VirtualNetworkConfiguration { return v.VirtualNetworkConfiguration }).(VirtualNetworkConfigurationPtrOutput)
}

func (o AdditionalLocationOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalLocation) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

type AdditionalLocationArrayOutput struct{ *pulumi.OutputState }

func (AdditionalLocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalLocation)(nil)).Elem()
}

func (o AdditionalLocationArrayOutput) ToAdditionalLocationArrayOutput() AdditionalLocationArrayOutput {
	return o
}

func (o AdditionalLocationArrayOutput) ToAdditionalLocationArrayOutputWithContext(ctx context.Context) AdditionalLocationArrayOutput {
	return o
}

func (o AdditionalLocationArrayOutput) Index(i pulumi.IntInput) AdditionalLocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalLocation {
		return vs[0].([]AdditionalLocation)[vs[1].(int)]
	}).(AdditionalLocationOutput)
}

type AdditionalLocationResponse struct {
	DisableGateway              *bool                                     `pulumi:"disableGateway"`
	GatewayRegionalUrl          string                                    `pulumi:"gatewayRegionalUrl"`
	Location                    string                                    `pulumi:"location"`
	PlatformVersion             string                                    `pulumi:"platformVersion"`
	PrivateIPAddresses          []string                                  `pulumi:"privateIPAddresses"`
	PublicIPAddresses           []string                                  `pulumi:"publicIPAddresses"`
	PublicIpAddressId           *string                                   `pulumi:"publicIpAddressId"`
	Sku                         ApiManagementServiceSkuPropertiesResponse `pulumi:"sku"`
	VirtualNetworkConfiguration *VirtualNetworkConfigurationResponse      `pulumi:"virtualNetworkConfiguration"`
	Zones                       []string                                  `pulumi:"zones"`
}


func (val *AdditionalLocationResponse) Defaults() *AdditionalLocationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DisableGateway) {
		disableGateway_ := false
		tmp.DisableGateway = &disableGateway_
	}
	return &tmp
}

type AdditionalLocationResponseOutput struct{ *pulumi.OutputState }

func (AdditionalLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalLocationResponse)(nil)).Elem()
}

func (o AdditionalLocationResponseOutput) ToAdditionalLocationResponseOutput() AdditionalLocationResponseOutput {
	return o
}

func (o AdditionalLocationResponseOutput) ToAdditionalLocationResponseOutputWithContext(ctx context.Context) AdditionalLocationResponseOutput {
	return o
}

func (o AdditionalLocationResponseOutput) DisableGateway() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) *bool { return v.DisableGateway }).(pulumi.BoolPtrOutput)
}

func (o AdditionalLocationResponseOutput) GatewayRegionalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) string { return v.GatewayRegionalUrl }).(pulumi.StringOutput)
}

func (o AdditionalLocationResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o AdditionalLocationResponseOutput) PlatformVersion() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) string { return v.PlatformVersion }).(pulumi.StringOutput)
}

func (o AdditionalLocationResponseOutput) PrivateIPAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) []string { return v.PrivateIPAddresses }).(pulumi.StringArrayOutput)
}

func (o AdditionalLocationResponseOutput) PublicIPAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) []string { return v.PublicIPAddresses }).(pulumi.StringArrayOutput)
}

func (o AdditionalLocationResponseOutput) PublicIpAddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) *string { return v.PublicIpAddressId }).(pulumi.StringPtrOutput)
}

func (o AdditionalLocationResponseOutput) Sku() ApiManagementServiceSkuPropertiesResponseOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) ApiManagementServiceSkuPropertiesResponse { return v.Sku }).(ApiManagementServiceSkuPropertiesResponseOutput)
}

func (o AdditionalLocationResponseOutput) VirtualNetworkConfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) *VirtualNetworkConfigurationResponse {
		return v.VirtualNetworkConfiguration
	}).(VirtualNetworkConfigurationResponsePtrOutput)
}

func (o AdditionalLocationResponseOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

type AdditionalLocationResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalLocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalLocationResponse)(nil)).Elem()
}

func (o AdditionalLocationResponseArrayOutput) ToAdditionalLocationResponseArrayOutput() AdditionalLocationResponseArrayOutput {
	return o
}

func (o AdditionalLocationResponseArrayOutput) ToAdditionalLocationResponseArrayOutputWithContext(ctx context.Context) AdditionalLocationResponseArrayOutput {
	return o
}

func (o AdditionalLocationResponseArrayOutput) Index(i pulumi.IntInput) AdditionalLocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalLocationResponse {
		return vs[0].([]AdditionalLocationResponse)[vs[1].(int)]
	}).(AdditionalLocationResponseOutput)
}

type ApiContactInformation struct {
	Email *string `pulumi:"email"`
	Name  *string `pulumi:"name"`
	Url   *string `pulumi:"url"`
}





type ApiContactInformationInput interface {
	pulumi.Input

	ToApiContactInformationOutput() ApiContactInformationOutput
	ToApiContactInformationOutputWithContext(context.Context) ApiContactInformationOutput
}

type ApiContactInformationArgs struct {
	Email pulumi.StringPtrInput `pulumi:"email"`
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Url   pulumi.StringPtrInput `pulumi:"url"`
}

func (ApiContactInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiContactInformation)(nil)).Elem()
}

func (i ApiContactInformationArgs) ToApiContactInformationOutput() ApiContactInformationOutput {
	return i.ToApiContactInformationOutputWithContext(context.Background())
}

func (i ApiContactInformationArgs) ToApiContactInformationOutputWithContext(ctx context.Context) ApiContactInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiContactInformationOutput)
}

func (i ApiContactInformationArgs) ToApiContactInformationPtrOutput() ApiContactInformationPtrOutput {
	return i.ToApiContactInformationPtrOutputWithContext(context.Background())
}

func (i ApiContactInformationArgs) ToApiContactInformationPtrOutputWithContext(ctx context.Context) ApiContactInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiContactInformationOutput).ToApiContactInformationPtrOutputWithContext(ctx)
}









type ApiContactInformationPtrInput interface {
	pulumi.Input

	ToApiContactInformationPtrOutput() ApiContactInformationPtrOutput
	ToApiContactInformationPtrOutputWithContext(context.Context) ApiContactInformationPtrOutput
}

type apiContactInformationPtrType ApiContactInformationArgs

func ApiContactInformationPtr(v *ApiContactInformationArgs) ApiContactInformationPtrInput {
	return (*apiContactInformationPtrType)(v)
}

func (*apiContactInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiContactInformation)(nil)).Elem()
}

func (i *apiContactInformationPtrType) ToApiContactInformationPtrOutput() ApiContactInformationPtrOutput {
	return i.ToApiContactInformationPtrOutputWithContext(context.Background())
}

func (i *apiContactInformationPtrType) ToApiContactInformationPtrOutputWithContext(ctx context.Context) ApiContactInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiContactInformationPtrOutput)
}

type ApiContactInformationOutput struct{ *pulumi.OutputState }

func (ApiContactInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiContactInformation)(nil)).Elem()
}

func (o ApiContactInformationOutput) ToApiContactInformationOutput() ApiContactInformationOutput {
	return o
}

func (o ApiContactInformationOutput) ToApiContactInformationOutputWithContext(ctx context.Context) ApiContactInformationOutput {
	return o
}

func (o ApiContactInformationOutput) ToApiContactInformationPtrOutput() ApiContactInformationPtrOutput {
	return o.ToApiContactInformationPtrOutputWithContext(context.Background())
}

func (o ApiContactInformationOutput) ToApiContactInformationPtrOutputWithContext(ctx context.Context) ApiContactInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiContactInformation) *ApiContactInformation {
		return &v
	}).(ApiContactInformationPtrOutput)
}

func (o ApiContactInformationOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiContactInformation) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o ApiContactInformationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiContactInformation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiContactInformationOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiContactInformation) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiContactInformationPtrOutput struct{ *pulumi.OutputState }

func (ApiContactInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiContactInformation)(nil)).Elem()
}

func (o ApiContactInformationPtrOutput) ToApiContactInformationPtrOutput() ApiContactInformationPtrOutput {
	return o
}

func (o ApiContactInformationPtrOutput) ToApiContactInformationPtrOutputWithContext(ctx context.Context) ApiContactInformationPtrOutput {
	return o
}

func (o ApiContactInformationPtrOutput) Elem() ApiContactInformationOutput {
	return o.ApplyT(func(v *ApiContactInformation) ApiContactInformation {
		if v != nil {
			return *v
		}
		var ret ApiContactInformation
		return ret
	}).(ApiContactInformationOutput)
}

func (o ApiContactInformationPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiContactInformation) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o ApiContactInformationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiContactInformation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiContactInformationPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiContactInformation) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApiContactInformationResponse struct {
	Email *string `pulumi:"email"`
	Name  *string `pulumi:"name"`
	Url   *string `pulumi:"url"`
}

type ApiContactInformationResponseOutput struct{ *pulumi.OutputState }

func (ApiContactInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiContactInformationResponse)(nil)).Elem()
}

func (o ApiContactInformationResponseOutput) ToApiContactInformationResponseOutput() ApiContactInformationResponseOutput {
	return o
}

func (o ApiContactInformationResponseOutput) ToApiContactInformationResponseOutputWithContext(ctx context.Context) ApiContactInformationResponseOutput {
	return o
}

func (o ApiContactInformationResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiContactInformationResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o ApiContactInformationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiContactInformationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiContactInformationResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiContactInformationResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiContactInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiContactInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiContactInformationResponse)(nil)).Elem()
}

func (o ApiContactInformationResponsePtrOutput) ToApiContactInformationResponsePtrOutput() ApiContactInformationResponsePtrOutput {
	return o
}

func (o ApiContactInformationResponsePtrOutput) ToApiContactInformationResponsePtrOutputWithContext(ctx context.Context) ApiContactInformationResponsePtrOutput {
	return o
}

func (o ApiContactInformationResponsePtrOutput) Elem() ApiContactInformationResponseOutput {
	return o.ApplyT(func(v *ApiContactInformationResponse) ApiContactInformationResponse {
		if v != nil {
			return *v
		}
		var ret ApiContactInformationResponse
		return ret
	}).(ApiContactInformationResponseOutput)
}

func (o ApiContactInformationResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiContactInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o ApiContactInformationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiContactInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiContactInformationResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiContactInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApiCreateOrUpdatePropertiesWsdlSelector struct {
	WsdlEndpointName *string `pulumi:"wsdlEndpointName"`
	WsdlServiceName  *string `pulumi:"wsdlServiceName"`
}





type ApiCreateOrUpdatePropertiesWsdlSelectorInput interface {
	pulumi.Input

	ToApiCreateOrUpdatePropertiesWsdlSelectorOutput() ApiCreateOrUpdatePropertiesWsdlSelectorOutput
	ToApiCreateOrUpdatePropertiesWsdlSelectorOutputWithContext(context.Context) ApiCreateOrUpdatePropertiesWsdlSelectorOutput
}

type ApiCreateOrUpdatePropertiesWsdlSelectorArgs struct {
	WsdlEndpointName pulumi.StringPtrInput `pulumi:"wsdlEndpointName"`
	WsdlServiceName  pulumi.StringPtrInput `pulumi:"wsdlServiceName"`
}

func (ApiCreateOrUpdatePropertiesWsdlSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiCreateOrUpdatePropertiesWsdlSelector)(nil)).Elem()
}

func (i ApiCreateOrUpdatePropertiesWsdlSelectorArgs) ToApiCreateOrUpdatePropertiesWsdlSelectorOutput() ApiCreateOrUpdatePropertiesWsdlSelectorOutput {
	return i.ToApiCreateOrUpdatePropertiesWsdlSelectorOutputWithContext(context.Background())
}

func (i ApiCreateOrUpdatePropertiesWsdlSelectorArgs) ToApiCreateOrUpdatePropertiesWsdlSelectorOutputWithContext(ctx context.Context) ApiCreateOrUpdatePropertiesWsdlSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiCreateOrUpdatePropertiesWsdlSelectorOutput)
}

func (i ApiCreateOrUpdatePropertiesWsdlSelectorArgs) ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput() ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput {
	return i.ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutputWithContext(context.Background())
}

func (i ApiCreateOrUpdatePropertiesWsdlSelectorArgs) ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutputWithContext(ctx context.Context) ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiCreateOrUpdatePropertiesWsdlSelectorOutput).ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutputWithContext(ctx)
}









type ApiCreateOrUpdatePropertiesWsdlSelectorPtrInput interface {
	pulumi.Input

	ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput() ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput
	ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutputWithContext(context.Context) ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput
}

type apiCreateOrUpdatePropertiesWsdlSelectorPtrType ApiCreateOrUpdatePropertiesWsdlSelectorArgs

func ApiCreateOrUpdatePropertiesWsdlSelectorPtr(v *ApiCreateOrUpdatePropertiesWsdlSelectorArgs) ApiCreateOrUpdatePropertiesWsdlSelectorPtrInput {
	return (*apiCreateOrUpdatePropertiesWsdlSelectorPtrType)(v)
}

func (*apiCreateOrUpdatePropertiesWsdlSelectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiCreateOrUpdatePropertiesWsdlSelector)(nil)).Elem()
}

func (i *apiCreateOrUpdatePropertiesWsdlSelectorPtrType) ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput() ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput {
	return i.ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutputWithContext(context.Background())
}

func (i *apiCreateOrUpdatePropertiesWsdlSelectorPtrType) ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutputWithContext(ctx context.Context) ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput)
}

type ApiCreateOrUpdatePropertiesWsdlSelectorOutput struct{ *pulumi.OutputState }

func (ApiCreateOrUpdatePropertiesWsdlSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiCreateOrUpdatePropertiesWsdlSelector)(nil)).Elem()
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorOutput) ToApiCreateOrUpdatePropertiesWsdlSelectorOutput() ApiCreateOrUpdatePropertiesWsdlSelectorOutput {
	return o
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorOutput) ToApiCreateOrUpdatePropertiesWsdlSelectorOutputWithContext(ctx context.Context) ApiCreateOrUpdatePropertiesWsdlSelectorOutput {
	return o
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorOutput) ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput() ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput {
	return o.ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutputWithContext(context.Background())
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorOutput) ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutputWithContext(ctx context.Context) ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiCreateOrUpdatePropertiesWsdlSelector) *ApiCreateOrUpdatePropertiesWsdlSelector {
		return &v
	}).(ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput)
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorOutput) WsdlEndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiCreateOrUpdatePropertiesWsdlSelector) *string { return v.WsdlEndpointName }).(pulumi.StringPtrOutput)
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorOutput) WsdlServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiCreateOrUpdatePropertiesWsdlSelector) *string { return v.WsdlServiceName }).(pulumi.StringPtrOutput)
}

type ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput struct{ *pulumi.OutputState }

func (ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiCreateOrUpdatePropertiesWsdlSelector)(nil)).Elem()
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput) ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput() ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput {
	return o
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput) ToApiCreateOrUpdatePropertiesWsdlSelectorPtrOutputWithContext(ctx context.Context) ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput {
	return o
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput) Elem() ApiCreateOrUpdatePropertiesWsdlSelectorOutput {
	return o.ApplyT(func(v *ApiCreateOrUpdatePropertiesWsdlSelector) ApiCreateOrUpdatePropertiesWsdlSelector {
		if v != nil {
			return *v
		}
		var ret ApiCreateOrUpdatePropertiesWsdlSelector
		return ret
	}).(ApiCreateOrUpdatePropertiesWsdlSelectorOutput)
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput) WsdlEndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiCreateOrUpdatePropertiesWsdlSelector) *string {
		if v == nil {
			return nil
		}
		return v.WsdlEndpointName
	}).(pulumi.StringPtrOutput)
}

func (o ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput) WsdlServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiCreateOrUpdatePropertiesWsdlSelector) *string {
		if v == nil {
			return nil
		}
		return v.WsdlServiceName
	}).(pulumi.StringPtrOutput)
}

type ApiLicenseInformation struct {
	Name *string `pulumi:"name"`
	Url  *string `pulumi:"url"`
}





type ApiLicenseInformationInput interface {
	pulumi.Input

	ToApiLicenseInformationOutput() ApiLicenseInformationOutput
	ToApiLicenseInformationOutputWithContext(context.Context) ApiLicenseInformationOutput
}

type ApiLicenseInformationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Url  pulumi.StringPtrInput `pulumi:"url"`
}

func (ApiLicenseInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiLicenseInformation)(nil)).Elem()
}

func (i ApiLicenseInformationArgs) ToApiLicenseInformationOutput() ApiLicenseInformationOutput {
	return i.ToApiLicenseInformationOutputWithContext(context.Background())
}

func (i ApiLicenseInformationArgs) ToApiLicenseInformationOutputWithContext(ctx context.Context) ApiLicenseInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiLicenseInformationOutput)
}

func (i ApiLicenseInformationArgs) ToApiLicenseInformationPtrOutput() ApiLicenseInformationPtrOutput {
	return i.ToApiLicenseInformationPtrOutputWithContext(context.Background())
}

func (i ApiLicenseInformationArgs) ToApiLicenseInformationPtrOutputWithContext(ctx context.Context) ApiLicenseInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiLicenseInformationOutput).ToApiLicenseInformationPtrOutputWithContext(ctx)
}









type ApiLicenseInformationPtrInput interface {
	pulumi.Input

	ToApiLicenseInformationPtrOutput() ApiLicenseInformationPtrOutput
	ToApiLicenseInformationPtrOutputWithContext(context.Context) ApiLicenseInformationPtrOutput
}

type apiLicenseInformationPtrType ApiLicenseInformationArgs

func ApiLicenseInformationPtr(v *ApiLicenseInformationArgs) ApiLicenseInformationPtrInput {
	return (*apiLicenseInformationPtrType)(v)
}

func (*apiLicenseInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiLicenseInformation)(nil)).Elem()
}

func (i *apiLicenseInformationPtrType) ToApiLicenseInformationPtrOutput() ApiLicenseInformationPtrOutput {
	return i.ToApiLicenseInformationPtrOutputWithContext(context.Background())
}

func (i *apiLicenseInformationPtrType) ToApiLicenseInformationPtrOutputWithContext(ctx context.Context) ApiLicenseInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiLicenseInformationPtrOutput)
}

type ApiLicenseInformationOutput struct{ *pulumi.OutputState }

func (ApiLicenseInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiLicenseInformation)(nil)).Elem()
}

func (o ApiLicenseInformationOutput) ToApiLicenseInformationOutput() ApiLicenseInformationOutput {
	return o
}

func (o ApiLicenseInformationOutput) ToApiLicenseInformationOutputWithContext(ctx context.Context) ApiLicenseInformationOutput {
	return o
}

func (o ApiLicenseInformationOutput) ToApiLicenseInformationPtrOutput() ApiLicenseInformationPtrOutput {
	return o.ToApiLicenseInformationPtrOutputWithContext(context.Background())
}

func (o ApiLicenseInformationOutput) ToApiLicenseInformationPtrOutputWithContext(ctx context.Context) ApiLicenseInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiLicenseInformation) *ApiLicenseInformation {
		return &v
	}).(ApiLicenseInformationPtrOutput)
}

func (o ApiLicenseInformationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiLicenseInformation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiLicenseInformationOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiLicenseInformation) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiLicenseInformationPtrOutput struct{ *pulumi.OutputState }

func (ApiLicenseInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiLicenseInformation)(nil)).Elem()
}

func (o ApiLicenseInformationPtrOutput) ToApiLicenseInformationPtrOutput() ApiLicenseInformationPtrOutput {
	return o
}

func (o ApiLicenseInformationPtrOutput) ToApiLicenseInformationPtrOutputWithContext(ctx context.Context) ApiLicenseInformationPtrOutput {
	return o
}

func (o ApiLicenseInformationPtrOutput) Elem() ApiLicenseInformationOutput {
	return o.ApplyT(func(v *ApiLicenseInformation) ApiLicenseInformation {
		if v != nil {
			return *v
		}
		var ret ApiLicenseInformation
		return ret
	}).(ApiLicenseInformationOutput)
}

func (o ApiLicenseInformationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiLicenseInformation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiLicenseInformationPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiLicenseInformation) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApiLicenseInformationResponse struct {
	Name *string `pulumi:"name"`
	Url  *string `pulumi:"url"`
}

type ApiLicenseInformationResponseOutput struct{ *pulumi.OutputState }

func (ApiLicenseInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiLicenseInformationResponse)(nil)).Elem()
}

func (o ApiLicenseInformationResponseOutput) ToApiLicenseInformationResponseOutput() ApiLicenseInformationResponseOutput {
	return o
}

func (o ApiLicenseInformationResponseOutput) ToApiLicenseInformationResponseOutputWithContext(ctx context.Context) ApiLicenseInformationResponseOutput {
	return o
}

func (o ApiLicenseInformationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiLicenseInformationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiLicenseInformationResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiLicenseInformationResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiLicenseInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiLicenseInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiLicenseInformationResponse)(nil)).Elem()
}

func (o ApiLicenseInformationResponsePtrOutput) ToApiLicenseInformationResponsePtrOutput() ApiLicenseInformationResponsePtrOutput {
	return o
}

func (o ApiLicenseInformationResponsePtrOutput) ToApiLicenseInformationResponsePtrOutputWithContext(ctx context.Context) ApiLicenseInformationResponsePtrOutput {
	return o
}

func (o ApiLicenseInformationResponsePtrOutput) Elem() ApiLicenseInformationResponseOutput {
	return o.ApplyT(func(v *ApiLicenseInformationResponse) ApiLicenseInformationResponse {
		if v != nil {
			return *v
		}
		var ret ApiLicenseInformationResponse
		return ret
	}).(ApiLicenseInformationResponseOutput)
}

func (o ApiLicenseInformationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiLicenseInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiLicenseInformationResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiLicenseInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApiManagementServiceIdentity struct {
	Type                   string                            `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityProperties `pulumi:"userAssignedIdentities"`
}





type ApiManagementServiceIdentityInput interface {
	pulumi.Input

	ToApiManagementServiceIdentityOutput() ApiManagementServiceIdentityOutput
	ToApiManagementServiceIdentityOutputWithContext(context.Context) ApiManagementServiceIdentityOutput
}

type ApiManagementServiceIdentityArgs struct {
	Type                   pulumi.StringInput             `pulumi:"type"`
	UserAssignedIdentities UserIdentityPropertiesMapInput `pulumi:"userAssignedIdentities"`
}

func (ApiManagementServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementServiceIdentity)(nil)).Elem()
}

func (i ApiManagementServiceIdentityArgs) ToApiManagementServiceIdentityOutput() ApiManagementServiceIdentityOutput {
	return i.ToApiManagementServiceIdentityOutputWithContext(context.Background())
}

func (i ApiManagementServiceIdentityArgs) ToApiManagementServiceIdentityOutputWithContext(ctx context.Context) ApiManagementServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceIdentityOutput)
}

func (i ApiManagementServiceIdentityArgs) ToApiManagementServiceIdentityPtrOutput() ApiManagementServiceIdentityPtrOutput {
	return i.ToApiManagementServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ApiManagementServiceIdentityArgs) ToApiManagementServiceIdentityPtrOutputWithContext(ctx context.Context) ApiManagementServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceIdentityOutput).ToApiManagementServiceIdentityPtrOutputWithContext(ctx)
}









type ApiManagementServiceIdentityPtrInput interface {
	pulumi.Input

	ToApiManagementServiceIdentityPtrOutput() ApiManagementServiceIdentityPtrOutput
	ToApiManagementServiceIdentityPtrOutputWithContext(context.Context) ApiManagementServiceIdentityPtrOutput
}

type apiManagementServiceIdentityPtrType ApiManagementServiceIdentityArgs

func ApiManagementServiceIdentityPtr(v *ApiManagementServiceIdentityArgs) ApiManagementServiceIdentityPtrInput {
	return (*apiManagementServiceIdentityPtrType)(v)
}

func (*apiManagementServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementServiceIdentity)(nil)).Elem()
}

func (i *apiManagementServiceIdentityPtrType) ToApiManagementServiceIdentityPtrOutput() ApiManagementServiceIdentityPtrOutput {
	return i.ToApiManagementServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *apiManagementServiceIdentityPtrType) ToApiManagementServiceIdentityPtrOutputWithContext(ctx context.Context) ApiManagementServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceIdentityPtrOutput)
}

type ApiManagementServiceIdentityOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementServiceIdentity)(nil)).Elem()
}

func (o ApiManagementServiceIdentityOutput) ToApiManagementServiceIdentityOutput() ApiManagementServiceIdentityOutput {
	return o
}

func (o ApiManagementServiceIdentityOutput) ToApiManagementServiceIdentityOutputWithContext(ctx context.Context) ApiManagementServiceIdentityOutput {
	return o
}

func (o ApiManagementServiceIdentityOutput) ToApiManagementServiceIdentityPtrOutput() ApiManagementServiceIdentityPtrOutput {
	return o.ToApiManagementServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ApiManagementServiceIdentityOutput) ToApiManagementServiceIdentityPtrOutputWithContext(ctx context.Context) ApiManagementServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiManagementServiceIdentity) *ApiManagementServiceIdentity {
		return &v
	}).(ApiManagementServiceIdentityPtrOutput)
}

func (o ApiManagementServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ApiManagementServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ApiManagementServiceIdentityOutput) UserAssignedIdentities() UserIdentityPropertiesMapOutput {
	return o.ApplyT(func(v ApiManagementServiceIdentity) map[string]UserIdentityProperties {
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesMapOutput)
}

type ApiManagementServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementServiceIdentity)(nil)).Elem()
}

func (o ApiManagementServiceIdentityPtrOutput) ToApiManagementServiceIdentityPtrOutput() ApiManagementServiceIdentityPtrOutput {
	return o
}

func (o ApiManagementServiceIdentityPtrOutput) ToApiManagementServiceIdentityPtrOutputWithContext(ctx context.Context) ApiManagementServiceIdentityPtrOutput {
	return o
}

func (o ApiManagementServiceIdentityPtrOutput) Elem() ApiManagementServiceIdentityOutput {
	return o.ApplyT(func(v *ApiManagementServiceIdentity) ApiManagementServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ApiManagementServiceIdentity
		return ret
	}).(ApiManagementServiceIdentityOutput)
}

func (o ApiManagementServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ApiManagementServiceIdentityPtrOutput) UserAssignedIdentities() UserIdentityPropertiesMapOutput {
	return o.ApplyT(func(v *ApiManagementServiceIdentity) map[string]UserIdentityProperties {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesMapOutput)
}

type ApiManagementServiceIdentityResponse struct {
	PrincipalId            string                                    `pulumi:"principalId"`
	TenantId               string                                    `pulumi:"tenantId"`
	Type                   string                                    `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityPropertiesResponse `pulumi:"userAssignedIdentities"`
}

type ApiManagementServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementServiceIdentityResponse)(nil)).Elem()
}

func (o ApiManagementServiceIdentityResponseOutput) ToApiManagementServiceIdentityResponseOutput() ApiManagementServiceIdentityResponseOutput {
	return o
}

func (o ApiManagementServiceIdentityResponseOutput) ToApiManagementServiceIdentityResponseOutputWithContext(ctx context.Context) ApiManagementServiceIdentityResponseOutput {
	return o
}

func (o ApiManagementServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApiManagementServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ApiManagementServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ApiManagementServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ApiManagementServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ApiManagementServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ApiManagementServiceIdentityResponseOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v ApiManagementServiceIdentityResponse) map[string]UserIdentityPropertiesResponse {
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type ApiManagementServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiManagementServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementServiceIdentityResponse)(nil)).Elem()
}

func (o ApiManagementServiceIdentityResponsePtrOutput) ToApiManagementServiceIdentityResponsePtrOutput() ApiManagementServiceIdentityResponsePtrOutput {
	return o
}

func (o ApiManagementServiceIdentityResponsePtrOutput) ToApiManagementServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ApiManagementServiceIdentityResponsePtrOutput {
	return o
}

func (o ApiManagementServiceIdentityResponsePtrOutput) Elem() ApiManagementServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ApiManagementServiceIdentityResponse) ApiManagementServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ApiManagementServiceIdentityResponse
		return ret
	}).(ApiManagementServiceIdentityResponseOutput)
}

func (o ApiManagementServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ApiManagementServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ApiManagementServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ApiManagementServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v *ApiManagementServiceIdentityResponse) map[string]UserIdentityPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type ApiManagementServiceSkuProperties struct {
	Capacity int    `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}





type ApiManagementServiceSkuPropertiesInput interface {
	pulumi.Input

	ToApiManagementServiceSkuPropertiesOutput() ApiManagementServiceSkuPropertiesOutput
	ToApiManagementServiceSkuPropertiesOutputWithContext(context.Context) ApiManagementServiceSkuPropertiesOutput
}

type ApiManagementServiceSkuPropertiesArgs struct {
	Capacity pulumi.IntInput    `pulumi:"capacity"`
	Name     pulumi.StringInput `pulumi:"name"`
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

func (o ApiManagementServiceSkuPropertiesOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v ApiManagementServiceSkuProperties) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o ApiManagementServiceSkuPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApiManagementServiceSkuProperties) string { return v.Name }).(pulumi.StringOutput)
}

type ApiManagementServiceSkuPropertiesResponse struct {
	Capacity int    `pulumi:"capacity"`
	Name     string `pulumi:"name"`
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

func (o ApiManagementServiceSkuPropertiesResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v ApiManagementServiceSkuPropertiesResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o ApiManagementServiceSkuPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApiManagementServiceSkuPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ApiVersionConstraint struct {
	MinApiVersion *string `pulumi:"minApiVersion"`
}





type ApiVersionConstraintInput interface {
	pulumi.Input

	ToApiVersionConstraintOutput() ApiVersionConstraintOutput
	ToApiVersionConstraintOutputWithContext(context.Context) ApiVersionConstraintOutput
}

type ApiVersionConstraintArgs struct {
	MinApiVersion pulumi.StringPtrInput `pulumi:"minApiVersion"`
}

func (ApiVersionConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionConstraint)(nil)).Elem()
}

func (i ApiVersionConstraintArgs) ToApiVersionConstraintOutput() ApiVersionConstraintOutput {
	return i.ToApiVersionConstraintOutputWithContext(context.Background())
}

func (i ApiVersionConstraintArgs) ToApiVersionConstraintOutputWithContext(ctx context.Context) ApiVersionConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionConstraintOutput)
}

func (i ApiVersionConstraintArgs) ToApiVersionConstraintPtrOutput() ApiVersionConstraintPtrOutput {
	return i.ToApiVersionConstraintPtrOutputWithContext(context.Background())
}

func (i ApiVersionConstraintArgs) ToApiVersionConstraintPtrOutputWithContext(ctx context.Context) ApiVersionConstraintPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionConstraintOutput).ToApiVersionConstraintPtrOutputWithContext(ctx)
}









type ApiVersionConstraintPtrInput interface {
	pulumi.Input

	ToApiVersionConstraintPtrOutput() ApiVersionConstraintPtrOutput
	ToApiVersionConstraintPtrOutputWithContext(context.Context) ApiVersionConstraintPtrOutput
}

type apiVersionConstraintPtrType ApiVersionConstraintArgs

func ApiVersionConstraintPtr(v *ApiVersionConstraintArgs) ApiVersionConstraintPtrInput {
	return (*apiVersionConstraintPtrType)(v)
}

func (*apiVersionConstraintPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionConstraint)(nil)).Elem()
}

func (i *apiVersionConstraintPtrType) ToApiVersionConstraintPtrOutput() ApiVersionConstraintPtrOutput {
	return i.ToApiVersionConstraintPtrOutputWithContext(context.Background())
}

func (i *apiVersionConstraintPtrType) ToApiVersionConstraintPtrOutputWithContext(ctx context.Context) ApiVersionConstraintPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionConstraintPtrOutput)
}

type ApiVersionConstraintOutput struct{ *pulumi.OutputState }

func (ApiVersionConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionConstraint)(nil)).Elem()
}

func (o ApiVersionConstraintOutput) ToApiVersionConstraintOutput() ApiVersionConstraintOutput {
	return o
}

func (o ApiVersionConstraintOutput) ToApiVersionConstraintOutputWithContext(ctx context.Context) ApiVersionConstraintOutput {
	return o
}

func (o ApiVersionConstraintOutput) ToApiVersionConstraintPtrOutput() ApiVersionConstraintPtrOutput {
	return o.ToApiVersionConstraintPtrOutputWithContext(context.Background())
}

func (o ApiVersionConstraintOutput) ToApiVersionConstraintPtrOutputWithContext(ctx context.Context) ApiVersionConstraintPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiVersionConstraint) *ApiVersionConstraint {
		return &v
	}).(ApiVersionConstraintPtrOutput)
}

func (o ApiVersionConstraintOutput) MinApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionConstraint) *string { return v.MinApiVersion }).(pulumi.StringPtrOutput)
}

type ApiVersionConstraintPtrOutput struct{ *pulumi.OutputState }

func (ApiVersionConstraintPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionConstraint)(nil)).Elem()
}

func (o ApiVersionConstraintPtrOutput) ToApiVersionConstraintPtrOutput() ApiVersionConstraintPtrOutput {
	return o
}

func (o ApiVersionConstraintPtrOutput) ToApiVersionConstraintPtrOutputWithContext(ctx context.Context) ApiVersionConstraintPtrOutput {
	return o
}

func (o ApiVersionConstraintPtrOutput) Elem() ApiVersionConstraintOutput {
	return o.ApplyT(func(v *ApiVersionConstraint) ApiVersionConstraint {
		if v != nil {
			return *v
		}
		var ret ApiVersionConstraint
		return ret
	}).(ApiVersionConstraintOutput)
}

func (o ApiVersionConstraintPtrOutput) MinApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionConstraint) *string {
		if v == nil {
			return nil
		}
		return v.MinApiVersion
	}).(pulumi.StringPtrOutput)
}

type ApiVersionConstraintResponse struct {
	MinApiVersion *string `pulumi:"minApiVersion"`
}

type ApiVersionConstraintResponseOutput struct{ *pulumi.OutputState }

func (ApiVersionConstraintResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionConstraintResponse)(nil)).Elem()
}

func (o ApiVersionConstraintResponseOutput) ToApiVersionConstraintResponseOutput() ApiVersionConstraintResponseOutput {
	return o
}

func (o ApiVersionConstraintResponseOutput) ToApiVersionConstraintResponseOutputWithContext(ctx context.Context) ApiVersionConstraintResponseOutput {
	return o
}

func (o ApiVersionConstraintResponseOutput) MinApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionConstraintResponse) *string { return v.MinApiVersion }).(pulumi.StringPtrOutput)
}

type ApiVersionConstraintResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiVersionConstraintResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionConstraintResponse)(nil)).Elem()
}

func (o ApiVersionConstraintResponsePtrOutput) ToApiVersionConstraintResponsePtrOutput() ApiVersionConstraintResponsePtrOutput {
	return o
}

func (o ApiVersionConstraintResponsePtrOutput) ToApiVersionConstraintResponsePtrOutputWithContext(ctx context.Context) ApiVersionConstraintResponsePtrOutput {
	return o
}

func (o ApiVersionConstraintResponsePtrOutput) Elem() ApiVersionConstraintResponseOutput {
	return o.ApplyT(func(v *ApiVersionConstraintResponse) ApiVersionConstraintResponse {
		if v != nil {
			return *v
		}
		var ret ApiVersionConstraintResponse
		return ret
	}).(ApiVersionConstraintResponseOutput)
}

func (o ApiVersionConstraintResponsePtrOutput) MinApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionConstraintResponse) *string {
		if v == nil {
			return nil
		}
		return v.MinApiVersion
	}).(pulumi.StringPtrOutput)
}

type ApiVersionSetContractDetails struct {
	Description       *string `pulumi:"description"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	VersionQueryName  *string `pulumi:"versionQueryName"`
	VersioningScheme  *string `pulumi:"versioningScheme"`
}





type ApiVersionSetContractDetailsInput interface {
	pulumi.Input

	ToApiVersionSetContractDetailsOutput() ApiVersionSetContractDetailsOutput
	ToApiVersionSetContractDetailsOutputWithContext(context.Context) ApiVersionSetContractDetailsOutput
}

type ApiVersionSetContractDetailsArgs struct {
	Description       pulumi.StringPtrInput `pulumi:"description"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	VersionHeaderName pulumi.StringPtrInput `pulumi:"versionHeaderName"`
	VersionQueryName  pulumi.StringPtrInput `pulumi:"versionQueryName"`
	VersioningScheme  pulumi.StringPtrInput `pulumi:"versioningScheme"`
}

func (ApiVersionSetContractDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSetContractDetails)(nil)).Elem()
}

func (i ApiVersionSetContractDetailsArgs) ToApiVersionSetContractDetailsOutput() ApiVersionSetContractDetailsOutput {
	return i.ToApiVersionSetContractDetailsOutputWithContext(context.Background())
}

func (i ApiVersionSetContractDetailsArgs) ToApiVersionSetContractDetailsOutputWithContext(ctx context.Context) ApiVersionSetContractDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetContractDetailsOutput)
}

func (i ApiVersionSetContractDetailsArgs) ToApiVersionSetContractDetailsPtrOutput() ApiVersionSetContractDetailsPtrOutput {
	return i.ToApiVersionSetContractDetailsPtrOutputWithContext(context.Background())
}

func (i ApiVersionSetContractDetailsArgs) ToApiVersionSetContractDetailsPtrOutputWithContext(ctx context.Context) ApiVersionSetContractDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetContractDetailsOutput).ToApiVersionSetContractDetailsPtrOutputWithContext(ctx)
}









type ApiVersionSetContractDetailsPtrInput interface {
	pulumi.Input

	ToApiVersionSetContractDetailsPtrOutput() ApiVersionSetContractDetailsPtrOutput
	ToApiVersionSetContractDetailsPtrOutputWithContext(context.Context) ApiVersionSetContractDetailsPtrOutput
}

type apiVersionSetContractDetailsPtrType ApiVersionSetContractDetailsArgs

func ApiVersionSetContractDetailsPtr(v *ApiVersionSetContractDetailsArgs) ApiVersionSetContractDetailsPtrInput {
	return (*apiVersionSetContractDetailsPtrType)(v)
}

func (*apiVersionSetContractDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSetContractDetails)(nil)).Elem()
}

func (i *apiVersionSetContractDetailsPtrType) ToApiVersionSetContractDetailsPtrOutput() ApiVersionSetContractDetailsPtrOutput {
	return i.ToApiVersionSetContractDetailsPtrOutputWithContext(context.Background())
}

func (i *apiVersionSetContractDetailsPtrType) ToApiVersionSetContractDetailsPtrOutputWithContext(ctx context.Context) ApiVersionSetContractDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetContractDetailsPtrOutput)
}

type ApiVersionSetContractDetailsOutput struct{ *pulumi.OutputState }

func (ApiVersionSetContractDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSetContractDetails)(nil)).Elem()
}

func (o ApiVersionSetContractDetailsOutput) ToApiVersionSetContractDetailsOutput() ApiVersionSetContractDetailsOutput {
	return o
}

func (o ApiVersionSetContractDetailsOutput) ToApiVersionSetContractDetailsOutputWithContext(ctx context.Context) ApiVersionSetContractDetailsOutput {
	return o
}

func (o ApiVersionSetContractDetailsOutput) ToApiVersionSetContractDetailsPtrOutput() ApiVersionSetContractDetailsPtrOutput {
	return o.ToApiVersionSetContractDetailsPtrOutputWithContext(context.Background())
}

func (o ApiVersionSetContractDetailsOutput) ToApiVersionSetContractDetailsPtrOutputWithContext(ctx context.Context) ApiVersionSetContractDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiVersionSetContractDetails) *ApiVersionSetContractDetails {
		return &v
	}).(ApiVersionSetContractDetailsPtrOutput)
}

func (o ApiVersionSetContractDetailsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetails) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetails) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetails) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetails) *string { return v.VersionHeaderName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetails) *string { return v.VersionQueryName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsOutput) VersioningScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetails) *string { return v.VersioningScheme }).(pulumi.StringPtrOutput)
}

type ApiVersionSetContractDetailsPtrOutput struct{ *pulumi.OutputState }

func (ApiVersionSetContractDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSetContractDetails)(nil)).Elem()
}

func (o ApiVersionSetContractDetailsPtrOutput) ToApiVersionSetContractDetailsPtrOutput() ApiVersionSetContractDetailsPtrOutput {
	return o
}

func (o ApiVersionSetContractDetailsPtrOutput) ToApiVersionSetContractDetailsPtrOutputWithContext(ctx context.Context) ApiVersionSetContractDetailsPtrOutput {
	return o
}

func (o ApiVersionSetContractDetailsPtrOutput) Elem() ApiVersionSetContractDetailsOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetails) ApiVersionSetContractDetails {
		if v != nil {
			return *v
		}
		var ret ApiVersionSetContractDetails
		return ret
	}).(ApiVersionSetContractDetailsOutput)
}

func (o ApiVersionSetContractDetailsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetails) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetails) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetails) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsPtrOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetails) *string {
		if v == nil {
			return nil
		}
		return v.VersionHeaderName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsPtrOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetails) *string {
		if v == nil {
			return nil
		}
		return v.VersionQueryName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsPtrOutput) VersioningScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetails) *string {
		if v == nil {
			return nil
		}
		return v.VersioningScheme
	}).(pulumi.StringPtrOutput)
}

type ApiVersionSetContractDetailsResponse struct {
	Description       *string `pulumi:"description"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	VersionQueryName  *string `pulumi:"versionQueryName"`
	VersioningScheme  *string `pulumi:"versioningScheme"`
}

type ApiVersionSetContractDetailsResponseOutput struct{ *pulumi.OutputState }

func (ApiVersionSetContractDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSetContractDetailsResponse)(nil)).Elem()
}

func (o ApiVersionSetContractDetailsResponseOutput) ToApiVersionSetContractDetailsResponseOutput() ApiVersionSetContractDetailsResponseOutput {
	return o
}

func (o ApiVersionSetContractDetailsResponseOutput) ToApiVersionSetContractDetailsResponseOutputWithContext(ctx context.Context) ApiVersionSetContractDetailsResponseOutput {
	return o
}

func (o ApiVersionSetContractDetailsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetailsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetailsResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetailsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponseOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetailsResponse) *string { return v.VersionHeaderName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponseOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetailsResponse) *string { return v.VersionQueryName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponseOutput) VersioningScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractDetailsResponse) *string { return v.VersioningScheme }).(pulumi.StringPtrOutput)
}

type ApiVersionSetContractDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiVersionSetContractDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSetContractDetailsResponse)(nil)).Elem()
}

func (o ApiVersionSetContractDetailsResponsePtrOutput) ToApiVersionSetContractDetailsResponsePtrOutput() ApiVersionSetContractDetailsResponsePtrOutput {
	return o
}

func (o ApiVersionSetContractDetailsResponsePtrOutput) ToApiVersionSetContractDetailsResponsePtrOutputWithContext(ctx context.Context) ApiVersionSetContractDetailsResponsePtrOutput {
	return o
}

func (o ApiVersionSetContractDetailsResponsePtrOutput) Elem() ApiVersionSetContractDetailsResponseOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetailsResponse) ApiVersionSetContractDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ApiVersionSetContractDetailsResponse
		return ret
	}).(ApiVersionSetContractDetailsResponseOutput)
}

func (o ApiVersionSetContractDetailsResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponsePtrOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.VersionHeaderName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponsePtrOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.VersionQueryName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractDetailsResponsePtrOutput) VersioningScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.VersioningScheme
	}).(pulumi.StringPtrOutput)
}

type ArmIdWrapperResponse struct {
	Id string `pulumi:"id"`
}

type ArmIdWrapperResponseOutput struct{ *pulumi.OutputState }

func (ArmIdWrapperResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdWrapperResponse)(nil)).Elem()
}

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponseOutput() ArmIdWrapperResponseOutput {
	return o
}

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponseOutputWithContext(ctx context.Context) ArmIdWrapperResponseOutput {
	return o
}

func (o ArmIdWrapperResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ArmIdWrapperResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ArmIdWrapperResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmIdWrapperResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdWrapperResponse)(nil)).Elem()
}

func (o ArmIdWrapperResponsePtrOutput) ToArmIdWrapperResponsePtrOutput() ArmIdWrapperResponsePtrOutput {
	return o
}

func (o ArmIdWrapperResponsePtrOutput) ToArmIdWrapperResponsePtrOutputWithContext(ctx context.Context) ArmIdWrapperResponsePtrOutput {
	return o
}

func (o ArmIdWrapperResponsePtrOutput) Elem() ArmIdWrapperResponseOutput {
	return o.ApplyT(func(v *ArmIdWrapperResponse) ArmIdWrapperResponse {
		if v != nil {
			return *v
		}
		var ret ArmIdWrapperResponse
		return ret
	}).(ArmIdWrapperResponseOutput)
}

func (o ArmIdWrapperResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdWrapperResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type AuthenticationSettingsContract struct {
	OAuth2 *OAuth2AuthenticationSettingsContract `pulumi:"oAuth2"`
	Openid *OpenIdAuthenticationSettingsContract `pulumi:"openid"`
}





type AuthenticationSettingsContractInput interface {
	pulumi.Input

	ToAuthenticationSettingsContractOutput() AuthenticationSettingsContractOutput
	ToAuthenticationSettingsContractOutputWithContext(context.Context) AuthenticationSettingsContractOutput
}

type AuthenticationSettingsContractArgs struct {
	OAuth2 OAuth2AuthenticationSettingsContractPtrInput `pulumi:"oAuth2"`
	Openid OpenIdAuthenticationSettingsContractPtrInput `pulumi:"openid"`
}

func (AuthenticationSettingsContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationSettingsContract)(nil)).Elem()
}

func (i AuthenticationSettingsContractArgs) ToAuthenticationSettingsContractOutput() AuthenticationSettingsContractOutput {
	return i.ToAuthenticationSettingsContractOutputWithContext(context.Background())
}

func (i AuthenticationSettingsContractArgs) ToAuthenticationSettingsContractOutputWithContext(ctx context.Context) AuthenticationSettingsContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationSettingsContractOutput)
}

func (i AuthenticationSettingsContractArgs) ToAuthenticationSettingsContractPtrOutput() AuthenticationSettingsContractPtrOutput {
	return i.ToAuthenticationSettingsContractPtrOutputWithContext(context.Background())
}

func (i AuthenticationSettingsContractArgs) ToAuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) AuthenticationSettingsContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationSettingsContractOutput).ToAuthenticationSettingsContractPtrOutputWithContext(ctx)
}









type AuthenticationSettingsContractPtrInput interface {
	pulumi.Input

	ToAuthenticationSettingsContractPtrOutput() AuthenticationSettingsContractPtrOutput
	ToAuthenticationSettingsContractPtrOutputWithContext(context.Context) AuthenticationSettingsContractPtrOutput
}

type authenticationSettingsContractPtrType AuthenticationSettingsContractArgs

func AuthenticationSettingsContractPtr(v *AuthenticationSettingsContractArgs) AuthenticationSettingsContractPtrInput {
	return (*authenticationSettingsContractPtrType)(v)
}

func (*authenticationSettingsContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationSettingsContract)(nil)).Elem()
}

func (i *authenticationSettingsContractPtrType) ToAuthenticationSettingsContractPtrOutput() AuthenticationSettingsContractPtrOutput {
	return i.ToAuthenticationSettingsContractPtrOutputWithContext(context.Background())
}

func (i *authenticationSettingsContractPtrType) ToAuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) AuthenticationSettingsContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationSettingsContractPtrOutput)
}

type AuthenticationSettingsContractOutput struct{ *pulumi.OutputState }

func (AuthenticationSettingsContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationSettingsContract)(nil)).Elem()
}

func (o AuthenticationSettingsContractOutput) ToAuthenticationSettingsContractOutput() AuthenticationSettingsContractOutput {
	return o
}

func (o AuthenticationSettingsContractOutput) ToAuthenticationSettingsContractOutputWithContext(ctx context.Context) AuthenticationSettingsContractOutput {
	return o
}

func (o AuthenticationSettingsContractOutput) ToAuthenticationSettingsContractPtrOutput() AuthenticationSettingsContractPtrOutput {
	return o.ToAuthenticationSettingsContractPtrOutputWithContext(context.Background())
}

func (o AuthenticationSettingsContractOutput) ToAuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) AuthenticationSettingsContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationSettingsContract) *AuthenticationSettingsContract {
		return &v
	}).(AuthenticationSettingsContractPtrOutput)
}

func (o AuthenticationSettingsContractOutput) OAuth2() OAuth2AuthenticationSettingsContractPtrOutput {
	return o.ApplyT(func(v AuthenticationSettingsContract) *OAuth2AuthenticationSettingsContract { return v.OAuth2 }).(OAuth2AuthenticationSettingsContractPtrOutput)
}

func (o AuthenticationSettingsContractOutput) Openid() OpenIdAuthenticationSettingsContractPtrOutput {
	return o.ApplyT(func(v AuthenticationSettingsContract) *OpenIdAuthenticationSettingsContract { return v.Openid }).(OpenIdAuthenticationSettingsContractPtrOutput)
}

type AuthenticationSettingsContractPtrOutput struct{ *pulumi.OutputState }

func (AuthenticationSettingsContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationSettingsContract)(nil)).Elem()
}

func (o AuthenticationSettingsContractPtrOutput) ToAuthenticationSettingsContractPtrOutput() AuthenticationSettingsContractPtrOutput {
	return o
}

func (o AuthenticationSettingsContractPtrOutput) ToAuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) AuthenticationSettingsContractPtrOutput {
	return o
}

func (o AuthenticationSettingsContractPtrOutput) Elem() AuthenticationSettingsContractOutput {
	return o.ApplyT(func(v *AuthenticationSettingsContract) AuthenticationSettingsContract {
		if v != nil {
			return *v
		}
		var ret AuthenticationSettingsContract
		return ret
	}).(AuthenticationSettingsContractOutput)
}

func (o AuthenticationSettingsContractPtrOutput) OAuth2() OAuth2AuthenticationSettingsContractPtrOutput {
	return o.ApplyT(func(v *AuthenticationSettingsContract) *OAuth2AuthenticationSettingsContract {
		if v == nil {
			return nil
		}
		return v.OAuth2
	}).(OAuth2AuthenticationSettingsContractPtrOutput)
}

func (o AuthenticationSettingsContractPtrOutput) Openid() OpenIdAuthenticationSettingsContractPtrOutput {
	return o.ApplyT(func(v *AuthenticationSettingsContract) *OpenIdAuthenticationSettingsContract {
		if v == nil {
			return nil
		}
		return v.Openid
	}).(OpenIdAuthenticationSettingsContractPtrOutput)
}

type AuthenticationSettingsContractResponse struct {
	OAuth2 *OAuth2AuthenticationSettingsContractResponse `pulumi:"oAuth2"`
	Openid *OpenIdAuthenticationSettingsContractResponse `pulumi:"openid"`
}

type AuthenticationSettingsContractResponseOutput struct{ *pulumi.OutputState }

func (AuthenticationSettingsContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationSettingsContractResponse)(nil)).Elem()
}

func (o AuthenticationSettingsContractResponseOutput) ToAuthenticationSettingsContractResponseOutput() AuthenticationSettingsContractResponseOutput {
	return o
}

func (o AuthenticationSettingsContractResponseOutput) ToAuthenticationSettingsContractResponseOutputWithContext(ctx context.Context) AuthenticationSettingsContractResponseOutput {
	return o
}

func (o AuthenticationSettingsContractResponseOutput) OAuth2() OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v AuthenticationSettingsContractResponse) *OAuth2AuthenticationSettingsContractResponse {
		return v.OAuth2
	}).(OAuth2AuthenticationSettingsContractResponsePtrOutput)
}

func (o AuthenticationSettingsContractResponseOutput) Openid() OpenIdAuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v AuthenticationSettingsContractResponse) *OpenIdAuthenticationSettingsContractResponse {
		return v.Openid
	}).(OpenIdAuthenticationSettingsContractResponsePtrOutput)
}

type AuthenticationSettingsContractResponsePtrOutput struct{ *pulumi.OutputState }

func (AuthenticationSettingsContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationSettingsContractResponse)(nil)).Elem()
}

func (o AuthenticationSettingsContractResponsePtrOutput) ToAuthenticationSettingsContractResponsePtrOutput() AuthenticationSettingsContractResponsePtrOutput {
	return o
}

func (o AuthenticationSettingsContractResponsePtrOutput) ToAuthenticationSettingsContractResponsePtrOutputWithContext(ctx context.Context) AuthenticationSettingsContractResponsePtrOutput {
	return o
}

func (o AuthenticationSettingsContractResponsePtrOutput) Elem() AuthenticationSettingsContractResponseOutput {
	return o.ApplyT(func(v *AuthenticationSettingsContractResponse) AuthenticationSettingsContractResponse {
		if v != nil {
			return *v
		}
		var ret AuthenticationSettingsContractResponse
		return ret
	}).(AuthenticationSettingsContractResponseOutput)
}

func (o AuthenticationSettingsContractResponsePtrOutput) OAuth2() OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v *AuthenticationSettingsContractResponse) *OAuth2AuthenticationSettingsContractResponse {
		if v == nil {
			return nil
		}
		return v.OAuth2
	}).(OAuth2AuthenticationSettingsContractResponsePtrOutput)
}

func (o AuthenticationSettingsContractResponsePtrOutput) Openid() OpenIdAuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v *AuthenticationSettingsContractResponse) *OpenIdAuthenticationSettingsContractResponse {
		if v == nil {
			return nil
		}
		return v.Openid
	}).(OpenIdAuthenticationSettingsContractResponsePtrOutput)
}

type BackendAuthorizationHeaderCredentials struct {
	Parameter string `pulumi:"parameter"`
	Scheme    string `pulumi:"scheme"`
}





type BackendAuthorizationHeaderCredentialsInput interface {
	pulumi.Input

	ToBackendAuthorizationHeaderCredentialsOutput() BackendAuthorizationHeaderCredentialsOutput
	ToBackendAuthorizationHeaderCredentialsOutputWithContext(context.Context) BackendAuthorizationHeaderCredentialsOutput
}

type BackendAuthorizationHeaderCredentialsArgs struct {
	Parameter pulumi.StringInput `pulumi:"parameter"`
	Scheme    pulumi.StringInput `pulumi:"scheme"`
}

func (BackendAuthorizationHeaderCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAuthorizationHeaderCredentials)(nil)).Elem()
}

func (i BackendAuthorizationHeaderCredentialsArgs) ToBackendAuthorizationHeaderCredentialsOutput() BackendAuthorizationHeaderCredentialsOutput {
	return i.ToBackendAuthorizationHeaderCredentialsOutputWithContext(context.Background())
}

func (i BackendAuthorizationHeaderCredentialsArgs) ToBackendAuthorizationHeaderCredentialsOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAuthorizationHeaderCredentialsOutput)
}

func (i BackendAuthorizationHeaderCredentialsArgs) ToBackendAuthorizationHeaderCredentialsPtrOutput() BackendAuthorizationHeaderCredentialsPtrOutput {
	return i.ToBackendAuthorizationHeaderCredentialsPtrOutputWithContext(context.Background())
}

func (i BackendAuthorizationHeaderCredentialsArgs) ToBackendAuthorizationHeaderCredentialsPtrOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAuthorizationHeaderCredentialsOutput).ToBackendAuthorizationHeaderCredentialsPtrOutputWithContext(ctx)
}









type BackendAuthorizationHeaderCredentialsPtrInput interface {
	pulumi.Input

	ToBackendAuthorizationHeaderCredentialsPtrOutput() BackendAuthorizationHeaderCredentialsPtrOutput
	ToBackendAuthorizationHeaderCredentialsPtrOutputWithContext(context.Context) BackendAuthorizationHeaderCredentialsPtrOutput
}

type backendAuthorizationHeaderCredentialsPtrType BackendAuthorizationHeaderCredentialsArgs

func BackendAuthorizationHeaderCredentialsPtr(v *BackendAuthorizationHeaderCredentialsArgs) BackendAuthorizationHeaderCredentialsPtrInput {
	return (*backendAuthorizationHeaderCredentialsPtrType)(v)
}

func (*backendAuthorizationHeaderCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendAuthorizationHeaderCredentials)(nil)).Elem()
}

func (i *backendAuthorizationHeaderCredentialsPtrType) ToBackendAuthorizationHeaderCredentialsPtrOutput() BackendAuthorizationHeaderCredentialsPtrOutput {
	return i.ToBackendAuthorizationHeaderCredentialsPtrOutputWithContext(context.Background())
}

func (i *backendAuthorizationHeaderCredentialsPtrType) ToBackendAuthorizationHeaderCredentialsPtrOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAuthorizationHeaderCredentialsPtrOutput)
}

type BackendAuthorizationHeaderCredentialsOutput struct{ *pulumi.OutputState }

func (BackendAuthorizationHeaderCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAuthorizationHeaderCredentials)(nil)).Elem()
}

func (o BackendAuthorizationHeaderCredentialsOutput) ToBackendAuthorizationHeaderCredentialsOutput() BackendAuthorizationHeaderCredentialsOutput {
	return o
}

func (o BackendAuthorizationHeaderCredentialsOutput) ToBackendAuthorizationHeaderCredentialsOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsOutput {
	return o
}

func (o BackendAuthorizationHeaderCredentialsOutput) ToBackendAuthorizationHeaderCredentialsPtrOutput() BackendAuthorizationHeaderCredentialsPtrOutput {
	return o.ToBackendAuthorizationHeaderCredentialsPtrOutputWithContext(context.Background())
}

func (o BackendAuthorizationHeaderCredentialsOutput) ToBackendAuthorizationHeaderCredentialsPtrOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendAuthorizationHeaderCredentials) *BackendAuthorizationHeaderCredentials {
		return &v
	}).(BackendAuthorizationHeaderCredentialsPtrOutput)
}

func (o BackendAuthorizationHeaderCredentialsOutput) Parameter() pulumi.StringOutput {
	return o.ApplyT(func(v BackendAuthorizationHeaderCredentials) string { return v.Parameter }).(pulumi.StringOutput)
}

func (o BackendAuthorizationHeaderCredentialsOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v BackendAuthorizationHeaderCredentials) string { return v.Scheme }).(pulumi.StringOutput)
}

type BackendAuthorizationHeaderCredentialsPtrOutput struct{ *pulumi.OutputState }

func (BackendAuthorizationHeaderCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendAuthorizationHeaderCredentials)(nil)).Elem()
}

func (o BackendAuthorizationHeaderCredentialsPtrOutput) ToBackendAuthorizationHeaderCredentialsPtrOutput() BackendAuthorizationHeaderCredentialsPtrOutput {
	return o
}

func (o BackendAuthorizationHeaderCredentialsPtrOutput) ToBackendAuthorizationHeaderCredentialsPtrOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsPtrOutput {
	return o
}

func (o BackendAuthorizationHeaderCredentialsPtrOutput) Elem() BackendAuthorizationHeaderCredentialsOutput {
	return o.ApplyT(func(v *BackendAuthorizationHeaderCredentials) BackendAuthorizationHeaderCredentials {
		if v != nil {
			return *v
		}
		var ret BackendAuthorizationHeaderCredentials
		return ret
	}).(BackendAuthorizationHeaderCredentialsOutput)
}

func (o BackendAuthorizationHeaderCredentialsPtrOutput) Parameter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendAuthorizationHeaderCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Parameter
	}).(pulumi.StringPtrOutput)
}

func (o BackendAuthorizationHeaderCredentialsPtrOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendAuthorizationHeaderCredentials) *string {
		if v == nil {
			return nil
		}
		return &v.Scheme
	}).(pulumi.StringPtrOutput)
}

type BackendAuthorizationHeaderCredentialsResponse struct {
	Parameter string `pulumi:"parameter"`
	Scheme    string `pulumi:"scheme"`
}

type BackendAuthorizationHeaderCredentialsResponseOutput struct{ *pulumi.OutputState }

func (BackendAuthorizationHeaderCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAuthorizationHeaderCredentialsResponse)(nil)).Elem()
}

func (o BackendAuthorizationHeaderCredentialsResponseOutput) ToBackendAuthorizationHeaderCredentialsResponseOutput() BackendAuthorizationHeaderCredentialsResponseOutput {
	return o
}

func (o BackendAuthorizationHeaderCredentialsResponseOutput) ToBackendAuthorizationHeaderCredentialsResponseOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsResponseOutput {
	return o
}

func (o BackendAuthorizationHeaderCredentialsResponseOutput) Parameter() pulumi.StringOutput {
	return o.ApplyT(func(v BackendAuthorizationHeaderCredentialsResponse) string { return v.Parameter }).(pulumi.StringOutput)
}

func (o BackendAuthorizationHeaderCredentialsResponseOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v BackendAuthorizationHeaderCredentialsResponse) string { return v.Scheme }).(pulumi.StringOutput)
}

type BackendAuthorizationHeaderCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (BackendAuthorizationHeaderCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendAuthorizationHeaderCredentialsResponse)(nil)).Elem()
}

func (o BackendAuthorizationHeaderCredentialsResponsePtrOutput) ToBackendAuthorizationHeaderCredentialsResponsePtrOutput() BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return o
}

func (o BackendAuthorizationHeaderCredentialsResponsePtrOutput) ToBackendAuthorizationHeaderCredentialsResponsePtrOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return o
}

func (o BackendAuthorizationHeaderCredentialsResponsePtrOutput) Elem() BackendAuthorizationHeaderCredentialsResponseOutput {
	return o.ApplyT(func(v *BackendAuthorizationHeaderCredentialsResponse) BackendAuthorizationHeaderCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret BackendAuthorizationHeaderCredentialsResponse
		return ret
	}).(BackendAuthorizationHeaderCredentialsResponseOutput)
}

func (o BackendAuthorizationHeaderCredentialsResponsePtrOutput) Parameter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendAuthorizationHeaderCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Parameter
	}).(pulumi.StringPtrOutput)
}

func (o BackendAuthorizationHeaderCredentialsResponsePtrOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendAuthorizationHeaderCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Scheme
	}).(pulumi.StringPtrOutput)
}

type BackendCredentialsContract struct {
	Authorization  *BackendAuthorizationHeaderCredentials `pulumi:"authorization"`
	Certificate    []string                               `pulumi:"certificate"`
	CertificateIds []string                               `pulumi:"certificateIds"`
	Header         map[string][]string                    `pulumi:"header"`
	Query          map[string][]string                    `pulumi:"query"`
}





type BackendCredentialsContractInput interface {
	pulumi.Input

	ToBackendCredentialsContractOutput() BackendCredentialsContractOutput
	ToBackendCredentialsContractOutputWithContext(context.Context) BackendCredentialsContractOutput
}

type BackendCredentialsContractArgs struct {
	Authorization  BackendAuthorizationHeaderCredentialsPtrInput `pulumi:"authorization"`
	Certificate    pulumi.StringArrayInput                       `pulumi:"certificate"`
	CertificateIds pulumi.StringArrayInput                       `pulumi:"certificateIds"`
	Header         pulumi.StringArrayMapInput                    `pulumi:"header"`
	Query          pulumi.StringArrayMapInput                    `pulumi:"query"`
}

func (BackendCredentialsContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendCredentialsContract)(nil)).Elem()
}

func (i BackendCredentialsContractArgs) ToBackendCredentialsContractOutput() BackendCredentialsContractOutput {
	return i.ToBackendCredentialsContractOutputWithContext(context.Background())
}

func (i BackendCredentialsContractArgs) ToBackendCredentialsContractOutputWithContext(ctx context.Context) BackendCredentialsContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendCredentialsContractOutput)
}

func (i BackendCredentialsContractArgs) ToBackendCredentialsContractPtrOutput() BackendCredentialsContractPtrOutput {
	return i.ToBackendCredentialsContractPtrOutputWithContext(context.Background())
}

func (i BackendCredentialsContractArgs) ToBackendCredentialsContractPtrOutputWithContext(ctx context.Context) BackendCredentialsContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendCredentialsContractOutput).ToBackendCredentialsContractPtrOutputWithContext(ctx)
}









type BackendCredentialsContractPtrInput interface {
	pulumi.Input

	ToBackendCredentialsContractPtrOutput() BackendCredentialsContractPtrOutput
	ToBackendCredentialsContractPtrOutputWithContext(context.Context) BackendCredentialsContractPtrOutput
}

type backendCredentialsContractPtrType BackendCredentialsContractArgs

func BackendCredentialsContractPtr(v *BackendCredentialsContractArgs) BackendCredentialsContractPtrInput {
	return (*backendCredentialsContractPtrType)(v)
}

func (*backendCredentialsContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendCredentialsContract)(nil)).Elem()
}

func (i *backendCredentialsContractPtrType) ToBackendCredentialsContractPtrOutput() BackendCredentialsContractPtrOutput {
	return i.ToBackendCredentialsContractPtrOutputWithContext(context.Background())
}

func (i *backendCredentialsContractPtrType) ToBackendCredentialsContractPtrOutputWithContext(ctx context.Context) BackendCredentialsContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendCredentialsContractPtrOutput)
}

type BackendCredentialsContractOutput struct{ *pulumi.OutputState }

func (BackendCredentialsContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendCredentialsContract)(nil)).Elem()
}

func (o BackendCredentialsContractOutput) ToBackendCredentialsContractOutput() BackendCredentialsContractOutput {
	return o
}

func (o BackendCredentialsContractOutput) ToBackendCredentialsContractOutputWithContext(ctx context.Context) BackendCredentialsContractOutput {
	return o
}

func (o BackendCredentialsContractOutput) ToBackendCredentialsContractPtrOutput() BackendCredentialsContractPtrOutput {
	return o.ToBackendCredentialsContractPtrOutputWithContext(context.Background())
}

func (o BackendCredentialsContractOutput) ToBackendCredentialsContractPtrOutputWithContext(ctx context.Context) BackendCredentialsContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendCredentialsContract) *BackendCredentialsContract {
		return &v
	}).(BackendCredentialsContractPtrOutput)
}

func (o BackendCredentialsContractOutput) Authorization() BackendAuthorizationHeaderCredentialsPtrOutput {
	return o.ApplyT(func(v BackendCredentialsContract) *BackendAuthorizationHeaderCredentials { return v.Authorization }).(BackendAuthorizationHeaderCredentialsPtrOutput)
}

func (o BackendCredentialsContractOutput) Certificate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackendCredentialsContract) []string { return v.Certificate }).(pulumi.StringArrayOutput)
}

func (o BackendCredentialsContractOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackendCredentialsContract) []string { return v.CertificateIds }).(pulumi.StringArrayOutput)
}

func (o BackendCredentialsContractOutput) Header() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v BackendCredentialsContract) map[string][]string { return v.Header }).(pulumi.StringArrayMapOutput)
}

func (o BackendCredentialsContractOutput) Query() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v BackendCredentialsContract) map[string][]string { return v.Query }).(pulumi.StringArrayMapOutput)
}

type BackendCredentialsContractPtrOutput struct{ *pulumi.OutputState }

func (BackendCredentialsContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendCredentialsContract)(nil)).Elem()
}

func (o BackendCredentialsContractPtrOutput) ToBackendCredentialsContractPtrOutput() BackendCredentialsContractPtrOutput {
	return o
}

func (o BackendCredentialsContractPtrOutput) ToBackendCredentialsContractPtrOutputWithContext(ctx context.Context) BackendCredentialsContractPtrOutput {
	return o
}

func (o BackendCredentialsContractPtrOutput) Elem() BackendCredentialsContractOutput {
	return o.ApplyT(func(v *BackendCredentialsContract) BackendCredentialsContract {
		if v != nil {
			return *v
		}
		var ret BackendCredentialsContract
		return ret
	}).(BackendCredentialsContractOutput)
}

func (o BackendCredentialsContractPtrOutput) Authorization() BackendAuthorizationHeaderCredentialsPtrOutput {
	return o.ApplyT(func(v *BackendCredentialsContract) *BackendAuthorizationHeaderCredentials {
		if v == nil {
			return nil
		}
		return v.Authorization
	}).(BackendAuthorizationHeaderCredentialsPtrOutput)
}

func (o BackendCredentialsContractPtrOutput) Certificate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendCredentialsContract) []string {
		if v == nil {
			return nil
		}
		return v.Certificate
	}).(pulumi.StringArrayOutput)
}

func (o BackendCredentialsContractPtrOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendCredentialsContract) []string {
		if v == nil {
			return nil
		}
		return v.CertificateIds
	}).(pulumi.StringArrayOutput)
}

func (o BackendCredentialsContractPtrOutput) Header() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *BackendCredentialsContract) map[string][]string {
		if v == nil {
			return nil
		}
		return v.Header
	}).(pulumi.StringArrayMapOutput)
}

func (o BackendCredentialsContractPtrOutput) Query() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *BackendCredentialsContract) map[string][]string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringArrayMapOutput)
}

type BackendCredentialsContractResponse struct {
	Authorization  *BackendAuthorizationHeaderCredentialsResponse `pulumi:"authorization"`
	Certificate    []string                                       `pulumi:"certificate"`
	CertificateIds []string                                       `pulumi:"certificateIds"`
	Header         map[string][]string                            `pulumi:"header"`
	Query          map[string][]string                            `pulumi:"query"`
}

type BackendCredentialsContractResponseOutput struct{ *pulumi.OutputState }

func (BackendCredentialsContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendCredentialsContractResponse)(nil)).Elem()
}

func (o BackendCredentialsContractResponseOutput) ToBackendCredentialsContractResponseOutput() BackendCredentialsContractResponseOutput {
	return o
}

func (o BackendCredentialsContractResponseOutput) ToBackendCredentialsContractResponseOutputWithContext(ctx context.Context) BackendCredentialsContractResponseOutput {
	return o
}

func (o BackendCredentialsContractResponseOutput) Authorization() BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return o.ApplyT(func(v BackendCredentialsContractResponse) *BackendAuthorizationHeaderCredentialsResponse {
		return v.Authorization
	}).(BackendAuthorizationHeaderCredentialsResponsePtrOutput)
}

func (o BackendCredentialsContractResponseOutput) Certificate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackendCredentialsContractResponse) []string { return v.Certificate }).(pulumi.StringArrayOutput)
}

func (o BackendCredentialsContractResponseOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackendCredentialsContractResponse) []string { return v.CertificateIds }).(pulumi.StringArrayOutput)
}

func (o BackendCredentialsContractResponseOutput) Header() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v BackendCredentialsContractResponse) map[string][]string { return v.Header }).(pulumi.StringArrayMapOutput)
}

func (o BackendCredentialsContractResponseOutput) Query() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v BackendCredentialsContractResponse) map[string][]string { return v.Query }).(pulumi.StringArrayMapOutput)
}

type BackendCredentialsContractResponsePtrOutput struct{ *pulumi.OutputState }

func (BackendCredentialsContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendCredentialsContractResponse)(nil)).Elem()
}

func (o BackendCredentialsContractResponsePtrOutput) ToBackendCredentialsContractResponsePtrOutput() BackendCredentialsContractResponsePtrOutput {
	return o
}

func (o BackendCredentialsContractResponsePtrOutput) ToBackendCredentialsContractResponsePtrOutputWithContext(ctx context.Context) BackendCredentialsContractResponsePtrOutput {
	return o
}

func (o BackendCredentialsContractResponsePtrOutput) Elem() BackendCredentialsContractResponseOutput {
	return o.ApplyT(func(v *BackendCredentialsContractResponse) BackendCredentialsContractResponse {
		if v != nil {
			return *v
		}
		var ret BackendCredentialsContractResponse
		return ret
	}).(BackendCredentialsContractResponseOutput)
}

func (o BackendCredentialsContractResponsePtrOutput) Authorization() BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return o.ApplyT(func(v *BackendCredentialsContractResponse) *BackendAuthorizationHeaderCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.Authorization
	}).(BackendAuthorizationHeaderCredentialsResponsePtrOutput)
}

func (o BackendCredentialsContractResponsePtrOutput) Certificate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendCredentialsContractResponse) []string {
		if v == nil {
			return nil
		}
		return v.Certificate
	}).(pulumi.StringArrayOutput)
}

func (o BackendCredentialsContractResponsePtrOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendCredentialsContractResponse) []string {
		if v == nil {
			return nil
		}
		return v.CertificateIds
	}).(pulumi.StringArrayOutput)
}

func (o BackendCredentialsContractResponsePtrOutput) Header() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *BackendCredentialsContractResponse) map[string][]string {
		if v == nil {
			return nil
		}
		return v.Header
	}).(pulumi.StringArrayMapOutput)
}

func (o BackendCredentialsContractResponsePtrOutput) Query() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *BackendCredentialsContractResponse) map[string][]string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringArrayMapOutput)
}

type BackendProperties struct {
	ServiceFabricCluster *BackendServiceFabricClusterProperties `pulumi:"serviceFabricCluster"`
}





type BackendPropertiesInput interface {
	pulumi.Input

	ToBackendPropertiesOutput() BackendPropertiesOutput
	ToBackendPropertiesOutputWithContext(context.Context) BackendPropertiesOutput
}

type BackendPropertiesArgs struct {
	ServiceFabricCluster BackendServiceFabricClusterPropertiesPtrInput `pulumi:"serviceFabricCluster"`
}

func (BackendPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendProperties)(nil)).Elem()
}

func (i BackendPropertiesArgs) ToBackendPropertiesOutput() BackendPropertiesOutput {
	return i.ToBackendPropertiesOutputWithContext(context.Background())
}

func (i BackendPropertiesArgs) ToBackendPropertiesOutputWithContext(ctx context.Context) BackendPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPropertiesOutput)
}

func (i BackendPropertiesArgs) ToBackendPropertiesPtrOutput() BackendPropertiesPtrOutput {
	return i.ToBackendPropertiesPtrOutputWithContext(context.Background())
}

func (i BackendPropertiesArgs) ToBackendPropertiesPtrOutputWithContext(ctx context.Context) BackendPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPropertiesOutput).ToBackendPropertiesPtrOutputWithContext(ctx)
}









type BackendPropertiesPtrInput interface {
	pulumi.Input

	ToBackendPropertiesPtrOutput() BackendPropertiesPtrOutput
	ToBackendPropertiesPtrOutputWithContext(context.Context) BackendPropertiesPtrOutput
}

type backendPropertiesPtrType BackendPropertiesArgs

func BackendPropertiesPtr(v *BackendPropertiesArgs) BackendPropertiesPtrInput {
	return (*backendPropertiesPtrType)(v)
}

func (*backendPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendProperties)(nil)).Elem()
}

func (i *backendPropertiesPtrType) ToBackendPropertiesPtrOutput() BackendPropertiesPtrOutput {
	return i.ToBackendPropertiesPtrOutputWithContext(context.Background())
}

func (i *backendPropertiesPtrType) ToBackendPropertiesPtrOutputWithContext(ctx context.Context) BackendPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPropertiesPtrOutput)
}

type BackendPropertiesOutput struct{ *pulumi.OutputState }

func (BackendPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendProperties)(nil)).Elem()
}

func (o BackendPropertiesOutput) ToBackendPropertiesOutput() BackendPropertiesOutput {
	return o
}

func (o BackendPropertiesOutput) ToBackendPropertiesOutputWithContext(ctx context.Context) BackendPropertiesOutput {
	return o
}

func (o BackendPropertiesOutput) ToBackendPropertiesPtrOutput() BackendPropertiesPtrOutput {
	return o.ToBackendPropertiesPtrOutputWithContext(context.Background())
}

func (o BackendPropertiesOutput) ToBackendPropertiesPtrOutputWithContext(ctx context.Context) BackendPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendProperties) *BackendProperties {
		return &v
	}).(BackendPropertiesPtrOutput)
}

func (o BackendPropertiesOutput) ServiceFabricCluster() BackendServiceFabricClusterPropertiesPtrOutput {
	return o.ApplyT(func(v BackendProperties) *BackendServiceFabricClusterProperties { return v.ServiceFabricCluster }).(BackendServiceFabricClusterPropertiesPtrOutput)
}

type BackendPropertiesPtrOutput struct{ *pulumi.OutputState }

func (BackendPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendProperties)(nil)).Elem()
}

func (o BackendPropertiesPtrOutput) ToBackendPropertiesPtrOutput() BackendPropertiesPtrOutput {
	return o
}

func (o BackendPropertiesPtrOutput) ToBackendPropertiesPtrOutputWithContext(ctx context.Context) BackendPropertiesPtrOutput {
	return o
}

func (o BackendPropertiesPtrOutput) Elem() BackendPropertiesOutput {
	return o.ApplyT(func(v *BackendProperties) BackendProperties {
		if v != nil {
			return *v
		}
		var ret BackendProperties
		return ret
	}).(BackendPropertiesOutput)
}

func (o BackendPropertiesPtrOutput) ServiceFabricCluster() BackendServiceFabricClusterPropertiesPtrOutput {
	return o.ApplyT(func(v *BackendProperties) *BackendServiceFabricClusterProperties {
		if v == nil {
			return nil
		}
		return v.ServiceFabricCluster
	}).(BackendServiceFabricClusterPropertiesPtrOutput)
}

type BackendPropertiesResponse struct {
	ServiceFabricCluster *BackendServiceFabricClusterPropertiesResponse `pulumi:"serviceFabricCluster"`
}

type BackendPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BackendPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPropertiesResponse)(nil)).Elem()
}

func (o BackendPropertiesResponseOutput) ToBackendPropertiesResponseOutput() BackendPropertiesResponseOutput {
	return o
}

func (o BackendPropertiesResponseOutput) ToBackendPropertiesResponseOutputWithContext(ctx context.Context) BackendPropertiesResponseOutput {
	return o
}

func (o BackendPropertiesResponseOutput) ServiceFabricCluster() BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v BackendPropertiesResponse) *BackendServiceFabricClusterPropertiesResponse {
		return v.ServiceFabricCluster
	}).(BackendServiceFabricClusterPropertiesResponsePtrOutput)
}

type BackendProxyContract struct {
	Password *string `pulumi:"password"`
	Url      string  `pulumi:"url"`
	Username *string `pulumi:"username"`
}





type BackendProxyContractInput interface {
	pulumi.Input

	ToBackendProxyContractOutput() BackendProxyContractOutput
	ToBackendProxyContractOutputWithContext(context.Context) BackendProxyContractOutput
}

type BackendProxyContractArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Url      pulumi.StringInput    `pulumi:"url"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (BackendProxyContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendProxyContract)(nil)).Elem()
}

func (i BackendProxyContractArgs) ToBackendProxyContractOutput() BackendProxyContractOutput {
	return i.ToBackendProxyContractOutputWithContext(context.Background())
}

func (i BackendProxyContractArgs) ToBackendProxyContractOutputWithContext(ctx context.Context) BackendProxyContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendProxyContractOutput)
}

func (i BackendProxyContractArgs) ToBackendProxyContractPtrOutput() BackendProxyContractPtrOutput {
	return i.ToBackendProxyContractPtrOutputWithContext(context.Background())
}

func (i BackendProxyContractArgs) ToBackendProxyContractPtrOutputWithContext(ctx context.Context) BackendProxyContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendProxyContractOutput).ToBackendProxyContractPtrOutputWithContext(ctx)
}









type BackendProxyContractPtrInput interface {
	pulumi.Input

	ToBackendProxyContractPtrOutput() BackendProxyContractPtrOutput
	ToBackendProxyContractPtrOutputWithContext(context.Context) BackendProxyContractPtrOutput
}

type backendProxyContractPtrType BackendProxyContractArgs

func BackendProxyContractPtr(v *BackendProxyContractArgs) BackendProxyContractPtrInput {
	return (*backendProxyContractPtrType)(v)
}

func (*backendProxyContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendProxyContract)(nil)).Elem()
}

func (i *backendProxyContractPtrType) ToBackendProxyContractPtrOutput() BackendProxyContractPtrOutput {
	return i.ToBackendProxyContractPtrOutputWithContext(context.Background())
}

func (i *backendProxyContractPtrType) ToBackendProxyContractPtrOutputWithContext(ctx context.Context) BackendProxyContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendProxyContractPtrOutput)
}

type BackendProxyContractOutput struct{ *pulumi.OutputState }

func (BackendProxyContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendProxyContract)(nil)).Elem()
}

func (o BackendProxyContractOutput) ToBackendProxyContractOutput() BackendProxyContractOutput {
	return o
}

func (o BackendProxyContractOutput) ToBackendProxyContractOutputWithContext(ctx context.Context) BackendProxyContractOutput {
	return o
}

func (o BackendProxyContractOutput) ToBackendProxyContractPtrOutput() BackendProxyContractPtrOutput {
	return o.ToBackendProxyContractPtrOutputWithContext(context.Background())
}

func (o BackendProxyContractOutput) ToBackendProxyContractPtrOutputWithContext(ctx context.Context) BackendProxyContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendProxyContract) *BackendProxyContract {
		return &v
	}).(BackendProxyContractPtrOutput)
}

func (o BackendProxyContractOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendProxyContract) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o BackendProxyContractOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v BackendProxyContract) string { return v.Url }).(pulumi.StringOutput)
}

func (o BackendProxyContractOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendProxyContract) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type BackendProxyContractPtrOutput struct{ *pulumi.OutputState }

func (BackendProxyContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendProxyContract)(nil)).Elem()
}

func (o BackendProxyContractPtrOutput) ToBackendProxyContractPtrOutput() BackendProxyContractPtrOutput {
	return o
}

func (o BackendProxyContractPtrOutput) ToBackendProxyContractPtrOutputWithContext(ctx context.Context) BackendProxyContractPtrOutput {
	return o
}

func (o BackendProxyContractPtrOutput) Elem() BackendProxyContractOutput {
	return o.ApplyT(func(v *BackendProxyContract) BackendProxyContract {
		if v != nil {
			return *v
		}
		var ret BackendProxyContract
		return ret
	}).(BackendProxyContractOutput)
}

func (o BackendProxyContractPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendProxyContract) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o BackendProxyContractPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendProxyContract) *string {
		if v == nil {
			return nil
		}
		return &v.Url
	}).(pulumi.StringPtrOutput)
}

func (o BackendProxyContractPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendProxyContract) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type BackendProxyContractResponse struct {
	Password *string `pulumi:"password"`
	Url      string  `pulumi:"url"`
	Username *string `pulumi:"username"`
}

type BackendProxyContractResponseOutput struct{ *pulumi.OutputState }

func (BackendProxyContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendProxyContractResponse)(nil)).Elem()
}

func (o BackendProxyContractResponseOutput) ToBackendProxyContractResponseOutput() BackendProxyContractResponseOutput {
	return o
}

func (o BackendProxyContractResponseOutput) ToBackendProxyContractResponseOutputWithContext(ctx context.Context) BackendProxyContractResponseOutput {
	return o
}

func (o BackendProxyContractResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendProxyContractResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o BackendProxyContractResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v BackendProxyContractResponse) string { return v.Url }).(pulumi.StringOutput)
}

func (o BackendProxyContractResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendProxyContractResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type BackendProxyContractResponsePtrOutput struct{ *pulumi.OutputState }

func (BackendProxyContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendProxyContractResponse)(nil)).Elem()
}

func (o BackendProxyContractResponsePtrOutput) ToBackendProxyContractResponsePtrOutput() BackendProxyContractResponsePtrOutput {
	return o
}

func (o BackendProxyContractResponsePtrOutput) ToBackendProxyContractResponsePtrOutputWithContext(ctx context.Context) BackendProxyContractResponsePtrOutput {
	return o
}

func (o BackendProxyContractResponsePtrOutput) Elem() BackendProxyContractResponseOutput {
	return o.ApplyT(func(v *BackendProxyContractResponse) BackendProxyContractResponse {
		if v != nil {
			return *v
		}
		var ret BackendProxyContractResponse
		return ret
	}).(BackendProxyContractResponseOutput)
}

func (o BackendProxyContractResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendProxyContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o BackendProxyContractResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendProxyContractResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Url
	}).(pulumi.StringPtrOutput)
}

func (o BackendProxyContractResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendProxyContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type BackendServiceFabricClusterProperties struct {
	ClientCertificateId           *string               `pulumi:"clientCertificateId"`
	ClientCertificatethumbprint   *string               `pulumi:"clientCertificatethumbprint"`
	ManagementEndpoints           []string              `pulumi:"managementEndpoints"`
	MaxPartitionResolutionRetries *int                  `pulumi:"maxPartitionResolutionRetries"`
	ServerCertificateThumbprints  []string              `pulumi:"serverCertificateThumbprints"`
	ServerX509Names               []X509CertificateName `pulumi:"serverX509Names"`
}





type BackendServiceFabricClusterPropertiesInput interface {
	pulumi.Input

	ToBackendServiceFabricClusterPropertiesOutput() BackendServiceFabricClusterPropertiesOutput
	ToBackendServiceFabricClusterPropertiesOutputWithContext(context.Context) BackendServiceFabricClusterPropertiesOutput
}

type BackendServiceFabricClusterPropertiesArgs struct {
	ClientCertificateId           pulumi.StringPtrInput         `pulumi:"clientCertificateId"`
	ClientCertificatethumbprint   pulumi.StringPtrInput         `pulumi:"clientCertificatethumbprint"`
	ManagementEndpoints           pulumi.StringArrayInput       `pulumi:"managementEndpoints"`
	MaxPartitionResolutionRetries pulumi.IntPtrInput            `pulumi:"maxPartitionResolutionRetries"`
	ServerCertificateThumbprints  pulumi.StringArrayInput       `pulumi:"serverCertificateThumbprints"`
	ServerX509Names               X509CertificateNameArrayInput `pulumi:"serverX509Names"`
}

func (BackendServiceFabricClusterPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceFabricClusterProperties)(nil)).Elem()
}

func (i BackendServiceFabricClusterPropertiesArgs) ToBackendServiceFabricClusterPropertiesOutput() BackendServiceFabricClusterPropertiesOutput {
	return i.ToBackendServiceFabricClusterPropertiesOutputWithContext(context.Background())
}

func (i BackendServiceFabricClusterPropertiesArgs) ToBackendServiceFabricClusterPropertiesOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceFabricClusterPropertiesOutput)
}

func (i BackendServiceFabricClusterPropertiesArgs) ToBackendServiceFabricClusterPropertiesPtrOutput() BackendServiceFabricClusterPropertiesPtrOutput {
	return i.ToBackendServiceFabricClusterPropertiesPtrOutputWithContext(context.Background())
}

func (i BackendServiceFabricClusterPropertiesArgs) ToBackendServiceFabricClusterPropertiesPtrOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceFabricClusterPropertiesOutput).ToBackendServiceFabricClusterPropertiesPtrOutputWithContext(ctx)
}









type BackendServiceFabricClusterPropertiesPtrInput interface {
	pulumi.Input

	ToBackendServiceFabricClusterPropertiesPtrOutput() BackendServiceFabricClusterPropertiesPtrOutput
	ToBackendServiceFabricClusterPropertiesPtrOutputWithContext(context.Context) BackendServiceFabricClusterPropertiesPtrOutput
}

type backendServiceFabricClusterPropertiesPtrType BackendServiceFabricClusterPropertiesArgs

func BackendServiceFabricClusterPropertiesPtr(v *BackendServiceFabricClusterPropertiesArgs) BackendServiceFabricClusterPropertiesPtrInput {
	return (*backendServiceFabricClusterPropertiesPtrType)(v)
}

func (*backendServiceFabricClusterPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceFabricClusterProperties)(nil)).Elem()
}

func (i *backendServiceFabricClusterPropertiesPtrType) ToBackendServiceFabricClusterPropertiesPtrOutput() BackendServiceFabricClusterPropertiesPtrOutput {
	return i.ToBackendServiceFabricClusterPropertiesPtrOutputWithContext(context.Background())
}

func (i *backendServiceFabricClusterPropertiesPtrType) ToBackendServiceFabricClusterPropertiesPtrOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceFabricClusterPropertiesPtrOutput)
}

type BackendServiceFabricClusterPropertiesOutput struct{ *pulumi.OutputState }

func (BackendServiceFabricClusterPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceFabricClusterProperties)(nil)).Elem()
}

func (o BackendServiceFabricClusterPropertiesOutput) ToBackendServiceFabricClusterPropertiesOutput() BackendServiceFabricClusterPropertiesOutput {
	return o
}

func (o BackendServiceFabricClusterPropertiesOutput) ToBackendServiceFabricClusterPropertiesOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesOutput {
	return o
}

func (o BackendServiceFabricClusterPropertiesOutput) ToBackendServiceFabricClusterPropertiesPtrOutput() BackendServiceFabricClusterPropertiesPtrOutput {
	return o.ToBackendServiceFabricClusterPropertiesPtrOutputWithContext(context.Background())
}

func (o BackendServiceFabricClusterPropertiesOutput) ToBackendServiceFabricClusterPropertiesPtrOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendServiceFabricClusterProperties) *BackendServiceFabricClusterProperties {
		return &v
	}).(BackendServiceFabricClusterPropertiesPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesOutput) ClientCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterProperties) *string { return v.ClientCertificateId }).(pulumi.StringPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesOutput) ClientCertificatethumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterProperties) *string { return v.ClientCertificatethumbprint }).(pulumi.StringPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesOutput) ManagementEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterProperties) []string { return v.ManagementEndpoints }).(pulumi.StringArrayOutput)
}

func (o BackendServiceFabricClusterPropertiesOutput) MaxPartitionResolutionRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterProperties) *int { return v.MaxPartitionResolutionRetries }).(pulumi.IntPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesOutput) ServerCertificateThumbprints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterProperties) []string { return v.ServerCertificateThumbprints }).(pulumi.StringArrayOutput)
}

func (o BackendServiceFabricClusterPropertiesOutput) ServerX509Names() X509CertificateNameArrayOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterProperties) []X509CertificateName { return v.ServerX509Names }).(X509CertificateNameArrayOutput)
}

type BackendServiceFabricClusterPropertiesPtrOutput struct{ *pulumi.OutputState }

func (BackendServiceFabricClusterPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceFabricClusterProperties)(nil)).Elem()
}

func (o BackendServiceFabricClusterPropertiesPtrOutput) ToBackendServiceFabricClusterPropertiesPtrOutput() BackendServiceFabricClusterPropertiesPtrOutput {
	return o
}

func (o BackendServiceFabricClusterPropertiesPtrOutput) ToBackendServiceFabricClusterPropertiesPtrOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesPtrOutput {
	return o
}

func (o BackendServiceFabricClusterPropertiesPtrOutput) Elem() BackendServiceFabricClusterPropertiesOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterProperties) BackendServiceFabricClusterProperties {
		if v != nil {
			return *v
		}
		var ret BackendServiceFabricClusterProperties
		return ret
	}).(BackendServiceFabricClusterPropertiesOutput)
}

func (o BackendServiceFabricClusterPropertiesPtrOutput) ClientCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientCertificateId
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesPtrOutput) ClientCertificatethumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientCertificatethumbprint
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesPtrOutput) ManagementEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterProperties) []string {
		if v == nil {
			return nil
		}
		return v.ManagementEndpoints
	}).(pulumi.StringArrayOutput)
}

func (o BackendServiceFabricClusterPropertiesPtrOutput) MaxPartitionResolutionRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterProperties) *int {
		if v == nil {
			return nil
		}
		return v.MaxPartitionResolutionRetries
	}).(pulumi.IntPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesPtrOutput) ServerCertificateThumbprints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterProperties) []string {
		if v == nil {
			return nil
		}
		return v.ServerCertificateThumbprints
	}).(pulumi.StringArrayOutput)
}

func (o BackendServiceFabricClusterPropertiesPtrOutput) ServerX509Names() X509CertificateNameArrayOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterProperties) []X509CertificateName {
		if v == nil {
			return nil
		}
		return v.ServerX509Names
	}).(X509CertificateNameArrayOutput)
}

type BackendServiceFabricClusterPropertiesResponse struct {
	ClientCertificateId           *string                       `pulumi:"clientCertificateId"`
	ClientCertificatethumbprint   *string                       `pulumi:"clientCertificatethumbprint"`
	ManagementEndpoints           []string                      `pulumi:"managementEndpoints"`
	MaxPartitionResolutionRetries *int                          `pulumi:"maxPartitionResolutionRetries"`
	ServerCertificateThumbprints  []string                      `pulumi:"serverCertificateThumbprints"`
	ServerX509Names               []X509CertificateNameResponse `pulumi:"serverX509Names"`
}

type BackendServiceFabricClusterPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BackendServiceFabricClusterPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceFabricClusterPropertiesResponse)(nil)).Elem()
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) ToBackendServiceFabricClusterPropertiesResponseOutput() BackendServiceFabricClusterPropertiesResponseOutput {
	return o
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) ToBackendServiceFabricClusterPropertiesResponseOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesResponseOutput {
	return o
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) ClientCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterPropertiesResponse) *string { return v.ClientCertificateId }).(pulumi.StringPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) ClientCertificatethumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterPropertiesResponse) *string { return v.ClientCertificatethumbprint }).(pulumi.StringPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) ManagementEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterPropertiesResponse) []string { return v.ManagementEndpoints }).(pulumi.StringArrayOutput)
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) MaxPartitionResolutionRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterPropertiesResponse) *int { return v.MaxPartitionResolutionRetries }).(pulumi.IntPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) ServerCertificateThumbprints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterPropertiesResponse) []string { return v.ServerCertificateThumbprints }).(pulumi.StringArrayOutput)
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) ServerX509Names() X509CertificateNameResponseArrayOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterPropertiesResponse) []X509CertificateNameResponse {
		return v.ServerX509Names
	}).(X509CertificateNameResponseArrayOutput)
}

type BackendServiceFabricClusterPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (BackendServiceFabricClusterPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceFabricClusterPropertiesResponse)(nil)).Elem()
}

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) ToBackendServiceFabricClusterPropertiesResponsePtrOutput() BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return o
}

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) ToBackendServiceFabricClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return o
}

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) Elem() BackendServiceFabricClusterPropertiesResponseOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterPropertiesResponse) BackendServiceFabricClusterPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret BackendServiceFabricClusterPropertiesResponse
		return ret
	}).(BackendServiceFabricClusterPropertiesResponseOutput)
}

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) ClientCertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientCertificateId
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) ClientCertificatethumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientCertificatethumbprint
	}).(pulumi.StringPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) ManagementEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ManagementEndpoints
	}).(pulumi.StringArrayOutput)
}

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) MaxPartitionResolutionRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxPartitionResolutionRetries
	}).(pulumi.IntPtrOutput)
}

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) ServerCertificateThumbprints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ServerCertificateThumbprints
	}).(pulumi.StringArrayOutput)
}

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) ServerX509Names() X509CertificateNameResponseArrayOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterPropertiesResponse) []X509CertificateNameResponse {
		if v == nil {
			return nil
		}
		return v.ServerX509Names
	}).(X509CertificateNameResponseArrayOutput)
}

type BackendTlsProperties struct {
	ValidateCertificateChain *bool `pulumi:"validateCertificateChain"`
	ValidateCertificateName  *bool `pulumi:"validateCertificateName"`
}


func (val *BackendTlsProperties) Defaults() *BackendTlsProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ValidateCertificateChain) {
		validateCertificateChain_ := true
		tmp.ValidateCertificateChain = &validateCertificateChain_
	}
	if isZero(tmp.ValidateCertificateName) {
		validateCertificateName_ := true
		tmp.ValidateCertificateName = &validateCertificateName_
	}
	return &tmp
}





type BackendTlsPropertiesInput interface {
	pulumi.Input

	ToBackendTlsPropertiesOutput() BackendTlsPropertiesOutput
	ToBackendTlsPropertiesOutputWithContext(context.Context) BackendTlsPropertiesOutput
}

type BackendTlsPropertiesArgs struct {
	ValidateCertificateChain pulumi.BoolPtrInput `pulumi:"validateCertificateChain"`
	ValidateCertificateName  pulumi.BoolPtrInput `pulumi:"validateCertificateName"`
}

func (BackendTlsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendTlsProperties)(nil)).Elem()
}

func (i BackendTlsPropertiesArgs) ToBackendTlsPropertiesOutput() BackendTlsPropertiesOutput {
	return i.ToBackendTlsPropertiesOutputWithContext(context.Background())
}

func (i BackendTlsPropertiesArgs) ToBackendTlsPropertiesOutputWithContext(ctx context.Context) BackendTlsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendTlsPropertiesOutput)
}

func (i BackendTlsPropertiesArgs) ToBackendTlsPropertiesPtrOutput() BackendTlsPropertiesPtrOutput {
	return i.ToBackendTlsPropertiesPtrOutputWithContext(context.Background())
}

func (i BackendTlsPropertiesArgs) ToBackendTlsPropertiesPtrOutputWithContext(ctx context.Context) BackendTlsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendTlsPropertiesOutput).ToBackendTlsPropertiesPtrOutputWithContext(ctx)
}









type BackendTlsPropertiesPtrInput interface {
	pulumi.Input

	ToBackendTlsPropertiesPtrOutput() BackendTlsPropertiesPtrOutput
	ToBackendTlsPropertiesPtrOutputWithContext(context.Context) BackendTlsPropertiesPtrOutput
}

type backendTlsPropertiesPtrType BackendTlsPropertiesArgs

func BackendTlsPropertiesPtr(v *BackendTlsPropertiesArgs) BackendTlsPropertiesPtrInput {
	return (*backendTlsPropertiesPtrType)(v)
}

func (*backendTlsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendTlsProperties)(nil)).Elem()
}

func (i *backendTlsPropertiesPtrType) ToBackendTlsPropertiesPtrOutput() BackendTlsPropertiesPtrOutput {
	return i.ToBackendTlsPropertiesPtrOutputWithContext(context.Background())
}

func (i *backendTlsPropertiesPtrType) ToBackendTlsPropertiesPtrOutputWithContext(ctx context.Context) BackendTlsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendTlsPropertiesPtrOutput)
}

type BackendTlsPropertiesOutput struct{ *pulumi.OutputState }

func (BackendTlsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendTlsProperties)(nil)).Elem()
}

func (o BackendTlsPropertiesOutput) ToBackendTlsPropertiesOutput() BackendTlsPropertiesOutput {
	return o
}

func (o BackendTlsPropertiesOutput) ToBackendTlsPropertiesOutputWithContext(ctx context.Context) BackendTlsPropertiesOutput {
	return o
}

func (o BackendTlsPropertiesOutput) ToBackendTlsPropertiesPtrOutput() BackendTlsPropertiesPtrOutput {
	return o.ToBackendTlsPropertiesPtrOutputWithContext(context.Background())
}

func (o BackendTlsPropertiesOutput) ToBackendTlsPropertiesPtrOutputWithContext(ctx context.Context) BackendTlsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendTlsProperties) *BackendTlsProperties {
		return &v
	}).(BackendTlsPropertiesPtrOutput)
}

func (o BackendTlsPropertiesOutput) ValidateCertificateChain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BackendTlsProperties) *bool { return v.ValidateCertificateChain }).(pulumi.BoolPtrOutput)
}

func (o BackendTlsPropertiesOutput) ValidateCertificateName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BackendTlsProperties) *bool { return v.ValidateCertificateName }).(pulumi.BoolPtrOutput)
}

type BackendTlsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (BackendTlsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendTlsProperties)(nil)).Elem()
}

func (o BackendTlsPropertiesPtrOutput) ToBackendTlsPropertiesPtrOutput() BackendTlsPropertiesPtrOutput {
	return o
}

func (o BackendTlsPropertiesPtrOutput) ToBackendTlsPropertiesPtrOutputWithContext(ctx context.Context) BackendTlsPropertiesPtrOutput {
	return o
}

func (o BackendTlsPropertiesPtrOutput) Elem() BackendTlsPropertiesOutput {
	return o.ApplyT(func(v *BackendTlsProperties) BackendTlsProperties {
		if v != nil {
			return *v
		}
		var ret BackendTlsProperties
		return ret
	}).(BackendTlsPropertiesOutput)
}

func (o BackendTlsPropertiesPtrOutput) ValidateCertificateChain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackendTlsProperties) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateCertificateChain
	}).(pulumi.BoolPtrOutput)
}

func (o BackendTlsPropertiesPtrOutput) ValidateCertificateName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackendTlsProperties) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateCertificateName
	}).(pulumi.BoolPtrOutput)
}

type BackendTlsPropertiesResponse struct {
	ValidateCertificateChain *bool `pulumi:"validateCertificateChain"`
	ValidateCertificateName  *bool `pulumi:"validateCertificateName"`
}


func (val *BackendTlsPropertiesResponse) Defaults() *BackendTlsPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ValidateCertificateChain) {
		validateCertificateChain_ := true
		tmp.ValidateCertificateChain = &validateCertificateChain_
	}
	if isZero(tmp.ValidateCertificateName) {
		validateCertificateName_ := true
		tmp.ValidateCertificateName = &validateCertificateName_
	}
	return &tmp
}

type BackendTlsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BackendTlsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendTlsPropertiesResponse)(nil)).Elem()
}

func (o BackendTlsPropertiesResponseOutput) ToBackendTlsPropertiesResponseOutput() BackendTlsPropertiesResponseOutput {
	return o
}

func (o BackendTlsPropertiesResponseOutput) ToBackendTlsPropertiesResponseOutputWithContext(ctx context.Context) BackendTlsPropertiesResponseOutput {
	return o
}

func (o BackendTlsPropertiesResponseOutput) ValidateCertificateChain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BackendTlsPropertiesResponse) *bool { return v.ValidateCertificateChain }).(pulumi.BoolPtrOutput)
}

func (o BackendTlsPropertiesResponseOutput) ValidateCertificateName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BackendTlsPropertiesResponse) *bool { return v.ValidateCertificateName }).(pulumi.BoolPtrOutput)
}

type BackendTlsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (BackendTlsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendTlsPropertiesResponse)(nil)).Elem()
}

func (o BackendTlsPropertiesResponsePtrOutput) ToBackendTlsPropertiesResponsePtrOutput() BackendTlsPropertiesResponsePtrOutput {
	return o
}

func (o BackendTlsPropertiesResponsePtrOutput) ToBackendTlsPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendTlsPropertiesResponsePtrOutput {
	return o
}

func (o BackendTlsPropertiesResponsePtrOutput) Elem() BackendTlsPropertiesResponseOutput {
	return o.ApplyT(func(v *BackendTlsPropertiesResponse) BackendTlsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret BackendTlsPropertiesResponse
		return ret
	}).(BackendTlsPropertiesResponseOutput)
}

func (o BackendTlsPropertiesResponsePtrOutput) ValidateCertificateChain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackendTlsPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateCertificateChain
	}).(pulumi.BoolPtrOutput)
}

func (o BackendTlsPropertiesResponsePtrOutput) ValidateCertificateName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackendTlsPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateCertificateName
	}).(pulumi.BoolPtrOutput)
}

type BodyDiagnosticSettings struct {
	Bytes *int `pulumi:"bytes"`
}





type BodyDiagnosticSettingsInput interface {
	pulumi.Input

	ToBodyDiagnosticSettingsOutput() BodyDiagnosticSettingsOutput
	ToBodyDiagnosticSettingsOutputWithContext(context.Context) BodyDiagnosticSettingsOutput
}

type BodyDiagnosticSettingsArgs struct {
	Bytes pulumi.IntPtrInput `pulumi:"bytes"`
}

func (BodyDiagnosticSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BodyDiagnosticSettings)(nil)).Elem()
}

func (i BodyDiagnosticSettingsArgs) ToBodyDiagnosticSettingsOutput() BodyDiagnosticSettingsOutput {
	return i.ToBodyDiagnosticSettingsOutputWithContext(context.Background())
}

func (i BodyDiagnosticSettingsArgs) ToBodyDiagnosticSettingsOutputWithContext(ctx context.Context) BodyDiagnosticSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BodyDiagnosticSettingsOutput)
}

func (i BodyDiagnosticSettingsArgs) ToBodyDiagnosticSettingsPtrOutput() BodyDiagnosticSettingsPtrOutput {
	return i.ToBodyDiagnosticSettingsPtrOutputWithContext(context.Background())
}

func (i BodyDiagnosticSettingsArgs) ToBodyDiagnosticSettingsPtrOutputWithContext(ctx context.Context) BodyDiagnosticSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BodyDiagnosticSettingsOutput).ToBodyDiagnosticSettingsPtrOutputWithContext(ctx)
}









type BodyDiagnosticSettingsPtrInput interface {
	pulumi.Input

	ToBodyDiagnosticSettingsPtrOutput() BodyDiagnosticSettingsPtrOutput
	ToBodyDiagnosticSettingsPtrOutputWithContext(context.Context) BodyDiagnosticSettingsPtrOutput
}

type bodyDiagnosticSettingsPtrType BodyDiagnosticSettingsArgs

func BodyDiagnosticSettingsPtr(v *BodyDiagnosticSettingsArgs) BodyDiagnosticSettingsPtrInput {
	return (*bodyDiagnosticSettingsPtrType)(v)
}

func (*bodyDiagnosticSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BodyDiagnosticSettings)(nil)).Elem()
}

func (i *bodyDiagnosticSettingsPtrType) ToBodyDiagnosticSettingsPtrOutput() BodyDiagnosticSettingsPtrOutput {
	return i.ToBodyDiagnosticSettingsPtrOutputWithContext(context.Background())
}

func (i *bodyDiagnosticSettingsPtrType) ToBodyDiagnosticSettingsPtrOutputWithContext(ctx context.Context) BodyDiagnosticSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BodyDiagnosticSettingsPtrOutput)
}

type BodyDiagnosticSettingsOutput struct{ *pulumi.OutputState }

func (BodyDiagnosticSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BodyDiagnosticSettings)(nil)).Elem()
}

func (o BodyDiagnosticSettingsOutput) ToBodyDiagnosticSettingsOutput() BodyDiagnosticSettingsOutput {
	return o
}

func (o BodyDiagnosticSettingsOutput) ToBodyDiagnosticSettingsOutputWithContext(ctx context.Context) BodyDiagnosticSettingsOutput {
	return o
}

func (o BodyDiagnosticSettingsOutput) ToBodyDiagnosticSettingsPtrOutput() BodyDiagnosticSettingsPtrOutput {
	return o.ToBodyDiagnosticSettingsPtrOutputWithContext(context.Background())
}

func (o BodyDiagnosticSettingsOutput) ToBodyDiagnosticSettingsPtrOutputWithContext(ctx context.Context) BodyDiagnosticSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BodyDiagnosticSettings) *BodyDiagnosticSettings {
		return &v
	}).(BodyDiagnosticSettingsPtrOutput)
}

func (o BodyDiagnosticSettingsOutput) Bytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BodyDiagnosticSettings) *int { return v.Bytes }).(pulumi.IntPtrOutput)
}

type BodyDiagnosticSettingsPtrOutput struct{ *pulumi.OutputState }

func (BodyDiagnosticSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BodyDiagnosticSettings)(nil)).Elem()
}

func (o BodyDiagnosticSettingsPtrOutput) ToBodyDiagnosticSettingsPtrOutput() BodyDiagnosticSettingsPtrOutput {
	return o
}

func (o BodyDiagnosticSettingsPtrOutput) ToBodyDiagnosticSettingsPtrOutputWithContext(ctx context.Context) BodyDiagnosticSettingsPtrOutput {
	return o
}

func (o BodyDiagnosticSettingsPtrOutput) Elem() BodyDiagnosticSettingsOutput {
	return o.ApplyT(func(v *BodyDiagnosticSettings) BodyDiagnosticSettings {
		if v != nil {
			return *v
		}
		var ret BodyDiagnosticSettings
		return ret
	}).(BodyDiagnosticSettingsOutput)
}

func (o BodyDiagnosticSettingsPtrOutput) Bytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BodyDiagnosticSettings) *int {
		if v == nil {
			return nil
		}
		return v.Bytes
	}).(pulumi.IntPtrOutput)
}

type BodyDiagnosticSettingsResponse struct {
	Bytes *int `pulumi:"bytes"`
}

type BodyDiagnosticSettingsResponseOutput struct{ *pulumi.OutputState }

func (BodyDiagnosticSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BodyDiagnosticSettingsResponse)(nil)).Elem()
}

func (o BodyDiagnosticSettingsResponseOutput) ToBodyDiagnosticSettingsResponseOutput() BodyDiagnosticSettingsResponseOutput {
	return o
}

func (o BodyDiagnosticSettingsResponseOutput) ToBodyDiagnosticSettingsResponseOutputWithContext(ctx context.Context) BodyDiagnosticSettingsResponseOutput {
	return o
}

func (o BodyDiagnosticSettingsResponseOutput) Bytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BodyDiagnosticSettingsResponse) *int { return v.Bytes }).(pulumi.IntPtrOutput)
}

type BodyDiagnosticSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (BodyDiagnosticSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BodyDiagnosticSettingsResponse)(nil)).Elem()
}

func (o BodyDiagnosticSettingsResponsePtrOutput) ToBodyDiagnosticSettingsResponsePtrOutput() BodyDiagnosticSettingsResponsePtrOutput {
	return o
}

func (o BodyDiagnosticSettingsResponsePtrOutput) ToBodyDiagnosticSettingsResponsePtrOutputWithContext(ctx context.Context) BodyDiagnosticSettingsResponsePtrOutput {
	return o
}

func (o BodyDiagnosticSettingsResponsePtrOutput) Elem() BodyDiagnosticSettingsResponseOutput {
	return o.ApplyT(func(v *BodyDiagnosticSettingsResponse) BodyDiagnosticSettingsResponse {
		if v != nil {
			return *v
		}
		var ret BodyDiagnosticSettingsResponse
		return ret
	}).(BodyDiagnosticSettingsResponseOutput)
}

func (o BodyDiagnosticSettingsResponsePtrOutput) Bytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BodyDiagnosticSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Bytes
	}).(pulumi.IntPtrOutput)
}

type CertificateConfiguration struct {
	Certificate         *CertificateInformation `pulumi:"certificate"`
	CertificatePassword *string                 `pulumi:"certificatePassword"`
	EncodedCertificate  *string                 `pulumi:"encodedCertificate"`
	StoreName           string                  `pulumi:"storeName"`
}





type CertificateConfigurationInput interface {
	pulumi.Input

	ToCertificateConfigurationOutput() CertificateConfigurationOutput
	ToCertificateConfigurationOutputWithContext(context.Context) CertificateConfigurationOutput
}

type CertificateConfigurationArgs struct {
	Certificate         CertificateInformationPtrInput `pulumi:"certificate"`
	CertificatePassword pulumi.StringPtrInput          `pulumi:"certificatePassword"`
	EncodedCertificate  pulumi.StringPtrInput          `pulumi:"encodedCertificate"`
	StoreName           pulumi.StringInput             `pulumi:"storeName"`
}

func (CertificateConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateConfiguration)(nil)).Elem()
}

func (i CertificateConfigurationArgs) ToCertificateConfigurationOutput() CertificateConfigurationOutput {
	return i.ToCertificateConfigurationOutputWithContext(context.Background())
}

func (i CertificateConfigurationArgs) ToCertificateConfigurationOutputWithContext(ctx context.Context) CertificateConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateConfigurationOutput)
}





type CertificateConfigurationArrayInput interface {
	pulumi.Input

	ToCertificateConfigurationArrayOutput() CertificateConfigurationArrayOutput
	ToCertificateConfigurationArrayOutputWithContext(context.Context) CertificateConfigurationArrayOutput
}

type CertificateConfigurationArray []CertificateConfigurationInput

func (CertificateConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateConfiguration)(nil)).Elem()
}

func (i CertificateConfigurationArray) ToCertificateConfigurationArrayOutput() CertificateConfigurationArrayOutput {
	return i.ToCertificateConfigurationArrayOutputWithContext(context.Background())
}

func (i CertificateConfigurationArray) ToCertificateConfigurationArrayOutputWithContext(ctx context.Context) CertificateConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateConfigurationArrayOutput)
}

type CertificateConfigurationOutput struct{ *pulumi.OutputState }

func (CertificateConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateConfiguration)(nil)).Elem()
}

func (o CertificateConfigurationOutput) ToCertificateConfigurationOutput() CertificateConfigurationOutput {
	return o
}

func (o CertificateConfigurationOutput) ToCertificateConfigurationOutputWithContext(ctx context.Context) CertificateConfigurationOutput {
	return o
}

func (o CertificateConfigurationOutput) Certificate() CertificateInformationPtrOutput {
	return o.ApplyT(func(v CertificateConfiguration) *CertificateInformation { return v.Certificate }).(CertificateInformationPtrOutput)
}

func (o CertificateConfigurationOutput) CertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateConfiguration) *string { return v.CertificatePassword }).(pulumi.StringPtrOutput)
}

func (o CertificateConfigurationOutput) EncodedCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateConfiguration) *string { return v.EncodedCertificate }).(pulumi.StringPtrOutput)
}

func (o CertificateConfigurationOutput) StoreName() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateConfiguration) string { return v.StoreName }).(pulumi.StringOutput)
}

type CertificateConfigurationArrayOutput struct{ *pulumi.OutputState }

func (CertificateConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateConfiguration)(nil)).Elem()
}

func (o CertificateConfigurationArrayOutput) ToCertificateConfigurationArrayOutput() CertificateConfigurationArrayOutput {
	return o
}

func (o CertificateConfigurationArrayOutput) ToCertificateConfigurationArrayOutputWithContext(ctx context.Context) CertificateConfigurationArrayOutput {
	return o
}

func (o CertificateConfigurationArrayOutput) Index(i pulumi.IntInput) CertificateConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateConfiguration {
		return vs[0].([]CertificateConfiguration)[vs[1].(int)]
	}).(CertificateConfigurationOutput)
}

type CertificateConfigurationResponse struct {
	Certificate         *CertificateInformationResponse `pulumi:"certificate"`
	CertificatePassword *string                         `pulumi:"certificatePassword"`
	EncodedCertificate  *string                         `pulumi:"encodedCertificate"`
	StoreName           string                          `pulumi:"storeName"`
}

type CertificateConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CertificateConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateConfigurationResponse)(nil)).Elem()
}

func (o CertificateConfigurationResponseOutput) ToCertificateConfigurationResponseOutput() CertificateConfigurationResponseOutput {
	return o
}

func (o CertificateConfigurationResponseOutput) ToCertificateConfigurationResponseOutputWithContext(ctx context.Context) CertificateConfigurationResponseOutput {
	return o
}

func (o CertificateConfigurationResponseOutput) Certificate() CertificateInformationResponsePtrOutput {
	return o.ApplyT(func(v CertificateConfigurationResponse) *CertificateInformationResponse { return v.Certificate }).(CertificateInformationResponsePtrOutput)
}

func (o CertificateConfigurationResponseOutput) CertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateConfigurationResponse) *string { return v.CertificatePassword }).(pulumi.StringPtrOutput)
}

func (o CertificateConfigurationResponseOutput) EncodedCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateConfigurationResponse) *string { return v.EncodedCertificate }).(pulumi.StringPtrOutput)
}

func (o CertificateConfigurationResponseOutput) StoreName() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateConfigurationResponse) string { return v.StoreName }).(pulumi.StringOutput)
}

type CertificateConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (CertificateConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateConfigurationResponse)(nil)).Elem()
}

func (o CertificateConfigurationResponseArrayOutput) ToCertificateConfigurationResponseArrayOutput() CertificateConfigurationResponseArrayOutput {
	return o
}

func (o CertificateConfigurationResponseArrayOutput) ToCertificateConfigurationResponseArrayOutputWithContext(ctx context.Context) CertificateConfigurationResponseArrayOutput {
	return o
}

func (o CertificateConfigurationResponseArrayOutput) Index(i pulumi.IntInput) CertificateConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateConfigurationResponse {
		return vs[0].([]CertificateConfigurationResponse)[vs[1].(int)]
	}).(CertificateConfigurationResponseOutput)
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

func (i CertificateInformationArgs) ToCertificateInformationPtrOutput() CertificateInformationPtrOutput {
	return i.ToCertificateInformationPtrOutputWithContext(context.Background())
}

func (i CertificateInformationArgs) ToCertificateInformationPtrOutputWithContext(ctx context.Context) CertificateInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateInformationOutput).ToCertificateInformationPtrOutputWithContext(ctx)
}









type CertificateInformationPtrInput interface {
	pulumi.Input

	ToCertificateInformationPtrOutput() CertificateInformationPtrOutput
	ToCertificateInformationPtrOutputWithContext(context.Context) CertificateInformationPtrOutput
}

type certificateInformationPtrType CertificateInformationArgs

func CertificateInformationPtr(v *CertificateInformationArgs) CertificateInformationPtrInput {
	return (*certificateInformationPtrType)(v)
}

func (*certificateInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateInformation)(nil)).Elem()
}

func (i *certificateInformationPtrType) ToCertificateInformationPtrOutput() CertificateInformationPtrOutput {
	return i.ToCertificateInformationPtrOutputWithContext(context.Background())
}

func (i *certificateInformationPtrType) ToCertificateInformationPtrOutputWithContext(ctx context.Context) CertificateInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateInformationPtrOutput)
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

func (o CertificateInformationOutput) ToCertificateInformationPtrOutput() CertificateInformationPtrOutput {
	return o.ToCertificateInformationPtrOutputWithContext(context.Background())
}

func (o CertificateInformationOutput) ToCertificateInformationPtrOutputWithContext(ctx context.Context) CertificateInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateInformation) *CertificateInformation {
		return &v
	}).(CertificateInformationPtrOutput)
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

type CertificateInformationPtrOutput struct{ *pulumi.OutputState }

func (CertificateInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateInformation)(nil)).Elem()
}

func (o CertificateInformationPtrOutput) ToCertificateInformationPtrOutput() CertificateInformationPtrOutput {
	return o
}

func (o CertificateInformationPtrOutput) ToCertificateInformationPtrOutputWithContext(ctx context.Context) CertificateInformationPtrOutput {
	return o
}

func (o CertificateInformationPtrOutput) Elem() CertificateInformationOutput {
	return o.ApplyT(func(v *CertificateInformation) CertificateInformation {
		if v != nil {
			return *v
		}
		var ret CertificateInformation
		return ret
	}).(CertificateInformationOutput)
}

func (o CertificateInformationPtrOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateInformation) *string {
		if v == nil {
			return nil
		}
		return &v.Expiry
	}).(pulumi.StringPtrOutput)
}

func (o CertificateInformationPtrOutput) Subject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateInformation) *string {
		if v == nil {
			return nil
		}
		return &v.Subject
	}).(pulumi.StringPtrOutput)
}

func (o CertificateInformationPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateInformation) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type CertificateInformationResponse struct {
	Expiry     string `pulumi:"expiry"`
	Subject    string `pulumi:"subject"`
	Thumbprint string `pulumi:"thumbprint"`
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

type CertificateInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (CertificateInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateInformationResponse)(nil)).Elem()
}

func (o CertificateInformationResponsePtrOutput) ToCertificateInformationResponsePtrOutput() CertificateInformationResponsePtrOutput {
	return o
}

func (o CertificateInformationResponsePtrOutput) ToCertificateInformationResponsePtrOutputWithContext(ctx context.Context) CertificateInformationResponsePtrOutput {
	return o
}

func (o CertificateInformationResponsePtrOutput) Elem() CertificateInformationResponseOutput {
	return o.ApplyT(func(v *CertificateInformationResponse) CertificateInformationResponse {
		if v != nil {
			return *v
		}
		var ret CertificateInformationResponse
		return ret
	}).(CertificateInformationResponseOutput)
}

func (o CertificateInformationResponsePtrOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Expiry
	}).(pulumi.StringPtrOutput)
}

func (o CertificateInformationResponsePtrOutput) Subject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Subject
	}).(pulumi.StringPtrOutput)
}

func (o CertificateInformationResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type DataMasking struct {
	Headers     []DataMaskingEntity `pulumi:"headers"`
	QueryParams []DataMaskingEntity `pulumi:"queryParams"`
}





type DataMaskingInput interface {
	pulumi.Input

	ToDataMaskingOutput() DataMaskingOutput
	ToDataMaskingOutputWithContext(context.Context) DataMaskingOutput
}

type DataMaskingArgs struct {
	Headers     DataMaskingEntityArrayInput `pulumi:"headers"`
	QueryParams DataMaskingEntityArrayInput `pulumi:"queryParams"`
}

func (DataMaskingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMasking)(nil)).Elem()
}

func (i DataMaskingArgs) ToDataMaskingOutput() DataMaskingOutput {
	return i.ToDataMaskingOutputWithContext(context.Background())
}

func (i DataMaskingArgs) ToDataMaskingOutputWithContext(ctx context.Context) DataMaskingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataMaskingOutput)
}

func (i DataMaskingArgs) ToDataMaskingPtrOutput() DataMaskingPtrOutput {
	return i.ToDataMaskingPtrOutputWithContext(context.Background())
}

func (i DataMaskingArgs) ToDataMaskingPtrOutputWithContext(ctx context.Context) DataMaskingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataMaskingOutput).ToDataMaskingPtrOutputWithContext(ctx)
}









type DataMaskingPtrInput interface {
	pulumi.Input

	ToDataMaskingPtrOutput() DataMaskingPtrOutput
	ToDataMaskingPtrOutputWithContext(context.Context) DataMaskingPtrOutput
}

type dataMaskingPtrType DataMaskingArgs

func DataMaskingPtr(v *DataMaskingArgs) DataMaskingPtrInput {
	return (*dataMaskingPtrType)(v)
}

func (*dataMaskingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataMasking)(nil)).Elem()
}

func (i *dataMaskingPtrType) ToDataMaskingPtrOutput() DataMaskingPtrOutput {
	return i.ToDataMaskingPtrOutputWithContext(context.Background())
}

func (i *dataMaskingPtrType) ToDataMaskingPtrOutputWithContext(ctx context.Context) DataMaskingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataMaskingPtrOutput)
}

type DataMaskingOutput struct{ *pulumi.OutputState }

func (DataMaskingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMasking)(nil)).Elem()
}

func (o DataMaskingOutput) ToDataMaskingOutput() DataMaskingOutput {
	return o
}

func (o DataMaskingOutput) ToDataMaskingOutputWithContext(ctx context.Context) DataMaskingOutput {
	return o
}

func (o DataMaskingOutput) ToDataMaskingPtrOutput() DataMaskingPtrOutput {
	return o.ToDataMaskingPtrOutputWithContext(context.Background())
}

func (o DataMaskingOutput) ToDataMaskingPtrOutputWithContext(ctx context.Context) DataMaskingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataMasking) *DataMasking {
		return &v
	}).(DataMaskingPtrOutput)
}

func (o DataMaskingOutput) Headers() DataMaskingEntityArrayOutput {
	return o.ApplyT(func(v DataMasking) []DataMaskingEntity { return v.Headers }).(DataMaskingEntityArrayOutput)
}

func (o DataMaskingOutput) QueryParams() DataMaskingEntityArrayOutput {
	return o.ApplyT(func(v DataMasking) []DataMaskingEntity { return v.QueryParams }).(DataMaskingEntityArrayOutput)
}

type DataMaskingPtrOutput struct{ *pulumi.OutputState }

func (DataMaskingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataMasking)(nil)).Elem()
}

func (o DataMaskingPtrOutput) ToDataMaskingPtrOutput() DataMaskingPtrOutput {
	return o
}

func (o DataMaskingPtrOutput) ToDataMaskingPtrOutputWithContext(ctx context.Context) DataMaskingPtrOutput {
	return o
}

func (o DataMaskingPtrOutput) Elem() DataMaskingOutput {
	return o.ApplyT(func(v *DataMasking) DataMasking {
		if v != nil {
			return *v
		}
		var ret DataMasking
		return ret
	}).(DataMaskingOutput)
}

func (o DataMaskingPtrOutput) Headers() DataMaskingEntityArrayOutput {
	return o.ApplyT(func(v *DataMasking) []DataMaskingEntity {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(DataMaskingEntityArrayOutput)
}

func (o DataMaskingPtrOutput) QueryParams() DataMaskingEntityArrayOutput {
	return o.ApplyT(func(v *DataMasking) []DataMaskingEntity {
		if v == nil {
			return nil
		}
		return v.QueryParams
	}).(DataMaskingEntityArrayOutput)
}

type DataMaskingEntity struct {
	Mode  *string `pulumi:"mode"`
	Value *string `pulumi:"value"`
}





type DataMaskingEntityInput interface {
	pulumi.Input

	ToDataMaskingEntityOutput() DataMaskingEntityOutput
	ToDataMaskingEntityOutputWithContext(context.Context) DataMaskingEntityOutput
}

type DataMaskingEntityArgs struct {
	Mode  pulumi.StringPtrInput `pulumi:"mode"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (DataMaskingEntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingEntity)(nil)).Elem()
}

func (i DataMaskingEntityArgs) ToDataMaskingEntityOutput() DataMaskingEntityOutput {
	return i.ToDataMaskingEntityOutputWithContext(context.Background())
}

func (i DataMaskingEntityArgs) ToDataMaskingEntityOutputWithContext(ctx context.Context) DataMaskingEntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataMaskingEntityOutput)
}





type DataMaskingEntityArrayInput interface {
	pulumi.Input

	ToDataMaskingEntityArrayOutput() DataMaskingEntityArrayOutput
	ToDataMaskingEntityArrayOutputWithContext(context.Context) DataMaskingEntityArrayOutput
}

type DataMaskingEntityArray []DataMaskingEntityInput

func (DataMaskingEntityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataMaskingEntity)(nil)).Elem()
}

func (i DataMaskingEntityArray) ToDataMaskingEntityArrayOutput() DataMaskingEntityArrayOutput {
	return i.ToDataMaskingEntityArrayOutputWithContext(context.Background())
}

func (i DataMaskingEntityArray) ToDataMaskingEntityArrayOutputWithContext(ctx context.Context) DataMaskingEntityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataMaskingEntityArrayOutput)
}

type DataMaskingEntityOutput struct{ *pulumi.OutputState }

func (DataMaskingEntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingEntity)(nil)).Elem()
}

func (o DataMaskingEntityOutput) ToDataMaskingEntityOutput() DataMaskingEntityOutput {
	return o
}

func (o DataMaskingEntityOutput) ToDataMaskingEntityOutputWithContext(ctx context.Context) DataMaskingEntityOutput {
	return o
}

func (o DataMaskingEntityOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataMaskingEntity) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o DataMaskingEntityOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataMaskingEntity) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DataMaskingEntityArrayOutput struct{ *pulumi.OutputState }

func (DataMaskingEntityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataMaskingEntity)(nil)).Elem()
}

func (o DataMaskingEntityArrayOutput) ToDataMaskingEntityArrayOutput() DataMaskingEntityArrayOutput {
	return o
}

func (o DataMaskingEntityArrayOutput) ToDataMaskingEntityArrayOutputWithContext(ctx context.Context) DataMaskingEntityArrayOutput {
	return o
}

func (o DataMaskingEntityArrayOutput) Index(i pulumi.IntInput) DataMaskingEntityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataMaskingEntity {
		return vs[0].([]DataMaskingEntity)[vs[1].(int)]
	}).(DataMaskingEntityOutput)
}

type DataMaskingEntityResponse struct {
	Mode  *string `pulumi:"mode"`
	Value *string `pulumi:"value"`
}

type DataMaskingEntityResponseOutput struct{ *pulumi.OutputState }

func (DataMaskingEntityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingEntityResponse)(nil)).Elem()
}

func (o DataMaskingEntityResponseOutput) ToDataMaskingEntityResponseOutput() DataMaskingEntityResponseOutput {
	return o
}

func (o DataMaskingEntityResponseOutput) ToDataMaskingEntityResponseOutputWithContext(ctx context.Context) DataMaskingEntityResponseOutput {
	return o
}

func (o DataMaskingEntityResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataMaskingEntityResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o DataMaskingEntityResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataMaskingEntityResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DataMaskingEntityResponseArrayOutput struct{ *pulumi.OutputState }

func (DataMaskingEntityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataMaskingEntityResponse)(nil)).Elem()
}

func (o DataMaskingEntityResponseArrayOutput) ToDataMaskingEntityResponseArrayOutput() DataMaskingEntityResponseArrayOutput {
	return o
}

func (o DataMaskingEntityResponseArrayOutput) ToDataMaskingEntityResponseArrayOutputWithContext(ctx context.Context) DataMaskingEntityResponseArrayOutput {
	return o
}

func (o DataMaskingEntityResponseArrayOutput) Index(i pulumi.IntInput) DataMaskingEntityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataMaskingEntityResponse {
		return vs[0].([]DataMaskingEntityResponse)[vs[1].(int)]
	}).(DataMaskingEntityResponseOutput)
}

type DataMaskingResponse struct {
	Headers     []DataMaskingEntityResponse `pulumi:"headers"`
	QueryParams []DataMaskingEntityResponse `pulumi:"queryParams"`
}

type DataMaskingResponseOutput struct{ *pulumi.OutputState }

func (DataMaskingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataMaskingResponse)(nil)).Elem()
}

func (o DataMaskingResponseOutput) ToDataMaskingResponseOutput() DataMaskingResponseOutput {
	return o
}

func (o DataMaskingResponseOutput) ToDataMaskingResponseOutputWithContext(ctx context.Context) DataMaskingResponseOutput {
	return o
}

func (o DataMaskingResponseOutput) Headers() DataMaskingEntityResponseArrayOutput {
	return o.ApplyT(func(v DataMaskingResponse) []DataMaskingEntityResponse { return v.Headers }).(DataMaskingEntityResponseArrayOutput)
}

func (o DataMaskingResponseOutput) QueryParams() DataMaskingEntityResponseArrayOutput {
	return o.ApplyT(func(v DataMaskingResponse) []DataMaskingEntityResponse { return v.QueryParams }).(DataMaskingEntityResponseArrayOutput)
}

type DataMaskingResponsePtrOutput struct{ *pulumi.OutputState }

func (DataMaskingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataMaskingResponse)(nil)).Elem()
}

func (o DataMaskingResponsePtrOutput) ToDataMaskingResponsePtrOutput() DataMaskingResponsePtrOutput {
	return o
}

func (o DataMaskingResponsePtrOutput) ToDataMaskingResponsePtrOutputWithContext(ctx context.Context) DataMaskingResponsePtrOutput {
	return o
}

func (o DataMaskingResponsePtrOutput) Elem() DataMaskingResponseOutput {
	return o.ApplyT(func(v *DataMaskingResponse) DataMaskingResponse {
		if v != nil {
			return *v
		}
		var ret DataMaskingResponse
		return ret
	}).(DataMaskingResponseOutput)
}

func (o DataMaskingResponsePtrOutput) Headers() DataMaskingEntityResponseArrayOutput {
	return o.ApplyT(func(v *DataMaskingResponse) []DataMaskingEntityResponse {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(DataMaskingEntityResponseArrayOutput)
}

func (o DataMaskingResponsePtrOutput) QueryParams() DataMaskingEntityResponseArrayOutput {
	return o.ApplyT(func(v *DataMaskingResponse) []DataMaskingEntityResponse {
		if v == nil {
			return nil
		}
		return v.QueryParams
	}).(DataMaskingEntityResponseArrayOutput)
}

type EmailTemplateParametersContractProperties struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Title       *string `pulumi:"title"`
}





type EmailTemplateParametersContractPropertiesInput interface {
	pulumi.Input

	ToEmailTemplateParametersContractPropertiesOutput() EmailTemplateParametersContractPropertiesOutput
	ToEmailTemplateParametersContractPropertiesOutputWithContext(context.Context) EmailTemplateParametersContractPropertiesOutput
}

type EmailTemplateParametersContractPropertiesArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	Title       pulumi.StringPtrInput `pulumi:"title"`
}

func (EmailTemplateParametersContractPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailTemplateParametersContractProperties)(nil)).Elem()
}

func (i EmailTemplateParametersContractPropertiesArgs) ToEmailTemplateParametersContractPropertiesOutput() EmailTemplateParametersContractPropertiesOutput {
	return i.ToEmailTemplateParametersContractPropertiesOutputWithContext(context.Background())
}

func (i EmailTemplateParametersContractPropertiesArgs) ToEmailTemplateParametersContractPropertiesOutputWithContext(ctx context.Context) EmailTemplateParametersContractPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailTemplateParametersContractPropertiesOutput)
}





type EmailTemplateParametersContractPropertiesArrayInput interface {
	pulumi.Input

	ToEmailTemplateParametersContractPropertiesArrayOutput() EmailTemplateParametersContractPropertiesArrayOutput
	ToEmailTemplateParametersContractPropertiesArrayOutputWithContext(context.Context) EmailTemplateParametersContractPropertiesArrayOutput
}

type EmailTemplateParametersContractPropertiesArray []EmailTemplateParametersContractPropertiesInput

func (EmailTemplateParametersContractPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailTemplateParametersContractProperties)(nil)).Elem()
}

func (i EmailTemplateParametersContractPropertiesArray) ToEmailTemplateParametersContractPropertiesArrayOutput() EmailTemplateParametersContractPropertiesArrayOutput {
	return i.ToEmailTemplateParametersContractPropertiesArrayOutputWithContext(context.Background())
}

func (i EmailTemplateParametersContractPropertiesArray) ToEmailTemplateParametersContractPropertiesArrayOutputWithContext(ctx context.Context) EmailTemplateParametersContractPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailTemplateParametersContractPropertiesArrayOutput)
}

type EmailTemplateParametersContractPropertiesOutput struct{ *pulumi.OutputState }

func (EmailTemplateParametersContractPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailTemplateParametersContractProperties)(nil)).Elem()
}

func (o EmailTemplateParametersContractPropertiesOutput) ToEmailTemplateParametersContractPropertiesOutput() EmailTemplateParametersContractPropertiesOutput {
	return o
}

func (o EmailTemplateParametersContractPropertiesOutput) ToEmailTemplateParametersContractPropertiesOutputWithContext(ctx context.Context) EmailTemplateParametersContractPropertiesOutput {
	return o
}

func (o EmailTemplateParametersContractPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmailTemplateParametersContractProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EmailTemplateParametersContractPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmailTemplateParametersContractProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EmailTemplateParametersContractPropertiesOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmailTemplateParametersContractProperties) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type EmailTemplateParametersContractPropertiesArrayOutput struct{ *pulumi.OutputState }

func (EmailTemplateParametersContractPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailTemplateParametersContractProperties)(nil)).Elem()
}

func (o EmailTemplateParametersContractPropertiesArrayOutput) ToEmailTemplateParametersContractPropertiesArrayOutput() EmailTemplateParametersContractPropertiesArrayOutput {
	return o
}

func (o EmailTemplateParametersContractPropertiesArrayOutput) ToEmailTemplateParametersContractPropertiesArrayOutputWithContext(ctx context.Context) EmailTemplateParametersContractPropertiesArrayOutput {
	return o
}

func (o EmailTemplateParametersContractPropertiesArrayOutput) Index(i pulumi.IntInput) EmailTemplateParametersContractPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmailTemplateParametersContractProperties {
		return vs[0].([]EmailTemplateParametersContractProperties)[vs[1].(int)]
	}).(EmailTemplateParametersContractPropertiesOutput)
}

type EmailTemplateParametersContractPropertiesResponse struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Title       *string `pulumi:"title"`
}

type EmailTemplateParametersContractPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EmailTemplateParametersContractPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailTemplateParametersContractPropertiesResponse)(nil)).Elem()
}

func (o EmailTemplateParametersContractPropertiesResponseOutput) ToEmailTemplateParametersContractPropertiesResponseOutput() EmailTemplateParametersContractPropertiesResponseOutput {
	return o
}

func (o EmailTemplateParametersContractPropertiesResponseOutput) ToEmailTemplateParametersContractPropertiesResponseOutputWithContext(ctx context.Context) EmailTemplateParametersContractPropertiesResponseOutput {
	return o
}

func (o EmailTemplateParametersContractPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmailTemplateParametersContractPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EmailTemplateParametersContractPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmailTemplateParametersContractPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EmailTemplateParametersContractPropertiesResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmailTemplateParametersContractPropertiesResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type EmailTemplateParametersContractPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (EmailTemplateParametersContractPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailTemplateParametersContractPropertiesResponse)(nil)).Elem()
}

func (o EmailTemplateParametersContractPropertiesResponseArrayOutput) ToEmailTemplateParametersContractPropertiesResponseArrayOutput() EmailTemplateParametersContractPropertiesResponseArrayOutput {
	return o
}

func (o EmailTemplateParametersContractPropertiesResponseArrayOutput) ToEmailTemplateParametersContractPropertiesResponseArrayOutputWithContext(ctx context.Context) EmailTemplateParametersContractPropertiesResponseArrayOutput {
	return o
}

func (o EmailTemplateParametersContractPropertiesResponseArrayOutput) Index(i pulumi.IntInput) EmailTemplateParametersContractPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmailTemplateParametersContractPropertiesResponse {
		return vs[0].([]EmailTemplateParametersContractPropertiesResponse)[vs[1].(int)]
	}).(EmailTemplateParametersContractPropertiesResponseOutput)
}

type GroupContractPropertiesResponse struct {
	BuiltIn     bool    `pulumi:"builtIn"`
	Description *string `pulumi:"description"`
	DisplayName string  `pulumi:"displayName"`
	ExternalId  *string `pulumi:"externalId"`
	Type        *string `pulumi:"type"`
}

type GroupContractPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GroupContractPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupContractPropertiesResponse)(nil)).Elem()
}

func (o GroupContractPropertiesResponseOutput) ToGroupContractPropertiesResponseOutput() GroupContractPropertiesResponseOutput {
	return o
}

func (o GroupContractPropertiesResponseOutput) ToGroupContractPropertiesResponseOutputWithContext(ctx context.Context) GroupContractPropertiesResponseOutput {
	return o
}

func (o GroupContractPropertiesResponseOutput) BuiltIn() pulumi.BoolOutput {
	return o.ApplyT(func(v GroupContractPropertiesResponse) bool { return v.BuiltIn }).(pulumi.BoolOutput)
}

func (o GroupContractPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupContractPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GroupContractPropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GroupContractPropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o GroupContractPropertiesResponseOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupContractPropertiesResponse) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

func (o GroupContractPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupContractPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type GroupContractPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (GroupContractPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupContractPropertiesResponse)(nil)).Elem()
}

func (o GroupContractPropertiesResponseArrayOutput) ToGroupContractPropertiesResponseArrayOutput() GroupContractPropertiesResponseArrayOutput {
	return o
}

func (o GroupContractPropertiesResponseArrayOutput) ToGroupContractPropertiesResponseArrayOutputWithContext(ctx context.Context) GroupContractPropertiesResponseArrayOutput {
	return o
}

func (o GroupContractPropertiesResponseArrayOutput) Index(i pulumi.IntInput) GroupContractPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupContractPropertiesResponse {
		return vs[0].([]GroupContractPropertiesResponse)[vs[1].(int)]
	}).(GroupContractPropertiesResponseOutput)
}

type HostnameConfiguration struct {
	Certificate                *CertificateInformation `pulumi:"certificate"`
	CertificatePassword        *string                 `pulumi:"certificatePassword"`
	CertificateSource          *string                 `pulumi:"certificateSource"`
	CertificateStatus          *string                 `pulumi:"certificateStatus"`
	DefaultSslBinding          *bool                   `pulumi:"defaultSslBinding"`
	EncodedCertificate         *string                 `pulumi:"encodedCertificate"`
	HostName                   string                  `pulumi:"hostName"`
	IdentityClientId           *string                 `pulumi:"identityClientId"`
	KeyVaultId                 *string                 `pulumi:"keyVaultId"`
	NegotiateClientCertificate *bool                   `pulumi:"negotiateClientCertificate"`
	Type                       string                  `pulumi:"type"`
}


func (val *HostnameConfiguration) Defaults() *HostnameConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultSslBinding) {
		defaultSslBinding_ := false
		tmp.DefaultSslBinding = &defaultSslBinding_
	}
	if isZero(tmp.NegotiateClientCertificate) {
		negotiateClientCertificate_ := false
		tmp.NegotiateClientCertificate = &negotiateClientCertificate_
	}
	return &tmp
}





type HostnameConfigurationInput interface {
	pulumi.Input

	ToHostnameConfigurationOutput() HostnameConfigurationOutput
	ToHostnameConfigurationOutputWithContext(context.Context) HostnameConfigurationOutput
}

type HostnameConfigurationArgs struct {
	Certificate                CertificateInformationPtrInput `pulumi:"certificate"`
	CertificatePassword        pulumi.StringPtrInput          `pulumi:"certificatePassword"`
	CertificateSource          pulumi.StringPtrInput          `pulumi:"certificateSource"`
	CertificateStatus          pulumi.StringPtrInput          `pulumi:"certificateStatus"`
	DefaultSslBinding          pulumi.BoolPtrInput            `pulumi:"defaultSslBinding"`
	EncodedCertificate         pulumi.StringPtrInput          `pulumi:"encodedCertificate"`
	HostName                   pulumi.StringInput             `pulumi:"hostName"`
	IdentityClientId           pulumi.StringPtrInput          `pulumi:"identityClientId"`
	KeyVaultId                 pulumi.StringPtrInput          `pulumi:"keyVaultId"`
	NegotiateClientCertificate pulumi.BoolPtrInput            `pulumi:"negotiateClientCertificate"`
	Type                       pulumi.StringInput             `pulumi:"type"`
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

func (o HostnameConfigurationOutput) Certificate() CertificateInformationPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *CertificateInformation { return v.Certificate }).(CertificateInformationPtrOutput)
}

func (o HostnameConfigurationOutput) CertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *string { return v.CertificatePassword }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationOutput) CertificateSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *string { return v.CertificateSource }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationOutput) CertificateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *string { return v.CertificateStatus }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationOutput) DefaultSslBinding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *bool { return v.DefaultSslBinding }).(pulumi.BoolPtrOutput)
}

func (o HostnameConfigurationOutput) EncodedCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *string { return v.EncodedCertificate }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v HostnameConfiguration) string { return v.HostName }).(pulumi.StringOutput)
}

func (o HostnameConfigurationOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationOutput) NegotiateClientCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *bool { return v.NegotiateClientCertificate }).(pulumi.BoolPtrOutput)
}

func (o HostnameConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HostnameConfiguration) string { return v.Type }).(pulumi.StringOutput)
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
	Certificate                *CertificateInformationResponse `pulumi:"certificate"`
	CertificatePassword        *string                         `pulumi:"certificatePassword"`
	CertificateSource          *string                         `pulumi:"certificateSource"`
	CertificateStatus          *string                         `pulumi:"certificateStatus"`
	DefaultSslBinding          *bool                           `pulumi:"defaultSslBinding"`
	EncodedCertificate         *string                         `pulumi:"encodedCertificate"`
	HostName                   string                          `pulumi:"hostName"`
	IdentityClientId           *string                         `pulumi:"identityClientId"`
	KeyVaultId                 *string                         `pulumi:"keyVaultId"`
	NegotiateClientCertificate *bool                           `pulumi:"negotiateClientCertificate"`
	Type                       string                          `pulumi:"type"`
}


func (val *HostnameConfigurationResponse) Defaults() *HostnameConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultSslBinding) {
		defaultSslBinding_ := false
		tmp.DefaultSslBinding = &defaultSslBinding_
	}
	if isZero(tmp.NegotiateClientCertificate) {
		negotiateClientCertificate_ := false
		tmp.NegotiateClientCertificate = &negotiateClientCertificate_
	}
	return &tmp
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

func (o HostnameConfigurationResponseOutput) Certificate() CertificateInformationResponsePtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *CertificateInformationResponse { return v.Certificate }).(CertificateInformationResponsePtrOutput)
}

func (o HostnameConfigurationResponseOutput) CertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *string { return v.CertificatePassword }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationResponseOutput) CertificateSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *string { return v.CertificateSource }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationResponseOutput) CertificateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *string { return v.CertificateStatus }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationResponseOutput) DefaultSslBinding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *bool { return v.DefaultSslBinding }).(pulumi.BoolPtrOutput)
}

func (o HostnameConfigurationResponseOutput) EncodedCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *string { return v.EncodedCertificate }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationResponseOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) string { return v.HostName }).(pulumi.StringOutput)
}

func (o HostnameConfigurationResponseOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationResponseOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationResponseOutput) NegotiateClientCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *bool { return v.NegotiateClientCertificate }).(pulumi.BoolPtrOutput)
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

type HttpMessageDiagnostic struct {
	Body        *BodyDiagnosticSettings `pulumi:"body"`
	DataMasking *DataMasking            `pulumi:"dataMasking"`
	Headers     []string                `pulumi:"headers"`
}





type HttpMessageDiagnosticInput interface {
	pulumi.Input

	ToHttpMessageDiagnosticOutput() HttpMessageDiagnosticOutput
	ToHttpMessageDiagnosticOutputWithContext(context.Context) HttpMessageDiagnosticOutput
}

type HttpMessageDiagnosticArgs struct {
	Body        BodyDiagnosticSettingsPtrInput `pulumi:"body"`
	DataMasking DataMaskingPtrInput            `pulumi:"dataMasking"`
	Headers     pulumi.StringArrayInput        `pulumi:"headers"`
}

func (HttpMessageDiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpMessageDiagnostic)(nil)).Elem()
}

func (i HttpMessageDiagnosticArgs) ToHttpMessageDiagnosticOutput() HttpMessageDiagnosticOutput {
	return i.ToHttpMessageDiagnosticOutputWithContext(context.Background())
}

func (i HttpMessageDiagnosticArgs) ToHttpMessageDiagnosticOutputWithContext(ctx context.Context) HttpMessageDiagnosticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpMessageDiagnosticOutput)
}

func (i HttpMessageDiagnosticArgs) ToHttpMessageDiagnosticPtrOutput() HttpMessageDiagnosticPtrOutput {
	return i.ToHttpMessageDiagnosticPtrOutputWithContext(context.Background())
}

func (i HttpMessageDiagnosticArgs) ToHttpMessageDiagnosticPtrOutputWithContext(ctx context.Context) HttpMessageDiagnosticPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpMessageDiagnosticOutput).ToHttpMessageDiagnosticPtrOutputWithContext(ctx)
}









type HttpMessageDiagnosticPtrInput interface {
	pulumi.Input

	ToHttpMessageDiagnosticPtrOutput() HttpMessageDiagnosticPtrOutput
	ToHttpMessageDiagnosticPtrOutputWithContext(context.Context) HttpMessageDiagnosticPtrOutput
}

type httpMessageDiagnosticPtrType HttpMessageDiagnosticArgs

func HttpMessageDiagnosticPtr(v *HttpMessageDiagnosticArgs) HttpMessageDiagnosticPtrInput {
	return (*httpMessageDiagnosticPtrType)(v)
}

func (*httpMessageDiagnosticPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpMessageDiagnostic)(nil)).Elem()
}

func (i *httpMessageDiagnosticPtrType) ToHttpMessageDiagnosticPtrOutput() HttpMessageDiagnosticPtrOutput {
	return i.ToHttpMessageDiagnosticPtrOutputWithContext(context.Background())
}

func (i *httpMessageDiagnosticPtrType) ToHttpMessageDiagnosticPtrOutputWithContext(ctx context.Context) HttpMessageDiagnosticPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpMessageDiagnosticPtrOutput)
}

type HttpMessageDiagnosticOutput struct{ *pulumi.OutputState }

func (HttpMessageDiagnosticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpMessageDiagnostic)(nil)).Elem()
}

func (o HttpMessageDiagnosticOutput) ToHttpMessageDiagnosticOutput() HttpMessageDiagnosticOutput {
	return o
}

func (o HttpMessageDiagnosticOutput) ToHttpMessageDiagnosticOutputWithContext(ctx context.Context) HttpMessageDiagnosticOutput {
	return o
}

func (o HttpMessageDiagnosticOutput) ToHttpMessageDiagnosticPtrOutput() HttpMessageDiagnosticPtrOutput {
	return o.ToHttpMessageDiagnosticPtrOutputWithContext(context.Background())
}

func (o HttpMessageDiagnosticOutput) ToHttpMessageDiagnosticPtrOutputWithContext(ctx context.Context) HttpMessageDiagnosticPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpMessageDiagnostic) *HttpMessageDiagnostic {
		return &v
	}).(HttpMessageDiagnosticPtrOutput)
}

func (o HttpMessageDiagnosticOutput) Body() BodyDiagnosticSettingsPtrOutput {
	return o.ApplyT(func(v HttpMessageDiagnostic) *BodyDiagnosticSettings { return v.Body }).(BodyDiagnosticSettingsPtrOutput)
}

func (o HttpMessageDiagnosticOutput) DataMasking() DataMaskingPtrOutput {
	return o.ApplyT(func(v HttpMessageDiagnostic) *DataMasking { return v.DataMasking }).(DataMaskingPtrOutput)
}

func (o HttpMessageDiagnosticOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HttpMessageDiagnostic) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

type HttpMessageDiagnosticPtrOutput struct{ *pulumi.OutputState }

func (HttpMessageDiagnosticPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpMessageDiagnostic)(nil)).Elem()
}

func (o HttpMessageDiagnosticPtrOutput) ToHttpMessageDiagnosticPtrOutput() HttpMessageDiagnosticPtrOutput {
	return o
}

func (o HttpMessageDiagnosticPtrOutput) ToHttpMessageDiagnosticPtrOutputWithContext(ctx context.Context) HttpMessageDiagnosticPtrOutput {
	return o
}

func (o HttpMessageDiagnosticPtrOutput) Elem() HttpMessageDiagnosticOutput {
	return o.ApplyT(func(v *HttpMessageDiagnostic) HttpMessageDiagnostic {
		if v != nil {
			return *v
		}
		var ret HttpMessageDiagnostic
		return ret
	}).(HttpMessageDiagnosticOutput)
}

func (o HttpMessageDiagnosticPtrOutput) Body() BodyDiagnosticSettingsPtrOutput {
	return o.ApplyT(func(v *HttpMessageDiagnostic) *BodyDiagnosticSettings {
		if v == nil {
			return nil
		}
		return v.Body
	}).(BodyDiagnosticSettingsPtrOutput)
}

func (o HttpMessageDiagnosticPtrOutput) DataMasking() DataMaskingPtrOutput {
	return o.ApplyT(func(v *HttpMessageDiagnostic) *DataMasking {
		if v == nil {
			return nil
		}
		return v.DataMasking
	}).(DataMaskingPtrOutput)
}

func (o HttpMessageDiagnosticPtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HttpMessageDiagnostic) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

type HttpMessageDiagnosticResponse struct {
	Body        *BodyDiagnosticSettingsResponse `pulumi:"body"`
	DataMasking *DataMaskingResponse            `pulumi:"dataMasking"`
	Headers     []string                        `pulumi:"headers"`
}

type HttpMessageDiagnosticResponseOutput struct{ *pulumi.OutputState }

func (HttpMessageDiagnosticResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpMessageDiagnosticResponse)(nil)).Elem()
}

func (o HttpMessageDiagnosticResponseOutput) ToHttpMessageDiagnosticResponseOutput() HttpMessageDiagnosticResponseOutput {
	return o
}

func (o HttpMessageDiagnosticResponseOutput) ToHttpMessageDiagnosticResponseOutputWithContext(ctx context.Context) HttpMessageDiagnosticResponseOutput {
	return o
}

func (o HttpMessageDiagnosticResponseOutput) Body() BodyDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v HttpMessageDiagnosticResponse) *BodyDiagnosticSettingsResponse { return v.Body }).(BodyDiagnosticSettingsResponsePtrOutput)
}

func (o HttpMessageDiagnosticResponseOutput) DataMasking() DataMaskingResponsePtrOutput {
	return o.ApplyT(func(v HttpMessageDiagnosticResponse) *DataMaskingResponse { return v.DataMasking }).(DataMaskingResponsePtrOutput)
}

func (o HttpMessageDiagnosticResponseOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HttpMessageDiagnosticResponse) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

type HttpMessageDiagnosticResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpMessageDiagnosticResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpMessageDiagnosticResponse)(nil)).Elem()
}

func (o HttpMessageDiagnosticResponsePtrOutput) ToHttpMessageDiagnosticResponsePtrOutput() HttpMessageDiagnosticResponsePtrOutput {
	return o
}

func (o HttpMessageDiagnosticResponsePtrOutput) ToHttpMessageDiagnosticResponsePtrOutputWithContext(ctx context.Context) HttpMessageDiagnosticResponsePtrOutput {
	return o
}

func (o HttpMessageDiagnosticResponsePtrOutput) Elem() HttpMessageDiagnosticResponseOutput {
	return o.ApplyT(func(v *HttpMessageDiagnosticResponse) HttpMessageDiagnosticResponse {
		if v != nil {
			return *v
		}
		var ret HttpMessageDiagnosticResponse
		return ret
	}).(HttpMessageDiagnosticResponseOutput)
}

func (o HttpMessageDiagnosticResponsePtrOutput) Body() BodyDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v *HttpMessageDiagnosticResponse) *BodyDiagnosticSettingsResponse {
		if v == nil {
			return nil
		}
		return v.Body
	}).(BodyDiagnosticSettingsResponsePtrOutput)
}

func (o HttpMessageDiagnosticResponsePtrOutput) DataMasking() DataMaskingResponsePtrOutput {
	return o.ApplyT(func(v *HttpMessageDiagnosticResponse) *DataMaskingResponse {
		if v == nil {
			return nil
		}
		return v.DataMasking
	}).(DataMaskingResponsePtrOutput)
}

func (o HttpMessageDiagnosticResponsePtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HttpMessageDiagnosticResponse) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

type KeyVaultContractCreateProperties struct {
	IdentityClientId *string `pulumi:"identityClientId"`
	SecretIdentifier *string `pulumi:"secretIdentifier"`
}





type KeyVaultContractCreatePropertiesInput interface {
	pulumi.Input

	ToKeyVaultContractCreatePropertiesOutput() KeyVaultContractCreatePropertiesOutput
	ToKeyVaultContractCreatePropertiesOutputWithContext(context.Context) KeyVaultContractCreatePropertiesOutput
}

type KeyVaultContractCreatePropertiesArgs struct {
	IdentityClientId pulumi.StringPtrInput `pulumi:"identityClientId"`
	SecretIdentifier pulumi.StringPtrInput `pulumi:"secretIdentifier"`
}

func (KeyVaultContractCreatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultContractCreateProperties)(nil)).Elem()
}

func (i KeyVaultContractCreatePropertiesArgs) ToKeyVaultContractCreatePropertiesOutput() KeyVaultContractCreatePropertiesOutput {
	return i.ToKeyVaultContractCreatePropertiesOutputWithContext(context.Background())
}

func (i KeyVaultContractCreatePropertiesArgs) ToKeyVaultContractCreatePropertiesOutputWithContext(ctx context.Context) KeyVaultContractCreatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultContractCreatePropertiesOutput)
}

func (i KeyVaultContractCreatePropertiesArgs) ToKeyVaultContractCreatePropertiesPtrOutput() KeyVaultContractCreatePropertiesPtrOutput {
	return i.ToKeyVaultContractCreatePropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultContractCreatePropertiesArgs) ToKeyVaultContractCreatePropertiesPtrOutputWithContext(ctx context.Context) KeyVaultContractCreatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultContractCreatePropertiesOutput).ToKeyVaultContractCreatePropertiesPtrOutputWithContext(ctx)
}









type KeyVaultContractCreatePropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultContractCreatePropertiesPtrOutput() KeyVaultContractCreatePropertiesPtrOutput
	ToKeyVaultContractCreatePropertiesPtrOutputWithContext(context.Context) KeyVaultContractCreatePropertiesPtrOutput
}

type keyVaultContractCreatePropertiesPtrType KeyVaultContractCreatePropertiesArgs

func KeyVaultContractCreatePropertiesPtr(v *KeyVaultContractCreatePropertiesArgs) KeyVaultContractCreatePropertiesPtrInput {
	return (*keyVaultContractCreatePropertiesPtrType)(v)
}

func (*keyVaultContractCreatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultContractCreateProperties)(nil)).Elem()
}

func (i *keyVaultContractCreatePropertiesPtrType) ToKeyVaultContractCreatePropertiesPtrOutput() KeyVaultContractCreatePropertiesPtrOutput {
	return i.ToKeyVaultContractCreatePropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultContractCreatePropertiesPtrType) ToKeyVaultContractCreatePropertiesPtrOutputWithContext(ctx context.Context) KeyVaultContractCreatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultContractCreatePropertiesPtrOutput)
}

type KeyVaultContractCreatePropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultContractCreatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultContractCreateProperties)(nil)).Elem()
}

func (o KeyVaultContractCreatePropertiesOutput) ToKeyVaultContractCreatePropertiesOutput() KeyVaultContractCreatePropertiesOutput {
	return o
}

func (o KeyVaultContractCreatePropertiesOutput) ToKeyVaultContractCreatePropertiesOutputWithContext(ctx context.Context) KeyVaultContractCreatePropertiesOutput {
	return o
}

func (o KeyVaultContractCreatePropertiesOutput) ToKeyVaultContractCreatePropertiesPtrOutput() KeyVaultContractCreatePropertiesPtrOutput {
	return o.ToKeyVaultContractCreatePropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultContractCreatePropertiesOutput) ToKeyVaultContractCreatePropertiesPtrOutputWithContext(ctx context.Context) KeyVaultContractCreatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultContractCreateProperties) *KeyVaultContractCreateProperties {
		return &v
	}).(KeyVaultContractCreatePropertiesPtrOutput)
}

func (o KeyVaultContractCreatePropertiesOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultContractCreateProperties) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o KeyVaultContractCreatePropertiesOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultContractCreateProperties) *string { return v.SecretIdentifier }).(pulumi.StringPtrOutput)
}

type KeyVaultContractCreatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultContractCreatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultContractCreateProperties)(nil)).Elem()
}

func (o KeyVaultContractCreatePropertiesPtrOutput) ToKeyVaultContractCreatePropertiesPtrOutput() KeyVaultContractCreatePropertiesPtrOutput {
	return o
}

func (o KeyVaultContractCreatePropertiesPtrOutput) ToKeyVaultContractCreatePropertiesPtrOutputWithContext(ctx context.Context) KeyVaultContractCreatePropertiesPtrOutput {
	return o
}

func (o KeyVaultContractCreatePropertiesPtrOutput) Elem() KeyVaultContractCreatePropertiesOutput {
	return o.ApplyT(func(v *KeyVaultContractCreateProperties) KeyVaultContractCreateProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultContractCreateProperties
		return ret
	}).(KeyVaultContractCreatePropertiesOutput)
}

func (o KeyVaultContractCreatePropertiesPtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultContractCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultContractCreatePropertiesPtrOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultContractCreateProperties) *string {
		if v == nil {
			return nil
		}
		return v.SecretIdentifier
	}).(pulumi.StringPtrOutput)
}

type KeyVaultContractPropertiesResponse struct {
	IdentityClientId *string                                             `pulumi:"identityClientId"`
	LastStatus       *KeyVaultLastAccessStatusContractPropertiesResponse `pulumi:"lastStatus"`
	SecretIdentifier *string                                             `pulumi:"secretIdentifier"`
}

type KeyVaultContractPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultContractPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultContractPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultContractPropertiesResponseOutput) ToKeyVaultContractPropertiesResponseOutput() KeyVaultContractPropertiesResponseOutput {
	return o
}

func (o KeyVaultContractPropertiesResponseOutput) ToKeyVaultContractPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultContractPropertiesResponseOutput {
	return o
}

func (o KeyVaultContractPropertiesResponseOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultContractPropertiesResponse) *string { return v.IdentityClientId }).(pulumi.StringPtrOutput)
}

func (o KeyVaultContractPropertiesResponseOutput) LastStatus() KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput {
	return o.ApplyT(func(v KeyVaultContractPropertiesResponse) *KeyVaultLastAccessStatusContractPropertiesResponse {
		return v.LastStatus
	}).(KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput)
}

func (o KeyVaultContractPropertiesResponseOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultContractPropertiesResponse) *string { return v.SecretIdentifier }).(pulumi.StringPtrOutput)
}

type KeyVaultContractPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultContractPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultContractPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultContractPropertiesResponsePtrOutput) ToKeyVaultContractPropertiesResponsePtrOutput() KeyVaultContractPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultContractPropertiesResponsePtrOutput) ToKeyVaultContractPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultContractPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultContractPropertiesResponsePtrOutput) Elem() KeyVaultContractPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultContractPropertiesResponse) KeyVaultContractPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultContractPropertiesResponse
		return ret
	}).(KeyVaultContractPropertiesResponseOutput)
}

func (o KeyVaultContractPropertiesResponsePtrOutput) IdentityClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultContractPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdentityClientId
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultContractPropertiesResponsePtrOutput) LastStatus() KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultContractPropertiesResponse) *KeyVaultLastAccessStatusContractPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.LastStatus
	}).(KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput)
}

func (o KeyVaultContractPropertiesResponsePtrOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultContractPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecretIdentifier
	}).(pulumi.StringPtrOutput)
}

type KeyVaultLastAccessStatusContractPropertiesResponse struct {
	Code         *string `pulumi:"code"`
	Message      *string `pulumi:"message"`
	TimeStampUtc *string `pulumi:"timeStampUtc"`
}

type KeyVaultLastAccessStatusContractPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultLastAccessStatusContractPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultLastAccessStatusContractPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultLastAccessStatusContractPropertiesResponseOutput) ToKeyVaultLastAccessStatusContractPropertiesResponseOutput() KeyVaultLastAccessStatusContractPropertiesResponseOutput {
	return o
}

func (o KeyVaultLastAccessStatusContractPropertiesResponseOutput) ToKeyVaultLastAccessStatusContractPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultLastAccessStatusContractPropertiesResponseOutput {
	return o
}

func (o KeyVaultLastAccessStatusContractPropertiesResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultLastAccessStatusContractPropertiesResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o KeyVaultLastAccessStatusContractPropertiesResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultLastAccessStatusContractPropertiesResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o KeyVaultLastAccessStatusContractPropertiesResponseOutput) TimeStampUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultLastAccessStatusContractPropertiesResponse) *string { return v.TimeStampUtc }).(pulumi.StringPtrOutput)
}

type KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultLastAccessStatusContractPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput) ToKeyVaultLastAccessStatusContractPropertiesResponsePtrOutput() KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput) ToKeyVaultLastAccessStatusContractPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput) Elem() KeyVaultLastAccessStatusContractPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultLastAccessStatusContractPropertiesResponse) KeyVaultLastAccessStatusContractPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultLastAccessStatusContractPropertiesResponse
		return ret
	}).(KeyVaultLastAccessStatusContractPropertiesResponseOutput)
}

func (o KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultLastAccessStatusContractPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultLastAccessStatusContractPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput) TimeStampUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultLastAccessStatusContractPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeStampUtc
	}).(pulumi.StringPtrOutput)
}

type OAuth2AuthenticationSettingsContract struct {
	AuthorizationServerId *string `pulumi:"authorizationServerId"`
	Scope                 *string `pulumi:"scope"`
}





type OAuth2AuthenticationSettingsContractInput interface {
	pulumi.Input

	ToOAuth2AuthenticationSettingsContractOutput() OAuth2AuthenticationSettingsContractOutput
	ToOAuth2AuthenticationSettingsContractOutputWithContext(context.Context) OAuth2AuthenticationSettingsContractOutput
}

type OAuth2AuthenticationSettingsContractArgs struct {
	AuthorizationServerId pulumi.StringPtrInput `pulumi:"authorizationServerId"`
	Scope                 pulumi.StringPtrInput `pulumi:"scope"`
}

func (OAuth2AuthenticationSettingsContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OAuth2AuthenticationSettingsContract)(nil)).Elem()
}

func (i OAuth2AuthenticationSettingsContractArgs) ToOAuth2AuthenticationSettingsContractOutput() OAuth2AuthenticationSettingsContractOutput {
	return i.ToOAuth2AuthenticationSettingsContractOutputWithContext(context.Background())
}

func (i OAuth2AuthenticationSettingsContractArgs) ToOAuth2AuthenticationSettingsContractOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OAuth2AuthenticationSettingsContractOutput)
}

func (i OAuth2AuthenticationSettingsContractArgs) ToOAuth2AuthenticationSettingsContractPtrOutput() OAuth2AuthenticationSettingsContractPtrOutput {
	return i.ToOAuth2AuthenticationSettingsContractPtrOutputWithContext(context.Background())
}

func (i OAuth2AuthenticationSettingsContractArgs) ToOAuth2AuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OAuth2AuthenticationSettingsContractOutput).ToOAuth2AuthenticationSettingsContractPtrOutputWithContext(ctx)
}









type OAuth2AuthenticationSettingsContractPtrInput interface {
	pulumi.Input

	ToOAuth2AuthenticationSettingsContractPtrOutput() OAuth2AuthenticationSettingsContractPtrOutput
	ToOAuth2AuthenticationSettingsContractPtrOutputWithContext(context.Context) OAuth2AuthenticationSettingsContractPtrOutput
}

type oauth2AuthenticationSettingsContractPtrType OAuth2AuthenticationSettingsContractArgs

func OAuth2AuthenticationSettingsContractPtr(v *OAuth2AuthenticationSettingsContractArgs) OAuth2AuthenticationSettingsContractPtrInput {
	return (*oauth2AuthenticationSettingsContractPtrType)(v)
}

func (*oauth2AuthenticationSettingsContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OAuth2AuthenticationSettingsContract)(nil)).Elem()
}

func (i *oauth2AuthenticationSettingsContractPtrType) ToOAuth2AuthenticationSettingsContractPtrOutput() OAuth2AuthenticationSettingsContractPtrOutput {
	return i.ToOAuth2AuthenticationSettingsContractPtrOutputWithContext(context.Background())
}

func (i *oauth2AuthenticationSettingsContractPtrType) ToOAuth2AuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OAuth2AuthenticationSettingsContractPtrOutput)
}

type OAuth2AuthenticationSettingsContractOutput struct{ *pulumi.OutputState }

func (OAuth2AuthenticationSettingsContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OAuth2AuthenticationSettingsContract)(nil)).Elem()
}

func (o OAuth2AuthenticationSettingsContractOutput) ToOAuth2AuthenticationSettingsContractOutput() OAuth2AuthenticationSettingsContractOutput {
	return o
}

func (o OAuth2AuthenticationSettingsContractOutput) ToOAuth2AuthenticationSettingsContractOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractOutput {
	return o
}

func (o OAuth2AuthenticationSettingsContractOutput) ToOAuth2AuthenticationSettingsContractPtrOutput() OAuth2AuthenticationSettingsContractPtrOutput {
	return o.ToOAuth2AuthenticationSettingsContractPtrOutputWithContext(context.Background())
}

func (o OAuth2AuthenticationSettingsContractOutput) ToOAuth2AuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OAuth2AuthenticationSettingsContract) *OAuth2AuthenticationSettingsContract {
		return &v
	}).(OAuth2AuthenticationSettingsContractPtrOutput)
}

func (o OAuth2AuthenticationSettingsContractOutput) AuthorizationServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OAuth2AuthenticationSettingsContract) *string { return v.AuthorizationServerId }).(pulumi.StringPtrOutput)
}

func (o OAuth2AuthenticationSettingsContractOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OAuth2AuthenticationSettingsContract) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

type OAuth2AuthenticationSettingsContractPtrOutput struct{ *pulumi.OutputState }

func (OAuth2AuthenticationSettingsContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OAuth2AuthenticationSettingsContract)(nil)).Elem()
}

func (o OAuth2AuthenticationSettingsContractPtrOutput) ToOAuth2AuthenticationSettingsContractPtrOutput() OAuth2AuthenticationSettingsContractPtrOutput {
	return o
}

func (o OAuth2AuthenticationSettingsContractPtrOutput) ToOAuth2AuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractPtrOutput {
	return o
}

func (o OAuth2AuthenticationSettingsContractPtrOutput) Elem() OAuth2AuthenticationSettingsContractOutput {
	return o.ApplyT(func(v *OAuth2AuthenticationSettingsContract) OAuth2AuthenticationSettingsContract {
		if v != nil {
			return *v
		}
		var ret OAuth2AuthenticationSettingsContract
		return ret
	}).(OAuth2AuthenticationSettingsContractOutput)
}

func (o OAuth2AuthenticationSettingsContractPtrOutput) AuthorizationServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OAuth2AuthenticationSettingsContract) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationServerId
	}).(pulumi.StringPtrOutput)
}

func (o OAuth2AuthenticationSettingsContractPtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OAuth2AuthenticationSettingsContract) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

type OAuth2AuthenticationSettingsContractResponse struct {
	AuthorizationServerId *string `pulumi:"authorizationServerId"`
	Scope                 *string `pulumi:"scope"`
}

type OAuth2AuthenticationSettingsContractResponseOutput struct{ *pulumi.OutputState }

func (OAuth2AuthenticationSettingsContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OAuth2AuthenticationSettingsContractResponse)(nil)).Elem()
}

func (o OAuth2AuthenticationSettingsContractResponseOutput) ToOAuth2AuthenticationSettingsContractResponseOutput() OAuth2AuthenticationSettingsContractResponseOutput {
	return o
}

func (o OAuth2AuthenticationSettingsContractResponseOutput) ToOAuth2AuthenticationSettingsContractResponseOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractResponseOutput {
	return o
}

func (o OAuth2AuthenticationSettingsContractResponseOutput) AuthorizationServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OAuth2AuthenticationSettingsContractResponse) *string { return v.AuthorizationServerId }).(pulumi.StringPtrOutput)
}

func (o OAuth2AuthenticationSettingsContractResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OAuth2AuthenticationSettingsContractResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

type OAuth2AuthenticationSettingsContractResponsePtrOutput struct{ *pulumi.OutputState }

func (OAuth2AuthenticationSettingsContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OAuth2AuthenticationSettingsContractResponse)(nil)).Elem()
}

func (o OAuth2AuthenticationSettingsContractResponsePtrOutput) ToOAuth2AuthenticationSettingsContractResponsePtrOutput() OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return o
}

func (o OAuth2AuthenticationSettingsContractResponsePtrOutput) ToOAuth2AuthenticationSettingsContractResponsePtrOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return o
}

func (o OAuth2AuthenticationSettingsContractResponsePtrOutput) Elem() OAuth2AuthenticationSettingsContractResponseOutput {
	return o.ApplyT(func(v *OAuth2AuthenticationSettingsContractResponse) OAuth2AuthenticationSettingsContractResponse {
		if v != nil {
			return *v
		}
		var ret OAuth2AuthenticationSettingsContractResponse
		return ret
	}).(OAuth2AuthenticationSettingsContractResponseOutput)
}

func (o OAuth2AuthenticationSettingsContractResponsePtrOutput) AuthorizationServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OAuth2AuthenticationSettingsContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationServerId
	}).(pulumi.StringPtrOutput)
}

func (o OAuth2AuthenticationSettingsContractResponsePtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OAuth2AuthenticationSettingsContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

type OpenIdAuthenticationSettingsContract struct {
	BearerTokenSendingMethods []string `pulumi:"bearerTokenSendingMethods"`
	OpenidProviderId          *string  `pulumi:"openidProviderId"`
}





type OpenIdAuthenticationSettingsContractInput interface {
	pulumi.Input

	ToOpenIdAuthenticationSettingsContractOutput() OpenIdAuthenticationSettingsContractOutput
	ToOpenIdAuthenticationSettingsContractOutputWithContext(context.Context) OpenIdAuthenticationSettingsContractOutput
}

type OpenIdAuthenticationSettingsContractArgs struct {
	BearerTokenSendingMethods pulumi.StringArrayInput `pulumi:"bearerTokenSendingMethods"`
	OpenidProviderId          pulumi.StringPtrInput   `pulumi:"openidProviderId"`
}

func (OpenIdAuthenticationSettingsContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdAuthenticationSettingsContract)(nil)).Elem()
}

func (i OpenIdAuthenticationSettingsContractArgs) ToOpenIdAuthenticationSettingsContractOutput() OpenIdAuthenticationSettingsContractOutput {
	return i.ToOpenIdAuthenticationSettingsContractOutputWithContext(context.Background())
}

func (i OpenIdAuthenticationSettingsContractArgs) ToOpenIdAuthenticationSettingsContractOutputWithContext(ctx context.Context) OpenIdAuthenticationSettingsContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdAuthenticationSettingsContractOutput)
}

func (i OpenIdAuthenticationSettingsContractArgs) ToOpenIdAuthenticationSettingsContractPtrOutput() OpenIdAuthenticationSettingsContractPtrOutput {
	return i.ToOpenIdAuthenticationSettingsContractPtrOutputWithContext(context.Background())
}

func (i OpenIdAuthenticationSettingsContractArgs) ToOpenIdAuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) OpenIdAuthenticationSettingsContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdAuthenticationSettingsContractOutput).ToOpenIdAuthenticationSettingsContractPtrOutputWithContext(ctx)
}









type OpenIdAuthenticationSettingsContractPtrInput interface {
	pulumi.Input

	ToOpenIdAuthenticationSettingsContractPtrOutput() OpenIdAuthenticationSettingsContractPtrOutput
	ToOpenIdAuthenticationSettingsContractPtrOutputWithContext(context.Context) OpenIdAuthenticationSettingsContractPtrOutput
}

type openIdAuthenticationSettingsContractPtrType OpenIdAuthenticationSettingsContractArgs

func OpenIdAuthenticationSettingsContractPtr(v *OpenIdAuthenticationSettingsContractArgs) OpenIdAuthenticationSettingsContractPtrInput {
	return (*openIdAuthenticationSettingsContractPtrType)(v)
}

func (*openIdAuthenticationSettingsContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdAuthenticationSettingsContract)(nil)).Elem()
}

func (i *openIdAuthenticationSettingsContractPtrType) ToOpenIdAuthenticationSettingsContractPtrOutput() OpenIdAuthenticationSettingsContractPtrOutput {
	return i.ToOpenIdAuthenticationSettingsContractPtrOutputWithContext(context.Background())
}

func (i *openIdAuthenticationSettingsContractPtrType) ToOpenIdAuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) OpenIdAuthenticationSettingsContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdAuthenticationSettingsContractPtrOutput)
}

type OpenIdAuthenticationSettingsContractOutput struct{ *pulumi.OutputState }

func (OpenIdAuthenticationSettingsContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdAuthenticationSettingsContract)(nil)).Elem()
}

func (o OpenIdAuthenticationSettingsContractOutput) ToOpenIdAuthenticationSettingsContractOutput() OpenIdAuthenticationSettingsContractOutput {
	return o
}

func (o OpenIdAuthenticationSettingsContractOutput) ToOpenIdAuthenticationSettingsContractOutputWithContext(ctx context.Context) OpenIdAuthenticationSettingsContractOutput {
	return o
}

func (o OpenIdAuthenticationSettingsContractOutput) ToOpenIdAuthenticationSettingsContractPtrOutput() OpenIdAuthenticationSettingsContractPtrOutput {
	return o.ToOpenIdAuthenticationSettingsContractPtrOutputWithContext(context.Background())
}

func (o OpenIdAuthenticationSettingsContractOutput) ToOpenIdAuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) OpenIdAuthenticationSettingsContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenIdAuthenticationSettingsContract) *OpenIdAuthenticationSettingsContract {
		return &v
	}).(OpenIdAuthenticationSettingsContractPtrOutput)
}

func (o OpenIdAuthenticationSettingsContractOutput) BearerTokenSendingMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OpenIdAuthenticationSettingsContract) []string { return v.BearerTokenSendingMethods }).(pulumi.StringArrayOutput)
}

func (o OpenIdAuthenticationSettingsContractOutput) OpenidProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdAuthenticationSettingsContract) *string { return v.OpenidProviderId }).(pulumi.StringPtrOutput)
}

type OpenIdAuthenticationSettingsContractPtrOutput struct{ *pulumi.OutputState }

func (OpenIdAuthenticationSettingsContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdAuthenticationSettingsContract)(nil)).Elem()
}

func (o OpenIdAuthenticationSettingsContractPtrOutput) ToOpenIdAuthenticationSettingsContractPtrOutput() OpenIdAuthenticationSettingsContractPtrOutput {
	return o
}

func (o OpenIdAuthenticationSettingsContractPtrOutput) ToOpenIdAuthenticationSettingsContractPtrOutputWithContext(ctx context.Context) OpenIdAuthenticationSettingsContractPtrOutput {
	return o
}

func (o OpenIdAuthenticationSettingsContractPtrOutput) Elem() OpenIdAuthenticationSettingsContractOutput {
	return o.ApplyT(func(v *OpenIdAuthenticationSettingsContract) OpenIdAuthenticationSettingsContract {
		if v != nil {
			return *v
		}
		var ret OpenIdAuthenticationSettingsContract
		return ret
	}).(OpenIdAuthenticationSettingsContractOutput)
}

func (o OpenIdAuthenticationSettingsContractPtrOutput) BearerTokenSendingMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OpenIdAuthenticationSettingsContract) []string {
		if v == nil {
			return nil
		}
		return v.BearerTokenSendingMethods
	}).(pulumi.StringArrayOutput)
}

func (o OpenIdAuthenticationSettingsContractPtrOutput) OpenidProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdAuthenticationSettingsContract) *string {
		if v == nil {
			return nil
		}
		return v.OpenidProviderId
	}).(pulumi.StringPtrOutput)
}

type OpenIdAuthenticationSettingsContractResponse struct {
	BearerTokenSendingMethods []string `pulumi:"bearerTokenSendingMethods"`
	OpenidProviderId          *string  `pulumi:"openidProviderId"`
}

type OpenIdAuthenticationSettingsContractResponseOutput struct{ *pulumi.OutputState }

func (OpenIdAuthenticationSettingsContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdAuthenticationSettingsContractResponse)(nil)).Elem()
}

func (o OpenIdAuthenticationSettingsContractResponseOutput) ToOpenIdAuthenticationSettingsContractResponseOutput() OpenIdAuthenticationSettingsContractResponseOutput {
	return o
}

func (o OpenIdAuthenticationSettingsContractResponseOutput) ToOpenIdAuthenticationSettingsContractResponseOutputWithContext(ctx context.Context) OpenIdAuthenticationSettingsContractResponseOutput {
	return o
}

func (o OpenIdAuthenticationSettingsContractResponseOutput) BearerTokenSendingMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OpenIdAuthenticationSettingsContractResponse) []string { return v.BearerTokenSendingMethods }).(pulumi.StringArrayOutput)
}

func (o OpenIdAuthenticationSettingsContractResponseOutput) OpenidProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdAuthenticationSettingsContractResponse) *string { return v.OpenidProviderId }).(pulumi.StringPtrOutput)
}

type OpenIdAuthenticationSettingsContractResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenIdAuthenticationSettingsContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdAuthenticationSettingsContractResponse)(nil)).Elem()
}

func (o OpenIdAuthenticationSettingsContractResponsePtrOutput) ToOpenIdAuthenticationSettingsContractResponsePtrOutput() OpenIdAuthenticationSettingsContractResponsePtrOutput {
	return o
}

func (o OpenIdAuthenticationSettingsContractResponsePtrOutput) ToOpenIdAuthenticationSettingsContractResponsePtrOutputWithContext(ctx context.Context) OpenIdAuthenticationSettingsContractResponsePtrOutput {
	return o
}

func (o OpenIdAuthenticationSettingsContractResponsePtrOutput) Elem() OpenIdAuthenticationSettingsContractResponseOutput {
	return o.ApplyT(func(v *OpenIdAuthenticationSettingsContractResponse) OpenIdAuthenticationSettingsContractResponse {
		if v != nil {
			return *v
		}
		var ret OpenIdAuthenticationSettingsContractResponse
		return ret
	}).(OpenIdAuthenticationSettingsContractResponseOutput)
}

func (o OpenIdAuthenticationSettingsContractResponsePtrOutput) BearerTokenSendingMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OpenIdAuthenticationSettingsContractResponse) []string {
		if v == nil {
			return nil
		}
		return v.BearerTokenSendingMethods
	}).(pulumi.StringArrayOutput)
}

func (o OpenIdAuthenticationSettingsContractResponsePtrOutput) OpenidProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdAuthenticationSettingsContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.OpenidProviderId
	}).(pulumi.StringPtrOutput)
}

type ParameterContract struct {
	DefaultValue *string                             `pulumi:"defaultValue"`
	Description  *string                             `pulumi:"description"`
	Examples     map[string]ParameterExampleContract `pulumi:"examples"`
	Name         string                              `pulumi:"name"`
	Required     *bool                               `pulumi:"required"`
	SchemaId     *string                             `pulumi:"schemaId"`
	Type         string                              `pulumi:"type"`
	TypeName     *string                             `pulumi:"typeName"`
	Values       []string                            `pulumi:"values"`
}





type ParameterContractInput interface {
	pulumi.Input

	ToParameterContractOutput() ParameterContractOutput
	ToParameterContractOutputWithContext(context.Context) ParameterContractOutput
}

type ParameterContractArgs struct {
	DefaultValue pulumi.StringPtrInput            `pulumi:"defaultValue"`
	Description  pulumi.StringPtrInput            `pulumi:"description"`
	Examples     ParameterExampleContractMapInput `pulumi:"examples"`
	Name         pulumi.StringInput               `pulumi:"name"`
	Required     pulumi.BoolPtrInput              `pulumi:"required"`
	SchemaId     pulumi.StringPtrInput            `pulumi:"schemaId"`
	Type         pulumi.StringInput               `pulumi:"type"`
	TypeName     pulumi.StringPtrInput            `pulumi:"typeName"`
	Values       pulumi.StringArrayInput          `pulumi:"values"`
}

func (ParameterContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterContract)(nil)).Elem()
}

func (i ParameterContractArgs) ToParameterContractOutput() ParameterContractOutput {
	return i.ToParameterContractOutputWithContext(context.Background())
}

func (i ParameterContractArgs) ToParameterContractOutputWithContext(ctx context.Context) ParameterContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterContractOutput)
}





type ParameterContractArrayInput interface {
	pulumi.Input

	ToParameterContractArrayOutput() ParameterContractArrayOutput
	ToParameterContractArrayOutputWithContext(context.Context) ParameterContractArrayOutput
}

type ParameterContractArray []ParameterContractInput

func (ParameterContractArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterContract)(nil)).Elem()
}

func (i ParameterContractArray) ToParameterContractArrayOutput() ParameterContractArrayOutput {
	return i.ToParameterContractArrayOutputWithContext(context.Background())
}

func (i ParameterContractArray) ToParameterContractArrayOutputWithContext(ctx context.Context) ParameterContractArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterContractArrayOutput)
}

type ParameterContractOutput struct{ *pulumi.OutputState }

func (ParameterContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterContract)(nil)).Elem()
}

func (o ParameterContractOutput) ToParameterContractOutput() ParameterContractOutput {
	return o
}

func (o ParameterContractOutput) ToParameterContractOutputWithContext(ctx context.Context) ParameterContractOutput {
	return o
}

func (o ParameterContractOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterContract) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o ParameterContractOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterContract) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterContractOutput) Examples() ParameterExampleContractMapOutput {
	return o.ApplyT(func(v ParameterContract) map[string]ParameterExampleContract { return v.Examples }).(ParameterExampleContractMapOutput)
}

func (o ParameterContractOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterContract) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterContractOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ParameterContract) *bool { return v.Required }).(pulumi.BoolPtrOutput)
}

func (o ParameterContractOutput) SchemaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterContract) *string { return v.SchemaId }).(pulumi.StringPtrOutput)
}

func (o ParameterContractOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterContract) string { return v.Type }).(pulumi.StringOutput)
}

func (o ParameterContractOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterContract) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

func (o ParameterContractOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ParameterContract) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ParameterContractArrayOutput struct{ *pulumi.OutputState }

func (ParameterContractArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterContract)(nil)).Elem()
}

func (o ParameterContractArrayOutput) ToParameterContractArrayOutput() ParameterContractArrayOutput {
	return o
}

func (o ParameterContractArrayOutput) ToParameterContractArrayOutputWithContext(ctx context.Context) ParameterContractArrayOutput {
	return o
}

func (o ParameterContractArrayOutput) Index(i pulumi.IntInput) ParameterContractOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterContract {
		return vs[0].([]ParameterContract)[vs[1].(int)]
	}).(ParameterContractOutput)
}

type ParameterContractResponse struct {
	DefaultValue *string                                     `pulumi:"defaultValue"`
	Description  *string                                     `pulumi:"description"`
	Examples     map[string]ParameterExampleContractResponse `pulumi:"examples"`
	Name         string                                      `pulumi:"name"`
	Required     *bool                                       `pulumi:"required"`
	SchemaId     *string                                     `pulumi:"schemaId"`
	Type         string                                      `pulumi:"type"`
	TypeName     *string                                     `pulumi:"typeName"`
	Values       []string                                    `pulumi:"values"`
}

type ParameterContractResponseOutput struct{ *pulumi.OutputState }

func (ParameterContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterContractResponse)(nil)).Elem()
}

func (o ParameterContractResponseOutput) ToParameterContractResponseOutput() ParameterContractResponseOutput {
	return o
}

func (o ParameterContractResponseOutput) ToParameterContractResponseOutputWithContext(ctx context.Context) ParameterContractResponseOutput {
	return o
}

func (o ParameterContractResponseOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterContractResponse) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o ParameterContractResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterContractResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterContractResponseOutput) Examples() ParameterExampleContractResponseMapOutput {
	return o.ApplyT(func(v ParameterContractResponse) map[string]ParameterExampleContractResponse { return v.Examples }).(ParameterExampleContractResponseMapOutput)
}

func (o ParameterContractResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterContractResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterContractResponseOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ParameterContractResponse) *bool { return v.Required }).(pulumi.BoolPtrOutput)
}

func (o ParameterContractResponseOutput) SchemaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterContractResponse) *string { return v.SchemaId }).(pulumi.StringPtrOutput)
}

func (o ParameterContractResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterContractResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ParameterContractResponseOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterContractResponse) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

func (o ParameterContractResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ParameterContractResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ParameterContractResponseArrayOutput struct{ *pulumi.OutputState }

func (ParameterContractResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterContractResponse)(nil)).Elem()
}

func (o ParameterContractResponseArrayOutput) ToParameterContractResponseArrayOutput() ParameterContractResponseArrayOutput {
	return o
}

func (o ParameterContractResponseArrayOutput) ToParameterContractResponseArrayOutputWithContext(ctx context.Context) ParameterContractResponseArrayOutput {
	return o
}

func (o ParameterContractResponseArrayOutput) Index(i pulumi.IntInput) ParameterContractResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterContractResponse {
		return vs[0].([]ParameterContractResponse)[vs[1].(int)]
	}).(ParameterContractResponseOutput)
}

type ParameterExampleContract struct {
	Description   *string     `pulumi:"description"`
	ExternalValue *string     `pulumi:"externalValue"`
	Summary       *string     `pulumi:"summary"`
	Value         interface{} `pulumi:"value"`
}





type ParameterExampleContractInput interface {
	pulumi.Input

	ToParameterExampleContractOutput() ParameterExampleContractOutput
	ToParameterExampleContractOutputWithContext(context.Context) ParameterExampleContractOutput
}

type ParameterExampleContractArgs struct {
	Description   pulumi.StringPtrInput `pulumi:"description"`
	ExternalValue pulumi.StringPtrInput `pulumi:"externalValue"`
	Summary       pulumi.StringPtrInput `pulumi:"summary"`
	Value         pulumi.Input          `pulumi:"value"`
}

func (ParameterExampleContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterExampleContract)(nil)).Elem()
}

func (i ParameterExampleContractArgs) ToParameterExampleContractOutput() ParameterExampleContractOutput {
	return i.ToParameterExampleContractOutputWithContext(context.Background())
}

func (i ParameterExampleContractArgs) ToParameterExampleContractOutputWithContext(ctx context.Context) ParameterExampleContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterExampleContractOutput)
}





type ParameterExampleContractMapInput interface {
	pulumi.Input

	ToParameterExampleContractMapOutput() ParameterExampleContractMapOutput
	ToParameterExampleContractMapOutputWithContext(context.Context) ParameterExampleContractMapOutput
}

type ParameterExampleContractMap map[string]ParameterExampleContractInput

func (ParameterExampleContractMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterExampleContract)(nil)).Elem()
}

func (i ParameterExampleContractMap) ToParameterExampleContractMapOutput() ParameterExampleContractMapOutput {
	return i.ToParameterExampleContractMapOutputWithContext(context.Background())
}

func (i ParameterExampleContractMap) ToParameterExampleContractMapOutputWithContext(ctx context.Context) ParameterExampleContractMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterExampleContractMapOutput)
}

type ParameterExampleContractOutput struct{ *pulumi.OutputState }

func (ParameterExampleContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterExampleContract)(nil)).Elem()
}

func (o ParameterExampleContractOutput) ToParameterExampleContractOutput() ParameterExampleContractOutput {
	return o
}

func (o ParameterExampleContractOutput) ToParameterExampleContractOutputWithContext(ctx context.Context) ParameterExampleContractOutput {
	return o
}

func (o ParameterExampleContractOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterExampleContract) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterExampleContractOutput) ExternalValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterExampleContract) *string { return v.ExternalValue }).(pulumi.StringPtrOutput)
}

func (o ParameterExampleContractOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterExampleContract) *string { return v.Summary }).(pulumi.StringPtrOutput)
}

func (o ParameterExampleContractOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterExampleContract) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ParameterExampleContractMapOutput struct{ *pulumi.OutputState }

func (ParameterExampleContractMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterExampleContract)(nil)).Elem()
}

func (o ParameterExampleContractMapOutput) ToParameterExampleContractMapOutput() ParameterExampleContractMapOutput {
	return o
}

func (o ParameterExampleContractMapOutput) ToParameterExampleContractMapOutputWithContext(ctx context.Context) ParameterExampleContractMapOutput {
	return o
}

func (o ParameterExampleContractMapOutput) MapIndex(k pulumi.StringInput) ParameterExampleContractOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterExampleContract {
		return vs[0].(map[string]ParameterExampleContract)[vs[1].(string)]
	}).(ParameterExampleContractOutput)
}

type ParameterExampleContractResponse struct {
	Description   *string     `pulumi:"description"`
	ExternalValue *string     `pulumi:"externalValue"`
	Summary       *string     `pulumi:"summary"`
	Value         interface{} `pulumi:"value"`
}

type ParameterExampleContractResponseOutput struct{ *pulumi.OutputState }

func (ParameterExampleContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterExampleContractResponse)(nil)).Elem()
}

func (o ParameterExampleContractResponseOutput) ToParameterExampleContractResponseOutput() ParameterExampleContractResponseOutput {
	return o
}

func (o ParameterExampleContractResponseOutput) ToParameterExampleContractResponseOutputWithContext(ctx context.Context) ParameterExampleContractResponseOutput {
	return o
}

func (o ParameterExampleContractResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterExampleContractResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterExampleContractResponseOutput) ExternalValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterExampleContractResponse) *string { return v.ExternalValue }).(pulumi.StringPtrOutput)
}

func (o ParameterExampleContractResponseOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterExampleContractResponse) *string { return v.Summary }).(pulumi.StringPtrOutput)
}

func (o ParameterExampleContractResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterExampleContractResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ParameterExampleContractResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterExampleContractResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterExampleContractResponse)(nil)).Elem()
}

func (o ParameterExampleContractResponseMapOutput) ToParameterExampleContractResponseMapOutput() ParameterExampleContractResponseMapOutput {
	return o
}

func (o ParameterExampleContractResponseMapOutput) ToParameterExampleContractResponseMapOutputWithContext(ctx context.Context) ParameterExampleContractResponseMapOutput {
	return o
}

func (o ParameterExampleContractResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterExampleContractResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterExampleContractResponse {
		return vs[0].(map[string]ParameterExampleContractResponse)[vs[1].(string)]
	}).(ParameterExampleContractResponseOutput)
}

type PipelineDiagnosticSettings struct {
	Request  *HttpMessageDiagnostic `pulumi:"request"`
	Response *HttpMessageDiagnostic `pulumi:"response"`
}





type PipelineDiagnosticSettingsInput interface {
	pulumi.Input

	ToPipelineDiagnosticSettingsOutput() PipelineDiagnosticSettingsOutput
	ToPipelineDiagnosticSettingsOutputWithContext(context.Context) PipelineDiagnosticSettingsOutput
}

type PipelineDiagnosticSettingsArgs struct {
	Request  HttpMessageDiagnosticPtrInput `pulumi:"request"`
	Response HttpMessageDiagnosticPtrInput `pulumi:"response"`
}

func (PipelineDiagnosticSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineDiagnosticSettings)(nil)).Elem()
}

func (i PipelineDiagnosticSettingsArgs) ToPipelineDiagnosticSettingsOutput() PipelineDiagnosticSettingsOutput {
	return i.ToPipelineDiagnosticSettingsOutputWithContext(context.Background())
}

func (i PipelineDiagnosticSettingsArgs) ToPipelineDiagnosticSettingsOutputWithContext(ctx context.Context) PipelineDiagnosticSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineDiagnosticSettingsOutput)
}

func (i PipelineDiagnosticSettingsArgs) ToPipelineDiagnosticSettingsPtrOutput() PipelineDiagnosticSettingsPtrOutput {
	return i.ToPipelineDiagnosticSettingsPtrOutputWithContext(context.Background())
}

func (i PipelineDiagnosticSettingsArgs) ToPipelineDiagnosticSettingsPtrOutputWithContext(ctx context.Context) PipelineDiagnosticSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineDiagnosticSettingsOutput).ToPipelineDiagnosticSettingsPtrOutputWithContext(ctx)
}









type PipelineDiagnosticSettingsPtrInput interface {
	pulumi.Input

	ToPipelineDiagnosticSettingsPtrOutput() PipelineDiagnosticSettingsPtrOutput
	ToPipelineDiagnosticSettingsPtrOutputWithContext(context.Context) PipelineDiagnosticSettingsPtrOutput
}

type pipelineDiagnosticSettingsPtrType PipelineDiagnosticSettingsArgs

func PipelineDiagnosticSettingsPtr(v *PipelineDiagnosticSettingsArgs) PipelineDiagnosticSettingsPtrInput {
	return (*pipelineDiagnosticSettingsPtrType)(v)
}

func (*pipelineDiagnosticSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineDiagnosticSettings)(nil)).Elem()
}

func (i *pipelineDiagnosticSettingsPtrType) ToPipelineDiagnosticSettingsPtrOutput() PipelineDiagnosticSettingsPtrOutput {
	return i.ToPipelineDiagnosticSettingsPtrOutputWithContext(context.Background())
}

func (i *pipelineDiagnosticSettingsPtrType) ToPipelineDiagnosticSettingsPtrOutputWithContext(ctx context.Context) PipelineDiagnosticSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineDiagnosticSettingsPtrOutput)
}

type PipelineDiagnosticSettingsOutput struct{ *pulumi.OutputState }

func (PipelineDiagnosticSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineDiagnosticSettings)(nil)).Elem()
}

func (o PipelineDiagnosticSettingsOutput) ToPipelineDiagnosticSettingsOutput() PipelineDiagnosticSettingsOutput {
	return o
}

func (o PipelineDiagnosticSettingsOutput) ToPipelineDiagnosticSettingsOutputWithContext(ctx context.Context) PipelineDiagnosticSettingsOutput {
	return o
}

func (o PipelineDiagnosticSettingsOutput) ToPipelineDiagnosticSettingsPtrOutput() PipelineDiagnosticSettingsPtrOutput {
	return o.ToPipelineDiagnosticSettingsPtrOutputWithContext(context.Background())
}

func (o PipelineDiagnosticSettingsOutput) ToPipelineDiagnosticSettingsPtrOutputWithContext(ctx context.Context) PipelineDiagnosticSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineDiagnosticSettings) *PipelineDiagnosticSettings {
		return &v
	}).(PipelineDiagnosticSettingsPtrOutput)
}

func (o PipelineDiagnosticSettingsOutput) Request() HttpMessageDiagnosticPtrOutput {
	return o.ApplyT(func(v PipelineDiagnosticSettings) *HttpMessageDiagnostic { return v.Request }).(HttpMessageDiagnosticPtrOutput)
}

func (o PipelineDiagnosticSettingsOutput) Response() HttpMessageDiagnosticPtrOutput {
	return o.ApplyT(func(v PipelineDiagnosticSettings) *HttpMessageDiagnostic { return v.Response }).(HttpMessageDiagnosticPtrOutput)
}

type PipelineDiagnosticSettingsPtrOutput struct{ *pulumi.OutputState }

func (PipelineDiagnosticSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineDiagnosticSettings)(nil)).Elem()
}

func (o PipelineDiagnosticSettingsPtrOutput) ToPipelineDiagnosticSettingsPtrOutput() PipelineDiagnosticSettingsPtrOutput {
	return o
}

func (o PipelineDiagnosticSettingsPtrOutput) ToPipelineDiagnosticSettingsPtrOutputWithContext(ctx context.Context) PipelineDiagnosticSettingsPtrOutput {
	return o
}

func (o PipelineDiagnosticSettingsPtrOutput) Elem() PipelineDiagnosticSettingsOutput {
	return o.ApplyT(func(v *PipelineDiagnosticSettings) PipelineDiagnosticSettings {
		if v != nil {
			return *v
		}
		var ret PipelineDiagnosticSettings
		return ret
	}).(PipelineDiagnosticSettingsOutput)
}

func (o PipelineDiagnosticSettingsPtrOutput) Request() HttpMessageDiagnosticPtrOutput {
	return o.ApplyT(func(v *PipelineDiagnosticSettings) *HttpMessageDiagnostic {
		if v == nil {
			return nil
		}
		return v.Request
	}).(HttpMessageDiagnosticPtrOutput)
}

func (o PipelineDiagnosticSettingsPtrOutput) Response() HttpMessageDiagnosticPtrOutput {
	return o.ApplyT(func(v *PipelineDiagnosticSettings) *HttpMessageDiagnostic {
		if v == nil {
			return nil
		}
		return v.Response
	}).(HttpMessageDiagnosticPtrOutput)
}

type PipelineDiagnosticSettingsResponse struct {
	Request  *HttpMessageDiagnosticResponse `pulumi:"request"`
	Response *HttpMessageDiagnosticResponse `pulumi:"response"`
}

type PipelineDiagnosticSettingsResponseOutput struct{ *pulumi.OutputState }

func (PipelineDiagnosticSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineDiagnosticSettingsResponse)(nil)).Elem()
}

func (o PipelineDiagnosticSettingsResponseOutput) ToPipelineDiagnosticSettingsResponseOutput() PipelineDiagnosticSettingsResponseOutput {
	return o
}

func (o PipelineDiagnosticSettingsResponseOutput) ToPipelineDiagnosticSettingsResponseOutputWithContext(ctx context.Context) PipelineDiagnosticSettingsResponseOutput {
	return o
}

func (o PipelineDiagnosticSettingsResponseOutput) Request() HttpMessageDiagnosticResponsePtrOutput {
	return o.ApplyT(func(v PipelineDiagnosticSettingsResponse) *HttpMessageDiagnosticResponse { return v.Request }).(HttpMessageDiagnosticResponsePtrOutput)
}

func (o PipelineDiagnosticSettingsResponseOutput) Response() HttpMessageDiagnosticResponsePtrOutput {
	return o.ApplyT(func(v PipelineDiagnosticSettingsResponse) *HttpMessageDiagnosticResponse { return v.Response }).(HttpMessageDiagnosticResponsePtrOutput)
}

type PipelineDiagnosticSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineDiagnosticSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineDiagnosticSettingsResponse)(nil)).Elem()
}

func (o PipelineDiagnosticSettingsResponsePtrOutput) ToPipelineDiagnosticSettingsResponsePtrOutput() PipelineDiagnosticSettingsResponsePtrOutput {
	return o
}

func (o PipelineDiagnosticSettingsResponsePtrOutput) ToPipelineDiagnosticSettingsResponsePtrOutputWithContext(ctx context.Context) PipelineDiagnosticSettingsResponsePtrOutput {
	return o
}

func (o PipelineDiagnosticSettingsResponsePtrOutput) Elem() PipelineDiagnosticSettingsResponseOutput {
	return o.ApplyT(func(v *PipelineDiagnosticSettingsResponse) PipelineDiagnosticSettingsResponse {
		if v != nil {
			return *v
		}
		var ret PipelineDiagnosticSettingsResponse
		return ret
	}).(PipelineDiagnosticSettingsResponseOutput)
}

func (o PipelineDiagnosticSettingsResponsePtrOutput) Request() HttpMessageDiagnosticResponsePtrOutput {
	return o.ApplyT(func(v *PipelineDiagnosticSettingsResponse) *HttpMessageDiagnosticResponse {
		if v == nil {
			return nil
		}
		return v.Request
	}).(HttpMessageDiagnosticResponsePtrOutput)
}

func (o PipelineDiagnosticSettingsResponsePtrOutput) Response() HttpMessageDiagnosticResponsePtrOutput {
	return o.ApplyT(func(v *PipelineDiagnosticSettingsResponse) *HttpMessageDiagnosticResponse {
		if v == nil {
			return nil
		}
		return v.Response
	}).(HttpMessageDiagnosticResponsePtrOutput)
}

type PrivateEndpointConnectionRequestProperties struct {
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionRequestPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionRequestPropertiesOutput() PrivateEndpointConnectionRequestPropertiesOutput
	ToPrivateEndpointConnectionRequestPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionRequestPropertiesOutput
}

type PrivateEndpointConnectionRequestPropertiesArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionRequestPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionRequestProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionRequestPropertiesArgs) ToPrivateEndpointConnectionRequestPropertiesOutput() PrivateEndpointConnectionRequestPropertiesOutput {
	return i.ToPrivateEndpointConnectionRequestPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionRequestPropertiesArgs) ToPrivateEndpointConnectionRequestPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionRequestPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionRequestPropertiesOutput)
}

func (i PrivateEndpointConnectionRequestPropertiesArgs) ToPrivateEndpointConnectionRequestPropertiesPtrOutput() PrivateEndpointConnectionRequestPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionRequestPropertiesArgs) ToPrivateEndpointConnectionRequestPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionRequestPropertiesOutput).ToPrivateEndpointConnectionRequestPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionRequestPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionRequestPropertiesPtrOutput() PrivateEndpointConnectionRequestPropertiesPtrOutput
	ToPrivateEndpointConnectionRequestPropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionRequestPropertiesPtrOutput
}

type privateEndpointConnectionRequestPropertiesPtrType PrivateEndpointConnectionRequestPropertiesArgs

func PrivateEndpointConnectionRequestPropertiesPtr(v *PrivateEndpointConnectionRequestPropertiesArgs) PrivateEndpointConnectionRequestPropertiesPtrInput {
	return (*privateEndpointConnectionRequestPropertiesPtrType)(v)
}

func (*privateEndpointConnectionRequestPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionRequestProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionRequestPropertiesPtrType) ToPrivateEndpointConnectionRequestPropertiesPtrOutput() PrivateEndpointConnectionRequestPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionRequestPropertiesPtrType) ToPrivateEndpointConnectionRequestPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionRequestPropertiesPtrOutput)
}

type PrivateEndpointConnectionRequestPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionRequestPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionRequestProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionRequestPropertiesOutput) ToPrivateEndpointConnectionRequestPropertiesOutput() PrivateEndpointConnectionRequestPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionRequestPropertiesOutput) ToPrivateEndpointConnectionRequestPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionRequestPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionRequestPropertiesOutput) ToPrivateEndpointConnectionRequestPropertiesPtrOutput() PrivateEndpointConnectionRequestPropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionRequestPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionRequestPropertiesOutput) ToPrivateEndpointConnectionRequestPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionRequestPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionRequestProperties) *PrivateEndpointConnectionRequestProperties {
		return &v
	}).(PrivateEndpointConnectionRequestPropertiesPtrOutput)
}

func (o PrivateEndpointConnectionRequestPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionRequestProperties) *PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionRequestPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionRequestPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionRequestProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionRequestPropertiesPtrOutput) ToPrivateEndpointConnectionRequestPropertiesPtrOutput() PrivateEndpointConnectionRequestPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionRequestPropertiesPtrOutput) ToPrivateEndpointConnectionRequestPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionRequestPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionRequestPropertiesPtrOutput) Elem() PrivateEndpointConnectionRequestPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionRequestProperties) PrivateEndpointConnectionRequestProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionRequestProperties
		return ret
	}).(PrivateEndpointConnectionRequestPropertiesOutput)
}

func (o PrivateEndpointConnectionRequestPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionRequestProperties) *PrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RemotePrivateEndpointConnectionWrapper struct {
	Id                                *string                           `pulumi:"id"`
	Name                              *string                           `pulumi:"name"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	Type                              *string                           `pulumi:"type"`
}





type RemotePrivateEndpointConnectionWrapperInput interface {
	pulumi.Input

	ToRemotePrivateEndpointConnectionWrapperOutput() RemotePrivateEndpointConnectionWrapperOutput
	ToRemotePrivateEndpointConnectionWrapperOutputWithContext(context.Context) RemotePrivateEndpointConnectionWrapperOutput
}

type RemotePrivateEndpointConnectionWrapperArgs struct {
	Id                                pulumi.StringPtrInput                  `pulumi:"id"`
	Name                              pulumi.StringPtrInput                  `pulumi:"name"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
	Type                              pulumi.StringPtrInput                  `pulumi:"type"`
}

func (RemotePrivateEndpointConnectionWrapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpointConnectionWrapper)(nil)).Elem()
}

func (i RemotePrivateEndpointConnectionWrapperArgs) ToRemotePrivateEndpointConnectionWrapperOutput() RemotePrivateEndpointConnectionWrapperOutput {
	return i.ToRemotePrivateEndpointConnectionWrapperOutputWithContext(context.Background())
}

func (i RemotePrivateEndpointConnectionWrapperArgs) ToRemotePrivateEndpointConnectionWrapperOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionWrapperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemotePrivateEndpointConnectionWrapperOutput)
}





type RemotePrivateEndpointConnectionWrapperArrayInput interface {
	pulumi.Input

	ToRemotePrivateEndpointConnectionWrapperArrayOutput() RemotePrivateEndpointConnectionWrapperArrayOutput
	ToRemotePrivateEndpointConnectionWrapperArrayOutputWithContext(context.Context) RemotePrivateEndpointConnectionWrapperArrayOutput
}

type RemotePrivateEndpointConnectionWrapperArray []RemotePrivateEndpointConnectionWrapperInput

func (RemotePrivateEndpointConnectionWrapperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemotePrivateEndpointConnectionWrapper)(nil)).Elem()
}

func (i RemotePrivateEndpointConnectionWrapperArray) ToRemotePrivateEndpointConnectionWrapperArrayOutput() RemotePrivateEndpointConnectionWrapperArrayOutput {
	return i.ToRemotePrivateEndpointConnectionWrapperArrayOutputWithContext(context.Background())
}

func (i RemotePrivateEndpointConnectionWrapperArray) ToRemotePrivateEndpointConnectionWrapperArrayOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionWrapperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemotePrivateEndpointConnectionWrapperArrayOutput)
}

type RemotePrivateEndpointConnectionWrapperOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionWrapperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpointConnectionWrapper)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionWrapperOutput) ToRemotePrivateEndpointConnectionWrapperOutput() RemotePrivateEndpointConnectionWrapperOutput {
	return o
}

func (o RemotePrivateEndpointConnectionWrapperOutput) ToRemotePrivateEndpointConnectionWrapperOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionWrapperOutput {
	return o
}

func (o RemotePrivateEndpointConnectionWrapperOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapper) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionWrapperOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapper) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionWrapperOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapper) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o RemotePrivateEndpointConnectionWrapperOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapper) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type RemotePrivateEndpointConnectionWrapperArrayOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionWrapperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemotePrivateEndpointConnectionWrapper)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionWrapperArrayOutput) ToRemotePrivateEndpointConnectionWrapperArrayOutput() RemotePrivateEndpointConnectionWrapperArrayOutput {
	return o
}

func (o RemotePrivateEndpointConnectionWrapperArrayOutput) ToRemotePrivateEndpointConnectionWrapperArrayOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionWrapperArrayOutput {
	return o
}

func (o RemotePrivateEndpointConnectionWrapperArrayOutput) Index(i pulumi.IntInput) RemotePrivateEndpointConnectionWrapperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RemotePrivateEndpointConnectionWrapper {
		return vs[0].([]RemotePrivateEndpointConnectionWrapper)[vs[1].(int)]
	}).(RemotePrivateEndpointConnectionWrapperOutput)
}

type RemotePrivateEndpointConnectionWrapperResponse struct {
	GroupIds                          []string                                  `pulumi:"groupIds"`
	Id                                *string                                   `pulumi:"id"`
	Name                              *string                                   `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse                     `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              *string                                   `pulumi:"type"`
}

type RemotePrivateEndpointConnectionWrapperResponseOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionWrapperResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpointConnectionWrapperResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionWrapperResponseOutput) ToRemotePrivateEndpointConnectionWrapperResponseOutput() RemotePrivateEndpointConnectionWrapperResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionWrapperResponseOutput) ToRemotePrivateEndpointConnectionWrapperResponseOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionWrapperResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionWrapperResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapperResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o RemotePrivateEndpointConnectionWrapperResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapperResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionWrapperResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapperResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionWrapperResponseOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapperResponse) *ArmIdWrapperResponse { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionWrapperResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapperResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o RemotePrivateEndpointConnectionWrapperResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapperResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RemotePrivateEndpointConnectionWrapperResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionWrapperResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type RemotePrivateEndpointConnectionWrapperResponseArrayOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionWrapperResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemotePrivateEndpointConnectionWrapperResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionWrapperResponseArrayOutput) ToRemotePrivateEndpointConnectionWrapperResponseArrayOutput() RemotePrivateEndpointConnectionWrapperResponseArrayOutput {
	return o
}

func (o RemotePrivateEndpointConnectionWrapperResponseArrayOutput) ToRemotePrivateEndpointConnectionWrapperResponseArrayOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionWrapperResponseArrayOutput {
	return o
}

func (o RemotePrivateEndpointConnectionWrapperResponseArrayOutput) Index(i pulumi.IntInput) RemotePrivateEndpointConnectionWrapperResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RemotePrivateEndpointConnectionWrapperResponse {
		return vs[0].([]RemotePrivateEndpointConnectionWrapperResponse)[vs[1].(int)]
	}).(RemotePrivateEndpointConnectionWrapperResponseOutput)
}

type RepresentationContract struct {
	ContentType    string                              `pulumi:"contentType"`
	Examples       map[string]ParameterExampleContract `pulumi:"examples"`
	FormParameters []ParameterContract                 `pulumi:"formParameters"`
	SchemaId       *string                             `pulumi:"schemaId"`
	TypeName       *string                             `pulumi:"typeName"`
}





type RepresentationContractInput interface {
	pulumi.Input

	ToRepresentationContractOutput() RepresentationContractOutput
	ToRepresentationContractOutputWithContext(context.Context) RepresentationContractOutput
}

type RepresentationContractArgs struct {
	ContentType    pulumi.StringInput               `pulumi:"contentType"`
	Examples       ParameterExampleContractMapInput `pulumi:"examples"`
	FormParameters ParameterContractArrayInput      `pulumi:"formParameters"`
	SchemaId       pulumi.StringPtrInput            `pulumi:"schemaId"`
	TypeName       pulumi.StringPtrInput            `pulumi:"typeName"`
}

func (RepresentationContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepresentationContract)(nil)).Elem()
}

func (i RepresentationContractArgs) ToRepresentationContractOutput() RepresentationContractOutput {
	return i.ToRepresentationContractOutputWithContext(context.Background())
}

func (i RepresentationContractArgs) ToRepresentationContractOutputWithContext(ctx context.Context) RepresentationContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepresentationContractOutput)
}





type RepresentationContractArrayInput interface {
	pulumi.Input

	ToRepresentationContractArrayOutput() RepresentationContractArrayOutput
	ToRepresentationContractArrayOutputWithContext(context.Context) RepresentationContractArrayOutput
}

type RepresentationContractArray []RepresentationContractInput

func (RepresentationContractArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepresentationContract)(nil)).Elem()
}

func (i RepresentationContractArray) ToRepresentationContractArrayOutput() RepresentationContractArrayOutput {
	return i.ToRepresentationContractArrayOutputWithContext(context.Background())
}

func (i RepresentationContractArray) ToRepresentationContractArrayOutputWithContext(ctx context.Context) RepresentationContractArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepresentationContractArrayOutput)
}

type RepresentationContractOutput struct{ *pulumi.OutputState }

func (RepresentationContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepresentationContract)(nil)).Elem()
}

func (o RepresentationContractOutput) ToRepresentationContractOutput() RepresentationContractOutput {
	return o
}

func (o RepresentationContractOutput) ToRepresentationContractOutputWithContext(ctx context.Context) RepresentationContractOutput {
	return o
}

func (o RepresentationContractOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v RepresentationContract) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o RepresentationContractOutput) Examples() ParameterExampleContractMapOutput {
	return o.ApplyT(func(v RepresentationContract) map[string]ParameterExampleContract { return v.Examples }).(ParameterExampleContractMapOutput)
}

func (o RepresentationContractOutput) FormParameters() ParameterContractArrayOutput {
	return o.ApplyT(func(v RepresentationContract) []ParameterContract { return v.FormParameters }).(ParameterContractArrayOutput)
}

func (o RepresentationContractOutput) SchemaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepresentationContract) *string { return v.SchemaId }).(pulumi.StringPtrOutput)
}

func (o RepresentationContractOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepresentationContract) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

type RepresentationContractArrayOutput struct{ *pulumi.OutputState }

func (RepresentationContractArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepresentationContract)(nil)).Elem()
}

func (o RepresentationContractArrayOutput) ToRepresentationContractArrayOutput() RepresentationContractArrayOutput {
	return o
}

func (o RepresentationContractArrayOutput) ToRepresentationContractArrayOutputWithContext(ctx context.Context) RepresentationContractArrayOutput {
	return o
}

func (o RepresentationContractArrayOutput) Index(i pulumi.IntInput) RepresentationContractOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepresentationContract {
		return vs[0].([]RepresentationContract)[vs[1].(int)]
	}).(RepresentationContractOutput)
}

type RepresentationContractResponse struct {
	ContentType    string                                      `pulumi:"contentType"`
	Examples       map[string]ParameterExampleContractResponse `pulumi:"examples"`
	FormParameters []ParameterContractResponse                 `pulumi:"formParameters"`
	SchemaId       *string                                     `pulumi:"schemaId"`
	TypeName       *string                                     `pulumi:"typeName"`
}

type RepresentationContractResponseOutput struct{ *pulumi.OutputState }

func (RepresentationContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepresentationContractResponse)(nil)).Elem()
}

func (o RepresentationContractResponseOutput) ToRepresentationContractResponseOutput() RepresentationContractResponseOutput {
	return o
}

func (o RepresentationContractResponseOutput) ToRepresentationContractResponseOutputWithContext(ctx context.Context) RepresentationContractResponseOutput {
	return o
}

func (o RepresentationContractResponseOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v RepresentationContractResponse) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o RepresentationContractResponseOutput) Examples() ParameterExampleContractResponseMapOutput {
	return o.ApplyT(func(v RepresentationContractResponse) map[string]ParameterExampleContractResponse { return v.Examples }).(ParameterExampleContractResponseMapOutput)
}

func (o RepresentationContractResponseOutput) FormParameters() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v RepresentationContractResponse) []ParameterContractResponse { return v.FormParameters }).(ParameterContractResponseArrayOutput)
}

func (o RepresentationContractResponseOutput) SchemaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepresentationContractResponse) *string { return v.SchemaId }).(pulumi.StringPtrOutput)
}

func (o RepresentationContractResponseOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepresentationContractResponse) *string { return v.TypeName }).(pulumi.StringPtrOutput)
}

type RepresentationContractResponseArrayOutput struct{ *pulumi.OutputState }

func (RepresentationContractResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepresentationContractResponse)(nil)).Elem()
}

func (o RepresentationContractResponseArrayOutput) ToRepresentationContractResponseArrayOutput() RepresentationContractResponseArrayOutput {
	return o
}

func (o RepresentationContractResponseArrayOutput) ToRepresentationContractResponseArrayOutputWithContext(ctx context.Context) RepresentationContractResponseArrayOutput {
	return o
}

func (o RepresentationContractResponseArrayOutput) Index(i pulumi.IntInput) RepresentationContractResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepresentationContractResponse {
		return vs[0].([]RepresentationContractResponse)[vs[1].(int)]
	}).(RepresentationContractResponseOutput)
}

type RequestContract struct {
	Description     *string                  `pulumi:"description"`
	Headers         []ParameterContract      `pulumi:"headers"`
	QueryParameters []ParameterContract      `pulumi:"queryParameters"`
	Representations []RepresentationContract `pulumi:"representations"`
}





type RequestContractInput interface {
	pulumi.Input

	ToRequestContractOutput() RequestContractOutput
	ToRequestContractOutputWithContext(context.Context) RequestContractOutput
}

type RequestContractArgs struct {
	Description     pulumi.StringPtrInput            `pulumi:"description"`
	Headers         ParameterContractArrayInput      `pulumi:"headers"`
	QueryParameters ParameterContractArrayInput      `pulumi:"queryParameters"`
	Representations RepresentationContractArrayInput `pulumi:"representations"`
}

func (RequestContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestContract)(nil)).Elem()
}

func (i RequestContractArgs) ToRequestContractOutput() RequestContractOutput {
	return i.ToRequestContractOutputWithContext(context.Background())
}

func (i RequestContractArgs) ToRequestContractOutputWithContext(ctx context.Context) RequestContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestContractOutput)
}

func (i RequestContractArgs) ToRequestContractPtrOutput() RequestContractPtrOutput {
	return i.ToRequestContractPtrOutputWithContext(context.Background())
}

func (i RequestContractArgs) ToRequestContractPtrOutputWithContext(ctx context.Context) RequestContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestContractOutput).ToRequestContractPtrOutputWithContext(ctx)
}









type RequestContractPtrInput interface {
	pulumi.Input

	ToRequestContractPtrOutput() RequestContractPtrOutput
	ToRequestContractPtrOutputWithContext(context.Context) RequestContractPtrOutput
}

type requestContractPtrType RequestContractArgs

func RequestContractPtr(v *RequestContractArgs) RequestContractPtrInput {
	return (*requestContractPtrType)(v)
}

func (*requestContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestContract)(nil)).Elem()
}

func (i *requestContractPtrType) ToRequestContractPtrOutput() RequestContractPtrOutput {
	return i.ToRequestContractPtrOutputWithContext(context.Background())
}

func (i *requestContractPtrType) ToRequestContractPtrOutputWithContext(ctx context.Context) RequestContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestContractPtrOutput)
}

type RequestContractOutput struct{ *pulumi.OutputState }

func (RequestContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestContract)(nil)).Elem()
}

func (o RequestContractOutput) ToRequestContractOutput() RequestContractOutput {
	return o
}

func (o RequestContractOutput) ToRequestContractOutputWithContext(ctx context.Context) RequestContractOutput {
	return o
}

func (o RequestContractOutput) ToRequestContractPtrOutput() RequestContractPtrOutput {
	return o.ToRequestContractPtrOutputWithContext(context.Background())
}

func (o RequestContractOutput) ToRequestContractPtrOutputWithContext(ctx context.Context) RequestContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequestContract) *RequestContract {
		return &v
	}).(RequestContractPtrOutput)
}

func (o RequestContractOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestContract) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RequestContractOutput) Headers() ParameterContractArrayOutput {
	return o.ApplyT(func(v RequestContract) []ParameterContract { return v.Headers }).(ParameterContractArrayOutput)
}

func (o RequestContractOutput) QueryParameters() ParameterContractArrayOutput {
	return o.ApplyT(func(v RequestContract) []ParameterContract { return v.QueryParameters }).(ParameterContractArrayOutput)
}

func (o RequestContractOutput) Representations() RepresentationContractArrayOutput {
	return o.ApplyT(func(v RequestContract) []RepresentationContract { return v.Representations }).(RepresentationContractArrayOutput)
}

type RequestContractPtrOutput struct{ *pulumi.OutputState }

func (RequestContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestContract)(nil)).Elem()
}

func (o RequestContractPtrOutput) ToRequestContractPtrOutput() RequestContractPtrOutput {
	return o
}

func (o RequestContractPtrOutput) ToRequestContractPtrOutputWithContext(ctx context.Context) RequestContractPtrOutput {
	return o
}

func (o RequestContractPtrOutput) Elem() RequestContractOutput {
	return o.ApplyT(func(v *RequestContract) RequestContract {
		if v != nil {
			return *v
		}
		var ret RequestContract
		return ret
	}).(RequestContractOutput)
}

func (o RequestContractPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestContract) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o RequestContractPtrOutput) Headers() ParameterContractArrayOutput {
	return o.ApplyT(func(v *RequestContract) []ParameterContract {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(ParameterContractArrayOutput)
}

func (o RequestContractPtrOutput) QueryParameters() ParameterContractArrayOutput {
	return o.ApplyT(func(v *RequestContract) []ParameterContract {
		if v == nil {
			return nil
		}
		return v.QueryParameters
	}).(ParameterContractArrayOutput)
}

func (o RequestContractPtrOutput) Representations() RepresentationContractArrayOutput {
	return o.ApplyT(func(v *RequestContract) []RepresentationContract {
		if v == nil {
			return nil
		}
		return v.Representations
	}).(RepresentationContractArrayOutput)
}

type RequestContractResponse struct {
	Description     *string                          `pulumi:"description"`
	Headers         []ParameterContractResponse      `pulumi:"headers"`
	QueryParameters []ParameterContractResponse      `pulumi:"queryParameters"`
	Representations []RepresentationContractResponse `pulumi:"representations"`
}

type RequestContractResponseOutput struct{ *pulumi.OutputState }

func (RequestContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestContractResponse)(nil)).Elem()
}

func (o RequestContractResponseOutput) ToRequestContractResponseOutput() RequestContractResponseOutput {
	return o
}

func (o RequestContractResponseOutput) ToRequestContractResponseOutputWithContext(ctx context.Context) RequestContractResponseOutput {
	return o
}

func (o RequestContractResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestContractResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RequestContractResponseOutput) Headers() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v RequestContractResponse) []ParameterContractResponse { return v.Headers }).(ParameterContractResponseArrayOutput)
}

func (o RequestContractResponseOutput) QueryParameters() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v RequestContractResponse) []ParameterContractResponse { return v.QueryParameters }).(ParameterContractResponseArrayOutput)
}

func (o RequestContractResponseOutput) Representations() RepresentationContractResponseArrayOutput {
	return o.ApplyT(func(v RequestContractResponse) []RepresentationContractResponse { return v.Representations }).(RepresentationContractResponseArrayOutput)
}

type RequestContractResponsePtrOutput struct{ *pulumi.OutputState }

func (RequestContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestContractResponse)(nil)).Elem()
}

func (o RequestContractResponsePtrOutput) ToRequestContractResponsePtrOutput() RequestContractResponsePtrOutput {
	return o
}

func (o RequestContractResponsePtrOutput) ToRequestContractResponsePtrOutputWithContext(ctx context.Context) RequestContractResponsePtrOutput {
	return o
}

func (o RequestContractResponsePtrOutput) Elem() RequestContractResponseOutput {
	return o.ApplyT(func(v *RequestContractResponse) RequestContractResponse {
		if v != nil {
			return *v
		}
		var ret RequestContractResponse
		return ret
	}).(RequestContractResponseOutput)
}

func (o RequestContractResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o RequestContractResponsePtrOutput) Headers() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v *RequestContractResponse) []ParameterContractResponse {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(ParameterContractResponseArrayOutput)
}

func (o RequestContractResponsePtrOutput) QueryParameters() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v *RequestContractResponse) []ParameterContractResponse {
		if v == nil {
			return nil
		}
		return v.QueryParameters
	}).(ParameterContractResponseArrayOutput)
}

func (o RequestContractResponsePtrOutput) Representations() RepresentationContractResponseArrayOutput {
	return o.ApplyT(func(v *RequestContractResponse) []RepresentationContractResponse {
		if v == nil {
			return nil
		}
		return v.Representations
	}).(RepresentationContractResponseArrayOutput)
}

type ResourceLocationDataContract struct {
	City            *string `pulumi:"city"`
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	District        *string `pulumi:"district"`
	Name            string  `pulumi:"name"`
}





type ResourceLocationDataContractInput interface {
	pulumi.Input

	ToResourceLocationDataContractOutput() ResourceLocationDataContractOutput
	ToResourceLocationDataContractOutputWithContext(context.Context) ResourceLocationDataContractOutput
}

type ResourceLocationDataContractArgs struct {
	City            pulumi.StringPtrInput `pulumi:"city"`
	CountryOrRegion pulumi.StringPtrInput `pulumi:"countryOrRegion"`
	District        pulumi.StringPtrInput `pulumi:"district"`
	Name            pulumi.StringInput    `pulumi:"name"`
}

func (ResourceLocationDataContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLocationDataContract)(nil)).Elem()
}

func (i ResourceLocationDataContractArgs) ToResourceLocationDataContractOutput() ResourceLocationDataContractOutput {
	return i.ToResourceLocationDataContractOutputWithContext(context.Background())
}

func (i ResourceLocationDataContractArgs) ToResourceLocationDataContractOutputWithContext(ctx context.Context) ResourceLocationDataContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLocationDataContractOutput)
}

func (i ResourceLocationDataContractArgs) ToResourceLocationDataContractPtrOutput() ResourceLocationDataContractPtrOutput {
	return i.ToResourceLocationDataContractPtrOutputWithContext(context.Background())
}

func (i ResourceLocationDataContractArgs) ToResourceLocationDataContractPtrOutputWithContext(ctx context.Context) ResourceLocationDataContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLocationDataContractOutput).ToResourceLocationDataContractPtrOutputWithContext(ctx)
}









type ResourceLocationDataContractPtrInput interface {
	pulumi.Input

	ToResourceLocationDataContractPtrOutput() ResourceLocationDataContractPtrOutput
	ToResourceLocationDataContractPtrOutputWithContext(context.Context) ResourceLocationDataContractPtrOutput
}

type resourceLocationDataContractPtrType ResourceLocationDataContractArgs

func ResourceLocationDataContractPtr(v *ResourceLocationDataContractArgs) ResourceLocationDataContractPtrInput {
	return (*resourceLocationDataContractPtrType)(v)
}

func (*resourceLocationDataContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLocationDataContract)(nil)).Elem()
}

func (i *resourceLocationDataContractPtrType) ToResourceLocationDataContractPtrOutput() ResourceLocationDataContractPtrOutput {
	return i.ToResourceLocationDataContractPtrOutputWithContext(context.Background())
}

func (i *resourceLocationDataContractPtrType) ToResourceLocationDataContractPtrOutputWithContext(ctx context.Context) ResourceLocationDataContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLocationDataContractPtrOutput)
}

type ResourceLocationDataContractOutput struct{ *pulumi.OutputState }

func (ResourceLocationDataContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLocationDataContract)(nil)).Elem()
}

func (o ResourceLocationDataContractOutput) ToResourceLocationDataContractOutput() ResourceLocationDataContractOutput {
	return o
}

func (o ResourceLocationDataContractOutput) ToResourceLocationDataContractOutputWithContext(ctx context.Context) ResourceLocationDataContractOutput {
	return o
}

func (o ResourceLocationDataContractOutput) ToResourceLocationDataContractPtrOutput() ResourceLocationDataContractPtrOutput {
	return o.ToResourceLocationDataContractPtrOutputWithContext(context.Background())
}

func (o ResourceLocationDataContractOutput) ToResourceLocationDataContractPtrOutputWithContext(ctx context.Context) ResourceLocationDataContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceLocationDataContract) *ResourceLocationDataContract {
		return &v
	}).(ResourceLocationDataContractPtrOutput)
}

func (o ResourceLocationDataContractOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLocationDataContract) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLocationDataContract) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLocationDataContract) *string { return v.District }).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceLocationDataContract) string { return v.Name }).(pulumi.StringOutput)
}

type ResourceLocationDataContractPtrOutput struct{ *pulumi.OutputState }

func (ResourceLocationDataContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLocationDataContract)(nil)).Elem()
}

func (o ResourceLocationDataContractPtrOutput) ToResourceLocationDataContractPtrOutput() ResourceLocationDataContractPtrOutput {
	return o
}

func (o ResourceLocationDataContractPtrOutput) ToResourceLocationDataContractPtrOutputWithContext(ctx context.Context) ResourceLocationDataContractPtrOutput {
	return o
}

func (o ResourceLocationDataContractPtrOutput) Elem() ResourceLocationDataContractOutput {
	return o.ApplyT(func(v *ResourceLocationDataContract) ResourceLocationDataContract {
		if v != nil {
			return *v
		}
		var ret ResourceLocationDataContract
		return ret
	}).(ResourceLocationDataContractOutput)
}

func (o ResourceLocationDataContractPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceLocationDataContract) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractPtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceLocationDataContract) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractPtrOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceLocationDataContract) *string {
		if v == nil {
			return nil
		}
		return v.District
	}).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceLocationDataContract) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ResourceLocationDataContractResponse struct {
	City            *string `pulumi:"city"`
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	District        *string `pulumi:"district"`
	Name            string  `pulumi:"name"`
}

type ResourceLocationDataContractResponseOutput struct{ *pulumi.OutputState }

func (ResourceLocationDataContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLocationDataContractResponse)(nil)).Elem()
}

func (o ResourceLocationDataContractResponseOutput) ToResourceLocationDataContractResponseOutput() ResourceLocationDataContractResponseOutput {
	return o
}

func (o ResourceLocationDataContractResponseOutput) ToResourceLocationDataContractResponseOutputWithContext(ctx context.Context) ResourceLocationDataContractResponseOutput {
	return o
}

func (o ResourceLocationDataContractResponseOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLocationDataContractResponse) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractResponseOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLocationDataContractResponse) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractResponseOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLocationDataContractResponse) *string { return v.District }).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceLocationDataContractResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ResourceLocationDataContractResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceLocationDataContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLocationDataContractResponse)(nil)).Elem()
}

func (o ResourceLocationDataContractResponsePtrOutput) ToResourceLocationDataContractResponsePtrOutput() ResourceLocationDataContractResponsePtrOutput {
	return o
}

func (o ResourceLocationDataContractResponsePtrOutput) ToResourceLocationDataContractResponsePtrOutputWithContext(ctx context.Context) ResourceLocationDataContractResponsePtrOutput {
	return o
}

func (o ResourceLocationDataContractResponsePtrOutput) Elem() ResourceLocationDataContractResponseOutput {
	return o.ApplyT(func(v *ResourceLocationDataContractResponse) ResourceLocationDataContractResponse {
		if v != nil {
			return *v
		}
		var ret ResourceLocationDataContractResponse
		return ret
	}).(ResourceLocationDataContractResponseOutput)
}

func (o ResourceLocationDataContractResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceLocationDataContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractResponsePtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceLocationDataContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractResponsePtrOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceLocationDataContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.District
	}).(pulumi.StringPtrOutput)
}

func (o ResourceLocationDataContractResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceLocationDataContractResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ResponseContract struct {
	Description     *string                  `pulumi:"description"`
	Headers         []ParameterContract      `pulumi:"headers"`
	Representations []RepresentationContract `pulumi:"representations"`
	StatusCode      int                      `pulumi:"statusCode"`
}





type ResponseContractInput interface {
	pulumi.Input

	ToResponseContractOutput() ResponseContractOutput
	ToResponseContractOutputWithContext(context.Context) ResponseContractOutput
}

type ResponseContractArgs struct {
	Description     pulumi.StringPtrInput            `pulumi:"description"`
	Headers         ParameterContractArrayInput      `pulumi:"headers"`
	Representations RepresentationContractArrayInput `pulumi:"representations"`
	StatusCode      pulumi.IntInput                  `pulumi:"statusCode"`
}

func (ResponseContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseContract)(nil)).Elem()
}

func (i ResponseContractArgs) ToResponseContractOutput() ResponseContractOutput {
	return i.ToResponseContractOutputWithContext(context.Background())
}

func (i ResponseContractArgs) ToResponseContractOutputWithContext(ctx context.Context) ResponseContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseContractOutput)
}





type ResponseContractArrayInput interface {
	pulumi.Input

	ToResponseContractArrayOutput() ResponseContractArrayOutput
	ToResponseContractArrayOutputWithContext(context.Context) ResponseContractArrayOutput
}

type ResponseContractArray []ResponseContractInput

func (ResponseContractArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResponseContract)(nil)).Elem()
}

func (i ResponseContractArray) ToResponseContractArrayOutput() ResponseContractArrayOutput {
	return i.ToResponseContractArrayOutputWithContext(context.Background())
}

func (i ResponseContractArray) ToResponseContractArrayOutputWithContext(ctx context.Context) ResponseContractArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseContractArrayOutput)
}

type ResponseContractOutput struct{ *pulumi.OutputState }

func (ResponseContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseContract)(nil)).Elem()
}

func (o ResponseContractOutput) ToResponseContractOutput() ResponseContractOutput {
	return o
}

func (o ResponseContractOutput) ToResponseContractOutputWithContext(ctx context.Context) ResponseContractOutput {
	return o
}

func (o ResponseContractOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseContract) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ResponseContractOutput) Headers() ParameterContractArrayOutput {
	return o.ApplyT(func(v ResponseContract) []ParameterContract { return v.Headers }).(ParameterContractArrayOutput)
}

func (o ResponseContractOutput) Representations() RepresentationContractArrayOutput {
	return o.ApplyT(func(v ResponseContract) []RepresentationContract { return v.Representations }).(RepresentationContractArrayOutput)
}

func (o ResponseContractOutput) StatusCode() pulumi.IntOutput {
	return o.ApplyT(func(v ResponseContract) int { return v.StatusCode }).(pulumi.IntOutput)
}

type ResponseContractArrayOutput struct{ *pulumi.OutputState }

func (ResponseContractArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResponseContract)(nil)).Elem()
}

func (o ResponseContractArrayOutput) ToResponseContractArrayOutput() ResponseContractArrayOutput {
	return o
}

func (o ResponseContractArrayOutput) ToResponseContractArrayOutputWithContext(ctx context.Context) ResponseContractArrayOutput {
	return o
}

func (o ResponseContractArrayOutput) Index(i pulumi.IntInput) ResponseContractOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResponseContract {
		return vs[0].([]ResponseContract)[vs[1].(int)]
	}).(ResponseContractOutput)
}

type ResponseContractResponse struct {
	Description     *string                          `pulumi:"description"`
	Headers         []ParameterContractResponse      `pulumi:"headers"`
	Representations []RepresentationContractResponse `pulumi:"representations"`
	StatusCode      int                              `pulumi:"statusCode"`
}

type ResponseContractResponseOutput struct{ *pulumi.OutputState }

func (ResponseContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseContractResponse)(nil)).Elem()
}

func (o ResponseContractResponseOutput) ToResponseContractResponseOutput() ResponseContractResponseOutput {
	return o
}

func (o ResponseContractResponseOutput) ToResponseContractResponseOutputWithContext(ctx context.Context) ResponseContractResponseOutput {
	return o
}

func (o ResponseContractResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseContractResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ResponseContractResponseOutput) Headers() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v ResponseContractResponse) []ParameterContractResponse { return v.Headers }).(ParameterContractResponseArrayOutput)
}

func (o ResponseContractResponseOutput) Representations() RepresentationContractResponseArrayOutput {
	return o.ApplyT(func(v ResponseContractResponse) []RepresentationContractResponse { return v.Representations }).(RepresentationContractResponseArrayOutput)
}

func (o ResponseContractResponseOutput) StatusCode() pulumi.IntOutput {
	return o.ApplyT(func(v ResponseContractResponse) int { return v.StatusCode }).(pulumi.IntOutput)
}

type ResponseContractResponseArrayOutput struct{ *pulumi.OutputState }

func (ResponseContractResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResponseContractResponse)(nil)).Elem()
}

func (o ResponseContractResponseArrayOutput) ToResponseContractResponseArrayOutput() ResponseContractResponseArrayOutput {
	return o
}

func (o ResponseContractResponseArrayOutput) ToResponseContractResponseArrayOutputWithContext(ctx context.Context) ResponseContractResponseArrayOutput {
	return o
}

func (o ResponseContractResponseArrayOutput) Index(i pulumi.IntInput) ResponseContractResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResponseContractResponse {
		return vs[0].([]ResponseContractResponse)[vs[1].(int)]
	}).(ResponseContractResponseOutput)
}

type SamplingSettings struct {
	Percentage   *float64 `pulumi:"percentage"`
	SamplingType *string  `pulumi:"samplingType"`
}





type SamplingSettingsInput interface {
	pulumi.Input

	ToSamplingSettingsOutput() SamplingSettingsOutput
	ToSamplingSettingsOutputWithContext(context.Context) SamplingSettingsOutput
}

type SamplingSettingsArgs struct {
	Percentage   pulumi.Float64PtrInput `pulumi:"percentage"`
	SamplingType pulumi.StringPtrInput  `pulumi:"samplingType"`
}

func (SamplingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SamplingSettings)(nil)).Elem()
}

func (i SamplingSettingsArgs) ToSamplingSettingsOutput() SamplingSettingsOutput {
	return i.ToSamplingSettingsOutputWithContext(context.Background())
}

func (i SamplingSettingsArgs) ToSamplingSettingsOutputWithContext(ctx context.Context) SamplingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamplingSettingsOutput)
}

func (i SamplingSettingsArgs) ToSamplingSettingsPtrOutput() SamplingSettingsPtrOutput {
	return i.ToSamplingSettingsPtrOutputWithContext(context.Background())
}

func (i SamplingSettingsArgs) ToSamplingSettingsPtrOutputWithContext(ctx context.Context) SamplingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamplingSettingsOutput).ToSamplingSettingsPtrOutputWithContext(ctx)
}









type SamplingSettingsPtrInput interface {
	pulumi.Input

	ToSamplingSettingsPtrOutput() SamplingSettingsPtrOutput
	ToSamplingSettingsPtrOutputWithContext(context.Context) SamplingSettingsPtrOutput
}

type samplingSettingsPtrType SamplingSettingsArgs

func SamplingSettingsPtr(v *SamplingSettingsArgs) SamplingSettingsPtrInput {
	return (*samplingSettingsPtrType)(v)
}

func (*samplingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SamplingSettings)(nil)).Elem()
}

func (i *samplingSettingsPtrType) ToSamplingSettingsPtrOutput() SamplingSettingsPtrOutput {
	return i.ToSamplingSettingsPtrOutputWithContext(context.Background())
}

func (i *samplingSettingsPtrType) ToSamplingSettingsPtrOutputWithContext(ctx context.Context) SamplingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamplingSettingsPtrOutput)
}

type SamplingSettingsOutput struct{ *pulumi.OutputState }

func (SamplingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SamplingSettings)(nil)).Elem()
}

func (o SamplingSettingsOutput) ToSamplingSettingsOutput() SamplingSettingsOutput {
	return o
}

func (o SamplingSettingsOutput) ToSamplingSettingsOutputWithContext(ctx context.Context) SamplingSettingsOutput {
	return o
}

func (o SamplingSettingsOutput) ToSamplingSettingsPtrOutput() SamplingSettingsPtrOutput {
	return o.ToSamplingSettingsPtrOutputWithContext(context.Background())
}

func (o SamplingSettingsOutput) ToSamplingSettingsPtrOutputWithContext(ctx context.Context) SamplingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SamplingSettings) *SamplingSettings {
		return &v
	}).(SamplingSettingsPtrOutput)
}

func (o SamplingSettingsOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SamplingSettings) *float64 { return v.Percentage }).(pulumi.Float64PtrOutput)
}

func (o SamplingSettingsOutput) SamplingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SamplingSettings) *string { return v.SamplingType }).(pulumi.StringPtrOutput)
}

type SamplingSettingsPtrOutput struct{ *pulumi.OutputState }

func (SamplingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamplingSettings)(nil)).Elem()
}

func (o SamplingSettingsPtrOutput) ToSamplingSettingsPtrOutput() SamplingSettingsPtrOutput {
	return o
}

func (o SamplingSettingsPtrOutput) ToSamplingSettingsPtrOutputWithContext(ctx context.Context) SamplingSettingsPtrOutput {
	return o
}

func (o SamplingSettingsPtrOutput) Elem() SamplingSettingsOutput {
	return o.ApplyT(func(v *SamplingSettings) SamplingSettings {
		if v != nil {
			return *v
		}
		var ret SamplingSettings
		return ret
	}).(SamplingSettingsOutput)
}

func (o SamplingSettingsPtrOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SamplingSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.Float64PtrOutput)
}

func (o SamplingSettingsPtrOutput) SamplingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SamplingSettings) *string {
		if v == nil {
			return nil
		}
		return v.SamplingType
	}).(pulumi.StringPtrOutput)
}

type SamplingSettingsResponse struct {
	Percentage   *float64 `pulumi:"percentage"`
	SamplingType *string  `pulumi:"samplingType"`
}

type SamplingSettingsResponseOutput struct{ *pulumi.OutputState }

func (SamplingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SamplingSettingsResponse)(nil)).Elem()
}

func (o SamplingSettingsResponseOutput) ToSamplingSettingsResponseOutput() SamplingSettingsResponseOutput {
	return o
}

func (o SamplingSettingsResponseOutput) ToSamplingSettingsResponseOutputWithContext(ctx context.Context) SamplingSettingsResponseOutput {
	return o
}

func (o SamplingSettingsResponseOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SamplingSettingsResponse) *float64 { return v.Percentage }).(pulumi.Float64PtrOutput)
}

func (o SamplingSettingsResponseOutput) SamplingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SamplingSettingsResponse) *string { return v.SamplingType }).(pulumi.StringPtrOutput)
}

type SamplingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SamplingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamplingSettingsResponse)(nil)).Elem()
}

func (o SamplingSettingsResponsePtrOutput) ToSamplingSettingsResponsePtrOutput() SamplingSettingsResponsePtrOutput {
	return o
}

func (o SamplingSettingsResponsePtrOutput) ToSamplingSettingsResponsePtrOutputWithContext(ctx context.Context) SamplingSettingsResponsePtrOutput {
	return o
}

func (o SamplingSettingsResponsePtrOutput) Elem() SamplingSettingsResponseOutput {
	return o.ApplyT(func(v *SamplingSettingsResponse) SamplingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret SamplingSettingsResponse
		return ret
	}).(SamplingSettingsResponseOutput)
}

func (o SamplingSettingsResponsePtrOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SamplingSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.Float64PtrOutput)
}

func (o SamplingSettingsResponsePtrOutput) SamplingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SamplingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SamplingType
	}).(pulumi.StringPtrOutput)
}

type SubscriptionKeyParameterNamesContract struct {
	Header *string `pulumi:"header"`
	Query  *string `pulumi:"query"`
}





type SubscriptionKeyParameterNamesContractInput interface {
	pulumi.Input

	ToSubscriptionKeyParameterNamesContractOutput() SubscriptionKeyParameterNamesContractOutput
	ToSubscriptionKeyParameterNamesContractOutputWithContext(context.Context) SubscriptionKeyParameterNamesContractOutput
}

type SubscriptionKeyParameterNamesContractArgs struct {
	Header pulumi.StringPtrInput `pulumi:"header"`
	Query  pulumi.StringPtrInput `pulumi:"query"`
}

func (SubscriptionKeyParameterNamesContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionKeyParameterNamesContract)(nil)).Elem()
}

func (i SubscriptionKeyParameterNamesContractArgs) ToSubscriptionKeyParameterNamesContractOutput() SubscriptionKeyParameterNamesContractOutput {
	return i.ToSubscriptionKeyParameterNamesContractOutputWithContext(context.Background())
}

func (i SubscriptionKeyParameterNamesContractArgs) ToSubscriptionKeyParameterNamesContractOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionKeyParameterNamesContractOutput)
}

func (i SubscriptionKeyParameterNamesContractArgs) ToSubscriptionKeyParameterNamesContractPtrOutput() SubscriptionKeyParameterNamesContractPtrOutput {
	return i.ToSubscriptionKeyParameterNamesContractPtrOutputWithContext(context.Background())
}

func (i SubscriptionKeyParameterNamesContractArgs) ToSubscriptionKeyParameterNamesContractPtrOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionKeyParameterNamesContractOutput).ToSubscriptionKeyParameterNamesContractPtrOutputWithContext(ctx)
}









type SubscriptionKeyParameterNamesContractPtrInput interface {
	pulumi.Input

	ToSubscriptionKeyParameterNamesContractPtrOutput() SubscriptionKeyParameterNamesContractPtrOutput
	ToSubscriptionKeyParameterNamesContractPtrOutputWithContext(context.Context) SubscriptionKeyParameterNamesContractPtrOutput
}

type subscriptionKeyParameterNamesContractPtrType SubscriptionKeyParameterNamesContractArgs

func SubscriptionKeyParameterNamesContractPtr(v *SubscriptionKeyParameterNamesContractArgs) SubscriptionKeyParameterNamesContractPtrInput {
	return (*subscriptionKeyParameterNamesContractPtrType)(v)
}

func (*subscriptionKeyParameterNamesContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionKeyParameterNamesContract)(nil)).Elem()
}

func (i *subscriptionKeyParameterNamesContractPtrType) ToSubscriptionKeyParameterNamesContractPtrOutput() SubscriptionKeyParameterNamesContractPtrOutput {
	return i.ToSubscriptionKeyParameterNamesContractPtrOutputWithContext(context.Background())
}

func (i *subscriptionKeyParameterNamesContractPtrType) ToSubscriptionKeyParameterNamesContractPtrOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionKeyParameterNamesContractPtrOutput)
}

type SubscriptionKeyParameterNamesContractOutput struct{ *pulumi.OutputState }

func (SubscriptionKeyParameterNamesContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionKeyParameterNamesContract)(nil)).Elem()
}

func (o SubscriptionKeyParameterNamesContractOutput) ToSubscriptionKeyParameterNamesContractOutput() SubscriptionKeyParameterNamesContractOutput {
	return o
}

func (o SubscriptionKeyParameterNamesContractOutput) ToSubscriptionKeyParameterNamesContractOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractOutput {
	return o
}

func (o SubscriptionKeyParameterNamesContractOutput) ToSubscriptionKeyParameterNamesContractPtrOutput() SubscriptionKeyParameterNamesContractPtrOutput {
	return o.ToSubscriptionKeyParameterNamesContractPtrOutputWithContext(context.Background())
}

func (o SubscriptionKeyParameterNamesContractOutput) ToSubscriptionKeyParameterNamesContractPtrOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionKeyParameterNamesContract) *SubscriptionKeyParameterNamesContract {
		return &v
	}).(SubscriptionKeyParameterNamesContractPtrOutput)
}

func (o SubscriptionKeyParameterNamesContractOutput) Header() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionKeyParameterNamesContract) *string { return v.Header }).(pulumi.StringPtrOutput)
}

func (o SubscriptionKeyParameterNamesContractOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionKeyParameterNamesContract) *string { return v.Query }).(pulumi.StringPtrOutput)
}

type SubscriptionKeyParameterNamesContractPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionKeyParameterNamesContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionKeyParameterNamesContract)(nil)).Elem()
}

func (o SubscriptionKeyParameterNamesContractPtrOutput) ToSubscriptionKeyParameterNamesContractPtrOutput() SubscriptionKeyParameterNamesContractPtrOutput {
	return o
}

func (o SubscriptionKeyParameterNamesContractPtrOutput) ToSubscriptionKeyParameterNamesContractPtrOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractPtrOutput {
	return o
}

func (o SubscriptionKeyParameterNamesContractPtrOutput) Elem() SubscriptionKeyParameterNamesContractOutput {
	return o.ApplyT(func(v *SubscriptionKeyParameterNamesContract) SubscriptionKeyParameterNamesContract {
		if v != nil {
			return *v
		}
		var ret SubscriptionKeyParameterNamesContract
		return ret
	}).(SubscriptionKeyParameterNamesContractOutput)
}

func (o SubscriptionKeyParameterNamesContractPtrOutput) Header() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionKeyParameterNamesContract) *string {
		if v == nil {
			return nil
		}
		return v.Header
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionKeyParameterNamesContractPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionKeyParameterNamesContract) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

type SubscriptionKeyParameterNamesContractResponse struct {
	Header *string `pulumi:"header"`
	Query  *string `pulumi:"query"`
}

type SubscriptionKeyParameterNamesContractResponseOutput struct{ *pulumi.OutputState }

func (SubscriptionKeyParameterNamesContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionKeyParameterNamesContractResponse)(nil)).Elem()
}

func (o SubscriptionKeyParameterNamesContractResponseOutput) ToSubscriptionKeyParameterNamesContractResponseOutput() SubscriptionKeyParameterNamesContractResponseOutput {
	return o
}

func (o SubscriptionKeyParameterNamesContractResponseOutput) ToSubscriptionKeyParameterNamesContractResponseOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractResponseOutput {
	return o
}

func (o SubscriptionKeyParameterNamesContractResponseOutput) Header() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionKeyParameterNamesContractResponse) *string { return v.Header }).(pulumi.StringPtrOutput)
}

func (o SubscriptionKeyParameterNamesContractResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionKeyParameterNamesContractResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

type SubscriptionKeyParameterNamesContractResponsePtrOutput struct{ *pulumi.OutputState }

func (SubscriptionKeyParameterNamesContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionKeyParameterNamesContractResponse)(nil)).Elem()
}

func (o SubscriptionKeyParameterNamesContractResponsePtrOutput) ToSubscriptionKeyParameterNamesContractResponsePtrOutput() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o
}

func (o SubscriptionKeyParameterNamesContractResponsePtrOutput) ToSubscriptionKeyParameterNamesContractResponsePtrOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o
}

func (o SubscriptionKeyParameterNamesContractResponsePtrOutput) Elem() SubscriptionKeyParameterNamesContractResponseOutput {
	return o.ApplyT(func(v *SubscriptionKeyParameterNamesContractResponse) SubscriptionKeyParameterNamesContractResponse {
		if v != nil {
			return *v
		}
		var ret SubscriptionKeyParameterNamesContractResponse
		return ret
	}).(SubscriptionKeyParameterNamesContractResponseOutput)
}

func (o SubscriptionKeyParameterNamesContractResponsePtrOutput) Header() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionKeyParameterNamesContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.Header
	}).(pulumi.StringPtrOutput)
}

func (o SubscriptionKeyParameterNamesContractResponsePtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscriptionKeyParameterNamesContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
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

type TokenBodyParameterContract struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type TokenBodyParameterContractInput interface {
	pulumi.Input

	ToTokenBodyParameterContractOutput() TokenBodyParameterContractOutput
	ToTokenBodyParameterContractOutputWithContext(context.Context) TokenBodyParameterContractOutput
}

type TokenBodyParameterContractArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TokenBodyParameterContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenBodyParameterContract)(nil)).Elem()
}

func (i TokenBodyParameterContractArgs) ToTokenBodyParameterContractOutput() TokenBodyParameterContractOutput {
	return i.ToTokenBodyParameterContractOutputWithContext(context.Background())
}

func (i TokenBodyParameterContractArgs) ToTokenBodyParameterContractOutputWithContext(ctx context.Context) TokenBodyParameterContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenBodyParameterContractOutput)
}





type TokenBodyParameterContractArrayInput interface {
	pulumi.Input

	ToTokenBodyParameterContractArrayOutput() TokenBodyParameterContractArrayOutput
	ToTokenBodyParameterContractArrayOutputWithContext(context.Context) TokenBodyParameterContractArrayOutput
}

type TokenBodyParameterContractArray []TokenBodyParameterContractInput

func (TokenBodyParameterContractArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenBodyParameterContract)(nil)).Elem()
}

func (i TokenBodyParameterContractArray) ToTokenBodyParameterContractArrayOutput() TokenBodyParameterContractArrayOutput {
	return i.ToTokenBodyParameterContractArrayOutputWithContext(context.Background())
}

func (i TokenBodyParameterContractArray) ToTokenBodyParameterContractArrayOutputWithContext(ctx context.Context) TokenBodyParameterContractArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenBodyParameterContractArrayOutput)
}

type TokenBodyParameterContractOutput struct{ *pulumi.OutputState }

func (TokenBodyParameterContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenBodyParameterContract)(nil)).Elem()
}

func (o TokenBodyParameterContractOutput) ToTokenBodyParameterContractOutput() TokenBodyParameterContractOutput {
	return o
}

func (o TokenBodyParameterContractOutput) ToTokenBodyParameterContractOutputWithContext(ctx context.Context) TokenBodyParameterContractOutput {
	return o
}

func (o TokenBodyParameterContractOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TokenBodyParameterContract) string { return v.Name }).(pulumi.StringOutput)
}

func (o TokenBodyParameterContractOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TokenBodyParameterContract) string { return v.Value }).(pulumi.StringOutput)
}

type TokenBodyParameterContractArrayOutput struct{ *pulumi.OutputState }

func (TokenBodyParameterContractArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenBodyParameterContract)(nil)).Elem()
}

func (o TokenBodyParameterContractArrayOutput) ToTokenBodyParameterContractArrayOutput() TokenBodyParameterContractArrayOutput {
	return o
}

func (o TokenBodyParameterContractArrayOutput) ToTokenBodyParameterContractArrayOutputWithContext(ctx context.Context) TokenBodyParameterContractArrayOutput {
	return o
}

func (o TokenBodyParameterContractArrayOutput) Index(i pulumi.IntInput) TokenBodyParameterContractOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenBodyParameterContract {
		return vs[0].([]TokenBodyParameterContract)[vs[1].(int)]
	}).(TokenBodyParameterContractOutput)
}

type TokenBodyParameterContractResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type TokenBodyParameterContractResponseOutput struct{ *pulumi.OutputState }

func (TokenBodyParameterContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenBodyParameterContractResponse)(nil)).Elem()
}

func (o TokenBodyParameterContractResponseOutput) ToTokenBodyParameterContractResponseOutput() TokenBodyParameterContractResponseOutput {
	return o
}

func (o TokenBodyParameterContractResponseOutput) ToTokenBodyParameterContractResponseOutputWithContext(ctx context.Context) TokenBodyParameterContractResponseOutput {
	return o
}

func (o TokenBodyParameterContractResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TokenBodyParameterContractResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TokenBodyParameterContractResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TokenBodyParameterContractResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TokenBodyParameterContractResponseArrayOutput struct{ *pulumi.OutputState }

func (TokenBodyParameterContractResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenBodyParameterContractResponse)(nil)).Elem()
}

func (o TokenBodyParameterContractResponseArrayOutput) ToTokenBodyParameterContractResponseArrayOutput() TokenBodyParameterContractResponseArrayOutput {
	return o
}

func (o TokenBodyParameterContractResponseArrayOutput) ToTokenBodyParameterContractResponseArrayOutputWithContext(ctx context.Context) TokenBodyParameterContractResponseArrayOutput {
	return o
}

func (o TokenBodyParameterContractResponseArrayOutput) Index(i pulumi.IntInput) TokenBodyParameterContractResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenBodyParameterContractResponse {
		return vs[0].([]TokenBodyParameterContractResponse)[vs[1].(int)]
	}).(TokenBodyParameterContractResponseOutput)
}

type UserIdentityContract struct {
	Id       *string `pulumi:"id"`
	Provider *string `pulumi:"provider"`
}





type UserIdentityContractInput interface {
	pulumi.Input

	ToUserIdentityContractOutput() UserIdentityContractOutput
	ToUserIdentityContractOutputWithContext(context.Context) UserIdentityContractOutput
}

type UserIdentityContractArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Provider pulumi.StringPtrInput `pulumi:"provider"`
}

func (UserIdentityContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityContract)(nil)).Elem()
}

func (i UserIdentityContractArgs) ToUserIdentityContractOutput() UserIdentityContractOutput {
	return i.ToUserIdentityContractOutputWithContext(context.Background())
}

func (i UserIdentityContractArgs) ToUserIdentityContractOutputWithContext(ctx context.Context) UserIdentityContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityContractOutput)
}





type UserIdentityContractArrayInput interface {
	pulumi.Input

	ToUserIdentityContractArrayOutput() UserIdentityContractArrayOutput
	ToUserIdentityContractArrayOutputWithContext(context.Context) UserIdentityContractArrayOutput
}

type UserIdentityContractArray []UserIdentityContractInput

func (UserIdentityContractArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserIdentityContract)(nil)).Elem()
}

func (i UserIdentityContractArray) ToUserIdentityContractArrayOutput() UserIdentityContractArrayOutput {
	return i.ToUserIdentityContractArrayOutputWithContext(context.Background())
}

func (i UserIdentityContractArray) ToUserIdentityContractArrayOutputWithContext(ctx context.Context) UserIdentityContractArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityContractArrayOutput)
}

type UserIdentityContractOutput struct{ *pulumi.OutputState }

func (UserIdentityContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityContract)(nil)).Elem()
}

func (o UserIdentityContractOutput) ToUserIdentityContractOutput() UserIdentityContractOutput {
	return o
}

func (o UserIdentityContractOutput) ToUserIdentityContractOutputWithContext(ctx context.Context) UserIdentityContractOutput {
	return o
}

func (o UserIdentityContractOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityContract) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o UserIdentityContractOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityContract) *string { return v.Provider }).(pulumi.StringPtrOutput)
}

type UserIdentityContractArrayOutput struct{ *pulumi.OutputState }

func (UserIdentityContractArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserIdentityContract)(nil)).Elem()
}

func (o UserIdentityContractArrayOutput) ToUserIdentityContractArrayOutput() UserIdentityContractArrayOutput {
	return o
}

func (o UserIdentityContractArrayOutput) ToUserIdentityContractArrayOutputWithContext(ctx context.Context) UserIdentityContractArrayOutput {
	return o
}

func (o UserIdentityContractArrayOutput) Index(i pulumi.IntInput) UserIdentityContractOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserIdentityContract {
		return vs[0].([]UserIdentityContract)[vs[1].(int)]
	}).(UserIdentityContractOutput)
}

type UserIdentityContractResponse struct {
	Id       *string `pulumi:"id"`
	Provider *string `pulumi:"provider"`
}

type UserIdentityContractResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityContractResponse)(nil)).Elem()
}

func (o UserIdentityContractResponseOutput) ToUserIdentityContractResponseOutput() UserIdentityContractResponseOutput {
	return o
}

func (o UserIdentityContractResponseOutput) ToUserIdentityContractResponseOutputWithContext(ctx context.Context) UserIdentityContractResponseOutput {
	return o
}

func (o UserIdentityContractResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityContractResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o UserIdentityContractResponseOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityContractResponse) *string { return v.Provider }).(pulumi.StringPtrOutput)
}

type UserIdentityContractResponseArrayOutput struct{ *pulumi.OutputState }

func (UserIdentityContractResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserIdentityContractResponse)(nil)).Elem()
}

func (o UserIdentityContractResponseArrayOutput) ToUserIdentityContractResponseArrayOutput() UserIdentityContractResponseArrayOutput {
	return o
}

func (o UserIdentityContractResponseArrayOutput) ToUserIdentityContractResponseArrayOutputWithContext(ctx context.Context) UserIdentityContractResponseArrayOutput {
	return o
}

func (o UserIdentityContractResponseArrayOutput) Index(i pulumi.IntInput) UserIdentityContractResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserIdentityContractResponse {
		return vs[0].([]UserIdentityContractResponse)[vs[1].(int)]
	}).(UserIdentityContractResponseOutput)
}

type UserIdentityProperties struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserIdentityPropertiesInput interface {
	pulumi.Input

	ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput
	ToUserIdentityPropertiesOutputWithContext(context.Context) UserIdentityPropertiesOutput
}

type UserIdentityPropertiesArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserIdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityProperties)(nil)).Elem()
}

func (i UserIdentityPropertiesArgs) ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput {
	return i.ToUserIdentityPropertiesOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesArgs) ToUserIdentityPropertiesOutputWithContext(ctx context.Context) UserIdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesOutput)
}





type UserIdentityPropertiesMapInput interface {
	pulumi.Input

	ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput
	ToUserIdentityPropertiesMapOutputWithContext(context.Context) UserIdentityPropertiesMapOutput
}

type UserIdentityPropertiesMap map[string]UserIdentityPropertiesInput

func (UserIdentityPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityProperties)(nil)).Elem()
}

func (i UserIdentityPropertiesMap) ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput {
	return i.ToUserIdentityPropertiesMapOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesMap) ToUserIdentityPropertiesMapOutputWithContext(ctx context.Context) UserIdentityPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesMapOutput)
}

type UserIdentityPropertiesOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityProperties)(nil)).Elem()
}

func (o UserIdentityPropertiesOutput) ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput {
	return o
}

func (o UserIdentityPropertiesOutput) ToUserIdentityPropertiesOutputWithContext(ctx context.Context) UserIdentityPropertiesOutput {
	return o
}

func (o UserIdentityPropertiesOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityProperties) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityPropertiesOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityProperties) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserIdentityPropertiesMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityProperties)(nil)).Elem()
}

func (o UserIdentityPropertiesMapOutput) ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput {
	return o
}

func (o UserIdentityPropertiesMapOutput) ToUserIdentityPropertiesMapOutputWithContext(ctx context.Context) UserIdentityPropertiesMapOutput {
	return o
}

func (o UserIdentityPropertiesMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityProperties {
		return vs[0].(map[string]UserIdentityProperties)[vs[1].(string)]
	}).(UserIdentityPropertiesOutput)
}

type UserIdentityPropertiesResponse struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}

type UserIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityPropertiesResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserIdentityPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityPropertiesResponse {
		return vs[0].(map[string]UserIdentityPropertiesResponse)[vs[1].(string)]
	}).(UserIdentityPropertiesResponseOutput)
}

type VirtualNetworkConfiguration struct {
	SubnetResourceId *string `pulumi:"subnetResourceId"`
}





type VirtualNetworkConfigurationInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationOutput() VirtualNetworkConfigurationOutput
	ToVirtualNetworkConfigurationOutputWithContext(context.Context) VirtualNetworkConfigurationOutput
}

type VirtualNetworkConfigurationArgs struct {
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

func (o VirtualNetworkConfigurationPtrOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SubnetResourceId
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigurationResponse struct {
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	Subnetname       string  `pulumi:"subnetname"`
	Vnetid           string  `pulumi:"vnetid"`
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

type X509CertificateName struct {
	IssuerCertificateThumbprint *string `pulumi:"issuerCertificateThumbprint"`
	Name                        *string `pulumi:"name"`
}





type X509CertificateNameInput interface {
	pulumi.Input

	ToX509CertificateNameOutput() X509CertificateNameOutput
	ToX509CertificateNameOutputWithContext(context.Context) X509CertificateNameOutput
}

type X509CertificateNameArgs struct {
	IssuerCertificateThumbprint pulumi.StringPtrInput `pulumi:"issuerCertificateThumbprint"`
	Name                        pulumi.StringPtrInput `pulumi:"name"`
}

func (X509CertificateNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X509CertificateName)(nil)).Elem()
}

func (i X509CertificateNameArgs) ToX509CertificateNameOutput() X509CertificateNameOutput {
	return i.ToX509CertificateNameOutputWithContext(context.Background())
}

func (i X509CertificateNameArgs) ToX509CertificateNameOutputWithContext(ctx context.Context) X509CertificateNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X509CertificateNameOutput)
}





type X509CertificateNameArrayInput interface {
	pulumi.Input

	ToX509CertificateNameArrayOutput() X509CertificateNameArrayOutput
	ToX509CertificateNameArrayOutputWithContext(context.Context) X509CertificateNameArrayOutput
}

type X509CertificateNameArray []X509CertificateNameInput

func (X509CertificateNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X509CertificateName)(nil)).Elem()
}

func (i X509CertificateNameArray) ToX509CertificateNameArrayOutput() X509CertificateNameArrayOutput {
	return i.ToX509CertificateNameArrayOutputWithContext(context.Background())
}

func (i X509CertificateNameArray) ToX509CertificateNameArrayOutputWithContext(ctx context.Context) X509CertificateNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X509CertificateNameArrayOutput)
}

type X509CertificateNameOutput struct{ *pulumi.OutputState }

func (X509CertificateNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X509CertificateName)(nil)).Elem()
}

func (o X509CertificateNameOutput) ToX509CertificateNameOutput() X509CertificateNameOutput {
	return o
}

func (o X509CertificateNameOutput) ToX509CertificateNameOutputWithContext(ctx context.Context) X509CertificateNameOutput {
	return o
}

func (o X509CertificateNameOutput) IssuerCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X509CertificateName) *string { return v.IssuerCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o X509CertificateNameOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X509CertificateName) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type X509CertificateNameArrayOutput struct{ *pulumi.OutputState }

func (X509CertificateNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X509CertificateName)(nil)).Elem()
}

func (o X509CertificateNameArrayOutput) ToX509CertificateNameArrayOutput() X509CertificateNameArrayOutput {
	return o
}

func (o X509CertificateNameArrayOutput) ToX509CertificateNameArrayOutputWithContext(ctx context.Context) X509CertificateNameArrayOutput {
	return o
}

func (o X509CertificateNameArrayOutput) Index(i pulumi.IntInput) X509CertificateNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X509CertificateName {
		return vs[0].([]X509CertificateName)[vs[1].(int)]
	}).(X509CertificateNameOutput)
}

type X509CertificateNameResponse struct {
	IssuerCertificateThumbprint *string `pulumi:"issuerCertificateThumbprint"`
	Name                        *string `pulumi:"name"`
}

type X509CertificateNameResponseOutput struct{ *pulumi.OutputState }

func (X509CertificateNameResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X509CertificateNameResponse)(nil)).Elem()
}

func (o X509CertificateNameResponseOutput) ToX509CertificateNameResponseOutput() X509CertificateNameResponseOutput {
	return o
}

func (o X509CertificateNameResponseOutput) ToX509CertificateNameResponseOutputWithContext(ctx context.Context) X509CertificateNameResponseOutput {
	return o
}

func (o X509CertificateNameResponseOutput) IssuerCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X509CertificateNameResponse) *string { return v.IssuerCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o X509CertificateNameResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X509CertificateNameResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type X509CertificateNameResponseArrayOutput struct{ *pulumi.OutputState }

func (X509CertificateNameResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X509CertificateNameResponse)(nil)).Elem()
}

func (o X509CertificateNameResponseArrayOutput) ToX509CertificateNameResponseArrayOutput() X509CertificateNameResponseArrayOutput {
	return o
}

func (o X509CertificateNameResponseArrayOutput) ToX509CertificateNameResponseArrayOutputWithContext(ctx context.Context) X509CertificateNameResponseArrayOutput {
	return o
}

func (o X509CertificateNameResponseArrayOutput) Index(i pulumi.IntInput) X509CertificateNameResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X509CertificateNameResponse {
		return vs[0].([]X509CertificateNameResponse)[vs[1].(int)]
	}).(X509CertificateNameResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AdditionalLocationOutput{})
	pulumi.RegisterOutputType(AdditionalLocationArrayOutput{})
	pulumi.RegisterOutputType(AdditionalLocationResponseOutput{})
	pulumi.RegisterOutputType(AdditionalLocationResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiContactInformationOutput{})
	pulumi.RegisterOutputType(ApiContactInformationPtrOutput{})
	pulumi.RegisterOutputType(ApiContactInformationResponseOutput{})
	pulumi.RegisterOutputType(ApiContactInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiCreateOrUpdatePropertiesWsdlSelectorOutput{})
	pulumi.RegisterOutputType(ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput{})
	pulumi.RegisterOutputType(ApiLicenseInformationOutput{})
	pulumi.RegisterOutputType(ApiLicenseInformationPtrOutput{})
	pulumi.RegisterOutputType(ApiLicenseInformationResponseOutput{})
	pulumi.RegisterOutputType(ApiLicenseInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceIdentityOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApiVersionConstraintOutput{})
	pulumi.RegisterOutputType(ApiVersionConstraintPtrOutput{})
	pulumi.RegisterOutputType(ApiVersionConstraintResponseOutput{})
	pulumi.RegisterOutputType(ApiVersionConstraintResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiVersionSetContractDetailsOutput{})
	pulumi.RegisterOutputType(ApiVersionSetContractDetailsPtrOutput{})
	pulumi.RegisterOutputType(ApiVersionSetContractDetailsResponseOutput{})
	pulumi.RegisterOutputType(ApiVersionSetContractDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmIdWrapperResponseOutput{})
	pulumi.RegisterOutputType(ArmIdWrapperResponsePtrOutput{})
	pulumi.RegisterOutputType(AuthenticationSettingsContractOutput{})
	pulumi.RegisterOutputType(AuthenticationSettingsContractPtrOutput{})
	pulumi.RegisterOutputType(AuthenticationSettingsContractResponseOutput{})
	pulumi.RegisterOutputType(AuthenticationSettingsContractResponsePtrOutput{})
	pulumi.RegisterOutputType(BackendAuthorizationHeaderCredentialsOutput{})
	pulumi.RegisterOutputType(BackendAuthorizationHeaderCredentialsPtrOutput{})
	pulumi.RegisterOutputType(BackendAuthorizationHeaderCredentialsResponseOutput{})
	pulumi.RegisterOutputType(BackendAuthorizationHeaderCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(BackendCredentialsContractOutput{})
	pulumi.RegisterOutputType(BackendCredentialsContractPtrOutput{})
	pulumi.RegisterOutputType(BackendCredentialsContractResponseOutput{})
	pulumi.RegisterOutputType(BackendCredentialsContractResponsePtrOutput{})
	pulumi.RegisterOutputType(BackendPropertiesOutput{})
	pulumi.RegisterOutputType(BackendPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BackendPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BackendProxyContractOutput{})
	pulumi.RegisterOutputType(BackendProxyContractPtrOutput{})
	pulumi.RegisterOutputType(BackendProxyContractResponseOutput{})
	pulumi.RegisterOutputType(BackendProxyContractResponsePtrOutput{})
	pulumi.RegisterOutputType(BackendServiceFabricClusterPropertiesOutput{})
	pulumi.RegisterOutputType(BackendServiceFabricClusterPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BackendServiceFabricClusterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BackendServiceFabricClusterPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(BackendTlsPropertiesOutput{})
	pulumi.RegisterOutputType(BackendTlsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BackendTlsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BackendTlsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(BodyDiagnosticSettingsOutput{})
	pulumi.RegisterOutputType(BodyDiagnosticSettingsPtrOutput{})
	pulumi.RegisterOutputType(BodyDiagnosticSettingsResponseOutput{})
	pulumi.RegisterOutputType(BodyDiagnosticSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CertificateConfigurationOutput{})
	pulumi.RegisterOutputType(CertificateConfigurationArrayOutput{})
	pulumi.RegisterOutputType(CertificateConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CertificateConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(CertificateInformationOutput{})
	pulumi.RegisterOutputType(CertificateInformationPtrOutput{})
	pulumi.RegisterOutputType(CertificateInformationResponseOutput{})
	pulumi.RegisterOutputType(CertificateInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(DataMaskingOutput{})
	pulumi.RegisterOutputType(DataMaskingPtrOutput{})
	pulumi.RegisterOutputType(DataMaskingEntityOutput{})
	pulumi.RegisterOutputType(DataMaskingEntityArrayOutput{})
	pulumi.RegisterOutputType(DataMaskingEntityResponseOutput{})
	pulumi.RegisterOutputType(DataMaskingEntityResponseArrayOutput{})
	pulumi.RegisterOutputType(DataMaskingResponseOutput{})
	pulumi.RegisterOutputType(DataMaskingResponsePtrOutput{})
	pulumi.RegisterOutputType(EmailTemplateParametersContractPropertiesOutput{})
	pulumi.RegisterOutputType(EmailTemplateParametersContractPropertiesArrayOutput{})
	pulumi.RegisterOutputType(EmailTemplateParametersContractPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EmailTemplateParametersContractPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(GroupContractPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GroupContractPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationArrayOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationResponseOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(HttpMessageDiagnosticOutput{})
	pulumi.RegisterOutputType(HttpMessageDiagnosticPtrOutput{})
	pulumi.RegisterOutputType(HttpMessageDiagnosticResponseOutput{})
	pulumi.RegisterOutputType(HttpMessageDiagnosticResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultContractCreatePropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultContractCreatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultContractPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultContractPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultLastAccessStatusContractPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultLastAccessStatusContractPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(OAuth2AuthenticationSettingsContractOutput{})
	pulumi.RegisterOutputType(OAuth2AuthenticationSettingsContractPtrOutput{})
	pulumi.RegisterOutputType(OAuth2AuthenticationSettingsContractResponseOutput{})
	pulumi.RegisterOutputType(OAuth2AuthenticationSettingsContractResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenIdAuthenticationSettingsContractOutput{})
	pulumi.RegisterOutputType(OpenIdAuthenticationSettingsContractPtrOutput{})
	pulumi.RegisterOutputType(OpenIdAuthenticationSettingsContractResponseOutput{})
	pulumi.RegisterOutputType(OpenIdAuthenticationSettingsContractResponsePtrOutput{})
	pulumi.RegisterOutputType(ParameterContractOutput{})
	pulumi.RegisterOutputType(ParameterContractArrayOutput{})
	pulumi.RegisterOutputType(ParameterContractResponseOutput{})
	pulumi.RegisterOutputType(ParameterContractResponseArrayOutput{})
	pulumi.RegisterOutputType(ParameterExampleContractOutput{})
	pulumi.RegisterOutputType(ParameterExampleContractMapOutput{})
	pulumi.RegisterOutputType(ParameterExampleContractResponseOutput{})
	pulumi.RegisterOutputType(ParameterExampleContractResponseMapOutput{})
	pulumi.RegisterOutputType(PipelineDiagnosticSettingsOutput{})
	pulumi.RegisterOutputType(PipelineDiagnosticSettingsPtrOutput{})
	pulumi.RegisterOutputType(PipelineDiagnosticSettingsResponseOutput{})
	pulumi.RegisterOutputType(PipelineDiagnosticSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionRequestPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionWrapperOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionWrapperArrayOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionWrapperResponseOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionWrapperResponseArrayOutput{})
	pulumi.RegisterOutputType(RepresentationContractOutput{})
	pulumi.RegisterOutputType(RepresentationContractArrayOutput{})
	pulumi.RegisterOutputType(RepresentationContractResponseOutput{})
	pulumi.RegisterOutputType(RepresentationContractResponseArrayOutput{})
	pulumi.RegisterOutputType(RequestContractOutput{})
	pulumi.RegisterOutputType(RequestContractPtrOutput{})
	pulumi.RegisterOutputType(RequestContractResponseOutput{})
	pulumi.RegisterOutputType(RequestContractResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceLocationDataContractOutput{})
	pulumi.RegisterOutputType(ResourceLocationDataContractPtrOutput{})
	pulumi.RegisterOutputType(ResourceLocationDataContractResponseOutput{})
	pulumi.RegisterOutputType(ResourceLocationDataContractResponsePtrOutput{})
	pulumi.RegisterOutputType(ResponseContractOutput{})
	pulumi.RegisterOutputType(ResponseContractArrayOutput{})
	pulumi.RegisterOutputType(ResponseContractResponseOutput{})
	pulumi.RegisterOutputType(ResponseContractResponseArrayOutput{})
	pulumi.RegisterOutputType(SamplingSettingsOutput{})
	pulumi.RegisterOutputType(SamplingSettingsPtrOutput{})
	pulumi.RegisterOutputType(SamplingSettingsResponseOutput{})
	pulumi.RegisterOutputType(SamplingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SubscriptionKeyParameterNamesContractOutput{})
	pulumi.RegisterOutputType(SubscriptionKeyParameterNamesContractPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionKeyParameterNamesContractResponseOutput{})
	pulumi.RegisterOutputType(SubscriptionKeyParameterNamesContractResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TokenBodyParameterContractOutput{})
	pulumi.RegisterOutputType(TokenBodyParameterContractArrayOutput{})
	pulumi.RegisterOutputType(TokenBodyParameterContractResponseOutput{})
	pulumi.RegisterOutputType(TokenBodyParameterContractResponseArrayOutput{})
	pulumi.RegisterOutputType(UserIdentityContractOutput{})
	pulumi.RegisterOutputType(UserIdentityContractArrayOutput{})
	pulumi.RegisterOutputType(UserIdentityContractResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityContractResponseArrayOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesMapOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(X509CertificateNameOutput{})
	pulumi.RegisterOutputType(X509CertificateNameArrayOutput{})
	pulumi.RegisterOutputType(X509CertificateNameResponseOutput{})
	pulumi.RegisterOutputType(X509CertificateNameResponseArrayOutput{})
}
