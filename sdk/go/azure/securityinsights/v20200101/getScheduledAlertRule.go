


package v20200101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledAlertRule(ctx *pulumi.Context, args *LookupScheduledAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupScheduledAlertRuleResult, error) {
	var rv LookupScheduledAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20200101:getScheduledAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupScheduledAlertRuleResult struct {
	AlertRuleTemplateName *string  `pulumi:"alertRuleTemplateName"`
	Description           *string  `pulumi:"description"`
	DisplayName           string   `pulumi:"displayName"`
	Enabled               bool     `pulumi:"enabled"`
	Etag                  *string  `pulumi:"etag"`
	Id                    string   `pulumi:"id"`
	Kind                  string   `pulumi:"kind"`
	LastModifiedUtc       string   `pulumi:"lastModifiedUtc"`
	Name                  string   `pulumi:"name"`
	Query                 string   `pulumi:"query"`
	QueryFrequency        string   `pulumi:"queryFrequency"`
	QueryPeriod           string   `pulumi:"queryPeriod"`
	Severity              string   `pulumi:"severity"`
	SuppressionDuration   string   `pulumi:"suppressionDuration"`
	SuppressionEnabled    bool     `pulumi:"suppressionEnabled"`
	Tactics               []string `pulumi:"tactics"`
	TriggerOperator       string   `pulumi:"triggerOperator"`
	TriggerThreshold      int      `pulumi:"triggerThreshold"`
	Type                  string   `pulumi:"type"`
}
