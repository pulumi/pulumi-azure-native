package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualNetwork(ctx, "virtualNetwork", &network.VirtualNetworkArgs{
			AddressSpace: &network.AddressSpaceArgs{
				AddressPrefixes: pulumi.StringArray{
					pulumi.String("10.0.0.0/16"),
				},
			},
			Location:          pulumi.String("eastus2euap"),
			ResourceGroupName: pulumi.String("vnetTest"),
			Subnets: []network.SubnetTypeArgs{
				{
					AddressPrefix: pulumi.String("10.0.0.0/16"),
					Name:          pulumi.String("test-1"),
					ServiceEndpointPolicies: network.ServiceEndpointPolicyTypeArray{
						{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/vnetTest/providers/Microsoft.Network/serviceEndpointPolicies/ServiceEndpointPolicy1"),
						},
					},
					ServiceEndpoints: network.ServiceEndpointPropertiesFormatArray{
						{
							Service: pulumi.String("Microsoft.Storage"),
						},
					},
				},
			},
			VirtualNetworkName: pulumi.String("vnet1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
