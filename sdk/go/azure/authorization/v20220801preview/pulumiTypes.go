


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyVariableColumn struct {
	ColumnName string `pulumi:"columnName"`
}





type PolicyVariableColumnInput interface {
	pulumi.Input

	ToPolicyVariableColumnOutput() PolicyVariableColumnOutput
	ToPolicyVariableColumnOutputWithContext(context.Context) PolicyVariableColumnOutput
}

type PolicyVariableColumnArgs struct {
	ColumnName pulumi.StringInput `pulumi:"columnName"`
}

func (PolicyVariableColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableColumn)(nil)).Elem()
}

func (i PolicyVariableColumnArgs) ToPolicyVariableColumnOutput() PolicyVariableColumnOutput {
	return i.ToPolicyVariableColumnOutputWithContext(context.Background())
}

func (i PolicyVariableColumnArgs) ToPolicyVariableColumnOutputWithContext(ctx context.Context) PolicyVariableColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVariableColumnOutput)
}





type PolicyVariableColumnArrayInput interface {
	pulumi.Input

	ToPolicyVariableColumnArrayOutput() PolicyVariableColumnArrayOutput
	ToPolicyVariableColumnArrayOutputWithContext(context.Context) PolicyVariableColumnArrayOutput
}

type PolicyVariableColumnArray []PolicyVariableColumnInput

func (PolicyVariableColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableColumn)(nil)).Elem()
}

func (i PolicyVariableColumnArray) ToPolicyVariableColumnArrayOutput() PolicyVariableColumnArrayOutput {
	return i.ToPolicyVariableColumnArrayOutputWithContext(context.Background())
}

func (i PolicyVariableColumnArray) ToPolicyVariableColumnArrayOutputWithContext(ctx context.Context) PolicyVariableColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVariableColumnArrayOutput)
}

type PolicyVariableColumnOutput struct{ *pulumi.OutputState }

func (PolicyVariableColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableColumn)(nil)).Elem()
}

func (o PolicyVariableColumnOutput) ToPolicyVariableColumnOutput() PolicyVariableColumnOutput {
	return o
}

func (o PolicyVariableColumnOutput) ToPolicyVariableColumnOutputWithContext(ctx context.Context) PolicyVariableColumnOutput {
	return o
}

func (o PolicyVariableColumnOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyVariableColumn) string { return v.ColumnName }).(pulumi.StringOutput)
}

type PolicyVariableColumnArrayOutput struct{ *pulumi.OutputState }

func (PolicyVariableColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableColumn)(nil)).Elem()
}

func (o PolicyVariableColumnArrayOutput) ToPolicyVariableColumnArrayOutput() PolicyVariableColumnArrayOutput {
	return o
}

func (o PolicyVariableColumnArrayOutput) ToPolicyVariableColumnArrayOutputWithContext(ctx context.Context) PolicyVariableColumnArrayOutput {
	return o
}

func (o PolicyVariableColumnArrayOutput) Index(i pulumi.IntInput) PolicyVariableColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyVariableColumn {
		return vs[0].([]PolicyVariableColumn)[vs[1].(int)]
	}).(PolicyVariableColumnOutput)
}

type PolicyVariableColumnResponse struct {
	ColumnName string `pulumi:"columnName"`
}

type PolicyVariableColumnResponseOutput struct{ *pulumi.OutputState }

func (PolicyVariableColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableColumnResponse)(nil)).Elem()
}

func (o PolicyVariableColumnResponseOutput) ToPolicyVariableColumnResponseOutput() PolicyVariableColumnResponseOutput {
	return o
}

func (o PolicyVariableColumnResponseOutput) ToPolicyVariableColumnResponseOutputWithContext(ctx context.Context) PolicyVariableColumnResponseOutput {
	return o
}

func (o PolicyVariableColumnResponseOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyVariableColumnResponse) string { return v.ColumnName }).(pulumi.StringOutput)
}

type PolicyVariableColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (PolicyVariableColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableColumnResponse)(nil)).Elem()
}

func (o PolicyVariableColumnResponseArrayOutput) ToPolicyVariableColumnResponseArrayOutput() PolicyVariableColumnResponseArrayOutput {
	return o
}

func (o PolicyVariableColumnResponseArrayOutput) ToPolicyVariableColumnResponseArrayOutputWithContext(ctx context.Context) PolicyVariableColumnResponseArrayOutput {
	return o
}

func (o PolicyVariableColumnResponseArrayOutput) Index(i pulumi.IntInput) PolicyVariableColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyVariableColumnResponse {
		return vs[0].([]PolicyVariableColumnResponse)[vs[1].(int)]
	}).(PolicyVariableColumnResponseOutput)
}

type PolicyVariableValueColumnValue struct {
	ColumnName  string      `pulumi:"columnName"`
	ColumnValue interface{} `pulumi:"columnValue"`
}





type PolicyVariableValueColumnValueInput interface {
	pulumi.Input

	ToPolicyVariableValueColumnValueOutput() PolicyVariableValueColumnValueOutput
	ToPolicyVariableValueColumnValueOutputWithContext(context.Context) PolicyVariableValueColumnValueOutput
}

type PolicyVariableValueColumnValueArgs struct {
	ColumnName  pulumi.StringInput `pulumi:"columnName"`
	ColumnValue pulumi.Input       `pulumi:"columnValue"`
}

func (PolicyVariableValueColumnValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableValueColumnValue)(nil)).Elem()
}

func (i PolicyVariableValueColumnValueArgs) ToPolicyVariableValueColumnValueOutput() PolicyVariableValueColumnValueOutput {
	return i.ToPolicyVariableValueColumnValueOutputWithContext(context.Background())
}

func (i PolicyVariableValueColumnValueArgs) ToPolicyVariableValueColumnValueOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVariableValueColumnValueOutput)
}





type PolicyVariableValueColumnValueArrayInput interface {
	pulumi.Input

	ToPolicyVariableValueColumnValueArrayOutput() PolicyVariableValueColumnValueArrayOutput
	ToPolicyVariableValueColumnValueArrayOutputWithContext(context.Context) PolicyVariableValueColumnValueArrayOutput
}

type PolicyVariableValueColumnValueArray []PolicyVariableValueColumnValueInput

func (PolicyVariableValueColumnValueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableValueColumnValue)(nil)).Elem()
}

func (i PolicyVariableValueColumnValueArray) ToPolicyVariableValueColumnValueArrayOutput() PolicyVariableValueColumnValueArrayOutput {
	return i.ToPolicyVariableValueColumnValueArrayOutputWithContext(context.Background())
}

func (i PolicyVariableValueColumnValueArray) ToPolicyVariableValueColumnValueArrayOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVariableValueColumnValueArrayOutput)
}

type PolicyVariableValueColumnValueOutput struct{ *pulumi.OutputState }

func (PolicyVariableValueColumnValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableValueColumnValue)(nil)).Elem()
}

func (o PolicyVariableValueColumnValueOutput) ToPolicyVariableValueColumnValueOutput() PolicyVariableValueColumnValueOutput {
	return o
}

func (o PolicyVariableValueColumnValueOutput) ToPolicyVariableValueColumnValueOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueOutput {
	return o
}

func (o PolicyVariableValueColumnValueOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyVariableValueColumnValue) string { return v.ColumnName }).(pulumi.StringOutput)
}

func (o PolicyVariableValueColumnValueOutput) ColumnValue() pulumi.AnyOutput {
	return o.ApplyT(func(v PolicyVariableValueColumnValue) interface{} { return v.ColumnValue }).(pulumi.AnyOutput)
}

type PolicyVariableValueColumnValueArrayOutput struct{ *pulumi.OutputState }

func (PolicyVariableValueColumnValueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableValueColumnValue)(nil)).Elem()
}

func (o PolicyVariableValueColumnValueArrayOutput) ToPolicyVariableValueColumnValueArrayOutput() PolicyVariableValueColumnValueArrayOutput {
	return o
}

func (o PolicyVariableValueColumnValueArrayOutput) ToPolicyVariableValueColumnValueArrayOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueArrayOutput {
	return o
}

func (o PolicyVariableValueColumnValueArrayOutput) Index(i pulumi.IntInput) PolicyVariableValueColumnValueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyVariableValueColumnValue {
		return vs[0].([]PolicyVariableValueColumnValue)[vs[1].(int)]
	}).(PolicyVariableValueColumnValueOutput)
}

type PolicyVariableValueColumnValueResponse struct {
	ColumnName  string      `pulumi:"columnName"`
	ColumnValue interface{} `pulumi:"columnValue"`
}

type PolicyVariableValueColumnValueResponseOutput struct{ *pulumi.OutputState }

func (PolicyVariableValueColumnValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableValueColumnValueResponse)(nil)).Elem()
}

func (o PolicyVariableValueColumnValueResponseOutput) ToPolicyVariableValueColumnValueResponseOutput() PolicyVariableValueColumnValueResponseOutput {
	return o
}

func (o PolicyVariableValueColumnValueResponseOutput) ToPolicyVariableValueColumnValueResponseOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueResponseOutput {
	return o
}

func (o PolicyVariableValueColumnValueResponseOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyVariableValueColumnValueResponse) string { return v.ColumnName }).(pulumi.StringOutput)
}

func (o PolicyVariableValueColumnValueResponseOutput) ColumnValue() pulumi.AnyOutput {
	return o.ApplyT(func(v PolicyVariableValueColumnValueResponse) interface{} { return v.ColumnValue }).(pulumi.AnyOutput)
}

type PolicyVariableValueColumnValueResponseArrayOutput struct{ *pulumi.OutputState }

func (PolicyVariableValueColumnValueResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableValueColumnValueResponse)(nil)).Elem()
}

func (o PolicyVariableValueColumnValueResponseArrayOutput) ToPolicyVariableValueColumnValueResponseArrayOutput() PolicyVariableValueColumnValueResponseArrayOutput {
	return o
}

func (o PolicyVariableValueColumnValueResponseArrayOutput) ToPolicyVariableValueColumnValueResponseArrayOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueResponseArrayOutput {
	return o
}

func (o PolicyVariableValueColumnValueResponseArrayOutput) Index(i pulumi.IntInput) PolicyVariableValueColumnValueResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyVariableValueColumnValueResponse {
		return vs[0].([]PolicyVariableValueColumnValueResponse)[vs[1].(int)]
	}).(PolicyVariableValueColumnValueResponseOutput)
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
	pulumi.RegisterOutputType(PolicyVariableColumnOutput{})
	pulumi.RegisterOutputType(PolicyVariableColumnArrayOutput{})
	pulumi.RegisterOutputType(PolicyVariableColumnResponseOutput{})
	pulumi.RegisterOutputType(PolicyVariableColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(PolicyVariableValueColumnValueOutput{})
	pulumi.RegisterOutputType(PolicyVariableValueColumnValueArrayOutput{})
	pulumi.RegisterOutputType(PolicyVariableValueColumnValueResponseOutput{})
	pulumi.RegisterOutputType(PolicyVariableValueColumnValueResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
