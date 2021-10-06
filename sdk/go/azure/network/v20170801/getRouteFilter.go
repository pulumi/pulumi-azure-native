


package v20170801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouteFilter(ctx *pulumi.Context, args *LookupRouteFilterArgs, opts ...pulumi.InvokeOption) (*LookupRouteFilterResult, error) {
	var rv LookupRouteFilterResult
	err := ctx.Invoke("azure-native:network/v20170801:getRouteFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteFilterArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RouteFilterName   string  `pulumi:"routeFilterName"`
}


type LookupRouteFilterResult struct {
	Etag              string                               `pulumi:"etag"`
	Id                *string                              `pulumi:"id"`
	Location          string                               `pulumi:"location"`
	Name              string                               `pulumi:"name"`
	Peerings          []ExpressRouteCircuitPeeringResponse `pulumi:"peerings"`
	ProvisioningState string                               `pulumi:"provisioningState"`
	Rules             []RouteFilterRuleResponse            `pulumi:"rules"`
	Tags              map[string]string                    `pulumi:"tags"`
	Type              string                               `pulumi:"type"`
}
