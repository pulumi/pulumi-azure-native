


package v20201015preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemTopicEventSubscription(ctx *pulumi.Context, args *LookupSystemTopicEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSystemTopicEventSubscriptionResult, error) {
	var rv LookupSystemTopicEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:getSystemTopicEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSystemTopicEventSubscriptionArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SystemTopicName       string `pulumi:"systemTopicName"`
}


type LookupSystemTopicEventSubscriptionResult struct {
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


func (val *LookupSystemTopicEventSubscriptionResult) Defaults() *LookupSystemTopicEventSubscriptionResult {
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

func LookupSystemTopicEventSubscriptionOutput(ctx *pulumi.Context, args LookupSystemTopicEventSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupSystemTopicEventSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemTopicEventSubscriptionResult, error) {
			args := v.(LookupSystemTopicEventSubscriptionArgs)
			r, err := LookupSystemTopicEventSubscription(ctx, &args, opts...)
			var s LookupSystemTopicEventSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemTopicEventSubscriptionResultOutput)
}

type LookupSystemTopicEventSubscriptionOutputArgs struct {
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	SystemTopicName       pulumi.StringInput `pulumi:"systemTopicName"`
}

func (LookupSystemTopicEventSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTopicEventSubscriptionArgs)(nil)).Elem()
}


type LookupSystemTopicEventSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupSystemTopicEventSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTopicEventSubscriptionResult)(nil)).Elem()
}

func (o LookupSystemTopicEventSubscriptionResultOutput) ToLookupSystemTopicEventSubscriptionResultOutput() LookupSystemTopicEventSubscriptionResultOutput {
	return o
}

func (o LookupSystemTopicEventSubscriptionResultOutput) ToLookupSystemTopicEventSubscriptionResultOutputWithContext(ctx context.Context) LookupSystemTopicEventSubscriptionResultOutput {
	return o
}

func (o LookupSystemTopicEventSubscriptionResultOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) *StorageBlobDeadLetterDestinationResponse {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) *DeadLetterWithResourceIdentityResponse {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) *DeliveryWithResourceIdentityResponse {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) interface{} { return v.Destination }).(pulumi.AnyOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) *string { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) *string { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) *EventSubscriptionFilterResponse { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) *RetryPolicyResponse { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) string { return v.Topic }).(pulumi.StringOutput)
}

func (o LookupSystemTopicEventSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicEventSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemTopicEventSubscriptionResultOutput{})
}
