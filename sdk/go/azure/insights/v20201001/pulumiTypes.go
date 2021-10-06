


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionGroup struct {
	ActionGroupId     string            `pulumi:"actionGroupId"`
	WebhookProperties map[string]string `pulumi:"webhookProperties"`
}





type ActionGroupInput interface {
	pulumi.Input

	ToActionGroupOutput() ActionGroupOutput
	ToActionGroupOutputWithContext(context.Context) ActionGroupOutput
}

type ActionGroupArgs struct {
	ActionGroupId     pulumi.StringInput    `pulumi:"actionGroupId"`
	WebhookProperties pulumi.StringMapInput `pulumi:"webhookProperties"`
}

func (ActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroup)(nil)).Elem()
}

func (i ActionGroupArgs) ToActionGroupOutput() ActionGroupOutput {
	return i.ToActionGroupOutputWithContext(context.Background())
}

func (i ActionGroupArgs) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupOutput)
}





type ActionGroupArrayInput interface {
	pulumi.Input

	ToActionGroupArrayOutput() ActionGroupArrayOutput
	ToActionGroupArrayOutputWithContext(context.Context) ActionGroupArrayOutput
}

type ActionGroupArray []ActionGroupInput

func (ActionGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionGroup)(nil)).Elem()
}

func (i ActionGroupArray) ToActionGroupArrayOutput() ActionGroupArrayOutput {
	return i.ToActionGroupArrayOutputWithContext(context.Background())
}

func (i ActionGroupArray) ToActionGroupArrayOutputWithContext(ctx context.Context) ActionGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupArrayOutput)
}

type ActionGroupOutput struct{ *pulumi.OutputState }

func (ActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroup)(nil)).Elem()
}

func (o ActionGroupOutput) ToActionGroupOutput() ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) ActionGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroup) string { return v.ActionGroupId }).(pulumi.StringOutput)
}

func (o ActionGroupOutput) WebhookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ActionGroup) map[string]string { return v.WebhookProperties }).(pulumi.StringMapOutput)
}

type ActionGroupArrayOutput struct{ *pulumi.OutputState }

func (ActionGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActionGroup)(nil)).Elem()
}

func (o ActionGroupArrayOutput) ToActionGroupArrayOutput() ActionGroupArrayOutput {
	return o
}

func (o ActionGroupArrayOutput) ToActionGroupArrayOutputWithContext(ctx context.Context) ActionGroupArrayOutput {
	return o
}

func (o ActionGroupArrayOutput) Index(i pulumi.IntInput) ActionGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActionGroup {
		return vs[0].([]ActionGroup)[vs[1].(int)]
	}).(ActionGroupOutput)
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
	ActionGroups []ActionGroup `pulumi:"actionGroups"`
}





type ActionListInput interface {
	pulumi.Input

	ToActionListOutput() ActionListOutput
	ToActionListOutputWithContext(context.Context) ActionListOutput
}

type ActionListArgs struct {
	ActionGroups ActionGroupArrayInput `pulumi:"actionGroups"`
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

func (o ActionListOutput) ActionGroups() ActionGroupArrayOutput {
	return o.ApplyT(func(v ActionList) []ActionGroup { return v.ActionGroups }).(ActionGroupArrayOutput)
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

func (o ActionListPtrOutput) ActionGroups() ActionGroupArrayOutput {
	return o.ApplyT(func(v *ActionList) []ActionGroup {
		if v == nil {
			return nil
		}
		return v.ActionGroups
	}).(ActionGroupArrayOutput)
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

func init() {
	pulumi.RegisterOutputType(ActionGroupOutput{})
	pulumi.RegisterOutputType(ActionGroupArrayOutput{})
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
}
