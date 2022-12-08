


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupP2sVpnGateway(ctx *pulumi.Context, args *LookupP2sVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupP2sVpnGatewayResult, error) {
	var rv LookupP2sVpnGatewayResult
	err := ctx.Invoke("azure-native:network/v20220501:getP2sVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupP2sVpnGatewayArgs struct {
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupP2sVpnGatewayResult struct {
	CustomDnsServers            []string                             `pulumi:"customDnsServers"`
	Etag                        string                               `pulumi:"etag"`
	Id                          *string                              `pulumi:"id"`
	IsRoutingPreferenceInternet *bool                                `pulumi:"isRoutingPreferenceInternet"`
	Location                    string                               `pulumi:"location"`
	Name                        string                               `pulumi:"name"`
	P2SConnectionConfigurations []P2SConnectionConfigurationResponse `pulumi:"p2SConnectionConfigurations"`
	ProvisioningState           string                               `pulumi:"provisioningState"`
	Tags                        map[string]string                    `pulumi:"tags"`
	Type                        string                               `pulumi:"type"`
	VirtualHub                  *SubResourceResponse                 `pulumi:"virtualHub"`
	VpnClientConnectionHealth   VpnClientConnectionHealthResponse    `pulumi:"vpnClientConnectionHealth"`
	VpnGatewayScaleUnit         *int                                 `pulumi:"vpnGatewayScaleUnit"`
	VpnServerConfiguration      *SubResourceResponse                 `pulumi:"vpnServerConfiguration"`
}

func LookupP2sVpnGatewayOutput(ctx *pulumi.Context, args LookupP2sVpnGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupP2sVpnGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupP2sVpnGatewayResult, error) {
			args := v.(LookupP2sVpnGatewayArgs)
			r, err := LookupP2sVpnGateway(ctx, &args, opts...)
			var s LookupP2sVpnGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupP2sVpnGatewayResultOutput)
}

type LookupP2sVpnGatewayOutputArgs struct {
	GatewayName       pulumi.StringInput `pulumi:"gatewayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupP2sVpnGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupP2sVpnGatewayArgs)(nil)).Elem()
}


type LookupP2sVpnGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupP2sVpnGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupP2sVpnGatewayResult)(nil)).Elem()
}

func (o LookupP2sVpnGatewayResultOutput) ToLookupP2sVpnGatewayResultOutput() LookupP2sVpnGatewayResultOutput {
	return o
}

func (o LookupP2sVpnGatewayResultOutput) ToLookupP2sVpnGatewayResultOutputWithContext(ctx context.Context) LookupP2sVpnGatewayResultOutput {
	return o
}

func (o LookupP2sVpnGatewayResultOutput) CustomDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) []string { return v.CustomDnsServers }).(pulumi.StringArrayOutput)
}

func (o LookupP2sVpnGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupP2sVpnGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupP2sVpnGatewayResultOutput) IsRoutingPreferenceInternet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *bool { return v.IsRoutingPreferenceInternet }).(pulumi.BoolPtrOutput)
}

func (o LookupP2sVpnGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupP2sVpnGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupP2sVpnGatewayResultOutput) P2SConnectionConfigurations() P2SConnectionConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) []P2SConnectionConfigurationResponse {
		return v.P2SConnectionConfigurations
	}).(P2SConnectionConfigurationResponseArrayOutput)
}

func (o LookupP2sVpnGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupP2sVpnGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupP2sVpnGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupP2sVpnGatewayResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func (o LookupP2sVpnGatewayResultOutput) VpnClientConnectionHealth() VpnClientConnectionHealthResponseOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) VpnClientConnectionHealthResponse {
		return v.VpnClientConnectionHealth
	}).(VpnClientConnectionHealthResponseOutput)
}

func (o LookupP2sVpnGatewayResultOutput) VpnGatewayScaleUnit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *int { return v.VpnGatewayScaleUnit }).(pulumi.IntPtrOutput)
}

func (o LookupP2sVpnGatewayResultOutput) VpnServerConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupP2sVpnGatewayResult) *SubResourceResponse { return v.VpnServerConfiguration }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupP2sVpnGatewayResultOutput{})
}
