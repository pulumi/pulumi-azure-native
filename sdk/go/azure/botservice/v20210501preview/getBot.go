


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBot(ctx *pulumi.Context, args *LookupBotArgs, opts ...pulumi.InvokeOption) (*LookupBotResult, error) {
	var rv LookupBotResult
	err := ctx.Invoke("azure-native:botservice/v20210501preview:getBot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupBotResult struct {
	Etag       *string               `pulumi:"etag"`
	Id         string                `pulumi:"id"`
	Kind       *string               `pulumi:"kind"`
	Location   *string               `pulumi:"location"`
	Name       string                `pulumi:"name"`
	Properties BotPropertiesResponse `pulumi:"properties"`
	Sku        *SkuResponse          `pulumi:"sku"`
	Tags       map[string]string     `pulumi:"tags"`
	Type       string                `pulumi:"type"`
}
