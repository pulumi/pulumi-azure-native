


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabaseAccountConnectionStrings(ctx *pulumi.Context, args *ListDatabaseAccountConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListDatabaseAccountConnectionStringsResult, error) {
	var rv ListDatabaseAccountConnectionStringsResult
	err := ctx.Invoke("azure-native:documentdb/v20210401preview:listDatabaseAccountConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabaseAccountConnectionStringsArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDatabaseAccountConnectionStringsResult struct {
	ConnectionStrings []DatabaseAccountConnectionStringResponse `pulumi:"connectionStrings"`
}
