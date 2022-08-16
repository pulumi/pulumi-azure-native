


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkPeering(ctx *pulumi.Context, args *LookupVirtualNetworkPeeringArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkPeeringResult, error) {
	var rv LookupVirtualNetworkPeeringResult
	err := ctx.Invoke("azure-native:network/v20180601:getVirtualNetworkPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkPeeringArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkName        string `pulumi:"virtualNetworkName"`
	VirtualNetworkPeeringName string `pulumi:"virtualNetworkPeeringName"`
}


type LookupVirtualNetworkPeeringResult struct {
	AllowForwardedTraffic     *bool                 `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool                 `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool                 `pulumi:"allowVirtualNetworkAccess"`
	Etag                      *string               `pulumi:"etag"`
	Id                        *string               `pulumi:"id"`
	Name                      *string               `pulumi:"name"`
	PeeringState              *string               `pulumi:"peeringState"`
	ProvisioningState         *string               `pulumi:"provisioningState"`
	RemoteAddressSpace        *AddressSpaceResponse `pulumi:"remoteAddressSpace"`
	RemoteVirtualNetwork      *SubResourceResponse  `pulumi:"remoteVirtualNetwork"`
	UseRemoteGateways         *bool                 `pulumi:"useRemoteGateways"`
}

func LookupVirtualNetworkPeeringOutput(ctx *pulumi.Context, args LookupVirtualNetworkPeeringOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkPeeringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkPeeringResult, error) {
			args := v.(LookupVirtualNetworkPeeringArgs)
			r, err := LookupVirtualNetworkPeering(ctx, &args, opts...)
			var s LookupVirtualNetworkPeeringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkPeeringResultOutput)
}

type LookupVirtualNetworkPeeringOutputArgs struct {
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkName        pulumi.StringInput `pulumi:"virtualNetworkName"`
	VirtualNetworkPeeringName pulumi.StringInput `pulumi:"virtualNetworkPeeringName"`
}

func (LookupVirtualNetworkPeeringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkPeeringArgs)(nil)).Elem()
}


type LookupVirtualNetworkPeeringResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkPeeringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkPeeringResult)(nil)).Elem()
}

func (o LookupVirtualNetworkPeeringResultOutput) ToLookupVirtualNetworkPeeringResultOutput() LookupVirtualNetworkPeeringResultOutput {
	return o
}

func (o LookupVirtualNetworkPeeringResultOutput) ToLookupVirtualNetworkPeeringResultOutputWithContext(ctx context.Context) LookupVirtualNetworkPeeringResultOutput {
	return o
}

func (o LookupVirtualNetworkPeeringResultOutput) AllowForwardedTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *bool { return v.AllowForwardedTraffic }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) AllowGatewayTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *bool { return v.AllowGatewayTransit }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) AllowVirtualNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *bool { return v.AllowVirtualNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) PeeringState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *string { return v.PeeringState }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) RemoteAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *AddressSpaceResponse { return v.RemoteAddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) RemoteVirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *SubResourceResponse { return v.RemoteVirtualNetwork }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualNetworkPeeringResultOutput) UseRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkPeeringResult) *bool { return v.UseRemoteGateways }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkPeeringResultOutput{})
}
