


package v20201030preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnterprisePolicyIdentity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type EnterprisePolicyIdentityInput interface {
	pulumi.Input

	ToEnterprisePolicyIdentityOutput() EnterprisePolicyIdentityOutput
	ToEnterprisePolicyIdentityOutputWithContext(context.Context) EnterprisePolicyIdentityOutput
}

type EnterprisePolicyIdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (EnterprisePolicyIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterprisePolicyIdentity)(nil)).Elem()
}

func (i EnterprisePolicyIdentityArgs) ToEnterprisePolicyIdentityOutput() EnterprisePolicyIdentityOutput {
	return i.ToEnterprisePolicyIdentityOutputWithContext(context.Background())
}

func (i EnterprisePolicyIdentityArgs) ToEnterprisePolicyIdentityOutputWithContext(ctx context.Context) EnterprisePolicyIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePolicyIdentityOutput)
}

func (i EnterprisePolicyIdentityArgs) ToEnterprisePolicyIdentityPtrOutput() EnterprisePolicyIdentityPtrOutput {
	return i.ToEnterprisePolicyIdentityPtrOutputWithContext(context.Background())
}

func (i EnterprisePolicyIdentityArgs) ToEnterprisePolicyIdentityPtrOutputWithContext(ctx context.Context) EnterprisePolicyIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePolicyIdentityOutput).ToEnterprisePolicyIdentityPtrOutputWithContext(ctx)
}









type EnterprisePolicyIdentityPtrInput interface {
	pulumi.Input

	ToEnterprisePolicyIdentityPtrOutput() EnterprisePolicyIdentityPtrOutput
	ToEnterprisePolicyIdentityPtrOutputWithContext(context.Context) EnterprisePolicyIdentityPtrOutput
}

type enterprisePolicyIdentityPtrType EnterprisePolicyIdentityArgs

func EnterprisePolicyIdentityPtr(v *EnterprisePolicyIdentityArgs) EnterprisePolicyIdentityPtrInput {
	return (*enterprisePolicyIdentityPtrType)(v)
}

func (*enterprisePolicyIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePolicyIdentity)(nil)).Elem()
}

func (i *enterprisePolicyIdentityPtrType) ToEnterprisePolicyIdentityPtrOutput() EnterprisePolicyIdentityPtrOutput {
	return i.ToEnterprisePolicyIdentityPtrOutputWithContext(context.Background())
}

func (i *enterprisePolicyIdentityPtrType) ToEnterprisePolicyIdentityPtrOutputWithContext(ctx context.Context) EnterprisePolicyIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePolicyIdentityPtrOutput)
}

type EnterprisePolicyIdentityOutput struct{ *pulumi.OutputState }

func (EnterprisePolicyIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterprisePolicyIdentity)(nil)).Elem()
}

func (o EnterprisePolicyIdentityOutput) ToEnterprisePolicyIdentityOutput() EnterprisePolicyIdentityOutput {
	return o
}

func (o EnterprisePolicyIdentityOutput) ToEnterprisePolicyIdentityOutputWithContext(ctx context.Context) EnterprisePolicyIdentityOutput {
	return o
}

func (o EnterprisePolicyIdentityOutput) ToEnterprisePolicyIdentityPtrOutput() EnterprisePolicyIdentityPtrOutput {
	return o.ToEnterprisePolicyIdentityPtrOutputWithContext(context.Background())
}

func (o EnterprisePolicyIdentityOutput) ToEnterprisePolicyIdentityPtrOutputWithContext(ctx context.Context) EnterprisePolicyIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnterprisePolicyIdentity) *EnterprisePolicyIdentity {
		return &v
	}).(EnterprisePolicyIdentityPtrOutput)
}

func (o EnterprisePolicyIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v EnterprisePolicyIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type EnterprisePolicyIdentityPtrOutput struct{ *pulumi.OutputState }

func (EnterprisePolicyIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePolicyIdentity)(nil)).Elem()
}

func (o EnterprisePolicyIdentityPtrOutput) ToEnterprisePolicyIdentityPtrOutput() EnterprisePolicyIdentityPtrOutput {
	return o
}

func (o EnterprisePolicyIdentityPtrOutput) ToEnterprisePolicyIdentityPtrOutputWithContext(ctx context.Context) EnterprisePolicyIdentityPtrOutput {
	return o
}

func (o EnterprisePolicyIdentityPtrOutput) Elem() EnterprisePolicyIdentityOutput {
	return o.ApplyT(func(v *EnterprisePolicyIdentity) EnterprisePolicyIdentity {
		if v != nil {
			return *v
		}
		var ret EnterprisePolicyIdentity
		return ret
	}).(EnterprisePolicyIdentityOutput)
}

func (o EnterprisePolicyIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *EnterprisePolicyIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type EnterprisePolicyIdentityResponse struct {
	SystemAssignedIdentityPrincipalId string  `pulumi:"systemAssignedIdentityPrincipalId"`
	TenantId                          string  `pulumi:"tenantId"`
	Type                              *string `pulumi:"type"`
}





type EnterprisePolicyIdentityResponseInput interface {
	pulumi.Input

	ToEnterprisePolicyIdentityResponseOutput() EnterprisePolicyIdentityResponseOutput
	ToEnterprisePolicyIdentityResponseOutputWithContext(context.Context) EnterprisePolicyIdentityResponseOutput
}

type EnterprisePolicyIdentityResponseArgs struct {
	SystemAssignedIdentityPrincipalId pulumi.StringInput    `pulumi:"systemAssignedIdentityPrincipalId"`
	TenantId                          pulumi.StringInput    `pulumi:"tenantId"`
	Type                              pulumi.StringPtrInput `pulumi:"type"`
}

func (EnterprisePolicyIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterprisePolicyIdentityResponse)(nil)).Elem()
}

func (i EnterprisePolicyIdentityResponseArgs) ToEnterprisePolicyIdentityResponseOutput() EnterprisePolicyIdentityResponseOutput {
	return i.ToEnterprisePolicyIdentityResponseOutputWithContext(context.Background())
}

func (i EnterprisePolicyIdentityResponseArgs) ToEnterprisePolicyIdentityResponseOutputWithContext(ctx context.Context) EnterprisePolicyIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePolicyIdentityResponseOutput)
}

func (i EnterprisePolicyIdentityResponseArgs) ToEnterprisePolicyIdentityResponsePtrOutput() EnterprisePolicyIdentityResponsePtrOutput {
	return i.ToEnterprisePolicyIdentityResponsePtrOutputWithContext(context.Background())
}

func (i EnterprisePolicyIdentityResponseArgs) ToEnterprisePolicyIdentityResponsePtrOutputWithContext(ctx context.Context) EnterprisePolicyIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePolicyIdentityResponseOutput).ToEnterprisePolicyIdentityResponsePtrOutputWithContext(ctx)
}









type EnterprisePolicyIdentityResponsePtrInput interface {
	pulumi.Input

	ToEnterprisePolicyIdentityResponsePtrOutput() EnterprisePolicyIdentityResponsePtrOutput
	ToEnterprisePolicyIdentityResponsePtrOutputWithContext(context.Context) EnterprisePolicyIdentityResponsePtrOutput
}

type enterprisePolicyIdentityResponsePtrType EnterprisePolicyIdentityResponseArgs

func EnterprisePolicyIdentityResponsePtr(v *EnterprisePolicyIdentityResponseArgs) EnterprisePolicyIdentityResponsePtrInput {
	return (*enterprisePolicyIdentityResponsePtrType)(v)
}

func (*enterprisePolicyIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePolicyIdentityResponse)(nil)).Elem()
}

func (i *enterprisePolicyIdentityResponsePtrType) ToEnterprisePolicyIdentityResponsePtrOutput() EnterprisePolicyIdentityResponsePtrOutput {
	return i.ToEnterprisePolicyIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *enterprisePolicyIdentityResponsePtrType) ToEnterprisePolicyIdentityResponsePtrOutputWithContext(ctx context.Context) EnterprisePolicyIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePolicyIdentityResponsePtrOutput)
}

type EnterprisePolicyIdentityResponseOutput struct{ *pulumi.OutputState }

func (EnterprisePolicyIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterprisePolicyIdentityResponse)(nil)).Elem()
}

func (o EnterprisePolicyIdentityResponseOutput) ToEnterprisePolicyIdentityResponseOutput() EnterprisePolicyIdentityResponseOutput {
	return o
}

func (o EnterprisePolicyIdentityResponseOutput) ToEnterprisePolicyIdentityResponseOutputWithContext(ctx context.Context) EnterprisePolicyIdentityResponseOutput {
	return o
}

func (o EnterprisePolicyIdentityResponseOutput) ToEnterprisePolicyIdentityResponsePtrOutput() EnterprisePolicyIdentityResponsePtrOutput {
	return o.ToEnterprisePolicyIdentityResponsePtrOutputWithContext(context.Background())
}

func (o EnterprisePolicyIdentityResponseOutput) ToEnterprisePolicyIdentityResponsePtrOutputWithContext(ctx context.Context) EnterprisePolicyIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnterprisePolicyIdentityResponse) *EnterprisePolicyIdentityResponse {
		return &v
	}).(EnterprisePolicyIdentityResponsePtrOutput)
}

func (o EnterprisePolicyIdentityResponseOutput) SystemAssignedIdentityPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EnterprisePolicyIdentityResponse) string { return v.SystemAssignedIdentityPrincipalId }).(pulumi.StringOutput)
}

func (o EnterprisePolicyIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v EnterprisePolicyIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o EnterprisePolicyIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterprisePolicyIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type EnterprisePolicyIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (EnterprisePolicyIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePolicyIdentityResponse)(nil)).Elem()
}

func (o EnterprisePolicyIdentityResponsePtrOutput) ToEnterprisePolicyIdentityResponsePtrOutput() EnterprisePolicyIdentityResponsePtrOutput {
	return o
}

func (o EnterprisePolicyIdentityResponsePtrOutput) ToEnterprisePolicyIdentityResponsePtrOutputWithContext(ctx context.Context) EnterprisePolicyIdentityResponsePtrOutput {
	return o
}

func (o EnterprisePolicyIdentityResponsePtrOutput) Elem() EnterprisePolicyIdentityResponseOutput {
	return o.ApplyT(func(v *EnterprisePolicyIdentityResponse) EnterprisePolicyIdentityResponse {
		if v != nil {
			return *v
		}
		var ret EnterprisePolicyIdentityResponse
		return ret
	}).(EnterprisePolicyIdentityResponseOutput)
}

func (o EnterprisePolicyIdentityResponsePtrOutput) SystemAssignedIdentityPrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicyIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SystemAssignedIdentityPrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o EnterprisePolicyIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicyIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o EnterprisePolicyIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicyIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type KeyProperties struct {
	Name    *string `pulumi:"name"`
	Version *string `pulumi:"version"`
}





type KeyPropertiesInput interface {
	pulumi.Input

	ToKeyPropertiesOutput() KeyPropertiesOutput
	ToKeyPropertiesOutputWithContext(context.Context) KeyPropertiesOutput
}

type KeyPropertiesArgs struct {
	Name    pulumi.StringPtrInput `pulumi:"name"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (KeyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyProperties)(nil)).Elem()
}

func (i KeyPropertiesArgs) ToKeyPropertiesOutput() KeyPropertiesOutput {
	return i.ToKeyPropertiesOutputWithContext(context.Background())
}

func (i KeyPropertiesArgs) ToKeyPropertiesOutputWithContext(ctx context.Context) KeyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPropertiesOutput)
}

func (i KeyPropertiesArgs) ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput {
	return i.ToKeyPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyPropertiesArgs) ToKeyPropertiesPtrOutputWithContext(ctx context.Context) KeyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPropertiesOutput).ToKeyPropertiesPtrOutputWithContext(ctx)
}









type KeyPropertiesPtrInput interface {
	pulumi.Input

	ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput
	ToKeyPropertiesPtrOutputWithContext(context.Context) KeyPropertiesPtrOutput
}

type keyPropertiesPtrType KeyPropertiesArgs

func KeyPropertiesPtr(v *KeyPropertiesArgs) KeyPropertiesPtrInput {
	return (*keyPropertiesPtrType)(v)
}

func (*keyPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyProperties)(nil)).Elem()
}

func (i *keyPropertiesPtrType) ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput {
	return i.ToKeyPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyPropertiesPtrType) ToKeyPropertiesPtrOutputWithContext(ctx context.Context) KeyPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPropertiesPtrOutput)
}

type KeyPropertiesOutput struct{ *pulumi.OutputState }

func (KeyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyProperties)(nil)).Elem()
}

func (o KeyPropertiesOutput) ToKeyPropertiesOutput() KeyPropertiesOutput {
	return o
}

func (o KeyPropertiesOutput) ToKeyPropertiesOutputWithContext(ctx context.Context) KeyPropertiesOutput {
	return o
}

func (o KeyPropertiesOutput) ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput {
	return o.ToKeyPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyPropertiesOutput) ToKeyPropertiesPtrOutputWithContext(ctx context.Context) KeyPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyProperties) *KeyProperties {
		return &v
	}).(KeyPropertiesPtrOutput)
}

func (o KeyPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o KeyPropertiesOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyProperties) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type KeyPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyProperties)(nil)).Elem()
}

func (o KeyPropertiesPtrOutput) ToKeyPropertiesPtrOutput() KeyPropertiesPtrOutput {
	return o
}

func (o KeyPropertiesPtrOutput) ToKeyPropertiesPtrOutputWithContext(ctx context.Context) KeyPropertiesPtrOutput {
	return o
}

func (o KeyPropertiesPtrOutput) Elem() KeyPropertiesOutput {
	return o.ApplyT(func(v *KeyProperties) KeyProperties {
		if v != nil {
			return *v
		}
		var ret KeyProperties
		return ret
	}).(KeyPropertiesOutput)
}

func (o KeyPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o KeyPropertiesPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type KeyPropertiesResponse struct {
	Name    *string `pulumi:"name"`
	Version *string `pulumi:"version"`
}





type KeyPropertiesResponseInput interface {
	pulumi.Input

	ToKeyPropertiesResponseOutput() KeyPropertiesResponseOutput
	ToKeyPropertiesResponseOutputWithContext(context.Context) KeyPropertiesResponseOutput
}

type KeyPropertiesResponseArgs struct {
	Name    pulumi.StringPtrInput `pulumi:"name"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (KeyPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPropertiesResponse)(nil)).Elem()
}

func (i KeyPropertiesResponseArgs) ToKeyPropertiesResponseOutput() KeyPropertiesResponseOutput {
	return i.ToKeyPropertiesResponseOutputWithContext(context.Background())
}

func (i KeyPropertiesResponseArgs) ToKeyPropertiesResponseOutputWithContext(ctx context.Context) KeyPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPropertiesResponseOutput)
}

func (i KeyPropertiesResponseArgs) ToKeyPropertiesResponsePtrOutput() KeyPropertiesResponsePtrOutput {
	return i.ToKeyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KeyPropertiesResponseArgs) ToKeyPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPropertiesResponseOutput).ToKeyPropertiesResponsePtrOutputWithContext(ctx)
}









type KeyPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKeyPropertiesResponsePtrOutput() KeyPropertiesResponsePtrOutput
	ToKeyPropertiesResponsePtrOutputWithContext(context.Context) KeyPropertiesResponsePtrOutput
}

type keyPropertiesResponsePtrType KeyPropertiesResponseArgs

func KeyPropertiesResponsePtr(v *KeyPropertiesResponseArgs) KeyPropertiesResponsePtrInput {
	return (*keyPropertiesResponsePtrType)(v)
}

func (*keyPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPropertiesResponse)(nil)).Elem()
}

func (i *keyPropertiesResponsePtrType) ToKeyPropertiesResponsePtrOutput() KeyPropertiesResponsePtrOutput {
	return i.ToKeyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *keyPropertiesResponsePtrType) ToKeyPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPropertiesResponsePtrOutput)
}

type KeyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPropertiesResponse)(nil)).Elem()
}

func (o KeyPropertiesResponseOutput) ToKeyPropertiesResponseOutput() KeyPropertiesResponseOutput {
	return o
}

func (o KeyPropertiesResponseOutput) ToKeyPropertiesResponseOutputWithContext(ctx context.Context) KeyPropertiesResponseOutput {
	return o
}

func (o KeyPropertiesResponseOutput) ToKeyPropertiesResponsePtrOutput() KeyPropertiesResponsePtrOutput {
	return o.ToKeyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KeyPropertiesResponseOutput) ToKeyPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyPropertiesResponse) *KeyPropertiesResponse {
		return &v
	}).(KeyPropertiesResponsePtrOutput)
}

func (o KeyPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o KeyPropertiesResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyPropertiesResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type KeyPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPropertiesResponse)(nil)).Elem()
}

func (o KeyPropertiesResponsePtrOutput) ToKeyPropertiesResponsePtrOutput() KeyPropertiesResponsePtrOutput {
	return o
}

func (o KeyPropertiesResponsePtrOutput) ToKeyPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyPropertiesResponsePtrOutput {
	return o
}

func (o KeyPropertiesResponsePtrOutput) Elem() KeyPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyPropertiesResponse) KeyPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyPropertiesResponse
		return ret
	}).(KeyPropertiesResponseOutput)
}

func (o KeyPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o KeyPropertiesResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type KeyVaultProperties struct {
	Id  *string        `pulumi:"id"`
	Key *KeyProperties `pulumi:"key"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	Id  pulumi.StringPtrInput `pulumi:"id"`
	Key KeyPropertiesPtrInput `pulumi:"key"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) Key() KeyPropertiesPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *KeyProperties { return v.Key }).(KeyPropertiesPtrOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) Key() KeyPropertiesPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *KeyProperties {
		if v == nil {
			return nil
		}
		return v.Key
	}).(KeyPropertiesPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	Id  *string                `pulumi:"id"`
	Key *KeyPropertiesResponse `pulumi:"key"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	Id  pulumi.StringPtrInput         `pulumi:"id"`
	Key KeyPropertiesResponsePtrInput `pulumi:"key"`
}

func (KeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return i.ToKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput)
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput).ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type KeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput
	ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) KeyVaultPropertiesResponsePtrOutput
}

type keyVaultPropertiesResponsePtrType KeyVaultPropertiesResponseArgs

func KeyVaultPropertiesResponsePtr(v *KeyVaultPropertiesResponseArgs) KeyVaultPropertiesResponsePtrInput {
	return (*keyVaultPropertiesResponsePtrType)(v)
}

func (*keyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponsePtrOutput)
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultPropertiesResponse) *KeyVaultPropertiesResponse {
		return &v
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) Key() KeyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *KeyPropertiesResponse { return v.Key }).(KeyPropertiesResponsePtrOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) Key() KeyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *KeyPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Key
	}).(KeyPropertiesResponsePtrOutput)
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

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
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

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
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
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
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

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
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
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PropertiesEncryption struct {
	KeyVault *KeyVaultProperties `pulumi:"keyVault"`
	State    *string             `pulumi:"state"`
}





type PropertiesEncryptionInput interface {
	pulumi.Input

	ToPropertiesEncryptionOutput() PropertiesEncryptionOutput
	ToPropertiesEncryptionOutputWithContext(context.Context) PropertiesEncryptionOutput
}

type PropertiesEncryptionArgs struct {
	KeyVault KeyVaultPropertiesPtrInput `pulumi:"keyVault"`
	State    pulumi.StringPtrInput      `pulumi:"state"`
}

func (PropertiesEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesEncryption)(nil)).Elem()
}

func (i PropertiesEncryptionArgs) ToPropertiesEncryptionOutput() PropertiesEncryptionOutput {
	return i.ToPropertiesEncryptionOutputWithContext(context.Background())
}

func (i PropertiesEncryptionArgs) ToPropertiesEncryptionOutputWithContext(ctx context.Context) PropertiesEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesEncryptionOutput)
}

func (i PropertiesEncryptionArgs) ToPropertiesEncryptionPtrOutput() PropertiesEncryptionPtrOutput {
	return i.ToPropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (i PropertiesEncryptionArgs) ToPropertiesEncryptionPtrOutputWithContext(ctx context.Context) PropertiesEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesEncryptionOutput).ToPropertiesEncryptionPtrOutputWithContext(ctx)
}









type PropertiesEncryptionPtrInput interface {
	pulumi.Input

	ToPropertiesEncryptionPtrOutput() PropertiesEncryptionPtrOutput
	ToPropertiesEncryptionPtrOutputWithContext(context.Context) PropertiesEncryptionPtrOutput
}

type propertiesEncryptionPtrType PropertiesEncryptionArgs

func PropertiesEncryptionPtr(v *PropertiesEncryptionArgs) PropertiesEncryptionPtrInput {
	return (*propertiesEncryptionPtrType)(v)
}

func (*propertiesEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesEncryption)(nil)).Elem()
}

func (i *propertiesEncryptionPtrType) ToPropertiesEncryptionPtrOutput() PropertiesEncryptionPtrOutput {
	return i.ToPropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (i *propertiesEncryptionPtrType) ToPropertiesEncryptionPtrOutputWithContext(ctx context.Context) PropertiesEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesEncryptionPtrOutput)
}

type PropertiesEncryptionOutput struct{ *pulumi.OutputState }

func (PropertiesEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesEncryption)(nil)).Elem()
}

func (o PropertiesEncryptionOutput) ToPropertiesEncryptionOutput() PropertiesEncryptionOutput {
	return o
}

func (o PropertiesEncryptionOutput) ToPropertiesEncryptionOutputWithContext(ctx context.Context) PropertiesEncryptionOutput {
	return o
}

func (o PropertiesEncryptionOutput) ToPropertiesEncryptionPtrOutput() PropertiesEncryptionPtrOutput {
	return o.ToPropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (o PropertiesEncryptionOutput) ToPropertiesEncryptionPtrOutputWithContext(ctx context.Context) PropertiesEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PropertiesEncryption) *PropertiesEncryption {
		return &v
	}).(PropertiesEncryptionPtrOutput)
}

func (o PropertiesEncryptionOutput) KeyVault() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v PropertiesEncryption) *KeyVaultProperties { return v.KeyVault }).(KeyVaultPropertiesPtrOutput)
}

func (o PropertiesEncryptionOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertiesEncryption) *string { return v.State }).(pulumi.StringPtrOutput)
}

type PropertiesEncryptionPtrOutput struct{ *pulumi.OutputState }

func (PropertiesEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesEncryption)(nil)).Elem()
}

func (o PropertiesEncryptionPtrOutput) ToPropertiesEncryptionPtrOutput() PropertiesEncryptionPtrOutput {
	return o
}

func (o PropertiesEncryptionPtrOutput) ToPropertiesEncryptionPtrOutputWithContext(ctx context.Context) PropertiesEncryptionPtrOutput {
	return o
}

func (o PropertiesEncryptionPtrOutput) Elem() PropertiesEncryptionOutput {
	return o.ApplyT(func(v *PropertiesEncryption) PropertiesEncryption {
		if v != nil {
			return *v
		}
		var ret PropertiesEncryption
		return ret
	}).(PropertiesEncryptionOutput)
}

func (o PropertiesEncryptionPtrOutput) KeyVault() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *PropertiesEncryption) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVault
	}).(KeyVaultPropertiesPtrOutput)
}

func (o PropertiesEncryptionPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PropertiesEncryption) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type PropertiesLockbox struct {
	State *string `pulumi:"state"`
}





type PropertiesLockboxInput interface {
	pulumi.Input

	ToPropertiesLockboxOutput() PropertiesLockboxOutput
	ToPropertiesLockboxOutputWithContext(context.Context) PropertiesLockboxOutput
}

type PropertiesLockboxArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (PropertiesLockboxArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesLockbox)(nil)).Elem()
}

func (i PropertiesLockboxArgs) ToPropertiesLockboxOutput() PropertiesLockboxOutput {
	return i.ToPropertiesLockboxOutputWithContext(context.Background())
}

func (i PropertiesLockboxArgs) ToPropertiesLockboxOutputWithContext(ctx context.Context) PropertiesLockboxOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesLockboxOutput)
}

func (i PropertiesLockboxArgs) ToPropertiesLockboxPtrOutput() PropertiesLockboxPtrOutput {
	return i.ToPropertiesLockboxPtrOutputWithContext(context.Background())
}

func (i PropertiesLockboxArgs) ToPropertiesLockboxPtrOutputWithContext(ctx context.Context) PropertiesLockboxPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesLockboxOutput).ToPropertiesLockboxPtrOutputWithContext(ctx)
}









type PropertiesLockboxPtrInput interface {
	pulumi.Input

	ToPropertiesLockboxPtrOutput() PropertiesLockboxPtrOutput
	ToPropertiesLockboxPtrOutputWithContext(context.Context) PropertiesLockboxPtrOutput
}

type propertiesLockboxPtrType PropertiesLockboxArgs

func PropertiesLockboxPtr(v *PropertiesLockboxArgs) PropertiesLockboxPtrInput {
	return (*propertiesLockboxPtrType)(v)
}

func (*propertiesLockboxPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesLockbox)(nil)).Elem()
}

func (i *propertiesLockboxPtrType) ToPropertiesLockboxPtrOutput() PropertiesLockboxPtrOutput {
	return i.ToPropertiesLockboxPtrOutputWithContext(context.Background())
}

func (i *propertiesLockboxPtrType) ToPropertiesLockboxPtrOutputWithContext(ctx context.Context) PropertiesLockboxPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesLockboxPtrOutput)
}

type PropertiesLockboxOutput struct{ *pulumi.OutputState }

func (PropertiesLockboxOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesLockbox)(nil)).Elem()
}

func (o PropertiesLockboxOutput) ToPropertiesLockboxOutput() PropertiesLockboxOutput {
	return o
}

func (o PropertiesLockboxOutput) ToPropertiesLockboxOutputWithContext(ctx context.Context) PropertiesLockboxOutput {
	return o
}

func (o PropertiesLockboxOutput) ToPropertiesLockboxPtrOutput() PropertiesLockboxPtrOutput {
	return o.ToPropertiesLockboxPtrOutputWithContext(context.Background())
}

func (o PropertiesLockboxOutput) ToPropertiesLockboxPtrOutputWithContext(ctx context.Context) PropertiesLockboxPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PropertiesLockbox) *PropertiesLockbox {
		return &v
	}).(PropertiesLockboxPtrOutput)
}

func (o PropertiesLockboxOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertiesLockbox) *string { return v.State }).(pulumi.StringPtrOutput)
}

type PropertiesLockboxPtrOutput struct{ *pulumi.OutputState }

func (PropertiesLockboxPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesLockbox)(nil)).Elem()
}

func (o PropertiesLockboxPtrOutput) ToPropertiesLockboxPtrOutput() PropertiesLockboxPtrOutput {
	return o
}

func (o PropertiesLockboxPtrOutput) ToPropertiesLockboxPtrOutputWithContext(ctx context.Context) PropertiesLockboxPtrOutput {
	return o
}

func (o PropertiesLockboxPtrOutput) Elem() PropertiesLockboxOutput {
	return o.ApplyT(func(v *PropertiesLockbox) PropertiesLockbox {
		if v != nil {
			return *v
		}
		var ret PropertiesLockbox
		return ret
	}).(PropertiesLockboxOutput)
}

func (o PropertiesLockboxPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PropertiesLockbox) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type PropertiesNetworkInjection struct {
	VirtualNetworks *VirtualNetworkPropertiesList `pulumi:"virtualNetworks"`
}





type PropertiesNetworkInjectionInput interface {
	pulumi.Input

	ToPropertiesNetworkInjectionOutput() PropertiesNetworkInjectionOutput
	ToPropertiesNetworkInjectionOutputWithContext(context.Context) PropertiesNetworkInjectionOutput
}

type PropertiesNetworkInjectionArgs struct {
	VirtualNetworks VirtualNetworkPropertiesListPtrInput `pulumi:"virtualNetworks"`
}

func (PropertiesNetworkInjectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesNetworkInjection)(nil)).Elem()
}

func (i PropertiesNetworkInjectionArgs) ToPropertiesNetworkInjectionOutput() PropertiesNetworkInjectionOutput {
	return i.ToPropertiesNetworkInjectionOutputWithContext(context.Background())
}

func (i PropertiesNetworkInjectionArgs) ToPropertiesNetworkInjectionOutputWithContext(ctx context.Context) PropertiesNetworkInjectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesNetworkInjectionOutput)
}

func (i PropertiesNetworkInjectionArgs) ToPropertiesNetworkInjectionPtrOutput() PropertiesNetworkInjectionPtrOutput {
	return i.ToPropertiesNetworkInjectionPtrOutputWithContext(context.Background())
}

func (i PropertiesNetworkInjectionArgs) ToPropertiesNetworkInjectionPtrOutputWithContext(ctx context.Context) PropertiesNetworkInjectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesNetworkInjectionOutput).ToPropertiesNetworkInjectionPtrOutputWithContext(ctx)
}









type PropertiesNetworkInjectionPtrInput interface {
	pulumi.Input

	ToPropertiesNetworkInjectionPtrOutput() PropertiesNetworkInjectionPtrOutput
	ToPropertiesNetworkInjectionPtrOutputWithContext(context.Context) PropertiesNetworkInjectionPtrOutput
}

type propertiesNetworkInjectionPtrType PropertiesNetworkInjectionArgs

func PropertiesNetworkInjectionPtr(v *PropertiesNetworkInjectionArgs) PropertiesNetworkInjectionPtrInput {
	return (*propertiesNetworkInjectionPtrType)(v)
}

func (*propertiesNetworkInjectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesNetworkInjection)(nil)).Elem()
}

func (i *propertiesNetworkInjectionPtrType) ToPropertiesNetworkInjectionPtrOutput() PropertiesNetworkInjectionPtrOutput {
	return i.ToPropertiesNetworkInjectionPtrOutputWithContext(context.Background())
}

func (i *propertiesNetworkInjectionPtrType) ToPropertiesNetworkInjectionPtrOutputWithContext(ctx context.Context) PropertiesNetworkInjectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesNetworkInjectionPtrOutput)
}

type PropertiesNetworkInjectionOutput struct{ *pulumi.OutputState }

func (PropertiesNetworkInjectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesNetworkInjection)(nil)).Elem()
}

func (o PropertiesNetworkInjectionOutput) ToPropertiesNetworkInjectionOutput() PropertiesNetworkInjectionOutput {
	return o
}

func (o PropertiesNetworkInjectionOutput) ToPropertiesNetworkInjectionOutputWithContext(ctx context.Context) PropertiesNetworkInjectionOutput {
	return o
}

func (o PropertiesNetworkInjectionOutput) ToPropertiesNetworkInjectionPtrOutput() PropertiesNetworkInjectionPtrOutput {
	return o.ToPropertiesNetworkInjectionPtrOutputWithContext(context.Background())
}

func (o PropertiesNetworkInjectionOutput) ToPropertiesNetworkInjectionPtrOutputWithContext(ctx context.Context) PropertiesNetworkInjectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PropertiesNetworkInjection) *PropertiesNetworkInjection {
		return &v
	}).(PropertiesNetworkInjectionPtrOutput)
}

func (o PropertiesNetworkInjectionOutput) VirtualNetworks() VirtualNetworkPropertiesListPtrOutput {
	return o.ApplyT(func(v PropertiesNetworkInjection) *VirtualNetworkPropertiesList { return v.VirtualNetworks }).(VirtualNetworkPropertiesListPtrOutput)
}

type PropertiesNetworkInjectionPtrOutput struct{ *pulumi.OutputState }

func (PropertiesNetworkInjectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesNetworkInjection)(nil)).Elem()
}

func (o PropertiesNetworkInjectionPtrOutput) ToPropertiesNetworkInjectionPtrOutput() PropertiesNetworkInjectionPtrOutput {
	return o
}

func (o PropertiesNetworkInjectionPtrOutput) ToPropertiesNetworkInjectionPtrOutputWithContext(ctx context.Context) PropertiesNetworkInjectionPtrOutput {
	return o
}

func (o PropertiesNetworkInjectionPtrOutput) Elem() PropertiesNetworkInjectionOutput {
	return o.ApplyT(func(v *PropertiesNetworkInjection) PropertiesNetworkInjection {
		if v != nil {
			return *v
		}
		var ret PropertiesNetworkInjection
		return ret
	}).(PropertiesNetworkInjectionOutput)
}

func (o PropertiesNetworkInjectionPtrOutput) VirtualNetworks() VirtualNetworkPropertiesListPtrOutput {
	return o.ApplyT(func(v *PropertiesNetworkInjection) *VirtualNetworkPropertiesList {
		if v == nil {
			return nil
		}
		return v.VirtualNetworks
	}).(VirtualNetworkPropertiesListPtrOutput)
}

type PropertiesResponseEncryption struct {
	KeyVault *KeyVaultPropertiesResponse `pulumi:"keyVault"`
	State    *string                     `pulumi:"state"`
}





type PropertiesResponseEncryptionInput interface {
	pulumi.Input

	ToPropertiesResponseEncryptionOutput() PropertiesResponseEncryptionOutput
	ToPropertiesResponseEncryptionOutputWithContext(context.Context) PropertiesResponseEncryptionOutput
}

type PropertiesResponseEncryptionArgs struct {
	KeyVault KeyVaultPropertiesResponsePtrInput `pulumi:"keyVault"`
	State    pulumi.StringPtrInput              `pulumi:"state"`
}

func (PropertiesResponseEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesResponseEncryption)(nil)).Elem()
}

func (i PropertiesResponseEncryptionArgs) ToPropertiesResponseEncryptionOutput() PropertiesResponseEncryptionOutput {
	return i.ToPropertiesResponseEncryptionOutputWithContext(context.Background())
}

func (i PropertiesResponseEncryptionArgs) ToPropertiesResponseEncryptionOutputWithContext(ctx context.Context) PropertiesResponseEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesResponseEncryptionOutput)
}

func (i PropertiesResponseEncryptionArgs) ToPropertiesResponseEncryptionPtrOutput() PropertiesResponseEncryptionPtrOutput {
	return i.ToPropertiesResponseEncryptionPtrOutputWithContext(context.Background())
}

func (i PropertiesResponseEncryptionArgs) ToPropertiesResponseEncryptionPtrOutputWithContext(ctx context.Context) PropertiesResponseEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesResponseEncryptionOutput).ToPropertiesResponseEncryptionPtrOutputWithContext(ctx)
}









type PropertiesResponseEncryptionPtrInput interface {
	pulumi.Input

	ToPropertiesResponseEncryptionPtrOutput() PropertiesResponseEncryptionPtrOutput
	ToPropertiesResponseEncryptionPtrOutputWithContext(context.Context) PropertiesResponseEncryptionPtrOutput
}

type propertiesResponseEncryptionPtrType PropertiesResponseEncryptionArgs

func PropertiesResponseEncryptionPtr(v *PropertiesResponseEncryptionArgs) PropertiesResponseEncryptionPtrInput {
	return (*propertiesResponseEncryptionPtrType)(v)
}

func (*propertiesResponseEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesResponseEncryption)(nil)).Elem()
}

func (i *propertiesResponseEncryptionPtrType) ToPropertiesResponseEncryptionPtrOutput() PropertiesResponseEncryptionPtrOutput {
	return i.ToPropertiesResponseEncryptionPtrOutputWithContext(context.Background())
}

func (i *propertiesResponseEncryptionPtrType) ToPropertiesResponseEncryptionPtrOutputWithContext(ctx context.Context) PropertiesResponseEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesResponseEncryptionPtrOutput)
}

type PropertiesResponseEncryptionOutput struct{ *pulumi.OutputState }

func (PropertiesResponseEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesResponseEncryption)(nil)).Elem()
}

func (o PropertiesResponseEncryptionOutput) ToPropertiesResponseEncryptionOutput() PropertiesResponseEncryptionOutput {
	return o
}

func (o PropertiesResponseEncryptionOutput) ToPropertiesResponseEncryptionOutputWithContext(ctx context.Context) PropertiesResponseEncryptionOutput {
	return o
}

func (o PropertiesResponseEncryptionOutput) ToPropertiesResponseEncryptionPtrOutput() PropertiesResponseEncryptionPtrOutput {
	return o.ToPropertiesResponseEncryptionPtrOutputWithContext(context.Background())
}

func (o PropertiesResponseEncryptionOutput) ToPropertiesResponseEncryptionPtrOutputWithContext(ctx context.Context) PropertiesResponseEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PropertiesResponseEncryption) *PropertiesResponseEncryption {
		return &v
	}).(PropertiesResponseEncryptionPtrOutput)
}

func (o PropertiesResponseEncryptionOutput) KeyVault() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PropertiesResponseEncryption) *KeyVaultPropertiesResponse { return v.KeyVault }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o PropertiesResponseEncryptionOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertiesResponseEncryption) *string { return v.State }).(pulumi.StringPtrOutput)
}

type PropertiesResponseEncryptionPtrOutput struct{ *pulumi.OutputState }

func (PropertiesResponseEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesResponseEncryption)(nil)).Elem()
}

func (o PropertiesResponseEncryptionPtrOutput) ToPropertiesResponseEncryptionPtrOutput() PropertiesResponseEncryptionPtrOutput {
	return o
}

func (o PropertiesResponseEncryptionPtrOutput) ToPropertiesResponseEncryptionPtrOutputWithContext(ctx context.Context) PropertiesResponseEncryptionPtrOutput {
	return o
}

func (o PropertiesResponseEncryptionPtrOutput) Elem() PropertiesResponseEncryptionOutput {
	return o.ApplyT(func(v *PropertiesResponseEncryption) PropertiesResponseEncryption {
		if v != nil {
			return *v
		}
		var ret PropertiesResponseEncryption
		return ret
	}).(PropertiesResponseEncryptionOutput)
}

func (o PropertiesResponseEncryptionPtrOutput) KeyVault() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PropertiesResponseEncryption) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVault
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o PropertiesResponseEncryptionPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PropertiesResponseEncryption) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type PropertiesResponseLockbox struct {
	State *string `pulumi:"state"`
}





type PropertiesResponseLockboxInput interface {
	pulumi.Input

	ToPropertiesResponseLockboxOutput() PropertiesResponseLockboxOutput
	ToPropertiesResponseLockboxOutputWithContext(context.Context) PropertiesResponseLockboxOutput
}

type PropertiesResponseLockboxArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (PropertiesResponseLockboxArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesResponseLockbox)(nil)).Elem()
}

func (i PropertiesResponseLockboxArgs) ToPropertiesResponseLockboxOutput() PropertiesResponseLockboxOutput {
	return i.ToPropertiesResponseLockboxOutputWithContext(context.Background())
}

func (i PropertiesResponseLockboxArgs) ToPropertiesResponseLockboxOutputWithContext(ctx context.Context) PropertiesResponseLockboxOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesResponseLockboxOutput)
}

func (i PropertiesResponseLockboxArgs) ToPropertiesResponseLockboxPtrOutput() PropertiesResponseLockboxPtrOutput {
	return i.ToPropertiesResponseLockboxPtrOutputWithContext(context.Background())
}

func (i PropertiesResponseLockboxArgs) ToPropertiesResponseLockboxPtrOutputWithContext(ctx context.Context) PropertiesResponseLockboxPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesResponseLockboxOutput).ToPropertiesResponseLockboxPtrOutputWithContext(ctx)
}









type PropertiesResponseLockboxPtrInput interface {
	pulumi.Input

	ToPropertiesResponseLockboxPtrOutput() PropertiesResponseLockboxPtrOutput
	ToPropertiesResponseLockboxPtrOutputWithContext(context.Context) PropertiesResponseLockboxPtrOutput
}

type propertiesResponseLockboxPtrType PropertiesResponseLockboxArgs

func PropertiesResponseLockboxPtr(v *PropertiesResponseLockboxArgs) PropertiesResponseLockboxPtrInput {
	return (*propertiesResponseLockboxPtrType)(v)
}

func (*propertiesResponseLockboxPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesResponseLockbox)(nil)).Elem()
}

func (i *propertiesResponseLockboxPtrType) ToPropertiesResponseLockboxPtrOutput() PropertiesResponseLockboxPtrOutput {
	return i.ToPropertiesResponseLockboxPtrOutputWithContext(context.Background())
}

func (i *propertiesResponseLockboxPtrType) ToPropertiesResponseLockboxPtrOutputWithContext(ctx context.Context) PropertiesResponseLockboxPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesResponseLockboxPtrOutput)
}

type PropertiesResponseLockboxOutput struct{ *pulumi.OutputState }

func (PropertiesResponseLockboxOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesResponseLockbox)(nil)).Elem()
}

func (o PropertiesResponseLockboxOutput) ToPropertiesResponseLockboxOutput() PropertiesResponseLockboxOutput {
	return o
}

func (o PropertiesResponseLockboxOutput) ToPropertiesResponseLockboxOutputWithContext(ctx context.Context) PropertiesResponseLockboxOutput {
	return o
}

func (o PropertiesResponseLockboxOutput) ToPropertiesResponseLockboxPtrOutput() PropertiesResponseLockboxPtrOutput {
	return o.ToPropertiesResponseLockboxPtrOutputWithContext(context.Background())
}

func (o PropertiesResponseLockboxOutput) ToPropertiesResponseLockboxPtrOutputWithContext(ctx context.Context) PropertiesResponseLockboxPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PropertiesResponseLockbox) *PropertiesResponseLockbox {
		return &v
	}).(PropertiesResponseLockboxPtrOutput)
}

func (o PropertiesResponseLockboxOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertiesResponseLockbox) *string { return v.State }).(pulumi.StringPtrOutput)
}

type PropertiesResponseLockboxPtrOutput struct{ *pulumi.OutputState }

func (PropertiesResponseLockboxPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesResponseLockbox)(nil)).Elem()
}

func (o PropertiesResponseLockboxPtrOutput) ToPropertiesResponseLockboxPtrOutput() PropertiesResponseLockboxPtrOutput {
	return o
}

func (o PropertiesResponseLockboxPtrOutput) ToPropertiesResponseLockboxPtrOutputWithContext(ctx context.Context) PropertiesResponseLockboxPtrOutput {
	return o
}

func (o PropertiesResponseLockboxPtrOutput) Elem() PropertiesResponseLockboxOutput {
	return o.ApplyT(func(v *PropertiesResponseLockbox) PropertiesResponseLockbox {
		if v != nil {
			return *v
		}
		var ret PropertiesResponseLockbox
		return ret
	}).(PropertiesResponseLockboxOutput)
}

func (o PropertiesResponseLockboxPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PropertiesResponseLockbox) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type PropertiesResponseNetworkInjection struct {
	VirtualNetworks *VirtualNetworkPropertiesListResponse `pulumi:"virtualNetworks"`
}





type PropertiesResponseNetworkInjectionInput interface {
	pulumi.Input

	ToPropertiesResponseNetworkInjectionOutput() PropertiesResponseNetworkInjectionOutput
	ToPropertiesResponseNetworkInjectionOutputWithContext(context.Context) PropertiesResponseNetworkInjectionOutput
}

type PropertiesResponseNetworkInjectionArgs struct {
	VirtualNetworks VirtualNetworkPropertiesListResponsePtrInput `pulumi:"virtualNetworks"`
}

func (PropertiesResponseNetworkInjectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesResponseNetworkInjection)(nil)).Elem()
}

func (i PropertiesResponseNetworkInjectionArgs) ToPropertiesResponseNetworkInjectionOutput() PropertiesResponseNetworkInjectionOutput {
	return i.ToPropertiesResponseNetworkInjectionOutputWithContext(context.Background())
}

func (i PropertiesResponseNetworkInjectionArgs) ToPropertiesResponseNetworkInjectionOutputWithContext(ctx context.Context) PropertiesResponseNetworkInjectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesResponseNetworkInjectionOutput)
}

func (i PropertiesResponseNetworkInjectionArgs) ToPropertiesResponseNetworkInjectionPtrOutput() PropertiesResponseNetworkInjectionPtrOutput {
	return i.ToPropertiesResponseNetworkInjectionPtrOutputWithContext(context.Background())
}

func (i PropertiesResponseNetworkInjectionArgs) ToPropertiesResponseNetworkInjectionPtrOutputWithContext(ctx context.Context) PropertiesResponseNetworkInjectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesResponseNetworkInjectionOutput).ToPropertiesResponseNetworkInjectionPtrOutputWithContext(ctx)
}









type PropertiesResponseNetworkInjectionPtrInput interface {
	pulumi.Input

	ToPropertiesResponseNetworkInjectionPtrOutput() PropertiesResponseNetworkInjectionPtrOutput
	ToPropertiesResponseNetworkInjectionPtrOutputWithContext(context.Context) PropertiesResponseNetworkInjectionPtrOutput
}

type propertiesResponseNetworkInjectionPtrType PropertiesResponseNetworkInjectionArgs

func PropertiesResponseNetworkInjectionPtr(v *PropertiesResponseNetworkInjectionArgs) PropertiesResponseNetworkInjectionPtrInput {
	return (*propertiesResponseNetworkInjectionPtrType)(v)
}

func (*propertiesResponseNetworkInjectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesResponseNetworkInjection)(nil)).Elem()
}

func (i *propertiesResponseNetworkInjectionPtrType) ToPropertiesResponseNetworkInjectionPtrOutput() PropertiesResponseNetworkInjectionPtrOutput {
	return i.ToPropertiesResponseNetworkInjectionPtrOutputWithContext(context.Background())
}

func (i *propertiesResponseNetworkInjectionPtrType) ToPropertiesResponseNetworkInjectionPtrOutputWithContext(ctx context.Context) PropertiesResponseNetworkInjectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertiesResponseNetworkInjectionPtrOutput)
}

type PropertiesResponseNetworkInjectionOutput struct{ *pulumi.OutputState }

func (PropertiesResponseNetworkInjectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertiesResponseNetworkInjection)(nil)).Elem()
}

func (o PropertiesResponseNetworkInjectionOutput) ToPropertiesResponseNetworkInjectionOutput() PropertiesResponseNetworkInjectionOutput {
	return o
}

func (o PropertiesResponseNetworkInjectionOutput) ToPropertiesResponseNetworkInjectionOutputWithContext(ctx context.Context) PropertiesResponseNetworkInjectionOutput {
	return o
}

func (o PropertiesResponseNetworkInjectionOutput) ToPropertiesResponseNetworkInjectionPtrOutput() PropertiesResponseNetworkInjectionPtrOutput {
	return o.ToPropertiesResponseNetworkInjectionPtrOutputWithContext(context.Background())
}

func (o PropertiesResponseNetworkInjectionOutput) ToPropertiesResponseNetworkInjectionPtrOutputWithContext(ctx context.Context) PropertiesResponseNetworkInjectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PropertiesResponseNetworkInjection) *PropertiesResponseNetworkInjection {
		return &v
	}).(PropertiesResponseNetworkInjectionPtrOutput)
}

func (o PropertiesResponseNetworkInjectionOutput) VirtualNetworks() VirtualNetworkPropertiesListResponsePtrOutput {
	return o.ApplyT(func(v PropertiesResponseNetworkInjection) *VirtualNetworkPropertiesListResponse {
		return v.VirtualNetworks
	}).(VirtualNetworkPropertiesListResponsePtrOutput)
}

type PropertiesResponseNetworkInjectionPtrOutput struct{ *pulumi.OutputState }

func (PropertiesResponseNetworkInjectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertiesResponseNetworkInjection)(nil)).Elem()
}

func (o PropertiesResponseNetworkInjectionPtrOutput) ToPropertiesResponseNetworkInjectionPtrOutput() PropertiesResponseNetworkInjectionPtrOutput {
	return o
}

func (o PropertiesResponseNetworkInjectionPtrOutput) ToPropertiesResponseNetworkInjectionPtrOutputWithContext(ctx context.Context) PropertiesResponseNetworkInjectionPtrOutput {
	return o
}

func (o PropertiesResponseNetworkInjectionPtrOutput) Elem() PropertiesResponseNetworkInjectionOutput {
	return o.ApplyT(func(v *PropertiesResponseNetworkInjection) PropertiesResponseNetworkInjection {
		if v != nil {
			return *v
		}
		var ret PropertiesResponseNetworkInjection
		return ret
	}).(PropertiesResponseNetworkInjectionOutput)
}

func (o PropertiesResponseNetworkInjectionPtrOutput) VirtualNetworks() VirtualNetworkPropertiesListResponsePtrOutput {
	return o.ApplyT(func(v *PropertiesResponseNetworkInjection) *VirtualNetworkPropertiesListResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworks
	}).(VirtualNetworkPropertiesListResponsePtrOutput)
}

type SubnetProperties struct {
	Name *string `pulumi:"name"`
}





type SubnetPropertiesInput interface {
	pulumi.Input

	ToSubnetPropertiesOutput() SubnetPropertiesOutput
	ToSubnetPropertiesOutputWithContext(context.Context) SubnetPropertiesOutput
}

type SubnetPropertiesArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SubnetPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetProperties)(nil)).Elem()
}

func (i SubnetPropertiesArgs) ToSubnetPropertiesOutput() SubnetPropertiesOutput {
	return i.ToSubnetPropertiesOutputWithContext(context.Background())
}

func (i SubnetPropertiesArgs) ToSubnetPropertiesOutputWithContext(ctx context.Context) SubnetPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPropertiesOutput)
}

func (i SubnetPropertiesArgs) ToSubnetPropertiesPtrOutput() SubnetPropertiesPtrOutput {
	return i.ToSubnetPropertiesPtrOutputWithContext(context.Background())
}

func (i SubnetPropertiesArgs) ToSubnetPropertiesPtrOutputWithContext(ctx context.Context) SubnetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPropertiesOutput).ToSubnetPropertiesPtrOutputWithContext(ctx)
}









type SubnetPropertiesPtrInput interface {
	pulumi.Input

	ToSubnetPropertiesPtrOutput() SubnetPropertiesPtrOutput
	ToSubnetPropertiesPtrOutputWithContext(context.Context) SubnetPropertiesPtrOutput
}

type subnetPropertiesPtrType SubnetPropertiesArgs

func SubnetPropertiesPtr(v *SubnetPropertiesArgs) SubnetPropertiesPtrInput {
	return (*subnetPropertiesPtrType)(v)
}

func (*subnetPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetProperties)(nil)).Elem()
}

func (i *subnetPropertiesPtrType) ToSubnetPropertiesPtrOutput() SubnetPropertiesPtrOutput {
	return i.ToSubnetPropertiesPtrOutputWithContext(context.Background())
}

func (i *subnetPropertiesPtrType) ToSubnetPropertiesPtrOutputWithContext(ctx context.Context) SubnetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPropertiesPtrOutput)
}

type SubnetPropertiesOutput struct{ *pulumi.OutputState }

func (SubnetPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetProperties)(nil)).Elem()
}

func (o SubnetPropertiesOutput) ToSubnetPropertiesOutput() SubnetPropertiesOutput {
	return o
}

func (o SubnetPropertiesOutput) ToSubnetPropertiesOutputWithContext(ctx context.Context) SubnetPropertiesOutput {
	return o
}

func (o SubnetPropertiesOutput) ToSubnetPropertiesPtrOutput() SubnetPropertiesPtrOutput {
	return o.ToSubnetPropertiesPtrOutputWithContext(context.Background())
}

func (o SubnetPropertiesOutput) ToSubnetPropertiesPtrOutputWithContext(ctx context.Context) SubnetPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetProperties) *SubnetProperties {
		return &v
	}).(SubnetPropertiesPtrOutput)
}

func (o SubnetPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SubnetPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SubnetPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetProperties)(nil)).Elem()
}

func (o SubnetPropertiesPtrOutput) ToSubnetPropertiesPtrOutput() SubnetPropertiesPtrOutput {
	return o
}

func (o SubnetPropertiesPtrOutput) ToSubnetPropertiesPtrOutputWithContext(ctx context.Context) SubnetPropertiesPtrOutput {
	return o
}

func (o SubnetPropertiesPtrOutput) Elem() SubnetPropertiesOutput {
	return o.ApplyT(func(v *SubnetProperties) SubnetProperties {
		if v != nil {
			return *v
		}
		var ret SubnetProperties
		return ret
	}).(SubnetPropertiesOutput)
}

func (o SubnetPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SubnetPropertiesResponse struct {
	Name *string `pulumi:"name"`
}





type SubnetPropertiesResponseInput interface {
	pulumi.Input

	ToSubnetPropertiesResponseOutput() SubnetPropertiesResponseOutput
	ToSubnetPropertiesResponseOutputWithContext(context.Context) SubnetPropertiesResponseOutput
}

type SubnetPropertiesResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SubnetPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetPropertiesResponse)(nil)).Elem()
}

func (i SubnetPropertiesResponseArgs) ToSubnetPropertiesResponseOutput() SubnetPropertiesResponseOutput {
	return i.ToSubnetPropertiesResponseOutputWithContext(context.Background())
}

func (i SubnetPropertiesResponseArgs) ToSubnetPropertiesResponseOutputWithContext(ctx context.Context) SubnetPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPropertiesResponseOutput)
}

func (i SubnetPropertiesResponseArgs) ToSubnetPropertiesResponsePtrOutput() SubnetPropertiesResponsePtrOutput {
	return i.ToSubnetPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SubnetPropertiesResponseArgs) ToSubnetPropertiesResponsePtrOutputWithContext(ctx context.Context) SubnetPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPropertiesResponseOutput).ToSubnetPropertiesResponsePtrOutputWithContext(ctx)
}









type SubnetPropertiesResponsePtrInput interface {
	pulumi.Input

	ToSubnetPropertiesResponsePtrOutput() SubnetPropertiesResponsePtrOutput
	ToSubnetPropertiesResponsePtrOutputWithContext(context.Context) SubnetPropertiesResponsePtrOutput
}

type subnetPropertiesResponsePtrType SubnetPropertiesResponseArgs

func SubnetPropertiesResponsePtr(v *SubnetPropertiesResponseArgs) SubnetPropertiesResponsePtrInput {
	return (*subnetPropertiesResponsePtrType)(v)
}

func (*subnetPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetPropertiesResponse)(nil)).Elem()
}

func (i *subnetPropertiesResponsePtrType) ToSubnetPropertiesResponsePtrOutput() SubnetPropertiesResponsePtrOutput {
	return i.ToSubnetPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *subnetPropertiesResponsePtrType) ToSubnetPropertiesResponsePtrOutputWithContext(ctx context.Context) SubnetPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPropertiesResponsePtrOutput)
}

type SubnetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SubnetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetPropertiesResponse)(nil)).Elem()
}

func (o SubnetPropertiesResponseOutput) ToSubnetPropertiesResponseOutput() SubnetPropertiesResponseOutput {
	return o
}

func (o SubnetPropertiesResponseOutput) ToSubnetPropertiesResponseOutputWithContext(ctx context.Context) SubnetPropertiesResponseOutput {
	return o
}

func (o SubnetPropertiesResponseOutput) ToSubnetPropertiesResponsePtrOutput() SubnetPropertiesResponsePtrOutput {
	return o.ToSubnetPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SubnetPropertiesResponseOutput) ToSubnetPropertiesResponsePtrOutputWithContext(ctx context.Context) SubnetPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetPropertiesResponse) *SubnetPropertiesResponse {
		return &v
	}).(SubnetPropertiesResponsePtrOutput)
}

func (o SubnetPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SubnetPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetPropertiesResponse)(nil)).Elem()
}

func (o SubnetPropertiesResponsePtrOutput) ToSubnetPropertiesResponsePtrOutput() SubnetPropertiesResponsePtrOutput {
	return o
}

func (o SubnetPropertiesResponsePtrOutput) ToSubnetPropertiesResponsePtrOutputWithContext(ctx context.Context) SubnetPropertiesResponsePtrOutput {
	return o
}

func (o SubnetPropertiesResponsePtrOutput) Elem() SubnetPropertiesResponseOutput {
	return o.ApplyT(func(v *SubnetPropertiesResponse) SubnetPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SubnetPropertiesResponse
		return ret
	}).(SubnetPropertiesResponseOutput)
}

func (o SubnetPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
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

type VirtualNetworkProperties struct {
	Id     *string           `pulumi:"id"`
	Subnet *SubnetProperties `pulumi:"subnet"`
}





type VirtualNetworkPropertiesInput interface {
	pulumi.Input

	ToVirtualNetworkPropertiesOutput() VirtualNetworkPropertiesOutput
	ToVirtualNetworkPropertiesOutputWithContext(context.Context) VirtualNetworkPropertiesOutput
}

type VirtualNetworkPropertiesArgs struct {
	Id     pulumi.StringPtrInput    `pulumi:"id"`
	Subnet SubnetPropertiesPtrInput `pulumi:"subnet"`
}

func (VirtualNetworkPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProperties)(nil)).Elem()
}

func (i VirtualNetworkPropertiesArgs) ToVirtualNetworkPropertiesOutput() VirtualNetworkPropertiesOutput {
	return i.ToVirtualNetworkPropertiesOutputWithContext(context.Background())
}

func (i VirtualNetworkPropertiesArgs) ToVirtualNetworkPropertiesOutputWithContext(ctx context.Context) VirtualNetworkPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesOutput)
}





type VirtualNetworkPropertiesArrayInput interface {
	pulumi.Input

	ToVirtualNetworkPropertiesArrayOutput() VirtualNetworkPropertiesArrayOutput
	ToVirtualNetworkPropertiesArrayOutputWithContext(context.Context) VirtualNetworkPropertiesArrayOutput
}

type VirtualNetworkPropertiesArray []VirtualNetworkPropertiesInput

func (VirtualNetworkPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkProperties)(nil)).Elem()
}

func (i VirtualNetworkPropertiesArray) ToVirtualNetworkPropertiesArrayOutput() VirtualNetworkPropertiesArrayOutput {
	return i.ToVirtualNetworkPropertiesArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkPropertiesArray) ToVirtualNetworkPropertiesArrayOutputWithContext(ctx context.Context) VirtualNetworkPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesArrayOutput)
}

type VirtualNetworkPropertiesOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProperties)(nil)).Elem()
}

func (o VirtualNetworkPropertiesOutput) ToVirtualNetworkPropertiesOutput() VirtualNetworkPropertiesOutput {
	return o
}

func (o VirtualNetworkPropertiesOutput) ToVirtualNetworkPropertiesOutputWithContext(ctx context.Context) VirtualNetworkPropertiesOutput {
	return o
}

func (o VirtualNetworkPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPropertiesOutput) Subnet() SubnetPropertiesPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProperties) *SubnetProperties { return v.Subnet }).(SubnetPropertiesPtrOutput)
}

type VirtualNetworkPropertiesArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkProperties)(nil)).Elem()
}

func (o VirtualNetworkPropertiesArrayOutput) ToVirtualNetworkPropertiesArrayOutput() VirtualNetworkPropertiesArrayOutput {
	return o
}

func (o VirtualNetworkPropertiesArrayOutput) ToVirtualNetworkPropertiesArrayOutputWithContext(ctx context.Context) VirtualNetworkPropertiesArrayOutput {
	return o
}

func (o VirtualNetworkPropertiesArrayOutput) Index(i pulumi.IntInput) VirtualNetworkPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkProperties {
		return vs[0].([]VirtualNetworkProperties)[vs[1].(int)]
	}).(VirtualNetworkPropertiesOutput)
}

type VirtualNetworkPropertiesList struct {
	NextLink *string                    `pulumi:"nextLink"`
	Value    []VirtualNetworkProperties `pulumi:"value"`
}





type VirtualNetworkPropertiesListInput interface {
	pulumi.Input

	ToVirtualNetworkPropertiesListOutput() VirtualNetworkPropertiesListOutput
	ToVirtualNetworkPropertiesListOutputWithContext(context.Context) VirtualNetworkPropertiesListOutput
}

type VirtualNetworkPropertiesListArgs struct {
	NextLink pulumi.StringPtrInput              `pulumi:"nextLink"`
	Value    VirtualNetworkPropertiesArrayInput `pulumi:"value"`
}

func (VirtualNetworkPropertiesListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPropertiesList)(nil)).Elem()
}

func (i VirtualNetworkPropertiesListArgs) ToVirtualNetworkPropertiesListOutput() VirtualNetworkPropertiesListOutput {
	return i.ToVirtualNetworkPropertiesListOutputWithContext(context.Background())
}

func (i VirtualNetworkPropertiesListArgs) ToVirtualNetworkPropertiesListOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesListOutput)
}

func (i VirtualNetworkPropertiesListArgs) ToVirtualNetworkPropertiesListPtrOutput() VirtualNetworkPropertiesListPtrOutput {
	return i.ToVirtualNetworkPropertiesListPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkPropertiesListArgs) ToVirtualNetworkPropertiesListPtrOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesListOutput).ToVirtualNetworkPropertiesListPtrOutputWithContext(ctx)
}









type VirtualNetworkPropertiesListPtrInput interface {
	pulumi.Input

	ToVirtualNetworkPropertiesListPtrOutput() VirtualNetworkPropertiesListPtrOutput
	ToVirtualNetworkPropertiesListPtrOutputWithContext(context.Context) VirtualNetworkPropertiesListPtrOutput
}

type virtualNetworkPropertiesListPtrType VirtualNetworkPropertiesListArgs

func VirtualNetworkPropertiesListPtr(v *VirtualNetworkPropertiesListArgs) VirtualNetworkPropertiesListPtrInput {
	return (*virtualNetworkPropertiesListPtrType)(v)
}

func (*virtualNetworkPropertiesListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPropertiesList)(nil)).Elem()
}

func (i *virtualNetworkPropertiesListPtrType) ToVirtualNetworkPropertiesListPtrOutput() VirtualNetworkPropertiesListPtrOutput {
	return i.ToVirtualNetworkPropertiesListPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkPropertiesListPtrType) ToVirtualNetworkPropertiesListPtrOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesListPtrOutput)
}

type VirtualNetworkPropertiesListOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPropertiesListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPropertiesList)(nil)).Elem()
}

func (o VirtualNetworkPropertiesListOutput) ToVirtualNetworkPropertiesListOutput() VirtualNetworkPropertiesListOutput {
	return o
}

func (o VirtualNetworkPropertiesListOutput) ToVirtualNetworkPropertiesListOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListOutput {
	return o
}

func (o VirtualNetworkPropertiesListOutput) ToVirtualNetworkPropertiesListPtrOutput() VirtualNetworkPropertiesListPtrOutput {
	return o.ToVirtualNetworkPropertiesListPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPropertiesListOutput) ToVirtualNetworkPropertiesListPtrOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPropertiesList) *VirtualNetworkPropertiesList {
		return &v
	}).(VirtualNetworkPropertiesListPtrOutput)
}

func (o VirtualNetworkPropertiesListOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPropertiesList) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPropertiesListOutput) Value() VirtualNetworkPropertiesArrayOutput {
	return o.ApplyT(func(v VirtualNetworkPropertiesList) []VirtualNetworkProperties { return v.Value }).(VirtualNetworkPropertiesArrayOutput)
}

type VirtualNetworkPropertiesListPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPropertiesListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPropertiesList)(nil)).Elem()
}

func (o VirtualNetworkPropertiesListPtrOutput) ToVirtualNetworkPropertiesListPtrOutput() VirtualNetworkPropertiesListPtrOutput {
	return o
}

func (o VirtualNetworkPropertiesListPtrOutput) ToVirtualNetworkPropertiesListPtrOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListPtrOutput {
	return o
}

func (o VirtualNetworkPropertiesListPtrOutput) Elem() VirtualNetworkPropertiesListOutput {
	return o.ApplyT(func(v *VirtualNetworkPropertiesList) VirtualNetworkPropertiesList {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPropertiesList
		return ret
	}).(VirtualNetworkPropertiesListOutput)
}

func (o VirtualNetworkPropertiesListPtrOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPropertiesList) *string {
		if v == nil {
			return nil
		}
		return v.NextLink
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPropertiesListPtrOutput) Value() VirtualNetworkPropertiesArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkPropertiesList) []VirtualNetworkProperties {
		if v == nil {
			return nil
		}
		return v.Value
	}).(VirtualNetworkPropertiesArrayOutput)
}

type VirtualNetworkPropertiesListResponse struct {
	NextLink *string                            `pulumi:"nextLink"`
	Value    []VirtualNetworkPropertiesResponse `pulumi:"value"`
}





type VirtualNetworkPropertiesListResponseInput interface {
	pulumi.Input

	ToVirtualNetworkPropertiesListResponseOutput() VirtualNetworkPropertiesListResponseOutput
	ToVirtualNetworkPropertiesListResponseOutputWithContext(context.Context) VirtualNetworkPropertiesListResponseOutput
}

type VirtualNetworkPropertiesListResponseArgs struct {
	NextLink pulumi.StringPtrInput                      `pulumi:"nextLink"`
	Value    VirtualNetworkPropertiesResponseArrayInput `pulumi:"value"`
}

func (VirtualNetworkPropertiesListResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPropertiesListResponse)(nil)).Elem()
}

func (i VirtualNetworkPropertiesListResponseArgs) ToVirtualNetworkPropertiesListResponseOutput() VirtualNetworkPropertiesListResponseOutput {
	return i.ToVirtualNetworkPropertiesListResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkPropertiesListResponseArgs) ToVirtualNetworkPropertiesListResponseOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesListResponseOutput)
}

func (i VirtualNetworkPropertiesListResponseArgs) ToVirtualNetworkPropertiesListResponsePtrOutput() VirtualNetworkPropertiesListResponsePtrOutput {
	return i.ToVirtualNetworkPropertiesListResponsePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkPropertiesListResponseArgs) ToVirtualNetworkPropertiesListResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesListResponseOutput).ToVirtualNetworkPropertiesListResponsePtrOutputWithContext(ctx)
}









type VirtualNetworkPropertiesListResponsePtrInput interface {
	pulumi.Input

	ToVirtualNetworkPropertiesListResponsePtrOutput() VirtualNetworkPropertiesListResponsePtrOutput
	ToVirtualNetworkPropertiesListResponsePtrOutputWithContext(context.Context) VirtualNetworkPropertiesListResponsePtrOutput
}

type virtualNetworkPropertiesListResponsePtrType VirtualNetworkPropertiesListResponseArgs

func VirtualNetworkPropertiesListResponsePtr(v *VirtualNetworkPropertiesListResponseArgs) VirtualNetworkPropertiesListResponsePtrInput {
	return (*virtualNetworkPropertiesListResponsePtrType)(v)
}

func (*virtualNetworkPropertiesListResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPropertiesListResponse)(nil)).Elem()
}

func (i *virtualNetworkPropertiesListResponsePtrType) ToVirtualNetworkPropertiesListResponsePtrOutput() VirtualNetworkPropertiesListResponsePtrOutput {
	return i.ToVirtualNetworkPropertiesListResponsePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkPropertiesListResponsePtrType) ToVirtualNetworkPropertiesListResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesListResponsePtrOutput)
}

type VirtualNetworkPropertiesListResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPropertiesListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPropertiesListResponse)(nil)).Elem()
}

func (o VirtualNetworkPropertiesListResponseOutput) ToVirtualNetworkPropertiesListResponseOutput() VirtualNetworkPropertiesListResponseOutput {
	return o
}

func (o VirtualNetworkPropertiesListResponseOutput) ToVirtualNetworkPropertiesListResponseOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListResponseOutput {
	return o
}

func (o VirtualNetworkPropertiesListResponseOutput) ToVirtualNetworkPropertiesListResponsePtrOutput() VirtualNetworkPropertiesListResponsePtrOutput {
	return o.ToVirtualNetworkPropertiesListResponsePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkPropertiesListResponseOutput) ToVirtualNetworkPropertiesListResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkPropertiesListResponse) *VirtualNetworkPropertiesListResponse {
		return &v
	}).(VirtualNetworkPropertiesListResponsePtrOutput)
}

func (o VirtualNetworkPropertiesListResponseOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPropertiesListResponse) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPropertiesListResponseOutput) Value() VirtualNetworkPropertiesResponseArrayOutput {
	return o.ApplyT(func(v VirtualNetworkPropertiesListResponse) []VirtualNetworkPropertiesResponse { return v.Value }).(VirtualNetworkPropertiesResponseArrayOutput)
}

type VirtualNetworkPropertiesListResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPropertiesListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPropertiesListResponse)(nil)).Elem()
}

func (o VirtualNetworkPropertiesListResponsePtrOutput) ToVirtualNetworkPropertiesListResponsePtrOutput() VirtualNetworkPropertiesListResponsePtrOutput {
	return o
}

func (o VirtualNetworkPropertiesListResponsePtrOutput) ToVirtualNetworkPropertiesListResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkPropertiesListResponsePtrOutput {
	return o
}

func (o VirtualNetworkPropertiesListResponsePtrOutput) Elem() VirtualNetworkPropertiesListResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkPropertiesListResponse) VirtualNetworkPropertiesListResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkPropertiesListResponse
		return ret
	}).(VirtualNetworkPropertiesListResponseOutput)
}

func (o VirtualNetworkPropertiesListResponsePtrOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPropertiesListResponse) *string {
		if v == nil {
			return nil
		}
		return v.NextLink
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPropertiesListResponsePtrOutput) Value() VirtualNetworkPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkPropertiesListResponse) []VirtualNetworkPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Value
	}).(VirtualNetworkPropertiesResponseArrayOutput)
}

type VirtualNetworkPropertiesResponse struct {
	Id     *string                   `pulumi:"id"`
	Subnet *SubnetPropertiesResponse `pulumi:"subnet"`
}





type VirtualNetworkPropertiesResponseInput interface {
	pulumi.Input

	ToVirtualNetworkPropertiesResponseOutput() VirtualNetworkPropertiesResponseOutput
	ToVirtualNetworkPropertiesResponseOutputWithContext(context.Context) VirtualNetworkPropertiesResponseOutput
}

type VirtualNetworkPropertiesResponseArgs struct {
	Id     pulumi.StringPtrInput            `pulumi:"id"`
	Subnet SubnetPropertiesResponsePtrInput `pulumi:"subnet"`
}

func (VirtualNetworkPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPropertiesResponse)(nil)).Elem()
}

func (i VirtualNetworkPropertiesResponseArgs) ToVirtualNetworkPropertiesResponseOutput() VirtualNetworkPropertiesResponseOutput {
	return i.ToVirtualNetworkPropertiesResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkPropertiesResponseArgs) ToVirtualNetworkPropertiesResponseOutputWithContext(ctx context.Context) VirtualNetworkPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesResponseOutput)
}





type VirtualNetworkPropertiesResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkPropertiesResponseArrayOutput() VirtualNetworkPropertiesResponseArrayOutput
	ToVirtualNetworkPropertiesResponseArrayOutputWithContext(context.Context) VirtualNetworkPropertiesResponseArrayOutput
}

type VirtualNetworkPropertiesResponseArray []VirtualNetworkPropertiesResponseInput

func (VirtualNetworkPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPropertiesResponse)(nil)).Elem()
}

func (i VirtualNetworkPropertiesResponseArray) ToVirtualNetworkPropertiesResponseArrayOutput() VirtualNetworkPropertiesResponseArrayOutput {
	return i.ToVirtualNetworkPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkPropertiesResponseArray) ToVirtualNetworkPropertiesResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPropertiesResponseArrayOutput)
}

type VirtualNetworkPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPropertiesResponse)(nil)).Elem()
}

func (o VirtualNetworkPropertiesResponseOutput) ToVirtualNetworkPropertiesResponseOutput() VirtualNetworkPropertiesResponseOutput {
	return o
}

func (o VirtualNetworkPropertiesResponseOutput) ToVirtualNetworkPropertiesResponseOutputWithContext(ctx context.Context) VirtualNetworkPropertiesResponseOutput {
	return o
}

func (o VirtualNetworkPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPropertiesResponseOutput) Subnet() SubnetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPropertiesResponse) *SubnetPropertiesResponse { return v.Subnet }).(SubnetPropertiesResponsePtrOutput)
}

type VirtualNetworkPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPropertiesResponse)(nil)).Elem()
}

func (o VirtualNetworkPropertiesResponseArrayOutput) ToVirtualNetworkPropertiesResponseArrayOutput() VirtualNetworkPropertiesResponseArrayOutput {
	return o
}

func (o VirtualNetworkPropertiesResponseArrayOutput) ToVirtualNetworkPropertiesResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkPropertiesResponseArrayOutput {
	return o
}

func (o VirtualNetworkPropertiesResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkPropertiesResponse {
		return vs[0].([]VirtualNetworkPropertiesResponse)[vs[1].(int)]
	}).(VirtualNetworkPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(EnterprisePolicyIdentityOutput{})
	pulumi.RegisterOutputType(EnterprisePolicyIdentityPtrOutput{})
	pulumi.RegisterOutputType(EnterprisePolicyIdentityResponseOutput{})
	pulumi.RegisterOutputType(EnterprisePolicyIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyPropertiesOutput{})
	pulumi.RegisterOutputType(KeyPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PropertiesEncryptionOutput{})
	pulumi.RegisterOutputType(PropertiesEncryptionPtrOutput{})
	pulumi.RegisterOutputType(PropertiesLockboxOutput{})
	pulumi.RegisterOutputType(PropertiesLockboxPtrOutput{})
	pulumi.RegisterOutputType(PropertiesNetworkInjectionOutput{})
	pulumi.RegisterOutputType(PropertiesNetworkInjectionPtrOutput{})
	pulumi.RegisterOutputType(PropertiesResponseEncryptionOutput{})
	pulumi.RegisterOutputType(PropertiesResponseEncryptionPtrOutput{})
	pulumi.RegisterOutputType(PropertiesResponseLockboxOutput{})
	pulumi.RegisterOutputType(PropertiesResponseLockboxPtrOutput{})
	pulumi.RegisterOutputType(PropertiesResponseNetworkInjectionOutput{})
	pulumi.RegisterOutputType(PropertiesResponseNetworkInjectionPtrOutput{})
	pulumi.RegisterOutputType(SubnetPropertiesOutput{})
	pulumi.RegisterOutputType(SubnetPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SubnetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SubnetPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPropertiesOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPropertiesArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPropertiesListOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPropertiesListPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPropertiesListResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPropertiesListResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPropertiesResponseArrayOutput{})
}
