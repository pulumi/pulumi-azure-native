


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListShareSynchronizations(ctx *pulumi.Context, args *ListShareSynchronizationsArgs, opts ...pulumi.InvokeOption) (*ListShareSynchronizationsResult, error) {
	var rv ListShareSynchronizationsResult
	err := ctx.Invoke("azure-native:datashare/v20191101:listShareSynchronizations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListShareSynchronizationsArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Filter            *string `pulumi:"filter"`
	Orderby           *string `pulumi:"orderby"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ShareName         string  `pulumi:"shareName"`
	SkipToken         *string `pulumi:"skipToken"`
}


type ListShareSynchronizationsResult struct {
	NextLink *string                        `pulumi:"nextLink"`
	Value    []ShareSynchronizationResponse `pulumi:"value"`
}
