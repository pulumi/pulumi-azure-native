package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datalakestore/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datalakestore.NewVirtualNetworkRule(ctx, "virtualNetworkRule", &datalakestore.VirtualNetworkRuleArgs{
			AccountName:            pulumi.String("contosoadla"),
			ResourceGroupName:      pulumi.String("contosorg"),
			SubnetId:               pulumi.String("test_subnetId"),
			VirtualNetworkRuleName: pulumi.String("test_virtual_network_rules_name"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
