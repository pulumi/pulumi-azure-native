


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context, args *LookupFirewallPolicyRuleCollectionGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyRuleCollectionGroupResult, error) {
	var rv LookupFirewallPolicyRuleCollectionGroupResult
	err := ctx.Invoke("azure-native:network/v20200801:getFirewallPolicyRuleCollectionGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallPolicyRuleCollectionGroupArgs struct {
	FirewallPolicyName      string `pulumi:"firewallPolicyName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	RuleCollectionGroupName string `pulumi:"ruleCollectionGroupName"`
}


type LookupFirewallPolicyRuleCollectionGroupResult struct {
	Etag              string        `pulumi:"etag"`
	Id                *string       `pulumi:"id"`
	Name              *string       `pulumi:"name"`
	Priority          *int          `pulumi:"priority"`
	ProvisioningState string        `pulumi:"provisioningState"`
	RuleCollections   []interface{} `pulumi:"ruleCollections"`
	Type              string        `pulumi:"type"`
}
