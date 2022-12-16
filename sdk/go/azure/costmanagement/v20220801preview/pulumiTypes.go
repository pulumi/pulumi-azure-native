


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KpiProperties struct {
	Enabled *bool   `pulumi:"enabled"`
	Id      *string `pulumi:"id"`
	Type    *string `pulumi:"type"`
}





type KpiPropertiesInput interface {
	pulumi.Input

	ToKpiPropertiesOutput() KpiPropertiesOutput
	ToKpiPropertiesOutputWithContext(context.Context) KpiPropertiesOutput
}

type KpiPropertiesArgs struct {
	Enabled pulumi.BoolPtrInput   `pulumi:"enabled"`
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Type    pulumi.StringPtrInput `pulumi:"type"`
}

func (KpiPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiProperties)(nil)).Elem()
}

func (i KpiPropertiesArgs) ToKpiPropertiesOutput() KpiPropertiesOutput {
	return i.ToKpiPropertiesOutputWithContext(context.Background())
}

func (i KpiPropertiesArgs) ToKpiPropertiesOutputWithContext(ctx context.Context) KpiPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiPropertiesOutput)
}





type KpiPropertiesArrayInput interface {
	pulumi.Input

	ToKpiPropertiesArrayOutput() KpiPropertiesArrayOutput
	ToKpiPropertiesArrayOutputWithContext(context.Context) KpiPropertiesArrayOutput
}

type KpiPropertiesArray []KpiPropertiesInput

func (KpiPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiProperties)(nil)).Elem()
}

func (i KpiPropertiesArray) ToKpiPropertiesArrayOutput() KpiPropertiesArrayOutput {
	return i.ToKpiPropertiesArrayOutputWithContext(context.Background())
}

func (i KpiPropertiesArray) ToKpiPropertiesArrayOutputWithContext(ctx context.Context) KpiPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiPropertiesArrayOutput)
}

type KpiPropertiesOutput struct{ *pulumi.OutputState }

func (KpiPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiProperties)(nil)).Elem()
}

func (o KpiPropertiesOutput) ToKpiPropertiesOutput() KpiPropertiesOutput {
	return o
}

func (o KpiPropertiesOutput) ToKpiPropertiesOutputWithContext(ctx context.Context) KpiPropertiesOutput {
	return o
}

func (o KpiPropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KpiProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o KpiPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KpiPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type KpiPropertiesArrayOutput struct{ *pulumi.OutputState }

func (KpiPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiProperties)(nil)).Elem()
}

func (o KpiPropertiesArrayOutput) ToKpiPropertiesArrayOutput() KpiPropertiesArrayOutput {
	return o
}

func (o KpiPropertiesArrayOutput) ToKpiPropertiesArrayOutputWithContext(ctx context.Context) KpiPropertiesArrayOutput {
	return o
}

func (o KpiPropertiesArrayOutput) Index(i pulumi.IntInput) KpiPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiProperties {
		return vs[0].([]KpiProperties)[vs[1].(int)]
	}).(KpiPropertiesOutput)
}

type KpiPropertiesResponse struct {
	Enabled *bool   `pulumi:"enabled"`
	Id      *string `pulumi:"id"`
	Type    *string `pulumi:"type"`
}

type KpiPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KpiPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiPropertiesResponse)(nil)).Elem()
}

func (o KpiPropertiesResponseOutput) ToKpiPropertiesResponseOutput() KpiPropertiesResponseOutput {
	return o
}

func (o KpiPropertiesResponseOutput) ToKpiPropertiesResponseOutputWithContext(ctx context.Context) KpiPropertiesResponseOutput {
	return o
}

func (o KpiPropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KpiPropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o KpiPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KpiPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type KpiPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (KpiPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiPropertiesResponse)(nil)).Elem()
}

func (o KpiPropertiesResponseArrayOutput) ToKpiPropertiesResponseArrayOutput() KpiPropertiesResponseArrayOutput {
	return o
}

func (o KpiPropertiesResponseArrayOutput) ToKpiPropertiesResponseArrayOutputWithContext(ctx context.Context) KpiPropertiesResponseArrayOutput {
	return o
}

func (o KpiPropertiesResponseArrayOutput) Index(i pulumi.IntInput) KpiPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiPropertiesResponse {
		return vs[0].([]KpiPropertiesResponse)[vs[1].(int)]
	}).(KpiPropertiesResponseOutput)
}

type PivotProperties struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type PivotPropertiesInput interface {
	pulumi.Input

	ToPivotPropertiesOutput() PivotPropertiesOutput
	ToPivotPropertiesOutputWithContext(context.Context) PivotPropertiesOutput
}

type PivotPropertiesArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PivotPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PivotProperties)(nil)).Elem()
}

func (i PivotPropertiesArgs) ToPivotPropertiesOutput() PivotPropertiesOutput {
	return i.ToPivotPropertiesOutputWithContext(context.Background())
}

func (i PivotPropertiesArgs) ToPivotPropertiesOutputWithContext(ctx context.Context) PivotPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PivotPropertiesOutput)
}





type PivotPropertiesArrayInput interface {
	pulumi.Input

	ToPivotPropertiesArrayOutput() PivotPropertiesArrayOutput
	ToPivotPropertiesArrayOutputWithContext(context.Context) PivotPropertiesArrayOutput
}

type PivotPropertiesArray []PivotPropertiesInput

func (PivotPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PivotProperties)(nil)).Elem()
}

func (i PivotPropertiesArray) ToPivotPropertiesArrayOutput() PivotPropertiesArrayOutput {
	return i.ToPivotPropertiesArrayOutputWithContext(context.Background())
}

func (i PivotPropertiesArray) ToPivotPropertiesArrayOutputWithContext(ctx context.Context) PivotPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PivotPropertiesArrayOutput)
}

type PivotPropertiesOutput struct{ *pulumi.OutputState }

func (PivotPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PivotProperties)(nil)).Elem()
}

func (o PivotPropertiesOutput) ToPivotPropertiesOutput() PivotPropertiesOutput {
	return o
}

func (o PivotPropertiesOutput) ToPivotPropertiesOutputWithContext(ctx context.Context) PivotPropertiesOutput {
	return o
}

func (o PivotPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PivotProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PivotPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PivotProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PivotPropertiesArrayOutput struct{ *pulumi.OutputState }

func (PivotPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PivotProperties)(nil)).Elem()
}

func (o PivotPropertiesArrayOutput) ToPivotPropertiesArrayOutput() PivotPropertiesArrayOutput {
	return o
}

func (o PivotPropertiesArrayOutput) ToPivotPropertiesArrayOutputWithContext(ctx context.Context) PivotPropertiesArrayOutput {
	return o
}

func (o PivotPropertiesArrayOutput) Index(i pulumi.IntInput) PivotPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PivotProperties {
		return vs[0].([]PivotProperties)[vs[1].(int)]
	}).(PivotPropertiesOutput)
}

type PivotPropertiesResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type PivotPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PivotPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PivotPropertiesResponse)(nil)).Elem()
}

func (o PivotPropertiesResponseOutput) ToPivotPropertiesResponseOutput() PivotPropertiesResponseOutput {
	return o
}

func (o PivotPropertiesResponseOutput) ToPivotPropertiesResponseOutputWithContext(ctx context.Context) PivotPropertiesResponseOutput {
	return o
}

func (o PivotPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PivotPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PivotPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PivotPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PivotPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (PivotPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PivotPropertiesResponse)(nil)).Elem()
}

func (o PivotPropertiesResponseArrayOutput) ToPivotPropertiesResponseArrayOutput() PivotPropertiesResponseArrayOutput {
	return o
}

func (o PivotPropertiesResponseArrayOutput) ToPivotPropertiesResponseArrayOutputWithContext(ctx context.Context) PivotPropertiesResponseArrayOutput {
	return o
}

func (o PivotPropertiesResponseArrayOutput) Index(i pulumi.IntInput) PivotPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PivotPropertiesResponse {
		return vs[0].([]PivotPropertiesResponse)[vs[1].(int)]
	}).(PivotPropertiesResponseOutput)
}

type ReportConfigAggregation struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}





type ReportConfigAggregationInput interface {
	pulumi.Input

	ToReportConfigAggregationOutput() ReportConfigAggregationOutput
	ToReportConfigAggregationOutputWithContext(context.Context) ReportConfigAggregationOutput
}

type ReportConfigAggregationArgs struct {
	Function pulumi.StringInput `pulumi:"function"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (ReportConfigAggregationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigAggregation)(nil)).Elem()
}

func (i ReportConfigAggregationArgs) ToReportConfigAggregationOutput() ReportConfigAggregationOutput {
	return i.ToReportConfigAggregationOutputWithContext(context.Background())
}

func (i ReportConfigAggregationArgs) ToReportConfigAggregationOutputWithContext(ctx context.Context) ReportConfigAggregationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigAggregationOutput)
}





type ReportConfigAggregationMapInput interface {
	pulumi.Input

	ToReportConfigAggregationMapOutput() ReportConfigAggregationMapOutput
	ToReportConfigAggregationMapOutputWithContext(context.Context) ReportConfigAggregationMapOutput
}

type ReportConfigAggregationMap map[string]ReportConfigAggregationInput

func (ReportConfigAggregationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportConfigAggregation)(nil)).Elem()
}

func (i ReportConfigAggregationMap) ToReportConfigAggregationMapOutput() ReportConfigAggregationMapOutput {
	return i.ToReportConfigAggregationMapOutputWithContext(context.Background())
}

func (i ReportConfigAggregationMap) ToReportConfigAggregationMapOutputWithContext(ctx context.Context) ReportConfigAggregationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigAggregationMapOutput)
}

type ReportConfigAggregationOutput struct{ *pulumi.OutputState }

func (ReportConfigAggregationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigAggregation)(nil)).Elem()
}

func (o ReportConfigAggregationOutput) ToReportConfigAggregationOutput() ReportConfigAggregationOutput {
	return o
}

func (o ReportConfigAggregationOutput) ToReportConfigAggregationOutputWithContext(ctx context.Context) ReportConfigAggregationOutput {
	return o
}

func (o ReportConfigAggregationOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigAggregation) string { return v.Function }).(pulumi.StringOutput)
}

func (o ReportConfigAggregationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigAggregation) string { return v.Name }).(pulumi.StringOutput)
}

type ReportConfigAggregationMapOutput struct{ *pulumi.OutputState }

func (ReportConfigAggregationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportConfigAggregation)(nil)).Elem()
}

func (o ReportConfigAggregationMapOutput) ToReportConfigAggregationMapOutput() ReportConfigAggregationMapOutput {
	return o
}

func (o ReportConfigAggregationMapOutput) ToReportConfigAggregationMapOutputWithContext(ctx context.Context) ReportConfigAggregationMapOutput {
	return o
}

func (o ReportConfigAggregationMapOutput) MapIndex(k pulumi.StringInput) ReportConfigAggregationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReportConfigAggregation {
		return vs[0].(map[string]ReportConfigAggregation)[vs[1].(string)]
	}).(ReportConfigAggregationOutput)
}

type ReportConfigAggregationResponse struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}

type ReportConfigAggregationResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigAggregationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigAggregationResponse)(nil)).Elem()
}

func (o ReportConfigAggregationResponseOutput) ToReportConfigAggregationResponseOutput() ReportConfigAggregationResponseOutput {
	return o
}

func (o ReportConfigAggregationResponseOutput) ToReportConfigAggregationResponseOutputWithContext(ctx context.Context) ReportConfigAggregationResponseOutput {
	return o
}

func (o ReportConfigAggregationResponseOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigAggregationResponse) string { return v.Function }).(pulumi.StringOutput)
}

func (o ReportConfigAggregationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigAggregationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ReportConfigAggregationResponseMapOutput struct{ *pulumi.OutputState }

func (ReportConfigAggregationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportConfigAggregationResponse)(nil)).Elem()
}

func (o ReportConfigAggregationResponseMapOutput) ToReportConfigAggregationResponseMapOutput() ReportConfigAggregationResponseMapOutput {
	return o
}

func (o ReportConfigAggregationResponseMapOutput) ToReportConfigAggregationResponseMapOutputWithContext(ctx context.Context) ReportConfigAggregationResponseMapOutput {
	return o
}

func (o ReportConfigAggregationResponseMapOutput) MapIndex(k pulumi.StringInput) ReportConfigAggregationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReportConfigAggregationResponse {
		return vs[0].(map[string]ReportConfigAggregationResponse)[vs[1].(string)]
	}).(ReportConfigAggregationResponseOutput)
}

type ReportConfigComparisonExpression struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type ReportConfigComparisonExpressionInput interface {
	pulumi.Input

	ToReportConfigComparisonExpressionOutput() ReportConfigComparisonExpressionOutput
	ToReportConfigComparisonExpressionOutputWithContext(context.Context) ReportConfigComparisonExpressionOutput
}

type ReportConfigComparisonExpressionArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (ReportConfigComparisonExpressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigComparisonExpression)(nil)).Elem()
}

func (i ReportConfigComparisonExpressionArgs) ToReportConfigComparisonExpressionOutput() ReportConfigComparisonExpressionOutput {
	return i.ToReportConfigComparisonExpressionOutputWithContext(context.Background())
}

func (i ReportConfigComparisonExpressionArgs) ToReportConfigComparisonExpressionOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigComparisonExpressionOutput)
}

func (i ReportConfigComparisonExpressionArgs) ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput {
	return i.ToReportConfigComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i ReportConfigComparisonExpressionArgs) ToReportConfigComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigComparisonExpressionOutput).ToReportConfigComparisonExpressionPtrOutputWithContext(ctx)
}









type ReportConfigComparisonExpressionPtrInput interface {
	pulumi.Input

	ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput
	ToReportConfigComparisonExpressionPtrOutputWithContext(context.Context) ReportConfigComparisonExpressionPtrOutput
}

type reportConfigComparisonExpressionPtrType ReportConfigComparisonExpressionArgs

func ReportConfigComparisonExpressionPtr(v *ReportConfigComparisonExpressionArgs) ReportConfigComparisonExpressionPtrInput {
	return (*reportConfigComparisonExpressionPtrType)(v)
}

func (*reportConfigComparisonExpressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigComparisonExpression)(nil)).Elem()
}

func (i *reportConfigComparisonExpressionPtrType) ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput {
	return i.ToReportConfigComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i *reportConfigComparisonExpressionPtrType) ToReportConfigComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigComparisonExpressionPtrOutput)
}

type ReportConfigComparisonExpressionOutput struct{ *pulumi.OutputState }

func (ReportConfigComparisonExpressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigComparisonExpression)(nil)).Elem()
}

func (o ReportConfigComparisonExpressionOutput) ToReportConfigComparisonExpressionOutput() ReportConfigComparisonExpressionOutput {
	return o
}

func (o ReportConfigComparisonExpressionOutput) ToReportConfigComparisonExpressionOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionOutput {
	return o
}

func (o ReportConfigComparisonExpressionOutput) ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput {
	return o.ToReportConfigComparisonExpressionPtrOutputWithContext(context.Background())
}

func (o ReportConfigComparisonExpressionOutput) ToReportConfigComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigComparisonExpression) *ReportConfigComparisonExpression {
		return &v
	}).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigComparisonExpressionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpression) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportConfigComparisonExpressionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpression) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ReportConfigComparisonExpressionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpression) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ReportConfigComparisonExpressionPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigComparisonExpressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigComparisonExpression)(nil)).Elem()
}

func (o ReportConfigComparisonExpressionPtrOutput) ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput {
	return o
}

func (o ReportConfigComparisonExpressionPtrOutput) ToReportConfigComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionPtrOutput {
	return o
}

func (o ReportConfigComparisonExpressionPtrOutput) Elem() ReportConfigComparisonExpressionOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpression) ReportConfigComparisonExpression {
		if v != nil {
			return *v
		}
		var ret ReportConfigComparisonExpression
		return ret
	}).(ReportConfigComparisonExpressionOutput)
}

func (o ReportConfigComparisonExpressionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigComparisonExpressionPtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigComparisonExpressionPtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpression) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type ReportConfigComparisonExpressionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type ReportConfigComparisonExpressionResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigComparisonExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigComparisonExpressionResponse)(nil)).Elem()
}

func (o ReportConfigComparisonExpressionResponseOutput) ToReportConfigComparisonExpressionResponseOutput() ReportConfigComparisonExpressionResponseOutput {
	return o
}

func (o ReportConfigComparisonExpressionResponseOutput) ToReportConfigComparisonExpressionResponseOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionResponseOutput {
	return o
}

func (o ReportConfigComparisonExpressionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpressionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportConfigComparisonExpressionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpressionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ReportConfigComparisonExpressionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpressionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ReportConfigComparisonExpressionResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigComparisonExpressionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigComparisonExpressionResponse)(nil)).Elem()
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) ToReportConfigComparisonExpressionResponsePtrOutput() ReportConfigComparisonExpressionResponsePtrOutput {
	return o
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) ToReportConfigComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionResponsePtrOutput {
	return o
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) Elem() ReportConfigComparisonExpressionResponseOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpressionResponse) ReportConfigComparisonExpressionResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigComparisonExpressionResponse
		return ret
	}).(ReportConfigComparisonExpressionResponseOutput)
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpressionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type ReportConfigDataset struct {
	Aggregation   map[string]ReportConfigAggregation `pulumi:"aggregation"`
	Configuration *ReportConfigDatasetConfiguration  `pulumi:"configuration"`
	Filter        *ReportConfigFilter                `pulumi:"filter"`
	Granularity   *string                            `pulumi:"granularity"`
	Grouping      []ReportConfigGrouping             `pulumi:"grouping"`
	Sorting       []ReportConfigSorting              `pulumi:"sorting"`
}





type ReportConfigDatasetInput interface {
	pulumi.Input

	ToReportConfigDatasetOutput() ReportConfigDatasetOutput
	ToReportConfigDatasetOutputWithContext(context.Context) ReportConfigDatasetOutput
}

type ReportConfigDatasetArgs struct {
	Aggregation   ReportConfigAggregationMapInput          `pulumi:"aggregation"`
	Configuration ReportConfigDatasetConfigurationPtrInput `pulumi:"configuration"`
	Filter        ReportConfigFilterPtrInput               `pulumi:"filter"`
	Granularity   pulumi.StringPtrInput                    `pulumi:"granularity"`
	Grouping      ReportConfigGroupingArrayInput           `pulumi:"grouping"`
	Sorting       ReportConfigSortingArrayInput            `pulumi:"sorting"`
}

func (ReportConfigDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDataset)(nil)).Elem()
}

func (i ReportConfigDatasetArgs) ToReportConfigDatasetOutput() ReportConfigDatasetOutput {
	return i.ToReportConfigDatasetOutputWithContext(context.Background())
}

func (i ReportConfigDatasetArgs) ToReportConfigDatasetOutputWithContext(ctx context.Context) ReportConfigDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetOutput)
}

func (i ReportConfigDatasetArgs) ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput {
	return i.ToReportConfigDatasetPtrOutputWithContext(context.Background())
}

func (i ReportConfigDatasetArgs) ToReportConfigDatasetPtrOutputWithContext(ctx context.Context) ReportConfigDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetOutput).ToReportConfigDatasetPtrOutputWithContext(ctx)
}









type ReportConfigDatasetPtrInput interface {
	pulumi.Input

	ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput
	ToReportConfigDatasetPtrOutputWithContext(context.Context) ReportConfigDatasetPtrOutput
}

type reportConfigDatasetPtrType ReportConfigDatasetArgs

func ReportConfigDatasetPtr(v *ReportConfigDatasetArgs) ReportConfigDatasetPtrInput {
	return (*reportConfigDatasetPtrType)(v)
}

func (*reportConfigDatasetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDataset)(nil)).Elem()
}

func (i *reportConfigDatasetPtrType) ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput {
	return i.ToReportConfigDatasetPtrOutputWithContext(context.Background())
}

func (i *reportConfigDatasetPtrType) ToReportConfigDatasetPtrOutputWithContext(ctx context.Context) ReportConfigDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetPtrOutput)
}

type ReportConfigDatasetOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDataset)(nil)).Elem()
}

func (o ReportConfigDatasetOutput) ToReportConfigDatasetOutput() ReportConfigDatasetOutput {
	return o
}

func (o ReportConfigDatasetOutput) ToReportConfigDatasetOutputWithContext(ctx context.Context) ReportConfigDatasetOutput {
	return o
}

func (o ReportConfigDatasetOutput) ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput {
	return o.ToReportConfigDatasetPtrOutputWithContext(context.Background())
}

func (o ReportConfigDatasetOutput) ToReportConfigDatasetPtrOutputWithContext(ctx context.Context) ReportConfigDatasetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigDataset) *ReportConfigDataset {
		return &v
	}).(ReportConfigDatasetPtrOutput)
}

func (o ReportConfigDatasetOutput) Aggregation() ReportConfigAggregationMapOutput {
	return o.ApplyT(func(v ReportConfigDataset) map[string]ReportConfigAggregation { return v.Aggregation }).(ReportConfigAggregationMapOutput)
}

func (o ReportConfigDatasetOutput) Configuration() ReportConfigDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v ReportConfigDataset) *ReportConfigDatasetConfiguration { return v.Configuration }).(ReportConfigDatasetConfigurationPtrOutput)
}

func (o ReportConfigDatasetOutput) Filter() ReportConfigFilterPtrOutput {
	return o.ApplyT(func(v ReportConfigDataset) *ReportConfigFilter { return v.Filter }).(ReportConfigFilterPtrOutput)
}

func (o ReportConfigDatasetOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigDataset) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o ReportConfigDatasetOutput) Grouping() ReportConfigGroupingArrayOutput {
	return o.ApplyT(func(v ReportConfigDataset) []ReportConfigGrouping { return v.Grouping }).(ReportConfigGroupingArrayOutput)
}

func (o ReportConfigDatasetOutput) Sorting() ReportConfigSortingArrayOutput {
	return o.ApplyT(func(v ReportConfigDataset) []ReportConfigSorting { return v.Sorting }).(ReportConfigSortingArrayOutput)
}

type ReportConfigDatasetPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDataset)(nil)).Elem()
}

func (o ReportConfigDatasetPtrOutput) ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput {
	return o
}

func (o ReportConfigDatasetPtrOutput) ToReportConfigDatasetPtrOutputWithContext(ctx context.Context) ReportConfigDatasetPtrOutput {
	return o
}

func (o ReportConfigDatasetPtrOutput) Elem() ReportConfigDatasetOutput {
	return o.ApplyT(func(v *ReportConfigDataset) ReportConfigDataset {
		if v != nil {
			return *v
		}
		var ret ReportConfigDataset
		return ret
	}).(ReportConfigDatasetOutput)
}

func (o ReportConfigDatasetPtrOutput) Aggregation() ReportConfigAggregationMapOutput {
	return o.ApplyT(func(v *ReportConfigDataset) map[string]ReportConfigAggregation {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(ReportConfigAggregationMapOutput)
}

func (o ReportConfigDatasetPtrOutput) Configuration() ReportConfigDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v *ReportConfigDataset) *ReportConfigDatasetConfiguration {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ReportConfigDatasetConfigurationPtrOutput)
}

func (o ReportConfigDatasetPtrOutput) Filter() ReportConfigFilterPtrOutput {
	return o.ApplyT(func(v *ReportConfigDataset) *ReportConfigFilter {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(ReportConfigFilterPtrOutput)
}

func (o ReportConfigDatasetPtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigDataset) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigDatasetPtrOutput) Grouping() ReportConfigGroupingArrayOutput {
	return o.ApplyT(func(v *ReportConfigDataset) []ReportConfigGrouping {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(ReportConfigGroupingArrayOutput)
}

func (o ReportConfigDatasetPtrOutput) Sorting() ReportConfigSortingArrayOutput {
	return o.ApplyT(func(v *ReportConfigDataset) []ReportConfigSorting {
		if v == nil {
			return nil
		}
		return v.Sorting
	}).(ReportConfigSortingArrayOutput)
}

type ReportConfigDatasetConfiguration struct {
	Columns []string `pulumi:"columns"`
}





type ReportConfigDatasetConfigurationInput interface {
	pulumi.Input

	ToReportConfigDatasetConfigurationOutput() ReportConfigDatasetConfigurationOutput
	ToReportConfigDatasetConfigurationOutputWithContext(context.Context) ReportConfigDatasetConfigurationOutput
}

type ReportConfigDatasetConfigurationArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (ReportConfigDatasetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetConfiguration)(nil)).Elem()
}

func (i ReportConfigDatasetConfigurationArgs) ToReportConfigDatasetConfigurationOutput() ReportConfigDatasetConfigurationOutput {
	return i.ToReportConfigDatasetConfigurationOutputWithContext(context.Background())
}

func (i ReportConfigDatasetConfigurationArgs) ToReportConfigDatasetConfigurationOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetConfigurationOutput)
}

func (i ReportConfigDatasetConfigurationArgs) ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput {
	return i.ToReportConfigDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i ReportConfigDatasetConfigurationArgs) ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetConfigurationOutput).ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx)
}









type ReportConfigDatasetConfigurationPtrInput interface {
	pulumi.Input

	ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput
	ToReportConfigDatasetConfigurationPtrOutputWithContext(context.Context) ReportConfigDatasetConfigurationPtrOutput
}

type reportConfigDatasetConfigurationPtrType ReportConfigDatasetConfigurationArgs

func ReportConfigDatasetConfigurationPtr(v *ReportConfigDatasetConfigurationArgs) ReportConfigDatasetConfigurationPtrInput {
	return (*reportConfigDatasetConfigurationPtrType)(v)
}

func (*reportConfigDatasetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetConfiguration)(nil)).Elem()
}

func (i *reportConfigDatasetConfigurationPtrType) ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput {
	return i.ToReportConfigDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i *reportConfigDatasetConfigurationPtrType) ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetConfigurationPtrOutput)
}

type ReportConfigDatasetConfigurationOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetConfiguration)(nil)).Elem()
}

func (o ReportConfigDatasetConfigurationOutput) ToReportConfigDatasetConfigurationOutput() ReportConfigDatasetConfigurationOutput {
	return o
}

func (o ReportConfigDatasetConfigurationOutput) ToReportConfigDatasetConfigurationOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationOutput {
	return o
}

func (o ReportConfigDatasetConfigurationOutput) ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput {
	return o.ToReportConfigDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (o ReportConfigDatasetConfigurationOutput) ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigDatasetConfiguration) *ReportConfigDatasetConfiguration {
		return &v
	}).(ReportConfigDatasetConfigurationPtrOutput)
}

func (o ReportConfigDatasetConfigurationOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportConfigDatasetConfiguration) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ReportConfigDatasetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetConfiguration)(nil)).Elem()
}

func (o ReportConfigDatasetConfigurationPtrOutput) ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput {
	return o
}

func (o ReportConfigDatasetConfigurationPtrOutput) ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationPtrOutput {
	return o
}

func (o ReportConfigDatasetConfigurationPtrOutput) Elem() ReportConfigDatasetConfigurationOutput {
	return o.ApplyT(func(v *ReportConfigDatasetConfiguration) ReportConfigDatasetConfiguration {
		if v != nil {
			return *v
		}
		var ret ReportConfigDatasetConfiguration
		return ret
	}).(ReportConfigDatasetConfigurationOutput)
}

func (o ReportConfigDatasetConfigurationPtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportConfigDatasetConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ReportConfigDatasetConfigurationResponse struct {
	Columns []string `pulumi:"columns"`
}

type ReportConfigDatasetConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetConfigurationResponse)(nil)).Elem()
}

func (o ReportConfigDatasetConfigurationResponseOutput) ToReportConfigDatasetConfigurationResponseOutput() ReportConfigDatasetConfigurationResponseOutput {
	return o
}

func (o ReportConfigDatasetConfigurationResponseOutput) ToReportConfigDatasetConfigurationResponseOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationResponseOutput {
	return o
}

func (o ReportConfigDatasetConfigurationResponseOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportConfigDatasetConfigurationResponse) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ReportConfigDatasetConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetConfigurationResponse)(nil)).Elem()
}

func (o ReportConfigDatasetConfigurationResponsePtrOutput) ToReportConfigDatasetConfigurationResponsePtrOutput() ReportConfigDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ReportConfigDatasetConfigurationResponsePtrOutput) ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ReportConfigDatasetConfigurationResponsePtrOutput) Elem() ReportConfigDatasetConfigurationResponseOutput {
	return o.ApplyT(func(v *ReportConfigDatasetConfigurationResponse) ReportConfigDatasetConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigDatasetConfigurationResponse
		return ret
	}).(ReportConfigDatasetConfigurationResponseOutput)
}

func (o ReportConfigDatasetConfigurationResponsePtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportConfigDatasetConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ReportConfigDatasetResponse struct {
	Aggregation   map[string]ReportConfigAggregationResponse `pulumi:"aggregation"`
	Configuration *ReportConfigDatasetConfigurationResponse  `pulumi:"configuration"`
	Filter        *ReportConfigFilterResponse                `pulumi:"filter"`
	Granularity   *string                                    `pulumi:"granularity"`
	Grouping      []ReportConfigGroupingResponse             `pulumi:"grouping"`
	Sorting       []ReportConfigSortingResponse              `pulumi:"sorting"`
}

type ReportConfigDatasetResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetResponse)(nil)).Elem()
}

func (o ReportConfigDatasetResponseOutput) ToReportConfigDatasetResponseOutput() ReportConfigDatasetResponseOutput {
	return o
}

func (o ReportConfigDatasetResponseOutput) ToReportConfigDatasetResponseOutputWithContext(ctx context.Context) ReportConfigDatasetResponseOutput {
	return o
}

func (o ReportConfigDatasetResponseOutput) Aggregation() ReportConfigAggregationResponseMapOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) map[string]ReportConfigAggregationResponse { return v.Aggregation }).(ReportConfigAggregationResponseMapOutput)
}

func (o ReportConfigDatasetResponseOutput) Configuration() ReportConfigDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) *ReportConfigDatasetConfigurationResponse { return v.Configuration }).(ReportConfigDatasetConfigurationResponsePtrOutput)
}

func (o ReportConfigDatasetResponseOutput) Filter() ReportConfigFilterResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) *ReportConfigFilterResponse { return v.Filter }).(ReportConfigFilterResponsePtrOutput)
}

func (o ReportConfigDatasetResponseOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o ReportConfigDatasetResponseOutput) Grouping() ReportConfigGroupingResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) []ReportConfigGroupingResponse { return v.Grouping }).(ReportConfigGroupingResponseArrayOutput)
}

func (o ReportConfigDatasetResponseOutput) Sorting() ReportConfigSortingResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) []ReportConfigSortingResponse { return v.Sorting }).(ReportConfigSortingResponseArrayOutput)
}

type ReportConfigDatasetResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetResponse)(nil)).Elem()
}

func (o ReportConfigDatasetResponsePtrOutput) ToReportConfigDatasetResponsePtrOutput() ReportConfigDatasetResponsePtrOutput {
	return o
}

func (o ReportConfigDatasetResponsePtrOutput) ToReportConfigDatasetResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetResponsePtrOutput {
	return o
}

func (o ReportConfigDatasetResponsePtrOutput) Elem() ReportConfigDatasetResponseOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) ReportConfigDatasetResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigDatasetResponse
		return ret
	}).(ReportConfigDatasetResponseOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Aggregation() ReportConfigAggregationResponseMapOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) map[string]ReportConfigAggregationResponse {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(ReportConfigAggregationResponseMapOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Configuration() ReportConfigDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) *ReportConfigDatasetConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ReportConfigDatasetConfigurationResponsePtrOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Filter() ReportConfigFilterResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) *ReportConfigFilterResponse {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(ReportConfigFilterResponsePtrOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Grouping() ReportConfigGroupingResponseArrayOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) []ReportConfigGroupingResponse {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(ReportConfigGroupingResponseArrayOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Sorting() ReportConfigSortingResponseArrayOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) []ReportConfigSortingResponse {
		if v == nil {
			return nil
		}
		return v.Sorting
	}).(ReportConfigSortingResponseArrayOutput)
}

type ReportConfigFilter struct {
	And        []ReportConfigFilter              `pulumi:"and"`
	Dimensions *ReportConfigComparisonExpression `pulumi:"dimensions"`
	Or         []ReportConfigFilter              `pulumi:"or"`
	Tags       *ReportConfigComparisonExpression `pulumi:"tags"`
}





type ReportConfigFilterInput interface {
	pulumi.Input

	ToReportConfigFilterOutput() ReportConfigFilterOutput
	ToReportConfigFilterOutputWithContext(context.Context) ReportConfigFilterOutput
}

type ReportConfigFilterArgs struct {
	And        ReportConfigFilterArrayInput             `pulumi:"and"`
	Dimensions ReportConfigComparisonExpressionPtrInput `pulumi:"dimensions"`
	Or         ReportConfigFilterArrayInput             `pulumi:"or"`
	Tags       ReportConfigComparisonExpressionPtrInput `pulumi:"tags"`
}

func (ReportConfigFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigFilter)(nil)).Elem()
}

func (i ReportConfigFilterArgs) ToReportConfigFilterOutput() ReportConfigFilterOutput {
	return i.ToReportConfigFilterOutputWithContext(context.Background())
}

func (i ReportConfigFilterArgs) ToReportConfigFilterOutputWithContext(ctx context.Context) ReportConfigFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterOutput)
}

func (i ReportConfigFilterArgs) ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput {
	return i.ToReportConfigFilterPtrOutputWithContext(context.Background())
}

func (i ReportConfigFilterArgs) ToReportConfigFilterPtrOutputWithContext(ctx context.Context) ReportConfigFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterOutput).ToReportConfigFilterPtrOutputWithContext(ctx)
}









type ReportConfigFilterPtrInput interface {
	pulumi.Input

	ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput
	ToReportConfigFilterPtrOutputWithContext(context.Context) ReportConfigFilterPtrOutput
}

type reportConfigFilterPtrType ReportConfigFilterArgs

func ReportConfigFilterPtr(v *ReportConfigFilterArgs) ReportConfigFilterPtrInput {
	return (*reportConfigFilterPtrType)(v)
}

func (*reportConfigFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigFilter)(nil)).Elem()
}

func (i *reportConfigFilterPtrType) ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput {
	return i.ToReportConfigFilterPtrOutputWithContext(context.Background())
}

func (i *reportConfigFilterPtrType) ToReportConfigFilterPtrOutputWithContext(ctx context.Context) ReportConfigFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterPtrOutput)
}





type ReportConfigFilterArrayInput interface {
	pulumi.Input

	ToReportConfigFilterArrayOutput() ReportConfigFilterArrayOutput
	ToReportConfigFilterArrayOutputWithContext(context.Context) ReportConfigFilterArrayOutput
}

type ReportConfigFilterArray []ReportConfigFilterInput

func (ReportConfigFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigFilter)(nil)).Elem()
}

func (i ReportConfigFilterArray) ToReportConfigFilterArrayOutput() ReportConfigFilterArrayOutput {
	return i.ToReportConfigFilterArrayOutputWithContext(context.Background())
}

func (i ReportConfigFilterArray) ToReportConfigFilterArrayOutputWithContext(ctx context.Context) ReportConfigFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterArrayOutput)
}

type ReportConfigFilterOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigFilter)(nil)).Elem()
}

func (o ReportConfigFilterOutput) ToReportConfigFilterOutput() ReportConfigFilterOutput {
	return o
}

func (o ReportConfigFilterOutput) ToReportConfigFilterOutputWithContext(ctx context.Context) ReportConfigFilterOutput {
	return o
}

func (o ReportConfigFilterOutput) ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput {
	return o.ToReportConfigFilterPtrOutputWithContext(context.Background())
}

func (o ReportConfigFilterOutput) ToReportConfigFilterPtrOutputWithContext(ctx context.Context) ReportConfigFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigFilter) *ReportConfigFilter {
		return &v
	}).(ReportConfigFilterPtrOutput)
}

func (o ReportConfigFilterOutput) And() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v ReportConfigFilter) []ReportConfigFilter { return v.And }).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterOutput) Dimensions() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportConfigFilter) *ReportConfigComparisonExpression { return v.Dimensions }).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterOutput) Or() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v ReportConfigFilter) []ReportConfigFilter { return v.Or }).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterOutput) Tags() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportConfigFilter) *ReportConfigComparisonExpression { return v.Tags }).(ReportConfigComparisonExpressionPtrOutput)
}

type ReportConfigFilterPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigFilter)(nil)).Elem()
}

func (o ReportConfigFilterPtrOutput) ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput {
	return o
}

func (o ReportConfigFilterPtrOutput) ToReportConfigFilterPtrOutputWithContext(ctx context.Context) ReportConfigFilterPtrOutput {
	return o
}

func (o ReportConfigFilterPtrOutput) Elem() ReportConfigFilterOutput {
	return o.ApplyT(func(v *ReportConfigFilter) ReportConfigFilter {
		if v != nil {
			return *v
		}
		var ret ReportConfigFilter
		return ret
	}).(ReportConfigFilterOutput)
}

func (o ReportConfigFilterPtrOutput) And() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilter) []ReportConfigFilter {
		if v == nil {
			return nil
		}
		return v.And
	}).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterPtrOutput) Dimensions() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportConfigFilter) *ReportConfigComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterPtrOutput) Or() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilter) []ReportConfigFilter {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterPtrOutput) Tags() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportConfigFilter) *ReportConfigComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(ReportConfigComparisonExpressionPtrOutput)
}

type ReportConfigFilterArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigFilter)(nil)).Elem()
}

func (o ReportConfigFilterArrayOutput) ToReportConfigFilterArrayOutput() ReportConfigFilterArrayOutput {
	return o
}

func (o ReportConfigFilterArrayOutput) ToReportConfigFilterArrayOutputWithContext(ctx context.Context) ReportConfigFilterArrayOutput {
	return o
}

func (o ReportConfigFilterArrayOutput) Index(i pulumi.IntInput) ReportConfigFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigFilter {
		return vs[0].([]ReportConfigFilter)[vs[1].(int)]
	}).(ReportConfigFilterOutput)
}

type ReportConfigFilterResponse struct {
	And        []ReportConfigFilterResponse              `pulumi:"and"`
	Dimensions *ReportConfigComparisonExpressionResponse `pulumi:"dimensions"`
	Or         []ReportConfigFilterResponse              `pulumi:"or"`
	Tags       *ReportConfigComparisonExpressionResponse `pulumi:"tags"`
}

type ReportConfigFilterResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigFilterResponse)(nil)).Elem()
}

func (o ReportConfigFilterResponseOutput) ToReportConfigFilterResponseOutput() ReportConfigFilterResponseOutput {
	return o
}

func (o ReportConfigFilterResponseOutput) ToReportConfigFilterResponseOutputWithContext(ctx context.Context) ReportConfigFilterResponseOutput {
	return o
}

func (o ReportConfigFilterResponseOutput) And() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) []ReportConfigFilterResponse { return v.And }).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponseOutput) Dimensions() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse { return v.Dimensions }).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponseOutput) Or() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) []ReportConfigFilterResponse { return v.Or }).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponseOutput) Tags() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse { return v.Tags }).(ReportConfigComparisonExpressionResponsePtrOutput)
}

type ReportConfigFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigFilterResponse)(nil)).Elem()
}

func (o ReportConfigFilterResponsePtrOutput) ToReportConfigFilterResponsePtrOutput() ReportConfigFilterResponsePtrOutput {
	return o
}

func (o ReportConfigFilterResponsePtrOutput) ToReportConfigFilterResponsePtrOutputWithContext(ctx context.Context) ReportConfigFilterResponsePtrOutput {
	return o
}

func (o ReportConfigFilterResponsePtrOutput) Elem() ReportConfigFilterResponseOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) ReportConfigFilterResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigFilterResponse
		return ret
	}).(ReportConfigFilterResponseOutput)
}

func (o ReportConfigFilterResponsePtrOutput) And() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) []ReportConfigFilterResponse {
		if v == nil {
			return nil
		}
		return v.And
	}).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponsePtrOutput) Dimensions() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponsePtrOutput) Or() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) []ReportConfigFilterResponse {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponsePtrOutput) Tags() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(ReportConfigComparisonExpressionResponsePtrOutput)
}

type ReportConfigFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigFilterResponse)(nil)).Elem()
}

func (o ReportConfigFilterResponseArrayOutput) ToReportConfigFilterResponseArrayOutput() ReportConfigFilterResponseArrayOutput {
	return o
}

func (o ReportConfigFilterResponseArrayOutput) ToReportConfigFilterResponseArrayOutputWithContext(ctx context.Context) ReportConfigFilterResponseArrayOutput {
	return o
}

func (o ReportConfigFilterResponseArrayOutput) Index(i pulumi.IntInput) ReportConfigFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigFilterResponse {
		return vs[0].([]ReportConfigFilterResponse)[vs[1].(int)]
	}).(ReportConfigFilterResponseOutput)
}

type ReportConfigGrouping struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type ReportConfigGroupingInput interface {
	pulumi.Input

	ToReportConfigGroupingOutput() ReportConfigGroupingOutput
	ToReportConfigGroupingOutputWithContext(context.Context) ReportConfigGroupingOutput
}

type ReportConfigGroupingArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ReportConfigGroupingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigGrouping)(nil)).Elem()
}

func (i ReportConfigGroupingArgs) ToReportConfigGroupingOutput() ReportConfigGroupingOutput {
	return i.ToReportConfigGroupingOutputWithContext(context.Background())
}

func (i ReportConfigGroupingArgs) ToReportConfigGroupingOutputWithContext(ctx context.Context) ReportConfigGroupingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigGroupingOutput)
}





type ReportConfigGroupingArrayInput interface {
	pulumi.Input

	ToReportConfigGroupingArrayOutput() ReportConfigGroupingArrayOutput
	ToReportConfigGroupingArrayOutputWithContext(context.Context) ReportConfigGroupingArrayOutput
}

type ReportConfigGroupingArray []ReportConfigGroupingInput

func (ReportConfigGroupingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigGrouping)(nil)).Elem()
}

func (i ReportConfigGroupingArray) ToReportConfigGroupingArrayOutput() ReportConfigGroupingArrayOutput {
	return i.ToReportConfigGroupingArrayOutputWithContext(context.Background())
}

func (i ReportConfigGroupingArray) ToReportConfigGroupingArrayOutputWithContext(ctx context.Context) ReportConfigGroupingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigGroupingArrayOutput)
}

type ReportConfigGroupingOutput struct{ *pulumi.OutputState }

func (ReportConfigGroupingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigGrouping)(nil)).Elem()
}

func (o ReportConfigGroupingOutput) ToReportConfigGroupingOutput() ReportConfigGroupingOutput {
	return o
}

func (o ReportConfigGroupingOutput) ToReportConfigGroupingOutputWithContext(ctx context.Context) ReportConfigGroupingOutput {
	return o
}

func (o ReportConfigGroupingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGrouping) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportConfigGroupingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGrouping) string { return v.Type }).(pulumi.StringOutput)
}

type ReportConfigGroupingArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigGroupingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigGrouping)(nil)).Elem()
}

func (o ReportConfigGroupingArrayOutput) ToReportConfigGroupingArrayOutput() ReportConfigGroupingArrayOutput {
	return o
}

func (o ReportConfigGroupingArrayOutput) ToReportConfigGroupingArrayOutputWithContext(ctx context.Context) ReportConfigGroupingArrayOutput {
	return o
}

func (o ReportConfigGroupingArrayOutput) Index(i pulumi.IntInput) ReportConfigGroupingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigGrouping {
		return vs[0].([]ReportConfigGrouping)[vs[1].(int)]
	}).(ReportConfigGroupingOutput)
}

type ReportConfigGroupingResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

type ReportConfigGroupingResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigGroupingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigGroupingResponse)(nil)).Elem()
}

func (o ReportConfigGroupingResponseOutput) ToReportConfigGroupingResponseOutput() ReportConfigGroupingResponseOutput {
	return o
}

func (o ReportConfigGroupingResponseOutput) ToReportConfigGroupingResponseOutputWithContext(ctx context.Context) ReportConfigGroupingResponseOutput {
	return o
}

func (o ReportConfigGroupingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGroupingResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportConfigGroupingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGroupingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ReportConfigGroupingResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigGroupingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigGroupingResponse)(nil)).Elem()
}

func (o ReportConfigGroupingResponseArrayOutput) ToReportConfigGroupingResponseArrayOutput() ReportConfigGroupingResponseArrayOutput {
	return o
}

func (o ReportConfigGroupingResponseArrayOutput) ToReportConfigGroupingResponseArrayOutputWithContext(ctx context.Context) ReportConfigGroupingResponseArrayOutput {
	return o
}

func (o ReportConfigGroupingResponseArrayOutput) Index(i pulumi.IntInput) ReportConfigGroupingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigGroupingResponse {
		return vs[0].([]ReportConfigGroupingResponse)[vs[1].(int)]
	}).(ReportConfigGroupingResponseOutput)
}

type ReportConfigSorting struct {
	Direction *string `pulumi:"direction"`
	Name      string  `pulumi:"name"`
}





type ReportConfigSortingInput interface {
	pulumi.Input

	ToReportConfigSortingOutput() ReportConfigSortingOutput
	ToReportConfigSortingOutputWithContext(context.Context) ReportConfigSortingOutput
}

type ReportConfigSortingArgs struct {
	Direction pulumi.StringPtrInput `pulumi:"direction"`
	Name      pulumi.StringInput    `pulumi:"name"`
}

func (ReportConfigSortingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigSorting)(nil)).Elem()
}

func (i ReportConfigSortingArgs) ToReportConfigSortingOutput() ReportConfigSortingOutput {
	return i.ToReportConfigSortingOutputWithContext(context.Background())
}

func (i ReportConfigSortingArgs) ToReportConfigSortingOutputWithContext(ctx context.Context) ReportConfigSortingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigSortingOutput)
}





type ReportConfigSortingArrayInput interface {
	pulumi.Input

	ToReportConfigSortingArrayOutput() ReportConfigSortingArrayOutput
	ToReportConfigSortingArrayOutputWithContext(context.Context) ReportConfigSortingArrayOutput
}

type ReportConfigSortingArray []ReportConfigSortingInput

func (ReportConfigSortingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigSorting)(nil)).Elem()
}

func (i ReportConfigSortingArray) ToReportConfigSortingArrayOutput() ReportConfigSortingArrayOutput {
	return i.ToReportConfigSortingArrayOutputWithContext(context.Background())
}

func (i ReportConfigSortingArray) ToReportConfigSortingArrayOutputWithContext(ctx context.Context) ReportConfigSortingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigSortingArrayOutput)
}

type ReportConfigSortingOutput struct{ *pulumi.OutputState }

func (ReportConfigSortingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigSorting)(nil)).Elem()
}

func (o ReportConfigSortingOutput) ToReportConfigSortingOutput() ReportConfigSortingOutput {
	return o
}

func (o ReportConfigSortingOutput) ToReportConfigSortingOutputWithContext(ctx context.Context) ReportConfigSortingOutput {
	return o
}

func (o ReportConfigSortingOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigSorting) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o ReportConfigSortingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigSorting) string { return v.Name }).(pulumi.StringOutput)
}

type ReportConfigSortingArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigSortingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigSorting)(nil)).Elem()
}

func (o ReportConfigSortingArrayOutput) ToReportConfigSortingArrayOutput() ReportConfigSortingArrayOutput {
	return o
}

func (o ReportConfigSortingArrayOutput) ToReportConfigSortingArrayOutputWithContext(ctx context.Context) ReportConfigSortingArrayOutput {
	return o
}

func (o ReportConfigSortingArrayOutput) Index(i pulumi.IntInput) ReportConfigSortingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigSorting {
		return vs[0].([]ReportConfigSorting)[vs[1].(int)]
	}).(ReportConfigSortingOutput)
}

type ReportConfigSortingResponse struct {
	Direction *string `pulumi:"direction"`
	Name      string  `pulumi:"name"`
}

type ReportConfigSortingResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigSortingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigSortingResponse)(nil)).Elem()
}

func (o ReportConfigSortingResponseOutput) ToReportConfigSortingResponseOutput() ReportConfigSortingResponseOutput {
	return o
}

func (o ReportConfigSortingResponseOutput) ToReportConfigSortingResponseOutputWithContext(ctx context.Context) ReportConfigSortingResponseOutput {
	return o
}

func (o ReportConfigSortingResponseOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigSortingResponse) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o ReportConfigSortingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigSortingResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ReportConfigSortingResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigSortingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigSortingResponse)(nil)).Elem()
}

func (o ReportConfigSortingResponseArrayOutput) ToReportConfigSortingResponseArrayOutput() ReportConfigSortingResponseArrayOutput {
	return o
}

func (o ReportConfigSortingResponseArrayOutput) ToReportConfigSortingResponseArrayOutputWithContext(ctx context.Context) ReportConfigSortingResponseArrayOutput {
	return o
}

func (o ReportConfigSortingResponseArrayOutput) Index(i pulumi.IntInput) ReportConfigSortingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigSortingResponse {
		return vs[0].([]ReportConfigSortingResponse)[vs[1].(int)]
	}).(ReportConfigSortingResponseOutput)
}

type ReportConfigTimePeriod struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type ReportConfigTimePeriodInput interface {
	pulumi.Input

	ToReportConfigTimePeriodOutput() ReportConfigTimePeriodOutput
	ToReportConfigTimePeriodOutputWithContext(context.Context) ReportConfigTimePeriodOutput
}

type ReportConfigTimePeriodArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (ReportConfigTimePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigTimePeriod)(nil)).Elem()
}

func (i ReportConfigTimePeriodArgs) ToReportConfigTimePeriodOutput() ReportConfigTimePeriodOutput {
	return i.ToReportConfigTimePeriodOutputWithContext(context.Background())
}

func (i ReportConfigTimePeriodArgs) ToReportConfigTimePeriodOutputWithContext(ctx context.Context) ReportConfigTimePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigTimePeriodOutput)
}

func (i ReportConfigTimePeriodArgs) ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput {
	return i.ToReportConfigTimePeriodPtrOutputWithContext(context.Background())
}

func (i ReportConfigTimePeriodArgs) ToReportConfigTimePeriodPtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigTimePeriodOutput).ToReportConfigTimePeriodPtrOutputWithContext(ctx)
}









type ReportConfigTimePeriodPtrInput interface {
	pulumi.Input

	ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput
	ToReportConfigTimePeriodPtrOutputWithContext(context.Context) ReportConfigTimePeriodPtrOutput
}

type reportConfigTimePeriodPtrType ReportConfigTimePeriodArgs

func ReportConfigTimePeriodPtr(v *ReportConfigTimePeriodArgs) ReportConfigTimePeriodPtrInput {
	return (*reportConfigTimePeriodPtrType)(v)
}

func (*reportConfigTimePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigTimePeriod)(nil)).Elem()
}

func (i *reportConfigTimePeriodPtrType) ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput {
	return i.ToReportConfigTimePeriodPtrOutputWithContext(context.Background())
}

func (i *reportConfigTimePeriodPtrType) ToReportConfigTimePeriodPtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigTimePeriodPtrOutput)
}

type ReportConfigTimePeriodOutput struct{ *pulumi.OutputState }

func (ReportConfigTimePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigTimePeriod)(nil)).Elem()
}

func (o ReportConfigTimePeriodOutput) ToReportConfigTimePeriodOutput() ReportConfigTimePeriodOutput {
	return o
}

func (o ReportConfigTimePeriodOutput) ToReportConfigTimePeriodOutputWithContext(ctx context.Context) ReportConfigTimePeriodOutput {
	return o
}

func (o ReportConfigTimePeriodOutput) ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput {
	return o.ToReportConfigTimePeriodPtrOutputWithContext(context.Background())
}

func (o ReportConfigTimePeriodOutput) ToReportConfigTimePeriodPtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigTimePeriod) *ReportConfigTimePeriod {
		return &v
	}).(ReportConfigTimePeriodPtrOutput)
}

func (o ReportConfigTimePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigTimePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportConfigTimePeriodOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigTimePeriod) string { return v.To }).(pulumi.StringOutput)
}

type ReportConfigTimePeriodPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigTimePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigTimePeriod)(nil)).Elem()
}

func (o ReportConfigTimePeriodPtrOutput) ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput {
	return o
}

func (o ReportConfigTimePeriodPtrOutput) ToReportConfigTimePeriodPtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodPtrOutput {
	return o
}

func (o ReportConfigTimePeriodPtrOutput) Elem() ReportConfigTimePeriodOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriod) ReportConfigTimePeriod {
		if v != nil {
			return *v
		}
		var ret ReportConfigTimePeriod
		return ret
	}).(ReportConfigTimePeriodOutput)
}

func (o ReportConfigTimePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigTimePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type ReportConfigTimePeriodResponse struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}

type ReportConfigTimePeriodResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigTimePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigTimePeriodResponse)(nil)).Elem()
}

func (o ReportConfigTimePeriodResponseOutput) ToReportConfigTimePeriodResponseOutput() ReportConfigTimePeriodResponseOutput {
	return o
}

func (o ReportConfigTimePeriodResponseOutput) ToReportConfigTimePeriodResponseOutputWithContext(ctx context.Context) ReportConfigTimePeriodResponseOutput {
	return o
}

func (o ReportConfigTimePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigTimePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportConfigTimePeriodResponseOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigTimePeriodResponse) string { return v.To }).(pulumi.StringOutput)
}

type ReportConfigTimePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigTimePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigTimePeriodResponse)(nil)).Elem()
}

func (o ReportConfigTimePeriodResponsePtrOutput) ToReportConfigTimePeriodResponsePtrOutput() ReportConfigTimePeriodResponsePtrOutput {
	return o
}

func (o ReportConfigTimePeriodResponsePtrOutput) ToReportConfigTimePeriodResponsePtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodResponsePtrOutput {
	return o
}

func (o ReportConfigTimePeriodResponsePtrOutput) Elem() ReportConfigTimePeriodResponseOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriodResponse) ReportConfigTimePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigTimePeriodResponse
		return ret
	}).(ReportConfigTimePeriodResponseOutput)
}

func (o ReportConfigTimePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigTimePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(KpiPropertiesOutput{})
	pulumi.RegisterOutputType(KpiPropertiesArrayOutput{})
	pulumi.RegisterOutputType(KpiPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KpiPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(PivotPropertiesOutput{})
	pulumi.RegisterOutputType(PivotPropertiesArrayOutput{})
	pulumi.RegisterOutputType(PivotPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PivotPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigAggregationOutput{})
	pulumi.RegisterOutputType(ReportConfigAggregationMapOutput{})
	pulumi.RegisterOutputType(ReportConfigAggregationResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigAggregationResponseMapOutput{})
	pulumi.RegisterOutputType(ReportConfigComparisonExpressionOutput{})
	pulumi.RegisterOutputType(ReportConfigComparisonExpressionPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigComparisonExpressionResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigComparisonExpressionResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetConfigurationOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigGroupingOutput{})
	pulumi.RegisterOutputType(ReportConfigGroupingArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigGroupingResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigGroupingResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigSortingOutput{})
	pulumi.RegisterOutputType(ReportConfigSortingArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigSortingResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigSortingResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodResponsePtrOutput{})
}
