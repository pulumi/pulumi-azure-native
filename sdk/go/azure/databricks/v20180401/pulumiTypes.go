


package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddressSpace struct {
	AddressPrefixes []string `pulumi:"addressPrefixes"`
}





type AddressSpaceInput interface {
	pulumi.Input

	ToAddressSpaceOutput() AddressSpaceOutput
	ToAddressSpaceOutputWithContext(context.Context) AddressSpaceOutput
}

type AddressSpaceArgs struct {
	AddressPrefixes pulumi.StringArrayInput `pulumi:"addressPrefixes"`
}

func (AddressSpaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpace)(nil)).Elem()
}

func (i AddressSpaceArgs) ToAddressSpaceOutput() AddressSpaceOutput {
	return i.ToAddressSpaceOutputWithContext(context.Background())
}

func (i AddressSpaceArgs) ToAddressSpaceOutputWithContext(ctx context.Context) AddressSpaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceOutput)
}

func (i AddressSpaceArgs) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return i.ToAddressSpacePtrOutputWithContext(context.Background())
}

func (i AddressSpaceArgs) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceOutput).ToAddressSpacePtrOutputWithContext(ctx)
}









type AddressSpacePtrInput interface {
	pulumi.Input

	ToAddressSpacePtrOutput() AddressSpacePtrOutput
	ToAddressSpacePtrOutputWithContext(context.Context) AddressSpacePtrOutput
}

type addressSpacePtrType AddressSpaceArgs

func AddressSpacePtr(v *AddressSpaceArgs) AddressSpacePtrInput {
	return (*addressSpacePtrType)(v)
}

func (*addressSpacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpace)(nil)).Elem()
}

func (i *addressSpacePtrType) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return i.ToAddressSpacePtrOutputWithContext(context.Background())
}

func (i *addressSpacePtrType) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpacePtrOutput)
}

type AddressSpaceOutput struct{ *pulumi.OutputState }

func (AddressSpaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpace)(nil)).Elem()
}

func (o AddressSpaceOutput) ToAddressSpaceOutput() AddressSpaceOutput {
	return o
}

func (o AddressSpaceOutput) ToAddressSpaceOutputWithContext(ctx context.Context) AddressSpaceOutput {
	return o
}

func (o AddressSpaceOutput) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return o.ToAddressSpacePtrOutputWithContext(context.Background())
}

func (o AddressSpaceOutput) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressSpace) *AddressSpace {
		return &v
	}).(AddressSpacePtrOutput)
}

func (o AddressSpaceOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddressSpace) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

type AddressSpacePtrOutput struct{ *pulumi.OutputState }

func (AddressSpacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpace)(nil)).Elem()
}

func (o AddressSpacePtrOutput) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return o
}

func (o AddressSpacePtrOutput) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return o
}

func (o AddressSpacePtrOutput) Elem() AddressSpaceOutput {
	return o.ApplyT(func(v *AddressSpace) AddressSpace {
		if v != nil {
			return *v
		}
		var ret AddressSpace
		return ret
	}).(AddressSpaceOutput)
}

func (o AddressSpacePtrOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AddressSpace) []string {
		if v == nil {
			return nil
		}
		return v.AddressPrefixes
	}).(pulumi.StringArrayOutput)
}

type AddressSpaceResponse struct {
	AddressPrefixes []string `pulumi:"addressPrefixes"`
}

type AddressSpaceResponseOutput struct{ *pulumi.OutputState }

func (AddressSpaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpaceResponse)(nil)).Elem()
}

func (o AddressSpaceResponseOutput) ToAddressSpaceResponseOutput() AddressSpaceResponseOutput {
	return o
}

func (o AddressSpaceResponseOutput) ToAddressSpaceResponseOutputWithContext(ctx context.Context) AddressSpaceResponseOutput {
	return o
}

func (o AddressSpaceResponseOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddressSpaceResponse) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

type AddressSpaceResponsePtrOutput struct{ *pulumi.OutputState }

func (AddressSpaceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpaceResponse)(nil)).Elem()
}

func (o AddressSpaceResponsePtrOutput) ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput {
	return o
}

func (o AddressSpaceResponsePtrOutput) ToAddressSpaceResponsePtrOutputWithContext(ctx context.Context) AddressSpaceResponsePtrOutput {
	return o
}

func (o AddressSpaceResponsePtrOutput) Elem() AddressSpaceResponseOutput {
	return o.ApplyT(func(v *AddressSpaceResponse) AddressSpaceResponse {
		if v != nil {
			return *v
		}
		var ret AddressSpaceResponse
		return ret
	}).(AddressSpaceResponseOutput)
}

func (o AddressSpaceResponsePtrOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AddressSpaceResponse) []string {
		if v == nil {
			return nil
		}
		return v.AddressPrefixes
	}).(pulumi.StringArrayOutput)
}

type CreatedByResponse struct {
	ApplicationId string `pulumi:"applicationId"`
	Oid           string `pulumi:"oid"`
	Puid          string `pulumi:"puid"`
}

type CreatedByResponseOutput struct{ *pulumi.OutputState }

func (CreatedByResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatedByResponse)(nil)).Elem()
}

func (o CreatedByResponseOutput) ToCreatedByResponseOutput() CreatedByResponseOutput {
	return o
}

func (o CreatedByResponseOutput) ToCreatedByResponseOutputWithContext(ctx context.Context) CreatedByResponseOutput {
	return o
}

func (o CreatedByResponseOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v CreatedByResponse) string { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o CreatedByResponseOutput) Oid() pulumi.StringOutput {
	return o.ApplyT(func(v CreatedByResponse) string { return v.Oid }).(pulumi.StringOutput)
}

func (o CreatedByResponseOutput) Puid() pulumi.StringOutput {
	return o.ApplyT(func(v CreatedByResponse) string { return v.Puid }).(pulumi.StringOutput)
}

type CreatedByResponsePtrOutput struct{ *pulumi.OutputState }

func (CreatedByResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreatedByResponse)(nil)).Elem()
}

func (o CreatedByResponsePtrOutput) ToCreatedByResponsePtrOutput() CreatedByResponsePtrOutput {
	return o
}

func (o CreatedByResponsePtrOutput) ToCreatedByResponsePtrOutputWithContext(ctx context.Context) CreatedByResponsePtrOutput {
	return o
}

func (o CreatedByResponsePtrOutput) Elem() CreatedByResponseOutput {
	return o.ApplyT(func(v *CreatedByResponse) CreatedByResponse {
		if v != nil {
			return *v
		}
		var ret CreatedByResponse
		return ret
	}).(CreatedByResponseOutput)
}

func (o CreatedByResponsePtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreatedByResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o CreatedByResponsePtrOutput) Oid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreatedByResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Oid
	}).(pulumi.StringPtrOutput)
}

func (o CreatedByResponsePtrOutput) Puid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreatedByResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Puid
	}).(pulumi.StringPtrOutput)
}

type Encryption struct {
	KeyName     *string `pulumi:"keyName"`
	KeySource   *string `pulumi:"keySource"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}


func (val *Encryption) Defaults() *Encryption {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		keySource_ := "Default"
		tmp.KeySource = &keySource_
	}
	return &tmp
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeySource   pulumi.StringPtrInput `pulumi:"keySource"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}


func (val *EncryptionArgs) Defaults() *EncryptionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		tmp.KeySource = pulumi.StringPtr("Default")
	}
	return &tmp
}
func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o EncryptionOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o EncryptionOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type EncryptionResponse struct {
	KeyName     *string `pulumi:"keyName"`
	KeySource   *string `pulumi:"keySource"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}


func (val *EncryptionResponse) Defaults() *EncryptionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.KeySource) {
		keySource_ := "Default"
		tmp.KeySource = &keySource_
	}
	return &tmp
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o EncryptionResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o EncryptionResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityConfigurationResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type ManagedIdentityConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityConfigurationResponse)(nil)).Elem()
}

func (o ManagedIdentityConfigurationResponseOutput) ToManagedIdentityConfigurationResponseOutput() ManagedIdentityConfigurationResponseOutput {
	return o
}

func (o ManagedIdentityConfigurationResponseOutput) ToManagedIdentityConfigurationResponseOutputWithContext(ctx context.Context) ManagedIdentityConfigurationResponseOutput {
	return o
}

func (o ManagedIdentityConfigurationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityConfigurationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedIdentityConfigurationResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityConfigurationResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedIdentityConfigurationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityConfigurationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ManagedIdentityConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityConfigurationResponse)(nil)).Elem()
}

func (o ManagedIdentityConfigurationResponsePtrOutput) ToManagedIdentityConfigurationResponsePtrOutput() ManagedIdentityConfigurationResponsePtrOutput {
	return o
}

func (o ManagedIdentityConfigurationResponsePtrOutput) ToManagedIdentityConfigurationResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityConfigurationResponsePtrOutput {
	return o
}

func (o ManagedIdentityConfigurationResponsePtrOutput) Elem() ManagedIdentityConfigurationResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityConfigurationResponse) ManagedIdentityConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityConfigurationResponse
		return ret
	}).(ManagedIdentityConfigurationResponseOutput)
}

func (o ManagedIdentityConfigurationResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityConfigurationResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityConfigurationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork struct {
	Id *string `pulumi:"id"`
}





type VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput
	ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutputWithContext(context.Context) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput
}

type VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork)(nil)).Elem()
}

func (i VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput)
}

func (i VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput).ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutputWithContext(ctx)
}









type VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput
	ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutputWithContext(context.Context) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput
}

type virtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrType VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkArgs

func VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtr(v *VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkArgs) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrInput {
	return (*virtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrType)(v)
}

func (*virtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork)(nil)).Elem()
}

func (i *virtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrType) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrType) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput {
	return o.ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork) *VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork {
		return &v
	}).(VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput)
}

func (o VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput) ToVirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput) Elem() VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork) VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork
		return ret
	}).(VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput)
}

func (o VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork struct {
	Id *string `pulumi:"id"`
}





type VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput
	ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutputWithContext(context.Context) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput
}

type VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork)(nil)).Elem()
}

func (i VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput)
}

type VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork struct {
	Id *string `pulumi:"id"`
}

type VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput) Elem() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork
		return ret
	}).(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput)
}

func (o VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork struct {
	Id *string `pulumi:"id"`
}

type VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type WorkspaceCustomBooleanParameter struct {
	Value bool `pulumi:"value"`
}





type WorkspaceCustomBooleanParameterInput interface {
	pulumi.Input

	ToWorkspaceCustomBooleanParameterOutput() WorkspaceCustomBooleanParameterOutput
	ToWorkspaceCustomBooleanParameterOutputWithContext(context.Context) WorkspaceCustomBooleanParameterOutput
}

type WorkspaceCustomBooleanParameterArgs struct {
	Value pulumi.BoolInput `pulumi:"value"`
}

func (WorkspaceCustomBooleanParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomBooleanParameter)(nil)).Elem()
}

func (i WorkspaceCustomBooleanParameterArgs) ToWorkspaceCustomBooleanParameterOutput() WorkspaceCustomBooleanParameterOutput {
	return i.ToWorkspaceCustomBooleanParameterOutputWithContext(context.Background())
}

func (i WorkspaceCustomBooleanParameterArgs) ToWorkspaceCustomBooleanParameterOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomBooleanParameterOutput)
}

func (i WorkspaceCustomBooleanParameterArgs) ToWorkspaceCustomBooleanParameterPtrOutput() WorkspaceCustomBooleanParameterPtrOutput {
	return i.ToWorkspaceCustomBooleanParameterPtrOutputWithContext(context.Background())
}

func (i WorkspaceCustomBooleanParameterArgs) ToWorkspaceCustomBooleanParameterPtrOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomBooleanParameterOutput).ToWorkspaceCustomBooleanParameterPtrOutputWithContext(ctx)
}









type WorkspaceCustomBooleanParameterPtrInput interface {
	pulumi.Input

	ToWorkspaceCustomBooleanParameterPtrOutput() WorkspaceCustomBooleanParameterPtrOutput
	ToWorkspaceCustomBooleanParameterPtrOutputWithContext(context.Context) WorkspaceCustomBooleanParameterPtrOutput
}

type workspaceCustomBooleanParameterPtrType WorkspaceCustomBooleanParameterArgs

func WorkspaceCustomBooleanParameterPtr(v *WorkspaceCustomBooleanParameterArgs) WorkspaceCustomBooleanParameterPtrInput {
	return (*workspaceCustomBooleanParameterPtrType)(v)
}

func (*workspaceCustomBooleanParameterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomBooleanParameter)(nil)).Elem()
}

func (i *workspaceCustomBooleanParameterPtrType) ToWorkspaceCustomBooleanParameterPtrOutput() WorkspaceCustomBooleanParameterPtrOutput {
	return i.ToWorkspaceCustomBooleanParameterPtrOutputWithContext(context.Background())
}

func (i *workspaceCustomBooleanParameterPtrType) ToWorkspaceCustomBooleanParameterPtrOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomBooleanParameterPtrOutput)
}

type WorkspaceCustomBooleanParameterOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomBooleanParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomBooleanParameter)(nil)).Elem()
}

func (o WorkspaceCustomBooleanParameterOutput) ToWorkspaceCustomBooleanParameterOutput() WorkspaceCustomBooleanParameterOutput {
	return o
}

func (o WorkspaceCustomBooleanParameterOutput) ToWorkspaceCustomBooleanParameterOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterOutput {
	return o
}

func (o WorkspaceCustomBooleanParameterOutput) ToWorkspaceCustomBooleanParameterPtrOutput() WorkspaceCustomBooleanParameterPtrOutput {
	return o.ToWorkspaceCustomBooleanParameterPtrOutputWithContext(context.Background())
}

func (o WorkspaceCustomBooleanParameterOutput) ToWorkspaceCustomBooleanParameterPtrOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCustomBooleanParameter) *WorkspaceCustomBooleanParameter {
		return &v
	}).(WorkspaceCustomBooleanParameterPtrOutput)
}

func (o WorkspaceCustomBooleanParameterOutput) Value() pulumi.BoolOutput {
	return o.ApplyT(func(v WorkspaceCustomBooleanParameter) bool { return v.Value }).(pulumi.BoolOutput)
}

type WorkspaceCustomBooleanParameterPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomBooleanParameterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomBooleanParameter)(nil)).Elem()
}

func (o WorkspaceCustomBooleanParameterPtrOutput) ToWorkspaceCustomBooleanParameterPtrOutput() WorkspaceCustomBooleanParameterPtrOutput {
	return o
}

func (o WorkspaceCustomBooleanParameterPtrOutput) ToWorkspaceCustomBooleanParameterPtrOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterPtrOutput {
	return o
}

func (o WorkspaceCustomBooleanParameterPtrOutput) Elem() WorkspaceCustomBooleanParameterOutput {
	return o.ApplyT(func(v *WorkspaceCustomBooleanParameter) WorkspaceCustomBooleanParameter {
		if v != nil {
			return *v
		}
		var ret WorkspaceCustomBooleanParameter
		return ret
	}).(WorkspaceCustomBooleanParameterOutput)
}

func (o WorkspaceCustomBooleanParameterPtrOutput) Value() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomBooleanParameter) *bool {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.BoolPtrOutput)
}

type WorkspaceCustomBooleanParameterResponse struct {
	Type  string `pulumi:"type"`
	Value bool   `pulumi:"value"`
}

type WorkspaceCustomBooleanParameterResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomBooleanParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomBooleanParameterResponse)(nil)).Elem()
}

func (o WorkspaceCustomBooleanParameterResponseOutput) ToWorkspaceCustomBooleanParameterResponseOutput() WorkspaceCustomBooleanParameterResponseOutput {
	return o
}

func (o WorkspaceCustomBooleanParameterResponseOutput) ToWorkspaceCustomBooleanParameterResponseOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterResponseOutput {
	return o
}

func (o WorkspaceCustomBooleanParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCustomBooleanParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o WorkspaceCustomBooleanParameterResponseOutput) Value() pulumi.BoolOutput {
	return o.ApplyT(func(v WorkspaceCustomBooleanParameterResponse) bool { return v.Value }).(pulumi.BoolOutput)
}

type WorkspaceCustomBooleanParameterResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomBooleanParameterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomBooleanParameterResponse)(nil)).Elem()
}

func (o WorkspaceCustomBooleanParameterResponsePtrOutput) ToWorkspaceCustomBooleanParameterResponsePtrOutput() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o
}

func (o WorkspaceCustomBooleanParameterResponsePtrOutput) ToWorkspaceCustomBooleanParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o
}

func (o WorkspaceCustomBooleanParameterResponsePtrOutput) Elem() WorkspaceCustomBooleanParameterResponseOutput {
	return o.ApplyT(func(v *WorkspaceCustomBooleanParameterResponse) WorkspaceCustomBooleanParameterResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceCustomBooleanParameterResponse
		return ret
	}).(WorkspaceCustomBooleanParameterResponseOutput)
}

func (o WorkspaceCustomBooleanParameterResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomBooleanParameterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceCustomBooleanParameterResponsePtrOutput) Value() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomBooleanParameterResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.BoolPtrOutput)
}

type WorkspaceCustomObjectParameterResponse struct {
	Type  string      `pulumi:"type"`
	Value interface{} `pulumi:"value"`
}

type WorkspaceCustomObjectParameterResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomObjectParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomObjectParameterResponse)(nil)).Elem()
}

func (o WorkspaceCustomObjectParameterResponseOutput) ToWorkspaceCustomObjectParameterResponseOutput() WorkspaceCustomObjectParameterResponseOutput {
	return o
}

func (o WorkspaceCustomObjectParameterResponseOutput) ToWorkspaceCustomObjectParameterResponseOutputWithContext(ctx context.Context) WorkspaceCustomObjectParameterResponseOutput {
	return o
}

func (o WorkspaceCustomObjectParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCustomObjectParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o WorkspaceCustomObjectParameterResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkspaceCustomObjectParameterResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type WorkspaceCustomObjectParameterResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomObjectParameterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomObjectParameterResponse)(nil)).Elem()
}

func (o WorkspaceCustomObjectParameterResponsePtrOutput) ToWorkspaceCustomObjectParameterResponsePtrOutput() WorkspaceCustomObjectParameterResponsePtrOutput {
	return o
}

func (o WorkspaceCustomObjectParameterResponsePtrOutput) ToWorkspaceCustomObjectParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomObjectParameterResponsePtrOutput {
	return o
}

func (o WorkspaceCustomObjectParameterResponsePtrOutput) Elem() WorkspaceCustomObjectParameterResponseOutput {
	return o.ApplyT(func(v *WorkspaceCustomObjectParameterResponse) WorkspaceCustomObjectParameterResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceCustomObjectParameterResponse
		return ret
	}).(WorkspaceCustomObjectParameterResponseOutput)
}

func (o WorkspaceCustomObjectParameterResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomObjectParameterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceCustomObjectParameterResponsePtrOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v *WorkspaceCustomObjectParameterResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.AnyOutput)
}

type WorkspaceCustomParameters struct {
	AmlWorkspaceId                  *WorkspaceCustomStringParameter  `pulumi:"amlWorkspaceId"`
	CustomPrivateSubnetName         *WorkspaceCustomStringParameter  `pulumi:"customPrivateSubnetName"`
	CustomPublicSubnetName          *WorkspaceCustomStringParameter  `pulumi:"customPublicSubnetName"`
	CustomVirtualNetworkId          *WorkspaceCustomStringParameter  `pulumi:"customVirtualNetworkId"`
	EnableNoPublicIp                *WorkspaceCustomBooleanParameter `pulumi:"enableNoPublicIp"`
	Encryption                      *WorkspaceEncryptionParameter    `pulumi:"encryption"`
	LoadBalancerBackendPoolName     *WorkspaceCustomStringParameter  `pulumi:"loadBalancerBackendPoolName"`
	LoadBalancerId                  *WorkspaceCustomStringParameter  `pulumi:"loadBalancerId"`
	NatGatewayName                  *WorkspaceCustomStringParameter  `pulumi:"natGatewayName"`
	PrepareEncryption               *WorkspaceCustomBooleanParameter `pulumi:"prepareEncryption"`
	PublicIpName                    *WorkspaceCustomStringParameter  `pulumi:"publicIpName"`
	RequireInfrastructureEncryption *WorkspaceCustomBooleanParameter `pulumi:"requireInfrastructureEncryption"`
	StorageAccountName              *WorkspaceCustomStringParameter  `pulumi:"storageAccountName"`
	StorageAccountSkuName           *WorkspaceCustomStringParameter  `pulumi:"storageAccountSkuName"`
	VnetAddressPrefix               *WorkspaceCustomStringParameter  `pulumi:"vnetAddressPrefix"`
}


func (val *WorkspaceCustomParameters) Defaults() *WorkspaceCustomParameters {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = tmp.Encryption.Defaults()

	return &tmp
}





type WorkspaceCustomParametersInput interface {
	pulumi.Input

	ToWorkspaceCustomParametersOutput() WorkspaceCustomParametersOutput
	ToWorkspaceCustomParametersOutputWithContext(context.Context) WorkspaceCustomParametersOutput
}

type WorkspaceCustomParametersArgs struct {
	AmlWorkspaceId                  WorkspaceCustomStringParameterPtrInput  `pulumi:"amlWorkspaceId"`
	CustomPrivateSubnetName         WorkspaceCustomStringParameterPtrInput  `pulumi:"customPrivateSubnetName"`
	CustomPublicSubnetName          WorkspaceCustomStringParameterPtrInput  `pulumi:"customPublicSubnetName"`
	CustomVirtualNetworkId          WorkspaceCustomStringParameterPtrInput  `pulumi:"customVirtualNetworkId"`
	EnableNoPublicIp                WorkspaceCustomBooleanParameterPtrInput `pulumi:"enableNoPublicIp"`
	Encryption                      WorkspaceEncryptionParameterPtrInput    `pulumi:"encryption"`
	LoadBalancerBackendPoolName     WorkspaceCustomStringParameterPtrInput  `pulumi:"loadBalancerBackendPoolName"`
	LoadBalancerId                  WorkspaceCustomStringParameterPtrInput  `pulumi:"loadBalancerId"`
	NatGatewayName                  WorkspaceCustomStringParameterPtrInput  `pulumi:"natGatewayName"`
	PrepareEncryption               WorkspaceCustomBooleanParameterPtrInput `pulumi:"prepareEncryption"`
	PublicIpName                    WorkspaceCustomStringParameterPtrInput  `pulumi:"publicIpName"`
	RequireInfrastructureEncryption WorkspaceCustomBooleanParameterPtrInput `pulumi:"requireInfrastructureEncryption"`
	StorageAccountName              WorkspaceCustomStringParameterPtrInput  `pulumi:"storageAccountName"`
	StorageAccountSkuName           WorkspaceCustomStringParameterPtrInput  `pulumi:"storageAccountSkuName"`
	VnetAddressPrefix               WorkspaceCustomStringParameterPtrInput  `pulumi:"vnetAddressPrefix"`
}


func (val *WorkspaceCustomParametersArgs) Defaults() *WorkspaceCustomParametersArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (WorkspaceCustomParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomParameters)(nil)).Elem()
}

func (i WorkspaceCustomParametersArgs) ToWorkspaceCustomParametersOutput() WorkspaceCustomParametersOutput {
	return i.ToWorkspaceCustomParametersOutputWithContext(context.Background())
}

func (i WorkspaceCustomParametersArgs) ToWorkspaceCustomParametersOutputWithContext(ctx context.Context) WorkspaceCustomParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomParametersOutput)
}

func (i WorkspaceCustomParametersArgs) ToWorkspaceCustomParametersPtrOutput() WorkspaceCustomParametersPtrOutput {
	return i.ToWorkspaceCustomParametersPtrOutputWithContext(context.Background())
}

func (i WorkspaceCustomParametersArgs) ToWorkspaceCustomParametersPtrOutputWithContext(ctx context.Context) WorkspaceCustomParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomParametersOutput).ToWorkspaceCustomParametersPtrOutputWithContext(ctx)
}









type WorkspaceCustomParametersPtrInput interface {
	pulumi.Input

	ToWorkspaceCustomParametersPtrOutput() WorkspaceCustomParametersPtrOutput
	ToWorkspaceCustomParametersPtrOutputWithContext(context.Context) WorkspaceCustomParametersPtrOutput
}

type workspaceCustomParametersPtrType WorkspaceCustomParametersArgs

func WorkspaceCustomParametersPtr(v *WorkspaceCustomParametersArgs) WorkspaceCustomParametersPtrInput {
	return (*workspaceCustomParametersPtrType)(v)
}

func (*workspaceCustomParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomParameters)(nil)).Elem()
}

func (i *workspaceCustomParametersPtrType) ToWorkspaceCustomParametersPtrOutput() WorkspaceCustomParametersPtrOutput {
	return i.ToWorkspaceCustomParametersPtrOutputWithContext(context.Background())
}

func (i *workspaceCustomParametersPtrType) ToWorkspaceCustomParametersPtrOutputWithContext(ctx context.Context) WorkspaceCustomParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomParametersPtrOutput)
}

type WorkspaceCustomParametersOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomParameters)(nil)).Elem()
}

func (o WorkspaceCustomParametersOutput) ToWorkspaceCustomParametersOutput() WorkspaceCustomParametersOutput {
	return o
}

func (o WorkspaceCustomParametersOutput) ToWorkspaceCustomParametersOutputWithContext(ctx context.Context) WorkspaceCustomParametersOutput {
	return o
}

func (o WorkspaceCustomParametersOutput) ToWorkspaceCustomParametersPtrOutput() WorkspaceCustomParametersPtrOutput {
	return o.ToWorkspaceCustomParametersPtrOutputWithContext(context.Background())
}

func (o WorkspaceCustomParametersOutput) ToWorkspaceCustomParametersPtrOutputWithContext(ctx context.Context) WorkspaceCustomParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCustomParameters) *WorkspaceCustomParameters {
		return &v
	}).(WorkspaceCustomParametersPtrOutput)
}

func (o WorkspaceCustomParametersOutput) AmlWorkspaceId() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.AmlWorkspaceId }).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) CustomPrivateSubnetName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.CustomPrivateSubnetName }).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) CustomPublicSubnetName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.CustomPublicSubnetName }).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) CustomVirtualNetworkId() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.CustomVirtualNetworkId }).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) EnableNoPublicIp() WorkspaceCustomBooleanParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomBooleanParameter { return v.EnableNoPublicIp }).(WorkspaceCustomBooleanParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) Encryption() WorkspaceEncryptionParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceEncryptionParameter { return v.Encryption }).(WorkspaceEncryptionParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) LoadBalancerBackendPoolName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		return v.LoadBalancerBackendPoolName
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) LoadBalancerId() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.LoadBalancerId }).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) NatGatewayName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.NatGatewayName }).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) PrepareEncryption() WorkspaceCustomBooleanParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomBooleanParameter { return v.PrepareEncryption }).(WorkspaceCustomBooleanParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) PublicIpName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.PublicIpName }).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) RequireInfrastructureEncryption() WorkspaceCustomBooleanParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomBooleanParameter {
		return v.RequireInfrastructureEncryption
	}).(WorkspaceCustomBooleanParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) StorageAccountName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.StorageAccountName }).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) StorageAccountSkuName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.StorageAccountSkuName }).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersOutput) VnetAddressPrefix() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParameters) *WorkspaceCustomStringParameter { return v.VnetAddressPrefix }).(WorkspaceCustomStringParameterPtrOutput)
}

type WorkspaceCustomParametersPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomParameters)(nil)).Elem()
}

func (o WorkspaceCustomParametersPtrOutput) ToWorkspaceCustomParametersPtrOutput() WorkspaceCustomParametersPtrOutput {
	return o
}

func (o WorkspaceCustomParametersPtrOutput) ToWorkspaceCustomParametersPtrOutputWithContext(ctx context.Context) WorkspaceCustomParametersPtrOutput {
	return o
}

func (o WorkspaceCustomParametersPtrOutput) Elem() WorkspaceCustomParametersOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) WorkspaceCustomParameters {
		if v != nil {
			return *v
		}
		var ret WorkspaceCustomParameters
		return ret
	}).(WorkspaceCustomParametersOutput)
}

func (o WorkspaceCustomParametersPtrOutput) AmlWorkspaceId() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.AmlWorkspaceId
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) CustomPrivateSubnetName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.CustomPrivateSubnetName
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) CustomPublicSubnetName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.CustomPublicSubnetName
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) CustomVirtualNetworkId() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.CustomVirtualNetworkId
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) EnableNoPublicIp() WorkspaceCustomBooleanParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomBooleanParameter {
		if v == nil {
			return nil
		}
		return v.EnableNoPublicIp
	}).(WorkspaceCustomBooleanParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) Encryption() WorkspaceEncryptionParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceEncryptionParameter {
		if v == nil {
			return nil
		}
		return v.Encryption
	}).(WorkspaceEncryptionParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) LoadBalancerBackendPoolName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.LoadBalancerBackendPoolName
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) LoadBalancerId() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.LoadBalancerId
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) NatGatewayName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.NatGatewayName
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) PrepareEncryption() WorkspaceCustomBooleanParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomBooleanParameter {
		if v == nil {
			return nil
		}
		return v.PrepareEncryption
	}).(WorkspaceCustomBooleanParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) PublicIpName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.PublicIpName
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) RequireInfrastructureEncryption() WorkspaceCustomBooleanParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomBooleanParameter {
		if v == nil {
			return nil
		}
		return v.RequireInfrastructureEncryption
	}).(WorkspaceCustomBooleanParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) StorageAccountName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) StorageAccountSkuName() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.StorageAccountSkuName
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomParametersPtrOutput) VnetAddressPrefix() WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomStringParameter {
		if v == nil {
			return nil
		}
		return v.VnetAddressPrefix
	}).(WorkspaceCustomStringParameterPtrOutput)
}

type WorkspaceCustomParametersResponse struct {
	AmlWorkspaceId                  *WorkspaceCustomStringParameterResponse  `pulumi:"amlWorkspaceId"`
	CustomPrivateSubnetName         *WorkspaceCustomStringParameterResponse  `pulumi:"customPrivateSubnetName"`
	CustomPublicSubnetName          *WorkspaceCustomStringParameterResponse  `pulumi:"customPublicSubnetName"`
	CustomVirtualNetworkId          *WorkspaceCustomStringParameterResponse  `pulumi:"customVirtualNetworkId"`
	EnableNoPublicIp                *WorkspaceCustomBooleanParameterResponse `pulumi:"enableNoPublicIp"`
	Encryption                      *WorkspaceEncryptionParameterResponse    `pulumi:"encryption"`
	LoadBalancerBackendPoolName     *WorkspaceCustomStringParameterResponse  `pulumi:"loadBalancerBackendPoolName"`
	LoadBalancerId                  *WorkspaceCustomStringParameterResponse  `pulumi:"loadBalancerId"`
	NatGatewayName                  *WorkspaceCustomStringParameterResponse  `pulumi:"natGatewayName"`
	PrepareEncryption               *WorkspaceCustomBooleanParameterResponse `pulumi:"prepareEncryption"`
	PublicIpName                    *WorkspaceCustomStringParameterResponse  `pulumi:"publicIpName"`
	RequireInfrastructureEncryption *WorkspaceCustomBooleanParameterResponse `pulumi:"requireInfrastructureEncryption"`
	ResourceTags                    WorkspaceCustomObjectParameterResponse   `pulumi:"resourceTags"`
	StorageAccountName              *WorkspaceCustomStringParameterResponse  `pulumi:"storageAccountName"`
	StorageAccountSkuName           *WorkspaceCustomStringParameterResponse  `pulumi:"storageAccountSkuName"`
	VnetAddressPrefix               *WorkspaceCustomStringParameterResponse  `pulumi:"vnetAddressPrefix"`
}


func (val *WorkspaceCustomParametersResponse) Defaults() *WorkspaceCustomParametersResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = tmp.Encryption.Defaults()

	return &tmp
}

type WorkspaceCustomParametersResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomParametersResponse)(nil)).Elem()
}

func (o WorkspaceCustomParametersResponseOutput) ToWorkspaceCustomParametersResponseOutput() WorkspaceCustomParametersResponseOutput {
	return o
}

func (o WorkspaceCustomParametersResponseOutput) ToWorkspaceCustomParametersResponseOutputWithContext(ctx context.Context) WorkspaceCustomParametersResponseOutput {
	return o
}

func (o WorkspaceCustomParametersResponseOutput) AmlWorkspaceId() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.AmlWorkspaceId
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) CustomPrivateSubnetName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.CustomPrivateSubnetName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) CustomPublicSubnetName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.CustomPublicSubnetName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) CustomVirtualNetworkId() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.CustomVirtualNetworkId
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) EnableNoPublicIp() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomBooleanParameterResponse {
		return v.EnableNoPublicIp
	}).(WorkspaceCustomBooleanParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) Encryption() WorkspaceEncryptionParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceEncryptionParameterResponse { return v.Encryption }).(WorkspaceEncryptionParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) LoadBalancerBackendPoolName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.LoadBalancerBackendPoolName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) LoadBalancerId() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.LoadBalancerId
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) NatGatewayName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.NatGatewayName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) PrepareEncryption() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomBooleanParameterResponse {
		return v.PrepareEncryption
	}).(WorkspaceCustomBooleanParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) PublicIpName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.PublicIpName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) RequireInfrastructureEncryption() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomBooleanParameterResponse {
		return v.RequireInfrastructureEncryption
	}).(WorkspaceCustomBooleanParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) ResourceTags() WorkspaceCustomObjectParameterResponseOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) WorkspaceCustomObjectParameterResponse {
		return v.ResourceTags
	}).(WorkspaceCustomObjectParameterResponseOutput)
}

func (o WorkspaceCustomParametersResponseOutput) StorageAccountName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.StorageAccountName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) StorageAccountSkuName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.StorageAccountSkuName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponseOutput) VnetAddressPrefix() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		return v.VnetAddressPrefix
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

type WorkspaceCustomParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomParametersResponse)(nil)).Elem()
}

func (o WorkspaceCustomParametersResponsePtrOutput) ToWorkspaceCustomParametersResponsePtrOutput() WorkspaceCustomParametersResponsePtrOutput {
	return o
}

func (o WorkspaceCustomParametersResponsePtrOutput) ToWorkspaceCustomParametersResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomParametersResponsePtrOutput {
	return o
}

func (o WorkspaceCustomParametersResponsePtrOutput) Elem() WorkspaceCustomParametersResponseOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) WorkspaceCustomParametersResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceCustomParametersResponse
		return ret
	}).(WorkspaceCustomParametersResponseOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) AmlWorkspaceId() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.AmlWorkspaceId
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) CustomPrivateSubnetName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.CustomPrivateSubnetName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) CustomPublicSubnetName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.CustomPublicSubnetName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) CustomVirtualNetworkId() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.CustomVirtualNetworkId
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) EnableNoPublicIp() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomBooleanParameterResponse {
		if v == nil {
			return nil
		}
		return v.EnableNoPublicIp
	}).(WorkspaceCustomBooleanParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) Encryption() WorkspaceEncryptionParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceEncryptionParameterResponse {
		if v == nil {
			return nil
		}
		return v.Encryption
	}).(WorkspaceEncryptionParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) LoadBalancerBackendPoolName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.LoadBalancerBackendPoolName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) LoadBalancerId() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.LoadBalancerId
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) NatGatewayName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.NatGatewayName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) PrepareEncryption() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomBooleanParameterResponse {
		if v == nil {
			return nil
		}
		return v.PrepareEncryption
	}).(WorkspaceCustomBooleanParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) PublicIpName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.PublicIpName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) RequireInfrastructureEncryption() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomBooleanParameterResponse {
		if v == nil {
			return nil
		}
		return v.RequireInfrastructureEncryption
	}).(WorkspaceCustomBooleanParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) ResourceTags() WorkspaceCustomObjectParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomObjectParameterResponse {
		if v == nil {
			return nil
		}
		return &v.ResourceTags
	}).(WorkspaceCustomObjectParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) StorageAccountName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) StorageAccountSkuName() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.StorageAccountSkuName
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

func (o WorkspaceCustomParametersResponsePtrOutput) VnetAddressPrefix() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomParametersResponse) *WorkspaceCustomStringParameterResponse {
		if v == nil {
			return nil
		}
		return v.VnetAddressPrefix
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
}

type WorkspaceCustomStringParameter struct {
	Value string `pulumi:"value"`
}





type WorkspaceCustomStringParameterInput interface {
	pulumi.Input

	ToWorkspaceCustomStringParameterOutput() WorkspaceCustomStringParameterOutput
	ToWorkspaceCustomStringParameterOutputWithContext(context.Context) WorkspaceCustomStringParameterOutput
}

type WorkspaceCustomStringParameterArgs struct {
	Value pulumi.StringInput `pulumi:"value"`
}

func (WorkspaceCustomStringParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomStringParameter)(nil)).Elem()
}

func (i WorkspaceCustomStringParameterArgs) ToWorkspaceCustomStringParameterOutput() WorkspaceCustomStringParameterOutput {
	return i.ToWorkspaceCustomStringParameterOutputWithContext(context.Background())
}

func (i WorkspaceCustomStringParameterArgs) ToWorkspaceCustomStringParameterOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomStringParameterOutput)
}

func (i WorkspaceCustomStringParameterArgs) ToWorkspaceCustomStringParameterPtrOutput() WorkspaceCustomStringParameterPtrOutput {
	return i.ToWorkspaceCustomStringParameterPtrOutputWithContext(context.Background())
}

func (i WorkspaceCustomStringParameterArgs) ToWorkspaceCustomStringParameterPtrOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomStringParameterOutput).ToWorkspaceCustomStringParameterPtrOutputWithContext(ctx)
}









type WorkspaceCustomStringParameterPtrInput interface {
	pulumi.Input

	ToWorkspaceCustomStringParameterPtrOutput() WorkspaceCustomStringParameterPtrOutput
	ToWorkspaceCustomStringParameterPtrOutputWithContext(context.Context) WorkspaceCustomStringParameterPtrOutput
}

type workspaceCustomStringParameterPtrType WorkspaceCustomStringParameterArgs

func WorkspaceCustomStringParameterPtr(v *WorkspaceCustomStringParameterArgs) WorkspaceCustomStringParameterPtrInput {
	return (*workspaceCustomStringParameterPtrType)(v)
}

func (*workspaceCustomStringParameterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomStringParameter)(nil)).Elem()
}

func (i *workspaceCustomStringParameterPtrType) ToWorkspaceCustomStringParameterPtrOutput() WorkspaceCustomStringParameterPtrOutput {
	return i.ToWorkspaceCustomStringParameterPtrOutputWithContext(context.Background())
}

func (i *workspaceCustomStringParameterPtrType) ToWorkspaceCustomStringParameterPtrOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomStringParameterPtrOutput)
}

type WorkspaceCustomStringParameterOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomStringParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomStringParameter)(nil)).Elem()
}

func (o WorkspaceCustomStringParameterOutput) ToWorkspaceCustomStringParameterOutput() WorkspaceCustomStringParameterOutput {
	return o
}

func (o WorkspaceCustomStringParameterOutput) ToWorkspaceCustomStringParameterOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterOutput {
	return o
}

func (o WorkspaceCustomStringParameterOutput) ToWorkspaceCustomStringParameterPtrOutput() WorkspaceCustomStringParameterPtrOutput {
	return o.ToWorkspaceCustomStringParameterPtrOutputWithContext(context.Background())
}

func (o WorkspaceCustomStringParameterOutput) ToWorkspaceCustomStringParameterPtrOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCustomStringParameter) *WorkspaceCustomStringParameter {
		return &v
	}).(WorkspaceCustomStringParameterPtrOutput)
}

func (o WorkspaceCustomStringParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCustomStringParameter) string { return v.Value }).(pulumi.StringOutput)
}

type WorkspaceCustomStringParameterPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomStringParameterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomStringParameter)(nil)).Elem()
}

func (o WorkspaceCustomStringParameterPtrOutput) ToWorkspaceCustomStringParameterPtrOutput() WorkspaceCustomStringParameterPtrOutput {
	return o
}

func (o WorkspaceCustomStringParameterPtrOutput) ToWorkspaceCustomStringParameterPtrOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterPtrOutput {
	return o
}

func (o WorkspaceCustomStringParameterPtrOutput) Elem() WorkspaceCustomStringParameterOutput {
	return o.ApplyT(func(v *WorkspaceCustomStringParameter) WorkspaceCustomStringParameter {
		if v != nil {
			return *v
		}
		var ret WorkspaceCustomStringParameter
		return ret
	}).(WorkspaceCustomStringParameterOutput)
}

func (o WorkspaceCustomStringParameterPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomStringParameter) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type WorkspaceCustomStringParameterResponse struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type WorkspaceCustomStringParameterResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomStringParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomStringParameterResponse)(nil)).Elem()
}

func (o WorkspaceCustomStringParameterResponseOutput) ToWorkspaceCustomStringParameterResponseOutput() WorkspaceCustomStringParameterResponseOutput {
	return o
}

func (o WorkspaceCustomStringParameterResponseOutput) ToWorkspaceCustomStringParameterResponseOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterResponseOutput {
	return o
}

func (o WorkspaceCustomStringParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCustomStringParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o WorkspaceCustomStringParameterResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCustomStringParameterResponse) string { return v.Value }).(pulumi.StringOutput)
}

type WorkspaceCustomStringParameterResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCustomStringParameterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomStringParameterResponse)(nil)).Elem()
}

func (o WorkspaceCustomStringParameterResponsePtrOutput) ToWorkspaceCustomStringParameterResponsePtrOutput() WorkspaceCustomStringParameterResponsePtrOutput {
	return o
}

func (o WorkspaceCustomStringParameterResponsePtrOutput) ToWorkspaceCustomStringParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterResponsePtrOutput {
	return o
}

func (o WorkspaceCustomStringParameterResponsePtrOutput) Elem() WorkspaceCustomStringParameterResponseOutput {
	return o.ApplyT(func(v *WorkspaceCustomStringParameterResponse) WorkspaceCustomStringParameterResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceCustomStringParameterResponse
		return ret
	}).(WorkspaceCustomStringParameterResponseOutput)
}

func (o WorkspaceCustomStringParameterResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomStringParameterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceCustomStringParameterResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCustomStringParameterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type WorkspaceEncryptionParameter struct {
	Value *Encryption `pulumi:"value"`
}


func (val *WorkspaceEncryptionParameter) Defaults() *WorkspaceEncryptionParameter {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Value = tmp.Value.Defaults()

	return &tmp
}





type WorkspaceEncryptionParameterInput interface {
	pulumi.Input

	ToWorkspaceEncryptionParameterOutput() WorkspaceEncryptionParameterOutput
	ToWorkspaceEncryptionParameterOutputWithContext(context.Context) WorkspaceEncryptionParameterOutput
}

type WorkspaceEncryptionParameterArgs struct {
	Value EncryptionPtrInput `pulumi:"value"`
}


func (val *WorkspaceEncryptionParameterArgs) Defaults() *WorkspaceEncryptionParameterArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (WorkspaceEncryptionParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceEncryptionParameter)(nil)).Elem()
}

func (i WorkspaceEncryptionParameterArgs) ToWorkspaceEncryptionParameterOutput() WorkspaceEncryptionParameterOutput {
	return i.ToWorkspaceEncryptionParameterOutputWithContext(context.Background())
}

func (i WorkspaceEncryptionParameterArgs) ToWorkspaceEncryptionParameterOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceEncryptionParameterOutput)
}

func (i WorkspaceEncryptionParameterArgs) ToWorkspaceEncryptionParameterPtrOutput() WorkspaceEncryptionParameterPtrOutput {
	return i.ToWorkspaceEncryptionParameterPtrOutputWithContext(context.Background())
}

func (i WorkspaceEncryptionParameterArgs) ToWorkspaceEncryptionParameterPtrOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceEncryptionParameterOutput).ToWorkspaceEncryptionParameterPtrOutputWithContext(ctx)
}









type WorkspaceEncryptionParameterPtrInput interface {
	pulumi.Input

	ToWorkspaceEncryptionParameterPtrOutput() WorkspaceEncryptionParameterPtrOutput
	ToWorkspaceEncryptionParameterPtrOutputWithContext(context.Context) WorkspaceEncryptionParameterPtrOutput
}

type workspaceEncryptionParameterPtrType WorkspaceEncryptionParameterArgs

func WorkspaceEncryptionParameterPtr(v *WorkspaceEncryptionParameterArgs) WorkspaceEncryptionParameterPtrInput {
	return (*workspaceEncryptionParameterPtrType)(v)
}

func (*workspaceEncryptionParameterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceEncryptionParameter)(nil)).Elem()
}

func (i *workspaceEncryptionParameterPtrType) ToWorkspaceEncryptionParameterPtrOutput() WorkspaceEncryptionParameterPtrOutput {
	return i.ToWorkspaceEncryptionParameterPtrOutputWithContext(context.Background())
}

func (i *workspaceEncryptionParameterPtrType) ToWorkspaceEncryptionParameterPtrOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceEncryptionParameterPtrOutput)
}

type WorkspaceEncryptionParameterOutput struct{ *pulumi.OutputState }

func (WorkspaceEncryptionParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceEncryptionParameter)(nil)).Elem()
}

func (o WorkspaceEncryptionParameterOutput) ToWorkspaceEncryptionParameterOutput() WorkspaceEncryptionParameterOutput {
	return o
}

func (o WorkspaceEncryptionParameterOutput) ToWorkspaceEncryptionParameterOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterOutput {
	return o
}

func (o WorkspaceEncryptionParameterOutput) ToWorkspaceEncryptionParameterPtrOutput() WorkspaceEncryptionParameterPtrOutput {
	return o.ToWorkspaceEncryptionParameterPtrOutputWithContext(context.Background())
}

func (o WorkspaceEncryptionParameterOutput) ToWorkspaceEncryptionParameterPtrOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceEncryptionParameter) *WorkspaceEncryptionParameter {
		return &v
	}).(WorkspaceEncryptionParameterPtrOutput)
}

func (o WorkspaceEncryptionParameterOutput) Value() EncryptionPtrOutput {
	return o.ApplyT(func(v WorkspaceEncryptionParameter) *Encryption { return v.Value }).(EncryptionPtrOutput)
}

type WorkspaceEncryptionParameterPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceEncryptionParameterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceEncryptionParameter)(nil)).Elem()
}

func (o WorkspaceEncryptionParameterPtrOutput) ToWorkspaceEncryptionParameterPtrOutput() WorkspaceEncryptionParameterPtrOutput {
	return o
}

func (o WorkspaceEncryptionParameterPtrOutput) ToWorkspaceEncryptionParameterPtrOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterPtrOutput {
	return o
}

func (o WorkspaceEncryptionParameterPtrOutput) Elem() WorkspaceEncryptionParameterOutput {
	return o.ApplyT(func(v *WorkspaceEncryptionParameter) WorkspaceEncryptionParameter {
		if v != nil {
			return *v
		}
		var ret WorkspaceEncryptionParameter
		return ret
	}).(WorkspaceEncryptionParameterOutput)
}

func (o WorkspaceEncryptionParameterPtrOutput) Value() EncryptionPtrOutput {
	return o.ApplyT(func(v *WorkspaceEncryptionParameter) *Encryption {
		if v == nil {
			return nil
		}
		return v.Value
	}).(EncryptionPtrOutput)
}

type WorkspaceEncryptionParameterResponse struct {
	Type  string              `pulumi:"type"`
	Value *EncryptionResponse `pulumi:"value"`
}


func (val *WorkspaceEncryptionParameterResponse) Defaults() *WorkspaceEncryptionParameterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Value = tmp.Value.Defaults()

	return &tmp
}

type WorkspaceEncryptionParameterResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceEncryptionParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceEncryptionParameterResponse)(nil)).Elem()
}

func (o WorkspaceEncryptionParameterResponseOutput) ToWorkspaceEncryptionParameterResponseOutput() WorkspaceEncryptionParameterResponseOutput {
	return o
}

func (o WorkspaceEncryptionParameterResponseOutput) ToWorkspaceEncryptionParameterResponseOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterResponseOutput {
	return o
}

func (o WorkspaceEncryptionParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceEncryptionParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o WorkspaceEncryptionParameterResponseOutput) Value() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v WorkspaceEncryptionParameterResponse) *EncryptionResponse { return v.Value }).(EncryptionResponsePtrOutput)
}

type WorkspaceEncryptionParameterResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceEncryptionParameterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceEncryptionParameterResponse)(nil)).Elem()
}

func (o WorkspaceEncryptionParameterResponsePtrOutput) ToWorkspaceEncryptionParameterResponsePtrOutput() WorkspaceEncryptionParameterResponsePtrOutput {
	return o
}

func (o WorkspaceEncryptionParameterResponsePtrOutput) ToWorkspaceEncryptionParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterResponsePtrOutput {
	return o
}

func (o WorkspaceEncryptionParameterResponsePtrOutput) Elem() WorkspaceEncryptionParameterResponseOutput {
	return o.ApplyT(func(v *WorkspaceEncryptionParameterResponse) WorkspaceEncryptionParameterResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceEncryptionParameterResponse
		return ret
	}).(WorkspaceEncryptionParameterResponseOutput)
}

func (o WorkspaceEncryptionParameterResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceEncryptionParameterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceEncryptionParameterResponsePtrOutput) Value() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *WorkspaceEncryptionParameterResponse) *EncryptionResponse {
		if v == nil {
			return nil
		}
		return v.Value
	}).(EncryptionResponsePtrOutput)
}

type WorkspaceProviderAuthorization struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type WorkspaceProviderAuthorizationInput interface {
	pulumi.Input

	ToWorkspaceProviderAuthorizationOutput() WorkspaceProviderAuthorizationOutput
	ToWorkspaceProviderAuthorizationOutputWithContext(context.Context) WorkspaceProviderAuthorizationOutput
}

type WorkspaceProviderAuthorizationArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (WorkspaceProviderAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceProviderAuthorization)(nil)).Elem()
}

func (i WorkspaceProviderAuthorizationArgs) ToWorkspaceProviderAuthorizationOutput() WorkspaceProviderAuthorizationOutput {
	return i.ToWorkspaceProviderAuthorizationOutputWithContext(context.Background())
}

func (i WorkspaceProviderAuthorizationArgs) ToWorkspaceProviderAuthorizationOutputWithContext(ctx context.Context) WorkspaceProviderAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceProviderAuthorizationOutput)
}





type WorkspaceProviderAuthorizationArrayInput interface {
	pulumi.Input

	ToWorkspaceProviderAuthorizationArrayOutput() WorkspaceProviderAuthorizationArrayOutput
	ToWorkspaceProviderAuthorizationArrayOutputWithContext(context.Context) WorkspaceProviderAuthorizationArrayOutput
}

type WorkspaceProviderAuthorizationArray []WorkspaceProviderAuthorizationInput

func (WorkspaceProviderAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkspaceProviderAuthorization)(nil)).Elem()
}

func (i WorkspaceProviderAuthorizationArray) ToWorkspaceProviderAuthorizationArrayOutput() WorkspaceProviderAuthorizationArrayOutput {
	return i.ToWorkspaceProviderAuthorizationArrayOutputWithContext(context.Background())
}

func (i WorkspaceProviderAuthorizationArray) ToWorkspaceProviderAuthorizationArrayOutputWithContext(ctx context.Context) WorkspaceProviderAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceProviderAuthorizationArrayOutput)
}

type WorkspaceProviderAuthorizationOutput struct{ *pulumi.OutputState }

func (WorkspaceProviderAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceProviderAuthorization)(nil)).Elem()
}

func (o WorkspaceProviderAuthorizationOutput) ToWorkspaceProviderAuthorizationOutput() WorkspaceProviderAuthorizationOutput {
	return o
}

func (o WorkspaceProviderAuthorizationOutput) ToWorkspaceProviderAuthorizationOutputWithContext(ctx context.Context) WorkspaceProviderAuthorizationOutput {
	return o
}

func (o WorkspaceProviderAuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceProviderAuthorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o WorkspaceProviderAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceProviderAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type WorkspaceProviderAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceProviderAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkspaceProviderAuthorization)(nil)).Elem()
}

func (o WorkspaceProviderAuthorizationArrayOutput) ToWorkspaceProviderAuthorizationArrayOutput() WorkspaceProviderAuthorizationArrayOutput {
	return o
}

func (o WorkspaceProviderAuthorizationArrayOutput) ToWorkspaceProviderAuthorizationArrayOutputWithContext(ctx context.Context) WorkspaceProviderAuthorizationArrayOutput {
	return o
}

func (o WorkspaceProviderAuthorizationArrayOutput) Index(i pulumi.IntInput) WorkspaceProviderAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkspaceProviderAuthorization {
		return vs[0].([]WorkspaceProviderAuthorization)[vs[1].(int)]
	}).(WorkspaceProviderAuthorizationOutput)
}

type WorkspaceProviderAuthorizationResponse struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}

type WorkspaceProviderAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceProviderAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceProviderAuthorizationResponse)(nil)).Elem()
}

func (o WorkspaceProviderAuthorizationResponseOutput) ToWorkspaceProviderAuthorizationResponseOutput() WorkspaceProviderAuthorizationResponseOutput {
	return o
}

func (o WorkspaceProviderAuthorizationResponseOutput) ToWorkspaceProviderAuthorizationResponseOutputWithContext(ctx context.Context) WorkspaceProviderAuthorizationResponseOutput {
	return o
}

func (o WorkspaceProviderAuthorizationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceProviderAuthorizationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o WorkspaceProviderAuthorizationResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceProviderAuthorizationResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type WorkspaceProviderAuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceProviderAuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkspaceProviderAuthorizationResponse)(nil)).Elem()
}

func (o WorkspaceProviderAuthorizationResponseArrayOutput) ToWorkspaceProviderAuthorizationResponseArrayOutput() WorkspaceProviderAuthorizationResponseArrayOutput {
	return o
}

func (o WorkspaceProviderAuthorizationResponseArrayOutput) ToWorkspaceProviderAuthorizationResponseArrayOutputWithContext(ctx context.Context) WorkspaceProviderAuthorizationResponseArrayOutput {
	return o
}

func (o WorkspaceProviderAuthorizationResponseArrayOutput) Index(i pulumi.IntInput) WorkspaceProviderAuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkspaceProviderAuthorizationResponse {
		return vs[0].([]WorkspaceProviderAuthorizationResponse)[vs[1].(int)]
	}).(WorkspaceProviderAuthorizationResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressSpaceOutput{})
	pulumi.RegisterOutputType(AddressSpacePtrOutput{})
	pulumi.RegisterOutputType(AddressSpaceResponseOutput{})
	pulumi.RegisterOutputType(AddressSpaceResponsePtrOutput{})
	pulumi.RegisterOutputType(CreatedByResponseOutput{})
	pulumi.RegisterOutputType(CreatedByResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomBooleanParameterOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomBooleanParameterPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomBooleanParameterResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomBooleanParameterResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomObjectParameterResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomObjectParameterResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomParametersOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomParametersPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomParametersResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomStringParameterOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomStringParameterPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomStringParameterResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceCustomStringParameterResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceEncryptionParameterOutput{})
	pulumi.RegisterOutputType(WorkspaceEncryptionParameterPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceEncryptionParameterResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceEncryptionParameterResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceProviderAuthorizationOutput{})
	pulumi.RegisterOutputType(WorkspaceProviderAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceProviderAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceProviderAuthorizationResponseArrayOutput{})
}
