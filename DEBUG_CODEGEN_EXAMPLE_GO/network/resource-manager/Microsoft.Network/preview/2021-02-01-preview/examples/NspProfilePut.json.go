package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNspProfile(ctx, "nspProfile", &network.NspProfileArgs{
			NetworkSecurityPerimeterName: pulumi.String("nsp1"),
			ProfileName:                  pulumi.String("profile1"),
			ResourceGroupName:            pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
