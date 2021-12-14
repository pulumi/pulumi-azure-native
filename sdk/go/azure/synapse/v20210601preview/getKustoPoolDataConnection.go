


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupKustoPoolDataConnection(ctx *pulumi.Context, args *LookupKustoPoolDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolDataConnectionResult, error) {
	var rv LookupKustoPoolDataConnectionResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getKustoPoolDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolDataConnectionArgs struct {
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	KustoPoolName      string `pulumi:"kustoPoolName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	WorkspaceName      string `pulumi:"workspaceName"`
}


type LookupKustoPoolDataConnectionResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}
