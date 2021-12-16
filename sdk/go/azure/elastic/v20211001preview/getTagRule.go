


package v20211001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagRule(ctx *pulumi.Context, args *LookupTagRuleArgs, opts ...pulumi.InvokeOption) (*LookupTagRuleResult, error) {
	var rv LookupTagRuleResult
	err := ctx.Invoke("azure-native:elastic/v20211001preview:getTagRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagRuleArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleSetName       string `pulumi:"ruleSetName"`
}


type LookupTagRuleResult struct {
	Id         string                               `pulumi:"id"`
	Name       string                               `pulumi:"name"`
	Properties MonitoringTagRulesPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                   `pulumi:"systemData"`
	Type       string                               `pulumi:"type"`
}
