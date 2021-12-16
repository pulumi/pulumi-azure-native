


package databricks

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





type AddressSpaceResponseInput interface {
	pulumi.Input

	ToAddressSpaceResponseOutput() AddressSpaceResponseOutput
	ToAddressSpaceResponseOutputWithContext(context.Context) AddressSpaceResponseOutput
}

type AddressSpaceResponseArgs struct {
	AddressPrefixes pulumi.StringArrayInput `pulumi:"addressPrefixes"`
}

func (AddressSpaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpaceResponse)(nil)).Elem()
}

func (i AddressSpaceResponseArgs) ToAddressSpaceResponseOutput() AddressSpaceResponseOutput {
	return i.ToAddressSpaceResponseOutputWithContext(context.Background())
}

func (i AddressSpaceResponseArgs) ToAddressSpaceResponseOutputWithContext(ctx context.Context) AddressSpaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceResponseOutput)
}

func (i AddressSpaceResponseArgs) ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput {
	return i.ToAddressSpaceResponsePtrOutputWithContext(context.Background())
}

func (i AddressSpaceResponseArgs) ToAddressSpaceResponsePtrOutputWithContext(ctx context.Context) AddressSpaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceResponseOutput).ToAddressSpaceResponsePtrOutputWithContext(ctx)
}









type AddressSpaceResponsePtrInput interface {
	pulumi.Input

	ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput
	ToAddressSpaceResponsePtrOutputWithContext(context.Context) AddressSpaceResponsePtrOutput
}

type addressSpaceResponsePtrType AddressSpaceResponseArgs

func AddressSpaceResponsePtr(v *AddressSpaceResponseArgs) AddressSpaceResponsePtrInput {
	return (*addressSpaceResponsePtrType)(v)
}

func (*addressSpaceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpaceResponse)(nil)).Elem()
}

func (i *addressSpaceResponsePtrType) ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput {
	return i.ToAddressSpaceResponsePtrOutputWithContext(context.Background())
}

func (i *addressSpaceResponsePtrType) ToAddressSpaceResponsePtrOutputWithContext(ctx context.Context) AddressSpaceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceResponsePtrOutput)
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

func (o AddressSpaceResponseOutput) ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput {
	return o.ToAddressSpaceResponsePtrOutputWithContext(context.Background())
}

func (o AddressSpaceResponseOutput) ToAddressSpaceResponsePtrOutputWithContext(ctx context.Context) AddressSpaceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressSpaceResponse) *AddressSpaceResponse {
		return &v
	}).(AddressSpaceResponsePtrOutput)
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





type CreatedByResponseInput interface {
	pulumi.Input

	ToCreatedByResponseOutput() CreatedByResponseOutput
	ToCreatedByResponseOutputWithContext(context.Context) CreatedByResponseOutput
}

type CreatedByResponseArgs struct {
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	Oid           pulumi.StringInput `pulumi:"oid"`
	Puid          pulumi.StringInput `pulumi:"puid"`
}

func (CreatedByResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatedByResponse)(nil)).Elem()
}

func (i CreatedByResponseArgs) ToCreatedByResponseOutput() CreatedByResponseOutput {
	return i.ToCreatedByResponseOutputWithContext(context.Background())
}

func (i CreatedByResponseArgs) ToCreatedByResponseOutputWithContext(ctx context.Context) CreatedByResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatedByResponseOutput)
}

func (i CreatedByResponseArgs) ToCreatedByResponsePtrOutput() CreatedByResponsePtrOutput {
	return i.ToCreatedByResponsePtrOutputWithContext(context.Background())
}

func (i CreatedByResponseArgs) ToCreatedByResponsePtrOutputWithContext(ctx context.Context) CreatedByResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatedByResponseOutput).ToCreatedByResponsePtrOutputWithContext(ctx)
}









type CreatedByResponsePtrInput interface {
	pulumi.Input

	ToCreatedByResponsePtrOutput() CreatedByResponsePtrOutput
	ToCreatedByResponsePtrOutputWithContext(context.Context) CreatedByResponsePtrOutput
}

type createdByResponsePtrType CreatedByResponseArgs

func CreatedByResponsePtr(v *CreatedByResponseArgs) CreatedByResponsePtrInput {
	return (*createdByResponsePtrType)(v)
}

func (*createdByResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreatedByResponse)(nil)).Elem()
}

func (i *createdByResponsePtrType) ToCreatedByResponsePtrOutput() CreatedByResponsePtrOutput {
	return i.ToCreatedByResponsePtrOutputWithContext(context.Background())
}

func (i *createdByResponsePtrType) ToCreatedByResponsePtrOutputWithContext(ctx context.Context) CreatedByResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatedByResponsePtrOutput)
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

func (o CreatedByResponseOutput) ToCreatedByResponsePtrOutput() CreatedByResponsePtrOutput {
	return o.ToCreatedByResponsePtrOutputWithContext(context.Background())
}

func (o CreatedByResponseOutput) ToCreatedByResponsePtrOutputWithContext(ctx context.Context) CreatedByResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreatedByResponse) *CreatedByResponse {
		return &v
	}).(CreatedByResponsePtrOutput)
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





type EncryptionResponseInput interface {
	pulumi.Input

	ToEncryptionResponseOutput() EncryptionResponseOutput
	ToEncryptionResponseOutputWithContext(context.Context) EncryptionResponseOutput
}

type EncryptionResponseArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeySource   pulumi.StringPtrInput `pulumi:"keySource"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (EncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return i.ToEncryptionResponseOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput)
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput).ToEncryptionResponsePtrOutputWithContext(ctx)
}









type EncryptionResponsePtrInput interface {
	pulumi.Input

	ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput
	ToEncryptionResponsePtrOutputWithContext(context.Context) EncryptionResponsePtrOutput
}

type encryptionResponsePtrType EncryptionResponseArgs

func EncryptionResponsePtr(v *EncryptionResponseArgs) EncryptionResponsePtrInput {
	return (*encryptionResponsePtrType)(v)
}

func (*encryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponsePtrOutput)
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

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionResponse) *EncryptionResponse {
		return &v
	}).(EncryptionResponsePtrOutput)
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





type ManagedIdentityConfigurationResponseInput interface {
	pulumi.Input

	ToManagedIdentityConfigurationResponseOutput() ManagedIdentityConfigurationResponseOutput
	ToManagedIdentityConfigurationResponseOutputWithContext(context.Context) ManagedIdentityConfigurationResponseOutput
}

type ManagedIdentityConfigurationResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (ManagedIdentityConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityConfigurationResponse)(nil)).Elem()
}

func (i ManagedIdentityConfigurationResponseArgs) ToManagedIdentityConfigurationResponseOutput() ManagedIdentityConfigurationResponseOutput {
	return i.ToManagedIdentityConfigurationResponseOutputWithContext(context.Background())
}

func (i ManagedIdentityConfigurationResponseArgs) ToManagedIdentityConfigurationResponseOutputWithContext(ctx context.Context) ManagedIdentityConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityConfigurationResponseOutput)
}

func (i ManagedIdentityConfigurationResponseArgs) ToManagedIdentityConfigurationResponsePtrOutput() ManagedIdentityConfigurationResponsePtrOutput {
	return i.ToManagedIdentityConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i ManagedIdentityConfigurationResponseArgs) ToManagedIdentityConfigurationResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityConfigurationResponseOutput).ToManagedIdentityConfigurationResponsePtrOutputWithContext(ctx)
}









type ManagedIdentityConfigurationResponsePtrInput interface {
	pulumi.Input

	ToManagedIdentityConfigurationResponsePtrOutput() ManagedIdentityConfigurationResponsePtrOutput
	ToManagedIdentityConfigurationResponsePtrOutputWithContext(context.Context) ManagedIdentityConfigurationResponsePtrOutput
}

type managedIdentityConfigurationResponsePtrType ManagedIdentityConfigurationResponseArgs

func ManagedIdentityConfigurationResponsePtr(v *ManagedIdentityConfigurationResponseArgs) ManagedIdentityConfigurationResponsePtrInput {
	return (*managedIdentityConfigurationResponsePtrType)(v)
}

func (*managedIdentityConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityConfigurationResponse)(nil)).Elem()
}

func (i *managedIdentityConfigurationResponsePtrType) ToManagedIdentityConfigurationResponsePtrOutput() ManagedIdentityConfigurationResponsePtrOutput {
	return i.ToManagedIdentityConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *managedIdentityConfigurationResponsePtrType) ToManagedIdentityConfigurationResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityConfigurationResponsePtrOutput)
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

func (o ManagedIdentityConfigurationResponseOutput) ToManagedIdentityConfigurationResponsePtrOutput() ManagedIdentityConfigurationResponsePtrOutput {
	return o.ToManagedIdentityConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityConfigurationResponseOutput) ToManagedIdentityConfigurationResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityConfigurationResponse) *ManagedIdentityConfigurationResponse {
		return &v
	}).(ManagedIdentityConfigurationResponsePtrOutput)
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

type PrivateEndpointConnectionProperties struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
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

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
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

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput
	ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponseOutput
}

type PrivateEndpointConnectionPropertiesResponseArgs struct {
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
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

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
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

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStateResponse {
		if v == nil {
			return nil
		}
		return &v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
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

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
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
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         string  `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionRequired pulumi.StringPtrInput `pulumi:"actionRequired"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Status         pulumi.StringInput    `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
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

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
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
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         string  `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionRequired pulumi.StringPtrInput `pulumi:"actionRequired"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Status         pulumi.StringInput    `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
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





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
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

func (i VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput).ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutputWithContext(ctx)
}









type VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput
	ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutputWithContext(context.Context) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput
}

type virtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrType VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkArgs

func VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtr(v *VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkArgs) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrInput {
	return (*virtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrType)(v)
}

func (*virtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork)(nil)).Elem()
}

func (i *virtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrType) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrType) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput)
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

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput {
	return o.ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork) *VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork {
		return &v
	}).(VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput)
}

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput) ToVirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput) Elem() VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork) VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork
		return ret
	}).(VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput)
}

func (o VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork struct {
	Id *string `pulumi:"id"`
}





type VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput
	ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutputWithContext(context.Context) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput
}

type VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork)(nil)).Elem()
}

func (i VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput)
}

func (i VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput).ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutputWithContext(ctx)
}









type VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput
	ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutputWithContext(context.Context) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput
}

type virtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrType VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkArgs

func VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtr(v *VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkArgs) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrInput {
	return (*virtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrType)(v)
}

func (*virtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork)(nil)).Elem()
}

func (i *virtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrType) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrType) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput)
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

func (o VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput {
	return o.ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork) *VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork {
		return &v
	}).(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput)
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





type VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput
	ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutputWithContext(context.Context) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput
}

type VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork)(nil)).Elem()
}

func (i VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput)
}

func (i VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkArgs) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput).ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutputWithContext(ctx)
}









type VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput
	ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutputWithContext(context.Context) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput
}

type virtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrType VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkArgs

func VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtr(v *VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkArgs) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrInput {
	return (*virtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrType)(v)
}

func (*virtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork)(nil)).Elem()
}

func (i *virtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrType) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput {
	return i.ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrType) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput)
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

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput {
	return o.ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork) *VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork {
		return &v
	}).(VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput)
}

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput) ToVirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutputWithContext(ctx context.Context) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput {
	return o
}

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput) Elem() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork
		return ret
	}).(VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput)
}

func (o VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
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





type WorkspaceCustomBooleanParameterResponseInput interface {
	pulumi.Input

	ToWorkspaceCustomBooleanParameterResponseOutput() WorkspaceCustomBooleanParameterResponseOutput
	ToWorkspaceCustomBooleanParameterResponseOutputWithContext(context.Context) WorkspaceCustomBooleanParameterResponseOutput
}

type WorkspaceCustomBooleanParameterResponseArgs struct {
	Type  pulumi.StringInput `pulumi:"type"`
	Value pulumi.BoolInput   `pulumi:"value"`
}

func (WorkspaceCustomBooleanParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomBooleanParameterResponse)(nil)).Elem()
}

func (i WorkspaceCustomBooleanParameterResponseArgs) ToWorkspaceCustomBooleanParameterResponseOutput() WorkspaceCustomBooleanParameterResponseOutput {
	return i.ToWorkspaceCustomBooleanParameterResponseOutputWithContext(context.Background())
}

func (i WorkspaceCustomBooleanParameterResponseArgs) ToWorkspaceCustomBooleanParameterResponseOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomBooleanParameterResponseOutput)
}

func (i WorkspaceCustomBooleanParameterResponseArgs) ToWorkspaceCustomBooleanParameterResponsePtrOutput() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return i.ToWorkspaceCustomBooleanParameterResponsePtrOutputWithContext(context.Background())
}

func (i WorkspaceCustomBooleanParameterResponseArgs) ToWorkspaceCustomBooleanParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomBooleanParameterResponseOutput).ToWorkspaceCustomBooleanParameterResponsePtrOutputWithContext(ctx)
}









type WorkspaceCustomBooleanParameterResponsePtrInput interface {
	pulumi.Input

	ToWorkspaceCustomBooleanParameterResponsePtrOutput() WorkspaceCustomBooleanParameterResponsePtrOutput
	ToWorkspaceCustomBooleanParameterResponsePtrOutputWithContext(context.Context) WorkspaceCustomBooleanParameterResponsePtrOutput
}

type workspaceCustomBooleanParameterResponsePtrType WorkspaceCustomBooleanParameterResponseArgs

func WorkspaceCustomBooleanParameterResponsePtr(v *WorkspaceCustomBooleanParameterResponseArgs) WorkspaceCustomBooleanParameterResponsePtrInput {
	return (*workspaceCustomBooleanParameterResponsePtrType)(v)
}

func (*workspaceCustomBooleanParameterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomBooleanParameterResponse)(nil)).Elem()
}

func (i *workspaceCustomBooleanParameterResponsePtrType) ToWorkspaceCustomBooleanParameterResponsePtrOutput() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return i.ToWorkspaceCustomBooleanParameterResponsePtrOutputWithContext(context.Background())
}

func (i *workspaceCustomBooleanParameterResponsePtrType) ToWorkspaceCustomBooleanParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomBooleanParameterResponsePtrOutput)
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

func (o WorkspaceCustomBooleanParameterResponseOutput) ToWorkspaceCustomBooleanParameterResponsePtrOutput() WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o.ToWorkspaceCustomBooleanParameterResponsePtrOutputWithContext(context.Background())
}

func (o WorkspaceCustomBooleanParameterResponseOutput) ToWorkspaceCustomBooleanParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomBooleanParameterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCustomBooleanParameterResponse) *WorkspaceCustomBooleanParameterResponse {
		return &v
	}).(WorkspaceCustomBooleanParameterResponsePtrOutput)
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





type WorkspaceCustomObjectParameterResponseInput interface {
	pulumi.Input

	ToWorkspaceCustomObjectParameterResponseOutput() WorkspaceCustomObjectParameterResponseOutput
	ToWorkspaceCustomObjectParameterResponseOutputWithContext(context.Context) WorkspaceCustomObjectParameterResponseOutput
}

type WorkspaceCustomObjectParameterResponseArgs struct {
	Type  pulumi.StringInput `pulumi:"type"`
	Value pulumi.Input       `pulumi:"value"`
}

func (WorkspaceCustomObjectParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomObjectParameterResponse)(nil)).Elem()
}

func (i WorkspaceCustomObjectParameterResponseArgs) ToWorkspaceCustomObjectParameterResponseOutput() WorkspaceCustomObjectParameterResponseOutput {
	return i.ToWorkspaceCustomObjectParameterResponseOutputWithContext(context.Background())
}

func (i WorkspaceCustomObjectParameterResponseArgs) ToWorkspaceCustomObjectParameterResponseOutputWithContext(ctx context.Context) WorkspaceCustomObjectParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomObjectParameterResponseOutput)
}

func (i WorkspaceCustomObjectParameterResponseArgs) ToWorkspaceCustomObjectParameterResponsePtrOutput() WorkspaceCustomObjectParameterResponsePtrOutput {
	return i.ToWorkspaceCustomObjectParameterResponsePtrOutputWithContext(context.Background())
}

func (i WorkspaceCustomObjectParameterResponseArgs) ToWorkspaceCustomObjectParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomObjectParameterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomObjectParameterResponseOutput).ToWorkspaceCustomObjectParameterResponsePtrOutputWithContext(ctx)
}









type WorkspaceCustomObjectParameterResponsePtrInput interface {
	pulumi.Input

	ToWorkspaceCustomObjectParameterResponsePtrOutput() WorkspaceCustomObjectParameterResponsePtrOutput
	ToWorkspaceCustomObjectParameterResponsePtrOutputWithContext(context.Context) WorkspaceCustomObjectParameterResponsePtrOutput
}

type workspaceCustomObjectParameterResponsePtrType WorkspaceCustomObjectParameterResponseArgs

func WorkspaceCustomObjectParameterResponsePtr(v *WorkspaceCustomObjectParameterResponseArgs) WorkspaceCustomObjectParameterResponsePtrInput {
	return (*workspaceCustomObjectParameterResponsePtrType)(v)
}

func (*workspaceCustomObjectParameterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomObjectParameterResponse)(nil)).Elem()
}

func (i *workspaceCustomObjectParameterResponsePtrType) ToWorkspaceCustomObjectParameterResponsePtrOutput() WorkspaceCustomObjectParameterResponsePtrOutput {
	return i.ToWorkspaceCustomObjectParameterResponsePtrOutputWithContext(context.Background())
}

func (i *workspaceCustomObjectParameterResponsePtrType) ToWorkspaceCustomObjectParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomObjectParameterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomObjectParameterResponsePtrOutput)
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

func (o WorkspaceCustomObjectParameterResponseOutput) ToWorkspaceCustomObjectParameterResponsePtrOutput() WorkspaceCustomObjectParameterResponsePtrOutput {
	return o.ToWorkspaceCustomObjectParameterResponsePtrOutputWithContext(context.Background())
}

func (o WorkspaceCustomObjectParameterResponseOutput) ToWorkspaceCustomObjectParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomObjectParameterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCustomObjectParameterResponse) *WorkspaceCustomObjectParameterResponse {
		return &v
	}).(WorkspaceCustomObjectParameterResponsePtrOutput)
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





type WorkspaceCustomParametersResponseInput interface {
	pulumi.Input

	ToWorkspaceCustomParametersResponseOutput() WorkspaceCustomParametersResponseOutput
	ToWorkspaceCustomParametersResponseOutputWithContext(context.Context) WorkspaceCustomParametersResponseOutput
}

type WorkspaceCustomParametersResponseArgs struct {
	AmlWorkspaceId                  WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"amlWorkspaceId"`
	CustomPrivateSubnetName         WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"customPrivateSubnetName"`
	CustomPublicSubnetName          WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"customPublicSubnetName"`
	CustomVirtualNetworkId          WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"customVirtualNetworkId"`
	EnableNoPublicIp                WorkspaceCustomBooleanParameterResponsePtrInput `pulumi:"enableNoPublicIp"`
	Encryption                      WorkspaceEncryptionParameterResponsePtrInput    `pulumi:"encryption"`
	LoadBalancerBackendPoolName     WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"loadBalancerBackendPoolName"`
	LoadBalancerId                  WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"loadBalancerId"`
	NatGatewayName                  WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"natGatewayName"`
	PrepareEncryption               WorkspaceCustomBooleanParameterResponsePtrInput `pulumi:"prepareEncryption"`
	PublicIpName                    WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"publicIpName"`
	RequireInfrastructureEncryption WorkspaceCustomBooleanParameterResponsePtrInput `pulumi:"requireInfrastructureEncryption"`
	ResourceTags                    WorkspaceCustomObjectParameterResponseInput     `pulumi:"resourceTags"`
	StorageAccountName              WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"storageAccountName"`
	StorageAccountSkuName           WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"storageAccountSkuName"`
	VnetAddressPrefix               WorkspaceCustomStringParameterResponsePtrInput  `pulumi:"vnetAddressPrefix"`
}

func (WorkspaceCustomParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomParametersResponse)(nil)).Elem()
}

func (i WorkspaceCustomParametersResponseArgs) ToWorkspaceCustomParametersResponseOutput() WorkspaceCustomParametersResponseOutput {
	return i.ToWorkspaceCustomParametersResponseOutputWithContext(context.Background())
}

func (i WorkspaceCustomParametersResponseArgs) ToWorkspaceCustomParametersResponseOutputWithContext(ctx context.Context) WorkspaceCustomParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomParametersResponseOutput)
}

func (i WorkspaceCustomParametersResponseArgs) ToWorkspaceCustomParametersResponsePtrOutput() WorkspaceCustomParametersResponsePtrOutput {
	return i.ToWorkspaceCustomParametersResponsePtrOutputWithContext(context.Background())
}

func (i WorkspaceCustomParametersResponseArgs) ToWorkspaceCustomParametersResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomParametersResponseOutput).ToWorkspaceCustomParametersResponsePtrOutputWithContext(ctx)
}









type WorkspaceCustomParametersResponsePtrInput interface {
	pulumi.Input

	ToWorkspaceCustomParametersResponsePtrOutput() WorkspaceCustomParametersResponsePtrOutput
	ToWorkspaceCustomParametersResponsePtrOutputWithContext(context.Context) WorkspaceCustomParametersResponsePtrOutput
}

type workspaceCustomParametersResponsePtrType WorkspaceCustomParametersResponseArgs

func WorkspaceCustomParametersResponsePtr(v *WorkspaceCustomParametersResponseArgs) WorkspaceCustomParametersResponsePtrInput {
	return (*workspaceCustomParametersResponsePtrType)(v)
}

func (*workspaceCustomParametersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomParametersResponse)(nil)).Elem()
}

func (i *workspaceCustomParametersResponsePtrType) ToWorkspaceCustomParametersResponsePtrOutput() WorkspaceCustomParametersResponsePtrOutput {
	return i.ToWorkspaceCustomParametersResponsePtrOutputWithContext(context.Background())
}

func (i *workspaceCustomParametersResponsePtrType) ToWorkspaceCustomParametersResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomParametersResponsePtrOutput)
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

func (o WorkspaceCustomParametersResponseOutput) ToWorkspaceCustomParametersResponsePtrOutput() WorkspaceCustomParametersResponsePtrOutput {
	return o.ToWorkspaceCustomParametersResponsePtrOutputWithContext(context.Background())
}

func (o WorkspaceCustomParametersResponseOutput) ToWorkspaceCustomParametersResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomParametersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCustomParametersResponse) *WorkspaceCustomParametersResponse {
		return &v
	}).(WorkspaceCustomParametersResponsePtrOutput)
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





type WorkspaceCustomStringParameterResponseInput interface {
	pulumi.Input

	ToWorkspaceCustomStringParameterResponseOutput() WorkspaceCustomStringParameterResponseOutput
	ToWorkspaceCustomStringParameterResponseOutputWithContext(context.Context) WorkspaceCustomStringParameterResponseOutput
}

type WorkspaceCustomStringParameterResponseArgs struct {
	Type  pulumi.StringInput `pulumi:"type"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (WorkspaceCustomStringParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCustomStringParameterResponse)(nil)).Elem()
}

func (i WorkspaceCustomStringParameterResponseArgs) ToWorkspaceCustomStringParameterResponseOutput() WorkspaceCustomStringParameterResponseOutput {
	return i.ToWorkspaceCustomStringParameterResponseOutputWithContext(context.Background())
}

func (i WorkspaceCustomStringParameterResponseArgs) ToWorkspaceCustomStringParameterResponseOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomStringParameterResponseOutput)
}

func (i WorkspaceCustomStringParameterResponseArgs) ToWorkspaceCustomStringParameterResponsePtrOutput() WorkspaceCustomStringParameterResponsePtrOutput {
	return i.ToWorkspaceCustomStringParameterResponsePtrOutputWithContext(context.Background())
}

func (i WorkspaceCustomStringParameterResponseArgs) ToWorkspaceCustomStringParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomStringParameterResponseOutput).ToWorkspaceCustomStringParameterResponsePtrOutputWithContext(ctx)
}









type WorkspaceCustomStringParameterResponsePtrInput interface {
	pulumi.Input

	ToWorkspaceCustomStringParameterResponsePtrOutput() WorkspaceCustomStringParameterResponsePtrOutput
	ToWorkspaceCustomStringParameterResponsePtrOutputWithContext(context.Context) WorkspaceCustomStringParameterResponsePtrOutput
}

type workspaceCustomStringParameterResponsePtrType WorkspaceCustomStringParameterResponseArgs

func WorkspaceCustomStringParameterResponsePtr(v *WorkspaceCustomStringParameterResponseArgs) WorkspaceCustomStringParameterResponsePtrInput {
	return (*workspaceCustomStringParameterResponsePtrType)(v)
}

func (*workspaceCustomStringParameterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCustomStringParameterResponse)(nil)).Elem()
}

func (i *workspaceCustomStringParameterResponsePtrType) ToWorkspaceCustomStringParameterResponsePtrOutput() WorkspaceCustomStringParameterResponsePtrOutput {
	return i.ToWorkspaceCustomStringParameterResponsePtrOutputWithContext(context.Background())
}

func (i *workspaceCustomStringParameterResponsePtrType) ToWorkspaceCustomStringParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCustomStringParameterResponsePtrOutput)
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

func (o WorkspaceCustomStringParameterResponseOutput) ToWorkspaceCustomStringParameterResponsePtrOutput() WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ToWorkspaceCustomStringParameterResponsePtrOutputWithContext(context.Background())
}

func (o WorkspaceCustomStringParameterResponseOutput) ToWorkspaceCustomStringParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceCustomStringParameterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCustomStringParameterResponse) *WorkspaceCustomStringParameterResponse {
		return &v
	}).(WorkspaceCustomStringParameterResponsePtrOutput)
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





type WorkspaceEncryptionParameterResponseInput interface {
	pulumi.Input

	ToWorkspaceEncryptionParameterResponseOutput() WorkspaceEncryptionParameterResponseOutput
	ToWorkspaceEncryptionParameterResponseOutputWithContext(context.Context) WorkspaceEncryptionParameterResponseOutput
}

type WorkspaceEncryptionParameterResponseArgs struct {
	Type  pulumi.StringInput         `pulumi:"type"`
	Value EncryptionResponsePtrInput `pulumi:"value"`
}

func (WorkspaceEncryptionParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceEncryptionParameterResponse)(nil)).Elem()
}

func (i WorkspaceEncryptionParameterResponseArgs) ToWorkspaceEncryptionParameterResponseOutput() WorkspaceEncryptionParameterResponseOutput {
	return i.ToWorkspaceEncryptionParameterResponseOutputWithContext(context.Background())
}

func (i WorkspaceEncryptionParameterResponseArgs) ToWorkspaceEncryptionParameterResponseOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceEncryptionParameterResponseOutput)
}

func (i WorkspaceEncryptionParameterResponseArgs) ToWorkspaceEncryptionParameterResponsePtrOutput() WorkspaceEncryptionParameterResponsePtrOutput {
	return i.ToWorkspaceEncryptionParameterResponsePtrOutputWithContext(context.Background())
}

func (i WorkspaceEncryptionParameterResponseArgs) ToWorkspaceEncryptionParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceEncryptionParameterResponseOutput).ToWorkspaceEncryptionParameterResponsePtrOutputWithContext(ctx)
}









type WorkspaceEncryptionParameterResponsePtrInput interface {
	pulumi.Input

	ToWorkspaceEncryptionParameterResponsePtrOutput() WorkspaceEncryptionParameterResponsePtrOutput
	ToWorkspaceEncryptionParameterResponsePtrOutputWithContext(context.Context) WorkspaceEncryptionParameterResponsePtrOutput
}

type workspaceEncryptionParameterResponsePtrType WorkspaceEncryptionParameterResponseArgs

func WorkspaceEncryptionParameterResponsePtr(v *WorkspaceEncryptionParameterResponseArgs) WorkspaceEncryptionParameterResponsePtrInput {
	return (*workspaceEncryptionParameterResponsePtrType)(v)
}

func (*workspaceEncryptionParameterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceEncryptionParameterResponse)(nil)).Elem()
}

func (i *workspaceEncryptionParameterResponsePtrType) ToWorkspaceEncryptionParameterResponsePtrOutput() WorkspaceEncryptionParameterResponsePtrOutput {
	return i.ToWorkspaceEncryptionParameterResponsePtrOutputWithContext(context.Background())
}

func (i *workspaceEncryptionParameterResponsePtrType) ToWorkspaceEncryptionParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceEncryptionParameterResponsePtrOutput)
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

func (o WorkspaceEncryptionParameterResponseOutput) ToWorkspaceEncryptionParameterResponsePtrOutput() WorkspaceEncryptionParameterResponsePtrOutput {
	return o.ToWorkspaceEncryptionParameterResponsePtrOutputWithContext(context.Background())
}

func (o WorkspaceEncryptionParameterResponseOutput) ToWorkspaceEncryptionParameterResponsePtrOutputWithContext(ctx context.Context) WorkspaceEncryptionParameterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceEncryptionParameterResponse) *WorkspaceEncryptionParameterResponse {
		return &v
	}).(WorkspaceEncryptionParameterResponsePtrOutput)
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





type WorkspaceProviderAuthorizationResponseInput interface {
	pulumi.Input

	ToWorkspaceProviderAuthorizationResponseOutput() WorkspaceProviderAuthorizationResponseOutput
	ToWorkspaceProviderAuthorizationResponseOutputWithContext(context.Context) WorkspaceProviderAuthorizationResponseOutput
}

type WorkspaceProviderAuthorizationResponseArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (WorkspaceProviderAuthorizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceProviderAuthorizationResponse)(nil)).Elem()
}

func (i WorkspaceProviderAuthorizationResponseArgs) ToWorkspaceProviderAuthorizationResponseOutput() WorkspaceProviderAuthorizationResponseOutput {
	return i.ToWorkspaceProviderAuthorizationResponseOutputWithContext(context.Background())
}

func (i WorkspaceProviderAuthorizationResponseArgs) ToWorkspaceProviderAuthorizationResponseOutputWithContext(ctx context.Context) WorkspaceProviderAuthorizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceProviderAuthorizationResponseOutput)
}





type WorkspaceProviderAuthorizationResponseArrayInput interface {
	pulumi.Input

	ToWorkspaceProviderAuthorizationResponseArrayOutput() WorkspaceProviderAuthorizationResponseArrayOutput
	ToWorkspaceProviderAuthorizationResponseArrayOutputWithContext(context.Context) WorkspaceProviderAuthorizationResponseArrayOutput
}

type WorkspaceProviderAuthorizationResponseArray []WorkspaceProviderAuthorizationResponseInput

func (WorkspaceProviderAuthorizationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkspaceProviderAuthorizationResponse)(nil)).Elem()
}

func (i WorkspaceProviderAuthorizationResponseArray) ToWorkspaceProviderAuthorizationResponseArrayOutput() WorkspaceProviderAuthorizationResponseArrayOutput {
	return i.ToWorkspaceProviderAuthorizationResponseArrayOutputWithContext(context.Background())
}

func (i WorkspaceProviderAuthorizationResponseArray) ToWorkspaceProviderAuthorizationResponseArrayOutputWithContext(ctx context.Context) WorkspaceProviderAuthorizationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceProviderAuthorizationResponseArrayOutput)
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
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkPtrOutput{})
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
