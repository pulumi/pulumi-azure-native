


package recoveryservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationProtectedItem(ctx *pulumi.Context, args *LookupReplicationProtectedItemArgs, opts ...pulumi.InvokeOption) (*LookupReplicationProtectedItemResult, error) {
	var rv LookupReplicationProtectedItemResult
	err := ctx.Invoke("azure-native:recoveryservices:getReplicationProtectedItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationProtectedItemArgs struct {
	FabricName                  string `pulumi:"fabricName"`
	ProtectionContainerName     string `pulumi:"protectionContainerName"`
	ReplicatedProtectedItemName string `pulumi:"replicatedProtectedItemName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	ResourceName                string `pulumi:"resourceName"`
}


type LookupReplicationProtectedItemResult struct {
	Id         string                                     `pulumi:"id"`
	Location   *string                                    `pulumi:"location"`
	Name       string                                     `pulumi:"name"`
	Properties ReplicationProtectedItemPropertiesResponse `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}
