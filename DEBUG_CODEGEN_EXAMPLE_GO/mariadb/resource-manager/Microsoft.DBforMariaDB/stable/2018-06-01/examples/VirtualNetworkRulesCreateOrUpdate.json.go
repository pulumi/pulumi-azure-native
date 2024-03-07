package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformariadb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformariadb.NewVirtualNetworkRule(ctx, "virtualNetworkRule", &dbformariadb.VirtualNetworkRuleArgs{
			IgnoreMissingVnetServiceEndpoint: pulumi.Bool(false),
			ResourceGroupName:                pulumi.String("TestGroup"),
			ServerName:                       pulumi.String("vnet-test-svr"),
			VirtualNetworkRuleName:           pulumi.String("vnet-firewall-rule"),
			VirtualNetworkSubnetId:           pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
