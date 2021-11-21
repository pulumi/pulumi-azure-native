


package v20211015

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCassandraDataCenter(ctx *pulumi.Context, args *LookupCassandraDataCenterArgs, opts ...pulumi.InvokeOption) (*LookupCassandraDataCenterResult, error) {
	var rv LookupCassandraDataCenterResult
	err := ctx.Invoke("azure-native:documentdb/v20211015:getCassandraDataCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCassandraDataCenterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	DataCenterName    string `pulumi:"dataCenterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCassandraDataCenterResult struct {
	Id         string                               `pulumi:"id"`
	Name       string                               `pulumi:"name"`
	Properties DataCenterResourceResponseProperties `pulumi:"properties"`
	Type       string                               `pulumi:"type"`
}
