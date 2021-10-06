


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMicrosoftSecurityIncidentCreationAlertRule(ctx *pulumi.Context, args *LookupMicrosoftSecurityIncidentCreationAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupMicrosoftSecurityIncidentCreationAlertRuleResult, error) {
	var rv LookupMicrosoftSecurityIncidentCreationAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights:getMicrosoftSecurityIncidentCreationAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMicrosoftSecurityIncidentCreationAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMicrosoftSecurityIncidentCreationAlertRuleResult struct {
	AlertRuleTemplateName     *string  `pulumi:"alertRuleTemplateName"`
	Description               *string  `pulumi:"description"`
	DisplayName               string   `pulumi:"displayName"`
	DisplayNamesExcludeFilter []string `pulumi:"displayNamesExcludeFilter"`
	DisplayNamesFilter        []string `pulumi:"displayNamesFilter"`
	Enabled                   bool     `pulumi:"enabled"`
	Etag                      *string  `pulumi:"etag"`
	Id                        string   `pulumi:"id"`
	Kind                      string   `pulumi:"kind"`
	LastModifiedUtc           string   `pulumi:"lastModifiedUtc"`
	Name                      string   `pulumi:"name"`
	ProductFilter             string   `pulumi:"productFilter"`
	SeveritiesFilter          []string `pulumi:"severitiesFilter"`
	Type                      string   `pulumi:"type"`
}
