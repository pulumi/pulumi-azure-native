


package v20210622

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EncryptionProperties struct {
	Identity           *EncryptionPropertiesIdentity `pulumi:"identity"`
	KeySource          *EncryptionKeySourceType      `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultProperties           `pulumi:"keyVaultProperties"`
}





type EncryptionPropertiesInput interface {
	pulumi.Input

	ToEncryptionPropertiesOutput() EncryptionPropertiesOutput
	ToEncryptionPropertiesOutputWithContext(context.Context) EncryptionPropertiesOutput
}

type EncryptionPropertiesArgs struct {
	Identity           EncryptionPropertiesIdentityPtrInput `pulumi:"identity"`
	KeySource          EncryptionKeySourceTypePtrInput      `pulumi:"keySource"`
	KeyVaultProperties KeyVaultPropertiesPtrInput           `pulumi:"keyVaultProperties"`
}

func (EncryptionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperties)(nil)).Elem()
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesOutput() EncryptionPropertiesOutput {
	return i.ToEncryptionPropertiesOutputWithContext(context.Background())
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesOutputWithContext(ctx context.Context) EncryptionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesOutput)
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return i.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesOutput).ToEncryptionPropertiesPtrOutputWithContext(ctx)
}









type EncryptionPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput
	ToEncryptionPropertiesPtrOutputWithContext(context.Context) EncryptionPropertiesPtrOutput
}

type encryptionPropertiesPtrType EncryptionPropertiesArgs

func EncryptionPropertiesPtr(v *EncryptionPropertiesArgs) EncryptionPropertiesPtrInput {
	return (*encryptionPropertiesPtrType)(v)
}

func (*encryptionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperties)(nil)).Elem()
}

func (i *encryptionPropertiesPtrType) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return i.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertiesPtrType) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesPtrOutput)
}

type EncryptionPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperties)(nil)).Elem()
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesOutput() EncryptionPropertiesOutput {
	return o
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesOutputWithContext(ctx context.Context) EncryptionPropertiesOutput {
	return o
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return o.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionProperties) *EncryptionProperties {
		return &v
	}).(EncryptionPropertiesPtrOutput)
}

func (o EncryptionPropertiesOutput) Identity() EncryptionPropertiesIdentityPtrOutput {
	return o.ApplyT(func(v EncryptionProperties) *EncryptionPropertiesIdentity { return v.Identity }).(EncryptionPropertiesIdentityPtrOutput)
}

func (o EncryptionPropertiesOutput) KeySource() EncryptionKeySourceTypePtrOutput {
	return o.ApplyT(func(v EncryptionProperties) *EncryptionKeySourceType { return v.KeySource }).(EncryptionKeySourceTypePtrOutput)
}

func (o EncryptionPropertiesOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v EncryptionProperties) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

type EncryptionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperties)(nil)).Elem()
}

func (o EncryptionPropertiesPtrOutput) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return o
}

func (o EncryptionPropertiesPtrOutput) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return o
}

func (o EncryptionPropertiesPtrOutput) Elem() EncryptionPropertiesOutput {
	return o.ApplyT(func(v *EncryptionProperties) EncryptionProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionProperties
		return ret
	}).(EncryptionPropertiesOutput)
}

func (o EncryptionPropertiesPtrOutput) Identity() EncryptionPropertiesIdentityPtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *EncryptionPropertiesIdentity {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EncryptionPropertiesIdentityPtrOutput)
}

func (o EncryptionPropertiesPtrOutput) KeySource() EncryptionKeySourceTypePtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *EncryptionKeySourceType {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(EncryptionKeySourceTypePtrOutput)
}

func (o EncryptionPropertiesPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

type EncryptionPropertiesIdentity struct {
	UserAssignedIdentity interface{} `pulumi:"userAssignedIdentity"`
}





type EncryptionPropertiesIdentityInput interface {
	pulumi.Input

	ToEncryptionPropertiesIdentityOutput() EncryptionPropertiesIdentityOutput
	ToEncryptionPropertiesIdentityOutputWithContext(context.Context) EncryptionPropertiesIdentityOutput
}

type EncryptionPropertiesIdentityArgs struct {
	UserAssignedIdentity pulumi.Input `pulumi:"userAssignedIdentity"`
}

func (EncryptionPropertiesIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesIdentity)(nil)).Elem()
}

func (i EncryptionPropertiesIdentityArgs) ToEncryptionPropertiesIdentityOutput() EncryptionPropertiesIdentityOutput {
	return i.ToEncryptionPropertiesIdentityOutputWithContext(context.Background())
}

func (i EncryptionPropertiesIdentityArgs) ToEncryptionPropertiesIdentityOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesIdentityOutput)
}

func (i EncryptionPropertiesIdentityArgs) ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput {
	return i.ToEncryptionPropertiesIdentityPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertiesIdentityArgs) ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesIdentityOutput).ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx)
}









type EncryptionPropertiesIdentityPtrInput interface {
	pulumi.Input

	ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput
	ToEncryptionPropertiesIdentityPtrOutputWithContext(context.Context) EncryptionPropertiesIdentityPtrOutput
}

type encryptionPropertiesIdentityPtrType EncryptionPropertiesIdentityArgs

func EncryptionPropertiesIdentityPtr(v *EncryptionPropertiesIdentityArgs) EncryptionPropertiesIdentityPtrInput {
	return (*encryptionPropertiesIdentityPtrType)(v)
}

func (*encryptionPropertiesIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesIdentity)(nil)).Elem()
}

func (i *encryptionPropertiesIdentityPtrType) ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput {
	return i.ToEncryptionPropertiesIdentityPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertiesIdentityPtrType) ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesIdentityPtrOutput)
}

type EncryptionPropertiesIdentityOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesIdentity)(nil)).Elem()
}

func (o EncryptionPropertiesIdentityOutput) ToEncryptionPropertiesIdentityOutput() EncryptionPropertiesIdentityOutput {
	return o
}

func (o EncryptionPropertiesIdentityOutput) ToEncryptionPropertiesIdentityOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityOutput {
	return o
}

func (o EncryptionPropertiesIdentityOutput) ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput {
	return o.ToEncryptionPropertiesIdentityPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertiesIdentityOutput) ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPropertiesIdentity) *EncryptionPropertiesIdentity {
		return &v
	}).(EncryptionPropertiesIdentityPtrOutput)
}

func (o EncryptionPropertiesIdentityOutput) UserAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v EncryptionPropertiesIdentity) interface{} { return v.UserAssignedIdentity }).(pulumi.AnyOutput)
}

type EncryptionPropertiesIdentityPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesIdentity)(nil)).Elem()
}

func (o EncryptionPropertiesIdentityPtrOutput) ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput {
	return o
}

func (o EncryptionPropertiesIdentityPtrOutput) ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityPtrOutput {
	return o
}

func (o EncryptionPropertiesIdentityPtrOutput) Elem() EncryptionPropertiesIdentityOutput {
	return o.ApplyT(func(v *EncryptionPropertiesIdentity) EncryptionPropertiesIdentity {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesIdentity
		return ret
	}).(EncryptionPropertiesIdentityOutput)
}

func (o EncryptionPropertiesIdentityPtrOutput) UserAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v *EncryptionPropertiesIdentity) interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.AnyOutput)
}

type EncryptionPropertiesResponse struct {
	Identity           *EncryptionPropertiesResponseIdentity `pulumi:"identity"`
	KeySource          *string                               `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultPropertiesResponse           `pulumi:"keyVaultProperties"`
}

type EncryptionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesResponse)(nil)).Elem()
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponseOutput() EncryptionPropertiesResponseOutput {
	return o
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponseOutputWithContext(ctx context.Context) EncryptionPropertiesResponseOutput {
	return o
}

func (o EncryptionPropertiesResponseOutput) Identity() EncryptionPropertiesResponseIdentityPtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) *EncryptionPropertiesResponseIdentity { return v.Identity }).(EncryptionPropertiesResponseIdentityPtrOutput)
}

func (o EncryptionPropertiesResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

type EncryptionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesResponse)(nil)).Elem()
}

func (o EncryptionPropertiesResponsePtrOutput) ToEncryptionPropertiesResponsePtrOutput() EncryptionPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesResponsePtrOutput) ToEncryptionPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesResponsePtrOutput) Elem() EncryptionPropertiesResponseOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) EncryptionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesResponse
		return ret
	}).(EncryptionPropertiesResponseOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) Identity() EncryptionPropertiesResponseIdentityPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *EncryptionPropertiesResponseIdentity {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EncryptionPropertiesResponseIdentityPtrOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

type EncryptionPropertiesResponseIdentity struct {
	UserAssignedIdentity interface{} `pulumi:"userAssignedIdentity"`
}

type EncryptionPropertiesResponseIdentityOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesResponseIdentity)(nil)).Elem()
}

func (o EncryptionPropertiesResponseIdentityOutput) ToEncryptionPropertiesResponseIdentityOutput() EncryptionPropertiesResponseIdentityOutput {
	return o
}

func (o EncryptionPropertiesResponseIdentityOutput) ToEncryptionPropertiesResponseIdentityOutputWithContext(ctx context.Context) EncryptionPropertiesResponseIdentityOutput {
	return o
}

func (o EncryptionPropertiesResponseIdentityOutput) UserAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponseIdentity) interface{} { return v.UserAssignedIdentity }).(pulumi.AnyOutput)
}

type EncryptionPropertiesResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesResponseIdentity)(nil)).Elem()
}

func (o EncryptionPropertiesResponseIdentityPtrOutput) ToEncryptionPropertiesResponseIdentityPtrOutput() EncryptionPropertiesResponseIdentityPtrOutput {
	return o
}

func (o EncryptionPropertiesResponseIdentityPtrOutput) ToEncryptionPropertiesResponseIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesResponseIdentityPtrOutput {
	return o
}

func (o EncryptionPropertiesResponseIdentityPtrOutput) Elem() EncryptionPropertiesResponseIdentityOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponseIdentity) EncryptionPropertiesResponseIdentity {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesResponseIdentity
		return ret
	}).(EncryptionPropertiesResponseIdentityOutput)
}

func (o EncryptionPropertiesResponseIdentityPtrOutput) UserAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponseIdentity) interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.AnyOutput)
}

type HybridRunbookWorkerLegacyResponse struct {
	Ip               *string `pulumi:"ip"`
	LastSeenDateTime *string `pulumi:"lastSeenDateTime"`
	Name             *string `pulumi:"name"`
	RegistrationTime *string `pulumi:"registrationTime"`
}

type HybridRunbookWorkerLegacyResponseOutput struct{ *pulumi.OutputState }

func (HybridRunbookWorkerLegacyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridRunbookWorkerLegacyResponse)(nil)).Elem()
}

func (o HybridRunbookWorkerLegacyResponseOutput) ToHybridRunbookWorkerLegacyResponseOutput() HybridRunbookWorkerLegacyResponseOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseOutput) ToHybridRunbookWorkerLegacyResponseOutputWithContext(ctx context.Context) HybridRunbookWorkerLegacyResponseOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o HybridRunbookWorkerLegacyResponseOutput) LastSeenDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.LastSeenDateTime }).(pulumi.StringPtrOutput)
}

func (o HybridRunbookWorkerLegacyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HybridRunbookWorkerLegacyResponseOutput) RegistrationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.RegistrationTime }).(pulumi.StringPtrOutput)
}

type HybridRunbookWorkerLegacyResponseArrayOutput struct{ *pulumi.OutputState }

func (HybridRunbookWorkerLegacyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HybridRunbookWorkerLegacyResponse)(nil)).Elem()
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) ToHybridRunbookWorkerLegacyResponseArrayOutput() HybridRunbookWorkerLegacyResponseArrayOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) ToHybridRunbookWorkerLegacyResponseArrayOutputWithContext(ctx context.Context) HybridRunbookWorkerLegacyResponseArrayOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) Index(i pulumi.IntInput) HybridRunbookWorkerLegacyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HybridRunbookWorkerLegacyResponse {
		return vs[0].([]HybridRunbookWorkerLegacyResponse)[vs[1].(int)]
	}).(HybridRunbookWorkerLegacyResponseOutput)
}

type Identity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
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

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
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

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                            `pulumi:"principalId"`
	TenantId               string                                            `pulumi:"tenantId"`
	Type                   *string                                           `pulumi:"type"`
	UserAssignedIdentities map[string]IdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
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

func (o IdentityResponseOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
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

func (o IdentityResponsePtrOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
}

type IdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type IdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutput() IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o IdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type IdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutput() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) IdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]IdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(IdentityResponseUserAssignedIdentitiesOutput)
}

type KeyResponse struct {
	KeyName     string `pulumi:"keyName"`
	Permissions string `pulumi:"permissions"`
	Value       string `pulumi:"value"`
}

type KeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVersion  *string `pulumi:"keyVersion"`
	KeyvaultUri *string `pulumi:"keyvaultUri"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
	KeyvaultUri pulumi.StringPtrInput `pulumi:"keyvaultUri"`
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

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyvaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyvaultUri }).(pulumi.StringPtrOutput)
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

func (o KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyvaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyvaultUri
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVersion  *string `pulumi:"keyVersion"`
	KeyvaultUri *string `pulumi:"keyvaultUri"`
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

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyvaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyvaultUri }).(pulumi.StringPtrOutput)
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

func (o KeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyvaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyvaultUri
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                             `pulumi:"id"`
	Name                              string                                             `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	Type                              string                                             `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointPropertyResponse { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
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

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string  `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
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

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
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
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationProperty struct {
	Name *string `pulumi:"name"`
}





type RunAsCredentialAssociationPropertyInput interface {
	pulumi.Input

	ToRunAsCredentialAssociationPropertyOutput() RunAsCredentialAssociationPropertyOutput
	ToRunAsCredentialAssociationPropertyOutputWithContext(context.Context) RunAsCredentialAssociationPropertyOutput
}

type RunAsCredentialAssociationPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (RunAsCredentialAssociationPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyOutput() RunAsCredentialAssociationPropertyOutput {
	return i.ToRunAsCredentialAssociationPropertyOutputWithContext(context.Background())
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsCredentialAssociationPropertyOutput)
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return i.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsCredentialAssociationPropertyOutput).ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx)
}









type RunAsCredentialAssociationPropertyPtrInput interface {
	pulumi.Input

	ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput
	ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Context) RunAsCredentialAssociationPropertyPtrOutput
}

type runAsCredentialAssociationPropertyPtrType RunAsCredentialAssociationPropertyArgs

func RunAsCredentialAssociationPropertyPtr(v *RunAsCredentialAssociationPropertyArgs) RunAsCredentialAssociationPropertyPtrInput {
	return (*runAsCredentialAssociationPropertyPtrType)(v)
}

func (*runAsCredentialAssociationPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (i *runAsCredentialAssociationPropertyPtrType) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return i.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i *runAsCredentialAssociationPropertyPtrType) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsCredentialAssociationPropertyPtrOutput)
}

type RunAsCredentialAssociationPropertyOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyOutput() RunAsCredentialAssociationPropertyOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return o.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Background())
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunAsCredentialAssociationProperty) *RunAsCredentialAssociationProperty {
		return &v
	}).(RunAsCredentialAssociationPropertyPtrOutput)
}

func (o RunAsCredentialAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunAsCredentialAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationPropertyPtrOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyPtrOutput) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyPtrOutput) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyPtrOutput) Elem() RunAsCredentialAssociationPropertyOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationProperty) RunAsCredentialAssociationProperty {
		if v != nil {
			return *v
		}
		var ret RunAsCredentialAssociationProperty
		return ret
	}).(RunAsCredentialAssociationPropertyOutput)
}

func (o RunAsCredentialAssociationPropertyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationProperty) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
}

type RunAsCredentialAssociationPropertyResponseOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsCredentialAssociationPropertyResponse)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyResponseOutput) ToRunAsCredentialAssociationPropertyResponseOutput() RunAsCredentialAssociationPropertyResponseOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponseOutput) ToRunAsCredentialAssociationPropertyResponseOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyResponseOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunAsCredentialAssociationPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunAsCredentialAssociationPropertyResponse)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) ToRunAsCredentialAssociationPropertyResponsePtrOutput() RunAsCredentialAssociationPropertyResponsePtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) ToRunAsCredentialAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyResponsePtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) Elem() RunAsCredentialAssociationPropertyResponseOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationPropertyResponse) RunAsCredentialAssociationPropertyResponse {
		if v != nil {
			return *v
		}
		var ret RunAsCredentialAssociationPropertyResponse
		return ret
	}).(RunAsCredentialAssociationPropertyResponseOutput)
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
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
	pulumi.RegisterOutputType(EncryptionPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesIdentityOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesIdentityPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponseIdentityOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(HybridRunbookWorkerLegacyResponseOutput{})
	pulumi.RegisterOutputType(HybridRunbookWorkerLegacyResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyPtrOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
