


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomainTopicEventSubscription(ctx *pulumi.Context, args *LookupDomainTopicEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupDomainTopicEventSubscriptionResult, error) {
	var rv LookupDomainTopicEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20211015preview:getDomainTopicEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainTopicEventSubscriptionArgs struct {
	DomainName            string `pulumi:"domainName"`
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	TopicName             string `pulumi:"topicName"`
}


type LookupDomainTopicEventSubscriptionResult struct {
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


func (val *LookupDomainTopicEventSubscriptionResult) Defaults() *LookupDomainTopicEventSubscriptionResult {
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

func LookupDomainTopicEventSubscriptionOutput(ctx *pulumi.Context, args LookupDomainTopicEventSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupDomainTopicEventSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainTopicEventSubscriptionResult, error) {
			args := v.(LookupDomainTopicEventSubscriptionArgs)
			r, err := LookupDomainTopicEventSubscription(ctx, &args, opts...)
			var s LookupDomainTopicEventSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainTopicEventSubscriptionResultOutput)
}

type LookupDomainTopicEventSubscriptionOutputArgs struct {
	DomainName            pulumi.StringInput `pulumi:"domainName"`
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName             pulumi.StringInput `pulumi:"topicName"`
}

func (LookupDomainTopicEventSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainTopicEventSubscriptionArgs)(nil)).Elem()
}


type LookupDomainTopicEventSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupDomainTopicEventSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainTopicEventSubscriptionResult)(nil)).Elem()
}

func (o LookupDomainTopicEventSubscriptionResultOutput) ToLookupDomainTopicEventSubscriptionResultOutput() LookupDomainTopicEventSubscriptionResultOutput {
	return o
}

func (o LookupDomainTopicEventSubscriptionResultOutput) ToLookupDomainTopicEventSubscriptionResultOutputWithContext(ctx context.Context) LookupDomainTopicEventSubscriptionResultOutput {
	return o
}

func (o LookupDomainTopicEventSubscriptionResultOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) *StorageBlobDeadLetterDestinationResponse {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) *DeadLetterWithResourceIdentityResponse {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) *DeliveryWithResourceIdentityResponse {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) interface{} { return v.Destination }).(pulumi.AnyOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) *string { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) *string { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) *EventSubscriptionFilterResponse { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) *RetryPolicyResponse { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) string { return v.Topic }).(pulumi.StringOutput)
}

func (o LookupDomainTopicEventSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicEventSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainTopicEventSubscriptionResultOutput{})
}
