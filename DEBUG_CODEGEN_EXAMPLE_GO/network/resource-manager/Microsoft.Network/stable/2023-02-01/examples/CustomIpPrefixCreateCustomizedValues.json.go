package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewCustomIPPrefix(ctx, "customIPPrefix", &network.CustomIPPrefixArgs{
			Cidr:               pulumi.String("0.0.0.0/24"),
			CustomIpPrefixName: pulumi.String("test-customipprefix"),
			Location:           pulumi.String("westus"),
			ResourceGroupName:  pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
