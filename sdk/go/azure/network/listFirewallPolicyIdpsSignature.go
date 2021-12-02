


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListFirewallPolicyIdpsSignature(ctx *pulumi.Context, args *ListFirewallPolicyIdpsSignatureArgs, opts ...pulumi.InvokeOption) (*ListFirewallPolicyIdpsSignatureResult, error) {
	var rv ListFirewallPolicyIdpsSignatureResult
	err := ctx.Invoke("azure-native:network:listFirewallPolicyIdpsSignature", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFirewallPolicyIdpsSignatureArgs struct {
	Filters            []FilterItems `pulumi:"filters"`
	FirewallPolicyName string        `pulumi:"firewallPolicyName"`
	OrderBy            *OrderBy      `pulumi:"orderBy"`
	ResourceGroupName  string        `pulumi:"resourceGroupName"`
	ResultsPerPage     *int          `pulumi:"resultsPerPage"`
	Search             *string       `pulumi:"search"`
	Skip               *int          `pulumi:"skip"`
}


type ListFirewallPolicyIdpsSignatureResult struct {
	MatchingRecordsCount *float64                    `pulumi:"matchingRecordsCount"`
	Signatures           []SingleQueryResultResponse `pulumi:"signatures"`
}
