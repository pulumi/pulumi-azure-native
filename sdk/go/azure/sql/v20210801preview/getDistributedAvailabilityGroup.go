


package v20210801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDistributedAvailabilityGroup(ctx *pulumi.Context, args *LookupDistributedAvailabilityGroupArgs, opts ...pulumi.InvokeOption) (*LookupDistributedAvailabilityGroupResult, error) {
	var rv LookupDistributedAvailabilityGroupResult
	err := ctx.Invoke("azure-native:sql/v20210801preview:getDistributedAvailabilityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDistributedAvailabilityGroupArgs struct {
	DistributedAvailabilityGroupName string `pulumi:"distributedAvailabilityGroupName"`
	ManagedInstanceName              string `pulumi:"managedInstanceName"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
}


type LookupDistributedAvailabilityGroupResult struct {
	DistributedAvailabilityGroupId string  `pulumi:"distributedAvailabilityGroupId"`
	Id                             string  `pulumi:"id"`
	LastHardenedLsn                string  `pulumi:"lastHardenedLsn"`
	LinkState                      string  `pulumi:"linkState"`
	Name                           string  `pulumi:"name"`
	PrimaryAvailabilityGroupName   *string `pulumi:"primaryAvailabilityGroupName"`
	ReplicationMode                *string `pulumi:"replicationMode"`
	SecondaryAvailabilityGroupName *string `pulumi:"secondaryAvailabilityGroupName"`
	SourceEndpoint                 *string `pulumi:"sourceEndpoint"`
	SourceReplicaId                string  `pulumi:"sourceReplicaId"`
	TargetDatabase                 *string `pulumi:"targetDatabase"`
	TargetReplicaId                string  `pulumi:"targetReplicaId"`
	Type                           string  `pulumi:"type"`
}
