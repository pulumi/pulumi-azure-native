package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewIpAllocation(ctx, "ipAllocation", &network.IpAllocationArgs{
			AllocationTags: pulumi.StringMap{
				"VNetID": pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
			},
			IpAllocationName:  pulumi.String("test-ipallocation"),
			Location:          pulumi.String("centraluseuap"),
			Prefix:            pulumi.String("3.2.5.0/24"),
			ResourceGroupName: pulumi.String("rg1"),
			Type:              pulumi.String("Hypernet"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
