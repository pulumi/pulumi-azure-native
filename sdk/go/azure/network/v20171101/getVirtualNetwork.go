


package v20171101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure-native:network/v20171101:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkArgs struct {
	Expand             *string `pulumi:"expand"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type LookupVirtualNetworkResult struct {
	AddressSpace           *AddressSpaceResponse           `pulumi:"addressSpace"`
	DhcpOptions            *DhcpOptionsResponse            `pulumi:"dhcpOptions"`
	EnableDdosProtection   *bool                           `pulumi:"enableDdosProtection"`
	EnableVmProtection     *bool                           `pulumi:"enableVmProtection"`
	Etag                   *string                         `pulumi:"etag"`
	Id                     *string                         `pulumi:"id"`
	Location               *string                         `pulumi:"location"`
	Name                   string                          `pulumi:"name"`
	ProvisioningState      *string                         `pulumi:"provisioningState"`
	ResourceGuid           *string                         `pulumi:"resourceGuid"`
	Subnets                []SubnetResponse                `pulumi:"subnets"`
	Tags                   map[string]string               `pulumi:"tags"`
	Type                   string                          `pulumi:"type"`
	VirtualNetworkPeerings []VirtualNetworkPeeringResponse `pulumi:"virtualNetworkPeerings"`
}
