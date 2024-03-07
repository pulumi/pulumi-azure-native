package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewWorkloadNetworkPortMirroring(ctx, "workloadNetworkPortMirroring", &avs.WorkloadNetworkPortMirroringArgs{
			Destination:       pulumi.String("vmGroup2"),
			Direction:         pulumi.String("BIDIRECTIONAL"),
			DisplayName:       pulumi.String("portMirroring1"),
			PortMirroringId:   pulumi.String("portMirroring1"),
			PrivateCloudName:  pulumi.String("cloud1"),
			ResourceGroupName: pulumi.String("group1"),
			Revision:          pulumi.Float64(1),
			Source:            pulumi.String("vmGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
