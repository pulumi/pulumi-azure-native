


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInstanceFailoverGroup(ctx *pulumi.Context, args *LookupInstanceFailoverGroupArgs, opts ...pulumi.InvokeOption) (*LookupInstanceFailoverGroupResult, error) {
	var rv LookupInstanceFailoverGroupResult
	err := ctx.Invoke("azure-native:sql:getInstanceFailoverGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceFailoverGroupArgs struct {
	FailoverGroupName string `pulumi:"failoverGroupName"`
	LocationName      string `pulumi:"locationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInstanceFailoverGroupResult struct {
	Id                   string                                         `pulumi:"id"`
	ManagedInstancePairs []ManagedInstancePairInfoResponse              `pulumi:"managedInstancePairs"`
	Name                 string                                         `pulumi:"name"`
	PartnerRegions       []PartnerRegionInfoResponse                    `pulumi:"partnerRegions"`
	ReadOnlyEndpoint     *InstanceFailoverGroupReadOnlyEndpointResponse `pulumi:"readOnlyEndpoint"`
	ReadWriteEndpoint    InstanceFailoverGroupReadWriteEndpointResponse `pulumi:"readWriteEndpoint"`
	ReplicationRole      string                                         `pulumi:"replicationRole"`
	ReplicationState     string                                         `pulumi:"replicationState"`
	Type                 string                                         `pulumi:"type"`
}
