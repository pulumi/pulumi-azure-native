package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNetworkVirtualAppliance(ctx, "networkVirtualAppliance", &network.NetworkVirtualApplianceArgs{
			Delegation: &network.DelegationPropertiesArgs{
				ServiceName: pulumi.String("PaloAltoNetworks.Cloudngfw/firewalls"),
			},
			Location:                    pulumi.String("West US"),
			NetworkVirtualApplianceName: pulumi.String("nva"),
			ResourceGroupName:           pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			VirtualHub: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
