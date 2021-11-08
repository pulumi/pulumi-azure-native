


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





type ActionGroupResponseInput interface {
	pulumi.Input

	ToActionGroupResponseOutput() ActionGroupResponseOutput
	ToActionGroupResponseOutputWithContext(context.Context) ActionGroupResponseOutput
}

type ActionGroupResponseArgs struct {
	ActionGroupId     pulumi.StringInput    `pulumi:"actionGroupId"`
	WebhookProperties pulumi.StringMapInput `pulumi:"webhookProperties"`
}

func (ActionGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupResponse)(nil)).Elem()
}

func (i ActionGroupResponseArgs) ToActionGroupResponseOutput() ActionGroupResponseOutput {
	return i.ToActionGroupResponseOutputWithContext(context.Background())
}

func (i ActionGroupResponseArgs) ToActionGroupResponseOutputWithContext(ctx context.Context) ActionGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupResponseOutput)
}





type ActionGroupResponseArrayInput interface {
	pulumi.Input

	ToActionGroupResponseArrayOutput() ActionGroupResponseArrayOutput
	ToActionGroupResponseArrayOutputWithContext(context.Context) ActionGroupResponseArrayOutput
}

type ActionGroupResponseArray []ActionGroupResponseInput

func (ActionGroupResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionGroupResponse)(nil)).Elem()
}

func (i ActionGroupResponseArray) ToActionGroupResponseArrayOutput() ActionGroupResponseArrayOutput {
	return i.ToActionGroupResponseArrayOutputWithContext(context.Background())
}

func (i ActionGroupResponseArray) ToActionGroupResponseArrayOutputWithContext(ctx context.Context) ActionGroupResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupResponseArrayOutput)
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

func (i ActionListArgs) ToActionListPtrOutput() ActionListPtrOutput {
	return i.ToActionListPtrOutputWithContext(context.Background())
}

func (i ActionListArgs) ToActionListPtrOutputWithContext(ctx context.Context) ActionListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionListOutput).ToActionListPtrOutputWithContext(ctx)
}









type ActionListPtrInput interface {
	pulumi.Input

	ToActionListPtrOutput() ActionListPtrOutput
	ToActionListPtrOutputWithContext(context.Context) ActionListPtrOutput
}

type actionListPtrType ActionListArgs

func ActionListPtr(v *ActionListArgs) ActionListPtrInput {
	return (*actionListPtrType)(v)
}

func (*actionListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionList)(nil)).Elem()
}

func (i *actionListPtrType) ToActionListPtrOutput() ActionListPtrOutput {
	return i.ToActionListPtrOutputWithContext(context.Background())
}

func (i *actionListPtrType) ToActionListPtrOutputWithContext(ctx context.Context) ActionListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionListPtrOutput)
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

func (o ActionListOutput) ToActionListPtrOutput() ActionListPtrOutput {
	return o.ToActionListPtrOutputWithContext(context.Background())
}

func (o ActionListOutput) ToActionListPtrOutputWithContext(ctx context.Context) ActionListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionList) *ActionList {
		return &v
	}).(ActionListPtrOutput)
}

func (o ActionListOutput) ActionGroups() ActionGroupTypeArrayOutput {
	return o.ApplyT(func(v ActionList) []ActionGroupType { return v.ActionGroups }).(ActionGroupTypeArrayOutput)
}

type ActionListPtrOutput struct{ *pulumi.OutputState }

func (ActionListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionList)(nil)).Elem()
}

func (o ActionListPtrOutput) ToActionListPtrOutput() ActionListPtrOutput {
	return o
}

func (o ActionListPtrOutput) ToActionListPtrOutputWithContext(ctx context.Context) ActionListPtrOutput {
	return o
}

func (o ActionListPtrOutput) Elem() ActionListOutput {
	return o.ApplyT(func(v *ActionList) ActionList {
		if v != nil {
			return *v
		}
		var ret ActionList
		return ret
	}).(ActionListOutput)
}

func (o ActionListPtrOutput) ActionGroups() ActionGroupTypeArrayOutput {
	return o.ApplyT(func(v *ActionList) []ActionGroupType {
		if v == nil {
			return nil
		}
		return v.ActionGroups
	}).(ActionGroupTypeArrayOutput)
}

type ActionListResponse struct {
	ActionGroups []ActionGroupResponse `pulumi:"actionGroups"`
}





type ActionListResponseInput interface {
	pulumi.Input

	ToActionListResponseOutput() ActionListResponseOutput
	ToActionListResponseOutputWithContext(context.Context) ActionListResponseOutput
}

type ActionListResponseArgs struct {
	ActionGroups ActionGroupResponseArrayInput `pulumi:"actionGroups"`
}

func (ActionListResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionListResponse)(nil)).Elem()
}

func (i ActionListResponseArgs) ToActionListResponseOutput() ActionListResponseOutput {
	return i.ToActionListResponseOutputWithContext(context.Background())
}

func (i ActionListResponseArgs) ToActionListResponseOutputWithContext(ctx context.Context) ActionListResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionListResponseOutput)
}

func (i ActionListResponseArgs) ToActionListResponsePtrOutput() ActionListResponsePtrOutput {
	return i.ToActionListResponsePtrOutputWithContext(context.Background())
}

func (i ActionListResponseArgs) ToActionListResponsePtrOutputWithContext(ctx context.Context) ActionListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionListResponseOutput).ToActionListResponsePtrOutputWithContext(ctx)
}









type ActionListResponsePtrInput interface {
	pulumi.Input

	ToActionListResponsePtrOutput() ActionListResponsePtrOutput
	ToActionListResponsePtrOutputWithContext(context.Context) ActionListResponsePtrOutput
}

type actionListResponsePtrType ActionListResponseArgs

func ActionListResponsePtr(v *ActionListResponseArgs) ActionListResponsePtrInput {
	return (*actionListResponsePtrType)(v)
}

func (*actionListResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionListResponse)(nil)).Elem()
}

func (i *actionListResponsePtrType) ToActionListResponsePtrOutput() ActionListResponsePtrOutput {
	return i.ToActionListResponsePtrOutputWithContext(context.Background())
}

func (i *actionListResponsePtrType) ToActionListResponsePtrOutputWithContext(ctx context.Context) ActionListResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionListResponsePtrOutput)
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

func (o ActionListResponseOutput) ToActionListResponsePtrOutput() ActionListResponsePtrOutput {
	return o.ToActionListResponsePtrOutputWithContext(context.Background())
}

func (o ActionListResponseOutput) ToActionListResponsePtrOutputWithContext(ctx context.Context) ActionListResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionListResponse) *ActionListResponse {
		return &v
	}).(ActionListResponsePtrOutput)
}

func (o ActionListResponseOutput) ActionGroups() ActionGroupResponseArrayOutput {
	return o.ApplyT(func(v ActionListResponse) []ActionGroupResponse { return v.ActionGroups }).(ActionGroupResponseArrayOutput)
}

type ActionListResponsePtrOutput struct{ *pulumi.OutputState }

func (ActionListResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionListResponse)(nil)).Elem()
}

func (o ActionListResponsePtrOutput) ToActionListResponsePtrOutput() ActionListResponsePtrOutput {
	return o
}

func (o ActionListResponsePtrOutput) ToActionListResponsePtrOutputWithContext(ctx context.Context) ActionListResponsePtrOutput {
	return o
}

func (o ActionListResponsePtrOutput) Elem() ActionListResponseOutput {
	return o.ApplyT(func(v *ActionListResponse) ActionListResponse {
		if v != nil {
			return *v
		}
		var ret ActionListResponse
		return ret
	}).(ActionListResponseOutput)
}

func (o ActionListResponsePtrOutput) ActionGroups() ActionGroupResponseArrayOutput {
	return o.ApplyT(func(v *ActionListResponse) []ActionGroupResponse {
		if v == nil {
			return nil
		}
		return v.ActionGroups
	}).(ActionGroupResponseArrayOutput)
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

func (i AlertRuleAllOfConditionArgs) ToAlertRuleAllOfConditionPtrOutput() AlertRuleAllOfConditionPtrOutput {
	return i.ToAlertRuleAllOfConditionPtrOutputWithContext(context.Background())
}

func (i AlertRuleAllOfConditionArgs) ToAlertRuleAllOfConditionPtrOutputWithContext(ctx context.Context) AlertRuleAllOfConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAllOfConditionOutput).ToAlertRuleAllOfConditionPtrOutputWithContext(ctx)
}









type AlertRuleAllOfConditionPtrInput interface {
	pulumi.Input

	ToAlertRuleAllOfConditionPtrOutput() AlertRuleAllOfConditionPtrOutput
	ToAlertRuleAllOfConditionPtrOutputWithContext(context.Context) AlertRuleAllOfConditionPtrOutput
}

type alertRuleAllOfConditionPtrType AlertRuleAllOfConditionArgs

func AlertRuleAllOfConditionPtr(v *AlertRuleAllOfConditionArgs) AlertRuleAllOfConditionPtrInput {
	return (*alertRuleAllOfConditionPtrType)(v)
}

func (*alertRuleAllOfConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleAllOfCondition)(nil)).Elem()
}

func (i *alertRuleAllOfConditionPtrType) ToAlertRuleAllOfConditionPtrOutput() AlertRuleAllOfConditionPtrOutput {
	return i.ToAlertRuleAllOfConditionPtrOutputWithContext(context.Background())
}

func (i *alertRuleAllOfConditionPtrType) ToAlertRuleAllOfConditionPtrOutputWithContext(ctx context.Context) AlertRuleAllOfConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAllOfConditionPtrOutput)
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

func (o AlertRuleAllOfConditionOutput) ToAlertRuleAllOfConditionPtrOutput() AlertRuleAllOfConditionPtrOutput {
	return o.ToAlertRuleAllOfConditionPtrOutputWithContext(context.Background())
}

func (o AlertRuleAllOfConditionOutput) ToAlertRuleAllOfConditionPtrOutputWithContext(ctx context.Context) AlertRuleAllOfConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertRuleAllOfCondition) *AlertRuleAllOfCondition {
		return &v
	}).(AlertRuleAllOfConditionPtrOutput)
}

func (o AlertRuleAllOfConditionOutput) AllOf() AlertRuleAnyOfOrLeafConditionArrayOutput {
	return o.ApplyT(func(v AlertRuleAllOfCondition) []AlertRuleAnyOfOrLeafCondition { return v.AllOf }).(AlertRuleAnyOfOrLeafConditionArrayOutput)
}

type AlertRuleAllOfConditionPtrOutput struct{ *pulumi.OutputState }

func (AlertRuleAllOfConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleAllOfCondition)(nil)).Elem()
}

func (o AlertRuleAllOfConditionPtrOutput) ToAlertRuleAllOfConditionPtrOutput() AlertRuleAllOfConditionPtrOutput {
	return o
}

func (o AlertRuleAllOfConditionPtrOutput) ToAlertRuleAllOfConditionPtrOutputWithContext(ctx context.Context) AlertRuleAllOfConditionPtrOutput {
	return o
}

func (o AlertRuleAllOfConditionPtrOutput) Elem() AlertRuleAllOfConditionOutput {
	return o.ApplyT(func(v *AlertRuleAllOfCondition) AlertRuleAllOfCondition {
		if v != nil {
			return *v
		}
		var ret AlertRuleAllOfCondition
		return ret
	}).(AlertRuleAllOfConditionOutput)
}

func (o AlertRuleAllOfConditionPtrOutput) AllOf() AlertRuleAnyOfOrLeafConditionArrayOutput {
	return o.ApplyT(func(v *AlertRuleAllOfCondition) []AlertRuleAnyOfOrLeafCondition {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(AlertRuleAnyOfOrLeafConditionArrayOutput)
}

type AlertRuleAllOfConditionResponse struct {
	AllOf []AlertRuleAnyOfOrLeafConditionResponse `pulumi:"allOf"`
}





type AlertRuleAllOfConditionResponseInput interface {
	pulumi.Input

	ToAlertRuleAllOfConditionResponseOutput() AlertRuleAllOfConditionResponseOutput
	ToAlertRuleAllOfConditionResponseOutputWithContext(context.Context) AlertRuleAllOfConditionResponseOutput
}

type AlertRuleAllOfConditionResponseArgs struct {
	AllOf AlertRuleAnyOfOrLeafConditionResponseArrayInput `pulumi:"allOf"`
}

func (AlertRuleAllOfConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleAllOfConditionResponse)(nil)).Elem()
}

func (i AlertRuleAllOfConditionResponseArgs) ToAlertRuleAllOfConditionResponseOutput() AlertRuleAllOfConditionResponseOutput {
	return i.ToAlertRuleAllOfConditionResponseOutputWithContext(context.Background())
}

func (i AlertRuleAllOfConditionResponseArgs) ToAlertRuleAllOfConditionResponseOutputWithContext(ctx context.Context) AlertRuleAllOfConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAllOfConditionResponseOutput)
}

func (i AlertRuleAllOfConditionResponseArgs) ToAlertRuleAllOfConditionResponsePtrOutput() AlertRuleAllOfConditionResponsePtrOutput {
	return i.ToAlertRuleAllOfConditionResponsePtrOutputWithContext(context.Background())
}

func (i AlertRuleAllOfConditionResponseArgs) ToAlertRuleAllOfConditionResponsePtrOutputWithContext(ctx context.Context) AlertRuleAllOfConditionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAllOfConditionResponseOutput).ToAlertRuleAllOfConditionResponsePtrOutputWithContext(ctx)
}









type AlertRuleAllOfConditionResponsePtrInput interface {
	pulumi.Input

	ToAlertRuleAllOfConditionResponsePtrOutput() AlertRuleAllOfConditionResponsePtrOutput
	ToAlertRuleAllOfConditionResponsePtrOutputWithContext(context.Context) AlertRuleAllOfConditionResponsePtrOutput
}

type alertRuleAllOfConditionResponsePtrType AlertRuleAllOfConditionResponseArgs

func AlertRuleAllOfConditionResponsePtr(v *AlertRuleAllOfConditionResponseArgs) AlertRuleAllOfConditionResponsePtrInput {
	return (*alertRuleAllOfConditionResponsePtrType)(v)
}

func (*alertRuleAllOfConditionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleAllOfConditionResponse)(nil)).Elem()
}

func (i *alertRuleAllOfConditionResponsePtrType) ToAlertRuleAllOfConditionResponsePtrOutput() AlertRuleAllOfConditionResponsePtrOutput {
	return i.ToAlertRuleAllOfConditionResponsePtrOutputWithContext(context.Background())
}

func (i *alertRuleAllOfConditionResponsePtrType) ToAlertRuleAllOfConditionResponsePtrOutputWithContext(ctx context.Context) AlertRuleAllOfConditionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAllOfConditionResponsePtrOutput)
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

func (o AlertRuleAllOfConditionResponseOutput) ToAlertRuleAllOfConditionResponsePtrOutput() AlertRuleAllOfConditionResponsePtrOutput {
	return o.ToAlertRuleAllOfConditionResponsePtrOutputWithContext(context.Background())
}

func (o AlertRuleAllOfConditionResponseOutput) ToAlertRuleAllOfConditionResponsePtrOutputWithContext(ctx context.Context) AlertRuleAllOfConditionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertRuleAllOfConditionResponse) *AlertRuleAllOfConditionResponse {
		return &v
	}).(AlertRuleAllOfConditionResponsePtrOutput)
}

func (o AlertRuleAllOfConditionResponseOutput) AllOf() AlertRuleAnyOfOrLeafConditionResponseArrayOutput {
	return o.ApplyT(func(v AlertRuleAllOfConditionResponse) []AlertRuleAnyOfOrLeafConditionResponse { return v.AllOf }).(AlertRuleAnyOfOrLeafConditionResponseArrayOutput)
}

type AlertRuleAllOfConditionResponsePtrOutput struct{ *pulumi.OutputState }

func (AlertRuleAllOfConditionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleAllOfConditionResponse)(nil)).Elem()
}

func (o AlertRuleAllOfConditionResponsePtrOutput) ToAlertRuleAllOfConditionResponsePtrOutput() AlertRuleAllOfConditionResponsePtrOutput {
	return o
}

func (o AlertRuleAllOfConditionResponsePtrOutput) ToAlertRuleAllOfConditionResponsePtrOutputWithContext(ctx context.Context) AlertRuleAllOfConditionResponsePtrOutput {
	return o
}

func (o AlertRuleAllOfConditionResponsePtrOutput) Elem() AlertRuleAllOfConditionResponseOutput {
	return o.ApplyT(func(v *AlertRuleAllOfConditionResponse) AlertRuleAllOfConditionResponse {
		if v != nil {
			return *v
		}
		var ret AlertRuleAllOfConditionResponse
		return ret
	}).(AlertRuleAllOfConditionResponseOutput)
}

func (o AlertRuleAllOfConditionResponsePtrOutput) AllOf() AlertRuleAnyOfOrLeafConditionResponseArrayOutput {
	return o.ApplyT(func(v *AlertRuleAllOfConditionResponse) []AlertRuleAnyOfOrLeafConditionResponse {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(AlertRuleAnyOfOrLeafConditionResponseArrayOutput)
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





type AlertRuleAnyOfOrLeafConditionResponseInput interface {
	pulumi.Input

	ToAlertRuleAnyOfOrLeafConditionResponseOutput() AlertRuleAnyOfOrLeafConditionResponseOutput
	ToAlertRuleAnyOfOrLeafConditionResponseOutputWithContext(context.Context) AlertRuleAnyOfOrLeafConditionResponseOutput
}

type AlertRuleAnyOfOrLeafConditionResponseArgs struct {
	AnyOf       AlertRuleLeafConditionResponseArrayInput `pulumi:"anyOf"`
	ContainsAny pulumi.StringArrayInput                  `pulumi:"containsAny"`
	Equals      pulumi.StringPtrInput                    `pulumi:"equals"`
	Field       pulumi.StringPtrInput                    `pulumi:"field"`
}

func (AlertRuleAnyOfOrLeafConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleAnyOfOrLeafConditionResponse)(nil)).Elem()
}

func (i AlertRuleAnyOfOrLeafConditionResponseArgs) ToAlertRuleAnyOfOrLeafConditionResponseOutput() AlertRuleAnyOfOrLeafConditionResponseOutput {
	return i.ToAlertRuleAnyOfOrLeafConditionResponseOutputWithContext(context.Background())
}

func (i AlertRuleAnyOfOrLeafConditionResponseArgs) ToAlertRuleAnyOfOrLeafConditionResponseOutputWithContext(ctx context.Context) AlertRuleAnyOfOrLeafConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAnyOfOrLeafConditionResponseOutput)
}





type AlertRuleAnyOfOrLeafConditionResponseArrayInput interface {
	pulumi.Input

	ToAlertRuleAnyOfOrLeafConditionResponseArrayOutput() AlertRuleAnyOfOrLeafConditionResponseArrayOutput
	ToAlertRuleAnyOfOrLeafConditionResponseArrayOutputWithContext(context.Context) AlertRuleAnyOfOrLeafConditionResponseArrayOutput
}

type AlertRuleAnyOfOrLeafConditionResponseArray []AlertRuleAnyOfOrLeafConditionResponseInput

func (AlertRuleAnyOfOrLeafConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertRuleAnyOfOrLeafConditionResponse)(nil)).Elem()
}

func (i AlertRuleAnyOfOrLeafConditionResponseArray) ToAlertRuleAnyOfOrLeafConditionResponseArrayOutput() AlertRuleAnyOfOrLeafConditionResponseArrayOutput {
	return i.ToAlertRuleAnyOfOrLeafConditionResponseArrayOutputWithContext(context.Background())
}

func (i AlertRuleAnyOfOrLeafConditionResponseArray) ToAlertRuleAnyOfOrLeafConditionResponseArrayOutputWithContext(ctx context.Context) AlertRuleAnyOfOrLeafConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleAnyOfOrLeafConditionResponseArrayOutput)
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





type AlertRuleLeafConditionResponseInput interface {
	pulumi.Input

	ToAlertRuleLeafConditionResponseOutput() AlertRuleLeafConditionResponseOutput
	ToAlertRuleLeafConditionResponseOutputWithContext(context.Context) AlertRuleLeafConditionResponseOutput
}

type AlertRuleLeafConditionResponseArgs struct {
	ContainsAny pulumi.StringArrayInput `pulumi:"containsAny"`
	Equals      pulumi.StringPtrInput   `pulumi:"equals"`
	Field       pulumi.StringPtrInput   `pulumi:"field"`
}

func (AlertRuleLeafConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleLeafConditionResponse)(nil)).Elem()
}

func (i AlertRuleLeafConditionResponseArgs) ToAlertRuleLeafConditionResponseOutput() AlertRuleLeafConditionResponseOutput {
	return i.ToAlertRuleLeafConditionResponseOutputWithContext(context.Background())
}

func (i AlertRuleLeafConditionResponseArgs) ToAlertRuleLeafConditionResponseOutputWithContext(ctx context.Context) AlertRuleLeafConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleLeafConditionResponseOutput)
}





type AlertRuleLeafConditionResponseArrayInput interface {
	pulumi.Input

	ToAlertRuleLeafConditionResponseArrayOutput() AlertRuleLeafConditionResponseArrayOutput
	ToAlertRuleLeafConditionResponseArrayOutputWithContext(context.Context) AlertRuleLeafConditionResponseArrayOutput
}

type AlertRuleLeafConditionResponseArray []AlertRuleLeafConditionResponseInput

func (AlertRuleLeafConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertRuleLeafConditionResponse)(nil)).Elem()
}

func (i AlertRuleLeafConditionResponseArray) ToAlertRuleLeafConditionResponseArrayOutput() AlertRuleLeafConditionResponseArrayOutput {
	return i.ToAlertRuleLeafConditionResponseArrayOutputWithContext(context.Background())
}

func (i AlertRuleLeafConditionResponseArray) ToAlertRuleLeafConditionResponseArrayOutputWithContext(ctx context.Context) AlertRuleLeafConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleLeafConditionResponseArrayOutput)
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





type AlertingActionInput interface {
	pulumi.Input

	ToAlertingActionOutput() AlertingActionOutput
	ToAlertingActionOutputWithContext(context.Context) AlertingActionOutput
}

type AlertingActionArgs struct {
	AznsAction      AzNsActionGroupPtrInput `pulumi:"aznsAction"`
	OdataType       pulumi.StringInput      `pulumi:"odataType"`
	Severity        pulumi.StringInput      `pulumi:"severity"`
	ThrottlingInMin pulumi.IntPtrInput      `pulumi:"throttlingInMin"`
	Trigger         TriggerConditionInput   `pulumi:"trigger"`
}

func (AlertingActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertingAction)(nil)).Elem()
}

func (i AlertingActionArgs) ToAlertingActionOutput() AlertingActionOutput {
	return i.ToAlertingActionOutputWithContext(context.Background())
}

func (i AlertingActionArgs) ToAlertingActionOutputWithContext(ctx context.Context) AlertingActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertingActionOutput)
}

type AlertingActionOutput struct{ *pulumi.OutputState }

func (AlertingActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertingAction)(nil)).Elem()
}

func (o AlertingActionOutput) ToAlertingActionOutput() AlertingActionOutput {
	return o
}

func (o AlertingActionOutput) ToAlertingActionOutputWithContext(ctx context.Context) AlertingActionOutput {
	return o
}

func (o AlertingActionOutput) AznsAction() AzNsActionGroupPtrOutput {
	return o.ApplyT(func(v AlertingAction) *AzNsActionGroup { return v.AznsAction }).(AzNsActionGroupPtrOutput)
}

func (o AlertingActionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AlertingAction) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AlertingActionOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v AlertingAction) string { return v.Severity }).(pulumi.StringOutput)
}

func (o AlertingActionOutput) ThrottlingInMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AlertingAction) *int { return v.ThrottlingInMin }).(pulumi.IntPtrOutput)
}

func (o AlertingActionOutput) Trigger() TriggerConditionOutput {
	return o.ApplyT(func(v AlertingAction) TriggerCondition { return v.Trigger }).(TriggerConditionOutput)
}

type AlertingActionResponse struct {
	AznsAction      *AzNsActionGroupResponse `pulumi:"aznsAction"`
	OdataType       string                   `pulumi:"odataType"`
	Severity        string                   `pulumi:"severity"`
	ThrottlingInMin *int                     `pulumi:"throttlingInMin"`
	Trigger         TriggerConditionResponse `pulumi:"trigger"`
}





type AlertingActionResponseInput interface {
	pulumi.Input

	ToAlertingActionResponseOutput() AlertingActionResponseOutput
	ToAlertingActionResponseOutputWithContext(context.Context) AlertingActionResponseOutput
}

type AlertingActionResponseArgs struct {
	AznsAction      AzNsActionGroupResponsePtrInput `pulumi:"aznsAction"`
	OdataType       pulumi.StringInput              `pulumi:"odataType"`
	Severity        pulumi.StringInput              `pulumi:"severity"`
	ThrottlingInMin pulumi.IntPtrInput              `pulumi:"throttlingInMin"`
	Trigger         TriggerConditionResponseInput   `pulumi:"trigger"`
}

func (AlertingActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertingActionResponse)(nil)).Elem()
}

func (i AlertingActionResponseArgs) ToAlertingActionResponseOutput() AlertingActionResponseOutput {
	return i.ToAlertingActionResponseOutputWithContext(context.Background())
}

func (i AlertingActionResponseArgs) ToAlertingActionResponseOutputWithContext(ctx context.Context) AlertingActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertingActionResponseOutput)
}

type AlertingActionResponseOutput struct{ *pulumi.OutputState }

func (AlertingActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertingActionResponse)(nil)).Elem()
}

func (o AlertingActionResponseOutput) ToAlertingActionResponseOutput() AlertingActionResponseOutput {
	return o
}

func (o AlertingActionResponseOutput) ToAlertingActionResponseOutputWithContext(ctx context.Context) AlertingActionResponseOutput {
	return o
}

func (o AlertingActionResponseOutput) AznsAction() AzNsActionGroupResponsePtrOutput {
	return o.ApplyT(func(v AlertingActionResponse) *AzNsActionGroupResponse { return v.AznsAction }).(AzNsActionGroupResponsePtrOutput)
}

func (o AlertingActionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v AlertingActionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o AlertingActionResponseOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v AlertingActionResponse) string { return v.Severity }).(pulumi.StringOutput)
}

func (o AlertingActionResponseOutput) ThrottlingInMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AlertingActionResponse) *int { return v.ThrottlingInMin }).(pulumi.IntPtrOutput)
}

func (o AlertingActionResponseOutput) Trigger() TriggerConditionResponseOutput {
	return o.ApplyT(func(v AlertingActionResponse) TriggerConditionResponse { return v.Trigger }).(TriggerConditionResponseOutput)
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





type ApplicationInsightsComponentAnalyticsItemPropertiesResponseInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput
}

type ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs struct {
	FunctionAlias pulumi.StringPtrInput `pulumi:"functionAlias"`
}

func (ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesResponseOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput)
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput).ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput
	ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput
}

type applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs

func ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtr(v *ApplicationInsightsComponentAnalyticsItemPropertiesResponseArgs) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrInput {
	return (*applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType)(v)
}

func (*applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return i.ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentAnalyticsItemPropertiesResponsePtrType) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput)
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

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o.ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *ApplicationInsightsComponentAnalyticsItemPropertiesResponse {
		return &v
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput)
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *string { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

type ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentAnalyticsItemPropertiesResponse)(nil)).Elem()
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput() ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) ToApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput {
	return o
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) Elem() ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemPropertiesResponse) ApplicationInsightsComponentAnalyticsItemPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentAnalyticsItemPropertiesResponse
		return ret
	}).(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput)
}

func (o ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentAnalyticsItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FunctionAlias
	}).(pulumi.StringPtrOutput)
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





type ApplicationInsightsComponentDataVolumeCapResponseInput interface {
	pulumi.Input

	ToApplicationInsightsComponentDataVolumeCapResponseOutput() ApplicationInsightsComponentDataVolumeCapResponseOutput
	ToApplicationInsightsComponentDataVolumeCapResponseOutputWithContext(context.Context) ApplicationInsightsComponentDataVolumeCapResponseOutput
}

type ApplicationInsightsComponentDataVolumeCapResponseArgs struct {
	Cap                                  pulumi.Float64PtrInput `pulumi:"cap"`
	MaxHistoryCap                        pulumi.Float64Input    `pulumi:"maxHistoryCap"`
	ResetTime                            pulumi.IntInput        `pulumi:"resetTime"`
	StopSendNotificationWhenHitCap       pulumi.BoolPtrInput    `pulumi:"stopSendNotificationWhenHitCap"`
	StopSendNotificationWhenHitThreshold pulumi.BoolPtrInput    `pulumi:"stopSendNotificationWhenHitThreshold"`
	WarningThreshold                     pulumi.IntPtrInput     `pulumi:"warningThreshold"`
}

func (ApplicationInsightsComponentDataVolumeCapResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentDataVolumeCapResponse)(nil)).Elem()
}

func (i ApplicationInsightsComponentDataVolumeCapResponseArgs) ToApplicationInsightsComponentDataVolumeCapResponseOutput() ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapResponseOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentDataVolumeCapResponseArgs) ToApplicationInsightsComponentDataVolumeCapResponseOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapResponseOutput)
}

func (i ApplicationInsightsComponentDataVolumeCapResponseArgs) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentDataVolumeCapResponseArgs) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapResponseOutput).ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentDataVolumeCapResponsePtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput
	ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput
}

type applicationInsightsComponentDataVolumeCapResponsePtrType ApplicationInsightsComponentDataVolumeCapResponseArgs

func ApplicationInsightsComponentDataVolumeCapResponsePtr(v *ApplicationInsightsComponentDataVolumeCapResponseArgs) ApplicationInsightsComponentDataVolumeCapResponsePtrInput {
	return (*applicationInsightsComponentDataVolumeCapResponsePtrType)(v)
}

func (*applicationInsightsComponentDataVolumeCapResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentDataVolumeCapResponse)(nil)).Elem()
}

func (i *applicationInsightsComponentDataVolumeCapResponsePtrType) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return i.ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentDataVolumeCapResponsePtrType) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentDataVolumeCapResponsePtrOutput)
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

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutput() ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return o.ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentDataVolumeCapResponseOutput) ToApplicationInsightsComponentDataVolumeCapResponsePtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentDataVolumeCapResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentDataVolumeCapResponse) *ApplicationInsightsComponentDataVolumeCapResponse {
		return &v
	}).(ApplicationInsightsComponentDataVolumeCapResponsePtrOutput)
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





type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput
}

type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs struct {
	Description                pulumi.StringPtrInput `pulumi:"description"`
	DisplayName                pulumi.StringPtrInput `pulumi:"displayName"`
	HelpUrl                    pulumi.StringPtrInput `pulumi:"helpUrl"`
	IsEnabledByDefault         pulumi.BoolPtrInput   `pulumi:"isEnabledByDefault"`
	IsHidden                   pulumi.BoolPtrInput   `pulumi:"isHidden"`
	IsInPreview                pulumi.BoolPtrInput   `pulumi:"isInPreview"`
	Name                       pulumi.StringPtrInput `pulumi:"name"`
	SupportsEmailNotifications pulumi.BoolPtrInput   `pulumi:"supportsEmailNotifications"`
}

func (ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions)(nil)).Elem()
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput)
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput).ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx)
}









type ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput
}

type applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs

func ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtr(v *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsArgs) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrInput {
	return (*applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType)(v)
}

func (*applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions)(nil)).Elem()
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput)
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

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return o.ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions) *ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitions {
		return &v
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationResponseRuleDefinitionsPtrOutput)
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





type ArmRoleReceiverResponseInput interface {
	pulumi.Input

	ToArmRoleReceiverResponseOutput() ArmRoleReceiverResponseOutput
	ToArmRoleReceiverResponseOutputWithContext(context.Context) ArmRoleReceiverResponseOutput
}

type ArmRoleReceiverResponseArgs struct {
	Name                 pulumi.StringInput  `pulumi:"name"`
	RoleId               pulumi.StringInput  `pulumi:"roleId"`
	UseCommonAlertSchema pulumi.BoolPtrInput `pulumi:"useCommonAlertSchema"`
}

func (ArmRoleReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmRoleReceiverResponse)(nil)).Elem()
}

func (i ArmRoleReceiverResponseArgs) ToArmRoleReceiverResponseOutput() ArmRoleReceiverResponseOutput {
	return i.ToArmRoleReceiverResponseOutputWithContext(context.Background())
}

func (i ArmRoleReceiverResponseArgs) ToArmRoleReceiverResponseOutputWithContext(ctx context.Context) ArmRoleReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRoleReceiverResponseOutput)
}





type ArmRoleReceiverResponseArrayInput interface {
	pulumi.Input

	ToArmRoleReceiverResponseArrayOutput() ArmRoleReceiverResponseArrayOutput
	ToArmRoleReceiverResponseArrayOutputWithContext(context.Context) ArmRoleReceiverResponseArrayOutput
}

type ArmRoleReceiverResponseArray []ArmRoleReceiverResponseInput

func (ArmRoleReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmRoleReceiverResponse)(nil)).Elem()
}

func (i ArmRoleReceiverResponseArray) ToArmRoleReceiverResponseArrayOutput() ArmRoleReceiverResponseArrayOutput {
	return i.ToArmRoleReceiverResponseArrayOutputWithContext(context.Background())
}

func (i ArmRoleReceiverResponseArray) ToArmRoleReceiverResponseArrayOutputWithContext(ctx context.Context) ArmRoleReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmRoleReceiverResponseArrayOutput)
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





type AutomationRunbookReceiverResponseInput interface {
	pulumi.Input

	ToAutomationRunbookReceiverResponseOutput() AutomationRunbookReceiverResponseOutput
	ToAutomationRunbookReceiverResponseOutputWithContext(context.Context) AutomationRunbookReceiverResponseOutput
}

type AutomationRunbookReceiverResponseArgs struct {
	AutomationAccountId  pulumi.StringInput    `pulumi:"automationAccountId"`
	IsGlobalRunbook      pulumi.BoolInput      `pulumi:"isGlobalRunbook"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
	RunbookName          pulumi.StringInput    `pulumi:"runbookName"`
	ServiceUri           pulumi.StringPtrInput `pulumi:"serviceUri"`
	UseCommonAlertSchema pulumi.BoolPtrInput   `pulumi:"useCommonAlertSchema"`
	WebhookResourceId    pulumi.StringInput    `pulumi:"webhookResourceId"`
}

func (AutomationRunbookReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRunbookReceiverResponse)(nil)).Elem()
}

func (i AutomationRunbookReceiverResponseArgs) ToAutomationRunbookReceiverResponseOutput() AutomationRunbookReceiverResponseOutput {
	return i.ToAutomationRunbookReceiverResponseOutputWithContext(context.Background())
}

func (i AutomationRunbookReceiverResponseArgs) ToAutomationRunbookReceiverResponseOutputWithContext(ctx context.Context) AutomationRunbookReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRunbookReceiverResponseOutput)
}





type AutomationRunbookReceiverResponseArrayInput interface {
	pulumi.Input

	ToAutomationRunbookReceiverResponseArrayOutput() AutomationRunbookReceiverResponseArrayOutput
	ToAutomationRunbookReceiverResponseArrayOutputWithContext(context.Context) AutomationRunbookReceiverResponseArrayOutput
}

type AutomationRunbookReceiverResponseArray []AutomationRunbookReceiverResponseInput

func (AutomationRunbookReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRunbookReceiverResponse)(nil)).Elem()
}

func (i AutomationRunbookReceiverResponseArray) ToAutomationRunbookReceiverResponseArrayOutput() AutomationRunbookReceiverResponseArrayOutput {
	return i.ToAutomationRunbookReceiverResponseArrayOutputWithContext(context.Background())
}

func (i AutomationRunbookReceiverResponseArray) ToAutomationRunbookReceiverResponseArrayOutputWithContext(ctx context.Context) AutomationRunbookReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRunbookReceiverResponseArrayOutput)
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





type AutoscaleNotificationResponseInput interface {
	pulumi.Input

	ToAutoscaleNotificationResponseOutput() AutoscaleNotificationResponseOutput
	ToAutoscaleNotificationResponseOutputWithContext(context.Context) AutoscaleNotificationResponseOutput
}

type AutoscaleNotificationResponseArgs struct {
	Email     EmailNotificationResponsePtrInput     `pulumi:"email"`
	Operation pulumi.StringInput                    `pulumi:"operation"`
	Webhooks  WebhookNotificationResponseArrayInput `pulumi:"webhooks"`
}

func (AutoscaleNotificationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleNotificationResponse)(nil)).Elem()
}

func (i AutoscaleNotificationResponseArgs) ToAutoscaleNotificationResponseOutput() AutoscaleNotificationResponseOutput {
	return i.ToAutoscaleNotificationResponseOutputWithContext(context.Background())
}

func (i AutoscaleNotificationResponseArgs) ToAutoscaleNotificationResponseOutputWithContext(ctx context.Context) AutoscaleNotificationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleNotificationResponseOutput)
}





type AutoscaleNotificationResponseArrayInput interface {
	pulumi.Input

	ToAutoscaleNotificationResponseArrayOutput() AutoscaleNotificationResponseArrayOutput
	ToAutoscaleNotificationResponseArrayOutputWithContext(context.Context) AutoscaleNotificationResponseArrayOutput
}

type AutoscaleNotificationResponseArray []AutoscaleNotificationResponseInput

func (AutoscaleNotificationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleNotificationResponse)(nil)).Elem()
}

func (i AutoscaleNotificationResponseArray) ToAutoscaleNotificationResponseArrayOutput() AutoscaleNotificationResponseArrayOutput {
	return i.ToAutoscaleNotificationResponseArrayOutputWithContext(context.Background())
}

func (i AutoscaleNotificationResponseArray) ToAutoscaleNotificationResponseArrayOutputWithContext(ctx context.Context) AutoscaleNotificationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleNotificationResponseArrayOutput)
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





type AutoscaleProfileResponseInput interface {
	pulumi.Input

	ToAutoscaleProfileResponseOutput() AutoscaleProfileResponseOutput
	ToAutoscaleProfileResponseOutputWithContext(context.Context) AutoscaleProfileResponseOutput
}

type AutoscaleProfileResponseArgs struct {
	Capacity   ScaleCapacityResponseInput  `pulumi:"capacity"`
	FixedDate  TimeWindowResponsePtrInput  `pulumi:"fixedDate"`
	Name       pulumi.StringInput          `pulumi:"name"`
	Recurrence RecurrenceResponsePtrInput  `pulumi:"recurrence"`
	Rules      ScaleRuleResponseArrayInput `pulumi:"rules"`
}

func (AutoscaleProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleProfileResponse)(nil)).Elem()
}

func (i AutoscaleProfileResponseArgs) ToAutoscaleProfileResponseOutput() AutoscaleProfileResponseOutput {
	return i.ToAutoscaleProfileResponseOutputWithContext(context.Background())
}

func (i AutoscaleProfileResponseArgs) ToAutoscaleProfileResponseOutputWithContext(ctx context.Context) AutoscaleProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleProfileResponseOutput)
}





type AutoscaleProfileResponseArrayInput interface {
	pulumi.Input

	ToAutoscaleProfileResponseArrayOutput() AutoscaleProfileResponseArrayOutput
	ToAutoscaleProfileResponseArrayOutputWithContext(context.Context) AutoscaleProfileResponseArrayOutput
}

type AutoscaleProfileResponseArray []AutoscaleProfileResponseInput

func (AutoscaleProfileResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoscaleProfileResponse)(nil)).Elem()
}

func (i AutoscaleProfileResponseArray) ToAutoscaleProfileResponseArrayOutput() AutoscaleProfileResponseArrayOutput {
	return i.ToAutoscaleProfileResponseArrayOutputWithContext(context.Background())
}

func (i AutoscaleProfileResponseArray) ToAutoscaleProfileResponseArrayOutputWithContext(ctx context.Context) AutoscaleProfileResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleProfileResponseArrayOutput)
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





type AzNsActionGroupInput interface {
	pulumi.Input

	ToAzNsActionGroupOutput() AzNsActionGroupOutput
	ToAzNsActionGroupOutputWithContext(context.Context) AzNsActionGroupOutput
}

type AzNsActionGroupArgs struct {
	ActionGroup          pulumi.StringArrayInput `pulumi:"actionGroup"`
	CustomWebhookPayload pulumi.StringPtrInput   `pulumi:"customWebhookPayload"`
	EmailSubject         pulumi.StringPtrInput   `pulumi:"emailSubject"`
}

func (AzNsActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzNsActionGroup)(nil)).Elem()
}

func (i AzNsActionGroupArgs) ToAzNsActionGroupOutput() AzNsActionGroupOutput {
	return i.ToAzNsActionGroupOutputWithContext(context.Background())
}

func (i AzNsActionGroupArgs) ToAzNsActionGroupOutputWithContext(ctx context.Context) AzNsActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzNsActionGroupOutput)
}

func (i AzNsActionGroupArgs) ToAzNsActionGroupPtrOutput() AzNsActionGroupPtrOutput {
	return i.ToAzNsActionGroupPtrOutputWithContext(context.Background())
}

func (i AzNsActionGroupArgs) ToAzNsActionGroupPtrOutputWithContext(ctx context.Context) AzNsActionGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzNsActionGroupOutput).ToAzNsActionGroupPtrOutputWithContext(ctx)
}









type AzNsActionGroupPtrInput interface {
	pulumi.Input

	ToAzNsActionGroupPtrOutput() AzNsActionGroupPtrOutput
	ToAzNsActionGroupPtrOutputWithContext(context.Context) AzNsActionGroupPtrOutput
}

type azNsActionGroupPtrType AzNsActionGroupArgs

func AzNsActionGroupPtr(v *AzNsActionGroupArgs) AzNsActionGroupPtrInput {
	return (*azNsActionGroupPtrType)(v)
}

func (*azNsActionGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzNsActionGroup)(nil)).Elem()
}

func (i *azNsActionGroupPtrType) ToAzNsActionGroupPtrOutput() AzNsActionGroupPtrOutput {
	return i.ToAzNsActionGroupPtrOutputWithContext(context.Background())
}

func (i *azNsActionGroupPtrType) ToAzNsActionGroupPtrOutputWithContext(ctx context.Context) AzNsActionGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzNsActionGroupPtrOutput)
}

type AzNsActionGroupOutput struct{ *pulumi.OutputState }

func (AzNsActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzNsActionGroup)(nil)).Elem()
}

func (o AzNsActionGroupOutput) ToAzNsActionGroupOutput() AzNsActionGroupOutput {
	return o
}

func (o AzNsActionGroupOutput) ToAzNsActionGroupOutputWithContext(ctx context.Context) AzNsActionGroupOutput {
	return o
}

func (o AzNsActionGroupOutput) ToAzNsActionGroupPtrOutput() AzNsActionGroupPtrOutput {
	return o.ToAzNsActionGroupPtrOutputWithContext(context.Background())
}

func (o AzNsActionGroupOutput) ToAzNsActionGroupPtrOutputWithContext(ctx context.Context) AzNsActionGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzNsActionGroup) *AzNsActionGroup {
		return &v
	}).(AzNsActionGroupPtrOutput)
}

func (o AzNsActionGroupOutput) ActionGroup() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzNsActionGroup) []string { return v.ActionGroup }).(pulumi.StringArrayOutput)
}

func (o AzNsActionGroupOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzNsActionGroup) *string { return v.CustomWebhookPayload }).(pulumi.StringPtrOutput)
}

func (o AzNsActionGroupOutput) EmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzNsActionGroup) *string { return v.EmailSubject }).(pulumi.StringPtrOutput)
}

type AzNsActionGroupPtrOutput struct{ *pulumi.OutputState }

func (AzNsActionGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzNsActionGroup)(nil)).Elem()
}

func (o AzNsActionGroupPtrOutput) ToAzNsActionGroupPtrOutput() AzNsActionGroupPtrOutput {
	return o
}

func (o AzNsActionGroupPtrOutput) ToAzNsActionGroupPtrOutputWithContext(ctx context.Context) AzNsActionGroupPtrOutput {
	return o
}

func (o AzNsActionGroupPtrOutput) Elem() AzNsActionGroupOutput {
	return o.ApplyT(func(v *AzNsActionGroup) AzNsActionGroup {
		if v != nil {
			return *v
		}
		var ret AzNsActionGroup
		return ret
	}).(AzNsActionGroupOutput)
}

func (o AzNsActionGroupPtrOutput) ActionGroup() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzNsActionGroup) []string {
		if v == nil {
			return nil
		}
		return v.ActionGroup
	}).(pulumi.StringArrayOutput)
}

func (o AzNsActionGroupPtrOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzNsActionGroup) *string {
		if v == nil {
			return nil
		}
		return v.CustomWebhookPayload
	}).(pulumi.StringPtrOutput)
}

func (o AzNsActionGroupPtrOutput) EmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzNsActionGroup) *string {
		if v == nil {
			return nil
		}
		return v.EmailSubject
	}).(pulumi.StringPtrOutput)
}

type AzNsActionGroupResponse struct {
	ActionGroup          []string `pulumi:"actionGroup"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	EmailSubject         *string  `pulumi:"emailSubject"`
}





type AzNsActionGroupResponseInput interface {
	pulumi.Input

	ToAzNsActionGroupResponseOutput() AzNsActionGroupResponseOutput
	ToAzNsActionGroupResponseOutputWithContext(context.Context) AzNsActionGroupResponseOutput
}

type AzNsActionGroupResponseArgs struct {
	ActionGroup          pulumi.StringArrayInput `pulumi:"actionGroup"`
	CustomWebhookPayload pulumi.StringPtrInput   `pulumi:"customWebhookPayload"`
	EmailSubject         pulumi.StringPtrInput   `pulumi:"emailSubject"`
}

func (AzNsActionGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzNsActionGroupResponse)(nil)).Elem()
}

func (i AzNsActionGroupResponseArgs) ToAzNsActionGroupResponseOutput() AzNsActionGroupResponseOutput {
	return i.ToAzNsActionGroupResponseOutputWithContext(context.Background())
}

func (i AzNsActionGroupResponseArgs) ToAzNsActionGroupResponseOutputWithContext(ctx context.Context) AzNsActionGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzNsActionGroupResponseOutput)
}

func (i AzNsActionGroupResponseArgs) ToAzNsActionGroupResponsePtrOutput() AzNsActionGroupResponsePtrOutput {
	return i.ToAzNsActionGroupResponsePtrOutputWithContext(context.Background())
}

func (i AzNsActionGroupResponseArgs) ToAzNsActionGroupResponsePtrOutputWithContext(ctx context.Context) AzNsActionGroupResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzNsActionGroupResponseOutput).ToAzNsActionGroupResponsePtrOutputWithContext(ctx)
}









type AzNsActionGroupResponsePtrInput interface {
	pulumi.Input

	ToAzNsActionGroupResponsePtrOutput() AzNsActionGroupResponsePtrOutput
	ToAzNsActionGroupResponsePtrOutputWithContext(context.Context) AzNsActionGroupResponsePtrOutput
}

type azNsActionGroupResponsePtrType AzNsActionGroupResponseArgs

func AzNsActionGroupResponsePtr(v *AzNsActionGroupResponseArgs) AzNsActionGroupResponsePtrInput {
	return (*azNsActionGroupResponsePtrType)(v)
}

func (*azNsActionGroupResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzNsActionGroupResponse)(nil)).Elem()
}

func (i *azNsActionGroupResponsePtrType) ToAzNsActionGroupResponsePtrOutput() AzNsActionGroupResponsePtrOutput {
	return i.ToAzNsActionGroupResponsePtrOutputWithContext(context.Background())
}

func (i *azNsActionGroupResponsePtrType) ToAzNsActionGroupResponsePtrOutputWithContext(ctx context.Context) AzNsActionGroupResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzNsActionGroupResponsePtrOutput)
}

type AzNsActionGroupResponseOutput struct{ *pulumi.OutputState }

func (AzNsActionGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzNsActionGroupResponse)(nil)).Elem()
}

func (o AzNsActionGroupResponseOutput) ToAzNsActionGroupResponseOutput() AzNsActionGroupResponseOutput {
	return o
}

func (o AzNsActionGroupResponseOutput) ToAzNsActionGroupResponseOutputWithContext(ctx context.Context) AzNsActionGroupResponseOutput {
	return o
}

func (o AzNsActionGroupResponseOutput) ToAzNsActionGroupResponsePtrOutput() AzNsActionGroupResponsePtrOutput {
	return o.ToAzNsActionGroupResponsePtrOutputWithContext(context.Background())
}

func (o AzNsActionGroupResponseOutput) ToAzNsActionGroupResponsePtrOutputWithContext(ctx context.Context) AzNsActionGroupResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzNsActionGroupResponse) *AzNsActionGroupResponse {
		return &v
	}).(AzNsActionGroupResponsePtrOutput)
}

func (o AzNsActionGroupResponseOutput) ActionGroup() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzNsActionGroupResponse) []string { return v.ActionGroup }).(pulumi.StringArrayOutput)
}

func (o AzNsActionGroupResponseOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzNsActionGroupResponse) *string { return v.CustomWebhookPayload }).(pulumi.StringPtrOutput)
}

func (o AzNsActionGroupResponseOutput) EmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzNsActionGroupResponse) *string { return v.EmailSubject }).(pulumi.StringPtrOutput)
}

type AzNsActionGroupResponsePtrOutput struct{ *pulumi.OutputState }

func (AzNsActionGroupResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzNsActionGroupResponse)(nil)).Elem()
}

func (o AzNsActionGroupResponsePtrOutput) ToAzNsActionGroupResponsePtrOutput() AzNsActionGroupResponsePtrOutput {
	return o
}

func (o AzNsActionGroupResponsePtrOutput) ToAzNsActionGroupResponsePtrOutputWithContext(ctx context.Context) AzNsActionGroupResponsePtrOutput {
	return o
}

func (o AzNsActionGroupResponsePtrOutput) Elem() AzNsActionGroupResponseOutput {
	return o.ApplyT(func(v *AzNsActionGroupResponse) AzNsActionGroupResponse {
		if v != nil {
			return *v
		}
		var ret AzNsActionGroupResponse
		return ret
	}).(AzNsActionGroupResponseOutput)
}

func (o AzNsActionGroupResponsePtrOutput) ActionGroup() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzNsActionGroupResponse) []string {
		if v == nil {
			return nil
		}
		return v.ActionGroup
	}).(pulumi.StringArrayOutput)
}

func (o AzNsActionGroupResponsePtrOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzNsActionGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomWebhookPayload
	}).(pulumi.StringPtrOutput)
}

func (o AzNsActionGroupResponsePtrOutput) EmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzNsActionGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.EmailSubject
	}).(pulumi.StringPtrOutput)
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





type AzureAppPushReceiverResponseInput interface {
	pulumi.Input

	ToAzureAppPushReceiverResponseOutput() AzureAppPushReceiverResponseOutput
	ToAzureAppPushReceiverResponseOutputWithContext(context.Context) AzureAppPushReceiverResponseOutput
}

type AzureAppPushReceiverResponseArgs struct {
	EmailAddress pulumi.StringInput `pulumi:"emailAddress"`
	Name         pulumi.StringInput `pulumi:"name"`
}

func (AzureAppPushReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureAppPushReceiverResponse)(nil)).Elem()
}

func (i AzureAppPushReceiverResponseArgs) ToAzureAppPushReceiverResponseOutput() AzureAppPushReceiverResponseOutput {
	return i.ToAzureAppPushReceiverResponseOutputWithContext(context.Background())
}

func (i AzureAppPushReceiverResponseArgs) ToAzureAppPushReceiverResponseOutputWithContext(ctx context.Context) AzureAppPushReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureAppPushReceiverResponseOutput)
}





type AzureAppPushReceiverResponseArrayInput interface {
	pulumi.Input

	ToAzureAppPushReceiverResponseArrayOutput() AzureAppPushReceiverResponseArrayOutput
	ToAzureAppPushReceiverResponseArrayOutputWithContext(context.Context) AzureAppPushReceiverResponseArrayOutput
}

type AzureAppPushReceiverResponseArray []AzureAppPushReceiverResponseInput

func (AzureAppPushReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureAppPushReceiverResponse)(nil)).Elem()
}

func (i AzureAppPushReceiverResponseArray) ToAzureAppPushReceiverResponseArrayOutput() AzureAppPushReceiverResponseArrayOutput {
	return i.ToAzureAppPushReceiverResponseArrayOutputWithContext(context.Background())
}

func (i AzureAppPushReceiverResponseArray) ToAzureAppPushReceiverResponseArrayOutputWithContext(ctx context.Context) AzureAppPushReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureAppPushReceiverResponseArrayOutput)
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





type AzureFunctionReceiverResponseInput interface {
	pulumi.Input

	ToAzureFunctionReceiverResponseOutput() AzureFunctionReceiverResponseOutput
	ToAzureFunctionReceiverResponseOutputWithContext(context.Context) AzureFunctionReceiverResponseOutput
}

type AzureFunctionReceiverResponseArgs struct {
	FunctionAppResourceId pulumi.StringInput  `pulumi:"functionAppResourceId"`
	FunctionName          pulumi.StringInput  `pulumi:"functionName"`
	HttpTriggerUrl        pulumi.StringInput  `pulumi:"httpTriggerUrl"`
	Name                  pulumi.StringInput  `pulumi:"name"`
	UseCommonAlertSchema  pulumi.BoolPtrInput `pulumi:"useCommonAlertSchema"`
}

func (AzureFunctionReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionReceiverResponse)(nil)).Elem()
}

func (i AzureFunctionReceiverResponseArgs) ToAzureFunctionReceiverResponseOutput() AzureFunctionReceiverResponseOutput {
	return i.ToAzureFunctionReceiverResponseOutputWithContext(context.Background())
}

func (i AzureFunctionReceiverResponseArgs) ToAzureFunctionReceiverResponseOutputWithContext(ctx context.Context) AzureFunctionReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFunctionReceiverResponseOutput)
}





type AzureFunctionReceiverResponseArrayInput interface {
	pulumi.Input

	ToAzureFunctionReceiverResponseArrayOutput() AzureFunctionReceiverResponseArrayOutput
	ToAzureFunctionReceiverResponseArrayOutputWithContext(context.Context) AzureFunctionReceiverResponseArrayOutput
}

type AzureFunctionReceiverResponseArray []AzureFunctionReceiverResponseInput

func (AzureFunctionReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureFunctionReceiverResponse)(nil)).Elem()
}

func (i AzureFunctionReceiverResponseArray) ToAzureFunctionReceiverResponseArrayOutput() AzureFunctionReceiverResponseArrayOutput {
	return i.ToAzureFunctionReceiverResponseArrayOutputWithContext(context.Background())
}

func (i AzureFunctionReceiverResponseArray) ToAzureFunctionReceiverResponseArrayOutputWithContext(ctx context.Context) AzureFunctionReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFunctionReceiverResponseArrayOutput)
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





type CriteriaInput interface {
	pulumi.Input

	ToCriteriaOutput() CriteriaOutput
	ToCriteriaOutputWithContext(context.Context) CriteriaOutput
}

type CriteriaArgs struct {
	Dimensions DimensionArrayInput `pulumi:"dimensions"`
	MetricName pulumi.StringInput  `pulumi:"metricName"`
}

func (CriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Criteria)(nil)).Elem()
}

func (i CriteriaArgs) ToCriteriaOutput() CriteriaOutput {
	return i.ToCriteriaOutputWithContext(context.Background())
}

func (i CriteriaArgs) ToCriteriaOutputWithContext(ctx context.Context) CriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CriteriaOutput)
}





type CriteriaArrayInput interface {
	pulumi.Input

	ToCriteriaArrayOutput() CriteriaArrayOutput
	ToCriteriaArrayOutputWithContext(context.Context) CriteriaArrayOutput
}

type CriteriaArray []CriteriaInput

func (CriteriaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Criteria)(nil)).Elem()
}

func (i CriteriaArray) ToCriteriaArrayOutput() CriteriaArrayOutput {
	return i.ToCriteriaArrayOutputWithContext(context.Background())
}

func (i CriteriaArray) ToCriteriaArrayOutputWithContext(ctx context.Context) CriteriaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CriteriaArrayOutput)
}

type CriteriaOutput struct{ *pulumi.OutputState }

func (CriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Criteria)(nil)).Elem()
}

func (o CriteriaOutput) ToCriteriaOutput() CriteriaOutput {
	return o
}

func (o CriteriaOutput) ToCriteriaOutputWithContext(ctx context.Context) CriteriaOutput {
	return o
}

func (o CriteriaOutput) Dimensions() DimensionArrayOutput {
	return o.ApplyT(func(v Criteria) []Dimension { return v.Dimensions }).(DimensionArrayOutput)
}

func (o CriteriaOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v Criteria) string { return v.MetricName }).(pulumi.StringOutput)
}

type CriteriaArrayOutput struct{ *pulumi.OutputState }

func (CriteriaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Criteria)(nil)).Elem()
}

func (o CriteriaArrayOutput) ToCriteriaArrayOutput() CriteriaArrayOutput {
	return o
}

func (o CriteriaArrayOutput) ToCriteriaArrayOutputWithContext(ctx context.Context) CriteriaArrayOutput {
	return o
}

func (o CriteriaArrayOutput) Index(i pulumi.IntInput) CriteriaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Criteria {
		return vs[0].([]Criteria)[vs[1].(int)]
	}).(CriteriaOutput)
}

type CriteriaResponse struct {
	Dimensions []DimensionResponse `pulumi:"dimensions"`
	MetricName string              `pulumi:"metricName"`
}





type CriteriaResponseInput interface {
	pulumi.Input

	ToCriteriaResponseOutput() CriteriaResponseOutput
	ToCriteriaResponseOutputWithContext(context.Context) CriteriaResponseOutput
}

type CriteriaResponseArgs struct {
	Dimensions DimensionResponseArrayInput `pulumi:"dimensions"`
	MetricName pulumi.StringInput          `pulumi:"metricName"`
}

func (CriteriaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CriteriaResponse)(nil)).Elem()
}

func (i CriteriaResponseArgs) ToCriteriaResponseOutput() CriteriaResponseOutput {
	return i.ToCriteriaResponseOutputWithContext(context.Background())
}

func (i CriteriaResponseArgs) ToCriteriaResponseOutputWithContext(ctx context.Context) CriteriaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CriteriaResponseOutput)
}





type CriteriaResponseArrayInput interface {
	pulumi.Input

	ToCriteriaResponseArrayOutput() CriteriaResponseArrayOutput
	ToCriteriaResponseArrayOutputWithContext(context.Context) CriteriaResponseArrayOutput
}

type CriteriaResponseArray []CriteriaResponseInput

func (CriteriaResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CriteriaResponse)(nil)).Elem()
}

func (i CriteriaResponseArray) ToCriteriaResponseArrayOutput() CriteriaResponseArrayOutput {
	return i.ToCriteriaResponseArrayOutputWithContext(context.Background())
}

func (i CriteriaResponseArray) ToCriteriaResponseArrayOutputWithContext(ctx context.Context) CriteriaResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CriteriaResponseArrayOutput)
}

type CriteriaResponseOutput struct{ *pulumi.OutputState }

func (CriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CriteriaResponse)(nil)).Elem()
}

func (o CriteriaResponseOutput) ToCriteriaResponseOutput() CriteriaResponseOutput {
	return o
}

func (o CriteriaResponseOutput) ToCriteriaResponseOutputWithContext(ctx context.Context) CriteriaResponseOutput {
	return o
}

func (o CriteriaResponseOutput) Dimensions() DimensionResponseArrayOutput {
	return o.ApplyT(func(v CriteriaResponse) []DimensionResponse { return v.Dimensions }).(DimensionResponseArrayOutput)
}

func (o CriteriaResponseOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v CriteriaResponse) string { return v.MetricName }).(pulumi.StringOutput)
}

type CriteriaResponseArrayOutput struct{ *pulumi.OutputState }

func (CriteriaResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CriteriaResponse)(nil)).Elem()
}

func (o CriteriaResponseArrayOutput) ToCriteriaResponseArrayOutput() CriteriaResponseArrayOutput {
	return o
}

func (o CriteriaResponseArrayOutput) ToCriteriaResponseArrayOutputWithContext(ctx context.Context) CriteriaResponseArrayOutput {
	return o
}

func (o CriteriaResponseArrayOutput) Index(i pulumi.IntInput) CriteriaResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CriteriaResponse {
		return vs[0].([]CriteriaResponse)[vs[1].(int)]
	}).(CriteriaResponseOutput)
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





type DataCollectionEndpointResourceResponseSystemDataInput interface {
	pulumi.Input

	ToDataCollectionEndpointResourceResponseSystemDataOutput() DataCollectionEndpointResourceResponseSystemDataOutput
	ToDataCollectionEndpointResourceResponseSystemDataOutputWithContext(context.Context) DataCollectionEndpointResourceResponseSystemDataOutput
}

type DataCollectionEndpointResourceResponseSystemDataArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (DataCollectionEndpointResourceResponseSystemDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResourceResponseSystemData)(nil)).Elem()
}

func (i DataCollectionEndpointResourceResponseSystemDataArgs) ToDataCollectionEndpointResourceResponseSystemDataOutput() DataCollectionEndpointResourceResponseSystemDataOutput {
	return i.ToDataCollectionEndpointResourceResponseSystemDataOutputWithContext(context.Background())
}

func (i DataCollectionEndpointResourceResponseSystemDataArgs) ToDataCollectionEndpointResourceResponseSystemDataOutputWithContext(ctx context.Context) DataCollectionEndpointResourceResponseSystemDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResourceResponseSystemDataOutput)
}

func (i DataCollectionEndpointResourceResponseSystemDataArgs) ToDataCollectionEndpointResourceResponseSystemDataPtrOutput() DataCollectionEndpointResourceResponseSystemDataPtrOutput {
	return i.ToDataCollectionEndpointResourceResponseSystemDataPtrOutputWithContext(context.Background())
}

func (i DataCollectionEndpointResourceResponseSystemDataArgs) ToDataCollectionEndpointResourceResponseSystemDataPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResourceResponseSystemDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResourceResponseSystemDataOutput).ToDataCollectionEndpointResourceResponseSystemDataPtrOutputWithContext(ctx)
}









type DataCollectionEndpointResourceResponseSystemDataPtrInput interface {
	pulumi.Input

	ToDataCollectionEndpointResourceResponseSystemDataPtrOutput() DataCollectionEndpointResourceResponseSystemDataPtrOutput
	ToDataCollectionEndpointResourceResponseSystemDataPtrOutputWithContext(context.Context) DataCollectionEndpointResourceResponseSystemDataPtrOutput
}

type dataCollectionEndpointResourceResponseSystemDataPtrType DataCollectionEndpointResourceResponseSystemDataArgs

func DataCollectionEndpointResourceResponseSystemDataPtr(v *DataCollectionEndpointResourceResponseSystemDataArgs) DataCollectionEndpointResourceResponseSystemDataPtrInput {
	return (*dataCollectionEndpointResourceResponseSystemDataPtrType)(v)
}

func (*dataCollectionEndpointResourceResponseSystemDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResourceResponseSystemData)(nil)).Elem()
}

func (i *dataCollectionEndpointResourceResponseSystemDataPtrType) ToDataCollectionEndpointResourceResponseSystemDataPtrOutput() DataCollectionEndpointResourceResponseSystemDataPtrOutput {
	return i.ToDataCollectionEndpointResourceResponseSystemDataPtrOutputWithContext(context.Background())
}

func (i *dataCollectionEndpointResourceResponseSystemDataPtrType) ToDataCollectionEndpointResourceResponseSystemDataPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResourceResponseSystemDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResourceResponseSystemDataPtrOutput)
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

func (o DataCollectionEndpointResourceResponseSystemDataOutput) ToDataCollectionEndpointResourceResponseSystemDataPtrOutput() DataCollectionEndpointResourceResponseSystemDataPtrOutput {
	return o.ToDataCollectionEndpointResourceResponseSystemDataPtrOutputWithContext(context.Background())
}

func (o DataCollectionEndpointResourceResponseSystemDataOutput) ToDataCollectionEndpointResourceResponseSystemDataPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResourceResponseSystemDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionEndpointResourceResponseSystemData) *DataCollectionEndpointResourceResponseSystemData {
		return &v
	}).(DataCollectionEndpointResourceResponseSystemDataPtrOutput)
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

type DataCollectionEndpointResourceResponseSystemDataPtrOutput struct{ *pulumi.OutputState }

func (DataCollectionEndpointResourceResponseSystemDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResourceResponseSystemData)(nil)).Elem()
}

func (o DataCollectionEndpointResourceResponseSystemDataPtrOutput) ToDataCollectionEndpointResourceResponseSystemDataPtrOutput() DataCollectionEndpointResourceResponseSystemDataPtrOutput {
	return o
}

func (o DataCollectionEndpointResourceResponseSystemDataPtrOutput) ToDataCollectionEndpointResourceResponseSystemDataPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResourceResponseSystemDataPtrOutput {
	return o
}

func (o DataCollectionEndpointResourceResponseSystemDataPtrOutput) Elem() DataCollectionEndpointResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResourceResponseSystemData) DataCollectionEndpointResourceResponseSystemData {
		if v != nil {
			return *v
		}
		var ret DataCollectionEndpointResourceResponseSystemData
		return ret
	}).(DataCollectionEndpointResourceResponseSystemDataOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataPtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResourceResponseSystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataPtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResourceResponseSystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataPtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResourceResponseSystemData) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataPtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResourceResponseSystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataPtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResourceResponseSystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o DataCollectionEndpointResourceResponseSystemDataPtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionEndpointResourceResponseSystemData) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type DataCollectionEndpointResponseConfigurationAccess struct {
	Endpoint string `pulumi:"endpoint"`
}





type DataCollectionEndpointResponseConfigurationAccessInput interface {
	pulumi.Input

	ToDataCollectionEndpointResponseConfigurationAccessOutput() DataCollectionEndpointResponseConfigurationAccessOutput
	ToDataCollectionEndpointResponseConfigurationAccessOutputWithContext(context.Context) DataCollectionEndpointResponseConfigurationAccessOutput
}

type DataCollectionEndpointResponseConfigurationAccessArgs struct {
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
}

func (DataCollectionEndpointResponseConfigurationAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResponseConfigurationAccess)(nil)).Elem()
}

func (i DataCollectionEndpointResponseConfigurationAccessArgs) ToDataCollectionEndpointResponseConfigurationAccessOutput() DataCollectionEndpointResponseConfigurationAccessOutput {
	return i.ToDataCollectionEndpointResponseConfigurationAccessOutputWithContext(context.Background())
}

func (i DataCollectionEndpointResponseConfigurationAccessArgs) ToDataCollectionEndpointResponseConfigurationAccessOutputWithContext(ctx context.Context) DataCollectionEndpointResponseConfigurationAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResponseConfigurationAccessOutput)
}

func (i DataCollectionEndpointResponseConfigurationAccessArgs) ToDataCollectionEndpointResponseConfigurationAccessPtrOutput() DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return i.ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(context.Background())
}

func (i DataCollectionEndpointResponseConfigurationAccessArgs) ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResponseConfigurationAccessOutput).ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(ctx)
}









type DataCollectionEndpointResponseConfigurationAccessPtrInput interface {
	pulumi.Input

	ToDataCollectionEndpointResponseConfigurationAccessPtrOutput() DataCollectionEndpointResponseConfigurationAccessPtrOutput
	ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(context.Context) DataCollectionEndpointResponseConfigurationAccessPtrOutput
}

type dataCollectionEndpointResponseConfigurationAccessPtrType DataCollectionEndpointResponseConfigurationAccessArgs

func DataCollectionEndpointResponseConfigurationAccessPtr(v *DataCollectionEndpointResponseConfigurationAccessArgs) DataCollectionEndpointResponseConfigurationAccessPtrInput {
	return (*dataCollectionEndpointResponseConfigurationAccessPtrType)(v)
}

func (*dataCollectionEndpointResponseConfigurationAccessPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResponseConfigurationAccess)(nil)).Elem()
}

func (i *dataCollectionEndpointResponseConfigurationAccessPtrType) ToDataCollectionEndpointResponseConfigurationAccessPtrOutput() DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return i.ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(context.Background())
}

func (i *dataCollectionEndpointResponseConfigurationAccessPtrType) ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResponseConfigurationAccessPtrOutput)
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

func (o DataCollectionEndpointResponseConfigurationAccessOutput) ToDataCollectionEndpointResponseConfigurationAccessPtrOutput() DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return o.ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(context.Background())
}

func (o DataCollectionEndpointResponseConfigurationAccessOutput) ToDataCollectionEndpointResponseConfigurationAccessPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseConfigurationAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionEndpointResponseConfigurationAccess) *DataCollectionEndpointResponseConfigurationAccess {
		return &v
	}).(DataCollectionEndpointResponseConfigurationAccessPtrOutput)
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





type DataCollectionEndpointResponseLogsIngestionInput interface {
	pulumi.Input

	ToDataCollectionEndpointResponseLogsIngestionOutput() DataCollectionEndpointResponseLogsIngestionOutput
	ToDataCollectionEndpointResponseLogsIngestionOutputWithContext(context.Context) DataCollectionEndpointResponseLogsIngestionOutput
}

type DataCollectionEndpointResponseLogsIngestionArgs struct {
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
}

func (DataCollectionEndpointResponseLogsIngestionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResponseLogsIngestion)(nil)).Elem()
}

func (i DataCollectionEndpointResponseLogsIngestionArgs) ToDataCollectionEndpointResponseLogsIngestionOutput() DataCollectionEndpointResponseLogsIngestionOutput {
	return i.ToDataCollectionEndpointResponseLogsIngestionOutputWithContext(context.Background())
}

func (i DataCollectionEndpointResponseLogsIngestionArgs) ToDataCollectionEndpointResponseLogsIngestionOutputWithContext(ctx context.Context) DataCollectionEndpointResponseLogsIngestionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResponseLogsIngestionOutput)
}

func (i DataCollectionEndpointResponseLogsIngestionArgs) ToDataCollectionEndpointResponseLogsIngestionPtrOutput() DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return i.ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(context.Background())
}

func (i DataCollectionEndpointResponseLogsIngestionArgs) ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResponseLogsIngestionOutput).ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(ctx)
}









type DataCollectionEndpointResponseLogsIngestionPtrInput interface {
	pulumi.Input

	ToDataCollectionEndpointResponseLogsIngestionPtrOutput() DataCollectionEndpointResponseLogsIngestionPtrOutput
	ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(context.Context) DataCollectionEndpointResponseLogsIngestionPtrOutput
}

type dataCollectionEndpointResponseLogsIngestionPtrType DataCollectionEndpointResponseLogsIngestionArgs

func DataCollectionEndpointResponseLogsIngestionPtr(v *DataCollectionEndpointResponseLogsIngestionArgs) DataCollectionEndpointResponseLogsIngestionPtrInput {
	return (*dataCollectionEndpointResponseLogsIngestionPtrType)(v)
}

func (*dataCollectionEndpointResponseLogsIngestionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResponseLogsIngestion)(nil)).Elem()
}

func (i *dataCollectionEndpointResponseLogsIngestionPtrType) ToDataCollectionEndpointResponseLogsIngestionPtrOutput() DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return i.ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(context.Background())
}

func (i *dataCollectionEndpointResponseLogsIngestionPtrType) ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResponseLogsIngestionPtrOutput)
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

func (o DataCollectionEndpointResponseLogsIngestionOutput) ToDataCollectionEndpointResponseLogsIngestionPtrOutput() DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return o.ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(context.Background())
}

func (o DataCollectionEndpointResponseLogsIngestionOutput) ToDataCollectionEndpointResponseLogsIngestionPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseLogsIngestionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionEndpointResponseLogsIngestion) *DataCollectionEndpointResponseLogsIngestion {
		return &v
	}).(DataCollectionEndpointResponseLogsIngestionPtrOutput)
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





type DataCollectionEndpointResponseNetworkAclsInput interface {
	pulumi.Input

	ToDataCollectionEndpointResponseNetworkAclsOutput() DataCollectionEndpointResponseNetworkAclsOutput
	ToDataCollectionEndpointResponseNetworkAclsOutputWithContext(context.Context) DataCollectionEndpointResponseNetworkAclsOutput
}

type DataCollectionEndpointResponseNetworkAclsArgs struct {
	PublicNetworkAccess pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (DataCollectionEndpointResponseNetworkAclsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionEndpointResponseNetworkAcls)(nil)).Elem()
}

func (i DataCollectionEndpointResponseNetworkAclsArgs) ToDataCollectionEndpointResponseNetworkAclsOutput() DataCollectionEndpointResponseNetworkAclsOutput {
	return i.ToDataCollectionEndpointResponseNetworkAclsOutputWithContext(context.Background())
}

func (i DataCollectionEndpointResponseNetworkAclsArgs) ToDataCollectionEndpointResponseNetworkAclsOutputWithContext(ctx context.Context) DataCollectionEndpointResponseNetworkAclsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResponseNetworkAclsOutput)
}

func (i DataCollectionEndpointResponseNetworkAclsArgs) ToDataCollectionEndpointResponseNetworkAclsPtrOutput() DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return i.ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(context.Background())
}

func (i DataCollectionEndpointResponseNetworkAclsArgs) ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResponseNetworkAclsOutput).ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(ctx)
}









type DataCollectionEndpointResponseNetworkAclsPtrInput interface {
	pulumi.Input

	ToDataCollectionEndpointResponseNetworkAclsPtrOutput() DataCollectionEndpointResponseNetworkAclsPtrOutput
	ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(context.Context) DataCollectionEndpointResponseNetworkAclsPtrOutput
}

type dataCollectionEndpointResponseNetworkAclsPtrType DataCollectionEndpointResponseNetworkAclsArgs

func DataCollectionEndpointResponseNetworkAclsPtr(v *DataCollectionEndpointResponseNetworkAclsArgs) DataCollectionEndpointResponseNetworkAclsPtrInput {
	return (*dataCollectionEndpointResponseNetworkAclsPtrType)(v)
}

func (*dataCollectionEndpointResponseNetworkAclsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionEndpointResponseNetworkAcls)(nil)).Elem()
}

func (i *dataCollectionEndpointResponseNetworkAclsPtrType) ToDataCollectionEndpointResponseNetworkAclsPtrOutput() DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return i.ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(context.Background())
}

func (i *dataCollectionEndpointResponseNetworkAclsPtrType) ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionEndpointResponseNetworkAclsPtrOutput)
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

func (o DataCollectionEndpointResponseNetworkAclsOutput) ToDataCollectionEndpointResponseNetworkAclsPtrOutput() DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return o.ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(context.Background())
}

func (o DataCollectionEndpointResponseNetworkAclsOutput) ToDataCollectionEndpointResponseNetworkAclsPtrOutputWithContext(ctx context.Context) DataCollectionEndpointResponseNetworkAclsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionEndpointResponseNetworkAcls) *DataCollectionEndpointResponseNetworkAcls {
		return &v
	}).(DataCollectionEndpointResponseNetworkAclsPtrOutput)
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





type DataCollectionRuleResponseDataSourcesInput interface {
	pulumi.Input

	ToDataCollectionRuleResponseDataSourcesOutput() DataCollectionRuleResponseDataSourcesOutput
	ToDataCollectionRuleResponseDataSourcesOutputWithContext(context.Context) DataCollectionRuleResponseDataSourcesOutput
}

type DataCollectionRuleResponseDataSourcesArgs struct {
	Extensions          ExtensionDataSourceResponseArrayInput       `pulumi:"extensions"`
	PerformanceCounters PerfCounterDataSourceResponseArrayInput     `pulumi:"performanceCounters"`
	Syslog              SyslogDataSourceResponseArrayInput          `pulumi:"syslog"`
	WindowsEventLogs    WindowsEventLogDataSourceResponseArrayInput `pulumi:"windowsEventLogs"`
}

func (DataCollectionRuleResponseDataSourcesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleResponseDataSources)(nil)).Elem()
}

func (i DataCollectionRuleResponseDataSourcesArgs) ToDataCollectionRuleResponseDataSourcesOutput() DataCollectionRuleResponseDataSourcesOutput {
	return i.ToDataCollectionRuleResponseDataSourcesOutputWithContext(context.Background())
}

func (i DataCollectionRuleResponseDataSourcesArgs) ToDataCollectionRuleResponseDataSourcesOutputWithContext(ctx context.Context) DataCollectionRuleResponseDataSourcesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleResponseDataSourcesOutput)
}

func (i DataCollectionRuleResponseDataSourcesArgs) ToDataCollectionRuleResponseDataSourcesPtrOutput() DataCollectionRuleResponseDataSourcesPtrOutput {
	return i.ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(context.Background())
}

func (i DataCollectionRuleResponseDataSourcesArgs) ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDataSourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleResponseDataSourcesOutput).ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(ctx)
}









type DataCollectionRuleResponseDataSourcesPtrInput interface {
	pulumi.Input

	ToDataCollectionRuleResponseDataSourcesPtrOutput() DataCollectionRuleResponseDataSourcesPtrOutput
	ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(context.Context) DataCollectionRuleResponseDataSourcesPtrOutput
}

type dataCollectionRuleResponseDataSourcesPtrType DataCollectionRuleResponseDataSourcesArgs

func DataCollectionRuleResponseDataSourcesPtr(v *DataCollectionRuleResponseDataSourcesArgs) DataCollectionRuleResponseDataSourcesPtrInput {
	return (*dataCollectionRuleResponseDataSourcesPtrType)(v)
}

func (*dataCollectionRuleResponseDataSourcesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleResponseDataSources)(nil)).Elem()
}

func (i *dataCollectionRuleResponseDataSourcesPtrType) ToDataCollectionRuleResponseDataSourcesPtrOutput() DataCollectionRuleResponseDataSourcesPtrOutput {
	return i.ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(context.Background())
}

func (i *dataCollectionRuleResponseDataSourcesPtrType) ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDataSourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleResponseDataSourcesPtrOutput)
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

func (o DataCollectionRuleResponseDataSourcesOutput) ToDataCollectionRuleResponseDataSourcesPtrOutput() DataCollectionRuleResponseDataSourcesPtrOutput {
	return o.ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(context.Background())
}

func (o DataCollectionRuleResponseDataSourcesOutput) ToDataCollectionRuleResponseDataSourcesPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDataSourcesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionRuleResponseDataSources) *DataCollectionRuleResponseDataSources {
		return &v
	}).(DataCollectionRuleResponseDataSourcesPtrOutput)
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





type DataCollectionRuleResponseDestinationsInput interface {
	pulumi.Input

	ToDataCollectionRuleResponseDestinationsOutput() DataCollectionRuleResponseDestinationsOutput
	ToDataCollectionRuleResponseDestinationsOutputWithContext(context.Context) DataCollectionRuleResponseDestinationsOutput
}

type DataCollectionRuleResponseDestinationsArgs struct {
	AzureMonitorMetrics DestinationsSpecResponseAzureMonitorMetricsPtrInput `pulumi:"azureMonitorMetrics"`
	LogAnalytics        LogAnalyticsDestinationResponseArrayInput           `pulumi:"logAnalytics"`
}

func (DataCollectionRuleResponseDestinationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCollectionRuleResponseDestinations)(nil)).Elem()
}

func (i DataCollectionRuleResponseDestinationsArgs) ToDataCollectionRuleResponseDestinationsOutput() DataCollectionRuleResponseDestinationsOutput {
	return i.ToDataCollectionRuleResponseDestinationsOutputWithContext(context.Background())
}

func (i DataCollectionRuleResponseDestinationsArgs) ToDataCollectionRuleResponseDestinationsOutputWithContext(ctx context.Context) DataCollectionRuleResponseDestinationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleResponseDestinationsOutput)
}

func (i DataCollectionRuleResponseDestinationsArgs) ToDataCollectionRuleResponseDestinationsPtrOutput() DataCollectionRuleResponseDestinationsPtrOutput {
	return i.ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(context.Background())
}

func (i DataCollectionRuleResponseDestinationsArgs) ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDestinationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleResponseDestinationsOutput).ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(ctx)
}









type DataCollectionRuleResponseDestinationsPtrInput interface {
	pulumi.Input

	ToDataCollectionRuleResponseDestinationsPtrOutput() DataCollectionRuleResponseDestinationsPtrOutput
	ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(context.Context) DataCollectionRuleResponseDestinationsPtrOutput
}

type dataCollectionRuleResponseDestinationsPtrType DataCollectionRuleResponseDestinationsArgs

func DataCollectionRuleResponseDestinationsPtr(v *DataCollectionRuleResponseDestinationsArgs) DataCollectionRuleResponseDestinationsPtrInput {
	return (*dataCollectionRuleResponseDestinationsPtrType)(v)
}

func (*dataCollectionRuleResponseDestinationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRuleResponseDestinations)(nil)).Elem()
}

func (i *dataCollectionRuleResponseDestinationsPtrType) ToDataCollectionRuleResponseDestinationsPtrOutput() DataCollectionRuleResponseDestinationsPtrOutput {
	return i.ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(context.Background())
}

func (i *dataCollectionRuleResponseDestinationsPtrType) ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDestinationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleResponseDestinationsPtrOutput)
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

func (o DataCollectionRuleResponseDestinationsOutput) ToDataCollectionRuleResponseDestinationsPtrOutput() DataCollectionRuleResponseDestinationsPtrOutput {
	return o.ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(context.Background())
}

func (o DataCollectionRuleResponseDestinationsOutput) ToDataCollectionRuleResponseDestinationsPtrOutputWithContext(ctx context.Context) DataCollectionRuleResponseDestinationsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCollectionRuleResponseDestinations) *DataCollectionRuleResponseDestinations {
		return &v
	}).(DataCollectionRuleResponseDestinationsPtrOutput)
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





type DataFlowResponseInput interface {
	pulumi.Input

	ToDataFlowResponseOutput() DataFlowResponseOutput
	ToDataFlowResponseOutputWithContext(context.Context) DataFlowResponseOutput
}

type DataFlowResponseArgs struct {
	Destinations pulumi.StringArrayInput `pulumi:"destinations"`
	Streams      pulumi.StringArrayInput `pulumi:"streams"`
}

func (DataFlowResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlowResponse)(nil)).Elem()
}

func (i DataFlowResponseArgs) ToDataFlowResponseOutput() DataFlowResponseOutput {
	return i.ToDataFlowResponseOutputWithContext(context.Background())
}

func (i DataFlowResponseArgs) ToDataFlowResponseOutputWithContext(ctx context.Context) DataFlowResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowResponseOutput)
}





type DataFlowResponseArrayInput interface {
	pulumi.Input

	ToDataFlowResponseArrayOutput() DataFlowResponseArrayOutput
	ToDataFlowResponseArrayOutputWithContext(context.Context) DataFlowResponseArrayOutput
}

type DataFlowResponseArray []DataFlowResponseInput

func (DataFlowResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataFlowResponse)(nil)).Elem()
}

func (i DataFlowResponseArray) ToDataFlowResponseArrayOutput() DataFlowResponseArrayOutput {
	return i.ToDataFlowResponseArrayOutputWithContext(context.Background())
}

func (i DataFlowResponseArray) ToDataFlowResponseArrayOutputWithContext(ctx context.Context) DataFlowResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowResponseArrayOutput)
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





type DataSourceConfigurationResponseInput interface {
	pulumi.Input

	ToDataSourceConfigurationResponseOutput() DataSourceConfigurationResponseOutput
	ToDataSourceConfigurationResponseOutputWithContext(context.Context) DataSourceConfigurationResponseOutput
}

type DataSourceConfigurationResponseArgs struct {
	EventLogs    EventLogConfigurationResponseArrayInput           `pulumi:"eventLogs"`
	PerfCounters PerformanceCounterConfigurationResponseArrayInput `pulumi:"perfCounters"`
	Providers    EtwProviderConfigurationResponseArrayInput        `pulumi:"providers"`
}

func (DataSourceConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceConfigurationResponse)(nil)).Elem()
}

func (i DataSourceConfigurationResponseArgs) ToDataSourceConfigurationResponseOutput() DataSourceConfigurationResponseOutput {
	return i.ToDataSourceConfigurationResponseOutputWithContext(context.Background())
}

func (i DataSourceConfigurationResponseArgs) ToDataSourceConfigurationResponseOutputWithContext(ctx context.Context) DataSourceConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceConfigurationResponseOutput)
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





type DataSourceResponseInput interface {
	pulumi.Input

	ToDataSourceResponseOutput() DataSourceResponseOutput
	ToDataSourceResponseOutputWithContext(context.Context) DataSourceResponseOutput
}

type DataSourceResponseArgs struct {
	Configuration DataSourceConfigurationResponseInput `pulumi:"configuration"`
	Kind          pulumi.StringInput                   `pulumi:"kind"`
	Sinks         SinkConfigurationResponseArrayInput  `pulumi:"sinks"`
}

func (DataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSourceResponse)(nil)).Elem()
}

func (i DataSourceResponseArgs) ToDataSourceResponseOutput() DataSourceResponseOutput {
	return i.ToDataSourceResponseOutputWithContext(context.Background())
}

func (i DataSourceResponseArgs) ToDataSourceResponseOutputWithContext(ctx context.Context) DataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceResponseOutput)
}





type DataSourceResponseArrayInput interface {
	pulumi.Input

	ToDataSourceResponseArrayOutput() DataSourceResponseArrayOutput
	ToDataSourceResponseArrayOutputWithContext(context.Context) DataSourceResponseArrayOutput
}

type DataSourceResponseArray []DataSourceResponseInput

func (DataSourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataSourceResponse)(nil)).Elem()
}

func (i DataSourceResponseArray) ToDataSourceResponseArrayOutput() DataSourceResponseArrayOutput {
	return i.ToDataSourceResponseArrayOutputWithContext(context.Background())
}

func (i DataSourceResponseArray) ToDataSourceResponseArrayOutputWithContext(ctx context.Context) DataSourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceResponseArrayOutput)
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





type DestinationsSpecResponseAzureMonitorMetricsInput interface {
	pulumi.Input

	ToDestinationsSpecResponseAzureMonitorMetricsOutput() DestinationsSpecResponseAzureMonitorMetricsOutput
	ToDestinationsSpecResponseAzureMonitorMetricsOutputWithContext(context.Context) DestinationsSpecResponseAzureMonitorMetricsOutput
}

type DestinationsSpecResponseAzureMonitorMetricsArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DestinationsSpecResponseAzureMonitorMetricsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DestinationsSpecResponseAzureMonitorMetrics)(nil)).Elem()
}

func (i DestinationsSpecResponseAzureMonitorMetricsArgs) ToDestinationsSpecResponseAzureMonitorMetricsOutput() DestinationsSpecResponseAzureMonitorMetricsOutput {
	return i.ToDestinationsSpecResponseAzureMonitorMetricsOutputWithContext(context.Background())
}

func (i DestinationsSpecResponseAzureMonitorMetricsArgs) ToDestinationsSpecResponseAzureMonitorMetricsOutputWithContext(ctx context.Context) DestinationsSpecResponseAzureMonitorMetricsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationsSpecResponseAzureMonitorMetricsOutput)
}

func (i DestinationsSpecResponseAzureMonitorMetricsArgs) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutput() DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return i.ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(context.Background())
}

func (i DestinationsSpecResponseAzureMonitorMetricsArgs) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationsSpecResponseAzureMonitorMetricsOutput).ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(ctx)
}









type DestinationsSpecResponseAzureMonitorMetricsPtrInput interface {
	pulumi.Input

	ToDestinationsSpecResponseAzureMonitorMetricsPtrOutput() DestinationsSpecResponseAzureMonitorMetricsPtrOutput
	ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(context.Context) DestinationsSpecResponseAzureMonitorMetricsPtrOutput
}

type destinationsSpecResponseAzureMonitorMetricsPtrType DestinationsSpecResponseAzureMonitorMetricsArgs

func DestinationsSpecResponseAzureMonitorMetricsPtr(v *DestinationsSpecResponseAzureMonitorMetricsArgs) DestinationsSpecResponseAzureMonitorMetricsPtrInput {
	return (*destinationsSpecResponseAzureMonitorMetricsPtrType)(v)
}

func (*destinationsSpecResponseAzureMonitorMetricsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DestinationsSpecResponseAzureMonitorMetrics)(nil)).Elem()
}

func (i *destinationsSpecResponseAzureMonitorMetricsPtrType) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutput() DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return i.ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(context.Background())
}

func (i *destinationsSpecResponseAzureMonitorMetricsPtrType) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DestinationsSpecResponseAzureMonitorMetricsPtrOutput)
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

func (o DestinationsSpecResponseAzureMonitorMetricsOutput) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutput() DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o.ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(context.Background())
}

func (o DestinationsSpecResponseAzureMonitorMetricsOutput) ToDestinationsSpecResponseAzureMonitorMetricsPtrOutputWithContext(ctx context.Context) DestinationsSpecResponseAzureMonitorMetricsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DestinationsSpecResponseAzureMonitorMetrics) *DestinationsSpecResponseAzureMonitorMetrics {
		return &v
	}).(DestinationsSpecResponseAzureMonitorMetricsPtrOutput)
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





type DimensionResponseInput interface {
	pulumi.Input

	ToDimensionResponseOutput() DimensionResponseOutput
	ToDimensionResponseOutputWithContext(context.Context) DimensionResponseOutput
}

type DimensionResponseArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (DimensionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DimensionResponse)(nil)).Elem()
}

func (i DimensionResponseArgs) ToDimensionResponseOutput() DimensionResponseOutput {
	return i.ToDimensionResponseOutputWithContext(context.Background())
}

func (i DimensionResponseArgs) ToDimensionResponseOutputWithContext(ctx context.Context) DimensionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DimensionResponseOutput)
}





type DimensionResponseArrayInput interface {
	pulumi.Input

	ToDimensionResponseArrayOutput() DimensionResponseArrayOutput
	ToDimensionResponseArrayOutputWithContext(context.Context) DimensionResponseArrayOutput
}

type DimensionResponseArray []DimensionResponseInput

func (DimensionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DimensionResponse)(nil)).Elem()
}

func (i DimensionResponseArray) ToDimensionResponseArrayOutput() DimensionResponseArrayOutput {
	return i.ToDimensionResponseArrayOutputWithContext(context.Background())
}

func (i DimensionResponseArray) ToDimensionResponseArrayOutputWithContext(ctx context.Context) DimensionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DimensionResponseArrayOutput)
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





type DynamicMetricCriteriaInput interface {
	pulumi.Input

	ToDynamicMetricCriteriaOutput() DynamicMetricCriteriaOutput
	ToDynamicMetricCriteriaOutputWithContext(context.Context) DynamicMetricCriteriaOutput
}

type DynamicMetricCriteriaArgs struct {
	AlertSensitivity     pulumi.StringInput                  `pulumi:"alertSensitivity"`
	CriterionType        pulumi.StringInput                  `pulumi:"criterionType"`
	Dimensions           MetricDimensionArrayInput           `pulumi:"dimensions"`
	FailingPeriods       DynamicThresholdFailingPeriodsInput `pulumi:"failingPeriods"`
	IgnoreDataBefore     pulumi.StringPtrInput               `pulumi:"ignoreDataBefore"`
	MetricName           pulumi.StringInput                  `pulumi:"metricName"`
	MetricNamespace      pulumi.StringPtrInput               `pulumi:"metricNamespace"`
	Name                 pulumi.StringInput                  `pulumi:"name"`
	Operator             pulumi.StringInput                  `pulumi:"operator"`
	SkipMetricValidation pulumi.BoolPtrInput                 `pulumi:"skipMetricValidation"`
	TimeAggregation      pulumi.StringInput                  `pulumi:"timeAggregation"`
}

func (DynamicMetricCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicMetricCriteria)(nil)).Elem()
}

func (i DynamicMetricCriteriaArgs) ToDynamicMetricCriteriaOutput() DynamicMetricCriteriaOutput {
	return i.ToDynamicMetricCriteriaOutputWithContext(context.Background())
}

func (i DynamicMetricCriteriaArgs) ToDynamicMetricCriteriaOutputWithContext(ctx context.Context) DynamicMetricCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicMetricCriteriaOutput)
}

type DynamicMetricCriteriaOutput struct{ *pulumi.OutputState }

func (DynamicMetricCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicMetricCriteria)(nil)).Elem()
}

func (o DynamicMetricCriteriaOutput) ToDynamicMetricCriteriaOutput() DynamicMetricCriteriaOutput {
	return o
}

func (o DynamicMetricCriteriaOutput) ToDynamicMetricCriteriaOutputWithContext(ctx context.Context) DynamicMetricCriteriaOutput {
	return o
}

func (o DynamicMetricCriteriaOutput) AlertSensitivity() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) string { return v.AlertSensitivity }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaOutput) CriterionType() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) string { return v.CriterionType }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaOutput) Dimensions() MetricDimensionArrayOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) []MetricDimension { return v.Dimensions }).(MetricDimensionArrayOutput)
}

func (o DynamicMetricCriteriaOutput) FailingPeriods() DynamicThresholdFailingPeriodsOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) DynamicThresholdFailingPeriods { return v.FailingPeriods }).(DynamicThresholdFailingPeriodsOutput)
}

func (o DynamicMetricCriteriaOutput) IgnoreDataBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) *string { return v.IgnoreDataBefore }).(pulumi.StringPtrOutput)
}

func (o DynamicMetricCriteriaOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o DynamicMetricCriteriaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) string { return v.Name }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) string { return v.Operator }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaOutput) SkipMetricValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) *bool { return v.SkipMetricValidation }).(pulumi.BoolPtrOutput)
}

func (o DynamicMetricCriteriaOutput) TimeAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteria) string { return v.TimeAggregation }).(pulumi.StringOutput)
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





type DynamicMetricCriteriaResponseInput interface {
	pulumi.Input

	ToDynamicMetricCriteriaResponseOutput() DynamicMetricCriteriaResponseOutput
	ToDynamicMetricCriteriaResponseOutputWithContext(context.Context) DynamicMetricCriteriaResponseOutput
}

type DynamicMetricCriteriaResponseArgs struct {
	AlertSensitivity     pulumi.StringInput                          `pulumi:"alertSensitivity"`
	CriterionType        pulumi.StringInput                          `pulumi:"criterionType"`
	Dimensions           MetricDimensionResponseArrayInput           `pulumi:"dimensions"`
	FailingPeriods       DynamicThresholdFailingPeriodsResponseInput `pulumi:"failingPeriods"`
	IgnoreDataBefore     pulumi.StringPtrInput                       `pulumi:"ignoreDataBefore"`
	MetricName           pulumi.StringInput                          `pulumi:"metricName"`
	MetricNamespace      pulumi.StringPtrInput                       `pulumi:"metricNamespace"`
	Name                 pulumi.StringInput                          `pulumi:"name"`
	Operator             pulumi.StringInput                          `pulumi:"operator"`
	SkipMetricValidation pulumi.BoolPtrInput                         `pulumi:"skipMetricValidation"`
	TimeAggregation      pulumi.StringInput                          `pulumi:"timeAggregation"`
}

func (DynamicMetricCriteriaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicMetricCriteriaResponse)(nil)).Elem()
}

func (i DynamicMetricCriteriaResponseArgs) ToDynamicMetricCriteriaResponseOutput() DynamicMetricCriteriaResponseOutput {
	return i.ToDynamicMetricCriteriaResponseOutputWithContext(context.Background())
}

func (i DynamicMetricCriteriaResponseArgs) ToDynamicMetricCriteriaResponseOutputWithContext(ctx context.Context) DynamicMetricCriteriaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicMetricCriteriaResponseOutput)
}

type DynamicMetricCriteriaResponseOutput struct{ *pulumi.OutputState }

func (DynamicMetricCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicMetricCriteriaResponse)(nil)).Elem()
}

func (o DynamicMetricCriteriaResponseOutput) ToDynamicMetricCriteriaResponseOutput() DynamicMetricCriteriaResponseOutput {
	return o
}

func (o DynamicMetricCriteriaResponseOutput) ToDynamicMetricCriteriaResponseOutputWithContext(ctx context.Context) DynamicMetricCriteriaResponseOutput {
	return o
}

func (o DynamicMetricCriteriaResponseOutput) AlertSensitivity() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) string { return v.AlertSensitivity }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaResponseOutput) CriterionType() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) string { return v.CriterionType }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaResponseOutput) Dimensions() MetricDimensionResponseArrayOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) []MetricDimensionResponse { return v.Dimensions }).(MetricDimensionResponseArrayOutput)
}

func (o DynamicMetricCriteriaResponseOutput) FailingPeriods() DynamicThresholdFailingPeriodsResponseOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) DynamicThresholdFailingPeriodsResponse { return v.FailingPeriods }).(DynamicThresholdFailingPeriodsResponseOutput)
}

func (o DynamicMetricCriteriaResponseOutput) IgnoreDataBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) *string { return v.IgnoreDataBefore }).(pulumi.StringPtrOutput)
}

func (o DynamicMetricCriteriaResponseOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaResponseOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o DynamicMetricCriteriaResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o DynamicMetricCriteriaResponseOutput) SkipMetricValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) *bool { return v.SkipMetricValidation }).(pulumi.BoolPtrOutput)
}

func (o DynamicMetricCriteriaResponseOutput) TimeAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v DynamicMetricCriteriaResponse) string { return v.TimeAggregation }).(pulumi.StringOutput)
}

type DynamicThresholdFailingPeriods struct {
	MinFailingPeriodsToAlert  float64 `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods float64 `pulumi:"numberOfEvaluationPeriods"`
}





type DynamicThresholdFailingPeriodsInput interface {
	pulumi.Input

	ToDynamicThresholdFailingPeriodsOutput() DynamicThresholdFailingPeriodsOutput
	ToDynamicThresholdFailingPeriodsOutputWithContext(context.Context) DynamicThresholdFailingPeriodsOutput
}

type DynamicThresholdFailingPeriodsArgs struct {
	MinFailingPeriodsToAlert  pulumi.Float64Input `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods pulumi.Float64Input `pulumi:"numberOfEvaluationPeriods"`
}

func (DynamicThresholdFailingPeriodsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicThresholdFailingPeriods)(nil)).Elem()
}

func (i DynamicThresholdFailingPeriodsArgs) ToDynamicThresholdFailingPeriodsOutput() DynamicThresholdFailingPeriodsOutput {
	return i.ToDynamicThresholdFailingPeriodsOutputWithContext(context.Background())
}

func (i DynamicThresholdFailingPeriodsArgs) ToDynamicThresholdFailingPeriodsOutputWithContext(ctx context.Context) DynamicThresholdFailingPeriodsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicThresholdFailingPeriodsOutput)
}

type DynamicThresholdFailingPeriodsOutput struct{ *pulumi.OutputState }

func (DynamicThresholdFailingPeriodsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicThresholdFailingPeriods)(nil)).Elem()
}

func (o DynamicThresholdFailingPeriodsOutput) ToDynamicThresholdFailingPeriodsOutput() DynamicThresholdFailingPeriodsOutput {
	return o
}

func (o DynamicThresholdFailingPeriodsOutput) ToDynamicThresholdFailingPeriodsOutputWithContext(ctx context.Context) DynamicThresholdFailingPeriodsOutput {
	return o
}

func (o DynamicThresholdFailingPeriodsOutput) MinFailingPeriodsToAlert() pulumi.Float64Output {
	return o.ApplyT(func(v DynamicThresholdFailingPeriods) float64 { return v.MinFailingPeriodsToAlert }).(pulumi.Float64Output)
}

func (o DynamicThresholdFailingPeriodsOutput) NumberOfEvaluationPeriods() pulumi.Float64Output {
	return o.ApplyT(func(v DynamicThresholdFailingPeriods) float64 { return v.NumberOfEvaluationPeriods }).(pulumi.Float64Output)
}

type DynamicThresholdFailingPeriodsResponse struct {
	MinFailingPeriodsToAlert  float64 `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods float64 `pulumi:"numberOfEvaluationPeriods"`
}





type DynamicThresholdFailingPeriodsResponseInput interface {
	pulumi.Input

	ToDynamicThresholdFailingPeriodsResponseOutput() DynamicThresholdFailingPeriodsResponseOutput
	ToDynamicThresholdFailingPeriodsResponseOutputWithContext(context.Context) DynamicThresholdFailingPeriodsResponseOutput
}

type DynamicThresholdFailingPeriodsResponseArgs struct {
	MinFailingPeriodsToAlert  pulumi.Float64Input `pulumi:"minFailingPeriodsToAlert"`
	NumberOfEvaluationPeriods pulumi.Float64Input `pulumi:"numberOfEvaluationPeriods"`
}

func (DynamicThresholdFailingPeriodsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicThresholdFailingPeriodsResponse)(nil)).Elem()
}

func (i DynamicThresholdFailingPeriodsResponseArgs) ToDynamicThresholdFailingPeriodsResponseOutput() DynamicThresholdFailingPeriodsResponseOutput {
	return i.ToDynamicThresholdFailingPeriodsResponseOutputWithContext(context.Background())
}

func (i DynamicThresholdFailingPeriodsResponseArgs) ToDynamicThresholdFailingPeriodsResponseOutputWithContext(ctx context.Context) DynamicThresholdFailingPeriodsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicThresholdFailingPeriodsResponseOutput)
}

type DynamicThresholdFailingPeriodsResponseOutput struct{ *pulumi.OutputState }

func (DynamicThresholdFailingPeriodsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DynamicThresholdFailingPeriodsResponse)(nil)).Elem()
}

func (o DynamicThresholdFailingPeriodsResponseOutput) ToDynamicThresholdFailingPeriodsResponseOutput() DynamicThresholdFailingPeriodsResponseOutput {
	return o
}

func (o DynamicThresholdFailingPeriodsResponseOutput) ToDynamicThresholdFailingPeriodsResponseOutputWithContext(ctx context.Context) DynamicThresholdFailingPeriodsResponseOutput {
	return o
}

func (o DynamicThresholdFailingPeriodsResponseOutput) MinFailingPeriodsToAlert() pulumi.Float64Output {
	return o.ApplyT(func(v DynamicThresholdFailingPeriodsResponse) float64 { return v.MinFailingPeriodsToAlert }).(pulumi.Float64Output)
}

func (o DynamicThresholdFailingPeriodsResponseOutput) NumberOfEvaluationPeriods() pulumi.Float64Output {
	return o.ApplyT(func(v DynamicThresholdFailingPeriodsResponse) float64 { return v.NumberOfEvaluationPeriods }).(pulumi.Float64Output)
}

type EmailNotification struct {
	CustomEmails                       []string `pulumi:"customEmails"`
	SendToSubscriptionAdministrator    *bool    `pulumi:"sendToSubscriptionAdministrator"`
	SendToSubscriptionCoAdministrators *bool    `pulumi:"sendToSubscriptionCoAdministrators"`
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





type EmailNotificationResponseInput interface {
	pulumi.Input

	ToEmailNotificationResponseOutput() EmailNotificationResponseOutput
	ToEmailNotificationResponseOutputWithContext(context.Context) EmailNotificationResponseOutput
}

type EmailNotificationResponseArgs struct {
	CustomEmails                       pulumi.StringArrayInput `pulumi:"customEmails"`
	SendToSubscriptionAdministrator    pulumi.BoolPtrInput     `pulumi:"sendToSubscriptionAdministrator"`
	SendToSubscriptionCoAdministrators pulumi.BoolPtrInput     `pulumi:"sendToSubscriptionCoAdministrators"`
}

func (EmailNotificationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailNotificationResponse)(nil)).Elem()
}

func (i EmailNotificationResponseArgs) ToEmailNotificationResponseOutput() EmailNotificationResponseOutput {
	return i.ToEmailNotificationResponseOutputWithContext(context.Background())
}

func (i EmailNotificationResponseArgs) ToEmailNotificationResponseOutputWithContext(ctx context.Context) EmailNotificationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailNotificationResponseOutput)
}

func (i EmailNotificationResponseArgs) ToEmailNotificationResponsePtrOutput() EmailNotificationResponsePtrOutput {
	return i.ToEmailNotificationResponsePtrOutputWithContext(context.Background())
}

func (i EmailNotificationResponseArgs) ToEmailNotificationResponsePtrOutputWithContext(ctx context.Context) EmailNotificationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailNotificationResponseOutput).ToEmailNotificationResponsePtrOutputWithContext(ctx)
}









type EmailNotificationResponsePtrInput interface {
	pulumi.Input

	ToEmailNotificationResponsePtrOutput() EmailNotificationResponsePtrOutput
	ToEmailNotificationResponsePtrOutputWithContext(context.Context) EmailNotificationResponsePtrOutput
}

type emailNotificationResponsePtrType EmailNotificationResponseArgs

func EmailNotificationResponsePtr(v *EmailNotificationResponseArgs) EmailNotificationResponsePtrInput {
	return (*emailNotificationResponsePtrType)(v)
}

func (*emailNotificationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailNotificationResponse)(nil)).Elem()
}

func (i *emailNotificationResponsePtrType) ToEmailNotificationResponsePtrOutput() EmailNotificationResponsePtrOutput {
	return i.ToEmailNotificationResponsePtrOutputWithContext(context.Background())
}

func (i *emailNotificationResponsePtrType) ToEmailNotificationResponsePtrOutputWithContext(ctx context.Context) EmailNotificationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailNotificationResponsePtrOutput)
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

func (o EmailNotificationResponseOutput) ToEmailNotificationResponsePtrOutput() EmailNotificationResponsePtrOutput {
	return o.ToEmailNotificationResponsePtrOutputWithContext(context.Background())
}

func (o EmailNotificationResponseOutput) ToEmailNotificationResponsePtrOutputWithContext(ctx context.Context) EmailNotificationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EmailNotificationResponse) *EmailNotificationResponse {
		return &v
	}).(EmailNotificationResponsePtrOutput)
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





type EmailReceiverResponseInput interface {
	pulumi.Input

	ToEmailReceiverResponseOutput() EmailReceiverResponseOutput
	ToEmailReceiverResponseOutputWithContext(context.Context) EmailReceiverResponseOutput
}

type EmailReceiverResponseArgs struct {
	EmailAddress         pulumi.StringInput  `pulumi:"emailAddress"`
	Name                 pulumi.StringInput  `pulumi:"name"`
	Status               pulumi.StringInput  `pulumi:"status"`
	UseCommonAlertSchema pulumi.BoolPtrInput `pulumi:"useCommonAlertSchema"`
}

func (EmailReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailReceiverResponse)(nil)).Elem()
}

func (i EmailReceiverResponseArgs) ToEmailReceiverResponseOutput() EmailReceiverResponseOutput {
	return i.ToEmailReceiverResponseOutputWithContext(context.Background())
}

func (i EmailReceiverResponseArgs) ToEmailReceiverResponseOutputWithContext(ctx context.Context) EmailReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailReceiverResponseOutput)
}





type EmailReceiverResponseArrayInput interface {
	pulumi.Input

	ToEmailReceiverResponseArrayOutput() EmailReceiverResponseArrayOutput
	ToEmailReceiverResponseArrayOutputWithContext(context.Context) EmailReceiverResponseArrayOutput
}

type EmailReceiverResponseArray []EmailReceiverResponseInput

func (EmailReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EmailReceiverResponse)(nil)).Elem()
}

func (i EmailReceiverResponseArray) ToEmailReceiverResponseArrayOutput() EmailReceiverResponseArrayOutput {
	return i.ToEmailReceiverResponseArrayOutputWithContext(context.Background())
}

func (i EmailReceiverResponseArray) ToEmailReceiverResponseArrayOutputWithContext(ctx context.Context) EmailReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailReceiverResponseArrayOutput)
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





type EtwEventConfigurationResponseInput interface {
	pulumi.Input

	ToEtwEventConfigurationResponseOutput() EtwEventConfigurationResponseOutput
	ToEtwEventConfigurationResponseOutputWithContext(context.Context) EtwEventConfigurationResponseOutput
}

type EtwEventConfigurationResponseArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	Id     pulumi.IntInput       `pulumi:"id"`
	Name   pulumi.StringInput    `pulumi:"name"`
}

func (EtwEventConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EtwEventConfigurationResponse)(nil)).Elem()
}

func (i EtwEventConfigurationResponseArgs) ToEtwEventConfigurationResponseOutput() EtwEventConfigurationResponseOutput {
	return i.ToEtwEventConfigurationResponseOutputWithContext(context.Background())
}

func (i EtwEventConfigurationResponseArgs) ToEtwEventConfigurationResponseOutputWithContext(ctx context.Context) EtwEventConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtwEventConfigurationResponseOutput)
}





type EtwEventConfigurationResponseArrayInput interface {
	pulumi.Input

	ToEtwEventConfigurationResponseArrayOutput() EtwEventConfigurationResponseArrayOutput
	ToEtwEventConfigurationResponseArrayOutputWithContext(context.Context) EtwEventConfigurationResponseArrayOutput
}

type EtwEventConfigurationResponseArray []EtwEventConfigurationResponseInput

func (EtwEventConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtwEventConfigurationResponse)(nil)).Elem()
}

func (i EtwEventConfigurationResponseArray) ToEtwEventConfigurationResponseArrayOutput() EtwEventConfigurationResponseArrayOutput {
	return i.ToEtwEventConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i EtwEventConfigurationResponseArray) ToEtwEventConfigurationResponseArrayOutputWithContext(ctx context.Context) EtwEventConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtwEventConfigurationResponseArrayOutput)
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





type EtwProviderConfigurationResponseInput interface {
	pulumi.Input

	ToEtwProviderConfigurationResponseOutput() EtwProviderConfigurationResponseOutput
	ToEtwProviderConfigurationResponseOutputWithContext(context.Context) EtwProviderConfigurationResponseOutput
}

type EtwProviderConfigurationResponseArgs struct {
	Events EtwEventConfigurationResponseArrayInput `pulumi:"events"`
	Id     pulumi.StringInput                      `pulumi:"id"`
}

func (EtwProviderConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EtwProviderConfigurationResponse)(nil)).Elem()
}

func (i EtwProviderConfigurationResponseArgs) ToEtwProviderConfigurationResponseOutput() EtwProviderConfigurationResponseOutput {
	return i.ToEtwProviderConfigurationResponseOutputWithContext(context.Background())
}

func (i EtwProviderConfigurationResponseArgs) ToEtwProviderConfigurationResponseOutputWithContext(ctx context.Context) EtwProviderConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtwProviderConfigurationResponseOutput)
}





type EtwProviderConfigurationResponseArrayInput interface {
	pulumi.Input

	ToEtwProviderConfigurationResponseArrayOutput() EtwProviderConfigurationResponseArrayOutput
	ToEtwProviderConfigurationResponseArrayOutputWithContext(context.Context) EtwProviderConfigurationResponseArrayOutput
}

type EtwProviderConfigurationResponseArray []EtwProviderConfigurationResponseInput

func (EtwProviderConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EtwProviderConfigurationResponse)(nil)).Elem()
}

func (i EtwProviderConfigurationResponseArray) ToEtwProviderConfigurationResponseArrayOutput() EtwProviderConfigurationResponseArrayOutput {
	return i.ToEtwProviderConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i EtwProviderConfigurationResponseArray) ToEtwProviderConfigurationResponseArrayOutputWithContext(ctx context.Context) EtwProviderConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EtwProviderConfigurationResponseArrayOutput)
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





type EventLogConfigurationResponseInput interface {
	pulumi.Input

	ToEventLogConfigurationResponseOutput() EventLogConfigurationResponseOutput
	ToEventLogConfigurationResponseOutputWithContext(context.Context) EventLogConfigurationResponseOutput
}

type EventLogConfigurationResponseArgs struct {
	Filter  pulumi.StringPtrInput `pulumi:"filter"`
	LogName pulumi.StringInput    `pulumi:"logName"`
}

func (EventLogConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventLogConfigurationResponse)(nil)).Elem()
}

func (i EventLogConfigurationResponseArgs) ToEventLogConfigurationResponseOutput() EventLogConfigurationResponseOutput {
	return i.ToEventLogConfigurationResponseOutputWithContext(context.Background())
}

func (i EventLogConfigurationResponseArgs) ToEventLogConfigurationResponseOutputWithContext(ctx context.Context) EventLogConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventLogConfigurationResponseOutput)
}





type EventLogConfigurationResponseArrayInput interface {
	pulumi.Input

	ToEventLogConfigurationResponseArrayOutput() EventLogConfigurationResponseArrayOutput
	ToEventLogConfigurationResponseArrayOutputWithContext(context.Context) EventLogConfigurationResponseArrayOutput
}

type EventLogConfigurationResponseArray []EventLogConfigurationResponseInput

func (EventLogConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventLogConfigurationResponse)(nil)).Elem()
}

func (i EventLogConfigurationResponseArray) ToEventLogConfigurationResponseArrayOutput() EventLogConfigurationResponseArrayOutput {
	return i.ToEventLogConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i EventLogConfigurationResponseArray) ToEventLogConfigurationResponseArrayOutputWithContext(ctx context.Context) EventLogConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventLogConfigurationResponseArrayOutput)
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





type ExtensionDataSourceResponseInput interface {
	pulumi.Input

	ToExtensionDataSourceResponseOutput() ExtensionDataSourceResponseOutput
	ToExtensionDataSourceResponseOutputWithContext(context.Context) ExtensionDataSourceResponseOutput
}

type ExtensionDataSourceResponseArgs struct {
	ExtensionName     pulumi.StringInput      `pulumi:"extensionName"`
	ExtensionSettings pulumi.Input            `pulumi:"extensionSettings"`
	InputDataSources  pulumi.StringArrayInput `pulumi:"inputDataSources"`
	Name              pulumi.StringPtrInput   `pulumi:"name"`
	Streams           pulumi.StringArrayInput `pulumi:"streams"`
}

func (ExtensionDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionDataSourceResponse)(nil)).Elem()
}

func (i ExtensionDataSourceResponseArgs) ToExtensionDataSourceResponseOutput() ExtensionDataSourceResponseOutput {
	return i.ToExtensionDataSourceResponseOutputWithContext(context.Background())
}

func (i ExtensionDataSourceResponseArgs) ToExtensionDataSourceResponseOutputWithContext(ctx context.Context) ExtensionDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionDataSourceResponseOutput)
}





type ExtensionDataSourceResponseArrayInput interface {
	pulumi.Input

	ToExtensionDataSourceResponseArrayOutput() ExtensionDataSourceResponseArrayOutput
	ToExtensionDataSourceResponseArrayOutputWithContext(context.Context) ExtensionDataSourceResponseArrayOutput
}

type ExtensionDataSourceResponseArray []ExtensionDataSourceResponseInput

func (ExtensionDataSourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionDataSourceResponse)(nil)).Elem()
}

func (i ExtensionDataSourceResponseArray) ToExtensionDataSourceResponseArrayOutput() ExtensionDataSourceResponseArrayOutput {
	return i.ToExtensionDataSourceResponseArrayOutputWithContext(context.Background())
}

func (i ExtensionDataSourceResponseArray) ToExtensionDataSourceResponseArrayOutputWithContext(ctx context.Context) ExtensionDataSourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionDataSourceResponseArrayOutput)
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





type ItsmReceiverResponseInput interface {
	pulumi.Input

	ToItsmReceiverResponseOutput() ItsmReceiverResponseOutput
	ToItsmReceiverResponseOutputWithContext(context.Context) ItsmReceiverResponseOutput
}

type ItsmReceiverResponseArgs struct {
	ConnectionId        pulumi.StringInput `pulumi:"connectionId"`
	Name                pulumi.StringInput `pulumi:"name"`
	Region              pulumi.StringInput `pulumi:"region"`
	TicketConfiguration pulumi.StringInput `pulumi:"ticketConfiguration"`
	WorkspaceId         pulumi.StringInput `pulumi:"workspaceId"`
}

func (ItsmReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ItsmReceiverResponse)(nil)).Elem()
}

func (i ItsmReceiverResponseArgs) ToItsmReceiverResponseOutput() ItsmReceiverResponseOutput {
	return i.ToItsmReceiverResponseOutputWithContext(context.Background())
}

func (i ItsmReceiverResponseArgs) ToItsmReceiverResponseOutputWithContext(ctx context.Context) ItsmReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItsmReceiverResponseOutput)
}





type ItsmReceiverResponseArrayInput interface {
	pulumi.Input

	ToItsmReceiverResponseArrayOutput() ItsmReceiverResponseArrayOutput
	ToItsmReceiverResponseArrayOutputWithContext(context.Context) ItsmReceiverResponseArrayOutput
}

type ItsmReceiverResponseArray []ItsmReceiverResponseInput

func (ItsmReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ItsmReceiverResponse)(nil)).Elem()
}

func (i ItsmReceiverResponseArray) ToItsmReceiverResponseArrayOutput() ItsmReceiverResponseArrayOutput {
	return i.ToItsmReceiverResponseArrayOutputWithContext(context.Background())
}

func (i ItsmReceiverResponseArray) ToItsmReceiverResponseArrayOutputWithContext(ctx context.Context) ItsmReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ItsmReceiverResponseArrayOutput)
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





type LocationThresholdRuleConditionInput interface {
	pulumi.Input

	ToLocationThresholdRuleConditionOutput() LocationThresholdRuleConditionOutput
	ToLocationThresholdRuleConditionOutputWithContext(context.Context) LocationThresholdRuleConditionOutput
}

type LocationThresholdRuleConditionArgs struct {
	DataSource          pulumi.Input          `pulumi:"dataSource"`
	FailedLocationCount pulumi.IntInput       `pulumi:"failedLocationCount"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	WindowSize          pulumi.StringPtrInput `pulumi:"windowSize"`
}

func (LocationThresholdRuleConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationThresholdRuleCondition)(nil)).Elem()
}

func (i LocationThresholdRuleConditionArgs) ToLocationThresholdRuleConditionOutput() LocationThresholdRuleConditionOutput {
	return i.ToLocationThresholdRuleConditionOutputWithContext(context.Background())
}

func (i LocationThresholdRuleConditionArgs) ToLocationThresholdRuleConditionOutputWithContext(ctx context.Context) LocationThresholdRuleConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationThresholdRuleConditionOutput)
}

type LocationThresholdRuleConditionOutput struct{ *pulumi.OutputState }

func (LocationThresholdRuleConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationThresholdRuleCondition)(nil)).Elem()
}

func (o LocationThresholdRuleConditionOutput) ToLocationThresholdRuleConditionOutput() LocationThresholdRuleConditionOutput {
	return o
}

func (o LocationThresholdRuleConditionOutput) ToLocationThresholdRuleConditionOutputWithContext(ctx context.Context) LocationThresholdRuleConditionOutput {
	return o
}

func (o LocationThresholdRuleConditionOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v LocationThresholdRuleCondition) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o LocationThresholdRuleConditionOutput) FailedLocationCount() pulumi.IntOutput {
	return o.ApplyT(func(v LocationThresholdRuleCondition) int { return v.FailedLocationCount }).(pulumi.IntOutput)
}

func (o LocationThresholdRuleConditionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v LocationThresholdRuleCondition) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o LocationThresholdRuleConditionOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationThresholdRuleCondition) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

type LocationThresholdRuleConditionResponse struct {
	DataSource          interface{} `pulumi:"dataSource"`
	FailedLocationCount int         `pulumi:"failedLocationCount"`
	OdataType           string      `pulumi:"odataType"`
	WindowSize          *string     `pulumi:"windowSize"`
}





type LocationThresholdRuleConditionResponseInput interface {
	pulumi.Input

	ToLocationThresholdRuleConditionResponseOutput() LocationThresholdRuleConditionResponseOutput
	ToLocationThresholdRuleConditionResponseOutputWithContext(context.Context) LocationThresholdRuleConditionResponseOutput
}

type LocationThresholdRuleConditionResponseArgs struct {
	DataSource          pulumi.Input          `pulumi:"dataSource"`
	FailedLocationCount pulumi.IntInput       `pulumi:"failedLocationCount"`
	OdataType           pulumi.StringInput    `pulumi:"odataType"`
	WindowSize          pulumi.StringPtrInput `pulumi:"windowSize"`
}

func (LocationThresholdRuleConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationThresholdRuleConditionResponse)(nil)).Elem()
}

func (i LocationThresholdRuleConditionResponseArgs) ToLocationThresholdRuleConditionResponseOutput() LocationThresholdRuleConditionResponseOutput {
	return i.ToLocationThresholdRuleConditionResponseOutputWithContext(context.Background())
}

func (i LocationThresholdRuleConditionResponseArgs) ToLocationThresholdRuleConditionResponseOutputWithContext(ctx context.Context) LocationThresholdRuleConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationThresholdRuleConditionResponseOutput)
}

type LocationThresholdRuleConditionResponseOutput struct{ *pulumi.OutputState }

func (LocationThresholdRuleConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationThresholdRuleConditionResponse)(nil)).Elem()
}

func (o LocationThresholdRuleConditionResponseOutput) ToLocationThresholdRuleConditionResponseOutput() LocationThresholdRuleConditionResponseOutput {
	return o
}

func (o LocationThresholdRuleConditionResponseOutput) ToLocationThresholdRuleConditionResponseOutputWithContext(ctx context.Context) LocationThresholdRuleConditionResponseOutput {
	return o
}

func (o LocationThresholdRuleConditionResponseOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v LocationThresholdRuleConditionResponse) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o LocationThresholdRuleConditionResponseOutput) FailedLocationCount() pulumi.IntOutput {
	return o.ApplyT(func(v LocationThresholdRuleConditionResponse) int { return v.FailedLocationCount }).(pulumi.IntOutput)
}

func (o LocationThresholdRuleConditionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v LocationThresholdRuleConditionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o LocationThresholdRuleConditionResponseOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationThresholdRuleConditionResponse) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
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





type LogAnalyticsDestinationResponseInput interface {
	pulumi.Input

	ToLogAnalyticsDestinationResponseOutput() LogAnalyticsDestinationResponseOutput
	ToLogAnalyticsDestinationResponseOutputWithContext(context.Context) LogAnalyticsDestinationResponseOutput
}

type LogAnalyticsDestinationResponseArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	WorkspaceId         pulumi.StringInput    `pulumi:"workspaceId"`
	WorkspaceResourceId pulumi.StringPtrInput `pulumi:"workspaceResourceId"`
}

func (LogAnalyticsDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsDestinationResponse)(nil)).Elem()
}

func (i LogAnalyticsDestinationResponseArgs) ToLogAnalyticsDestinationResponseOutput() LogAnalyticsDestinationResponseOutput {
	return i.ToLogAnalyticsDestinationResponseOutputWithContext(context.Background())
}

func (i LogAnalyticsDestinationResponseArgs) ToLogAnalyticsDestinationResponseOutputWithContext(ctx context.Context) LogAnalyticsDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsDestinationResponseOutput)
}





type LogAnalyticsDestinationResponseArrayInput interface {
	pulumi.Input

	ToLogAnalyticsDestinationResponseArrayOutput() LogAnalyticsDestinationResponseArrayOutput
	ToLogAnalyticsDestinationResponseArrayOutputWithContext(context.Context) LogAnalyticsDestinationResponseArrayOutput
}

type LogAnalyticsDestinationResponseArray []LogAnalyticsDestinationResponseInput

func (LogAnalyticsDestinationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogAnalyticsDestinationResponse)(nil)).Elem()
}

func (i LogAnalyticsDestinationResponseArray) ToLogAnalyticsDestinationResponseArrayOutput() LogAnalyticsDestinationResponseArrayOutput {
	return i.ToLogAnalyticsDestinationResponseArrayOutputWithContext(context.Background())
}

func (i LogAnalyticsDestinationResponseArray) ToLogAnalyticsDestinationResponseArrayOutputWithContext(ctx context.Context) LogAnalyticsDestinationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsDestinationResponseArrayOutput)
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





type LogMetricTriggerInput interface {
	pulumi.Input

	ToLogMetricTriggerOutput() LogMetricTriggerOutput
	ToLogMetricTriggerOutputWithContext(context.Context) LogMetricTriggerOutput
}

type LogMetricTriggerArgs struct {
	MetricColumn      pulumi.StringPtrInput  `pulumi:"metricColumn"`
	MetricTriggerType pulumi.StringPtrInput  `pulumi:"metricTriggerType"`
	Threshold         pulumi.Float64PtrInput `pulumi:"threshold"`
	ThresholdOperator pulumi.StringPtrInput  `pulumi:"thresholdOperator"`
}

func (LogMetricTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMetricTrigger)(nil)).Elem()
}

func (i LogMetricTriggerArgs) ToLogMetricTriggerOutput() LogMetricTriggerOutput {
	return i.ToLogMetricTriggerOutputWithContext(context.Background())
}

func (i LogMetricTriggerArgs) ToLogMetricTriggerOutputWithContext(ctx context.Context) LogMetricTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricTriggerOutput)
}

func (i LogMetricTriggerArgs) ToLogMetricTriggerPtrOutput() LogMetricTriggerPtrOutput {
	return i.ToLogMetricTriggerPtrOutputWithContext(context.Background())
}

func (i LogMetricTriggerArgs) ToLogMetricTriggerPtrOutputWithContext(ctx context.Context) LogMetricTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricTriggerOutput).ToLogMetricTriggerPtrOutputWithContext(ctx)
}









type LogMetricTriggerPtrInput interface {
	pulumi.Input

	ToLogMetricTriggerPtrOutput() LogMetricTriggerPtrOutput
	ToLogMetricTriggerPtrOutputWithContext(context.Context) LogMetricTriggerPtrOutput
}

type logMetricTriggerPtrType LogMetricTriggerArgs

func LogMetricTriggerPtr(v *LogMetricTriggerArgs) LogMetricTriggerPtrInput {
	return (*logMetricTriggerPtrType)(v)
}

func (*logMetricTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMetricTrigger)(nil)).Elem()
}

func (i *logMetricTriggerPtrType) ToLogMetricTriggerPtrOutput() LogMetricTriggerPtrOutput {
	return i.ToLogMetricTriggerPtrOutputWithContext(context.Background())
}

func (i *logMetricTriggerPtrType) ToLogMetricTriggerPtrOutputWithContext(ctx context.Context) LogMetricTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricTriggerPtrOutput)
}

type LogMetricTriggerOutput struct{ *pulumi.OutputState }

func (LogMetricTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMetricTrigger)(nil)).Elem()
}

func (o LogMetricTriggerOutput) ToLogMetricTriggerOutput() LogMetricTriggerOutput {
	return o
}

func (o LogMetricTriggerOutput) ToLogMetricTriggerOutputWithContext(ctx context.Context) LogMetricTriggerOutput {
	return o
}

func (o LogMetricTriggerOutput) ToLogMetricTriggerPtrOutput() LogMetricTriggerPtrOutput {
	return o.ToLogMetricTriggerPtrOutputWithContext(context.Background())
}

func (o LogMetricTriggerOutput) ToLogMetricTriggerPtrOutputWithContext(ctx context.Context) LogMetricTriggerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogMetricTrigger) *LogMetricTrigger {
		return &v
	}).(LogMetricTriggerPtrOutput)
}

func (o LogMetricTriggerOutput) MetricColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogMetricTrigger) *string { return v.MetricColumn }).(pulumi.StringPtrOutput)
}

func (o LogMetricTriggerOutput) MetricTriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogMetricTrigger) *string { return v.MetricTriggerType }).(pulumi.StringPtrOutput)
}

func (o LogMetricTriggerOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LogMetricTrigger) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

func (o LogMetricTriggerOutput) ThresholdOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogMetricTrigger) *string { return v.ThresholdOperator }).(pulumi.StringPtrOutput)
}

type LogMetricTriggerPtrOutput struct{ *pulumi.OutputState }

func (LogMetricTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMetricTrigger)(nil)).Elem()
}

func (o LogMetricTriggerPtrOutput) ToLogMetricTriggerPtrOutput() LogMetricTriggerPtrOutput {
	return o
}

func (o LogMetricTriggerPtrOutput) ToLogMetricTriggerPtrOutputWithContext(ctx context.Context) LogMetricTriggerPtrOutput {
	return o
}

func (o LogMetricTriggerPtrOutput) Elem() LogMetricTriggerOutput {
	return o.ApplyT(func(v *LogMetricTrigger) LogMetricTrigger {
		if v != nil {
			return *v
		}
		var ret LogMetricTrigger
		return ret
	}).(LogMetricTriggerOutput)
}

func (o LogMetricTriggerPtrOutput) MetricColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogMetricTrigger) *string {
		if v == nil {
			return nil
		}
		return v.MetricColumn
	}).(pulumi.StringPtrOutput)
}

func (o LogMetricTriggerPtrOutput) MetricTriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogMetricTrigger) *string {
		if v == nil {
			return nil
		}
		return v.MetricTriggerType
	}).(pulumi.StringPtrOutput)
}

func (o LogMetricTriggerPtrOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LogMetricTrigger) *float64 {
		if v == nil {
			return nil
		}
		return v.Threshold
	}).(pulumi.Float64PtrOutput)
}

func (o LogMetricTriggerPtrOutput) ThresholdOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogMetricTrigger) *string {
		if v == nil {
			return nil
		}
		return v.ThresholdOperator
	}).(pulumi.StringPtrOutput)
}

type LogMetricTriggerResponse struct {
	MetricColumn      *string  `pulumi:"metricColumn"`
	MetricTriggerType *string  `pulumi:"metricTriggerType"`
	Threshold         *float64 `pulumi:"threshold"`
	ThresholdOperator *string  `pulumi:"thresholdOperator"`
}





type LogMetricTriggerResponseInput interface {
	pulumi.Input

	ToLogMetricTriggerResponseOutput() LogMetricTriggerResponseOutput
	ToLogMetricTriggerResponseOutputWithContext(context.Context) LogMetricTriggerResponseOutput
}

type LogMetricTriggerResponseArgs struct {
	MetricColumn      pulumi.StringPtrInput  `pulumi:"metricColumn"`
	MetricTriggerType pulumi.StringPtrInput  `pulumi:"metricTriggerType"`
	Threshold         pulumi.Float64PtrInput `pulumi:"threshold"`
	ThresholdOperator pulumi.StringPtrInput  `pulumi:"thresholdOperator"`
}

func (LogMetricTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMetricTriggerResponse)(nil)).Elem()
}

func (i LogMetricTriggerResponseArgs) ToLogMetricTriggerResponseOutput() LogMetricTriggerResponseOutput {
	return i.ToLogMetricTriggerResponseOutputWithContext(context.Background())
}

func (i LogMetricTriggerResponseArgs) ToLogMetricTriggerResponseOutputWithContext(ctx context.Context) LogMetricTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricTriggerResponseOutput)
}

func (i LogMetricTriggerResponseArgs) ToLogMetricTriggerResponsePtrOutput() LogMetricTriggerResponsePtrOutput {
	return i.ToLogMetricTriggerResponsePtrOutputWithContext(context.Background())
}

func (i LogMetricTriggerResponseArgs) ToLogMetricTriggerResponsePtrOutputWithContext(ctx context.Context) LogMetricTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricTriggerResponseOutput).ToLogMetricTriggerResponsePtrOutputWithContext(ctx)
}









type LogMetricTriggerResponsePtrInput interface {
	pulumi.Input

	ToLogMetricTriggerResponsePtrOutput() LogMetricTriggerResponsePtrOutput
	ToLogMetricTriggerResponsePtrOutputWithContext(context.Context) LogMetricTriggerResponsePtrOutput
}

type logMetricTriggerResponsePtrType LogMetricTriggerResponseArgs

func LogMetricTriggerResponsePtr(v *LogMetricTriggerResponseArgs) LogMetricTriggerResponsePtrInput {
	return (*logMetricTriggerResponsePtrType)(v)
}

func (*logMetricTriggerResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMetricTriggerResponse)(nil)).Elem()
}

func (i *logMetricTriggerResponsePtrType) ToLogMetricTriggerResponsePtrOutput() LogMetricTriggerResponsePtrOutput {
	return i.ToLogMetricTriggerResponsePtrOutputWithContext(context.Background())
}

func (i *logMetricTriggerResponsePtrType) ToLogMetricTriggerResponsePtrOutputWithContext(ctx context.Context) LogMetricTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogMetricTriggerResponsePtrOutput)
}

type LogMetricTriggerResponseOutput struct{ *pulumi.OutputState }

func (LogMetricTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogMetricTriggerResponse)(nil)).Elem()
}

func (o LogMetricTriggerResponseOutput) ToLogMetricTriggerResponseOutput() LogMetricTriggerResponseOutput {
	return o
}

func (o LogMetricTriggerResponseOutput) ToLogMetricTriggerResponseOutputWithContext(ctx context.Context) LogMetricTriggerResponseOutput {
	return o
}

func (o LogMetricTriggerResponseOutput) ToLogMetricTriggerResponsePtrOutput() LogMetricTriggerResponsePtrOutput {
	return o.ToLogMetricTriggerResponsePtrOutputWithContext(context.Background())
}

func (o LogMetricTriggerResponseOutput) ToLogMetricTriggerResponsePtrOutputWithContext(ctx context.Context) LogMetricTriggerResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogMetricTriggerResponse) *LogMetricTriggerResponse {
		return &v
	}).(LogMetricTriggerResponsePtrOutput)
}

func (o LogMetricTriggerResponseOutput) MetricColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogMetricTriggerResponse) *string { return v.MetricColumn }).(pulumi.StringPtrOutput)
}

func (o LogMetricTriggerResponseOutput) MetricTriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogMetricTriggerResponse) *string { return v.MetricTriggerType }).(pulumi.StringPtrOutput)
}

func (o LogMetricTriggerResponseOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LogMetricTriggerResponse) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

func (o LogMetricTriggerResponseOutput) ThresholdOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogMetricTriggerResponse) *string { return v.ThresholdOperator }).(pulumi.StringPtrOutput)
}

type LogMetricTriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (LogMetricTriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogMetricTriggerResponse)(nil)).Elem()
}

func (o LogMetricTriggerResponsePtrOutput) ToLogMetricTriggerResponsePtrOutput() LogMetricTriggerResponsePtrOutput {
	return o
}

func (o LogMetricTriggerResponsePtrOutput) ToLogMetricTriggerResponsePtrOutputWithContext(ctx context.Context) LogMetricTriggerResponsePtrOutput {
	return o
}

func (o LogMetricTriggerResponsePtrOutput) Elem() LogMetricTriggerResponseOutput {
	return o.ApplyT(func(v *LogMetricTriggerResponse) LogMetricTriggerResponse {
		if v != nil {
			return *v
		}
		var ret LogMetricTriggerResponse
		return ret
	}).(LogMetricTriggerResponseOutput)
}

func (o LogMetricTriggerResponsePtrOutput) MetricColumn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogMetricTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.MetricColumn
	}).(pulumi.StringPtrOutput)
}

func (o LogMetricTriggerResponsePtrOutput) MetricTriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogMetricTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.MetricTriggerType
	}).(pulumi.StringPtrOutput)
}

func (o LogMetricTriggerResponsePtrOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LogMetricTriggerResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Threshold
	}).(pulumi.Float64PtrOutput)
}

func (o LogMetricTriggerResponsePtrOutput) ThresholdOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogMetricTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.ThresholdOperator
	}).(pulumi.StringPtrOutput)
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





type LogSettingsResponseInput interface {
	pulumi.Input

	ToLogSettingsResponseOutput() LogSettingsResponseOutput
	ToLogSettingsResponseOutputWithContext(context.Context) LogSettingsResponseOutput
}

type LogSettingsResponseArgs struct {
	Category        pulumi.StringPtrInput           `pulumi:"category"`
	Enabled         pulumi.BoolInput                `pulumi:"enabled"`
	RetentionPolicy RetentionPolicyResponsePtrInput `pulumi:"retentionPolicy"`
}

func (LogSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettingsResponse)(nil)).Elem()
}

func (i LogSettingsResponseArgs) ToLogSettingsResponseOutput() LogSettingsResponseOutput {
	return i.ToLogSettingsResponseOutputWithContext(context.Background())
}

func (i LogSettingsResponseArgs) ToLogSettingsResponseOutputWithContext(ctx context.Context) LogSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingsResponseOutput)
}





type LogSettingsResponseArrayInput interface {
	pulumi.Input

	ToLogSettingsResponseArrayOutput() LogSettingsResponseArrayOutput
	ToLogSettingsResponseArrayOutputWithContext(context.Context) LogSettingsResponseArrayOutput
}

type LogSettingsResponseArray []LogSettingsResponseInput

func (LogSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettingsResponse)(nil)).Elem()
}

func (i LogSettingsResponseArray) ToLogSettingsResponseArrayOutput() LogSettingsResponseArrayOutput {
	return i.ToLogSettingsResponseArrayOutputWithContext(context.Background())
}

func (i LogSettingsResponseArray) ToLogSettingsResponseArrayOutputWithContext(ctx context.Context) LogSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingsResponseArrayOutput)
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





type LogToMetricActionInput interface {
	pulumi.Input

	ToLogToMetricActionOutput() LogToMetricActionOutput
	ToLogToMetricActionOutputWithContext(context.Context) LogToMetricActionOutput
}

type LogToMetricActionArgs struct {
	Criteria  CriteriaArrayInput `pulumi:"criteria"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (LogToMetricActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogToMetricAction)(nil)).Elem()
}

func (i LogToMetricActionArgs) ToLogToMetricActionOutput() LogToMetricActionOutput {
	return i.ToLogToMetricActionOutputWithContext(context.Background())
}

func (i LogToMetricActionArgs) ToLogToMetricActionOutputWithContext(ctx context.Context) LogToMetricActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogToMetricActionOutput)
}

type LogToMetricActionOutput struct{ *pulumi.OutputState }

func (LogToMetricActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogToMetricAction)(nil)).Elem()
}

func (o LogToMetricActionOutput) ToLogToMetricActionOutput() LogToMetricActionOutput {
	return o
}

func (o LogToMetricActionOutput) ToLogToMetricActionOutputWithContext(ctx context.Context) LogToMetricActionOutput {
	return o
}

func (o LogToMetricActionOutput) Criteria() CriteriaArrayOutput {
	return o.ApplyT(func(v LogToMetricAction) []Criteria { return v.Criteria }).(CriteriaArrayOutput)
}

func (o LogToMetricActionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v LogToMetricAction) string { return v.OdataType }).(pulumi.StringOutput)
}

type LogToMetricActionResponse struct {
	Criteria  []CriteriaResponse `pulumi:"criteria"`
	OdataType string             `pulumi:"odataType"`
}





type LogToMetricActionResponseInput interface {
	pulumi.Input

	ToLogToMetricActionResponseOutput() LogToMetricActionResponseOutput
	ToLogToMetricActionResponseOutputWithContext(context.Context) LogToMetricActionResponseOutput
}

type LogToMetricActionResponseArgs struct {
	Criteria  CriteriaResponseArrayInput `pulumi:"criteria"`
	OdataType pulumi.StringInput         `pulumi:"odataType"`
}

func (LogToMetricActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogToMetricActionResponse)(nil)).Elem()
}

func (i LogToMetricActionResponseArgs) ToLogToMetricActionResponseOutput() LogToMetricActionResponseOutput {
	return i.ToLogToMetricActionResponseOutputWithContext(context.Background())
}

func (i LogToMetricActionResponseArgs) ToLogToMetricActionResponseOutputWithContext(ctx context.Context) LogToMetricActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogToMetricActionResponseOutput)
}

type LogToMetricActionResponseOutput struct{ *pulumi.OutputState }

func (LogToMetricActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogToMetricActionResponse)(nil)).Elem()
}

func (o LogToMetricActionResponseOutput) ToLogToMetricActionResponseOutput() LogToMetricActionResponseOutput {
	return o
}

func (o LogToMetricActionResponseOutput) ToLogToMetricActionResponseOutputWithContext(ctx context.Context) LogToMetricActionResponseOutput {
	return o
}

func (o LogToMetricActionResponseOutput) Criteria() CriteriaResponseArrayOutput {
	return o.ApplyT(func(v LogToMetricActionResponse) []CriteriaResponse { return v.Criteria }).(CriteriaResponseArrayOutput)
}

func (o LogToMetricActionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v LogToMetricActionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type LogicAppReceiver struct {
	CallbackUrl          string `pulumi:"callbackUrl"`
	Name                 string `pulumi:"name"`
	ResourceId           string `pulumi:"resourceId"`
	UseCommonAlertSchema *bool  `pulumi:"useCommonAlertSchema"`
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





type LogicAppReceiverResponseInput interface {
	pulumi.Input

	ToLogicAppReceiverResponseOutput() LogicAppReceiverResponseOutput
	ToLogicAppReceiverResponseOutputWithContext(context.Context) LogicAppReceiverResponseOutput
}

type LogicAppReceiverResponseArgs struct {
	CallbackUrl          pulumi.StringInput  `pulumi:"callbackUrl"`
	Name                 pulumi.StringInput  `pulumi:"name"`
	ResourceId           pulumi.StringInput  `pulumi:"resourceId"`
	UseCommonAlertSchema pulumi.BoolPtrInput `pulumi:"useCommonAlertSchema"`
}

func (LogicAppReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogicAppReceiverResponse)(nil)).Elem()
}

func (i LogicAppReceiverResponseArgs) ToLogicAppReceiverResponseOutput() LogicAppReceiverResponseOutput {
	return i.ToLogicAppReceiverResponseOutputWithContext(context.Background())
}

func (i LogicAppReceiverResponseArgs) ToLogicAppReceiverResponseOutputWithContext(ctx context.Context) LogicAppReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogicAppReceiverResponseOutput)
}





type LogicAppReceiverResponseArrayInput interface {
	pulumi.Input

	ToLogicAppReceiverResponseArrayOutput() LogicAppReceiverResponseArrayOutput
	ToLogicAppReceiverResponseArrayOutputWithContext(context.Context) LogicAppReceiverResponseArrayOutput
}

type LogicAppReceiverResponseArray []LogicAppReceiverResponseInput

func (LogicAppReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogicAppReceiverResponse)(nil)).Elem()
}

func (i LogicAppReceiverResponseArray) ToLogicAppReceiverResponseArrayOutput() LogicAppReceiverResponseArrayOutput {
	return i.ToLogicAppReceiverResponseArrayOutputWithContext(context.Background())
}

func (i LogicAppReceiverResponseArray) ToLogicAppReceiverResponseArrayOutputWithContext(ctx context.Context) LogicAppReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogicAppReceiverResponseArrayOutput)
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





type ManagementEventAggregationConditionInput interface {
	pulumi.Input

	ToManagementEventAggregationConditionOutput() ManagementEventAggregationConditionOutput
	ToManagementEventAggregationConditionOutputWithContext(context.Context) ManagementEventAggregationConditionOutput
}

type ManagementEventAggregationConditionArgs struct {
	Operator   ConditionOperatorPtrInput `pulumi:"operator"`
	Threshold  pulumi.Float64PtrInput    `pulumi:"threshold"`
	WindowSize pulumi.StringPtrInput     `pulumi:"windowSize"`
}

func (ManagementEventAggregationConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventAggregationCondition)(nil)).Elem()
}

func (i ManagementEventAggregationConditionArgs) ToManagementEventAggregationConditionOutput() ManagementEventAggregationConditionOutput {
	return i.ToManagementEventAggregationConditionOutputWithContext(context.Background())
}

func (i ManagementEventAggregationConditionArgs) ToManagementEventAggregationConditionOutputWithContext(ctx context.Context) ManagementEventAggregationConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionOutput)
}

func (i ManagementEventAggregationConditionArgs) ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput {
	return i.ToManagementEventAggregationConditionPtrOutputWithContext(context.Background())
}

func (i ManagementEventAggregationConditionArgs) ToManagementEventAggregationConditionPtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionOutput).ToManagementEventAggregationConditionPtrOutputWithContext(ctx)
}









type ManagementEventAggregationConditionPtrInput interface {
	pulumi.Input

	ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput
	ToManagementEventAggregationConditionPtrOutputWithContext(context.Context) ManagementEventAggregationConditionPtrOutput
}

type managementEventAggregationConditionPtrType ManagementEventAggregationConditionArgs

func ManagementEventAggregationConditionPtr(v *ManagementEventAggregationConditionArgs) ManagementEventAggregationConditionPtrInput {
	return (*managementEventAggregationConditionPtrType)(v)
}

func (*managementEventAggregationConditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementEventAggregationCondition)(nil)).Elem()
}

func (i *managementEventAggregationConditionPtrType) ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput {
	return i.ToManagementEventAggregationConditionPtrOutputWithContext(context.Background())
}

func (i *managementEventAggregationConditionPtrType) ToManagementEventAggregationConditionPtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionPtrOutput)
}

type ManagementEventAggregationConditionOutput struct{ *pulumi.OutputState }

func (ManagementEventAggregationConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventAggregationCondition)(nil)).Elem()
}

func (o ManagementEventAggregationConditionOutput) ToManagementEventAggregationConditionOutput() ManagementEventAggregationConditionOutput {
	return o
}

func (o ManagementEventAggregationConditionOutput) ToManagementEventAggregationConditionOutputWithContext(ctx context.Context) ManagementEventAggregationConditionOutput {
	return o
}

func (o ManagementEventAggregationConditionOutput) ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput {
	return o.ToManagementEventAggregationConditionPtrOutputWithContext(context.Background())
}

func (o ManagementEventAggregationConditionOutput) ToManagementEventAggregationConditionPtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementEventAggregationCondition) *ManagementEventAggregationCondition {
		return &v
	}).(ManagementEventAggregationConditionPtrOutput)
}

func (o ManagementEventAggregationConditionOutput) Operator() ConditionOperatorPtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationCondition) *ConditionOperator { return v.Operator }).(ConditionOperatorPtrOutput)
}

func (o ManagementEventAggregationConditionOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationCondition) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

func (o ManagementEventAggregationConditionOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationCondition) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

type ManagementEventAggregationConditionPtrOutput struct{ *pulumi.OutputState }

func (ManagementEventAggregationConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementEventAggregationCondition)(nil)).Elem()
}

func (o ManagementEventAggregationConditionPtrOutput) ToManagementEventAggregationConditionPtrOutput() ManagementEventAggregationConditionPtrOutput {
	return o
}

func (o ManagementEventAggregationConditionPtrOutput) ToManagementEventAggregationConditionPtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionPtrOutput {
	return o
}

func (o ManagementEventAggregationConditionPtrOutput) Elem() ManagementEventAggregationConditionOutput {
	return o.ApplyT(func(v *ManagementEventAggregationCondition) ManagementEventAggregationCondition {
		if v != nil {
			return *v
		}
		var ret ManagementEventAggregationCondition
		return ret
	}).(ManagementEventAggregationConditionOutput)
}

func (o ManagementEventAggregationConditionPtrOutput) Operator() ConditionOperatorPtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationCondition) *ConditionOperator {
		if v == nil {
			return nil
		}
		return v.Operator
	}).(ConditionOperatorPtrOutput)
}

func (o ManagementEventAggregationConditionPtrOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationCondition) *float64 {
		if v == nil {
			return nil
		}
		return v.Threshold
	}).(pulumi.Float64PtrOutput)
}

func (o ManagementEventAggregationConditionPtrOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationCondition) *string {
		if v == nil {
			return nil
		}
		return v.WindowSize
	}).(pulumi.StringPtrOutput)
}

type ManagementEventAggregationConditionResponse struct {
	Operator   *string  `pulumi:"operator"`
	Threshold  *float64 `pulumi:"threshold"`
	WindowSize *string  `pulumi:"windowSize"`
}





type ManagementEventAggregationConditionResponseInput interface {
	pulumi.Input

	ToManagementEventAggregationConditionResponseOutput() ManagementEventAggregationConditionResponseOutput
	ToManagementEventAggregationConditionResponseOutputWithContext(context.Context) ManagementEventAggregationConditionResponseOutput
}

type ManagementEventAggregationConditionResponseArgs struct {
	Operator   pulumi.StringPtrInput  `pulumi:"operator"`
	Threshold  pulumi.Float64PtrInput `pulumi:"threshold"`
	WindowSize pulumi.StringPtrInput  `pulumi:"windowSize"`
}

func (ManagementEventAggregationConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventAggregationConditionResponse)(nil)).Elem()
}

func (i ManagementEventAggregationConditionResponseArgs) ToManagementEventAggregationConditionResponseOutput() ManagementEventAggregationConditionResponseOutput {
	return i.ToManagementEventAggregationConditionResponseOutputWithContext(context.Background())
}

func (i ManagementEventAggregationConditionResponseArgs) ToManagementEventAggregationConditionResponseOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionResponseOutput)
}

func (i ManagementEventAggregationConditionResponseArgs) ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput {
	return i.ToManagementEventAggregationConditionResponsePtrOutputWithContext(context.Background())
}

func (i ManagementEventAggregationConditionResponseArgs) ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionResponseOutput).ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx)
}









type ManagementEventAggregationConditionResponsePtrInput interface {
	pulumi.Input

	ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput
	ToManagementEventAggregationConditionResponsePtrOutputWithContext(context.Context) ManagementEventAggregationConditionResponsePtrOutput
}

type managementEventAggregationConditionResponsePtrType ManagementEventAggregationConditionResponseArgs

func ManagementEventAggregationConditionResponsePtr(v *ManagementEventAggregationConditionResponseArgs) ManagementEventAggregationConditionResponsePtrInput {
	return (*managementEventAggregationConditionResponsePtrType)(v)
}

func (*managementEventAggregationConditionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementEventAggregationConditionResponse)(nil)).Elem()
}

func (i *managementEventAggregationConditionResponsePtrType) ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput {
	return i.ToManagementEventAggregationConditionResponsePtrOutputWithContext(context.Background())
}

func (i *managementEventAggregationConditionResponsePtrType) ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventAggregationConditionResponsePtrOutput)
}

type ManagementEventAggregationConditionResponseOutput struct{ *pulumi.OutputState }

func (ManagementEventAggregationConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventAggregationConditionResponse)(nil)).Elem()
}

func (o ManagementEventAggregationConditionResponseOutput) ToManagementEventAggregationConditionResponseOutput() ManagementEventAggregationConditionResponseOutput {
	return o
}

func (o ManagementEventAggregationConditionResponseOutput) ToManagementEventAggregationConditionResponseOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponseOutput {
	return o
}

func (o ManagementEventAggregationConditionResponseOutput) ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput {
	return o.ToManagementEventAggregationConditionResponsePtrOutputWithContext(context.Background())
}

func (o ManagementEventAggregationConditionResponseOutput) ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagementEventAggregationConditionResponse) *ManagementEventAggregationConditionResponse {
		return &v
	}).(ManagementEventAggregationConditionResponsePtrOutput)
}

func (o ManagementEventAggregationConditionResponseOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationConditionResponse) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o ManagementEventAggregationConditionResponseOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationConditionResponse) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

func (o ManagementEventAggregationConditionResponseOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementEventAggregationConditionResponse) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

type ManagementEventAggregationConditionResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementEventAggregationConditionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementEventAggregationConditionResponse)(nil)).Elem()
}

func (o ManagementEventAggregationConditionResponsePtrOutput) ToManagementEventAggregationConditionResponsePtrOutput() ManagementEventAggregationConditionResponsePtrOutput {
	return o
}

func (o ManagementEventAggregationConditionResponsePtrOutput) ToManagementEventAggregationConditionResponsePtrOutputWithContext(ctx context.Context) ManagementEventAggregationConditionResponsePtrOutput {
	return o
}

func (o ManagementEventAggregationConditionResponsePtrOutput) Elem() ManagementEventAggregationConditionResponseOutput {
	return o.ApplyT(func(v *ManagementEventAggregationConditionResponse) ManagementEventAggregationConditionResponse {
		if v != nil {
			return *v
		}
		var ret ManagementEventAggregationConditionResponse
		return ret
	}).(ManagementEventAggregationConditionResponseOutput)
}

func (o ManagementEventAggregationConditionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationConditionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ManagementEventAggregationConditionResponsePtrOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationConditionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Threshold
	}).(pulumi.Float64PtrOutput)
}

func (o ManagementEventAggregationConditionResponsePtrOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementEventAggregationConditionResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowSize
	}).(pulumi.StringPtrOutput)
}

type ManagementEventRuleCondition struct {
	Aggregation *ManagementEventAggregationCondition `pulumi:"aggregation"`
	DataSource  interface{}                          `pulumi:"dataSource"`
	OdataType   string                               `pulumi:"odataType"`
}





type ManagementEventRuleConditionInput interface {
	pulumi.Input

	ToManagementEventRuleConditionOutput() ManagementEventRuleConditionOutput
	ToManagementEventRuleConditionOutputWithContext(context.Context) ManagementEventRuleConditionOutput
}

type ManagementEventRuleConditionArgs struct {
	Aggregation ManagementEventAggregationConditionPtrInput `pulumi:"aggregation"`
	DataSource  pulumi.Input                                `pulumi:"dataSource"`
	OdataType   pulumi.StringInput                          `pulumi:"odataType"`
}

func (ManagementEventRuleConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventRuleCondition)(nil)).Elem()
}

func (i ManagementEventRuleConditionArgs) ToManagementEventRuleConditionOutput() ManagementEventRuleConditionOutput {
	return i.ToManagementEventRuleConditionOutputWithContext(context.Background())
}

func (i ManagementEventRuleConditionArgs) ToManagementEventRuleConditionOutputWithContext(ctx context.Context) ManagementEventRuleConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventRuleConditionOutput)
}

type ManagementEventRuleConditionOutput struct{ *pulumi.OutputState }

func (ManagementEventRuleConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventRuleCondition)(nil)).Elem()
}

func (o ManagementEventRuleConditionOutput) ToManagementEventRuleConditionOutput() ManagementEventRuleConditionOutput {
	return o
}

func (o ManagementEventRuleConditionOutput) ToManagementEventRuleConditionOutputWithContext(ctx context.Context) ManagementEventRuleConditionOutput {
	return o
}

func (o ManagementEventRuleConditionOutput) Aggregation() ManagementEventAggregationConditionPtrOutput {
	return o.ApplyT(func(v ManagementEventRuleCondition) *ManagementEventAggregationCondition { return v.Aggregation }).(ManagementEventAggregationConditionPtrOutput)
}

func (o ManagementEventRuleConditionOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v ManagementEventRuleCondition) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o ManagementEventRuleConditionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementEventRuleCondition) string { return v.OdataType }).(pulumi.StringOutput)
}

type ManagementEventRuleConditionResponse struct {
	Aggregation *ManagementEventAggregationConditionResponse `pulumi:"aggregation"`
	DataSource  interface{}                                  `pulumi:"dataSource"`
	OdataType   string                                       `pulumi:"odataType"`
}





type ManagementEventRuleConditionResponseInput interface {
	pulumi.Input

	ToManagementEventRuleConditionResponseOutput() ManagementEventRuleConditionResponseOutput
	ToManagementEventRuleConditionResponseOutputWithContext(context.Context) ManagementEventRuleConditionResponseOutput
}

type ManagementEventRuleConditionResponseArgs struct {
	Aggregation ManagementEventAggregationConditionResponsePtrInput `pulumi:"aggregation"`
	DataSource  pulumi.Input                                        `pulumi:"dataSource"`
	OdataType   pulumi.StringInput                                  `pulumi:"odataType"`
}

func (ManagementEventRuleConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventRuleConditionResponse)(nil)).Elem()
}

func (i ManagementEventRuleConditionResponseArgs) ToManagementEventRuleConditionResponseOutput() ManagementEventRuleConditionResponseOutput {
	return i.ToManagementEventRuleConditionResponseOutputWithContext(context.Background())
}

func (i ManagementEventRuleConditionResponseArgs) ToManagementEventRuleConditionResponseOutputWithContext(ctx context.Context) ManagementEventRuleConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementEventRuleConditionResponseOutput)
}

type ManagementEventRuleConditionResponseOutput struct{ *pulumi.OutputState }

func (ManagementEventRuleConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementEventRuleConditionResponse)(nil)).Elem()
}

func (o ManagementEventRuleConditionResponseOutput) ToManagementEventRuleConditionResponseOutput() ManagementEventRuleConditionResponseOutput {
	return o
}

func (o ManagementEventRuleConditionResponseOutput) ToManagementEventRuleConditionResponseOutputWithContext(ctx context.Context) ManagementEventRuleConditionResponseOutput {
	return o
}

func (o ManagementEventRuleConditionResponseOutput) Aggregation() ManagementEventAggregationConditionResponsePtrOutput {
	return o.ApplyT(func(v ManagementEventRuleConditionResponse) *ManagementEventAggregationConditionResponse {
		return v.Aggregation
	}).(ManagementEventAggregationConditionResponsePtrOutput)
}

func (o ManagementEventRuleConditionResponseOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v ManagementEventRuleConditionResponse) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o ManagementEventRuleConditionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementEventRuleConditionResponse) string { return v.OdataType }).(pulumi.StringOutput)
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





type ManagementGroupLogSettingsResponseInput interface {
	pulumi.Input

	ToManagementGroupLogSettingsResponseOutput() ManagementGroupLogSettingsResponseOutput
	ToManagementGroupLogSettingsResponseOutputWithContext(context.Context) ManagementGroupLogSettingsResponseOutput
}

type ManagementGroupLogSettingsResponseArgs struct {
	Category pulumi.StringInput `pulumi:"category"`
	Enabled  pulumi.BoolInput   `pulumi:"enabled"`
}

func (ManagementGroupLogSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupLogSettingsResponse)(nil)).Elem()
}

func (i ManagementGroupLogSettingsResponseArgs) ToManagementGroupLogSettingsResponseOutput() ManagementGroupLogSettingsResponseOutput {
	return i.ToManagementGroupLogSettingsResponseOutputWithContext(context.Background())
}

func (i ManagementGroupLogSettingsResponseArgs) ToManagementGroupLogSettingsResponseOutputWithContext(ctx context.Context) ManagementGroupLogSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupLogSettingsResponseOutput)
}





type ManagementGroupLogSettingsResponseArrayInput interface {
	pulumi.Input

	ToManagementGroupLogSettingsResponseArrayOutput() ManagementGroupLogSettingsResponseArrayOutput
	ToManagementGroupLogSettingsResponseArrayOutputWithContext(context.Context) ManagementGroupLogSettingsResponseArrayOutput
}

type ManagementGroupLogSettingsResponseArray []ManagementGroupLogSettingsResponseInput

func (ManagementGroupLogSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupLogSettingsResponse)(nil)).Elem()
}

func (i ManagementGroupLogSettingsResponseArray) ToManagementGroupLogSettingsResponseArrayOutput() ManagementGroupLogSettingsResponseArrayOutput {
	return i.ToManagementGroupLogSettingsResponseArrayOutputWithContext(context.Background())
}

func (i ManagementGroupLogSettingsResponseArray) ToManagementGroupLogSettingsResponseArrayOutputWithContext(ctx context.Context) ManagementGroupLogSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupLogSettingsResponseArrayOutput)
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





type MetricAlertActionResponseInput interface {
	pulumi.Input

	ToMetricAlertActionResponseOutput() MetricAlertActionResponseOutput
	ToMetricAlertActionResponseOutputWithContext(context.Context) MetricAlertActionResponseOutput
}

type MetricAlertActionResponseArgs struct {
	ActionGroupId     pulumi.StringPtrInput `pulumi:"actionGroupId"`
	WebHookProperties pulumi.StringMapInput `pulumi:"webHookProperties"`
}

func (MetricAlertActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertActionResponse)(nil)).Elem()
}

func (i MetricAlertActionResponseArgs) ToMetricAlertActionResponseOutput() MetricAlertActionResponseOutput {
	return i.ToMetricAlertActionResponseOutputWithContext(context.Background())
}

func (i MetricAlertActionResponseArgs) ToMetricAlertActionResponseOutputWithContext(ctx context.Context) MetricAlertActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertActionResponseOutput)
}





type MetricAlertActionResponseArrayInput interface {
	pulumi.Input

	ToMetricAlertActionResponseArrayOutput() MetricAlertActionResponseArrayOutput
	ToMetricAlertActionResponseArrayOutputWithContext(context.Context) MetricAlertActionResponseArrayOutput
}

type MetricAlertActionResponseArray []MetricAlertActionResponseInput

func (MetricAlertActionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricAlertActionResponse)(nil)).Elem()
}

func (i MetricAlertActionResponseArray) ToMetricAlertActionResponseArrayOutput() MetricAlertActionResponseArrayOutput {
	return i.ToMetricAlertActionResponseArrayOutputWithContext(context.Background())
}

func (i MetricAlertActionResponseArray) ToMetricAlertActionResponseArrayOutputWithContext(ctx context.Context) MetricAlertActionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertActionResponseArrayOutput)
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





type MetricAlertMultipleResourceMultipleMetricCriteriaInput interface {
	pulumi.Input

	ToMetricAlertMultipleResourceMultipleMetricCriteriaOutput() MetricAlertMultipleResourceMultipleMetricCriteriaOutput
	ToMetricAlertMultipleResourceMultipleMetricCriteriaOutputWithContext(context.Context) MetricAlertMultipleResourceMultipleMetricCriteriaOutput
}

type MetricAlertMultipleResourceMultipleMetricCriteriaArgs struct {
	AllOf     pulumi.ArrayInput  `pulumi:"allOf"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (MetricAlertMultipleResourceMultipleMetricCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertMultipleResourceMultipleMetricCriteria)(nil)).Elem()
}

func (i MetricAlertMultipleResourceMultipleMetricCriteriaArgs) ToMetricAlertMultipleResourceMultipleMetricCriteriaOutput() MetricAlertMultipleResourceMultipleMetricCriteriaOutput {
	return i.ToMetricAlertMultipleResourceMultipleMetricCriteriaOutputWithContext(context.Background())
}

func (i MetricAlertMultipleResourceMultipleMetricCriteriaArgs) ToMetricAlertMultipleResourceMultipleMetricCriteriaOutputWithContext(ctx context.Context) MetricAlertMultipleResourceMultipleMetricCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertMultipleResourceMultipleMetricCriteriaOutput)
}

type MetricAlertMultipleResourceMultipleMetricCriteriaOutput struct{ *pulumi.OutputState }

func (MetricAlertMultipleResourceMultipleMetricCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertMultipleResourceMultipleMetricCriteria)(nil)).Elem()
}

func (o MetricAlertMultipleResourceMultipleMetricCriteriaOutput) ToMetricAlertMultipleResourceMultipleMetricCriteriaOutput() MetricAlertMultipleResourceMultipleMetricCriteriaOutput {
	return o
}

func (o MetricAlertMultipleResourceMultipleMetricCriteriaOutput) ToMetricAlertMultipleResourceMultipleMetricCriteriaOutputWithContext(ctx context.Context) MetricAlertMultipleResourceMultipleMetricCriteriaOutput {
	return o
}

func (o MetricAlertMultipleResourceMultipleMetricCriteriaOutput) AllOf() pulumi.ArrayOutput {
	return o.ApplyT(func(v MetricAlertMultipleResourceMultipleMetricCriteria) []interface{} { return v.AllOf }).(pulumi.ArrayOutput)
}

func (o MetricAlertMultipleResourceMultipleMetricCriteriaOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MetricAlertMultipleResourceMultipleMetricCriteria) string { return v.OdataType }).(pulumi.StringOutput)
}

type MetricAlertMultipleResourceMultipleMetricCriteriaResponse struct {
	AllOf     []interface{} `pulumi:"allOf"`
	OdataType string        `pulumi:"odataType"`
}





type MetricAlertMultipleResourceMultipleMetricCriteriaResponseInput interface {
	pulumi.Input

	ToMetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput() MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput
	ToMetricAlertMultipleResourceMultipleMetricCriteriaResponseOutputWithContext(context.Context) MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput
}

type MetricAlertMultipleResourceMultipleMetricCriteriaResponseArgs struct {
	AllOf     pulumi.ArrayInput  `pulumi:"allOf"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (MetricAlertMultipleResourceMultipleMetricCriteriaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertMultipleResourceMultipleMetricCriteriaResponse)(nil)).Elem()
}

func (i MetricAlertMultipleResourceMultipleMetricCriteriaResponseArgs) ToMetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput() MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput {
	return i.ToMetricAlertMultipleResourceMultipleMetricCriteriaResponseOutputWithContext(context.Background())
}

func (i MetricAlertMultipleResourceMultipleMetricCriteriaResponseArgs) ToMetricAlertMultipleResourceMultipleMetricCriteriaResponseOutputWithContext(ctx context.Context) MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput)
}

type MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput struct{ *pulumi.OutputState }

func (MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertMultipleResourceMultipleMetricCriteriaResponse)(nil)).Elem()
}

func (o MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput) ToMetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput() MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput {
	return o
}

func (o MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput) ToMetricAlertMultipleResourceMultipleMetricCriteriaResponseOutputWithContext(ctx context.Context) MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput {
	return o
}

func (o MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput) AllOf() pulumi.ArrayOutput {
	return o.ApplyT(func(v MetricAlertMultipleResourceMultipleMetricCriteriaResponse) []interface{} { return v.AllOf }).(pulumi.ArrayOutput)
}

func (o MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MetricAlertMultipleResourceMultipleMetricCriteriaResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

type MetricAlertSingleResourceMultipleMetricCriteria struct {
	AllOf     []MetricCriteria `pulumi:"allOf"`
	OdataType string           `pulumi:"odataType"`
}





type MetricAlertSingleResourceMultipleMetricCriteriaInput interface {
	pulumi.Input

	ToMetricAlertSingleResourceMultipleMetricCriteriaOutput() MetricAlertSingleResourceMultipleMetricCriteriaOutput
	ToMetricAlertSingleResourceMultipleMetricCriteriaOutputWithContext(context.Context) MetricAlertSingleResourceMultipleMetricCriteriaOutput
}

type MetricAlertSingleResourceMultipleMetricCriteriaArgs struct {
	AllOf     MetricCriteriaArrayInput `pulumi:"allOf"`
	OdataType pulumi.StringInput       `pulumi:"odataType"`
}

func (MetricAlertSingleResourceMultipleMetricCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertSingleResourceMultipleMetricCriteria)(nil)).Elem()
}

func (i MetricAlertSingleResourceMultipleMetricCriteriaArgs) ToMetricAlertSingleResourceMultipleMetricCriteriaOutput() MetricAlertSingleResourceMultipleMetricCriteriaOutput {
	return i.ToMetricAlertSingleResourceMultipleMetricCriteriaOutputWithContext(context.Background())
}

func (i MetricAlertSingleResourceMultipleMetricCriteriaArgs) ToMetricAlertSingleResourceMultipleMetricCriteriaOutputWithContext(ctx context.Context) MetricAlertSingleResourceMultipleMetricCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertSingleResourceMultipleMetricCriteriaOutput)
}

type MetricAlertSingleResourceMultipleMetricCriteriaOutput struct{ *pulumi.OutputState }

func (MetricAlertSingleResourceMultipleMetricCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertSingleResourceMultipleMetricCriteria)(nil)).Elem()
}

func (o MetricAlertSingleResourceMultipleMetricCriteriaOutput) ToMetricAlertSingleResourceMultipleMetricCriteriaOutput() MetricAlertSingleResourceMultipleMetricCriteriaOutput {
	return o
}

func (o MetricAlertSingleResourceMultipleMetricCriteriaOutput) ToMetricAlertSingleResourceMultipleMetricCriteriaOutputWithContext(ctx context.Context) MetricAlertSingleResourceMultipleMetricCriteriaOutput {
	return o
}

func (o MetricAlertSingleResourceMultipleMetricCriteriaOutput) AllOf() MetricCriteriaArrayOutput {
	return o.ApplyT(func(v MetricAlertSingleResourceMultipleMetricCriteria) []MetricCriteria { return v.AllOf }).(MetricCriteriaArrayOutput)
}

func (o MetricAlertSingleResourceMultipleMetricCriteriaOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MetricAlertSingleResourceMultipleMetricCriteria) string { return v.OdataType }).(pulumi.StringOutput)
}

type MetricAlertSingleResourceMultipleMetricCriteriaResponse struct {
	AllOf     []MetricCriteriaResponse `pulumi:"allOf"`
	OdataType string                   `pulumi:"odataType"`
}





type MetricAlertSingleResourceMultipleMetricCriteriaResponseInput interface {
	pulumi.Input

	ToMetricAlertSingleResourceMultipleMetricCriteriaResponseOutput() MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput
	ToMetricAlertSingleResourceMultipleMetricCriteriaResponseOutputWithContext(context.Context) MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput
}

type MetricAlertSingleResourceMultipleMetricCriteriaResponseArgs struct {
	AllOf     MetricCriteriaResponseArrayInput `pulumi:"allOf"`
	OdataType pulumi.StringInput               `pulumi:"odataType"`
}

func (MetricAlertSingleResourceMultipleMetricCriteriaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertSingleResourceMultipleMetricCriteriaResponse)(nil)).Elem()
}

func (i MetricAlertSingleResourceMultipleMetricCriteriaResponseArgs) ToMetricAlertSingleResourceMultipleMetricCriteriaResponseOutput() MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput {
	return i.ToMetricAlertSingleResourceMultipleMetricCriteriaResponseOutputWithContext(context.Background())
}

func (i MetricAlertSingleResourceMultipleMetricCriteriaResponseArgs) ToMetricAlertSingleResourceMultipleMetricCriteriaResponseOutputWithContext(ctx context.Context) MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput)
}

type MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput struct{ *pulumi.OutputState }

func (MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricAlertSingleResourceMultipleMetricCriteriaResponse)(nil)).Elem()
}

func (o MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput) ToMetricAlertSingleResourceMultipleMetricCriteriaResponseOutput() MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput {
	return o
}

func (o MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput) ToMetricAlertSingleResourceMultipleMetricCriteriaResponseOutputWithContext(ctx context.Context) MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput {
	return o
}

func (o MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput) AllOf() MetricCriteriaResponseArrayOutput {
	return o.ApplyT(func(v MetricAlertSingleResourceMultipleMetricCriteriaResponse) []MetricCriteriaResponse {
		return v.AllOf
	}).(MetricCriteriaResponseArrayOutput)
}

func (o MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v MetricAlertSingleResourceMultipleMetricCriteriaResponse) string { return v.OdataType }).(pulumi.StringOutput)
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





type MetricCriteriaInput interface {
	pulumi.Input

	ToMetricCriteriaOutput() MetricCriteriaOutput
	ToMetricCriteriaOutputWithContext(context.Context) MetricCriteriaOutput
}

type MetricCriteriaArgs struct {
	CriterionType        pulumi.StringInput        `pulumi:"criterionType"`
	Dimensions           MetricDimensionArrayInput `pulumi:"dimensions"`
	MetricName           pulumi.StringInput        `pulumi:"metricName"`
	MetricNamespace      pulumi.StringPtrInput     `pulumi:"metricNamespace"`
	Name                 pulumi.StringInput        `pulumi:"name"`
	Operator             pulumi.StringInput        `pulumi:"operator"`
	SkipMetricValidation pulumi.BoolPtrInput       `pulumi:"skipMetricValidation"`
	Threshold            pulumi.Float64Input       `pulumi:"threshold"`
	TimeAggregation      pulumi.StringInput        `pulumi:"timeAggregation"`
}

func (MetricCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCriteria)(nil)).Elem()
}

func (i MetricCriteriaArgs) ToMetricCriteriaOutput() MetricCriteriaOutput {
	return i.ToMetricCriteriaOutputWithContext(context.Background())
}

func (i MetricCriteriaArgs) ToMetricCriteriaOutputWithContext(ctx context.Context) MetricCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCriteriaOutput)
}





type MetricCriteriaArrayInput interface {
	pulumi.Input

	ToMetricCriteriaArrayOutput() MetricCriteriaArrayOutput
	ToMetricCriteriaArrayOutputWithContext(context.Context) MetricCriteriaArrayOutput
}

type MetricCriteriaArray []MetricCriteriaInput

func (MetricCriteriaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCriteria)(nil)).Elem()
}

func (i MetricCriteriaArray) ToMetricCriteriaArrayOutput() MetricCriteriaArrayOutput {
	return i.ToMetricCriteriaArrayOutputWithContext(context.Background())
}

func (i MetricCriteriaArray) ToMetricCriteriaArrayOutputWithContext(ctx context.Context) MetricCriteriaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCriteriaArrayOutput)
}

type MetricCriteriaOutput struct{ *pulumi.OutputState }

func (MetricCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCriteria)(nil)).Elem()
}

func (o MetricCriteriaOutput) ToMetricCriteriaOutput() MetricCriteriaOutput {
	return o
}

func (o MetricCriteriaOutput) ToMetricCriteriaOutputWithContext(ctx context.Context) MetricCriteriaOutput {
	return o
}

func (o MetricCriteriaOutput) CriterionType() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteria) string { return v.CriterionType }).(pulumi.StringOutput)
}

func (o MetricCriteriaOutput) Dimensions() MetricDimensionArrayOutput {
	return o.ApplyT(func(v MetricCriteria) []MetricDimension { return v.Dimensions }).(MetricDimensionArrayOutput)
}

func (o MetricCriteriaOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteria) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o MetricCriteriaOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricCriteria) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o MetricCriteriaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteria) string { return v.Name }).(pulumi.StringOutput)
}

func (o MetricCriteriaOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteria) string { return v.Operator }).(pulumi.StringOutput)
}

func (o MetricCriteriaOutput) SkipMetricValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MetricCriteria) *bool { return v.SkipMetricValidation }).(pulumi.BoolPtrOutput)
}

func (o MetricCriteriaOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v MetricCriteria) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o MetricCriteriaOutput) TimeAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteria) string { return v.TimeAggregation }).(pulumi.StringOutput)
}

type MetricCriteriaArrayOutput struct{ *pulumi.OutputState }

func (MetricCriteriaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCriteria)(nil)).Elem()
}

func (o MetricCriteriaArrayOutput) ToMetricCriteriaArrayOutput() MetricCriteriaArrayOutput {
	return o
}

func (o MetricCriteriaArrayOutput) ToMetricCriteriaArrayOutputWithContext(ctx context.Context) MetricCriteriaArrayOutput {
	return o
}

func (o MetricCriteriaArrayOutput) Index(i pulumi.IntInput) MetricCriteriaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricCriteria {
		return vs[0].([]MetricCriteria)[vs[1].(int)]
	}).(MetricCriteriaOutput)
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





type MetricCriteriaResponseInput interface {
	pulumi.Input

	ToMetricCriteriaResponseOutput() MetricCriteriaResponseOutput
	ToMetricCriteriaResponseOutputWithContext(context.Context) MetricCriteriaResponseOutput
}

type MetricCriteriaResponseArgs struct {
	CriterionType        pulumi.StringInput                `pulumi:"criterionType"`
	Dimensions           MetricDimensionResponseArrayInput `pulumi:"dimensions"`
	MetricName           pulumi.StringInput                `pulumi:"metricName"`
	MetricNamespace      pulumi.StringPtrInput             `pulumi:"metricNamespace"`
	Name                 pulumi.StringInput                `pulumi:"name"`
	Operator             pulumi.StringInput                `pulumi:"operator"`
	SkipMetricValidation pulumi.BoolPtrInput               `pulumi:"skipMetricValidation"`
	Threshold            pulumi.Float64Input               `pulumi:"threshold"`
	TimeAggregation      pulumi.StringInput                `pulumi:"timeAggregation"`
}

func (MetricCriteriaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCriteriaResponse)(nil)).Elem()
}

func (i MetricCriteriaResponseArgs) ToMetricCriteriaResponseOutput() MetricCriteriaResponseOutput {
	return i.ToMetricCriteriaResponseOutputWithContext(context.Background())
}

func (i MetricCriteriaResponseArgs) ToMetricCriteriaResponseOutputWithContext(ctx context.Context) MetricCriteriaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCriteriaResponseOutput)
}





type MetricCriteriaResponseArrayInput interface {
	pulumi.Input

	ToMetricCriteriaResponseArrayOutput() MetricCriteriaResponseArrayOutput
	ToMetricCriteriaResponseArrayOutputWithContext(context.Context) MetricCriteriaResponseArrayOutput
}

type MetricCriteriaResponseArray []MetricCriteriaResponseInput

func (MetricCriteriaResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCriteriaResponse)(nil)).Elem()
}

func (i MetricCriteriaResponseArray) ToMetricCriteriaResponseArrayOutput() MetricCriteriaResponseArrayOutput {
	return i.ToMetricCriteriaResponseArrayOutputWithContext(context.Background())
}

func (i MetricCriteriaResponseArray) ToMetricCriteriaResponseArrayOutputWithContext(ctx context.Context) MetricCriteriaResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricCriteriaResponseArrayOutput)
}

type MetricCriteriaResponseOutput struct{ *pulumi.OutputState }

func (MetricCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricCriteriaResponse)(nil)).Elem()
}

func (o MetricCriteriaResponseOutput) ToMetricCriteriaResponseOutput() MetricCriteriaResponseOutput {
	return o
}

func (o MetricCriteriaResponseOutput) ToMetricCriteriaResponseOutputWithContext(ctx context.Context) MetricCriteriaResponseOutput {
	return o
}

func (o MetricCriteriaResponseOutput) CriterionType() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteriaResponse) string { return v.CriterionType }).(pulumi.StringOutput)
}

func (o MetricCriteriaResponseOutput) Dimensions() MetricDimensionResponseArrayOutput {
	return o.ApplyT(func(v MetricCriteriaResponse) []MetricDimensionResponse { return v.Dimensions }).(MetricDimensionResponseArrayOutput)
}

func (o MetricCriteriaResponseOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteriaResponse) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o MetricCriteriaResponseOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricCriteriaResponse) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o MetricCriteriaResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteriaResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MetricCriteriaResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteriaResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o MetricCriteriaResponseOutput) SkipMetricValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MetricCriteriaResponse) *bool { return v.SkipMetricValidation }).(pulumi.BoolPtrOutput)
}

func (o MetricCriteriaResponseOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v MetricCriteriaResponse) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o MetricCriteriaResponseOutput) TimeAggregation() pulumi.StringOutput {
	return o.ApplyT(func(v MetricCriteriaResponse) string { return v.TimeAggregation }).(pulumi.StringOutput)
}

type MetricCriteriaResponseArrayOutput struct{ *pulumi.OutputState }

func (MetricCriteriaResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricCriteriaResponse)(nil)).Elem()
}

func (o MetricCriteriaResponseArrayOutput) ToMetricCriteriaResponseArrayOutput() MetricCriteriaResponseArrayOutput {
	return o
}

func (o MetricCriteriaResponseArrayOutput) ToMetricCriteriaResponseArrayOutputWithContext(ctx context.Context) MetricCriteriaResponseArrayOutput {
	return o
}

func (o MetricCriteriaResponseArrayOutput) Index(i pulumi.IntInput) MetricCriteriaResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricCriteriaResponse {
		return vs[0].([]MetricCriteriaResponse)[vs[1].(int)]
	}).(MetricCriteriaResponseOutput)
}

type MetricDimension struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type MetricDimensionInput interface {
	pulumi.Input

	ToMetricDimensionOutput() MetricDimensionOutput
	ToMetricDimensionOutputWithContext(context.Context) MetricDimensionOutput
}

type MetricDimensionArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (MetricDimensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricDimension)(nil)).Elem()
}

func (i MetricDimensionArgs) ToMetricDimensionOutput() MetricDimensionOutput {
	return i.ToMetricDimensionOutputWithContext(context.Background())
}

func (i MetricDimensionArgs) ToMetricDimensionOutputWithContext(ctx context.Context) MetricDimensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricDimensionOutput)
}





type MetricDimensionArrayInput interface {
	pulumi.Input

	ToMetricDimensionArrayOutput() MetricDimensionArrayOutput
	ToMetricDimensionArrayOutputWithContext(context.Context) MetricDimensionArrayOutput
}

type MetricDimensionArray []MetricDimensionInput

func (MetricDimensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricDimension)(nil)).Elem()
}

func (i MetricDimensionArray) ToMetricDimensionArrayOutput() MetricDimensionArrayOutput {
	return i.ToMetricDimensionArrayOutputWithContext(context.Background())
}

func (i MetricDimensionArray) ToMetricDimensionArrayOutputWithContext(ctx context.Context) MetricDimensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricDimensionArrayOutput)
}

type MetricDimensionOutput struct{ *pulumi.OutputState }

func (MetricDimensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricDimension)(nil)).Elem()
}

func (o MetricDimensionOutput) ToMetricDimensionOutput() MetricDimensionOutput {
	return o
}

func (o MetricDimensionOutput) ToMetricDimensionOutputWithContext(ctx context.Context) MetricDimensionOutput {
	return o
}

func (o MetricDimensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MetricDimension) string { return v.Name }).(pulumi.StringOutput)
}

func (o MetricDimensionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v MetricDimension) string { return v.Operator }).(pulumi.StringOutput)
}

func (o MetricDimensionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MetricDimension) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type MetricDimensionArrayOutput struct{ *pulumi.OutputState }

func (MetricDimensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricDimension)(nil)).Elem()
}

func (o MetricDimensionArrayOutput) ToMetricDimensionArrayOutput() MetricDimensionArrayOutput {
	return o
}

func (o MetricDimensionArrayOutput) ToMetricDimensionArrayOutputWithContext(ctx context.Context) MetricDimensionArrayOutput {
	return o
}

func (o MetricDimensionArrayOutput) Index(i pulumi.IntInput) MetricDimensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricDimension {
		return vs[0].([]MetricDimension)[vs[1].(int)]
	}).(MetricDimensionOutput)
}

type MetricDimensionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type MetricDimensionResponseInput interface {
	pulumi.Input

	ToMetricDimensionResponseOutput() MetricDimensionResponseOutput
	ToMetricDimensionResponseOutputWithContext(context.Context) MetricDimensionResponseOutput
}

type MetricDimensionResponseArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (MetricDimensionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricDimensionResponse)(nil)).Elem()
}

func (i MetricDimensionResponseArgs) ToMetricDimensionResponseOutput() MetricDimensionResponseOutput {
	return i.ToMetricDimensionResponseOutputWithContext(context.Background())
}

func (i MetricDimensionResponseArgs) ToMetricDimensionResponseOutputWithContext(ctx context.Context) MetricDimensionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricDimensionResponseOutput)
}





type MetricDimensionResponseArrayInput interface {
	pulumi.Input

	ToMetricDimensionResponseArrayOutput() MetricDimensionResponseArrayOutput
	ToMetricDimensionResponseArrayOutputWithContext(context.Context) MetricDimensionResponseArrayOutput
}

type MetricDimensionResponseArray []MetricDimensionResponseInput

func (MetricDimensionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricDimensionResponse)(nil)).Elem()
}

func (i MetricDimensionResponseArray) ToMetricDimensionResponseArrayOutput() MetricDimensionResponseArrayOutput {
	return i.ToMetricDimensionResponseArrayOutputWithContext(context.Background())
}

func (i MetricDimensionResponseArray) ToMetricDimensionResponseArrayOutputWithContext(ctx context.Context) MetricDimensionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricDimensionResponseArrayOutput)
}

type MetricDimensionResponseOutput struct{ *pulumi.OutputState }

func (MetricDimensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricDimensionResponse)(nil)).Elem()
}

func (o MetricDimensionResponseOutput) ToMetricDimensionResponseOutput() MetricDimensionResponseOutput {
	return o
}

func (o MetricDimensionResponseOutput) ToMetricDimensionResponseOutputWithContext(ctx context.Context) MetricDimensionResponseOutput {
	return o
}

func (o MetricDimensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MetricDimensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o MetricDimensionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v MetricDimensionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o MetricDimensionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MetricDimensionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type MetricDimensionResponseArrayOutput struct{ *pulumi.OutputState }

func (MetricDimensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricDimensionResponse)(nil)).Elem()
}

func (o MetricDimensionResponseArrayOutput) ToMetricDimensionResponseArrayOutput() MetricDimensionResponseArrayOutput {
	return o
}

func (o MetricDimensionResponseArrayOutput) ToMetricDimensionResponseArrayOutputWithContext(ctx context.Context) MetricDimensionResponseArrayOutput {
	return o
}

func (o MetricDimensionResponseArrayOutput) Index(i pulumi.IntInput) MetricDimensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricDimensionResponse {
		return vs[0].([]MetricDimensionResponse)[vs[1].(int)]
	}).(MetricDimensionResponseOutput)
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





type MetricSettingsResponseInput interface {
	pulumi.Input

	ToMetricSettingsResponseOutput() MetricSettingsResponseOutput
	ToMetricSettingsResponseOutputWithContext(context.Context) MetricSettingsResponseOutput
}

type MetricSettingsResponseArgs struct {
	Category        pulumi.StringPtrInput           `pulumi:"category"`
	Enabled         pulumi.BoolInput                `pulumi:"enabled"`
	RetentionPolicy RetentionPolicyResponsePtrInput `pulumi:"retentionPolicy"`
	TimeGrain       pulumi.StringPtrInput           `pulumi:"timeGrain"`
}

func (MetricSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricSettingsResponse)(nil)).Elem()
}

func (i MetricSettingsResponseArgs) ToMetricSettingsResponseOutput() MetricSettingsResponseOutput {
	return i.ToMetricSettingsResponseOutputWithContext(context.Background())
}

func (i MetricSettingsResponseArgs) ToMetricSettingsResponseOutputWithContext(ctx context.Context) MetricSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricSettingsResponseOutput)
}





type MetricSettingsResponseArrayInput interface {
	pulumi.Input

	ToMetricSettingsResponseArrayOutput() MetricSettingsResponseArrayOutput
	ToMetricSettingsResponseArrayOutputWithContext(context.Context) MetricSettingsResponseArrayOutput
}

type MetricSettingsResponseArray []MetricSettingsResponseInput

func (MetricSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricSettingsResponse)(nil)).Elem()
}

func (i MetricSettingsResponseArray) ToMetricSettingsResponseArrayOutput() MetricSettingsResponseArrayOutput {
	return i.ToMetricSettingsResponseArrayOutputWithContext(context.Background())
}

func (i MetricSettingsResponseArray) ToMetricSettingsResponseArrayOutputWithContext(ctx context.Context) MetricSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricSettingsResponseArrayOutput)
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





type MetricTriggerResponseInput interface {
	pulumi.Input

	ToMetricTriggerResponseOutput() MetricTriggerResponseOutput
	ToMetricTriggerResponseOutputWithContext(context.Context) MetricTriggerResponseOutput
}

type MetricTriggerResponseArgs struct {
	Dimensions             ScaleRuleMetricDimensionResponseArrayInput `pulumi:"dimensions"`
	DividePerInstance      pulumi.BoolPtrInput                        `pulumi:"dividePerInstance"`
	MetricName             pulumi.StringInput                         `pulumi:"metricName"`
	MetricNamespace        pulumi.StringPtrInput                      `pulumi:"metricNamespace"`
	MetricResourceLocation pulumi.StringPtrInput                      `pulumi:"metricResourceLocation"`
	MetricResourceUri      pulumi.StringInput                         `pulumi:"metricResourceUri"`
	Operator               pulumi.StringInput                         `pulumi:"operator"`
	Statistic              pulumi.StringInput                         `pulumi:"statistic"`
	Threshold              pulumi.Float64Input                        `pulumi:"threshold"`
	TimeAggregation        pulumi.StringInput                         `pulumi:"timeAggregation"`
	TimeGrain              pulumi.StringInput                         `pulumi:"timeGrain"`
	TimeWindow             pulumi.StringInput                         `pulumi:"timeWindow"`
}

func (MetricTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricTriggerResponse)(nil)).Elem()
}

func (i MetricTriggerResponseArgs) ToMetricTriggerResponseOutput() MetricTriggerResponseOutput {
	return i.ToMetricTriggerResponseOutputWithContext(context.Background())
}

func (i MetricTriggerResponseArgs) ToMetricTriggerResponseOutputWithContext(ctx context.Context) MetricTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricTriggerResponseOutput)
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





type MyWorkbookManagedIdentityResponseInput interface {
	pulumi.Input

	ToMyWorkbookManagedIdentityResponseOutput() MyWorkbookManagedIdentityResponseOutput
	ToMyWorkbookManagedIdentityResponseOutputWithContext(context.Context) MyWorkbookManagedIdentityResponseOutput
}

type MyWorkbookManagedIdentityResponseArgs struct {
	Type                   pulumi.StringPtrInput                            `pulumi:"type"`
	UserAssignedIdentities MyWorkbookUserAssignedIdentitiesResponsePtrInput `pulumi:"userAssignedIdentities"`
}

func (MyWorkbookManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookManagedIdentityResponse)(nil)).Elem()
}

func (i MyWorkbookManagedIdentityResponseArgs) ToMyWorkbookManagedIdentityResponseOutput() MyWorkbookManagedIdentityResponseOutput {
	return i.ToMyWorkbookManagedIdentityResponseOutputWithContext(context.Background())
}

func (i MyWorkbookManagedIdentityResponseArgs) ToMyWorkbookManagedIdentityResponseOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityResponseOutput)
}

func (i MyWorkbookManagedIdentityResponseArgs) ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput {
	return i.ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i MyWorkbookManagedIdentityResponseArgs) ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityResponseOutput).ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx)
}









type MyWorkbookManagedIdentityResponsePtrInput interface {
	pulumi.Input

	ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput
	ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(context.Context) MyWorkbookManagedIdentityResponsePtrOutput
}

type myWorkbookManagedIdentityResponsePtrType MyWorkbookManagedIdentityResponseArgs

func MyWorkbookManagedIdentityResponsePtr(v *MyWorkbookManagedIdentityResponseArgs) MyWorkbookManagedIdentityResponsePtrInput {
	return (*myWorkbookManagedIdentityResponsePtrType)(v)
}

func (*myWorkbookManagedIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookManagedIdentityResponse)(nil)).Elem()
}

func (i *myWorkbookManagedIdentityResponsePtrType) ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput {
	return i.ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *myWorkbookManagedIdentityResponsePtrType) ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityResponsePtrOutput)
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

func (o MyWorkbookManagedIdentityResponseOutput) ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput {
	return o.ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (o MyWorkbookManagedIdentityResponseOutput) ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MyWorkbookManagedIdentityResponse) *MyWorkbookManagedIdentityResponse {
		return &v
	}).(MyWorkbookManagedIdentityResponsePtrOutput)
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





type MyWorkbookUserAssignedIdentitiesResponseInput interface {
	pulumi.Input

	ToMyWorkbookUserAssignedIdentitiesResponseOutput() MyWorkbookUserAssignedIdentitiesResponseOutput
	ToMyWorkbookUserAssignedIdentitiesResponseOutputWithContext(context.Context) MyWorkbookUserAssignedIdentitiesResponseOutput
}

type MyWorkbookUserAssignedIdentitiesResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
}

func (MyWorkbookUserAssignedIdentitiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (i MyWorkbookUserAssignedIdentitiesResponseArgs) ToMyWorkbookUserAssignedIdentitiesResponseOutput() MyWorkbookUserAssignedIdentitiesResponseOutput {
	return i.ToMyWorkbookUserAssignedIdentitiesResponseOutputWithContext(context.Background())
}

func (i MyWorkbookUserAssignedIdentitiesResponseArgs) ToMyWorkbookUserAssignedIdentitiesResponseOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookUserAssignedIdentitiesResponseOutput)
}

func (i MyWorkbookUserAssignedIdentitiesResponseArgs) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return i.ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (i MyWorkbookUserAssignedIdentitiesResponseArgs) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookUserAssignedIdentitiesResponseOutput).ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx)
}









type MyWorkbookUserAssignedIdentitiesResponsePtrInput interface {
	pulumi.Input

	ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput
	ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput
}

type myWorkbookUserAssignedIdentitiesResponsePtrType MyWorkbookUserAssignedIdentitiesResponseArgs

func MyWorkbookUserAssignedIdentitiesResponsePtr(v *MyWorkbookUserAssignedIdentitiesResponseArgs) MyWorkbookUserAssignedIdentitiesResponsePtrInput {
	return (*myWorkbookUserAssignedIdentitiesResponsePtrType)(v)
}

func (*myWorkbookUserAssignedIdentitiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (i *myWorkbookUserAssignedIdentitiesResponsePtrType) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return i.ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (i *myWorkbookUserAssignedIdentitiesResponsePtrType) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookUserAssignedIdentitiesResponsePtrOutput)
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

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MyWorkbookUserAssignedIdentitiesResponse) *MyWorkbookUserAssignedIdentitiesResponse {
		return &v
	}).(MyWorkbookUserAssignedIdentitiesResponsePtrOutput)
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





type PerfCounterDataSourceResponseInput interface {
	pulumi.Input

	ToPerfCounterDataSourceResponseOutput() PerfCounterDataSourceResponseOutput
	ToPerfCounterDataSourceResponseOutputWithContext(context.Context) PerfCounterDataSourceResponseOutput
}

type PerfCounterDataSourceResponseArgs struct {
	CounterSpecifiers          pulumi.StringArrayInput `pulumi:"counterSpecifiers"`
	Name                       pulumi.StringPtrInput   `pulumi:"name"`
	SamplingFrequencyInSeconds pulumi.IntPtrInput      `pulumi:"samplingFrequencyInSeconds"`
	Streams                    pulumi.StringArrayInput `pulumi:"streams"`
}

func (PerfCounterDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerfCounterDataSourceResponse)(nil)).Elem()
}

func (i PerfCounterDataSourceResponseArgs) ToPerfCounterDataSourceResponseOutput() PerfCounterDataSourceResponseOutput {
	return i.ToPerfCounterDataSourceResponseOutputWithContext(context.Background())
}

func (i PerfCounterDataSourceResponseArgs) ToPerfCounterDataSourceResponseOutputWithContext(ctx context.Context) PerfCounterDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerfCounterDataSourceResponseOutput)
}





type PerfCounterDataSourceResponseArrayInput interface {
	pulumi.Input

	ToPerfCounterDataSourceResponseArrayOutput() PerfCounterDataSourceResponseArrayOutput
	ToPerfCounterDataSourceResponseArrayOutputWithContext(context.Context) PerfCounterDataSourceResponseArrayOutput
}

type PerfCounterDataSourceResponseArray []PerfCounterDataSourceResponseInput

func (PerfCounterDataSourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerfCounterDataSourceResponse)(nil)).Elem()
}

func (i PerfCounterDataSourceResponseArray) ToPerfCounterDataSourceResponseArrayOutput() PerfCounterDataSourceResponseArrayOutput {
	return i.ToPerfCounterDataSourceResponseArrayOutputWithContext(context.Background())
}

func (i PerfCounterDataSourceResponseArray) ToPerfCounterDataSourceResponseArrayOutputWithContext(ctx context.Context) PerfCounterDataSourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerfCounterDataSourceResponseArrayOutput)
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





type PerformanceCounterConfigurationResponseInput interface {
	pulumi.Input

	ToPerformanceCounterConfigurationResponseOutput() PerformanceCounterConfigurationResponseOutput
	ToPerformanceCounterConfigurationResponseOutputWithContext(context.Context) PerformanceCounterConfigurationResponseOutput
}

type PerformanceCounterConfigurationResponseArgs struct {
	Instance       pulumi.StringPtrInput `pulumi:"instance"`
	Name           pulumi.StringInput    `pulumi:"name"`
	SamplingPeriod pulumi.StringInput    `pulumi:"samplingPeriod"`
}

func (PerformanceCounterConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PerformanceCounterConfigurationResponse)(nil)).Elem()
}

func (i PerformanceCounterConfigurationResponseArgs) ToPerformanceCounterConfigurationResponseOutput() PerformanceCounterConfigurationResponseOutput {
	return i.ToPerformanceCounterConfigurationResponseOutputWithContext(context.Background())
}

func (i PerformanceCounterConfigurationResponseArgs) ToPerformanceCounterConfigurationResponseOutputWithContext(ctx context.Context) PerformanceCounterConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerformanceCounterConfigurationResponseOutput)
}





type PerformanceCounterConfigurationResponseArrayInput interface {
	pulumi.Input

	ToPerformanceCounterConfigurationResponseArrayOutput() PerformanceCounterConfigurationResponseArrayOutput
	ToPerformanceCounterConfigurationResponseArrayOutputWithContext(context.Context) PerformanceCounterConfigurationResponseArrayOutput
}

type PerformanceCounterConfigurationResponseArray []PerformanceCounterConfigurationResponseInput

func (PerformanceCounterConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PerformanceCounterConfigurationResponse)(nil)).Elem()
}

func (i PerformanceCounterConfigurationResponseArray) ToPerformanceCounterConfigurationResponseArrayOutput() PerformanceCounterConfigurationResponseArrayOutput {
	return i.ToPerformanceCounterConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i PerformanceCounterConfigurationResponseArray) ToPerformanceCounterConfigurationResponseArrayOutputWithContext(ctx context.Context) PerformanceCounterConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PerformanceCounterConfigurationResponseArrayOutput)
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





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                                        `pulumi:"id"`
	Name                              pulumi.StringInput                                        `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointPropertyResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                        `pulumi:"provisioningState"`
	Type                              pulumi.StringInput                                        `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
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





type PrivateEndpointPropertyResponseInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput
	ToPrivateEndpointPropertyResponseOutputWithContext(context.Context) PrivateEndpointPropertyResponseOutput
}

type PrivateEndpointPropertyResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return i.ToPrivateEndpointPropertyResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput)
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput).ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput
	ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Context) PrivateEndpointPropertyResponsePtrOutput
}

type privateEndpointPropertyResponsePtrType PrivateEndpointPropertyResponseArgs

func PrivateEndpointPropertyResponsePtr(v *PrivateEndpointPropertyResponseArgs) PrivateEndpointPropertyResponsePtrInput {
	return (*privateEndpointPropertyResponsePtrType)(v)
}

func (*privateEndpointPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponsePtrOutput)
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

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointPropertyResponse) *PrivateEndpointPropertyResponse {
		return &v
	}).(PrivateEndpointPropertyResponsePtrOutput)
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





type PrivateLinkScopedResourceResponseInput interface {
	pulumi.Input

	ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput
	ToPrivateLinkScopedResourceResponseOutputWithContext(context.Context) PrivateLinkScopedResourceResponseOutput
}

type PrivateLinkScopedResourceResponseArgs struct {
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	ScopeId    pulumi.StringPtrInput `pulumi:"scopeId"`
}

func (PrivateLinkScopedResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (i PrivateLinkScopedResourceResponseArgs) ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput {
	return i.ToPrivateLinkScopedResourceResponseOutputWithContext(context.Background())
}

func (i PrivateLinkScopedResourceResponseArgs) ToPrivateLinkScopedResourceResponseOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopedResourceResponseOutput)
}





type PrivateLinkScopedResourceResponseArrayInput interface {
	pulumi.Input

	ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput
	ToPrivateLinkScopedResourceResponseArrayOutputWithContext(context.Context) PrivateLinkScopedResourceResponseArrayOutput
}

type PrivateLinkScopedResourceResponseArray []PrivateLinkScopedResourceResponseInput

func (PrivateLinkScopedResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (i PrivateLinkScopedResourceResponseArray) ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput {
	return i.ToPrivateLinkScopedResourceResponseArrayOutputWithContext(context.Background())
}

func (i PrivateLinkScopedResourceResponseArray) ToPrivateLinkScopedResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopedResourceResponseArrayOutput)
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





type PrivateLinkServiceConnectionStatePropertyResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput
	ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput
}

type PrivateLinkServiceConnectionStatePropertyResponseArgs struct {
	ActionsRequired pulumi.StringInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput `pulumi:"description"`
	Status          pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput).ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
}

type privateLinkServiceConnectionStatePropertyResponsePtrType PrivateLinkServiceConnectionStatePropertyResponseArgs

func PrivateLinkServiceConnectionStatePropertyResponsePtr(v *PrivateLinkServiceConnectionStatePropertyResponseArgs) PrivateLinkServiceConnectionStatePropertyResponsePtrInput {
	return (*privateLinkServiceConnectionStatePropertyResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
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

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatePropertyResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
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





type RecurrenceResponseInput interface {
	pulumi.Input

	ToRecurrenceResponseOutput() RecurrenceResponseOutput
	ToRecurrenceResponseOutputWithContext(context.Context) RecurrenceResponseOutput
}

type RecurrenceResponseArgs struct {
	Frequency pulumi.StringInput             `pulumi:"frequency"`
	Schedule  RecurrentScheduleResponseInput `pulumi:"schedule"`
}

func (RecurrenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceResponse)(nil)).Elem()
}

func (i RecurrenceResponseArgs) ToRecurrenceResponseOutput() RecurrenceResponseOutput {
	return i.ToRecurrenceResponseOutputWithContext(context.Background())
}

func (i RecurrenceResponseArgs) ToRecurrenceResponseOutputWithContext(ctx context.Context) RecurrenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceResponseOutput)
}

func (i RecurrenceResponseArgs) ToRecurrenceResponsePtrOutput() RecurrenceResponsePtrOutput {
	return i.ToRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (i RecurrenceResponseArgs) ToRecurrenceResponsePtrOutputWithContext(ctx context.Context) RecurrenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceResponseOutput).ToRecurrenceResponsePtrOutputWithContext(ctx)
}









type RecurrenceResponsePtrInput interface {
	pulumi.Input

	ToRecurrenceResponsePtrOutput() RecurrenceResponsePtrOutput
	ToRecurrenceResponsePtrOutputWithContext(context.Context) RecurrenceResponsePtrOutput
}

type recurrenceResponsePtrType RecurrenceResponseArgs

func RecurrenceResponsePtr(v *RecurrenceResponseArgs) RecurrenceResponsePtrInput {
	return (*recurrenceResponsePtrType)(v)
}

func (*recurrenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceResponse)(nil)).Elem()
}

func (i *recurrenceResponsePtrType) ToRecurrenceResponsePtrOutput() RecurrenceResponsePtrOutput {
	return i.ToRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (i *recurrenceResponsePtrType) ToRecurrenceResponsePtrOutputWithContext(ctx context.Context) RecurrenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceResponsePtrOutput)
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

func (o RecurrenceResponseOutput) ToRecurrenceResponsePtrOutput() RecurrenceResponsePtrOutput {
	return o.ToRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (o RecurrenceResponseOutput) ToRecurrenceResponsePtrOutputWithContext(ctx context.Context) RecurrenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrenceResponse) *RecurrenceResponse {
		return &v
	}).(RecurrenceResponsePtrOutput)
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





type RecurrentScheduleResponseInput interface {
	pulumi.Input

	ToRecurrentScheduleResponseOutput() RecurrentScheduleResponseOutput
	ToRecurrentScheduleResponseOutputWithContext(context.Context) RecurrentScheduleResponseOutput
}

type RecurrentScheduleResponseArgs struct {
	Days     pulumi.StringArrayInput `pulumi:"days"`
	Hours    pulumi.IntArrayInput    `pulumi:"hours"`
	Minutes  pulumi.IntArrayInput    `pulumi:"minutes"`
	TimeZone pulumi.StringInput      `pulumi:"timeZone"`
}

func (RecurrentScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrentScheduleResponse)(nil)).Elem()
}

func (i RecurrentScheduleResponseArgs) ToRecurrentScheduleResponseOutput() RecurrentScheduleResponseOutput {
	return i.ToRecurrentScheduleResponseOutputWithContext(context.Background())
}

func (i RecurrentScheduleResponseArgs) ToRecurrentScheduleResponseOutputWithContext(ctx context.Context) RecurrentScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrentScheduleResponseOutput)
}

func (i RecurrentScheduleResponseArgs) ToRecurrentScheduleResponsePtrOutput() RecurrentScheduleResponsePtrOutput {
	return i.ToRecurrentScheduleResponsePtrOutputWithContext(context.Background())
}

func (i RecurrentScheduleResponseArgs) ToRecurrentScheduleResponsePtrOutputWithContext(ctx context.Context) RecurrentScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrentScheduleResponseOutput).ToRecurrentScheduleResponsePtrOutputWithContext(ctx)
}









type RecurrentScheduleResponsePtrInput interface {
	pulumi.Input

	ToRecurrentScheduleResponsePtrOutput() RecurrentScheduleResponsePtrOutput
	ToRecurrentScheduleResponsePtrOutputWithContext(context.Context) RecurrentScheduleResponsePtrOutput
}

type recurrentScheduleResponsePtrType RecurrentScheduleResponseArgs

func RecurrentScheduleResponsePtr(v *RecurrentScheduleResponseArgs) RecurrentScheduleResponsePtrInput {
	return (*recurrentScheduleResponsePtrType)(v)
}

func (*recurrentScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrentScheduleResponse)(nil)).Elem()
}

func (i *recurrentScheduleResponsePtrType) ToRecurrentScheduleResponsePtrOutput() RecurrentScheduleResponsePtrOutput {
	return i.ToRecurrentScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *recurrentScheduleResponsePtrType) ToRecurrentScheduleResponsePtrOutputWithContext(ctx context.Context) RecurrentScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrentScheduleResponsePtrOutput)
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

func (o RecurrentScheduleResponseOutput) ToRecurrentScheduleResponsePtrOutput() RecurrentScheduleResponsePtrOutput {
	return o.ToRecurrentScheduleResponsePtrOutputWithContext(context.Background())
}

func (o RecurrentScheduleResponseOutput) ToRecurrentScheduleResponsePtrOutputWithContext(ctx context.Context) RecurrentScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrentScheduleResponse) *RecurrentScheduleResponse {
		return &v
	}).(RecurrentScheduleResponsePtrOutput)
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





type RetentionPolicyResponseInput interface {
	pulumi.Input

	ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput
	ToRetentionPolicyResponseOutputWithContext(context.Context) RetentionPolicyResponseOutput
}

type RetentionPolicyResponseArgs struct {
	Days    pulumi.IntInput  `pulumi:"days"`
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (RetentionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return i.ToRetentionPolicyResponseOutputWithContext(context.Background())
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyResponseOutput)
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return i.ToRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyResponseOutput).ToRetentionPolicyResponsePtrOutputWithContext(ctx)
}









type RetentionPolicyResponsePtrInput interface {
	pulumi.Input

	ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput
	ToRetentionPolicyResponsePtrOutputWithContext(context.Context) RetentionPolicyResponsePtrOutput
}

type retentionPolicyResponsePtrType RetentionPolicyResponseArgs

func RetentionPolicyResponsePtr(v *RetentionPolicyResponseArgs) RetentionPolicyResponsePtrInput {
	return (*retentionPolicyResponsePtrType)(v)
}

func (*retentionPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (i *retentionPolicyResponsePtrType) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return i.ToRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *retentionPolicyResponsePtrType) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyResponsePtrOutput)
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

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o.ToRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicyResponse) *RetentionPolicyResponse {
		return &v
	}).(RetentionPolicyResponsePtrOutput)
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





type RuleEmailActionInput interface {
	pulumi.Input

	ToRuleEmailActionOutput() RuleEmailActionOutput
	ToRuleEmailActionOutputWithContext(context.Context) RuleEmailActionOutput
}

type RuleEmailActionArgs struct {
	CustomEmails        pulumi.StringArrayInput `pulumi:"customEmails"`
	OdataType           pulumi.StringInput      `pulumi:"odataType"`
	SendToServiceOwners pulumi.BoolPtrInput     `pulumi:"sendToServiceOwners"`
}

func (RuleEmailActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleEmailAction)(nil)).Elem()
}

func (i RuleEmailActionArgs) ToRuleEmailActionOutput() RuleEmailActionOutput {
	return i.ToRuleEmailActionOutputWithContext(context.Background())
}

func (i RuleEmailActionArgs) ToRuleEmailActionOutputWithContext(ctx context.Context) RuleEmailActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleEmailActionOutput)
}

type RuleEmailActionOutput struct{ *pulumi.OutputState }

func (RuleEmailActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleEmailAction)(nil)).Elem()
}

func (o RuleEmailActionOutput) ToRuleEmailActionOutput() RuleEmailActionOutput {
	return o
}

func (o RuleEmailActionOutput) ToRuleEmailActionOutputWithContext(ctx context.Context) RuleEmailActionOutput {
	return o
}

func (o RuleEmailActionOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleEmailAction) []string { return v.CustomEmails }).(pulumi.StringArrayOutput)
}

func (o RuleEmailActionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleEmailAction) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleEmailActionOutput) SendToServiceOwners() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleEmailAction) *bool { return v.SendToServiceOwners }).(pulumi.BoolPtrOutput)
}

type RuleEmailActionResponse struct {
	CustomEmails        []string `pulumi:"customEmails"`
	OdataType           string   `pulumi:"odataType"`
	SendToServiceOwners *bool    `pulumi:"sendToServiceOwners"`
}





type RuleEmailActionResponseInput interface {
	pulumi.Input

	ToRuleEmailActionResponseOutput() RuleEmailActionResponseOutput
	ToRuleEmailActionResponseOutputWithContext(context.Context) RuleEmailActionResponseOutput
}

type RuleEmailActionResponseArgs struct {
	CustomEmails        pulumi.StringArrayInput `pulumi:"customEmails"`
	OdataType           pulumi.StringInput      `pulumi:"odataType"`
	SendToServiceOwners pulumi.BoolPtrInput     `pulumi:"sendToServiceOwners"`
}

func (RuleEmailActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleEmailActionResponse)(nil)).Elem()
}

func (i RuleEmailActionResponseArgs) ToRuleEmailActionResponseOutput() RuleEmailActionResponseOutput {
	return i.ToRuleEmailActionResponseOutputWithContext(context.Background())
}

func (i RuleEmailActionResponseArgs) ToRuleEmailActionResponseOutputWithContext(ctx context.Context) RuleEmailActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleEmailActionResponseOutput)
}

type RuleEmailActionResponseOutput struct{ *pulumi.OutputState }

func (RuleEmailActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleEmailActionResponse)(nil)).Elem()
}

func (o RuleEmailActionResponseOutput) ToRuleEmailActionResponseOutput() RuleEmailActionResponseOutput {
	return o
}

func (o RuleEmailActionResponseOutput) ToRuleEmailActionResponseOutputWithContext(ctx context.Context) RuleEmailActionResponseOutput {
	return o
}

func (o RuleEmailActionResponseOutput) CustomEmails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RuleEmailActionResponse) []string { return v.CustomEmails }).(pulumi.StringArrayOutput)
}

func (o RuleEmailActionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleEmailActionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleEmailActionResponseOutput) SendToServiceOwners() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleEmailActionResponse) *bool { return v.SendToServiceOwners }).(pulumi.BoolPtrOutput)
}

type RuleManagementEventClaimsDataSource struct {
	EmailAddress *string `pulumi:"emailAddress"`
}





type RuleManagementEventClaimsDataSourceInput interface {
	pulumi.Input

	ToRuleManagementEventClaimsDataSourceOutput() RuleManagementEventClaimsDataSourceOutput
	ToRuleManagementEventClaimsDataSourceOutputWithContext(context.Context) RuleManagementEventClaimsDataSourceOutput
}

type RuleManagementEventClaimsDataSourceArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
}

func (RuleManagementEventClaimsDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventClaimsDataSource)(nil)).Elem()
}

func (i RuleManagementEventClaimsDataSourceArgs) ToRuleManagementEventClaimsDataSourceOutput() RuleManagementEventClaimsDataSourceOutput {
	return i.ToRuleManagementEventClaimsDataSourceOutputWithContext(context.Background())
}

func (i RuleManagementEventClaimsDataSourceArgs) ToRuleManagementEventClaimsDataSourceOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceOutput)
}

func (i RuleManagementEventClaimsDataSourceArgs) ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput {
	return i.ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(context.Background())
}

func (i RuleManagementEventClaimsDataSourceArgs) ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceOutput).ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx)
}









type RuleManagementEventClaimsDataSourcePtrInput interface {
	pulumi.Input

	ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput
	ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(context.Context) RuleManagementEventClaimsDataSourcePtrOutput
}

type ruleManagementEventClaimsDataSourcePtrType RuleManagementEventClaimsDataSourceArgs

func RuleManagementEventClaimsDataSourcePtr(v *RuleManagementEventClaimsDataSourceArgs) RuleManagementEventClaimsDataSourcePtrInput {
	return (*ruleManagementEventClaimsDataSourcePtrType)(v)
}

func (*ruleManagementEventClaimsDataSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleManagementEventClaimsDataSource)(nil)).Elem()
}

func (i *ruleManagementEventClaimsDataSourcePtrType) ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput {
	return i.ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(context.Background())
}

func (i *ruleManagementEventClaimsDataSourcePtrType) ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourcePtrOutput)
}

type RuleManagementEventClaimsDataSourceOutput struct{ *pulumi.OutputState }

func (RuleManagementEventClaimsDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventClaimsDataSource)(nil)).Elem()
}

func (o RuleManagementEventClaimsDataSourceOutput) ToRuleManagementEventClaimsDataSourceOutput() RuleManagementEventClaimsDataSourceOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceOutput) ToRuleManagementEventClaimsDataSourceOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceOutput) ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput {
	return o.ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(context.Background())
}

func (o RuleManagementEventClaimsDataSourceOutput) ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleManagementEventClaimsDataSource) *RuleManagementEventClaimsDataSource {
		return &v
	}).(RuleManagementEventClaimsDataSourcePtrOutput)
}

func (o RuleManagementEventClaimsDataSourceOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventClaimsDataSource) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

type RuleManagementEventClaimsDataSourcePtrOutput struct{ *pulumi.OutputState }

func (RuleManagementEventClaimsDataSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleManagementEventClaimsDataSource)(nil)).Elem()
}

func (o RuleManagementEventClaimsDataSourcePtrOutput) ToRuleManagementEventClaimsDataSourcePtrOutput() RuleManagementEventClaimsDataSourcePtrOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourcePtrOutput) ToRuleManagementEventClaimsDataSourcePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourcePtrOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourcePtrOutput) Elem() RuleManagementEventClaimsDataSourceOutput {
	return o.ApplyT(func(v *RuleManagementEventClaimsDataSource) RuleManagementEventClaimsDataSource {
		if v != nil {
			return *v
		}
		var ret RuleManagementEventClaimsDataSource
		return ret
	}).(RuleManagementEventClaimsDataSourceOutput)
}

func (o RuleManagementEventClaimsDataSourcePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleManagementEventClaimsDataSource) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

type RuleManagementEventClaimsDataSourceResponse struct {
	EmailAddress *string `pulumi:"emailAddress"`
}





type RuleManagementEventClaimsDataSourceResponseInput interface {
	pulumi.Input

	ToRuleManagementEventClaimsDataSourceResponseOutput() RuleManagementEventClaimsDataSourceResponseOutput
	ToRuleManagementEventClaimsDataSourceResponseOutputWithContext(context.Context) RuleManagementEventClaimsDataSourceResponseOutput
}

type RuleManagementEventClaimsDataSourceResponseArgs struct {
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
}

func (RuleManagementEventClaimsDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventClaimsDataSourceResponse)(nil)).Elem()
}

func (i RuleManagementEventClaimsDataSourceResponseArgs) ToRuleManagementEventClaimsDataSourceResponseOutput() RuleManagementEventClaimsDataSourceResponseOutput {
	return i.ToRuleManagementEventClaimsDataSourceResponseOutputWithContext(context.Background())
}

func (i RuleManagementEventClaimsDataSourceResponseArgs) ToRuleManagementEventClaimsDataSourceResponseOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceResponseOutput)
}

func (i RuleManagementEventClaimsDataSourceResponseArgs) ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return i.ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(context.Background())
}

func (i RuleManagementEventClaimsDataSourceResponseArgs) ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceResponseOutput).ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx)
}









type RuleManagementEventClaimsDataSourceResponsePtrInput interface {
	pulumi.Input

	ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput
	ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput
}

type ruleManagementEventClaimsDataSourceResponsePtrType RuleManagementEventClaimsDataSourceResponseArgs

func RuleManagementEventClaimsDataSourceResponsePtr(v *RuleManagementEventClaimsDataSourceResponseArgs) RuleManagementEventClaimsDataSourceResponsePtrInput {
	return (*ruleManagementEventClaimsDataSourceResponsePtrType)(v)
}

func (*ruleManagementEventClaimsDataSourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleManagementEventClaimsDataSourceResponse)(nil)).Elem()
}

func (i *ruleManagementEventClaimsDataSourceResponsePtrType) ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return i.ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(context.Background())
}

func (i *ruleManagementEventClaimsDataSourceResponsePtrType) ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventClaimsDataSourceResponsePtrOutput)
}

type RuleManagementEventClaimsDataSourceResponseOutput struct{ *pulumi.OutputState }

func (RuleManagementEventClaimsDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventClaimsDataSourceResponse)(nil)).Elem()
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) ToRuleManagementEventClaimsDataSourceResponseOutput() RuleManagementEventClaimsDataSourceResponseOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) ToRuleManagementEventClaimsDataSourceResponseOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponseOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o.ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(context.Background())
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RuleManagementEventClaimsDataSourceResponse) *RuleManagementEventClaimsDataSourceResponse {
		return &v
	}).(RuleManagementEventClaimsDataSourceResponsePtrOutput)
}

func (o RuleManagementEventClaimsDataSourceResponseOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventClaimsDataSourceResponse) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

type RuleManagementEventClaimsDataSourceResponsePtrOutput struct{ *pulumi.OutputState }

func (RuleManagementEventClaimsDataSourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleManagementEventClaimsDataSourceResponse)(nil)).Elem()
}

func (o RuleManagementEventClaimsDataSourceResponsePtrOutput) ToRuleManagementEventClaimsDataSourceResponsePtrOutput() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceResponsePtrOutput) ToRuleManagementEventClaimsDataSourceResponsePtrOutputWithContext(ctx context.Context) RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o
}

func (o RuleManagementEventClaimsDataSourceResponsePtrOutput) Elem() RuleManagementEventClaimsDataSourceResponseOutput {
	return o.ApplyT(func(v *RuleManagementEventClaimsDataSourceResponse) RuleManagementEventClaimsDataSourceResponse {
		if v != nil {
			return *v
		}
		var ret RuleManagementEventClaimsDataSourceResponse
		return ret
	}).(RuleManagementEventClaimsDataSourceResponseOutput)
}

func (o RuleManagementEventClaimsDataSourceResponsePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleManagementEventClaimsDataSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
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





type RuleManagementEventDataSourceInput interface {
	pulumi.Input

	ToRuleManagementEventDataSourceOutput() RuleManagementEventDataSourceOutput
	ToRuleManagementEventDataSourceOutputWithContext(context.Context) RuleManagementEventDataSourceOutput
}

type RuleManagementEventDataSourceArgs struct {
	Claims               RuleManagementEventClaimsDataSourcePtrInput `pulumi:"claims"`
	EventName            pulumi.StringPtrInput                       `pulumi:"eventName"`
	EventSource          pulumi.StringPtrInput                       `pulumi:"eventSource"`
	LegacyResourceId     pulumi.StringPtrInput                       `pulumi:"legacyResourceId"`
	Level                pulumi.StringPtrInput                       `pulumi:"level"`
	MetricNamespace      pulumi.StringPtrInput                       `pulumi:"metricNamespace"`
	OdataType            pulumi.StringInput                          `pulumi:"odataType"`
	OperationName        pulumi.StringPtrInput                       `pulumi:"operationName"`
	ResourceGroupName    pulumi.StringPtrInput                       `pulumi:"resourceGroupName"`
	ResourceLocation     pulumi.StringPtrInput                       `pulumi:"resourceLocation"`
	ResourceProviderName pulumi.StringPtrInput                       `pulumi:"resourceProviderName"`
	ResourceUri          pulumi.StringPtrInput                       `pulumi:"resourceUri"`
	Status               pulumi.StringPtrInput                       `pulumi:"status"`
	SubStatus            pulumi.StringPtrInput                       `pulumi:"subStatus"`
}

func (RuleManagementEventDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventDataSource)(nil)).Elem()
}

func (i RuleManagementEventDataSourceArgs) ToRuleManagementEventDataSourceOutput() RuleManagementEventDataSourceOutput {
	return i.ToRuleManagementEventDataSourceOutputWithContext(context.Background())
}

func (i RuleManagementEventDataSourceArgs) ToRuleManagementEventDataSourceOutputWithContext(ctx context.Context) RuleManagementEventDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventDataSourceOutput)
}

type RuleManagementEventDataSourceOutput struct{ *pulumi.OutputState }

func (RuleManagementEventDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventDataSource)(nil)).Elem()
}

func (o RuleManagementEventDataSourceOutput) ToRuleManagementEventDataSourceOutput() RuleManagementEventDataSourceOutput {
	return o
}

func (o RuleManagementEventDataSourceOutput) ToRuleManagementEventDataSourceOutputWithContext(ctx context.Context) RuleManagementEventDataSourceOutput {
	return o
}

func (o RuleManagementEventDataSourceOutput) Claims() RuleManagementEventClaimsDataSourcePtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *RuleManagementEventClaimsDataSource { return v.Claims }).(RuleManagementEventClaimsDataSourcePtrOutput)
}

func (o RuleManagementEventDataSourceOutput) EventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.EventName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) EventSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.EventSource }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) LegacyResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.LegacyResourceId }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleManagementEventDataSourceOutput) OperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.OperationName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) ResourceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.ResourceProviderName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceOutput) SubStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSource) *string { return v.SubStatus }).(pulumi.StringPtrOutput)
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





type RuleManagementEventDataSourceResponseInput interface {
	pulumi.Input

	ToRuleManagementEventDataSourceResponseOutput() RuleManagementEventDataSourceResponseOutput
	ToRuleManagementEventDataSourceResponseOutputWithContext(context.Context) RuleManagementEventDataSourceResponseOutput
}

type RuleManagementEventDataSourceResponseArgs struct {
	Claims               RuleManagementEventClaimsDataSourceResponsePtrInput `pulumi:"claims"`
	EventName            pulumi.StringPtrInput                               `pulumi:"eventName"`
	EventSource          pulumi.StringPtrInput                               `pulumi:"eventSource"`
	LegacyResourceId     pulumi.StringPtrInput                               `pulumi:"legacyResourceId"`
	Level                pulumi.StringPtrInput                               `pulumi:"level"`
	MetricNamespace      pulumi.StringPtrInput                               `pulumi:"metricNamespace"`
	OdataType            pulumi.StringInput                                  `pulumi:"odataType"`
	OperationName        pulumi.StringPtrInput                               `pulumi:"operationName"`
	ResourceGroupName    pulumi.StringPtrInput                               `pulumi:"resourceGroupName"`
	ResourceLocation     pulumi.StringPtrInput                               `pulumi:"resourceLocation"`
	ResourceProviderName pulumi.StringPtrInput                               `pulumi:"resourceProviderName"`
	ResourceUri          pulumi.StringPtrInput                               `pulumi:"resourceUri"`
	Status               pulumi.StringPtrInput                               `pulumi:"status"`
	SubStatus            pulumi.StringPtrInput                               `pulumi:"subStatus"`
}

func (RuleManagementEventDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventDataSourceResponse)(nil)).Elem()
}

func (i RuleManagementEventDataSourceResponseArgs) ToRuleManagementEventDataSourceResponseOutput() RuleManagementEventDataSourceResponseOutput {
	return i.ToRuleManagementEventDataSourceResponseOutputWithContext(context.Background())
}

func (i RuleManagementEventDataSourceResponseArgs) ToRuleManagementEventDataSourceResponseOutputWithContext(ctx context.Context) RuleManagementEventDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleManagementEventDataSourceResponseOutput)
}

type RuleManagementEventDataSourceResponseOutput struct{ *pulumi.OutputState }

func (RuleManagementEventDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleManagementEventDataSourceResponse)(nil)).Elem()
}

func (o RuleManagementEventDataSourceResponseOutput) ToRuleManagementEventDataSourceResponseOutput() RuleManagementEventDataSourceResponseOutput {
	return o
}

func (o RuleManagementEventDataSourceResponseOutput) ToRuleManagementEventDataSourceResponseOutputWithContext(ctx context.Context) RuleManagementEventDataSourceResponseOutput {
	return o
}

func (o RuleManagementEventDataSourceResponseOutput) Claims() RuleManagementEventClaimsDataSourceResponsePtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *RuleManagementEventClaimsDataSourceResponse {
		return v.Claims
	}).(RuleManagementEventClaimsDataSourceResponsePtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) EventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.EventName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) EventSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.EventSource }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) LegacyResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.LegacyResourceId }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) OperationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.OperationName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) ResourceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.ResourceProviderName }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o RuleManagementEventDataSourceResponseOutput) SubStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleManagementEventDataSourceResponse) *string { return v.SubStatus }).(pulumi.StringPtrOutput)
}

type RuleMetricDataSource struct {
	LegacyResourceId *string `pulumi:"legacyResourceId"`
	MetricName       *string `pulumi:"metricName"`
	MetricNamespace  *string `pulumi:"metricNamespace"`
	OdataType        string  `pulumi:"odataType"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceUri      *string `pulumi:"resourceUri"`
}





type RuleMetricDataSourceInput interface {
	pulumi.Input

	ToRuleMetricDataSourceOutput() RuleMetricDataSourceOutput
	ToRuleMetricDataSourceOutputWithContext(context.Context) RuleMetricDataSourceOutput
}

type RuleMetricDataSourceArgs struct {
	LegacyResourceId pulumi.StringPtrInput `pulumi:"legacyResourceId"`
	MetricName       pulumi.StringPtrInput `pulumi:"metricName"`
	MetricNamespace  pulumi.StringPtrInput `pulumi:"metricNamespace"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	ResourceLocation pulumi.StringPtrInput `pulumi:"resourceLocation"`
	ResourceUri      pulumi.StringPtrInput `pulumi:"resourceUri"`
}

func (RuleMetricDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleMetricDataSource)(nil)).Elem()
}

func (i RuleMetricDataSourceArgs) ToRuleMetricDataSourceOutput() RuleMetricDataSourceOutput {
	return i.ToRuleMetricDataSourceOutputWithContext(context.Background())
}

func (i RuleMetricDataSourceArgs) ToRuleMetricDataSourceOutputWithContext(ctx context.Context) RuleMetricDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMetricDataSourceOutput)
}

type RuleMetricDataSourceOutput struct{ *pulumi.OutputState }

func (RuleMetricDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleMetricDataSource)(nil)).Elem()
}

func (o RuleMetricDataSourceOutput) ToRuleMetricDataSourceOutput() RuleMetricDataSourceOutput {
	return o
}

func (o RuleMetricDataSourceOutput) ToRuleMetricDataSourceOutputWithContext(ctx context.Context) RuleMetricDataSourceOutput {
	return o
}

func (o RuleMetricDataSourceOutput) LegacyResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.LegacyResourceId }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleMetricDataSource) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleMetricDataSourceOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSource) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

type RuleMetricDataSourceResponse struct {
	LegacyResourceId *string `pulumi:"legacyResourceId"`
	MetricName       *string `pulumi:"metricName"`
	MetricNamespace  *string `pulumi:"metricNamespace"`
	OdataType        string  `pulumi:"odataType"`
	ResourceLocation *string `pulumi:"resourceLocation"`
	ResourceUri      *string `pulumi:"resourceUri"`
}





type RuleMetricDataSourceResponseInput interface {
	pulumi.Input

	ToRuleMetricDataSourceResponseOutput() RuleMetricDataSourceResponseOutput
	ToRuleMetricDataSourceResponseOutputWithContext(context.Context) RuleMetricDataSourceResponseOutput
}

type RuleMetricDataSourceResponseArgs struct {
	LegacyResourceId pulumi.StringPtrInput `pulumi:"legacyResourceId"`
	MetricName       pulumi.StringPtrInput `pulumi:"metricName"`
	MetricNamespace  pulumi.StringPtrInput `pulumi:"metricNamespace"`
	OdataType        pulumi.StringInput    `pulumi:"odataType"`
	ResourceLocation pulumi.StringPtrInput `pulumi:"resourceLocation"`
	ResourceUri      pulumi.StringPtrInput `pulumi:"resourceUri"`
}

func (RuleMetricDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleMetricDataSourceResponse)(nil)).Elem()
}

func (i RuleMetricDataSourceResponseArgs) ToRuleMetricDataSourceResponseOutput() RuleMetricDataSourceResponseOutput {
	return i.ToRuleMetricDataSourceResponseOutputWithContext(context.Background())
}

func (i RuleMetricDataSourceResponseArgs) ToRuleMetricDataSourceResponseOutputWithContext(ctx context.Context) RuleMetricDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMetricDataSourceResponseOutput)
}

type RuleMetricDataSourceResponseOutput struct{ *pulumi.OutputState }

func (RuleMetricDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleMetricDataSourceResponse)(nil)).Elem()
}

func (o RuleMetricDataSourceResponseOutput) ToRuleMetricDataSourceResponseOutput() RuleMetricDataSourceResponseOutput {
	return o
}

func (o RuleMetricDataSourceResponseOutput) ToRuleMetricDataSourceResponseOutputWithContext(ctx context.Context) RuleMetricDataSourceResponseOutput {
	return o
}

func (o RuleMetricDataSourceResponseOutput) LegacyResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.LegacyResourceId }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceResponseOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceResponseOutput) MetricNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.MetricNamespace }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleMetricDataSourceResponseOutput) ResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.ResourceLocation }).(pulumi.StringPtrOutput)
}

func (o RuleMetricDataSourceResponseOutput) ResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleMetricDataSourceResponse) *string { return v.ResourceUri }).(pulumi.StringPtrOutput)
}

type RuleWebhookAction struct {
	OdataType  string            `pulumi:"odataType"`
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}





type RuleWebhookActionInput interface {
	pulumi.Input

	ToRuleWebhookActionOutput() RuleWebhookActionOutput
	ToRuleWebhookActionOutputWithContext(context.Context) RuleWebhookActionOutput
}

type RuleWebhookActionArgs struct {
	OdataType  pulumi.StringInput    `pulumi:"odataType"`
	Properties pulumi.StringMapInput `pulumi:"properties"`
	ServiceUri pulumi.StringPtrInput `pulumi:"serviceUri"`
}

func (RuleWebhookActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleWebhookAction)(nil)).Elem()
}

func (i RuleWebhookActionArgs) ToRuleWebhookActionOutput() RuleWebhookActionOutput {
	return i.ToRuleWebhookActionOutputWithContext(context.Background())
}

func (i RuleWebhookActionArgs) ToRuleWebhookActionOutputWithContext(ctx context.Context) RuleWebhookActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleWebhookActionOutput)
}

type RuleWebhookActionOutput struct{ *pulumi.OutputState }

func (RuleWebhookActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleWebhookAction)(nil)).Elem()
}

func (o RuleWebhookActionOutput) ToRuleWebhookActionOutput() RuleWebhookActionOutput {
	return o
}

func (o RuleWebhookActionOutput) ToRuleWebhookActionOutputWithContext(ctx context.Context) RuleWebhookActionOutput {
	return o
}

func (o RuleWebhookActionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleWebhookAction) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleWebhookActionOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v RuleWebhookAction) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o RuleWebhookActionOutput) ServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleWebhookAction) *string { return v.ServiceUri }).(pulumi.StringPtrOutput)
}

type RuleWebhookActionResponse struct {
	OdataType  string            `pulumi:"odataType"`
	Properties map[string]string `pulumi:"properties"`
	ServiceUri *string           `pulumi:"serviceUri"`
}





type RuleWebhookActionResponseInput interface {
	pulumi.Input

	ToRuleWebhookActionResponseOutput() RuleWebhookActionResponseOutput
	ToRuleWebhookActionResponseOutputWithContext(context.Context) RuleWebhookActionResponseOutput
}

type RuleWebhookActionResponseArgs struct {
	OdataType  pulumi.StringInput    `pulumi:"odataType"`
	Properties pulumi.StringMapInput `pulumi:"properties"`
	ServiceUri pulumi.StringPtrInput `pulumi:"serviceUri"`
}

func (RuleWebhookActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleWebhookActionResponse)(nil)).Elem()
}

func (i RuleWebhookActionResponseArgs) ToRuleWebhookActionResponseOutput() RuleWebhookActionResponseOutput {
	return i.ToRuleWebhookActionResponseOutputWithContext(context.Background())
}

func (i RuleWebhookActionResponseArgs) ToRuleWebhookActionResponseOutputWithContext(ctx context.Context) RuleWebhookActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleWebhookActionResponseOutput)
}

type RuleWebhookActionResponseOutput struct{ *pulumi.OutputState }

func (RuleWebhookActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleWebhookActionResponse)(nil)).Elem()
}

func (o RuleWebhookActionResponseOutput) ToRuleWebhookActionResponseOutput() RuleWebhookActionResponseOutput {
	return o
}

func (o RuleWebhookActionResponseOutput) ToRuleWebhookActionResponseOutputWithContext(ctx context.Context) RuleWebhookActionResponseOutput {
	return o
}

func (o RuleWebhookActionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleWebhookActionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o RuleWebhookActionResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v RuleWebhookActionResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o RuleWebhookActionResponseOutput) ServiceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleWebhookActionResponse) *string { return v.ServiceUri }).(pulumi.StringPtrOutput)
}

type ScaleAction struct {
	Cooldown  string         `pulumi:"cooldown"`
	Direction ScaleDirection `pulumi:"direction"`
	Type      ScaleType      `pulumi:"type"`
	Value     *string        `pulumi:"value"`
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





type ScaleActionResponseInput interface {
	pulumi.Input

	ToScaleActionResponseOutput() ScaleActionResponseOutput
	ToScaleActionResponseOutputWithContext(context.Context) ScaleActionResponseOutput
}

type ScaleActionResponseArgs struct {
	Cooldown  pulumi.StringInput    `pulumi:"cooldown"`
	Direction pulumi.StringInput    `pulumi:"direction"`
	Type      pulumi.StringInput    `pulumi:"type"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (ScaleActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleActionResponse)(nil)).Elem()
}

func (i ScaleActionResponseArgs) ToScaleActionResponseOutput() ScaleActionResponseOutput {
	return i.ToScaleActionResponseOutputWithContext(context.Background())
}

func (i ScaleActionResponseArgs) ToScaleActionResponseOutputWithContext(ctx context.Context) ScaleActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleActionResponseOutput)
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





type ScaleCapacityResponseInput interface {
	pulumi.Input

	ToScaleCapacityResponseOutput() ScaleCapacityResponseOutput
	ToScaleCapacityResponseOutputWithContext(context.Context) ScaleCapacityResponseOutput
}

type ScaleCapacityResponseArgs struct {
	Default pulumi.StringInput `pulumi:"default"`
	Maximum pulumi.StringInput `pulumi:"maximum"`
	Minimum pulumi.StringInput `pulumi:"minimum"`
}

func (ScaleCapacityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleCapacityResponse)(nil)).Elem()
}

func (i ScaleCapacityResponseArgs) ToScaleCapacityResponseOutput() ScaleCapacityResponseOutput {
	return i.ToScaleCapacityResponseOutputWithContext(context.Background())
}

func (i ScaleCapacityResponseArgs) ToScaleCapacityResponseOutputWithContext(ctx context.Context) ScaleCapacityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleCapacityResponseOutput)
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





type ScaleRuleInput interface {
	pulumi.Input

	ToScaleRuleOutput() ScaleRuleOutput
	ToScaleRuleOutputWithContext(context.Context) ScaleRuleOutput
}

type ScaleRuleArgs struct {
	MetricTrigger MetricTriggerInput `pulumi:"metricTrigger"`
	ScaleAction   ScaleActionInput   `pulumi:"scaleAction"`
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





type ScaleRuleMetricDimensionResponseInput interface {
	pulumi.Input

	ToScaleRuleMetricDimensionResponseOutput() ScaleRuleMetricDimensionResponseOutput
	ToScaleRuleMetricDimensionResponseOutputWithContext(context.Context) ScaleRuleMetricDimensionResponseOutput
}

type ScaleRuleMetricDimensionResponseArgs struct {
	DimensionName pulumi.StringInput      `pulumi:"dimensionName"`
	Operator      pulumi.StringInput      `pulumi:"operator"`
	Values        pulumi.StringArrayInput `pulumi:"values"`
}

func (ScaleRuleMetricDimensionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleMetricDimensionResponse)(nil)).Elem()
}

func (i ScaleRuleMetricDimensionResponseArgs) ToScaleRuleMetricDimensionResponseOutput() ScaleRuleMetricDimensionResponseOutput {
	return i.ToScaleRuleMetricDimensionResponseOutputWithContext(context.Background())
}

func (i ScaleRuleMetricDimensionResponseArgs) ToScaleRuleMetricDimensionResponseOutputWithContext(ctx context.Context) ScaleRuleMetricDimensionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleMetricDimensionResponseOutput)
}





type ScaleRuleMetricDimensionResponseArrayInput interface {
	pulumi.Input

	ToScaleRuleMetricDimensionResponseArrayOutput() ScaleRuleMetricDimensionResponseArrayOutput
	ToScaleRuleMetricDimensionResponseArrayOutputWithContext(context.Context) ScaleRuleMetricDimensionResponseArrayOutput
}

type ScaleRuleMetricDimensionResponseArray []ScaleRuleMetricDimensionResponseInput

func (ScaleRuleMetricDimensionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleMetricDimensionResponse)(nil)).Elem()
}

func (i ScaleRuleMetricDimensionResponseArray) ToScaleRuleMetricDimensionResponseArrayOutput() ScaleRuleMetricDimensionResponseArrayOutput {
	return i.ToScaleRuleMetricDimensionResponseArrayOutputWithContext(context.Background())
}

func (i ScaleRuleMetricDimensionResponseArray) ToScaleRuleMetricDimensionResponseArrayOutputWithContext(ctx context.Context) ScaleRuleMetricDimensionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleMetricDimensionResponseArrayOutput)
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





type ScaleRuleResponseInput interface {
	pulumi.Input

	ToScaleRuleResponseOutput() ScaleRuleResponseOutput
	ToScaleRuleResponseOutputWithContext(context.Context) ScaleRuleResponseOutput
}

type ScaleRuleResponseArgs struct {
	MetricTrigger MetricTriggerResponseInput `pulumi:"metricTrigger"`
	ScaleAction   ScaleActionResponseInput   `pulumi:"scaleAction"`
}

func (ScaleRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleResponse)(nil)).Elem()
}

func (i ScaleRuleResponseArgs) ToScaleRuleResponseOutput() ScaleRuleResponseOutput {
	return i.ToScaleRuleResponseOutputWithContext(context.Background())
}

func (i ScaleRuleResponseArgs) ToScaleRuleResponseOutputWithContext(ctx context.Context) ScaleRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleResponseOutput)
}





type ScaleRuleResponseArrayInput interface {
	pulumi.Input

	ToScaleRuleResponseArrayOutput() ScaleRuleResponseArrayOutput
	ToScaleRuleResponseArrayOutputWithContext(context.Context) ScaleRuleResponseArrayOutput
}

type ScaleRuleResponseArray []ScaleRuleResponseInput

func (ScaleRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleResponse)(nil)).Elem()
}

func (i ScaleRuleResponseArray) ToScaleRuleResponseArrayOutput() ScaleRuleResponseArrayOutput {
	return i.ToScaleRuleResponseArrayOutputWithContext(context.Background())
}

func (i ScaleRuleResponseArray) ToScaleRuleResponseArrayOutputWithContext(ctx context.Context) ScaleRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleResponseArrayOutput)
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





type ScheduleResponseInput interface {
	pulumi.Input

	ToScheduleResponseOutput() ScheduleResponseOutput
	ToScheduleResponseOutputWithContext(context.Context) ScheduleResponseOutput
}

type ScheduleResponseArgs struct {
	FrequencyInMinutes  pulumi.IntInput `pulumi:"frequencyInMinutes"`
	TimeWindowInMinutes pulumi.IntInput `pulumi:"timeWindowInMinutes"`
}

func (ScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (i ScheduleResponseArgs) ToScheduleResponseOutput() ScheduleResponseOutput {
	return i.ToScheduleResponseOutputWithContext(context.Background())
}

func (i ScheduleResponseArgs) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleResponseOutput)
}

func (i ScheduleResponseArgs) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return i.ToScheduleResponsePtrOutputWithContext(context.Background())
}

func (i ScheduleResponseArgs) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleResponseOutput).ToScheduleResponsePtrOutputWithContext(ctx)
}









type ScheduleResponsePtrInput interface {
	pulumi.Input

	ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput
	ToScheduleResponsePtrOutputWithContext(context.Context) ScheduleResponsePtrOutput
}

type scheduleResponsePtrType ScheduleResponseArgs

func ScheduleResponsePtr(v *ScheduleResponseArgs) ScheduleResponsePtrInput {
	return (*scheduleResponsePtrType)(v)
}

func (*scheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResponse)(nil)).Elem()
}

func (i *scheduleResponsePtrType) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return i.ToScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *scheduleResponsePtrType) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleResponsePtrOutput)
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

func (o ScheduleResponseOutput) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return o.ToScheduleResponsePtrOutputWithContext(context.Background())
}

func (o ScheduleResponseOutput) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleResponse) *ScheduleResponse {
		return &v
	}).(ScheduleResponsePtrOutput)
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





type SinkConfigurationResponseInput interface {
	pulumi.Input

	ToSinkConfigurationResponseOutput() SinkConfigurationResponseOutput
	ToSinkConfigurationResponseOutputWithContext(context.Context) SinkConfigurationResponseOutput
}

type SinkConfigurationResponseArgs struct {
	Kind pulumi.StringInput `pulumi:"kind"`
}

func (SinkConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SinkConfigurationResponse)(nil)).Elem()
}

func (i SinkConfigurationResponseArgs) ToSinkConfigurationResponseOutput() SinkConfigurationResponseOutput {
	return i.ToSinkConfigurationResponseOutputWithContext(context.Background())
}

func (i SinkConfigurationResponseArgs) ToSinkConfigurationResponseOutputWithContext(ctx context.Context) SinkConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SinkConfigurationResponseOutput)
}





type SinkConfigurationResponseArrayInput interface {
	pulumi.Input

	ToSinkConfigurationResponseArrayOutput() SinkConfigurationResponseArrayOutput
	ToSinkConfigurationResponseArrayOutputWithContext(context.Context) SinkConfigurationResponseArrayOutput
}

type SinkConfigurationResponseArray []SinkConfigurationResponseInput

func (SinkConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SinkConfigurationResponse)(nil)).Elem()
}

func (i SinkConfigurationResponseArray) ToSinkConfigurationResponseArrayOutput() SinkConfigurationResponseArrayOutput {
	return i.ToSinkConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i SinkConfigurationResponseArray) ToSinkConfigurationResponseArrayOutputWithContext(ctx context.Context) SinkConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SinkConfigurationResponseArrayOutput)
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





type SmsReceiverResponseInput interface {
	pulumi.Input

	ToSmsReceiverResponseOutput() SmsReceiverResponseOutput
	ToSmsReceiverResponseOutputWithContext(context.Context) SmsReceiverResponseOutput
}

type SmsReceiverResponseArgs struct {
	CountryCode pulumi.StringInput `pulumi:"countryCode"`
	Name        pulumi.StringInput `pulumi:"name"`
	PhoneNumber pulumi.StringInput `pulumi:"phoneNumber"`
	Status      pulumi.StringInput `pulumi:"status"`
}

func (SmsReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsReceiverResponse)(nil)).Elem()
}

func (i SmsReceiverResponseArgs) ToSmsReceiverResponseOutput() SmsReceiverResponseOutput {
	return i.ToSmsReceiverResponseOutputWithContext(context.Background())
}

func (i SmsReceiverResponseArgs) ToSmsReceiverResponseOutputWithContext(ctx context.Context) SmsReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsReceiverResponseOutput)
}





type SmsReceiverResponseArrayInput interface {
	pulumi.Input

	ToSmsReceiverResponseArrayOutput() SmsReceiverResponseArrayOutput
	ToSmsReceiverResponseArrayOutputWithContext(context.Context) SmsReceiverResponseArrayOutput
}

type SmsReceiverResponseArray []SmsReceiverResponseInput

func (SmsReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SmsReceiverResponse)(nil)).Elem()
}

func (i SmsReceiverResponseArray) ToSmsReceiverResponseArrayOutput() SmsReceiverResponseArrayOutput {
	return i.ToSmsReceiverResponseArrayOutputWithContext(context.Background())
}

func (i SmsReceiverResponseArray) ToSmsReceiverResponseArrayOutputWithContext(ctx context.Context) SmsReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsReceiverResponseArrayOutput)
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

func (i SourceArgs) ToSourcePtrOutput() SourcePtrOutput {
	return i.ToSourcePtrOutputWithContext(context.Background())
}

func (i SourceArgs) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceOutput).ToSourcePtrOutputWithContext(ctx)
}









type SourcePtrInput interface {
	pulumi.Input

	ToSourcePtrOutput() SourcePtrOutput
	ToSourcePtrOutputWithContext(context.Context) SourcePtrOutput
}

type sourcePtrType SourceArgs

func SourcePtr(v *SourceArgs) SourcePtrInput {
	return (*sourcePtrType)(v)
}

func (*sourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Source)(nil)).Elem()
}

func (i *sourcePtrType) ToSourcePtrOutput() SourcePtrOutput {
	return i.ToSourcePtrOutputWithContext(context.Background())
}

func (i *sourcePtrType) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourcePtrOutput)
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

func (o SourceOutput) ToSourcePtrOutput() SourcePtrOutput {
	return o.ToSourcePtrOutputWithContext(context.Background())
}

func (o SourceOutput) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Source) *Source {
		return &v
	}).(SourcePtrOutput)
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

type SourcePtrOutput struct{ *pulumi.OutputState }

func (SourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Source)(nil)).Elem()
}

func (o SourcePtrOutput) ToSourcePtrOutput() SourcePtrOutput {
	return o
}

func (o SourcePtrOutput) ToSourcePtrOutputWithContext(ctx context.Context) SourcePtrOutput {
	return o
}

func (o SourcePtrOutput) Elem() SourceOutput {
	return o.ApplyT(func(v *Source) Source {
		if v != nil {
			return *v
		}
		var ret Source
		return ret
	}).(SourceOutput)
}

func (o SourcePtrOutput) AuthorizedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Source) []string {
		if v == nil {
			return nil
		}
		return v.AuthorizedResources
	}).(pulumi.StringArrayOutput)
}

func (o SourcePtrOutput) DataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Source) *string {
		if v == nil {
			return nil
		}
		return &v.DataSourceId
	}).(pulumi.StringPtrOutput)
}

func (o SourcePtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Source) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o SourcePtrOutput) QueryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Source) *string {
		if v == nil {
			return nil
		}
		return v.QueryType
	}).(pulumi.StringPtrOutput)
}

type SourceResponse struct {
	AuthorizedResources []string `pulumi:"authorizedResources"`
	DataSourceId        string   `pulumi:"dataSourceId"`
	Query               *string  `pulumi:"query"`
	QueryType           *string  `pulumi:"queryType"`
}





type SourceResponseInput interface {
	pulumi.Input

	ToSourceResponseOutput() SourceResponseOutput
	ToSourceResponseOutputWithContext(context.Context) SourceResponseOutput
}

type SourceResponseArgs struct {
	AuthorizedResources pulumi.StringArrayInput `pulumi:"authorizedResources"`
	DataSourceId        pulumi.StringInput      `pulumi:"dataSourceId"`
	Query               pulumi.StringPtrInput   `pulumi:"query"`
	QueryType           pulumi.StringPtrInput   `pulumi:"queryType"`
}

func (SourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceResponse)(nil)).Elem()
}

func (i SourceResponseArgs) ToSourceResponseOutput() SourceResponseOutput {
	return i.ToSourceResponseOutputWithContext(context.Background())
}

func (i SourceResponseArgs) ToSourceResponseOutputWithContext(ctx context.Context) SourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceResponseOutput)
}

func (i SourceResponseArgs) ToSourceResponsePtrOutput() SourceResponsePtrOutput {
	return i.ToSourceResponsePtrOutputWithContext(context.Background())
}

func (i SourceResponseArgs) ToSourceResponsePtrOutputWithContext(ctx context.Context) SourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceResponseOutput).ToSourceResponsePtrOutputWithContext(ctx)
}









type SourceResponsePtrInput interface {
	pulumi.Input

	ToSourceResponsePtrOutput() SourceResponsePtrOutput
	ToSourceResponsePtrOutputWithContext(context.Context) SourceResponsePtrOutput
}

type sourceResponsePtrType SourceResponseArgs

func SourceResponsePtr(v *SourceResponseArgs) SourceResponsePtrInput {
	return (*sourceResponsePtrType)(v)
}

func (*sourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceResponse)(nil)).Elem()
}

func (i *sourceResponsePtrType) ToSourceResponsePtrOutput() SourceResponsePtrOutput {
	return i.ToSourceResponsePtrOutputWithContext(context.Background())
}

func (i *sourceResponsePtrType) ToSourceResponsePtrOutputWithContext(ctx context.Context) SourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceResponsePtrOutput)
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

func (o SourceResponseOutput) ToSourceResponsePtrOutput() SourceResponsePtrOutput {
	return o.ToSourceResponsePtrOutputWithContext(context.Background())
}

func (o SourceResponseOutput) ToSourceResponsePtrOutputWithContext(ctx context.Context) SourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceResponse) *SourceResponse {
		return &v
	}).(SourceResponsePtrOutput)
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

type SourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceResponse)(nil)).Elem()
}

func (o SourceResponsePtrOutput) ToSourceResponsePtrOutput() SourceResponsePtrOutput {
	return o
}

func (o SourceResponsePtrOutput) ToSourceResponsePtrOutputWithContext(ctx context.Context) SourceResponsePtrOutput {
	return o
}

func (o SourceResponsePtrOutput) Elem() SourceResponseOutput {
	return o.ApplyT(func(v *SourceResponse) SourceResponse {
		if v != nil {
			return *v
		}
		var ret SourceResponse
		return ret
	}).(SourceResponseOutput)
}

func (o SourceResponsePtrOutput) AuthorizedResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SourceResponse) []string {
		if v == nil {
			return nil
		}
		return v.AuthorizedResources
	}).(pulumi.StringArrayOutput)
}

func (o SourceResponsePtrOutput) DataSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataSourceId
	}).(pulumi.StringPtrOutput)
}

func (o SourceResponsePtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o SourceResponsePtrOutput) QueryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.QueryType
	}).(pulumi.StringPtrOutput)
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





type SubscriptionLogSettingsResponseInput interface {
	pulumi.Input

	ToSubscriptionLogSettingsResponseOutput() SubscriptionLogSettingsResponseOutput
	ToSubscriptionLogSettingsResponseOutputWithContext(context.Context) SubscriptionLogSettingsResponseOutput
}

type SubscriptionLogSettingsResponseArgs struct {
	Category pulumi.StringPtrInput `pulumi:"category"`
	Enabled  pulumi.BoolInput      `pulumi:"enabled"`
}

func (SubscriptionLogSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionLogSettingsResponse)(nil)).Elem()
}

func (i SubscriptionLogSettingsResponseArgs) ToSubscriptionLogSettingsResponseOutput() SubscriptionLogSettingsResponseOutput {
	return i.ToSubscriptionLogSettingsResponseOutputWithContext(context.Background())
}

func (i SubscriptionLogSettingsResponseArgs) ToSubscriptionLogSettingsResponseOutputWithContext(ctx context.Context) SubscriptionLogSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionLogSettingsResponseOutput)
}





type SubscriptionLogSettingsResponseArrayInput interface {
	pulumi.Input

	ToSubscriptionLogSettingsResponseArrayOutput() SubscriptionLogSettingsResponseArrayOutput
	ToSubscriptionLogSettingsResponseArrayOutputWithContext(context.Context) SubscriptionLogSettingsResponseArrayOutput
}

type SubscriptionLogSettingsResponseArray []SubscriptionLogSettingsResponseInput

func (SubscriptionLogSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionLogSettingsResponse)(nil)).Elem()
}

func (i SubscriptionLogSettingsResponseArray) ToSubscriptionLogSettingsResponseArrayOutput() SubscriptionLogSettingsResponseArrayOutput {
	return i.ToSubscriptionLogSettingsResponseArrayOutputWithContext(context.Background())
}

func (i SubscriptionLogSettingsResponseArray) ToSubscriptionLogSettingsResponseArrayOutputWithContext(ctx context.Context) SubscriptionLogSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionLogSettingsResponseArrayOutput)
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





type SyslogDataSourceResponseInput interface {
	pulumi.Input

	ToSyslogDataSourceResponseOutput() SyslogDataSourceResponseOutput
	ToSyslogDataSourceResponseOutputWithContext(context.Context) SyslogDataSourceResponseOutput
}

type SyslogDataSourceResponseArgs struct {
	FacilityNames pulumi.StringArrayInput `pulumi:"facilityNames"`
	LogLevels     pulumi.StringArrayInput `pulumi:"logLevels"`
	Name          pulumi.StringPtrInput   `pulumi:"name"`
	Streams       pulumi.StringArrayInput `pulumi:"streams"`
}

func (SyslogDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyslogDataSourceResponse)(nil)).Elem()
}

func (i SyslogDataSourceResponseArgs) ToSyslogDataSourceResponseOutput() SyslogDataSourceResponseOutput {
	return i.ToSyslogDataSourceResponseOutputWithContext(context.Background())
}

func (i SyslogDataSourceResponseArgs) ToSyslogDataSourceResponseOutputWithContext(ctx context.Context) SyslogDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyslogDataSourceResponseOutput)
}





type SyslogDataSourceResponseArrayInput interface {
	pulumi.Input

	ToSyslogDataSourceResponseArrayOutput() SyslogDataSourceResponseArrayOutput
	ToSyslogDataSourceResponseArrayOutputWithContext(context.Context) SyslogDataSourceResponseArrayOutput
}

type SyslogDataSourceResponseArray []SyslogDataSourceResponseInput

func (SyslogDataSourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyslogDataSourceResponse)(nil)).Elem()
}

func (i SyslogDataSourceResponseArray) ToSyslogDataSourceResponseArrayOutput() SyslogDataSourceResponseArrayOutput {
	return i.ToSyslogDataSourceResponseArrayOutputWithContext(context.Background())
}

func (i SyslogDataSourceResponseArray) ToSyslogDataSourceResponseArrayOutputWithContext(ctx context.Context) SyslogDataSourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyslogDataSourceResponseArrayOutput)
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





type ThresholdRuleConditionInput interface {
	pulumi.Input

	ToThresholdRuleConditionOutput() ThresholdRuleConditionOutput
	ToThresholdRuleConditionOutputWithContext(context.Context) ThresholdRuleConditionOutput
}

type ThresholdRuleConditionArgs struct {
	DataSource      pulumi.Input                    `pulumi:"dataSource"`
	OdataType       pulumi.StringInput              `pulumi:"odataType"`
	Operator        ConditionOperatorInput          `pulumi:"operator"`
	Threshold       pulumi.Float64Input             `pulumi:"threshold"`
	TimeAggregation TimeAggregationOperatorPtrInput `pulumi:"timeAggregation"`
	WindowSize      pulumi.StringPtrInput           `pulumi:"windowSize"`
}

func (ThresholdRuleConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdRuleCondition)(nil)).Elem()
}

func (i ThresholdRuleConditionArgs) ToThresholdRuleConditionOutput() ThresholdRuleConditionOutput {
	return i.ToThresholdRuleConditionOutputWithContext(context.Background())
}

func (i ThresholdRuleConditionArgs) ToThresholdRuleConditionOutputWithContext(ctx context.Context) ThresholdRuleConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdRuleConditionOutput)
}

type ThresholdRuleConditionOutput struct{ *pulumi.OutputState }

func (ThresholdRuleConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdRuleCondition)(nil)).Elem()
}

func (o ThresholdRuleConditionOutput) ToThresholdRuleConditionOutput() ThresholdRuleConditionOutput {
	return o
}

func (o ThresholdRuleConditionOutput) ToThresholdRuleConditionOutputWithContext(ctx context.Context) ThresholdRuleConditionOutput {
	return o
}

func (o ThresholdRuleConditionOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o ThresholdRuleConditionOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ThresholdRuleConditionOutput) Operator() ConditionOperatorOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) ConditionOperator { return v.Operator }).(ConditionOperatorOutput)
}

func (o ThresholdRuleConditionOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v ThresholdRuleCondition) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o ThresholdRuleConditionOutput) TimeAggregation() TimeAggregationOperatorPtrOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) *TimeAggregationOperator { return v.TimeAggregation }).(TimeAggregationOperatorPtrOutput)
}

func (o ThresholdRuleConditionOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThresholdRuleCondition) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
}

type ThresholdRuleConditionResponse struct {
	DataSource      interface{} `pulumi:"dataSource"`
	OdataType       string      `pulumi:"odataType"`
	Operator        string      `pulumi:"operator"`
	Threshold       float64     `pulumi:"threshold"`
	TimeAggregation *string     `pulumi:"timeAggregation"`
	WindowSize      *string     `pulumi:"windowSize"`
}





type ThresholdRuleConditionResponseInput interface {
	pulumi.Input

	ToThresholdRuleConditionResponseOutput() ThresholdRuleConditionResponseOutput
	ToThresholdRuleConditionResponseOutputWithContext(context.Context) ThresholdRuleConditionResponseOutput
}

type ThresholdRuleConditionResponseArgs struct {
	DataSource      pulumi.Input          `pulumi:"dataSource"`
	OdataType       pulumi.StringInput    `pulumi:"odataType"`
	Operator        pulumi.StringInput    `pulumi:"operator"`
	Threshold       pulumi.Float64Input   `pulumi:"threshold"`
	TimeAggregation pulumi.StringPtrInput `pulumi:"timeAggregation"`
	WindowSize      pulumi.StringPtrInput `pulumi:"windowSize"`
}

func (ThresholdRuleConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdRuleConditionResponse)(nil)).Elem()
}

func (i ThresholdRuleConditionResponseArgs) ToThresholdRuleConditionResponseOutput() ThresholdRuleConditionResponseOutput {
	return i.ToThresholdRuleConditionResponseOutputWithContext(context.Background())
}

func (i ThresholdRuleConditionResponseArgs) ToThresholdRuleConditionResponseOutputWithContext(ctx context.Context) ThresholdRuleConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThresholdRuleConditionResponseOutput)
}

type ThresholdRuleConditionResponseOutput struct{ *pulumi.OutputState }

func (ThresholdRuleConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdRuleConditionResponse)(nil)).Elem()
}

func (o ThresholdRuleConditionResponseOutput) ToThresholdRuleConditionResponseOutput() ThresholdRuleConditionResponseOutput {
	return o
}

func (o ThresholdRuleConditionResponseOutput) ToThresholdRuleConditionResponseOutputWithContext(ctx context.Context) ThresholdRuleConditionResponseOutput {
	return o
}

func (o ThresholdRuleConditionResponseOutput) DataSource() pulumi.AnyOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) interface{} { return v.DataSource }).(pulumi.AnyOutput)
}

func (o ThresholdRuleConditionResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o ThresholdRuleConditionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ThresholdRuleConditionResponseOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o ThresholdRuleConditionResponseOutput) TimeAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) *string { return v.TimeAggregation }).(pulumi.StringPtrOutput)
}

func (o ThresholdRuleConditionResponseOutput) WindowSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThresholdRuleConditionResponse) *string { return v.WindowSize }).(pulumi.StringPtrOutput)
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





type TimeWindowResponseInput interface {
	pulumi.Input

	ToTimeWindowResponseOutput() TimeWindowResponseOutput
	ToTimeWindowResponseOutputWithContext(context.Context) TimeWindowResponseOutput
}

type TimeWindowResponseArgs struct {
	End      pulumi.StringInput    `pulumi:"end"`
	Start    pulumi.StringInput    `pulumi:"start"`
	TimeZone pulumi.StringPtrInput `pulumi:"timeZone"`
}

func (TimeWindowResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeWindowResponse)(nil)).Elem()
}

func (i TimeWindowResponseArgs) ToTimeWindowResponseOutput() TimeWindowResponseOutput {
	return i.ToTimeWindowResponseOutputWithContext(context.Background())
}

func (i TimeWindowResponseArgs) ToTimeWindowResponseOutputWithContext(ctx context.Context) TimeWindowResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowResponseOutput)
}

func (i TimeWindowResponseArgs) ToTimeWindowResponsePtrOutput() TimeWindowResponsePtrOutput {
	return i.ToTimeWindowResponsePtrOutputWithContext(context.Background())
}

func (i TimeWindowResponseArgs) ToTimeWindowResponsePtrOutputWithContext(ctx context.Context) TimeWindowResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowResponseOutput).ToTimeWindowResponsePtrOutputWithContext(ctx)
}









type TimeWindowResponsePtrInput interface {
	pulumi.Input

	ToTimeWindowResponsePtrOutput() TimeWindowResponsePtrOutput
	ToTimeWindowResponsePtrOutputWithContext(context.Context) TimeWindowResponsePtrOutput
}

type timeWindowResponsePtrType TimeWindowResponseArgs

func TimeWindowResponsePtr(v *TimeWindowResponseArgs) TimeWindowResponsePtrInput {
	return (*timeWindowResponsePtrType)(v)
}

func (*timeWindowResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeWindowResponse)(nil)).Elem()
}

func (i *timeWindowResponsePtrType) ToTimeWindowResponsePtrOutput() TimeWindowResponsePtrOutput {
	return i.ToTimeWindowResponsePtrOutputWithContext(context.Background())
}

func (i *timeWindowResponsePtrType) ToTimeWindowResponsePtrOutputWithContext(ctx context.Context) TimeWindowResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeWindowResponsePtrOutput)
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

func (o TimeWindowResponseOutput) ToTimeWindowResponsePtrOutput() TimeWindowResponsePtrOutput {
	return o.ToTimeWindowResponsePtrOutputWithContext(context.Background())
}

func (o TimeWindowResponseOutput) ToTimeWindowResponsePtrOutputWithContext(ctx context.Context) TimeWindowResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeWindowResponse) *TimeWindowResponse {
		return &v
	}).(TimeWindowResponsePtrOutput)
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





type TriggerConditionInput interface {
	pulumi.Input

	ToTriggerConditionOutput() TriggerConditionOutput
	ToTriggerConditionOutputWithContext(context.Context) TriggerConditionOutput
}

type TriggerConditionArgs struct {
	MetricTrigger     LogMetricTriggerPtrInput `pulumi:"metricTrigger"`
	Threshold         pulumi.Float64Input      `pulumi:"threshold"`
	ThresholdOperator pulumi.StringInput       `pulumi:"thresholdOperator"`
}

func (TriggerConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerCondition)(nil)).Elem()
}

func (i TriggerConditionArgs) ToTriggerConditionOutput() TriggerConditionOutput {
	return i.ToTriggerConditionOutputWithContext(context.Background())
}

func (i TriggerConditionArgs) ToTriggerConditionOutputWithContext(ctx context.Context) TriggerConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerConditionOutput)
}

type TriggerConditionOutput struct{ *pulumi.OutputState }

func (TriggerConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerCondition)(nil)).Elem()
}

func (o TriggerConditionOutput) ToTriggerConditionOutput() TriggerConditionOutput {
	return o
}

func (o TriggerConditionOutput) ToTriggerConditionOutputWithContext(ctx context.Context) TriggerConditionOutput {
	return o
}

func (o TriggerConditionOutput) MetricTrigger() LogMetricTriggerPtrOutput {
	return o.ApplyT(func(v TriggerCondition) *LogMetricTrigger { return v.MetricTrigger }).(LogMetricTriggerPtrOutput)
}

func (o TriggerConditionOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v TriggerCondition) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o TriggerConditionOutput) ThresholdOperator() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerCondition) string { return v.ThresholdOperator }).(pulumi.StringOutput)
}

type TriggerConditionResponse struct {
	MetricTrigger     *LogMetricTriggerResponse `pulumi:"metricTrigger"`
	Threshold         float64                   `pulumi:"threshold"`
	ThresholdOperator string                    `pulumi:"thresholdOperator"`
}





type TriggerConditionResponseInput interface {
	pulumi.Input

	ToTriggerConditionResponseOutput() TriggerConditionResponseOutput
	ToTriggerConditionResponseOutputWithContext(context.Context) TriggerConditionResponseOutput
}

type TriggerConditionResponseArgs struct {
	MetricTrigger     LogMetricTriggerResponsePtrInput `pulumi:"metricTrigger"`
	Threshold         pulumi.Float64Input              `pulumi:"threshold"`
	ThresholdOperator pulumi.StringInput               `pulumi:"thresholdOperator"`
}

func (TriggerConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerConditionResponse)(nil)).Elem()
}

func (i TriggerConditionResponseArgs) ToTriggerConditionResponseOutput() TriggerConditionResponseOutput {
	return i.ToTriggerConditionResponseOutputWithContext(context.Background())
}

func (i TriggerConditionResponseArgs) ToTriggerConditionResponseOutputWithContext(ctx context.Context) TriggerConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerConditionResponseOutput)
}

type TriggerConditionResponseOutput struct{ *pulumi.OutputState }

func (TriggerConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerConditionResponse)(nil)).Elem()
}

func (o TriggerConditionResponseOutput) ToTriggerConditionResponseOutput() TriggerConditionResponseOutput {
	return o
}

func (o TriggerConditionResponseOutput) ToTriggerConditionResponseOutputWithContext(ctx context.Context) TriggerConditionResponseOutput {
	return o
}

func (o TriggerConditionResponseOutput) MetricTrigger() LogMetricTriggerResponsePtrOutput {
	return o.ApplyT(func(v TriggerConditionResponse) *LogMetricTriggerResponse { return v.MetricTrigger }).(LogMetricTriggerResponsePtrOutput)
}

func (o TriggerConditionResponseOutput) Threshold() pulumi.Float64Output {
	return o.ApplyT(func(v TriggerConditionResponse) float64 { return v.Threshold }).(pulumi.Float64Output)
}

func (o TriggerConditionResponseOutput) ThresholdOperator() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerConditionResponse) string { return v.ThresholdOperator }).(pulumi.StringOutput)
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





type VoiceReceiverResponseInput interface {
	pulumi.Input

	ToVoiceReceiverResponseOutput() VoiceReceiverResponseOutput
	ToVoiceReceiverResponseOutputWithContext(context.Context) VoiceReceiverResponseOutput
}

type VoiceReceiverResponseArgs struct {
	CountryCode pulumi.StringInput `pulumi:"countryCode"`
	Name        pulumi.StringInput `pulumi:"name"`
	PhoneNumber pulumi.StringInput `pulumi:"phoneNumber"`
}

func (VoiceReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VoiceReceiverResponse)(nil)).Elem()
}

func (i VoiceReceiverResponseArgs) ToVoiceReceiverResponseOutput() VoiceReceiverResponseOutput {
	return i.ToVoiceReceiverResponseOutputWithContext(context.Background())
}

func (i VoiceReceiverResponseArgs) ToVoiceReceiverResponseOutputWithContext(ctx context.Context) VoiceReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceReceiverResponseOutput)
}





type VoiceReceiverResponseArrayInput interface {
	pulumi.Input

	ToVoiceReceiverResponseArrayOutput() VoiceReceiverResponseArrayOutput
	ToVoiceReceiverResponseArrayOutputWithContext(context.Context) VoiceReceiverResponseArrayOutput
}

type VoiceReceiverResponseArray []VoiceReceiverResponseInput

func (VoiceReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VoiceReceiverResponse)(nil)).Elem()
}

func (i VoiceReceiverResponseArray) ToVoiceReceiverResponseArrayOutput() VoiceReceiverResponseArrayOutput {
	return i.ToVoiceReceiverResponseArrayOutputWithContext(context.Background())
}

func (i VoiceReceiverResponseArray) ToVoiceReceiverResponseArrayOutputWithContext(ctx context.Context) VoiceReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VoiceReceiverResponseArrayOutput)
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





type WebTestGeolocationResponseInput interface {
	pulumi.Input

	ToWebTestGeolocationResponseOutput() WebTestGeolocationResponseOutput
	ToWebTestGeolocationResponseOutputWithContext(context.Context) WebTestGeolocationResponseOutput
}

type WebTestGeolocationResponseArgs struct {
	Location pulumi.StringPtrInput `pulumi:"location"`
}

func (WebTestGeolocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestGeolocationResponse)(nil)).Elem()
}

func (i WebTestGeolocationResponseArgs) ToWebTestGeolocationResponseOutput() WebTestGeolocationResponseOutput {
	return i.ToWebTestGeolocationResponseOutputWithContext(context.Background())
}

func (i WebTestGeolocationResponseArgs) ToWebTestGeolocationResponseOutputWithContext(ctx context.Context) WebTestGeolocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationResponseOutput)
}





type WebTestGeolocationResponseArrayInput interface {
	pulumi.Input

	ToWebTestGeolocationResponseArrayOutput() WebTestGeolocationResponseArrayOutput
	ToWebTestGeolocationResponseArrayOutputWithContext(context.Context) WebTestGeolocationResponseArrayOutput
}

type WebTestGeolocationResponseArray []WebTestGeolocationResponseInput

func (WebTestGeolocationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTestGeolocationResponse)(nil)).Elem()
}

func (i WebTestGeolocationResponseArray) ToWebTestGeolocationResponseArrayOutput() WebTestGeolocationResponseArrayOutput {
	return i.ToWebTestGeolocationResponseArrayOutputWithContext(context.Background())
}

func (i WebTestGeolocationResponseArray) ToWebTestGeolocationResponseArrayOutputWithContext(ctx context.Context) WebTestGeolocationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestGeolocationResponseArrayOutput)
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





type WebTestPropertiesResponseConfigurationInput interface {
	pulumi.Input

	ToWebTestPropertiesResponseConfigurationOutput() WebTestPropertiesResponseConfigurationOutput
	ToWebTestPropertiesResponseConfigurationOutputWithContext(context.Context) WebTestPropertiesResponseConfigurationOutput
}

type WebTestPropertiesResponseConfigurationArgs struct {
	WebTest pulumi.StringPtrInput `pulumi:"webTest"`
}

func (WebTestPropertiesResponseConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (i WebTestPropertiesResponseConfigurationArgs) ToWebTestPropertiesResponseConfigurationOutput() WebTestPropertiesResponseConfigurationOutput {
	return i.ToWebTestPropertiesResponseConfigurationOutputWithContext(context.Background())
}

func (i WebTestPropertiesResponseConfigurationArgs) ToWebTestPropertiesResponseConfigurationOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesResponseConfigurationOutput)
}

func (i WebTestPropertiesResponseConfigurationArgs) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return i.ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(context.Background())
}

func (i WebTestPropertiesResponseConfigurationArgs) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesResponseConfigurationOutput).ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx)
}









type WebTestPropertiesResponseConfigurationPtrInput interface {
	pulumi.Input

	ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput
	ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(context.Context) WebTestPropertiesResponseConfigurationPtrOutput
}

type webTestPropertiesResponseConfigurationPtrType WebTestPropertiesResponseConfigurationArgs

func WebTestPropertiesResponseConfigurationPtr(v *WebTestPropertiesResponseConfigurationArgs) WebTestPropertiesResponseConfigurationPtrInput {
	return (*webTestPropertiesResponseConfigurationPtrType)(v)
}

func (*webTestPropertiesResponseConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestPropertiesResponseConfiguration)(nil)).Elem()
}

func (i *webTestPropertiesResponseConfigurationPtrType) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return i.ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(context.Background())
}

func (i *webTestPropertiesResponseConfigurationPtrType) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestPropertiesResponseConfigurationPtrOutput)
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

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationPtrOutput() WebTestPropertiesResponseConfigurationPtrOutput {
	return o.ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(context.Background())
}

func (o WebTestPropertiesResponseConfigurationOutput) ToWebTestPropertiesResponseConfigurationPtrOutputWithContext(ctx context.Context) WebTestPropertiesResponseConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTestPropertiesResponseConfiguration) *WebTestPropertiesResponseConfiguration {
		return &v
	}).(WebTestPropertiesResponseConfigurationPtrOutput)
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





type WebhookNotificationResponseInput interface {
	pulumi.Input

	ToWebhookNotificationResponseOutput() WebhookNotificationResponseOutput
	ToWebhookNotificationResponseOutputWithContext(context.Context) WebhookNotificationResponseOutput
}

type WebhookNotificationResponseArgs struct {
	Properties pulumi.StringMapInput `pulumi:"properties"`
	ServiceUri pulumi.StringPtrInput `pulumi:"serviceUri"`
}

func (WebhookNotificationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookNotificationResponse)(nil)).Elem()
}

func (i WebhookNotificationResponseArgs) ToWebhookNotificationResponseOutput() WebhookNotificationResponseOutput {
	return i.ToWebhookNotificationResponseOutputWithContext(context.Background())
}

func (i WebhookNotificationResponseArgs) ToWebhookNotificationResponseOutputWithContext(ctx context.Context) WebhookNotificationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookNotificationResponseOutput)
}





type WebhookNotificationResponseArrayInput interface {
	pulumi.Input

	ToWebhookNotificationResponseArrayOutput() WebhookNotificationResponseArrayOutput
	ToWebhookNotificationResponseArrayOutputWithContext(context.Context) WebhookNotificationResponseArrayOutput
}

type WebhookNotificationResponseArray []WebhookNotificationResponseInput

func (WebhookNotificationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebhookNotificationResponse)(nil)).Elem()
}

func (i WebhookNotificationResponseArray) ToWebhookNotificationResponseArrayOutput() WebhookNotificationResponseArrayOutput {
	return i.ToWebhookNotificationResponseArrayOutputWithContext(context.Background())
}

func (i WebhookNotificationResponseArray) ToWebhookNotificationResponseArrayOutputWithContext(ctx context.Context) WebhookNotificationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookNotificationResponseArrayOutput)
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





type WebhookReceiverResponseInput interface {
	pulumi.Input

	ToWebhookReceiverResponseOutput() WebhookReceiverResponseOutput
	ToWebhookReceiverResponseOutputWithContext(context.Context) WebhookReceiverResponseOutput
}

type WebhookReceiverResponseArgs struct {
	IdentifierUri        pulumi.StringPtrInput `pulumi:"identifierUri"`
	Name                 pulumi.StringInput    `pulumi:"name"`
	ObjectId             pulumi.StringPtrInput `pulumi:"objectId"`
	ServiceUri           pulumi.StringInput    `pulumi:"serviceUri"`
	TenantId             pulumi.StringPtrInput `pulumi:"tenantId"`
	UseAadAuth           pulumi.BoolPtrInput   `pulumi:"useAadAuth"`
	UseCommonAlertSchema pulumi.BoolPtrInput   `pulumi:"useCommonAlertSchema"`
}

func (WebhookReceiverResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookReceiverResponse)(nil)).Elem()
}

func (i WebhookReceiverResponseArgs) ToWebhookReceiverResponseOutput() WebhookReceiverResponseOutput {
	return i.ToWebhookReceiverResponseOutputWithContext(context.Background())
}

func (i WebhookReceiverResponseArgs) ToWebhookReceiverResponseOutputWithContext(ctx context.Context) WebhookReceiverResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookReceiverResponseOutput)
}





type WebhookReceiverResponseArrayInput interface {
	pulumi.Input

	ToWebhookReceiverResponseArrayOutput() WebhookReceiverResponseArrayOutput
	ToWebhookReceiverResponseArrayOutputWithContext(context.Context) WebhookReceiverResponseArrayOutput
}

type WebhookReceiverResponseArray []WebhookReceiverResponseInput

func (WebhookReceiverResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebhookReceiverResponse)(nil)).Elem()
}

func (i WebhookReceiverResponseArray) ToWebhookReceiverResponseArrayOutput() WebhookReceiverResponseArrayOutput {
	return i.ToWebhookReceiverResponseArrayOutputWithContext(context.Background())
}

func (i WebhookReceiverResponseArray) ToWebhookReceiverResponseArrayOutputWithContext(ctx context.Context) WebhookReceiverResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookReceiverResponseArrayOutput)
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





type WebtestLocationAvailabilityCriteriaInput interface {
	pulumi.Input

	ToWebtestLocationAvailabilityCriteriaOutput() WebtestLocationAvailabilityCriteriaOutput
	ToWebtestLocationAvailabilityCriteriaOutputWithContext(context.Context) WebtestLocationAvailabilityCriteriaOutput
}

type WebtestLocationAvailabilityCriteriaArgs struct {
	ComponentId         pulumi.StringInput  `pulumi:"componentId"`
	FailedLocationCount pulumi.Float64Input `pulumi:"failedLocationCount"`
	OdataType           pulumi.StringInput  `pulumi:"odataType"`
	WebTestId           pulumi.StringInput  `pulumi:"webTestId"`
}

func (WebtestLocationAvailabilityCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebtestLocationAvailabilityCriteria)(nil)).Elem()
}

func (i WebtestLocationAvailabilityCriteriaArgs) ToWebtestLocationAvailabilityCriteriaOutput() WebtestLocationAvailabilityCriteriaOutput {
	return i.ToWebtestLocationAvailabilityCriteriaOutputWithContext(context.Background())
}

func (i WebtestLocationAvailabilityCriteriaArgs) ToWebtestLocationAvailabilityCriteriaOutputWithContext(ctx context.Context) WebtestLocationAvailabilityCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebtestLocationAvailabilityCriteriaOutput)
}

type WebtestLocationAvailabilityCriteriaOutput struct{ *pulumi.OutputState }

func (WebtestLocationAvailabilityCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebtestLocationAvailabilityCriteria)(nil)).Elem()
}

func (o WebtestLocationAvailabilityCriteriaOutput) ToWebtestLocationAvailabilityCriteriaOutput() WebtestLocationAvailabilityCriteriaOutput {
	return o
}

func (o WebtestLocationAvailabilityCriteriaOutput) ToWebtestLocationAvailabilityCriteriaOutputWithContext(ctx context.Context) WebtestLocationAvailabilityCriteriaOutput {
	return o
}

func (o WebtestLocationAvailabilityCriteriaOutput) ComponentId() pulumi.StringOutput {
	return o.ApplyT(func(v WebtestLocationAvailabilityCriteria) string { return v.ComponentId }).(pulumi.StringOutput)
}

func (o WebtestLocationAvailabilityCriteriaOutput) FailedLocationCount() pulumi.Float64Output {
	return o.ApplyT(func(v WebtestLocationAvailabilityCriteria) float64 { return v.FailedLocationCount }).(pulumi.Float64Output)
}

func (o WebtestLocationAvailabilityCriteriaOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v WebtestLocationAvailabilityCriteria) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o WebtestLocationAvailabilityCriteriaOutput) WebTestId() pulumi.StringOutput {
	return o.ApplyT(func(v WebtestLocationAvailabilityCriteria) string { return v.WebTestId }).(pulumi.StringOutput)
}

type WebtestLocationAvailabilityCriteriaResponse struct {
	ComponentId         string  `pulumi:"componentId"`
	FailedLocationCount float64 `pulumi:"failedLocationCount"`
	OdataType           string  `pulumi:"odataType"`
	WebTestId           string  `pulumi:"webTestId"`
}





type WebtestLocationAvailabilityCriteriaResponseInput interface {
	pulumi.Input

	ToWebtestLocationAvailabilityCriteriaResponseOutput() WebtestLocationAvailabilityCriteriaResponseOutput
	ToWebtestLocationAvailabilityCriteriaResponseOutputWithContext(context.Context) WebtestLocationAvailabilityCriteriaResponseOutput
}

type WebtestLocationAvailabilityCriteriaResponseArgs struct {
	ComponentId         pulumi.StringInput  `pulumi:"componentId"`
	FailedLocationCount pulumi.Float64Input `pulumi:"failedLocationCount"`
	OdataType           pulumi.StringInput  `pulumi:"odataType"`
	WebTestId           pulumi.StringInput  `pulumi:"webTestId"`
}

func (WebtestLocationAvailabilityCriteriaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebtestLocationAvailabilityCriteriaResponse)(nil)).Elem()
}

func (i WebtestLocationAvailabilityCriteriaResponseArgs) ToWebtestLocationAvailabilityCriteriaResponseOutput() WebtestLocationAvailabilityCriteriaResponseOutput {
	return i.ToWebtestLocationAvailabilityCriteriaResponseOutputWithContext(context.Background())
}

func (i WebtestLocationAvailabilityCriteriaResponseArgs) ToWebtestLocationAvailabilityCriteriaResponseOutputWithContext(ctx context.Context) WebtestLocationAvailabilityCriteriaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebtestLocationAvailabilityCriteriaResponseOutput)
}

type WebtestLocationAvailabilityCriteriaResponseOutput struct{ *pulumi.OutputState }

func (WebtestLocationAvailabilityCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebtestLocationAvailabilityCriteriaResponse)(nil)).Elem()
}

func (o WebtestLocationAvailabilityCriteriaResponseOutput) ToWebtestLocationAvailabilityCriteriaResponseOutput() WebtestLocationAvailabilityCriteriaResponseOutput {
	return o
}

func (o WebtestLocationAvailabilityCriteriaResponseOutput) ToWebtestLocationAvailabilityCriteriaResponseOutputWithContext(ctx context.Context) WebtestLocationAvailabilityCriteriaResponseOutput {
	return o
}

func (o WebtestLocationAvailabilityCriteriaResponseOutput) ComponentId() pulumi.StringOutput {
	return o.ApplyT(func(v WebtestLocationAvailabilityCriteriaResponse) string { return v.ComponentId }).(pulumi.StringOutput)
}

func (o WebtestLocationAvailabilityCriteriaResponseOutput) FailedLocationCount() pulumi.Float64Output {
	return o.ApplyT(func(v WebtestLocationAvailabilityCriteriaResponse) float64 { return v.FailedLocationCount }).(pulumi.Float64Output)
}

func (o WebtestLocationAvailabilityCriteriaResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v WebtestLocationAvailabilityCriteriaResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o WebtestLocationAvailabilityCriteriaResponseOutput) WebTestId() pulumi.StringOutput {
	return o.ApplyT(func(v WebtestLocationAvailabilityCriteriaResponse) string { return v.WebTestId }).(pulumi.StringOutput)
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





type WindowsEventLogDataSourceResponseInput interface {
	pulumi.Input

	ToWindowsEventLogDataSourceResponseOutput() WindowsEventLogDataSourceResponseOutput
	ToWindowsEventLogDataSourceResponseOutputWithContext(context.Context) WindowsEventLogDataSourceResponseOutput
}

type WindowsEventLogDataSourceResponseArgs struct {
	Name         pulumi.StringPtrInput   `pulumi:"name"`
	Streams      pulumi.StringArrayInput `pulumi:"streams"`
	XPathQueries pulumi.StringArrayInput `pulumi:"xPathQueries"`
}

func (WindowsEventLogDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsEventLogDataSourceResponse)(nil)).Elem()
}

func (i WindowsEventLogDataSourceResponseArgs) ToWindowsEventLogDataSourceResponseOutput() WindowsEventLogDataSourceResponseOutput {
	return i.ToWindowsEventLogDataSourceResponseOutputWithContext(context.Background())
}

func (i WindowsEventLogDataSourceResponseArgs) ToWindowsEventLogDataSourceResponseOutputWithContext(ctx context.Context) WindowsEventLogDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsEventLogDataSourceResponseOutput)
}





type WindowsEventLogDataSourceResponseArrayInput interface {
	pulumi.Input

	ToWindowsEventLogDataSourceResponseArrayOutput() WindowsEventLogDataSourceResponseArrayOutput
	ToWindowsEventLogDataSourceResponseArrayOutputWithContext(context.Context) WindowsEventLogDataSourceResponseArrayOutput
}

type WindowsEventLogDataSourceResponseArray []WindowsEventLogDataSourceResponseInput

func (WindowsEventLogDataSourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WindowsEventLogDataSourceResponse)(nil)).Elem()
}

func (i WindowsEventLogDataSourceResponseArray) ToWindowsEventLogDataSourceResponseArrayOutput() WindowsEventLogDataSourceResponseArrayOutput {
	return i.ToWindowsEventLogDataSourceResponseArrayOutputWithContext(context.Background())
}

func (i WindowsEventLogDataSourceResponseArray) ToWindowsEventLogDataSourceResponseArrayOutputWithContext(ctx context.Context) WindowsEventLogDataSourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsEventLogDataSourceResponseArrayOutput)
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





type WorkbookManagedIdentityResponseInput interface {
	pulumi.Input

	ToWorkbookManagedIdentityResponseOutput() WorkbookManagedIdentityResponseOutput
	ToWorkbookManagedIdentityResponseOutputWithContext(context.Context) WorkbookManagedIdentityResponseOutput
}

type WorkbookManagedIdentityResponseArgs struct {
	Type                   pulumi.StringPtrInput                          `pulumi:"type"`
	UserAssignedIdentities WorkbookUserAssignedIdentitiesResponsePtrInput `pulumi:"userAssignedIdentities"`
}

func (WorkbookManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookManagedIdentityResponse)(nil)).Elem()
}

func (i WorkbookManagedIdentityResponseArgs) ToWorkbookManagedIdentityResponseOutput() WorkbookManagedIdentityResponseOutput {
	return i.ToWorkbookManagedIdentityResponseOutputWithContext(context.Background())
}

func (i WorkbookManagedIdentityResponseArgs) ToWorkbookManagedIdentityResponseOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityResponseOutput)
}

func (i WorkbookManagedIdentityResponseArgs) ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput {
	return i.ToWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i WorkbookManagedIdentityResponseArgs) ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityResponseOutput).ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx)
}









type WorkbookManagedIdentityResponsePtrInput interface {
	pulumi.Input

	ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput
	ToWorkbookManagedIdentityResponsePtrOutputWithContext(context.Context) WorkbookManagedIdentityResponsePtrOutput
}

type workbookManagedIdentityResponsePtrType WorkbookManagedIdentityResponseArgs

func WorkbookManagedIdentityResponsePtr(v *WorkbookManagedIdentityResponseArgs) WorkbookManagedIdentityResponsePtrInput {
	return (*workbookManagedIdentityResponsePtrType)(v)
}

func (*workbookManagedIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookManagedIdentityResponse)(nil)).Elem()
}

func (i *workbookManagedIdentityResponsePtrType) ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput {
	return i.ToWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *workbookManagedIdentityResponsePtrType) ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityResponsePtrOutput)
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

func (o WorkbookManagedIdentityResponseOutput) ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput {
	return o.ToWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (o WorkbookManagedIdentityResponseOutput) ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkbookManagedIdentityResponse) *WorkbookManagedIdentityResponse {
		return &v
	}).(WorkbookManagedIdentityResponsePtrOutput)
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





type WorkbookTemplateGalleryResponseInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryResponseOutput() WorkbookTemplateGalleryResponseOutput
	ToWorkbookTemplateGalleryResponseOutputWithContext(context.Context) WorkbookTemplateGalleryResponseOutput
}

type WorkbookTemplateGalleryResponseArgs struct {
	Category     pulumi.StringPtrInput `pulumi:"category"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Order        pulumi.IntPtrInput    `pulumi:"order"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (WorkbookTemplateGalleryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateGalleryResponseArgs) ToWorkbookTemplateGalleryResponseOutput() WorkbookTemplateGalleryResponseOutput {
	return i.ToWorkbookTemplateGalleryResponseOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryResponseArgs) ToWorkbookTemplateGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryResponseOutput)
}





type WorkbookTemplateGalleryResponseArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryResponseArrayOutput() WorkbookTemplateGalleryResponseArrayOutput
	ToWorkbookTemplateGalleryResponseArrayOutputWithContext(context.Context) WorkbookTemplateGalleryResponseArrayOutput
}

type WorkbookTemplateGalleryResponseArray []WorkbookTemplateGalleryResponseInput

func (WorkbookTemplateGalleryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateGalleryResponseArray) ToWorkbookTemplateGalleryResponseArrayOutput() WorkbookTemplateGalleryResponseArrayOutput {
	return i.ToWorkbookTemplateGalleryResponseArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryResponseArray) ToWorkbookTemplateGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryResponseArrayOutput)
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





type WorkbookTemplateLocalizedGalleryResponseInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryResponseOutput() WorkbookTemplateLocalizedGalleryResponseOutput
	ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryResponseOutput
}

type WorkbookTemplateLocalizedGalleryResponseArgs struct {
	Galleries    WorkbookTemplateGalleryResponseArrayInput `pulumi:"galleries"`
	TemplateData pulumi.Input                              `pulumi:"templateData"`
}

func (WorkbookTemplateLocalizedGalleryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryResponseArgs) ToWorkbookTemplateLocalizedGalleryResponseOutput() WorkbookTemplateLocalizedGalleryResponseOutput {
	return i.ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryResponseArgs) ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryResponseOutput)
}





type WorkbookTemplateLocalizedGalleryResponseArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryResponseArrayOutput() WorkbookTemplateLocalizedGalleryResponseArrayOutput
	ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryResponseArrayOutput
}

type WorkbookTemplateLocalizedGalleryResponseArray []WorkbookTemplateLocalizedGalleryResponseInput

func (WorkbookTemplateLocalizedGalleryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryResponseArray) ToWorkbookTemplateLocalizedGalleryResponseArrayOutput() WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return i.ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryResponseArray) ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryResponseArrayOutput)
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





type WorkbookUserAssignedIdentitiesResponseInput interface {
	pulumi.Input

	ToWorkbookUserAssignedIdentitiesResponseOutput() WorkbookUserAssignedIdentitiesResponseOutput
	ToWorkbookUserAssignedIdentitiesResponseOutputWithContext(context.Context) WorkbookUserAssignedIdentitiesResponseOutput
}

type WorkbookUserAssignedIdentitiesResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
}

func (WorkbookUserAssignedIdentitiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (i WorkbookUserAssignedIdentitiesResponseArgs) ToWorkbookUserAssignedIdentitiesResponseOutput() WorkbookUserAssignedIdentitiesResponseOutput {
	return i.ToWorkbookUserAssignedIdentitiesResponseOutputWithContext(context.Background())
}

func (i WorkbookUserAssignedIdentitiesResponseArgs) ToWorkbookUserAssignedIdentitiesResponseOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookUserAssignedIdentitiesResponseOutput)
}

func (i WorkbookUserAssignedIdentitiesResponseArgs) ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return i.ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (i WorkbookUserAssignedIdentitiesResponseArgs) ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookUserAssignedIdentitiesResponseOutput).ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx)
}









type WorkbookUserAssignedIdentitiesResponsePtrInput interface {
	pulumi.Input

	ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput
	ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput
}

type workbookUserAssignedIdentitiesResponsePtrType WorkbookUserAssignedIdentitiesResponseArgs

func WorkbookUserAssignedIdentitiesResponsePtr(v *WorkbookUserAssignedIdentitiesResponseArgs) WorkbookUserAssignedIdentitiesResponsePtrInput {
	return (*workbookUserAssignedIdentitiesResponsePtrType)(v)
}

func (*workbookUserAssignedIdentitiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (i *workbookUserAssignedIdentitiesResponsePtrType) ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return i.ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (i *workbookUserAssignedIdentitiesResponsePtrType) ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookUserAssignedIdentitiesResponsePtrOutput)
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

func (o WorkbookUserAssignedIdentitiesResponseOutput) ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkbookUserAssignedIdentitiesResponse) *WorkbookUserAssignedIdentitiesResponse {
		return &v
	}).(WorkbookUserAssignedIdentitiesResponsePtrOutput)
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
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryArray)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArrayMap) ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return i.ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArrayMap) ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryArrayMapOutput)
}

type WorkbookTemplateLocalizedGalleryArrayMapOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkbookTemplateLocalizedGalleryArray)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) MapIndex(k pulumi.StringInput) WorkbookTemplateLocalizedGalleryArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGalleryArray {
		return vs[0].(map[string]WorkbookTemplateLocalizedGalleryArray)[vs[1].(string)]
	}).(WorkbookTemplateLocalizedGalleryArrayOutput)
}





type WorkbookTemplateLocalizedGalleryArrayMapInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput
	ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput
}

type WorkbookTemplateLocalizedGalleryResponseArrayMap map[string]WorkbookTemplateLocalizedGalleryResponseArrayInput

func (WorkbookTemplateLocalizedGalleryResponseArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponseArray)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryResponseArrayMap) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return i.ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryResponseArrayMap) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryResponseArrayMapOutput)
}

type WorkbookTemplateLocalizedGalleryResponseArrayMapOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkbookTemplateLocalizedGalleryResponseArray)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) MapIndex(k pulumi.StringInput) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGalleryResponseArray {
		return vs[0].(map[string]WorkbookTemplateLocalizedGalleryResponseArray)[vs[1].(string)]
	}).(WorkbookTemplateLocalizedGalleryResponseArrayOutput)
}





type WorkbookTemplateLocalizedGalleryResponseArrayMapInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput
	ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput
}

func init() {
	pulumi.RegisterOutputType(ActionGroupTypeOutput{})
	pulumi.RegisterOutputType(ActionGroupTypeArrayOutput{})
	pulumi.RegisterOutputType(ActionGroupResponseOutput{})
	pulumi.RegisterOutputType(ActionGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(ActionListOutput{})
	pulumi.RegisterOutputType(ActionListPtrOutput{})
	pulumi.RegisterOutputType(ActionListResponseOutput{})
	pulumi.RegisterOutputType(ActionListResponsePtrOutput{})
	pulumi.RegisterOutputType(AlertRuleAllOfConditionOutput{})
	pulumi.RegisterOutputType(AlertRuleAllOfConditionPtrOutput{})
	pulumi.RegisterOutputType(AlertRuleAllOfConditionResponseOutput{})
	pulumi.RegisterOutputType(AlertRuleAllOfConditionResponsePtrOutput{})
	pulumi.RegisterOutputType(AlertRuleAnyOfOrLeafConditionOutput{})
	pulumi.RegisterOutputType(AlertRuleAnyOfOrLeafConditionArrayOutput{})
	pulumi.RegisterOutputType(AlertRuleAnyOfOrLeafConditionResponseOutput{})
	pulumi.RegisterOutputType(AlertRuleAnyOfOrLeafConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(AlertRuleLeafConditionOutput{})
	pulumi.RegisterOutputType(AlertRuleLeafConditionArrayOutput{})
	pulumi.RegisterOutputType(AlertRuleLeafConditionResponseOutput{})
	pulumi.RegisterOutputType(AlertRuleLeafConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(AlertingActionOutput{})
	pulumi.RegisterOutputType(AlertingActionResponseOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentAnalyticsItemPropertiesResponsePtrOutput{})
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
	pulumi.RegisterOutputType(AzNsActionGroupOutput{})
	pulumi.RegisterOutputType(AzNsActionGroupPtrOutput{})
	pulumi.RegisterOutputType(AzNsActionGroupResponseOutput{})
	pulumi.RegisterOutputType(AzNsActionGroupResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverArrayOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverResponseOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverArrayOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverResponseOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(CriteriaOutput{})
	pulumi.RegisterOutputType(CriteriaArrayOutput{})
	pulumi.RegisterOutputType(CriteriaResponseOutput{})
	pulumi.RegisterOutputType(CriteriaResponseArrayOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointNetworkAclsOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointNetworkAclsPtrOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResourceResponseSystemDataOutput{})
	pulumi.RegisterOutputType(DataCollectionEndpointResourceResponseSystemDataPtrOutput{})
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
	pulumi.RegisterOutputType(DimensionOutput{})
	pulumi.RegisterOutputType(DimensionArrayOutput{})
	pulumi.RegisterOutputType(DimensionResponseOutput{})
	pulumi.RegisterOutputType(DimensionResponseArrayOutput{})
	pulumi.RegisterOutputType(DynamicMetricCriteriaOutput{})
	pulumi.RegisterOutputType(DynamicMetricCriteriaResponseOutput{})
	pulumi.RegisterOutputType(DynamicThresholdFailingPeriodsOutput{})
	pulumi.RegisterOutputType(DynamicThresholdFailingPeriodsResponseOutput{})
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
	pulumi.RegisterOutputType(LocationThresholdRuleConditionOutput{})
	pulumi.RegisterOutputType(LocationThresholdRuleConditionResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationArrayOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsDestinationResponseArrayOutput{})
	pulumi.RegisterOutputType(LogMetricTriggerOutput{})
	pulumi.RegisterOutputType(LogMetricTriggerPtrOutput{})
	pulumi.RegisterOutputType(LogMetricTriggerResponseOutput{})
	pulumi.RegisterOutputType(LogMetricTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(LogSettingsOutput{})
	pulumi.RegisterOutputType(LogSettingsArrayOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(LogToMetricActionOutput{})
	pulumi.RegisterOutputType(LogToMetricActionResponseOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverArrayOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverResponseOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionPtrOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionResponseOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementEventRuleConditionOutput{})
	pulumi.RegisterOutputType(ManagementEventRuleConditionResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricAlertActionOutput{})
	pulumi.RegisterOutputType(MetricAlertActionArrayOutput{})
	pulumi.RegisterOutputType(MetricAlertActionResponseOutput{})
	pulumi.RegisterOutputType(MetricAlertActionResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricAlertMultipleResourceMultipleMetricCriteriaOutput{})
	pulumi.RegisterOutputType(MetricAlertMultipleResourceMultipleMetricCriteriaResponseOutput{})
	pulumi.RegisterOutputType(MetricAlertSingleResourceMultipleMetricCriteriaOutput{})
	pulumi.RegisterOutputType(MetricAlertSingleResourceMultipleMetricCriteriaResponseOutput{})
	pulumi.RegisterOutputType(MetricCriteriaOutput{})
	pulumi.RegisterOutputType(MetricCriteriaArrayOutput{})
	pulumi.RegisterOutputType(MetricCriteriaResponseOutput{})
	pulumi.RegisterOutputType(MetricCriteriaResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricDimensionOutput{})
	pulumi.RegisterOutputType(MetricDimensionArrayOutput{})
	pulumi.RegisterOutputType(MetricDimensionResponseOutput{})
	pulumi.RegisterOutputType(MetricDimensionResponseArrayOutput{})
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
	pulumi.RegisterOutputType(RuleEmailActionOutput{})
	pulumi.RegisterOutputType(RuleEmailActionResponseOutput{})
	pulumi.RegisterOutputType(RuleManagementEventClaimsDataSourceOutput{})
	pulumi.RegisterOutputType(RuleManagementEventClaimsDataSourcePtrOutput{})
	pulumi.RegisterOutputType(RuleManagementEventClaimsDataSourceResponseOutput{})
	pulumi.RegisterOutputType(RuleManagementEventClaimsDataSourceResponsePtrOutput{})
	pulumi.RegisterOutputType(RuleManagementEventDataSourceOutput{})
	pulumi.RegisterOutputType(RuleManagementEventDataSourceResponseOutput{})
	pulumi.RegisterOutputType(RuleMetricDataSourceOutput{})
	pulumi.RegisterOutputType(RuleMetricDataSourceResponseOutput{})
	pulumi.RegisterOutputType(RuleWebhookActionOutput{})
	pulumi.RegisterOutputType(RuleWebhookActionResponseOutput{})
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
	pulumi.RegisterOutputType(SourcePtrOutput{})
	pulumi.RegisterOutputType(SourceResponseOutput{})
	pulumi.RegisterOutputType(SourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsResponseOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceArrayOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceResponseOutput{})
	pulumi.RegisterOutputType(SyslogDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(ThresholdRuleConditionOutput{})
	pulumi.RegisterOutputType(ThresholdRuleConditionResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowOutput{})
	pulumi.RegisterOutputType(TimeWindowPtrOutput{})
	pulumi.RegisterOutputType(TimeWindowResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(TriggerConditionOutput{})
	pulumi.RegisterOutputType(TriggerConditionResponseOutput{})
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
	pulumi.RegisterOutputType(WebtestLocationAvailabilityCriteriaOutput{})
	pulumi.RegisterOutputType(WebtestLocationAvailabilityCriteriaResponseOutput{})
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
