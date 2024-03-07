package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewWorkloadNetworkVMGroup(ctx, "workloadNetworkVMGroup", &avs.WorkloadNetworkVMGroupArgs{
			DisplayName: pulumi.String("vmGroup1"),
			Members: pulumi.StringArray{
				pulumi.String("564d43da-fefc-2a3b-1d92-42855622fa50"),
			},
			PrivateCloudName:  pulumi.String("cloud1"),
			ResourceGroupName: pulumi.String("group1"),
			Revision:          pulumi.Float64(1),
			VmGroupId:         pulumi.String("vmGroup1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
