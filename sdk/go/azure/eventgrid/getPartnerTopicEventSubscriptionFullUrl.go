


package eventgrid

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPartnerTopicEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetPartnerTopicEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetPartnerTopicEventSubscriptionFullUrlResult, error) {
	var rv GetPartnerTopicEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid:getPartnerTopicEventSubscriptionFullUrl", args, &rv, opts...)
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
