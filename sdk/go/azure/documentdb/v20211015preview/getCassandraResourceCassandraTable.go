


package v20211015preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCassandraResourceCassandraTable(ctx *pulumi.Context, args *LookupCassandraResourceCassandraTableArgs, opts ...pulumi.InvokeOption) (*LookupCassandraResourceCassandraTableResult, error) {
	var rv LookupCassandraResourceCassandraTableResult
	err := ctx.Invoke("azure-native:documentdb/v20211015preview:getCassandraResourceCassandraTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraResourceCassandraTableArgs struct {
	AccountName       string `pulumi:"accountName"`
	KeyspaceName      string `pulumi:"keyspaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TableName         string `pulumi:"tableName"`
}


type LookupCassandraResourceCassandraTableResult struct {
	Id       string                                       `pulumi:"id"`
	Identity *ManagedServiceIdentityResponse              `pulumi:"identity"`
	Location *string                                      `pulumi:"location"`
	Name     string                                       `pulumi:"name"`
	Options  *CassandraTableGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *CassandraTableGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                            `pulumi:"tags"`
	Type     string                                       `pulumi:"type"`
}
