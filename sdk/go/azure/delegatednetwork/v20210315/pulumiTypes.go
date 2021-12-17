


package v20210315

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ControllerDetailsType struct {
	Id *string `pulumi:"id"`
}





type ControllerDetailsTypeInput interface {
	pulumi.Input

	ToControllerDetailsTypeOutput() ControllerDetailsTypeOutput
	ToControllerDetailsTypeOutputWithContext(context.Context) ControllerDetailsTypeOutput
}

type ControllerDetailsTypeArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ControllerDetailsTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ControllerDetailsType)(nil)).Elem()
}

func (i ControllerDetailsTypeArgs) ToControllerDetailsTypeOutput() ControllerDetailsTypeOutput {
	return i.ToControllerDetailsTypeOutputWithContext(context.Background())
}

func (i ControllerDetailsTypeArgs) ToControllerDetailsTypeOutputWithContext(ctx context.Context) ControllerDetailsTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerDetailsTypeOutput)
}

func (i ControllerDetailsTypeArgs) ToControllerDetailsTypePtrOutput() ControllerDetailsTypePtrOutput {
	return i.ToControllerDetailsTypePtrOutputWithContext(context.Background())
}

func (i ControllerDetailsTypeArgs) ToControllerDetailsTypePtrOutputWithContext(ctx context.Context) ControllerDetailsTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerDetailsTypeOutput).ToControllerDetailsTypePtrOutputWithContext(ctx)
}









type ControllerDetailsTypePtrInput interface {
	pulumi.Input

	ToControllerDetailsTypePtrOutput() ControllerDetailsTypePtrOutput
	ToControllerDetailsTypePtrOutputWithContext(context.Context) ControllerDetailsTypePtrOutput
}

type controllerDetailsTypePtrType ControllerDetailsTypeArgs

func ControllerDetailsTypePtr(v *ControllerDetailsTypeArgs) ControllerDetailsTypePtrInput {
	return (*controllerDetailsTypePtrType)(v)
}

func (*controllerDetailsTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ControllerDetailsType)(nil)).Elem()
}

func (i *controllerDetailsTypePtrType) ToControllerDetailsTypePtrOutput() ControllerDetailsTypePtrOutput {
	return i.ToControllerDetailsTypePtrOutputWithContext(context.Background())
}

func (i *controllerDetailsTypePtrType) ToControllerDetailsTypePtrOutputWithContext(ctx context.Context) ControllerDetailsTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ControllerDetailsTypePtrOutput)
}

type ControllerDetailsTypeOutput struct{ *pulumi.OutputState }

func (ControllerDetailsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControllerDetailsType)(nil)).Elem()
}

func (o ControllerDetailsTypeOutput) ToControllerDetailsTypeOutput() ControllerDetailsTypeOutput {
	return o
}

func (o ControllerDetailsTypeOutput) ToControllerDetailsTypeOutputWithContext(ctx context.Context) ControllerDetailsTypeOutput {
	return o
}

func (o ControllerDetailsTypeOutput) ToControllerDetailsTypePtrOutput() ControllerDetailsTypePtrOutput {
	return o.ToControllerDetailsTypePtrOutputWithContext(context.Background())
}

func (o ControllerDetailsTypeOutput) ToControllerDetailsTypePtrOutputWithContext(ctx context.Context) ControllerDetailsTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ControllerDetailsType) *ControllerDetailsType {
		return &v
	}).(ControllerDetailsTypePtrOutput)
}

func (o ControllerDetailsTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControllerDetailsType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ControllerDetailsTypePtrOutput struct{ *pulumi.OutputState }

func (ControllerDetailsTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControllerDetailsType)(nil)).Elem()
}

func (o ControllerDetailsTypePtrOutput) ToControllerDetailsTypePtrOutput() ControllerDetailsTypePtrOutput {
	return o
}

func (o ControllerDetailsTypePtrOutput) ToControllerDetailsTypePtrOutputWithContext(ctx context.Context) ControllerDetailsTypePtrOutput {
	return o
}

func (o ControllerDetailsTypePtrOutput) Elem() ControllerDetailsTypeOutput {
	return o.ApplyT(func(v *ControllerDetailsType) ControllerDetailsType {
		if v != nil {
			return *v
		}
		var ret ControllerDetailsType
		return ret
	}).(ControllerDetailsTypeOutput)
}

func (o ControllerDetailsTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControllerDetailsType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ControllerDetailsResponse struct {
	Id *string `pulumi:"id"`
}

type ControllerDetailsResponseOutput struct{ *pulumi.OutputState }

func (ControllerDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControllerDetailsResponse)(nil)).Elem()
}

func (o ControllerDetailsResponseOutput) ToControllerDetailsResponseOutput() ControllerDetailsResponseOutput {
	return o
}

func (o ControllerDetailsResponseOutput) ToControllerDetailsResponseOutputWithContext(ctx context.Context) ControllerDetailsResponseOutput {
	return o
}

func (o ControllerDetailsResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ControllerDetailsResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ControllerDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ControllerDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ControllerDetailsResponse)(nil)).Elem()
}

func (o ControllerDetailsResponsePtrOutput) ToControllerDetailsResponsePtrOutput() ControllerDetailsResponsePtrOutput {
	return o
}

func (o ControllerDetailsResponsePtrOutput) ToControllerDetailsResponsePtrOutputWithContext(ctx context.Context) ControllerDetailsResponsePtrOutput {
	return o
}

func (o ControllerDetailsResponsePtrOutput) Elem() ControllerDetailsResponseOutput {
	return o.ApplyT(func(v *ControllerDetailsResponse) ControllerDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ControllerDetailsResponse
		return ret
	}).(ControllerDetailsResponseOutput)
}

func (o ControllerDetailsResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ControllerDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type OrchestratorIdentity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type OrchestratorIdentityInput interface {
	pulumi.Input

	ToOrchestratorIdentityOutput() OrchestratorIdentityOutput
	ToOrchestratorIdentityOutputWithContext(context.Context) OrchestratorIdentityOutput
}

type OrchestratorIdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (OrchestratorIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratorIdentity)(nil)).Elem()
}

func (i OrchestratorIdentityArgs) ToOrchestratorIdentityOutput() OrchestratorIdentityOutput {
	return i.ToOrchestratorIdentityOutputWithContext(context.Background())
}

func (i OrchestratorIdentityArgs) ToOrchestratorIdentityOutputWithContext(ctx context.Context) OrchestratorIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratorIdentityOutput)
}

func (i OrchestratorIdentityArgs) ToOrchestratorIdentityPtrOutput() OrchestratorIdentityPtrOutput {
	return i.ToOrchestratorIdentityPtrOutputWithContext(context.Background())
}

func (i OrchestratorIdentityArgs) ToOrchestratorIdentityPtrOutputWithContext(ctx context.Context) OrchestratorIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratorIdentityOutput).ToOrchestratorIdentityPtrOutputWithContext(ctx)
}









type OrchestratorIdentityPtrInput interface {
	pulumi.Input

	ToOrchestratorIdentityPtrOutput() OrchestratorIdentityPtrOutput
	ToOrchestratorIdentityPtrOutputWithContext(context.Context) OrchestratorIdentityPtrOutput
}

type orchestratorIdentityPtrType OrchestratorIdentityArgs

func OrchestratorIdentityPtr(v *OrchestratorIdentityArgs) OrchestratorIdentityPtrInput {
	return (*orchestratorIdentityPtrType)(v)
}

func (*orchestratorIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratorIdentity)(nil)).Elem()
}

func (i *orchestratorIdentityPtrType) ToOrchestratorIdentityPtrOutput() OrchestratorIdentityPtrOutput {
	return i.ToOrchestratorIdentityPtrOutputWithContext(context.Background())
}

func (i *orchestratorIdentityPtrType) ToOrchestratorIdentityPtrOutputWithContext(ctx context.Context) OrchestratorIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrchestratorIdentityPtrOutput)
}

type OrchestratorIdentityOutput struct{ *pulumi.OutputState }

func (OrchestratorIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratorIdentity)(nil)).Elem()
}

func (o OrchestratorIdentityOutput) ToOrchestratorIdentityOutput() OrchestratorIdentityOutput {
	return o
}

func (o OrchestratorIdentityOutput) ToOrchestratorIdentityOutputWithContext(ctx context.Context) OrchestratorIdentityOutput {
	return o
}

func (o OrchestratorIdentityOutput) ToOrchestratorIdentityPtrOutput() OrchestratorIdentityPtrOutput {
	return o.ToOrchestratorIdentityPtrOutputWithContext(context.Background())
}

func (o OrchestratorIdentityOutput) ToOrchestratorIdentityPtrOutputWithContext(ctx context.Context) OrchestratorIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OrchestratorIdentity) *OrchestratorIdentity {
		return &v
	}).(OrchestratorIdentityPtrOutput)
}

func (o OrchestratorIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v OrchestratorIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type OrchestratorIdentityPtrOutput struct{ *pulumi.OutputState }

func (OrchestratorIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratorIdentity)(nil)).Elem()
}

func (o OrchestratorIdentityPtrOutput) ToOrchestratorIdentityPtrOutput() OrchestratorIdentityPtrOutput {
	return o
}

func (o OrchestratorIdentityPtrOutput) ToOrchestratorIdentityPtrOutputWithContext(ctx context.Context) OrchestratorIdentityPtrOutput {
	return o
}

func (o OrchestratorIdentityPtrOutput) Elem() OrchestratorIdentityOutput {
	return o.ApplyT(func(v *OrchestratorIdentity) OrchestratorIdentity {
		if v != nil {
			return *v
		}
		var ret OrchestratorIdentity
		return ret
	}).(OrchestratorIdentityOutput)
}

func (o OrchestratorIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *OrchestratorIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type OrchestratorIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type OrchestratorIdentityResponseOutput struct{ *pulumi.OutputState }

func (OrchestratorIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrchestratorIdentityResponse)(nil)).Elem()
}

func (o OrchestratorIdentityResponseOutput) ToOrchestratorIdentityResponseOutput() OrchestratorIdentityResponseOutput {
	return o
}

func (o OrchestratorIdentityResponseOutput) ToOrchestratorIdentityResponseOutputWithContext(ctx context.Context) OrchestratorIdentityResponseOutput {
	return o
}

func (o OrchestratorIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v OrchestratorIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o OrchestratorIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v OrchestratorIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o OrchestratorIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrchestratorIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type OrchestratorIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (OrchestratorIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrchestratorIdentityResponse)(nil)).Elem()
}

func (o OrchestratorIdentityResponsePtrOutput) ToOrchestratorIdentityResponsePtrOutput() OrchestratorIdentityResponsePtrOutput {
	return o
}

func (o OrchestratorIdentityResponsePtrOutput) ToOrchestratorIdentityResponsePtrOutputWithContext(ctx context.Context) OrchestratorIdentityResponsePtrOutput {
	return o
}

func (o OrchestratorIdentityResponsePtrOutput) Elem() OrchestratorIdentityResponseOutput {
	return o.ApplyT(func(v *OrchestratorIdentityResponse) OrchestratorIdentityResponse {
		if v != nil {
			return *v
		}
		var ret OrchestratorIdentityResponse
		return ret
	}).(OrchestratorIdentityResponseOutput)
}

func (o OrchestratorIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrchestratorIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o OrchestratorIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrchestratorIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o OrchestratorIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrchestratorIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SubnetDetails struct {
	Id *string `pulumi:"id"`
}





type SubnetDetailsInput interface {
	pulumi.Input

	ToSubnetDetailsOutput() SubnetDetailsOutput
	ToSubnetDetailsOutputWithContext(context.Context) SubnetDetailsOutput
}

type SubnetDetailsArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubnetDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetDetails)(nil)).Elem()
}

func (i SubnetDetailsArgs) ToSubnetDetailsOutput() SubnetDetailsOutput {
	return i.ToSubnetDetailsOutputWithContext(context.Background())
}

func (i SubnetDetailsArgs) ToSubnetDetailsOutputWithContext(ctx context.Context) SubnetDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetDetailsOutput)
}

func (i SubnetDetailsArgs) ToSubnetDetailsPtrOutput() SubnetDetailsPtrOutput {
	return i.ToSubnetDetailsPtrOutputWithContext(context.Background())
}

func (i SubnetDetailsArgs) ToSubnetDetailsPtrOutputWithContext(ctx context.Context) SubnetDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetDetailsOutput).ToSubnetDetailsPtrOutputWithContext(ctx)
}









type SubnetDetailsPtrInput interface {
	pulumi.Input

	ToSubnetDetailsPtrOutput() SubnetDetailsPtrOutput
	ToSubnetDetailsPtrOutputWithContext(context.Context) SubnetDetailsPtrOutput
}

type subnetDetailsPtrType SubnetDetailsArgs

func SubnetDetailsPtr(v *SubnetDetailsArgs) SubnetDetailsPtrInput {
	return (*subnetDetailsPtrType)(v)
}

func (*subnetDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetDetails)(nil)).Elem()
}

func (i *subnetDetailsPtrType) ToSubnetDetailsPtrOutput() SubnetDetailsPtrOutput {
	return i.ToSubnetDetailsPtrOutputWithContext(context.Background())
}

func (i *subnetDetailsPtrType) ToSubnetDetailsPtrOutputWithContext(ctx context.Context) SubnetDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetDetailsPtrOutput)
}

type SubnetDetailsOutput struct{ *pulumi.OutputState }

func (SubnetDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetDetails)(nil)).Elem()
}

func (o SubnetDetailsOutput) ToSubnetDetailsOutput() SubnetDetailsOutput {
	return o
}

func (o SubnetDetailsOutput) ToSubnetDetailsOutputWithContext(ctx context.Context) SubnetDetailsOutput {
	return o
}

func (o SubnetDetailsOutput) ToSubnetDetailsPtrOutput() SubnetDetailsPtrOutput {
	return o.ToSubnetDetailsPtrOutputWithContext(context.Background())
}

func (o SubnetDetailsOutput) ToSubnetDetailsPtrOutputWithContext(ctx context.Context) SubnetDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetDetails) *SubnetDetails {
		return &v
	}).(SubnetDetailsPtrOutput)
}

func (o SubnetDetailsOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetDetails) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubnetDetailsPtrOutput struct{ *pulumi.OutputState }

func (SubnetDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetDetails)(nil)).Elem()
}

func (o SubnetDetailsPtrOutput) ToSubnetDetailsPtrOutput() SubnetDetailsPtrOutput {
	return o
}

func (o SubnetDetailsPtrOutput) ToSubnetDetailsPtrOutputWithContext(ctx context.Context) SubnetDetailsPtrOutput {
	return o
}

func (o SubnetDetailsPtrOutput) Elem() SubnetDetailsOutput {
	return o.ApplyT(func(v *SubnetDetails) SubnetDetails {
		if v != nil {
			return *v
		}
		var ret SubnetDetails
		return ret
	}).(SubnetDetailsOutput)
}

func (o SubnetDetailsPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetDetails) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubnetDetailsResponse struct {
	Id *string `pulumi:"id"`
}

type SubnetDetailsResponseOutput struct{ *pulumi.OutputState }

func (SubnetDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetDetailsResponse)(nil)).Elem()
}

func (o SubnetDetailsResponseOutput) ToSubnetDetailsResponseOutput() SubnetDetailsResponseOutput {
	return o
}

func (o SubnetDetailsResponseOutput) ToSubnetDetailsResponseOutputWithContext(ctx context.Context) SubnetDetailsResponseOutput {
	return o
}

func (o SubnetDetailsResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetDetailsResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubnetDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetDetailsResponse)(nil)).Elem()
}

func (o SubnetDetailsResponsePtrOutput) ToSubnetDetailsResponsePtrOutput() SubnetDetailsResponsePtrOutput {
	return o
}

func (o SubnetDetailsResponsePtrOutput) ToSubnetDetailsResponsePtrOutputWithContext(ctx context.Context) SubnetDetailsResponsePtrOutput {
	return o
}

func (o SubnetDetailsResponsePtrOutput) Elem() SubnetDetailsResponseOutput {
	return o.ApplyT(func(v *SubnetDetailsResponse) SubnetDetailsResponse {
		if v != nil {
			return *v
		}
		var ret SubnetDetailsResponse
		return ret
	}).(SubnetDetailsResponseOutput)
}

func (o SubnetDetailsResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ControllerDetailsTypeOutput{})
	pulumi.RegisterOutputType(ControllerDetailsTypePtrOutput{})
	pulumi.RegisterOutputType(ControllerDetailsResponseOutput{})
	pulumi.RegisterOutputType(ControllerDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(OrchestratorIdentityOutput{})
	pulumi.RegisterOutputType(OrchestratorIdentityPtrOutput{})
	pulumi.RegisterOutputType(OrchestratorIdentityResponseOutput{})
	pulumi.RegisterOutputType(OrchestratorIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetDetailsOutput{})
	pulumi.RegisterOutputType(SubnetDetailsPtrOutput{})
	pulumi.RegisterOutputType(SubnetDetailsResponseOutput{})
	pulumi.RegisterOutputType(SubnetDetailsResponsePtrOutput{})
}
