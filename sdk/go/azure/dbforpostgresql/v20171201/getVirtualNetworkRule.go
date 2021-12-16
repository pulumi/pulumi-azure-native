


package v20171201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkRule(ctx *pulumi.Context, args *LookupVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkRuleResult, error) {
	var rv LookupVirtualNetworkRuleResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20171201:getVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkRuleArgs struct {
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ServerName             string `pulumi:"serverName"`
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}


type LookupVirtualNetworkRuleResult struct {
	Id                               string `pulumi:"id"`
	IgnoreMissingVnetServiceEndpoint *bool  `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Name                             string `pulumi:"name"`
	State                            string `pulumi:"state"`
	Type                             string `pulumi:"type"`
	VirtualNetworkSubnetId           string `pulumi:"virtualNetworkSubnetId"`
}
