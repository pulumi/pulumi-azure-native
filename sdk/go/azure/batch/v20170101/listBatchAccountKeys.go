


package v20170101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBatchAccountKeys(ctx *pulumi.Context, args *ListBatchAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListBatchAccountKeysResult, error) {
	var rv ListBatchAccountKeysResult
	err := ctx.Invoke("azure-native:batch/v20170101:listBatchAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBatchAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListBatchAccountKeysResult struct {
	Primary   string `pulumi:"primary"`
	Secondary string `pulumi:"secondary"`
}
