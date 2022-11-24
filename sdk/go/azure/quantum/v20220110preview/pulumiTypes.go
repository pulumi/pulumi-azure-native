


package v20220110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Provider struct {
	ApplicationName   *string `pulumi:"applicationName"`
	InstanceUri       *string `pulumi:"instanceUri"`
	ProviderId        *string `pulumi:"providerId"`
	ProviderSku       *string `pulumi:"providerSku"`
	ProvisioningState *string `pulumi:"provisioningState"`
	ResourceUsageId   *string `pulumi:"resourceUsageId"`
}





type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(context.Context) ProviderOutput
}

type ProviderArgs struct {
	ApplicationName   pulumi.StringPtrInput `pulumi:"applicationName"`
	InstanceUri       pulumi.StringPtrInput `pulumi:"instanceUri"`
	ProviderId        pulumi.StringPtrInput `pulumi:"providerId"`
	ProviderSku       pulumi.StringPtrInput `pulumi:"providerSku"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	ResourceUsageId   pulumi.StringPtrInput `pulumi:"resourceUsageId"`
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (i ProviderArgs) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i ProviderArgs) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}





type ProviderArrayInput interface {
	pulumi.Input

	ToProviderArrayOutput() ProviderArrayOutput
	ToProviderArrayOutputWithContext(context.Context) ProviderArrayOutput
}

type ProviderArray []ProviderInput

func (ProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Provider)(nil)).Elem()
}

func (i ProviderArray) ToProviderArrayOutput() ProviderArrayOutput {
	return i.ToProviderArrayOutputWithContext(context.Background())
}

func (i ProviderArray) ToProviderArrayOutputWithContext(ctx context.Context) ProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderArrayOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ApplicationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ApplicationName }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) InstanceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.InstanceUri }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) ProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ProviderId }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) ProviderSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ProviderSku }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) ResourceUsageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ResourceUsageId }).(pulumi.StringPtrOutput)
}

type ProviderArrayOutput struct{ *pulumi.OutputState }

func (ProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Provider)(nil)).Elem()
}

func (o ProviderArrayOutput) ToProviderArrayOutput() ProviderArrayOutput {
	return o
}

func (o ProviderArrayOutput) ToProviderArrayOutputWithContext(ctx context.Context) ProviderArrayOutput {
	return o
}

func (o ProviderArrayOutput) Index(i pulumi.IntInput) ProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Provider {
		return vs[0].([]Provider)[vs[1].(int)]
	}).(ProviderOutput)
}

type ProviderResponse struct {
	ApplicationName   *string `pulumi:"applicationName"`
	InstanceUri       *string `pulumi:"instanceUri"`
	ProviderId        *string `pulumi:"providerId"`
	ProviderSku       *string `pulumi:"providerSku"`
	ProvisioningState *string `pulumi:"provisioningState"`
	ResourceUsageId   *string `pulumi:"resourceUsageId"`
}

type ProviderResponseOutput struct{ *pulumi.OutputState }

func (ProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderResponse)(nil)).Elem()
}

func (o ProviderResponseOutput) ToProviderResponseOutput() ProviderResponseOutput {
	return o
}

func (o ProviderResponseOutput) ToProviderResponseOutputWithContext(ctx context.Context) ProviderResponseOutput {
	return o
}

func (o ProviderResponseOutput) ApplicationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ApplicationName }).(pulumi.StringPtrOutput)
}

func (o ProviderResponseOutput) InstanceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.InstanceUri }).(pulumi.StringPtrOutput)
}

func (o ProviderResponseOutput) ProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ProviderId }).(pulumi.StringPtrOutput)
}

func (o ProviderResponseOutput) ProviderSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ProviderSku }).(pulumi.StringPtrOutput)
}

func (o ProviderResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ProviderResponseOutput) ResourceUsageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ResourceUsageId }).(pulumi.StringPtrOutput)
}

type ProviderResponseArrayOutput struct{ *pulumi.OutputState }

func (ProviderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderResponse)(nil)).Elem()
}

func (o ProviderResponseArrayOutput) ToProviderResponseArrayOutput() ProviderResponseArrayOutput {
	return o
}

func (o ProviderResponseArrayOutput) ToProviderResponseArrayOutputWithContext(ctx context.Context) ProviderResponseArrayOutput {
	return o
}

func (o ProviderResponseArrayOutput) Index(i pulumi.IntInput) ProviderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderResponse {
		return vs[0].([]ProviderResponse)[vs[1].(int)]
	}).(ProviderResponseOutput)
}

type QuantumWorkspaceIdentity struct {
	Type *string `pulumi:"type"`
}





type QuantumWorkspaceIdentityInput interface {
	pulumi.Input

	ToQuantumWorkspaceIdentityOutput() QuantumWorkspaceIdentityOutput
	ToQuantumWorkspaceIdentityOutputWithContext(context.Context) QuantumWorkspaceIdentityOutput
}

type QuantumWorkspaceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (QuantumWorkspaceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QuantumWorkspaceIdentity)(nil)).Elem()
}

func (i QuantumWorkspaceIdentityArgs) ToQuantumWorkspaceIdentityOutput() QuantumWorkspaceIdentityOutput {
	return i.ToQuantumWorkspaceIdentityOutputWithContext(context.Background())
}

func (i QuantumWorkspaceIdentityArgs) ToQuantumWorkspaceIdentityOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuantumWorkspaceIdentityOutput)
}

func (i QuantumWorkspaceIdentityArgs) ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput {
	return i.ToQuantumWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (i QuantumWorkspaceIdentityArgs) ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuantumWorkspaceIdentityOutput).ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx)
}









type QuantumWorkspaceIdentityPtrInput interface {
	pulumi.Input

	ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput
	ToQuantumWorkspaceIdentityPtrOutputWithContext(context.Context) QuantumWorkspaceIdentityPtrOutput
}

type quantumWorkspaceIdentityPtrType QuantumWorkspaceIdentityArgs

func QuantumWorkspaceIdentityPtr(v *QuantumWorkspaceIdentityArgs) QuantumWorkspaceIdentityPtrInput {
	return (*quantumWorkspaceIdentityPtrType)(v)
}

func (*quantumWorkspaceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuantumWorkspaceIdentity)(nil)).Elem()
}

func (i *quantumWorkspaceIdentityPtrType) ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput {
	return i.ToQuantumWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (i *quantumWorkspaceIdentityPtrType) ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuantumWorkspaceIdentityPtrOutput)
}

type QuantumWorkspaceIdentityOutput struct{ *pulumi.OutputState }

func (QuantumWorkspaceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuantumWorkspaceIdentity)(nil)).Elem()
}

func (o QuantumWorkspaceIdentityOutput) ToQuantumWorkspaceIdentityOutput() QuantumWorkspaceIdentityOutput {
	return o
}

func (o QuantumWorkspaceIdentityOutput) ToQuantumWorkspaceIdentityOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityOutput {
	return o
}

func (o QuantumWorkspaceIdentityOutput) ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput {
	return o.ToQuantumWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (o QuantumWorkspaceIdentityOutput) ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QuantumWorkspaceIdentity) *QuantumWorkspaceIdentity {
		return &v
	}).(QuantumWorkspaceIdentityPtrOutput)
}

func (o QuantumWorkspaceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuantumWorkspaceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type QuantumWorkspaceIdentityPtrOutput struct{ *pulumi.OutputState }

func (QuantumWorkspaceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuantumWorkspaceIdentity)(nil)).Elem()
}

func (o QuantumWorkspaceIdentityPtrOutput) ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput {
	return o
}

func (o QuantumWorkspaceIdentityPtrOutput) ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityPtrOutput {
	return o
}

func (o QuantumWorkspaceIdentityPtrOutput) Elem() QuantumWorkspaceIdentityOutput {
	return o.ApplyT(func(v *QuantumWorkspaceIdentity) QuantumWorkspaceIdentity {
		if v != nil {
			return *v
		}
		var ret QuantumWorkspaceIdentity
		return ret
	}).(QuantumWorkspaceIdentityOutput)
}

func (o QuantumWorkspaceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuantumWorkspaceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type QuantumWorkspaceResponseIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type QuantumWorkspaceResponseIdentityOutput struct{ *pulumi.OutputState }

func (QuantumWorkspaceResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuantumWorkspaceResponseIdentity)(nil)).Elem()
}

func (o QuantumWorkspaceResponseIdentityOutput) ToQuantumWorkspaceResponseIdentityOutput() QuantumWorkspaceResponseIdentityOutput {
	return o
}

func (o QuantumWorkspaceResponseIdentityOutput) ToQuantumWorkspaceResponseIdentityOutputWithContext(ctx context.Context) QuantumWorkspaceResponseIdentityOutput {
	return o
}

func (o QuantumWorkspaceResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v QuantumWorkspaceResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o QuantumWorkspaceResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v QuantumWorkspaceResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o QuantumWorkspaceResponseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuantumWorkspaceResponseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type QuantumWorkspaceResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (QuantumWorkspaceResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuantumWorkspaceResponseIdentity)(nil)).Elem()
}

func (o QuantumWorkspaceResponseIdentityPtrOutput) ToQuantumWorkspaceResponseIdentityPtrOutput() QuantumWorkspaceResponseIdentityPtrOutput {
	return o
}

func (o QuantumWorkspaceResponseIdentityPtrOutput) ToQuantumWorkspaceResponseIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceResponseIdentityPtrOutput {
	return o
}

func (o QuantumWorkspaceResponseIdentityPtrOutput) Elem() QuantumWorkspaceResponseIdentityOutput {
	return o.ApplyT(func(v *QuantumWorkspaceResponseIdentity) QuantumWorkspaceResponseIdentity {
		if v != nil {
			return *v
		}
		var ret QuantumWorkspaceResponseIdentity
		return ret
	}).(QuantumWorkspaceResponseIdentityOutput)
}

func (o QuantumWorkspaceResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuantumWorkspaceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o QuantumWorkspaceResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuantumWorkspaceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o QuantumWorkspaceResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuantumWorkspaceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
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
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderArrayOutput{})
	pulumi.RegisterOutputType(ProviderResponseOutput{})
	pulumi.RegisterOutputType(ProviderResponseArrayOutput{})
	pulumi.RegisterOutputType(QuantumWorkspaceIdentityOutput{})
	pulumi.RegisterOutputType(QuantumWorkspaceIdentityPtrOutput{})
	pulumi.RegisterOutputType(QuantumWorkspaceResponseIdentityOutput{})
	pulumi.RegisterOutputType(QuantumWorkspaceResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
