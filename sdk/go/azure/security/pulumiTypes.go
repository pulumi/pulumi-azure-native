


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdaptiveApplicationControlIssueSummaryResponse struct {
	Issue       *string  `pulumi:"issue"`
	NumberOfVms *float64 `pulumi:"numberOfVms"`
}





type AdaptiveApplicationControlIssueSummaryResponseInput interface {
	pulumi.Input

	ToAdaptiveApplicationControlIssueSummaryResponseOutput() AdaptiveApplicationControlIssueSummaryResponseOutput
	ToAdaptiveApplicationControlIssueSummaryResponseOutputWithContext(context.Context) AdaptiveApplicationControlIssueSummaryResponseOutput
}

type AdaptiveApplicationControlIssueSummaryResponseArgs struct {
	Issue       pulumi.StringPtrInput  `pulumi:"issue"`
	NumberOfVms pulumi.Float64PtrInput `pulumi:"numberOfVms"`
}

func (AdaptiveApplicationControlIssueSummaryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdaptiveApplicationControlIssueSummaryResponse)(nil)).Elem()
}

func (i AdaptiveApplicationControlIssueSummaryResponseArgs) ToAdaptiveApplicationControlIssueSummaryResponseOutput() AdaptiveApplicationControlIssueSummaryResponseOutput {
	return i.ToAdaptiveApplicationControlIssueSummaryResponseOutputWithContext(context.Background())
}

func (i AdaptiveApplicationControlIssueSummaryResponseArgs) ToAdaptiveApplicationControlIssueSummaryResponseOutputWithContext(ctx context.Context) AdaptiveApplicationControlIssueSummaryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdaptiveApplicationControlIssueSummaryResponseOutput)
}





type AdaptiveApplicationControlIssueSummaryResponseArrayInput interface {
	pulumi.Input

	ToAdaptiveApplicationControlIssueSummaryResponseArrayOutput() AdaptiveApplicationControlIssueSummaryResponseArrayOutput
	ToAdaptiveApplicationControlIssueSummaryResponseArrayOutputWithContext(context.Context) AdaptiveApplicationControlIssueSummaryResponseArrayOutput
}

type AdaptiveApplicationControlIssueSummaryResponseArray []AdaptiveApplicationControlIssueSummaryResponseInput

func (AdaptiveApplicationControlIssueSummaryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdaptiveApplicationControlIssueSummaryResponse)(nil)).Elem()
}

func (i AdaptiveApplicationControlIssueSummaryResponseArray) ToAdaptiveApplicationControlIssueSummaryResponseArrayOutput() AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
	return i.ToAdaptiveApplicationControlIssueSummaryResponseArrayOutputWithContext(context.Background())
}

func (i AdaptiveApplicationControlIssueSummaryResponseArray) ToAdaptiveApplicationControlIssueSummaryResponseArrayOutputWithContext(ctx context.Context) AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdaptiveApplicationControlIssueSummaryResponseArrayOutput)
}

type AdaptiveApplicationControlIssueSummaryResponseOutput struct{ *pulumi.OutputState }

func (AdaptiveApplicationControlIssueSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdaptiveApplicationControlIssueSummaryResponse)(nil)).Elem()
}

func (o AdaptiveApplicationControlIssueSummaryResponseOutput) ToAdaptiveApplicationControlIssueSummaryResponseOutput() AdaptiveApplicationControlIssueSummaryResponseOutput {
	return o
}

func (o AdaptiveApplicationControlIssueSummaryResponseOutput) ToAdaptiveApplicationControlIssueSummaryResponseOutputWithContext(ctx context.Context) AdaptiveApplicationControlIssueSummaryResponseOutput {
	return o
}

func (o AdaptiveApplicationControlIssueSummaryResponseOutput) Issue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdaptiveApplicationControlIssueSummaryResponse) *string { return v.Issue }).(pulumi.StringPtrOutput)
}

func (o AdaptiveApplicationControlIssueSummaryResponseOutput) NumberOfVms() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AdaptiveApplicationControlIssueSummaryResponse) *float64 { return v.NumberOfVms }).(pulumi.Float64PtrOutput)
}

type AdaptiveApplicationControlIssueSummaryResponseArrayOutput struct{ *pulumi.OutputState }

func (AdaptiveApplicationControlIssueSummaryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdaptiveApplicationControlIssueSummaryResponse)(nil)).Elem()
}

func (o AdaptiveApplicationControlIssueSummaryResponseArrayOutput) ToAdaptiveApplicationControlIssueSummaryResponseArrayOutput() AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
	return o
}

func (o AdaptiveApplicationControlIssueSummaryResponseArrayOutput) ToAdaptiveApplicationControlIssueSummaryResponseArrayOutputWithContext(ctx context.Context) AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
	return o
}

func (o AdaptiveApplicationControlIssueSummaryResponseArrayOutput) Index(i pulumi.IntInput) AdaptiveApplicationControlIssueSummaryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdaptiveApplicationControlIssueSummaryResponse {
		return vs[0].([]AdaptiveApplicationControlIssueSummaryResponse)[vs[1].(int)]
	}).(AdaptiveApplicationControlIssueSummaryResponseOutput)
}

type AdditionalWorkspacesProperties struct {
	DataTypes []string `pulumi:"dataTypes"`
	Type      *string  `pulumi:"type"`
	Workspace *string  `pulumi:"workspace"`
}





type AdditionalWorkspacesPropertiesInput interface {
	pulumi.Input

	ToAdditionalWorkspacesPropertiesOutput() AdditionalWorkspacesPropertiesOutput
	ToAdditionalWorkspacesPropertiesOutputWithContext(context.Context) AdditionalWorkspacesPropertiesOutput
}

type AdditionalWorkspacesPropertiesArgs struct {
	DataTypes pulumi.StringArrayInput `pulumi:"dataTypes"`
	Type      pulumi.StringPtrInput   `pulumi:"type"`
	Workspace pulumi.StringPtrInput   `pulumi:"workspace"`
}

func (AdditionalWorkspacesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspacesProperties)(nil)).Elem()
}

func (i AdditionalWorkspacesPropertiesArgs) ToAdditionalWorkspacesPropertiesOutput() AdditionalWorkspacesPropertiesOutput {
	return i.ToAdditionalWorkspacesPropertiesOutputWithContext(context.Background())
}

func (i AdditionalWorkspacesPropertiesArgs) ToAdditionalWorkspacesPropertiesOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalWorkspacesPropertiesOutput)
}





type AdditionalWorkspacesPropertiesArrayInput interface {
	pulumi.Input

	ToAdditionalWorkspacesPropertiesArrayOutput() AdditionalWorkspacesPropertiesArrayOutput
	ToAdditionalWorkspacesPropertiesArrayOutputWithContext(context.Context) AdditionalWorkspacesPropertiesArrayOutput
}

type AdditionalWorkspacesPropertiesArray []AdditionalWorkspacesPropertiesInput

func (AdditionalWorkspacesPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalWorkspacesProperties)(nil)).Elem()
}

func (i AdditionalWorkspacesPropertiesArray) ToAdditionalWorkspacesPropertiesArrayOutput() AdditionalWorkspacesPropertiesArrayOutput {
	return i.ToAdditionalWorkspacesPropertiesArrayOutputWithContext(context.Background())
}

func (i AdditionalWorkspacesPropertiesArray) ToAdditionalWorkspacesPropertiesArrayOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalWorkspacesPropertiesArrayOutput)
}

type AdditionalWorkspacesPropertiesOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspacesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspacesProperties)(nil)).Elem()
}

func (o AdditionalWorkspacesPropertiesOutput) ToAdditionalWorkspacesPropertiesOutput() AdditionalWorkspacesPropertiesOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesOutput) ToAdditionalWorkspacesPropertiesOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesOutput) DataTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalWorkspacesProperties) []string { return v.DataTypes }).(pulumi.StringArrayOutput)
}

func (o AdditionalWorkspacesPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalWorkspacesProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o AdditionalWorkspacesPropertiesOutput) Workspace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalWorkspacesProperties) *string { return v.Workspace }).(pulumi.StringPtrOutput)
}

type AdditionalWorkspacesPropertiesArrayOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspacesPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalWorkspacesProperties)(nil)).Elem()
}

func (o AdditionalWorkspacesPropertiesArrayOutput) ToAdditionalWorkspacesPropertiesArrayOutput() AdditionalWorkspacesPropertiesArrayOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesArrayOutput) ToAdditionalWorkspacesPropertiesArrayOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesArrayOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesArrayOutput) Index(i pulumi.IntInput) AdditionalWorkspacesPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalWorkspacesProperties {
		return vs[0].([]AdditionalWorkspacesProperties)[vs[1].(int)]
	}).(AdditionalWorkspacesPropertiesOutput)
}

type AdditionalWorkspacesPropertiesResponse struct {
	DataTypes []string `pulumi:"dataTypes"`
	Type      *string  `pulumi:"type"`
	Workspace *string  `pulumi:"workspace"`
}





type AdditionalWorkspacesPropertiesResponseInput interface {
	pulumi.Input

	ToAdditionalWorkspacesPropertiesResponseOutput() AdditionalWorkspacesPropertiesResponseOutput
	ToAdditionalWorkspacesPropertiesResponseOutputWithContext(context.Context) AdditionalWorkspacesPropertiesResponseOutput
}

type AdditionalWorkspacesPropertiesResponseArgs struct {
	DataTypes pulumi.StringArrayInput `pulumi:"dataTypes"`
	Type      pulumi.StringPtrInput   `pulumi:"type"`
	Workspace pulumi.StringPtrInput   `pulumi:"workspace"`
}

func (AdditionalWorkspacesPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspacesPropertiesResponse)(nil)).Elem()
}

func (i AdditionalWorkspacesPropertiesResponseArgs) ToAdditionalWorkspacesPropertiesResponseOutput() AdditionalWorkspacesPropertiesResponseOutput {
	return i.ToAdditionalWorkspacesPropertiesResponseOutputWithContext(context.Background())
}

func (i AdditionalWorkspacesPropertiesResponseArgs) ToAdditionalWorkspacesPropertiesResponseOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalWorkspacesPropertiesResponseOutput)
}





type AdditionalWorkspacesPropertiesResponseArrayInput interface {
	pulumi.Input

	ToAdditionalWorkspacesPropertiesResponseArrayOutput() AdditionalWorkspacesPropertiesResponseArrayOutput
	ToAdditionalWorkspacesPropertiesResponseArrayOutputWithContext(context.Context) AdditionalWorkspacesPropertiesResponseArrayOutput
}

type AdditionalWorkspacesPropertiesResponseArray []AdditionalWorkspacesPropertiesResponseInput

func (AdditionalWorkspacesPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalWorkspacesPropertiesResponse)(nil)).Elem()
}

func (i AdditionalWorkspacesPropertiesResponseArray) ToAdditionalWorkspacesPropertiesResponseArrayOutput() AdditionalWorkspacesPropertiesResponseArrayOutput {
	return i.ToAdditionalWorkspacesPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i AdditionalWorkspacesPropertiesResponseArray) ToAdditionalWorkspacesPropertiesResponseArrayOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalWorkspacesPropertiesResponseArrayOutput)
}

type AdditionalWorkspacesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspacesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspacesPropertiesResponse)(nil)).Elem()
}

func (o AdditionalWorkspacesPropertiesResponseOutput) ToAdditionalWorkspacesPropertiesResponseOutput() AdditionalWorkspacesPropertiesResponseOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesResponseOutput) ToAdditionalWorkspacesPropertiesResponseOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesResponseOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesResponseOutput) DataTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdditionalWorkspacesPropertiesResponse) []string { return v.DataTypes }).(pulumi.StringArrayOutput)
}

func (o AdditionalWorkspacesPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalWorkspacesPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o AdditionalWorkspacesPropertiesResponseOutput) Workspace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalWorkspacesPropertiesResponse) *string { return v.Workspace }).(pulumi.StringPtrOutput)
}

type AdditionalWorkspacesPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspacesPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalWorkspacesPropertiesResponse)(nil)).Elem()
}

func (o AdditionalWorkspacesPropertiesResponseArrayOutput) ToAdditionalWorkspacesPropertiesResponseArrayOutput() AdditionalWorkspacesPropertiesResponseArrayOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesResponseArrayOutput) ToAdditionalWorkspacesPropertiesResponseArrayOutputWithContext(ctx context.Context) AdditionalWorkspacesPropertiesResponseArrayOutput {
	return o
}

func (o AdditionalWorkspacesPropertiesResponseArrayOutput) Index(i pulumi.IntInput) AdditionalWorkspacesPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalWorkspacesPropertiesResponse {
		return vs[0].([]AdditionalWorkspacesPropertiesResponse)[vs[1].(int)]
	}).(AdditionalWorkspacesPropertiesResponseOutput)
}

type AllowlistCustomAlertRule struct {
	AllowlistValues []string `pulumi:"allowlistValues"`
	IsEnabled       bool     `pulumi:"isEnabled"`
	RuleType        string   `pulumi:"ruleType"`
}





type AllowlistCustomAlertRuleInput interface {
	pulumi.Input

	ToAllowlistCustomAlertRuleOutput() AllowlistCustomAlertRuleOutput
	ToAllowlistCustomAlertRuleOutputWithContext(context.Context) AllowlistCustomAlertRuleOutput
}

type AllowlistCustomAlertRuleArgs struct {
	AllowlistValues pulumi.StringArrayInput `pulumi:"allowlistValues"`
	IsEnabled       pulumi.BoolInput        `pulumi:"isEnabled"`
	RuleType        pulumi.StringInput      `pulumi:"ruleType"`
}

func (AllowlistCustomAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistCustomAlertRule)(nil)).Elem()
}

func (i AllowlistCustomAlertRuleArgs) ToAllowlistCustomAlertRuleOutput() AllowlistCustomAlertRuleOutput {
	return i.ToAllowlistCustomAlertRuleOutputWithContext(context.Background())
}

func (i AllowlistCustomAlertRuleArgs) ToAllowlistCustomAlertRuleOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistCustomAlertRuleOutput)
}





type AllowlistCustomAlertRuleArrayInput interface {
	pulumi.Input

	ToAllowlistCustomAlertRuleArrayOutput() AllowlistCustomAlertRuleArrayOutput
	ToAllowlistCustomAlertRuleArrayOutputWithContext(context.Context) AllowlistCustomAlertRuleArrayOutput
}

type AllowlistCustomAlertRuleArray []AllowlistCustomAlertRuleInput

func (AllowlistCustomAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowlistCustomAlertRule)(nil)).Elem()
}

func (i AllowlistCustomAlertRuleArray) ToAllowlistCustomAlertRuleArrayOutput() AllowlistCustomAlertRuleArrayOutput {
	return i.ToAllowlistCustomAlertRuleArrayOutputWithContext(context.Background())
}

func (i AllowlistCustomAlertRuleArray) ToAllowlistCustomAlertRuleArrayOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistCustomAlertRuleArrayOutput)
}

type AllowlistCustomAlertRuleOutput struct{ *pulumi.OutputState }

func (AllowlistCustomAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistCustomAlertRule)(nil)).Elem()
}

func (o AllowlistCustomAlertRuleOutput) ToAllowlistCustomAlertRuleOutput() AllowlistCustomAlertRuleOutput {
	return o
}

func (o AllowlistCustomAlertRuleOutput) ToAllowlistCustomAlertRuleOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleOutput {
	return o
}

func (o AllowlistCustomAlertRuleOutput) AllowlistValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRule) []string { return v.AllowlistValues }).(pulumi.StringArrayOutput)
}

func (o AllowlistCustomAlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRule) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o AllowlistCustomAlertRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type AllowlistCustomAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (AllowlistCustomAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowlistCustomAlertRule)(nil)).Elem()
}

func (o AllowlistCustomAlertRuleArrayOutput) ToAllowlistCustomAlertRuleArrayOutput() AllowlistCustomAlertRuleArrayOutput {
	return o
}

func (o AllowlistCustomAlertRuleArrayOutput) ToAllowlistCustomAlertRuleArrayOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleArrayOutput {
	return o
}

func (o AllowlistCustomAlertRuleArrayOutput) Index(i pulumi.IntInput) AllowlistCustomAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AllowlistCustomAlertRule {
		return vs[0].([]AllowlistCustomAlertRule)[vs[1].(int)]
	}).(AllowlistCustomAlertRuleOutput)
}

type AllowlistCustomAlertRuleResponse struct {
	AllowlistValues []string `pulumi:"allowlistValues"`
	Description     string   `pulumi:"description"`
	DisplayName     string   `pulumi:"displayName"`
	IsEnabled       bool     `pulumi:"isEnabled"`
	RuleType        string   `pulumi:"ruleType"`
	ValueType       string   `pulumi:"valueType"`
}





type AllowlistCustomAlertRuleResponseInput interface {
	pulumi.Input

	ToAllowlistCustomAlertRuleResponseOutput() AllowlistCustomAlertRuleResponseOutput
	ToAllowlistCustomAlertRuleResponseOutputWithContext(context.Context) AllowlistCustomAlertRuleResponseOutput
}

type AllowlistCustomAlertRuleResponseArgs struct {
	AllowlistValues pulumi.StringArrayInput `pulumi:"allowlistValues"`
	Description     pulumi.StringInput      `pulumi:"description"`
	DisplayName     pulumi.StringInput      `pulumi:"displayName"`
	IsEnabled       pulumi.BoolInput        `pulumi:"isEnabled"`
	RuleType        pulumi.StringInput      `pulumi:"ruleType"`
	ValueType       pulumi.StringInput      `pulumi:"valueType"`
}

func (AllowlistCustomAlertRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistCustomAlertRuleResponse)(nil)).Elem()
}

func (i AllowlistCustomAlertRuleResponseArgs) ToAllowlistCustomAlertRuleResponseOutput() AllowlistCustomAlertRuleResponseOutput {
	return i.ToAllowlistCustomAlertRuleResponseOutputWithContext(context.Background())
}

func (i AllowlistCustomAlertRuleResponseArgs) ToAllowlistCustomAlertRuleResponseOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistCustomAlertRuleResponseOutput)
}





type AllowlistCustomAlertRuleResponseArrayInput interface {
	pulumi.Input

	ToAllowlistCustomAlertRuleResponseArrayOutput() AllowlistCustomAlertRuleResponseArrayOutput
	ToAllowlistCustomAlertRuleResponseArrayOutputWithContext(context.Context) AllowlistCustomAlertRuleResponseArrayOutput
}

type AllowlistCustomAlertRuleResponseArray []AllowlistCustomAlertRuleResponseInput

func (AllowlistCustomAlertRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowlistCustomAlertRuleResponse)(nil)).Elem()
}

func (i AllowlistCustomAlertRuleResponseArray) ToAllowlistCustomAlertRuleResponseArrayOutput() AllowlistCustomAlertRuleResponseArrayOutput {
	return i.ToAllowlistCustomAlertRuleResponseArrayOutputWithContext(context.Background())
}

func (i AllowlistCustomAlertRuleResponseArray) ToAllowlistCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistCustomAlertRuleResponseArrayOutput)
}

type AllowlistCustomAlertRuleResponseOutput struct{ *pulumi.OutputState }

func (AllowlistCustomAlertRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistCustomAlertRuleResponse)(nil)).Elem()
}

func (o AllowlistCustomAlertRuleResponseOutput) ToAllowlistCustomAlertRuleResponseOutput() AllowlistCustomAlertRuleResponseOutput {
	return o
}

func (o AllowlistCustomAlertRuleResponseOutput) ToAllowlistCustomAlertRuleResponseOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleResponseOutput {
	return o
}

func (o AllowlistCustomAlertRuleResponseOutput) AllowlistValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) []string { return v.AllowlistValues }).(pulumi.StringArrayOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o AllowlistCustomAlertRuleResponseOutput) ValueType() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistCustomAlertRuleResponse) string { return v.ValueType }).(pulumi.StringOutput)
}

type AllowlistCustomAlertRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (AllowlistCustomAlertRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AllowlistCustomAlertRuleResponse)(nil)).Elem()
}

func (o AllowlistCustomAlertRuleResponseArrayOutput) ToAllowlistCustomAlertRuleResponseArrayOutput() AllowlistCustomAlertRuleResponseArrayOutput {
	return o
}

func (o AllowlistCustomAlertRuleResponseArrayOutput) ToAllowlistCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) AllowlistCustomAlertRuleResponseArrayOutput {
	return o
}

func (o AllowlistCustomAlertRuleResponseArrayOutput) Index(i pulumi.IntInput) AllowlistCustomAlertRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AllowlistCustomAlertRuleResponse {
		return vs[0].([]AllowlistCustomAlertRuleResponse)[vs[1].(int)]
	}).(AllowlistCustomAlertRuleResponseOutput)
}

type AssessmentLinksResponse struct {
	AzurePortalUri string `pulumi:"azurePortalUri"`
}





type AssessmentLinksResponseInput interface {
	pulumi.Input

	ToAssessmentLinksResponseOutput() AssessmentLinksResponseOutput
	ToAssessmentLinksResponseOutputWithContext(context.Context) AssessmentLinksResponseOutput
}

type AssessmentLinksResponseArgs struct {
	AzurePortalUri pulumi.StringInput `pulumi:"azurePortalUri"`
}

func (AssessmentLinksResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentLinksResponse)(nil)).Elem()
}

func (i AssessmentLinksResponseArgs) ToAssessmentLinksResponseOutput() AssessmentLinksResponseOutput {
	return i.ToAssessmentLinksResponseOutputWithContext(context.Background())
}

func (i AssessmentLinksResponseArgs) ToAssessmentLinksResponseOutputWithContext(ctx context.Context) AssessmentLinksResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentLinksResponseOutput)
}

func (i AssessmentLinksResponseArgs) ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput {
	return i.ToAssessmentLinksResponsePtrOutputWithContext(context.Background())
}

func (i AssessmentLinksResponseArgs) ToAssessmentLinksResponsePtrOutputWithContext(ctx context.Context) AssessmentLinksResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentLinksResponseOutput).ToAssessmentLinksResponsePtrOutputWithContext(ctx)
}









type AssessmentLinksResponsePtrInput interface {
	pulumi.Input

	ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput
	ToAssessmentLinksResponsePtrOutputWithContext(context.Context) AssessmentLinksResponsePtrOutput
}

type assessmentLinksResponsePtrType AssessmentLinksResponseArgs

func AssessmentLinksResponsePtr(v *AssessmentLinksResponseArgs) AssessmentLinksResponsePtrInput {
	return (*assessmentLinksResponsePtrType)(v)
}

func (*assessmentLinksResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentLinksResponse)(nil)).Elem()
}

func (i *assessmentLinksResponsePtrType) ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput {
	return i.ToAssessmentLinksResponsePtrOutputWithContext(context.Background())
}

func (i *assessmentLinksResponsePtrType) ToAssessmentLinksResponsePtrOutputWithContext(ctx context.Context) AssessmentLinksResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentLinksResponsePtrOutput)
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

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput {
	return o.ToAssessmentLinksResponsePtrOutputWithContext(context.Background())
}

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponsePtrOutputWithContext(ctx context.Context) AssessmentLinksResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentLinksResponse) *AssessmentLinksResponse {
		return &v
	}).(AssessmentLinksResponsePtrOutput)
}

func (o AssessmentLinksResponseOutput) AzurePortalUri() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentLinksResponse) string { return v.AzurePortalUri }).(pulumi.StringOutput)
}

type AssessmentLinksResponsePtrOutput struct{ *pulumi.OutputState }

func (AssessmentLinksResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentLinksResponse)(nil)).Elem()
}

func (o AssessmentLinksResponsePtrOutput) ToAssessmentLinksResponsePtrOutput() AssessmentLinksResponsePtrOutput {
	return o
}

func (o AssessmentLinksResponsePtrOutput) ToAssessmentLinksResponsePtrOutputWithContext(ctx context.Context) AssessmentLinksResponsePtrOutput {
	return o
}

func (o AssessmentLinksResponsePtrOutput) Elem() AssessmentLinksResponseOutput {
	return o.ApplyT(func(v *AssessmentLinksResponse) AssessmentLinksResponse {
		if v != nil {
			return *v
		}
		var ret AssessmentLinksResponse
		return ret
	}).(AssessmentLinksResponseOutput)
}

func (o AssessmentLinksResponsePtrOutput) AzurePortalUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentLinksResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzurePortalUri
	}).(pulumi.StringPtrOutput)
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

func (i AssessmentStatusArgs) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return i.ToAssessmentStatusPtrOutputWithContext(context.Background())
}

func (i AssessmentStatusArgs) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusOutput).ToAssessmentStatusPtrOutputWithContext(ctx)
}









type AssessmentStatusPtrInput interface {
	pulumi.Input

	ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput
	ToAssessmentStatusPtrOutputWithContext(context.Context) AssessmentStatusPtrOutput
}

type assessmentStatusPtrType AssessmentStatusArgs

func AssessmentStatusPtr(v *AssessmentStatusArgs) AssessmentStatusPtrInput {
	return (*assessmentStatusPtrType)(v)
}

func (*assessmentStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatus)(nil)).Elem()
}

func (i *assessmentStatusPtrType) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return i.ToAssessmentStatusPtrOutputWithContext(context.Background())
}

func (i *assessmentStatusPtrType) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusPtrOutput)
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

func (o AssessmentStatusOutput) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return o.ToAssessmentStatusPtrOutputWithContext(context.Background())
}

func (o AssessmentStatusOutput) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentStatus) *AssessmentStatus {
		return &v
	}).(AssessmentStatusPtrOutput)
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

type AssessmentStatusPtrOutput struct{ *pulumi.OutputState }

func (AssessmentStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatus)(nil)).Elem()
}

func (o AssessmentStatusPtrOutput) ToAssessmentStatusPtrOutput() AssessmentStatusPtrOutput {
	return o
}

func (o AssessmentStatusPtrOutput) ToAssessmentStatusPtrOutputWithContext(ctx context.Context) AssessmentStatusPtrOutput {
	return o
}

func (o AssessmentStatusPtrOutput) Elem() AssessmentStatusOutput {
	return o.ApplyT(func(v *AssessmentStatus) AssessmentStatus {
		if v != nil {
			return *v
		}
		var ret AssessmentStatus
		return ret
	}).(AssessmentStatusOutput)
}

func (o AssessmentStatusPtrOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatus) *string {
		if v == nil {
			return nil
		}
		return v.Cause
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatus) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

type AssessmentStatusResponse struct {
	Cause       *string `pulumi:"cause"`
	Code        string  `pulumi:"code"`
	Description *string `pulumi:"description"`
}





type AssessmentStatusResponseInput interface {
	pulumi.Input

	ToAssessmentStatusResponseOutput() AssessmentStatusResponseOutput
	ToAssessmentStatusResponseOutputWithContext(context.Context) AssessmentStatusResponseOutput
}

type AssessmentStatusResponseArgs struct {
	Cause       pulumi.StringPtrInput `pulumi:"cause"`
	Code        pulumi.StringInput    `pulumi:"code"`
	Description pulumi.StringPtrInput `pulumi:"description"`
}

func (AssessmentStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatusResponse)(nil)).Elem()
}

func (i AssessmentStatusResponseArgs) ToAssessmentStatusResponseOutput() AssessmentStatusResponseOutput {
	return i.ToAssessmentStatusResponseOutputWithContext(context.Background())
}

func (i AssessmentStatusResponseArgs) ToAssessmentStatusResponseOutputWithContext(ctx context.Context) AssessmentStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusResponseOutput)
}

func (i AssessmentStatusResponseArgs) ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput {
	return i.ToAssessmentStatusResponsePtrOutputWithContext(context.Background())
}

func (i AssessmentStatusResponseArgs) ToAssessmentStatusResponsePtrOutputWithContext(ctx context.Context) AssessmentStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusResponseOutput).ToAssessmentStatusResponsePtrOutputWithContext(ctx)
}









type AssessmentStatusResponsePtrInput interface {
	pulumi.Input

	ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput
	ToAssessmentStatusResponsePtrOutputWithContext(context.Context) AssessmentStatusResponsePtrOutput
}

type assessmentStatusResponsePtrType AssessmentStatusResponseArgs

func AssessmentStatusResponsePtr(v *AssessmentStatusResponseArgs) AssessmentStatusResponsePtrInput {
	return (*assessmentStatusResponsePtrType)(v)
}

func (*assessmentStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatusResponse)(nil)).Elem()
}

func (i *assessmentStatusResponsePtrType) ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput {
	return i.ToAssessmentStatusResponsePtrOutputWithContext(context.Background())
}

func (i *assessmentStatusResponsePtrType) ToAssessmentStatusResponsePtrOutputWithContext(ctx context.Context) AssessmentStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusResponsePtrOutput)
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

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput {
	return o.ToAssessmentStatusResponsePtrOutputWithContext(context.Background())
}

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponsePtrOutputWithContext(ctx context.Context) AssessmentStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentStatusResponse) *AssessmentStatusResponse {
		return &v
	}).(AssessmentStatusResponsePtrOutput)
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

type AssessmentStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (AssessmentStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStatusResponse)(nil)).Elem()
}

func (o AssessmentStatusResponsePtrOutput) ToAssessmentStatusResponsePtrOutput() AssessmentStatusResponsePtrOutput {
	return o
}

func (o AssessmentStatusResponsePtrOutput) ToAssessmentStatusResponsePtrOutputWithContext(ctx context.Context) AssessmentStatusResponsePtrOutput {
	return o
}

func (o AssessmentStatusResponsePtrOutput) Elem() AssessmentStatusResponseOutput {
	return o.ApplyT(func(v *AssessmentStatusResponse) AssessmentStatusResponse {
		if v != nil {
			return *v
		}
		var ret AssessmentStatusResponse
		return ret
	}).(AssessmentStatusResponseOutput)
}

func (o AssessmentStatusResponsePtrOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cause
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

type AssignedComponentItem struct {
	Key *string `pulumi:"key"`
}





type AssignedComponentItemInput interface {
	pulumi.Input

	ToAssignedComponentItemOutput() AssignedComponentItemOutput
	ToAssignedComponentItemOutputWithContext(context.Context) AssignedComponentItemOutput
}

type AssignedComponentItemArgs struct {
	Key pulumi.StringPtrInput `pulumi:"key"`
}

func (AssignedComponentItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedComponentItem)(nil)).Elem()
}

func (i AssignedComponentItemArgs) ToAssignedComponentItemOutput() AssignedComponentItemOutput {
	return i.ToAssignedComponentItemOutputWithContext(context.Background())
}

func (i AssignedComponentItemArgs) ToAssignedComponentItemOutputWithContext(ctx context.Context) AssignedComponentItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedComponentItemOutput)
}

func (i AssignedComponentItemArgs) ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput {
	return i.ToAssignedComponentItemPtrOutputWithContext(context.Background())
}

func (i AssignedComponentItemArgs) ToAssignedComponentItemPtrOutputWithContext(ctx context.Context) AssignedComponentItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedComponentItemOutput).ToAssignedComponentItemPtrOutputWithContext(ctx)
}









type AssignedComponentItemPtrInput interface {
	pulumi.Input

	ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput
	ToAssignedComponentItemPtrOutputWithContext(context.Context) AssignedComponentItemPtrOutput
}

type assignedComponentItemPtrType AssignedComponentItemArgs

func AssignedComponentItemPtr(v *AssignedComponentItemArgs) AssignedComponentItemPtrInput {
	return (*assignedComponentItemPtrType)(v)
}

func (*assignedComponentItemPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedComponentItem)(nil)).Elem()
}

func (i *assignedComponentItemPtrType) ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput {
	return i.ToAssignedComponentItemPtrOutputWithContext(context.Background())
}

func (i *assignedComponentItemPtrType) ToAssignedComponentItemPtrOutputWithContext(ctx context.Context) AssignedComponentItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedComponentItemPtrOutput)
}

type AssignedComponentItemOutput struct{ *pulumi.OutputState }

func (AssignedComponentItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedComponentItem)(nil)).Elem()
}

func (o AssignedComponentItemOutput) ToAssignedComponentItemOutput() AssignedComponentItemOutput {
	return o
}

func (o AssignedComponentItemOutput) ToAssignedComponentItemOutputWithContext(ctx context.Context) AssignedComponentItemOutput {
	return o
}

func (o AssignedComponentItemOutput) ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput {
	return o.ToAssignedComponentItemPtrOutputWithContext(context.Background())
}

func (o AssignedComponentItemOutput) ToAssignedComponentItemPtrOutputWithContext(ctx context.Context) AssignedComponentItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignedComponentItem) *AssignedComponentItem {
		return &v
	}).(AssignedComponentItemPtrOutput)
}

func (o AssignedComponentItemOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignedComponentItem) *string { return v.Key }).(pulumi.StringPtrOutput)
}

type AssignedComponentItemPtrOutput struct{ *pulumi.OutputState }

func (AssignedComponentItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedComponentItem)(nil)).Elem()
}

func (o AssignedComponentItemPtrOutput) ToAssignedComponentItemPtrOutput() AssignedComponentItemPtrOutput {
	return o
}

func (o AssignedComponentItemPtrOutput) ToAssignedComponentItemPtrOutputWithContext(ctx context.Context) AssignedComponentItemPtrOutput {
	return o
}

func (o AssignedComponentItemPtrOutput) Elem() AssignedComponentItemOutput {
	return o.ApplyT(func(v *AssignedComponentItem) AssignedComponentItem {
		if v != nil {
			return *v
		}
		var ret AssignedComponentItem
		return ret
	}).(AssignedComponentItemOutput)
}

func (o AssignedComponentItemPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedComponentItem) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

type AssignedComponentItemResponse struct {
	Key *string `pulumi:"key"`
}





type AssignedComponentItemResponseInput interface {
	pulumi.Input

	ToAssignedComponentItemResponseOutput() AssignedComponentItemResponseOutput
	ToAssignedComponentItemResponseOutputWithContext(context.Context) AssignedComponentItemResponseOutput
}

type AssignedComponentItemResponseArgs struct {
	Key pulumi.StringPtrInput `pulumi:"key"`
}

func (AssignedComponentItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedComponentItemResponse)(nil)).Elem()
}

func (i AssignedComponentItemResponseArgs) ToAssignedComponentItemResponseOutput() AssignedComponentItemResponseOutput {
	return i.ToAssignedComponentItemResponseOutputWithContext(context.Background())
}

func (i AssignedComponentItemResponseArgs) ToAssignedComponentItemResponseOutputWithContext(ctx context.Context) AssignedComponentItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedComponentItemResponseOutput)
}

func (i AssignedComponentItemResponseArgs) ToAssignedComponentItemResponsePtrOutput() AssignedComponentItemResponsePtrOutput {
	return i.ToAssignedComponentItemResponsePtrOutputWithContext(context.Background())
}

func (i AssignedComponentItemResponseArgs) ToAssignedComponentItemResponsePtrOutputWithContext(ctx context.Context) AssignedComponentItemResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedComponentItemResponseOutput).ToAssignedComponentItemResponsePtrOutputWithContext(ctx)
}









type AssignedComponentItemResponsePtrInput interface {
	pulumi.Input

	ToAssignedComponentItemResponsePtrOutput() AssignedComponentItemResponsePtrOutput
	ToAssignedComponentItemResponsePtrOutputWithContext(context.Context) AssignedComponentItemResponsePtrOutput
}

type assignedComponentItemResponsePtrType AssignedComponentItemResponseArgs

func AssignedComponentItemResponsePtr(v *AssignedComponentItemResponseArgs) AssignedComponentItemResponsePtrInput {
	return (*assignedComponentItemResponsePtrType)(v)
}

func (*assignedComponentItemResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedComponentItemResponse)(nil)).Elem()
}

func (i *assignedComponentItemResponsePtrType) ToAssignedComponentItemResponsePtrOutput() AssignedComponentItemResponsePtrOutput {
	return i.ToAssignedComponentItemResponsePtrOutputWithContext(context.Background())
}

func (i *assignedComponentItemResponsePtrType) ToAssignedComponentItemResponsePtrOutputWithContext(ctx context.Context) AssignedComponentItemResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedComponentItemResponsePtrOutput)
}

type AssignedComponentItemResponseOutput struct{ *pulumi.OutputState }

func (AssignedComponentItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedComponentItemResponse)(nil)).Elem()
}

func (o AssignedComponentItemResponseOutput) ToAssignedComponentItemResponseOutput() AssignedComponentItemResponseOutput {
	return o
}

func (o AssignedComponentItemResponseOutput) ToAssignedComponentItemResponseOutputWithContext(ctx context.Context) AssignedComponentItemResponseOutput {
	return o
}

func (o AssignedComponentItemResponseOutput) ToAssignedComponentItemResponsePtrOutput() AssignedComponentItemResponsePtrOutput {
	return o.ToAssignedComponentItemResponsePtrOutputWithContext(context.Background())
}

func (o AssignedComponentItemResponseOutput) ToAssignedComponentItemResponsePtrOutputWithContext(ctx context.Context) AssignedComponentItemResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignedComponentItemResponse) *AssignedComponentItemResponse {
		return &v
	}).(AssignedComponentItemResponsePtrOutput)
}

func (o AssignedComponentItemResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignedComponentItemResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

type AssignedComponentItemResponsePtrOutput struct{ *pulumi.OutputState }

func (AssignedComponentItemResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedComponentItemResponse)(nil)).Elem()
}

func (o AssignedComponentItemResponsePtrOutput) ToAssignedComponentItemResponsePtrOutput() AssignedComponentItemResponsePtrOutput {
	return o
}

func (o AssignedComponentItemResponsePtrOutput) ToAssignedComponentItemResponsePtrOutputWithContext(ctx context.Context) AssignedComponentItemResponsePtrOutput {
	return o
}

func (o AssignedComponentItemResponsePtrOutput) Elem() AssignedComponentItemResponseOutput {
	return o.ApplyT(func(v *AssignedComponentItemResponse) AssignedComponentItemResponse {
		if v != nil {
			return *v
		}
		var ret AssignedComponentItemResponse
		return ret
	}).(AssignedComponentItemResponseOutput)
}

func (o AssignedComponentItemResponsePtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedComponentItemResponse) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

type AssignedStandardItem struct {
	Id *string `pulumi:"id"`
}





type AssignedStandardItemInput interface {
	pulumi.Input

	ToAssignedStandardItemOutput() AssignedStandardItemOutput
	ToAssignedStandardItemOutputWithContext(context.Context) AssignedStandardItemOutput
}

type AssignedStandardItemArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (AssignedStandardItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedStandardItem)(nil)).Elem()
}

func (i AssignedStandardItemArgs) ToAssignedStandardItemOutput() AssignedStandardItemOutput {
	return i.ToAssignedStandardItemOutputWithContext(context.Background())
}

func (i AssignedStandardItemArgs) ToAssignedStandardItemOutputWithContext(ctx context.Context) AssignedStandardItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedStandardItemOutput)
}

func (i AssignedStandardItemArgs) ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput {
	return i.ToAssignedStandardItemPtrOutputWithContext(context.Background())
}

func (i AssignedStandardItemArgs) ToAssignedStandardItemPtrOutputWithContext(ctx context.Context) AssignedStandardItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedStandardItemOutput).ToAssignedStandardItemPtrOutputWithContext(ctx)
}









type AssignedStandardItemPtrInput interface {
	pulumi.Input

	ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput
	ToAssignedStandardItemPtrOutputWithContext(context.Context) AssignedStandardItemPtrOutput
}

type assignedStandardItemPtrType AssignedStandardItemArgs

func AssignedStandardItemPtr(v *AssignedStandardItemArgs) AssignedStandardItemPtrInput {
	return (*assignedStandardItemPtrType)(v)
}

func (*assignedStandardItemPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedStandardItem)(nil)).Elem()
}

func (i *assignedStandardItemPtrType) ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput {
	return i.ToAssignedStandardItemPtrOutputWithContext(context.Background())
}

func (i *assignedStandardItemPtrType) ToAssignedStandardItemPtrOutputWithContext(ctx context.Context) AssignedStandardItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedStandardItemPtrOutput)
}

type AssignedStandardItemOutput struct{ *pulumi.OutputState }

func (AssignedStandardItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedStandardItem)(nil)).Elem()
}

func (o AssignedStandardItemOutput) ToAssignedStandardItemOutput() AssignedStandardItemOutput {
	return o
}

func (o AssignedStandardItemOutput) ToAssignedStandardItemOutputWithContext(ctx context.Context) AssignedStandardItemOutput {
	return o
}

func (o AssignedStandardItemOutput) ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput {
	return o.ToAssignedStandardItemPtrOutputWithContext(context.Background())
}

func (o AssignedStandardItemOutput) ToAssignedStandardItemPtrOutputWithContext(ctx context.Context) AssignedStandardItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignedStandardItem) *AssignedStandardItem {
		return &v
	}).(AssignedStandardItemPtrOutput)
}

func (o AssignedStandardItemOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignedStandardItem) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type AssignedStandardItemPtrOutput struct{ *pulumi.OutputState }

func (AssignedStandardItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedStandardItem)(nil)).Elem()
}

func (o AssignedStandardItemPtrOutput) ToAssignedStandardItemPtrOutput() AssignedStandardItemPtrOutput {
	return o
}

func (o AssignedStandardItemPtrOutput) ToAssignedStandardItemPtrOutputWithContext(ctx context.Context) AssignedStandardItemPtrOutput {
	return o
}

func (o AssignedStandardItemPtrOutput) Elem() AssignedStandardItemOutput {
	return o.ApplyT(func(v *AssignedStandardItem) AssignedStandardItem {
		if v != nil {
			return *v
		}
		var ret AssignedStandardItem
		return ret
	}).(AssignedStandardItemOutput)
}

func (o AssignedStandardItemPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedStandardItem) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type AssignedStandardItemResponse struct {
	Id *string `pulumi:"id"`
}





type AssignedStandardItemResponseInput interface {
	pulumi.Input

	ToAssignedStandardItemResponseOutput() AssignedStandardItemResponseOutput
	ToAssignedStandardItemResponseOutputWithContext(context.Context) AssignedStandardItemResponseOutput
}

type AssignedStandardItemResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (AssignedStandardItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedStandardItemResponse)(nil)).Elem()
}

func (i AssignedStandardItemResponseArgs) ToAssignedStandardItemResponseOutput() AssignedStandardItemResponseOutput {
	return i.ToAssignedStandardItemResponseOutputWithContext(context.Background())
}

func (i AssignedStandardItemResponseArgs) ToAssignedStandardItemResponseOutputWithContext(ctx context.Context) AssignedStandardItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedStandardItemResponseOutput)
}

func (i AssignedStandardItemResponseArgs) ToAssignedStandardItemResponsePtrOutput() AssignedStandardItemResponsePtrOutput {
	return i.ToAssignedStandardItemResponsePtrOutputWithContext(context.Background())
}

func (i AssignedStandardItemResponseArgs) ToAssignedStandardItemResponsePtrOutputWithContext(ctx context.Context) AssignedStandardItemResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedStandardItemResponseOutput).ToAssignedStandardItemResponsePtrOutputWithContext(ctx)
}









type AssignedStandardItemResponsePtrInput interface {
	pulumi.Input

	ToAssignedStandardItemResponsePtrOutput() AssignedStandardItemResponsePtrOutput
	ToAssignedStandardItemResponsePtrOutputWithContext(context.Context) AssignedStandardItemResponsePtrOutput
}

type assignedStandardItemResponsePtrType AssignedStandardItemResponseArgs

func AssignedStandardItemResponsePtr(v *AssignedStandardItemResponseArgs) AssignedStandardItemResponsePtrInput {
	return (*assignedStandardItemResponsePtrType)(v)
}

func (*assignedStandardItemResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedStandardItemResponse)(nil)).Elem()
}

func (i *assignedStandardItemResponsePtrType) ToAssignedStandardItemResponsePtrOutput() AssignedStandardItemResponsePtrOutput {
	return i.ToAssignedStandardItemResponsePtrOutputWithContext(context.Background())
}

func (i *assignedStandardItemResponsePtrType) ToAssignedStandardItemResponsePtrOutputWithContext(ctx context.Context) AssignedStandardItemResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignedStandardItemResponsePtrOutput)
}

type AssignedStandardItemResponseOutput struct{ *pulumi.OutputState }

func (AssignedStandardItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignedStandardItemResponse)(nil)).Elem()
}

func (o AssignedStandardItemResponseOutput) ToAssignedStandardItemResponseOutput() AssignedStandardItemResponseOutput {
	return o
}

func (o AssignedStandardItemResponseOutput) ToAssignedStandardItemResponseOutputWithContext(ctx context.Context) AssignedStandardItemResponseOutput {
	return o
}

func (o AssignedStandardItemResponseOutput) ToAssignedStandardItemResponsePtrOutput() AssignedStandardItemResponsePtrOutput {
	return o.ToAssignedStandardItemResponsePtrOutputWithContext(context.Background())
}

func (o AssignedStandardItemResponseOutput) ToAssignedStandardItemResponsePtrOutputWithContext(ctx context.Context) AssignedStandardItemResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignedStandardItemResponse) *AssignedStandardItemResponse {
		return &v
	}).(AssignedStandardItemResponsePtrOutput)
}

func (o AssignedStandardItemResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignedStandardItemResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type AssignedStandardItemResponsePtrOutput struct{ *pulumi.OutputState }

func (AssignedStandardItemResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignedStandardItemResponse)(nil)).Elem()
}

func (o AssignedStandardItemResponsePtrOutput) ToAssignedStandardItemResponsePtrOutput() AssignedStandardItemResponsePtrOutput {
	return o
}

func (o AssignedStandardItemResponsePtrOutput) ToAssignedStandardItemResponsePtrOutputWithContext(ctx context.Context) AssignedStandardItemResponsePtrOutput {
	return o
}

func (o AssignedStandardItemResponsePtrOutput) Elem() AssignedStandardItemResponseOutput {
	return o.ApplyT(func(v *AssignedStandardItemResponse) AssignedStandardItemResponse {
		if v != nil {
			return *v
		}
		var ret AssignedStandardItemResponse
		return ret
	}).(AssignedStandardItemResponseOutput)
}

func (o AssignedStandardItemResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignedStandardItemResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type AssignmentPropertiesAdditionalData struct {
	ExemptionCategory *string `pulumi:"exemptionCategory"`
}





type AssignmentPropertiesAdditionalDataInput interface {
	pulumi.Input

	ToAssignmentPropertiesAdditionalDataOutput() AssignmentPropertiesAdditionalDataOutput
	ToAssignmentPropertiesAdditionalDataOutputWithContext(context.Context) AssignmentPropertiesAdditionalDataOutput
}

type AssignmentPropertiesAdditionalDataArgs struct {
	ExemptionCategory pulumi.StringPtrInput `pulumi:"exemptionCategory"`
}

func (AssignmentPropertiesAdditionalDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPropertiesAdditionalData)(nil)).Elem()
}

func (i AssignmentPropertiesAdditionalDataArgs) ToAssignmentPropertiesAdditionalDataOutput() AssignmentPropertiesAdditionalDataOutput {
	return i.ToAssignmentPropertiesAdditionalDataOutputWithContext(context.Background())
}

func (i AssignmentPropertiesAdditionalDataArgs) ToAssignmentPropertiesAdditionalDataOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPropertiesAdditionalDataOutput)
}

func (i AssignmentPropertiesAdditionalDataArgs) ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput {
	return i.ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(context.Background())
}

func (i AssignmentPropertiesAdditionalDataArgs) ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPropertiesAdditionalDataOutput).ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx)
}









type AssignmentPropertiesAdditionalDataPtrInput interface {
	pulumi.Input

	ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput
	ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(context.Context) AssignmentPropertiesAdditionalDataPtrOutput
}

type assignmentPropertiesAdditionalDataPtrType AssignmentPropertiesAdditionalDataArgs

func AssignmentPropertiesAdditionalDataPtr(v *AssignmentPropertiesAdditionalDataArgs) AssignmentPropertiesAdditionalDataPtrInput {
	return (*assignmentPropertiesAdditionalDataPtrType)(v)
}

func (*assignmentPropertiesAdditionalDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentPropertiesAdditionalData)(nil)).Elem()
}

func (i *assignmentPropertiesAdditionalDataPtrType) ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput {
	return i.ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(context.Background())
}

func (i *assignmentPropertiesAdditionalDataPtrType) ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPropertiesAdditionalDataPtrOutput)
}

type AssignmentPropertiesAdditionalDataOutput struct{ *pulumi.OutputState }

func (AssignmentPropertiesAdditionalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPropertiesAdditionalData)(nil)).Elem()
}

func (o AssignmentPropertiesAdditionalDataOutput) ToAssignmentPropertiesAdditionalDataOutput() AssignmentPropertiesAdditionalDataOutput {
	return o
}

func (o AssignmentPropertiesAdditionalDataOutput) ToAssignmentPropertiesAdditionalDataOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataOutput {
	return o
}

func (o AssignmentPropertiesAdditionalDataOutput) ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput {
	return o.ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(context.Background())
}

func (o AssignmentPropertiesAdditionalDataOutput) ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignmentPropertiesAdditionalData) *AssignmentPropertiesAdditionalData {
		return &v
	}).(AssignmentPropertiesAdditionalDataPtrOutput)
}

func (o AssignmentPropertiesAdditionalDataOutput) ExemptionCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignmentPropertiesAdditionalData) *string { return v.ExemptionCategory }).(pulumi.StringPtrOutput)
}

type AssignmentPropertiesAdditionalDataPtrOutput struct{ *pulumi.OutputState }

func (AssignmentPropertiesAdditionalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentPropertiesAdditionalData)(nil)).Elem()
}

func (o AssignmentPropertiesAdditionalDataPtrOutput) ToAssignmentPropertiesAdditionalDataPtrOutput() AssignmentPropertiesAdditionalDataPtrOutput {
	return o
}

func (o AssignmentPropertiesAdditionalDataPtrOutput) ToAssignmentPropertiesAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesAdditionalDataPtrOutput {
	return o
}

func (o AssignmentPropertiesAdditionalDataPtrOutput) Elem() AssignmentPropertiesAdditionalDataOutput {
	return o.ApplyT(func(v *AssignmentPropertiesAdditionalData) AssignmentPropertiesAdditionalData {
		if v != nil {
			return *v
		}
		var ret AssignmentPropertiesAdditionalData
		return ret
	}).(AssignmentPropertiesAdditionalDataOutput)
}

func (o AssignmentPropertiesAdditionalDataPtrOutput) ExemptionCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentPropertiesAdditionalData) *string {
		if v == nil {
			return nil
		}
		return v.ExemptionCategory
	}).(pulumi.StringPtrOutput)
}

type AssignmentPropertiesResponseAdditionalData struct {
	ExemptionCategory *string `pulumi:"exemptionCategory"`
}





type AssignmentPropertiesResponseAdditionalDataInput interface {
	pulumi.Input

	ToAssignmentPropertiesResponseAdditionalDataOutput() AssignmentPropertiesResponseAdditionalDataOutput
	ToAssignmentPropertiesResponseAdditionalDataOutputWithContext(context.Context) AssignmentPropertiesResponseAdditionalDataOutput
}

type AssignmentPropertiesResponseAdditionalDataArgs struct {
	ExemptionCategory pulumi.StringPtrInput `pulumi:"exemptionCategory"`
}

func (AssignmentPropertiesResponseAdditionalDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPropertiesResponseAdditionalData)(nil)).Elem()
}

func (i AssignmentPropertiesResponseAdditionalDataArgs) ToAssignmentPropertiesResponseAdditionalDataOutput() AssignmentPropertiesResponseAdditionalDataOutput {
	return i.ToAssignmentPropertiesResponseAdditionalDataOutputWithContext(context.Background())
}

func (i AssignmentPropertiesResponseAdditionalDataArgs) ToAssignmentPropertiesResponseAdditionalDataOutputWithContext(ctx context.Context) AssignmentPropertiesResponseAdditionalDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPropertiesResponseAdditionalDataOutput)
}

func (i AssignmentPropertiesResponseAdditionalDataArgs) ToAssignmentPropertiesResponseAdditionalDataPtrOutput() AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return i.ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(context.Background())
}

func (i AssignmentPropertiesResponseAdditionalDataArgs) ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPropertiesResponseAdditionalDataOutput).ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(ctx)
}









type AssignmentPropertiesResponseAdditionalDataPtrInput interface {
	pulumi.Input

	ToAssignmentPropertiesResponseAdditionalDataPtrOutput() AssignmentPropertiesResponseAdditionalDataPtrOutput
	ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(context.Context) AssignmentPropertiesResponseAdditionalDataPtrOutput
}

type assignmentPropertiesResponseAdditionalDataPtrType AssignmentPropertiesResponseAdditionalDataArgs

func AssignmentPropertiesResponseAdditionalDataPtr(v *AssignmentPropertiesResponseAdditionalDataArgs) AssignmentPropertiesResponseAdditionalDataPtrInput {
	return (*assignmentPropertiesResponseAdditionalDataPtrType)(v)
}

func (*assignmentPropertiesResponseAdditionalDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentPropertiesResponseAdditionalData)(nil)).Elem()
}

func (i *assignmentPropertiesResponseAdditionalDataPtrType) ToAssignmentPropertiesResponseAdditionalDataPtrOutput() AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return i.ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(context.Background())
}

func (i *assignmentPropertiesResponseAdditionalDataPtrType) ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentPropertiesResponseAdditionalDataPtrOutput)
}

type AssignmentPropertiesResponseAdditionalDataOutput struct{ *pulumi.OutputState }

func (AssignmentPropertiesResponseAdditionalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentPropertiesResponseAdditionalData)(nil)).Elem()
}

func (o AssignmentPropertiesResponseAdditionalDataOutput) ToAssignmentPropertiesResponseAdditionalDataOutput() AssignmentPropertiesResponseAdditionalDataOutput {
	return o
}

func (o AssignmentPropertiesResponseAdditionalDataOutput) ToAssignmentPropertiesResponseAdditionalDataOutputWithContext(ctx context.Context) AssignmentPropertiesResponseAdditionalDataOutput {
	return o
}

func (o AssignmentPropertiesResponseAdditionalDataOutput) ToAssignmentPropertiesResponseAdditionalDataPtrOutput() AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return o.ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(context.Background())
}

func (o AssignmentPropertiesResponseAdditionalDataOutput) ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignmentPropertiesResponseAdditionalData) *AssignmentPropertiesResponseAdditionalData {
		return &v
	}).(AssignmentPropertiesResponseAdditionalDataPtrOutput)
}

func (o AssignmentPropertiesResponseAdditionalDataOutput) ExemptionCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignmentPropertiesResponseAdditionalData) *string { return v.ExemptionCategory }).(pulumi.StringPtrOutput)
}

type AssignmentPropertiesResponseAdditionalDataPtrOutput struct{ *pulumi.OutputState }

func (AssignmentPropertiesResponseAdditionalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentPropertiesResponseAdditionalData)(nil)).Elem()
}

func (o AssignmentPropertiesResponseAdditionalDataPtrOutput) ToAssignmentPropertiesResponseAdditionalDataPtrOutput() AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return o
}

func (o AssignmentPropertiesResponseAdditionalDataPtrOutput) ToAssignmentPropertiesResponseAdditionalDataPtrOutputWithContext(ctx context.Context) AssignmentPropertiesResponseAdditionalDataPtrOutput {
	return o
}

func (o AssignmentPropertiesResponseAdditionalDataPtrOutput) Elem() AssignmentPropertiesResponseAdditionalDataOutput {
	return o.ApplyT(func(v *AssignmentPropertiesResponseAdditionalData) AssignmentPropertiesResponseAdditionalData {
		if v != nil {
			return *v
		}
		var ret AssignmentPropertiesResponseAdditionalData
		return ret
	}).(AssignmentPropertiesResponseAdditionalDataOutput)
}

func (o AssignmentPropertiesResponseAdditionalDataPtrOutput) ExemptionCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentPropertiesResponseAdditionalData) *string {
		if v == nil {
			return nil
		}
		return v.ExemptionCategory
	}).(pulumi.StringPtrOutput)
}

type AutomationActionEventHub struct {
	ActionType         string  `pulumi:"actionType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EventHubResourceId *string `pulumi:"eventHubResourceId"`
}





type AutomationActionEventHubInput interface {
	pulumi.Input

	ToAutomationActionEventHubOutput() AutomationActionEventHubOutput
	ToAutomationActionEventHubOutputWithContext(context.Context) AutomationActionEventHubOutput
}

type AutomationActionEventHubArgs struct {
	ActionType         pulumi.StringInput    `pulumi:"actionType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	EventHubResourceId pulumi.StringPtrInput `pulumi:"eventHubResourceId"`
}

func (AutomationActionEventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionEventHub)(nil)).Elem()
}

func (i AutomationActionEventHubArgs) ToAutomationActionEventHubOutput() AutomationActionEventHubOutput {
	return i.ToAutomationActionEventHubOutputWithContext(context.Background())
}

func (i AutomationActionEventHubArgs) ToAutomationActionEventHubOutputWithContext(ctx context.Context) AutomationActionEventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionEventHubOutput)
}

type AutomationActionEventHubOutput struct{ *pulumi.OutputState }

func (AutomationActionEventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionEventHub)(nil)).Elem()
}

func (o AutomationActionEventHubOutput) ToAutomationActionEventHubOutput() AutomationActionEventHubOutput {
	return o
}

func (o AutomationActionEventHubOutput) ToAutomationActionEventHubOutputWithContext(ctx context.Context) AutomationActionEventHubOutput {
	return o
}

func (o AutomationActionEventHubOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionEventHub) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionEventHubOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionEventHub) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o AutomationActionEventHubOutput) EventHubResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionEventHub) *string { return v.EventHubResourceId }).(pulumi.StringPtrOutput)
}

type AutomationActionEventHubResponse struct {
	ActionType         string  `pulumi:"actionType"`
	ConnectionString   *string `pulumi:"connectionString"`
	EventHubResourceId *string `pulumi:"eventHubResourceId"`
	SasPolicyName      string  `pulumi:"sasPolicyName"`
}





type AutomationActionEventHubResponseInput interface {
	pulumi.Input

	ToAutomationActionEventHubResponseOutput() AutomationActionEventHubResponseOutput
	ToAutomationActionEventHubResponseOutputWithContext(context.Context) AutomationActionEventHubResponseOutput
}

type AutomationActionEventHubResponseArgs struct {
	ActionType         pulumi.StringInput    `pulumi:"actionType"`
	ConnectionString   pulumi.StringPtrInput `pulumi:"connectionString"`
	EventHubResourceId pulumi.StringPtrInput `pulumi:"eventHubResourceId"`
	SasPolicyName      pulumi.StringInput    `pulumi:"sasPolicyName"`
}

func (AutomationActionEventHubResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionEventHubResponse)(nil)).Elem()
}

func (i AutomationActionEventHubResponseArgs) ToAutomationActionEventHubResponseOutput() AutomationActionEventHubResponseOutput {
	return i.ToAutomationActionEventHubResponseOutputWithContext(context.Background())
}

func (i AutomationActionEventHubResponseArgs) ToAutomationActionEventHubResponseOutputWithContext(ctx context.Context) AutomationActionEventHubResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionEventHubResponseOutput)
}

type AutomationActionEventHubResponseOutput struct{ *pulumi.OutputState }

func (AutomationActionEventHubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionEventHubResponse)(nil)).Elem()
}

func (o AutomationActionEventHubResponseOutput) ToAutomationActionEventHubResponseOutput() AutomationActionEventHubResponseOutput {
	return o
}

func (o AutomationActionEventHubResponseOutput) ToAutomationActionEventHubResponseOutputWithContext(ctx context.Context) AutomationActionEventHubResponseOutput {
	return o
}

func (o AutomationActionEventHubResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionEventHubResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionEventHubResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionEventHubResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o AutomationActionEventHubResponseOutput) EventHubResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionEventHubResponse) *string { return v.EventHubResourceId }).(pulumi.StringPtrOutput)
}

func (o AutomationActionEventHubResponseOutput) SasPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionEventHubResponse) string { return v.SasPolicyName }).(pulumi.StringOutput)
}

type AutomationActionLogicApp struct {
	ActionType         string  `pulumi:"actionType"`
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	Uri                *string `pulumi:"uri"`
}





type AutomationActionLogicAppInput interface {
	pulumi.Input

	ToAutomationActionLogicAppOutput() AutomationActionLogicAppOutput
	ToAutomationActionLogicAppOutputWithContext(context.Context) AutomationActionLogicAppOutput
}

type AutomationActionLogicAppArgs struct {
	ActionType         pulumi.StringInput    `pulumi:"actionType"`
	LogicAppResourceId pulumi.StringPtrInput `pulumi:"logicAppResourceId"`
	Uri                pulumi.StringPtrInput `pulumi:"uri"`
}

func (AutomationActionLogicAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionLogicApp)(nil)).Elem()
}

func (i AutomationActionLogicAppArgs) ToAutomationActionLogicAppOutput() AutomationActionLogicAppOutput {
	return i.ToAutomationActionLogicAppOutputWithContext(context.Background())
}

func (i AutomationActionLogicAppArgs) ToAutomationActionLogicAppOutputWithContext(ctx context.Context) AutomationActionLogicAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionLogicAppOutput)
}

type AutomationActionLogicAppOutput struct{ *pulumi.OutputState }

func (AutomationActionLogicAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionLogicApp)(nil)).Elem()
}

func (o AutomationActionLogicAppOutput) ToAutomationActionLogicAppOutput() AutomationActionLogicAppOutput {
	return o
}

func (o AutomationActionLogicAppOutput) ToAutomationActionLogicAppOutputWithContext(ctx context.Context) AutomationActionLogicAppOutput {
	return o
}

func (o AutomationActionLogicAppOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionLogicApp) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionLogicAppOutput) LogicAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionLogicApp) *string { return v.LogicAppResourceId }).(pulumi.StringPtrOutput)
}

func (o AutomationActionLogicAppOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionLogicApp) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type AutomationActionLogicAppResponse struct {
	ActionType         string  `pulumi:"actionType"`
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	Uri                *string `pulumi:"uri"`
}





type AutomationActionLogicAppResponseInput interface {
	pulumi.Input

	ToAutomationActionLogicAppResponseOutput() AutomationActionLogicAppResponseOutput
	ToAutomationActionLogicAppResponseOutputWithContext(context.Context) AutomationActionLogicAppResponseOutput
}

type AutomationActionLogicAppResponseArgs struct {
	ActionType         pulumi.StringInput    `pulumi:"actionType"`
	LogicAppResourceId pulumi.StringPtrInput `pulumi:"logicAppResourceId"`
	Uri                pulumi.StringPtrInput `pulumi:"uri"`
}

func (AutomationActionLogicAppResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionLogicAppResponse)(nil)).Elem()
}

func (i AutomationActionLogicAppResponseArgs) ToAutomationActionLogicAppResponseOutput() AutomationActionLogicAppResponseOutput {
	return i.ToAutomationActionLogicAppResponseOutputWithContext(context.Background())
}

func (i AutomationActionLogicAppResponseArgs) ToAutomationActionLogicAppResponseOutputWithContext(ctx context.Context) AutomationActionLogicAppResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionLogicAppResponseOutput)
}

type AutomationActionLogicAppResponseOutput struct{ *pulumi.OutputState }

func (AutomationActionLogicAppResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionLogicAppResponse)(nil)).Elem()
}

func (o AutomationActionLogicAppResponseOutput) ToAutomationActionLogicAppResponseOutput() AutomationActionLogicAppResponseOutput {
	return o
}

func (o AutomationActionLogicAppResponseOutput) ToAutomationActionLogicAppResponseOutputWithContext(ctx context.Context) AutomationActionLogicAppResponseOutput {
	return o
}

func (o AutomationActionLogicAppResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionLogicAppResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionLogicAppResponseOutput) LogicAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionLogicAppResponse) *string { return v.LogicAppResourceId }).(pulumi.StringPtrOutput)
}

func (o AutomationActionLogicAppResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionLogicAppResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type AutomationActionWorkspace struct {
	ActionType          string  `pulumi:"actionType"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}





type AutomationActionWorkspaceInput interface {
	pulumi.Input

	ToAutomationActionWorkspaceOutput() AutomationActionWorkspaceOutput
	ToAutomationActionWorkspaceOutputWithContext(context.Context) AutomationActionWorkspaceOutput
}

type AutomationActionWorkspaceArgs struct {
	ActionType          pulumi.StringInput    `pulumi:"actionType"`
	WorkspaceResourceId pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (AutomationActionWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionWorkspace)(nil)).Elem()
}

func (i AutomationActionWorkspaceArgs) ToAutomationActionWorkspaceOutput() AutomationActionWorkspaceOutput {
	return i.ToAutomationActionWorkspaceOutputWithContext(context.Background())
}

func (i AutomationActionWorkspaceArgs) ToAutomationActionWorkspaceOutputWithContext(ctx context.Context) AutomationActionWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionWorkspaceOutput)
}

type AutomationActionWorkspaceOutput struct{ *pulumi.OutputState }

func (AutomationActionWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionWorkspace)(nil)).Elem()
}

func (o AutomationActionWorkspaceOutput) ToAutomationActionWorkspaceOutput() AutomationActionWorkspaceOutput {
	return o
}

func (o AutomationActionWorkspaceOutput) ToAutomationActionWorkspaceOutputWithContext(ctx context.Context) AutomationActionWorkspaceOutput {
	return o
}

func (o AutomationActionWorkspaceOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionWorkspace) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionWorkspaceOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionWorkspace) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type AutomationActionWorkspaceResponse struct {
	ActionType          string  `pulumi:"actionType"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}





type AutomationActionWorkspaceResponseInput interface {
	pulumi.Input

	ToAutomationActionWorkspaceResponseOutput() AutomationActionWorkspaceResponseOutput
	ToAutomationActionWorkspaceResponseOutputWithContext(context.Context) AutomationActionWorkspaceResponseOutput
}

type AutomationActionWorkspaceResponseArgs struct {
	ActionType          pulumi.StringInput    `pulumi:"actionType"`
	WorkspaceResourceId pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (AutomationActionWorkspaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionWorkspaceResponse)(nil)).Elem()
}

func (i AutomationActionWorkspaceResponseArgs) ToAutomationActionWorkspaceResponseOutput() AutomationActionWorkspaceResponseOutput {
	return i.ToAutomationActionWorkspaceResponseOutputWithContext(context.Background())
}

func (i AutomationActionWorkspaceResponseArgs) ToAutomationActionWorkspaceResponseOutputWithContext(ctx context.Context) AutomationActionWorkspaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationActionWorkspaceResponseOutput)
}

type AutomationActionWorkspaceResponseOutput struct{ *pulumi.OutputState }

func (AutomationActionWorkspaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationActionWorkspaceResponse)(nil)).Elem()
}

func (o AutomationActionWorkspaceResponseOutput) ToAutomationActionWorkspaceResponseOutput() AutomationActionWorkspaceResponseOutput {
	return o
}

func (o AutomationActionWorkspaceResponseOutput) ToAutomationActionWorkspaceResponseOutputWithContext(ctx context.Context) AutomationActionWorkspaceResponseOutput {
	return o
}

func (o AutomationActionWorkspaceResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationActionWorkspaceResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationActionWorkspaceResponseOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationActionWorkspaceResponse) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
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





type AutomationRuleSetResponseInput interface {
	pulumi.Input

	ToAutomationRuleSetResponseOutput() AutomationRuleSetResponseOutput
	ToAutomationRuleSetResponseOutputWithContext(context.Context) AutomationRuleSetResponseOutput
}

type AutomationRuleSetResponseArgs struct {
	Rules AutomationTriggeringRuleResponseArrayInput `pulumi:"rules"`
}

func (AutomationRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleSetResponse)(nil)).Elem()
}

func (i AutomationRuleSetResponseArgs) ToAutomationRuleSetResponseOutput() AutomationRuleSetResponseOutput {
	return i.ToAutomationRuleSetResponseOutputWithContext(context.Background())
}

func (i AutomationRuleSetResponseArgs) ToAutomationRuleSetResponseOutputWithContext(ctx context.Context) AutomationRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleSetResponseOutput)
}





type AutomationRuleSetResponseArrayInput interface {
	pulumi.Input

	ToAutomationRuleSetResponseArrayOutput() AutomationRuleSetResponseArrayOutput
	ToAutomationRuleSetResponseArrayOutputWithContext(context.Context) AutomationRuleSetResponseArrayOutput
}

type AutomationRuleSetResponseArray []AutomationRuleSetResponseInput

func (AutomationRuleSetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRuleSetResponse)(nil)).Elem()
}

func (i AutomationRuleSetResponseArray) ToAutomationRuleSetResponseArrayOutput() AutomationRuleSetResponseArrayOutput {
	return i.ToAutomationRuleSetResponseArrayOutputWithContext(context.Background())
}

func (i AutomationRuleSetResponseArray) ToAutomationRuleSetResponseArrayOutputWithContext(ctx context.Context) AutomationRuleSetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleSetResponseArrayOutput)
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





type AutomationScopeResponseInput interface {
	pulumi.Input

	ToAutomationScopeResponseOutput() AutomationScopeResponseOutput
	ToAutomationScopeResponseOutputWithContext(context.Context) AutomationScopeResponseOutput
}

type AutomationScopeResponseArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	ScopePath   pulumi.StringPtrInput `pulumi:"scopePath"`
}

func (AutomationScopeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationScopeResponse)(nil)).Elem()
}

func (i AutomationScopeResponseArgs) ToAutomationScopeResponseOutput() AutomationScopeResponseOutput {
	return i.ToAutomationScopeResponseOutputWithContext(context.Background())
}

func (i AutomationScopeResponseArgs) ToAutomationScopeResponseOutputWithContext(ctx context.Context) AutomationScopeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationScopeResponseOutput)
}





type AutomationScopeResponseArrayInput interface {
	pulumi.Input

	ToAutomationScopeResponseArrayOutput() AutomationScopeResponseArrayOutput
	ToAutomationScopeResponseArrayOutputWithContext(context.Context) AutomationScopeResponseArrayOutput
}

type AutomationScopeResponseArray []AutomationScopeResponseInput

func (AutomationScopeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationScopeResponse)(nil)).Elem()
}

func (i AutomationScopeResponseArray) ToAutomationScopeResponseArrayOutput() AutomationScopeResponseArrayOutput {
	return i.ToAutomationScopeResponseArrayOutputWithContext(context.Background())
}

func (i AutomationScopeResponseArray) ToAutomationScopeResponseArrayOutputWithContext(ctx context.Context) AutomationScopeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationScopeResponseArrayOutput)
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





type AutomationSourceResponseInput interface {
	pulumi.Input

	ToAutomationSourceResponseOutput() AutomationSourceResponseOutput
	ToAutomationSourceResponseOutputWithContext(context.Context) AutomationSourceResponseOutput
}

type AutomationSourceResponseArgs struct {
	EventSource pulumi.StringPtrInput               `pulumi:"eventSource"`
	RuleSets    AutomationRuleSetResponseArrayInput `pulumi:"ruleSets"`
}

func (AutomationSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationSourceResponse)(nil)).Elem()
}

func (i AutomationSourceResponseArgs) ToAutomationSourceResponseOutput() AutomationSourceResponseOutput {
	return i.ToAutomationSourceResponseOutputWithContext(context.Background())
}

func (i AutomationSourceResponseArgs) ToAutomationSourceResponseOutputWithContext(ctx context.Context) AutomationSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationSourceResponseOutput)
}





type AutomationSourceResponseArrayInput interface {
	pulumi.Input

	ToAutomationSourceResponseArrayOutput() AutomationSourceResponseArrayOutput
	ToAutomationSourceResponseArrayOutputWithContext(context.Context) AutomationSourceResponseArrayOutput
}

type AutomationSourceResponseArray []AutomationSourceResponseInput

func (AutomationSourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationSourceResponse)(nil)).Elem()
}

func (i AutomationSourceResponseArray) ToAutomationSourceResponseArrayOutput() AutomationSourceResponseArrayOutput {
	return i.ToAutomationSourceResponseArrayOutputWithContext(context.Background())
}

func (i AutomationSourceResponseArray) ToAutomationSourceResponseArrayOutputWithContext(ctx context.Context) AutomationSourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationSourceResponseArrayOutput)
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





type AutomationTriggeringRuleResponseInput interface {
	pulumi.Input

	ToAutomationTriggeringRuleResponseOutput() AutomationTriggeringRuleResponseOutput
	ToAutomationTriggeringRuleResponseOutputWithContext(context.Context) AutomationTriggeringRuleResponseOutput
}

type AutomationTriggeringRuleResponseArgs struct {
	ExpectedValue pulumi.StringPtrInput `pulumi:"expectedValue"`
	Operator      pulumi.StringPtrInput `pulumi:"operator"`
	PropertyJPath pulumi.StringPtrInput `pulumi:"propertyJPath"`
	PropertyType  pulumi.StringPtrInput `pulumi:"propertyType"`
}

func (AutomationTriggeringRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationTriggeringRuleResponse)(nil)).Elem()
}

func (i AutomationTriggeringRuleResponseArgs) ToAutomationTriggeringRuleResponseOutput() AutomationTriggeringRuleResponseOutput {
	return i.ToAutomationTriggeringRuleResponseOutputWithContext(context.Background())
}

func (i AutomationTriggeringRuleResponseArgs) ToAutomationTriggeringRuleResponseOutputWithContext(ctx context.Context) AutomationTriggeringRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationTriggeringRuleResponseOutput)
}





type AutomationTriggeringRuleResponseArrayInput interface {
	pulumi.Input

	ToAutomationTriggeringRuleResponseArrayOutput() AutomationTriggeringRuleResponseArrayOutput
	ToAutomationTriggeringRuleResponseArrayOutputWithContext(context.Context) AutomationTriggeringRuleResponseArrayOutput
}

type AutomationTriggeringRuleResponseArray []AutomationTriggeringRuleResponseInput

func (AutomationTriggeringRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationTriggeringRuleResponse)(nil)).Elem()
}

func (i AutomationTriggeringRuleResponseArray) ToAutomationTriggeringRuleResponseArrayOutput() AutomationTriggeringRuleResponseArrayOutput {
	return i.ToAutomationTriggeringRuleResponseArrayOutputWithContext(context.Background())
}

func (i AutomationTriggeringRuleResponseArray) ToAutomationTriggeringRuleResponseArrayOutputWithContext(ctx context.Context) AutomationTriggeringRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationTriggeringRuleResponseArrayOutput)
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

type AwAssumeRoleAuthenticationDetailsProperties struct {
	AuthenticationType string `pulumi:"authenticationType"`
	AwsAssumeRoleArn   string `pulumi:"awsAssumeRoleArn"`
	AwsExternalId      string `pulumi:"awsExternalId"`
}





type AwAssumeRoleAuthenticationDetailsPropertiesInput interface {
	pulumi.Input

	ToAwAssumeRoleAuthenticationDetailsPropertiesOutput() AwAssumeRoleAuthenticationDetailsPropertiesOutput
	ToAwAssumeRoleAuthenticationDetailsPropertiesOutputWithContext(context.Context) AwAssumeRoleAuthenticationDetailsPropertiesOutput
}

type AwAssumeRoleAuthenticationDetailsPropertiesArgs struct {
	AuthenticationType pulumi.StringInput `pulumi:"authenticationType"`
	AwsAssumeRoleArn   pulumi.StringInput `pulumi:"awsAssumeRoleArn"`
	AwsExternalId      pulumi.StringInput `pulumi:"awsExternalId"`
}

func (AwAssumeRoleAuthenticationDetailsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwAssumeRoleAuthenticationDetailsProperties)(nil)).Elem()
}

func (i AwAssumeRoleAuthenticationDetailsPropertiesArgs) ToAwAssumeRoleAuthenticationDetailsPropertiesOutput() AwAssumeRoleAuthenticationDetailsPropertiesOutput {
	return i.ToAwAssumeRoleAuthenticationDetailsPropertiesOutputWithContext(context.Background())
}

func (i AwAssumeRoleAuthenticationDetailsPropertiesArgs) ToAwAssumeRoleAuthenticationDetailsPropertiesOutputWithContext(ctx context.Context) AwAssumeRoleAuthenticationDetailsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwAssumeRoleAuthenticationDetailsPropertiesOutput)
}

type AwAssumeRoleAuthenticationDetailsPropertiesOutput struct{ *pulumi.OutputState }

func (AwAssumeRoleAuthenticationDetailsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwAssumeRoleAuthenticationDetailsProperties)(nil)).Elem()
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesOutput) ToAwAssumeRoleAuthenticationDetailsPropertiesOutput() AwAssumeRoleAuthenticationDetailsPropertiesOutput {
	return o
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesOutput) ToAwAssumeRoleAuthenticationDetailsPropertiesOutputWithContext(ctx context.Context) AwAssumeRoleAuthenticationDetailsPropertiesOutput {
	return o
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v AwAssumeRoleAuthenticationDetailsProperties) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesOutput) AwsAssumeRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v AwAssumeRoleAuthenticationDetailsProperties) string { return v.AwsAssumeRoleArn }).(pulumi.StringOutput)
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesOutput) AwsExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v AwAssumeRoleAuthenticationDetailsProperties) string { return v.AwsExternalId }).(pulumi.StringOutput)
}

type AwAssumeRoleAuthenticationDetailsPropertiesResponse struct {
	AccountId                       string   `pulumi:"accountId"`
	AuthenticationProvisioningState string   `pulumi:"authenticationProvisioningState"`
	AuthenticationType              string   `pulumi:"authenticationType"`
	AwsAssumeRoleArn                string   `pulumi:"awsAssumeRoleArn"`
	AwsExternalId                   string   `pulumi:"awsExternalId"`
	GrantedPermissions              []string `pulumi:"grantedPermissions"`
}





type AwAssumeRoleAuthenticationDetailsPropertiesResponseInput interface {
	pulumi.Input

	ToAwAssumeRoleAuthenticationDetailsPropertiesResponseOutput() AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput
	ToAwAssumeRoleAuthenticationDetailsPropertiesResponseOutputWithContext(context.Context) AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput
}

type AwAssumeRoleAuthenticationDetailsPropertiesResponseArgs struct {
	AccountId                       pulumi.StringInput      `pulumi:"accountId"`
	AuthenticationProvisioningState pulumi.StringInput      `pulumi:"authenticationProvisioningState"`
	AuthenticationType              pulumi.StringInput      `pulumi:"authenticationType"`
	AwsAssumeRoleArn                pulumi.StringInput      `pulumi:"awsAssumeRoleArn"`
	AwsExternalId                   pulumi.StringInput      `pulumi:"awsExternalId"`
	GrantedPermissions              pulumi.StringArrayInput `pulumi:"grantedPermissions"`
}

func (AwAssumeRoleAuthenticationDetailsPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwAssumeRoleAuthenticationDetailsPropertiesResponse)(nil)).Elem()
}

func (i AwAssumeRoleAuthenticationDetailsPropertiesResponseArgs) ToAwAssumeRoleAuthenticationDetailsPropertiesResponseOutput() AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput {
	return i.ToAwAssumeRoleAuthenticationDetailsPropertiesResponseOutputWithContext(context.Background())
}

func (i AwAssumeRoleAuthenticationDetailsPropertiesResponseArgs) ToAwAssumeRoleAuthenticationDetailsPropertiesResponseOutputWithContext(ctx context.Context) AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput)
}

type AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwAssumeRoleAuthenticationDetailsPropertiesResponse)(nil)).Elem()
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput) ToAwAssumeRoleAuthenticationDetailsPropertiesResponseOutput() AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput {
	return o
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput) ToAwAssumeRoleAuthenticationDetailsPropertiesResponseOutputWithContext(ctx context.Context) AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput {
	return o
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AwAssumeRoleAuthenticationDetailsPropertiesResponse) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput) AuthenticationProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AwAssumeRoleAuthenticationDetailsPropertiesResponse) string {
		return v.AuthenticationProvisioningState
	}).(pulumi.StringOutput)
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v AwAssumeRoleAuthenticationDetailsPropertiesResponse) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput) AwsAssumeRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v AwAssumeRoleAuthenticationDetailsPropertiesResponse) string { return v.AwsAssumeRoleArn }).(pulumi.StringOutput)
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput) AwsExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v AwAssumeRoleAuthenticationDetailsPropertiesResponse) string { return v.AwsExternalId }).(pulumi.StringOutput)
}

func (o AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput) GrantedPermissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AwAssumeRoleAuthenticationDetailsPropertiesResponse) []string { return v.GrantedPermissions }).(pulumi.StringArrayOutput)
}

type AwsCredsAuthenticationDetailsProperties struct {
	AuthenticationType string `pulumi:"authenticationType"`
	AwsAccessKeyId     string `pulumi:"awsAccessKeyId"`
	AwsSecretAccessKey string `pulumi:"awsSecretAccessKey"`
}





type AwsCredsAuthenticationDetailsPropertiesInput interface {
	pulumi.Input

	ToAwsCredsAuthenticationDetailsPropertiesOutput() AwsCredsAuthenticationDetailsPropertiesOutput
	ToAwsCredsAuthenticationDetailsPropertiesOutputWithContext(context.Context) AwsCredsAuthenticationDetailsPropertiesOutput
}

type AwsCredsAuthenticationDetailsPropertiesArgs struct {
	AuthenticationType pulumi.StringInput `pulumi:"authenticationType"`
	AwsAccessKeyId     pulumi.StringInput `pulumi:"awsAccessKeyId"`
	AwsSecretAccessKey pulumi.StringInput `pulumi:"awsSecretAccessKey"`
}

func (AwsCredsAuthenticationDetailsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCredsAuthenticationDetailsProperties)(nil)).Elem()
}

func (i AwsCredsAuthenticationDetailsPropertiesArgs) ToAwsCredsAuthenticationDetailsPropertiesOutput() AwsCredsAuthenticationDetailsPropertiesOutput {
	return i.ToAwsCredsAuthenticationDetailsPropertiesOutputWithContext(context.Background())
}

func (i AwsCredsAuthenticationDetailsPropertiesArgs) ToAwsCredsAuthenticationDetailsPropertiesOutputWithContext(ctx context.Context) AwsCredsAuthenticationDetailsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCredsAuthenticationDetailsPropertiesOutput)
}

type AwsCredsAuthenticationDetailsPropertiesOutput struct{ *pulumi.OutputState }

func (AwsCredsAuthenticationDetailsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCredsAuthenticationDetailsProperties)(nil)).Elem()
}

func (o AwsCredsAuthenticationDetailsPropertiesOutput) ToAwsCredsAuthenticationDetailsPropertiesOutput() AwsCredsAuthenticationDetailsPropertiesOutput {
	return o
}

func (o AwsCredsAuthenticationDetailsPropertiesOutput) ToAwsCredsAuthenticationDetailsPropertiesOutputWithContext(ctx context.Context) AwsCredsAuthenticationDetailsPropertiesOutput {
	return o
}

func (o AwsCredsAuthenticationDetailsPropertiesOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCredsAuthenticationDetailsProperties) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o AwsCredsAuthenticationDetailsPropertiesOutput) AwsAccessKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCredsAuthenticationDetailsProperties) string { return v.AwsAccessKeyId }).(pulumi.StringOutput)
}

func (o AwsCredsAuthenticationDetailsPropertiesOutput) AwsSecretAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCredsAuthenticationDetailsProperties) string { return v.AwsSecretAccessKey }).(pulumi.StringOutput)
}

type AwsCredsAuthenticationDetailsPropertiesResponse struct {
	AccountId                       string   `pulumi:"accountId"`
	AuthenticationProvisioningState string   `pulumi:"authenticationProvisioningState"`
	AuthenticationType              string   `pulumi:"authenticationType"`
	AwsAccessKeyId                  string   `pulumi:"awsAccessKeyId"`
	AwsSecretAccessKey              string   `pulumi:"awsSecretAccessKey"`
	GrantedPermissions              []string `pulumi:"grantedPermissions"`
}





type AwsCredsAuthenticationDetailsPropertiesResponseInput interface {
	pulumi.Input

	ToAwsCredsAuthenticationDetailsPropertiesResponseOutput() AwsCredsAuthenticationDetailsPropertiesResponseOutput
	ToAwsCredsAuthenticationDetailsPropertiesResponseOutputWithContext(context.Context) AwsCredsAuthenticationDetailsPropertiesResponseOutput
}

type AwsCredsAuthenticationDetailsPropertiesResponseArgs struct {
	AccountId                       pulumi.StringInput      `pulumi:"accountId"`
	AuthenticationProvisioningState pulumi.StringInput      `pulumi:"authenticationProvisioningState"`
	AuthenticationType              pulumi.StringInput      `pulumi:"authenticationType"`
	AwsAccessKeyId                  pulumi.StringInput      `pulumi:"awsAccessKeyId"`
	AwsSecretAccessKey              pulumi.StringInput      `pulumi:"awsSecretAccessKey"`
	GrantedPermissions              pulumi.StringArrayInput `pulumi:"grantedPermissions"`
}

func (AwsCredsAuthenticationDetailsPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCredsAuthenticationDetailsPropertiesResponse)(nil)).Elem()
}

func (i AwsCredsAuthenticationDetailsPropertiesResponseArgs) ToAwsCredsAuthenticationDetailsPropertiesResponseOutput() AwsCredsAuthenticationDetailsPropertiesResponseOutput {
	return i.ToAwsCredsAuthenticationDetailsPropertiesResponseOutputWithContext(context.Background())
}

func (i AwsCredsAuthenticationDetailsPropertiesResponseArgs) ToAwsCredsAuthenticationDetailsPropertiesResponseOutputWithContext(ctx context.Context) AwsCredsAuthenticationDetailsPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCredsAuthenticationDetailsPropertiesResponseOutput)
}

type AwsCredsAuthenticationDetailsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AwsCredsAuthenticationDetailsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCredsAuthenticationDetailsPropertiesResponse)(nil)).Elem()
}

func (o AwsCredsAuthenticationDetailsPropertiesResponseOutput) ToAwsCredsAuthenticationDetailsPropertiesResponseOutput() AwsCredsAuthenticationDetailsPropertiesResponseOutput {
	return o
}

func (o AwsCredsAuthenticationDetailsPropertiesResponseOutput) ToAwsCredsAuthenticationDetailsPropertiesResponseOutputWithContext(ctx context.Context) AwsCredsAuthenticationDetailsPropertiesResponseOutput {
	return o
}

func (o AwsCredsAuthenticationDetailsPropertiesResponseOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCredsAuthenticationDetailsPropertiesResponse) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o AwsCredsAuthenticationDetailsPropertiesResponseOutput) AuthenticationProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCredsAuthenticationDetailsPropertiesResponse) string {
		return v.AuthenticationProvisioningState
	}).(pulumi.StringOutput)
}

func (o AwsCredsAuthenticationDetailsPropertiesResponseOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCredsAuthenticationDetailsPropertiesResponse) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o AwsCredsAuthenticationDetailsPropertiesResponseOutput) AwsAccessKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCredsAuthenticationDetailsPropertiesResponse) string { return v.AwsAccessKeyId }).(pulumi.StringOutput)
}

func (o AwsCredsAuthenticationDetailsPropertiesResponseOutput) AwsSecretAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCredsAuthenticationDetailsPropertiesResponse) string { return v.AwsSecretAccessKey }).(pulumi.StringOutput)
}

func (o AwsCredsAuthenticationDetailsPropertiesResponseOutput) GrantedPermissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AwsCredsAuthenticationDetailsPropertiesResponse) []string { return v.GrantedPermissions }).(pulumi.StringArrayOutput)
}

type AzureResourceDetails struct {
	Source string `pulumi:"source"`
}





type AzureResourceDetailsInput interface {
	pulumi.Input

	ToAzureResourceDetailsOutput() AzureResourceDetailsOutput
	ToAzureResourceDetailsOutputWithContext(context.Context) AzureResourceDetailsOutput
}

type AzureResourceDetailsArgs struct {
	Source pulumi.StringInput `pulumi:"source"`
}

func (AzureResourceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceDetails)(nil)).Elem()
}

func (i AzureResourceDetailsArgs) ToAzureResourceDetailsOutput() AzureResourceDetailsOutput {
	return i.ToAzureResourceDetailsOutputWithContext(context.Background())
}

func (i AzureResourceDetailsArgs) ToAzureResourceDetailsOutputWithContext(ctx context.Context) AzureResourceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureResourceDetailsOutput)
}

type AzureResourceDetailsOutput struct{ *pulumi.OutputState }

func (AzureResourceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceDetails)(nil)).Elem()
}

func (o AzureResourceDetailsOutput) ToAzureResourceDetailsOutput() AzureResourceDetailsOutput {
	return o
}

func (o AzureResourceDetailsOutput) ToAzureResourceDetailsOutputWithContext(ctx context.Context) AzureResourceDetailsOutput {
	return o
}

func (o AzureResourceDetailsOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceDetails) string { return v.Source }).(pulumi.StringOutput)
}

type AzureResourceDetailsResponse struct {
	Id     string `pulumi:"id"`
	Source string `pulumi:"source"`
}





type AzureResourceDetailsResponseInput interface {
	pulumi.Input

	ToAzureResourceDetailsResponseOutput() AzureResourceDetailsResponseOutput
	ToAzureResourceDetailsResponseOutputWithContext(context.Context) AzureResourceDetailsResponseOutput
}

type AzureResourceDetailsResponseArgs struct {
	Id     pulumi.StringInput `pulumi:"id"`
	Source pulumi.StringInput `pulumi:"source"`
}

func (AzureResourceDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceDetailsResponse)(nil)).Elem()
}

func (i AzureResourceDetailsResponseArgs) ToAzureResourceDetailsResponseOutput() AzureResourceDetailsResponseOutput {
	return i.ToAzureResourceDetailsResponseOutputWithContext(context.Background())
}

func (i AzureResourceDetailsResponseArgs) ToAzureResourceDetailsResponseOutputWithContext(ctx context.Context) AzureResourceDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureResourceDetailsResponseOutput)
}

type AzureResourceDetailsResponseOutput struct{ *pulumi.OutputState }

func (AzureResourceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceDetailsResponse)(nil)).Elem()
}

func (o AzureResourceDetailsResponseOutput) ToAzureResourceDetailsResponseOutput() AzureResourceDetailsResponseOutput {
	return o
}

func (o AzureResourceDetailsResponseOutput) ToAzureResourceDetailsResponseOutputWithContext(ctx context.Context) AzureResourceDetailsResponseOutput {
	return o
}

func (o AzureResourceDetailsResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceDetailsResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o AzureResourceDetailsResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceDetailsResponse) string { return v.Source }).(pulumi.StringOutput)
}

type CspmMonitorAwsOffering struct {
	NativeCloudConnection *CspmMonitorAwsOfferingNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                       `pulumi:"offeringType"`
}





type CspmMonitorAwsOfferingInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingOutput() CspmMonitorAwsOfferingOutput
	ToCspmMonitorAwsOfferingOutputWithContext(context.Context) CspmMonitorAwsOfferingOutput
}

type CspmMonitorAwsOfferingArgs struct {
	NativeCloudConnection CspmMonitorAwsOfferingNativeCloudConnectionPtrInput `pulumi:"nativeCloudConnection"`
	OfferingType          pulumi.StringInput                                  `pulumi:"offeringType"`
}

func (CspmMonitorAwsOfferingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOffering)(nil)).Elem()
}

func (i CspmMonitorAwsOfferingArgs) ToCspmMonitorAwsOfferingOutput() CspmMonitorAwsOfferingOutput {
	return i.ToCspmMonitorAwsOfferingOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingArgs) ToCspmMonitorAwsOfferingOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingOutput)
}

type CspmMonitorAwsOfferingOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOffering)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingOutput) ToCspmMonitorAwsOfferingOutput() CspmMonitorAwsOfferingOutput {
	return o
}

func (o CspmMonitorAwsOfferingOutput) ToCspmMonitorAwsOfferingOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingOutput {
	return o
}

func (o CspmMonitorAwsOfferingOutput) NativeCloudConnection() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o.ApplyT(func(v CspmMonitorAwsOffering) *CspmMonitorAwsOfferingNativeCloudConnection {
		return v.NativeCloudConnection
	}).(CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput)
}

func (o CspmMonitorAwsOfferingOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v CspmMonitorAwsOffering) string { return v.OfferingType }).(pulumi.StringOutput)
}

type CspmMonitorAwsOfferingNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type CspmMonitorAwsOfferingNativeCloudConnectionInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingNativeCloudConnectionOutput() CspmMonitorAwsOfferingNativeCloudConnectionOutput
	ToCspmMonitorAwsOfferingNativeCloudConnectionOutputWithContext(context.Context) CspmMonitorAwsOfferingNativeCloudConnectionOutput
}

type CspmMonitorAwsOfferingNativeCloudConnectionArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (CspmMonitorAwsOfferingNativeCloudConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingNativeCloudConnection)(nil)).Elem()
}

func (i CspmMonitorAwsOfferingNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingNativeCloudConnectionOutput() CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return i.ToCspmMonitorAwsOfferingNativeCloudConnectionOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingNativeCloudConnectionOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingNativeCloudConnectionOutput)
}

func (i CspmMonitorAwsOfferingNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return i.ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingNativeCloudConnectionOutput).ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx)
}









type CspmMonitorAwsOfferingNativeCloudConnectionPtrInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput
	ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput
}

type cspmMonitorAwsOfferingNativeCloudConnectionPtrType CspmMonitorAwsOfferingNativeCloudConnectionArgs

func CspmMonitorAwsOfferingNativeCloudConnectionPtr(v *CspmMonitorAwsOfferingNativeCloudConnectionArgs) CspmMonitorAwsOfferingNativeCloudConnectionPtrInput {
	return (*cspmMonitorAwsOfferingNativeCloudConnectionPtrType)(v)
}

func (*cspmMonitorAwsOfferingNativeCloudConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CspmMonitorAwsOfferingNativeCloudConnection)(nil)).Elem()
}

func (i *cspmMonitorAwsOfferingNativeCloudConnectionPtrType) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return i.ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (i *cspmMonitorAwsOfferingNativeCloudConnectionPtrType) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput)
}

type CspmMonitorAwsOfferingNativeCloudConnectionOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingNativeCloudConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingNativeCloudConnection)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionOutput() CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return o
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return o
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o.ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CspmMonitorAwsOfferingNativeCloudConnection) *CspmMonitorAwsOfferingNativeCloudConnection {
		return &v
	}).(CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput)
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingNativeCloudConnection) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CspmMonitorAwsOfferingNativeCloudConnection)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) ToCspmMonitorAwsOfferingNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput {
	return o
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) Elem() CspmMonitorAwsOfferingNativeCloudConnectionOutput {
	return o.ApplyT(func(v *CspmMonitorAwsOfferingNativeCloudConnection) CspmMonitorAwsOfferingNativeCloudConnection {
		if v != nil {
			return *v
		}
		var ret CspmMonitorAwsOfferingNativeCloudConnection
		return ret
	}).(CspmMonitorAwsOfferingNativeCloudConnectionOutput)
}

func (o CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CspmMonitorAwsOfferingNativeCloudConnection) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type CspmMonitorAwsOfferingResponse struct {
	Description           string                                               `pulumi:"description"`
	NativeCloudConnection *CspmMonitorAwsOfferingResponseNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                               `pulumi:"offeringType"`
}





type CspmMonitorAwsOfferingResponseInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingResponseOutput() CspmMonitorAwsOfferingResponseOutput
	ToCspmMonitorAwsOfferingResponseOutputWithContext(context.Context) CspmMonitorAwsOfferingResponseOutput
}

type CspmMonitorAwsOfferingResponseArgs struct {
	Description           pulumi.StringInput                                          `pulumi:"description"`
	NativeCloudConnection CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrInput `pulumi:"nativeCloudConnection"`
	OfferingType          pulumi.StringInput                                          `pulumi:"offeringType"`
}

func (CspmMonitorAwsOfferingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingResponse)(nil)).Elem()
}

func (i CspmMonitorAwsOfferingResponseArgs) ToCspmMonitorAwsOfferingResponseOutput() CspmMonitorAwsOfferingResponseOutput {
	return i.ToCspmMonitorAwsOfferingResponseOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingResponseArgs) ToCspmMonitorAwsOfferingResponseOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingResponseOutput)
}

type CspmMonitorAwsOfferingResponseOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingResponse)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingResponseOutput) ToCspmMonitorAwsOfferingResponseOutput() CspmMonitorAwsOfferingResponseOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseOutput) ToCspmMonitorAwsOfferingResponseOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o CspmMonitorAwsOfferingResponseOutput) NativeCloudConnection() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingResponse) *CspmMonitorAwsOfferingResponseNativeCloudConnection {
		return v.NativeCloudConnection
	}).(CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput)
}

func (o CspmMonitorAwsOfferingResponseOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingResponse) string { return v.OfferingType }).(pulumi.StringOutput)
}

type CspmMonitorAwsOfferingResponseNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type CspmMonitorAwsOfferingResponseNativeCloudConnectionInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput
	ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutputWithContext(context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput
}

type CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingResponseNativeCloudConnection)(nil)).Elem()
}

func (i CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return i.ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput)
}

func (i CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return i.ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (i CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput).ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx)
}









type CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrInput interface {
	pulumi.Input

	ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput
	ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput
}

type cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs

func CspmMonitorAwsOfferingResponseNativeCloudConnectionPtr(v *CspmMonitorAwsOfferingResponseNativeCloudConnectionArgs) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrInput {
	return (*cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType)(v)
}

func (*cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CspmMonitorAwsOfferingResponseNativeCloudConnection)(nil)).Elem()
}

func (i *cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return i.ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (i *cspmMonitorAwsOfferingResponseNativeCloudConnectionPtrType) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput)
}

type CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CspmMonitorAwsOfferingResponseNativeCloudConnection)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o.ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(context.Background())
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CspmMonitorAwsOfferingResponseNativeCloudConnection) *CspmMonitorAwsOfferingResponseNativeCloudConnection {
		return &v
	}).(CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput)
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CspmMonitorAwsOfferingResponseNativeCloudConnection) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput struct{ *pulumi.OutputState }

func (CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CspmMonitorAwsOfferingResponseNativeCloudConnection)(nil)).Elem()
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput() CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) ToCspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutputWithContext(ctx context.Context) CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput {
	return o
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) Elem() CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput {
	return o.ApplyT(func(v *CspmMonitorAwsOfferingResponseNativeCloudConnection) CspmMonitorAwsOfferingResponseNativeCloudConnection {
		if v != nil {
			return *v
		}
		var ret CspmMonitorAwsOfferingResponseNativeCloudConnection
		return ret
	}).(CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput)
}

func (o CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CspmMonitorAwsOfferingResponseNativeCloudConnection) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOffering struct {
	CloudWatchToKinesis   *DefenderForContainersAwsOfferingCloudWatchToKinesis   `pulumi:"cloudWatchToKinesis"`
	KinesisToS3           *DefenderForContainersAwsOfferingKinesisToS3           `pulumi:"kinesisToS3"`
	KubernetesScubaReader *DefenderForContainersAwsOfferingKubernetesScubaReader `pulumi:"kubernetesScubaReader"`
	KubernetesService     *DefenderForContainersAwsOfferingKubernetesService     `pulumi:"kubernetesService"`
	OfferingType          string                                                 `pulumi:"offeringType"`
}





type DefenderForContainersAwsOfferingInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingOutput() DefenderForContainersAwsOfferingOutput
	ToDefenderForContainersAwsOfferingOutputWithContext(context.Context) DefenderForContainersAwsOfferingOutput
}

type DefenderForContainersAwsOfferingArgs struct {
	CloudWatchToKinesis   DefenderForContainersAwsOfferingCloudWatchToKinesisPtrInput   `pulumi:"cloudWatchToKinesis"`
	KinesisToS3           DefenderForContainersAwsOfferingKinesisToS3PtrInput           `pulumi:"kinesisToS3"`
	KubernetesScubaReader DefenderForContainersAwsOfferingKubernetesScubaReaderPtrInput `pulumi:"kubernetesScubaReader"`
	KubernetesService     DefenderForContainersAwsOfferingKubernetesServicePtrInput     `pulumi:"kubernetesService"`
	OfferingType          pulumi.StringInput                                            `pulumi:"offeringType"`
}

func (DefenderForContainersAwsOfferingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOffering)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingArgs) ToDefenderForContainersAwsOfferingOutput() DefenderForContainersAwsOfferingOutput {
	return i.ToDefenderForContainersAwsOfferingOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingArgs) ToDefenderForContainersAwsOfferingOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingOutput)
}

type DefenderForContainersAwsOfferingOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOffering)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingOutput) ToDefenderForContainersAwsOfferingOutput() DefenderForContainersAwsOfferingOutput {
	return o
}

func (o DefenderForContainersAwsOfferingOutput) ToDefenderForContainersAwsOfferingOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingOutput {
	return o
}

func (o DefenderForContainersAwsOfferingOutput) CloudWatchToKinesis() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) *DefenderForContainersAwsOfferingCloudWatchToKinesis {
		return v.CloudWatchToKinesis
	}).(DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput)
}

func (o DefenderForContainersAwsOfferingOutput) KinesisToS3() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) *DefenderForContainersAwsOfferingKinesisToS3 {
		return v.KinesisToS3
	}).(DefenderForContainersAwsOfferingKinesisToS3PtrOutput)
}

func (o DefenderForContainersAwsOfferingOutput) KubernetesScubaReader() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) *DefenderForContainersAwsOfferingKubernetesScubaReader {
		return v.KubernetesScubaReader
	}).(DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput)
}

func (o DefenderForContainersAwsOfferingOutput) KubernetesService() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) *DefenderForContainersAwsOfferingKubernetesService {
		return v.KubernetesService
	}).(DefenderForContainersAwsOfferingKubernetesServicePtrOutput)
}

func (o DefenderForContainersAwsOfferingOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOffering) string { return v.OfferingType }).(pulumi.StringOutput)
}

type DefenderForContainersAwsOfferingCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingCloudWatchToKinesisInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisOutput
	ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutputWithContext(context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisOutput
}

type DefenderForContainersAwsOfferingCloudWatchToKinesisArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingCloudWatchToKinesis)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return i.ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingCloudWatchToKinesisOutput)
}

func (i DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return i.ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingCloudWatchToKinesisOutput).ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingCloudWatchToKinesisPtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput
	ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput
}

type defenderForContainersAwsOfferingCloudWatchToKinesisPtrType DefenderForContainersAwsOfferingCloudWatchToKinesisArgs

func DefenderForContainersAwsOfferingCloudWatchToKinesisPtr(v *DefenderForContainersAwsOfferingCloudWatchToKinesisArgs) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrInput {
	return (*defenderForContainersAwsOfferingCloudWatchToKinesisPtrType)(v)
}

func (*defenderForContainersAwsOfferingCloudWatchToKinesisPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingCloudWatchToKinesis)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingCloudWatchToKinesisPtrType) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return i.ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingCloudWatchToKinesisPtrType) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput)
}

type DefenderForContainersAwsOfferingCloudWatchToKinesisOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingCloudWatchToKinesis)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return o
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return o
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o.ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingCloudWatchToKinesis) *DefenderForContainersAwsOfferingCloudWatchToKinesis {
		return &v
	}).(DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput)
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingCloudWatchToKinesis) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingCloudWatchToKinesis)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) ToDefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) Elem() DefenderForContainersAwsOfferingCloudWatchToKinesisOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingCloudWatchToKinesis) DefenderForContainersAwsOfferingCloudWatchToKinesis {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingCloudWatchToKinesis
		return ret
	}).(DefenderForContainersAwsOfferingCloudWatchToKinesisOutput)
}

func (o DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingCloudWatchToKinesis) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingKinesisToS3Input interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKinesisToS3Output() DefenderForContainersAwsOfferingKinesisToS3Output
	ToDefenderForContainersAwsOfferingKinesisToS3OutputWithContext(context.Context) DefenderForContainersAwsOfferingKinesisToS3Output
}

type DefenderForContainersAwsOfferingKinesisToS3Args struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingKinesisToS3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKinesisToS3)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingKinesisToS3Args) ToDefenderForContainersAwsOfferingKinesisToS3Output() DefenderForContainersAwsOfferingKinesisToS3Output {
	return i.ToDefenderForContainersAwsOfferingKinesisToS3OutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKinesisToS3Args) ToDefenderForContainersAwsOfferingKinesisToS3OutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3Output {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKinesisToS3Output)
}

func (i DefenderForContainersAwsOfferingKinesisToS3Args) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return i.ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKinesisToS3Args) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKinesisToS3Output).ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingKinesisToS3PtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput
	ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput
}

type defenderForContainersAwsOfferingKinesisToS3PtrType DefenderForContainersAwsOfferingKinesisToS3Args

func DefenderForContainersAwsOfferingKinesisToS3Ptr(v *DefenderForContainersAwsOfferingKinesisToS3Args) DefenderForContainersAwsOfferingKinesisToS3PtrInput {
	return (*defenderForContainersAwsOfferingKinesisToS3PtrType)(v)
}

func (*defenderForContainersAwsOfferingKinesisToS3PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKinesisToS3)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingKinesisToS3PtrType) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return i.ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingKinesisToS3PtrType) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKinesisToS3PtrOutput)
}

type DefenderForContainersAwsOfferingKinesisToS3Output struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKinesisToS3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKinesisToS3)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) ToDefenderForContainersAwsOfferingKinesisToS3Output() DefenderForContainersAwsOfferingKinesisToS3Output {
	return o
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) ToDefenderForContainersAwsOfferingKinesisToS3OutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3Output {
	return o
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o.ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingKinesisToS3) *DefenderForContainersAwsOfferingKinesisToS3 {
		return &v
	}).(DefenderForContainersAwsOfferingKinesisToS3PtrOutput)
}

func (o DefenderForContainersAwsOfferingKinesisToS3Output) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingKinesisToS3) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKinesisToS3PtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKinesisToS3PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKinesisToS3)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKinesisToS3PtrOutput) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutput() DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKinesisToS3PtrOutput) ToDefenderForContainersAwsOfferingKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKinesisToS3PtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKinesisToS3PtrOutput) Elem() DefenderForContainersAwsOfferingKinesisToS3Output {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKinesisToS3) DefenderForContainersAwsOfferingKinesisToS3 {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingKinesisToS3
		return ret
	}).(DefenderForContainersAwsOfferingKinesisToS3Output)
}

func (o DefenderForContainersAwsOfferingKinesisToS3PtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKinesisToS3) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingKubernetesScubaReaderInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderOutput
	ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutputWithContext(context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderOutput
}

type DefenderForContainersAwsOfferingKubernetesScubaReaderArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesScubaReader)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesScubaReaderOutput)
}

func (i DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesScubaReaderOutput).ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingKubernetesScubaReaderPtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput
	ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput
}

type defenderForContainersAwsOfferingKubernetesScubaReaderPtrType DefenderForContainersAwsOfferingKubernetesScubaReaderArgs

func DefenderForContainersAwsOfferingKubernetesScubaReaderPtr(v *DefenderForContainersAwsOfferingKubernetesScubaReaderArgs) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrInput {
	return (*defenderForContainersAwsOfferingKubernetesScubaReaderPtrType)(v)
}

func (*defenderForContainersAwsOfferingKubernetesScubaReaderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKubernetesScubaReader)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingKubernetesScubaReaderPtrType) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingKubernetesScubaReaderPtrType) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesScubaReaderOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesScubaReader)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o.ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingKubernetesScubaReader) *DefenderForContainersAwsOfferingKubernetesScubaReader {
		return &v
	}).(DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput)
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingKubernetesScubaReader) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKubernetesScubaReader)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) ToDefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) Elem() DefenderForContainersAwsOfferingKubernetesScubaReaderOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKubernetesScubaReader) DefenderForContainersAwsOfferingKubernetesScubaReader {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingKubernetesScubaReader
		return ret
	}).(DefenderForContainersAwsOfferingKubernetesScubaReaderOutput)
}

func (o DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKubernetesScubaReader) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingKubernetesServiceInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKubernetesServiceOutput() DefenderForContainersAwsOfferingKubernetesServiceOutput
	ToDefenderForContainersAwsOfferingKubernetesServiceOutputWithContext(context.Context) DefenderForContainersAwsOfferingKubernetesServiceOutput
}

type DefenderForContainersAwsOfferingKubernetesServiceArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingKubernetesServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesService)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingKubernetesServiceArgs) ToDefenderForContainersAwsOfferingKubernetesServiceOutput() DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesServiceOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKubernetesServiceArgs) ToDefenderForContainersAwsOfferingKubernetesServiceOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesServiceOutput)
}

func (i DefenderForContainersAwsOfferingKubernetesServiceArgs) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingKubernetesServiceArgs) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesServiceOutput).ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingKubernetesServicePtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput
	ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput
}

type defenderForContainersAwsOfferingKubernetesServicePtrType DefenderForContainersAwsOfferingKubernetesServiceArgs

func DefenderForContainersAwsOfferingKubernetesServicePtr(v *DefenderForContainersAwsOfferingKubernetesServiceArgs) DefenderForContainersAwsOfferingKubernetesServicePtrInput {
	return (*defenderForContainersAwsOfferingKubernetesServicePtrType)(v)
}

func (*defenderForContainersAwsOfferingKubernetesServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKubernetesService)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingKubernetesServicePtrType) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return i.ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingKubernetesServicePtrType) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingKubernetesServicePtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesServiceOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKubernetesServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingKubernetesService)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) ToDefenderForContainersAwsOfferingKubernetesServiceOutput() DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) ToDefenderForContainersAwsOfferingKubernetesServiceOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o.ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingKubernetesService) *DefenderForContainersAwsOfferingKubernetesService {
		return &v
	}).(DefenderForContainersAwsOfferingKubernetesServicePtrOutput)
}

func (o DefenderForContainersAwsOfferingKubernetesServiceOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingKubernetesService) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingKubernetesServicePtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingKubernetesServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingKubernetesService)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingKubernetesServicePtrOutput) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutput() DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesServicePtrOutput) ToDefenderForContainersAwsOfferingKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingKubernetesServicePtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingKubernetesServicePtrOutput) Elem() DefenderForContainersAwsOfferingKubernetesServiceOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKubernetesService) DefenderForContainersAwsOfferingKubernetesService {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingKubernetesService
		return ret
	}).(DefenderForContainersAwsOfferingKubernetesServiceOutput)
}

func (o DefenderForContainersAwsOfferingKubernetesServicePtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingKubernetesService) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponse struct {
	CloudWatchToKinesis   *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis   `pulumi:"cloudWatchToKinesis"`
	Description           string                                                         `pulumi:"description"`
	KinesisToS3           *DefenderForContainersAwsOfferingResponseKinesisToS3           `pulumi:"kinesisToS3"`
	KubernetesScubaReader *DefenderForContainersAwsOfferingResponseKubernetesScubaReader `pulumi:"kubernetesScubaReader"`
	KubernetesService     *DefenderForContainersAwsOfferingResponseKubernetesService     `pulumi:"kubernetesService"`
	OfferingType          string                                                         `pulumi:"offeringType"`
}





type DefenderForContainersAwsOfferingResponseInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseOutput() DefenderForContainersAwsOfferingResponseOutput
	ToDefenderForContainersAwsOfferingResponseOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseOutput
}

type DefenderForContainersAwsOfferingResponseArgs struct {
	CloudWatchToKinesis   DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrInput   `pulumi:"cloudWatchToKinesis"`
	Description           pulumi.StringInput                                                    `pulumi:"description"`
	KinesisToS3           DefenderForContainersAwsOfferingResponseKinesisToS3PtrInput           `pulumi:"kinesisToS3"`
	KubernetesScubaReader DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrInput `pulumi:"kubernetesScubaReader"`
	KubernetesService     DefenderForContainersAwsOfferingResponseKubernetesServicePtrInput     `pulumi:"kubernetesService"`
	OfferingType          pulumi.StringInput                                                    `pulumi:"offeringType"`
}

func (DefenderForContainersAwsOfferingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponse)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseArgs) ToDefenderForContainersAwsOfferingResponseOutput() DefenderForContainersAwsOfferingResponseOutput {
	return i.ToDefenderForContainersAwsOfferingResponseOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseArgs) ToDefenderForContainersAwsOfferingResponseOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseOutput)
}

type DefenderForContainersAwsOfferingResponseOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponse)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseOutput) ToDefenderForContainersAwsOfferingResponseOutput() DefenderForContainersAwsOfferingResponseOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseOutput) ToDefenderForContainersAwsOfferingResponseOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseOutput) CloudWatchToKinesis() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis {
		return v.CloudWatchToKinesis
	}).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) KinesisToS3() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) *DefenderForContainersAwsOfferingResponseKinesisToS3 {
		return v.KinesisToS3
	}).(DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) KubernetesScubaReader() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) *DefenderForContainersAwsOfferingResponseKubernetesScubaReader {
		return v.KubernetesScubaReader
	}).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) KubernetesService() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) *DefenderForContainersAwsOfferingResponseKubernetesService {
		return v.KubernetesService
	}).(DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponse) string { return v.OfferingType }).(pulumi.StringOutput)
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput
	ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseCloudWatchToKinesis)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return i.ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput)
}

func (i DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput).ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput
	ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput
}

type defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs

func DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtr(v *DefenderForContainersAwsOfferingResponseCloudWatchToKinesisArgs) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrInput {
	return (*defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType)(v)
}

func (*defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseCloudWatchToKinesis)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrType) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput)
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseCloudWatchToKinesis)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o.ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingResponseCloudWatchToKinesis) *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis {
		return &v
	}).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponseCloudWatchToKinesis) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseCloudWatchToKinesis)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) ToDefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) Elem() DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis) DefenderForContainersAwsOfferingResponseCloudWatchToKinesis {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingResponseCloudWatchToKinesis
		return ret
	}).(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput)
}

func (o DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingResponseKinesisToS3Input interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKinesisToS3Output() DefenderForContainersAwsOfferingResponseKinesisToS3Output
	ToDefenderForContainersAwsOfferingResponseKinesisToS3OutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3Output
}

type DefenderForContainersAwsOfferingResponseKinesisToS3Args struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingResponseKinesisToS3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKinesisToS3)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseKinesisToS3Args) ToDefenderForContainersAwsOfferingResponseKinesisToS3Output() DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return i.ToDefenderForContainersAwsOfferingResponseKinesisToS3OutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKinesisToS3Args) ToDefenderForContainersAwsOfferingResponseKinesisToS3OutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKinesisToS3Output)
}

func (i DefenderForContainersAwsOfferingResponseKinesisToS3Args) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKinesisToS3Args) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKinesisToS3Output).ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingResponseKinesisToS3PtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput
	ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput
}

type defenderForContainersAwsOfferingResponseKinesisToS3PtrType DefenderForContainersAwsOfferingResponseKinesisToS3Args

func DefenderForContainersAwsOfferingResponseKinesisToS3Ptr(v *DefenderForContainersAwsOfferingResponseKinesisToS3Args) DefenderForContainersAwsOfferingResponseKinesisToS3PtrInput {
	return (*defenderForContainersAwsOfferingResponseKinesisToS3PtrType)(v)
}

func (*defenderForContainersAwsOfferingResponseKinesisToS3PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKinesisToS3)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingResponseKinesisToS3PtrType) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingResponseKinesisToS3PtrType) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput)
}

type DefenderForContainersAwsOfferingResponseKinesisToS3Output struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKinesisToS3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKinesisToS3)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) ToDefenderForContainersAwsOfferingResponseKinesisToS3Output() DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) ToDefenderForContainersAwsOfferingResponseKinesisToS3OutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o.ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingResponseKinesisToS3) *DefenderForContainersAwsOfferingResponseKinesisToS3 {
		return &v
	}).(DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3Output) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponseKinesisToS3) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKinesisToS3)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput() DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) ToDefenderForContainersAwsOfferingResponseKinesisToS3PtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) Elem() DefenderForContainersAwsOfferingResponseKinesisToS3Output {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKinesisToS3) DefenderForContainersAwsOfferingResponseKinesisToS3 {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingResponseKinesisToS3
		return ret
	}).(DefenderForContainersAwsOfferingResponseKinesisToS3Output)
}

func (o DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKinesisToS3) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput
	ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesScubaReader)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput)
}

func (i DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput).ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput
	ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput
}

type defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs

func DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtr(v *DefenderForContainersAwsOfferingResponseKubernetesScubaReaderArgs) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrInput {
	return (*defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType)(v)
}

func (*defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKubernetesScubaReader)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrType) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesScubaReader)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o.ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingResponseKubernetesScubaReader) *DefenderForContainersAwsOfferingResponseKubernetesScubaReader {
		return &v
	}).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponseKubernetesScubaReader) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKubernetesScubaReader)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) ToDefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) Elem() DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKubernetesScubaReader) DefenderForContainersAwsOfferingResponseKubernetesScubaReader {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingResponseKubernetesScubaReader
		return ret
	}).(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput)
}

func (o DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKubernetesScubaReader) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForContainersAwsOfferingResponseKubernetesServiceInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutput() DefenderForContainersAwsOfferingResponseKubernetesServiceOutput
	ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKubernetesServiceOutput
}

type DefenderForContainersAwsOfferingResponseKubernetesServiceArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesService)(nil)).Elem()
}

func (i DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutput() DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesServiceOutput)
}

func (i DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(context.Background())
}

func (i DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesServiceOutput).ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx)
}









type DefenderForContainersAwsOfferingResponseKubernetesServicePtrInput interface {
	pulumi.Input

	ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput
	ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput
}

type defenderForContainersAwsOfferingResponseKubernetesServicePtrType DefenderForContainersAwsOfferingResponseKubernetesServiceArgs

func DefenderForContainersAwsOfferingResponseKubernetesServicePtr(v *DefenderForContainersAwsOfferingResponseKubernetesServiceArgs) DefenderForContainersAwsOfferingResponseKubernetesServicePtrInput {
	return (*defenderForContainersAwsOfferingResponseKubernetesServicePtrType)(v)
}

func (*defenderForContainersAwsOfferingResponseKubernetesServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKubernetesService)(nil)).Elem()
}

func (i *defenderForContainersAwsOfferingResponseKubernetesServicePtrType) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return i.ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(context.Background())
}

func (i *defenderForContainersAwsOfferingResponseKubernetesServicePtrType) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesServiceOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForContainersAwsOfferingResponseKubernetesService)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutput() DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServiceOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o.ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(context.Background())
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForContainersAwsOfferingResponseKubernetesService) *DefenderForContainersAwsOfferingResponseKubernetesService {
		return &v
	}).(DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput)
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServiceOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForContainersAwsOfferingResponseKubernetesService) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput struct{ *pulumi.OutputState }

func (DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForContainersAwsOfferingResponseKubernetesService)(nil)).Elem()
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput() DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) ToDefenderForContainersAwsOfferingResponseKubernetesServicePtrOutputWithContext(ctx context.Context) DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput {
	return o
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) Elem() DefenderForContainersAwsOfferingResponseKubernetesServiceOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKubernetesService) DefenderForContainersAwsOfferingResponseKubernetesService {
		if v != nil {
			return *v
		}
		var ret DefenderForContainersAwsOfferingResponseKubernetesService
		return ret
	}).(DefenderForContainersAwsOfferingResponseKubernetesServiceOutput)
}

func (o DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForContainersAwsOfferingResponseKubernetesService) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOffering struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingDefenderForServers  `pulumi:"defenderForServers"`
	OfferingType        string                                            `pulumi:"offeringType"`
}





type DefenderForServersAwsOfferingInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingOutput() DefenderForServersAwsOfferingOutput
	ToDefenderForServersAwsOfferingOutputWithContext(context.Context) DefenderForServersAwsOfferingOutput
}

type DefenderForServersAwsOfferingArgs struct {
	ArcAutoProvisioning DefenderForServersAwsOfferingArcAutoProvisioningPtrInput `pulumi:"arcAutoProvisioning"`
	DefenderForServers  DefenderForServersAwsOfferingDefenderForServersPtrInput  `pulumi:"defenderForServers"`
	OfferingType        pulumi.StringInput                                       `pulumi:"offeringType"`
}

func (DefenderForServersAwsOfferingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOffering)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingArgs) ToDefenderForServersAwsOfferingOutput() DefenderForServersAwsOfferingOutput {
	return i.ToDefenderForServersAwsOfferingOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingArgs) ToDefenderForServersAwsOfferingOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingOutput)
}

type DefenderForServersAwsOfferingOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOffering)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingOutput) ToDefenderForServersAwsOfferingOutput() DefenderForServersAwsOfferingOutput {
	return o
}

func (o DefenderForServersAwsOfferingOutput) ToDefenderForServersAwsOfferingOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingOutput {
	return o
}

func (o DefenderForServersAwsOfferingOutput) ArcAutoProvisioning() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOffering) *DefenderForServersAwsOfferingArcAutoProvisioning {
		return v.ArcAutoProvisioning
	}).(DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput)
}

func (o DefenderForServersAwsOfferingOutput) DefenderForServers() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOffering) *DefenderForServersAwsOfferingDefenderForServers {
		return v.DefenderForServers
	}).(DefenderForServersAwsOfferingDefenderForServersPtrOutput)
}

func (o DefenderForServersAwsOfferingOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForServersAwsOffering) string { return v.OfferingType }).(pulumi.StringOutput)
}

type DefenderForServersAwsOfferingArcAutoProvisioning struct {
	Enabled                        *bool                                                        `pulumi:"enabled"`
	ServicePrincipalSecretMetadata *DefenderForServersAwsOfferingServicePrincipalSecretMetadata `pulumi:"servicePrincipalSecretMetadata"`
}





type DefenderForServersAwsOfferingArcAutoProvisioningInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingArcAutoProvisioningOutput() DefenderForServersAwsOfferingArcAutoProvisioningOutput
	ToDefenderForServersAwsOfferingArcAutoProvisioningOutputWithContext(context.Context) DefenderForServersAwsOfferingArcAutoProvisioningOutput
}

type DefenderForServersAwsOfferingArcAutoProvisioningArgs struct {
	Enabled                        pulumi.BoolPtrInput                                                 `pulumi:"enabled"`
	ServicePrincipalSecretMetadata DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrInput `pulumi:"servicePrincipalSecretMetadata"`
}

func (DefenderForServersAwsOfferingArcAutoProvisioningArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingArcAutoProvisioning)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingArcAutoProvisioningOutput() DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return i.ToDefenderForServersAwsOfferingArcAutoProvisioningOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingArcAutoProvisioningOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingArcAutoProvisioningOutput)
}

func (i DefenderForServersAwsOfferingArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return i.ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingArcAutoProvisioningOutput).ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingArcAutoProvisioningPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput
	ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput
}

type defenderForServersAwsOfferingArcAutoProvisioningPtrType DefenderForServersAwsOfferingArcAutoProvisioningArgs

func DefenderForServersAwsOfferingArcAutoProvisioningPtr(v *DefenderForServersAwsOfferingArcAutoProvisioningArgs) DefenderForServersAwsOfferingArcAutoProvisioningPtrInput {
	return (*defenderForServersAwsOfferingArcAutoProvisioningPtrType)(v)
}

func (*defenderForServersAwsOfferingArcAutoProvisioningPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingArcAutoProvisioning)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingArcAutoProvisioningPtrType) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return i.ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingArcAutoProvisioningPtrType) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput)
}

type DefenderForServersAwsOfferingArcAutoProvisioningOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingArcAutoProvisioningOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingArcAutoProvisioning)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningOutput() DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return o
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return o
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o.ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingArcAutoProvisioning) *DefenderForServersAwsOfferingArcAutoProvisioning {
		return &v
	}).(DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput)
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingArcAutoProvisioning) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningOutput) ServicePrincipalSecretMetadata() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingArcAutoProvisioning) *DefenderForServersAwsOfferingServicePrincipalSecretMetadata {
		return v.ServicePrincipalSecretMetadata
	}).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingArcAutoProvisioning)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) ToDefenderForServersAwsOfferingArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) Elem() DefenderForServersAwsOfferingArcAutoProvisioningOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingArcAutoProvisioning) DefenderForServersAwsOfferingArcAutoProvisioning {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingArcAutoProvisioning
		return ret
	}).(DefenderForServersAwsOfferingArcAutoProvisioningOutput)
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingArcAutoProvisioning) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput) ServicePrincipalSecretMetadata() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingArcAutoProvisioning) *DefenderForServersAwsOfferingServicePrincipalSecretMetadata {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalSecretMetadata
	}).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForServersAwsOfferingDefenderForServersInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingDefenderForServersOutput() DefenderForServersAwsOfferingDefenderForServersOutput
	ToDefenderForServersAwsOfferingDefenderForServersOutputWithContext(context.Context) DefenderForServersAwsOfferingDefenderForServersOutput
}

type DefenderForServersAwsOfferingDefenderForServersArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForServersAwsOfferingDefenderForServersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingDefenderForServers)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingDefenderForServersArgs) ToDefenderForServersAwsOfferingDefenderForServersOutput() DefenderForServersAwsOfferingDefenderForServersOutput {
	return i.ToDefenderForServersAwsOfferingDefenderForServersOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingDefenderForServersArgs) ToDefenderForServersAwsOfferingDefenderForServersOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingDefenderForServersOutput)
}

func (i DefenderForServersAwsOfferingDefenderForServersArgs) ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return i.ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingDefenderForServersArgs) ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingDefenderForServersOutput).ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingDefenderForServersPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput
	ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput
}

type defenderForServersAwsOfferingDefenderForServersPtrType DefenderForServersAwsOfferingDefenderForServersArgs

func DefenderForServersAwsOfferingDefenderForServersPtr(v *DefenderForServersAwsOfferingDefenderForServersArgs) DefenderForServersAwsOfferingDefenderForServersPtrInput {
	return (*defenderForServersAwsOfferingDefenderForServersPtrType)(v)
}

func (*defenderForServersAwsOfferingDefenderForServersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingDefenderForServers)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingDefenderForServersPtrType) ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return i.ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingDefenderForServersPtrType) ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingDefenderForServersPtrOutput)
}

type DefenderForServersAwsOfferingDefenderForServersOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingDefenderForServersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingDefenderForServers)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) ToDefenderForServersAwsOfferingDefenderForServersOutput() DefenderForServersAwsOfferingDefenderForServersOutput {
	return o
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) ToDefenderForServersAwsOfferingDefenderForServersOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersOutput {
	return o
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o.ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingDefenderForServers) *DefenderForServersAwsOfferingDefenderForServers {
		return &v
	}).(DefenderForServersAwsOfferingDefenderForServersPtrOutput)
}

func (o DefenderForServersAwsOfferingDefenderForServersOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingDefenderForServers) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingDefenderForServersPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingDefenderForServersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingDefenderForServers)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingDefenderForServersPtrOutput) ToDefenderForServersAwsOfferingDefenderForServersPtrOutput() DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingDefenderForServersPtrOutput) ToDefenderForServersAwsOfferingDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingDefenderForServersPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingDefenderForServersPtrOutput) Elem() DefenderForServersAwsOfferingDefenderForServersOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingDefenderForServers) DefenderForServersAwsOfferingDefenderForServers {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingDefenderForServers
		return ret
	}).(DefenderForServersAwsOfferingDefenderForServersOutput)
}

func (o DefenderForServersAwsOfferingDefenderForServersPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingDefenderForServers) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingResponse struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingResponseArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingResponseDefenderForServers  `pulumi:"defenderForServers"`
	Description         string                                                    `pulumi:"description"`
	OfferingType        string                                                    `pulumi:"offeringType"`
}





type DefenderForServersAwsOfferingResponseInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseOutput() DefenderForServersAwsOfferingResponseOutput
	ToDefenderForServersAwsOfferingResponseOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseOutput
}

type DefenderForServersAwsOfferingResponseArgs struct {
	ArcAutoProvisioning DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrInput `pulumi:"arcAutoProvisioning"`
	DefenderForServers  DefenderForServersAwsOfferingResponseDefenderForServersPtrInput  `pulumi:"defenderForServers"`
	Description         pulumi.StringInput                                               `pulumi:"description"`
	OfferingType        pulumi.StringInput                                               `pulumi:"offeringType"`
}

func (DefenderForServersAwsOfferingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponse)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingResponseArgs) ToDefenderForServersAwsOfferingResponseOutput() DefenderForServersAwsOfferingResponseOutput {
	return i.ToDefenderForServersAwsOfferingResponseOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseArgs) ToDefenderForServersAwsOfferingResponseOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseOutput)
}

type DefenderForServersAwsOfferingResponseOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponse)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseOutput) ToDefenderForServersAwsOfferingResponseOutput() DefenderForServersAwsOfferingResponseOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseOutput) ToDefenderForServersAwsOfferingResponseOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseOutput) ArcAutoProvisioning() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponse) *DefenderForServersAwsOfferingResponseArcAutoProvisioning {
		return v.ArcAutoProvisioning
	}).(DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseOutput) DefenderForServers() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponse) *DefenderForServersAwsOfferingResponseDefenderForServers {
		return v.DefenderForServers
	}).(DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o DefenderForServersAwsOfferingResponseOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponse) string { return v.OfferingType }).(pulumi.StringOutput)
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioning struct {
	Enabled                        *bool                                                                `pulumi:"enabled"`
	ServicePrincipalSecretMetadata *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata `pulumi:"servicePrincipalSecretMetadata"`
}





type DefenderForServersAwsOfferingResponseArcAutoProvisioningInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput
	ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs struct {
	Enabled                        pulumi.BoolPtrInput                                                         `pulumi:"enabled"`
	ServicePrincipalSecretMetadata DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrInput `pulumi:"servicePrincipalSecretMetadata"`
}

func (DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseArcAutoProvisioning)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return i.ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput)
}

func (i DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput).ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput
	ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput
}

type defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs

func DefenderForServersAwsOfferingResponseArcAutoProvisioningPtr(v *DefenderForServersAwsOfferingResponseArcAutoProvisioningArgs) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrInput {
	return (*defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType)(v)
}

func (*defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseArcAutoProvisioning)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingResponseArcAutoProvisioningPtrType) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput)
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseArcAutoProvisioning)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o.ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingResponseArcAutoProvisioning) *DefenderForServersAwsOfferingResponseArcAutoProvisioning {
		return &v
	}).(DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseArcAutoProvisioning) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput) ServicePrincipalSecretMetadata() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseArcAutoProvisioning) *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata {
		return v.ServicePrincipalSecretMetadata
	}).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseArcAutoProvisioning)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput() DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) ToDefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) Elem() DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseArcAutoProvisioning) DefenderForServersAwsOfferingResponseArcAutoProvisioning {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingResponseArcAutoProvisioning
		return ret
	}).(DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput)
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseArcAutoProvisioning) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput) ServicePrincipalSecretMetadata() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseArcAutoProvisioning) *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalSecretMetadata
	}).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingResponseDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}





type DefenderForServersAwsOfferingResponseDefenderForServersInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseDefenderForServersOutput() DefenderForServersAwsOfferingResponseDefenderForServersOutput
	ToDefenderForServersAwsOfferingResponseDefenderForServersOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseDefenderForServersOutput
}

type DefenderForServersAwsOfferingResponseDefenderForServersArgs struct {
	CloudRoleArn pulumi.StringPtrInput `pulumi:"cloudRoleArn"`
}

func (DefenderForServersAwsOfferingResponseDefenderForServersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseDefenderForServers)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingResponseDefenderForServersArgs) ToDefenderForServersAwsOfferingResponseDefenderForServersOutput() DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return i.ToDefenderForServersAwsOfferingResponseDefenderForServersOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseDefenderForServersArgs) ToDefenderForServersAwsOfferingResponseDefenderForServersOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseDefenderForServersOutput)
}

func (i DefenderForServersAwsOfferingResponseDefenderForServersArgs) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseDefenderForServersArgs) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseDefenderForServersOutput).ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingResponseDefenderForServersPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput
	ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput
}

type defenderForServersAwsOfferingResponseDefenderForServersPtrType DefenderForServersAwsOfferingResponseDefenderForServersArgs

func DefenderForServersAwsOfferingResponseDefenderForServersPtr(v *DefenderForServersAwsOfferingResponseDefenderForServersArgs) DefenderForServersAwsOfferingResponseDefenderForServersPtrInput {
	return (*defenderForServersAwsOfferingResponseDefenderForServersPtrType)(v)
}

func (*defenderForServersAwsOfferingResponseDefenderForServersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseDefenderForServers)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingResponseDefenderForServersPtrType) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingResponseDefenderForServersPtrType) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput)
}

type DefenderForServersAwsOfferingResponseDefenderForServersOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseDefenderForServersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseDefenderForServers)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersOutput() DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o.ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingResponseDefenderForServers) *DefenderForServersAwsOfferingResponseDefenderForServers {
		return &v
	}).(DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseDefenderForServers) *string { return v.CloudRoleArn }).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseDefenderForServers)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutput() DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) ToDefenderForServersAwsOfferingResponseDefenderForServersPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) Elem() DefenderForServersAwsOfferingResponseDefenderForServersOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseDefenderForServers) DefenderForServersAwsOfferingResponseDefenderForServers {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingResponseDefenderForServers
		return ret
	}).(DefenderForServersAwsOfferingResponseDefenderForServersOutput)
}

func (o DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput) CloudRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseDefenderForServers) *string {
		if v == nil {
			return nil
		}
		return v.CloudRoleArn
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata struct {
	ExpiryDate           *string `pulumi:"expiryDate"`
	ParameterNameInStore *string `pulumi:"parameterNameInStore"`
	ParameterStoreRegion *string `pulumi:"parameterStoreRegion"`
}





type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput
	ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs struct {
	ExpiryDate           pulumi.StringPtrInput `pulumi:"expiryDate"`
	ParameterNameInStore pulumi.StringPtrInput `pulumi:"parameterNameInStore"`
	ParameterStoreRegion pulumi.StringPtrInput `pulumi:"parameterStoreRegion"`
}

func (DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return i.ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput)
}

func (i DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput).ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput
	ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput
}

type defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs

func DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtr(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataArgs) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrInput {
	return (*defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType)(v)
}

func (*defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return i.ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrType) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o.ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata {
		return &v
	}).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		return v.ExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ParameterNameInStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		return v.ParameterNameInStore
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput) ParameterStoreRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		return v.ParameterStoreRegion
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ToDefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) Elem() DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata
		return ret
	}).(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ParameterNameInStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ParameterNameInStore
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput) ParameterStoreRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ParameterStoreRegion
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadata struct {
	ExpiryDate           *string `pulumi:"expiryDate"`
	ParameterNameInStore *string `pulumi:"parameterNameInStore"`
	ParameterStoreRegion *string `pulumi:"parameterStoreRegion"`
}





type DefenderForServersAwsOfferingServicePrincipalSecretMetadataInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput
	ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutputWithContext(context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs struct {
	ExpiryDate           pulumi.StringPtrInput `pulumi:"expiryDate"`
	ParameterNameInStore pulumi.StringPtrInput `pulumi:"parameterNameInStore"`
	ParameterStoreRegion pulumi.StringPtrInput `pulumi:"parameterStoreRegion"`
}

func (DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingServicePrincipalSecretMetadata)(nil)).Elem()
}

func (i DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return i.ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput)
}

func (i DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return i.ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (i DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput).ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx)
}









type DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrInput interface {
	pulumi.Input

	ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput
	ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput
}

type defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs

func DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtr(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadataArgs) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrInput {
	return (*defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType)(v)
}

func (*defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingServicePrincipalSecretMetadata)(nil)).Elem()
}

func (i *defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return i.ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (i *defenderForServersAwsOfferingServicePrincipalSecretMetadataPtrType) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput)
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefenderForServersAwsOfferingServicePrincipalSecretMetadata)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return o
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return o
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o.ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(context.Background())
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *DefenderForServersAwsOfferingServicePrincipalSecretMetadata {
		return &v
	}).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string { return v.ExpiryDate }).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ParameterNameInStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		return v.ParameterNameInStore
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput) ParameterStoreRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		return v.ParameterStoreRegion
	}).(pulumi.StringPtrOutput)
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput struct{ *pulumi.OutputState }

func (DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefenderForServersAwsOfferingServicePrincipalSecretMetadata)(nil)).Elem()
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput() DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ToDefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutputWithContext(ctx context.Context) DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput {
	return o
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) Elem() DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadata) DefenderForServersAwsOfferingServicePrincipalSecretMetadata {
		if v != nil {
			return *v
		}
		var ret DefenderForServersAwsOfferingServicePrincipalSecretMetadata
		return ret
	}).(DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ParameterNameInStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ParameterNameInStore
	}).(pulumi.StringPtrOutput)
}

func (o DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput) ParameterStoreRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefenderForServersAwsOfferingServicePrincipalSecretMetadata) *string {
		if v == nil {
			return nil
		}
		return v.ParameterStoreRegion
	}).(pulumi.StringPtrOutput)
}

type DenylistCustomAlertRule struct {
	DenylistValues []string `pulumi:"denylistValues"`
	IsEnabled      bool     `pulumi:"isEnabled"`
	RuleType       string   `pulumi:"ruleType"`
}





type DenylistCustomAlertRuleInput interface {
	pulumi.Input

	ToDenylistCustomAlertRuleOutput() DenylistCustomAlertRuleOutput
	ToDenylistCustomAlertRuleOutputWithContext(context.Context) DenylistCustomAlertRuleOutput
}

type DenylistCustomAlertRuleArgs struct {
	DenylistValues pulumi.StringArrayInput `pulumi:"denylistValues"`
	IsEnabled      pulumi.BoolInput        `pulumi:"isEnabled"`
	RuleType       pulumi.StringInput      `pulumi:"ruleType"`
}

func (DenylistCustomAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DenylistCustomAlertRule)(nil)).Elem()
}

func (i DenylistCustomAlertRuleArgs) ToDenylistCustomAlertRuleOutput() DenylistCustomAlertRuleOutput {
	return i.ToDenylistCustomAlertRuleOutputWithContext(context.Background())
}

func (i DenylistCustomAlertRuleArgs) ToDenylistCustomAlertRuleOutputWithContext(ctx context.Context) DenylistCustomAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DenylistCustomAlertRuleOutput)
}





type DenylistCustomAlertRuleArrayInput interface {
	pulumi.Input

	ToDenylistCustomAlertRuleArrayOutput() DenylistCustomAlertRuleArrayOutput
	ToDenylistCustomAlertRuleArrayOutputWithContext(context.Context) DenylistCustomAlertRuleArrayOutput
}

type DenylistCustomAlertRuleArray []DenylistCustomAlertRuleInput

func (DenylistCustomAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DenylistCustomAlertRule)(nil)).Elem()
}

func (i DenylistCustomAlertRuleArray) ToDenylistCustomAlertRuleArrayOutput() DenylistCustomAlertRuleArrayOutput {
	return i.ToDenylistCustomAlertRuleArrayOutputWithContext(context.Background())
}

func (i DenylistCustomAlertRuleArray) ToDenylistCustomAlertRuleArrayOutputWithContext(ctx context.Context) DenylistCustomAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DenylistCustomAlertRuleArrayOutput)
}

type DenylistCustomAlertRuleOutput struct{ *pulumi.OutputState }

func (DenylistCustomAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DenylistCustomAlertRule)(nil)).Elem()
}

func (o DenylistCustomAlertRuleOutput) ToDenylistCustomAlertRuleOutput() DenylistCustomAlertRuleOutput {
	return o
}

func (o DenylistCustomAlertRuleOutput) ToDenylistCustomAlertRuleOutputWithContext(ctx context.Context) DenylistCustomAlertRuleOutput {
	return o
}

func (o DenylistCustomAlertRuleOutput) DenylistValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DenylistCustomAlertRule) []string { return v.DenylistValues }).(pulumi.StringArrayOutput)
}

func (o DenylistCustomAlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DenylistCustomAlertRule) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o DenylistCustomAlertRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type DenylistCustomAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (DenylistCustomAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DenylistCustomAlertRule)(nil)).Elem()
}

func (o DenylistCustomAlertRuleArrayOutput) ToDenylistCustomAlertRuleArrayOutput() DenylistCustomAlertRuleArrayOutput {
	return o
}

func (o DenylistCustomAlertRuleArrayOutput) ToDenylistCustomAlertRuleArrayOutputWithContext(ctx context.Context) DenylistCustomAlertRuleArrayOutput {
	return o
}

func (o DenylistCustomAlertRuleArrayOutput) Index(i pulumi.IntInput) DenylistCustomAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DenylistCustomAlertRule {
		return vs[0].([]DenylistCustomAlertRule)[vs[1].(int)]
	}).(DenylistCustomAlertRuleOutput)
}

type DenylistCustomAlertRuleResponse struct {
	DenylistValues []string `pulumi:"denylistValues"`
	Description    string   `pulumi:"description"`
	DisplayName    string   `pulumi:"displayName"`
	IsEnabled      bool     `pulumi:"isEnabled"`
	RuleType       string   `pulumi:"ruleType"`
	ValueType      string   `pulumi:"valueType"`
}





type DenylistCustomAlertRuleResponseInput interface {
	pulumi.Input

	ToDenylistCustomAlertRuleResponseOutput() DenylistCustomAlertRuleResponseOutput
	ToDenylistCustomAlertRuleResponseOutputWithContext(context.Context) DenylistCustomAlertRuleResponseOutput
}

type DenylistCustomAlertRuleResponseArgs struct {
	DenylistValues pulumi.StringArrayInput `pulumi:"denylistValues"`
	Description    pulumi.StringInput      `pulumi:"description"`
	DisplayName    pulumi.StringInput      `pulumi:"displayName"`
	IsEnabled      pulumi.BoolInput        `pulumi:"isEnabled"`
	RuleType       pulumi.StringInput      `pulumi:"ruleType"`
	ValueType      pulumi.StringInput      `pulumi:"valueType"`
}

func (DenylistCustomAlertRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DenylistCustomAlertRuleResponse)(nil)).Elem()
}

func (i DenylistCustomAlertRuleResponseArgs) ToDenylistCustomAlertRuleResponseOutput() DenylistCustomAlertRuleResponseOutput {
	return i.ToDenylistCustomAlertRuleResponseOutputWithContext(context.Background())
}

func (i DenylistCustomAlertRuleResponseArgs) ToDenylistCustomAlertRuleResponseOutputWithContext(ctx context.Context) DenylistCustomAlertRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DenylistCustomAlertRuleResponseOutput)
}





type DenylistCustomAlertRuleResponseArrayInput interface {
	pulumi.Input

	ToDenylistCustomAlertRuleResponseArrayOutput() DenylistCustomAlertRuleResponseArrayOutput
	ToDenylistCustomAlertRuleResponseArrayOutputWithContext(context.Context) DenylistCustomAlertRuleResponseArrayOutput
}

type DenylistCustomAlertRuleResponseArray []DenylistCustomAlertRuleResponseInput

func (DenylistCustomAlertRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DenylistCustomAlertRuleResponse)(nil)).Elem()
}

func (i DenylistCustomAlertRuleResponseArray) ToDenylistCustomAlertRuleResponseArrayOutput() DenylistCustomAlertRuleResponseArrayOutput {
	return i.ToDenylistCustomAlertRuleResponseArrayOutputWithContext(context.Background())
}

func (i DenylistCustomAlertRuleResponseArray) ToDenylistCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) DenylistCustomAlertRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DenylistCustomAlertRuleResponseArrayOutput)
}

type DenylistCustomAlertRuleResponseOutput struct{ *pulumi.OutputState }

func (DenylistCustomAlertRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DenylistCustomAlertRuleResponse)(nil)).Elem()
}

func (o DenylistCustomAlertRuleResponseOutput) ToDenylistCustomAlertRuleResponseOutput() DenylistCustomAlertRuleResponseOutput {
	return o
}

func (o DenylistCustomAlertRuleResponseOutput) ToDenylistCustomAlertRuleResponseOutputWithContext(ctx context.Context) DenylistCustomAlertRuleResponseOutput {
	return o
}

func (o DenylistCustomAlertRuleResponseOutput) DenylistValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) []string { return v.DenylistValues }).(pulumi.StringArrayOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o DenylistCustomAlertRuleResponseOutput) ValueType() pulumi.StringOutput {
	return o.ApplyT(func(v DenylistCustomAlertRuleResponse) string { return v.ValueType }).(pulumi.StringOutput)
}

type DenylistCustomAlertRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (DenylistCustomAlertRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DenylistCustomAlertRuleResponse)(nil)).Elem()
}

func (o DenylistCustomAlertRuleResponseArrayOutput) ToDenylistCustomAlertRuleResponseArrayOutput() DenylistCustomAlertRuleResponseArrayOutput {
	return o
}

func (o DenylistCustomAlertRuleResponseArrayOutput) ToDenylistCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) DenylistCustomAlertRuleResponseArrayOutput {
	return o
}

func (o DenylistCustomAlertRuleResponseArrayOutput) Index(i pulumi.IntInput) DenylistCustomAlertRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DenylistCustomAlertRuleResponse {
		return vs[0].([]DenylistCustomAlertRuleResponse)[vs[1].(int)]
	}).(DenylistCustomAlertRuleResponseOutput)
}

type GcpCredentialsDetailsProperties struct {
	AuthProviderX509CertUrl string `pulumi:"authProviderX509CertUrl"`
	AuthUri                 string `pulumi:"authUri"`
	AuthenticationType      string `pulumi:"authenticationType"`
	ClientEmail             string `pulumi:"clientEmail"`
	ClientId                string `pulumi:"clientId"`
	ClientX509CertUrl       string `pulumi:"clientX509CertUrl"`
	OrganizationId          string `pulumi:"organizationId"`
	PrivateKey              string `pulumi:"privateKey"`
	PrivateKeyId            string `pulumi:"privateKeyId"`
	ProjectId               string `pulumi:"projectId"`
	TokenUri                string `pulumi:"tokenUri"`
	Type                    string `pulumi:"type"`
}





type GcpCredentialsDetailsPropertiesInput interface {
	pulumi.Input

	ToGcpCredentialsDetailsPropertiesOutput() GcpCredentialsDetailsPropertiesOutput
	ToGcpCredentialsDetailsPropertiesOutputWithContext(context.Context) GcpCredentialsDetailsPropertiesOutput
}

type GcpCredentialsDetailsPropertiesArgs struct {
	AuthProviderX509CertUrl pulumi.StringInput `pulumi:"authProviderX509CertUrl"`
	AuthUri                 pulumi.StringInput `pulumi:"authUri"`
	AuthenticationType      pulumi.StringInput `pulumi:"authenticationType"`
	ClientEmail             pulumi.StringInput `pulumi:"clientEmail"`
	ClientId                pulumi.StringInput `pulumi:"clientId"`
	ClientX509CertUrl       pulumi.StringInput `pulumi:"clientX509CertUrl"`
	OrganizationId          pulumi.StringInput `pulumi:"organizationId"`
	PrivateKey              pulumi.StringInput `pulumi:"privateKey"`
	PrivateKeyId            pulumi.StringInput `pulumi:"privateKeyId"`
	ProjectId               pulumi.StringInput `pulumi:"projectId"`
	TokenUri                pulumi.StringInput `pulumi:"tokenUri"`
	Type                    pulumi.StringInput `pulumi:"type"`
}

func (GcpCredentialsDetailsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GcpCredentialsDetailsProperties)(nil)).Elem()
}

func (i GcpCredentialsDetailsPropertiesArgs) ToGcpCredentialsDetailsPropertiesOutput() GcpCredentialsDetailsPropertiesOutput {
	return i.ToGcpCredentialsDetailsPropertiesOutputWithContext(context.Background())
}

func (i GcpCredentialsDetailsPropertiesArgs) ToGcpCredentialsDetailsPropertiesOutputWithContext(ctx context.Context) GcpCredentialsDetailsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpCredentialsDetailsPropertiesOutput)
}

type GcpCredentialsDetailsPropertiesOutput struct{ *pulumi.OutputState }

func (GcpCredentialsDetailsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GcpCredentialsDetailsProperties)(nil)).Elem()
}

func (o GcpCredentialsDetailsPropertiesOutput) ToGcpCredentialsDetailsPropertiesOutput() GcpCredentialsDetailsPropertiesOutput {
	return o
}

func (o GcpCredentialsDetailsPropertiesOutput) ToGcpCredentialsDetailsPropertiesOutputWithContext(ctx context.Context) GcpCredentialsDetailsPropertiesOutput {
	return o
}

func (o GcpCredentialsDetailsPropertiesOutput) AuthProviderX509CertUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.AuthProviderX509CertUrl }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) AuthUri() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.AuthUri }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) ClientEmail() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.ClientEmail }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) ClientX509CertUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.ClientX509CertUrl }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) PrivateKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.PrivateKeyId }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) TokenUri() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.TokenUri }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsProperties) string { return v.Type }).(pulumi.StringOutput)
}

type GcpCredentialsDetailsPropertiesResponse struct {
	AuthProviderX509CertUrl         string   `pulumi:"authProviderX509CertUrl"`
	AuthUri                         string   `pulumi:"authUri"`
	AuthenticationProvisioningState string   `pulumi:"authenticationProvisioningState"`
	AuthenticationType              string   `pulumi:"authenticationType"`
	ClientEmail                     string   `pulumi:"clientEmail"`
	ClientId                        string   `pulumi:"clientId"`
	ClientX509CertUrl               string   `pulumi:"clientX509CertUrl"`
	GrantedPermissions              []string `pulumi:"grantedPermissions"`
	OrganizationId                  string   `pulumi:"organizationId"`
	PrivateKey                      string   `pulumi:"privateKey"`
	PrivateKeyId                    string   `pulumi:"privateKeyId"`
	ProjectId                       string   `pulumi:"projectId"`
	TokenUri                        string   `pulumi:"tokenUri"`
	Type                            string   `pulumi:"type"`
}





type GcpCredentialsDetailsPropertiesResponseInput interface {
	pulumi.Input

	ToGcpCredentialsDetailsPropertiesResponseOutput() GcpCredentialsDetailsPropertiesResponseOutput
	ToGcpCredentialsDetailsPropertiesResponseOutputWithContext(context.Context) GcpCredentialsDetailsPropertiesResponseOutput
}

type GcpCredentialsDetailsPropertiesResponseArgs struct {
	AuthProviderX509CertUrl         pulumi.StringInput      `pulumi:"authProviderX509CertUrl"`
	AuthUri                         pulumi.StringInput      `pulumi:"authUri"`
	AuthenticationProvisioningState pulumi.StringInput      `pulumi:"authenticationProvisioningState"`
	AuthenticationType              pulumi.StringInput      `pulumi:"authenticationType"`
	ClientEmail                     pulumi.StringInput      `pulumi:"clientEmail"`
	ClientId                        pulumi.StringInput      `pulumi:"clientId"`
	ClientX509CertUrl               pulumi.StringInput      `pulumi:"clientX509CertUrl"`
	GrantedPermissions              pulumi.StringArrayInput `pulumi:"grantedPermissions"`
	OrganizationId                  pulumi.StringInput      `pulumi:"organizationId"`
	PrivateKey                      pulumi.StringInput      `pulumi:"privateKey"`
	PrivateKeyId                    pulumi.StringInput      `pulumi:"privateKeyId"`
	ProjectId                       pulumi.StringInput      `pulumi:"projectId"`
	TokenUri                        pulumi.StringInput      `pulumi:"tokenUri"`
	Type                            pulumi.StringInput      `pulumi:"type"`
}

func (GcpCredentialsDetailsPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GcpCredentialsDetailsPropertiesResponse)(nil)).Elem()
}

func (i GcpCredentialsDetailsPropertiesResponseArgs) ToGcpCredentialsDetailsPropertiesResponseOutput() GcpCredentialsDetailsPropertiesResponseOutput {
	return i.ToGcpCredentialsDetailsPropertiesResponseOutputWithContext(context.Background())
}

func (i GcpCredentialsDetailsPropertiesResponseArgs) ToGcpCredentialsDetailsPropertiesResponseOutputWithContext(ctx context.Context) GcpCredentialsDetailsPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpCredentialsDetailsPropertiesResponseOutput)
}

type GcpCredentialsDetailsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GcpCredentialsDetailsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GcpCredentialsDetailsPropertiesResponse)(nil)).Elem()
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) ToGcpCredentialsDetailsPropertiesResponseOutput() GcpCredentialsDetailsPropertiesResponseOutput {
	return o
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) ToGcpCredentialsDetailsPropertiesResponseOutputWithContext(ctx context.Context) GcpCredentialsDetailsPropertiesResponseOutput {
	return o
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) AuthProviderX509CertUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.AuthProviderX509CertUrl }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) AuthUri() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.AuthUri }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) AuthenticationProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.AuthenticationProvisioningState }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) ClientEmail() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.ClientEmail }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) ClientX509CertUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.ClientX509CertUrl }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) GrantedPermissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) []string { return v.GrantedPermissions }).(pulumi.StringArrayOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) PrivateKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.PrivateKeyId }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) TokenUri() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.TokenUri }).(pulumi.StringOutput)
}

func (o GcpCredentialsDetailsPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GcpCredentialsDetailsPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type HybridComputeSettingsProperties struct {
	AutoProvision     string                      `pulumi:"autoProvision"`
	ProxyServer       *ProxyServerProperties      `pulumi:"proxyServer"`
	Region            *string                     `pulumi:"region"`
	ResourceGroupName *string                     `pulumi:"resourceGroupName"`
	ServicePrincipal  *ServicePrincipalProperties `pulumi:"servicePrincipal"`
}





type HybridComputeSettingsPropertiesInput interface {
	pulumi.Input

	ToHybridComputeSettingsPropertiesOutput() HybridComputeSettingsPropertiesOutput
	ToHybridComputeSettingsPropertiesOutputWithContext(context.Context) HybridComputeSettingsPropertiesOutput
}

type HybridComputeSettingsPropertiesArgs struct {
	AutoProvision     pulumi.StringInput                 `pulumi:"autoProvision"`
	ProxyServer       ProxyServerPropertiesPtrInput      `pulumi:"proxyServer"`
	Region            pulumi.StringPtrInput              `pulumi:"region"`
	ResourceGroupName pulumi.StringPtrInput              `pulumi:"resourceGroupName"`
	ServicePrincipal  ServicePrincipalPropertiesPtrInput `pulumi:"servicePrincipal"`
}

func (HybridComputeSettingsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputeSettingsProperties)(nil)).Elem()
}

func (i HybridComputeSettingsPropertiesArgs) ToHybridComputeSettingsPropertiesOutput() HybridComputeSettingsPropertiesOutput {
	return i.ToHybridComputeSettingsPropertiesOutputWithContext(context.Background())
}

func (i HybridComputeSettingsPropertiesArgs) ToHybridComputeSettingsPropertiesOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputeSettingsPropertiesOutput)
}

func (i HybridComputeSettingsPropertiesArgs) ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput {
	return i.ToHybridComputeSettingsPropertiesPtrOutputWithContext(context.Background())
}

func (i HybridComputeSettingsPropertiesArgs) ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputeSettingsPropertiesOutput).ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx)
}









type HybridComputeSettingsPropertiesPtrInput interface {
	pulumi.Input

	ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput
	ToHybridComputeSettingsPropertiesPtrOutputWithContext(context.Context) HybridComputeSettingsPropertiesPtrOutput
}

type hybridComputeSettingsPropertiesPtrType HybridComputeSettingsPropertiesArgs

func HybridComputeSettingsPropertiesPtr(v *HybridComputeSettingsPropertiesArgs) HybridComputeSettingsPropertiesPtrInput {
	return (*hybridComputeSettingsPropertiesPtrType)(v)
}

func (*hybridComputeSettingsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputeSettingsProperties)(nil)).Elem()
}

func (i *hybridComputeSettingsPropertiesPtrType) ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput {
	return i.ToHybridComputeSettingsPropertiesPtrOutputWithContext(context.Background())
}

func (i *hybridComputeSettingsPropertiesPtrType) ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputeSettingsPropertiesPtrOutput)
}

type HybridComputeSettingsPropertiesOutput struct{ *pulumi.OutputState }

func (HybridComputeSettingsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputeSettingsProperties)(nil)).Elem()
}

func (o HybridComputeSettingsPropertiesOutput) ToHybridComputeSettingsPropertiesOutput() HybridComputeSettingsPropertiesOutput {
	return o
}

func (o HybridComputeSettingsPropertiesOutput) ToHybridComputeSettingsPropertiesOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesOutput {
	return o
}

func (o HybridComputeSettingsPropertiesOutput) ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput {
	return o.ToHybridComputeSettingsPropertiesPtrOutputWithContext(context.Background())
}

func (o HybridComputeSettingsPropertiesOutput) ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HybridComputeSettingsProperties) *HybridComputeSettingsProperties {
		return &v
	}).(HybridComputeSettingsPropertiesPtrOutput)
}

func (o HybridComputeSettingsPropertiesOutput) AutoProvision() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) string { return v.AutoProvision }).(pulumi.StringOutput)
}

func (o HybridComputeSettingsPropertiesOutput) ProxyServer() ProxyServerPropertiesPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) *ProxyServerProperties { return v.ProxyServer }).(ProxyServerPropertiesPtrOutput)
}

func (o HybridComputeSettingsPropertiesOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesOutput) ServicePrincipal() ServicePrincipalPropertiesPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsProperties) *ServicePrincipalProperties { return v.ServicePrincipal }).(ServicePrincipalPropertiesPtrOutput)
}

type HybridComputeSettingsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (HybridComputeSettingsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputeSettingsProperties)(nil)).Elem()
}

func (o HybridComputeSettingsPropertiesPtrOutput) ToHybridComputeSettingsPropertiesPtrOutput() HybridComputeSettingsPropertiesPtrOutput {
	return o
}

func (o HybridComputeSettingsPropertiesPtrOutput) ToHybridComputeSettingsPropertiesPtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesPtrOutput {
	return o
}

func (o HybridComputeSettingsPropertiesPtrOutput) Elem() HybridComputeSettingsPropertiesOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) HybridComputeSettingsProperties {
		if v != nil {
			return *v
		}
		var ret HybridComputeSettingsProperties
		return ret
	}).(HybridComputeSettingsPropertiesOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) AutoProvision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AutoProvision
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) ProxyServer() ProxyServerPropertiesPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *ProxyServerProperties {
		if v == nil {
			return nil
		}
		return v.ProxyServer
	}).(ProxyServerPropertiesPtrOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *string {
		if v == nil {
			return nil
		}
		return v.Region
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesPtrOutput) ServicePrincipal() ServicePrincipalPropertiesPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsProperties) *ServicePrincipalProperties {
		if v == nil {
			return nil
		}
		return v.ServicePrincipal
	}).(ServicePrincipalPropertiesPtrOutput)
}

type HybridComputeSettingsPropertiesResponse struct {
	AutoProvision                  string                              `pulumi:"autoProvision"`
	HybridComputeProvisioningState string                              `pulumi:"hybridComputeProvisioningState"`
	ProxyServer                    *ProxyServerPropertiesResponse      `pulumi:"proxyServer"`
	Region                         *string                             `pulumi:"region"`
	ResourceGroupName              *string                             `pulumi:"resourceGroupName"`
	ServicePrincipal               *ServicePrincipalPropertiesResponse `pulumi:"servicePrincipal"`
}





type HybridComputeSettingsPropertiesResponseInput interface {
	pulumi.Input

	ToHybridComputeSettingsPropertiesResponseOutput() HybridComputeSettingsPropertiesResponseOutput
	ToHybridComputeSettingsPropertiesResponseOutputWithContext(context.Context) HybridComputeSettingsPropertiesResponseOutput
}

type HybridComputeSettingsPropertiesResponseArgs struct {
	AutoProvision                  pulumi.StringInput                         `pulumi:"autoProvision"`
	HybridComputeProvisioningState pulumi.StringInput                         `pulumi:"hybridComputeProvisioningState"`
	ProxyServer                    ProxyServerPropertiesResponsePtrInput      `pulumi:"proxyServer"`
	Region                         pulumi.StringPtrInput                      `pulumi:"region"`
	ResourceGroupName              pulumi.StringPtrInput                      `pulumi:"resourceGroupName"`
	ServicePrincipal               ServicePrincipalPropertiesResponsePtrInput `pulumi:"servicePrincipal"`
}

func (HybridComputeSettingsPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputeSettingsPropertiesResponse)(nil)).Elem()
}

func (i HybridComputeSettingsPropertiesResponseArgs) ToHybridComputeSettingsPropertiesResponseOutput() HybridComputeSettingsPropertiesResponseOutput {
	return i.ToHybridComputeSettingsPropertiesResponseOutputWithContext(context.Background())
}

func (i HybridComputeSettingsPropertiesResponseArgs) ToHybridComputeSettingsPropertiesResponseOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputeSettingsPropertiesResponseOutput)
}

func (i HybridComputeSettingsPropertiesResponseArgs) ToHybridComputeSettingsPropertiesResponsePtrOutput() HybridComputeSettingsPropertiesResponsePtrOutput {
	return i.ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i HybridComputeSettingsPropertiesResponseArgs) ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputeSettingsPropertiesResponseOutput).ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(ctx)
}









type HybridComputeSettingsPropertiesResponsePtrInput interface {
	pulumi.Input

	ToHybridComputeSettingsPropertiesResponsePtrOutput() HybridComputeSettingsPropertiesResponsePtrOutput
	ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(context.Context) HybridComputeSettingsPropertiesResponsePtrOutput
}

type hybridComputeSettingsPropertiesResponsePtrType HybridComputeSettingsPropertiesResponseArgs

func HybridComputeSettingsPropertiesResponsePtr(v *HybridComputeSettingsPropertiesResponseArgs) HybridComputeSettingsPropertiesResponsePtrInput {
	return (*hybridComputeSettingsPropertiesResponsePtrType)(v)
}

func (*hybridComputeSettingsPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputeSettingsPropertiesResponse)(nil)).Elem()
}

func (i *hybridComputeSettingsPropertiesResponsePtrType) ToHybridComputeSettingsPropertiesResponsePtrOutput() HybridComputeSettingsPropertiesResponsePtrOutput {
	return i.ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *hybridComputeSettingsPropertiesResponsePtrType) ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridComputeSettingsPropertiesResponsePtrOutput)
}

type HybridComputeSettingsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (HybridComputeSettingsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridComputeSettingsPropertiesResponse)(nil)).Elem()
}

func (o HybridComputeSettingsPropertiesResponseOutput) ToHybridComputeSettingsPropertiesResponseOutput() HybridComputeSettingsPropertiesResponseOutput {
	return o
}

func (o HybridComputeSettingsPropertiesResponseOutput) ToHybridComputeSettingsPropertiesResponseOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesResponseOutput {
	return o
}

func (o HybridComputeSettingsPropertiesResponseOutput) ToHybridComputeSettingsPropertiesResponsePtrOutput() HybridComputeSettingsPropertiesResponsePtrOutput {
	return o.ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o HybridComputeSettingsPropertiesResponseOutput) ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HybridComputeSettingsPropertiesResponse) *HybridComputeSettingsPropertiesResponse {
		return &v
	}).(HybridComputeSettingsPropertiesResponsePtrOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) AutoProvision() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) string { return v.AutoProvision }).(pulumi.StringOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) HybridComputeProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) string { return v.HybridComputeProvisioningState }).(pulumi.StringOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) ProxyServer() ProxyServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) *ProxyServerPropertiesResponse { return v.ProxyServer }).(ProxyServerPropertiesResponsePtrOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponseOutput) ServicePrincipal() ServicePrincipalPropertiesResponsePtrOutput {
	return o.ApplyT(func(v HybridComputeSettingsPropertiesResponse) *ServicePrincipalPropertiesResponse {
		return v.ServicePrincipal
	}).(ServicePrincipalPropertiesResponsePtrOutput)
}

type HybridComputeSettingsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (HybridComputeSettingsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridComputeSettingsPropertiesResponse)(nil)).Elem()
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ToHybridComputeSettingsPropertiesResponsePtrOutput() HybridComputeSettingsPropertiesResponsePtrOutput {
	return o
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ToHybridComputeSettingsPropertiesResponsePtrOutputWithContext(ctx context.Context) HybridComputeSettingsPropertiesResponsePtrOutput {
	return o
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) Elem() HybridComputeSettingsPropertiesResponseOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) HybridComputeSettingsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret HybridComputeSettingsPropertiesResponse
		return ret
	}).(HybridComputeSettingsPropertiesResponseOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) AutoProvision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AutoProvision
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) HybridComputeProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HybridComputeProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ProxyServer() ProxyServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *ProxyServerPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ProxyServer
	}).(ProxyServerPropertiesResponsePtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Region
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o HybridComputeSettingsPropertiesResponsePtrOutput) ServicePrincipal() ServicePrincipalPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *HybridComputeSettingsPropertiesResponse) *ServicePrincipalPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ServicePrincipal
	}).(ServicePrincipalPropertiesResponsePtrOutput)
}

type IngestionConnectionStringResponse struct {
	Location string `pulumi:"location"`
	Value    string `pulumi:"value"`
}





type IngestionConnectionStringResponseInput interface {
	pulumi.Input

	ToIngestionConnectionStringResponseOutput() IngestionConnectionStringResponseOutput
	ToIngestionConnectionStringResponseOutputWithContext(context.Context) IngestionConnectionStringResponseOutput
}

type IngestionConnectionStringResponseArgs struct {
	Location pulumi.StringInput `pulumi:"location"`
	Value    pulumi.StringInput `pulumi:"value"`
}

func (IngestionConnectionStringResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionConnectionStringResponse)(nil)).Elem()
}

func (i IngestionConnectionStringResponseArgs) ToIngestionConnectionStringResponseOutput() IngestionConnectionStringResponseOutput {
	return i.ToIngestionConnectionStringResponseOutputWithContext(context.Background())
}

func (i IngestionConnectionStringResponseArgs) ToIngestionConnectionStringResponseOutputWithContext(ctx context.Context) IngestionConnectionStringResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionConnectionStringResponseOutput)
}





type IngestionConnectionStringResponseArrayInput interface {
	pulumi.Input

	ToIngestionConnectionStringResponseArrayOutput() IngestionConnectionStringResponseArrayOutput
	ToIngestionConnectionStringResponseArrayOutputWithContext(context.Context) IngestionConnectionStringResponseArrayOutput
}

type IngestionConnectionStringResponseArray []IngestionConnectionStringResponseInput

func (IngestionConnectionStringResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngestionConnectionStringResponse)(nil)).Elem()
}

func (i IngestionConnectionStringResponseArray) ToIngestionConnectionStringResponseArrayOutput() IngestionConnectionStringResponseArrayOutput {
	return i.ToIngestionConnectionStringResponseArrayOutputWithContext(context.Background())
}

func (i IngestionConnectionStringResponseArray) ToIngestionConnectionStringResponseArrayOutputWithContext(ctx context.Context) IngestionConnectionStringResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngestionConnectionStringResponseArrayOutput)
}

type IngestionConnectionStringResponseOutput struct{ *pulumi.OutputState }

func (IngestionConnectionStringResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngestionConnectionStringResponse)(nil)).Elem()
}

func (o IngestionConnectionStringResponseOutput) ToIngestionConnectionStringResponseOutput() IngestionConnectionStringResponseOutput {
	return o
}

func (o IngestionConnectionStringResponseOutput) ToIngestionConnectionStringResponseOutputWithContext(ctx context.Context) IngestionConnectionStringResponseOutput {
	return o
}

func (o IngestionConnectionStringResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v IngestionConnectionStringResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o IngestionConnectionStringResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v IngestionConnectionStringResponse) string { return v.Value }).(pulumi.StringOutput)
}

type IngestionConnectionStringResponseArrayOutput struct{ *pulumi.OutputState }

func (IngestionConnectionStringResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IngestionConnectionStringResponse)(nil)).Elem()
}

func (o IngestionConnectionStringResponseArrayOutput) ToIngestionConnectionStringResponseArrayOutput() IngestionConnectionStringResponseArrayOutput {
	return o
}

func (o IngestionConnectionStringResponseArrayOutput) ToIngestionConnectionStringResponseArrayOutputWithContext(ctx context.Context) IngestionConnectionStringResponseArrayOutput {
	return o
}

func (o IngestionConnectionStringResponseArrayOutput) Index(i pulumi.IntInput) IngestionConnectionStringResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IngestionConnectionStringResponse {
		return vs[0].([]IngestionConnectionStringResponse)[vs[1].(int)]
	}).(IngestionConnectionStringResponseOutput)
}

type JitNetworkAccessPolicyVirtualMachine struct {
	Id              string                     `pulumi:"id"`
	Ports           []JitNetworkAccessPortRule `pulumi:"ports"`
	PublicIpAddress *string                    `pulumi:"publicIpAddress"`
}





type JitNetworkAccessPolicyVirtualMachineInput interface {
	pulumi.Input

	ToJitNetworkAccessPolicyVirtualMachineOutput() JitNetworkAccessPolicyVirtualMachineOutput
	ToJitNetworkAccessPolicyVirtualMachineOutputWithContext(context.Context) JitNetworkAccessPolicyVirtualMachineOutput
}

type JitNetworkAccessPolicyVirtualMachineArgs struct {
	Id              pulumi.StringInput                 `pulumi:"id"`
	Ports           JitNetworkAccessPortRuleArrayInput `pulumi:"ports"`
	PublicIpAddress pulumi.StringPtrInput              `pulumi:"publicIpAddress"`
}

func (JitNetworkAccessPolicyVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPolicyVirtualMachine)(nil)).Elem()
}

func (i JitNetworkAccessPolicyVirtualMachineArgs) ToJitNetworkAccessPolicyVirtualMachineOutput() JitNetworkAccessPolicyVirtualMachineOutput {
	return i.ToJitNetworkAccessPolicyVirtualMachineOutputWithContext(context.Background())
}

func (i JitNetworkAccessPolicyVirtualMachineArgs) ToJitNetworkAccessPolicyVirtualMachineOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPolicyVirtualMachineOutput)
}





type JitNetworkAccessPolicyVirtualMachineArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessPolicyVirtualMachineArrayOutput() JitNetworkAccessPolicyVirtualMachineArrayOutput
	ToJitNetworkAccessPolicyVirtualMachineArrayOutputWithContext(context.Context) JitNetworkAccessPolicyVirtualMachineArrayOutput
}

type JitNetworkAccessPolicyVirtualMachineArray []JitNetworkAccessPolicyVirtualMachineInput

func (JitNetworkAccessPolicyVirtualMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPolicyVirtualMachine)(nil)).Elem()
}

func (i JitNetworkAccessPolicyVirtualMachineArray) ToJitNetworkAccessPolicyVirtualMachineArrayOutput() JitNetworkAccessPolicyVirtualMachineArrayOutput {
	return i.ToJitNetworkAccessPolicyVirtualMachineArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessPolicyVirtualMachineArray) ToJitNetworkAccessPolicyVirtualMachineArrayOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPolicyVirtualMachineArrayOutput)
}

type JitNetworkAccessPolicyVirtualMachineOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPolicyVirtualMachine)(nil)).Elem()
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) ToJitNetworkAccessPolicyVirtualMachineOutput() JitNetworkAccessPolicyVirtualMachineOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) ToJitNetworkAccessPolicyVirtualMachineOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachine) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) Ports() JitNetworkAccessPortRuleArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachine) []JitNetworkAccessPortRule { return v.Ports }).(JitNetworkAccessPortRuleArrayOutput)
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachine) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

type JitNetworkAccessPolicyVirtualMachineArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyVirtualMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPolicyVirtualMachine)(nil)).Elem()
}

func (o JitNetworkAccessPolicyVirtualMachineArrayOutput) ToJitNetworkAccessPolicyVirtualMachineArrayOutput() JitNetworkAccessPolicyVirtualMachineArrayOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineArrayOutput) ToJitNetworkAccessPolicyVirtualMachineArrayOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineArrayOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessPolicyVirtualMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessPolicyVirtualMachine {
		return vs[0].([]JitNetworkAccessPolicyVirtualMachine)[vs[1].(int)]
	}).(JitNetworkAccessPolicyVirtualMachineOutput)
}

type JitNetworkAccessPolicyVirtualMachineResponse struct {
	Id              string                             `pulumi:"id"`
	Ports           []JitNetworkAccessPortRuleResponse `pulumi:"ports"`
	PublicIpAddress *string                            `pulumi:"publicIpAddress"`
}





type JitNetworkAccessPolicyVirtualMachineResponseInput interface {
	pulumi.Input

	ToJitNetworkAccessPolicyVirtualMachineResponseOutput() JitNetworkAccessPolicyVirtualMachineResponseOutput
	ToJitNetworkAccessPolicyVirtualMachineResponseOutputWithContext(context.Context) JitNetworkAccessPolicyVirtualMachineResponseOutput
}

type JitNetworkAccessPolicyVirtualMachineResponseArgs struct {
	Id              pulumi.StringInput                         `pulumi:"id"`
	Ports           JitNetworkAccessPortRuleResponseArrayInput `pulumi:"ports"`
	PublicIpAddress pulumi.StringPtrInput                      `pulumi:"publicIpAddress"`
}

func (JitNetworkAccessPolicyVirtualMachineResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPolicyVirtualMachineResponse)(nil)).Elem()
}

func (i JitNetworkAccessPolicyVirtualMachineResponseArgs) ToJitNetworkAccessPolicyVirtualMachineResponseOutput() JitNetworkAccessPolicyVirtualMachineResponseOutput {
	return i.ToJitNetworkAccessPolicyVirtualMachineResponseOutputWithContext(context.Background())
}

func (i JitNetworkAccessPolicyVirtualMachineResponseArgs) ToJitNetworkAccessPolicyVirtualMachineResponseOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPolicyVirtualMachineResponseOutput)
}





type JitNetworkAccessPolicyVirtualMachineResponseArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessPolicyVirtualMachineResponseArrayOutput() JitNetworkAccessPolicyVirtualMachineResponseArrayOutput
	ToJitNetworkAccessPolicyVirtualMachineResponseArrayOutputWithContext(context.Context) JitNetworkAccessPolicyVirtualMachineResponseArrayOutput
}

type JitNetworkAccessPolicyVirtualMachineResponseArray []JitNetworkAccessPolicyVirtualMachineResponseInput

func (JitNetworkAccessPolicyVirtualMachineResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPolicyVirtualMachineResponse)(nil)).Elem()
}

func (i JitNetworkAccessPolicyVirtualMachineResponseArray) ToJitNetworkAccessPolicyVirtualMachineResponseArrayOutput() JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return i.ToJitNetworkAccessPolicyVirtualMachineResponseArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessPolicyVirtualMachineResponseArray) ToJitNetworkAccessPolicyVirtualMachineResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPolicyVirtualMachineResponseArrayOutput)
}

type JitNetworkAccessPolicyVirtualMachineResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyVirtualMachineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPolicyVirtualMachineResponse)(nil)).Elem()
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) ToJitNetworkAccessPolicyVirtualMachineResponseOutput() JitNetworkAccessPolicyVirtualMachineResponseOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) ToJitNetworkAccessPolicyVirtualMachineResponseOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineResponseOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachineResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) Ports() JitNetworkAccessPortRuleResponseArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachineResponse) []JitNetworkAccessPortRuleResponse {
		return v.Ports
	}).(JitNetworkAccessPortRuleResponseArrayOutput)
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachineResponse) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

type JitNetworkAccessPolicyVirtualMachineResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyVirtualMachineResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPolicyVirtualMachineResponse)(nil)).Elem()
}

func (o JitNetworkAccessPolicyVirtualMachineResponseArrayOutput) ToJitNetworkAccessPolicyVirtualMachineResponseArrayOutput() JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineResponseArrayOutput) ToJitNetworkAccessPolicyVirtualMachineResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessPolicyVirtualMachineResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessPolicyVirtualMachineResponse {
		return vs[0].([]JitNetworkAccessPolicyVirtualMachineResponse)[vs[1].(int)]
	}).(JitNetworkAccessPolicyVirtualMachineResponseOutput)
}

type JitNetworkAccessPortRule struct {
	AllowedSourceAddressPrefix   *string  `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes []string `pulumi:"allowedSourceAddressPrefixes"`
	MaxRequestAccessDuration     string   `pulumi:"maxRequestAccessDuration"`
	Number                       int      `pulumi:"number"`
	Protocol                     string   `pulumi:"protocol"`
}





type JitNetworkAccessPortRuleInput interface {
	pulumi.Input

	ToJitNetworkAccessPortRuleOutput() JitNetworkAccessPortRuleOutput
	ToJitNetworkAccessPortRuleOutputWithContext(context.Context) JitNetworkAccessPortRuleOutput
}

type JitNetworkAccessPortRuleArgs struct {
	AllowedSourceAddressPrefix   pulumi.StringPtrInput   `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes pulumi.StringArrayInput `pulumi:"allowedSourceAddressPrefixes"`
	MaxRequestAccessDuration     pulumi.StringInput      `pulumi:"maxRequestAccessDuration"`
	Number                       pulumi.IntInput         `pulumi:"number"`
	Protocol                     pulumi.StringInput      `pulumi:"protocol"`
}

func (JitNetworkAccessPortRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPortRule)(nil)).Elem()
}

func (i JitNetworkAccessPortRuleArgs) ToJitNetworkAccessPortRuleOutput() JitNetworkAccessPortRuleOutput {
	return i.ToJitNetworkAccessPortRuleOutputWithContext(context.Background())
}

func (i JitNetworkAccessPortRuleArgs) ToJitNetworkAccessPortRuleOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPortRuleOutput)
}





type JitNetworkAccessPortRuleArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessPortRuleArrayOutput() JitNetworkAccessPortRuleArrayOutput
	ToJitNetworkAccessPortRuleArrayOutputWithContext(context.Context) JitNetworkAccessPortRuleArrayOutput
}

type JitNetworkAccessPortRuleArray []JitNetworkAccessPortRuleInput

func (JitNetworkAccessPortRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPortRule)(nil)).Elem()
}

func (i JitNetworkAccessPortRuleArray) ToJitNetworkAccessPortRuleArrayOutput() JitNetworkAccessPortRuleArrayOutput {
	return i.ToJitNetworkAccessPortRuleArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessPortRuleArray) ToJitNetworkAccessPortRuleArrayOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPortRuleArrayOutput)
}

type JitNetworkAccessPortRuleOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPortRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPortRule)(nil)).Elem()
}

func (o JitNetworkAccessPortRuleOutput) ToJitNetworkAccessPortRuleOutput() JitNetworkAccessPortRuleOutput {
	return o
}

func (o JitNetworkAccessPortRuleOutput) ToJitNetworkAccessPortRuleOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleOutput {
	return o
}

func (o JitNetworkAccessPortRuleOutput) AllowedSourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) *string { return v.AllowedSourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessPortRuleOutput) AllowedSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) []string { return v.AllowedSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o JitNetworkAccessPortRuleOutput) MaxRequestAccessDuration() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) string { return v.MaxRequestAccessDuration }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPortRuleOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) int { return v.Number }).(pulumi.IntOutput)
}

func (o JitNetworkAccessPortRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) string { return v.Protocol }).(pulumi.StringOutput)
}

type JitNetworkAccessPortRuleArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPortRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPortRule)(nil)).Elem()
}

func (o JitNetworkAccessPortRuleArrayOutput) ToJitNetworkAccessPortRuleArrayOutput() JitNetworkAccessPortRuleArrayOutput {
	return o
}

func (o JitNetworkAccessPortRuleArrayOutput) ToJitNetworkAccessPortRuleArrayOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleArrayOutput {
	return o
}

func (o JitNetworkAccessPortRuleArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessPortRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessPortRule {
		return vs[0].([]JitNetworkAccessPortRule)[vs[1].(int)]
	}).(JitNetworkAccessPortRuleOutput)
}

type JitNetworkAccessPortRuleResponse struct {
	AllowedSourceAddressPrefix   *string  `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes []string `pulumi:"allowedSourceAddressPrefixes"`
	MaxRequestAccessDuration     string   `pulumi:"maxRequestAccessDuration"`
	Number                       int      `pulumi:"number"`
	Protocol                     string   `pulumi:"protocol"`
}





type JitNetworkAccessPortRuleResponseInput interface {
	pulumi.Input

	ToJitNetworkAccessPortRuleResponseOutput() JitNetworkAccessPortRuleResponseOutput
	ToJitNetworkAccessPortRuleResponseOutputWithContext(context.Context) JitNetworkAccessPortRuleResponseOutput
}

type JitNetworkAccessPortRuleResponseArgs struct {
	AllowedSourceAddressPrefix   pulumi.StringPtrInput   `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes pulumi.StringArrayInput `pulumi:"allowedSourceAddressPrefixes"`
	MaxRequestAccessDuration     pulumi.StringInput      `pulumi:"maxRequestAccessDuration"`
	Number                       pulumi.IntInput         `pulumi:"number"`
	Protocol                     pulumi.StringInput      `pulumi:"protocol"`
}

func (JitNetworkAccessPortRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPortRuleResponse)(nil)).Elem()
}

func (i JitNetworkAccessPortRuleResponseArgs) ToJitNetworkAccessPortRuleResponseOutput() JitNetworkAccessPortRuleResponseOutput {
	return i.ToJitNetworkAccessPortRuleResponseOutputWithContext(context.Background())
}

func (i JitNetworkAccessPortRuleResponseArgs) ToJitNetworkAccessPortRuleResponseOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPortRuleResponseOutput)
}





type JitNetworkAccessPortRuleResponseArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessPortRuleResponseArrayOutput() JitNetworkAccessPortRuleResponseArrayOutput
	ToJitNetworkAccessPortRuleResponseArrayOutputWithContext(context.Context) JitNetworkAccessPortRuleResponseArrayOutput
}

type JitNetworkAccessPortRuleResponseArray []JitNetworkAccessPortRuleResponseInput

func (JitNetworkAccessPortRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPortRuleResponse)(nil)).Elem()
}

func (i JitNetworkAccessPortRuleResponseArray) ToJitNetworkAccessPortRuleResponseArrayOutput() JitNetworkAccessPortRuleResponseArrayOutput {
	return i.ToJitNetworkAccessPortRuleResponseArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessPortRuleResponseArray) ToJitNetworkAccessPortRuleResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPortRuleResponseArrayOutput)
}

type JitNetworkAccessPortRuleResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPortRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPortRuleResponse)(nil)).Elem()
}

func (o JitNetworkAccessPortRuleResponseOutput) ToJitNetworkAccessPortRuleResponseOutput() JitNetworkAccessPortRuleResponseOutput {
	return o
}

func (o JitNetworkAccessPortRuleResponseOutput) ToJitNetworkAccessPortRuleResponseOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleResponseOutput {
	return o
}

func (o JitNetworkAccessPortRuleResponseOutput) AllowedSourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) *string { return v.AllowedSourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessPortRuleResponseOutput) AllowedSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) []string { return v.AllowedSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o JitNetworkAccessPortRuleResponseOutput) MaxRequestAccessDuration() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) string { return v.MaxRequestAccessDuration }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPortRuleResponseOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) int { return v.Number }).(pulumi.IntOutput)
}

func (o JitNetworkAccessPortRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

type JitNetworkAccessPortRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPortRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPortRuleResponse)(nil)).Elem()
}

func (o JitNetworkAccessPortRuleResponseArrayOutput) ToJitNetworkAccessPortRuleResponseArrayOutput() JitNetworkAccessPortRuleResponseArrayOutput {
	return o
}

func (o JitNetworkAccessPortRuleResponseArrayOutput) ToJitNetworkAccessPortRuleResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleResponseArrayOutput {
	return o
}

func (o JitNetworkAccessPortRuleResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessPortRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessPortRuleResponse {
		return vs[0].([]JitNetworkAccessPortRuleResponse)[vs[1].(int)]
	}).(JitNetworkAccessPortRuleResponseOutput)
}

type JitNetworkAccessRequest struct {
	Justification   *string                                 `pulumi:"justification"`
	Requestor       string                                  `pulumi:"requestor"`
	StartTimeUtc    string                                  `pulumi:"startTimeUtc"`
	VirtualMachines []JitNetworkAccessRequestVirtualMachine `pulumi:"virtualMachines"`
}





type JitNetworkAccessRequestInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestOutput() JitNetworkAccessRequestOutput
	ToJitNetworkAccessRequestOutputWithContext(context.Context) JitNetworkAccessRequestOutput
}

type JitNetworkAccessRequestArgs struct {
	Justification   pulumi.StringPtrInput                           `pulumi:"justification"`
	Requestor       pulumi.StringInput                              `pulumi:"requestor"`
	StartTimeUtc    pulumi.StringInput                              `pulumi:"startTimeUtc"`
	VirtualMachines JitNetworkAccessRequestVirtualMachineArrayInput `pulumi:"virtualMachines"`
}

func (JitNetworkAccessRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequest)(nil)).Elem()
}

func (i JitNetworkAccessRequestArgs) ToJitNetworkAccessRequestOutput() JitNetworkAccessRequestOutput {
	return i.ToJitNetworkAccessRequestOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestArgs) ToJitNetworkAccessRequestOutputWithContext(ctx context.Context) JitNetworkAccessRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestOutput)
}





type JitNetworkAccessRequestArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestArrayOutput() JitNetworkAccessRequestArrayOutput
	ToJitNetworkAccessRequestArrayOutputWithContext(context.Context) JitNetworkAccessRequestArrayOutput
}

type JitNetworkAccessRequestArray []JitNetworkAccessRequestInput

func (JitNetworkAccessRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequest)(nil)).Elem()
}

func (i JitNetworkAccessRequestArray) ToJitNetworkAccessRequestArrayOutput() JitNetworkAccessRequestArrayOutput {
	return i.ToJitNetworkAccessRequestArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestArray) ToJitNetworkAccessRequestArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestArrayOutput)
}

type JitNetworkAccessRequestOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequest)(nil)).Elem()
}

func (o JitNetworkAccessRequestOutput) ToJitNetworkAccessRequestOutput() JitNetworkAccessRequestOutput {
	return o
}

func (o JitNetworkAccessRequestOutput) ToJitNetworkAccessRequestOutputWithContext(ctx context.Context) JitNetworkAccessRequestOutput {
	return o
}

func (o JitNetworkAccessRequestOutput) Justification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequest) *string { return v.Justification }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessRequestOutput) Requestor() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequest) string { return v.Requestor }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequest) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestOutput) VirtualMachines() JitNetworkAccessRequestVirtualMachineArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequest) []JitNetworkAccessRequestVirtualMachine { return v.VirtualMachines }).(JitNetworkAccessRequestVirtualMachineArrayOutput)
}

type JitNetworkAccessRequestArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequest)(nil)).Elem()
}

func (o JitNetworkAccessRequestArrayOutput) ToJitNetworkAccessRequestArrayOutput() JitNetworkAccessRequestArrayOutput {
	return o
}

func (o JitNetworkAccessRequestArrayOutput) ToJitNetworkAccessRequestArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestArrayOutput {
	return o
}

func (o JitNetworkAccessRequestArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequest {
		return vs[0].([]JitNetworkAccessRequest)[vs[1].(int)]
	}).(JitNetworkAccessRequestOutput)
}

type JitNetworkAccessRequestPort struct {
	AllowedSourceAddressPrefix   *string  `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes []string `pulumi:"allowedSourceAddressPrefixes"`
	EndTimeUtc                   string   `pulumi:"endTimeUtc"`
	MappedPort                   *int     `pulumi:"mappedPort"`
	Number                       int      `pulumi:"number"`
	Status                       string   `pulumi:"status"`
	StatusReason                 string   `pulumi:"statusReason"`
}





type JitNetworkAccessRequestPortInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestPortOutput() JitNetworkAccessRequestPortOutput
	ToJitNetworkAccessRequestPortOutputWithContext(context.Context) JitNetworkAccessRequestPortOutput
}

type JitNetworkAccessRequestPortArgs struct {
	AllowedSourceAddressPrefix   pulumi.StringPtrInput   `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes pulumi.StringArrayInput `pulumi:"allowedSourceAddressPrefixes"`
	EndTimeUtc                   pulumi.StringInput      `pulumi:"endTimeUtc"`
	MappedPort                   pulumi.IntPtrInput      `pulumi:"mappedPort"`
	Number                       pulumi.IntInput         `pulumi:"number"`
	Status                       pulumi.StringInput      `pulumi:"status"`
	StatusReason                 pulumi.StringInput      `pulumi:"statusReason"`
}

func (JitNetworkAccessRequestPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestPort)(nil)).Elem()
}

func (i JitNetworkAccessRequestPortArgs) ToJitNetworkAccessRequestPortOutput() JitNetworkAccessRequestPortOutput {
	return i.ToJitNetworkAccessRequestPortOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestPortArgs) ToJitNetworkAccessRequestPortOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestPortOutput)
}





type JitNetworkAccessRequestPortArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestPortArrayOutput() JitNetworkAccessRequestPortArrayOutput
	ToJitNetworkAccessRequestPortArrayOutputWithContext(context.Context) JitNetworkAccessRequestPortArrayOutput
}

type JitNetworkAccessRequestPortArray []JitNetworkAccessRequestPortInput

func (JitNetworkAccessRequestPortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestPort)(nil)).Elem()
}

func (i JitNetworkAccessRequestPortArray) ToJitNetworkAccessRequestPortArrayOutput() JitNetworkAccessRequestPortArrayOutput {
	return i.ToJitNetworkAccessRequestPortArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestPortArray) ToJitNetworkAccessRequestPortArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestPortArrayOutput)
}

type JitNetworkAccessRequestPortOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestPort)(nil)).Elem()
}

func (o JitNetworkAccessRequestPortOutput) ToJitNetworkAccessRequestPortOutput() JitNetworkAccessRequestPortOutput {
	return o
}

func (o JitNetworkAccessRequestPortOutput) ToJitNetworkAccessRequestPortOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortOutput {
	return o
}

func (o JitNetworkAccessRequestPortOutput) AllowedSourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) *string { return v.AllowedSourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessRequestPortOutput) AllowedSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) []string { return v.AllowedSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o JitNetworkAccessRequestPortOutput) EndTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) string { return v.EndTimeUtc }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestPortOutput) MappedPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) *int { return v.MappedPort }).(pulumi.IntPtrOutput)
}

func (o JitNetworkAccessRequestPortOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) int { return v.Number }).(pulumi.IntOutput)
}

func (o JitNetworkAccessRequestPortOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) string { return v.Status }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestPortOutput) StatusReason() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) string { return v.StatusReason }).(pulumi.StringOutput)
}

type JitNetworkAccessRequestPortArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestPortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestPort)(nil)).Elem()
}

func (o JitNetworkAccessRequestPortArrayOutput) ToJitNetworkAccessRequestPortArrayOutput() JitNetworkAccessRequestPortArrayOutput {
	return o
}

func (o JitNetworkAccessRequestPortArrayOutput) ToJitNetworkAccessRequestPortArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortArrayOutput {
	return o
}

func (o JitNetworkAccessRequestPortArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestPortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestPort {
		return vs[0].([]JitNetworkAccessRequestPort)[vs[1].(int)]
	}).(JitNetworkAccessRequestPortOutput)
}

type JitNetworkAccessRequestPortResponse struct {
	AllowedSourceAddressPrefix   *string  `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes []string `pulumi:"allowedSourceAddressPrefixes"`
	EndTimeUtc                   string   `pulumi:"endTimeUtc"`
	MappedPort                   *int     `pulumi:"mappedPort"`
	Number                       int      `pulumi:"number"`
	Status                       string   `pulumi:"status"`
	StatusReason                 string   `pulumi:"statusReason"`
}





type JitNetworkAccessRequestPortResponseInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestPortResponseOutput() JitNetworkAccessRequestPortResponseOutput
	ToJitNetworkAccessRequestPortResponseOutputWithContext(context.Context) JitNetworkAccessRequestPortResponseOutput
}

type JitNetworkAccessRequestPortResponseArgs struct {
	AllowedSourceAddressPrefix   pulumi.StringPtrInput   `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes pulumi.StringArrayInput `pulumi:"allowedSourceAddressPrefixes"`
	EndTimeUtc                   pulumi.StringInput      `pulumi:"endTimeUtc"`
	MappedPort                   pulumi.IntPtrInput      `pulumi:"mappedPort"`
	Number                       pulumi.IntInput         `pulumi:"number"`
	Status                       pulumi.StringInput      `pulumi:"status"`
	StatusReason                 pulumi.StringInput      `pulumi:"statusReason"`
}

func (JitNetworkAccessRequestPortResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestPortResponse)(nil)).Elem()
}

func (i JitNetworkAccessRequestPortResponseArgs) ToJitNetworkAccessRequestPortResponseOutput() JitNetworkAccessRequestPortResponseOutput {
	return i.ToJitNetworkAccessRequestPortResponseOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestPortResponseArgs) ToJitNetworkAccessRequestPortResponseOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestPortResponseOutput)
}





type JitNetworkAccessRequestPortResponseArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestPortResponseArrayOutput() JitNetworkAccessRequestPortResponseArrayOutput
	ToJitNetworkAccessRequestPortResponseArrayOutputWithContext(context.Context) JitNetworkAccessRequestPortResponseArrayOutput
}

type JitNetworkAccessRequestPortResponseArray []JitNetworkAccessRequestPortResponseInput

func (JitNetworkAccessRequestPortResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestPortResponse)(nil)).Elem()
}

func (i JitNetworkAccessRequestPortResponseArray) ToJitNetworkAccessRequestPortResponseArrayOutput() JitNetworkAccessRequestPortResponseArrayOutput {
	return i.ToJitNetworkAccessRequestPortResponseArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestPortResponseArray) ToJitNetworkAccessRequestPortResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestPortResponseArrayOutput)
}

type JitNetworkAccessRequestPortResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestPortResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestPortResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestPortResponseOutput) ToJitNetworkAccessRequestPortResponseOutput() JitNetworkAccessRequestPortResponseOutput {
	return o
}

func (o JitNetworkAccessRequestPortResponseOutput) ToJitNetworkAccessRequestPortResponseOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortResponseOutput {
	return o
}

func (o JitNetworkAccessRequestPortResponseOutput) AllowedSourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) *string { return v.AllowedSourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) AllowedSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) []string { return v.AllowedSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) EndTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) string { return v.EndTimeUtc }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) MappedPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) *int { return v.MappedPort }).(pulumi.IntPtrOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) int { return v.Number }).(pulumi.IntOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) StatusReason() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) string { return v.StatusReason }).(pulumi.StringOutput)
}

type JitNetworkAccessRequestPortResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestPortResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestPortResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestPortResponseArrayOutput) ToJitNetworkAccessRequestPortResponseArrayOutput() JitNetworkAccessRequestPortResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestPortResponseArrayOutput) ToJitNetworkAccessRequestPortResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestPortResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestPortResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestPortResponse {
		return vs[0].([]JitNetworkAccessRequestPortResponse)[vs[1].(int)]
	}).(JitNetworkAccessRequestPortResponseOutput)
}

type JitNetworkAccessRequestResponse struct {
	Justification   *string                                         `pulumi:"justification"`
	Requestor       string                                          `pulumi:"requestor"`
	StartTimeUtc    string                                          `pulumi:"startTimeUtc"`
	VirtualMachines []JitNetworkAccessRequestVirtualMachineResponse `pulumi:"virtualMachines"`
}





type JitNetworkAccessRequestResponseInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestResponseOutput() JitNetworkAccessRequestResponseOutput
	ToJitNetworkAccessRequestResponseOutputWithContext(context.Context) JitNetworkAccessRequestResponseOutput
}

type JitNetworkAccessRequestResponseArgs struct {
	Justification   pulumi.StringPtrInput                                   `pulumi:"justification"`
	Requestor       pulumi.StringInput                                      `pulumi:"requestor"`
	StartTimeUtc    pulumi.StringInput                                      `pulumi:"startTimeUtc"`
	VirtualMachines JitNetworkAccessRequestVirtualMachineResponseArrayInput `pulumi:"virtualMachines"`
}

func (JitNetworkAccessRequestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestResponse)(nil)).Elem()
}

func (i JitNetworkAccessRequestResponseArgs) ToJitNetworkAccessRequestResponseOutput() JitNetworkAccessRequestResponseOutput {
	return i.ToJitNetworkAccessRequestResponseOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestResponseArgs) ToJitNetworkAccessRequestResponseOutputWithContext(ctx context.Context) JitNetworkAccessRequestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestResponseOutput)
}





type JitNetworkAccessRequestResponseArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestResponseArrayOutput() JitNetworkAccessRequestResponseArrayOutput
	ToJitNetworkAccessRequestResponseArrayOutputWithContext(context.Context) JitNetworkAccessRequestResponseArrayOutput
}

type JitNetworkAccessRequestResponseArray []JitNetworkAccessRequestResponseInput

func (JitNetworkAccessRequestResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestResponse)(nil)).Elem()
}

func (i JitNetworkAccessRequestResponseArray) ToJitNetworkAccessRequestResponseArrayOutput() JitNetworkAccessRequestResponseArrayOutput {
	return i.ToJitNetworkAccessRequestResponseArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestResponseArray) ToJitNetworkAccessRequestResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestResponseArrayOutput)
}

type JitNetworkAccessRequestResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestResponseOutput) ToJitNetworkAccessRequestResponseOutput() JitNetworkAccessRequestResponseOutput {
	return o
}

func (o JitNetworkAccessRequestResponseOutput) ToJitNetworkAccessRequestResponseOutputWithContext(ctx context.Context) JitNetworkAccessRequestResponseOutput {
	return o
}

func (o JitNetworkAccessRequestResponseOutput) Justification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestResponse) *string { return v.Justification }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessRequestResponseOutput) Requestor() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestResponse) string { return v.Requestor }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestResponseOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestResponse) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestResponseOutput) VirtualMachines() JitNetworkAccessRequestVirtualMachineResponseArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestResponse) []JitNetworkAccessRequestVirtualMachineResponse {
		return v.VirtualMachines
	}).(JitNetworkAccessRequestVirtualMachineResponseArrayOutput)
}

type JitNetworkAccessRequestResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestResponseArrayOutput) ToJitNetworkAccessRequestResponseArrayOutput() JitNetworkAccessRequestResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestResponseArrayOutput) ToJitNetworkAccessRequestResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestResponse {
		return vs[0].([]JitNetworkAccessRequestResponse)[vs[1].(int)]
	}).(JitNetworkAccessRequestResponseOutput)
}

type JitNetworkAccessRequestVirtualMachine struct {
	Id    string                        `pulumi:"id"`
	Ports []JitNetworkAccessRequestPort `pulumi:"ports"`
}





type JitNetworkAccessRequestVirtualMachineInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestVirtualMachineOutput() JitNetworkAccessRequestVirtualMachineOutput
	ToJitNetworkAccessRequestVirtualMachineOutputWithContext(context.Context) JitNetworkAccessRequestVirtualMachineOutput
}

type JitNetworkAccessRequestVirtualMachineArgs struct {
	Id    pulumi.StringInput                    `pulumi:"id"`
	Ports JitNetworkAccessRequestPortArrayInput `pulumi:"ports"`
}

func (JitNetworkAccessRequestVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestVirtualMachine)(nil)).Elem()
}

func (i JitNetworkAccessRequestVirtualMachineArgs) ToJitNetworkAccessRequestVirtualMachineOutput() JitNetworkAccessRequestVirtualMachineOutput {
	return i.ToJitNetworkAccessRequestVirtualMachineOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestVirtualMachineArgs) ToJitNetworkAccessRequestVirtualMachineOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestVirtualMachineOutput)
}





type JitNetworkAccessRequestVirtualMachineArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestVirtualMachineArrayOutput() JitNetworkAccessRequestVirtualMachineArrayOutput
	ToJitNetworkAccessRequestVirtualMachineArrayOutputWithContext(context.Context) JitNetworkAccessRequestVirtualMachineArrayOutput
}

type JitNetworkAccessRequestVirtualMachineArray []JitNetworkAccessRequestVirtualMachineInput

func (JitNetworkAccessRequestVirtualMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestVirtualMachine)(nil)).Elem()
}

func (i JitNetworkAccessRequestVirtualMachineArray) ToJitNetworkAccessRequestVirtualMachineArrayOutput() JitNetworkAccessRequestVirtualMachineArrayOutput {
	return i.ToJitNetworkAccessRequestVirtualMachineArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestVirtualMachineArray) ToJitNetworkAccessRequestVirtualMachineArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestVirtualMachineArrayOutput)
}

type JitNetworkAccessRequestVirtualMachineOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestVirtualMachine)(nil)).Elem()
}

func (o JitNetworkAccessRequestVirtualMachineOutput) ToJitNetworkAccessRequestVirtualMachineOutput() JitNetworkAccessRequestVirtualMachineOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineOutput) ToJitNetworkAccessRequestVirtualMachineOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestVirtualMachine) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestVirtualMachineOutput) Ports() JitNetworkAccessRequestPortArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestVirtualMachine) []JitNetworkAccessRequestPort { return v.Ports }).(JitNetworkAccessRequestPortArrayOutput)
}

type JitNetworkAccessRequestVirtualMachineArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestVirtualMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestVirtualMachine)(nil)).Elem()
}

func (o JitNetworkAccessRequestVirtualMachineArrayOutput) ToJitNetworkAccessRequestVirtualMachineArrayOutput() JitNetworkAccessRequestVirtualMachineArrayOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineArrayOutput) ToJitNetworkAccessRequestVirtualMachineArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineArrayOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestVirtualMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestVirtualMachine {
		return vs[0].([]JitNetworkAccessRequestVirtualMachine)[vs[1].(int)]
	}).(JitNetworkAccessRequestVirtualMachineOutput)
}

type JitNetworkAccessRequestVirtualMachineResponse struct {
	Id    string                                `pulumi:"id"`
	Ports []JitNetworkAccessRequestPortResponse `pulumi:"ports"`
}





type JitNetworkAccessRequestVirtualMachineResponseInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestVirtualMachineResponseOutput() JitNetworkAccessRequestVirtualMachineResponseOutput
	ToJitNetworkAccessRequestVirtualMachineResponseOutputWithContext(context.Context) JitNetworkAccessRequestVirtualMachineResponseOutput
}

type JitNetworkAccessRequestVirtualMachineResponseArgs struct {
	Id    pulumi.StringInput                            `pulumi:"id"`
	Ports JitNetworkAccessRequestPortResponseArrayInput `pulumi:"ports"`
}

func (JitNetworkAccessRequestVirtualMachineResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestVirtualMachineResponse)(nil)).Elem()
}

func (i JitNetworkAccessRequestVirtualMachineResponseArgs) ToJitNetworkAccessRequestVirtualMachineResponseOutput() JitNetworkAccessRequestVirtualMachineResponseOutput {
	return i.ToJitNetworkAccessRequestVirtualMachineResponseOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestVirtualMachineResponseArgs) ToJitNetworkAccessRequestVirtualMachineResponseOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestVirtualMachineResponseOutput)
}





type JitNetworkAccessRequestVirtualMachineResponseArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestVirtualMachineResponseArrayOutput() JitNetworkAccessRequestVirtualMachineResponseArrayOutput
	ToJitNetworkAccessRequestVirtualMachineResponseArrayOutputWithContext(context.Context) JitNetworkAccessRequestVirtualMachineResponseArrayOutput
}

type JitNetworkAccessRequestVirtualMachineResponseArray []JitNetworkAccessRequestVirtualMachineResponseInput

func (JitNetworkAccessRequestVirtualMachineResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestVirtualMachineResponse)(nil)).Elem()
}

func (i JitNetworkAccessRequestVirtualMachineResponseArray) ToJitNetworkAccessRequestVirtualMachineResponseArrayOutput() JitNetworkAccessRequestVirtualMachineResponseArrayOutput {
	return i.ToJitNetworkAccessRequestVirtualMachineResponseArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestVirtualMachineResponseArray) ToJitNetworkAccessRequestVirtualMachineResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestVirtualMachineResponseArrayOutput)
}

type JitNetworkAccessRequestVirtualMachineResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestVirtualMachineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestVirtualMachineResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestVirtualMachineResponseOutput) ToJitNetworkAccessRequestVirtualMachineResponseOutput() JitNetworkAccessRequestVirtualMachineResponseOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineResponseOutput) ToJitNetworkAccessRequestVirtualMachineResponseOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineResponseOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestVirtualMachineResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestVirtualMachineResponseOutput) Ports() JitNetworkAccessRequestPortResponseArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestVirtualMachineResponse) []JitNetworkAccessRequestPortResponse {
		return v.Ports
	}).(JitNetworkAccessRequestPortResponseArrayOutput)
}

type JitNetworkAccessRequestVirtualMachineResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestVirtualMachineResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestVirtualMachineResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestVirtualMachineResponseArrayOutput) ToJitNetworkAccessRequestVirtualMachineResponseArrayOutput() JitNetworkAccessRequestVirtualMachineResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineResponseArrayOutput) ToJitNetworkAccessRequestVirtualMachineResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestVirtualMachineResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestVirtualMachineResponse {
		return vs[0].([]JitNetworkAccessRequestVirtualMachineResponse)[vs[1].(int)]
	}).(JitNetworkAccessRequestVirtualMachineResponseOutput)
}

type OnPremiseResourceDetails struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}





type OnPremiseResourceDetailsInput interface {
	pulumi.Input

	ToOnPremiseResourceDetailsOutput() OnPremiseResourceDetailsOutput
	ToOnPremiseResourceDetailsOutputWithContext(context.Context) OnPremiseResourceDetailsOutput
}

type OnPremiseResourceDetailsArgs struct {
	MachineName      pulumi.StringInput `pulumi:"machineName"`
	Source           pulumi.StringInput `pulumi:"source"`
	SourceComputerId pulumi.StringInput `pulumi:"sourceComputerId"`
	Vmuuid           pulumi.StringInput `pulumi:"vmuuid"`
	WorkspaceId      pulumi.StringInput `pulumi:"workspaceId"`
}

func (OnPremiseResourceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseResourceDetails)(nil)).Elem()
}

func (i OnPremiseResourceDetailsArgs) ToOnPremiseResourceDetailsOutput() OnPremiseResourceDetailsOutput {
	return i.ToOnPremiseResourceDetailsOutputWithContext(context.Background())
}

func (i OnPremiseResourceDetailsArgs) ToOnPremiseResourceDetailsOutputWithContext(ctx context.Context) OnPremiseResourceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremiseResourceDetailsOutput)
}

type OnPremiseResourceDetailsOutput struct{ *pulumi.OutputState }

func (OnPremiseResourceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseResourceDetails)(nil)).Elem()
}

func (o OnPremiseResourceDetailsOutput) ToOnPremiseResourceDetailsOutput() OnPremiseResourceDetailsOutput {
	return o
}

func (o OnPremiseResourceDetailsOutput) ToOnPremiseResourceDetailsOutputWithContext(ctx context.Context) OnPremiseResourceDetailsOutput {
	return o
}

func (o OnPremiseResourceDetailsOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.Source }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsOutput) SourceComputerId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.SourceComputerId }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsOutput) Vmuuid() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.Vmuuid }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetails) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

type OnPremiseResourceDetailsResponse struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}





type OnPremiseResourceDetailsResponseInput interface {
	pulumi.Input

	ToOnPremiseResourceDetailsResponseOutput() OnPremiseResourceDetailsResponseOutput
	ToOnPremiseResourceDetailsResponseOutputWithContext(context.Context) OnPremiseResourceDetailsResponseOutput
}

type OnPremiseResourceDetailsResponseArgs struct {
	MachineName      pulumi.StringInput `pulumi:"machineName"`
	Source           pulumi.StringInput `pulumi:"source"`
	SourceComputerId pulumi.StringInput `pulumi:"sourceComputerId"`
	Vmuuid           pulumi.StringInput `pulumi:"vmuuid"`
	WorkspaceId      pulumi.StringInput `pulumi:"workspaceId"`
}

func (OnPremiseResourceDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseResourceDetailsResponse)(nil)).Elem()
}

func (i OnPremiseResourceDetailsResponseArgs) ToOnPremiseResourceDetailsResponseOutput() OnPremiseResourceDetailsResponseOutput {
	return i.ToOnPremiseResourceDetailsResponseOutputWithContext(context.Background())
}

func (i OnPremiseResourceDetailsResponseArgs) ToOnPremiseResourceDetailsResponseOutputWithContext(ctx context.Context) OnPremiseResourceDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremiseResourceDetailsResponseOutput)
}

type OnPremiseResourceDetailsResponseOutput struct{ *pulumi.OutputState }

func (OnPremiseResourceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseResourceDetailsResponse)(nil)).Elem()
}

func (o OnPremiseResourceDetailsResponseOutput) ToOnPremiseResourceDetailsResponseOutput() OnPremiseResourceDetailsResponseOutput {
	return o
}

func (o OnPremiseResourceDetailsResponseOutput) ToOnPremiseResourceDetailsResponseOutputWithContext(ctx context.Context) OnPremiseResourceDetailsResponseOutput {
	return o
}

func (o OnPremiseResourceDetailsResponseOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.Source }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsResponseOutput) SourceComputerId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.SourceComputerId }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsResponseOutput) Vmuuid() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.Vmuuid }).(pulumi.StringOutput)
}

func (o OnPremiseResourceDetailsResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseResourceDetailsResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
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





type OnPremiseSqlResourceDetailsInput interface {
	pulumi.Input

	ToOnPremiseSqlResourceDetailsOutput() OnPremiseSqlResourceDetailsOutput
	ToOnPremiseSqlResourceDetailsOutputWithContext(context.Context) OnPremiseSqlResourceDetailsOutput
}

type OnPremiseSqlResourceDetailsArgs struct {
	DatabaseName     pulumi.StringInput `pulumi:"databaseName"`
	MachineName      pulumi.StringInput `pulumi:"machineName"`
	ServerName       pulumi.StringInput `pulumi:"serverName"`
	Source           pulumi.StringInput `pulumi:"source"`
	SourceComputerId pulumi.StringInput `pulumi:"sourceComputerId"`
	Vmuuid           pulumi.StringInput `pulumi:"vmuuid"`
	WorkspaceId      pulumi.StringInput `pulumi:"workspaceId"`
}

func (OnPremiseSqlResourceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSqlResourceDetails)(nil)).Elem()
}

func (i OnPremiseSqlResourceDetailsArgs) ToOnPremiseSqlResourceDetailsOutput() OnPremiseSqlResourceDetailsOutput {
	return i.ToOnPremiseSqlResourceDetailsOutputWithContext(context.Background())
}

func (i OnPremiseSqlResourceDetailsArgs) ToOnPremiseSqlResourceDetailsOutputWithContext(ctx context.Context) OnPremiseSqlResourceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremiseSqlResourceDetailsOutput)
}

type OnPremiseSqlResourceDetailsOutput struct{ *pulumi.OutputState }

func (OnPremiseSqlResourceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSqlResourceDetails)(nil)).Elem()
}

func (o OnPremiseSqlResourceDetailsOutput) ToOnPremiseSqlResourceDetailsOutput() OnPremiseSqlResourceDetailsOutput {
	return o
}

func (o OnPremiseSqlResourceDetailsOutput) ToOnPremiseSqlResourceDetailsOutputWithContext(ctx context.Context) OnPremiseSqlResourceDetailsOutput {
	return o
}

func (o OnPremiseSqlResourceDetailsOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.ServerName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.Source }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) SourceComputerId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.SourceComputerId }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) Vmuuid() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.Vmuuid }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetails) string { return v.WorkspaceId }).(pulumi.StringOutput)
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





type OnPremiseSqlResourceDetailsResponseInput interface {
	pulumi.Input

	ToOnPremiseSqlResourceDetailsResponseOutput() OnPremiseSqlResourceDetailsResponseOutput
	ToOnPremiseSqlResourceDetailsResponseOutputWithContext(context.Context) OnPremiseSqlResourceDetailsResponseOutput
}

type OnPremiseSqlResourceDetailsResponseArgs struct {
	DatabaseName     pulumi.StringInput `pulumi:"databaseName"`
	MachineName      pulumi.StringInput `pulumi:"machineName"`
	ServerName       pulumi.StringInput `pulumi:"serverName"`
	Source           pulumi.StringInput `pulumi:"source"`
	SourceComputerId pulumi.StringInput `pulumi:"sourceComputerId"`
	Vmuuid           pulumi.StringInput `pulumi:"vmuuid"`
	WorkspaceId      pulumi.StringInput `pulumi:"workspaceId"`
}

func (OnPremiseSqlResourceDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSqlResourceDetailsResponse)(nil)).Elem()
}

func (i OnPremiseSqlResourceDetailsResponseArgs) ToOnPremiseSqlResourceDetailsResponseOutput() OnPremiseSqlResourceDetailsResponseOutput {
	return i.ToOnPremiseSqlResourceDetailsResponseOutputWithContext(context.Background())
}

func (i OnPremiseSqlResourceDetailsResponseArgs) ToOnPremiseSqlResourceDetailsResponseOutputWithContext(ctx context.Context) OnPremiseSqlResourceDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnPremiseSqlResourceDetailsResponseOutput)
}

type OnPremiseSqlResourceDetailsResponseOutput struct{ *pulumi.OutputState }

func (OnPremiseSqlResourceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnPremiseSqlResourceDetailsResponse)(nil)).Elem()
}

func (o OnPremiseSqlResourceDetailsResponseOutput) ToOnPremiseSqlResourceDetailsResponseOutput() OnPremiseSqlResourceDetailsResponseOutput {
	return o
}

func (o OnPremiseSqlResourceDetailsResponseOutput) ToOnPremiseSqlResourceDetailsResponseOutputWithContext(ctx context.Context) OnPremiseSqlResourceDetailsResponseOutput {
	return o
}

func (o OnPremiseSqlResourceDetailsResponseOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.ServerName }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.Source }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) SourceComputerId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.SourceComputerId }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) Vmuuid() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.Vmuuid }).(pulumi.StringOutput)
}

func (o OnPremiseSqlResourceDetailsResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v OnPremiseSqlResourceDetailsResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

type PathRecommendation struct {
	Action              *string              `pulumi:"action"`
	Common              *bool                `pulumi:"common"`
	ConfigurationStatus *string              `pulumi:"configurationStatus"`
	FileType            *string              `pulumi:"fileType"`
	Path                *string              `pulumi:"path"`
	PublisherInfo       *PublisherInfo       `pulumi:"publisherInfo"`
	Type                *string              `pulumi:"type"`
	UserSids            []string             `pulumi:"userSids"`
	Usernames           []UserRecommendation `pulumi:"usernames"`
}





type PathRecommendationInput interface {
	pulumi.Input

	ToPathRecommendationOutput() PathRecommendationOutput
	ToPathRecommendationOutputWithContext(context.Context) PathRecommendationOutput
}

type PathRecommendationArgs struct {
	Action              pulumi.StringPtrInput        `pulumi:"action"`
	Common              pulumi.BoolPtrInput          `pulumi:"common"`
	ConfigurationStatus pulumi.StringPtrInput        `pulumi:"configurationStatus"`
	FileType            pulumi.StringPtrInput        `pulumi:"fileType"`
	Path                pulumi.StringPtrInput        `pulumi:"path"`
	PublisherInfo       PublisherInfoPtrInput        `pulumi:"publisherInfo"`
	Type                pulumi.StringPtrInput        `pulumi:"type"`
	UserSids            pulumi.StringArrayInput      `pulumi:"userSids"`
	Usernames           UserRecommendationArrayInput `pulumi:"usernames"`
}

func (PathRecommendationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PathRecommendation)(nil)).Elem()
}

func (i PathRecommendationArgs) ToPathRecommendationOutput() PathRecommendationOutput {
	return i.ToPathRecommendationOutputWithContext(context.Background())
}

func (i PathRecommendationArgs) ToPathRecommendationOutputWithContext(ctx context.Context) PathRecommendationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PathRecommendationOutput)
}





type PathRecommendationArrayInput interface {
	pulumi.Input

	ToPathRecommendationArrayOutput() PathRecommendationArrayOutput
	ToPathRecommendationArrayOutputWithContext(context.Context) PathRecommendationArrayOutput
}

type PathRecommendationArray []PathRecommendationInput

func (PathRecommendationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PathRecommendation)(nil)).Elem()
}

func (i PathRecommendationArray) ToPathRecommendationArrayOutput() PathRecommendationArrayOutput {
	return i.ToPathRecommendationArrayOutputWithContext(context.Background())
}

func (i PathRecommendationArray) ToPathRecommendationArrayOutputWithContext(ctx context.Context) PathRecommendationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PathRecommendationArrayOutput)
}

type PathRecommendationOutput struct{ *pulumi.OutputState }

func (PathRecommendationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PathRecommendation)(nil)).Elem()
}

func (o PathRecommendationOutput) ToPathRecommendationOutput() PathRecommendationOutput {
	return o
}

func (o PathRecommendationOutput) ToPathRecommendationOutputWithContext(ctx context.Context) PathRecommendationOutput {
	return o
}

func (o PathRecommendationOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) Common() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *bool { return v.Common }).(pulumi.BoolPtrOutput)
}

func (o PathRecommendationOutput) ConfigurationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.ConfigurationStatus }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) FileType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.FileType }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) PublisherInfo() PublisherInfoPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *PublisherInfo { return v.PublisherInfo }).(PublisherInfoPtrOutput)
}

func (o PathRecommendationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) UserSids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PathRecommendation) []string { return v.UserSids }).(pulumi.StringArrayOutput)
}

func (o PathRecommendationOutput) Usernames() UserRecommendationArrayOutput {
	return o.ApplyT(func(v PathRecommendation) []UserRecommendation { return v.Usernames }).(UserRecommendationArrayOutput)
}

type PathRecommendationArrayOutput struct{ *pulumi.OutputState }

func (PathRecommendationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PathRecommendation)(nil)).Elem()
}

func (o PathRecommendationArrayOutput) ToPathRecommendationArrayOutput() PathRecommendationArrayOutput {
	return o
}

func (o PathRecommendationArrayOutput) ToPathRecommendationArrayOutputWithContext(ctx context.Context) PathRecommendationArrayOutput {
	return o
}

func (o PathRecommendationArrayOutput) Index(i pulumi.IntInput) PathRecommendationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PathRecommendation {
		return vs[0].([]PathRecommendation)[vs[1].(int)]
	}).(PathRecommendationOutput)
}

type PathRecommendationResponse struct {
	Action              *string                      `pulumi:"action"`
	Common              *bool                        `pulumi:"common"`
	ConfigurationStatus *string                      `pulumi:"configurationStatus"`
	FileType            *string                      `pulumi:"fileType"`
	Path                *string                      `pulumi:"path"`
	PublisherInfo       *PublisherInfoResponse       `pulumi:"publisherInfo"`
	Type                *string                      `pulumi:"type"`
	UserSids            []string                     `pulumi:"userSids"`
	Usernames           []UserRecommendationResponse `pulumi:"usernames"`
}





type PathRecommendationResponseInput interface {
	pulumi.Input

	ToPathRecommendationResponseOutput() PathRecommendationResponseOutput
	ToPathRecommendationResponseOutputWithContext(context.Context) PathRecommendationResponseOutput
}

type PathRecommendationResponseArgs struct {
	Action              pulumi.StringPtrInput                `pulumi:"action"`
	Common              pulumi.BoolPtrInput                  `pulumi:"common"`
	ConfigurationStatus pulumi.StringPtrInput                `pulumi:"configurationStatus"`
	FileType            pulumi.StringPtrInput                `pulumi:"fileType"`
	Path                pulumi.StringPtrInput                `pulumi:"path"`
	PublisherInfo       PublisherInfoResponsePtrInput        `pulumi:"publisherInfo"`
	Type                pulumi.StringPtrInput                `pulumi:"type"`
	UserSids            pulumi.StringArrayInput              `pulumi:"userSids"`
	Usernames           UserRecommendationResponseArrayInput `pulumi:"usernames"`
}

func (PathRecommendationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PathRecommendationResponse)(nil)).Elem()
}

func (i PathRecommendationResponseArgs) ToPathRecommendationResponseOutput() PathRecommendationResponseOutput {
	return i.ToPathRecommendationResponseOutputWithContext(context.Background())
}

func (i PathRecommendationResponseArgs) ToPathRecommendationResponseOutputWithContext(ctx context.Context) PathRecommendationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PathRecommendationResponseOutput)
}





type PathRecommendationResponseArrayInput interface {
	pulumi.Input

	ToPathRecommendationResponseArrayOutput() PathRecommendationResponseArrayOutput
	ToPathRecommendationResponseArrayOutputWithContext(context.Context) PathRecommendationResponseArrayOutput
}

type PathRecommendationResponseArray []PathRecommendationResponseInput

func (PathRecommendationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PathRecommendationResponse)(nil)).Elem()
}

func (i PathRecommendationResponseArray) ToPathRecommendationResponseArrayOutput() PathRecommendationResponseArrayOutput {
	return i.ToPathRecommendationResponseArrayOutputWithContext(context.Background())
}

func (i PathRecommendationResponseArray) ToPathRecommendationResponseArrayOutputWithContext(ctx context.Context) PathRecommendationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PathRecommendationResponseArrayOutput)
}

type PathRecommendationResponseOutput struct{ *pulumi.OutputState }

func (PathRecommendationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PathRecommendationResponse)(nil)).Elem()
}

func (o PathRecommendationResponseOutput) ToPathRecommendationResponseOutput() PathRecommendationResponseOutput {
	return o
}

func (o PathRecommendationResponseOutput) ToPathRecommendationResponseOutputWithContext(ctx context.Context) PathRecommendationResponseOutput {
	return o
}

func (o PathRecommendationResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) Common() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *bool { return v.Common }).(pulumi.BoolPtrOutput)
}

func (o PathRecommendationResponseOutput) ConfigurationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.ConfigurationStatus }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) FileType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.FileType }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) PublisherInfo() PublisherInfoResponsePtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *PublisherInfoResponse { return v.PublisherInfo }).(PublisherInfoResponsePtrOutput)
}

func (o PathRecommendationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) UserSids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PathRecommendationResponse) []string { return v.UserSids }).(pulumi.StringArrayOutput)
}

func (o PathRecommendationResponseOutput) Usernames() UserRecommendationResponseArrayOutput {
	return o.ApplyT(func(v PathRecommendationResponse) []UserRecommendationResponse { return v.Usernames }).(UserRecommendationResponseArrayOutput)
}

type PathRecommendationResponseArrayOutput struct{ *pulumi.OutputState }

func (PathRecommendationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PathRecommendationResponse)(nil)).Elem()
}

func (o PathRecommendationResponseArrayOutput) ToPathRecommendationResponseArrayOutput() PathRecommendationResponseArrayOutput {
	return o
}

func (o PathRecommendationResponseArrayOutput) ToPathRecommendationResponseArrayOutputWithContext(ctx context.Context) PathRecommendationResponseArrayOutput {
	return o
}

func (o PathRecommendationResponseArrayOutput) Index(i pulumi.IntInput) PathRecommendationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PathRecommendationResponse {
		return vs[0].([]PathRecommendationResponse)[vs[1].(int)]
	}).(PathRecommendationResponseOutput)
}

type ProtectionMode struct {
	Exe        *string `pulumi:"exe"`
	Executable *string `pulumi:"executable"`
	Msi        *string `pulumi:"msi"`
	Script     *string `pulumi:"script"`
}





type ProtectionModeInput interface {
	pulumi.Input

	ToProtectionModeOutput() ProtectionModeOutput
	ToProtectionModeOutputWithContext(context.Context) ProtectionModeOutput
}

type ProtectionModeArgs struct {
	Exe        pulumi.StringPtrInput `pulumi:"exe"`
	Executable pulumi.StringPtrInput `pulumi:"executable"`
	Msi        pulumi.StringPtrInput `pulumi:"msi"`
	Script     pulumi.StringPtrInput `pulumi:"script"`
}

func (ProtectionModeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionMode)(nil)).Elem()
}

func (i ProtectionModeArgs) ToProtectionModeOutput() ProtectionModeOutput {
	return i.ToProtectionModeOutputWithContext(context.Background())
}

func (i ProtectionModeArgs) ToProtectionModeOutputWithContext(ctx context.Context) ProtectionModeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeOutput)
}

func (i ProtectionModeArgs) ToProtectionModePtrOutput() ProtectionModePtrOutput {
	return i.ToProtectionModePtrOutputWithContext(context.Background())
}

func (i ProtectionModeArgs) ToProtectionModePtrOutputWithContext(ctx context.Context) ProtectionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeOutput).ToProtectionModePtrOutputWithContext(ctx)
}









type ProtectionModePtrInput interface {
	pulumi.Input

	ToProtectionModePtrOutput() ProtectionModePtrOutput
	ToProtectionModePtrOutputWithContext(context.Context) ProtectionModePtrOutput
}

type protectionModePtrType ProtectionModeArgs

func ProtectionModePtr(v *ProtectionModeArgs) ProtectionModePtrInput {
	return (*protectionModePtrType)(v)
}

func (*protectionModePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionMode)(nil)).Elem()
}

func (i *protectionModePtrType) ToProtectionModePtrOutput() ProtectionModePtrOutput {
	return i.ToProtectionModePtrOutputWithContext(context.Background())
}

func (i *protectionModePtrType) ToProtectionModePtrOutputWithContext(ctx context.Context) ProtectionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModePtrOutput)
}

type ProtectionModeOutput struct{ *pulumi.OutputState }

func (ProtectionModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionMode)(nil)).Elem()
}

func (o ProtectionModeOutput) ToProtectionModeOutput() ProtectionModeOutput {
	return o
}

func (o ProtectionModeOutput) ToProtectionModeOutputWithContext(ctx context.Context) ProtectionModeOutput {
	return o
}

func (o ProtectionModeOutput) ToProtectionModePtrOutput() ProtectionModePtrOutput {
	return o.ToProtectionModePtrOutputWithContext(context.Background())
}

func (o ProtectionModeOutput) ToProtectionModePtrOutputWithContext(ctx context.Context) ProtectionModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectionMode) *ProtectionMode {
		return &v
	}).(ProtectionModePtrOutput)
}

func (o ProtectionModeOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionMode) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeOutput) Executable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionMode) *string { return v.Executable }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeOutput) Msi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionMode) *string { return v.Msi }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionMode) *string { return v.Script }).(pulumi.StringPtrOutput)
}

type ProtectionModePtrOutput struct{ *pulumi.OutputState }

func (ProtectionModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionMode)(nil)).Elem()
}

func (o ProtectionModePtrOutput) ToProtectionModePtrOutput() ProtectionModePtrOutput {
	return o
}

func (o ProtectionModePtrOutput) ToProtectionModePtrOutputWithContext(ctx context.Context) ProtectionModePtrOutput {
	return o
}

func (o ProtectionModePtrOutput) Elem() ProtectionModeOutput {
	return o.ApplyT(func(v *ProtectionMode) ProtectionMode {
		if v != nil {
			return *v
		}
		var ret ProtectionMode
		return ret
	}).(ProtectionModeOutput)
}

func (o ProtectionModePtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModePtrOutput) Executable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) *string {
		if v == nil {
			return nil
		}
		return v.Executable
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModePtrOutput) Msi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) *string {
		if v == nil {
			return nil
		}
		return v.Msi
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModePtrOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) *string {
		if v == nil {
			return nil
		}
		return v.Script
	}).(pulumi.StringPtrOutput)
}

type ProtectionModeResponse struct {
	Exe        *string `pulumi:"exe"`
	Executable *string `pulumi:"executable"`
	Msi        *string `pulumi:"msi"`
	Script     *string `pulumi:"script"`
}





type ProtectionModeResponseInput interface {
	pulumi.Input

	ToProtectionModeResponseOutput() ProtectionModeResponseOutput
	ToProtectionModeResponseOutputWithContext(context.Context) ProtectionModeResponseOutput
}

type ProtectionModeResponseArgs struct {
	Exe        pulumi.StringPtrInput `pulumi:"exe"`
	Executable pulumi.StringPtrInput `pulumi:"executable"`
	Msi        pulumi.StringPtrInput `pulumi:"msi"`
	Script     pulumi.StringPtrInput `pulumi:"script"`
}

func (ProtectionModeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionModeResponse)(nil)).Elem()
}

func (i ProtectionModeResponseArgs) ToProtectionModeResponseOutput() ProtectionModeResponseOutput {
	return i.ToProtectionModeResponseOutputWithContext(context.Background())
}

func (i ProtectionModeResponseArgs) ToProtectionModeResponseOutputWithContext(ctx context.Context) ProtectionModeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeResponseOutput)
}

func (i ProtectionModeResponseArgs) ToProtectionModeResponsePtrOutput() ProtectionModeResponsePtrOutput {
	return i.ToProtectionModeResponsePtrOutputWithContext(context.Background())
}

func (i ProtectionModeResponseArgs) ToProtectionModeResponsePtrOutputWithContext(ctx context.Context) ProtectionModeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeResponseOutput).ToProtectionModeResponsePtrOutputWithContext(ctx)
}









type ProtectionModeResponsePtrInput interface {
	pulumi.Input

	ToProtectionModeResponsePtrOutput() ProtectionModeResponsePtrOutput
	ToProtectionModeResponsePtrOutputWithContext(context.Context) ProtectionModeResponsePtrOutput
}

type protectionModeResponsePtrType ProtectionModeResponseArgs

func ProtectionModeResponsePtr(v *ProtectionModeResponseArgs) ProtectionModeResponsePtrInput {
	return (*protectionModeResponsePtrType)(v)
}

func (*protectionModeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionModeResponse)(nil)).Elem()
}

func (i *protectionModeResponsePtrType) ToProtectionModeResponsePtrOutput() ProtectionModeResponsePtrOutput {
	return i.ToProtectionModeResponsePtrOutputWithContext(context.Background())
}

func (i *protectionModeResponsePtrType) ToProtectionModeResponsePtrOutputWithContext(ctx context.Context) ProtectionModeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeResponsePtrOutput)
}

type ProtectionModeResponseOutput struct{ *pulumi.OutputState }

func (ProtectionModeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionModeResponse)(nil)).Elem()
}

func (o ProtectionModeResponseOutput) ToProtectionModeResponseOutput() ProtectionModeResponseOutput {
	return o
}

func (o ProtectionModeResponseOutput) ToProtectionModeResponseOutputWithContext(ctx context.Context) ProtectionModeResponseOutput {
	return o
}

func (o ProtectionModeResponseOutput) ToProtectionModeResponsePtrOutput() ProtectionModeResponsePtrOutput {
	return o.ToProtectionModeResponsePtrOutputWithContext(context.Background())
}

func (o ProtectionModeResponseOutput) ToProtectionModeResponsePtrOutputWithContext(ctx context.Context) ProtectionModeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectionModeResponse) *ProtectionModeResponse {
		return &v
	}).(ProtectionModeResponsePtrOutput)
}

func (o ProtectionModeResponseOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionModeResponse) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponseOutput) Executable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionModeResponse) *string { return v.Executable }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponseOutput) Msi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionModeResponse) *string { return v.Msi }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponseOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionModeResponse) *string { return v.Script }).(pulumi.StringPtrOutput)
}

type ProtectionModeResponsePtrOutput struct{ *pulumi.OutputState }

func (ProtectionModeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionModeResponse)(nil)).Elem()
}

func (o ProtectionModeResponsePtrOutput) ToProtectionModeResponsePtrOutput() ProtectionModeResponsePtrOutput {
	return o
}

func (o ProtectionModeResponsePtrOutput) ToProtectionModeResponsePtrOutputWithContext(ctx context.Context) ProtectionModeResponsePtrOutput {
	return o
}

func (o ProtectionModeResponsePtrOutput) Elem() ProtectionModeResponseOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) ProtectionModeResponse {
		if v != nil {
			return *v
		}
		var ret ProtectionModeResponse
		return ret
	}).(ProtectionModeResponseOutput)
}

func (o ProtectionModeResponsePtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponsePtrOutput) Executable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Executable
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponsePtrOutput) Msi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Msi
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponsePtrOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Script
	}).(pulumi.StringPtrOutput)
}

type ProxyServerProperties struct {
	Ip   *string `pulumi:"ip"`
	Port *string `pulumi:"port"`
}





type ProxyServerPropertiesInput interface {
	pulumi.Input

	ToProxyServerPropertiesOutput() ProxyServerPropertiesOutput
	ToProxyServerPropertiesOutputWithContext(context.Context) ProxyServerPropertiesOutput
}

type ProxyServerPropertiesArgs struct {
	Ip   pulumi.StringPtrInput `pulumi:"ip"`
	Port pulumi.StringPtrInput `pulumi:"port"`
}

func (ProxyServerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyServerProperties)(nil)).Elem()
}

func (i ProxyServerPropertiesArgs) ToProxyServerPropertiesOutput() ProxyServerPropertiesOutput {
	return i.ToProxyServerPropertiesOutputWithContext(context.Background())
}

func (i ProxyServerPropertiesArgs) ToProxyServerPropertiesOutputWithContext(ctx context.Context) ProxyServerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyServerPropertiesOutput)
}

func (i ProxyServerPropertiesArgs) ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput {
	return i.ToProxyServerPropertiesPtrOutputWithContext(context.Background())
}

func (i ProxyServerPropertiesArgs) ToProxyServerPropertiesPtrOutputWithContext(ctx context.Context) ProxyServerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyServerPropertiesOutput).ToProxyServerPropertiesPtrOutputWithContext(ctx)
}









type ProxyServerPropertiesPtrInput interface {
	pulumi.Input

	ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput
	ToProxyServerPropertiesPtrOutputWithContext(context.Context) ProxyServerPropertiesPtrOutput
}

type proxyServerPropertiesPtrType ProxyServerPropertiesArgs

func ProxyServerPropertiesPtr(v *ProxyServerPropertiesArgs) ProxyServerPropertiesPtrInput {
	return (*proxyServerPropertiesPtrType)(v)
}

func (*proxyServerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyServerProperties)(nil)).Elem()
}

func (i *proxyServerPropertiesPtrType) ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput {
	return i.ToProxyServerPropertiesPtrOutputWithContext(context.Background())
}

func (i *proxyServerPropertiesPtrType) ToProxyServerPropertiesPtrOutputWithContext(ctx context.Context) ProxyServerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyServerPropertiesPtrOutput)
}

type ProxyServerPropertiesOutput struct{ *pulumi.OutputState }

func (ProxyServerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyServerProperties)(nil)).Elem()
}

func (o ProxyServerPropertiesOutput) ToProxyServerPropertiesOutput() ProxyServerPropertiesOutput {
	return o
}

func (o ProxyServerPropertiesOutput) ToProxyServerPropertiesOutputWithContext(ctx context.Context) ProxyServerPropertiesOutput {
	return o
}

func (o ProxyServerPropertiesOutput) ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput {
	return o.ToProxyServerPropertiesPtrOutputWithContext(context.Background())
}

func (o ProxyServerPropertiesOutput) ToProxyServerPropertiesPtrOutputWithContext(ctx context.Context) ProxyServerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProxyServerProperties) *ProxyServerProperties {
		return &v
	}).(ProxyServerPropertiesPtrOutput)
}

func (o ProxyServerPropertiesOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyServerProperties) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o ProxyServerPropertiesOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyServerProperties) *string { return v.Port }).(pulumi.StringPtrOutput)
}

type ProxyServerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ProxyServerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyServerProperties)(nil)).Elem()
}

func (o ProxyServerPropertiesPtrOutput) ToProxyServerPropertiesPtrOutput() ProxyServerPropertiesPtrOutput {
	return o
}

func (o ProxyServerPropertiesPtrOutput) ToProxyServerPropertiesPtrOutputWithContext(ctx context.Context) ProxyServerPropertiesPtrOutput {
	return o
}

func (o ProxyServerPropertiesPtrOutput) Elem() ProxyServerPropertiesOutput {
	return o.ApplyT(func(v *ProxyServerProperties) ProxyServerProperties {
		if v != nil {
			return *v
		}
		var ret ProxyServerProperties
		return ret
	}).(ProxyServerPropertiesOutput)
}

func (o ProxyServerPropertiesPtrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyServerProperties) *string {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(pulumi.StringPtrOutput)
}

func (o ProxyServerPropertiesPtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyServerProperties) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

type ProxyServerPropertiesResponse struct {
	Ip   *string `pulumi:"ip"`
	Port *string `pulumi:"port"`
}





type ProxyServerPropertiesResponseInput interface {
	pulumi.Input

	ToProxyServerPropertiesResponseOutput() ProxyServerPropertiesResponseOutput
	ToProxyServerPropertiesResponseOutputWithContext(context.Context) ProxyServerPropertiesResponseOutput
}

type ProxyServerPropertiesResponseArgs struct {
	Ip   pulumi.StringPtrInput `pulumi:"ip"`
	Port pulumi.StringPtrInput `pulumi:"port"`
}

func (ProxyServerPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyServerPropertiesResponse)(nil)).Elem()
}

func (i ProxyServerPropertiesResponseArgs) ToProxyServerPropertiesResponseOutput() ProxyServerPropertiesResponseOutput {
	return i.ToProxyServerPropertiesResponseOutputWithContext(context.Background())
}

func (i ProxyServerPropertiesResponseArgs) ToProxyServerPropertiesResponseOutputWithContext(ctx context.Context) ProxyServerPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyServerPropertiesResponseOutput)
}

func (i ProxyServerPropertiesResponseArgs) ToProxyServerPropertiesResponsePtrOutput() ProxyServerPropertiesResponsePtrOutput {
	return i.ToProxyServerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ProxyServerPropertiesResponseArgs) ToProxyServerPropertiesResponsePtrOutputWithContext(ctx context.Context) ProxyServerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyServerPropertiesResponseOutput).ToProxyServerPropertiesResponsePtrOutputWithContext(ctx)
}









type ProxyServerPropertiesResponsePtrInput interface {
	pulumi.Input

	ToProxyServerPropertiesResponsePtrOutput() ProxyServerPropertiesResponsePtrOutput
	ToProxyServerPropertiesResponsePtrOutputWithContext(context.Context) ProxyServerPropertiesResponsePtrOutput
}

type proxyServerPropertiesResponsePtrType ProxyServerPropertiesResponseArgs

func ProxyServerPropertiesResponsePtr(v *ProxyServerPropertiesResponseArgs) ProxyServerPropertiesResponsePtrInput {
	return (*proxyServerPropertiesResponsePtrType)(v)
}

func (*proxyServerPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyServerPropertiesResponse)(nil)).Elem()
}

func (i *proxyServerPropertiesResponsePtrType) ToProxyServerPropertiesResponsePtrOutput() ProxyServerPropertiesResponsePtrOutput {
	return i.ToProxyServerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *proxyServerPropertiesResponsePtrType) ToProxyServerPropertiesResponsePtrOutputWithContext(ctx context.Context) ProxyServerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyServerPropertiesResponsePtrOutput)
}

type ProxyServerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ProxyServerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProxyServerPropertiesResponse)(nil)).Elem()
}

func (o ProxyServerPropertiesResponseOutput) ToProxyServerPropertiesResponseOutput() ProxyServerPropertiesResponseOutput {
	return o
}

func (o ProxyServerPropertiesResponseOutput) ToProxyServerPropertiesResponseOutputWithContext(ctx context.Context) ProxyServerPropertiesResponseOutput {
	return o
}

func (o ProxyServerPropertiesResponseOutput) ToProxyServerPropertiesResponsePtrOutput() ProxyServerPropertiesResponsePtrOutput {
	return o.ToProxyServerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ProxyServerPropertiesResponseOutput) ToProxyServerPropertiesResponsePtrOutputWithContext(ctx context.Context) ProxyServerPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProxyServerPropertiesResponse) *ProxyServerPropertiesResponse {
		return &v
	}).(ProxyServerPropertiesResponsePtrOutput)
}

func (o ProxyServerPropertiesResponseOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyServerPropertiesResponse) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o ProxyServerPropertiesResponseOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProxyServerPropertiesResponse) *string { return v.Port }).(pulumi.StringPtrOutput)
}

type ProxyServerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ProxyServerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProxyServerPropertiesResponse)(nil)).Elem()
}

func (o ProxyServerPropertiesResponsePtrOutput) ToProxyServerPropertiesResponsePtrOutput() ProxyServerPropertiesResponsePtrOutput {
	return o
}

func (o ProxyServerPropertiesResponsePtrOutput) ToProxyServerPropertiesResponsePtrOutputWithContext(ctx context.Context) ProxyServerPropertiesResponsePtrOutput {
	return o
}

func (o ProxyServerPropertiesResponsePtrOutput) Elem() ProxyServerPropertiesResponseOutput {
	return o.ApplyT(func(v *ProxyServerPropertiesResponse) ProxyServerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ProxyServerPropertiesResponse
		return ret
	}).(ProxyServerPropertiesResponseOutput)
}

func (o ProxyServerPropertiesResponsePtrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(pulumi.StringPtrOutput)
}

func (o ProxyServerPropertiesResponsePtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProxyServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

type PublisherInfo struct {
	BinaryName    *string `pulumi:"binaryName"`
	ProductName   *string `pulumi:"productName"`
	PublisherName *string `pulumi:"publisherName"`
	Version       *string `pulumi:"version"`
}





type PublisherInfoInput interface {
	pulumi.Input

	ToPublisherInfoOutput() PublisherInfoOutput
	ToPublisherInfoOutputWithContext(context.Context) PublisherInfoOutput
}

type PublisherInfoArgs struct {
	BinaryName    pulumi.StringPtrInput `pulumi:"binaryName"`
	ProductName   pulumi.StringPtrInput `pulumi:"productName"`
	PublisherName pulumi.StringPtrInput `pulumi:"publisherName"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (PublisherInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublisherInfo)(nil)).Elem()
}

func (i PublisherInfoArgs) ToPublisherInfoOutput() PublisherInfoOutput {
	return i.ToPublisherInfoOutputWithContext(context.Background())
}

func (i PublisherInfoArgs) ToPublisherInfoOutputWithContext(ctx context.Context) PublisherInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherInfoOutput)
}

func (i PublisherInfoArgs) ToPublisherInfoPtrOutput() PublisherInfoPtrOutput {
	return i.ToPublisherInfoPtrOutputWithContext(context.Background())
}

func (i PublisherInfoArgs) ToPublisherInfoPtrOutputWithContext(ctx context.Context) PublisherInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherInfoOutput).ToPublisherInfoPtrOutputWithContext(ctx)
}









type PublisherInfoPtrInput interface {
	pulumi.Input

	ToPublisherInfoPtrOutput() PublisherInfoPtrOutput
	ToPublisherInfoPtrOutputWithContext(context.Context) PublisherInfoPtrOutput
}

type publisherInfoPtrType PublisherInfoArgs

func PublisherInfoPtr(v *PublisherInfoArgs) PublisherInfoPtrInput {
	return (*publisherInfoPtrType)(v)
}

func (*publisherInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublisherInfo)(nil)).Elem()
}

func (i *publisherInfoPtrType) ToPublisherInfoPtrOutput() PublisherInfoPtrOutput {
	return i.ToPublisherInfoPtrOutputWithContext(context.Background())
}

func (i *publisherInfoPtrType) ToPublisherInfoPtrOutputWithContext(ctx context.Context) PublisherInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherInfoPtrOutput)
}

type PublisherInfoOutput struct{ *pulumi.OutputState }

func (PublisherInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublisherInfo)(nil)).Elem()
}

func (o PublisherInfoOutput) ToPublisherInfoOutput() PublisherInfoOutput {
	return o
}

func (o PublisherInfoOutput) ToPublisherInfoOutputWithContext(ctx context.Context) PublisherInfoOutput {
	return o
}

func (o PublisherInfoOutput) ToPublisherInfoPtrOutput() PublisherInfoPtrOutput {
	return o.ToPublisherInfoPtrOutputWithContext(context.Background())
}

func (o PublisherInfoOutput) ToPublisherInfoPtrOutputWithContext(ctx context.Context) PublisherInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublisherInfo) *PublisherInfo {
		return &v
	}).(PublisherInfoPtrOutput)
}

func (o PublisherInfoOutput) BinaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfo) *string { return v.BinaryName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfo) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoOutput) PublisherName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfo) *string { return v.PublisherName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfo) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type PublisherInfoPtrOutput struct{ *pulumi.OutputState }

func (PublisherInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublisherInfo)(nil)).Elem()
}

func (o PublisherInfoPtrOutput) ToPublisherInfoPtrOutput() PublisherInfoPtrOutput {
	return o
}

func (o PublisherInfoPtrOutput) ToPublisherInfoPtrOutputWithContext(ctx context.Context) PublisherInfoPtrOutput {
	return o
}

func (o PublisherInfoPtrOutput) Elem() PublisherInfoOutput {
	return o.ApplyT(func(v *PublisherInfo) PublisherInfo {
		if v != nil {
			return *v
		}
		var ret PublisherInfo
		return ret
	}).(PublisherInfoOutput)
}

func (o PublisherInfoPtrOutput) BinaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfo) *string {
		if v == nil {
			return nil
		}
		return v.BinaryName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoPtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfo) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoPtrOutput) PublisherName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfo) *string {
		if v == nil {
			return nil
		}
		return v.PublisherName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfo) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type PublisherInfoResponse struct {
	BinaryName    *string `pulumi:"binaryName"`
	ProductName   *string `pulumi:"productName"`
	PublisherName *string `pulumi:"publisherName"`
	Version       *string `pulumi:"version"`
}





type PublisherInfoResponseInput interface {
	pulumi.Input

	ToPublisherInfoResponseOutput() PublisherInfoResponseOutput
	ToPublisherInfoResponseOutputWithContext(context.Context) PublisherInfoResponseOutput
}

type PublisherInfoResponseArgs struct {
	BinaryName    pulumi.StringPtrInput `pulumi:"binaryName"`
	ProductName   pulumi.StringPtrInput `pulumi:"productName"`
	PublisherName pulumi.StringPtrInput `pulumi:"publisherName"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (PublisherInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublisherInfoResponse)(nil)).Elem()
}

func (i PublisherInfoResponseArgs) ToPublisherInfoResponseOutput() PublisherInfoResponseOutput {
	return i.ToPublisherInfoResponseOutputWithContext(context.Background())
}

func (i PublisherInfoResponseArgs) ToPublisherInfoResponseOutputWithContext(ctx context.Context) PublisherInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherInfoResponseOutput)
}

func (i PublisherInfoResponseArgs) ToPublisherInfoResponsePtrOutput() PublisherInfoResponsePtrOutput {
	return i.ToPublisherInfoResponsePtrOutputWithContext(context.Background())
}

func (i PublisherInfoResponseArgs) ToPublisherInfoResponsePtrOutputWithContext(ctx context.Context) PublisherInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherInfoResponseOutput).ToPublisherInfoResponsePtrOutputWithContext(ctx)
}









type PublisherInfoResponsePtrInput interface {
	pulumi.Input

	ToPublisherInfoResponsePtrOutput() PublisherInfoResponsePtrOutput
	ToPublisherInfoResponsePtrOutputWithContext(context.Context) PublisherInfoResponsePtrOutput
}

type publisherInfoResponsePtrType PublisherInfoResponseArgs

func PublisherInfoResponsePtr(v *PublisherInfoResponseArgs) PublisherInfoResponsePtrInput {
	return (*publisherInfoResponsePtrType)(v)
}

func (*publisherInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublisherInfoResponse)(nil)).Elem()
}

func (i *publisherInfoResponsePtrType) ToPublisherInfoResponsePtrOutput() PublisherInfoResponsePtrOutput {
	return i.ToPublisherInfoResponsePtrOutputWithContext(context.Background())
}

func (i *publisherInfoResponsePtrType) ToPublisherInfoResponsePtrOutputWithContext(ctx context.Context) PublisherInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherInfoResponsePtrOutput)
}

type PublisherInfoResponseOutput struct{ *pulumi.OutputState }

func (PublisherInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublisherInfoResponse)(nil)).Elem()
}

func (o PublisherInfoResponseOutput) ToPublisherInfoResponseOutput() PublisherInfoResponseOutput {
	return o
}

func (o PublisherInfoResponseOutput) ToPublisherInfoResponseOutputWithContext(ctx context.Context) PublisherInfoResponseOutput {
	return o
}

func (o PublisherInfoResponseOutput) ToPublisherInfoResponsePtrOutput() PublisherInfoResponsePtrOutput {
	return o.ToPublisherInfoResponsePtrOutputWithContext(context.Background())
}

func (o PublisherInfoResponseOutput) ToPublisherInfoResponsePtrOutputWithContext(ctx context.Context) PublisherInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublisherInfoResponse) *PublisherInfoResponse {
		return &v
	}).(PublisherInfoResponsePtrOutput)
}

func (o PublisherInfoResponseOutput) BinaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfoResponse) *string { return v.BinaryName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponseOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfoResponse) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponseOutput) PublisherName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfoResponse) *string { return v.PublisherName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfoResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type PublisherInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (PublisherInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublisherInfoResponse)(nil)).Elem()
}

func (o PublisherInfoResponsePtrOutput) ToPublisherInfoResponsePtrOutput() PublisherInfoResponsePtrOutput {
	return o
}

func (o PublisherInfoResponsePtrOutput) ToPublisherInfoResponsePtrOutputWithContext(ctx context.Context) PublisherInfoResponsePtrOutput {
	return o
}

func (o PublisherInfoResponsePtrOutput) Elem() PublisherInfoResponseOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) PublisherInfoResponse {
		if v != nil {
			return *v
		}
		var ret PublisherInfoResponse
		return ret
	}).(PublisherInfoResponseOutput)
}

func (o PublisherInfoResponsePtrOutput) BinaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.BinaryName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponsePtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponsePtrOutput) PublisherName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublisherName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type RecommendationConfigurationProperties struct {
	RecommendationType string `pulumi:"recommendationType"`
	Status             string `pulumi:"status"`
}





type RecommendationConfigurationPropertiesInput interface {
	pulumi.Input

	ToRecommendationConfigurationPropertiesOutput() RecommendationConfigurationPropertiesOutput
	ToRecommendationConfigurationPropertiesOutputWithContext(context.Context) RecommendationConfigurationPropertiesOutput
}

type RecommendationConfigurationPropertiesArgs struct {
	RecommendationType pulumi.StringInput `pulumi:"recommendationType"`
	Status             pulumi.StringInput `pulumi:"status"`
}

func (RecommendationConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigurationProperties)(nil)).Elem()
}

func (i RecommendationConfigurationPropertiesArgs) ToRecommendationConfigurationPropertiesOutput() RecommendationConfigurationPropertiesOutput {
	return i.ToRecommendationConfigurationPropertiesOutputWithContext(context.Background())
}

func (i RecommendationConfigurationPropertiesArgs) ToRecommendationConfigurationPropertiesOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendationConfigurationPropertiesOutput)
}





type RecommendationConfigurationPropertiesArrayInput interface {
	pulumi.Input

	ToRecommendationConfigurationPropertiesArrayOutput() RecommendationConfigurationPropertiesArrayOutput
	ToRecommendationConfigurationPropertiesArrayOutputWithContext(context.Context) RecommendationConfigurationPropertiesArrayOutput
}

type RecommendationConfigurationPropertiesArray []RecommendationConfigurationPropertiesInput

func (RecommendationConfigurationPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendationConfigurationProperties)(nil)).Elem()
}

func (i RecommendationConfigurationPropertiesArray) ToRecommendationConfigurationPropertiesArrayOutput() RecommendationConfigurationPropertiesArrayOutput {
	return i.ToRecommendationConfigurationPropertiesArrayOutputWithContext(context.Background())
}

func (i RecommendationConfigurationPropertiesArray) ToRecommendationConfigurationPropertiesArrayOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendationConfigurationPropertiesArrayOutput)
}

type RecommendationConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (RecommendationConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigurationProperties)(nil)).Elem()
}

func (o RecommendationConfigurationPropertiesOutput) ToRecommendationConfigurationPropertiesOutput() RecommendationConfigurationPropertiesOutput {
	return o
}

func (o RecommendationConfigurationPropertiesOutput) ToRecommendationConfigurationPropertiesOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesOutput {
	return o
}

func (o RecommendationConfigurationPropertiesOutput) RecommendationType() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationProperties) string { return v.RecommendationType }).(pulumi.StringOutput)
}

func (o RecommendationConfigurationPropertiesOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationProperties) string { return v.Status }).(pulumi.StringOutput)
}

type RecommendationConfigurationPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RecommendationConfigurationPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendationConfigurationProperties)(nil)).Elem()
}

func (o RecommendationConfigurationPropertiesArrayOutput) ToRecommendationConfigurationPropertiesArrayOutput() RecommendationConfigurationPropertiesArrayOutput {
	return o
}

func (o RecommendationConfigurationPropertiesArrayOutput) ToRecommendationConfigurationPropertiesArrayOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesArrayOutput {
	return o
}

func (o RecommendationConfigurationPropertiesArrayOutput) Index(i pulumi.IntInput) RecommendationConfigurationPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendationConfigurationProperties {
		return vs[0].([]RecommendationConfigurationProperties)[vs[1].(int)]
	}).(RecommendationConfigurationPropertiesOutput)
}

type RecommendationConfigurationPropertiesResponse struct {
	Name               string `pulumi:"name"`
	RecommendationType string `pulumi:"recommendationType"`
	Status             string `pulumi:"status"`
}





type RecommendationConfigurationPropertiesResponseInput interface {
	pulumi.Input

	ToRecommendationConfigurationPropertiesResponseOutput() RecommendationConfigurationPropertiesResponseOutput
	ToRecommendationConfigurationPropertiesResponseOutputWithContext(context.Context) RecommendationConfigurationPropertiesResponseOutput
}

type RecommendationConfigurationPropertiesResponseArgs struct {
	Name               pulumi.StringInput `pulumi:"name"`
	RecommendationType pulumi.StringInput `pulumi:"recommendationType"`
	Status             pulumi.StringInput `pulumi:"status"`
}

func (RecommendationConfigurationPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigurationPropertiesResponse)(nil)).Elem()
}

func (i RecommendationConfigurationPropertiesResponseArgs) ToRecommendationConfigurationPropertiesResponseOutput() RecommendationConfigurationPropertiesResponseOutput {
	return i.ToRecommendationConfigurationPropertiesResponseOutputWithContext(context.Background())
}

func (i RecommendationConfigurationPropertiesResponseArgs) ToRecommendationConfigurationPropertiesResponseOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendationConfigurationPropertiesResponseOutput)
}





type RecommendationConfigurationPropertiesResponseArrayInput interface {
	pulumi.Input

	ToRecommendationConfigurationPropertiesResponseArrayOutput() RecommendationConfigurationPropertiesResponseArrayOutput
	ToRecommendationConfigurationPropertiesResponseArrayOutputWithContext(context.Context) RecommendationConfigurationPropertiesResponseArrayOutput
}

type RecommendationConfigurationPropertiesResponseArray []RecommendationConfigurationPropertiesResponseInput

func (RecommendationConfigurationPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendationConfigurationPropertiesResponse)(nil)).Elem()
}

func (i RecommendationConfigurationPropertiesResponseArray) ToRecommendationConfigurationPropertiesResponseArrayOutput() RecommendationConfigurationPropertiesResponseArrayOutput {
	return i.ToRecommendationConfigurationPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i RecommendationConfigurationPropertiesResponseArray) ToRecommendationConfigurationPropertiesResponseArrayOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendationConfigurationPropertiesResponseArrayOutput)
}

type RecommendationConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RecommendationConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigurationPropertiesResponse)(nil)).Elem()
}

func (o RecommendationConfigurationPropertiesResponseOutput) ToRecommendationConfigurationPropertiesResponseOutput() RecommendationConfigurationPropertiesResponseOutput {
	return o
}

func (o RecommendationConfigurationPropertiesResponseOutput) ToRecommendationConfigurationPropertiesResponseOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesResponseOutput {
	return o
}

func (o RecommendationConfigurationPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RecommendationConfigurationPropertiesResponseOutput) RecommendationType() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationPropertiesResponse) string { return v.RecommendationType }).(pulumi.StringOutput)
}

func (o RecommendationConfigurationPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v RecommendationConfigurationPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type RecommendationConfigurationPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RecommendationConfigurationPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecommendationConfigurationPropertiesResponse)(nil)).Elem()
}

func (o RecommendationConfigurationPropertiesResponseArrayOutput) ToRecommendationConfigurationPropertiesResponseArrayOutput() RecommendationConfigurationPropertiesResponseArrayOutput {
	return o
}

func (o RecommendationConfigurationPropertiesResponseArrayOutput) ToRecommendationConfigurationPropertiesResponseArrayOutputWithContext(ctx context.Context) RecommendationConfigurationPropertiesResponseArrayOutput {
	return o
}

func (o RecommendationConfigurationPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RecommendationConfigurationPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecommendationConfigurationPropertiesResponse {
		return vs[0].([]RecommendationConfigurationPropertiesResponse)[vs[1].(int)]
	}).(RecommendationConfigurationPropertiesResponseOutput)
}

type RuleResultsPropertiesResponse struct {
	Results [][]string `pulumi:"results"`
}





type RuleResultsPropertiesResponseInput interface {
	pulumi.Input

	ToRuleResultsPropertiesResponseOutput() RuleResultsPropertiesResponseOutput
	ToRuleResultsPropertiesResponseOutputWithContext(context.Context) RuleResultsPropertiesResponseOutput
}

type RuleResultsPropertiesResponseArgs struct {
	Results pulumi.StringArrayArrayInput `pulumi:"results"`
}

func (RuleResultsPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleResultsPropertiesResponse)(nil)).Elem()
}

func (i RuleResultsPropertiesResponseArgs) ToRuleResultsPropertiesResponseOutput() RuleResultsPropertiesResponseOutput {
	return i.ToRuleResultsPropertiesResponseOutputWithContext(context.Background())
}

func (i RuleResultsPropertiesResponseArgs) ToRuleResultsPropertiesResponseOutputWithContext(ctx context.Context) RuleResultsPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleResultsPropertiesResponseOutput)
}

func (i RuleResultsPropertiesResponseArgs) ToRuleResultsPropertiesResponsePtrOutput() RuleResultsPropertiesResponsePtrOutput {
	return i.ToRuleResultsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i RuleResultsPropertiesResponseArgs) ToRuleResultsPropertiesResponsePtrOutputWithContext(ctx context.Context) RuleResultsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleResultsPropertiesResponseOutput).ToRuleResultsPropertiesResponsePtrOutputWithContext(ctx)
}









type RuleResultsPropertiesResponsePtrInput interface {
	pulumi.Input

	ToRuleResultsPropertiesResponsePtrOutput() RuleResultsPropertiesResponsePtrOutput
	ToRuleResultsPropertiesResponsePtrOutputWithContext(context.Context) RuleResultsPropertiesResponsePtrOutput
}

type ruleResultsPropertiesResponsePtrType RuleResultsPropertiesResponseArgs

func RuleResultsPropertiesResponsePtr(v *RuleResultsPropertiesResponseArgs) RuleResultsPropertiesResponsePtrInput {
	return (*ruleResultsPropertiesResponsePtrType)(v)
}

func (*ruleResultsPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleResultsPropertiesResponse)(nil)).Elem()
}

func (i *ruleResultsPropertiesResponsePtrType) ToRuleResultsPropertiesResponsePtrOutput() RuleResultsPropertiesResponsePtrOutput {
	return i.ToRuleResultsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *ruleResultsPropertiesResponsePtrType) ToRuleResultsPropertiesResponsePtrOutputWithContext(ctx context.Context) RuleResultsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleResultsPropertiesResponsePtrOutput)
}

type RuleResultsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RuleResultsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleResultsPropertiesResponse)(nil)).Elem()
}

func (o RuleResultsPropertiesResponseOutput) ToRuleResultsPropertiesResponseOutput() RuleResultsPropertiesResponseOutput {
	return o
}

func (o RuleResultsPropertiesResponseOutput) ToRuleResultsPropertiesResponseOutputWithContext(ctx context.Context) RuleResultsPropertiesResponseOutput {
	return o
}

func (o RuleResultsPropertiesResponseOutput) ToRuleResultsPropertiesResponsePtrOutput() RuleResultsPropertiesResponsePtrOutput {
	return o.ToRuleResultsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o RuleResultsPropertiesResponseOutput) ToRuleResultsPropertiesResponsePtrOutputWithContext(ctx context.Context) RuleResultsPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleResultsPropertiesResponse) *RuleResultsPropertiesResponse {
		return &v
	}).(RuleResultsPropertiesResponsePtrOutput)
}

func (o RuleResultsPropertiesResponseOutput) Results() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v RuleResultsPropertiesResponse) [][]string { return v.Results }).(pulumi.StringArrayArrayOutput)
}

type RuleResultsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RuleResultsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleResultsPropertiesResponse)(nil)).Elem()
}

func (o RuleResultsPropertiesResponsePtrOutput) ToRuleResultsPropertiesResponsePtrOutput() RuleResultsPropertiesResponsePtrOutput {
	return o
}

func (o RuleResultsPropertiesResponsePtrOutput) ToRuleResultsPropertiesResponsePtrOutputWithContext(ctx context.Context) RuleResultsPropertiesResponsePtrOutput {
	return o
}

func (o RuleResultsPropertiesResponsePtrOutput) Elem() RuleResultsPropertiesResponseOutput {
	return o.ApplyT(func(v *RuleResultsPropertiesResponse) RuleResultsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret RuleResultsPropertiesResponse
		return ret
	}).(RuleResultsPropertiesResponseOutput)
}

func (o RuleResultsPropertiesResponsePtrOutput) Results() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v *RuleResultsPropertiesResponse) [][]string {
		if v == nil {
			return nil
		}
		return v.Results
	}).(pulumi.StringArrayArrayOutput)
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





type ScopeElementResponseInput interface {
	pulumi.Input

	ToScopeElementResponseOutput() ScopeElementResponseOutput
	ToScopeElementResponseOutputWithContext(context.Context) ScopeElementResponseOutput
}

type ScopeElementResponseArgs struct {
	Field pulumi.StringPtrInput `pulumi:"field"`
}

func (ScopeElementResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeElementResponse)(nil)).Elem()
}

func (i ScopeElementResponseArgs) ToScopeElementResponseOutput() ScopeElementResponseOutput {
	return i.ToScopeElementResponseOutputWithContext(context.Background())
}

func (i ScopeElementResponseArgs) ToScopeElementResponseOutputWithContext(ctx context.Context) ScopeElementResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeElementResponseOutput)
}





type ScopeElementResponseArrayInput interface {
	pulumi.Input

	ToScopeElementResponseArrayOutput() ScopeElementResponseArrayOutput
	ToScopeElementResponseArrayOutputWithContext(context.Context) ScopeElementResponseArrayOutput
}

type ScopeElementResponseArray []ScopeElementResponseInput

func (ScopeElementResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScopeElementResponse)(nil)).Elem()
}

func (i ScopeElementResponseArray) ToScopeElementResponseArrayOutput() ScopeElementResponseArrayOutput {
	return i.ToScopeElementResponseArrayOutputWithContext(context.Background())
}

func (i ScopeElementResponseArray) ToScopeElementResponseArrayOutputWithContext(ctx context.Context) ScopeElementResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeElementResponseArrayOutput)
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

type SecurityAssessmentMetadataPartnerData struct {
	PartnerName string  `pulumi:"partnerName"`
	ProductName *string `pulumi:"productName"`
	Secret      string  `pulumi:"secret"`
}





type SecurityAssessmentMetadataPartnerDataInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPartnerDataOutput() SecurityAssessmentMetadataPartnerDataOutput
	ToSecurityAssessmentMetadataPartnerDataOutputWithContext(context.Context) SecurityAssessmentMetadataPartnerDataOutput
}

type SecurityAssessmentMetadataPartnerDataArgs struct {
	PartnerName pulumi.StringInput    `pulumi:"partnerName"`
	ProductName pulumi.StringPtrInput `pulumi:"productName"`
	Secret      pulumi.StringInput    `pulumi:"secret"`
}

func (SecurityAssessmentMetadataPartnerDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataOutput() SecurityAssessmentMetadataPartnerDataOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataOutput)
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataOutput).ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx)
}









type SecurityAssessmentMetadataPartnerDataPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput
	ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput
}

type securityAssessmentMetadataPartnerDataPtrType SecurityAssessmentMetadataPartnerDataArgs

func SecurityAssessmentMetadataPartnerDataPtr(v *SecurityAssessmentMetadataPartnerDataArgs) SecurityAssessmentMetadataPartnerDataPtrInput {
	return (*securityAssessmentMetadataPartnerDataPtrType)(v)
}

func (*securityAssessmentMetadataPartnerDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (i *securityAssessmentMetadataPartnerDataPtrType) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentMetadataPartnerDataPtrType) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

type SecurityAssessmentMetadataPartnerDataOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataOutput() SecurityAssessmentMetadataPartnerDataOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentMetadataPartnerData) *SecurityAssessmentMetadataPartnerData {
		return &v
	}).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerData) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerData) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerData) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentMetadataPartnerDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) Elem() SecurityAssessmentMetadataPartnerDataOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) SecurityAssessmentMetadataPartnerData {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPartnerData
		return ret
	}).(SecurityAssessmentMetadataPartnerDataOutput)
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPartnerDataResponse struct {
	PartnerName string  `pulumi:"partnerName"`
	ProductName *string `pulumi:"productName"`
	Secret      string  `pulumi:"secret"`
}





type SecurityAssessmentMetadataPartnerDataResponseInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPartnerDataResponseOutput() SecurityAssessmentMetadataPartnerDataResponseOutput
	ToSecurityAssessmentMetadataPartnerDataResponseOutputWithContext(context.Context) SecurityAssessmentMetadataPartnerDataResponseOutput
}

type SecurityAssessmentMetadataPartnerDataResponseArgs struct {
	PartnerName pulumi.StringInput    `pulumi:"partnerName"`
	ProductName pulumi.StringPtrInput `pulumi:"productName"`
	Secret      pulumi.StringInput    `pulumi:"secret"`
}

func (SecurityAssessmentMetadataPartnerDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerDataResponse)(nil)).Elem()
}

func (i SecurityAssessmentMetadataPartnerDataResponseArgs) ToSecurityAssessmentMetadataPartnerDataResponseOutput() SecurityAssessmentMetadataPartnerDataResponseOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataResponseOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPartnerDataResponseArgs) ToSecurityAssessmentMetadataPartnerDataResponseOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataResponseOutput)
}

func (i SecurityAssessmentMetadataPartnerDataResponseArgs) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutput() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPartnerDataResponseArgs) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataResponseOutput).ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(ctx)
}









type SecurityAssessmentMetadataPartnerDataResponsePtrInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPartnerDataResponsePtrOutput() SecurityAssessmentMetadataPartnerDataResponsePtrOutput
	ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(context.Context) SecurityAssessmentMetadataPartnerDataResponsePtrOutput
}

type securityAssessmentMetadataPartnerDataResponsePtrType SecurityAssessmentMetadataPartnerDataResponseArgs

func SecurityAssessmentMetadataPartnerDataResponsePtr(v *SecurityAssessmentMetadataPartnerDataResponseArgs) SecurityAssessmentMetadataPartnerDataResponsePtrInput {
	return (*securityAssessmentMetadataPartnerDataResponsePtrType)(v)
}

func (*securityAssessmentMetadataPartnerDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerDataResponse)(nil)).Elem()
}

func (i *securityAssessmentMetadataPartnerDataResponsePtrType) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutput() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(context.Background())
}

func (i *securityAssessmentMetadataPartnerDataResponsePtrType) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

type SecurityAssessmentMetadataPartnerDataResponseOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ToSecurityAssessmentMetadataPartnerDataResponseOutput() SecurityAssessmentMetadataPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ToSecurityAssessmentMetadataPartnerDataResponseOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutput() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentMetadataPartnerDataResponse) *SecurityAssessmentMetadataPartnerDataResponse {
		return &v
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerDataResponse) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerDataResponse) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerDataResponse) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentMetadataPartnerDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutput() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) Elem() SecurityAssessmentMetadataPartnerDataResponseOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) SecurityAssessmentMetadataPartnerDataResponse {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPartnerDataResponse
		return ret
	}).(SecurityAssessmentMetadataPartnerDataResponseOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataProperties struct {
	AssessmentType         string                                 `pulumi:"assessmentType"`
	Categories             []string                               `pulumi:"categories"`
	Description            *string                                `pulumi:"description"`
	DisplayName            string                                 `pulumi:"displayName"`
	ImplementationEffort   *string                                `pulumi:"implementationEffort"`
	PartnerData            *SecurityAssessmentMetadataPartnerData `pulumi:"partnerData"`
	Preview                *bool                                  `pulumi:"preview"`
	RemediationDescription *string                                `pulumi:"remediationDescription"`
	Severity               string                                 `pulumi:"severity"`
	Threats                []string                               `pulumi:"threats"`
	UserImpact             *string                                `pulumi:"userImpact"`
}





type SecurityAssessmentMetadataPropertiesInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesOutput() SecurityAssessmentMetadataPropertiesOutput
	ToSecurityAssessmentMetadataPropertiesOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesOutput
}

type SecurityAssessmentMetadataPropertiesArgs struct {
	AssessmentType         pulumi.StringInput                            `pulumi:"assessmentType"`
	Categories             pulumi.StringArrayInput                       `pulumi:"categories"`
	Description            pulumi.StringPtrInput                         `pulumi:"description"`
	DisplayName            pulumi.StringInput                            `pulumi:"displayName"`
	ImplementationEffort   pulumi.StringPtrInput                         `pulumi:"implementationEffort"`
	PartnerData            SecurityAssessmentMetadataPartnerDataPtrInput `pulumi:"partnerData"`
	Preview                pulumi.BoolPtrInput                           `pulumi:"preview"`
	RemediationDescription pulumi.StringPtrInput                         `pulumi:"remediationDescription"`
	Severity               pulumi.StringInput                            `pulumi:"severity"`
	Threats                pulumi.StringArrayInput                       `pulumi:"threats"`
	UserImpact             pulumi.StringPtrInput                         `pulumi:"userImpact"`
}

func (SecurityAssessmentMetadataPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesOutput() SecurityAssessmentMetadataPropertiesOutput {
	return i.ToSecurityAssessmentMetadataPropertiesOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesOutput)
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesOutput).ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx)
}









type SecurityAssessmentMetadataPropertiesPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput
	ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesPtrOutput
}

type securityAssessmentMetadataPropertiesPtrType SecurityAssessmentMetadataPropertiesArgs

func SecurityAssessmentMetadataPropertiesPtr(v *SecurityAssessmentMetadataPropertiesArgs) SecurityAssessmentMetadataPropertiesPtrInput {
	return (*securityAssessmentMetadataPropertiesPtrType)(v)
}

func (*securityAssessmentMetadataPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (i *securityAssessmentMetadataPropertiesPtrType) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentMetadataPropertiesPtrType) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesPtrOutput)
}

type SecurityAssessmentMetadataPropertiesOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesOutput() SecurityAssessmentMetadataPropertiesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return o.ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentMetadataProperties) *SecurityAssessmentMetadataProperties {
		return &v
	}).(SecurityAssessmentMetadataPropertiesPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) string { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) PartnerData() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *SecurityAssessmentMetadataPartnerData {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) string { return v.Severity }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) []string { return v.Threats }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.UserImpact }).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Elem() SecurityAssessmentMetadataPropertiesOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) SecurityAssessmentMetadataProperties {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataProperties
		return ret
	}).(SecurityAssessmentMetadataPropertiesOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) AssessmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AssessmentType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationEffort
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) PartnerData() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *SecurityAssessmentMetadataPartnerData {
		if v == nil {
			return nil
		}
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Preview
	}).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.RemediationDescription
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Severity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) []string {
		if v == nil {
			return nil
		}
		return v.Threats
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.UserImpact
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponse struct {
	AssessmentType         string                                         `pulumi:"assessmentType"`
	Categories             []string                                       `pulumi:"categories"`
	Description            *string                                        `pulumi:"description"`
	DisplayName            string                                         `pulumi:"displayName"`
	ImplementationEffort   *string                                        `pulumi:"implementationEffort"`
	PartnerData            *SecurityAssessmentMetadataPartnerDataResponse `pulumi:"partnerData"`
	PolicyDefinitionId     string                                         `pulumi:"policyDefinitionId"`
	Preview                *bool                                          `pulumi:"preview"`
	RemediationDescription *string                                        `pulumi:"remediationDescription"`
	Severity               string                                         `pulumi:"severity"`
	Threats                []string                                       `pulumi:"threats"`
	UserImpact             *string                                        `pulumi:"userImpact"`
}





type SecurityAssessmentMetadataPropertiesResponseInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesResponseOutput() SecurityAssessmentMetadataPropertiesResponseOutput
	ToSecurityAssessmentMetadataPropertiesResponseOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesResponseOutput
}

type SecurityAssessmentMetadataPropertiesResponseArgs struct {
	AssessmentType         pulumi.StringInput                                    `pulumi:"assessmentType"`
	Categories             pulumi.StringArrayInput                               `pulumi:"categories"`
	Description            pulumi.StringPtrInput                                 `pulumi:"description"`
	DisplayName            pulumi.StringInput                                    `pulumi:"displayName"`
	ImplementationEffort   pulumi.StringPtrInput                                 `pulumi:"implementationEffort"`
	PartnerData            SecurityAssessmentMetadataPartnerDataResponsePtrInput `pulumi:"partnerData"`
	PolicyDefinitionId     pulumi.StringInput                                    `pulumi:"policyDefinitionId"`
	Preview                pulumi.BoolPtrInput                                   `pulumi:"preview"`
	RemediationDescription pulumi.StringPtrInput                                 `pulumi:"remediationDescription"`
	Severity               pulumi.StringInput                                    `pulumi:"severity"`
	Threats                pulumi.StringArrayInput                               `pulumi:"threats"`
	UserImpact             pulumi.StringPtrInput                                 `pulumi:"userImpact"`
}

func (SecurityAssessmentMetadataPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPropertiesResponse)(nil)).Elem()
}

func (i SecurityAssessmentMetadataPropertiesResponseArgs) ToSecurityAssessmentMetadataPropertiesResponseOutput() SecurityAssessmentMetadataPropertiesResponseOutput {
	return i.ToSecurityAssessmentMetadataPropertiesResponseOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesResponseArgs) ToSecurityAssessmentMetadataPropertiesResponseOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesResponseOutput)
}

func (i SecurityAssessmentMetadataPropertiesResponseArgs) ToSecurityAssessmentMetadataPropertiesResponsePtrOutput() SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesResponseArgs) ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesResponseOutput).ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(ctx)
}









type SecurityAssessmentMetadataPropertiesResponsePtrInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesResponsePtrOutput() SecurityAssessmentMetadataPropertiesResponsePtrOutput
	ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesResponsePtrOutput
}

type securityAssessmentMetadataPropertiesResponsePtrType SecurityAssessmentMetadataPropertiesResponseArgs

func SecurityAssessmentMetadataPropertiesResponsePtr(v *SecurityAssessmentMetadataPropertiesResponseArgs) SecurityAssessmentMetadataPropertiesResponsePtrInput {
	return (*securityAssessmentMetadataPropertiesResponsePtrType)(v)
}

func (*securityAssessmentMetadataPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPropertiesResponse)(nil)).Elem()
}

func (i *securityAssessmentMetadataPropertiesResponsePtrType) ToSecurityAssessmentMetadataPropertiesResponsePtrOutput() SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *securityAssessmentMetadataPropertiesResponsePtrType) ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesResponsePtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPropertiesResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ToSecurityAssessmentMetadataPropertiesResponseOutput() SecurityAssessmentMetadataPropertiesResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ToSecurityAssessmentMetadataPropertiesResponseOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ToSecurityAssessmentMetadataPropertiesResponsePtrOutput() SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o.ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentMetadataPropertiesResponse) *SecurityAssessmentMetadataPropertiesResponse {
		return &v
	}).(SecurityAssessmentMetadataPropertiesResponsePtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *SecurityAssessmentMetadataPartnerDataResponse {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.Severity }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) []string { return v.Threats }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.UserImpact }).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPropertiesResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) ToSecurityAssessmentMetadataPropertiesResponsePtrOutput() SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Elem() SecurityAssessmentMetadataPropertiesResponseOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) SecurityAssessmentMetadataPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPropertiesResponse
		return ret
	}).(SecurityAssessmentMetadataPropertiesResponseOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) AssessmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AssessmentType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationEffort
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *SecurityAssessmentMetadataPartnerDataResponse {
		if v == nil {
			return nil
		}
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PolicyDefinitionId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Preview
	}).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RemediationDescription
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Severity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Threats
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserImpact
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentPartnerData struct {
	PartnerName string `pulumi:"partnerName"`
	Secret      string `pulumi:"secret"`
}





type SecurityAssessmentPartnerDataInput interface {
	pulumi.Input

	ToSecurityAssessmentPartnerDataOutput() SecurityAssessmentPartnerDataOutput
	ToSecurityAssessmentPartnerDataOutputWithContext(context.Context) SecurityAssessmentPartnerDataOutput
}

type SecurityAssessmentPartnerDataArgs struct {
	PartnerName pulumi.StringInput `pulumi:"partnerName"`
	Secret      pulumi.StringInput `pulumi:"secret"`
}

func (SecurityAssessmentPartnerDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerData)(nil)).Elem()
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataOutput() SecurityAssessmentPartnerDataOutput {
	return i.ToSecurityAssessmentPartnerDataOutputWithContext(context.Background())
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataOutput)
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return i.ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataOutput).ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx)
}









type SecurityAssessmentPartnerDataPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput
	ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Context) SecurityAssessmentPartnerDataPtrOutput
}

type securityAssessmentPartnerDataPtrType SecurityAssessmentPartnerDataArgs

func SecurityAssessmentPartnerDataPtr(v *SecurityAssessmentPartnerDataArgs) SecurityAssessmentPartnerDataPtrInput {
	return (*securityAssessmentPartnerDataPtrType)(v)
}

func (*securityAssessmentPartnerDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerData)(nil)).Elem()
}

func (i *securityAssessmentPartnerDataPtrType) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return i.ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentPartnerDataPtrType) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataPtrOutput)
}

type SecurityAssessmentPartnerDataOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataOutput() SecurityAssessmentPartnerDataOutput {
	return o
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataOutput {
	return o
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return o.ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentPartnerData) *SecurityAssessmentPartnerData {
		return &v
	}).(SecurityAssessmentPartnerDataPtrOutput)
}

func (o SecurityAssessmentPartnerDataOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerData) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentPartnerDataOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerData) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentPartnerDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataPtrOutput) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataPtrOutput) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataPtrOutput) Elem() SecurityAssessmentPartnerDataOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerData) SecurityAssessmentPartnerData {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentPartnerData
		return ret
	}).(SecurityAssessmentPartnerDataOutput)
}

func (o SecurityAssessmentPartnerDataPtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentPartnerDataPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentPartnerDataResponse struct {
	PartnerName string `pulumi:"partnerName"`
	Secret      string `pulumi:"secret"`
}





type SecurityAssessmentPartnerDataResponseInput interface {
	pulumi.Input

	ToSecurityAssessmentPartnerDataResponseOutput() SecurityAssessmentPartnerDataResponseOutput
	ToSecurityAssessmentPartnerDataResponseOutputWithContext(context.Context) SecurityAssessmentPartnerDataResponseOutput
}

type SecurityAssessmentPartnerDataResponseArgs struct {
	PartnerName pulumi.StringInput `pulumi:"partnerName"`
	Secret      pulumi.StringInput `pulumi:"secret"`
}

func (SecurityAssessmentPartnerDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerDataResponse)(nil)).Elem()
}

func (i SecurityAssessmentPartnerDataResponseArgs) ToSecurityAssessmentPartnerDataResponseOutput() SecurityAssessmentPartnerDataResponseOutput {
	return i.ToSecurityAssessmentPartnerDataResponseOutputWithContext(context.Background())
}

func (i SecurityAssessmentPartnerDataResponseArgs) ToSecurityAssessmentPartnerDataResponseOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataResponseOutput)
}

func (i SecurityAssessmentPartnerDataResponseArgs) ToSecurityAssessmentPartnerDataResponsePtrOutput() SecurityAssessmentPartnerDataResponsePtrOutput {
	return i.ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentPartnerDataResponseArgs) ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataResponseOutput).ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(ctx)
}









type SecurityAssessmentPartnerDataResponsePtrInput interface {
	pulumi.Input

	ToSecurityAssessmentPartnerDataResponsePtrOutput() SecurityAssessmentPartnerDataResponsePtrOutput
	ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(context.Context) SecurityAssessmentPartnerDataResponsePtrOutput
}

type securityAssessmentPartnerDataResponsePtrType SecurityAssessmentPartnerDataResponseArgs

func SecurityAssessmentPartnerDataResponsePtr(v *SecurityAssessmentPartnerDataResponseArgs) SecurityAssessmentPartnerDataResponsePtrInput {
	return (*securityAssessmentPartnerDataResponsePtrType)(v)
}

func (*securityAssessmentPartnerDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerDataResponse)(nil)).Elem()
}

func (i *securityAssessmentPartnerDataResponsePtrType) ToSecurityAssessmentPartnerDataResponsePtrOutput() SecurityAssessmentPartnerDataResponsePtrOutput {
	return i.ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(context.Background())
}

func (i *securityAssessmentPartnerDataResponsePtrType) ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataResponsePtrOutput)
}

type SecurityAssessmentPartnerDataResponseOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataResponseOutput) ToSecurityAssessmentPartnerDataResponseOutput() SecurityAssessmentPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponseOutput) ToSecurityAssessmentPartnerDataResponseOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponseOutput) ToSecurityAssessmentPartnerDataResponsePtrOutput() SecurityAssessmentPartnerDataResponsePtrOutput {
	return o.ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentPartnerDataResponseOutput) ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentPartnerDataResponse) *SecurityAssessmentPartnerDataResponse {
		return &v
	}).(SecurityAssessmentPartnerDataResponsePtrOutput)
}

func (o SecurityAssessmentPartnerDataResponseOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerDataResponse) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentPartnerDataResponseOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerDataResponse) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentPartnerDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) ToSecurityAssessmentPartnerDataResponsePtrOutput() SecurityAssessmentPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) Elem() SecurityAssessmentPartnerDataResponseOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerDataResponse) SecurityAssessmentPartnerDataResponse {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentPartnerDataResponse
		return ret
	}).(SecurityAssessmentPartnerDataResponseOutput)
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesOrganizationalData struct {
	ExcludedAccountIds         []string `pulumi:"excludedAccountIds"`
	OrganizationMembershipType *string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string  `pulumi:"parentHierarchyId"`
	StacksetName               *string  `pulumi:"stacksetName"`
}





type SecurityConnectorPropertiesOrganizationalDataInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesOrganizationalDataOutput() SecurityConnectorPropertiesOrganizationalDataOutput
	ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(context.Context) SecurityConnectorPropertiesOrganizationalDataOutput
}

type SecurityConnectorPropertiesOrganizationalDataArgs struct {
	ExcludedAccountIds         pulumi.StringArrayInput `pulumi:"excludedAccountIds"`
	OrganizationMembershipType pulumi.StringPtrInput   `pulumi:"organizationMembershipType"`
	ParentHierarchyId          pulumi.StringPtrInput   `pulumi:"parentHierarchyId"`
	StacksetName               pulumi.StringPtrInput   `pulumi:"stacksetName"`
}

func (SecurityConnectorPropertiesOrganizationalDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataOutput() SecurityConnectorPropertiesOrganizationalDataOutput {
	return i.ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesOrganizationalDataOutput)
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesOrganizationalDataArgs) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesOrganizationalDataOutput).ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx)
}









type SecurityConnectorPropertiesOrganizationalDataPtrInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput
	ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput
}

type securityConnectorPropertiesOrganizationalDataPtrType SecurityConnectorPropertiesOrganizationalDataArgs

func SecurityConnectorPropertiesOrganizationalDataPtr(v *SecurityConnectorPropertiesOrganizationalDataArgs) SecurityConnectorPropertiesOrganizationalDataPtrInput {
	return (*securityConnectorPropertiesOrganizationalDataPtrType)(v)
}

func (*securityConnectorPropertiesOrganizationalDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (i *securityConnectorPropertiesOrganizationalDataPtrType) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i *securityConnectorPropertiesOrganizationalDataPtrType) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesOrganizationalDataPtrOutput)
}

type SecurityConnectorPropertiesOrganizationalDataOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesOrganizationalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataOutput() SecurityConnectorPropertiesOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o.ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(context.Background())
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityConnectorPropertiesOrganizationalData) *SecurityConnectorPropertiesOrganizationalData {
		return &v
	}).(SecurityConnectorPropertiesOrganizationalDataPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ExcludedAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) []string { return v.ExcludedAccountIds }).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) *string { return v.OrganizationMembershipType }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) *string { return v.ParentHierarchyId }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesOrganizationalData) *string { return v.StacksetName }).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesOrganizationalDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesOrganizationalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutput() SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) Elem() SecurityConnectorPropertiesOrganizationalDataOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) SecurityConnectorPropertiesOrganizationalData {
		if v != nil {
			return *v
		}
		var ret SecurityConnectorPropertiesOrganizationalData
		return ret
	}).(SecurityConnectorPropertiesOrganizationalDataOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ExcludedAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedAccountIds
	}).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationMembershipType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.ParentHierarchyId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesOrganizationalDataPtrOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.StacksetName
	}).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesResponseOrganizationalData struct {
	ExcludedAccountIds         []string `pulumi:"excludedAccountIds"`
	OrganizationMembershipType *string  `pulumi:"organizationMembershipType"`
	ParentHierarchyId          *string  `pulumi:"parentHierarchyId"`
	StacksetName               *string  `pulumi:"stacksetName"`
}





type SecurityConnectorPropertiesResponseOrganizationalDataInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesResponseOrganizationalDataOutput() SecurityConnectorPropertiesResponseOrganizationalDataOutput
	ToSecurityConnectorPropertiesResponseOrganizationalDataOutputWithContext(context.Context) SecurityConnectorPropertiesResponseOrganizationalDataOutput
}

type SecurityConnectorPropertiesResponseOrganizationalDataArgs struct {
	ExcludedAccountIds         pulumi.StringArrayInput `pulumi:"excludedAccountIds"`
	OrganizationMembershipType pulumi.StringPtrInput   `pulumi:"organizationMembershipType"`
	ParentHierarchyId          pulumi.StringPtrInput   `pulumi:"parentHierarchyId"`
	StacksetName               pulumi.StringPtrInput   `pulumi:"stacksetName"`
}

func (SecurityConnectorPropertiesResponseOrganizationalDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (i SecurityConnectorPropertiesResponseOrganizationalDataArgs) ToSecurityConnectorPropertiesResponseOrganizationalDataOutput() SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return i.ToSecurityConnectorPropertiesResponseOrganizationalDataOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesResponseOrganizationalDataArgs) ToSecurityConnectorPropertiesResponseOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesResponseOrganizationalDataOutput)
}

func (i SecurityConnectorPropertiesResponseOrganizationalDataArgs) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i SecurityConnectorPropertiesResponseOrganizationalDataArgs) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesResponseOrganizationalDataOutput).ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx)
}









type SecurityConnectorPropertiesResponseOrganizationalDataPtrInput interface {
	pulumi.Input

	ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput
	ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput
}

type securityConnectorPropertiesResponseOrganizationalDataPtrType SecurityConnectorPropertiesResponseOrganizationalDataArgs

func SecurityConnectorPropertiesResponseOrganizationalDataPtr(v *SecurityConnectorPropertiesResponseOrganizationalDataArgs) SecurityConnectorPropertiesResponseOrganizationalDataPtrInput {
	return (*securityConnectorPropertiesResponseOrganizationalDataPtrType)(v)
}

func (*securityConnectorPropertiesResponseOrganizationalDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (i *securityConnectorPropertiesResponseOrganizationalDataPtrType) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return i.ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(context.Background())
}

func (i *securityConnectorPropertiesResponseOrganizationalDataPtrType) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput)
}

type SecurityConnectorPropertiesResponseOrganizationalDataOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesResponseOrganizationalDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataOutput() SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o.ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(context.Background())
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityConnectorPropertiesResponseOrganizationalData) *SecurityConnectorPropertiesResponseOrganizationalData {
		return &v
	}).(SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ExcludedAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) []string { return v.ExcludedAccountIds }).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) *string {
		return v.OrganizationMembershipType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) *string { return v.ParentHierarchyId }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityConnectorPropertiesResponseOrganizationalData) *string { return v.StacksetName }).(pulumi.StringPtrOutput)
}

type SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorPropertiesResponseOrganizationalData)(nil)).Elem()
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutput() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ToSecurityConnectorPropertiesResponseOrganizationalDataPtrOutputWithContext(ctx context.Context) SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) Elem() SecurityConnectorPropertiesResponseOrganizationalDataOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) SecurityConnectorPropertiesResponseOrganizationalData {
		if v != nil {
			return *v
		}
		var ret SecurityConnectorPropertiesResponseOrganizationalData
		return ret
	}).(SecurityConnectorPropertiesResponseOrganizationalDataOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ExcludedAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedAccountIds
	}).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) OrganizationMembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationMembershipType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) ParentHierarchyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.ParentHierarchyId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput) StacksetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorPropertiesResponseOrganizationalData) *string {
		if v == nil {
			return nil
		}
		return v.StacksetName
	}).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesAlertNotifications struct {
	MinimalSeverity *string `pulumi:"minimalSeverity"`
	State           *string `pulumi:"state"`
}





type SecurityContactPropertiesAlertNotificationsInput interface {
	pulumi.Input

	ToSecurityContactPropertiesAlertNotificationsOutput() SecurityContactPropertiesAlertNotificationsOutput
	ToSecurityContactPropertiesAlertNotificationsOutputWithContext(context.Context) SecurityContactPropertiesAlertNotificationsOutput
}

type SecurityContactPropertiesAlertNotificationsArgs struct {
	MinimalSeverity pulumi.StringPtrInput `pulumi:"minimalSeverity"`
	State           pulumi.StringPtrInput `pulumi:"state"`
}

func (SecurityContactPropertiesAlertNotificationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesAlertNotifications)(nil)).Elem()
}

func (i SecurityContactPropertiesAlertNotificationsArgs) ToSecurityContactPropertiesAlertNotificationsOutput() SecurityContactPropertiesAlertNotificationsOutput {
	return i.ToSecurityContactPropertiesAlertNotificationsOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesAlertNotificationsArgs) ToSecurityContactPropertiesAlertNotificationsOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesAlertNotificationsOutput)
}

func (i SecurityContactPropertiesAlertNotificationsArgs) ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput {
	return i.ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesAlertNotificationsArgs) ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesAlertNotificationsOutput).ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx)
}









type SecurityContactPropertiesAlertNotificationsPtrInput interface {
	pulumi.Input

	ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput
	ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput
}

type securityContactPropertiesAlertNotificationsPtrType SecurityContactPropertiesAlertNotificationsArgs

func SecurityContactPropertiesAlertNotificationsPtr(v *SecurityContactPropertiesAlertNotificationsArgs) SecurityContactPropertiesAlertNotificationsPtrInput {
	return (*securityContactPropertiesAlertNotificationsPtrType)(v)
}

func (*securityContactPropertiesAlertNotificationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesAlertNotifications)(nil)).Elem()
}

func (i *securityContactPropertiesAlertNotificationsPtrType) ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput {
	return i.ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(context.Background())
}

func (i *securityContactPropertiesAlertNotificationsPtrType) ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesAlertNotificationsPtrOutput)
}

type SecurityContactPropertiesAlertNotificationsOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesAlertNotificationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesAlertNotifications)(nil)).Elem()
}

func (o SecurityContactPropertiesAlertNotificationsOutput) ToSecurityContactPropertiesAlertNotificationsOutput() SecurityContactPropertiesAlertNotificationsOutput {
	return o
}

func (o SecurityContactPropertiesAlertNotificationsOutput) ToSecurityContactPropertiesAlertNotificationsOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsOutput {
	return o
}

func (o SecurityContactPropertiesAlertNotificationsOutput) ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput {
	return o.ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(context.Background())
}

func (o SecurityContactPropertiesAlertNotificationsOutput) ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityContactPropertiesAlertNotifications) *SecurityContactPropertiesAlertNotifications {
		return &v
	}).(SecurityContactPropertiesAlertNotificationsPtrOutput)
}

func (o SecurityContactPropertiesAlertNotificationsOutput) MinimalSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesAlertNotifications) *string { return v.MinimalSeverity }).(pulumi.StringPtrOutput)
}

func (o SecurityContactPropertiesAlertNotificationsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesAlertNotifications) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesAlertNotificationsPtrOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesAlertNotificationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesAlertNotifications)(nil)).Elem()
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) ToSecurityContactPropertiesAlertNotificationsPtrOutput() SecurityContactPropertiesAlertNotificationsPtrOutput {
	return o
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) ToSecurityContactPropertiesAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesAlertNotificationsPtrOutput {
	return o
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) Elem() SecurityContactPropertiesAlertNotificationsOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesAlertNotifications) SecurityContactPropertiesAlertNotifications {
		if v != nil {
			return *v
		}
		var ret SecurityContactPropertiesAlertNotifications
		return ret
	}).(SecurityContactPropertiesAlertNotificationsOutput)
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) MinimalSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesAlertNotifications) *string {
		if v == nil {
			return nil
		}
		return v.MinimalSeverity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityContactPropertiesAlertNotificationsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesAlertNotifications) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesNotificationsByRole struct {
	Roles []string `pulumi:"roles"`
	State *string  `pulumi:"state"`
}





type SecurityContactPropertiesNotificationsByRoleInput interface {
	pulumi.Input

	ToSecurityContactPropertiesNotificationsByRoleOutput() SecurityContactPropertiesNotificationsByRoleOutput
	ToSecurityContactPropertiesNotificationsByRoleOutputWithContext(context.Context) SecurityContactPropertiesNotificationsByRoleOutput
}

type SecurityContactPropertiesNotificationsByRoleArgs struct {
	Roles pulumi.StringArrayInput `pulumi:"roles"`
	State pulumi.StringPtrInput   `pulumi:"state"`
}

func (SecurityContactPropertiesNotificationsByRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesNotificationsByRole)(nil)).Elem()
}

func (i SecurityContactPropertiesNotificationsByRoleArgs) ToSecurityContactPropertiesNotificationsByRoleOutput() SecurityContactPropertiesNotificationsByRoleOutput {
	return i.ToSecurityContactPropertiesNotificationsByRoleOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesNotificationsByRoleArgs) ToSecurityContactPropertiesNotificationsByRoleOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesNotificationsByRoleOutput)
}

func (i SecurityContactPropertiesNotificationsByRoleArgs) ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput {
	return i.ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesNotificationsByRoleArgs) ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesNotificationsByRoleOutput).ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx)
}









type SecurityContactPropertiesNotificationsByRolePtrInput interface {
	pulumi.Input

	ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput
	ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput
}

type securityContactPropertiesNotificationsByRolePtrType SecurityContactPropertiesNotificationsByRoleArgs

func SecurityContactPropertiesNotificationsByRolePtr(v *SecurityContactPropertiesNotificationsByRoleArgs) SecurityContactPropertiesNotificationsByRolePtrInput {
	return (*securityContactPropertiesNotificationsByRolePtrType)(v)
}

func (*securityContactPropertiesNotificationsByRolePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesNotificationsByRole)(nil)).Elem()
}

func (i *securityContactPropertiesNotificationsByRolePtrType) ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput {
	return i.ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(context.Background())
}

func (i *securityContactPropertiesNotificationsByRolePtrType) ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesNotificationsByRolePtrOutput)
}

type SecurityContactPropertiesNotificationsByRoleOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesNotificationsByRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesNotificationsByRole)(nil)).Elem()
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) ToSecurityContactPropertiesNotificationsByRoleOutput() SecurityContactPropertiesNotificationsByRoleOutput {
	return o
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) ToSecurityContactPropertiesNotificationsByRoleOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRoleOutput {
	return o
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput {
	return o.ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(context.Background())
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityContactPropertiesNotificationsByRole) *SecurityContactPropertiesNotificationsByRole {
		return &v
	}).(SecurityContactPropertiesNotificationsByRolePtrOutput)
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityContactPropertiesNotificationsByRole) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o SecurityContactPropertiesNotificationsByRoleOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesNotificationsByRole) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesNotificationsByRolePtrOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesNotificationsByRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesNotificationsByRole)(nil)).Elem()
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) ToSecurityContactPropertiesNotificationsByRolePtrOutput() SecurityContactPropertiesNotificationsByRolePtrOutput {
	return o
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) ToSecurityContactPropertiesNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesNotificationsByRolePtrOutput {
	return o
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) Elem() SecurityContactPropertiesNotificationsByRoleOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesNotificationsByRole) SecurityContactPropertiesNotificationsByRole {
		if v != nil {
			return *v
		}
		var ret SecurityContactPropertiesNotificationsByRole
		return ret
	}).(SecurityContactPropertiesNotificationsByRoleOutput)
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesNotificationsByRole) []string {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(pulumi.StringArrayOutput)
}

func (o SecurityContactPropertiesNotificationsByRolePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesNotificationsByRole) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesResponseAlertNotifications struct {
	MinimalSeverity *string `pulumi:"minimalSeverity"`
	State           *string `pulumi:"state"`
}





type SecurityContactPropertiesResponseAlertNotificationsInput interface {
	pulumi.Input

	ToSecurityContactPropertiesResponseAlertNotificationsOutput() SecurityContactPropertiesResponseAlertNotificationsOutput
	ToSecurityContactPropertiesResponseAlertNotificationsOutputWithContext(context.Context) SecurityContactPropertiesResponseAlertNotificationsOutput
}

type SecurityContactPropertiesResponseAlertNotificationsArgs struct {
	MinimalSeverity pulumi.StringPtrInput `pulumi:"minimalSeverity"`
	State           pulumi.StringPtrInput `pulumi:"state"`
}

func (SecurityContactPropertiesResponseAlertNotificationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesResponseAlertNotifications)(nil)).Elem()
}

func (i SecurityContactPropertiesResponseAlertNotificationsArgs) ToSecurityContactPropertiesResponseAlertNotificationsOutput() SecurityContactPropertiesResponseAlertNotificationsOutput {
	return i.ToSecurityContactPropertiesResponseAlertNotificationsOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesResponseAlertNotificationsArgs) ToSecurityContactPropertiesResponseAlertNotificationsOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseAlertNotificationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesResponseAlertNotificationsOutput)
}

func (i SecurityContactPropertiesResponseAlertNotificationsArgs) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutput() SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return i.ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesResponseAlertNotificationsArgs) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesResponseAlertNotificationsOutput).ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(ctx)
}









type SecurityContactPropertiesResponseAlertNotificationsPtrInput interface {
	pulumi.Input

	ToSecurityContactPropertiesResponseAlertNotificationsPtrOutput() SecurityContactPropertiesResponseAlertNotificationsPtrOutput
	ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(context.Context) SecurityContactPropertiesResponseAlertNotificationsPtrOutput
}

type securityContactPropertiesResponseAlertNotificationsPtrType SecurityContactPropertiesResponseAlertNotificationsArgs

func SecurityContactPropertiesResponseAlertNotificationsPtr(v *SecurityContactPropertiesResponseAlertNotificationsArgs) SecurityContactPropertiesResponseAlertNotificationsPtrInput {
	return (*securityContactPropertiesResponseAlertNotificationsPtrType)(v)
}

func (*securityContactPropertiesResponseAlertNotificationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesResponseAlertNotifications)(nil)).Elem()
}

func (i *securityContactPropertiesResponseAlertNotificationsPtrType) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutput() SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return i.ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(context.Background())
}

func (i *securityContactPropertiesResponseAlertNotificationsPtrType) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesResponseAlertNotificationsPtrOutput)
}

type SecurityContactPropertiesResponseAlertNotificationsOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesResponseAlertNotificationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesResponseAlertNotifications)(nil)).Elem()
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) ToSecurityContactPropertiesResponseAlertNotificationsOutput() SecurityContactPropertiesResponseAlertNotificationsOutput {
	return o
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) ToSecurityContactPropertiesResponseAlertNotificationsOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseAlertNotificationsOutput {
	return o
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutput() SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return o.ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(context.Background())
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityContactPropertiesResponseAlertNotifications) *SecurityContactPropertiesResponseAlertNotifications {
		return &v
	}).(SecurityContactPropertiesResponseAlertNotificationsPtrOutput)
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) MinimalSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesResponseAlertNotifications) *string { return v.MinimalSeverity }).(pulumi.StringPtrOutput)
}

func (o SecurityContactPropertiesResponseAlertNotificationsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesResponseAlertNotifications) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesResponseAlertNotificationsPtrOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesResponseAlertNotificationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesResponseAlertNotifications)(nil)).Elem()
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutput() SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return o
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) ToSecurityContactPropertiesResponseAlertNotificationsPtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseAlertNotificationsPtrOutput {
	return o
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) Elem() SecurityContactPropertiesResponseAlertNotificationsOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseAlertNotifications) SecurityContactPropertiesResponseAlertNotifications {
		if v != nil {
			return *v
		}
		var ret SecurityContactPropertiesResponseAlertNotifications
		return ret
	}).(SecurityContactPropertiesResponseAlertNotificationsOutput)
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) MinimalSeverity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseAlertNotifications) *string {
		if v == nil {
			return nil
		}
		return v.MinimalSeverity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityContactPropertiesResponseAlertNotificationsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseAlertNotifications) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesResponseNotificationsByRole struct {
	Roles []string `pulumi:"roles"`
	State *string  `pulumi:"state"`
}





type SecurityContactPropertiesResponseNotificationsByRoleInput interface {
	pulumi.Input

	ToSecurityContactPropertiesResponseNotificationsByRoleOutput() SecurityContactPropertiesResponseNotificationsByRoleOutput
	ToSecurityContactPropertiesResponseNotificationsByRoleOutputWithContext(context.Context) SecurityContactPropertiesResponseNotificationsByRoleOutput
}

type SecurityContactPropertiesResponseNotificationsByRoleArgs struct {
	Roles pulumi.StringArrayInput `pulumi:"roles"`
	State pulumi.StringPtrInput   `pulumi:"state"`
}

func (SecurityContactPropertiesResponseNotificationsByRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesResponseNotificationsByRole)(nil)).Elem()
}

func (i SecurityContactPropertiesResponseNotificationsByRoleArgs) ToSecurityContactPropertiesResponseNotificationsByRoleOutput() SecurityContactPropertiesResponseNotificationsByRoleOutput {
	return i.ToSecurityContactPropertiesResponseNotificationsByRoleOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesResponseNotificationsByRoleArgs) ToSecurityContactPropertiesResponseNotificationsByRoleOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseNotificationsByRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesResponseNotificationsByRoleOutput)
}

func (i SecurityContactPropertiesResponseNotificationsByRoleArgs) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutput() SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return i.ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(context.Background())
}

func (i SecurityContactPropertiesResponseNotificationsByRoleArgs) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesResponseNotificationsByRoleOutput).ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(ctx)
}









type SecurityContactPropertiesResponseNotificationsByRolePtrInput interface {
	pulumi.Input

	ToSecurityContactPropertiesResponseNotificationsByRolePtrOutput() SecurityContactPropertiesResponseNotificationsByRolePtrOutput
	ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(context.Context) SecurityContactPropertiesResponseNotificationsByRolePtrOutput
}

type securityContactPropertiesResponseNotificationsByRolePtrType SecurityContactPropertiesResponseNotificationsByRoleArgs

func SecurityContactPropertiesResponseNotificationsByRolePtr(v *SecurityContactPropertiesResponseNotificationsByRoleArgs) SecurityContactPropertiesResponseNotificationsByRolePtrInput {
	return (*securityContactPropertiesResponseNotificationsByRolePtrType)(v)
}

func (*securityContactPropertiesResponseNotificationsByRolePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesResponseNotificationsByRole)(nil)).Elem()
}

func (i *securityContactPropertiesResponseNotificationsByRolePtrType) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutput() SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return i.ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(context.Background())
}

func (i *securityContactPropertiesResponseNotificationsByRolePtrType) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactPropertiesResponseNotificationsByRolePtrOutput)
}

type SecurityContactPropertiesResponseNotificationsByRoleOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesResponseNotificationsByRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityContactPropertiesResponseNotificationsByRole)(nil)).Elem()
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) ToSecurityContactPropertiesResponseNotificationsByRoleOutput() SecurityContactPropertiesResponseNotificationsByRoleOutput {
	return o
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) ToSecurityContactPropertiesResponseNotificationsByRoleOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseNotificationsByRoleOutput {
	return o
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutput() SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return o.ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(context.Background())
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityContactPropertiesResponseNotificationsByRole) *SecurityContactPropertiesResponseNotificationsByRole {
		return &v
	}).(SecurityContactPropertiesResponseNotificationsByRolePtrOutput)
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityContactPropertiesResponseNotificationsByRole) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o SecurityContactPropertiesResponseNotificationsByRoleOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityContactPropertiesResponseNotificationsByRole) *string { return v.State }).(pulumi.StringPtrOutput)
}

type SecurityContactPropertiesResponseNotificationsByRolePtrOutput struct{ *pulumi.OutputState }

func (SecurityContactPropertiesResponseNotificationsByRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContactPropertiesResponseNotificationsByRole)(nil)).Elem()
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutput() SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return o
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) ToSecurityContactPropertiesResponseNotificationsByRolePtrOutputWithContext(ctx context.Context) SecurityContactPropertiesResponseNotificationsByRolePtrOutput {
	return o
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) Elem() SecurityContactPropertiesResponseNotificationsByRoleOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseNotificationsByRole) SecurityContactPropertiesResponseNotificationsByRole {
		if v != nil {
			return *v
		}
		var ret SecurityContactPropertiesResponseNotificationsByRole
		return ret
	}).(SecurityContactPropertiesResponseNotificationsByRoleOutput)
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseNotificationsByRole) []string {
		if v == nil {
			return nil
		}
		return v.Roles
	}).(pulumi.StringArrayOutput)
}

func (o SecurityContactPropertiesResponseNotificationsByRolePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContactPropertiesResponseNotificationsByRole) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalProperties struct {
	ApplicationId *string `pulumi:"applicationId"`
	Secret        *string `pulumi:"secret"`
}





type ServicePrincipalPropertiesInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesOutput() ServicePrincipalPropertiesOutput
	ToServicePrincipalPropertiesOutputWithContext(context.Context) ServicePrincipalPropertiesOutput
}

type ServicePrincipalPropertiesArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	Secret        pulumi.StringPtrInput `pulumi:"secret"`
}

func (ServicePrincipalPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalProperties)(nil)).Elem()
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesOutput() ServicePrincipalPropertiesOutput {
	return i.ToServicePrincipalPropertiesOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesOutputWithContext(ctx context.Context) ServicePrincipalPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesOutput)
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return i.ToServicePrincipalPropertiesPtrOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesArgs) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesOutput).ToServicePrincipalPropertiesPtrOutputWithContext(ctx)
}









type ServicePrincipalPropertiesPtrInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput
	ToServicePrincipalPropertiesPtrOutputWithContext(context.Context) ServicePrincipalPropertiesPtrOutput
}

type servicePrincipalPropertiesPtrType ServicePrincipalPropertiesArgs

func ServicePrincipalPropertiesPtr(v *ServicePrincipalPropertiesArgs) ServicePrincipalPropertiesPtrInput {
	return (*servicePrincipalPropertiesPtrType)(v)
}

func (*servicePrincipalPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalProperties)(nil)).Elem()
}

func (i *servicePrincipalPropertiesPtrType) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return i.ToServicePrincipalPropertiesPtrOutputWithContext(context.Background())
}

func (i *servicePrincipalPropertiesPtrType) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesPtrOutput)
}

type ServicePrincipalPropertiesOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalProperties)(nil)).Elem()
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesOutput() ServicePrincipalPropertiesOutput {
	return o
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesOutputWithContext(ctx context.Context) ServicePrincipalPropertiesOutput {
	return o
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return o.ToServicePrincipalPropertiesPtrOutputWithContext(context.Background())
}

func (o ServicePrincipalPropertiesOutput) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePrincipalProperties) *ServicePrincipalProperties {
		return &v
	}).(ServicePrincipalPropertiesPtrOutput)
}

func (o ServicePrincipalPropertiesOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalProperties) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalProperties) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ServicePrincipalPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalProperties)(nil)).Elem()
}

func (o ServicePrincipalPropertiesPtrOutput) ToServicePrincipalPropertiesPtrOutput() ServicePrincipalPropertiesPtrOutput {
	return o
}

func (o ServicePrincipalPropertiesPtrOutput) ToServicePrincipalPropertiesPtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesPtrOutput {
	return o
}

func (o ServicePrincipalPropertiesPtrOutput) Elem() ServicePrincipalPropertiesOutput {
	return o.ApplyT(func(v *ServicePrincipalProperties) ServicePrincipalProperties {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalProperties
		return ret
	}).(ServicePrincipalPropertiesOutput)
}

func (o ServicePrincipalPropertiesPtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalProperties) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalPropertiesResponse struct {
	ApplicationId *string `pulumi:"applicationId"`
	Secret        *string `pulumi:"secret"`
}





type ServicePrincipalPropertiesResponseInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesResponseOutput() ServicePrincipalPropertiesResponseOutput
	ToServicePrincipalPropertiesResponseOutputWithContext(context.Context) ServicePrincipalPropertiesResponseOutput
}

type ServicePrincipalPropertiesResponseArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	Secret        pulumi.StringPtrInput `pulumi:"secret"`
}

func (ServicePrincipalPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (i ServicePrincipalPropertiesResponseArgs) ToServicePrincipalPropertiesResponseOutput() ServicePrincipalPropertiesResponseOutput {
	return i.ToServicePrincipalPropertiesResponseOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesResponseArgs) ToServicePrincipalPropertiesResponseOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesResponseOutput)
}

func (i ServicePrincipalPropertiesResponseArgs) ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput {
	return i.ToServicePrincipalPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ServicePrincipalPropertiesResponseArgs) ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesResponseOutput).ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx)
}









type ServicePrincipalPropertiesResponsePtrInput interface {
	pulumi.Input

	ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput
	ToServicePrincipalPropertiesResponsePtrOutputWithContext(context.Context) ServicePrincipalPropertiesResponsePtrOutput
}

type servicePrincipalPropertiesResponsePtrType ServicePrincipalPropertiesResponseArgs

func ServicePrincipalPropertiesResponsePtr(v *ServicePrincipalPropertiesResponseArgs) ServicePrincipalPropertiesResponsePtrInput {
	return (*servicePrincipalPropertiesResponsePtrType)(v)
}

func (*servicePrincipalPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (i *servicePrincipalPropertiesResponsePtrType) ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput {
	return i.ToServicePrincipalPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *servicePrincipalPropertiesResponsePtrType) ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalPropertiesResponsePtrOutput)
}

type ServicePrincipalPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponseOutput() ServicePrincipalPropertiesResponseOutput {
	return o
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponseOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponseOutput {
	return o
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput {
	return o.ToServicePrincipalPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ServicePrincipalPropertiesResponseOutput) ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePrincipalPropertiesResponse) *ServicePrincipalPropertiesResponse {
		return &v
	}).(ServicePrincipalPropertiesResponsePtrOutput)
}

func (o ServicePrincipalPropertiesResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalPropertiesResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesResponseOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalPropertiesResponse) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ServicePrincipalPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalPropertiesResponse)(nil)).Elem()
}

func (o ServicePrincipalPropertiesResponsePtrOutput) ToServicePrincipalPropertiesResponsePtrOutput() ServicePrincipalPropertiesResponsePtrOutput {
	return o
}

func (o ServicePrincipalPropertiesResponsePtrOutput) ToServicePrincipalPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalPropertiesResponsePtrOutput {
	return o
}

func (o ServicePrincipalPropertiesResponsePtrOutput) Elem() ServicePrincipalPropertiesResponseOutput {
	return o.ApplyT(func(v *ServicePrincipalPropertiesResponse) ServicePrincipalPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalPropertiesResponse
		return ret
	}).(ServicePrincipalPropertiesResponseOutput)
}

func (o ServicePrincipalPropertiesResponsePtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalPropertiesResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type StandardComponentProperties struct {
	Key *string `pulumi:"key"`
}





type StandardComponentPropertiesInput interface {
	pulumi.Input

	ToStandardComponentPropertiesOutput() StandardComponentPropertiesOutput
	ToStandardComponentPropertiesOutputWithContext(context.Context) StandardComponentPropertiesOutput
}

type StandardComponentPropertiesArgs struct {
	Key pulumi.StringPtrInput `pulumi:"key"`
}

func (StandardComponentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardComponentProperties)(nil)).Elem()
}

func (i StandardComponentPropertiesArgs) ToStandardComponentPropertiesOutput() StandardComponentPropertiesOutput {
	return i.ToStandardComponentPropertiesOutputWithContext(context.Background())
}

func (i StandardComponentPropertiesArgs) ToStandardComponentPropertiesOutputWithContext(ctx context.Context) StandardComponentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardComponentPropertiesOutput)
}





type StandardComponentPropertiesArrayInput interface {
	pulumi.Input

	ToStandardComponentPropertiesArrayOutput() StandardComponentPropertiesArrayOutput
	ToStandardComponentPropertiesArrayOutputWithContext(context.Context) StandardComponentPropertiesArrayOutput
}

type StandardComponentPropertiesArray []StandardComponentPropertiesInput

func (StandardComponentPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardComponentProperties)(nil)).Elem()
}

func (i StandardComponentPropertiesArray) ToStandardComponentPropertiesArrayOutput() StandardComponentPropertiesArrayOutput {
	return i.ToStandardComponentPropertiesArrayOutputWithContext(context.Background())
}

func (i StandardComponentPropertiesArray) ToStandardComponentPropertiesArrayOutputWithContext(ctx context.Context) StandardComponentPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardComponentPropertiesArrayOutput)
}

type StandardComponentPropertiesOutput struct{ *pulumi.OutputState }

func (StandardComponentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardComponentProperties)(nil)).Elem()
}

func (o StandardComponentPropertiesOutput) ToStandardComponentPropertiesOutput() StandardComponentPropertiesOutput {
	return o
}

func (o StandardComponentPropertiesOutput) ToStandardComponentPropertiesOutputWithContext(ctx context.Context) StandardComponentPropertiesOutput {
	return o
}

func (o StandardComponentPropertiesOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StandardComponentProperties) *string { return v.Key }).(pulumi.StringPtrOutput)
}

type StandardComponentPropertiesArrayOutput struct{ *pulumi.OutputState }

func (StandardComponentPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardComponentProperties)(nil)).Elem()
}

func (o StandardComponentPropertiesArrayOutput) ToStandardComponentPropertiesArrayOutput() StandardComponentPropertiesArrayOutput {
	return o
}

func (o StandardComponentPropertiesArrayOutput) ToStandardComponentPropertiesArrayOutputWithContext(ctx context.Context) StandardComponentPropertiesArrayOutput {
	return o
}

func (o StandardComponentPropertiesArrayOutput) Index(i pulumi.IntInput) StandardComponentPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StandardComponentProperties {
		return vs[0].([]StandardComponentProperties)[vs[1].(int)]
	}).(StandardComponentPropertiesOutput)
}

type StandardComponentPropertiesResponse struct {
	Key *string `pulumi:"key"`
}





type StandardComponentPropertiesResponseInput interface {
	pulumi.Input

	ToStandardComponentPropertiesResponseOutput() StandardComponentPropertiesResponseOutput
	ToStandardComponentPropertiesResponseOutputWithContext(context.Context) StandardComponentPropertiesResponseOutput
}

type StandardComponentPropertiesResponseArgs struct {
	Key pulumi.StringPtrInput `pulumi:"key"`
}

func (StandardComponentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardComponentPropertiesResponse)(nil)).Elem()
}

func (i StandardComponentPropertiesResponseArgs) ToStandardComponentPropertiesResponseOutput() StandardComponentPropertiesResponseOutput {
	return i.ToStandardComponentPropertiesResponseOutputWithContext(context.Background())
}

func (i StandardComponentPropertiesResponseArgs) ToStandardComponentPropertiesResponseOutputWithContext(ctx context.Context) StandardComponentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardComponentPropertiesResponseOutput)
}





type StandardComponentPropertiesResponseArrayInput interface {
	pulumi.Input

	ToStandardComponentPropertiesResponseArrayOutput() StandardComponentPropertiesResponseArrayOutput
	ToStandardComponentPropertiesResponseArrayOutputWithContext(context.Context) StandardComponentPropertiesResponseArrayOutput
}

type StandardComponentPropertiesResponseArray []StandardComponentPropertiesResponseInput

func (StandardComponentPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardComponentPropertiesResponse)(nil)).Elem()
}

func (i StandardComponentPropertiesResponseArray) ToStandardComponentPropertiesResponseArrayOutput() StandardComponentPropertiesResponseArrayOutput {
	return i.ToStandardComponentPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i StandardComponentPropertiesResponseArray) ToStandardComponentPropertiesResponseArrayOutputWithContext(ctx context.Context) StandardComponentPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardComponentPropertiesResponseArrayOutput)
}

type StandardComponentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StandardComponentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardComponentPropertiesResponse)(nil)).Elem()
}

func (o StandardComponentPropertiesResponseOutput) ToStandardComponentPropertiesResponseOutput() StandardComponentPropertiesResponseOutput {
	return o
}

func (o StandardComponentPropertiesResponseOutput) ToStandardComponentPropertiesResponseOutputWithContext(ctx context.Context) StandardComponentPropertiesResponseOutput {
	return o
}

func (o StandardComponentPropertiesResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StandardComponentPropertiesResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

type StandardComponentPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (StandardComponentPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardComponentPropertiesResponse)(nil)).Elem()
}

func (o StandardComponentPropertiesResponseArrayOutput) ToStandardComponentPropertiesResponseArrayOutput() StandardComponentPropertiesResponseArrayOutput {
	return o
}

func (o StandardComponentPropertiesResponseArrayOutput) ToStandardComponentPropertiesResponseArrayOutputWithContext(ctx context.Context) StandardComponentPropertiesResponseArrayOutput {
	return o
}

func (o StandardComponentPropertiesResponseArrayOutput) Index(i pulumi.IntInput) StandardComponentPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StandardComponentPropertiesResponse {
		return vs[0].([]StandardComponentPropertiesResponse)[vs[1].(int)]
	}).(StandardComponentPropertiesResponseOutput)
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





type SuppressionAlertsScopeResponseInput interface {
	pulumi.Input

	ToSuppressionAlertsScopeResponseOutput() SuppressionAlertsScopeResponseOutput
	ToSuppressionAlertsScopeResponseOutputWithContext(context.Context) SuppressionAlertsScopeResponseOutput
}

type SuppressionAlertsScopeResponseArgs struct {
	AllOf ScopeElementResponseArrayInput `pulumi:"allOf"`
}

func (SuppressionAlertsScopeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionAlertsScopeResponse)(nil)).Elem()
}

func (i SuppressionAlertsScopeResponseArgs) ToSuppressionAlertsScopeResponseOutput() SuppressionAlertsScopeResponseOutput {
	return i.ToSuppressionAlertsScopeResponseOutputWithContext(context.Background())
}

func (i SuppressionAlertsScopeResponseArgs) ToSuppressionAlertsScopeResponseOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeResponseOutput)
}

func (i SuppressionAlertsScopeResponseArgs) ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput {
	return i.ToSuppressionAlertsScopeResponsePtrOutputWithContext(context.Background())
}

func (i SuppressionAlertsScopeResponseArgs) ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeResponseOutput).ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx)
}









type SuppressionAlertsScopeResponsePtrInput interface {
	pulumi.Input

	ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput
	ToSuppressionAlertsScopeResponsePtrOutputWithContext(context.Context) SuppressionAlertsScopeResponsePtrOutput
}

type suppressionAlertsScopeResponsePtrType SuppressionAlertsScopeResponseArgs

func SuppressionAlertsScopeResponsePtr(v *SuppressionAlertsScopeResponseArgs) SuppressionAlertsScopeResponsePtrInput {
	return (*suppressionAlertsScopeResponsePtrType)(v)
}

func (*suppressionAlertsScopeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionAlertsScopeResponse)(nil)).Elem()
}

func (i *suppressionAlertsScopeResponsePtrType) ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput {
	return i.ToSuppressionAlertsScopeResponsePtrOutputWithContext(context.Background())
}

func (i *suppressionAlertsScopeResponsePtrType) ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionAlertsScopeResponsePtrOutput)
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

func (o SuppressionAlertsScopeResponseOutput) ToSuppressionAlertsScopeResponsePtrOutput() SuppressionAlertsScopeResponsePtrOutput {
	return o.ToSuppressionAlertsScopeResponsePtrOutputWithContext(context.Background())
}

func (o SuppressionAlertsScopeResponseOutput) ToSuppressionAlertsScopeResponsePtrOutputWithContext(ctx context.Context) SuppressionAlertsScopeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SuppressionAlertsScopeResponse) *SuppressionAlertsScopeResponse {
		return &v
	}).(SuppressionAlertsScopeResponsePtrOutput)
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

type ThresholdCustomAlertRule struct {
	IsEnabled    bool   `pulumi:"isEnabled"`
	MaxThreshold int    `pulumi:"maxThreshold"`
	MinThreshold int    `pulumi:"minThreshold"`
	RuleType     string `pulumi:"ruleType"`
}





type ThresholdCustomAlertRuleInput interface {
	pulumi.Input

	ToThresholdCustomAlertRuleOutput() ThresholdCustomAlertRuleOutput
	ToThresholdCustomAlertRuleOutputWithContext(context.Context) ThresholdCustomAlertRuleOutput
}

type ThresholdCustomAlertRuleArgs struct {
	IsEnabled    pulumi.BoolInput   `pulumi:"isEnabled"`
	MaxThreshold pulumi.IntInput    `pulumi:"maxThreshold"`
	MinThreshold pulumi.IntInput    `pulumi:"minThreshold"`
	RuleType     pulumi.StringInput `pulumi:"ruleType"`
}

func (ThresholdCustomAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdCustomAlertRule)(nil)).Elem()
}

func (i ThresholdCustomAlertRuleArgs) ToThresholdCustomAlertRuleOutput() ThresholdCustomAlertRuleOutput {
	return i.ToThresholdCustomAlertRuleOutputWithContext(context.Background())
}

func (i ThresholdCustomAlertRuleArgs) ToThresholdCustomAlertRuleOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdCustomAlertRuleOutput)
}





type ThresholdCustomAlertRuleArrayInput interface {
	pulumi.Input

	ToThresholdCustomAlertRuleArrayOutput() ThresholdCustomAlertRuleArrayOutput
	ToThresholdCustomAlertRuleArrayOutputWithContext(context.Context) ThresholdCustomAlertRuleArrayOutput
}

type ThresholdCustomAlertRuleArray []ThresholdCustomAlertRuleInput

func (ThresholdCustomAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThresholdCustomAlertRule)(nil)).Elem()
}

func (i ThresholdCustomAlertRuleArray) ToThresholdCustomAlertRuleArrayOutput() ThresholdCustomAlertRuleArrayOutput {
	return i.ToThresholdCustomAlertRuleArrayOutputWithContext(context.Background())
}

func (i ThresholdCustomAlertRuleArray) ToThresholdCustomAlertRuleArrayOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdCustomAlertRuleArrayOutput)
}

type ThresholdCustomAlertRuleOutput struct{ *pulumi.OutputState }

func (ThresholdCustomAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdCustomAlertRule)(nil)).Elem()
}

func (o ThresholdCustomAlertRuleOutput) ToThresholdCustomAlertRuleOutput() ThresholdCustomAlertRuleOutput {
	return o
}

func (o ThresholdCustomAlertRuleOutput) ToThresholdCustomAlertRuleOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleOutput {
	return o
}

func (o ThresholdCustomAlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRule) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o ThresholdCustomAlertRuleOutput) MaxThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRule) int { return v.MaxThreshold }).(pulumi.IntOutput)
}

func (o ThresholdCustomAlertRuleOutput) MinThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRule) int { return v.MinThreshold }).(pulumi.IntOutput)
}

func (o ThresholdCustomAlertRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type ThresholdCustomAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (ThresholdCustomAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThresholdCustomAlertRule)(nil)).Elem()
}

func (o ThresholdCustomAlertRuleArrayOutput) ToThresholdCustomAlertRuleArrayOutput() ThresholdCustomAlertRuleArrayOutput {
	return o
}

func (o ThresholdCustomAlertRuleArrayOutput) ToThresholdCustomAlertRuleArrayOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleArrayOutput {
	return o
}

func (o ThresholdCustomAlertRuleArrayOutput) Index(i pulumi.IntInput) ThresholdCustomAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThresholdCustomAlertRule {
		return vs[0].([]ThresholdCustomAlertRule)[vs[1].(int)]
	}).(ThresholdCustomAlertRuleOutput)
}

type ThresholdCustomAlertRuleResponse struct {
	Description  string `pulumi:"description"`
	DisplayName  string `pulumi:"displayName"`
	IsEnabled    bool   `pulumi:"isEnabled"`
	MaxThreshold int    `pulumi:"maxThreshold"`
	MinThreshold int    `pulumi:"minThreshold"`
	RuleType     string `pulumi:"ruleType"`
}





type ThresholdCustomAlertRuleResponseInput interface {
	pulumi.Input

	ToThresholdCustomAlertRuleResponseOutput() ThresholdCustomAlertRuleResponseOutput
	ToThresholdCustomAlertRuleResponseOutputWithContext(context.Context) ThresholdCustomAlertRuleResponseOutput
}

type ThresholdCustomAlertRuleResponseArgs struct {
	Description  pulumi.StringInput `pulumi:"description"`
	DisplayName  pulumi.StringInput `pulumi:"displayName"`
	IsEnabled    pulumi.BoolInput   `pulumi:"isEnabled"`
	MaxThreshold pulumi.IntInput    `pulumi:"maxThreshold"`
	MinThreshold pulumi.IntInput    `pulumi:"minThreshold"`
	RuleType     pulumi.StringInput `pulumi:"ruleType"`
}

func (ThresholdCustomAlertRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdCustomAlertRuleResponse)(nil)).Elem()
}

func (i ThresholdCustomAlertRuleResponseArgs) ToThresholdCustomAlertRuleResponseOutput() ThresholdCustomAlertRuleResponseOutput {
	return i.ToThresholdCustomAlertRuleResponseOutputWithContext(context.Background())
}

func (i ThresholdCustomAlertRuleResponseArgs) ToThresholdCustomAlertRuleResponseOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdCustomAlertRuleResponseOutput)
}





type ThresholdCustomAlertRuleResponseArrayInput interface {
	pulumi.Input

	ToThresholdCustomAlertRuleResponseArrayOutput() ThresholdCustomAlertRuleResponseArrayOutput
	ToThresholdCustomAlertRuleResponseArrayOutputWithContext(context.Context) ThresholdCustomAlertRuleResponseArrayOutput
}

type ThresholdCustomAlertRuleResponseArray []ThresholdCustomAlertRuleResponseInput

func (ThresholdCustomAlertRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThresholdCustomAlertRuleResponse)(nil)).Elem()
}

func (i ThresholdCustomAlertRuleResponseArray) ToThresholdCustomAlertRuleResponseArrayOutput() ThresholdCustomAlertRuleResponseArrayOutput {
	return i.ToThresholdCustomAlertRuleResponseArrayOutputWithContext(context.Background())
}

func (i ThresholdCustomAlertRuleResponseArray) ToThresholdCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdCustomAlertRuleResponseArrayOutput)
}

type ThresholdCustomAlertRuleResponseOutput struct{ *pulumi.OutputState }

func (ThresholdCustomAlertRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdCustomAlertRuleResponse)(nil)).Elem()
}

func (o ThresholdCustomAlertRuleResponseOutput) ToThresholdCustomAlertRuleResponseOutput() ThresholdCustomAlertRuleResponseOutput {
	return o
}

func (o ThresholdCustomAlertRuleResponseOutput) ToThresholdCustomAlertRuleResponseOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleResponseOutput {
	return o
}

func (o ThresholdCustomAlertRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) MaxThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) int { return v.MaxThreshold }).(pulumi.IntOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) MinThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) int { return v.MinThreshold }).(pulumi.IntOutput)
}

func (o ThresholdCustomAlertRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdCustomAlertRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

type ThresholdCustomAlertRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ThresholdCustomAlertRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThresholdCustomAlertRuleResponse)(nil)).Elem()
}

func (o ThresholdCustomAlertRuleResponseArrayOutput) ToThresholdCustomAlertRuleResponseArrayOutput() ThresholdCustomAlertRuleResponseArrayOutput {
	return o
}

func (o ThresholdCustomAlertRuleResponseArrayOutput) ToThresholdCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) ThresholdCustomAlertRuleResponseArrayOutput {
	return o
}

func (o ThresholdCustomAlertRuleResponseArrayOutput) Index(i pulumi.IntInput) ThresholdCustomAlertRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThresholdCustomAlertRuleResponse {
		return vs[0].([]ThresholdCustomAlertRuleResponse)[vs[1].(int)]
	}).(ThresholdCustomAlertRuleResponseOutput)
}

type TimeWindowCustomAlertRule struct {
	IsEnabled      bool   `pulumi:"isEnabled"`
	MaxThreshold   int    `pulumi:"maxThreshold"`
	MinThreshold   int    `pulumi:"minThreshold"`
	RuleType       string `pulumi:"ruleType"`
	TimeWindowSize string `pulumi:"timeWindowSize"`
}





type TimeWindowCustomAlertRuleInput interface {
	pulumi.Input

	ToTimeWindowCustomAlertRuleOutput() TimeWindowCustomAlertRuleOutput
	ToTimeWindowCustomAlertRuleOutputWithContext(context.Context) TimeWindowCustomAlertRuleOutput
}

type TimeWindowCustomAlertRuleArgs struct {
	IsEnabled      pulumi.BoolInput   `pulumi:"isEnabled"`
	MaxThreshold   pulumi.IntInput    `pulumi:"maxThreshold"`
	MinThreshold   pulumi.IntInput    `pulumi:"minThreshold"`
	RuleType       pulumi.StringInput `pulumi:"ruleType"`
	TimeWindowSize pulumi.StringInput `pulumi:"timeWindowSize"`
}

func (TimeWindowCustomAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowCustomAlertRule)(nil)).Elem()
}

func (i TimeWindowCustomAlertRuleArgs) ToTimeWindowCustomAlertRuleOutput() TimeWindowCustomAlertRuleOutput {
	return i.ToTimeWindowCustomAlertRuleOutputWithContext(context.Background())
}

func (i TimeWindowCustomAlertRuleArgs) ToTimeWindowCustomAlertRuleOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowCustomAlertRuleOutput)
}





type TimeWindowCustomAlertRuleArrayInput interface {
	pulumi.Input

	ToTimeWindowCustomAlertRuleArrayOutput() TimeWindowCustomAlertRuleArrayOutput
	ToTimeWindowCustomAlertRuleArrayOutputWithContext(context.Context) TimeWindowCustomAlertRuleArrayOutput
}

type TimeWindowCustomAlertRuleArray []TimeWindowCustomAlertRuleInput

func (TimeWindowCustomAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeWindowCustomAlertRule)(nil)).Elem()
}

func (i TimeWindowCustomAlertRuleArray) ToTimeWindowCustomAlertRuleArrayOutput() TimeWindowCustomAlertRuleArrayOutput {
	return i.ToTimeWindowCustomAlertRuleArrayOutputWithContext(context.Background())
}

func (i TimeWindowCustomAlertRuleArray) ToTimeWindowCustomAlertRuleArrayOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowCustomAlertRuleArrayOutput)
}

type TimeWindowCustomAlertRuleOutput struct{ *pulumi.OutputState }

func (TimeWindowCustomAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowCustomAlertRule)(nil)).Elem()
}

func (o TimeWindowCustomAlertRuleOutput) ToTimeWindowCustomAlertRuleOutput() TimeWindowCustomAlertRuleOutput {
	return o
}

func (o TimeWindowCustomAlertRuleOutput) ToTimeWindowCustomAlertRuleOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleOutput {
	return o
}

func (o TimeWindowCustomAlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o TimeWindowCustomAlertRuleOutput) MaxThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) int { return v.MaxThreshold }).(pulumi.IntOutput)
}

func (o TimeWindowCustomAlertRuleOutput) MinThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) int { return v.MinThreshold }).(pulumi.IntOutput)
}

func (o TimeWindowCustomAlertRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o TimeWindowCustomAlertRuleOutput) TimeWindowSize() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRule) string { return v.TimeWindowSize }).(pulumi.StringOutput)
}

type TimeWindowCustomAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (TimeWindowCustomAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeWindowCustomAlertRule)(nil)).Elem()
}

func (o TimeWindowCustomAlertRuleArrayOutput) ToTimeWindowCustomAlertRuleArrayOutput() TimeWindowCustomAlertRuleArrayOutput {
	return o
}

func (o TimeWindowCustomAlertRuleArrayOutput) ToTimeWindowCustomAlertRuleArrayOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleArrayOutput {
	return o
}

func (o TimeWindowCustomAlertRuleArrayOutput) Index(i pulumi.IntInput) TimeWindowCustomAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeWindowCustomAlertRule {
		return vs[0].([]TimeWindowCustomAlertRule)[vs[1].(int)]
	}).(TimeWindowCustomAlertRuleOutput)
}

type TimeWindowCustomAlertRuleResponse struct {
	Description    string `pulumi:"description"`
	DisplayName    string `pulumi:"displayName"`
	IsEnabled      bool   `pulumi:"isEnabled"`
	MaxThreshold   int    `pulumi:"maxThreshold"`
	MinThreshold   int    `pulumi:"minThreshold"`
	RuleType       string `pulumi:"ruleType"`
	TimeWindowSize string `pulumi:"timeWindowSize"`
}





type TimeWindowCustomAlertRuleResponseInput interface {
	pulumi.Input

	ToTimeWindowCustomAlertRuleResponseOutput() TimeWindowCustomAlertRuleResponseOutput
	ToTimeWindowCustomAlertRuleResponseOutputWithContext(context.Context) TimeWindowCustomAlertRuleResponseOutput
}

type TimeWindowCustomAlertRuleResponseArgs struct {
	Description    pulumi.StringInput `pulumi:"description"`
	DisplayName    pulumi.StringInput `pulumi:"displayName"`
	IsEnabled      pulumi.BoolInput   `pulumi:"isEnabled"`
	MaxThreshold   pulumi.IntInput    `pulumi:"maxThreshold"`
	MinThreshold   pulumi.IntInput    `pulumi:"minThreshold"`
	RuleType       pulumi.StringInput `pulumi:"ruleType"`
	TimeWindowSize pulumi.StringInput `pulumi:"timeWindowSize"`
}

func (TimeWindowCustomAlertRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowCustomAlertRuleResponse)(nil)).Elem()
}

func (i TimeWindowCustomAlertRuleResponseArgs) ToTimeWindowCustomAlertRuleResponseOutput() TimeWindowCustomAlertRuleResponseOutput {
	return i.ToTimeWindowCustomAlertRuleResponseOutputWithContext(context.Background())
}

func (i TimeWindowCustomAlertRuleResponseArgs) ToTimeWindowCustomAlertRuleResponseOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowCustomAlertRuleResponseOutput)
}





type TimeWindowCustomAlertRuleResponseArrayInput interface {
	pulumi.Input

	ToTimeWindowCustomAlertRuleResponseArrayOutput() TimeWindowCustomAlertRuleResponseArrayOutput
	ToTimeWindowCustomAlertRuleResponseArrayOutputWithContext(context.Context) TimeWindowCustomAlertRuleResponseArrayOutput
}

type TimeWindowCustomAlertRuleResponseArray []TimeWindowCustomAlertRuleResponseInput

func (TimeWindowCustomAlertRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeWindowCustomAlertRuleResponse)(nil)).Elem()
}

func (i TimeWindowCustomAlertRuleResponseArray) ToTimeWindowCustomAlertRuleResponseArrayOutput() TimeWindowCustomAlertRuleResponseArrayOutput {
	return i.ToTimeWindowCustomAlertRuleResponseArrayOutputWithContext(context.Background())
}

func (i TimeWindowCustomAlertRuleResponseArray) ToTimeWindowCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowCustomAlertRuleResponseArrayOutput)
}

type TimeWindowCustomAlertRuleResponseOutput struct{ *pulumi.OutputState }

func (TimeWindowCustomAlertRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowCustomAlertRuleResponse)(nil)).Elem()
}

func (o TimeWindowCustomAlertRuleResponseOutput) ToTimeWindowCustomAlertRuleResponseOutput() TimeWindowCustomAlertRuleResponseOutput {
	return o
}

func (o TimeWindowCustomAlertRuleResponseOutput) ToTimeWindowCustomAlertRuleResponseOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleResponseOutput {
	return o
}

func (o TimeWindowCustomAlertRuleResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) MaxThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) int { return v.MaxThreshold }).(pulumi.IntOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) MinThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) int { return v.MinThreshold }).(pulumi.IntOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o TimeWindowCustomAlertRuleResponseOutput) TimeWindowSize() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowCustomAlertRuleResponse) string { return v.TimeWindowSize }).(pulumi.StringOutput)
}

type TimeWindowCustomAlertRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (TimeWindowCustomAlertRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimeWindowCustomAlertRuleResponse)(nil)).Elem()
}

func (o TimeWindowCustomAlertRuleResponseArrayOutput) ToTimeWindowCustomAlertRuleResponseArrayOutput() TimeWindowCustomAlertRuleResponseArrayOutput {
	return o
}

func (o TimeWindowCustomAlertRuleResponseArrayOutput) ToTimeWindowCustomAlertRuleResponseArrayOutputWithContext(ctx context.Context) TimeWindowCustomAlertRuleResponseArrayOutput {
	return o
}

func (o TimeWindowCustomAlertRuleResponseArrayOutput) Index(i pulumi.IntInput) TimeWindowCustomAlertRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimeWindowCustomAlertRuleResponse {
		return vs[0].([]TimeWindowCustomAlertRuleResponse)[vs[1].(int)]
	}).(TimeWindowCustomAlertRuleResponseOutput)
}

type UserDefinedResourcesProperties struct {
	Query              string   `pulumi:"query"`
	QuerySubscriptions []string `pulumi:"querySubscriptions"`
}





type UserDefinedResourcesPropertiesInput interface {
	pulumi.Input

	ToUserDefinedResourcesPropertiesOutput() UserDefinedResourcesPropertiesOutput
	ToUserDefinedResourcesPropertiesOutputWithContext(context.Context) UserDefinedResourcesPropertiesOutput
}

type UserDefinedResourcesPropertiesArgs struct {
	Query              pulumi.StringInput      `pulumi:"query"`
	QuerySubscriptions pulumi.StringArrayInput `pulumi:"querySubscriptions"`
}

func (UserDefinedResourcesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDefinedResourcesProperties)(nil)).Elem()
}

func (i UserDefinedResourcesPropertiesArgs) ToUserDefinedResourcesPropertiesOutput() UserDefinedResourcesPropertiesOutput {
	return i.ToUserDefinedResourcesPropertiesOutputWithContext(context.Background())
}

func (i UserDefinedResourcesPropertiesArgs) ToUserDefinedResourcesPropertiesOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesOutput)
}

func (i UserDefinedResourcesPropertiesArgs) ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput {
	return i.ToUserDefinedResourcesPropertiesPtrOutputWithContext(context.Background())
}

func (i UserDefinedResourcesPropertiesArgs) ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesOutput).ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx)
}









type UserDefinedResourcesPropertiesPtrInput interface {
	pulumi.Input

	ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput
	ToUserDefinedResourcesPropertiesPtrOutputWithContext(context.Context) UserDefinedResourcesPropertiesPtrOutput
}

type userDefinedResourcesPropertiesPtrType UserDefinedResourcesPropertiesArgs

func UserDefinedResourcesPropertiesPtr(v *UserDefinedResourcesPropertiesArgs) UserDefinedResourcesPropertiesPtrInput {
	return (*userDefinedResourcesPropertiesPtrType)(v)
}

func (*userDefinedResourcesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedResourcesProperties)(nil)).Elem()
}

func (i *userDefinedResourcesPropertiesPtrType) ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput {
	return i.ToUserDefinedResourcesPropertiesPtrOutputWithContext(context.Background())
}

func (i *userDefinedResourcesPropertiesPtrType) ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesPtrOutput)
}

type UserDefinedResourcesPropertiesOutput struct{ *pulumi.OutputState }

func (UserDefinedResourcesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDefinedResourcesProperties)(nil)).Elem()
}

func (o UserDefinedResourcesPropertiesOutput) ToUserDefinedResourcesPropertiesOutput() UserDefinedResourcesPropertiesOutput {
	return o
}

func (o UserDefinedResourcesPropertiesOutput) ToUserDefinedResourcesPropertiesOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesOutput {
	return o
}

func (o UserDefinedResourcesPropertiesOutput) ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput {
	return o.ToUserDefinedResourcesPropertiesPtrOutputWithContext(context.Background())
}

func (o UserDefinedResourcesPropertiesOutput) ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserDefinedResourcesProperties) *UserDefinedResourcesProperties {
		return &v
	}).(UserDefinedResourcesPropertiesPtrOutput)
}

func (o UserDefinedResourcesPropertiesOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v UserDefinedResourcesProperties) string { return v.Query }).(pulumi.StringOutput)
}

func (o UserDefinedResourcesPropertiesOutput) QuerySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UserDefinedResourcesProperties) []string { return v.QuerySubscriptions }).(pulumi.StringArrayOutput)
}

type UserDefinedResourcesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (UserDefinedResourcesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedResourcesProperties)(nil)).Elem()
}

func (o UserDefinedResourcesPropertiesPtrOutput) ToUserDefinedResourcesPropertiesPtrOutput() UserDefinedResourcesPropertiesPtrOutput {
	return o
}

func (o UserDefinedResourcesPropertiesPtrOutput) ToUserDefinedResourcesPropertiesPtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesPtrOutput {
	return o
}

func (o UserDefinedResourcesPropertiesPtrOutput) Elem() UserDefinedResourcesPropertiesOutput {
	return o.ApplyT(func(v *UserDefinedResourcesProperties) UserDefinedResourcesProperties {
		if v != nil {
			return *v
		}
		var ret UserDefinedResourcesProperties
		return ret
	}).(UserDefinedResourcesPropertiesOutput)
}

func (o UserDefinedResourcesPropertiesPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDefinedResourcesProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Query
	}).(pulumi.StringPtrOutput)
}

func (o UserDefinedResourcesPropertiesPtrOutput) QuerySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserDefinedResourcesProperties) []string {
		if v == nil {
			return nil
		}
		return v.QuerySubscriptions
	}).(pulumi.StringArrayOutput)
}

type UserDefinedResourcesPropertiesResponse struct {
	Query              string   `pulumi:"query"`
	QuerySubscriptions []string `pulumi:"querySubscriptions"`
}





type UserDefinedResourcesPropertiesResponseInput interface {
	pulumi.Input

	ToUserDefinedResourcesPropertiesResponseOutput() UserDefinedResourcesPropertiesResponseOutput
	ToUserDefinedResourcesPropertiesResponseOutputWithContext(context.Context) UserDefinedResourcesPropertiesResponseOutput
}

type UserDefinedResourcesPropertiesResponseArgs struct {
	Query              pulumi.StringInput      `pulumi:"query"`
	QuerySubscriptions pulumi.StringArrayInput `pulumi:"querySubscriptions"`
}

func (UserDefinedResourcesPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDefinedResourcesPropertiesResponse)(nil)).Elem()
}

func (i UserDefinedResourcesPropertiesResponseArgs) ToUserDefinedResourcesPropertiesResponseOutput() UserDefinedResourcesPropertiesResponseOutput {
	return i.ToUserDefinedResourcesPropertiesResponseOutputWithContext(context.Background())
}

func (i UserDefinedResourcesPropertiesResponseArgs) ToUserDefinedResourcesPropertiesResponseOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesResponseOutput)
}

func (i UserDefinedResourcesPropertiesResponseArgs) ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput {
	return i.ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i UserDefinedResourcesPropertiesResponseArgs) ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesResponseOutput).ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx)
}









type UserDefinedResourcesPropertiesResponsePtrInput interface {
	pulumi.Input

	ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput
	ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(context.Context) UserDefinedResourcesPropertiesResponsePtrOutput
}

type userDefinedResourcesPropertiesResponsePtrType UserDefinedResourcesPropertiesResponseArgs

func UserDefinedResourcesPropertiesResponsePtr(v *UserDefinedResourcesPropertiesResponseArgs) UserDefinedResourcesPropertiesResponsePtrInput {
	return (*userDefinedResourcesPropertiesResponsePtrType)(v)
}

func (*userDefinedResourcesPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedResourcesPropertiesResponse)(nil)).Elem()
}

func (i *userDefinedResourcesPropertiesResponsePtrType) ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput {
	return i.ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *userDefinedResourcesPropertiesResponsePtrType) ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDefinedResourcesPropertiesResponsePtrOutput)
}

type UserDefinedResourcesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserDefinedResourcesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserDefinedResourcesPropertiesResponse)(nil)).Elem()
}

func (o UserDefinedResourcesPropertiesResponseOutput) ToUserDefinedResourcesPropertiesResponseOutput() UserDefinedResourcesPropertiesResponseOutput {
	return o
}

func (o UserDefinedResourcesPropertiesResponseOutput) ToUserDefinedResourcesPropertiesResponseOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponseOutput {
	return o
}

func (o UserDefinedResourcesPropertiesResponseOutput) ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput {
	return o.ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o UserDefinedResourcesPropertiesResponseOutput) ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserDefinedResourcesPropertiesResponse) *UserDefinedResourcesPropertiesResponse {
		return &v
	}).(UserDefinedResourcesPropertiesResponsePtrOutput)
}

func (o UserDefinedResourcesPropertiesResponseOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v UserDefinedResourcesPropertiesResponse) string { return v.Query }).(pulumi.StringOutput)
}

func (o UserDefinedResourcesPropertiesResponseOutput) QuerySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UserDefinedResourcesPropertiesResponse) []string { return v.QuerySubscriptions }).(pulumi.StringArrayOutput)
}

type UserDefinedResourcesPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (UserDefinedResourcesPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDefinedResourcesPropertiesResponse)(nil)).Elem()
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) ToUserDefinedResourcesPropertiesResponsePtrOutput() UserDefinedResourcesPropertiesResponsePtrOutput {
	return o
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) ToUserDefinedResourcesPropertiesResponsePtrOutputWithContext(ctx context.Context) UserDefinedResourcesPropertiesResponsePtrOutput {
	return o
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) Elem() UserDefinedResourcesPropertiesResponseOutput {
	return o.ApplyT(func(v *UserDefinedResourcesPropertiesResponse) UserDefinedResourcesPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret UserDefinedResourcesPropertiesResponse
		return ret
	}).(UserDefinedResourcesPropertiesResponseOutput)
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDefinedResourcesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Query
	}).(pulumi.StringPtrOutput)
}

func (o UserDefinedResourcesPropertiesResponsePtrOutput) QuerySubscriptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UserDefinedResourcesPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.QuerySubscriptions
	}).(pulumi.StringArrayOutput)
}

type UserRecommendation struct {
	RecommendationAction *string `pulumi:"recommendationAction"`
	Username             *string `pulumi:"username"`
}





type UserRecommendationInput interface {
	pulumi.Input

	ToUserRecommendationOutput() UserRecommendationOutput
	ToUserRecommendationOutputWithContext(context.Context) UserRecommendationOutput
}

type UserRecommendationArgs struct {
	RecommendationAction pulumi.StringPtrInput `pulumi:"recommendationAction"`
	Username             pulumi.StringPtrInput `pulumi:"username"`
}

func (UserRecommendationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRecommendation)(nil)).Elem()
}

func (i UserRecommendationArgs) ToUserRecommendationOutput() UserRecommendationOutput {
	return i.ToUserRecommendationOutputWithContext(context.Background())
}

func (i UserRecommendationArgs) ToUserRecommendationOutputWithContext(ctx context.Context) UserRecommendationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRecommendationOutput)
}





type UserRecommendationArrayInput interface {
	pulumi.Input

	ToUserRecommendationArrayOutput() UserRecommendationArrayOutput
	ToUserRecommendationArrayOutputWithContext(context.Context) UserRecommendationArrayOutput
}

type UserRecommendationArray []UserRecommendationInput

func (UserRecommendationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserRecommendation)(nil)).Elem()
}

func (i UserRecommendationArray) ToUserRecommendationArrayOutput() UserRecommendationArrayOutput {
	return i.ToUserRecommendationArrayOutputWithContext(context.Background())
}

func (i UserRecommendationArray) ToUserRecommendationArrayOutputWithContext(ctx context.Context) UserRecommendationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRecommendationArrayOutput)
}

type UserRecommendationOutput struct{ *pulumi.OutputState }

func (UserRecommendationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRecommendation)(nil)).Elem()
}

func (o UserRecommendationOutput) ToUserRecommendationOutput() UserRecommendationOutput {
	return o
}

func (o UserRecommendationOutput) ToUserRecommendationOutputWithContext(ctx context.Context) UserRecommendationOutput {
	return o
}

func (o UserRecommendationOutput) RecommendationAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserRecommendation) *string { return v.RecommendationAction }).(pulumi.StringPtrOutput)
}

func (o UserRecommendationOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserRecommendation) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type UserRecommendationArrayOutput struct{ *pulumi.OutputState }

func (UserRecommendationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserRecommendation)(nil)).Elem()
}

func (o UserRecommendationArrayOutput) ToUserRecommendationArrayOutput() UserRecommendationArrayOutput {
	return o
}

func (o UserRecommendationArrayOutput) ToUserRecommendationArrayOutputWithContext(ctx context.Context) UserRecommendationArrayOutput {
	return o
}

func (o UserRecommendationArrayOutput) Index(i pulumi.IntInput) UserRecommendationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserRecommendation {
		return vs[0].([]UserRecommendation)[vs[1].(int)]
	}).(UserRecommendationOutput)
}

type UserRecommendationResponse struct {
	RecommendationAction *string `pulumi:"recommendationAction"`
	Username             *string `pulumi:"username"`
}





type UserRecommendationResponseInput interface {
	pulumi.Input

	ToUserRecommendationResponseOutput() UserRecommendationResponseOutput
	ToUserRecommendationResponseOutputWithContext(context.Context) UserRecommendationResponseOutput
}

type UserRecommendationResponseArgs struct {
	RecommendationAction pulumi.StringPtrInput `pulumi:"recommendationAction"`
	Username             pulumi.StringPtrInput `pulumi:"username"`
}

func (UserRecommendationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRecommendationResponse)(nil)).Elem()
}

func (i UserRecommendationResponseArgs) ToUserRecommendationResponseOutput() UserRecommendationResponseOutput {
	return i.ToUserRecommendationResponseOutputWithContext(context.Background())
}

func (i UserRecommendationResponseArgs) ToUserRecommendationResponseOutputWithContext(ctx context.Context) UserRecommendationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRecommendationResponseOutput)
}





type UserRecommendationResponseArrayInput interface {
	pulumi.Input

	ToUserRecommendationResponseArrayOutput() UserRecommendationResponseArrayOutput
	ToUserRecommendationResponseArrayOutputWithContext(context.Context) UserRecommendationResponseArrayOutput
}

type UserRecommendationResponseArray []UserRecommendationResponseInput

func (UserRecommendationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserRecommendationResponse)(nil)).Elem()
}

func (i UserRecommendationResponseArray) ToUserRecommendationResponseArrayOutput() UserRecommendationResponseArrayOutput {
	return i.ToUserRecommendationResponseArrayOutputWithContext(context.Background())
}

func (i UserRecommendationResponseArray) ToUserRecommendationResponseArrayOutputWithContext(ctx context.Context) UserRecommendationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRecommendationResponseArrayOutput)
}

type UserRecommendationResponseOutput struct{ *pulumi.OutputState }

func (UserRecommendationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRecommendationResponse)(nil)).Elem()
}

func (o UserRecommendationResponseOutput) ToUserRecommendationResponseOutput() UserRecommendationResponseOutput {
	return o
}

func (o UserRecommendationResponseOutput) ToUserRecommendationResponseOutputWithContext(ctx context.Context) UserRecommendationResponseOutput {
	return o
}

func (o UserRecommendationResponseOutput) RecommendationAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserRecommendationResponse) *string { return v.RecommendationAction }).(pulumi.StringPtrOutput)
}

func (o UserRecommendationResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserRecommendationResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type UserRecommendationResponseArrayOutput struct{ *pulumi.OutputState }

func (UserRecommendationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserRecommendationResponse)(nil)).Elem()
}

func (o UserRecommendationResponseArrayOutput) ToUserRecommendationResponseArrayOutput() UserRecommendationResponseArrayOutput {
	return o
}

func (o UserRecommendationResponseArrayOutput) ToUserRecommendationResponseArrayOutputWithContext(ctx context.Context) UserRecommendationResponseArrayOutput {
	return o
}

func (o UserRecommendationResponseArrayOutput) Index(i pulumi.IntInput) UserRecommendationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserRecommendationResponse {
		return vs[0].([]UserRecommendationResponse)[vs[1].(int)]
	}).(UserRecommendationResponseOutput)
}

type VmRecommendation struct {
	ConfigurationStatus  *string `pulumi:"configurationStatus"`
	EnforcementSupport   *string `pulumi:"enforcementSupport"`
	RecommendationAction *string `pulumi:"recommendationAction"`
	ResourceId           *string `pulumi:"resourceId"`
}





type VmRecommendationInput interface {
	pulumi.Input

	ToVmRecommendationOutput() VmRecommendationOutput
	ToVmRecommendationOutputWithContext(context.Context) VmRecommendationOutput
}

type VmRecommendationArgs struct {
	ConfigurationStatus  pulumi.StringPtrInput `pulumi:"configurationStatus"`
	EnforcementSupport   pulumi.StringPtrInput `pulumi:"enforcementSupport"`
	RecommendationAction pulumi.StringPtrInput `pulumi:"recommendationAction"`
	ResourceId           pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (VmRecommendationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmRecommendation)(nil)).Elem()
}

func (i VmRecommendationArgs) ToVmRecommendationOutput() VmRecommendationOutput {
	return i.ToVmRecommendationOutputWithContext(context.Background())
}

func (i VmRecommendationArgs) ToVmRecommendationOutputWithContext(ctx context.Context) VmRecommendationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmRecommendationOutput)
}





type VmRecommendationArrayInput interface {
	pulumi.Input

	ToVmRecommendationArrayOutput() VmRecommendationArrayOutput
	ToVmRecommendationArrayOutputWithContext(context.Context) VmRecommendationArrayOutput
}

type VmRecommendationArray []VmRecommendationInput

func (VmRecommendationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmRecommendation)(nil)).Elem()
}

func (i VmRecommendationArray) ToVmRecommendationArrayOutput() VmRecommendationArrayOutput {
	return i.ToVmRecommendationArrayOutputWithContext(context.Background())
}

func (i VmRecommendationArray) ToVmRecommendationArrayOutputWithContext(ctx context.Context) VmRecommendationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmRecommendationArrayOutput)
}

type VmRecommendationOutput struct{ *pulumi.OutputState }

func (VmRecommendationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmRecommendation)(nil)).Elem()
}

func (o VmRecommendationOutput) ToVmRecommendationOutput() VmRecommendationOutput {
	return o
}

func (o VmRecommendationOutput) ToVmRecommendationOutputWithContext(ctx context.Context) VmRecommendationOutput {
	return o
}

func (o VmRecommendationOutput) ConfigurationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendation) *string { return v.ConfigurationStatus }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationOutput) EnforcementSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendation) *string { return v.EnforcementSupport }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationOutput) RecommendationAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendation) *string { return v.RecommendationAction }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendation) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VmRecommendationArrayOutput struct{ *pulumi.OutputState }

func (VmRecommendationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmRecommendation)(nil)).Elem()
}

func (o VmRecommendationArrayOutput) ToVmRecommendationArrayOutput() VmRecommendationArrayOutput {
	return o
}

func (o VmRecommendationArrayOutput) ToVmRecommendationArrayOutputWithContext(ctx context.Context) VmRecommendationArrayOutput {
	return o
}

func (o VmRecommendationArrayOutput) Index(i pulumi.IntInput) VmRecommendationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmRecommendation {
		return vs[0].([]VmRecommendation)[vs[1].(int)]
	}).(VmRecommendationOutput)
}

type VmRecommendationResponse struct {
	ConfigurationStatus  *string `pulumi:"configurationStatus"`
	EnforcementSupport   *string `pulumi:"enforcementSupport"`
	RecommendationAction *string `pulumi:"recommendationAction"`
	ResourceId           *string `pulumi:"resourceId"`
}





type VmRecommendationResponseInput interface {
	pulumi.Input

	ToVmRecommendationResponseOutput() VmRecommendationResponseOutput
	ToVmRecommendationResponseOutputWithContext(context.Context) VmRecommendationResponseOutput
}

type VmRecommendationResponseArgs struct {
	ConfigurationStatus  pulumi.StringPtrInput `pulumi:"configurationStatus"`
	EnforcementSupport   pulumi.StringPtrInput `pulumi:"enforcementSupport"`
	RecommendationAction pulumi.StringPtrInput `pulumi:"recommendationAction"`
	ResourceId           pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (VmRecommendationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmRecommendationResponse)(nil)).Elem()
}

func (i VmRecommendationResponseArgs) ToVmRecommendationResponseOutput() VmRecommendationResponseOutput {
	return i.ToVmRecommendationResponseOutputWithContext(context.Background())
}

func (i VmRecommendationResponseArgs) ToVmRecommendationResponseOutputWithContext(ctx context.Context) VmRecommendationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmRecommendationResponseOutput)
}





type VmRecommendationResponseArrayInput interface {
	pulumi.Input

	ToVmRecommendationResponseArrayOutput() VmRecommendationResponseArrayOutput
	ToVmRecommendationResponseArrayOutputWithContext(context.Context) VmRecommendationResponseArrayOutput
}

type VmRecommendationResponseArray []VmRecommendationResponseInput

func (VmRecommendationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmRecommendationResponse)(nil)).Elem()
}

func (i VmRecommendationResponseArray) ToVmRecommendationResponseArrayOutput() VmRecommendationResponseArrayOutput {
	return i.ToVmRecommendationResponseArrayOutputWithContext(context.Background())
}

func (i VmRecommendationResponseArray) ToVmRecommendationResponseArrayOutputWithContext(ctx context.Context) VmRecommendationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmRecommendationResponseArrayOutput)
}

type VmRecommendationResponseOutput struct{ *pulumi.OutputState }

func (VmRecommendationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmRecommendationResponse)(nil)).Elem()
}

func (o VmRecommendationResponseOutput) ToVmRecommendationResponseOutput() VmRecommendationResponseOutput {
	return o
}

func (o VmRecommendationResponseOutput) ToVmRecommendationResponseOutputWithContext(ctx context.Context) VmRecommendationResponseOutput {
	return o
}

func (o VmRecommendationResponseOutput) ConfigurationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendationResponse) *string { return v.ConfigurationStatus }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationResponseOutput) EnforcementSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendationResponse) *string { return v.EnforcementSupport }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationResponseOutput) RecommendationAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendationResponse) *string { return v.RecommendationAction }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VmRecommendationResponseArrayOutput struct{ *pulumi.OutputState }

func (VmRecommendationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmRecommendationResponse)(nil)).Elem()
}

func (o VmRecommendationResponseArrayOutput) ToVmRecommendationResponseArrayOutput() VmRecommendationResponseArrayOutput {
	return o
}

func (o VmRecommendationResponseArrayOutput) ToVmRecommendationResponseArrayOutputWithContext(ctx context.Context) VmRecommendationResponseArrayOutput {
	return o
}

func (o VmRecommendationResponseArrayOutput) Index(i pulumi.IntInput) VmRecommendationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmRecommendationResponse {
		return vs[0].([]VmRecommendationResponse)[vs[1].(int)]
	}).(VmRecommendationResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AdaptiveApplicationControlIssueSummaryResponseOutput{})
	pulumi.RegisterOutputType(AdaptiveApplicationControlIssueSummaryResponseArrayOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspacesPropertiesOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspacesPropertiesArrayOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspacesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspacesPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(AllowlistCustomAlertRuleOutput{})
	pulumi.RegisterOutputType(AllowlistCustomAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(AllowlistCustomAlertRuleResponseOutput{})
	pulumi.RegisterOutputType(AllowlistCustomAlertRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(AssessmentLinksResponseOutput{})
	pulumi.RegisterOutputType(AssessmentLinksResponsePtrOutput{})
	pulumi.RegisterOutputType(AssessmentStatusOutput{})
	pulumi.RegisterOutputType(AssessmentStatusPtrOutput{})
	pulumi.RegisterOutputType(AssessmentStatusResponseOutput{})
	pulumi.RegisterOutputType(AssessmentStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(AssignedComponentItemOutput{})
	pulumi.RegisterOutputType(AssignedComponentItemPtrOutput{})
	pulumi.RegisterOutputType(AssignedComponentItemResponseOutput{})
	pulumi.RegisterOutputType(AssignedComponentItemResponsePtrOutput{})
	pulumi.RegisterOutputType(AssignedStandardItemOutput{})
	pulumi.RegisterOutputType(AssignedStandardItemPtrOutput{})
	pulumi.RegisterOutputType(AssignedStandardItemResponseOutput{})
	pulumi.RegisterOutputType(AssignedStandardItemResponsePtrOutput{})
	pulumi.RegisterOutputType(AssignmentPropertiesAdditionalDataOutput{})
	pulumi.RegisterOutputType(AssignmentPropertiesAdditionalDataPtrOutput{})
	pulumi.RegisterOutputType(AssignmentPropertiesResponseAdditionalDataOutput{})
	pulumi.RegisterOutputType(AssignmentPropertiesResponseAdditionalDataPtrOutput{})
	pulumi.RegisterOutputType(AutomationActionEventHubOutput{})
	pulumi.RegisterOutputType(AutomationActionEventHubResponseOutput{})
	pulumi.RegisterOutputType(AutomationActionLogicAppOutput{})
	pulumi.RegisterOutputType(AutomationActionLogicAppResponseOutput{})
	pulumi.RegisterOutputType(AutomationActionWorkspaceOutput{})
	pulumi.RegisterOutputType(AutomationActionWorkspaceResponseOutput{})
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
	pulumi.RegisterOutputType(AwAssumeRoleAuthenticationDetailsPropertiesOutput{})
	pulumi.RegisterOutputType(AwAssumeRoleAuthenticationDetailsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AwsCredsAuthenticationDetailsPropertiesOutput{})
	pulumi.RegisterOutputType(AwsCredsAuthenticationDetailsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AzureResourceDetailsOutput{})
	pulumi.RegisterOutputType(AzureResourceDetailsResponseOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingNativeCloudConnectionOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingNativeCloudConnectionPtrOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingResponseOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingResponseNativeCloudConnectionOutput{})
	pulumi.RegisterOutputType(CspmMonitorAwsOfferingResponseNativeCloudConnectionPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingCloudWatchToKinesisOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingCloudWatchToKinesisPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKinesisToS3Output{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKinesisToS3PtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKubernetesScubaReaderOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKubernetesScubaReaderPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKubernetesServiceOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingKubernetesServicePtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseCloudWatchToKinesisPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKinesisToS3Output{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKinesisToS3PtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKubernetesScubaReaderPtrOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKubernetesServiceOutput{})
	pulumi.RegisterOutputType(DefenderForContainersAwsOfferingResponseKubernetesServicePtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingArcAutoProvisioningOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingArcAutoProvisioningPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingDefenderForServersOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingDefenderForServersPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseArcAutoProvisioningOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseArcAutoProvisioningPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseDefenderForServersOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseDefenderForServersPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadataPtrOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingServicePrincipalSecretMetadataOutput{})
	pulumi.RegisterOutputType(DefenderForServersAwsOfferingServicePrincipalSecretMetadataPtrOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleResponseOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(GcpCredentialsDetailsPropertiesOutput{})
	pulumi.RegisterOutputType(GcpCredentialsDetailsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HybridComputeSettingsPropertiesOutput{})
	pulumi.RegisterOutputType(HybridComputeSettingsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(HybridComputeSettingsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HybridComputeSettingsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IngestionConnectionStringResponseOutput{})
	pulumi.RegisterOutputType(IngestionConnectionStringResponseArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPolicyVirtualMachineOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPolicyVirtualMachineArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPolicyVirtualMachineResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPolicyVirtualMachineResponseArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPortRuleOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPortRuleArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPortRuleResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPortRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestPortOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestPortArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestPortResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestPortResponseArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestResponseArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestVirtualMachineOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestVirtualMachineArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestVirtualMachineResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestVirtualMachineResponseArrayOutput{})
	pulumi.RegisterOutputType(OnPremiseResourceDetailsOutput{})
	pulumi.RegisterOutputType(OnPremiseResourceDetailsResponseOutput{})
	pulumi.RegisterOutputType(OnPremiseSqlResourceDetailsOutput{})
	pulumi.RegisterOutputType(OnPremiseSqlResourceDetailsResponseOutput{})
	pulumi.RegisterOutputType(PathRecommendationOutput{})
	pulumi.RegisterOutputType(PathRecommendationArrayOutput{})
	pulumi.RegisterOutputType(PathRecommendationResponseOutput{})
	pulumi.RegisterOutputType(PathRecommendationResponseArrayOutput{})
	pulumi.RegisterOutputType(ProtectionModeOutput{})
	pulumi.RegisterOutputType(ProtectionModePtrOutput{})
	pulumi.RegisterOutputType(ProtectionModeResponseOutput{})
	pulumi.RegisterOutputType(ProtectionModeResponsePtrOutput{})
	pulumi.RegisterOutputType(ProxyServerPropertiesOutput{})
	pulumi.RegisterOutputType(ProxyServerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ProxyServerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProxyServerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PublisherInfoOutput{})
	pulumi.RegisterOutputType(PublisherInfoPtrOutput{})
	pulumi.RegisterOutputType(PublisherInfoResponseOutput{})
	pulumi.RegisterOutputType(PublisherInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(RecommendationConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(RecommendationConfigurationPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RecommendationConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RecommendationConfigurationPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(RuleResultsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RuleResultsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeElementOutput{})
	pulumi.RegisterOutputType(ScopeElementArrayOutput{})
	pulumi.RegisterOutputType(ScopeElementResponseOutput{})
	pulumi.RegisterOutputType(ScopeElementResponseArrayOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesOrganizationalDataOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesOrganizationalDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesResponseOrganizationalDataOutput{})
	pulumi.RegisterOutputType(SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesAlertNotificationsOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesAlertNotificationsPtrOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesNotificationsByRoleOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesNotificationsByRolePtrOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesResponseAlertNotificationsOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesResponseAlertNotificationsPtrOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesResponseNotificationsByRoleOutput{})
	pulumi.RegisterOutputType(SecurityContactPropertiesResponseNotificationsByRolePtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServicePrincipalPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StandardComponentPropertiesOutput{})
	pulumi.RegisterOutputType(StandardComponentPropertiesArrayOutput{})
	pulumi.RegisterOutputType(StandardComponentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StandardComponentPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopeOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopePtrOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopeResponseOutput{})
	pulumi.RegisterOutputType(SuppressionAlertsScopeResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(ThresholdCustomAlertRuleOutput{})
	pulumi.RegisterOutputType(ThresholdCustomAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(ThresholdCustomAlertRuleResponseOutput{})
	pulumi.RegisterOutputType(ThresholdCustomAlertRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(TimeWindowCustomAlertRuleOutput{})
	pulumi.RegisterOutputType(TimeWindowCustomAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(TimeWindowCustomAlertRuleResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowCustomAlertRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(UserDefinedResourcesPropertiesOutput{})
	pulumi.RegisterOutputType(UserDefinedResourcesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(UserDefinedResourcesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserDefinedResourcesPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(UserRecommendationOutput{})
	pulumi.RegisterOutputType(UserRecommendationArrayOutput{})
	pulumi.RegisterOutputType(UserRecommendationResponseOutput{})
	pulumi.RegisterOutputType(UserRecommendationResponseArrayOutput{})
	pulumi.RegisterOutputType(VmRecommendationOutput{})
	pulumi.RegisterOutputType(VmRecommendationArrayOutput{})
	pulumi.RegisterOutputType(VmRecommendationResponseOutput{})
	pulumi.RegisterOutputType(VmRecommendationResponseArrayOutput{})
}
