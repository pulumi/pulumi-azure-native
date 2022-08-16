


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


func (val *AdditionalWorkspacesProperties) Defaults() *AdditionalWorkspacesProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "Sentinel"
		tmp.Type = &type_
	}
	return &tmp
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


func (val *AdditionalWorkspacesPropertiesArgs) Defaults() *AdditionalWorkspacesPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		tmp.Type = pulumi.StringPtr("Sentinel")
	}
	return &tmp
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


func (val *AdditionalWorkspacesPropertiesResponse) Defaults() *AdditionalWorkspacesPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "Sentinel"
		tmp.Type = &type_
	}
	return &tmp
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

type AwAssumeRoleAuthenticationDetailsProperties struct {
	AuthenticationType string `pulumi:"authenticationType"`
	AwsAssumeRoleArn   string `pulumi:"awsAssumeRoleArn"`
	AwsExternalId      string `pulumi:"awsExternalId"`
}

type AwAssumeRoleAuthenticationDetailsPropertiesResponse struct {
	AccountId                       string   `pulumi:"accountId"`
	AuthenticationProvisioningState string   `pulumi:"authenticationProvisioningState"`
	AuthenticationType              string   `pulumi:"authenticationType"`
	AwsAssumeRoleArn                string   `pulumi:"awsAssumeRoleArn"`
	AwsExternalId                   string   `pulumi:"awsExternalId"`
	GrantedPermissions              []string `pulumi:"grantedPermissions"`
}

type AwsCredsAuthenticationDetailsProperties struct {
	AuthenticationType string `pulumi:"authenticationType"`
	AwsAccessKeyId     string `pulumi:"awsAccessKeyId"`
	AwsSecretAccessKey string `pulumi:"awsSecretAccessKey"`
}

type AwsCredsAuthenticationDetailsPropertiesResponse struct {
	AccountId                       string   `pulumi:"accountId"`
	AuthenticationProvisioningState string   `pulumi:"authenticationProvisioningState"`
	AuthenticationType              string   `pulumi:"authenticationType"`
	AwsAccessKeyId                  string   `pulumi:"awsAccessKeyId"`
	AwsSecretAccessKey              string   `pulumi:"awsSecretAccessKey"`
	GrantedPermissions              []string `pulumi:"grantedPermissions"`
}

type AzureResourceDetails struct {
	Source string `pulumi:"source"`
}

type AzureResourceDetailsResponse struct {
	Id     string `pulumi:"id"`
	Source string `pulumi:"source"`
}

type CspmMonitorAwsOffering struct {
	NativeCloudConnection *CspmMonitorAwsOfferingNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                       `pulumi:"offeringType"`
}

type CspmMonitorAwsOfferingNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type CspmMonitorAwsOfferingResponse struct {
	Description           string                                               `pulumi:"description"`
	NativeCloudConnection *CspmMonitorAwsOfferingResponseNativeCloudConnection `pulumi:"nativeCloudConnection"`
	OfferingType          string                                               `pulumi:"offeringType"`
}

type CspmMonitorAwsOfferingResponseNativeCloudConnection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOffering struct {
	CloudWatchToKinesis   *DefenderForContainersAwsOfferingCloudWatchToKinesis   `pulumi:"cloudWatchToKinesis"`
	KinesisToS3           *DefenderForContainersAwsOfferingKinesisToS3           `pulumi:"kinesisToS3"`
	KubernetesScubaReader *DefenderForContainersAwsOfferingKubernetesScubaReader `pulumi:"kubernetesScubaReader"`
	KubernetesService     *DefenderForContainersAwsOfferingKubernetesService     `pulumi:"kubernetesService"`
	OfferingType          string                                                 `pulumi:"offeringType"`
}

type DefenderForContainersAwsOfferingCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponse struct {
	CloudWatchToKinesis   *DefenderForContainersAwsOfferingResponseCloudWatchToKinesis   `pulumi:"cloudWatchToKinesis"`
	Description           string                                                         `pulumi:"description"`
	KinesisToS3           *DefenderForContainersAwsOfferingResponseKinesisToS3           `pulumi:"kinesisToS3"`
	KubernetesScubaReader *DefenderForContainersAwsOfferingResponseKubernetesScubaReader `pulumi:"kubernetesScubaReader"`
	KubernetesService     *DefenderForContainersAwsOfferingResponseKubernetesService     `pulumi:"kubernetesService"`
	OfferingType          string                                                         `pulumi:"offeringType"`
}

type DefenderForContainersAwsOfferingResponseCloudWatchToKinesis struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseKinesisToS3 struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseKubernetesScubaReader struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForContainersAwsOfferingResponseKubernetesService struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForServersAwsOffering struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingDefenderForServers  `pulumi:"defenderForServers"`
	OfferingType        string                                            `pulumi:"offeringType"`
}

type DefenderForServersAwsOfferingArcAutoProvisioning struct {
	Enabled                        *bool                                                        `pulumi:"enabled"`
	ServicePrincipalSecretMetadata *DefenderForServersAwsOfferingServicePrincipalSecretMetadata `pulumi:"servicePrincipalSecretMetadata"`
}

type DefenderForServersAwsOfferingDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForServersAwsOfferingResponse struct {
	ArcAutoProvisioning *DefenderForServersAwsOfferingResponseArcAutoProvisioning `pulumi:"arcAutoProvisioning"`
	DefenderForServers  *DefenderForServersAwsOfferingResponseDefenderForServers  `pulumi:"defenderForServers"`
	Description         string                                                    `pulumi:"description"`
	OfferingType        string                                                    `pulumi:"offeringType"`
}

type DefenderForServersAwsOfferingResponseArcAutoProvisioning struct {
	Enabled                        *bool                                                                `pulumi:"enabled"`
	ServicePrincipalSecretMetadata *DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata `pulumi:"servicePrincipalSecretMetadata"`
}

type DefenderForServersAwsOfferingResponseDefenderForServers struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type DefenderForServersAwsOfferingResponseServicePrincipalSecretMetadata struct {
	ExpiryDate           *string `pulumi:"expiryDate"`
	ParameterNameInStore *string `pulumi:"parameterNameInStore"`
	ParameterStoreRegion *string `pulumi:"parameterStoreRegion"`
}

type DefenderForServersAwsOfferingServicePrincipalSecretMetadata struct {
	ExpiryDate           *string `pulumi:"expiryDate"`
	ParameterNameInStore *string `pulumi:"parameterNameInStore"`
	ParameterStoreRegion *string `pulumi:"parameterStoreRegion"`
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

type InformationProtectionAwsOffering struct {
	InformationProtection *InformationProtectionAwsOfferingInformationProtection `pulumi:"informationProtection"`
	OfferingType          string                                                 `pulumi:"offeringType"`
}

type InformationProtectionAwsOfferingInformationProtection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type InformationProtectionAwsOfferingResponse struct {
	Description           string                                                         `pulumi:"description"`
	InformationProtection *InformationProtectionAwsOfferingResponseInformationProtection `pulumi:"informationProtection"`
	OfferingType          string                                                         `pulumi:"offeringType"`
}

type InformationProtectionAwsOfferingResponseInformationProtection struct {
	CloudRoleArn *string `pulumi:"cloudRoleArn"`
}

type IngestionConnectionStringResponse struct {
	Location string `pulumi:"location"`
	Value    string `pulumi:"value"`
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


func (val *RecommendationConfigurationProperties) Defaults() *RecommendationConfigurationProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = "Enabled"
	}
	return &tmp
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


func (val *RecommendationConfigurationPropertiesArgs) Defaults() *RecommendationConfigurationPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = pulumi.String("Enabled")
	}
	return &tmp
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


func (val *RecommendationConfigurationPropertiesResponse) Defaults() *RecommendationConfigurationPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = "Enabled"
	}
	return &tmp
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

func (o RuleResultsPropertiesResponseOutput) Results() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v RuleResultsPropertiesResponse) [][]string { return v.Results }).(pulumi.StringArrayArrayOutput)
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
	pulumi.RegisterOutputType(AssessmentStatusOutput{})
	pulumi.RegisterOutputType(AssessmentStatusResponseOutput{})
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
	pulumi.RegisterOutputType(DenylistCustomAlertRuleOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleResponseOutput{})
	pulumi.RegisterOutputType(DenylistCustomAlertRuleResponseArrayOutput{})
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
