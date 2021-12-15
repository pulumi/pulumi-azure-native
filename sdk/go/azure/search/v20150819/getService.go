


package v20150819

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:search/v20150819:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SearchServiceName string `pulumi:"searchServiceName"`
}


type LookupServiceResult struct {
	HostingMode       *string           `pulumi:"hostingMode"`
	Id                string            `pulumi:"id"`
	Identity          *IdentityResponse `pulumi:"identity"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	PartitionCount    *int              `pulumi:"partitionCount"`
	ProvisioningState string            `pulumi:"provisioningState"`
	ReplicaCount      *int              `pulumi:"replicaCount"`
	Sku               *SkuResponse      `pulumi:"sku"`
	Status            string            `pulumi:"status"`
	StatusDetails     string            `pulumi:"statusDetails"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}


func (val *LookupServiceResult) Defaults() *LookupServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HostingMode) {
		hostingMode_ := "default"
		tmp.HostingMode = &hostingMode_
	}
	if isZero(tmp.PartitionCount) {
		partitionCount_ := 1
		tmp.PartitionCount = &partitionCount_
	}
	if isZero(tmp.ReplicaCount) {
		replicaCount_ := 1
		tmp.ReplicaCount = &replicaCount_
	}
	return &tmp
}
