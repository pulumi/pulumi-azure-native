


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemTopicEventSubscription(ctx *pulumi.Context, args *LookupSystemTopicEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSystemTopicEventSubscriptionResult, error) {
	var rv LookupSystemTopicEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20211201:getSystemTopicEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSystemTopicEventSubscriptionArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SystemTopicName       string `pulumi:"systemTopicName"`
}


type LookupSystemTopicEventSubscriptionResult struct {
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
	SystemData            SystemDataResponse                        `pulumi:"systemData"`
	Topic                 string                                    `pulumi:"topic"`
	Type                  string                                    `pulumi:"type"`
}
