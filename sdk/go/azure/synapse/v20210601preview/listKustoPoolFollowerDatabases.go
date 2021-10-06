


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListKustoPoolFollowerDatabases(ctx *pulumi.Context, args *ListKustoPoolFollowerDatabasesArgs, opts ...pulumi.InvokeOption) (*ListKustoPoolFollowerDatabasesResult, error) {
	var rv ListKustoPoolFollowerDatabasesResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:listKustoPoolFollowerDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKustoPoolFollowerDatabasesArgs struct {
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListKustoPoolFollowerDatabasesResult struct {
	Value []FollowerDatabaseDefinitionResponse `pulumi:"value"`
}
