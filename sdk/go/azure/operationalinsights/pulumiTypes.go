


package operationalinsights

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

type Column struct {
	DataTypeHint *string `pulumi:"dataTypeHint"`
	Description  *string `pulumi:"description"`
	DisplayName  *string `pulumi:"displayName"`
	Name         *string `pulumi:"name"`
	Type         *string `pulumi:"type"`
}





type ColumnInput interface {
	pulumi.Input

	ToColumnOutput() ColumnOutput
	ToColumnOutputWithContext(context.Context) ColumnOutput
}

type ColumnArgs struct {
	DataTypeHint pulumi.StringPtrInput `pulumi:"dataTypeHint"`
	Description  pulumi.StringPtrInput `pulumi:"description"`
	DisplayName  pulumi.StringPtrInput `pulumi:"displayName"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (ColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Column)(nil)).Elem()
}

func (i ColumnArgs) ToColumnOutput() ColumnOutput {
	return i.ToColumnOutputWithContext(context.Background())
}

func (i ColumnArgs) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnOutput)
}





type ColumnArrayInput interface {
	pulumi.Input

	ToColumnArrayOutput() ColumnArrayOutput
	ToColumnArrayOutputWithContext(context.Context) ColumnArrayOutput
}

type ColumnArray []ColumnInput

func (ColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Column)(nil)).Elem()
}

func (i ColumnArray) ToColumnArrayOutput() ColumnArrayOutput {
	return i.ToColumnArrayOutputWithContext(context.Background())
}

func (i ColumnArray) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnArrayOutput)
}

type ColumnOutput struct{ *pulumi.OutputState }

func (ColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Column)(nil)).Elem()
}

func (o ColumnOutput) ToColumnOutput() ColumnOutput {
	return o
}

func (o ColumnOutput) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return o
}

func (o ColumnOutput) DataTypeHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.DataTypeHint }).(pulumi.StringPtrOutput)
}

func (o ColumnOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ColumnOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ColumnOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnArrayOutput struct{ *pulumi.OutputState }

func (ColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Column)(nil)).Elem()
}

func (o ColumnArrayOutput) ToColumnArrayOutput() ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) Index(i pulumi.IntInput) ColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Column {
		return vs[0].([]Column)[vs[1].(int)]
	}).(ColumnOutput)
}

type ColumnResponse struct {
	DataTypeHint     *string `pulumi:"dataTypeHint"`
	Description      *string `pulumi:"description"`
	DisplayName      *string `pulumi:"displayName"`
	IsDefaultDisplay bool    `pulumi:"isDefaultDisplay"`
	IsHidden         bool    `pulumi:"isHidden"`
	Name             *string `pulumi:"name"`
	Type             *string `pulumi:"type"`
}

type ColumnResponseOutput struct{ *pulumi.OutputState }

func (ColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnResponse)(nil)).Elem()
}

func (o ColumnResponseOutput) ToColumnResponseOutput() ColumnResponseOutput {
	return o
}

func (o ColumnResponseOutput) ToColumnResponseOutputWithContext(ctx context.Context) ColumnResponseOutput {
	return o
}

func (o ColumnResponseOutput) DataTypeHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.DataTypeHint }).(pulumi.StringPtrOutput)
}

func (o ColumnResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ColumnResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ColumnResponseOutput) IsDefaultDisplay() pulumi.BoolOutput {
	return o.ApplyT(func(v ColumnResponse) bool { return v.IsDefaultDisplay }).(pulumi.BoolOutput)
}

func (o ColumnResponseOutput) IsHidden() pulumi.BoolOutput {
	return o.ApplyT(func(v ColumnResponse) bool { return v.IsHidden }).(pulumi.BoolOutput)
}

func (o ColumnResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (ColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ColumnResponse)(nil)).Elem()
}

func (o ColumnResponseArrayOutput) ToColumnResponseArrayOutput() ColumnResponseArrayOutput {
	return o
}

func (o ColumnResponseArrayOutput) ToColumnResponseArrayOutputWithContext(ctx context.Context) ColumnResponseArrayOutput {
	return o
}

func (o ColumnResponseArrayOutput) Index(i pulumi.IntInput) ColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ColumnResponse {
		return vs[0].([]ColumnResponse)[vs[1].(int)]
	}).(ColumnResponseOutput)
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

type LogAnalyticsQueryPackQueryPropertiesRelated struct {
	Categories    []string `pulumi:"categories"`
	ResourceTypes []string `pulumi:"resourceTypes"`
	Solutions     []string `pulumi:"solutions"`
}





type LogAnalyticsQueryPackQueryPropertiesRelatedInput interface {
	pulumi.Input

	ToLogAnalyticsQueryPackQueryPropertiesRelatedOutput() LogAnalyticsQueryPackQueryPropertiesRelatedOutput
	ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedOutput
}

type LogAnalyticsQueryPackQueryPropertiesRelatedArgs struct {
	Categories    pulumi.StringArrayInput `pulumi:"categories"`
	ResourceTypes pulumi.StringArrayInput `pulumi:"resourceTypes"`
	Solutions     pulumi.StringArrayInput `pulumi:"solutions"`
}

func (LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutput() LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return i.ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(context.Background())
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsQueryPackQueryPropertiesRelatedOutput)
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return i.ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Background())
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsQueryPackQueryPropertiesRelatedOutput).ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx)
}









type LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput interface {
	pulumi.Input

	ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput
	ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput
}

type logAnalyticsQueryPackQueryPropertiesRelatedPtrType LogAnalyticsQueryPackQueryPropertiesRelatedArgs

func LogAnalyticsQueryPackQueryPropertiesRelatedPtr(v *LogAnalyticsQueryPackQueryPropertiesRelatedArgs) LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput {
	return (*logAnalyticsQueryPackQueryPropertiesRelatedPtrType)(v)
}

func (*logAnalyticsQueryPackQueryPropertiesRelatedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (i *logAnalyticsQueryPackQueryPropertiesRelatedPtrType) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return i.ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsQueryPackQueryPropertiesRelatedPtrType) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput)
}

type LogAnalyticsQueryPackQueryPropertiesRelatedOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutput() LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o.ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsQueryPackQueryPropertiesRelated) *LogAnalyticsQueryPackQueryPropertiesRelated {
		return &v
	}).(LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesRelated) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesRelated) []string { return v.ResourceTypes }).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesRelated) []string { return v.Solutions }).(pulumi.StringArrayOutput)
}

type LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) Elem() LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) LogAnalyticsQueryPackQueryPropertiesRelated {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsQueryPackQueryPropertiesRelated
		return ret
	}).(LogAnalyticsQueryPackQueryPropertiesRelatedOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) []string {
		if v == nil {
			return nil
		}
		return v.ResourceTypes
	}).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) []string {
		if v == nil {
			return nil
		}
		return v.Solutions
	}).(pulumi.StringArrayOutput)
}

type LogAnalyticsQueryPackQueryPropertiesResponseRelated struct {
	Categories    []string `pulumi:"categories"`
	ResourceTypes []string `pulumi:"resourceTypes"`
	Solutions     []string `pulumi:"solutions"`
}

type LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsQueryPackQueryPropertiesResponseRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput() LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string { return v.ResourceTypes }).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string { return v.Solutions }).(pulumi.StringArrayOutput)
}

type LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsQueryPackQueryPropertiesResponseRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) Elem() LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) LogAnalyticsQueryPackQueryPropertiesResponseRelated {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsQueryPackQueryPropertiesResponseRelated
		return ret
	}).(LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string {
		if v == nil {
			return nil
		}
		return v.ResourceTypes
	}).(pulumi.StringArrayOutput)
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string {
		if v == nil {
			return nil
		}
		return v.Solutions
	}).(pulumi.StringArrayOutput)
}

type MachineReferenceWithHints struct {
	Id   string `pulumi:"id"`
	Kind string `pulumi:"kind"`
}





type MachineReferenceWithHintsInput interface {
	pulumi.Input

	ToMachineReferenceWithHintsOutput() MachineReferenceWithHintsOutput
	ToMachineReferenceWithHintsOutputWithContext(context.Context) MachineReferenceWithHintsOutput
}

type MachineReferenceWithHintsArgs struct {
	Id   pulumi.StringInput `pulumi:"id"`
	Kind pulumi.StringInput `pulumi:"kind"`
}

func (MachineReferenceWithHintsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineReferenceWithHints)(nil)).Elem()
}

func (i MachineReferenceWithHintsArgs) ToMachineReferenceWithHintsOutput() MachineReferenceWithHintsOutput {
	return i.ToMachineReferenceWithHintsOutputWithContext(context.Background())
}

func (i MachineReferenceWithHintsArgs) ToMachineReferenceWithHintsOutputWithContext(ctx context.Context) MachineReferenceWithHintsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineReferenceWithHintsOutput)
}





type MachineReferenceWithHintsArrayInput interface {
	pulumi.Input

	ToMachineReferenceWithHintsArrayOutput() MachineReferenceWithHintsArrayOutput
	ToMachineReferenceWithHintsArrayOutputWithContext(context.Context) MachineReferenceWithHintsArrayOutput
}

type MachineReferenceWithHintsArray []MachineReferenceWithHintsInput

func (MachineReferenceWithHintsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineReferenceWithHints)(nil)).Elem()
}

func (i MachineReferenceWithHintsArray) ToMachineReferenceWithHintsArrayOutput() MachineReferenceWithHintsArrayOutput {
	return i.ToMachineReferenceWithHintsArrayOutputWithContext(context.Background())
}

func (i MachineReferenceWithHintsArray) ToMachineReferenceWithHintsArrayOutputWithContext(ctx context.Context) MachineReferenceWithHintsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineReferenceWithHintsArrayOutput)
}

type MachineReferenceWithHintsOutput struct{ *pulumi.OutputState }

func (MachineReferenceWithHintsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineReferenceWithHints)(nil)).Elem()
}

func (o MachineReferenceWithHintsOutput) ToMachineReferenceWithHintsOutput() MachineReferenceWithHintsOutput {
	return o
}

func (o MachineReferenceWithHintsOutput) ToMachineReferenceWithHintsOutputWithContext(ctx context.Context) MachineReferenceWithHintsOutput {
	return o
}

func (o MachineReferenceWithHintsOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHints) string { return v.Id }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHints) string { return v.Kind }).(pulumi.StringOutput)
}

type MachineReferenceWithHintsArrayOutput struct{ *pulumi.OutputState }

func (MachineReferenceWithHintsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineReferenceWithHints)(nil)).Elem()
}

func (o MachineReferenceWithHintsArrayOutput) ToMachineReferenceWithHintsArrayOutput() MachineReferenceWithHintsArrayOutput {
	return o
}

func (o MachineReferenceWithHintsArrayOutput) ToMachineReferenceWithHintsArrayOutputWithContext(ctx context.Context) MachineReferenceWithHintsArrayOutput {
	return o
}

func (o MachineReferenceWithHintsArrayOutput) Index(i pulumi.IntInput) MachineReferenceWithHintsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineReferenceWithHints {
		return vs[0].([]MachineReferenceWithHints)[vs[1].(int)]
	}).(MachineReferenceWithHintsOutput)
}

type MachineReferenceWithHintsResponse struct {
	DisplayNameHint string `pulumi:"displayNameHint"`
	Id              string `pulumi:"id"`
	Kind            string `pulumi:"kind"`
	Name            string `pulumi:"name"`
	OsFamilyHint    string `pulumi:"osFamilyHint"`
	Type            string `pulumi:"type"`
}

type MachineReferenceWithHintsResponseOutput struct{ *pulumi.OutputState }

func (MachineReferenceWithHintsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineReferenceWithHintsResponse)(nil)).Elem()
}

func (o MachineReferenceWithHintsResponseOutput) ToMachineReferenceWithHintsResponseOutput() MachineReferenceWithHintsResponseOutput {
	return o
}

func (o MachineReferenceWithHintsResponseOutput) ToMachineReferenceWithHintsResponseOutputWithContext(ctx context.Context) MachineReferenceWithHintsResponseOutput {
	return o
}

func (o MachineReferenceWithHintsResponseOutput) DisplayNameHint() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.DisplayNameHint }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) OsFamilyHint() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.OsFamilyHint }).(pulumi.StringOutput)
}

func (o MachineReferenceWithHintsResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineReferenceWithHintsResponse) string { return v.Type }).(pulumi.StringOutput)
}

type MachineReferenceWithHintsResponseArrayOutput struct{ *pulumi.OutputState }

func (MachineReferenceWithHintsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineReferenceWithHintsResponse)(nil)).Elem()
}

func (o MachineReferenceWithHintsResponseArrayOutput) ToMachineReferenceWithHintsResponseArrayOutput() MachineReferenceWithHintsResponseArrayOutput {
	return o
}

func (o MachineReferenceWithHintsResponseArrayOutput) ToMachineReferenceWithHintsResponseArrayOutputWithContext(ctx context.Context) MachineReferenceWithHintsResponseArrayOutput {
	return o
}

func (o MachineReferenceWithHintsResponseArrayOutput) Index(i pulumi.IntInput) MachineReferenceWithHintsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineReferenceWithHintsResponse {
		return vs[0].([]MachineReferenceWithHintsResponse)[vs[1].(int)]
	}).(MachineReferenceWithHintsResponseOutput)
}

type PrivateLinkScopedResourceResponse struct {
	ResourceId *string `pulumi:"resourceId"`
	ScopeId    *string `pulumi:"scopeId"`
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

type RestoredLogs struct {
	EndRestoreTime   *string `pulumi:"endRestoreTime"`
	SourceTable      *string `pulumi:"sourceTable"`
	StartRestoreTime *string `pulumi:"startRestoreTime"`
}





type RestoredLogsInput interface {
	pulumi.Input

	ToRestoredLogsOutput() RestoredLogsOutput
	ToRestoredLogsOutputWithContext(context.Context) RestoredLogsOutput
}

type RestoredLogsArgs struct {
	EndRestoreTime   pulumi.StringPtrInput `pulumi:"endRestoreTime"`
	SourceTable      pulumi.StringPtrInput `pulumi:"sourceTable"`
	StartRestoreTime pulumi.StringPtrInput `pulumi:"startRestoreTime"`
}

func (RestoredLogsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RestoredLogs)(nil)).Elem()
}

func (i RestoredLogsArgs) ToRestoredLogsOutput() RestoredLogsOutput {
	return i.ToRestoredLogsOutputWithContext(context.Background())
}

func (i RestoredLogsArgs) ToRestoredLogsOutputWithContext(ctx context.Context) RestoredLogsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoredLogsOutput)
}

func (i RestoredLogsArgs) ToRestoredLogsPtrOutput() RestoredLogsPtrOutput {
	return i.ToRestoredLogsPtrOutputWithContext(context.Background())
}

func (i RestoredLogsArgs) ToRestoredLogsPtrOutputWithContext(ctx context.Context) RestoredLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoredLogsOutput).ToRestoredLogsPtrOutputWithContext(ctx)
}









type RestoredLogsPtrInput interface {
	pulumi.Input

	ToRestoredLogsPtrOutput() RestoredLogsPtrOutput
	ToRestoredLogsPtrOutputWithContext(context.Context) RestoredLogsPtrOutput
}

type restoredLogsPtrType RestoredLogsArgs

func RestoredLogsPtr(v *RestoredLogsArgs) RestoredLogsPtrInput {
	return (*restoredLogsPtrType)(v)
}

func (*restoredLogsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RestoredLogs)(nil)).Elem()
}

func (i *restoredLogsPtrType) ToRestoredLogsPtrOutput() RestoredLogsPtrOutput {
	return i.ToRestoredLogsPtrOutputWithContext(context.Background())
}

func (i *restoredLogsPtrType) ToRestoredLogsPtrOutputWithContext(ctx context.Context) RestoredLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoredLogsPtrOutput)
}

type RestoredLogsOutput struct{ *pulumi.OutputState }

func (RestoredLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestoredLogs)(nil)).Elem()
}

func (o RestoredLogsOutput) ToRestoredLogsOutput() RestoredLogsOutput {
	return o
}

func (o RestoredLogsOutput) ToRestoredLogsOutputWithContext(ctx context.Context) RestoredLogsOutput {
	return o
}

func (o RestoredLogsOutput) ToRestoredLogsPtrOutput() RestoredLogsPtrOutput {
	return o.ToRestoredLogsPtrOutputWithContext(context.Background())
}

func (o RestoredLogsOutput) ToRestoredLogsPtrOutputWithContext(ctx context.Context) RestoredLogsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestoredLogs) *RestoredLogs {
		return &v
	}).(RestoredLogsPtrOutput)
}

func (o RestoredLogsOutput) EndRestoreTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoredLogs) *string { return v.EndRestoreTime }).(pulumi.StringPtrOutput)
}

func (o RestoredLogsOutput) SourceTable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoredLogs) *string { return v.SourceTable }).(pulumi.StringPtrOutput)
}

func (o RestoredLogsOutput) StartRestoreTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoredLogs) *string { return v.StartRestoreTime }).(pulumi.StringPtrOutput)
}

type RestoredLogsPtrOutput struct{ *pulumi.OutputState }

func (RestoredLogsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestoredLogs)(nil)).Elem()
}

func (o RestoredLogsPtrOutput) ToRestoredLogsPtrOutput() RestoredLogsPtrOutput {
	return o
}

func (o RestoredLogsPtrOutput) ToRestoredLogsPtrOutputWithContext(ctx context.Context) RestoredLogsPtrOutput {
	return o
}

func (o RestoredLogsPtrOutput) Elem() RestoredLogsOutput {
	return o.ApplyT(func(v *RestoredLogs) RestoredLogs {
		if v != nil {
			return *v
		}
		var ret RestoredLogs
		return ret
	}).(RestoredLogsOutput)
}

func (o RestoredLogsPtrOutput) EndRestoreTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoredLogs) *string {
		if v == nil {
			return nil
		}
		return v.EndRestoreTime
	}).(pulumi.StringPtrOutput)
}

func (o RestoredLogsPtrOutput) SourceTable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoredLogs) *string {
		if v == nil {
			return nil
		}
		return v.SourceTable
	}).(pulumi.StringPtrOutput)
}

func (o RestoredLogsPtrOutput) StartRestoreTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoredLogs) *string {
		if v == nil {
			return nil
		}
		return v.StartRestoreTime
	}).(pulumi.StringPtrOutput)
}

type RestoredLogsResponse struct {
	EndRestoreTime   *string `pulumi:"endRestoreTime"`
	SourceTable      *string `pulumi:"sourceTable"`
	StartRestoreTime *string `pulumi:"startRestoreTime"`
}

type RestoredLogsResponseOutput struct{ *pulumi.OutputState }

func (RestoredLogsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestoredLogsResponse)(nil)).Elem()
}

func (o RestoredLogsResponseOutput) ToRestoredLogsResponseOutput() RestoredLogsResponseOutput {
	return o
}

func (o RestoredLogsResponseOutput) ToRestoredLogsResponseOutputWithContext(ctx context.Context) RestoredLogsResponseOutput {
	return o
}

func (o RestoredLogsResponseOutput) EndRestoreTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoredLogsResponse) *string { return v.EndRestoreTime }).(pulumi.StringPtrOutput)
}

func (o RestoredLogsResponseOutput) SourceTable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoredLogsResponse) *string { return v.SourceTable }).(pulumi.StringPtrOutput)
}

func (o RestoredLogsResponseOutput) StartRestoreTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoredLogsResponse) *string { return v.StartRestoreTime }).(pulumi.StringPtrOutput)
}

type RestoredLogsResponsePtrOutput struct{ *pulumi.OutputState }

func (RestoredLogsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestoredLogsResponse)(nil)).Elem()
}

func (o RestoredLogsResponsePtrOutput) ToRestoredLogsResponsePtrOutput() RestoredLogsResponsePtrOutput {
	return o
}

func (o RestoredLogsResponsePtrOutput) ToRestoredLogsResponsePtrOutputWithContext(ctx context.Context) RestoredLogsResponsePtrOutput {
	return o
}

func (o RestoredLogsResponsePtrOutput) Elem() RestoredLogsResponseOutput {
	return o.ApplyT(func(v *RestoredLogsResponse) RestoredLogsResponse {
		if v != nil {
			return *v
		}
		var ret RestoredLogsResponse
		return ret
	}).(RestoredLogsResponseOutput)
}

func (o RestoredLogsResponsePtrOutput) EndRestoreTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoredLogsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndRestoreTime
	}).(pulumi.StringPtrOutput)
}

func (o RestoredLogsResponsePtrOutput) SourceTable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoredLogsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceTable
	}).(pulumi.StringPtrOutput)
}

func (o RestoredLogsResponsePtrOutput) StartRestoreTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoredLogsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartRestoreTime
	}).(pulumi.StringPtrOutput)
}

type ResultStatisticsResponse struct {
	IngestedRecords int     `pulumi:"ingestedRecords"`
	Progress        float64 `pulumi:"progress"`
}

type ResultStatisticsResponseOutput struct{ *pulumi.OutputState }

func (ResultStatisticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResultStatisticsResponse)(nil)).Elem()
}

func (o ResultStatisticsResponseOutput) ToResultStatisticsResponseOutput() ResultStatisticsResponseOutput {
	return o
}

func (o ResultStatisticsResponseOutput) ToResultStatisticsResponseOutputWithContext(ctx context.Context) ResultStatisticsResponseOutput {
	return o
}

func (o ResultStatisticsResponseOutput) IngestedRecords() pulumi.IntOutput {
	return o.ApplyT(func(v ResultStatisticsResponse) int { return v.IngestedRecords }).(pulumi.IntOutput)
}

func (o ResultStatisticsResponseOutput) Progress() pulumi.Float64Output {
	return o.ApplyT(func(v ResultStatisticsResponse) float64 { return v.Progress }).(pulumi.Float64Output)
}

type ResultStatisticsResponsePtrOutput struct{ *pulumi.OutputState }

func (ResultStatisticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResultStatisticsResponse)(nil)).Elem()
}

func (o ResultStatisticsResponsePtrOutput) ToResultStatisticsResponsePtrOutput() ResultStatisticsResponsePtrOutput {
	return o
}

func (o ResultStatisticsResponsePtrOutput) ToResultStatisticsResponsePtrOutputWithContext(ctx context.Context) ResultStatisticsResponsePtrOutput {
	return o
}

func (o ResultStatisticsResponsePtrOutput) Elem() ResultStatisticsResponseOutput {
	return o.ApplyT(func(v *ResultStatisticsResponse) ResultStatisticsResponse {
		if v != nil {
			return *v
		}
		var ret ResultStatisticsResponse
		return ret
	}).(ResultStatisticsResponseOutput)
}

func (o ResultStatisticsResponsePtrOutput) IngestedRecords() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResultStatisticsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.IngestedRecords
	}).(pulumi.IntPtrOutput)
}

func (o ResultStatisticsResponsePtrOutput) Progress() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResultStatisticsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Progress
	}).(pulumi.Float64PtrOutput)
}

type Schema struct {
	Columns     []Column `pulumi:"columns"`
	Description *string  `pulumi:"description"`
	DisplayName *string  `pulumi:"displayName"`
	Name        *string  `pulumi:"name"`
}





type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(context.Context) SchemaOutput
}

type SchemaArgs struct {
	Columns     ColumnArrayInput      `pulumi:"columns"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Schema)(nil)).Elem()
}

func (i SchemaArgs) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i SchemaArgs) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

func (i SchemaArgs) ToSchemaPtrOutput() SchemaPtrOutput {
	return i.ToSchemaPtrOutputWithContext(context.Background())
}

func (i SchemaArgs) ToSchemaPtrOutputWithContext(ctx context.Context) SchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput).ToSchemaPtrOutputWithContext(ctx)
}









type SchemaPtrInput interface {
	pulumi.Input

	ToSchemaPtrOutput() SchemaPtrOutput
	ToSchemaPtrOutputWithContext(context.Context) SchemaPtrOutput
}

type schemaPtrType SchemaArgs

func SchemaPtr(v *SchemaArgs) SchemaPtrInput {
	return (*schemaPtrType)(v)
}

func (*schemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (i *schemaPtrType) ToSchemaPtrOutput() SchemaPtrOutput {
	return i.ToSchemaPtrOutputWithContext(context.Background())
}

func (i *schemaPtrType) ToSchemaPtrOutputWithContext(ctx context.Context) SchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaPtrOutput)
}

type SchemaOutput struct{ *pulumi.OutputState }

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schema)(nil)).Elem()
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaPtrOutput() SchemaPtrOutput {
	return o.ToSchemaPtrOutputWithContext(context.Background())
}

func (o SchemaOutput) ToSchemaPtrOutputWithContext(ctx context.Context) SchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Schema) *Schema {
		return &v
	}).(SchemaPtrOutput)
}

func (o SchemaOutput) Columns() ColumnArrayOutput {
	return o.ApplyT(func(v Schema) []Column { return v.Columns }).(ColumnArrayOutput)
}

func (o SchemaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Schema) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SchemaOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Schema) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o SchemaOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Schema) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SchemaPtrOutput struct{ *pulumi.OutputState }

func (SchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (o SchemaPtrOutput) ToSchemaPtrOutput() SchemaPtrOutput {
	return o
}

func (o SchemaPtrOutput) ToSchemaPtrOutputWithContext(ctx context.Context) SchemaPtrOutput {
	return o
}

func (o SchemaPtrOutput) Elem() SchemaOutput {
	return o.ApplyT(func(v *Schema) Schema {
		if v != nil {
			return *v
		}
		var ret Schema
		return ret
	}).(SchemaOutput)
}

func (o SchemaPtrOutput) Columns() ColumnArrayOutput {
	return o.ApplyT(func(v *Schema) []Column {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(ColumnArrayOutput)
}

func (o SchemaPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SchemaPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SchemaPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SchemaResponse struct {
	Categories      []string              `pulumi:"categories"`
	Columns         []ColumnResponse      `pulumi:"columns"`
	Description     *string               `pulumi:"description"`
	DisplayName     *string               `pulumi:"displayName"`
	Labels          []string              `pulumi:"labels"`
	Name            *string               `pulumi:"name"`
	RestoredLogs    RestoredLogsResponse  `pulumi:"restoredLogs"`
	SearchResults   SearchResultsResponse `pulumi:"searchResults"`
	Solutions       []string              `pulumi:"solutions"`
	Source          string                `pulumi:"source"`
	StandardColumns []ColumnResponse      `pulumi:"standardColumns"`
	TableSubType    string                `pulumi:"tableSubType"`
	TableType       string                `pulumi:"tableType"`
}

type SchemaResponseOutput struct{ *pulumi.OutputState }

func (SchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaResponse)(nil)).Elem()
}

func (o SchemaResponseOutput) ToSchemaResponseOutput() SchemaResponseOutput {
	return o
}

func (o SchemaResponseOutput) ToSchemaResponseOutputWithContext(ctx context.Context) SchemaResponseOutput {
	return o
}

func (o SchemaResponseOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SchemaResponse) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o SchemaResponseOutput) Columns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v SchemaResponse) []ColumnResponse { return v.Columns }).(ColumnResponseArrayOutput)
}

func (o SchemaResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchemaResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SchemaResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchemaResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o SchemaResponseOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SchemaResponse) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o SchemaResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchemaResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SchemaResponseOutput) RestoredLogs() RestoredLogsResponseOutput {
	return o.ApplyT(func(v SchemaResponse) RestoredLogsResponse { return v.RestoredLogs }).(RestoredLogsResponseOutput)
}

func (o SchemaResponseOutput) SearchResults() SearchResultsResponseOutput {
	return o.ApplyT(func(v SchemaResponse) SearchResultsResponse { return v.SearchResults }).(SearchResultsResponseOutput)
}

func (o SchemaResponseOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SchemaResponse) []string { return v.Solutions }).(pulumi.StringArrayOutput)
}

func (o SchemaResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v SchemaResponse) string { return v.Source }).(pulumi.StringOutput)
}

func (o SchemaResponseOutput) StandardColumns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v SchemaResponse) []ColumnResponse { return v.StandardColumns }).(ColumnResponseArrayOutput)
}

func (o SchemaResponseOutput) TableSubType() pulumi.StringOutput {
	return o.ApplyT(func(v SchemaResponse) string { return v.TableSubType }).(pulumi.StringOutput)
}

func (o SchemaResponseOutput) TableType() pulumi.StringOutput {
	return o.ApplyT(func(v SchemaResponse) string { return v.TableType }).(pulumi.StringOutput)
}

type SchemaResponsePtrOutput struct{ *pulumi.OutputState }

func (SchemaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaResponse)(nil)).Elem()
}

func (o SchemaResponsePtrOutput) ToSchemaResponsePtrOutput() SchemaResponsePtrOutput {
	return o
}

func (o SchemaResponsePtrOutput) ToSchemaResponsePtrOutputWithContext(ctx context.Context) SchemaResponsePtrOutput {
	return o
}

func (o SchemaResponsePtrOutput) Elem() SchemaResponseOutput {
	return o.ApplyT(func(v *SchemaResponse) SchemaResponse {
		if v != nil {
			return *v
		}
		var ret SchemaResponse
		return ret
	}).(SchemaResponseOutput)
}

func (o SchemaResponsePtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SchemaResponse) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o SchemaResponsePtrOutput) Columns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v *SchemaResponse) []ColumnResponse {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(ColumnResponseArrayOutput)
}

func (o SchemaResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SchemaResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SchemaResponsePtrOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SchemaResponse) []string {
		if v == nil {
			return nil
		}
		return v.Labels
	}).(pulumi.StringArrayOutput)
}

func (o SchemaResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SchemaResponsePtrOutput) RestoredLogs() RestoredLogsResponsePtrOutput {
	return o.ApplyT(func(v *SchemaResponse) *RestoredLogsResponse {
		if v == nil {
			return nil
		}
		return &v.RestoredLogs
	}).(RestoredLogsResponsePtrOutput)
}

func (o SchemaResponsePtrOutput) SearchResults() SearchResultsResponsePtrOutput {
	return o.ApplyT(func(v *SchemaResponse) *SearchResultsResponse {
		if v == nil {
			return nil
		}
		return &v.SearchResults
	}).(SearchResultsResponsePtrOutput)
}

func (o SchemaResponsePtrOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SchemaResponse) []string {
		if v == nil {
			return nil
		}
		return v.Solutions
	}).(pulumi.StringArrayOutput)
}

func (o SchemaResponsePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Source
	}).(pulumi.StringPtrOutput)
}

func (o SchemaResponsePtrOutput) StandardColumns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v *SchemaResponse) []ColumnResponse {
		if v == nil {
			return nil
		}
		return v.StandardColumns
	}).(ColumnResponseArrayOutput)
}

func (o SchemaResponsePtrOutput) TableSubType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TableSubType
	}).(pulumi.StringPtrOutput)
}

func (o SchemaResponsePtrOutput) TableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TableType
	}).(pulumi.StringPtrOutput)
}

type SearchResults struct {
	Description     *string `pulumi:"description"`
	EndSearchTime   *string `pulumi:"endSearchTime"`
	Limit           *int    `pulumi:"limit"`
	Query           *string `pulumi:"query"`
	StartSearchTime *string `pulumi:"startSearchTime"`
}





type SearchResultsInput interface {
	pulumi.Input

	ToSearchResultsOutput() SearchResultsOutput
	ToSearchResultsOutputWithContext(context.Context) SearchResultsOutput
}

type SearchResultsArgs struct {
	Description     pulumi.StringPtrInput `pulumi:"description"`
	EndSearchTime   pulumi.StringPtrInput `pulumi:"endSearchTime"`
	Limit           pulumi.IntPtrInput    `pulumi:"limit"`
	Query           pulumi.StringPtrInput `pulumi:"query"`
	StartSearchTime pulumi.StringPtrInput `pulumi:"startSearchTime"`
}

func (SearchResultsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchResults)(nil)).Elem()
}

func (i SearchResultsArgs) ToSearchResultsOutput() SearchResultsOutput {
	return i.ToSearchResultsOutputWithContext(context.Background())
}

func (i SearchResultsArgs) ToSearchResultsOutputWithContext(ctx context.Context) SearchResultsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchResultsOutput)
}

func (i SearchResultsArgs) ToSearchResultsPtrOutput() SearchResultsPtrOutput {
	return i.ToSearchResultsPtrOutputWithContext(context.Background())
}

func (i SearchResultsArgs) ToSearchResultsPtrOutputWithContext(ctx context.Context) SearchResultsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchResultsOutput).ToSearchResultsPtrOutputWithContext(ctx)
}









type SearchResultsPtrInput interface {
	pulumi.Input

	ToSearchResultsPtrOutput() SearchResultsPtrOutput
	ToSearchResultsPtrOutputWithContext(context.Context) SearchResultsPtrOutput
}

type searchResultsPtrType SearchResultsArgs

func SearchResultsPtr(v *SearchResultsArgs) SearchResultsPtrInput {
	return (*searchResultsPtrType)(v)
}

func (*searchResultsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchResults)(nil)).Elem()
}

func (i *searchResultsPtrType) ToSearchResultsPtrOutput() SearchResultsPtrOutput {
	return i.ToSearchResultsPtrOutputWithContext(context.Background())
}

func (i *searchResultsPtrType) ToSearchResultsPtrOutputWithContext(ctx context.Context) SearchResultsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SearchResultsPtrOutput)
}

type SearchResultsOutput struct{ *pulumi.OutputState }

func (SearchResultsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchResults)(nil)).Elem()
}

func (o SearchResultsOutput) ToSearchResultsOutput() SearchResultsOutput {
	return o
}

func (o SearchResultsOutput) ToSearchResultsOutputWithContext(ctx context.Context) SearchResultsOutput {
	return o
}

func (o SearchResultsOutput) ToSearchResultsPtrOutput() SearchResultsPtrOutput {
	return o.ToSearchResultsPtrOutputWithContext(context.Background())
}

func (o SearchResultsOutput) ToSearchResultsPtrOutputWithContext(ctx context.Context) SearchResultsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SearchResults) *SearchResults {
		return &v
	}).(SearchResultsPtrOutput)
}

func (o SearchResultsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchResults) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SearchResultsOutput) EndSearchTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchResults) *string { return v.EndSearchTime }).(pulumi.StringPtrOutput)
}

func (o SearchResultsOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SearchResults) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o SearchResultsOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchResults) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o SearchResultsOutput) StartSearchTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchResults) *string { return v.StartSearchTime }).(pulumi.StringPtrOutput)
}

type SearchResultsPtrOutput struct{ *pulumi.OutputState }

func (SearchResultsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchResults)(nil)).Elem()
}

func (o SearchResultsPtrOutput) ToSearchResultsPtrOutput() SearchResultsPtrOutput {
	return o
}

func (o SearchResultsPtrOutput) ToSearchResultsPtrOutputWithContext(ctx context.Context) SearchResultsPtrOutput {
	return o
}

func (o SearchResultsPtrOutput) Elem() SearchResultsOutput {
	return o.ApplyT(func(v *SearchResults) SearchResults {
		if v != nil {
			return *v
		}
		var ret SearchResults
		return ret
	}).(SearchResultsOutput)
}

func (o SearchResultsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchResults) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SearchResultsPtrOutput) EndSearchTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchResults) *string {
		if v == nil {
			return nil
		}
		return v.EndSearchTime
	}).(pulumi.StringPtrOutput)
}

func (o SearchResultsPtrOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SearchResults) *int {
		if v == nil {
			return nil
		}
		return v.Limit
	}).(pulumi.IntPtrOutput)
}

func (o SearchResultsPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchResults) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o SearchResultsPtrOutput) StartSearchTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchResults) *string {
		if v == nil {
			return nil
		}
		return v.StartSearchTime
	}).(pulumi.StringPtrOutput)
}

type SearchResultsResponse struct {
	Description     *string `pulumi:"description"`
	EndSearchTime   *string `pulumi:"endSearchTime"`
	Limit           *int    `pulumi:"limit"`
	Query           *string `pulumi:"query"`
	SourceTable     string  `pulumi:"sourceTable"`
	StartSearchTime *string `pulumi:"startSearchTime"`
}

type SearchResultsResponseOutput struct{ *pulumi.OutputState }

func (SearchResultsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SearchResultsResponse)(nil)).Elem()
}

func (o SearchResultsResponseOutput) ToSearchResultsResponseOutput() SearchResultsResponseOutput {
	return o
}

func (o SearchResultsResponseOutput) ToSearchResultsResponseOutputWithContext(ctx context.Context) SearchResultsResponseOutput {
	return o
}

func (o SearchResultsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchResultsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SearchResultsResponseOutput) EndSearchTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchResultsResponse) *string { return v.EndSearchTime }).(pulumi.StringPtrOutput)
}

func (o SearchResultsResponseOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SearchResultsResponse) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o SearchResultsResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchResultsResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o SearchResultsResponseOutput) SourceTable() pulumi.StringOutput {
	return o.ApplyT(func(v SearchResultsResponse) string { return v.SourceTable }).(pulumi.StringOutput)
}

func (o SearchResultsResponseOutput) StartSearchTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SearchResultsResponse) *string { return v.StartSearchTime }).(pulumi.StringPtrOutput)
}

type SearchResultsResponsePtrOutput struct{ *pulumi.OutputState }

func (SearchResultsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SearchResultsResponse)(nil)).Elem()
}

func (o SearchResultsResponsePtrOutput) ToSearchResultsResponsePtrOutput() SearchResultsResponsePtrOutput {
	return o
}

func (o SearchResultsResponsePtrOutput) ToSearchResultsResponsePtrOutputWithContext(ctx context.Context) SearchResultsResponsePtrOutput {
	return o
}

func (o SearchResultsResponsePtrOutput) Elem() SearchResultsResponseOutput {
	return o.ApplyT(func(v *SearchResultsResponse) SearchResultsResponse {
		if v != nil {
			return *v
		}
		var ret SearchResultsResponse
		return ret
	}).(SearchResultsResponseOutput)
}

func (o SearchResultsResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchResultsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SearchResultsResponsePtrOutput) EndSearchTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchResultsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndSearchTime
	}).(pulumi.StringPtrOutput)
}

func (o SearchResultsResponsePtrOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SearchResultsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Limit
	}).(pulumi.IntPtrOutput)
}

func (o SearchResultsResponsePtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchResultsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o SearchResultsResponsePtrOutput) SourceTable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchResultsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceTable
	}).(pulumi.StringPtrOutput)
}

func (o SearchResultsResponsePtrOutput) StartSearchTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SearchResultsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartSearchTime
	}).(pulumi.StringPtrOutput)
}

type StorageAccount struct {
	Id  string `pulumi:"id"`
	Key string `pulumi:"key"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Id  pulumi.StringInput `pulumi:"id"`
	Key pulumi.StringInput `pulumi:"key"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Key }).(pulumi.StringOutput)
}

type StorageAccountResponse struct {
	Id  string `pulumi:"id"`
	Key string `pulumi:"key"`
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Key }).(pulumi.StringOutput)
}

type StorageInsightStatusResponse struct {
	Description *string `pulumi:"description"`
	State       string  `pulumi:"state"`
}

type StorageInsightStatusResponseOutput struct{ *pulumi.OutputState }

func (StorageInsightStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageInsightStatusResponse)(nil)).Elem()
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponseOutput() StorageInsightStatusResponseOutput {
	return o
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponseOutputWithContext(ctx context.Context) StorageInsightStatusResponseOutput {
	return o
}

func (o StorageInsightStatusResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageInsightStatusResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StorageInsightStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v StorageInsightStatusResponse) string { return v.State }).(pulumi.StringOutput)
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

type Tag struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(context.Context) TagOutput
}

type TagArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (i TagArgs) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i TagArgs) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}





type TagArrayInput interface {
	pulumi.Input

	ToTagArrayOutput() TagArrayOutput
	ToTagArrayOutputWithContext(context.Context) TagArrayOutput
}

type TagArray []TagInput

func (TagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil)).Elem()
}

func (i TagArray) ToTagArrayOutput() TagArrayOutput {
	return i.ToTagArrayOutputWithContext(context.Background())
}

func (i TagArray) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagArrayOutput)
}

type TagOutput struct{ *pulumi.OutputState }

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

func (o TagOutput) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return o
}

func (o TagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Tag) string { return v.Name }).(pulumi.StringOutput)
}

func (o TagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v Tag) string { return v.Value }).(pulumi.StringOutput)
}

type TagArrayOutput struct{ *pulumi.OutputState }

func (TagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil)).Elem()
}

func (o TagArrayOutput) ToTagArrayOutput() TagArrayOutput {
	return o
}

func (o TagArrayOutput) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return o
}

func (o TagArrayOutput) Index(i pulumi.IntInput) TagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Tag {
		return vs[0].([]Tag)[vs[1].(int)]
	}).(TagOutput)
}

type TagResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type TagResponseOutput struct{ *pulumi.OutputState }

func (TagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagResponse)(nil)).Elem()
}

func (o TagResponseOutput) ToTagResponseOutput() TagResponseOutput {
	return o
}

func (o TagResponseOutput) ToTagResponseOutputWithContext(ctx context.Context) TagResponseOutput {
	return o
}

func (o TagResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TagResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TagResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TagResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TagResponseArrayOutput struct{ *pulumi.OutputState }

func (TagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagResponse)(nil)).Elem()
}

func (o TagResponseArrayOutput) ToTagResponseArrayOutput() TagResponseArrayOutput {
	return o
}

func (o TagResponseArrayOutput) ToTagResponseArrayOutputWithContext(ctx context.Context) TagResponseArrayOutput {
	return o
}

func (o TagResponseArrayOutput) Index(i pulumi.IntInput) TagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagResponse {
		return vs[0].([]TagResponse)[vs[1].(int)]
	}).(TagResponseOutput)
}

type UserIdentityPropertiesResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
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
	pulumi.RegisterOutputType(AssociatedWorkspaceResponseOutput{})
	pulumi.RegisterOutputType(AssociatedWorkspaceResponseArrayOutput{})
	pulumi.RegisterOutputType(CapacityReservationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CapacityReservationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuOutput{})
	pulumi.RegisterOutputType(ClusterSkuPtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponseOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ColumnOutput{})
	pulumi.RegisterOutputType(ColumnArrayOutput{})
	pulumi.RegisterOutputType(ColumnResponseOutput{})
	pulumi.RegisterOutputType(ColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesRelatedOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput{})
	pulumi.RegisterOutputType(MachineReferenceWithHintsOutput{})
	pulumi.RegisterOutputType(MachineReferenceWithHintsArrayOutput{})
	pulumi.RegisterOutputType(MachineReferenceWithHintsResponseOutput{})
	pulumi.RegisterOutputType(MachineReferenceWithHintsResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(RestoredLogsOutput{})
	pulumi.RegisterOutputType(RestoredLogsPtrOutput{})
	pulumi.RegisterOutputType(RestoredLogsResponseOutput{})
	pulumi.RegisterOutputType(RestoredLogsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResultStatisticsResponseOutput{})
	pulumi.RegisterOutputType(ResultStatisticsResponsePtrOutput{})
	pulumi.RegisterOutputType(SchemaOutput{})
	pulumi.RegisterOutputType(SchemaPtrOutput{})
	pulumi.RegisterOutputType(SchemaResponseOutput{})
	pulumi.RegisterOutputType(SchemaResponsePtrOutput{})
	pulumi.RegisterOutputType(SearchResultsOutput{})
	pulumi.RegisterOutputType(SearchResultsPtrOutput{})
	pulumi.RegisterOutputType(SearchResultsResponseOutput{})
	pulumi.RegisterOutputType(SearchResultsResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageInsightStatusResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TagOutput{})
	pulumi.RegisterOutputType(TagArrayOutput{})
	pulumi.RegisterOutputType(TagResponseOutput{})
	pulumi.RegisterOutputType(TagResponseArrayOutput{})
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
