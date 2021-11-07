


package v20140401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlertRule(ctx *pulumi.Context, args *LookupAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupAlertRuleResult, error) {
	var rv LookupAlertRuleResult
	err := ctx.Invoke("azure-native:insights/v20140401:getAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupAlertRuleResult struct {
	Action            interface{}       `pulumi:"action"`
	Actions           []interface{}     `pulumi:"actions"`
	Condition         interface{}       `pulumi:"condition"`
	Description       *string           `pulumi:"description"`
	Id                string            `pulumi:"id"`
	IsEnabled         bool              `pulumi:"isEnabled"`
	LastUpdatedTime   string            `pulumi:"lastUpdatedTime"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}
