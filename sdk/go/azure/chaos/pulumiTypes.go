


package chaos

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Branch struct {
	Actions []interface{} `pulumi:"actions"`
	Name    string        `pulumi:"name"`
}





type BranchInput interface {
	pulumi.Input

	ToBranchOutput() BranchOutput
	ToBranchOutputWithContext(context.Context) BranchOutput
}

type BranchArgs struct {
	Actions pulumi.ArrayInput  `pulumi:"actions"`
	Name    pulumi.StringInput `pulumi:"name"`
}

func (BranchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Branch)(nil)).Elem()
}

func (i BranchArgs) ToBranchOutput() BranchOutput {
	return i.ToBranchOutputWithContext(context.Background())
}

func (i BranchArgs) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchOutput)
}





type BranchArrayInput interface {
	pulumi.Input

	ToBranchArrayOutput() BranchArrayOutput
	ToBranchArrayOutputWithContext(context.Context) BranchArrayOutput
}

type BranchArray []BranchInput

func (BranchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Branch)(nil)).Elem()
}

func (i BranchArray) ToBranchArrayOutput() BranchArrayOutput {
	return i.ToBranchArrayOutputWithContext(context.Background())
}

func (i BranchArray) ToBranchArrayOutputWithContext(ctx context.Context) BranchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchArrayOutput)
}

type BranchOutput struct{ *pulumi.OutputState }

func (BranchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Branch)(nil)).Elem()
}

func (o BranchOutput) ToBranchOutput() BranchOutput {
	return o
}

func (o BranchOutput) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return o
}

func (o BranchOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v Branch) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o BranchOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Branch) string { return v.Name }).(pulumi.StringOutput)
}

type BranchArrayOutput struct{ *pulumi.OutputState }

func (BranchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Branch)(nil)).Elem()
}

func (o BranchArrayOutput) ToBranchArrayOutput() BranchArrayOutput {
	return o
}

func (o BranchArrayOutput) ToBranchArrayOutputWithContext(ctx context.Context) BranchArrayOutput {
	return o
}

func (o BranchArrayOutput) Index(i pulumi.IntInput) BranchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Branch {
		return vs[0].([]Branch)[vs[1].(int)]
	}).(BranchOutput)
}

type BranchResponse struct {
	Actions []interface{} `pulumi:"actions"`
	Name    string        `pulumi:"name"`
}





type BranchResponseInput interface {
	pulumi.Input

	ToBranchResponseOutput() BranchResponseOutput
	ToBranchResponseOutputWithContext(context.Context) BranchResponseOutput
}

type BranchResponseArgs struct {
	Actions pulumi.ArrayInput  `pulumi:"actions"`
	Name    pulumi.StringInput `pulumi:"name"`
}

func (BranchResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchResponse)(nil)).Elem()
}

func (i BranchResponseArgs) ToBranchResponseOutput() BranchResponseOutput {
	return i.ToBranchResponseOutputWithContext(context.Background())
}

func (i BranchResponseArgs) ToBranchResponseOutputWithContext(ctx context.Context) BranchResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchResponseOutput)
}





type BranchResponseArrayInput interface {
	pulumi.Input

	ToBranchResponseArrayOutput() BranchResponseArrayOutput
	ToBranchResponseArrayOutputWithContext(context.Context) BranchResponseArrayOutput
}

type BranchResponseArray []BranchResponseInput

func (BranchResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BranchResponse)(nil)).Elem()
}

func (i BranchResponseArray) ToBranchResponseArrayOutput() BranchResponseArrayOutput {
	return i.ToBranchResponseArrayOutputWithContext(context.Background())
}

func (i BranchResponseArray) ToBranchResponseArrayOutputWithContext(ctx context.Context) BranchResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchResponseArrayOutput)
}

type BranchResponseOutput struct{ *pulumi.OutputState }

func (BranchResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BranchResponse)(nil)).Elem()
}

func (o BranchResponseOutput) ToBranchResponseOutput() BranchResponseOutput {
	return o
}

func (o BranchResponseOutput) ToBranchResponseOutputWithContext(ctx context.Context) BranchResponseOutput {
	return o
}

func (o BranchResponseOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v BranchResponse) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o BranchResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BranchResponse) string { return v.Name }).(pulumi.StringOutput)
}

type BranchResponseArrayOutput struct{ *pulumi.OutputState }

func (BranchResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BranchResponse)(nil)).Elem()
}

func (o BranchResponseArrayOutput) ToBranchResponseArrayOutput() BranchResponseArrayOutput {
	return o
}

func (o BranchResponseArrayOutput) ToBranchResponseArrayOutputWithContext(ctx context.Context) BranchResponseArrayOutput {
	return o
}

func (o BranchResponseArrayOutput) Index(i pulumi.IntInput) BranchResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BranchResponse {
		return vs[0].([]BranchResponse)[vs[1].(int)]
	}).(BranchResponseOutput)
}

type CapabilityPropertiesResponse struct {
	Description      string `pulumi:"description"`
	ParametersSchema string `pulumi:"parametersSchema"`
	Publisher        string `pulumi:"publisher"`
	TargetType       string `pulumi:"targetType"`
	Urn              string `pulumi:"urn"`
}





type CapabilityPropertiesResponseInput interface {
	pulumi.Input

	ToCapabilityPropertiesResponseOutput() CapabilityPropertiesResponseOutput
	ToCapabilityPropertiesResponseOutputWithContext(context.Context) CapabilityPropertiesResponseOutput
}

type CapabilityPropertiesResponseArgs struct {
	Description      pulumi.StringInput `pulumi:"description"`
	ParametersSchema pulumi.StringInput `pulumi:"parametersSchema"`
	Publisher        pulumi.StringInput `pulumi:"publisher"`
	TargetType       pulumi.StringInput `pulumi:"targetType"`
	Urn              pulumi.StringInput `pulumi:"urn"`
}

func (CapabilityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityPropertiesResponse)(nil)).Elem()
}

func (i CapabilityPropertiesResponseArgs) ToCapabilityPropertiesResponseOutput() CapabilityPropertiesResponseOutput {
	return i.ToCapabilityPropertiesResponseOutputWithContext(context.Background())
}

func (i CapabilityPropertiesResponseArgs) ToCapabilityPropertiesResponseOutputWithContext(ctx context.Context) CapabilityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityPropertiesResponseOutput)
}

func (i CapabilityPropertiesResponseArgs) ToCapabilityPropertiesResponsePtrOutput() CapabilityPropertiesResponsePtrOutput {
	return i.ToCapabilityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CapabilityPropertiesResponseArgs) ToCapabilityPropertiesResponsePtrOutputWithContext(ctx context.Context) CapabilityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityPropertiesResponseOutput).ToCapabilityPropertiesResponsePtrOutputWithContext(ctx)
}









type CapabilityPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCapabilityPropertiesResponsePtrOutput() CapabilityPropertiesResponsePtrOutput
	ToCapabilityPropertiesResponsePtrOutputWithContext(context.Context) CapabilityPropertiesResponsePtrOutput
}

type capabilityPropertiesResponsePtrType CapabilityPropertiesResponseArgs

func CapabilityPropertiesResponsePtr(v *CapabilityPropertiesResponseArgs) CapabilityPropertiesResponsePtrInput {
	return (*capabilityPropertiesResponsePtrType)(v)
}

func (*capabilityPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CapabilityPropertiesResponse)(nil)).Elem()
}

func (i *capabilityPropertiesResponsePtrType) ToCapabilityPropertiesResponsePtrOutput() CapabilityPropertiesResponsePtrOutput {
	return i.ToCapabilityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *capabilityPropertiesResponsePtrType) ToCapabilityPropertiesResponsePtrOutputWithContext(ctx context.Context) CapabilityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityPropertiesResponsePtrOutput)
}

type CapabilityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CapabilityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityPropertiesResponse)(nil)).Elem()
}

func (o CapabilityPropertiesResponseOutput) ToCapabilityPropertiesResponseOutput() CapabilityPropertiesResponseOutput {
	return o
}

func (o CapabilityPropertiesResponseOutput) ToCapabilityPropertiesResponseOutputWithContext(ctx context.Context) CapabilityPropertiesResponseOutput {
	return o
}

func (o CapabilityPropertiesResponseOutput) ToCapabilityPropertiesResponsePtrOutput() CapabilityPropertiesResponsePtrOutput {
	return o.ToCapabilityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CapabilityPropertiesResponseOutput) ToCapabilityPropertiesResponsePtrOutputWithContext(ctx context.Context) CapabilityPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CapabilityPropertiesResponse) *CapabilityPropertiesResponse {
		return &v
	}).(CapabilityPropertiesResponsePtrOutput)
}

func (o CapabilityPropertiesResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v CapabilityPropertiesResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o CapabilityPropertiesResponseOutput) ParametersSchema() pulumi.StringOutput {
	return o.ApplyT(func(v CapabilityPropertiesResponse) string { return v.ParametersSchema }).(pulumi.StringOutput)
}

func (o CapabilityPropertiesResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v CapabilityPropertiesResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o CapabilityPropertiesResponseOutput) TargetType() pulumi.StringOutput {
	return o.ApplyT(func(v CapabilityPropertiesResponse) string { return v.TargetType }).(pulumi.StringOutput)
}

func (o CapabilityPropertiesResponseOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v CapabilityPropertiesResponse) string { return v.Urn }).(pulumi.StringOutput)
}

type CapabilityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CapabilityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapabilityPropertiesResponse)(nil)).Elem()
}

func (o CapabilityPropertiesResponsePtrOutput) ToCapabilityPropertiesResponsePtrOutput() CapabilityPropertiesResponsePtrOutput {
	return o
}

func (o CapabilityPropertiesResponsePtrOutput) ToCapabilityPropertiesResponsePtrOutputWithContext(ctx context.Context) CapabilityPropertiesResponsePtrOutput {
	return o
}

func (o CapabilityPropertiesResponsePtrOutput) Elem() CapabilityPropertiesResponseOutput {
	return o.ApplyT(func(v *CapabilityPropertiesResponse) CapabilityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CapabilityPropertiesResponse
		return ret
	}).(CapabilityPropertiesResponseOutput)
}

func (o CapabilityPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapabilityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o CapabilityPropertiesResponsePtrOutput) ParametersSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapabilityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ParametersSchema
	}).(pulumi.StringPtrOutput)
}

func (o CapabilityPropertiesResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapabilityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o CapabilityPropertiesResponsePtrOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapabilityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetType
	}).(pulumi.StringPtrOutput)
}

func (o CapabilityPropertiesResponsePtrOutput) Urn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapabilityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Urn
	}).(pulumi.StringPtrOutput)
}

type ContinuousAction struct {
	Duration   string         `pulumi:"duration"`
	Name       string         `pulumi:"name"`
	Parameters []KeyValuePair `pulumi:"parameters"`
	SelectorId string         `pulumi:"selectorId"`
	Type       string         `pulumi:"type"`
}





type ContinuousActionInput interface {
	pulumi.Input

	ToContinuousActionOutput() ContinuousActionOutput
	ToContinuousActionOutputWithContext(context.Context) ContinuousActionOutput
}

type ContinuousActionArgs struct {
	Duration   pulumi.StringInput     `pulumi:"duration"`
	Name       pulumi.StringInput     `pulumi:"name"`
	Parameters KeyValuePairArrayInput `pulumi:"parameters"`
	SelectorId pulumi.StringInput     `pulumi:"selectorId"`
	Type       pulumi.StringInput     `pulumi:"type"`
}

func (ContinuousActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousAction)(nil)).Elem()
}

func (i ContinuousActionArgs) ToContinuousActionOutput() ContinuousActionOutput {
	return i.ToContinuousActionOutputWithContext(context.Background())
}

func (i ContinuousActionArgs) ToContinuousActionOutputWithContext(ctx context.Context) ContinuousActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContinuousActionOutput)
}

type ContinuousActionOutput struct{ *pulumi.OutputState }

func (ContinuousActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousAction)(nil)).Elem()
}

func (o ContinuousActionOutput) ToContinuousActionOutput() ContinuousActionOutput {
	return o
}

func (o ContinuousActionOutput) ToContinuousActionOutputWithContext(ctx context.Context) ContinuousActionOutput {
	return o
}

func (o ContinuousActionOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousAction) string { return v.Duration }).(pulumi.StringOutput)
}

func (o ContinuousActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContinuousActionOutput) Parameters() KeyValuePairArrayOutput {
	return o.ApplyT(func(v ContinuousAction) []KeyValuePair { return v.Parameters }).(KeyValuePairArrayOutput)
}

func (o ContinuousActionOutput) SelectorId() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousAction) string { return v.SelectorId }).(pulumi.StringOutput)
}

func (o ContinuousActionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousAction) string { return v.Type }).(pulumi.StringOutput)
}

type ContinuousActionResponse struct {
	Duration   string                 `pulumi:"duration"`
	Name       string                 `pulumi:"name"`
	Parameters []KeyValuePairResponse `pulumi:"parameters"`
	SelectorId string                 `pulumi:"selectorId"`
	Type       string                 `pulumi:"type"`
}





type ContinuousActionResponseInput interface {
	pulumi.Input

	ToContinuousActionResponseOutput() ContinuousActionResponseOutput
	ToContinuousActionResponseOutputWithContext(context.Context) ContinuousActionResponseOutput
}

type ContinuousActionResponseArgs struct {
	Duration   pulumi.StringInput             `pulumi:"duration"`
	Name       pulumi.StringInput             `pulumi:"name"`
	Parameters KeyValuePairResponseArrayInput `pulumi:"parameters"`
	SelectorId pulumi.StringInput             `pulumi:"selectorId"`
	Type       pulumi.StringInput             `pulumi:"type"`
}

func (ContinuousActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousActionResponse)(nil)).Elem()
}

func (i ContinuousActionResponseArgs) ToContinuousActionResponseOutput() ContinuousActionResponseOutput {
	return i.ToContinuousActionResponseOutputWithContext(context.Background())
}

func (i ContinuousActionResponseArgs) ToContinuousActionResponseOutputWithContext(ctx context.Context) ContinuousActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContinuousActionResponseOutput)
}

type ContinuousActionResponseOutput struct{ *pulumi.OutputState }

func (ContinuousActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousActionResponse)(nil)).Elem()
}

func (o ContinuousActionResponseOutput) ToContinuousActionResponseOutput() ContinuousActionResponseOutput {
	return o
}

func (o ContinuousActionResponseOutput) ToContinuousActionResponseOutputWithContext(ctx context.Context) ContinuousActionResponseOutput {
	return o
}

func (o ContinuousActionResponseOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousActionResponse) string { return v.Duration }).(pulumi.StringOutput)
}

func (o ContinuousActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContinuousActionResponseOutput) Parameters() KeyValuePairResponseArrayOutput {
	return o.ApplyT(func(v ContinuousActionResponse) []KeyValuePairResponse { return v.Parameters }).(KeyValuePairResponseArrayOutput)
}

func (o ContinuousActionResponseOutput) SelectorId() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousActionResponse) string { return v.SelectorId }).(pulumi.StringOutput)
}

func (o ContinuousActionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousActionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type DelayAction struct {
	Duration string `pulumi:"duration"`
	Name     string `pulumi:"name"`
	Type     string `pulumi:"type"`
}





type DelayActionInput interface {
	pulumi.Input

	ToDelayActionOutput() DelayActionOutput
	ToDelayActionOutputWithContext(context.Context) DelayActionOutput
}

type DelayActionArgs struct {
	Duration pulumi.StringInput `pulumi:"duration"`
	Name     pulumi.StringInput `pulumi:"name"`
	Type     pulumi.StringInput `pulumi:"type"`
}

func (DelayActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DelayAction)(nil)).Elem()
}

func (i DelayActionArgs) ToDelayActionOutput() DelayActionOutput {
	return i.ToDelayActionOutputWithContext(context.Background())
}

func (i DelayActionArgs) ToDelayActionOutputWithContext(ctx context.Context) DelayActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelayActionOutput)
}

type DelayActionOutput struct{ *pulumi.OutputState }

func (DelayActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DelayAction)(nil)).Elem()
}

func (o DelayActionOutput) ToDelayActionOutput() DelayActionOutput {
	return o
}

func (o DelayActionOutput) ToDelayActionOutputWithContext(ctx context.Context) DelayActionOutput {
	return o
}

func (o DelayActionOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v DelayAction) string { return v.Duration }).(pulumi.StringOutput)
}

func (o DelayActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DelayAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o DelayActionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DelayAction) string { return v.Type }).(pulumi.StringOutput)
}

type DelayActionResponse struct {
	Duration string `pulumi:"duration"`
	Name     string `pulumi:"name"`
	Type     string `pulumi:"type"`
}





type DelayActionResponseInput interface {
	pulumi.Input

	ToDelayActionResponseOutput() DelayActionResponseOutput
	ToDelayActionResponseOutputWithContext(context.Context) DelayActionResponseOutput
}

type DelayActionResponseArgs struct {
	Duration pulumi.StringInput `pulumi:"duration"`
	Name     pulumi.StringInput `pulumi:"name"`
	Type     pulumi.StringInput `pulumi:"type"`
}

func (DelayActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DelayActionResponse)(nil)).Elem()
}

func (i DelayActionResponseArgs) ToDelayActionResponseOutput() DelayActionResponseOutput {
	return i.ToDelayActionResponseOutputWithContext(context.Background())
}

func (i DelayActionResponseArgs) ToDelayActionResponseOutputWithContext(ctx context.Context) DelayActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelayActionResponseOutput)
}

type DelayActionResponseOutput struct{ *pulumi.OutputState }

func (DelayActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DelayActionResponse)(nil)).Elem()
}

func (o DelayActionResponseOutput) ToDelayActionResponseOutput() DelayActionResponseOutput {
	return o
}

func (o DelayActionResponseOutput) ToDelayActionResponseOutputWithContext(ctx context.Context) DelayActionResponseOutput {
	return o
}

func (o DelayActionResponseOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v DelayActionResponse) string { return v.Duration }).(pulumi.StringOutput)
}

func (o DelayActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DelayActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DelayActionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DelayActionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type DiscreteAction struct {
	Name       string         `pulumi:"name"`
	Parameters []KeyValuePair `pulumi:"parameters"`
	SelectorId string         `pulumi:"selectorId"`
	Type       string         `pulumi:"type"`
}





type DiscreteActionInput interface {
	pulumi.Input

	ToDiscreteActionOutput() DiscreteActionOutput
	ToDiscreteActionOutputWithContext(context.Context) DiscreteActionOutput
}

type DiscreteActionArgs struct {
	Name       pulumi.StringInput     `pulumi:"name"`
	Parameters KeyValuePairArrayInput `pulumi:"parameters"`
	SelectorId pulumi.StringInput     `pulumi:"selectorId"`
	Type       pulumi.StringInput     `pulumi:"type"`
}

func (DiscreteActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiscreteAction)(nil)).Elem()
}

func (i DiscreteActionArgs) ToDiscreteActionOutput() DiscreteActionOutput {
	return i.ToDiscreteActionOutputWithContext(context.Background())
}

func (i DiscreteActionArgs) ToDiscreteActionOutputWithContext(ctx context.Context) DiscreteActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiscreteActionOutput)
}

type DiscreteActionOutput struct{ *pulumi.OutputState }

func (DiscreteActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiscreteAction)(nil)).Elem()
}

func (o DiscreteActionOutput) ToDiscreteActionOutput() DiscreteActionOutput {
	return o
}

func (o DiscreteActionOutput) ToDiscreteActionOutputWithContext(ctx context.Context) DiscreteActionOutput {
	return o
}

func (o DiscreteActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DiscreteAction) string { return v.Name }).(pulumi.StringOutput)
}

func (o DiscreteActionOutput) Parameters() KeyValuePairArrayOutput {
	return o.ApplyT(func(v DiscreteAction) []KeyValuePair { return v.Parameters }).(KeyValuePairArrayOutput)
}

func (o DiscreteActionOutput) SelectorId() pulumi.StringOutput {
	return o.ApplyT(func(v DiscreteAction) string { return v.SelectorId }).(pulumi.StringOutput)
}

func (o DiscreteActionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DiscreteAction) string { return v.Type }).(pulumi.StringOutput)
}

type DiscreteActionResponse struct {
	Name       string                 `pulumi:"name"`
	Parameters []KeyValuePairResponse `pulumi:"parameters"`
	SelectorId string                 `pulumi:"selectorId"`
	Type       string                 `pulumi:"type"`
}





type DiscreteActionResponseInput interface {
	pulumi.Input

	ToDiscreteActionResponseOutput() DiscreteActionResponseOutput
	ToDiscreteActionResponseOutputWithContext(context.Context) DiscreteActionResponseOutput
}

type DiscreteActionResponseArgs struct {
	Name       pulumi.StringInput             `pulumi:"name"`
	Parameters KeyValuePairResponseArrayInput `pulumi:"parameters"`
	SelectorId pulumi.StringInput             `pulumi:"selectorId"`
	Type       pulumi.StringInput             `pulumi:"type"`
}

func (DiscreteActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiscreteActionResponse)(nil)).Elem()
}

func (i DiscreteActionResponseArgs) ToDiscreteActionResponseOutput() DiscreteActionResponseOutput {
	return i.ToDiscreteActionResponseOutputWithContext(context.Background())
}

func (i DiscreteActionResponseArgs) ToDiscreteActionResponseOutputWithContext(ctx context.Context) DiscreteActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiscreteActionResponseOutput)
}

type DiscreteActionResponseOutput struct{ *pulumi.OutputState }

func (DiscreteActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiscreteActionResponse)(nil)).Elem()
}

func (o DiscreteActionResponseOutput) ToDiscreteActionResponseOutput() DiscreteActionResponseOutput {
	return o
}

func (o DiscreteActionResponseOutput) ToDiscreteActionResponseOutputWithContext(ctx context.Context) DiscreteActionResponseOutput {
	return o
}

func (o DiscreteActionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DiscreteActionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DiscreteActionResponseOutput) Parameters() KeyValuePairResponseArrayOutput {
	return o.ApplyT(func(v DiscreteActionResponse) []KeyValuePairResponse { return v.Parameters }).(KeyValuePairResponseArrayOutput)
}

func (o DiscreteActionResponseOutput) SelectorId() pulumi.StringOutput {
	return o.ApplyT(func(v DiscreteActionResponse) string { return v.SelectorId }).(pulumi.StringOutput)
}

func (o DiscreteActionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DiscreteActionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ExperimentProperties struct {
	Selectors       []Selector `pulumi:"selectors"`
	StartOnCreation *bool      `pulumi:"startOnCreation"`
	Steps           []Step     `pulumi:"steps"`
}





type ExperimentPropertiesInput interface {
	pulumi.Input

	ToExperimentPropertiesOutput() ExperimentPropertiesOutput
	ToExperimentPropertiesOutputWithContext(context.Context) ExperimentPropertiesOutput
}

type ExperimentPropertiesArgs struct {
	Selectors       SelectorArrayInput  `pulumi:"selectors"`
	StartOnCreation pulumi.BoolPtrInput `pulumi:"startOnCreation"`
	Steps           StepArrayInput      `pulumi:"steps"`
}

func (ExperimentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExperimentProperties)(nil)).Elem()
}

func (i ExperimentPropertiesArgs) ToExperimentPropertiesOutput() ExperimentPropertiesOutput {
	return i.ToExperimentPropertiesOutputWithContext(context.Background())
}

func (i ExperimentPropertiesArgs) ToExperimentPropertiesOutputWithContext(ctx context.Context) ExperimentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentPropertiesOutput)
}

func (i ExperimentPropertiesArgs) ToExperimentPropertiesPtrOutput() ExperimentPropertiesPtrOutput {
	return i.ToExperimentPropertiesPtrOutputWithContext(context.Background())
}

func (i ExperimentPropertiesArgs) ToExperimentPropertiesPtrOutputWithContext(ctx context.Context) ExperimentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentPropertiesOutput).ToExperimentPropertiesPtrOutputWithContext(ctx)
}









type ExperimentPropertiesPtrInput interface {
	pulumi.Input

	ToExperimentPropertiesPtrOutput() ExperimentPropertiesPtrOutput
	ToExperimentPropertiesPtrOutputWithContext(context.Context) ExperimentPropertiesPtrOutput
}

type experimentPropertiesPtrType ExperimentPropertiesArgs

func ExperimentPropertiesPtr(v *ExperimentPropertiesArgs) ExperimentPropertiesPtrInput {
	return (*experimentPropertiesPtrType)(v)
}

func (*experimentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentProperties)(nil)).Elem()
}

func (i *experimentPropertiesPtrType) ToExperimentPropertiesPtrOutput() ExperimentPropertiesPtrOutput {
	return i.ToExperimentPropertiesPtrOutputWithContext(context.Background())
}

func (i *experimentPropertiesPtrType) ToExperimentPropertiesPtrOutputWithContext(ctx context.Context) ExperimentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentPropertiesPtrOutput)
}

type ExperimentPropertiesOutput struct{ *pulumi.OutputState }

func (ExperimentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExperimentProperties)(nil)).Elem()
}

func (o ExperimentPropertiesOutput) ToExperimentPropertiesOutput() ExperimentPropertiesOutput {
	return o
}

func (o ExperimentPropertiesOutput) ToExperimentPropertiesOutputWithContext(ctx context.Context) ExperimentPropertiesOutput {
	return o
}

func (o ExperimentPropertiesOutput) ToExperimentPropertiesPtrOutput() ExperimentPropertiesPtrOutput {
	return o.ToExperimentPropertiesPtrOutputWithContext(context.Background())
}

func (o ExperimentPropertiesOutput) ToExperimentPropertiesPtrOutputWithContext(ctx context.Context) ExperimentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExperimentProperties) *ExperimentProperties {
		return &v
	}).(ExperimentPropertiesPtrOutput)
}

func (o ExperimentPropertiesOutput) Selectors() SelectorArrayOutput {
	return o.ApplyT(func(v ExperimentProperties) []Selector { return v.Selectors }).(SelectorArrayOutput)
}

func (o ExperimentPropertiesOutput) StartOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExperimentProperties) *bool { return v.StartOnCreation }).(pulumi.BoolPtrOutput)
}

func (o ExperimentPropertiesOutput) Steps() StepArrayOutput {
	return o.ApplyT(func(v ExperimentProperties) []Step { return v.Steps }).(StepArrayOutput)
}

type ExperimentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ExperimentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentProperties)(nil)).Elem()
}

func (o ExperimentPropertiesPtrOutput) ToExperimentPropertiesPtrOutput() ExperimentPropertiesPtrOutput {
	return o
}

func (o ExperimentPropertiesPtrOutput) ToExperimentPropertiesPtrOutputWithContext(ctx context.Context) ExperimentPropertiesPtrOutput {
	return o
}

func (o ExperimentPropertiesPtrOutput) Elem() ExperimentPropertiesOutput {
	return o.ApplyT(func(v *ExperimentProperties) ExperimentProperties {
		if v != nil {
			return *v
		}
		var ret ExperimentProperties
		return ret
	}).(ExperimentPropertiesOutput)
}

func (o ExperimentPropertiesPtrOutput) Selectors() SelectorArrayOutput {
	return o.ApplyT(func(v *ExperimentProperties) []Selector {
		if v == nil {
			return nil
		}
		return v.Selectors
	}).(SelectorArrayOutput)
}

func (o ExperimentPropertiesPtrOutput) StartOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExperimentProperties) *bool {
		if v == nil {
			return nil
		}
		return v.StartOnCreation
	}).(pulumi.BoolPtrOutput)
}

func (o ExperimentPropertiesPtrOutput) Steps() StepArrayOutput {
	return o.ApplyT(func(v *ExperimentProperties) []Step {
		if v == nil {
			return nil
		}
		return v.Steps
	}).(StepArrayOutput)
}

type ExperimentPropertiesResponse struct {
	Selectors       []SelectorResponse `pulumi:"selectors"`
	StartOnCreation *bool              `pulumi:"startOnCreation"`
	Steps           []StepResponse     `pulumi:"steps"`
}





type ExperimentPropertiesResponseInput interface {
	pulumi.Input

	ToExperimentPropertiesResponseOutput() ExperimentPropertiesResponseOutput
	ToExperimentPropertiesResponseOutputWithContext(context.Context) ExperimentPropertiesResponseOutput
}

type ExperimentPropertiesResponseArgs struct {
	Selectors       SelectorResponseArrayInput `pulumi:"selectors"`
	StartOnCreation pulumi.BoolPtrInput        `pulumi:"startOnCreation"`
	Steps           StepResponseArrayInput     `pulumi:"steps"`
}

func (ExperimentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExperimentPropertiesResponse)(nil)).Elem()
}

func (i ExperimentPropertiesResponseArgs) ToExperimentPropertiesResponseOutput() ExperimentPropertiesResponseOutput {
	return i.ToExperimentPropertiesResponseOutputWithContext(context.Background())
}

func (i ExperimentPropertiesResponseArgs) ToExperimentPropertiesResponseOutputWithContext(ctx context.Context) ExperimentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentPropertiesResponseOutput)
}

func (i ExperimentPropertiesResponseArgs) ToExperimentPropertiesResponsePtrOutput() ExperimentPropertiesResponsePtrOutput {
	return i.ToExperimentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ExperimentPropertiesResponseArgs) ToExperimentPropertiesResponsePtrOutputWithContext(ctx context.Context) ExperimentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentPropertiesResponseOutput).ToExperimentPropertiesResponsePtrOutputWithContext(ctx)
}









type ExperimentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToExperimentPropertiesResponsePtrOutput() ExperimentPropertiesResponsePtrOutput
	ToExperimentPropertiesResponsePtrOutputWithContext(context.Context) ExperimentPropertiesResponsePtrOutput
}

type experimentPropertiesResponsePtrType ExperimentPropertiesResponseArgs

func ExperimentPropertiesResponsePtr(v *ExperimentPropertiesResponseArgs) ExperimentPropertiesResponsePtrInput {
	return (*experimentPropertiesResponsePtrType)(v)
}

func (*experimentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentPropertiesResponse)(nil)).Elem()
}

func (i *experimentPropertiesResponsePtrType) ToExperimentPropertiesResponsePtrOutput() ExperimentPropertiesResponsePtrOutput {
	return i.ToExperimentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *experimentPropertiesResponsePtrType) ToExperimentPropertiesResponsePtrOutputWithContext(ctx context.Context) ExperimentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentPropertiesResponsePtrOutput)
}

type ExperimentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ExperimentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExperimentPropertiesResponse)(nil)).Elem()
}

func (o ExperimentPropertiesResponseOutput) ToExperimentPropertiesResponseOutput() ExperimentPropertiesResponseOutput {
	return o
}

func (o ExperimentPropertiesResponseOutput) ToExperimentPropertiesResponseOutputWithContext(ctx context.Context) ExperimentPropertiesResponseOutput {
	return o
}

func (o ExperimentPropertiesResponseOutput) ToExperimentPropertiesResponsePtrOutput() ExperimentPropertiesResponsePtrOutput {
	return o.ToExperimentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ExperimentPropertiesResponseOutput) ToExperimentPropertiesResponsePtrOutputWithContext(ctx context.Context) ExperimentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExperimentPropertiesResponse) *ExperimentPropertiesResponse {
		return &v
	}).(ExperimentPropertiesResponsePtrOutput)
}

func (o ExperimentPropertiesResponseOutput) Selectors() SelectorResponseArrayOutput {
	return o.ApplyT(func(v ExperimentPropertiesResponse) []SelectorResponse { return v.Selectors }).(SelectorResponseArrayOutput)
}

func (o ExperimentPropertiesResponseOutput) StartOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExperimentPropertiesResponse) *bool { return v.StartOnCreation }).(pulumi.BoolPtrOutput)
}

func (o ExperimentPropertiesResponseOutput) Steps() StepResponseArrayOutput {
	return o.ApplyT(func(v ExperimentPropertiesResponse) []StepResponse { return v.Steps }).(StepResponseArrayOutput)
}

type ExperimentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ExperimentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentPropertiesResponse)(nil)).Elem()
}

func (o ExperimentPropertiesResponsePtrOutput) ToExperimentPropertiesResponsePtrOutput() ExperimentPropertiesResponsePtrOutput {
	return o
}

func (o ExperimentPropertiesResponsePtrOutput) ToExperimentPropertiesResponsePtrOutputWithContext(ctx context.Context) ExperimentPropertiesResponsePtrOutput {
	return o
}

func (o ExperimentPropertiesResponsePtrOutput) Elem() ExperimentPropertiesResponseOutput {
	return o.ApplyT(func(v *ExperimentPropertiesResponse) ExperimentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ExperimentPropertiesResponse
		return ret
	}).(ExperimentPropertiesResponseOutput)
}

func (o ExperimentPropertiesResponsePtrOutput) Selectors() SelectorResponseArrayOutput {
	return o.ApplyT(func(v *ExperimentPropertiesResponse) []SelectorResponse {
		if v == nil {
			return nil
		}
		return v.Selectors
	}).(SelectorResponseArrayOutput)
}

func (o ExperimentPropertiesResponsePtrOutput) StartOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExperimentPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.StartOnCreation
	}).(pulumi.BoolPtrOutput)
}

func (o ExperimentPropertiesResponsePtrOutput) Steps() StepResponseArrayOutput {
	return o.ApplyT(func(v *ExperimentPropertiesResponse) []StepResponse {
		if v == nil {
			return nil
		}
		return v.Steps
	}).(StepResponseArrayOutput)
}

type KeyValuePair struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}





type KeyValuePairInput interface {
	pulumi.Input

	ToKeyValuePairOutput() KeyValuePairOutput
	ToKeyValuePairOutputWithContext(context.Context) KeyValuePairOutput
}

type KeyValuePairArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (KeyValuePairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyValuePair)(nil)).Elem()
}

func (i KeyValuePairArgs) ToKeyValuePairOutput() KeyValuePairOutput {
	return i.ToKeyValuePairOutputWithContext(context.Background())
}

func (i KeyValuePairArgs) ToKeyValuePairOutputWithContext(ctx context.Context) KeyValuePairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyValuePairOutput)
}





type KeyValuePairArrayInput interface {
	pulumi.Input

	ToKeyValuePairArrayOutput() KeyValuePairArrayOutput
	ToKeyValuePairArrayOutputWithContext(context.Context) KeyValuePairArrayOutput
}

type KeyValuePairArray []KeyValuePairInput

func (KeyValuePairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyValuePair)(nil)).Elem()
}

func (i KeyValuePairArray) ToKeyValuePairArrayOutput() KeyValuePairArrayOutput {
	return i.ToKeyValuePairArrayOutputWithContext(context.Background())
}

func (i KeyValuePairArray) ToKeyValuePairArrayOutputWithContext(ctx context.Context) KeyValuePairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyValuePairArrayOutput)
}

type KeyValuePairOutput struct{ *pulumi.OutputState }

func (KeyValuePairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyValuePair)(nil)).Elem()
}

func (o KeyValuePairOutput) ToKeyValuePairOutput() KeyValuePairOutput {
	return o
}

func (o KeyValuePairOutput) ToKeyValuePairOutputWithContext(ctx context.Context) KeyValuePairOutput {
	return o
}

func (o KeyValuePairOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v KeyValuePair) string { return v.Key }).(pulumi.StringOutput)
}

func (o KeyValuePairOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v KeyValuePair) string { return v.Value }).(pulumi.StringOutput)
}

type KeyValuePairArrayOutput struct{ *pulumi.OutputState }

func (KeyValuePairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyValuePair)(nil)).Elem()
}

func (o KeyValuePairArrayOutput) ToKeyValuePairArrayOutput() KeyValuePairArrayOutput {
	return o
}

func (o KeyValuePairArrayOutput) ToKeyValuePairArrayOutputWithContext(ctx context.Context) KeyValuePairArrayOutput {
	return o
}

func (o KeyValuePairArrayOutput) Index(i pulumi.IntInput) KeyValuePairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyValuePair {
		return vs[0].([]KeyValuePair)[vs[1].(int)]
	}).(KeyValuePairOutput)
}

type KeyValuePairResponse struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}





type KeyValuePairResponseInput interface {
	pulumi.Input

	ToKeyValuePairResponseOutput() KeyValuePairResponseOutput
	ToKeyValuePairResponseOutputWithContext(context.Context) KeyValuePairResponseOutput
}

type KeyValuePairResponseArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (KeyValuePairResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyValuePairResponse)(nil)).Elem()
}

func (i KeyValuePairResponseArgs) ToKeyValuePairResponseOutput() KeyValuePairResponseOutput {
	return i.ToKeyValuePairResponseOutputWithContext(context.Background())
}

func (i KeyValuePairResponseArgs) ToKeyValuePairResponseOutputWithContext(ctx context.Context) KeyValuePairResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyValuePairResponseOutput)
}





type KeyValuePairResponseArrayInput interface {
	pulumi.Input

	ToKeyValuePairResponseArrayOutput() KeyValuePairResponseArrayOutput
	ToKeyValuePairResponseArrayOutputWithContext(context.Context) KeyValuePairResponseArrayOutput
}

type KeyValuePairResponseArray []KeyValuePairResponseInput

func (KeyValuePairResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyValuePairResponse)(nil)).Elem()
}

func (i KeyValuePairResponseArray) ToKeyValuePairResponseArrayOutput() KeyValuePairResponseArrayOutput {
	return i.ToKeyValuePairResponseArrayOutputWithContext(context.Background())
}

func (i KeyValuePairResponseArray) ToKeyValuePairResponseArrayOutputWithContext(ctx context.Context) KeyValuePairResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyValuePairResponseArrayOutput)
}

type KeyValuePairResponseOutput struct{ *pulumi.OutputState }

func (KeyValuePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyValuePairResponse)(nil)).Elem()
}

func (o KeyValuePairResponseOutput) ToKeyValuePairResponseOutput() KeyValuePairResponseOutput {
	return o
}

func (o KeyValuePairResponseOutput) ToKeyValuePairResponseOutputWithContext(ctx context.Context) KeyValuePairResponseOutput {
	return o
}

func (o KeyValuePairResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v KeyValuePairResponse) string { return v.Key }).(pulumi.StringOutput)
}

func (o KeyValuePairResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v KeyValuePairResponse) string { return v.Value }).(pulumi.StringOutput)
}

type KeyValuePairResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyValuePairResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyValuePairResponse)(nil)).Elem()
}

func (o KeyValuePairResponseArrayOutput) ToKeyValuePairResponseArrayOutput() KeyValuePairResponseArrayOutput {
	return o
}

func (o KeyValuePairResponseArrayOutput) ToKeyValuePairResponseArrayOutputWithContext(ctx context.Context) KeyValuePairResponseArrayOutput {
	return o
}

func (o KeyValuePairResponseArrayOutput) Index(i pulumi.IntInput) KeyValuePairResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyValuePairResponse {
		return vs[0].([]KeyValuePairResponse)[vs[1].(int)]
	}).(KeyValuePairResponseOutput)
}

type ResourceIdentity struct {
	Type ResourceIdentityType `pulumi:"type"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type ResourceIdentityTypeInput `pulumi:"type"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v ResourceIdentity) ResourceIdentityType { return v.Type }).(ResourceIdentityTypeOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}





type ResourceIdentityResponseInput interface {
	pulumi.Input

	ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput
	ToResourceIdentityResponseOutputWithContext(context.Context) ResourceIdentityResponseOutput
}

type ResourceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (ResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return i.ToResourceIdentityResponseOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput)
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput).ToResourceIdentityResponsePtrOutputWithContext(ctx)
}









type ResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput
	ToResourceIdentityResponsePtrOutputWithContext(context.Context) ResourceIdentityResponsePtrOutput
}

type resourceIdentityResponsePtrType ResourceIdentityResponseArgs

func ResourceIdentityResponsePtr(v *ResourceIdentityResponseArgs) ResourceIdentityResponsePtrInput {
	return (*resourceIdentityResponsePtrType)(v)
}

func (*resourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponsePtrOutput)
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityResponse) *ResourceIdentityResponse {
		return &v
	}).(ResourceIdentityResponsePtrOutput)
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type Selector struct {
	Id      string            `pulumi:"id"`
	Targets []TargetReference `pulumi:"targets"`
	Type    SelectorType      `pulumi:"type"`
}





type SelectorInput interface {
	pulumi.Input

	ToSelectorOutput() SelectorOutput
	ToSelectorOutputWithContext(context.Context) SelectorOutput
}

type SelectorArgs struct {
	Id      pulumi.StringInput        `pulumi:"id"`
	Targets TargetReferenceArrayInput `pulumi:"targets"`
	Type    SelectorTypeInput         `pulumi:"type"`
}

func (SelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Selector)(nil)).Elem()
}

func (i SelectorArgs) ToSelectorOutput() SelectorOutput {
	return i.ToSelectorOutputWithContext(context.Background())
}

func (i SelectorArgs) ToSelectorOutputWithContext(ctx context.Context) SelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorOutput)
}





type SelectorArrayInput interface {
	pulumi.Input

	ToSelectorArrayOutput() SelectorArrayOutput
	ToSelectorArrayOutputWithContext(context.Context) SelectorArrayOutput
}

type SelectorArray []SelectorInput

func (SelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Selector)(nil)).Elem()
}

func (i SelectorArray) ToSelectorArrayOutput() SelectorArrayOutput {
	return i.ToSelectorArrayOutputWithContext(context.Background())
}

func (i SelectorArray) ToSelectorArrayOutputWithContext(ctx context.Context) SelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorArrayOutput)
}

type SelectorOutput struct{ *pulumi.OutputState }

func (SelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Selector)(nil)).Elem()
}

func (o SelectorOutput) ToSelectorOutput() SelectorOutput {
	return o
}

func (o SelectorOutput) ToSelectorOutputWithContext(ctx context.Context) SelectorOutput {
	return o
}

func (o SelectorOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Selector) string { return v.Id }).(pulumi.StringOutput)
}

func (o SelectorOutput) Targets() TargetReferenceArrayOutput {
	return o.ApplyT(func(v Selector) []TargetReference { return v.Targets }).(TargetReferenceArrayOutput)
}

func (o SelectorOutput) Type() SelectorTypeOutput {
	return o.ApplyT(func(v Selector) SelectorType { return v.Type }).(SelectorTypeOutput)
}

type SelectorArrayOutput struct{ *pulumi.OutputState }

func (SelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Selector)(nil)).Elem()
}

func (o SelectorArrayOutput) ToSelectorArrayOutput() SelectorArrayOutput {
	return o
}

func (o SelectorArrayOutput) ToSelectorArrayOutputWithContext(ctx context.Context) SelectorArrayOutput {
	return o
}

func (o SelectorArrayOutput) Index(i pulumi.IntInput) SelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Selector {
		return vs[0].([]Selector)[vs[1].(int)]
	}).(SelectorOutput)
}

type SelectorResponse struct {
	Id      string                    `pulumi:"id"`
	Targets []TargetReferenceResponse `pulumi:"targets"`
	Type    string                    `pulumi:"type"`
}





type SelectorResponseInput interface {
	pulumi.Input

	ToSelectorResponseOutput() SelectorResponseOutput
	ToSelectorResponseOutputWithContext(context.Context) SelectorResponseOutput
}

type SelectorResponseArgs struct {
	Id      pulumi.StringInput                `pulumi:"id"`
	Targets TargetReferenceResponseArrayInput `pulumi:"targets"`
	Type    pulumi.StringInput                `pulumi:"type"`
}

func (SelectorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectorResponse)(nil)).Elem()
}

func (i SelectorResponseArgs) ToSelectorResponseOutput() SelectorResponseOutput {
	return i.ToSelectorResponseOutputWithContext(context.Background())
}

func (i SelectorResponseArgs) ToSelectorResponseOutputWithContext(ctx context.Context) SelectorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorResponseOutput)
}





type SelectorResponseArrayInput interface {
	pulumi.Input

	ToSelectorResponseArrayOutput() SelectorResponseArrayOutput
	ToSelectorResponseArrayOutputWithContext(context.Context) SelectorResponseArrayOutput
}

type SelectorResponseArray []SelectorResponseInput

func (SelectorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelectorResponse)(nil)).Elem()
}

func (i SelectorResponseArray) ToSelectorResponseArrayOutput() SelectorResponseArrayOutput {
	return i.ToSelectorResponseArrayOutputWithContext(context.Background())
}

func (i SelectorResponseArray) ToSelectorResponseArrayOutputWithContext(ctx context.Context) SelectorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorResponseArrayOutput)
}

type SelectorResponseOutput struct{ *pulumi.OutputState }

func (SelectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectorResponse)(nil)).Elem()
}

func (o SelectorResponseOutput) ToSelectorResponseOutput() SelectorResponseOutput {
	return o
}

func (o SelectorResponseOutput) ToSelectorResponseOutputWithContext(ctx context.Context) SelectorResponseOutput {
	return o
}

func (o SelectorResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SelectorResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o SelectorResponseOutput) Targets() TargetReferenceResponseArrayOutput {
	return o.ApplyT(func(v SelectorResponse) []TargetReferenceResponse { return v.Targets }).(TargetReferenceResponseArrayOutput)
}

func (o SelectorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SelectorResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SelectorResponseArrayOutput struct{ *pulumi.OutputState }

func (SelectorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelectorResponse)(nil)).Elem()
}

func (o SelectorResponseArrayOutput) ToSelectorResponseArrayOutput() SelectorResponseArrayOutput {
	return o
}

func (o SelectorResponseArrayOutput) ToSelectorResponseArrayOutputWithContext(ctx context.Context) SelectorResponseArrayOutput {
	return o
}

func (o SelectorResponseArrayOutput) Index(i pulumi.IntInput) SelectorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SelectorResponse {
		return vs[0].([]SelectorResponse)[vs[1].(int)]
	}).(SelectorResponseOutput)
}

type Step struct {
	Branches []Branch `pulumi:"branches"`
	Name     string   `pulumi:"name"`
}





type StepInput interface {
	pulumi.Input

	ToStepOutput() StepOutput
	ToStepOutputWithContext(context.Context) StepOutput
}

type StepArgs struct {
	Branches BranchArrayInput   `pulumi:"branches"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (StepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Step)(nil)).Elem()
}

func (i StepArgs) ToStepOutput() StepOutput {
	return i.ToStepOutputWithContext(context.Background())
}

func (i StepArgs) ToStepOutputWithContext(ctx context.Context) StepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepOutput)
}





type StepArrayInput interface {
	pulumi.Input

	ToStepArrayOutput() StepArrayOutput
	ToStepArrayOutputWithContext(context.Context) StepArrayOutput
}

type StepArray []StepInput

func (StepArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Step)(nil)).Elem()
}

func (i StepArray) ToStepArrayOutput() StepArrayOutput {
	return i.ToStepArrayOutputWithContext(context.Background())
}

func (i StepArray) ToStepArrayOutputWithContext(ctx context.Context) StepArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepArrayOutput)
}

type StepOutput struct{ *pulumi.OutputState }

func (StepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Step)(nil)).Elem()
}

func (o StepOutput) ToStepOutput() StepOutput {
	return o
}

func (o StepOutput) ToStepOutputWithContext(ctx context.Context) StepOutput {
	return o
}

func (o StepOutput) Branches() BranchArrayOutput {
	return o.ApplyT(func(v Step) []Branch { return v.Branches }).(BranchArrayOutput)
}

func (o StepOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Step) string { return v.Name }).(pulumi.StringOutput)
}

type StepArrayOutput struct{ *pulumi.OutputState }

func (StepArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Step)(nil)).Elem()
}

func (o StepArrayOutput) ToStepArrayOutput() StepArrayOutput {
	return o
}

func (o StepArrayOutput) ToStepArrayOutputWithContext(ctx context.Context) StepArrayOutput {
	return o
}

func (o StepArrayOutput) Index(i pulumi.IntInput) StepOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Step {
		return vs[0].([]Step)[vs[1].(int)]
	}).(StepOutput)
}

type StepResponse struct {
	Branches []BranchResponse `pulumi:"branches"`
	Name     string           `pulumi:"name"`
}





type StepResponseInput interface {
	pulumi.Input

	ToStepResponseOutput() StepResponseOutput
	ToStepResponseOutputWithContext(context.Context) StepResponseOutput
}

type StepResponseArgs struct {
	Branches BranchResponseArrayInput `pulumi:"branches"`
	Name     pulumi.StringInput       `pulumi:"name"`
}

func (StepResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StepResponse)(nil)).Elem()
}

func (i StepResponseArgs) ToStepResponseOutput() StepResponseOutput {
	return i.ToStepResponseOutputWithContext(context.Background())
}

func (i StepResponseArgs) ToStepResponseOutputWithContext(ctx context.Context) StepResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepResponseOutput)
}





type StepResponseArrayInput interface {
	pulumi.Input

	ToStepResponseArrayOutput() StepResponseArrayOutput
	ToStepResponseArrayOutputWithContext(context.Context) StepResponseArrayOutput
}

type StepResponseArray []StepResponseInput

func (StepResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepResponse)(nil)).Elem()
}

func (i StepResponseArray) ToStepResponseArrayOutput() StepResponseArrayOutput {
	return i.ToStepResponseArrayOutputWithContext(context.Background())
}

func (i StepResponseArray) ToStepResponseArrayOutputWithContext(ctx context.Context) StepResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StepResponseArrayOutput)
}

type StepResponseOutput struct{ *pulumi.OutputState }

func (StepResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StepResponse)(nil)).Elem()
}

func (o StepResponseOutput) ToStepResponseOutput() StepResponseOutput {
	return o
}

func (o StepResponseOutput) ToStepResponseOutputWithContext(ctx context.Context) StepResponseOutput {
	return o
}

func (o StepResponseOutput) Branches() BranchResponseArrayOutput {
	return o.ApplyT(func(v StepResponse) []BranchResponse { return v.Branches }).(BranchResponseArrayOutput)
}

func (o StepResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StepResponse) string { return v.Name }).(pulumi.StringOutput)
}

type StepResponseArrayOutput struct{ *pulumi.OutputState }

func (StepResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StepResponse)(nil)).Elem()
}

func (o StepResponseArrayOutput) ToStepResponseArrayOutput() StepResponseArrayOutput {
	return o
}

func (o StepResponseArrayOutput) ToStepResponseArrayOutputWithContext(ctx context.Context) StepResponseArrayOutput {
	return o
}

func (o StepResponseArrayOutput) Index(i pulumi.IntInput) StepResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StepResponse {
		return vs[0].([]StepResponse)[vs[1].(int)]
	}).(StepResponseOutput)
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

type TargetReference struct {
	Id   string              `pulumi:"id"`
	Type TargetReferenceType `pulumi:"type"`
}





type TargetReferenceInput interface {
	pulumi.Input

	ToTargetReferenceOutput() TargetReferenceOutput
	ToTargetReferenceOutputWithContext(context.Context) TargetReferenceOutput
}

type TargetReferenceArgs struct {
	Id   pulumi.StringInput       `pulumi:"id"`
	Type TargetReferenceTypeInput `pulumi:"type"`
}

func (TargetReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetReference)(nil)).Elem()
}

func (i TargetReferenceArgs) ToTargetReferenceOutput() TargetReferenceOutput {
	return i.ToTargetReferenceOutputWithContext(context.Background())
}

func (i TargetReferenceArgs) ToTargetReferenceOutputWithContext(ctx context.Context) TargetReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetReferenceOutput)
}





type TargetReferenceArrayInput interface {
	pulumi.Input

	ToTargetReferenceArrayOutput() TargetReferenceArrayOutput
	ToTargetReferenceArrayOutputWithContext(context.Context) TargetReferenceArrayOutput
}

type TargetReferenceArray []TargetReferenceInput

func (TargetReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetReference)(nil)).Elem()
}

func (i TargetReferenceArray) ToTargetReferenceArrayOutput() TargetReferenceArrayOutput {
	return i.ToTargetReferenceArrayOutputWithContext(context.Background())
}

func (i TargetReferenceArray) ToTargetReferenceArrayOutputWithContext(ctx context.Context) TargetReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetReferenceArrayOutput)
}

type TargetReferenceOutput struct{ *pulumi.OutputState }

func (TargetReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetReference)(nil)).Elem()
}

func (o TargetReferenceOutput) ToTargetReferenceOutput() TargetReferenceOutput {
	return o
}

func (o TargetReferenceOutput) ToTargetReferenceOutputWithContext(ctx context.Context) TargetReferenceOutput {
	return o
}

func (o TargetReferenceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TargetReference) string { return v.Id }).(pulumi.StringOutput)
}

func (o TargetReferenceOutput) Type() TargetReferenceTypeOutput {
	return o.ApplyT(func(v TargetReference) TargetReferenceType { return v.Type }).(TargetReferenceTypeOutput)
}

type TargetReferenceArrayOutput struct{ *pulumi.OutputState }

func (TargetReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetReference)(nil)).Elem()
}

func (o TargetReferenceArrayOutput) ToTargetReferenceArrayOutput() TargetReferenceArrayOutput {
	return o
}

func (o TargetReferenceArrayOutput) ToTargetReferenceArrayOutputWithContext(ctx context.Context) TargetReferenceArrayOutput {
	return o
}

func (o TargetReferenceArrayOutput) Index(i pulumi.IntInput) TargetReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetReference {
		return vs[0].([]TargetReference)[vs[1].(int)]
	}).(TargetReferenceOutput)
}

type TargetReferenceResponse struct {
	Id   string `pulumi:"id"`
	Type string `pulumi:"type"`
}





type TargetReferenceResponseInput interface {
	pulumi.Input

	ToTargetReferenceResponseOutput() TargetReferenceResponseOutput
	ToTargetReferenceResponseOutputWithContext(context.Context) TargetReferenceResponseOutput
}

type TargetReferenceResponseArgs struct {
	Id   pulumi.StringInput `pulumi:"id"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (TargetReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetReferenceResponse)(nil)).Elem()
}

func (i TargetReferenceResponseArgs) ToTargetReferenceResponseOutput() TargetReferenceResponseOutput {
	return i.ToTargetReferenceResponseOutputWithContext(context.Background())
}

func (i TargetReferenceResponseArgs) ToTargetReferenceResponseOutputWithContext(ctx context.Context) TargetReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetReferenceResponseOutput)
}





type TargetReferenceResponseArrayInput interface {
	pulumi.Input

	ToTargetReferenceResponseArrayOutput() TargetReferenceResponseArrayOutput
	ToTargetReferenceResponseArrayOutputWithContext(context.Context) TargetReferenceResponseArrayOutput
}

type TargetReferenceResponseArray []TargetReferenceResponseInput

func (TargetReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetReferenceResponse)(nil)).Elem()
}

func (i TargetReferenceResponseArray) ToTargetReferenceResponseArrayOutput() TargetReferenceResponseArrayOutput {
	return i.ToTargetReferenceResponseArrayOutputWithContext(context.Background())
}

func (i TargetReferenceResponseArray) ToTargetReferenceResponseArrayOutputWithContext(ctx context.Context) TargetReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetReferenceResponseArrayOutput)
}

type TargetReferenceResponseOutput struct{ *pulumi.OutputState }

func (TargetReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetReferenceResponse)(nil)).Elem()
}

func (o TargetReferenceResponseOutput) ToTargetReferenceResponseOutput() TargetReferenceResponseOutput {
	return o
}

func (o TargetReferenceResponseOutput) ToTargetReferenceResponseOutputWithContext(ctx context.Context) TargetReferenceResponseOutput {
	return o
}

func (o TargetReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TargetReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o TargetReferenceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TargetReferenceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TargetReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetReferenceResponse)(nil)).Elem()
}

func (o TargetReferenceResponseArrayOutput) ToTargetReferenceResponseArrayOutput() TargetReferenceResponseArrayOutput {
	return o
}

func (o TargetReferenceResponseArrayOutput) ToTargetReferenceResponseArrayOutputWithContext(ctx context.Context) TargetReferenceResponseArrayOutput {
	return o
}

func (o TargetReferenceResponseArrayOutput) Index(i pulumi.IntInput) TargetReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetReferenceResponse {
		return vs[0].([]TargetReferenceResponse)[vs[1].(int)]
	}).(TargetReferenceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(BranchOutput{})
	pulumi.RegisterOutputType(BranchArrayOutput{})
	pulumi.RegisterOutputType(BranchResponseOutput{})
	pulumi.RegisterOutputType(BranchResponseArrayOutput{})
	pulumi.RegisterOutputType(CapabilityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CapabilityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ContinuousActionOutput{})
	pulumi.RegisterOutputType(ContinuousActionResponseOutput{})
	pulumi.RegisterOutputType(DelayActionOutput{})
	pulumi.RegisterOutputType(DelayActionResponseOutput{})
	pulumi.RegisterOutputType(DiscreteActionOutput{})
	pulumi.RegisterOutputType(DiscreteActionResponseOutput{})
	pulumi.RegisterOutputType(ExperimentPropertiesOutput{})
	pulumi.RegisterOutputType(ExperimentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ExperimentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ExperimentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyValuePairOutput{})
	pulumi.RegisterOutputType(KeyValuePairArrayOutput{})
	pulumi.RegisterOutputType(KeyValuePairResponseOutput{})
	pulumi.RegisterOutputType(KeyValuePairResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SelectorOutput{})
	pulumi.RegisterOutputType(SelectorArrayOutput{})
	pulumi.RegisterOutputType(SelectorResponseOutput{})
	pulumi.RegisterOutputType(SelectorResponseArrayOutput{})
	pulumi.RegisterOutputType(StepOutput{})
	pulumi.RegisterOutputType(StepArrayOutput{})
	pulumi.RegisterOutputType(StepResponseOutput{})
	pulumi.RegisterOutputType(StepResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetReferenceOutput{})
	pulumi.RegisterOutputType(TargetReferenceArrayOutput{})
	pulumi.RegisterOutputType(TargetReferenceResponseOutput{})
	pulumi.RegisterOutputType(TargetReferenceResponseArrayOutput{})
}
