


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCassandraCluster(ctx *pulumi.Context, args *LookupCassandraClusterArgs, opts ...pulumi.InvokeOption) (*LookupCassandraClusterResult, error) {
	var rv LookupCassandraClusterResult
	err := ctx.Invoke("azure-native:documentdb/v20210401preview:getCassandraCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCassandraClusterResult struct {
	Id         string                            `pulumi:"id"`
	Identity   *ManagedServiceIdentityResponse   `pulumi:"identity"`
	Location   *string                           `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties ClusterResourceResponseProperties `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}
