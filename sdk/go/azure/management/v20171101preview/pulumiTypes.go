


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





type ManagementGroupChildInfoResponseInput interface {
	pulumi.Input

	ToManagementGroupChildInfoResponseOutput() ManagementGroupChildInfoResponseOutput
	ToManagementGroupChildInfoResponseOutputWithContext(context.Context) ManagementGroupChildInfoResponseOutput
}

type ManagementGroupChildInfoResponseArgs struct {
	ChildId     pulumi.StringPtrInput                      `pulumi:"childId"`
	ChildType   pulumi.StringPtrInput                      `pulumi:"childType"`
	Children    ManagementGroupChildInfoResponseArrayInput `pulumi:"children"`
	DisplayName pulumi.StringPtrInput                      `pulumi:"displayName"`
}

func (ManagementGroupChildInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupChildInfoResponse)(nil)).Elem()
}

func (i ManagementGroupChildInfoResponseArgs) ToManagementGroupChildInfoResponseOutput() ManagementGroupChildInfoResponseOutput {
	return i.ToManagementGroupChildInfoResponseOutputWithContext(context.Background())
}

func (i ManagementGroupChildInfoResponseArgs) ToManagementGroupChildInfoResponseOutputWithContext(ctx context.Context) ManagementGroupChildInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupChildInfoResponseOutput)
}





type ManagementGroupChildInfoResponseArrayInput interface {
	pulumi.Input

	ToManagementGroupChildInfoResponseArrayOutput() ManagementGroupChildInfoResponseArrayOutput
	ToManagementGroupChildInfoResponseArrayOutputWithContext(context.Context) ManagementGroupChildInfoResponseArrayOutput
}

type ManagementGroupChildInfoResponseArray []ManagementGroupChildInfoResponseInput

func (ManagementGroupChildInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupChildInfoResponse)(nil)).Elem()
}

func (i ManagementGroupChildInfoResponseArray) ToManagementGroupChildInfoResponseArrayOutput() ManagementGroupChildInfoResponseArrayOutput {
	return i.ToManagementGroupChildInfoResponseArrayOutputWithContext(context.Background())
}

func (i ManagementGroupChildInfoResponseArray) ToManagementGroupChildInfoResponseArrayOutputWithContext(ctx context.Context) ManagementGroupChildInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupChildInfoResponseArrayOutput)
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





type ManagementGroupDetailsResponseInput interface {
	pulumi.Input

	ToManagementGroupDetailsResponseOutput() ManagementGroupDetailsResponseOutput
	ToManagementGroupDetailsResponseOutputWithContext(context.Context) ManagementGroupDetailsResponseOutput
}

type ManagementGroupDetailsResponseArgs struct {
	Parent      ParentGroupInfoResponsePtrInput `pulumi:"parent"`
	UpdatedBy   pulumi.StringPtrInput           `pulumi:"updatedBy"`
	UpdatedTime pulumi.StringPtrInput           `pulumi:"updatedTime"`
	Version     pulumi.Float64PtrInput          `pulumi:"version"`
}

func (ManagementGroupDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupDetailsResponse)(nil)).Elem()
}

func (i ManagementGroupDetailsResponseArgs) ToManagementGroupDetailsResponseOutput() ManagementGroupDetailsResponseOutput {
	return i.ToManagementGroupDetailsResponseOutputWithContext(context.Background())
}

func (i ManagementGroupDetailsResponseArgs) ToManagementGroupDetailsResponseOutputWithContext(ctx context.Context) ManagementGroupDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupDetailsResponseOutput)
}

func (i ManagementGroupDetailsResponseArgs) ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput {
	return i.ToManagementGroupDetailsResponsePtrOutputWithContext(context.Background())
}

func (i ManagementGroupDetailsResponseArgs) ToManagementGroupDetailsResponsePtrOutputWithContext(ctx context.Context) ManagementGroupDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupDetailsResponseOutput).ToManagementGroupDetailsResponsePtrOutputWithContext(ctx)
}









type ManagementGroupDetailsResponsePtrInput interface {
	pulumi.Input

	ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput
	ToManagementGroupDetailsResponsePtrOutputWithContext(context.Context) ManagementGroupDetailsResponsePtrOutput
}

type managementGroupDetailsResponsePtrType ManagementGroupDetailsResponseArgs

func ManagementGroupDetailsResponsePtr(v *ManagementGroupDetailsResponseArgs) ManagementGroupDetailsResponsePtrInput {
	return (*managementGroupDetailsResponsePtrType)(v)
}

func (*managementGroupDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupDetailsResponse)(nil)).Elem()
}

func (i *managementGroupDetailsResponsePtrType) ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput {
	return i.ToManagementGroupDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *managementGroupDetailsResponsePtrType) ToManagementGroupDetailsResponsePtrOutputWithContext(ctx context.Context) ManagementGroupDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupDetailsResponsePtrOutput)
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

func (o ManagementGroupDetailsResponseOutput) ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput {
	return o.ToManagementGroupDetailsResponsePtrOutputWithContext(context.Background())
}

func (o ManagementGroupDetailsResponseOutput) ToManagementGroupDetailsResponsePtrOutputWithContext(ctx context.Context) ManagementGroupDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementGroupDetailsResponse) *ManagementGroupDetailsResponse {
		return &v
	}).(ManagementGroupDetailsResponsePtrOutput)
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





type ParentGroupInfoResponseInput interface {
	pulumi.Input

	ToParentGroupInfoResponseOutput() ParentGroupInfoResponseOutput
	ToParentGroupInfoResponseOutputWithContext(context.Context) ParentGroupInfoResponseOutput
}

type ParentGroupInfoResponseArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	ParentId    pulumi.StringPtrInput `pulumi:"parentId"`
}

func (ParentGroupInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentGroupInfoResponse)(nil)).Elem()
}

func (i ParentGroupInfoResponseArgs) ToParentGroupInfoResponseOutput() ParentGroupInfoResponseOutput {
	return i.ToParentGroupInfoResponseOutputWithContext(context.Background())
}

func (i ParentGroupInfoResponseArgs) ToParentGroupInfoResponseOutputWithContext(ctx context.Context) ParentGroupInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentGroupInfoResponseOutput)
}

func (i ParentGroupInfoResponseArgs) ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput {
	return i.ToParentGroupInfoResponsePtrOutputWithContext(context.Background())
}

func (i ParentGroupInfoResponseArgs) ToParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) ParentGroupInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentGroupInfoResponseOutput).ToParentGroupInfoResponsePtrOutputWithContext(ctx)
}









type ParentGroupInfoResponsePtrInput interface {
	pulumi.Input

	ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput
	ToParentGroupInfoResponsePtrOutputWithContext(context.Context) ParentGroupInfoResponsePtrOutput
}

type parentGroupInfoResponsePtrType ParentGroupInfoResponseArgs

func ParentGroupInfoResponsePtr(v *ParentGroupInfoResponseArgs) ParentGroupInfoResponsePtrInput {
	return (*parentGroupInfoResponsePtrType)(v)
}

func (*parentGroupInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParentGroupInfoResponse)(nil)).Elem()
}

func (i *parentGroupInfoResponsePtrType) ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput {
	return i.ToParentGroupInfoResponsePtrOutputWithContext(context.Background())
}

func (i *parentGroupInfoResponsePtrType) ToParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) ParentGroupInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentGroupInfoResponsePtrOutput)
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

func (o ParentGroupInfoResponseOutput) ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput {
	return o.ToParentGroupInfoResponsePtrOutputWithContext(context.Background())
}

func (o ParentGroupInfoResponseOutput) ToParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) ParentGroupInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParentGroupInfoResponse) *ParentGroupInfoResponse {
		return &v
	}).(ParentGroupInfoResponsePtrOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementGroupChildInfoResponseInput)(nil)).Elem(), ManagementGroupChildInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementGroupChildInfoResponseArrayInput)(nil)).Elem(), ManagementGroupChildInfoResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementGroupDetailsResponseInput)(nil)).Elem(), ManagementGroupDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagementGroupDetailsResponsePtrInput)(nil)).Elem(), ManagementGroupDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParentGroupInfoResponseInput)(nil)).Elem(), ParentGroupInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParentGroupInfoResponsePtrInput)(nil)).Elem(), ParentGroupInfoResponseArgs{})
	pulumi.RegisterOutputType(ManagementGroupChildInfoResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupChildInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupDetailsResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ParentGroupInfoResponseOutput{})
	pulumi.RegisterOutputType(ParentGroupInfoResponsePtrOutput{})
}
