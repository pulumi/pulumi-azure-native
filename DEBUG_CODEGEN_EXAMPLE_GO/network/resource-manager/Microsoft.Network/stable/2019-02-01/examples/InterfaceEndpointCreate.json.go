package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewInterfaceEndpoint(ctx, "interfaceEndpoint", &network.InterfaceEndpointArgs{
			EndpointService: &network.EndpointServiceArgs{
				Id: pulumi.String("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Provider/resourceType/resourceName"),
			},
			Fqdn:                  pulumi.String("uniqueIdentifier.fqdn.windows.net"),
			InterfaceEndpointName: pulumi.String("testIe"),
			ResourceGroupName:     pulumi.String("rg1"),
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
