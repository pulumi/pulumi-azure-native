


package v20170615preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventSubscription(ctx *pulumi.Context, args *LookupEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupEventSubscriptionResult, error) {
	var rv LookupEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20170615preview:getEventSubscription", args, &rv, opts...)
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
	Destination       *EventSubscriptionDestinationResponse `pulumi:"destination"`
	Filter            *EventSubscriptionFilterResponse      `pulumi:"filter"`
	Id                string                                `pulumi:"id"`
	Labels            []string                              `pulumi:"labels"`
	Name              string                                `pulumi:"name"`
	ProvisioningState string                                `pulumi:"provisioningState"`
	Topic             string                                `pulumi:"topic"`
	Type              string                                `pulumi:"type"`
}


func (val *LookupEventSubscriptionResult) Defaults() *LookupEventSubscriptionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Filter = tmp.Filter.Defaults()

	return &tmp
}
