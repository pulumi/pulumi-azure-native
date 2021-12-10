


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListProductFamilies(ctx *pulumi.Context, args *ListProductFamiliesArgs, opts ...pulumi.InvokeOption) (*ListProductFamiliesResult, error) {
	var rv ListProductFamiliesResult
	err := ctx.Invoke("azure-native:edgeorder/v20211201:listProductFamilies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductFamiliesArgs struct {
	CustomerSubscriptionDetails *CustomerSubscriptionDetails    `pulumi:"customerSubscriptionDetails"`
	Expand                      *string                         `pulumi:"expand"`
	FilterableProperties        map[string][]FilterableProperty `pulumi:"filterableProperties"`
	SkipToken                   *string                         `pulumi:"skipToken"`
}


type ListProductFamiliesResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []ProductFamilyResponse `pulumi:"value"`
}
