package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cache.NewFirewallRule(ctx, "firewallRule", &cache.FirewallRuleArgs{
			CacheName:         pulumi.String("cache1"),
			EndIP:             pulumi.String("192.168.1.4"),
			ResourceGroupName: pulumi.String("rg1"),
			RuleName:          pulumi.String("rule1"),
			StartIP:           pulumi.String("192.168.1.1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
