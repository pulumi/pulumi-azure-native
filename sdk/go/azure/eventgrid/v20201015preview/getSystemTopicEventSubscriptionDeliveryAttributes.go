


package v20201015preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemTopicEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetSystemTopicEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetSystemTopicEventSubscriptionDeliveryAttributesResult, error) {
	var rv GetSystemTopicEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:getSystemTopicEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSystemTopicEventSubscriptionDeliveryAttributesArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SystemTopicName       string `pulumi:"systemTopicName"`
}


type GetSystemTopicEventSubscriptionDeliveryAttributesResult struct {
	Value []interface{} `pulumi:"value"`
}
