


package alertsmanagement

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSmartDetectorAlertRule(ctx *pulumi.Context, args *LookupSmartDetectorAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupSmartDetectorAlertRuleResult, error) {
	var rv LookupSmartDetectorAlertRuleResult
	err := ctx.Invoke("azure-native:alertsmanagement:getSmartDetectorAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSmartDetectorAlertRuleArgs struct {
	AlertRuleName     string `pulumi:"alertRuleName"`
	ExpandDetector    *bool  `pulumi:"expandDetector"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSmartDetectorAlertRuleResult struct {
	ActionGroups ActionGroupsInformationResponse `pulumi:"actionGroups"`
	Description  *string                         `pulumi:"description"`
	Detector     DetectorResponse                `pulumi:"detector"`
	Frequency    string                          `pulumi:"frequency"`
	Id           string                          `pulumi:"id"`
	Location     *string                         `pulumi:"location"`
	Name         string                          `pulumi:"name"`
	Scope        []string                        `pulumi:"scope"`
	Severity     string                          `pulumi:"severity"`
	State        string                          `pulumi:"state"`
	Tags         map[string]string               `pulumi:"tags"`
	Throttling   *ThrottlingInformationResponse  `pulumi:"throttling"`
	Type         string                          `pulumi:"type"`
}


func (val *LookupSmartDetectorAlertRuleResult) Defaults() *LookupSmartDetectorAlertRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}
