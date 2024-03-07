package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewAddon(ctx, "addon", &avs.AddonArgs{
			AddonName:        pulumi.String("arc"),
			PrivateCloudName: pulumi.String("cloud1"),
			Properties: avs.AddonArcProperties{
				AddonType: "Arc",
				VCenter:   "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg_test/providers/Microsoft.ConnectedVMwarevSphere/VCenters/test-vcenter",
			},
			ResourceGroupName: pulumi.String("group1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
