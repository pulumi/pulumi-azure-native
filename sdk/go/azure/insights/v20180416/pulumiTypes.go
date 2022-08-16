


package v20180416

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertingAction struct {
	AznsAction      *AzNsActionGroup `pulumi:"aznsAction"`
	OdataType       string           `pulumi:"odataType"`
	Severity        string           `pulumi:"severity"`
	ThrottlingInMin *int             `pulumi:"throttlingInMin"`
	Trigger         TriggerCondition `pulumi:"trigger"`
}

type AlertingActionResponse struct {
	AznsAction      *AzNsActionGroupResponse `pulumi:"aznsAction"`
	OdataType       string                   `pulumi:"odataType"`
	Severity        string                   `pulumi:"severity"`
	ThrottlingInMin *int                     `pulumi:"throttlingInMin"`
	Trigger         TriggerConditionResponse `pulumi:"trigger"`
}

type AzNsActionGroup struct {
	ActionGroup          []string `pulumi:"actionGroup"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	EmailSubject         *string  `pulumi:"emailSubject"`
}

type AzNsActionGroupResponse struct {
	ActionGroup          []string `pulumi:"actionGroup"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	EmailSubject         *string  `pulumi:"emailSubject"`
}

type Criteria struct {
	Dimensions []Dimension `pulumi:"dimensions"`
	MetricName string      `pulumi:"metricName"`
}

type CriteriaResponse struct {
	Dimensions []DimensionResponse `pulumi:"dimensions"`
	MetricName string              `pulumi:"metricName"`
}

type Dimension struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type DimensionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type LogMetricTrigger struct {
	MetricColumn      *string  `pulumi:"metricColumn"`
	MetricTriggerType *string  `pulumi:"metricTriggerType"`
	Threshold         *float64 `pulumi:"threshold"`
	ThresholdOperator *string  `pulumi:"thresholdOperator"`
}

type LogMetricTriggerResponse struct {
	MetricColumn      *string  `pulumi:"metricColumn"`
	MetricTriggerType *string  `pulumi:"metricTriggerType"`
	Threshold         *float64 `pulumi:"threshold"`
	ThresholdOperator *string  `pulumi:"thresholdOperator"`
}

type LogToMetricAction struct {
	Criteria  []Criteria `pulumi:"criteria"`
	OdataType string     `pulumi:"odataType"`
}

type LogToMetricActionResponse struct {
	Criteria  []CriteriaResponse `pulumi:"criteria"`
	OdataType string             `pulumi:"odataType"`
}

type Schedule struct {
	FrequencyInMinutes  int `pulumi:"frequencyInMinutes"`
	TimeWindowInMinutes int `pulumi:"timeWindowInMinutes"`
}





type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(context.Context) ScheduleOutput
}

type ScheduleArgs struct {
	FrequencyInMinutes  pulumi.IntInput `pulumi:"frequencyInMinutes"`
	TimeWindowInMinutes pulumi.IntInput `pulumi:"timeWindowInMinutes"`
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (i ScheduleArgs) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

func (i ScheduleArgs) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput).ToSchedulePtrOutputWithContext(ctx)
}









type SchedulePtrInput interface {
	pulumi.Input

	ToSchedulePtrOutput() SchedulePtrOutput
	ToSchedulePtrOutputWithContext(context.Context) SchedulePtrOutput
}

type schedulePtrType ScheduleArgs

func SchedulePtr(v *ScheduleArgs) SchedulePtrInput {
	return (*schedulePtrType)(v)
}

func (*schedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (i *schedulePtrType) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i *schedulePtrType) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulePtrOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o.ToSchedulePtrOutputWithContext(context.Background())
}

func (o ScheduleOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Schedule) *Schedule {
		return &v
	}).(SchedulePtrOutput)
}

func (o ScheduleOutput) FrequencyInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v Schedule) int { return v.FrequencyInMinutes }).(pulumi.IntOutput)
}

func (o ScheduleOutput) TimeWindowInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v Schedule) int { return v.TimeWindowInMinutes }).(pulumi.IntOutput)
}

type SchedulePtrOutput struct{ *pulumi.OutputState }

func (SchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (o SchedulePtrOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) Elem() ScheduleOutput {
	return o.ApplyT(func(v *Schedule) Schedule {
		if v != nil {
			return *v
		}
		var ret Schedule
		return ret
	}).(ScheduleOutput)
}

func (o SchedulePtrOutput) FrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Schedule) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o SchedulePtrOutput) TimeWindowInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Schedule) *int {
		if v == nil {
			return nil
		}
		return &v.TimeWindowInMinutes
	}).(pulumi.IntPtrOutput)
}

type ScheduleResponse struct {
	FrequencyInMinutes  int `pulumi:"frequencyInMinutes"`
	TimeWindowInMinutes int `pulumi:"timeWindowInMinutes"`
}

type ScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseOutput) ToScheduleResponseOutput() ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) FrequencyInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleResponse) int { return v.FrequencyInMinutes }).(pulumi.IntOutput)
}

func (o ScheduleResponseOutput) TimeWindowInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleResponse) int { return v.TimeWindowInMinutes }).(pulumi.IntOutput)
}

type ScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) Elem() ScheduleResponseOutput {
	return o.ApplyT(func(v *ScheduleResponse) ScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ScheduleResponse
		return ret
	}).(ScheduleResponseOutput)
}

func (o ScheduleResponsePtrOutput) FrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o ScheduleResponsePtrOutput) TimeWindowInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TimeWindowInMinutes
	}).(pulumi.IntPtrOutput)
}

type Source struct {
	AuthorizedResources []string `pulumi:"authorizedResources"`
	DataSourceId        string   `pulumi:"dataSourceId"`
	Query               *string  `pulumi:"query"`
	QueryType           *string  `pulumi:"queryType"`
}





type SourceInput interface {
	pulumi.Input

	ToSourceOutput() SourceOutput
	ToSourceOutputWithContext(context.Context) SourceOutput
}

type SourceArgs struct {
	AuthorizedResources pulumi.StringArrayInput `pulumi:"authorizedResources"`
	DataSourceId        pulumi.StringInput      `pulumi:"dataSourceId"`
	Query               pulumi.StringPtrInput   `pulumi:"query"`
	QueryType           pulumi.StringPtrInput   `pulumi:"queryType"`
}

func (SourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (i SourceArgs) ToSourceOutput() SourceOutput {
	return i.ToSourceOutputWithContext(context.Background())
}

func (i SourceArgs) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutput)
}

type SourceOutput struct{ *pulumi.OutputState }

func (SourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (o SourceOutput) ToSourceOutput() SourceOutput {
	return o
}

func (o SourceOutput) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return o
}

func (o SourceOutput) AuthorizedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Source) []string { return v.AuthorizedResources }).(pulumi.StringArrayOutput)
}

func (o SourceOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v Source) string { return v.DataSourceId }).(pulumi.StringOutput)
}

func (o SourceOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Source) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o SourceOutput) QueryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Source) *string { return v.QueryType }).(pulumi.StringPtrOutput)
}

type SourceResponse struct {
	AuthorizedResources []string `pulumi:"authorizedResources"`
	DataSourceId        string   `pulumi:"dataSourceId"`
	Query               *string  `pulumi:"query"`
	QueryType           *string  `pulumi:"queryType"`
}

type SourceResponseOutput struct{ *pulumi.OutputState }

func (SourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceResponse)(nil)).Elem()
}

func (o SourceResponseOutput) ToSourceResponseOutput() SourceResponseOutput {
	return o
}

func (o SourceResponseOutput) ToSourceResponseOutputWithContext(ctx context.Context) SourceResponseOutput {
	return o
}

func (o SourceResponseOutput) AuthorizedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SourceResponse) []string { return v.AuthorizedResources }).(pulumi.StringArrayOutput)
}

func (o SourceResponseOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SourceResponse) string { return v.DataSourceId }).(pulumi.StringOutput)
}

func (o SourceResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o SourceResponseOutput) QueryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceResponse) *string { return v.QueryType }).(pulumi.StringPtrOutput)
}

type TriggerCondition struct {
	MetricTrigger     *LogMetricTrigger `pulumi:"metricTrigger"`
	Threshold         float64           `pulumi:"threshold"`
	ThresholdOperator string            `pulumi:"thresholdOperator"`
}

type TriggerConditionResponse struct {
	MetricTrigger     *LogMetricTriggerResponse `pulumi:"metricTrigger"`
	Threshold         float64                   `pulumi:"threshold"`
	ThresholdOperator string                    `pulumi:"thresholdOperator"`
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(SchedulePtrOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceOutput{})
	pulumi.RegisterOutputType(SourceResponseOutput{})
}
