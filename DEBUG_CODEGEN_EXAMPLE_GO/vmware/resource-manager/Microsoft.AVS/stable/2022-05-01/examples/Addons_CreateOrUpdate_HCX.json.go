package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/avs/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := avs.NewAddon(ctx, "addon", &avs.AddonArgs{
			AddonName:        pulumi.String("hcx"),
			PrivateCloudName: pulumi.String("cloud1"),
			Properties: avs.AddonHcxProperties{
				AddonType: "HCX",
				Offer:     "VMware MaaS Cloud Provider (Enterprise)",
			},
			ResourceGroupName: pulumi.String("group1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
