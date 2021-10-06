


package webpubsub

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebPubSubHub(ctx *pulumi.Context, args *LookupWebPubSubHubArgs, opts ...pulumi.InvokeOption) (*LookupWebPubSubHubResult, error) {
	var rv LookupWebPubSubHubResult
	err := ctx.Invoke("azure-native:webpubsub:getWebPubSubHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebPubSubHubArgs struct {
	HubName           string `pulumi:"hubName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWebPubSubHubResult struct {
	Id         string                         `pulumi:"id"`
	Name       string                         `pulumi:"name"`
	Properties WebPubSubHubPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse             `pulumi:"systemData"`
	Type       string                         `pulumi:"type"`
}
