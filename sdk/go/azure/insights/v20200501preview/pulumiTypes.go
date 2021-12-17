


package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Action struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}





type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

type ActionArgs struct {
	ActionGroupId     pulumi.StringPtrInput `pulumi:"actionGroupId"`
	WebHookProperties pulumi.StringMapInput `pulumi:"webHookProperties"`
}

func (ActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (i ActionArgs) ToActionOutput() ActionOutput {
	return i.ToActionOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput)
}





type ActionArrayInput interface {
	pulumi.Input

	ToActionArrayOutput() ActionArrayOutput
	ToActionArrayOutputWithContext(context.Context) ActionArrayOutput
}

type ActionArray []ActionInput

func (ActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Action)(nil)).Elem()
}

func (i ActionArray) ToActionArrayOutput() ActionArrayOutput {
	return i.ToActionArrayOutputWithContext(context.Background())
}

func (i ActionArray) ToActionArrayOutputWithContext(ctx context.Context) ActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionArrayOutput)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

func (o ActionOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Action) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o ActionOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v Action) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type ActionArrayOutput struct{ *pulumi.OutputState }

func (ActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Action)(nil)).Elem()
}

func (o ActionArrayOutput) ToActionArrayOutput() ActionArrayOutput {
	return o
}

func (o ActionArrayOutput) ToActionArrayOutputWithContext(ctx context.Context) ActionArrayOutput {
	return o
}

func (o ActionArrayOutput) Index(i pulumi.IntInput) ActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Action {
		return vs[0].([]Action)[vs[1].(int)]
	}).(ActionOutput)
}

type ActionResponse struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}

type ActionResponseOutput struct{ *pulumi.OutputState }

func (ActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionResponse)(nil)).Elem()
}

func (o ActionResponseOutput) ToActionResponseOutput() ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) ToActionResponseOutputWithContext(ctx context.Context) ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionResponse) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o ActionResponseOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ActionResponse) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type ActionResponseArrayOutput struct{ *pulumi.OutputState }

func (ActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionResponse)(nil)).Elem()
}

func (o ActionResponseArrayOutput) ToActionResponseArrayOutput() ActionResponseArrayOutput {
	return o
}

func (o ActionResponseArrayOutput) ToActionResponseArrayOutputWithContext(ctx context.Context) ActionResponseArrayOutput {
	return o
}

func (o ActionResponseArrayOutput) Index(i pulumi.IntInput) ActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActionResponse {
		return vs[0].([]ActionResponse)[vs[1].(int)]
	}).(ActionResponseOutput)
}

type Condition struct {
	Dimensions          []Dimension              `pulumi:"dimensions"`
	FailingPeriods      *ConditionFailingPeriods `pulumi:"failingPeriods"`
	MetricMeasureColumn *string                  `pulumi:"metricMeasureColumn"`
	Operator            string                   `pulumi:"operator"`
	Query               *string                  `pulumi:"query"`
	ResourceIdColumn    *string                  `pulumi:"resourceIdColumn"`
	Threshold           float64                  `pulumi:"threshold"`
	TimeAggregation     string                   `pulumi:"timeAggregation"`
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
	Operator            pulumi.StringInput              `pulumi:"operator"`
	Query               pulumi.StringPtrInput           `pulumi:"query"`
	ResourceIdColumn    pulumi.StringPtrInput           `pulumi:"resourceIdColumn"`
	Threshold           pulumi.Float64Input             `pulumi:"threshold"`
	TimeAggregation     pulumi.StringInput              `pulumi:"timeAggregation"`
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

func (o ConditionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v Condition) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ConditionOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) ResourceIdColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.ResourceIdColumn }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v Condition) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o ConditionOutput) TimeAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v Condition) string { return v.TimeAggregation }).(pulumi.StringOutput)
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
	Operator            string                           `pulumi:"operator"`
	Query               *string                          `pulumi:"query"`
	ResourceIdColumn    *string                          `pulumi:"resourceIdColumn"`
	Threshold           float64                          `pulumi:"threshold"`
	TimeAggregation     string                           `pulumi:"timeAggregation"`
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

func (o ConditionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ConditionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ConditionResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) ResourceIdColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.ResourceIdColumn }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v ConditionResponse) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o ConditionResponseOutput) TimeAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v ConditionResponse) string { return v.TimeAggregation }).(pulumi.StringOutput)
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

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionArrayOutput{})
	pulumi.RegisterOutputType(ActionResponseOutput{})
	pulumi.RegisterOutputType(ActionResponseArrayOutput{})
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
	pulumi.RegisterOutputType(ScheduledQueryRuleCriteriaResponseOutput{})
}
