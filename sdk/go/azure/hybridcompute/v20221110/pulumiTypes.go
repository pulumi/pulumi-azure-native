


package v20221110

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgentConfigurationResponse struct {
	ConfigMode                string                           `pulumi:"configMode"`
	ExtensionsAllowList       []ConfigurationExtensionResponse `pulumi:"extensionsAllowList"`
	ExtensionsBlockList       []ConfigurationExtensionResponse `pulumi:"extensionsBlockList"`
	ExtensionsEnabled         string                           `pulumi:"extensionsEnabled"`
	GuestConfigurationEnabled string                           `pulumi:"guestConfigurationEnabled"`
	IncomingConnectionsPorts  []string                         `pulumi:"incomingConnectionsPorts"`
	ProxyBypass               []string                         `pulumi:"proxyBypass"`
	ProxyUrl                  string                           `pulumi:"proxyUrl"`
}

type AgentConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AgentConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentConfigurationResponse)(nil)).Elem()
}

func (o AgentConfigurationResponseOutput) ToAgentConfigurationResponseOutput() AgentConfigurationResponseOutput {
	return o
}

func (o AgentConfigurationResponseOutput) ToAgentConfigurationResponseOutputWithContext(ctx context.Context) AgentConfigurationResponseOutput {
	return o
}

func (o AgentConfigurationResponseOutput) ConfigMode() pulumi.StringOutput {
	return o.ApplyT(func(v AgentConfigurationResponse) string { return v.ConfigMode }).(pulumi.StringOutput)
}

func (o AgentConfigurationResponseOutput) ExtensionsAllowList() ConfigurationExtensionResponseArrayOutput {
	return o.ApplyT(func(v AgentConfigurationResponse) []ConfigurationExtensionResponse { return v.ExtensionsAllowList }).(ConfigurationExtensionResponseArrayOutput)
}

func (o AgentConfigurationResponseOutput) ExtensionsBlockList() ConfigurationExtensionResponseArrayOutput {
	return o.ApplyT(func(v AgentConfigurationResponse) []ConfigurationExtensionResponse { return v.ExtensionsBlockList }).(ConfigurationExtensionResponseArrayOutput)
}

func (o AgentConfigurationResponseOutput) ExtensionsEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v AgentConfigurationResponse) string { return v.ExtensionsEnabled }).(pulumi.StringOutput)
}

func (o AgentConfigurationResponseOutput) GuestConfigurationEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v AgentConfigurationResponse) string { return v.GuestConfigurationEnabled }).(pulumi.StringOutput)
}

func (o AgentConfigurationResponseOutput) IncomingConnectionsPorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AgentConfigurationResponse) []string { return v.IncomingConnectionsPorts }).(pulumi.StringArrayOutput)
}

func (o AgentConfigurationResponseOutput) ProxyBypass() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AgentConfigurationResponse) []string { return v.ProxyBypass }).(pulumi.StringArrayOutput)
}

func (o AgentConfigurationResponseOutput) ProxyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v AgentConfigurationResponse) string { return v.ProxyUrl }).(pulumi.StringOutput)
}

type CloudMetadataResponse struct {
	Provider string `pulumi:"provider"`
}

type CloudMetadataResponseOutput struct{ *pulumi.OutputState }

func (CloudMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudMetadataResponse)(nil)).Elem()
}

func (o CloudMetadataResponseOutput) ToCloudMetadataResponseOutput() CloudMetadataResponseOutput {
	return o
}

func (o CloudMetadataResponseOutput) ToCloudMetadataResponseOutputWithContext(ctx context.Context) CloudMetadataResponseOutput {
	return o
}

func (o CloudMetadataResponseOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v CloudMetadataResponse) string { return v.Provider }).(pulumi.StringOutput)
}

type CloudMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudMetadataResponse)(nil)).Elem()
}

func (o CloudMetadataResponsePtrOutput) ToCloudMetadataResponsePtrOutput() CloudMetadataResponsePtrOutput {
	return o
}

func (o CloudMetadataResponsePtrOutput) ToCloudMetadataResponsePtrOutputWithContext(ctx context.Context) CloudMetadataResponsePtrOutput {
	return o
}

func (o CloudMetadataResponsePtrOutput) Elem() CloudMetadataResponseOutput {
	return o.ApplyT(func(v *CloudMetadataResponse) CloudMetadataResponse {
		if v != nil {
			return *v
		}
		var ret CloudMetadataResponse
		return ret
	}).(CloudMetadataResponseOutput)
}

func (o CloudMetadataResponsePtrOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Provider
	}).(pulumi.StringPtrOutput)
}

type ConfigurationExtensionResponse struct {
	Publisher string `pulumi:"publisher"`
	Type      string `pulumi:"type"`
}

type ConfigurationExtensionResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationExtensionResponse)(nil)).Elem()
}

func (o ConfigurationExtensionResponseOutput) ToConfigurationExtensionResponseOutput() ConfigurationExtensionResponseOutput {
	return o
}

func (o ConfigurationExtensionResponseOutput) ToConfigurationExtensionResponseOutputWithContext(ctx context.Context) ConfigurationExtensionResponseOutput {
	return o
}

func (o ConfigurationExtensionResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationExtensionResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o ConfigurationExtensionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationExtensionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ConfigurationExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationExtensionResponse)(nil)).Elem()
}

func (o ConfigurationExtensionResponseArrayOutput) ToConfigurationExtensionResponseArrayOutput() ConfigurationExtensionResponseArrayOutput {
	return o
}

func (o ConfigurationExtensionResponseArrayOutput) ToConfigurationExtensionResponseArrayOutputWithContext(ctx context.Context) ConfigurationExtensionResponseArrayOutput {
	return o
}

func (o ConfigurationExtensionResponseArrayOutput) Index(i pulumi.IntInput) ConfigurationExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationExtensionResponse {
		return vs[0].([]ConfigurationExtensionResponse)[vs[1].(int)]
	}).(ConfigurationExtensionResponseOutput)
}

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

type ErrorAdditionalInfoResponseOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o ErrorAdditionalInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ErrorAdditionalInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) Index(i pulumi.IntInput) ErrorAdditionalInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorAdditionalInfoResponse {
		return vs[0].([]ErrorAdditionalInfoResponse)[vs[1].(int)]
	}).(ErrorAdditionalInfoResponseOutput)
}

type ErrorDetailResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorDetailResponse         `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type HybridComputePrivateLinkScopeProperties struct {
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}





type HybridComputePrivateLinkScopePropertiesInput interface {
	pulumi.Input

	ToHybridComputePrivateLinkScopePropertiesOutput() HybridComputePrivateLinkScopePropertiesOutput
	ToHybridComputePrivateLinkScopePropertiesOutputWithContext(context.Context) HybridComputePrivateLinkScopePropertiesOutput
}

type HybridComputePrivateLinkScopePropertiesArgs struct {
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (HybridComputePrivateLinkScopePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputePrivateLinkScopeProperties)(nil)).Elem()
}

func (i HybridComputePrivateLinkScopePropertiesArgs) ToHybridComputePrivateLinkScopePropertiesOutput() HybridComputePrivateLinkScopePropertiesOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesOutputWithContext(context.Background())
}

func (i HybridComputePrivateLinkScopePropertiesArgs) ToHybridComputePrivateLinkScopePropertiesOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesOutput)
}

func (i HybridComputePrivateLinkScopePropertiesArgs) ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(context.Background())
}

func (i HybridComputePrivateLinkScopePropertiesArgs) ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesOutput).ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx)
}









type HybridComputePrivateLinkScopePropertiesPtrInput interface {
	pulumi.Input

	ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput
	ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput
}

type hybridComputePrivateLinkScopePropertiesPtrType HybridComputePrivateLinkScopePropertiesArgs

func HybridComputePrivateLinkScopePropertiesPtr(v *HybridComputePrivateLinkScopePropertiesArgs) HybridComputePrivateLinkScopePropertiesPtrInput {
	return (*hybridComputePrivateLinkScopePropertiesPtrType)(v)
}

func (*hybridComputePrivateLinkScopePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputePrivateLinkScopeProperties)(nil)).Elem()
}

func (i *hybridComputePrivateLinkScopePropertiesPtrType) ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(context.Background())
}

func (i *hybridComputePrivateLinkScopePropertiesPtrType) ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesPtrOutput)
}

type HybridComputePrivateLinkScopePropertiesOutput struct{ *pulumi.OutputState }

func (HybridComputePrivateLinkScopePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputePrivateLinkScopeProperties)(nil)).Elem()
}

func (o HybridComputePrivateLinkScopePropertiesOutput) ToHybridComputePrivateLinkScopePropertiesOutput() HybridComputePrivateLinkScopePropertiesOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesOutput) ToHybridComputePrivateLinkScopePropertiesOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesOutput) ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput {
	return o.ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(context.Background())
}

func (o HybridComputePrivateLinkScopePropertiesOutput) ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HybridComputePrivateLinkScopeProperties) *HybridComputePrivateLinkScopeProperties {
		return &v
	}).(HybridComputePrivateLinkScopePropertiesPtrOutput)
}

func (o HybridComputePrivateLinkScopePropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputePrivateLinkScopeProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type HybridComputePrivateLinkScopePropertiesPtrOutput struct{ *pulumi.OutputState }

func (HybridComputePrivateLinkScopePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputePrivateLinkScopeProperties)(nil)).Elem()
}

func (o HybridComputePrivateLinkScopePropertiesPtrOutput) ToHybridComputePrivateLinkScopePropertiesPtrOutput() HybridComputePrivateLinkScopePropertiesPtrOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesPtrOutput) ToHybridComputePrivateLinkScopePropertiesPtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesPtrOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesPtrOutput) Elem() HybridComputePrivateLinkScopePropertiesOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopeProperties) HybridComputePrivateLinkScopeProperties {
		if v != nil {
			return *v
		}
		var ret HybridComputePrivateLinkScopeProperties
		return ret
	}).(HybridComputePrivateLinkScopePropertiesOutput)
}

func (o HybridComputePrivateLinkScopePropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopeProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type HybridComputePrivateLinkScopePropertiesResponse struct {
	PrivateEndpointConnections []PrivateEndpointConnectionDataModelResponse `pulumi:"privateEndpointConnections"`
	PrivateLinkScopeId         string                                       `pulumi:"privateLinkScopeId"`
	ProvisioningState          string                                       `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                      `pulumi:"publicNetworkAccess"`
}

type HybridComputePrivateLinkScopePropertiesResponseOutput struct{ *pulumi.OutputState }

func (HybridComputePrivateLinkScopePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputePrivateLinkScopePropertiesResponse)(nil)).Elem()
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ToHybridComputePrivateLinkScopePropertiesResponseOutput() HybridComputePrivateLinkScopePropertiesResponseOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ToHybridComputePrivateLinkScopePropertiesResponseOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponseOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionDataModelResponseArrayOutput {
	return o.ApplyT(func(v HybridComputePrivateLinkScopePropertiesResponse) []PrivateEndpointConnectionDataModelResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionDataModelResponseArrayOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) PrivateLinkScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputePrivateLinkScopePropertiesResponse) string { return v.PrivateLinkScopeId }).(pulumi.StringOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputePrivateLinkScopePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputePrivateLinkScopePropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type LocationData struct {
	City            *string `pulumi:"city"`
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	District        *string `pulumi:"district"`
	Name            string  `pulumi:"name"`
}





type LocationDataInput interface {
	pulumi.Input

	ToLocationDataOutput() LocationDataOutput
	ToLocationDataOutputWithContext(context.Context) LocationDataOutput
}

type LocationDataArgs struct {
	City            pulumi.StringPtrInput `pulumi:"city"`
	CountryOrRegion pulumi.StringPtrInput `pulumi:"countryOrRegion"`
	District        pulumi.StringPtrInput `pulumi:"district"`
	Name            pulumi.StringInput    `pulumi:"name"`
}

func (LocationDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationData)(nil)).Elem()
}

func (i LocationDataArgs) ToLocationDataOutput() LocationDataOutput {
	return i.ToLocationDataOutputWithContext(context.Background())
}

func (i LocationDataArgs) ToLocationDataOutputWithContext(ctx context.Context) LocationDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataOutput)
}

func (i LocationDataArgs) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return i.ToLocationDataPtrOutputWithContext(context.Background())
}

func (i LocationDataArgs) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataOutput).ToLocationDataPtrOutputWithContext(ctx)
}









type LocationDataPtrInput interface {
	pulumi.Input

	ToLocationDataPtrOutput() LocationDataPtrOutput
	ToLocationDataPtrOutputWithContext(context.Context) LocationDataPtrOutput
}

type locationDataPtrType LocationDataArgs

func LocationDataPtr(v *LocationDataArgs) LocationDataPtrInput {
	return (*locationDataPtrType)(v)
}

func (*locationDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationData)(nil)).Elem()
}

func (i *locationDataPtrType) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return i.ToLocationDataPtrOutputWithContext(context.Background())
}

func (i *locationDataPtrType) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataPtrOutput)
}

type LocationDataOutput struct{ *pulumi.OutputState }

func (LocationDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationData)(nil)).Elem()
}

func (o LocationDataOutput) ToLocationDataOutput() LocationDataOutput {
	return o
}

func (o LocationDataOutput) ToLocationDataOutputWithContext(ctx context.Context) LocationDataOutput {
	return o
}

func (o LocationDataOutput) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return o.ToLocationDataPtrOutputWithContext(context.Background())
}

func (o LocationDataOutput) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocationData) *LocationData {
		return &v
	}).(LocationDataPtrOutput)
}

func (o LocationDataOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o LocationDataOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o LocationDataOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.District }).(pulumi.StringPtrOutput)
}

func (o LocationDataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LocationData) string { return v.Name }).(pulumi.StringOutput)
}

type LocationDataPtrOutput struct{ *pulumi.OutputState }

func (LocationDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationData)(nil)).Elem()
}

func (o LocationDataPtrOutput) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return o
}

func (o LocationDataPtrOutput) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return o
}

func (o LocationDataPtrOutput) Elem() LocationDataOutput {
	return o.ApplyT(func(v *LocationData) LocationData {
		if v != nil {
			return *v
		}
		var ret LocationData
		return ret
	}).(LocationDataOutput)
}

func (o LocationDataPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataPtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataPtrOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.District
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type LocationDataResponse struct {
	City            *string `pulumi:"city"`
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	District        *string `pulumi:"district"`
	Name            string  `pulumi:"name"`
}

type LocationDataResponseOutput struct{ *pulumi.OutputState }

func (LocationDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationDataResponse)(nil)).Elem()
}

func (o LocationDataResponseOutput) ToLocationDataResponseOutput() LocationDataResponseOutput {
	return o
}

func (o LocationDataResponseOutput) ToLocationDataResponseOutputWithContext(ctx context.Context) LocationDataResponseOutput {
	return o
}

func (o LocationDataResponseOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationDataResponse) *string { return v.City }).(pulumi.StringPtrOutput)
}

func (o LocationDataResponseOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationDataResponse) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

func (o LocationDataResponseOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationDataResponse) *string { return v.District }).(pulumi.StringPtrOutput)
}

func (o LocationDataResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LocationDataResponse) string { return v.Name }).(pulumi.StringOutput)
}

type LocationDataResponsePtrOutput struct{ *pulumi.OutputState }

func (LocationDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationDataResponse)(nil)).Elem()
}

func (o LocationDataResponsePtrOutput) ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput {
	return o
}

func (o LocationDataResponsePtrOutput) ToLocationDataResponsePtrOutputWithContext(ctx context.Context) LocationDataResponsePtrOutput {
	return o
}

func (o LocationDataResponsePtrOutput) Elem() LocationDataResponseOutput {
	return o.ApplyT(func(v *LocationDataResponse) LocationDataResponse {
		if v != nil {
			return *v
		}
		var ret LocationDataResponse
		return ret
	}).(LocationDataResponseOutput)
}

func (o LocationDataResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataResponsePtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataResponsePtrOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.District
	}).(pulumi.StringPtrOutput)
}

func (o LocationDataResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceView struct {
	Name               *string                             `pulumi:"name"`
	Status             *MachineExtensionInstanceViewStatus `pulumi:"status"`
	Type               *string                             `pulumi:"type"`
	TypeHandlerVersion *string                             `pulumi:"typeHandlerVersion"`
}





type MachineExtensionInstanceViewInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewOutput() MachineExtensionInstanceViewOutput
	ToMachineExtensionInstanceViewOutputWithContext(context.Context) MachineExtensionInstanceViewOutput
}

type MachineExtensionInstanceViewArgs struct {
	Name               pulumi.StringPtrInput                      `pulumi:"name"`
	Status             MachineExtensionInstanceViewStatusPtrInput `pulumi:"status"`
	Type               pulumi.StringPtrInput                      `pulumi:"type"`
	TypeHandlerVersion pulumi.StringPtrInput                      `pulumi:"typeHandlerVersion"`
}

func (MachineExtensionInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceView)(nil)).Elem()
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewOutput() MachineExtensionInstanceViewOutput {
	return i.ToMachineExtensionInstanceViewOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewOutputWithContext(ctx context.Context) MachineExtensionInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewOutput)
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput {
	return i.ToMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewOutput).ToMachineExtensionInstanceViewPtrOutputWithContext(ctx)
}









type MachineExtensionInstanceViewPtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput
	ToMachineExtensionInstanceViewPtrOutputWithContext(context.Context) MachineExtensionInstanceViewPtrOutput
}

type machineExtensionInstanceViewPtrType MachineExtensionInstanceViewArgs

func MachineExtensionInstanceViewPtr(v *MachineExtensionInstanceViewArgs) MachineExtensionInstanceViewPtrInput {
	return (*machineExtensionInstanceViewPtrType)(v)
}

func (*machineExtensionInstanceViewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceView)(nil)).Elem()
}

func (i *machineExtensionInstanceViewPtrType) ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput {
	return i.ToMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewPtrType) ToMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewPtrOutput)
}





type MachineExtensionInstanceViewArrayInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewArrayOutput() MachineExtensionInstanceViewArrayOutput
	ToMachineExtensionInstanceViewArrayOutputWithContext(context.Context) MachineExtensionInstanceViewArrayOutput
}

type MachineExtensionInstanceViewArray []MachineExtensionInstanceViewInput

func (MachineExtensionInstanceViewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceView)(nil)).Elem()
}

func (i MachineExtensionInstanceViewArray) ToMachineExtensionInstanceViewArrayOutput() MachineExtensionInstanceViewArrayOutput {
	return i.ToMachineExtensionInstanceViewArrayOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewArray) ToMachineExtensionInstanceViewArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewArrayOutput)
}

type MachineExtensionInstanceViewOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceView)(nil)).Elem()
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewOutput() MachineExtensionInstanceViewOutput {
	return o
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewOutputWithContext(ctx context.Context) MachineExtensionInstanceViewOutput {
	return o
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput {
	return o.ToMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionInstanceView) *MachineExtensionInstanceView {
		return &v
	}).(MachineExtensionInstanceViewPtrOutput)
}

func (o MachineExtensionInstanceViewOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewOutput) Status() MachineExtensionInstanceViewStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *MachineExtensionInstanceViewStatus { return v.Status }).(MachineExtensionInstanceViewStatusPtrOutput)
}

func (o MachineExtensionInstanceViewOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceView)(nil)).Elem()
}

func (o MachineExtensionInstanceViewPtrOutput) ToMachineExtensionInstanceViewPtrOutput() MachineExtensionInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewPtrOutput) ToMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewPtrOutput) Elem() MachineExtensionInstanceViewOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) MachineExtensionInstanceView {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceView
		return ret
	}).(MachineExtensionInstanceViewOutput)
}

func (o MachineExtensionInstanceViewPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewPtrOutput) Status() MachineExtensionInstanceViewStatusPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) *MachineExtensionInstanceViewStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(MachineExtensionInstanceViewStatusPtrOutput)
}

func (o MachineExtensionInstanceViewPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewArrayOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceView)(nil)).Elem()
}

func (o MachineExtensionInstanceViewArrayOutput) ToMachineExtensionInstanceViewArrayOutput() MachineExtensionInstanceViewArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewArrayOutput) ToMachineExtensionInstanceViewArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewArrayOutput) Index(i pulumi.IntInput) MachineExtensionInstanceViewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineExtensionInstanceView {
		return vs[0].([]MachineExtensionInstanceView)[vs[1].(int)]
	}).(MachineExtensionInstanceViewOutput)
}

type MachineExtensionInstanceViewResponse struct {
	Name               *string                                     `pulumi:"name"`
	Status             *MachineExtensionInstanceViewResponseStatus `pulumi:"status"`
	Type               *string                                     `pulumi:"type"`
	TypeHandlerVersion *string                                     `pulumi:"typeHandlerVersion"`
}

type MachineExtensionInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponseOutput() MachineExtensionInstanceViewResponseOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponseOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *MachineExtensionInstanceViewResponseStatus {
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponsePtrOutput) ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponsePtrOutput) ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponsePtrOutput) Elem() MachineExtensionInstanceViewResponseOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) MachineExtensionInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceViewResponse
		return ret
	}).(MachineExtensionInstanceViewResponseOutput)
}

func (o MachineExtensionInstanceViewResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponsePtrOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) *MachineExtensionInstanceViewResponseStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionInstanceViewResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponsePtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseArrayOutput) ToMachineExtensionInstanceViewResponseArrayOutput() MachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseArrayOutput) ToMachineExtensionInstanceViewResponseArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) MachineExtensionInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineExtensionInstanceViewResponse {
		return vs[0].([]MachineExtensionInstanceViewResponse)[vs[1].(int)]
	}).(MachineExtensionInstanceViewResponseOutput)
}

type MachineExtensionInstanceViewResponseStatus struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}

type MachineExtensionInstanceViewResponseStatusOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewResponseStatusPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Elem() MachineExtensionInstanceViewResponseStatusOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) MachineExtensionInstanceViewResponseStatus {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceViewResponseStatus
		return ret
	}).(MachineExtensionInstanceViewResponseStatusOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewStatus struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}





type MachineExtensionInstanceViewStatusInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewStatusOutput() MachineExtensionInstanceViewStatusOutput
	ToMachineExtensionInstanceViewStatusOutputWithContext(context.Context) MachineExtensionInstanceViewStatusOutput
}

type MachineExtensionInstanceViewStatusArgs struct {
	Code          pulumi.StringPtrInput `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput `pulumi:"displayStatus"`
	Level         pulumi.StringPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput `pulumi:"message"`
	Time          pulumi.StringPtrInput `pulumi:"time"`
}

func (MachineExtensionInstanceViewStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewStatus)(nil)).Elem()
}

func (i MachineExtensionInstanceViewStatusArgs) ToMachineExtensionInstanceViewStatusOutput() MachineExtensionInstanceViewStatusOutput {
	return i.ToMachineExtensionInstanceViewStatusOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewStatusArgs) ToMachineExtensionInstanceViewStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewStatusOutput)
}

func (i MachineExtensionInstanceViewStatusArgs) ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewStatusPtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewStatusArgs) ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewStatusOutput).ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx)
}









type MachineExtensionInstanceViewStatusPtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput
	ToMachineExtensionInstanceViewStatusPtrOutputWithContext(context.Context) MachineExtensionInstanceViewStatusPtrOutput
}

type machineExtensionInstanceViewStatusPtrType MachineExtensionInstanceViewStatusArgs

func MachineExtensionInstanceViewStatusPtr(v *MachineExtensionInstanceViewStatusArgs) MachineExtensionInstanceViewStatusPtrInput {
	return (*machineExtensionInstanceViewStatusPtrType)(v)
}

func (*machineExtensionInstanceViewStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewStatus)(nil)).Elem()
}

func (i *machineExtensionInstanceViewStatusPtrType) ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewStatusPtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewStatusPtrType) ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewStatusPtrOutput)
}

type MachineExtensionInstanceViewStatusOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewStatusOutput) ToMachineExtensionInstanceViewStatusOutput() MachineExtensionInstanceViewStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewStatusOutput) ToMachineExtensionInstanceViewStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewStatusOutput) ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput {
	return o.ToMachineExtensionInstanceViewStatusPtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewStatusOutput) ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionInstanceViewStatus) *MachineExtensionInstanceViewStatus {
		return &v
	}).(MachineExtensionInstanceViewStatusPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewStatusPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewStatusPtrOutput) ToMachineExtensionInstanceViewStatusPtrOutput() MachineExtensionInstanceViewStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewStatusPtrOutput) ToMachineExtensionInstanceViewStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Elem() MachineExtensionInstanceViewStatusOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) MachineExtensionInstanceViewStatus {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceViewStatus
		return ret
	}).(MachineExtensionInstanceViewStatusOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewStatusPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionResponse struct {
	AutoUpgradeMinorVersion *bool                                 `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade  *bool                                 `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag          *string                               `pulumi:"forceUpdateTag"`
	Id                      string                                `pulumi:"id"`
	InstanceView            *MachineExtensionInstanceViewResponse `pulumi:"instanceView"`
	Location                string                                `pulumi:"location"`
	Name                    string                                `pulumi:"name"`
	ProtectedSettings       interface{}                           `pulumi:"protectedSettings"`
	ProvisioningState       string                                `pulumi:"provisioningState"`
	Publisher               *string                               `pulumi:"publisher"`
	Settings                interface{}                           `pulumi:"settings"`
	SystemData              SystemDataResponse                    `pulumi:"systemData"`
	Tags                    map[string]string                     `pulumi:"tags"`
	Type                    string                                `pulumi:"type"`
	TypeHandlerVersion      *string                               `pulumi:"typeHandlerVersion"`
}

type MachineExtensionResponseOutput struct{ *pulumi.OutputState }

func (MachineExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionResponse)(nil)).Elem()
}

func (o MachineExtensionResponseOutput) ToMachineExtensionResponseOutput() MachineExtensionResponseOutput {
	return o
}

func (o MachineExtensionResponseOutput) ToMachineExtensionResponseOutputWithContext(ctx context.Context) MachineExtensionResponseOutput {
	return o
}

func (o MachineExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MachineExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o MachineExtensionResponseOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MachineExtensionResponse) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o MachineExtensionResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o MachineExtensionResponseOutput) InstanceView() MachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v MachineExtensionResponse) *MachineExtensionInstanceViewResponse { return v.InstanceView }).(MachineExtensionInstanceViewResponsePtrOutput)
}

func (o MachineExtensionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o MachineExtensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MachineExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v MachineExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o MachineExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MachineExtensionResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v MachineExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o MachineExtensionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v MachineExtensionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MachineExtensionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v MachineExtensionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MachineExtensionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o MachineExtensionResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type MachineExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (MachineExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionResponse)(nil)).Elem()
}

func (o MachineExtensionResponseArrayOutput) ToMachineExtensionResponseArrayOutput() MachineExtensionResponseArrayOutput {
	return o
}

func (o MachineExtensionResponseArrayOutput) ToMachineExtensionResponseArrayOutputWithContext(ctx context.Context) MachineExtensionResponseArrayOutput {
	return o
}

func (o MachineExtensionResponseArrayOutput) Index(i pulumi.IntInput) MachineExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineExtensionResponse {
		return vs[0].([]MachineExtensionResponse)[vs[1].(int)]
	}).(MachineExtensionResponseOutput)
}

type OSProfile struct {
	LinuxConfiguration   *OSProfileLinuxConfiguration   `pulumi:"linuxConfiguration"`
	WindowsConfiguration *OSProfileWindowsConfiguration `pulumi:"windowsConfiguration"`
}





type OSProfileInput interface {
	pulumi.Input

	ToOSProfileOutput() OSProfileOutput
	ToOSProfileOutputWithContext(context.Context) OSProfileOutput
}

type OSProfileArgs struct {
	LinuxConfiguration   OSProfileLinuxConfigurationPtrInput   `pulumi:"linuxConfiguration"`
	WindowsConfiguration OSProfileWindowsConfigurationPtrInput `pulumi:"windowsConfiguration"`
}

func (OSProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (i OSProfileArgs) ToOSProfileOutput() OSProfileOutput {
	return i.ToOSProfileOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput)
}

func (i OSProfileArgs) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput).ToOSProfilePtrOutputWithContext(ctx)
}









type OSProfilePtrInput interface {
	pulumi.Input

	ToOSProfilePtrOutput() OSProfilePtrOutput
	ToOSProfilePtrOutputWithContext(context.Context) OSProfilePtrOutput
}

type osprofilePtrType OSProfileArgs

func OSProfilePtr(v *OSProfileArgs) OSProfilePtrInput {
	return (*osprofilePtrType)(v)
}

func (*osprofilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (i *osprofilePtrType) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i *osprofilePtrType) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfilePtrOutput)
}

type OSProfileOutput struct{ *pulumi.OutputState }

func (OSProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (o OSProfileOutput) ToOSProfileOutput() OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o.ToOSProfilePtrOutputWithContext(context.Background())
}

func (o OSProfileOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSProfile) *OSProfile {
		return &v
	}).(OSProfilePtrOutput)
}

func (o OSProfileOutput) LinuxConfiguration() OSProfileLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfile) *OSProfileLinuxConfiguration { return v.LinuxConfiguration }).(OSProfileLinuxConfigurationPtrOutput)
}

func (o OSProfileOutput) WindowsConfiguration() OSProfileWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfile) *OSProfileWindowsConfiguration { return v.WindowsConfiguration }).(OSProfileWindowsConfigurationPtrOutput)
}

type OSProfilePtrOutput struct{ *pulumi.OutputState }

func (OSProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) Elem() OSProfileOutput {
	return o.ApplyT(func(v *OSProfile) OSProfile {
		if v != nil {
			return *v
		}
		var ret OSProfile
		return ret
	}).(OSProfileOutput)
}

func (o OSProfilePtrOutput) LinuxConfiguration() OSProfileLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfile) *OSProfileLinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(OSProfileLinuxConfigurationPtrOutput)
}

func (o OSProfilePtrOutput) WindowsConfiguration() OSProfileWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfile) *OSProfileWindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(OSProfileWindowsConfigurationPtrOutput)
}

type OSProfileLinuxConfiguration struct {
	AssessmentMode *string `pulumi:"assessmentMode"`
	PatchMode      *string `pulumi:"patchMode"`
}





type OSProfileLinuxConfigurationInput interface {
	pulumi.Input

	ToOSProfileLinuxConfigurationOutput() OSProfileLinuxConfigurationOutput
	ToOSProfileLinuxConfigurationOutputWithContext(context.Context) OSProfileLinuxConfigurationOutput
}

type OSProfileLinuxConfigurationArgs struct {
	AssessmentMode pulumi.StringPtrInput `pulumi:"assessmentMode"`
	PatchMode      pulumi.StringPtrInput `pulumi:"patchMode"`
}

func (OSProfileLinuxConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileLinuxConfiguration)(nil)).Elem()
}

func (i OSProfileLinuxConfigurationArgs) ToOSProfileLinuxConfigurationOutput() OSProfileLinuxConfigurationOutput {
	return i.ToOSProfileLinuxConfigurationOutputWithContext(context.Background())
}

func (i OSProfileLinuxConfigurationArgs) ToOSProfileLinuxConfigurationOutputWithContext(ctx context.Context) OSProfileLinuxConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileLinuxConfigurationOutput)
}

func (i OSProfileLinuxConfigurationArgs) ToOSProfileLinuxConfigurationPtrOutput() OSProfileLinuxConfigurationPtrOutput {
	return i.ToOSProfileLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i OSProfileLinuxConfigurationArgs) ToOSProfileLinuxConfigurationPtrOutputWithContext(ctx context.Context) OSProfileLinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileLinuxConfigurationOutput).ToOSProfileLinuxConfigurationPtrOutputWithContext(ctx)
}









type OSProfileLinuxConfigurationPtrInput interface {
	pulumi.Input

	ToOSProfileLinuxConfigurationPtrOutput() OSProfileLinuxConfigurationPtrOutput
	ToOSProfileLinuxConfigurationPtrOutputWithContext(context.Context) OSProfileLinuxConfigurationPtrOutput
}

type osprofileLinuxConfigurationPtrType OSProfileLinuxConfigurationArgs

func OSProfileLinuxConfigurationPtr(v *OSProfileLinuxConfigurationArgs) OSProfileLinuxConfigurationPtrInput {
	return (*osprofileLinuxConfigurationPtrType)(v)
}

func (*osprofileLinuxConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileLinuxConfiguration)(nil)).Elem()
}

func (i *osprofileLinuxConfigurationPtrType) ToOSProfileLinuxConfigurationPtrOutput() OSProfileLinuxConfigurationPtrOutput {
	return i.ToOSProfileLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i *osprofileLinuxConfigurationPtrType) ToOSProfileLinuxConfigurationPtrOutputWithContext(ctx context.Context) OSProfileLinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileLinuxConfigurationPtrOutput)
}

type OSProfileLinuxConfigurationOutput struct{ *pulumi.OutputState }

func (OSProfileLinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileLinuxConfiguration)(nil)).Elem()
}

func (o OSProfileLinuxConfigurationOutput) ToOSProfileLinuxConfigurationOutput() OSProfileLinuxConfigurationOutput {
	return o
}

func (o OSProfileLinuxConfigurationOutput) ToOSProfileLinuxConfigurationOutputWithContext(ctx context.Context) OSProfileLinuxConfigurationOutput {
	return o
}

func (o OSProfileLinuxConfigurationOutput) ToOSProfileLinuxConfigurationPtrOutput() OSProfileLinuxConfigurationPtrOutput {
	return o.ToOSProfileLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (o OSProfileLinuxConfigurationOutput) ToOSProfileLinuxConfigurationPtrOutputWithContext(ctx context.Context) OSProfileLinuxConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSProfileLinuxConfiguration) *OSProfileLinuxConfiguration {
		return &v
	}).(OSProfileLinuxConfigurationPtrOutput)
}

func (o OSProfileLinuxConfigurationOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileLinuxConfiguration) *string { return v.AssessmentMode }).(pulumi.StringPtrOutput)
}

func (o OSProfileLinuxConfigurationOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileLinuxConfiguration) *string { return v.PatchMode }).(pulumi.StringPtrOutput)
}

type OSProfileLinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (OSProfileLinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileLinuxConfiguration)(nil)).Elem()
}

func (o OSProfileLinuxConfigurationPtrOutput) ToOSProfileLinuxConfigurationPtrOutput() OSProfileLinuxConfigurationPtrOutput {
	return o
}

func (o OSProfileLinuxConfigurationPtrOutput) ToOSProfileLinuxConfigurationPtrOutputWithContext(ctx context.Context) OSProfileLinuxConfigurationPtrOutput {
	return o
}

func (o OSProfileLinuxConfigurationPtrOutput) Elem() OSProfileLinuxConfigurationOutput {
	return o.ApplyT(func(v *OSProfileLinuxConfiguration) OSProfileLinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret OSProfileLinuxConfiguration
		return ret
	}).(OSProfileLinuxConfigurationOutput)
}

func (o OSProfileLinuxConfigurationPtrOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileLinuxConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentMode
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileLinuxConfigurationPtrOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileLinuxConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PatchMode
	}).(pulumi.StringPtrOutput)
}

type OSProfileResponse struct {
	ComputerName         string                                 `pulumi:"computerName"`
	LinuxConfiguration   *OSProfileResponseLinuxConfiguration   `pulumi:"linuxConfiguration"`
	WindowsConfiguration *OSProfileResponseWindowsConfiguration `pulumi:"windowsConfiguration"`
}

type OSProfileResponseOutput struct{ *pulumi.OutputState }

func (OSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutput() OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutputWithContext(ctx context.Context) OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v OSProfileResponse) string { return v.ComputerName }).(pulumi.StringOutput)
}

func (o OSProfileResponseOutput) LinuxConfiguration() OSProfileResponseLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *OSProfileResponseLinuxConfiguration { return v.LinuxConfiguration }).(OSProfileResponseLinuxConfigurationPtrOutput)
}

func (o OSProfileResponseOutput) WindowsConfiguration() OSProfileResponseWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *OSProfileResponseWindowsConfiguration { return v.WindowsConfiguration }).(OSProfileResponseWindowsConfigurationPtrOutput)
}

type OSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) Elem() OSProfileResponseOutput {
	return o.ApplyT(func(v *OSProfileResponse) OSProfileResponse {
		if v != nil {
			return *v
		}
		var ret OSProfileResponse
		return ret
	}).(OSProfileResponseOutput)
}

func (o OSProfileResponsePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) LinuxConfiguration() OSProfileResponseLinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *OSProfileResponseLinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(OSProfileResponseLinuxConfigurationPtrOutput)
}

func (o OSProfileResponsePtrOutput) WindowsConfiguration() OSProfileResponseWindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *OSProfileResponseWindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(OSProfileResponseWindowsConfigurationPtrOutput)
}

type OSProfileResponseLinuxConfiguration struct {
	AssessmentMode *string `pulumi:"assessmentMode"`
	PatchMode      *string `pulumi:"patchMode"`
}

type OSProfileResponseLinuxConfigurationOutput struct{ *pulumi.OutputState }

func (OSProfileResponseLinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponseLinuxConfiguration)(nil)).Elem()
}

func (o OSProfileResponseLinuxConfigurationOutput) ToOSProfileResponseLinuxConfigurationOutput() OSProfileResponseLinuxConfigurationOutput {
	return o
}

func (o OSProfileResponseLinuxConfigurationOutput) ToOSProfileResponseLinuxConfigurationOutputWithContext(ctx context.Context) OSProfileResponseLinuxConfigurationOutput {
	return o
}

func (o OSProfileResponseLinuxConfigurationOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponseLinuxConfiguration) *string { return v.AssessmentMode }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseLinuxConfigurationOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponseLinuxConfiguration) *string { return v.PatchMode }).(pulumi.StringPtrOutput)
}

type OSProfileResponseLinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (OSProfileResponseLinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponseLinuxConfiguration)(nil)).Elem()
}

func (o OSProfileResponseLinuxConfigurationPtrOutput) ToOSProfileResponseLinuxConfigurationPtrOutput() OSProfileResponseLinuxConfigurationPtrOutput {
	return o
}

func (o OSProfileResponseLinuxConfigurationPtrOutput) ToOSProfileResponseLinuxConfigurationPtrOutputWithContext(ctx context.Context) OSProfileResponseLinuxConfigurationPtrOutput {
	return o
}

func (o OSProfileResponseLinuxConfigurationPtrOutput) Elem() OSProfileResponseLinuxConfigurationOutput {
	return o.ApplyT(func(v *OSProfileResponseLinuxConfiguration) OSProfileResponseLinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret OSProfileResponseLinuxConfiguration
		return ret
	}).(OSProfileResponseLinuxConfigurationOutput)
}

func (o OSProfileResponseLinuxConfigurationPtrOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponseLinuxConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentMode
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseLinuxConfigurationPtrOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponseLinuxConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PatchMode
	}).(pulumi.StringPtrOutput)
}

type OSProfileResponseWindowsConfiguration struct {
	AssessmentMode *string `pulumi:"assessmentMode"`
	PatchMode      *string `pulumi:"patchMode"`
}

type OSProfileResponseWindowsConfigurationOutput struct{ *pulumi.OutputState }

func (OSProfileResponseWindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponseWindowsConfiguration)(nil)).Elem()
}

func (o OSProfileResponseWindowsConfigurationOutput) ToOSProfileResponseWindowsConfigurationOutput() OSProfileResponseWindowsConfigurationOutput {
	return o
}

func (o OSProfileResponseWindowsConfigurationOutput) ToOSProfileResponseWindowsConfigurationOutputWithContext(ctx context.Context) OSProfileResponseWindowsConfigurationOutput {
	return o
}

func (o OSProfileResponseWindowsConfigurationOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponseWindowsConfiguration) *string { return v.AssessmentMode }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseWindowsConfigurationOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponseWindowsConfiguration) *string { return v.PatchMode }).(pulumi.StringPtrOutput)
}

type OSProfileResponseWindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (OSProfileResponseWindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponseWindowsConfiguration)(nil)).Elem()
}

func (o OSProfileResponseWindowsConfigurationPtrOutput) ToOSProfileResponseWindowsConfigurationPtrOutput() OSProfileResponseWindowsConfigurationPtrOutput {
	return o
}

func (o OSProfileResponseWindowsConfigurationPtrOutput) ToOSProfileResponseWindowsConfigurationPtrOutputWithContext(ctx context.Context) OSProfileResponseWindowsConfigurationPtrOutput {
	return o
}

func (o OSProfileResponseWindowsConfigurationPtrOutput) Elem() OSProfileResponseWindowsConfigurationOutput {
	return o.ApplyT(func(v *OSProfileResponseWindowsConfiguration) OSProfileResponseWindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret OSProfileResponseWindowsConfiguration
		return ret
	}).(OSProfileResponseWindowsConfigurationOutput)
}

func (o OSProfileResponseWindowsConfigurationPtrOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponseWindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentMode
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseWindowsConfigurationPtrOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponseWindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PatchMode
	}).(pulumi.StringPtrOutput)
}

type OSProfileWindowsConfiguration struct {
	AssessmentMode *string `pulumi:"assessmentMode"`
	PatchMode      *string `pulumi:"patchMode"`
}





type OSProfileWindowsConfigurationInput interface {
	pulumi.Input

	ToOSProfileWindowsConfigurationOutput() OSProfileWindowsConfigurationOutput
	ToOSProfileWindowsConfigurationOutputWithContext(context.Context) OSProfileWindowsConfigurationOutput
}

type OSProfileWindowsConfigurationArgs struct {
	AssessmentMode pulumi.StringPtrInput `pulumi:"assessmentMode"`
	PatchMode      pulumi.StringPtrInput `pulumi:"patchMode"`
}

func (OSProfileWindowsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileWindowsConfiguration)(nil)).Elem()
}

func (i OSProfileWindowsConfigurationArgs) ToOSProfileWindowsConfigurationOutput() OSProfileWindowsConfigurationOutput {
	return i.ToOSProfileWindowsConfigurationOutputWithContext(context.Background())
}

func (i OSProfileWindowsConfigurationArgs) ToOSProfileWindowsConfigurationOutputWithContext(ctx context.Context) OSProfileWindowsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileWindowsConfigurationOutput)
}

func (i OSProfileWindowsConfigurationArgs) ToOSProfileWindowsConfigurationPtrOutput() OSProfileWindowsConfigurationPtrOutput {
	return i.ToOSProfileWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i OSProfileWindowsConfigurationArgs) ToOSProfileWindowsConfigurationPtrOutputWithContext(ctx context.Context) OSProfileWindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileWindowsConfigurationOutput).ToOSProfileWindowsConfigurationPtrOutputWithContext(ctx)
}









type OSProfileWindowsConfigurationPtrInput interface {
	pulumi.Input

	ToOSProfileWindowsConfigurationPtrOutput() OSProfileWindowsConfigurationPtrOutput
	ToOSProfileWindowsConfigurationPtrOutputWithContext(context.Context) OSProfileWindowsConfigurationPtrOutput
}

type osprofileWindowsConfigurationPtrType OSProfileWindowsConfigurationArgs

func OSProfileWindowsConfigurationPtr(v *OSProfileWindowsConfigurationArgs) OSProfileWindowsConfigurationPtrInput {
	return (*osprofileWindowsConfigurationPtrType)(v)
}

func (*osprofileWindowsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileWindowsConfiguration)(nil)).Elem()
}

func (i *osprofileWindowsConfigurationPtrType) ToOSProfileWindowsConfigurationPtrOutput() OSProfileWindowsConfigurationPtrOutput {
	return i.ToOSProfileWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i *osprofileWindowsConfigurationPtrType) ToOSProfileWindowsConfigurationPtrOutputWithContext(ctx context.Context) OSProfileWindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileWindowsConfigurationPtrOutput)
}

type OSProfileWindowsConfigurationOutput struct{ *pulumi.OutputState }

func (OSProfileWindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileWindowsConfiguration)(nil)).Elem()
}

func (o OSProfileWindowsConfigurationOutput) ToOSProfileWindowsConfigurationOutput() OSProfileWindowsConfigurationOutput {
	return o
}

func (o OSProfileWindowsConfigurationOutput) ToOSProfileWindowsConfigurationOutputWithContext(ctx context.Context) OSProfileWindowsConfigurationOutput {
	return o
}

func (o OSProfileWindowsConfigurationOutput) ToOSProfileWindowsConfigurationPtrOutput() OSProfileWindowsConfigurationPtrOutput {
	return o.ToOSProfileWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (o OSProfileWindowsConfigurationOutput) ToOSProfileWindowsConfigurationPtrOutputWithContext(ctx context.Context) OSProfileWindowsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSProfileWindowsConfiguration) *OSProfileWindowsConfiguration {
		return &v
	}).(OSProfileWindowsConfigurationPtrOutput)
}

func (o OSProfileWindowsConfigurationOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileWindowsConfiguration) *string { return v.AssessmentMode }).(pulumi.StringPtrOutput)
}

func (o OSProfileWindowsConfigurationOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileWindowsConfiguration) *string { return v.PatchMode }).(pulumi.StringPtrOutput)
}

type OSProfileWindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (OSProfileWindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileWindowsConfiguration)(nil)).Elem()
}

func (o OSProfileWindowsConfigurationPtrOutput) ToOSProfileWindowsConfigurationPtrOutput() OSProfileWindowsConfigurationPtrOutput {
	return o
}

func (o OSProfileWindowsConfigurationPtrOutput) ToOSProfileWindowsConfigurationPtrOutputWithContext(ctx context.Context) OSProfileWindowsConfigurationPtrOutput {
	return o
}

func (o OSProfileWindowsConfigurationPtrOutput) Elem() OSProfileWindowsConfigurationOutput {
	return o.ApplyT(func(v *OSProfileWindowsConfiguration) OSProfileWindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret OSProfileWindowsConfiguration
		return ret
	}).(OSProfileWindowsConfigurationOutput)
}

func (o OSProfileWindowsConfigurationPtrOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileWindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentMode
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileWindowsConfigurationPtrOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileWindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PatchMode
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionDataModelResponse struct {
	Id         string                                       `pulumi:"id"`
	Name       string                                       `pulumi:"name"`
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                       `pulumi:"type"`
}

type PrivateEndpointConnectionDataModelResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionDataModelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionDataModelResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionDataModelResponseOutput) ToPrivateEndpointConnectionDataModelResponseOutput() PrivateEndpointConnectionDataModelResponseOutput {
	return o
}

func (o PrivateEndpointConnectionDataModelResponseOutput) ToPrivateEndpointConnectionDataModelResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionDataModelResponseOutput {
	return o
}

func (o PrivateEndpointConnectionDataModelResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionDataModelResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionDataModelResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionDataModelResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionDataModelResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionDataModelResponse) *PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

func (o PrivateEndpointConnectionDataModelResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionDataModelResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionDataModelResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionDataModelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionDataModelResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionDataModelResponseArrayOutput) ToPrivateEndpointConnectionDataModelResponseArrayOutput() PrivateEndpointConnectionDataModelResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionDataModelResponseArrayOutput) ToPrivateEndpointConnectionDataModelResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionDataModelResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionDataModelResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionDataModelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionDataModelResponse {
		return vs[0].([]PrivateEndpointConnectionDataModelResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionDataModelResponseOutput)
}

type PrivateEndpointConnectionProperties struct {
	PrivateEndpoint                   *PrivateEndpointProperty                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	PrivateEndpoint                   PrivateEndpointPropertyPtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyPtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput).ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput
	ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPtrOutput
}

type privateEndpointConnectionPropertiesPtrType PrivateEndpointConnectionPropertiesArgs

func PrivateEndpointConnectionPropertiesPtr(v *PrivateEndpointConnectionPropertiesArgs) PrivateEndpointConnectionPropertiesPtrInput {
	return (*privateEndpointConnectionPropertiesPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPtrOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionProperties {
		return &v
	}).(PrivateEndpointConnectionPropertiesPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateEndpoint() PrivateEndpointPropertyPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateEndpointProperty { return v.PrivateEndpoint }).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionStateProperty {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateEndpointConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) Elem() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) PrivateEndpointConnectionProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionProperties
		return ret
	}).(PrivateEndpointConnectionPropertiesOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateEndpoint() PrivateEndpointPropertyPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateEndpointProperty {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionStateProperty {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	GroupIds                          []string                                           `pulumi:"groupIds"`
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                             `pulumi:"provisioningState"`
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointPropertyResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) PrivateEndpointConnectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponse
		return ret
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointPropertyResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointProperty struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput
	ToPrivateEndpointPropertyOutputWithContext(context.Context) PrivateEndpointPropertyOutput
}

type PrivateEndpointPropertyArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return i.ToPrivateEndpointPropertyOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput)
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput).ToPrivateEndpointPropertyPtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput
	ToPrivateEndpointPropertyPtrOutputWithContext(context.Context) PrivateEndpointPropertyPtrOutput
}

type privateEndpointPropertyPtrType PrivateEndpointPropertyArgs

func PrivateEndpointPropertyPtr(v *PrivateEndpointPropertyArgs) PrivateEndpointPropertyPtrInput {
	return (*privateEndpointPropertyPtrType)(v)
}

func (*privateEndpointPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyPtrOutput)
}

type PrivateEndpointPropertyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointProperty) *PrivateEndpointProperty {
		return &v
	}).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointPropertyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointProperty) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) Elem() PrivateEndpointPropertyOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) PrivateEndpointProperty {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointProperty
		return ret
	}).(PrivateEndpointPropertyOutput)
}

func (o PrivateEndpointPropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) Elem() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) PrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointPropertyResponse
		return ret
	}).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateProperty struct {
	Description string `pulumi:"description"`
	Status      string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput
	ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyOutput
}

type PrivateLinkServiceConnectionStatePropertyArgs struct {
	Description pulumi.StringInput `pulumi:"description"`
	Status      pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput).ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput
	ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput
}

type privateLinkServiceConnectionStatePropertyPtrType PrivateLinkServiceConnectionStatePropertyArgs

func PrivateLinkServiceConnectionStatePropertyPtr(v *PrivateLinkServiceConnectionStatePropertyArgs) PrivateLinkServiceConnectionStatePropertyPtrInput {
	return (*privateLinkServiceConnectionStatePropertyPtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateProperty) *PrivateLinkServiceConnectionStateProperty {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) PrivateLinkServiceConnectionStateProperty {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateProperty
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type ServiceStatus struct {
	StartupType *string `pulumi:"startupType"`
	Status      *string `pulumi:"status"`
}





type ServiceStatusInput interface {
	pulumi.Input

	ToServiceStatusOutput() ServiceStatusOutput
	ToServiceStatusOutputWithContext(context.Context) ServiceStatusOutput
}

type ServiceStatusArgs struct {
	StartupType pulumi.StringPtrInput `pulumi:"startupType"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (ServiceStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceStatus)(nil)).Elem()
}

func (i ServiceStatusArgs) ToServiceStatusOutput() ServiceStatusOutput {
	return i.ToServiceStatusOutputWithContext(context.Background())
}

func (i ServiceStatusArgs) ToServiceStatusOutputWithContext(ctx context.Context) ServiceStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceStatusOutput)
}

func (i ServiceStatusArgs) ToServiceStatusPtrOutput() ServiceStatusPtrOutput {
	return i.ToServiceStatusPtrOutputWithContext(context.Background())
}

func (i ServiceStatusArgs) ToServiceStatusPtrOutputWithContext(ctx context.Context) ServiceStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceStatusOutput).ToServiceStatusPtrOutputWithContext(ctx)
}









type ServiceStatusPtrInput interface {
	pulumi.Input

	ToServiceStatusPtrOutput() ServiceStatusPtrOutput
	ToServiceStatusPtrOutputWithContext(context.Context) ServiceStatusPtrOutput
}

type serviceStatusPtrType ServiceStatusArgs

func ServiceStatusPtr(v *ServiceStatusArgs) ServiceStatusPtrInput {
	return (*serviceStatusPtrType)(v)
}

func (*serviceStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceStatus)(nil)).Elem()
}

func (i *serviceStatusPtrType) ToServiceStatusPtrOutput() ServiceStatusPtrOutput {
	return i.ToServiceStatusPtrOutputWithContext(context.Background())
}

func (i *serviceStatusPtrType) ToServiceStatusPtrOutputWithContext(ctx context.Context) ServiceStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceStatusPtrOutput)
}

type ServiceStatusOutput struct{ *pulumi.OutputState }

func (ServiceStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceStatus)(nil)).Elem()
}

func (o ServiceStatusOutput) ToServiceStatusOutput() ServiceStatusOutput {
	return o
}

func (o ServiceStatusOutput) ToServiceStatusOutputWithContext(ctx context.Context) ServiceStatusOutput {
	return o
}

func (o ServiceStatusOutput) ToServiceStatusPtrOutput() ServiceStatusPtrOutput {
	return o.ToServiceStatusPtrOutputWithContext(context.Background())
}

func (o ServiceStatusOutput) ToServiceStatusPtrOutputWithContext(ctx context.Context) ServiceStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceStatus) *ServiceStatus {
		return &v
	}).(ServiceStatusPtrOutput)
}

func (o ServiceStatusOutput) StartupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceStatus) *string { return v.StartupType }).(pulumi.StringPtrOutput)
}

func (o ServiceStatusOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceStatus) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ServiceStatusPtrOutput struct{ *pulumi.OutputState }

func (ServiceStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceStatus)(nil)).Elem()
}

func (o ServiceStatusPtrOutput) ToServiceStatusPtrOutput() ServiceStatusPtrOutput {
	return o
}

func (o ServiceStatusPtrOutput) ToServiceStatusPtrOutputWithContext(ctx context.Context) ServiceStatusPtrOutput {
	return o
}

func (o ServiceStatusPtrOutput) Elem() ServiceStatusOutput {
	return o.ApplyT(func(v *ServiceStatus) ServiceStatus {
		if v != nil {
			return *v
		}
		var ret ServiceStatus
		return ret
	}).(ServiceStatusOutput)
}

func (o ServiceStatusPtrOutput) StartupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceStatus) *string {
		if v == nil {
			return nil
		}
		return v.StartupType
	}).(pulumi.StringPtrOutput)
}

func (o ServiceStatusPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceStatus) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ServiceStatusResponse struct {
	StartupType *string `pulumi:"startupType"`
	Status      *string `pulumi:"status"`
}

type ServiceStatusResponseOutput struct{ *pulumi.OutputState }

func (ServiceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceStatusResponse)(nil)).Elem()
}

func (o ServiceStatusResponseOutput) ToServiceStatusResponseOutput() ServiceStatusResponseOutput {
	return o
}

func (o ServiceStatusResponseOutput) ToServiceStatusResponseOutputWithContext(ctx context.Context) ServiceStatusResponseOutput {
	return o
}

func (o ServiceStatusResponseOutput) StartupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceStatusResponse) *string { return v.StartupType }).(pulumi.StringPtrOutput)
}

func (o ServiceStatusResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceStatusResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ServiceStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceStatusResponse)(nil)).Elem()
}

func (o ServiceStatusResponsePtrOutput) ToServiceStatusResponsePtrOutput() ServiceStatusResponsePtrOutput {
	return o
}

func (o ServiceStatusResponsePtrOutput) ToServiceStatusResponsePtrOutputWithContext(ctx context.Context) ServiceStatusResponsePtrOutput {
	return o
}

func (o ServiceStatusResponsePtrOutput) Elem() ServiceStatusResponseOutput {
	return o.ApplyT(func(v *ServiceStatusResponse) ServiceStatusResponse {
		if v != nil {
			return *v
		}
		var ret ServiceStatusResponse
		return ret
	}).(ServiceStatusResponseOutput)
}

func (o ServiceStatusResponsePtrOutput) StartupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartupType
	}).(pulumi.StringPtrOutput)
}

func (o ServiceStatusResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ServiceStatuses struct {
	ExtensionService          *ServiceStatus `pulumi:"extensionService"`
	GuestConfigurationService *ServiceStatus `pulumi:"guestConfigurationService"`
}





type ServiceStatusesInput interface {
	pulumi.Input

	ToServiceStatusesOutput() ServiceStatusesOutput
	ToServiceStatusesOutputWithContext(context.Context) ServiceStatusesOutput
}

type ServiceStatusesArgs struct {
	ExtensionService          ServiceStatusPtrInput `pulumi:"extensionService"`
	GuestConfigurationService ServiceStatusPtrInput `pulumi:"guestConfigurationService"`
}

func (ServiceStatusesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceStatuses)(nil)).Elem()
}

func (i ServiceStatusesArgs) ToServiceStatusesOutput() ServiceStatusesOutput {
	return i.ToServiceStatusesOutputWithContext(context.Background())
}

func (i ServiceStatusesArgs) ToServiceStatusesOutputWithContext(ctx context.Context) ServiceStatusesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceStatusesOutput)
}

func (i ServiceStatusesArgs) ToServiceStatusesPtrOutput() ServiceStatusesPtrOutput {
	return i.ToServiceStatusesPtrOutputWithContext(context.Background())
}

func (i ServiceStatusesArgs) ToServiceStatusesPtrOutputWithContext(ctx context.Context) ServiceStatusesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceStatusesOutput).ToServiceStatusesPtrOutputWithContext(ctx)
}









type ServiceStatusesPtrInput interface {
	pulumi.Input

	ToServiceStatusesPtrOutput() ServiceStatusesPtrOutput
	ToServiceStatusesPtrOutputWithContext(context.Context) ServiceStatusesPtrOutput
}

type serviceStatusesPtrType ServiceStatusesArgs

func ServiceStatusesPtr(v *ServiceStatusesArgs) ServiceStatusesPtrInput {
	return (*serviceStatusesPtrType)(v)
}

func (*serviceStatusesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceStatuses)(nil)).Elem()
}

func (i *serviceStatusesPtrType) ToServiceStatusesPtrOutput() ServiceStatusesPtrOutput {
	return i.ToServiceStatusesPtrOutputWithContext(context.Background())
}

func (i *serviceStatusesPtrType) ToServiceStatusesPtrOutputWithContext(ctx context.Context) ServiceStatusesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceStatusesPtrOutput)
}

type ServiceStatusesOutput struct{ *pulumi.OutputState }

func (ServiceStatusesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceStatuses)(nil)).Elem()
}

func (o ServiceStatusesOutput) ToServiceStatusesOutput() ServiceStatusesOutput {
	return o
}

func (o ServiceStatusesOutput) ToServiceStatusesOutputWithContext(ctx context.Context) ServiceStatusesOutput {
	return o
}

func (o ServiceStatusesOutput) ToServiceStatusesPtrOutput() ServiceStatusesPtrOutput {
	return o.ToServiceStatusesPtrOutputWithContext(context.Background())
}

func (o ServiceStatusesOutput) ToServiceStatusesPtrOutputWithContext(ctx context.Context) ServiceStatusesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceStatuses) *ServiceStatuses {
		return &v
	}).(ServiceStatusesPtrOutput)
}

func (o ServiceStatusesOutput) ExtensionService() ServiceStatusPtrOutput {
	return o.ApplyT(func(v ServiceStatuses) *ServiceStatus { return v.ExtensionService }).(ServiceStatusPtrOutput)
}

func (o ServiceStatusesOutput) GuestConfigurationService() ServiceStatusPtrOutput {
	return o.ApplyT(func(v ServiceStatuses) *ServiceStatus { return v.GuestConfigurationService }).(ServiceStatusPtrOutput)
}

type ServiceStatusesPtrOutput struct{ *pulumi.OutputState }

func (ServiceStatusesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceStatuses)(nil)).Elem()
}

func (o ServiceStatusesPtrOutput) ToServiceStatusesPtrOutput() ServiceStatusesPtrOutput {
	return o
}

func (o ServiceStatusesPtrOutput) ToServiceStatusesPtrOutputWithContext(ctx context.Context) ServiceStatusesPtrOutput {
	return o
}

func (o ServiceStatusesPtrOutput) Elem() ServiceStatusesOutput {
	return o.ApplyT(func(v *ServiceStatuses) ServiceStatuses {
		if v != nil {
			return *v
		}
		var ret ServiceStatuses
		return ret
	}).(ServiceStatusesOutput)
}

func (o ServiceStatusesPtrOutput) ExtensionService() ServiceStatusPtrOutput {
	return o.ApplyT(func(v *ServiceStatuses) *ServiceStatus {
		if v == nil {
			return nil
		}
		return v.ExtensionService
	}).(ServiceStatusPtrOutput)
}

func (o ServiceStatusesPtrOutput) GuestConfigurationService() ServiceStatusPtrOutput {
	return o.ApplyT(func(v *ServiceStatuses) *ServiceStatus {
		if v == nil {
			return nil
		}
		return v.GuestConfigurationService
	}).(ServiceStatusPtrOutput)
}

type ServiceStatusesResponse struct {
	ExtensionService          *ServiceStatusResponse `pulumi:"extensionService"`
	GuestConfigurationService *ServiceStatusResponse `pulumi:"guestConfigurationService"`
}

type ServiceStatusesResponseOutput struct{ *pulumi.OutputState }

func (ServiceStatusesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceStatusesResponse)(nil)).Elem()
}

func (o ServiceStatusesResponseOutput) ToServiceStatusesResponseOutput() ServiceStatusesResponseOutput {
	return o
}

func (o ServiceStatusesResponseOutput) ToServiceStatusesResponseOutputWithContext(ctx context.Context) ServiceStatusesResponseOutput {
	return o
}

func (o ServiceStatusesResponseOutput) ExtensionService() ServiceStatusResponsePtrOutput {
	return o.ApplyT(func(v ServiceStatusesResponse) *ServiceStatusResponse { return v.ExtensionService }).(ServiceStatusResponsePtrOutput)
}

func (o ServiceStatusesResponseOutput) GuestConfigurationService() ServiceStatusResponsePtrOutput {
	return o.ApplyT(func(v ServiceStatusesResponse) *ServiceStatusResponse { return v.GuestConfigurationService }).(ServiceStatusResponsePtrOutput)
}

type ServiceStatusesResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceStatusesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceStatusesResponse)(nil)).Elem()
}

func (o ServiceStatusesResponsePtrOutput) ToServiceStatusesResponsePtrOutput() ServiceStatusesResponsePtrOutput {
	return o
}

func (o ServiceStatusesResponsePtrOutput) ToServiceStatusesResponsePtrOutputWithContext(ctx context.Context) ServiceStatusesResponsePtrOutput {
	return o
}

func (o ServiceStatusesResponsePtrOutput) Elem() ServiceStatusesResponseOutput {
	return o.ApplyT(func(v *ServiceStatusesResponse) ServiceStatusesResponse {
		if v != nil {
			return *v
		}
		var ret ServiceStatusesResponse
		return ret
	}).(ServiceStatusesResponseOutput)
}

func (o ServiceStatusesResponsePtrOutput) ExtensionService() ServiceStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServiceStatusesResponse) *ServiceStatusResponse {
		if v == nil {
			return nil
		}
		return v.ExtensionService
	}).(ServiceStatusResponsePtrOutput)
}

func (o ServiceStatusesResponsePtrOutput) GuestConfigurationService() ServiceStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServiceStatusesResponse) *ServiceStatusResponse {
		if v == nil {
			return nil
		}
		return v.GuestConfigurationService
	}).(ServiceStatusResponsePtrOutput)
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

func init() {
	pulumi.RegisterOutputType(AgentConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CloudMetadataResponseOutput{})
	pulumi.RegisterOutputType(CloudMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationExtensionResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesPtrOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(LocationDataOutput{})
	pulumi.RegisterOutputType(LocationDataPtrOutput{})
	pulumi.RegisterOutputType(LocationDataResponseOutput{})
	pulumi.RegisterOutputType(LocationDataResponsePtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewArrayOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewStatusOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewStatusPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionResponseOutput{})
	pulumi.RegisterOutputType(MachineExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(OSProfileOutput{})
	pulumi.RegisterOutputType(OSProfilePtrOutput{})
	pulumi.RegisterOutputType(OSProfileLinuxConfigurationOutput{})
	pulumi.RegisterOutputType(OSProfileLinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseLinuxConfigurationOutput{})
	pulumi.RegisterOutputType(OSProfileResponseLinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseWindowsConfigurationOutput{})
	pulumi.RegisterOutputType(OSProfileResponseWindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(OSProfileWindowsConfigurationOutput{})
	pulumi.RegisterOutputType(OSProfileWindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionDataModelResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionDataModelResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceStatusOutput{})
	pulumi.RegisterOutputType(ServiceStatusPtrOutput{})
	pulumi.RegisterOutputType(ServiceStatusResponseOutput{})
	pulumi.RegisterOutputType(ServiceStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceStatusesOutput{})
	pulumi.RegisterOutputType(ServiceStatusesPtrOutput{})
	pulumi.RegisterOutputType(ServiceStatusesResponseOutput{})
	pulumi.RegisterOutputType(ServiceStatusesResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
