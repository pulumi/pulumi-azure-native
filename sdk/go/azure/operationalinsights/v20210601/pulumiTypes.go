


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssociatedWorkspaceResponse struct {
	AssociateDate string `pulumi:"associateDate"`
	ResourceId    string `pulumi:"resourceId"`
	WorkspaceId   string `pulumi:"workspaceId"`
	WorkspaceName string `pulumi:"workspaceName"`
}





type AssociatedWorkspaceResponseInput interface {
	pulumi.Input

	ToAssociatedWorkspaceResponseOutput() AssociatedWorkspaceResponseOutput
	ToAssociatedWorkspaceResponseOutputWithContext(context.Context) AssociatedWorkspaceResponseOutput
}

type AssociatedWorkspaceResponseArgs struct {
	AssociateDate pulumi.StringInput `pulumi:"associateDate"`
	ResourceId    pulumi.StringInput `pulumi:"resourceId"`
	WorkspaceId   pulumi.StringInput `pulumi:"workspaceId"`
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (AssociatedWorkspaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociatedWorkspaceResponse)(nil)).Elem()
}

func (i AssociatedWorkspaceResponseArgs) ToAssociatedWorkspaceResponseOutput() AssociatedWorkspaceResponseOutput {
	return i.ToAssociatedWorkspaceResponseOutputWithContext(context.Background())
}

func (i AssociatedWorkspaceResponseArgs) ToAssociatedWorkspaceResponseOutputWithContext(ctx context.Context) AssociatedWorkspaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociatedWorkspaceResponseOutput)
}





type AssociatedWorkspaceResponseArrayInput interface {
	pulumi.Input

	ToAssociatedWorkspaceResponseArrayOutput() AssociatedWorkspaceResponseArrayOutput
	ToAssociatedWorkspaceResponseArrayOutputWithContext(context.Context) AssociatedWorkspaceResponseArrayOutput
}

type AssociatedWorkspaceResponseArray []AssociatedWorkspaceResponseInput

func (AssociatedWorkspaceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssociatedWorkspaceResponse)(nil)).Elem()
}

func (i AssociatedWorkspaceResponseArray) ToAssociatedWorkspaceResponseArrayOutput() AssociatedWorkspaceResponseArrayOutput {
	return i.ToAssociatedWorkspaceResponseArrayOutputWithContext(context.Background())
}

func (i AssociatedWorkspaceResponseArray) ToAssociatedWorkspaceResponseArrayOutputWithContext(ctx context.Context) AssociatedWorkspaceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociatedWorkspaceResponseArrayOutput)
}

type AssociatedWorkspaceResponseOutput struct{ *pulumi.OutputState }

func (AssociatedWorkspaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociatedWorkspaceResponse)(nil)).Elem()
}

func (o AssociatedWorkspaceResponseOutput) ToAssociatedWorkspaceResponseOutput() AssociatedWorkspaceResponseOutput {
	return o
}

func (o AssociatedWorkspaceResponseOutput) ToAssociatedWorkspaceResponseOutputWithContext(ctx context.Context) AssociatedWorkspaceResponseOutput {
	return o
}

func (o AssociatedWorkspaceResponseOutput) AssociateDate() pulumi.StringOutput {
	return o.ApplyT(func(v AssociatedWorkspaceResponse) string { return v.AssociateDate }).(pulumi.StringOutput)
}

func (o AssociatedWorkspaceResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AssociatedWorkspaceResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o AssociatedWorkspaceResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v AssociatedWorkspaceResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o AssociatedWorkspaceResponseOutput) WorkspaceName() pulumi.StringOutput {
	return o.ApplyT(func(v AssociatedWorkspaceResponse) string { return v.WorkspaceName }).(pulumi.StringOutput)
}

type AssociatedWorkspaceResponseArrayOutput struct{ *pulumi.OutputState }

func (AssociatedWorkspaceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssociatedWorkspaceResponse)(nil)).Elem()
}

func (o AssociatedWorkspaceResponseArrayOutput) ToAssociatedWorkspaceResponseArrayOutput() AssociatedWorkspaceResponseArrayOutput {
	return o
}

func (o AssociatedWorkspaceResponseArrayOutput) ToAssociatedWorkspaceResponseArrayOutputWithContext(ctx context.Context) AssociatedWorkspaceResponseArrayOutput {
	return o
}

func (o AssociatedWorkspaceResponseArrayOutput) Index(i pulumi.IntInput) AssociatedWorkspaceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AssociatedWorkspaceResponse {
		return vs[0].([]AssociatedWorkspaceResponse)[vs[1].(int)]
	}).(AssociatedWorkspaceResponseOutput)
}

type CapacityReservationPropertiesResponse struct {
	LastSkuUpdate string  `pulumi:"lastSkuUpdate"`
	MinCapacity   float64 `pulumi:"minCapacity"`
}





type CapacityReservationPropertiesResponseInput interface {
	pulumi.Input

	ToCapacityReservationPropertiesResponseOutput() CapacityReservationPropertiesResponseOutput
	ToCapacityReservationPropertiesResponseOutputWithContext(context.Context) CapacityReservationPropertiesResponseOutput
}

type CapacityReservationPropertiesResponseArgs struct {
	LastSkuUpdate pulumi.StringInput  `pulumi:"lastSkuUpdate"`
	MinCapacity   pulumi.Float64Input `pulumi:"minCapacity"`
}

func (CapacityReservationPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservationPropertiesResponse)(nil)).Elem()
}

func (i CapacityReservationPropertiesResponseArgs) ToCapacityReservationPropertiesResponseOutput() CapacityReservationPropertiesResponseOutput {
	return i.ToCapacityReservationPropertiesResponseOutputWithContext(context.Background())
}

func (i CapacityReservationPropertiesResponseArgs) ToCapacityReservationPropertiesResponseOutputWithContext(ctx context.Context) CapacityReservationPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationPropertiesResponseOutput)
}

func (i CapacityReservationPropertiesResponseArgs) ToCapacityReservationPropertiesResponsePtrOutput() CapacityReservationPropertiesResponsePtrOutput {
	return i.ToCapacityReservationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CapacityReservationPropertiesResponseArgs) ToCapacityReservationPropertiesResponsePtrOutputWithContext(ctx context.Context) CapacityReservationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationPropertiesResponseOutput).ToCapacityReservationPropertiesResponsePtrOutputWithContext(ctx)
}









type CapacityReservationPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCapacityReservationPropertiesResponsePtrOutput() CapacityReservationPropertiesResponsePtrOutput
	ToCapacityReservationPropertiesResponsePtrOutputWithContext(context.Context) CapacityReservationPropertiesResponsePtrOutput
}

type capacityReservationPropertiesResponsePtrType CapacityReservationPropertiesResponseArgs

func CapacityReservationPropertiesResponsePtr(v *CapacityReservationPropertiesResponseArgs) CapacityReservationPropertiesResponsePtrInput {
	return (*capacityReservationPropertiesResponsePtrType)(v)
}

func (*capacityReservationPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationPropertiesResponse)(nil)).Elem()
}

func (i *capacityReservationPropertiesResponsePtrType) ToCapacityReservationPropertiesResponsePtrOutput() CapacityReservationPropertiesResponsePtrOutput {
	return i.ToCapacityReservationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *capacityReservationPropertiesResponsePtrType) ToCapacityReservationPropertiesResponsePtrOutputWithContext(ctx context.Context) CapacityReservationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationPropertiesResponsePtrOutput)
}

type CapacityReservationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CapacityReservationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservationPropertiesResponse)(nil)).Elem()
}

func (o CapacityReservationPropertiesResponseOutput) ToCapacityReservationPropertiesResponseOutput() CapacityReservationPropertiesResponseOutput {
	return o
}

func (o CapacityReservationPropertiesResponseOutput) ToCapacityReservationPropertiesResponseOutputWithContext(ctx context.Context) CapacityReservationPropertiesResponseOutput {
	return o
}

func (o CapacityReservationPropertiesResponseOutput) ToCapacityReservationPropertiesResponsePtrOutput() CapacityReservationPropertiesResponsePtrOutput {
	return o.ToCapacityReservationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CapacityReservationPropertiesResponseOutput) ToCapacityReservationPropertiesResponsePtrOutputWithContext(ctx context.Context) CapacityReservationPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CapacityReservationPropertiesResponse) *CapacityReservationPropertiesResponse {
		return &v
	}).(CapacityReservationPropertiesResponsePtrOutput)
}

func (o CapacityReservationPropertiesResponseOutput) LastSkuUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v CapacityReservationPropertiesResponse) string { return v.LastSkuUpdate }).(pulumi.StringOutput)
}

func (o CapacityReservationPropertiesResponseOutput) MinCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v CapacityReservationPropertiesResponse) float64 { return v.MinCapacity }).(pulumi.Float64Output)
}

type CapacityReservationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CapacityReservationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationPropertiesResponse)(nil)).Elem()
}

func (o CapacityReservationPropertiesResponsePtrOutput) ToCapacityReservationPropertiesResponsePtrOutput() CapacityReservationPropertiesResponsePtrOutput {
	return o
}

func (o CapacityReservationPropertiesResponsePtrOutput) ToCapacityReservationPropertiesResponsePtrOutputWithContext(ctx context.Context) CapacityReservationPropertiesResponsePtrOutput {
	return o
}

func (o CapacityReservationPropertiesResponsePtrOutput) Elem() CapacityReservationPropertiesResponseOutput {
	return o.ApplyT(func(v *CapacityReservationPropertiesResponse) CapacityReservationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CapacityReservationPropertiesResponse
		return ret
	}).(CapacityReservationPropertiesResponseOutput)
}

func (o CapacityReservationPropertiesResponsePtrOutput) LastSkuUpdate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapacityReservationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSkuUpdate
	}).(pulumi.StringPtrOutput)
}

func (o CapacityReservationPropertiesResponsePtrOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CapacityReservationPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MinCapacity
	}).(pulumi.Float64PtrOutput)
}

type ClusterSku struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
}





type ClusterSkuInput interface {
	pulumi.Input

	ToClusterSkuOutput() ClusterSkuOutput
	ToClusterSkuOutputWithContext(context.Context) ClusterSkuOutput
}

type ClusterSkuArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
}

func (ClusterSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (i ClusterSkuArgs) ToClusterSkuOutput() ClusterSkuOutput {
	return i.ToClusterSkuOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput)
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput).ToClusterSkuPtrOutputWithContext(ctx)
}









type ClusterSkuPtrInput interface {
	pulumi.Input

	ToClusterSkuPtrOutput() ClusterSkuPtrOutput
	ToClusterSkuPtrOutputWithContext(context.Context) ClusterSkuPtrOutput
}

type clusterSkuPtrType ClusterSkuArgs

func ClusterSkuPtr(v *ClusterSkuArgs) ClusterSkuPtrInput {
	return (*clusterSkuPtrType)(v)
}

func (*clusterSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuPtrOutput)
}

type ClusterSkuOutput struct{ *pulumi.OutputState }

func (ClusterSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (o ClusterSkuOutput) ToClusterSkuOutput() ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSku) *ClusterSku {
		return &v
	}).(ClusterSkuPtrOutput)
}

func (o ClusterSkuOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ClusterSku) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o ClusterSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ClusterSkuPtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) Elem() ClusterSkuOutput {
	return o.ApplyT(func(v *ClusterSku) ClusterSku {
		if v != nil {
			return *v
		}
		var ret ClusterSku
		return ret
	}).(ClusterSkuOutput)
}

func (o ClusterSkuPtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ClusterSku) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o ClusterSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ClusterSkuResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
}





type ClusterSkuResponseInput interface {
	pulumi.Input

	ToClusterSkuResponseOutput() ClusterSkuResponseOutput
	ToClusterSkuResponseOutputWithContext(context.Context) ClusterSkuResponseOutput
}

type ClusterSkuResponseArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
}

func (ClusterSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuResponse)(nil)).Elem()
}

func (i ClusterSkuResponseArgs) ToClusterSkuResponseOutput() ClusterSkuResponseOutput {
	return i.ToClusterSkuResponseOutputWithContext(context.Background())
}

func (i ClusterSkuResponseArgs) ToClusterSkuResponseOutputWithContext(ctx context.Context) ClusterSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuResponseOutput)
}

func (i ClusterSkuResponseArgs) ToClusterSkuResponsePtrOutput() ClusterSkuResponsePtrOutput {
	return i.ToClusterSkuResponsePtrOutputWithContext(context.Background())
}

func (i ClusterSkuResponseArgs) ToClusterSkuResponsePtrOutputWithContext(ctx context.Context) ClusterSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuResponseOutput).ToClusterSkuResponsePtrOutputWithContext(ctx)
}









type ClusterSkuResponsePtrInput interface {
	pulumi.Input

	ToClusterSkuResponsePtrOutput() ClusterSkuResponsePtrOutput
	ToClusterSkuResponsePtrOutputWithContext(context.Context) ClusterSkuResponsePtrOutput
}

type clusterSkuResponsePtrType ClusterSkuResponseArgs

func ClusterSkuResponsePtr(v *ClusterSkuResponseArgs) ClusterSkuResponsePtrInput {
	return (*clusterSkuResponsePtrType)(v)
}

func (*clusterSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSkuResponse)(nil)).Elem()
}

func (i *clusterSkuResponsePtrType) ToClusterSkuResponsePtrOutput() ClusterSkuResponsePtrOutput {
	return i.ToClusterSkuResponsePtrOutputWithContext(context.Background())
}

func (i *clusterSkuResponsePtrType) ToClusterSkuResponsePtrOutputWithContext(ctx context.Context) ClusterSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuResponsePtrOutput)
}

type ClusterSkuResponseOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutput() ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutputWithContext(ctx context.Context) ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponsePtrOutput() ClusterSkuResponsePtrOutput {
	return o.ToClusterSkuResponsePtrOutputWithContext(context.Background())
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponsePtrOutputWithContext(ctx context.Context) ClusterSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSkuResponse) *ClusterSkuResponse {
		return &v
	}).(ClusterSkuResponsePtrOutput)
}

func (o ClusterSkuResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ClusterSkuResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o ClusterSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ClusterSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutput() ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutputWithContext(ctx context.Context) ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) Elem() ClusterSkuResponseOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) ClusterSkuResponse {
		if v != nil {
			return *v
		}
		var ret ClusterSkuResponse
		return ret
	}).(ClusterSkuResponseOutput)
}

func (o ClusterSkuResponsePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o ClusterSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type Identity struct {
	Type                   IdentityType           `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   IdentityTypeInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput   `pulumi:"userAssignedIdentities"`
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

func (o IdentityOutput) Type() IdentityTypeOutput {
	return o.ApplyT(func(v Identity) IdentityType { return v.Type }).(IdentityTypeOutput)
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

func (o IdentityPtrOutput) Type() IdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *IdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(IdentityTypePtrOutput)
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
	PrincipalId            string                                    `pulumi:"principalId"`
	TenantId               string                                    `pulumi:"tenantId"`
	Type                   string                                    `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityPropertiesResponse `pulumi:"userAssignedIdentities"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                     `pulumi:"principalId"`
	TenantId               pulumi.StringInput                     `pulumi:"tenantId"`
	Type                   pulumi.StringInput                     `pulumi:"type"`
	UserAssignedIdentities UserIdentityPropertiesResponseMapInput `pulumi:"userAssignedIdentities"`
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

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]UserIdentityPropertiesResponse { return v.UserAssignedIdentities }).(UserIdentityPropertiesResponseMapOutput)
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
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]UserIdentityPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type KeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyRsaSize  *int    `pulumi:"keyRsaSize"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyRsaSize  pulumi.IntPtrInput    `pulumi:"keyRsaSize"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
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

func (o KeyVaultPropertiesOutput) KeyRsaSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *int { return v.KeyRsaSize }).(pulumi.IntPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
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

func (o KeyVaultPropertiesPtrOutput) KeyRsaSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *int {
		if v == nil {
			return nil
		}
		return v.KeyRsaSize
	}).(pulumi.IntPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
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

type KeyVaultPropertiesResponse struct {
	KeyName     *string `pulumi:"keyName"`
	KeyRsaSize  *int    `pulumi:"keyRsaSize"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyRsaSize  pulumi.IntPtrInput    `pulumi:"keyRsaSize"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
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

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyRsaSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *int { return v.KeyRsaSize }).(pulumi.IntPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
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

func (o KeyVaultPropertiesResponsePtrOutput) KeyRsaSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.KeyRsaSize
	}).(pulumi.IntPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
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

type PrivateLinkScopedResourceResponse struct {
	ResourceId *string `pulumi:"resourceId"`
	ScopeId    *string `pulumi:"scopeId"`
}





type PrivateLinkScopedResourceResponseInput interface {
	pulumi.Input

	ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput
	ToPrivateLinkScopedResourceResponseOutputWithContext(context.Context) PrivateLinkScopedResourceResponseOutput
}

type PrivateLinkScopedResourceResponseArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	ScopeId    pulumi.StringPtrInput `pulumi:"scopeId"`
}

func (PrivateLinkScopedResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (i PrivateLinkScopedResourceResponseArgs) ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput {
	return i.ToPrivateLinkScopedResourceResponseOutputWithContext(context.Background())
}

func (i PrivateLinkScopedResourceResponseArgs) ToPrivateLinkScopedResourceResponseOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopedResourceResponseOutput)
}





type PrivateLinkScopedResourceResponseArrayInput interface {
	pulumi.Input

	ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput
	ToPrivateLinkScopedResourceResponseArrayOutputWithContext(context.Context) PrivateLinkScopedResourceResponseArrayOutput
}

type PrivateLinkScopedResourceResponseArray []PrivateLinkScopedResourceResponseInput

func (PrivateLinkScopedResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (i PrivateLinkScopedResourceResponseArray) ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput {
	return i.ToPrivateLinkScopedResourceResponseArrayOutputWithContext(context.Background())
}

func (i PrivateLinkScopedResourceResponseArray) ToPrivateLinkScopedResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopedResourceResponseArrayOutput)
}

type PrivateLinkScopedResourceResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkScopedResourceResponseOutput) ScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ScopeId }).(pulumi.StringPtrOutput)
}

type PrivateLinkScopedResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkScopedResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkScopedResourceResponse {
		return vs[0].([]PrivateLinkScopedResourceResponse)[vs[1].(int)]
	}).(PrivateLinkScopedResourceResponseOutput)
}

type UserIdentityPropertiesResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserIdentityPropertiesResponseInput interface {
	pulumi.Input

	ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput
	ToUserIdentityPropertiesResponseOutputWithContext(context.Context) UserIdentityPropertiesResponseOutput
}

type UserIdentityPropertiesResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (UserIdentityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (i UserIdentityPropertiesResponseArgs) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return i.ToUserIdentityPropertiesResponseOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesResponseArgs) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesResponseOutput)
}





type UserIdentityPropertiesResponseMapInput interface {
	pulumi.Input

	ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput
	ToUserIdentityPropertiesResponseMapOutputWithContext(context.Context) UserIdentityPropertiesResponseMapOutput
}

type UserIdentityPropertiesResponseMap map[string]UserIdentityPropertiesResponseInput

func (UserIdentityPropertiesResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (i UserIdentityPropertiesResponseMap) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return i.ToUserIdentityPropertiesResponseMapOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesResponseMap) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesResponseMapOutput)
}

type UserIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserIdentityPropertiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserIdentityPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityPropertiesResponse {
		return vs[0].(map[string]UserIdentityPropertiesResponse)[vs[1].(string)]
	}).(UserIdentityPropertiesResponseOutput)
}

type WorkspaceCapping struct {
	DailyQuotaGb *float64 `pulumi:"dailyQuotaGb"`
}





type WorkspaceCappingInput interface {
	pulumi.Input

	ToWorkspaceCappingOutput() WorkspaceCappingOutput
	ToWorkspaceCappingOutputWithContext(context.Context) WorkspaceCappingOutput
}

type WorkspaceCappingArgs struct {
	DailyQuotaGb pulumi.Float64PtrInput `pulumi:"dailyQuotaGb"`
}

func (WorkspaceCappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCapping)(nil)).Elem()
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingOutput() WorkspaceCappingOutput {
	return i.ToWorkspaceCappingOutputWithContext(context.Background())
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingOutputWithContext(ctx context.Context) WorkspaceCappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingOutput)
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return i.ToWorkspaceCappingPtrOutputWithContext(context.Background())
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingOutput).ToWorkspaceCappingPtrOutputWithContext(ctx)
}









type WorkspaceCappingPtrInput interface {
	pulumi.Input

	ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput
	ToWorkspaceCappingPtrOutputWithContext(context.Context) WorkspaceCappingPtrOutput
}

type workspaceCappingPtrType WorkspaceCappingArgs

func WorkspaceCappingPtr(v *WorkspaceCappingArgs) WorkspaceCappingPtrInput {
	return (*workspaceCappingPtrType)(v)
}

func (*workspaceCappingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCapping)(nil)).Elem()
}

func (i *workspaceCappingPtrType) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return i.ToWorkspaceCappingPtrOutputWithContext(context.Background())
}

func (i *workspaceCappingPtrType) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingPtrOutput)
}

type WorkspaceCappingOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCapping)(nil)).Elem()
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingOutput() WorkspaceCappingOutput {
	return o
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingOutputWithContext(ctx context.Context) WorkspaceCappingOutput {
	return o
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return o.ToWorkspaceCappingPtrOutputWithContext(context.Background())
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCapping) *WorkspaceCapping {
		return &v
	}).(WorkspaceCappingPtrOutput)
}

func (o WorkspaceCappingOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkspaceCapping) *float64 { return v.DailyQuotaGb }).(pulumi.Float64PtrOutput)
}

type WorkspaceCappingPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCapping)(nil)).Elem()
}

func (o WorkspaceCappingPtrOutput) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return o
}

func (o WorkspaceCappingPtrOutput) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return o
}

func (o WorkspaceCappingPtrOutput) Elem() WorkspaceCappingOutput {
	return o.ApplyT(func(v *WorkspaceCapping) WorkspaceCapping {
		if v != nil {
			return *v
		}
		var ret WorkspaceCapping
		return ret
	}).(WorkspaceCappingOutput)
}

func (o WorkspaceCappingPtrOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkspaceCapping) *float64 {
		if v == nil {
			return nil
		}
		return v.DailyQuotaGb
	}).(pulumi.Float64PtrOutput)
}

type WorkspaceCappingResponse struct {
	DailyQuotaGb        *float64 `pulumi:"dailyQuotaGb"`
	DataIngestionStatus string   `pulumi:"dataIngestionStatus"`
	QuotaNextResetTime  string   `pulumi:"quotaNextResetTime"`
}





type WorkspaceCappingResponseInput interface {
	pulumi.Input

	ToWorkspaceCappingResponseOutput() WorkspaceCappingResponseOutput
	ToWorkspaceCappingResponseOutputWithContext(context.Context) WorkspaceCappingResponseOutput
}

type WorkspaceCappingResponseArgs struct {
	DailyQuotaGb        pulumi.Float64PtrInput `pulumi:"dailyQuotaGb"`
	DataIngestionStatus pulumi.StringInput     `pulumi:"dataIngestionStatus"`
	QuotaNextResetTime  pulumi.StringInput     `pulumi:"quotaNextResetTime"`
}

func (WorkspaceCappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCappingResponse)(nil)).Elem()
}

func (i WorkspaceCappingResponseArgs) ToWorkspaceCappingResponseOutput() WorkspaceCappingResponseOutput {
	return i.ToWorkspaceCappingResponseOutputWithContext(context.Background())
}

func (i WorkspaceCappingResponseArgs) ToWorkspaceCappingResponseOutputWithContext(ctx context.Context) WorkspaceCappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingResponseOutput)
}

func (i WorkspaceCappingResponseArgs) ToWorkspaceCappingResponsePtrOutput() WorkspaceCappingResponsePtrOutput {
	return i.ToWorkspaceCappingResponsePtrOutputWithContext(context.Background())
}

func (i WorkspaceCappingResponseArgs) ToWorkspaceCappingResponsePtrOutputWithContext(ctx context.Context) WorkspaceCappingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingResponseOutput).ToWorkspaceCappingResponsePtrOutputWithContext(ctx)
}









type WorkspaceCappingResponsePtrInput interface {
	pulumi.Input

	ToWorkspaceCappingResponsePtrOutput() WorkspaceCappingResponsePtrOutput
	ToWorkspaceCappingResponsePtrOutputWithContext(context.Context) WorkspaceCappingResponsePtrOutput
}

type workspaceCappingResponsePtrType WorkspaceCappingResponseArgs

func WorkspaceCappingResponsePtr(v *WorkspaceCappingResponseArgs) WorkspaceCappingResponsePtrInput {
	return (*workspaceCappingResponsePtrType)(v)
}

func (*workspaceCappingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCappingResponse)(nil)).Elem()
}

func (i *workspaceCappingResponsePtrType) ToWorkspaceCappingResponsePtrOutput() WorkspaceCappingResponsePtrOutput {
	return i.ToWorkspaceCappingResponsePtrOutputWithContext(context.Background())
}

func (i *workspaceCappingResponsePtrType) ToWorkspaceCappingResponsePtrOutputWithContext(ctx context.Context) WorkspaceCappingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingResponsePtrOutput)
}

type WorkspaceCappingResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCappingResponse)(nil)).Elem()
}

func (o WorkspaceCappingResponseOutput) ToWorkspaceCappingResponseOutput() WorkspaceCappingResponseOutput {
	return o
}

func (o WorkspaceCappingResponseOutput) ToWorkspaceCappingResponseOutputWithContext(ctx context.Context) WorkspaceCappingResponseOutput {
	return o
}

func (o WorkspaceCappingResponseOutput) ToWorkspaceCappingResponsePtrOutput() WorkspaceCappingResponsePtrOutput {
	return o.ToWorkspaceCappingResponsePtrOutputWithContext(context.Background())
}

func (o WorkspaceCappingResponseOutput) ToWorkspaceCappingResponsePtrOutputWithContext(ctx context.Context) WorkspaceCappingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCappingResponse) *WorkspaceCappingResponse {
		return &v
	}).(WorkspaceCappingResponsePtrOutput)
}

func (o WorkspaceCappingResponseOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkspaceCappingResponse) *float64 { return v.DailyQuotaGb }).(pulumi.Float64PtrOutput)
}

func (o WorkspaceCappingResponseOutput) DataIngestionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCappingResponse) string { return v.DataIngestionStatus }).(pulumi.StringOutput)
}

func (o WorkspaceCappingResponseOutput) QuotaNextResetTime() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCappingResponse) string { return v.QuotaNextResetTime }).(pulumi.StringOutput)
}

type WorkspaceCappingResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCappingResponse)(nil)).Elem()
}

func (o WorkspaceCappingResponsePtrOutput) ToWorkspaceCappingResponsePtrOutput() WorkspaceCappingResponsePtrOutput {
	return o
}

func (o WorkspaceCappingResponsePtrOutput) ToWorkspaceCappingResponsePtrOutputWithContext(ctx context.Context) WorkspaceCappingResponsePtrOutput {
	return o
}

func (o WorkspaceCappingResponsePtrOutput) Elem() WorkspaceCappingResponseOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) WorkspaceCappingResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceCappingResponse
		return ret
	}).(WorkspaceCappingResponseOutput)
}

func (o WorkspaceCappingResponsePtrOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.DailyQuotaGb
	}).(pulumi.Float64PtrOutput)
}

func (o WorkspaceCappingResponsePtrOutput) DataIngestionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataIngestionStatus
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceCappingResponsePtrOutput) QuotaNextResetTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.QuotaNextResetTime
	}).(pulumi.StringPtrOutput)
}

type WorkspaceFeatures struct {
	ClusterResourceId                           *string `pulumi:"clusterResourceId"`
	DisableLocalAuth                            *bool   `pulumi:"disableLocalAuth"`
	EnableDataExport                            *bool   `pulumi:"enableDataExport"`
	EnableLogAccessUsingOnlyResourcePermissions *bool   `pulumi:"enableLogAccessUsingOnlyResourcePermissions"`
	ImmediatePurgeDataOn30Days                  *bool   `pulumi:"immediatePurgeDataOn30Days"`
}





type WorkspaceFeaturesInput interface {
	pulumi.Input

	ToWorkspaceFeaturesOutput() WorkspaceFeaturesOutput
	ToWorkspaceFeaturesOutputWithContext(context.Context) WorkspaceFeaturesOutput
}

type WorkspaceFeaturesArgs struct {
	ClusterResourceId                           pulumi.StringPtrInput `pulumi:"clusterResourceId"`
	DisableLocalAuth                            pulumi.BoolPtrInput   `pulumi:"disableLocalAuth"`
	EnableDataExport                            pulumi.BoolPtrInput   `pulumi:"enableDataExport"`
	EnableLogAccessUsingOnlyResourcePermissions pulumi.BoolPtrInput   `pulumi:"enableLogAccessUsingOnlyResourcePermissions"`
	ImmediatePurgeDataOn30Days                  pulumi.BoolPtrInput   `pulumi:"immediatePurgeDataOn30Days"`
}

func (WorkspaceFeaturesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceFeatures)(nil)).Elem()
}

func (i WorkspaceFeaturesArgs) ToWorkspaceFeaturesOutput() WorkspaceFeaturesOutput {
	return i.ToWorkspaceFeaturesOutputWithContext(context.Background())
}

func (i WorkspaceFeaturesArgs) ToWorkspaceFeaturesOutputWithContext(ctx context.Context) WorkspaceFeaturesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceFeaturesOutput)
}

func (i WorkspaceFeaturesArgs) ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput {
	return i.ToWorkspaceFeaturesPtrOutputWithContext(context.Background())
}

func (i WorkspaceFeaturesArgs) ToWorkspaceFeaturesPtrOutputWithContext(ctx context.Context) WorkspaceFeaturesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceFeaturesOutput).ToWorkspaceFeaturesPtrOutputWithContext(ctx)
}









type WorkspaceFeaturesPtrInput interface {
	pulumi.Input

	ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput
	ToWorkspaceFeaturesPtrOutputWithContext(context.Context) WorkspaceFeaturesPtrOutput
}

type workspaceFeaturesPtrType WorkspaceFeaturesArgs

func WorkspaceFeaturesPtr(v *WorkspaceFeaturesArgs) WorkspaceFeaturesPtrInput {
	return (*workspaceFeaturesPtrType)(v)
}

func (*workspaceFeaturesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceFeatures)(nil)).Elem()
}

func (i *workspaceFeaturesPtrType) ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput {
	return i.ToWorkspaceFeaturesPtrOutputWithContext(context.Background())
}

func (i *workspaceFeaturesPtrType) ToWorkspaceFeaturesPtrOutputWithContext(ctx context.Context) WorkspaceFeaturesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceFeaturesPtrOutput)
}

type WorkspaceFeaturesOutput struct{ *pulumi.OutputState }

func (WorkspaceFeaturesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceFeatures)(nil)).Elem()
}

func (o WorkspaceFeaturesOutput) ToWorkspaceFeaturesOutput() WorkspaceFeaturesOutput {
	return o
}

func (o WorkspaceFeaturesOutput) ToWorkspaceFeaturesOutputWithContext(ctx context.Context) WorkspaceFeaturesOutput {
	return o
}

func (o WorkspaceFeaturesOutput) ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput {
	return o.ToWorkspaceFeaturesPtrOutputWithContext(context.Background())
}

func (o WorkspaceFeaturesOutput) ToWorkspaceFeaturesPtrOutputWithContext(ctx context.Context) WorkspaceFeaturesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceFeatures) *WorkspaceFeatures {
		return &v
	}).(WorkspaceFeaturesPtrOutput)
}

func (o WorkspaceFeaturesOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *string { return v.ClusterResourceId }).(pulumi.StringPtrOutput)
}

func (o WorkspaceFeaturesOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesOutput) EnableDataExport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *bool { return v.EnableDataExport }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesOutput) EnableLogAccessUsingOnlyResourcePermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *bool { return v.EnableLogAccessUsingOnlyResourcePermissions }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *bool { return v.ImmediatePurgeDataOn30Days }).(pulumi.BoolPtrOutput)
}

type WorkspaceFeaturesPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceFeaturesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceFeatures)(nil)).Elem()
}

func (o WorkspaceFeaturesPtrOutput) ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput {
	return o
}

func (o WorkspaceFeaturesPtrOutput) ToWorkspaceFeaturesPtrOutputWithContext(ctx context.Context) WorkspaceFeaturesPtrOutput {
	return o
}

func (o WorkspaceFeaturesPtrOutput) Elem() WorkspaceFeaturesOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) WorkspaceFeatures {
		if v != nil {
			return *v
		}
		var ret WorkspaceFeatures
		return ret
	}).(WorkspaceFeaturesOutput)
}

func (o WorkspaceFeaturesPtrOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *string {
		if v == nil {
			return nil
		}
		return v.ClusterResourceId
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceFeaturesPtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesPtrOutput) EnableDataExport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDataExport
	}).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesPtrOutput) EnableLogAccessUsingOnlyResourcePermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *bool {
		if v == nil {
			return nil
		}
		return v.EnableLogAccessUsingOnlyResourcePermissions
	}).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesPtrOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *bool {
		if v == nil {
			return nil
		}
		return v.ImmediatePurgeDataOn30Days
	}).(pulumi.BoolPtrOutput)
}

type WorkspaceFeaturesResponse struct {
	ClusterResourceId                           *string `pulumi:"clusterResourceId"`
	DisableLocalAuth                            *bool   `pulumi:"disableLocalAuth"`
	EnableDataExport                            *bool   `pulumi:"enableDataExport"`
	EnableLogAccessUsingOnlyResourcePermissions *bool   `pulumi:"enableLogAccessUsingOnlyResourcePermissions"`
	ImmediatePurgeDataOn30Days                  *bool   `pulumi:"immediatePurgeDataOn30Days"`
}





type WorkspaceFeaturesResponseInput interface {
	pulumi.Input

	ToWorkspaceFeaturesResponseOutput() WorkspaceFeaturesResponseOutput
	ToWorkspaceFeaturesResponseOutputWithContext(context.Context) WorkspaceFeaturesResponseOutput
}

type WorkspaceFeaturesResponseArgs struct {
	ClusterResourceId                           pulumi.StringPtrInput `pulumi:"clusterResourceId"`
	DisableLocalAuth                            pulumi.BoolPtrInput   `pulumi:"disableLocalAuth"`
	EnableDataExport                            pulumi.BoolPtrInput   `pulumi:"enableDataExport"`
	EnableLogAccessUsingOnlyResourcePermissions pulumi.BoolPtrInput   `pulumi:"enableLogAccessUsingOnlyResourcePermissions"`
	ImmediatePurgeDataOn30Days                  pulumi.BoolPtrInput   `pulumi:"immediatePurgeDataOn30Days"`
}

func (WorkspaceFeaturesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceFeaturesResponse)(nil)).Elem()
}

func (i WorkspaceFeaturesResponseArgs) ToWorkspaceFeaturesResponseOutput() WorkspaceFeaturesResponseOutput {
	return i.ToWorkspaceFeaturesResponseOutputWithContext(context.Background())
}

func (i WorkspaceFeaturesResponseArgs) ToWorkspaceFeaturesResponseOutputWithContext(ctx context.Context) WorkspaceFeaturesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceFeaturesResponseOutput)
}

func (i WorkspaceFeaturesResponseArgs) ToWorkspaceFeaturesResponsePtrOutput() WorkspaceFeaturesResponsePtrOutput {
	return i.ToWorkspaceFeaturesResponsePtrOutputWithContext(context.Background())
}

func (i WorkspaceFeaturesResponseArgs) ToWorkspaceFeaturesResponsePtrOutputWithContext(ctx context.Context) WorkspaceFeaturesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceFeaturesResponseOutput).ToWorkspaceFeaturesResponsePtrOutputWithContext(ctx)
}









type WorkspaceFeaturesResponsePtrInput interface {
	pulumi.Input

	ToWorkspaceFeaturesResponsePtrOutput() WorkspaceFeaturesResponsePtrOutput
	ToWorkspaceFeaturesResponsePtrOutputWithContext(context.Context) WorkspaceFeaturesResponsePtrOutput
}

type workspaceFeaturesResponsePtrType WorkspaceFeaturesResponseArgs

func WorkspaceFeaturesResponsePtr(v *WorkspaceFeaturesResponseArgs) WorkspaceFeaturesResponsePtrInput {
	return (*workspaceFeaturesResponsePtrType)(v)
}

func (*workspaceFeaturesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceFeaturesResponse)(nil)).Elem()
}

func (i *workspaceFeaturesResponsePtrType) ToWorkspaceFeaturesResponsePtrOutput() WorkspaceFeaturesResponsePtrOutput {
	return i.ToWorkspaceFeaturesResponsePtrOutputWithContext(context.Background())
}

func (i *workspaceFeaturesResponsePtrType) ToWorkspaceFeaturesResponsePtrOutputWithContext(ctx context.Context) WorkspaceFeaturesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceFeaturesResponsePtrOutput)
}

type WorkspaceFeaturesResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceFeaturesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceFeaturesResponse)(nil)).Elem()
}

func (o WorkspaceFeaturesResponseOutput) ToWorkspaceFeaturesResponseOutput() WorkspaceFeaturesResponseOutput {
	return o
}

func (o WorkspaceFeaturesResponseOutput) ToWorkspaceFeaturesResponseOutputWithContext(ctx context.Context) WorkspaceFeaturesResponseOutput {
	return o
}

func (o WorkspaceFeaturesResponseOutput) ToWorkspaceFeaturesResponsePtrOutput() WorkspaceFeaturesResponsePtrOutput {
	return o.ToWorkspaceFeaturesResponsePtrOutputWithContext(context.Background())
}

func (o WorkspaceFeaturesResponseOutput) ToWorkspaceFeaturesResponsePtrOutputWithContext(ctx context.Context) WorkspaceFeaturesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceFeaturesResponse) *WorkspaceFeaturesResponse {
		return &v
	}).(WorkspaceFeaturesResponsePtrOutput)
}

func (o WorkspaceFeaturesResponseOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *string { return v.ClusterResourceId }).(pulumi.StringPtrOutput)
}

func (o WorkspaceFeaturesResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesResponseOutput) EnableDataExport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *bool { return v.EnableDataExport }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesResponseOutput) EnableLogAccessUsingOnlyResourcePermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *bool { return v.EnableLogAccessUsingOnlyResourcePermissions }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesResponseOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *bool { return v.ImmediatePurgeDataOn30Days }).(pulumi.BoolPtrOutput)
}

type WorkspaceFeaturesResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceFeaturesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceFeaturesResponse)(nil)).Elem()
}

func (o WorkspaceFeaturesResponsePtrOutput) ToWorkspaceFeaturesResponsePtrOutput() WorkspaceFeaturesResponsePtrOutput {
	return o
}

func (o WorkspaceFeaturesResponsePtrOutput) ToWorkspaceFeaturesResponsePtrOutputWithContext(ctx context.Context) WorkspaceFeaturesResponsePtrOutput {
	return o
}

func (o WorkspaceFeaturesResponsePtrOutput) Elem() WorkspaceFeaturesResponseOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) WorkspaceFeaturesResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceFeaturesResponse
		return ret
	}).(WorkspaceFeaturesResponseOutput)
}

func (o WorkspaceFeaturesResponsePtrOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterResourceId
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceFeaturesResponsePtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesResponsePtrOutput) EnableDataExport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDataExport
	}).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesResponsePtrOutput) EnableLogAccessUsingOnlyResourcePermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableLogAccessUsingOnlyResourcePermissions
	}).(pulumi.BoolPtrOutput)
}

func (o WorkspaceFeaturesResponsePtrOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ImmediatePurgeDataOn30Days
	}).(pulumi.BoolPtrOutput)
}

type WorkspaceSku struct {
	CapacityReservationLevel *int   `pulumi:"capacityReservationLevel"`
	Name                     string `pulumi:"name"`
}





type WorkspaceSkuInput interface {
	pulumi.Input

	ToWorkspaceSkuOutput() WorkspaceSkuOutput
	ToWorkspaceSkuOutputWithContext(context.Context) WorkspaceSkuOutput
}

type WorkspaceSkuArgs struct {
	CapacityReservationLevel pulumi.IntPtrInput `pulumi:"capacityReservationLevel"`
	Name                     pulumi.StringInput `pulumi:"name"`
}

func (WorkspaceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSku)(nil)).Elem()
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuOutput() WorkspaceSkuOutput {
	return i.ToWorkspaceSkuOutputWithContext(context.Background())
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuOutputWithContext(ctx context.Context) WorkspaceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuOutput)
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return i.ToWorkspaceSkuPtrOutputWithContext(context.Background())
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuOutput).ToWorkspaceSkuPtrOutputWithContext(ctx)
}









type WorkspaceSkuPtrInput interface {
	pulumi.Input

	ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput
	ToWorkspaceSkuPtrOutputWithContext(context.Context) WorkspaceSkuPtrOutput
}

type workspaceSkuPtrType WorkspaceSkuArgs

func WorkspaceSkuPtr(v *WorkspaceSkuArgs) WorkspaceSkuPtrInput {
	return (*workspaceSkuPtrType)(v)
}

func (*workspaceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSku)(nil)).Elem()
}

func (i *workspaceSkuPtrType) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return i.ToWorkspaceSkuPtrOutputWithContext(context.Background())
}

func (i *workspaceSkuPtrType) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuPtrOutput)
}

type WorkspaceSkuOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSku)(nil)).Elem()
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuOutput() WorkspaceSkuOutput {
	return o
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuOutputWithContext(ctx context.Context) WorkspaceSkuOutput {
	return o
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return o.ToWorkspaceSkuPtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceSku) *WorkspaceSku {
		return &v
	}).(WorkspaceSkuPtrOutput)
}

func (o WorkspaceSkuOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkspaceSku) *int { return v.CapacityReservationLevel }).(pulumi.IntPtrOutput)
}

func (o WorkspaceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceSku) string { return v.Name }).(pulumi.StringOutput)
}

type WorkspaceSkuPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSku)(nil)).Elem()
}

func (o WorkspaceSkuPtrOutput) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return o
}

func (o WorkspaceSkuPtrOutput) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return o
}

func (o WorkspaceSkuPtrOutput) Elem() WorkspaceSkuOutput {
	return o.ApplyT(func(v *WorkspaceSku) WorkspaceSku {
		if v != nil {
			return *v
		}
		var ret WorkspaceSku
		return ret
	}).(WorkspaceSkuOutput)
}

func (o WorkspaceSkuPtrOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkspaceSku) *int {
		if v == nil {
			return nil
		}
		return v.CapacityReservationLevel
	}).(pulumi.IntPtrOutput)
}

func (o WorkspaceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type WorkspaceSkuResponse struct {
	CapacityReservationLevel *int   `pulumi:"capacityReservationLevel"`
	LastSkuUpdate            string `pulumi:"lastSkuUpdate"`
	Name                     string `pulumi:"name"`
}





type WorkspaceSkuResponseInput interface {
	pulumi.Input

	ToWorkspaceSkuResponseOutput() WorkspaceSkuResponseOutput
	ToWorkspaceSkuResponseOutputWithContext(context.Context) WorkspaceSkuResponseOutput
}

type WorkspaceSkuResponseArgs struct {
	CapacityReservationLevel pulumi.IntPtrInput `pulumi:"capacityReservationLevel"`
	LastSkuUpdate            pulumi.StringInput `pulumi:"lastSkuUpdate"`
	Name                     pulumi.StringInput `pulumi:"name"`
}

func (WorkspaceSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSkuResponse)(nil)).Elem()
}

func (i WorkspaceSkuResponseArgs) ToWorkspaceSkuResponseOutput() WorkspaceSkuResponseOutput {
	return i.ToWorkspaceSkuResponseOutputWithContext(context.Background())
}

func (i WorkspaceSkuResponseArgs) ToWorkspaceSkuResponseOutputWithContext(ctx context.Context) WorkspaceSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuResponseOutput)
}

func (i WorkspaceSkuResponseArgs) ToWorkspaceSkuResponsePtrOutput() WorkspaceSkuResponsePtrOutput {
	return i.ToWorkspaceSkuResponsePtrOutputWithContext(context.Background())
}

func (i WorkspaceSkuResponseArgs) ToWorkspaceSkuResponsePtrOutputWithContext(ctx context.Context) WorkspaceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuResponseOutput).ToWorkspaceSkuResponsePtrOutputWithContext(ctx)
}









type WorkspaceSkuResponsePtrInput interface {
	pulumi.Input

	ToWorkspaceSkuResponsePtrOutput() WorkspaceSkuResponsePtrOutput
	ToWorkspaceSkuResponsePtrOutputWithContext(context.Context) WorkspaceSkuResponsePtrOutput
}

type workspaceSkuResponsePtrType WorkspaceSkuResponseArgs

func WorkspaceSkuResponsePtr(v *WorkspaceSkuResponseArgs) WorkspaceSkuResponsePtrInput {
	return (*workspaceSkuResponsePtrType)(v)
}

func (*workspaceSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSkuResponse)(nil)).Elem()
}

func (i *workspaceSkuResponsePtrType) ToWorkspaceSkuResponsePtrOutput() WorkspaceSkuResponsePtrOutput {
	return i.ToWorkspaceSkuResponsePtrOutputWithContext(context.Background())
}

func (i *workspaceSkuResponsePtrType) ToWorkspaceSkuResponsePtrOutputWithContext(ctx context.Context) WorkspaceSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuResponsePtrOutput)
}

type WorkspaceSkuResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSkuResponse)(nil)).Elem()
}

func (o WorkspaceSkuResponseOutput) ToWorkspaceSkuResponseOutput() WorkspaceSkuResponseOutput {
	return o
}

func (o WorkspaceSkuResponseOutput) ToWorkspaceSkuResponseOutputWithContext(ctx context.Context) WorkspaceSkuResponseOutput {
	return o
}

func (o WorkspaceSkuResponseOutput) ToWorkspaceSkuResponsePtrOutput() WorkspaceSkuResponsePtrOutput {
	return o.ToWorkspaceSkuResponsePtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuResponseOutput) ToWorkspaceSkuResponsePtrOutputWithContext(ctx context.Context) WorkspaceSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceSkuResponse) *WorkspaceSkuResponse {
		return &v
	}).(WorkspaceSkuResponsePtrOutput)
}

func (o WorkspaceSkuResponseOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) *int { return v.CapacityReservationLevel }).(pulumi.IntPtrOutput)
}

func (o WorkspaceSkuResponseOutput) LastSkuUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) string { return v.LastSkuUpdate }).(pulumi.StringOutput)
}

func (o WorkspaceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type WorkspaceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSkuResponse)(nil)).Elem()
}

func (o WorkspaceSkuResponsePtrOutput) ToWorkspaceSkuResponsePtrOutput() WorkspaceSkuResponsePtrOutput {
	return o
}

func (o WorkspaceSkuResponsePtrOutput) ToWorkspaceSkuResponsePtrOutputWithContext(ctx context.Context) WorkspaceSkuResponsePtrOutput {
	return o
}

func (o WorkspaceSkuResponsePtrOutput) Elem() WorkspaceSkuResponseOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) WorkspaceSkuResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceSkuResponse
		return ret
	}).(WorkspaceSkuResponseOutput)
}

func (o WorkspaceSkuResponsePtrOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.CapacityReservationLevel
	}).(pulumi.IntPtrOutput)
}

func (o WorkspaceSkuResponsePtrOutput) LastSkuUpdate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSkuUpdate
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssociatedWorkspaceResponseInput)(nil)).Elem(), AssociatedWorkspaceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssociatedWorkspaceResponseArrayInput)(nil)).Elem(), AssociatedWorkspaceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CapacityReservationPropertiesResponseInput)(nil)).Elem(), CapacityReservationPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CapacityReservationPropertiesResponsePtrInput)(nil)).Elem(), CapacityReservationPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSkuInput)(nil)).Elem(), ClusterSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSkuPtrInput)(nil)).Elem(), ClusterSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSkuResponseInput)(nil)).Elem(), ClusterSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSkuResponsePtrInput)(nil)).Elem(), ClusterSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPtrInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponsePtrInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesInput)(nil)).Elem(), KeyVaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesPtrInput)(nil)).Elem(), KeyVaultPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesResponseInput)(nil)).Elem(), KeyVaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultPropertiesResponsePtrInput)(nil)).Elem(), KeyVaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkScopedResourceResponseInput)(nil)).Elem(), PrivateLinkScopedResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkScopedResourceResponseArrayInput)(nil)).Elem(), PrivateLinkScopedResourceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityPropertiesResponseInput)(nil)).Elem(), UserIdentityPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityPropertiesResponseMapInput)(nil)).Elem(), UserIdentityPropertiesResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceCappingInput)(nil)).Elem(), WorkspaceCappingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceCappingPtrInput)(nil)).Elem(), WorkspaceCappingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceCappingResponseInput)(nil)).Elem(), WorkspaceCappingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceCappingResponsePtrInput)(nil)).Elem(), WorkspaceCappingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceFeaturesInput)(nil)).Elem(), WorkspaceFeaturesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceFeaturesPtrInput)(nil)).Elem(), WorkspaceFeaturesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceFeaturesResponseInput)(nil)).Elem(), WorkspaceFeaturesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceFeaturesResponsePtrInput)(nil)).Elem(), WorkspaceFeaturesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceSkuInput)(nil)).Elem(), WorkspaceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceSkuPtrInput)(nil)).Elem(), WorkspaceSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceSkuResponseInput)(nil)).Elem(), WorkspaceSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceSkuResponsePtrInput)(nil)).Elem(), WorkspaceSkuResponseArgs{})
	pulumi.RegisterOutputType(AssociatedWorkspaceResponseOutput{})
	pulumi.RegisterOutputType(AssociatedWorkspaceResponseArrayOutput{})
	pulumi.RegisterOutputType(CapacityReservationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CapacityReservationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuOutput{})
	pulumi.RegisterOutputType(ClusterSkuPtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponseOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceFeaturesOutput{})
	pulumi.RegisterOutputType(WorkspaceFeaturesPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceFeaturesResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceFeaturesResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuResponsePtrOutput{})
}
