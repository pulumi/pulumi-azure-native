


package v20161101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateFirewallRuleWithAccountParameters struct {
	EndIpAddress   string `pulumi:"endIpAddress"`
	Name           string `pulumi:"name"`
	StartIpAddress string `pulumi:"startIpAddress"`
}





type CreateFirewallRuleWithAccountParametersInput interface {
	pulumi.Input

	ToCreateFirewallRuleWithAccountParametersOutput() CreateFirewallRuleWithAccountParametersOutput
	ToCreateFirewallRuleWithAccountParametersOutputWithContext(context.Context) CreateFirewallRuleWithAccountParametersOutput
}

type CreateFirewallRuleWithAccountParametersArgs struct {
	EndIpAddress   pulumi.StringInput `pulumi:"endIpAddress"`
	Name           pulumi.StringInput `pulumi:"name"`
	StartIpAddress pulumi.StringInput `pulumi:"startIpAddress"`
}

func (CreateFirewallRuleWithAccountParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateFirewallRuleWithAccountParameters)(nil)).Elem()
}

func (i CreateFirewallRuleWithAccountParametersArgs) ToCreateFirewallRuleWithAccountParametersOutput() CreateFirewallRuleWithAccountParametersOutput {
	return i.ToCreateFirewallRuleWithAccountParametersOutputWithContext(context.Background())
}

func (i CreateFirewallRuleWithAccountParametersArgs) ToCreateFirewallRuleWithAccountParametersOutputWithContext(ctx context.Context) CreateFirewallRuleWithAccountParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateFirewallRuleWithAccountParametersOutput)
}





type CreateFirewallRuleWithAccountParametersArrayInput interface {
	pulumi.Input

	ToCreateFirewallRuleWithAccountParametersArrayOutput() CreateFirewallRuleWithAccountParametersArrayOutput
	ToCreateFirewallRuleWithAccountParametersArrayOutputWithContext(context.Context) CreateFirewallRuleWithAccountParametersArrayOutput
}

type CreateFirewallRuleWithAccountParametersArray []CreateFirewallRuleWithAccountParametersInput

func (CreateFirewallRuleWithAccountParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateFirewallRuleWithAccountParameters)(nil)).Elem()
}

func (i CreateFirewallRuleWithAccountParametersArray) ToCreateFirewallRuleWithAccountParametersArrayOutput() CreateFirewallRuleWithAccountParametersArrayOutput {
	return i.ToCreateFirewallRuleWithAccountParametersArrayOutputWithContext(context.Background())
}

func (i CreateFirewallRuleWithAccountParametersArray) ToCreateFirewallRuleWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateFirewallRuleWithAccountParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateFirewallRuleWithAccountParametersArrayOutput)
}

type CreateFirewallRuleWithAccountParametersOutput struct{ *pulumi.OutputState }

func (CreateFirewallRuleWithAccountParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateFirewallRuleWithAccountParameters)(nil)).Elem()
}

func (o CreateFirewallRuleWithAccountParametersOutput) ToCreateFirewallRuleWithAccountParametersOutput() CreateFirewallRuleWithAccountParametersOutput {
	return o
}

func (o CreateFirewallRuleWithAccountParametersOutput) ToCreateFirewallRuleWithAccountParametersOutputWithContext(ctx context.Context) CreateFirewallRuleWithAccountParametersOutput {
	return o
}

func (o CreateFirewallRuleWithAccountParametersOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v CreateFirewallRuleWithAccountParameters) string { return v.EndIpAddress }).(pulumi.StringOutput)
}

func (o CreateFirewallRuleWithAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CreateFirewallRuleWithAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

func (o CreateFirewallRuleWithAccountParametersOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v CreateFirewallRuleWithAccountParameters) string { return v.StartIpAddress }).(pulumi.StringOutput)
}

type CreateFirewallRuleWithAccountParametersArrayOutput struct{ *pulumi.OutputState }

func (CreateFirewallRuleWithAccountParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateFirewallRuleWithAccountParameters)(nil)).Elem()
}

func (o CreateFirewallRuleWithAccountParametersArrayOutput) ToCreateFirewallRuleWithAccountParametersArrayOutput() CreateFirewallRuleWithAccountParametersArrayOutput {
	return o
}

func (o CreateFirewallRuleWithAccountParametersArrayOutput) ToCreateFirewallRuleWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateFirewallRuleWithAccountParametersArrayOutput {
	return o
}

func (o CreateFirewallRuleWithAccountParametersArrayOutput) Index(i pulumi.IntInput) CreateFirewallRuleWithAccountParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CreateFirewallRuleWithAccountParameters {
		return vs[0].([]CreateFirewallRuleWithAccountParameters)[vs[1].(int)]
	}).(CreateFirewallRuleWithAccountParametersOutput)
}

type CreateTrustedIdProviderWithAccountParameters struct {
	IdProvider string `pulumi:"idProvider"`
	Name       string `pulumi:"name"`
}





type CreateTrustedIdProviderWithAccountParametersInput interface {
	pulumi.Input

	ToCreateTrustedIdProviderWithAccountParametersOutput() CreateTrustedIdProviderWithAccountParametersOutput
	ToCreateTrustedIdProviderWithAccountParametersOutputWithContext(context.Context) CreateTrustedIdProviderWithAccountParametersOutput
}

type CreateTrustedIdProviderWithAccountParametersArgs struct {
	IdProvider pulumi.StringInput `pulumi:"idProvider"`
	Name       pulumi.StringInput `pulumi:"name"`
}

func (CreateTrustedIdProviderWithAccountParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateTrustedIdProviderWithAccountParameters)(nil)).Elem()
}

func (i CreateTrustedIdProviderWithAccountParametersArgs) ToCreateTrustedIdProviderWithAccountParametersOutput() CreateTrustedIdProviderWithAccountParametersOutput {
	return i.ToCreateTrustedIdProviderWithAccountParametersOutputWithContext(context.Background())
}

func (i CreateTrustedIdProviderWithAccountParametersArgs) ToCreateTrustedIdProviderWithAccountParametersOutputWithContext(ctx context.Context) CreateTrustedIdProviderWithAccountParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateTrustedIdProviderWithAccountParametersOutput)
}





type CreateTrustedIdProviderWithAccountParametersArrayInput interface {
	pulumi.Input

	ToCreateTrustedIdProviderWithAccountParametersArrayOutput() CreateTrustedIdProviderWithAccountParametersArrayOutput
	ToCreateTrustedIdProviderWithAccountParametersArrayOutputWithContext(context.Context) CreateTrustedIdProviderWithAccountParametersArrayOutput
}

type CreateTrustedIdProviderWithAccountParametersArray []CreateTrustedIdProviderWithAccountParametersInput

func (CreateTrustedIdProviderWithAccountParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateTrustedIdProviderWithAccountParameters)(nil)).Elem()
}

func (i CreateTrustedIdProviderWithAccountParametersArray) ToCreateTrustedIdProviderWithAccountParametersArrayOutput() CreateTrustedIdProviderWithAccountParametersArrayOutput {
	return i.ToCreateTrustedIdProviderWithAccountParametersArrayOutputWithContext(context.Background())
}

func (i CreateTrustedIdProviderWithAccountParametersArray) ToCreateTrustedIdProviderWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateTrustedIdProviderWithAccountParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateTrustedIdProviderWithAccountParametersArrayOutput)
}

type CreateTrustedIdProviderWithAccountParametersOutput struct{ *pulumi.OutputState }

func (CreateTrustedIdProviderWithAccountParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateTrustedIdProviderWithAccountParameters)(nil)).Elem()
}

func (o CreateTrustedIdProviderWithAccountParametersOutput) ToCreateTrustedIdProviderWithAccountParametersOutput() CreateTrustedIdProviderWithAccountParametersOutput {
	return o
}

func (o CreateTrustedIdProviderWithAccountParametersOutput) ToCreateTrustedIdProviderWithAccountParametersOutputWithContext(ctx context.Context) CreateTrustedIdProviderWithAccountParametersOutput {
	return o
}

func (o CreateTrustedIdProviderWithAccountParametersOutput) IdProvider() pulumi.StringOutput {
	return o.ApplyT(func(v CreateTrustedIdProviderWithAccountParameters) string { return v.IdProvider }).(pulumi.StringOutput)
}

func (o CreateTrustedIdProviderWithAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CreateTrustedIdProviderWithAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

type CreateTrustedIdProviderWithAccountParametersArrayOutput struct{ *pulumi.OutputState }

func (CreateTrustedIdProviderWithAccountParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateTrustedIdProviderWithAccountParameters)(nil)).Elem()
}

func (o CreateTrustedIdProviderWithAccountParametersArrayOutput) ToCreateTrustedIdProviderWithAccountParametersArrayOutput() CreateTrustedIdProviderWithAccountParametersArrayOutput {
	return o
}

func (o CreateTrustedIdProviderWithAccountParametersArrayOutput) ToCreateTrustedIdProviderWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateTrustedIdProviderWithAccountParametersArrayOutput {
	return o
}

func (o CreateTrustedIdProviderWithAccountParametersArrayOutput) Index(i pulumi.IntInput) CreateTrustedIdProviderWithAccountParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CreateTrustedIdProviderWithAccountParameters {
		return vs[0].([]CreateTrustedIdProviderWithAccountParameters)[vs[1].(int)]
	}).(CreateTrustedIdProviderWithAccountParametersOutput)
}

type CreateVirtualNetworkRuleWithAccountParameters struct {
	Name     string `pulumi:"name"`
	SubnetId string `pulumi:"subnetId"`
}





type CreateVirtualNetworkRuleWithAccountParametersInput interface {
	pulumi.Input

	ToCreateVirtualNetworkRuleWithAccountParametersOutput() CreateVirtualNetworkRuleWithAccountParametersOutput
	ToCreateVirtualNetworkRuleWithAccountParametersOutputWithContext(context.Context) CreateVirtualNetworkRuleWithAccountParametersOutput
}

type CreateVirtualNetworkRuleWithAccountParametersArgs struct {
	Name     pulumi.StringInput `pulumi:"name"`
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
}

func (CreateVirtualNetworkRuleWithAccountParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateVirtualNetworkRuleWithAccountParameters)(nil)).Elem()
}

func (i CreateVirtualNetworkRuleWithAccountParametersArgs) ToCreateVirtualNetworkRuleWithAccountParametersOutput() CreateVirtualNetworkRuleWithAccountParametersOutput {
	return i.ToCreateVirtualNetworkRuleWithAccountParametersOutputWithContext(context.Background())
}

func (i CreateVirtualNetworkRuleWithAccountParametersArgs) ToCreateVirtualNetworkRuleWithAccountParametersOutputWithContext(ctx context.Context) CreateVirtualNetworkRuleWithAccountParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateVirtualNetworkRuleWithAccountParametersOutput)
}





type CreateVirtualNetworkRuleWithAccountParametersArrayInput interface {
	pulumi.Input

	ToCreateVirtualNetworkRuleWithAccountParametersArrayOutput() CreateVirtualNetworkRuleWithAccountParametersArrayOutput
	ToCreateVirtualNetworkRuleWithAccountParametersArrayOutputWithContext(context.Context) CreateVirtualNetworkRuleWithAccountParametersArrayOutput
}

type CreateVirtualNetworkRuleWithAccountParametersArray []CreateVirtualNetworkRuleWithAccountParametersInput

func (CreateVirtualNetworkRuleWithAccountParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateVirtualNetworkRuleWithAccountParameters)(nil)).Elem()
}

func (i CreateVirtualNetworkRuleWithAccountParametersArray) ToCreateVirtualNetworkRuleWithAccountParametersArrayOutput() CreateVirtualNetworkRuleWithAccountParametersArrayOutput {
	return i.ToCreateVirtualNetworkRuleWithAccountParametersArrayOutputWithContext(context.Background())
}

func (i CreateVirtualNetworkRuleWithAccountParametersArray) ToCreateVirtualNetworkRuleWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateVirtualNetworkRuleWithAccountParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateVirtualNetworkRuleWithAccountParametersArrayOutput)
}

type CreateVirtualNetworkRuleWithAccountParametersOutput struct{ *pulumi.OutputState }

func (CreateVirtualNetworkRuleWithAccountParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateVirtualNetworkRuleWithAccountParameters)(nil)).Elem()
}

func (o CreateVirtualNetworkRuleWithAccountParametersOutput) ToCreateVirtualNetworkRuleWithAccountParametersOutput() CreateVirtualNetworkRuleWithAccountParametersOutput {
	return o
}

func (o CreateVirtualNetworkRuleWithAccountParametersOutput) ToCreateVirtualNetworkRuleWithAccountParametersOutputWithContext(ctx context.Context) CreateVirtualNetworkRuleWithAccountParametersOutput {
	return o
}

func (o CreateVirtualNetworkRuleWithAccountParametersOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CreateVirtualNetworkRuleWithAccountParameters) string { return v.Name }).(pulumi.StringOutput)
}

func (o CreateVirtualNetworkRuleWithAccountParametersOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v CreateVirtualNetworkRuleWithAccountParameters) string { return v.SubnetId }).(pulumi.StringOutput)
}

type CreateVirtualNetworkRuleWithAccountParametersArrayOutput struct{ *pulumi.OutputState }

func (CreateVirtualNetworkRuleWithAccountParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CreateVirtualNetworkRuleWithAccountParameters)(nil)).Elem()
}

func (o CreateVirtualNetworkRuleWithAccountParametersArrayOutput) ToCreateVirtualNetworkRuleWithAccountParametersArrayOutput() CreateVirtualNetworkRuleWithAccountParametersArrayOutput {
	return o
}

func (o CreateVirtualNetworkRuleWithAccountParametersArrayOutput) ToCreateVirtualNetworkRuleWithAccountParametersArrayOutputWithContext(ctx context.Context) CreateVirtualNetworkRuleWithAccountParametersArrayOutput {
	return o
}

func (o CreateVirtualNetworkRuleWithAccountParametersArrayOutput) Index(i pulumi.IntInput) CreateVirtualNetworkRuleWithAccountParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CreateVirtualNetworkRuleWithAccountParameters {
		return vs[0].([]CreateVirtualNetworkRuleWithAccountParameters)[vs[1].(int)]
	}).(CreateVirtualNetworkRuleWithAccountParametersOutput)
}

type EncryptionConfig struct {
	KeyVaultMetaInfo *KeyVaultMetaInfo    `pulumi:"keyVaultMetaInfo"`
	Type             EncryptionConfigType `pulumi:"type"`
}





type EncryptionConfigInput interface {
	pulumi.Input

	ToEncryptionConfigOutput() EncryptionConfigOutput
	ToEncryptionConfigOutputWithContext(context.Context) EncryptionConfigOutput
}

type EncryptionConfigArgs struct {
	KeyVaultMetaInfo KeyVaultMetaInfoPtrInput  `pulumi:"keyVaultMetaInfo"`
	Type             EncryptionConfigTypeInput `pulumi:"type"`
}

func (EncryptionConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionConfig)(nil)).Elem()
}

func (i EncryptionConfigArgs) ToEncryptionConfigOutput() EncryptionConfigOutput {
	return i.ToEncryptionConfigOutputWithContext(context.Background())
}

func (i EncryptionConfigArgs) ToEncryptionConfigOutputWithContext(ctx context.Context) EncryptionConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionConfigOutput)
}

func (i EncryptionConfigArgs) ToEncryptionConfigPtrOutput() EncryptionConfigPtrOutput {
	return i.ToEncryptionConfigPtrOutputWithContext(context.Background())
}

func (i EncryptionConfigArgs) ToEncryptionConfigPtrOutputWithContext(ctx context.Context) EncryptionConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionConfigOutput).ToEncryptionConfigPtrOutputWithContext(ctx)
}









type EncryptionConfigPtrInput interface {
	pulumi.Input

	ToEncryptionConfigPtrOutput() EncryptionConfigPtrOutput
	ToEncryptionConfigPtrOutputWithContext(context.Context) EncryptionConfigPtrOutput
}

type encryptionConfigPtrType EncryptionConfigArgs

func EncryptionConfigPtr(v *EncryptionConfigArgs) EncryptionConfigPtrInput {
	return (*encryptionConfigPtrType)(v)
}

func (*encryptionConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionConfig)(nil)).Elem()
}

func (i *encryptionConfigPtrType) ToEncryptionConfigPtrOutput() EncryptionConfigPtrOutput {
	return i.ToEncryptionConfigPtrOutputWithContext(context.Background())
}

func (i *encryptionConfigPtrType) ToEncryptionConfigPtrOutputWithContext(ctx context.Context) EncryptionConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionConfigPtrOutput)
}

type EncryptionConfigOutput struct{ *pulumi.OutputState }

func (EncryptionConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionConfig)(nil)).Elem()
}

func (o EncryptionConfigOutput) ToEncryptionConfigOutput() EncryptionConfigOutput {
	return o
}

func (o EncryptionConfigOutput) ToEncryptionConfigOutputWithContext(ctx context.Context) EncryptionConfigOutput {
	return o
}

func (o EncryptionConfigOutput) ToEncryptionConfigPtrOutput() EncryptionConfigPtrOutput {
	return o.ToEncryptionConfigPtrOutputWithContext(context.Background())
}

func (o EncryptionConfigOutput) ToEncryptionConfigPtrOutputWithContext(ctx context.Context) EncryptionConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionConfig) *EncryptionConfig {
		return &v
	}).(EncryptionConfigPtrOutput)
}

func (o EncryptionConfigOutput) KeyVaultMetaInfo() KeyVaultMetaInfoPtrOutput {
	return o.ApplyT(func(v EncryptionConfig) *KeyVaultMetaInfo { return v.KeyVaultMetaInfo }).(KeyVaultMetaInfoPtrOutput)
}

func (o EncryptionConfigOutput) Type() EncryptionConfigTypeOutput {
	return o.ApplyT(func(v EncryptionConfig) EncryptionConfigType { return v.Type }).(EncryptionConfigTypeOutput)
}

type EncryptionConfigPtrOutput struct{ *pulumi.OutputState }

func (EncryptionConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionConfig)(nil)).Elem()
}

func (o EncryptionConfigPtrOutput) ToEncryptionConfigPtrOutput() EncryptionConfigPtrOutput {
	return o
}

func (o EncryptionConfigPtrOutput) ToEncryptionConfigPtrOutputWithContext(ctx context.Context) EncryptionConfigPtrOutput {
	return o
}

func (o EncryptionConfigPtrOutput) Elem() EncryptionConfigOutput {
	return o.ApplyT(func(v *EncryptionConfig) EncryptionConfig {
		if v != nil {
			return *v
		}
		var ret EncryptionConfig
		return ret
	}).(EncryptionConfigOutput)
}

func (o EncryptionConfigPtrOutput) KeyVaultMetaInfo() KeyVaultMetaInfoPtrOutput {
	return o.ApplyT(func(v *EncryptionConfig) *KeyVaultMetaInfo {
		if v == nil {
			return nil
		}
		return v.KeyVaultMetaInfo
	}).(KeyVaultMetaInfoPtrOutput)
}

func (o EncryptionConfigPtrOutput) Type() EncryptionConfigTypePtrOutput {
	return o.ApplyT(func(v *EncryptionConfig) *EncryptionConfigType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(EncryptionConfigTypePtrOutput)
}

type EncryptionConfigResponse struct {
	KeyVaultMetaInfo *KeyVaultMetaInfoResponse `pulumi:"keyVaultMetaInfo"`
	Type             string                    `pulumi:"type"`
}

type EncryptionConfigResponseOutput struct{ *pulumi.OutputState }

func (EncryptionConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionConfigResponse)(nil)).Elem()
}

func (o EncryptionConfigResponseOutput) ToEncryptionConfigResponseOutput() EncryptionConfigResponseOutput {
	return o
}

func (o EncryptionConfigResponseOutput) ToEncryptionConfigResponseOutputWithContext(ctx context.Context) EncryptionConfigResponseOutput {
	return o
}

func (o EncryptionConfigResponseOutput) KeyVaultMetaInfo() KeyVaultMetaInfoResponsePtrOutput {
	return o.ApplyT(func(v EncryptionConfigResponse) *KeyVaultMetaInfoResponse { return v.KeyVaultMetaInfo }).(KeyVaultMetaInfoResponsePtrOutput)
}

func (o EncryptionConfigResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionConfigResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EncryptionIdentity struct {
	Type EncryptionIdentityType `pulumi:"type"`
}





type EncryptionIdentityInput interface {
	pulumi.Input

	ToEncryptionIdentityOutput() EncryptionIdentityOutput
	ToEncryptionIdentityOutputWithContext(context.Context) EncryptionIdentityOutput
}

type EncryptionIdentityArgs struct {
	Type EncryptionIdentityTypeInput `pulumi:"type"`
}

func (EncryptionIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionIdentity)(nil)).Elem()
}

func (i EncryptionIdentityArgs) ToEncryptionIdentityOutput() EncryptionIdentityOutput {
	return i.ToEncryptionIdentityOutputWithContext(context.Background())
}

func (i EncryptionIdentityArgs) ToEncryptionIdentityOutputWithContext(ctx context.Context) EncryptionIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionIdentityOutput)
}

func (i EncryptionIdentityArgs) ToEncryptionIdentityPtrOutput() EncryptionIdentityPtrOutput {
	return i.ToEncryptionIdentityPtrOutputWithContext(context.Background())
}

func (i EncryptionIdentityArgs) ToEncryptionIdentityPtrOutputWithContext(ctx context.Context) EncryptionIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionIdentityOutput).ToEncryptionIdentityPtrOutputWithContext(ctx)
}









type EncryptionIdentityPtrInput interface {
	pulumi.Input

	ToEncryptionIdentityPtrOutput() EncryptionIdentityPtrOutput
	ToEncryptionIdentityPtrOutputWithContext(context.Context) EncryptionIdentityPtrOutput
}

type encryptionIdentityPtrType EncryptionIdentityArgs

func EncryptionIdentityPtr(v *EncryptionIdentityArgs) EncryptionIdentityPtrInput {
	return (*encryptionIdentityPtrType)(v)
}

func (*encryptionIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionIdentity)(nil)).Elem()
}

func (i *encryptionIdentityPtrType) ToEncryptionIdentityPtrOutput() EncryptionIdentityPtrOutput {
	return i.ToEncryptionIdentityPtrOutputWithContext(context.Background())
}

func (i *encryptionIdentityPtrType) ToEncryptionIdentityPtrOutputWithContext(ctx context.Context) EncryptionIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionIdentityPtrOutput)
}

type EncryptionIdentityOutput struct{ *pulumi.OutputState }

func (EncryptionIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionIdentity)(nil)).Elem()
}

func (o EncryptionIdentityOutput) ToEncryptionIdentityOutput() EncryptionIdentityOutput {
	return o
}

func (o EncryptionIdentityOutput) ToEncryptionIdentityOutputWithContext(ctx context.Context) EncryptionIdentityOutput {
	return o
}

func (o EncryptionIdentityOutput) ToEncryptionIdentityPtrOutput() EncryptionIdentityPtrOutput {
	return o.ToEncryptionIdentityPtrOutputWithContext(context.Background())
}

func (o EncryptionIdentityOutput) ToEncryptionIdentityPtrOutputWithContext(ctx context.Context) EncryptionIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionIdentity) *EncryptionIdentity {
		return &v
	}).(EncryptionIdentityPtrOutput)
}

func (o EncryptionIdentityOutput) Type() EncryptionIdentityTypeOutput {
	return o.ApplyT(func(v EncryptionIdentity) EncryptionIdentityType { return v.Type }).(EncryptionIdentityTypeOutput)
}

type EncryptionIdentityPtrOutput struct{ *pulumi.OutputState }

func (EncryptionIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionIdentity)(nil)).Elem()
}

func (o EncryptionIdentityPtrOutput) ToEncryptionIdentityPtrOutput() EncryptionIdentityPtrOutput {
	return o
}

func (o EncryptionIdentityPtrOutput) ToEncryptionIdentityPtrOutputWithContext(ctx context.Context) EncryptionIdentityPtrOutput {
	return o
}

func (o EncryptionIdentityPtrOutput) Elem() EncryptionIdentityOutput {
	return o.ApplyT(func(v *EncryptionIdentity) EncryptionIdentity {
		if v != nil {
			return *v
		}
		var ret EncryptionIdentity
		return ret
	}).(EncryptionIdentityOutput)
}

func (o EncryptionIdentityPtrOutput) Type() EncryptionIdentityTypePtrOutput {
	return o.ApplyT(func(v *EncryptionIdentity) *EncryptionIdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(EncryptionIdentityTypePtrOutput)
}

type EncryptionIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type EncryptionIdentityResponseOutput struct{ *pulumi.OutputState }

func (EncryptionIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionIdentityResponse)(nil)).Elem()
}

func (o EncryptionIdentityResponseOutput) ToEncryptionIdentityResponseOutput() EncryptionIdentityResponseOutput {
	return o
}

func (o EncryptionIdentityResponseOutput) ToEncryptionIdentityResponseOutputWithContext(ctx context.Context) EncryptionIdentityResponseOutput {
	return o
}

func (o EncryptionIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EncryptionIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o EncryptionIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type FirewallRuleResponse struct {
	EndIpAddress   string `pulumi:"endIpAddress"`
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	StartIpAddress string `pulumi:"startIpAddress"`
	Type           string `pulumi:"type"`
}

type FirewallRuleResponseOutput struct{ *pulumi.OutputState }

func (FirewallRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallRuleResponse)(nil)).Elem()
}

func (o FirewallRuleResponseOutput) ToFirewallRuleResponseOutput() FirewallRuleResponseOutput {
	return o
}

func (o FirewallRuleResponseOutput) ToFirewallRuleResponseOutputWithContext(ctx context.Context) FirewallRuleResponseOutput {
	return o
}

func (o FirewallRuleResponseOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.EndIpAddress }).(pulumi.StringOutput)
}

func (o FirewallRuleResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o FirewallRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallRuleResponseOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.StartIpAddress }).(pulumi.StringOutput)
}

func (o FirewallRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FirewallRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

type FirewallRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallRuleResponse)(nil)).Elem()
}

func (o FirewallRuleResponseArrayOutput) ToFirewallRuleResponseArrayOutput() FirewallRuleResponseArrayOutput {
	return o
}

func (o FirewallRuleResponseArrayOutput) ToFirewallRuleResponseArrayOutputWithContext(ctx context.Context) FirewallRuleResponseArrayOutput {
	return o
}

func (o FirewallRuleResponseArrayOutput) Index(i pulumi.IntInput) FirewallRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallRuleResponse {
		return vs[0].([]FirewallRuleResponse)[vs[1].(int)]
	}).(FirewallRuleResponseOutput)
}

type KeyVaultMetaInfo struct {
	EncryptionKeyName    string `pulumi:"encryptionKeyName"`
	EncryptionKeyVersion string `pulumi:"encryptionKeyVersion"`
	KeyVaultResourceId   string `pulumi:"keyVaultResourceId"`
}





type KeyVaultMetaInfoInput interface {
	pulumi.Input

	ToKeyVaultMetaInfoOutput() KeyVaultMetaInfoOutput
	ToKeyVaultMetaInfoOutputWithContext(context.Context) KeyVaultMetaInfoOutput
}

type KeyVaultMetaInfoArgs struct {
	EncryptionKeyName    pulumi.StringInput `pulumi:"encryptionKeyName"`
	EncryptionKeyVersion pulumi.StringInput `pulumi:"encryptionKeyVersion"`
	KeyVaultResourceId   pulumi.StringInput `pulumi:"keyVaultResourceId"`
}

func (KeyVaultMetaInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultMetaInfo)(nil)).Elem()
}

func (i KeyVaultMetaInfoArgs) ToKeyVaultMetaInfoOutput() KeyVaultMetaInfoOutput {
	return i.ToKeyVaultMetaInfoOutputWithContext(context.Background())
}

func (i KeyVaultMetaInfoArgs) ToKeyVaultMetaInfoOutputWithContext(ctx context.Context) KeyVaultMetaInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultMetaInfoOutput)
}

func (i KeyVaultMetaInfoArgs) ToKeyVaultMetaInfoPtrOutput() KeyVaultMetaInfoPtrOutput {
	return i.ToKeyVaultMetaInfoPtrOutputWithContext(context.Background())
}

func (i KeyVaultMetaInfoArgs) ToKeyVaultMetaInfoPtrOutputWithContext(ctx context.Context) KeyVaultMetaInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultMetaInfoOutput).ToKeyVaultMetaInfoPtrOutputWithContext(ctx)
}









type KeyVaultMetaInfoPtrInput interface {
	pulumi.Input

	ToKeyVaultMetaInfoPtrOutput() KeyVaultMetaInfoPtrOutput
	ToKeyVaultMetaInfoPtrOutputWithContext(context.Context) KeyVaultMetaInfoPtrOutput
}

type keyVaultMetaInfoPtrType KeyVaultMetaInfoArgs

func KeyVaultMetaInfoPtr(v *KeyVaultMetaInfoArgs) KeyVaultMetaInfoPtrInput {
	return (*keyVaultMetaInfoPtrType)(v)
}

func (*keyVaultMetaInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultMetaInfo)(nil)).Elem()
}

func (i *keyVaultMetaInfoPtrType) ToKeyVaultMetaInfoPtrOutput() KeyVaultMetaInfoPtrOutput {
	return i.ToKeyVaultMetaInfoPtrOutputWithContext(context.Background())
}

func (i *keyVaultMetaInfoPtrType) ToKeyVaultMetaInfoPtrOutputWithContext(ctx context.Context) KeyVaultMetaInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultMetaInfoPtrOutput)
}

type KeyVaultMetaInfoOutput struct{ *pulumi.OutputState }

func (KeyVaultMetaInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultMetaInfo)(nil)).Elem()
}

func (o KeyVaultMetaInfoOutput) ToKeyVaultMetaInfoOutput() KeyVaultMetaInfoOutput {
	return o
}

func (o KeyVaultMetaInfoOutput) ToKeyVaultMetaInfoOutputWithContext(ctx context.Context) KeyVaultMetaInfoOutput {
	return o
}

func (o KeyVaultMetaInfoOutput) ToKeyVaultMetaInfoPtrOutput() KeyVaultMetaInfoPtrOutput {
	return o.ToKeyVaultMetaInfoPtrOutputWithContext(context.Background())
}

func (o KeyVaultMetaInfoOutput) ToKeyVaultMetaInfoPtrOutputWithContext(ctx context.Context) KeyVaultMetaInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultMetaInfo) *KeyVaultMetaInfo {
		return &v
	}).(KeyVaultMetaInfoPtrOutput)
}

func (o KeyVaultMetaInfoOutput) EncryptionKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultMetaInfo) string { return v.EncryptionKeyName }).(pulumi.StringOutput)
}

func (o KeyVaultMetaInfoOutput) EncryptionKeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultMetaInfo) string { return v.EncryptionKeyVersion }).(pulumi.StringOutput)
}

func (o KeyVaultMetaInfoOutput) KeyVaultResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultMetaInfo) string { return v.KeyVaultResourceId }).(pulumi.StringOutput)
}

type KeyVaultMetaInfoPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultMetaInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultMetaInfo)(nil)).Elem()
}

func (o KeyVaultMetaInfoPtrOutput) ToKeyVaultMetaInfoPtrOutput() KeyVaultMetaInfoPtrOutput {
	return o
}

func (o KeyVaultMetaInfoPtrOutput) ToKeyVaultMetaInfoPtrOutputWithContext(ctx context.Context) KeyVaultMetaInfoPtrOutput {
	return o
}

func (o KeyVaultMetaInfoPtrOutput) Elem() KeyVaultMetaInfoOutput {
	return o.ApplyT(func(v *KeyVaultMetaInfo) KeyVaultMetaInfo {
		if v != nil {
			return *v
		}
		var ret KeyVaultMetaInfo
		return ret
	}).(KeyVaultMetaInfoOutput)
}

func (o KeyVaultMetaInfoPtrOutput) EncryptionKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultMetaInfo) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionKeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultMetaInfoPtrOutput) EncryptionKeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultMetaInfo) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionKeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultMetaInfoPtrOutput) KeyVaultResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultMetaInfo) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultResourceId
	}).(pulumi.StringPtrOutput)
}

type KeyVaultMetaInfoResponse struct {
	EncryptionKeyName    string `pulumi:"encryptionKeyName"`
	EncryptionKeyVersion string `pulumi:"encryptionKeyVersion"`
	KeyVaultResourceId   string `pulumi:"keyVaultResourceId"`
}

type KeyVaultMetaInfoResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultMetaInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultMetaInfoResponse)(nil)).Elem()
}

func (o KeyVaultMetaInfoResponseOutput) ToKeyVaultMetaInfoResponseOutput() KeyVaultMetaInfoResponseOutput {
	return o
}

func (o KeyVaultMetaInfoResponseOutput) ToKeyVaultMetaInfoResponseOutputWithContext(ctx context.Context) KeyVaultMetaInfoResponseOutput {
	return o
}

func (o KeyVaultMetaInfoResponseOutput) EncryptionKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultMetaInfoResponse) string { return v.EncryptionKeyName }).(pulumi.StringOutput)
}

func (o KeyVaultMetaInfoResponseOutput) EncryptionKeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultMetaInfoResponse) string { return v.EncryptionKeyVersion }).(pulumi.StringOutput)
}

func (o KeyVaultMetaInfoResponseOutput) KeyVaultResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultMetaInfoResponse) string { return v.KeyVaultResourceId }).(pulumi.StringOutput)
}

type KeyVaultMetaInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultMetaInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultMetaInfoResponse)(nil)).Elem()
}

func (o KeyVaultMetaInfoResponsePtrOutput) ToKeyVaultMetaInfoResponsePtrOutput() KeyVaultMetaInfoResponsePtrOutput {
	return o
}

func (o KeyVaultMetaInfoResponsePtrOutput) ToKeyVaultMetaInfoResponsePtrOutputWithContext(ctx context.Context) KeyVaultMetaInfoResponsePtrOutput {
	return o
}

func (o KeyVaultMetaInfoResponsePtrOutput) Elem() KeyVaultMetaInfoResponseOutput {
	return o.ApplyT(func(v *KeyVaultMetaInfoResponse) KeyVaultMetaInfoResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultMetaInfoResponse
		return ret
	}).(KeyVaultMetaInfoResponseOutput)
}

func (o KeyVaultMetaInfoResponsePtrOutput) EncryptionKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultMetaInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionKeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultMetaInfoResponsePtrOutput) EncryptionKeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultMetaInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionKeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultMetaInfoResponsePtrOutput) KeyVaultResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultMetaInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultResourceId
	}).(pulumi.StringPtrOutput)
}

type TrustedIdProviderResponse struct {
	Id         string `pulumi:"id"`
	IdProvider string `pulumi:"idProvider"`
	Name       string `pulumi:"name"`
	Type       string `pulumi:"type"`
}

type TrustedIdProviderResponseOutput struct{ *pulumi.OutputState }

func (TrustedIdProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedIdProviderResponse)(nil)).Elem()
}

func (o TrustedIdProviderResponseOutput) ToTrustedIdProviderResponseOutput() TrustedIdProviderResponseOutput {
	return o
}

func (o TrustedIdProviderResponseOutput) ToTrustedIdProviderResponseOutputWithContext(ctx context.Context) TrustedIdProviderResponseOutput {
	return o
}

func (o TrustedIdProviderResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TrustedIdProviderResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o TrustedIdProviderResponseOutput) IdProvider() pulumi.StringOutput {
	return o.ApplyT(func(v TrustedIdProviderResponse) string { return v.IdProvider }).(pulumi.StringOutput)
}

func (o TrustedIdProviderResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TrustedIdProviderResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TrustedIdProviderResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TrustedIdProviderResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TrustedIdProviderResponseArrayOutput struct{ *pulumi.OutputState }

func (TrustedIdProviderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedIdProviderResponse)(nil)).Elem()
}

func (o TrustedIdProviderResponseArrayOutput) ToTrustedIdProviderResponseArrayOutput() TrustedIdProviderResponseArrayOutput {
	return o
}

func (o TrustedIdProviderResponseArrayOutput) ToTrustedIdProviderResponseArrayOutputWithContext(ctx context.Context) TrustedIdProviderResponseArrayOutput {
	return o
}

func (o TrustedIdProviderResponseArrayOutput) Index(i pulumi.IntInput) TrustedIdProviderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrustedIdProviderResponse {
		return vs[0].([]TrustedIdProviderResponse)[vs[1].(int)]
	}).(TrustedIdProviderResponseOutput)
}

type VirtualNetworkRuleResponse struct {
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	SubnetId string `pulumi:"subnetId"`
	Type     string `pulumi:"type"`
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CreateFirewallRuleWithAccountParametersOutput{})
	pulumi.RegisterOutputType(CreateFirewallRuleWithAccountParametersArrayOutput{})
	pulumi.RegisterOutputType(CreateTrustedIdProviderWithAccountParametersOutput{})
	pulumi.RegisterOutputType(CreateTrustedIdProviderWithAccountParametersArrayOutput{})
	pulumi.RegisterOutputType(CreateVirtualNetworkRuleWithAccountParametersOutput{})
	pulumi.RegisterOutputType(CreateVirtualNetworkRuleWithAccountParametersArrayOutput{})
	pulumi.RegisterOutputType(EncryptionConfigOutput{})
	pulumi.RegisterOutputType(EncryptionConfigPtrOutput{})
	pulumi.RegisterOutputType(EncryptionConfigResponseOutput{})
	pulumi.RegisterOutputType(EncryptionIdentityOutput{})
	pulumi.RegisterOutputType(EncryptionIdentityPtrOutput{})
	pulumi.RegisterOutputType(EncryptionIdentityResponseOutput{})
	pulumi.RegisterOutputType(FirewallRuleResponseOutput{})
	pulumi.RegisterOutputType(FirewallRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultMetaInfoOutput{})
	pulumi.RegisterOutputType(KeyVaultMetaInfoPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultMetaInfoResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultMetaInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(TrustedIdProviderResponseOutput{})
	pulumi.RegisterOutputType(TrustedIdProviderResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
