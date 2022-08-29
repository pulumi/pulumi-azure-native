


package v20160901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplianceArtifact struct {
	Name *string                `pulumi:"name"`
	Type *ApplianceArtifactType `pulumi:"type"`
	Uri  *string                `pulumi:"uri"`
}





type ApplianceArtifactInput interface {
	pulumi.Input

	ToApplianceArtifactOutput() ApplianceArtifactOutput
	ToApplianceArtifactOutputWithContext(context.Context) ApplianceArtifactOutput
}

type ApplianceArtifactArgs struct {
	Name pulumi.StringPtrInput         `pulumi:"name"`
	Type ApplianceArtifactTypePtrInput `pulumi:"type"`
	Uri  pulumi.StringPtrInput         `pulumi:"uri"`
}

func (ApplianceArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceArtifact)(nil)).Elem()
}

func (i ApplianceArtifactArgs) ToApplianceArtifactOutput() ApplianceArtifactOutput {
	return i.ToApplianceArtifactOutputWithContext(context.Background())
}

func (i ApplianceArtifactArgs) ToApplianceArtifactOutputWithContext(ctx context.Context) ApplianceArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceArtifactOutput)
}





type ApplianceArtifactArrayInput interface {
	pulumi.Input

	ToApplianceArtifactArrayOutput() ApplianceArtifactArrayOutput
	ToApplianceArtifactArrayOutputWithContext(context.Context) ApplianceArtifactArrayOutput
}

type ApplianceArtifactArray []ApplianceArtifactInput

func (ApplianceArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceArtifact)(nil)).Elem()
}

func (i ApplianceArtifactArray) ToApplianceArtifactArrayOutput() ApplianceArtifactArrayOutput {
	return i.ToApplianceArtifactArrayOutputWithContext(context.Background())
}

func (i ApplianceArtifactArray) ToApplianceArtifactArrayOutputWithContext(ctx context.Context) ApplianceArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceArtifactArrayOutput)
}

type ApplianceArtifactOutput struct{ *pulumi.OutputState }

func (ApplianceArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceArtifact)(nil)).Elem()
}

func (o ApplianceArtifactOutput) ToApplianceArtifactOutput() ApplianceArtifactOutput {
	return o
}

func (o ApplianceArtifactOutput) ToApplianceArtifactOutputWithContext(ctx context.Context) ApplianceArtifactOutput {
	return o
}

func (o ApplianceArtifactOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplianceArtifact) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplianceArtifactOutput) Type() ApplianceArtifactTypePtrOutput {
	return o.ApplyT(func(v ApplianceArtifact) *ApplianceArtifactType { return v.Type }).(ApplianceArtifactTypePtrOutput)
}

func (o ApplianceArtifactOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplianceArtifact) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ApplianceArtifactArrayOutput struct{ *pulumi.OutputState }

func (ApplianceArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceArtifact)(nil)).Elem()
}

func (o ApplianceArtifactArrayOutput) ToApplianceArtifactArrayOutput() ApplianceArtifactArrayOutput {
	return o
}

func (o ApplianceArtifactArrayOutput) ToApplianceArtifactArrayOutputWithContext(ctx context.Context) ApplianceArtifactArrayOutput {
	return o
}

func (o ApplianceArtifactArrayOutput) Index(i pulumi.IntInput) ApplianceArtifactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplianceArtifact {
		return vs[0].([]ApplianceArtifact)[vs[1].(int)]
	}).(ApplianceArtifactOutput)
}

type ApplianceArtifactResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
	Uri  *string `pulumi:"uri"`
}

type ApplianceArtifactResponseOutput struct{ *pulumi.OutputState }

func (ApplianceArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceArtifactResponse)(nil)).Elem()
}

func (o ApplianceArtifactResponseOutput) ToApplianceArtifactResponseOutput() ApplianceArtifactResponseOutput {
	return o
}

func (o ApplianceArtifactResponseOutput) ToApplianceArtifactResponseOutputWithContext(ctx context.Context) ApplianceArtifactResponseOutput {
	return o
}

func (o ApplianceArtifactResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplianceArtifactResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplianceArtifactResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplianceArtifactResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ApplianceArtifactResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplianceArtifactResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ApplianceArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplianceArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceArtifactResponse)(nil)).Elem()
}

func (o ApplianceArtifactResponseArrayOutput) ToApplianceArtifactResponseArrayOutput() ApplianceArtifactResponseArrayOutput {
	return o
}

func (o ApplianceArtifactResponseArrayOutput) ToApplianceArtifactResponseArrayOutputWithContext(ctx context.Context) ApplianceArtifactResponseArrayOutput {
	return o
}

func (o ApplianceArtifactResponseArrayOutput) Index(i pulumi.IntInput) ApplianceArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplianceArtifactResponse {
		return vs[0].([]ApplianceArtifactResponse)[vs[1].(int)]
	}).(ApplianceArtifactResponseOutput)
}

type ApplianceProviderAuthorization struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}





type ApplianceProviderAuthorizationInput interface {
	pulumi.Input

	ToApplianceProviderAuthorizationOutput() ApplianceProviderAuthorizationOutput
	ToApplianceProviderAuthorizationOutputWithContext(context.Context) ApplianceProviderAuthorizationOutput
}

type ApplianceProviderAuthorizationArgs struct {
	PrincipalId      pulumi.StringInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringInput `pulumi:"roleDefinitionId"`
}

func (ApplianceProviderAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceProviderAuthorization)(nil)).Elem()
}

func (i ApplianceProviderAuthorizationArgs) ToApplianceProviderAuthorizationOutput() ApplianceProviderAuthorizationOutput {
	return i.ToApplianceProviderAuthorizationOutputWithContext(context.Background())
}

func (i ApplianceProviderAuthorizationArgs) ToApplianceProviderAuthorizationOutputWithContext(ctx context.Context) ApplianceProviderAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceProviderAuthorizationOutput)
}





type ApplianceProviderAuthorizationArrayInput interface {
	pulumi.Input

	ToApplianceProviderAuthorizationArrayOutput() ApplianceProviderAuthorizationArrayOutput
	ToApplianceProviderAuthorizationArrayOutputWithContext(context.Context) ApplianceProviderAuthorizationArrayOutput
}

type ApplianceProviderAuthorizationArray []ApplianceProviderAuthorizationInput

func (ApplianceProviderAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceProviderAuthorization)(nil)).Elem()
}

func (i ApplianceProviderAuthorizationArray) ToApplianceProviderAuthorizationArrayOutput() ApplianceProviderAuthorizationArrayOutput {
	return i.ToApplianceProviderAuthorizationArrayOutputWithContext(context.Background())
}

func (i ApplianceProviderAuthorizationArray) ToApplianceProviderAuthorizationArrayOutputWithContext(ctx context.Context) ApplianceProviderAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceProviderAuthorizationArrayOutput)
}

type ApplianceProviderAuthorizationOutput struct{ *pulumi.OutputState }

func (ApplianceProviderAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceProviderAuthorization)(nil)).Elem()
}

func (o ApplianceProviderAuthorizationOutput) ToApplianceProviderAuthorizationOutput() ApplianceProviderAuthorizationOutput {
	return o
}

func (o ApplianceProviderAuthorizationOutput) ToApplianceProviderAuthorizationOutputWithContext(ctx context.Context) ApplianceProviderAuthorizationOutput {
	return o
}

func (o ApplianceProviderAuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceProviderAuthorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ApplianceProviderAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceProviderAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type ApplianceProviderAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (ApplianceProviderAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceProviderAuthorization)(nil)).Elem()
}

func (o ApplianceProviderAuthorizationArrayOutput) ToApplianceProviderAuthorizationArrayOutput() ApplianceProviderAuthorizationArrayOutput {
	return o
}

func (o ApplianceProviderAuthorizationArrayOutput) ToApplianceProviderAuthorizationArrayOutputWithContext(ctx context.Context) ApplianceProviderAuthorizationArrayOutput {
	return o
}

func (o ApplianceProviderAuthorizationArrayOutput) Index(i pulumi.IntInput) ApplianceProviderAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplianceProviderAuthorization {
		return vs[0].([]ApplianceProviderAuthorization)[vs[1].(int)]
	}).(ApplianceProviderAuthorizationOutput)
}

type ApplianceProviderAuthorizationResponse struct {
	PrincipalId      string `pulumi:"principalId"`
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
}

type ApplianceProviderAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (ApplianceProviderAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplianceProviderAuthorizationResponse)(nil)).Elem()
}

func (o ApplianceProviderAuthorizationResponseOutput) ToApplianceProviderAuthorizationResponseOutput() ApplianceProviderAuthorizationResponseOutput {
	return o
}

func (o ApplianceProviderAuthorizationResponseOutput) ToApplianceProviderAuthorizationResponseOutputWithContext(ctx context.Context) ApplianceProviderAuthorizationResponseOutput {
	return o
}

func (o ApplianceProviderAuthorizationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceProviderAuthorizationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ApplianceProviderAuthorizationResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplianceProviderAuthorizationResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type ApplianceProviderAuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplianceProviderAuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplianceProviderAuthorizationResponse)(nil)).Elem()
}

func (o ApplianceProviderAuthorizationResponseArrayOutput) ToApplianceProviderAuthorizationResponseArrayOutput() ApplianceProviderAuthorizationResponseArrayOutput {
	return o
}

func (o ApplianceProviderAuthorizationResponseArrayOutput) ToApplianceProviderAuthorizationResponseArrayOutputWithContext(ctx context.Context) ApplianceProviderAuthorizationResponseArrayOutput {
	return o
}

func (o ApplianceProviderAuthorizationResponseArrayOutput) Index(i pulumi.IntInput) ApplianceProviderAuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplianceProviderAuthorizationResponse {
		return vs[0].([]ApplianceProviderAuthorizationResponse)[vs[1].(int)]
	}).(ApplianceProviderAuthorizationResponseOutput)
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
	pulumi.RegisterOutputType(ApplianceArtifactOutput{})
	pulumi.RegisterOutputType(ApplianceArtifactArrayOutput{})
	pulumi.RegisterOutputType(ApplianceArtifactResponseOutput{})
	pulumi.RegisterOutputType(ApplianceArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplianceProviderAuthorizationOutput{})
	pulumi.RegisterOutputType(ApplianceProviderAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(ApplianceProviderAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(ApplianceProviderAuthorizationResponseArrayOutput{})
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
