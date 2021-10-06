


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActiveSecurityAdminRule(ctx *pulumi.Context, args *ListActiveSecurityAdminRuleArgs, opts ...pulumi.InvokeOption) (*ListActiveSecurityAdminRuleResult, error) {
	var rv ListActiveSecurityAdminRuleResult
	err := ctx.Invoke("azure-native:network:listActiveSecurityAdminRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveSecurityAdminRuleArgs struct {
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListActiveSecurityAdminRuleResult struct {
	SkipToken *string       `pulumi:"skipToken"`
	Value     []interface{} `pulumi:"value"`
}
