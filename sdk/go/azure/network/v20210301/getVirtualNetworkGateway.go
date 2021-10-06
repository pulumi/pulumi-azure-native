


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkGateway(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayResult, error) {
	var rv LookupVirtualNetworkGatewayResult
	err := ctx.Invoke("azure-native:network/v20210301:getVirtualNetworkGateway", args, &rv, opts...)
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
	ActiveActive                    *bool                                          `pulumi:"activeActive"`
	BgpSettings                     *BgpSettingsResponse                           `pulumi:"bgpSettings"`
	CustomRoutes                    *AddressSpaceResponse                          `pulumi:"customRoutes"`
	EnableBgp                       *bool                                          `pulumi:"enableBgp"`
	EnableBgpRouteTranslationForNat *bool                                          `pulumi:"enableBgpRouteTranslationForNat"`
	EnableDnsForwarding             *bool                                          `pulumi:"enableDnsForwarding"`
	EnablePrivateIpAddress          *bool                                          `pulumi:"enablePrivateIpAddress"`
	Etag                            string                                         `pulumi:"etag"`
	ExtendedLocation                *ExtendedLocationResponse                      `pulumi:"extendedLocation"`
	GatewayDefaultSite              *SubResourceResponse                           `pulumi:"gatewayDefaultSite"`
	GatewayType                     *string                                        `pulumi:"gatewayType"`
	Id                              *string                                        `pulumi:"id"`
	InboundDnsForwardingEndpoint    string                                         `pulumi:"inboundDnsForwardingEndpoint"`
	IpConfigurations                []VirtualNetworkGatewayIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location                        *string                                        `pulumi:"location"`
	Name                            string                                         `pulumi:"name"`
	NatRules                        []VirtualNetworkGatewayNatRuleResponse         `pulumi:"natRules"`
	ProvisioningState               string                                         `pulumi:"provisioningState"`
	ResourceGuid                    string                                         `pulumi:"resourceGuid"`
	Sku                             *VirtualNetworkGatewaySkuResponse              `pulumi:"sku"`
	Tags                            map[string]string                              `pulumi:"tags"`
	Type                            string                                         `pulumi:"type"`
	VNetExtendedLocationResourceId  *string                                        `pulumi:"vNetExtendedLocationResourceId"`
	VpnClientConfiguration          *VpnClientConfigurationResponse                `pulumi:"vpnClientConfiguration"`
	VpnGatewayGeneration            *string                                        `pulumi:"vpnGatewayGeneration"`
	VpnType                         *string                                        `pulumi:"vpnType"`
}
