


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallPolicy(ctx *pulumi.Context, args *LookupFirewallPolicyArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyResult, error) {
	var rv LookupFirewallPolicyResult
	err := ctx.Invoke("azure-native:network/v20200301:getFirewallPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallPolicyArgs struct {
	Expand             *string `pulumi:"expand"`
	FirewallPolicyName string  `pulumi:"firewallPolicyName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupFirewallPolicyResult struct {
	BasePolicy        *SubResourceResponse  `pulumi:"basePolicy"`
	ChildPolicies     []SubResourceResponse `pulumi:"childPolicies"`
	Etag              string                `pulumi:"etag"`
	Firewalls         []SubResourceResponse `pulumi:"firewalls"`
	Id                *string               `pulumi:"id"`
	Location          *string               `pulumi:"location"`
	Name              string                `pulumi:"name"`
	ProvisioningState string                `pulumi:"provisioningState"`
	RuleGroups        []SubResourceResponse `pulumi:"ruleGroups"`
	Tags              map[string]string     `pulumi:"tags"`
	ThreatIntelMode   *string               `pulumi:"threatIntelMode"`
	Type              string                `pulumi:"type"`
}
