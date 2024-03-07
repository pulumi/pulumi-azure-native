package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNetworkInterfaceTapConfiguration(ctx, "networkInterfaceTapConfiguration", &network.NetworkInterfaceTapConfigurationArgs{
			NetworkInterfaceName: pulumi.String("mynic"),
			ResourceGroupName:    pulumi.String("testrg"),
			TapConfigurationName: pulumi.String("tapconfiguration1"),
			VirtualNetworkTap: &network.VirtualNetworkTapTypeArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworkTaps/testvtap"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
