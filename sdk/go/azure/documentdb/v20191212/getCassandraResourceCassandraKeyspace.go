


package v20191212

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCassandraResourceCassandraKeyspace(ctx *pulumi.Context, args *LookupCassandraResourceCassandraKeyspaceArgs, opts ...pulumi.InvokeOption) (*LookupCassandraResourceCassandraKeyspaceResult, error) {
	var rv LookupCassandraResourceCassandraKeyspaceResult
	err := ctx.Invoke("azure-native:documentdb/v20191212:getCassandraResourceCassandraKeyspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraResourceCassandraKeyspaceArgs struct {
	AccountName       string `pulumi:"accountName"`
	KeyspaceName      string `pulumi:"keyspaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCassandraResourceCassandraKeyspaceResult struct {
	Id       string                                          `pulumi:"id"`
	Location *string                                         `pulumi:"location"`
	Name     string                                          `pulumi:"name"`
	Resource *CassandraKeyspaceGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                               `pulumi:"tags"`
	Type     string                                          `pulumi:"type"`
}
