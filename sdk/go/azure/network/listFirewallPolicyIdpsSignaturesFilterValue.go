


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFirewallPolicyIdpsSignaturesFilterValue(ctx *pulumi.Context, args *ListFirewallPolicyIdpsSignaturesFilterValueArgs, opts ...pulumi.InvokeOption) (*ListFirewallPolicyIdpsSignaturesFilterValueResult, error) {
	var rv ListFirewallPolicyIdpsSignaturesFilterValueResult
	err := ctx.Invoke("azure-native:network:listFirewallPolicyIdpsSignaturesFilterValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFirewallPolicyIdpsSignaturesFilterValueArgs struct {
	FilterName         *string `pulumi:"filterName"`
	FirewallPolicyName string  `pulumi:"firewallPolicyName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type ListFirewallPolicyIdpsSignaturesFilterValueResult struct {
	FilterValues []string `pulumi:"filterValues"`
}
