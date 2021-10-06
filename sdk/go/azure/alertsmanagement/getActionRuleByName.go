


package alertsmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupActionRuleByName(ctx *pulumi.Context, args *LookupActionRuleByNameArgs, opts ...pulumi.InvokeOption) (*LookupActionRuleByNameResult, error) {
	var rv LookupActionRuleByNameResult
	err := ctx.Invoke("azure-native:alertsmanagement:getActionRuleByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupActionRuleByNameArgs struct {
	ActionRuleName    string `pulumi:"actionRuleName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupActionRuleByNameResult struct {
	Id         string            `pulumi:"id"`
	Location   string            `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}
