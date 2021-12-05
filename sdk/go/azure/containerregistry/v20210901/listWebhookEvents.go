


package v20210901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebhookEvents(ctx *pulumi.Context, args *ListWebhookEventsArgs, opts ...pulumi.InvokeOption) (*ListWebhookEventsResult, error) {
	var rv ListWebhookEventsResult
	err := ctx.Invoke("azure-native:containerregistry/v20210901:listWebhookEvents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebhookEventsArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WebhookName       string `pulumi:"webhookName"`
}


type ListWebhookEventsResult struct {
	NextLink *string         `pulumi:"nextLink"`
	Value    []EventResponse `pulumi:"value"`
}
