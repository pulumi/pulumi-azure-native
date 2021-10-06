


package v20180712

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	var rv LookupChannelResult
	err := ctx.Invoke("azure-native:botservice/v20180712:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelArgs struct {
	ChannelName       string `pulumi:"channelName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupChannelResult struct {
	Etag       *string           `pulumi:"etag"`
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}
