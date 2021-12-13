


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssessmentLinksResponse struct {
	AzurePortalUri string `pulumi:"azurePortalUri"`
}

type AssessmentLinksResponseOutput struct{ *pulumi.OutputState }

func (AssessmentLinksResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentLinksResponse)(nil)).Elem()
}

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponseOutput() AssessmentLinksResponseOutput {
	return o
}

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponseOutputWithContext(ctx context.Context) AssessmentLinksResponseOutput {
	return o
}

func (o AssessmentLinksResponseOutput) AzurePortalUri() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentLinksResponse) string { return v.AzurePortalUri }).(pulumi.StringOutput)
}

type AssessmentStatus struct {
	Cause       *string `pulumi:"cause"`
	Code        string  `pulumi:"code"`
	Description *string `pulumi:"description"`
}





type AssessmentStatusInput interface {
	pulumi.Input

	ToAssessmentStatusOutput() AssessmentStatusOutput
	ToAssessmentStatusOutputWithContext(context.Context) AssessmentStatusOutput
}

type AssessmentStatusArgs struct {
	Cause       pulumi.StringPtrInput `pulumi:"cause"`
	Code        pulumi.StringInput    `pulumi:"code"`
	Description pulumi.StringPtrInput `pulumi:"description"`
}

func (AssessmentStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatus)(nil)).Elem()
}

func (i AssessmentStatusArgs) ToAssessmentStatusOutput() AssessmentStatusOutput {
	return i.ToAssessmentStatusOutputWithContext(context.Background())
}

func (i AssessmentStatusArgs) ToAssessmentStatusOutputWithContext(ctx context.Context) AssessmentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusOutput)
}

type AssessmentStatusOutput struct{ *pulumi.OutputState }

func (AssessmentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatus)(nil)).Elem()
}

func (o AssessmentStatusOutput) ToAssessmentStatusOutput() AssessmentStatusOutput {
	return o
}

func (o AssessmentStatusOutput) ToAssessmentStatusOutputWithContext(ctx context.Context) AssessmentStatusOutput {
	return o
}

func (o AssessmentStatusOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatus) *string { return v.Cause }).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentStatus) string { return v.Code }).(pulumi.StringOutput)
}

func (o AssessmentStatusOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatus) *string { return v.Description }).(pulumi.StringPtrOutput)
}

type AssessmentStatusResponse struct {
	Cause       *string `pulumi:"cause"`
	Code        string  `pulumi:"code"`
	Description *string `pulumi:"description"`
}

type AssessmentStatusResponseOutput struct{ *pulumi.OutputState }

func (AssessmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatusResponse)(nil)).Elem()
}

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponseOutput() AssessmentStatusResponseOutput {
	return o
}

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponseOutputWithContext(ctx context.Context) AssessmentStatusResponseOutput {
	return o
}

func (o AssessmentStatusResponseOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatusResponse) *string { return v.Cause }).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentStatusResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o AssessmentStatusResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatusResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

type AutomationActionEventHub struct {
	ActionType         string  `pulumi:"actionType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EventHubResourceId *string `pulumi:"eventHubResourceId"`
}

type AutomationActionEventHubResponse struct {
	ActionType         string  `pulumi:"actionType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EventHubResourceId *string `pulumi:"eventHubResourceId"`
	SasPolicyName      string  `pulumi:"sasPolicyName"`
}

type AutomationActionLogicApp struct {
	ActionType         string  `pulumi:"actionType"`
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	Uri                *string `pulumi:"uri"`
}

type AutomationActionLogicAppResponse struct {
	ActionType         string  `pulumi:"actionType"`
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	Uri                *string `pulumi:"uri"`
}

type AutomationActionWorkspace struct {
	ActionType          string  `pulumi:"actionType"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}

type AutomationActionWorkspaceResponse struct {
	ActionType          string  `pulumi:"actionType"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}

type AutomationRuleSet struct {
	Rules []AutomationTriggeringRule `pulumi:"rules"`
}





type AutomationRuleSetInput interface {
	pulumi.Input

	ToAutomationRuleSetOutput() AutomationRuleSetOutput
	ToAutomationRuleSetOutputWithContext(context.Context) AutomationRuleSetOutput
}

type AutomationRuleSetArgs struct {
	Rules AutomationTriggeringRuleArrayInput `pulumi:"rules"`
}

func (AutomationRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleSet)(nil)).Elem()
}

func (i AutomationRuleSetArgs) ToAutomationRuleSetOutput() AutomationRuleSetOutput {
	return i.ToAutomationRuleSetOutputWithContext(context.Background())
}

func (i AutomationRuleSetArgs) ToAutomationRuleSetOutputWithContext(ctx context.Context) AutomationRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleSetOutput)
}





type AutomationRuleSetArrayInput interface {
	pulumi.Input

	ToAutomationRuleSetArrayOutput() AutomationRuleSetArrayOutput
	ToAutomationRuleSetArrayOutputWithContext(context.Context) AutomationRuleSetArrayOutput
}

type AutomationRuleSetArray []AutomationRuleSetInput

func (AutomationRuleSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRuleSet)(nil)).Elem()
}

func (i AutomationRuleSetArray) ToAutomationRuleSetArrayOutput() AutomationRuleSetArrayOutput {
	return i.ToAutomationRuleSetArrayOutputWithContext(context.Background())
}

func (i AutomationRuleSetArray) ToAutomationRuleSetArrayOutputWithContext(ctx context.Context) AutomationRuleSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleSetArrayOutput)
}

type AutomationRuleSetOutput struct{ *pulumi.OutputState }

func (AutomationRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleSet)(nil)).Elem()
}

func (o AutomationRuleSetOutput) ToAutomationRuleSetOutput() AutomationRuleSetOutput {
	return o
}

func (o AutomationRuleSetOutput) ToAutomationRuleSetOutputWithContext(ctx context.Context) AutomationRuleSetOutput {
	return o
}

func (o AutomationRuleSetOutput) Rules() AutomationTriggeringRuleArrayOutput {
	return o.ApplyT(func(v AutomationRuleSet) []AutomationTriggeringRule { return v.Rules }).(AutomationTriggeringRuleArrayOutput)
}

type AutomationRuleSetArrayOutput struct{ *pulumi.OutputState }

func (AutomationRuleSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRuleSet)(nil)).Elem()
}

func (o AutomationRuleSetArrayOutput) ToAutomationRuleSetArrayOutput() AutomationRuleSetArrayOutput {
	return o
}

func (o AutomationRuleSetArrayOutput) ToAutomationRuleSetArrayOutputWithContext(ctx context.Context) AutomationRuleSetArrayOutput {
	return o
}

func (o AutomationRuleSetArrayOutput) Index(i pulumi.IntInput) AutomationRuleSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationRuleSet {
		return vs[0].([]AutomationRuleSet)[vs[1].(int)]
	}).(AutomationRuleSetOutput)
}

type AutomationRuleSetResponse struct {
	Rules []AutomationTriggeringRuleResponse `pulumi:"rules"`
}

type AutomationRuleSetResponseOutput struct{ *pulumi.OutputState }

func (AutomationRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleSetResponse)(nil)).Elem()
}

func (o AutomationRuleSetResponseOutput) ToAutomationRuleSetResponseOutput() AutomationRuleSetResponseOutput {
	return o
}

func (o AutomationRuleSetResponseOutput) ToAutomationRuleSetResponseOutputWithContext(ctx context.Context) AutomationRuleSetResponseOutput {
	return o
}

func (o AutomationRuleSetResponseOutput) Rules() AutomationTriggeringRuleResponseArrayOutput {
	return o.ApplyT(func(v AutomationRuleSetResponse) []AutomationTriggeringRuleResponse { return v.Rules }).(AutomationTriggeringRuleResponseArrayOutput)
}

type AutomationRuleSetResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationRuleSetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRuleSetResponse)(nil)).Elem()
}

func (o AutomationRuleSetResponseArrayOutput) ToAutomationRuleSetResponseArrayOutput() AutomationRuleSetResponseArrayOutput {
	return o
}

func (o AutomationRuleSetResponseArrayOutput) ToAutomationRuleSetResponseArrayOutputWithContext(ctx context.Context) AutomationRuleSetResponseArrayOutput {
	return o
}

func (o AutomationRuleSetResponseArrayOutput) Index(i pulumi.IntInput) AutomationRuleSetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationRuleSetResponse {
		return vs[0].([]AutomationRuleSetResponse)[vs[1].(int)]
	}).(AutomationRuleSetResponseOutput)
}

type AutomationScope struct {
	Description *string `pulumi:"description"`
	ScopePath   *string `pulumi:"scopePath"`
}





type AutomationScopeInput interface {
	pulumi.Input

	ToAutomationScopeOutput() AutomationScopeOutput
	ToAutomationScopeOutputWithContext(context.Context) AutomationScopeOutput
}

type AutomationScopeArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	ScopePath   pulumi.StringPtrInput `pulumi:"scopePath"`
}

func (AutomationScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationScope)(nil)).Elem()
}

func (i AutomationScopeArgs) ToAutomationScopeOutput() AutomationScopeOutput {
	return i.ToAutomationScopeOutputWithContext(context.Background())
}

func (i AutomationScopeArgs) ToAutomationScopeOutputWithContext(ctx context.Context) AutomationScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationScopeOutput)
}





type AutomationScopeArrayInput interface {
	pulumi.Input

	ToAutomationScopeArrayOutput() AutomationScopeArrayOutput
	ToAutomationScopeArrayOutputWithContext(context.Context) AutomationScopeArrayOutput
}

type AutomationScopeArray []AutomationScopeInput

func (AutomationScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationScope)(nil)).Elem()
}

func (i AutomationScopeArray) ToAutomationScopeArrayOutput() AutomationScopeArrayOutput {
	return i.ToAutomationScopeArrayOutputWithContext(context.Background())
}

func (i AutomationScopeArray) ToAutomationScopeArrayOutputWithContext(ctx context.Context) AutomationScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationScopeArrayOutput)
}

type AutomationScopeOutput struct{ *pulumi.OutputState }

func (AutomationScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationScope)(nil)).Elem()
}

func (o AutomationScopeOutput) ToAutomationScopeOutput() AutomationScopeOutput {
	return o
}

func (o AutomationScopeOutput) ToAutomationScopeOutputWithContext(ctx context.Context) AutomationScopeOutput {
	return o
}

func (o AutomationScopeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationScope) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AutomationScopeOutput) ScopePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationScope) *string { return v.ScopePath }).(pulumi.StringPtrOutput)
}

type AutomationScopeArrayOutput struct{ *pulumi.OutputState }

func (AutomationScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationScope)(nil)).Elem()
}

func (o AutomationScopeArrayOutput) ToAutomationScopeArrayOutput() AutomationScopeArrayOutput {
	return o
}

func (o AutomationScopeArrayOutput) ToAutomationScopeArrayOutputWithContext(ctx context.Context) AutomationScopeArrayOutput {
	return o
}

func (o AutomationScopeArrayOutput) Index(i pulumi.IntInput) AutomationScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationScope {
		return vs[0].([]AutomationScope)[vs[1].(int)]
	}).(AutomationScopeOutput)
}

type AutomationScopeResponse struct {
	Description *string `pulumi:"description"`
	ScopePath   *string `pulumi:"scopePath"`
}

type AutomationScopeResponseOutput struct{ *pulumi.OutputState }

func (AutomationScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationScopeResponse)(nil)).Elem()
}

func (o AutomationScopeResponseOutput) ToAutomationScopeResponseOutput() AutomationScopeResponseOutput {
	return o
}

func (o AutomationScopeResponseOutput) ToAutomationScopeResponseOutputWithContext(ctx context.Context) AutomationScopeResponseOutput {
	return o
}

func (o AutomationScopeResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationScopeResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AutomationScopeResponseOutput) ScopePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationScopeResponse) *string { return v.ScopePath }).(pulumi.StringPtrOutput)
}

type AutomationScopeResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationScopeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationScopeResponse)(nil)).Elem()
}

func (o AutomationScopeResponseArrayOutput) ToAutomationScopeResponseArrayOutput() AutomationScopeResponseArrayOutput {
	return o
}

func (o AutomationScopeResponseArrayOutput) ToAutomationScopeResponseArrayOutputWithContext(ctx context.Context) AutomationScopeResponseArrayOutput {
	return o
}

func (o AutomationScopeResponseArrayOutput) Index(i pulumi.IntInput) AutomationScopeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationScopeResponse {
		return vs[0].([]AutomationScopeResponse)[vs[1].(int)]
	}).(AutomationScopeResponseOutput)
}

type AutomationSource struct {
	EventSource *string             `pulumi:"eventSource"`
	RuleSets    []AutomationRuleSet `pulumi:"ruleSets"`
}





type AutomationSourceInput interface {
	pulumi.Input

	ToAutomationSourceOutput() AutomationSourceOutput
	ToAutomationSourceOutputWithContext(context.Context) AutomationSourceOutput
}

type AutomationSourceArgs struct {
	EventSource pulumi.StringPtrInput       `pulumi:"eventSource"`
	RuleSets    AutomationRuleSetArrayInput `pulumi:"ruleSets"`
}

func (AutomationSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationSource)(nil)).Elem()
}

func (i AutomationSourceArgs) ToAutomationSourceOutput() AutomationSourceOutput {
	return i.ToAutomationSourceOutputWithContext(context.Background())
}

func (i AutomationSourceArgs) ToAutomationSourceOutputWithContext(ctx context.Context) AutomationSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationSourceOutput)
}





type AutomationSourceArrayInput interface {
	pulumi.Input

	ToAutomationSourceArrayOutput() AutomationSourceArrayOutput
	ToAutomationSourceArrayOutputWithContext(context.Context) AutomationSourceArrayOutput
}

type AutomationSourceArray []AutomationSourceInput

func (AutomationSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationSource)(nil)).Elem()
}

func (i AutomationSourceArray) ToAutomationSourceArrayOutput() AutomationSourceArrayOutput {
	return i.ToAutomationSourceArrayOutputWithContext(context.Background())
}

func (i AutomationSourceArray) ToAutomationSourceArrayOutputWithContext(ctx context.Context) AutomationSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationSourceArrayOutput)
}

type AutomationSourceOutput struct{ *pulumi.OutputState }

func (AutomationSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationSource)(nil)).Elem()
}

func (o AutomationSourceOutput) ToAutomationSourceOutput() AutomationSourceOutput {
	return o
}

func (o AutomationSourceOutput) ToAutomationSourceOutputWithContext(ctx context.Context) AutomationSourceOutput {
	return o
}

func (o AutomationSourceOutput) EventSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationSource) *string { return v.EventSource }).(pulumi.StringPtrOutput)
}

func (o AutomationSourceOutput) RuleSets() AutomationRuleSetArrayOutput {
	return o.ApplyT(func(v AutomationSource) []AutomationRuleSet { return v.RuleSets }).(AutomationRuleSetArrayOutput)
}

type AutomationSourceArrayOutput struct{ *pulumi.OutputState }

func (AutomationSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationSource)(nil)).Elem()
}

func (o AutomationSourceArrayOutput) ToAutomationSourceArrayOutput() AutomationSourceArrayOutput {
	return o
}

func (o AutomationSourceArrayOutput) ToAutomationSourceArrayOutputWithContext(ctx context.Context) AutomationSourceArrayOutput {
	return o
}

func (o AutomationSourceArrayOutput) Index(i pulumi.IntInput) AutomationSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationSource {
		return vs[0].([]AutomationSource)[vs[1].(int)]
	}).(AutomationSourceOutput)
}

type AutomationSourceResponse struct {
	EventSource *string                     `pulumi:"eventSource"`
	RuleSets    []AutomationRuleSetResponse `pulumi:"ruleSets"`
}

type AutomationSourceResponseOutput struct{ *pulumi.OutputState }

func (AutomationSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationSourceResponse)(nil)).Elem()
}

func (o AutomationSourceResponseOutput) ToAutomationSourceResponseOutput() AutomationSourceResponseOutput {
	return o
}

func (o AutomationSourceResponseOutput) ToAutomationSourceResponseOutputWithContext(ctx context.Context) AutomationSourceResponseOutput {
	return o
}

func (o AutomationSourceResponseOutput) EventSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationSourceResponse) *string { return v.EventSource }).(pulumi.StringPtrOutput)
}

func (o AutomationSourceResponseOutput) RuleSets() AutomationRuleSetResponseArrayOutput {
	return o.ApplyT(func(v AutomationSourceResponse) []AutomationRuleSetResponse { return v.RuleSets }).(AutomationRuleSetResponseArrayOutput)
}

type AutomationSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationSourceResponse)(nil)).Elem()
}

func (o AutomationSourceResponseArrayOutput) ToAutomationSourceResponseArrayOutput() AutomationSourceResponseArrayOutput {
	return o
}

func (o AutomationSourceResponseArrayOutput) ToAutomationSourceResponseArrayOutputWithContext(ctx context.Context) AutomationSourceResponseArrayOutput {
	return o
}

func (o AutomationSourceResponseArrayOutput) Index(i pulumi.IntInput) AutomationSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationSourceResponse {
		return vs[0].([]AutomationSourceResponse)[vs[1].(int)]
	}).(AutomationSourceResponseOutput)
}

type AutomationTriggeringRule struct {
	ExpectedValue *string `pulumi:"expectedValue"`
	Operator      *string `pulumi:"operator"`
	PropertyJPath *string `pulumi:"propertyJPath"`
	PropertyType  *string `pulumi:"propertyType"`
}





type AutomationTriggeringRuleInput interface {
	pulumi.Input

	ToAutomationTriggeringRuleOutput() AutomationTriggeringRuleOutput
	ToAutomationTriggeringRuleOutputWithContext(context.Context) AutomationTriggeringRuleOutput
}

type AutomationTriggeringRuleArgs struct {
	ExpectedValue pulumi.StringPtrInput `pulumi:"expectedValue"`
	Operator      pulumi.StringPtrInput `pulumi:"operator"`
	PropertyJPath pulumi.StringPtrInput `pulumi:"propertyJPath"`
	PropertyType  pulumi.StringPtrInput `pulumi:"propertyType"`
}

func (AutomationTriggeringRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationTriggeringRule)(nil)).Elem()
}

func (i AutomationTriggeringRuleArgs) ToAutomationTriggeringRuleOutput() AutomationTriggeringRuleOutput {
	return i.ToAutomationTriggeringRuleOutputWithContext(context.Background())
}

func (i AutomationTriggeringRuleArgs) ToAutomationTriggeringRuleOutputWithContext(ctx context.Context) AutomationTriggeringRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationTriggeringRuleOutput)
}





type AutomationTriggeringRuleArrayInput interface {
	pulumi.Input

	ToAutomationTriggeringRuleArrayOutput() AutomationTriggeringRuleArrayOutput
	ToAutomationTriggeringRuleArrayOutputWithContext(context.Context) AutomationTriggeringRuleArrayOutput
}

type AutomationTriggeringRuleArray []AutomationTriggeringRuleInput

func (AutomationTriggeringRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationTriggeringRule)(nil)).Elem()
}

func (i AutomationTriggeringRuleArray) ToAutomationTriggeringRuleArrayOutput() AutomationTriggeringRuleArrayOutput {
	return i.ToAutomationTriggeringRuleArrayOutputWithContext(context.Background())
}

func (i AutomationTriggeringRuleArray) ToAutomationTriggeringRuleArrayOutputWithContext(ctx context.Context) AutomationTriggeringRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationTriggeringRuleArrayOutput)
}

type AutomationTriggeringRuleOutput struct{ *pulumi.OutputState }

func (AutomationTriggeringRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationTriggeringRule)(nil)).Elem()
}

func (o AutomationTriggeringRuleOutput) ToAutomationTriggeringRuleOutput() AutomationTriggeringRuleOutput {
	return o
}

func (o AutomationTriggeringRuleOutput) ToAutomationTriggeringRuleOutputWithContext(ctx context.Context) AutomationTriggeringRuleOutput {
	return o
}

func (o AutomationTriggeringRuleOutput) ExpectedValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRule) *string { return v.ExpectedValue }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRule) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleOutput) PropertyJPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRule) *string { return v.PropertyJPath }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleOutput) PropertyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRule) *string { return v.PropertyType }).(pulumi.StringPtrOutput)
}

type AutomationTriggeringRuleArrayOutput struct{ *pulumi.OutputState }

func (AutomationTriggeringRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationTriggeringRule)(nil)).Elem()
}

func (o AutomationTriggeringRuleArrayOutput) ToAutomationTriggeringRuleArrayOutput() AutomationTriggeringRuleArrayOutput {
	return o
}

func (o AutomationTriggeringRuleArrayOutput) ToAutomationTriggeringRuleArrayOutputWithContext(ctx context.Context) AutomationTriggeringRuleArrayOutput {
	return o
}

func (o AutomationTriggeringRuleArrayOutput) Index(i pulumi.IntInput) AutomationTriggeringRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationTriggeringRule {
		return vs[0].([]AutomationTriggeringRule)[vs[1].(int)]
	}).(AutomationTriggeringRuleOutput)
}

type AutomationTriggeringRuleResponse struct {
	ExpectedValue *string `pulumi:"expectedValue"`
	Operator      *string `pulumi:"operator"`
	PropertyJPath *string `pulumi:"propertyJPath"`
	PropertyType  *string `pulumi:"propertyType"`
}

type AutomationTriggeringRuleResponseOutput struct{ *pulumi.OutputState }

func (AutomationTriggeringRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationTriggeringRuleResponse)(nil)).Elem()
}

func (o AutomationTriggeringRuleResponseOutput) ToAutomationTriggeringRuleResponseOutput() AutomationTriggeringRuleResponseOutput {
	return o
}

func (o AutomationTriggeringRuleResponseOutput) ToAutomationTriggeringRuleResponseOutputWithContext(ctx context.Context) AutomationTriggeringRuleResponseOutput {
	return o
}

func (o AutomationTriggeringRuleResponseOutput) ExpectedValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRuleResponse) *string { return v.ExpectedValue }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleResponseOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRuleResponse) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleResponseOutput) PropertyJPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRuleResponse) *string { return v.PropertyJPath }).(pulumi.StringPtrOutput)
}

func (o AutomationTriggeringRuleResponseOutput) PropertyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationTriggeringRuleResponse) *string { return v.PropertyType }).(pulumi.StringPtrOutput)
}

type AutomationTriggeringRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationTriggeringRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationTriggeringRuleResponse)(nil)).Elem()
}

func (o AutomationTriggeringRuleResponseArrayOutput) ToAutomationTriggeringRuleResponseArrayOutput() AutomationTriggeringRuleResponseArrayOutput {
	return o
}

func (o AutomationTriggeringRuleResponseArrayOutput) ToAutomationTriggeringRuleResponseArrayOutputWithContext(ctx context.Context) AutomationTriggeringRuleResponseArrayOutput {
	return o
}

func (o AutomationTriggeringRuleResponseArrayOutput) Index(i pulumi.IntInput) AutomationTriggeringRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationTriggeringRuleResponse {
		return vs[0].([]AutomationTriggeringRuleResponse)[vs[1].(int)]
	}).(AutomationTriggeringRuleResponseOutput)
}

type AzureResourceDetails struct {
	Source string `pulumi:"source"`
}

type AzureResourceDetailsResponse struct {
	Id     string `pulumi:"id"`
	Source string `pulumi:"source"`
}

type OnPremiseResourceDetails struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type OnPremiseResourceDetailsResponse struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type OnPremiseSqlResourceDetails struct {
	DatabaseName     string `pulumi:"databaseName"`
	MachineName      string `pulumi:"machineName"`
	ServerName       string `pulumi:"serverName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type OnPremiseSqlResourceDetailsResponse struct {
	DatabaseName     string `pulumi:"databaseName"`
	MachineName      string `pulumi:"machineName"`
	ServerName       string `pulumi:"serverName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type ScopeElement struct {
	Field *string `pulumi:"field"`
}





type ScopeElementInput interface {
	pulumi.Input

	ToScopeElementOutput() ScopeElementOutput
	ToScopeElementOutputWithContext(context.Context) ScopeElementOutput
}

type ScopeElementArgs struct {
	Field pulumi.StringPtrInput `pulumi:"field"`
}

func (ScopeElementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeElement)(nil)).Elem()
}

func (i ScopeElementArgs) ToScopeElementOutput() ScopeElementOutput {
	return i.ToScopeElementOutputWithContext(context.Background())
}

func (i ScopeElementArgs) ToScopeElementOutputWithContext(ctx context.Context) ScopeElementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeElementOutput)
}





type ScopeElementArrayInput interface {
	pulumi.Input

	ToScopeElementArrayOutput() ScopeElementArrayOutput
	ToScopeElementArrayOutputWithContext(context.Context) ScopeElementArrayOutput
}

type ScopeElementArray []ScopeElementInput

func (ScopeElementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScopeElement)(nil)).Elem()
}

func (i ScopeElementArray) ToScopeElementArrayOutput() ScopeElementArrayOutput {
	return i.ToScopeElementArrayOutputWithContext(context.Background())
}

func (i ScopeElementArray) ToScopeElementArrayOutputWithContext(ctx context.Context) ScopeElementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeElementArrayOutput)
}

type ScopeElementOutput struct{ *pulumi.OutputState }

func (ScopeElementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeElement)(nil)).Elem()
}

func (o ScopeElementOutput) ToScopeElementOutput() ScopeElementOutput {
	return o
}

func (o ScopeElementOutput) ToScopeElementOutputWithContext(ctx context.Context) ScopeElementOutput {
	return o
}

func (o ScopeElementOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeElement) *string { return v.Field }).(pulumi.StringPtrOutput)
}

type ScopeElementArrayOutput struct{ *pulumi.OutputState }

func (ScopeElementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScopeElement)(nil)).Elem()
}

func (o ScopeElementArrayOutput) ToScopeElementArrayOutput() ScopeElementArrayOutput {
	return o
}

func (o ScopeElementArrayOutput) ToScopeElementArrayOutputWithContext(ctx context.Context) ScopeElementArrayOutput {
	return o
}

func (o ScopeElementArrayOutput) Index(i pulumi.IntInput) ScopeElementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScopeElement {
		return vs[0].([]ScopeElement)[vs[1].(int)]
	}).(ScopeElementOutput)
}

type ScopeElementResponse struct {
	Field *string `pulumi:"field"`
}

type ScopeElementResponseOutput struct{ *pulumi.OutputState }

func (ScopeElementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeElementResponse)(nil)).Elem()
}

func (o ScopeElementResponseOutput) ToScopeElementResponseOutput() ScopeElementResponseOutput {
	return o
}

func (o ScopeElementResponseOutput) ToScopeElementResponseOutputWithContext(ctx context.Context) ScopeElementResponseOutput {
	return o
}

func (o ScopeElementResponseOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeElementResponse) *string { return v.Field }).(pulumi.StringPtrOutput)
}

type ScopeElementResponseArrayOutput struct{ *pulumi.OutputState }

func (ScopeElementResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScopeElementResponse)(nil)).Elem()
}

func (o ScopeElementResponseArrayOutput) ToScopeElementResponseArrayOutput() ScopeElementResponseArrayOutput {
	return o
}

func (o ScopeElementResponseArrayOutput) ToScopeElementResponseArrayOutputWithContext(ctx context.Context) ScopeElementResponseArrayOutput {
	return o
}

func (o ScopeElementResponseArrayOutput) Index(i pulumi.IntInput) ScopeElementResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScopeElementResponse {
		return vs[0].([]ScopeElementResponse)[vs[1].(int)]
	}).(ScopeElementResponseOutput)
}

type SuppressionAlertsScope struct {
	AllOf []ScopeElement `pulumi:"allOf"`
}





type SuppressionAlertsScopeInput interface {
	pulumi.Input

	ToSuppressionAlertsScopeOutput() SuppressionAlertsScopeOutput
	ToSuppressionAlertsScopeOutputWithContext(context.Context) SuppressionAlertsScopeOutput
}

type SuppressionAlertsScopeArgs struct {
	AllOf ScopeElementArrayInput `pulumi:"allOf"`
}

func (SuppressionAlertsScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionAlertsScope)(nil)).Elem()
}

func (i SuppressionAlertsScopeArgs) ToSuppressionAlertsScopeOutput() SuppressionAlertsScopeOutput {
	return i.ToSuppressionAlertsScopeOutputWithContext(context.Background())
}

func (i SuppressionAlertsScopeArgs) ToSuppressionAlertsScopeOutputWithContext(ctx context.Context) SuppressionAlertsScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeOutput)
}

func (i SuppressionAlertsScopeArgs) ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput {
	return i.ToSuppressionAlertsScopePtrOutputWithContext(context.Background())
}

func (i SuppressionAlertsScopeArgs) ToSuppressionAlertsScopePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeOutput).ToSuppressionAlertsScopePtrOutputWithContext(ctx)
}









type SuppressionAlertsScopePtrInput interface {
	pulumi.Input

	ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput
	ToSuppressionAlertsScopePtrOutputWithContext(context.Context) SuppressionAlertsScopePtrOutput
}

type suppressionAlertsScopePtrType SuppressionAlertsScopeArgs

func SuppressionAlertsScopePtr(v *SuppressionAlertsScopeArgs) SuppressionAlertsScopePtrInput {
	return (*suppressionAlertsScopePtrType)(v)
}

func (*suppressionAlertsScopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionAlertsScope)(nil)).Elem()
}

func (i *suppressionAlertsScopePtrType) ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput {
	return i.ToSuppressionAlertsScopePtrOutputWithContext(context.Background())
}

func (i *suppressionAlertsScopePtrType) ToSuppressionAlertsScopePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopePtrOutput)
}

type SuppressionAlertsScopeOutput struct{ *pulumi.OutputState }

func (SuppressionAlertsScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionAlertsScope)(nil)).Elem()
}

func (o SuppressionAlertsScopeOutput) ToSuppressionAlertsScopeOutput() SuppressionAlertsScopeOutput {
	return o
}

func (o SuppressionAlertsScopeOutput) ToSuppressionAlertsScopeOutputWithContext(ctx context.Context) SuppressionAlertsScopeOutput {
	return o
}

func (o SuppressionAlertsScopeOutput) ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput {
	return o.ToSuppressionAlertsScopePtrOutputWithContext(context.Background())
}

func (o SuppressionAlertsScopeOutput) ToSuppressionAlertsScopePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SuppressionAlertsScope) *SuppressionAlertsScope {
		return &v
	}).(SuppressionAlertsScopePtrOutput)
}

func (o SuppressionAlertsScopeOutput) AllOf() ScopeElementArrayOutput {
	return o.ApplyT(func(v SuppressionAlertsScope) []ScopeElement { return v.AllOf }).(ScopeElementArrayOutput)
}

type SuppressionAlertsScopePtrOutput struct{ *pulumi.OutputState }

func (SuppressionAlertsScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionAlertsScope)(nil)).Elem()
}

func (o SuppressionAlertsScopePtrOutput) ToSuppressionAlertsScopePtrOutput() SuppressionAlertsScopePtrOutput {
	return o
}

func (o SuppressionAlertsScopePtrOutput) ToSuppressionAlertsScopePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopePtrOutput {
	return o
}

func (o SuppressionAlertsScopePtrOutput) Elem() SuppressionAlertsScopeOutput {
	return o.ApplyT(func(v *SuppressionAlertsScope) SuppressionAlertsScope {
		if v != nil {
			return *v
		}
		var ret SuppressionAlertsScope
		return ret
	}).(SuppressionAlertsScopeOutput)
}

func (o SuppressionAlertsScopePtrOutput) AllOf() ScopeElementArrayOutput {
	return o.ApplyT(func(v *SuppressionAlertsScope) []ScopeElement {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(ScopeElementArrayOutput)
}

type SuppressionAlertsScopeResponse struct {
	AllOf []ScopeElementResponse `pulumi:"allOf"`
}

type SuppressionAlertsScopeResponseOutput struct{ *pulumi.OutputState }

func (SuppressionAlertsScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionAlertsScopeResponse)(nil)).Elem()
}

func (o SuppressionAlertsScopeResponseOutput) ToSuppressionAlertsScopeResponseOutput() SuppressionAlertsScopeResponseOutput {
	return o
}

func (o SuppressionAlertsScopeResponseOutput) ToSuppressionAlertsScopeResponseOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponseOutput {
	return o
}

func (o SuppressionAlertsScopeResponseOutput) AllOf() ScopeElementResponseArrayOutput {
	return o.ApplyT(func(v SuppressionAlertsScopeResponse) []ScopeElementResponse { return v.AllOf }).(ScopeElementResponseArrayOutput)
}

type SuppressionAlertsScopeResponsePtrOutput struct{ *pulumi.OutputState }

func (SuppressionAlertsScopeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionAlertsScopeResponse)(nil)).Elem()
}

func (o SuppressionAlertsScopeResponsePtrOutput) ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput {
	return o
}

func (o SuppressionAlertsScopeResponsePtrOutput) ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponsePtrOutput {
	return o
}

func (o SuppressionAlertsScopeResponsePtrOutput) Elem() SuppressionAlertsScopeResponseOutput {
	return o.ApplyT(func(v *SuppressionAlertsScopeResponse) SuppressionAlertsScopeResponse {
		if v != nil {
			return *v
		}
		var ret SuppressionAlertsScopeResponse
		return ret
	}).(SuppressionAlertsScopeResponseOutput)
}

func (o SuppressionAlertsScopeResponsePtrOutput) AllOf() ScopeElementResponseArrayOutput {
	return o.ApplyT(func(v *SuppressionAlertsScopeResponse) []ScopeElementResponse {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(ScopeElementResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentLinksResponseOutput{})
	pulumi.RegisterOutputType(AssessmentStatusOutput{})
	pulumi.RegisterOutputType(AssessmentStatusResponseOutput{})
	pulumi.RegisterOutputType(AutomationRuleSetOutput{})
	pulumi.RegisterOutputType(AutomationRuleSetArrayOutput{})
	pulumi.RegisterOutputType(AutomationRuleSetResponseOutput{})
	pulumi.RegisterOutputType(AutomationRuleSetResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationScopeOutput{})
	pulumi.RegisterOutputType(AutomationScopeArrayOutput{})
	pulumi.RegisterOutputType(AutomationScopeResponseOutput{})
	pulumi.RegisterOutputType(AutomationScopeResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationSourceOutput{})
	pulumi.RegisterOutputType(AutomationSourceArrayOutput{})
	pulumi.RegisterOutputType(AutomationSourceResponseOutput{})
	pulumi.RegisterOutputType(AutomationSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationTriggeringRuleOutput{})
	pulumi.RegisterOutputType(AutomationTriggeringRuleArrayOutput{})
	pulumi.RegisterOutputType(AutomationTriggeringRuleResponseOutput{})
	pulumi.RegisterOutputType(AutomationTriggeringRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ScopeElementOutput{})
	pulumi.RegisterOutputType(ScopeElementArrayOutput{})
	pulumi.RegisterOutputType(ScopeElementResponseOutput{})
	pulumi.RegisterOutputType(ScopeElementResponseArrayOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopeOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopePtrOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopeResponseOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopeResponsePtrOutput{})
}
