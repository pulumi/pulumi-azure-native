


package v20181101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("azure-native:network/v20181101:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerArgs struct {
	Expand            *string `pulumi:"expand"`
	LoadBalancerName  string  `pulumi:"loadBalancerName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupLoadBalancerResult struct {
	BackendAddressPools      []BackendAddressPoolResponse      `pulumi:"backendAddressPools"`
	Etag                     *string                           `pulumi:"etag"`
	FrontendIPConfigurations []FrontendIPConfigurationResponse `pulumi:"frontendIPConfigurations"`
	Id                       *string                           `pulumi:"id"`
	InboundNatPools          []InboundNatPoolResponse          `pulumi:"inboundNatPools"`
	InboundNatRules          []InboundNatRuleResponse          `pulumi:"inboundNatRules"`
	LoadBalancingRules       []LoadBalancingRuleResponse       `pulumi:"loadBalancingRules"`
	Location                 *string                           `pulumi:"location"`
	Name                     string                            `pulumi:"name"`
	OutboundRules            []OutboundRuleResponse            `pulumi:"outboundRules"`
	Probes                   []ProbeResponse                   `pulumi:"probes"`
	ProvisioningState        *string                           `pulumi:"provisioningState"`
	ResourceGuid             *string                           `pulumi:"resourceGuid"`
	Sku                      *LoadBalancerSkuResponse          `pulumi:"sku"`
	Tags                     map[string]string                 `pulumi:"tags"`
	Type                     string                            `pulumi:"type"`
}
