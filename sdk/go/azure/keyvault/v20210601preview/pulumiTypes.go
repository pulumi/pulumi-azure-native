


package v20210601preview

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





type ActionResponseInput interface {
	pulumi.Input

	ToActionResponseOutput() ActionResponseOutput
	ToActionResponseOutputWithContext(context.Context) ActionResponseOutput
}

type ActionResponseArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionResponse)(nil)).Elem()
}

func (i ActionResponseArgs) ToActionResponseOutput() ActionResponseOutput {
	return i.ToActionResponseOutputWithContext(context.Background())
}

func (i ActionResponseArgs) ToActionResponseOutputWithContext(ctx context.Context) ActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionResponseOutput)
}

func (i ActionResponseArgs) ToActionResponsePtrOutput() ActionResponsePtrOutput {
	return i.ToActionResponsePtrOutputWithContext(context.Background())
}

func (i ActionResponseArgs) ToActionResponsePtrOutputWithContext(ctx context.Context) ActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionResponseOutput).ToActionResponsePtrOutputWithContext(ctx)
}









type ActionResponsePtrInput interface {
	pulumi.Input

	ToActionResponsePtrOutput() ActionResponsePtrOutput
	ToActionResponsePtrOutputWithContext(context.Context) ActionResponsePtrOutput
}

type actionResponsePtrType ActionResponseArgs

func ActionResponsePtr(v *ActionResponseArgs) ActionResponsePtrInput {
	return (*actionResponsePtrType)(v)
}

func (*actionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionResponse)(nil)).Elem()
}

func (i *actionResponsePtrType) ToActionResponsePtrOutput() ActionResponsePtrOutput {
	return i.ToActionResponsePtrOutputWithContext(context.Background())
}

func (i *actionResponsePtrType) ToActionResponsePtrOutputWithContext(ctx context.Context) ActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionResponsePtrOutput)
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

func (o ActionResponseOutput) ToActionResponsePtrOutput() ActionResponsePtrOutput {
	return o.ToActionResponsePtrOutputWithContext(context.Background())
}

func (o ActionResponseOutput) ToActionResponsePtrOutputWithContext(ctx context.Context) ActionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionResponse) *ActionResponse {
		return &v
	}).(ActionResponsePtrOutput)
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
	Enabled    *bool    `pulumi:"enabled"`
	Expires    *float64 `pulumi:"expires"`
	Exportable *bool    `pulumi:"exportable"`
	NotBefore  *float64 `pulumi:"notBefore"`
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





type KeyAttributesResponseInput interface {
	pulumi.Input

	ToKeyAttributesResponseOutput() KeyAttributesResponseOutput
	ToKeyAttributesResponseOutputWithContext(context.Context) KeyAttributesResponseOutput
}

type KeyAttributesResponseArgs struct {
	Created       pulumi.Float64Input    `pulumi:"created"`
	Enabled       pulumi.BoolPtrInput    `pulumi:"enabled"`
	Expires       pulumi.Float64PtrInput `pulumi:"expires"`
	Exportable    pulumi.BoolPtrInput    `pulumi:"exportable"`
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

func (o KeyPropertiesOutput) ReleasePolicy() KeyReleasePolicyPtrOutput {
	return o.ApplyT(func(v KeyProperties) *KeyReleasePolicy { return v.ReleasePolicy }).(KeyReleasePolicyPtrOutput)
}

func (o KeyPropertiesOutput) RotationPolicy() RotationPolicyPtrOutput {
	return o.ApplyT(func(v KeyProperties) *RotationPolicy { return v.RotationPolicy }).(RotationPolicyPtrOutput)
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

func (o KeyPropertiesPtrOutput) ReleasePolicy() KeyReleasePolicyPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *KeyReleasePolicy {
		if v == nil {
			return nil
		}
		return v.ReleasePolicy
	}).(KeyReleasePolicyPtrOutput)
}

func (o KeyPropertiesPtrOutput) RotationPolicy() RotationPolicyPtrOutput {
	return o.ApplyT(func(v *KeyProperties) *RotationPolicy {
		if v == nil {
			return nil
		}
		return v.RotationPolicy
	}).(RotationPolicyPtrOutput)
}

type KeyReleasePolicy struct {
	ContentType *string `pulumi:"contentType"`
	Data        *string `pulumi:"data"`
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





type KeyReleasePolicyResponseInput interface {
	pulumi.Input

	ToKeyReleasePolicyResponseOutput() KeyReleasePolicyResponseOutput
	ToKeyReleasePolicyResponseOutputWithContext(context.Context) KeyReleasePolicyResponseOutput
}

type KeyReleasePolicyResponseArgs struct {
	ContentType pulumi.StringPtrInput `pulumi:"contentType"`
	Data        pulumi.StringPtrInput `pulumi:"data"`
}

func (KeyReleasePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyReleasePolicyResponse)(nil)).Elem()
}

func (i KeyReleasePolicyResponseArgs) ToKeyReleasePolicyResponseOutput() KeyReleasePolicyResponseOutput {
	return i.ToKeyReleasePolicyResponseOutputWithContext(context.Background())
}

func (i KeyReleasePolicyResponseArgs) ToKeyReleasePolicyResponseOutputWithContext(ctx context.Context) KeyReleasePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyReleasePolicyResponseOutput)
}

func (i KeyReleasePolicyResponseArgs) ToKeyReleasePolicyResponsePtrOutput() KeyReleasePolicyResponsePtrOutput {
	return i.ToKeyReleasePolicyResponsePtrOutputWithContext(context.Background())
}

func (i KeyReleasePolicyResponseArgs) ToKeyReleasePolicyResponsePtrOutputWithContext(ctx context.Context) KeyReleasePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyReleasePolicyResponseOutput).ToKeyReleasePolicyResponsePtrOutputWithContext(ctx)
}









type KeyReleasePolicyResponsePtrInput interface {
	pulumi.Input

	ToKeyReleasePolicyResponsePtrOutput() KeyReleasePolicyResponsePtrOutput
	ToKeyReleasePolicyResponsePtrOutputWithContext(context.Context) KeyReleasePolicyResponsePtrOutput
}

type keyReleasePolicyResponsePtrType KeyReleasePolicyResponseArgs

func KeyReleasePolicyResponsePtr(v *KeyReleasePolicyResponseArgs) KeyReleasePolicyResponsePtrInput {
	return (*keyReleasePolicyResponsePtrType)(v)
}

func (*keyReleasePolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyReleasePolicyResponse)(nil)).Elem()
}

func (i *keyReleasePolicyResponsePtrType) ToKeyReleasePolicyResponsePtrOutput() KeyReleasePolicyResponsePtrOutput {
	return i.ToKeyReleasePolicyResponsePtrOutputWithContext(context.Background())
}

func (i *keyReleasePolicyResponsePtrType) ToKeyReleasePolicyResponsePtrOutputWithContext(ctx context.Context) KeyReleasePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyReleasePolicyResponsePtrOutput)
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

func (o KeyReleasePolicyResponseOutput) ToKeyReleasePolicyResponsePtrOutput() KeyReleasePolicyResponsePtrOutput {
	return o.ToKeyReleasePolicyResponsePtrOutputWithContext(context.Background())
}

func (o KeyReleasePolicyResponseOutput) ToKeyReleasePolicyResponsePtrOutputWithContext(ctx context.Context) KeyReleasePolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyReleasePolicyResponse) *KeyReleasePolicyResponse {
		return &v
	}).(KeyReleasePolicyResponsePtrOutput)
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





type KeyRotationPolicyAttributesResponseInput interface {
	pulumi.Input

	ToKeyRotationPolicyAttributesResponseOutput() KeyRotationPolicyAttributesResponseOutput
	ToKeyRotationPolicyAttributesResponseOutputWithContext(context.Context) KeyRotationPolicyAttributesResponseOutput
}

type KeyRotationPolicyAttributesResponseArgs struct {
	Created    pulumi.Float64Input   `pulumi:"created"`
	ExpiryTime pulumi.StringPtrInput `pulumi:"expiryTime"`
	Updated    pulumi.Float64Input   `pulumi:"updated"`
}

func (KeyRotationPolicyAttributesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyRotationPolicyAttributesResponse)(nil)).Elem()
}

func (i KeyRotationPolicyAttributesResponseArgs) ToKeyRotationPolicyAttributesResponseOutput() KeyRotationPolicyAttributesResponseOutput {
	return i.ToKeyRotationPolicyAttributesResponseOutputWithContext(context.Background())
}

func (i KeyRotationPolicyAttributesResponseArgs) ToKeyRotationPolicyAttributesResponseOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRotationPolicyAttributesResponseOutput)
}

func (i KeyRotationPolicyAttributesResponseArgs) ToKeyRotationPolicyAttributesResponsePtrOutput() KeyRotationPolicyAttributesResponsePtrOutput {
	return i.ToKeyRotationPolicyAttributesResponsePtrOutputWithContext(context.Background())
}

func (i KeyRotationPolicyAttributesResponseArgs) ToKeyRotationPolicyAttributesResponsePtrOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRotationPolicyAttributesResponseOutput).ToKeyRotationPolicyAttributesResponsePtrOutputWithContext(ctx)
}









type KeyRotationPolicyAttributesResponsePtrInput interface {
	pulumi.Input

	ToKeyRotationPolicyAttributesResponsePtrOutput() KeyRotationPolicyAttributesResponsePtrOutput
	ToKeyRotationPolicyAttributesResponsePtrOutputWithContext(context.Context) KeyRotationPolicyAttributesResponsePtrOutput
}

type keyRotationPolicyAttributesResponsePtrType KeyRotationPolicyAttributesResponseArgs

func KeyRotationPolicyAttributesResponsePtr(v *KeyRotationPolicyAttributesResponseArgs) KeyRotationPolicyAttributesResponsePtrInput {
	return (*keyRotationPolicyAttributesResponsePtrType)(v)
}

func (*keyRotationPolicyAttributesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRotationPolicyAttributesResponse)(nil)).Elem()
}

func (i *keyRotationPolicyAttributesResponsePtrType) ToKeyRotationPolicyAttributesResponsePtrOutput() KeyRotationPolicyAttributesResponsePtrOutput {
	return i.ToKeyRotationPolicyAttributesResponsePtrOutputWithContext(context.Background())
}

func (i *keyRotationPolicyAttributesResponsePtrType) ToKeyRotationPolicyAttributesResponsePtrOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRotationPolicyAttributesResponsePtrOutput)
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

func (o KeyRotationPolicyAttributesResponseOutput) ToKeyRotationPolicyAttributesResponsePtrOutput() KeyRotationPolicyAttributesResponsePtrOutput {
	return o.ToKeyRotationPolicyAttributesResponsePtrOutputWithContext(context.Background())
}

func (o KeyRotationPolicyAttributesResponseOutput) ToKeyRotationPolicyAttributesResponsePtrOutputWithContext(ctx context.Context) KeyRotationPolicyAttributesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyRotationPolicyAttributesResponse) *KeyRotationPolicyAttributesResponse {
		return &v
	}).(KeyRotationPolicyAttributesResponsePtrOutput)
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





type LifetimeActionResponseInput interface {
	pulumi.Input

	ToLifetimeActionResponseOutput() LifetimeActionResponseOutput
	ToLifetimeActionResponseOutputWithContext(context.Context) LifetimeActionResponseOutput
}

type LifetimeActionResponseArgs struct {
	Action  ActionResponsePtrInput  `pulumi:"action"`
	Trigger TriggerResponsePtrInput `pulumi:"trigger"`
}

func (LifetimeActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LifetimeActionResponse)(nil)).Elem()
}

func (i LifetimeActionResponseArgs) ToLifetimeActionResponseOutput() LifetimeActionResponseOutput {
	return i.ToLifetimeActionResponseOutputWithContext(context.Background())
}

func (i LifetimeActionResponseArgs) ToLifetimeActionResponseOutputWithContext(ctx context.Context) LifetimeActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifetimeActionResponseOutput)
}





type LifetimeActionResponseArrayInput interface {
	pulumi.Input

	ToLifetimeActionResponseArrayOutput() LifetimeActionResponseArrayOutput
	ToLifetimeActionResponseArrayOutputWithContext(context.Context) LifetimeActionResponseArrayOutput
}

type LifetimeActionResponseArray []LifetimeActionResponseInput

func (LifetimeActionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LifetimeActionResponse)(nil)).Elem()
}

func (i LifetimeActionResponseArray) ToLifetimeActionResponseArrayOutput() LifetimeActionResponseArrayOutput {
	return i.ToLifetimeActionResponseArrayOutputWithContext(context.Background())
}

func (i LifetimeActionResponseArray) ToLifetimeActionResponseArrayOutputWithContext(ctx context.Context) LifetimeActionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifetimeActionResponseArrayOutput)
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





type MHSMIPRuleResponseInput interface {
	pulumi.Input

	ToMHSMIPRuleResponseOutput() MHSMIPRuleResponseOutput
	ToMHSMIPRuleResponseOutputWithContext(context.Context) MHSMIPRuleResponseOutput
}

type MHSMIPRuleResponseArgs struct {
	Value pulumi.StringInput `pulumi:"value"`
}

func (MHSMIPRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMIPRuleResponse)(nil)).Elem()
}

func (i MHSMIPRuleResponseArgs) ToMHSMIPRuleResponseOutput() MHSMIPRuleResponseOutput {
	return i.ToMHSMIPRuleResponseOutputWithContext(context.Background())
}

func (i MHSMIPRuleResponseArgs) ToMHSMIPRuleResponseOutputWithContext(ctx context.Context) MHSMIPRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMIPRuleResponseOutput)
}





type MHSMIPRuleResponseArrayInput interface {
	pulumi.Input

	ToMHSMIPRuleResponseArrayOutput() MHSMIPRuleResponseArrayOutput
	ToMHSMIPRuleResponseArrayOutputWithContext(context.Context) MHSMIPRuleResponseArrayOutput
}

type MHSMIPRuleResponseArray []MHSMIPRuleResponseInput

func (MHSMIPRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMIPRuleResponse)(nil)).Elem()
}

func (i MHSMIPRuleResponseArray) ToMHSMIPRuleResponseArrayOutput() MHSMIPRuleResponseArrayOutput {
	return i.ToMHSMIPRuleResponseArrayOutputWithContext(context.Background())
}

func (i MHSMIPRuleResponseArray) ToMHSMIPRuleResponseArrayOutputWithContext(ctx context.Context) MHSMIPRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMIPRuleResponseArrayOutput)
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





type MHSMNetworkRuleSetResponseInput interface {
	pulumi.Input

	ToMHSMNetworkRuleSetResponseOutput() MHSMNetworkRuleSetResponseOutput
	ToMHSMNetworkRuleSetResponseOutputWithContext(context.Context) MHSMNetworkRuleSetResponseOutput
}

type MHSMNetworkRuleSetResponseArgs struct {
	Bypass              pulumi.StringPtrInput                    `pulumi:"bypass"`
	DefaultAction       pulumi.StringPtrInput                    `pulumi:"defaultAction"`
	IpRules             MHSMIPRuleResponseArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules MHSMVirtualNetworkRuleResponseArrayInput `pulumi:"virtualNetworkRules"`
}

func (MHSMNetworkRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMNetworkRuleSetResponse)(nil)).Elem()
}

func (i MHSMNetworkRuleSetResponseArgs) ToMHSMNetworkRuleSetResponseOutput() MHSMNetworkRuleSetResponseOutput {
	return i.ToMHSMNetworkRuleSetResponseOutputWithContext(context.Background())
}

func (i MHSMNetworkRuleSetResponseArgs) ToMHSMNetworkRuleSetResponseOutputWithContext(ctx context.Context) MHSMNetworkRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMNetworkRuleSetResponseOutput)
}

func (i MHSMNetworkRuleSetResponseArgs) ToMHSMNetworkRuleSetResponsePtrOutput() MHSMNetworkRuleSetResponsePtrOutput {
	return i.ToMHSMNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i MHSMNetworkRuleSetResponseArgs) ToMHSMNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) MHSMNetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMNetworkRuleSetResponseOutput).ToMHSMNetworkRuleSetResponsePtrOutputWithContext(ctx)
}









type MHSMNetworkRuleSetResponsePtrInput interface {
	pulumi.Input

	ToMHSMNetworkRuleSetResponsePtrOutput() MHSMNetworkRuleSetResponsePtrOutput
	ToMHSMNetworkRuleSetResponsePtrOutputWithContext(context.Context) MHSMNetworkRuleSetResponsePtrOutput
}

type mhsmnetworkRuleSetResponsePtrType MHSMNetworkRuleSetResponseArgs

func MHSMNetworkRuleSetResponsePtr(v *MHSMNetworkRuleSetResponseArgs) MHSMNetworkRuleSetResponsePtrInput {
	return (*mhsmnetworkRuleSetResponsePtrType)(v)
}

func (*mhsmnetworkRuleSetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMNetworkRuleSetResponse)(nil)).Elem()
}

func (i *mhsmnetworkRuleSetResponsePtrType) ToMHSMNetworkRuleSetResponsePtrOutput() MHSMNetworkRuleSetResponsePtrOutput {
	return i.ToMHSMNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i *mhsmnetworkRuleSetResponsePtrType) ToMHSMNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) MHSMNetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMNetworkRuleSetResponsePtrOutput)
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

func (o MHSMNetworkRuleSetResponseOutput) ToMHSMNetworkRuleSetResponsePtrOutput() MHSMNetworkRuleSetResponsePtrOutput {
	return o.ToMHSMNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (o MHSMNetworkRuleSetResponseOutput) ToMHSMNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) MHSMNetworkRuleSetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MHSMNetworkRuleSetResponse) *MHSMNetworkRuleSetResponse {
		return &v
	}).(MHSMNetworkRuleSetResponsePtrOutput)
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
	PrivateEndpoint                   *MHSMPrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *MHSMPrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                         `pulumi:"provisioningState"`
}





type MHSMPrivateEndpointConnectionItemResponseInput interface {
	pulumi.Input

	ToMHSMPrivateEndpointConnectionItemResponseOutput() MHSMPrivateEndpointConnectionItemResponseOutput
	ToMHSMPrivateEndpointConnectionItemResponseOutputWithContext(context.Context) MHSMPrivateEndpointConnectionItemResponseOutput
}

type MHSMPrivateEndpointConnectionItemResponseArgs struct {
	PrivateEndpoint                   MHSMPrivateEndpointResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState MHSMPrivateLinkServiceConnectionStateResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                    `pulumi:"provisioningState"`
}

func (MHSMPrivateEndpointConnectionItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateEndpointConnectionItemResponse)(nil)).Elem()
}

func (i MHSMPrivateEndpointConnectionItemResponseArgs) ToMHSMPrivateEndpointConnectionItemResponseOutput() MHSMPrivateEndpointConnectionItemResponseOutput {
	return i.ToMHSMPrivateEndpointConnectionItemResponseOutputWithContext(context.Background())
}

func (i MHSMPrivateEndpointConnectionItemResponseArgs) ToMHSMPrivateEndpointConnectionItemResponseOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateEndpointConnectionItemResponseOutput)
}





type MHSMPrivateEndpointConnectionItemResponseArrayInput interface {
	pulumi.Input

	ToMHSMPrivateEndpointConnectionItemResponseArrayOutput() MHSMPrivateEndpointConnectionItemResponseArrayOutput
	ToMHSMPrivateEndpointConnectionItemResponseArrayOutputWithContext(context.Context) MHSMPrivateEndpointConnectionItemResponseArrayOutput
}

type MHSMPrivateEndpointConnectionItemResponseArray []MHSMPrivateEndpointConnectionItemResponseInput

func (MHSMPrivateEndpointConnectionItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMPrivateEndpointConnectionItemResponse)(nil)).Elem()
}

func (i MHSMPrivateEndpointConnectionItemResponseArray) ToMHSMPrivateEndpointConnectionItemResponseArrayOutput() MHSMPrivateEndpointConnectionItemResponseArrayOutput {
	return i.ToMHSMPrivateEndpointConnectionItemResponseArrayOutputWithContext(context.Background())
}

func (i MHSMPrivateEndpointConnectionItemResponseArray) ToMHSMPrivateEndpointConnectionItemResponseArrayOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateEndpointConnectionItemResponseArrayOutput)
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





type MHSMPrivateEndpointResponseInput interface {
	pulumi.Input

	ToMHSMPrivateEndpointResponseOutput() MHSMPrivateEndpointResponseOutput
	ToMHSMPrivateEndpointResponseOutputWithContext(context.Context) MHSMPrivateEndpointResponseOutput
}

type MHSMPrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (MHSMPrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateEndpointResponse)(nil)).Elem()
}

func (i MHSMPrivateEndpointResponseArgs) ToMHSMPrivateEndpointResponseOutput() MHSMPrivateEndpointResponseOutput {
	return i.ToMHSMPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i MHSMPrivateEndpointResponseArgs) ToMHSMPrivateEndpointResponseOutputWithContext(ctx context.Context) MHSMPrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateEndpointResponseOutput)
}

func (i MHSMPrivateEndpointResponseArgs) ToMHSMPrivateEndpointResponsePtrOutput() MHSMPrivateEndpointResponsePtrOutput {
	return i.ToMHSMPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i MHSMPrivateEndpointResponseArgs) ToMHSMPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) MHSMPrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateEndpointResponseOutput).ToMHSMPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type MHSMPrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToMHSMPrivateEndpointResponsePtrOutput() MHSMPrivateEndpointResponsePtrOutput
	ToMHSMPrivateEndpointResponsePtrOutputWithContext(context.Context) MHSMPrivateEndpointResponsePtrOutput
}

type mhsmprivateEndpointResponsePtrType MHSMPrivateEndpointResponseArgs

func MHSMPrivateEndpointResponsePtr(v *MHSMPrivateEndpointResponseArgs) MHSMPrivateEndpointResponsePtrInput {
	return (*mhsmprivateEndpointResponsePtrType)(v)
}

func (*mhsmprivateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMPrivateEndpointResponse)(nil)).Elem()
}

func (i *mhsmprivateEndpointResponsePtrType) ToMHSMPrivateEndpointResponsePtrOutput() MHSMPrivateEndpointResponsePtrOutput {
	return i.ToMHSMPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *mhsmprivateEndpointResponsePtrType) ToMHSMPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) MHSMPrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateEndpointResponsePtrOutput)
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

func (o MHSMPrivateEndpointResponseOutput) ToMHSMPrivateEndpointResponsePtrOutput() MHSMPrivateEndpointResponsePtrOutput {
	return o.ToMHSMPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o MHSMPrivateEndpointResponseOutput) ToMHSMPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) MHSMPrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MHSMPrivateEndpointResponse) *MHSMPrivateEndpointResponse {
		return &v
	}).(MHSMPrivateEndpointResponsePtrOutput)
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





type MHSMPrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToMHSMPrivateLinkServiceConnectionStateResponseOutput() MHSMPrivateLinkServiceConnectionStateResponseOutput
	ToMHSMPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) MHSMPrivateLinkServiceConnectionStateResponseOutput
}

type MHSMPrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (MHSMPrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i MHSMPrivateLinkServiceConnectionStateResponseArgs) ToMHSMPrivateLinkServiceConnectionStateResponseOutput() MHSMPrivateLinkServiceConnectionStateResponseOutput {
	return i.ToMHSMPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i MHSMPrivateLinkServiceConnectionStateResponseArgs) ToMHSMPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateLinkServiceConnectionStateResponseOutput)
}

func (i MHSMPrivateLinkServiceConnectionStateResponseArgs) ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutput() MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i MHSMPrivateLinkServiceConnectionStateResponseArgs) ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateLinkServiceConnectionStateResponseOutput).ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type MHSMPrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutput() MHSMPrivateLinkServiceConnectionStateResponsePtrOutput
	ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) MHSMPrivateLinkServiceConnectionStateResponsePtrOutput
}

type mhsmprivateLinkServiceConnectionStateResponsePtrType MHSMPrivateLinkServiceConnectionStateResponseArgs

func MHSMPrivateLinkServiceConnectionStateResponsePtr(v *MHSMPrivateLinkServiceConnectionStateResponseArgs) MHSMPrivateLinkServiceConnectionStateResponsePtrInput {
	return (*mhsmprivateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*mhsmprivateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMPrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *mhsmprivateLinkServiceConnectionStateResponsePtrType) ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutput() MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *mhsmprivateLinkServiceConnectionStateResponsePtrType) ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateLinkServiceConnectionStateResponsePtrOutput)
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

func (o MHSMPrivateLinkServiceConnectionStateResponseOutput) ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutput() MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o MHSMPrivateLinkServiceConnectionStateResponseOutput) ToMHSMPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MHSMPrivateLinkServiceConnectionStateResponse) *MHSMPrivateLinkServiceConnectionStateResponse {
		return &v
	}).(MHSMPrivateLinkServiceConnectionStateResponsePtrOutput)
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





type MHSMVirtualNetworkRuleResponseInput interface {
	pulumi.Input

	ToMHSMVirtualNetworkRuleResponseOutput() MHSMVirtualNetworkRuleResponseOutput
	ToMHSMVirtualNetworkRuleResponseOutputWithContext(context.Context) MHSMVirtualNetworkRuleResponseOutput
}

type MHSMVirtualNetworkRuleResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (MHSMVirtualNetworkRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMVirtualNetworkRuleResponse)(nil)).Elem()
}

func (i MHSMVirtualNetworkRuleResponseArgs) ToMHSMVirtualNetworkRuleResponseOutput() MHSMVirtualNetworkRuleResponseOutput {
	return i.ToMHSMVirtualNetworkRuleResponseOutputWithContext(context.Background())
}

func (i MHSMVirtualNetworkRuleResponseArgs) ToMHSMVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) MHSMVirtualNetworkRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMVirtualNetworkRuleResponseOutput)
}





type MHSMVirtualNetworkRuleResponseArrayInput interface {
	pulumi.Input

	ToMHSMVirtualNetworkRuleResponseArrayOutput() MHSMVirtualNetworkRuleResponseArrayOutput
	ToMHSMVirtualNetworkRuleResponseArrayOutputWithContext(context.Context) MHSMVirtualNetworkRuleResponseArrayOutput
}

type MHSMVirtualNetworkRuleResponseArray []MHSMVirtualNetworkRuleResponseInput

func (MHSMVirtualNetworkRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MHSMVirtualNetworkRuleResponse)(nil)).Elem()
}

func (i MHSMVirtualNetworkRuleResponseArray) ToMHSMVirtualNetworkRuleResponseArrayOutput() MHSMVirtualNetworkRuleResponseArrayOutput {
	return i.ToMHSMVirtualNetworkRuleResponseArrayOutputWithContext(context.Background())
}

func (i MHSMVirtualNetworkRuleResponseArray) ToMHSMVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) MHSMVirtualNetworkRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMVirtualNetworkRuleResponseArrayOutput)
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
	CreateMode                 *string                                     `pulumi:"createMode"`
	EnablePurgeProtection      *bool                                       `pulumi:"enablePurgeProtection"`
	EnableSoftDelete           *bool                                       `pulumi:"enableSoftDelete"`
	HsmUri                     string                                      `pulumi:"hsmUri"`
	InitialAdminObjectIds      []string                                    `pulumi:"initialAdminObjectIds"`
	NetworkAcls                *MHSMNetworkRuleSetResponse                 `pulumi:"networkAcls"`
	PrivateEndpointConnections []MHSMPrivateEndpointConnectionItemResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                      `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                     `pulumi:"publicNetworkAccess"`
	ScheduledPurgeDate         string                                      `pulumi:"scheduledPurgeDate"`
	SoftDeleteRetentionInDays  *int                                        `pulumi:"softDeleteRetentionInDays"`
	StatusMessage              string                                      `pulumi:"statusMessage"`
	TenantId                   *string                                     `pulumi:"tenantId"`
}





type ManagedHsmPropertiesResponseInput interface {
	pulumi.Input

	ToManagedHsmPropertiesResponseOutput() ManagedHsmPropertiesResponseOutput
	ToManagedHsmPropertiesResponseOutputWithContext(context.Context) ManagedHsmPropertiesResponseOutput
}

type ManagedHsmPropertiesResponseArgs struct {
	CreateMode                 pulumi.StringPtrInput                               `pulumi:"createMode"`
	EnablePurgeProtection      pulumi.BoolPtrInput                                 `pulumi:"enablePurgeProtection"`
	EnableSoftDelete           pulumi.BoolPtrInput                                 `pulumi:"enableSoftDelete"`
	HsmUri                     pulumi.StringInput                                  `pulumi:"hsmUri"`
	InitialAdminObjectIds      pulumi.StringArrayInput                             `pulumi:"initialAdminObjectIds"`
	NetworkAcls                MHSMNetworkRuleSetResponsePtrInput                  `pulumi:"networkAcls"`
	PrivateEndpointConnections MHSMPrivateEndpointConnectionItemResponseArrayInput `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringInput                                  `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrInput                               `pulumi:"publicNetworkAccess"`
	ScheduledPurgeDate         pulumi.StringInput                                  `pulumi:"scheduledPurgeDate"`
	SoftDeleteRetentionInDays  pulumi.IntPtrInput                                  `pulumi:"softDeleteRetentionInDays"`
	StatusMessage              pulumi.StringInput                                  `pulumi:"statusMessage"`
	TenantId                   pulumi.StringPtrInput                               `pulumi:"tenantId"`
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

func (o ManagedHsmPropertiesResponsePtrOutput) NetworkAcls() MHSMNetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *MHSMNetworkRuleSetResponse {
		if v == nil {
			return nil
		}
		return v.NetworkAcls
	}).(MHSMNetworkRuleSetResponsePtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) PrivateEndpointConnections() MHSMPrivateEndpointConnectionItemResponseArrayOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) []MHSMPrivateEndpointConnectionItemResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(MHSMPrivateEndpointConnectionItemResponseArrayOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o ManagedHsmPropertiesResponsePtrOutput) ScheduledPurgeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedHsmPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ScheduledPurgeDate
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





type RotationPolicyResponseInput interface {
	pulumi.Input

	ToRotationPolicyResponseOutput() RotationPolicyResponseOutput
	ToRotationPolicyResponseOutputWithContext(context.Context) RotationPolicyResponseOutput
}

type RotationPolicyResponseArgs struct {
	Attributes      KeyRotationPolicyAttributesResponsePtrInput `pulumi:"attributes"`
	LifetimeActions LifetimeActionResponseArrayInput            `pulumi:"lifetimeActions"`
}

func (RotationPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RotationPolicyResponse)(nil)).Elem()
}

func (i RotationPolicyResponseArgs) ToRotationPolicyResponseOutput() RotationPolicyResponseOutput {
	return i.ToRotationPolicyResponseOutputWithContext(context.Background())
}

func (i RotationPolicyResponseArgs) ToRotationPolicyResponseOutputWithContext(ctx context.Context) RotationPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RotationPolicyResponseOutput)
}

func (i RotationPolicyResponseArgs) ToRotationPolicyResponsePtrOutput() RotationPolicyResponsePtrOutput {
	return i.ToRotationPolicyResponsePtrOutputWithContext(context.Background())
}

func (i RotationPolicyResponseArgs) ToRotationPolicyResponsePtrOutputWithContext(ctx context.Context) RotationPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RotationPolicyResponseOutput).ToRotationPolicyResponsePtrOutputWithContext(ctx)
}









type RotationPolicyResponsePtrInput interface {
	pulumi.Input

	ToRotationPolicyResponsePtrOutput() RotationPolicyResponsePtrOutput
	ToRotationPolicyResponsePtrOutputWithContext(context.Context) RotationPolicyResponsePtrOutput
}

type rotationPolicyResponsePtrType RotationPolicyResponseArgs

func RotationPolicyResponsePtr(v *RotationPolicyResponseArgs) RotationPolicyResponsePtrInput {
	return (*rotationPolicyResponsePtrType)(v)
}

func (*rotationPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RotationPolicyResponse)(nil)).Elem()
}

func (i *rotationPolicyResponsePtrType) ToRotationPolicyResponsePtrOutput() RotationPolicyResponsePtrOutput {
	return i.ToRotationPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *rotationPolicyResponsePtrType) ToRotationPolicyResponsePtrOutputWithContext(ctx context.Context) RotationPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RotationPolicyResponsePtrOutput)
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

func (o RotationPolicyResponseOutput) ToRotationPolicyResponsePtrOutput() RotationPolicyResponsePtrOutput {
	return o.ToRotationPolicyResponsePtrOutputWithContext(context.Background())
}

func (o RotationPolicyResponseOutput) ToRotationPolicyResponsePtrOutputWithContext(ctx context.Context) RotationPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RotationPolicyResponse) *RotationPolicyResponse {
		return &v
	}).(RotationPolicyResponsePtrOutput)
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





type TriggerResponseInput interface {
	pulumi.Input

	ToTriggerResponseOutput() TriggerResponseOutput
	ToTriggerResponseOutputWithContext(context.Context) TriggerResponseOutput
}

type TriggerResponseArgs struct {
	TimeAfterCreate  pulumi.StringPtrInput `pulumi:"timeAfterCreate"`
	TimeBeforeExpiry pulumi.StringPtrInput `pulumi:"timeBeforeExpiry"`
}

func (TriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerResponse)(nil)).Elem()
}

func (i TriggerResponseArgs) ToTriggerResponseOutput() TriggerResponseOutput {
	return i.ToTriggerResponseOutputWithContext(context.Background())
}

func (i TriggerResponseArgs) ToTriggerResponseOutputWithContext(ctx context.Context) TriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerResponseOutput)
}

func (i TriggerResponseArgs) ToTriggerResponsePtrOutput() TriggerResponsePtrOutput {
	return i.ToTriggerResponsePtrOutputWithContext(context.Background())
}

func (i TriggerResponseArgs) ToTriggerResponsePtrOutputWithContext(ctx context.Context) TriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerResponseOutput).ToTriggerResponsePtrOutputWithContext(ctx)
}









type TriggerResponsePtrInput interface {
	pulumi.Input

	ToTriggerResponsePtrOutput() TriggerResponsePtrOutput
	ToTriggerResponsePtrOutputWithContext(context.Context) TriggerResponsePtrOutput
}

type triggerResponsePtrType TriggerResponseArgs

func TriggerResponsePtr(v *TriggerResponseArgs) TriggerResponsePtrInput {
	return (*triggerResponsePtrType)(v)
}

func (*triggerResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerResponse)(nil)).Elem()
}

func (i *triggerResponsePtrType) ToTriggerResponsePtrOutput() TriggerResponsePtrOutput {
	return i.ToTriggerResponsePtrOutputWithContext(context.Background())
}

func (i *triggerResponsePtrType) ToTriggerResponsePtrOutputWithContext(ctx context.Context) TriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerResponsePtrOutput)
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

func (o TriggerResponseOutput) ToTriggerResponsePtrOutput() TriggerResponsePtrOutput {
	return o.ToTriggerResponsePtrOutputWithContext(context.Background())
}

func (o TriggerResponseOutput) ToTriggerResponsePtrOutputWithContext(ctx context.Context) TriggerResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TriggerResponse) *TriggerResponse {
		return &v
	}).(TriggerResponsePtrOutput)
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

func (o VaultPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
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





type VaultPropertiesResponseInput interface {
	pulumi.Input

	ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput
	ToVaultPropertiesResponseOutputWithContext(context.Context) VaultPropertiesResponseOutput
}

type VaultPropertiesResponseArgs struct {
	AccessPolicies               AccessPolicyEntryResponseArrayInput             `pulumi:"accessPolicies"`
	EnablePurgeProtection        pulumi.BoolPtrInput                             `pulumi:"enablePurgeProtection"`
	EnableRbacAuthorization      pulumi.BoolPtrInput                             `pulumi:"enableRbacAuthorization"`
	EnableSoftDelete             pulumi.BoolPtrInput                             `pulumi:"enableSoftDelete"`
	EnabledForDeployment         pulumi.BoolPtrInput                             `pulumi:"enabledForDeployment"`
	EnabledForDiskEncryption     pulumi.BoolPtrInput                             `pulumi:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment pulumi.BoolPtrInput                             `pulumi:"enabledForTemplateDeployment"`
	HsmPoolResourceId            pulumi.StringInput                              `pulumi:"hsmPoolResourceId"`
	NetworkAcls                  NetworkRuleSetResponsePtrInput                  `pulumi:"networkAcls"`
	PrivateEndpointConnections   PrivateEndpointConnectionItemResponseArrayInput `pulumi:"privateEndpointConnections"`
	ProvisioningState            pulumi.StringPtrInput                           `pulumi:"provisioningState"`
	PublicNetworkAccess          pulumi.StringPtrInput                           `pulumi:"publicNetworkAccess"`
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

func (o VaultPropertiesResponsePtrOutput) HsmPoolResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HsmPoolResourceId
	}).(pulumi.StringPtrOutput)
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

func (o VaultPropertiesResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
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
	pulumi.RegisterInputType(reflect.TypeOf((*ActionInput)(nil)).Elem(), ActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionPtrInput)(nil)).Elem(), ActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionResponseInput)(nil)).Elem(), ActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionResponsePtrInput)(nil)).Elem(), ActionResponseArgs{})
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
	pulumi.RegisterInputType(reflect.TypeOf((*KeyReleasePolicyInput)(nil)).Elem(), KeyReleasePolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyReleasePolicyPtrInput)(nil)).Elem(), KeyReleasePolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyReleasePolicyResponseInput)(nil)).Elem(), KeyReleasePolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyReleasePolicyResponsePtrInput)(nil)).Elem(), KeyReleasePolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRotationPolicyAttributesInput)(nil)).Elem(), KeyRotationPolicyAttributesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRotationPolicyAttributesPtrInput)(nil)).Elem(), KeyRotationPolicyAttributesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRotationPolicyAttributesResponseInput)(nil)).Elem(), KeyRotationPolicyAttributesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyRotationPolicyAttributesResponsePtrInput)(nil)).Elem(), KeyRotationPolicyAttributesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LifetimeActionInput)(nil)).Elem(), LifetimeActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LifetimeActionArrayInput)(nil)).Elem(), LifetimeActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LifetimeActionResponseInput)(nil)).Elem(), LifetimeActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LifetimeActionResponseArrayInput)(nil)).Elem(), LifetimeActionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMIPRuleInput)(nil)).Elem(), MHSMIPRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMIPRuleArrayInput)(nil)).Elem(), MHSMIPRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMIPRuleResponseInput)(nil)).Elem(), MHSMIPRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMIPRuleResponseArrayInput)(nil)).Elem(), MHSMIPRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMNetworkRuleSetInput)(nil)).Elem(), MHSMNetworkRuleSetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMNetworkRuleSetPtrInput)(nil)).Elem(), MHSMNetworkRuleSetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMNetworkRuleSetResponseInput)(nil)).Elem(), MHSMNetworkRuleSetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMNetworkRuleSetResponsePtrInput)(nil)).Elem(), MHSMNetworkRuleSetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMPrivateEndpointConnectionItemResponseInput)(nil)).Elem(), MHSMPrivateEndpointConnectionItemResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMPrivateEndpointConnectionItemResponseArrayInput)(nil)).Elem(), MHSMPrivateEndpointConnectionItemResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMPrivateEndpointResponseInput)(nil)).Elem(), MHSMPrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMPrivateEndpointResponsePtrInput)(nil)).Elem(), MHSMPrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMPrivateLinkServiceConnectionStateInput)(nil)).Elem(), MHSMPrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMPrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), MHSMPrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMPrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), MHSMPrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMPrivateLinkServiceConnectionStateResponsePtrInput)(nil)).Elem(), MHSMPrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMVirtualNetworkRuleInput)(nil)).Elem(), MHSMVirtualNetworkRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMVirtualNetworkRuleArrayInput)(nil)).Elem(), MHSMVirtualNetworkRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMVirtualNetworkRuleResponseInput)(nil)).Elem(), MHSMVirtualNetworkRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MHSMVirtualNetworkRuleResponseArrayInput)(nil)).Elem(), MHSMVirtualNetworkRuleResponseArray{})
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
	pulumi.RegisterInputType(reflect.TypeOf((*RotationPolicyInput)(nil)).Elem(), RotationPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RotationPolicyPtrInput)(nil)).Elem(), RotationPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RotationPolicyResponseInput)(nil)).Elem(), RotationPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RotationPolicyResponsePtrInput)(nil)).Elem(), RotationPolicyResponseArgs{})
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
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), TriggerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerPtrInput)(nil)).Elem(), TriggerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerResponseInput)(nil)).Elem(), TriggerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerResponsePtrInput)(nil)).Elem(), TriggerResponseArgs{})
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
	pulumi.RegisterOutputType(KeyPropertiesPtrOutput{})
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
	pulumi.RegisterOutputType(RotationPolicyOutput{})
	pulumi.RegisterOutputType(RotationPolicyPtrOutput{})
	pulumi.RegisterOutputType(RotationPolicyResponseOutput{})
	pulumi.RegisterOutputType(RotationPolicyResponsePtrOutput{})
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
	pulumi.RegisterOutputType(TriggerOutput{})
	pulumi.RegisterOutputType(TriggerPtrOutput{})
	pulumi.RegisterOutputType(TriggerResponseOutput{})
	pulumi.RegisterOutputType(TriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesOutput{})
	pulumi.RegisterOutputType(VaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
