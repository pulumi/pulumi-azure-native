package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPrivateEndpoint(ctx, "privateEndpoint", &network.PrivateEndpointArgs{
			CustomNetworkInterfaceName: pulumi.String("testPeNic"),
			IpConfigurations: []network.PrivateEndpointIPConfigurationArgs{
				{
					GroupId:          pulumi.String("file"),
					MemberName:       pulumi.String("file"),
					Name:             pulumi.String("pestaticconfig"),
					PrivateIPAddress: pulumi.String("192.168.0.6"),
				},
			},
			Location:            pulumi.String("eastus2euap"),
			PrivateEndpointName: pulumi.String("testPe"),
			PrivateLinkServiceConnections: []network.PrivateLinkServiceConnectionArgs{
				{
					GroupIds: pulumi.StringArray{
						pulumi.String("groupIdFromResource"),
					},
					PrivateLinkServiceId: pulumi.String("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
					RequestMessage:       pulumi.String("Please approve my connection."),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			Subnet: &network.SubnetTypeArgs{
				Id: pulumi.String("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
