


package v20210801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOutboundFirewallRule(ctx *pulumi.Context, args *LookupOutboundFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupOutboundFirewallRuleResult, error) {
	var rv LookupOutboundFirewallRuleResult
	err := ctx.Invoke("azure-native:sql/v20210801preview:getOutboundFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutboundFirewallRuleArgs struct {
	OutboundRuleFqdn  string `pulumi:"outboundRuleFqdn"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupOutboundFirewallRuleResult struct {
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ProvisioningState string `pulumi:"provisioningState"`
	Type              string `pulumi:"type"`
}
