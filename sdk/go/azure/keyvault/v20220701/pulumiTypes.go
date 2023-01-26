


package v20220701

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

type Action struct {
	Type *KeyRotationPolicyActionType `pulumi:"type"`
}





type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

type ActionArgs struct {
	Type KeyRotationPolicyActionTypePtrInput `pulumi:"type"`
}

func (ActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (i ActionArgs) ToActionOutput() ActionOutput {
	return i.ToActionOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput)
}

func (i ActionArgs) ToActionPtrOutput() ActionPtrOutput {
	return i.ToActionPtrOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput).ToActionPtrOutputWithContext(ctx)
}









type ActionPtrInput interface {
	pulumi.Input

	ToActionPtrOutput() ActionPtrOutput
	ToActionPtrOutputWithContext(context.Context) ActionPtrOutput
}

type actionPtrType ActionArgs

func ActionPtr(v *ActionArgs) ActionPtrInput {
	return (*actionPtrType)(v)
}

func (*actionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (i *actionPtrType) ToActionPtrOutput() ActionPtrOutput {
	return i.ToActionPtrOutputWithContext(context.Background())
}

func (i *actionPtrType) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionPtrOutput)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

func (o ActionOutput) ToActionPtrOutput() ActionPtrOutput {
	return o.ToActionPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Action) *Action {
		return &v
	}).(ActionPtrOutput)
}

func (o ActionOutput) Type() KeyRotationPolicyActionTypePtrOutput {
	return o.ApplyT(func(v Action) *KeyRotationPolicyActionType { return v.Type }).(KeyRotationPolicyActionTypePtrOutput)
}

type ActionPtrOutput struct{ *pulumi.OutputState }

func (ActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionPtrOutput) ToActionPtrOutput() ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) Elem() ActionOutput {
	return o.ApplyT(func(v *Action) Action {
		if v != nil {
			return *v
		}
		var ret Action
		return ret
	}).(ActionOutput)
}

func (o ActionPtrOutput) Type() KeyRotationPolicyActionTypePtrOutput {
	return o.ApplyT(func(v *Action) *KeyRotationPolicyActionType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(KeyRotationPolicyActionTypePtrOutput)
}

type ActionResponse struct {
	Type *string `pulumi:"type"`
}

type ActionResponseOutput struct{ *pulumi.OutputState }

func (ActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionResponse)(nil)).Elem()
}

func (o ActionResponseOutput) ToActionResponseOutput() ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) ToActionResponseOutputWithContext(ctx context.Context) ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ActionResponsePtrOutput struct{ *pulumi.OutputState }

func (ActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionResponse)(nil)).Elem()
}

func (o ActionResponsePtrOutput) ToActionResponsePtrOutput() ActionResponsePtrOutput {
	return o
}

func (o ActionResponsePtrOutput) ToActionResponsePtrOutputWithContext(ctx context.Context) ActionResponsePtrOutput {
	return o
}

func (o ActionResponsePtrOutput) Elem() ActionResponseOutput {
	return o.ApplyT(func(v *ActionResponse) ActionResponse {
		if v != nil {
			return *v
		}
		var ret ActionResponse
		return ret
	}).(ActionResponseOutput)
}

func (o ActionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
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
	Enabled    *bool    `pulumi:"enabled"`
	Expires    *float64 `pulumi:"expires"`
	Exportable *bool    `pulumi:"exportable"`
	NotBefore  *float64 `pulumi:"notBefore"`
}


func (val *KeyAttributes) Defaults() *KeyAttributes {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Exportable) {
		exportable_ := false
		tmp.Exportable = &exportable_
	}
	return &tmp
}





type KeyAttributesInput interface {
	pulumi.Input

	ToKeyAttributesOutput() KeyAttributesOutput
	ToKeyAttributesOutputWithContext(context.Context) KeyAttributesOutput
}

type KeyAttributesArgs struct {
	Enabled    pulumi.BoolPtrInput    `pulumi:"enabled"`
	Expires    pulumi.Float64PtrInput `pulumi:"expires"`
	Exportable pulumi.BoolPtrInput    `pulumi:"exportable"`
	NotBefore  pulumi.Float64PtrInput `pulumi:"notBefore"`
}


func (val *KeyAttributesArgs) Defaults() *KeyAttributesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Exportable) {
		tmp.Exportable = pulumi.BoolPtr(false)
	}
	return &tmp
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

func (o KeyAttributesOutput) Exportable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyAttributes) *bool { return v.Exportable }).(pulumi.BoolPtrOutput)
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

func (o KeyAttributesPtrOutput) Exportable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyAttributes) *bool {
		if v == nil {
			return nil
		}
		return v.Exportable
	}).(pulumi.BoolPtrOutput)
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
	Exportable    *bool    `pulumi:"exportable"`
	NotBefore     *float64 `pulumi:"notBefore"`
	RecoveryLevel string   `pulumi:"recoveryLevel"`
	Updated       float64  `pulumi:"updated"`
}


func (val *KeyAttributesResponse) Defaults() *KeyAttributesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Exportable) {
		exportable_ := false
		tmp.Exportable = &exportable_
	}
	return &tmp
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

func (o KeyAttributesResponseOutput) Created() pulumi.Float64Output {
	return o.ApplyT(func(v KeyAttributesResponse) float64 { return v.Created }).(pulumi.Float64Output)
}

func (o KeyAttributesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyAttributesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o KeyAttributesResponseOutput) Expires() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KeyAttributesResponse) *float64 { return v.Expires }).(pulumi.Float64PtrOutput)
}

func (o KeyAttributesResponseOutput) Exportable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyAttributesResponse) *bool { return v.Exportable }).(pulumi.BoolPtrOutput)
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

func (o KeyAttributesResponsePtrOutput) Exportable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyAttributesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Exportable
	}).(pulumi.BoolPtrOutput)
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
	Attributes     *KeyAttributes    `pulumi:"attributes"`
	CurveName      *string           `pulumi:"curveName"`
	KeyOps         []string          `pulumi:"keyOps"`
	KeySize        *int              `pulumi:"keySize"`
	Kty            *string           `pulumi:"kty"`
	ReleasePolicy  *KeyReleasePolicy `pulumi:"releasePolicy"`
	RotationPolicy *RotationPolicy   `pulumi:"rotationPolicy"`
}


func (val *KeyProperties) Defaults() *KeyProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Attributes = tmp.Attributes.Defaults()

	tmp.ReleasePolicy = tmp.ReleasePolicy.Defaults()

	return &tmp
}





type KeyPropertiesInput interface {
	pulumi.Input

	ToKeyPropertiesOutput() KeyPropertiesOutput
	ToKeyPropertiesOutputWithContext(context.Context) KeyPropertiesOutput
}

type KeyPropertiesArgs struct {
	Attributes     KeyAttributesPtrInput    `pulumi:"attributes"`
	CurveName      pulumi.StringPtrInput    `pulumi:"curveName"`
	KeyOps         pulumi.StringArrayInput  `pulumi:"keyOps"`
	KeySize        pulumi.IntPtrInput       `pulumi:"keySize"`
	Kty            pulumi.StringPtrInput    `pulumi:"kty"`
	ReleasePolicy  KeyReleasePolicyPtrInput `pulumi:"releasePolicy"`
	RotationPolicy RotationPolicyPtrInput   `pulumi:"rotationPolicy"`
}


func (val *KeyPropertiesArgs) Defaults() *KeyPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
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

func (o KeyPropertiesOutput) ReleasePolicy() KeyReleasePolicyPtrOutput {
	return o.ApplyT(func(v KeyProperties) *KeyReleasePolicy { return v.ReleasePolicy }).(KeyReleasePolicyPtrOutput)
}

func (o KeyPropertiesOutput) RotationPolicy() RotationPolicyPtrOutput {
	return o.ApplyT(func(v KeyProperties) *RotationPolicy { return v.RotationPolicy }).(RotationPolicyPtrOutput)
}

type KeyReleasePolicy struct {
	ContentType *string `pulumi:"contentType"`
	Data        *string `pulumi:"data"`
}


func (val *KeyReleasePolicy) Defaults() *KeyReleasePolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ContentType) {
		contentType_ := "application/json; charset=utf-8"
		tmp.ContentType = &contentType_
	}
	return &tmp
}





type KeyReleasePolicyInput interface {
	pulumi.Input

	ToKeyReleasePolicyOutput() KeyReleasePolicyOutput
	ToKeyReleasePolicyOutputWithContext(context.Context) KeyReleasePolicyOutput
}

type KeyReleasePolicyArgs struct {
	ContentType pulumi.StringPtrInput `pulumi:"contentType"`
	Data        pulumi.StringPtrInput `pulumi:"data"`
}


func (val *KeyReleasePolicyArgs) Defaults() *KeyReleasePolicyArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ContentType) {
		tmp.ContentType = pulumi.StringPtr("application/json; charset=utf-8")
	}
	return &tmp
}
func (KeyReleasePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyReleasePolicy)(nil)).Elem()
}

func (i KeyReleasePolicyArgs) ToKeyReleasePolicyOutput() KeyReleasePolicyOutput {
	return i.ToKeyReleasePolicyOutputWithContext(context.Background())
}

func (i KeyReleasePolicyArgs) ToKeyReleasePolicyOutputWithContext(ctx context.Context) KeyReleasePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyReleasePolicyOutput)
}

func (i KeyReleasePolicyArgs) ToKeyReleasePolicyPtrOutput() KeyReleasePolicyPtrOutput {
	return i.ToKeyReleasePolicyPtrOutputWithContext(context.Background())
}

func (i KeyReleasePolicyArgs) ToKeyReleasePolicyPtrOutputWithContext(ctx context.Context) KeyReleasePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyReleasePolicyOutput).ToKeyReleasePolicyPtrOutputWithContext(ctx)
}









type KeyReleasePolicyPtrInput interface {
	pulumi.Input

	ToKeyReleasePolicyPtrOutput() KeyReleasePolicyPtrOutput
	ToKeyReleasePolicyPtrOutputWithContext(context.Context) KeyReleasePolicyPtrOutput
}

type keyReleasePolicyPtrType KeyReleasePolicyArgs

func KeyReleasePolicyPtr(v *KeyReleasePolicyArgs) KeyReleasePolicyPtrInput {
	return (*keyReleasePolicyPtrType)(v)
}

func (*keyReleasePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyReleasePolicy)(nil)).Elem()
}

func (i *keyReleasePolicyPtrType) ToKeyReleasePolicyPtrOutput() KeyReleasePolicyPtrOutput {
	return i.ToKeyReleasePolicyPtrOutputWithContext(context.Background())
}

func (i *keyReleasePolicyPtrType) ToKeyReleasePolicyPtrOutputWithContext(ctx context.Context) KeyReleasePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyReleasePolicyPtrOutput)
}

type KeyReleasePolicyOutput struct{ *pulumi.OutputState }

func (KeyReleasePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyReleasePolicy)(nil)).Elem()
}

func (o KeyReleasePolicyOutput) ToKeyReleasePolicyOutput() KeyReleasePolicyOutput {
	return o
}

func (o KeyReleasePolicyOutput) ToKeyReleasePolicyOutputWithContext(ctx context.Context) KeyReleasePolicyOutput {
	return o
}

func (o KeyReleasePolicyOutput) ToKeyReleasePolicyPtrOutput() KeyReleasePolicyPtrOutput {
	return o.ToKeyReleasePolicyPtrOutputWithContext(context.Background())
}

func (o KeyReleasePolicyOutput) ToKeyReleasePolicyPtrOutputWithContext(ctx context.Context) KeyReleasePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyReleasePolicy) *KeyReleasePolicy {
		return &v
	}).(KeyReleasePolicyPtrOutput)
}

func (o KeyReleasePolicyOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyReleasePolicy) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o KeyReleasePolicyOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyReleasePolicy) *string { return v.Data }).(pulumi.StringPtrOutput)
}

type KeyReleasePolicyPtrOutput struct{ *pulumi.OutputState }

func (KeyReleasePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyReleasePolicy)(nil)).Elem()
}

func (o KeyReleasePolicyPtrOutput) ToKeyReleasePolicyPtrOutput() KeyReleasePolicyPtrOutput {
	return o
}

func (o KeyReleasePolicyPtrOutput) ToKeyReleasePolicyPtrOutputWithContext(ctx context.Context) KeyReleasePolicyPtrOutput {
	return o
}

func (o KeyReleasePolicyPtrOutput) Elem() KeyReleasePolicyOutput {
	return o.ApplyT(func(v *KeyReleasePolicy) KeyReleasePolicy {
		if v != nil {
			return *v
		}
		var ret KeyReleasePolicy
		return ret
	}).(KeyReleasePolicyOutput)
}

func (o KeyReleasePolicyPtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyReleasePolicy) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o KeyReleasePolicyPtrOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyReleasePolicy) *string {
		if v == nil {
			return nil
		}
		return v.Data
	}).(pulumi.StringPtrOutput)
}

type KeyReleasePolicyResponse struct {
	ContentType *string `pulumi:"contentType"`
	Data        *string `pulumi:"data"`
}


func (val *KeyReleasePolicyResponse) Defaults() *KeyReleasePolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ContentType) {
		contentType_ := "application/json; charset=utf-8"
		tmp.ContentType = &contentType_
	}
	return &tmp
}

type KeyReleasePolicyResponseOutput struct{ *pulumi.OutputState }

func (KeyReleasePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyReleasePolicyResponse)(nil)).Elem()
}

func (o KeyReleasePolicyResponseOutput) ToKeyReleasePolicyResponseOutput() KeyReleasePolicyResponseOutput {
	return o
}

func (o KeyReleasePolicyResponseOutput) ToKeyReleasePolicyResponseOutputWithContext(ctx context.Context) KeyReleasePolicyResponseOutput {
	return o
}

func (o KeyReleasePolicyResponseOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyReleasePolicyResponse) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o KeyReleasePolicyResponseOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyReleasePolicyResponse) *string { return v.Data }).(pulumi.StringPtrOutput)
}

type KeyReleasePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyReleasePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyReleasePolicyResponse)(nil)).Elem()
}

func (o KeyReleasePolicyResponsePtrOutput) ToKeyReleasePolicyResponsePtrOutput() KeyReleasePolicyResponsePtrOutput {
	return o
}

func (o KeyReleasePolicyResponsePtrOutput) ToKeyReleasePolicyResponsePtrOutputWithContext(ctx context.Context) KeyReleasePolicyResponsePtrOutput {
	return o
}

func (o KeyReleasePolicyResponsePtrOutput) Elem() KeyReleasePolicyResponseOutput {
	return o.ApplyT(func(v *KeyReleasePolicyResponse) KeyReleasePolicyResponse {
		if v != nil {
			return *v
		}
		var ret KeyReleasePolicyResponse
		return ret
	}).(KeyReleasePolicyResponseOutput)
}

func (o KeyReleasePolicyResponsePtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyReleasePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o KeyReleasePolicyResponsePtrOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyReleasePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Data
	}).(pulumi.StringPtrOutput)
}

type KeyRotationPolicyAttributes struct {
	ExpiryTime *string `pulumi:"expiryTime"`
}





type KeyRotationPolicyAttributesInput interface {
	pulumi.Input

	ToKeyRotationPolicyAttributesOutput() KeyRotationPolicyAttributesOutput
	ToKeyRotationPolicyAttributesOutputWithContext(context.Context) KeyRotationPolicyAttributesOutput
}

type KeyRotationPolicyAttributesArgs struct {
	ExpiryTime pulumi.StringPtrInput `pulumi:"expiryTime"`
}

func (KeyRotationPolicyAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyRotationPolicyAttributes)(nil)).Elem()
}

func (i KeyRotationPolicyAttributesArgs) ToKeyRotationPolicyAttributesOutput() KeyRotationPolicyAttributesOutput {
	return i.ToKeyRotationPolicyAttributesOutputWithContext(context.Background())
}

func (i KeyRotationPolicyAttributesArgs) ToKeyRotationPolicyAttributesOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRotationPolicyAttributesOutput)
}

func (i KeyRotationPolicyAttributesArgs) ToKeyRotationPolicyAttributesPtrOutput() KeyRotationPolicyAttributesPtrOutput {
	return i.ToKeyRotationPolicyAttributesPtrOutputWithContext(context.Background())
}

func (i KeyRotationPolicyAttributesArgs) ToKeyRotationPolicyAttributesPtrOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRotationPolicyAttributesOutput).ToKeyRotationPolicyAttributesPtrOutputWithContext(ctx)
}









type KeyRotationPolicyAttributesPtrInput interface {
	pulumi.Input

	ToKeyRotationPolicyAttributesPtrOutput() KeyRotationPolicyAttributesPtrOutput
	ToKeyRotationPolicyAttributesPtrOutputWithContext(context.Context) KeyRotationPolicyAttributesPtrOutput
}

type keyRotationPolicyAttributesPtrType KeyRotationPolicyAttributesArgs

func KeyRotationPolicyAttributesPtr(v *KeyRotationPolicyAttributesArgs) KeyRotationPolicyAttributesPtrInput {
	return (*keyRotationPolicyAttributesPtrType)(v)
}

func (*keyRotationPolicyAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRotationPolicyAttributes)(nil)).Elem()
}

func (i *keyRotationPolicyAttributesPtrType) ToKeyRotationPolicyAttributesPtrOutput() KeyRotationPolicyAttributesPtrOutput {
	return i.ToKeyRotationPolicyAttributesPtrOutputWithContext(context.Background())
}

func (i *keyRotationPolicyAttributesPtrType) ToKeyRotationPolicyAttributesPtrOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRotationPolicyAttributesPtrOutput)
}

type KeyRotationPolicyAttributesOutput struct{ *pulumi.OutputState }

func (KeyRotationPolicyAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyRotationPolicyAttributes)(nil)).Elem()
}

func (o KeyRotationPolicyAttributesOutput) ToKeyRotationPolicyAttributesOutput() KeyRotationPolicyAttributesOutput {
	return o
}

func (o KeyRotationPolicyAttributesOutput) ToKeyRotationPolicyAttributesOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesOutput {
	return o
}

func (o KeyRotationPolicyAttributesOutput) ToKeyRotationPolicyAttributesPtrOutput() KeyRotationPolicyAttributesPtrOutput {
	return o.ToKeyRotationPolicyAttributesPtrOutputWithContext(context.Background())
}

func (o KeyRotationPolicyAttributesOutput) ToKeyRotationPolicyAttributesPtrOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyRotationPolicyAttributes) *KeyRotationPolicyAttributes {
		return &v
	}).(KeyRotationPolicyAttributesPtrOutput)
}

func (o KeyRotationPolicyAttributesOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyRotationPolicyAttributes) *string { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

type KeyRotationPolicyAttributesPtrOutput struct{ *pulumi.OutputState }

func (KeyRotationPolicyAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRotationPolicyAttributes)(nil)).Elem()
}

func (o KeyRotationPolicyAttributesPtrOutput) ToKeyRotationPolicyAttributesPtrOutput() KeyRotationPolicyAttributesPtrOutput {
	return o
}

func (o KeyRotationPolicyAttributesPtrOutput) ToKeyRotationPolicyAttributesPtrOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesPtrOutput {
	return o
}

func (o KeyRotationPolicyAttributesPtrOutput) Elem() KeyRotationPolicyAttributesOutput {
	return o.ApplyT(func(v *KeyRotationPolicyAttributes) KeyRotationPolicyAttributes {
		if v != nil {
			return *v
		}
		var ret KeyRotationPolicyAttributes
		return ret
	}).(KeyRotationPolicyAttributesOutput)
}

func (o KeyRotationPolicyAttributesPtrOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyRotationPolicyAttributes) *string {
		if v == nil {
			return nil
		}
		return v.ExpiryTime
	}).(pulumi.StringPtrOutput)
}

type KeyRotationPolicyAttributesResponse struct {
	Created    float64 `pulumi:"created"`
	ExpiryTime *string `pulumi:"expiryTime"`
	Updated    float64 `pulumi:"updated"`
}

type KeyRotationPolicyAttributesResponseOutput struct{ *pulumi.OutputState }

func (KeyRotationPolicyAttributesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyRotationPolicyAttributesResponse)(nil)).Elem()
}

func (o KeyRotationPolicyAttributesResponseOutput) ToKeyRotationPolicyAttributesResponseOutput() KeyRotationPolicyAttributesResponseOutput {
	return o
}

func (o KeyRotationPolicyAttributesResponseOutput) ToKeyRotationPolicyAttributesResponseOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesResponseOutput {
	return o
}

func (o KeyRotationPolicyAttributesResponseOutput) Created() pulumi.Float64Output {
	return o.ApplyT(func(v KeyRotationPolicyAttributesResponse) float64 { return v.Created }).(pulumi.Float64Output)
}

func (o KeyRotationPolicyAttributesResponseOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyRotationPolicyAttributesResponse) *string { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

func (o KeyRotationPolicyAttributesResponseOutput) Updated() pulumi.Float64Output {
	return o.ApplyT(func(v KeyRotationPolicyAttributesResponse) float64 { return v.Updated }).(pulumi.Float64Output)
}

type KeyRotationPolicyAttributesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyRotationPolicyAttributesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRotationPolicyAttributesResponse)(nil)).Elem()
}

func (o KeyRotationPolicyAttributesResponsePtrOutput) ToKeyRotationPolicyAttributesResponsePtrOutput() KeyRotationPolicyAttributesResponsePtrOutput {
	return o
}

func (o KeyRotationPolicyAttributesResponsePtrOutput) ToKeyRotationPolicyAttributesResponsePtrOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesResponsePtrOutput {
	return o
}

func (o KeyRotationPolicyAttributesResponsePtrOutput) Elem() KeyRotationPolicyAttributesResponseOutput {
	return o.ApplyT(func(v *KeyRotationPolicyAttributesResponse) KeyRotationPolicyAttributesResponse {
		if v != nil {
			return *v
		}
		var ret KeyRotationPolicyAttributesResponse
		return ret
	}).(KeyRotationPolicyAttributesResponseOutput)
}

func (o KeyRotationPolicyAttributesResponsePtrOutput) Created() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyRotationPolicyAttributesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Created
	}).(pulumi.Float64PtrOutput)
}

func (o KeyRotationPolicyAttributesResponsePtrOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyRotationPolicyAttributesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpiryTime
	}).(pulumi.StringPtrOutput)
}

func (o KeyRotationPolicyAttributesResponsePtrOutput) Updated() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyRotationPolicyAttributesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Updated
	}).(pulumi.Float64PtrOutput)
}

type LifetimeAction struct {
	Action  *Action  `pulumi:"action"`
	Trigger *Trigger `pulumi:"trigger"`
}





type LifetimeActionInput interface {
	pulumi.Input

	ToLifetimeActionOutput() LifetimeActionOutput
	ToLifetimeActionOutputWithContext(context.Context) LifetimeActionOutput
}

type LifetimeActionArgs struct {
	Action  ActionPtrInput  `pulumi:"action"`
	Trigger TriggerPtrInput `pulumi:"trigger"`
}

func (LifetimeActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LifetimeAction)(nil)).Elem()
}

func (i LifetimeActionArgs) ToLifetimeActionOutput() LifetimeActionOutput {
	return i.ToLifetimeActionOutputWithContext(context.Background())
}

func (i LifetimeActionArgs) ToLifetimeActionOutputWithContext(ctx context.Context) LifetimeActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifetimeActionOutput)
}





type LifetimeActionArrayInput interface {
	pulumi.Input

	ToLifetimeActionArrayOutput() LifetimeActionArrayOutput
	ToLifetimeActionArrayOutputWithContext(context.Context) LifetimeActionArrayOutput
}

type LifetimeActionArray []LifetimeActionInput

func (LifetimeActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LifetimeAction)(nil)).Elem()
}

func (i LifetimeActionArray) ToLifetimeActionArrayOutput() LifetimeActionArrayOutput {
	return i.ToLifetimeActionArrayOutputWithContext(context.Background())
}

func (i LifetimeActionArray) ToLifetimeActionArrayOutputWithContext(ctx context.Context) LifetimeActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifetimeActionArrayOutput)
}

type LifetimeActionOutput struct{ *pulumi.OutputState }

func (LifetimeActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LifetimeAction)(nil)).Elem()
}

func (o LifetimeActionOutput) ToLifetimeActionOutput() LifetimeActionOutput {
	return o
}

func (o LifetimeActionOutput) ToLifetimeActionOutputWithContext(ctx context.Context) LifetimeActionOutput {
	return o
}

func (o LifetimeActionOutput) Action() ActionPtrOutput {
	return o.ApplyT(func(v LifetimeAction) *Action { return v.Action }).(ActionPtrOutput)
}

func (o LifetimeActionOutput) Trigger() TriggerPtrOutput {
	return o.ApplyT(func(v LifetimeAction) *Trigger { return v.Trigger }).(TriggerPtrOutput)
}

type LifetimeActionArrayOutput struct{ *pulumi.OutputState }

func (LifetimeActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LifetimeAction)(nil)).Elem()
}

func (o LifetimeActionArrayOutput) ToLifetimeActionArrayOutput() LifetimeActionArrayOutput {
	return o
}

func (o LifetimeActionArrayOutput) ToLifetimeActionArrayOutputWithContext(ctx context.Context) LifetimeActionArrayOutput {
	return o
}

func (o LifetimeActionArrayOutput) Index(i pulumi.IntInput) LifetimeActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LifetimeAction {
		return vs[0].([]LifetimeAction)[vs[1].(int)]
	}).(LifetimeActionOutput)
}

type LifetimeActionResponse struct {
	Action  *ActionResponse  `pulumi:"action"`
	Trigger *TriggerResponse `pulumi:"trigger"`
}

type LifetimeActionResponseOutput struct{ *pulumi.OutputState }

func (LifetimeActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LifetimeActionResponse)(nil)).Elem()
}

func (o LifetimeActionResponseOutput) ToLifetimeActionResponseOutput() LifetimeActionResponseOutput {
	return o
}

func (o LifetimeActionResponseOutput) ToLifetimeActionResponseOutputWithContext(ctx context.Context) LifetimeActionResponseOutput {
	return o
}

func (o LifetimeActionResponseOutput) Action() ActionResponsePtrOutput {
	return o.ApplyT(func(v LifetimeActionResponse) *ActionResponse { return v.Action }).(ActionResponsePtrOutput)
}

func (o LifetimeActionResponseOutput) Trigger() TriggerResponsePtrOutput {
	return o.ApplyT(func(v LifetimeActionResponse) *TriggerResponse { return v.Trigger }).(TriggerResponsePtrOutput)
}

type LifetimeActionResponseArrayOutput struct{ *pulumi.OutputState }

func (LifetimeActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LifetimeActionResponse)(nil)).Elem()
}

func (o LifetimeActionResponseArrayOutput) ToLifetimeActionResponseArrayOutput() LifetimeActionResponseArrayOutput {
	return o
}

func (o LifetimeActionResponseArrayOutput) ToLifetimeActionResponseArrayOutputWithContext(ctx context.Context) LifetimeActionResponseArrayOutput {
	return o
}

func (o LifetimeActionResponseArrayOutput) Index(i pulumi.IntInput) LifetimeActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LifetimeActionResponse {
		return vs[0].([]LifetimeActionResponse)[vs[1].(int)]
	}).(LifetimeActionResponseOutput)
}

type MHSMIPRule struct {
	Value string `pulumi:"value"`
}





type MHSMIPRuleInput interface {
	pulumi.Input

	ToMHSMIPRuleOutput() MHSMIPRuleOutput
	ToMHSMIPRuleOutputWithContext(context.Context) MHSMIPRuleOutput
}

type MHSMIPRuleArgs struct {
	Value pulumi.StringInput `pulumi:"value"`
}

func (MHSMIPRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMIPRule)(nil)).Elem()
}

func (i MHSMIPRuleArgs) ToMHSMIPRuleOutput() MHSMIPRuleOutput {
	return i.ToMHSMIPRuleOutputWithContext(context.Background())
}

func (i MHSMIPRuleArgs) ToMHSMIPRuleOutputWithContext(ctx context.Context) MHSMIPRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMIPRuleOutput)
}





type MHSMIPRuleArrayInput interface {
	pulumi.Input

	ToMHSMIPRuleArrayOutput() MHSMIPRuleArrayOutput
	ToMHSMIPRuleArrayOutputWithContext(context.Context) MHSMIPRuleArrayOutput
}

type MHSMIPRuleArray []MHSMIPRuleInput

func (MHSMIPRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMIPRule)(nil)).Elem()
}

func (i MHSMIPRuleArray) ToMHSMIPRuleArrayOutput() MHSMIPRuleArrayOutput {
	return i.ToMHSMIPRuleArrayOutputWithContext(context.Background())
}

func (i MHSMIPRuleArray) ToMHSMIPRuleArrayOutputWithContext(ctx context.Context) MHSMIPRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMIPRuleArrayOutput)
}

type MHSMIPRuleOutput struct{ *pulumi.OutputState }

func (MHSMIPRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMIPRule)(nil)).Elem()
}

func (o MHSMIPRuleOutput) ToMHSMIPRuleOutput() MHSMIPRuleOutput {
	return o
}

func (o MHSMIPRuleOutput) ToMHSMIPRuleOutputWithContext(ctx context.Context) MHSMIPRuleOutput {
	return o
}

func (o MHSMIPRuleOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v MHSMIPRule) string { return v.Value }).(pulumi.StringOutput)
}

type MHSMIPRuleArrayOutput struct{ *pulumi.OutputState }

func (MHSMIPRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMIPRule)(nil)).Elem()
}

func (o MHSMIPRuleArrayOutput) ToMHSMIPRuleArrayOutput() MHSMIPRuleArrayOutput {
	return o
}

func (o MHSMIPRuleArrayOutput) ToMHSMIPRuleArrayOutputWithContext(ctx context.Context) MHSMIPRuleArrayOutput {
	return o
}

func (o MHSMIPRuleArrayOutput) Index(i pulumi.IntInput) MHSMIPRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MHSMIPRule {
		return vs[0].([]MHSMIPRule)[vs[1].(int)]
	}).(MHSMIPRuleOutput)
}

type MHSMIPRuleResponse struct {
	Value string `pulumi:"value"`
}

type MHSMIPRuleResponseOutput struct{ *pulumi.OutputState }

func (MHSMIPRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMIPRuleResponse)(nil)).Elem()
}

func (o MHSMIPRuleResponseOutput) ToMHSMIPRuleResponseOutput() MHSMIPRuleResponseOutput {
	return o
}

func (o MHSMIPRuleResponseOutput) ToMHSMIPRuleResponseOutputWithContext(ctx context.Context) MHSMIPRuleResponseOutput {
	return o
}

func (o MHSMIPRuleResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v MHSMIPRuleResponse) string { return v.Value }).(pulumi.StringOutput)
}

type MHSMIPRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (MHSMIPRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMIPRuleResponse)(nil)).Elem()
}

func (o MHSMIPRuleResponseArrayOutput) ToMHSMIPRuleResponseArrayOutput() MHSMIPRuleResponseArrayOutput {
	return o
}

func (o MHSMIPRuleResponseArrayOutput) ToMHSMIPRuleResponseArrayOutputWithContext(ctx context.Context) MHSMIPRuleResponseArrayOutput {
	return o
}

func (o MHSMIPRuleResponseArrayOutput) Index(i pulumi.IntInput) MHSMIPRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MHSMIPRuleResponse {
		return vs[0].([]MHSMIPRuleResponse)[vs[1].(int)]
	}).(MHSMIPRuleResponseOutput)
}

type MHSMNetworkRuleSet struct {
	Bypass              *string                  `pulumi:"bypass"`
	DefaultAction       *string                  `pulumi:"defaultAction"`
	IpRules             []MHSMIPRule             `pulumi:"ipRules"`
	VirtualNetworkRules []MHSMVirtualNetworkRule `pulumi:"virtualNetworkRules"`
}





type MHSMNetworkRuleSetInput interface {
	pulumi.Input

	ToMHSMNetworkRuleSetOutput() MHSMNetworkRuleSetOutput
	ToMHSMNetworkRuleSetOutputWithContext(context.Context) MHSMNetworkRuleSetOutput
}

type MHSMNetworkRuleSetArgs struct {
	Bypass              pulumi.StringPtrInput            `pulumi:"bypass"`
	DefaultAction       pulumi.StringPtrInput            `pulumi:"defaultAction"`
	IpRules             MHSMIPRuleArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules MHSMVirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (MHSMNetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMNetworkRuleSet)(nil)).Elem()
}

func (i MHSMNetworkRuleSetArgs) ToMHSMNetworkRuleSetOutput() MHSMNetworkRuleSetOutput {
	return i.ToMHSMNetworkRuleSetOutputWithContext(context.Background())
}

func (i MHSMNetworkRuleSetArgs) ToMHSMNetworkRuleSetOutputWithContext(ctx context.Context) MHSMNetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMNetworkRuleSetOutput)
}

func (i MHSMNetworkRuleSetArgs) ToMHSMNetworkRuleSetPtrOutput() MHSMNetworkRuleSetPtrOutput {
	return i.ToMHSMNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i MHSMNetworkRuleSetArgs) ToMHSMNetworkRuleSetPtrOutputWithContext(ctx context.Context) MHSMNetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMNetworkRuleSetOutput).ToMHSMNetworkRuleSetPtrOutputWithContext(ctx)
}









type MHSMNetworkRuleSetPtrInput interface {
	pulumi.Input

	ToMHSMNetworkRuleSetPtrOutput() MHSMNetworkRuleSetPtrOutput
	ToMHSMNetworkRuleSetPtrOutputWithContext(context.Context) MHSMNetworkRuleSetPtrOutput
}

type mhsmnetworkRuleSetPtrType MHSMNetworkRuleSetArgs

func MHSMNetworkRuleSetPtr(v *MHSMNetworkRuleSetArgs) MHSMNetworkRuleSetPtrInput {
	return (*mhsmnetworkRuleSetPtrType)(v)
}

func (*mhsmnetworkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMNetworkRuleSet)(nil)).Elem()
}

func (i *mhsmnetworkRuleSetPtrType) ToMHSMNetworkRuleSetPtrOutput() MHSMNetworkRuleSetPtrOutput {
	return i.ToMHSMNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *mhsmnetworkRuleSetPtrType) ToMHSMNetworkRuleSetPtrOutputWithContext(ctx context.Context) MHSMNetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMNetworkRuleSetPtrOutput)
}

type MHSMNetworkRuleSetOutput struct{ *pulumi.OutputState }

func (MHSMNetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMNetworkRuleSet)(nil)).Elem()
}

func (o MHSMNetworkRuleSetOutput) ToMHSMNetworkRuleSetOutput() MHSMNetworkRuleSetOutput {
	return o
}

func (o MHSMNetworkRuleSetOutput) ToMHSMNetworkRuleSetOutputWithContext(ctx context.Context) MHSMNetworkRuleSetOutput {
	return o
}

func (o MHSMNetworkRuleSetOutput) ToMHSMNetworkRuleSetPtrOutput() MHSMNetworkRuleSetPtrOutput {
	return o.ToMHSMNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o MHSMNetworkRuleSetOutput) ToMHSMNetworkRuleSetPtrOutputWithContext(ctx context.Context) MHSMNetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MHSMNetworkRuleSet) *MHSMNetworkRuleSet {
		return &v
	}).(MHSMNetworkRuleSetPtrOutput)
}

func (o MHSMNetworkRuleSetOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMNetworkRuleSet) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o MHSMNetworkRuleSetOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMNetworkRuleSet) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o MHSMNetworkRuleSetOutput) IpRules() MHSMIPRuleArrayOutput {
	return o.ApplyT(func(v MHSMNetworkRuleSet) []MHSMIPRule { return v.IpRules }).(MHSMIPRuleArrayOutput)
}

func (o MHSMNetworkRuleSetOutput) VirtualNetworkRules() MHSMVirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v MHSMNetworkRuleSet) []MHSMVirtualNetworkRule { return v.VirtualNetworkRules }).(MHSMVirtualNetworkRuleArrayOutput)
}

type MHSMNetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (MHSMNetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMNetworkRuleSet)(nil)).Elem()
}

func (o MHSMNetworkRuleSetPtrOutput) ToMHSMNetworkRuleSetPtrOutput() MHSMNetworkRuleSetPtrOutput {
	return o
}

func (o MHSMNetworkRuleSetPtrOutput) ToMHSMNetworkRuleSetPtrOutputWithContext(ctx context.Context) MHSMNetworkRuleSetPtrOutput {
	return o
}

func (o MHSMNetworkRuleSetPtrOutput) Elem() MHSMNetworkRuleSetOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSet) MHSMNetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret MHSMNetworkRuleSet
		return ret
	}).(MHSMNetworkRuleSetOutput)
}

func (o MHSMNetworkRuleSetPtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o MHSMNetworkRuleSetPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o MHSMNetworkRuleSetPtrOutput) IpRules() MHSMIPRuleArrayOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSet) []MHSMIPRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(MHSMIPRuleArrayOutput)
}

func (o MHSMNetworkRuleSetPtrOutput) VirtualNetworkRules() MHSMVirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSet) []MHSMVirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(MHSMVirtualNetworkRuleArrayOutput)
}

type MHSMNetworkRuleSetResponse struct {
	Bypass              *string                          `pulumi:"bypass"`
	DefaultAction       *string                          `pulumi:"defaultAction"`
	IpRules             []MHSMIPRuleResponse             `pulumi:"ipRules"`
	VirtualNetworkRules []MHSMVirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}

type MHSMNetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (MHSMNetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMNetworkRuleSetResponse)(nil)).Elem()
}

func (o MHSMNetworkRuleSetResponseOutput) ToMHSMNetworkRuleSetResponseOutput() MHSMNetworkRuleSetResponseOutput {
	return o
}

func (o MHSMNetworkRuleSetResponseOutput) ToMHSMNetworkRuleSetResponseOutputWithContext(ctx context.Context) MHSMNetworkRuleSetResponseOutput {
	return o
}

func (o MHSMNetworkRuleSetResponseOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMNetworkRuleSetResponse) *string { return v.Bypass }).(pulumi.StringPtrOutput)
}

func (o MHSMNetworkRuleSetResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMNetworkRuleSetResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o MHSMNetworkRuleSetResponseOutput) IpRules() MHSMIPRuleResponseArrayOutput {
	return o.ApplyT(func(v MHSMNetworkRuleSetResponse) []MHSMIPRuleResponse { return v.IpRules }).(MHSMIPRuleResponseArrayOutput)
}

func (o MHSMNetworkRuleSetResponseOutput) VirtualNetworkRules() MHSMVirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v MHSMNetworkRuleSetResponse) []MHSMVirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(MHSMVirtualNetworkRuleResponseArrayOutput)
}

type MHSMNetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (MHSMNetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMNetworkRuleSetResponse)(nil)).Elem()
}

func (o MHSMNetworkRuleSetResponsePtrOutput) ToMHSMNetworkRuleSetResponsePtrOutput() MHSMNetworkRuleSetResponsePtrOutput {
	return o
}

func (o MHSMNetworkRuleSetResponsePtrOutput) ToMHSMNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) MHSMNetworkRuleSetResponsePtrOutput {
	return o
}

func (o MHSMNetworkRuleSetResponsePtrOutput) Elem() MHSMNetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSetResponse) MHSMNetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret MHSMNetworkRuleSetResponse
		return ret
	}).(MHSMNetworkRuleSetResponseOutput)
}

func (o MHSMNetworkRuleSetResponsePtrOutput) Bypass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Bypass
	}).(pulumi.StringPtrOutput)
}

func (o MHSMNetworkRuleSetResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o MHSMNetworkRuleSetResponsePtrOutput) IpRules() MHSMIPRuleResponseArrayOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSetResponse) []MHSMIPRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(MHSMIPRuleResponseArrayOutput)
}

func (o MHSMNetworkRuleSetResponsePtrOutput) VirtualNetworkRules() MHSMVirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *MHSMNetworkRuleSetResponse) []MHSMVirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(MHSMVirtualNetworkRuleResponseArrayOutput)
}

type MHSMPrivateEndpointConnectionItemResponse struct {
	Etag                              *string                                        `pulumi:"etag"`
	Id                                *string                                        `pulumi:"id"`
	PrivateEndpoint                   *MHSMPrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *MHSMPrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                         `pulumi:"provisioningState"`
}

type MHSMPrivateEndpointConnectionItemResponseOutput struct{ *pulumi.OutputState }

func (MHSMPrivateEndpointConnectionItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateEndpointConnectionItemResponse)(nil)).Elem()
}

func (o MHSMPrivateEndpointConnectionItemResponseOutput) ToMHSMPrivateEndpointConnectionItemResponseOutput() MHSMPrivateEndpointConnectionItemResponseOutput {
	return o
}

func (o MHSMPrivateEndpointConnectionItemResponseOutput) ToMHSMPrivateEndpointConnectionItemResponseOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionItemResponseOutput {
	return o
}

func (o MHSMPrivateEndpointConnectionItemResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMPrivateEndpointConnectionItemResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateEndpointConnectionItemResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMPrivateEndpointConnectionItemResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateEndpointConnectionItemResponseOutput) PrivateEndpoint() MHSMPrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v MHSMPrivateEndpointConnectionItemResponse) *MHSMPrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(MHSMPrivateEndpointResponsePtrOutput)
}

func (o MHSMPrivateEndpointConnectionItemResponseOutput) PrivateLinkServiceConnectionState() MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v MHSMPrivateEndpointConnectionItemResponse) *MHSMPrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(MHSMPrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o MHSMPrivateEndpointConnectionItemResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MHSMPrivateEndpointConnectionItemResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type MHSMPrivateEndpointConnectionItemResponseArrayOutput struct{ *pulumi.OutputState }

func (MHSMPrivateEndpointConnectionItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMPrivateEndpointConnectionItemResponse)(nil)).Elem()
}

func (o MHSMPrivateEndpointConnectionItemResponseArrayOutput) ToMHSMPrivateEndpointConnectionItemResponseArrayOutput() MHSMPrivateEndpointConnectionItemResponseArrayOutput {
	return o
}

func (o MHSMPrivateEndpointConnectionItemResponseArrayOutput) ToMHSMPrivateEndpointConnectionItemResponseArrayOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionItemResponseArrayOutput {
	return o
}

func (o MHSMPrivateEndpointConnectionItemResponseArrayOutput) Index(i pulumi.IntInput) MHSMPrivateEndpointConnectionItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MHSMPrivateEndpointConnectionItemResponse {
		return vs[0].([]MHSMPrivateEndpointConnectionItemResponse)[vs[1].(int)]
	}).(MHSMPrivateEndpointConnectionItemResponseOutput)
}

type MHSMPrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type MHSMPrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (MHSMPrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateEndpointResponse)(nil)).Elem()
}

func (o MHSMPrivateEndpointResponseOutput) ToMHSMPrivateEndpointResponseOutput() MHSMPrivateEndpointResponseOutput {
	return o
}

func (o MHSMPrivateEndpointResponseOutput) ToMHSMPrivateEndpointResponseOutputWithContext(ctx context.Context) MHSMPrivateEndpointResponseOutput {
	return o
}

func (o MHSMPrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MHSMPrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type MHSMPrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (MHSMPrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMPrivateEndpointResponse)(nil)).Elem()
}

func (o MHSMPrivateEndpointResponsePtrOutput) ToMHSMPrivateEndpointResponsePtrOutput() MHSMPrivateEndpointResponsePtrOutput {
	return o
}

func (o MHSMPrivateEndpointResponsePtrOutput) ToMHSMPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) MHSMPrivateEndpointResponsePtrOutput {
	return o
}

func (o MHSMPrivateEndpointResponsePtrOutput) Elem() MHSMPrivateEndpointResponseOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointResponse) MHSMPrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret MHSMPrivateEndpointResponse
		return ret
	}).(MHSMPrivateEndpointResponseOutput)
}

func (o MHSMPrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type MHSMPrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type MHSMPrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToMHSMPrivateLinkServiceConnectionStateOutput() MHSMPrivateLinkServiceConnectionStateOutput
	ToMHSMPrivateLinkServiceConnectionStateOutputWithContext(context.Context) MHSMPrivateLinkServiceConnectionStateOutput
}

type MHSMPrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (MHSMPrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i MHSMPrivateLinkServiceConnectionStateArgs) ToMHSMPrivateLinkServiceConnectionStateOutput() MHSMPrivateLinkServiceConnectionStateOutput {
	return i.ToMHSMPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i MHSMPrivateLinkServiceConnectionStateArgs) ToMHSMPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateLinkServiceConnectionStateOutput)
}

func (i MHSMPrivateLinkServiceConnectionStateArgs) ToMHSMPrivateLinkServiceConnectionStatePtrOutput() MHSMPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToMHSMPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i MHSMPrivateLinkServiceConnectionStateArgs) ToMHSMPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateLinkServiceConnectionStateOutput).ToMHSMPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type MHSMPrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToMHSMPrivateLinkServiceConnectionStatePtrOutput() MHSMPrivateLinkServiceConnectionStatePtrOutput
	ToMHSMPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) MHSMPrivateLinkServiceConnectionStatePtrOutput
}

type mhsmprivateLinkServiceConnectionStatePtrType MHSMPrivateLinkServiceConnectionStateArgs

func MHSMPrivateLinkServiceConnectionStatePtr(v *MHSMPrivateLinkServiceConnectionStateArgs) MHSMPrivateLinkServiceConnectionStatePtrInput {
	return (*mhsmprivateLinkServiceConnectionStatePtrType)(v)
}

func (*mhsmprivateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *mhsmprivateLinkServiceConnectionStatePtrType) ToMHSMPrivateLinkServiceConnectionStatePtrOutput() MHSMPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToMHSMPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *mhsmprivateLinkServiceConnectionStatePtrType) ToMHSMPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateLinkServiceConnectionStatePtrOutput)
}

type MHSMPrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (MHSMPrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o MHSMPrivateLinkServiceConnectionStateOutput) ToMHSMPrivateLinkServiceConnectionStateOutput() MHSMPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o MHSMPrivateLinkServiceConnectionStateOutput) ToMHSMPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o MHSMPrivateLinkServiceConnectionStateOutput) ToMHSMPrivateLinkServiceConnectionStatePtrOutput() MHSMPrivateLinkServiceConnectionStatePtrOutput {
	return o.ToMHSMPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o MHSMPrivateLinkServiceConnectionStateOutput) ToMHSMPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MHSMPrivateLinkServiceConnectionState) *MHSMPrivateLinkServiceConnectionState {
		return &v
	}).(MHSMPrivateLinkServiceConnectionStatePtrOutput)
}

func (o MHSMPrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMPrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMPrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMPrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type MHSMPrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (MHSMPrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o MHSMPrivateLinkServiceConnectionStatePtrOutput) ToMHSMPrivateLinkServiceConnectionStatePtrOutput() MHSMPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o MHSMPrivateLinkServiceConnectionStatePtrOutput) ToMHSMPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o MHSMPrivateLinkServiceConnectionStatePtrOutput) Elem() MHSMPrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *MHSMPrivateLinkServiceConnectionState) MHSMPrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret MHSMPrivateLinkServiceConnectionState
		return ret
	}).(MHSMPrivateLinkServiceConnectionStateOutput)
}

func (o MHSMPrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type MHSMPrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type MHSMPrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (MHSMPrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o MHSMPrivateLinkServiceConnectionStateResponseOutput) ToMHSMPrivateLinkServiceConnectionStateResponseOutput() MHSMPrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o MHSMPrivateLinkServiceConnectionStateResponseOutput) ToMHSMPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o MHSMPrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMPrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMPrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MHSMPrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type MHSMPrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (MHSMPrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMPrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o MHSMPrivateLinkServiceConnectionStateResponsePtrOutput) ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutput() MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o MHSMPrivateLinkServiceConnectionStateResponsePtrOutput) ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o MHSMPrivateLinkServiceConnectionStateResponsePtrOutput) Elem() MHSMPrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *MHSMPrivateLinkServiceConnectionStateResponse) MHSMPrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret MHSMPrivateLinkServiceConnectionStateResponse
		return ret
	}).(MHSMPrivateLinkServiceConnectionStateResponseOutput)
}

func (o MHSMPrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMPrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMPrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o MHSMPrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMPrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type MHSMVirtualNetworkRule struct {
	Id string `pulumi:"id"`
}





type MHSMVirtualNetworkRuleInput interface {
	pulumi.Input

	ToMHSMVirtualNetworkRuleOutput() MHSMVirtualNetworkRuleOutput
	ToMHSMVirtualNetworkRuleOutputWithContext(context.Context) MHSMVirtualNetworkRuleOutput
}

type MHSMVirtualNetworkRuleArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (MHSMVirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMVirtualNetworkRule)(nil)).Elem()
}

func (i MHSMVirtualNetworkRuleArgs) ToMHSMVirtualNetworkRuleOutput() MHSMVirtualNetworkRuleOutput {
	return i.ToMHSMVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i MHSMVirtualNetworkRuleArgs) ToMHSMVirtualNetworkRuleOutputWithContext(ctx context.Context) MHSMVirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMVirtualNetworkRuleOutput)
}





type MHSMVirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToMHSMVirtualNetworkRuleArrayOutput() MHSMVirtualNetworkRuleArrayOutput
	ToMHSMVirtualNetworkRuleArrayOutputWithContext(context.Context) MHSMVirtualNetworkRuleArrayOutput
}

type MHSMVirtualNetworkRuleArray []MHSMVirtualNetworkRuleInput

func (MHSMVirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMVirtualNetworkRule)(nil)).Elem()
}

func (i MHSMVirtualNetworkRuleArray) ToMHSMVirtualNetworkRuleArrayOutput() MHSMVirtualNetworkRuleArrayOutput {
	return i.ToMHSMVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i MHSMVirtualNetworkRuleArray) ToMHSMVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) MHSMVirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMVirtualNetworkRuleArrayOutput)
}

type MHSMVirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (MHSMVirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMVirtualNetworkRule)(nil)).Elem()
}

func (o MHSMVirtualNetworkRuleOutput) ToMHSMVirtualNetworkRuleOutput() MHSMVirtualNetworkRuleOutput {
	return o
}

func (o MHSMVirtualNetworkRuleOutput) ToMHSMVirtualNetworkRuleOutputWithContext(ctx context.Context) MHSMVirtualNetworkRuleOutput {
	return o
}

func (o MHSMVirtualNetworkRuleOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MHSMVirtualNetworkRule) string { return v.Id }).(pulumi.StringOutput)
}

type MHSMVirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (MHSMVirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMVirtualNetworkRule)(nil)).Elem()
}

func (o MHSMVirtualNetworkRuleArrayOutput) ToMHSMVirtualNetworkRuleArrayOutput() MHSMVirtualNetworkRuleArrayOutput {
	return o
}

func (o MHSMVirtualNetworkRuleArrayOutput) ToMHSMVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) MHSMVirtualNetworkRuleArrayOutput {
	return o
}

func (o MHSMVirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) MHSMVirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MHSMVirtualNetworkRule {
		return vs[0].([]MHSMVirtualNetworkRule)[vs[1].(int)]
	}).(MHSMVirtualNetworkRuleOutput)
}

type MHSMVirtualNetworkRuleResponse struct {
	Id string `pulumi:"id"`
}

type MHSMVirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (MHSMVirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMVirtualNetworkRuleResponse)(nil)).Elem()
}

func (o MHSMVirtualNetworkRuleResponseOutput) ToMHSMVirtualNetworkRuleResponseOutput() MHSMVirtualNetworkRuleResponseOutput {
	return o
}

func (o MHSMVirtualNetworkRuleResponseOutput) ToMHSMVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) MHSMVirtualNetworkRuleResponseOutput {
	return o
}

func (o MHSMVirtualNetworkRuleResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MHSMVirtualNetworkRuleResponse) string { return v.Id }).(pulumi.StringOutput)
}

type MHSMVirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (MHSMVirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMVirtualNetworkRuleResponse)(nil)).Elem()
}

func (o MHSMVirtualNetworkRuleResponseArrayOutput) ToMHSMVirtualNetworkRuleResponseArrayOutput() MHSMVirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o MHSMVirtualNetworkRuleResponseArrayOutput) ToMHSMVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) MHSMVirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o MHSMVirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) MHSMVirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MHSMVirtualNetworkRuleResponse {
		return vs[0].([]MHSMVirtualNetworkRuleResponse)[vs[1].(int)]
	}).(MHSMVirtualNetworkRuleResponseOutput)
}

type ManagedHSMSecurityDomainPropertiesResponse struct {
	ActivationStatus        string `pulumi:"activationStatus"`
	ActivationStatusMessage string `pulumi:"activationStatusMessage"`
}

type ManagedHSMSecurityDomainPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagedHSMSecurityDomainPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHSMSecurityDomainPropertiesResponse)(nil)).Elem()
}

func (o ManagedHSMSecurityDomainPropertiesResponseOutput) ToManagedHSMSecurityDomainPropertiesResponseOutput() ManagedHSMSecurityDomainPropertiesResponseOutput {
	return o
}

func (o ManagedHSMSecurityDomainPropertiesResponseOutput) ToManagedHSMSecurityDomainPropertiesResponseOutputWithContext(ctx context.Context) ManagedHSMSecurityDomainPropertiesResponseOutput {
	return o
}

func (o ManagedHSMSecurityDomainPropertiesResponseOutput) ActivationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHSMSecurityDomainPropertiesResponse) string { return v.ActivationStatus }).(pulumi.StringOutput)
}

func (o ManagedHSMSecurityDomainPropertiesResponseOutput) ActivationStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHSMSecurityDomainPropertiesResponse) string { return v.ActivationStatusMessage }).(pulumi.StringOutput)
}

type ManagedHsmProperties struct {
	CreateMode                *CreateMode         `pulumi:"createMode"`
	EnablePurgeProtection     *bool               `pulumi:"enablePurgeProtection"`
	EnableSoftDelete          *bool               `pulumi:"enableSoftDelete"`
	InitialAdminObjectIds     []string            `pulumi:"initialAdminObjectIds"`
	NetworkAcls               *MHSMNetworkRuleSet `pulumi:"networkAcls"`
	PublicNetworkAccess       *string             `pulumi:"publicNetworkAccess"`
	SoftDeleteRetentionInDays *int                `pulumi:"softDeleteRetentionInDays"`
	TenantId                  *string             `pulumi:"tenantId"`
}


func (val *ManagedHsmProperties) Defaults() *ManagedHsmProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnablePurgeProtection) {
		enablePurgeProtection_ := true
		tmp.EnablePurgeProtection = &enablePurgeProtection_
	}
	if isZero(tmp.EnableSoftDelete) {
		enableSoftDelete_ := true
		tmp.EnableSoftDelete = &enableSoftDelete_
	}
	if isZero(tmp.SoftDeleteRetentionInDays) {
		softDeleteRetentionInDays_ := 90
		tmp.SoftDeleteRetentionInDays = &softDeleteRetentionInDays_
	}
	return &tmp
}





type ManagedHsmPropertiesInput interface {
	pulumi.Input

	ToManagedHsmPropertiesOutput() ManagedHsmPropertiesOutput
	ToManagedHsmPropertiesOutputWithContext(context.Context) ManagedHsmPropertiesOutput
}

type ManagedHsmPropertiesArgs struct {
	CreateMode                CreateModePtrInput         `pulumi:"createMode"`
	EnablePurgeProtection     pulumi.BoolPtrInput        `pulumi:"enablePurgeProtection"`
	EnableSoftDelete          pulumi.BoolPtrInput        `pulumi:"enableSoftDelete"`
	InitialAdminObjectIds     pulumi.StringArrayInput    `pulumi:"initialAdminObjectIds"`
	NetworkAcls               MHSMNetworkRuleSetPtrInput `pulumi:"networkAcls"`
	PublicNetworkAccess       pulumi.StringPtrInput      `pulumi:"publicNetworkAccess"`
	SoftDeleteRetentionInDays pulumi.IntPtrInput         `pulumi:"softDeleteRetentionInDays"`
	TenantId                  pulumi.StringPtrInput      `pulumi:"tenantId"`
}


func (val *ManagedHsmPropertiesArgs) Defaults() *ManagedHsmPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnablePurgeProtection) {
		tmp.EnablePurgeProtection = pulumi.BoolPtr(true)
	}
	if isZero(tmp.EnableSoftDelete) {
		tmp.EnableSoftDelete = pulumi.BoolPtr(true)
	}
	if isZero(tmp.SoftDeleteRetentionInDays) {
		tmp.SoftDeleteRetentionInDays = pulumi.IntPtr(90)
	}
	return &tmp
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

func (o ManagedHsmPropertiesOutput) NetworkAcls() MHSMNetworkRuleSetPtrOutput {
	return o.ApplyT(func(v ManagedHsmProperties) *MHSMNetworkRuleSet { return v.NetworkAcls }).(MHSMNetworkRuleSetPtrOutput)
}

func (o ManagedHsmPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedHsmProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
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

func (o ManagedHsmPropertiesPtrOutput) NetworkAcls() MHSMNetworkRuleSetPtrOutput {
	return o.ApplyT(func(v *ManagedHsmProperties) *MHSMNetworkRuleSet {
		if v == nil {
			return nil
		}
		return v.NetworkAcls
	}).(MHSMNetworkRuleSetPtrOutput)
}

func (o ManagedHsmPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
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
	EnablePurgeProtection      *bool                                       `pulumi:"enablePurgeProtection"`
	EnableSoftDelete           *bool                                       `pulumi:"enableSoftDelete"`
	HsmUri                     string                                      `pulumi:"hsmUri"`
	InitialAdminObjectIds      []string                                    `pulumi:"initialAdminObjectIds"`
	NetworkAcls                *MHSMNetworkRuleSetResponse                 `pulumi:"networkAcls"`
	PrivateEndpointConnections []MHSMPrivateEndpointConnectionItemResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                      `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                     `pulumi:"publicNetworkAccess"`
	ScheduledPurgeDate         string                                      `pulumi:"scheduledPurgeDate"`
	SecurityDomainProperties   ManagedHSMSecurityDomainPropertiesResponse  `pulumi:"securityDomainProperties"`
	SoftDeleteRetentionInDays  *int                                        `pulumi:"softDeleteRetentionInDays"`
	StatusMessage              string                                      `pulumi:"statusMessage"`
	TenantId                   *string                                     `pulumi:"tenantId"`
}


func (val *ManagedHsmPropertiesResponse) Defaults() *ManagedHsmPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnablePurgeProtection) {
		enablePurgeProtection_ := true
		tmp.EnablePurgeProtection = &enablePurgeProtection_
	}
	if isZero(tmp.EnableSoftDelete) {
		enableSoftDelete_ := true
		tmp.EnableSoftDelete = &enableSoftDelete_
	}
	if isZero(tmp.SoftDeleteRetentionInDays) {
		softDeleteRetentionInDays_ := 90
		tmp.SoftDeleteRetentionInDays = &softDeleteRetentionInDays_
	}
	return &tmp
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

func (o ManagedHsmPropertiesResponseOutput) NetworkAcls() MHSMNetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) *MHSMNetworkRuleSetResponse { return v.NetworkAcls }).(MHSMNetworkRuleSetResponsePtrOutput)
}

func (o ManagedHsmPropertiesResponseOutput) PrivateEndpointConnections() MHSMPrivateEndpointConnectionItemResponseArrayOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) []MHSMPrivateEndpointConnectionItemResponse {
		return v.PrivateEndpointConnections
	}).(MHSMPrivateEndpointConnectionItemResponseArrayOutput)
}

func (o ManagedHsmPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedHsmPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ManagedHsmPropertiesResponseOutput) ScheduledPurgeDate() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) string { return v.ScheduledPurgeDate }).(pulumi.StringOutput)
}

func (o ManagedHsmPropertiesResponseOutput) SecurityDomainProperties() ManagedHSMSecurityDomainPropertiesResponseOutput {
	return o.ApplyT(func(v ManagedHsmPropertiesResponse) ManagedHSMSecurityDomainPropertiesResponse {
		return v.SecurityDomainProperties
	}).(ManagedHSMSecurityDomainPropertiesResponseOutput)
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

type RotationPolicy struct {
	Attributes      *KeyRotationPolicyAttributes `pulumi:"attributes"`
	LifetimeActions []LifetimeAction             `pulumi:"lifetimeActions"`
}





type RotationPolicyInput interface {
	pulumi.Input

	ToRotationPolicyOutput() RotationPolicyOutput
	ToRotationPolicyOutputWithContext(context.Context) RotationPolicyOutput
}

type RotationPolicyArgs struct {
	Attributes      KeyRotationPolicyAttributesPtrInput `pulumi:"attributes"`
	LifetimeActions LifetimeActionArrayInput            `pulumi:"lifetimeActions"`
}

func (RotationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RotationPolicy)(nil)).Elem()
}

func (i RotationPolicyArgs) ToRotationPolicyOutput() RotationPolicyOutput {
	return i.ToRotationPolicyOutputWithContext(context.Background())
}

func (i RotationPolicyArgs) ToRotationPolicyOutputWithContext(ctx context.Context) RotationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RotationPolicyOutput)
}

func (i RotationPolicyArgs) ToRotationPolicyPtrOutput() RotationPolicyPtrOutput {
	return i.ToRotationPolicyPtrOutputWithContext(context.Background())
}

func (i RotationPolicyArgs) ToRotationPolicyPtrOutputWithContext(ctx context.Context) RotationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RotationPolicyOutput).ToRotationPolicyPtrOutputWithContext(ctx)
}









type RotationPolicyPtrInput interface {
	pulumi.Input

	ToRotationPolicyPtrOutput() RotationPolicyPtrOutput
	ToRotationPolicyPtrOutputWithContext(context.Context) RotationPolicyPtrOutput
}

type rotationPolicyPtrType RotationPolicyArgs

func RotationPolicyPtr(v *RotationPolicyArgs) RotationPolicyPtrInput {
	return (*rotationPolicyPtrType)(v)
}

func (*rotationPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RotationPolicy)(nil)).Elem()
}

func (i *rotationPolicyPtrType) ToRotationPolicyPtrOutput() RotationPolicyPtrOutput {
	return i.ToRotationPolicyPtrOutputWithContext(context.Background())
}

func (i *rotationPolicyPtrType) ToRotationPolicyPtrOutputWithContext(ctx context.Context) RotationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RotationPolicyPtrOutput)
}

type RotationPolicyOutput struct{ *pulumi.OutputState }

func (RotationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RotationPolicy)(nil)).Elem()
}

func (o RotationPolicyOutput) ToRotationPolicyOutput() RotationPolicyOutput {
	return o
}

func (o RotationPolicyOutput) ToRotationPolicyOutputWithContext(ctx context.Context) RotationPolicyOutput {
	return o
}

func (o RotationPolicyOutput) ToRotationPolicyPtrOutput() RotationPolicyPtrOutput {
	return o.ToRotationPolicyPtrOutputWithContext(context.Background())
}

func (o RotationPolicyOutput) ToRotationPolicyPtrOutputWithContext(ctx context.Context) RotationPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RotationPolicy) *RotationPolicy {
		return &v
	}).(RotationPolicyPtrOutput)
}

func (o RotationPolicyOutput) Attributes() KeyRotationPolicyAttributesPtrOutput {
	return o.ApplyT(func(v RotationPolicy) *KeyRotationPolicyAttributes { return v.Attributes }).(KeyRotationPolicyAttributesPtrOutput)
}

func (o RotationPolicyOutput) LifetimeActions() LifetimeActionArrayOutput {
	return o.ApplyT(func(v RotationPolicy) []LifetimeAction { return v.LifetimeActions }).(LifetimeActionArrayOutput)
}

type RotationPolicyPtrOutput struct{ *pulumi.OutputState }

func (RotationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RotationPolicy)(nil)).Elem()
}

func (o RotationPolicyPtrOutput) ToRotationPolicyPtrOutput() RotationPolicyPtrOutput {
	return o
}

func (o RotationPolicyPtrOutput) ToRotationPolicyPtrOutputWithContext(ctx context.Context) RotationPolicyPtrOutput {
	return o
}

func (o RotationPolicyPtrOutput) Elem() RotationPolicyOutput {
	return o.ApplyT(func(v *RotationPolicy) RotationPolicy {
		if v != nil {
			return *v
		}
		var ret RotationPolicy
		return ret
	}).(RotationPolicyOutput)
}

func (o RotationPolicyPtrOutput) Attributes() KeyRotationPolicyAttributesPtrOutput {
	return o.ApplyT(func(v *RotationPolicy) *KeyRotationPolicyAttributes {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(KeyRotationPolicyAttributesPtrOutput)
}

func (o RotationPolicyPtrOutput) LifetimeActions() LifetimeActionArrayOutput {
	return o.ApplyT(func(v *RotationPolicy) []LifetimeAction {
		if v == nil {
			return nil
		}
		return v.LifetimeActions
	}).(LifetimeActionArrayOutput)
}

type RotationPolicyResponse struct {
	Attributes      *KeyRotationPolicyAttributesResponse `pulumi:"attributes"`
	LifetimeActions []LifetimeActionResponse             `pulumi:"lifetimeActions"`
}

type RotationPolicyResponseOutput struct{ *pulumi.OutputState }

func (RotationPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RotationPolicyResponse)(nil)).Elem()
}

func (o RotationPolicyResponseOutput) ToRotationPolicyResponseOutput() RotationPolicyResponseOutput {
	return o
}

func (o RotationPolicyResponseOutput) ToRotationPolicyResponseOutputWithContext(ctx context.Context) RotationPolicyResponseOutput {
	return o
}

func (o RotationPolicyResponseOutput) Attributes() KeyRotationPolicyAttributesResponsePtrOutput {
	return o.ApplyT(func(v RotationPolicyResponse) *KeyRotationPolicyAttributesResponse { return v.Attributes }).(KeyRotationPolicyAttributesResponsePtrOutput)
}

func (o RotationPolicyResponseOutput) LifetimeActions() LifetimeActionResponseArrayOutput {
	return o.ApplyT(func(v RotationPolicyResponse) []LifetimeActionResponse { return v.LifetimeActions }).(LifetimeActionResponseArrayOutput)
}

type RotationPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RotationPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RotationPolicyResponse)(nil)).Elem()
}

func (o RotationPolicyResponsePtrOutput) ToRotationPolicyResponsePtrOutput() RotationPolicyResponsePtrOutput {
	return o
}

func (o RotationPolicyResponsePtrOutput) ToRotationPolicyResponsePtrOutputWithContext(ctx context.Context) RotationPolicyResponsePtrOutput {
	return o
}

func (o RotationPolicyResponsePtrOutput) Elem() RotationPolicyResponseOutput {
	return o.ApplyT(func(v *RotationPolicyResponse) RotationPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RotationPolicyResponse
		return ret
	}).(RotationPolicyResponseOutput)
}

func (o RotationPolicyResponsePtrOutput) Attributes() KeyRotationPolicyAttributesResponsePtrOutput {
	return o.ApplyT(func(v *RotationPolicyResponse) *KeyRotationPolicyAttributesResponse {
		if v == nil {
			return nil
		}
		return v.Attributes
	}).(KeyRotationPolicyAttributesResponsePtrOutput)
}

func (o RotationPolicyResponsePtrOutput) LifetimeActions() LifetimeActionResponseArrayOutput {
	return o.ApplyT(func(v *RotationPolicyResponse) []LifetimeActionResponse {
		if v == nil {
			return nil
		}
		return v.LifetimeActions
	}).(LifetimeActionResponseArrayOutput)
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

type Trigger struct {
	TimeAfterCreate  *string `pulumi:"timeAfterCreate"`
	TimeBeforeExpiry *string `pulumi:"timeBeforeExpiry"`
}





type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(context.Context) TriggerOutput
}

type TriggerArgs struct {
	TimeAfterCreate  pulumi.StringPtrInput `pulumi:"timeAfterCreate"`
	TimeBeforeExpiry pulumi.StringPtrInput `pulumi:"timeBeforeExpiry"`
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil)).Elem()
}

func (i TriggerArgs) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i TriggerArgs) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

func (i TriggerArgs) ToTriggerPtrOutput() TriggerPtrOutput {
	return i.ToTriggerPtrOutputWithContext(context.Background())
}

func (i TriggerArgs) ToTriggerPtrOutputWithContext(ctx context.Context) TriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput).ToTriggerPtrOutputWithContext(ctx)
}









type TriggerPtrInput interface {
	pulumi.Input

	ToTriggerPtrOutput() TriggerPtrOutput
	ToTriggerPtrOutputWithContext(context.Context) TriggerPtrOutput
}

type triggerPtrType TriggerArgs

func TriggerPtr(v *TriggerArgs) TriggerPtrInput {
	return (*triggerPtrType)(v)
}

func (*triggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *triggerPtrType) ToTriggerPtrOutput() TriggerPtrOutput {
	return i.ToTriggerPtrOutputWithContext(context.Background())
}

func (i *triggerPtrType) ToTriggerPtrOutputWithContext(ctx context.Context) TriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPtrOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerPtrOutput() TriggerPtrOutput {
	return o.ToTriggerPtrOutputWithContext(context.Background())
}

func (o TriggerOutput) ToTriggerPtrOutputWithContext(ctx context.Context) TriggerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Trigger) *Trigger {
		return &v
	}).(TriggerPtrOutput)
}

func (o TriggerOutput) TimeAfterCreate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Trigger) *string { return v.TimeAfterCreate }).(pulumi.StringPtrOutput)
}

func (o TriggerOutput) TimeBeforeExpiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Trigger) *string { return v.TimeBeforeExpiry }).(pulumi.StringPtrOutput)
}

type TriggerPtrOutput struct{ *pulumi.OutputState }

func (TriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (o TriggerPtrOutput) ToTriggerPtrOutput() TriggerPtrOutput {
	return o
}

func (o TriggerPtrOutput) ToTriggerPtrOutputWithContext(ctx context.Context) TriggerPtrOutput {
	return o
}

func (o TriggerPtrOutput) Elem() TriggerOutput {
	return o.ApplyT(func(v *Trigger) Trigger {
		if v != nil {
			return *v
		}
		var ret Trigger
		return ret
	}).(TriggerOutput)
}

func (o TriggerPtrOutput) TimeAfterCreate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeAfterCreate
	}).(pulumi.StringPtrOutput)
}

func (o TriggerPtrOutput) TimeBeforeExpiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeBeforeExpiry
	}).(pulumi.StringPtrOutput)
}

type TriggerResponse struct {
	TimeAfterCreate  *string `pulumi:"timeAfterCreate"`
	TimeBeforeExpiry *string `pulumi:"timeBeforeExpiry"`
}

type TriggerResponseOutput struct{ *pulumi.OutputState }

func (TriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerResponse)(nil)).Elem()
}

func (o TriggerResponseOutput) ToTriggerResponseOutput() TriggerResponseOutput {
	return o
}

func (o TriggerResponseOutput) ToTriggerResponseOutputWithContext(ctx context.Context) TriggerResponseOutput {
	return o
}

func (o TriggerResponseOutput) TimeAfterCreate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerResponse) *string { return v.TimeAfterCreate }).(pulumi.StringPtrOutput)
}

func (o TriggerResponseOutput) TimeBeforeExpiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerResponse) *string { return v.TimeBeforeExpiry }).(pulumi.StringPtrOutput)
}

type TriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (TriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerResponse)(nil)).Elem()
}

func (o TriggerResponsePtrOutput) ToTriggerResponsePtrOutput() TriggerResponsePtrOutput {
	return o
}

func (o TriggerResponsePtrOutput) ToTriggerResponsePtrOutputWithContext(ctx context.Context) TriggerResponsePtrOutput {
	return o
}

func (o TriggerResponsePtrOutput) Elem() TriggerResponseOutput {
	return o.ApplyT(func(v *TriggerResponse) TriggerResponse {
		if v != nil {
			return *v
		}
		var ret TriggerResponse
		return ret
	}).(TriggerResponseOutput)
}

func (o TriggerResponsePtrOutput) TimeAfterCreate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeAfterCreate
	}).(pulumi.StringPtrOutput)
}

func (o TriggerResponsePtrOutput) TimeBeforeExpiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeBeforeExpiry
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
	PublicNetworkAccess          *string             `pulumi:"publicNetworkAccess"`
	Sku                          Sku                 `pulumi:"sku"`
	SoftDeleteRetentionInDays    *int                `pulumi:"softDeleteRetentionInDays"`
	TenantId                     string              `pulumi:"tenantId"`
	VaultUri                     *string             `pulumi:"vaultUri"`
}


func (val *VaultProperties) Defaults() *VaultProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableRbacAuthorization) {
		enableRbacAuthorization_ := false
		tmp.EnableRbacAuthorization = &enableRbacAuthorization_
	}
	if isZero(tmp.EnableSoftDelete) {
		enableSoftDelete_ := true
		tmp.EnableSoftDelete = &enableSoftDelete_
	}
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	if isZero(tmp.SoftDeleteRetentionInDays) {
		softDeleteRetentionInDays_ := 90
		tmp.SoftDeleteRetentionInDays = &softDeleteRetentionInDays_
	}
	return &tmp
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
	PublicNetworkAccess          pulumi.StringPtrInput       `pulumi:"publicNetworkAccess"`
	Sku                          SkuInput                    `pulumi:"sku"`
	SoftDeleteRetentionInDays    pulumi.IntPtrInput          `pulumi:"softDeleteRetentionInDays"`
	TenantId                     pulumi.StringInput          `pulumi:"tenantId"`
	VaultUri                     pulumi.StringPtrInput       `pulumi:"vaultUri"`
}


func (val *VaultPropertiesArgs) Defaults() *VaultPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableRbacAuthorization) {
		tmp.EnableRbacAuthorization = pulumi.BoolPtr(false)
	}
	if isZero(tmp.EnableSoftDelete) {
		tmp.EnableSoftDelete = pulumi.BoolPtr(true)
	}
	if isZero(tmp.PublicNetworkAccess) {
		tmp.PublicNetworkAccess = pulumi.StringPtr("enabled")
	}
	if isZero(tmp.SoftDeleteRetentionInDays) {
		tmp.SoftDeleteRetentionInDays = pulumi.IntPtr(90)
	}
	return &tmp
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

func (o VaultPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
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

type VaultPropertiesResponse struct {
	AccessPolicies               []AccessPolicyEntryResponse             `pulumi:"accessPolicies"`
	EnablePurgeProtection        *bool                                   `pulumi:"enablePurgeProtection"`
	EnableRbacAuthorization      *bool                                   `pulumi:"enableRbacAuthorization"`
	EnableSoftDelete             *bool                                   `pulumi:"enableSoftDelete"`
	EnabledForDeployment         *bool                                   `pulumi:"enabledForDeployment"`
	EnabledForDiskEncryption     *bool                                   `pulumi:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment *bool                                   `pulumi:"enabledForTemplateDeployment"`
	HsmPoolResourceId            string                                  `pulumi:"hsmPoolResourceId"`
	NetworkAcls                  *NetworkRuleSetResponse                 `pulumi:"networkAcls"`
	PrivateEndpointConnections   []PrivateEndpointConnectionItemResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState            *string                                 `pulumi:"provisioningState"`
	PublicNetworkAccess          *string                                 `pulumi:"publicNetworkAccess"`
	Sku                          SkuResponse                             `pulumi:"sku"`
	SoftDeleteRetentionInDays    *int                                    `pulumi:"softDeleteRetentionInDays"`
	TenantId                     string                                  `pulumi:"tenantId"`
	VaultUri                     *string                                 `pulumi:"vaultUri"`
}


func (val *VaultPropertiesResponse) Defaults() *VaultPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableRbacAuthorization) {
		enableRbacAuthorization_ := false
		tmp.EnableRbacAuthorization = &enableRbacAuthorization_
	}
	if isZero(tmp.EnableSoftDelete) {
		enableSoftDelete_ := true
		tmp.EnableSoftDelete = &enableSoftDelete_
	}
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	if isZero(tmp.SoftDeleteRetentionInDays) {
		softDeleteRetentionInDays_ := 90
		tmp.SoftDeleteRetentionInDays = &softDeleteRetentionInDays_
	}
	return &tmp
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

func (o VaultPropertiesResponseOutput) HsmPoolResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.HsmPoolResourceId }).(pulumi.StringOutput)
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

func (o VaultPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(AccessPolicyEntryOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryArrayOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryResponseOutput{})
	pulumi.RegisterOutputType(AccessPolicyEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(ActionResponseOutput{})
	pulumi.RegisterOutputType(ActionResponsePtrOutput{})
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyAttributesOutput{})
	pulumi.RegisterOutputType(KeyAttributesPtrOutput{})
	pulumi.RegisterOutputType(KeyAttributesResponseOutput{})
	pulumi.RegisterOutputType(KeyAttributesResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyPropertiesOutput{})
	pulumi.RegisterOutputType(KeyReleasePolicyOutput{})
	pulumi.RegisterOutputType(KeyReleasePolicyPtrOutput{})
	pulumi.RegisterOutputType(KeyReleasePolicyResponseOutput{})
	pulumi.RegisterOutputType(KeyReleasePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyRotationPolicyAttributesOutput{})
	pulumi.RegisterOutputType(KeyRotationPolicyAttributesPtrOutput{})
	pulumi.RegisterOutputType(KeyRotationPolicyAttributesResponseOutput{})
	pulumi.RegisterOutputType(KeyRotationPolicyAttributesResponsePtrOutput{})
	pulumi.RegisterOutputType(LifetimeActionOutput{})
	pulumi.RegisterOutputType(LifetimeActionArrayOutput{})
	pulumi.RegisterOutputType(LifetimeActionResponseOutput{})
	pulumi.RegisterOutputType(LifetimeActionResponseArrayOutput{})
	pulumi.RegisterOutputType(MHSMIPRuleOutput{})
	pulumi.RegisterOutputType(MHSMIPRuleArrayOutput{})
	pulumi.RegisterOutputType(MHSMIPRuleResponseOutput{})
	pulumi.RegisterOutputType(MHSMIPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(MHSMNetworkRuleSetOutput{})
	pulumi.RegisterOutputType(MHSMNetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(MHSMNetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(MHSMNetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(MHSMPrivateEndpointConnectionItemResponseOutput{})
	pulumi.RegisterOutputType(MHSMPrivateEndpointConnectionItemResponseArrayOutput{})
	pulumi.RegisterOutputType(MHSMPrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(MHSMPrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(MHSMPrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(MHSMPrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(MHSMPrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(MHSMPrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(MHSMVirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(MHSMVirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(MHSMVirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(MHSMVirtualNetworkRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedHSMSecurityDomainPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagedHsmPropertiesOutput{})
	pulumi.RegisterOutputType(ManagedHsmPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedHsmPropertiesResponseOutput{})
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
	pulumi.RegisterOutputType(RotationPolicyOutput{})
	pulumi.RegisterOutputType(RotationPolicyPtrOutput{})
	pulumi.RegisterOutputType(RotationPolicyResponseOutput{})
	pulumi.RegisterOutputType(RotationPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(SecretAttributesOutput{})
	pulumi.RegisterOutputType(SecretAttributesPtrOutput{})
	pulumi.RegisterOutputType(SecretAttributesResponseOutput{})
	pulumi.RegisterOutputType(SecretAttributesResponsePtrOutput{})
	pulumi.RegisterOutputType(SecretPropertiesOutput{})
	pulumi.RegisterOutputType(SecretPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TriggerOutput{})
	pulumi.RegisterOutputType(TriggerPtrOutput{})
	pulumi.RegisterOutputType(TriggerResponseOutput{})
	pulumi.RegisterOutputType(TriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
