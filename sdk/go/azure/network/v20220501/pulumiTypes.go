


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomRule struct {
	Action                     string                    `pulumi:"action"`
	EnabledState               *string                   `pulumi:"enabledState"`
	MatchConditions            []FrontDoorMatchCondition `pulumi:"matchConditions"`
	Name                       *string                   `pulumi:"name"`
	Priority                   int                       `pulumi:"priority"`
	RateLimitDurationInMinutes *int                      `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         *int                      `pulumi:"rateLimitThreshold"`
	RuleType                   string                    `pulumi:"ruleType"`
}





type CustomRuleInput interface {
	pulumi.Input

	ToCustomRuleOutput() CustomRuleOutput
	ToCustomRuleOutputWithContext(context.Context) CustomRuleOutput
}

type CustomRuleArgs struct {
	Action                     pulumi.StringInput                `pulumi:"action"`
	EnabledState               pulumi.StringPtrInput             `pulumi:"enabledState"`
	MatchConditions            FrontDoorMatchConditionArrayInput `pulumi:"matchConditions"`
	Name                       pulumi.StringPtrInput             `pulumi:"name"`
	Priority                   pulumi.IntInput                   `pulumi:"priority"`
	RateLimitDurationInMinutes pulumi.IntPtrInput                `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         pulumi.IntPtrInput                `pulumi:"rateLimitThreshold"`
	RuleType                   pulumi.StringInput                `pulumi:"ruleType"`
}

func (CustomRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRule)(nil)).Elem()
}

func (i CustomRuleArgs) ToCustomRuleOutput() CustomRuleOutput {
	return i.ToCustomRuleOutputWithContext(context.Background())
}

func (i CustomRuleArgs) ToCustomRuleOutputWithContext(ctx context.Context) CustomRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleOutput)
}





type CustomRuleArrayInput interface {
	pulumi.Input

	ToCustomRuleArrayOutput() CustomRuleArrayOutput
	ToCustomRuleArrayOutputWithContext(context.Context) CustomRuleArrayOutput
}

type CustomRuleArray []CustomRuleInput

func (CustomRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRule)(nil)).Elem()
}

func (i CustomRuleArray) ToCustomRuleArrayOutput() CustomRuleArrayOutput {
	return i.ToCustomRuleArrayOutputWithContext(context.Background())
}

func (i CustomRuleArray) ToCustomRuleArrayOutputWithContext(ctx context.Context) CustomRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleArrayOutput)
}

type CustomRuleOutput struct{ *pulumi.OutputState }

func (CustomRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRule)(nil)).Elem()
}

func (o CustomRuleOutput) ToCustomRuleOutput() CustomRuleOutput {
	return o
}

func (o CustomRuleOutput) ToCustomRuleOutputWithContext(ctx context.Context) CustomRuleOutput {
	return o
}

func (o CustomRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRule) string { return v.Action }).(pulumi.StringOutput)
}

func (o CustomRuleOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRule) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o CustomRuleOutput) MatchConditions() FrontDoorMatchConditionArrayOutput {
	return o.ApplyT(func(v CustomRule) []FrontDoorMatchCondition { return v.MatchConditions }).(FrontDoorMatchConditionArrayOutput)
}

func (o CustomRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v CustomRule) int { return v.Priority }).(pulumi.IntOutput)
}

func (o CustomRuleOutput) RateLimitDurationInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomRule) *int { return v.RateLimitDurationInMinutes }).(pulumi.IntPtrOutput)
}

func (o CustomRuleOutput) RateLimitThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomRule) *int { return v.RateLimitThreshold }).(pulumi.IntPtrOutput)
}

func (o CustomRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type CustomRuleArrayOutput struct{ *pulumi.OutputState }

func (CustomRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRule)(nil)).Elem()
}

func (o CustomRuleArrayOutput) ToCustomRuleArrayOutput() CustomRuleArrayOutput {
	return o
}

func (o CustomRuleArrayOutput) ToCustomRuleArrayOutputWithContext(ctx context.Context) CustomRuleArrayOutput {
	return o
}

func (o CustomRuleArrayOutput) Index(i pulumi.IntInput) CustomRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRule {
		return vs[0].([]CustomRule)[vs[1].(int)]
	}).(CustomRuleOutput)
}

type CustomRuleList struct {
	Rules []CustomRule `pulumi:"rules"`
}





type CustomRuleListInput interface {
	pulumi.Input

	ToCustomRuleListOutput() CustomRuleListOutput
	ToCustomRuleListOutputWithContext(context.Context) CustomRuleListOutput
}

type CustomRuleListArgs struct {
	Rules CustomRuleArrayInput `pulumi:"rules"`
}

func (CustomRuleListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleList)(nil)).Elem()
}

func (i CustomRuleListArgs) ToCustomRuleListOutput() CustomRuleListOutput {
	return i.ToCustomRuleListOutputWithContext(context.Background())
}

func (i CustomRuleListArgs) ToCustomRuleListOutputWithContext(ctx context.Context) CustomRuleListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListOutput)
}

func (i CustomRuleListArgs) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return i.ToCustomRuleListPtrOutputWithContext(context.Background())
}

func (i CustomRuleListArgs) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListOutput).ToCustomRuleListPtrOutputWithContext(ctx)
}









type CustomRuleListPtrInput interface {
	pulumi.Input

	ToCustomRuleListPtrOutput() CustomRuleListPtrOutput
	ToCustomRuleListPtrOutputWithContext(context.Context) CustomRuleListPtrOutput
}

type customRuleListPtrType CustomRuleListArgs

func CustomRuleListPtr(v *CustomRuleListArgs) CustomRuleListPtrInput {
	return (*customRuleListPtrType)(v)
}

func (*customRuleListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleList)(nil)).Elem()
}

func (i *customRuleListPtrType) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return i.ToCustomRuleListPtrOutputWithContext(context.Background())
}

func (i *customRuleListPtrType) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomRuleListPtrOutput)
}

type CustomRuleListOutput struct{ *pulumi.OutputState }

func (CustomRuleListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleList)(nil)).Elem()
}

func (o CustomRuleListOutput) ToCustomRuleListOutput() CustomRuleListOutput {
	return o
}

func (o CustomRuleListOutput) ToCustomRuleListOutputWithContext(ctx context.Context) CustomRuleListOutput {
	return o
}

func (o CustomRuleListOutput) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return o.ToCustomRuleListPtrOutputWithContext(context.Background())
}

func (o CustomRuleListOutput) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomRuleList) *CustomRuleList {
		return &v
	}).(CustomRuleListPtrOutput)
}

func (o CustomRuleListOutput) Rules() CustomRuleArrayOutput {
	return o.ApplyT(func(v CustomRuleList) []CustomRule { return v.Rules }).(CustomRuleArrayOutput)
}

type CustomRuleListPtrOutput struct{ *pulumi.OutputState }

func (CustomRuleListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleList)(nil)).Elem()
}

func (o CustomRuleListPtrOutput) ToCustomRuleListPtrOutput() CustomRuleListPtrOutput {
	return o
}

func (o CustomRuleListPtrOutput) ToCustomRuleListPtrOutputWithContext(ctx context.Context) CustomRuleListPtrOutput {
	return o
}

func (o CustomRuleListPtrOutput) Elem() CustomRuleListOutput {
	return o.ApplyT(func(v *CustomRuleList) CustomRuleList {
		if v != nil {
			return *v
		}
		var ret CustomRuleList
		return ret
	}).(CustomRuleListOutput)
}

func (o CustomRuleListPtrOutput) Rules() CustomRuleArrayOutput {
	return o.ApplyT(func(v *CustomRuleList) []CustomRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(CustomRuleArrayOutput)
}

type CustomRuleListResponse struct {
	Rules []CustomRuleResponse `pulumi:"rules"`
}

type CustomRuleListResponseOutput struct{ *pulumi.OutputState }

func (CustomRuleListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleListResponse)(nil)).Elem()
}

func (o CustomRuleListResponseOutput) ToCustomRuleListResponseOutput() CustomRuleListResponseOutput {
	return o
}

func (o CustomRuleListResponseOutput) ToCustomRuleListResponseOutputWithContext(ctx context.Context) CustomRuleListResponseOutput {
	return o
}

func (o CustomRuleListResponseOutput) Rules() CustomRuleResponseArrayOutput {
	return o.ApplyT(func(v CustomRuleListResponse) []CustomRuleResponse { return v.Rules }).(CustomRuleResponseArrayOutput)
}

type CustomRuleListResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomRuleListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomRuleListResponse)(nil)).Elem()
}

func (o CustomRuleListResponsePtrOutput) ToCustomRuleListResponsePtrOutput() CustomRuleListResponsePtrOutput {
	return o
}

func (o CustomRuleListResponsePtrOutput) ToCustomRuleListResponsePtrOutputWithContext(ctx context.Context) CustomRuleListResponsePtrOutput {
	return o
}

func (o CustomRuleListResponsePtrOutput) Elem() CustomRuleListResponseOutput {
	return o.ApplyT(func(v *CustomRuleListResponse) CustomRuleListResponse {
		if v != nil {
			return *v
		}
		var ret CustomRuleListResponse
		return ret
	}).(CustomRuleListResponseOutput)
}

func (o CustomRuleListResponsePtrOutput) Rules() CustomRuleResponseArrayOutput {
	return o.ApplyT(func(v *CustomRuleListResponse) []CustomRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(CustomRuleResponseArrayOutput)
}

type CustomRuleResponse struct {
	Action                     string                            `pulumi:"action"`
	EnabledState               *string                           `pulumi:"enabledState"`
	MatchConditions            []FrontDoorMatchConditionResponse `pulumi:"matchConditions"`
	Name                       *string                           `pulumi:"name"`
	Priority                   int                               `pulumi:"priority"`
	RateLimitDurationInMinutes *int                              `pulumi:"rateLimitDurationInMinutes"`
	RateLimitThreshold         *int                              `pulumi:"rateLimitThreshold"`
	RuleType                   string                            `pulumi:"ruleType"`
}

type CustomRuleResponseOutput struct{ *pulumi.OutputState }

func (CustomRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomRuleResponse)(nil)).Elem()
}

func (o CustomRuleResponseOutput) ToCustomRuleResponseOutput() CustomRuleResponseOutput {
	return o
}

func (o CustomRuleResponseOutput) ToCustomRuleResponseOutputWithContext(ctx context.Context) CustomRuleResponseOutput {
	return o
}

func (o CustomRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o CustomRuleResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRuleResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o CustomRuleResponseOutput) MatchConditions() FrontDoorMatchConditionResponseArrayOutput {
	return o.ApplyT(func(v CustomRuleResponse) []FrontDoorMatchConditionResponse { return v.MatchConditions }).(FrontDoorMatchConditionResponseArrayOutput)
}

func (o CustomRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CustomRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v CustomRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o CustomRuleResponseOutput) RateLimitDurationInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomRuleResponse) *int { return v.RateLimitDurationInMinutes }).(pulumi.IntPtrOutput)
}

func (o CustomRuleResponseOutput) RateLimitThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CustomRuleResponse) *int { return v.RateLimitThreshold }).(pulumi.IntPtrOutput)
}

func (o CustomRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v CustomRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

type CustomRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (CustomRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomRuleResponse)(nil)).Elem()
}

func (o CustomRuleResponseArrayOutput) ToCustomRuleResponseArrayOutput() CustomRuleResponseArrayOutput {
	return o
}

func (o CustomRuleResponseArrayOutput) ToCustomRuleResponseArrayOutputWithContext(ctx context.Context) CustomRuleResponseArrayOutput {
	return o
}

func (o CustomRuleResponseArrayOutput) Index(i pulumi.IntInput) CustomRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomRuleResponse {
		return vs[0].([]CustomRuleResponse)[vs[1].(int)]
	}).(CustomRuleResponseOutput)
}

type FrontDoorManagedRuleGroupOverride struct {
	Exclusions    []ManagedRuleExclusion         `pulumi:"exclusions"`
	RuleGroupName string                         `pulumi:"ruleGroupName"`
	Rules         []FrontDoorManagedRuleOverride `pulumi:"rules"`
}





type FrontDoorManagedRuleGroupOverrideInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleGroupOverrideOutput() FrontDoorManagedRuleGroupOverrideOutput
	ToFrontDoorManagedRuleGroupOverrideOutputWithContext(context.Context) FrontDoorManagedRuleGroupOverrideOutput
}

type FrontDoorManagedRuleGroupOverrideArgs struct {
	Exclusions    ManagedRuleExclusionArrayInput         `pulumi:"exclusions"`
	RuleGroupName pulumi.StringInput                     `pulumi:"ruleGroupName"`
	Rules         FrontDoorManagedRuleOverrideArrayInput `pulumi:"rules"`
}

func (FrontDoorManagedRuleGroupOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleGroupOverride)(nil)).Elem()
}

func (i FrontDoorManagedRuleGroupOverrideArgs) ToFrontDoorManagedRuleGroupOverrideOutput() FrontDoorManagedRuleGroupOverrideOutput {
	return i.ToFrontDoorManagedRuleGroupOverrideOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleGroupOverrideArgs) ToFrontDoorManagedRuleGroupOverrideOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleGroupOverrideOutput)
}





type FrontDoorManagedRuleGroupOverrideArrayInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleGroupOverrideArrayOutput() FrontDoorManagedRuleGroupOverrideArrayOutput
	ToFrontDoorManagedRuleGroupOverrideArrayOutputWithContext(context.Context) FrontDoorManagedRuleGroupOverrideArrayOutput
}

type FrontDoorManagedRuleGroupOverrideArray []FrontDoorManagedRuleGroupOverrideInput

func (FrontDoorManagedRuleGroupOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleGroupOverride)(nil)).Elem()
}

func (i FrontDoorManagedRuleGroupOverrideArray) ToFrontDoorManagedRuleGroupOverrideArrayOutput() FrontDoorManagedRuleGroupOverrideArrayOutput {
	return i.ToFrontDoorManagedRuleGroupOverrideArrayOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleGroupOverrideArray) ToFrontDoorManagedRuleGroupOverrideArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleGroupOverrideArrayOutput)
}

type FrontDoorManagedRuleGroupOverrideOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleGroupOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleGroupOverride)(nil)).Elem()
}

func (o FrontDoorManagedRuleGroupOverrideOutput) ToFrontDoorManagedRuleGroupOverrideOutput() FrontDoorManagedRuleGroupOverrideOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideOutput) ToFrontDoorManagedRuleGroupOverrideOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideOutput) Exclusions() ManagedRuleExclusionArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverride) []ManagedRuleExclusion { return v.Exclusions }).(ManagedRuleExclusionArrayOutput)
}

func (o FrontDoorManagedRuleGroupOverrideOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverride) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o FrontDoorManagedRuleGroupOverrideOutput) Rules() FrontDoorManagedRuleOverrideArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverride) []FrontDoorManagedRuleOverride { return v.Rules }).(FrontDoorManagedRuleOverrideArrayOutput)
}

type FrontDoorManagedRuleGroupOverrideArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleGroupOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleGroupOverride)(nil)).Elem()
}

func (o FrontDoorManagedRuleGroupOverrideArrayOutput) ToFrontDoorManagedRuleGroupOverrideArrayOutput() FrontDoorManagedRuleGroupOverrideArrayOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideArrayOutput) ToFrontDoorManagedRuleGroupOverrideArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideArrayOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleGroupOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleGroupOverride {
		return vs[0].([]FrontDoorManagedRuleGroupOverride)[vs[1].(int)]
	}).(FrontDoorManagedRuleGroupOverrideOutput)
}

type FrontDoorManagedRuleGroupOverrideResponse struct {
	Exclusions    []ManagedRuleExclusionResponse         `pulumi:"exclusions"`
	RuleGroupName string                                 `pulumi:"ruleGroupName"`
	Rules         []FrontDoorManagedRuleOverrideResponse `pulumi:"rules"`
}

type FrontDoorManagedRuleGroupOverrideResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleGroupOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleGroupOverrideResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) ToFrontDoorManagedRuleGroupOverrideResponseOutput() FrontDoorManagedRuleGroupOverrideResponseOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) ToFrontDoorManagedRuleGroupOverrideResponseOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideResponseOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) Exclusions() ManagedRuleExclusionResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverrideResponse) []ManagedRuleExclusionResponse { return v.Exclusions }).(ManagedRuleExclusionResponseArrayOutput)
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverrideResponse) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o FrontDoorManagedRuleGroupOverrideResponseOutput) Rules() FrontDoorManagedRuleOverrideResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleGroupOverrideResponse) []FrontDoorManagedRuleOverrideResponse {
		return v.Rules
	}).(FrontDoorManagedRuleOverrideResponseArrayOutput)
}

type FrontDoorManagedRuleGroupOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleGroupOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleGroupOverrideResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleGroupOverrideResponseArrayOutput) ToFrontDoorManagedRuleGroupOverrideResponseArrayOutput() FrontDoorManagedRuleGroupOverrideResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideResponseArrayOutput) ToFrontDoorManagedRuleGroupOverrideResponseArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleGroupOverrideResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleGroupOverrideResponseArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleGroupOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleGroupOverrideResponse {
		return vs[0].([]FrontDoorManagedRuleGroupOverrideResponse)[vs[1].(int)]
	}).(FrontDoorManagedRuleGroupOverrideResponseOutput)
}

type FrontDoorManagedRuleOverride struct {
	Action       *string                `pulumi:"action"`
	EnabledState *string                `pulumi:"enabledState"`
	Exclusions   []ManagedRuleExclusion `pulumi:"exclusions"`
	RuleId       string                 `pulumi:"ruleId"`
}





type FrontDoorManagedRuleOverrideInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleOverrideOutput() FrontDoorManagedRuleOverrideOutput
	ToFrontDoorManagedRuleOverrideOutputWithContext(context.Context) FrontDoorManagedRuleOverrideOutput
}

type FrontDoorManagedRuleOverrideArgs struct {
	Action       pulumi.StringPtrInput          `pulumi:"action"`
	EnabledState pulumi.StringPtrInput          `pulumi:"enabledState"`
	Exclusions   ManagedRuleExclusionArrayInput `pulumi:"exclusions"`
	RuleId       pulumi.StringInput             `pulumi:"ruleId"`
}

func (FrontDoorManagedRuleOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleOverride)(nil)).Elem()
}

func (i FrontDoorManagedRuleOverrideArgs) ToFrontDoorManagedRuleOverrideOutput() FrontDoorManagedRuleOverrideOutput {
	return i.ToFrontDoorManagedRuleOverrideOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleOverrideArgs) ToFrontDoorManagedRuleOverrideOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleOverrideOutput)
}





type FrontDoorManagedRuleOverrideArrayInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleOverrideArrayOutput() FrontDoorManagedRuleOverrideArrayOutput
	ToFrontDoorManagedRuleOverrideArrayOutputWithContext(context.Context) FrontDoorManagedRuleOverrideArrayOutput
}

type FrontDoorManagedRuleOverrideArray []FrontDoorManagedRuleOverrideInput

func (FrontDoorManagedRuleOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleOverride)(nil)).Elem()
}

func (i FrontDoorManagedRuleOverrideArray) ToFrontDoorManagedRuleOverrideArrayOutput() FrontDoorManagedRuleOverrideArrayOutput {
	return i.ToFrontDoorManagedRuleOverrideArrayOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleOverrideArray) ToFrontDoorManagedRuleOverrideArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleOverrideArrayOutput)
}

type FrontDoorManagedRuleOverrideOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleOverride)(nil)).Elem()
}

func (o FrontDoorManagedRuleOverrideOutput) ToFrontDoorManagedRuleOverrideOutput() FrontDoorManagedRuleOverrideOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideOutput) ToFrontDoorManagedRuleOverrideOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverride) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleOverrideOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverride) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleOverrideOutput) Exclusions() ManagedRuleExclusionArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverride) []ManagedRuleExclusion { return v.Exclusions }).(ManagedRuleExclusionArrayOutput)
}

func (o FrontDoorManagedRuleOverrideOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverride) string { return v.RuleId }).(pulumi.StringOutput)
}

type FrontDoorManagedRuleOverrideArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleOverride)(nil)).Elem()
}

func (o FrontDoorManagedRuleOverrideArrayOutput) ToFrontDoorManagedRuleOverrideArrayOutput() FrontDoorManagedRuleOverrideArrayOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideArrayOutput) ToFrontDoorManagedRuleOverrideArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideArrayOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleOverride {
		return vs[0].([]FrontDoorManagedRuleOverride)[vs[1].(int)]
	}).(FrontDoorManagedRuleOverrideOutput)
}

type FrontDoorManagedRuleOverrideResponse struct {
	Action       *string                        `pulumi:"action"`
	EnabledState *string                        `pulumi:"enabledState"`
	Exclusions   []ManagedRuleExclusionResponse `pulumi:"exclusions"`
	RuleId       string                         `pulumi:"ruleId"`
}

type FrontDoorManagedRuleOverrideResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleOverrideResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleOverrideResponseOutput) ToFrontDoorManagedRuleOverrideResponseOutput() FrontDoorManagedRuleOverrideResponseOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideResponseOutput) ToFrontDoorManagedRuleOverrideResponseOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideResponseOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverrideResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleOverrideResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverrideResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleOverrideResponseOutput) Exclusions() ManagedRuleExclusionResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverrideResponse) []ManagedRuleExclusionResponse { return v.Exclusions }).(ManagedRuleExclusionResponseArrayOutput)
}

func (o FrontDoorManagedRuleOverrideResponseOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleOverrideResponse) string { return v.RuleId }).(pulumi.StringOutput)
}

type FrontDoorManagedRuleOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleOverrideResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleOverrideResponseArrayOutput) ToFrontDoorManagedRuleOverrideResponseArrayOutput() FrontDoorManagedRuleOverrideResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideResponseArrayOutput) ToFrontDoorManagedRuleOverrideResponseArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleOverrideResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleOverrideResponseArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleOverrideResponse {
		return vs[0].([]FrontDoorManagedRuleOverrideResponse)[vs[1].(int)]
	}).(FrontDoorManagedRuleOverrideResponseOutput)
}

type FrontDoorManagedRuleSet struct {
	Exclusions         []ManagedRuleExclusion              `pulumi:"exclusions"`
	RuleGroupOverrides []FrontDoorManagedRuleGroupOverride `pulumi:"ruleGroupOverrides"`
	RuleSetAction      *string                             `pulumi:"ruleSetAction"`
	RuleSetType        string                              `pulumi:"ruleSetType"`
	RuleSetVersion     string                              `pulumi:"ruleSetVersion"`
}





type FrontDoorManagedRuleSetInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleSetOutput() FrontDoorManagedRuleSetOutput
	ToFrontDoorManagedRuleSetOutputWithContext(context.Context) FrontDoorManagedRuleSetOutput
}

type FrontDoorManagedRuleSetArgs struct {
	Exclusions         ManagedRuleExclusionArrayInput              `pulumi:"exclusions"`
	RuleGroupOverrides FrontDoorManagedRuleGroupOverrideArrayInput `pulumi:"ruleGroupOverrides"`
	RuleSetAction      pulumi.StringPtrInput                       `pulumi:"ruleSetAction"`
	RuleSetType        pulumi.StringInput                          `pulumi:"ruleSetType"`
	RuleSetVersion     pulumi.StringInput                          `pulumi:"ruleSetVersion"`
}

func (FrontDoorManagedRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleSet)(nil)).Elem()
}

func (i FrontDoorManagedRuleSetArgs) ToFrontDoorManagedRuleSetOutput() FrontDoorManagedRuleSetOutput {
	return i.ToFrontDoorManagedRuleSetOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleSetArgs) ToFrontDoorManagedRuleSetOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleSetOutput)
}





type FrontDoorManagedRuleSetArrayInput interface {
	pulumi.Input

	ToFrontDoorManagedRuleSetArrayOutput() FrontDoorManagedRuleSetArrayOutput
	ToFrontDoorManagedRuleSetArrayOutputWithContext(context.Context) FrontDoorManagedRuleSetArrayOutput
}

type FrontDoorManagedRuleSetArray []FrontDoorManagedRuleSetInput

func (FrontDoorManagedRuleSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleSet)(nil)).Elem()
}

func (i FrontDoorManagedRuleSetArray) ToFrontDoorManagedRuleSetArrayOutput() FrontDoorManagedRuleSetArrayOutput {
	return i.ToFrontDoorManagedRuleSetArrayOutputWithContext(context.Background())
}

func (i FrontDoorManagedRuleSetArray) ToFrontDoorManagedRuleSetArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorManagedRuleSetArrayOutput)
}

type FrontDoorManagedRuleSetOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleSet)(nil)).Elem()
}

func (o FrontDoorManagedRuleSetOutput) ToFrontDoorManagedRuleSetOutput() FrontDoorManagedRuleSetOutput {
	return o
}

func (o FrontDoorManagedRuleSetOutput) ToFrontDoorManagedRuleSetOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetOutput {
	return o
}

func (o FrontDoorManagedRuleSetOutput) Exclusions() ManagedRuleExclusionArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) []ManagedRuleExclusion { return v.Exclusions }).(ManagedRuleExclusionArrayOutput)
}

func (o FrontDoorManagedRuleSetOutput) RuleGroupOverrides() FrontDoorManagedRuleGroupOverrideArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) []FrontDoorManagedRuleGroupOverride { return v.RuleGroupOverrides }).(FrontDoorManagedRuleGroupOverrideArrayOutput)
}

func (o FrontDoorManagedRuleSetOutput) RuleSetAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) *string { return v.RuleSetAction }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleSetOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o FrontDoorManagedRuleSetOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSet) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type FrontDoorManagedRuleSetArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleSet)(nil)).Elem()
}

func (o FrontDoorManagedRuleSetArrayOutput) ToFrontDoorManagedRuleSetArrayOutput() FrontDoorManagedRuleSetArrayOutput {
	return o
}

func (o FrontDoorManagedRuleSetArrayOutput) ToFrontDoorManagedRuleSetArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetArrayOutput {
	return o
}

func (o FrontDoorManagedRuleSetArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleSet {
		return vs[0].([]FrontDoorManagedRuleSet)[vs[1].(int)]
	}).(FrontDoorManagedRuleSetOutput)
}

type FrontDoorManagedRuleSetResponse struct {
	Exclusions         []ManagedRuleExclusionResponse              `pulumi:"exclusions"`
	RuleGroupOverrides []FrontDoorManagedRuleGroupOverrideResponse `pulumi:"ruleGroupOverrides"`
	RuleSetAction      *string                                     `pulumi:"ruleSetAction"`
	RuleSetType        string                                      `pulumi:"ruleSetType"`
	RuleSetVersion     string                                      `pulumi:"ruleSetVersion"`
}

type FrontDoorManagedRuleSetResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorManagedRuleSetResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleSetResponseOutput) ToFrontDoorManagedRuleSetResponseOutput() FrontDoorManagedRuleSetResponseOutput {
	return o
}

func (o FrontDoorManagedRuleSetResponseOutput) ToFrontDoorManagedRuleSetResponseOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetResponseOutput {
	return o
}

func (o FrontDoorManagedRuleSetResponseOutput) Exclusions() ManagedRuleExclusionResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) []ManagedRuleExclusionResponse { return v.Exclusions }).(ManagedRuleExclusionResponseArrayOutput)
}

func (o FrontDoorManagedRuleSetResponseOutput) RuleGroupOverrides() FrontDoorManagedRuleGroupOverrideResponseArrayOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) []FrontDoorManagedRuleGroupOverrideResponse {
		return v.RuleGroupOverrides
	}).(FrontDoorManagedRuleGroupOverrideResponseArrayOutput)
}

func (o FrontDoorManagedRuleSetResponseOutput) RuleSetAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) *string { return v.RuleSetAction }).(pulumi.StringPtrOutput)
}

func (o FrontDoorManagedRuleSetResponseOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o FrontDoorManagedRuleSetResponseOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorManagedRuleSetResponse) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type FrontDoorManagedRuleSetResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorManagedRuleSetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorManagedRuleSetResponse)(nil)).Elem()
}

func (o FrontDoorManagedRuleSetResponseArrayOutput) ToFrontDoorManagedRuleSetResponseArrayOutput() FrontDoorManagedRuleSetResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleSetResponseArrayOutput) ToFrontDoorManagedRuleSetResponseArrayOutputWithContext(ctx context.Context) FrontDoorManagedRuleSetResponseArrayOutput {
	return o
}

func (o FrontDoorManagedRuleSetResponseArrayOutput) Index(i pulumi.IntInput) FrontDoorManagedRuleSetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorManagedRuleSetResponse {
		return vs[0].([]FrontDoorManagedRuleSetResponse)[vs[1].(int)]
	}).(FrontDoorManagedRuleSetResponseOutput)
}

type FrontDoorMatchCondition struct {
	MatchValue      []string `pulumi:"matchValue"`
	MatchVariable   string   `pulumi:"matchVariable"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}





type FrontDoorMatchConditionInput interface {
	pulumi.Input

	ToFrontDoorMatchConditionOutput() FrontDoorMatchConditionOutput
	ToFrontDoorMatchConditionOutputWithContext(context.Context) FrontDoorMatchConditionOutput
}

type FrontDoorMatchConditionArgs struct {
	MatchValue      pulumi.StringArrayInput `pulumi:"matchValue"`
	MatchVariable   pulumi.StringInput      `pulumi:"matchVariable"`
	NegateCondition pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	Operator        pulumi.StringInput      `pulumi:"operator"`
	Selector        pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms      pulumi.StringArrayInput `pulumi:"transforms"`
}

func (FrontDoorMatchConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorMatchCondition)(nil)).Elem()
}

func (i FrontDoorMatchConditionArgs) ToFrontDoorMatchConditionOutput() FrontDoorMatchConditionOutput {
	return i.ToFrontDoorMatchConditionOutputWithContext(context.Background())
}

func (i FrontDoorMatchConditionArgs) ToFrontDoorMatchConditionOutputWithContext(ctx context.Context) FrontDoorMatchConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorMatchConditionOutput)
}





type FrontDoorMatchConditionArrayInput interface {
	pulumi.Input

	ToFrontDoorMatchConditionArrayOutput() FrontDoorMatchConditionArrayOutput
	ToFrontDoorMatchConditionArrayOutputWithContext(context.Context) FrontDoorMatchConditionArrayOutput
}

type FrontDoorMatchConditionArray []FrontDoorMatchConditionInput

func (FrontDoorMatchConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorMatchCondition)(nil)).Elem()
}

func (i FrontDoorMatchConditionArray) ToFrontDoorMatchConditionArrayOutput() FrontDoorMatchConditionArrayOutput {
	return i.ToFrontDoorMatchConditionArrayOutputWithContext(context.Background())
}

func (i FrontDoorMatchConditionArray) ToFrontDoorMatchConditionArrayOutputWithContext(ctx context.Context) FrontDoorMatchConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorMatchConditionArrayOutput)
}

type FrontDoorMatchConditionOutput struct{ *pulumi.OutputState }

func (FrontDoorMatchConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorMatchCondition)(nil)).Elem()
}

func (o FrontDoorMatchConditionOutput) ToFrontDoorMatchConditionOutput() FrontDoorMatchConditionOutput {
	return o
}

func (o FrontDoorMatchConditionOutput) ToFrontDoorMatchConditionOutputWithContext(ctx context.Context) FrontDoorMatchConditionOutput {
	return o
}

func (o FrontDoorMatchConditionOutput) MatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) []string { return v.MatchValue }).(pulumi.StringArrayOutput)
}

func (o FrontDoorMatchConditionOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o FrontDoorMatchConditionOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o FrontDoorMatchConditionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) string { return v.Operator }).(pulumi.StringOutput)
}

func (o FrontDoorMatchConditionOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o FrontDoorMatchConditionOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontDoorMatchCondition) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type FrontDoorMatchConditionArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorMatchConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorMatchCondition)(nil)).Elem()
}

func (o FrontDoorMatchConditionArrayOutput) ToFrontDoorMatchConditionArrayOutput() FrontDoorMatchConditionArrayOutput {
	return o
}

func (o FrontDoorMatchConditionArrayOutput) ToFrontDoorMatchConditionArrayOutputWithContext(ctx context.Context) FrontDoorMatchConditionArrayOutput {
	return o
}

func (o FrontDoorMatchConditionArrayOutput) Index(i pulumi.IntInput) FrontDoorMatchConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorMatchCondition {
		return vs[0].([]FrontDoorMatchCondition)[vs[1].(int)]
	}).(FrontDoorMatchConditionOutput)
}

type FrontDoorMatchConditionResponse struct {
	MatchValue      []string `pulumi:"matchValue"`
	MatchVariable   string   `pulumi:"matchVariable"`
	NegateCondition *bool    `pulumi:"negateCondition"`
	Operator        string   `pulumi:"operator"`
	Selector        *string  `pulumi:"selector"`
	Transforms      []string `pulumi:"transforms"`
}

type FrontDoorMatchConditionResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorMatchConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorMatchConditionResponse)(nil)).Elem()
}

func (o FrontDoorMatchConditionResponseOutput) ToFrontDoorMatchConditionResponseOutput() FrontDoorMatchConditionResponseOutput {
	return o
}

func (o FrontDoorMatchConditionResponseOutput) ToFrontDoorMatchConditionResponseOutputWithContext(ctx context.Context) FrontDoorMatchConditionResponseOutput {
	return o
}

func (o FrontDoorMatchConditionResponseOutput) MatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) []string { return v.MatchValue }).(pulumi.StringArrayOutput)
}

func (o FrontDoorMatchConditionResponseOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o FrontDoorMatchConditionResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o FrontDoorMatchConditionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o FrontDoorMatchConditionResponseOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o FrontDoorMatchConditionResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FrontDoorMatchConditionResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type FrontDoorMatchConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontDoorMatchConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontDoorMatchConditionResponse)(nil)).Elem()
}

func (o FrontDoorMatchConditionResponseArrayOutput) ToFrontDoorMatchConditionResponseArrayOutput() FrontDoorMatchConditionResponseArrayOutput {
	return o
}

func (o FrontDoorMatchConditionResponseArrayOutput) ToFrontDoorMatchConditionResponseArrayOutputWithContext(ctx context.Context) FrontDoorMatchConditionResponseArrayOutput {
	return o
}

func (o FrontDoorMatchConditionResponseArrayOutput) Index(i pulumi.IntInput) FrontDoorMatchConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontDoorMatchConditionResponse {
		return vs[0].([]FrontDoorMatchConditionResponse)[vs[1].(int)]
	}).(FrontDoorMatchConditionResponseOutput)
}

type FrontDoorPolicySettings struct {
	CustomBlockResponseBody       *string `pulumi:"customBlockResponseBody"`
	CustomBlockResponseStatusCode *int    `pulumi:"customBlockResponseStatusCode"`
	EnabledState                  *string `pulumi:"enabledState"`
	Mode                          *string `pulumi:"mode"`
	RedirectUrl                   *string `pulumi:"redirectUrl"`
	RequestBodyCheck              *string `pulumi:"requestBodyCheck"`
}





type FrontDoorPolicySettingsInput interface {
	pulumi.Input

	ToFrontDoorPolicySettingsOutput() FrontDoorPolicySettingsOutput
	ToFrontDoorPolicySettingsOutputWithContext(context.Context) FrontDoorPolicySettingsOutput
}

type FrontDoorPolicySettingsArgs struct {
	CustomBlockResponseBody       pulumi.StringPtrInput `pulumi:"customBlockResponseBody"`
	CustomBlockResponseStatusCode pulumi.IntPtrInput    `pulumi:"customBlockResponseStatusCode"`
	EnabledState                  pulumi.StringPtrInput `pulumi:"enabledState"`
	Mode                          pulumi.StringPtrInput `pulumi:"mode"`
	RedirectUrl                   pulumi.StringPtrInput `pulumi:"redirectUrl"`
	RequestBodyCheck              pulumi.StringPtrInput `pulumi:"requestBodyCheck"`
}

func (FrontDoorPolicySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorPolicySettings)(nil)).Elem()
}

func (i FrontDoorPolicySettingsArgs) ToFrontDoorPolicySettingsOutput() FrontDoorPolicySettingsOutput {
	return i.ToFrontDoorPolicySettingsOutputWithContext(context.Background())
}

func (i FrontDoorPolicySettingsArgs) ToFrontDoorPolicySettingsOutputWithContext(ctx context.Context) FrontDoorPolicySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorPolicySettingsOutput)
}

func (i FrontDoorPolicySettingsArgs) ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput {
	return i.ToFrontDoorPolicySettingsPtrOutputWithContext(context.Background())
}

func (i FrontDoorPolicySettingsArgs) ToFrontDoorPolicySettingsPtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorPolicySettingsOutput).ToFrontDoorPolicySettingsPtrOutputWithContext(ctx)
}









type FrontDoorPolicySettingsPtrInput interface {
	pulumi.Input

	ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput
	ToFrontDoorPolicySettingsPtrOutputWithContext(context.Context) FrontDoorPolicySettingsPtrOutput
}

type frontDoorPolicySettingsPtrType FrontDoorPolicySettingsArgs

func FrontDoorPolicySettingsPtr(v *FrontDoorPolicySettingsArgs) FrontDoorPolicySettingsPtrInput {
	return (*frontDoorPolicySettingsPtrType)(v)
}

func (*frontDoorPolicySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontDoorPolicySettings)(nil)).Elem()
}

func (i *frontDoorPolicySettingsPtrType) ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput {
	return i.ToFrontDoorPolicySettingsPtrOutputWithContext(context.Background())
}

func (i *frontDoorPolicySettingsPtrType) ToFrontDoorPolicySettingsPtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontDoorPolicySettingsPtrOutput)
}

type FrontDoorPolicySettingsOutput struct{ *pulumi.OutputState }

func (FrontDoorPolicySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorPolicySettings)(nil)).Elem()
}

func (o FrontDoorPolicySettingsOutput) ToFrontDoorPolicySettingsOutput() FrontDoorPolicySettingsOutput {
	return o
}

func (o FrontDoorPolicySettingsOutput) ToFrontDoorPolicySettingsOutputWithContext(ctx context.Context) FrontDoorPolicySettingsOutput {
	return o
}

func (o FrontDoorPolicySettingsOutput) ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput {
	return o.ToFrontDoorPolicySettingsPtrOutputWithContext(context.Background())
}

func (o FrontDoorPolicySettingsOutput) ToFrontDoorPolicySettingsPtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontDoorPolicySettings) *FrontDoorPolicySettings {
		return &v
	}).(FrontDoorPolicySettingsPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) CustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.CustomBlockResponseBody }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) CustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *int { return v.CustomBlockResponseStatusCode }).(pulumi.IntPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsOutput) RequestBodyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettings) *string { return v.RequestBodyCheck }).(pulumi.StringPtrOutput)
}

type FrontDoorPolicySettingsPtrOutput struct{ *pulumi.OutputState }

func (FrontDoorPolicySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontDoorPolicySettings)(nil)).Elem()
}

func (o FrontDoorPolicySettingsPtrOutput) ToFrontDoorPolicySettingsPtrOutput() FrontDoorPolicySettingsPtrOutput {
	return o
}

func (o FrontDoorPolicySettingsPtrOutput) ToFrontDoorPolicySettingsPtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsPtrOutput {
	return o
}

func (o FrontDoorPolicySettingsPtrOutput) Elem() FrontDoorPolicySettingsOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) FrontDoorPolicySettings {
		if v != nil {
			return *v
		}
		var ret FrontDoorPolicySettings
		return ret
	}).(FrontDoorPolicySettingsOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) CustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.CustomBlockResponseBody
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) CustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *int {
		if v == nil {
			return nil
		}
		return v.CustomBlockResponseStatusCode
	}).(pulumi.IntPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.EnabledState
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsPtrOutput) RequestBodyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettings) *string {
		if v == nil {
			return nil
		}
		return v.RequestBodyCheck
	}).(pulumi.StringPtrOutput)
}

type FrontDoorPolicySettingsResponse struct {
	CustomBlockResponseBody       *string `pulumi:"customBlockResponseBody"`
	CustomBlockResponseStatusCode *int    `pulumi:"customBlockResponseStatusCode"`
	EnabledState                  *string `pulumi:"enabledState"`
	Mode                          *string `pulumi:"mode"`
	RedirectUrl                   *string `pulumi:"redirectUrl"`
	RequestBodyCheck              *string `pulumi:"requestBodyCheck"`
}

type FrontDoorPolicySettingsResponseOutput struct{ *pulumi.OutputState }

func (FrontDoorPolicySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontDoorPolicySettingsResponse)(nil)).Elem()
}

func (o FrontDoorPolicySettingsResponseOutput) ToFrontDoorPolicySettingsResponseOutput() FrontDoorPolicySettingsResponseOutput {
	return o
}

func (o FrontDoorPolicySettingsResponseOutput) ToFrontDoorPolicySettingsResponseOutputWithContext(ctx context.Context) FrontDoorPolicySettingsResponseOutput {
	return o
}

func (o FrontDoorPolicySettingsResponseOutput) CustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.CustomBlockResponseBody }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) CustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *int { return v.CustomBlockResponseStatusCode }).(pulumi.IntPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponseOutput) RequestBodyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontDoorPolicySettingsResponse) *string { return v.RequestBodyCheck }).(pulumi.StringPtrOutput)
}

type FrontDoorPolicySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (FrontDoorPolicySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontDoorPolicySettingsResponse)(nil)).Elem()
}

func (o FrontDoorPolicySettingsResponsePtrOutput) ToFrontDoorPolicySettingsResponsePtrOutput() FrontDoorPolicySettingsResponsePtrOutput {
	return o
}

func (o FrontDoorPolicySettingsResponsePtrOutput) ToFrontDoorPolicySettingsResponsePtrOutputWithContext(ctx context.Context) FrontDoorPolicySettingsResponsePtrOutput {
	return o
}

func (o FrontDoorPolicySettingsResponsePtrOutput) Elem() FrontDoorPolicySettingsResponseOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) FrontDoorPolicySettingsResponse {
		if v != nil {
			return *v
		}
		var ret FrontDoorPolicySettingsResponse
		return ret
	}).(FrontDoorPolicySettingsResponseOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) CustomBlockResponseBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomBlockResponseBody
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) CustomBlockResponseStatusCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.CustomBlockResponseStatusCode
	}).(pulumi.IntPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EnabledState
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RedirectUrl
	}).(pulumi.StringPtrOutput)
}

func (o FrontDoorPolicySettingsResponsePtrOutput) RequestBodyCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontDoorPolicySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RequestBodyCheck
	}).(pulumi.StringPtrOutput)
}

type FrontendEndpointLinkResponse struct {
	Id *string `pulumi:"id"`
}

type FrontendEndpointLinkResponseOutput struct{ *pulumi.OutputState }

func (FrontendEndpointLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointLinkResponse)(nil)).Elem()
}

func (o FrontendEndpointLinkResponseOutput) ToFrontendEndpointLinkResponseOutput() FrontendEndpointLinkResponseOutput {
	return o
}

func (o FrontendEndpointLinkResponseOutput) ToFrontendEndpointLinkResponseOutputWithContext(ctx context.Context) FrontendEndpointLinkResponseOutput {
	return o
}

func (o FrontendEndpointLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type FrontendEndpointLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontendEndpointLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendEndpointLinkResponse)(nil)).Elem()
}

func (o FrontendEndpointLinkResponseArrayOutput) ToFrontendEndpointLinkResponseArrayOutput() FrontendEndpointLinkResponseArrayOutput {
	return o
}

func (o FrontendEndpointLinkResponseArrayOutput) ToFrontendEndpointLinkResponseArrayOutputWithContext(ctx context.Context) FrontendEndpointLinkResponseArrayOutput {
	return o
}

func (o FrontendEndpointLinkResponseArrayOutput) Index(i pulumi.IntInput) FrontendEndpointLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendEndpointLinkResponse {
		return vs[0].([]FrontendEndpointLinkResponse)[vs[1].(int)]
	}).(FrontendEndpointLinkResponseOutput)
}

type ManagedRuleExclusion struct {
	MatchVariable         string `pulumi:"matchVariable"`
	Selector              string `pulumi:"selector"`
	SelectorMatchOperator string `pulumi:"selectorMatchOperator"`
}





type ManagedRuleExclusionInput interface {
	pulumi.Input

	ToManagedRuleExclusionOutput() ManagedRuleExclusionOutput
	ToManagedRuleExclusionOutputWithContext(context.Context) ManagedRuleExclusionOutput
}

type ManagedRuleExclusionArgs struct {
	MatchVariable         pulumi.StringInput `pulumi:"matchVariable"`
	Selector              pulumi.StringInput `pulumi:"selector"`
	SelectorMatchOperator pulumi.StringInput `pulumi:"selectorMatchOperator"`
}

func (ManagedRuleExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleExclusion)(nil)).Elem()
}

func (i ManagedRuleExclusionArgs) ToManagedRuleExclusionOutput() ManagedRuleExclusionOutput {
	return i.ToManagedRuleExclusionOutputWithContext(context.Background())
}

func (i ManagedRuleExclusionArgs) ToManagedRuleExclusionOutputWithContext(ctx context.Context) ManagedRuleExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleExclusionOutput)
}





type ManagedRuleExclusionArrayInput interface {
	pulumi.Input

	ToManagedRuleExclusionArrayOutput() ManagedRuleExclusionArrayOutput
	ToManagedRuleExclusionArrayOutputWithContext(context.Context) ManagedRuleExclusionArrayOutput
}

type ManagedRuleExclusionArray []ManagedRuleExclusionInput

func (ManagedRuleExclusionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleExclusion)(nil)).Elem()
}

func (i ManagedRuleExclusionArray) ToManagedRuleExclusionArrayOutput() ManagedRuleExclusionArrayOutput {
	return i.ToManagedRuleExclusionArrayOutputWithContext(context.Background())
}

func (i ManagedRuleExclusionArray) ToManagedRuleExclusionArrayOutputWithContext(ctx context.Context) ManagedRuleExclusionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleExclusionArrayOutput)
}

type ManagedRuleExclusionOutput struct{ *pulumi.OutputState }

func (ManagedRuleExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleExclusion)(nil)).Elem()
}

func (o ManagedRuleExclusionOutput) ToManagedRuleExclusionOutput() ManagedRuleExclusionOutput {
	return o
}

func (o ManagedRuleExclusionOutput) ToManagedRuleExclusionOutputWithContext(ctx context.Context) ManagedRuleExclusionOutput {
	return o
}

func (o ManagedRuleExclusionOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusion) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o ManagedRuleExclusionOutput) Selector() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusion) string { return v.Selector }).(pulumi.StringOutput)
}

func (o ManagedRuleExclusionOutput) SelectorMatchOperator() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusion) string { return v.SelectorMatchOperator }).(pulumi.StringOutput)
}

type ManagedRuleExclusionArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleExclusionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleExclusion)(nil)).Elem()
}

func (o ManagedRuleExclusionArrayOutput) ToManagedRuleExclusionArrayOutput() ManagedRuleExclusionArrayOutput {
	return o
}

func (o ManagedRuleExclusionArrayOutput) ToManagedRuleExclusionArrayOutputWithContext(ctx context.Context) ManagedRuleExclusionArrayOutput {
	return o
}

func (o ManagedRuleExclusionArrayOutput) Index(i pulumi.IntInput) ManagedRuleExclusionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleExclusion {
		return vs[0].([]ManagedRuleExclusion)[vs[1].(int)]
	}).(ManagedRuleExclusionOutput)
}

type ManagedRuleExclusionResponse struct {
	MatchVariable         string `pulumi:"matchVariable"`
	Selector              string `pulumi:"selector"`
	SelectorMatchOperator string `pulumi:"selectorMatchOperator"`
}

type ManagedRuleExclusionResponseOutput struct{ *pulumi.OutputState }

func (ManagedRuleExclusionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleExclusionResponse)(nil)).Elem()
}

func (o ManagedRuleExclusionResponseOutput) ToManagedRuleExclusionResponseOutput() ManagedRuleExclusionResponseOutput {
	return o
}

func (o ManagedRuleExclusionResponseOutput) ToManagedRuleExclusionResponseOutputWithContext(ctx context.Context) ManagedRuleExclusionResponseOutput {
	return o
}

func (o ManagedRuleExclusionResponseOutput) MatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusionResponse) string { return v.MatchVariable }).(pulumi.StringOutput)
}

func (o ManagedRuleExclusionResponseOutput) Selector() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusionResponse) string { return v.Selector }).(pulumi.StringOutput)
}

func (o ManagedRuleExclusionResponseOutput) SelectorMatchOperator() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedRuleExclusionResponse) string { return v.SelectorMatchOperator }).(pulumi.StringOutput)
}

type ManagedRuleExclusionResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedRuleExclusionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedRuleExclusionResponse)(nil)).Elem()
}

func (o ManagedRuleExclusionResponseArrayOutput) ToManagedRuleExclusionResponseArrayOutput() ManagedRuleExclusionResponseArrayOutput {
	return o
}

func (o ManagedRuleExclusionResponseArrayOutput) ToManagedRuleExclusionResponseArrayOutputWithContext(ctx context.Context) ManagedRuleExclusionResponseArrayOutput {
	return o
}

func (o ManagedRuleExclusionResponseArrayOutput) Index(i pulumi.IntInput) ManagedRuleExclusionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedRuleExclusionResponse {
		return vs[0].([]ManagedRuleExclusionResponse)[vs[1].(int)]
	}).(ManagedRuleExclusionResponseOutput)
}

type ManagedRuleSetList struct {
	ManagedRuleSets []FrontDoorManagedRuleSet `pulumi:"managedRuleSets"`
}





type ManagedRuleSetListInput interface {
	pulumi.Input

	ToManagedRuleSetListOutput() ManagedRuleSetListOutput
	ToManagedRuleSetListOutputWithContext(context.Context) ManagedRuleSetListOutput
}

type ManagedRuleSetListArgs struct {
	ManagedRuleSets FrontDoorManagedRuleSetArrayInput `pulumi:"managedRuleSets"`
}

func (ManagedRuleSetListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetList)(nil)).Elem()
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListOutput() ManagedRuleSetListOutput {
	return i.ToManagedRuleSetListOutputWithContext(context.Background())
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListOutputWithContext(ctx context.Context) ManagedRuleSetListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListOutput)
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return i.ToManagedRuleSetListPtrOutputWithContext(context.Background())
}

func (i ManagedRuleSetListArgs) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListOutput).ToManagedRuleSetListPtrOutputWithContext(ctx)
}









type ManagedRuleSetListPtrInput interface {
	pulumi.Input

	ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput
	ToManagedRuleSetListPtrOutputWithContext(context.Context) ManagedRuleSetListPtrOutput
}

type managedRuleSetListPtrType ManagedRuleSetListArgs

func ManagedRuleSetListPtr(v *ManagedRuleSetListArgs) ManagedRuleSetListPtrInput {
	return (*managedRuleSetListPtrType)(v)
}

func (*managedRuleSetListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetList)(nil)).Elem()
}

func (i *managedRuleSetListPtrType) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return i.ToManagedRuleSetListPtrOutputWithContext(context.Background())
}

func (i *managedRuleSetListPtrType) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedRuleSetListPtrOutput)
}

type ManagedRuleSetListOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetList)(nil)).Elem()
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListOutput() ManagedRuleSetListOutput {
	return o
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListOutputWithContext(ctx context.Context) ManagedRuleSetListOutput {
	return o
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return o.ToManagedRuleSetListPtrOutputWithContext(context.Background())
}

func (o ManagedRuleSetListOutput) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedRuleSetList) *ManagedRuleSetList {
		return &v
	}).(ManagedRuleSetListPtrOutput)
}

func (o ManagedRuleSetListOutput) ManagedRuleSets() FrontDoorManagedRuleSetArrayOutput {
	return o.ApplyT(func(v ManagedRuleSetList) []FrontDoorManagedRuleSet { return v.ManagedRuleSets }).(FrontDoorManagedRuleSetArrayOutput)
}

type ManagedRuleSetListPtrOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetList)(nil)).Elem()
}

func (o ManagedRuleSetListPtrOutput) ToManagedRuleSetListPtrOutput() ManagedRuleSetListPtrOutput {
	return o
}

func (o ManagedRuleSetListPtrOutput) ToManagedRuleSetListPtrOutputWithContext(ctx context.Context) ManagedRuleSetListPtrOutput {
	return o
}

func (o ManagedRuleSetListPtrOutput) Elem() ManagedRuleSetListOutput {
	return o.ApplyT(func(v *ManagedRuleSetList) ManagedRuleSetList {
		if v != nil {
			return *v
		}
		var ret ManagedRuleSetList
		return ret
	}).(ManagedRuleSetListOutput)
}

func (o ManagedRuleSetListPtrOutput) ManagedRuleSets() FrontDoorManagedRuleSetArrayOutput {
	return o.ApplyT(func(v *ManagedRuleSetList) []FrontDoorManagedRuleSet {
		if v == nil {
			return nil
		}
		return v.ManagedRuleSets
	}).(FrontDoorManagedRuleSetArrayOutput)
}

type ManagedRuleSetListResponse struct {
	ManagedRuleSets []FrontDoorManagedRuleSetResponse `pulumi:"managedRuleSets"`
}

type ManagedRuleSetListResponseOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedRuleSetListResponse)(nil)).Elem()
}

func (o ManagedRuleSetListResponseOutput) ToManagedRuleSetListResponseOutput() ManagedRuleSetListResponseOutput {
	return o
}

func (o ManagedRuleSetListResponseOutput) ToManagedRuleSetListResponseOutputWithContext(ctx context.Context) ManagedRuleSetListResponseOutput {
	return o
}

func (o ManagedRuleSetListResponseOutput) ManagedRuleSets() FrontDoorManagedRuleSetResponseArrayOutput {
	return o.ApplyT(func(v ManagedRuleSetListResponse) []FrontDoorManagedRuleSetResponse { return v.ManagedRuleSets }).(FrontDoorManagedRuleSetResponseArrayOutput)
}

type ManagedRuleSetListResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedRuleSetListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedRuleSetListResponse)(nil)).Elem()
}

func (o ManagedRuleSetListResponsePtrOutput) ToManagedRuleSetListResponsePtrOutput() ManagedRuleSetListResponsePtrOutput {
	return o
}

func (o ManagedRuleSetListResponsePtrOutput) ToManagedRuleSetListResponsePtrOutputWithContext(ctx context.Context) ManagedRuleSetListResponsePtrOutput {
	return o
}

func (o ManagedRuleSetListResponsePtrOutput) Elem() ManagedRuleSetListResponseOutput {
	return o.ApplyT(func(v *ManagedRuleSetListResponse) ManagedRuleSetListResponse {
		if v != nil {
			return *v
		}
		var ret ManagedRuleSetListResponse
		return ret
	}).(ManagedRuleSetListResponseOutput)
}

func (o ManagedRuleSetListResponsePtrOutput) ManagedRuleSets() FrontDoorManagedRuleSetResponseArrayOutput {
	return o.ApplyT(func(v *ManagedRuleSetListResponse) []FrontDoorManagedRuleSetResponse {
		if v == nil {
			return nil
		}
		return v.ManagedRuleSets
	}).(FrontDoorManagedRuleSetResponseArrayOutput)
}

type RoutingRuleLinkResponse struct {
	Id *string `pulumi:"id"`
}

type RoutingRuleLinkResponseOutput struct{ *pulumi.OutputState }

func (RoutingRuleLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleLinkResponse)(nil)).Elem()
}

func (o RoutingRuleLinkResponseOutput) ToRoutingRuleLinkResponseOutput() RoutingRuleLinkResponseOutput {
	return o
}

func (o RoutingRuleLinkResponseOutput) ToRoutingRuleLinkResponseOutputWithContext(ctx context.Context) RoutingRuleLinkResponseOutput {
	return o
}

func (o RoutingRuleLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type RoutingRuleLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingRuleLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingRuleLinkResponse)(nil)).Elem()
}

func (o RoutingRuleLinkResponseArrayOutput) ToRoutingRuleLinkResponseArrayOutput() RoutingRuleLinkResponseArrayOutput {
	return o
}

func (o RoutingRuleLinkResponseArrayOutput) ToRoutingRuleLinkResponseArrayOutputWithContext(ctx context.Context) RoutingRuleLinkResponseArrayOutput {
	return o
}

func (o RoutingRuleLinkResponseArrayOutput) Index(i pulumi.IntInput) RoutingRuleLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingRuleLinkResponse {
		return vs[0].([]RoutingRuleLinkResponse)[vs[1].(int)]
	}).(RoutingRuleLinkResponseOutput)
}

type SecurityPolicyLinkResponse struct {
	Id *string `pulumi:"id"`
}

type SecurityPolicyLinkResponseOutput struct{ *pulumi.OutputState }

func (SecurityPolicyLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicyLinkResponse)(nil)).Elem()
}

func (o SecurityPolicyLinkResponseOutput) ToSecurityPolicyLinkResponseOutput() SecurityPolicyLinkResponseOutput {
	return o
}

func (o SecurityPolicyLinkResponseOutput) ToSecurityPolicyLinkResponseOutputWithContext(ctx context.Context) SecurityPolicyLinkResponseOutput {
	return o
}

func (o SecurityPolicyLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityPolicyLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SecurityPolicyLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (SecurityPolicyLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityPolicyLinkResponse)(nil)).Elem()
}

func (o SecurityPolicyLinkResponseArrayOutput) ToSecurityPolicyLinkResponseArrayOutput() SecurityPolicyLinkResponseArrayOutput {
	return o
}

func (o SecurityPolicyLinkResponseArrayOutput) ToSecurityPolicyLinkResponseArrayOutputWithContext(ctx context.Context) SecurityPolicyLinkResponseArrayOutput {
	return o
}

func (o SecurityPolicyLinkResponseArrayOutput) Index(i pulumi.IntInput) SecurityPolicyLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityPolicyLinkResponse {
		return vs[0].([]SecurityPolicyLinkResponse)[vs[1].(int)]
	}).(SecurityPolicyLinkResponseOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomRuleOutput{})
	pulumi.RegisterOutputType(CustomRuleArrayOutput{})
	pulumi.RegisterOutputType(CustomRuleListOutput{})
	pulumi.RegisterOutputType(CustomRuleListPtrOutput{})
	pulumi.RegisterOutputType(CustomRuleListResponseOutput{})
	pulumi.RegisterOutputType(CustomRuleListResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomRuleResponseOutput{})
	pulumi.RegisterOutputType(CustomRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleGroupOverrideOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleGroupOverrideArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleGroupOverrideResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleGroupOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleOverrideOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleOverrideArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleOverrideResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleSetOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleSetArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleSetResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorManagedRuleSetResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorMatchConditionOutput{})
	pulumi.RegisterOutputType(FrontDoorMatchConditionArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorMatchConditionResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorMatchConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontDoorPolicySettingsOutput{})
	pulumi.RegisterOutputType(FrontDoorPolicySettingsPtrOutput{})
	pulumi.RegisterOutputType(FrontDoorPolicySettingsResponseOutput{})
	pulumi.RegisterOutputType(FrontDoorPolicySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(FrontendEndpointLinkResponseOutput{})
	pulumi.RegisterOutputType(FrontendEndpointLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleExclusionOutput{})
	pulumi.RegisterOutputType(ManagedRuleExclusionArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleExclusionResponseOutput{})
	pulumi.RegisterOutputType(ManagedRuleExclusionResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListPtrOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListResponseOutput{})
	pulumi.RegisterOutputType(ManagedRuleSetListResponsePtrOutput{})
	pulumi.RegisterOutputType(RoutingRuleLinkResponseOutput{})
	pulumi.RegisterOutputType(RoutingRuleLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(SecurityPolicyLinkResponseOutput{})
	pulumi.RegisterOutputType(SecurityPolicyLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
