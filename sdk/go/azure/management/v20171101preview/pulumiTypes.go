


package v20171101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementGroupChildInfoResponse struct {
	ChildId     *string                            `pulumi:"childId"`
	ChildType   *string                            `pulumi:"childType"`
	Children    []ManagementGroupChildInfoResponse `pulumi:"children"`
	DisplayName *string                            `pulumi:"displayName"`
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

func (o ManagementGroupChildInfoResponseOutput) ChildId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.ChildId }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupChildInfoResponseOutput) ChildType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.ChildType }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupChildInfoResponseOutput) Children() ManagementGroupChildInfoResponseArrayOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) []ManagementGroupChildInfoResponse { return v.Children }).(ManagementGroupChildInfoResponseArrayOutput)
}

func (o ManagementGroupChildInfoResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
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
	Parent      *ParentGroupInfoResponse `pulumi:"parent"`
	UpdatedBy   *string                  `pulumi:"updatedBy"`
	UpdatedTime *string                  `pulumi:"updatedTime"`
	Version     *float64                 `pulumi:"version"`
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

func (o ManagementGroupDetailsResponseOutput) Parent() ParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *ParentGroupInfoResponse { return v.Parent }).(ParentGroupInfoResponsePtrOutput)
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

func (o ManagementGroupDetailsResponsePtrOutput) Parent() ParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) *ParentGroupInfoResponse {
		if v == nil {
			return nil
		}
		return v.Parent
	}).(ParentGroupInfoResponsePtrOutput)
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

type ParentGroupInfoResponse struct {
	DisplayName *string `pulumi:"displayName"`
	ParentId    *string `pulumi:"parentId"`
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

func (o ParentGroupInfoResponseOutput) ParentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentGroupInfoResponse) *string { return v.ParentId }).(pulumi.StringPtrOutput)
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

func (o ParentGroupInfoResponsePtrOutput) ParentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParentGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ParentId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupChildInfoResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupChildInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupDetailsResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ParentGroupInfoResponseOutput{})
	pulumi.RegisterOutputType(ParentGroupInfoResponsePtrOutput{})
}
