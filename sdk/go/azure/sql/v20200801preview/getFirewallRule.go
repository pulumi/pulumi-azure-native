


package v20200801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallRule(ctx *pulumi.Context, args *LookupFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupFirewallRuleResult, error) {
	var rv LookupFirewallRuleResult
	err := ctx.Invoke("azure-native:sql/v20200801preview:getFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallRuleArgs struct {
	FirewallRuleName  string `pulumi:"firewallRuleName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupFirewallRuleResult struct {
	EndIpAddress   *string `pulumi:"endIpAddress"`
	Id             string  `pulumi:"id"`
	Name           *string `pulumi:"name"`
	StartIpAddress *string `pulumi:"startIpAddress"`
	Type           string  `pulumi:"type"`
}
