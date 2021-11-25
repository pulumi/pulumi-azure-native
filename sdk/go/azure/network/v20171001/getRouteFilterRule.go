


package v20171001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouteFilterRule(ctx *pulumi.Context, args *LookupRouteFilterRuleArgs, opts ...pulumi.InvokeOption) (*LookupRouteFilterRuleResult, error) {
	var rv LookupRouteFilterRuleResult
	err := ctx.Invoke("azure-native:network/v20171001:getRouteFilterRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteFilterRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteFilterName   string `pulumi:"routeFilterName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupRouteFilterRuleResult struct {
	Access              string            `pulumi:"access"`
	Communities         []string          `pulumi:"communities"`
	Etag                string            `pulumi:"etag"`
	Id                  *string           `pulumi:"id"`
	Location            *string           `pulumi:"location"`
	Name                *string           `pulumi:"name"`
	ProvisioningState   string            `pulumi:"provisioningState"`
	RouteFilterRuleType string            `pulumi:"routeFilterRuleType"`
	Tags                map[string]string `pulumi:"tags"`
}
