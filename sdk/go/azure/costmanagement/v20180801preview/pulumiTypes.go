


package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectorCollectionErrorInfoResponse struct {
	ErrorCode      string `pulumi:"errorCode"`
	ErrorMessage   string `pulumi:"errorMessage"`
	ErrorStartTime string `pulumi:"errorStartTime"`
}

type ConnectorCollectionErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (ConnectorCollectionErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorCollectionErrorInfoResponse)(nil)).Elem()
}

func (o ConnectorCollectionErrorInfoResponseOutput) ToConnectorCollectionErrorInfoResponseOutput() ConnectorCollectionErrorInfoResponseOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponseOutput) ToConnectorCollectionErrorInfoResponseOutputWithContext(ctx context.Context) ConnectorCollectionErrorInfoResponseOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorCode() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorCode }).(pulumi.StringOutput)
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorStartTime }).(pulumi.StringOutput)
}

type ConnectorCollectionErrorInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectorCollectionErrorInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorCollectionErrorInfoResponse)(nil)).Elem()
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ToConnectorCollectionErrorInfoResponsePtrOutput() ConnectorCollectionErrorInfoResponsePtrOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ToConnectorCollectionErrorInfoResponsePtrOutputWithContext(ctx context.Context) ConnectorCollectionErrorInfoResponsePtrOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) Elem() ConnectorCollectionErrorInfoResponseOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) ConnectorCollectionErrorInfoResponse {
		if v != nil {
			return *v
		}
		var ret ConnectorCollectionErrorInfoResponse
		return ret
	}).(ConnectorCollectionErrorInfoResponseOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorCode
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorStartTime
	}).(pulumi.StringPtrOutput)
}

type ConnectorCollectionInfoResponse struct {
	Error             *ConnectorCollectionErrorInfoResponse `pulumi:"error"`
	LastRun           string                                `pulumi:"lastRun"`
	LastUpdated       string                                `pulumi:"lastUpdated"`
	SourceLastUpdated string                                `pulumi:"sourceLastUpdated"`
}

type ConnectorCollectionInfoResponseOutput struct{ *pulumi.OutputState }

func (ConnectorCollectionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorCollectionInfoResponse)(nil)).Elem()
}

func (o ConnectorCollectionInfoResponseOutput) ToConnectorCollectionInfoResponseOutput() ConnectorCollectionInfoResponseOutput {
	return o
}

func (o ConnectorCollectionInfoResponseOutput) ToConnectorCollectionInfoResponseOutputWithContext(ctx context.Context) ConnectorCollectionInfoResponseOutput {
	return o
}

func (o ConnectorCollectionInfoResponseOutput) Error() ConnectorCollectionErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) *ConnectorCollectionErrorInfoResponse { return v.Error }).(ConnectorCollectionErrorInfoResponsePtrOutput)
}

func (o ConnectorCollectionInfoResponseOutput) LastRun() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) string { return v.LastRun }).(pulumi.StringOutput)
}

func (o ConnectorCollectionInfoResponseOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) string { return v.LastUpdated }).(pulumi.StringOutput)
}

func (o ConnectorCollectionInfoResponseOutput) SourceLastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) string { return v.SourceLastUpdated }).(pulumi.StringOutput)
}

type ReportAggregation struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}





type ReportAggregationInput interface {
	pulumi.Input

	ToReportAggregationOutput() ReportAggregationOutput
	ToReportAggregationOutputWithContext(context.Context) ReportAggregationOutput
}

type ReportAggregationArgs struct {
	Function pulumi.StringInput `pulumi:"function"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (ReportAggregationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportAggregation)(nil)).Elem()
}

func (i ReportAggregationArgs) ToReportAggregationOutput() ReportAggregationOutput {
	return i.ToReportAggregationOutputWithContext(context.Background())
}

func (i ReportAggregationArgs) ToReportAggregationOutputWithContext(ctx context.Context) ReportAggregationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportAggregationOutput)
}





type ReportAggregationMapInput interface {
	pulumi.Input

	ToReportAggregationMapOutput() ReportAggregationMapOutput
	ToReportAggregationMapOutputWithContext(context.Context) ReportAggregationMapOutput
}

type ReportAggregationMap map[string]ReportAggregationInput

func (ReportAggregationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportAggregation)(nil)).Elem()
}

func (i ReportAggregationMap) ToReportAggregationMapOutput() ReportAggregationMapOutput {
	return i.ToReportAggregationMapOutputWithContext(context.Background())
}

func (i ReportAggregationMap) ToReportAggregationMapOutputWithContext(ctx context.Context) ReportAggregationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportAggregationMapOutput)
}

type ReportAggregationOutput struct{ *pulumi.OutputState }

func (ReportAggregationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportAggregation)(nil)).Elem()
}

func (o ReportAggregationOutput) ToReportAggregationOutput() ReportAggregationOutput {
	return o
}

func (o ReportAggregationOutput) ToReportAggregationOutputWithContext(ctx context.Context) ReportAggregationOutput {
	return o
}

func (o ReportAggregationOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v ReportAggregation) string { return v.Function }).(pulumi.StringOutput)
}

func (o ReportAggregationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportAggregation) string { return v.Name }).(pulumi.StringOutput)
}

type ReportAggregationMapOutput struct{ *pulumi.OutputState }

func (ReportAggregationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportAggregation)(nil)).Elem()
}

func (o ReportAggregationMapOutput) ToReportAggregationMapOutput() ReportAggregationMapOutput {
	return o
}

func (o ReportAggregationMapOutput) ToReportAggregationMapOutputWithContext(ctx context.Context) ReportAggregationMapOutput {
	return o
}

func (o ReportAggregationMapOutput) MapIndex(k pulumi.StringInput) ReportAggregationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReportAggregation {
		return vs[0].(map[string]ReportAggregation)[vs[1].(string)]
	}).(ReportAggregationOutput)
}

type ReportAggregationResponse struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}

type ReportAggregationResponseOutput struct{ *pulumi.OutputState }

func (ReportAggregationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportAggregationResponse)(nil)).Elem()
}

func (o ReportAggregationResponseOutput) ToReportAggregationResponseOutput() ReportAggregationResponseOutput {
	return o
}

func (o ReportAggregationResponseOutput) ToReportAggregationResponseOutputWithContext(ctx context.Context) ReportAggregationResponseOutput {
	return o
}

func (o ReportAggregationResponseOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v ReportAggregationResponse) string { return v.Function }).(pulumi.StringOutput)
}

func (o ReportAggregationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportAggregationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ReportAggregationResponseMapOutput struct{ *pulumi.OutputState }

func (ReportAggregationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportAggregationResponse)(nil)).Elem()
}

func (o ReportAggregationResponseMapOutput) ToReportAggregationResponseMapOutput() ReportAggregationResponseMapOutput {
	return o
}

func (o ReportAggregationResponseMapOutput) ToReportAggregationResponseMapOutputWithContext(ctx context.Context) ReportAggregationResponseMapOutput {
	return o
}

func (o ReportAggregationResponseMapOutput) MapIndex(k pulumi.StringInput) ReportAggregationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReportAggregationResponse {
		return vs[0].(map[string]ReportAggregationResponse)[vs[1].(string)]
	}).(ReportAggregationResponseOutput)
}

type ReportComparisonExpression struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type ReportComparisonExpressionInput interface {
	pulumi.Input

	ToReportComparisonExpressionOutput() ReportComparisonExpressionOutput
	ToReportComparisonExpressionOutputWithContext(context.Context) ReportComparisonExpressionOutput
}

type ReportComparisonExpressionArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (ReportComparisonExpressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportComparisonExpression)(nil)).Elem()
}

func (i ReportComparisonExpressionArgs) ToReportComparisonExpressionOutput() ReportComparisonExpressionOutput {
	return i.ToReportComparisonExpressionOutputWithContext(context.Background())
}

func (i ReportComparisonExpressionArgs) ToReportComparisonExpressionOutputWithContext(ctx context.Context) ReportComparisonExpressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportComparisonExpressionOutput)
}

func (i ReportComparisonExpressionArgs) ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput {
	return i.ToReportComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i ReportComparisonExpressionArgs) ToReportComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportComparisonExpressionOutput).ToReportComparisonExpressionPtrOutputWithContext(ctx)
}









type ReportComparisonExpressionPtrInput interface {
	pulumi.Input

	ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput
	ToReportComparisonExpressionPtrOutputWithContext(context.Context) ReportComparisonExpressionPtrOutput
}

type reportComparisonExpressionPtrType ReportComparisonExpressionArgs

func ReportComparisonExpressionPtr(v *ReportComparisonExpressionArgs) ReportComparisonExpressionPtrInput {
	return (*reportComparisonExpressionPtrType)(v)
}

func (*reportComparisonExpressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportComparisonExpression)(nil)).Elem()
}

func (i *reportComparisonExpressionPtrType) ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput {
	return i.ToReportComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i *reportComparisonExpressionPtrType) ToReportComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportComparisonExpressionPtrOutput)
}

type ReportComparisonExpressionOutput struct{ *pulumi.OutputState }

func (ReportComparisonExpressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportComparisonExpression)(nil)).Elem()
}

func (o ReportComparisonExpressionOutput) ToReportComparisonExpressionOutput() ReportComparisonExpressionOutput {
	return o
}

func (o ReportComparisonExpressionOutput) ToReportComparisonExpressionOutputWithContext(ctx context.Context) ReportComparisonExpressionOutput {
	return o
}

func (o ReportComparisonExpressionOutput) ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput {
	return o.ToReportComparisonExpressionPtrOutputWithContext(context.Background())
}

func (o ReportComparisonExpressionOutput) ToReportComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportComparisonExpressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportComparisonExpression) *ReportComparisonExpression {
		return &v
	}).(ReportComparisonExpressionPtrOutput)
}

func (o ReportComparisonExpressionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportComparisonExpression) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportComparisonExpressionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ReportComparisonExpression) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ReportComparisonExpressionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportComparisonExpression) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ReportComparisonExpressionPtrOutput struct{ *pulumi.OutputState }

func (ReportComparisonExpressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportComparisonExpression)(nil)).Elem()
}

func (o ReportComparisonExpressionPtrOutput) ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput {
	return o
}

func (o ReportComparisonExpressionPtrOutput) ToReportComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportComparisonExpressionPtrOutput {
	return o
}

func (o ReportComparisonExpressionPtrOutput) Elem() ReportComparisonExpressionOutput {
	return o.ApplyT(func(v *ReportComparisonExpression) ReportComparisonExpression {
		if v != nil {
			return *v
		}
		var ret ReportComparisonExpression
		return ret
	}).(ReportComparisonExpressionOutput)
}

func (o ReportComparisonExpressionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ReportComparisonExpressionPtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ReportComparisonExpressionPtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportComparisonExpression) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type ReportComparisonExpressionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type ReportComparisonExpressionResponseOutput struct{ *pulumi.OutputState }

func (ReportComparisonExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportComparisonExpressionResponse)(nil)).Elem()
}

func (o ReportComparisonExpressionResponseOutput) ToReportComparisonExpressionResponseOutput() ReportComparisonExpressionResponseOutput {
	return o
}

func (o ReportComparisonExpressionResponseOutput) ToReportComparisonExpressionResponseOutputWithContext(ctx context.Context) ReportComparisonExpressionResponseOutput {
	return o
}

func (o ReportComparisonExpressionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportComparisonExpressionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportComparisonExpressionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ReportComparisonExpressionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ReportComparisonExpressionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportComparisonExpressionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ReportComparisonExpressionResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportComparisonExpressionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportComparisonExpressionResponse)(nil)).Elem()
}

func (o ReportComparisonExpressionResponsePtrOutput) ToReportComparisonExpressionResponsePtrOutput() ReportComparisonExpressionResponsePtrOutput {
	return o
}

func (o ReportComparisonExpressionResponsePtrOutput) ToReportComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) ReportComparisonExpressionResponsePtrOutput {
	return o
}

func (o ReportComparisonExpressionResponsePtrOutput) Elem() ReportComparisonExpressionResponseOutput {
	return o.ApplyT(func(v *ReportComparisonExpressionResponse) ReportComparisonExpressionResponse {
		if v != nil {
			return *v
		}
		var ret ReportComparisonExpressionResponse
		return ret
	}).(ReportComparisonExpressionResponseOutput)
}

func (o ReportComparisonExpressionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ReportComparisonExpressionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ReportComparisonExpressionResponsePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportComparisonExpressionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type ReportDataset struct {
	Aggregation   map[string]ReportAggregation `pulumi:"aggregation"`
	Configuration *ReportDatasetConfiguration  `pulumi:"configuration"`
	Filter        *ReportFilter                `pulumi:"filter"`
	Granularity   *string                      `pulumi:"granularity"`
	Grouping      []ReportGrouping             `pulumi:"grouping"`
}





type ReportDatasetInput interface {
	pulumi.Input

	ToReportDatasetOutput() ReportDatasetOutput
	ToReportDatasetOutputWithContext(context.Context) ReportDatasetOutput
}

type ReportDatasetArgs struct {
	Aggregation   ReportAggregationMapInput          `pulumi:"aggregation"`
	Configuration ReportDatasetConfigurationPtrInput `pulumi:"configuration"`
	Filter        ReportFilterPtrInput               `pulumi:"filter"`
	Granularity   pulumi.StringPtrInput              `pulumi:"granularity"`
	Grouping      ReportGroupingArrayInput           `pulumi:"grouping"`
}

func (ReportDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDataset)(nil)).Elem()
}

func (i ReportDatasetArgs) ToReportDatasetOutput() ReportDatasetOutput {
	return i.ToReportDatasetOutputWithContext(context.Background())
}

func (i ReportDatasetArgs) ToReportDatasetOutputWithContext(ctx context.Context) ReportDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetOutput)
}

func (i ReportDatasetArgs) ToReportDatasetPtrOutput() ReportDatasetPtrOutput {
	return i.ToReportDatasetPtrOutputWithContext(context.Background())
}

func (i ReportDatasetArgs) ToReportDatasetPtrOutputWithContext(ctx context.Context) ReportDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetOutput).ToReportDatasetPtrOutputWithContext(ctx)
}









type ReportDatasetPtrInput interface {
	pulumi.Input

	ToReportDatasetPtrOutput() ReportDatasetPtrOutput
	ToReportDatasetPtrOutputWithContext(context.Context) ReportDatasetPtrOutput
}

type reportDatasetPtrType ReportDatasetArgs

func ReportDatasetPtr(v *ReportDatasetArgs) ReportDatasetPtrInput {
	return (*reportDatasetPtrType)(v)
}

func (*reportDatasetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDataset)(nil)).Elem()
}

func (i *reportDatasetPtrType) ToReportDatasetPtrOutput() ReportDatasetPtrOutput {
	return i.ToReportDatasetPtrOutputWithContext(context.Background())
}

func (i *reportDatasetPtrType) ToReportDatasetPtrOutputWithContext(ctx context.Context) ReportDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetPtrOutput)
}

type ReportDatasetOutput struct{ *pulumi.OutputState }

func (ReportDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDataset)(nil)).Elem()
}

func (o ReportDatasetOutput) ToReportDatasetOutput() ReportDatasetOutput {
	return o
}

func (o ReportDatasetOutput) ToReportDatasetOutputWithContext(ctx context.Context) ReportDatasetOutput {
	return o
}

func (o ReportDatasetOutput) ToReportDatasetPtrOutput() ReportDatasetPtrOutput {
	return o.ToReportDatasetPtrOutputWithContext(context.Background())
}

func (o ReportDatasetOutput) ToReportDatasetPtrOutputWithContext(ctx context.Context) ReportDatasetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportDataset) *ReportDataset {
		return &v
	}).(ReportDatasetPtrOutput)
}

func (o ReportDatasetOutput) Aggregation() ReportAggregationMapOutput {
	return o.ApplyT(func(v ReportDataset) map[string]ReportAggregation { return v.Aggregation }).(ReportAggregationMapOutput)
}

func (o ReportDatasetOutput) Configuration() ReportDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v ReportDataset) *ReportDatasetConfiguration { return v.Configuration }).(ReportDatasetConfigurationPtrOutput)
}

func (o ReportDatasetOutput) Filter() ReportFilterPtrOutput {
	return o.ApplyT(func(v ReportDataset) *ReportFilter { return v.Filter }).(ReportFilterPtrOutput)
}

func (o ReportDatasetOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportDataset) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o ReportDatasetOutput) Grouping() ReportGroupingArrayOutput {
	return o.ApplyT(func(v ReportDataset) []ReportGrouping { return v.Grouping }).(ReportGroupingArrayOutput)
}

type ReportDatasetPtrOutput struct{ *pulumi.OutputState }

func (ReportDatasetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDataset)(nil)).Elem()
}

func (o ReportDatasetPtrOutput) ToReportDatasetPtrOutput() ReportDatasetPtrOutput {
	return o
}

func (o ReportDatasetPtrOutput) ToReportDatasetPtrOutputWithContext(ctx context.Context) ReportDatasetPtrOutput {
	return o
}

func (o ReportDatasetPtrOutput) Elem() ReportDatasetOutput {
	return o.ApplyT(func(v *ReportDataset) ReportDataset {
		if v != nil {
			return *v
		}
		var ret ReportDataset
		return ret
	}).(ReportDatasetOutput)
}

func (o ReportDatasetPtrOutput) Aggregation() ReportAggregationMapOutput {
	return o.ApplyT(func(v *ReportDataset) map[string]ReportAggregation {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(ReportAggregationMapOutput)
}

func (o ReportDatasetPtrOutput) Configuration() ReportDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v *ReportDataset) *ReportDatasetConfiguration {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ReportDatasetConfigurationPtrOutput)
}

func (o ReportDatasetPtrOutput) Filter() ReportFilterPtrOutput {
	return o.ApplyT(func(v *ReportDataset) *ReportFilter {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(ReportFilterPtrOutput)
}

func (o ReportDatasetPtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportDataset) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o ReportDatasetPtrOutput) Grouping() ReportGroupingArrayOutput {
	return o.ApplyT(func(v *ReportDataset) []ReportGrouping {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(ReportGroupingArrayOutput)
}

type ReportDatasetConfiguration struct {
	Columns []string `pulumi:"columns"`
}





type ReportDatasetConfigurationInput interface {
	pulumi.Input

	ToReportDatasetConfigurationOutput() ReportDatasetConfigurationOutput
	ToReportDatasetConfigurationOutputWithContext(context.Context) ReportDatasetConfigurationOutput
}

type ReportDatasetConfigurationArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (ReportDatasetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDatasetConfiguration)(nil)).Elem()
}

func (i ReportDatasetConfigurationArgs) ToReportDatasetConfigurationOutput() ReportDatasetConfigurationOutput {
	return i.ToReportDatasetConfigurationOutputWithContext(context.Background())
}

func (i ReportDatasetConfigurationArgs) ToReportDatasetConfigurationOutputWithContext(ctx context.Context) ReportDatasetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetConfigurationOutput)
}

func (i ReportDatasetConfigurationArgs) ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput {
	return i.ToReportDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i ReportDatasetConfigurationArgs) ToReportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetConfigurationOutput).ToReportDatasetConfigurationPtrOutputWithContext(ctx)
}









type ReportDatasetConfigurationPtrInput interface {
	pulumi.Input

	ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput
	ToReportDatasetConfigurationPtrOutputWithContext(context.Context) ReportDatasetConfigurationPtrOutput
}

type reportDatasetConfigurationPtrType ReportDatasetConfigurationArgs

func ReportDatasetConfigurationPtr(v *ReportDatasetConfigurationArgs) ReportDatasetConfigurationPtrInput {
	return (*reportDatasetConfigurationPtrType)(v)
}

func (*reportDatasetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDatasetConfiguration)(nil)).Elem()
}

func (i *reportDatasetConfigurationPtrType) ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput {
	return i.ToReportDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i *reportDatasetConfigurationPtrType) ToReportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetConfigurationPtrOutput)
}

type ReportDatasetConfigurationOutput struct{ *pulumi.OutputState }

func (ReportDatasetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDatasetConfiguration)(nil)).Elem()
}

func (o ReportDatasetConfigurationOutput) ToReportDatasetConfigurationOutput() ReportDatasetConfigurationOutput {
	return o
}

func (o ReportDatasetConfigurationOutput) ToReportDatasetConfigurationOutputWithContext(ctx context.Context) ReportDatasetConfigurationOutput {
	return o
}

func (o ReportDatasetConfigurationOutput) ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput {
	return o.ToReportDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (o ReportDatasetConfigurationOutput) ToReportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportDatasetConfiguration) *ReportDatasetConfiguration {
		return &v
	}).(ReportDatasetConfigurationPtrOutput)
}

func (o ReportDatasetConfigurationOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportDatasetConfiguration) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ReportDatasetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ReportDatasetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDatasetConfiguration)(nil)).Elem()
}

func (o ReportDatasetConfigurationPtrOutput) ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput {
	return o
}

func (o ReportDatasetConfigurationPtrOutput) ToReportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationPtrOutput {
	return o
}

func (o ReportDatasetConfigurationPtrOutput) Elem() ReportDatasetConfigurationOutput {
	return o.ApplyT(func(v *ReportDatasetConfiguration) ReportDatasetConfiguration {
		if v != nil {
			return *v
		}
		var ret ReportDatasetConfiguration
		return ret
	}).(ReportDatasetConfigurationOutput)
}

func (o ReportDatasetConfigurationPtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportDatasetConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ReportDatasetConfigurationResponse struct {
	Columns []string `pulumi:"columns"`
}

type ReportDatasetConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ReportDatasetConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDatasetConfigurationResponse)(nil)).Elem()
}

func (o ReportDatasetConfigurationResponseOutput) ToReportDatasetConfigurationResponseOutput() ReportDatasetConfigurationResponseOutput {
	return o
}

func (o ReportDatasetConfigurationResponseOutput) ToReportDatasetConfigurationResponseOutputWithContext(ctx context.Context) ReportDatasetConfigurationResponseOutput {
	return o
}

func (o ReportDatasetConfigurationResponseOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportDatasetConfigurationResponse) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ReportDatasetConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportDatasetConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDatasetConfigurationResponse)(nil)).Elem()
}

func (o ReportDatasetConfigurationResponsePtrOutput) ToReportDatasetConfigurationResponsePtrOutput() ReportDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ReportDatasetConfigurationResponsePtrOutput) ToReportDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ReportDatasetConfigurationResponsePtrOutput) Elem() ReportDatasetConfigurationResponseOutput {
	return o.ApplyT(func(v *ReportDatasetConfigurationResponse) ReportDatasetConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ReportDatasetConfigurationResponse
		return ret
	}).(ReportDatasetConfigurationResponseOutput)
}

func (o ReportDatasetConfigurationResponsePtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportDatasetConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ReportDatasetResponse struct {
	Aggregation   map[string]ReportAggregationResponse `pulumi:"aggregation"`
	Configuration *ReportDatasetConfigurationResponse  `pulumi:"configuration"`
	Filter        *ReportFilterResponse                `pulumi:"filter"`
	Granularity   *string                              `pulumi:"granularity"`
	Grouping      []ReportGroupingResponse             `pulumi:"grouping"`
}

type ReportDatasetResponseOutput struct{ *pulumi.OutputState }

func (ReportDatasetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDatasetResponse)(nil)).Elem()
}

func (o ReportDatasetResponseOutput) ToReportDatasetResponseOutput() ReportDatasetResponseOutput {
	return o
}

func (o ReportDatasetResponseOutput) ToReportDatasetResponseOutputWithContext(ctx context.Context) ReportDatasetResponseOutput {
	return o
}

func (o ReportDatasetResponseOutput) Aggregation() ReportAggregationResponseMapOutput {
	return o.ApplyT(func(v ReportDatasetResponse) map[string]ReportAggregationResponse { return v.Aggregation }).(ReportAggregationResponseMapOutput)
}

func (o ReportDatasetResponseOutput) Configuration() ReportDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ReportDatasetResponse) *ReportDatasetConfigurationResponse { return v.Configuration }).(ReportDatasetConfigurationResponsePtrOutput)
}

func (o ReportDatasetResponseOutput) Filter() ReportFilterResponsePtrOutput {
	return o.ApplyT(func(v ReportDatasetResponse) *ReportFilterResponse { return v.Filter }).(ReportFilterResponsePtrOutput)
}

func (o ReportDatasetResponseOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportDatasetResponse) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o ReportDatasetResponseOutput) Grouping() ReportGroupingResponseArrayOutput {
	return o.ApplyT(func(v ReportDatasetResponse) []ReportGroupingResponse { return v.Grouping }).(ReportGroupingResponseArrayOutput)
}

type ReportDatasetResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportDatasetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDatasetResponse)(nil)).Elem()
}

func (o ReportDatasetResponsePtrOutput) ToReportDatasetResponsePtrOutput() ReportDatasetResponsePtrOutput {
	return o
}

func (o ReportDatasetResponsePtrOutput) ToReportDatasetResponsePtrOutputWithContext(ctx context.Context) ReportDatasetResponsePtrOutput {
	return o
}

func (o ReportDatasetResponsePtrOutput) Elem() ReportDatasetResponseOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) ReportDatasetResponse {
		if v != nil {
			return *v
		}
		var ret ReportDatasetResponse
		return ret
	}).(ReportDatasetResponseOutput)
}

func (o ReportDatasetResponsePtrOutput) Aggregation() ReportAggregationResponseMapOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) map[string]ReportAggregationResponse {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(ReportAggregationResponseMapOutput)
}

func (o ReportDatasetResponsePtrOutput) Configuration() ReportDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) *ReportDatasetConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ReportDatasetConfigurationResponsePtrOutput)
}

func (o ReportDatasetResponsePtrOutput) Filter() ReportFilterResponsePtrOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) *ReportFilterResponse {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(ReportFilterResponsePtrOutput)
}

func (o ReportDatasetResponsePtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o ReportDatasetResponsePtrOutput) Grouping() ReportGroupingResponseArrayOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) []ReportGroupingResponse {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(ReportGroupingResponseArrayOutput)
}

type ReportDefinition struct {
	Dataset    *ReportDataset    `pulumi:"dataset"`
	TimePeriod *ReportTimePeriod `pulumi:"timePeriod"`
	Timeframe  string            `pulumi:"timeframe"`
	Type       string            `pulumi:"type"`
}





type ReportDefinitionInput interface {
	pulumi.Input

	ToReportDefinitionOutput() ReportDefinitionOutput
	ToReportDefinitionOutputWithContext(context.Context) ReportDefinitionOutput
}

type ReportDefinitionArgs struct {
	Dataset    ReportDatasetPtrInput    `pulumi:"dataset"`
	TimePeriod ReportTimePeriodPtrInput `pulumi:"timePeriod"`
	Timeframe  pulumi.StringInput       `pulumi:"timeframe"`
	Type       pulumi.StringInput       `pulumi:"type"`
}

func (ReportDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDefinition)(nil)).Elem()
}

func (i ReportDefinitionArgs) ToReportDefinitionOutput() ReportDefinitionOutput {
	return i.ToReportDefinitionOutputWithContext(context.Background())
}

func (i ReportDefinitionArgs) ToReportDefinitionOutputWithContext(ctx context.Context) ReportDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDefinitionOutput)
}

type ReportDefinitionOutput struct{ *pulumi.OutputState }

func (ReportDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDefinition)(nil)).Elem()
}

func (o ReportDefinitionOutput) ToReportDefinitionOutput() ReportDefinitionOutput {
	return o
}

func (o ReportDefinitionOutput) ToReportDefinitionOutputWithContext(ctx context.Context) ReportDefinitionOutput {
	return o
}

func (o ReportDefinitionOutput) Dataset() ReportDatasetPtrOutput {
	return o.ApplyT(func(v ReportDefinition) *ReportDataset { return v.Dataset }).(ReportDatasetPtrOutput)
}

func (o ReportDefinitionOutput) TimePeriod() ReportTimePeriodPtrOutput {
	return o.ApplyT(func(v ReportDefinition) *ReportTimePeriod { return v.TimePeriod }).(ReportTimePeriodPtrOutput)
}

func (o ReportDefinitionOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDefinition) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ReportDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type ReportDefinitionResponse struct {
	Dataset    *ReportDatasetResponse    `pulumi:"dataset"`
	TimePeriod *ReportTimePeriodResponse `pulumi:"timePeriod"`
	Timeframe  string                    `pulumi:"timeframe"`
	Type       string                    `pulumi:"type"`
}

type ReportDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ReportDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDefinitionResponse)(nil)).Elem()
}

func (o ReportDefinitionResponseOutput) ToReportDefinitionResponseOutput() ReportDefinitionResponseOutput {
	return o
}

func (o ReportDefinitionResponseOutput) ToReportDefinitionResponseOutputWithContext(ctx context.Context) ReportDefinitionResponseOutput {
	return o
}

func (o ReportDefinitionResponseOutput) Dataset() ReportDatasetResponsePtrOutput {
	return o.ApplyT(func(v ReportDefinitionResponse) *ReportDatasetResponse { return v.Dataset }).(ReportDatasetResponsePtrOutput)
}

func (o ReportDefinitionResponseOutput) TimePeriod() ReportTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v ReportDefinitionResponse) *ReportTimePeriodResponse { return v.TimePeriod }).(ReportTimePeriodResponsePtrOutput)
}

func (o ReportDefinitionResponseOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDefinitionResponse) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ReportDefinitionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDefinitionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ReportDeliveryDestination struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}





type ReportDeliveryDestinationInput interface {
	pulumi.Input

	ToReportDeliveryDestinationOutput() ReportDeliveryDestinationOutput
	ToReportDeliveryDestinationOutputWithContext(context.Context) ReportDeliveryDestinationOutput
}

type ReportDeliveryDestinationArgs struct {
	Container      pulumi.StringInput    `pulumi:"container"`
	ResourceId     pulumi.StringInput    `pulumi:"resourceId"`
	RootFolderPath pulumi.StringPtrInput `pulumi:"rootFolderPath"`
}

func (ReportDeliveryDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryDestination)(nil)).Elem()
}

func (i ReportDeliveryDestinationArgs) ToReportDeliveryDestinationOutput() ReportDeliveryDestinationOutput {
	return i.ToReportDeliveryDestinationOutputWithContext(context.Background())
}

func (i ReportDeliveryDestinationArgs) ToReportDeliveryDestinationOutputWithContext(ctx context.Context) ReportDeliveryDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDeliveryDestinationOutput)
}

type ReportDeliveryDestinationOutput struct{ *pulumi.OutputState }

func (ReportDeliveryDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryDestination)(nil)).Elem()
}

func (o ReportDeliveryDestinationOutput) ToReportDeliveryDestinationOutput() ReportDeliveryDestinationOutput {
	return o
}

func (o ReportDeliveryDestinationOutput) ToReportDeliveryDestinationOutputWithContext(ctx context.Context) ReportDeliveryDestinationOutput {
	return o
}

func (o ReportDeliveryDestinationOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDeliveryDestination) string { return v.Container }).(pulumi.StringOutput)
}

func (o ReportDeliveryDestinationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDeliveryDestination) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ReportDeliveryDestinationOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportDeliveryDestination) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ReportDeliveryDestinationResponse struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}

type ReportDeliveryDestinationResponseOutput struct{ *pulumi.OutputState }

func (ReportDeliveryDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryDestinationResponse)(nil)).Elem()
}

func (o ReportDeliveryDestinationResponseOutput) ToReportDeliveryDestinationResponseOutput() ReportDeliveryDestinationResponseOutput {
	return o
}

func (o ReportDeliveryDestinationResponseOutput) ToReportDeliveryDestinationResponseOutputWithContext(ctx context.Context) ReportDeliveryDestinationResponseOutput {
	return o
}

func (o ReportDeliveryDestinationResponseOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDeliveryDestinationResponse) string { return v.Container }).(pulumi.StringOutput)
}

func (o ReportDeliveryDestinationResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDeliveryDestinationResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ReportDeliveryDestinationResponseOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportDeliveryDestinationResponse) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ReportDeliveryInfo struct {
	Destination ReportDeliveryDestination `pulumi:"destination"`
}





type ReportDeliveryInfoInput interface {
	pulumi.Input

	ToReportDeliveryInfoOutput() ReportDeliveryInfoOutput
	ToReportDeliveryInfoOutputWithContext(context.Context) ReportDeliveryInfoOutput
}

type ReportDeliveryInfoArgs struct {
	Destination ReportDeliveryDestinationInput `pulumi:"destination"`
}

func (ReportDeliveryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryInfo)(nil)).Elem()
}

func (i ReportDeliveryInfoArgs) ToReportDeliveryInfoOutput() ReportDeliveryInfoOutput {
	return i.ToReportDeliveryInfoOutputWithContext(context.Background())
}

func (i ReportDeliveryInfoArgs) ToReportDeliveryInfoOutputWithContext(ctx context.Context) ReportDeliveryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDeliveryInfoOutput)
}

type ReportDeliveryInfoOutput struct{ *pulumi.OutputState }

func (ReportDeliveryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryInfo)(nil)).Elem()
}

func (o ReportDeliveryInfoOutput) ToReportDeliveryInfoOutput() ReportDeliveryInfoOutput {
	return o
}

func (o ReportDeliveryInfoOutput) ToReportDeliveryInfoOutputWithContext(ctx context.Context) ReportDeliveryInfoOutput {
	return o
}

func (o ReportDeliveryInfoOutput) Destination() ReportDeliveryDestinationOutput {
	return o.ApplyT(func(v ReportDeliveryInfo) ReportDeliveryDestination { return v.Destination }).(ReportDeliveryDestinationOutput)
}

type ReportDeliveryInfoResponse struct {
	Destination ReportDeliveryDestinationResponse `pulumi:"destination"`
}

type ReportDeliveryInfoResponseOutput struct{ *pulumi.OutputState }

func (ReportDeliveryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryInfoResponse)(nil)).Elem()
}

func (o ReportDeliveryInfoResponseOutput) ToReportDeliveryInfoResponseOutput() ReportDeliveryInfoResponseOutput {
	return o
}

func (o ReportDeliveryInfoResponseOutput) ToReportDeliveryInfoResponseOutputWithContext(ctx context.Context) ReportDeliveryInfoResponseOutput {
	return o
}

func (o ReportDeliveryInfoResponseOutput) Destination() ReportDeliveryDestinationResponseOutput {
	return o.ApplyT(func(v ReportDeliveryInfoResponse) ReportDeliveryDestinationResponse { return v.Destination }).(ReportDeliveryDestinationResponseOutput)
}

type ReportFilter struct {
	And       []ReportFilter              `pulumi:"and"`
	Dimension *ReportComparisonExpression `pulumi:"dimension"`
	Not       *ReportFilter               `pulumi:"not"`
	Or        []ReportFilter              `pulumi:"or"`
	Tag       *ReportComparisonExpression `pulumi:"tag"`
}





type ReportFilterInput interface {
	pulumi.Input

	ToReportFilterOutput() ReportFilterOutput
	ToReportFilterOutputWithContext(context.Context) ReportFilterOutput
}

type ReportFilterArgs struct {
	And       ReportFilterArrayInput             `pulumi:"and"`
	Dimension ReportComparisonExpressionPtrInput `pulumi:"dimension"`
	Not       ReportFilterPtrInput               `pulumi:"not"`
	Or        ReportFilterArrayInput             `pulumi:"or"`
	Tag       ReportComparisonExpressionPtrInput `pulumi:"tag"`
}

func (ReportFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportFilter)(nil)).Elem()
}

func (i ReportFilterArgs) ToReportFilterOutput() ReportFilterOutput {
	return i.ToReportFilterOutputWithContext(context.Background())
}

func (i ReportFilterArgs) ToReportFilterOutputWithContext(ctx context.Context) ReportFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportFilterOutput)
}

func (i ReportFilterArgs) ToReportFilterPtrOutput() ReportFilterPtrOutput {
	return i.ToReportFilterPtrOutputWithContext(context.Background())
}

func (i ReportFilterArgs) ToReportFilterPtrOutputWithContext(ctx context.Context) ReportFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportFilterOutput).ToReportFilterPtrOutputWithContext(ctx)
}









type ReportFilterPtrInput interface {
	pulumi.Input

	ToReportFilterPtrOutput() ReportFilterPtrOutput
	ToReportFilterPtrOutputWithContext(context.Context) ReportFilterPtrOutput
}

type reportFilterPtrType ReportFilterArgs

func ReportFilterPtr(v *ReportFilterArgs) ReportFilterPtrInput {
	return (*reportFilterPtrType)(v)
}

func (*reportFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportFilter)(nil)).Elem()
}

func (i *reportFilterPtrType) ToReportFilterPtrOutput() ReportFilterPtrOutput {
	return i.ToReportFilterPtrOutputWithContext(context.Background())
}

func (i *reportFilterPtrType) ToReportFilterPtrOutputWithContext(ctx context.Context) ReportFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportFilterPtrOutput)
}





type ReportFilterArrayInput interface {
	pulumi.Input

	ToReportFilterArrayOutput() ReportFilterArrayOutput
	ToReportFilterArrayOutputWithContext(context.Context) ReportFilterArrayOutput
}

type ReportFilterArray []ReportFilterInput

func (ReportFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportFilter)(nil)).Elem()
}

func (i ReportFilterArray) ToReportFilterArrayOutput() ReportFilterArrayOutput {
	return i.ToReportFilterArrayOutputWithContext(context.Background())
}

func (i ReportFilterArray) ToReportFilterArrayOutputWithContext(ctx context.Context) ReportFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportFilterArrayOutput)
}

type ReportFilterOutput struct{ *pulumi.OutputState }

func (ReportFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportFilter)(nil)).Elem()
}

func (o ReportFilterOutput) ToReportFilterOutput() ReportFilterOutput {
	return o
}

func (o ReportFilterOutput) ToReportFilterOutputWithContext(ctx context.Context) ReportFilterOutput {
	return o
}

func (o ReportFilterOutput) ToReportFilterPtrOutput() ReportFilterPtrOutput {
	return o.ToReportFilterPtrOutputWithContext(context.Background())
}

func (o ReportFilterOutput) ToReportFilterPtrOutputWithContext(ctx context.Context) ReportFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportFilter) *ReportFilter {
		return &v
	}).(ReportFilterPtrOutput)
}

func (o ReportFilterOutput) And() ReportFilterArrayOutput {
	return o.ApplyT(func(v ReportFilter) []ReportFilter { return v.And }).(ReportFilterArrayOutput)
}

func (o ReportFilterOutput) Dimension() ReportComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportFilter) *ReportComparisonExpression { return v.Dimension }).(ReportComparisonExpressionPtrOutput)
}

func (o ReportFilterOutput) Not() ReportFilterPtrOutput {
	return o.ApplyT(func(v ReportFilter) *ReportFilter { return v.Not }).(ReportFilterPtrOutput)
}

func (o ReportFilterOutput) Or() ReportFilterArrayOutput {
	return o.ApplyT(func(v ReportFilter) []ReportFilter { return v.Or }).(ReportFilterArrayOutput)
}

func (o ReportFilterOutput) Tag() ReportComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportFilter) *ReportComparisonExpression { return v.Tag }).(ReportComparisonExpressionPtrOutput)
}

type ReportFilterPtrOutput struct{ *pulumi.OutputState }

func (ReportFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportFilter)(nil)).Elem()
}

func (o ReportFilterPtrOutput) ToReportFilterPtrOutput() ReportFilterPtrOutput {
	return o
}

func (o ReportFilterPtrOutput) ToReportFilterPtrOutputWithContext(ctx context.Context) ReportFilterPtrOutput {
	return o
}

func (o ReportFilterPtrOutput) Elem() ReportFilterOutput {
	return o.ApplyT(func(v *ReportFilter) ReportFilter {
		if v != nil {
			return *v
		}
		var ret ReportFilter
		return ret
	}).(ReportFilterOutput)
}

func (o ReportFilterPtrOutput) And() ReportFilterArrayOutput {
	return o.ApplyT(func(v *ReportFilter) []ReportFilter {
		if v == nil {
			return nil
		}
		return v.And
	}).(ReportFilterArrayOutput)
}

func (o ReportFilterPtrOutput) Dimension() ReportComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportFilter) *ReportComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Dimension
	}).(ReportComparisonExpressionPtrOutput)
}

func (o ReportFilterPtrOutput) Not() ReportFilterPtrOutput {
	return o.ApplyT(func(v *ReportFilter) *ReportFilter {
		if v == nil {
			return nil
		}
		return v.Not
	}).(ReportFilterPtrOutput)
}

func (o ReportFilterPtrOutput) Or() ReportFilterArrayOutput {
	return o.ApplyT(func(v *ReportFilter) []ReportFilter {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportFilterArrayOutput)
}

func (o ReportFilterPtrOutput) Tag() ReportComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportFilter) *ReportComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(ReportComparisonExpressionPtrOutput)
}

type ReportFilterArrayOutput struct{ *pulumi.OutputState }

func (ReportFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportFilter)(nil)).Elem()
}

func (o ReportFilterArrayOutput) ToReportFilterArrayOutput() ReportFilterArrayOutput {
	return o
}

func (o ReportFilterArrayOutput) ToReportFilterArrayOutputWithContext(ctx context.Context) ReportFilterArrayOutput {
	return o
}

func (o ReportFilterArrayOutput) Index(i pulumi.IntInput) ReportFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportFilter {
		return vs[0].([]ReportFilter)[vs[1].(int)]
	}).(ReportFilterOutput)
}

type ReportFilterResponse struct {
	And       []ReportFilterResponse              `pulumi:"and"`
	Dimension *ReportComparisonExpressionResponse `pulumi:"dimension"`
	Not       *ReportFilterResponse               `pulumi:"not"`
	Or        []ReportFilterResponse              `pulumi:"or"`
	Tag       *ReportComparisonExpressionResponse `pulumi:"tag"`
}

type ReportFilterResponseOutput struct{ *pulumi.OutputState }

func (ReportFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportFilterResponse)(nil)).Elem()
}

func (o ReportFilterResponseOutput) ToReportFilterResponseOutput() ReportFilterResponseOutput {
	return o
}

func (o ReportFilterResponseOutput) ToReportFilterResponseOutputWithContext(ctx context.Context) ReportFilterResponseOutput {
	return o
}

func (o ReportFilterResponseOutput) And() ReportFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportFilterResponse) []ReportFilterResponse { return v.And }).(ReportFilterResponseArrayOutput)
}

func (o ReportFilterResponseOutput) Dimension() ReportComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportFilterResponse) *ReportComparisonExpressionResponse { return v.Dimension }).(ReportComparisonExpressionResponsePtrOutput)
}

func (o ReportFilterResponseOutput) Not() ReportFilterResponsePtrOutput {
	return o.ApplyT(func(v ReportFilterResponse) *ReportFilterResponse { return v.Not }).(ReportFilterResponsePtrOutput)
}

func (o ReportFilterResponseOutput) Or() ReportFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportFilterResponse) []ReportFilterResponse { return v.Or }).(ReportFilterResponseArrayOutput)
}

func (o ReportFilterResponseOutput) Tag() ReportComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportFilterResponse) *ReportComparisonExpressionResponse { return v.Tag }).(ReportComparisonExpressionResponsePtrOutput)
}

type ReportFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportFilterResponse)(nil)).Elem()
}

func (o ReportFilterResponsePtrOutput) ToReportFilterResponsePtrOutput() ReportFilterResponsePtrOutput {
	return o
}

func (o ReportFilterResponsePtrOutput) ToReportFilterResponsePtrOutputWithContext(ctx context.Context) ReportFilterResponsePtrOutput {
	return o
}

func (o ReportFilterResponsePtrOutput) Elem() ReportFilterResponseOutput {
	return o.ApplyT(func(v *ReportFilterResponse) ReportFilterResponse {
		if v != nil {
			return *v
		}
		var ret ReportFilterResponse
		return ret
	}).(ReportFilterResponseOutput)
}

func (o ReportFilterResponsePtrOutput) And() ReportFilterResponseArrayOutput {
	return o.ApplyT(func(v *ReportFilterResponse) []ReportFilterResponse {
		if v == nil {
			return nil
		}
		return v.And
	}).(ReportFilterResponseArrayOutput)
}

func (o ReportFilterResponsePtrOutput) Dimension() ReportComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportFilterResponse) *ReportComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Dimension
	}).(ReportComparisonExpressionResponsePtrOutput)
}

func (o ReportFilterResponsePtrOutput) Not() ReportFilterResponsePtrOutput {
	return o.ApplyT(func(v *ReportFilterResponse) *ReportFilterResponse {
		if v == nil {
			return nil
		}
		return v.Not
	}).(ReportFilterResponsePtrOutput)
}

func (o ReportFilterResponsePtrOutput) Or() ReportFilterResponseArrayOutput {
	return o.ApplyT(func(v *ReportFilterResponse) []ReportFilterResponse {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportFilterResponseArrayOutput)
}

func (o ReportFilterResponsePtrOutput) Tag() ReportComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportFilterResponse) *ReportComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(ReportComparisonExpressionResponsePtrOutput)
}

type ReportFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportFilterResponse)(nil)).Elem()
}

func (o ReportFilterResponseArrayOutput) ToReportFilterResponseArrayOutput() ReportFilterResponseArrayOutput {
	return o
}

func (o ReportFilterResponseArrayOutput) ToReportFilterResponseArrayOutputWithContext(ctx context.Context) ReportFilterResponseArrayOutput {
	return o
}

func (o ReportFilterResponseArrayOutput) Index(i pulumi.IntInput) ReportFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportFilterResponse {
		return vs[0].([]ReportFilterResponse)[vs[1].(int)]
	}).(ReportFilterResponseOutput)
}

type ReportGrouping struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type ReportGroupingInput interface {
	pulumi.Input

	ToReportGroupingOutput() ReportGroupingOutput
	ToReportGroupingOutputWithContext(context.Context) ReportGroupingOutput
}

type ReportGroupingArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ReportGroupingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportGrouping)(nil)).Elem()
}

func (i ReportGroupingArgs) ToReportGroupingOutput() ReportGroupingOutput {
	return i.ToReportGroupingOutputWithContext(context.Background())
}

func (i ReportGroupingArgs) ToReportGroupingOutputWithContext(ctx context.Context) ReportGroupingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportGroupingOutput)
}





type ReportGroupingArrayInput interface {
	pulumi.Input

	ToReportGroupingArrayOutput() ReportGroupingArrayOutput
	ToReportGroupingArrayOutputWithContext(context.Context) ReportGroupingArrayOutput
}

type ReportGroupingArray []ReportGroupingInput

func (ReportGroupingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportGrouping)(nil)).Elem()
}

func (i ReportGroupingArray) ToReportGroupingArrayOutput() ReportGroupingArrayOutput {
	return i.ToReportGroupingArrayOutputWithContext(context.Background())
}

func (i ReportGroupingArray) ToReportGroupingArrayOutputWithContext(ctx context.Context) ReportGroupingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportGroupingArrayOutput)
}

type ReportGroupingOutput struct{ *pulumi.OutputState }

func (ReportGroupingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportGrouping)(nil)).Elem()
}

func (o ReportGroupingOutput) ToReportGroupingOutput() ReportGroupingOutput {
	return o
}

func (o ReportGroupingOutput) ToReportGroupingOutputWithContext(ctx context.Context) ReportGroupingOutput {
	return o
}

func (o ReportGroupingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportGrouping) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportGroupingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportGrouping) string { return v.Type }).(pulumi.StringOutput)
}

type ReportGroupingArrayOutput struct{ *pulumi.OutputState }

func (ReportGroupingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportGrouping)(nil)).Elem()
}

func (o ReportGroupingArrayOutput) ToReportGroupingArrayOutput() ReportGroupingArrayOutput {
	return o
}

func (o ReportGroupingArrayOutput) ToReportGroupingArrayOutputWithContext(ctx context.Context) ReportGroupingArrayOutput {
	return o
}

func (o ReportGroupingArrayOutput) Index(i pulumi.IntInput) ReportGroupingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportGrouping {
		return vs[0].([]ReportGrouping)[vs[1].(int)]
	}).(ReportGroupingOutput)
}

type ReportGroupingResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

type ReportGroupingResponseOutput struct{ *pulumi.OutputState }

func (ReportGroupingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportGroupingResponse)(nil)).Elem()
}

func (o ReportGroupingResponseOutput) ToReportGroupingResponseOutput() ReportGroupingResponseOutput {
	return o
}

func (o ReportGroupingResponseOutput) ToReportGroupingResponseOutputWithContext(ctx context.Context) ReportGroupingResponseOutput {
	return o
}

func (o ReportGroupingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportGroupingResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportGroupingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportGroupingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ReportGroupingResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportGroupingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportGroupingResponse)(nil)).Elem()
}

func (o ReportGroupingResponseArrayOutput) ToReportGroupingResponseArrayOutput() ReportGroupingResponseArrayOutput {
	return o
}

func (o ReportGroupingResponseArrayOutput) ToReportGroupingResponseArrayOutputWithContext(ctx context.Context) ReportGroupingResponseArrayOutput {
	return o
}

func (o ReportGroupingResponseArrayOutput) Index(i pulumi.IntInput) ReportGroupingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportGroupingResponse {
		return vs[0].([]ReportGroupingResponse)[vs[1].(int)]
	}).(ReportGroupingResponseOutput)
}

type ReportRecurrencePeriod struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}





type ReportRecurrencePeriodInput interface {
	pulumi.Input

	ToReportRecurrencePeriodOutput() ReportRecurrencePeriodOutput
	ToReportRecurrencePeriodOutputWithContext(context.Context) ReportRecurrencePeriodOutput
}

type ReportRecurrencePeriodArgs struct {
	From pulumi.StringInput    `pulumi:"from"`
	To   pulumi.StringPtrInput `pulumi:"to"`
}

func (ReportRecurrencePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportRecurrencePeriod)(nil)).Elem()
}

func (i ReportRecurrencePeriodArgs) ToReportRecurrencePeriodOutput() ReportRecurrencePeriodOutput {
	return i.ToReportRecurrencePeriodOutputWithContext(context.Background())
}

func (i ReportRecurrencePeriodArgs) ToReportRecurrencePeriodOutputWithContext(ctx context.Context) ReportRecurrencePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportRecurrencePeriodOutput)
}

func (i ReportRecurrencePeriodArgs) ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput {
	return i.ToReportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i ReportRecurrencePeriodArgs) ToReportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportRecurrencePeriodOutput).ToReportRecurrencePeriodPtrOutputWithContext(ctx)
}









type ReportRecurrencePeriodPtrInput interface {
	pulumi.Input

	ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput
	ToReportRecurrencePeriodPtrOutputWithContext(context.Context) ReportRecurrencePeriodPtrOutput
}

type reportRecurrencePeriodPtrType ReportRecurrencePeriodArgs

func ReportRecurrencePeriodPtr(v *ReportRecurrencePeriodArgs) ReportRecurrencePeriodPtrInput {
	return (*reportRecurrencePeriodPtrType)(v)
}

func (*reportRecurrencePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportRecurrencePeriod)(nil)).Elem()
}

func (i *reportRecurrencePeriodPtrType) ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput {
	return i.ToReportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i *reportRecurrencePeriodPtrType) ToReportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportRecurrencePeriodPtrOutput)
}

type ReportRecurrencePeriodOutput struct{ *pulumi.OutputState }

func (ReportRecurrencePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportRecurrencePeriod)(nil)).Elem()
}

func (o ReportRecurrencePeriodOutput) ToReportRecurrencePeriodOutput() ReportRecurrencePeriodOutput {
	return o
}

func (o ReportRecurrencePeriodOutput) ToReportRecurrencePeriodOutputWithContext(ctx context.Context) ReportRecurrencePeriodOutput {
	return o
}

func (o ReportRecurrencePeriodOutput) ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput {
	return o.ToReportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (o ReportRecurrencePeriodOutput) ToReportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportRecurrencePeriod) *ReportRecurrencePeriod {
		return &v
	}).(ReportRecurrencePeriodPtrOutput)
}

func (o ReportRecurrencePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportRecurrencePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportRecurrencePeriodOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportRecurrencePeriod) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ReportRecurrencePeriodPtrOutput struct{ *pulumi.OutputState }

func (ReportRecurrencePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportRecurrencePeriod)(nil)).Elem()
}

func (o ReportRecurrencePeriodPtrOutput) ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput {
	return o
}

func (o ReportRecurrencePeriodPtrOutput) ToReportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodPtrOutput {
	return o
}

func (o ReportRecurrencePeriodPtrOutput) Elem() ReportRecurrencePeriodOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriod) ReportRecurrencePeriod {
		if v != nil {
			return *v
		}
		var ret ReportRecurrencePeriod
		return ret
	}).(ReportRecurrencePeriodOutput)
}

func (o ReportRecurrencePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportRecurrencePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ReportRecurrencePeriodResponse struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}

type ReportRecurrencePeriodResponseOutput struct{ *pulumi.OutputState }

func (ReportRecurrencePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportRecurrencePeriodResponse)(nil)).Elem()
}

func (o ReportRecurrencePeriodResponseOutput) ToReportRecurrencePeriodResponseOutput() ReportRecurrencePeriodResponseOutput {
	return o
}

func (o ReportRecurrencePeriodResponseOutput) ToReportRecurrencePeriodResponseOutputWithContext(ctx context.Context) ReportRecurrencePeriodResponseOutput {
	return o
}

func (o ReportRecurrencePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportRecurrencePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportRecurrencePeriodResponseOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportRecurrencePeriodResponse) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ReportRecurrencePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportRecurrencePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportRecurrencePeriodResponse)(nil)).Elem()
}

func (o ReportRecurrencePeriodResponsePtrOutput) ToReportRecurrencePeriodResponsePtrOutput() ReportRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ReportRecurrencePeriodResponsePtrOutput) ToReportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ReportRecurrencePeriodResponsePtrOutput) Elem() ReportRecurrencePeriodResponseOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriodResponse) ReportRecurrencePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ReportRecurrencePeriodResponse
		return ret
	}).(ReportRecurrencePeriodResponseOutput)
}

func (o ReportRecurrencePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportRecurrencePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ReportSchedule struct {
	Recurrence       string                  `pulumi:"recurrence"`
	RecurrencePeriod *ReportRecurrencePeriod `pulumi:"recurrencePeriod"`
	Status           *string                 `pulumi:"status"`
}





type ReportScheduleInput interface {
	pulumi.Input

	ToReportScheduleOutput() ReportScheduleOutput
	ToReportScheduleOutputWithContext(context.Context) ReportScheduleOutput
}

type ReportScheduleArgs struct {
	Recurrence       pulumi.StringInput             `pulumi:"recurrence"`
	RecurrencePeriod ReportRecurrencePeriodPtrInput `pulumi:"recurrencePeriod"`
	Status           pulumi.StringPtrInput          `pulumi:"status"`
}

func (ReportScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportSchedule)(nil)).Elem()
}

func (i ReportScheduleArgs) ToReportScheduleOutput() ReportScheduleOutput {
	return i.ToReportScheduleOutputWithContext(context.Background())
}

func (i ReportScheduleArgs) ToReportScheduleOutputWithContext(ctx context.Context) ReportScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportScheduleOutput)
}

func (i ReportScheduleArgs) ToReportSchedulePtrOutput() ReportSchedulePtrOutput {
	return i.ToReportSchedulePtrOutputWithContext(context.Background())
}

func (i ReportScheduleArgs) ToReportSchedulePtrOutputWithContext(ctx context.Context) ReportSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportScheduleOutput).ToReportSchedulePtrOutputWithContext(ctx)
}









type ReportSchedulePtrInput interface {
	pulumi.Input

	ToReportSchedulePtrOutput() ReportSchedulePtrOutput
	ToReportSchedulePtrOutputWithContext(context.Context) ReportSchedulePtrOutput
}

type reportSchedulePtrType ReportScheduleArgs

func ReportSchedulePtr(v *ReportScheduleArgs) ReportSchedulePtrInput {
	return (*reportSchedulePtrType)(v)
}

func (*reportSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportSchedule)(nil)).Elem()
}

func (i *reportSchedulePtrType) ToReportSchedulePtrOutput() ReportSchedulePtrOutput {
	return i.ToReportSchedulePtrOutputWithContext(context.Background())
}

func (i *reportSchedulePtrType) ToReportSchedulePtrOutputWithContext(ctx context.Context) ReportSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportSchedulePtrOutput)
}

type ReportScheduleOutput struct{ *pulumi.OutputState }

func (ReportScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportSchedule)(nil)).Elem()
}

func (o ReportScheduleOutput) ToReportScheduleOutput() ReportScheduleOutput {
	return o
}

func (o ReportScheduleOutput) ToReportScheduleOutputWithContext(ctx context.Context) ReportScheduleOutput {
	return o
}

func (o ReportScheduleOutput) ToReportSchedulePtrOutput() ReportSchedulePtrOutput {
	return o.ToReportSchedulePtrOutputWithContext(context.Background())
}

func (o ReportScheduleOutput) ToReportSchedulePtrOutputWithContext(ctx context.Context) ReportSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportSchedule) *ReportSchedule {
		return &v
	}).(ReportSchedulePtrOutput)
}

func (o ReportScheduleOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ReportSchedule) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ReportScheduleOutput) RecurrencePeriod() ReportRecurrencePeriodPtrOutput {
	return o.ApplyT(func(v ReportSchedule) *ReportRecurrencePeriod { return v.RecurrencePeriod }).(ReportRecurrencePeriodPtrOutput)
}

func (o ReportScheduleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportSchedule) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ReportSchedulePtrOutput struct{ *pulumi.OutputState }

func (ReportSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportSchedule)(nil)).Elem()
}

func (o ReportSchedulePtrOutput) ToReportSchedulePtrOutput() ReportSchedulePtrOutput {
	return o
}

func (o ReportSchedulePtrOutput) ToReportSchedulePtrOutputWithContext(ctx context.Context) ReportSchedulePtrOutput {
	return o
}

func (o ReportSchedulePtrOutput) Elem() ReportScheduleOutput {
	return o.ApplyT(func(v *ReportSchedule) ReportSchedule {
		if v != nil {
			return *v
		}
		var ret ReportSchedule
		return ret
	}).(ReportScheduleOutput)
}

func (o ReportSchedulePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportSchedule) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ReportSchedulePtrOutput) RecurrencePeriod() ReportRecurrencePeriodPtrOutput {
	return o.ApplyT(func(v *ReportSchedule) *ReportRecurrencePeriod {
		if v == nil {
			return nil
		}
		return v.RecurrencePeriod
	}).(ReportRecurrencePeriodPtrOutput)
}

func (o ReportSchedulePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportSchedule) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ReportScheduleResponse struct {
	Recurrence       string                          `pulumi:"recurrence"`
	RecurrencePeriod *ReportRecurrencePeriodResponse `pulumi:"recurrencePeriod"`
	Status           *string                         `pulumi:"status"`
}

type ReportScheduleResponseOutput struct{ *pulumi.OutputState }

func (ReportScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportScheduleResponse)(nil)).Elem()
}

func (o ReportScheduleResponseOutput) ToReportScheduleResponseOutput() ReportScheduleResponseOutput {
	return o
}

func (o ReportScheduleResponseOutput) ToReportScheduleResponseOutputWithContext(ctx context.Context) ReportScheduleResponseOutput {
	return o
}

func (o ReportScheduleResponseOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ReportScheduleResponse) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ReportScheduleResponseOutput) RecurrencePeriod() ReportRecurrencePeriodResponsePtrOutput {
	return o.ApplyT(func(v ReportScheduleResponse) *ReportRecurrencePeriodResponse { return v.RecurrencePeriod }).(ReportRecurrencePeriodResponsePtrOutput)
}

func (o ReportScheduleResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportScheduleResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ReportScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportScheduleResponse)(nil)).Elem()
}

func (o ReportScheduleResponsePtrOutput) ToReportScheduleResponsePtrOutput() ReportScheduleResponsePtrOutput {
	return o
}

func (o ReportScheduleResponsePtrOutput) ToReportScheduleResponsePtrOutputWithContext(ctx context.Context) ReportScheduleResponsePtrOutput {
	return o
}

func (o ReportScheduleResponsePtrOutput) Elem() ReportScheduleResponseOutput {
	return o.ApplyT(func(v *ReportScheduleResponse) ReportScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ReportScheduleResponse
		return ret
	}).(ReportScheduleResponseOutput)
}

func (o ReportScheduleResponsePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ReportScheduleResponsePtrOutput) RecurrencePeriod() ReportRecurrencePeriodResponsePtrOutput {
	return o.ApplyT(func(v *ReportScheduleResponse) *ReportRecurrencePeriodResponse {
		if v == nil {
			return nil
		}
		return v.RecurrencePeriod
	}).(ReportRecurrencePeriodResponsePtrOutput)
}

func (o ReportScheduleResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ReportTimePeriod struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type ReportTimePeriodInput interface {
	pulumi.Input

	ToReportTimePeriodOutput() ReportTimePeriodOutput
	ToReportTimePeriodOutputWithContext(context.Context) ReportTimePeriodOutput
}

type ReportTimePeriodArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (ReportTimePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportTimePeriod)(nil)).Elem()
}

func (i ReportTimePeriodArgs) ToReportTimePeriodOutput() ReportTimePeriodOutput {
	return i.ToReportTimePeriodOutputWithContext(context.Background())
}

func (i ReportTimePeriodArgs) ToReportTimePeriodOutputWithContext(ctx context.Context) ReportTimePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportTimePeriodOutput)
}

func (i ReportTimePeriodArgs) ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput {
	return i.ToReportTimePeriodPtrOutputWithContext(context.Background())
}

func (i ReportTimePeriodArgs) ToReportTimePeriodPtrOutputWithContext(ctx context.Context) ReportTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportTimePeriodOutput).ToReportTimePeriodPtrOutputWithContext(ctx)
}









type ReportTimePeriodPtrInput interface {
	pulumi.Input

	ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput
	ToReportTimePeriodPtrOutputWithContext(context.Context) ReportTimePeriodPtrOutput
}

type reportTimePeriodPtrType ReportTimePeriodArgs

func ReportTimePeriodPtr(v *ReportTimePeriodArgs) ReportTimePeriodPtrInput {
	return (*reportTimePeriodPtrType)(v)
}

func (*reportTimePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportTimePeriod)(nil)).Elem()
}

func (i *reportTimePeriodPtrType) ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput {
	return i.ToReportTimePeriodPtrOutputWithContext(context.Background())
}

func (i *reportTimePeriodPtrType) ToReportTimePeriodPtrOutputWithContext(ctx context.Context) ReportTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportTimePeriodPtrOutput)
}

type ReportTimePeriodOutput struct{ *pulumi.OutputState }

func (ReportTimePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportTimePeriod)(nil)).Elem()
}

func (o ReportTimePeriodOutput) ToReportTimePeriodOutput() ReportTimePeriodOutput {
	return o
}

func (o ReportTimePeriodOutput) ToReportTimePeriodOutputWithContext(ctx context.Context) ReportTimePeriodOutput {
	return o
}

func (o ReportTimePeriodOutput) ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput {
	return o.ToReportTimePeriodPtrOutputWithContext(context.Background())
}

func (o ReportTimePeriodOutput) ToReportTimePeriodPtrOutputWithContext(ctx context.Context) ReportTimePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportTimePeriod) *ReportTimePeriod {
		return &v
	}).(ReportTimePeriodPtrOutput)
}

func (o ReportTimePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportTimePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportTimePeriodOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ReportTimePeriod) string { return v.To }).(pulumi.StringOutput)
}

type ReportTimePeriodPtrOutput struct{ *pulumi.OutputState }

func (ReportTimePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportTimePeriod)(nil)).Elem()
}

func (o ReportTimePeriodPtrOutput) ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput {
	return o
}

func (o ReportTimePeriodPtrOutput) ToReportTimePeriodPtrOutputWithContext(ctx context.Context) ReportTimePeriodPtrOutput {
	return o
}

func (o ReportTimePeriodPtrOutput) Elem() ReportTimePeriodOutput {
	return o.ApplyT(func(v *ReportTimePeriod) ReportTimePeriod {
		if v != nil {
			return *v
		}
		var ret ReportTimePeriod
		return ret
	}).(ReportTimePeriodOutput)
}

func (o ReportTimePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportTimePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type ReportTimePeriodResponse struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}

type ReportTimePeriodResponseOutput struct{ *pulumi.OutputState }

func (ReportTimePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportTimePeriodResponse)(nil)).Elem()
}

func (o ReportTimePeriodResponseOutput) ToReportTimePeriodResponseOutput() ReportTimePeriodResponseOutput {
	return o
}

func (o ReportTimePeriodResponseOutput) ToReportTimePeriodResponseOutputWithContext(ctx context.Context) ReportTimePeriodResponseOutput {
	return o
}

func (o ReportTimePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportTimePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportTimePeriodResponseOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ReportTimePeriodResponse) string { return v.To }).(pulumi.StringOutput)
}

type ReportTimePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportTimePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportTimePeriodResponse)(nil)).Elem()
}

func (o ReportTimePeriodResponsePtrOutput) ToReportTimePeriodResponsePtrOutput() ReportTimePeriodResponsePtrOutput {
	return o
}

func (o ReportTimePeriodResponsePtrOutput) ToReportTimePeriodResponsePtrOutputWithContext(ctx context.Context) ReportTimePeriodResponsePtrOutput {
	return o
}

func (o ReportTimePeriodResponsePtrOutput) Elem() ReportTimePeriodResponseOutput {
	return o.ApplyT(func(v *ReportTimePeriodResponse) ReportTimePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ReportTimePeriodResponse
		return ret
	}).(ReportTimePeriodResponseOutput)
}

func (o ReportTimePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportTimePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectorCollectionErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(ConnectorCollectionErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectorCollectionInfoResponseOutput{})
	pulumi.RegisterOutputType(ReportAggregationOutput{})
	pulumi.RegisterOutputType(ReportAggregationMapOutput{})
	pulumi.RegisterOutputType(ReportAggregationResponseOutput{})
	pulumi.RegisterOutputType(ReportAggregationResponseMapOutput{})
	pulumi.RegisterOutputType(ReportComparisonExpressionOutput{})
	pulumi.RegisterOutputType(ReportComparisonExpressionPtrOutput{})
	pulumi.RegisterOutputType(ReportComparisonExpressionResponseOutput{})
	pulumi.RegisterOutputType(ReportComparisonExpressionResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportDatasetOutput{})
	pulumi.RegisterOutputType(ReportDatasetPtrOutput{})
	pulumi.RegisterOutputType(ReportDatasetConfigurationOutput{})
	pulumi.RegisterOutputType(ReportDatasetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ReportDatasetConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ReportDatasetConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportDatasetResponseOutput{})
	pulumi.RegisterOutputType(ReportDatasetResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportDefinitionOutput{})
	pulumi.RegisterOutputType(ReportDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ReportDeliveryDestinationOutput{})
	pulumi.RegisterOutputType(ReportDeliveryDestinationResponseOutput{})
	pulumi.RegisterOutputType(ReportDeliveryInfoOutput{})
	pulumi.RegisterOutputType(ReportDeliveryInfoResponseOutput{})
	pulumi.RegisterOutputType(ReportFilterOutput{})
	pulumi.RegisterOutputType(ReportFilterPtrOutput{})
	pulumi.RegisterOutputType(ReportFilterArrayOutput{})
	pulumi.RegisterOutputType(ReportFilterResponseOutput{})
	pulumi.RegisterOutputType(ReportFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportGroupingOutput{})
	pulumi.RegisterOutputType(ReportGroupingArrayOutput{})
	pulumi.RegisterOutputType(ReportGroupingResponseOutput{})
	pulumi.RegisterOutputType(ReportGroupingResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportRecurrencePeriodOutput{})
	pulumi.RegisterOutputType(ReportRecurrencePeriodPtrOutput{})
	pulumi.RegisterOutputType(ReportRecurrencePeriodResponseOutput{})
	pulumi.RegisterOutputType(ReportRecurrencePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportScheduleOutput{})
	pulumi.RegisterOutputType(ReportSchedulePtrOutput{})
	pulumi.RegisterOutputType(ReportScheduleResponseOutput{})
	pulumi.RegisterOutputType(ReportScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportTimePeriodOutput{})
	pulumi.RegisterOutputType(ReportTimePeriodPtrOutput{})
	pulumi.RegisterOutputType(ReportTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(ReportTimePeriodResponsePtrOutput{})
}
