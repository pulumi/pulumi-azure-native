


package v20150401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccountMongoDBDatabase(ctx *pulumi.Context, args *LookupDatabaseAccountMongoDBDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountMongoDBDatabaseResult, error) {
	var rv LookupDatabaseAccountMongoDBDatabaseResult
	err := ctx.Invoke("azure-native:documentdb/v20150401:getDatabaseAccountMongoDBDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountMongoDBDatabaseArgs struct {
	AccountName       string `pulumi:"accountName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountMongoDBDatabaseResult struct {
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}
