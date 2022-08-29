


package iotcentral

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppSkuInfo struct {
	Name string `pulumi:"name"`
}





type AppSkuInfoInput interface {
	pulumi.Input

	ToAppSkuInfoOutput() AppSkuInfoOutput
	ToAppSkuInfoOutputWithContext(context.Context) AppSkuInfoOutput
}

type AppSkuInfoArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (AppSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return i.ToAppSkuInfoOutputWithContext(context.Background())
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoOutput)
}

type AppSkuInfoOutput struct{ *pulumi.OutputState }

func (AppSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return o
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return o
}

func (o AppSkuInfoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfo) string { return v.Name }).(pulumi.StringOutput)
}

type AppSkuInfoResponse struct {
	Name string `pulumi:"name"`
}

type AppSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (AppSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfoResponse)(nil)).Elem()
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutput() AppSkuInfoResponseOutput {
	return o
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutputWithContext(ctx context.Context) AppSkuInfoResponseOutput {
	return o
}

func (o AppSkuInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfoResponse) string { return v.Name }).(pulumi.StringOutput)
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

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
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

type SystemAssignedServiceIdentity struct {
	Type string `pulumi:"type"`
}





type SystemAssignedServiceIdentityInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput
	ToSystemAssignedServiceIdentityOutputWithContext(context.Context) SystemAssignedServiceIdentityOutput
}

type SystemAssignedServiceIdentityArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (SystemAssignedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentity)(nil)).Elem()
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput {
	return i.ToSystemAssignedServiceIdentityOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityOutput)
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return i.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityOutput).ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx)
}









type SystemAssignedServiceIdentityPtrInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput
	ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Context) SystemAssignedServiceIdentityPtrOutput
}

type systemAssignedServiceIdentityPtrType SystemAssignedServiceIdentityArgs

func SystemAssignedServiceIdentityPtr(v *SystemAssignedServiceIdentityArgs) SystemAssignedServiceIdentityPtrInput {
	return (*systemAssignedServiceIdentityPtrType)(v)
}

func (*systemAssignedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentity)(nil)).Elem()
}

func (i *systemAssignedServiceIdentityPtrType) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return i.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *systemAssignedServiceIdentityPtrType) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityPtrOutput)
}

type SystemAssignedServiceIdentityOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentity)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput {
	return o
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityOutput {
	return o
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return o.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemAssignedServiceIdentity) *SystemAssignedServiceIdentity {
		return &v
	}).(SystemAssignedServiceIdentityPtrOutput)
}

func (o SystemAssignedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type SystemAssignedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentity)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityPtrOutput) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityPtrOutput) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityPtrOutput) Elem() SystemAssignedServiceIdentityOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentity) SystemAssignedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentity
		return ret
	}).(SystemAssignedServiceIdentityOutput)
}

func (o SystemAssignedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type SystemAssignedServiceIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type SystemAssignedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponseOutput() SystemAssignedServiceIdentityResponseOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponseOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponseOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o SystemAssignedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o SystemAssignedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SystemAssignedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) ToSystemAssignedServiceIdentityResponsePtrOutput() SystemAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) Elem() SystemAssignedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) SystemAssignedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentityResponse
		return ret
	}).(SystemAssignedServiceIdentityResponseOutput)
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
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
	pulumi.RegisterOutputType(AppSkuInfoOutput{})
	pulumi.RegisterOutputType(AppSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
