


package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabaseAccountKeys(ctx *pulumi.Context, args *ListDatabaseAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListDatabaseAccountKeysResult, error) {
	var rv ListDatabaseAccountKeysResult
	err := ctx.Invoke("azure-native:documentdb/v20200601preview:listDatabaseAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDatabaseAccountKeysResult struct {
	PrimaryMasterKey           string `pulumi:"primaryMasterKey"`
	PrimaryReadonlyMasterKey   string `pulumi:"primaryReadonlyMasterKey"`
	SecondaryMasterKey         string `pulumi:"secondaryMasterKey"`
	SecondaryReadonlyMasterKey string `pulumi:"secondaryReadonlyMasterKey"`
}
