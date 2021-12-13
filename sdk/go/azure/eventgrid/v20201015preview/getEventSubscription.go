


package v20201015preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventSubscription(ctx *pulumi.Context, args *LookupEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupEventSubscriptionResult, error) {
	var rv LookupEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:getEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEventSubscriptionArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	Scope                 string `pulumi:"scope"`
}


type LookupEventSubscriptionResult struct {
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


func (val *LookupEventSubscriptionResult) Defaults() *LookupEventSubscriptionResult {
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
