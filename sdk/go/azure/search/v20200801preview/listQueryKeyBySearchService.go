


package v20200801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListQueryKeyBySearchService(ctx *pulumi.Context, args *ListQueryKeyBySearchServiceArgs, opts ...pulumi.InvokeOption) (*ListQueryKeyBySearchServiceResult, error) {
	var rv ListQueryKeyBySearchServiceResult
	err := ctx.Invoke("azure-native:search/v20200801preview:listQueryKeyBySearchService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListQueryKeyBySearchServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SearchServiceName string `pulumi:"searchServiceName"`
}


type ListQueryKeyBySearchServiceResult struct {
	NextLink string             `pulumi:"nextLink"`
	Value    []QueryKeyResponse `pulumi:"value"`
}
