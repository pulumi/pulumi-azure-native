


package v20170301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdditionalLocation struct {
	Location                    string                            `pulumi:"location"`
	Sku                         ApiManagementServiceSkuProperties `pulumi:"sku"`
	VirtualNetworkConfiguration *VirtualNetworkConfiguration      `pulumi:"virtualNetworkConfiguration"`
}


func (val *AdditionalLocation) Defaults() *AdditionalLocation {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = *tmp.Sku.Defaults()

	return &tmp
}





type AdditionalLocationInput interface {
	pulumi.Input

	ToAdditionalLocationOutput() AdditionalLocationOutput
	ToAdditionalLocationOutputWithContext(context.Context) AdditionalLocationOutput
}

type AdditionalLocationArgs struct {
	Location                    pulumi.StringInput                     `pulumi:"location"`
	Sku                         ApiManagementServiceSkuPropertiesInput `pulumi:"sku"`
	VirtualNetworkConfiguration VirtualNetworkConfigurationPtrInput    `pulumi:"virtualNetworkConfiguration"`
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

func (o AdditionalLocationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalLocation) string { return v.Location }).(pulumi.StringOutput)
}

func (o AdditionalLocationOutput) Sku() ApiManagementServiceSkuPropertiesOutput {
	return o.ApplyT(func(v AdditionalLocation) ApiManagementServiceSkuProperties { return v.Sku }).(ApiManagementServiceSkuPropertiesOutput)
}

func (o AdditionalLocationOutput) VirtualNetworkConfiguration() VirtualNetworkConfigurationPtrOutput {
	return o.ApplyT(func(v AdditionalLocation) *VirtualNetworkConfiguration { return v.VirtualNetworkConfiguration }).(VirtualNetworkConfigurationPtrOutput)
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
	GatewayRegionalUrl          string                                    `pulumi:"gatewayRegionalUrl"`
	Location                    string                                    `pulumi:"location"`
	Sku                         ApiManagementServiceSkuPropertiesResponse `pulumi:"sku"`
	StaticIps                   []string                                  `pulumi:"staticIps"`
	VirtualNetworkConfiguration *VirtualNetworkConfigurationResponse      `pulumi:"virtualNetworkConfiguration"`
}


func (val *AdditionalLocationResponse) Defaults() *AdditionalLocationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = *tmp.Sku.Defaults()

	return &tmp
}





type AdditionalLocationResponseInput interface {
	pulumi.Input

	ToAdditionalLocationResponseOutput() AdditionalLocationResponseOutput
	ToAdditionalLocationResponseOutputWithContext(context.Context) AdditionalLocationResponseOutput
}

type AdditionalLocationResponseArgs struct {
	GatewayRegionalUrl          pulumi.StringInput                             `pulumi:"gatewayRegionalUrl"`
	Location                    pulumi.StringInput                             `pulumi:"location"`
	Sku                         ApiManagementServiceSkuPropertiesResponseInput `pulumi:"sku"`
	StaticIps                   pulumi.StringArrayInput                        `pulumi:"staticIps"`
	VirtualNetworkConfiguration VirtualNetworkConfigurationResponsePtrInput    `pulumi:"virtualNetworkConfiguration"`
}

func (AdditionalLocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalLocationResponse)(nil)).Elem()
}

func (i AdditionalLocationResponseArgs) ToAdditionalLocationResponseOutput() AdditionalLocationResponseOutput {
	return i.ToAdditionalLocationResponseOutputWithContext(context.Background())
}

func (i AdditionalLocationResponseArgs) ToAdditionalLocationResponseOutputWithContext(ctx context.Context) AdditionalLocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalLocationResponseOutput)
}





type AdditionalLocationResponseArrayInput interface {
	pulumi.Input

	ToAdditionalLocationResponseArrayOutput() AdditionalLocationResponseArrayOutput
	ToAdditionalLocationResponseArrayOutputWithContext(context.Context) AdditionalLocationResponseArrayOutput
}

type AdditionalLocationResponseArray []AdditionalLocationResponseInput

func (AdditionalLocationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalLocationResponse)(nil)).Elem()
}

func (i AdditionalLocationResponseArray) ToAdditionalLocationResponseArrayOutput() AdditionalLocationResponseArrayOutput {
	return i.ToAdditionalLocationResponseArrayOutputWithContext(context.Background())
}

func (i AdditionalLocationResponseArray) ToAdditionalLocationResponseArrayOutputWithContext(ctx context.Context) AdditionalLocationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalLocationResponseArrayOutput)
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

func (o AdditionalLocationResponseOutput) GatewayRegionalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) string { return v.GatewayRegionalUrl }).(pulumi.StringOutput)
}

func (o AdditionalLocationResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o AdditionalLocationResponseOutput) Sku() ApiManagementServiceSkuPropertiesResponseOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) ApiManagementServiceSkuPropertiesResponse { return v.Sku }).(ApiManagementServiceSkuPropertiesResponseOutput)
}

func (o AdditionalLocationResponseOutput) StaticIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) []string { return v.StaticIps }).(pulumi.StringArrayOutput)
}

func (o AdditionalLocationResponseOutput) VirtualNetworkConfiguration() VirtualNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AdditionalLocationResponse) *VirtualNetworkConfigurationResponse {
		return v.VirtualNetworkConfiguration
	}).(VirtualNetworkConfigurationResponsePtrOutput)
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

type ApiManagementServiceIdentity struct {
	Type string `pulumi:"type"`
}





type ApiManagementServiceIdentityInput interface {
	pulumi.Input

	ToApiManagementServiceIdentityOutput() ApiManagementServiceIdentityOutput
	ToApiManagementServiceIdentityOutputWithContext(context.Context) ApiManagementServiceIdentityOutput
}

type ApiManagementServiceIdentityArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
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

type ApiManagementServiceIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}





type ApiManagementServiceIdentityResponseInput interface {
	pulumi.Input

	ToApiManagementServiceIdentityResponseOutput() ApiManagementServiceIdentityResponseOutput
	ToApiManagementServiceIdentityResponseOutputWithContext(context.Context) ApiManagementServiceIdentityResponseOutput
}

type ApiManagementServiceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (ApiManagementServiceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementServiceIdentityResponse)(nil)).Elem()
}

func (i ApiManagementServiceIdentityResponseArgs) ToApiManagementServiceIdentityResponseOutput() ApiManagementServiceIdentityResponseOutput {
	return i.ToApiManagementServiceIdentityResponseOutputWithContext(context.Background())
}

func (i ApiManagementServiceIdentityResponseArgs) ToApiManagementServiceIdentityResponseOutputWithContext(ctx context.Context) ApiManagementServiceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceIdentityResponseOutput)
}

func (i ApiManagementServiceIdentityResponseArgs) ToApiManagementServiceIdentityResponsePtrOutput() ApiManagementServiceIdentityResponsePtrOutput {
	return i.ToApiManagementServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ApiManagementServiceIdentityResponseArgs) ToApiManagementServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ApiManagementServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceIdentityResponseOutput).ToApiManagementServiceIdentityResponsePtrOutputWithContext(ctx)
}









type ApiManagementServiceIdentityResponsePtrInput interface {
	pulumi.Input

	ToApiManagementServiceIdentityResponsePtrOutput() ApiManagementServiceIdentityResponsePtrOutput
	ToApiManagementServiceIdentityResponsePtrOutputWithContext(context.Context) ApiManagementServiceIdentityResponsePtrOutput
}

type apiManagementServiceIdentityResponsePtrType ApiManagementServiceIdentityResponseArgs

func ApiManagementServiceIdentityResponsePtr(v *ApiManagementServiceIdentityResponseArgs) ApiManagementServiceIdentityResponsePtrInput {
	return (*apiManagementServiceIdentityResponsePtrType)(v)
}

func (*apiManagementServiceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementServiceIdentityResponse)(nil)).Elem()
}

func (i *apiManagementServiceIdentityResponsePtrType) ToApiManagementServiceIdentityResponsePtrOutput() ApiManagementServiceIdentityResponsePtrOutput {
	return i.ToApiManagementServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *apiManagementServiceIdentityResponsePtrType) ToApiManagementServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ApiManagementServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementServiceIdentityResponsePtrOutput)
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

func (o ApiManagementServiceIdentityResponseOutput) ToApiManagementServiceIdentityResponsePtrOutput() ApiManagementServiceIdentityResponsePtrOutput {
	return o.ToApiManagementServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ApiManagementServiceIdentityResponseOutput) ToApiManagementServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ApiManagementServiceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiManagementServiceIdentityResponse) *ApiManagementServiceIdentityResponse {
		return &v
	}).(ApiManagementServiceIdentityResponsePtrOutput)
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

type ApiManagementServiceSkuProperties struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}


func (val *ApiManagementServiceSkuProperties) Defaults() *ApiManagementServiceSkuProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Capacity) {
		capacity_ := 1
		tmp.Capacity = &capacity_
	}
	return &tmp
}





type ApiManagementServiceSkuPropertiesInput interface {
	pulumi.Input

	ToApiManagementServiceSkuPropertiesOutput() ApiManagementServiceSkuPropertiesOutput
	ToApiManagementServiceSkuPropertiesOutputWithContext(context.Context) ApiManagementServiceSkuPropertiesOutput
}

type ApiManagementServiceSkuPropertiesArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
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

func (o ApiManagementServiceSkuPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApiManagementServiceSkuProperties) string { return v.Name }).(pulumi.StringOutput)
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

func (o ApiManagementServiceSkuPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementServiceSkuProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ApiManagementServiceSkuPropertiesResponse struct {
	Capacity *int   `pulumi:"capacity"`
	Name     string `pulumi:"name"`
}


func (val *ApiManagementServiceSkuPropertiesResponse) Defaults() *ApiManagementServiceSkuPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Capacity) {
		capacity_ := 1
		tmp.Capacity = &capacity_
	}
	return &tmp
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

type ApiVersionSetContract struct {
	Description       *string `pulumi:"description"`
	DisplayName       string  `pulumi:"displayName"`
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	VersionQueryName  *string `pulumi:"versionQueryName"`
	VersioningScheme  string  `pulumi:"versioningScheme"`
}





type ApiVersionSetContractInput interface {
	pulumi.Input

	ToApiVersionSetContractOutput() ApiVersionSetContractOutput
	ToApiVersionSetContractOutputWithContext(context.Context) ApiVersionSetContractOutput
}

type ApiVersionSetContractArgs struct {
	Description       pulumi.StringPtrInput `pulumi:"description"`
	DisplayName       pulumi.StringInput    `pulumi:"displayName"`
	VersionHeaderName pulumi.StringPtrInput `pulumi:"versionHeaderName"`
	VersionQueryName  pulumi.StringPtrInput `pulumi:"versionQueryName"`
	VersioningScheme  pulumi.StringInput    `pulumi:"versioningScheme"`
}

func (ApiVersionSetContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSetContract)(nil)).Elem()
}

func (i ApiVersionSetContractArgs) ToApiVersionSetContractOutput() ApiVersionSetContractOutput {
	return i.ToApiVersionSetContractOutputWithContext(context.Background())
}

func (i ApiVersionSetContractArgs) ToApiVersionSetContractOutputWithContext(ctx context.Context) ApiVersionSetContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetContractOutput)
}

func (i ApiVersionSetContractArgs) ToApiVersionSetContractPtrOutput() ApiVersionSetContractPtrOutput {
	return i.ToApiVersionSetContractPtrOutputWithContext(context.Background())
}

func (i ApiVersionSetContractArgs) ToApiVersionSetContractPtrOutputWithContext(ctx context.Context) ApiVersionSetContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetContractOutput).ToApiVersionSetContractPtrOutputWithContext(ctx)
}









type ApiVersionSetContractPtrInput interface {
	pulumi.Input

	ToApiVersionSetContractPtrOutput() ApiVersionSetContractPtrOutput
	ToApiVersionSetContractPtrOutputWithContext(context.Context) ApiVersionSetContractPtrOutput
}

type apiVersionSetContractPtrType ApiVersionSetContractArgs

func ApiVersionSetContractPtr(v *ApiVersionSetContractArgs) ApiVersionSetContractPtrInput {
	return (*apiVersionSetContractPtrType)(v)
}

func (*apiVersionSetContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSetContract)(nil)).Elem()
}

func (i *apiVersionSetContractPtrType) ToApiVersionSetContractPtrOutput() ApiVersionSetContractPtrOutput {
	return i.ToApiVersionSetContractPtrOutputWithContext(context.Background())
}

func (i *apiVersionSetContractPtrType) ToApiVersionSetContractPtrOutputWithContext(ctx context.Context) ApiVersionSetContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetContractPtrOutput)
}

type ApiVersionSetContractOutput struct{ *pulumi.OutputState }

func (ApiVersionSetContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSetContract)(nil)).Elem()
}

func (o ApiVersionSetContractOutput) ToApiVersionSetContractOutput() ApiVersionSetContractOutput {
	return o
}

func (o ApiVersionSetContractOutput) ToApiVersionSetContractOutputWithContext(ctx context.Context) ApiVersionSetContractOutput {
	return o
}

func (o ApiVersionSetContractOutput) ToApiVersionSetContractPtrOutput() ApiVersionSetContractPtrOutput {
	return o.ToApiVersionSetContractPtrOutputWithContext(context.Background())
}

func (o ApiVersionSetContractOutput) ToApiVersionSetContractPtrOutputWithContext(ctx context.Context) ApiVersionSetContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiVersionSetContract) *ApiVersionSetContract {
		return &v
	}).(ApiVersionSetContractPtrOutput)
}

func (o ApiVersionSetContractOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContract) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ApiVersionSetContract) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ApiVersionSetContractOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContract) *string { return v.VersionHeaderName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContract) *string { return v.VersionQueryName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractOutput) VersioningScheme() pulumi.StringOutput {
	return o.ApplyT(func(v ApiVersionSetContract) string { return v.VersioningScheme }).(pulumi.StringOutput)
}

type ApiVersionSetContractPtrOutput struct{ *pulumi.OutputState }

func (ApiVersionSetContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSetContract)(nil)).Elem()
}

func (o ApiVersionSetContractPtrOutput) ToApiVersionSetContractPtrOutput() ApiVersionSetContractPtrOutput {
	return o
}

func (o ApiVersionSetContractPtrOutput) ToApiVersionSetContractPtrOutputWithContext(ctx context.Context) ApiVersionSetContractPtrOutput {
	return o
}

func (o ApiVersionSetContractPtrOutput) Elem() ApiVersionSetContractOutput {
	return o.ApplyT(func(v *ApiVersionSetContract) ApiVersionSetContract {
		if v != nil {
			return *v
		}
		var ret ApiVersionSetContract
		return ret
	}).(ApiVersionSetContractOutput)
}

func (o ApiVersionSetContractPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContract) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContract) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractPtrOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContract) *string {
		if v == nil {
			return nil
		}
		return v.VersionHeaderName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractPtrOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContract) *string {
		if v == nil {
			return nil
		}
		return v.VersionQueryName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractPtrOutput) VersioningScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContract) *string {
		if v == nil {
			return nil
		}
		return &v.VersioningScheme
	}).(pulumi.StringPtrOutput)
}

type ApiVersionSetContractResponse struct {
	Description       *string `pulumi:"description"`
	DisplayName       string  `pulumi:"displayName"`
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	Type              string  `pulumi:"type"`
	VersionHeaderName *string `pulumi:"versionHeaderName"`
	VersionQueryName  *string `pulumi:"versionQueryName"`
	VersioningScheme  string  `pulumi:"versioningScheme"`
}





type ApiVersionSetContractResponseInput interface {
	pulumi.Input

	ToApiVersionSetContractResponseOutput() ApiVersionSetContractResponseOutput
	ToApiVersionSetContractResponseOutputWithContext(context.Context) ApiVersionSetContractResponseOutput
}

type ApiVersionSetContractResponseArgs struct {
	Description       pulumi.StringPtrInput `pulumi:"description"`
	DisplayName       pulumi.StringInput    `pulumi:"displayName"`
	Id                pulumi.StringInput    `pulumi:"id"`
	Name              pulumi.StringInput    `pulumi:"name"`
	Type              pulumi.StringInput    `pulumi:"type"`
	VersionHeaderName pulumi.StringPtrInput `pulumi:"versionHeaderName"`
	VersionQueryName  pulumi.StringPtrInput `pulumi:"versionQueryName"`
	VersioningScheme  pulumi.StringInput    `pulumi:"versioningScheme"`
}

func (ApiVersionSetContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSetContractResponse)(nil)).Elem()
}

func (i ApiVersionSetContractResponseArgs) ToApiVersionSetContractResponseOutput() ApiVersionSetContractResponseOutput {
	return i.ToApiVersionSetContractResponseOutputWithContext(context.Background())
}

func (i ApiVersionSetContractResponseArgs) ToApiVersionSetContractResponseOutputWithContext(ctx context.Context) ApiVersionSetContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetContractResponseOutput)
}

func (i ApiVersionSetContractResponseArgs) ToApiVersionSetContractResponsePtrOutput() ApiVersionSetContractResponsePtrOutput {
	return i.ToApiVersionSetContractResponsePtrOutputWithContext(context.Background())
}

func (i ApiVersionSetContractResponseArgs) ToApiVersionSetContractResponsePtrOutputWithContext(ctx context.Context) ApiVersionSetContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetContractResponseOutput).ToApiVersionSetContractResponsePtrOutputWithContext(ctx)
}









type ApiVersionSetContractResponsePtrInput interface {
	pulumi.Input

	ToApiVersionSetContractResponsePtrOutput() ApiVersionSetContractResponsePtrOutput
	ToApiVersionSetContractResponsePtrOutputWithContext(context.Context) ApiVersionSetContractResponsePtrOutput
}

type apiVersionSetContractResponsePtrType ApiVersionSetContractResponseArgs

func ApiVersionSetContractResponsePtr(v *ApiVersionSetContractResponseArgs) ApiVersionSetContractResponsePtrInput {
	return (*apiVersionSetContractResponsePtrType)(v)
}

func (*apiVersionSetContractResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSetContractResponse)(nil)).Elem()
}

func (i *apiVersionSetContractResponsePtrType) ToApiVersionSetContractResponsePtrOutput() ApiVersionSetContractResponsePtrOutput {
	return i.ToApiVersionSetContractResponsePtrOutputWithContext(context.Background())
}

func (i *apiVersionSetContractResponsePtrType) ToApiVersionSetContractResponsePtrOutputWithContext(ctx context.Context) ApiVersionSetContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiVersionSetContractResponsePtrOutput)
}

type ApiVersionSetContractResponseOutput struct{ *pulumi.OutputState }

func (ApiVersionSetContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiVersionSetContractResponse)(nil)).Elem()
}

func (o ApiVersionSetContractResponseOutput) ToApiVersionSetContractResponseOutput() ApiVersionSetContractResponseOutput {
	return o
}

func (o ApiVersionSetContractResponseOutput) ToApiVersionSetContractResponseOutputWithContext(ctx context.Context) ApiVersionSetContractResponseOutput {
	return o
}

func (o ApiVersionSetContractResponseOutput) ToApiVersionSetContractResponsePtrOutput() ApiVersionSetContractResponsePtrOutput {
	return o.ToApiVersionSetContractResponsePtrOutputWithContext(context.Background())
}

func (o ApiVersionSetContractResponseOutput) ToApiVersionSetContractResponsePtrOutputWithContext(ctx context.Context) ApiVersionSetContractResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiVersionSetContractResponse) *ApiVersionSetContractResponse {
		return &v
	}).(ApiVersionSetContractResponsePtrOutput)
}

func (o ApiVersionSetContractResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ApiVersionSetContractResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ApiVersionSetContractResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ApiVersionSetContractResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ApiVersionSetContractResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApiVersionSetContractResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApiVersionSetContractResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ApiVersionSetContractResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ApiVersionSetContractResponseOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractResponse) *string { return v.VersionHeaderName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponseOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiVersionSetContractResponse) *string { return v.VersionQueryName }).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponseOutput) VersioningScheme() pulumi.StringOutput {
	return o.ApplyT(func(v ApiVersionSetContractResponse) string { return v.VersioningScheme }).(pulumi.StringOutput)
}

type ApiVersionSetContractResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiVersionSetContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiVersionSetContractResponse)(nil)).Elem()
}

func (o ApiVersionSetContractResponsePtrOutput) ToApiVersionSetContractResponsePtrOutput() ApiVersionSetContractResponsePtrOutput {
	return o
}

func (o ApiVersionSetContractResponsePtrOutput) ToApiVersionSetContractResponsePtrOutputWithContext(ctx context.Context) ApiVersionSetContractResponsePtrOutput {
	return o
}

func (o ApiVersionSetContractResponsePtrOutput) Elem() ApiVersionSetContractResponseOutput {
	return o.ApplyT(func(v *ApiVersionSetContractResponse) ApiVersionSetContractResponse {
		if v != nil {
			return *v
		}
		var ret ApiVersionSetContractResponse
		return ret
	}).(ApiVersionSetContractResponseOutput)
}

func (o ApiVersionSetContractResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponsePtrOutput) VersionHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.VersionHeaderName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponsePtrOutput) VersionQueryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.VersionQueryName
	}).(pulumi.StringPtrOutput)
}

func (o ApiVersionSetContractResponsePtrOutput) VersioningScheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiVersionSetContractResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VersioningScheme
	}).(pulumi.StringPtrOutput)
}

type AuthenticationSettingsContract struct {
	OAuth2 *OAuth2AuthenticationSettingsContract `pulumi:"oAuth2"`
}





type AuthenticationSettingsContractInput interface {
	pulumi.Input

	ToAuthenticationSettingsContractOutput() AuthenticationSettingsContractOutput
	ToAuthenticationSettingsContractOutputWithContext(context.Context) AuthenticationSettingsContractOutput
}

type AuthenticationSettingsContractArgs struct {
	OAuth2 OAuth2AuthenticationSettingsContractPtrInput `pulumi:"oAuth2"`
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

type AuthenticationSettingsContractResponse struct {
	OAuth2 *OAuth2AuthenticationSettingsContractResponse `pulumi:"oAuth2"`
}





type AuthenticationSettingsContractResponseInput interface {
	pulumi.Input

	ToAuthenticationSettingsContractResponseOutput() AuthenticationSettingsContractResponseOutput
	ToAuthenticationSettingsContractResponseOutputWithContext(context.Context) AuthenticationSettingsContractResponseOutput
}

type AuthenticationSettingsContractResponseArgs struct {
	OAuth2 OAuth2AuthenticationSettingsContractResponsePtrInput `pulumi:"oAuth2"`
}

func (AuthenticationSettingsContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationSettingsContractResponse)(nil)).Elem()
}

func (i AuthenticationSettingsContractResponseArgs) ToAuthenticationSettingsContractResponseOutput() AuthenticationSettingsContractResponseOutput {
	return i.ToAuthenticationSettingsContractResponseOutputWithContext(context.Background())
}

func (i AuthenticationSettingsContractResponseArgs) ToAuthenticationSettingsContractResponseOutputWithContext(ctx context.Context) AuthenticationSettingsContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationSettingsContractResponseOutput)
}

func (i AuthenticationSettingsContractResponseArgs) ToAuthenticationSettingsContractResponsePtrOutput() AuthenticationSettingsContractResponsePtrOutput {
	return i.ToAuthenticationSettingsContractResponsePtrOutputWithContext(context.Background())
}

func (i AuthenticationSettingsContractResponseArgs) ToAuthenticationSettingsContractResponsePtrOutputWithContext(ctx context.Context) AuthenticationSettingsContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationSettingsContractResponseOutput).ToAuthenticationSettingsContractResponsePtrOutputWithContext(ctx)
}









type AuthenticationSettingsContractResponsePtrInput interface {
	pulumi.Input

	ToAuthenticationSettingsContractResponsePtrOutput() AuthenticationSettingsContractResponsePtrOutput
	ToAuthenticationSettingsContractResponsePtrOutputWithContext(context.Context) AuthenticationSettingsContractResponsePtrOutput
}

type authenticationSettingsContractResponsePtrType AuthenticationSettingsContractResponseArgs

func AuthenticationSettingsContractResponsePtr(v *AuthenticationSettingsContractResponseArgs) AuthenticationSettingsContractResponsePtrInput {
	return (*authenticationSettingsContractResponsePtrType)(v)
}

func (*authenticationSettingsContractResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationSettingsContractResponse)(nil)).Elem()
}

func (i *authenticationSettingsContractResponsePtrType) ToAuthenticationSettingsContractResponsePtrOutput() AuthenticationSettingsContractResponsePtrOutput {
	return i.ToAuthenticationSettingsContractResponsePtrOutputWithContext(context.Background())
}

func (i *authenticationSettingsContractResponsePtrType) ToAuthenticationSettingsContractResponsePtrOutputWithContext(ctx context.Context) AuthenticationSettingsContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationSettingsContractResponsePtrOutput)
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

func (o AuthenticationSettingsContractResponseOutput) ToAuthenticationSettingsContractResponsePtrOutput() AuthenticationSettingsContractResponsePtrOutput {
	return o.ToAuthenticationSettingsContractResponsePtrOutputWithContext(context.Background())
}

func (o AuthenticationSettingsContractResponseOutput) ToAuthenticationSettingsContractResponsePtrOutputWithContext(ctx context.Context) AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationSettingsContractResponse) *AuthenticationSettingsContractResponse {
		return &v
	}).(AuthenticationSettingsContractResponsePtrOutput)
}

func (o AuthenticationSettingsContractResponseOutput) OAuth2() OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v AuthenticationSettingsContractResponse) *OAuth2AuthenticationSettingsContractResponse {
		return v.OAuth2
	}).(OAuth2AuthenticationSettingsContractResponsePtrOutput)
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





type BackendAuthorizationHeaderCredentialsResponseInput interface {
	pulumi.Input

	ToBackendAuthorizationHeaderCredentialsResponseOutput() BackendAuthorizationHeaderCredentialsResponseOutput
	ToBackendAuthorizationHeaderCredentialsResponseOutputWithContext(context.Context) BackendAuthorizationHeaderCredentialsResponseOutput
}

type BackendAuthorizationHeaderCredentialsResponseArgs struct {
	Parameter pulumi.StringInput `pulumi:"parameter"`
	Scheme    pulumi.StringInput `pulumi:"scheme"`
}

func (BackendAuthorizationHeaderCredentialsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAuthorizationHeaderCredentialsResponse)(nil)).Elem()
}

func (i BackendAuthorizationHeaderCredentialsResponseArgs) ToBackendAuthorizationHeaderCredentialsResponseOutput() BackendAuthorizationHeaderCredentialsResponseOutput {
	return i.ToBackendAuthorizationHeaderCredentialsResponseOutputWithContext(context.Background())
}

func (i BackendAuthorizationHeaderCredentialsResponseArgs) ToBackendAuthorizationHeaderCredentialsResponseOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAuthorizationHeaderCredentialsResponseOutput)
}

func (i BackendAuthorizationHeaderCredentialsResponseArgs) ToBackendAuthorizationHeaderCredentialsResponsePtrOutput() BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return i.ToBackendAuthorizationHeaderCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i BackendAuthorizationHeaderCredentialsResponseArgs) ToBackendAuthorizationHeaderCredentialsResponsePtrOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAuthorizationHeaderCredentialsResponseOutput).ToBackendAuthorizationHeaderCredentialsResponsePtrOutputWithContext(ctx)
}









type BackendAuthorizationHeaderCredentialsResponsePtrInput interface {
	pulumi.Input

	ToBackendAuthorizationHeaderCredentialsResponsePtrOutput() BackendAuthorizationHeaderCredentialsResponsePtrOutput
	ToBackendAuthorizationHeaderCredentialsResponsePtrOutputWithContext(context.Context) BackendAuthorizationHeaderCredentialsResponsePtrOutput
}

type backendAuthorizationHeaderCredentialsResponsePtrType BackendAuthorizationHeaderCredentialsResponseArgs

func BackendAuthorizationHeaderCredentialsResponsePtr(v *BackendAuthorizationHeaderCredentialsResponseArgs) BackendAuthorizationHeaderCredentialsResponsePtrInput {
	return (*backendAuthorizationHeaderCredentialsResponsePtrType)(v)
}

func (*backendAuthorizationHeaderCredentialsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendAuthorizationHeaderCredentialsResponse)(nil)).Elem()
}

func (i *backendAuthorizationHeaderCredentialsResponsePtrType) ToBackendAuthorizationHeaderCredentialsResponsePtrOutput() BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return i.ToBackendAuthorizationHeaderCredentialsResponsePtrOutputWithContext(context.Background())
}

func (i *backendAuthorizationHeaderCredentialsResponsePtrType) ToBackendAuthorizationHeaderCredentialsResponsePtrOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAuthorizationHeaderCredentialsResponsePtrOutput)
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

func (o BackendAuthorizationHeaderCredentialsResponseOutput) ToBackendAuthorizationHeaderCredentialsResponsePtrOutput() BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return o.ToBackendAuthorizationHeaderCredentialsResponsePtrOutputWithContext(context.Background())
}

func (o BackendAuthorizationHeaderCredentialsResponseOutput) ToBackendAuthorizationHeaderCredentialsResponsePtrOutputWithContext(ctx context.Context) BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendAuthorizationHeaderCredentialsResponse) *BackendAuthorizationHeaderCredentialsResponse {
		return &v
	}).(BackendAuthorizationHeaderCredentialsResponsePtrOutput)
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
	Authorization *BackendAuthorizationHeaderCredentials `pulumi:"authorization"`
	Certificate   []string                               `pulumi:"certificate"`
	Header        map[string][]string                    `pulumi:"header"`
	Query         map[string][]string                    `pulumi:"query"`
}





type BackendCredentialsContractInput interface {
	pulumi.Input

	ToBackendCredentialsContractOutput() BackendCredentialsContractOutput
	ToBackendCredentialsContractOutputWithContext(context.Context) BackendCredentialsContractOutput
}

type BackendCredentialsContractArgs struct {
	Authorization BackendAuthorizationHeaderCredentialsPtrInput `pulumi:"authorization"`
	Certificate   pulumi.StringArrayInput                       `pulumi:"certificate"`
	Header        pulumi.StringArrayMapInput                    `pulumi:"header"`
	Query         pulumi.StringArrayMapInput                    `pulumi:"query"`
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
	Authorization *BackendAuthorizationHeaderCredentialsResponse `pulumi:"authorization"`
	Certificate   []string                                       `pulumi:"certificate"`
	Header        map[string][]string                            `pulumi:"header"`
	Query         map[string][]string                            `pulumi:"query"`
}





type BackendCredentialsContractResponseInput interface {
	pulumi.Input

	ToBackendCredentialsContractResponseOutput() BackendCredentialsContractResponseOutput
	ToBackendCredentialsContractResponseOutputWithContext(context.Context) BackendCredentialsContractResponseOutput
}

type BackendCredentialsContractResponseArgs struct {
	Authorization BackendAuthorizationHeaderCredentialsResponsePtrInput `pulumi:"authorization"`
	Certificate   pulumi.StringArrayInput                               `pulumi:"certificate"`
	Header        pulumi.StringArrayMapInput                            `pulumi:"header"`
	Query         pulumi.StringArrayMapInput                            `pulumi:"query"`
}

func (BackendCredentialsContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendCredentialsContractResponse)(nil)).Elem()
}

func (i BackendCredentialsContractResponseArgs) ToBackendCredentialsContractResponseOutput() BackendCredentialsContractResponseOutput {
	return i.ToBackendCredentialsContractResponseOutputWithContext(context.Background())
}

func (i BackendCredentialsContractResponseArgs) ToBackendCredentialsContractResponseOutputWithContext(ctx context.Context) BackendCredentialsContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendCredentialsContractResponseOutput)
}

func (i BackendCredentialsContractResponseArgs) ToBackendCredentialsContractResponsePtrOutput() BackendCredentialsContractResponsePtrOutput {
	return i.ToBackendCredentialsContractResponsePtrOutputWithContext(context.Background())
}

func (i BackendCredentialsContractResponseArgs) ToBackendCredentialsContractResponsePtrOutputWithContext(ctx context.Context) BackendCredentialsContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendCredentialsContractResponseOutput).ToBackendCredentialsContractResponsePtrOutputWithContext(ctx)
}









type BackendCredentialsContractResponsePtrInput interface {
	pulumi.Input

	ToBackendCredentialsContractResponsePtrOutput() BackendCredentialsContractResponsePtrOutput
	ToBackendCredentialsContractResponsePtrOutputWithContext(context.Context) BackendCredentialsContractResponsePtrOutput
}

type backendCredentialsContractResponsePtrType BackendCredentialsContractResponseArgs

func BackendCredentialsContractResponsePtr(v *BackendCredentialsContractResponseArgs) BackendCredentialsContractResponsePtrInput {
	return (*backendCredentialsContractResponsePtrType)(v)
}

func (*backendCredentialsContractResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendCredentialsContractResponse)(nil)).Elem()
}

func (i *backendCredentialsContractResponsePtrType) ToBackendCredentialsContractResponsePtrOutput() BackendCredentialsContractResponsePtrOutput {
	return i.ToBackendCredentialsContractResponsePtrOutputWithContext(context.Background())
}

func (i *backendCredentialsContractResponsePtrType) ToBackendCredentialsContractResponsePtrOutputWithContext(ctx context.Context) BackendCredentialsContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendCredentialsContractResponsePtrOutput)
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

func (o BackendCredentialsContractResponseOutput) ToBackendCredentialsContractResponsePtrOutput() BackendCredentialsContractResponsePtrOutput {
	return o.ToBackendCredentialsContractResponsePtrOutputWithContext(context.Background())
}

func (o BackendCredentialsContractResponseOutput) ToBackendCredentialsContractResponsePtrOutputWithContext(ctx context.Context) BackendCredentialsContractResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendCredentialsContractResponse) *BackendCredentialsContractResponse {
		return &v
	}).(BackendCredentialsContractResponsePtrOutput)
}

func (o BackendCredentialsContractResponseOutput) Authorization() BackendAuthorizationHeaderCredentialsResponsePtrOutput {
	return o.ApplyT(func(v BackendCredentialsContractResponse) *BackendAuthorizationHeaderCredentialsResponse {
		return v.Authorization
	}).(BackendAuthorizationHeaderCredentialsResponsePtrOutput)
}

func (o BackendCredentialsContractResponseOutput) Certificate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BackendCredentialsContractResponse) []string { return v.Certificate }).(pulumi.StringArrayOutput)
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





type BackendPropertiesResponseInput interface {
	pulumi.Input

	ToBackendPropertiesResponseOutput() BackendPropertiesResponseOutput
	ToBackendPropertiesResponseOutputWithContext(context.Context) BackendPropertiesResponseOutput
}

type BackendPropertiesResponseArgs struct {
	ServiceFabricCluster BackendServiceFabricClusterPropertiesResponsePtrInput `pulumi:"serviceFabricCluster"`
}

func (BackendPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPropertiesResponse)(nil)).Elem()
}

func (i BackendPropertiesResponseArgs) ToBackendPropertiesResponseOutput() BackendPropertiesResponseOutput {
	return i.ToBackendPropertiesResponseOutputWithContext(context.Background())
}

func (i BackendPropertiesResponseArgs) ToBackendPropertiesResponseOutputWithContext(ctx context.Context) BackendPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPropertiesResponseOutput)
}

func (i BackendPropertiesResponseArgs) ToBackendPropertiesResponsePtrOutput() BackendPropertiesResponsePtrOutput {
	return i.ToBackendPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i BackendPropertiesResponseArgs) ToBackendPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPropertiesResponseOutput).ToBackendPropertiesResponsePtrOutputWithContext(ctx)
}









type BackendPropertiesResponsePtrInput interface {
	pulumi.Input

	ToBackendPropertiesResponsePtrOutput() BackendPropertiesResponsePtrOutput
	ToBackendPropertiesResponsePtrOutputWithContext(context.Context) BackendPropertiesResponsePtrOutput
}

type backendPropertiesResponsePtrType BackendPropertiesResponseArgs

func BackendPropertiesResponsePtr(v *BackendPropertiesResponseArgs) BackendPropertiesResponsePtrInput {
	return (*backendPropertiesResponsePtrType)(v)
}

func (*backendPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendPropertiesResponse)(nil)).Elem()
}

func (i *backendPropertiesResponsePtrType) ToBackendPropertiesResponsePtrOutput() BackendPropertiesResponsePtrOutput {
	return i.ToBackendPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *backendPropertiesResponsePtrType) ToBackendPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPropertiesResponsePtrOutput)
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

func (o BackendPropertiesResponseOutput) ToBackendPropertiesResponsePtrOutput() BackendPropertiesResponsePtrOutput {
	return o.ToBackendPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o BackendPropertiesResponseOutput) ToBackendPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendPropertiesResponse) *BackendPropertiesResponse {
		return &v
	}).(BackendPropertiesResponsePtrOutput)
}

func (o BackendPropertiesResponseOutput) ServiceFabricCluster() BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v BackendPropertiesResponse) *BackendServiceFabricClusterPropertiesResponse {
		return v.ServiceFabricCluster
	}).(BackendServiceFabricClusterPropertiesResponsePtrOutput)
}

type BackendPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (BackendPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendPropertiesResponse)(nil)).Elem()
}

func (o BackendPropertiesResponsePtrOutput) ToBackendPropertiesResponsePtrOutput() BackendPropertiesResponsePtrOutput {
	return o
}

func (o BackendPropertiesResponsePtrOutput) ToBackendPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendPropertiesResponsePtrOutput {
	return o
}

func (o BackendPropertiesResponsePtrOutput) Elem() BackendPropertiesResponseOutput {
	return o.ApplyT(func(v *BackendPropertiesResponse) BackendPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret BackendPropertiesResponse
		return ret
	}).(BackendPropertiesResponseOutput)
}

func (o BackendPropertiesResponsePtrOutput) ServiceFabricCluster() BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BackendPropertiesResponse) *BackendServiceFabricClusterPropertiesResponse {
		if v == nil {
			return nil
		}
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





type BackendProxyContractResponseInput interface {
	pulumi.Input

	ToBackendProxyContractResponseOutput() BackendProxyContractResponseOutput
	ToBackendProxyContractResponseOutputWithContext(context.Context) BackendProxyContractResponseOutput
}

type BackendProxyContractResponseArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Url      pulumi.StringInput    `pulumi:"url"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (BackendProxyContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendProxyContractResponse)(nil)).Elem()
}

func (i BackendProxyContractResponseArgs) ToBackendProxyContractResponseOutput() BackendProxyContractResponseOutput {
	return i.ToBackendProxyContractResponseOutputWithContext(context.Background())
}

func (i BackendProxyContractResponseArgs) ToBackendProxyContractResponseOutputWithContext(ctx context.Context) BackendProxyContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendProxyContractResponseOutput)
}

func (i BackendProxyContractResponseArgs) ToBackendProxyContractResponsePtrOutput() BackendProxyContractResponsePtrOutput {
	return i.ToBackendProxyContractResponsePtrOutputWithContext(context.Background())
}

func (i BackendProxyContractResponseArgs) ToBackendProxyContractResponsePtrOutputWithContext(ctx context.Context) BackendProxyContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendProxyContractResponseOutput).ToBackendProxyContractResponsePtrOutputWithContext(ctx)
}









type BackendProxyContractResponsePtrInput interface {
	pulumi.Input

	ToBackendProxyContractResponsePtrOutput() BackendProxyContractResponsePtrOutput
	ToBackendProxyContractResponsePtrOutputWithContext(context.Context) BackendProxyContractResponsePtrOutput
}

type backendProxyContractResponsePtrType BackendProxyContractResponseArgs

func BackendProxyContractResponsePtr(v *BackendProxyContractResponseArgs) BackendProxyContractResponsePtrInput {
	return (*backendProxyContractResponsePtrType)(v)
}

func (*backendProxyContractResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendProxyContractResponse)(nil)).Elem()
}

func (i *backendProxyContractResponsePtrType) ToBackendProxyContractResponsePtrOutput() BackendProxyContractResponsePtrOutput {
	return i.ToBackendProxyContractResponsePtrOutputWithContext(context.Background())
}

func (i *backendProxyContractResponsePtrType) ToBackendProxyContractResponsePtrOutputWithContext(ctx context.Context) BackendProxyContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendProxyContractResponsePtrOutput)
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

func (o BackendProxyContractResponseOutput) ToBackendProxyContractResponsePtrOutput() BackendProxyContractResponsePtrOutput {
	return o.ToBackendProxyContractResponsePtrOutputWithContext(context.Background())
}

func (o BackendProxyContractResponseOutput) ToBackendProxyContractResponsePtrOutputWithContext(ctx context.Context) BackendProxyContractResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendProxyContractResponse) *BackendProxyContractResponse {
		return &v
	}).(BackendProxyContractResponsePtrOutput)
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
	ClientCertificatethumbprint   string                `pulumi:"clientCertificatethumbprint"`
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
	ClientCertificatethumbprint   pulumi.StringInput            `pulumi:"clientCertificatethumbprint"`
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

func (o BackendServiceFabricClusterPropertiesOutput) ClientCertificatethumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterProperties) string { return v.ClientCertificatethumbprint }).(pulumi.StringOutput)
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

func (o BackendServiceFabricClusterPropertiesPtrOutput) ClientCertificatethumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ClientCertificatethumbprint
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
	ClientCertificatethumbprint   string                        `pulumi:"clientCertificatethumbprint"`
	ManagementEndpoints           []string                      `pulumi:"managementEndpoints"`
	MaxPartitionResolutionRetries *int                          `pulumi:"maxPartitionResolutionRetries"`
	ServerCertificateThumbprints  []string                      `pulumi:"serverCertificateThumbprints"`
	ServerX509Names               []X509CertificateNameResponse `pulumi:"serverX509Names"`
}





type BackendServiceFabricClusterPropertiesResponseInput interface {
	pulumi.Input

	ToBackendServiceFabricClusterPropertiesResponseOutput() BackendServiceFabricClusterPropertiesResponseOutput
	ToBackendServiceFabricClusterPropertiesResponseOutputWithContext(context.Context) BackendServiceFabricClusterPropertiesResponseOutput
}

type BackendServiceFabricClusterPropertiesResponseArgs struct {
	ClientCertificatethumbprint   pulumi.StringInput                    `pulumi:"clientCertificatethumbprint"`
	ManagementEndpoints           pulumi.StringArrayInput               `pulumi:"managementEndpoints"`
	MaxPartitionResolutionRetries pulumi.IntPtrInput                    `pulumi:"maxPartitionResolutionRetries"`
	ServerCertificateThumbprints  pulumi.StringArrayInput               `pulumi:"serverCertificateThumbprints"`
	ServerX509Names               X509CertificateNameResponseArrayInput `pulumi:"serverX509Names"`
}

func (BackendServiceFabricClusterPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendServiceFabricClusterPropertiesResponse)(nil)).Elem()
}

func (i BackendServiceFabricClusterPropertiesResponseArgs) ToBackendServiceFabricClusterPropertiesResponseOutput() BackendServiceFabricClusterPropertiesResponseOutput {
	return i.ToBackendServiceFabricClusterPropertiesResponseOutputWithContext(context.Background())
}

func (i BackendServiceFabricClusterPropertiesResponseArgs) ToBackendServiceFabricClusterPropertiesResponseOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceFabricClusterPropertiesResponseOutput)
}

func (i BackendServiceFabricClusterPropertiesResponseArgs) ToBackendServiceFabricClusterPropertiesResponsePtrOutput() BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return i.ToBackendServiceFabricClusterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i BackendServiceFabricClusterPropertiesResponseArgs) ToBackendServiceFabricClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceFabricClusterPropertiesResponseOutput).ToBackendServiceFabricClusterPropertiesResponsePtrOutputWithContext(ctx)
}









type BackendServiceFabricClusterPropertiesResponsePtrInput interface {
	pulumi.Input

	ToBackendServiceFabricClusterPropertiesResponsePtrOutput() BackendServiceFabricClusterPropertiesResponsePtrOutput
	ToBackendServiceFabricClusterPropertiesResponsePtrOutputWithContext(context.Context) BackendServiceFabricClusterPropertiesResponsePtrOutput
}

type backendServiceFabricClusterPropertiesResponsePtrType BackendServiceFabricClusterPropertiesResponseArgs

func BackendServiceFabricClusterPropertiesResponsePtr(v *BackendServiceFabricClusterPropertiesResponseArgs) BackendServiceFabricClusterPropertiesResponsePtrInput {
	return (*backendServiceFabricClusterPropertiesResponsePtrType)(v)
}

func (*backendServiceFabricClusterPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendServiceFabricClusterPropertiesResponse)(nil)).Elem()
}

func (i *backendServiceFabricClusterPropertiesResponsePtrType) ToBackendServiceFabricClusterPropertiesResponsePtrOutput() BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return i.ToBackendServiceFabricClusterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *backendServiceFabricClusterPropertiesResponsePtrType) ToBackendServiceFabricClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendServiceFabricClusterPropertiesResponsePtrOutput)
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

func (o BackendServiceFabricClusterPropertiesResponseOutput) ToBackendServiceFabricClusterPropertiesResponsePtrOutput() BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return o.ToBackendServiceFabricClusterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) ToBackendServiceFabricClusterPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendServiceFabricClusterPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendServiceFabricClusterPropertiesResponse) *BackendServiceFabricClusterPropertiesResponse {
		return &v
	}).(BackendServiceFabricClusterPropertiesResponsePtrOutput)
}

func (o BackendServiceFabricClusterPropertiesResponseOutput) ClientCertificatethumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v BackendServiceFabricClusterPropertiesResponse) string { return v.ClientCertificatethumbprint }).(pulumi.StringOutput)
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

func (o BackendServiceFabricClusterPropertiesResponsePtrOutput) ClientCertificatethumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendServiceFabricClusterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientCertificatethumbprint
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





type BackendTlsPropertiesResponseInput interface {
	pulumi.Input

	ToBackendTlsPropertiesResponseOutput() BackendTlsPropertiesResponseOutput
	ToBackendTlsPropertiesResponseOutputWithContext(context.Context) BackendTlsPropertiesResponseOutput
}

type BackendTlsPropertiesResponseArgs struct {
	ValidateCertificateChain pulumi.BoolPtrInput `pulumi:"validateCertificateChain"`
	ValidateCertificateName  pulumi.BoolPtrInput `pulumi:"validateCertificateName"`
}

func (BackendTlsPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendTlsPropertiesResponse)(nil)).Elem()
}

func (i BackendTlsPropertiesResponseArgs) ToBackendTlsPropertiesResponseOutput() BackendTlsPropertiesResponseOutput {
	return i.ToBackendTlsPropertiesResponseOutputWithContext(context.Background())
}

func (i BackendTlsPropertiesResponseArgs) ToBackendTlsPropertiesResponseOutputWithContext(ctx context.Context) BackendTlsPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendTlsPropertiesResponseOutput)
}

func (i BackendTlsPropertiesResponseArgs) ToBackendTlsPropertiesResponsePtrOutput() BackendTlsPropertiesResponsePtrOutput {
	return i.ToBackendTlsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i BackendTlsPropertiesResponseArgs) ToBackendTlsPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendTlsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendTlsPropertiesResponseOutput).ToBackendTlsPropertiesResponsePtrOutputWithContext(ctx)
}









type BackendTlsPropertiesResponsePtrInput interface {
	pulumi.Input

	ToBackendTlsPropertiesResponsePtrOutput() BackendTlsPropertiesResponsePtrOutput
	ToBackendTlsPropertiesResponsePtrOutputWithContext(context.Context) BackendTlsPropertiesResponsePtrOutput
}

type backendTlsPropertiesResponsePtrType BackendTlsPropertiesResponseArgs

func BackendTlsPropertiesResponsePtr(v *BackendTlsPropertiesResponseArgs) BackendTlsPropertiesResponsePtrInput {
	return (*backendTlsPropertiesResponsePtrType)(v)
}

func (*backendTlsPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendTlsPropertiesResponse)(nil)).Elem()
}

func (i *backendTlsPropertiesResponsePtrType) ToBackendTlsPropertiesResponsePtrOutput() BackendTlsPropertiesResponsePtrOutput {
	return i.ToBackendTlsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *backendTlsPropertiesResponsePtrType) ToBackendTlsPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendTlsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendTlsPropertiesResponsePtrOutput)
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

func (o BackendTlsPropertiesResponseOutput) ToBackendTlsPropertiesResponsePtrOutput() BackendTlsPropertiesResponsePtrOutput {
	return o.ToBackendTlsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o BackendTlsPropertiesResponseOutput) ToBackendTlsPropertiesResponsePtrOutputWithContext(ctx context.Context) BackendTlsPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendTlsPropertiesResponse) *BackendTlsPropertiesResponse {
		return &v
	}).(BackendTlsPropertiesResponsePtrOutput)
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

type CertificateConfiguration struct {
	CertificatePassword *string `pulumi:"certificatePassword"`
	EncodedCertificate  *string `pulumi:"encodedCertificate"`
	StoreName           string  `pulumi:"storeName"`
}





type CertificateConfigurationInput interface {
	pulumi.Input

	ToCertificateConfigurationOutput() CertificateConfigurationOutput
	ToCertificateConfigurationOutputWithContext(context.Context) CertificateConfigurationOutput
}

type CertificateConfigurationArgs struct {
	CertificatePassword pulumi.StringPtrInput `pulumi:"certificatePassword"`
	EncodedCertificate  pulumi.StringPtrInput `pulumi:"encodedCertificate"`
	StoreName           pulumi.StringInput    `pulumi:"storeName"`
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
	Certificate         CertificateInformationResponse `pulumi:"certificate"`
	CertificatePassword *string                        `pulumi:"certificatePassword"`
	EncodedCertificate  *string                        `pulumi:"encodedCertificate"`
	StoreName           string                         `pulumi:"storeName"`
}





type CertificateConfigurationResponseInput interface {
	pulumi.Input

	ToCertificateConfigurationResponseOutput() CertificateConfigurationResponseOutput
	ToCertificateConfigurationResponseOutputWithContext(context.Context) CertificateConfigurationResponseOutput
}

type CertificateConfigurationResponseArgs struct {
	Certificate         CertificateInformationResponseInput `pulumi:"certificate"`
	CertificatePassword pulumi.StringPtrInput               `pulumi:"certificatePassword"`
	EncodedCertificate  pulumi.StringPtrInput               `pulumi:"encodedCertificate"`
	StoreName           pulumi.StringInput                  `pulumi:"storeName"`
}

func (CertificateConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateConfigurationResponse)(nil)).Elem()
}

func (i CertificateConfigurationResponseArgs) ToCertificateConfigurationResponseOutput() CertificateConfigurationResponseOutput {
	return i.ToCertificateConfigurationResponseOutputWithContext(context.Background())
}

func (i CertificateConfigurationResponseArgs) ToCertificateConfigurationResponseOutputWithContext(ctx context.Context) CertificateConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateConfigurationResponseOutput)
}





type CertificateConfigurationResponseArrayInput interface {
	pulumi.Input

	ToCertificateConfigurationResponseArrayOutput() CertificateConfigurationResponseArrayOutput
	ToCertificateConfigurationResponseArrayOutputWithContext(context.Context) CertificateConfigurationResponseArrayOutput
}

type CertificateConfigurationResponseArray []CertificateConfigurationResponseInput

func (CertificateConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateConfigurationResponse)(nil)).Elem()
}

func (i CertificateConfigurationResponseArray) ToCertificateConfigurationResponseArrayOutput() CertificateConfigurationResponseArrayOutput {
	return i.ToCertificateConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i CertificateConfigurationResponseArray) ToCertificateConfigurationResponseArrayOutputWithContext(ctx context.Context) CertificateConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateConfigurationResponseArrayOutput)
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

func (o CertificateConfigurationResponseOutput) Certificate() CertificateInformationResponseOutput {
	return o.ApplyT(func(v CertificateConfigurationResponse) CertificateInformationResponse { return v.Certificate }).(CertificateInformationResponseOutput)
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





type EmailTemplateParametersContractPropertiesResponseInput interface {
	pulumi.Input

	ToEmailTemplateParametersContractPropertiesResponseOutput() EmailTemplateParametersContractPropertiesResponseOutput
	ToEmailTemplateParametersContractPropertiesResponseOutputWithContext(context.Context) EmailTemplateParametersContractPropertiesResponseOutput
}

type EmailTemplateParametersContractPropertiesResponseArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	Title       pulumi.StringPtrInput `pulumi:"title"`
}

func (EmailTemplateParametersContractPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailTemplateParametersContractPropertiesResponse)(nil)).Elem()
}

func (i EmailTemplateParametersContractPropertiesResponseArgs) ToEmailTemplateParametersContractPropertiesResponseOutput() EmailTemplateParametersContractPropertiesResponseOutput {
	return i.ToEmailTemplateParametersContractPropertiesResponseOutputWithContext(context.Background())
}

func (i EmailTemplateParametersContractPropertiesResponseArgs) ToEmailTemplateParametersContractPropertiesResponseOutputWithContext(ctx context.Context) EmailTemplateParametersContractPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailTemplateParametersContractPropertiesResponseOutput)
}





type EmailTemplateParametersContractPropertiesResponseArrayInput interface {
	pulumi.Input

	ToEmailTemplateParametersContractPropertiesResponseArrayOutput() EmailTemplateParametersContractPropertiesResponseArrayOutput
	ToEmailTemplateParametersContractPropertiesResponseArrayOutputWithContext(context.Context) EmailTemplateParametersContractPropertiesResponseArrayOutput
}

type EmailTemplateParametersContractPropertiesResponseArray []EmailTemplateParametersContractPropertiesResponseInput

func (EmailTemplateParametersContractPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailTemplateParametersContractPropertiesResponse)(nil)).Elem()
}

func (i EmailTemplateParametersContractPropertiesResponseArray) ToEmailTemplateParametersContractPropertiesResponseArrayOutput() EmailTemplateParametersContractPropertiesResponseArrayOutput {
	return i.ToEmailTemplateParametersContractPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i EmailTemplateParametersContractPropertiesResponseArray) ToEmailTemplateParametersContractPropertiesResponseArrayOutputWithContext(ctx context.Context) EmailTemplateParametersContractPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailTemplateParametersContractPropertiesResponseArrayOutput)
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

type GroupContractResponse struct {
	BuiltIn     bool    `pulumi:"builtIn"`
	Description *string `pulumi:"description"`
	DisplayName string  `pulumi:"displayName"`
	ExternalId  *string `pulumi:"externalId"`
	Id          string  `pulumi:"id"`
	Name        string  `pulumi:"name"`
	Type        string  `pulumi:"type"`
}





type GroupContractResponseInput interface {
	pulumi.Input

	ToGroupContractResponseOutput() GroupContractResponseOutput
	ToGroupContractResponseOutputWithContext(context.Context) GroupContractResponseOutput
}

type GroupContractResponseArgs struct {
	BuiltIn     pulumi.BoolInput      `pulumi:"builtIn"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	DisplayName pulumi.StringInput    `pulumi:"displayName"`
	ExternalId  pulumi.StringPtrInput `pulumi:"externalId"`
	Id          pulumi.StringInput    `pulumi:"id"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (GroupContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupContractResponse)(nil)).Elem()
}

func (i GroupContractResponseArgs) ToGroupContractResponseOutput() GroupContractResponseOutput {
	return i.ToGroupContractResponseOutputWithContext(context.Background())
}

func (i GroupContractResponseArgs) ToGroupContractResponseOutputWithContext(ctx context.Context) GroupContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupContractResponseOutput)
}





type GroupContractResponseArrayInput interface {
	pulumi.Input

	ToGroupContractResponseArrayOutput() GroupContractResponseArrayOutput
	ToGroupContractResponseArrayOutputWithContext(context.Context) GroupContractResponseArrayOutput
}

type GroupContractResponseArray []GroupContractResponseInput

func (GroupContractResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupContractResponse)(nil)).Elem()
}

func (i GroupContractResponseArray) ToGroupContractResponseArrayOutput() GroupContractResponseArrayOutput {
	return i.ToGroupContractResponseArrayOutputWithContext(context.Background())
}

func (i GroupContractResponseArray) ToGroupContractResponseArrayOutputWithContext(ctx context.Context) GroupContractResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupContractResponseArrayOutput)
}

type GroupContractResponseOutput struct{ *pulumi.OutputState }

func (GroupContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupContractResponse)(nil)).Elem()
}

func (o GroupContractResponseOutput) ToGroupContractResponseOutput() GroupContractResponseOutput {
	return o
}

func (o GroupContractResponseOutput) ToGroupContractResponseOutputWithContext(ctx context.Context) GroupContractResponseOutput {
	return o
}

func (o GroupContractResponseOutput) BuiltIn() pulumi.BoolOutput {
	return o.ApplyT(func(v GroupContractResponse) bool { return v.BuiltIn }).(pulumi.BoolOutput)
}

func (o GroupContractResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupContractResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GroupContractResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GroupContractResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o GroupContractResponseOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupContractResponse) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

func (o GroupContractResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GroupContractResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o GroupContractResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GroupContractResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o GroupContractResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GroupContractResponse) string { return v.Type }).(pulumi.StringOutput)
}

type GroupContractResponseArrayOutput struct{ *pulumi.OutputState }

func (GroupContractResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupContractResponse)(nil)).Elem()
}

func (o GroupContractResponseArrayOutput) ToGroupContractResponseArrayOutput() GroupContractResponseArrayOutput {
	return o
}

func (o GroupContractResponseArrayOutput) ToGroupContractResponseArrayOutputWithContext(ctx context.Context) GroupContractResponseArrayOutput {
	return o
}

func (o GroupContractResponseArrayOutput) Index(i pulumi.IntInput) GroupContractResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupContractResponse {
		return vs[0].([]GroupContractResponse)[vs[1].(int)]
	}).(GroupContractResponseOutput)
}

type HostnameConfiguration struct {
	CertificatePassword        *string      `pulumi:"certificatePassword"`
	DefaultSslBinding          *bool        `pulumi:"defaultSslBinding"`
	EncodedCertificate         *string      `pulumi:"encodedCertificate"`
	HostName                   string       `pulumi:"hostName"`
	KeyVaultId                 *string      `pulumi:"keyVaultId"`
	NegotiateClientCertificate *bool        `pulumi:"negotiateClientCertificate"`
	Type                       HostnameType `pulumi:"type"`
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
	CertificatePassword        pulumi.StringPtrInput `pulumi:"certificatePassword"`
	DefaultSslBinding          pulumi.BoolPtrInput   `pulumi:"defaultSslBinding"`
	EncodedCertificate         pulumi.StringPtrInput `pulumi:"encodedCertificate"`
	HostName                   pulumi.StringInput    `pulumi:"hostName"`
	KeyVaultId                 pulumi.StringPtrInput `pulumi:"keyVaultId"`
	NegotiateClientCertificate pulumi.BoolPtrInput   `pulumi:"negotiateClientCertificate"`
	Type                       HostnameTypeInput     `pulumi:"type"`
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

func (o HostnameConfigurationOutput) CertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *string { return v.CertificatePassword }).(pulumi.StringPtrOutput)
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

func (o HostnameConfigurationOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o HostnameConfigurationOutput) NegotiateClientCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostnameConfiguration) *bool { return v.NegotiateClientCertificate }).(pulumi.BoolPtrOutput)
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
	Certificate                CertificateInformationResponse `pulumi:"certificate"`
	CertificatePassword        *string                        `pulumi:"certificatePassword"`
	DefaultSslBinding          *bool                          `pulumi:"defaultSslBinding"`
	EncodedCertificate         *string                        `pulumi:"encodedCertificate"`
	HostName                   string                         `pulumi:"hostName"`
	KeyVaultId                 *string                        `pulumi:"keyVaultId"`
	NegotiateClientCertificate *bool                          `pulumi:"negotiateClientCertificate"`
	Type                       string                         `pulumi:"type"`
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





type HostnameConfigurationResponseInput interface {
	pulumi.Input

	ToHostnameConfigurationResponseOutput() HostnameConfigurationResponseOutput
	ToHostnameConfigurationResponseOutputWithContext(context.Context) HostnameConfigurationResponseOutput
}

type HostnameConfigurationResponseArgs struct {
	Certificate                CertificateInformationResponseInput `pulumi:"certificate"`
	CertificatePassword        pulumi.StringPtrInput               `pulumi:"certificatePassword"`
	DefaultSslBinding          pulumi.BoolPtrInput                 `pulumi:"defaultSslBinding"`
	EncodedCertificate         pulumi.StringPtrInput               `pulumi:"encodedCertificate"`
	HostName                   pulumi.StringInput                  `pulumi:"hostName"`
	KeyVaultId                 pulumi.StringPtrInput               `pulumi:"keyVaultId"`
	NegotiateClientCertificate pulumi.BoolPtrInput                 `pulumi:"negotiateClientCertificate"`
	Type                       pulumi.StringInput                  `pulumi:"type"`
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

func (o HostnameConfigurationResponseOutput) CertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostnameConfigurationResponse) *string { return v.CertificatePassword }).(pulumi.StringPtrOutput)
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

type LoggerSamplingContract struct {
	EvaluationInterval         *string  `pulumi:"evaluationInterval"`
	InitialPercentage          *float64 `pulumi:"initialPercentage"`
	MaxPercentage              *float64 `pulumi:"maxPercentage"`
	MaxTelemetryItemsPerSecond *int     `pulumi:"maxTelemetryItemsPerSecond"`
	MinPercentage              *float64 `pulumi:"minPercentage"`
	MovingAverageRatio         *float64 `pulumi:"movingAverageRatio"`
	Percentage                 *float64 `pulumi:"percentage"`
	PercentageDecreaseTimeout  *string  `pulumi:"percentageDecreaseTimeout"`
	PercentageIncreaseTimeout  *string  `pulumi:"percentageIncreaseTimeout"`
	SamplingType               *string  `pulumi:"samplingType"`
}





type LoggerSamplingContractInput interface {
	pulumi.Input

	ToLoggerSamplingContractOutput() LoggerSamplingContractOutput
	ToLoggerSamplingContractOutputWithContext(context.Context) LoggerSamplingContractOutput
}

type LoggerSamplingContractArgs struct {
	EvaluationInterval         pulumi.StringPtrInput  `pulumi:"evaluationInterval"`
	InitialPercentage          pulumi.Float64PtrInput `pulumi:"initialPercentage"`
	MaxPercentage              pulumi.Float64PtrInput `pulumi:"maxPercentage"`
	MaxTelemetryItemsPerSecond pulumi.IntPtrInput     `pulumi:"maxTelemetryItemsPerSecond"`
	MinPercentage              pulumi.Float64PtrInput `pulumi:"minPercentage"`
	MovingAverageRatio         pulumi.Float64PtrInput `pulumi:"movingAverageRatio"`
	Percentage                 pulumi.Float64PtrInput `pulumi:"percentage"`
	PercentageDecreaseTimeout  pulumi.StringPtrInput  `pulumi:"percentageDecreaseTimeout"`
	PercentageIncreaseTimeout  pulumi.StringPtrInput  `pulumi:"percentageIncreaseTimeout"`
	SamplingType               pulumi.StringPtrInput  `pulumi:"samplingType"`
}

func (LoggerSamplingContractArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggerSamplingContract)(nil)).Elem()
}

func (i LoggerSamplingContractArgs) ToLoggerSamplingContractOutput() LoggerSamplingContractOutput {
	return i.ToLoggerSamplingContractOutputWithContext(context.Background())
}

func (i LoggerSamplingContractArgs) ToLoggerSamplingContractOutputWithContext(ctx context.Context) LoggerSamplingContractOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggerSamplingContractOutput)
}

func (i LoggerSamplingContractArgs) ToLoggerSamplingContractPtrOutput() LoggerSamplingContractPtrOutput {
	return i.ToLoggerSamplingContractPtrOutputWithContext(context.Background())
}

func (i LoggerSamplingContractArgs) ToLoggerSamplingContractPtrOutputWithContext(ctx context.Context) LoggerSamplingContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggerSamplingContractOutput).ToLoggerSamplingContractPtrOutputWithContext(ctx)
}









type LoggerSamplingContractPtrInput interface {
	pulumi.Input

	ToLoggerSamplingContractPtrOutput() LoggerSamplingContractPtrOutput
	ToLoggerSamplingContractPtrOutputWithContext(context.Context) LoggerSamplingContractPtrOutput
}

type loggerSamplingContractPtrType LoggerSamplingContractArgs

func LoggerSamplingContractPtr(v *LoggerSamplingContractArgs) LoggerSamplingContractPtrInput {
	return (*loggerSamplingContractPtrType)(v)
}

func (*loggerSamplingContractPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggerSamplingContract)(nil)).Elem()
}

func (i *loggerSamplingContractPtrType) ToLoggerSamplingContractPtrOutput() LoggerSamplingContractPtrOutput {
	return i.ToLoggerSamplingContractPtrOutputWithContext(context.Background())
}

func (i *loggerSamplingContractPtrType) ToLoggerSamplingContractPtrOutputWithContext(ctx context.Context) LoggerSamplingContractPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggerSamplingContractPtrOutput)
}

type LoggerSamplingContractOutput struct{ *pulumi.OutputState }

func (LoggerSamplingContractOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggerSamplingContract)(nil)).Elem()
}

func (o LoggerSamplingContractOutput) ToLoggerSamplingContractOutput() LoggerSamplingContractOutput {
	return o
}

func (o LoggerSamplingContractOutput) ToLoggerSamplingContractOutputWithContext(ctx context.Context) LoggerSamplingContractOutput {
	return o
}

func (o LoggerSamplingContractOutput) ToLoggerSamplingContractPtrOutput() LoggerSamplingContractPtrOutput {
	return o.ToLoggerSamplingContractPtrOutputWithContext(context.Background())
}

func (o LoggerSamplingContractOutput) ToLoggerSamplingContractPtrOutputWithContext(ctx context.Context) LoggerSamplingContractPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoggerSamplingContract) *LoggerSamplingContract {
		return &v
	}).(LoggerSamplingContractPtrOutput)
}

func (o LoggerSamplingContractOutput) EvaluationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *string { return v.EvaluationInterval }).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractOutput) InitialPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *float64 { return v.InitialPercentage }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractOutput) MaxPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *float64 { return v.MaxPercentage }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractOutput) MaxTelemetryItemsPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *int { return v.MaxTelemetryItemsPerSecond }).(pulumi.IntPtrOutput)
}

func (o LoggerSamplingContractOutput) MinPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *float64 { return v.MinPercentage }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractOutput) MovingAverageRatio() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *float64 { return v.MovingAverageRatio }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *float64 { return v.Percentage }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractOutput) PercentageDecreaseTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *string { return v.PercentageDecreaseTimeout }).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractOutput) PercentageIncreaseTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *string { return v.PercentageIncreaseTimeout }).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractOutput) SamplingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContract) *string { return v.SamplingType }).(pulumi.StringPtrOutput)
}

type LoggerSamplingContractPtrOutput struct{ *pulumi.OutputState }

func (LoggerSamplingContractPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggerSamplingContract)(nil)).Elem()
}

func (o LoggerSamplingContractPtrOutput) ToLoggerSamplingContractPtrOutput() LoggerSamplingContractPtrOutput {
	return o
}

func (o LoggerSamplingContractPtrOutput) ToLoggerSamplingContractPtrOutputWithContext(ctx context.Context) LoggerSamplingContractPtrOutput {
	return o
}

func (o LoggerSamplingContractPtrOutput) Elem() LoggerSamplingContractOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) LoggerSamplingContract {
		if v != nil {
			return *v
		}
		var ret LoggerSamplingContract
		return ret
	}).(LoggerSamplingContractOutput)
}

func (o LoggerSamplingContractPtrOutput) EvaluationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *string {
		if v == nil {
			return nil
		}
		return v.EvaluationInterval
	}).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractPtrOutput) InitialPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *float64 {
		if v == nil {
			return nil
		}
		return v.InitialPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractPtrOutput) MaxPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractPtrOutput) MaxTelemetryItemsPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *int {
		if v == nil {
			return nil
		}
		return v.MaxTelemetryItemsPerSecond
	}).(pulumi.IntPtrOutput)
}

func (o LoggerSamplingContractPtrOutput) MinPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *float64 {
		if v == nil {
			return nil
		}
		return v.MinPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractPtrOutput) MovingAverageRatio() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *float64 {
		if v == nil {
			return nil
		}
		return v.MovingAverageRatio
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractPtrOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *float64 {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractPtrOutput) PercentageDecreaseTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *string {
		if v == nil {
			return nil
		}
		return v.PercentageDecreaseTimeout
	}).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractPtrOutput) PercentageIncreaseTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *string {
		if v == nil {
			return nil
		}
		return v.PercentageIncreaseTimeout
	}).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractPtrOutput) SamplingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContract) *string {
		if v == nil {
			return nil
		}
		return v.SamplingType
	}).(pulumi.StringPtrOutput)
}

type LoggerSamplingContractResponse struct {
	EvaluationInterval         *string  `pulumi:"evaluationInterval"`
	InitialPercentage          *float64 `pulumi:"initialPercentage"`
	MaxPercentage              *float64 `pulumi:"maxPercentage"`
	MaxTelemetryItemsPerSecond *int     `pulumi:"maxTelemetryItemsPerSecond"`
	MinPercentage              *float64 `pulumi:"minPercentage"`
	MovingAverageRatio         *float64 `pulumi:"movingAverageRatio"`
	Percentage                 *float64 `pulumi:"percentage"`
	PercentageDecreaseTimeout  *string  `pulumi:"percentageDecreaseTimeout"`
	PercentageIncreaseTimeout  *string  `pulumi:"percentageIncreaseTimeout"`
	SamplingType               *string  `pulumi:"samplingType"`
}





type LoggerSamplingContractResponseInput interface {
	pulumi.Input

	ToLoggerSamplingContractResponseOutput() LoggerSamplingContractResponseOutput
	ToLoggerSamplingContractResponseOutputWithContext(context.Context) LoggerSamplingContractResponseOutput
}

type LoggerSamplingContractResponseArgs struct {
	EvaluationInterval         pulumi.StringPtrInput  `pulumi:"evaluationInterval"`
	InitialPercentage          pulumi.Float64PtrInput `pulumi:"initialPercentage"`
	MaxPercentage              pulumi.Float64PtrInput `pulumi:"maxPercentage"`
	MaxTelemetryItemsPerSecond pulumi.IntPtrInput     `pulumi:"maxTelemetryItemsPerSecond"`
	MinPercentage              pulumi.Float64PtrInput `pulumi:"minPercentage"`
	MovingAverageRatio         pulumi.Float64PtrInput `pulumi:"movingAverageRatio"`
	Percentage                 pulumi.Float64PtrInput `pulumi:"percentage"`
	PercentageDecreaseTimeout  pulumi.StringPtrInput  `pulumi:"percentageDecreaseTimeout"`
	PercentageIncreaseTimeout  pulumi.StringPtrInput  `pulumi:"percentageIncreaseTimeout"`
	SamplingType               pulumi.StringPtrInput  `pulumi:"samplingType"`
}

func (LoggerSamplingContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggerSamplingContractResponse)(nil)).Elem()
}

func (i LoggerSamplingContractResponseArgs) ToLoggerSamplingContractResponseOutput() LoggerSamplingContractResponseOutput {
	return i.ToLoggerSamplingContractResponseOutputWithContext(context.Background())
}

func (i LoggerSamplingContractResponseArgs) ToLoggerSamplingContractResponseOutputWithContext(ctx context.Context) LoggerSamplingContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggerSamplingContractResponseOutput)
}

func (i LoggerSamplingContractResponseArgs) ToLoggerSamplingContractResponsePtrOutput() LoggerSamplingContractResponsePtrOutput {
	return i.ToLoggerSamplingContractResponsePtrOutputWithContext(context.Background())
}

func (i LoggerSamplingContractResponseArgs) ToLoggerSamplingContractResponsePtrOutputWithContext(ctx context.Context) LoggerSamplingContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggerSamplingContractResponseOutput).ToLoggerSamplingContractResponsePtrOutputWithContext(ctx)
}









type LoggerSamplingContractResponsePtrInput interface {
	pulumi.Input

	ToLoggerSamplingContractResponsePtrOutput() LoggerSamplingContractResponsePtrOutput
	ToLoggerSamplingContractResponsePtrOutputWithContext(context.Context) LoggerSamplingContractResponsePtrOutput
}

type loggerSamplingContractResponsePtrType LoggerSamplingContractResponseArgs

func LoggerSamplingContractResponsePtr(v *LoggerSamplingContractResponseArgs) LoggerSamplingContractResponsePtrInput {
	return (*loggerSamplingContractResponsePtrType)(v)
}

func (*loggerSamplingContractResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggerSamplingContractResponse)(nil)).Elem()
}

func (i *loggerSamplingContractResponsePtrType) ToLoggerSamplingContractResponsePtrOutput() LoggerSamplingContractResponsePtrOutput {
	return i.ToLoggerSamplingContractResponsePtrOutputWithContext(context.Background())
}

func (i *loggerSamplingContractResponsePtrType) ToLoggerSamplingContractResponsePtrOutputWithContext(ctx context.Context) LoggerSamplingContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggerSamplingContractResponsePtrOutput)
}

type LoggerSamplingContractResponseOutput struct{ *pulumi.OutputState }

func (LoggerSamplingContractResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggerSamplingContractResponse)(nil)).Elem()
}

func (o LoggerSamplingContractResponseOutput) ToLoggerSamplingContractResponseOutput() LoggerSamplingContractResponseOutput {
	return o
}

func (o LoggerSamplingContractResponseOutput) ToLoggerSamplingContractResponseOutputWithContext(ctx context.Context) LoggerSamplingContractResponseOutput {
	return o
}

func (o LoggerSamplingContractResponseOutput) ToLoggerSamplingContractResponsePtrOutput() LoggerSamplingContractResponsePtrOutput {
	return o.ToLoggerSamplingContractResponsePtrOutputWithContext(context.Background())
}

func (o LoggerSamplingContractResponseOutput) ToLoggerSamplingContractResponsePtrOutputWithContext(ctx context.Context) LoggerSamplingContractResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoggerSamplingContractResponse) *LoggerSamplingContractResponse {
		return &v
	}).(LoggerSamplingContractResponsePtrOutput)
}

func (o LoggerSamplingContractResponseOutput) EvaluationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *string { return v.EvaluationInterval }).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractResponseOutput) InitialPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *float64 { return v.InitialPercentage }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponseOutput) MaxPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *float64 { return v.MaxPercentage }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponseOutput) MaxTelemetryItemsPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *int { return v.MaxTelemetryItemsPerSecond }).(pulumi.IntPtrOutput)
}

func (o LoggerSamplingContractResponseOutput) MinPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *float64 { return v.MinPercentage }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponseOutput) MovingAverageRatio() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *float64 { return v.MovingAverageRatio }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponseOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *float64 { return v.Percentage }).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponseOutput) PercentageDecreaseTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *string { return v.PercentageDecreaseTimeout }).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractResponseOutput) PercentageIncreaseTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *string { return v.PercentageIncreaseTimeout }).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractResponseOutput) SamplingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggerSamplingContractResponse) *string { return v.SamplingType }).(pulumi.StringPtrOutput)
}

type LoggerSamplingContractResponsePtrOutput struct{ *pulumi.OutputState }

func (LoggerSamplingContractResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggerSamplingContractResponse)(nil)).Elem()
}

func (o LoggerSamplingContractResponsePtrOutput) ToLoggerSamplingContractResponsePtrOutput() LoggerSamplingContractResponsePtrOutput {
	return o
}

func (o LoggerSamplingContractResponsePtrOutput) ToLoggerSamplingContractResponsePtrOutputWithContext(ctx context.Context) LoggerSamplingContractResponsePtrOutput {
	return o
}

func (o LoggerSamplingContractResponsePtrOutput) Elem() LoggerSamplingContractResponseOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) LoggerSamplingContractResponse {
		if v != nil {
			return *v
		}
		var ret LoggerSamplingContractResponse
		return ret
	}).(LoggerSamplingContractResponseOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) EvaluationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.EvaluationInterval
	}).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) InitialPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.InitialPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) MaxPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) MaxTelemetryItemsPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxTelemetryItemsPerSecond
	}).(pulumi.IntPtrOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) MinPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MinPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) MovingAverageRatio() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MovingAverageRatio
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) Percentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.Float64PtrOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) PercentageDecreaseTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.PercentageDecreaseTimeout
	}).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) PercentageIncreaseTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.PercentageIncreaseTimeout
	}).(pulumi.StringPtrOutput)
}

func (o LoggerSamplingContractResponsePtrOutput) SamplingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggerSamplingContractResponse) *string {
		if v == nil {
			return nil
		}
		return v.SamplingType
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





type OAuth2AuthenticationSettingsContractResponseInput interface {
	pulumi.Input

	ToOAuth2AuthenticationSettingsContractResponseOutput() OAuth2AuthenticationSettingsContractResponseOutput
	ToOAuth2AuthenticationSettingsContractResponseOutputWithContext(context.Context) OAuth2AuthenticationSettingsContractResponseOutput
}

type OAuth2AuthenticationSettingsContractResponseArgs struct {
	AuthorizationServerId pulumi.StringPtrInput `pulumi:"authorizationServerId"`
	Scope                 pulumi.StringPtrInput `pulumi:"scope"`
}

func (OAuth2AuthenticationSettingsContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OAuth2AuthenticationSettingsContractResponse)(nil)).Elem()
}

func (i OAuth2AuthenticationSettingsContractResponseArgs) ToOAuth2AuthenticationSettingsContractResponseOutput() OAuth2AuthenticationSettingsContractResponseOutput {
	return i.ToOAuth2AuthenticationSettingsContractResponseOutputWithContext(context.Background())
}

func (i OAuth2AuthenticationSettingsContractResponseArgs) ToOAuth2AuthenticationSettingsContractResponseOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OAuth2AuthenticationSettingsContractResponseOutput)
}

func (i OAuth2AuthenticationSettingsContractResponseArgs) ToOAuth2AuthenticationSettingsContractResponsePtrOutput() OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return i.ToOAuth2AuthenticationSettingsContractResponsePtrOutputWithContext(context.Background())
}

func (i OAuth2AuthenticationSettingsContractResponseArgs) ToOAuth2AuthenticationSettingsContractResponsePtrOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OAuth2AuthenticationSettingsContractResponseOutput).ToOAuth2AuthenticationSettingsContractResponsePtrOutputWithContext(ctx)
}









type OAuth2AuthenticationSettingsContractResponsePtrInput interface {
	pulumi.Input

	ToOAuth2AuthenticationSettingsContractResponsePtrOutput() OAuth2AuthenticationSettingsContractResponsePtrOutput
	ToOAuth2AuthenticationSettingsContractResponsePtrOutputWithContext(context.Context) OAuth2AuthenticationSettingsContractResponsePtrOutput
}

type oauth2AuthenticationSettingsContractResponsePtrType OAuth2AuthenticationSettingsContractResponseArgs

func OAuth2AuthenticationSettingsContractResponsePtr(v *OAuth2AuthenticationSettingsContractResponseArgs) OAuth2AuthenticationSettingsContractResponsePtrInput {
	return (*oauth2AuthenticationSettingsContractResponsePtrType)(v)
}

func (*oauth2AuthenticationSettingsContractResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OAuth2AuthenticationSettingsContractResponse)(nil)).Elem()
}

func (i *oauth2AuthenticationSettingsContractResponsePtrType) ToOAuth2AuthenticationSettingsContractResponsePtrOutput() OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return i.ToOAuth2AuthenticationSettingsContractResponsePtrOutputWithContext(context.Background())
}

func (i *oauth2AuthenticationSettingsContractResponsePtrType) ToOAuth2AuthenticationSettingsContractResponsePtrOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OAuth2AuthenticationSettingsContractResponsePtrOutput)
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

func (o OAuth2AuthenticationSettingsContractResponseOutput) ToOAuth2AuthenticationSettingsContractResponsePtrOutput() OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return o.ToOAuth2AuthenticationSettingsContractResponsePtrOutputWithContext(context.Background())
}

func (o OAuth2AuthenticationSettingsContractResponseOutput) ToOAuth2AuthenticationSettingsContractResponsePtrOutputWithContext(ctx context.Context) OAuth2AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OAuth2AuthenticationSettingsContractResponse) *OAuth2AuthenticationSettingsContractResponse {
		return &v
	}).(OAuth2AuthenticationSettingsContractResponsePtrOutput)
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

type ParameterContract struct {
	DefaultValue *string  `pulumi:"defaultValue"`
	Description  *string  `pulumi:"description"`
	Name         string   `pulumi:"name"`
	Required     *bool    `pulumi:"required"`
	Type         string   `pulumi:"type"`
	Values       []string `pulumi:"values"`
}





type ParameterContractInput interface {
	pulumi.Input

	ToParameterContractOutput() ParameterContractOutput
	ToParameterContractOutputWithContext(context.Context) ParameterContractOutput
}

type ParameterContractArgs struct {
	DefaultValue pulumi.StringPtrInput   `pulumi:"defaultValue"`
	Description  pulumi.StringPtrInput   `pulumi:"description"`
	Name         pulumi.StringInput      `pulumi:"name"`
	Required     pulumi.BoolPtrInput     `pulumi:"required"`
	Type         pulumi.StringInput      `pulumi:"type"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
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

func (o ParameterContractOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterContract) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterContractOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ParameterContract) *bool { return v.Required }).(pulumi.BoolPtrOutput)
}

func (o ParameterContractOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterContract) string { return v.Type }).(pulumi.StringOutput)
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
	DefaultValue *string  `pulumi:"defaultValue"`
	Description  *string  `pulumi:"description"`
	Name         string   `pulumi:"name"`
	Required     *bool    `pulumi:"required"`
	Type         string   `pulumi:"type"`
	Values       []string `pulumi:"values"`
}





type ParameterContractResponseInput interface {
	pulumi.Input

	ToParameterContractResponseOutput() ParameterContractResponseOutput
	ToParameterContractResponseOutputWithContext(context.Context) ParameterContractResponseOutput
}

type ParameterContractResponseArgs struct {
	DefaultValue pulumi.StringPtrInput   `pulumi:"defaultValue"`
	Description  pulumi.StringPtrInput   `pulumi:"description"`
	Name         pulumi.StringInput      `pulumi:"name"`
	Required     pulumi.BoolPtrInput     `pulumi:"required"`
	Type         pulumi.StringInput      `pulumi:"type"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (ParameterContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterContractResponse)(nil)).Elem()
}

func (i ParameterContractResponseArgs) ToParameterContractResponseOutput() ParameterContractResponseOutput {
	return i.ToParameterContractResponseOutputWithContext(context.Background())
}

func (i ParameterContractResponseArgs) ToParameterContractResponseOutputWithContext(ctx context.Context) ParameterContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterContractResponseOutput)
}





type ParameterContractResponseArrayInput interface {
	pulumi.Input

	ToParameterContractResponseArrayOutput() ParameterContractResponseArrayOutput
	ToParameterContractResponseArrayOutputWithContext(context.Context) ParameterContractResponseArrayOutput
}

type ParameterContractResponseArray []ParameterContractResponseInput

func (ParameterContractResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterContractResponse)(nil)).Elem()
}

func (i ParameterContractResponseArray) ToParameterContractResponseArrayOutput() ParameterContractResponseArrayOutput {
	return i.ToParameterContractResponseArrayOutputWithContext(context.Background())
}

func (i ParameterContractResponseArray) ToParameterContractResponseArrayOutputWithContext(ctx context.Context) ParameterContractResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterContractResponseArrayOutput)
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

func (o ParameterContractResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterContractResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterContractResponseOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ParameterContractResponse) *bool { return v.Required }).(pulumi.BoolPtrOutput)
}

func (o ParameterContractResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterContractResponse) string { return v.Type }).(pulumi.StringOutput)
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

type RepresentationContract struct {
	ContentType    string              `pulumi:"contentType"`
	FormParameters []ParameterContract `pulumi:"formParameters"`
	Sample         *string             `pulumi:"sample"`
	SchemaId       *string             `pulumi:"schemaId"`
	TypeName       *string             `pulumi:"typeName"`
}





type RepresentationContractInput interface {
	pulumi.Input

	ToRepresentationContractOutput() RepresentationContractOutput
	ToRepresentationContractOutputWithContext(context.Context) RepresentationContractOutput
}

type RepresentationContractArgs struct {
	ContentType    pulumi.StringInput          `pulumi:"contentType"`
	FormParameters ParameterContractArrayInput `pulumi:"formParameters"`
	Sample         pulumi.StringPtrInput       `pulumi:"sample"`
	SchemaId       pulumi.StringPtrInput       `pulumi:"schemaId"`
	TypeName       pulumi.StringPtrInput       `pulumi:"typeName"`
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

func (o RepresentationContractOutput) FormParameters() ParameterContractArrayOutput {
	return o.ApplyT(func(v RepresentationContract) []ParameterContract { return v.FormParameters }).(ParameterContractArrayOutput)
}

func (o RepresentationContractOutput) Sample() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepresentationContract) *string { return v.Sample }).(pulumi.StringPtrOutput)
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
	ContentType    string                      `pulumi:"contentType"`
	FormParameters []ParameterContractResponse `pulumi:"formParameters"`
	Sample         *string                     `pulumi:"sample"`
	SchemaId       *string                     `pulumi:"schemaId"`
	TypeName       *string                     `pulumi:"typeName"`
}





type RepresentationContractResponseInput interface {
	pulumi.Input

	ToRepresentationContractResponseOutput() RepresentationContractResponseOutput
	ToRepresentationContractResponseOutputWithContext(context.Context) RepresentationContractResponseOutput
}

type RepresentationContractResponseArgs struct {
	ContentType    pulumi.StringInput                  `pulumi:"contentType"`
	FormParameters ParameterContractResponseArrayInput `pulumi:"formParameters"`
	Sample         pulumi.StringPtrInput               `pulumi:"sample"`
	SchemaId       pulumi.StringPtrInput               `pulumi:"schemaId"`
	TypeName       pulumi.StringPtrInput               `pulumi:"typeName"`
}

func (RepresentationContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepresentationContractResponse)(nil)).Elem()
}

func (i RepresentationContractResponseArgs) ToRepresentationContractResponseOutput() RepresentationContractResponseOutput {
	return i.ToRepresentationContractResponseOutputWithContext(context.Background())
}

func (i RepresentationContractResponseArgs) ToRepresentationContractResponseOutputWithContext(ctx context.Context) RepresentationContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepresentationContractResponseOutput)
}





type RepresentationContractResponseArrayInput interface {
	pulumi.Input

	ToRepresentationContractResponseArrayOutput() RepresentationContractResponseArrayOutput
	ToRepresentationContractResponseArrayOutputWithContext(context.Context) RepresentationContractResponseArrayOutput
}

type RepresentationContractResponseArray []RepresentationContractResponseInput

func (RepresentationContractResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepresentationContractResponse)(nil)).Elem()
}

func (i RepresentationContractResponseArray) ToRepresentationContractResponseArrayOutput() RepresentationContractResponseArrayOutput {
	return i.ToRepresentationContractResponseArrayOutputWithContext(context.Background())
}

func (i RepresentationContractResponseArray) ToRepresentationContractResponseArrayOutputWithContext(ctx context.Context) RepresentationContractResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepresentationContractResponseArrayOutput)
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

func (o RepresentationContractResponseOutput) FormParameters() ParameterContractResponseArrayOutput {
	return o.ApplyT(func(v RepresentationContractResponse) []ParameterContractResponse { return v.FormParameters }).(ParameterContractResponseArrayOutput)
}

func (o RepresentationContractResponseOutput) Sample() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepresentationContractResponse) *string { return v.Sample }).(pulumi.StringPtrOutput)
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





type RequestContractResponseInput interface {
	pulumi.Input

	ToRequestContractResponseOutput() RequestContractResponseOutput
	ToRequestContractResponseOutputWithContext(context.Context) RequestContractResponseOutput
}

type RequestContractResponseArgs struct {
	Description     pulumi.StringPtrInput                    `pulumi:"description"`
	Headers         ParameterContractResponseArrayInput      `pulumi:"headers"`
	QueryParameters ParameterContractResponseArrayInput      `pulumi:"queryParameters"`
	Representations RepresentationContractResponseArrayInput `pulumi:"representations"`
}

func (RequestContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestContractResponse)(nil)).Elem()
}

func (i RequestContractResponseArgs) ToRequestContractResponseOutput() RequestContractResponseOutput {
	return i.ToRequestContractResponseOutputWithContext(context.Background())
}

func (i RequestContractResponseArgs) ToRequestContractResponseOutputWithContext(ctx context.Context) RequestContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestContractResponseOutput)
}

func (i RequestContractResponseArgs) ToRequestContractResponsePtrOutput() RequestContractResponsePtrOutput {
	return i.ToRequestContractResponsePtrOutputWithContext(context.Background())
}

func (i RequestContractResponseArgs) ToRequestContractResponsePtrOutputWithContext(ctx context.Context) RequestContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestContractResponseOutput).ToRequestContractResponsePtrOutputWithContext(ctx)
}









type RequestContractResponsePtrInput interface {
	pulumi.Input

	ToRequestContractResponsePtrOutput() RequestContractResponsePtrOutput
	ToRequestContractResponsePtrOutputWithContext(context.Context) RequestContractResponsePtrOutput
}

type requestContractResponsePtrType RequestContractResponseArgs

func RequestContractResponsePtr(v *RequestContractResponseArgs) RequestContractResponsePtrInput {
	return (*requestContractResponsePtrType)(v)
}

func (*requestContractResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestContractResponse)(nil)).Elem()
}

func (i *requestContractResponsePtrType) ToRequestContractResponsePtrOutput() RequestContractResponsePtrOutput {
	return i.ToRequestContractResponsePtrOutputWithContext(context.Background())
}

func (i *requestContractResponsePtrType) ToRequestContractResponsePtrOutputWithContext(ctx context.Context) RequestContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestContractResponsePtrOutput)
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

func (o RequestContractResponseOutput) ToRequestContractResponsePtrOutput() RequestContractResponsePtrOutput {
	return o.ToRequestContractResponsePtrOutputWithContext(context.Background())
}

func (o RequestContractResponseOutput) ToRequestContractResponsePtrOutputWithContext(ctx context.Context) RequestContractResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequestContractResponse) *RequestContractResponse {
		return &v
	}).(RequestContractResponsePtrOutput)
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





type ResponseContractResponseInput interface {
	pulumi.Input

	ToResponseContractResponseOutput() ResponseContractResponseOutput
	ToResponseContractResponseOutputWithContext(context.Context) ResponseContractResponseOutput
}

type ResponseContractResponseArgs struct {
	Description     pulumi.StringPtrInput                    `pulumi:"description"`
	Headers         ParameterContractResponseArrayInput      `pulumi:"headers"`
	Representations RepresentationContractResponseArrayInput `pulumi:"representations"`
	StatusCode      pulumi.IntInput                          `pulumi:"statusCode"`
}

func (ResponseContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseContractResponse)(nil)).Elem()
}

func (i ResponseContractResponseArgs) ToResponseContractResponseOutput() ResponseContractResponseOutput {
	return i.ToResponseContractResponseOutputWithContext(context.Background())
}

func (i ResponseContractResponseArgs) ToResponseContractResponseOutputWithContext(ctx context.Context) ResponseContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseContractResponseOutput)
}





type ResponseContractResponseArrayInput interface {
	pulumi.Input

	ToResponseContractResponseArrayOutput() ResponseContractResponseArrayOutput
	ToResponseContractResponseArrayOutputWithContext(context.Context) ResponseContractResponseArrayOutput
}

type ResponseContractResponseArray []ResponseContractResponseInput

func (ResponseContractResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResponseContractResponse)(nil)).Elem()
}

func (i ResponseContractResponseArray) ToResponseContractResponseArrayOutput() ResponseContractResponseArrayOutput {
	return i.ToResponseContractResponseArrayOutputWithContext(context.Background())
}

func (i ResponseContractResponseArray) ToResponseContractResponseArrayOutputWithContext(ctx context.Context) ResponseContractResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseContractResponseArrayOutput)
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





type SubscriptionKeyParameterNamesContractResponseInput interface {
	pulumi.Input

	ToSubscriptionKeyParameterNamesContractResponseOutput() SubscriptionKeyParameterNamesContractResponseOutput
	ToSubscriptionKeyParameterNamesContractResponseOutputWithContext(context.Context) SubscriptionKeyParameterNamesContractResponseOutput
}

type SubscriptionKeyParameterNamesContractResponseArgs struct {
	Header pulumi.StringPtrInput `pulumi:"header"`
	Query  pulumi.StringPtrInput `pulumi:"query"`
}

func (SubscriptionKeyParameterNamesContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionKeyParameterNamesContractResponse)(nil)).Elem()
}

func (i SubscriptionKeyParameterNamesContractResponseArgs) ToSubscriptionKeyParameterNamesContractResponseOutput() SubscriptionKeyParameterNamesContractResponseOutput {
	return i.ToSubscriptionKeyParameterNamesContractResponseOutputWithContext(context.Background())
}

func (i SubscriptionKeyParameterNamesContractResponseArgs) ToSubscriptionKeyParameterNamesContractResponseOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionKeyParameterNamesContractResponseOutput)
}

func (i SubscriptionKeyParameterNamesContractResponseArgs) ToSubscriptionKeyParameterNamesContractResponsePtrOutput() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return i.ToSubscriptionKeyParameterNamesContractResponsePtrOutputWithContext(context.Background())
}

func (i SubscriptionKeyParameterNamesContractResponseArgs) ToSubscriptionKeyParameterNamesContractResponsePtrOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionKeyParameterNamesContractResponseOutput).ToSubscriptionKeyParameterNamesContractResponsePtrOutputWithContext(ctx)
}









type SubscriptionKeyParameterNamesContractResponsePtrInput interface {
	pulumi.Input

	ToSubscriptionKeyParameterNamesContractResponsePtrOutput() SubscriptionKeyParameterNamesContractResponsePtrOutput
	ToSubscriptionKeyParameterNamesContractResponsePtrOutputWithContext(context.Context) SubscriptionKeyParameterNamesContractResponsePtrOutput
}

type subscriptionKeyParameterNamesContractResponsePtrType SubscriptionKeyParameterNamesContractResponseArgs

func SubscriptionKeyParameterNamesContractResponsePtr(v *SubscriptionKeyParameterNamesContractResponseArgs) SubscriptionKeyParameterNamesContractResponsePtrInput {
	return (*subscriptionKeyParameterNamesContractResponsePtrType)(v)
}

func (*subscriptionKeyParameterNamesContractResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionKeyParameterNamesContractResponse)(nil)).Elem()
}

func (i *subscriptionKeyParameterNamesContractResponsePtrType) ToSubscriptionKeyParameterNamesContractResponsePtrOutput() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return i.ToSubscriptionKeyParameterNamesContractResponsePtrOutputWithContext(context.Background())
}

func (i *subscriptionKeyParameterNamesContractResponsePtrType) ToSubscriptionKeyParameterNamesContractResponsePtrOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
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

func (o SubscriptionKeyParameterNamesContractResponseOutput) ToSubscriptionKeyParameterNamesContractResponsePtrOutput() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ToSubscriptionKeyParameterNamesContractResponsePtrOutputWithContext(context.Background())
}

func (o SubscriptionKeyParameterNamesContractResponseOutput) ToSubscriptionKeyParameterNamesContractResponsePtrOutputWithContext(ctx context.Context) SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionKeyParameterNamesContractResponse) *SubscriptionKeyParameterNamesContractResponse {
		return &v
	}).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
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





type TokenBodyParameterContractResponseInput interface {
	pulumi.Input

	ToTokenBodyParameterContractResponseOutput() TokenBodyParameterContractResponseOutput
	ToTokenBodyParameterContractResponseOutputWithContext(context.Context) TokenBodyParameterContractResponseOutput
}

type TokenBodyParameterContractResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TokenBodyParameterContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenBodyParameterContractResponse)(nil)).Elem()
}

func (i TokenBodyParameterContractResponseArgs) ToTokenBodyParameterContractResponseOutput() TokenBodyParameterContractResponseOutput {
	return i.ToTokenBodyParameterContractResponseOutputWithContext(context.Background())
}

func (i TokenBodyParameterContractResponseArgs) ToTokenBodyParameterContractResponseOutputWithContext(ctx context.Context) TokenBodyParameterContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenBodyParameterContractResponseOutput)
}





type TokenBodyParameterContractResponseArrayInput interface {
	pulumi.Input

	ToTokenBodyParameterContractResponseArrayOutput() TokenBodyParameterContractResponseArrayOutput
	ToTokenBodyParameterContractResponseArrayOutputWithContext(context.Context) TokenBodyParameterContractResponseArrayOutput
}

type TokenBodyParameterContractResponseArray []TokenBodyParameterContractResponseInput

func (TokenBodyParameterContractResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenBodyParameterContractResponse)(nil)).Elem()
}

func (i TokenBodyParameterContractResponseArray) ToTokenBodyParameterContractResponseArrayOutput() TokenBodyParameterContractResponseArrayOutput {
	return i.ToTokenBodyParameterContractResponseArrayOutputWithContext(context.Background())
}

func (i TokenBodyParameterContractResponseArray) ToTokenBodyParameterContractResponseArrayOutputWithContext(ctx context.Context) TokenBodyParameterContractResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenBodyParameterContractResponseArrayOutput)
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

type UserIdentityContractResponse struct {
	Id       *string `pulumi:"id"`
	Provider *string `pulumi:"provider"`
}





type UserIdentityContractResponseInput interface {
	pulumi.Input

	ToUserIdentityContractResponseOutput() UserIdentityContractResponseOutput
	ToUserIdentityContractResponseOutputWithContext(context.Context) UserIdentityContractResponseOutput
}

type UserIdentityContractResponseArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Provider pulumi.StringPtrInput `pulumi:"provider"`
}

func (UserIdentityContractResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityContractResponse)(nil)).Elem()
}

func (i UserIdentityContractResponseArgs) ToUserIdentityContractResponseOutput() UserIdentityContractResponseOutput {
	return i.ToUserIdentityContractResponseOutputWithContext(context.Background())
}

func (i UserIdentityContractResponseArgs) ToUserIdentityContractResponseOutputWithContext(ctx context.Context) UserIdentityContractResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityContractResponseOutput)
}





type UserIdentityContractResponseArrayInput interface {
	pulumi.Input

	ToUserIdentityContractResponseArrayOutput() UserIdentityContractResponseArrayOutput
	ToUserIdentityContractResponseArrayOutputWithContext(context.Context) UserIdentityContractResponseArrayOutput
}

type UserIdentityContractResponseArray []UserIdentityContractResponseInput

func (UserIdentityContractResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserIdentityContractResponse)(nil)).Elem()
}

func (i UserIdentityContractResponseArray) ToUserIdentityContractResponseArrayOutput() UserIdentityContractResponseArrayOutput {
	return i.ToUserIdentityContractResponseArrayOutputWithContext(context.Background())
}

func (i UserIdentityContractResponseArray) ToUserIdentityContractResponseArrayOutputWithContext(ctx context.Context) UserIdentityContractResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityContractResponseArrayOutput)
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





type VirtualNetworkConfigurationResponseInput interface {
	pulumi.Input

	ToVirtualNetworkConfigurationResponseOutput() VirtualNetworkConfigurationResponseOutput
	ToVirtualNetworkConfigurationResponseOutputWithContext(context.Context) VirtualNetworkConfigurationResponseOutput
}

type VirtualNetworkConfigurationResponseArgs struct {
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





type X509CertificateNameResponseInput interface {
	pulumi.Input

	ToX509CertificateNameResponseOutput() X509CertificateNameResponseOutput
	ToX509CertificateNameResponseOutputWithContext(context.Context) X509CertificateNameResponseOutput
}

type X509CertificateNameResponseArgs struct {
	IssuerCertificateThumbprint pulumi.StringPtrInput `pulumi:"issuerCertificateThumbprint"`
	Name                        pulumi.StringPtrInput `pulumi:"name"`
}

func (X509CertificateNameResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X509CertificateNameResponse)(nil)).Elem()
}

func (i X509CertificateNameResponseArgs) ToX509CertificateNameResponseOutput() X509CertificateNameResponseOutput {
	return i.ToX509CertificateNameResponseOutputWithContext(context.Background())
}

func (i X509CertificateNameResponseArgs) ToX509CertificateNameResponseOutputWithContext(ctx context.Context) X509CertificateNameResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X509CertificateNameResponseOutput)
}





type X509CertificateNameResponseArrayInput interface {
	pulumi.Input

	ToX509CertificateNameResponseArrayOutput() X509CertificateNameResponseArrayOutput
	ToX509CertificateNameResponseArrayOutputWithContext(context.Context) X509CertificateNameResponseArrayOutput
}

type X509CertificateNameResponseArray []X509CertificateNameResponseInput

func (X509CertificateNameResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X509CertificateNameResponse)(nil)).Elem()
}

func (i X509CertificateNameResponseArray) ToX509CertificateNameResponseArrayOutput() X509CertificateNameResponseArrayOutput {
	return i.ToX509CertificateNameResponseArrayOutputWithContext(context.Background())
}

func (i X509CertificateNameResponseArray) ToX509CertificateNameResponseArrayOutputWithContext(ctx context.Context) X509CertificateNameResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X509CertificateNameResponseArrayOutput)
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
	pulumi.RegisterOutputType(ApiCreateOrUpdatePropertiesWsdlSelectorOutput{})
	pulumi.RegisterOutputType(ApiCreateOrUpdatePropertiesWsdlSelectorPtrOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceIdentityOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApiManagementServiceSkuPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiVersionSetContractOutput{})
	pulumi.RegisterOutputType(ApiVersionSetContractPtrOutput{})
	pulumi.RegisterOutputType(ApiVersionSetContractResponseOutput{})
	pulumi.RegisterOutputType(ApiVersionSetContractResponsePtrOutput{})
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
	pulumi.RegisterOutputType(BackendPropertiesResponsePtrOutput{})
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
	pulumi.RegisterOutputType(CertificateConfigurationOutput{})
	pulumi.RegisterOutputType(CertificateConfigurationArrayOutput{})
	pulumi.RegisterOutputType(CertificateConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CertificateConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(CertificateInformationResponseOutput{})
	pulumi.RegisterOutputType(EmailTemplateParametersContractPropertiesOutput{})
	pulumi.RegisterOutputType(EmailTemplateParametersContractPropertiesArrayOutput{})
	pulumi.RegisterOutputType(EmailTemplateParametersContractPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EmailTemplateParametersContractPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(GroupContractResponseOutput{})
	pulumi.RegisterOutputType(GroupContractResponseArrayOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationArrayOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationResponseOutput{})
	pulumi.RegisterOutputType(HostnameConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(LoggerSamplingContractOutput{})
	pulumi.RegisterOutputType(LoggerSamplingContractPtrOutput{})
	pulumi.RegisterOutputType(LoggerSamplingContractResponseOutput{})
	pulumi.RegisterOutputType(LoggerSamplingContractResponsePtrOutput{})
	pulumi.RegisterOutputType(OAuth2AuthenticationSettingsContractOutput{})
	pulumi.RegisterOutputType(OAuth2AuthenticationSettingsContractPtrOutput{})
	pulumi.RegisterOutputType(OAuth2AuthenticationSettingsContractResponseOutput{})
	pulumi.RegisterOutputType(OAuth2AuthenticationSettingsContractResponsePtrOutput{})
	pulumi.RegisterOutputType(ParameterContractOutput{})
	pulumi.RegisterOutputType(ParameterContractArrayOutput{})
	pulumi.RegisterOutputType(ParameterContractResponseOutput{})
	pulumi.RegisterOutputType(ParameterContractResponseArrayOutput{})
	pulumi.RegisterOutputType(RepresentationContractOutput{})
	pulumi.RegisterOutputType(RepresentationContractArrayOutput{})
	pulumi.RegisterOutputType(RepresentationContractResponseOutput{})
	pulumi.RegisterOutputType(RepresentationContractResponseArrayOutput{})
	pulumi.RegisterOutputType(RequestContractOutput{})
	pulumi.RegisterOutputType(RequestContractPtrOutput{})
	pulumi.RegisterOutputType(RequestContractResponseOutput{})
	pulumi.RegisterOutputType(RequestContractResponsePtrOutput{})
	pulumi.RegisterOutputType(ResponseContractOutput{})
	pulumi.RegisterOutputType(ResponseContractArrayOutput{})
	pulumi.RegisterOutputType(ResponseContractResponseOutput{})
	pulumi.RegisterOutputType(ResponseContractResponseArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionKeyParameterNamesContractOutput{})
	pulumi.RegisterOutputType(SubscriptionKeyParameterNamesContractPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionKeyParameterNamesContractResponseOutput{})
	pulumi.RegisterOutputType(SubscriptionKeyParameterNamesContractResponsePtrOutput{})
	pulumi.RegisterOutputType(TokenBodyParameterContractOutput{})
	pulumi.RegisterOutputType(TokenBodyParameterContractArrayOutput{})
	pulumi.RegisterOutputType(TokenBodyParameterContractResponseOutput{})
	pulumi.RegisterOutputType(TokenBodyParameterContractResponseArrayOutput{})
	pulumi.RegisterOutputType(UserIdentityContractResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityContractResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(X509CertificateNameOutput{})
	pulumi.RegisterOutputType(X509CertificateNameArrayOutput{})
	pulumi.RegisterOutputType(X509CertificateNameResponseOutput{})
	pulumi.RegisterOutputType(X509CertificateNameResponseArrayOutput{})
}
