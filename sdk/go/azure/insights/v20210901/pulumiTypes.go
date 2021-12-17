


package v20210901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

type EventHubReceiver struct {
	EventHubName         string  `pulumi:"eventHubName"`
	EventHubNameSpace    string  `pulumi:"eventHubNameSpace"`
	Name                 string  `pulumi:"name"`
	SubscriptionId       string  `pulumi:"subscriptionId"`
	TenantId             *string `pulumi:"tenantId"`
	UseCommonAlertSchema *bool   `pulumi:"useCommonAlertSchema"`
}


func (val *EventHubReceiver) Defaults() *EventHubReceiver {
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





type EventHubReceiverInput interface {
	pulumi.Input

	ToEventHubReceiverOutput() EventHubReceiverOutput
	ToEventHubReceiverOutputWithContext(context.Context) EventHubReceiverOutput
}

type EventHubReceiverArgs struct {
	EventHubName         pulumi.StringInput    `pulumi:"eventHubName"`
	EventHubNameSpace    pulumi.StringInput    `pulumi:"eventHubNameSpace"`
	Name                 pulumi.StringInput    `pulumi:"name"`
	SubscriptionId       pulumi.StringInput    `pulumi:"subscriptionId"`
	TenantId             pulumi.StringPtrInput `pulumi:"tenantId"`
	UseCommonAlertSchema pulumi.BoolPtrInput   `pulumi:"useCommonAlertSchema"`
}

func (EventHubReceiverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubReceiver)(nil)).Elem()
}

func (i EventHubReceiverArgs) ToEventHubReceiverOutput() EventHubReceiverOutput {
	return i.ToEventHubReceiverOutputWithContext(context.Background())
}

func (i EventHubReceiverArgs) ToEventHubReceiverOutputWithContext(ctx context.Context) EventHubReceiverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubReceiverOutput)
}





type EventHubReceiverArrayInput interface {
	pulumi.Input

	ToEventHubReceiverArrayOutput() EventHubReceiverArrayOutput
	ToEventHubReceiverArrayOutputWithContext(context.Context) EventHubReceiverArrayOutput
}

type EventHubReceiverArray []EventHubReceiverInput

func (EventHubReceiverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHubReceiver)(nil)).Elem()
}

func (i EventHubReceiverArray) ToEventHubReceiverArrayOutput() EventHubReceiverArrayOutput {
	return i.ToEventHubReceiverArrayOutputWithContext(context.Background())
}

func (i EventHubReceiverArray) ToEventHubReceiverArrayOutputWithContext(ctx context.Context) EventHubReceiverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubReceiverArrayOutput)
}

type EventHubReceiverOutput struct{ *pulumi.OutputState }

func (EventHubReceiverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubReceiver)(nil)).Elem()
}

func (o EventHubReceiverOutput) ToEventHubReceiverOutput() EventHubReceiverOutput {
	return o
}

func (o EventHubReceiverOutput) ToEventHubReceiverOutputWithContext(ctx context.Context) EventHubReceiverOutput {
	return o
}

func (o EventHubReceiverOutput) EventHubName() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubReceiver) string { return v.EventHubName }).(pulumi.StringOutput)
}

func (o EventHubReceiverOutput) EventHubNameSpace() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubReceiver) string { return v.EventHubNameSpace }).(pulumi.StringOutput)
}

func (o EventHubReceiverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubReceiver) string { return v.Name }).(pulumi.StringOutput)
}

func (o EventHubReceiverOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubReceiver) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o EventHubReceiverOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubReceiver) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o EventHubReceiverOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventHubReceiver) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type EventHubReceiverArrayOutput struct{ *pulumi.OutputState }

func (EventHubReceiverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHubReceiver)(nil)).Elem()
}

func (o EventHubReceiverArrayOutput) ToEventHubReceiverArrayOutput() EventHubReceiverArrayOutput {
	return o
}

func (o EventHubReceiverArrayOutput) ToEventHubReceiverArrayOutputWithContext(ctx context.Context) EventHubReceiverArrayOutput {
	return o
}

func (o EventHubReceiverArrayOutput) Index(i pulumi.IntInput) EventHubReceiverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventHubReceiver {
		return vs[0].([]EventHubReceiver)[vs[1].(int)]
	}).(EventHubReceiverOutput)
}

type EventHubReceiverResponse struct {
	EventHubName         string  `pulumi:"eventHubName"`
	EventHubNameSpace    string  `pulumi:"eventHubNameSpace"`
	Name                 string  `pulumi:"name"`
	SubscriptionId       string  `pulumi:"subscriptionId"`
	TenantId             *string `pulumi:"tenantId"`
	UseCommonAlertSchema *bool   `pulumi:"useCommonAlertSchema"`
}


func (val *EventHubReceiverResponse) Defaults() *EventHubReceiverResponse {
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

type EventHubReceiverResponseOutput struct{ *pulumi.OutputState }

func (EventHubReceiverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubReceiverResponse)(nil)).Elem()
}

func (o EventHubReceiverResponseOutput) ToEventHubReceiverResponseOutput() EventHubReceiverResponseOutput {
	return o
}

func (o EventHubReceiverResponseOutput) ToEventHubReceiverResponseOutputWithContext(ctx context.Context) EventHubReceiverResponseOutput {
	return o
}

func (o EventHubReceiverResponseOutput) EventHubName() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubReceiverResponse) string { return v.EventHubName }).(pulumi.StringOutput)
}

func (o EventHubReceiverResponseOutput) EventHubNameSpace() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubReceiverResponse) string { return v.EventHubNameSpace }).(pulumi.StringOutput)
}

func (o EventHubReceiverResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubReceiverResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EventHubReceiverResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubReceiverResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o EventHubReceiverResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubReceiverResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o EventHubReceiverResponseOutput) UseCommonAlertSchema() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventHubReceiverResponse) *bool { return v.UseCommonAlertSchema }).(pulumi.BoolPtrOutput)
}

type EventHubReceiverResponseArrayOutput struct{ *pulumi.OutputState }

func (EventHubReceiverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHubReceiverResponse)(nil)).Elem()
}

func (o EventHubReceiverResponseArrayOutput) ToEventHubReceiverResponseArrayOutput() EventHubReceiverResponseArrayOutput {
	return o
}

func (o EventHubReceiverResponseArrayOutput) ToEventHubReceiverResponseArrayOutputWithContext(ctx context.Context) EventHubReceiverResponseArrayOutput {
	return o
}

func (o EventHubReceiverResponseArrayOutput) Index(i pulumi.IntInput) EventHubReceiverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventHubReceiverResponse {
		return vs[0].([]EventHubReceiverResponse)[vs[1].(int)]
	}).(EventHubReceiverResponseOutput)
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

func init() {
	pulumi.RegisterOutputType(ArmRoleReceiverOutput{})
	pulumi.RegisterOutputType(ArmRoleReceiverArrayOutput{})
	pulumi.RegisterOutputType(ArmRoleReceiverResponseOutput{})
	pulumi.RegisterOutputType(ArmRoleReceiverResponseArrayOutput{})
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
	pulumi.RegisterOutputType(EmailReceiverOutput{})
	pulumi.RegisterOutputType(EmailReceiverArrayOutput{})
	pulumi.RegisterOutputType(EmailReceiverResponseOutput{})
	pulumi.RegisterOutputType(EmailReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(EventHubReceiverOutput{})
	pulumi.RegisterOutputType(EventHubReceiverArrayOutput{})
	pulumi.RegisterOutputType(EventHubReceiverResponseOutput{})
	pulumi.RegisterOutputType(EventHubReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(ItsmReceiverOutput{})
	pulumi.RegisterOutputType(ItsmReceiverArrayOutput{})
	pulumi.RegisterOutputType(ItsmReceiverResponseOutput{})
	pulumi.RegisterOutputType(ItsmReceiverResponseArrayOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverArrayOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverResponseOutput{})
	pulumi.RegisterOutputType(LogicAppReceiverResponseArrayOutput{})
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
}
