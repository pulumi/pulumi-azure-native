package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNetworkSecurityPerimeter(ctx, "networkSecurityPerimeter", &network.NetworkSecurityPerimeterArgs{
			Description:                  pulumi.String("Description of TestNetworkSecurityPerimeter"),
			DisplayName:                  pulumi.String("TestNetworkSecurityPerimeter"),
			NetworkSecurityPerimeterName: pulumi.String("nsp1"),
			ResourceGroupName:            pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
