package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNspAccessRule(ctx, "nspAccessRule", &network.NspAccessRuleArgs{
			AccessRuleName: pulumi.String("accessRule1"),
			AddressPrefixes: pulumi.StringArray{
				pulumi.String("10.11.0.0/16"),
				pulumi.String("10.10.1.0/24"),
			},
			Direction:                    pulumi.String("Inbound"),
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
