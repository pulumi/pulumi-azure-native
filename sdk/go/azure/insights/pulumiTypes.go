


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionGroupType struct {
	ActionGroupId     string            `pulumi:"actionGroupId"`
	WebhookProperties map[string]string `pulumi:"webhookProperties"`
}





type ActionGroupTypeInput interface {
	pulumi.Input

	ToActionGroupTypeOutput() ActionGroupTypeOutput
	ToActionGroupTypeOutputWithContext(context.Context) ActionGroupTypeOutput
}

type ActionGroupTypeArgs struct {
	ActionGroupId     pulumi.StringInput    `pulumi:"actionGroupId"`
	WebhookProperties pulumi.StringMapInput `pulumi:"webhookProperties"`
}

func (ActionGroupTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupType)(nil)).Elem()
}

func (i ActionGroupTypeArgs) ToActionGroupTypeOutput() ActionGroupTypeOutput {
	return i.ToActionGroupTypeOutputWithContext(context.Background())
}

func (i ActionGroupTypeArgs) ToActionGroupTypeOutputWithContext(ctx context.Context) ActionGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupTypeOutput)
}





type ActionGroupTypeArrayInput interface {
	pulumi.Input

	ToActionGroupTypeArrayOutput() ActionGroupTypeArrayOutput
	ToActionGroupTypeArrayOutputWithContext(context.Context) ActionGroupTypeArrayOutput
}

type ActionGroupTypeArray []ActionGroupTypeInput

func (ActionGroupTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionGroupType)(nil)).Elem()
}

func (i ActionGroupTypeArray) ToActionGroupTypeArrayOutput() ActionGroupTypeArrayOutput {
	return i.ToActionGroupTypeArrayOutputWithContext(context.Background())
}

func (i ActionGroupTypeArray) ToActionGroupTypeArrayOutputWithContext(ctx context.Context) ActionGroupTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupTypeArrayOutput)
}

type ActionGroupTypeOutput struct{ *pulumi.OutputState }

func (ActionGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupType)(nil)).Elem()
}

func (o ActionGroupTypeOutput) ToActionGroupTypeOutput() ActionGroupTypeOutput {
	return o
}

func (o ActionGroupTypeOutput) ToActionGroupTypeOutputWithContext(ctx context.Context) ActionGroupTypeOutput {
	return o
}

func (o ActionGroupTypeOutput) ActionGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroupType) string { return v.ActionGroupId }).(pulumi.StringOutput)
}

func (o ActionGroupTypeOutput) WebhookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ActionGroupType) map[string]string { return v.WebhookProperties }).(pulumi.StringMapOutput)
}

type ActionGroupTypeArrayOutput struct{ *pulumi.OutputState }

func (ActionGroupTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionGroupType)(nil)).Elem()
}

func (o ActionGroupTypeArrayOutput) ToActionGroupTypeArrayOutput() ActionGroupTypeArrayOutput {
	return o
}

func (o ActionGroupTypeArrayOutput) ToActionGroupTypeArrayOutputWithContext(ctx context.Context) ActionGroupTypeArrayOutput {
	return o
}

func (o ActionGroupTypeArrayOutput) Index(i pulumi.IntInput) ActionGroupTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActionGroupType {
		return vs[0].([]ActionGroupType)[vs[1].(int)]
	}).(ActionGroupTypeOutput)
}

type ActionGroupResponse struct {
	ActionGroupId     string            `pulumi:"actionGroupId"`
	WebhookProperties map[string]string `pulumi:"webhookProperties"`
}

type ActionGroupResponseOutput struct{ *pulumi.OutputState }

func (ActionGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupResponse)(nil)).Elem()
}

func (o ActionGroupResponseOutput) ToActionGroupResponseOutput() ActionGroupResponseOutput {
	return o
}

func (o ActionGroupResponseOutput) ToActionGroupResponseOutputWithContext(ctx context.Context) ActionGroupResponseOutput {
	return o
}

func (o ActionGroupResponseOutput) ActionGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroupResponse) string { return v.ActionGroupId }).(pulumi.StringOutput)
}

func (o ActionGroupResponseOutput) WebhookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ActionGroupResponse) map[string]string { return v.WebhookProperties }).(pulumi.StringMapOutput)
}

type ActionGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (ActionGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionGroupResponse)(nil)).Elem()
}

func (o ActionGroupResponseArrayOutput) ToActionGroupResponseArrayOutput() ActionGroupResponseArrayOutput {
	return o
}

func (o ActionGroupResponseArrayOutput) ToActionGroupResponseArrayOutputWithContext(ctx context.Context) ActionGroupResponseArrayOutput {
	return o
}

func (o ActionGroupResponseArrayOutput) Index(i pulumi.IntInput) ActionGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActionGroupResponse {
		return vs[0].([]ActionGroupResponse)[vs[1].(int)]
	}).(ActionGroupResponseOutput)
}

type ActionList struct {
	ActionGroups []ActionGroupType `pulumi:"actionGroups"`
}





type ActionListInput interface {
	pulumi.Input

	ToActionListOutput() ActionListOutput
	ToActionListOutputWithContext(context.Context) ActionListOutput
}

type ActionListArgs struct {
	ActionGroups ActionGroupTypeArrayInput `pulumi:"actionGroups"`
}

func (ActionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionList)(nil)).Elem()
}

func (i ActionListArgs) ToActionListOutput() ActionListOutput {
	return i.ToActionListOutputWithContext(context.Background())
}

func (i ActionListArgs) ToActionListOutputWithContext(ctx context.Context) ActionListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionListOutput)
}

type ActionListOutput struct{ *pulumi.OutputState }

func (ActionListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionList)(nil)).Elem()
}

func (o ActionListOutput) ToActionListOutput() ActionListOutput {
	return o
}

func (o ActionListOutput) ToActionListOutputWithContext(ctx context.Context) ActionListOutput {
	return o
}

func (o ActionListOutput) ActionGroups() ActionGroupTypeArrayOutput {
	return o.ApplyT(func(v ActionList) []ActionGroupType { return v.ActionGroups }).(ActionGroupTypeArrayOutput)
}

type ActionListResponse struct {
	ActionGroups []ActionGroupResponse `pulumi:"actionGroups"`
}

type ActionListResponseOutput struct{ *pulumi.OutputState }

func (ActionListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionListResponse)(nil)).Elem()
}

func (o ActionListResponseOutput) ToActionListResponseOutput() ActionListResponseOutput {
	return o
}

func (o ActionListResponseOutput) ToActionListResponseOutputWithContext(ctx context.Context) ActionListResponseOutput {
	return o
}

func (o ActionListResponseOutput) ActionGroups() ActionGroupResponseArrayOutput {
	return o.ApplyT(func(v ActionListResponse) []ActionGroupResponse { return v.ActionGroups }).(ActionGroupResponseArrayOutput)
}

type AlertRuleAllOfCondition struct {
	AllOf []AlertRuleAnyOfOrLeafCondition `pulumi:"allOf"`
}





type AlertRuleAllOfConditionInput interface {
	pulumi.Input

	ToAlertRuleAllOfConditionOutput() AlertRuleAllOfConditionOutput
	ToAlertRuleAllOfConditionOutputWithContext(context.Context) AlertRuleAllOfConditionOutput
}

type AlertRuleAllOfConditionArgs struct {
	AllOf AlertRuleAnyOfOrLeafConditionArrayInput `pulumi:"allOf"`
}

func (AlertRuleAllOfConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleAllOfCondition)(nil)).Elem()
}

func (i AlertRuleAllOfConditionArgs) ToAlertRuleAllOfConditionOutput() AlertRuleAllOfConditionOutput {
	return i.ToAlertRuleAllOfConditionOutputWithContext(context.Background())
}

func (i AlertRuleAllOfConditionArgs) ToAlertRuleAllOfConditionOutputWithContext(ctx context.Context) AlertRuleAllOfConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAllOfConditionOutput)
}

type AlertRuleAllOfConditionOutput struct{ *pulumi.OutputState }

func (AlertRuleAllOfConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleAllOfCondition)(nil)).Elem()
}

func (o AlertRuleAllOfConditionOutput) ToAlertRuleAllOfConditionOutput() AlertRuleAllOfConditionOutput {
	return o
}

func (o AlertRuleAllOfConditionOutput) ToAlertRuleAllOfConditionOutputWithContext(ctx context.Context) AlertRuleAllOfConditionOutput {
	return o
}

func (o AlertRuleAllOfConditionOutput) AllOf() AlertRuleAnyOfOrLeafConditionArrayOutput {
	return o.ApplyT(func(v AlertRuleAllOfCondition) []AlertRuleAnyOfOrLeafCondition { return v.AllOf }).(AlertRuleAnyOfOrLeafConditionArrayOutput)
}

type AlertRuleAllOfConditionResponse struct {
	AllOf []AlertRuleAnyOfOrLeafConditionResponse `pulumi:"allOf"`
}

type AlertRuleAllOfConditionResponseOutput struct{ *pulumi.OutputState }

func (AlertRuleAllOfConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleAllOfConditionResponse)(nil)).Elem()
}

func (o AlertRuleAllOfConditionResponseOutput) ToAlertRuleAllOfConditionResponseOutput() AlertRuleAllOfConditionResponseOutput {
	return o
}

func (o AlertRuleAllOfConditionResponseOutput) ToAlertRuleAllOfConditionResponseOutputWithContext(ctx context.Context) AlertRuleAllOfConditionResponseOutput {
	return o
}

func (o AlertRuleAllOfConditionResponseOutput) AllOf() AlertRuleAnyOfOrLeafConditionResponseArrayOutput {
	return o.ApplyT(func(v AlertRuleAllOfConditionResponse) []AlertRuleAnyOfOrLeafConditionResponse { return v.AllOf }).(AlertRuleAnyOfOrLeafConditionResponseArrayOutput)
}

type AlertRuleAnyOfOrLeafCondition struct {
	AnyOf       []AlertRuleLeafCondition `pulumi:"anyOf"`
	ContainsAny []string                 `pulumi:"containsAny"`
	Equals      *string                  `pulumi:"equals"`
	Field       *string                  `pulumi:"field"`
}





type AlertRuleAnyOfOrLeafConditionInput interface {
	pulumi.Input

	ToAlertRuleAnyOfOrLeafConditionOutput() AlertRuleAnyOfOrLeafConditionOutput
	ToAlertRuleAnyOfOrLeafConditionOutputWithContext(context.Context) AlertRuleAnyOfOrLeafConditionOutput
}

type AlertRuleAnyOfOrLeafConditionArgs struct {
	AnyOf       AlertRuleLeafConditionArrayInput `pulumi:"anyOf"`
	ContainsAny pulumi.StringArrayInput          `pulumi:"containsAny"`
	Equals      pulumi.StringPtrInput            `pulumi:"equals"`
	Field       pulumi.StringPtrInput            `pulumi:"field"`
}

func (AlertRuleAnyOfOrLeafConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleAnyOfOrLeafCondition)(nil)).Elem()
}

func (i AlertRuleAnyOfOrLeafConditionArgs) ToAlertRuleAnyOfOrLeafConditionOutput() AlertRuleAnyOfOrLeafConditionOutput {
	return i.ToAlertRuleAnyOfOrLeafConditionOutputWithContext(context.Background())
}

func (i AlertRuleAnyOfOrLeafConditionArgs) ToAlertRuleAnyOfOrLeafConditionOutputWithContext(ctx context.Context) AlertRuleAnyOfOrLeafConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAnyOfOrLeafConditionOutput)
}





type AlertRuleAnyOfOrLeafConditionArrayInput interface {
	pulumi.Input

	ToAlertRuleAnyOfOrLeafConditionArrayOutput() AlertRuleAnyOfOrLeafConditionArrayOutput
	ToAlertRuleAnyOfOrLeafConditionArrayOutputWithContext(context.Context) AlertRuleAnyOfOrLeafConditionArrayOutput
}

type AlertRuleAnyOfOrLeafConditionArray []AlertRuleAnyOfOrLeafConditionInput

func (AlertRuleAnyOfOrLeafConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertRuleAnyOfOrLeafCondition)(nil)).Elem()
}

func (i AlertRuleAnyOfOrLeafConditionArray) ToAlertRuleAnyOfOrLeafConditionArrayOutput() AlertRuleAnyOfOrLeafConditionArrayOutput {
	return i.ToAlertRuleAnyOfOrLeafConditionArrayOutputWithContext(context.Background())
}

func (i AlertRuleAnyOfOrLeafConditionArray) ToAlertRuleAnyOfOrLeafConditionArrayOutputWithContext(ctx context.Context) AlertRuleAnyOfOrLeafConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAnyOfOrLeafConditionArrayOutput)
}

type AlertRuleAnyOfOrLeafConditionOutput struct{ *pulumi.OutputState }

func (AlertRuleAnyOfOrLeafConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleAnyOfOrLeafCondition)(nil)).Elem()
}

func (o AlertRuleAnyOfOrLeafConditionOutput) ToAlertRuleAnyOfOrLeafConditionOutput() AlertRuleAnyOfOrLeafConditionOutput {
	return o
}

func (o AlertRuleAnyOfOrLeafConditionOutput) ToAlertRuleAnyOfOrLeafConditionOutputWithContext(ctx context.Context) AlertRuleAnyOfOrLeafConditionOutput {
	return o
}

func (o AlertRuleAnyOfOrLeafConditionOutput) AnyOf() AlertRuleLeafConditionArrayOutput {
	return o.ApplyT(func(v AlertRuleAnyOfOrLeafCondition) []AlertRuleLeafCondition { return v.AnyOf }).(AlertRuleLeafConditionArrayOutput)
}

func (o AlertRuleAnyOfOrLeafConditionOutput) ContainsAny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertRuleAnyOfOrLeafCondition) []string { return v.ContainsAny }).(pulumi.StringArrayOutput)
}

func (o AlertRuleAnyOfOrLeafConditionOutput) Equals() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertRuleAnyOfOrLeafCondition) *string { return v.Equals }).(pulumi.StringPtrOutput)
}

func (o AlertRuleAnyOfOrLeafConditionOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertRuleAnyOfOrLeafCondition) *string { return v.Field }).(pulumi.StringPtrOutput)
}

type AlertRuleAnyOfOrLeafConditionArrayOutput struct{ *pulumi.OutputState }

func (AlertRuleAnyOfOrLeafConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertRuleAnyOfOrLeafCondition)(nil)).Elem()
}

func (o AlertRuleAnyOfOrLeafConditionArrayOutput) ToAlertRuleAnyOfOrLeafConditionArrayOutput() AlertRuleAnyOfOrLeafConditionArrayOutput {
	return o
}

func (o AlertRuleAnyOfOrLeafConditionArrayOutput) ToAlertRuleAnyOfOrLeafConditionArrayOutputWithContext(ctx context.Context) AlertRuleAnyOfOrLeafConditionArrayOutput {
	return o
}

func (o AlertRuleAnyOfOrLeafConditionArrayOutput) Index(i pulumi.IntInput) AlertRuleAnyOfOrLeafConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertRuleAnyOfOrLeafCondition {
		return vs[0].([]AlertRuleAnyOfOrLeafCondition)[vs[1].(int)]
	}).(AlertRuleAnyOfOrLeafConditionOutput)
}

type AlertRuleAnyOfOrLeafConditionResponse struct {
	AnyOf       []AlertRuleLeafConditionResponse `pulumi:"anyOf"`
	ContainsAny []string                         `pulumi:"containsAny"`
	Equals      *string                          `pulumi:"equals"`
	Field       *string                          `pulumi:"field"`
}

type AlertRuleAnyOfOrLeafConditionResponseOutput struct{ *pulumi.OutputState }

func (AlertRuleAnyOfOrLeafConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleAnyOfOrLeafConditionResponse)(nil)).Elem()
}

func (o AlertRuleAnyOfOrLeafConditionResponseOutput) ToAlertRuleAnyOfOrLeafConditionResponseOutput() AlertRuleAnyOfOrLeafConditionResponseOutput {
	return o
}

func (o AlertRuleAnyOfOrLeafConditionResponseOutput) ToAlertRuleAnyOfOrLeafConditionResponseOutputWithContext(ctx context.Context) AlertRuleAnyOfOrLeafConditionResponseOutput {
	return o
}

func (o AlertRuleAnyOfOrLeafConditionResponseOutput) AnyOf() AlertRuleLeafConditionResponseArrayOutput {
	return o.ApplyT(func(v AlertRuleAnyOfOrLeafConditionResponse) []AlertRuleLeafConditionResponse { return v.AnyOf }).(AlertRuleLeafConditionResponseArrayOutput)
}

func (o AlertRuleAnyOfOrLeafConditionResponseOutput) ContainsAny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertRuleAnyOfOrLeafConditionResponse) []string { return v.ContainsAny }).(pulumi.StringArrayOutput)
}

func (o AlertRuleAnyOfOrLeafConditionResponseOutput) Equals() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertRuleAnyOfOrLeafConditionResponse) *string { return v.Equals }).(pulumi.StringPtrOutput)
}

func (o AlertRuleAnyOfOrLeafConditionResponseOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertRuleAnyOfOrLeafConditionResponse) *string { return v.Field }).(pulumi.StringPtrOutput)
}

type AlertRuleAnyOfOrLeafConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (AlertRuleAnyOfOrLeafConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertRuleAnyOfOrLeafConditionResponse)(nil)).Elem()
}

func (o AlertRuleAnyOfOrLeafConditionResponseArrayOutput) ToAlertRuleAnyOfOrLeafConditionResponseArrayOutput() AlertRuleAnyOfOrLeafConditionResponseArrayOutput {
	return o
}

func (o AlertRuleAnyOfOrLeafConditionResponseArrayOutput) ToAlertRuleAnyOfOrLeafConditionResponseArrayOutputWithContext(ctx context.Context) AlertRuleAnyOfOrLeafConditionResponseArrayOutput {
	return o
}

func (o AlertRuleAnyOfOrLeafConditionResponseArrayOutput) Index(i pulumi.IntInput) AlertRuleAnyOfOrLeafConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertRuleAnyOfOrLeafConditionResponse {
		return vs[0].([]AlertRuleAnyOfOrLeafConditionResponse)[vs[1].(int)]
	}).(AlertRuleAnyOfOrLeafConditionResponseOutput)
}

type AlertRuleLeafCondition struct {
	ContainsAny []string `pulumi:"containsAny"`
	Equals      *string  `pulumi:"equals"`
	Field       *string  `pulumi:"field"`
}





type AlertRuleLeafConditionInput interface {
	pulumi.Input

	ToAlertRuleLeafConditionOutput() AlertRuleLeafConditionOutput
	ToAlertRuleLeafConditionOutputWithContext(context.Context) AlertRuleLeafConditionOutput
}

type AlertRuleLeafConditionArgs struct {
	ContainsAny pulumi.StringArrayInput `pulumi:"containsAny"`
	Equals      pulumi.StringPtrInput   `pulumi:"equals"`
	Field       pulumi.StringPtrInput   `pulumi:"field"`
}

func (AlertRuleLeafConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleLeafCondition)(nil)).Elem()
}

func (i AlertRuleLeafConditionArgs) ToAlertRuleLeafConditionOutput() AlertRuleLeafConditionOutput {
	return i.ToAlertRuleLeafConditionOutputWithContext(context.Background())
}

func (i AlertRuleLeafConditionArgs) ToAlertRuleLeafConditionOutputWithContext(ctx context.Context) AlertRuleLeafConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleLeafConditionOutput)
}





type AlertRuleLeafConditionArrayInput interface {
	pulumi.Input

	ToAlertRuleLeafConditionArrayOutput() AlertRuleLeafConditionArrayOutput
	ToAlertRuleLeafConditionArrayOutputWithContext(context.Context) AlertRuleLeafConditionArrayOutput
}

type AlertRuleLeafConditionArray []AlertRuleLeafConditionInput

func (AlertRuleLeafConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertRuleLeafCondition)(nil)).Elem()
}

func (i AlertRuleLeafConditionArray) ToAlertRuleLeafConditionArrayOutput() AlertRuleLeafConditionArrayOutput {
	return i.ToAlertRuleLeafConditionArrayOutputWithContext(context.Background())
}

func (i AlertRuleLeafConditionArray) ToAlertRuleLeafConditionArrayOutputWithContext(ctx context.Context) AlertRuleLeafConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleLeafConditionArrayOutput)
}

type AlertRuleLeafConditionOutput struct{ *pulumi.OutputState }

func (AlertRuleLeafConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleLeafCondition)(nil)).Elem()
}

func (o AlertRuleLeafConditionOutput) ToAlertRuleLeafConditionOutput() AlertRuleLeafConditionOutput {
	return o
}

func (o AlertRuleLeafConditionOutput) ToAlertRuleLeafConditionOutputWithContext(ctx context.Context) AlertRuleLeafConditionOutput {
	return o
}

func (o AlertRuleLeafConditionOutput) ContainsAny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertRuleLeafCondition) []string { return v.ContainsAny }).(pulumi.StringArrayOutput)
}

func (o AlertRuleLeafConditionOutput) Equals() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertRuleLeafCondition) *string { return v.Equals }).(pulumi.StringPtrOutput)
}

func (o AlertRuleLeafConditionOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertRuleLeafCondition) *string { return v.Field }).(pulumi.StringPtrOutput)
}

type AlertRuleLeafConditionArrayOutput struct{ *pulumi.OutputState }

func (AlertRuleLeafConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertRuleLeafCondition)(nil)).Elem()
}

func (o AlertRuleLeafConditionArrayOutput) ToAlertRuleLeafConditionArrayOutput() AlertRuleLeafConditionArrayOutput {
	return o
}

func (o AlertRuleLeafConditionArrayOutput) ToAlertRuleLeafConditionArrayOutputWithContext(ctx context.Context) AlertRuleLeafConditionArrayOutput {
	return o
}

func (o AlertRuleLeafConditionArrayOutput) Index(i pulumi.IntInput) AlertRuleLeafConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertRuleLeafCondition {
		return vs[0].([]AlertRuleLeafCondition)[vs[1].(int)]
	}).(AlertRuleLeafConditionOutput)
}

type AlertRuleLeafConditionResponse struct {
	ContainsAny []string `pulumi:"containsAny"`
	Equals      *string  `pulumi:"equals"`
	Field       *string  `pulumi:"field"`
}

type AlertRuleLeafConditionResponseOutput struct{ *pulumi.OutputState }

func (AlertRuleLeafConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleLeafConditionResponse)(nil)).Elem()
}

func (o AlertRuleLeafConditionResponseOutput) ToAlertRuleLeafConditionResponseOutput() AlertRuleLeafConditionResponseOutput {
	return o
}

func (o AlertRuleLeafConditionResponseOutput) ToAlertRuleLeafConditionResponseOutputWithContext(ctx context.Context) AlertRuleLeafConditionResponseOutput {
	return o
}

func (o AlertRuleLeafConditionResponseOutput) ContainsAny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertRuleLeafConditionResponse) []string { return v.ContainsAny }).(pulumi.StringArrayOutput)
}

func (o AlertRuleLeafConditionResponseOutput) Equals() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertRuleLeafConditionResponse) *string { return v.Equals }).(pulumi.StringPtrOutput)
}

func (o AlertRuleLeafConditionResponseOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertRuleLeafConditionResponse) *string { return v.Field }).(pulumi.StringPtrOutput)
}

type AlertRuleLeafConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (AlertRuleLeafConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertRuleLeafConditionResponse)(nil)).Elem()
}

func (o AlertRuleLeafConditionResponseArrayOutput) ToAlertRuleLeafConditionResponseArrayOutput() AlertRuleLeafConditionResponseArrayOutput {
	return o
}

func (o AlertRuleLeafConditionResponseArrayOutput) ToAlertRuleLeafConditionResponseArrayOutputWithContext(ctx context.Context) AlertRuleLeafConditionResponseArrayOutput {
	return o
}

func (o AlertRuleLeafConditionResponseArrayOutput) Index(i pulumi.IntInput) AlertRuleLeafConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertRuleLeafConditionResponse {
		return vs[0].([]AlertRuleLeafConditionResponse)[vs[1].(int)]
	}).(AlertRuleLeafConditionResponseOutput)
}

type AlertingAction struct {
	AznsAction      *AzNsActionGroup `pulumi:"aznsAction"`
	OdataType       string           `pulumi:"odataType"`
	Severity        string           `pulumi:"severity"`
	ThrottlingInMin *int             `pulumi:"throttlingInMin"`
	Trigger         TriggerCondition `pulumi:"trigger"`
}

type AlertingActionResponse struct {
	AznsAction      *AzNsActionGroupResponse `pulumi:"aznsAction"`
	OdataType       string                   `pulumi:"odataType"`
	Severity        string                   `pulumi:"severity"`
	ThrottlingInMin *int                     `pulumi:"throttlingInMin"`
	Trigger         TriggerConditionResponse `pulumi:"trigger"`
}

type ApplicationInsightsComponentAnalyticsItemProperties struct {
	FunctionAlias *string `pulumi:"functionAlias"`
}





type ApplicationInsightsComponentAnalyticsItemPropertiesInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesOutput() ApplicationInsightsComponentAnalyticsItemPropertiesOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesOutput
}

type ApplicationInsightsComponentAnalyticsItemPropertiesArgs struct {
	FunctionAlias pulumi.StringPtrInput `pulumi:"functionAlias"`
}

func (ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesOutput() ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesOutput)
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesOutput).ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput
}

type applicationInsightsComponentAnalyticsItemPropertiesPtrType ApplicationInsightsComponentAnalyticsItemPropertiesArgs

func ApplicationInsightsComponentAnalyticsItemPropertiesPtr(v *ApplicationInsightsComponentAnalyticsItemPropertiesArgs) ApplicationInsightsComponentAnalyticsItemPropertiesPtrInput {
	return (*applicationInsightsComponentAnalyticsItemPropertiesPtrType)(v)
}

func (*applicationInsightsComponentAnalyticsItemPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesPtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesPtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesOutput() ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return o.ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentAnalyticsItemProperties) *ApplicationInsightsComponentAnalyticsItemProperties {
		return &v
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput)
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentAnalyticsItemProperties) *string { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemProperties)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) Elem() ApplicationInsightsComponentAnalyticsItemPropertiesOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemProperties) ApplicationInsightsComponentAnalyticsItemProperties {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentAnalyticsItemProperties
		return ret
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesOutput)
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemProperties) *string {
		if v == nil {
			return nil
		}
		return v.FunctionAlias
	}).(pulumi.StringPtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesResponse struct {
	FunctionAlias *string `pulumi:"functionAlias"`
}

type ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *string { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCap struct {
	Cap                                  *float64 `pulumi:"cap"`
	StopSendNotificationWhenHitCap       *bool    `pulumi:"stopSendNotificationWhenHitCap"`
	StopSendNotificationWhenHitThreshold *bool    `pulumi:"stopSendNotificationWhenHitThreshold"`
	WarningThreshold                     *int     `pulumi:"warningThreshold"`
}





type ApplicationInsightsComponentDataVolumeCapInput interface {
	pulumi.Input

	ToApplicationInsightsComponentDataVolumeCapOutput() ApplicationInsightsComponentDataVolumeCapOutput
	ToApplicationInsightsComponentDataVolumeCapOutputWithContext(context.Context) ApplicationInsightsComponentDataVolumeCapOutput
}

type ApplicationInsightsComponentDataVolumeCapArgs struct {
	Cap                                  pulumi.Float64PtrInput `pulumi:"cap"`
	StopSendNotificationWhenHitCap       pulumi.BoolPtrInput    `pulumi:"stopSendNotificationWhenHitCap"`
	StopSendNotificationWhenHitThreshold pulumi.BoolPtrInput    `pulumi:"stopSendNotificationWhenHitThreshold"`
	WarningThreshold                     pulumi.IntPtrInput     `pulumi:"warningThreshold"`
}

func (ApplicationInsightsComponentDataVolumeCapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentDataVolumeCap)(nil)).Elem()
}

func (i ApplicationInsightsComponentDataVolumeCapArgs) ToApplicationInsightsComponentDataVolumeCapOutput() ApplicationInsightsComponentDataVolumeCapOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentDataVolumeCapArgs) ToApplicationInsightsComponentDataVolumeCapOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapOutput)
}

func (i ApplicationInsightsComponentDataVolumeCapArgs) ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentDataVolumeCapArgs) ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapOutput).ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentDataVolumeCapPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput
	ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput
}

type applicationInsightsComponentDataVolumeCapPtrType ApplicationInsightsComponentDataVolumeCapArgs

func ApplicationInsightsComponentDataVolumeCapPtr(v *ApplicationInsightsComponentDataVolumeCapArgs) ApplicationInsightsComponentDataVolumeCapPtrInput {
	return (*applicationInsightsComponentDataVolumeCapPtrType)(v)
}

func (*applicationInsightsComponentDataVolumeCapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentDataVolumeCap)(nil)).Elem()
}

func (i *applicationInsightsComponentDataVolumeCapPtrType) ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentDataVolumeCapPtrType) ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCapOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentDataVolumeCapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentDataVolumeCap)(nil)).Elem()
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) ToApplicationInsightsComponentDataVolumeCapOutput() ApplicationInsightsComponentDataVolumeCapOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) ToApplicationInsightsComponentDataVolumeCapOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return o.ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentDataVolumeCap) *ApplicationInsightsComponentDataVolumeCap {
		return &v
	}).(ApplicationInsightsComponentDataVolumeCapPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) Cap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCap) *float64 { return v.Cap }).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) StopSendNotificationWhenHitCap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCap) *bool { return v.StopSendNotificationWhenHitCap }).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) StopSendNotificationWhenHitThreshold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCap) *bool { return v.StopSendNotificationWhenHitThreshold }).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapOutput) WarningThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCap) *int { return v.WarningThreshold }).(pulumi.IntPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCapPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentDataVolumeCapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentDataVolumeCap)(nil)).Elem()
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) ToApplicationInsightsComponentDataVolumeCapPtrOutput() ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) ToApplicationInsightsComponentDataVolumeCapPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapPtrOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) Elem() ApplicationInsightsComponentDataVolumeCapOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) ApplicationInsightsComponentDataVolumeCap {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentDataVolumeCap
		return ret
	}).(ApplicationInsightsComponentDataVolumeCapOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) Cap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) *float64 {
		if v == nil {
			return nil
		}
		return v.Cap
	}).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) StopSendNotificationWhenHitCap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) *bool {
		if v == nil {
			return nil
		}
		return v.StopSendNotificationWhenHitCap
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) StopSendNotificationWhenHitThreshold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) *bool {
		if v == nil {
			return nil
		}
		return v.StopSendNotificationWhenHitThreshold
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapPtrOutput) WarningThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCap) *int {
		if v == nil {
			return nil
		}
		return v.WarningThreshold
	}).(pulumi.IntPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCapResponse struct {
	Cap                                  *float64 `pulumi:"cap"`
	MaxHistoryCap                        float64  `pulumi:"maxHistoryCap"`
	ResetTime                            int      `pulumi:"resetTime"`
	StopSendNotificationWhenHitCap       *bool    `pulumi:"stopSendNotificationWhenHitCap"`
	StopSendNotificationWhenHitThreshold *bool    `pulumi:"stopSendNotificationWhenHitThreshold"`
	WarningThreshold                     *int     `pulumi:"warningThreshold"`
}

type ApplicationInsightsComponentDataVolumeCapResponseOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentDataVolumeCapResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentDataVolumeCapResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ToApplicationInsightsComponentDataVolumeCapResponseOutput() ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ToApplicationInsightsComponentDataVolumeCapResponseOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) Cap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) *float64 { return v.Cap }).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) MaxHistoryCap() pulumi.Float64Output {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) float64 { return v.MaxHistoryCap }).(pulumi.Float64Output)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ResetTime() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) int { return v.ResetTime }).(pulumi.IntOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) StopSendNotificationWhenHitCap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) *bool {
		return v.StopSendNotificationWhenHitCap
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) StopSendNotificationWhenHitThreshold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) *bool {
		return v.StopSendNotificationWhenHitThreshold
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) WarningThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentDataVolumeCapResponse) *int { return v.WarningThreshold }).(pulumi.IntPtrOutput)
}

type ApplicationInsightsComponentDataVolumeCapResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentDataVolumeCapResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) Elem() ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) ApplicationInsightsComponentDataVolumeCapResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentDataVolumeCapResponse
		return ret
	}).(ApplicationInsightsComponentDataVolumeCapResponseOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) Cap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Cap
	}).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) MaxHistoryCap() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MaxHistoryCap
	}).(pulumi.Float64PtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) ResetTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ResetTime
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) StopSendNotificationWhenHitCap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *bool {
		if v == nil {
			return nil
		}
		return v.StopSendNotificationWhenHitCap
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) StopSendNotificationWhenHitThreshold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *bool {
		if v == nil {
			return nil
		}
		return v.StopSendNotificationWhenHitThreshold
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentDataVolumeCapResponsePtrOutput) WarningThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentDataVolumeCapResponse) *int {
		if v == nil {
			return nil
		}
		return v.WarningThreshold
	}).(pulumi.IntPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions struct {
	Description                *string `pulumi:"description"`
	DisplayName                *string `pulumi:"displayName"`
	HelpUrl                    *string `pulumi:"helpUrl"`
	IsEnabledByDefault         *bool   `pulumi:"isEnabledByDefault"`
	IsHidden                   *bool   `pulumi:"isHidden"`
	IsInPreview                *bool   `pulumi:"isInPreview"`
	Name                       *string `pulumi:"name"`
	SupportsEmailNotifications *bool   `pulumi:"supportsEmailNotifications"`
}

type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) Elem() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions
		return ret
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions struct {
	Description                *string `pulumi:"description"`
	DisplayName                *string `pulumi:"displayName"`
	HelpUrl                    *string `pulumi:"helpUrl"`
	IsEnabledByDefault         *bool   `pulumi:"isEnabledByDefault"`
	IsHidden                   *bool   `pulumi:"isHidden"`
	IsInPreview                *bool   `pulumi:"isInPreview"`
	Name                       *string `pulumi:"name"`
	SupportsEmailNotifications *bool   `pulumi:"supportsEmailNotifications"`
}





type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput
}

type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs struct {
	Description                pulumi.StringPtrInput `pulumi:"description"`
	DisplayName                pulumi.StringPtrInput `pulumi:"displayName"`
	HelpUrl                    pulumi.StringPtrInput `pulumi:"helpUrl"`
	IsEnabledByDefault         pulumi.BoolPtrInput   `pulumi:"isEnabledByDefault"`
	IsHidden                   pulumi.BoolPtrInput   `pulumi:"isHidden"`
	IsInPreview                pulumi.BoolPtrInput   `pulumi:"isInPreview"`
	Name                       pulumi.StringPtrInput `pulumi:"name"`
	SupportsEmailNotifications pulumi.BoolPtrInput   `pulumi:"supportsEmailNotifications"`
}

func (ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions)(nil)).Elem()
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput)
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput).ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput
}

type applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs

func ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtr(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsArgs) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrInput {
	return (*applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType)(v)
}

func (*applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions)(nil)).Elem()
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return o.ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions {
		return &v
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) Elem() ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions
		return ret
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type ArmRoleReceiver struct {
	Name                 string `pulumi:"name"`
	RoleId               string `pulumi:"roleId"`
	UseCommonAlertSchema *bool  `pulumi:"useCommonAlertSchema"`
}


func (val *ArmRoleReceiver) Defaults() *ArmRoleReceiver {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}





type ArmRoleReceiverInput interface {
	pulumi.Input

	ToArmRoleReceiverOutput() ArmRoleReceiverOutput
	ToArmRoleReceiverOutputWithContext(context.Context) ArmRoleReceiverOutput
}

type ArmRoleReceiverArgs struct {
	Name                 pulumi.StringInput  `pulumi:"name"`
	RoleId               pulumi.StringInput  `pulumi:"roleId"`
	UseCommonAlertSchema pulumi.BoolPtrInput `pulumi:"useCommonAlertSchema"`
}


func (val *ArmRoleReceiverArgs) Defaults() *ArmRoleReceiverArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		tmp.UseCommonAlertSchema = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (ArmRoleReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmRoleReceiver)(nil)).Elem()
}

func (i ArmRoleReceiverArgs) ToArmRoleReceiverOutput() ArmRoleReceiverOutput {
	return i.ToArmRoleReceiverOutputWithContext(context.Background())
}

func (i ArmRoleReceiverArgs) ToArmRoleReceiverOutputWithContext(ctx context.Context) ArmRoleReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRoleReceiverOutput)
}





type ArmRoleReceiverArrayInput interface {
	pulumi.Input

	ToArmRoleReceiverArrayOutput() ArmRoleReceiverArrayOutput
	ToArmRoleReceiverArrayOutputWithContext(context.Context) ArmRoleReceiverArrayOutput
}

type ArmRoleReceiverArray []ArmRoleReceiverInput

func (ArmRoleReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmRoleReceiver)(nil)).Elem()
}

func (i ArmRoleReceiverArray) ToArmRoleReceiverArrayOutput() ArmRoleReceiverArrayOutput {
	return i.ToArmRoleReceiverArrayOutputWithContext(context.Background())
}

func (i ArmRoleReceiverArray) ToArmRoleReceiverArrayOutputWithContext(ctx context.Context) ArmRoleReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRoleReceiverArrayOutput)
}

type ArmRoleReceiverOutput struct{ *pulumi.OutputState }

func (ArmRoleReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmRoleReceiver)(nil)).Elem()
}

func (o ArmRoleReceiverOutput) ToArmRoleReceiverOutput() ArmRoleReceiverOutput {
	return o
}

func (o ArmRoleReceiverOutput) ToArmRoleReceiverOutputWithContext(ctx context.Context) ArmRoleReceiverOutput {
	return o
}

func (o ArmRoleReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ArmRoleReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o ArmRoleReceiverOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmRoleReceiver) string { return v.RoleId }).(pulumi.StringOutput)
}

func (o ArmRoleReceiverOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ArmRoleReceiver) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type ArmRoleReceiverArrayOutput struct{ *pulumi.OutputState }

func (ArmRoleReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmRoleReceiver)(nil)).Elem()
}

func (o ArmRoleReceiverArrayOutput) ToArmRoleReceiverArrayOutput() ArmRoleReceiverArrayOutput {
	return o
}

func (o ArmRoleReceiverArrayOutput) ToArmRoleReceiverArrayOutputWithContext(ctx context.Context) ArmRoleReceiverArrayOutput {
	return o
}

func (o ArmRoleReceiverArrayOutput) Index(i pulumi.IntInput) ArmRoleReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArmRoleReceiver {
		return vs[0].([]ArmRoleReceiver)[vs[1].(int)]
	}).(ArmRoleReceiverOutput)
}

type ArmRoleReceiverResponse struct {
	Name                 string `pulumi:"name"`
	RoleId               string `pulumi:"roleId"`
	UseCommonAlertSchema *bool  `pulumi:"useCommonAlertSchema"`
}


func (val *ArmRoleReceiverResponse) Defaults() *ArmRoleReceiverResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}

type ArmRoleReceiverResponseOutput struct{ *pulumi.OutputState }

func (ArmRoleReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmRoleReceiverResponse)(nil)).Elem()
}

func (o ArmRoleReceiverResponseOutput) ToArmRoleReceiverResponseOutput() ArmRoleReceiverResponseOutput {
	return o
}

func (o ArmRoleReceiverResponseOutput) ToArmRoleReceiverResponseOutputWithContext(ctx context.Context) ArmRoleReceiverResponseOutput {
	return o
}

func (o ArmRoleReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ArmRoleReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ArmRoleReceiverResponseOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmRoleReceiverResponse) string { return v.RoleId }).(pulumi.StringOutput)
}

func (o ArmRoleReceiverResponseOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ArmRoleReceiverResponse) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type ArmRoleReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (ArmRoleReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmRoleReceiverResponse)(nil)).Elem()
}

func (o ArmRoleReceiverResponseArrayOutput) ToArmRoleReceiverResponseArrayOutput() ArmRoleReceiverResponseArrayOutput {
	return o
}

func (o ArmRoleReceiverResponseArrayOutput) ToArmRoleReceiverResponseArrayOutputWithContext(ctx context.Context) ArmRoleReceiverResponseArrayOutput {
	return o
}

func (o ArmRoleReceiverResponseArrayOutput) Index(i pulumi.IntInput) ArmRoleReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArmRoleReceiverResponse {
		return vs[0].([]ArmRoleReceiverResponse)[vs[1].(int)]
	}).(ArmRoleReceiverResponseOutput)
}

type AutomationRunbookReceiver struct {
	AutomationAccountId  string  `pulumi:"automationAccountId"`
	IsGlobalRunbook      bool    `pulumi:"isGlobalRunbook"`
	Name                 *string `pulumi:"name"`
	RunbookName          string  `pulumi:"runbookName"`
	ServiceUri           *string `pulumi:"serviceUri"`
	UseCommonAlertSchema *bool   `pulumi:"useCommonAlertSchema"`
	WebhookResourceId    string  `pulumi:"webhookResourceId"`
}


func (val *AutomationRunbookReceiver) Defaults() *AutomationRunbookReceiver {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}





type AutomationRunbookReceiverInput interface {
	pulumi.Input

	ToAutomationRunbookReceiverOutput() AutomationRunbookReceiverOutput
	ToAutomationRunbookReceiverOutputWithContext(context.Context) AutomationRunbookReceiverOutput
}

type AutomationRunbookReceiverArgs struct {
	AutomationAccountId  pulumi.StringInput    `pulumi:"automationAccountId"`
	IsGlobalRunbook      pulumi.BoolInput      `pulumi:"isGlobalRunbook"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
	RunbookName          pulumi.StringInput    `pulumi:"runbookName"`
	ServiceUri           pulumi.StringPtrInput `pulumi:"serviceUri"`
	UseCommonAlertSchema pulumi.BoolPtrInput   `pulumi:"useCommonAlertSchema"`
	WebhookResourceId    pulumi.StringInput    `pulumi:"webhookResourceId"`
}


func (val *AutomationRunbookReceiverArgs) Defaults() *AutomationRunbookReceiverArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		tmp.UseCommonAlertSchema = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (AutomationRunbookReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRunbookReceiver)(nil)).Elem()
}

func (i AutomationRunbookReceiverArgs) ToAutomationRunbookReceiverOutput() AutomationRunbookReceiverOutput {
	return i.ToAutomationRunbookReceiverOutputWithContext(context.Background())
}

func (i AutomationRunbookReceiverArgs) ToAutomationRunbookReceiverOutputWithContext(ctx context.Context) AutomationRunbookReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRunbookReceiverOutput)
}





type AutomationRunbookReceiverArrayInput interface {
	pulumi.Input

	ToAutomationRunbookReceiverArrayOutput() AutomationRunbookReceiverArrayOutput
	ToAutomationRunbookReceiverArrayOutputWithContext(context.Context) AutomationRunbookReceiverArrayOutput
}

type AutomationRunbookReceiverArray []AutomationRunbookReceiverInput

func (AutomationRunbookReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRunbookReceiver)(nil)).Elem()
}

func (i AutomationRunbookReceiverArray) ToAutomationRunbookReceiverArrayOutput() AutomationRunbookReceiverArrayOutput {
	return i.ToAutomationRunbookReceiverArrayOutputWithContext(context.Background())
}

func (i AutomationRunbookReceiverArray) ToAutomationRunbookReceiverArrayOutputWithContext(ctx context.Context) AutomationRunbookReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRunbookReceiverArrayOutput)
}

type AutomationRunbookReceiverOutput struct{ *pulumi.OutputState }

func (AutomationRunbookReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRunbookReceiver)(nil)).Elem()
}

func (o AutomationRunbookReceiverOutput) ToAutomationRunbookReceiverOutput() AutomationRunbookReceiverOutput {
	return o
}

func (o AutomationRunbookReceiverOutput) ToAutomationRunbookReceiverOutputWithContext(ctx context.Context) AutomationRunbookReceiverOutput {
	return o
}

func (o AutomationRunbookReceiverOutput) AutomationAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRunbookReceiver) string { return v.AutomationAccountId }).(pulumi.StringOutput)
}

func (o AutomationRunbookReceiverOutput) IsGlobalRunbook() pulumi.BoolOutput {
	return o.ApplyT(func(v AutomationRunbookReceiver) bool { return v.IsGlobalRunbook }).(pulumi.BoolOutput)
}

func (o AutomationRunbookReceiverOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRunbookReceiver) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AutomationRunbookReceiverOutput) RunbookName() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRunbookReceiver) string { return v.RunbookName }).(pulumi.StringOutput)
}

func (o AutomationRunbookReceiverOutput) ServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRunbookReceiver) *string { return v.ServiceUri }).(pulumi.StringPtrOutput)
}

func (o AutomationRunbookReceiverOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomationRunbookReceiver) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

func (o AutomationRunbookReceiverOutput) WebhookResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRunbookReceiver) string { return v.WebhookResourceId }).(pulumi.StringOutput)
}

type AutomationRunbookReceiverArrayOutput struct{ *pulumi.OutputState }

func (AutomationRunbookReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRunbookReceiver)(nil)).Elem()
}

func (o AutomationRunbookReceiverArrayOutput) ToAutomationRunbookReceiverArrayOutput() AutomationRunbookReceiverArrayOutput {
	return o
}

func (o AutomationRunbookReceiverArrayOutput) ToAutomationRunbookReceiverArrayOutputWithContext(ctx context.Context) AutomationRunbookReceiverArrayOutput {
	return o
}

func (o AutomationRunbookReceiverArrayOutput) Index(i pulumi.IntInput) AutomationRunbookReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationRunbookReceiver {
		return vs[0].([]AutomationRunbookReceiver)[vs[1].(int)]
	}).(AutomationRunbookReceiverOutput)
}

type AutomationRunbookReceiverResponse struct {
	AutomationAccountId  string  `pulumi:"automationAccountId"`
	IsGlobalRunbook      bool    `pulumi:"isGlobalRunbook"`
	Name                 *string `pulumi:"name"`
	RunbookName          string  `pulumi:"runbookName"`
	ServiceUri           *string `pulumi:"serviceUri"`
	UseCommonAlertSchema *bool   `pulumi:"useCommonAlertSchema"`
	WebhookResourceId    string  `pulumi:"webhookResourceId"`
}


func (val *AutomationRunbookReceiverResponse) Defaults() *AutomationRunbookReceiverResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}

type AutomationRunbookReceiverResponseOutput struct{ *pulumi.OutputState }

func (AutomationRunbookReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRunbookReceiverResponse)(nil)).Elem()
}

func (o AutomationRunbookReceiverResponseOutput) ToAutomationRunbookReceiverResponseOutput() AutomationRunbookReceiverResponseOutput {
	return o
}

func (o AutomationRunbookReceiverResponseOutput) ToAutomationRunbookReceiverResponseOutputWithContext(ctx context.Context) AutomationRunbookReceiverResponseOutput {
	return o
}

func (o AutomationRunbookReceiverResponseOutput) AutomationAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRunbookReceiverResponse) string { return v.AutomationAccountId }).(pulumi.StringOutput)
}

func (o AutomationRunbookReceiverResponseOutput) IsGlobalRunbook() pulumi.BoolOutput {
	return o.ApplyT(func(v AutomationRunbookReceiverResponse) bool { return v.IsGlobalRunbook }).(pulumi.BoolOutput)
}

func (o AutomationRunbookReceiverResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRunbookReceiverResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AutomationRunbookReceiverResponseOutput) RunbookName() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRunbookReceiverResponse) string { return v.RunbookName }).(pulumi.StringOutput)
}

func (o AutomationRunbookReceiverResponseOutput) ServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRunbookReceiverResponse) *string { return v.ServiceUri }).(pulumi.StringPtrOutput)
}

func (o AutomationRunbookReceiverResponseOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomationRunbookReceiverResponse) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

func (o AutomationRunbookReceiverResponseOutput) WebhookResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRunbookReceiverResponse) string { return v.WebhookResourceId }).(pulumi.StringOutput)
}

type AutomationRunbookReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationRunbookReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRunbookReceiverResponse)(nil)).Elem()
}

func (o AutomationRunbookReceiverResponseArrayOutput) ToAutomationRunbookReceiverResponseArrayOutput() AutomationRunbookReceiverResponseArrayOutput {
	return o
}

func (o AutomationRunbookReceiverResponseArrayOutput) ToAutomationRunbookReceiverResponseArrayOutputWithContext(ctx context.Context) AutomationRunbookReceiverResponseArrayOutput {
	return o
}

func (o AutomationRunbookReceiverResponseArrayOutput) Index(i pulumi.IntInput) AutomationRunbookReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationRunbookReceiverResponse {
		return vs[0].([]AutomationRunbookReceiverResponse)[vs[1].(int)]
	}).(AutomationRunbookReceiverResponseOutput)
}

type AutoscaleNotification struct {
	Email     *EmailNotification    `pulumi:"email"`
	Operation OperationType         `pulumi:"operation"`
	Webhooks  []WebhookNotification `pulumi:"webhooks"`
}


func (val *AutoscaleNotification) Defaults() *AutoscaleNotification {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Email = tmp.Email.Defaults()

	return &tmp
}





type AutoscaleNotificationInput interface {
	pulumi.Input

	ToAutoscaleNotificationOutput() AutoscaleNotificationOutput
	ToAutoscaleNotificationOutputWithContext(context.Context) AutoscaleNotificationOutput
}

type AutoscaleNotificationArgs struct {
	Email     EmailNotificationPtrInput     `pulumi:"email"`
	Operation OperationTypeInput            `pulumi:"operation"`
	Webhooks  WebhookNotificationArrayInput `pulumi:"webhooks"`
}


func (val *AutoscaleNotificationArgs) Defaults() *AutoscaleNotificationArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (AutoscaleNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleNotification)(nil)).Elem()
}

func (i AutoscaleNotificationArgs) ToAutoscaleNotificationOutput() AutoscaleNotificationOutput {
	return i.ToAutoscaleNotificationOutputWithContext(context.Background())
}

func (i AutoscaleNotificationArgs) ToAutoscaleNotificationOutputWithContext(ctx context.Context) AutoscaleNotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleNotificationOutput)
}





type AutoscaleNotificationArrayInput interface {
	pulumi.Input

	ToAutoscaleNotificationArrayOutput() AutoscaleNotificationArrayOutput
	ToAutoscaleNotificationArrayOutputWithContext(context.Context) AutoscaleNotificationArrayOutput
}

type AutoscaleNotificationArray []AutoscaleNotificationInput

func (AutoscaleNotificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleNotification)(nil)).Elem()
}

func (i AutoscaleNotificationArray) ToAutoscaleNotificationArrayOutput() AutoscaleNotificationArrayOutput {
	return i.ToAutoscaleNotificationArrayOutputWithContext(context.Background())
}

func (i AutoscaleNotificationArray) ToAutoscaleNotificationArrayOutputWithContext(ctx context.Context) AutoscaleNotificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleNotificationArrayOutput)
}

type AutoscaleNotificationOutput struct{ *pulumi.OutputState }

func (AutoscaleNotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleNotification)(nil)).Elem()
}

func (o AutoscaleNotificationOutput) ToAutoscaleNotificationOutput() AutoscaleNotificationOutput {
	return o
}

func (o AutoscaleNotificationOutput) ToAutoscaleNotificationOutputWithContext(ctx context.Context) AutoscaleNotificationOutput {
	return o
}

func (o AutoscaleNotificationOutput) Email() EmailNotificationPtrOutput {
	return o.ApplyT(func(v AutoscaleNotification) *EmailNotification { return v.Email }).(EmailNotificationPtrOutput)
}

func (o AutoscaleNotificationOutput) Operation() OperationTypeOutput {
	return o.ApplyT(func(v AutoscaleNotification) OperationType { return v.Operation }).(OperationTypeOutput)
}

func (o AutoscaleNotificationOutput) Webhooks() WebhookNotificationArrayOutput {
	return o.ApplyT(func(v AutoscaleNotification) []WebhookNotification { return v.Webhooks }).(WebhookNotificationArrayOutput)
}

type AutoscaleNotificationArrayOutput struct{ *pulumi.OutputState }

func (AutoscaleNotificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleNotification)(nil)).Elem()
}

func (o AutoscaleNotificationArrayOutput) ToAutoscaleNotificationArrayOutput() AutoscaleNotificationArrayOutput {
	return o
}

func (o AutoscaleNotificationArrayOutput) ToAutoscaleNotificationArrayOutputWithContext(ctx context.Context) AutoscaleNotificationArrayOutput {
	return o
}

func (o AutoscaleNotificationArrayOutput) Index(i pulumi.IntInput) AutoscaleNotificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoscaleNotification {
		return vs[0].([]AutoscaleNotification)[vs[1].(int)]
	}).(AutoscaleNotificationOutput)
}

type AutoscaleNotificationResponse struct {
	Email     *EmailNotificationResponse    `pulumi:"email"`
	Operation string                        `pulumi:"operation"`
	Webhooks  []WebhookNotificationResponse `pulumi:"webhooks"`
}


func (val *AutoscaleNotificationResponse) Defaults() *AutoscaleNotificationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Email = tmp.Email.Defaults()

	return &tmp
}

type AutoscaleNotificationResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleNotificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleNotificationResponse)(nil)).Elem()
}

func (o AutoscaleNotificationResponseOutput) ToAutoscaleNotificationResponseOutput() AutoscaleNotificationResponseOutput {
	return o
}

func (o AutoscaleNotificationResponseOutput) ToAutoscaleNotificationResponseOutputWithContext(ctx context.Context) AutoscaleNotificationResponseOutput {
	return o
}

func (o AutoscaleNotificationResponseOutput) Email() EmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v AutoscaleNotificationResponse) *EmailNotificationResponse { return v.Email }).(EmailNotificationResponsePtrOutput)
}

func (o AutoscaleNotificationResponseOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v AutoscaleNotificationResponse) string { return v.Operation }).(pulumi.StringOutput)
}

func (o AutoscaleNotificationResponseOutput) Webhooks() WebhookNotificationResponseArrayOutput {
	return o.ApplyT(func(v AutoscaleNotificationResponse) []WebhookNotificationResponse { return v.Webhooks }).(WebhookNotificationResponseArrayOutput)
}

type AutoscaleNotificationResponseArrayOutput struct{ *pulumi.OutputState }

func (AutoscaleNotificationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleNotificationResponse)(nil)).Elem()
}

func (o AutoscaleNotificationResponseArrayOutput) ToAutoscaleNotificationResponseArrayOutput() AutoscaleNotificationResponseArrayOutput {
	return o
}

func (o AutoscaleNotificationResponseArrayOutput) ToAutoscaleNotificationResponseArrayOutputWithContext(ctx context.Context) AutoscaleNotificationResponseArrayOutput {
	return o
}

func (o AutoscaleNotificationResponseArrayOutput) Index(i pulumi.IntInput) AutoscaleNotificationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoscaleNotificationResponse {
		return vs[0].([]AutoscaleNotificationResponse)[vs[1].(int)]
	}).(AutoscaleNotificationResponseOutput)
}

type AutoscaleProfile struct {
	Capacity   ScaleCapacity `pulumi:"capacity"`
	FixedDate  *TimeWindow   `pulumi:"fixedDate"`
	Name       string        `pulumi:"name"`
	Recurrence *Recurrence   `pulumi:"recurrence"`
	Rules      []ScaleRule   `pulumi:"rules"`
}





type AutoscaleProfileInput interface {
	pulumi.Input

	ToAutoscaleProfileOutput() AutoscaleProfileOutput
	ToAutoscaleProfileOutputWithContext(context.Context) AutoscaleProfileOutput
}

type AutoscaleProfileArgs struct {
	Capacity   ScaleCapacityInput  `pulumi:"capacity"`
	FixedDate  TimeWindowPtrInput  `pulumi:"fixedDate"`
	Name       pulumi.StringInput  `pulumi:"name"`
	Recurrence RecurrencePtrInput  `pulumi:"recurrence"`
	Rules      ScaleRuleArrayInput `pulumi:"rules"`
}

func (AutoscaleProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleProfile)(nil)).Elem()
}

func (i AutoscaleProfileArgs) ToAutoscaleProfileOutput() AutoscaleProfileOutput {
	return i.ToAutoscaleProfileOutputWithContext(context.Background())
}

func (i AutoscaleProfileArgs) ToAutoscaleProfileOutputWithContext(ctx context.Context) AutoscaleProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleProfileOutput)
}





type AutoscaleProfileArrayInput interface {
	pulumi.Input

	ToAutoscaleProfileArrayOutput() AutoscaleProfileArrayOutput
	ToAutoscaleProfileArrayOutputWithContext(context.Context) AutoscaleProfileArrayOutput
}

type AutoscaleProfileArray []AutoscaleProfileInput

func (AutoscaleProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleProfile)(nil)).Elem()
}

func (i AutoscaleProfileArray) ToAutoscaleProfileArrayOutput() AutoscaleProfileArrayOutput {
	return i.ToAutoscaleProfileArrayOutputWithContext(context.Background())
}

func (i AutoscaleProfileArray) ToAutoscaleProfileArrayOutputWithContext(ctx context.Context) AutoscaleProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleProfileArrayOutput)
}

type AutoscaleProfileOutput struct{ *pulumi.OutputState }

func (AutoscaleProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleProfile)(nil)).Elem()
}

func (o AutoscaleProfileOutput) ToAutoscaleProfileOutput() AutoscaleProfileOutput {
	return o
}

func (o AutoscaleProfileOutput) ToAutoscaleProfileOutputWithContext(ctx context.Context) AutoscaleProfileOutput {
	return o
}

func (o AutoscaleProfileOutput) Capacity() ScaleCapacityOutput {
	return o.ApplyT(func(v AutoscaleProfile) ScaleCapacity { return v.Capacity }).(ScaleCapacityOutput)
}

func (o AutoscaleProfileOutput) FixedDate() TimeWindowPtrOutput {
	return o.ApplyT(func(v AutoscaleProfile) *TimeWindow { return v.FixedDate }).(TimeWindowPtrOutput)
}

func (o AutoscaleProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoscaleProfile) string { return v.Name }).(pulumi.StringOutput)
}

func (o AutoscaleProfileOutput) Recurrence() RecurrencePtrOutput {
	return o.ApplyT(func(v AutoscaleProfile) *Recurrence { return v.Recurrence }).(RecurrencePtrOutput)
}

func (o AutoscaleProfileOutput) Rules() ScaleRuleArrayOutput {
	return o.ApplyT(func(v AutoscaleProfile) []ScaleRule { return v.Rules }).(ScaleRuleArrayOutput)
}

type AutoscaleProfileArrayOutput struct{ *pulumi.OutputState }

func (AutoscaleProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleProfile)(nil)).Elem()
}

func (o AutoscaleProfileArrayOutput) ToAutoscaleProfileArrayOutput() AutoscaleProfileArrayOutput {
	return o
}

func (o AutoscaleProfileArrayOutput) ToAutoscaleProfileArrayOutputWithContext(ctx context.Context) AutoscaleProfileArrayOutput {
	return o
}

func (o AutoscaleProfileArrayOutput) Index(i pulumi.IntInput) AutoscaleProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoscaleProfile {
		return vs[0].([]AutoscaleProfile)[vs[1].(int)]
	}).(AutoscaleProfileOutput)
}

type AutoscaleProfileResponse struct {
	Capacity   ScaleCapacityResponse `pulumi:"capacity"`
	FixedDate  *TimeWindowResponse   `pulumi:"fixedDate"`
	Name       string                `pulumi:"name"`
	Recurrence *RecurrenceResponse   `pulumi:"recurrence"`
	Rules      []ScaleRuleResponse   `pulumi:"rules"`
}

type AutoscaleProfileResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleProfileResponse)(nil)).Elem()
}

func (o AutoscaleProfileResponseOutput) ToAutoscaleProfileResponseOutput() AutoscaleProfileResponseOutput {
	return o
}

func (o AutoscaleProfileResponseOutput) ToAutoscaleProfileResponseOutputWithContext(ctx context.Context) AutoscaleProfileResponseOutput {
	return o
}

func (o AutoscaleProfileResponseOutput) Capacity() ScaleCapacityResponseOutput {
	return o.ApplyT(func(v AutoscaleProfileResponse) ScaleCapacityResponse { return v.Capacity }).(ScaleCapacityResponseOutput)
}

func (o AutoscaleProfileResponseOutput) FixedDate() TimeWindowResponsePtrOutput {
	return o.ApplyT(func(v AutoscaleProfileResponse) *TimeWindowResponse { return v.FixedDate }).(TimeWindowResponsePtrOutput)
}

func (o AutoscaleProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoscaleProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AutoscaleProfileResponseOutput) Recurrence() RecurrenceResponsePtrOutput {
	return o.ApplyT(func(v AutoscaleProfileResponse) *RecurrenceResponse { return v.Recurrence }).(RecurrenceResponsePtrOutput)
}

func (o AutoscaleProfileResponseOutput) Rules() ScaleRuleResponseArrayOutput {
	return o.ApplyT(func(v AutoscaleProfileResponse) []ScaleRuleResponse { return v.Rules }).(ScaleRuleResponseArrayOutput)
}

type AutoscaleProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (AutoscaleProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleProfileResponse)(nil)).Elem()
}

func (o AutoscaleProfileResponseArrayOutput) ToAutoscaleProfileResponseArrayOutput() AutoscaleProfileResponseArrayOutput {
	return o
}

func (o AutoscaleProfileResponseArrayOutput) ToAutoscaleProfileResponseArrayOutputWithContext(ctx context.Context) AutoscaleProfileResponseArrayOutput {
	return o
}

func (o AutoscaleProfileResponseArrayOutput) Index(i pulumi.IntInput) AutoscaleProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoscaleProfileResponse {
		return vs[0].([]AutoscaleProfileResponse)[vs[1].(int)]
	}).(AutoscaleProfileResponseOutput)
}

type AzNsActionGroup struct {
	ActionGroup          []string `pulumi:"actionGroup"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	EmailSubject         *string  `pulumi:"emailSubject"`
}

type AzNsActionGroupResponse struct {
	ActionGroup          []string `pulumi:"actionGroup"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	EmailSubject         *string  `pulumi:"emailSubject"`
}

type AzureAppPushReceiver struct {
	EmailAddress string `pulumi:"emailAddress"`
	Name         string `pulumi:"name"`
}





type AzureAppPushReceiverInput interface {
	pulumi.Input

	ToAzureAppPushReceiverOutput() AzureAppPushReceiverOutput
	ToAzureAppPushReceiverOutputWithContext(context.Context) AzureAppPushReceiverOutput
}

type AzureAppPushReceiverArgs struct {
	EmailAddress pulumi.StringInput `pulumi:"emailAddress"`
	Name         pulumi.StringInput `pulumi:"name"`
}

func (AzureAppPushReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureAppPushReceiver)(nil)).Elem()
}

func (i AzureAppPushReceiverArgs) ToAzureAppPushReceiverOutput() AzureAppPushReceiverOutput {
	return i.ToAzureAppPushReceiverOutputWithContext(context.Background())
}

func (i AzureAppPushReceiverArgs) ToAzureAppPushReceiverOutputWithContext(ctx context.Context) AzureAppPushReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureAppPushReceiverOutput)
}





type AzureAppPushReceiverArrayInput interface {
	pulumi.Input

	ToAzureAppPushReceiverArrayOutput() AzureAppPushReceiverArrayOutput
	ToAzureAppPushReceiverArrayOutputWithContext(context.Context) AzureAppPushReceiverArrayOutput
}

type AzureAppPushReceiverArray []AzureAppPushReceiverInput

func (AzureAppPushReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureAppPushReceiver)(nil)).Elem()
}

func (i AzureAppPushReceiverArray) ToAzureAppPushReceiverArrayOutput() AzureAppPushReceiverArrayOutput {
	return i.ToAzureAppPushReceiverArrayOutputWithContext(context.Background())
}

func (i AzureAppPushReceiverArray) ToAzureAppPushReceiverArrayOutputWithContext(ctx context.Context) AzureAppPushReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureAppPushReceiverArrayOutput)
}

type AzureAppPushReceiverOutput struct{ *pulumi.OutputState }

func (AzureAppPushReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureAppPushReceiver)(nil)).Elem()
}

func (o AzureAppPushReceiverOutput) ToAzureAppPushReceiverOutput() AzureAppPushReceiverOutput {
	return o
}

func (o AzureAppPushReceiverOutput) ToAzureAppPushReceiverOutputWithContext(ctx context.Context) AzureAppPushReceiverOutput {
	return o
}

func (o AzureAppPushReceiverOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v AzureAppPushReceiver) string { return v.EmailAddress }).(pulumi.StringOutput)
}

func (o AzureAppPushReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureAppPushReceiver) string { return v.Name }).(pulumi.StringOutput)
}

type AzureAppPushReceiverArrayOutput struct{ *pulumi.OutputState }

func (AzureAppPushReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureAppPushReceiver)(nil)).Elem()
}

func (o AzureAppPushReceiverArrayOutput) ToAzureAppPushReceiverArrayOutput() AzureAppPushReceiverArrayOutput {
	return o
}

func (o AzureAppPushReceiverArrayOutput) ToAzureAppPushReceiverArrayOutputWithContext(ctx context.Context) AzureAppPushReceiverArrayOutput {
	return o
}

func (o AzureAppPushReceiverArrayOutput) Index(i pulumi.IntInput) AzureAppPushReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureAppPushReceiver {
		return vs[0].([]AzureAppPushReceiver)[vs[1].(int)]
	}).(AzureAppPushReceiverOutput)
}

type AzureAppPushReceiverResponse struct {
	EmailAddress string `pulumi:"emailAddress"`
	Name         string `pulumi:"name"`
}

type AzureAppPushReceiverResponseOutput struct{ *pulumi.OutputState }

func (AzureAppPushReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureAppPushReceiverResponse)(nil)).Elem()
}

func (o AzureAppPushReceiverResponseOutput) ToAzureAppPushReceiverResponseOutput() AzureAppPushReceiverResponseOutput {
	return o
}

func (o AzureAppPushReceiverResponseOutput) ToAzureAppPushReceiverResponseOutputWithContext(ctx context.Context) AzureAppPushReceiverResponseOutput {
	return o
}

func (o AzureAppPushReceiverResponseOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v AzureAppPushReceiverResponse) string { return v.EmailAddress }).(pulumi.StringOutput)
}

func (o AzureAppPushReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureAppPushReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

type AzureAppPushReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureAppPushReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureAppPushReceiverResponse)(nil)).Elem()
}

func (o AzureAppPushReceiverResponseArrayOutput) ToAzureAppPushReceiverResponseArrayOutput() AzureAppPushReceiverResponseArrayOutput {
	return o
}

func (o AzureAppPushReceiverResponseArrayOutput) ToAzureAppPushReceiverResponseArrayOutputWithContext(ctx context.Context) AzureAppPushReceiverResponseArrayOutput {
	return o
}

func (o AzureAppPushReceiverResponseArrayOutput) Index(i pulumi.IntInput) AzureAppPushReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureAppPushReceiverResponse {
		return vs[0].([]AzureAppPushReceiverResponse)[vs[1].(int)]
	}).(AzureAppPushReceiverResponseOutput)
}

type AzureFunctionReceiver struct {
	FunctionAppResourceId string `pulumi:"functionAppResourceId"`
	FunctionName          string `pulumi:"functionName"`
	HttpTriggerUrl        string `pulumi:"httpTriggerUrl"`
	Name                  string `pulumi:"name"`
	UseCommonAlertSchema  *bool  `pulumi:"useCommonAlertSchema"`
}


func (val *AzureFunctionReceiver) Defaults() *AzureFunctionReceiver {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}





type AzureFunctionReceiverInput interface {
	pulumi.Input

	ToAzureFunctionReceiverOutput() AzureFunctionReceiverOutput
	ToAzureFunctionReceiverOutputWithContext(context.Context) AzureFunctionReceiverOutput
}

type AzureFunctionReceiverArgs struct {
	FunctionAppResourceId pulumi.StringInput  `pulumi:"functionAppResourceId"`
	FunctionName          pulumi.StringInput  `pulumi:"functionName"`
	HttpTriggerUrl        pulumi.StringInput  `pulumi:"httpTriggerUrl"`
	Name                  pulumi.StringInput  `pulumi:"name"`
	UseCommonAlertSchema  pulumi.BoolPtrInput `pulumi:"useCommonAlertSchema"`
}


func (val *AzureFunctionReceiverArgs) Defaults() *AzureFunctionReceiverArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		tmp.UseCommonAlertSchema = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (AzureFunctionReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionReceiver)(nil)).Elem()
}

func (i AzureFunctionReceiverArgs) ToAzureFunctionReceiverOutput() AzureFunctionReceiverOutput {
	return i.ToAzureFunctionReceiverOutputWithContext(context.Background())
}

func (i AzureFunctionReceiverArgs) ToAzureFunctionReceiverOutputWithContext(ctx context.Context) AzureFunctionReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFunctionReceiverOutput)
}





type AzureFunctionReceiverArrayInput interface {
	pulumi.Input

	ToAzureFunctionReceiverArrayOutput() AzureFunctionReceiverArrayOutput
	ToAzureFunctionReceiverArrayOutputWithContext(context.Context) AzureFunctionReceiverArrayOutput
}

type AzureFunctionReceiverArray []AzureFunctionReceiverInput

func (AzureFunctionReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureFunctionReceiver)(nil)).Elem()
}

func (i AzureFunctionReceiverArray) ToAzureFunctionReceiverArrayOutput() AzureFunctionReceiverArrayOutput {
	return i.ToAzureFunctionReceiverArrayOutputWithContext(context.Background())
}

func (i AzureFunctionReceiverArray) ToAzureFunctionReceiverArrayOutputWithContext(ctx context.Context) AzureFunctionReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFunctionReceiverArrayOutput)
}

type AzureFunctionReceiverOutput struct{ *pulumi.OutputState }

func (AzureFunctionReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionReceiver)(nil)).Elem()
}

func (o AzureFunctionReceiverOutput) ToAzureFunctionReceiverOutput() AzureFunctionReceiverOutput {
	return o
}

func (o AzureFunctionReceiverOutput) ToAzureFunctionReceiverOutputWithContext(ctx context.Context) AzureFunctionReceiverOutput {
	return o
}

func (o AzureFunctionReceiverOutput) FunctionAppResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionReceiver) string { return v.FunctionAppResourceId }).(pulumi.StringOutput)
}

func (o AzureFunctionReceiverOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionReceiver) string { return v.FunctionName }).(pulumi.StringOutput)
}

func (o AzureFunctionReceiverOutput) HttpTriggerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionReceiver) string { return v.HttpTriggerUrl }).(pulumi.StringOutput)
}

func (o AzureFunctionReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureFunctionReceiverOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFunctionReceiver) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type AzureFunctionReceiverArrayOutput struct{ *pulumi.OutputState }

func (AzureFunctionReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureFunctionReceiver)(nil)).Elem()
}

func (o AzureFunctionReceiverArrayOutput) ToAzureFunctionReceiverArrayOutput() AzureFunctionReceiverArrayOutput {
	return o
}

func (o AzureFunctionReceiverArrayOutput) ToAzureFunctionReceiverArrayOutputWithContext(ctx context.Context) AzureFunctionReceiverArrayOutput {
	return o
}

func (o AzureFunctionReceiverArrayOutput) Index(i pulumi.IntInput) AzureFunctionReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureFunctionReceiver {
		return vs[0].([]AzureFunctionReceiver)[vs[1].(int)]
	}).(AzureFunctionReceiverOutput)
}

type AzureFunctionReceiverResponse struct {
	FunctionAppResourceId string `pulumi:"functionAppResourceId"`
	FunctionName          string `pulumi:"functionName"`
	HttpTriggerUrl        string `pulumi:"httpTriggerUrl"`
	Name                  string `pulumi:"name"`
	UseCommonAlertSchema  *bool  `pulumi:"useCommonAlertSchema"`
}


func (val *AzureFunctionReceiverResponse) Defaults() *AzureFunctionReceiverResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}

type AzureFunctionReceiverResponseOutput struct{ *pulumi.OutputState }

func (AzureFunctionReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionReceiverResponse)(nil)).Elem()
}

func (o AzureFunctionReceiverResponseOutput) ToAzureFunctionReceiverResponseOutput() AzureFunctionReceiverResponseOutput {
	return o
}

func (o AzureFunctionReceiverResponseOutput) ToAzureFunctionReceiverResponseOutputWithContext(ctx context.Context) AzureFunctionReceiverResponseOutput {
	return o
}

func (o AzureFunctionReceiverResponseOutput) FunctionAppResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionReceiverResponse) string { return v.FunctionAppResourceId }).(pulumi.StringOutput)
}

func (o AzureFunctionReceiverResponseOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionReceiverResponse) string { return v.FunctionName }).(pulumi.StringOutput)
}

func (o AzureFunctionReceiverResponseOutput) HttpTriggerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionReceiverResponse) string { return v.HttpTriggerUrl }).(pulumi.StringOutput)
}

func (o AzureFunctionReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureFunctionReceiverResponseOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureFunctionReceiverResponse) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type AzureFunctionReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureFunctionReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureFunctionReceiverResponse)(nil)).Elem()
}

func (o AzureFunctionReceiverResponseArrayOutput) ToAzureFunctionReceiverResponseArrayOutput() AzureFunctionReceiverResponseArrayOutput {
	return o
}

func (o AzureFunctionReceiverResponseArrayOutput) ToAzureFunctionReceiverResponseArrayOutputWithContext(ctx context.Context) AzureFunctionReceiverResponseArrayOutput {
	return o
}

func (o AzureFunctionReceiverResponseArrayOutput) Index(i pulumi.IntInput) AzureFunctionReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureFunctionReceiverResponse {
		return vs[0].([]AzureFunctionReceiverResponse)[vs[1].(int)]
	}).(AzureFunctionReceiverResponseOutput)
}

type Criteria struct {
	Dimensions []Dimension `pulumi:"dimensions"`
	MetricName string      `pulumi:"metricName"`
}

type CriteriaResponse struct {
	Dimensions []DimensionResponse `pulumi:"dimensions"`
	MetricName string              `pulumi:"metricName"`
}

type DataCollectionEndpointNetworkAcls struct {
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}





type DataCollectionEndpointNetworkAclsInput interface {
	pulumi.Input

	ToDataCollectionEndpointNetworkAclsOutput() DataCollectionEndpointNetworkAclsOutput
	ToDataCollectionEndpointNetworkAclsOutputWithContext(context.Context) DataCollectionEndpointNetworkAclsOutput
}

type DataCollectionEndpointNetworkAclsArgs struct {
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (DataCollectionEndpointNetworkAclsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointNetworkAcls)(nil)).Elem()
}

func (i DataCollectionEndpointNetworkAclsArgs) ToDataCollectionEndpointNetworkAclsOutput() DataCollectionEndpointNetworkAclsOutput {
	return i.ToDataCollectionEndpointNetworkAclsOutputWithContext(context.Background())
}

func (i DataCollectionEndpointNetworkAclsArgs) ToDataCollectionEndpointNetworkAclsOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointNetworkAclsOutput)
}

func (i DataCollectionEndpointNetworkAclsArgs) ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput {
	return i.ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(context.Background())
}

func (i DataCollectionEndpointNetworkAclsArgs) ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointNetworkAclsOutput).ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx)
}









type DataCollectionEndpointNetworkAclsPtrInput interface {
	pulumi.Input

	ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput
	ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(context.Context) DataCollectionEndpointNetworkAclsPtrOutput
}

type dataCollectionEndpointNetworkAclsPtrType DataCollectionEndpointNetworkAclsArgs

func DataCollectionEndpointNetworkAclsPtr(v *DataCollectionEndpointNetworkAclsArgs) DataCollectionEndpointNetworkAclsPtrInput {
	return (*dataCollectionEndpointNetworkAclsPtrType)(v)
}

func (*dataCollectionEndpointNetworkAclsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointNetworkAcls)(nil)).Elem()
}

func (i *dataCollectionEndpointNetworkAclsPtrType) ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput {
	return i.ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(context.Background())
}

func (i *dataCollectionEndpointNetworkAclsPtrType) ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointNetworkAclsPtrOutput)
}

type DataCollectionEndpointNetworkAclsOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointNetworkAclsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointNetworkAcls)(nil)).Elem()
}

func (o DataCollectionEndpointNetworkAclsOutput) ToDataCollectionEndpointNetworkAclsOutput() DataCollectionEndpointNetworkAclsOutput {
	return o
}

func (o DataCollectionEndpointNetworkAclsOutput) ToDataCollectionEndpointNetworkAclsOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsOutput {
	return o
}

func (o DataCollectionEndpointNetworkAclsOutput) ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput {
	return o.ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(context.Background())
}

func (o DataCollectionEndpointNetworkAclsOutput) ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionEndpointNetworkAcls) *DataCollectionEndpointNetworkAcls {
		return &v
	}).(DataCollectionEndpointNetworkAclsPtrOutput)
}

func (o DataCollectionEndpointNetworkAclsOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointNetworkAcls) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointNetworkAclsPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointNetworkAclsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointNetworkAcls)(nil)).Elem()
}

func (o DataCollectionEndpointNetworkAclsPtrOutput) ToDataCollectionEndpointNetworkAclsPtrOutput() DataCollectionEndpointNetworkAclsPtrOutput {
	return o
}

func (o DataCollectionEndpointNetworkAclsPtrOutput) ToDataCollectionEndpointNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointNetworkAclsPtrOutput {
	return o
}

func (o DataCollectionEndpointNetworkAclsPtrOutput) Elem() DataCollectionEndpointNetworkAclsOutput {
	return o.ApplyT(func(v *DataCollectionEndpointNetworkAcls) DataCollectionEndpointNetworkAcls {
		if v != nil {
			return *v
		}
		var ret DataCollectionEndpointNetworkAcls
		return ret
	}).(DataCollectionEndpointNetworkAclsOutput)
}

func (o DataCollectionEndpointNetworkAclsPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointNetworkAcls) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResourceResponseSystemData struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type DataCollectionEndpointResourceResponseSystemDataOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResourceResponseSystemDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResourceResponseSystemData)(nil)).Elem()
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) ToDataCollectionEndpointResourceResponseSystemDataOutput() DataCollectionEndpointResourceResponseSystemDataOutput {
	return o
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) ToDataCollectionEndpointResourceResponseSystemDataOutputWithContext(ctx context.Context) DataCollectionEndpointResourceResponseSystemDataOutput {
	return o
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResourceResponseSystemData) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResponseConfigurationAccess struct {
	Endpoint string `pulumi:"endpoint"`
}

type DataCollectionEndpointResponseConfigurationAccessOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseConfigurationAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResponseConfigurationAccess)(nil)).Elem()
}

func (o DataCollectionEndpointResponseConfigurationAccessOutput) ToDataCollectionEndpointResponseConfigurationAccessOutput() DataCollectionEndpointResponseConfigurationAccessOutput {
	return o
}

func (o DataCollectionEndpointResponseConfigurationAccessOutput) ToDataCollectionEndpointResponseConfigurationAccessOutputWithContext(ctx context.Context) DataCollectionEndpointResponseConfigurationAccessOutput {
	return o
}

func (o DataCollectionEndpointResponseConfigurationAccessOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DataCollectionEndpointResponseConfigurationAccess) string { return v.Endpoint }).(pulumi.StringOutput)
}

type DataCollectionEndpointResponseConfigurationAccessPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseConfigurationAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResponseConfigurationAccess)(nil)).Elem()
}

func (o DataCollectionEndpointResponseConfigurationAccessPtrOutput) ToDataCollectionEndpointResponseConfigurationAccessPtrOutput() DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseConfigurationAccessPtrOutput) ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseConfigurationAccessPtrOutput) Elem() DataCollectionEndpointResponseConfigurationAccessOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseConfigurationAccess) DataCollectionEndpointResponseConfigurationAccess {
		if v != nil {
			return *v
		}
		var ret DataCollectionEndpointResponseConfigurationAccess
		return ret
	}).(DataCollectionEndpointResponseConfigurationAccessOutput)
}

func (o DataCollectionEndpointResponseConfigurationAccessPtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseConfigurationAccess) *string {
		if v == nil {
			return nil
		}
		return &v.Endpoint
	}).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResponseLogsIngestion struct {
	Endpoint string `pulumi:"endpoint"`
}

type DataCollectionEndpointResponseLogsIngestionOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseLogsIngestionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResponseLogsIngestion)(nil)).Elem()
}

func (o DataCollectionEndpointResponseLogsIngestionOutput) ToDataCollectionEndpointResponseLogsIngestionOutput() DataCollectionEndpointResponseLogsIngestionOutput {
	return o
}

func (o DataCollectionEndpointResponseLogsIngestionOutput) ToDataCollectionEndpointResponseLogsIngestionOutputWithContext(ctx context.Context) DataCollectionEndpointResponseLogsIngestionOutput {
	return o
}

func (o DataCollectionEndpointResponseLogsIngestionOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v DataCollectionEndpointResponseLogsIngestion) string { return v.Endpoint }).(pulumi.StringOutput)
}

type DataCollectionEndpointResponseLogsIngestionPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseLogsIngestionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResponseLogsIngestion)(nil)).Elem()
}

func (o DataCollectionEndpointResponseLogsIngestionPtrOutput) ToDataCollectionEndpointResponseLogsIngestionPtrOutput() DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseLogsIngestionPtrOutput) ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseLogsIngestionPtrOutput) Elem() DataCollectionEndpointResponseLogsIngestionOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseLogsIngestion) DataCollectionEndpointResponseLogsIngestion {
		if v != nil {
			return *v
		}
		var ret DataCollectionEndpointResponseLogsIngestion
		return ret
	}).(DataCollectionEndpointResponseLogsIngestionOutput)
}

func (o DataCollectionEndpointResponseLogsIngestionPtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseLogsIngestion) *string {
		if v == nil {
			return nil
		}
		return &v.Endpoint
	}).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResponseNetworkAcls struct {
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
}

type DataCollectionEndpointResponseNetworkAclsOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseNetworkAclsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResponseNetworkAcls)(nil)).Elem()
}

func (o DataCollectionEndpointResponseNetworkAclsOutput) ToDataCollectionEndpointResponseNetworkAclsOutput() DataCollectionEndpointResponseNetworkAclsOutput {
	return o
}

func (o DataCollectionEndpointResponseNetworkAclsOutput) ToDataCollectionEndpointResponseNetworkAclsOutputWithContext(ctx context.Context) DataCollectionEndpointResponseNetworkAclsOutput {
	return o
}

func (o DataCollectionEndpointResponseNetworkAclsOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCollectionEndpointResponseNetworkAcls) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResponseNetworkAclsPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResponseNetworkAclsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResponseNetworkAcls)(nil)).Elem()
}

func (o DataCollectionEndpointResponseNetworkAclsPtrOutput) ToDataCollectionEndpointResponseNetworkAclsPtrOutput() DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseNetworkAclsPtrOutput) ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return o
}

func (o DataCollectionEndpointResponseNetworkAclsPtrOutput) Elem() DataCollectionEndpointResponseNetworkAclsOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseNetworkAcls) DataCollectionEndpointResponseNetworkAcls {
		if v != nil {
			return *v
		}
		var ret DataCollectionEndpointResponseNetworkAcls
		return ret
	}).(DataCollectionEndpointResponseNetworkAclsOutput)
}

func (o DataCollectionEndpointResponseNetworkAclsPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResponseNetworkAcls) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type DataCollectionRuleDataSources struct {
	Extensions          []ExtensionDataSource       `pulumi:"extensions"`
	PerformanceCounters []PerfCounterDataSource     `pulumi:"performanceCounters"`
	Syslog              []SyslogDataSource          `pulumi:"syslog"`
	WindowsEventLogs    []WindowsEventLogDataSource `pulumi:"windowsEventLogs"`
}





type DataCollectionRuleDataSourcesInput interface {
	pulumi.Input

	ToDataCollectionRuleDataSourcesOutput() DataCollectionRuleDataSourcesOutput
	ToDataCollectionRuleDataSourcesOutputWithContext(context.Context) DataCollectionRuleDataSourcesOutput
}

type DataCollectionRuleDataSourcesArgs struct {
	Extensions          ExtensionDataSourceArrayInput       `pulumi:"extensions"`
	PerformanceCounters PerfCounterDataSourceArrayInput     `pulumi:"performanceCounters"`
	Syslog              SyslogDataSourceArrayInput          `pulumi:"syslog"`
	WindowsEventLogs    WindowsEventLogDataSourceArrayInput `pulumi:"windowsEventLogs"`
}

func (DataCollectionRuleDataSourcesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleDataSources)(nil)).Elem()
}

func (i DataCollectionRuleDataSourcesArgs) ToDataCollectionRuleDataSourcesOutput() DataCollectionRuleDataSourcesOutput {
	return i.ToDataCollectionRuleDataSourcesOutputWithContext(context.Background())
}

func (i DataCollectionRuleDataSourcesArgs) ToDataCollectionRuleDataSourcesOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDataSourcesOutput)
}

func (i DataCollectionRuleDataSourcesArgs) ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput {
	return i.ToDataCollectionRuleDataSourcesPtrOutputWithContext(context.Background())
}

func (i DataCollectionRuleDataSourcesArgs) ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDataSourcesOutput).ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx)
}









type DataCollectionRuleDataSourcesPtrInput interface {
	pulumi.Input

	ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput
	ToDataCollectionRuleDataSourcesPtrOutputWithContext(context.Context) DataCollectionRuleDataSourcesPtrOutput
}

type dataCollectionRuleDataSourcesPtrType DataCollectionRuleDataSourcesArgs

func DataCollectionRuleDataSourcesPtr(v *DataCollectionRuleDataSourcesArgs) DataCollectionRuleDataSourcesPtrInput {
	return (*dataCollectionRuleDataSourcesPtrType)(v)
}

func (*dataCollectionRuleDataSourcesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleDataSources)(nil)).Elem()
}

func (i *dataCollectionRuleDataSourcesPtrType) ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput {
	return i.ToDataCollectionRuleDataSourcesPtrOutputWithContext(context.Background())
}

func (i *dataCollectionRuleDataSourcesPtrType) ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDataSourcesPtrOutput)
}

type DataCollectionRuleDataSourcesOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleDataSourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleDataSources)(nil)).Elem()
}

func (o DataCollectionRuleDataSourcesOutput) ToDataCollectionRuleDataSourcesOutput() DataCollectionRuleDataSourcesOutput {
	return o
}

func (o DataCollectionRuleDataSourcesOutput) ToDataCollectionRuleDataSourcesOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesOutput {
	return o
}

func (o DataCollectionRuleDataSourcesOutput) ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput {
	return o.ToDataCollectionRuleDataSourcesPtrOutputWithContext(context.Background())
}

func (o DataCollectionRuleDataSourcesOutput) ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionRuleDataSources) *DataCollectionRuleDataSources {
		return &v
	}).(DataCollectionRuleDataSourcesPtrOutput)
}

func (o DataCollectionRuleDataSourcesOutput) Extensions() ExtensionDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []ExtensionDataSource { return v.Extensions }).(ExtensionDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesOutput) PerformanceCounters() PerfCounterDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []PerfCounterDataSource { return v.PerformanceCounters }).(PerfCounterDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesOutput) Syslog() SyslogDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []SyslogDataSource { return v.Syslog }).(SyslogDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesOutput) WindowsEventLogs() WindowsEventLogDataSourceArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDataSources) []WindowsEventLogDataSource { return v.WindowsEventLogs }).(WindowsEventLogDataSourceArrayOutput)
}

type DataCollectionRuleDataSourcesPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleDataSourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleDataSources)(nil)).Elem()
}

func (o DataCollectionRuleDataSourcesPtrOutput) ToDataCollectionRuleDataSourcesPtrOutput() DataCollectionRuleDataSourcesPtrOutput {
	return o
}

func (o DataCollectionRuleDataSourcesPtrOutput) ToDataCollectionRuleDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleDataSourcesPtrOutput {
	return o
}

func (o DataCollectionRuleDataSourcesPtrOutput) Elem() DataCollectionRuleDataSourcesOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) DataCollectionRuleDataSources {
		if v != nil {
			return *v
		}
		var ret DataCollectionRuleDataSources
		return ret
	}).(DataCollectionRuleDataSourcesOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) Extensions() ExtensionDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []ExtensionDataSource {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(ExtensionDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) PerformanceCounters() PerfCounterDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []PerfCounterDataSource {
		if v == nil {
			return nil
		}
		return v.PerformanceCounters
	}).(PerfCounterDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) Syslog() SyslogDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []SyslogDataSource {
		if v == nil {
			return nil
		}
		return v.Syslog
	}).(SyslogDataSourceArrayOutput)
}

func (o DataCollectionRuleDataSourcesPtrOutput) WindowsEventLogs() WindowsEventLogDataSourceArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDataSources) []WindowsEventLogDataSource {
		if v == nil {
			return nil
		}
		return v.WindowsEventLogs
	}).(WindowsEventLogDataSourceArrayOutput)
}

type DataCollectionRuleDestinations struct {
	AzureMonitorMetrics *DestinationsSpecAzureMonitorMetrics `pulumi:"azureMonitorMetrics"`
	LogAnalytics        []LogAnalyticsDestination            `pulumi:"logAnalytics"`
}





type DataCollectionRuleDestinationsInput interface {
	pulumi.Input

	ToDataCollectionRuleDestinationsOutput() DataCollectionRuleDestinationsOutput
	ToDataCollectionRuleDestinationsOutputWithContext(context.Context) DataCollectionRuleDestinationsOutput
}

type DataCollectionRuleDestinationsArgs struct {
	AzureMonitorMetrics DestinationsSpecAzureMonitorMetricsPtrInput `pulumi:"azureMonitorMetrics"`
	LogAnalytics        LogAnalyticsDestinationArrayInput           `pulumi:"logAnalytics"`
}

func (DataCollectionRuleDestinationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleDestinations)(nil)).Elem()
}

func (i DataCollectionRuleDestinationsArgs) ToDataCollectionRuleDestinationsOutput() DataCollectionRuleDestinationsOutput {
	return i.ToDataCollectionRuleDestinationsOutputWithContext(context.Background())
}

func (i DataCollectionRuleDestinationsArgs) ToDataCollectionRuleDestinationsOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDestinationsOutput)
}

func (i DataCollectionRuleDestinationsArgs) ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput {
	return i.ToDataCollectionRuleDestinationsPtrOutputWithContext(context.Background())
}

func (i DataCollectionRuleDestinationsArgs) ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDestinationsOutput).ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx)
}









type DataCollectionRuleDestinationsPtrInput interface {
	pulumi.Input

	ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput
	ToDataCollectionRuleDestinationsPtrOutputWithContext(context.Context) DataCollectionRuleDestinationsPtrOutput
}

type dataCollectionRuleDestinationsPtrType DataCollectionRuleDestinationsArgs

func DataCollectionRuleDestinationsPtr(v *DataCollectionRuleDestinationsArgs) DataCollectionRuleDestinationsPtrInput {
	return (*dataCollectionRuleDestinationsPtrType)(v)
}

func (*dataCollectionRuleDestinationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleDestinations)(nil)).Elem()
}

func (i *dataCollectionRuleDestinationsPtrType) ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput {
	return i.ToDataCollectionRuleDestinationsPtrOutputWithContext(context.Background())
}

func (i *dataCollectionRuleDestinationsPtrType) ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleDestinationsPtrOutput)
}

type DataCollectionRuleDestinationsOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleDestinationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleDestinations)(nil)).Elem()
}

func (o DataCollectionRuleDestinationsOutput) ToDataCollectionRuleDestinationsOutput() DataCollectionRuleDestinationsOutput {
	return o
}

func (o DataCollectionRuleDestinationsOutput) ToDataCollectionRuleDestinationsOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsOutput {
	return o
}

func (o DataCollectionRuleDestinationsOutput) ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput {
	return o.ToDataCollectionRuleDestinationsPtrOutputWithContext(context.Background())
}

func (o DataCollectionRuleDestinationsOutput) ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionRuleDestinations) *DataCollectionRuleDestinations {
		return &v
	}).(DataCollectionRuleDestinationsPtrOutput)
}

func (o DataCollectionRuleDestinationsOutput) AzureMonitorMetrics() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleDestinations) *DestinationsSpecAzureMonitorMetrics {
		return v.AzureMonitorMetrics
	}).(DestinationsSpecAzureMonitorMetricsPtrOutput)
}

func (o DataCollectionRuleDestinationsOutput) LogAnalytics() LogAnalyticsDestinationArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleDestinations) []LogAnalyticsDestination { return v.LogAnalytics }).(LogAnalyticsDestinationArrayOutput)
}

type DataCollectionRuleDestinationsPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleDestinationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleDestinations)(nil)).Elem()
}

func (o DataCollectionRuleDestinationsPtrOutput) ToDataCollectionRuleDestinationsPtrOutput() DataCollectionRuleDestinationsPtrOutput {
	return o
}

func (o DataCollectionRuleDestinationsPtrOutput) ToDataCollectionRuleDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleDestinationsPtrOutput {
	return o
}

func (o DataCollectionRuleDestinationsPtrOutput) Elem() DataCollectionRuleDestinationsOutput {
	return o.ApplyT(func(v *DataCollectionRuleDestinations) DataCollectionRuleDestinations {
		if v != nil {
			return *v
		}
		var ret DataCollectionRuleDestinations
		return ret
	}).(DataCollectionRuleDestinationsOutput)
}

func (o DataCollectionRuleDestinationsPtrOutput) AzureMonitorMetrics() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o.ApplyT(func(v *DataCollectionRuleDestinations) *DestinationsSpecAzureMonitorMetrics {
		if v == nil {
			return nil
		}
		return v.AzureMonitorMetrics
	}).(DestinationsSpecAzureMonitorMetricsPtrOutput)
}

func (o DataCollectionRuleDestinationsPtrOutput) LogAnalytics() LogAnalyticsDestinationArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleDestinations) []LogAnalyticsDestination {
		if v == nil {
			return nil
		}
		return v.LogAnalytics
	}).(LogAnalyticsDestinationArrayOutput)
}

type DataCollectionRuleResponseDataSources struct {
	Extensions          []ExtensionDataSourceResponse       `pulumi:"extensions"`
	PerformanceCounters []PerfCounterDataSourceResponse     `pulumi:"performanceCounters"`
	Syslog              []SyslogDataSourceResponse          `pulumi:"syslog"`
	WindowsEventLogs    []WindowsEventLogDataSourceResponse `pulumi:"windowsEventLogs"`
}

type DataCollectionRuleResponseDataSourcesOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResponseDataSourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleResponseDataSources)(nil)).Elem()
}

func (o DataCollectionRuleResponseDataSourcesOutput) ToDataCollectionRuleResponseDataSourcesOutput() DataCollectionRuleResponseDataSourcesOutput {
	return o
}

func (o DataCollectionRuleResponseDataSourcesOutput) ToDataCollectionRuleResponseDataSourcesOutputWithContext(ctx context.Context) DataCollectionRuleResponseDataSourcesOutput {
	return o
}

func (o DataCollectionRuleResponseDataSourcesOutput) Extensions() ExtensionDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []ExtensionDataSourceResponse { return v.Extensions }).(ExtensionDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesOutput) PerformanceCounters() PerfCounterDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []PerfCounterDataSourceResponse {
		return v.PerformanceCounters
	}).(PerfCounterDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesOutput) Syslog() SyslogDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []SyslogDataSourceResponse { return v.Syslog }).(SyslogDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesOutput) WindowsEventLogs() WindowsEventLogDataSourceResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDataSources) []WindowsEventLogDataSourceResponse {
		return v.WindowsEventLogs
	}).(WindowsEventLogDataSourceResponseArrayOutput)
}

type DataCollectionRuleResponseDataSourcesPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResponseDataSourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleResponseDataSources)(nil)).Elem()
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) ToDataCollectionRuleResponseDataSourcesPtrOutput() DataCollectionRuleResponseDataSourcesPtrOutput {
	return o
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDataSourcesPtrOutput {
	return o
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) Elem() DataCollectionRuleResponseDataSourcesOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) DataCollectionRuleResponseDataSources {
		if v != nil {
			return *v
		}
		var ret DataCollectionRuleResponseDataSources
		return ret
	}).(DataCollectionRuleResponseDataSourcesOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) Extensions() ExtensionDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []ExtensionDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(ExtensionDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) PerformanceCounters() PerfCounterDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []PerfCounterDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.PerformanceCounters
	}).(PerfCounterDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) Syslog() SyslogDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []SyslogDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.Syslog
	}).(SyslogDataSourceResponseArrayOutput)
}

func (o DataCollectionRuleResponseDataSourcesPtrOutput) WindowsEventLogs() WindowsEventLogDataSourceResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDataSources) []WindowsEventLogDataSourceResponse {
		if v == nil {
			return nil
		}
		return v.WindowsEventLogs
	}).(WindowsEventLogDataSourceResponseArrayOutput)
}

type DataCollectionRuleResponseDestinations struct {
	AzureMonitorMetrics *DestinationsSpecResponseAzureMonitorMetrics `pulumi:"azureMonitorMetrics"`
	LogAnalytics        []LogAnalyticsDestinationResponse            `pulumi:"logAnalytics"`
}

type DataCollectionRuleResponseDestinationsOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResponseDestinationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleResponseDestinations)(nil)).Elem()
}

func (o DataCollectionRuleResponseDestinationsOutput) ToDataCollectionRuleResponseDestinationsOutput() DataCollectionRuleResponseDestinationsOutput {
	return o
}

func (o DataCollectionRuleResponseDestinationsOutput) ToDataCollectionRuleResponseDestinationsOutputWithContext(ctx context.Context) DataCollectionRuleResponseDestinationsOutput {
	return o
}

func (o DataCollectionRuleResponseDestinationsOutput) AzureMonitorMetrics() DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDestinations) *DestinationsSpecResponseAzureMonitorMetrics {
		return v.AzureMonitorMetrics
	}).(DestinationsSpecResponseAzureMonitorMetricsPtrOutput)
}

func (o DataCollectionRuleResponseDestinationsOutput) LogAnalytics() LogAnalyticsDestinationResponseArrayOutput {
	return o.ApplyT(func(v DataCollectionRuleResponseDestinations) []LogAnalyticsDestinationResponse {
		return v.LogAnalytics
	}).(LogAnalyticsDestinationResponseArrayOutput)
}

type DataCollectionRuleResponseDestinationsPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleResponseDestinationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleResponseDestinations)(nil)).Elem()
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) ToDataCollectionRuleResponseDestinationsPtrOutput() DataCollectionRuleResponseDestinationsPtrOutput {
	return o
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDestinationsPtrOutput {
	return o
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) Elem() DataCollectionRuleResponseDestinationsOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDestinations) DataCollectionRuleResponseDestinations {
		if v != nil {
			return *v
		}
		var ret DataCollectionRuleResponseDestinations
		return ret
	}).(DataCollectionRuleResponseDestinationsOutput)
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) AzureMonitorMetrics() DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDestinations) *DestinationsSpecResponseAzureMonitorMetrics {
		if v == nil {
			return nil
		}
		return v.AzureMonitorMetrics
	}).(DestinationsSpecResponseAzureMonitorMetricsPtrOutput)
}

func (o DataCollectionRuleResponseDestinationsPtrOutput) LogAnalytics() LogAnalyticsDestinationResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRuleResponseDestinations) []LogAnalyticsDestinationResponse {
		if v == nil {
			return nil
		}
		return v.LogAnalytics
	}).(LogAnalyticsDestinationResponseArrayOutput)
}

type DataFlow struct {
	Destinations []string `pulumi:"destinations"`
	Streams      []string `pulumi:"streams"`
}





type DataFlowInput interface {
	pulumi.Input

	ToDataFlowOutput() DataFlowOutput
	ToDataFlowOutputWithContext(context.Context) DataFlowOutput
}

type DataFlowArgs struct {
	Destinations pulumi.StringArrayInput `pulumi:"destinations"`
	Streams      pulumi.StringArrayInput `pulumi:"streams"`
}

func (DataFlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlow)(nil)).Elem()
}

func (i DataFlowArgs) ToDataFlowOutput() DataFlowOutput {
	return i.ToDataFlowOutputWithContext(context.Background())
}

func (i DataFlowArgs) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowOutput)
}





type DataFlowArrayInput interface {
	pulumi.Input

	ToDataFlowArrayOutput() DataFlowArrayOutput
	ToDataFlowArrayOutputWithContext(context.Context) DataFlowArrayOutput
}

type DataFlowArray []DataFlowInput

func (DataFlowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataFlow)(nil)).Elem()
}

func (i DataFlowArray) ToDataFlowArrayOutput() DataFlowArrayOutput {
	return i.ToDataFlowArrayOutputWithContext(context.Background())
}

func (i DataFlowArray) ToDataFlowArrayOutputWithContext(ctx context.Context) DataFlowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowArrayOutput)
}

type DataFlowOutput struct{ *pulumi.OutputState }

func (DataFlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlow)(nil)).Elem()
}

func (o DataFlowOutput) ToDataFlowOutput() DataFlowOutput {
	return o
}

func (o DataFlowOutput) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return o
}

func (o DataFlowOutput) Destinations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlow) []string { return v.Destinations }).(pulumi.StringArrayOutput)
}

func (o DataFlowOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlow) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type DataFlowArrayOutput struct{ *pulumi.OutputState }

func (DataFlowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataFlow)(nil)).Elem()
}

func (o DataFlowArrayOutput) ToDataFlowArrayOutput() DataFlowArrayOutput {
	return o
}

func (o DataFlowArrayOutput) ToDataFlowArrayOutputWithContext(ctx context.Context) DataFlowArrayOutput {
	return o
}

func (o DataFlowArrayOutput) Index(i pulumi.IntInput) DataFlowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataFlow {
		return vs[0].([]DataFlow)[vs[1].(int)]
	}).(DataFlowOutput)
}

type DataFlowResponse struct {
	Destinations []string `pulumi:"destinations"`
	Streams      []string `pulumi:"streams"`
}

type DataFlowResponseOutput struct{ *pulumi.OutputState }

func (DataFlowResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlowResponse)(nil)).Elem()
}

func (o DataFlowResponseOutput) ToDataFlowResponseOutput() DataFlowResponseOutput {
	return o
}

func (o DataFlowResponseOutput) ToDataFlowResponseOutputWithContext(ctx context.Context) DataFlowResponseOutput {
	return o
}

func (o DataFlowResponseOutput) Destinations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlowResponse) []string { return v.Destinations }).(pulumi.StringArrayOutput)
}

func (o DataFlowResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DataFlowResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type DataFlowResponseArrayOutput struct{ *pulumi.OutputState }

func (DataFlowResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataFlowResponse)(nil)).Elem()
}

func (o DataFlowResponseArrayOutput) ToDataFlowResponseArrayOutput() DataFlowResponseArrayOutput {
	return o
}

func (o DataFlowResponseArrayOutput) ToDataFlowResponseArrayOutputWithContext(ctx context.Context) DataFlowResponseArrayOutput {
	return o
}

func (o DataFlowResponseArrayOutput) Index(i pulumi.IntInput) DataFlowResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataFlowResponse {
		return vs[0].([]DataFlowResponse)[vs[1].(int)]
	}).(DataFlowResponseOutput)
}

type DataSource struct {
	Configuration DataSourceConfiguration `pulumi:"configuration"`
	Kind          string                  `pulumi:"kind"`
	Sinks         []SinkConfiguration     `pulumi:"sinks"`
}





type DataSourceInput interface {
	pulumi.Input

	ToDataSourceOutput() DataSourceOutput
	ToDataSourceOutputWithContext(context.Context) DataSourceOutput
}

type DataSourceArgs struct {
	Configuration DataSourceConfigurationInput `pulumi:"configuration"`
	Kind          pulumi.StringInput           `pulumi:"kind"`
	Sinks         SinkConfigurationArrayInput  `pulumi:"sinks"`
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSource)(nil)).Elem()
}

func (i DataSourceArgs) ToDataSourceOutput() DataSourceOutput {
	return i.ToDataSourceOutputWithContext(context.Background())
}

func (i DataSourceArgs) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceOutput)
}





type DataSourceArrayInput interface {
	pulumi.Input

	ToDataSourceArrayOutput() DataSourceArrayOutput
	ToDataSourceArrayOutputWithContext(context.Context) DataSourceArrayOutput
}

type DataSourceArray []DataSourceInput

func (DataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSource)(nil)).Elem()
}

func (i DataSourceArray) ToDataSourceArrayOutput() DataSourceArrayOutput {
	return i.ToDataSourceArrayOutputWithContext(context.Background())
}

func (i DataSourceArray) ToDataSourceArrayOutputWithContext(ctx context.Context) DataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceArrayOutput)
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSource)(nil)).Elem()
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

func (o DataSourceOutput) Configuration() DataSourceConfigurationOutput {
	return o.ApplyT(func(v DataSource) DataSourceConfiguration { return v.Configuration }).(DataSourceConfigurationOutput)
}

func (o DataSourceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v DataSource) string { return v.Kind }).(pulumi.StringOutput)
}

func (o DataSourceOutput) Sinks() SinkConfigurationArrayOutput {
	return o.ApplyT(func(v DataSource) []SinkConfiguration { return v.Sinks }).(SinkConfigurationArrayOutput)
}

type DataSourceArrayOutput struct{ *pulumi.OutputState }

func (DataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSource)(nil)).Elem()
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutput() DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) ToDataSourceArrayOutputWithContext(ctx context.Context) DataSourceArrayOutput {
	return o
}

func (o DataSourceArrayOutput) Index(i pulumi.IntInput) DataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataSource {
		return vs[0].([]DataSource)[vs[1].(int)]
	}).(DataSourceOutput)
}

type DataSourceConfiguration struct {
	EventLogs    []EventLogConfiguration           `pulumi:"eventLogs"`
	PerfCounters []PerformanceCounterConfiguration `pulumi:"perfCounters"`
	Providers    []EtwProviderConfiguration        `pulumi:"providers"`
}





type DataSourceConfigurationInput interface {
	pulumi.Input

	ToDataSourceConfigurationOutput() DataSourceConfigurationOutput
	ToDataSourceConfigurationOutputWithContext(context.Context) DataSourceConfigurationOutput
}

type DataSourceConfigurationArgs struct {
	EventLogs    EventLogConfigurationArrayInput           `pulumi:"eventLogs"`
	PerfCounters PerformanceCounterConfigurationArrayInput `pulumi:"perfCounters"`
	Providers    EtwProviderConfigurationArrayInput        `pulumi:"providers"`
}

func (DataSourceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceConfiguration)(nil)).Elem()
}

func (i DataSourceConfigurationArgs) ToDataSourceConfigurationOutput() DataSourceConfigurationOutput {
	return i.ToDataSourceConfigurationOutputWithContext(context.Background())
}

func (i DataSourceConfigurationArgs) ToDataSourceConfigurationOutputWithContext(ctx context.Context) DataSourceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceConfigurationOutput)
}

type DataSourceConfigurationOutput struct{ *pulumi.OutputState }

func (DataSourceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceConfiguration)(nil)).Elem()
}

func (o DataSourceConfigurationOutput) ToDataSourceConfigurationOutput() DataSourceConfigurationOutput {
	return o
}

func (o DataSourceConfigurationOutput) ToDataSourceConfigurationOutputWithContext(ctx context.Context) DataSourceConfigurationOutput {
	return o
}

func (o DataSourceConfigurationOutput) EventLogs() EventLogConfigurationArrayOutput {
	return o.ApplyT(func(v DataSourceConfiguration) []EventLogConfiguration { return v.EventLogs }).(EventLogConfigurationArrayOutput)
}

func (o DataSourceConfigurationOutput) PerfCounters() PerformanceCounterConfigurationArrayOutput {
	return o.ApplyT(func(v DataSourceConfiguration) []PerformanceCounterConfiguration { return v.PerfCounters }).(PerformanceCounterConfigurationArrayOutput)
}

func (o DataSourceConfigurationOutput) Providers() EtwProviderConfigurationArrayOutput {
	return o.ApplyT(func(v DataSourceConfiguration) []EtwProviderConfiguration { return v.Providers }).(EtwProviderConfigurationArrayOutput)
}

type DataSourceConfigurationResponse struct {
	EventLogs    []EventLogConfigurationResponse           `pulumi:"eventLogs"`
	PerfCounters []PerformanceCounterConfigurationResponse `pulumi:"perfCounters"`
	Providers    []EtwProviderConfigurationResponse        `pulumi:"providers"`
}

type DataSourceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (DataSourceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceConfigurationResponse)(nil)).Elem()
}

func (o DataSourceConfigurationResponseOutput) ToDataSourceConfigurationResponseOutput() DataSourceConfigurationResponseOutput {
	return o
}

func (o DataSourceConfigurationResponseOutput) ToDataSourceConfigurationResponseOutputWithContext(ctx context.Context) DataSourceConfigurationResponseOutput {
	return o
}

func (o DataSourceConfigurationResponseOutput) EventLogs() EventLogConfigurationResponseArrayOutput {
	return o.ApplyT(func(v DataSourceConfigurationResponse) []EventLogConfigurationResponse { return v.EventLogs }).(EventLogConfigurationResponseArrayOutput)
}

func (o DataSourceConfigurationResponseOutput) PerfCounters() PerformanceCounterConfigurationResponseArrayOutput {
	return o.ApplyT(func(v DataSourceConfigurationResponse) []PerformanceCounterConfigurationResponse {
		return v.PerfCounters
	}).(PerformanceCounterConfigurationResponseArrayOutput)
}

func (o DataSourceConfigurationResponseOutput) Providers() EtwProviderConfigurationResponseArrayOutput {
	return o.ApplyT(func(v DataSourceConfigurationResponse) []EtwProviderConfigurationResponse { return v.Providers }).(EtwProviderConfigurationResponseArrayOutput)
}

type DataSourceResponse struct {
	Configuration DataSourceConfigurationResponse `pulumi:"configuration"`
	Kind          string                          `pulumi:"kind"`
	Sinks         []SinkConfigurationResponse     `pulumi:"sinks"`
}

type DataSourceResponseOutput struct{ *pulumi.OutputState }

func (DataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceResponse)(nil)).Elem()
}

func (o DataSourceResponseOutput) ToDataSourceResponseOutput() DataSourceResponseOutput {
	return o
}

func (o DataSourceResponseOutput) ToDataSourceResponseOutputWithContext(ctx context.Context) DataSourceResponseOutput {
	return o
}

func (o DataSourceResponseOutput) Configuration() DataSourceConfigurationResponseOutput {
	return o.ApplyT(func(v DataSourceResponse) DataSourceConfigurationResponse { return v.Configuration }).(DataSourceConfigurationResponseOutput)
}

func (o DataSourceResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v DataSourceResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o DataSourceResponseOutput) Sinks() SinkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v DataSourceResponse) []SinkConfigurationResponse { return v.Sinks }).(SinkConfigurationResponseArrayOutput)
}

type DataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (DataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSourceResponse)(nil)).Elem()
}

func (o DataSourceResponseArrayOutput) ToDataSourceResponseArrayOutput() DataSourceResponseArrayOutput {
	return o
}

func (o DataSourceResponseArrayOutput) ToDataSourceResponseArrayOutputWithContext(ctx context.Context) DataSourceResponseArrayOutput {
	return o
}

func (o DataSourceResponseArrayOutput) Index(i pulumi.IntInput) DataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataSourceResponse {
		return vs[0].([]DataSourceResponse)[vs[1].(int)]
	}).(DataSourceResponseOutput)
}

type DestinationsSpecAzureMonitorMetrics struct {
	Name *string `pulumi:"name"`
}





type DestinationsSpecAzureMonitorMetricsInput interface {
	pulumi.Input

	ToDestinationsSpecAzureMonitorMetricsOutput() DestinationsSpecAzureMonitorMetricsOutput
	ToDestinationsSpecAzureMonitorMetricsOutputWithContext(context.Context) DestinationsSpecAzureMonitorMetricsOutput
}

type DestinationsSpecAzureMonitorMetricsArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DestinationsSpecAzureMonitorMetricsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationsSpecAzureMonitorMetrics)(nil)).Elem()
}

func (i DestinationsSpecAzureMonitorMetricsArgs) ToDestinationsSpecAzureMonitorMetricsOutput() DestinationsSpecAzureMonitorMetricsOutput {
	return i.ToDestinationsSpecAzureMonitorMetricsOutputWithContext(context.Background())
}

func (i DestinationsSpecAzureMonitorMetricsArgs) ToDestinationsSpecAzureMonitorMetricsOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationsSpecAzureMonitorMetricsOutput)
}

func (i DestinationsSpecAzureMonitorMetricsArgs) ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return i.ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(context.Background())
}

func (i DestinationsSpecAzureMonitorMetricsArgs) ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationsSpecAzureMonitorMetricsOutput).ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx)
}









type DestinationsSpecAzureMonitorMetricsPtrInput interface {
	pulumi.Input

	ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput
	ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput
}

type destinationsSpecAzureMonitorMetricsPtrType DestinationsSpecAzureMonitorMetricsArgs

func DestinationsSpecAzureMonitorMetricsPtr(v *DestinationsSpecAzureMonitorMetricsArgs) DestinationsSpecAzureMonitorMetricsPtrInput {
	return (*destinationsSpecAzureMonitorMetricsPtrType)(v)
}

func (*destinationsSpecAzureMonitorMetricsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationsSpecAzureMonitorMetrics)(nil)).Elem()
}

func (i *destinationsSpecAzureMonitorMetricsPtrType) ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return i.ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(context.Background())
}

func (i *destinationsSpecAzureMonitorMetricsPtrType) ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationsSpecAzureMonitorMetricsPtrOutput)
}

type DestinationsSpecAzureMonitorMetricsOutput struct{ *pulumi.OutputState }

func (DestinationsSpecAzureMonitorMetricsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationsSpecAzureMonitorMetrics)(nil)).Elem()
}

func (o DestinationsSpecAzureMonitorMetricsOutput) ToDestinationsSpecAzureMonitorMetricsOutput() DestinationsSpecAzureMonitorMetricsOutput {
	return o
}

func (o DestinationsSpecAzureMonitorMetricsOutput) ToDestinationsSpecAzureMonitorMetricsOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsOutput {
	return o
}

func (o DestinationsSpecAzureMonitorMetricsOutput) ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o.ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(context.Background())
}

func (o DestinationsSpecAzureMonitorMetricsOutput) ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DestinationsSpecAzureMonitorMetrics) *DestinationsSpecAzureMonitorMetrics {
		return &v
	}).(DestinationsSpecAzureMonitorMetricsPtrOutput)
}

func (o DestinationsSpecAzureMonitorMetricsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationsSpecAzureMonitorMetrics) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DestinationsSpecAzureMonitorMetricsPtrOutput struct{ *pulumi.OutputState }

func (DestinationsSpecAzureMonitorMetricsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationsSpecAzureMonitorMetrics)(nil)).Elem()
}

func (o DestinationsSpecAzureMonitorMetricsPtrOutput) ToDestinationsSpecAzureMonitorMetricsPtrOutput() DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o
}

func (o DestinationsSpecAzureMonitorMetricsPtrOutput) ToDestinationsSpecAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecAzureMonitorMetricsPtrOutput {
	return o
}

func (o DestinationsSpecAzureMonitorMetricsPtrOutput) Elem() DestinationsSpecAzureMonitorMetricsOutput {
	return o.ApplyT(func(v *DestinationsSpecAzureMonitorMetrics) DestinationsSpecAzureMonitorMetrics {
		if v != nil {
			return *v
		}
		var ret DestinationsSpecAzureMonitorMetrics
		return ret
	}).(DestinationsSpecAzureMonitorMetricsOutput)
}

func (o DestinationsSpecAzureMonitorMetricsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationsSpecAzureMonitorMetrics) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type DestinationsSpecResponseAzureMonitorMetrics struct {
	Name *string `pulumi:"name"`
}

type DestinationsSpecResponseAzureMonitorMetricsOutput struct{ *pulumi.OutputState }

func (DestinationsSpecResponseAzureMonitorMetricsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationsSpecResponseAzureMonitorMetrics)(nil)).Elem()
}

func (o DestinationsSpecResponseAzureMonitorMetricsOutput) ToDestinationsSpecResponseAzureMonitorMetricsOutput() DestinationsSpecResponseAzureMonitorMetricsOutput {
	return o
}

func (o DestinationsSpecResponseAzureMonitorMetricsOutput) ToDestinationsSpecResponseAzureMonitorMetricsOutputWithContext(ctx context.Context) DestinationsSpecResponseAzureMonitorMetricsOutput {
	return o
}

func (o DestinationsSpecResponseAzureMonitorMetricsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DestinationsSpecResponseAzureMonitorMetrics) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DestinationsSpecResponseAzureMonitorMetricsPtrOutput struct{ *pulumi.OutputState }

func (DestinationsSpecResponseAzureMonitorMetricsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationsSpecResponseAzureMonitorMetrics)(nil)).Elem()
}

func (o DestinationsSpecResponseAzureMonitorMetricsPtrOutput) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutput() DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o
}

func (o DestinationsSpecResponseAzureMonitorMetricsPtrOutput) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o
}

func (o DestinationsSpecResponseAzureMonitorMetricsPtrOutput) Elem() DestinationsSpecResponseAzureMonitorMetricsOutput {
	return o.ApplyT(func(v *DestinationsSpecResponseAzureMonitorMetrics) DestinationsSpecResponseAzureMonitorMetrics {
		if v != nil {
			return *v
		}
		var ret DestinationsSpecResponseAzureMonitorMetrics
		return ret
	}).(DestinationsSpecResponseAzureMonitorMetricsOutput)
}

func (o DestinationsSpecResponseAzureMonitorMetricsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DestinationsSpecResponseAzureMonitorMetrics) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type Dimension struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type DimensionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type DynamicMetricCriteria struct {
	AlertSensitivity     string                         `pulumi:"alertSensitivity"`
	CriterionType        string                         `pulumi:"criterionType"`
	Dimensions           []MetricDimension              `pulumi:"dimensions"`
	FailingPeriods       DynamicThresholdFailingPeriods `pulumi:"failingPeriods"`
	IgnoreDataBefore     *string                        `pulumi:"ignoreDataBefore"`
	MetricName           string                         `pulumi:"metricName"`
	MetricNamespace      *string                        `pulumi:"metricNamespace"`
	Name                 string                         `pulumi:"name"`
	Operator             string                         `pulumi:"operator"`
	SkipMetricValidation *bool                          `pulumi:"skipMetricValidation"`
	TimeAggregation      string                         `pulumi:"timeAggregation"`
}

type DynamicMetricCriteriaResponse struct {
	AlertSensitivity     string                                 `pulumi:"alertSensitivity"`
	CriterionType        string                                 `pulumi:"criterionType"`
	Dimensions           []MetricDimensionResponse              `pulumi:"dimensions"`
	FailingPeriods       DynamicThresholdFailingPeriodsResponse `pulumi:"failingPeriods"`
	IgnoreDataBefore     *string                                `pulumi:"ignoreDataBefore"`
	MetricName           string                                 `pulumi:"metricName"`
	MetricNamespace      *string                                `pulumi:"metricNamespace"`
	Name                 string                                 `pulumi:"name"`
	Operator             string                                 `pulumi:"operator"`
	SkipMetricValidation *bool                                  `pulumi:"skipMetricValidation"`
	TimeAggregation      string                                 `pulumi:"timeAggregation"`
}

type DynamicThresholdFailingPeriods struct {
	MinFailingPeriodsToAlert  float64 `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods float64 `pulumi:"numberOfEvaluationPeriods"`
}

type DynamicThresholdFailingPeriodsResponse struct {
	MinFailingPeriodsToAlert  float64 `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods float64 `pulumi:"numberOfEvaluationPeriods"`
}

type EmailNotification struct {
	CustomEmails                       []string `pulumi:"customEmails"`
	SendToSubscriptionAdministrator    *bool    `pulumi:"sendToSubscriptionAdministrator"`
	SendToSubscriptionCoAdministrators *bool    `pulumi:"sendToSubscriptionCoAdministrators"`
}


func (val *EmailNotification) Defaults() *EmailNotification {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SendToSubscriptionAdministrator) {
		sendToSubscriptionAdministrator_ := false
		tmp.SendToSubscriptionAdministrator = &sendToSubscriptionAdministrator_
	}
	if isZero(tmp.SendToSubscriptionCoAdministrators) {
		sendToSubscriptionCoAdministrators_ := false
		tmp.SendToSubscriptionCoAdministrators = &sendToSubscriptionCoAdministrators_
	}
	return &tmp
}





type EmailNotificationInput interface {
	pulumi.Input

	ToEmailNotificationOutput() EmailNotificationOutput
	ToEmailNotificationOutputWithContext(context.Context) EmailNotificationOutput
}

type EmailNotificationArgs struct {
	CustomEmails                       pulumi.StringArrayInput `pulumi:"customEmails"`
	SendToSubscriptionAdministrator    pulumi.BoolPtrInput     `pulumi:"sendToSubscriptionAdministrator"`
	SendToSubscriptionCoAdministrators pulumi.BoolPtrInput     `pulumi:"sendToSubscriptionCoAdministrators"`
}


func (val *EmailNotificationArgs) Defaults() *EmailNotificationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SendToSubscriptionAdministrator) {
		tmp.SendToSubscriptionAdministrator = pulumi.BoolPtr(false)
	}
	if isZero(tmp.SendToSubscriptionCoAdministrators) {
		tmp.SendToSubscriptionCoAdministrators = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (EmailNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailNotification)(nil)).Elem()
}

func (i EmailNotificationArgs) ToEmailNotificationOutput() EmailNotificationOutput {
	return i.ToEmailNotificationOutputWithContext(context.Background())
}

func (i EmailNotificationArgs) ToEmailNotificationOutputWithContext(ctx context.Context) EmailNotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailNotificationOutput)
}

func (i EmailNotificationArgs) ToEmailNotificationPtrOutput() EmailNotificationPtrOutput {
	return i.ToEmailNotificationPtrOutputWithContext(context.Background())
}

func (i EmailNotificationArgs) ToEmailNotificationPtrOutputWithContext(ctx context.Context) EmailNotificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailNotificationOutput).ToEmailNotificationPtrOutputWithContext(ctx)
}









type EmailNotificationPtrInput interface {
	pulumi.Input

	ToEmailNotificationPtrOutput() EmailNotificationPtrOutput
	ToEmailNotificationPtrOutputWithContext(context.Context) EmailNotificationPtrOutput
}

type emailNotificationPtrType EmailNotificationArgs

func EmailNotificationPtr(v *EmailNotificationArgs) EmailNotificationPtrInput {
	return (*emailNotificationPtrType)(v)
}

func (*emailNotificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailNotification)(nil)).Elem()
}

func (i *emailNotificationPtrType) ToEmailNotificationPtrOutput() EmailNotificationPtrOutput {
	return i.ToEmailNotificationPtrOutputWithContext(context.Background())
}

func (i *emailNotificationPtrType) ToEmailNotificationPtrOutputWithContext(ctx context.Context) EmailNotificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailNotificationPtrOutput)
}

type EmailNotificationOutput struct{ *pulumi.OutputState }

func (EmailNotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailNotification)(nil)).Elem()
}

func (o EmailNotificationOutput) ToEmailNotificationOutput() EmailNotificationOutput {
	return o
}

func (o EmailNotificationOutput) ToEmailNotificationOutputWithContext(ctx context.Context) EmailNotificationOutput {
	return o
}

func (o EmailNotificationOutput) ToEmailNotificationPtrOutput() EmailNotificationPtrOutput {
	return o.ToEmailNotificationPtrOutputWithContext(context.Background())
}

func (o EmailNotificationOutput) ToEmailNotificationPtrOutputWithContext(ctx context.Context) EmailNotificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EmailNotification) *EmailNotification {
		return &v
	}).(EmailNotificationPtrOutput)
}

func (o EmailNotificationOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EmailNotification) []string { return v.CustomEmails }).(pulumi.StringArrayOutput)
}

func (o EmailNotificationOutput) SendToSubscriptionAdministrator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EmailNotification) *bool { return v.SendToSubscriptionAdministrator }).(pulumi.BoolPtrOutput)
}

func (o EmailNotificationOutput) SendToSubscriptionCoAdministrators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EmailNotification) *bool { return v.SendToSubscriptionCoAdministrators }).(pulumi.BoolPtrOutput)
}

type EmailNotificationPtrOutput struct{ *pulumi.OutputState }

func (EmailNotificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailNotification)(nil)).Elem()
}

func (o EmailNotificationPtrOutput) ToEmailNotificationPtrOutput() EmailNotificationPtrOutput {
	return o
}

func (o EmailNotificationPtrOutput) ToEmailNotificationPtrOutputWithContext(ctx context.Context) EmailNotificationPtrOutput {
	return o
}

func (o EmailNotificationPtrOutput) Elem() EmailNotificationOutput {
	return o.ApplyT(func(v *EmailNotification) EmailNotification {
		if v != nil {
			return *v
		}
		var ret EmailNotification
		return ret
	}).(EmailNotificationOutput)
}

func (o EmailNotificationPtrOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EmailNotification) []string {
		if v == nil {
			return nil
		}
		return v.CustomEmails
	}).(pulumi.StringArrayOutput)
}

func (o EmailNotificationPtrOutput) SendToSubscriptionAdministrator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EmailNotification) *bool {
		if v == nil {
			return nil
		}
		return v.SendToSubscriptionAdministrator
	}).(pulumi.BoolPtrOutput)
}

func (o EmailNotificationPtrOutput) SendToSubscriptionCoAdministrators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EmailNotification) *bool {
		if v == nil {
			return nil
		}
		return v.SendToSubscriptionCoAdministrators
	}).(pulumi.BoolPtrOutput)
}

type EmailNotificationResponse struct {
	CustomEmails                       []string `pulumi:"customEmails"`
	SendToSubscriptionAdministrator    *bool    `pulumi:"sendToSubscriptionAdministrator"`
	SendToSubscriptionCoAdministrators *bool    `pulumi:"sendToSubscriptionCoAdministrators"`
}


func (val *EmailNotificationResponse) Defaults() *EmailNotificationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SendToSubscriptionAdministrator) {
		sendToSubscriptionAdministrator_ := false
		tmp.SendToSubscriptionAdministrator = &sendToSubscriptionAdministrator_
	}
	if isZero(tmp.SendToSubscriptionCoAdministrators) {
		sendToSubscriptionCoAdministrators_ := false
		tmp.SendToSubscriptionCoAdministrators = &sendToSubscriptionCoAdministrators_
	}
	return &tmp
}

type EmailNotificationResponseOutput struct{ *pulumi.OutputState }

func (EmailNotificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailNotificationResponse)(nil)).Elem()
}

func (o EmailNotificationResponseOutput) ToEmailNotificationResponseOutput() EmailNotificationResponseOutput {
	return o
}

func (o EmailNotificationResponseOutput) ToEmailNotificationResponseOutputWithContext(ctx context.Context) EmailNotificationResponseOutput {
	return o
}

func (o EmailNotificationResponseOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EmailNotificationResponse) []string { return v.CustomEmails }).(pulumi.StringArrayOutput)
}

func (o EmailNotificationResponseOutput) SendToSubscriptionAdministrator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EmailNotificationResponse) *bool { return v.SendToSubscriptionAdministrator }).(pulumi.BoolPtrOutput)
}

func (o EmailNotificationResponseOutput) SendToSubscriptionCoAdministrators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EmailNotificationResponse) *bool { return v.SendToSubscriptionCoAdministrators }).(pulumi.BoolPtrOutput)
}

type EmailNotificationResponsePtrOutput struct{ *pulumi.OutputState }

func (EmailNotificationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailNotificationResponse)(nil)).Elem()
}

func (o EmailNotificationResponsePtrOutput) ToEmailNotificationResponsePtrOutput() EmailNotificationResponsePtrOutput {
	return o
}

func (o EmailNotificationResponsePtrOutput) ToEmailNotificationResponsePtrOutputWithContext(ctx context.Context) EmailNotificationResponsePtrOutput {
	return o
}

func (o EmailNotificationResponsePtrOutput) Elem() EmailNotificationResponseOutput {
	return o.ApplyT(func(v *EmailNotificationResponse) EmailNotificationResponse {
		if v != nil {
			return *v
		}
		var ret EmailNotificationResponse
		return ret
	}).(EmailNotificationResponseOutput)
}

func (o EmailNotificationResponsePtrOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EmailNotificationResponse) []string {
		if v == nil {
			return nil
		}
		return v.CustomEmails
	}).(pulumi.StringArrayOutput)
}

func (o EmailNotificationResponsePtrOutput) SendToSubscriptionAdministrator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EmailNotificationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendToSubscriptionAdministrator
	}).(pulumi.BoolPtrOutput)
}

func (o EmailNotificationResponsePtrOutput) SendToSubscriptionCoAdministrators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EmailNotificationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SendToSubscriptionCoAdministrators
	}).(pulumi.BoolPtrOutput)
}

type EmailReceiver struct {
	EmailAddress         string `pulumi:"emailAddress"`
	Name                 string `pulumi:"name"`
	UseCommonAlertSchema *bool  `pulumi:"useCommonAlertSchema"`
}


func (val *EmailReceiver) Defaults() *EmailReceiver {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}





type EmailReceiverInput interface {
	pulumi.Input

	ToEmailReceiverOutput() EmailReceiverOutput
	ToEmailReceiverOutputWithContext(context.Context) EmailReceiverOutput
}

type EmailReceiverArgs struct {
	EmailAddress         pulumi.StringInput  `pulumi:"emailAddress"`
	Name                 pulumi.StringInput  `pulumi:"name"`
	UseCommonAlertSchema pulumi.BoolPtrInput `pulumi:"useCommonAlertSchema"`
}


func (val *EmailReceiverArgs) Defaults() *EmailReceiverArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		tmp.UseCommonAlertSchema = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (EmailReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailReceiver)(nil)).Elem()
}

func (i EmailReceiverArgs) ToEmailReceiverOutput() EmailReceiverOutput {
	return i.ToEmailReceiverOutputWithContext(context.Background())
}

func (i EmailReceiverArgs) ToEmailReceiverOutputWithContext(ctx context.Context) EmailReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailReceiverOutput)
}





type EmailReceiverArrayInput interface {
	pulumi.Input

	ToEmailReceiverArrayOutput() EmailReceiverArrayOutput
	ToEmailReceiverArrayOutputWithContext(context.Context) EmailReceiverArrayOutput
}

type EmailReceiverArray []EmailReceiverInput

func (EmailReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailReceiver)(nil)).Elem()
}

func (i EmailReceiverArray) ToEmailReceiverArrayOutput() EmailReceiverArrayOutput {
	return i.ToEmailReceiverArrayOutputWithContext(context.Background())
}

func (i EmailReceiverArray) ToEmailReceiverArrayOutputWithContext(ctx context.Context) EmailReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailReceiverArrayOutput)
}

type EmailReceiverOutput struct{ *pulumi.OutputState }

func (EmailReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailReceiver)(nil)).Elem()
}

func (o EmailReceiverOutput) ToEmailReceiverOutput() EmailReceiverOutput {
	return o
}

func (o EmailReceiverOutput) ToEmailReceiverOutputWithContext(ctx context.Context) EmailReceiverOutput {
	return o
}

func (o EmailReceiverOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v EmailReceiver) string { return v.EmailAddress }).(pulumi.StringOutput)
}

func (o EmailReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EmailReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o EmailReceiverOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EmailReceiver) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type EmailReceiverArrayOutput struct{ *pulumi.OutputState }

func (EmailReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailReceiver)(nil)).Elem()
}

func (o EmailReceiverArrayOutput) ToEmailReceiverArrayOutput() EmailReceiverArrayOutput {
	return o
}

func (o EmailReceiverArrayOutput) ToEmailReceiverArrayOutputWithContext(ctx context.Context) EmailReceiverArrayOutput {
	return o
}

func (o EmailReceiverArrayOutput) Index(i pulumi.IntInput) EmailReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmailReceiver {
		return vs[0].([]EmailReceiver)[vs[1].(int)]
	}).(EmailReceiverOutput)
}

type EmailReceiverResponse struct {
	EmailAddress         string `pulumi:"emailAddress"`
	Name                 string `pulumi:"name"`
	Status               string `pulumi:"status"`
	UseCommonAlertSchema *bool  `pulumi:"useCommonAlertSchema"`
}


func (val *EmailReceiverResponse) Defaults() *EmailReceiverResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}

type EmailReceiverResponseOutput struct{ *pulumi.OutputState }

func (EmailReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailReceiverResponse)(nil)).Elem()
}

func (o EmailReceiverResponseOutput) ToEmailReceiverResponseOutput() EmailReceiverResponseOutput {
	return o
}

func (o EmailReceiverResponseOutput) ToEmailReceiverResponseOutputWithContext(ctx context.Context) EmailReceiverResponseOutput {
	return o
}

func (o EmailReceiverResponseOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v EmailReceiverResponse) string { return v.EmailAddress }).(pulumi.StringOutput)
}

func (o EmailReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EmailReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EmailReceiverResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v EmailReceiverResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o EmailReceiverResponseOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EmailReceiverResponse) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type EmailReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (EmailReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailReceiverResponse)(nil)).Elem()
}

func (o EmailReceiverResponseArrayOutput) ToEmailReceiverResponseArrayOutput() EmailReceiverResponseArrayOutput {
	return o
}

func (o EmailReceiverResponseArrayOutput) ToEmailReceiverResponseArrayOutputWithContext(ctx context.Context) EmailReceiverResponseArrayOutput {
	return o
}

func (o EmailReceiverResponseArrayOutput) Index(i pulumi.IntInput) EmailReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EmailReceiverResponse {
		return vs[0].([]EmailReceiverResponse)[vs[1].(int)]
	}).(EmailReceiverResponseOutput)
}

type EtwEventConfiguration struct {
	Filter *string `pulumi:"filter"`
	Id     int     `pulumi:"id"`
	Name   string  `pulumi:"name"`
}





type EtwEventConfigurationInput interface {
	pulumi.Input

	ToEtwEventConfigurationOutput() EtwEventConfigurationOutput
	ToEtwEventConfigurationOutputWithContext(context.Context) EtwEventConfigurationOutput
}

type EtwEventConfigurationArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	Id     pulumi.IntInput       `pulumi:"id"`
	Name   pulumi.StringInput    `pulumi:"name"`
}

func (EtwEventConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EtwEventConfiguration)(nil)).Elem()
}

func (i EtwEventConfigurationArgs) ToEtwEventConfigurationOutput() EtwEventConfigurationOutput {
	return i.ToEtwEventConfigurationOutputWithContext(context.Background())
}

func (i EtwEventConfigurationArgs) ToEtwEventConfigurationOutputWithContext(ctx context.Context) EtwEventConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtwEventConfigurationOutput)
}





type EtwEventConfigurationArrayInput interface {
	pulumi.Input

	ToEtwEventConfigurationArrayOutput() EtwEventConfigurationArrayOutput
	ToEtwEventConfigurationArrayOutputWithContext(context.Context) EtwEventConfigurationArrayOutput
}

type EtwEventConfigurationArray []EtwEventConfigurationInput

func (EtwEventConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtwEventConfiguration)(nil)).Elem()
}

func (i EtwEventConfigurationArray) ToEtwEventConfigurationArrayOutput() EtwEventConfigurationArrayOutput {
	return i.ToEtwEventConfigurationArrayOutputWithContext(context.Background())
}

func (i EtwEventConfigurationArray) ToEtwEventConfigurationArrayOutputWithContext(ctx context.Context) EtwEventConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtwEventConfigurationArrayOutput)
}

type EtwEventConfigurationOutput struct{ *pulumi.OutputState }

func (EtwEventConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EtwEventConfiguration)(nil)).Elem()
}

func (o EtwEventConfigurationOutput) ToEtwEventConfigurationOutput() EtwEventConfigurationOutput {
	return o
}

func (o EtwEventConfigurationOutput) ToEtwEventConfigurationOutputWithContext(ctx context.Context) EtwEventConfigurationOutput {
	return o
}

func (o EtwEventConfigurationOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EtwEventConfiguration) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o EtwEventConfigurationOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v EtwEventConfiguration) int { return v.Id }).(pulumi.IntOutput)
}

func (o EtwEventConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EtwEventConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

type EtwEventConfigurationArrayOutput struct{ *pulumi.OutputState }

func (EtwEventConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtwEventConfiguration)(nil)).Elem()
}

func (o EtwEventConfigurationArrayOutput) ToEtwEventConfigurationArrayOutput() EtwEventConfigurationArrayOutput {
	return o
}

func (o EtwEventConfigurationArrayOutput) ToEtwEventConfigurationArrayOutputWithContext(ctx context.Context) EtwEventConfigurationArrayOutput {
	return o
}

func (o EtwEventConfigurationArrayOutput) Index(i pulumi.IntInput) EtwEventConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EtwEventConfiguration {
		return vs[0].([]EtwEventConfiguration)[vs[1].(int)]
	}).(EtwEventConfigurationOutput)
}

type EtwEventConfigurationResponse struct {
	Filter *string `pulumi:"filter"`
	Id     int     `pulumi:"id"`
	Name   string  `pulumi:"name"`
}

type EtwEventConfigurationResponseOutput struct{ *pulumi.OutputState }

func (EtwEventConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EtwEventConfigurationResponse)(nil)).Elem()
}

func (o EtwEventConfigurationResponseOutput) ToEtwEventConfigurationResponseOutput() EtwEventConfigurationResponseOutput {
	return o
}

func (o EtwEventConfigurationResponseOutput) ToEtwEventConfigurationResponseOutputWithContext(ctx context.Context) EtwEventConfigurationResponseOutput {
	return o
}

func (o EtwEventConfigurationResponseOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EtwEventConfigurationResponse) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o EtwEventConfigurationResponseOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v EtwEventConfigurationResponse) int { return v.Id }).(pulumi.IntOutput)
}

func (o EtwEventConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EtwEventConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type EtwEventConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (EtwEventConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtwEventConfigurationResponse)(nil)).Elem()
}

func (o EtwEventConfigurationResponseArrayOutput) ToEtwEventConfigurationResponseArrayOutput() EtwEventConfigurationResponseArrayOutput {
	return o
}

func (o EtwEventConfigurationResponseArrayOutput) ToEtwEventConfigurationResponseArrayOutputWithContext(ctx context.Context) EtwEventConfigurationResponseArrayOutput {
	return o
}

func (o EtwEventConfigurationResponseArrayOutput) Index(i pulumi.IntInput) EtwEventConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EtwEventConfigurationResponse {
		return vs[0].([]EtwEventConfigurationResponse)[vs[1].(int)]
	}).(EtwEventConfigurationResponseOutput)
}

type EtwProviderConfiguration struct {
	Events []EtwEventConfiguration `pulumi:"events"`
	Id     string                  `pulumi:"id"`
}





type EtwProviderConfigurationInput interface {
	pulumi.Input

	ToEtwProviderConfigurationOutput() EtwProviderConfigurationOutput
	ToEtwProviderConfigurationOutputWithContext(context.Context) EtwProviderConfigurationOutput
}

type EtwProviderConfigurationArgs struct {
	Events EtwEventConfigurationArrayInput `pulumi:"events"`
	Id     pulumi.StringInput              `pulumi:"id"`
}

func (EtwProviderConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EtwProviderConfiguration)(nil)).Elem()
}

func (i EtwProviderConfigurationArgs) ToEtwProviderConfigurationOutput() EtwProviderConfigurationOutput {
	return i.ToEtwProviderConfigurationOutputWithContext(context.Background())
}

func (i EtwProviderConfigurationArgs) ToEtwProviderConfigurationOutputWithContext(ctx context.Context) EtwProviderConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtwProviderConfigurationOutput)
}





type EtwProviderConfigurationArrayInput interface {
	pulumi.Input

	ToEtwProviderConfigurationArrayOutput() EtwProviderConfigurationArrayOutput
	ToEtwProviderConfigurationArrayOutputWithContext(context.Context) EtwProviderConfigurationArrayOutput
}

type EtwProviderConfigurationArray []EtwProviderConfigurationInput

func (EtwProviderConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtwProviderConfiguration)(nil)).Elem()
}

func (i EtwProviderConfigurationArray) ToEtwProviderConfigurationArrayOutput() EtwProviderConfigurationArrayOutput {
	return i.ToEtwProviderConfigurationArrayOutputWithContext(context.Background())
}

func (i EtwProviderConfigurationArray) ToEtwProviderConfigurationArrayOutputWithContext(ctx context.Context) EtwProviderConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtwProviderConfigurationArrayOutput)
}

type EtwProviderConfigurationOutput struct{ *pulumi.OutputState }

func (EtwProviderConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EtwProviderConfiguration)(nil)).Elem()
}

func (o EtwProviderConfigurationOutput) ToEtwProviderConfigurationOutput() EtwProviderConfigurationOutput {
	return o
}

func (o EtwProviderConfigurationOutput) ToEtwProviderConfigurationOutputWithContext(ctx context.Context) EtwProviderConfigurationOutput {
	return o
}

func (o EtwProviderConfigurationOutput) Events() EtwEventConfigurationArrayOutput {
	return o.ApplyT(func(v EtwProviderConfiguration) []EtwEventConfiguration { return v.Events }).(EtwEventConfigurationArrayOutput)
}

func (o EtwProviderConfigurationOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v EtwProviderConfiguration) string { return v.Id }).(pulumi.StringOutput)
}

type EtwProviderConfigurationArrayOutput struct{ *pulumi.OutputState }

func (EtwProviderConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtwProviderConfiguration)(nil)).Elem()
}

func (o EtwProviderConfigurationArrayOutput) ToEtwProviderConfigurationArrayOutput() EtwProviderConfigurationArrayOutput {
	return o
}

func (o EtwProviderConfigurationArrayOutput) ToEtwProviderConfigurationArrayOutputWithContext(ctx context.Context) EtwProviderConfigurationArrayOutput {
	return o
}

func (o EtwProviderConfigurationArrayOutput) Index(i pulumi.IntInput) EtwProviderConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EtwProviderConfiguration {
		return vs[0].([]EtwProviderConfiguration)[vs[1].(int)]
	}).(EtwProviderConfigurationOutput)
}

type EtwProviderConfigurationResponse struct {
	Events []EtwEventConfigurationResponse `pulumi:"events"`
	Id     string                          `pulumi:"id"`
}

type EtwProviderConfigurationResponseOutput struct{ *pulumi.OutputState }

func (EtwProviderConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EtwProviderConfigurationResponse)(nil)).Elem()
}

func (o EtwProviderConfigurationResponseOutput) ToEtwProviderConfigurationResponseOutput() EtwProviderConfigurationResponseOutput {
	return o
}

func (o EtwProviderConfigurationResponseOutput) ToEtwProviderConfigurationResponseOutputWithContext(ctx context.Context) EtwProviderConfigurationResponseOutput {
	return o
}

func (o EtwProviderConfigurationResponseOutput) Events() EtwEventConfigurationResponseArrayOutput {
	return o.ApplyT(func(v EtwProviderConfigurationResponse) []EtwEventConfigurationResponse { return v.Events }).(EtwEventConfigurationResponseArrayOutput)
}

func (o EtwProviderConfigurationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v EtwProviderConfigurationResponse) string { return v.Id }).(pulumi.StringOutput)
}

type EtwProviderConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (EtwProviderConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtwProviderConfigurationResponse)(nil)).Elem()
}

func (o EtwProviderConfigurationResponseArrayOutput) ToEtwProviderConfigurationResponseArrayOutput() EtwProviderConfigurationResponseArrayOutput {
	return o
}

func (o EtwProviderConfigurationResponseArrayOutput) ToEtwProviderConfigurationResponseArrayOutputWithContext(ctx context.Context) EtwProviderConfigurationResponseArrayOutput {
	return o
}

func (o EtwProviderConfigurationResponseArrayOutput) Index(i pulumi.IntInput) EtwProviderConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EtwProviderConfigurationResponse {
		return vs[0].([]EtwProviderConfigurationResponse)[vs[1].(int)]
	}).(EtwProviderConfigurationResponseOutput)
}

type EventLogConfiguration struct {
	Filter  *string `pulumi:"filter"`
	LogName string  `pulumi:"logName"`
}





type EventLogConfigurationInput interface {
	pulumi.Input

	ToEventLogConfigurationOutput() EventLogConfigurationOutput
	ToEventLogConfigurationOutputWithContext(context.Context) EventLogConfigurationOutput
}

type EventLogConfigurationArgs struct {
	Filter  pulumi.StringPtrInput `pulumi:"filter"`
	LogName pulumi.StringInput    `pulumi:"logName"`
}

func (EventLogConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventLogConfiguration)(nil)).Elem()
}

func (i EventLogConfigurationArgs) ToEventLogConfigurationOutput() EventLogConfigurationOutput {
	return i.ToEventLogConfigurationOutputWithContext(context.Background())
}

func (i EventLogConfigurationArgs) ToEventLogConfigurationOutputWithContext(ctx context.Context) EventLogConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventLogConfigurationOutput)
}





type EventLogConfigurationArrayInput interface {
	pulumi.Input

	ToEventLogConfigurationArrayOutput() EventLogConfigurationArrayOutput
	ToEventLogConfigurationArrayOutputWithContext(context.Context) EventLogConfigurationArrayOutput
}

type EventLogConfigurationArray []EventLogConfigurationInput

func (EventLogConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventLogConfiguration)(nil)).Elem()
}

func (i EventLogConfigurationArray) ToEventLogConfigurationArrayOutput() EventLogConfigurationArrayOutput {
	return i.ToEventLogConfigurationArrayOutputWithContext(context.Background())
}

func (i EventLogConfigurationArray) ToEventLogConfigurationArrayOutputWithContext(ctx context.Context) EventLogConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventLogConfigurationArrayOutput)
}

type EventLogConfigurationOutput struct{ *pulumi.OutputState }

func (EventLogConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventLogConfiguration)(nil)).Elem()
}

func (o EventLogConfigurationOutput) ToEventLogConfigurationOutput() EventLogConfigurationOutput {
	return o
}

func (o EventLogConfigurationOutput) ToEventLogConfigurationOutputWithContext(ctx context.Context) EventLogConfigurationOutput {
	return o
}

func (o EventLogConfigurationOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventLogConfiguration) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o EventLogConfigurationOutput) LogName() pulumi.StringOutput {
	return o.ApplyT(func(v EventLogConfiguration) string { return v.LogName }).(pulumi.StringOutput)
}

type EventLogConfigurationArrayOutput struct{ *pulumi.OutputState }

func (EventLogConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventLogConfiguration)(nil)).Elem()
}

func (o EventLogConfigurationArrayOutput) ToEventLogConfigurationArrayOutput() EventLogConfigurationArrayOutput {
	return o
}

func (o EventLogConfigurationArrayOutput) ToEventLogConfigurationArrayOutputWithContext(ctx context.Context) EventLogConfigurationArrayOutput {
	return o
}

func (o EventLogConfigurationArrayOutput) Index(i pulumi.IntInput) EventLogConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventLogConfiguration {
		return vs[0].([]EventLogConfiguration)[vs[1].(int)]
	}).(EventLogConfigurationOutput)
}

type EventLogConfigurationResponse struct {
	Filter  *string `pulumi:"filter"`
	LogName string  `pulumi:"logName"`
}

type EventLogConfigurationResponseOutput struct{ *pulumi.OutputState }

func (EventLogConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventLogConfigurationResponse)(nil)).Elem()
}

func (o EventLogConfigurationResponseOutput) ToEventLogConfigurationResponseOutput() EventLogConfigurationResponseOutput {
	return o
}

func (o EventLogConfigurationResponseOutput) ToEventLogConfigurationResponseOutputWithContext(ctx context.Context) EventLogConfigurationResponseOutput {
	return o
}

func (o EventLogConfigurationResponseOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventLogConfigurationResponse) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o EventLogConfigurationResponseOutput) LogName() pulumi.StringOutput {
	return o.ApplyT(func(v EventLogConfigurationResponse) string { return v.LogName }).(pulumi.StringOutput)
}

type EventLogConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (EventLogConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventLogConfigurationResponse)(nil)).Elem()
}

func (o EventLogConfigurationResponseArrayOutput) ToEventLogConfigurationResponseArrayOutput() EventLogConfigurationResponseArrayOutput {
	return o
}

func (o EventLogConfigurationResponseArrayOutput) ToEventLogConfigurationResponseArrayOutputWithContext(ctx context.Context) EventLogConfigurationResponseArrayOutput {
	return o
}

func (o EventLogConfigurationResponseArrayOutput) Index(i pulumi.IntInput) EventLogConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventLogConfigurationResponse {
		return vs[0].([]EventLogConfigurationResponse)[vs[1].(int)]
	}).(EventLogConfigurationResponseOutput)
}

type ExtensionDataSource struct {
	ExtensionName     string      `pulumi:"extensionName"`
	ExtensionSettings interface{} `pulumi:"extensionSettings"`
	InputDataSources  []string    `pulumi:"inputDataSources"`
	Name              *string     `pulumi:"name"`
	Streams           []string    `pulumi:"streams"`
}





type ExtensionDataSourceInput interface {
	pulumi.Input

	ToExtensionDataSourceOutput() ExtensionDataSourceOutput
	ToExtensionDataSourceOutputWithContext(context.Context) ExtensionDataSourceOutput
}

type ExtensionDataSourceArgs struct {
	ExtensionName     pulumi.StringInput      `pulumi:"extensionName"`
	ExtensionSettings pulumi.Input            `pulumi:"extensionSettings"`
	InputDataSources  pulumi.StringArrayInput `pulumi:"inputDataSources"`
	Name              pulumi.StringPtrInput   `pulumi:"name"`
	Streams           pulumi.StringArrayInput `pulumi:"streams"`
}

func (ExtensionDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionDataSource)(nil)).Elem()
}

func (i ExtensionDataSourceArgs) ToExtensionDataSourceOutput() ExtensionDataSourceOutput {
	return i.ToExtensionDataSourceOutputWithContext(context.Background())
}

func (i ExtensionDataSourceArgs) ToExtensionDataSourceOutputWithContext(ctx context.Context) ExtensionDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionDataSourceOutput)
}





type ExtensionDataSourceArrayInput interface {
	pulumi.Input

	ToExtensionDataSourceArrayOutput() ExtensionDataSourceArrayOutput
	ToExtensionDataSourceArrayOutputWithContext(context.Context) ExtensionDataSourceArrayOutput
}

type ExtensionDataSourceArray []ExtensionDataSourceInput

func (ExtensionDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionDataSource)(nil)).Elem()
}

func (i ExtensionDataSourceArray) ToExtensionDataSourceArrayOutput() ExtensionDataSourceArrayOutput {
	return i.ToExtensionDataSourceArrayOutputWithContext(context.Background())
}

func (i ExtensionDataSourceArray) ToExtensionDataSourceArrayOutputWithContext(ctx context.Context) ExtensionDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionDataSourceArrayOutput)
}

type ExtensionDataSourceOutput struct{ *pulumi.OutputState }

func (ExtensionDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionDataSource)(nil)).Elem()
}

func (o ExtensionDataSourceOutput) ToExtensionDataSourceOutput() ExtensionDataSourceOutput {
	return o
}

func (o ExtensionDataSourceOutput) ToExtensionDataSourceOutputWithContext(ctx context.Context) ExtensionDataSourceOutput {
	return o
}

func (o ExtensionDataSourceOutput) ExtensionName() pulumi.StringOutput {
	return o.ApplyT(func(v ExtensionDataSource) string { return v.ExtensionName }).(pulumi.StringOutput)
}

func (o ExtensionDataSourceOutput) ExtensionSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v ExtensionDataSource) interface{} { return v.ExtensionSettings }).(pulumi.AnyOutput)
}

func (o ExtensionDataSourceOutput) InputDataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExtensionDataSource) []string { return v.InputDataSources }).(pulumi.StringArrayOutput)
}

func (o ExtensionDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtensionDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExtensionDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type ExtensionDataSourceArrayOutput struct{ *pulumi.OutputState }

func (ExtensionDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionDataSource)(nil)).Elem()
}

func (o ExtensionDataSourceArrayOutput) ToExtensionDataSourceArrayOutput() ExtensionDataSourceArrayOutput {
	return o
}

func (o ExtensionDataSourceArrayOutput) ToExtensionDataSourceArrayOutputWithContext(ctx context.Context) ExtensionDataSourceArrayOutput {
	return o
}

func (o ExtensionDataSourceArrayOutput) Index(i pulumi.IntInput) ExtensionDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionDataSource {
		return vs[0].([]ExtensionDataSource)[vs[1].(int)]
	}).(ExtensionDataSourceOutput)
}

type ExtensionDataSourceResponse struct {
	ExtensionName     string      `pulumi:"extensionName"`
	ExtensionSettings interface{} `pulumi:"extensionSettings"`
	InputDataSources  []string    `pulumi:"inputDataSources"`
	Name              *string     `pulumi:"name"`
	Streams           []string    `pulumi:"streams"`
}

type ExtensionDataSourceResponseOutput struct{ *pulumi.OutputState }

func (ExtensionDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionDataSourceResponse)(nil)).Elem()
}

func (o ExtensionDataSourceResponseOutput) ToExtensionDataSourceResponseOutput() ExtensionDataSourceResponseOutput {
	return o
}

func (o ExtensionDataSourceResponseOutput) ToExtensionDataSourceResponseOutputWithContext(ctx context.Context) ExtensionDataSourceResponseOutput {
	return o
}

func (o ExtensionDataSourceResponseOutput) ExtensionName() pulumi.StringOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) string { return v.ExtensionName }).(pulumi.StringOutput)
}

func (o ExtensionDataSourceResponseOutput) ExtensionSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) interface{} { return v.ExtensionSettings }).(pulumi.AnyOutput)
}

func (o ExtensionDataSourceResponseOutput) InputDataSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) []string { return v.InputDataSources }).(pulumi.StringArrayOutput)
}

func (o ExtensionDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtensionDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExtensionDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type ExtensionDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (ExtensionDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionDataSourceResponse)(nil)).Elem()
}

func (o ExtensionDataSourceResponseArrayOutput) ToExtensionDataSourceResponseArrayOutput() ExtensionDataSourceResponseArrayOutput {
	return o
}

func (o ExtensionDataSourceResponseArrayOutput) ToExtensionDataSourceResponseArrayOutputWithContext(ctx context.Context) ExtensionDataSourceResponseArrayOutput {
	return o
}

func (o ExtensionDataSourceResponseArrayOutput) Index(i pulumi.IntInput) ExtensionDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionDataSourceResponse {
		return vs[0].([]ExtensionDataSourceResponse)[vs[1].(int)]
	}).(ExtensionDataSourceResponseOutput)
}

type ItsmReceiver struct {
	ConnectionId        string `pulumi:"connectionId"`
	Name                string `pulumi:"name"`
	Region              string `pulumi:"region"`
	TicketConfiguration string `pulumi:"ticketConfiguration"`
	WorkspaceId         string `pulumi:"workspaceId"`
}





type ItsmReceiverInput interface {
	pulumi.Input

	ToItsmReceiverOutput() ItsmReceiverOutput
	ToItsmReceiverOutputWithContext(context.Context) ItsmReceiverOutput
}

type ItsmReceiverArgs struct {
	ConnectionId        pulumi.StringInput `pulumi:"connectionId"`
	Name                pulumi.StringInput `pulumi:"name"`
	Region              pulumi.StringInput `pulumi:"region"`
	TicketConfiguration pulumi.StringInput `pulumi:"ticketConfiguration"`
	WorkspaceId         pulumi.StringInput `pulumi:"workspaceId"`
}

func (ItsmReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ItsmReceiver)(nil)).Elem()
}

func (i ItsmReceiverArgs) ToItsmReceiverOutput() ItsmReceiverOutput {
	return i.ToItsmReceiverOutputWithContext(context.Background())
}

func (i ItsmReceiverArgs) ToItsmReceiverOutputWithContext(ctx context.Context) ItsmReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItsmReceiverOutput)
}





type ItsmReceiverArrayInput interface {
	pulumi.Input

	ToItsmReceiverArrayOutput() ItsmReceiverArrayOutput
	ToItsmReceiverArrayOutputWithContext(context.Context) ItsmReceiverArrayOutput
}

type ItsmReceiverArray []ItsmReceiverInput

func (ItsmReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ItsmReceiver)(nil)).Elem()
}

func (i ItsmReceiverArray) ToItsmReceiverArrayOutput() ItsmReceiverArrayOutput {
	return i.ToItsmReceiverArrayOutputWithContext(context.Background())
}

func (i ItsmReceiverArray) ToItsmReceiverArrayOutputWithContext(ctx context.Context) ItsmReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItsmReceiverArrayOutput)
}

type ItsmReceiverOutput struct{ *pulumi.OutputState }

func (ItsmReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ItsmReceiver)(nil)).Elem()
}

func (o ItsmReceiverOutput) ToItsmReceiverOutput() ItsmReceiverOutput {
	return o
}

func (o ItsmReceiverOutput) ToItsmReceiverOutputWithContext(ctx context.Context) ItsmReceiverOutput {
	return o
}

func (o ItsmReceiverOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiver) string { return v.ConnectionId }).(pulumi.StringOutput)
}

func (o ItsmReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o ItsmReceiverOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiver) string { return v.Region }).(pulumi.StringOutput)
}

func (o ItsmReceiverOutput) TicketConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiver) string { return v.TicketConfiguration }).(pulumi.StringOutput)
}

func (o ItsmReceiverOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiver) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

type ItsmReceiverArrayOutput struct{ *pulumi.OutputState }

func (ItsmReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ItsmReceiver)(nil)).Elem()
}

func (o ItsmReceiverArrayOutput) ToItsmReceiverArrayOutput() ItsmReceiverArrayOutput {
	return o
}

func (o ItsmReceiverArrayOutput) ToItsmReceiverArrayOutputWithContext(ctx context.Context) ItsmReceiverArrayOutput {
	return o
}

func (o ItsmReceiverArrayOutput) Index(i pulumi.IntInput) ItsmReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ItsmReceiver {
		return vs[0].([]ItsmReceiver)[vs[1].(int)]
	}).(ItsmReceiverOutput)
}

type ItsmReceiverResponse struct {
	ConnectionId        string `pulumi:"connectionId"`
	Name                string `pulumi:"name"`
	Region              string `pulumi:"region"`
	TicketConfiguration string `pulumi:"ticketConfiguration"`
	WorkspaceId         string `pulumi:"workspaceId"`
}

type ItsmReceiverResponseOutput struct{ *pulumi.OutputState }

func (ItsmReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ItsmReceiverResponse)(nil)).Elem()
}

func (o ItsmReceiverResponseOutput) ToItsmReceiverResponseOutput() ItsmReceiverResponseOutput {
	return o
}

func (o ItsmReceiverResponseOutput) ToItsmReceiverResponseOutputWithContext(ctx context.Context) ItsmReceiverResponseOutput {
	return o
}

func (o ItsmReceiverResponseOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiverResponse) string { return v.ConnectionId }).(pulumi.StringOutput)
}

func (o ItsmReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ItsmReceiverResponseOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiverResponse) string { return v.Region }).(pulumi.StringOutput)
}

func (o ItsmReceiverResponseOutput) TicketConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiverResponse) string { return v.TicketConfiguration }).(pulumi.StringOutput)
}

func (o ItsmReceiverResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v ItsmReceiverResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

type ItsmReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (ItsmReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ItsmReceiverResponse)(nil)).Elem()
}

func (o ItsmReceiverResponseArrayOutput) ToItsmReceiverResponseArrayOutput() ItsmReceiverResponseArrayOutput {
	return o
}

func (o ItsmReceiverResponseArrayOutput) ToItsmReceiverResponseArrayOutputWithContext(ctx context.Context) ItsmReceiverResponseArrayOutput {
	return o
}

func (o ItsmReceiverResponseArrayOutput) Index(i pulumi.IntInput) ItsmReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ItsmReceiverResponse {
		return vs[0].([]ItsmReceiverResponse)[vs[1].(int)]
	}).(ItsmReceiverResponseOutput)
}

type LocationThresholdRuleCondition struct {
	DataSource          interface{} `pulumi:"dataSource"`
	FailedLocationCount int         `pulumi:"failedLocationCount"`
	OdataType           string      `pulumi:"odataType"`
	WindowSize          *string     `pulumi:"windowSize"`
}

type LocationThresholdRuleConditionResponse struct {
	DataSource          interface{} `pulumi:"dataSource"`
	FailedLocationCount int         `pulumi:"failedLocationCount"`
	OdataType           string      `pulumi:"odataType"`
	WindowSize          *string     `pulumi:"windowSize"`
}

type LogAnalyticsDestination struct {
	Name                *string `pulumi:"name"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}





type LogAnalyticsDestinationInput interface {
	pulumi.Input

	ToLogAnalyticsDestinationOutput() LogAnalyticsDestinationOutput
	ToLogAnalyticsDestinationOutputWithContext(context.Context) LogAnalyticsDestinationOutput
}

type LogAnalyticsDestinationArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	WorkspaceResourceId pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (LogAnalyticsDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsDestination)(nil)).Elem()
}

func (i LogAnalyticsDestinationArgs) ToLogAnalyticsDestinationOutput() LogAnalyticsDestinationOutput {
	return i.ToLogAnalyticsDestinationOutputWithContext(context.Background())
}

func (i LogAnalyticsDestinationArgs) ToLogAnalyticsDestinationOutputWithContext(ctx context.Context) LogAnalyticsDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsDestinationOutput)
}





type LogAnalyticsDestinationArrayInput interface {
	pulumi.Input

	ToLogAnalyticsDestinationArrayOutput() LogAnalyticsDestinationArrayOutput
	ToLogAnalyticsDestinationArrayOutputWithContext(context.Context) LogAnalyticsDestinationArrayOutput
}

type LogAnalyticsDestinationArray []LogAnalyticsDestinationInput

func (LogAnalyticsDestinationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogAnalyticsDestination)(nil)).Elem()
}

func (i LogAnalyticsDestinationArray) ToLogAnalyticsDestinationArrayOutput() LogAnalyticsDestinationArrayOutput {
	return i.ToLogAnalyticsDestinationArrayOutputWithContext(context.Background())
}

func (i LogAnalyticsDestinationArray) ToLogAnalyticsDestinationArrayOutputWithContext(ctx context.Context) LogAnalyticsDestinationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsDestinationArrayOutput)
}

type LogAnalyticsDestinationOutput struct{ *pulumi.OutputState }

func (LogAnalyticsDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsDestination)(nil)).Elem()
}

func (o LogAnalyticsDestinationOutput) ToLogAnalyticsDestinationOutput() LogAnalyticsDestinationOutput {
	return o
}

func (o LogAnalyticsDestinationOutput) ToLogAnalyticsDestinationOutputWithContext(ctx context.Context) LogAnalyticsDestinationOutput {
	return o
}

func (o LogAnalyticsDestinationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsDestination) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsDestinationOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsDestination) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsDestinationArrayOutput struct{ *pulumi.OutputState }

func (LogAnalyticsDestinationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogAnalyticsDestination)(nil)).Elem()
}

func (o LogAnalyticsDestinationArrayOutput) ToLogAnalyticsDestinationArrayOutput() LogAnalyticsDestinationArrayOutput {
	return o
}

func (o LogAnalyticsDestinationArrayOutput) ToLogAnalyticsDestinationArrayOutputWithContext(ctx context.Context) LogAnalyticsDestinationArrayOutput {
	return o
}

func (o LogAnalyticsDestinationArrayOutput) Index(i pulumi.IntInput) LogAnalyticsDestinationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogAnalyticsDestination {
		return vs[0].([]LogAnalyticsDestination)[vs[1].(int)]
	}).(LogAnalyticsDestinationOutput)
}

type LogAnalyticsDestinationResponse struct {
	Name                *string `pulumi:"name"`
	WorkspaceId         string  `pulumi:"workspaceId"`
	WorkspaceResourceId *string `pulumi:"workspaceResourceId"`
}

type LogAnalyticsDestinationResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsDestinationResponse)(nil)).Elem()
}

func (o LogAnalyticsDestinationResponseOutput) ToLogAnalyticsDestinationResponseOutput() LogAnalyticsDestinationResponseOutput {
	return o
}

func (o LogAnalyticsDestinationResponseOutput) ToLogAnalyticsDestinationResponseOutputWithContext(ctx context.Context) LogAnalyticsDestinationResponseOutput {
	return o
}

func (o LogAnalyticsDestinationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsDestinationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsDestinationResponseOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalyticsDestinationResponse) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func (o LogAnalyticsDestinationResponseOutput) WorkspaceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsDestinationResponse) *string { return v.WorkspaceResourceId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsDestinationResponseArrayOutput struct{ *pulumi.OutputState }

func (LogAnalyticsDestinationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogAnalyticsDestinationResponse)(nil)).Elem()
}

func (o LogAnalyticsDestinationResponseArrayOutput) ToLogAnalyticsDestinationResponseArrayOutput() LogAnalyticsDestinationResponseArrayOutput {
	return o
}

func (o LogAnalyticsDestinationResponseArrayOutput) ToLogAnalyticsDestinationResponseArrayOutputWithContext(ctx context.Context) LogAnalyticsDestinationResponseArrayOutput {
	return o
}

func (o LogAnalyticsDestinationResponseArrayOutput) Index(i pulumi.IntInput) LogAnalyticsDestinationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogAnalyticsDestinationResponse {
		return vs[0].([]LogAnalyticsDestinationResponse)[vs[1].(int)]
	}).(LogAnalyticsDestinationResponseOutput)
}

type LogMetricTrigger struct {
	MetricColumn      *string  `pulumi:"metricColumn"`
	MetricTriggerType *string  `pulumi:"metricTriggerType"`
	Threshold         *float64 `pulumi:"threshold"`
	ThresholdOperator *string  `pulumi:"thresholdOperator"`
}

type LogMetricTriggerResponse struct {
	MetricColumn      *string  `pulumi:"metricColumn"`
	MetricTriggerType *string  `pulumi:"metricTriggerType"`
	Threshold         *float64 `pulumi:"threshold"`
	ThresholdOperator *string  `pulumi:"thresholdOperator"`
}

type LogSettings struct {
	Category        *string          `pulumi:"category"`
	Enabled         bool             `pulumi:"enabled"`
	RetentionPolicy *RetentionPolicy `pulumi:"retentionPolicy"`
}





type LogSettingsInput interface {
	pulumi.Input

	ToLogSettingsOutput() LogSettingsOutput
	ToLogSettingsOutputWithContext(context.Context) LogSettingsOutput
}

type LogSettingsArgs struct {
	Category        pulumi.StringPtrInput   `pulumi:"category"`
	Enabled         pulumi.BoolInput        `pulumi:"enabled"`
	RetentionPolicy RetentionPolicyPtrInput `pulumi:"retentionPolicy"`
}

func (LogSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettings)(nil)).Elem()
}

func (i LogSettingsArgs) ToLogSettingsOutput() LogSettingsOutput {
	return i.ToLogSettingsOutputWithContext(context.Background())
}

func (i LogSettingsArgs) ToLogSettingsOutputWithContext(ctx context.Context) LogSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingsOutput)
}





type LogSettingsArrayInput interface {
	pulumi.Input

	ToLogSettingsArrayOutput() LogSettingsArrayOutput
	ToLogSettingsArrayOutputWithContext(context.Context) LogSettingsArrayOutput
}

type LogSettingsArray []LogSettingsInput

func (LogSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettings)(nil)).Elem()
}

func (i LogSettingsArray) ToLogSettingsArrayOutput() LogSettingsArrayOutput {
	return i.ToLogSettingsArrayOutputWithContext(context.Background())
}

func (i LogSettingsArray) ToLogSettingsArrayOutputWithContext(ctx context.Context) LogSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingsArrayOutput)
}

type LogSettingsOutput struct{ *pulumi.OutputState }

func (LogSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettings)(nil)).Elem()
}

func (o LogSettingsOutput) ToLogSettingsOutput() LogSettingsOutput {
	return o
}

func (o LogSettingsOutput) ToLogSettingsOutputWithContext(ctx context.Context) LogSettingsOutput {
	return o
}

func (o LogSettingsOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogSettings) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o LogSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LogSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LogSettingsOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v LogSettings) *RetentionPolicy { return v.RetentionPolicy }).(RetentionPolicyPtrOutput)
}

type LogSettingsArrayOutput struct{ *pulumi.OutputState }

func (LogSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettings)(nil)).Elem()
}

func (o LogSettingsArrayOutput) ToLogSettingsArrayOutput() LogSettingsArrayOutput {
	return o
}

func (o LogSettingsArrayOutput) ToLogSettingsArrayOutputWithContext(ctx context.Context) LogSettingsArrayOutput {
	return o
}

func (o LogSettingsArrayOutput) Index(i pulumi.IntInput) LogSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogSettings {
		return vs[0].([]LogSettings)[vs[1].(int)]
	}).(LogSettingsOutput)
}

type LogSettingsResponse struct {
	Category        *string                  `pulumi:"category"`
	Enabled         bool                     `pulumi:"enabled"`
	RetentionPolicy *RetentionPolicyResponse `pulumi:"retentionPolicy"`
}

type LogSettingsResponseOutput struct{ *pulumi.OutputState }

func (LogSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettingsResponse)(nil)).Elem()
}

func (o LogSettingsResponseOutput) ToLogSettingsResponseOutput() LogSettingsResponseOutput {
	return o
}

func (o LogSettingsResponseOutput) ToLogSettingsResponseOutputWithContext(ctx context.Context) LogSettingsResponseOutput {
	return o
}

func (o LogSettingsResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogSettingsResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o LogSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LogSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LogSettingsResponseOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LogSettingsResponse) *RetentionPolicyResponse { return v.RetentionPolicy }).(RetentionPolicyResponsePtrOutput)
}

type LogSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (LogSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettingsResponse)(nil)).Elem()
}

func (o LogSettingsResponseArrayOutput) ToLogSettingsResponseArrayOutput() LogSettingsResponseArrayOutput {
	return o
}

func (o LogSettingsResponseArrayOutput) ToLogSettingsResponseArrayOutputWithContext(ctx context.Context) LogSettingsResponseArrayOutput {
	return o
}

func (o LogSettingsResponseArrayOutput) Index(i pulumi.IntInput) LogSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogSettingsResponse {
		return vs[0].([]LogSettingsResponse)[vs[1].(int)]
	}).(LogSettingsResponseOutput)
}

type LogToMetricAction struct {
	Criteria  []Criteria `pulumi:"criteria"`
	OdataType string     `pulumi:"odataType"`
}

type LogToMetricActionResponse struct {
	Criteria  []CriteriaResponse `pulumi:"criteria"`
	OdataType string             `pulumi:"odataType"`
}

type LogicAppReceiver struct {
	CallbackUrl          string `pulumi:"callbackUrl"`
	Name                 string `pulumi:"name"`
	ResourceId           string `pulumi:"resourceId"`
	UseCommonAlertSchema *bool  `pulumi:"useCommonAlertSchema"`
}


func (val *LogicAppReceiver) Defaults() *LogicAppReceiver {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}





type LogicAppReceiverInput interface {
	pulumi.Input

	ToLogicAppReceiverOutput() LogicAppReceiverOutput
	ToLogicAppReceiverOutputWithContext(context.Context) LogicAppReceiverOutput
}

type LogicAppReceiverArgs struct {
	CallbackUrl          pulumi.StringInput  `pulumi:"callbackUrl"`
	Name                 pulumi.StringInput  `pulumi:"name"`
	ResourceId           pulumi.StringInput  `pulumi:"resourceId"`
	UseCommonAlertSchema pulumi.BoolPtrInput `pulumi:"useCommonAlertSchema"`
}


func (val *LogicAppReceiverArgs) Defaults() *LogicAppReceiverArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		tmp.UseCommonAlertSchema = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (LogicAppReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogicAppReceiver)(nil)).Elem()
}

func (i LogicAppReceiverArgs) ToLogicAppReceiverOutput() LogicAppReceiverOutput {
	return i.ToLogicAppReceiverOutputWithContext(context.Background())
}

func (i LogicAppReceiverArgs) ToLogicAppReceiverOutputWithContext(ctx context.Context) LogicAppReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogicAppReceiverOutput)
}





type LogicAppReceiverArrayInput interface {
	pulumi.Input

	ToLogicAppReceiverArrayOutput() LogicAppReceiverArrayOutput
	ToLogicAppReceiverArrayOutputWithContext(context.Context) LogicAppReceiverArrayOutput
}

type LogicAppReceiverArray []LogicAppReceiverInput

func (LogicAppReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogicAppReceiver)(nil)).Elem()
}

func (i LogicAppReceiverArray) ToLogicAppReceiverArrayOutput() LogicAppReceiverArrayOutput {
	return i.ToLogicAppReceiverArrayOutputWithContext(context.Background())
}

func (i LogicAppReceiverArray) ToLogicAppReceiverArrayOutputWithContext(ctx context.Context) LogicAppReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogicAppReceiverArrayOutput)
}

type LogicAppReceiverOutput struct{ *pulumi.OutputState }

func (LogicAppReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogicAppReceiver)(nil)).Elem()
}

func (o LogicAppReceiverOutput) ToLogicAppReceiverOutput() LogicAppReceiverOutput {
	return o
}

func (o LogicAppReceiverOutput) ToLogicAppReceiverOutputWithContext(ctx context.Context) LogicAppReceiverOutput {
	return o
}

func (o LogicAppReceiverOutput) CallbackUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LogicAppReceiver) string { return v.CallbackUrl }).(pulumi.StringOutput)
}

func (o LogicAppReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LogicAppReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o LogicAppReceiverOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LogicAppReceiver) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o LogicAppReceiverOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogicAppReceiver) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type LogicAppReceiverArrayOutput struct{ *pulumi.OutputState }

func (LogicAppReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogicAppReceiver)(nil)).Elem()
}

func (o LogicAppReceiverArrayOutput) ToLogicAppReceiverArrayOutput() LogicAppReceiverArrayOutput {
	return o
}

func (o LogicAppReceiverArrayOutput) ToLogicAppReceiverArrayOutputWithContext(ctx context.Context) LogicAppReceiverArrayOutput {
	return o
}

func (o LogicAppReceiverArrayOutput) Index(i pulumi.IntInput) LogicAppReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogicAppReceiver {
		return vs[0].([]LogicAppReceiver)[vs[1].(int)]
	}).(LogicAppReceiverOutput)
}

type LogicAppReceiverResponse struct {
	CallbackUrl          string `pulumi:"callbackUrl"`
	Name                 string `pulumi:"name"`
	ResourceId           string `pulumi:"resourceId"`
	UseCommonAlertSchema *bool  `pulumi:"useCommonAlertSchema"`
}


func (val *LogicAppReceiverResponse) Defaults() *LogicAppReceiverResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}

type LogicAppReceiverResponseOutput struct{ *pulumi.OutputState }

func (LogicAppReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogicAppReceiverResponse)(nil)).Elem()
}

func (o LogicAppReceiverResponseOutput) ToLogicAppReceiverResponseOutput() LogicAppReceiverResponseOutput {
	return o
}

func (o LogicAppReceiverResponseOutput) ToLogicAppReceiverResponseOutputWithContext(ctx context.Context) LogicAppReceiverResponseOutput {
	return o
}

func (o LogicAppReceiverResponseOutput) CallbackUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LogicAppReceiverResponse) string { return v.CallbackUrl }).(pulumi.StringOutput)
}

func (o LogicAppReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LogicAppReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o LogicAppReceiverResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LogicAppReceiverResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o LogicAppReceiverResponseOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LogicAppReceiverResponse) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type LogicAppReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (LogicAppReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogicAppReceiverResponse)(nil)).Elem()
}

func (o LogicAppReceiverResponseArrayOutput) ToLogicAppReceiverResponseArrayOutput() LogicAppReceiverResponseArrayOutput {
	return o
}

func (o LogicAppReceiverResponseArrayOutput) ToLogicAppReceiverResponseArrayOutputWithContext(ctx context.Context) LogicAppReceiverResponseArrayOutput {
	return o
}

func (o LogicAppReceiverResponseArrayOutput) Index(i pulumi.IntInput) LogicAppReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogicAppReceiverResponse {
		return vs[0].([]LogicAppReceiverResponse)[vs[1].(int)]
	}).(LogicAppReceiverResponseOutput)
}

type ManagementEventAggregationCondition struct {
	Operator   *ConditionOperator `pulumi:"operator"`
	Threshold  *float64           `pulumi:"threshold"`
	WindowSize *string            `pulumi:"windowSize"`
}

type ManagementEventAggregationConditionResponse struct {
	Operator   *string  `pulumi:"operator"`
	Threshold  *float64 `pulumi:"threshold"`
	WindowSize *string  `pulumi:"windowSize"`
}

type ManagementEventRuleCondition struct {
	Aggregation *ManagementEventAggregationCondition `pulumi:"aggregation"`
	DataSource  interface{}                          `pulumi:"dataSource"`
	OdataType   string                               `pulumi:"odataType"`
}

type ManagementEventRuleConditionResponse struct {
	Aggregation *ManagementEventAggregationConditionResponse `pulumi:"aggregation"`
	DataSource  interface{}                                  `pulumi:"dataSource"`
	OdataType   string                                       `pulumi:"odataType"`
}

type ManagementGroupLogSettings struct {
	Category string `pulumi:"category"`
	Enabled  bool   `pulumi:"enabled"`
}





type ManagementGroupLogSettingsInput interface {
	pulumi.Input

	ToManagementGroupLogSettingsOutput() ManagementGroupLogSettingsOutput
	ToManagementGroupLogSettingsOutputWithContext(context.Context) ManagementGroupLogSettingsOutput
}

type ManagementGroupLogSettingsArgs struct {
	Category pulumi.StringInput `pulumi:"category"`
	Enabled  pulumi.BoolInput   `pulumi:"enabled"`
}

func (ManagementGroupLogSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupLogSettings)(nil)).Elem()
}

func (i ManagementGroupLogSettingsArgs) ToManagementGroupLogSettingsOutput() ManagementGroupLogSettingsOutput {
	return i.ToManagementGroupLogSettingsOutputWithContext(context.Background())
}

func (i ManagementGroupLogSettingsArgs) ToManagementGroupLogSettingsOutputWithContext(ctx context.Context) ManagementGroupLogSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupLogSettingsOutput)
}





type ManagementGroupLogSettingsArrayInput interface {
	pulumi.Input

	ToManagementGroupLogSettingsArrayOutput() ManagementGroupLogSettingsArrayOutput
	ToManagementGroupLogSettingsArrayOutputWithContext(context.Context) ManagementGroupLogSettingsArrayOutput
}

type ManagementGroupLogSettingsArray []ManagementGroupLogSettingsInput

func (ManagementGroupLogSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupLogSettings)(nil)).Elem()
}

func (i ManagementGroupLogSettingsArray) ToManagementGroupLogSettingsArrayOutput() ManagementGroupLogSettingsArrayOutput {
	return i.ToManagementGroupLogSettingsArrayOutputWithContext(context.Background())
}

func (i ManagementGroupLogSettingsArray) ToManagementGroupLogSettingsArrayOutputWithContext(ctx context.Context) ManagementGroupLogSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupLogSettingsArrayOutput)
}

type ManagementGroupLogSettingsOutput struct{ *pulumi.OutputState }

func (ManagementGroupLogSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupLogSettings)(nil)).Elem()
}

func (o ManagementGroupLogSettingsOutput) ToManagementGroupLogSettingsOutput() ManagementGroupLogSettingsOutput {
	return o
}

func (o ManagementGroupLogSettingsOutput) ToManagementGroupLogSettingsOutputWithContext(ctx context.Context) ManagementGroupLogSettingsOutput {
	return o
}

func (o ManagementGroupLogSettingsOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementGroupLogSettings) string { return v.Category }).(pulumi.StringOutput)
}

func (o ManagementGroupLogSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagementGroupLogSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ManagementGroupLogSettingsArrayOutput struct{ *pulumi.OutputState }

func (ManagementGroupLogSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupLogSettings)(nil)).Elem()
}

func (o ManagementGroupLogSettingsArrayOutput) ToManagementGroupLogSettingsArrayOutput() ManagementGroupLogSettingsArrayOutput {
	return o
}

func (o ManagementGroupLogSettingsArrayOutput) ToManagementGroupLogSettingsArrayOutputWithContext(ctx context.Context) ManagementGroupLogSettingsArrayOutput {
	return o
}

func (o ManagementGroupLogSettingsArrayOutput) Index(i pulumi.IntInput) ManagementGroupLogSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementGroupLogSettings {
		return vs[0].([]ManagementGroupLogSettings)[vs[1].(int)]
	}).(ManagementGroupLogSettingsOutput)
}

type ManagementGroupLogSettingsResponse struct {
	Category string `pulumi:"category"`
	Enabled  bool   `pulumi:"enabled"`
}

type ManagementGroupLogSettingsResponseOutput struct{ *pulumi.OutputState }

func (ManagementGroupLogSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupLogSettingsResponse)(nil)).Elem()
}

func (o ManagementGroupLogSettingsResponseOutput) ToManagementGroupLogSettingsResponseOutput() ManagementGroupLogSettingsResponseOutput {
	return o
}

func (o ManagementGroupLogSettingsResponseOutput) ToManagementGroupLogSettingsResponseOutputWithContext(ctx context.Context) ManagementGroupLogSettingsResponseOutput {
	return o
}

func (o ManagementGroupLogSettingsResponseOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementGroupLogSettingsResponse) string { return v.Category }).(pulumi.StringOutput)
}

func (o ManagementGroupLogSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagementGroupLogSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ManagementGroupLogSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementGroupLogSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupLogSettingsResponse)(nil)).Elem()
}

func (o ManagementGroupLogSettingsResponseArrayOutput) ToManagementGroupLogSettingsResponseArrayOutput() ManagementGroupLogSettingsResponseArrayOutput {
	return o
}

func (o ManagementGroupLogSettingsResponseArrayOutput) ToManagementGroupLogSettingsResponseArrayOutputWithContext(ctx context.Context) ManagementGroupLogSettingsResponseArrayOutput {
	return o
}

func (o ManagementGroupLogSettingsResponseArrayOutput) Index(i pulumi.IntInput) ManagementGroupLogSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementGroupLogSettingsResponse {
		return vs[0].([]ManagementGroupLogSettingsResponse)[vs[1].(int)]
	}).(ManagementGroupLogSettingsResponseOutput)
}

type MetricAlertAction struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}





type MetricAlertActionInput interface {
	pulumi.Input

	ToMetricAlertActionOutput() MetricAlertActionOutput
	ToMetricAlertActionOutputWithContext(context.Context) MetricAlertActionOutput
}

type MetricAlertActionArgs struct {
	ActionGroupId     pulumi.StringPtrInput `pulumi:"actionGroupId"`
	WebHookProperties pulumi.StringMapInput `pulumi:"webHookProperties"`
}

func (MetricAlertActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertAction)(nil)).Elem()
}

func (i MetricAlertActionArgs) ToMetricAlertActionOutput() MetricAlertActionOutput {
	return i.ToMetricAlertActionOutputWithContext(context.Background())
}

func (i MetricAlertActionArgs) ToMetricAlertActionOutputWithContext(ctx context.Context) MetricAlertActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertActionOutput)
}





type MetricAlertActionArrayInput interface {
	pulumi.Input

	ToMetricAlertActionArrayOutput() MetricAlertActionArrayOutput
	ToMetricAlertActionArrayOutputWithContext(context.Context) MetricAlertActionArrayOutput
}

type MetricAlertActionArray []MetricAlertActionInput

func (MetricAlertActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricAlertAction)(nil)).Elem()
}

func (i MetricAlertActionArray) ToMetricAlertActionArrayOutput() MetricAlertActionArrayOutput {
	return i.ToMetricAlertActionArrayOutputWithContext(context.Background())
}

func (i MetricAlertActionArray) ToMetricAlertActionArrayOutputWithContext(ctx context.Context) MetricAlertActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertActionArrayOutput)
}

type MetricAlertActionOutput struct{ *pulumi.OutputState }

func (MetricAlertActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertAction)(nil)).Elem()
}

func (o MetricAlertActionOutput) ToMetricAlertActionOutput() MetricAlertActionOutput {
	return o
}

func (o MetricAlertActionOutput) ToMetricAlertActionOutputWithContext(ctx context.Context) MetricAlertActionOutput {
	return o
}

func (o MetricAlertActionOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricAlertAction) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o MetricAlertActionOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v MetricAlertAction) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type MetricAlertActionArrayOutput struct{ *pulumi.OutputState }

func (MetricAlertActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricAlertAction)(nil)).Elem()
}

func (o MetricAlertActionArrayOutput) ToMetricAlertActionArrayOutput() MetricAlertActionArrayOutput {
	return o
}

func (o MetricAlertActionArrayOutput) ToMetricAlertActionArrayOutputWithContext(ctx context.Context) MetricAlertActionArrayOutput {
	return o
}

func (o MetricAlertActionArrayOutput) Index(i pulumi.IntInput) MetricAlertActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricAlertAction {
		return vs[0].([]MetricAlertAction)[vs[1].(int)]
	}).(MetricAlertActionOutput)
}

type MetricAlertActionResponse struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}

type MetricAlertActionResponseOutput struct{ *pulumi.OutputState }

func (MetricAlertActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertActionResponse)(nil)).Elem()
}

func (o MetricAlertActionResponseOutput) ToMetricAlertActionResponseOutput() MetricAlertActionResponseOutput {
	return o
}

func (o MetricAlertActionResponseOutput) ToMetricAlertActionResponseOutputWithContext(ctx context.Context) MetricAlertActionResponseOutput {
	return o
}

func (o MetricAlertActionResponseOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricAlertActionResponse) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o MetricAlertActionResponseOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v MetricAlertActionResponse) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type MetricAlertActionResponseArrayOutput struct{ *pulumi.OutputState }

func (MetricAlertActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricAlertActionResponse)(nil)).Elem()
}

func (o MetricAlertActionResponseArrayOutput) ToMetricAlertActionResponseArrayOutput() MetricAlertActionResponseArrayOutput {
	return o
}

func (o MetricAlertActionResponseArrayOutput) ToMetricAlertActionResponseArrayOutputWithContext(ctx context.Context) MetricAlertActionResponseArrayOutput {
	return o
}

func (o MetricAlertActionResponseArrayOutput) Index(i pulumi.IntInput) MetricAlertActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricAlertActionResponse {
		return vs[0].([]MetricAlertActionResponse)[vs[1].(int)]
	}).(MetricAlertActionResponseOutput)
}

type MetricAlertMultipleResourceMultipleMetricCriteria struct {
	AllOf     []interface{} `pulumi:"allOf"`
	OdataType string        `pulumi:"odataType"`
}

type MetricAlertMultipleResourceMultipleMetricCriteriaResponse struct {
	AllOf     []interface{} `pulumi:"allOf"`
	OdataType string        `pulumi:"odataType"`
}

type MetricAlertSingleResourceMultipleMetricCriteria struct {
	AllOf     []MetricCriteria `pulumi:"allOf"`
	OdataType string           `pulumi:"odataType"`
}

type MetricAlertSingleResourceMultipleMetricCriteriaResponse struct {
	AllOf     []MetricCriteriaResponse `pulumi:"allOf"`
	OdataType string                   `pulumi:"odataType"`
}

type MetricCriteria struct {
	CriterionType        string            `pulumi:"criterionType"`
	Dimensions           []MetricDimension `pulumi:"dimensions"`
	MetricName           string            `pulumi:"metricName"`
	MetricNamespace      *string           `pulumi:"metricNamespace"`
	Name                 string            `pulumi:"name"`
	Operator             string            `pulumi:"operator"`
	SkipMetricValidation *bool             `pulumi:"skipMetricValidation"`
	Threshold            float64           `pulumi:"threshold"`
	TimeAggregation      string            `pulumi:"timeAggregation"`
}

type MetricCriteriaResponse struct {
	CriterionType        string                    `pulumi:"criterionType"`
	Dimensions           []MetricDimensionResponse `pulumi:"dimensions"`
	MetricName           string                    `pulumi:"metricName"`
	MetricNamespace      *string                   `pulumi:"metricNamespace"`
	Name                 string                    `pulumi:"name"`
	Operator             string                    `pulumi:"operator"`
	SkipMetricValidation *bool                     `pulumi:"skipMetricValidation"`
	Threshold            float64                   `pulumi:"threshold"`
	TimeAggregation      string                    `pulumi:"timeAggregation"`
}

type MetricDimension struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type MetricDimensionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type MetricSettings struct {
	Category        *string          `pulumi:"category"`
	Enabled         bool             `pulumi:"enabled"`
	RetentionPolicy *RetentionPolicy `pulumi:"retentionPolicy"`
	TimeGrain       *string          `pulumi:"timeGrain"`
}





type MetricSettingsInput interface {
	pulumi.Input

	ToMetricSettingsOutput() MetricSettingsOutput
	ToMetricSettingsOutputWithContext(context.Context) MetricSettingsOutput
}

type MetricSettingsArgs struct {
	Category        pulumi.StringPtrInput   `pulumi:"category"`
	Enabled         pulumi.BoolInput        `pulumi:"enabled"`
	RetentionPolicy RetentionPolicyPtrInput `pulumi:"retentionPolicy"`
	TimeGrain       pulumi.StringPtrInput   `pulumi:"timeGrain"`
}

func (MetricSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricSettings)(nil)).Elem()
}

func (i MetricSettingsArgs) ToMetricSettingsOutput() MetricSettingsOutput {
	return i.ToMetricSettingsOutputWithContext(context.Background())
}

func (i MetricSettingsArgs) ToMetricSettingsOutputWithContext(ctx context.Context) MetricSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricSettingsOutput)
}





type MetricSettingsArrayInput interface {
	pulumi.Input

	ToMetricSettingsArrayOutput() MetricSettingsArrayOutput
	ToMetricSettingsArrayOutputWithContext(context.Context) MetricSettingsArrayOutput
}

type MetricSettingsArray []MetricSettingsInput

func (MetricSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricSettings)(nil)).Elem()
}

func (i MetricSettingsArray) ToMetricSettingsArrayOutput() MetricSettingsArrayOutput {
	return i.ToMetricSettingsArrayOutputWithContext(context.Background())
}

func (i MetricSettingsArray) ToMetricSettingsArrayOutputWithContext(ctx context.Context) MetricSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricSettingsArrayOutput)
}

type MetricSettingsOutput struct{ *pulumi.OutputState }

func (MetricSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricSettings)(nil)).Elem()
}

func (o MetricSettingsOutput) ToMetricSettingsOutput() MetricSettingsOutput {
	return o
}

func (o MetricSettingsOutput) ToMetricSettingsOutputWithContext(ctx context.Context) MetricSettingsOutput {
	return o
}

func (o MetricSettingsOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricSettings) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o MetricSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v MetricSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o MetricSettingsOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v MetricSettings) *RetentionPolicy { return v.RetentionPolicy }).(RetentionPolicyPtrOutput)
}

func (o MetricSettingsOutput) TimeGrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricSettings) *string { return v.TimeGrain }).(pulumi.StringPtrOutput)
}

type MetricSettingsArrayOutput struct{ *pulumi.OutputState }

func (MetricSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricSettings)(nil)).Elem()
}

func (o MetricSettingsArrayOutput) ToMetricSettingsArrayOutput() MetricSettingsArrayOutput {
	return o
}

func (o MetricSettingsArrayOutput) ToMetricSettingsArrayOutputWithContext(ctx context.Context) MetricSettingsArrayOutput {
	return o
}

func (o MetricSettingsArrayOutput) Index(i pulumi.IntInput) MetricSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricSettings {
		return vs[0].([]MetricSettings)[vs[1].(int)]
	}).(MetricSettingsOutput)
}

type MetricSettingsResponse struct {
	Category        *string                  `pulumi:"category"`
	Enabled         bool                     `pulumi:"enabled"`
	RetentionPolicy *RetentionPolicyResponse `pulumi:"retentionPolicy"`
	TimeGrain       *string                  `pulumi:"timeGrain"`
}

type MetricSettingsResponseOutput struct{ *pulumi.OutputState }

func (MetricSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricSettingsResponse)(nil)).Elem()
}

func (o MetricSettingsResponseOutput) ToMetricSettingsResponseOutput() MetricSettingsResponseOutput {
	return o
}

func (o MetricSettingsResponseOutput) ToMetricSettingsResponseOutputWithContext(ctx context.Context) MetricSettingsResponseOutput {
	return o
}

func (o MetricSettingsResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricSettingsResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o MetricSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v MetricSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o MetricSettingsResponseOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v MetricSettingsResponse) *RetentionPolicyResponse { return v.RetentionPolicy }).(RetentionPolicyResponsePtrOutput)
}

func (o MetricSettingsResponseOutput) TimeGrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricSettingsResponse) *string { return v.TimeGrain }).(pulumi.StringPtrOutput)
}

type MetricSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (MetricSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricSettingsResponse)(nil)).Elem()
}

func (o MetricSettingsResponseArrayOutput) ToMetricSettingsResponseArrayOutput() MetricSettingsResponseArrayOutput {
	return o
}

func (o MetricSettingsResponseArrayOutput) ToMetricSettingsResponseArrayOutputWithContext(ctx context.Context) MetricSettingsResponseArrayOutput {
	return o
}

func (o MetricSettingsResponseArrayOutput) Index(i pulumi.IntInput) MetricSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricSettingsResponse {
		return vs[0].([]MetricSettingsResponse)[vs[1].(int)]
	}).(MetricSettingsResponseOutput)
}

type MetricTrigger struct {
	Dimensions             []ScaleRuleMetricDimension `pulumi:"dimensions"`
	DividePerInstance      *bool                      `pulumi:"dividePerInstance"`
	MetricName             string                     `pulumi:"metricName"`
	MetricNamespace        *string                    `pulumi:"metricNamespace"`
	MetricResourceLocation *string                    `pulumi:"metricResourceLocation"`
	MetricResourceUri      string                     `pulumi:"metricResourceUri"`
	Operator               ComparisonOperationType    `pulumi:"operator"`
	Statistic              MetricStatisticType        `pulumi:"statistic"`
	Threshold              float64                    `pulumi:"threshold"`
	TimeAggregation        TimeAggregationType        `pulumi:"timeAggregation"`
	TimeGrain              string                     `pulumi:"timeGrain"`
	TimeWindow             string                     `pulumi:"timeWindow"`
}





type MetricTriggerInput interface {
	pulumi.Input

	ToMetricTriggerOutput() MetricTriggerOutput
	ToMetricTriggerOutputWithContext(context.Context) MetricTriggerOutput
}

type MetricTriggerArgs struct {
	Dimensions             ScaleRuleMetricDimensionArrayInput `pulumi:"dimensions"`
	DividePerInstance      pulumi.BoolPtrInput                `pulumi:"dividePerInstance"`
	MetricName             pulumi.StringInput                 `pulumi:"metricName"`
	MetricNamespace        pulumi.StringPtrInput              `pulumi:"metricNamespace"`
	MetricResourceLocation pulumi.StringPtrInput              `pulumi:"metricResourceLocation"`
	MetricResourceUri      pulumi.StringInput                 `pulumi:"metricResourceUri"`
	Operator               ComparisonOperationTypeInput       `pulumi:"operator"`
	Statistic              MetricStatisticTypeInput           `pulumi:"statistic"`
	Threshold              pulumi.Float64Input                `pulumi:"threshold"`
	TimeAggregation        TimeAggregationTypeInput           `pulumi:"timeAggregation"`
	TimeGrain              pulumi.StringInput                 `pulumi:"timeGrain"`
	TimeWindow             pulumi.StringInput                 `pulumi:"timeWindow"`
}

func (MetricTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricTrigger)(nil)).Elem()
}

func (i MetricTriggerArgs) ToMetricTriggerOutput() MetricTriggerOutput {
	return i.ToMetricTriggerOutputWithContext(context.Background())
}

func (i MetricTriggerArgs) ToMetricTriggerOutputWithContext(ctx context.Context) MetricTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricTriggerOutput)
}

type MetricTriggerOutput struct{ *pulumi.OutputState }

func (MetricTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricTrigger)(nil)).Elem()
}

func (o MetricTriggerOutput) ToMetricTriggerOutput() MetricTriggerOutput {
	return o
}

func (o MetricTriggerOutput) ToMetricTriggerOutputWithContext(ctx context.Context) MetricTriggerOutput {
	return o
}

func (o MetricTriggerOutput) Dimensions() ScaleRuleMetricDimensionArrayOutput {
	return o.ApplyT(func(v MetricTrigger) []ScaleRuleMetricDimension { return v.Dimensions }).(ScaleRuleMetricDimensionArrayOutput)
}

func (o MetricTriggerOutput) DividePerInstance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MetricTrigger) *bool { return v.DividePerInstance }).(pulumi.BoolPtrOutput)
}

func (o MetricTriggerOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTrigger) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o MetricTriggerOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricTrigger) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o MetricTriggerOutput) MetricResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricTrigger) *string { return v.MetricResourceLocation }).(pulumi.StringPtrOutput)
}

func (o MetricTriggerOutput) MetricResourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTrigger) string { return v.MetricResourceUri }).(pulumi.StringOutput)
}

func (o MetricTriggerOutput) Operator() ComparisonOperationTypeOutput {
	return o.ApplyT(func(v MetricTrigger) ComparisonOperationType { return v.Operator }).(ComparisonOperationTypeOutput)
}

func (o MetricTriggerOutput) Statistic() MetricStatisticTypeOutput {
	return o.ApplyT(func(v MetricTrigger) MetricStatisticType { return v.Statistic }).(MetricStatisticTypeOutput)
}

func (o MetricTriggerOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v MetricTrigger) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o MetricTriggerOutput) TimeAggregation() TimeAggregationTypeOutput {
	return o.ApplyT(func(v MetricTrigger) TimeAggregationType { return v.TimeAggregation }).(TimeAggregationTypeOutput)
}

func (o MetricTriggerOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTrigger) string { return v.TimeGrain }).(pulumi.StringOutput)
}

func (o MetricTriggerOutput) TimeWindow() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTrigger) string { return v.TimeWindow }).(pulumi.StringOutput)
}

type MetricTriggerResponse struct {
	Dimensions             []ScaleRuleMetricDimensionResponse `pulumi:"dimensions"`
	DividePerInstance      *bool                              `pulumi:"dividePerInstance"`
	MetricName             string                             `pulumi:"metricName"`
	MetricNamespace        *string                            `pulumi:"metricNamespace"`
	MetricResourceLocation *string                            `pulumi:"metricResourceLocation"`
	MetricResourceUri      string                             `pulumi:"metricResourceUri"`
	Operator               string                             `pulumi:"operator"`
	Statistic              string                             `pulumi:"statistic"`
	Threshold              float64                            `pulumi:"threshold"`
	TimeAggregation        string                             `pulumi:"timeAggregation"`
	TimeGrain              string                             `pulumi:"timeGrain"`
	TimeWindow             string                             `pulumi:"timeWindow"`
}

type MetricTriggerResponseOutput struct{ *pulumi.OutputState }

func (MetricTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricTriggerResponse)(nil)).Elem()
}

func (o MetricTriggerResponseOutput) ToMetricTriggerResponseOutput() MetricTriggerResponseOutput {
	return o
}

func (o MetricTriggerResponseOutput) ToMetricTriggerResponseOutputWithContext(ctx context.Context) MetricTriggerResponseOutput {
	return o
}

func (o MetricTriggerResponseOutput) Dimensions() ScaleRuleMetricDimensionResponseArrayOutput {
	return o.ApplyT(func(v MetricTriggerResponse) []ScaleRuleMetricDimensionResponse { return v.Dimensions }).(ScaleRuleMetricDimensionResponseArrayOutput)
}

func (o MetricTriggerResponseOutput) DividePerInstance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MetricTriggerResponse) *bool { return v.DividePerInstance }).(pulumi.BoolPtrOutput)
}

func (o MetricTriggerResponseOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTriggerResponse) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o MetricTriggerResponseOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricTriggerResponse) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o MetricTriggerResponseOutput) MetricResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricTriggerResponse) *string { return v.MetricResourceLocation }).(pulumi.StringPtrOutput)
}

func (o MetricTriggerResponseOutput) MetricResourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTriggerResponse) string { return v.MetricResourceUri }).(pulumi.StringOutput)
}

func (o MetricTriggerResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTriggerResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o MetricTriggerResponseOutput) Statistic() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTriggerResponse) string { return v.Statistic }).(pulumi.StringOutput)
}

func (o MetricTriggerResponseOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v MetricTriggerResponse) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o MetricTriggerResponseOutput) TimeAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTriggerResponse) string { return v.TimeAggregation }).(pulumi.StringOutput)
}

func (o MetricTriggerResponseOutput) TimeGrain() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTriggerResponse) string { return v.TimeGrain }).(pulumi.StringOutput)
}

func (o MetricTriggerResponseOutput) TimeWindow() pulumi.StringOutput {
	return o.ApplyT(func(v MetricTriggerResponse) string { return v.TimeWindow }).(pulumi.StringOutput)
}

type MyWorkbookManagedIdentity struct {
	Type *string `pulumi:"type"`
}





type MyWorkbookManagedIdentityInput interface {
	pulumi.Input

	ToMyWorkbookManagedIdentityOutput() MyWorkbookManagedIdentityOutput
	ToMyWorkbookManagedIdentityOutputWithContext(context.Context) MyWorkbookManagedIdentityOutput
}

type MyWorkbookManagedIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (MyWorkbookManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookManagedIdentity)(nil)).Elem()
}

func (i MyWorkbookManagedIdentityArgs) ToMyWorkbookManagedIdentityOutput() MyWorkbookManagedIdentityOutput {
	return i.ToMyWorkbookManagedIdentityOutputWithContext(context.Background())
}

func (i MyWorkbookManagedIdentityArgs) ToMyWorkbookManagedIdentityOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityOutput)
}

func (i MyWorkbookManagedIdentityArgs) ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput {
	return i.ToMyWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (i MyWorkbookManagedIdentityArgs) ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityOutput).ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx)
}









type MyWorkbookManagedIdentityPtrInput interface {
	pulumi.Input

	ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput
	ToMyWorkbookManagedIdentityPtrOutputWithContext(context.Context) MyWorkbookManagedIdentityPtrOutput
}

type myWorkbookManagedIdentityPtrType MyWorkbookManagedIdentityArgs

func MyWorkbookManagedIdentityPtr(v *MyWorkbookManagedIdentityArgs) MyWorkbookManagedIdentityPtrInput {
	return (*myWorkbookManagedIdentityPtrType)(v)
}

func (*myWorkbookManagedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookManagedIdentity)(nil)).Elem()
}

func (i *myWorkbookManagedIdentityPtrType) ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput {
	return i.ToMyWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *myWorkbookManagedIdentityPtrType) ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityPtrOutput)
}

type MyWorkbookManagedIdentityOutput struct{ *pulumi.OutputState }

func (MyWorkbookManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookManagedIdentity)(nil)).Elem()
}

func (o MyWorkbookManagedIdentityOutput) ToMyWorkbookManagedIdentityOutput() MyWorkbookManagedIdentityOutput {
	return o
}

func (o MyWorkbookManagedIdentityOutput) ToMyWorkbookManagedIdentityOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityOutput {
	return o
}

func (o MyWorkbookManagedIdentityOutput) ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput {
	return o.ToMyWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (o MyWorkbookManagedIdentityOutput) ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MyWorkbookManagedIdentity) *MyWorkbookManagedIdentity {
		return &v
	}).(MyWorkbookManagedIdentityPtrOutput)
}

func (o MyWorkbookManagedIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MyWorkbookManagedIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type MyWorkbookManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (MyWorkbookManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookManagedIdentity)(nil)).Elem()
}

func (o MyWorkbookManagedIdentityPtrOutput) ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput {
	return o
}

func (o MyWorkbookManagedIdentityPtrOutput) ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityPtrOutput {
	return o
}

func (o MyWorkbookManagedIdentityPtrOutput) Elem() MyWorkbookManagedIdentityOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentity) MyWorkbookManagedIdentity {
		if v != nil {
			return *v
		}
		var ret MyWorkbookManagedIdentity
		return ret
	}).(MyWorkbookManagedIdentityOutput)
}

func (o MyWorkbookManagedIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type MyWorkbookManagedIdentityResponse struct {
	Type                   *string                                   `pulumi:"type"`
	UserAssignedIdentities *MyWorkbookUserAssignedIdentitiesResponse `pulumi:"userAssignedIdentities"`
}

type MyWorkbookManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (MyWorkbookManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookManagedIdentityResponse)(nil)).Elem()
}

func (o MyWorkbookManagedIdentityResponseOutput) ToMyWorkbookManagedIdentityResponseOutput() MyWorkbookManagedIdentityResponseOutput {
	return o
}

func (o MyWorkbookManagedIdentityResponseOutput) ToMyWorkbookManagedIdentityResponseOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponseOutput {
	return o
}

func (o MyWorkbookManagedIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MyWorkbookManagedIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o MyWorkbookManagedIdentityResponseOutput) UserAssignedIdentities() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyT(func(v MyWorkbookManagedIdentityResponse) *MyWorkbookUserAssignedIdentitiesResponse {
		return v.UserAssignedIdentities
	}).(MyWorkbookUserAssignedIdentitiesResponsePtrOutput)
}

type MyWorkbookManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (MyWorkbookManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookManagedIdentityResponse)(nil)).Elem()
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput {
	return o
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponsePtrOutput {
	return o
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) Elem() MyWorkbookManagedIdentityResponseOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentityResponse) MyWorkbookManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret MyWorkbookManagedIdentityResponse
		return ret
	}).(MyWorkbookManagedIdentityResponseOutput)
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) UserAssignedIdentities() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentityResponse) *MyWorkbookUserAssignedIdentitiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(MyWorkbookUserAssignedIdentitiesResponsePtrOutput)
}

type MyWorkbookUserAssignedIdentitiesResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
}

type MyWorkbookUserAssignedIdentitiesResponseOutput struct{ *pulumi.OutputState }

func (MyWorkbookUserAssignedIdentitiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) ToMyWorkbookUserAssignedIdentitiesResponseOutput() MyWorkbookUserAssignedIdentitiesResponseOutput {
	return o
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) ToMyWorkbookUserAssignedIdentitiesResponseOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponseOutput {
	return o
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v MyWorkbookUserAssignedIdentitiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v MyWorkbookUserAssignedIdentitiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type MyWorkbookUserAssignedIdentitiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MyWorkbookUserAssignedIdentitiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) Elem() MyWorkbookUserAssignedIdentitiesResponseOutput {
	return o.ApplyT(func(v *MyWorkbookUserAssignedIdentitiesResponse) MyWorkbookUserAssignedIdentitiesResponse {
		if v != nil {
			return *v
		}
		var ret MyWorkbookUserAssignedIdentitiesResponse
		return ret
	}).(MyWorkbookUserAssignedIdentitiesResponseOutput)
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

type PerfCounterDataSource struct {
	CounterSpecifiers          []string `pulumi:"counterSpecifiers"`
	Name                       *string  `pulumi:"name"`
	SamplingFrequencyInSeconds *int     `pulumi:"samplingFrequencyInSeconds"`
	Streams                    []string `pulumi:"streams"`
}





type PerfCounterDataSourceInput interface {
	pulumi.Input

	ToPerfCounterDataSourceOutput() PerfCounterDataSourceOutput
	ToPerfCounterDataSourceOutputWithContext(context.Context) PerfCounterDataSourceOutput
}

type PerfCounterDataSourceArgs struct {
	CounterSpecifiers          pulumi.StringArrayInput `pulumi:"counterSpecifiers"`
	Name                       pulumi.StringPtrInput   `pulumi:"name"`
	SamplingFrequencyInSeconds pulumi.IntPtrInput      `pulumi:"samplingFrequencyInSeconds"`
	Streams                    pulumi.StringArrayInput `pulumi:"streams"`
}

func (PerfCounterDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerfCounterDataSource)(nil)).Elem()
}

func (i PerfCounterDataSourceArgs) ToPerfCounterDataSourceOutput() PerfCounterDataSourceOutput {
	return i.ToPerfCounterDataSourceOutputWithContext(context.Background())
}

func (i PerfCounterDataSourceArgs) ToPerfCounterDataSourceOutputWithContext(ctx context.Context) PerfCounterDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerfCounterDataSourceOutput)
}





type PerfCounterDataSourceArrayInput interface {
	pulumi.Input

	ToPerfCounterDataSourceArrayOutput() PerfCounterDataSourceArrayOutput
	ToPerfCounterDataSourceArrayOutputWithContext(context.Context) PerfCounterDataSourceArrayOutput
}

type PerfCounterDataSourceArray []PerfCounterDataSourceInput

func (PerfCounterDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerfCounterDataSource)(nil)).Elem()
}

func (i PerfCounterDataSourceArray) ToPerfCounterDataSourceArrayOutput() PerfCounterDataSourceArrayOutput {
	return i.ToPerfCounterDataSourceArrayOutputWithContext(context.Background())
}

func (i PerfCounterDataSourceArray) ToPerfCounterDataSourceArrayOutputWithContext(ctx context.Context) PerfCounterDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerfCounterDataSourceArrayOutput)
}

type PerfCounterDataSourceOutput struct{ *pulumi.OutputState }

func (PerfCounterDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerfCounterDataSource)(nil)).Elem()
}

func (o PerfCounterDataSourceOutput) ToPerfCounterDataSourceOutput() PerfCounterDataSourceOutput {
	return o
}

func (o PerfCounterDataSourceOutput) ToPerfCounterDataSourceOutputWithContext(ctx context.Context) PerfCounterDataSourceOutput {
	return o
}

func (o PerfCounterDataSourceOutput) CounterSpecifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PerfCounterDataSource) []string { return v.CounterSpecifiers }).(pulumi.StringArrayOutput)
}

func (o PerfCounterDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PerfCounterDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PerfCounterDataSourceOutput) SamplingFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PerfCounterDataSource) *int { return v.SamplingFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o PerfCounterDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PerfCounterDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type PerfCounterDataSourceArrayOutput struct{ *pulumi.OutputState }

func (PerfCounterDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerfCounterDataSource)(nil)).Elem()
}

func (o PerfCounterDataSourceArrayOutput) ToPerfCounterDataSourceArrayOutput() PerfCounterDataSourceArrayOutput {
	return o
}

func (o PerfCounterDataSourceArrayOutput) ToPerfCounterDataSourceArrayOutputWithContext(ctx context.Context) PerfCounterDataSourceArrayOutput {
	return o
}

func (o PerfCounterDataSourceArrayOutput) Index(i pulumi.IntInput) PerfCounterDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerfCounterDataSource {
		return vs[0].([]PerfCounterDataSource)[vs[1].(int)]
	}).(PerfCounterDataSourceOutput)
}

type PerfCounterDataSourceResponse struct {
	CounterSpecifiers          []string `pulumi:"counterSpecifiers"`
	Name                       *string  `pulumi:"name"`
	SamplingFrequencyInSeconds *int     `pulumi:"samplingFrequencyInSeconds"`
	Streams                    []string `pulumi:"streams"`
}

type PerfCounterDataSourceResponseOutput struct{ *pulumi.OutputState }

func (PerfCounterDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerfCounterDataSourceResponse)(nil)).Elem()
}

func (o PerfCounterDataSourceResponseOutput) ToPerfCounterDataSourceResponseOutput() PerfCounterDataSourceResponseOutput {
	return o
}

func (o PerfCounterDataSourceResponseOutput) ToPerfCounterDataSourceResponseOutputWithContext(ctx context.Context) PerfCounterDataSourceResponseOutput {
	return o
}

func (o PerfCounterDataSourceResponseOutput) CounterSpecifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PerfCounterDataSourceResponse) []string { return v.CounterSpecifiers }).(pulumi.StringArrayOutput)
}

func (o PerfCounterDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PerfCounterDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PerfCounterDataSourceResponseOutput) SamplingFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PerfCounterDataSourceResponse) *int { return v.SamplingFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o PerfCounterDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PerfCounterDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type PerfCounterDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (PerfCounterDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerfCounterDataSourceResponse)(nil)).Elem()
}

func (o PerfCounterDataSourceResponseArrayOutput) ToPerfCounterDataSourceResponseArrayOutput() PerfCounterDataSourceResponseArrayOutput {
	return o
}

func (o PerfCounterDataSourceResponseArrayOutput) ToPerfCounterDataSourceResponseArrayOutputWithContext(ctx context.Context) PerfCounterDataSourceResponseArrayOutput {
	return o
}

func (o PerfCounterDataSourceResponseArrayOutput) Index(i pulumi.IntInput) PerfCounterDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerfCounterDataSourceResponse {
		return vs[0].([]PerfCounterDataSourceResponse)[vs[1].(int)]
	}).(PerfCounterDataSourceResponseOutput)
}

type PerformanceCounterConfiguration struct {
	Instance       *string `pulumi:"instance"`
	Name           string  `pulumi:"name"`
	SamplingPeriod string  `pulumi:"samplingPeriod"`
}





type PerformanceCounterConfigurationInput interface {
	pulumi.Input

	ToPerformanceCounterConfigurationOutput() PerformanceCounterConfigurationOutput
	ToPerformanceCounterConfigurationOutputWithContext(context.Context) PerformanceCounterConfigurationOutput
}

type PerformanceCounterConfigurationArgs struct {
	Instance       pulumi.StringPtrInput `pulumi:"instance"`
	Name           pulumi.StringInput    `pulumi:"name"`
	SamplingPeriod pulumi.StringInput    `pulumi:"samplingPeriod"`
}

func (PerformanceCounterConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerformanceCounterConfiguration)(nil)).Elem()
}

func (i PerformanceCounterConfigurationArgs) ToPerformanceCounterConfigurationOutput() PerformanceCounterConfigurationOutput {
	return i.ToPerformanceCounterConfigurationOutputWithContext(context.Background())
}

func (i PerformanceCounterConfigurationArgs) ToPerformanceCounterConfigurationOutputWithContext(ctx context.Context) PerformanceCounterConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerformanceCounterConfigurationOutput)
}





type PerformanceCounterConfigurationArrayInput interface {
	pulumi.Input

	ToPerformanceCounterConfigurationArrayOutput() PerformanceCounterConfigurationArrayOutput
	ToPerformanceCounterConfigurationArrayOutputWithContext(context.Context) PerformanceCounterConfigurationArrayOutput
}

type PerformanceCounterConfigurationArray []PerformanceCounterConfigurationInput

func (PerformanceCounterConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerformanceCounterConfiguration)(nil)).Elem()
}

func (i PerformanceCounterConfigurationArray) ToPerformanceCounterConfigurationArrayOutput() PerformanceCounterConfigurationArrayOutput {
	return i.ToPerformanceCounterConfigurationArrayOutputWithContext(context.Background())
}

func (i PerformanceCounterConfigurationArray) ToPerformanceCounterConfigurationArrayOutputWithContext(ctx context.Context) PerformanceCounterConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerformanceCounterConfigurationArrayOutput)
}

type PerformanceCounterConfigurationOutput struct{ *pulumi.OutputState }

func (PerformanceCounterConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerformanceCounterConfiguration)(nil)).Elem()
}

func (o PerformanceCounterConfigurationOutput) ToPerformanceCounterConfigurationOutput() PerformanceCounterConfigurationOutput {
	return o
}

func (o PerformanceCounterConfigurationOutput) ToPerformanceCounterConfigurationOutputWithContext(ctx context.Context) PerformanceCounterConfigurationOutput {
	return o
}

func (o PerformanceCounterConfigurationOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PerformanceCounterConfiguration) *string { return v.Instance }).(pulumi.StringPtrOutput)
}

func (o PerformanceCounterConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PerformanceCounterConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o PerformanceCounterConfigurationOutput) SamplingPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v PerformanceCounterConfiguration) string { return v.SamplingPeriod }).(pulumi.StringOutput)
}

type PerformanceCounterConfigurationArrayOutput struct{ *pulumi.OutputState }

func (PerformanceCounterConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerformanceCounterConfiguration)(nil)).Elem()
}

func (o PerformanceCounterConfigurationArrayOutput) ToPerformanceCounterConfigurationArrayOutput() PerformanceCounterConfigurationArrayOutput {
	return o
}

func (o PerformanceCounterConfigurationArrayOutput) ToPerformanceCounterConfigurationArrayOutputWithContext(ctx context.Context) PerformanceCounterConfigurationArrayOutput {
	return o
}

func (o PerformanceCounterConfigurationArrayOutput) Index(i pulumi.IntInput) PerformanceCounterConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerformanceCounterConfiguration {
		return vs[0].([]PerformanceCounterConfiguration)[vs[1].(int)]
	}).(PerformanceCounterConfigurationOutput)
}

type PerformanceCounterConfigurationResponse struct {
	Instance       *string `pulumi:"instance"`
	Name           string  `pulumi:"name"`
	SamplingPeriod string  `pulumi:"samplingPeriod"`
}

type PerformanceCounterConfigurationResponseOutput struct{ *pulumi.OutputState }

func (PerformanceCounterConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PerformanceCounterConfigurationResponse)(nil)).Elem()
}

func (o PerformanceCounterConfigurationResponseOutput) ToPerformanceCounterConfigurationResponseOutput() PerformanceCounterConfigurationResponseOutput {
	return o
}

func (o PerformanceCounterConfigurationResponseOutput) ToPerformanceCounterConfigurationResponseOutputWithContext(ctx context.Context) PerformanceCounterConfigurationResponseOutput {
	return o
}

func (o PerformanceCounterConfigurationResponseOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PerformanceCounterConfigurationResponse) *string { return v.Instance }).(pulumi.StringPtrOutput)
}

func (o PerformanceCounterConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PerformanceCounterConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PerformanceCounterConfigurationResponseOutput) SamplingPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v PerformanceCounterConfigurationResponse) string { return v.SamplingPeriod }).(pulumi.StringOutput)
}

type PerformanceCounterConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (PerformanceCounterConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerformanceCounterConfigurationResponse)(nil)).Elem()
}

func (o PerformanceCounterConfigurationResponseArrayOutput) ToPerformanceCounterConfigurationResponseArrayOutput() PerformanceCounterConfigurationResponseArrayOutput {
	return o
}

func (o PerformanceCounterConfigurationResponseArrayOutput) ToPerformanceCounterConfigurationResponseArrayOutputWithContext(ctx context.Context) PerformanceCounterConfigurationResponseArrayOutput {
	return o
}

func (o PerformanceCounterConfigurationResponseArrayOutput) Index(i pulumi.IntInput) PerformanceCounterConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PerformanceCounterConfigurationResponse {
		return vs[0].([]PerformanceCounterConfigurationResponse)[vs[1].(int)]
	}).(PerformanceCounterConfigurationResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                             `pulumi:"id"`
	Name                              string                                             `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                             `pulumi:"provisioningState"`
	Type                              string                                             `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointPropertyResponse { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointProperty struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput
	ToPrivateEndpointPropertyOutputWithContext(context.Context) PrivateEndpointPropertyOutput
}

type PrivateEndpointPropertyArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return i.ToPrivateEndpointPropertyOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput)
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput).ToPrivateEndpointPropertyPtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput
	ToPrivateEndpointPropertyPtrOutputWithContext(context.Context) PrivateEndpointPropertyPtrOutput
}

type privateEndpointPropertyPtrType PrivateEndpointPropertyArgs

func PrivateEndpointPropertyPtr(v *PrivateEndpointPropertyArgs) PrivateEndpointPropertyPtrInput {
	return (*privateEndpointPropertyPtrType)(v)
}

func (*privateEndpointPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyPtrOutput)
}

type PrivateEndpointPropertyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointProperty) *PrivateEndpointProperty {
		return &v
	}).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointPropertyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointProperty) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) Elem() PrivateEndpointPropertyOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) PrivateEndpointProperty {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointProperty
		return ret
	}).(PrivateEndpointPropertyOutput)
}

func (o PrivateEndpointPropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) Elem() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) PrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointPropertyResponse
		return ret
	}).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkScopedResourceResponse struct {
	ResourceId *string `pulumi:"resourceId"`
	ScopeId    *string `pulumi:"scopeId"`
}

type PrivateLinkScopedResourceResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkScopedResourceResponseOutput) ScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ScopeId }).(pulumi.StringPtrOutput)
}

type PrivateLinkScopedResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkScopedResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkScopedResourceResponse {
		return vs[0].([]PrivateLinkScopedResourceResponse)[vs[1].(int)]
	}).(PrivateLinkScopedResourceResponseOutput)
}

type PrivateLinkServiceConnectionStateProperty struct {
	Description string `pulumi:"description"`
	Status      string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput
	ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyOutput
}

type PrivateLinkServiceConnectionStatePropertyArgs struct {
	Description pulumi.StringInput `pulumi:"description"`
	Status      pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput).ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput
	ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput
}

type privateLinkServiceConnectionStatePropertyPtrType PrivateLinkServiceConnectionStatePropertyArgs

func PrivateLinkServiceConnectionStatePropertyPtr(v *PrivateLinkServiceConnectionStatePropertyArgs) PrivateLinkServiceConnectionStatePropertyPtrInput {
	return (*privateLinkServiceConnectionStatePropertyPtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateProperty) *PrivateLinkServiceConnectionStateProperty {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) PrivateLinkServiceConnectionStateProperty {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateProperty
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type Recurrence struct {
	Frequency RecurrenceFrequency `pulumi:"frequency"`
	Schedule  RecurrentSchedule   `pulumi:"schedule"`
}





type RecurrenceInput interface {
	pulumi.Input

	ToRecurrenceOutput() RecurrenceOutput
	ToRecurrenceOutputWithContext(context.Context) RecurrenceOutput
}

type RecurrenceArgs struct {
	Frequency RecurrenceFrequencyInput `pulumi:"frequency"`
	Schedule  RecurrentScheduleInput   `pulumi:"schedule"`
}

func (RecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Recurrence)(nil)).Elem()
}

func (i RecurrenceArgs) ToRecurrenceOutput() RecurrenceOutput {
	return i.ToRecurrenceOutputWithContext(context.Background())
}

func (i RecurrenceArgs) ToRecurrenceOutputWithContext(ctx context.Context) RecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceOutput)
}

func (i RecurrenceArgs) ToRecurrencePtrOutput() RecurrencePtrOutput {
	return i.ToRecurrencePtrOutputWithContext(context.Background())
}

func (i RecurrenceArgs) ToRecurrencePtrOutputWithContext(ctx context.Context) RecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceOutput).ToRecurrencePtrOutputWithContext(ctx)
}









type RecurrencePtrInput interface {
	pulumi.Input

	ToRecurrencePtrOutput() RecurrencePtrOutput
	ToRecurrencePtrOutputWithContext(context.Context) RecurrencePtrOutput
}

type recurrencePtrType RecurrenceArgs

func RecurrencePtr(v *RecurrenceArgs) RecurrencePtrInput {
	return (*recurrencePtrType)(v)
}

func (*recurrencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Recurrence)(nil)).Elem()
}

func (i *recurrencePtrType) ToRecurrencePtrOutput() RecurrencePtrOutput {
	return i.ToRecurrencePtrOutputWithContext(context.Background())
}

func (i *recurrencePtrType) ToRecurrencePtrOutputWithContext(ctx context.Context) RecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrencePtrOutput)
}

type RecurrenceOutput struct{ *pulumi.OutputState }

func (RecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Recurrence)(nil)).Elem()
}

func (o RecurrenceOutput) ToRecurrenceOutput() RecurrenceOutput {
	return o
}

func (o RecurrenceOutput) ToRecurrenceOutputWithContext(ctx context.Context) RecurrenceOutput {
	return o
}

func (o RecurrenceOutput) ToRecurrencePtrOutput() RecurrencePtrOutput {
	return o.ToRecurrencePtrOutputWithContext(context.Background())
}

func (o RecurrenceOutput) ToRecurrencePtrOutputWithContext(ctx context.Context) RecurrencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Recurrence) *Recurrence {
		return &v
	}).(RecurrencePtrOutput)
}

func (o RecurrenceOutput) Frequency() RecurrenceFrequencyOutput {
	return o.ApplyT(func(v Recurrence) RecurrenceFrequency { return v.Frequency }).(RecurrenceFrequencyOutput)
}

func (o RecurrenceOutput) Schedule() RecurrentScheduleOutput {
	return o.ApplyT(func(v Recurrence) RecurrentSchedule { return v.Schedule }).(RecurrentScheduleOutput)
}

type RecurrencePtrOutput struct{ *pulumi.OutputState }

func (RecurrencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Recurrence)(nil)).Elem()
}

func (o RecurrencePtrOutput) ToRecurrencePtrOutput() RecurrencePtrOutput {
	return o
}

func (o RecurrencePtrOutput) ToRecurrencePtrOutputWithContext(ctx context.Context) RecurrencePtrOutput {
	return o
}

func (o RecurrencePtrOutput) Elem() RecurrenceOutput {
	return o.ApplyT(func(v *Recurrence) Recurrence {
		if v != nil {
			return *v
		}
		var ret Recurrence
		return ret
	}).(RecurrenceOutput)
}

func (o RecurrencePtrOutput) Frequency() RecurrenceFrequencyPtrOutput {
	return o.ApplyT(func(v *Recurrence) *RecurrenceFrequency {
		if v == nil {
			return nil
		}
		return &v.Frequency
	}).(RecurrenceFrequencyPtrOutput)
}

func (o RecurrencePtrOutput) Schedule() RecurrentSchedulePtrOutput {
	return o.ApplyT(func(v *Recurrence) *RecurrentSchedule {
		if v == nil {
			return nil
		}
		return &v.Schedule
	}).(RecurrentSchedulePtrOutput)
}

type RecurrenceResponse struct {
	Frequency string                    `pulumi:"frequency"`
	Schedule  RecurrentScheduleResponse `pulumi:"schedule"`
}

type RecurrenceResponseOutput struct{ *pulumi.OutputState }

func (RecurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceResponse)(nil)).Elem()
}

func (o RecurrenceResponseOutput) ToRecurrenceResponseOutput() RecurrenceResponseOutput {
	return o
}

func (o RecurrenceResponseOutput) ToRecurrenceResponseOutputWithContext(ctx context.Context) RecurrenceResponseOutput {
	return o
}

func (o RecurrenceResponseOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v RecurrenceResponse) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o RecurrenceResponseOutput) Schedule() RecurrentScheduleResponseOutput {
	return o.ApplyT(func(v RecurrenceResponse) RecurrentScheduleResponse { return v.Schedule }).(RecurrentScheduleResponseOutput)
}

type RecurrenceResponsePtrOutput struct{ *pulumi.OutputState }

func (RecurrenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceResponse)(nil)).Elem()
}

func (o RecurrenceResponsePtrOutput) ToRecurrenceResponsePtrOutput() RecurrenceResponsePtrOutput {
	return o
}

func (o RecurrenceResponsePtrOutput) ToRecurrenceResponsePtrOutputWithContext(ctx context.Context) RecurrenceResponsePtrOutput {
	return o
}

func (o RecurrenceResponsePtrOutput) Elem() RecurrenceResponseOutput {
	return o.ApplyT(func(v *RecurrenceResponse) RecurrenceResponse {
		if v != nil {
			return *v
		}
		var ret RecurrenceResponse
		return ret
	}).(RecurrenceResponseOutput)
}

func (o RecurrenceResponsePtrOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Frequency
	}).(pulumi.StringPtrOutput)
}

func (o RecurrenceResponsePtrOutput) Schedule() RecurrentScheduleResponsePtrOutput {
	return o.ApplyT(func(v *RecurrenceResponse) *RecurrentScheduleResponse {
		if v == nil {
			return nil
		}
		return &v.Schedule
	}).(RecurrentScheduleResponsePtrOutput)
}

type RecurrentSchedule struct {
	Days     []string `pulumi:"days"`
	Hours    []int    `pulumi:"hours"`
	Minutes  []int    `pulumi:"minutes"`
	TimeZone string   `pulumi:"timeZone"`
}





type RecurrentScheduleInput interface {
	pulumi.Input

	ToRecurrentScheduleOutput() RecurrentScheduleOutput
	ToRecurrentScheduleOutputWithContext(context.Context) RecurrentScheduleOutput
}

type RecurrentScheduleArgs struct {
	Days     pulumi.StringArrayInput `pulumi:"days"`
	Hours    pulumi.IntArrayInput    `pulumi:"hours"`
	Minutes  pulumi.IntArrayInput    `pulumi:"minutes"`
	TimeZone pulumi.StringInput      `pulumi:"timeZone"`
}

func (RecurrentScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrentSchedule)(nil)).Elem()
}

func (i RecurrentScheduleArgs) ToRecurrentScheduleOutput() RecurrentScheduleOutput {
	return i.ToRecurrentScheduleOutputWithContext(context.Background())
}

func (i RecurrentScheduleArgs) ToRecurrentScheduleOutputWithContext(ctx context.Context) RecurrentScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrentScheduleOutput)
}

func (i RecurrentScheduleArgs) ToRecurrentSchedulePtrOutput() RecurrentSchedulePtrOutput {
	return i.ToRecurrentSchedulePtrOutputWithContext(context.Background())
}

func (i RecurrentScheduleArgs) ToRecurrentSchedulePtrOutputWithContext(ctx context.Context) RecurrentSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrentScheduleOutput).ToRecurrentSchedulePtrOutputWithContext(ctx)
}









type RecurrentSchedulePtrInput interface {
	pulumi.Input

	ToRecurrentSchedulePtrOutput() RecurrentSchedulePtrOutput
	ToRecurrentSchedulePtrOutputWithContext(context.Context) RecurrentSchedulePtrOutput
}

type recurrentSchedulePtrType RecurrentScheduleArgs

func RecurrentSchedulePtr(v *RecurrentScheduleArgs) RecurrentSchedulePtrInput {
	return (*recurrentSchedulePtrType)(v)
}

func (*recurrentSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrentSchedule)(nil)).Elem()
}

func (i *recurrentSchedulePtrType) ToRecurrentSchedulePtrOutput() RecurrentSchedulePtrOutput {
	return i.ToRecurrentSchedulePtrOutputWithContext(context.Background())
}

func (i *recurrentSchedulePtrType) ToRecurrentSchedulePtrOutputWithContext(ctx context.Context) RecurrentSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrentSchedulePtrOutput)
}

type RecurrentScheduleOutput struct{ *pulumi.OutputState }

func (RecurrentScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrentSchedule)(nil)).Elem()
}

func (o RecurrentScheduleOutput) ToRecurrentScheduleOutput() RecurrentScheduleOutput {
	return o
}

func (o RecurrentScheduleOutput) ToRecurrentScheduleOutputWithContext(ctx context.Context) RecurrentScheduleOutput {
	return o
}

func (o RecurrentScheduleOutput) ToRecurrentSchedulePtrOutput() RecurrentSchedulePtrOutput {
	return o.ToRecurrentSchedulePtrOutputWithContext(context.Background())
}

func (o RecurrentScheduleOutput) ToRecurrentSchedulePtrOutputWithContext(ctx context.Context) RecurrentSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrentSchedule) *RecurrentSchedule {
		return &v
	}).(RecurrentSchedulePtrOutput)
}

func (o RecurrentScheduleOutput) Days() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecurrentSchedule) []string { return v.Days }).(pulumi.StringArrayOutput)
}

func (o RecurrentScheduleOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrentSchedule) []int { return v.Hours }).(pulumi.IntArrayOutput)
}

func (o RecurrentScheduleOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrentSchedule) []int { return v.Minutes }).(pulumi.IntArrayOutput)
}

func (o RecurrentScheduleOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v RecurrentSchedule) string { return v.TimeZone }).(pulumi.StringOutput)
}

type RecurrentSchedulePtrOutput struct{ *pulumi.OutputState }

func (RecurrentSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrentSchedule)(nil)).Elem()
}

func (o RecurrentSchedulePtrOutput) ToRecurrentSchedulePtrOutput() RecurrentSchedulePtrOutput {
	return o
}

func (o RecurrentSchedulePtrOutput) ToRecurrentSchedulePtrOutputWithContext(ctx context.Context) RecurrentSchedulePtrOutput {
	return o
}

func (o RecurrentSchedulePtrOutput) Elem() RecurrentScheduleOutput {
	return o.ApplyT(func(v *RecurrentSchedule) RecurrentSchedule {
		if v != nil {
			return *v
		}
		var ret RecurrentSchedule
		return ret
	}).(RecurrentScheduleOutput)
}

func (o RecurrentSchedulePtrOutput) Days() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RecurrentSchedule) []string {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.StringArrayOutput)
}

func (o RecurrentSchedulePtrOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrentSchedule) []int {
		if v == nil {
			return nil
		}
		return v.Hours
	}).(pulumi.IntArrayOutput)
}

func (o RecurrentSchedulePtrOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrentSchedule) []int {
		if v == nil {
			return nil
		}
		return v.Minutes
	}).(pulumi.IntArrayOutput)
}

func (o RecurrentSchedulePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecurrentSchedule) *string {
		if v == nil {
			return nil
		}
		return &v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type RecurrentScheduleResponse struct {
	Days     []string `pulumi:"days"`
	Hours    []int    `pulumi:"hours"`
	Minutes  []int    `pulumi:"minutes"`
	TimeZone string   `pulumi:"timeZone"`
}

type RecurrentScheduleResponseOutput struct{ *pulumi.OutputState }

func (RecurrentScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrentScheduleResponse)(nil)).Elem()
}

func (o RecurrentScheduleResponseOutput) ToRecurrentScheduleResponseOutput() RecurrentScheduleResponseOutput {
	return o
}

func (o RecurrentScheduleResponseOutput) ToRecurrentScheduleResponseOutputWithContext(ctx context.Context) RecurrentScheduleResponseOutput {
	return o
}

func (o RecurrentScheduleResponseOutput) Days() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecurrentScheduleResponse) []string { return v.Days }).(pulumi.StringArrayOutput)
}

func (o RecurrentScheduleResponseOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrentScheduleResponse) []int { return v.Hours }).(pulumi.IntArrayOutput)
}

func (o RecurrentScheduleResponseOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrentScheduleResponse) []int { return v.Minutes }).(pulumi.IntArrayOutput)
}

func (o RecurrentScheduleResponseOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v RecurrentScheduleResponse) string { return v.TimeZone }).(pulumi.StringOutput)
}

type RecurrentScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (RecurrentScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrentScheduleResponse)(nil)).Elem()
}

func (o RecurrentScheduleResponsePtrOutput) ToRecurrentScheduleResponsePtrOutput() RecurrentScheduleResponsePtrOutput {
	return o
}

func (o RecurrentScheduleResponsePtrOutput) ToRecurrentScheduleResponsePtrOutputWithContext(ctx context.Context) RecurrentScheduleResponsePtrOutput {
	return o
}

func (o RecurrentScheduleResponsePtrOutput) Elem() RecurrentScheduleResponseOutput {
	return o.ApplyT(func(v *RecurrentScheduleResponse) RecurrentScheduleResponse {
		if v != nil {
			return *v
		}
		var ret RecurrentScheduleResponse
		return ret
	}).(RecurrentScheduleResponseOutput)
}

func (o RecurrentScheduleResponsePtrOutput) Days() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RecurrentScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.StringArrayOutput)
}

func (o RecurrentScheduleResponsePtrOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrentScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.Hours
	}).(pulumi.IntArrayOutput)
}

func (o RecurrentScheduleResponsePtrOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrentScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.Minutes
	}).(pulumi.IntArrayOutput)
}

func (o RecurrentScheduleResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecurrentScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type RetentionPolicy struct {
	Days    int  `pulumi:"days"`
	Enabled bool `pulumi:"enabled"`
}





type RetentionPolicyInput interface {
	pulumi.Input

	ToRetentionPolicyOutput() RetentionPolicyOutput
	ToRetentionPolicyOutputWithContext(context.Context) RetentionPolicyOutput
}

type RetentionPolicyArgs struct {
	Days    pulumi.IntInput  `pulumi:"days"`
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (RetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return i.ToRetentionPolicyOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput)
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput).ToRetentionPolicyPtrOutputWithContext(ctx)
}









type RetentionPolicyPtrInput interface {
	pulumi.Input

	ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput
	ToRetentionPolicyPtrOutputWithContext(context.Context) RetentionPolicyPtrOutput
}

type retentionPolicyPtrType RetentionPolicyArgs

func RetentionPolicyPtr(v *RetentionPolicyArgs) RetentionPolicyPtrInput {
	return (*retentionPolicyPtrType)(v)
}

func (*retentionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyPtrOutput)
}

type RetentionPolicyOutput struct{ *pulumi.OutputState }

func (RetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicy) *RetentionPolicy {
		return &v
	}).(RetentionPolicyPtrOutput)
}

func (o RetentionPolicyOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicy) int { return v.Days }).(pulumi.IntOutput)
}

func (o RetentionPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicy) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RetentionPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) Elem() RetentionPolicyOutput {
	return o.ApplyT(func(v *RetentionPolicy) RetentionPolicy {
		if v != nil {
			return *v
		}
		var ret RetentionPolicy
		return ret
	}).(RetentionPolicyOutput)
}

func (o RetentionPolicyPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type RetentionPolicyResponse struct {
	Days    int  `pulumi:"days"`
	Enabled bool `pulumi:"enabled"`
}

type RetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) int { return v.Days }).(pulumi.IntOutput)
}

func (o RetentionPolicyResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RetentionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) Elem() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) RetentionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetentionPolicyResponse
		return ret
	}).(RetentionPolicyResponseOutput)
}

func (o RetentionPolicyResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type RuleEmailAction struct {
	CustomEmails        []string `pulumi:"customEmails"`
	OdataType           string   `pulumi:"odataType"`
	SendToServiceOwners *bool    `pulumi:"sendToServiceOwners"`
}

type RuleEmailActionResponse struct {
	CustomEmails        []string `pulumi:"customEmails"`
	OdataType           string   `pulumi:"odataType"`
	SendToServiceOwners *bool    `pulumi:"sendToServiceOwners"`
}

type RuleManagementEventClaimsDataSource struct {
	EmailAddress *string `pulumi:"emailAddress"`
}

type RuleManagementEventClaimsDataSourceResponse struct {
	EmailAddress *string `pulumi:"emailAddress"`
}

type RuleManagementEventDataSource struct {
	Claims               *RuleManagementEventClaimsDataSource `pulumi:"claims"`
	EventName            *string                              `pulumi:"eventName"`
	EventSource          *string                              `pulumi:"eventSource"`
	LegacyResourceId     *string                              `pulumi:"legacyResourceId"`
	Level                *string                              `pulumi:"level"`
	MetricNamespace      *string                              `pulumi:"metricNamespace"`
	OdataType            string                               `pulumi:"odataType"`
	OperationName        *string                              `pulumi:"operationName"`
	ResourceGroupName    *string                              `pulumi:"resourceGroupName"`
	ResourceLocation     *string                              `pulumi:"resourceLocation"`
	ResourceProviderName *string                              `pulumi:"resourceProviderName"`
	ResourceUri          *string                              `pulumi:"resourceUri"`
	Status               *string                              `pulumi:"status"`
	SubStatus            *string                              `pulumi:"subStatus"`
}

type RuleManagementEventDataSourceResponse struct {
	Claims               *RuleManagementEventClaimsDataSourceResponse `pulumi:"claims"`
	EventName            *string                                      `pulumi:"eventName"`
	EventSource          *string                                      `pulumi:"eventSource"`
	LegacyResourceId     *string                                      `pulumi:"legacyResourceId"`
	Level                *string                                      `pulumi:"level"`
	MetricNamespace      *string                                      `pulumi:"metricNamespace"`
	OdataType            string                                       `pulumi:"odataType"`
	OperationName        *string                                      `pulumi:"operationName"`
	ResourceGroupName    *string                                      `pulumi:"resourceGroupName"`
	ResourceLocation     *string                                      `pulumi:"resourceLocation"`
	ResourceProviderName *string                                      `pulumi:"resourceProviderName"`
	ResourceUri          *string                                      `pulumi:"resourceUri"`
	Status               *string                                      `pulumi:"status"`
	SubStatus            *string                                      `pulumi:"subStatus"`
}

type RuleMetricDataSource struct {
	LegacyResourceId *string `pulumi:"legacyResourceId"`
	MetricName       *string `pulumi:"metricName"`
	MetricNamespace  *string `pulumi:"metricNamespace"`
	OdataType        string  `pulumi:"odataType"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceUri      *string `pulumi:"resourceUri"`
}

type RuleMetricDataSourceResponse struct {
	LegacyResourceId *string `pulumi:"legacyResourceId"`
	MetricName       *string `pulumi:"metricName"`
	MetricNamespace  *string `pulumi:"metricNamespace"`
	OdataType        string  `pulumi:"odataType"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceUri      *string `pulumi:"resourceUri"`
}

type RuleWebhookAction struct {
	OdataType  string            `pulumi:"odataType"`
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}

type RuleWebhookActionResponse struct {
	OdataType  string            `pulumi:"odataType"`
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}

type ScaleAction struct {
	Cooldown  string         `pulumi:"cooldown"`
	Direction ScaleDirection `pulumi:"direction"`
	Type      ScaleType      `pulumi:"type"`
	Value     *string        `pulumi:"value"`
}


func (val *ScaleAction) Defaults() *ScaleAction {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Value) {
		value_ := "1"
		tmp.Value = &value_
	}
	return &tmp
}





type ScaleActionInput interface {
	pulumi.Input

	ToScaleActionOutput() ScaleActionOutput
	ToScaleActionOutputWithContext(context.Context) ScaleActionOutput
}

type ScaleActionArgs struct {
	Cooldown  pulumi.StringInput    `pulumi:"cooldown"`
	Direction ScaleDirectionInput   `pulumi:"direction"`
	Type      ScaleTypeInput        `pulumi:"type"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}


func (val *ScaleActionArgs) Defaults() *ScaleActionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Value) {
		tmp.Value = pulumi.StringPtr("1")
	}
	return &tmp
}
func (ScaleActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleAction)(nil)).Elem()
}

func (i ScaleActionArgs) ToScaleActionOutput() ScaleActionOutput {
	return i.ToScaleActionOutputWithContext(context.Background())
}

func (i ScaleActionArgs) ToScaleActionOutputWithContext(ctx context.Context) ScaleActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleActionOutput)
}

type ScaleActionOutput struct{ *pulumi.OutputState }

func (ScaleActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleAction)(nil)).Elem()
}

func (o ScaleActionOutput) ToScaleActionOutput() ScaleActionOutput {
	return o
}

func (o ScaleActionOutput) ToScaleActionOutputWithContext(ctx context.Context) ScaleActionOutput {
	return o
}

func (o ScaleActionOutput) Cooldown() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleAction) string { return v.Cooldown }).(pulumi.StringOutput)
}

func (o ScaleActionOutput) Direction() ScaleDirectionOutput {
	return o.ApplyT(func(v ScaleAction) ScaleDirection { return v.Direction }).(ScaleDirectionOutput)
}

func (o ScaleActionOutput) Type() ScaleTypeOutput {
	return o.ApplyT(func(v ScaleAction) ScaleType { return v.Type }).(ScaleTypeOutput)
}

func (o ScaleActionOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleAction) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ScaleActionResponse struct {
	Cooldown  string  `pulumi:"cooldown"`
	Direction string  `pulumi:"direction"`
	Type      string  `pulumi:"type"`
	Value     *string `pulumi:"value"`
}


func (val *ScaleActionResponse) Defaults() *ScaleActionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Value) {
		value_ := "1"
		tmp.Value = &value_
	}
	return &tmp
}

type ScaleActionResponseOutput struct{ *pulumi.OutputState }

func (ScaleActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleActionResponse)(nil)).Elem()
}

func (o ScaleActionResponseOutput) ToScaleActionResponseOutput() ScaleActionResponseOutput {
	return o
}

func (o ScaleActionResponseOutput) ToScaleActionResponseOutputWithContext(ctx context.Context) ScaleActionResponseOutput {
	return o
}

func (o ScaleActionResponseOutput) Cooldown() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleActionResponse) string { return v.Cooldown }).(pulumi.StringOutput)
}

func (o ScaleActionResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleActionResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o ScaleActionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleActionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ScaleActionResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleActionResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ScaleCapacity struct {
	Default string `pulumi:"default"`
	Maximum string `pulumi:"maximum"`
	Minimum string `pulumi:"minimum"`
}





type ScaleCapacityInput interface {
	pulumi.Input

	ToScaleCapacityOutput() ScaleCapacityOutput
	ToScaleCapacityOutputWithContext(context.Context) ScaleCapacityOutput
}

type ScaleCapacityArgs struct {
	Default pulumi.StringInput `pulumi:"default"`
	Maximum pulumi.StringInput `pulumi:"maximum"`
	Minimum pulumi.StringInput `pulumi:"minimum"`
}

func (ScaleCapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleCapacity)(nil)).Elem()
}

func (i ScaleCapacityArgs) ToScaleCapacityOutput() ScaleCapacityOutput {
	return i.ToScaleCapacityOutputWithContext(context.Background())
}

func (i ScaleCapacityArgs) ToScaleCapacityOutputWithContext(ctx context.Context) ScaleCapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleCapacityOutput)
}

type ScaleCapacityOutput struct{ *pulumi.OutputState }

func (ScaleCapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleCapacity)(nil)).Elem()
}

func (o ScaleCapacityOutput) ToScaleCapacityOutput() ScaleCapacityOutput {
	return o
}

func (o ScaleCapacityOutput) ToScaleCapacityOutputWithContext(ctx context.Context) ScaleCapacityOutput {
	return o
}

func (o ScaleCapacityOutput) Default() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleCapacity) string { return v.Default }).(pulumi.StringOutput)
}

func (o ScaleCapacityOutput) Maximum() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleCapacity) string { return v.Maximum }).(pulumi.StringOutput)
}

func (o ScaleCapacityOutput) Minimum() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleCapacity) string { return v.Minimum }).(pulumi.StringOutput)
}

type ScaleCapacityResponse struct {
	Default string `pulumi:"default"`
	Maximum string `pulumi:"maximum"`
	Minimum string `pulumi:"minimum"`
}

type ScaleCapacityResponseOutput struct{ *pulumi.OutputState }

func (ScaleCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleCapacityResponse)(nil)).Elem()
}

func (o ScaleCapacityResponseOutput) ToScaleCapacityResponseOutput() ScaleCapacityResponseOutput {
	return o
}

func (o ScaleCapacityResponseOutput) ToScaleCapacityResponseOutputWithContext(ctx context.Context) ScaleCapacityResponseOutput {
	return o
}

func (o ScaleCapacityResponseOutput) Default() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleCapacityResponse) string { return v.Default }).(pulumi.StringOutput)
}

func (o ScaleCapacityResponseOutput) Maximum() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleCapacityResponse) string { return v.Maximum }).(pulumi.StringOutput)
}

func (o ScaleCapacityResponseOutput) Minimum() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleCapacityResponse) string { return v.Minimum }).(pulumi.StringOutput)
}

type ScaleRule struct {
	MetricTrigger MetricTrigger `pulumi:"metricTrigger"`
	ScaleAction   ScaleAction   `pulumi:"scaleAction"`
}


func (val *ScaleRule) Defaults() *ScaleRule {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ScaleAction = *tmp.ScaleAction.Defaults()

	return &tmp
}





type ScaleRuleInput interface {
	pulumi.Input

	ToScaleRuleOutput() ScaleRuleOutput
	ToScaleRuleOutputWithContext(context.Context) ScaleRuleOutput
}

type ScaleRuleArgs struct {
	MetricTrigger MetricTriggerInput `pulumi:"metricTrigger"`
	ScaleAction   ScaleActionInput   `pulumi:"scaleAction"`
}


func (val *ScaleRuleArgs) Defaults() *ScaleRuleArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (ScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRule)(nil)).Elem()
}

func (i ScaleRuleArgs) ToScaleRuleOutput() ScaleRuleOutput {
	return i.ToScaleRuleOutputWithContext(context.Background())
}

func (i ScaleRuleArgs) ToScaleRuleOutputWithContext(ctx context.Context) ScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleOutput)
}





type ScaleRuleArrayInput interface {
	pulumi.Input

	ToScaleRuleArrayOutput() ScaleRuleArrayOutput
	ToScaleRuleArrayOutputWithContext(context.Context) ScaleRuleArrayOutput
}

type ScaleRuleArray []ScaleRuleInput

func (ScaleRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRule)(nil)).Elem()
}

func (i ScaleRuleArray) ToScaleRuleArrayOutput() ScaleRuleArrayOutput {
	return i.ToScaleRuleArrayOutputWithContext(context.Background())
}

func (i ScaleRuleArray) ToScaleRuleArrayOutputWithContext(ctx context.Context) ScaleRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleArrayOutput)
}

type ScaleRuleOutput struct{ *pulumi.OutputState }

func (ScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRule)(nil)).Elem()
}

func (o ScaleRuleOutput) ToScaleRuleOutput() ScaleRuleOutput {
	return o
}

func (o ScaleRuleOutput) ToScaleRuleOutputWithContext(ctx context.Context) ScaleRuleOutput {
	return o
}

func (o ScaleRuleOutput) MetricTrigger() MetricTriggerOutput {
	return o.ApplyT(func(v ScaleRule) MetricTrigger { return v.MetricTrigger }).(MetricTriggerOutput)
}

func (o ScaleRuleOutput) ScaleAction() ScaleActionOutput {
	return o.ApplyT(func(v ScaleRule) ScaleAction { return v.ScaleAction }).(ScaleActionOutput)
}

type ScaleRuleArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRule)(nil)).Elem()
}

func (o ScaleRuleArrayOutput) ToScaleRuleArrayOutput() ScaleRuleArrayOutput {
	return o
}

func (o ScaleRuleArrayOutput) ToScaleRuleArrayOutputWithContext(ctx context.Context) ScaleRuleArrayOutput {
	return o
}

func (o ScaleRuleArrayOutput) Index(i pulumi.IntInput) ScaleRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRule {
		return vs[0].([]ScaleRule)[vs[1].(int)]
	}).(ScaleRuleOutput)
}

type ScaleRuleMetricDimension struct {
	DimensionName string   `pulumi:"dimensionName"`
	Operator      string   `pulumi:"operator"`
	Values        []string `pulumi:"values"`
}





type ScaleRuleMetricDimensionInput interface {
	pulumi.Input

	ToScaleRuleMetricDimensionOutput() ScaleRuleMetricDimensionOutput
	ToScaleRuleMetricDimensionOutputWithContext(context.Context) ScaleRuleMetricDimensionOutput
}

type ScaleRuleMetricDimensionArgs struct {
	DimensionName pulumi.StringInput      `pulumi:"dimensionName"`
	Operator      pulumi.StringInput      `pulumi:"operator"`
	Values        pulumi.StringArrayInput `pulumi:"values"`
}

func (ScaleRuleMetricDimensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleMetricDimension)(nil)).Elem()
}

func (i ScaleRuleMetricDimensionArgs) ToScaleRuleMetricDimensionOutput() ScaleRuleMetricDimensionOutput {
	return i.ToScaleRuleMetricDimensionOutputWithContext(context.Background())
}

func (i ScaleRuleMetricDimensionArgs) ToScaleRuleMetricDimensionOutputWithContext(ctx context.Context) ScaleRuleMetricDimensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleMetricDimensionOutput)
}





type ScaleRuleMetricDimensionArrayInput interface {
	pulumi.Input

	ToScaleRuleMetricDimensionArrayOutput() ScaleRuleMetricDimensionArrayOutput
	ToScaleRuleMetricDimensionArrayOutputWithContext(context.Context) ScaleRuleMetricDimensionArrayOutput
}

type ScaleRuleMetricDimensionArray []ScaleRuleMetricDimensionInput

func (ScaleRuleMetricDimensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleMetricDimension)(nil)).Elem()
}

func (i ScaleRuleMetricDimensionArray) ToScaleRuleMetricDimensionArrayOutput() ScaleRuleMetricDimensionArrayOutput {
	return i.ToScaleRuleMetricDimensionArrayOutputWithContext(context.Background())
}

func (i ScaleRuleMetricDimensionArray) ToScaleRuleMetricDimensionArrayOutputWithContext(ctx context.Context) ScaleRuleMetricDimensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleMetricDimensionArrayOutput)
}

type ScaleRuleMetricDimensionOutput struct{ *pulumi.OutputState }

func (ScaleRuleMetricDimensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleMetricDimension)(nil)).Elem()
}

func (o ScaleRuleMetricDimensionOutput) ToScaleRuleMetricDimensionOutput() ScaleRuleMetricDimensionOutput {
	return o
}

func (o ScaleRuleMetricDimensionOutput) ToScaleRuleMetricDimensionOutputWithContext(ctx context.Context) ScaleRuleMetricDimensionOutput {
	return o
}

func (o ScaleRuleMetricDimensionOutput) DimensionName() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleRuleMetricDimension) string { return v.DimensionName }).(pulumi.StringOutput)
}

func (o ScaleRuleMetricDimensionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleRuleMetricDimension) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ScaleRuleMetricDimensionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScaleRuleMetricDimension) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ScaleRuleMetricDimensionArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleMetricDimensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleMetricDimension)(nil)).Elem()
}

func (o ScaleRuleMetricDimensionArrayOutput) ToScaleRuleMetricDimensionArrayOutput() ScaleRuleMetricDimensionArrayOutput {
	return o
}

func (o ScaleRuleMetricDimensionArrayOutput) ToScaleRuleMetricDimensionArrayOutputWithContext(ctx context.Context) ScaleRuleMetricDimensionArrayOutput {
	return o
}

func (o ScaleRuleMetricDimensionArrayOutput) Index(i pulumi.IntInput) ScaleRuleMetricDimensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRuleMetricDimension {
		return vs[0].([]ScaleRuleMetricDimension)[vs[1].(int)]
	}).(ScaleRuleMetricDimensionOutput)
}

type ScaleRuleMetricDimensionResponse struct {
	DimensionName string   `pulumi:"dimensionName"`
	Operator      string   `pulumi:"operator"`
	Values        []string `pulumi:"values"`
}

type ScaleRuleMetricDimensionResponseOutput struct{ *pulumi.OutputState }

func (ScaleRuleMetricDimensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleMetricDimensionResponse)(nil)).Elem()
}

func (o ScaleRuleMetricDimensionResponseOutput) ToScaleRuleMetricDimensionResponseOutput() ScaleRuleMetricDimensionResponseOutput {
	return o
}

func (o ScaleRuleMetricDimensionResponseOutput) ToScaleRuleMetricDimensionResponseOutputWithContext(ctx context.Context) ScaleRuleMetricDimensionResponseOutput {
	return o
}

func (o ScaleRuleMetricDimensionResponseOutput) DimensionName() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleRuleMetricDimensionResponse) string { return v.DimensionName }).(pulumi.StringOutput)
}

func (o ScaleRuleMetricDimensionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ScaleRuleMetricDimensionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ScaleRuleMetricDimensionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScaleRuleMetricDimensionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ScaleRuleMetricDimensionResponseArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleMetricDimensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleMetricDimensionResponse)(nil)).Elem()
}

func (o ScaleRuleMetricDimensionResponseArrayOutput) ToScaleRuleMetricDimensionResponseArrayOutput() ScaleRuleMetricDimensionResponseArrayOutput {
	return o
}

func (o ScaleRuleMetricDimensionResponseArrayOutput) ToScaleRuleMetricDimensionResponseArrayOutputWithContext(ctx context.Context) ScaleRuleMetricDimensionResponseArrayOutput {
	return o
}

func (o ScaleRuleMetricDimensionResponseArrayOutput) Index(i pulumi.IntInput) ScaleRuleMetricDimensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRuleMetricDimensionResponse {
		return vs[0].([]ScaleRuleMetricDimensionResponse)[vs[1].(int)]
	}).(ScaleRuleMetricDimensionResponseOutput)
}

type ScaleRuleResponse struct {
	MetricTrigger MetricTriggerResponse `pulumi:"metricTrigger"`
	ScaleAction   ScaleActionResponse   `pulumi:"scaleAction"`
}


func (val *ScaleRuleResponse) Defaults() *ScaleRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ScaleAction = *tmp.ScaleAction.Defaults()

	return &tmp
}

type ScaleRuleResponseOutput struct{ *pulumi.OutputState }

func (ScaleRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleResponse)(nil)).Elem()
}

func (o ScaleRuleResponseOutput) ToScaleRuleResponseOutput() ScaleRuleResponseOutput {
	return o
}

func (o ScaleRuleResponseOutput) ToScaleRuleResponseOutputWithContext(ctx context.Context) ScaleRuleResponseOutput {
	return o
}

func (o ScaleRuleResponseOutput) MetricTrigger() MetricTriggerResponseOutput {
	return o.ApplyT(func(v ScaleRuleResponse) MetricTriggerResponse { return v.MetricTrigger }).(MetricTriggerResponseOutput)
}

func (o ScaleRuleResponseOutput) ScaleAction() ScaleActionResponseOutput {
	return o.ApplyT(func(v ScaleRuleResponse) ScaleActionResponse { return v.ScaleAction }).(ScaleActionResponseOutput)
}

type ScaleRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleResponse)(nil)).Elem()
}

func (o ScaleRuleResponseArrayOutput) ToScaleRuleResponseArrayOutput() ScaleRuleResponseArrayOutput {
	return o
}

func (o ScaleRuleResponseArrayOutput) ToScaleRuleResponseArrayOutputWithContext(ctx context.Context) ScaleRuleResponseArrayOutput {
	return o
}

func (o ScaleRuleResponseArrayOutput) Index(i pulumi.IntInput) ScaleRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRuleResponse {
		return vs[0].([]ScaleRuleResponse)[vs[1].(int)]
	}).(ScaleRuleResponseOutput)
}

type Schedule struct {
	FrequencyInMinutes  int `pulumi:"frequencyInMinutes"`
	TimeWindowInMinutes int `pulumi:"timeWindowInMinutes"`
}





type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(context.Context) ScheduleOutput
}

type ScheduleArgs struct {
	FrequencyInMinutes  pulumi.IntInput `pulumi:"frequencyInMinutes"`
	TimeWindowInMinutes pulumi.IntInput `pulumi:"timeWindowInMinutes"`
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (i ScheduleArgs) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

func (i ScheduleArgs) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput).ToSchedulePtrOutputWithContext(ctx)
}









type SchedulePtrInput interface {
	pulumi.Input

	ToSchedulePtrOutput() SchedulePtrOutput
	ToSchedulePtrOutputWithContext(context.Context) SchedulePtrOutput
}

type schedulePtrType ScheduleArgs

func SchedulePtr(v *ScheduleArgs) SchedulePtrInput {
	return (*schedulePtrType)(v)
}

func (*schedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (i *schedulePtrType) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i *schedulePtrType) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulePtrOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o.ToSchedulePtrOutputWithContext(context.Background())
}

func (o ScheduleOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Schedule) *Schedule {
		return &v
	}).(SchedulePtrOutput)
}

func (o ScheduleOutput) FrequencyInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v Schedule) int { return v.FrequencyInMinutes }).(pulumi.IntOutput)
}

func (o ScheduleOutput) TimeWindowInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v Schedule) int { return v.TimeWindowInMinutes }).(pulumi.IntOutput)
}

type SchedulePtrOutput struct{ *pulumi.OutputState }

func (SchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (o SchedulePtrOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) Elem() ScheduleOutput {
	return o.ApplyT(func(v *Schedule) Schedule {
		if v != nil {
			return *v
		}
		var ret Schedule
		return ret
	}).(ScheduleOutput)
}

func (o SchedulePtrOutput) FrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Schedule) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o SchedulePtrOutput) TimeWindowInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Schedule) *int {
		if v == nil {
			return nil
		}
		return &v.TimeWindowInMinutes
	}).(pulumi.IntPtrOutput)
}

type ScheduleResponse struct {
	FrequencyInMinutes  int `pulumi:"frequencyInMinutes"`
	TimeWindowInMinutes int `pulumi:"timeWindowInMinutes"`
}

type ScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseOutput) ToScheduleResponseOutput() ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) FrequencyInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleResponse) int { return v.FrequencyInMinutes }).(pulumi.IntOutput)
}

func (o ScheduleResponseOutput) TimeWindowInMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleResponse) int { return v.TimeWindowInMinutes }).(pulumi.IntOutput)
}

type ScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) Elem() ScheduleResponseOutput {
	return o.ApplyT(func(v *ScheduleResponse) ScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ScheduleResponse
		return ret
	}).(ScheduleResponseOutput)
}

func (o ScheduleResponsePtrOutput) FrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o ScheduleResponsePtrOutput) TimeWindowInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TimeWindowInMinutes
	}).(pulumi.IntPtrOutput)
}

type SinkConfiguration struct {
	Kind string `pulumi:"kind"`
}





type SinkConfigurationInput interface {
	pulumi.Input

	ToSinkConfigurationOutput() SinkConfigurationOutput
	ToSinkConfigurationOutputWithContext(context.Context) SinkConfigurationOutput
}

type SinkConfigurationArgs struct {
	Kind pulumi.StringInput `pulumi:"kind"`
}

func (SinkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SinkConfiguration)(nil)).Elem()
}

func (i SinkConfigurationArgs) ToSinkConfigurationOutput() SinkConfigurationOutput {
	return i.ToSinkConfigurationOutputWithContext(context.Background())
}

func (i SinkConfigurationArgs) ToSinkConfigurationOutputWithContext(ctx context.Context) SinkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SinkConfigurationOutput)
}





type SinkConfigurationArrayInput interface {
	pulumi.Input

	ToSinkConfigurationArrayOutput() SinkConfigurationArrayOutput
	ToSinkConfigurationArrayOutputWithContext(context.Context) SinkConfigurationArrayOutput
}

type SinkConfigurationArray []SinkConfigurationInput

func (SinkConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SinkConfiguration)(nil)).Elem()
}

func (i SinkConfigurationArray) ToSinkConfigurationArrayOutput() SinkConfigurationArrayOutput {
	return i.ToSinkConfigurationArrayOutputWithContext(context.Background())
}

func (i SinkConfigurationArray) ToSinkConfigurationArrayOutputWithContext(ctx context.Context) SinkConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SinkConfigurationArrayOutput)
}

type SinkConfigurationOutput struct{ *pulumi.OutputState }

func (SinkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SinkConfiguration)(nil)).Elem()
}

func (o SinkConfigurationOutput) ToSinkConfigurationOutput() SinkConfigurationOutput {
	return o
}

func (o SinkConfigurationOutput) ToSinkConfigurationOutputWithContext(ctx context.Context) SinkConfigurationOutput {
	return o
}

func (o SinkConfigurationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v SinkConfiguration) string { return v.Kind }).(pulumi.StringOutput)
}

type SinkConfigurationArrayOutput struct{ *pulumi.OutputState }

func (SinkConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SinkConfiguration)(nil)).Elem()
}

func (o SinkConfigurationArrayOutput) ToSinkConfigurationArrayOutput() SinkConfigurationArrayOutput {
	return o
}

func (o SinkConfigurationArrayOutput) ToSinkConfigurationArrayOutputWithContext(ctx context.Context) SinkConfigurationArrayOutput {
	return o
}

func (o SinkConfigurationArrayOutput) Index(i pulumi.IntInput) SinkConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SinkConfiguration {
		return vs[0].([]SinkConfiguration)[vs[1].(int)]
	}).(SinkConfigurationOutput)
}

type SinkConfigurationResponse struct {
	Kind string `pulumi:"kind"`
}

type SinkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SinkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SinkConfigurationResponse)(nil)).Elem()
}

func (o SinkConfigurationResponseOutput) ToSinkConfigurationResponseOutput() SinkConfigurationResponseOutput {
	return o
}

func (o SinkConfigurationResponseOutput) ToSinkConfigurationResponseOutputWithContext(ctx context.Context) SinkConfigurationResponseOutput {
	return o
}

func (o SinkConfigurationResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v SinkConfigurationResponse) string { return v.Kind }).(pulumi.StringOutput)
}

type SinkConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (SinkConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SinkConfigurationResponse)(nil)).Elem()
}

func (o SinkConfigurationResponseArrayOutput) ToSinkConfigurationResponseArrayOutput() SinkConfigurationResponseArrayOutput {
	return o
}

func (o SinkConfigurationResponseArrayOutput) ToSinkConfigurationResponseArrayOutputWithContext(ctx context.Context) SinkConfigurationResponseArrayOutput {
	return o
}

func (o SinkConfigurationResponseArrayOutput) Index(i pulumi.IntInput) SinkConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SinkConfigurationResponse {
		return vs[0].([]SinkConfigurationResponse)[vs[1].(int)]
	}).(SinkConfigurationResponseOutput)
}

type SmsReceiver struct {
	CountryCode string `pulumi:"countryCode"`
	Name        string `pulumi:"name"`
	PhoneNumber string `pulumi:"phoneNumber"`
}





type SmsReceiverInput interface {
	pulumi.Input

	ToSmsReceiverOutput() SmsReceiverOutput
	ToSmsReceiverOutputWithContext(context.Context) SmsReceiverOutput
}

type SmsReceiverArgs struct {
	CountryCode pulumi.StringInput `pulumi:"countryCode"`
	Name        pulumi.StringInput `pulumi:"name"`
	PhoneNumber pulumi.StringInput `pulumi:"phoneNumber"`
}

func (SmsReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsReceiver)(nil)).Elem()
}

func (i SmsReceiverArgs) ToSmsReceiverOutput() SmsReceiverOutput {
	return i.ToSmsReceiverOutputWithContext(context.Background())
}

func (i SmsReceiverArgs) ToSmsReceiverOutputWithContext(ctx context.Context) SmsReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsReceiverOutput)
}





type SmsReceiverArrayInput interface {
	pulumi.Input

	ToSmsReceiverArrayOutput() SmsReceiverArrayOutput
	ToSmsReceiverArrayOutputWithContext(context.Context) SmsReceiverArrayOutput
}

type SmsReceiverArray []SmsReceiverInput

func (SmsReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SmsReceiver)(nil)).Elem()
}

func (i SmsReceiverArray) ToSmsReceiverArrayOutput() SmsReceiverArrayOutput {
	return i.ToSmsReceiverArrayOutputWithContext(context.Background())
}

func (i SmsReceiverArray) ToSmsReceiverArrayOutputWithContext(ctx context.Context) SmsReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsReceiverArrayOutput)
}

type SmsReceiverOutput struct{ *pulumi.OutputState }

func (SmsReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsReceiver)(nil)).Elem()
}

func (o SmsReceiverOutput) ToSmsReceiverOutput() SmsReceiverOutput {
	return o
}

func (o SmsReceiverOutput) ToSmsReceiverOutputWithContext(ctx context.Context) SmsReceiverOutput {
	return o
}

func (o SmsReceiverOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v SmsReceiver) string { return v.CountryCode }).(pulumi.StringOutput)
}

func (o SmsReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SmsReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o SmsReceiverOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v SmsReceiver) string { return v.PhoneNumber }).(pulumi.StringOutput)
}

type SmsReceiverArrayOutput struct{ *pulumi.OutputState }

func (SmsReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SmsReceiver)(nil)).Elem()
}

func (o SmsReceiverArrayOutput) ToSmsReceiverArrayOutput() SmsReceiverArrayOutput {
	return o
}

func (o SmsReceiverArrayOutput) ToSmsReceiverArrayOutputWithContext(ctx context.Context) SmsReceiverArrayOutput {
	return o
}

func (o SmsReceiverArrayOutput) Index(i pulumi.IntInput) SmsReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SmsReceiver {
		return vs[0].([]SmsReceiver)[vs[1].(int)]
	}).(SmsReceiverOutput)
}

type SmsReceiverResponse struct {
	CountryCode string `pulumi:"countryCode"`
	Name        string `pulumi:"name"`
	PhoneNumber string `pulumi:"phoneNumber"`
	Status      string `pulumi:"status"`
}

type SmsReceiverResponseOutput struct{ *pulumi.OutputState }

func (SmsReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsReceiverResponse)(nil)).Elem()
}

func (o SmsReceiverResponseOutput) ToSmsReceiverResponseOutput() SmsReceiverResponseOutput {
	return o
}

func (o SmsReceiverResponseOutput) ToSmsReceiverResponseOutputWithContext(ctx context.Context) SmsReceiverResponseOutput {
	return o
}

func (o SmsReceiverResponseOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v SmsReceiverResponse) string { return v.CountryCode }).(pulumi.StringOutput)
}

func (o SmsReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SmsReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SmsReceiverResponseOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v SmsReceiverResponse) string { return v.PhoneNumber }).(pulumi.StringOutput)
}

func (o SmsReceiverResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SmsReceiverResponse) string { return v.Status }).(pulumi.StringOutput)
}

type SmsReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (SmsReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SmsReceiverResponse)(nil)).Elem()
}

func (o SmsReceiverResponseArrayOutput) ToSmsReceiverResponseArrayOutput() SmsReceiverResponseArrayOutput {
	return o
}

func (o SmsReceiverResponseArrayOutput) ToSmsReceiverResponseArrayOutputWithContext(ctx context.Context) SmsReceiverResponseArrayOutput {
	return o
}

func (o SmsReceiverResponseArrayOutput) Index(i pulumi.IntInput) SmsReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SmsReceiverResponse {
		return vs[0].([]SmsReceiverResponse)[vs[1].(int)]
	}).(SmsReceiverResponseOutput)
}

type Source struct {
	AuthorizedResources []string `pulumi:"authorizedResources"`
	DataSourceId        string   `pulumi:"dataSourceId"`
	Query               *string  `pulumi:"query"`
	QueryType           *string  `pulumi:"queryType"`
}





type SourceInput interface {
	pulumi.Input

	ToSourceOutput() SourceOutput
	ToSourceOutputWithContext(context.Context) SourceOutput
}

type SourceArgs struct {
	AuthorizedResources pulumi.StringArrayInput `pulumi:"authorizedResources"`
	DataSourceId        pulumi.StringInput      `pulumi:"dataSourceId"`
	Query               pulumi.StringPtrInput   `pulumi:"query"`
	QueryType           pulumi.StringPtrInput   `pulumi:"queryType"`
}

func (SourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (i SourceArgs) ToSourceOutput() SourceOutput {
	return i.ToSourceOutputWithContext(context.Background())
}

func (i SourceArgs) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutput)
}

type SourceOutput struct{ *pulumi.OutputState }

func (SourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Source)(nil)).Elem()
}

func (o SourceOutput) ToSourceOutput() SourceOutput {
	return o
}

func (o SourceOutput) ToSourceOutputWithContext(ctx context.Context) SourceOutput {
	return o
}

func (o SourceOutput) AuthorizedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Source) []string { return v.AuthorizedResources }).(pulumi.StringArrayOutput)
}

func (o SourceOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v Source) string { return v.DataSourceId }).(pulumi.StringOutput)
}

func (o SourceOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Source) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o SourceOutput) QueryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Source) *string { return v.QueryType }).(pulumi.StringPtrOutput)
}

type SourceResponse struct {
	AuthorizedResources []string `pulumi:"authorizedResources"`
	DataSourceId        string   `pulumi:"dataSourceId"`
	Query               *string  `pulumi:"query"`
	QueryType           *string  `pulumi:"queryType"`
}

type SourceResponseOutput struct{ *pulumi.OutputState }

func (SourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceResponse)(nil)).Elem()
}

func (o SourceResponseOutput) ToSourceResponseOutput() SourceResponseOutput {
	return o
}

func (o SourceResponseOutput) ToSourceResponseOutputWithContext(ctx context.Context) SourceResponseOutput {
	return o
}

func (o SourceResponseOutput) AuthorizedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SourceResponse) []string { return v.AuthorizedResources }).(pulumi.StringArrayOutput)
}

func (o SourceResponseOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SourceResponse) string { return v.DataSourceId }).(pulumi.StringOutput)
}

func (o SourceResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o SourceResponseOutput) QueryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceResponse) *string { return v.QueryType }).(pulumi.StringPtrOutput)
}

type SubscriptionLogSettings struct {
	Category *string `pulumi:"category"`
	Enabled  bool    `pulumi:"enabled"`
}





type SubscriptionLogSettingsInput interface {
	pulumi.Input

	ToSubscriptionLogSettingsOutput() SubscriptionLogSettingsOutput
	ToSubscriptionLogSettingsOutputWithContext(context.Context) SubscriptionLogSettingsOutput
}

type SubscriptionLogSettingsArgs struct {
	Category pulumi.StringPtrInput `pulumi:"category"`
	Enabled  pulumi.BoolInput      `pulumi:"enabled"`
}

func (SubscriptionLogSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionLogSettings)(nil)).Elem()
}

func (i SubscriptionLogSettingsArgs) ToSubscriptionLogSettingsOutput() SubscriptionLogSettingsOutput {
	return i.ToSubscriptionLogSettingsOutputWithContext(context.Background())
}

func (i SubscriptionLogSettingsArgs) ToSubscriptionLogSettingsOutputWithContext(ctx context.Context) SubscriptionLogSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionLogSettingsOutput)
}





type SubscriptionLogSettingsArrayInput interface {
	pulumi.Input

	ToSubscriptionLogSettingsArrayOutput() SubscriptionLogSettingsArrayOutput
	ToSubscriptionLogSettingsArrayOutputWithContext(context.Context) SubscriptionLogSettingsArrayOutput
}

type SubscriptionLogSettingsArray []SubscriptionLogSettingsInput

func (SubscriptionLogSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionLogSettings)(nil)).Elem()
}

func (i SubscriptionLogSettingsArray) ToSubscriptionLogSettingsArrayOutput() SubscriptionLogSettingsArrayOutput {
	return i.ToSubscriptionLogSettingsArrayOutputWithContext(context.Background())
}

func (i SubscriptionLogSettingsArray) ToSubscriptionLogSettingsArrayOutputWithContext(ctx context.Context) SubscriptionLogSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionLogSettingsArrayOutput)
}

type SubscriptionLogSettingsOutput struct{ *pulumi.OutputState }

func (SubscriptionLogSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionLogSettings)(nil)).Elem()
}

func (o SubscriptionLogSettingsOutput) ToSubscriptionLogSettingsOutput() SubscriptionLogSettingsOutput {
	return o
}

func (o SubscriptionLogSettingsOutput) ToSubscriptionLogSettingsOutputWithContext(ctx context.Context) SubscriptionLogSettingsOutput {
	return o
}

func (o SubscriptionLogSettingsOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionLogSettings) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o SubscriptionLogSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SubscriptionLogSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type SubscriptionLogSettingsArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionLogSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionLogSettings)(nil)).Elem()
}

func (o SubscriptionLogSettingsArrayOutput) ToSubscriptionLogSettingsArrayOutput() SubscriptionLogSettingsArrayOutput {
	return o
}

func (o SubscriptionLogSettingsArrayOutput) ToSubscriptionLogSettingsArrayOutputWithContext(ctx context.Context) SubscriptionLogSettingsArrayOutput {
	return o
}

func (o SubscriptionLogSettingsArrayOutput) Index(i pulumi.IntInput) SubscriptionLogSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriptionLogSettings {
		return vs[0].([]SubscriptionLogSettings)[vs[1].(int)]
	}).(SubscriptionLogSettingsOutput)
}

type SubscriptionLogSettingsResponse struct {
	Category *string `pulumi:"category"`
	Enabled  bool    `pulumi:"enabled"`
}

type SubscriptionLogSettingsResponseOutput struct{ *pulumi.OutputState }

func (SubscriptionLogSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionLogSettingsResponse)(nil)).Elem()
}

func (o SubscriptionLogSettingsResponseOutput) ToSubscriptionLogSettingsResponseOutput() SubscriptionLogSettingsResponseOutput {
	return o
}

func (o SubscriptionLogSettingsResponseOutput) ToSubscriptionLogSettingsResponseOutputWithContext(ctx context.Context) SubscriptionLogSettingsResponseOutput {
	return o
}

func (o SubscriptionLogSettingsResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionLogSettingsResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o SubscriptionLogSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SubscriptionLogSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type SubscriptionLogSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionLogSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionLogSettingsResponse)(nil)).Elem()
}

func (o SubscriptionLogSettingsResponseArrayOutput) ToSubscriptionLogSettingsResponseArrayOutput() SubscriptionLogSettingsResponseArrayOutput {
	return o
}

func (o SubscriptionLogSettingsResponseArrayOutput) ToSubscriptionLogSettingsResponseArrayOutputWithContext(ctx context.Context) SubscriptionLogSettingsResponseArrayOutput {
	return o
}

func (o SubscriptionLogSettingsResponseArrayOutput) Index(i pulumi.IntInput) SubscriptionLogSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriptionLogSettingsResponse {
		return vs[0].([]SubscriptionLogSettingsResponse)[vs[1].(int)]
	}).(SubscriptionLogSettingsResponseOutput)
}

type SyslogDataSource struct {
	FacilityNames []string `pulumi:"facilityNames"`
	LogLevels     []string `pulumi:"logLevels"`
	Name          *string  `pulumi:"name"`
	Streams       []string `pulumi:"streams"`
}





type SyslogDataSourceInput interface {
	pulumi.Input

	ToSyslogDataSourceOutput() SyslogDataSourceOutput
	ToSyslogDataSourceOutputWithContext(context.Context) SyslogDataSourceOutput
}

type SyslogDataSourceArgs struct {
	FacilityNames pulumi.StringArrayInput `pulumi:"facilityNames"`
	LogLevels     pulumi.StringArrayInput `pulumi:"logLevels"`
	Name          pulumi.StringPtrInput   `pulumi:"name"`
	Streams       pulumi.StringArrayInput `pulumi:"streams"`
}

func (SyslogDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyslogDataSource)(nil)).Elem()
}

func (i SyslogDataSourceArgs) ToSyslogDataSourceOutput() SyslogDataSourceOutput {
	return i.ToSyslogDataSourceOutputWithContext(context.Background())
}

func (i SyslogDataSourceArgs) ToSyslogDataSourceOutputWithContext(ctx context.Context) SyslogDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyslogDataSourceOutput)
}





type SyslogDataSourceArrayInput interface {
	pulumi.Input

	ToSyslogDataSourceArrayOutput() SyslogDataSourceArrayOutput
	ToSyslogDataSourceArrayOutputWithContext(context.Context) SyslogDataSourceArrayOutput
}

type SyslogDataSourceArray []SyslogDataSourceInput

func (SyslogDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyslogDataSource)(nil)).Elem()
}

func (i SyslogDataSourceArray) ToSyslogDataSourceArrayOutput() SyslogDataSourceArrayOutput {
	return i.ToSyslogDataSourceArrayOutputWithContext(context.Background())
}

func (i SyslogDataSourceArray) ToSyslogDataSourceArrayOutputWithContext(ctx context.Context) SyslogDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyslogDataSourceArrayOutput)
}

type SyslogDataSourceOutput struct{ *pulumi.OutputState }

func (SyslogDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyslogDataSource)(nil)).Elem()
}

func (o SyslogDataSourceOutput) ToSyslogDataSourceOutput() SyslogDataSourceOutput {
	return o
}

func (o SyslogDataSourceOutput) ToSyslogDataSourceOutputWithContext(ctx context.Context) SyslogDataSourceOutput {
	return o
}

func (o SyslogDataSourceOutput) FacilityNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSource) []string { return v.FacilityNames }).(pulumi.StringArrayOutput)
}

func (o SyslogDataSourceOutput) LogLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSource) []string { return v.LogLevels }).(pulumi.StringArrayOutput)
}

func (o SyslogDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyslogDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SyslogDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type SyslogDataSourceArrayOutput struct{ *pulumi.OutputState }

func (SyslogDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyslogDataSource)(nil)).Elem()
}

func (o SyslogDataSourceArrayOutput) ToSyslogDataSourceArrayOutput() SyslogDataSourceArrayOutput {
	return o
}

func (o SyslogDataSourceArrayOutput) ToSyslogDataSourceArrayOutputWithContext(ctx context.Context) SyslogDataSourceArrayOutput {
	return o
}

func (o SyslogDataSourceArrayOutput) Index(i pulumi.IntInput) SyslogDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyslogDataSource {
		return vs[0].([]SyslogDataSource)[vs[1].(int)]
	}).(SyslogDataSourceOutput)
}

type SyslogDataSourceResponse struct {
	FacilityNames []string `pulumi:"facilityNames"`
	LogLevels     []string `pulumi:"logLevels"`
	Name          *string  `pulumi:"name"`
	Streams       []string `pulumi:"streams"`
}

type SyslogDataSourceResponseOutput struct{ *pulumi.OutputState }

func (SyslogDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyslogDataSourceResponse)(nil)).Elem()
}

func (o SyslogDataSourceResponseOutput) ToSyslogDataSourceResponseOutput() SyslogDataSourceResponseOutput {
	return o
}

func (o SyslogDataSourceResponseOutput) ToSyslogDataSourceResponseOutputWithContext(ctx context.Context) SyslogDataSourceResponseOutput {
	return o
}

func (o SyslogDataSourceResponseOutput) FacilityNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSourceResponse) []string { return v.FacilityNames }).(pulumi.StringArrayOutput)
}

func (o SyslogDataSourceResponseOutput) LogLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSourceResponse) []string { return v.LogLevels }).(pulumi.StringArrayOutput)
}

func (o SyslogDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyslogDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SyslogDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SyslogDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

type SyslogDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SyslogDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyslogDataSourceResponse)(nil)).Elem()
}

func (o SyslogDataSourceResponseArrayOutput) ToSyslogDataSourceResponseArrayOutput() SyslogDataSourceResponseArrayOutput {
	return o
}

func (o SyslogDataSourceResponseArrayOutput) ToSyslogDataSourceResponseArrayOutputWithContext(ctx context.Context) SyslogDataSourceResponseArrayOutput {
	return o
}

func (o SyslogDataSourceResponseArrayOutput) Index(i pulumi.IntInput) SyslogDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyslogDataSourceResponse {
		return vs[0].([]SyslogDataSourceResponse)[vs[1].(int)]
	}).(SyslogDataSourceResponseOutput)
}

type ThresholdRuleCondition struct {
	DataSource      interface{}              `pulumi:"dataSource"`
	OdataType       string                   `pulumi:"odataType"`
	Operator        ConditionOperator        `pulumi:"operator"`
	Threshold       float64                  `pulumi:"threshold"`
	TimeAggregation *TimeAggregationOperator `pulumi:"timeAggregation"`
	WindowSize      *string                  `pulumi:"windowSize"`
}

type ThresholdRuleConditionResponse struct {
	DataSource      interface{} `pulumi:"dataSource"`
	OdataType       string      `pulumi:"odataType"`
	Operator        string      `pulumi:"operator"`
	Threshold       float64     `pulumi:"threshold"`
	TimeAggregation *string     `pulumi:"timeAggregation"`
	WindowSize      *string     `pulumi:"windowSize"`
}

type TimeWindow struct {
	End      string  `pulumi:"end"`
	Start    string  `pulumi:"start"`
	TimeZone *string `pulumi:"timeZone"`
}





type TimeWindowInput interface {
	pulumi.Input

	ToTimeWindowOutput() TimeWindowOutput
	ToTimeWindowOutputWithContext(context.Context) TimeWindowOutput
}

type TimeWindowArgs struct {
	End      pulumi.StringInput    `pulumi:"end"`
	Start    pulumi.StringInput    `pulumi:"start"`
	TimeZone pulumi.StringPtrInput `pulumi:"timeZone"`
}

func (TimeWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindow)(nil)).Elem()
}

func (i TimeWindowArgs) ToTimeWindowOutput() TimeWindowOutput {
	return i.ToTimeWindowOutputWithContext(context.Background())
}

func (i TimeWindowArgs) ToTimeWindowOutputWithContext(ctx context.Context) TimeWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowOutput)
}

func (i TimeWindowArgs) ToTimeWindowPtrOutput() TimeWindowPtrOutput {
	return i.ToTimeWindowPtrOutputWithContext(context.Background())
}

func (i TimeWindowArgs) ToTimeWindowPtrOutputWithContext(ctx context.Context) TimeWindowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowOutput).ToTimeWindowPtrOutputWithContext(ctx)
}









type TimeWindowPtrInput interface {
	pulumi.Input

	ToTimeWindowPtrOutput() TimeWindowPtrOutput
	ToTimeWindowPtrOutputWithContext(context.Context) TimeWindowPtrOutput
}

type timeWindowPtrType TimeWindowArgs

func TimeWindowPtr(v *TimeWindowArgs) TimeWindowPtrInput {
	return (*timeWindowPtrType)(v)
}

func (*timeWindowPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeWindow)(nil)).Elem()
}

func (i *timeWindowPtrType) ToTimeWindowPtrOutput() TimeWindowPtrOutput {
	return i.ToTimeWindowPtrOutputWithContext(context.Background())
}

func (i *timeWindowPtrType) ToTimeWindowPtrOutputWithContext(ctx context.Context) TimeWindowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowPtrOutput)
}

type TimeWindowOutput struct{ *pulumi.OutputState }

func (TimeWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindow)(nil)).Elem()
}

func (o TimeWindowOutput) ToTimeWindowOutput() TimeWindowOutput {
	return o
}

func (o TimeWindowOutput) ToTimeWindowOutputWithContext(ctx context.Context) TimeWindowOutput {
	return o
}

func (o TimeWindowOutput) ToTimeWindowPtrOutput() TimeWindowPtrOutput {
	return o.ToTimeWindowPtrOutputWithContext(context.Background())
}

func (o TimeWindowOutput) ToTimeWindowPtrOutputWithContext(ctx context.Context) TimeWindowPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeWindow) *TimeWindow {
		return &v
	}).(TimeWindowPtrOutput)
}

func (o TimeWindowOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindow) string { return v.End }).(pulumi.StringOutput)
}

func (o TimeWindowOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindow) string { return v.Start }).(pulumi.StringOutput)
}

func (o TimeWindowOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeWindow) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type TimeWindowPtrOutput struct{ *pulumi.OutputState }

func (TimeWindowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeWindow)(nil)).Elem()
}

func (o TimeWindowPtrOutput) ToTimeWindowPtrOutput() TimeWindowPtrOutput {
	return o
}

func (o TimeWindowPtrOutput) ToTimeWindowPtrOutputWithContext(ctx context.Context) TimeWindowPtrOutput {
	return o
}

func (o TimeWindowPtrOutput) Elem() TimeWindowOutput {
	return o.ApplyT(func(v *TimeWindow) TimeWindow {
		if v != nil {
			return *v
		}
		var ret TimeWindow
		return ret
	}).(TimeWindowOutput)
}

func (o TimeWindowPtrOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeWindow) *string {
		if v == nil {
			return nil
		}
		return &v.End
	}).(pulumi.StringPtrOutput)
}

func (o TimeWindowPtrOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeWindow) *string {
		if v == nil {
			return nil
		}
		return &v.Start
	}).(pulumi.StringPtrOutput)
}

func (o TimeWindowPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeWindow) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type TimeWindowResponse struct {
	End      string  `pulumi:"end"`
	Start    string  `pulumi:"start"`
	TimeZone *string `pulumi:"timeZone"`
}

type TimeWindowResponseOutput struct{ *pulumi.OutputState }

func (TimeWindowResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowResponse)(nil)).Elem()
}

func (o TimeWindowResponseOutput) ToTimeWindowResponseOutput() TimeWindowResponseOutput {
	return o
}

func (o TimeWindowResponseOutput) ToTimeWindowResponseOutputWithContext(ctx context.Context) TimeWindowResponseOutput {
	return o
}

func (o TimeWindowResponseOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowResponse) string { return v.End }).(pulumi.StringOutput)
}

func (o TimeWindowResponseOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v TimeWindowResponse) string { return v.Start }).(pulumi.StringOutput)
}

func (o TimeWindowResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimeWindowResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type TimeWindowResponsePtrOutput struct{ *pulumi.OutputState }

func (TimeWindowResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeWindowResponse)(nil)).Elem()
}

func (o TimeWindowResponsePtrOutput) ToTimeWindowResponsePtrOutput() TimeWindowResponsePtrOutput {
	return o
}

func (o TimeWindowResponsePtrOutput) ToTimeWindowResponsePtrOutputWithContext(ctx context.Context) TimeWindowResponsePtrOutput {
	return o
}

func (o TimeWindowResponsePtrOutput) Elem() TimeWindowResponseOutput {
	return o.ApplyT(func(v *TimeWindowResponse) TimeWindowResponse {
		if v != nil {
			return *v
		}
		var ret TimeWindowResponse
		return ret
	}).(TimeWindowResponseOutput)
}

func (o TimeWindowResponsePtrOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeWindowResponse) *string {
		if v == nil {
			return nil
		}
		return &v.End
	}).(pulumi.StringPtrOutput)
}

func (o TimeWindowResponsePtrOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeWindowResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Start
	}).(pulumi.StringPtrOutput)
}

func (o TimeWindowResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeWindowResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type TriggerCondition struct {
	MetricTrigger     *LogMetricTrigger `pulumi:"metricTrigger"`
	Threshold         float64           `pulumi:"threshold"`
	ThresholdOperator string            `pulumi:"thresholdOperator"`
}

type TriggerConditionResponse struct {
	MetricTrigger     *LogMetricTriggerResponse `pulumi:"metricTrigger"`
	Threshold         float64                   `pulumi:"threshold"`
	ThresholdOperator string                    `pulumi:"thresholdOperator"`
}

type VoiceReceiver struct {
	CountryCode string `pulumi:"countryCode"`
	Name        string `pulumi:"name"`
	PhoneNumber string `pulumi:"phoneNumber"`
}





type VoiceReceiverInput interface {
	pulumi.Input

	ToVoiceReceiverOutput() VoiceReceiverOutput
	ToVoiceReceiverOutputWithContext(context.Context) VoiceReceiverOutput
}

type VoiceReceiverArgs struct {
	CountryCode pulumi.StringInput `pulumi:"countryCode"`
	Name        pulumi.StringInput `pulumi:"name"`
	PhoneNumber pulumi.StringInput `pulumi:"phoneNumber"`
}

func (VoiceReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VoiceReceiver)(nil)).Elem()
}

func (i VoiceReceiverArgs) ToVoiceReceiverOutput() VoiceReceiverOutput {
	return i.ToVoiceReceiverOutputWithContext(context.Background())
}

func (i VoiceReceiverArgs) ToVoiceReceiverOutputWithContext(ctx context.Context) VoiceReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceReceiverOutput)
}





type VoiceReceiverArrayInput interface {
	pulumi.Input

	ToVoiceReceiverArrayOutput() VoiceReceiverArrayOutput
	ToVoiceReceiverArrayOutputWithContext(context.Context) VoiceReceiverArrayOutput
}

type VoiceReceiverArray []VoiceReceiverInput

func (VoiceReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VoiceReceiver)(nil)).Elem()
}

func (i VoiceReceiverArray) ToVoiceReceiverArrayOutput() VoiceReceiverArrayOutput {
	return i.ToVoiceReceiverArrayOutputWithContext(context.Background())
}

func (i VoiceReceiverArray) ToVoiceReceiverArrayOutputWithContext(ctx context.Context) VoiceReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceReceiverArrayOutput)
}

type VoiceReceiverOutput struct{ *pulumi.OutputState }

func (VoiceReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VoiceReceiver)(nil)).Elem()
}

func (o VoiceReceiverOutput) ToVoiceReceiverOutput() VoiceReceiverOutput {
	return o
}

func (o VoiceReceiverOutput) ToVoiceReceiverOutputWithContext(ctx context.Context) VoiceReceiverOutput {
	return o
}

func (o VoiceReceiverOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v VoiceReceiver) string { return v.CountryCode }).(pulumi.StringOutput)
}

func (o VoiceReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VoiceReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o VoiceReceiverOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v VoiceReceiver) string { return v.PhoneNumber }).(pulumi.StringOutput)
}

type VoiceReceiverArrayOutput struct{ *pulumi.OutputState }

func (VoiceReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VoiceReceiver)(nil)).Elem()
}

func (o VoiceReceiverArrayOutput) ToVoiceReceiverArrayOutput() VoiceReceiverArrayOutput {
	return o
}

func (o VoiceReceiverArrayOutput) ToVoiceReceiverArrayOutputWithContext(ctx context.Context) VoiceReceiverArrayOutput {
	return o
}

func (o VoiceReceiverArrayOutput) Index(i pulumi.IntInput) VoiceReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VoiceReceiver {
		return vs[0].([]VoiceReceiver)[vs[1].(int)]
	}).(VoiceReceiverOutput)
}

type VoiceReceiverResponse struct {
	CountryCode string `pulumi:"countryCode"`
	Name        string `pulumi:"name"`
	PhoneNumber string `pulumi:"phoneNumber"`
}

type VoiceReceiverResponseOutput struct{ *pulumi.OutputState }

func (VoiceReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VoiceReceiverResponse)(nil)).Elem()
}

func (o VoiceReceiverResponseOutput) ToVoiceReceiverResponseOutput() VoiceReceiverResponseOutput {
	return o
}

func (o VoiceReceiverResponseOutput) ToVoiceReceiverResponseOutputWithContext(ctx context.Context) VoiceReceiverResponseOutput {
	return o
}

func (o VoiceReceiverResponseOutput) CountryCode() pulumi.StringOutput {
	return o.ApplyT(func(v VoiceReceiverResponse) string { return v.CountryCode }).(pulumi.StringOutput)
}

func (o VoiceReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VoiceReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VoiceReceiverResponseOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v VoiceReceiverResponse) string { return v.PhoneNumber }).(pulumi.StringOutput)
}

type VoiceReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (VoiceReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VoiceReceiverResponse)(nil)).Elem()
}

func (o VoiceReceiverResponseArrayOutput) ToVoiceReceiverResponseArrayOutput() VoiceReceiverResponseArrayOutput {
	return o
}

func (o VoiceReceiverResponseArrayOutput) ToVoiceReceiverResponseArrayOutputWithContext(ctx context.Context) VoiceReceiverResponseArrayOutput {
	return o
}

func (o VoiceReceiverResponseArrayOutput) Index(i pulumi.IntInput) VoiceReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VoiceReceiverResponse {
		return vs[0].([]VoiceReceiverResponse)[vs[1].(int)]
	}).(VoiceReceiverResponseOutput)
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

type WebhookNotification struct {
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}





type WebhookNotificationInput interface {
	pulumi.Input

	ToWebhookNotificationOutput() WebhookNotificationOutput
	ToWebhookNotificationOutputWithContext(context.Context) WebhookNotificationOutput
}

type WebhookNotificationArgs struct {
	Properties pulumi.StringMapInput `pulumi:"properties"`
	ServiceUri pulumi.StringPtrInput `pulumi:"serviceUri"`
}

func (WebhookNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookNotification)(nil)).Elem()
}

func (i WebhookNotificationArgs) ToWebhookNotificationOutput() WebhookNotificationOutput {
	return i.ToWebhookNotificationOutputWithContext(context.Background())
}

func (i WebhookNotificationArgs) ToWebhookNotificationOutputWithContext(ctx context.Context) WebhookNotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookNotificationOutput)
}





type WebhookNotificationArrayInput interface {
	pulumi.Input

	ToWebhookNotificationArrayOutput() WebhookNotificationArrayOutput
	ToWebhookNotificationArrayOutputWithContext(context.Context) WebhookNotificationArrayOutput
}

type WebhookNotificationArray []WebhookNotificationInput

func (WebhookNotificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebhookNotification)(nil)).Elem()
}

func (i WebhookNotificationArray) ToWebhookNotificationArrayOutput() WebhookNotificationArrayOutput {
	return i.ToWebhookNotificationArrayOutputWithContext(context.Background())
}

func (i WebhookNotificationArray) ToWebhookNotificationArrayOutputWithContext(ctx context.Context) WebhookNotificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookNotificationArrayOutput)
}

type WebhookNotificationOutput struct{ *pulumi.OutputState }

func (WebhookNotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookNotification)(nil)).Elem()
}

func (o WebhookNotificationOutput) ToWebhookNotificationOutput() WebhookNotificationOutput {
	return o
}

func (o WebhookNotificationOutput) ToWebhookNotificationOutputWithContext(ctx context.Context) WebhookNotificationOutput {
	return o
}

func (o WebhookNotificationOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v WebhookNotification) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o WebhookNotificationOutput) ServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookNotification) *string { return v.ServiceUri }).(pulumi.StringPtrOutput)
}

type WebhookNotificationArrayOutput struct{ *pulumi.OutputState }

func (WebhookNotificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebhookNotification)(nil)).Elem()
}

func (o WebhookNotificationArrayOutput) ToWebhookNotificationArrayOutput() WebhookNotificationArrayOutput {
	return o
}

func (o WebhookNotificationArrayOutput) ToWebhookNotificationArrayOutputWithContext(ctx context.Context) WebhookNotificationArrayOutput {
	return o
}

func (o WebhookNotificationArrayOutput) Index(i pulumi.IntInput) WebhookNotificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebhookNotification {
		return vs[0].([]WebhookNotification)[vs[1].(int)]
	}).(WebhookNotificationOutput)
}

type WebhookNotificationResponse struct {
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}

type WebhookNotificationResponseOutput struct{ *pulumi.OutputState }

func (WebhookNotificationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookNotificationResponse)(nil)).Elem()
}

func (o WebhookNotificationResponseOutput) ToWebhookNotificationResponseOutput() WebhookNotificationResponseOutput {
	return o
}

func (o WebhookNotificationResponseOutput) ToWebhookNotificationResponseOutputWithContext(ctx context.Context) WebhookNotificationResponseOutput {
	return o
}

func (o WebhookNotificationResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v WebhookNotificationResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o WebhookNotificationResponseOutput) ServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookNotificationResponse) *string { return v.ServiceUri }).(pulumi.StringPtrOutput)
}

type WebhookNotificationResponseArrayOutput struct{ *pulumi.OutputState }

func (WebhookNotificationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebhookNotificationResponse)(nil)).Elem()
}

func (o WebhookNotificationResponseArrayOutput) ToWebhookNotificationResponseArrayOutput() WebhookNotificationResponseArrayOutput {
	return o
}

func (o WebhookNotificationResponseArrayOutput) ToWebhookNotificationResponseArrayOutputWithContext(ctx context.Context) WebhookNotificationResponseArrayOutput {
	return o
}

func (o WebhookNotificationResponseArrayOutput) Index(i pulumi.IntInput) WebhookNotificationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebhookNotificationResponse {
		return vs[0].([]WebhookNotificationResponse)[vs[1].(int)]
	}).(WebhookNotificationResponseOutput)
}

type WebhookReceiver struct {
	IdentifierUri        *string `pulumi:"identifierUri"`
	Name                 string  `pulumi:"name"`
	ObjectId             *string `pulumi:"objectId"`
	ServiceUri           string  `pulumi:"serviceUri"`
	TenantId             *string `pulumi:"tenantId"`
	UseAadAuth           *bool   `pulumi:"useAadAuth"`
	UseCommonAlertSchema *bool   `pulumi:"useCommonAlertSchema"`
}


func (val *WebhookReceiver) Defaults() *WebhookReceiver {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseAadAuth) {
		useAadAuth_ := false
		tmp.UseAadAuth = &useAadAuth_
	}
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}





type WebhookReceiverInput interface {
	pulumi.Input

	ToWebhookReceiverOutput() WebhookReceiverOutput
	ToWebhookReceiverOutputWithContext(context.Context) WebhookReceiverOutput
}

type WebhookReceiverArgs struct {
	IdentifierUri        pulumi.StringPtrInput `pulumi:"identifierUri"`
	Name                 pulumi.StringInput    `pulumi:"name"`
	ObjectId             pulumi.StringPtrInput `pulumi:"objectId"`
	ServiceUri           pulumi.StringInput    `pulumi:"serviceUri"`
	TenantId             pulumi.StringPtrInput `pulumi:"tenantId"`
	UseAadAuth           pulumi.BoolPtrInput   `pulumi:"useAadAuth"`
	UseCommonAlertSchema pulumi.BoolPtrInput   `pulumi:"useCommonAlertSchema"`
}


func (val *WebhookReceiverArgs) Defaults() *WebhookReceiverArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseAadAuth) {
		tmp.UseAadAuth = pulumi.BoolPtr(false)
	}
	if isZero(tmp.UseCommonAlertSchema) {
		tmp.UseCommonAlertSchema = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (WebhookReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookReceiver)(nil)).Elem()
}

func (i WebhookReceiverArgs) ToWebhookReceiverOutput() WebhookReceiverOutput {
	return i.ToWebhookReceiverOutputWithContext(context.Background())
}

func (i WebhookReceiverArgs) ToWebhookReceiverOutputWithContext(ctx context.Context) WebhookReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookReceiverOutput)
}





type WebhookReceiverArrayInput interface {
	pulumi.Input

	ToWebhookReceiverArrayOutput() WebhookReceiverArrayOutput
	ToWebhookReceiverArrayOutputWithContext(context.Context) WebhookReceiverArrayOutput
}

type WebhookReceiverArray []WebhookReceiverInput

func (WebhookReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebhookReceiver)(nil)).Elem()
}

func (i WebhookReceiverArray) ToWebhookReceiverArrayOutput() WebhookReceiverArrayOutput {
	return i.ToWebhookReceiverArrayOutputWithContext(context.Background())
}

func (i WebhookReceiverArray) ToWebhookReceiverArrayOutputWithContext(ctx context.Context) WebhookReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookReceiverArrayOutput)
}

type WebhookReceiverOutput struct{ *pulumi.OutputState }

func (WebhookReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookReceiver)(nil)).Elem()
}

func (o WebhookReceiverOutput) ToWebhookReceiverOutput() WebhookReceiverOutput {
	return o
}

func (o WebhookReceiverOutput) ToWebhookReceiverOutputWithContext(ctx context.Context) WebhookReceiverOutput {
	return o
}

func (o WebhookReceiverOutput) IdentifierUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookReceiver) *string { return v.IdentifierUri }).(pulumi.StringPtrOutput)
}

func (o WebhookReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o WebhookReceiverOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookReceiver) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o WebhookReceiverOutput) ServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookReceiver) string { return v.ServiceUri }).(pulumi.StringOutput)
}

func (o WebhookReceiverOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookReceiver) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o WebhookReceiverOutput) UseAadAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebhookReceiver) *bool { return v.UseAadAuth }).(pulumi.BoolPtrOutput)
}

func (o WebhookReceiverOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebhookReceiver) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type WebhookReceiverArrayOutput struct{ *pulumi.OutputState }

func (WebhookReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebhookReceiver)(nil)).Elem()
}

func (o WebhookReceiverArrayOutput) ToWebhookReceiverArrayOutput() WebhookReceiverArrayOutput {
	return o
}

func (o WebhookReceiverArrayOutput) ToWebhookReceiverArrayOutputWithContext(ctx context.Context) WebhookReceiverArrayOutput {
	return o
}

func (o WebhookReceiverArrayOutput) Index(i pulumi.IntInput) WebhookReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebhookReceiver {
		return vs[0].([]WebhookReceiver)[vs[1].(int)]
	}).(WebhookReceiverOutput)
}

type WebhookReceiverResponse struct {
	IdentifierUri        *string `pulumi:"identifierUri"`
	Name                 string  `pulumi:"name"`
	ObjectId             *string `pulumi:"objectId"`
	ServiceUri           string  `pulumi:"serviceUri"`
	TenantId             *string `pulumi:"tenantId"`
	UseAadAuth           *bool   `pulumi:"useAadAuth"`
	UseCommonAlertSchema *bool   `pulumi:"useCommonAlertSchema"`
}


func (val *WebhookReceiverResponse) Defaults() *WebhookReceiverResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.UseAadAuth) {
		useAadAuth_ := false
		tmp.UseAadAuth = &useAadAuth_
	}
	if isZero(tmp.UseCommonAlertSchema) {
		useCommonAlertSchema_ := false
		tmp.UseCommonAlertSchema = &useCommonAlertSchema_
	}
	return &tmp
}

type WebhookReceiverResponseOutput struct{ *pulumi.OutputState }

func (WebhookReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookReceiverResponse)(nil)).Elem()
}

func (o WebhookReceiverResponseOutput) ToWebhookReceiverResponseOutput() WebhookReceiverResponseOutput {
	return o
}

func (o WebhookReceiverResponseOutput) ToWebhookReceiverResponseOutputWithContext(ctx context.Context) WebhookReceiverResponseOutput {
	return o
}

func (o WebhookReceiverResponseOutput) IdentifierUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookReceiverResponse) *string { return v.IdentifierUri }).(pulumi.StringPtrOutput)
}

func (o WebhookReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o WebhookReceiverResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookReceiverResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o WebhookReceiverResponseOutput) ServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookReceiverResponse) string { return v.ServiceUri }).(pulumi.StringOutput)
}

func (o WebhookReceiverResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookReceiverResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o WebhookReceiverResponseOutput) UseAadAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebhookReceiverResponse) *bool { return v.UseAadAuth }).(pulumi.BoolPtrOutput)
}

func (o WebhookReceiverResponseOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebhookReceiverResponse) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type WebhookReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (WebhookReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebhookReceiverResponse)(nil)).Elem()
}

func (o WebhookReceiverResponseArrayOutput) ToWebhookReceiverResponseArrayOutput() WebhookReceiverResponseArrayOutput {
	return o
}

func (o WebhookReceiverResponseArrayOutput) ToWebhookReceiverResponseArrayOutputWithContext(ctx context.Context) WebhookReceiverResponseArrayOutput {
	return o
}

func (o WebhookReceiverResponseArrayOutput) Index(i pulumi.IntInput) WebhookReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebhookReceiverResponse {
		return vs[0].([]WebhookReceiverResponse)[vs[1].(int)]
	}).(WebhookReceiverResponseOutput)
}

type WebtestLocationAvailabilityCriteria struct {
	ComponentId         string  `pulumi:"componentId"`
	FailedLocationCount float64 `pulumi:"failedLocationCount"`
	OdataType           string  `pulumi:"odataType"`
	WebTestId           string  `pulumi:"webTestId"`
}

type WebtestLocationAvailabilityCriteriaResponse struct {
	ComponentId         string  `pulumi:"componentId"`
	FailedLocationCount float64 `pulumi:"failedLocationCount"`
	OdataType           string  `pulumi:"odataType"`
	WebTestId           string  `pulumi:"webTestId"`
}

type WindowsEventLogDataSource struct {
	Name         *string  `pulumi:"name"`
	Streams      []string `pulumi:"streams"`
	XPathQueries []string `pulumi:"xPathQueries"`
}





type WindowsEventLogDataSourceInput interface {
	pulumi.Input

	ToWindowsEventLogDataSourceOutput() WindowsEventLogDataSourceOutput
	ToWindowsEventLogDataSourceOutputWithContext(context.Context) WindowsEventLogDataSourceOutput
}

type WindowsEventLogDataSourceArgs struct {
	Name         pulumi.StringPtrInput   `pulumi:"name"`
	Streams      pulumi.StringArrayInput `pulumi:"streams"`
	XPathQueries pulumi.StringArrayInput `pulumi:"xPathQueries"`
}

func (WindowsEventLogDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsEventLogDataSource)(nil)).Elem()
}

func (i WindowsEventLogDataSourceArgs) ToWindowsEventLogDataSourceOutput() WindowsEventLogDataSourceOutput {
	return i.ToWindowsEventLogDataSourceOutputWithContext(context.Background())
}

func (i WindowsEventLogDataSourceArgs) ToWindowsEventLogDataSourceOutputWithContext(ctx context.Context) WindowsEventLogDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsEventLogDataSourceOutput)
}





type WindowsEventLogDataSourceArrayInput interface {
	pulumi.Input

	ToWindowsEventLogDataSourceArrayOutput() WindowsEventLogDataSourceArrayOutput
	ToWindowsEventLogDataSourceArrayOutputWithContext(context.Context) WindowsEventLogDataSourceArrayOutput
}

type WindowsEventLogDataSourceArray []WindowsEventLogDataSourceInput

func (WindowsEventLogDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsEventLogDataSource)(nil)).Elem()
}

func (i WindowsEventLogDataSourceArray) ToWindowsEventLogDataSourceArrayOutput() WindowsEventLogDataSourceArrayOutput {
	return i.ToWindowsEventLogDataSourceArrayOutputWithContext(context.Background())
}

func (i WindowsEventLogDataSourceArray) ToWindowsEventLogDataSourceArrayOutputWithContext(ctx context.Context) WindowsEventLogDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsEventLogDataSourceArrayOutput)
}

type WindowsEventLogDataSourceOutput struct{ *pulumi.OutputState }

func (WindowsEventLogDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsEventLogDataSource)(nil)).Elem()
}

func (o WindowsEventLogDataSourceOutput) ToWindowsEventLogDataSourceOutput() WindowsEventLogDataSourceOutput {
	return o
}

func (o WindowsEventLogDataSourceOutput) ToWindowsEventLogDataSourceOutputWithContext(ctx context.Context) WindowsEventLogDataSourceOutput {
	return o
}

func (o WindowsEventLogDataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsEventLogDataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WindowsEventLogDataSourceOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsEventLogDataSource) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

func (o WindowsEventLogDataSourceOutput) XPathQueries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsEventLogDataSource) []string { return v.XPathQueries }).(pulumi.StringArrayOutput)
}

type WindowsEventLogDataSourceArrayOutput struct{ *pulumi.OutputState }

func (WindowsEventLogDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsEventLogDataSource)(nil)).Elem()
}

func (o WindowsEventLogDataSourceArrayOutput) ToWindowsEventLogDataSourceArrayOutput() WindowsEventLogDataSourceArrayOutput {
	return o
}

func (o WindowsEventLogDataSourceArrayOutput) ToWindowsEventLogDataSourceArrayOutputWithContext(ctx context.Context) WindowsEventLogDataSourceArrayOutput {
	return o
}

func (o WindowsEventLogDataSourceArrayOutput) Index(i pulumi.IntInput) WindowsEventLogDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WindowsEventLogDataSource {
		return vs[0].([]WindowsEventLogDataSource)[vs[1].(int)]
	}).(WindowsEventLogDataSourceOutput)
}

type WindowsEventLogDataSourceResponse struct {
	Name         *string  `pulumi:"name"`
	Streams      []string `pulumi:"streams"`
	XPathQueries []string `pulumi:"xPathQueries"`
}

type WindowsEventLogDataSourceResponseOutput struct{ *pulumi.OutputState }

func (WindowsEventLogDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsEventLogDataSourceResponse)(nil)).Elem()
}

func (o WindowsEventLogDataSourceResponseOutput) ToWindowsEventLogDataSourceResponseOutput() WindowsEventLogDataSourceResponseOutput {
	return o
}

func (o WindowsEventLogDataSourceResponseOutput) ToWindowsEventLogDataSourceResponseOutputWithContext(ctx context.Context) WindowsEventLogDataSourceResponseOutput {
	return o
}

func (o WindowsEventLogDataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsEventLogDataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WindowsEventLogDataSourceResponseOutput) Streams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsEventLogDataSourceResponse) []string { return v.Streams }).(pulumi.StringArrayOutput)
}

func (o WindowsEventLogDataSourceResponseOutput) XPathQueries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsEventLogDataSourceResponse) []string { return v.XPathQueries }).(pulumi.StringArrayOutput)
}

type WindowsEventLogDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (WindowsEventLogDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsEventLogDataSourceResponse)(nil)).Elem()
}

func (o WindowsEventLogDataSourceResponseArrayOutput) ToWindowsEventLogDataSourceResponseArrayOutput() WindowsEventLogDataSourceResponseArrayOutput {
	return o
}

func (o WindowsEventLogDataSourceResponseArrayOutput) ToWindowsEventLogDataSourceResponseArrayOutputWithContext(ctx context.Context) WindowsEventLogDataSourceResponseArrayOutput {
	return o
}

func (o WindowsEventLogDataSourceResponseArrayOutput) Index(i pulumi.IntInput) WindowsEventLogDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WindowsEventLogDataSourceResponse {
		return vs[0].([]WindowsEventLogDataSourceResponse)[vs[1].(int)]
	}).(WindowsEventLogDataSourceResponseOutput)
}

type WorkbookManagedIdentity struct {
	Type *string `pulumi:"type"`
}





type WorkbookManagedIdentityInput interface {
	pulumi.Input

	ToWorkbookManagedIdentityOutput() WorkbookManagedIdentityOutput
	ToWorkbookManagedIdentityOutputWithContext(context.Context) WorkbookManagedIdentityOutput
}

type WorkbookManagedIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (WorkbookManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookManagedIdentity)(nil)).Elem()
}

func (i WorkbookManagedIdentityArgs) ToWorkbookManagedIdentityOutput() WorkbookManagedIdentityOutput {
	return i.ToWorkbookManagedIdentityOutputWithContext(context.Background())
}

func (i WorkbookManagedIdentityArgs) ToWorkbookManagedIdentityOutputWithContext(ctx context.Context) WorkbookManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityOutput)
}

func (i WorkbookManagedIdentityArgs) ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput {
	return i.ToWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (i WorkbookManagedIdentityArgs) ToWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityOutput).ToWorkbookManagedIdentityPtrOutputWithContext(ctx)
}









type WorkbookManagedIdentityPtrInput interface {
	pulumi.Input

	ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput
	ToWorkbookManagedIdentityPtrOutputWithContext(context.Context) WorkbookManagedIdentityPtrOutput
}

type workbookManagedIdentityPtrType WorkbookManagedIdentityArgs

func WorkbookManagedIdentityPtr(v *WorkbookManagedIdentityArgs) WorkbookManagedIdentityPtrInput {
	return (*workbookManagedIdentityPtrType)(v)
}

func (*workbookManagedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookManagedIdentity)(nil)).Elem()
}

func (i *workbookManagedIdentityPtrType) ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput {
	return i.ToWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *workbookManagedIdentityPtrType) ToWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityPtrOutput)
}

type WorkbookManagedIdentityOutput struct{ *pulumi.OutputState }

func (WorkbookManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookManagedIdentity)(nil)).Elem()
}

func (o WorkbookManagedIdentityOutput) ToWorkbookManagedIdentityOutput() WorkbookManagedIdentityOutput {
	return o
}

func (o WorkbookManagedIdentityOutput) ToWorkbookManagedIdentityOutputWithContext(ctx context.Context) WorkbookManagedIdentityOutput {
	return o
}

func (o WorkbookManagedIdentityOutput) ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput {
	return o.ToWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (o WorkbookManagedIdentityOutput) ToWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkbookManagedIdentity) *WorkbookManagedIdentity {
		return &v
	}).(WorkbookManagedIdentityPtrOutput)
}

func (o WorkbookManagedIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookManagedIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkbookManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (WorkbookManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookManagedIdentity)(nil)).Elem()
}

func (o WorkbookManagedIdentityPtrOutput) ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput {
	return o
}

func (o WorkbookManagedIdentityPtrOutput) ToWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityPtrOutput {
	return o
}

func (o WorkbookManagedIdentityPtrOutput) Elem() WorkbookManagedIdentityOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentity) WorkbookManagedIdentity {
		if v != nil {
			return *v
		}
		var ret WorkbookManagedIdentity
		return ret
	}).(WorkbookManagedIdentityOutput)
}

func (o WorkbookManagedIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type WorkbookManagedIdentityResponse struct {
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities *WorkbookUserAssignedIdentitiesResponse `pulumi:"userAssignedIdentities"`
}

type WorkbookManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (WorkbookManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookManagedIdentityResponse)(nil)).Elem()
}

func (o WorkbookManagedIdentityResponseOutput) ToWorkbookManagedIdentityResponseOutput() WorkbookManagedIdentityResponseOutput {
	return o
}

func (o WorkbookManagedIdentityResponseOutput) ToWorkbookManagedIdentityResponseOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponseOutput {
	return o
}

func (o WorkbookManagedIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookManagedIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o WorkbookManagedIdentityResponseOutput) UserAssignedIdentities() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyT(func(v WorkbookManagedIdentityResponse) *WorkbookUserAssignedIdentitiesResponse {
		return v.UserAssignedIdentities
	}).(WorkbookUserAssignedIdentitiesResponsePtrOutput)
}

type WorkbookManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkbookManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookManagedIdentityResponse)(nil)).Elem()
}

func (o WorkbookManagedIdentityResponsePtrOutput) ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput {
	return o
}

func (o WorkbookManagedIdentityResponsePtrOutput) ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponsePtrOutput {
	return o
}

func (o WorkbookManagedIdentityResponsePtrOutput) Elem() WorkbookManagedIdentityResponseOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentityResponse) WorkbookManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret WorkbookManagedIdentityResponse
		return ret
	}).(WorkbookManagedIdentityResponseOutput)
}

func (o WorkbookManagedIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o WorkbookManagedIdentityResponsePtrOutput) UserAssignedIdentities() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentityResponse) *WorkbookUserAssignedIdentitiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(WorkbookUserAssignedIdentitiesResponsePtrOutput)
}

type WorkbookTemplateGallery struct {
	Category     *string `pulumi:"category"`
	Name         *string `pulumi:"name"`
	Order        *int    `pulumi:"order"`
	ResourceType *string `pulumi:"resourceType"`
	Type         *string `pulumi:"type"`
}





type WorkbookTemplateGalleryInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryOutput() WorkbookTemplateGalleryOutput
	ToWorkbookTemplateGalleryOutputWithContext(context.Context) WorkbookTemplateGalleryOutput
}

type WorkbookTemplateGalleryArgs struct {
	Category     pulumi.StringPtrInput `pulumi:"category"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Order        pulumi.IntPtrInput    `pulumi:"order"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (WorkbookTemplateGalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGallery)(nil)).Elem()
}

func (i WorkbookTemplateGalleryArgs) ToWorkbookTemplateGalleryOutput() WorkbookTemplateGalleryOutput {
	return i.ToWorkbookTemplateGalleryOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryArgs) ToWorkbookTemplateGalleryOutputWithContext(ctx context.Context) WorkbookTemplateGalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryOutput)
}





type WorkbookTemplateGalleryArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryArrayOutput() WorkbookTemplateGalleryArrayOutput
	ToWorkbookTemplateGalleryArrayOutputWithContext(context.Context) WorkbookTemplateGalleryArrayOutput
}

type WorkbookTemplateGalleryArray []WorkbookTemplateGalleryInput

func (WorkbookTemplateGalleryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGallery)(nil)).Elem()
}

func (i WorkbookTemplateGalleryArray) ToWorkbookTemplateGalleryArrayOutput() WorkbookTemplateGalleryArrayOutput {
	return i.ToWorkbookTemplateGalleryArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryArray) ToWorkbookTemplateGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryArrayOutput)
}

type WorkbookTemplateGalleryOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGallery)(nil)).Elem()
}

func (o WorkbookTemplateGalleryOutput) ToWorkbookTemplateGalleryOutput() WorkbookTemplateGalleryOutput {
	return o
}

func (o WorkbookTemplateGalleryOutput) ToWorkbookTemplateGalleryOutputWithContext(ctx context.Context) WorkbookTemplateGalleryOutput {
	return o
}

func (o WorkbookTemplateGalleryOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkbookTemplateGalleryArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGallery)(nil)).Elem()
}

func (o WorkbookTemplateGalleryArrayOutput) ToWorkbookTemplateGalleryArrayOutput() WorkbookTemplateGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryArrayOutput) ToWorkbookTemplateGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateGalleryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateGallery {
		return vs[0].([]WorkbookTemplateGallery)[vs[1].(int)]
	}).(WorkbookTemplateGalleryOutput)
}

type WorkbookTemplateGalleryResponse struct {
	Category     *string `pulumi:"category"`
	Name         *string `pulumi:"name"`
	Order        *int    `pulumi:"order"`
	ResourceType *string `pulumi:"resourceType"`
	Type         *string `pulumi:"type"`
}

type WorkbookTemplateGalleryResponseOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateGalleryResponseOutput) ToWorkbookTemplateGalleryResponseOutput() WorkbookTemplateGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseOutput) ToWorkbookTemplateGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkbookTemplateGalleryResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateGalleryResponseArrayOutput) ToWorkbookTemplateGalleryResponseArrayOutput() WorkbookTemplateGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseArrayOutput) ToWorkbookTemplateGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateGalleryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateGalleryResponse {
		return vs[0].([]WorkbookTemplateGalleryResponse)[vs[1].(int)]
	}).(WorkbookTemplateGalleryResponseOutput)
}

type WorkbookTemplateLocalizedGallery struct {
	Galleries    []WorkbookTemplateGallery `pulumi:"galleries"`
	TemplateData interface{}               `pulumi:"templateData"`
}





type WorkbookTemplateLocalizedGalleryInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryOutput() WorkbookTemplateLocalizedGalleryOutput
	ToWorkbookTemplateLocalizedGalleryOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryOutput
}

type WorkbookTemplateLocalizedGalleryArgs struct {
	Galleries    WorkbookTemplateGalleryArrayInput `pulumi:"galleries"`
	TemplateData pulumi.Input                      `pulumi:"templateData"`
}

func (WorkbookTemplateLocalizedGalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArgs) ToWorkbookTemplateLocalizedGalleryOutput() WorkbookTemplateLocalizedGalleryOutput {
	return i.ToWorkbookTemplateLocalizedGalleryOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArgs) ToWorkbookTemplateLocalizedGalleryOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryOutput)
}





type WorkbookTemplateLocalizedGalleryArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryArrayOutput() WorkbookTemplateLocalizedGalleryArrayOutput
	ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryArrayOutput
}

type WorkbookTemplateLocalizedGalleryArray []WorkbookTemplateLocalizedGalleryInput

func (WorkbookTemplateLocalizedGalleryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArray) ToWorkbookTemplateLocalizedGalleryArrayOutput() WorkbookTemplateLocalizedGalleryArrayOutput {
	return i.ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArray) ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryArrayOutput)
}

type WorkbookTemplateLocalizedGalleryOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryOutput) ToWorkbookTemplateLocalizedGalleryOutput() WorkbookTemplateLocalizedGalleryOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryOutput) ToWorkbookTemplateLocalizedGalleryOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryOutput) Galleries() WorkbookTemplateGalleryArrayOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGallery) []WorkbookTemplateGallery { return v.Galleries }).(WorkbookTemplateGalleryArrayOutput)
}

func (o WorkbookTemplateLocalizedGalleryOutput) TemplateData() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGallery) interface{} { return v.TemplateData }).(pulumi.AnyOutput)
}

type WorkbookTemplateLocalizedGalleryArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryArrayOutput) ToWorkbookTemplateLocalizedGalleryArrayOutput() WorkbookTemplateLocalizedGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayOutput) ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateLocalizedGalleryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGallery {
		return vs[0].([]WorkbookTemplateLocalizedGallery)[vs[1].(int)]
	}).(WorkbookTemplateLocalizedGalleryOutput)
}

type WorkbookTemplateLocalizedGalleryResponse struct {
	Galleries    []WorkbookTemplateGalleryResponse `pulumi:"galleries"`
	TemplateData interface{}                       `pulumi:"templateData"`
}

type WorkbookTemplateLocalizedGalleryResponseOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) ToWorkbookTemplateLocalizedGalleryResponseOutput() WorkbookTemplateLocalizedGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) Galleries() WorkbookTemplateGalleryResponseArrayOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGalleryResponse) []WorkbookTemplateGalleryResponse { return v.Galleries }).(WorkbookTemplateGalleryResponseArrayOutput)
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) TemplateData() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGalleryResponse) interface{} { return v.TemplateData }).(pulumi.AnyOutput)
}

type WorkbookTemplateLocalizedGalleryResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayOutput() WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateLocalizedGalleryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGalleryResponse {
		return vs[0].([]WorkbookTemplateLocalizedGalleryResponse)[vs[1].(int)]
	}).(WorkbookTemplateLocalizedGalleryResponseOutput)
}

type WorkbookUserAssignedIdentitiesResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
}

type WorkbookUserAssignedIdentitiesResponseOutput struct{ *pulumi.OutputState }

func (WorkbookUserAssignedIdentitiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) ToWorkbookUserAssignedIdentitiesResponseOutput() WorkbookUserAssignedIdentitiesResponseOutput {
	return o
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) ToWorkbookUserAssignedIdentitiesResponseOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponseOutput {
	return o
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookUserAssignedIdentitiesResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookUserAssignedIdentitiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookUserAssignedIdentitiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type WorkbookUserAssignedIdentitiesResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkbookUserAssignedIdentitiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) Elem() WorkbookUserAssignedIdentitiesResponseOutput {
	return o.ApplyT(func(v *WorkbookUserAssignedIdentitiesResponse) WorkbookUserAssignedIdentitiesResponse {
		if v != nil {
			return *v
		}
		var ret WorkbookUserAssignedIdentitiesResponse
		return ret
	}).(WorkbookUserAssignedIdentitiesResponseOutput)
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

type WorkbookTemplateLocalizedGalleryArrayMap map[string]WorkbookTemplateLocalizedGalleryArrayInput

func (WorkbookTemplateLocalizedGalleryArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string][]WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArrayMap) ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return i.ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArrayMap) ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryArrayMapOutput)
}





type WorkbookTemplateLocalizedGalleryArrayMapInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput
	ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput
}

type WorkbookTemplateLocalizedGalleryArrayMapOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string][]WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) MapIndex(k pulumi.StringInput) WorkbookTemplateLocalizedGalleryArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) []WorkbookTemplateLocalizedGallery {
		return vs[0].(map[string][]WorkbookTemplateLocalizedGallery)[vs[1].(string)]
	}).(WorkbookTemplateLocalizedGalleryArrayOutput)
}

type WorkbookTemplateLocalizedGalleryResponseArrayMapOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string][]WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) MapIndex(k pulumi.StringInput) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) []WorkbookTemplateLocalizedGalleryResponse {
		return vs[0].(map[string][]WorkbookTemplateLocalizedGalleryResponse)[vs[1].(string)]
	}).(WorkbookTemplateLocalizedGalleryResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionGroupTypeOutput{})
	pulumi.RegisterOutputType(ActionGroupTypeArrayOutput{})
	pulumi.RegisterOutputType(ActionGroupResponseOutput{})
	pulumi.RegisterOutputType(ActionGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(ActionListOutput{})
	pulumi.RegisterOutputType(ActionListResponseOutput{})
	pulumi.RegisterOutputType(AlertRuleAllOfConditionOutput{})
	pulumi.RegisterOutputType(AlertRuleAllOfConditionResponseOutput{})
	pulumi.RegisterOutputType(AlertRuleAnyOfOrLeafConditionOutput{})
	pulumi.RegisterOutputType(AlertRuleAnyOfOrLeafConditionArrayOutput{})
	pulumi.RegisterOutputType(AlertRuleAnyOfOrLeafConditionResponseOutput{})
	pulumi.RegisterOutputType(AlertRuleAnyOfOrLeafConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(AlertRuleLeafConditionOutput{})
	pulumi.RegisterOutputType(AlertRuleLeafConditionArrayOutput{})
	pulumi.RegisterOutputType(AlertRuleLeafConditionResponseOutput{})
	pulumi.RegisterOutputType(AlertRuleLeafConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentDataVolumeCapOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentDataVolumeCapPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentDataVolumeCapResponseOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentDataVolumeCapResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitionsPtrOutput{})
	pulumi.RegisterOutputType(ArmRoleReceiverOutput{})
	pulumi.RegisterOutputType(ArmRoleReceiverArrayOutput{})
	pulumi.RegisterOutputType(ArmRoleReceiverResponseOutput{})
	pulumi.RegisterOutputType(ArmRoleReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverArrayOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverResponseOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleNotificationOutput{})
	pulumi.RegisterOutputType(AutoscaleNotificationArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleNotificationResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleNotificationResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleProfileOutput{})
	pulumi.RegisterOutputType(AutoscaleProfileArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleProfileResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverArrayOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverResponseOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverArrayOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverResponseOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointNetworkAclsOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointNetworkAclsPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResourceResponseSystemDataOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseConfigurationAccessOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseConfigurationAccessPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseLogsIngestionOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseLogsIngestionPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseNetworkAclsOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResponseNetworkAclsPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDataSourcesOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDataSourcesPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDestinationsOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleDestinationsPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDataSourcesOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDataSourcesPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDestinationsOutput{})
	pulumi.RegisterOutputType(DataCollectionRuleResponseDestinationsPtrOutput{})
	pulumi.RegisterOutputType(DataFlowOutput{})
	pulumi.RegisterOutputType(DataFlowArrayOutput{})
	pulumi.RegisterOutputType(DataFlowResponseOutput{})
	pulumi.RegisterOutputType(DataFlowResponseArrayOutput{})
	pulumi.RegisterOutputType(DataSourceOutput{})
	pulumi.RegisterOutputType(DataSourceArrayOutput{})
	pulumi.RegisterOutputType(DataSourceConfigurationOutput{})
	pulumi.RegisterOutputType(DataSourceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(DataSourceResponseOutput{})
	pulumi.RegisterOutputType(DataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(DestinationsSpecAzureMonitorMetricsOutput{})
	pulumi.RegisterOutputType(DestinationsSpecAzureMonitorMetricsPtrOutput{})
	pulumi.RegisterOutputType(DestinationsSpecResponseAzureMonitorMetricsOutput{})
	pulumi.RegisterOutputType(DestinationsSpecResponseAzureMonitorMetricsPtrOutput{})
	pulumi.RegisterOutputType(EmailNotificationOutput{})
	pulumi.RegisterOutputType(EmailNotificationPtrOutput{})
	pulumi.RegisterOutputType(EmailNotificationResponseOutput{})
	pulumi.RegisterOutputType(EmailNotificationResponsePtrOutput{})
	pulumi.RegisterOutputType(EmailReceiverOutput{})
	pulumi.RegisterOutputType(EmailReceiverArrayOutput{})
	pulumi.RegisterOutputType(EmailReceiverResponseOutput{})
	pulumi.RegisterOutputType(EmailReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(EtwEventConfigurationOutput{})
	pulumi.RegisterOutputType(EtwEventConfigurationArrayOutput{})
	pulumi.RegisterOutputType(EtwEventConfigurationResponseOutput{})
	pulumi.RegisterOutputType(EtwEventConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(EtwProviderConfigurationOutput{})
	pulumi.RegisterOutputType(EtwProviderConfigurationArrayOutput{})
	pulumi.RegisterOutputType(EtwProviderConfigurationResponseOutput{})
	pulumi.RegisterOutputType(EtwProviderConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(EventLogConfigurationOutput{})
	pulumi.RegisterOutputType(EventLogConfigurationArrayOutput{})
	pulumi.RegisterOutputType(EventLogConfigurationResponseOutput{})
	pulumi.RegisterOutputType(EventLogConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtensionDataSourceOutput{})
	pulumi.RegisterOutputType(ExtensionDataSourceArrayOutput{})
	pulumi.RegisterOutputType(ExtensionDataSourceResponseOutput{})
	pulumi.RegisterOutputType(ExtensionDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(ItsmReceiverOutput{})
	pulumi.RegisterOutputType(ItsmReceiverArrayOutput{})
	pulumi.RegisterOutputType(ItsmReceiverResponseOutput{})
	pulumi.RegisterOutputType(ItsmReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationArrayOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationResponseArrayOutput{})
	pulumi.RegisterOutputType(LogSettingsOutput{})
	pulumi.RegisterOutputType(LogSettingsArrayOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverArrayOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverResponseOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricAlertActionOutput{})
	pulumi.RegisterOutputType(MetricAlertActionArrayOutput{})
	pulumi.RegisterOutputType(MetricAlertActionResponseOutput{})
	pulumi.RegisterOutputType(MetricAlertActionResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricSettingsOutput{})
	pulumi.RegisterOutputType(MetricSettingsArrayOutput{})
	pulumi.RegisterOutputType(MetricSettingsResponseOutput{})
	pulumi.RegisterOutputType(MetricSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricTriggerOutput{})
	pulumi.RegisterOutputType(MetricTriggerResponseOutput{})
	pulumi.RegisterOutputType(MyWorkbookManagedIdentityOutput{})
	pulumi.RegisterOutputType(MyWorkbookManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(MyWorkbookManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(MyWorkbookManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MyWorkbookUserAssignedIdentitiesResponseOutput{})
	pulumi.RegisterOutputType(MyWorkbookUserAssignedIdentitiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceArrayOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceResponseOutput{})
	pulumi.RegisterOutputType(PerfCounterDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(PerformanceCounterConfigurationOutput{})
	pulumi.RegisterOutputType(PerformanceCounterConfigurationArrayOutput{})
	pulumi.RegisterOutputType(PerformanceCounterConfigurationResponseOutput{})
	pulumi.RegisterOutputType(PerformanceCounterConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(RecurrenceOutput{})
	pulumi.RegisterOutputType(RecurrencePtrOutput{})
	pulumi.RegisterOutputType(RecurrenceResponseOutput{})
	pulumi.RegisterOutputType(RecurrenceResponsePtrOutput{})
	pulumi.RegisterOutputType(RecurrentScheduleOutput{})
	pulumi.RegisterOutputType(RecurrentSchedulePtrOutput{})
	pulumi.RegisterOutputType(RecurrentScheduleResponseOutput{})
	pulumi.RegisterOutputType(RecurrentScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyOutput{})
	pulumi.RegisterOutputType(RetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ScaleActionOutput{})
	pulumi.RegisterOutputType(ScaleActionResponseOutput{})
	pulumi.RegisterOutputType(ScaleCapacityOutput{})
	pulumi.RegisterOutputType(ScaleCapacityResponseOutput{})
	pulumi.RegisterOutputType(ScaleRuleOutput{})
	pulumi.RegisterOutputType(ScaleRuleArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleMetricDimensionOutput{})
	pulumi.RegisterOutputType(ScaleRuleMetricDimensionArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleMetricDimensionResponseOutput{})
	pulumi.RegisterOutputType(ScaleRuleMetricDimensionResponseArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(ScaleRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(SchedulePtrOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(SinkConfigurationOutput{})
	pulumi.RegisterOutputType(SinkConfigurationArrayOutput{})
	pulumi.RegisterOutputType(SinkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SinkConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SmsReceiverOutput{})
	pulumi.RegisterOutputType(SmsReceiverArrayOutput{})
	pulumi.RegisterOutputType(SmsReceiverResponseOutput{})
	pulumi.RegisterOutputType(SmsReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(SourceOutput{})
	pulumi.RegisterOutputType(SourceResponseOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsResponseOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceArrayOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceResponseOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(TimeWindowOutput{})
	pulumi.RegisterOutputType(TimeWindowPtrOutput{})
	pulumi.RegisterOutputType(TimeWindowResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(VoiceReceiverOutput{})
	pulumi.RegisterOutputType(VoiceReceiverArrayOutput{})
	pulumi.RegisterOutputType(VoiceReceiverResponseOutput{})
	pulumi.RegisterOutputType(VoiceReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationArrayOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationResponseOutput{})
	pulumi.RegisterOutputType(WebTestGeolocationResponseArrayOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesConfigurationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseConfigurationOutput{})
	pulumi.RegisterOutputType(WebTestPropertiesResponseConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WebhookNotificationOutput{})
	pulumi.RegisterOutputType(WebhookNotificationArrayOutput{})
	pulumi.RegisterOutputType(WebhookNotificationResponseOutput{})
	pulumi.RegisterOutputType(WebhookNotificationResponseArrayOutput{})
	pulumi.RegisterOutputType(WebhookReceiverOutput{})
	pulumi.RegisterOutputType(WebhookReceiverArrayOutput{})
	pulumi.RegisterOutputType(WebhookReceiverResponseOutput{})
	pulumi.RegisterOutputType(WebhookReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceArrayOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceResponseOutput{})
	pulumi.RegisterOutputType(WindowsEventLogDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryResponseOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryResponseOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkbookUserAssignedIdentitiesResponseOutput{})
	pulumi.RegisterOutputType(WorkbookUserAssignedIdentitiesResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryArrayMapOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryResponseArrayMapOutput{})
}
