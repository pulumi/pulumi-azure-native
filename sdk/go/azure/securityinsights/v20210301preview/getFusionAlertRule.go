


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFusionAlertRule(ctx *pulumi.Context, args *LookupFusionAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupFusionAlertRuleResult, error) {
	var rv LookupFusionAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getFusionAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFusionAlertRuleArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	RuleId                              string `pulumi:"ruleId"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupFusionAlertRuleResult struct {
	AlertRuleTemplateName string             `pulumi:"alertRuleTemplateName"`
	Description           string             `pulumi:"description"`
	DisplayName           string             `pulumi:"displayName"`
	Enabled               bool               `pulumi:"enabled"`
	Etag                  *string            `pulumi:"etag"`
	Id                    string             `pulumi:"id"`
	Kind                  string             `pulumi:"kind"`
	LastModifiedUtc       string             `pulumi:"lastModifiedUtc"`
	Name                  string             `pulumi:"name"`
	Severity              string             `pulumi:"severity"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	Tactics               []string           `pulumi:"tactics"`
	Type                  string             `pulumi:"type"`
}
