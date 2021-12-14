


package hdinsight

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:hdinsight:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupClusterResult struct {
	Etag       *string                      `pulumi:"etag"`
	Id         string                       `pulumi:"id"`
	Identity   *ClusterIdentityResponse     `pulumi:"identity"`
	Location   *string                      `pulumi:"location"`
	Name       string                       `pulumi:"name"`
	Properties ClusterGetPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string            `pulumi:"tags"`
	Type       string                       `pulumi:"type"`
}


func (val *LookupClusterResult) Defaults() *LookupClusterResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
