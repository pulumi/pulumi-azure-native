


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupActivityLogAlert(ctx *pulumi.Context, args *LookupActivityLogAlertArgs, opts ...pulumi.InvokeOption) (*LookupActivityLogAlertResult, error) {
	var rv LookupActivityLogAlertResult
	err := ctx.Invoke("azure-native:insights:getActivityLogAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupActivityLogAlertArgs struct {
	ActivityLogAlertName string `pulumi:"activityLogAlertName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupActivityLogAlertResult struct {
	Actions     ActionListResponse              `pulumi:"actions"`
	Condition   AlertRuleAllOfConditionResponse `pulumi:"condition"`
	Description *string                         `pulumi:"description"`
	Enabled     *bool                           `pulumi:"enabled"`
	Id          string                          `pulumi:"id"`
	Location    *string                         `pulumi:"location"`
	Name        string                          `pulumi:"name"`
	Scopes      []string                        `pulumi:"scopes"`
	Tags        map[string]string               `pulumi:"tags"`
	Type        string                          `pulumi:"type"`
}


func (val *LookupActivityLogAlertResult) Defaults() *LookupActivityLogAlertResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		enabled_ := true
		tmp.Enabled = &enabled_
	}
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}
