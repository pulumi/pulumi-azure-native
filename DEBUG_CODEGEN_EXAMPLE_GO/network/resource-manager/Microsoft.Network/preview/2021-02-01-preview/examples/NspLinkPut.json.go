package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNspLink(ctx, "nspLink", &network.NspLinkArgs{
			AutoApprovedRemotePerimeterResourceId: pulumi.String("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp2"),
			LinkName:                              pulumi.String("link1"),
			NetworkSecurityPerimeterName:          pulumi.String("nsp1"),
			ResourceGroupName:                     pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
