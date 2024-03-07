package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewSubnet(ctx, "subnet", &network.SubnetArgs{
			AddressPrefix:      pulumi.String("10.0.0.0/16"),
			ResourceGroupName:  pulumi.String("subnet-test"),
			SubnetName:         pulumi.String("subnet1"),
			VirtualNetworkName: pulumi.String("vnetname"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
