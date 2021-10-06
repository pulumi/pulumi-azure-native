


package v20160401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRedisFirewallRule(ctx *pulumi.Context, args *LookupRedisFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupRedisFirewallRuleResult, error) {
	var rv LookupRedisFirewallRuleResult
	err := ctx.Invoke("azure-native:cache/v20160401:getRedisFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRedisFirewallRuleArgs struct {
	CacheName         string `pulumi:"cacheName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupRedisFirewallRuleResult struct {
	EndIP   string `pulumi:"endIP"`
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	StartIP string `pulumi:"startIP"`
	Type    string `pulumi:"type"`
}
