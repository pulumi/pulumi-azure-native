package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPublicIPPrefix(ctx, "publicIPPrefix", &network.PublicIPPrefixArgs{
			Location:               pulumi.String("westus"),
			PrefixLength:           pulumi.Int(30),
			PublicIPAddressVersion: pulumi.String("IPv4"),
			PublicIpPrefixName:     pulumi.String("test-ipprefix"),
			ResourceGroupName:      pulumi.String("rg1"),
			Sku: &network.PublicIPPrefixSkuArgs{
				Name: pulumi.String("Standard"),
				Tier: pulumi.String("Regional"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
