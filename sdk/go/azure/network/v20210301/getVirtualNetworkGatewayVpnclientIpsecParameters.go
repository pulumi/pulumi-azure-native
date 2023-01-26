


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayVpnclientIpsecParameters(ctx *pulumi.Context, args *GetVirtualNetworkGatewayVpnclientIpsecParametersArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayVpnclientIpsecParametersResult, error) {
	var rv GetVirtualNetworkGatewayVpnclientIpsecParametersResult
	err := ctx.Invoke("azure-native:network/v20210301:getVirtualNetworkGatewayVpnclientIpsecParameters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayVpnclientIpsecParametersArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}


type GetVirtualNetworkGatewayVpnclientIpsecParametersResult struct {
	DhGroup             string `pulumi:"dhGroup"`
	IkeEncryption       string `pulumi:"ikeEncryption"`
	IkeIntegrity        string `pulumi:"ikeIntegrity"`
	IpsecEncryption     string `pulumi:"ipsecEncryption"`
	IpsecIntegrity      string `pulumi:"ipsecIntegrity"`
	PfsGroup            string `pulumi:"pfsGroup"`
	SaDataSizeKilobytes int    `pulumi:"saDataSizeKilobytes"`
	SaLifeTimeSeconds   int    `pulumi:"saLifeTimeSeconds"`
}

func GetVirtualNetworkGatewayVpnclientIpsecParametersOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayVpnclientIpsecParametersOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayVpnclientIpsecParametersResult, error) {
			args := v.(GetVirtualNetworkGatewayVpnclientIpsecParametersArgs)
			r, err := GetVirtualNetworkGatewayVpnclientIpsecParameters(ctx, &args, opts...)
			var s GetVirtualNetworkGatewayVpnclientIpsecParametersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput)
}

type GetVirtualNetworkGatewayVpnclientIpsecParametersOutputArgs struct {
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayVpnclientIpsecParametersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnclientIpsecParametersArgs)(nil)).Elem()
}


type GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnclientIpsecParametersResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) ToGetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput() GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) ToGetVirtualNetworkGatewayVpnclientIpsecParametersResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) DhGroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.DhGroup }).(pulumi.StringOutput)
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) IkeEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.IkeEncryption }).(pulumi.StringOutput)
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) IkeIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.IkeIntegrity }).(pulumi.StringOutput)
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) IpsecEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.IpsecEncryption }).(pulumi.StringOutput)
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) IpsecIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.IpsecIntegrity }).(pulumi.StringOutput)
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) PfsGroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.PfsGroup }).(pulumi.StringOutput)
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) SaDataSizeKilobytes() pulumi.IntOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) int { return v.SaDataSizeKilobytes }).(pulumi.IntOutput)
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) SaLifeTimeSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) int { return v.SaLifeTimeSeconds }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput{})
}
