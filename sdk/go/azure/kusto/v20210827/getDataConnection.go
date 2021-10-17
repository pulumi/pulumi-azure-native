


package v20210827

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDataConnection(ctx *pulumi.Context, args *LookupDataConnectionArgs, opts ...pulumi.InvokeOption) (*LookupDataConnectionResult, error) {
	var rv LookupDataConnectionResult
	err := ctx.Invoke("azure-native:kusto/v20210827:getDataConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataConnectionArgs struct {
	ClusterName        string `pulumi:"clusterName"`
	DataConnectionName string `pulumi:"dataConnectionName"`
	DatabaseName       string `pulumi:"databaseName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupDataConnectionResult struct {
	Id       string  `pulumi:"id"`
	Kind     string  `pulumi:"kind"`
	Location *string `pulumi:"location"`
	Name     string  `pulumi:"name"`
	Type     string  `pulumi:"type"`
}
