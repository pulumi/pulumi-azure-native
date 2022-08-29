


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BudgetComparisonExpression struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type BudgetComparisonExpressionInput interface {
	pulumi.Input

	ToBudgetComparisonExpressionOutput() BudgetComparisonExpressionOutput
	ToBudgetComparisonExpressionOutputWithContext(context.Context) BudgetComparisonExpressionOutput
}

type BudgetComparisonExpressionArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (BudgetComparisonExpressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetComparisonExpression)(nil)).Elem()
}

func (i BudgetComparisonExpressionArgs) ToBudgetComparisonExpressionOutput() BudgetComparisonExpressionOutput {
	return i.ToBudgetComparisonExpressionOutputWithContext(context.Background())
}

func (i BudgetComparisonExpressionArgs) ToBudgetComparisonExpressionOutputWithContext(ctx context.Context) BudgetComparisonExpressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetComparisonExpressionOutput)
}

func (i BudgetComparisonExpressionArgs) ToBudgetComparisonExpressionPtrOutput() BudgetComparisonExpressionPtrOutput {
	return i.ToBudgetComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i BudgetComparisonExpressionArgs) ToBudgetComparisonExpressionPtrOutputWithContext(ctx context.Context) BudgetComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetComparisonExpressionOutput).ToBudgetComparisonExpressionPtrOutputWithContext(ctx)
}









type BudgetComparisonExpressionPtrInput interface {
	pulumi.Input

	ToBudgetComparisonExpressionPtrOutput() BudgetComparisonExpressionPtrOutput
	ToBudgetComparisonExpressionPtrOutputWithContext(context.Context) BudgetComparisonExpressionPtrOutput
}

type budgetComparisonExpressionPtrType BudgetComparisonExpressionArgs

func BudgetComparisonExpressionPtr(v *BudgetComparisonExpressionArgs) BudgetComparisonExpressionPtrInput {
	return (*budgetComparisonExpressionPtrType)(v)
}

func (*budgetComparisonExpressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetComparisonExpression)(nil)).Elem()
}

func (i *budgetComparisonExpressionPtrType) ToBudgetComparisonExpressionPtrOutput() BudgetComparisonExpressionPtrOutput {
	return i.ToBudgetComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i *budgetComparisonExpressionPtrType) ToBudgetComparisonExpressionPtrOutputWithContext(ctx context.Context) BudgetComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetComparisonExpressionPtrOutput)
}

type BudgetComparisonExpressionOutput struct{ *pulumi.OutputState }

func (BudgetComparisonExpressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetComparisonExpression)(nil)).Elem()
}

func (o BudgetComparisonExpressionOutput) ToBudgetComparisonExpressionOutput() BudgetComparisonExpressionOutput {
	return o
}

func (o BudgetComparisonExpressionOutput) ToBudgetComparisonExpressionOutputWithContext(ctx context.Context) BudgetComparisonExpressionOutput {
	return o
}

func (o BudgetComparisonExpressionOutput) ToBudgetComparisonExpressionPtrOutput() BudgetComparisonExpressionPtrOutput {
	return o.ToBudgetComparisonExpressionPtrOutputWithContext(context.Background())
}

func (o BudgetComparisonExpressionOutput) ToBudgetComparisonExpressionPtrOutputWithContext(ctx context.Context) BudgetComparisonExpressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BudgetComparisonExpression) *BudgetComparisonExpression {
		return &v
	}).(BudgetComparisonExpressionPtrOutput)
}

func (o BudgetComparisonExpressionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BudgetComparisonExpression) string { return v.Name }).(pulumi.StringOutput)
}

func (o BudgetComparisonExpressionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v BudgetComparisonExpression) string { return v.Operator }).(pulumi.StringOutput)
}

func (o BudgetComparisonExpressionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BudgetComparisonExpression) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type BudgetComparisonExpressionPtrOutput struct{ *pulumi.OutputState }

func (BudgetComparisonExpressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetComparisonExpression)(nil)).Elem()
}

func (o BudgetComparisonExpressionPtrOutput) ToBudgetComparisonExpressionPtrOutput() BudgetComparisonExpressionPtrOutput {
	return o
}

func (o BudgetComparisonExpressionPtrOutput) ToBudgetComparisonExpressionPtrOutputWithContext(ctx context.Context) BudgetComparisonExpressionPtrOutput {
	return o
}

func (o BudgetComparisonExpressionPtrOutput) Elem() BudgetComparisonExpressionOutput {
	return o.ApplyT(func(v *BudgetComparisonExpression) BudgetComparisonExpression {
		if v != nil {
			return *v
		}
		var ret BudgetComparisonExpression
		return ret
	}).(BudgetComparisonExpressionOutput)
}

func (o BudgetComparisonExpressionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BudgetComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o BudgetComparisonExpressionPtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BudgetComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o BudgetComparisonExpressionPtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BudgetComparisonExpression) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type BudgetComparisonExpressionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type BudgetComparisonExpressionResponseOutput struct{ *pulumi.OutputState }

func (BudgetComparisonExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetComparisonExpressionResponse)(nil)).Elem()
}

func (o BudgetComparisonExpressionResponseOutput) ToBudgetComparisonExpressionResponseOutput() BudgetComparisonExpressionResponseOutput {
	return o
}

func (o BudgetComparisonExpressionResponseOutput) ToBudgetComparisonExpressionResponseOutputWithContext(ctx context.Context) BudgetComparisonExpressionResponseOutput {
	return o
}

func (o BudgetComparisonExpressionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BudgetComparisonExpressionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o BudgetComparisonExpressionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v BudgetComparisonExpressionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o BudgetComparisonExpressionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BudgetComparisonExpressionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type BudgetComparisonExpressionResponsePtrOutput struct{ *pulumi.OutputState }

func (BudgetComparisonExpressionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetComparisonExpressionResponse)(nil)).Elem()
}

func (o BudgetComparisonExpressionResponsePtrOutput) ToBudgetComparisonExpressionResponsePtrOutput() BudgetComparisonExpressionResponsePtrOutput {
	return o
}

func (o BudgetComparisonExpressionResponsePtrOutput) ToBudgetComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) BudgetComparisonExpressionResponsePtrOutput {
	return o
}

func (o BudgetComparisonExpressionResponsePtrOutput) Elem() BudgetComparisonExpressionResponseOutput {
	return o.ApplyT(func(v *BudgetComparisonExpressionResponse) BudgetComparisonExpressionResponse {
		if v != nil {
			return *v
		}
		var ret BudgetComparisonExpressionResponse
		return ret
	}).(BudgetComparisonExpressionResponseOutput)
}

func (o BudgetComparisonExpressionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BudgetComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o BudgetComparisonExpressionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BudgetComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o BudgetComparisonExpressionResponsePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BudgetComparisonExpressionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type BudgetFilter struct {
	And        []BudgetFilterProperties    `pulumi:"and"`
	Dimensions *BudgetComparisonExpression `pulumi:"dimensions"`
	Not        *BudgetFilterProperties     `pulumi:"not"`
	Tags       *BudgetComparisonExpression `pulumi:"tags"`
}





type BudgetFilterInput interface {
	pulumi.Input

	ToBudgetFilterOutput() BudgetFilterOutput
	ToBudgetFilterOutputWithContext(context.Context) BudgetFilterOutput
}

type BudgetFilterArgs struct {
	And        BudgetFilterPropertiesArrayInput   `pulumi:"and"`
	Dimensions BudgetComparisonExpressionPtrInput `pulumi:"dimensions"`
	Not        BudgetFilterPropertiesPtrInput     `pulumi:"not"`
	Tags       BudgetComparisonExpressionPtrInput `pulumi:"tags"`
}

func (BudgetFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetFilter)(nil)).Elem()
}

func (i BudgetFilterArgs) ToBudgetFilterOutput() BudgetFilterOutput {
	return i.ToBudgetFilterOutputWithContext(context.Background())
}

func (i BudgetFilterArgs) ToBudgetFilterOutputWithContext(ctx context.Context) BudgetFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetFilterOutput)
}

func (i BudgetFilterArgs) ToBudgetFilterPtrOutput() BudgetFilterPtrOutput {
	return i.ToBudgetFilterPtrOutputWithContext(context.Background())
}

func (i BudgetFilterArgs) ToBudgetFilterPtrOutputWithContext(ctx context.Context) BudgetFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetFilterOutput).ToBudgetFilterPtrOutputWithContext(ctx)
}









type BudgetFilterPtrInput interface {
	pulumi.Input

	ToBudgetFilterPtrOutput() BudgetFilterPtrOutput
	ToBudgetFilterPtrOutputWithContext(context.Context) BudgetFilterPtrOutput
}

type budgetFilterPtrType BudgetFilterArgs

func BudgetFilterPtr(v *BudgetFilterArgs) BudgetFilterPtrInput {
	return (*budgetFilterPtrType)(v)
}

func (*budgetFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetFilter)(nil)).Elem()
}

func (i *budgetFilterPtrType) ToBudgetFilterPtrOutput() BudgetFilterPtrOutput {
	return i.ToBudgetFilterPtrOutputWithContext(context.Background())
}

func (i *budgetFilterPtrType) ToBudgetFilterPtrOutputWithContext(ctx context.Context) BudgetFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetFilterPtrOutput)
}

type BudgetFilterOutput struct{ *pulumi.OutputState }

func (BudgetFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetFilter)(nil)).Elem()
}

func (o BudgetFilterOutput) ToBudgetFilterOutput() BudgetFilterOutput {
	return o
}

func (o BudgetFilterOutput) ToBudgetFilterOutputWithContext(ctx context.Context) BudgetFilterOutput {
	return o
}

func (o BudgetFilterOutput) ToBudgetFilterPtrOutput() BudgetFilterPtrOutput {
	return o.ToBudgetFilterPtrOutputWithContext(context.Background())
}

func (o BudgetFilterOutput) ToBudgetFilterPtrOutputWithContext(ctx context.Context) BudgetFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BudgetFilter) *BudgetFilter {
		return &v
	}).(BudgetFilterPtrOutput)
}

func (o BudgetFilterOutput) And() BudgetFilterPropertiesArrayOutput {
	return o.ApplyT(func(v BudgetFilter) []BudgetFilterProperties { return v.And }).(BudgetFilterPropertiesArrayOutput)
}

func (o BudgetFilterOutput) Dimensions() BudgetComparisonExpressionPtrOutput {
	return o.ApplyT(func(v BudgetFilter) *BudgetComparisonExpression { return v.Dimensions }).(BudgetComparisonExpressionPtrOutput)
}

func (o BudgetFilterOutput) Not() BudgetFilterPropertiesPtrOutput {
	return o.ApplyT(func(v BudgetFilter) *BudgetFilterProperties { return v.Not }).(BudgetFilterPropertiesPtrOutput)
}

func (o BudgetFilterOutput) Tags() BudgetComparisonExpressionPtrOutput {
	return o.ApplyT(func(v BudgetFilter) *BudgetComparisonExpression { return v.Tags }).(BudgetComparisonExpressionPtrOutput)
}

type BudgetFilterPtrOutput struct{ *pulumi.OutputState }

func (BudgetFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetFilter)(nil)).Elem()
}

func (o BudgetFilterPtrOutput) ToBudgetFilterPtrOutput() BudgetFilterPtrOutput {
	return o
}

func (o BudgetFilterPtrOutput) ToBudgetFilterPtrOutputWithContext(ctx context.Context) BudgetFilterPtrOutput {
	return o
}

func (o BudgetFilterPtrOutput) Elem() BudgetFilterOutput {
	return o.ApplyT(func(v *BudgetFilter) BudgetFilter {
		if v != nil {
			return *v
		}
		var ret BudgetFilter
		return ret
	}).(BudgetFilterOutput)
}

func (o BudgetFilterPtrOutput) And() BudgetFilterPropertiesArrayOutput {
	return o.ApplyT(func(v *BudgetFilter) []BudgetFilterProperties {
		if v == nil {
			return nil
		}
		return v.And
	}).(BudgetFilterPropertiesArrayOutput)
}

func (o BudgetFilterPtrOutput) Dimensions() BudgetComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *BudgetFilter) *BudgetComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(BudgetComparisonExpressionPtrOutput)
}

func (o BudgetFilterPtrOutput) Not() BudgetFilterPropertiesPtrOutput {
	return o.ApplyT(func(v *BudgetFilter) *BudgetFilterProperties {
		if v == nil {
			return nil
		}
		return v.Not
	}).(BudgetFilterPropertiesPtrOutput)
}

func (o BudgetFilterPtrOutput) Tags() BudgetComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *BudgetFilter) *BudgetComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(BudgetComparisonExpressionPtrOutput)
}

type BudgetFilterProperties struct {
	Dimensions *BudgetComparisonExpression `pulumi:"dimensions"`
	Tags       *BudgetComparisonExpression `pulumi:"tags"`
}





type BudgetFilterPropertiesInput interface {
	pulumi.Input

	ToBudgetFilterPropertiesOutput() BudgetFilterPropertiesOutput
	ToBudgetFilterPropertiesOutputWithContext(context.Context) BudgetFilterPropertiesOutput
}

type BudgetFilterPropertiesArgs struct {
	Dimensions BudgetComparisonExpressionPtrInput `pulumi:"dimensions"`
	Tags       BudgetComparisonExpressionPtrInput `pulumi:"tags"`
}

func (BudgetFilterPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetFilterProperties)(nil)).Elem()
}

func (i BudgetFilterPropertiesArgs) ToBudgetFilterPropertiesOutput() BudgetFilterPropertiesOutput {
	return i.ToBudgetFilterPropertiesOutputWithContext(context.Background())
}

func (i BudgetFilterPropertiesArgs) ToBudgetFilterPropertiesOutputWithContext(ctx context.Context) BudgetFilterPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetFilterPropertiesOutput)
}

func (i BudgetFilterPropertiesArgs) ToBudgetFilterPropertiesPtrOutput() BudgetFilterPropertiesPtrOutput {
	return i.ToBudgetFilterPropertiesPtrOutputWithContext(context.Background())
}

func (i BudgetFilterPropertiesArgs) ToBudgetFilterPropertiesPtrOutputWithContext(ctx context.Context) BudgetFilterPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetFilterPropertiesOutput).ToBudgetFilterPropertiesPtrOutputWithContext(ctx)
}









type BudgetFilterPropertiesPtrInput interface {
	pulumi.Input

	ToBudgetFilterPropertiesPtrOutput() BudgetFilterPropertiesPtrOutput
	ToBudgetFilterPropertiesPtrOutputWithContext(context.Context) BudgetFilterPropertiesPtrOutput
}

type budgetFilterPropertiesPtrType BudgetFilterPropertiesArgs

func BudgetFilterPropertiesPtr(v *BudgetFilterPropertiesArgs) BudgetFilterPropertiesPtrInput {
	return (*budgetFilterPropertiesPtrType)(v)
}

func (*budgetFilterPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetFilterProperties)(nil)).Elem()
}

func (i *budgetFilterPropertiesPtrType) ToBudgetFilterPropertiesPtrOutput() BudgetFilterPropertiesPtrOutput {
	return i.ToBudgetFilterPropertiesPtrOutputWithContext(context.Background())
}

func (i *budgetFilterPropertiesPtrType) ToBudgetFilterPropertiesPtrOutputWithContext(ctx context.Context) BudgetFilterPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetFilterPropertiesPtrOutput)
}





type BudgetFilterPropertiesArrayInput interface {
	pulumi.Input

	ToBudgetFilterPropertiesArrayOutput() BudgetFilterPropertiesArrayOutput
	ToBudgetFilterPropertiesArrayOutputWithContext(context.Context) BudgetFilterPropertiesArrayOutput
}

type BudgetFilterPropertiesArray []BudgetFilterPropertiesInput

func (BudgetFilterPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BudgetFilterProperties)(nil)).Elem()
}

func (i BudgetFilterPropertiesArray) ToBudgetFilterPropertiesArrayOutput() BudgetFilterPropertiesArrayOutput {
	return i.ToBudgetFilterPropertiesArrayOutputWithContext(context.Background())
}

func (i BudgetFilterPropertiesArray) ToBudgetFilterPropertiesArrayOutputWithContext(ctx context.Context) BudgetFilterPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetFilterPropertiesArrayOutput)
}

type BudgetFilterPropertiesOutput struct{ *pulumi.OutputState }

func (BudgetFilterPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetFilterProperties)(nil)).Elem()
}

func (o BudgetFilterPropertiesOutput) ToBudgetFilterPropertiesOutput() BudgetFilterPropertiesOutput {
	return o
}

func (o BudgetFilterPropertiesOutput) ToBudgetFilterPropertiesOutputWithContext(ctx context.Context) BudgetFilterPropertiesOutput {
	return o
}

func (o BudgetFilterPropertiesOutput) ToBudgetFilterPropertiesPtrOutput() BudgetFilterPropertiesPtrOutput {
	return o.ToBudgetFilterPropertiesPtrOutputWithContext(context.Background())
}

func (o BudgetFilterPropertiesOutput) ToBudgetFilterPropertiesPtrOutputWithContext(ctx context.Context) BudgetFilterPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BudgetFilterProperties) *BudgetFilterProperties {
		return &v
	}).(BudgetFilterPropertiesPtrOutput)
}

func (o BudgetFilterPropertiesOutput) Dimensions() BudgetComparisonExpressionPtrOutput {
	return o.ApplyT(func(v BudgetFilterProperties) *BudgetComparisonExpression { return v.Dimensions }).(BudgetComparisonExpressionPtrOutput)
}

func (o BudgetFilterPropertiesOutput) Tags() BudgetComparisonExpressionPtrOutput {
	return o.ApplyT(func(v BudgetFilterProperties) *BudgetComparisonExpression { return v.Tags }).(BudgetComparisonExpressionPtrOutput)
}

type BudgetFilterPropertiesPtrOutput struct{ *pulumi.OutputState }

func (BudgetFilterPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetFilterProperties)(nil)).Elem()
}

func (o BudgetFilterPropertiesPtrOutput) ToBudgetFilterPropertiesPtrOutput() BudgetFilterPropertiesPtrOutput {
	return o
}

func (o BudgetFilterPropertiesPtrOutput) ToBudgetFilterPropertiesPtrOutputWithContext(ctx context.Context) BudgetFilterPropertiesPtrOutput {
	return o
}

func (o BudgetFilterPropertiesPtrOutput) Elem() BudgetFilterPropertiesOutput {
	return o.ApplyT(func(v *BudgetFilterProperties) BudgetFilterProperties {
		if v != nil {
			return *v
		}
		var ret BudgetFilterProperties
		return ret
	}).(BudgetFilterPropertiesOutput)
}

func (o BudgetFilterPropertiesPtrOutput) Dimensions() BudgetComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *BudgetFilterProperties) *BudgetComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(BudgetComparisonExpressionPtrOutput)
}

func (o BudgetFilterPropertiesPtrOutput) Tags() BudgetComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *BudgetFilterProperties) *BudgetComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(BudgetComparisonExpressionPtrOutput)
}

type BudgetFilterPropertiesArrayOutput struct{ *pulumi.OutputState }

func (BudgetFilterPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BudgetFilterProperties)(nil)).Elem()
}

func (o BudgetFilterPropertiesArrayOutput) ToBudgetFilterPropertiesArrayOutput() BudgetFilterPropertiesArrayOutput {
	return o
}

func (o BudgetFilterPropertiesArrayOutput) ToBudgetFilterPropertiesArrayOutputWithContext(ctx context.Context) BudgetFilterPropertiesArrayOutput {
	return o
}

func (o BudgetFilterPropertiesArrayOutput) Index(i pulumi.IntInput) BudgetFilterPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BudgetFilterProperties {
		return vs[0].([]BudgetFilterProperties)[vs[1].(int)]
	}).(BudgetFilterPropertiesOutput)
}

type BudgetFilterPropertiesResponse struct {
	Dimensions *BudgetComparisonExpressionResponse `pulumi:"dimensions"`
	Tags       *BudgetComparisonExpressionResponse `pulumi:"tags"`
}

type BudgetFilterPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BudgetFilterPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetFilterPropertiesResponse)(nil)).Elem()
}

func (o BudgetFilterPropertiesResponseOutput) ToBudgetFilterPropertiesResponseOutput() BudgetFilterPropertiesResponseOutput {
	return o
}

func (o BudgetFilterPropertiesResponseOutput) ToBudgetFilterPropertiesResponseOutputWithContext(ctx context.Context) BudgetFilterPropertiesResponseOutput {
	return o
}

func (o BudgetFilterPropertiesResponseOutput) Dimensions() BudgetComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v BudgetFilterPropertiesResponse) *BudgetComparisonExpressionResponse { return v.Dimensions }).(BudgetComparisonExpressionResponsePtrOutput)
}

func (o BudgetFilterPropertiesResponseOutput) Tags() BudgetComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v BudgetFilterPropertiesResponse) *BudgetComparisonExpressionResponse { return v.Tags }).(BudgetComparisonExpressionResponsePtrOutput)
}

type BudgetFilterPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (BudgetFilterPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetFilterPropertiesResponse)(nil)).Elem()
}

func (o BudgetFilterPropertiesResponsePtrOutput) ToBudgetFilterPropertiesResponsePtrOutput() BudgetFilterPropertiesResponsePtrOutput {
	return o
}

func (o BudgetFilterPropertiesResponsePtrOutput) ToBudgetFilterPropertiesResponsePtrOutputWithContext(ctx context.Context) BudgetFilterPropertiesResponsePtrOutput {
	return o
}

func (o BudgetFilterPropertiesResponsePtrOutput) Elem() BudgetFilterPropertiesResponseOutput {
	return o.ApplyT(func(v *BudgetFilterPropertiesResponse) BudgetFilterPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret BudgetFilterPropertiesResponse
		return ret
	}).(BudgetFilterPropertiesResponseOutput)
}

func (o BudgetFilterPropertiesResponsePtrOutput) Dimensions() BudgetComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *BudgetFilterPropertiesResponse) *BudgetComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(BudgetComparisonExpressionResponsePtrOutput)
}

func (o BudgetFilterPropertiesResponsePtrOutput) Tags() BudgetComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *BudgetFilterPropertiesResponse) *BudgetComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(BudgetComparisonExpressionResponsePtrOutput)
}

type BudgetFilterPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (BudgetFilterPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BudgetFilterPropertiesResponse)(nil)).Elem()
}

func (o BudgetFilterPropertiesResponseArrayOutput) ToBudgetFilterPropertiesResponseArrayOutput() BudgetFilterPropertiesResponseArrayOutput {
	return o
}

func (o BudgetFilterPropertiesResponseArrayOutput) ToBudgetFilterPropertiesResponseArrayOutputWithContext(ctx context.Context) BudgetFilterPropertiesResponseArrayOutput {
	return o
}

func (o BudgetFilterPropertiesResponseArrayOutput) Index(i pulumi.IntInput) BudgetFilterPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BudgetFilterPropertiesResponse {
		return vs[0].([]BudgetFilterPropertiesResponse)[vs[1].(int)]
	}).(BudgetFilterPropertiesResponseOutput)
}

type BudgetFilterResponse struct {
	And        []BudgetFilterPropertiesResponse    `pulumi:"and"`
	Dimensions *BudgetComparisonExpressionResponse `pulumi:"dimensions"`
	Not        *BudgetFilterPropertiesResponse     `pulumi:"not"`
	Tags       *BudgetComparisonExpressionResponse `pulumi:"tags"`
}

type BudgetFilterResponseOutput struct{ *pulumi.OutputState }

func (BudgetFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetFilterResponse)(nil)).Elem()
}

func (o BudgetFilterResponseOutput) ToBudgetFilterResponseOutput() BudgetFilterResponseOutput {
	return o
}

func (o BudgetFilterResponseOutput) ToBudgetFilterResponseOutputWithContext(ctx context.Context) BudgetFilterResponseOutput {
	return o
}

func (o BudgetFilterResponseOutput) And() BudgetFilterPropertiesResponseArrayOutput {
	return o.ApplyT(func(v BudgetFilterResponse) []BudgetFilterPropertiesResponse { return v.And }).(BudgetFilterPropertiesResponseArrayOutput)
}

func (o BudgetFilterResponseOutput) Dimensions() BudgetComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v BudgetFilterResponse) *BudgetComparisonExpressionResponse { return v.Dimensions }).(BudgetComparisonExpressionResponsePtrOutput)
}

func (o BudgetFilterResponseOutput) Not() BudgetFilterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v BudgetFilterResponse) *BudgetFilterPropertiesResponse { return v.Not }).(BudgetFilterPropertiesResponsePtrOutput)
}

func (o BudgetFilterResponseOutput) Tags() BudgetComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v BudgetFilterResponse) *BudgetComparisonExpressionResponse { return v.Tags }).(BudgetComparisonExpressionResponsePtrOutput)
}

type BudgetFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (BudgetFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetFilterResponse)(nil)).Elem()
}

func (o BudgetFilterResponsePtrOutput) ToBudgetFilterResponsePtrOutput() BudgetFilterResponsePtrOutput {
	return o
}

func (o BudgetFilterResponsePtrOutput) ToBudgetFilterResponsePtrOutputWithContext(ctx context.Context) BudgetFilterResponsePtrOutput {
	return o
}

func (o BudgetFilterResponsePtrOutput) Elem() BudgetFilterResponseOutput {
	return o.ApplyT(func(v *BudgetFilterResponse) BudgetFilterResponse {
		if v != nil {
			return *v
		}
		var ret BudgetFilterResponse
		return ret
	}).(BudgetFilterResponseOutput)
}

func (o BudgetFilterResponsePtrOutput) And() BudgetFilterPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *BudgetFilterResponse) []BudgetFilterPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.And
	}).(BudgetFilterPropertiesResponseArrayOutput)
}

func (o BudgetFilterResponsePtrOutput) Dimensions() BudgetComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *BudgetFilterResponse) *BudgetComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(BudgetComparisonExpressionResponsePtrOutput)
}

func (o BudgetFilterResponsePtrOutput) Not() BudgetFilterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *BudgetFilterResponse) *BudgetFilterPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Not
	}).(BudgetFilterPropertiesResponsePtrOutput)
}

func (o BudgetFilterResponsePtrOutput) Tags() BudgetComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *BudgetFilterResponse) *BudgetComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(BudgetComparisonExpressionResponsePtrOutput)
}

type BudgetTimePeriod struct {
	EndDate   *string `pulumi:"endDate"`
	StartDate string  `pulumi:"startDate"`
}





type BudgetTimePeriodInput interface {
	pulumi.Input

	ToBudgetTimePeriodOutput() BudgetTimePeriodOutput
	ToBudgetTimePeriodOutputWithContext(context.Context) BudgetTimePeriodOutput
}

type BudgetTimePeriodArgs struct {
	EndDate   pulumi.StringPtrInput `pulumi:"endDate"`
	StartDate pulumi.StringInput    `pulumi:"startDate"`
}

func (BudgetTimePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetTimePeriod)(nil)).Elem()
}

func (i BudgetTimePeriodArgs) ToBudgetTimePeriodOutput() BudgetTimePeriodOutput {
	return i.ToBudgetTimePeriodOutputWithContext(context.Background())
}

func (i BudgetTimePeriodArgs) ToBudgetTimePeriodOutputWithContext(ctx context.Context) BudgetTimePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BudgetTimePeriodOutput)
}

type BudgetTimePeriodOutput struct{ *pulumi.OutputState }

func (BudgetTimePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetTimePeriod)(nil)).Elem()
}

func (o BudgetTimePeriodOutput) ToBudgetTimePeriodOutput() BudgetTimePeriodOutput {
	return o
}

func (o BudgetTimePeriodOutput) ToBudgetTimePeriodOutputWithContext(ctx context.Context) BudgetTimePeriodOutput {
	return o
}

func (o BudgetTimePeriodOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BudgetTimePeriod) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o BudgetTimePeriodOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v BudgetTimePeriod) string { return v.StartDate }).(pulumi.StringOutput)
}

type BudgetTimePeriodResponse struct {
	EndDate   *string `pulumi:"endDate"`
	StartDate string  `pulumi:"startDate"`
}

type BudgetTimePeriodResponseOutput struct{ *pulumi.OutputState }

func (BudgetTimePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetTimePeriodResponse)(nil)).Elem()
}

func (o BudgetTimePeriodResponseOutput) ToBudgetTimePeriodResponseOutput() BudgetTimePeriodResponseOutput {
	return o
}

func (o BudgetTimePeriodResponseOutput) ToBudgetTimePeriodResponseOutputWithContext(ctx context.Context) BudgetTimePeriodResponseOutput {
	return o
}

func (o BudgetTimePeriodResponseOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BudgetTimePeriodResponse) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o BudgetTimePeriodResponseOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v BudgetTimePeriodResponse) string { return v.StartDate }).(pulumi.StringOutput)
}

type CurrentSpendResponse struct {
	Amount float64 `pulumi:"amount"`
	Unit   string  `pulumi:"unit"`
}

type CurrentSpendResponseOutput struct{ *pulumi.OutputState }

func (CurrentSpendResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CurrentSpendResponse)(nil)).Elem()
}

func (o CurrentSpendResponseOutput) ToCurrentSpendResponseOutput() CurrentSpendResponseOutput {
	return o
}

func (o CurrentSpendResponseOutput) ToCurrentSpendResponseOutputWithContext(ctx context.Context) CurrentSpendResponseOutput {
	return o
}

func (o CurrentSpendResponseOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v CurrentSpendResponse) float64 { return v.Amount }).(pulumi.Float64Output)
}

func (o CurrentSpendResponseOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v CurrentSpendResponse) string { return v.Unit }).(pulumi.StringOutput)
}

type ForecastSpendResponse struct {
	Amount float64 `pulumi:"amount"`
	Unit   string  `pulumi:"unit"`
}

type ForecastSpendResponseOutput struct{ *pulumi.OutputState }

func (ForecastSpendResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForecastSpendResponse)(nil)).Elem()
}

func (o ForecastSpendResponseOutput) ToForecastSpendResponseOutput() ForecastSpendResponseOutput {
	return o
}

func (o ForecastSpendResponseOutput) ToForecastSpendResponseOutputWithContext(ctx context.Context) ForecastSpendResponseOutput {
	return o
}

func (o ForecastSpendResponseOutput) Amount() pulumi.Float64Output {
	return o.ApplyT(func(v ForecastSpendResponse) float64 { return v.Amount }).(pulumi.Float64Output)
}

func (o ForecastSpendResponseOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v ForecastSpendResponse) string { return v.Unit }).(pulumi.StringOutput)
}

type Notification struct {
	ContactEmails []string `pulumi:"contactEmails"`
	ContactGroups []string `pulumi:"contactGroups"`
	ContactRoles  []string `pulumi:"contactRoles"`
	Enabled       bool     `pulumi:"enabled"`
	Operator      string   `pulumi:"operator"`
	Threshold     float64  `pulumi:"threshold"`
	ThresholdType *string  `pulumi:"thresholdType"`
}


func (val *Notification) Defaults() *Notification {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ThresholdType) {
		thresholdType_ := "Actual"
		tmp.ThresholdType = &thresholdType_
	}
	return &tmp
}





type NotificationInput interface {
	pulumi.Input

	ToNotificationOutput() NotificationOutput
	ToNotificationOutputWithContext(context.Context) NotificationOutput
}

type NotificationArgs struct {
	ContactEmails pulumi.StringArrayInput `pulumi:"contactEmails"`
	ContactGroups pulumi.StringArrayInput `pulumi:"contactGroups"`
	ContactRoles  pulumi.StringArrayInput `pulumi:"contactRoles"`
	Enabled       pulumi.BoolInput        `pulumi:"enabled"`
	Operator      pulumi.StringInput      `pulumi:"operator"`
	Threshold     pulumi.Float64Input     `pulumi:"threshold"`
	ThresholdType pulumi.StringPtrInput   `pulumi:"thresholdType"`
}


func (val *NotificationArgs) Defaults() *NotificationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ThresholdType) {
		tmp.ThresholdType = pulumi.StringPtr("Actual")
	}
	return &tmp
}
func (NotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Notification)(nil)).Elem()
}

func (i NotificationArgs) ToNotificationOutput() NotificationOutput {
	return i.ToNotificationOutputWithContext(context.Background())
}

func (i NotificationArgs) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationOutput)
}





type NotificationMapInput interface {
	pulumi.Input

	ToNotificationMapOutput() NotificationMapOutput
	ToNotificationMapOutputWithContext(context.Context) NotificationMapOutput
}

type NotificationMap map[string]NotificationInput

func (NotificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Notification)(nil)).Elem()
}

func (i NotificationMap) ToNotificationMapOutput() NotificationMapOutput {
	return i.ToNotificationMapOutputWithContext(context.Background())
}

func (i NotificationMap) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationMapOutput)
}

type NotificationOutput struct{ *pulumi.OutputState }

func (NotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Notification)(nil)).Elem()
}

func (o NotificationOutput) ToNotificationOutput() NotificationOutput {
	return o
}

func (o NotificationOutput) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return o
}

func (o NotificationOutput) ContactEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Notification) []string { return v.ContactEmails }).(pulumi.StringArrayOutput)
}

func (o NotificationOutput) ContactGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Notification) []string { return v.ContactGroups }).(pulumi.StringArrayOutput)
}

func (o NotificationOutput) ContactRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Notification) []string { return v.ContactRoles }).(pulumi.StringArrayOutput)
}

func (o NotificationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v Notification) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o NotificationOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v Notification) string { return v.Operator }).(pulumi.StringOutput)
}

func (o NotificationOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v Notification) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o NotificationOutput) ThresholdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Notification) *string { return v.ThresholdType }).(pulumi.StringPtrOutput)
}

type NotificationMapOutput struct{ *pulumi.OutputState }

func (NotificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Notification)(nil)).Elem()
}

func (o NotificationMapOutput) ToNotificationMapOutput() NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) MapIndex(k pulumi.StringInput) NotificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Notification {
		return vs[0].(map[string]Notification)[vs[1].(string)]
	}).(NotificationOutput)
}

type NotificationResponse struct {
	ContactEmails []string `pulumi:"contactEmails"`
	ContactGroups []string `pulumi:"contactGroups"`
	ContactRoles  []string `pulumi:"contactRoles"`
	Enabled       bool     `pulumi:"enabled"`
	Operator      string   `pulumi:"operator"`
	Threshold     float64  `pulumi:"threshold"`
	ThresholdType *string  `pulumi:"thresholdType"`
}


func (val *NotificationResponse) Defaults() *NotificationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ThresholdType) {
		thresholdType_ := "Actual"
		tmp.ThresholdType = &thresholdType_
	}
	return &tmp
}

type NotificationResponseOutput struct{ *pulumi.OutputState }

func (NotificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationResponse)(nil)).Elem()
}

func (o NotificationResponseOutput) ToNotificationResponseOutput() NotificationResponseOutput {
	return o
}

func (o NotificationResponseOutput) ToNotificationResponseOutputWithContext(ctx context.Context) NotificationResponseOutput {
	return o
}

func (o NotificationResponseOutput) ContactEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationResponse) []string { return v.ContactEmails }).(pulumi.StringArrayOutput)
}

func (o NotificationResponseOutput) ContactGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationResponse) []string { return v.ContactGroups }).(pulumi.StringArrayOutput)
}

func (o NotificationResponseOutput) ContactRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationResponse) []string { return v.ContactRoles }).(pulumi.StringArrayOutput)
}

func (o NotificationResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v NotificationResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o NotificationResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o NotificationResponseOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v NotificationResponse) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o NotificationResponseOutput) ThresholdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationResponse) *string { return v.ThresholdType }).(pulumi.StringPtrOutput)
}

type NotificationResponseMapOutput struct{ *pulumi.OutputState }

func (NotificationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NotificationResponse)(nil)).Elem()
}

func (o NotificationResponseMapOutput) ToNotificationResponseMapOutput() NotificationResponseMapOutput {
	return o
}

func (o NotificationResponseMapOutput) ToNotificationResponseMapOutputWithContext(ctx context.Context) NotificationResponseMapOutput {
	return o
}

func (o NotificationResponseMapOutput) MapIndex(k pulumi.StringInput) NotificationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NotificationResponse {
		return vs[0].(map[string]NotificationResponse)[vs[1].(string)]
	}).(NotificationResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(BudgetComparisonExpressionOutput{})
	pulumi.RegisterOutputType(BudgetComparisonExpressionPtrOutput{})
	pulumi.RegisterOutputType(BudgetComparisonExpressionResponseOutput{})
	pulumi.RegisterOutputType(BudgetComparisonExpressionResponsePtrOutput{})
	pulumi.RegisterOutputType(BudgetFilterOutput{})
	pulumi.RegisterOutputType(BudgetFilterPtrOutput{})
	pulumi.RegisterOutputType(BudgetFilterPropertiesOutput{})
	pulumi.RegisterOutputType(BudgetFilterPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BudgetFilterPropertiesArrayOutput{})
	pulumi.RegisterOutputType(BudgetFilterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BudgetFilterPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(BudgetFilterPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(BudgetFilterResponseOutput{})
	pulumi.RegisterOutputType(BudgetFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(BudgetTimePeriodOutput{})
	pulumi.RegisterOutputType(BudgetTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(CurrentSpendResponseOutput{})
	pulumi.RegisterOutputType(ForecastSpendResponseOutput{})
	pulumi.RegisterOutputType(NotificationOutput{})
	pulumi.RegisterOutputType(NotificationMapOutput{})
	pulumi.RegisterOutputType(NotificationResponseOutput{})
	pulumi.RegisterOutputType(NotificationResponseMapOutput{})
}
