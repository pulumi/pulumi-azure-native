


package v20190301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAvailabilitySet(ctx *pulumi.Context, args *LookupAvailabilitySetArgs, opts ...pulumi.InvokeOption) (*LookupAvailabilitySetResult, error) {
	var rv LookupAvailabilitySetResult
	err := ctx.Invoke("azure-native:compute/v20190301:getAvailabilitySet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAvailabilitySetArgs struct {
	AvailabilitySetName string `pulumi:"availabilitySetName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupAvailabilitySetResult struct {
	Id                        string                       `pulumi:"id"`
	Location                  string                       `pulumi:"location"`
	Name                      string                       `pulumi:"name"`
	PlatformFaultDomainCount  *int                         `pulumi:"platformFaultDomainCount"`
	PlatformUpdateDomainCount *int                         `pulumi:"platformUpdateDomainCount"`
	ProximityPlacementGroup   *SubResourceResponse         `pulumi:"proximityPlacementGroup"`
	Sku                       *SkuResponse                 `pulumi:"sku"`
	Statuses                  []InstanceViewStatusResponse `pulumi:"statuses"`
	Tags                      map[string]string            `pulumi:"tags"`
	Type                      string                       `pulumi:"type"`
	VirtualMachines           []SubResourceResponse        `pulumi:"virtualMachines"`
}
