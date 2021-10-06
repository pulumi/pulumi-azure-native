


package v20200804preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHealthAlert(ctx *pulumi.Context, args *LookupHealthAlertArgs, opts ...pulumi.InvokeOption) (*LookupHealthAlertResult, error) {
	var rv LookupHealthAlertResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20200804preview:getHealthAlert", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHealthAlertArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupHealthAlertResult struct {
	Actions         []HealthAlertActionResponse `pulumi:"actions"`
	Criteria        HealthAlertCriteriaResponse `pulumi:"criteria"`
	Description     string                      `pulumi:"description"`
	Enabled         bool                        `pulumi:"enabled"`
	Id              string                      `pulumi:"id"`
	LastUpdatedTime string                      `pulumi:"lastUpdatedTime"`
	Location        string                      `pulumi:"location"`
	Name            string                      `pulumi:"name"`
	Scopes          []string                    `pulumi:"scopes"`
	Tags            map[string]string           `pulumi:"tags"`
	Type            string                      `pulumi:"type"`
}
