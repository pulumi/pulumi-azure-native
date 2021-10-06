


package v20180301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRedisKeys(ctx *pulumi.Context, args *ListRedisKeysArgs, opts ...pulumi.InvokeOption) (*ListRedisKeysResult, error) {
	var rv ListRedisKeysResult
	err := ctx.Invoke("azure-native:cache/v20180301:listRedisKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRedisKeysArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListRedisKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}
