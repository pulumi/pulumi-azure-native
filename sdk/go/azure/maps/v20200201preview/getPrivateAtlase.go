


package v20200201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateAtlase(ctx *pulumi.Context, args *LookupPrivateAtlaseArgs, opts ...pulumi.InvokeOption) (*LookupPrivateAtlaseResult, error) {
	var rv LookupPrivateAtlaseResult
	err := ctx.Invoke("azure-native:maps/v20200201preview:getPrivateAtlase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateAtlaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	PrivateAtlasName  string `pulumi:"privateAtlasName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPrivateAtlaseResult struct {
	Id         string                         `pulumi:"id"`
	Location   string                         `pulumi:"location"`
	Name       string                         `pulumi:"name"`
	Properties PrivateAtlasPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string              `pulumi:"tags"`
	Type       string                         `pulumi:"type"`
}
