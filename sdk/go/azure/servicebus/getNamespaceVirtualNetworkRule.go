


package servicebus

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespaceVirtualNetworkRule(ctx *pulumi.Context, args *LookupNamespaceVirtualNetworkRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceVirtualNetworkRuleResult, error) {
	var rv LookupNamespaceVirtualNetworkRuleResult
	err := ctx.Invoke("azure-native:servicebus:getNamespaceVirtualNetworkRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceVirtualNetworkRuleArgs struct {
	NamespaceName          string `pulumi:"namespaceName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	VirtualNetworkRuleName string `pulumi:"virtualNetworkRuleName"`
}


type LookupNamespaceVirtualNetworkRuleResult struct {
	Id                     string  `pulumi:"id"`
	Name                   string  `pulumi:"name"`
	Type                   string  `pulumi:"type"`
	VirtualNetworkSubnetId *string `pulumi:"virtualNetworkSubnetId"`
}
