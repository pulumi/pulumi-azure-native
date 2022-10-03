


package v20220615

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

type HeaderField struct {
	HeaderFieldName  *string `pulumi:"headerFieldName"`
	HeaderFieldValue *string `pulumi:"headerFieldValue"`
}





type HeaderFieldInput interface {
	pulumi.Input

	ToHeaderFieldOutput() HeaderFieldOutput
	ToHeaderFieldOutputWithContext(context.Context) HeaderFieldOutput
}

type HeaderFieldArgs struct {
	HeaderFieldName  pulumi.StringPtrInput `pulumi:"headerFieldName"`
	HeaderFieldValue pulumi.StringPtrInput `pulumi:"headerFieldValue"`
}

func (HeaderFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderField)(nil)).Elem()
}

func (i HeaderFieldArgs) ToHeaderFieldOutput() HeaderFieldOutput {
	return i.ToHeaderFieldOutputWithContext(context.Background())
}

func (i HeaderFieldArgs) ToHeaderFieldOutputWithContext(ctx context.Context) HeaderFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeaderFieldOutput)
}





type HeaderFieldArrayInput interface {
	pulumi.Input

	ToHeaderFieldArrayOutput() HeaderFieldArrayOutput
	ToHeaderFieldArrayOutputWithContext(context.Context) HeaderFieldArrayOutput
}

type HeaderFieldArray []HeaderFieldInput

func (HeaderFieldArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HeaderField)(nil)).Elem()
}

func (i HeaderFieldArray) ToHeaderFieldArrayOutput() HeaderFieldArrayOutput {
	return i.ToHeaderFieldArrayOutputWithContext(context.Background())
}

func (i HeaderFieldArray) ToHeaderFieldArrayOutputWithContext(ctx context.Context) HeaderFieldArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeaderFieldArrayOutput)
}

type HeaderFieldOutput struct{ *pulumi.OutputState }

func (HeaderFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderField)(nil)).Elem()
}

func (o HeaderFieldOutput) ToHeaderFieldOutput() HeaderFieldOutput {
	return o
}

func (o HeaderFieldOutput) ToHeaderFieldOutputWithContext(ctx context.Context) HeaderFieldOutput {
	return o
}

func (o HeaderFieldOutput) HeaderFieldName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderField) *string { return v.HeaderFieldName }).(pulumi.StringPtrOutput)
}

func (o HeaderFieldOutput) HeaderFieldValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderField) *string { return v.HeaderFieldValue }).(pulumi.StringPtrOutput)
}

type HeaderFieldArrayOutput struct{ *pulumi.OutputState }

func (HeaderFieldArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HeaderField)(nil)).Elem()
}

func (o HeaderFieldArrayOutput) ToHeaderFieldArrayOutput() HeaderFieldArrayOutput {
	return o
}

func (o HeaderFieldArrayOutput) ToHeaderFieldArrayOutputWithContext(ctx context.Context) HeaderFieldArrayOutput {
	return o
}

func (o HeaderFieldArrayOutput) Index(i pulumi.IntInput) HeaderFieldOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HeaderField {
		return vs[0].([]HeaderField)[vs[1].(int)]
	}).(HeaderFieldOutput)
}

type HeaderFieldResponse struct {
	HeaderFieldName  *string `pulumi:"headerFieldName"`
	HeaderFieldValue *string `pulumi:"headerFieldValue"`
}

type HeaderFieldResponseOutput struct{ *pulumi.OutputState }

func (HeaderFieldResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderFieldResponse)(nil)).Elem()
}

func (o HeaderFieldResponseOutput) ToHeaderFieldResponseOutput() HeaderFieldResponseOutput {
	return o
}

func (o HeaderFieldResponseOutput) ToHeaderFieldResponseOutputWithContext(ctx context.Context) HeaderFieldResponseOutput {
	return o
}

func (o HeaderFieldResponseOutput) HeaderFieldName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderFieldResponse) *string { return v.HeaderFieldName }).(pulumi.StringPtrOutput)
}

func (o HeaderFieldResponseOutput) HeaderFieldValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderFieldResponse) *string { return v.HeaderFieldValue }).(pulumi.StringPtrOutput)
}

type HeaderFieldResponseArrayOutput struct{ *pulumi.OutputState }

func (HeaderFieldResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HeaderFieldResponse)(nil)).Elem()
}

func (o HeaderFieldResponseArrayOutput) ToHeaderFieldResponseArrayOutput() HeaderFieldResponseArrayOutput {
	return o
}

func (o HeaderFieldResponseArrayOutput) ToHeaderFieldResponseArrayOutputWithContext(ctx context.Context) HeaderFieldResponseArrayOutput {
	return o
}

func (o HeaderFieldResponseArrayOutput) Index(i pulumi.IntInput) HeaderFieldResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HeaderFieldResponse {
		return vs[0].([]HeaderFieldResponse)[vs[1].(int)]
	}).(HeaderFieldResponseOutput)
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

type WebTestGeolocation struct {
	Location *string `pulumi:"location"`
}





type WebTestGeolocationInput interface {
	pulumi.Input

	ToWebTestGeolocationOutput() WebTestGeolocationOutput
	ToWebTestGeolocationOutputWithContext(context.Context) WebTestGeolocationOutput
}

type WebTestGeolocationArgs struct {
	Location pulumi.StringPtrInput `pulumi:"location"`
}

func (WebTestGeolocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocation)(nil)).Elem()
}

func (i WebTestGeolocationArgs) ToWebTestGeolocationOutput() WebTestGeolocationOutput {
	return i.ToWebTestGeolocationOutputWithContext(context.Background())
}

func (i WebTestGeolocationArgs) ToWebTestGeolocationOutputWithContext(ctx context.Context) WebTestGeolocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationOutput)
}





type WebTestGeolocationArrayInput interface {
	pulumi.Input

	ToWebTestGeolocationArrayOutput() WebTestGeolocationArrayOutput
	ToWebTestGeolocationArrayOutputWithContext(context.Context) WebTestGeolocationArrayOutput
}

type WebTestGeolocationArray []WebTestGeolocationInput

func (WebTestGeolocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocation)(nil)).Elem()
}

func (i WebTestGeolocationArray) ToWebTestGeolocationArrayOutput() WebTestGeolocationArrayOutput {
	return i.ToWebTestGeolocationArrayOutputWithContext(context.Background())
}

func (i WebTestGeolocationArray) ToWebTestGeolocationArrayOutputWithContext(ctx context.Context) WebTestGeolocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationArrayOutput)
}

type WebTestGeolocationOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocation)(nil)).Elem()
}

func (o WebTestGeolocationOutput) ToWebTestGeolocationOutput() WebTestGeolocationOutput {
	return o
}

func (o WebTestGeolocationOutput) ToWebTestGeolocationOutputWithContext(ctx context.Context) WebTestGeolocationOutput {
	return o
}

func (o WebTestGeolocationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestGeolocation) *string { return v.Location }).(pulumi.StringPtrOutput)
}

type WebTestGeolocationArrayOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocation)(nil)).Elem()
}

func (o WebTestGeolocationArrayOutput) ToWebTestGeolocationArrayOutput() WebTestGeolocationArrayOutput {
	return o
}

func (o WebTestGeolocationArrayOutput) ToWebTestGeolocationArrayOutputWithContext(ctx context.Context) WebTestGeolocationArrayOutput {
	return o
}

func (o WebTestGeolocationArrayOutput) Index(i pulumi.IntInput) WebTestGeolocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebTestGeolocation {
		return vs[0].([]WebTestGeolocation)[vs[1].(int)]
	}).(WebTestGeolocationOutput)
}

type WebTestGeolocationResponse struct {
	Location *string `pulumi:"location"`
}

type WebTestGeolocationResponseOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocationResponse)(nil)).Elem()
}

func (o WebTestGeolocationResponseOutput) ToWebTestGeolocationResponseOutput() WebTestGeolocationResponseOutput {
	return o
}

func (o WebTestGeolocationResponseOutput) ToWebTestGeolocationResponseOutputWithContext(ctx context.Context) WebTestGeolocationResponseOutput {
	return o
}

func (o WebTestGeolocationResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestGeolocationResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

type WebTestGeolocationResponseArrayOutput struct{ *pulumi.OutputState }

func (WebTestGeolocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocationResponse)(nil)).Elem()
}

func (o WebTestGeolocationResponseArrayOutput) ToWebTestGeolocationResponseArrayOutput() WebTestGeolocationResponseArrayOutput {
	return o
}

func (o WebTestGeolocationResponseArrayOutput) ToWebTestGeolocationResponseArrayOutputWithContext(ctx context.Context) WebTestGeolocationResponseArrayOutput {
	return o
}

func (o WebTestGeolocationResponseArrayOutput) Index(i pulumi.IntInput) WebTestGeolocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebTestGeolocationResponse {
		return vs[0].([]WebTestGeolocationResponse)[vs[1].(int)]
	}).(WebTestGeolocationResponseOutput)
}

type WebTestPropertiesConfiguration struct {
	WebTest *string `pulumi:"webTest"`
}





type WebTestPropertiesConfigurationInput interface {
	pulumi.Input

	ToWebTestPropertiesConfigurationOutput() WebTestPropertiesConfigurationOutput
	ToWebTestPropertiesConfigurationOutputWithContext(context.Context) WebTestPropertiesConfigurationOutput
}

type WebTestPropertiesConfigurationArgs struct {
	WebTest pulumi.StringPtrInput `pulumi:"webTest"`
}

func (WebTestPropertiesConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesConfiguration)(nil)).Elem()
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationOutput() WebTestPropertiesConfigurationOutput {
	return i.ToWebTestPropertiesConfigurationOutputWithContext(context.Background())
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesConfigurationOutput)
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return i.ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Background())
}

func (i WebTestPropertiesConfigurationArgs) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesConfigurationOutput).ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx)
}









type WebTestPropertiesConfigurationPtrInput interface {
	pulumi.Input

	ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput
	ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Context) WebTestPropertiesConfigurationPtrOutput
}

type webTestPropertiesConfigurationPtrType WebTestPropertiesConfigurationArgs

func WebTestPropertiesConfigurationPtr(v *WebTestPropertiesConfigurationArgs) WebTestPropertiesConfigurationPtrInput {
	return (*webTestPropertiesConfigurationPtrType)(v)
}

func (*webTestPropertiesConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesConfiguration)(nil)).Elem()
}

func (i *webTestPropertiesConfigurationPtrType) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return i.ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Background())
}

func (i *webTestPropertiesConfigurationPtrType) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesConfigurationPtrOutput)
}

type WebTestPropertiesConfigurationOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesConfigurationOutput) ToWebTestPropertiesConfigurationOutput() WebTestPropertiesConfigurationOutput {
	return o
}

func (o WebTestPropertiesConfigurationOutput) ToWebTestPropertiesConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationOutput {
	return o
}

func (o WebTestPropertiesConfigurationOutput) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return o.ToWebTestPropertiesConfigurationPtrOutputWithContext(context.Background())
}

func (o WebTestPropertiesConfigurationOutput) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTestPropertiesConfiguration) *WebTestPropertiesConfiguration {
		return &v
	}).(WebTestPropertiesConfigurationPtrOutput)
}

func (o WebTestPropertiesConfigurationOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesConfiguration) *string { return v.WebTest }).(pulumi.StringPtrOutput)
}

type WebTestPropertiesConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesConfigurationPtrOutput) ToWebTestPropertiesConfigurationPtrOutput() WebTestPropertiesConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesConfigurationPtrOutput) ToWebTestPropertiesConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesConfigurationPtrOutput) Elem() WebTestPropertiesConfigurationOutput {
	return o.ApplyT(func(v *WebTestPropertiesConfiguration) WebTestPropertiesConfiguration {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesConfiguration
		return ret
	}).(WebTestPropertiesConfigurationOutput)
}

func (o WebTestPropertiesConfigurationPtrOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.WebTest
	}).(pulumi.StringPtrOutput)
}

type WebTestPropertiesContentValidation struct {
	ContentMatch    *string `pulumi:"contentMatch"`
	IgnoreCase      *bool   `pulumi:"ignoreCase"`
	PassIfTextFound *bool   `pulumi:"passIfTextFound"`
}





type WebTestPropertiesContentValidationInput interface {
	pulumi.Input

	ToWebTestPropertiesContentValidationOutput() WebTestPropertiesContentValidationOutput
	ToWebTestPropertiesContentValidationOutputWithContext(context.Context) WebTestPropertiesContentValidationOutput
}

type WebTestPropertiesContentValidationArgs struct {
	ContentMatch    pulumi.StringPtrInput `pulumi:"contentMatch"`
	IgnoreCase      pulumi.BoolPtrInput   `pulumi:"ignoreCase"`
	PassIfTextFound pulumi.BoolPtrInput   `pulumi:"passIfTextFound"`
}

func (WebTestPropertiesContentValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesContentValidation)(nil)).Elem()
}

func (i WebTestPropertiesContentValidationArgs) ToWebTestPropertiesContentValidationOutput() WebTestPropertiesContentValidationOutput {
	return i.ToWebTestPropertiesContentValidationOutputWithContext(context.Background())
}

func (i WebTestPropertiesContentValidationArgs) ToWebTestPropertiesContentValidationOutputWithContext(ctx context.Context) WebTestPropertiesContentValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesContentValidationOutput)
}

func (i WebTestPropertiesContentValidationArgs) ToWebTestPropertiesContentValidationPtrOutput() WebTestPropertiesContentValidationPtrOutput {
	return i.ToWebTestPropertiesContentValidationPtrOutputWithContext(context.Background())
}

func (i WebTestPropertiesContentValidationArgs) ToWebTestPropertiesContentValidationPtrOutputWithContext(ctx context.Context) WebTestPropertiesContentValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesContentValidationOutput).ToWebTestPropertiesContentValidationPtrOutputWithContext(ctx)
}









type WebTestPropertiesContentValidationPtrInput interface {
	pulumi.Input

	ToWebTestPropertiesContentValidationPtrOutput() WebTestPropertiesContentValidationPtrOutput
	ToWebTestPropertiesContentValidationPtrOutputWithContext(context.Context) WebTestPropertiesContentValidationPtrOutput
}

type webTestPropertiesContentValidationPtrType WebTestPropertiesContentValidationArgs

func WebTestPropertiesContentValidationPtr(v *WebTestPropertiesContentValidationArgs) WebTestPropertiesContentValidationPtrInput {
	return (*webTestPropertiesContentValidationPtrType)(v)
}

func (*webTestPropertiesContentValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesContentValidation)(nil)).Elem()
}

func (i *webTestPropertiesContentValidationPtrType) ToWebTestPropertiesContentValidationPtrOutput() WebTestPropertiesContentValidationPtrOutput {
	return i.ToWebTestPropertiesContentValidationPtrOutputWithContext(context.Background())
}

func (i *webTestPropertiesContentValidationPtrType) ToWebTestPropertiesContentValidationPtrOutputWithContext(ctx context.Context) WebTestPropertiesContentValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesContentValidationPtrOutput)
}

type WebTestPropertiesContentValidationOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesContentValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesContentValidation)(nil)).Elem()
}

func (o WebTestPropertiesContentValidationOutput) ToWebTestPropertiesContentValidationOutput() WebTestPropertiesContentValidationOutput {
	return o
}

func (o WebTestPropertiesContentValidationOutput) ToWebTestPropertiesContentValidationOutputWithContext(ctx context.Context) WebTestPropertiesContentValidationOutput {
	return o
}

func (o WebTestPropertiesContentValidationOutput) ToWebTestPropertiesContentValidationPtrOutput() WebTestPropertiesContentValidationPtrOutput {
	return o.ToWebTestPropertiesContentValidationPtrOutputWithContext(context.Background())
}

func (o WebTestPropertiesContentValidationOutput) ToWebTestPropertiesContentValidationPtrOutputWithContext(ctx context.Context) WebTestPropertiesContentValidationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTestPropertiesContentValidation) *WebTestPropertiesContentValidation {
		return &v
	}).(WebTestPropertiesContentValidationPtrOutput)
}

func (o WebTestPropertiesContentValidationOutput) ContentMatch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesContentValidation) *string { return v.ContentMatch }).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesContentValidationOutput) IgnoreCase() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesContentValidation) *bool { return v.IgnoreCase }).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesContentValidationOutput) PassIfTextFound() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesContentValidation) *bool { return v.PassIfTextFound }).(pulumi.BoolPtrOutput)
}

type WebTestPropertiesContentValidationPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesContentValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesContentValidation)(nil)).Elem()
}

func (o WebTestPropertiesContentValidationPtrOutput) ToWebTestPropertiesContentValidationPtrOutput() WebTestPropertiesContentValidationPtrOutput {
	return o
}

func (o WebTestPropertiesContentValidationPtrOutput) ToWebTestPropertiesContentValidationPtrOutputWithContext(ctx context.Context) WebTestPropertiesContentValidationPtrOutput {
	return o
}

func (o WebTestPropertiesContentValidationPtrOutput) Elem() WebTestPropertiesContentValidationOutput {
	return o.ApplyT(func(v *WebTestPropertiesContentValidation) WebTestPropertiesContentValidation {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesContentValidation
		return ret
	}).(WebTestPropertiesContentValidationOutput)
}

func (o WebTestPropertiesContentValidationPtrOutput) ContentMatch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesContentValidation) *string {
		if v == nil {
			return nil
		}
		return v.ContentMatch
	}).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesContentValidationPtrOutput) IgnoreCase() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesContentValidation) *bool {
		if v == nil {
			return nil
		}
		return v.IgnoreCase
	}).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesContentValidationPtrOutput) PassIfTextFound() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesContentValidation) *bool {
		if v == nil {
			return nil
		}
		return v.PassIfTextFound
	}).(pulumi.BoolPtrOutput)
}

type WebTestPropertiesRequest struct {
	FollowRedirects        *bool         `pulumi:"followRedirects"`
	Headers                []HeaderField `pulumi:"headers"`
	HttpVerb               *string       `pulumi:"httpVerb"`
	ParseDependentRequests *bool         `pulumi:"parseDependentRequests"`
	RequestBody            *string       `pulumi:"requestBody"`
	RequestUrl             *string       `pulumi:"requestUrl"`
}





type WebTestPropertiesRequestInput interface {
	pulumi.Input

	ToWebTestPropertiesRequestOutput() WebTestPropertiesRequestOutput
	ToWebTestPropertiesRequestOutputWithContext(context.Context) WebTestPropertiesRequestOutput
}

type WebTestPropertiesRequestArgs struct {
	FollowRedirects        pulumi.BoolPtrInput   `pulumi:"followRedirects"`
	Headers                HeaderFieldArrayInput `pulumi:"headers"`
	HttpVerb               pulumi.StringPtrInput `pulumi:"httpVerb"`
	ParseDependentRequests pulumi.BoolPtrInput   `pulumi:"parseDependentRequests"`
	RequestBody            pulumi.StringPtrInput `pulumi:"requestBody"`
	RequestUrl             pulumi.StringPtrInput `pulumi:"requestUrl"`
}

func (WebTestPropertiesRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesRequest)(nil)).Elem()
}

func (i WebTestPropertiesRequestArgs) ToWebTestPropertiesRequestOutput() WebTestPropertiesRequestOutput {
	return i.ToWebTestPropertiesRequestOutputWithContext(context.Background())
}

func (i WebTestPropertiesRequestArgs) ToWebTestPropertiesRequestOutputWithContext(ctx context.Context) WebTestPropertiesRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesRequestOutput)
}

func (i WebTestPropertiesRequestArgs) ToWebTestPropertiesRequestPtrOutput() WebTestPropertiesRequestPtrOutput {
	return i.ToWebTestPropertiesRequestPtrOutputWithContext(context.Background())
}

func (i WebTestPropertiesRequestArgs) ToWebTestPropertiesRequestPtrOutputWithContext(ctx context.Context) WebTestPropertiesRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesRequestOutput).ToWebTestPropertiesRequestPtrOutputWithContext(ctx)
}









type WebTestPropertiesRequestPtrInput interface {
	pulumi.Input

	ToWebTestPropertiesRequestPtrOutput() WebTestPropertiesRequestPtrOutput
	ToWebTestPropertiesRequestPtrOutputWithContext(context.Context) WebTestPropertiesRequestPtrOutput
}

type webTestPropertiesRequestPtrType WebTestPropertiesRequestArgs

func WebTestPropertiesRequestPtr(v *WebTestPropertiesRequestArgs) WebTestPropertiesRequestPtrInput {
	return (*webTestPropertiesRequestPtrType)(v)
}

func (*webTestPropertiesRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesRequest)(nil)).Elem()
}

func (i *webTestPropertiesRequestPtrType) ToWebTestPropertiesRequestPtrOutput() WebTestPropertiesRequestPtrOutput {
	return i.ToWebTestPropertiesRequestPtrOutputWithContext(context.Background())
}

func (i *webTestPropertiesRequestPtrType) ToWebTestPropertiesRequestPtrOutputWithContext(ctx context.Context) WebTestPropertiesRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesRequestPtrOutput)
}

type WebTestPropertiesRequestOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesRequest)(nil)).Elem()
}

func (o WebTestPropertiesRequestOutput) ToWebTestPropertiesRequestOutput() WebTestPropertiesRequestOutput {
	return o
}

func (o WebTestPropertiesRequestOutput) ToWebTestPropertiesRequestOutputWithContext(ctx context.Context) WebTestPropertiesRequestOutput {
	return o
}

func (o WebTestPropertiesRequestOutput) ToWebTestPropertiesRequestPtrOutput() WebTestPropertiesRequestPtrOutput {
	return o.ToWebTestPropertiesRequestPtrOutputWithContext(context.Background())
}

func (o WebTestPropertiesRequestOutput) ToWebTestPropertiesRequestPtrOutputWithContext(ctx context.Context) WebTestPropertiesRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTestPropertiesRequest) *WebTestPropertiesRequest {
		return &v
	}).(WebTestPropertiesRequestPtrOutput)
}

func (o WebTestPropertiesRequestOutput) FollowRedirects() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesRequest) *bool { return v.FollowRedirects }).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesRequestOutput) Headers() HeaderFieldArrayOutput {
	return o.ApplyT(func(v WebTestPropertiesRequest) []HeaderField { return v.Headers }).(HeaderFieldArrayOutput)
}

func (o WebTestPropertiesRequestOutput) HttpVerb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesRequest) *string { return v.HttpVerb }).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesRequestOutput) ParseDependentRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesRequest) *bool { return v.ParseDependentRequests }).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesRequestOutput) RequestBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesRequest) *string { return v.RequestBody }).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesRequestOutput) RequestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesRequest) *string { return v.RequestUrl }).(pulumi.StringPtrOutput)
}

type WebTestPropertiesRequestPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesRequest)(nil)).Elem()
}

func (o WebTestPropertiesRequestPtrOutput) ToWebTestPropertiesRequestPtrOutput() WebTestPropertiesRequestPtrOutput {
	return o
}

func (o WebTestPropertiesRequestPtrOutput) ToWebTestPropertiesRequestPtrOutputWithContext(ctx context.Context) WebTestPropertiesRequestPtrOutput {
	return o
}

func (o WebTestPropertiesRequestPtrOutput) Elem() WebTestPropertiesRequestOutput {
	return o.ApplyT(func(v *WebTestPropertiesRequest) WebTestPropertiesRequest {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesRequest
		return ret
	}).(WebTestPropertiesRequestOutput)
}

func (o WebTestPropertiesRequestPtrOutput) FollowRedirects() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesRequest) *bool {
		if v == nil {
			return nil
		}
		return v.FollowRedirects
	}).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesRequestPtrOutput) Headers() HeaderFieldArrayOutput {
	return o.ApplyT(func(v *WebTestPropertiesRequest) []HeaderField {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(HeaderFieldArrayOutput)
}

func (o WebTestPropertiesRequestPtrOutput) HttpVerb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesRequest) *string {
		if v == nil {
			return nil
		}
		return v.HttpVerb
	}).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesRequestPtrOutput) ParseDependentRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesRequest) *bool {
		if v == nil {
			return nil
		}
		return v.ParseDependentRequests
	}).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesRequestPtrOutput) RequestBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesRequest) *string {
		if v == nil {
			return nil
		}
		return v.RequestBody
	}).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesRequestPtrOutput) RequestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesRequest) *string {
		if v == nil {
			return nil
		}
		return v.RequestUrl
	}).(pulumi.StringPtrOutput)
}

type WebTestPropertiesResponseConfiguration struct {
	WebTest *string `pulumi:"webTest"`
}

type WebTestPropertiesResponseConfigurationOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationOutput() WebTestPropertiesResponseConfigurationOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseConfiguration) *string { return v.WebTest }).(pulumi.StringPtrOutput)
}

type WebTestPropertiesResponseConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return o
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) Elem() WebTestPropertiesResponseConfigurationOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseConfiguration) WebTestPropertiesResponseConfiguration {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesResponseConfiguration
		return ret
	}).(WebTestPropertiesResponseConfigurationOutput)
}

func (o WebTestPropertiesResponseConfigurationPtrOutput) WebTest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.WebTest
	}).(pulumi.StringPtrOutput)
}

type WebTestPropertiesResponseContentValidation struct {
	ContentMatch    *string `pulumi:"contentMatch"`
	IgnoreCase      *bool   `pulumi:"ignoreCase"`
	PassIfTextFound *bool   `pulumi:"passIfTextFound"`
}

type WebTestPropertiesResponseContentValidationOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseContentValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesResponseContentValidation)(nil)).Elem()
}

func (o WebTestPropertiesResponseContentValidationOutput) ToWebTestPropertiesResponseContentValidationOutput() WebTestPropertiesResponseContentValidationOutput {
	return o
}

func (o WebTestPropertiesResponseContentValidationOutput) ToWebTestPropertiesResponseContentValidationOutputWithContext(ctx context.Context) WebTestPropertiesResponseContentValidationOutput {
	return o
}

func (o WebTestPropertiesResponseContentValidationOutput) ContentMatch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseContentValidation) *string { return v.ContentMatch }).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesResponseContentValidationOutput) IgnoreCase() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseContentValidation) *bool { return v.IgnoreCase }).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesResponseContentValidationOutput) PassIfTextFound() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseContentValidation) *bool { return v.PassIfTextFound }).(pulumi.BoolPtrOutput)
}

type WebTestPropertiesResponseContentValidationPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseContentValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesResponseContentValidation)(nil)).Elem()
}

func (o WebTestPropertiesResponseContentValidationPtrOutput) ToWebTestPropertiesResponseContentValidationPtrOutput() WebTestPropertiesResponseContentValidationPtrOutput {
	return o
}

func (o WebTestPropertiesResponseContentValidationPtrOutput) ToWebTestPropertiesResponseContentValidationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseContentValidationPtrOutput {
	return o
}

func (o WebTestPropertiesResponseContentValidationPtrOutput) Elem() WebTestPropertiesResponseContentValidationOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseContentValidation) WebTestPropertiesResponseContentValidation {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesResponseContentValidation
		return ret
	}).(WebTestPropertiesResponseContentValidationOutput)
}

func (o WebTestPropertiesResponseContentValidationPtrOutput) ContentMatch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseContentValidation) *string {
		if v == nil {
			return nil
		}
		return v.ContentMatch
	}).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesResponseContentValidationPtrOutput) IgnoreCase() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseContentValidation) *bool {
		if v == nil {
			return nil
		}
		return v.IgnoreCase
	}).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesResponseContentValidationPtrOutput) PassIfTextFound() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseContentValidation) *bool {
		if v == nil {
			return nil
		}
		return v.PassIfTextFound
	}).(pulumi.BoolPtrOutput)
}

type WebTestPropertiesResponseRequest struct {
	FollowRedirects        *bool                 `pulumi:"followRedirects"`
	Headers                []HeaderFieldResponse `pulumi:"headers"`
	HttpVerb               *string               `pulumi:"httpVerb"`
	ParseDependentRequests *bool                 `pulumi:"parseDependentRequests"`
	RequestBody            *string               `pulumi:"requestBody"`
	RequestUrl             *string               `pulumi:"requestUrl"`
}

type WebTestPropertiesResponseRequestOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesResponseRequest)(nil)).Elem()
}

func (o WebTestPropertiesResponseRequestOutput) ToWebTestPropertiesResponseRequestOutput() WebTestPropertiesResponseRequestOutput {
	return o
}

func (o WebTestPropertiesResponseRequestOutput) ToWebTestPropertiesResponseRequestOutputWithContext(ctx context.Context) WebTestPropertiesResponseRequestOutput {
	return o
}

func (o WebTestPropertiesResponseRequestOutput) FollowRedirects() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseRequest) *bool { return v.FollowRedirects }).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesResponseRequestOutput) Headers() HeaderFieldResponseArrayOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseRequest) []HeaderFieldResponse { return v.Headers }).(HeaderFieldResponseArrayOutput)
}

func (o WebTestPropertiesResponseRequestOutput) HttpVerb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseRequest) *string { return v.HttpVerb }).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesResponseRequestOutput) ParseDependentRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseRequest) *bool { return v.ParseDependentRequests }).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesResponseRequestOutput) RequestBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseRequest) *string { return v.RequestBody }).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesResponseRequestOutput) RequestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseRequest) *string { return v.RequestUrl }).(pulumi.StringPtrOutput)
}

type WebTestPropertiesResponseRequestPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesResponseRequest)(nil)).Elem()
}

func (o WebTestPropertiesResponseRequestPtrOutput) ToWebTestPropertiesResponseRequestPtrOutput() WebTestPropertiesResponseRequestPtrOutput {
	return o
}

func (o WebTestPropertiesResponseRequestPtrOutput) ToWebTestPropertiesResponseRequestPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseRequestPtrOutput {
	return o
}

func (o WebTestPropertiesResponseRequestPtrOutput) Elem() WebTestPropertiesResponseRequestOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseRequest) WebTestPropertiesResponseRequest {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesResponseRequest
		return ret
	}).(WebTestPropertiesResponseRequestOutput)
}

func (o WebTestPropertiesResponseRequestPtrOutput) FollowRedirects() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseRequest) *bool {
		if v == nil {
			return nil
		}
		return v.FollowRedirects
	}).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesResponseRequestPtrOutput) Headers() HeaderFieldResponseArrayOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseRequest) []HeaderFieldResponse {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(HeaderFieldResponseArrayOutput)
}

func (o WebTestPropertiesResponseRequestPtrOutput) HttpVerb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseRequest) *string {
		if v == nil {
			return nil
		}
		return v.HttpVerb
	}).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesResponseRequestPtrOutput) ParseDependentRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseRequest) *bool {
		if v == nil {
			return nil
		}
		return v.ParseDependentRequests
	}).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesResponseRequestPtrOutput) RequestBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseRequest) *string {
		if v == nil {
			return nil
		}
		return v.RequestBody
	}).(pulumi.StringPtrOutput)
}

func (o WebTestPropertiesResponseRequestPtrOutput) RequestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseRequest) *string {
		if v == nil {
			return nil
		}
		return v.RequestUrl
	}).(pulumi.StringPtrOutput)
}

type WebTestPropertiesResponseValidationRules struct {
	ContentValidation             *WebTestPropertiesResponseContentValidation `pulumi:"contentValidation"`
	ExpectedHttpStatusCode        *int                                        `pulumi:"expectedHttpStatusCode"`
	IgnoreHttpsStatusCode         *bool                                       `pulumi:"ignoreHttpsStatusCode"`
	SSLCertRemainingLifetimeCheck *int                                        `pulumi:"sSLCertRemainingLifetimeCheck"`
	SSLCheck                      *bool                                       `pulumi:"sSLCheck"`
}

type WebTestPropertiesResponseValidationRulesOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseValidationRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesResponseValidationRules)(nil)).Elem()
}

func (o WebTestPropertiesResponseValidationRulesOutput) ToWebTestPropertiesResponseValidationRulesOutput() WebTestPropertiesResponseValidationRulesOutput {
	return o
}

func (o WebTestPropertiesResponseValidationRulesOutput) ToWebTestPropertiesResponseValidationRulesOutputWithContext(ctx context.Context) WebTestPropertiesResponseValidationRulesOutput {
	return o
}

func (o WebTestPropertiesResponseValidationRulesOutput) ContentValidation() WebTestPropertiesResponseContentValidationPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseValidationRules) *WebTestPropertiesResponseContentValidation {
		return v.ContentValidation
	}).(WebTestPropertiesResponseContentValidationPtrOutput)
}

func (o WebTestPropertiesResponseValidationRulesOutput) ExpectedHttpStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseValidationRules) *int { return v.ExpectedHttpStatusCode }).(pulumi.IntPtrOutput)
}

func (o WebTestPropertiesResponseValidationRulesOutput) IgnoreHttpsStatusCode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseValidationRules) *bool { return v.IgnoreHttpsStatusCode }).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesResponseValidationRulesOutput) SSLCertRemainingLifetimeCheck() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseValidationRules) *int { return v.SSLCertRemainingLifetimeCheck }).(pulumi.IntPtrOutput)
}

func (o WebTestPropertiesResponseValidationRulesOutput) SSLCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesResponseValidationRules) *bool { return v.SSLCheck }).(pulumi.BoolPtrOutput)
}

type WebTestPropertiesResponseValidationRulesPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesResponseValidationRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesResponseValidationRules)(nil)).Elem()
}

func (o WebTestPropertiesResponseValidationRulesPtrOutput) ToWebTestPropertiesResponseValidationRulesPtrOutput() WebTestPropertiesResponseValidationRulesPtrOutput {
	return o
}

func (o WebTestPropertiesResponseValidationRulesPtrOutput) ToWebTestPropertiesResponseValidationRulesPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseValidationRulesPtrOutput {
	return o
}

func (o WebTestPropertiesResponseValidationRulesPtrOutput) Elem() WebTestPropertiesResponseValidationRulesOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseValidationRules) WebTestPropertiesResponseValidationRules {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesResponseValidationRules
		return ret
	}).(WebTestPropertiesResponseValidationRulesOutput)
}

func (o WebTestPropertiesResponseValidationRulesPtrOutput) ContentValidation() WebTestPropertiesResponseContentValidationPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseValidationRules) *WebTestPropertiesResponseContentValidation {
		if v == nil {
			return nil
		}
		return v.ContentValidation
	}).(WebTestPropertiesResponseContentValidationPtrOutput)
}

func (o WebTestPropertiesResponseValidationRulesPtrOutput) ExpectedHttpStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseValidationRules) *int {
		if v == nil {
			return nil
		}
		return v.ExpectedHttpStatusCode
	}).(pulumi.IntPtrOutput)
}

func (o WebTestPropertiesResponseValidationRulesPtrOutput) IgnoreHttpsStatusCode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseValidationRules) *bool {
		if v == nil {
			return nil
		}
		return v.IgnoreHttpsStatusCode
	}).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesResponseValidationRulesPtrOutput) SSLCertRemainingLifetimeCheck() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseValidationRules) *int {
		if v == nil {
			return nil
		}
		return v.SSLCertRemainingLifetimeCheck
	}).(pulumi.IntPtrOutput)
}

func (o WebTestPropertiesResponseValidationRulesPtrOutput) SSLCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesResponseValidationRules) *bool {
		if v == nil {
			return nil
		}
		return v.SSLCheck
	}).(pulumi.BoolPtrOutput)
}

type WebTestPropertiesValidationRules struct {
	ContentValidation             *WebTestPropertiesContentValidation `pulumi:"contentValidation"`
	ExpectedHttpStatusCode        *int                                `pulumi:"expectedHttpStatusCode"`
	IgnoreHttpsStatusCode         *bool                               `pulumi:"ignoreHttpsStatusCode"`
	SSLCertRemainingLifetimeCheck *int                                `pulumi:"sSLCertRemainingLifetimeCheck"`
	SSLCheck                      *bool                               `pulumi:"sSLCheck"`
}





type WebTestPropertiesValidationRulesInput interface {
	pulumi.Input

	ToWebTestPropertiesValidationRulesOutput() WebTestPropertiesValidationRulesOutput
	ToWebTestPropertiesValidationRulesOutputWithContext(context.Context) WebTestPropertiesValidationRulesOutput
}

type WebTestPropertiesValidationRulesArgs struct {
	ContentValidation             WebTestPropertiesContentValidationPtrInput `pulumi:"contentValidation"`
	ExpectedHttpStatusCode        pulumi.IntPtrInput                         `pulumi:"expectedHttpStatusCode"`
	IgnoreHttpsStatusCode         pulumi.BoolPtrInput                        `pulumi:"ignoreHttpsStatusCode"`
	SSLCertRemainingLifetimeCheck pulumi.IntPtrInput                         `pulumi:"sSLCertRemainingLifetimeCheck"`
	SSLCheck                      pulumi.BoolPtrInput                        `pulumi:"sSLCheck"`
}

func (WebTestPropertiesValidationRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesValidationRules)(nil)).Elem()
}

func (i WebTestPropertiesValidationRulesArgs) ToWebTestPropertiesValidationRulesOutput() WebTestPropertiesValidationRulesOutput {
	return i.ToWebTestPropertiesValidationRulesOutputWithContext(context.Background())
}

func (i WebTestPropertiesValidationRulesArgs) ToWebTestPropertiesValidationRulesOutputWithContext(ctx context.Context) WebTestPropertiesValidationRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesValidationRulesOutput)
}

func (i WebTestPropertiesValidationRulesArgs) ToWebTestPropertiesValidationRulesPtrOutput() WebTestPropertiesValidationRulesPtrOutput {
	return i.ToWebTestPropertiesValidationRulesPtrOutputWithContext(context.Background())
}

func (i WebTestPropertiesValidationRulesArgs) ToWebTestPropertiesValidationRulesPtrOutputWithContext(ctx context.Context) WebTestPropertiesValidationRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesValidationRulesOutput).ToWebTestPropertiesValidationRulesPtrOutputWithContext(ctx)
}









type WebTestPropertiesValidationRulesPtrInput interface {
	pulumi.Input

	ToWebTestPropertiesValidationRulesPtrOutput() WebTestPropertiesValidationRulesPtrOutput
	ToWebTestPropertiesValidationRulesPtrOutputWithContext(context.Context) WebTestPropertiesValidationRulesPtrOutput
}

type webTestPropertiesValidationRulesPtrType WebTestPropertiesValidationRulesArgs

func WebTestPropertiesValidationRulesPtr(v *WebTestPropertiesValidationRulesArgs) WebTestPropertiesValidationRulesPtrInput {
	return (*webTestPropertiesValidationRulesPtrType)(v)
}

func (*webTestPropertiesValidationRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesValidationRules)(nil)).Elem()
}

func (i *webTestPropertiesValidationRulesPtrType) ToWebTestPropertiesValidationRulesPtrOutput() WebTestPropertiesValidationRulesPtrOutput {
	return i.ToWebTestPropertiesValidationRulesPtrOutputWithContext(context.Background())
}

func (i *webTestPropertiesValidationRulesPtrType) ToWebTestPropertiesValidationRulesPtrOutputWithContext(ctx context.Context) WebTestPropertiesValidationRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesValidationRulesPtrOutput)
}

type WebTestPropertiesValidationRulesOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesValidationRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesValidationRules)(nil)).Elem()
}

func (o WebTestPropertiesValidationRulesOutput) ToWebTestPropertiesValidationRulesOutput() WebTestPropertiesValidationRulesOutput {
	return o
}

func (o WebTestPropertiesValidationRulesOutput) ToWebTestPropertiesValidationRulesOutputWithContext(ctx context.Context) WebTestPropertiesValidationRulesOutput {
	return o
}

func (o WebTestPropertiesValidationRulesOutput) ToWebTestPropertiesValidationRulesPtrOutput() WebTestPropertiesValidationRulesPtrOutput {
	return o.ToWebTestPropertiesValidationRulesPtrOutputWithContext(context.Background())
}

func (o WebTestPropertiesValidationRulesOutput) ToWebTestPropertiesValidationRulesPtrOutputWithContext(ctx context.Context) WebTestPropertiesValidationRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTestPropertiesValidationRules) *WebTestPropertiesValidationRules {
		return &v
	}).(WebTestPropertiesValidationRulesPtrOutput)
}

func (o WebTestPropertiesValidationRulesOutput) ContentValidation() WebTestPropertiesContentValidationPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesValidationRules) *WebTestPropertiesContentValidation {
		return v.ContentValidation
	}).(WebTestPropertiesContentValidationPtrOutput)
}

func (o WebTestPropertiesValidationRulesOutput) ExpectedHttpStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesValidationRules) *int { return v.ExpectedHttpStatusCode }).(pulumi.IntPtrOutput)
}

func (o WebTestPropertiesValidationRulesOutput) IgnoreHttpsStatusCode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesValidationRules) *bool { return v.IgnoreHttpsStatusCode }).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesValidationRulesOutput) SSLCertRemainingLifetimeCheck() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesValidationRules) *int { return v.SSLCertRemainingLifetimeCheck }).(pulumi.IntPtrOutput)
}

func (o WebTestPropertiesValidationRulesOutput) SSLCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebTestPropertiesValidationRules) *bool { return v.SSLCheck }).(pulumi.BoolPtrOutput)
}

type WebTestPropertiesValidationRulesPtrOutput struct{ *pulumi.OutputState }

func (WebTestPropertiesValidationRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesValidationRules)(nil)).Elem()
}

func (o WebTestPropertiesValidationRulesPtrOutput) ToWebTestPropertiesValidationRulesPtrOutput() WebTestPropertiesValidationRulesPtrOutput {
	return o
}

func (o WebTestPropertiesValidationRulesPtrOutput) ToWebTestPropertiesValidationRulesPtrOutputWithContext(ctx context.Context) WebTestPropertiesValidationRulesPtrOutput {
	return o
}

func (o WebTestPropertiesValidationRulesPtrOutput) Elem() WebTestPropertiesValidationRulesOutput {
	return o.ApplyT(func(v *WebTestPropertiesValidationRules) WebTestPropertiesValidationRules {
		if v != nil {
			return *v
		}
		var ret WebTestPropertiesValidationRules
		return ret
	}).(WebTestPropertiesValidationRulesOutput)
}

func (o WebTestPropertiesValidationRulesPtrOutput) ContentValidation() WebTestPropertiesContentValidationPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesValidationRules) *WebTestPropertiesContentValidation {
		if v == nil {
			return nil
		}
		return v.ContentValidation
	}).(WebTestPropertiesContentValidationPtrOutput)
}

func (o WebTestPropertiesValidationRulesPtrOutput) ExpectedHttpStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesValidationRules) *int {
		if v == nil {
			return nil
		}
		return v.ExpectedHttpStatusCode
	}).(pulumi.IntPtrOutput)
}

func (o WebTestPropertiesValidationRulesPtrOutput) IgnoreHttpsStatusCode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesValidationRules) *bool {
		if v == nil {
			return nil
		}
		return v.IgnoreHttpsStatusCode
	}).(pulumi.BoolPtrOutput)
}

func (o WebTestPropertiesValidationRulesPtrOutput) SSLCertRemainingLifetimeCheck() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesValidationRules) *int {
		if v == nil {
			return nil
		}
		return v.SSLCertRemainingLifetimeCheck
	}).(pulumi.IntPtrOutput)
}

func (o WebTestPropertiesValidationRulesPtrOutput) SSLCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTestPropertiesValidationRules) *bool {
		if v == nil {
			return nil
		}
		return v.SSLCheck
	}).(pulumi.BoolPtrOutput)
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
	pulumi.RegisterOutputType(HeaderFieldOutput{})
	pulumi.RegisterOutputType(HeaderFieldArrayOutput{})
	pulumi.RegisterOutputType(HeaderFieldResponseOutput{})
	pulumi.RegisterOutputType(HeaderFieldResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaOutput{})
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationArrayOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationResponseOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationResponseArrayOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesConfigurationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesContentValidationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesContentValidationPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesRequestOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesRequestPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseConfigurationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseContentValidationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseContentValidationPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseRequestOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseRequestPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseValidationRulesOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseValidationRulesPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesValidationRulesOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesValidationRulesPtrOutput{})
}
