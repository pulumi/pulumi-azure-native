


package v20201015preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPartnerTopicEventSubscriptionDeliveryAttributes(ctx *pulumi.Context, args *GetPartnerTopicEventSubscriptionDeliveryAttributesArgs, opts ...pulumi.InvokeOption) (*GetPartnerTopicEventSubscriptionDeliveryAttributesResult, error) {
	var rv GetPartnerTopicEventSubscriptionDeliveryAttributesResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:getPartnerTopicEventSubscriptionDeliveryAttributes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPartnerTopicEventSubscriptionDeliveryAttributesArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	PartnerTopicName      string `pulumi:"partnerTopicName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type GetPartnerTopicEventSubscriptionDeliveryAttributesResult struct {
	Value []interface{} `pulumi:"value"`
}
