


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRulesEngine(ctx *pulumi.Context, args *LookupRulesEngineArgs, opts ...pulumi.InvokeOption) (*LookupRulesEngineResult, error) {
	var rv LookupRulesEngineResult
	err := ctx.Invoke("azure-native:network:getRulesEngine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRulesEngineArgs struct {
	FrontDoorName     string `pulumi:"frontDoorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RulesEngineName   string `pulumi:"rulesEngineName"`
}


type LookupRulesEngineResult struct {
	Id            string                    `pulumi:"id"`
	Name          string                    `pulumi:"name"`
	ResourceState string                    `pulumi:"resourceState"`
	Rules         []RulesEngineRuleResponse `pulumi:"rules"`
	Type          string                    `pulumi:"type"`
}
