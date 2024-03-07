package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPrivateEndpoint(ctx, "privateEndpoint", &network.PrivateEndpointArgs{
			ApplicationSecurityGroups: network.ApplicationSecurityGroupTypeArray{
				&network.ApplicationSecurityGroupTypeArgs{
					Id: pulumi.String("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/applicationSecurityGroup/asg1"),
				},
			},
			Location:            pulumi.String("eastus2euap"),
			PrivateEndpointName: pulumi.String("testPe"),
			PrivateLinkServiceConnections: network.PrivateLinkServiceConnectionArray{
				&network.PrivateLinkServiceConnectionArgs{
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
