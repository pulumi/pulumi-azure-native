


package v20210701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAccountKeys(ctx *pulumi.Context, args *ListAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListAccountKeysResult, error) {
	var rv ListAccountKeysResult
	err := ctx.Invoke("azure-native:maps/v20210701preview:listAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAccountKeysResult struct {
	PrimaryKey              string `pulumi:"primaryKey"`
	PrimaryKeyLastUpdated   string `pulumi:"primaryKeyLastUpdated"`
	SecondaryKey            string `pulumi:"secondaryKey"`
	SecondaryKeyLastUpdated string `pulumi:"secondaryKeyLastUpdated"`
}
