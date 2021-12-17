


package v20200313

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:search/v20200313:getService", args, &rv, opts...)
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
	HostingMode                *string                             `pulumi:"hostingMode"`
	Id                         string                              `pulumi:"id"`
	Identity                   *IdentityResponse                   `pulumi:"identity"`
	Location                   *string                             `pulumi:"location"`
	Name                       string                              `pulumi:"name"`
	NetworkRuleSet             *NetworkRuleSetResponse             `pulumi:"networkRuleSet"`
	PartitionCount             *int                                `pulumi:"partitionCount"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                             `pulumi:"publicNetworkAccess"`
	ReplicaCount               *int                                `pulumi:"replicaCount"`
	Sku                        *SkuResponse                        `pulumi:"sku"`
	Status                     string                              `pulumi:"status"`
	StatusDetails              string                              `pulumi:"statusDetails"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
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
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	if isZero(tmp.ReplicaCount) {
		replicaCount_ := 1
		tmp.ReplicaCount = &replicaCount_
	}
	return &tmp
}
