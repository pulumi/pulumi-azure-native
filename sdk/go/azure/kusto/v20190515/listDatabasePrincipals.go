


package v20190515

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDatabasePrincipals(ctx *pulumi.Context, args *ListDatabasePrincipalsArgs, opts ...pulumi.InvokeOption) (*ListDatabasePrincipalsResult, error) {
	var rv ListDatabasePrincipalsResult
	err := ctx.Invoke("azure-native:kusto/v20190515:listDatabasePrincipals", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabasePrincipalsArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DatabaseName      string `pulumi:"databaseName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListDatabasePrincipalsResult struct {
	Value []DatabasePrincipalResponse `pulumi:"value"`
}
