


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActivityLogAlertActionGroup struct {
	ActionGroupId     string            `pulumi:"actionGroupId"`
	WebhookProperties map[string]string `pulumi:"webhookProperties"`
}





type ActivityLogAlertActionGroupInput interface {
	pulumi.Input

	ToActivityLogAlertActionGroupOutput() ActivityLogAlertActionGroupOutput
	ToActivityLogAlertActionGroupOutputWithContext(context.Context) ActivityLogAlertActionGroupOutput
}

type ActivityLogAlertActionGroupArgs struct {
	ActionGroupId     pulumi.StringInput    `pulumi:"actionGroupId"`
	WebhookProperties pulumi.StringMapInput `pulumi:"webhookProperties"`
}

func (ActivityLogAlertActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertActionGroup)(nil)).Elem()
}

func (i ActivityLogAlertActionGroupArgs) ToActivityLogAlertActionGroupOutput() ActivityLogAlertActionGroupOutput {
	return i.ToActivityLogAlertActionGroupOutputWithContext(context.Background())
}

func (i ActivityLogAlertActionGroupArgs) ToActivityLogAlertActionGroupOutputWithContext(ctx context.Context) ActivityLogAlertActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertActionGroupOutput)
}





type ActivityLogAlertActionGroupArrayInput interface {
	pulumi.Input

	ToActivityLogAlertActionGroupArrayOutput() ActivityLogAlertActionGroupArrayOutput
	ToActivityLogAlertActionGroupArrayOutputWithContext(context.Context) ActivityLogAlertActionGroupArrayOutput
}

type ActivityLogAlertActionGroupArray []ActivityLogAlertActionGroupInput

func (ActivityLogAlertActionGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActivityLogAlertActionGroup)(nil)).Elem()
}

func (i ActivityLogAlertActionGroupArray) ToActivityLogAlertActionGroupArrayOutput() ActivityLogAlertActionGroupArrayOutput {
	return i.ToActivityLogAlertActionGroupArrayOutputWithContext(context.Background())
}

func (i ActivityLogAlertActionGroupArray) ToActivityLogAlertActionGroupArrayOutputWithContext(ctx context.Context) ActivityLogAlertActionGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertActionGroupArrayOutput)
}

type ActivityLogAlertActionGroupOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertActionGroup)(nil)).Elem()
}

func (o ActivityLogAlertActionGroupOutput) ToActivityLogAlertActionGroupOutput() ActivityLogAlertActionGroupOutput {
	return o
}

func (o ActivityLogAlertActionGroupOutput) ToActivityLogAlertActionGroupOutputWithContext(ctx context.Context) ActivityLogAlertActionGroupOutput {
	return o
}

func (o ActivityLogAlertActionGroupOutput) ActionGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityLogAlertActionGroup) string { return v.ActionGroupId }).(pulumi.StringOutput)
}

func (o ActivityLogAlertActionGroupOutput) WebhookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ActivityLogAlertActionGroup) map[string]string { return v.WebhookProperties }).(pulumi.StringMapOutput)
}

type ActivityLogAlertActionGroupArrayOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertActionGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActivityLogAlertActionGroup)(nil)).Elem()
}

func (o ActivityLogAlertActionGroupArrayOutput) ToActivityLogAlertActionGroupArrayOutput() ActivityLogAlertActionGroupArrayOutput {
	return o
}

func (o ActivityLogAlertActionGroupArrayOutput) ToActivityLogAlertActionGroupArrayOutputWithContext(ctx context.Context) ActivityLogAlertActionGroupArrayOutput {
	return o
}

func (o ActivityLogAlertActionGroupArrayOutput) Index(i pulumi.IntInput) ActivityLogAlertActionGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActivityLogAlertActionGroup {
		return vs[0].([]ActivityLogAlertActionGroup)[vs[1].(int)]
	}).(ActivityLogAlertActionGroupOutput)
}

type ActivityLogAlertActionGroupResponse struct {
	ActionGroupId     string            `pulumi:"actionGroupId"`
	WebhookProperties map[string]string `pulumi:"webhookProperties"`
}

type ActivityLogAlertActionGroupResponseOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertActionGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertActionGroupResponse)(nil)).Elem()
}

func (o ActivityLogAlertActionGroupResponseOutput) ToActivityLogAlertActionGroupResponseOutput() ActivityLogAlertActionGroupResponseOutput {
	return o
}

func (o ActivityLogAlertActionGroupResponseOutput) ToActivityLogAlertActionGroupResponseOutputWithContext(ctx context.Context) ActivityLogAlertActionGroupResponseOutput {
	return o
}

func (o ActivityLogAlertActionGroupResponseOutput) ActionGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityLogAlertActionGroupResponse) string { return v.ActionGroupId }).(pulumi.StringOutput)
}

func (o ActivityLogAlertActionGroupResponseOutput) WebhookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ActivityLogAlertActionGroupResponse) map[string]string { return v.WebhookProperties }).(pulumi.StringMapOutput)
}

type ActivityLogAlertActionGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertActionGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActivityLogAlertActionGroupResponse)(nil)).Elem()
}

func (o ActivityLogAlertActionGroupResponseArrayOutput) ToActivityLogAlertActionGroupResponseArrayOutput() ActivityLogAlertActionGroupResponseArrayOutput {
	return o
}

func (o ActivityLogAlertActionGroupResponseArrayOutput) ToActivityLogAlertActionGroupResponseArrayOutputWithContext(ctx context.Context) ActivityLogAlertActionGroupResponseArrayOutput {
	return o
}

func (o ActivityLogAlertActionGroupResponseArrayOutput) Index(i pulumi.IntInput) ActivityLogAlertActionGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActivityLogAlertActionGroupResponse {
		return vs[0].([]ActivityLogAlertActionGroupResponse)[vs[1].(int)]
	}).(ActivityLogAlertActionGroupResponseOutput)
}

type ActivityLogAlertActionList struct {
	ActionGroups []ActivityLogAlertActionGroup `pulumi:"actionGroups"`
}





type ActivityLogAlertActionListInput interface {
	pulumi.Input

	ToActivityLogAlertActionListOutput() ActivityLogAlertActionListOutput
	ToActivityLogAlertActionListOutputWithContext(context.Context) ActivityLogAlertActionListOutput
}

type ActivityLogAlertActionListArgs struct {
	ActionGroups ActivityLogAlertActionGroupArrayInput `pulumi:"actionGroups"`
}

func (ActivityLogAlertActionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertActionList)(nil)).Elem()
}

func (i ActivityLogAlertActionListArgs) ToActivityLogAlertActionListOutput() ActivityLogAlertActionListOutput {
	return i.ToActivityLogAlertActionListOutputWithContext(context.Background())
}

func (i ActivityLogAlertActionListArgs) ToActivityLogAlertActionListOutputWithContext(ctx context.Context) ActivityLogAlertActionListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertActionListOutput)
}

type ActivityLogAlertActionListOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertActionListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertActionList)(nil)).Elem()
}

func (o ActivityLogAlertActionListOutput) ToActivityLogAlertActionListOutput() ActivityLogAlertActionListOutput {
	return o
}

func (o ActivityLogAlertActionListOutput) ToActivityLogAlertActionListOutputWithContext(ctx context.Context) ActivityLogAlertActionListOutput {
	return o
}

func (o ActivityLogAlertActionListOutput) ActionGroups() ActivityLogAlertActionGroupArrayOutput {
	return o.ApplyT(func(v ActivityLogAlertActionList) []ActivityLogAlertActionGroup { return v.ActionGroups }).(ActivityLogAlertActionGroupArrayOutput)
}

type ActivityLogAlertActionListResponse struct {
	ActionGroups []ActivityLogAlertActionGroupResponse `pulumi:"actionGroups"`
}

type ActivityLogAlertActionListResponseOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertActionListResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertActionListResponse)(nil)).Elem()
}

func (o ActivityLogAlertActionListResponseOutput) ToActivityLogAlertActionListResponseOutput() ActivityLogAlertActionListResponseOutput {
	return o
}

func (o ActivityLogAlertActionListResponseOutput) ToActivityLogAlertActionListResponseOutputWithContext(ctx context.Context) ActivityLogAlertActionListResponseOutput {
	return o
}

func (o ActivityLogAlertActionListResponseOutput) ActionGroups() ActivityLogAlertActionGroupResponseArrayOutput {
	return o.ApplyT(func(v ActivityLogAlertActionListResponse) []ActivityLogAlertActionGroupResponse {
		return v.ActionGroups
	}).(ActivityLogAlertActionGroupResponseArrayOutput)
}

type ActivityLogAlertAllOfCondition struct {
	AllOf []ActivityLogAlertLeafCondition `pulumi:"allOf"`
}





type ActivityLogAlertAllOfConditionInput interface {
	pulumi.Input

	ToActivityLogAlertAllOfConditionOutput() ActivityLogAlertAllOfConditionOutput
	ToActivityLogAlertAllOfConditionOutputWithContext(context.Context) ActivityLogAlertAllOfConditionOutput
}

type ActivityLogAlertAllOfConditionArgs struct {
	AllOf ActivityLogAlertLeafConditionArrayInput `pulumi:"allOf"`
}

func (ActivityLogAlertAllOfConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertAllOfCondition)(nil)).Elem()
}

func (i ActivityLogAlertAllOfConditionArgs) ToActivityLogAlertAllOfConditionOutput() ActivityLogAlertAllOfConditionOutput {
	return i.ToActivityLogAlertAllOfConditionOutputWithContext(context.Background())
}

func (i ActivityLogAlertAllOfConditionArgs) ToActivityLogAlertAllOfConditionOutputWithContext(ctx context.Context) ActivityLogAlertAllOfConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertAllOfConditionOutput)
}

type ActivityLogAlertAllOfConditionOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertAllOfConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertAllOfCondition)(nil)).Elem()
}

func (o ActivityLogAlertAllOfConditionOutput) ToActivityLogAlertAllOfConditionOutput() ActivityLogAlertAllOfConditionOutput {
	return o
}

func (o ActivityLogAlertAllOfConditionOutput) ToActivityLogAlertAllOfConditionOutputWithContext(ctx context.Context) ActivityLogAlertAllOfConditionOutput {
	return o
}

func (o ActivityLogAlertAllOfConditionOutput) AllOf() ActivityLogAlertLeafConditionArrayOutput {
	return o.ApplyT(func(v ActivityLogAlertAllOfCondition) []ActivityLogAlertLeafCondition { return v.AllOf }).(ActivityLogAlertLeafConditionArrayOutput)
}

type ActivityLogAlertAllOfConditionResponse struct {
	AllOf []ActivityLogAlertLeafConditionResponse `pulumi:"allOf"`
}

type ActivityLogAlertAllOfConditionResponseOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertAllOfConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertAllOfConditionResponse)(nil)).Elem()
}

func (o ActivityLogAlertAllOfConditionResponseOutput) ToActivityLogAlertAllOfConditionResponseOutput() ActivityLogAlertAllOfConditionResponseOutput {
	return o
}

func (o ActivityLogAlertAllOfConditionResponseOutput) ToActivityLogAlertAllOfConditionResponseOutputWithContext(ctx context.Context) ActivityLogAlertAllOfConditionResponseOutput {
	return o
}

func (o ActivityLogAlertAllOfConditionResponseOutput) AllOf() ActivityLogAlertLeafConditionResponseArrayOutput {
	return o.ApplyT(func(v ActivityLogAlertAllOfConditionResponse) []ActivityLogAlertLeafConditionResponse { return v.AllOf }).(ActivityLogAlertLeafConditionResponseArrayOutput)
}

type ActivityLogAlertLeafCondition struct {
	Equals string `pulumi:"equals"`
	Field  string `pulumi:"field"`
}





type ActivityLogAlertLeafConditionInput interface {
	pulumi.Input

	ToActivityLogAlertLeafConditionOutput() ActivityLogAlertLeafConditionOutput
	ToActivityLogAlertLeafConditionOutputWithContext(context.Context) ActivityLogAlertLeafConditionOutput
}

type ActivityLogAlertLeafConditionArgs struct {
	Equals pulumi.StringInput `pulumi:"equals"`
	Field  pulumi.StringInput `pulumi:"field"`
}

func (ActivityLogAlertLeafConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertLeafCondition)(nil)).Elem()
}

func (i ActivityLogAlertLeafConditionArgs) ToActivityLogAlertLeafConditionOutput() ActivityLogAlertLeafConditionOutput {
	return i.ToActivityLogAlertLeafConditionOutputWithContext(context.Background())
}

func (i ActivityLogAlertLeafConditionArgs) ToActivityLogAlertLeafConditionOutputWithContext(ctx context.Context) ActivityLogAlertLeafConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertLeafConditionOutput)
}





type ActivityLogAlertLeafConditionArrayInput interface {
	pulumi.Input

	ToActivityLogAlertLeafConditionArrayOutput() ActivityLogAlertLeafConditionArrayOutput
	ToActivityLogAlertLeafConditionArrayOutputWithContext(context.Context) ActivityLogAlertLeafConditionArrayOutput
}

type ActivityLogAlertLeafConditionArray []ActivityLogAlertLeafConditionInput

func (ActivityLogAlertLeafConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActivityLogAlertLeafCondition)(nil)).Elem()
}

func (i ActivityLogAlertLeafConditionArray) ToActivityLogAlertLeafConditionArrayOutput() ActivityLogAlertLeafConditionArrayOutput {
	return i.ToActivityLogAlertLeafConditionArrayOutputWithContext(context.Background())
}

func (i ActivityLogAlertLeafConditionArray) ToActivityLogAlertLeafConditionArrayOutputWithContext(ctx context.Context) ActivityLogAlertLeafConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityLogAlertLeafConditionArrayOutput)
}

type ActivityLogAlertLeafConditionOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertLeafConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertLeafCondition)(nil)).Elem()
}

func (o ActivityLogAlertLeafConditionOutput) ToActivityLogAlertLeafConditionOutput() ActivityLogAlertLeafConditionOutput {
	return o
}

func (o ActivityLogAlertLeafConditionOutput) ToActivityLogAlertLeafConditionOutputWithContext(ctx context.Context) ActivityLogAlertLeafConditionOutput {
	return o
}

func (o ActivityLogAlertLeafConditionOutput) Equals() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityLogAlertLeafCondition) string { return v.Equals }).(pulumi.StringOutput)
}

func (o ActivityLogAlertLeafConditionOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityLogAlertLeafCondition) string { return v.Field }).(pulumi.StringOutput)
}

type ActivityLogAlertLeafConditionArrayOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertLeafConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActivityLogAlertLeafCondition)(nil)).Elem()
}

func (o ActivityLogAlertLeafConditionArrayOutput) ToActivityLogAlertLeafConditionArrayOutput() ActivityLogAlertLeafConditionArrayOutput {
	return o
}

func (o ActivityLogAlertLeafConditionArrayOutput) ToActivityLogAlertLeafConditionArrayOutputWithContext(ctx context.Context) ActivityLogAlertLeafConditionArrayOutput {
	return o
}

func (o ActivityLogAlertLeafConditionArrayOutput) Index(i pulumi.IntInput) ActivityLogAlertLeafConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActivityLogAlertLeafCondition {
		return vs[0].([]ActivityLogAlertLeafCondition)[vs[1].(int)]
	}).(ActivityLogAlertLeafConditionOutput)
}

type ActivityLogAlertLeafConditionResponse struct {
	Equals string `pulumi:"equals"`
	Field  string `pulumi:"field"`
}

type ActivityLogAlertLeafConditionResponseOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertLeafConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityLogAlertLeafConditionResponse)(nil)).Elem()
}

func (o ActivityLogAlertLeafConditionResponseOutput) ToActivityLogAlertLeafConditionResponseOutput() ActivityLogAlertLeafConditionResponseOutput {
	return o
}

func (o ActivityLogAlertLeafConditionResponseOutput) ToActivityLogAlertLeafConditionResponseOutputWithContext(ctx context.Context) ActivityLogAlertLeafConditionResponseOutput {
	return o
}

func (o ActivityLogAlertLeafConditionResponseOutput) Equals() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityLogAlertLeafConditionResponse) string { return v.Equals }).(pulumi.StringOutput)
}

func (o ActivityLogAlertLeafConditionResponseOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityLogAlertLeafConditionResponse) string { return v.Field }).(pulumi.StringOutput)
}

type ActivityLogAlertLeafConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (ActivityLogAlertLeafConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ActivityLogAlertLeafConditionResponse)(nil)).Elem()
}

func (o ActivityLogAlertLeafConditionResponseArrayOutput) ToActivityLogAlertLeafConditionResponseArrayOutput() ActivityLogAlertLeafConditionResponseArrayOutput {
	return o
}

func (o ActivityLogAlertLeafConditionResponseArrayOutput) ToActivityLogAlertLeafConditionResponseArrayOutputWithContext(ctx context.Context) ActivityLogAlertLeafConditionResponseArrayOutput {
	return o
}

func (o ActivityLogAlertLeafConditionResponseArrayOutput) Index(i pulumi.IntInput) ActivityLogAlertLeafConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ActivityLogAlertLeafConditionResponse {
		return vs[0].([]ActivityLogAlertLeafConditionResponse)[vs[1].(int)]
	}).(ActivityLogAlertLeafConditionResponseOutput)
}

type AutomationRunbookReceiver struct {
	AutomationAccountId string  `pulumi:"automationAccountId"`
	IsGlobalRunbook     bool    `pulumi:"isGlobalRunbook"`
	Name                *string `pulumi:"name"`
	RunbookName         string  `pulumi:"runbookName"`
	ServiceUri          *string `pulumi:"serviceUri"`
	WebhookResourceId   string  `pulumi:"webhookResourceId"`
}





type AutomationRunbookReceiverInput interface {
	pulumi.Input

	ToAutomationRunbookReceiverOutput() AutomationRunbookReceiverOutput
	ToAutomationRunbookReceiverOutputWithContext(context.Context) AutomationRunbookReceiverOutput
}

type AutomationRunbookReceiverArgs struct {
	AutomationAccountId pulumi.StringInput    `pulumi:"automationAccountId"`
	IsGlobalRunbook     pulumi.BoolInput      `pulumi:"isGlobalRunbook"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	RunbookName         pulumi.StringInput    `pulumi:"runbookName"`
	ServiceUri          pulumi.StringPtrInput `pulumi:"serviceUri"`
	WebhookResourceId   pulumi.StringInput    `pulumi:"webhookResourceId"`
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
	AutomationAccountId string  `pulumi:"automationAccountId"`
	IsGlobalRunbook     bool    `pulumi:"isGlobalRunbook"`
	Name                *string `pulumi:"name"`
	RunbookName         string  `pulumi:"runbookName"`
	ServiceUri          *string `pulumi:"serviceUri"`
	WebhookResourceId   string  `pulumi:"webhookResourceId"`
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

type EmailReceiver struct {
	EmailAddress string `pulumi:"emailAddress"`
	Name         string `pulumi:"name"`
}





type EmailReceiverInput interface {
	pulumi.Input

	ToEmailReceiverOutput() EmailReceiverOutput
	ToEmailReceiverOutputWithContext(context.Context) EmailReceiverOutput
}

type EmailReceiverArgs struct {
	EmailAddress pulumi.StringInput `pulumi:"emailAddress"`
	Name         pulumi.StringInput `pulumi:"name"`
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
	EmailAddress string `pulumi:"emailAddress"`
	Name         string `pulumi:"name"`
	Status       string `pulumi:"status"`
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

type WebhookReceiver struct {
	Name       string `pulumi:"name"`
	ServiceUri string `pulumi:"serviceUri"`
}





type WebhookReceiverInput interface {
	pulumi.Input

	ToWebhookReceiverOutput() WebhookReceiverOutput
	ToWebhookReceiverOutputWithContext(context.Context) WebhookReceiverOutput
}

type WebhookReceiverArgs struct {
	Name       pulumi.StringInput `pulumi:"name"`
	ServiceUri pulumi.StringInput `pulumi:"serviceUri"`
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

func (o WebhookReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o WebhookReceiverOutput) ServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookReceiver) string { return v.ServiceUri }).(pulumi.StringOutput)
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
	Name       string `pulumi:"name"`
	ServiceUri string `pulumi:"serviceUri"`
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

func (o WebhookReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o WebhookReceiverResponseOutput) ServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v WebhookReceiverResponse) string { return v.ServiceUri }).(pulumi.StringOutput)
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

func init() {
	pulumi.RegisterOutputType(ActivityLogAlertActionGroupOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertActionGroupArrayOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertActionGroupResponseOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertActionGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertActionListOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertActionListResponseOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertAllOfConditionOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertAllOfConditionResponseOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertLeafConditionOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertLeafConditionArrayOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertLeafConditionResponseOutput{})
	pulumi.RegisterOutputType(ActivityLogAlertLeafConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverArrayOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverResponseOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverArrayOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverResponseOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(EmailReceiverOutput{})
	pulumi.RegisterOutputType(EmailReceiverArrayOutput{})
	pulumi.RegisterOutputType(EmailReceiverResponseOutput{})
	pulumi.RegisterOutputType(EmailReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(ItsmReceiverOutput{})
	pulumi.RegisterOutputType(ItsmReceiverArrayOutput{})
	pulumi.RegisterOutputType(ItsmReceiverResponseOutput{})
	pulumi.RegisterOutputType(ItsmReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(SmsReceiverOutput{})
	pulumi.RegisterOutputType(SmsReceiverArrayOutput{})
	pulumi.RegisterOutputType(SmsReceiverResponseOutput{})
	pulumi.RegisterOutputType(SmsReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(WebhookReceiverOutput{})
	pulumi.RegisterOutputType(WebhookReceiverArrayOutput{})
	pulumi.RegisterOutputType(WebhookReceiverResponseOutput{})
	pulumi.RegisterOutputType(WebhookReceiverResponseArrayOutput{})
}
