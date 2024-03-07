package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datalakeanalytics/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datalakeanalytics.NewFirewallRule(ctx, "firewallRule", &datalakeanalytics.FirewallRuleArgs{
			AccountName:       pulumi.String("contosoadla"),
			EndIpAddress:      pulumi.String("2.2.2.2"),
			FirewallRuleName:  pulumi.String("test_rule"),
			ResourceGroupName: pulumi.String("contosorg"),
			StartIpAddress:    pulumi.String("1.1.1.1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
