


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHost(ctx *pulumi.Context, args *LookupHostArgs, opts ...pulumi.InvokeOption) (*LookupHostResult, error) {
	var rv LookupHostResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20201001preview:getHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostArgs struct {
	HostName          string `pulumi:"hostName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupHostResult struct {
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
