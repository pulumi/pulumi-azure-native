package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualWan(ctx, "virtualWan", &network.VirtualWanArgs{
			DisableVpnEncryption: pulumi.Bool(false),
			Location:             pulumi.String("West US"),
			ResourceGroupName:    pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			Type:           pulumi.String("Basic"),
			VirtualWANName: pulumi.String("wan1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
