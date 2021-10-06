


package v20151106

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccountCassandraTable(ctx *pulumi.Context, args *LookupDatabaseAccountCassandraTableArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountCassandraTableResult, error) {
	var rv LookupDatabaseAccountCassandraTableResult
	err := ctx.Invoke("azure-native:documentdb/v20151106:getDatabaseAccountCassandraTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountCassandraTableArgs struct {
	AccountName       string `pulumi:"accountName"`
	KeyspaceName      string `pulumi:"keyspaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableName         string `pulumi:"tableName"`
}


type LookupDatabaseAccountCassandraTableResult struct {
	DefaultTtl *int                     `pulumi:"defaultTtl"`
	Id         string                   `pulumi:"id"`
	Location   *string                  `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Schema     *CassandraSchemaResponse `pulumi:"schema"`
	Tags       map[string]string        `pulumi:"tags"`
	Type       string                   `pulumi:"type"`
}
