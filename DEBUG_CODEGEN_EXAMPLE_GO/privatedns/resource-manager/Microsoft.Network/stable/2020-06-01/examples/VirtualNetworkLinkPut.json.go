package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualNetworkLink(ctx, "virtualNetworkLink", &network.VirtualNetworkLinkArgs{
			Location:            pulumi.String("Global"),
			PrivateZoneName:     pulumi.String("privatezone1.com"),
			RegistrationEnabled: pulumi.Bool(false),
			ResourceGroupName:   pulumi.String("resourceGroup1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			VirtualNetwork: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName"),
			},
			VirtualNetworkLinkName: pulumi.String("virtualNetworkLink1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
