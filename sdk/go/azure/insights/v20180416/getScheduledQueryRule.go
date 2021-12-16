


package v20180416

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledQueryRule(ctx *pulumi.Context, args *LookupScheduledQueryRuleArgs, opts ...pulumi.InvokeOption) (*LookupScheduledQueryRuleResult, error) {
	var rv LookupScheduledQueryRuleResult
	err := ctx.Invoke("azure-native:insights/v20180416:getScheduledQueryRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupScheduledQueryRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupScheduledQueryRuleResult struct {
	Action                   interface{}       `pulumi:"action"`
	AutoMitigate             *bool             `pulumi:"autoMitigate"`
	CreatedWithApiVersion    string            `pulumi:"createdWithApiVersion"`
	Description              *string           `pulumi:"description"`
	DisplayName              *string           `pulumi:"displayName"`
	Enabled                  *string           `pulumi:"enabled"`
	Etag                     string            `pulumi:"etag"`
	Id                       string            `pulumi:"id"`
	IsLegacyLogAnalyticsRule bool              `pulumi:"isLegacyLogAnalyticsRule"`
	Kind                     string            `pulumi:"kind"`
	LastUpdatedTime          string            `pulumi:"lastUpdatedTime"`
	Location                 string            `pulumi:"location"`
	Name                     string            `pulumi:"name"`
	ProvisioningState        string            `pulumi:"provisioningState"`
	Schedule                 *ScheduleResponse `pulumi:"schedule"`
	Source                   SourceResponse    `pulumi:"source"`
	Tags                     map[string]string `pulumi:"tags"`
	Type                     string            `pulumi:"type"`
}


func (val *LookupScheduledQueryRuleResult) Defaults() *LookupScheduledQueryRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutoMitigate) {
		autoMitigate_ := false
		tmp.AutoMitigate = &autoMitigate_
	}
	return &tmp
}
