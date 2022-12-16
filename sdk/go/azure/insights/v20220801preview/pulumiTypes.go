


package v20220801preview

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


func (val *Condition) Defaults() *Condition {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FailingPeriods = tmp.FailingPeriods.Defaults()

	return &tmp
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


func (val *ConditionArgs) Defaults() *ConditionArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
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


func (val *ConditionFailingPeriods) Defaults() *ConditionFailingPeriods {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MinFailingPeriodsToAlert) {
		minFailingPeriodsToAlert_ := 1.0
		tmp.MinFailingPeriodsToAlert = &minFailingPeriodsToAlert_
	}
	if isZero(tmp.NumberOfEvaluationPeriods) {
		numberOfEvaluationPeriods_ := 1.0
		tmp.NumberOfEvaluationPeriods = &numberOfEvaluationPeriods_
	}
	return &tmp
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


func (val *ConditionFailingPeriodsArgs) Defaults() *ConditionFailingPeriodsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MinFailingPeriodsToAlert) {
		tmp.MinFailingPeriodsToAlert = pulumi.Float64Ptr(1.0)
	}
	if isZero(tmp.NumberOfEvaluationPeriods) {
		tmp.NumberOfEvaluationPeriods = pulumi.Float64Ptr(1.0)
	}
	return &tmp
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


func (val *ConditionResponse) Defaults() *ConditionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FailingPeriods = tmp.FailingPeriods.Defaults()

	return &tmp
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


func (val *ConditionResponseFailingPeriods) Defaults() *ConditionResponseFailingPeriods {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MinFailingPeriodsToAlert) {
		minFailingPeriodsToAlert_ := 1.0
		tmp.MinFailingPeriodsToAlert = &minFailingPeriodsToAlert_
	}
	if isZero(tmp.NumberOfEvaluationPeriods) {
		numberOfEvaluationPeriods_ := 1.0
		tmp.NumberOfEvaluationPeriods = &numberOfEvaluationPeriods_
	}
	return &tmp
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

type Identity struct {
	Type                   IdentityType           `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   IdentityTypeInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput   `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() IdentityTypeOutput {
	return o.ApplyT(func(v Identity) IdentityType { return v.Type }).(IdentityTypeOutput)
}

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() IdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *IdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(IdentityTypePtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                    `pulumi:"principalId"`
	TenantId               string                                    `pulumi:"tenantId"`
	Type                   string                                    `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityPropertiesResponse `pulumi:"userAssignedIdentities"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]UserIdentityPropertiesResponse { return v.UserAssignedIdentities }).(UserIdentityPropertiesResponseMapOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]UserIdentityPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type RuleResolveConfiguration struct {
	AutoResolved  *bool   `pulumi:"autoResolved"`
	TimeToResolve *string `pulumi:"timeToResolve"`
}





type RuleResolveConfigurationInput interface {
	pulumi.Input

	ToRuleResolveConfigurationOutput() RuleResolveConfigurationOutput
	ToRuleResolveConfigurationOutputWithContext(context.Context) RuleResolveConfigurationOutput
}

type RuleResolveConfigurationArgs struct {
	AutoResolved  pulumi.BoolPtrInput   `pulumi:"autoResolved"`
	TimeToResolve pulumi.StringPtrInput `pulumi:"timeToResolve"`
}

func (RuleResolveConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleResolveConfiguration)(nil)).Elem()
}

func (i RuleResolveConfigurationArgs) ToRuleResolveConfigurationOutput() RuleResolveConfigurationOutput {
	return i.ToRuleResolveConfigurationOutputWithContext(context.Background())
}

func (i RuleResolveConfigurationArgs) ToRuleResolveConfigurationOutputWithContext(ctx context.Context) RuleResolveConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleResolveConfigurationOutput)
}

func (i RuleResolveConfigurationArgs) ToRuleResolveConfigurationPtrOutput() RuleResolveConfigurationPtrOutput {
	return i.ToRuleResolveConfigurationPtrOutputWithContext(context.Background())
}

func (i RuleResolveConfigurationArgs) ToRuleResolveConfigurationPtrOutputWithContext(ctx context.Context) RuleResolveConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleResolveConfigurationOutput).ToRuleResolveConfigurationPtrOutputWithContext(ctx)
}









type RuleResolveConfigurationPtrInput interface {
	pulumi.Input

	ToRuleResolveConfigurationPtrOutput() RuleResolveConfigurationPtrOutput
	ToRuleResolveConfigurationPtrOutputWithContext(context.Context) RuleResolveConfigurationPtrOutput
}

type ruleResolveConfigurationPtrType RuleResolveConfigurationArgs

func RuleResolveConfigurationPtr(v *RuleResolveConfigurationArgs) RuleResolveConfigurationPtrInput {
	return (*ruleResolveConfigurationPtrType)(v)
}

func (*ruleResolveConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleResolveConfiguration)(nil)).Elem()
}

func (i *ruleResolveConfigurationPtrType) ToRuleResolveConfigurationPtrOutput() RuleResolveConfigurationPtrOutput {
	return i.ToRuleResolveConfigurationPtrOutputWithContext(context.Background())
}

func (i *ruleResolveConfigurationPtrType) ToRuleResolveConfigurationPtrOutputWithContext(ctx context.Context) RuleResolveConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleResolveConfigurationPtrOutput)
}

type RuleResolveConfigurationOutput struct{ *pulumi.OutputState }

func (RuleResolveConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleResolveConfiguration)(nil)).Elem()
}

func (o RuleResolveConfigurationOutput) ToRuleResolveConfigurationOutput() RuleResolveConfigurationOutput {
	return o
}

func (o RuleResolveConfigurationOutput) ToRuleResolveConfigurationOutputWithContext(ctx context.Context) RuleResolveConfigurationOutput {
	return o
}

func (o RuleResolveConfigurationOutput) ToRuleResolveConfigurationPtrOutput() RuleResolveConfigurationPtrOutput {
	return o.ToRuleResolveConfigurationPtrOutputWithContext(context.Background())
}

func (o RuleResolveConfigurationOutput) ToRuleResolveConfigurationPtrOutputWithContext(ctx context.Context) RuleResolveConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleResolveConfiguration) *RuleResolveConfiguration {
		return &v
	}).(RuleResolveConfigurationPtrOutput)
}

func (o RuleResolveConfigurationOutput) AutoResolved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleResolveConfiguration) *bool { return v.AutoResolved }).(pulumi.BoolPtrOutput)
}

func (o RuleResolveConfigurationOutput) TimeToResolve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleResolveConfiguration) *string { return v.TimeToResolve }).(pulumi.StringPtrOutput)
}

type RuleResolveConfigurationPtrOutput struct{ *pulumi.OutputState }

func (RuleResolveConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleResolveConfiguration)(nil)).Elem()
}

func (o RuleResolveConfigurationPtrOutput) ToRuleResolveConfigurationPtrOutput() RuleResolveConfigurationPtrOutput {
	return o
}

func (o RuleResolveConfigurationPtrOutput) ToRuleResolveConfigurationPtrOutputWithContext(ctx context.Context) RuleResolveConfigurationPtrOutput {
	return o
}

func (o RuleResolveConfigurationPtrOutput) Elem() RuleResolveConfigurationOutput {
	return o.ApplyT(func(v *RuleResolveConfiguration) RuleResolveConfiguration {
		if v != nil {
			return *v
		}
		var ret RuleResolveConfiguration
		return ret
	}).(RuleResolveConfigurationOutput)
}

func (o RuleResolveConfigurationPtrOutput) AutoResolved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RuleResolveConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.AutoResolved
	}).(pulumi.BoolPtrOutput)
}

func (o RuleResolveConfigurationPtrOutput) TimeToResolve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleResolveConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TimeToResolve
	}).(pulumi.StringPtrOutput)
}

type RuleResolveConfigurationResponse struct {
	AutoResolved  *bool   `pulumi:"autoResolved"`
	TimeToResolve *string `pulumi:"timeToResolve"`
}

type RuleResolveConfigurationResponseOutput struct{ *pulumi.OutputState }

func (RuleResolveConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleResolveConfigurationResponse)(nil)).Elem()
}

func (o RuleResolveConfigurationResponseOutput) ToRuleResolveConfigurationResponseOutput() RuleResolveConfigurationResponseOutput {
	return o
}

func (o RuleResolveConfigurationResponseOutput) ToRuleResolveConfigurationResponseOutputWithContext(ctx context.Context) RuleResolveConfigurationResponseOutput {
	return o
}

func (o RuleResolveConfigurationResponseOutput) AutoResolved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleResolveConfigurationResponse) *bool { return v.AutoResolved }).(pulumi.BoolPtrOutput)
}

func (o RuleResolveConfigurationResponseOutput) TimeToResolve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleResolveConfigurationResponse) *string { return v.TimeToResolve }).(pulumi.StringPtrOutput)
}

type RuleResolveConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (RuleResolveConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleResolveConfigurationResponse)(nil)).Elem()
}

func (o RuleResolveConfigurationResponsePtrOutput) ToRuleResolveConfigurationResponsePtrOutput() RuleResolveConfigurationResponsePtrOutput {
	return o
}

func (o RuleResolveConfigurationResponsePtrOutput) ToRuleResolveConfigurationResponsePtrOutputWithContext(ctx context.Context) RuleResolveConfigurationResponsePtrOutput {
	return o
}

func (o RuleResolveConfigurationResponsePtrOutput) Elem() RuleResolveConfigurationResponseOutput {
	return o.ApplyT(func(v *RuleResolveConfigurationResponse) RuleResolveConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret RuleResolveConfigurationResponse
		return ret
	}).(RuleResolveConfigurationResponseOutput)
}

func (o RuleResolveConfigurationResponsePtrOutput) AutoResolved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RuleResolveConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutoResolved
	}).(pulumi.BoolPtrOutput)
}

func (o RuleResolveConfigurationResponsePtrOutput) TimeToResolve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleResolveConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeToResolve
	}).(pulumi.StringPtrOutput)
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

func (o ScheduledQueryRuleCriteriaOutput) AllOf() ConditionArrayOutput {
	return o.ApplyT(func(v ScheduledQueryRuleCriteria) []Condition { return v.AllOf }).(ConditionArrayOutput)
}

type ScheduledQueryRuleCriteriaResponse struct {
	AllOf []ConditionResponse `pulumi:"allOf"`
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

func (o ScheduledQueryRuleCriteriaResponseOutput) AllOf() ConditionResponseArrayOutput {
	return o.ApplyT(func(v ScheduledQueryRuleCriteriaResponse) []ConditionResponse { return v.AllOf }).(ConditionResponseArrayOutput)
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

type UserIdentityPropertiesResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserIdentityPropertiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserIdentityPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityPropertiesResponse {
		return vs[0].(map[string]UserIdentityPropertiesResponse)[vs[1].(string)]
	}).(UserIdentityPropertiesResponseOutput)
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
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(RuleResolveConfigurationOutput{})
	pulumi.RegisterOutputType(RuleResolveConfigurationPtrOutput{})
	pulumi.RegisterOutputType(RuleResolveConfigurationResponseOutput{})
	pulumi.RegisterOutputType(RuleResolveConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseMapOutput{})
}
