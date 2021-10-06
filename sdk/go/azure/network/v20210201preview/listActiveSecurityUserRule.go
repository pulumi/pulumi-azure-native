


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActiveSecurityUserRule(ctx *pulumi.Context, args *ListActiveSecurityUserRuleArgs, opts ...pulumi.InvokeOption) (*ListActiveSecurityUserRuleResult, error) {
	var rv ListActiveSecurityUserRuleResult
	err := ctx.Invoke("azure-native:network/v20210201preview:listActiveSecurityUserRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveSecurityUserRuleArgs struct {
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListActiveSecurityUserRuleResult struct {
	SkipToken *string       `pulumi:"skipToken"`
	Value     []interface{} `pulumi:"value"`
}
