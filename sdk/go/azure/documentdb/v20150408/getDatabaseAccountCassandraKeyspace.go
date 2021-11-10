


package v20150408

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatabaseAccountCassandraKeyspace(ctx *pulumi.Context, args *LookupDatabaseAccountCassandraKeyspaceArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseAccountCassandraKeyspaceResult, error) {
	var rv LookupDatabaseAccountCassandraKeyspaceResult
	err := ctx.Invoke("azure-native:documentdb/v20150408:getDatabaseAccountCassandraKeyspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseAccountCassandraKeyspaceArgs struct {
	AccountName       string `pulumi:"accountName"`
	KeyspaceName      string `pulumi:"keyspaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDatabaseAccountCassandraKeyspaceResult struct {
	Id       string            `pulumi:"id"`
	Location *string           `pulumi:"location"`
	Name     string            `pulumi:"name"`
	Tags     map[string]string `pulumi:"tags"`
	Type     string            `pulumi:"type"`
}
