


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayVpnclientConnectionHealth(ctx *pulumi.Context, args *GetVirtualNetworkGatewayVpnclientConnectionHealthArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayVpnclientConnectionHealthResult, error) {
	var rv GetVirtualNetworkGatewayVpnclientConnectionHealthResult
	err := ctx.Invoke("azure-native:network/v20191201:getVirtualNetworkGatewayVpnclientConnectionHealth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayVpnclientConnectionHealthArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}


type GetVirtualNetworkGatewayVpnclientConnectionHealthResult struct {
	Value []VpnClientConnectionHealthDetailResponse `pulumi:"value"`
}

func GetVirtualNetworkGatewayVpnclientConnectionHealthOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayVpnclientConnectionHealthOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayVpnclientConnectionHealthResult, error) {
			args := v.(GetVirtualNetworkGatewayVpnclientConnectionHealthArgs)
			r, err := GetVirtualNetworkGatewayVpnclientConnectionHealth(ctx, &args, opts...)
			var s GetVirtualNetworkGatewayVpnclientConnectionHealthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput)
}

type GetVirtualNetworkGatewayVpnclientConnectionHealthOutputArgs struct {
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayVpnclientConnectionHealthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnclientConnectionHealthArgs)(nil)).Elem()
}


type GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnclientConnectionHealthResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput) ToGetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput() GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput) ToGetVirtualNetworkGatewayVpnclientConnectionHealthResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput) Value() VpnClientConnectionHealthDetailResponseArrayOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientConnectionHealthResult) []VpnClientConnectionHealthDetailResponse {
		return v.Value
	}).(VpnClientConnectionHealthDetailResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayVpnclientConnectionHealthResultOutput{})
}
