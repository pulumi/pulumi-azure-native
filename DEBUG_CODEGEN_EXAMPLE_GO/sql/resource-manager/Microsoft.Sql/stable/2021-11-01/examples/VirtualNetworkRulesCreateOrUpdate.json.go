package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewVirtualNetworkRule(ctx, "virtualNetworkRule", &sql.VirtualNetworkRuleArgs{
			IgnoreMissingVnetServiceEndpoint: pulumi.Bool(false),
			ResourceGroupName:                pulumi.String("Default"),
			ServerName:                       pulumi.String("vnet-test-svr"),
			VirtualNetworkRuleName:           pulumi.String("vnet-firewall-rule"),
			VirtualNetworkSubnetId:           pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
