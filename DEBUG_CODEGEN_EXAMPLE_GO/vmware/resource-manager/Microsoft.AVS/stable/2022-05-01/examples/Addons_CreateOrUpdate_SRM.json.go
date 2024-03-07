package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewAddon(ctx, "addon", &avs.AddonArgs{
			AddonName:        pulumi.String("srm"),
			PrivateCloudName: pulumi.String("cloud1"),
			Properties: avs.AddonSrmProperties{
				AddonType:  "SRM",
				LicenseKey: "41915178-A8FF-4A4D-B683-6D735AF5E3F5",
			},
			ResourceGroupName: pulumi.String("group1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
