


package managednetwork

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedNetworkGroup(ctx *pulumi.Context, args *LookupManagedNetworkGroupArgs, opts ...pulumi.InvokeOption) (*LookupManagedNetworkGroupResult, error) {
	var rv LookupManagedNetworkGroupResult
	err := ctx.Invoke("azure-native:managednetwork:getManagedNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedNetworkGroupArgs struct {
	ManagedNetworkGroupName string `pulumi:"managedNetworkGroupName"`
	ManagedNetworkName      string `pulumi:"managedNetworkName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupManagedNetworkGroupResult struct {
	Etag              string               `pulumi:"etag"`
	Id                string               `pulumi:"id"`
	Kind              *string              `pulumi:"kind"`
	Location          *string              `pulumi:"location"`
	ManagementGroups  []ResourceIdResponse `pulumi:"managementGroups"`
	Name              string               `pulumi:"name"`
	ProvisioningState string               `pulumi:"provisioningState"`
	Subnets           []ResourceIdResponse `pulumi:"subnets"`
	Subscriptions     []ResourceIdResponse `pulumi:"subscriptions"`
	Type              string               `pulumi:"type"`
	VirtualNetworks   []ResourceIdResponse `pulumi:"virtualNetworks"`
}
