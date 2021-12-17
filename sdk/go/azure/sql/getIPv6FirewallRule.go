


package sql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIPv6FirewallRule(ctx *pulumi.Context, args *LookupIPv6FirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupIPv6FirewallRuleResult, error) {
	var rv LookupIPv6FirewallRuleResult
	err := ctx.Invoke("azure-native:sql:getIPv6FirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIPv6FirewallRuleArgs struct {
	FirewallRuleName  string `pulumi:"firewallRuleName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupIPv6FirewallRuleResult struct {
	EndIPv6Address   *string `pulumi:"endIPv6Address"`
	Id               string  `pulumi:"id"`
	Name             *string `pulumi:"name"`
	StartIPv6Address *string `pulumi:"startIPv6Address"`
	Type             string  `pulumi:"type"`
}
