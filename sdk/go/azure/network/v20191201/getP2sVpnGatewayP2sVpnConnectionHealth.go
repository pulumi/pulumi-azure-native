


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetP2sVpnGatewayP2sVpnConnectionHealth(ctx *pulumi.Context, args *GetP2sVpnGatewayP2sVpnConnectionHealthArgs, opts ...pulumi.InvokeOption) (*GetP2sVpnGatewayP2sVpnConnectionHealthResult, error) {
	var rv GetP2sVpnGatewayP2sVpnConnectionHealthResult
	err := ctx.Invoke("azure-native:network/v20191201:getP2sVpnGatewayP2sVpnConnectionHealth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetP2sVpnGatewayP2sVpnConnectionHealthArgs struct {
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetP2sVpnGatewayP2sVpnConnectionHealthResult struct {
	Etag                        string                               `pulumi:"etag"`
	Id                          *string                              `pulumi:"id"`
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

func GetP2sVpnGatewayP2sVpnConnectionHealthOutput(ctx *pulumi.Context, args GetP2sVpnGatewayP2sVpnConnectionHealthOutputArgs, opts ...pulumi.InvokeOption) GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetP2sVpnGatewayP2sVpnConnectionHealthResult, error) {
			args := v.(GetP2sVpnGatewayP2sVpnConnectionHealthArgs)
			r, err := GetP2sVpnGatewayP2sVpnConnectionHealth(ctx, &args, opts...)
			var s GetP2sVpnGatewayP2sVpnConnectionHealthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput)
}

type GetP2sVpnGatewayP2sVpnConnectionHealthOutputArgs struct {
	GatewayName       pulumi.StringInput `pulumi:"gatewayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetP2sVpnGatewayP2sVpnConnectionHealthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetP2sVpnGatewayP2sVpnConnectionHealthArgs)(nil)).Elem()
}


type GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput struct{ *pulumi.OutputState }

func (GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetP2sVpnGatewayP2sVpnConnectionHealthResult)(nil)).Elem()
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) ToGetP2sVpnGatewayP2sVpnConnectionHealthResultOutput() GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput {
	return o
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) ToGetP2sVpnGatewayP2sVpnConnectionHealthResultOutputWithContext(ctx context.Context) GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput {
	return o
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) P2SConnectionConfigurations() P2SConnectionConfigurationResponseArrayOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) []P2SConnectionConfigurationResponse {
		return v.P2SConnectionConfigurations
	}).(P2SConnectionConfigurationResponseArrayOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) VpnClientConnectionHealth() VpnClientConnectionHealthResponseOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) VpnClientConnectionHealthResponse {
		return v.VpnClientConnectionHealth
	}).(VpnClientConnectionHealthResponseOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) VpnGatewayScaleUnit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) *int { return v.VpnGatewayScaleUnit }).(pulumi.IntPtrOutput)
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) VpnServerConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) *SubResourceResponse {
		return v.VpnServerConfiguration
	}).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput{})
}
