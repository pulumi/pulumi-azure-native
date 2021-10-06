


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupKustoPoolDatabase(ctx *pulumi.Context, args *LookupKustoPoolDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolDatabaseResult, error) {
	var rv LookupKustoPoolDatabaseResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getKustoPoolDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolDatabaseArgs struct {
	DatabaseName      string `pulumi:"databaseName"`
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupKustoPoolDatabaseResult struct {
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Location   *string            `pulumi:"location"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}
