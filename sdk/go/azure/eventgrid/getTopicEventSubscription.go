


package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTopicEventSubscription(ctx *pulumi.Context, args *LookupTopicEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupTopicEventSubscriptionResult, error) {
	var rv LookupTopicEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid:getTopicEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupTopicEventSubscriptionArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TopicName             string `pulumi:"topicName"`
}


type LookupTopicEventSubscriptionResult struct {
	DeadLetterDestination          *StorageBlobDeadLetterDestinationResponse `pulumi:"deadLetterDestination"`
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentityResponse   `pulumi:"deadLetterWithResourceIdentity"`
	DeliveryWithResourceIdentity   *DeliveryWithResourceIdentityResponse     `pulumi:"deliveryWithResourceIdentity"`
	Destination                    interface{}                               `pulumi:"destination"`
	EventDeliverySchema            *string                                   `pulumi:"eventDeliverySchema"`
	ExpirationTimeUtc              *string                                   `pulumi:"expirationTimeUtc"`
	Filter                         *EventSubscriptionFilterResponse          `pulumi:"filter"`
	Id                             string                                    `pulumi:"id"`
	Labels                         []string                                  `pulumi:"labels"`
	Name                           string                                    `pulumi:"name"`
	ProvisioningState              string                                    `pulumi:"provisioningState"`
	RetryPolicy                    *RetryPolicyResponse                      `pulumi:"retryPolicy"`
	SystemData                     SystemDataResponse                        `pulumi:"systemData"`
	Topic                          string                                    `pulumi:"topic"`
	Type                           string                                    `pulumi:"type"`
}


func (val *LookupTopicEventSubscriptionResult) Defaults() *LookupTopicEventSubscriptionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EventDeliverySchema) {
		eventDeliverySchema_ := "EventGridSchema"
		tmp.EventDeliverySchema = &eventDeliverySchema_
	}
	tmp.Filter = tmp.Filter.Defaults()

	tmp.RetryPolicy = tmp.RetryPolicy.Defaults()

	return &tmp
}

func LookupTopicEventSubscriptionOutput(ctx *pulumi.Context, args LookupTopicEventSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupTopicEventSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicEventSubscriptionResult, error) {
			args := v.(LookupTopicEventSubscriptionArgs)
			r, err := LookupTopicEventSubscription(ctx, &args, opts...)
			var s LookupTopicEventSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicEventSubscriptionResultOutput)
}

type LookupTopicEventSubscriptionOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName             pulumi.StringInput `pulumi:"topicName"`
}

func (LookupTopicEventSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicEventSubscriptionArgs)(nil)).Elem()
}


type LookupTopicEventSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupTopicEventSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicEventSubscriptionResult)(nil)).Elem()
}

func (o LookupTopicEventSubscriptionResultOutput) ToLookupTopicEventSubscriptionResultOutput() LookupTopicEventSubscriptionResultOutput {
	return o
}

func (o LookupTopicEventSubscriptionResultOutput) ToLookupTopicEventSubscriptionResultOutputWithContext(ctx context.Context) LookupTopicEventSubscriptionResultOutput {
	return o
}

func (o LookupTopicEventSubscriptionResultOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) *StorageBlobDeadLetterDestinationResponse {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) *DeadLetterWithResourceIdentityResponse {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) *DeliveryWithResourceIdentityResponse {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) interface{} { return v.Destination }).(pulumi.AnyOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) *string { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) *string { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) *EventSubscriptionFilterResponse { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) *RetryPolicyResponse { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) string { return v.Topic }).(pulumi.StringOutput)
}

func (o LookupTopicEventSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicEventSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicEventSubscriptionResultOutput{})
}
