


package v20180201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkGateway(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayResult, error) {
	var rv LookupVirtualNetworkGatewayResult
	err := ctx.Invoke("azure-native:network/v20180201:getVirtualNetworkGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkGatewayArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}


type LookupVirtualNetworkGatewayResult struct {
	ActiveActive           *bool                                          `pulumi:"activeActive"`
	BgpSettings            *BgpSettingsResponse                           `pulumi:"bgpSettings"`
	EnableBgp              *bool                                          `pulumi:"enableBgp"`
	Etag                   *string                                        `pulumi:"etag"`
	GatewayDefaultSite     *SubResourceResponse                           `pulumi:"gatewayDefaultSite"`
	GatewayType            *string                                        `pulumi:"gatewayType"`
	Id                     *string                                        `pulumi:"id"`
	IpConfigurations       []VirtualNetworkGatewayIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location               *string                                        `pulumi:"location"`
	Name                   string                                         `pulumi:"name"`
	ProvisioningState      string                                         `pulumi:"provisioningState"`
	ResourceGuid           *string                                        `pulumi:"resourceGuid"`
	Sku                    *VirtualNetworkGatewaySkuResponse              `pulumi:"sku"`
	Tags                   map[string]string                              `pulumi:"tags"`
	Type                   string                                         `pulumi:"type"`
	VpnClientConfiguration *VpnClientConfigurationResponse                `pulumi:"vpnClientConfiguration"`
	VpnType                *string                                        `pulumi:"vpnType"`
}
