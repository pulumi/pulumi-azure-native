


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSpatialAnchorsAccountKeys(ctx *pulumi.Context, args *ListSpatialAnchorsAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListSpatialAnchorsAccountKeysResult, error) {
	var rv ListSpatialAnchorsAccountKeysResult
	err := ctx.Invoke("azure-native:mixedreality/v20210301preview:listSpatialAnchorsAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSpatialAnchorsAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSpatialAnchorsAccountKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}
