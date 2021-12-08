


package alertsmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrometheusRuleGroup(ctx *pulumi.Context, args *LookupPrometheusRuleGroupArgs, opts ...pulumi.InvokeOption) (*LookupPrometheusRuleGroupResult, error) {
	var rv LookupPrometheusRuleGroupResult
	err := ctx.Invoke("azure-native:alertsmanagement:getPrometheusRuleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrometheusRuleGroupArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleGroupName     string `pulumi:"ruleGroupName"`
}


type LookupPrometheusRuleGroupResult struct {
	Description *string                  `pulumi:"description"`
	Id          string                   `pulumi:"id"`
	Interval    *string                  `pulumi:"interval"`
	Location    string                   `pulumi:"location"`
	Name        string                   `pulumi:"name"`
	Rules       []PrometheusRuleResponse `pulumi:"rules"`
	Scopes      []string                 `pulumi:"scopes"`
	SystemData  SystemDataResponse       `pulumi:"systemData"`
	Tags        map[string]string        `pulumi:"tags"`
	Type        string                   `pulumi:"type"`
}
