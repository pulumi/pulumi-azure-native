


package avs

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkloadNetworkSegment(ctx *pulumi.Context, args *LookupWorkloadNetworkSegmentArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadNetworkSegmentResult, error) {
	var rv LookupWorkloadNetworkSegmentResult
	err := ctx.Invoke("azure-native:avs:getWorkloadNetworkSegment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadNetworkSegmentArgs struct {
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SegmentId         string `pulumi:"segmentId"`
}


type LookupWorkloadNetworkSegmentResult struct {
	ConnectedGateway  *string                                 `pulumi:"connectedGateway"`
	DisplayName       *string                                 `pulumi:"displayName"`
	Id                string                                  `pulumi:"id"`
	Name              string                                  `pulumi:"name"`
	PortVif           []WorkloadNetworkSegmentPortVifResponse `pulumi:"portVif"`
	ProvisioningState string                                  `pulumi:"provisioningState"`
	Revision          *float64                                `pulumi:"revision"`
	Status            string                                  `pulumi:"status"`
	Subnet            *WorkloadNetworkSegmentSubnetResponse   `pulumi:"subnet"`
	Type              string                                  `pulumi:"type"`
}
