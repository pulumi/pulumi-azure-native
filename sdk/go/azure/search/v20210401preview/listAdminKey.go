


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAdminKey(ctx *pulumi.Context, args *ListAdminKeyArgs, opts ...pulumi.InvokeOption) (*ListAdminKeyResult, error) {
	var rv ListAdminKeyResult
	err := ctx.Invoke("azure-native:search/v20210401preview:listAdminKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAdminKeyArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SearchServiceName string `pulumi:"searchServiceName"`
}


type ListAdminKeyResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}
