


package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetvNetPeering(ctx *pulumi.Context, args *GetvNetPeeringArgs, opts ...pulumi.InvokeOption) (*GetvNetPeeringResult, error) {
	var rv GetvNetPeeringResult
	err := ctx.Invoke("azure-native:databricks/v20180401:getvNetPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetvNetPeeringArgs struct {
	PeeringName       string `pulumi:"peeringName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type GetvNetPeeringResult struct {
	AllowForwardedTraffic     *bool                                                                  `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool                                                                  `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool                                                                  `pulumi:"allowVirtualNetworkAccess"`
	DatabricksAddressSpace    *AddressSpaceResponse                                                  `pulumi:"databricksAddressSpace"`
	DatabricksVirtualNetwork  *VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork `pulumi:"databricksVirtualNetwork"`
	Id                        string                                                                 `pulumi:"id"`
	Name                      string                                                                 `pulumi:"name"`
	PeeringState              string                                                                 `pulumi:"peeringState"`
	ProvisioningState         string                                                                 `pulumi:"provisioningState"`
	RemoteAddressSpace        *AddressSpaceResponse                                                  `pulumi:"remoteAddressSpace"`
	RemoteVirtualNetwork      VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork      `pulumi:"remoteVirtualNetwork"`
	Type                      string                                                                 `pulumi:"type"`
	UseRemoteGateways         *bool                                                                  `pulumi:"useRemoteGateways"`
}

func GetvNetPeeringOutput(ctx *pulumi.Context, args GetvNetPeeringOutputArgs, opts ...pulumi.InvokeOption) GetvNetPeeringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetvNetPeeringResult, error) {
			args := v.(GetvNetPeeringArgs)
			r, err := GetvNetPeering(ctx, &args, opts...)
			var s GetvNetPeeringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetvNetPeeringResultOutput)
}

type GetvNetPeeringOutputArgs struct {
	PeeringName       pulumi.StringInput `pulumi:"peeringName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetvNetPeeringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvNetPeeringArgs)(nil)).Elem()
}


type GetvNetPeeringResultOutput struct{ *pulumi.OutputState }

func (GetvNetPeeringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvNetPeeringResult)(nil)).Elem()
}

func (o GetvNetPeeringResultOutput) ToGetvNetPeeringResultOutput() GetvNetPeeringResultOutput {
	return o
}

func (o GetvNetPeeringResultOutput) ToGetvNetPeeringResultOutputWithContext(ctx context.Context) GetvNetPeeringResultOutput {
	return o
}

func (o GetvNetPeeringResultOutput) AllowForwardedTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) *bool { return v.AllowForwardedTraffic }).(pulumi.BoolPtrOutput)
}

func (o GetvNetPeeringResultOutput) AllowGatewayTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) *bool { return v.AllowGatewayTransit }).(pulumi.BoolPtrOutput)
}

func (o GetvNetPeeringResultOutput) AllowVirtualNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) *bool { return v.AllowVirtualNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o GetvNetPeeringResultOutput) DatabricksAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) *AddressSpaceResponse { return v.DatabricksAddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o GetvNetPeeringResultOutput) DatabricksVirtualNetwork() VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) *VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork {
		return v.DatabricksVirtualNetwork
	}).(VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput)
}

func (o GetvNetPeeringResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetvNetPeeringResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetvNetPeeringResultOutput) PeeringState() pulumi.StringOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) string { return v.PeeringState }).(pulumi.StringOutput)
}

func (o GetvNetPeeringResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetvNetPeeringResultOutput) RemoteAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) *AddressSpaceResponse { return v.RemoteAddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o GetvNetPeeringResultOutput) RemoteVirtualNetwork() VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork {
		return v.RemoteVirtualNetwork
	}).(VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput)
}

func (o GetvNetPeeringResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetvNetPeeringResultOutput) UseRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetvNetPeeringResult) *bool { return v.UseRemoteGateways }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetvNetPeeringResultOutput{})
}
