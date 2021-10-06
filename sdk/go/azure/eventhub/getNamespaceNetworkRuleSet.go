


package eventhub

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespaceNetworkRuleSet(ctx *pulumi.Context, args *LookupNamespaceNetworkRuleSetArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceNetworkRuleSetResult, error) {
	var rv LookupNamespaceNetworkRuleSetResult
	err := ctx.Invoke("azure-native:eventhub:getNamespaceNetworkRuleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceNetworkRuleSetArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNamespaceNetworkRuleSetResult struct {
	DefaultAction       *string                                `pulumi:"defaultAction"`
	Id                  string                                 `pulumi:"id"`
	IpRules             []NWRuleSetIpRulesResponse             `pulumi:"ipRules"`
	Name                string                                 `pulumi:"name"`
	Type                string                                 `pulumi:"type"`
	VirtualNetworkRules []NWRuleSetVirtualNetworkRulesResponse `pulumi:"virtualNetworkRules"`
}
