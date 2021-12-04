


package eventhub

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:eventhub:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupClusterResult struct {
	CreatedAt string              `pulumi:"createdAt"`
	Id        string              `pulumi:"id"`
	Location  *string             `pulumi:"location"`
	MetricId  string              `pulumi:"metricId"`
	Name      string              `pulumi:"name"`
	Sku       *ClusterSkuResponse `pulumi:"sku"`
	Status    string              `pulumi:"status"`
	Tags      map[string]string   `pulumi:"tags"`
	Type      string              `pulumi:"type"`
	UpdatedAt string              `pulumi:"updatedAt"`
}
