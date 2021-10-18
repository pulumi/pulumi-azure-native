


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetEventSubscriptionDeliveryAttributesResult, error) {
	var rv GetEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid/v20210601preview:getEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEventSubscriptionDeliveryAttributesArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	Scope                 string `pulumi:"scope"`
}


type GetEventSubscriptionDeliveryAttributesResult struct {
	Value []interface{} `pulumi:"value"`
}
