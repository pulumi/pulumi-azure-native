


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessPolicyEntry struct {
	ApplicationId *string     `pulumi:"applicationId"`
	ObjectId      string      `pulumi:"objectId"`
	Permissions   Permissions `pulumi:"permissions"`
	TenantId      string      `pulumi:"tenantId"`
}





type AccessPolicyEntryInput interface {
	pulumi.Input

	ToAccessPolicyEntryOutput() AccessPolicyEntryOutput
	ToAccessPolicyEntryOutputWithContext(context.Context) AccessPolicyEntryOutput
}

type AccessPolicyEntryArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	ObjectId      pulumi.StringInput    `pulumi:"objectId"`
	Permissions   PermissionsInput      `pulumi:"permissions"`
	TenantId      pulumi.StringInput    `pulumi:"tenantId"`
}

func (AccessPolicyEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyEntry)(nil)).Elem()
}

func (i AccessPolicyEntryArgs) ToAccessPolicyEntryOutput() AccessPolicyEntryOutput {
	return i.ToAccessPolicyEntryOutputWithContext(context.Background())
}

func (i AccessPolicyEntryArgs) ToAccessPolicyEntryOutputWithContext(ctx context.Context) AccessPolicyEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyEntryOutput)
}





type AccessPolicyEntryArrayInput interface {
	pulumi.Input

	ToAccessPolicyEntryArrayOutput() AccessPolicyEntryArrayOutput
	ToAccessPolicyEntryArrayOutputWithContext(context.Context) AccessPolicyEntryArrayOutput
}

type AccessPolicyEntryArray []AccessPolicyEntryInput

func (AccessPolicyEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyEntry)(nil)).Elem()
}

func (i AccessPolicyEntryArray) ToAccessPolicyEntryArrayOutput() AccessPolicyEntryArrayOutput {
	return i.ToAccessPolicyEntryArrayOutputWithContext(context.Background())
}

func (i AccessPolicyEntryArray) ToAccessPolicyEntryArrayOutputWithContext(ctx context.Context) AccessPolicyEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyEntryArrayOutput)
}

type AccessPolicyEntryOutput struct{ *pulumi.OutputState }

func (AccessPolicyEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyEntry)(nil)).Elem()
}

func (o AccessPolicyEntryOutput) ToAccessPolicyEntryOutput() AccessPolicyEntryOutput {
	return o
}

func (o AccessPolicyEntryOutput) ToAccessPolicyEntryOutputWithContext(ctx context.Context) AccessPolicyEntryOutput {
	return o
}

func (o AccessPolicyEntryOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessPolicyEntry) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o AccessPolicyEntryOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessPolicyEntry) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o AccessPolicyEntryOutput) Permissions() PermissionsOutput {
	return o.ApplyT(func(v AccessPolicyEntry) Permissions { return v.Permissions }).(PermissionsOutput)
}

func (o AccessPolicyEntryOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessPolicyEntry) string { return v.TenantId }).(pulumi.StringOutput)
}

type AccessPolicyEntryArrayOutput struct{ *pulumi.OutputState }

func (AccessPolicyEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyEntry)(nil)).Elem()
}

func (o AccessPolicyEntryArrayOutput) ToAccessPolicyEntryArrayOutput() AccessPolicyEntryArrayOutput {
	return o
}

func (o AccessPolicyEntryArrayOutput) ToAccessPolicyEntryArrayOutputWithContext(ctx context.Context) AccessPolicyEntryArrayOutput {
	return o
}

func (o AccessPolicyEntryArrayOutput) Index(i pulumi.IntInput) AccessPolicyEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessPolicyEntry {
		return vs[0].([]AccessPolicyEntry)[vs[1].(int)]
	}).(AccessPolicyEntryOutput)
}

type AccessPolicyEntryResponse struct {
	ApplicationId *string             `pulumi:"applicationId"`
	ObjectId      string              `pulumi:"objectId"`
	Permissions   PermissionsResponse `pulumi:"permissions"`
	TenantId      string              `pulumi:"tenantId"`
}

type AccessPolicyEntryResponseOutput struct{ *pulumi.OutputState }

func (AccessPolicyEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyEntryResponse)(nil)).Elem()
}

func (o AccessPolicyEntryResponseOutput) ToAccessPolicyEntryResponseOutput() AccessPolicyEntryResponseOutput {
	return o
}

func (o AccessPolicyEntryResponseOutput) ToAccessPolicyEntryResponseOutputWithContext(ctx context.Context) AccessPolicyEntryResponseOutput {
	return o
}

func (o AccessPolicyEntryResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessPolicyEntryResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o AccessPolicyEntryResponseOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessPolicyEntryResponse) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o AccessPolicyEntryResponseOutput) Permissions() PermissionsResponseOutput {
	return o.ApplyT(func(v AccessPolicyEntryResponse) PermissionsResponse { return v.Permissions }).(PermissionsResponseOutput)
}

func (o AccessPolicyEntryResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessPolicyEntryResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type AccessPolicyEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessPolicyEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyEntryResponse)(nil)).Elem()
}

func (o AccessPolicyEntryResponseArrayOutput) ToAccessPolicyEntryResponseArrayOutput() AccessPolicyEntryResponseArrayOutput {
	return o
}

func (o AccessPolicyEntryResponseArrayOutput) ToAccessPolicyEntryResponseArrayOutputWithContext(ctx context.Context) AccessPolicyEntryResponseArrayOutput {
	return o
}

func (o AccessPolicyEntryResponseArrayOutput) Index(i pulumi.IntInput) AccessPolicyEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessPolicyEntryResponse {
		return vs[0].([]AccessPolicyEntryResponse)[vs[1].(int)]
	}).(AccessPolicyEntryResponseOutput)
}

type Permissions struct {
	Certificates []string `pulumi:"certificates"`
	Keys         []string `pulumi:"keys"`
	Secrets      []string `pulumi:"secrets"`
	Storage      []string `pulumi:"storage"`
}





type PermissionsInput interface {
	pulumi.Input

	ToPermissionsOutput() PermissionsOutput
	ToPermissionsOutputWithContext(context.Context) PermissionsOutput
}

type PermissionsArgs struct {
	Certificates pulumi.StringArrayInput `pulumi:"certificates"`
	Keys         pulumi.StringArrayInput `pulumi:"keys"`
	Secrets      pulumi.StringArrayInput `pulumi:"secrets"`
	Storage      pulumi.StringArrayInput `pulumi:"storage"`
}

func (PermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Permissions)(nil)).Elem()
}

func (i PermissionsArgs) ToPermissionsOutput() PermissionsOutput {
	return i.ToPermissionsOutputWithContext(context.Background())
}

func (i PermissionsArgs) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsOutput)
}

type PermissionsOutput struct{ *pulumi.OutputState }

func (PermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Permissions)(nil)).Elem()
}

func (o PermissionsOutput) ToPermissionsOutput() PermissionsOutput {
	return o
}

func (o PermissionsOutput) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return o
}

func (o PermissionsOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permissions) []string { return v.Certificates }).(pulumi.StringArrayOutput)
}

func (o PermissionsOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permissions) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

func (o PermissionsOutput) Secrets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permissions) []string { return v.Secrets }).(pulumi.StringArrayOutput)
}

func (o PermissionsOutput) Storage() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permissions) []string { return v.Storage }).(pulumi.StringArrayOutput)
}

type PermissionsResponse struct {
	Certificates []string `pulumi:"certificates"`
	Keys         []string `pulumi:"keys"`
	Secrets      []string `pulumi:"secrets"`
	Storage      []string `pulumi:"storage"`
}

type PermissionsResponseOutput struct{ *pulumi.OutputState }

func (PermissionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResponse)(nil)).Elem()
}

func (o PermissionsResponseOutput) ToPermissionsResponseOutput() PermissionsResponseOutput {
	return o
}

func (o PermissionsResponseOutput) ToPermissionsResponseOutputWithContext(ctx context.Context) PermissionsResponseOutput {
	return o
}

func (o PermissionsResponseOutput) Certificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []string { return v.Certificates }).(pulumi.StringArrayOutput)
}

func (o PermissionsResponseOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

func (o PermissionsResponseOutput) Secrets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []string { return v.Secrets }).(pulumi.StringArrayOutput)
}

func (o PermissionsResponseOutput) Storage() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []string { return v.Storage }).(pulumi.StringArrayOutput)
}

type SecretAttributes struct {
	Enabled   *bool `pulumi:"enabled"`
	Expires   *int  `pulumi:"expires"`
	NotBefore *int  `pulumi:"notBefore"`
}





type SecretAttributesInput interface {
	pulumi.Input

	ToSecretAttributesOutput() SecretAttributesOutput
	ToSecretAttributesOutputWithContext(context.Context) SecretAttributesOutput
}

type SecretAttributesArgs struct {
	Enabled   pulumi.BoolPtrInput `pulumi:"enabled"`
	Expires   pulumi.IntPtrInput  `pulumi:"expires"`
	NotBefore pulumi.IntPtrInput  `pulumi:"notBefore"`
}

func (SecretAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretAttributes)(nil)).Elem()
}

func (i SecretAttributesArgs) ToSecretAttributesOutput() SecretAttributesOutput {
	return i.ToSecretAttributesOutputWithContext(context.Background())
}

func (i SecretAttributesArgs) ToSecretAttributesOutputWithContext(ctx context.Context) SecretAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretAttributesOutput)
}

func (i SecretAttributesArgs) ToSecretAttributesPtrOutput() SecretAttributesPtrOutput {
	return i.ToSecretAttributesPtrOutputWithContext(context.Background())
}

func (i SecretAttributesArgs) ToSecretAttributesPtrOutputWithContext(ctx context.Context) SecretAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretAttributesOutput).ToSecretAttributesPtrOutputWithContext(ctx)
}









type SecretAttributesPtrInput interface {
	pulumi.Input

	ToSecretAttributesPtrOutput() SecretAttributesPtrOutput
	ToSecretAttributesPtrOutputWithContext(context.Context) SecretAttributesPtrOutput
}

type secretAttributesPtrType SecretAttributesArgs

func SecretAttributesPtr(v *SecretAttributesArgs) SecretAttributesPtrInput {
	return (*secretAttributesPtrType)(v)
}

func (*secretAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretAttributes)(nil)).Elem()
}

func (i *secretAttributesPtrType) ToSecretAttributesPtrOutput() SecretAttributesPtrOutput {
	return i.ToSecretAttributesPtrOutputWithContext(context.Background())
}

func (i *secretAttributesPtrType) ToSecretAttributesPtrOutputWithContext(ctx context.Context) SecretAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretAttributesPtrOutput)
}

type SecretAttributesOutput struct{ *pulumi.OutputState }

func (SecretAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretAttributes)(nil)).Elem()
}

func (o SecretAttributesOutput) ToSecretAttributesOutput() SecretAttributesOutput {
	return o
}

func (o SecretAttributesOutput) ToSecretAttributesOutputWithContext(ctx context.Context) SecretAttributesOutput {
	return o
}

func (o SecretAttributesOutput) ToSecretAttributesPtrOutput() SecretAttributesPtrOutput {
	return o.ToSecretAttributesPtrOutputWithContext(context.Background())
}

func (o SecretAttributesOutput) ToSecretAttributesPtrOutputWithContext(ctx context.Context) SecretAttributesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretAttributes) *SecretAttributes {
		return &v
	}).(SecretAttributesPtrOutput)
}

func (o SecretAttributesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecretAttributes) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SecretAttributesOutput) Expires() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecretAttributes) *int { return v.Expires }).(pulumi.IntPtrOutput)
}

func (o SecretAttributesOutput) NotBefore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecretAttributes) *int { return v.NotBefore }).(pulumi.IntPtrOutput)
}

type SecretAttributesPtrOutput struct{ *pulumi.OutputState }

func (SecretAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretAttributes)(nil)).Elem()
}

func (o SecretAttributesPtrOutput) ToSecretAttributesPtrOutput() SecretAttributesPtrOutput {
	return o
}

func (o SecretAttributesPtrOutput) ToSecretAttributesPtrOutputWithContext(ctx context.Context) SecretAttributesPtrOutput {
	return o
}

func (o SecretAttributesPtrOutput) Elem() SecretAttributesOutput {
	return o.ApplyT(func(v *SecretAttributes) SecretAttributes {
		if v != nil {
			return *v
		}
		var ret SecretAttributes
		return ret
	}).(SecretAttributesOutput)
}

func (o SecretAttributesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretAttributes) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o SecretAttributesPtrOutput) Expires() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretAttributes) *int {
		if v == nil {
			return nil
		}
		return v.Expires
	}).(pulumi.IntPtrOutput)
}

func (o SecretAttributesPtrOutput) NotBefore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretAttributes) *int {
		if v == nil {
			return nil
		}
		return v.NotBefore
	}).(pulumi.IntPtrOutput)
}

type SecretAttributesResponse struct {
	Created   int   `pulumi:"created"`
	Enabled   *bool `pulumi:"enabled"`
	Expires   *int  `pulumi:"expires"`
	NotBefore *int  `pulumi:"notBefore"`
	Updated   int   `pulumi:"updated"`
}

type SecretAttributesResponseOutput struct{ *pulumi.OutputState }

func (SecretAttributesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretAttributesResponse)(nil)).Elem()
}

func (o SecretAttributesResponseOutput) ToSecretAttributesResponseOutput() SecretAttributesResponseOutput {
	return o
}

func (o SecretAttributesResponseOutput) ToSecretAttributesResponseOutputWithContext(ctx context.Context) SecretAttributesResponseOutput {
	return o
}

func (o SecretAttributesResponseOutput) Created() pulumi.IntOutput {
	return o.ApplyT(func(v SecretAttributesResponse) int { return v.Created }).(pulumi.IntOutput)
}

func (o SecretAttributesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecretAttributesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SecretAttributesResponseOutput) Expires() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecretAttributesResponse) *int { return v.Expires }).(pulumi.IntPtrOutput)
}

func (o SecretAttributesResponseOutput) NotBefore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecretAttributesResponse) *int { return v.NotBefore }).(pulumi.IntPtrOutput)
}

func (o SecretAttributesResponseOutput) Updated() pulumi.IntOutput {
	return o.ApplyT(func(v SecretAttributesResponse) int { return v.Updated }).(pulumi.IntOutput)
}

type SecretAttributesResponsePtrOutput struct{ *pulumi.OutputState }

func (SecretAttributesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretAttributesResponse)(nil)).Elem()
}

func (o SecretAttributesResponsePtrOutput) ToSecretAttributesResponsePtrOutput() SecretAttributesResponsePtrOutput {
	return o
}

func (o SecretAttributesResponsePtrOutput) ToSecretAttributesResponsePtrOutputWithContext(ctx context.Context) SecretAttributesResponsePtrOutput {
	return o
}

func (o SecretAttributesResponsePtrOutput) Elem() SecretAttributesResponseOutput {
	return o.ApplyT(func(v *SecretAttributesResponse) SecretAttributesResponse {
		if v != nil {
			return *v
		}
		var ret SecretAttributesResponse
		return ret
	}).(SecretAttributesResponseOutput)
}

func (o SecretAttributesResponsePtrOutput) Created() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretAttributesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Created
	}).(pulumi.IntPtrOutput)
}

func (o SecretAttributesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretAttributesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o SecretAttributesResponsePtrOutput) Expires() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretAttributesResponse) *int {
		if v == nil {
			return nil
		}
		return v.Expires
	}).(pulumi.IntPtrOutput)
}

func (o SecretAttributesResponsePtrOutput) NotBefore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretAttributesResponse) *int {
		if v == nil {
			return nil
		}
		return v.NotBefore
	}).(pulumi.IntPtrOutput)
}

func (o SecretAttributesResponsePtrOutput) Updated() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretAttributesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Updated
	}).(pulumi.IntPtrOutput)
}

type SecretProperties struct {
	Attributes  *SecretAttributes `pulumi:"attributes"`
	ContentType *string           `pulumi:"contentType"`
	Value       *string           `pulumi:"value"`
}





type SecretPropertiesInput interface {
	pulumi.Input

	ToSecretPropertiesOutput() SecretPropertiesOutput
	ToSecretPropertiesOutputWithContext(context.Context) SecretPropertiesOutput
}

type SecretPropertiesArgs struct {
	Attributes  SecretAttributesPtrInput `pulumi:"attributes"`
	ContentType pulumi.StringPtrInput    `pulumi:"contentType"`
	Value       pulumi.StringPtrInput    `pulumi:"value"`
}

func (SecretPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretProperties)(nil)).Elem()
}

func (i SecretPropertiesArgs) ToSecretPropertiesOutput() SecretPropertiesOutput {
	return i.ToSecretPropertiesOutputWithContext(context.Background())
}

func (i SecretPropertiesArgs) ToSecretPropertiesOutputWithContext(ctx context.Context) SecretPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPropertiesOutput)
}

type SecretPropertiesOutput struct{ *pulumi.OutputState }

func (SecretPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretProperties)(nil)).Elem()
}

func (o SecretPropertiesOutput) ToSecretPropertiesOutput() SecretPropertiesOutput {
	return o
}

func (o SecretPropertiesOutput) ToSecretPropertiesOutputWithContext(ctx context.Context) SecretPropertiesOutput {
	return o
}

func (o SecretPropertiesOutput) Attributes() SecretAttributesPtrOutput {
	return o.ApplyT(func(v SecretProperties) *SecretAttributes { return v.Attributes }).(SecretAttributesPtrOutput)
}

func (o SecretPropertiesOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretProperties) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o SecretPropertiesOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretProperties) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SecretPropertiesResponse struct {
	Attributes           *SecretAttributesResponse `pulumi:"attributes"`
	ContentType          *string                   `pulumi:"contentType"`
	SecretUri            string                    `pulumi:"secretUri"`
	SecretUriWithVersion string                    `pulumi:"secretUriWithVersion"`
	Value                *string                   `pulumi:"value"`
}

type SecretPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SecretPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretPropertiesResponse)(nil)).Elem()
}

func (o SecretPropertiesResponseOutput) ToSecretPropertiesResponseOutput() SecretPropertiesResponseOutput {
	return o
}

func (o SecretPropertiesResponseOutput) ToSecretPropertiesResponseOutputWithContext(ctx context.Context) SecretPropertiesResponseOutput {
	return o
}

func (o SecretPropertiesResponseOutput) Attributes() SecretAttributesResponsePtrOutput {
	return o.ApplyT(func(v SecretPropertiesResponse) *SecretAttributesResponse { return v.Attributes }).(SecretAttributesResponsePtrOutput)
}

func (o SecretPropertiesResponseOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretPropertiesResponse) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o SecretPropertiesResponseOutput) SecretUri() pulumi.StringOutput {
	return o.ApplyT(func(v SecretPropertiesResponse) string { return v.SecretUri }).(pulumi.StringOutput)
}

func (o SecretPropertiesResponseOutput) SecretUriWithVersion() pulumi.StringOutput {
	return o.ApplyT(func(v SecretPropertiesResponse) string { return v.SecretUriWithVersion }).(pulumi.StringOutput)
}

func (o SecretPropertiesResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretPropertiesResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type Sku struct {
	Family string  `pulumi:"family"`
	Name   SkuName `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Family pulumi.StringInput `pulumi:"family"`
	Name   SkuNameInput       `pulumi:"name"`
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

func (o SkuOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuOutput) Name() SkuNameOutput {
	return o.ApplyT(func(v Sku) SkuName { return v.Name }).(SkuNameOutput)
}

type SkuResponse struct {
	Family string `pulumi:"family"`
	Name   string `pulumi:"name"`
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

func (o SkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type VaultProperties struct {
	AccessPolicies               []AccessPolicyEntry `pulumi:"accessPolicies"`
	CreateMode                   *CreateMode         `pulumi:"createMode"`
	EnablePurgeProtection        *bool               `pulumi:"enablePurgeProtection"`
	EnableSoftDelete             *bool               `pulumi:"enableSoftDelete"`
	EnabledForDeployment         *bool               `pulumi:"enabledForDeployment"`
	EnabledForDiskEncryption     *bool               `pulumi:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment *bool               `pulumi:"enabledForTemplateDeployment"`
	Sku                          Sku                 `pulumi:"sku"`
	TenantId                     string              `pulumi:"tenantId"`
	VaultUri                     *string             `pulumi:"vaultUri"`
}





type VaultPropertiesInput interface {
	pulumi.Input

	ToVaultPropertiesOutput() VaultPropertiesOutput
	ToVaultPropertiesOutputWithContext(context.Context) VaultPropertiesOutput
}

type VaultPropertiesArgs struct {
	AccessPolicies               AccessPolicyEntryArrayInput `pulumi:"accessPolicies"`
	CreateMode                   CreateModePtrInput          `pulumi:"createMode"`
	EnablePurgeProtection        pulumi.BoolPtrInput         `pulumi:"enablePurgeProtection"`
	EnableSoftDelete             pulumi.BoolPtrInput         `pulumi:"enableSoftDelete"`
	EnabledForDeployment         pulumi.BoolPtrInput         `pulumi:"enabledForDeployment"`
	EnabledForDiskEncryption     pulumi.BoolPtrInput         `pulumi:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment pulumi.BoolPtrInput         `pulumi:"enabledForTemplateDeployment"`
	Sku                          SkuInput                    `pulumi:"sku"`
	TenantId                     pulumi.StringInput          `pulumi:"tenantId"`
	VaultUri                     pulumi.StringPtrInput       `pulumi:"vaultUri"`
}

func (VaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultProperties)(nil)).Elem()
}

func (i VaultPropertiesArgs) ToVaultPropertiesOutput() VaultPropertiesOutput {
	return i.ToVaultPropertiesOutputWithContext(context.Background())
}

func (i VaultPropertiesArgs) ToVaultPropertiesOutputWithContext(ctx context.Context) VaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesOutput)
}

type VaultPropertiesOutput struct{ *pulumi.OutputState }

func (VaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultProperties)(nil)).Elem()
}

func (o VaultPropertiesOutput) ToVaultPropertiesOutput() VaultPropertiesOutput {
	return o
}

func (o VaultPropertiesOutput) ToVaultPropertiesOutputWithContext(ctx context.Context) VaultPropertiesOutput {
	return o
}

func (o VaultPropertiesOutput) AccessPolicies() AccessPolicyEntryArrayOutput {
	return o.ApplyT(func(v VaultProperties) []AccessPolicyEntry { return v.AccessPolicies }).(AccessPolicyEntryArrayOutput)
}

func (o VaultPropertiesOutput) CreateMode() CreateModePtrOutput {
	return o.ApplyT(func(v VaultProperties) *CreateMode { return v.CreateMode }).(CreateModePtrOutput)
}

func (o VaultPropertiesOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultProperties) *bool { return v.EnablePurgeProtection }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultProperties) *bool { return v.EnableSoftDelete }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesOutput) EnabledForDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultProperties) *bool { return v.EnabledForDeployment }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesOutput) EnabledForDiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultProperties) *bool { return v.EnabledForDiskEncryption }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesOutput) EnabledForTemplateDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultProperties) *bool { return v.EnabledForTemplateDeployment }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesOutput) Sku() SkuOutput {
	return o.ApplyT(func(v VaultProperties) Sku { return v.Sku }).(SkuOutput)
}

func (o VaultPropertiesOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultProperties) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o VaultPropertiesOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultProperties) *string { return v.VaultUri }).(pulumi.StringPtrOutput)
}

type VaultPropertiesResponse struct {
	AccessPolicies               []AccessPolicyEntryResponse `pulumi:"accessPolicies"`
	CreateMode                   *string                     `pulumi:"createMode"`
	EnablePurgeProtection        *bool                       `pulumi:"enablePurgeProtection"`
	EnableSoftDelete             *bool                       `pulumi:"enableSoftDelete"`
	EnabledForDeployment         *bool                       `pulumi:"enabledForDeployment"`
	EnabledForDiskEncryption     *bool                       `pulumi:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment *bool                       `pulumi:"enabledForTemplateDeployment"`
	Sku                          SkuResponse                 `pulumi:"sku"`
	TenantId                     string                      `pulumi:"tenantId"`
	VaultUri                     *string                     `pulumi:"vaultUri"`
}

type VaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponse)(nil)).Elem()
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutputWithContext(ctx context.Context) VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) AccessPolicies() AccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) []AccessPolicyEntryResponse { return v.AccessPolicies }).(AccessPolicyEntryResponseArrayOutput)
}

func (o VaultPropertiesResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnablePurgeProtection }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponseOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnableSoftDelete }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponseOutput) EnabledForDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnabledForDeployment }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponseOutput) EnabledForDiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnabledForDiskEncryption }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponseOutput) EnabledForTemplateDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnabledForTemplateDeployment }).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponseOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o VaultPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *string { return v.VaultUri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessPolicyEntryOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryArrayOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryResponseOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(PermissionsOutput{})
	pulumi.RegisterOutputType(PermissionsResponseOutput{})
	pulumi.RegisterOutputType(SecretAttributesOutput{})
	pulumi.RegisterOutputType(SecretAttributesPtrOutput{})
	pulumi.RegisterOutputType(SecretAttributesResponseOutput{})
	pulumi.RegisterOutputType(SecretAttributesResponsePtrOutput{})
	pulumi.RegisterOutputType(SecretPropertiesOutput{})
	pulumi.RegisterOutputType(SecretPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseOutput{})
}
