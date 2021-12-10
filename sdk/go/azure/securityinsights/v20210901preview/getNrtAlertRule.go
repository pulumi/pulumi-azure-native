


package v20210901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNrtAlertRule(ctx *pulumi.Context, args *LookupNrtAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupNrtAlertRuleResult, error) {
	var rv LookupNrtAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getNrtAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNrtAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupNrtAlertRuleResult struct {
	AlertDetailsOverride  *AlertDetailsOverrideResponse  `pulumi:"alertDetailsOverride"`
	AlertRuleTemplateName *string                        `pulumi:"alertRuleTemplateName"`
	CustomDetails         map[string]string              `pulumi:"customDetails"`
	Description           *string                        `pulumi:"description"`
	DisplayName           string                         `pulumi:"displayName"`
	Enabled               bool                           `pulumi:"enabled"`
	EntityMappings        []EntityMappingResponse        `pulumi:"entityMappings"`
	Etag                  *string                        `pulumi:"etag"`
	Id                    string                         `pulumi:"id"`
	IncidentConfiguration *IncidentConfigurationResponse `pulumi:"incidentConfiguration"`
	Kind                  string                         `pulumi:"kind"`
	LastModifiedUtc       string                         `pulumi:"lastModifiedUtc"`
	Name                  string                         `pulumi:"name"`
	Query                 *string                        `pulumi:"query"`
	Severity              *string                        `pulumi:"severity"`
	SuppressionDuration   string                         `pulumi:"suppressionDuration"`
	SuppressionEnabled    bool                           `pulumi:"suppressionEnabled"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	Tactics               []string                       `pulumi:"tactics"`
	TemplateVersion       *string                        `pulumi:"templateVersion"`
	Type                  string                         `pulumi:"type"`
}
