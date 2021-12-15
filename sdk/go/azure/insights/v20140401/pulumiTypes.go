


package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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


func (val *ScaleRuleResponse) Defaults() *ScaleRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ScaleAction = *tmp.ScaleAction.Defaults()

	return &tmp
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

func init() {
	pulumi.RegisterOutputType(AutoscaleNotificationOutput{})
	pulumi.RegisterOutputType(AutoscaleNotificationArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleNotificationResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleNotificationResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleProfileOutput{})
	pulumi.RegisterOutputType(AutoscaleProfileArrayOutput{})
	pulumi.RegisterOutputType(AutoscaleProfileResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(EmailNotificationOutput{})
	pulumi.RegisterOutputType(EmailNotificationPtrOutput{})
	pulumi.RegisterOutputType(EmailNotificationResponseOutput{})
	pulumi.RegisterOutputType(EmailNotificationResponsePtrOutput{})
	pulumi.RegisterOutputType(LocationThresholdRuleConditionOutput{})
	pulumi.RegisterOutputType(LocationThresholdRuleConditionResponseOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionPtrOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionResponseOutput{})
	pulumi.RegisterOutputType(ManagementEventAggregationConditionResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagementEventRuleConditionOutput{})
	pulumi.RegisterOutputType(ManagementEventRuleConditionResponseOutput{})
	pulumi.RegisterOutputType(MetricTriggerOutput{})
	pulumi.RegisterOutputType(MetricTriggerResponseOutput{})
	pulumi.RegisterOutputType(RecurrenceOutput{})
	pulumi.RegisterOutputType(RecurrencePtrOutput{})
	pulumi.RegisterOutputType(RecurrenceResponseOutput{})
	pulumi.RegisterOutputType(RecurrenceResponsePtrOutput{})
	pulumi.RegisterOutputType(RecurrentScheduleOutput{})
	pulumi.RegisterOutputType(RecurrentSchedulePtrOutput{})
	pulumi.RegisterOutputType(RecurrentScheduleResponseOutput{})
	pulumi.RegisterOutputType(RecurrentScheduleResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ThresholdRuleConditionOutput{})
	pulumi.RegisterOutputType(ThresholdRuleConditionResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowOutput{})
	pulumi.RegisterOutputType(TimeWindowPtrOutput{})
	pulumi.RegisterOutputType(TimeWindowResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(WebhookNotificationOutput{})
	pulumi.RegisterOutputType(WebhookNotificationArrayOutput{})
	pulumi.RegisterOutputType(WebhookNotificationResponseOutput{})
	pulumi.RegisterOutputType(WebhookNotificationResponseArrayOutput{})
}
