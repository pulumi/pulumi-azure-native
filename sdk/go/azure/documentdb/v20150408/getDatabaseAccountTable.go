


package v20150408

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccountTable(ctx *pulumi.Context, args *LookupDatabaseAccountTableArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountTableResult, error) {
	var rv LookupDatabaseAccountTableResult
	err := ctx.Invoke("azure-native:documentdb/v20150408:getDatabaseAccountTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountTableArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableName         string `pulumi:"tableName"`
}


type LookupDatabaseAccountTableResult struct {
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}
