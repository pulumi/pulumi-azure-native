


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupActivityLogAlert(ctx *pulumi.Context, args *LookupActivityLogAlertArgs, opts ...pulumi.InvokeOption) (*LookupActivityLogAlertResult, error) {
	var rv LookupActivityLogAlertResult
	err := ctx.Invoke("azure-native:insights/v20201001:getActivityLogAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
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
