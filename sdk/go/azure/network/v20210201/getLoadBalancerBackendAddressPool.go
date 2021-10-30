


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoadBalancerBackendAddressPool(ctx *pulumi.Context, args *LookupLoadBalancerBackendAddressPoolArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerBackendAddressPoolResult, error) {
	var rv LookupLoadBalancerBackendAddressPoolResult
	err := ctx.Invoke("azure-native:network/v20210201:getLoadBalancerBackendAddressPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerBackendAddressPoolArgs struct {
	BackendAddressPoolName string `pulumi:"backendAddressPoolName"`
	LoadBalancerName       string `pulumi:"loadBalancerName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupLoadBalancerBackendAddressPoolResult struct {
	BackendIPConfigurations      []NetworkInterfaceIPConfigurationResponse    `pulumi:"backendIPConfigurations"`
	Etag                         string                                       `pulumi:"etag"`
	Id                           *string                                      `pulumi:"id"`
	LoadBalancerBackendAddresses []LoadBalancerBackendAddressResponse         `pulumi:"loadBalancerBackendAddresses"`
	LoadBalancingRules           []SubResourceResponse                        `pulumi:"loadBalancingRules"`
	Location                     *string                                      `pulumi:"location"`
	Name                         *string                                      `pulumi:"name"`
	OutboundRule                 SubResourceResponse                          `pulumi:"outboundRule"`
	OutboundRules                []SubResourceResponse                        `pulumi:"outboundRules"`
	ProvisioningState            string                                       `pulumi:"provisioningState"`
	TunnelInterfaces             []GatewayLoadBalancerTunnelInterfaceResponse `pulumi:"tunnelInterfaces"`
	Type                         string                                       `pulumi:"type"`
}
