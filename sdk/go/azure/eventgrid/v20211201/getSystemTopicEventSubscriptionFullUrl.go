


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemTopicEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetSystemTopicEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetSystemTopicEventSubscriptionFullUrlResult, error) {
	var rv GetSystemTopicEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20211201:getSystemTopicEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSystemTopicEventSubscriptionFullUrlArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	SystemTopicName       string `pulumi:"systemTopicName"`
}


type GetSystemTopicEventSubscriptionFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}
