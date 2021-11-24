


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplication(ctx *pulumi.Context, args *LookupReplicationArgs, opts ...pulumi.InvokeOption) (*LookupReplicationResult, error) {
	var rv LookupReplicationResult
	err := ctx.Invoke("azure-native:containerregistry/v20190501:getReplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ReplicationName   string `pulumi:"replicationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupReplicationResult struct {
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState string            `pulumi:"provisioningState"`
	Status            StatusResponse    `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
