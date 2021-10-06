


package v20190101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlertsSuppressionRule(ctx *pulumi.Context, args *LookupAlertsSuppressionRuleArgs, opts ...pulumi.InvokeOption) (*LookupAlertsSuppressionRuleResult, error) {
	var rv LookupAlertsSuppressionRuleResult
	err := ctx.Invoke("azure-native:security/v20190101preview:getAlertsSuppressionRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlertsSuppressionRuleArgs struct {
	AlertsSuppressionRuleName string `pulumi:"alertsSuppressionRuleName"`
}


type LookupAlertsSuppressionRuleResult struct {
	AlertType              string                          `pulumi:"alertType"`
	Comment                *string                         `pulumi:"comment"`
	ExpirationDateUtc      *string                         `pulumi:"expirationDateUtc"`
	Id                     string                          `pulumi:"id"`
	LastModifiedUtc        string                          `pulumi:"lastModifiedUtc"`
	Name                   string                          `pulumi:"name"`
	Reason                 string                          `pulumi:"reason"`
	State                  string                          `pulumi:"state"`
	SuppressionAlertsScope *SuppressionAlertsScopeResponse `pulumi:"suppressionAlertsScope"`
	Type                   string                          `pulumi:"type"`
}
