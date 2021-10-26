


package v20171201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListChannelWithKeys(ctx *pulumi.Context, args *ListChannelWithKeysArgs, opts ...pulumi.InvokeOption) (*ListChannelWithKeysResult, error) {
	var rv ListChannelWithKeysResult
	err := ctx.Invoke("azure-native:botservice/v20171201:listChannelWithKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListChannelWithKeysArgs struct {
	ChannelName       string `pulumi:"channelName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListChannelWithKeysResult struct {
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
