


package v20161101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkRule(ctx *pulumi.Context, args *LookupVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkRuleResult, error) {
	var rv LookupVirtualNetworkRuleResult
	err := ctx.Invoke("azure-native:datalakestore/v20161101:getVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkRuleArgs struct {
	AccountName            string `pulumi:"accountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}


type LookupVirtualNetworkRuleResult struct {
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	SubnetId string `pulumi:"subnetId"`
	Type     string `pulumi:"type"`
}
