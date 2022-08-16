


package v20220415preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplianceCredentialKubeconfigResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type ApplianceCredentialKubeconfigResponseOutput struct{ *pulumi.OutputState }

func (ApplianceCredentialKubeconfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceCredentialKubeconfigResponse)(nil)).Elem()
}

func (o ApplianceCredentialKubeconfigResponseOutput) ToApplianceCredentialKubeconfigResponseOutput() ApplianceCredentialKubeconfigResponseOutput {
	return o
}

func (o ApplianceCredentialKubeconfigResponseOutput) ToApplianceCredentialKubeconfigResponseOutputWithContext(ctx context.Context) ApplianceCredentialKubeconfigResponseOutput {
	return o
}

func (o ApplianceCredentialKubeconfigResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceCredentialKubeconfigResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplianceCredentialKubeconfigResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceCredentialKubeconfigResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ApplianceCredentialKubeconfigResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplianceCredentialKubeconfigResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceCredentialKubeconfigResponse)(nil)).Elem()
}

func (o ApplianceCredentialKubeconfigResponseArrayOutput) ToApplianceCredentialKubeconfigResponseArrayOutput() ApplianceCredentialKubeconfigResponseArrayOutput {
	return o
}

func (o ApplianceCredentialKubeconfigResponseArrayOutput) ToApplianceCredentialKubeconfigResponseArrayOutputWithContext(ctx context.Context) ApplianceCredentialKubeconfigResponseArrayOutput {
	return o
}

func (o ApplianceCredentialKubeconfigResponseArrayOutput) Index(i pulumi.IntInput) ApplianceCredentialKubeconfigResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplianceCredentialKubeconfigResponse {
		return vs[0].([]ApplianceCredentialKubeconfigResponse)[vs[1].(int)]
	}).(ApplianceCredentialKubeconfigResponseOutput)
}

type AppliancePropertiesInfrastructureConfig struct {
	Provider *string `pulumi:"provider"`
}





type AppliancePropertiesInfrastructureConfigInput interface {
	pulumi.Input

	ToAppliancePropertiesInfrastructureConfigOutput() AppliancePropertiesInfrastructureConfigOutput
	ToAppliancePropertiesInfrastructureConfigOutputWithContext(context.Context) AppliancePropertiesInfrastructureConfigOutput
}

type AppliancePropertiesInfrastructureConfigArgs struct {
	Provider pulumi.StringPtrInput `pulumi:"provider"`
}

func (AppliancePropertiesInfrastructureConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppliancePropertiesInfrastructureConfig)(nil)).Elem()
}

func (i AppliancePropertiesInfrastructureConfigArgs) ToAppliancePropertiesInfrastructureConfigOutput() AppliancePropertiesInfrastructureConfigOutput {
	return i.ToAppliancePropertiesInfrastructureConfigOutputWithContext(context.Background())
}

func (i AppliancePropertiesInfrastructureConfigArgs) ToAppliancePropertiesInfrastructureConfigOutputWithContext(ctx context.Context) AppliancePropertiesInfrastructureConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppliancePropertiesInfrastructureConfigOutput)
}

func (i AppliancePropertiesInfrastructureConfigArgs) ToAppliancePropertiesInfrastructureConfigPtrOutput() AppliancePropertiesInfrastructureConfigPtrOutput {
	return i.ToAppliancePropertiesInfrastructureConfigPtrOutputWithContext(context.Background())
}

func (i AppliancePropertiesInfrastructureConfigArgs) ToAppliancePropertiesInfrastructureConfigPtrOutputWithContext(ctx context.Context) AppliancePropertiesInfrastructureConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppliancePropertiesInfrastructureConfigOutput).ToAppliancePropertiesInfrastructureConfigPtrOutputWithContext(ctx)
}









type AppliancePropertiesInfrastructureConfigPtrInput interface {
	pulumi.Input

	ToAppliancePropertiesInfrastructureConfigPtrOutput() AppliancePropertiesInfrastructureConfigPtrOutput
	ToAppliancePropertiesInfrastructureConfigPtrOutputWithContext(context.Context) AppliancePropertiesInfrastructureConfigPtrOutput
}

type appliancePropertiesInfrastructureConfigPtrType AppliancePropertiesInfrastructureConfigArgs

func AppliancePropertiesInfrastructureConfigPtr(v *AppliancePropertiesInfrastructureConfigArgs) AppliancePropertiesInfrastructureConfigPtrInput {
	return (*appliancePropertiesInfrastructureConfigPtrType)(v)
}

func (*appliancePropertiesInfrastructureConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppliancePropertiesInfrastructureConfig)(nil)).Elem()
}

func (i *appliancePropertiesInfrastructureConfigPtrType) ToAppliancePropertiesInfrastructureConfigPtrOutput() AppliancePropertiesInfrastructureConfigPtrOutput {
	return i.ToAppliancePropertiesInfrastructureConfigPtrOutputWithContext(context.Background())
}

func (i *appliancePropertiesInfrastructureConfigPtrType) ToAppliancePropertiesInfrastructureConfigPtrOutputWithContext(ctx context.Context) AppliancePropertiesInfrastructureConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppliancePropertiesInfrastructureConfigPtrOutput)
}

type AppliancePropertiesInfrastructureConfigOutput struct{ *pulumi.OutputState }

func (AppliancePropertiesInfrastructureConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppliancePropertiesInfrastructureConfig)(nil)).Elem()
}

func (o AppliancePropertiesInfrastructureConfigOutput) ToAppliancePropertiesInfrastructureConfigOutput() AppliancePropertiesInfrastructureConfigOutput {
	return o
}

func (o AppliancePropertiesInfrastructureConfigOutput) ToAppliancePropertiesInfrastructureConfigOutputWithContext(ctx context.Context) AppliancePropertiesInfrastructureConfigOutput {
	return o
}

func (o AppliancePropertiesInfrastructureConfigOutput) ToAppliancePropertiesInfrastructureConfigPtrOutput() AppliancePropertiesInfrastructureConfigPtrOutput {
	return o.ToAppliancePropertiesInfrastructureConfigPtrOutputWithContext(context.Background())
}

func (o AppliancePropertiesInfrastructureConfigOutput) ToAppliancePropertiesInfrastructureConfigPtrOutputWithContext(ctx context.Context) AppliancePropertiesInfrastructureConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppliancePropertiesInfrastructureConfig) *AppliancePropertiesInfrastructureConfig {
		return &v
	}).(AppliancePropertiesInfrastructureConfigPtrOutput)
}

func (o AppliancePropertiesInfrastructureConfigOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppliancePropertiesInfrastructureConfig) *string { return v.Provider }).(pulumi.StringPtrOutput)
}

type AppliancePropertiesInfrastructureConfigPtrOutput struct{ *pulumi.OutputState }

func (AppliancePropertiesInfrastructureConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppliancePropertiesInfrastructureConfig)(nil)).Elem()
}

func (o AppliancePropertiesInfrastructureConfigPtrOutput) ToAppliancePropertiesInfrastructureConfigPtrOutput() AppliancePropertiesInfrastructureConfigPtrOutput {
	return o
}

func (o AppliancePropertiesInfrastructureConfigPtrOutput) ToAppliancePropertiesInfrastructureConfigPtrOutputWithContext(ctx context.Context) AppliancePropertiesInfrastructureConfigPtrOutput {
	return o
}

func (o AppliancePropertiesInfrastructureConfigPtrOutput) Elem() AppliancePropertiesInfrastructureConfigOutput {
	return o.ApplyT(func(v *AppliancePropertiesInfrastructureConfig) AppliancePropertiesInfrastructureConfig {
		if v != nil {
			return *v
		}
		var ret AppliancePropertiesInfrastructureConfig
		return ret
	}).(AppliancePropertiesInfrastructureConfigOutput)
}

func (o AppliancePropertiesInfrastructureConfigPtrOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppliancePropertiesInfrastructureConfig) *string {
		if v == nil {
			return nil
		}
		return v.Provider
	}).(pulumi.StringPtrOutput)
}

type AppliancePropertiesResponseInfrastructureConfig struct {
	Provider *string `pulumi:"provider"`
}

type AppliancePropertiesResponseInfrastructureConfigOutput struct{ *pulumi.OutputState }

func (AppliancePropertiesResponseInfrastructureConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppliancePropertiesResponseInfrastructureConfig)(nil)).Elem()
}

func (o AppliancePropertiesResponseInfrastructureConfigOutput) ToAppliancePropertiesResponseInfrastructureConfigOutput() AppliancePropertiesResponseInfrastructureConfigOutput {
	return o
}

func (o AppliancePropertiesResponseInfrastructureConfigOutput) ToAppliancePropertiesResponseInfrastructureConfigOutputWithContext(ctx context.Context) AppliancePropertiesResponseInfrastructureConfigOutput {
	return o
}

func (o AppliancePropertiesResponseInfrastructureConfigOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppliancePropertiesResponseInfrastructureConfig) *string { return v.Provider }).(pulumi.StringPtrOutput)
}

type AppliancePropertiesResponseInfrastructureConfigPtrOutput struct{ *pulumi.OutputState }

func (AppliancePropertiesResponseInfrastructureConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppliancePropertiesResponseInfrastructureConfig)(nil)).Elem()
}

func (o AppliancePropertiesResponseInfrastructureConfigPtrOutput) ToAppliancePropertiesResponseInfrastructureConfigPtrOutput() AppliancePropertiesResponseInfrastructureConfigPtrOutput {
	return o
}

func (o AppliancePropertiesResponseInfrastructureConfigPtrOutput) ToAppliancePropertiesResponseInfrastructureConfigPtrOutputWithContext(ctx context.Context) AppliancePropertiesResponseInfrastructureConfigPtrOutput {
	return o
}

func (o AppliancePropertiesResponseInfrastructureConfigPtrOutput) Elem() AppliancePropertiesResponseInfrastructureConfigOutput {
	return o.ApplyT(func(v *AppliancePropertiesResponseInfrastructureConfig) AppliancePropertiesResponseInfrastructureConfig {
		if v != nil {
			return *v
		}
		var ret AppliancePropertiesResponseInfrastructureConfig
		return ret
	}).(AppliancePropertiesResponseInfrastructureConfigOutput)
}

func (o AppliancePropertiesResponseInfrastructureConfigPtrOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppliancePropertiesResponseInfrastructureConfig) *string {
		if v == nil {
			return nil
		}
		return v.Provider
	}).(pulumi.StringPtrOutput)
}

type HybridConnectionConfigResponse struct {
	ExpirationTime       float64 `pulumi:"expirationTime"`
	HybridConnectionName string  `pulumi:"hybridConnectionName"`
	Relay                string  `pulumi:"relay"`
	Token                string  `pulumi:"token"`
}

type HybridConnectionConfigResponseOutput struct{ *pulumi.OutputState }

func (HybridConnectionConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionConfigResponse)(nil)).Elem()
}

func (o HybridConnectionConfigResponseOutput) ToHybridConnectionConfigResponseOutput() HybridConnectionConfigResponseOutput {
	return o
}

func (o HybridConnectionConfigResponseOutput) ToHybridConnectionConfigResponseOutputWithContext(ctx context.Context) HybridConnectionConfigResponseOutput {
	return o
}

func (o HybridConnectionConfigResponseOutput) ExpirationTime() pulumi.Float64Output {
	return o.ApplyT(func(v HybridConnectionConfigResponse) float64 { return v.ExpirationTime }).(pulumi.Float64Output)
}

func (o HybridConnectionConfigResponseOutput) HybridConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionConfigResponse) string { return v.HybridConnectionName }).(pulumi.StringOutput)
}

func (o HybridConnectionConfigResponseOutput) Relay() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionConfigResponse) string { return v.Relay }).(pulumi.StringOutput)
}

func (o HybridConnectionConfigResponseOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionConfigResponse) string { return v.Token }).(pulumi.StringOutput)
}

type Identity struct {
	Type *string `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
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

func (o IdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.Type }).(pulumi.StringPtrOutput)
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

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
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

type SSHKeyResponse struct {
	PrivateKey *string `pulumi:"privateKey"`
	PublicKey  *string `pulumi:"publicKey"`
}

type SSHKeyResponseOutput struct{ *pulumi.OutputState }

func (SSHKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SSHKeyResponse)(nil)).Elem()
}

func (o SSHKeyResponseOutput) ToSSHKeyResponseOutput() SSHKeyResponseOutput {
	return o
}

func (o SSHKeyResponseOutput) ToSSHKeyResponseOutputWithContext(ctx context.Context) SSHKeyResponseOutput {
	return o
}

func (o SSHKeyResponseOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SSHKeyResponse) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o SSHKeyResponseOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SSHKeyResponse) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

type SSHKeyResponseMapOutput struct{ *pulumi.OutputState }

func (SSHKeyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SSHKeyResponse)(nil)).Elem()
}

func (o SSHKeyResponseMapOutput) ToSSHKeyResponseMapOutput() SSHKeyResponseMapOutput {
	return o
}

func (o SSHKeyResponseMapOutput) ToSSHKeyResponseMapOutputWithContext(ctx context.Context) SSHKeyResponseMapOutput {
	return o
}

func (o SSHKeyResponseMapOutput) MapIndex(k pulumi.StringInput) SSHKeyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SSHKeyResponse {
		return vs[0].(map[string]SSHKeyResponse)[vs[1].(string)]
	}).(SSHKeyResponseOutput)
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
	pulumi.RegisterOutputType(ApplianceCredentialKubeconfigResponseOutput{})
	pulumi.RegisterOutputType(ApplianceCredentialKubeconfigResponseArrayOutput{})
	pulumi.RegisterOutputType(AppliancePropertiesInfrastructureConfigOutput{})
	pulumi.RegisterOutputType(AppliancePropertiesInfrastructureConfigPtrOutput{})
	pulumi.RegisterOutputType(AppliancePropertiesResponseInfrastructureConfigOutput{})
	pulumi.RegisterOutputType(AppliancePropertiesResponseInfrastructureConfigPtrOutput{})
	pulumi.RegisterOutputType(HybridConnectionConfigResponseOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SSHKeyResponseOutput{})
	pulumi.RegisterOutputType(SSHKeyResponseMapOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
