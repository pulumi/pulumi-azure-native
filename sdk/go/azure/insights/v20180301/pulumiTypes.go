


package v20180301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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





type AutomationRunbookReceiverResponseInput interface {
	pulumi.Input

	ToAutomationRunbookReceiverResponseOutput() AutomationRunbookReceiverResponseOutput
	ToAutomationRunbookReceiverResponseOutputWithContext(context.Context) AutomationRunbookReceiverResponseOutput
}

type AutomationRunbookReceiverResponseArgs struct {
	AutomationAccountId pulumi.StringInput    `pulumi:"automationAccountId"`
	IsGlobalRunbook     pulumi.BoolInput      `pulumi:"isGlobalRunbook"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	RunbookName         pulumi.StringInput    `pulumi:"runbookName"`
	ServiceUri          pulumi.StringPtrInput `pulumi:"serviceUri"`
	WebhookResourceId   pulumi.StringInput    `pulumi:"webhookResourceId"`
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
}





type AzureFunctionReceiverInput interface {
	pulumi.Input

	ToAzureFunctionReceiverOutput() AzureFunctionReceiverOutput
	ToAzureFunctionReceiverOutputWithContext(context.Context) AzureFunctionReceiverOutput
}

type AzureFunctionReceiverArgs struct {
	FunctionAppResourceId pulumi.StringInput `pulumi:"functionAppResourceId"`
	FunctionName          pulumi.StringInput `pulumi:"functionName"`
	HttpTriggerUrl        pulumi.StringInput `pulumi:"httpTriggerUrl"`
	Name                  pulumi.StringInput `pulumi:"name"`
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
}





type AzureFunctionReceiverResponseInput interface {
	pulumi.Input

	ToAzureFunctionReceiverResponseOutput() AzureFunctionReceiverResponseOutput
	ToAzureFunctionReceiverResponseOutputWithContext(context.Context) AzureFunctionReceiverResponseOutput
}

type AzureFunctionReceiverResponseArgs struct {
	FunctionAppResourceId pulumi.StringInput `pulumi:"functionAppResourceId"`
	FunctionName          pulumi.StringInput `pulumi:"functionName"`
	HttpTriggerUrl        pulumi.StringInput `pulumi:"httpTriggerUrl"`
	Name                  pulumi.StringInput `pulumi:"name"`
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





type EmailReceiverResponseInput interface {
	pulumi.Input

	ToEmailReceiverResponseOutput() EmailReceiverResponseOutput
	ToEmailReceiverResponseOutputWithContext(context.Context) EmailReceiverResponseOutput
}

type EmailReceiverResponseArgs struct {
	EmailAddress pulumi.StringInput `pulumi:"emailAddress"`
	Name         pulumi.StringInput `pulumi:"name"`
	Status       pulumi.StringInput `pulumi:"status"`
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

type LogicAppReceiver struct {
	CallbackUrl string `pulumi:"callbackUrl"`
	Name        string `pulumi:"name"`
	ResourceId  string `pulumi:"resourceId"`
}





type LogicAppReceiverInput interface {
	pulumi.Input

	ToLogicAppReceiverOutput() LogicAppReceiverOutput
	ToLogicAppReceiverOutputWithContext(context.Context) LogicAppReceiverOutput
}

type LogicAppReceiverArgs struct {
	CallbackUrl pulumi.StringInput `pulumi:"callbackUrl"`
	Name        pulumi.StringInput `pulumi:"name"`
	ResourceId  pulumi.StringInput `pulumi:"resourceId"`
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
	CallbackUrl string `pulumi:"callbackUrl"`
	Name        string `pulumi:"name"`
	ResourceId  string `pulumi:"resourceId"`
}





type LogicAppReceiverResponseInput interface {
	pulumi.Input

	ToLogicAppReceiverResponseOutput() LogicAppReceiverResponseOutput
	ToLogicAppReceiverResponseOutputWithContext(context.Context) LogicAppReceiverResponseOutput
}

type LogicAppReceiverResponseArgs struct {
	CallbackUrl pulumi.StringInput `pulumi:"callbackUrl"`
	Name        pulumi.StringInput `pulumi:"name"`
	ResourceId  pulumi.StringInput `pulumi:"resourceId"`
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





type WebhookReceiverResponseInput interface {
	pulumi.Input

	ToWebhookReceiverResponseOutput() WebhookReceiverResponseOutput
	ToWebhookReceiverResponseOutputWithContext(context.Context) WebhookReceiverResponseOutput
}

type WebhookReceiverResponseArgs struct {
	Name       pulumi.StringInput `pulumi:"name"`
	ServiceUri pulumi.StringInput `pulumi:"serviceUri"`
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRunbookReceiverInput)(nil)).Elem(), AutomationRunbookReceiverArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRunbookReceiverArrayInput)(nil)).Elem(), AutomationRunbookReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRunbookReceiverResponseInput)(nil)).Elem(), AutomationRunbookReceiverResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRunbookReceiverResponseArrayInput)(nil)).Elem(), AutomationRunbookReceiverResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureAppPushReceiverInput)(nil)).Elem(), AzureAppPushReceiverArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureAppPushReceiverArrayInput)(nil)).Elem(), AzureAppPushReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureAppPushReceiverResponseInput)(nil)).Elem(), AzureAppPushReceiverResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureAppPushReceiverResponseArrayInput)(nil)).Elem(), AzureAppPushReceiverResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFunctionReceiverInput)(nil)).Elem(), AzureFunctionReceiverArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFunctionReceiverArrayInput)(nil)).Elem(), AzureFunctionReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFunctionReceiverResponseInput)(nil)).Elem(), AzureFunctionReceiverResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFunctionReceiverResponseArrayInput)(nil)).Elem(), AzureFunctionReceiverResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicMetricCriteriaInput)(nil)).Elem(), DynamicMetricCriteriaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicMetricCriteriaResponseInput)(nil)).Elem(), DynamicMetricCriteriaResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicThresholdFailingPeriodsInput)(nil)).Elem(), DynamicThresholdFailingPeriodsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicThresholdFailingPeriodsResponseInput)(nil)).Elem(), DynamicThresholdFailingPeriodsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailReceiverInput)(nil)).Elem(), EmailReceiverArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailReceiverArrayInput)(nil)).Elem(), EmailReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailReceiverResponseInput)(nil)).Elem(), EmailReceiverResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailReceiverResponseArrayInput)(nil)).Elem(), EmailReceiverResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItsmReceiverInput)(nil)).Elem(), ItsmReceiverArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItsmReceiverArrayInput)(nil)).Elem(), ItsmReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItsmReceiverResponseInput)(nil)).Elem(), ItsmReceiverResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ItsmReceiverResponseArrayInput)(nil)).Elem(), ItsmReceiverResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogicAppReceiverInput)(nil)).Elem(), LogicAppReceiverArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogicAppReceiverArrayInput)(nil)).Elem(), LogicAppReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogicAppReceiverResponseInput)(nil)).Elem(), LogicAppReceiverResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogicAppReceiverResponseArrayInput)(nil)).Elem(), LogicAppReceiverResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricAlertActionInput)(nil)).Elem(), MetricAlertActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricAlertActionArrayInput)(nil)).Elem(), MetricAlertActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricAlertActionResponseInput)(nil)).Elem(), MetricAlertActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricAlertActionResponseArrayInput)(nil)).Elem(), MetricAlertActionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricAlertMultipleResourceMultipleMetricCriteriaInput)(nil)).Elem(), MetricAlertMultipleResourceMultipleMetricCriteriaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricAlertMultipleResourceMultipleMetricCriteriaResponseInput)(nil)).Elem(), MetricAlertMultipleResourceMultipleMetricCriteriaResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricAlertSingleResourceMultipleMetricCriteriaInput)(nil)).Elem(), MetricAlertSingleResourceMultipleMetricCriteriaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricAlertSingleResourceMultipleMetricCriteriaResponseInput)(nil)).Elem(), MetricAlertSingleResourceMultipleMetricCriteriaResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricCriteriaInput)(nil)).Elem(), MetricCriteriaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricCriteriaArrayInput)(nil)).Elem(), MetricCriteriaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricCriteriaResponseInput)(nil)).Elem(), MetricCriteriaResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricCriteriaResponseArrayInput)(nil)).Elem(), MetricCriteriaResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricDimensionInput)(nil)).Elem(), MetricDimensionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricDimensionArrayInput)(nil)).Elem(), MetricDimensionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricDimensionResponseInput)(nil)).Elem(), MetricDimensionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricDimensionResponseArrayInput)(nil)).Elem(), MetricDimensionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsReceiverInput)(nil)).Elem(), SmsReceiverArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsReceiverArrayInput)(nil)).Elem(), SmsReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsReceiverResponseInput)(nil)).Elem(), SmsReceiverResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsReceiverResponseArrayInput)(nil)).Elem(), SmsReceiverResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceReceiverInput)(nil)).Elem(), VoiceReceiverArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceReceiverArrayInput)(nil)).Elem(), VoiceReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceReceiverResponseInput)(nil)).Elem(), VoiceReceiverResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VoiceReceiverResponseArrayInput)(nil)).Elem(), VoiceReceiverResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookReceiverInput)(nil)).Elem(), WebhookReceiverArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookReceiverArrayInput)(nil)).Elem(), WebhookReceiverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookReceiverResponseInput)(nil)).Elem(), WebhookReceiverResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookReceiverResponseArrayInput)(nil)).Elem(), WebhookReceiverResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebtestLocationAvailabilityCriteriaInput)(nil)).Elem(), WebtestLocationAvailabilityCriteriaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebtestLocationAvailabilityCriteriaResponseInput)(nil)).Elem(), WebtestLocationAvailabilityCriteriaResponseArgs{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverArrayOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverResponseOutput{})
	pulumi.RegisterOutputType(AutomationRunbookReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverArrayOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverResponseOutput{})
	pulumi.RegisterOutputType(AzureAppPushReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverArrayOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverResponseOutput{})
	pulumi.RegisterOutputType(AzureFunctionReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(DynamicMetricCriteriaOutput{})
	pulumi.RegisterOutputType(DynamicMetricCriteriaResponseOutput{})
	pulumi.RegisterOutputType(DynamicThresholdFailingPeriodsOutput{})
	pulumi.RegisterOutputType(DynamicThresholdFailingPeriodsResponseOutput{})
	pulumi.RegisterOutputType(EmailReceiverOutput{})
	pulumi.RegisterOutputType(EmailReceiverArrayOutput{})
	pulumi.RegisterOutputType(EmailReceiverResponseOutput{})
	pulumi.RegisterOutputType(EmailReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(ItsmReceiverOutput{})
	pulumi.RegisterOutputType(ItsmReceiverArrayOutput{})
	pulumi.RegisterOutputType(ItsmReceiverResponseOutput{})
	pulumi.RegisterOutputType(ItsmReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverArrayOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverResponseOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverResponseArrayOutput{})
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
	pulumi.RegisterOutputType(SmsReceiverOutput{})
	pulumi.RegisterOutputType(SmsReceiverArrayOutput{})
	pulumi.RegisterOutputType(SmsReceiverResponseOutput{})
	pulumi.RegisterOutputType(SmsReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(VoiceReceiverOutput{})
	pulumi.RegisterOutputType(VoiceReceiverArrayOutput{})
	pulumi.RegisterOutputType(VoiceReceiverResponseOutput{})
	pulumi.RegisterOutputType(VoiceReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(WebhookReceiverOutput{})
	pulumi.RegisterOutputType(WebhookReceiverArrayOutput{})
	pulumi.RegisterOutputType(WebhookReceiverResponseOutput{})
	pulumi.RegisterOutputType(WebhookReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(WebtestLocationAvailabilityCriteriaOutput{})
	pulumi.RegisterOutputType(WebtestLocationAvailabilityCriteriaResponseOutput{})
}
