


package v20151001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMediaServiceKeys(ctx *pulumi.Context, args *ListMediaServiceKeysArgs, opts ...pulumi.InvokeOption) (*ListMediaServiceKeysResult, error) {
	var rv ListMediaServiceKeysResult
	err := ctx.Invoke("azure-native:media/v20151001:listMediaServiceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMediaServiceKeysArgs struct {
	MediaServiceName  string `pulumi:"mediaServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMediaServiceKeysResult struct {
	PrimaryAuthEndpoint   *string `pulumi:"primaryAuthEndpoint"`
	PrimaryKey            *string `pulumi:"primaryKey"`
	Scope                 *string `pulumi:"scope"`
	SecondaryAuthEndpoint *string `pulumi:"secondaryAuthEndpoint"`
	SecondaryKey          *string `pulumi:"secondaryKey"`
}
