


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateManagementGroupDetails struct {
	Parent *CreateParentGroupInfo `pulumi:"parent"`
}





type CreateManagementGroupDetailsInput interface {
	pulumi.Input

	ToCreateManagementGroupDetailsOutput() CreateManagementGroupDetailsOutput
	ToCreateManagementGroupDetailsOutputWithContext(context.Context) CreateManagementGroupDetailsOutput
}

type CreateManagementGroupDetailsArgs struct {
	Parent CreateParentGroupInfoPtrInput `pulumi:"parent"`
}

func (CreateManagementGroupDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateManagementGroupDetails)(nil)).Elem()
}

func (i CreateManagementGroupDetailsArgs) ToCreateManagementGroupDetailsOutput() CreateManagementGroupDetailsOutput {
	return i.ToCreateManagementGroupDetailsOutputWithContext(context.Background())
}

func (i CreateManagementGroupDetailsArgs) ToCreateManagementGroupDetailsOutputWithContext(ctx context.Context) CreateManagementGroupDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateManagementGroupDetailsOutput)
}

func (i CreateManagementGroupDetailsArgs) ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput {
	return i.ToCreateManagementGroupDetailsPtrOutputWithContext(context.Background())
}

func (i CreateManagementGroupDetailsArgs) ToCreateManagementGroupDetailsPtrOutputWithContext(ctx context.Context) CreateManagementGroupDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateManagementGroupDetailsOutput).ToCreateManagementGroupDetailsPtrOutputWithContext(ctx)
}









type CreateManagementGroupDetailsPtrInput interface {
	pulumi.Input

	ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput
	ToCreateManagementGroupDetailsPtrOutputWithContext(context.Context) CreateManagementGroupDetailsPtrOutput
}

type createManagementGroupDetailsPtrType CreateManagementGroupDetailsArgs

func CreateManagementGroupDetailsPtr(v *CreateManagementGroupDetailsArgs) CreateManagementGroupDetailsPtrInput {
	return (*createManagementGroupDetailsPtrType)(v)
}

func (*createManagementGroupDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateManagementGroupDetails)(nil)).Elem()
}

func (i *createManagementGroupDetailsPtrType) ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput {
	return i.ToCreateManagementGroupDetailsPtrOutputWithContext(context.Background())
}

func (i *createManagementGroupDetailsPtrType) ToCreateManagementGroupDetailsPtrOutputWithContext(ctx context.Context) CreateManagementGroupDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateManagementGroupDetailsPtrOutput)
}

type CreateManagementGroupDetailsOutput struct{ *pulumi.OutputState }

func (CreateManagementGroupDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateManagementGroupDetails)(nil)).Elem()
}

func (o CreateManagementGroupDetailsOutput) ToCreateManagementGroupDetailsOutput() CreateManagementGroupDetailsOutput {
	return o
}

func (o CreateManagementGroupDetailsOutput) ToCreateManagementGroupDetailsOutputWithContext(ctx context.Context) CreateManagementGroupDetailsOutput {
	return o
}

func (o CreateManagementGroupDetailsOutput) ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput {
	return o.ToCreateManagementGroupDetailsPtrOutputWithContext(context.Background())
}

func (o CreateManagementGroupDetailsOutput) ToCreateManagementGroupDetailsPtrOutputWithContext(ctx context.Context) CreateManagementGroupDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateManagementGroupDetails) *CreateManagementGroupDetails {
		return &v
	}).(CreateManagementGroupDetailsPtrOutput)
}

func (o CreateManagementGroupDetailsOutput) Parent() CreateParentGroupInfoPtrOutput {
	return o.ApplyT(func(v CreateManagementGroupDetails) *CreateParentGroupInfo { return v.Parent }).(CreateParentGroupInfoPtrOutput)
}

type CreateManagementGroupDetailsPtrOutput struct{ *pulumi.OutputState }

func (CreateManagementGroupDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateManagementGroupDetails)(nil)).Elem()
}

func (o CreateManagementGroupDetailsPtrOutput) ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput {
	return o
}

func (o CreateManagementGroupDetailsPtrOutput) ToCreateManagementGroupDetailsPtrOutputWithContext(ctx context.Context) CreateManagementGroupDetailsPtrOutput {
	return o
}

func (o CreateManagementGroupDetailsPtrOutput) Elem() CreateManagementGroupDetailsOutput {
	return o.ApplyT(func(v *CreateManagementGroupDetails) CreateManagementGroupDetails {
		if v != nil {
			return *v
		}
		var ret CreateManagementGroupDetails
		return ret
	}).(CreateManagementGroupDetailsOutput)
}

func (o CreateManagementGroupDetailsPtrOutput) Parent() CreateParentGroupInfoPtrOutput {
	return o.ApplyT(func(v *CreateManagementGroupDetails) *CreateParentGroupInfo {
		if v == nil {
			return nil
		}
		return v.Parent
	}).(CreateParentGroupInfoPtrOutput)
}

type CreateParentGroupInfo struct {
	Id *string `pulumi:"id"`
}





type CreateParentGroupInfoInput interface {
	pulumi.Input

	ToCreateParentGroupInfoOutput() CreateParentGroupInfoOutput
	ToCreateParentGroupInfoOutputWithContext(context.Context) CreateParentGroupInfoOutput
}

type CreateParentGroupInfoArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (CreateParentGroupInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateParentGroupInfo)(nil)).Elem()
}

func (i CreateParentGroupInfoArgs) ToCreateParentGroupInfoOutput() CreateParentGroupInfoOutput {
	return i.ToCreateParentGroupInfoOutputWithContext(context.Background())
}

func (i CreateParentGroupInfoArgs) ToCreateParentGroupInfoOutputWithContext(ctx context.Context) CreateParentGroupInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateParentGroupInfoOutput)
}

func (i CreateParentGroupInfoArgs) ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput {
	return i.ToCreateParentGroupInfoPtrOutputWithContext(context.Background())
}

func (i CreateParentGroupInfoArgs) ToCreateParentGroupInfoPtrOutputWithContext(ctx context.Context) CreateParentGroupInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateParentGroupInfoOutput).ToCreateParentGroupInfoPtrOutputWithContext(ctx)
}









type CreateParentGroupInfoPtrInput interface {
	pulumi.Input

	ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput
	ToCreateParentGroupInfoPtrOutputWithContext(context.Context) CreateParentGroupInfoPtrOutput
}

type createParentGroupInfoPtrType CreateParentGroupInfoArgs

func CreateParentGroupInfoPtr(v *CreateParentGroupInfoArgs) CreateParentGroupInfoPtrInput {
	return (*createParentGroupInfoPtrType)(v)
}

func (*createParentGroupInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateParentGroupInfo)(nil)).Elem()
}

func (i *createParentGroupInfoPtrType) ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput {
	return i.ToCreateParentGroupInfoPtrOutputWithContext(context.Background())
}

func (i *createParentGroupInfoPtrType) ToCreateParentGroupInfoPtrOutputWithContext(ctx context.Context) CreateParentGroupInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateParentGroupInfoPtrOutput)
}

type CreateParentGroupInfoOutput struct{ *pulumi.OutputState }

func (CreateParentGroupInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateParentGroupInfo)(nil)).Elem()
}

func (o CreateParentGroupInfoOutput) ToCreateParentGroupInfoOutput() CreateParentGroupInfoOutput {
	return o
}

func (o CreateParentGroupInfoOutput) ToCreateParentGroupInfoOutputWithContext(ctx context.Context) CreateParentGroupInfoOutput {
	return o
}

func (o CreateParentGroupInfoOutput) ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput {
	return o.ToCreateParentGroupInfoPtrOutputWithContext(context.Background())
}

func (o CreateParentGroupInfoOutput) ToCreateParentGroupInfoPtrOutputWithContext(ctx context.Context) CreateParentGroupInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateParentGroupInfo) *CreateParentGroupInfo {
		return &v
	}).(CreateParentGroupInfoPtrOutput)
}

func (o CreateParentGroupInfoOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateParentGroupInfo) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type CreateParentGroupInfoPtrOutput struct{ *pulumi.OutputState }

func (CreateParentGroupInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateParentGroupInfo)(nil)).Elem()
}

func (o CreateParentGroupInfoPtrOutput) ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput {
	return o
}

func (o CreateParentGroupInfoPtrOutput) ToCreateParentGroupInfoPtrOutputWithContext(ctx context.Context) CreateParentGroupInfoPtrOutput {
	return o
}

func (o CreateParentGroupInfoPtrOutput) Elem() CreateParentGroupInfoOutput {
	return o.ApplyT(func(v *CreateParentGroupInfo) CreateParentGroupInfo {
		if v != nil {
			return *v
		}
		var ret CreateParentGroupInfo
		return ret
	}).(CreateParentGroupInfoOutput)
}

func (o CreateParentGroupInfoPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateParentGroupInfo) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type DescendantParentGroupInfoResponse struct {
	Id *string `pulumi:"id"`
}

type DescendantParentGroupInfoResponseOutput struct{ *pulumi.OutputState }

func (DescendantParentGroupInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DescendantParentGroupInfoResponse)(nil)).Elem()
}

func (o DescendantParentGroupInfoResponseOutput) ToDescendantParentGroupInfoResponseOutput() DescendantParentGroupInfoResponseOutput {
	return o
}

func (o DescendantParentGroupInfoResponseOutput) ToDescendantParentGroupInfoResponseOutputWithContext(ctx context.Context) DescendantParentGroupInfoResponseOutput {
	return o
}

func (o DescendantParentGroupInfoResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DescendantParentGroupInfoResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type DescendantParentGroupInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (DescendantParentGroupInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DescendantParentGroupInfoResponse)(nil)).Elem()
}

func (o DescendantParentGroupInfoResponsePtrOutput) ToDescendantParentGroupInfoResponsePtrOutput() DescendantParentGroupInfoResponsePtrOutput {
	return o
}

func (o DescendantParentGroupInfoResponsePtrOutput) ToDescendantParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) DescendantParentGroupInfoResponsePtrOutput {
	return o
}

func (o DescendantParentGroupInfoResponsePtrOutput) Elem() DescendantParentGroupInfoResponseOutput {
	return o.ApplyT(func(v *DescendantParentGroupInfoResponse) DescendantParentGroupInfoResponse {
		if v != nil {
			return *v
		}
		var ret DescendantParentGroupInfoResponse
		return ret
	}).(DescendantParentGroupInfoResponseOutput)
}

func (o DescendantParentGroupInfoResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DescendantParentGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type EntityInfoResponse struct {
	DisplayName            *string                        `pulumi:"displayName"`
	Id                     string                         `pulumi:"id"`
	InheritedPermissions   *string                        `pulumi:"inheritedPermissions"`
	Name                   string                         `pulumi:"name"`
	NumberOfChildGroups    *int                           `pulumi:"numberOfChildGroups"`
	NumberOfChildren       *int                           `pulumi:"numberOfChildren"`
	NumberOfDescendants    *int                           `pulumi:"numberOfDescendants"`
	Parent                 *EntityParentGroupInfoResponse `pulumi:"parent"`
	ParentDisplayNameChain []string                       `pulumi:"parentDisplayNameChain"`
	ParentNameChain        []string                       `pulumi:"parentNameChain"`
	Permissions            *string                        `pulumi:"permissions"`
	TenantId               *string                        `pulumi:"tenantId"`
	Type                   string                         `pulumi:"type"`
}

type EntityParentGroupInfoResponse struct {
	Id *string `pulumi:"id"`
}

type ManagementGroupChildInfoResponse struct {
	Children    []ManagementGroupChildInfoResponse `pulumi:"children"`
	DisplayName *string                            `pulumi:"displayName"`
	Id          *string                            `pulumi:"id"`
	Name        *string                            `pulumi:"name"`
	Type        *string                            `pulumi:"type"`
}

type ManagementGroupChildInfoResponseOutput struct{ *pulumi.OutputState }

func (ManagementGroupChildInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupChildInfoResponse)(nil)).Elem()
}

func (o ManagementGroupChildInfoResponseOutput) ToManagementGroupChildInfoResponseOutput() ManagementGroupChildInfoResponseOutput {
	return o
}

func (o ManagementGroupChildInfoResponseOutput) ToManagementGroupChildInfoResponseOutputWithContext(ctx context.Context) ManagementGroupChildInfoResponseOutput {
	return o
}

func (o ManagementGroupChildInfoResponseOutput) Children() ManagementGroupChildInfoResponseArrayOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) []ManagementGroupChildInfoResponse { return v.Children }).(ManagementGroupChildInfoResponseArrayOutput)
}

func (o ManagementGroupChildInfoResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupChildInfoResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupChildInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupChildInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagementGroupChildInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementGroupChildInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupChildInfoResponse)(nil)).Elem()
}

func (o ManagementGroupChildInfoResponseArrayOutput) ToManagementGroupChildInfoResponseArrayOutput() ManagementGroupChildInfoResponseArrayOutput {
	return o
}

func (o ManagementGroupChildInfoResponseArrayOutput) ToManagementGroupChildInfoResponseArrayOutputWithContext(ctx context.Context) ManagementGroupChildInfoResponseArrayOutput {
	return o
}

func (o ManagementGroupChildInfoResponseArrayOutput) Index(i pulumi.IntInput) ManagementGroupChildInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementGroupChildInfoResponse {
		return vs[0].([]ManagementGroupChildInfoResponse)[vs[1].(int)]
	}).(ManagementGroupChildInfoResponseOutput)
}

type ManagementGroupDetailsResponse struct {
	ManagementGroupAncestors      []string                             `pulumi:"managementGroupAncestors"`
	ManagementGroupAncestorsChain []ManagementGroupPathElementResponse `pulumi:"managementGroupAncestorsChain"`
	Parent                        *ParentGroupInfoResponse             `pulumi:"parent"`
	Path                          []ManagementGroupPathElementResponse `pulumi:"path"`
	UpdatedBy                     *string                              `pulumi:"updatedBy"`
	UpdatedTime                   *string                              `pulumi:"updatedTime"`
	Version                       *float64                             `pulumi:"version"`
}

type ManagementGroupDetailsResponseOutput struct{ *pulumi.OutputState }

func (ManagementGroupDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupDetailsResponse)(nil)).Elem()
}

func (o ManagementGroupDetailsResponseOutput) ToManagementGroupDetailsResponseOutput() ManagementGroupDetailsResponseOutput {
	return o
}

func (o ManagementGroupDetailsResponseOutput) ToManagementGroupDetailsResponseOutputWithContext(ctx context.Context) ManagementGroupDetailsResponseOutput {
	return o
}

func (o ManagementGroupDetailsResponseOutput) ManagementGroupAncestors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) []string { return v.ManagementGroupAncestors }).(pulumi.StringArrayOutput)
}

func (o ManagementGroupDetailsResponseOutput) ManagementGroupAncestorsChain() ManagementGroupPathElementResponseArrayOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) []ManagementGroupPathElementResponse {
		return v.ManagementGroupAncestorsChain
	}).(ManagementGroupPathElementResponseArrayOutput)
}

func (o ManagementGroupDetailsResponseOutput) Parent() ParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *ParentGroupInfoResponse { return v.Parent }).(ParentGroupInfoResponsePtrOutput)
}

func (o ManagementGroupDetailsResponseOutput) Path() ManagementGroupPathElementResponseArrayOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) []ManagementGroupPathElementResponse { return v.Path }).(ManagementGroupPathElementResponseArrayOutput)
}

func (o ManagementGroupDetailsResponseOutput) UpdatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *string { return v.UpdatedBy }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupDetailsResponseOutput) UpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *string { return v.UpdatedTime }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupDetailsResponseOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *float64 { return v.Version }).(pulumi.Float64PtrOutput)
}

type ManagementGroupDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementGroupDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupDetailsResponse)(nil)).Elem()
}

func (o ManagementGroupDetailsResponsePtrOutput) ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput {
	return o
}

func (o ManagementGroupDetailsResponsePtrOutput) ToManagementGroupDetailsResponsePtrOutputWithContext(ctx context.Context) ManagementGroupDetailsResponsePtrOutput {
	return o
}

func (o ManagementGroupDetailsResponsePtrOutput) Elem() ManagementGroupDetailsResponseOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) ManagementGroupDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ManagementGroupDetailsResponse
		return ret
	}).(ManagementGroupDetailsResponseOutput)
}

func (o ManagementGroupDetailsResponsePtrOutput) ManagementGroupAncestors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.ManagementGroupAncestors
	}).(pulumi.StringArrayOutput)
}

func (o ManagementGroupDetailsResponsePtrOutput) ManagementGroupAncestorsChain() ManagementGroupPathElementResponseArrayOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) []ManagementGroupPathElementResponse {
		if v == nil {
			return nil
		}
		return v.ManagementGroupAncestorsChain
	}).(ManagementGroupPathElementResponseArrayOutput)
}

func (o ManagementGroupDetailsResponsePtrOutput) Parent() ParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) *ParentGroupInfoResponse {
		if v == nil {
			return nil
		}
		return v.Parent
	}).(ParentGroupInfoResponsePtrOutput)
}

func (o ManagementGroupDetailsResponsePtrOutput) Path() ManagementGroupPathElementResponseArrayOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) []ManagementGroupPathElementResponse {
		if v == nil {
			return nil
		}
		return v.Path
	}).(ManagementGroupPathElementResponseArrayOutput)
}

func (o ManagementGroupDetailsResponsePtrOutput) UpdatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpdatedBy
	}).(pulumi.StringPtrOutput)
}

func (o ManagementGroupDetailsResponsePtrOutput) UpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpdatedTime
	}).(pulumi.StringPtrOutput)
}

func (o ManagementGroupDetailsResponsePtrOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.Float64PtrOutput)
}

type ManagementGroupPathElementResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Name        *string `pulumi:"name"`
}

type ManagementGroupPathElementResponseOutput struct{ *pulumi.OutputState }

func (ManagementGroupPathElementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupPathElementResponse)(nil)).Elem()
}

func (o ManagementGroupPathElementResponseOutput) ToManagementGroupPathElementResponseOutput() ManagementGroupPathElementResponseOutput {
	return o
}

func (o ManagementGroupPathElementResponseOutput) ToManagementGroupPathElementResponseOutputWithContext(ctx context.Context) ManagementGroupPathElementResponseOutput {
	return o
}

func (o ManagementGroupPathElementResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupPathElementResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupPathElementResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupPathElementResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ManagementGroupPathElementResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementGroupPathElementResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupPathElementResponse)(nil)).Elem()
}

func (o ManagementGroupPathElementResponseArrayOutput) ToManagementGroupPathElementResponseArrayOutput() ManagementGroupPathElementResponseArrayOutput {
	return o
}

func (o ManagementGroupPathElementResponseArrayOutput) ToManagementGroupPathElementResponseArrayOutputWithContext(ctx context.Context) ManagementGroupPathElementResponseArrayOutput {
	return o
}

func (o ManagementGroupPathElementResponseArrayOutput) Index(i pulumi.IntInput) ManagementGroupPathElementResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementGroupPathElementResponse {
		return vs[0].([]ManagementGroupPathElementResponse)[vs[1].(int)]
	}).(ManagementGroupPathElementResponseOutput)
}

type ParentGroupInfoResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Id          *string `pulumi:"id"`
	Name        *string `pulumi:"name"`
}

type ParentGroupInfoResponseOutput struct{ *pulumi.OutputState }

func (ParentGroupInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentGroupInfoResponse)(nil)).Elem()
}

func (o ParentGroupInfoResponseOutput) ToParentGroupInfoResponseOutput() ParentGroupInfoResponseOutput {
	return o
}

func (o ParentGroupInfoResponseOutput) ToParentGroupInfoResponseOutputWithContext(ctx context.Context) ParentGroupInfoResponseOutput {
	return o
}

func (o ParentGroupInfoResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentGroupInfoResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ParentGroupInfoResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentGroupInfoResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ParentGroupInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentGroupInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ParentGroupInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ParentGroupInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParentGroupInfoResponse)(nil)).Elem()
}

func (o ParentGroupInfoResponsePtrOutput) ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput {
	return o
}

func (o ParentGroupInfoResponsePtrOutput) ToParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) ParentGroupInfoResponsePtrOutput {
	return o
}

func (o ParentGroupInfoResponsePtrOutput) Elem() ParentGroupInfoResponseOutput {
	return o.ApplyT(func(v *ParentGroupInfoResponse) ParentGroupInfoResponse {
		if v != nil {
			return *v
		}
		var ret ParentGroupInfoResponse
		return ret
	}).(ParentGroupInfoResponseOutput)
}

func (o ParentGroupInfoResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParentGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ParentGroupInfoResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParentGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ParentGroupInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParentGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CreateManagementGroupDetailsOutput{})
	pulumi.RegisterOutputType(CreateManagementGroupDetailsPtrOutput{})
	pulumi.RegisterOutputType(CreateParentGroupInfoOutput{})
	pulumi.RegisterOutputType(CreateParentGroupInfoPtrOutput{})
	pulumi.RegisterOutputType(DescendantParentGroupInfoResponseOutput{})
	pulumi.RegisterOutputType(DescendantParentGroupInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementGroupChildInfoResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupChildInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupDetailsResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementGroupPathElementResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupPathElementResponseArrayOutput{})
	pulumi.RegisterOutputType(ParentGroupInfoResponseOutput{})
	pulumi.RegisterOutputType(ParentGroupInfoResponsePtrOutput{})
}
