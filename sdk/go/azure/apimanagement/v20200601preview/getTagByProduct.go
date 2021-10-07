


package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagByProduct(ctx *pulumi.Context, args *LookupTagByProductArgs, opts ...pulumi.InvokeOption) (*LookupTagByProductResult, error) {
	var rv LookupTagByProductResult
	err := ctx.Invoke("azure-native:apimanagement/v20200601preview:getTagByProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagByProductArgs struct {
	ProductId         string `pulumi:"productId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	TagId             string `pulumi:"tagId"`
}


type LookupTagByProductResult struct {
	DisplayName string `pulumi:"displayName"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	Type        string `pulumi:"type"`
}
