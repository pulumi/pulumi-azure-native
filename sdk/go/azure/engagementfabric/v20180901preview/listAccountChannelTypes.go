


package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAccountChannelTypes(ctx *pulumi.Context, args *ListAccountChannelTypesArgs, opts ...pulumi.InvokeOption) (*ListAccountChannelTypesResult, error) {
	var rv ListAccountChannelTypesResult
	err := ctx.Invoke("azure-native:engagementfabric/v20180901preview:listAccountChannelTypes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountChannelTypesArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAccountChannelTypesResult struct {
	Value []ChannelTypeDescriptionResponse `pulumi:"value"`
}
