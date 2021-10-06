


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListShareSubscriptionSynchronizationDetails(ctx *pulumi.Context, args *ListShareSubscriptionSynchronizationDetailsArgs, opts ...pulumi.InvokeOption) (*ListShareSubscriptionSynchronizationDetailsResult, error) {
	var rv ListShareSubscriptionSynchronizationDetailsResult
	err := ctx.Invoke("azure-native:datashare/v20200901:listShareSubscriptionSynchronizationDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListShareSubscriptionSynchronizationDetailsArgs struct {
	AccountName           string  `pulumi:"accountName"`
	Filter                *string `pulumi:"filter"`
	Orderby               *string `pulumi:"orderby"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	SkipToken             *string `pulumi:"skipToken"`
	SynchronizationId     string  `pulumi:"synchronizationId"`
}


type ListShareSubscriptionSynchronizationDetailsResult struct {
	NextLink *string                          `pulumi:"nextLink"`
	Value    []SynchronizationDetailsResponse `pulumi:"value"`
}
