


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetEventSubscriptionFullUrl(ctx *pulumi.Context, args *GetEventSubscriptionFullUrlArgs, opts ...pulumi.InvokeOption) (*GetEventSubscriptionFullUrlResult, error) {
	var rv GetEventSubscriptionFullUrlResult
	err := ctx.Invoke("azure-native:eventgrid/v20200101preview:getEventSubscriptionFullUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEventSubscriptionFullUrlArgs struct {
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	Scope                 string `pulumi:"scope"`
}


type GetEventSubscriptionFullUrlResult struct {
	EndpointUrl *string `pulumi:"endpointUrl"`
}
