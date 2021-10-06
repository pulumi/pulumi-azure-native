


package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContentItem(ctx *pulumi.Context, args *LookupContentItemArgs, opts ...pulumi.InvokeOption) (*LookupContentItemResult, error) {
	var rv LookupContentItemResult
	err := ctx.Invoke("azure-native:apimanagement:getContentItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentItemArgs struct {
	ContentItemId     string `pulumi:"contentItemId"`
	ContentTypeId     string `pulumi:"contentTypeId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupContentItemResult struct {
	Id         string      `pulumi:"id"`
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}
