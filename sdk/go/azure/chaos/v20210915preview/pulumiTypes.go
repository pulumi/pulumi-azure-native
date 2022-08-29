


package v20210915preview

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

type ContinuousAction struct {
	Duration   string         `pulumi:"duration"`
	Name       string         `pulumi:"name"`
	Parameters []KeyValuePair `pulumi:"parameters"`
	SelectorId string         `pulumi:"selectorId"`
	Type       string         `pulumi:"type"`
}

type ContinuousActionResponse struct {
	Duration   string                 `pulumi:"duration"`
	Name       string                 `pulumi:"name"`
	Parameters []KeyValuePairResponse `pulumi:"parameters"`
	SelectorId string                 `pulumi:"selectorId"`
	Type       string                 `pulumi:"type"`
}

type DelayAction struct {
	Duration string `pulumi:"duration"`
	Name     string `pulumi:"name"`
	Type     string `pulumi:"type"`
}

type DelayActionResponse struct {
	Duration string `pulumi:"duration"`
	Name     string `pulumi:"name"`
	Type     string `pulumi:"type"`
}

type DiscreteAction struct {
	Name       string         `pulumi:"name"`
	Parameters []KeyValuePair `pulumi:"parameters"`
	SelectorId string         `pulumi:"selectorId"`
	Type       string         `pulumi:"type"`
}

type DiscreteActionResponse struct {
	Name       string                 `pulumi:"name"`
	Parameters []KeyValuePairResponse `pulumi:"parameters"`
	SelectorId string                 `pulumi:"selectorId"`
	Type       string                 `pulumi:"type"`
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

func (o ExperimentPropertiesOutput) Selectors() SelectorArrayOutput {
	return o.ApplyT(func(v ExperimentProperties) []Selector { return v.Selectors }).(SelectorArrayOutput)
}

func (o ExperimentPropertiesOutput) StartOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExperimentProperties) *bool { return v.StartOnCreation }).(pulumi.BoolPtrOutput)
}

func (o ExperimentPropertiesOutput) Steps() StepArrayOutput {
	return o.ApplyT(func(v ExperimentProperties) []Step { return v.Steps }).(StepArrayOutput)
}

type ExperimentPropertiesResponse struct {
	Selectors       []SelectorResponse `pulumi:"selectors"`
	StartOnCreation *bool              `pulumi:"startOnCreation"`
	Steps           []StepResponse     `pulumi:"steps"`
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

func (o ExperimentPropertiesResponseOutput) Selectors() SelectorResponseArrayOutput {
	return o.ApplyT(func(v ExperimentPropertiesResponse) []SelectorResponse { return v.Selectors }).(SelectorResponseArrayOutput)
}

func (o ExperimentPropertiesResponseOutput) StartOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ExperimentPropertiesResponse) *bool { return v.StartOnCreation }).(pulumi.BoolPtrOutput)
}

func (o ExperimentPropertiesResponseOutput) Steps() StepResponseArrayOutput {
	return o.ApplyT(func(v ExperimentPropertiesResponse) []StepResponse { return v.Steps }).(StepResponseArrayOutput)
}

type KeyValuePair struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

type KeyValuePairResponse struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
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
	pulumi.RegisterOutputType(ExperimentPropertiesOutput{})
	pulumi.RegisterOutputType(ExperimentPropertiesResponseOutput{})
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
	pulumi.RegisterOutputType(TargetReferenceOutput{})
	pulumi.RegisterOutputType(TargetReferenceArrayOutput{})
	pulumi.RegisterOutputType(TargetReferenceResponseOutput{})
	pulumi.RegisterOutputType(TargetReferenceResponseArrayOutput{})
}
