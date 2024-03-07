package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewPublicIPAddress(ctx, "publicIPAddress", &network.PublicIPAddressArgs{
			IdleTimeoutInMinutes:     pulumi.Int(10),
			Location:                 pulumi.String("eastus"),
			PublicIPAddressVersion:   pulumi.String("IPv4"),
			PublicIPAllocationMethod: pulumi.String("Static"),
			PublicIpAddressName:      pulumi.String("test-ip"),
			ResourceGroupName:        pulumi.String("rg1"),
			Sku: &network.PublicIPAddressSkuArgs{
				Name: pulumi.String("Standard"),
				Tier: pulumi.String("Global"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
