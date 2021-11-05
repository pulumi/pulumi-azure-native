


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Actions struct {
	ActionGroups     []string          `pulumi:"actionGroups"`
	CustomProperties map[string]string `pulumi:"customProperties"`
}





type ActionsInput interface {
	pulumi.Input

	ToActionsOutput() ActionsOutput
	ToActionsOutputWithContext(context.Context) ActionsOutput
}

type ActionsArgs struct {
	ActionGroups     pulumi.StringArrayInput `pulumi:"actionGroups"`
	CustomProperties pulumi.StringMapInput   `pulumi:"customProperties"`
}

func (ActionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Actions)(nil)).Elem()
}

func (i ActionsArgs) ToActionsOutput() ActionsOutput {
	return i.ToActionsOutputWithContext(context.Background())
}

func (i ActionsArgs) ToActionsOutputWithContext(ctx context.Context) ActionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOutput)
}

func (i ActionsArgs) ToActionsPtrOutput() ActionsPtrOutput {
	return i.ToActionsPtrOutputWithContext(context.Background())
}

func (i ActionsArgs) ToActionsPtrOutputWithContext(ctx context.Context) ActionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsOutput).ToActionsPtrOutputWithContext(ctx)
}









type ActionsPtrInput interface {
	pulumi.Input

	ToActionsPtrOutput() ActionsPtrOutput
	ToActionsPtrOutputWithContext(context.Context) ActionsPtrOutput
}

type actionsPtrType ActionsArgs

func ActionsPtr(v *ActionsArgs) ActionsPtrInput {
	return (*actionsPtrType)(v)
}

func (*actionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Actions)(nil)).Elem()
}

func (i *actionsPtrType) ToActionsPtrOutput() ActionsPtrOutput {
	return i.ToActionsPtrOutputWithContext(context.Background())
}

func (i *actionsPtrType) ToActionsPtrOutputWithContext(ctx context.Context) ActionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsPtrOutput)
}

type ActionsOutput struct{ *pulumi.OutputState }

func (ActionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Actions)(nil)).Elem()
}

func (o ActionsOutput) ToActionsOutput() ActionsOutput {
	return o
}

func (o ActionsOutput) ToActionsOutputWithContext(ctx context.Context) ActionsOutput {
	return o
}

func (o ActionsOutput) ToActionsPtrOutput() ActionsPtrOutput {
	return o.ToActionsPtrOutputWithContext(context.Background())
}

func (o ActionsOutput) ToActionsPtrOutputWithContext(ctx context.Context) ActionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Actions) *Actions {
		return &v
	}).(ActionsPtrOutput)
}

func (o ActionsOutput) ActionGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Actions) []string { return v.ActionGroups }).(pulumi.StringArrayOutput)
}

func (o ActionsOutput) CustomProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v Actions) map[string]string { return v.CustomProperties }).(pulumi.StringMapOutput)
}

type ActionsPtrOutput struct{ *pulumi.OutputState }

func (ActionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Actions)(nil)).Elem()
}

func (o ActionsPtrOutput) ToActionsPtrOutput() ActionsPtrOutput {
	return o
}

func (o ActionsPtrOutput) ToActionsPtrOutputWithContext(ctx context.Context) ActionsPtrOutput {
	return o
}

func (o ActionsPtrOutput) Elem() ActionsOutput {
	return o.ApplyT(func(v *Actions) Actions {
		if v != nil {
			return *v
		}
		var ret Actions
		return ret
	}).(ActionsOutput)
}

func (o ActionsPtrOutput) ActionGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Actions) []string {
		if v == nil {
			return nil
		}
		return v.ActionGroups
	}).(pulumi.StringArrayOutput)
}

func (o ActionsPtrOutput) CustomProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Actions) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomProperties
	}).(pulumi.StringMapOutput)
}

type ActionsResponse struct {
	ActionGroups     []string          `pulumi:"actionGroups"`
	CustomProperties map[string]string `pulumi:"customProperties"`
}





type ActionsResponseInput interface {
	pulumi.Input

	ToActionsResponseOutput() ActionsResponseOutput
	ToActionsResponseOutputWithContext(context.Context) ActionsResponseOutput
}

type ActionsResponseArgs struct {
	ActionGroups     pulumi.StringArrayInput `pulumi:"actionGroups"`
	CustomProperties pulumi.StringMapInput   `pulumi:"customProperties"`
}

func (ActionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionsResponse)(nil)).Elem()
}

func (i ActionsResponseArgs) ToActionsResponseOutput() ActionsResponseOutput {
	return i.ToActionsResponseOutputWithContext(context.Background())
}

func (i ActionsResponseArgs) ToActionsResponseOutputWithContext(ctx context.Context) ActionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsResponseOutput)
}

func (i ActionsResponseArgs) ToActionsResponsePtrOutput() ActionsResponsePtrOutput {
	return i.ToActionsResponsePtrOutputWithContext(context.Background())
}

func (i ActionsResponseArgs) ToActionsResponsePtrOutputWithContext(ctx context.Context) ActionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsResponseOutput).ToActionsResponsePtrOutputWithContext(ctx)
}









type ActionsResponsePtrInput interface {
	pulumi.Input

	ToActionsResponsePtrOutput() ActionsResponsePtrOutput
	ToActionsResponsePtrOutputWithContext(context.Context) ActionsResponsePtrOutput
}

type actionsResponsePtrType ActionsResponseArgs

func ActionsResponsePtr(v *ActionsResponseArgs) ActionsResponsePtrInput {
	return (*actionsResponsePtrType)(v)
}

func (*actionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsResponse)(nil)).Elem()
}

func (i *actionsResponsePtrType) ToActionsResponsePtrOutput() ActionsResponsePtrOutput {
	return i.ToActionsResponsePtrOutputWithContext(context.Background())
}

func (i *actionsResponsePtrType) ToActionsResponsePtrOutputWithContext(ctx context.Context) ActionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionsResponsePtrOutput)
}

type ActionsResponseOutput struct{ *pulumi.OutputState }

func (ActionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionsResponse)(nil)).Elem()
}

func (o ActionsResponseOutput) ToActionsResponseOutput() ActionsResponseOutput {
	return o
}

func (o ActionsResponseOutput) ToActionsResponseOutputWithContext(ctx context.Context) ActionsResponseOutput {
	return o
}

func (o ActionsResponseOutput) ToActionsResponsePtrOutput() ActionsResponsePtrOutput {
	return o.ToActionsResponsePtrOutputWithContext(context.Background())
}

func (o ActionsResponseOutput) ToActionsResponsePtrOutputWithContext(ctx context.Context) ActionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionsResponse) *ActionsResponse {
		return &v
	}).(ActionsResponsePtrOutput)
}

func (o ActionsResponseOutput) ActionGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActionsResponse) []string { return v.ActionGroups }).(pulumi.StringArrayOutput)
}

func (o ActionsResponseOutput) CustomProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ActionsResponse) map[string]string { return v.CustomProperties }).(pulumi.StringMapOutput)
}

type ActionsResponsePtrOutput struct{ *pulumi.OutputState }

func (ActionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionsResponse)(nil)).Elem()
}

func (o ActionsResponsePtrOutput) ToActionsResponsePtrOutput() ActionsResponsePtrOutput {
	return o
}

func (o ActionsResponsePtrOutput) ToActionsResponsePtrOutputWithContext(ctx context.Context) ActionsResponsePtrOutput {
	return o
}

func (o ActionsResponsePtrOutput) Elem() ActionsResponseOutput {
	return o.ApplyT(func(v *ActionsResponse) ActionsResponse {
		if v != nil {
			return *v
		}
		var ret ActionsResponse
		return ret
	}).(ActionsResponseOutput)
}

func (o ActionsResponsePtrOutput) ActionGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActionsResponse) []string {
		if v == nil {
			return nil
		}
		return v.ActionGroups
	}).(pulumi.StringArrayOutput)
}

func (o ActionsResponsePtrOutput) CustomProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ActionsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.CustomProperties
	}).(pulumi.StringMapOutput)
}

type Condition struct {
	Dimensions          []Dimension              `pulumi:"dimensions"`
	FailingPeriods      *ConditionFailingPeriods `pulumi:"failingPeriods"`
	MetricMeasureColumn *string                  `pulumi:"metricMeasureColumn"`
	MetricName          *string                  `pulumi:"metricName"`
	Operator            *string                  `pulumi:"operator"`
	Query               *string                  `pulumi:"query"`
	ResourceIdColumn    *string                  `pulumi:"resourceIdColumn"`
	Threshold           *float64                 `pulumi:"threshold"`
	TimeAggregation     *string                  `pulumi:"timeAggregation"`
}





type ConditionInput interface {
	pulumi.Input

	ToConditionOutput() ConditionOutput
	ToConditionOutputWithContext(context.Context) ConditionOutput
}

type ConditionArgs struct {
	Dimensions          DimensionArrayInput             `pulumi:"dimensions"`
	FailingPeriods      ConditionFailingPeriodsPtrInput `pulumi:"failingPeriods"`
	MetricMeasureColumn pulumi.StringPtrInput           `pulumi:"metricMeasureColumn"`
	MetricName          pulumi.StringPtrInput           `pulumi:"metricName"`
	Operator            pulumi.StringPtrInput           `pulumi:"operator"`
	Query               pulumi.StringPtrInput           `pulumi:"query"`
	ResourceIdColumn    pulumi.StringPtrInput           `pulumi:"resourceIdColumn"`
	Threshold           pulumi.Float64PtrInput          `pulumi:"threshold"`
	TimeAggregation     pulumi.StringPtrInput           `pulumi:"timeAggregation"`
}

func (ConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Condition)(nil)).Elem()
}

func (i ConditionArgs) ToConditionOutput() ConditionOutput {
	return i.ToConditionOutputWithContext(context.Background())
}

func (i ConditionArgs) ToConditionOutputWithContext(ctx context.Context) ConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionOutput)
}





type ConditionArrayInput interface {
	pulumi.Input

	ToConditionArrayOutput() ConditionArrayOutput
	ToConditionArrayOutputWithContext(context.Context) ConditionArrayOutput
}

type ConditionArray []ConditionInput

func (ConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Condition)(nil)).Elem()
}

func (i ConditionArray) ToConditionArrayOutput() ConditionArrayOutput {
	return i.ToConditionArrayOutputWithContext(context.Background())
}

func (i ConditionArray) ToConditionArrayOutputWithContext(ctx context.Context) ConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionArrayOutput)
}

type ConditionOutput struct{ *pulumi.OutputState }

func (ConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Condition)(nil)).Elem()
}

func (o ConditionOutput) ToConditionOutput() ConditionOutput {
	return o
}

func (o ConditionOutput) ToConditionOutputWithContext(ctx context.Context) ConditionOutput {
	return o
}

func (o ConditionOutput) Dimensions() DimensionArrayOutput {
	return o.ApplyT(func(v Condition) []Dimension { return v.Dimensions }).(DimensionArrayOutput)
}

func (o ConditionOutput) FailingPeriods() ConditionFailingPeriodsPtrOutput {
	return o.ApplyT(func(v Condition) *ConditionFailingPeriods { return v.FailingPeriods }).(ConditionFailingPeriodsPtrOutput)
}

func (o ConditionOutput) MetricMeasureColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.MetricMeasureColumn }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) ResourceIdColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.ResourceIdColumn }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Condition) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

func (o ConditionOutput) TimeAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.TimeAggregation }).(pulumi.StringPtrOutput)
}

type ConditionArrayOutput struct{ *pulumi.OutputState }

func (ConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Condition)(nil)).Elem()
}

func (o ConditionArrayOutput) ToConditionArrayOutput() ConditionArrayOutput {
	return o
}

func (o ConditionArrayOutput) ToConditionArrayOutputWithContext(ctx context.Context) ConditionArrayOutput {
	return o
}

func (o ConditionArrayOutput) Index(i pulumi.IntInput) ConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Condition {
		return vs[0].([]Condition)[vs[1].(int)]
	}).(ConditionOutput)
}

type ConditionFailingPeriods struct {
	MinFailingPeriodsToAlert  *float64 `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods *float64 `pulumi:"numberOfEvaluationPeriods"`
}





type ConditionFailingPeriodsInput interface {
	pulumi.Input

	ToConditionFailingPeriodsOutput() ConditionFailingPeriodsOutput
	ToConditionFailingPeriodsOutputWithContext(context.Context) ConditionFailingPeriodsOutput
}

type ConditionFailingPeriodsArgs struct {
	MinFailingPeriodsToAlert  pulumi.Float64PtrInput `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods pulumi.Float64PtrInput `pulumi:"numberOfEvaluationPeriods"`
}

func (ConditionFailingPeriodsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionFailingPeriods)(nil)).Elem()
}

func (i ConditionFailingPeriodsArgs) ToConditionFailingPeriodsOutput() ConditionFailingPeriodsOutput {
	return i.ToConditionFailingPeriodsOutputWithContext(context.Background())
}

func (i ConditionFailingPeriodsArgs) ToConditionFailingPeriodsOutputWithContext(ctx context.Context) ConditionFailingPeriodsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionFailingPeriodsOutput)
}

func (i ConditionFailingPeriodsArgs) ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput {
	return i.ToConditionFailingPeriodsPtrOutputWithContext(context.Background())
}

func (i ConditionFailingPeriodsArgs) ToConditionFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionFailingPeriodsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionFailingPeriodsOutput).ToConditionFailingPeriodsPtrOutputWithContext(ctx)
}









type ConditionFailingPeriodsPtrInput interface {
	pulumi.Input

	ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput
	ToConditionFailingPeriodsPtrOutputWithContext(context.Context) ConditionFailingPeriodsPtrOutput
}

type conditionFailingPeriodsPtrType ConditionFailingPeriodsArgs

func ConditionFailingPeriodsPtr(v *ConditionFailingPeriodsArgs) ConditionFailingPeriodsPtrInput {
	return (*conditionFailingPeriodsPtrType)(v)
}

func (*conditionFailingPeriodsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionFailingPeriods)(nil)).Elem()
}

func (i *conditionFailingPeriodsPtrType) ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput {
	return i.ToConditionFailingPeriodsPtrOutputWithContext(context.Background())
}

func (i *conditionFailingPeriodsPtrType) ToConditionFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionFailingPeriodsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionFailingPeriodsPtrOutput)
}

type ConditionFailingPeriodsOutput struct{ *pulumi.OutputState }

func (ConditionFailingPeriodsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionFailingPeriods)(nil)).Elem()
}

func (o ConditionFailingPeriodsOutput) ToConditionFailingPeriodsOutput() ConditionFailingPeriodsOutput {
	return o
}

func (o ConditionFailingPeriodsOutput) ToConditionFailingPeriodsOutputWithContext(ctx context.Context) ConditionFailingPeriodsOutput {
	return o
}

func (o ConditionFailingPeriodsOutput) ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput {
	return o.ToConditionFailingPeriodsPtrOutputWithContext(context.Background())
}

func (o ConditionFailingPeriodsOutput) ToConditionFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionFailingPeriodsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConditionFailingPeriods) *ConditionFailingPeriods {
		return &v
	}).(ConditionFailingPeriodsPtrOutput)
}

func (o ConditionFailingPeriodsOutput) MinFailingPeriodsToAlert() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConditionFailingPeriods) *float64 { return v.MinFailingPeriodsToAlert }).(pulumi.Float64PtrOutput)
}

func (o ConditionFailingPeriodsOutput) NumberOfEvaluationPeriods() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConditionFailingPeriods) *float64 { return v.NumberOfEvaluationPeriods }).(pulumi.Float64PtrOutput)
}

type ConditionFailingPeriodsPtrOutput struct{ *pulumi.OutputState }

func (ConditionFailingPeriodsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionFailingPeriods)(nil)).Elem()
}

func (o ConditionFailingPeriodsPtrOutput) ToConditionFailingPeriodsPtrOutput() ConditionFailingPeriodsPtrOutput {
	return o
}

func (o ConditionFailingPeriodsPtrOutput) ToConditionFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionFailingPeriodsPtrOutput {
	return o
}

func (o ConditionFailingPeriodsPtrOutput) Elem() ConditionFailingPeriodsOutput {
	return o.ApplyT(func(v *ConditionFailingPeriods) ConditionFailingPeriods {
		if v != nil {
			return *v
		}
		var ret ConditionFailingPeriods
		return ret
	}).(ConditionFailingPeriodsOutput)
}

func (o ConditionFailingPeriodsPtrOutput) MinFailingPeriodsToAlert() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConditionFailingPeriods) *float64 {
		if v == nil {
			return nil
		}
		return v.MinFailingPeriodsToAlert
	}).(pulumi.Float64PtrOutput)
}

func (o ConditionFailingPeriodsPtrOutput) NumberOfEvaluationPeriods() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConditionFailingPeriods) *float64 {
		if v == nil {
			return nil
		}
		return v.NumberOfEvaluationPeriods
	}).(pulumi.Float64PtrOutput)
}

type ConditionResponse struct {
	Dimensions          []DimensionResponse              `pulumi:"dimensions"`
	FailingPeriods      *ConditionResponseFailingPeriods `pulumi:"failingPeriods"`
	MetricMeasureColumn *string                          `pulumi:"metricMeasureColumn"`
	MetricName          *string                          `pulumi:"metricName"`
	Operator            *string                          `pulumi:"operator"`
	Query               *string                          `pulumi:"query"`
	ResourceIdColumn    *string                          `pulumi:"resourceIdColumn"`
	Threshold           *float64                         `pulumi:"threshold"`
	TimeAggregation     *string                          `pulumi:"timeAggregation"`
}





type ConditionResponseInput interface {
	pulumi.Input

	ToConditionResponseOutput() ConditionResponseOutput
	ToConditionResponseOutputWithContext(context.Context) ConditionResponseOutput
}

type ConditionResponseArgs struct {
	Dimensions          DimensionResponseArrayInput             `pulumi:"dimensions"`
	FailingPeriods      ConditionResponseFailingPeriodsPtrInput `pulumi:"failingPeriods"`
	MetricMeasureColumn pulumi.StringPtrInput                   `pulumi:"metricMeasureColumn"`
	MetricName          pulumi.StringPtrInput                   `pulumi:"metricName"`
	Operator            pulumi.StringPtrInput                   `pulumi:"operator"`
	Query               pulumi.StringPtrInput                   `pulumi:"query"`
	ResourceIdColumn    pulumi.StringPtrInput                   `pulumi:"resourceIdColumn"`
	Threshold           pulumi.Float64PtrInput                  `pulumi:"threshold"`
	TimeAggregation     pulumi.StringPtrInput                   `pulumi:"timeAggregation"`
}

func (ConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponse)(nil)).Elem()
}

func (i ConditionResponseArgs) ToConditionResponseOutput() ConditionResponseOutput {
	return i.ToConditionResponseOutputWithContext(context.Background())
}

func (i ConditionResponseArgs) ToConditionResponseOutputWithContext(ctx context.Context) ConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseOutput)
}





type ConditionResponseArrayInput interface {
	pulumi.Input

	ToConditionResponseArrayOutput() ConditionResponseArrayOutput
	ToConditionResponseArrayOutputWithContext(context.Context) ConditionResponseArrayOutput
}

type ConditionResponseArray []ConditionResponseInput

func (ConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConditionResponse)(nil)).Elem()
}

func (i ConditionResponseArray) ToConditionResponseArrayOutput() ConditionResponseArrayOutput {
	return i.ToConditionResponseArrayOutputWithContext(context.Background())
}

func (i ConditionResponseArray) ToConditionResponseArrayOutputWithContext(ctx context.Context) ConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseArrayOutput)
}

type ConditionResponseOutput struct{ *pulumi.OutputState }

func (ConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponse)(nil)).Elem()
}

func (o ConditionResponseOutput) ToConditionResponseOutput() ConditionResponseOutput {
	return o
}

func (o ConditionResponseOutput) ToConditionResponseOutputWithContext(ctx context.Context) ConditionResponseOutput {
	return o
}

func (o ConditionResponseOutput) Dimensions() DimensionResponseArrayOutput {
	return o.ApplyT(func(v ConditionResponse) []DimensionResponse { return v.Dimensions }).(DimensionResponseArrayOutput)
}

func (o ConditionResponseOutput) FailingPeriods() ConditionResponseFailingPeriodsPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *ConditionResponseFailingPeriods { return v.FailingPeriods }).(ConditionResponseFailingPeriodsPtrOutput)
}

func (o ConditionResponseOutput) MetricMeasureColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.MetricMeasureColumn }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) ResourceIdColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.ResourceIdColumn }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConditionResponse) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

func (o ConditionResponseOutput) TimeAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.TimeAggregation }).(pulumi.StringPtrOutput)
}

type ConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (ConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConditionResponse)(nil)).Elem()
}

func (o ConditionResponseArrayOutput) ToConditionResponseArrayOutput() ConditionResponseArrayOutput {
	return o
}

func (o ConditionResponseArrayOutput) ToConditionResponseArrayOutputWithContext(ctx context.Context) ConditionResponseArrayOutput {
	return o
}

func (o ConditionResponseArrayOutput) Index(i pulumi.IntInput) ConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConditionResponse {
		return vs[0].([]ConditionResponse)[vs[1].(int)]
	}).(ConditionResponseOutput)
}

type ConditionResponseFailingPeriods struct {
	MinFailingPeriodsToAlert  *float64 `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods *float64 `pulumi:"numberOfEvaluationPeriods"`
}





type ConditionResponseFailingPeriodsInput interface {
	pulumi.Input

	ToConditionResponseFailingPeriodsOutput() ConditionResponseFailingPeriodsOutput
	ToConditionResponseFailingPeriodsOutputWithContext(context.Context) ConditionResponseFailingPeriodsOutput
}

type ConditionResponseFailingPeriodsArgs struct {
	MinFailingPeriodsToAlert  pulumi.Float64PtrInput `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods pulumi.Float64PtrInput `pulumi:"numberOfEvaluationPeriods"`
}

func (ConditionResponseFailingPeriodsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponseFailingPeriods)(nil)).Elem()
}

func (i ConditionResponseFailingPeriodsArgs) ToConditionResponseFailingPeriodsOutput() ConditionResponseFailingPeriodsOutput {
	return i.ToConditionResponseFailingPeriodsOutputWithContext(context.Background())
}

func (i ConditionResponseFailingPeriodsArgs) ToConditionResponseFailingPeriodsOutputWithContext(ctx context.Context) ConditionResponseFailingPeriodsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseFailingPeriodsOutput)
}

func (i ConditionResponseFailingPeriodsArgs) ToConditionResponseFailingPeriodsPtrOutput() ConditionResponseFailingPeriodsPtrOutput {
	return i.ToConditionResponseFailingPeriodsPtrOutputWithContext(context.Background())
}

func (i ConditionResponseFailingPeriodsArgs) ToConditionResponseFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionResponseFailingPeriodsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseFailingPeriodsOutput).ToConditionResponseFailingPeriodsPtrOutputWithContext(ctx)
}









type ConditionResponseFailingPeriodsPtrInput interface {
	pulumi.Input

	ToConditionResponseFailingPeriodsPtrOutput() ConditionResponseFailingPeriodsPtrOutput
	ToConditionResponseFailingPeriodsPtrOutputWithContext(context.Context) ConditionResponseFailingPeriodsPtrOutput
}

type conditionResponseFailingPeriodsPtrType ConditionResponseFailingPeriodsArgs

func ConditionResponseFailingPeriodsPtr(v *ConditionResponseFailingPeriodsArgs) ConditionResponseFailingPeriodsPtrInput {
	return (*conditionResponseFailingPeriodsPtrType)(v)
}

func (*conditionResponseFailingPeriodsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionResponseFailingPeriods)(nil)).Elem()
}

func (i *conditionResponseFailingPeriodsPtrType) ToConditionResponseFailingPeriodsPtrOutput() ConditionResponseFailingPeriodsPtrOutput {
	return i.ToConditionResponseFailingPeriodsPtrOutputWithContext(context.Background())
}

func (i *conditionResponseFailingPeriodsPtrType) ToConditionResponseFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionResponseFailingPeriodsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseFailingPeriodsPtrOutput)
}

type ConditionResponseFailingPeriodsOutput struct{ *pulumi.OutputState }

func (ConditionResponseFailingPeriodsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponseFailingPeriods)(nil)).Elem()
}

func (o ConditionResponseFailingPeriodsOutput) ToConditionResponseFailingPeriodsOutput() ConditionResponseFailingPeriodsOutput {
	return o
}

func (o ConditionResponseFailingPeriodsOutput) ToConditionResponseFailingPeriodsOutputWithContext(ctx context.Context) ConditionResponseFailingPeriodsOutput {
	return o
}

func (o ConditionResponseFailingPeriodsOutput) ToConditionResponseFailingPeriodsPtrOutput() ConditionResponseFailingPeriodsPtrOutput {
	return o.ToConditionResponseFailingPeriodsPtrOutputWithContext(context.Background())
}

func (o ConditionResponseFailingPeriodsOutput) ToConditionResponseFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionResponseFailingPeriodsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConditionResponseFailingPeriods) *ConditionResponseFailingPeriods {
		return &v
	}).(ConditionResponseFailingPeriodsPtrOutput)
}

func (o ConditionResponseFailingPeriodsOutput) MinFailingPeriodsToAlert() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConditionResponseFailingPeriods) *float64 { return v.MinFailingPeriodsToAlert }).(pulumi.Float64PtrOutput)
}

func (o ConditionResponseFailingPeriodsOutput) NumberOfEvaluationPeriods() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConditionResponseFailingPeriods) *float64 { return v.NumberOfEvaluationPeriods }).(pulumi.Float64PtrOutput)
}

type ConditionResponseFailingPeriodsPtrOutput struct{ *pulumi.OutputState }

func (ConditionResponseFailingPeriodsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionResponseFailingPeriods)(nil)).Elem()
}

func (o ConditionResponseFailingPeriodsPtrOutput) ToConditionResponseFailingPeriodsPtrOutput() ConditionResponseFailingPeriodsPtrOutput {
	return o
}

func (o ConditionResponseFailingPeriodsPtrOutput) ToConditionResponseFailingPeriodsPtrOutputWithContext(ctx context.Context) ConditionResponseFailingPeriodsPtrOutput {
	return o
}

func (o ConditionResponseFailingPeriodsPtrOutput) Elem() ConditionResponseFailingPeriodsOutput {
	return o.ApplyT(func(v *ConditionResponseFailingPeriods) ConditionResponseFailingPeriods {
		if v != nil {
			return *v
		}
		var ret ConditionResponseFailingPeriods
		return ret
	}).(ConditionResponseFailingPeriodsOutput)
}

func (o ConditionResponseFailingPeriodsPtrOutput) MinFailingPeriodsToAlert() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConditionResponseFailingPeriods) *float64 {
		if v == nil {
			return nil
		}
		return v.MinFailingPeriodsToAlert
	}).(pulumi.Float64PtrOutput)
}

func (o ConditionResponseFailingPeriodsPtrOutput) NumberOfEvaluationPeriods() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConditionResponseFailingPeriods) *float64 {
		if v == nil {
			return nil
		}
		return v.NumberOfEvaluationPeriods
	}).(pulumi.Float64PtrOutput)
}

type Dimension struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type DimensionInput interface {
	pulumi.Input

	ToDimensionOutput() DimensionOutput
	ToDimensionOutputWithContext(context.Context) DimensionOutput
}

type DimensionArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (DimensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Dimension)(nil)).Elem()
}

func (i DimensionArgs) ToDimensionOutput() DimensionOutput {
	return i.ToDimensionOutputWithContext(context.Background())
}

func (i DimensionArgs) ToDimensionOutputWithContext(ctx context.Context) DimensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DimensionOutput)
}





type DimensionArrayInput interface {
	pulumi.Input

	ToDimensionArrayOutput() DimensionArrayOutput
	ToDimensionArrayOutputWithContext(context.Context) DimensionArrayOutput
}

type DimensionArray []DimensionInput

func (DimensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Dimension)(nil)).Elem()
}

func (i DimensionArray) ToDimensionArrayOutput() DimensionArrayOutput {
	return i.ToDimensionArrayOutputWithContext(context.Background())
}

func (i DimensionArray) ToDimensionArrayOutputWithContext(ctx context.Context) DimensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DimensionArrayOutput)
}

type DimensionOutput struct{ *pulumi.OutputState }

func (DimensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dimension)(nil)).Elem()
}

func (o DimensionOutput) ToDimensionOutput() DimensionOutput {
	return o
}

func (o DimensionOutput) ToDimensionOutputWithContext(ctx context.Context) DimensionOutput {
	return o
}

func (o DimensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Dimension) string { return v.Name }).(pulumi.StringOutput)
}

func (o DimensionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v Dimension) string { return v.Operator }).(pulumi.StringOutput)
}

func (o DimensionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Dimension) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type DimensionArrayOutput struct{ *pulumi.OutputState }

func (DimensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Dimension)(nil)).Elem()
}

func (o DimensionArrayOutput) ToDimensionArrayOutput() DimensionArrayOutput {
	return o
}

func (o DimensionArrayOutput) ToDimensionArrayOutputWithContext(ctx context.Context) DimensionArrayOutput {
	return o
}

func (o DimensionArrayOutput) Index(i pulumi.IntInput) DimensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Dimension {
		return vs[0].([]Dimension)[vs[1].(int)]
	}).(DimensionOutput)
}

type DimensionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type DimensionResponseInput interface {
	pulumi.Input

	ToDimensionResponseOutput() DimensionResponseOutput
	ToDimensionResponseOutputWithContext(context.Context) DimensionResponseOutput
}

type DimensionResponseArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (DimensionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DimensionResponse)(nil)).Elem()
}

func (i DimensionResponseArgs) ToDimensionResponseOutput() DimensionResponseOutput {
	return i.ToDimensionResponseOutputWithContext(context.Background())
}

func (i DimensionResponseArgs) ToDimensionResponseOutputWithContext(ctx context.Context) DimensionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DimensionResponseOutput)
}





type DimensionResponseArrayInput interface {
	pulumi.Input

	ToDimensionResponseArrayOutput() DimensionResponseArrayOutput
	ToDimensionResponseArrayOutputWithContext(context.Context) DimensionResponseArrayOutput
}

type DimensionResponseArray []DimensionResponseInput

func (DimensionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DimensionResponse)(nil)).Elem()
}

func (i DimensionResponseArray) ToDimensionResponseArrayOutput() DimensionResponseArrayOutput {
	return i.ToDimensionResponseArrayOutputWithContext(context.Background())
}

func (i DimensionResponseArray) ToDimensionResponseArrayOutputWithContext(ctx context.Context) DimensionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DimensionResponseArrayOutput)
}

type DimensionResponseOutput struct{ *pulumi.OutputState }

func (DimensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DimensionResponse)(nil)).Elem()
}

func (o DimensionResponseOutput) ToDimensionResponseOutput() DimensionResponseOutput {
	return o
}

func (o DimensionResponseOutput) ToDimensionResponseOutputWithContext(ctx context.Context) DimensionResponseOutput {
	return o
}

func (o DimensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DimensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DimensionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v DimensionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o DimensionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DimensionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type DimensionResponseArrayOutput struct{ *pulumi.OutputState }

func (DimensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DimensionResponse)(nil)).Elem()
}

func (o DimensionResponseArrayOutput) ToDimensionResponseArrayOutput() DimensionResponseArrayOutput {
	return o
}

func (o DimensionResponseArrayOutput) ToDimensionResponseArrayOutputWithContext(ctx context.Context) DimensionResponseArrayOutput {
	return o
}

func (o DimensionResponseArrayOutput) Index(i pulumi.IntInput) DimensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DimensionResponse {
		return vs[0].([]DimensionResponse)[vs[1].(int)]
	}).(DimensionResponseOutput)
}

type ScheduledQueryRuleCriteria struct {
	AllOf []Condition `pulumi:"allOf"`
}





type ScheduledQueryRuleCriteriaInput interface {
	pulumi.Input

	ToScheduledQueryRuleCriteriaOutput() ScheduledQueryRuleCriteriaOutput
	ToScheduledQueryRuleCriteriaOutputWithContext(context.Context) ScheduledQueryRuleCriteriaOutput
}

type ScheduledQueryRuleCriteriaArgs struct {
	AllOf ConditionArrayInput `pulumi:"allOf"`
}

func (ScheduledQueryRuleCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRuleCriteria)(nil)).Elem()
}

func (i ScheduledQueryRuleCriteriaArgs) ToScheduledQueryRuleCriteriaOutput() ScheduledQueryRuleCriteriaOutput {
	return i.ToScheduledQueryRuleCriteriaOutputWithContext(context.Background())
}

func (i ScheduledQueryRuleCriteriaArgs) ToScheduledQueryRuleCriteriaOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleCriteriaOutput)
}

func (i ScheduledQueryRuleCriteriaArgs) ToScheduledQueryRuleCriteriaPtrOutput() ScheduledQueryRuleCriteriaPtrOutput {
	return i.ToScheduledQueryRuleCriteriaPtrOutputWithContext(context.Background())
}

func (i ScheduledQueryRuleCriteriaArgs) ToScheduledQueryRuleCriteriaPtrOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleCriteriaOutput).ToScheduledQueryRuleCriteriaPtrOutputWithContext(ctx)
}









type ScheduledQueryRuleCriteriaPtrInput interface {
	pulumi.Input

	ToScheduledQueryRuleCriteriaPtrOutput() ScheduledQueryRuleCriteriaPtrOutput
	ToScheduledQueryRuleCriteriaPtrOutputWithContext(context.Context) ScheduledQueryRuleCriteriaPtrOutput
}

type scheduledQueryRuleCriteriaPtrType ScheduledQueryRuleCriteriaArgs

func ScheduledQueryRuleCriteriaPtr(v *ScheduledQueryRuleCriteriaArgs) ScheduledQueryRuleCriteriaPtrInput {
	return (*scheduledQueryRuleCriteriaPtrType)(v)
}

func (*scheduledQueryRuleCriteriaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryRuleCriteria)(nil)).Elem()
}

func (i *scheduledQueryRuleCriteriaPtrType) ToScheduledQueryRuleCriteriaPtrOutput() ScheduledQueryRuleCriteriaPtrOutput {
	return i.ToScheduledQueryRuleCriteriaPtrOutputWithContext(context.Background())
}

func (i *scheduledQueryRuleCriteriaPtrType) ToScheduledQueryRuleCriteriaPtrOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleCriteriaPtrOutput)
}

type ScheduledQueryRuleCriteriaOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRuleCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRuleCriteria)(nil)).Elem()
}

func (o ScheduledQueryRuleCriteriaOutput) ToScheduledQueryRuleCriteriaOutput() ScheduledQueryRuleCriteriaOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaOutput) ToScheduledQueryRuleCriteriaOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaOutput) ToScheduledQueryRuleCriteriaPtrOutput() ScheduledQueryRuleCriteriaPtrOutput {
	return o.ToScheduledQueryRuleCriteriaPtrOutputWithContext(context.Background())
}

func (o ScheduledQueryRuleCriteriaOutput) ToScheduledQueryRuleCriteriaPtrOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledQueryRuleCriteria) *ScheduledQueryRuleCriteria {
		return &v
	}).(ScheduledQueryRuleCriteriaPtrOutput)
}

func (o ScheduledQueryRuleCriteriaOutput) AllOf() ConditionArrayOutput {
	return o.ApplyT(func(v ScheduledQueryRuleCriteria) []Condition { return v.AllOf }).(ConditionArrayOutput)
}

type ScheduledQueryRuleCriteriaPtrOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRuleCriteriaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryRuleCriteria)(nil)).Elem()
}

func (o ScheduledQueryRuleCriteriaPtrOutput) ToScheduledQueryRuleCriteriaPtrOutput() ScheduledQueryRuleCriteriaPtrOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaPtrOutput) ToScheduledQueryRuleCriteriaPtrOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaPtrOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaPtrOutput) Elem() ScheduledQueryRuleCriteriaOutput {
	return o.ApplyT(func(v *ScheduledQueryRuleCriteria) ScheduledQueryRuleCriteria {
		if v != nil {
			return *v
		}
		var ret ScheduledQueryRuleCriteria
		return ret
	}).(ScheduledQueryRuleCriteriaOutput)
}

func (o ScheduledQueryRuleCriteriaPtrOutput) AllOf() ConditionArrayOutput {
	return o.ApplyT(func(v *ScheduledQueryRuleCriteria) []Condition {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(ConditionArrayOutput)
}

type ScheduledQueryRuleCriteriaResponse struct {
	AllOf []ConditionResponse `pulumi:"allOf"`
}





type ScheduledQueryRuleCriteriaResponseInput interface {
	pulumi.Input

	ToScheduledQueryRuleCriteriaResponseOutput() ScheduledQueryRuleCriteriaResponseOutput
	ToScheduledQueryRuleCriteriaResponseOutputWithContext(context.Context) ScheduledQueryRuleCriteriaResponseOutput
}

type ScheduledQueryRuleCriteriaResponseArgs struct {
	AllOf ConditionResponseArrayInput `pulumi:"allOf"`
}

func (ScheduledQueryRuleCriteriaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRuleCriteriaResponse)(nil)).Elem()
}

func (i ScheduledQueryRuleCriteriaResponseArgs) ToScheduledQueryRuleCriteriaResponseOutput() ScheduledQueryRuleCriteriaResponseOutput {
	return i.ToScheduledQueryRuleCriteriaResponseOutputWithContext(context.Background())
}

func (i ScheduledQueryRuleCriteriaResponseArgs) ToScheduledQueryRuleCriteriaResponseOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleCriteriaResponseOutput)
}

func (i ScheduledQueryRuleCriteriaResponseArgs) ToScheduledQueryRuleCriteriaResponsePtrOutput() ScheduledQueryRuleCriteriaResponsePtrOutput {
	return i.ToScheduledQueryRuleCriteriaResponsePtrOutputWithContext(context.Background())
}

func (i ScheduledQueryRuleCriteriaResponseArgs) ToScheduledQueryRuleCriteriaResponsePtrOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleCriteriaResponseOutput).ToScheduledQueryRuleCriteriaResponsePtrOutputWithContext(ctx)
}









type ScheduledQueryRuleCriteriaResponsePtrInput interface {
	pulumi.Input

	ToScheduledQueryRuleCriteriaResponsePtrOutput() ScheduledQueryRuleCriteriaResponsePtrOutput
	ToScheduledQueryRuleCriteriaResponsePtrOutputWithContext(context.Context) ScheduledQueryRuleCriteriaResponsePtrOutput
}

type scheduledQueryRuleCriteriaResponsePtrType ScheduledQueryRuleCriteriaResponseArgs

func ScheduledQueryRuleCriteriaResponsePtr(v *ScheduledQueryRuleCriteriaResponseArgs) ScheduledQueryRuleCriteriaResponsePtrInput {
	return (*scheduledQueryRuleCriteriaResponsePtrType)(v)
}

func (*scheduledQueryRuleCriteriaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryRuleCriteriaResponse)(nil)).Elem()
}

func (i *scheduledQueryRuleCriteriaResponsePtrType) ToScheduledQueryRuleCriteriaResponsePtrOutput() ScheduledQueryRuleCriteriaResponsePtrOutput {
	return i.ToScheduledQueryRuleCriteriaResponsePtrOutputWithContext(context.Background())
}

func (i *scheduledQueryRuleCriteriaResponsePtrType) ToScheduledQueryRuleCriteriaResponsePtrOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledQueryRuleCriteriaResponsePtrOutput)
}

type ScheduledQueryRuleCriteriaResponseOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRuleCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledQueryRuleCriteriaResponse)(nil)).Elem()
}

func (o ScheduledQueryRuleCriteriaResponseOutput) ToScheduledQueryRuleCriteriaResponseOutput() ScheduledQueryRuleCriteriaResponseOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaResponseOutput) ToScheduledQueryRuleCriteriaResponseOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaResponseOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaResponseOutput) ToScheduledQueryRuleCriteriaResponsePtrOutput() ScheduledQueryRuleCriteriaResponsePtrOutput {
	return o.ToScheduledQueryRuleCriteriaResponsePtrOutputWithContext(context.Background())
}

func (o ScheduledQueryRuleCriteriaResponseOutput) ToScheduledQueryRuleCriteriaResponsePtrOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledQueryRuleCriteriaResponse) *ScheduledQueryRuleCriteriaResponse {
		return &v
	}).(ScheduledQueryRuleCriteriaResponsePtrOutput)
}

func (o ScheduledQueryRuleCriteriaResponseOutput) AllOf() ConditionResponseArrayOutput {
	return o.ApplyT(func(v ScheduledQueryRuleCriteriaResponse) []ConditionResponse { return v.AllOf }).(ConditionResponseArrayOutput)
}

type ScheduledQueryRuleCriteriaResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduledQueryRuleCriteriaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledQueryRuleCriteriaResponse)(nil)).Elem()
}

func (o ScheduledQueryRuleCriteriaResponsePtrOutput) ToScheduledQueryRuleCriteriaResponsePtrOutput() ScheduledQueryRuleCriteriaResponsePtrOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaResponsePtrOutput) ToScheduledQueryRuleCriteriaResponsePtrOutputWithContext(ctx context.Context) ScheduledQueryRuleCriteriaResponsePtrOutput {
	return o
}

func (o ScheduledQueryRuleCriteriaResponsePtrOutput) Elem() ScheduledQueryRuleCriteriaResponseOutput {
	return o.ApplyT(func(v *ScheduledQueryRuleCriteriaResponse) ScheduledQueryRuleCriteriaResponse {
		if v != nil {
			return *v
		}
		var ret ScheduledQueryRuleCriteriaResponse
		return ret
	}).(ScheduledQueryRuleCriteriaResponseOutput)
}

func (o ScheduledQueryRuleCriteriaResponsePtrOutput) AllOf() ConditionResponseArrayOutput {
	return o.ApplyT(func(v *ScheduledQueryRuleCriteriaResponse) []ConditionResponse {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(ConditionResponseArrayOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionsOutput{})
	pulumi.RegisterOutputType(ActionsPtrOutput{})
	pulumi.RegisterOutputType(ActionsResponseOutput{})
	pulumi.RegisterOutputType(ActionsResponsePtrOutput{})
	pulumi.RegisterOutputType(ConditionOutput{})
	pulumi.RegisterOutputType(ConditionArrayOutput{})
	pulumi.RegisterOutputType(ConditionFailingPeriodsOutput{})
	pulumi.RegisterOutputType(ConditionFailingPeriodsPtrOutput{})
	pulumi.RegisterOutputType(ConditionResponseOutput{})
	pulumi.RegisterOutputType(ConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(ConditionResponseFailingPeriodsOutput{})
	pulumi.RegisterOutputType(ConditionResponseFailingPeriodsPtrOutput{})
	pulumi.RegisterOutputType(DimensionOutput{})
	pulumi.RegisterOutputType(DimensionArrayOutput{})
	pulumi.RegisterOutputType(DimensionResponseOutput{})
	pulumi.RegisterOutputType(DimensionResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaPtrOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaResponseOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
