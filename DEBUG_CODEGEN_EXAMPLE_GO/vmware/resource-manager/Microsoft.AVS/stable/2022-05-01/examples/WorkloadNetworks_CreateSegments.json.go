package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewWorkloadNetworkSegment(ctx, "workloadNetworkSegment", &avs.WorkloadNetworkSegmentArgs{
			ConnectedGateway:  pulumi.String("/infra/tier-1s/gateway"),
			DisplayName:       pulumi.String("segment1"),
			PrivateCloudName:  pulumi.String("cloud1"),
			ResourceGroupName: pulumi.String("group1"),
			Revision:          pulumi.Float64(1),
			SegmentId:         pulumi.String("segment1"),
			Subnet: &avs.WorkloadNetworkSegmentSubnetArgs{
				DhcpRanges: pulumi.StringArray{
					pulumi.String("40.20.0.0-40.20.0.1"),
				},
				GatewayAddress: pulumi.String("40.20.20.20/16"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
