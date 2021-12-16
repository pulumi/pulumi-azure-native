


package v20210501preview

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

type LogSettings struct {
	Category        *string          `pulumi:"category"`
	CategoryGroup   *string          `pulumi:"categoryGroup"`
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
	CategoryGroup   pulumi.StringPtrInput   `pulumi:"categoryGroup"`
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

func (o LogSettingsOutput) CategoryGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogSettings) *string { return v.CategoryGroup }).(pulumi.StringPtrOutput)
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
	CategoryGroup   *string                  `pulumi:"categoryGroup"`
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
	CategoryGroup   pulumi.StringPtrInput           `pulumi:"categoryGroup"`
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

func (o LogSettingsResponseOutput) CategoryGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogSettingsResponse) *string { return v.CategoryGroup }).(pulumi.StringPtrOutput)
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

type ManagementGroupLogSettings struct {
	Category      *string `pulumi:"category"`
	CategoryGroup *string `pulumi:"categoryGroup"`
	Enabled       bool    `pulumi:"enabled"`
}





type ManagementGroupLogSettingsInput interface {
	pulumi.Input

	ToManagementGroupLogSettingsOutput() ManagementGroupLogSettingsOutput
	ToManagementGroupLogSettingsOutputWithContext(context.Context) ManagementGroupLogSettingsOutput
}

type ManagementGroupLogSettingsArgs struct {
	Category      pulumi.StringPtrInput `pulumi:"category"`
	CategoryGroup pulumi.StringPtrInput `pulumi:"categoryGroup"`
	Enabled       pulumi.BoolInput      `pulumi:"enabled"`
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

func (o ManagementGroupLogSettingsOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupLogSettings) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupLogSettingsOutput) CategoryGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupLogSettings) *string { return v.CategoryGroup }).(pulumi.StringPtrOutput)
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
	Category      *string `pulumi:"category"`
	CategoryGroup *string `pulumi:"categoryGroup"`
	Enabled       bool    `pulumi:"enabled"`
}





type ManagementGroupLogSettingsResponseInput interface {
	pulumi.Input

	ToManagementGroupLogSettingsResponseOutput() ManagementGroupLogSettingsResponseOutput
	ToManagementGroupLogSettingsResponseOutputWithContext(context.Context) ManagementGroupLogSettingsResponseOutput
}

type ManagementGroupLogSettingsResponseArgs struct {
	Category      pulumi.StringPtrInput `pulumi:"category"`
	CategoryGroup pulumi.StringPtrInput `pulumi:"categoryGroup"`
	Enabled       pulumi.BoolInput      `pulumi:"enabled"`
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

func (o ManagementGroupLogSettingsResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupLogSettingsResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupLogSettingsResponseOutput) CategoryGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupLogSettingsResponse) *string { return v.CategoryGroup }).(pulumi.StringPtrOutput)
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

type PredictiveAutoscalePolicy struct {
	ScaleLookAheadTime *string                            `pulumi:"scaleLookAheadTime"`
	ScaleMode          PredictiveAutoscalePolicyScaleMode `pulumi:"scaleMode"`
}





type PredictiveAutoscalePolicyInput interface {
	pulumi.Input

	ToPredictiveAutoscalePolicyOutput() PredictiveAutoscalePolicyOutput
	ToPredictiveAutoscalePolicyOutputWithContext(context.Context) PredictiveAutoscalePolicyOutput
}

type PredictiveAutoscalePolicyArgs struct {
	ScaleLookAheadTime pulumi.StringPtrInput                   `pulumi:"scaleLookAheadTime"`
	ScaleMode          PredictiveAutoscalePolicyScaleModeInput `pulumi:"scaleMode"`
}

func (PredictiveAutoscalePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PredictiveAutoscalePolicy)(nil)).Elem()
}

func (i PredictiveAutoscalePolicyArgs) ToPredictiveAutoscalePolicyOutput() PredictiveAutoscalePolicyOutput {
	return i.ToPredictiveAutoscalePolicyOutputWithContext(context.Background())
}

func (i PredictiveAutoscalePolicyArgs) ToPredictiveAutoscalePolicyOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PredictiveAutoscalePolicyOutput)
}

func (i PredictiveAutoscalePolicyArgs) ToPredictiveAutoscalePolicyPtrOutput() PredictiveAutoscalePolicyPtrOutput {
	return i.ToPredictiveAutoscalePolicyPtrOutputWithContext(context.Background())
}

func (i PredictiveAutoscalePolicyArgs) ToPredictiveAutoscalePolicyPtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PredictiveAutoscalePolicyOutput).ToPredictiveAutoscalePolicyPtrOutputWithContext(ctx)
}









type PredictiveAutoscalePolicyPtrInput interface {
	pulumi.Input

	ToPredictiveAutoscalePolicyPtrOutput() PredictiveAutoscalePolicyPtrOutput
	ToPredictiveAutoscalePolicyPtrOutputWithContext(context.Context) PredictiveAutoscalePolicyPtrOutput
}

type predictiveAutoscalePolicyPtrType PredictiveAutoscalePolicyArgs

func PredictiveAutoscalePolicyPtr(v *PredictiveAutoscalePolicyArgs) PredictiveAutoscalePolicyPtrInput {
	return (*predictiveAutoscalePolicyPtrType)(v)
}

func (*predictiveAutoscalePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PredictiveAutoscalePolicy)(nil)).Elem()
}

func (i *predictiveAutoscalePolicyPtrType) ToPredictiveAutoscalePolicyPtrOutput() PredictiveAutoscalePolicyPtrOutput {
	return i.ToPredictiveAutoscalePolicyPtrOutputWithContext(context.Background())
}

func (i *predictiveAutoscalePolicyPtrType) ToPredictiveAutoscalePolicyPtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PredictiveAutoscalePolicyPtrOutput)
}

type PredictiveAutoscalePolicyOutput struct{ *pulumi.OutputState }

func (PredictiveAutoscalePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PredictiveAutoscalePolicy)(nil)).Elem()
}

func (o PredictiveAutoscalePolicyOutput) ToPredictiveAutoscalePolicyOutput() PredictiveAutoscalePolicyOutput {
	return o
}

func (o PredictiveAutoscalePolicyOutput) ToPredictiveAutoscalePolicyOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyOutput {
	return o
}

func (o PredictiveAutoscalePolicyOutput) ToPredictiveAutoscalePolicyPtrOutput() PredictiveAutoscalePolicyPtrOutput {
	return o.ToPredictiveAutoscalePolicyPtrOutputWithContext(context.Background())
}

func (o PredictiveAutoscalePolicyOutput) ToPredictiveAutoscalePolicyPtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PredictiveAutoscalePolicy) *PredictiveAutoscalePolicy {
		return &v
	}).(PredictiveAutoscalePolicyPtrOutput)
}

func (o PredictiveAutoscalePolicyOutput) ScaleLookAheadTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PredictiveAutoscalePolicy) *string { return v.ScaleLookAheadTime }).(pulumi.StringPtrOutput)
}

func (o PredictiveAutoscalePolicyOutput) ScaleMode() PredictiveAutoscalePolicyScaleModeOutput {
	return o.ApplyT(func(v PredictiveAutoscalePolicy) PredictiveAutoscalePolicyScaleMode { return v.ScaleMode }).(PredictiveAutoscalePolicyScaleModeOutput)
}

type PredictiveAutoscalePolicyPtrOutput struct{ *pulumi.OutputState }

func (PredictiveAutoscalePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PredictiveAutoscalePolicy)(nil)).Elem()
}

func (o PredictiveAutoscalePolicyPtrOutput) ToPredictiveAutoscalePolicyPtrOutput() PredictiveAutoscalePolicyPtrOutput {
	return o
}

func (o PredictiveAutoscalePolicyPtrOutput) ToPredictiveAutoscalePolicyPtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyPtrOutput {
	return o
}

func (o PredictiveAutoscalePolicyPtrOutput) Elem() PredictiveAutoscalePolicyOutput {
	return o.ApplyT(func(v *PredictiveAutoscalePolicy) PredictiveAutoscalePolicy {
		if v != nil {
			return *v
		}
		var ret PredictiveAutoscalePolicy
		return ret
	}).(PredictiveAutoscalePolicyOutput)
}

func (o PredictiveAutoscalePolicyPtrOutput) ScaleLookAheadTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PredictiveAutoscalePolicy) *string {
		if v == nil {
			return nil
		}
		return v.ScaleLookAheadTime
	}).(pulumi.StringPtrOutput)
}

func (o PredictiveAutoscalePolicyPtrOutput) ScaleMode() PredictiveAutoscalePolicyScaleModePtrOutput {
	return o.ApplyT(func(v *PredictiveAutoscalePolicy) *PredictiveAutoscalePolicyScaleMode {
		if v == nil {
			return nil
		}
		return &v.ScaleMode
	}).(PredictiveAutoscalePolicyScaleModePtrOutput)
}

type PredictiveAutoscalePolicyResponse struct {
	ScaleLookAheadTime *string `pulumi:"scaleLookAheadTime"`
	ScaleMode          string  `pulumi:"scaleMode"`
}





type PredictiveAutoscalePolicyResponseInput interface {
	pulumi.Input

	ToPredictiveAutoscalePolicyResponseOutput() PredictiveAutoscalePolicyResponseOutput
	ToPredictiveAutoscalePolicyResponseOutputWithContext(context.Context) PredictiveAutoscalePolicyResponseOutput
}

type PredictiveAutoscalePolicyResponseArgs struct {
	ScaleLookAheadTime pulumi.StringPtrInput `pulumi:"scaleLookAheadTime"`
	ScaleMode          pulumi.StringInput    `pulumi:"scaleMode"`
}

func (PredictiveAutoscalePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PredictiveAutoscalePolicyResponse)(nil)).Elem()
}

func (i PredictiveAutoscalePolicyResponseArgs) ToPredictiveAutoscalePolicyResponseOutput() PredictiveAutoscalePolicyResponseOutput {
	return i.ToPredictiveAutoscalePolicyResponseOutputWithContext(context.Background())
}

func (i PredictiveAutoscalePolicyResponseArgs) ToPredictiveAutoscalePolicyResponseOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PredictiveAutoscalePolicyResponseOutput)
}

func (i PredictiveAutoscalePolicyResponseArgs) ToPredictiveAutoscalePolicyResponsePtrOutput() PredictiveAutoscalePolicyResponsePtrOutput {
	return i.ToPredictiveAutoscalePolicyResponsePtrOutputWithContext(context.Background())
}

func (i PredictiveAutoscalePolicyResponseArgs) ToPredictiveAutoscalePolicyResponsePtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PredictiveAutoscalePolicyResponseOutput).ToPredictiveAutoscalePolicyResponsePtrOutputWithContext(ctx)
}









type PredictiveAutoscalePolicyResponsePtrInput interface {
	pulumi.Input

	ToPredictiveAutoscalePolicyResponsePtrOutput() PredictiveAutoscalePolicyResponsePtrOutput
	ToPredictiveAutoscalePolicyResponsePtrOutputWithContext(context.Context) PredictiveAutoscalePolicyResponsePtrOutput
}

type predictiveAutoscalePolicyResponsePtrType PredictiveAutoscalePolicyResponseArgs

func PredictiveAutoscalePolicyResponsePtr(v *PredictiveAutoscalePolicyResponseArgs) PredictiveAutoscalePolicyResponsePtrInput {
	return (*predictiveAutoscalePolicyResponsePtrType)(v)
}

func (*predictiveAutoscalePolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PredictiveAutoscalePolicyResponse)(nil)).Elem()
}

func (i *predictiveAutoscalePolicyResponsePtrType) ToPredictiveAutoscalePolicyResponsePtrOutput() PredictiveAutoscalePolicyResponsePtrOutput {
	return i.ToPredictiveAutoscalePolicyResponsePtrOutputWithContext(context.Background())
}

func (i *predictiveAutoscalePolicyResponsePtrType) ToPredictiveAutoscalePolicyResponsePtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PredictiveAutoscalePolicyResponsePtrOutput)
}

type PredictiveAutoscalePolicyResponseOutput struct{ *pulumi.OutputState }

func (PredictiveAutoscalePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PredictiveAutoscalePolicyResponse)(nil)).Elem()
}

func (o PredictiveAutoscalePolicyResponseOutput) ToPredictiveAutoscalePolicyResponseOutput() PredictiveAutoscalePolicyResponseOutput {
	return o
}

func (o PredictiveAutoscalePolicyResponseOutput) ToPredictiveAutoscalePolicyResponseOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyResponseOutput {
	return o
}

func (o PredictiveAutoscalePolicyResponseOutput) ToPredictiveAutoscalePolicyResponsePtrOutput() PredictiveAutoscalePolicyResponsePtrOutput {
	return o.ToPredictiveAutoscalePolicyResponsePtrOutputWithContext(context.Background())
}

func (o PredictiveAutoscalePolicyResponseOutput) ToPredictiveAutoscalePolicyResponsePtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PredictiveAutoscalePolicyResponse) *PredictiveAutoscalePolicyResponse {
		return &v
	}).(PredictiveAutoscalePolicyResponsePtrOutput)
}

func (o PredictiveAutoscalePolicyResponseOutput) ScaleLookAheadTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PredictiveAutoscalePolicyResponse) *string { return v.ScaleLookAheadTime }).(pulumi.StringPtrOutput)
}

func (o PredictiveAutoscalePolicyResponseOutput) ScaleMode() pulumi.StringOutput {
	return o.ApplyT(func(v PredictiveAutoscalePolicyResponse) string { return v.ScaleMode }).(pulumi.StringOutput)
}

type PredictiveAutoscalePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (PredictiveAutoscalePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PredictiveAutoscalePolicyResponse)(nil)).Elem()
}

func (o PredictiveAutoscalePolicyResponsePtrOutput) ToPredictiveAutoscalePolicyResponsePtrOutput() PredictiveAutoscalePolicyResponsePtrOutput {
	return o
}

func (o PredictiveAutoscalePolicyResponsePtrOutput) ToPredictiveAutoscalePolicyResponsePtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyResponsePtrOutput {
	return o
}

func (o PredictiveAutoscalePolicyResponsePtrOutput) Elem() PredictiveAutoscalePolicyResponseOutput {
	return o.ApplyT(func(v *PredictiveAutoscalePolicyResponse) PredictiveAutoscalePolicyResponse {
		if v != nil {
			return *v
		}
		var ret PredictiveAutoscalePolicyResponse
		return ret
	}).(PredictiveAutoscalePolicyResponseOutput)
}

func (o PredictiveAutoscalePolicyResponsePtrOutput) ScaleLookAheadTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PredictiveAutoscalePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScaleLookAheadTime
	}).(pulumi.StringPtrOutput)
}

func (o PredictiveAutoscalePolicyResponsePtrOutput) ScaleMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PredictiveAutoscalePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ScaleMode
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

type SubscriptionLogSettings struct {
	Category      *string `pulumi:"category"`
	CategoryGroup *string `pulumi:"categoryGroup"`
	Enabled       bool    `pulumi:"enabled"`
}





type SubscriptionLogSettingsInput interface {
	pulumi.Input

	ToSubscriptionLogSettingsOutput() SubscriptionLogSettingsOutput
	ToSubscriptionLogSettingsOutputWithContext(context.Context) SubscriptionLogSettingsOutput
}

type SubscriptionLogSettingsArgs struct {
	Category      pulumi.StringPtrInput `pulumi:"category"`
	CategoryGroup pulumi.StringPtrInput `pulumi:"categoryGroup"`
	Enabled       pulumi.BoolInput      `pulumi:"enabled"`
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

func (o SubscriptionLogSettingsOutput) CategoryGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionLogSettings) *string { return v.CategoryGroup }).(pulumi.StringPtrOutput)
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
	Category      *string `pulumi:"category"`
	CategoryGroup *string `pulumi:"categoryGroup"`
	Enabled       bool    `pulumi:"enabled"`
}





type SubscriptionLogSettingsResponseInput interface {
	pulumi.Input

	ToSubscriptionLogSettingsResponseOutput() SubscriptionLogSettingsResponseOutput
	ToSubscriptionLogSettingsResponseOutputWithContext(context.Context) SubscriptionLogSettingsResponseOutput
}

type SubscriptionLogSettingsResponseArgs struct {
	Category      pulumi.StringPtrInput `pulumi:"category"`
	CategoryGroup pulumi.StringPtrInput `pulumi:"categoryGroup"`
	Enabled       pulumi.BoolInput      `pulumi:"enabled"`
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

func (o SubscriptionLogSettingsResponseOutput) CategoryGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionLogSettingsResponse) *string { return v.CategoryGroup }).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(LogSettingsOutput{})
	pulumi.RegisterOutputType(LogSettingsArrayOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricSettingsOutput{})
	pulumi.RegisterOutputType(MetricSettingsArrayOutput{})
	pulumi.RegisterOutputType(MetricSettingsResponseOutput{})
	pulumi.RegisterOutputType(MetricSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricTriggerOutput{})
	pulumi.RegisterOutputType(MetricTriggerResponseOutput{})
	pulumi.RegisterOutputType(PredictiveAutoscalePolicyOutput{})
	pulumi.RegisterOutputType(PredictiveAutoscalePolicyPtrOutput{})
	pulumi.RegisterOutputType(PredictiveAutoscalePolicyResponseOutput{})
	pulumi.RegisterOutputType(PredictiveAutoscalePolicyResponsePtrOutput{})
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
	pulumi.RegisterOutputType(SubscriptionLogSettingsOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsResponseOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TimeWindowOutput{})
	pulumi.RegisterOutputType(TimeWindowPtrOutput{})
	pulumi.RegisterOutputType(TimeWindowResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(WebhookNotificationOutput{})
	pulumi.RegisterOutputType(WebhookNotificationArrayOutput{})
	pulumi.RegisterOutputType(WebhookNotificationResponseOutput{})
	pulumi.RegisterOutputType(WebhookNotificationResponseArrayOutput{})
}
