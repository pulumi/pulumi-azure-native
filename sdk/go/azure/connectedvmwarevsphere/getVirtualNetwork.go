


package connectedvmwarevsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}


type LookupVirtualNetworkResult struct {
	CustomResourceName string                    `pulumi:"customResourceName"`
	ExtendedLocation   *ExtendedLocationResponse `pulumi:"extendedLocation"`
	Id                 string                    `pulumi:"id"`
	InventoryItemId    *string                   `pulumi:"inventoryItemId"`
	Kind               *string                   `pulumi:"kind"`
	Location           string                    `pulumi:"location"`
	MoName             string                    `pulumi:"moName"`
	MoRefId            *string                   `pulumi:"moRefId"`
	Name               string                    `pulumi:"name"`
	ProvisioningState  string                    `pulumi:"provisioningState"`
	Statuses           []ResourceStatusResponse  `pulumi:"statuses"`
	SystemData         SystemDataResponse        `pulumi:"systemData"`
	Tags               map[string]string         `pulumi:"tags"`
	Type               string                    `pulumi:"type"`
	Uuid               string                    `pulumi:"uuid"`
	VCenterId          *string                   `pulumi:"vCenterId"`
}
