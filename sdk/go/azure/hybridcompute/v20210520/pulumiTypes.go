


package v20210520

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	PrivateLinkScopeId  string  `pulumi:"privateLinkScopeId"`
	ProvisioningState   string  `pulumi:"provisioningState"`
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
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

type MachineExtensionProperties struct {
	AutoUpgradeMinorVersion *bool                         `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string                       `pulumi:"forceUpdateTag"`
	InstanceView            *MachineExtensionInstanceView `pulumi:"instanceView"`
	ProtectedSettings       interface{}                   `pulumi:"protectedSettings"`
	Publisher               *string                       `pulumi:"publisher"`
	Settings                interface{}                   `pulumi:"settings"`
	Type                    *string                       `pulumi:"type"`
	TypeHandlerVersion      *string                       `pulumi:"typeHandlerVersion"`
}





type MachineExtensionPropertiesInput interface {
	pulumi.Input

	ToMachineExtensionPropertiesOutput() MachineExtensionPropertiesOutput
	ToMachineExtensionPropertiesOutputWithContext(context.Context) MachineExtensionPropertiesOutput
}

type MachineExtensionPropertiesArgs struct {
	AutoUpgradeMinorVersion pulumi.BoolPtrInput                  `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          pulumi.StringPtrInput                `pulumi:"forceUpdateTag"`
	InstanceView            MachineExtensionInstanceViewPtrInput `pulumi:"instanceView"`
	ProtectedSettings       pulumi.Input                         `pulumi:"protectedSettings"`
	Publisher               pulumi.StringPtrInput                `pulumi:"publisher"`
	Settings                pulumi.Input                         `pulumi:"settings"`
	Type                    pulumi.StringPtrInput                `pulumi:"type"`
	TypeHandlerVersion      pulumi.StringPtrInput                `pulumi:"typeHandlerVersion"`
}

func (MachineExtensionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionProperties)(nil)).Elem()
}

func (i MachineExtensionPropertiesArgs) ToMachineExtensionPropertiesOutput() MachineExtensionPropertiesOutput {
	return i.ToMachineExtensionPropertiesOutputWithContext(context.Background())
}

func (i MachineExtensionPropertiesArgs) ToMachineExtensionPropertiesOutputWithContext(ctx context.Context) MachineExtensionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionPropertiesOutput)
}

func (i MachineExtensionPropertiesArgs) ToMachineExtensionPropertiesPtrOutput() MachineExtensionPropertiesPtrOutput {
	return i.ToMachineExtensionPropertiesPtrOutputWithContext(context.Background())
}

func (i MachineExtensionPropertiesArgs) ToMachineExtensionPropertiesPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionPropertiesOutput).ToMachineExtensionPropertiesPtrOutputWithContext(ctx)
}









type MachineExtensionPropertiesPtrInput interface {
	pulumi.Input

	ToMachineExtensionPropertiesPtrOutput() MachineExtensionPropertiesPtrOutput
	ToMachineExtensionPropertiesPtrOutputWithContext(context.Context) MachineExtensionPropertiesPtrOutput
}

type machineExtensionPropertiesPtrType MachineExtensionPropertiesArgs

func MachineExtensionPropertiesPtr(v *MachineExtensionPropertiesArgs) MachineExtensionPropertiesPtrInput {
	return (*machineExtensionPropertiesPtrType)(v)
}

func (*machineExtensionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionProperties)(nil)).Elem()
}

func (i *machineExtensionPropertiesPtrType) ToMachineExtensionPropertiesPtrOutput() MachineExtensionPropertiesPtrOutput {
	return i.ToMachineExtensionPropertiesPtrOutputWithContext(context.Background())
}

func (i *machineExtensionPropertiesPtrType) ToMachineExtensionPropertiesPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionPropertiesPtrOutput)
}

type MachineExtensionPropertiesOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionProperties)(nil)).Elem()
}

func (o MachineExtensionPropertiesOutput) ToMachineExtensionPropertiesOutput() MachineExtensionPropertiesOutput {
	return o
}

func (o MachineExtensionPropertiesOutput) ToMachineExtensionPropertiesOutputWithContext(ctx context.Context) MachineExtensionPropertiesOutput {
	return o
}

func (o MachineExtensionPropertiesOutput) ToMachineExtensionPropertiesPtrOutput() MachineExtensionPropertiesPtrOutput {
	return o.ToMachineExtensionPropertiesPtrOutputWithContext(context.Background())
}

func (o MachineExtensionPropertiesOutput) ToMachineExtensionPropertiesPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionProperties) *MachineExtensionProperties {
		return &v
	}).(MachineExtensionPropertiesPtrOutput)
}

func (o MachineExtensionPropertiesOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MachineExtensionProperties) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o MachineExtensionPropertiesOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionProperties) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesOutput) InstanceView() MachineExtensionInstanceViewPtrOutput {
	return o.ApplyT(func(v MachineExtensionProperties) *MachineExtensionInstanceView { return v.InstanceView }).(MachineExtensionInstanceViewPtrOutput)
}

func (o MachineExtensionPropertiesOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v MachineExtensionProperties) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o MachineExtensionPropertiesOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionProperties) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v MachineExtensionProperties) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o MachineExtensionPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionProperties) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type MachineExtensionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionProperties)(nil)).Elem()
}

func (o MachineExtensionPropertiesPtrOutput) ToMachineExtensionPropertiesPtrOutput() MachineExtensionPropertiesPtrOutput {
	return o
}

func (o MachineExtensionPropertiesPtrOutput) ToMachineExtensionPropertiesPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesPtrOutput {
	return o
}

func (o MachineExtensionPropertiesPtrOutput) Elem() MachineExtensionPropertiesOutput {
	return o.ApplyT(func(v *MachineExtensionProperties) MachineExtensionProperties {
		if v != nil {
			return *v
		}
		var ret MachineExtensionProperties
		return ret
	}).(MachineExtensionPropertiesOutput)
}

func (o MachineExtensionPropertiesPtrOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MachineExtensionProperties) *bool {
		if v == nil {
			return nil
		}
		return v.AutoUpgradeMinorVersion
	}).(pulumi.BoolPtrOutput)
}

func (o MachineExtensionPropertiesPtrOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionProperties) *string {
		if v == nil {
			return nil
		}
		return v.ForceUpdateTag
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesPtrOutput) InstanceView() MachineExtensionInstanceViewPtrOutput {
	return o.ApplyT(func(v *MachineExtensionProperties) *MachineExtensionInstanceView {
		if v == nil {
			return nil
		}
		return v.InstanceView
	}).(MachineExtensionInstanceViewPtrOutput)
}

func (o MachineExtensionPropertiesPtrOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MachineExtensionProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.ProtectedSettings
	}).(pulumi.AnyOutput)
}

func (o MachineExtensionPropertiesPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesPtrOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MachineExtensionProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(pulumi.AnyOutput)
}

func (o MachineExtensionPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionProperties) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionPropertiesResponse struct {
	AutoUpgradeMinorVersion *bool                                 `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string                               `pulumi:"forceUpdateTag"`
	InstanceView            *MachineExtensionInstanceViewResponse `pulumi:"instanceView"`
	ProtectedSettings       interface{}                           `pulumi:"protectedSettings"`
	ProvisioningState       string                                `pulumi:"provisioningState"`
	Publisher               *string                               `pulumi:"publisher"`
	Settings                interface{}                           `pulumi:"settings"`
	Type                    *string                               `pulumi:"type"`
	TypeHandlerVersion      *string                               `pulumi:"typeHandlerVersion"`
}

type MachineExtensionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionPropertiesResponse)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponseOutput) ToMachineExtensionPropertiesResponseOutput() MachineExtensionPropertiesResponseOutput {
	return o
}

func (o MachineExtensionPropertiesResponseOutput) ToMachineExtensionPropertiesResponseOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseOutput {
	return o
}

func (o MachineExtensionPropertiesResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o MachineExtensionPropertiesResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseOutput) InstanceView() MachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponse) *MachineExtensionInstanceViewResponse {
		return v.InstanceView
	}).(MachineExtensionInstanceViewResponsePtrOutput)
}

func (o MachineExtensionPropertiesResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o MachineExtensionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MachineExtensionPropertiesResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o MachineExtensionPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type MachineProperties struct {
	ClientPublicKey            *string                        `pulumi:"clientPublicKey"`
	Extensions                 []MachineExtensionInstanceView `pulumi:"extensions"`
	LocationData               *LocationData                  `pulumi:"locationData"`
	ParentClusterResourceId    *string                        `pulumi:"parentClusterResourceId"`
	PrivateLinkScopeResourceId *string                        `pulumi:"privateLinkScopeResourceId"`
	VmId                       *string                        `pulumi:"vmId"`
}





type MachinePropertiesInput interface {
	pulumi.Input

	ToMachinePropertiesOutput() MachinePropertiesOutput
	ToMachinePropertiesOutputWithContext(context.Context) MachinePropertiesOutput
}

type MachinePropertiesArgs struct {
	ClientPublicKey            pulumi.StringPtrInput                  `pulumi:"clientPublicKey"`
	Extensions                 MachineExtensionInstanceViewArrayInput `pulumi:"extensions"`
	LocationData               LocationDataPtrInput                   `pulumi:"locationData"`
	ParentClusterResourceId    pulumi.StringPtrInput                  `pulumi:"parentClusterResourceId"`
	PrivateLinkScopeResourceId pulumi.StringPtrInput                  `pulumi:"privateLinkScopeResourceId"`
	VmId                       pulumi.StringPtrInput                  `pulumi:"vmId"`
}

func (MachinePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineProperties)(nil)).Elem()
}

func (i MachinePropertiesArgs) ToMachinePropertiesOutput() MachinePropertiesOutput {
	return i.ToMachinePropertiesOutputWithContext(context.Background())
}

func (i MachinePropertiesArgs) ToMachinePropertiesOutputWithContext(ctx context.Context) MachinePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePropertiesOutput)
}

func (i MachinePropertiesArgs) ToMachinePropertiesPtrOutput() MachinePropertiesPtrOutput {
	return i.ToMachinePropertiesPtrOutputWithContext(context.Background())
}

func (i MachinePropertiesArgs) ToMachinePropertiesPtrOutputWithContext(ctx context.Context) MachinePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePropertiesOutput).ToMachinePropertiesPtrOutputWithContext(ctx)
}









type MachinePropertiesPtrInput interface {
	pulumi.Input

	ToMachinePropertiesPtrOutput() MachinePropertiesPtrOutput
	ToMachinePropertiesPtrOutputWithContext(context.Context) MachinePropertiesPtrOutput
}

type machinePropertiesPtrType MachinePropertiesArgs

func MachinePropertiesPtr(v *MachinePropertiesArgs) MachinePropertiesPtrInput {
	return (*machinePropertiesPtrType)(v)
}

func (*machinePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineProperties)(nil)).Elem()
}

func (i *machinePropertiesPtrType) ToMachinePropertiesPtrOutput() MachinePropertiesPtrOutput {
	return i.ToMachinePropertiesPtrOutputWithContext(context.Background())
}

func (i *machinePropertiesPtrType) ToMachinePropertiesPtrOutputWithContext(ctx context.Context) MachinePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePropertiesPtrOutput)
}

type MachinePropertiesOutput struct{ *pulumi.OutputState }

func (MachinePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineProperties)(nil)).Elem()
}

func (o MachinePropertiesOutput) ToMachinePropertiesOutput() MachinePropertiesOutput {
	return o
}

func (o MachinePropertiesOutput) ToMachinePropertiesOutputWithContext(ctx context.Context) MachinePropertiesOutput {
	return o
}

func (o MachinePropertiesOutput) ToMachinePropertiesPtrOutput() MachinePropertiesPtrOutput {
	return o.ToMachinePropertiesPtrOutputWithContext(context.Background())
}

func (o MachinePropertiesOutput) ToMachinePropertiesPtrOutputWithContext(ctx context.Context) MachinePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineProperties) *MachineProperties {
		return &v
	}).(MachinePropertiesPtrOutput)
}

func (o MachinePropertiesOutput) ClientPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineProperties) *string { return v.ClientPublicKey }).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesOutput) Extensions() MachineExtensionInstanceViewArrayOutput {
	return o.ApplyT(func(v MachineProperties) []MachineExtensionInstanceView { return v.Extensions }).(MachineExtensionInstanceViewArrayOutput)
}

func (o MachinePropertiesOutput) LocationData() LocationDataPtrOutput {
	return o.ApplyT(func(v MachineProperties) *LocationData { return v.LocationData }).(LocationDataPtrOutput)
}

func (o MachinePropertiesOutput) ParentClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineProperties) *string { return v.ParentClusterResourceId }).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesOutput) PrivateLinkScopeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineProperties) *string { return v.PrivateLinkScopeResourceId }).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineProperties) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

type MachinePropertiesPtrOutput struct{ *pulumi.OutputState }

func (MachinePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineProperties)(nil)).Elem()
}

func (o MachinePropertiesPtrOutput) ToMachinePropertiesPtrOutput() MachinePropertiesPtrOutput {
	return o
}

func (o MachinePropertiesPtrOutput) ToMachinePropertiesPtrOutputWithContext(ctx context.Context) MachinePropertiesPtrOutput {
	return o
}

func (o MachinePropertiesPtrOutput) Elem() MachinePropertiesOutput {
	return o.ApplyT(func(v *MachineProperties) MachineProperties {
		if v != nil {
			return *v
		}
		var ret MachineProperties
		return ret
	}).(MachinePropertiesOutput)
}

func (o MachinePropertiesPtrOutput) ClientPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientPublicKey
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesPtrOutput) Extensions() MachineExtensionInstanceViewArrayOutput {
	return o.ApplyT(func(v *MachineProperties) []MachineExtensionInstanceView {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(MachineExtensionInstanceViewArrayOutput)
}

func (o MachinePropertiesPtrOutput) LocationData() LocationDataPtrOutput {
	return o.ApplyT(func(v *MachineProperties) *LocationData {
		if v == nil {
			return nil
		}
		return v.LocationData
	}).(LocationDataPtrOutput)
}

func (o MachinePropertiesPtrOutput) ParentClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineProperties) *string {
		if v == nil {
			return nil
		}
		return v.ParentClusterResourceId
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesPtrOutput) PrivateLinkScopeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLinkScopeResourceId
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesPtrOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineProperties) *string {
		if v == nil {
			return nil
		}
		return v.VmId
	}).(pulumi.StringPtrOutput)
}

type MachinePropertiesResponse struct {
	AdFqdn                     string                                 `pulumi:"adFqdn"`
	AgentVersion               string                                 `pulumi:"agentVersion"`
	ClientPublicKey            *string                                `pulumi:"clientPublicKey"`
	DetectedProperties         map[string]string                      `pulumi:"detectedProperties"`
	DisplayName                string                                 `pulumi:"displayName"`
	DnsFqdn                    string                                 `pulumi:"dnsFqdn"`
	DomainName                 string                                 `pulumi:"domainName"`
	ErrorDetails               []ErrorDetailResponse                  `pulumi:"errorDetails"`
	Extensions                 []MachineExtensionInstanceViewResponse `pulumi:"extensions"`
	LastStatusChange           string                                 `pulumi:"lastStatusChange"`
	LocationData               *LocationDataResponse                  `pulumi:"locationData"`
	MachineFqdn                string                                 `pulumi:"machineFqdn"`
	OsName                     string                                 `pulumi:"osName"`
	OsProfile                  OSProfileResponse                      `pulumi:"osProfile"`
	OsSku                      string                                 `pulumi:"osSku"`
	OsVersion                  string                                 `pulumi:"osVersion"`
	ParentClusterResourceId    *string                                `pulumi:"parentClusterResourceId"`
	PrivateLinkScopeResourceId *string                                `pulumi:"privateLinkScopeResourceId"`
	ProvisioningState          string                                 `pulumi:"provisioningState"`
	Status                     string                                 `pulumi:"status"`
	VmId                       *string                                `pulumi:"vmId"`
	VmUuid                     string                                 `pulumi:"vmUuid"`
}

type MachinePropertiesResponseOutput struct{ *pulumi.OutputState }

func (MachinePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachinePropertiesResponse)(nil)).Elem()
}

func (o MachinePropertiesResponseOutput) ToMachinePropertiesResponseOutput() MachinePropertiesResponseOutput {
	return o
}

func (o MachinePropertiesResponseOutput) ToMachinePropertiesResponseOutputWithContext(ctx context.Context) MachinePropertiesResponseOutput {
	return o
}

func (o MachinePropertiesResponseOutput) AdFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.AdFqdn }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) ClientPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) *string { return v.ClientPublicKey }).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponseOutput) DetectedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) map[string]string { return v.DetectedProperties }).(pulumi.StringMapOutput)
}

func (o MachinePropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) DnsFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.DnsFqdn }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) []ErrorDetailResponse { return v.ErrorDetails }).(ErrorDetailResponseArrayOutput)
}

func (o MachinePropertiesResponseOutput) Extensions() MachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) []MachineExtensionInstanceViewResponse { return v.Extensions }).(MachineExtensionInstanceViewResponseArrayOutput)
}

func (o MachinePropertiesResponseOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.LastStatusChange }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) LocationData() LocationDataResponsePtrOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) *LocationDataResponse { return v.LocationData }).(LocationDataResponsePtrOutput)
}

func (o MachinePropertiesResponseOutput) MachineFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.MachineFqdn }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.OsName }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) OsProfile() OSProfileResponseOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) OSProfileResponse { return v.OsProfile }).(OSProfileResponseOutput)
}

func (o MachinePropertiesResponseOutput) OsSku() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.OsSku }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.OsVersion }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) ParentClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) *string { return v.ParentClusterResourceId }).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponseOutput) PrivateLinkScopeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) *string { return v.PrivateLinkScopeResourceId }).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o MachinePropertiesResponseOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponseOutput) VmUuid() pulumi.StringOutput {
	return o.ApplyT(func(v MachinePropertiesResponse) string { return v.VmUuid }).(pulumi.StringOutput)
}

type OSProfileResponse struct {
	ComputerName string `pulumi:"computerName"`
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
	pulumi.RegisterOutputType(MachineExtensionPropertiesOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MachinePropertiesOutput{})
	pulumi.RegisterOutputType(MachinePropertiesPtrOutput{})
	pulumi.RegisterOutputType(MachinePropertiesResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
