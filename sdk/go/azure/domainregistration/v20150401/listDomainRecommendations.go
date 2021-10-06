


package v20150401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDomainRecommendations(ctx *pulumi.Context, args *ListDomainRecommendationsArgs, opts ...pulumi.InvokeOption) (*ListDomainRecommendationsResult, error) {
	var rv ListDomainRecommendationsResult
	err := ctx.Invoke("azure-native:domainregistration/v20150401:listDomainRecommendations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDomainRecommendationsArgs struct {
	Keywords                 *string `pulumi:"keywords"`
	MaxDomainRecommendations *int    `pulumi:"maxDomainRecommendations"`
}


type ListDomainRecommendationsResult struct {
	NextLink string                   `pulumi:"nextLink"`
	Value    []NameIdentifierResponse `pulumi:"value"`
}
