


package v20180531

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

type ReportConfigDefinition struct {
	Dataset    *ReportConfigDataset    `pulumi:"dataset"`
	TimePeriod *ReportConfigTimePeriod `pulumi:"timePeriod"`
	Timeframe  string                  `pulumi:"timeframe"`
	Type       string                  `pulumi:"type"`
}





type ReportConfigDefinitionInput interface {
	pulumi.Input

	ToReportConfigDefinitionOutput() ReportConfigDefinitionOutput
	ToReportConfigDefinitionOutputWithContext(context.Context) ReportConfigDefinitionOutput
}

type ReportConfigDefinitionArgs struct {
	Dataset    ReportConfigDatasetPtrInput    `pulumi:"dataset"`
	TimePeriod ReportConfigTimePeriodPtrInput `pulumi:"timePeriod"`
	Timeframe  pulumi.StringInput             `pulumi:"timeframe"`
	Type       pulumi.StringInput             `pulumi:"type"`
}

func (ReportConfigDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDefinition)(nil)).Elem()
}

func (i ReportConfigDefinitionArgs) ToReportConfigDefinitionOutput() ReportConfigDefinitionOutput {
	return i.ToReportConfigDefinitionOutputWithContext(context.Background())
}

func (i ReportConfigDefinitionArgs) ToReportConfigDefinitionOutputWithContext(ctx context.Context) ReportConfigDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDefinitionOutput)
}

type ReportConfigDefinitionOutput struct{ *pulumi.OutputState }

func (ReportConfigDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDefinition)(nil)).Elem()
}

func (o ReportConfigDefinitionOutput) ToReportConfigDefinitionOutput() ReportConfigDefinitionOutput {
	return o
}

func (o ReportConfigDefinitionOutput) ToReportConfigDefinitionOutputWithContext(ctx context.Context) ReportConfigDefinitionOutput {
	return o
}

func (o ReportConfigDefinitionOutput) Dataset() ReportConfigDatasetPtrOutput {
	return o.ApplyT(func(v ReportConfigDefinition) *ReportConfigDataset { return v.Dataset }).(ReportConfigDatasetPtrOutput)
}

func (o ReportConfigDefinitionOutput) TimePeriod() ReportConfigTimePeriodPtrOutput {
	return o.ApplyT(func(v ReportConfigDefinition) *ReportConfigTimePeriod { return v.TimePeriod }).(ReportConfigTimePeriodPtrOutput)
}

func (o ReportConfigDefinitionOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigDefinition) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ReportConfigDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type ReportConfigDefinitionResponse struct {
	Dataset    *ReportConfigDatasetResponse    `pulumi:"dataset"`
	TimePeriod *ReportConfigTimePeriodResponse `pulumi:"timePeriod"`
	Timeframe  string                          `pulumi:"timeframe"`
	Type       string                          `pulumi:"type"`
}

type ReportConfigDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDefinitionResponse)(nil)).Elem()
}

func (o ReportConfigDefinitionResponseOutput) ToReportConfigDefinitionResponseOutput() ReportConfigDefinitionResponseOutput {
	return o
}

func (o ReportConfigDefinitionResponseOutput) ToReportConfigDefinitionResponseOutputWithContext(ctx context.Context) ReportConfigDefinitionResponseOutput {
	return o
}

func (o ReportConfigDefinitionResponseOutput) Dataset() ReportConfigDatasetResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigDefinitionResponse) *ReportConfigDatasetResponse { return v.Dataset }).(ReportConfigDatasetResponsePtrOutput)
}

func (o ReportConfigDefinitionResponseOutput) TimePeriod() ReportConfigTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigDefinitionResponse) *ReportConfigTimePeriodResponse { return v.TimePeriod }).(ReportConfigTimePeriodResponsePtrOutput)
}

func (o ReportConfigDefinitionResponseOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigDefinitionResponse) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ReportConfigDefinitionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigDefinitionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ReportConfigDeliveryDestination struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}





type ReportConfigDeliveryDestinationInput interface {
	pulumi.Input

	ToReportConfigDeliveryDestinationOutput() ReportConfigDeliveryDestinationOutput
	ToReportConfigDeliveryDestinationOutputWithContext(context.Context) ReportConfigDeliveryDestinationOutput
}

type ReportConfigDeliveryDestinationArgs struct {
	Container      pulumi.StringInput    `pulumi:"container"`
	ResourceId     pulumi.StringInput    `pulumi:"resourceId"`
	RootFolderPath pulumi.StringPtrInput `pulumi:"rootFolderPath"`
}

func (ReportConfigDeliveryDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDeliveryDestination)(nil)).Elem()
}

func (i ReportConfigDeliveryDestinationArgs) ToReportConfigDeliveryDestinationOutput() ReportConfigDeliveryDestinationOutput {
	return i.ToReportConfigDeliveryDestinationOutputWithContext(context.Background())
}

func (i ReportConfigDeliveryDestinationArgs) ToReportConfigDeliveryDestinationOutputWithContext(ctx context.Context) ReportConfigDeliveryDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDeliveryDestinationOutput)
}

type ReportConfigDeliveryDestinationOutput struct{ *pulumi.OutputState }

func (ReportConfigDeliveryDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDeliveryDestination)(nil)).Elem()
}

func (o ReportConfigDeliveryDestinationOutput) ToReportConfigDeliveryDestinationOutput() ReportConfigDeliveryDestinationOutput {
	return o
}

func (o ReportConfigDeliveryDestinationOutput) ToReportConfigDeliveryDestinationOutputWithContext(ctx context.Context) ReportConfigDeliveryDestinationOutput {
	return o
}

func (o ReportConfigDeliveryDestinationOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigDeliveryDestination) string { return v.Container }).(pulumi.StringOutput)
}

func (o ReportConfigDeliveryDestinationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigDeliveryDestination) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ReportConfigDeliveryDestinationOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigDeliveryDestination) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ReportConfigDeliveryDestinationResponse struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}

type ReportConfigDeliveryDestinationResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigDeliveryDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDeliveryDestinationResponse)(nil)).Elem()
}

func (o ReportConfigDeliveryDestinationResponseOutput) ToReportConfigDeliveryDestinationResponseOutput() ReportConfigDeliveryDestinationResponseOutput {
	return o
}

func (o ReportConfigDeliveryDestinationResponseOutput) ToReportConfigDeliveryDestinationResponseOutputWithContext(ctx context.Context) ReportConfigDeliveryDestinationResponseOutput {
	return o
}

func (o ReportConfigDeliveryDestinationResponseOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigDeliveryDestinationResponse) string { return v.Container }).(pulumi.StringOutput)
}

func (o ReportConfigDeliveryDestinationResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigDeliveryDestinationResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ReportConfigDeliveryDestinationResponseOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigDeliveryDestinationResponse) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ReportConfigDeliveryInfo struct {
	Destination ReportConfigDeliveryDestination `pulumi:"destination"`
}





type ReportConfigDeliveryInfoInput interface {
	pulumi.Input

	ToReportConfigDeliveryInfoOutput() ReportConfigDeliveryInfoOutput
	ToReportConfigDeliveryInfoOutputWithContext(context.Context) ReportConfigDeliveryInfoOutput
}

type ReportConfigDeliveryInfoArgs struct {
	Destination ReportConfigDeliveryDestinationInput `pulumi:"destination"`
}

func (ReportConfigDeliveryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDeliveryInfo)(nil)).Elem()
}

func (i ReportConfigDeliveryInfoArgs) ToReportConfigDeliveryInfoOutput() ReportConfigDeliveryInfoOutput {
	return i.ToReportConfigDeliveryInfoOutputWithContext(context.Background())
}

func (i ReportConfigDeliveryInfoArgs) ToReportConfigDeliveryInfoOutputWithContext(ctx context.Context) ReportConfigDeliveryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDeliveryInfoOutput)
}

type ReportConfigDeliveryInfoOutput struct{ *pulumi.OutputState }

func (ReportConfigDeliveryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDeliveryInfo)(nil)).Elem()
}

func (o ReportConfigDeliveryInfoOutput) ToReportConfigDeliveryInfoOutput() ReportConfigDeliveryInfoOutput {
	return o
}

func (o ReportConfigDeliveryInfoOutput) ToReportConfigDeliveryInfoOutputWithContext(ctx context.Context) ReportConfigDeliveryInfoOutput {
	return o
}

func (o ReportConfigDeliveryInfoOutput) Destination() ReportConfigDeliveryDestinationOutput {
	return o.ApplyT(func(v ReportConfigDeliveryInfo) ReportConfigDeliveryDestination { return v.Destination }).(ReportConfigDeliveryDestinationOutput)
}

type ReportConfigDeliveryInfoResponse struct {
	Destination ReportConfigDeliveryDestinationResponse `pulumi:"destination"`
}

type ReportConfigDeliveryInfoResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigDeliveryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDeliveryInfoResponse)(nil)).Elem()
}

func (o ReportConfigDeliveryInfoResponseOutput) ToReportConfigDeliveryInfoResponseOutput() ReportConfigDeliveryInfoResponseOutput {
	return o
}

func (o ReportConfigDeliveryInfoResponseOutput) ToReportConfigDeliveryInfoResponseOutputWithContext(ctx context.Context) ReportConfigDeliveryInfoResponseOutput {
	return o
}

func (o ReportConfigDeliveryInfoResponseOutput) Destination() ReportConfigDeliveryDestinationResponseOutput {
	return o.ApplyT(func(v ReportConfigDeliveryInfoResponse) ReportConfigDeliveryDestinationResponse { return v.Destination }).(ReportConfigDeliveryDestinationResponseOutput)
}

type ReportConfigFilter struct {
	And       []ReportConfigFilter              `pulumi:"and"`
	Dimension *ReportConfigComparisonExpression `pulumi:"dimension"`
	Not       *ReportConfigFilter               `pulumi:"not"`
	Or        []ReportConfigFilter              `pulumi:"or"`
	Tag       *ReportConfigComparisonExpression `pulumi:"tag"`
}





type ReportConfigFilterInput interface {
	pulumi.Input

	ToReportConfigFilterOutput() ReportConfigFilterOutput
	ToReportConfigFilterOutputWithContext(context.Context) ReportConfigFilterOutput
}

type ReportConfigFilterArgs struct {
	And       ReportConfigFilterArrayInput             `pulumi:"and"`
	Dimension ReportConfigComparisonExpressionPtrInput `pulumi:"dimension"`
	Not       ReportConfigFilterPtrInput               `pulumi:"not"`
	Or        ReportConfigFilterArrayInput             `pulumi:"or"`
	Tag       ReportConfigComparisonExpressionPtrInput `pulumi:"tag"`
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

func (o ReportConfigFilterOutput) Dimension() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportConfigFilter) *ReportConfigComparisonExpression { return v.Dimension }).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterOutput) Not() ReportConfigFilterPtrOutput {
	return o.ApplyT(func(v ReportConfigFilter) *ReportConfigFilter { return v.Not }).(ReportConfigFilterPtrOutput)
}

func (o ReportConfigFilterOutput) Or() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v ReportConfigFilter) []ReportConfigFilter { return v.Or }).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterOutput) Tag() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportConfigFilter) *ReportConfigComparisonExpression { return v.Tag }).(ReportConfigComparisonExpressionPtrOutput)
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

func (o ReportConfigFilterPtrOutput) Dimension() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportConfigFilter) *ReportConfigComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Dimension
	}).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterPtrOutput) Not() ReportConfigFilterPtrOutput {
	return o.ApplyT(func(v *ReportConfigFilter) *ReportConfigFilter {
		if v == nil {
			return nil
		}
		return v.Not
	}).(ReportConfigFilterPtrOutput)
}

func (o ReportConfigFilterPtrOutput) Or() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilter) []ReportConfigFilter {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterPtrOutput) Tag() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportConfigFilter) *ReportConfigComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Tag
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
	And       []ReportConfigFilterResponse              `pulumi:"and"`
	Dimension *ReportConfigComparisonExpressionResponse `pulumi:"dimension"`
	Not       *ReportConfigFilterResponse               `pulumi:"not"`
	Or        []ReportConfigFilterResponse              `pulumi:"or"`
	Tag       *ReportConfigComparisonExpressionResponse `pulumi:"tag"`
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

func (o ReportConfigFilterResponseOutput) Dimension() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse { return v.Dimension }).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponseOutput) Not() ReportConfigFilterResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigFilterResponse { return v.Not }).(ReportConfigFilterResponsePtrOutput)
}

func (o ReportConfigFilterResponseOutput) Or() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) []ReportConfigFilterResponse { return v.Or }).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponseOutput) Tag() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse { return v.Tag }).(ReportConfigComparisonExpressionResponsePtrOutput)
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

func (o ReportConfigFilterResponsePtrOutput) Dimension() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Dimension
	}).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponsePtrOutput) Not() ReportConfigFilterResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) *ReportConfigFilterResponse {
		if v == nil {
			return nil
		}
		return v.Not
	}).(ReportConfigFilterResponsePtrOutput)
}

func (o ReportConfigFilterResponsePtrOutput) Or() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) []ReportConfigFilterResponse {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponsePtrOutput) Tag() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Tag
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
	ColumnType string `pulumi:"columnType"`
	Name       string `pulumi:"name"`
}





type ReportConfigGroupingInput interface {
	pulumi.Input

	ToReportConfigGroupingOutput() ReportConfigGroupingOutput
	ToReportConfigGroupingOutputWithContext(context.Context) ReportConfigGroupingOutput
}

type ReportConfigGroupingArgs struct {
	ColumnType pulumi.StringInput `pulumi:"columnType"`
	Name       pulumi.StringInput `pulumi:"name"`
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

func (o ReportConfigGroupingOutput) ColumnType() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGrouping) string { return v.ColumnType }).(pulumi.StringOutput)
}

func (o ReportConfigGroupingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGrouping) string { return v.Name }).(pulumi.StringOutput)
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
	ColumnType string `pulumi:"columnType"`
	Name       string `pulumi:"name"`
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

func (o ReportConfigGroupingResponseOutput) ColumnType() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGroupingResponse) string { return v.ColumnType }).(pulumi.StringOutput)
}

func (o ReportConfigGroupingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGroupingResponse) string { return v.Name }).(pulumi.StringOutput)
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

type ReportConfigRecurrencePeriod struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}





type ReportConfigRecurrencePeriodInput interface {
	pulumi.Input

	ToReportConfigRecurrencePeriodOutput() ReportConfigRecurrencePeriodOutput
	ToReportConfigRecurrencePeriodOutputWithContext(context.Context) ReportConfigRecurrencePeriodOutput
}

type ReportConfigRecurrencePeriodArgs struct {
	From pulumi.StringInput    `pulumi:"from"`
	To   pulumi.StringPtrInput `pulumi:"to"`
}

func (ReportConfigRecurrencePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigRecurrencePeriod)(nil)).Elem()
}

func (i ReportConfigRecurrencePeriodArgs) ToReportConfigRecurrencePeriodOutput() ReportConfigRecurrencePeriodOutput {
	return i.ToReportConfigRecurrencePeriodOutputWithContext(context.Background())
}

func (i ReportConfigRecurrencePeriodArgs) ToReportConfigRecurrencePeriodOutputWithContext(ctx context.Context) ReportConfigRecurrencePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigRecurrencePeriodOutput)
}

func (i ReportConfigRecurrencePeriodArgs) ToReportConfigRecurrencePeriodPtrOutput() ReportConfigRecurrencePeriodPtrOutput {
	return i.ToReportConfigRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i ReportConfigRecurrencePeriodArgs) ToReportConfigRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportConfigRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigRecurrencePeriodOutput).ToReportConfigRecurrencePeriodPtrOutputWithContext(ctx)
}









type ReportConfigRecurrencePeriodPtrInput interface {
	pulumi.Input

	ToReportConfigRecurrencePeriodPtrOutput() ReportConfigRecurrencePeriodPtrOutput
	ToReportConfigRecurrencePeriodPtrOutputWithContext(context.Context) ReportConfigRecurrencePeriodPtrOutput
}

type reportConfigRecurrencePeriodPtrType ReportConfigRecurrencePeriodArgs

func ReportConfigRecurrencePeriodPtr(v *ReportConfigRecurrencePeriodArgs) ReportConfigRecurrencePeriodPtrInput {
	return (*reportConfigRecurrencePeriodPtrType)(v)
}

func (*reportConfigRecurrencePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigRecurrencePeriod)(nil)).Elem()
}

func (i *reportConfigRecurrencePeriodPtrType) ToReportConfigRecurrencePeriodPtrOutput() ReportConfigRecurrencePeriodPtrOutput {
	return i.ToReportConfigRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i *reportConfigRecurrencePeriodPtrType) ToReportConfigRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportConfigRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigRecurrencePeriodPtrOutput)
}

type ReportConfigRecurrencePeriodOutput struct{ *pulumi.OutputState }

func (ReportConfigRecurrencePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigRecurrencePeriod)(nil)).Elem()
}

func (o ReportConfigRecurrencePeriodOutput) ToReportConfigRecurrencePeriodOutput() ReportConfigRecurrencePeriodOutput {
	return o
}

func (o ReportConfigRecurrencePeriodOutput) ToReportConfigRecurrencePeriodOutputWithContext(ctx context.Context) ReportConfigRecurrencePeriodOutput {
	return o
}

func (o ReportConfigRecurrencePeriodOutput) ToReportConfigRecurrencePeriodPtrOutput() ReportConfigRecurrencePeriodPtrOutput {
	return o.ToReportConfigRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (o ReportConfigRecurrencePeriodOutput) ToReportConfigRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportConfigRecurrencePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigRecurrencePeriod) *ReportConfigRecurrencePeriod {
		return &v
	}).(ReportConfigRecurrencePeriodPtrOutput)
}

func (o ReportConfigRecurrencePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigRecurrencePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportConfigRecurrencePeriodOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigRecurrencePeriod) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ReportConfigRecurrencePeriodPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigRecurrencePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigRecurrencePeriod)(nil)).Elem()
}

func (o ReportConfigRecurrencePeriodPtrOutput) ToReportConfigRecurrencePeriodPtrOutput() ReportConfigRecurrencePeriodPtrOutput {
	return o
}

func (o ReportConfigRecurrencePeriodPtrOutput) ToReportConfigRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportConfigRecurrencePeriodPtrOutput {
	return o
}

func (o ReportConfigRecurrencePeriodPtrOutput) Elem() ReportConfigRecurrencePeriodOutput {
	return o.ApplyT(func(v *ReportConfigRecurrencePeriod) ReportConfigRecurrencePeriod {
		if v != nil {
			return *v
		}
		var ret ReportConfigRecurrencePeriod
		return ret
	}).(ReportConfigRecurrencePeriodOutput)
}

func (o ReportConfigRecurrencePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigRecurrencePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ReportConfigRecurrencePeriodResponse struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}

type ReportConfigRecurrencePeriodResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigRecurrencePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigRecurrencePeriodResponse)(nil)).Elem()
}

func (o ReportConfigRecurrencePeriodResponseOutput) ToReportConfigRecurrencePeriodResponseOutput() ReportConfigRecurrencePeriodResponseOutput {
	return o
}

func (o ReportConfigRecurrencePeriodResponseOutput) ToReportConfigRecurrencePeriodResponseOutputWithContext(ctx context.Context) ReportConfigRecurrencePeriodResponseOutput {
	return o
}

func (o ReportConfigRecurrencePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigRecurrencePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportConfigRecurrencePeriodResponseOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigRecurrencePeriodResponse) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ReportConfigRecurrencePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigRecurrencePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigRecurrencePeriodResponse)(nil)).Elem()
}

func (o ReportConfigRecurrencePeriodResponsePtrOutput) ToReportConfigRecurrencePeriodResponsePtrOutput() ReportConfigRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ReportConfigRecurrencePeriodResponsePtrOutput) ToReportConfigRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ReportConfigRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ReportConfigRecurrencePeriodResponsePtrOutput) Elem() ReportConfigRecurrencePeriodResponseOutput {
	return o.ApplyT(func(v *ReportConfigRecurrencePeriodResponse) ReportConfigRecurrencePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigRecurrencePeriodResponse
		return ret
	}).(ReportConfigRecurrencePeriodResponseOutput)
}

func (o ReportConfigRecurrencePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigRecurrencePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ReportConfigSchedule struct {
	Recurrence       string                       `pulumi:"recurrence"`
	RecurrencePeriod ReportConfigRecurrencePeriod `pulumi:"recurrencePeriod"`
	Status           *string                      `pulumi:"status"`
}





type ReportConfigScheduleInput interface {
	pulumi.Input

	ToReportConfigScheduleOutput() ReportConfigScheduleOutput
	ToReportConfigScheduleOutputWithContext(context.Context) ReportConfigScheduleOutput
}

type ReportConfigScheduleArgs struct {
	Recurrence       pulumi.StringInput                `pulumi:"recurrence"`
	RecurrencePeriod ReportConfigRecurrencePeriodInput `pulumi:"recurrencePeriod"`
	Status           pulumi.StringPtrInput             `pulumi:"status"`
}

func (ReportConfigScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigSchedule)(nil)).Elem()
}

func (i ReportConfigScheduleArgs) ToReportConfigScheduleOutput() ReportConfigScheduleOutput {
	return i.ToReportConfigScheduleOutputWithContext(context.Background())
}

func (i ReportConfigScheduleArgs) ToReportConfigScheduleOutputWithContext(ctx context.Context) ReportConfigScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigScheduleOutput)
}

func (i ReportConfigScheduleArgs) ToReportConfigSchedulePtrOutput() ReportConfigSchedulePtrOutput {
	return i.ToReportConfigSchedulePtrOutputWithContext(context.Background())
}

func (i ReportConfigScheduleArgs) ToReportConfigSchedulePtrOutputWithContext(ctx context.Context) ReportConfigSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigScheduleOutput).ToReportConfigSchedulePtrOutputWithContext(ctx)
}









type ReportConfigSchedulePtrInput interface {
	pulumi.Input

	ToReportConfigSchedulePtrOutput() ReportConfigSchedulePtrOutput
	ToReportConfigSchedulePtrOutputWithContext(context.Context) ReportConfigSchedulePtrOutput
}

type reportConfigSchedulePtrType ReportConfigScheduleArgs

func ReportConfigSchedulePtr(v *ReportConfigScheduleArgs) ReportConfigSchedulePtrInput {
	return (*reportConfigSchedulePtrType)(v)
}

func (*reportConfigSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigSchedule)(nil)).Elem()
}

func (i *reportConfigSchedulePtrType) ToReportConfigSchedulePtrOutput() ReportConfigSchedulePtrOutput {
	return i.ToReportConfigSchedulePtrOutputWithContext(context.Background())
}

func (i *reportConfigSchedulePtrType) ToReportConfigSchedulePtrOutputWithContext(ctx context.Context) ReportConfigSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigSchedulePtrOutput)
}

type ReportConfigScheduleOutput struct{ *pulumi.OutputState }

func (ReportConfigScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigSchedule)(nil)).Elem()
}

func (o ReportConfigScheduleOutput) ToReportConfigScheduleOutput() ReportConfigScheduleOutput {
	return o
}

func (o ReportConfigScheduleOutput) ToReportConfigScheduleOutputWithContext(ctx context.Context) ReportConfigScheduleOutput {
	return o
}

func (o ReportConfigScheduleOutput) ToReportConfigSchedulePtrOutput() ReportConfigSchedulePtrOutput {
	return o.ToReportConfigSchedulePtrOutputWithContext(context.Background())
}

func (o ReportConfigScheduleOutput) ToReportConfigSchedulePtrOutputWithContext(ctx context.Context) ReportConfigSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigSchedule) *ReportConfigSchedule {
		return &v
	}).(ReportConfigSchedulePtrOutput)
}

func (o ReportConfigScheduleOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigSchedule) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ReportConfigScheduleOutput) RecurrencePeriod() ReportConfigRecurrencePeriodOutput {
	return o.ApplyT(func(v ReportConfigSchedule) ReportConfigRecurrencePeriod { return v.RecurrencePeriod }).(ReportConfigRecurrencePeriodOutput)
}

func (o ReportConfigScheduleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigSchedule) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ReportConfigSchedulePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigSchedule)(nil)).Elem()
}

func (o ReportConfigSchedulePtrOutput) ToReportConfigSchedulePtrOutput() ReportConfigSchedulePtrOutput {
	return o
}

func (o ReportConfigSchedulePtrOutput) ToReportConfigSchedulePtrOutputWithContext(ctx context.Context) ReportConfigSchedulePtrOutput {
	return o
}

func (o ReportConfigSchedulePtrOutput) Elem() ReportConfigScheduleOutput {
	return o.ApplyT(func(v *ReportConfigSchedule) ReportConfigSchedule {
		if v != nil {
			return *v
		}
		var ret ReportConfigSchedule
		return ret
	}).(ReportConfigScheduleOutput)
}

func (o ReportConfigSchedulePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigSchedule) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigSchedulePtrOutput) RecurrencePeriod() ReportConfigRecurrencePeriodPtrOutput {
	return o.ApplyT(func(v *ReportConfigSchedule) *ReportConfigRecurrencePeriod {
		if v == nil {
			return nil
		}
		return &v.RecurrencePeriod
	}).(ReportConfigRecurrencePeriodPtrOutput)
}

func (o ReportConfigSchedulePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigSchedule) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ReportConfigScheduleResponse struct {
	Recurrence       string                               `pulumi:"recurrence"`
	RecurrencePeriod ReportConfigRecurrencePeriodResponse `pulumi:"recurrencePeriod"`
	Status           *string                              `pulumi:"status"`
}

type ReportConfigScheduleResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigScheduleResponse)(nil)).Elem()
}

func (o ReportConfigScheduleResponseOutput) ToReportConfigScheduleResponseOutput() ReportConfigScheduleResponseOutput {
	return o
}

func (o ReportConfigScheduleResponseOutput) ToReportConfigScheduleResponseOutputWithContext(ctx context.Context) ReportConfigScheduleResponseOutput {
	return o
}

func (o ReportConfigScheduleResponseOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigScheduleResponse) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ReportConfigScheduleResponseOutput) RecurrencePeriod() ReportConfigRecurrencePeriodResponseOutput {
	return o.ApplyT(func(v ReportConfigScheduleResponse) ReportConfigRecurrencePeriodResponse { return v.RecurrencePeriod }).(ReportConfigRecurrencePeriodResponseOutput)
}

func (o ReportConfigScheduleResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigScheduleResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ReportConfigScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigScheduleResponse)(nil)).Elem()
}

func (o ReportConfigScheduleResponsePtrOutput) ToReportConfigScheduleResponsePtrOutput() ReportConfigScheduleResponsePtrOutput {
	return o
}

func (o ReportConfigScheduleResponsePtrOutput) ToReportConfigScheduleResponsePtrOutputWithContext(ctx context.Context) ReportConfigScheduleResponsePtrOutput {
	return o
}

func (o ReportConfigScheduleResponsePtrOutput) Elem() ReportConfigScheduleResponseOutput {
	return o.ApplyT(func(v *ReportConfigScheduleResponse) ReportConfigScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigScheduleResponse
		return ret
	}).(ReportConfigScheduleResponseOutput)
}

func (o ReportConfigScheduleResponsePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigScheduleResponsePtrOutput) RecurrencePeriod() ReportConfigRecurrencePeriodResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigScheduleResponse) *ReportConfigRecurrencePeriodResponse {
		if v == nil {
			return nil
		}
		return &v.RecurrencePeriod
	}).(ReportConfigRecurrencePeriodResponsePtrOutput)
}

func (o ReportConfigScheduleResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(ReportConfigDefinitionOutput{})
	pulumi.RegisterOutputType(ReportConfigDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigDeliveryDestinationOutput{})
	pulumi.RegisterOutputType(ReportConfigDeliveryDestinationResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigDeliveryInfoOutput{})
	pulumi.RegisterOutputType(ReportConfigDeliveryInfoResponseOutput{})
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
	pulumi.RegisterOutputType(ReportConfigRecurrencePeriodOutput{})
	pulumi.RegisterOutputType(ReportConfigRecurrencePeriodPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigRecurrencePeriodResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigRecurrencePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigScheduleOutput{})
	pulumi.RegisterOutputType(ReportConfigSchedulePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigScheduleResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodResponsePtrOutput{})
}
