package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewLoadBalancerBackendAddressPool(ctx, "loadBalancerBackendAddressPool", &network.LoadBalancerBackendAddressPoolArgs{
			BackendAddressPoolName: pulumi.String("backend"),
			LoadBalancerBackendAddresses: network.LoadBalancerBackendAddressArray{
				&network.LoadBalancerBackendAddressArgs{
					IpAddress: pulumi.String("10.0.0.4"),
					Name:      pulumi.String("address1"),
					VirtualNetwork: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
					},
				},
				&network.LoadBalancerBackendAddressArgs{
					IpAddress: pulumi.String("10.0.0.5"),
					Name:      pulumi.String("address2"),
					VirtualNetwork: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
					},
				},
			},
			LoadBalancerName:  pulumi.String("lb"),
			ResourceGroupName: pulumi.String("testrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
