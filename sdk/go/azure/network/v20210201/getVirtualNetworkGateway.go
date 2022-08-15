


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkGateway(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayResult, error) {
	var rv LookupVirtualNetworkGatewayResult
	err := ctx.Invoke("azure-native:network/v20210201:getVirtualNetworkGateway", args, &rv, opts...)
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

func LookupVirtualNetworkGatewayOutput(ctx *pulumi.Context, args LookupVirtualNetworkGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkGatewayResult, error) {
			args := v.(LookupVirtualNetworkGatewayArgs)
			r, err := LookupVirtualNetworkGateway(ctx, &args, opts...)
			var s LookupVirtualNetworkGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkGatewayResultOutput)
}

type LookupVirtualNetworkGatewayOutputArgs struct {
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (LookupVirtualNetworkGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayArgs)(nil)).Elem()
}


type LookupVirtualNetworkGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayResult)(nil)).Elem()
}

func (o LookupVirtualNetworkGatewayResultOutput) ToLookupVirtualNetworkGatewayResultOutput() LookupVirtualNetworkGatewayResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayResultOutput) ToLookupVirtualNetworkGatewayResultOutputWithContext(ctx context.Context) LookupVirtualNetworkGatewayResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayResultOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *bool { return v.ActiveActive }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *BgpSettingsResponse { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) CustomRoutes() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *AddressSpaceResponse { return v.CustomRoutes }).(AddressSpaceResponsePtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) EnableBgpRouteTranslationForNat() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *bool { return v.EnableBgpRouteTranslationForNat }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) EnableDnsForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *bool { return v.EnableDnsForwarding }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) EnablePrivateIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *bool { return v.EnablePrivateIpAddress }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) GatewayDefaultSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *SubResourceResponse { return v.GatewayDefaultSite }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *string { return v.GatewayType }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) InboundDnsForwardingEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.InboundDnsForwardingEndpoint }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) []VirtualNetworkGatewayIPConfigurationResponse {
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) NatRules() VirtualNetworkGatewayNatRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) []VirtualNetworkGatewayNatRuleResponse { return v.NatRules }).(VirtualNetworkGatewayNatRuleResponseArrayOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) Sku() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *VirtualNetworkGatewaySkuResponse { return v.Sku }).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) VNetExtendedLocationResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *string { return v.VNetExtendedLocationResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) VpnClientConfiguration() VpnClientConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *VpnClientConfigurationResponse {
		return v.VpnClientConfiguration
	}).(VpnClientConfigurationResponsePtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) VpnGatewayGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *string { return v.VpnGatewayGeneration }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayResultOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayResult) *string { return v.VpnType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkGatewayResultOutput{})
}
