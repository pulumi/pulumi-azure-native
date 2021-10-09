


package v20200401preview

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





type AccessPolicyEntryResponseInput interface {
	pulumi.Input

	ToAccessPolicyEntryResponseOutput() AccessPolicyEntryResponseOutput
	ToAccessPolicyEntryResponseOutputWithContext(context.Context) AccessPolicyEntryResponseOutput
}

type AccessPolicyEntryResponseArgs struct {
	ApplicationId pulumi.StringPtrInput    `pulumi:"applicationId"`
	ObjectId      pulumi.StringInput       `pulumi:"objectId"`
	Permissions   PermissionsResponseInput `pulumi:"permissions"`
	TenantId      pulumi.StringInput       `pulumi:"tenantId"`
}

func (AccessPolicyEntryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPolicyEntryResponse)(nil)).Elem()
}

func (i AccessPolicyEntryResponseArgs) ToAccessPolicyEntryResponseOutput() AccessPolicyEntryResponseOutput {
	return i.ToAccessPolicyEntryResponseOutputWithContext(context.Background())
}

func (i AccessPolicyEntryResponseArgs) ToAccessPolicyEntryResponseOutputWithContext(ctx context.Context) AccessPolicyEntryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyEntryResponseOutput)
}





type AccessPolicyEntryResponseArrayInput interface {
	pulumi.Input

	ToAccessPolicyEntryResponseArrayOutput() AccessPolicyEntryResponseArrayOutput
	ToAccessPolicyEntryResponseArrayOutputWithContext(context.Context) AccessPolicyEntryResponseArrayOutput
}

type AccessPolicyEntryResponseArray []AccessPolicyEntryResponseInput

func (AccessPolicyEntryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessPolicyEntryResponse)(nil)).Elem()
}

func (i AccessPolicyEntryResponseArray) ToAccessPolicyEntryResponseArrayOutput() AccessPolicyEntryResponseArrayOutput {
	return i.ToAccessPolicyEntryResponseArrayOutputWithContext(context.Background())
}

func (i AccessPolicyEntryResponseArray) ToAccessPolicyEntryResponseArrayOutputWithContext(ctx context.Context) AccessPolicyEntryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPolicyEntryResponseArrayOutput)
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

type IPRule struct {
	Value string `pulumi:"value"`
}





type IPRuleInput interface {
	pulumi.Input

	ToIPRuleOutput() IPRuleOutput
	ToIPRuleOutputWithContext(context.Context) IPRuleOutput
}

type IPRuleArgs struct {
	Value pulumi.StringInput `pulumi:"value"`
}

func (IPRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (i IPRuleArgs) ToIPRuleOutput() IPRuleOutput {
	return i.ToIPRuleOutputWithContext(context.Background())
}

func (i IPRuleArgs) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleOutput)
}





type IPRuleArrayInput interface {
	pulumi.Input

	ToIPRuleArrayOutput() IPRuleArrayOutput
	ToIPRuleArrayOutputWithContext(context.Context) IPRuleArrayOutput
}

type IPRuleArray []IPRuleInput

func (IPRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (i IPRuleArray) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return i.ToIPRuleArrayOutputWithContext(context.Background())
}

func (i IPRuleArray) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleArrayOutput)
}

type IPRuleOutput struct{ *pulumi.OutputState }

func (IPRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (o IPRuleOutput) ToIPRuleOutput() IPRuleOutput {
	return o
}

func (o IPRuleOutput) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return o
}

func (o IPRuleOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v IPRule) string { return v.Value }).(pulumi.StringOutput)
}

type IPRuleArrayOutput struct{ *pulumi.OutputState }

func (IPRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) Index(i pulumi.IntInput) IPRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRule {
		return vs[0].([]IPRule)[vs[1].(int)]
	}).(IPRuleOutput)
}

type IPRuleResponse struct {
	Value string `pulumi:"value"`
}





type IPRuleResponseInput interface {
	pulumi.Input

	ToIPRuleResponseOutput() IPRuleResponseOutput
	ToIPRuleResponseOutputWithContext(context.Context) IPRuleResponseOutput
}

type IPRuleResponseArgs struct {
	Value pulumi.StringInput `pulumi:"value"`
}

func (IPRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (i IPRuleResponseArgs) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return i.ToIPRuleResponseOutputWithContext(context.Background())
}

func (i IPRuleResponseArgs) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleResponseOutput)
}





type IPRuleResponseArrayInput interface {
	pulumi.Input

	ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput
	ToIPRuleResponseArrayOutputWithContext(context.Context) IPRuleResponseArrayOutput
}

type IPRuleResponseArray []IPRuleResponseInput

func (IPRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (i IPRuleResponseArray) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return i.ToIPRuleResponseArrayOutputWithContext(context.Background())
}

func (i IPRuleResponseArray) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleResponseArrayOutput)
}

type IPRuleResponseOutput struct{ *pulumi.OutputState }

func (IPRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v IPRuleResponse) string { return v.Value }).(pulumi.StringOutput)
}

type IPRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) Index(i pulumi.IntInput) IPRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRuleResponse {
		return vs[0].([]IPRuleResponse)[vs[1].(int)]
	}).(IPRuleResponseOutput)
}

type KeyAttributes struct {
	Enabled   *bool    `pulumi:"enabled"`
	Expires   *float64 `pulumi:"expires"`
	NotBefore *float64 `pulumi:"notBefore"`
}





type KeyAttributesInput interface {
	pulumi.Input

	ToKeyAttributesOutput() KeyAttributesOutput
	ToKeyAttributesOutputWithContext(context.Context) KeyAttributesOutput
}

type KeyAttributesArgs struct {
	Enabled   pulumi.BoolPtrInput    `pulumi:"enabled"`
	Expires   pulumi.Float64PtrInput `pulumi:"expires"`
	NotBefore pulumi.Float64PtrInput `pulumi:"notBefore"`
}

func (KeyAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAttributes)(nil)).Elem()
}

func (i KeyAttributesArgs) ToKeyAttributesOutput() KeyAttributesOutput {
	return i.ToKeyAttributesOutputWithContext(context.Background())
}

func (i KeyAttributesArgs) ToKeyAttributesOutputWithContext(ctx context.Context) KeyAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesOutput)
}

func (i KeyAttributesArgs) ToKeyAttributesPtrOutput() KeyAttributesPtrOutput {
	return i.ToKeyAttributesPtrOutputWithContext(context.Background())
}

func (i KeyAttributesArgs) ToKeyAttributesPtrOutputWithContext(ctx context.Context) KeyAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesOutput).ToKeyAttributesPtrOutputWithContext(ctx)
}









type KeyAttributesPtrInput interface {
	pulumi.Input

	ToKeyAttributesPtrOutput() KeyAttributesPtrOutput
	ToKeyAttributesPtrOutputWithContext(context.Context) KeyAttributesPtrOutput
}

type keyAttributesPtrType KeyAttributesArgs

func KeyAttributesPtr(v *KeyAttributesArgs) KeyAttributesPtrInput {
	return (*keyAttributesPtrType)(v)
}

func (*keyAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyAttributes)(nil)).Elem()
}

func (i *keyAttributesPtrType) ToKeyAttributesPtrOutput() KeyAttributesPtrOutput {
	return i.ToKeyAttributesPtrOutputWithContext(context.Background())
}

func (i *keyAttributesPtrType) ToKeyAttributesPtrOutputWithContext(ctx context.Context) KeyAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesPtrOutput)
}

type KeyAttributesOutput struct{ *pulumi.OutputState }

func (KeyAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAttributes)(nil)).Elem()
}

func (o KeyAttributesOutput) ToKeyAttributesOutput() KeyAttributesOutput {
	return o
}

func (o KeyAttributesOutput) ToKeyAttributesOutputWithContext(ctx context.Context) KeyAttributesOutput {
	return o
}

func (o KeyAttributesOutput) ToKeyAttributesPtrOutput() KeyAttributesPtrOutput {
	return o.ToKeyAttributesPtrOutputWithContext(context.Background())
}

func (o KeyAttributesOutput) ToKeyAttributesPtrOutputWithContext(ctx context.Context) KeyAttributesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyAttributes) *KeyAttributes {
		return &v
	}).(KeyAttributesPtrOutput)
}

func (o KeyAttributesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyAttributes) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o KeyAttributesOutput) Expires() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KeyAttributes) *float64 { return v.Expires }).(pulumi.Float64PtrOutput)
}

func (o KeyAttributesOutput) NotBefore() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KeyAttributes) *float64 { return v.NotBefore }).(pulumi.Float64PtrOutput)
}

type KeyAttributesPtrOutput struct{ *pulumi.OutputState }

func (KeyAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyAttributes)(nil)).Elem()
}

func (o KeyAttributesPtrOutput) ToKeyAttributesPtrOutput() KeyAttributesPtrOutput {
	return o
}

func (o KeyAttributesPtrOutput) ToKeyAttributesPtrOutputWithContext(ctx context.Context) KeyAttributesPtrOutput {
	return o
}

func (o KeyAttributesPtrOutput) Elem() KeyAttributesOutput {
	return o.ApplyT(func(v *KeyAttributes) KeyAttributes {
		if v != nil {
			return *v
		}
		var ret KeyAttributes
		return ret
	}).(KeyAttributesOutput)
}

func (o KeyAttributesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o KeyAttributesPtrOutput) Expires() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *float64 {
		if v == nil {
			return nil
		}
		return v.Expires
	}).(pulumi.Float64PtrOutput)
}

func (o KeyAttributesPtrOutput) NotBefore() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *float64 {
		if v == nil {
			return nil
		}
		return v.NotBefore
	}).(pulumi.Float64PtrOutput)
}

type KeyAttributesResponse struct {
	Created       float64  `pulumi:"created"`
	Enabled       *bool    `pulumi:"enabled"`
	Expires       *float64 `pulumi:"expires"`
	NotBefore     *float64 `pulumi:"notBefore"`
	RecoveryLevel string   `pulumi:"recoveryLevel"`
	Updated       float64  `pulumi:"updated"`
}





type KeyAttributesResponseInput interface {
	pulumi.Input

	ToKeyAttributesResponseOutput() KeyAttributesResponseOutput
	ToKeyAttributesResponseOutputWithContext(context.Context) KeyAttributesResponseOutput
}

type KeyAttributesResponseArgs struct {
	Created       pulumi.Float64Input    `pulumi:"created"`
	Enabled       pulumi.BoolPtrInput    `pulumi:"enabled"`
	Expires       pulumi.Float64PtrInput `pulumi:"expires"`
	NotBefore     pulumi.Float64PtrInput `pulumi:"notBefore"`
	RecoveryLevel pulumi.StringInput     `pulumi:"recoveryLevel"`
	Updated       pulumi.Float64Input    `pulumi:"updated"`
}

func (KeyAttributesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAttributesResponse)(nil)).Elem()
}

func (i KeyAttributesResponseArgs) ToKeyAttributesResponseOutput() KeyAttributesResponseOutput {
	return i.ToKeyAttributesResponseOutputWithContext(context.Background())
}

func (i KeyAttributesResponseArgs) ToKeyAttributesResponseOutputWithContext(ctx context.Context) KeyAttributesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesResponseOutput)
}

func (i KeyAttributesResponseArgs) ToKeyAttributesResponsePtrOutput() KeyAttributesResponsePtrOutput {
	return i.ToKeyAttributesResponsePtrOutputWithContext(context.Background())
}

func (i KeyAttributesResponseArgs) ToKeyAttributesResponsePtrOutputWithContext(ctx context.Context) KeyAttributesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesResponseOutput).ToKeyAttributesResponsePtrOutputWithContext(ctx)
}









type KeyAttributesResponsePtrInput interface {
	pulumi.Input

	ToKeyAttributesResponsePtrOutput() KeyAttributesResponsePtrOutput
	ToKeyAttributesResponsePtrOutputWithContext(context.Context) KeyAttributesResponsePtrOutput
}

type keyAttributesResponsePtrType KeyAttributesResponseArgs

func KeyAttributesResponsePtr(v *KeyAttributesResponseArgs) KeyAttributesResponsePtrInput {
	return (*keyAttributesResponsePtrType)(v)
}

func (*keyAttributesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyAttributesResponse)(nil)).Elem()
}

func (i *keyAttributesResponsePtrType) ToKeyAttributesResponsePtrOutput() KeyAttributesResponsePtrOutput {
	return i.ToKeyAttributesResponsePtrOutputWithContext(context.Background())
}

func (i *keyAttributesResponsePtrType) ToKeyAttributesResponsePtrOutputWithContext(ctx context.Context) KeyAttributesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAttributesResponsePtrOutput)
}

type KeyAttributesResponseOutput struct{ *pulumi.OutputState }

func (KeyAttributesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAttributesResponse)(nil)).Elem()
}

func (o KeyAttributesResponseOutput) ToKeyAttributesResponseOutput() KeyAttributesResponseOutput {
	return o
}

func (o KeyAttributesResponseOutput) ToKeyAttributesResponseOutputWithContext(ctx context.Context) KeyAttributesResponseOutput {
	return o
}

func (o KeyAttributesResponseOutput) ToKeyAttributesResponsePtrOutput() KeyAttributesResponsePtrOutput {
	return o.ToKeyAttributesResponsePtrOutputWithContext(context.Background())
}

func (o KeyAttributesResponseOutput) ToKeyAttributesResponsePtrOutputWithContext(ctx context.Context) KeyAttributesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyAttributesResponse) *KeyAttributesResponse {
		return &v
	}).(KeyAttributesResponsePtrOutput)
}

func (o KeyAttributesResponseOutput) Created() pulumi.Float64Output {
	return o.ApplyT(func(v KeyAttributesResponse) float64 { return v.Created }).(pulumi.Float64Output)
}

func (o KeyAttributesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyAttributesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o KeyAttributesResponseOutput) Expires() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KeyAttributesResponse) *float64 { return v.Expires }).(pulumi.Float64PtrOutput)
}

func (o KeyAttributesResponseOutput) NotBefore() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KeyAttributesResponse) *float64 { return v.NotBefore }).(pulumi.Float64PtrOutput)
}

func (o KeyAttributesResponseOutput) RecoveryLevel() pulumi.StringOutput {
	return o.ApplyT(func(v KeyAttributesResponse) string { return v.RecoveryLevel }).(pulumi.StringOutput)
}

func (o KeyAttributesResponseOutput) Updated() pulumi.Float64Output {
	return o.ApplyT(func(v KeyAttributesResponse) float64 { return v.Updated }).(pulumi.Float64Output)
}

type KeyAttributesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyAttributesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyAttributesResponse)(nil)).Elem()
}

func (o KeyAttributesResponsePtrOutput) ToKeyAttributesResponsePtrOutput() KeyAttributesResponsePtrOutput {
	return o
}

func (o KeyAttributesResponsePtrOutput) ToKeyAttributesResponsePtrOutputWithContext(ctx context.Context) KeyAttributesResponsePtrOutput {
	return o
}

func (o KeyAttributesResponsePtrOutput) Elem() KeyAttributesResponseOutput {
	return o.ApplyT(func(v *KeyAttributesResponse) KeyAttributesResponse {
		if v != nil {
			return *v
		}
		var ret KeyAttributesResponse
		return ret
	}).(KeyAttributesResponseOutput)
}

func (o KeyAttributesResponsePtrOutput) Created() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyAttributesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Created
	}).(pulumi.Float64PtrOutput)
}

func (o KeyAttributesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyAttributesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o KeyAttributesResponsePtrOutput) Expires() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyAttributesResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Expires
	}).(pulumi.Float64PtrOutput)
}

func (o KeyAttributesResponsePtrOutput) NotBefore() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyAttributesResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.NotBefore
	}).(pulumi.Float64PtrOutput)
}

func (o KeyAttributesResponsePtrOutput) RecoveryLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyAttributesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RecoveryLevel
	}).(pulumi.StringPtrOutput)
}

func (o KeyAttributesResponsePtrOutput) Updated() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyAttributesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Updated
	}).(pulumi.Float64PtrOutput)
}

type KeyProperties struct {
	Attributes *KeyAttributes `pulumi:"attributes"`
	CurveName  *string        `pulumi:"curveName"`
	KeyOps     []string       `pulumi:"keyOps"`
	KeySize    *int           `pulumi:"keySize"`
	Kty        *string        `pulumi:"kty"`
}





type KeyPropertiesInput interface {
	pulumi.Input

	ToKeyPropertiesOutput() KeyPropertiesOutput
	ToKeyPropertiesOutputWithContext(context.Context) KeyPropertiesOutput
}

type KeyPropertiesArgs struct {
	Attributes KeyAttributesPtrInput   `pulumi:"attributes"`
	CurveName  pulumi.StringPtrInput   `pulumi:"curveName"`
	KeyOps     pulumi.StringArrayInput `pulumi:"keyOps"`
	KeySize    pulumi.IntPtrInput      `pulumi:"keySize"`
	Kty        pulumi.StringPtrInput   `pulumi:"kty"`
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

func (o KeyPropertiesOutput) Attributes() KeyAttributesPtrOutput {
	return o.ApplyT(func(v KeyProperties) *KeyAttributes { return v.Attributes }).(KeyAttributesPtrOutput)
}

func (o KeyPropertiesOutput) CurveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyProperties) *string { return v.CurveName }).(pulumi.StringPtrOutput)
}

func (o KeyPropertiesOutput) KeyOps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KeyProperties) []string { return v.KeyOps }).(pulumi.StringArrayOutput)
}

func (o KeyPropertiesOutput) KeySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KeyProperties) *int { return v.KeySize }).(pulumi.IntPtrOutput)
}

func (o KeyPropertiesOutput) Kty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyProperties) *string { return v.Kty }).(pulumi.StringPtrOutput)
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

func (o KeyPropertiesPtrOutput) Attributes() KeyAttributesPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *KeyAttributes {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(KeyAttributesPtrOutput)
}

func (o KeyPropertiesPtrOutput) CurveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *string {
		if v == nil {
			return nil
		}
		return v.CurveName
	}).(pulumi.StringPtrOutput)
}

func (o KeyPropertiesPtrOutput) KeyOps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KeyProperties) []string {
		if v == nil {
			return nil
		}
		return v.KeyOps
	}).(pulumi.StringArrayOutput)
}

func (o KeyPropertiesPtrOutput) KeySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *int {
		if v == nil {
			return nil
		}
		return v.KeySize
	}).(pulumi.IntPtrOutput)
}

func (o KeyPropertiesPtrOutput) Kty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *string {
		if v == nil {
			return nil
		}
		return v.Kty
	}).(pulumi.StringPtrOutput)
}

type ManagedHsmProperties struct {
	CreateMode                *CreateMode `pulumi:"createMode"`
	EnablePurgeProtection     *bool       `pulumi:"enablePurgeProtection"`
	EnableSoftDelete          *bool       `pulumi:"enableSoftDelete"`
	InitialAdminObjectIds     []string    `pulumi:"initialAdminObjectIds"`
	SoftDeleteRetentionInDays *int        `pulumi:"softDeleteRetentionInDays"`
	TenantId                  *string     `pulumi:"tenantId"`
}





type ManagedHsmPropertiesInput interface {
	pulumi.Input

	ToManagedHsmPropertiesOutput() ManagedHsmPropertiesOutput
	ToManagedHsmPropertiesOutputWithContext(context.Context) ManagedHsmPropertiesOutput
}

type ManagedHsmPropertiesArgs struct {
	CreateMode                CreateModePtrInput      `pulumi:"createMode"`
	EnablePurgeProtection     pulumi.BoolPtrInput     `pulumi:"enablePurgeProtection"`
	EnableSoftDelete          pulumi.BoolPtrInput     `pulumi:"enableSoftDelete"`
	InitialAdminObjectIds     pulumi.StringArrayInput `pulumi:"initialAdminObjectIds"`
	SoftDeleteRetentionInDays pulumi.IntPtrInput      `pulumi:"softDeleteRetentionInDays"`
	TenantId                  pulumi.StringPtrInput   `pulumi:"tenantId"`
}

func (ManagedHsmPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHsmProperties)(nil)).Elem()
}

func (i ManagedHsmPropertiesArgs) ToManagedHsmPropertiesOutput() ManagedHsmPropertiesOutput {
	return i.ToManagedHsmPropertiesOutputWithContext(context.Background())
}

func (i ManagedHsmPropertiesArgs) ToManagedHsmPropertiesOutputWithContext(ctx context.Context) ManagedHsmPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmPropertiesOutput)
}

func (i ManagedHsmPropertiesArgs) ToManagedHsmPropertiesPtrOutput() ManagedHsmPropertiesPtrOutput {
	return i.ToManagedHsmPropertiesPtrOutputWithContext(context.Background())
}

func (i ManagedHsmPropertiesArgs) ToManagedHsmPropertiesPtrOutputWithContext(ctx context.Context) ManagedHsmPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmPropertiesOutput).ToManagedHsmPropertiesPtrOutputWithContext(ctx)
}









type ManagedHsmPropertiesPtrInput interface {
	pulumi.Input

	ToManagedHsmPropertiesPtrOutput() ManagedHsmPropertiesPtrOutput
	ToManagedHsmPropertiesPtrOutputWithContext(context.Context) ManagedHsmPropertiesPtrOutput
}

type managedHsmPropertiesPtrType ManagedHsmPropertiesArgs

func ManagedHsmPropertiesPtr(v *ManagedHsmPropertiesArgs) ManagedHsmPropertiesPtrInput {
	return (*managedHsmPropertiesPtrType)(v)
}

func (*managedHsmPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsmProperties)(nil)).Elem()
}

func (i *managedHsmPropertiesPtrType) ToManagedHsmPropertiesPtrOutput() ManagedHsmPropertiesPtrOutput {
	return i.ToManagedHsmPropertiesPtrOutputWithContext(context.Background())
}

func (i *managedHsmPropertiesPtrType) ToManagedHsmPropertiesPtrOutputWithContext(ctx context.Context) ManagedHsmPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmPropertiesPtrOutput)
}

type ManagedHsmPropertiesOutput struct{ *pulumi.OutputState }

func (ManagedHsmPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHsmProperties)(nil)).Elem()
}

func (o ManagedHsmPropertiesOutput) ToManagedHsmPropertiesOutput() ManagedHsmPropertiesOutput {
	return o
}

func (o ManagedHsmPropertiesOutput) ToManagedHsmPropertiesOutputWithContext(ctx context.Context) ManagedHsmPropertiesOutput {
	return o
}

func (o ManagedHsmPropertiesOutput) ToManagedHsmPropertiesPtrOutput() ManagedHsmPropertiesPtrOutput {
	return o.ToManagedHsmPropertiesPtrOutputWithContext(context.Background())
}

func (o ManagedHsmPropertiesOutput) ToManagedHsmPropertiesPtrOutputWithContext(ctx context.Context) ManagedHsmPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedHsmProperties) *ManagedHsmProperties {
		return &v
	}).(ManagedHsmPropertiesPtrOutput)
}

func (o ManagedHsmPropertiesOutput) CreateMode() CreateModePtrOutput {
	return o.ApplyT(func(v ManagedHsmProperties) *CreateMode { return v.CreateMode }).(CreateModePtrOutput)
}

func (o ManagedHsmPropertiesOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedHsmProperties) *bool { return v.EnablePurgeProtection }).(pulumi.BoolPtrOutput)
}

func (o ManagedHsmPropertiesOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedHsmProperties) *bool { return v.EnableSoftDelete }).(pulumi.BoolPtrOutput)
}

func (o ManagedHsmPropertiesOutput) InitialAdminObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedHsmProperties) []string { return v.InitialAdminObjectIds }).(pulumi.StringArrayOutput)
}

func (o ManagedHsmPropertiesOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedHsmProperties) *int { return v.SoftDeleteRetentionInDays }).(pulumi.IntPtrOutput)
}

func (o ManagedHsmPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedHsmProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ManagedHsmPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagedHsmPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsmProperties)(nil)).Elem()
}

func (o ManagedHsmPropertiesPtrOutput) ToManagedHsmPropertiesPtrOutput() ManagedHsmPropertiesPtrOutput {
	return o
}

func (o ManagedHsmPropertiesPtrOutput) ToManagedHsmPropertiesPtrOutputWithContext(ctx context.Context) ManagedHsmPropertiesPtrOutput {
	return o
}

func (o ManagedHsmPropertiesPtrOutput) Elem() ManagedHsmPropertiesOutput {
	return o.ApplyT(func(v *ManagedHsmProperties) ManagedHsmProperties {
		if v != nil {
			return *v
		}
		var ret ManagedHsmProperties
		return ret
	}).(ManagedHsmPropertiesOutput)
}

func (o ManagedHsmPropertiesPtrOutput) CreateMode() CreateModePtrOutput {
	return o.ApplyT(func(v *ManagedHsmProperties) *CreateMode {
		if v == nil {
			return nil
		}
		return v.CreateMode
	}).(CreateModePtrOutput)
}

func (o ManagedHsmPropertiesPtrOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedHsmProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePurgeProtection
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedHsmPropertiesPtrOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedHsmProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableSoftDelete
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedHsmPropertiesPtrOutput) InitialAdminObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedHsmProperties) []string {
		if v == nil {
			return nil
		}
		return v.InitialAdminObjectIds
	}).(pulumi.StringArrayOutput)
}

func (o ManagedHsmPropertiesPtrOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedHsmProperties) *int {
		if v == nil {
			return nil
		}
		return v.SoftDeleteRetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o ManagedHsmPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ManagedHsmPropertiesResponse struct {
	CreateMode                *string  `pulumi:"createMode"`
	EnablePurgeProtection     *bool    `pulumi:"enablePurgeProtection"`
	EnableSoftDelete          *bool    `pulumi:"enableSoftDelete"`
	HsmUri                    string   `pulumi:"hsmUri"`
	InitialAdminObjectIds     []string `pulumi:"initialAdminObjectIds"`
	ProvisioningState         string   `pulumi:"provisioningState"`
	SoftDeleteRetentionInDays *int     `pulumi:"softDeleteRetentionInDays"`
	StatusMessage             string   `pulumi:"statusMessage"`
	TenantId                  *string  `pulumi:"tenantId"`
}





type ManagedHsmPropertiesResponseInput interface {
	pulumi.Input

	ToManagedHsmPropertiesResponseOutput() ManagedHsmPropertiesResponseOutput
	ToManagedHsmPropertiesResponseOutputWithContext(context.Context) ManagedHsmPropertiesResponseOutput
}

type ManagedHsmPropertiesResponseArgs struct {
	CreateMode                pulumi.StringPtrInput   `pulumi:"createMode"`
	EnablePurgeProtection     pulumi.BoolPtrInput     `pulumi:"enablePurgeProtection"`
	EnableSoftDelete          pulumi.BoolPtrInput     `pulumi:"enableSoftDelete"`
	HsmUri                    pulumi.StringInput      `pulumi:"hsmUri"`
	InitialAdminObjectIds     pulumi.StringArrayInput `pulumi:"initialAdminObjectIds"`
	ProvisioningState         pulumi.StringInput      `pulumi:"provisioningState"`
	SoftDeleteRetentionInDays pulumi.IntPtrInput      `pulumi:"softDeleteRetentionInDays"`
	StatusMessage             pulumi.StringInput      `pulumi:"statusMessage"`
	TenantId                  pulumi.StringPtrInput   `pulumi:"tenantId"`
}

func (ManagedHsmPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHsmPropertiesResponse)(nil)).Elem()
}

func (i ManagedHsmPropertiesResponseArgs) ToManagedHsmPropertiesResponseOutput() ManagedHsmPropertiesResponseOutput {
	return i.ToManagedHsmPropertiesResponseOutputWithContext(context.Background())
}

func (i ManagedHsmPropertiesResponseArgs) ToManagedHsmPropertiesResponseOutputWithContext(ctx context.Context) ManagedHsmPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmPropertiesResponseOutput)
}

func (i ManagedHsmPropertiesResponseArgs) ToManagedHsmPropertiesResponsePtrOutput() ManagedHsmPropertiesResponsePtrOutput {
	return i.ToManagedHsmPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ManagedHsmPropertiesResponseArgs) ToManagedHsmPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedHsmPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmPropertiesResponseOutput).ToManagedHsmPropertiesResponsePtrOutputWithContext(ctx)
}









type ManagedHsmPropertiesResponsePtrInput interface {
	pulumi.Input

	ToManagedHsmPropertiesResponsePtrOutput() ManagedHsmPropertiesResponsePtrOutput
	ToManagedHsmPropertiesResponsePtrOutputWithContext(context.Context) ManagedHsmPropertiesResponsePtrOutput
}

type managedHsmPropertiesResponsePtrType ManagedHsmPropertiesResponseArgs

func ManagedHsmPropertiesResponsePtr(v *ManagedHsmPropertiesResponseArgs) ManagedHsmPropertiesResponsePtrInput {
	return (*managedHsmPropertiesResponsePtrType)(v)
}

func (*managedHsmPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsmPropertiesResponse)(nil)).Elem()
}

func (i *managedHsmPropertiesResponsePtrType) ToManagedHsmPropertiesResponsePtrOutput() ManagedHsmPropertiesResponsePtrOutput {
	return i.ToManagedHsmPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *managedHsmPropertiesResponsePtrType) ToManagedHsmPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedHsmPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmPropertiesResponsePtrOutput)
}

type ManagedHsmPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagedHsmPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHsmPropertiesResponse)(nil)).Elem()
}

func (o ManagedHsmPropertiesResponseOutput) ToManagedHsmPropertiesResponseOutput() ManagedHsmPropertiesResponseOutput {
	return o
}

func (o ManagedHsmPropertiesResponseOutput) ToManagedHsmPropertiesResponseOutputWithContext(ctx context.Context) ManagedHsmPropertiesResponseOutput {
	return o
}

func (o ManagedHsmPropertiesResponseOutput) ToManagedHsmPropertiesResponsePtrOutput() ManagedHsmPropertiesResponsePtrOutput {
	return o.ToManagedHsmPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ManagedHsmPropertiesResponseOutput) ToManagedHsmPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedHsmPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedHsmPropertiesResponse) *ManagedHsmPropertiesResponse {
		return &v
	}).(ManagedHsmPropertiesResponsePtrOutput)
}

func (o ManagedHsmPropertiesResponseOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) *string { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o ManagedHsmPropertiesResponseOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) *bool { return v.EnablePurgeProtection }).(pulumi.BoolPtrOutput)
}

func (o ManagedHsmPropertiesResponseOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) *bool { return v.EnableSoftDelete }).(pulumi.BoolPtrOutput)
}

func (o ManagedHsmPropertiesResponseOutput) HsmUri() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) string { return v.HsmUri }).(pulumi.StringOutput)
}

func (o ManagedHsmPropertiesResponseOutput) InitialAdminObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) []string { return v.InitialAdminObjectIds }).(pulumi.StringArrayOutput)
}

func (o ManagedHsmPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedHsmPropertiesResponseOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) *int { return v.SoftDeleteRetentionInDays }).(pulumi.IntPtrOutput)
}

func (o ManagedHsmPropertiesResponseOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) string { return v.StatusMessage }).(pulumi.StringOutput)
}

func (o ManagedHsmPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ManagedHsmPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedHsmPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsmPropertiesResponse)(nil)).Elem()
}

func (o ManagedHsmPropertiesResponsePtrOutput) ToManagedHsmPropertiesResponsePtrOutput() ManagedHsmPropertiesResponsePtrOutput {
	return o
}

func (o ManagedHsmPropertiesResponsePtrOutput) ToManagedHsmPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedHsmPropertiesResponsePtrOutput {
	return o
}

func (o ManagedHsmPropertiesResponsePtrOutput) Elem() ManagedHsmPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) ManagedHsmPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ManagedHsmPropertiesResponse
		return ret
	}).(ManagedHsmPropertiesResponseOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreateMode
	}).(pulumi.StringPtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePurgeProtection
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableSoftDelete
	}).(pulumi.BoolPtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) HsmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HsmUri
	}).(pulumi.StringPtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) InitialAdminObjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.InitialAdminObjectIds
	}).(pulumi.StringArrayOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.SoftDeleteRetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) StatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StatusMessage
	}).(pulumi.StringPtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ManagedHsmSku struct {
	Family string            `pulumi:"family"`
	Name   ManagedHsmSkuName `pulumi:"name"`
}





type ManagedHsmSkuInput interface {
	pulumi.Input

	ToManagedHsmSkuOutput() ManagedHsmSkuOutput
	ToManagedHsmSkuOutputWithContext(context.Context) ManagedHsmSkuOutput
}

type ManagedHsmSkuArgs struct {
	Family pulumi.StringInput     `pulumi:"family"`
	Name   ManagedHsmSkuNameInput `pulumi:"name"`
}

func (ManagedHsmSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHsmSku)(nil)).Elem()
}

func (i ManagedHsmSkuArgs) ToManagedHsmSkuOutput() ManagedHsmSkuOutput {
	return i.ToManagedHsmSkuOutputWithContext(context.Background())
}

func (i ManagedHsmSkuArgs) ToManagedHsmSkuOutputWithContext(ctx context.Context) ManagedHsmSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmSkuOutput)
}

func (i ManagedHsmSkuArgs) ToManagedHsmSkuPtrOutput() ManagedHsmSkuPtrOutput {
	return i.ToManagedHsmSkuPtrOutputWithContext(context.Background())
}

func (i ManagedHsmSkuArgs) ToManagedHsmSkuPtrOutputWithContext(ctx context.Context) ManagedHsmSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmSkuOutput).ToManagedHsmSkuPtrOutputWithContext(ctx)
}









type ManagedHsmSkuPtrInput interface {
	pulumi.Input

	ToManagedHsmSkuPtrOutput() ManagedHsmSkuPtrOutput
	ToManagedHsmSkuPtrOutputWithContext(context.Context) ManagedHsmSkuPtrOutput
}

type managedHsmSkuPtrType ManagedHsmSkuArgs

func ManagedHsmSkuPtr(v *ManagedHsmSkuArgs) ManagedHsmSkuPtrInput {
	return (*managedHsmSkuPtrType)(v)
}

func (*managedHsmSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsmSku)(nil)).Elem()
}

func (i *managedHsmSkuPtrType) ToManagedHsmSkuPtrOutput() ManagedHsmSkuPtrOutput {
	return i.ToManagedHsmSkuPtrOutputWithContext(context.Background())
}

func (i *managedHsmSkuPtrType) ToManagedHsmSkuPtrOutputWithContext(ctx context.Context) ManagedHsmSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmSkuPtrOutput)
}

type ManagedHsmSkuOutput struct{ *pulumi.OutputState }

func (ManagedHsmSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHsmSku)(nil)).Elem()
}

func (o ManagedHsmSkuOutput) ToManagedHsmSkuOutput() ManagedHsmSkuOutput {
	return o
}

func (o ManagedHsmSkuOutput) ToManagedHsmSkuOutputWithContext(ctx context.Context) ManagedHsmSkuOutput {
	return o
}

func (o ManagedHsmSkuOutput) ToManagedHsmSkuPtrOutput() ManagedHsmSkuPtrOutput {
	return o.ToManagedHsmSkuPtrOutputWithContext(context.Background())
}

func (o ManagedHsmSkuOutput) ToManagedHsmSkuPtrOutputWithContext(ctx context.Context) ManagedHsmSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedHsmSku) *ManagedHsmSku {
		return &v
	}).(ManagedHsmSkuPtrOutput)
}

func (o ManagedHsmSkuOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHsmSku) string { return v.Family }).(pulumi.StringOutput)
}

func (o ManagedHsmSkuOutput) Name() ManagedHsmSkuNameOutput {
	return o.ApplyT(func(v ManagedHsmSku) ManagedHsmSkuName { return v.Name }).(ManagedHsmSkuNameOutput)
}

type ManagedHsmSkuPtrOutput struct{ *pulumi.OutputState }

func (ManagedHsmSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsmSku)(nil)).Elem()
}

func (o ManagedHsmSkuPtrOutput) ToManagedHsmSkuPtrOutput() ManagedHsmSkuPtrOutput {
	return o
}

func (o ManagedHsmSkuPtrOutput) ToManagedHsmSkuPtrOutputWithContext(ctx context.Context) ManagedHsmSkuPtrOutput {
	return o
}

func (o ManagedHsmSkuPtrOutput) Elem() ManagedHsmSkuOutput {
	return o.ApplyT(func(v *ManagedHsmSku) ManagedHsmSku {
		if v != nil {
			return *v
		}
		var ret ManagedHsmSku
		return ret
	}).(ManagedHsmSkuOutput)
}

func (o ManagedHsmSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmSku) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ManagedHsmSkuPtrOutput) Name() ManagedHsmSkuNamePtrOutput {
	return o.ApplyT(func(v *ManagedHsmSku) *ManagedHsmSkuName {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(ManagedHsmSkuNamePtrOutput)
}

type ManagedHsmSkuResponse struct {
	Family string `pulumi:"family"`
	Name   string `pulumi:"name"`
}





type ManagedHsmSkuResponseInput interface {
	pulumi.Input

	ToManagedHsmSkuResponseOutput() ManagedHsmSkuResponseOutput
	ToManagedHsmSkuResponseOutputWithContext(context.Context) ManagedHsmSkuResponseOutput
}

type ManagedHsmSkuResponseArgs struct {
	Family pulumi.StringInput `pulumi:"family"`
	Name   pulumi.StringInput `pulumi:"name"`
}

func (ManagedHsmSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHsmSkuResponse)(nil)).Elem()
}

func (i ManagedHsmSkuResponseArgs) ToManagedHsmSkuResponseOutput() ManagedHsmSkuResponseOutput {
	return i.ToManagedHsmSkuResponseOutputWithContext(context.Background())
}

func (i ManagedHsmSkuResponseArgs) ToManagedHsmSkuResponseOutputWithContext(ctx context.Context) ManagedHsmSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmSkuResponseOutput)
}

func (i ManagedHsmSkuResponseArgs) ToManagedHsmSkuResponsePtrOutput() ManagedHsmSkuResponsePtrOutput {
	return i.ToManagedHsmSkuResponsePtrOutputWithContext(context.Background())
}

func (i ManagedHsmSkuResponseArgs) ToManagedHsmSkuResponsePtrOutputWithContext(ctx context.Context) ManagedHsmSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmSkuResponseOutput).ToManagedHsmSkuResponsePtrOutputWithContext(ctx)
}









type ManagedHsmSkuResponsePtrInput interface {
	pulumi.Input

	ToManagedHsmSkuResponsePtrOutput() ManagedHsmSkuResponsePtrOutput
	ToManagedHsmSkuResponsePtrOutputWithContext(context.Context) ManagedHsmSkuResponsePtrOutput
}

type managedHsmSkuResponsePtrType ManagedHsmSkuResponseArgs

func ManagedHsmSkuResponsePtr(v *ManagedHsmSkuResponseArgs) ManagedHsmSkuResponsePtrInput {
	return (*managedHsmSkuResponsePtrType)(v)
}

func (*managedHsmSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsmSkuResponse)(nil)).Elem()
}

func (i *managedHsmSkuResponsePtrType) ToManagedHsmSkuResponsePtrOutput() ManagedHsmSkuResponsePtrOutput {
	return i.ToManagedHsmSkuResponsePtrOutputWithContext(context.Background())
}

func (i *managedHsmSkuResponsePtrType) ToManagedHsmSkuResponsePtrOutputWithContext(ctx context.Context) ManagedHsmSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHsmSkuResponsePtrOutput)
}

type ManagedHsmSkuResponseOutput struct{ *pulumi.OutputState }

func (ManagedHsmSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHsmSkuResponse)(nil)).Elem()
}

func (o ManagedHsmSkuResponseOutput) ToManagedHsmSkuResponseOutput() ManagedHsmSkuResponseOutput {
	return o
}

func (o ManagedHsmSkuResponseOutput) ToManagedHsmSkuResponseOutputWithContext(ctx context.Context) ManagedHsmSkuResponseOutput {
	return o
}

func (o ManagedHsmSkuResponseOutput) ToManagedHsmSkuResponsePtrOutput() ManagedHsmSkuResponsePtrOutput {
	return o.ToManagedHsmSkuResponsePtrOutputWithContext(context.Background())
}

func (o ManagedHsmSkuResponseOutput) ToManagedHsmSkuResponsePtrOutputWithContext(ctx context.Context) ManagedHsmSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedHsmSkuResponse) *ManagedHsmSkuResponse {
		return &v
	}).(ManagedHsmSkuResponsePtrOutput)
}

func (o ManagedHsmSkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHsmSkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o ManagedHsmSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHsmSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ManagedHsmSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedHsmSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedHsmSkuResponse)(nil)).Elem()
}

func (o ManagedHsmSkuResponsePtrOutput) ToManagedHsmSkuResponsePtrOutput() ManagedHsmSkuResponsePtrOutput {
	return o
}

func (o ManagedHsmSkuResponsePtrOutput) ToManagedHsmSkuResponsePtrOutputWithContext(ctx context.Context) ManagedHsmSkuResponsePtrOutput {
	return o
}

func (o ManagedHsmSkuResponsePtrOutput) Elem() ManagedHsmSkuResponseOutput {
	return o.ApplyT(func(v *ManagedHsmSkuResponse) ManagedHsmSkuResponse {
		if v != nil {
			return *v
		}
		var ret ManagedHsmSkuResponse
		return ret
	}).(ManagedHsmSkuResponseOutput)
}

func (o ManagedHsmSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ManagedHsmSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type NetworkRuleSet struct {
	Bypass              *string              `pulumi:"bypass"`
	DefaultAction       *string              `pulumi:"defaultAction"`
	IpRules             []IPRule             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	Bypass              pulumi.StringPtrInput        `pulumi:"bypass"`
	DefaultAction       pulumi.StringPtrInput        `pulumi:"defaultAction"`
	IpRules             IPRuleArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSet) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSet) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IPRule { return v.IpRules }).(IPRuleArrayOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IPRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	Bypass              *string                      `pulumi:"bypass"`
	DefaultAction       *string                      `pulumi:"defaultAction"`
	IpRules             []IPRuleResponse             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}





type NetworkRuleSetResponseInput interface {
	pulumi.Input

	ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput
	ToNetworkRuleSetResponseOutputWithContext(context.Context) NetworkRuleSetResponseOutput
}

type NetworkRuleSetResponseArgs struct {
	Bypass              pulumi.StringPtrInput                `pulumi:"bypass"`
	DefaultAction       pulumi.StringPtrInput                `pulumi:"defaultAction"`
	IpRules             IPRuleResponseArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleResponseArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return i.ToNetworkRuleSetResponseOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput)
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput).ToNetworkRuleSetResponsePtrOutputWithContext(ctx)
}









type NetworkRuleSetResponsePtrInput interface {
	pulumi.Input

	ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput
	ToNetworkRuleSetResponsePtrOutputWithContext(context.Context) NetworkRuleSetResponsePtrOutput
}

type networkRuleSetResponsePtrType NetworkRuleSetResponseArgs

func NetworkRuleSetResponsePtr(v *NetworkRuleSetResponseArgs) NetworkRuleSetResponsePtrInput {
	return (*networkRuleSetResponsePtrType)(v)
}

func (*networkRuleSetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponsePtrOutput)
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSetResponse) *NetworkRuleSetResponse {
		return &v
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o NetworkRuleSetResponseOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IPRuleResponse { return v.IpRules }).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IPRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponsePtrOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []VirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleResponseArrayOutput)
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





type PermissionsResponseInput interface {
	pulumi.Input

	ToPermissionsResponseOutput() PermissionsResponseOutput
	ToPermissionsResponseOutputWithContext(context.Context) PermissionsResponseOutput
}

type PermissionsResponseArgs struct {
	Certificates pulumi.StringArrayInput `pulumi:"certificates"`
	Keys         pulumi.StringArrayInput `pulumi:"keys"`
	Secrets      pulumi.StringArrayInput `pulumi:"secrets"`
	Storage      pulumi.StringArrayInput `pulumi:"storage"`
}

func (PermissionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResponse)(nil)).Elem()
}

func (i PermissionsResponseArgs) ToPermissionsResponseOutput() PermissionsResponseOutput {
	return i.ToPermissionsResponseOutputWithContext(context.Background())
}

func (i PermissionsResponseArgs) ToPermissionsResponseOutputWithContext(ctx context.Context) PermissionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResponseOutput)
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

type PrivateEndpointConnectionItemResponse struct {
	Etag                              *string                                    `pulumi:"etag"`
	Id                                *string                                    `pulumi:"id"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionItemResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionItemResponseOutput() PrivateEndpointConnectionItemResponseOutput
	ToPrivateEndpointConnectionItemResponseOutputWithContext(context.Context) PrivateEndpointConnectionItemResponseOutput
}

type PrivateEndpointConnectionItemResponseArgs struct {
	Etag                              pulumi.StringPtrInput                             `pulumi:"etag"`
	Id                                pulumi.StringPtrInput                             `pulumi:"id"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionItemResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionItemResponseArgs) ToPrivateEndpointConnectionItemResponseOutput() PrivateEndpointConnectionItemResponseOutput {
	return i.ToPrivateEndpointConnectionItemResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionItemResponseArgs) ToPrivateEndpointConnectionItemResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionItemResponseOutput)
}





type PrivateEndpointConnectionItemResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionItemResponseArrayOutput() PrivateEndpointConnectionItemResponseArrayOutput
	ToPrivateEndpointConnectionItemResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionItemResponseArrayOutput
}

type PrivateEndpointConnectionItemResponseArray []PrivateEndpointConnectionItemResponseInput

func (PrivateEndpointConnectionItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionItemResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionItemResponseArray) ToPrivateEndpointConnectionItemResponseArrayOutput() PrivateEndpointConnectionItemResponseArrayOutput {
	return i.ToPrivateEndpointConnectionItemResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionItemResponseArray) ToPrivateEndpointConnectionItemResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionItemResponseArrayOutput)
}

type PrivateEndpointConnectionItemResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionItemResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionItemResponseOutput) ToPrivateEndpointConnectionItemResponseOutput() PrivateEndpointConnectionItemResponseOutput {
	return o
}

func (o PrivateEndpointConnectionItemResponseOutput) ToPrivateEndpointConnectionItemResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionItemResponseOutput {
	return o
}

func (o PrivateEndpointConnectionItemResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionItemResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionItemResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionItemResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionItemResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionItemResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionItemResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionItemResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionItemResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionItemResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionItemResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionItemResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionItemResponseArrayOutput) ToPrivateEndpointConnectionItemResponseArrayOutput() PrivateEndpointConnectionItemResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionItemResponseArrayOutput) ToPrivateEndpointConnectionItemResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionItemResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionItemResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionItemResponse {
		return vs[0].([]PrivateEndpointConnectionItemResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionItemResponseOutput)
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





type SecretAttributesResponseInput interface {
	pulumi.Input

	ToSecretAttributesResponseOutput() SecretAttributesResponseOutput
	ToSecretAttributesResponseOutputWithContext(context.Context) SecretAttributesResponseOutput
}

type SecretAttributesResponseArgs struct {
	Created   pulumi.IntInput     `pulumi:"created"`
	Enabled   pulumi.BoolPtrInput `pulumi:"enabled"`
	Expires   pulumi.IntPtrInput  `pulumi:"expires"`
	NotBefore pulumi.IntPtrInput  `pulumi:"notBefore"`
	Updated   pulumi.IntInput     `pulumi:"updated"`
}

func (SecretAttributesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretAttributesResponse)(nil)).Elem()
}

func (i SecretAttributesResponseArgs) ToSecretAttributesResponseOutput() SecretAttributesResponseOutput {
	return i.ToSecretAttributesResponseOutputWithContext(context.Background())
}

func (i SecretAttributesResponseArgs) ToSecretAttributesResponseOutputWithContext(ctx context.Context) SecretAttributesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretAttributesResponseOutput)
}

func (i SecretAttributesResponseArgs) ToSecretAttributesResponsePtrOutput() SecretAttributesResponsePtrOutput {
	return i.ToSecretAttributesResponsePtrOutputWithContext(context.Background())
}

func (i SecretAttributesResponseArgs) ToSecretAttributesResponsePtrOutputWithContext(ctx context.Context) SecretAttributesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretAttributesResponseOutput).ToSecretAttributesResponsePtrOutputWithContext(ctx)
}









type SecretAttributesResponsePtrInput interface {
	pulumi.Input

	ToSecretAttributesResponsePtrOutput() SecretAttributesResponsePtrOutput
	ToSecretAttributesResponsePtrOutputWithContext(context.Context) SecretAttributesResponsePtrOutput
}

type secretAttributesResponsePtrType SecretAttributesResponseArgs

func SecretAttributesResponsePtr(v *SecretAttributesResponseArgs) SecretAttributesResponsePtrInput {
	return (*secretAttributesResponsePtrType)(v)
}

func (*secretAttributesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretAttributesResponse)(nil)).Elem()
}

func (i *secretAttributesResponsePtrType) ToSecretAttributesResponsePtrOutput() SecretAttributesResponsePtrOutput {
	return i.ToSecretAttributesResponsePtrOutputWithContext(context.Background())
}

func (i *secretAttributesResponsePtrType) ToSecretAttributesResponsePtrOutputWithContext(ctx context.Context) SecretAttributesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretAttributesResponsePtrOutput)
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

func (o SecretAttributesResponseOutput) ToSecretAttributesResponsePtrOutput() SecretAttributesResponsePtrOutput {
	return o.ToSecretAttributesResponsePtrOutputWithContext(context.Background())
}

func (o SecretAttributesResponseOutput) ToSecretAttributesResponsePtrOutputWithContext(ctx context.Context) SecretAttributesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretAttributesResponse) *SecretAttributesResponse {
		return &v
	}).(SecretAttributesResponsePtrOutput)
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

func (i SecretPropertiesArgs) ToSecretPropertiesPtrOutput() SecretPropertiesPtrOutput {
	return i.ToSecretPropertiesPtrOutputWithContext(context.Background())
}

func (i SecretPropertiesArgs) ToSecretPropertiesPtrOutputWithContext(ctx context.Context) SecretPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPropertiesOutput).ToSecretPropertiesPtrOutputWithContext(ctx)
}









type SecretPropertiesPtrInput interface {
	pulumi.Input

	ToSecretPropertiesPtrOutput() SecretPropertiesPtrOutput
	ToSecretPropertiesPtrOutputWithContext(context.Context) SecretPropertiesPtrOutput
}

type secretPropertiesPtrType SecretPropertiesArgs

func SecretPropertiesPtr(v *SecretPropertiesArgs) SecretPropertiesPtrInput {
	return (*secretPropertiesPtrType)(v)
}

func (*secretPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretProperties)(nil)).Elem()
}

func (i *secretPropertiesPtrType) ToSecretPropertiesPtrOutput() SecretPropertiesPtrOutput {
	return i.ToSecretPropertiesPtrOutputWithContext(context.Background())
}

func (i *secretPropertiesPtrType) ToSecretPropertiesPtrOutputWithContext(ctx context.Context) SecretPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPropertiesPtrOutput)
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

func (o SecretPropertiesOutput) ToSecretPropertiesPtrOutput() SecretPropertiesPtrOutput {
	return o.ToSecretPropertiesPtrOutputWithContext(context.Background())
}

func (o SecretPropertiesOutput) ToSecretPropertiesPtrOutputWithContext(ctx context.Context) SecretPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretProperties) *SecretProperties {
		return &v
	}).(SecretPropertiesPtrOutput)
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

type SecretPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SecretPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretProperties)(nil)).Elem()
}

func (o SecretPropertiesPtrOutput) ToSecretPropertiesPtrOutput() SecretPropertiesPtrOutput {
	return o
}

func (o SecretPropertiesPtrOutput) ToSecretPropertiesPtrOutputWithContext(ctx context.Context) SecretPropertiesPtrOutput {
	return o
}

func (o SecretPropertiesPtrOutput) Elem() SecretPropertiesOutput {
	return o.ApplyT(func(v *SecretProperties) SecretProperties {
		if v != nil {
			return *v
		}
		var ret SecretProperties
		return ret
	}).(SecretPropertiesOutput)
}

func (o SecretPropertiesPtrOutput) Attributes() SecretAttributesPtrOutput {
	return o.ApplyT(func(v *SecretProperties) *SecretAttributes {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(SecretAttributesPtrOutput)
}

func (o SecretPropertiesPtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretProperties) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o SecretPropertiesPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretProperties) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type SecretPropertiesResponse struct {
	Attributes           *SecretAttributesResponse `pulumi:"attributes"`
	ContentType          *string                   `pulumi:"contentType"`
	SecretUri            string                    `pulumi:"secretUri"`
	SecretUriWithVersion string                    `pulumi:"secretUriWithVersion"`
	Value                *string                   `pulumi:"value"`
}





type SecretPropertiesResponseInput interface {
	pulumi.Input

	ToSecretPropertiesResponseOutput() SecretPropertiesResponseOutput
	ToSecretPropertiesResponseOutputWithContext(context.Context) SecretPropertiesResponseOutput
}

type SecretPropertiesResponseArgs struct {
	Attributes           SecretAttributesResponsePtrInput `pulumi:"attributes"`
	ContentType          pulumi.StringPtrInput            `pulumi:"contentType"`
	SecretUri            pulumi.StringInput               `pulumi:"secretUri"`
	SecretUriWithVersion pulumi.StringInput               `pulumi:"secretUriWithVersion"`
	Value                pulumi.StringPtrInput            `pulumi:"value"`
}

func (SecretPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretPropertiesResponse)(nil)).Elem()
}

func (i SecretPropertiesResponseArgs) ToSecretPropertiesResponseOutput() SecretPropertiesResponseOutput {
	return i.ToSecretPropertiesResponseOutputWithContext(context.Background())
}

func (i SecretPropertiesResponseArgs) ToSecretPropertiesResponseOutputWithContext(ctx context.Context) SecretPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPropertiesResponseOutput)
}

func (i SecretPropertiesResponseArgs) ToSecretPropertiesResponsePtrOutput() SecretPropertiesResponsePtrOutput {
	return i.ToSecretPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SecretPropertiesResponseArgs) ToSecretPropertiesResponsePtrOutputWithContext(ctx context.Context) SecretPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPropertiesResponseOutput).ToSecretPropertiesResponsePtrOutputWithContext(ctx)
}









type SecretPropertiesResponsePtrInput interface {
	pulumi.Input

	ToSecretPropertiesResponsePtrOutput() SecretPropertiesResponsePtrOutput
	ToSecretPropertiesResponsePtrOutputWithContext(context.Context) SecretPropertiesResponsePtrOutput
}

type secretPropertiesResponsePtrType SecretPropertiesResponseArgs

func SecretPropertiesResponsePtr(v *SecretPropertiesResponseArgs) SecretPropertiesResponsePtrInput {
	return (*secretPropertiesResponsePtrType)(v)
}

func (*secretPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretPropertiesResponse)(nil)).Elem()
}

func (i *secretPropertiesResponsePtrType) ToSecretPropertiesResponsePtrOutput() SecretPropertiesResponsePtrOutput {
	return i.ToSecretPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *secretPropertiesResponsePtrType) ToSecretPropertiesResponsePtrOutputWithContext(ctx context.Context) SecretPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPropertiesResponsePtrOutput)
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

func (o SecretPropertiesResponseOutput) ToSecretPropertiesResponsePtrOutput() SecretPropertiesResponsePtrOutput {
	return o.ToSecretPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SecretPropertiesResponseOutput) ToSecretPropertiesResponsePtrOutputWithContext(ctx context.Context) SecretPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretPropertiesResponse) *SecretPropertiesResponse {
		return &v
	}).(SecretPropertiesResponsePtrOutput)
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

type SecretPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SecretPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretPropertiesResponse)(nil)).Elem()
}

func (o SecretPropertiesResponsePtrOutput) ToSecretPropertiesResponsePtrOutput() SecretPropertiesResponsePtrOutput {
	return o
}

func (o SecretPropertiesResponsePtrOutput) ToSecretPropertiesResponsePtrOutputWithContext(ctx context.Context) SecretPropertiesResponsePtrOutput {
	return o
}

func (o SecretPropertiesResponsePtrOutput) Elem() SecretPropertiesResponseOutput {
	return o.ApplyT(func(v *SecretPropertiesResponse) SecretPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SecretPropertiesResponse
		return ret
	}).(SecretPropertiesResponseOutput)
}

func (o SecretPropertiesResponsePtrOutput) Attributes() SecretAttributesResponsePtrOutput {
	return o.ApplyT(func(v *SecretPropertiesResponse) *SecretAttributesResponse {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(SecretAttributesResponsePtrOutput)
}

func (o SecretPropertiesResponsePtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o SecretPropertiesResponsePtrOutput) SecretUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUri
	}).(pulumi.StringPtrOutput)
}

func (o SecretPropertiesResponsePtrOutput) SecretUriWithVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUriWithVersion
	}).(pulumi.StringPtrOutput)
}

func (o SecretPropertiesResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
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

func (o SkuOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuOutput) Name() SkuNameOutput {
	return o.ApplyT(func(v Sku) SkuName { return v.Name }).(SkuNameOutput)
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

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v *Sku) *SkuName {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(SkuNamePtrOutput)
}

type SkuResponse struct {
	Family string `pulumi:"family"`
	Name   string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Family pulumi.StringInput `pulumi:"family"`
	Name   pulumi.StringInput `pulumi:"name"`
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

func (o SkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Family }).(pulumi.StringOutput)
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

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
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

type VaultProperties struct {
	AccessPolicies               []AccessPolicyEntry `pulumi:"accessPolicies"`
	CreateMode                   *CreateMode         `pulumi:"createMode"`
	EnablePurgeProtection        *bool               `pulumi:"enablePurgeProtection"`
	EnableRbacAuthorization      *bool               `pulumi:"enableRbacAuthorization"`
	EnableSoftDelete             *bool               `pulumi:"enableSoftDelete"`
	EnabledForDeployment         *bool               `pulumi:"enabledForDeployment"`
	EnabledForDiskEncryption     *bool               `pulumi:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment *bool               `pulumi:"enabledForTemplateDeployment"`
	NetworkAcls                  *NetworkRuleSet     `pulumi:"networkAcls"`
	ProvisioningState            *string             `pulumi:"provisioningState"`
	Sku                          Sku                 `pulumi:"sku"`
	SoftDeleteRetentionInDays    *int                `pulumi:"softDeleteRetentionInDays"`
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
	EnableRbacAuthorization      pulumi.BoolPtrInput         `pulumi:"enableRbacAuthorization"`
	EnableSoftDelete             pulumi.BoolPtrInput         `pulumi:"enableSoftDelete"`
	EnabledForDeployment         pulumi.BoolPtrInput         `pulumi:"enabledForDeployment"`
	EnabledForDiskEncryption     pulumi.BoolPtrInput         `pulumi:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment pulumi.BoolPtrInput         `pulumi:"enabledForTemplateDeployment"`
	NetworkAcls                  NetworkRuleSetPtrInput      `pulumi:"networkAcls"`
	ProvisioningState            pulumi.StringPtrInput       `pulumi:"provisioningState"`
	Sku                          SkuInput                    `pulumi:"sku"`
	SoftDeleteRetentionInDays    pulumi.IntPtrInput          `pulumi:"softDeleteRetentionInDays"`
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

func (i VaultPropertiesArgs) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return i.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i VaultPropertiesArgs) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesOutput).ToVaultPropertiesPtrOutputWithContext(ctx)
}









type VaultPropertiesPtrInput interface {
	pulumi.Input

	ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput
	ToVaultPropertiesPtrOutputWithContext(context.Context) VaultPropertiesPtrOutput
}

type vaultPropertiesPtrType VaultPropertiesArgs

func VaultPropertiesPtr(v *VaultPropertiesArgs) VaultPropertiesPtrInput {
	return (*vaultPropertiesPtrType)(v)
}

func (*vaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultProperties)(nil)).Elem()
}

func (i *vaultPropertiesPtrType) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return i.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *vaultPropertiesPtrType) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesPtrOutput)
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

func (o VaultPropertiesOutput) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return o.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o VaultPropertiesOutput) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VaultProperties) *VaultProperties {
		return &v
	}).(VaultPropertiesPtrOutput)
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

func (o VaultPropertiesOutput) EnableRbacAuthorization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultProperties) *bool { return v.EnableRbacAuthorization }).(pulumi.BoolPtrOutput)
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

func (o VaultPropertiesOutput) NetworkAcls() NetworkRuleSetPtrOutput {
	return o.ApplyT(func(v VaultProperties) *NetworkRuleSet { return v.NetworkAcls }).(NetworkRuleSetPtrOutput)
}

func (o VaultPropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesOutput) Sku() SkuOutput {
	return o.ApplyT(func(v VaultProperties) Sku { return v.Sku }).(SkuOutput)
}

func (o VaultPropertiesOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VaultProperties) *int { return v.SoftDeleteRetentionInDays }).(pulumi.IntPtrOutput)
}

func (o VaultPropertiesOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultProperties) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o VaultPropertiesOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultProperties) *string { return v.VaultUri }).(pulumi.StringPtrOutput)
}

type VaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultProperties)(nil)).Elem()
}

func (o VaultPropertiesPtrOutput) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return o
}

func (o VaultPropertiesPtrOutput) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return o
}

func (o VaultPropertiesPtrOutput) Elem() VaultPropertiesOutput {
	return o.ApplyT(func(v *VaultProperties) VaultProperties {
		if v != nil {
			return *v
		}
		var ret VaultProperties
		return ret
	}).(VaultPropertiesOutput)
}

func (o VaultPropertiesPtrOutput) AccessPolicies() AccessPolicyEntryArrayOutput {
	return o.ApplyT(func(v *VaultProperties) []AccessPolicyEntry {
		if v == nil {
			return nil
		}
		return v.AccessPolicies
	}).(AccessPolicyEntryArrayOutput)
}

func (o VaultPropertiesPtrOutput) CreateMode() CreateModePtrOutput {
	return o.ApplyT(func(v *VaultProperties) *CreateMode {
		if v == nil {
			return nil
		}
		return v.CreateMode
	}).(CreateModePtrOutput)
}

func (o VaultPropertiesPtrOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePurgeProtection
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesPtrOutput) EnableRbacAuthorization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableRbacAuthorization
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesPtrOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableSoftDelete
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesPtrOutput) EnabledForDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnabledForDeployment
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesPtrOutput) EnabledForDiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnabledForDiskEncryption
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesPtrOutput) EnabledForTemplateDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnabledForTemplateDeployment
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesPtrOutput) NetworkAcls() NetworkRuleSetPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *NetworkRuleSet {
		if v == nil {
			return nil
		}
		return v.NetworkAcls
	}).(NetworkRuleSetPtrOutput)
}

func (o VaultPropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesPtrOutput) Sku() SkuPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *Sku {
		if v == nil {
			return nil
		}
		return &v.Sku
	}).(SkuPtrOutput)
}

func (o VaultPropertiesPtrOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *int {
		if v == nil {
			return nil
		}
		return v.SoftDeleteRetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o VaultPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesPtrOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.VaultUri
	}).(pulumi.StringPtrOutput)
}

type VaultPropertiesResponse struct {
	AccessPolicies               []AccessPolicyEntryResponse             `pulumi:"accessPolicies"`
	CreateMode                   *string                                 `pulumi:"createMode"`
	EnablePurgeProtection        *bool                                   `pulumi:"enablePurgeProtection"`
	EnableRbacAuthorization      *bool                                   `pulumi:"enableRbacAuthorization"`
	EnableSoftDelete             *bool                                   `pulumi:"enableSoftDelete"`
	EnabledForDeployment         *bool                                   `pulumi:"enabledForDeployment"`
	EnabledForDiskEncryption     *bool                                   `pulumi:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment *bool                                   `pulumi:"enabledForTemplateDeployment"`
	NetworkAcls                  *NetworkRuleSetResponse                 `pulumi:"networkAcls"`
	PrivateEndpointConnections   []PrivateEndpointConnectionItemResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState            *string                                 `pulumi:"provisioningState"`
	Sku                          SkuResponse                             `pulumi:"sku"`
	SoftDeleteRetentionInDays    *int                                    `pulumi:"softDeleteRetentionInDays"`
	TenantId                     string                                  `pulumi:"tenantId"`
	VaultUri                     *string                                 `pulumi:"vaultUri"`
}





type VaultPropertiesResponseInput interface {
	pulumi.Input

	ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput
	ToVaultPropertiesResponseOutputWithContext(context.Context) VaultPropertiesResponseOutput
}

type VaultPropertiesResponseArgs struct {
	AccessPolicies               AccessPolicyEntryResponseArrayInput             `pulumi:"accessPolicies"`
	CreateMode                   pulumi.StringPtrInput                           `pulumi:"createMode"`
	EnablePurgeProtection        pulumi.BoolPtrInput                             `pulumi:"enablePurgeProtection"`
	EnableRbacAuthorization      pulumi.BoolPtrInput                             `pulumi:"enableRbacAuthorization"`
	EnableSoftDelete             pulumi.BoolPtrInput                             `pulumi:"enableSoftDelete"`
	EnabledForDeployment         pulumi.BoolPtrInput                             `pulumi:"enabledForDeployment"`
	EnabledForDiskEncryption     pulumi.BoolPtrInput                             `pulumi:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment pulumi.BoolPtrInput                             `pulumi:"enabledForTemplateDeployment"`
	NetworkAcls                  NetworkRuleSetResponsePtrInput                  `pulumi:"networkAcls"`
	PrivateEndpointConnections   PrivateEndpointConnectionItemResponseArrayInput `pulumi:"privateEndpointConnections"`
	ProvisioningState            pulumi.StringPtrInput                           `pulumi:"provisioningState"`
	Sku                          SkuResponseInput                                `pulumi:"sku"`
	SoftDeleteRetentionInDays    pulumi.IntPtrInput                              `pulumi:"softDeleteRetentionInDays"`
	TenantId                     pulumi.StringInput                              `pulumi:"tenantId"`
	VaultUri                     pulumi.StringPtrInput                           `pulumi:"vaultUri"`
}

func (VaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponse)(nil)).Elem()
}

func (i VaultPropertiesResponseArgs) ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput {
	return i.ToVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i VaultPropertiesResponseArgs) ToVaultPropertiesResponseOutputWithContext(ctx context.Context) VaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesResponseOutput)
}

func (i VaultPropertiesResponseArgs) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return i.ToVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i VaultPropertiesResponseArgs) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesResponseOutput).ToVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type VaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput
	ToVaultPropertiesResponsePtrOutputWithContext(context.Context) VaultPropertiesResponsePtrOutput
}

type vaultPropertiesResponsePtrType VaultPropertiesResponseArgs

func VaultPropertiesResponsePtr(v *VaultPropertiesResponseArgs) VaultPropertiesResponsePtrInput {
	return (*vaultPropertiesResponsePtrType)(v)
}

func (*vaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponse)(nil)).Elem()
}

func (i *vaultPropertiesResponsePtrType) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return i.ToVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *vaultPropertiesResponsePtrType) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesResponsePtrOutput)
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

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return o.ToVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VaultPropertiesResponse) *VaultPropertiesResponse {
		return &v
	}).(VaultPropertiesResponsePtrOutput)
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

func (o VaultPropertiesResponseOutput) EnableRbacAuthorization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *bool { return v.EnableRbacAuthorization }).(pulumi.BoolPtrOutput)
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

func (o VaultPropertiesResponseOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *NetworkRuleSetResponse { return v.NetworkAcls }).(NetworkRuleSetResponsePtrOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionItemResponseArrayOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) []PrivateEndpointConnectionItemResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionItemResponseArrayOutput)
}

func (o VaultPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o VaultPropertiesResponseOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *int { return v.SoftDeleteRetentionInDays }).(pulumi.IntPtrOutput)
}

func (o VaultPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *string { return v.VaultUri }).(pulumi.StringPtrOutput)
}

type VaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponse)(nil)).Elem()
}

func (o VaultPropertiesResponsePtrOutput) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return o
}

func (o VaultPropertiesResponsePtrOutput) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return o
}

func (o VaultPropertiesResponsePtrOutput) Elem() VaultPropertiesResponseOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) VaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesResponse
		return ret
	}).(VaultPropertiesResponseOutput)
}

func (o VaultPropertiesResponsePtrOutput) AccessPolicies() AccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) []AccessPolicyEntryResponse {
		if v == nil {
			return nil
		}
		return v.AccessPolicies
	}).(AccessPolicyEntryResponseArrayOutput)
}

func (o VaultPropertiesResponsePtrOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreateMode
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePurgeProtection
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) EnableRbacAuthorization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableRbacAuthorization
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) EnableSoftDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableSoftDelete
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) EnabledForDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnabledForDeployment
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) EnabledForDiskEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnabledForDiskEncryption
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) EnabledForTemplateDeployment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnabledForTemplateDeployment
	}).(pulumi.BoolPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *NetworkRuleSetResponse {
		if v == nil {
			return nil
		}
		return v.NetworkAcls
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionItemResponseArrayOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) []PrivateEndpointConnectionItemResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionItemResponseArrayOutput)
}

func (o VaultPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *SkuResponse {
		if v == nil {
			return nil
		}
		return &v.Sku
	}).(SkuResponsePtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.SoftDeleteRetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VaultUri
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkRule struct {
	Id                               string `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint *bool  `pulumi:"ignoreMissingVnetServiceEndpoint"`
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Id                               pulumi.StringInput  `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput `pulumi:"ignoreMissingVnetServiceEndpoint"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkRuleOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Id                               string `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint *bool  `pulumi:"ignoreMissingVnetServiceEndpoint"`
}





type VirtualNetworkRuleResponseInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput
	ToVirtualNetworkRuleResponseOutputWithContext(context.Context) VirtualNetworkRuleResponseOutput
}

type VirtualNetworkRuleResponseArgs struct {
	Id                               pulumi.StringInput  `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput `pulumi:"ignoreMissingVnetServiceEndpoint"`
}

func (VirtualNetworkRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return i.ToVirtualNetworkRuleResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseOutput)
}





type VirtualNetworkRuleResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput
	ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Context) VirtualNetworkRuleResponseArrayOutput
}

type VirtualNetworkRuleResponseArray []VirtualNetworkRuleResponseInput

func (VirtualNetworkRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return i.ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseArrayOutput)
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

func (o VirtualNetworkRuleResponseOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyEntryInput)(nil)).Elem(), AccessPolicyEntryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyEntryArrayInput)(nil)).Elem(), AccessPolicyEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyEntryResponseInput)(nil)).Elem(), AccessPolicyEntryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPolicyEntryResponseArrayInput)(nil)).Elem(), AccessPolicyEntryResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRuleInput)(nil)).Elem(), IPRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRuleArrayInput)(nil)).Elem(), IPRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRuleResponseInput)(nil)).Elem(), IPRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPRuleResponseArrayInput)(nil)).Elem(), IPRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyAttributesInput)(nil)).Elem(), KeyAttributesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyAttributesPtrInput)(nil)).Elem(), KeyAttributesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyAttributesResponseInput)(nil)).Elem(), KeyAttributesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyAttributesResponsePtrInput)(nil)).Elem(), KeyAttributesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPropertiesInput)(nil)).Elem(), KeyPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPropertiesPtrInput)(nil)).Elem(), KeyPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedHsmPropertiesInput)(nil)).Elem(), ManagedHsmPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedHsmPropertiesPtrInput)(nil)).Elem(), ManagedHsmPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedHsmPropertiesResponseInput)(nil)).Elem(), ManagedHsmPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedHsmPropertiesResponsePtrInput)(nil)).Elem(), ManagedHsmPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedHsmSkuInput)(nil)).Elem(), ManagedHsmSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedHsmSkuPtrInput)(nil)).Elem(), ManagedHsmSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedHsmSkuResponseInput)(nil)).Elem(), ManagedHsmSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedHsmSkuResponsePtrInput)(nil)).Elem(), ManagedHsmSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleSetInput)(nil)).Elem(), NetworkRuleSetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleSetPtrInput)(nil)).Elem(), NetworkRuleSetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleSetResponseInput)(nil)).Elem(), NetworkRuleSetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleSetResponsePtrInput)(nil)).Elem(), NetworkRuleSetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsInput)(nil)).Elem(), PermissionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsResponseInput)(nil)).Elem(), PermissionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionItemResponseInput)(nil)).Elem(), PrivateEndpointConnectionItemResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionItemResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionItemResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponseInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponsePtrInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretAttributesInput)(nil)).Elem(), SecretAttributesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretAttributesPtrInput)(nil)).Elem(), SecretAttributesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretAttributesResponseInput)(nil)).Elem(), SecretAttributesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretAttributesResponsePtrInput)(nil)).Elem(), SecretAttributesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretPropertiesInput)(nil)).Elem(), SecretPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretPropertiesPtrInput)(nil)).Elem(), SecretPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretPropertiesResponseInput)(nil)).Elem(), SecretPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretPropertiesResponsePtrInput)(nil)).Elem(), SecretPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultPropertiesInput)(nil)).Elem(), VaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultPropertiesPtrInput)(nil)).Elem(), VaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultPropertiesResponseInput)(nil)).Elem(), VaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultPropertiesResponsePtrInput)(nil)).Elem(), VaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkRuleInput)(nil)).Elem(), VirtualNetworkRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkRuleArrayInput)(nil)).Elem(), VirtualNetworkRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkRuleResponseInput)(nil)).Elem(), VirtualNetworkRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkRuleResponseArrayInput)(nil)).Elem(), VirtualNetworkRuleResponseArray{})
	pulumi.RegisterOutputType(AccessPolicyEntryOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryArrayOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryResponseOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyAttributesOutput{})
	pulumi.RegisterOutputType(KeyAttributesPtrOutput{})
	pulumi.RegisterOutputType(KeyAttributesResponseOutput{})
	pulumi.RegisterOutputType(KeyAttributesResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyPropertiesOutput{})
	pulumi.RegisterOutputType(KeyPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedHsmPropertiesOutput{})
	pulumi.RegisterOutputType(ManagedHsmPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedHsmPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagedHsmPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedHsmSkuOutput{})
	pulumi.RegisterOutputType(ManagedHsmSkuPtrOutput{})
	pulumi.RegisterOutputType(ManagedHsmSkuResponseOutput{})
	pulumi.RegisterOutputType(ManagedHsmSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(PermissionsOutput{})
	pulumi.RegisterOutputType(PermissionsResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionItemResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionItemResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(SecretAttributesOutput{})
	pulumi.RegisterOutputType(SecretAttributesPtrOutput{})
	pulumi.RegisterOutputType(SecretAttributesResponseOutput{})
	pulumi.RegisterOutputType(SecretAttributesResponsePtrOutput{})
	pulumi.RegisterOutputType(SecretPropertiesOutput{})
	pulumi.RegisterOutputType(SecretPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SecretPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SecretPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesOutput{})
	pulumi.RegisterOutputType(VaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
