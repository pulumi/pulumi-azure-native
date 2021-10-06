


package v20201015preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPartnerTopicEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetPartnerTopicEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetPartnerTopicEventSubscriptionFullUrlResult, error) {
	var rv GetPartnerTopicEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20201015preview:getPartnerTopicEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPartnerTopicEventSubscriptionFullUrlArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	PartnerTopicName      string `pulumi:"partnerTopicName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type GetPartnerTopicEventSubscriptionFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}
