


package v20221001

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
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowOutput{})
	pulumi.RegisterOutputType(TimeWindowPtrOutput{})
	pulumi.RegisterOutputType(TimeWindowResponseOutput{})
	pulumi.RegisterOutputType(TimeWindowResponsePtrOutput{})
	pulumi.RegisterOutputType(WebhookNotificationOutput{})
	pulumi.RegisterOutputType(WebhookNotificationArrayOutput{})
	pulumi.RegisterOutputType(WebhookNotificationResponseOutput{})
	pulumi.RegisterOutputType(WebhookNotificationResponseArrayOutput{})
}
