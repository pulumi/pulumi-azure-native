


package v20170901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationArtifact struct {
	Name *string                  `pulumi:"name"`
	Type *ApplicationArtifactType `pulumi:"type"`
	Uri  *string                  `pulumi:"uri"`
}





type ApplicationArtifactInput interface {
	pulumi.Input

	ToApplicationArtifactOutput() ApplicationArtifactOutput
	ToApplicationArtifactOutputWithContext(context.Context) ApplicationArtifactOutput
}

type ApplicationArtifactArgs struct {
	Name pulumi.StringPtrInput           `pulumi:"name"`
	Type ApplicationArtifactTypePtrInput `pulumi:"type"`
	Uri  pulumi.StringPtrInput           `pulumi:"uri"`
}

func (ApplicationArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationArtifact)(nil)).Elem()
}

func (i ApplicationArtifactArgs) ToApplicationArtifactOutput() ApplicationArtifactOutput {
	return i.ToApplicationArtifactOutputWithContext(context.Background())
}

func (i ApplicationArtifactArgs) ToApplicationArtifactOutputWithContext(ctx context.Context) ApplicationArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArtifactOutput)
}





type ApplicationArtifactArrayInput interface {
	pulumi.Input

	ToApplicationArtifactArrayOutput() ApplicationArtifactArrayOutput
	ToApplicationArtifactArrayOutputWithContext(context.Context) ApplicationArtifactArrayOutput
}

type ApplicationArtifactArray []ApplicationArtifactInput

func (ApplicationArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationArtifact)(nil)).Elem()
}

func (i ApplicationArtifactArray) ToApplicationArtifactArrayOutput() ApplicationArtifactArrayOutput {
	return i.ToApplicationArtifactArrayOutputWithContext(context.Background())
}

func (i ApplicationArtifactArray) ToApplicationArtifactArrayOutputWithContext(ctx context.Context) ApplicationArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArtifactArrayOutput)
}

type ApplicationArtifactOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationArtifact)(nil)).Elem()
}

func (o ApplicationArtifactOutput) ToApplicationArtifactOutput() ApplicationArtifactOutput {
	return o
}

func (o ApplicationArtifactOutput) ToApplicationArtifactOutputWithContext(ctx context.Context) ApplicationArtifactOutput {
	return o
}

func (o ApplicationArtifactOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationArtifact) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationArtifactOutput) Type() ApplicationArtifactTypePtrOutput {
	return o.ApplyT(func(v ApplicationArtifact) *ApplicationArtifactType { return v.Type }).(ApplicationArtifactTypePtrOutput)
}

func (o ApplicationArtifactOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationArtifact) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ApplicationArtifactArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationArtifact)(nil)).Elem()
}

func (o ApplicationArtifactArrayOutput) ToApplicationArtifactArrayOutput() ApplicationArtifactArrayOutput {
	return o
}

func (o ApplicationArtifactArrayOutput) ToApplicationArtifactArrayOutputWithContext(ctx context.Context) ApplicationArtifactArrayOutput {
	return o
}

func (o ApplicationArtifactArrayOutput) Index(i pulumi.IntInput) ApplicationArtifactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationArtifact {
		return vs[0].([]ApplicationArtifact)[vs[1].(int)]
	}).(ApplicationArtifactOutput)
}

type ApplicationArtifactResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
	Uri  *string `pulumi:"uri"`
}





type ApplicationArtifactResponseInput interface {
	pulumi.Input

	ToApplicationArtifactResponseOutput() ApplicationArtifactResponseOutput
	ToApplicationArtifactResponseOutputWithContext(context.Context) ApplicationArtifactResponseOutput
}

type ApplicationArtifactResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
	Uri  pulumi.StringPtrInput `pulumi:"uri"`
}

func (ApplicationArtifactResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationArtifactResponse)(nil)).Elem()
}

func (i ApplicationArtifactResponseArgs) ToApplicationArtifactResponseOutput() ApplicationArtifactResponseOutput {
	return i.ToApplicationArtifactResponseOutputWithContext(context.Background())
}

func (i ApplicationArtifactResponseArgs) ToApplicationArtifactResponseOutputWithContext(ctx context.Context) ApplicationArtifactResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArtifactResponseOutput)
}





type ApplicationArtifactResponseArrayInput interface {
	pulumi.Input

	ToApplicationArtifactResponseArrayOutput() ApplicationArtifactResponseArrayOutput
	ToApplicationArtifactResponseArrayOutputWithContext(context.Context) ApplicationArtifactResponseArrayOutput
}

type ApplicationArtifactResponseArray []ApplicationArtifactResponseInput

func (ApplicationArtifactResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationArtifactResponse)(nil)).Elem()
}

func (i ApplicationArtifactResponseArray) ToApplicationArtifactResponseArrayOutput() ApplicationArtifactResponseArrayOutput {
	return i.ToApplicationArtifactResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationArtifactResponseArray) ToApplicationArtifactResponseArrayOutputWithContext(ctx context.Context) ApplicationArtifactResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArtifactResponseArrayOutput)
}

type ApplicationArtifactResponseOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationArtifactResponse)(nil)).Elem()
}

func (o ApplicationArtifactResponseOutput) ToApplicationArtifactResponseOutput() ApplicationArtifactResponseOutput {
	return o
}

func (o ApplicationArtifactResponseOutput) ToApplicationArtifactResponseOutputWithContext(ctx context.Context) ApplicationArtifactResponseOutput {
	return o
}

func (o ApplicationArtifactResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationArtifactResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationArtifactResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationArtifactResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ApplicationArtifactResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationArtifactResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ApplicationArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationArtifactResponse)(nil)).Elem()
}

func (o ApplicationArtifactResponseArrayOutput) ToApplicationArtifactResponseArrayOutput() ApplicationArtifactResponseArrayOutput {
	return o
}

func (o ApplicationArtifactResponseArrayOutput) ToApplicationArtifactResponseArrayOutputWithContext(ctx context.Context) ApplicationArtifactResponseArrayOutput {
	return o
}

func (o ApplicationArtifactResponseArrayOutput) Index(i pulumi.IntInput) ApplicationArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationArtifactResponse {
		return vs[0].([]ApplicationArtifactResponse)[vs[1].(int)]
	}).(ApplicationArtifactResponseOutput)
}

type ApplicationProviderAuthorization struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type ApplicationProviderAuthorizationInput interface {
	pulumi.Input

	ToApplicationProviderAuthorizationOutput() ApplicationProviderAuthorizationOutput
	ToApplicationProviderAuthorizationOutputWithContext(context.Context) ApplicationProviderAuthorizationOutput
}

type ApplicationProviderAuthorizationArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (ApplicationProviderAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationProviderAuthorization)(nil)).Elem()
}

func (i ApplicationProviderAuthorizationArgs) ToApplicationProviderAuthorizationOutput() ApplicationProviderAuthorizationOutput {
	return i.ToApplicationProviderAuthorizationOutputWithContext(context.Background())
}

func (i ApplicationProviderAuthorizationArgs) ToApplicationProviderAuthorizationOutputWithContext(ctx context.Context) ApplicationProviderAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProviderAuthorizationOutput)
}





type ApplicationProviderAuthorizationArrayInput interface {
	pulumi.Input

	ToApplicationProviderAuthorizationArrayOutput() ApplicationProviderAuthorizationArrayOutput
	ToApplicationProviderAuthorizationArrayOutputWithContext(context.Context) ApplicationProviderAuthorizationArrayOutput
}

type ApplicationProviderAuthorizationArray []ApplicationProviderAuthorizationInput

func (ApplicationProviderAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationProviderAuthorization)(nil)).Elem()
}

func (i ApplicationProviderAuthorizationArray) ToApplicationProviderAuthorizationArrayOutput() ApplicationProviderAuthorizationArrayOutput {
	return i.ToApplicationProviderAuthorizationArrayOutputWithContext(context.Background())
}

func (i ApplicationProviderAuthorizationArray) ToApplicationProviderAuthorizationArrayOutputWithContext(ctx context.Context) ApplicationProviderAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProviderAuthorizationArrayOutput)
}

type ApplicationProviderAuthorizationOutput struct{ *pulumi.OutputState }

func (ApplicationProviderAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationProviderAuthorization)(nil)).Elem()
}

func (o ApplicationProviderAuthorizationOutput) ToApplicationProviderAuthorizationOutput() ApplicationProviderAuthorizationOutput {
	return o
}

func (o ApplicationProviderAuthorizationOutput) ToApplicationProviderAuthorizationOutputWithContext(ctx context.Context) ApplicationProviderAuthorizationOutput {
	return o
}

func (o ApplicationProviderAuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationProviderAuthorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ApplicationProviderAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationProviderAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type ApplicationProviderAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationProviderAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationProviderAuthorization)(nil)).Elem()
}

func (o ApplicationProviderAuthorizationArrayOutput) ToApplicationProviderAuthorizationArrayOutput() ApplicationProviderAuthorizationArrayOutput {
	return o
}

func (o ApplicationProviderAuthorizationArrayOutput) ToApplicationProviderAuthorizationArrayOutputWithContext(ctx context.Context) ApplicationProviderAuthorizationArrayOutput {
	return o
}

func (o ApplicationProviderAuthorizationArrayOutput) Index(i pulumi.IntInput) ApplicationProviderAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationProviderAuthorization {
		return vs[0].([]ApplicationProviderAuthorization)[vs[1].(int)]
	}).(ApplicationProviderAuthorizationOutput)
}

type ApplicationProviderAuthorizationResponse struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type ApplicationProviderAuthorizationResponseInput interface {
	pulumi.Input

	ToApplicationProviderAuthorizationResponseOutput() ApplicationProviderAuthorizationResponseOutput
	ToApplicationProviderAuthorizationResponseOutputWithContext(context.Context) ApplicationProviderAuthorizationResponseOutput
}

type ApplicationProviderAuthorizationResponseArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (ApplicationProviderAuthorizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationProviderAuthorizationResponse)(nil)).Elem()
}

func (i ApplicationProviderAuthorizationResponseArgs) ToApplicationProviderAuthorizationResponseOutput() ApplicationProviderAuthorizationResponseOutput {
	return i.ToApplicationProviderAuthorizationResponseOutputWithContext(context.Background())
}

func (i ApplicationProviderAuthorizationResponseArgs) ToApplicationProviderAuthorizationResponseOutputWithContext(ctx context.Context) ApplicationProviderAuthorizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProviderAuthorizationResponseOutput)
}





type ApplicationProviderAuthorizationResponseArrayInput interface {
	pulumi.Input

	ToApplicationProviderAuthorizationResponseArrayOutput() ApplicationProviderAuthorizationResponseArrayOutput
	ToApplicationProviderAuthorizationResponseArrayOutputWithContext(context.Context) ApplicationProviderAuthorizationResponseArrayOutput
}

type ApplicationProviderAuthorizationResponseArray []ApplicationProviderAuthorizationResponseInput

func (ApplicationProviderAuthorizationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationProviderAuthorizationResponse)(nil)).Elem()
}

func (i ApplicationProviderAuthorizationResponseArray) ToApplicationProviderAuthorizationResponseArrayOutput() ApplicationProviderAuthorizationResponseArrayOutput {
	return i.ToApplicationProviderAuthorizationResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationProviderAuthorizationResponseArray) ToApplicationProviderAuthorizationResponseArrayOutputWithContext(ctx context.Context) ApplicationProviderAuthorizationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProviderAuthorizationResponseArrayOutput)
}

type ApplicationProviderAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationProviderAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationProviderAuthorizationResponse)(nil)).Elem()
}

func (o ApplicationProviderAuthorizationResponseOutput) ToApplicationProviderAuthorizationResponseOutput() ApplicationProviderAuthorizationResponseOutput {
	return o
}

func (o ApplicationProviderAuthorizationResponseOutput) ToApplicationProviderAuthorizationResponseOutputWithContext(ctx context.Context) ApplicationProviderAuthorizationResponseOutput {
	return o
}

func (o ApplicationProviderAuthorizationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationProviderAuthorizationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ApplicationProviderAuthorizationResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationProviderAuthorizationResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type ApplicationProviderAuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationProviderAuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationProviderAuthorizationResponse)(nil)).Elem()
}

func (o ApplicationProviderAuthorizationResponseArrayOutput) ToApplicationProviderAuthorizationResponseArrayOutput() ApplicationProviderAuthorizationResponseArrayOutput {
	return o
}

func (o ApplicationProviderAuthorizationResponseArrayOutput) ToApplicationProviderAuthorizationResponseArrayOutputWithContext(ctx context.Context) ApplicationProviderAuthorizationResponseArrayOutput {
	return o
}

func (o ApplicationProviderAuthorizationResponseArrayOutput) Index(i pulumi.IntInput) ApplicationProviderAuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationProviderAuthorizationResponse {
		return vs[0].([]ApplicationProviderAuthorizationResponse)[vs[1].(int)]
	}).(ApplicationProviderAuthorizationResponseOutput)
}

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
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

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
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

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
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

type Plan struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       string  `pulumi:"version"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name          pulumi.StringInput    `pulumi:"name"`
	Product       pulumi.StringInput    `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringInput    `pulumi:"publisher"`
	Version       pulumi.StringInput    `pulumi:"version"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i PlanArgs) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput).ToPlanPtrOutputWithContext(ctx)
}









type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func PlanPtr(v *PlanArgs) PlanPtrInput {
	return (*planPtrType)(v)
}

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Version }).(pulumi.StringOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type PlanResponse struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       string  `pulumi:"version"`
}





type PlanResponseInput interface {
	pulumi.Input

	ToPlanResponseOutput() PlanResponseOutput
	ToPlanResponseOutputWithContext(context.Context) PlanResponseOutput
}

type PlanResponseArgs struct {
	Name          pulumi.StringInput    `pulumi:"name"`
	Product       pulumi.StringInput    `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringInput    `pulumi:"publisher"`
	Version       pulumi.StringInput    `pulumi:"version"`
}

func (PlanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (i PlanResponseArgs) ToPlanResponseOutput() PlanResponseOutput {
	return i.ToPlanResponseOutputWithContext(context.Background())
}

func (i PlanResponseArgs) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseOutput)
}

func (i PlanResponseArgs) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return i.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (i PlanResponseArgs) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseOutput).ToPlanResponsePtrOutputWithContext(ctx)
}









type PlanResponsePtrInput interface {
	pulumi.Input

	ToPlanResponsePtrOutput() PlanResponsePtrOutput
	ToPlanResponsePtrOutputWithContext(context.Context) PlanResponsePtrOutput
}

type planResponsePtrType PlanResponseArgs

func PlanResponsePtr(v *PlanResponseArgs) PlanResponsePtrInput {
	return (*planResponsePtrType)(v)
}

func (*planResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (i *planResponsePtrType) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return i.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (i *planResponsePtrType) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponsePtrOutput)
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o.ToPlanResponsePtrOutputWithContext(context.Background())
}

func (o PlanResponseOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlanResponse) *PlanResponse {
		return &v
	}).(PlanResponsePtrOutput)
}

func (o PlanResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Version }).(pulumi.StringOutput)
}

type PlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) Elem() PlanResponseOutput {
	return o.ApplyT(func(v *PlanResponse) PlanResponse {
		if v != nil {
			return *v
		}
		var ret PlanResponse
		return ret
	}).(PlanResponseOutput)
}

func (o PlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Model    *string `pulumi:"model"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Model    pulumi.StringPtrInput `pulumi:"model"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Model }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Model
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

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
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
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Model    *string `pulumi:"model"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Model    pulumi.StringPtrInput `pulumi:"model"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Model }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Model
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

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArtifactInput)(nil)).Elem(), ApplicationArtifactArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArtifactArrayInput)(nil)).Elem(), ApplicationArtifactArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArtifactResponseInput)(nil)).Elem(), ApplicationArtifactResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArtifactResponseArrayInput)(nil)).Elem(), ApplicationArtifactResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationProviderAuthorizationInput)(nil)).Elem(), ApplicationProviderAuthorizationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationProviderAuthorizationArrayInput)(nil)).Elem(), ApplicationProviderAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationProviderAuthorizationResponseInput)(nil)).Elem(), ApplicationProviderAuthorizationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationProviderAuthorizationResponseArrayInput)(nil)).Elem(), ApplicationProviderAuthorizationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPtrInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponsePtrInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanInput)(nil)).Elem(), PlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanPtrInput)(nil)).Elem(), PlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanResponseInput)(nil)).Elem(), PlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanResponsePtrInput)(nil)).Elem(), PlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterOutputType(ApplicationArtifactOutput{})
	pulumi.RegisterOutputType(ApplicationArtifactArrayOutput{})
	pulumi.RegisterOutputType(ApplicationArtifactResponseOutput{})
	pulumi.RegisterOutputType(ApplicationArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationProviderAuthorizationOutput{})
	pulumi.RegisterOutputType(ApplicationProviderAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationProviderAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationProviderAuthorizationResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
