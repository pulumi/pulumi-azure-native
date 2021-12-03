


package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProximityPlacementGroup(ctx *pulumi.Context, args *LookupProximityPlacementGroupArgs, opts ...pulumi.InvokeOption) (*LookupProximityPlacementGroupResult, error) {
	var rv LookupProximityPlacementGroupResult
	err := ctx.Invoke("azure-native:compute/v20180601:getProximityPlacementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProximityPlacementGroupArgs struct {
	ProximityPlacementGroupName string `pulumi:"proximityPlacementGroupName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
}


type LookupProximityPlacementGroupResult struct {
	AvailabilitySets            []SubResourceResponse `pulumi:"availabilitySets"`
	Id                          string                `pulumi:"id"`
	Location                    string                `pulumi:"location"`
	Name                        string                `pulumi:"name"`
	ProximityPlacementGroupType *string               `pulumi:"proximityPlacementGroupType"`
	Tags                        map[string]string     `pulumi:"tags"`
	Type                        string                `pulumi:"type"`
	VirtualMachineScaleSets     []SubResourceResponse `pulumi:"virtualMachineScaleSets"`
	VirtualMachines             []SubResourceResponse `pulumi:"virtualMachines"`
}
