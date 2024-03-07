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
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("rg1"),
			Subnets: []network.SubnetTypeArgs{
				{
					AddressPrefixes: pulumi.StringArray{
						pulumi.String("10.0.0.0/28"),
						pulumi.String("10.0.1.0/28"),
					},
					Name: pulumi.String("test-2"),
				},
			},
			VirtualNetworkName: pulumi.String("test-vnet"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
