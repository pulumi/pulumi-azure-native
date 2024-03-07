package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewFirewallRule(ctx, "firewallRule", &sql.FirewallRuleArgs{
			EndIpAddress:      pulumi.String("0.0.0.1"),
			FirewallRuleName:  pulumi.String("firewallrulecrudtest-3927"),
			ResourceGroupName: pulumi.String("firewallrulecrudtest-12"),
			ServerName:        pulumi.String("firewallrulecrudtest-6285"),
			StartIpAddress:    pulumi.String("0.0.0.1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
