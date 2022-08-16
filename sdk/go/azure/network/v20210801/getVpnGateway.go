


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVpnGateway(ctx *pulumi.Context, args *LookupVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVpnGatewayResult, error) {
	var rv LookupVpnGatewayResult
	err := ctx.Invoke("azure-native:network/v20210801:getVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnGatewayArgs struct {
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupVpnGatewayResult struct {
	BgpSettings                     *BgpSettingsResponse                `pulumi:"bgpSettings"`
	Connections                     []VpnConnectionResponse             `pulumi:"connections"`
	EnableBgpRouteTranslationForNat *bool                               `pulumi:"enableBgpRouteTranslationForNat"`
	Etag                            string                              `pulumi:"etag"`
	Id                              *string                             `pulumi:"id"`
	IpConfigurations                []VpnGatewayIpConfigurationResponse `pulumi:"ipConfigurations"`
	IsRoutingPreferenceInternet     *bool                               `pulumi:"isRoutingPreferenceInternet"`
	Location                        string                              `pulumi:"location"`
	Name                            string                              `pulumi:"name"`
	NatRules                        []VpnGatewayNatRuleResponse         `pulumi:"natRules"`
	ProvisioningState               string                              `pulumi:"provisioningState"`
	Tags                            map[string]string                   `pulumi:"tags"`
	Type                            string                              `pulumi:"type"`
	VirtualHub                      *SubResourceResponse                `pulumi:"virtualHub"`
	VpnGatewayScaleUnit             *int                                `pulumi:"vpnGatewayScaleUnit"`
}

func LookupVpnGatewayOutput(ctx *pulumi.Context, args LookupVpnGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVpnGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnGatewayResult, error) {
			args := v.(LookupVpnGatewayArgs)
			r, err := LookupVpnGateway(ctx, &args, opts...)
			var s LookupVpnGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnGatewayResultOutput)
}

type LookupVpnGatewayOutputArgs struct {
	GatewayName       pulumi.StringInput `pulumi:"gatewayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVpnGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnGatewayArgs)(nil)).Elem()
}


type LookupVpnGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVpnGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnGatewayResult)(nil)).Elem()
}

func (o LookupVpnGatewayResultOutput) ToLookupVpnGatewayResultOutput() LookupVpnGatewayResultOutput {
	return o
}

func (o LookupVpnGatewayResultOutput) ToLookupVpnGatewayResultOutputWithContext(ctx context.Context) LookupVpnGatewayResultOutput {
	return o
}

func (o LookupVpnGatewayResultOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *BgpSettingsResponse { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o LookupVpnGatewayResultOutput) Connections() VpnConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) []VpnConnectionResponse { return v.Connections }).(VpnConnectionResponseArrayOutput)
}

func (o LookupVpnGatewayResultOutput) EnableBgpRouteTranslationForNat() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *bool { return v.EnableBgpRouteTranslationForNat }).(pulumi.BoolPtrOutput)
}

func (o LookupVpnGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVpnGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVpnGatewayResultOutput) IpConfigurations() VpnGatewayIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) []VpnGatewayIpConfigurationResponse { return v.IpConfigurations }).(VpnGatewayIpConfigurationResponseArrayOutput)
}

func (o LookupVpnGatewayResultOutput) IsRoutingPreferenceInternet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *bool { return v.IsRoutingPreferenceInternet }).(pulumi.BoolPtrOutput)
}

func (o LookupVpnGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVpnGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVpnGatewayResultOutput) NatRules() VpnGatewayNatRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) []VpnGatewayNatRuleResponse { return v.NatRules }).(VpnGatewayNatRuleResponseArrayOutput)
}

func (o LookupVpnGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVpnGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVpnGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVpnGatewayResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func (o LookupVpnGatewayResultOutput) VpnGatewayScaleUnit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVpnGatewayResult) *int { return v.VpnGatewayScaleUnit }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnGatewayResultOutput{})
}
