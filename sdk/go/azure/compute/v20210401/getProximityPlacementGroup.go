


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProximityPlacementGroup(ctx *pulumi.Context, args *LookupProximityPlacementGroupArgs, opts ...pulumi.InvokeOption) (*LookupProximityPlacementGroupResult, error) {
	var rv LookupProximityPlacementGroupResult
	err := ctx.Invoke("azure-native:compute/v20210401:getProximityPlacementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProximityPlacementGroupArgs struct {
	IncludeColocationStatus     *string `pulumi:"includeColocationStatus"`
	ProximityPlacementGroupName string  `pulumi:"proximityPlacementGroupName"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
}


type LookupProximityPlacementGroupResult struct {
	AvailabilitySets            []SubResourceWithColocationStatusResponse `pulumi:"availabilitySets"`
	ColocationStatus            *InstanceViewStatusResponse               `pulumi:"colocationStatus"`
	Id                          string                                    `pulumi:"id"`
	Location                    string                                    `pulumi:"location"`
	Name                        string                                    `pulumi:"name"`
	ProximityPlacementGroupType *string                                   `pulumi:"proximityPlacementGroupType"`
	Tags                        map[string]string                         `pulumi:"tags"`
	Type                        string                                    `pulumi:"type"`
	VirtualMachineScaleSets     []SubResourceWithColocationStatusResponse `pulumi:"virtualMachineScaleSets"`
	VirtualMachines             []SubResourceWithColocationStatusResponse `pulumi:"virtualMachines"`
}
