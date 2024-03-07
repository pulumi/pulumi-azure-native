package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewOutboundFirewallRule(ctx, "outboundFirewallRule", &sql.OutboundFirewallRuleArgs{
			OutboundRuleFqdn:  pulumi.String("server.database.windows.net"),
			ResourceGroupName: pulumi.String("sqlcrudtest-7398"),
			ServerName:        pulumi.String("sqlcrudtest-4645"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
