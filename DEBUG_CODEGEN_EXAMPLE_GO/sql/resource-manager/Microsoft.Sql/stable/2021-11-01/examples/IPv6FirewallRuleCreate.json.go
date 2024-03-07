package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewIPv6FirewallRule(ctx, "iPv6FirewallRule", &sql.IPv6FirewallRuleArgs{
			EndIPv6Address:    pulumi.String("0000:0000:0000:0000:0000:ffff:0000:0003"),
			FirewallRuleName:  pulumi.String("firewallrulecrudtest-5370"),
			ResourceGroupName: pulumi.String("firewallrulecrudtest-12"),
			ServerName:        pulumi.String("firewallrulecrudtest-6285"),
			StartIPv6Address:  pulumi.String("0000:0000:0000:0000:0000:ffff:0000:0003"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
