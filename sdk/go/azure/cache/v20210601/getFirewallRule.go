


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallRule(ctx *pulumi.Context, args *LookupFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupFirewallRuleResult, error) {
	var rv LookupFirewallRuleResult
	err := ctx.Invoke("azure-native:cache/v20210601:getFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallRuleArgs struct {
	CacheName         string `pulumi:"cacheName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupFirewallRuleResult struct {
	EndIP   string `pulumi:"endIP"`
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	StartIP string `pulumi:"startIP"`
	Type    string `pulumi:"type"`
}
