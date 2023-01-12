


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayBgpPeerStatus(ctx *pulumi.Context, args *GetVirtualNetworkGatewayBgpPeerStatusArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayBgpPeerStatusResult, error) {
	var rv GetVirtualNetworkGatewayBgpPeerStatusResult
	err := ctx.Invoke("azure-native:network/v20210801:getVirtualNetworkGatewayBgpPeerStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayBgpPeerStatusArgs struct {
	Peer                      *string `pulumi:"peer"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string  `pulumi:"virtualNetworkGatewayName"`
}


type GetVirtualNetworkGatewayBgpPeerStatusResult struct {
	Value []BgpPeerStatusResponse `pulumi:"value"`
}

func GetVirtualNetworkGatewayBgpPeerStatusOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayBgpPeerStatusOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayBgpPeerStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayBgpPeerStatusResult, error) {
			args := v.(GetVirtualNetworkGatewayBgpPeerStatusArgs)
			r, err := GetVirtualNetworkGatewayBgpPeerStatus(ctx, &args, opts...)
			var s GetVirtualNetworkGatewayBgpPeerStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualNetworkGatewayBgpPeerStatusResultOutput)
}

type GetVirtualNetworkGatewayBgpPeerStatusOutputArgs struct {
	Peer                      pulumi.StringPtrInput `pulumi:"peer"`
	ResourceGroupName         pulumi.StringInput    `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName pulumi.StringInput    `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayBgpPeerStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayBgpPeerStatusArgs)(nil)).Elem()
}


type GetVirtualNetworkGatewayBgpPeerStatusResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayBgpPeerStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayBgpPeerStatusResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayBgpPeerStatusResultOutput) ToGetVirtualNetworkGatewayBgpPeerStatusResultOutput() GetVirtualNetworkGatewayBgpPeerStatusResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayBgpPeerStatusResultOutput) ToGetVirtualNetworkGatewayBgpPeerStatusResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayBgpPeerStatusResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayBgpPeerStatusResultOutput) Value() BgpPeerStatusResponseArrayOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayBgpPeerStatusResult) []BgpPeerStatusResponse { return v.Value }).(BgpPeerStatusResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayBgpPeerStatusResultOutput{})
}
