


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupAlertRule(ctx *pulumi.Context, args *LookupAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupAlertRuleResult, error) {
	var rv LookupAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20200101:getAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAlertRuleResult struct {
	Etag *string `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Kind string  `pulumi:"kind"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}
