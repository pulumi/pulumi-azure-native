


package v20210401preview

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

type EncryptionEntitiesDefinition struct {
	ManagedServices *EncryptionV2 `pulumi:"managedServices"`
}





type EncryptionEntitiesDefinitionInput interface {
	pulumi.Input

	ToEncryptionEntitiesDefinitionOutput() EncryptionEntitiesDefinitionOutput
	ToEncryptionEntitiesDefinitionOutputWithContext(context.Context) EncryptionEntitiesDefinitionOutput
}

type EncryptionEntitiesDefinitionArgs struct {
	ManagedServices EncryptionV2PtrInput `pulumi:"managedServices"`
}

func (EncryptionEntitiesDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionEntitiesDefinition)(nil)).Elem()
}

func (i EncryptionEntitiesDefinitionArgs) ToEncryptionEntitiesDefinitionOutput() EncryptionEntitiesDefinitionOutput {
	return i.ToEncryptionEntitiesDefinitionOutputWithContext(context.Background())
}

func (i EncryptionEntitiesDefinitionArgs) ToEncryptionEntitiesDefinitionOutputWithContext(ctx context.Context) EncryptionEntitiesDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionEntitiesDefinitionOutput)
}

func (i EncryptionEntitiesDefinitionArgs) ToEncryptionEntitiesDefinitionPtrOutput() EncryptionEntitiesDefinitionPtrOutput {
	return i.ToEncryptionEntitiesDefinitionPtrOutputWithContext(context.Background())
}

func (i EncryptionEntitiesDefinitionArgs) ToEncryptionEntitiesDefinitionPtrOutputWithContext(ctx context.Context) EncryptionEntitiesDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionEntitiesDefinitionOutput).ToEncryptionEntitiesDefinitionPtrOutputWithContext(ctx)
}









type EncryptionEntitiesDefinitionPtrInput interface {
	pulumi.Input

	ToEncryptionEntitiesDefinitionPtrOutput() EncryptionEntitiesDefinitionPtrOutput
	ToEncryptionEntitiesDefinitionPtrOutputWithContext(context.Context) EncryptionEntitiesDefinitionPtrOutput
}

type encryptionEntitiesDefinitionPtrType EncryptionEntitiesDefinitionArgs

func EncryptionEntitiesDefinitionPtr(v *EncryptionEntitiesDefinitionArgs) EncryptionEntitiesDefinitionPtrInput {
	return (*encryptionEntitiesDefinitionPtrType)(v)
}

func (*encryptionEntitiesDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionEntitiesDefinition)(nil)).Elem()
}

func (i *encryptionEntitiesDefinitionPtrType) ToEncryptionEntitiesDefinitionPtrOutput() EncryptionEntitiesDefinitionPtrOutput {
	return i.ToEncryptionEntitiesDefinitionPtrOutputWithContext(context.Background())
}

func (i *encryptionEntitiesDefinitionPtrType) ToEncryptionEntitiesDefinitionPtrOutputWithContext(ctx context.Context) EncryptionEntitiesDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionEntitiesDefinitionPtrOutput)
}

type EncryptionEntitiesDefinitionOutput struct{ *pulumi.OutputState }

func (EncryptionEntitiesDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionEntitiesDefinition)(nil)).Elem()
}

func (o EncryptionEntitiesDefinitionOutput) ToEncryptionEntitiesDefinitionOutput() EncryptionEntitiesDefinitionOutput {
	return o
}

func (o EncryptionEntitiesDefinitionOutput) ToEncryptionEntitiesDefinitionOutputWithContext(ctx context.Context) EncryptionEntitiesDefinitionOutput {
	return o
}

func (o EncryptionEntitiesDefinitionOutput) ToEncryptionEntitiesDefinitionPtrOutput() EncryptionEntitiesDefinitionPtrOutput {
	return o.ToEncryptionEntitiesDefinitionPtrOutputWithContext(context.Background())
}

func (o EncryptionEntitiesDefinitionOutput) ToEncryptionEntitiesDefinitionPtrOutputWithContext(ctx context.Context) EncryptionEntitiesDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionEntitiesDefinition) *EncryptionEntitiesDefinition {
		return &v
	}).(EncryptionEntitiesDefinitionPtrOutput)
}

func (o EncryptionEntitiesDefinitionOutput) ManagedServices() EncryptionV2PtrOutput {
	return o.ApplyT(func(v EncryptionEntitiesDefinition) *EncryptionV2 { return v.ManagedServices }).(EncryptionV2PtrOutput)
}

type EncryptionEntitiesDefinitionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionEntitiesDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionEntitiesDefinition)(nil)).Elem()
}

func (o EncryptionEntitiesDefinitionPtrOutput) ToEncryptionEntitiesDefinitionPtrOutput() EncryptionEntitiesDefinitionPtrOutput {
	return o
}

func (o EncryptionEntitiesDefinitionPtrOutput) ToEncryptionEntitiesDefinitionPtrOutputWithContext(ctx context.Context) EncryptionEntitiesDefinitionPtrOutput {
	return o
}

func (o EncryptionEntitiesDefinitionPtrOutput) Elem() EncryptionEntitiesDefinitionOutput {
	return o.ApplyT(func(v *EncryptionEntitiesDefinition) EncryptionEntitiesDefinition {
		if v != nil {
			return *v
		}
		var ret EncryptionEntitiesDefinition
		return ret
	}).(EncryptionEntitiesDefinitionOutput)
}

func (o EncryptionEntitiesDefinitionPtrOutput) ManagedServices() EncryptionV2PtrOutput {
	return o.ApplyT(func(v *EncryptionEntitiesDefinition) *EncryptionV2 {
		if v == nil {
			return nil
		}
		return v.ManagedServices
	}).(EncryptionV2PtrOutput)
}

type EncryptionEntitiesDefinitionResponse struct {
	ManagedServices *EncryptionV2Response `pulumi:"managedServices"`
}

type EncryptionEntitiesDefinitionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionEntitiesDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionEntitiesDefinitionResponse)(nil)).Elem()
}

func (o EncryptionEntitiesDefinitionResponseOutput) ToEncryptionEntitiesDefinitionResponseOutput() EncryptionEntitiesDefinitionResponseOutput {
	return o
}

func (o EncryptionEntitiesDefinitionResponseOutput) ToEncryptionEntitiesDefinitionResponseOutputWithContext(ctx context.Context) EncryptionEntitiesDefinitionResponseOutput {
	return o
}

func (o EncryptionEntitiesDefinitionResponseOutput) ManagedServices() EncryptionV2ResponsePtrOutput {
	return o.ApplyT(func(v EncryptionEntitiesDefinitionResponse) *EncryptionV2Response { return v.ManagedServices }).(EncryptionV2ResponsePtrOutput)
}

type EncryptionEntitiesDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionEntitiesDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionEntitiesDefinitionResponse)(nil)).Elem()
}

func (o EncryptionEntitiesDefinitionResponsePtrOutput) ToEncryptionEntitiesDefinitionResponsePtrOutput() EncryptionEntitiesDefinitionResponsePtrOutput {
	return o
}

func (o EncryptionEntitiesDefinitionResponsePtrOutput) ToEncryptionEntitiesDefinitionResponsePtrOutputWithContext(ctx context.Context) EncryptionEntitiesDefinitionResponsePtrOutput {
	return o
}

func (o EncryptionEntitiesDefinitionResponsePtrOutput) Elem() EncryptionEntitiesDefinitionResponseOutput {
	return o.ApplyT(func(v *EncryptionEntitiesDefinitionResponse) EncryptionEntitiesDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionEntitiesDefinitionResponse
		return ret
	}).(EncryptionEntitiesDefinitionResponseOutput)
}

func (o EncryptionEntitiesDefinitionResponsePtrOutput) ManagedServices() EncryptionV2ResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionEntitiesDefinitionResponse) *EncryptionV2Response {
		if v == nil {
			return nil
		}
		return v.ManagedServices
	}).(EncryptionV2ResponsePtrOutput)
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

type EncryptionV2 struct {
	KeySource          string                          `pulumi:"keySource"`
	KeyVaultProperties *EncryptionV2KeyVaultProperties `pulumi:"keyVaultProperties"`
}





type EncryptionV2Input interface {
	pulumi.Input

	ToEncryptionV2Output() EncryptionV2Output
	ToEncryptionV2OutputWithContext(context.Context) EncryptionV2Output
}

type EncryptionV2Args struct {
	KeySource          pulumi.StringInput                     `pulumi:"keySource"`
	KeyVaultProperties EncryptionV2KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
}

func (EncryptionV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionV2)(nil)).Elem()
}

func (i EncryptionV2Args) ToEncryptionV2Output() EncryptionV2Output {
	return i.ToEncryptionV2OutputWithContext(context.Background())
}

func (i EncryptionV2Args) ToEncryptionV2OutputWithContext(ctx context.Context) EncryptionV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionV2Output)
}

func (i EncryptionV2Args) ToEncryptionV2PtrOutput() EncryptionV2PtrOutput {
	return i.ToEncryptionV2PtrOutputWithContext(context.Background())
}

func (i EncryptionV2Args) ToEncryptionV2PtrOutputWithContext(ctx context.Context) EncryptionV2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionV2Output).ToEncryptionV2PtrOutputWithContext(ctx)
}









type EncryptionV2PtrInput interface {
	pulumi.Input

	ToEncryptionV2PtrOutput() EncryptionV2PtrOutput
	ToEncryptionV2PtrOutputWithContext(context.Context) EncryptionV2PtrOutput
}

type encryptionV2PtrType EncryptionV2Args

func EncryptionV2Ptr(v *EncryptionV2Args) EncryptionV2PtrInput {
	return (*encryptionV2PtrType)(v)
}

func (*encryptionV2PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionV2)(nil)).Elem()
}

func (i *encryptionV2PtrType) ToEncryptionV2PtrOutput() EncryptionV2PtrOutput {
	return i.ToEncryptionV2PtrOutputWithContext(context.Background())
}

func (i *encryptionV2PtrType) ToEncryptionV2PtrOutputWithContext(ctx context.Context) EncryptionV2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionV2PtrOutput)
}

type EncryptionV2Output struct{ *pulumi.OutputState }

func (EncryptionV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionV2)(nil)).Elem()
}

func (o EncryptionV2Output) ToEncryptionV2Output() EncryptionV2Output {
	return o
}

func (o EncryptionV2Output) ToEncryptionV2OutputWithContext(ctx context.Context) EncryptionV2Output {
	return o
}

func (o EncryptionV2Output) ToEncryptionV2PtrOutput() EncryptionV2PtrOutput {
	return o.ToEncryptionV2PtrOutputWithContext(context.Background())
}

func (o EncryptionV2Output) ToEncryptionV2PtrOutputWithContext(ctx context.Context) EncryptionV2PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionV2) *EncryptionV2 {
		return &v
	}).(EncryptionV2PtrOutput)
}

func (o EncryptionV2Output) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionV2) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionV2Output) KeyVaultProperties() EncryptionV2KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v EncryptionV2) *EncryptionV2KeyVaultProperties { return v.KeyVaultProperties }).(EncryptionV2KeyVaultPropertiesPtrOutput)
}

type EncryptionV2PtrOutput struct{ *pulumi.OutputState }

func (EncryptionV2PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionV2)(nil)).Elem()
}

func (o EncryptionV2PtrOutput) ToEncryptionV2PtrOutput() EncryptionV2PtrOutput {
	return o
}

func (o EncryptionV2PtrOutput) ToEncryptionV2PtrOutputWithContext(ctx context.Context) EncryptionV2PtrOutput {
	return o
}

func (o EncryptionV2PtrOutput) Elem() EncryptionV2Output {
	return o.ApplyT(func(v *EncryptionV2) EncryptionV2 {
		if v != nil {
			return *v
		}
		var ret EncryptionV2
		return ret
	}).(EncryptionV2Output)
}

func (o EncryptionV2PtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionV2) *string {
		if v == nil {
			return nil
		}
		return &v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionV2PtrOutput) KeyVaultProperties() EncryptionV2KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *EncryptionV2) *EncryptionV2KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(EncryptionV2KeyVaultPropertiesPtrOutput)
}

type EncryptionV2KeyVaultProperties struct {
	KeyName     string `pulumi:"keyName"`
	KeyVaultUri string `pulumi:"keyVaultUri"`
	KeyVersion  string `pulumi:"keyVersion"`
}





type EncryptionV2KeyVaultPropertiesInput interface {
	pulumi.Input

	ToEncryptionV2KeyVaultPropertiesOutput() EncryptionV2KeyVaultPropertiesOutput
	ToEncryptionV2KeyVaultPropertiesOutputWithContext(context.Context) EncryptionV2KeyVaultPropertiesOutput
}

type EncryptionV2KeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringInput `pulumi:"keyName"`
	KeyVaultUri pulumi.StringInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringInput `pulumi:"keyVersion"`
}

func (EncryptionV2KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionV2KeyVaultProperties)(nil)).Elem()
}

func (i EncryptionV2KeyVaultPropertiesArgs) ToEncryptionV2KeyVaultPropertiesOutput() EncryptionV2KeyVaultPropertiesOutput {
	return i.ToEncryptionV2KeyVaultPropertiesOutputWithContext(context.Background())
}

func (i EncryptionV2KeyVaultPropertiesArgs) ToEncryptionV2KeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionV2KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionV2KeyVaultPropertiesOutput)
}

func (i EncryptionV2KeyVaultPropertiesArgs) ToEncryptionV2KeyVaultPropertiesPtrOutput() EncryptionV2KeyVaultPropertiesPtrOutput {
	return i.ToEncryptionV2KeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionV2KeyVaultPropertiesArgs) ToEncryptionV2KeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionV2KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionV2KeyVaultPropertiesOutput).ToEncryptionV2KeyVaultPropertiesPtrOutputWithContext(ctx)
}









type EncryptionV2KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionV2KeyVaultPropertiesPtrOutput() EncryptionV2KeyVaultPropertiesPtrOutput
	ToEncryptionV2KeyVaultPropertiesPtrOutputWithContext(context.Context) EncryptionV2KeyVaultPropertiesPtrOutput
}

type encryptionV2KeyVaultPropertiesPtrType EncryptionV2KeyVaultPropertiesArgs

func EncryptionV2KeyVaultPropertiesPtr(v *EncryptionV2KeyVaultPropertiesArgs) EncryptionV2KeyVaultPropertiesPtrInput {
	return (*encryptionV2KeyVaultPropertiesPtrType)(v)
}

func (*encryptionV2KeyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionV2KeyVaultProperties)(nil)).Elem()
}

func (i *encryptionV2KeyVaultPropertiesPtrType) ToEncryptionV2KeyVaultPropertiesPtrOutput() EncryptionV2KeyVaultPropertiesPtrOutput {
	return i.ToEncryptionV2KeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionV2KeyVaultPropertiesPtrType) ToEncryptionV2KeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionV2KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionV2KeyVaultPropertiesPtrOutput)
}

type EncryptionV2KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionV2KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionV2KeyVaultProperties)(nil)).Elem()
}

func (o EncryptionV2KeyVaultPropertiesOutput) ToEncryptionV2KeyVaultPropertiesOutput() EncryptionV2KeyVaultPropertiesOutput {
	return o
}

func (o EncryptionV2KeyVaultPropertiesOutput) ToEncryptionV2KeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionV2KeyVaultPropertiesOutput {
	return o
}

func (o EncryptionV2KeyVaultPropertiesOutput) ToEncryptionV2KeyVaultPropertiesPtrOutput() EncryptionV2KeyVaultPropertiesPtrOutput {
	return o.ToEncryptionV2KeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionV2KeyVaultPropertiesOutput) ToEncryptionV2KeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionV2KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionV2KeyVaultProperties) *EncryptionV2KeyVaultProperties {
		return &v
	}).(EncryptionV2KeyVaultPropertiesPtrOutput)
}

func (o EncryptionV2KeyVaultPropertiesOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionV2KeyVaultProperties) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o EncryptionV2KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionV2KeyVaultProperties) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o EncryptionV2KeyVaultPropertiesOutput) KeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionV2KeyVaultProperties) string { return v.KeyVersion }).(pulumi.StringOutput)
}

type EncryptionV2KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionV2KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionV2KeyVaultProperties)(nil)).Elem()
}

func (o EncryptionV2KeyVaultPropertiesPtrOutput) ToEncryptionV2KeyVaultPropertiesPtrOutput() EncryptionV2KeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionV2KeyVaultPropertiesPtrOutput) ToEncryptionV2KeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionV2KeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionV2KeyVaultPropertiesPtrOutput) Elem() EncryptionV2KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *EncryptionV2KeyVaultProperties) EncryptionV2KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionV2KeyVaultProperties
		return ret
	}).(EncryptionV2KeyVaultPropertiesOutput)
}

func (o EncryptionV2KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionV2KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionV2KeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionV2KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionV2KeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionV2KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type EncryptionV2Response struct {
	KeySource          string                                  `pulumi:"keySource"`
	KeyVaultProperties *EncryptionV2ResponseKeyVaultProperties `pulumi:"keyVaultProperties"`
}

type EncryptionV2ResponseOutput struct{ *pulumi.OutputState }

func (EncryptionV2ResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionV2Response)(nil)).Elem()
}

func (o EncryptionV2ResponseOutput) ToEncryptionV2ResponseOutput() EncryptionV2ResponseOutput {
	return o
}

func (o EncryptionV2ResponseOutput) ToEncryptionV2ResponseOutputWithContext(ctx context.Context) EncryptionV2ResponseOutput {
	return o
}

func (o EncryptionV2ResponseOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionV2Response) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionV2ResponseOutput) KeyVaultProperties() EncryptionV2ResponseKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v EncryptionV2Response) *EncryptionV2ResponseKeyVaultProperties { return v.KeyVaultProperties }).(EncryptionV2ResponseKeyVaultPropertiesPtrOutput)
}

type EncryptionV2ResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionV2ResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionV2Response)(nil)).Elem()
}

func (o EncryptionV2ResponsePtrOutput) ToEncryptionV2ResponsePtrOutput() EncryptionV2ResponsePtrOutput {
	return o
}

func (o EncryptionV2ResponsePtrOutput) ToEncryptionV2ResponsePtrOutputWithContext(ctx context.Context) EncryptionV2ResponsePtrOutput {
	return o
}

func (o EncryptionV2ResponsePtrOutput) Elem() EncryptionV2ResponseOutput {
	return o.ApplyT(func(v *EncryptionV2Response) EncryptionV2Response {
		if v != nil {
			return *v
		}
		var ret EncryptionV2Response
		return ret
	}).(EncryptionV2ResponseOutput)
}

func (o EncryptionV2ResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionV2Response) *string {
		if v == nil {
			return nil
		}
		return &v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionV2ResponsePtrOutput) KeyVaultProperties() EncryptionV2ResponseKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *EncryptionV2Response) *EncryptionV2ResponseKeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(EncryptionV2ResponseKeyVaultPropertiesPtrOutput)
}

type EncryptionV2ResponseKeyVaultProperties struct {
	KeyName     string `pulumi:"keyName"`
	KeyVaultUri string `pulumi:"keyVaultUri"`
	KeyVersion  string `pulumi:"keyVersion"`
}

type EncryptionV2ResponseKeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionV2ResponseKeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionV2ResponseKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionV2ResponseKeyVaultPropertiesOutput) ToEncryptionV2ResponseKeyVaultPropertiesOutput() EncryptionV2ResponseKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionV2ResponseKeyVaultPropertiesOutput) ToEncryptionV2ResponseKeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionV2ResponseKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionV2ResponseKeyVaultPropertiesOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionV2ResponseKeyVaultProperties) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o EncryptionV2ResponseKeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionV2ResponseKeyVaultProperties) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o EncryptionV2ResponseKeyVaultPropertiesOutput) KeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionV2ResponseKeyVaultProperties) string { return v.KeyVersion }).(pulumi.StringOutput)
}

type EncryptionV2ResponseKeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionV2ResponseKeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionV2ResponseKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionV2ResponseKeyVaultPropertiesPtrOutput) ToEncryptionV2ResponseKeyVaultPropertiesPtrOutput() EncryptionV2ResponseKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionV2ResponseKeyVaultPropertiesPtrOutput) ToEncryptionV2ResponseKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionV2ResponseKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionV2ResponseKeyVaultPropertiesPtrOutput) Elem() EncryptionV2ResponseKeyVaultPropertiesOutput {
	return o.ApplyT(func(v *EncryptionV2ResponseKeyVaultProperties) EncryptionV2ResponseKeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionV2ResponseKeyVaultProperties
		return ret
	}).(EncryptionV2ResponseKeyVaultPropertiesOutput)
}

func (o EncryptionV2ResponseKeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionV2ResponseKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionV2ResponseKeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionV2ResponseKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionV2ResponseKeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionV2ResponseKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVersion
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

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
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

type PrivateEndpointConnectionResponse struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
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

func (o PrivateLinkServiceConnectionStateOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         string  `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
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

type WorkspacePropertiesEncryption struct {
	Entities EncryptionEntitiesDefinition `pulumi:"entities"`
}





type WorkspacePropertiesEncryptionInput interface {
	pulumi.Input

	ToWorkspacePropertiesEncryptionOutput() WorkspacePropertiesEncryptionOutput
	ToWorkspacePropertiesEncryptionOutputWithContext(context.Context) WorkspacePropertiesEncryptionOutput
}

type WorkspacePropertiesEncryptionArgs struct {
	Entities EncryptionEntitiesDefinitionInput `pulumi:"entities"`
}

func (WorkspacePropertiesEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspacePropertiesEncryption)(nil)).Elem()
}

func (i WorkspacePropertiesEncryptionArgs) ToWorkspacePropertiesEncryptionOutput() WorkspacePropertiesEncryptionOutput {
	return i.ToWorkspacePropertiesEncryptionOutputWithContext(context.Background())
}

func (i WorkspacePropertiesEncryptionArgs) ToWorkspacePropertiesEncryptionOutputWithContext(ctx context.Context) WorkspacePropertiesEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePropertiesEncryptionOutput)
}

func (i WorkspacePropertiesEncryptionArgs) ToWorkspacePropertiesEncryptionPtrOutput() WorkspacePropertiesEncryptionPtrOutput {
	return i.ToWorkspacePropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (i WorkspacePropertiesEncryptionArgs) ToWorkspacePropertiesEncryptionPtrOutputWithContext(ctx context.Context) WorkspacePropertiesEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePropertiesEncryptionOutput).ToWorkspacePropertiesEncryptionPtrOutputWithContext(ctx)
}









type WorkspacePropertiesEncryptionPtrInput interface {
	pulumi.Input

	ToWorkspacePropertiesEncryptionPtrOutput() WorkspacePropertiesEncryptionPtrOutput
	ToWorkspacePropertiesEncryptionPtrOutputWithContext(context.Context) WorkspacePropertiesEncryptionPtrOutput
}

type workspacePropertiesEncryptionPtrType WorkspacePropertiesEncryptionArgs

func WorkspacePropertiesEncryptionPtr(v *WorkspacePropertiesEncryptionArgs) WorkspacePropertiesEncryptionPtrInput {
	return (*workspacePropertiesEncryptionPtrType)(v)
}

func (*workspacePropertiesEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePropertiesEncryption)(nil)).Elem()
}

func (i *workspacePropertiesEncryptionPtrType) ToWorkspacePropertiesEncryptionPtrOutput() WorkspacePropertiesEncryptionPtrOutput {
	return i.ToWorkspacePropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (i *workspacePropertiesEncryptionPtrType) ToWorkspacePropertiesEncryptionPtrOutputWithContext(ctx context.Context) WorkspacePropertiesEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePropertiesEncryptionPtrOutput)
}

type WorkspacePropertiesEncryptionOutput struct{ *pulumi.OutputState }

func (WorkspacePropertiesEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspacePropertiesEncryption)(nil)).Elem()
}

func (o WorkspacePropertiesEncryptionOutput) ToWorkspacePropertiesEncryptionOutput() WorkspacePropertiesEncryptionOutput {
	return o
}

func (o WorkspacePropertiesEncryptionOutput) ToWorkspacePropertiesEncryptionOutputWithContext(ctx context.Context) WorkspacePropertiesEncryptionOutput {
	return o
}

func (o WorkspacePropertiesEncryptionOutput) ToWorkspacePropertiesEncryptionPtrOutput() WorkspacePropertiesEncryptionPtrOutput {
	return o.ToWorkspacePropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (o WorkspacePropertiesEncryptionOutput) ToWorkspacePropertiesEncryptionPtrOutputWithContext(ctx context.Context) WorkspacePropertiesEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspacePropertiesEncryption) *WorkspacePropertiesEncryption {
		return &v
	}).(WorkspacePropertiesEncryptionPtrOutput)
}

func (o WorkspacePropertiesEncryptionOutput) Entities() EncryptionEntitiesDefinitionOutput {
	return o.ApplyT(func(v WorkspacePropertiesEncryption) EncryptionEntitiesDefinition { return v.Entities }).(EncryptionEntitiesDefinitionOutput)
}

type WorkspacePropertiesEncryptionPtrOutput struct{ *pulumi.OutputState }

func (WorkspacePropertiesEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePropertiesEncryption)(nil)).Elem()
}

func (o WorkspacePropertiesEncryptionPtrOutput) ToWorkspacePropertiesEncryptionPtrOutput() WorkspacePropertiesEncryptionPtrOutput {
	return o
}

func (o WorkspacePropertiesEncryptionPtrOutput) ToWorkspacePropertiesEncryptionPtrOutputWithContext(ctx context.Context) WorkspacePropertiesEncryptionPtrOutput {
	return o
}

func (o WorkspacePropertiesEncryptionPtrOutput) Elem() WorkspacePropertiesEncryptionOutput {
	return o.ApplyT(func(v *WorkspacePropertiesEncryption) WorkspacePropertiesEncryption {
		if v != nil {
			return *v
		}
		var ret WorkspacePropertiesEncryption
		return ret
	}).(WorkspacePropertiesEncryptionOutput)
}

func (o WorkspacePropertiesEncryptionPtrOutput) Entities() EncryptionEntitiesDefinitionPtrOutput {
	return o.ApplyT(func(v *WorkspacePropertiesEncryption) *EncryptionEntitiesDefinition {
		if v == nil {
			return nil
		}
		return &v.Entities
	}).(EncryptionEntitiesDefinitionPtrOutput)
}

type WorkspacePropertiesResponseEncryption struct {
	Entities EncryptionEntitiesDefinitionResponse `pulumi:"entities"`
}

type WorkspacePropertiesResponseEncryptionOutput struct{ *pulumi.OutputState }

func (WorkspacePropertiesResponseEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspacePropertiesResponseEncryption)(nil)).Elem()
}

func (o WorkspacePropertiesResponseEncryptionOutput) ToWorkspacePropertiesResponseEncryptionOutput() WorkspacePropertiesResponseEncryptionOutput {
	return o
}

func (o WorkspacePropertiesResponseEncryptionOutput) ToWorkspacePropertiesResponseEncryptionOutputWithContext(ctx context.Context) WorkspacePropertiesResponseEncryptionOutput {
	return o
}

func (o WorkspacePropertiesResponseEncryptionOutput) Entities() EncryptionEntitiesDefinitionResponseOutput {
	return o.ApplyT(func(v WorkspacePropertiesResponseEncryption) EncryptionEntitiesDefinitionResponse { return v.Entities }).(EncryptionEntitiesDefinitionResponseOutput)
}

type WorkspacePropertiesResponseEncryptionPtrOutput struct{ *pulumi.OutputState }

func (WorkspacePropertiesResponseEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePropertiesResponseEncryption)(nil)).Elem()
}

func (o WorkspacePropertiesResponseEncryptionPtrOutput) ToWorkspacePropertiesResponseEncryptionPtrOutput() WorkspacePropertiesResponseEncryptionPtrOutput {
	return o
}

func (o WorkspacePropertiesResponseEncryptionPtrOutput) ToWorkspacePropertiesResponseEncryptionPtrOutputWithContext(ctx context.Context) WorkspacePropertiesResponseEncryptionPtrOutput {
	return o
}

func (o WorkspacePropertiesResponseEncryptionPtrOutput) Elem() WorkspacePropertiesResponseEncryptionOutput {
	return o.ApplyT(func(v *WorkspacePropertiesResponseEncryption) WorkspacePropertiesResponseEncryption {
		if v != nil {
			return *v
		}
		var ret WorkspacePropertiesResponseEncryption
		return ret
	}).(WorkspacePropertiesResponseEncryptionOutput)
}

func (o WorkspacePropertiesResponseEncryptionPtrOutput) Entities() EncryptionEntitiesDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *WorkspacePropertiesResponseEncryption) *EncryptionEntitiesDefinitionResponse {
		if v == nil {
			return nil
		}
		return &v.Entities
	}).(EncryptionEntitiesDefinitionResponsePtrOutput)
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
	pulumi.RegisterOutputType(EncryptionEntitiesDefinitionOutput{})
	pulumi.RegisterOutputType(EncryptionEntitiesDefinitionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionEntitiesDefinitionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionEntitiesDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionV2Output{})
	pulumi.RegisterOutputType(EncryptionV2PtrOutput{})
	pulumi.RegisterOutputType(EncryptionV2KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionV2KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionV2ResponseOutput{})
	pulumi.RegisterOutputType(EncryptionV2ResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionV2ResponseKeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionV2ResponseKeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
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
	pulumi.RegisterOutputType(WorkspacePropertiesEncryptionOutput{})
	pulumi.RegisterOutputType(WorkspacePropertiesEncryptionPtrOutput{})
	pulumi.RegisterOutputType(WorkspacePropertiesResponseEncryptionOutput{})
	pulumi.RegisterOutputType(WorkspacePropertiesResponseEncryptionPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceProviderAuthorizationOutput{})
	pulumi.RegisterOutputType(WorkspaceProviderAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceProviderAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceProviderAuthorizationResponseArrayOutput{})
}
