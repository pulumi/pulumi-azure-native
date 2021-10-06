


package securityinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutomationRule(ctx *pulumi.Context, args *LookupAutomationRuleArgs, opts ...pulumi.InvokeOption) (*LookupAutomationRuleResult, error) {
	var rv LookupAutomationRuleResult
	err := ctx.Invoke("azure-native:securityinsights:getAutomationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutomationRuleArgs struct {
	AutomationRuleId                    string `pulumi:"automationRuleId"`
	OperationalInsightsResourceProvider string `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	WorkspaceName                       string `pulumi:"workspaceName"`
}


type LookupAutomationRuleResult struct {
	Actions             []interface{}                         `pulumi:"actions"`
	CreatedBy           ClientInfoResponse                    `pulumi:"createdBy"`
	CreatedTimeUtc      string                                `pulumi:"createdTimeUtc"`
	DisplayName         string                                `pulumi:"displayName"`
	Etag                *string                               `pulumi:"etag"`
	Id                  string                                `pulumi:"id"`
	LastModifiedBy      ClientInfoResponse                    `pulumi:"lastModifiedBy"`
	LastModifiedTimeUtc string                                `pulumi:"lastModifiedTimeUtc"`
	Name                string                                `pulumi:"name"`
	Order               int                                   `pulumi:"order"`
	TriggeringLogic     AutomationRuleTriggeringLogicResponse `pulumi:"triggeringLogic"`
	Type                string                                `pulumi:"type"`
}
