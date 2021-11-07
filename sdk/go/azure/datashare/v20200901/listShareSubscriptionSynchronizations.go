


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListShareSubscriptionSynchronizations(ctx *pulumi.Context, args *ListShareSubscriptionSynchronizationsArgs, opts ...pulumi.InvokeOption) (*ListShareSubscriptionSynchronizationsResult, error) {
	var rv ListShareSubscriptionSynchronizationsResult
	err := ctx.Invoke("azure-native:datashare/v20200901:listShareSubscriptionSynchronizations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListShareSubscriptionSynchronizationsArgs struct {
	AccountName           string  `pulumi:"accountName"`
	Filter                *string `pulumi:"filter"`
	Orderby               *string `pulumi:"orderby"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	SkipToken             *string `pulumi:"skipToken"`
}


type ListShareSubscriptionSynchronizationsResult struct {
	NextLink *string                                    `pulumi:"nextLink"`
	Value    []ShareSubscriptionSynchronizationResponse `pulumi:"value"`
}
