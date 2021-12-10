


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListObjectAnchorsAccountKeys(ctx *pulumi.Context, args *ListObjectAnchorsAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListObjectAnchorsAccountKeysResult, error) {
	var rv ListObjectAnchorsAccountKeysResult
	err := ctx.Invoke("azure-native:mixedreality/v20210301preview:listObjectAnchorsAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListObjectAnchorsAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListObjectAnchorsAccountKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}
