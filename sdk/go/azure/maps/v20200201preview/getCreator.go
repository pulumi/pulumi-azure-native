


package v20200201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCreator(ctx *pulumi.Context, args *LookupCreatorArgs, opts ...pulumi.InvokeOption) (*LookupCreatorResult, error) {
	var rv LookupCreatorResult
	err := ctx.Invoke("azure-native:maps/v20200201preview:getCreator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCreatorArgs struct {
	AccountName       string `pulumi:"accountName"`
	CreatorName       string `pulumi:"creatorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCreatorResult struct {
	Id         string                    `pulumi:"id"`
	Location   string                    `pulumi:"location"`
	Name       string                    `pulumi:"name"`
	Properties CreatorPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string         `pulumi:"tags"`
	Type       string                    `pulumi:"type"`
}
