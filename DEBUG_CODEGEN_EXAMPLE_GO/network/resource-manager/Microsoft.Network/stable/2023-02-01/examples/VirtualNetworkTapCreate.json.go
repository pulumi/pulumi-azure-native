package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualNetworkTap(ctx, "virtualNetworkTap", &network.VirtualNetworkTapArgs{
			DestinationNetworkInterfaceIPConfiguration: &network.NetworkInterfaceIPConfigurationArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/ipconfig1"),
			},
			Location:          pulumi.String("centraluseuap"),
			ResourceGroupName: pulumi.String("rg1"),
			TapName:           pulumi.String("test-vtap"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
