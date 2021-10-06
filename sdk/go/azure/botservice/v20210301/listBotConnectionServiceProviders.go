


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBotConnectionServiceProviders(ctx *pulumi.Context, args *ListBotConnectionServiceProvidersArgs, opts ...pulumi.InvokeOption) (*ListBotConnectionServiceProvidersResult, error) {
	var rv ListBotConnectionServiceProvidersResult
	err := ctx.Invoke("azure-native:botservice/v20210301:listBotConnectionServiceProviders", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBotConnectionServiceProvidersArgs struct {
}


type ListBotConnectionServiceProvidersResult struct {
	NextLink *string                   `pulumi:"nextLink"`
	Value    []ServiceProviderResponse `pulumi:"value"`
}
