


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssignedComponentItem struct {
	Key *string `pulumi:"key"`
}





type AssignedComponentItemInput interface {
	pulumi.Input

	ToAssignedComponentItemOutput() AssignedComponentItemOutput
	ToAssignedComponentItemOutputWithContext(context.Context) AssignedComponentItemOutput
}

type AssignedComponentItemArgs struct {
	Key pulumi.StringPtrInput `pulumi:"key"`
}

func (AssignedComponentItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedComponentItem)(nil)).Elem()
}

func (i AssignedComponentItemArgs) ToAssignedComponentItemOutput() AssignedComponentItemOutput {
	return i.ToAssignedComponentItemOutputWithContext(context.Background())
}

func (i AssignedComponentItemArgs) ToAssignedComponentItemOutputWithContext(ctx context.Context) AssignedComponentItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedComponentItemOutput)
}

func (i AssignedComponentItemArgs) ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput {
	return i.ToAssignedComponentItemPtrOutputWithContext(context.Background())
}

func (i AssignedComponentItemArgs) ToAssignedComponentItemPtrOutputWithContext(ctx context.Context) AssignedComponentItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedComponentItemOutput).ToAssignedComponentItemPtrOutputWithContext(ctx)
}









type AssignedComponentItemPtrInput interface {
	pulumi.Input

	ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput
	ToAssignedComponentItemPtrOutputWithContext(context.Context) AssignedComponentItemPtrOutput
}

type assignedComponentItemPtrType AssignedComponentItemArgs

func AssignedComponentItemPtr(v *AssignedComponentItemArgs) AssignedComponentItemPtrInput {
	return (*assignedComponentItemPtrType)(v)
}

func (*assignedComponentItemPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedComponentItem)(nil)).Elem()
}

func (i *assignedComponentItemPtrType) ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput {
	return i.ToAssignedComponentItemPtrOutputWithContext(context.Background())
}

func (i *assignedComponentItemPtrType) ToAssignedComponentItemPtrOutputWithContext(ctx context.Context) AssignedComponentItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedComponentItemPtrOutput)
}

type AssignedComponentItemOutput struct{ *pulumi.OutputState }

func (AssignedComponentItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedComponentItem)(nil)).Elem()
}

func (o AssignedComponentItemOutput) ToAssignedComponentItemOutput() AssignedComponentItemOutput {
	return o
}

func (o AssignedComponentItemOutput) ToAssignedComponentItemOutputWithContext(ctx context.Context) AssignedComponentItemOutput {
	return o
}

func (o AssignedComponentItemOutput) ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput {
	return o.ToAssignedComponentItemPtrOutputWithContext(context.Background())
}

func (o AssignedComponentItemOutput) ToAssignedComponentItemPtrOutputWithContext(ctx context.Context) AssignedComponentItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignedComponentItem) *AssignedComponentItem {
		return &v
	}).(AssignedComponentItemPtrOutput)
}

func (o AssignedComponentItemOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignedComponentItem) *string { return v.Key }).(pulumi.StringPtrOutput)
}

type AssignedComponentItemPtrOutput struct{ *pulumi.OutputState }

func (AssignedComponentItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedComponentItem)(nil)).Elem()
}

func (o AssignedComponentItemPtrOutput) ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput {
	return o
}

func (o AssignedComponentItemPtrOutput) ToAssignedComponentItemPtrOutputWithContext(ctx context.Context) AssignedComponentItemPtrOutput {
	return o
}

func (o AssignedComponentItemPtrOutput) Elem() AssignedComponentItemOutput {
	return o.ApplyT(func(v *AssignedComponentItem) AssignedComponentItem {
		if v != nil {
			return *v
		}
		var ret AssignedComponentItem
		return ret
	}).(AssignedComponentItemOutput)
}

func (o AssignedComponentItemPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedComponentItem) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

type AssignedComponentItemResponse struct {
	Key *string `pulumi:"key"`
}

type AssignedComponentItemResponseOutput struct{ *pulumi.OutputState }

func (AssignedComponentItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedComponentItemResponse)(nil)).Elem()
}

func (o AssignedComponentItemResponseOutput) ToAssignedComponentItemResponseOutput() AssignedComponentItemResponseOutput {
	return o
}

func (o AssignedComponentItemResponseOutput) ToAssignedComponentItemResponseOutputWithContext(ctx context.Context) AssignedComponentItemResponseOutput {
	return o
}

func (o AssignedComponentItemResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignedComponentItemResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

type AssignedComponentItemResponsePtrOutput struct{ *pulumi.OutputState }

func (AssignedComponentItemResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedComponentItemResponse)(nil)).Elem()
}

func (o AssignedComponentItemResponsePtrOutput) ToAssignedComponentItemResponsePtrOutput() AssignedComponentItemResponsePtrOutput {
	return o
}

func (o AssignedComponentItemResponsePtrOutput) ToAssignedComponentItemResponsePtrOutputWithContext(ctx context.Context) AssignedComponentItemResponsePtrOutput {
	return o
}

func (o AssignedComponentItemResponsePtrOutput) Elem() AssignedComponentItemResponseOutput {
	return o.ApplyT(func(v *AssignedComponentItemResponse) AssignedComponentItemResponse {
		if v != nil {
			return *v
		}
		var ret AssignedComponentItemResponse
		return ret
	}).(AssignedComponentItemResponseOutput)
}

func (o AssignedComponentItemResponsePtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedComponentItemResponse) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

type AssignedStandardItem struct {
	Id *string `pulumi:"id"`
}





type AssignedStandardItemInput interface {
	pulumi.Input

	ToAssignedStandardItemOutput() AssignedStandardItemOutput
	ToAssignedStandardItemOutputWithContext(context.Context) AssignedStandardItemOutput
}

type AssignedStandardItemArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (AssignedStandardItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedStandardItem)(nil)).Elem()
}

func (i AssignedStandardItemArgs) ToAssignedStandardItemOutput() AssignedStandardItemOutput {
	return i.ToAssignedStandardItemOutputWithContext(context.Background())
}

func (i AssignedStandardItemArgs) ToAssignedStandardItemOutputWithContext(ctx context.Context) AssignedStandardItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedStandardItemOutput)
}

func (i AssignedStandardItemArgs) ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput {
	return i.ToAssignedStandardItemPtrOutputWithContext(context.Background())
}

func (i AssignedStandardItemArgs) ToAssignedStandardItemPtrOutputWithContext(ctx context.Context) AssignedStandardItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedStandardItemOutput).ToAssignedStandardItemPtrOutputWithContext(ctx)
}









type AssignedStandardItemPtrInput interface {
	pulumi.Input

	ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput
	ToAssignedStandardItemPtrOutputWithContext(context.Context) AssignedStandardItemPtrOutput
}

type assignedStandardItemPtrType AssignedStandardItemArgs

func AssignedStandardItemPtr(v *AssignedStandardItemArgs) AssignedStandardItemPtrInput {
	return (*assignedStandardItemPtrType)(v)
}

func (*assignedStandardItemPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedStandardItem)(nil)).Elem()
}

func (i *assignedStandardItemPtrType) ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput {
	return i.ToAssignedStandardItemPtrOutputWithContext(context.Background())
}

func (i *assignedStandardItemPtrType) ToAssignedStandardItemPtrOutputWithContext(ctx context.Context) AssignedStandardItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedStandardItemPtrOutput)
}

type AssignedStandardItemOutput struct{ *pulumi.OutputState }

func (AssignedStandardItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedStandardItem)(nil)).Elem()
}

func (o AssignedStandardItemOutput) ToAssignedStandardItemOutput() AssignedStandardItemOutput {
	return o
}

func (o AssignedStandardItemOutput) ToAssignedStandardItemOutputWithContext(ctx context.Context) AssignedStandardItemOutput {
	return o
}

func (o AssignedStandardItemOutput) ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput {
	return o.ToAssignedStandardItemPtrOutputWithContext(context.Background())
}

func (o AssignedStandardItemOutput) ToAssignedStandardItemPtrOutputWithContext(ctx context.Context) AssignedStandardItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignedStandardItem) *AssignedStandardItem {
		return &v
	}).(AssignedStandardItemPtrOutput)
}

func (o AssignedStandardItemOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignedStandardItem) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type AssignedStandardItemPtrOutput struct{ *pulumi.OutputState }

func (AssignedStandardItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedStandardItem)(nil)).Elem()
}

func (o AssignedStandardItemPtrOutput) ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput {
	return o
}

func (o AssignedStandardItemPtrOutput) ToAssignedStandardItemPtrOutputWithContext(ctx context.Context) AssignedStandardItemPtrOutput {
	return o
}

func (o AssignedStandardItemPtrOutput) Elem() AssignedStandardItemOutput {
	return o.ApplyT(func(v *AssignedStandardItem) AssignedStandardItem {
		if v != nil {
			return *v
		}
		var ret AssignedStandardItem
		return ret
	}).(AssignedStandardItemOutput)
}

func (o AssignedStandardItemPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedStandardItem) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type AssignedStandardItemResponse struct {
	Id *string `pulumi:"id"`
}

type AssignedStandardItemResponseOutput struct{ *pulumi.OutputState }

func (AssignedStandardItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedStandardItemResponse)(nil)).Elem()
}

func (o AssignedStandardItemResponseOutput) ToAssignedStandardItemResponseOutput() AssignedStandardItemResponseOutput {
	return o
}

func (o AssignedStandardItemResponseOutput) ToAssignedStandardItemResponseOutputWithContext(ctx context.Context) AssignedStandardItemResponseOutput {
	return o
}

func (o AssignedStandardItemResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignedStandardItemResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type AssignedStandardItemResponsePtrOutput struct{ *pulumi.OutputState }

func (AssignedStandardItemResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedStandardItemResponse)(nil)).Elem()
}

func (o AssignedStandardItemResponsePtrOutput) ToAssignedStandardItemResponsePtrOutput() AssignedStandardItemResponsePtrOutput {
	return o
}

func (o AssignedStandardItemResponsePtrOutput) ToAssignedStandardItemResponsePtrOutputWithContext(ctx context.Context) AssignedStandardItemResponsePtrOutput {
	return o
}

func (o AssignedStandardItemResponsePtrOutput) Elem() AssignedStandardItemResponseOutput {
	return o.ApplyT(func(v *AssignedStandardItemResponse) AssignedStandardItemResponse {
		if v != nil {
			return *v
		}
		var ret AssignedStandardItemResponse
		return ret
	}).(AssignedStandardItemResponseOutput)
}

func (o AssignedStandardItemResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedStandardItemResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type AssignmentPropertiesAdditionalData struct {
	ExemptionCategory *string `pulumi:"exemptionCategory"`
}





type AssignmentPropertiesAdditionalDataInput interface {
	pulumi.Input

	ToAssignmentPropertiesAdditionalDataOutput() AssignmentPropertiesAdditionalDataOutput
	ToAssignmentPropertiesAdditionalDataOutputWithContext(context.Context) AssignmentPropertiesAdditionalDataOutput
}

type AssignmentPropertiesAdditionalDataArgs struct {
	ExemptionCategory pulumi.StringPtrInput `pulumi:"exemptionCategory"`
}

func (AssignmentPropertiesAdditionalDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPropertiesAdditionalData)(nil)).Elem()
}

func (i AssignmentPropertiesAdditionalDataArgs) ToAssignmentPropertiesAdditionalDataOutput() AssignmentPropertiesAdditionalDataOutput {
	return i.ToAssignmentPropertiesAdditionalDataOutputWithContext(context.Background())
}

func (i AssignmentPropertiesAdditionalDataArgs) ToAssignmentPropertiesAdditionalDataOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPropertiesAdditionalDataOutput)
}

func (i AssignmentPropertiesAdditionalDataArgs) ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput {
	return i.ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(context.Background())
}

func (i AssignmentPropertiesAdditionalDataArgs) ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPropertiesAdditionalDataOutput).ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx)
}









type AssignmentPropertiesAdditionalDataPtrInput interface {
	pulumi.Input

	ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput
	ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(context.Context) AssignmentPropertiesAdditionalDataPtrOutput
}

type assignmentPropertiesAdditionalDataPtrType AssignmentPropertiesAdditionalDataArgs

func AssignmentPropertiesAdditionalDataPtr(v *AssignmentPropertiesAdditionalDataArgs) AssignmentPropertiesAdditionalDataPtrInput {
	return (*assignmentPropertiesAdditionalDataPtrType)(v)
}

func (*assignmentPropertiesAdditionalDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentPropertiesAdditionalData)(nil)).Elem()
}

func (i *assignmentPropertiesAdditionalDataPtrType) ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput {
	return i.ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(context.Background())
}

func (i *assignmentPropertiesAdditionalDataPtrType) ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPropertiesAdditionalDataPtrOutput)
}

type AssignmentPropertiesAdditionalDataOutput struct{ *pulumi.OutputState }

func (AssignmentPropertiesAdditionalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPropertiesAdditionalData)(nil)).Elem()
}

func (o AssignmentPropertiesAdditionalDataOutput) ToAssignmentPropertiesAdditionalDataOutput() AssignmentPropertiesAdditionalDataOutput {
	return o
}

func (o AssignmentPropertiesAdditionalDataOutput) ToAssignmentPropertiesAdditionalDataOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataOutput {
	return o
}

func (o AssignmentPropertiesAdditionalDataOutput) ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput {
	return o.ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(context.Background())
}

func (o AssignmentPropertiesAdditionalDataOutput) ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignmentPropertiesAdditionalData) *AssignmentPropertiesAdditionalData {
		return &v
	}).(AssignmentPropertiesAdditionalDataPtrOutput)
}

func (o AssignmentPropertiesAdditionalDataOutput) ExemptionCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignmentPropertiesAdditionalData) *string { return v.ExemptionCategory }).(pulumi.StringPtrOutput)
}

type AssignmentPropertiesAdditionalDataPtrOutput struct{ *pulumi.OutputState }

func (AssignmentPropertiesAdditionalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentPropertiesAdditionalData)(nil)).Elem()
}

func (o AssignmentPropertiesAdditionalDataPtrOutput) ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput {
	return o
}

func (o AssignmentPropertiesAdditionalDataPtrOutput) ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataPtrOutput {
	return o
}

func (o AssignmentPropertiesAdditionalDataPtrOutput) Elem() AssignmentPropertiesAdditionalDataOutput {
	return o.ApplyT(func(v *AssignmentPropertiesAdditionalData) AssignmentPropertiesAdditionalData {
		if v != nil {
			return *v
		}
		var ret AssignmentPropertiesAdditionalData
		return ret
	}).(AssignmentPropertiesAdditionalDataOutput)
}

func (o AssignmentPropertiesAdditionalDataPtrOutput) ExemptionCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentPropertiesAdditionalData) *string {
		if v == nil {
			return nil
		}
		return v.ExemptionCategory
	}).(pulumi.StringPtrOutput)
}

type AssignmentPropertiesResponseAdditionalData struct {
	ExemptionCategory *string `pulumi:"exemptionCategory"`
}

type AssignmentPropertiesResponseAdditionalDataOutput struct{ *pulumi.OutputState }

func (AssignmentPropertiesResponseAdditionalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPropertiesResponseAdditionalData)(nil)).Elem()
}

func (o AssignmentPropertiesResponseAdditionalDataOutput) ToAssignmentPropertiesResponseAdditionalDataOutput() AssignmentPropertiesResponseAdditionalDataOutput {
	return o
}

func (o AssignmentPropertiesResponseAdditionalDataOutput) ToAssignmentPropertiesResponseAdditionalDataOutputWithContext(ctx context.Context) AssignmentPropertiesResponseAdditionalDataOutput {
	return o
}

func (o AssignmentPropertiesResponseAdditionalDataOutput) ExemptionCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignmentPropertiesResponseAdditionalData) *string { return v.ExemptionCategory }).(pulumi.StringPtrOutput)
}

type AssignmentPropertiesResponseAdditionalDataPtrOutput struct{ *pulumi.OutputState }

func (AssignmentPropertiesResponseAdditionalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentPropertiesResponseAdditionalData)(nil)).Elem()
}

func (o AssignmentPropertiesResponseAdditionalDataPtrOutput) ToAssignmentPropertiesResponseAdditionalDataPtrOutput() AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return o
}

func (o AssignmentPropertiesResponseAdditionalDataPtrOutput) ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return o
}

func (o AssignmentPropertiesResponseAdditionalDataPtrOutput) Elem() AssignmentPropertiesResponseAdditionalDataOutput {
	return o.ApplyT(func(v *AssignmentPropertiesResponseAdditionalData) AssignmentPropertiesResponseAdditionalData {
		if v != nil {
			return *v
		}
		var ret AssignmentPropertiesResponseAdditionalData
		return ret
	}).(AssignmentPropertiesResponseAdditionalDataOutput)
}

func (o AssignmentPropertiesResponseAdditionalDataPtrOutput) ExemptionCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentPropertiesResponseAdditionalData) *string {
		if v == nil {
			return nil
		}
		return v.ExemptionCategory
	}).(pulumi.StringPtrOutput)
}

type StandardComponentProperties struct {
	Key *string `pulumi:"key"`
}





type StandardComponentPropertiesInput interface {
	pulumi.Input

	ToStandardComponentPropertiesOutput() StandardComponentPropertiesOutput
	ToStandardComponentPropertiesOutputWithContext(context.Context) StandardComponentPropertiesOutput
}

type StandardComponentPropertiesArgs struct {
	Key pulumi.StringPtrInput `pulumi:"key"`
}

func (StandardComponentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardComponentProperties)(nil)).Elem()
}

func (i StandardComponentPropertiesArgs) ToStandardComponentPropertiesOutput() StandardComponentPropertiesOutput {
	return i.ToStandardComponentPropertiesOutputWithContext(context.Background())
}

func (i StandardComponentPropertiesArgs) ToStandardComponentPropertiesOutputWithContext(ctx context.Context) StandardComponentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardComponentPropertiesOutput)
}





type StandardComponentPropertiesArrayInput interface {
	pulumi.Input

	ToStandardComponentPropertiesArrayOutput() StandardComponentPropertiesArrayOutput
	ToStandardComponentPropertiesArrayOutputWithContext(context.Context) StandardComponentPropertiesArrayOutput
}

type StandardComponentPropertiesArray []StandardComponentPropertiesInput

func (StandardComponentPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardComponentProperties)(nil)).Elem()
}

func (i StandardComponentPropertiesArray) ToStandardComponentPropertiesArrayOutput() StandardComponentPropertiesArrayOutput {
	return i.ToStandardComponentPropertiesArrayOutputWithContext(context.Background())
}

func (i StandardComponentPropertiesArray) ToStandardComponentPropertiesArrayOutputWithContext(ctx context.Context) StandardComponentPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardComponentPropertiesArrayOutput)
}

type StandardComponentPropertiesOutput struct{ *pulumi.OutputState }

func (StandardComponentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardComponentProperties)(nil)).Elem()
}

func (o StandardComponentPropertiesOutput) ToStandardComponentPropertiesOutput() StandardComponentPropertiesOutput {
	return o
}

func (o StandardComponentPropertiesOutput) ToStandardComponentPropertiesOutputWithContext(ctx context.Context) StandardComponentPropertiesOutput {
	return o
}

func (o StandardComponentPropertiesOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StandardComponentProperties) *string { return v.Key }).(pulumi.StringPtrOutput)
}

type StandardComponentPropertiesArrayOutput struct{ *pulumi.OutputState }

func (StandardComponentPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardComponentProperties)(nil)).Elem()
}

func (o StandardComponentPropertiesArrayOutput) ToStandardComponentPropertiesArrayOutput() StandardComponentPropertiesArrayOutput {
	return o
}

func (o StandardComponentPropertiesArrayOutput) ToStandardComponentPropertiesArrayOutputWithContext(ctx context.Context) StandardComponentPropertiesArrayOutput {
	return o
}

func (o StandardComponentPropertiesArrayOutput) Index(i pulumi.IntInput) StandardComponentPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StandardComponentProperties {
		return vs[0].([]StandardComponentProperties)[vs[1].(int)]
	}).(StandardComponentPropertiesOutput)
}

type StandardComponentPropertiesResponse struct {
	Key *string `pulumi:"key"`
}

type StandardComponentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StandardComponentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardComponentPropertiesResponse)(nil)).Elem()
}

func (o StandardComponentPropertiesResponseOutput) ToStandardComponentPropertiesResponseOutput() StandardComponentPropertiesResponseOutput {
	return o
}

func (o StandardComponentPropertiesResponseOutput) ToStandardComponentPropertiesResponseOutputWithContext(ctx context.Context) StandardComponentPropertiesResponseOutput {
	return o
}

func (o StandardComponentPropertiesResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StandardComponentPropertiesResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

type StandardComponentPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (StandardComponentPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardComponentPropertiesResponse)(nil)).Elem()
}

func (o StandardComponentPropertiesResponseArrayOutput) ToStandardComponentPropertiesResponseArrayOutput() StandardComponentPropertiesResponseArrayOutput {
	return o
}

func (o StandardComponentPropertiesResponseArrayOutput) ToStandardComponentPropertiesResponseArrayOutputWithContext(ctx context.Context) StandardComponentPropertiesResponseArrayOutput {
	return o
}

func (o StandardComponentPropertiesResponseArrayOutput) Index(i pulumi.IntInput) StandardComponentPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StandardComponentPropertiesResponse {
		return vs[0].([]StandardComponentPropertiesResponse)[vs[1].(int)]
	}).(StandardComponentPropertiesResponseOutput)
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

func init() {
	pulumi.RegisterOutputType(AssignedComponentItemOutput{})
	pulumi.RegisterOutputType(AssignedComponentItemPtrOutput{})
	pulumi.RegisterOutputType(AssignedComponentItemResponseOutput{})
	pulumi.RegisterOutputType(AssignedComponentItemResponsePtrOutput{})
	pulumi.RegisterOutputType(AssignedStandardItemOutput{})
	pulumi.RegisterOutputType(AssignedStandardItemPtrOutput{})
	pulumi.RegisterOutputType(AssignedStandardItemResponseOutput{})
	pulumi.RegisterOutputType(AssignedStandardItemResponsePtrOutput{})
	pulumi.RegisterOutputType(AssignmentPropertiesAdditionalDataOutput{})
	pulumi.RegisterOutputType(AssignmentPropertiesAdditionalDataPtrOutput{})
	pulumi.RegisterOutputType(AssignmentPropertiesResponseAdditionalDataOutput{})
	pulumi.RegisterOutputType(AssignmentPropertiesResponseAdditionalDataPtrOutput{})
	pulumi.RegisterOutputType(StandardComponentPropertiesOutput{})
	pulumi.RegisterOutputType(StandardComponentPropertiesArrayOutput{})
	pulumi.RegisterOutputType(StandardComponentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StandardComponentPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
