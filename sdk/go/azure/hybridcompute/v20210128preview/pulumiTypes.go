


package v20210128preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}





type ErrorAdditionalInfoResponseInput interface {
	pulumi.Input

	ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput
	ToErrorAdditionalInfoResponseOutputWithContext(context.Context) ErrorAdditionalInfoResponseOutput
}

type ErrorAdditionalInfoResponseArgs struct {
	Info pulumi.Input       `pulumi:"info"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ErrorAdditionalInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (i ErrorAdditionalInfoResponseArgs) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return i.ToErrorAdditionalInfoResponseOutputWithContext(context.Background())
}

func (i ErrorAdditionalInfoResponseArgs) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorAdditionalInfoResponseOutput)
}





type ErrorAdditionalInfoResponseArrayInput interface {
	pulumi.Input

	ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput
	ToErrorAdditionalInfoResponseArrayOutputWithContext(context.Context) ErrorAdditionalInfoResponseArrayOutput
}

type ErrorAdditionalInfoResponseArray []ErrorAdditionalInfoResponseInput

func (ErrorAdditionalInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (i ErrorAdditionalInfoResponseArray) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return i.ToErrorAdditionalInfoResponseArrayOutputWithContext(context.Background())
}

func (i ErrorAdditionalInfoResponseArray) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorAdditionalInfoResponseArrayOutput)
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





type ErrorDetailResponseInput interface {
	pulumi.Input

	ToErrorDetailResponseOutput() ErrorDetailResponseOutput
	ToErrorDetailResponseOutputWithContext(context.Context) ErrorDetailResponseOutput
}

type ErrorDetailResponseArgs struct {
	AdditionalInfo ErrorAdditionalInfoResponseArrayInput `pulumi:"additionalInfo"`
	Code           pulumi.StringInput                    `pulumi:"code"`
	Details        ErrorDetailResponseArrayInput         `pulumi:"details"`
	Message        pulumi.StringInput                    `pulumi:"message"`
	Target         pulumi.StringInput                    `pulumi:"target"`
}

func (ErrorDetailResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return i.ToErrorDetailResponseOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArgs) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseOutput)
}





type ErrorDetailResponseArrayInput interface {
	pulumi.Input

	ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput
	ToErrorDetailResponseArrayOutputWithContext(context.Context) ErrorDetailResponseArrayOutput
}

type ErrorDetailResponseArray []ErrorDetailResponseInput

func (ErrorDetailResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return i.ToErrorDetailResponseArrayOutputWithContext(context.Background())
}

func (i ErrorDetailResponseArray) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailResponseArrayOutput)
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





type HybridComputePrivateLinkScopePropertiesResponseInput interface {
	pulumi.Input

	ToHybridComputePrivateLinkScopePropertiesResponseOutput() HybridComputePrivateLinkScopePropertiesResponseOutput
	ToHybridComputePrivateLinkScopePropertiesResponseOutputWithContext(context.Context) HybridComputePrivateLinkScopePropertiesResponseOutput
}

type HybridComputePrivateLinkScopePropertiesResponseArgs struct {
	PrivateLinkScopeId  pulumi.StringInput    `pulumi:"privateLinkScopeId"`
	ProvisioningState   pulumi.StringInput    `pulumi:"provisioningState"`
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (HybridComputePrivateLinkScopePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputePrivateLinkScopePropertiesResponse)(nil)).Elem()
}

func (i HybridComputePrivateLinkScopePropertiesResponseArgs) ToHybridComputePrivateLinkScopePropertiesResponseOutput() HybridComputePrivateLinkScopePropertiesResponseOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesResponseOutputWithContext(context.Background())
}

func (i HybridComputePrivateLinkScopePropertiesResponseArgs) ToHybridComputePrivateLinkScopePropertiesResponseOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesResponseOutput)
}

func (i HybridComputePrivateLinkScopePropertiesResponseArgs) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i HybridComputePrivateLinkScopePropertiesResponseArgs) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesResponseOutput).ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx)
}









type HybridComputePrivateLinkScopePropertiesResponsePtrInput interface {
	pulumi.Input

	ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput
	ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput
}

type hybridComputePrivateLinkScopePropertiesResponsePtrType HybridComputePrivateLinkScopePropertiesResponseArgs

func HybridComputePrivateLinkScopePropertiesResponsePtr(v *HybridComputePrivateLinkScopePropertiesResponseArgs) HybridComputePrivateLinkScopePropertiesResponsePtrInput {
	return (*hybridComputePrivateLinkScopePropertiesResponsePtrType)(v)
}

func (*hybridComputePrivateLinkScopePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputePrivateLinkScopePropertiesResponse)(nil)).Elem()
}

func (i *hybridComputePrivateLinkScopePropertiesResponsePtrType) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return i.ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *hybridComputePrivateLinkScopePropertiesResponsePtrType) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputePrivateLinkScopePropertiesResponsePtrOutput)
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

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return o.ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o HybridComputePrivateLinkScopePropertiesResponseOutput) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HybridComputePrivateLinkScopePropertiesResponse) *HybridComputePrivateLinkScopePropertiesResponse {
		return &v
	}).(HybridComputePrivateLinkScopePropertiesResponsePtrOutput)
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

type HybridComputePrivateLinkScopePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (HybridComputePrivateLinkScopePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputePrivateLinkScopePropertiesResponse)(nil)).Elem()
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutput() HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) ToHybridComputePrivateLinkScopePropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputePrivateLinkScopePropertiesResponsePtrOutput {
	return o
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) Elem() HybridComputePrivateLinkScopePropertiesResponseOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopePropertiesResponse) HybridComputePrivateLinkScopePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret HybridComputePrivateLinkScopePropertiesResponse
		return ret
	}).(HybridComputePrivateLinkScopePropertiesResponseOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) PrivateLinkScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkScopeId
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputePrivateLinkScopePropertiesResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputePrivateLinkScopePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
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





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
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

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
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





type LocationDataResponseInput interface {
	pulumi.Input

	ToLocationDataResponseOutput() LocationDataResponseOutput
	ToLocationDataResponseOutputWithContext(context.Context) LocationDataResponseOutput
}

type LocationDataResponseArgs struct {
	City            pulumi.StringPtrInput `pulumi:"city"`
	CountryOrRegion pulumi.StringPtrInput `pulumi:"countryOrRegion"`
	District        pulumi.StringPtrInput `pulumi:"district"`
	Name            pulumi.StringInput    `pulumi:"name"`
}

func (LocationDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationDataResponse)(nil)).Elem()
}

func (i LocationDataResponseArgs) ToLocationDataResponseOutput() LocationDataResponseOutput {
	return i.ToLocationDataResponseOutputWithContext(context.Background())
}

func (i LocationDataResponseArgs) ToLocationDataResponseOutputWithContext(ctx context.Context) LocationDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataResponseOutput)
}

func (i LocationDataResponseArgs) ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput {
	return i.ToLocationDataResponsePtrOutputWithContext(context.Background())
}

func (i LocationDataResponseArgs) ToLocationDataResponsePtrOutputWithContext(ctx context.Context) LocationDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataResponseOutput).ToLocationDataResponsePtrOutputWithContext(ctx)
}









type LocationDataResponsePtrInput interface {
	pulumi.Input

	ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput
	ToLocationDataResponsePtrOutputWithContext(context.Context) LocationDataResponsePtrOutput
}

type locationDataResponsePtrType LocationDataResponseArgs

func LocationDataResponsePtr(v *LocationDataResponseArgs) LocationDataResponsePtrInput {
	return (*locationDataResponsePtrType)(v)
}

func (*locationDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationDataResponse)(nil)).Elem()
}

func (i *locationDataResponsePtrType) ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput {
	return i.ToLocationDataResponsePtrOutputWithContext(context.Background())
}

func (i *locationDataResponsePtrType) ToLocationDataResponsePtrOutputWithContext(ctx context.Context) LocationDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataResponsePtrOutput)
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

func (o LocationDataResponseOutput) ToLocationDataResponsePtrOutput() LocationDataResponsePtrOutput {
	return o.ToLocationDataResponsePtrOutputWithContext(context.Background())
}

func (o LocationDataResponseOutput) ToLocationDataResponsePtrOutputWithContext(ctx context.Context) LocationDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocationDataResponse) *LocationDataResponse {
		return &v
	}).(LocationDataResponsePtrOutput)
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





type MachineExtensionInstanceViewResponseInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseOutput() MachineExtensionInstanceViewResponseOutput
	ToMachineExtensionInstanceViewResponseOutputWithContext(context.Context) MachineExtensionInstanceViewResponseOutput
}

type MachineExtensionInstanceViewResponseArgs struct {
	Name               pulumi.StringPtrInput                              `pulumi:"name"`
	Status             MachineExtensionInstanceViewResponseStatusPtrInput `pulumi:"status"`
	Type               pulumi.StringPtrInput                              `pulumi:"type"`
	TypeHandlerVersion pulumi.StringPtrInput                              `pulumi:"typeHandlerVersion"`
}

func (MachineExtensionInstanceViewResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponseOutput() MachineExtensionInstanceViewResponseOutput {
	return i.ToMachineExtensionInstanceViewResponseOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponseOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseOutput)
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput {
	return i.ToMachineExtensionInstanceViewResponsePtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseArgs) ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseOutput).ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx)
}









type MachineExtensionInstanceViewResponsePtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput
	ToMachineExtensionInstanceViewResponsePtrOutputWithContext(context.Context) MachineExtensionInstanceViewResponsePtrOutput
}

type machineExtensionInstanceViewResponsePtrType MachineExtensionInstanceViewResponseArgs

func MachineExtensionInstanceViewResponsePtr(v *MachineExtensionInstanceViewResponseArgs) MachineExtensionInstanceViewResponsePtrInput {
	return (*machineExtensionInstanceViewResponsePtrType)(v)
}

func (*machineExtensionInstanceViewResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (i *machineExtensionInstanceViewResponsePtrType) ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput {
	return i.ToMachineExtensionInstanceViewResponsePtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewResponsePtrType) ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponsePtrOutput)
}





type MachineExtensionInstanceViewResponseArrayInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseArrayOutput() MachineExtensionInstanceViewResponseArrayOutput
	ToMachineExtensionInstanceViewResponseArrayOutputWithContext(context.Context) MachineExtensionInstanceViewResponseArrayOutput
}

type MachineExtensionInstanceViewResponseArray []MachineExtensionInstanceViewResponseInput

func (MachineExtensionInstanceViewResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (i MachineExtensionInstanceViewResponseArray) ToMachineExtensionInstanceViewResponseArrayOutput() MachineExtensionInstanceViewResponseArrayOutput {
	return i.ToMachineExtensionInstanceViewResponseArrayOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseArray) ToMachineExtensionInstanceViewResponseArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseArrayOutput)
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

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponsePtrOutput() MachineExtensionInstanceViewResponsePtrOutput {
	return o.ToMachineExtensionInstanceViewResponsePtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewResponseOutput) ToMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionInstanceViewResponse) *MachineExtensionInstanceViewResponse {
		return &v
	}).(MachineExtensionInstanceViewResponsePtrOutput)
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





type MachineExtensionInstanceViewResponseStatusInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput
	ToMachineExtensionInstanceViewResponseStatusOutputWithContext(context.Context) MachineExtensionInstanceViewResponseStatusOutput
}

type MachineExtensionInstanceViewResponseStatusArgs struct {
	Code          pulumi.StringPtrInput `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput `pulumi:"displayStatus"`
	Level         pulumi.StringPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput `pulumi:"message"`
	Time          pulumi.StringPtrInput `pulumi:"time"`
}

func (MachineExtensionInstanceViewResponseStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput {
	return i.ToMachineExtensionInstanceViewResponseStatusOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseStatusOutput)
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewResponseStatusArgs) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseStatusOutput).ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx)
}









type MachineExtensionInstanceViewResponseStatusPtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput
	ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput
}

type machineExtensionInstanceViewResponseStatusPtrType MachineExtensionInstanceViewResponseStatusArgs

func MachineExtensionInstanceViewResponseStatusPtr(v *MachineExtensionInstanceViewResponseStatusArgs) MachineExtensionInstanceViewResponseStatusPtrInput {
	return (*machineExtensionInstanceViewResponseStatusPtrType)(v)
}

func (*machineExtensionInstanceViewResponseStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (i *machineExtensionInstanceViewResponseStatusPtrType) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return i.ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewResponseStatusPtrType) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewResponseStatusPtrOutput)
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

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionInstanceViewResponseStatus) *MachineExtensionInstanceViewResponseStatus {
		return &v
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
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





type MachineExtensionPropertiesResponseInput interface {
	pulumi.Input

	ToMachineExtensionPropertiesResponseOutput() MachineExtensionPropertiesResponseOutput
	ToMachineExtensionPropertiesResponseOutputWithContext(context.Context) MachineExtensionPropertiesResponseOutput
}

type MachineExtensionPropertiesResponseArgs struct {
	AutoUpgradeMinorVersion pulumi.BoolPtrInput                          `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          pulumi.StringPtrInput                        `pulumi:"forceUpdateTag"`
	InstanceView            MachineExtensionInstanceViewResponsePtrInput `pulumi:"instanceView"`
	ProtectedSettings       pulumi.Input                                 `pulumi:"protectedSettings"`
	ProvisioningState       pulumi.StringInput                           `pulumi:"provisioningState"`
	Publisher               pulumi.StringPtrInput                        `pulumi:"publisher"`
	Settings                pulumi.Input                                 `pulumi:"settings"`
	Type                    pulumi.StringPtrInput                        `pulumi:"type"`
	TypeHandlerVersion      pulumi.StringPtrInput                        `pulumi:"typeHandlerVersion"`
}

func (MachineExtensionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionPropertiesResponse)(nil)).Elem()
}

func (i MachineExtensionPropertiesResponseArgs) ToMachineExtensionPropertiesResponseOutput() MachineExtensionPropertiesResponseOutput {
	return i.ToMachineExtensionPropertiesResponseOutputWithContext(context.Background())
}

func (i MachineExtensionPropertiesResponseArgs) ToMachineExtensionPropertiesResponseOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionPropertiesResponseOutput)
}

func (i MachineExtensionPropertiesResponseArgs) ToMachineExtensionPropertiesResponsePtrOutput() MachineExtensionPropertiesResponsePtrOutput {
	return i.ToMachineExtensionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MachineExtensionPropertiesResponseArgs) ToMachineExtensionPropertiesResponsePtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionPropertiesResponseOutput).ToMachineExtensionPropertiesResponsePtrOutputWithContext(ctx)
}









type MachineExtensionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMachineExtensionPropertiesResponsePtrOutput() MachineExtensionPropertiesResponsePtrOutput
	ToMachineExtensionPropertiesResponsePtrOutputWithContext(context.Context) MachineExtensionPropertiesResponsePtrOutput
}

type machineExtensionPropertiesResponsePtrType MachineExtensionPropertiesResponseArgs

func MachineExtensionPropertiesResponsePtr(v *MachineExtensionPropertiesResponseArgs) MachineExtensionPropertiesResponsePtrInput {
	return (*machineExtensionPropertiesResponsePtrType)(v)
}

func (*machineExtensionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionPropertiesResponse)(nil)).Elem()
}

func (i *machineExtensionPropertiesResponsePtrType) ToMachineExtensionPropertiesResponsePtrOutput() MachineExtensionPropertiesResponsePtrOutput {
	return i.ToMachineExtensionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *machineExtensionPropertiesResponsePtrType) ToMachineExtensionPropertiesResponsePtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionPropertiesResponsePtrOutput)
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

func (o MachineExtensionPropertiesResponseOutput) ToMachineExtensionPropertiesResponsePtrOutput() MachineExtensionPropertiesResponsePtrOutput {
	return o.ToMachineExtensionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MachineExtensionPropertiesResponseOutput) ToMachineExtensionPropertiesResponsePtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachineExtensionPropertiesResponse) *MachineExtensionPropertiesResponse {
		return &v
	}).(MachineExtensionPropertiesResponsePtrOutput)
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

type MachineExtensionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionPropertiesResponse)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponsePtrOutput) ToMachineExtensionPropertiesResponsePtrOutput() MachineExtensionPropertiesResponsePtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponsePtrOutput) ToMachineExtensionPropertiesResponsePtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponsePtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponsePtrOutput) Elem() MachineExtensionPropertiesResponseOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) MachineExtensionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MachineExtensionPropertiesResponse
		return ret
	}).(MachineExtensionPropertiesResponseOutput)
}

func (o MachineExtensionPropertiesResponsePtrOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutoUpgradeMinorVersion
	}).(pulumi.BoolPtrOutput)
}

func (o MachineExtensionPropertiesResponsePtrOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ForceUpdateTag
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponsePtrOutput) InstanceView() MachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) *MachineExtensionInstanceViewResponse {
		if v == nil {
			return nil
		}
		return v.InstanceView
	}).(MachineExtensionInstanceViewResponsePtrOutput)
}

func (o MachineExtensionPropertiesResponsePtrOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ProtectedSettings
	}).(pulumi.AnyOutput)
}

func (o MachineExtensionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponsePtrOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Settings
	}).(pulumi.AnyOutput)
}

func (o MachineExtensionPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponsePtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type MachineProperties struct {
	ClientPublicKey            *string                        `pulumi:"clientPublicKey"`
	Extensions                 []MachineExtensionInstanceView `pulumi:"extensions"`
	LocationData               *LocationData                  `pulumi:"locationData"`
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
	PrivateLinkScopeResourceId *string                                `pulumi:"privateLinkScopeResourceId"`
	ProvisioningState          string                                 `pulumi:"provisioningState"`
	Status                     string                                 `pulumi:"status"`
	VmId                       *string                                `pulumi:"vmId"`
	VmUuid                     string                                 `pulumi:"vmUuid"`
}





type MachinePropertiesResponseInput interface {
	pulumi.Input

	ToMachinePropertiesResponseOutput() MachinePropertiesResponseOutput
	ToMachinePropertiesResponseOutputWithContext(context.Context) MachinePropertiesResponseOutput
}

type MachinePropertiesResponseArgs struct {
	AdFqdn                     pulumi.StringInput                             `pulumi:"adFqdn"`
	AgentVersion               pulumi.StringInput                             `pulumi:"agentVersion"`
	ClientPublicKey            pulumi.StringPtrInput                          `pulumi:"clientPublicKey"`
	DisplayName                pulumi.StringInput                             `pulumi:"displayName"`
	DnsFqdn                    pulumi.StringInput                             `pulumi:"dnsFqdn"`
	DomainName                 pulumi.StringInput                             `pulumi:"domainName"`
	ErrorDetails               ErrorDetailResponseArrayInput                  `pulumi:"errorDetails"`
	Extensions                 MachineExtensionInstanceViewResponseArrayInput `pulumi:"extensions"`
	LastStatusChange           pulumi.StringInput                             `pulumi:"lastStatusChange"`
	LocationData               LocationDataResponsePtrInput                   `pulumi:"locationData"`
	MachineFqdn                pulumi.StringInput                             `pulumi:"machineFqdn"`
	OsName                     pulumi.StringInput                             `pulumi:"osName"`
	OsProfile                  OSProfileResponseInput                         `pulumi:"osProfile"`
	OsSku                      pulumi.StringInput                             `pulumi:"osSku"`
	OsVersion                  pulumi.StringInput                             `pulumi:"osVersion"`
	PrivateLinkScopeResourceId pulumi.StringPtrInput                          `pulumi:"privateLinkScopeResourceId"`
	ProvisioningState          pulumi.StringInput                             `pulumi:"provisioningState"`
	Status                     pulumi.StringInput                             `pulumi:"status"`
	VmId                       pulumi.StringPtrInput                          `pulumi:"vmId"`
	VmUuid                     pulumi.StringInput                             `pulumi:"vmUuid"`
}

func (MachinePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachinePropertiesResponse)(nil)).Elem()
}

func (i MachinePropertiesResponseArgs) ToMachinePropertiesResponseOutput() MachinePropertiesResponseOutput {
	return i.ToMachinePropertiesResponseOutputWithContext(context.Background())
}

func (i MachinePropertiesResponseArgs) ToMachinePropertiesResponseOutputWithContext(ctx context.Context) MachinePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePropertiesResponseOutput)
}

func (i MachinePropertiesResponseArgs) ToMachinePropertiesResponsePtrOutput() MachinePropertiesResponsePtrOutput {
	return i.ToMachinePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MachinePropertiesResponseArgs) ToMachinePropertiesResponsePtrOutputWithContext(ctx context.Context) MachinePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePropertiesResponseOutput).ToMachinePropertiesResponsePtrOutputWithContext(ctx)
}









type MachinePropertiesResponsePtrInput interface {
	pulumi.Input

	ToMachinePropertiesResponsePtrOutput() MachinePropertiesResponsePtrOutput
	ToMachinePropertiesResponsePtrOutputWithContext(context.Context) MachinePropertiesResponsePtrOutput
}

type machinePropertiesResponsePtrType MachinePropertiesResponseArgs

func MachinePropertiesResponsePtr(v *MachinePropertiesResponseArgs) MachinePropertiesResponsePtrInput {
	return (*machinePropertiesResponsePtrType)(v)
}

func (*machinePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachinePropertiesResponse)(nil)).Elem()
}

func (i *machinePropertiesResponsePtrType) ToMachinePropertiesResponsePtrOutput() MachinePropertiesResponsePtrOutput {
	return i.ToMachinePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *machinePropertiesResponsePtrType) ToMachinePropertiesResponsePtrOutputWithContext(ctx context.Context) MachinePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachinePropertiesResponsePtrOutput)
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

func (o MachinePropertiesResponseOutput) ToMachinePropertiesResponsePtrOutput() MachinePropertiesResponsePtrOutput {
	return o.ToMachinePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MachinePropertiesResponseOutput) ToMachinePropertiesResponsePtrOutputWithContext(ctx context.Context) MachinePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MachinePropertiesResponse) *MachinePropertiesResponse {
		return &v
	}).(MachinePropertiesResponsePtrOutput)
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

type MachinePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MachinePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachinePropertiesResponse)(nil)).Elem()
}

func (o MachinePropertiesResponsePtrOutput) ToMachinePropertiesResponsePtrOutput() MachinePropertiesResponsePtrOutput {
	return o
}

func (o MachinePropertiesResponsePtrOutput) ToMachinePropertiesResponsePtrOutputWithContext(ctx context.Context) MachinePropertiesResponsePtrOutput {
	return o
}

func (o MachinePropertiesResponsePtrOutput) Elem() MachinePropertiesResponseOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) MachinePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MachinePropertiesResponse
		return ret
	}).(MachinePropertiesResponseOutput)
}

func (o MachinePropertiesResponsePtrOutput) AdFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdFqdn
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AgentVersion
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) ClientPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientPublicKey
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) DnsFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DnsFqdn
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainName
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) []ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.ErrorDetails
	}).(ErrorDetailResponseArrayOutput)
}

func (o MachinePropertiesResponsePtrOutput) Extensions() MachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) []MachineExtensionInstanceViewResponse {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(MachineExtensionInstanceViewResponseArrayOutput)
}

func (o MachinePropertiesResponsePtrOutput) LastStatusChange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastStatusChange
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) LocationData() LocationDataResponsePtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *LocationDataResponse {
		if v == nil {
			return nil
		}
		return v.LocationData
	}).(LocationDataResponsePtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) MachineFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MachineFqdn
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) OsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsName
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *OSProfileResponse {
		if v == nil {
			return nil
		}
		return &v.OsProfile
	}).(OSProfileResponsePtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) OsSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsSku
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsVersion
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) PrivateLinkScopeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateLinkScopeResourceId
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmId
	}).(pulumi.StringPtrOutput)
}

func (o MachinePropertiesResponsePtrOutput) VmUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachinePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VmUuid
	}).(pulumi.StringPtrOutput)
}

type OSProfileResponse struct {
	ComputerName string `pulumi:"computerName"`
}





type OSProfileResponseInput interface {
	pulumi.Input

	ToOSProfileResponseOutput() OSProfileResponseOutput
	ToOSProfileResponseOutputWithContext(context.Context) OSProfileResponseOutput
}

type OSProfileResponseArgs struct {
	ComputerName pulumi.StringInput `pulumi:"computerName"`
}

func (OSProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponse)(nil)).Elem()
}

func (i OSProfileResponseArgs) ToOSProfileResponseOutput() OSProfileResponseOutput {
	return i.ToOSProfileResponseOutputWithContext(context.Background())
}

func (i OSProfileResponseArgs) ToOSProfileResponseOutputWithContext(ctx context.Context) OSProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileResponseOutput)
}

func (i OSProfileResponseArgs) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return i.ToOSProfileResponsePtrOutputWithContext(context.Background())
}

func (i OSProfileResponseArgs) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileResponseOutput).ToOSProfileResponsePtrOutputWithContext(ctx)
}









type OSProfileResponsePtrInput interface {
	pulumi.Input

	ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput
	ToOSProfileResponsePtrOutputWithContext(context.Context) OSProfileResponsePtrOutput
}

type osprofileResponsePtrType OSProfileResponseArgs

func OSProfileResponsePtr(v *OSProfileResponseArgs) OSProfileResponsePtrInput {
	return (*osprofileResponsePtrType)(v)
}

func (*osprofileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponse)(nil)).Elem()
}

func (i *osprofileResponsePtrType) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return i.ToOSProfileResponsePtrOutputWithContext(context.Background())
}

func (i *osprofileResponsePtrType) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileResponsePtrOutput)
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

func (o OSProfileResponseOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o.ToOSProfileResponsePtrOutputWithContext(context.Background())
}

func (o OSProfileResponseOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSProfileResponse) *OSProfileResponse {
		return &v
	}).(OSProfileResponsePtrOutput)
}

func (o OSProfileResponseOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v OSProfileResponse) string { return v.ComputerName }).(pulumi.StringOutput)
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





type PrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput
	ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponseOutput
}

type PrivateEndpointConnectionPropertiesResponseArgs struct {
	PrivateEndpoint                   PrivateEndpointPropertyResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                        `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput).ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput
	ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput
}

type privateEndpointConnectionPropertiesResponsePtrType PrivateEndpointConnectionPropertiesResponseArgs

func PrivateEndpointConnectionPropertiesResponsePtr(v *PrivateEndpointConnectionPropertiesResponseArgs) PrivateEndpointConnectionPropertiesResponsePtrInput {
	return (*privateEndpointConnectionPropertiesResponsePtrType)(v)
}

func (*privateEndpointConnectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
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

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponse {
		return &v
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
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





type PrivateEndpointPropertyResponseInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput
	ToPrivateEndpointPropertyResponseOutputWithContext(context.Context) PrivateEndpointPropertyResponseOutput
}

type PrivateEndpointPropertyResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return i.ToPrivateEndpointPropertyResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput)
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput).ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput
	ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Context) PrivateEndpointPropertyResponsePtrOutput
}

type privateEndpointPropertyResponsePtrType PrivateEndpointPropertyResponseArgs

func PrivateEndpointPropertyResponsePtr(v *PrivateEndpointPropertyResponseArgs) PrivateEndpointPropertyResponsePtrInput {
	return (*privateEndpointPropertyResponsePtrType)(v)
}

func (*privateEndpointPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponsePtrOutput)
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

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointPropertyResponse) *PrivateEndpointPropertyResponse {
		return &v
	}).(PrivateEndpointPropertyResponsePtrOutput)
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





type PrivateLinkServiceConnectionStatePropertyResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput
	ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput
}

type PrivateLinkServiceConnectionStatePropertyResponseArgs struct {
	ActionsRequired pulumi.StringInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput `pulumi:"description"`
	Status          pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput).ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
}

type privateLinkServiceConnectionStatePropertyResponsePtrType PrivateLinkServiceConnectionStatePropertyResponseArgs

func PrivateLinkServiceConnectionStatePropertyResponsePtr(v *PrivateLinkServiceConnectionStatePropertyResponseArgs) PrivateLinkServiceConnectionStatePropertyResponsePtrInput {
	return (*privateLinkServiceConnectionStatePropertyResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
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

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatePropertyResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
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





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesPtrOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesResponseOutput{})
	pulumi.RegisterOutputType(HybridComputePrivateLinkScopePropertiesResponsePtrOutput{})
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
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MachinePropertiesOutput{})
	pulumi.RegisterOutputType(MachinePropertiesPtrOutput{})
	pulumi.RegisterOutputType(MachinePropertiesResponseOutput{})
	pulumi.RegisterOutputType(MachinePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponsePtrOutput{})
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
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
