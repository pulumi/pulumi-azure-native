


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkTap(ctx *pulumi.Context, args *LookupVirtualNetworkTapArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkTapResult, error) {
	var rv LookupVirtualNetworkTapResult
	err := ctx.Invoke("azure-native:network/v20210301:getVirtualNetworkTap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualNetworkTapArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TapName           string `pulumi:"tapName"`
}


type LookupVirtualNetworkTapResult struct {
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfigurationResponse           `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     *NetworkInterfaceIPConfigurationResponse   `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                *int                                       `pulumi:"destinationPort"`
	Etag                                           string                                     `pulumi:"etag"`
	Id                                             *string                                    `pulumi:"id"`
	Location                                       *string                                    `pulumi:"location"`
	Name                                           string                                     `pulumi:"name"`
	NetworkInterfaceTapConfigurations              []NetworkInterfaceTapConfigurationResponse `pulumi:"networkInterfaceTapConfigurations"`
	ProvisioningState                              string                                     `pulumi:"provisioningState"`
	ResourceGuid                                   string                                     `pulumi:"resourceGuid"`
	Tags                                           map[string]string                          `pulumi:"tags"`
	Type                                           string                                     `pulumi:"type"`
}


func (val *LookupVirtualNetworkTapResult) Defaults() *LookupVirtualNetworkTapResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DestinationLoadBalancerFrontEndIPConfiguration = tmp.DestinationLoadBalancerFrontEndIPConfiguration.Defaults()

	tmp.DestinationNetworkInterfaceIPConfiguration = tmp.DestinationNetworkInterfaceIPConfiguration.Defaults()

	return &tmp
}
