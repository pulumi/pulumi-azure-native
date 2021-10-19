


package mixedreality

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRemoteRenderingAccountKeys(ctx *pulumi.Context, args *ListRemoteRenderingAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListRemoteRenderingAccountKeysResult, error) {
	var rv ListRemoteRenderingAccountKeysResult
	err := ctx.Invoke("azure-native:mixedreality:listRemoteRenderingAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemoteRenderingAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListRemoteRenderingAccountKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}
