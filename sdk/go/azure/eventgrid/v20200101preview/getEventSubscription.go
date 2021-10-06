


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventSubscription(ctx *pulumi.Context, args *LookupEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupEventSubscriptionResult, error) {
	var rv LookupEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20200101preview:getEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventSubscriptionArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	Scope                 string `pulumi:"scope"`
}


type LookupEventSubscriptionResult struct {
	DeadLetterDestination *StorageBlobDeadLetterDestinationResponse `pulumi:"deadLetterDestination"`
	Destination           interface{}                               `pulumi:"destination"`
	EventDeliverySchema   *string                                   `pulumi:"eventDeliverySchema"`
	ExpirationTimeUtc     *string                                   `pulumi:"expirationTimeUtc"`
	Filter                *EventSubscriptionFilterResponse          `pulumi:"filter"`
	Id                    string                                    `pulumi:"id"`
	Labels                []string                                  `pulumi:"labels"`
	Name                  string                                    `pulumi:"name"`
	ProvisioningState     string                                    `pulumi:"provisioningState"`
	RetryPolicy           *RetryPolicyResponse                      `pulumi:"retryPolicy"`
	Topic                 string                                    `pulumi:"topic"`
	Type                  string                                    `pulumi:"type"`
}
