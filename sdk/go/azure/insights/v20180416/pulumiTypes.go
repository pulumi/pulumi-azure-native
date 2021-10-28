


package v20180416

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

func init() {
	pulumi.RegisterOutputType(AlertingActionOutput{})
	pulumi.RegisterOutputType(AlertingActionResponseOutput{})
	pulumi.RegisterOutputType(AzNsActionGroupOutput{})
	pulumi.RegisterOutputType(AzNsActionGroupPtrOutput{})
	pulumi.RegisterOutputType(AzNsActionGroupResponseOutput{})
	pulumi.RegisterOutputType(AzNsActionGroupResponsePtrOutput{})
	pulumi.RegisterOutputType(CriteriaOutput{})
	pulumi.RegisterOutputType(CriteriaArrayOutput{})
	pulumi.RegisterOutputType(CriteriaResponseOutput{})
	pulumi.RegisterOutputType(CriteriaResponseArrayOutput{})
	pulumi.RegisterOutputType(DimensionOutput{})
	pulumi.RegisterOutputType(DimensionArrayOutput{})
	pulumi.RegisterOutputType(DimensionResponseOutput{})
	pulumi.RegisterOutputType(DimensionResponseArrayOutput{})
	pulumi.RegisterOutputType(LogMetricTriggerOutput{})
	pulumi.RegisterOutputType(LogMetricTriggerPtrOutput{})
	pulumi.RegisterOutputType(LogMetricTriggerResponseOutput{})
	pulumi.RegisterOutputType(LogMetricTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(LogToMetricActionOutput{})
	pulumi.RegisterOutputType(LogToMetricActionResponseOutput{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(SchedulePtrOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceOutput{})
	pulumi.RegisterOutputType(SourcePtrOutput{})
	pulumi.RegisterOutputType(SourceResponseOutput{})
	pulumi.RegisterOutputType(SourceResponsePtrOutput{})
	pulumi.RegisterOutputType(TriggerConditionOutput{})
	pulumi.RegisterOutputType(TriggerConditionResponseOutput{})
}
