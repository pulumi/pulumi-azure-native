package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewFirewallRule(ctx, "firewallRule", &dbforpostgresql.FirewallRuleArgs{
			EndIpAddress:      pulumi.String("255.255.255.255"),
			FirewallRuleName:  pulumi.String("rule1"),
			ResourceGroupName: pulumi.String("testrg"),
			ServerName:        pulumi.String("testserver"),
			StartIpAddress:    pulumi.String("0.0.0.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
