package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicenetworking/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicenetworking.NewAssociationsInterface(ctx, "associationsInterface", &servicenetworking.AssociationsInterfaceArgs{
			AssociationName:   pulumi.String("as1"),
			AssociationType:   pulumi.String("subnets"),
			Location:          pulumi.String("NorthCentralUS"),
			ResourceGroupName: pulumi.String("rg1"),
			Subnet: &servicenetworking.AssociationSubnetArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-tc/subnets/tc-subnet"),
			},
			TrafficControllerName: pulumi.String("tc1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
