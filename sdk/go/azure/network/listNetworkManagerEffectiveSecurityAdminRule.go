


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNetworkManagerEffectiveSecurityAdminRule(ctx *pulumi.Context, args *ListNetworkManagerEffectiveSecurityAdminRuleArgs, opts ...pulumi.InvokeOption) (*ListNetworkManagerEffectiveSecurityAdminRuleResult, error) {
	var rv ListNetworkManagerEffectiveSecurityAdminRuleResult
	err := ctx.Invoke("azure-native:network:listNetworkManagerEffectiveSecurityAdminRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNetworkManagerEffectiveSecurityAdminRuleArgs struct {
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type ListNetworkManagerEffectiveSecurityAdminRuleResult struct {
	SkipToken *string       `pulumi:"skipToken"`
	Value     []interface{} `pulumi:"value"`
}
