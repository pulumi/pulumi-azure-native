


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledAlertRule(ctx *pulumi.Context, args *LookupScheduledAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupScheduledAlertRuleResult, error) {
	var rv LookupScheduledAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20210301preview:getScheduledAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledAlertRuleArgs struct {
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	RuleId                              string `pulumi:"ruleId"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupScheduledAlertRuleResult struct {
	AlertDetailsOverride  *AlertDetailsOverrideResponse  `pulumi:"alertDetailsOverride"`
	AlertRuleTemplateName *string                        `pulumi:"alertRuleTemplateName"`
	CustomDetails         map[string]string              `pulumi:"customDetails"`
	Description           *string                        `pulumi:"description"`
	DisplayName           string                         `pulumi:"displayName"`
	Enabled               bool                           `pulumi:"enabled"`
	EntityMappings        []EntityMappingResponse        `pulumi:"entityMappings"`
	Etag                  *string                        `pulumi:"etag"`
	EventGroupingSettings *EventGroupingSettingsResponse `pulumi:"eventGroupingSettings"`
	Id                    string                         `pulumi:"id"`
	IncidentConfiguration *IncidentConfigurationResponse `pulumi:"incidentConfiguration"`
	Kind                  string                         `pulumi:"kind"`
	LastModifiedUtc       string                         `pulumi:"lastModifiedUtc"`
	Name                  string                         `pulumi:"name"`
	Query                 string                         `pulumi:"query"`
	QueryFrequency        string                         `pulumi:"queryFrequency"`
	QueryPeriod           string                         `pulumi:"queryPeriod"`
	Severity              string                         `pulumi:"severity"`
	SuppressionDuration   string                         `pulumi:"suppressionDuration"`
	SuppressionEnabled    bool                           `pulumi:"suppressionEnabled"`
	SystemData            SystemDataResponse             `pulumi:"systemData"`
	Tactics               []string                       `pulumi:"tactics"`
	TriggerOperator       string                         `pulumi:"triggerOperator"`
	TriggerThreshold      int                            `pulumi:"triggerThreshold"`
	Type                  string                         `pulumi:"type"`
}
