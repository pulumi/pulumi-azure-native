


package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIpFirewallRule(ctx *pulumi.Context, args *LookupIpFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupIpFirewallRuleResult, error) {
	var rv LookupIpFirewallRuleResult
	err := ctx.Invoke("azure-native:synapse/v20190601preview:getIpFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpFirewallRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupIpFirewallRuleResult struct {
	EndIpAddress      *string `pulumi:"endIpAddress"`
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	StartIpAddress    *string `pulumi:"startIpAddress"`
	Type              string  `pulumi:"type"`
}
