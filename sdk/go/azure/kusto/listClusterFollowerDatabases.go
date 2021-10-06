


package kusto

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListClusterFollowerDatabases(ctx *pulumi.Context, args *ListClusterFollowerDatabasesArgs, opts ...pulumi.InvokeOption) (*ListClusterFollowerDatabasesResult, error) {
	var rv ListClusterFollowerDatabasesResult
	err := ctx.Invoke("azure-native:kusto:listClusterFollowerDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListClusterFollowerDatabasesArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListClusterFollowerDatabasesResult struct {
	Value []FollowerDatabaseDefinitionResponse `pulumi:"value"`
}
