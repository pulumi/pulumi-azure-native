


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInventoryItem(ctx *pulumi.Context, args *LookupInventoryItemArgs, opts ...pulumi.InvokeOption) (*LookupInventoryItemResult, error) {
	var rv LookupInventoryItemResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20201001preview:getInventoryItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInventoryItemArgs struct {
	InventoryItemName string `pulumi:"inventoryItemName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VcenterName       string `pulumi:"vcenterName"`
}


type LookupInventoryItemResult struct {
	Id                string             `pulumi:"id"`
	InventoryType     string             `pulumi:"inventoryType"`
	Kind              *string            `pulumi:"kind"`
	ManagedResourceId *string            `pulumi:"managedResourceId"`
	MoName            *string            `pulumi:"moName"`
	MoRefId           *string            `pulumi:"moRefId"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}
