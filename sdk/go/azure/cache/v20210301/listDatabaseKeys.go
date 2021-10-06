


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabaseKeys(ctx *pulumi.Context, args *ListDatabaseKeysArgs, opts ...pulumi.InvokeOption) (*ListDatabaseKeysResult, error) {
	var rv ListDatabaseKeysResult
	err := ctx.Invoke("azure-native:cache/v20210301:listDatabaseKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseKeysArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDatabaseKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}
