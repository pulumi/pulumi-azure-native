


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:streamanalytics/v20200301:getCluster", args, &rv, opts...)
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
	CapacityAllocated int                 `pulumi:"capacityAllocated"`
	CapacityAssigned  int                 `pulumi:"capacityAssigned"`
	ClusterId         string              `pulumi:"clusterId"`
	CreatedDate       string              `pulumi:"createdDate"`
	Etag              string              `pulumi:"etag"`
	Id                string              `pulumi:"id"`
	Location          *string             `pulumi:"location"`
	Name              string              `pulumi:"name"`
	ProvisioningState string              `pulumi:"provisioningState"`
	Sku               *ClusterSkuResponse `pulumi:"sku"`
	Tags              map[string]string   `pulumi:"tags"`
	Type              string              `pulumi:"type"`
}
