


package v20170801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupLocalNetworkGateway(ctx *pulumi.Context, args *LookupLocalNetworkGatewayArgs, opts ...pulumi.InvokeOption) (*LookupLocalNetworkGatewayResult, error) {
	var rv LookupLocalNetworkGatewayResult
	err := ctx.Invoke("azure-native:network/v20170801:getLocalNetworkGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocalNetworkGatewayArgs struct {
	LocalNetworkGatewayName string `pulumi:"localNetworkGatewayName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupLocalNetworkGatewayResult struct {
	BgpSettings              *BgpSettingsResponse  `pulumi:"bgpSettings"`
	Etag                     *string               `pulumi:"etag"`
	GatewayIpAddress         *string               `pulumi:"gatewayIpAddress"`
	Id                       *string               `pulumi:"id"`
	LocalNetworkAddressSpace *AddressSpaceResponse `pulumi:"localNetworkAddressSpace"`
	Location                 *string               `pulumi:"location"`
	Name                     string                `pulumi:"name"`
	ProvisioningState        string                `pulumi:"provisioningState"`
	ResourceGuid             *string               `pulumi:"resourceGuid"`
	Tags                     map[string]string     `pulumi:"tags"`
	Type                     string                `pulumi:"type"`
}

func LookupLocalNetworkGatewayOutput(ctx *pulumi.Context, args LookupLocalNetworkGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupLocalNetworkGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalNetworkGatewayResult, error) {
			args := v.(LookupLocalNetworkGatewayArgs)
			r, err := LookupLocalNetworkGateway(ctx, &args, opts...)
			var s LookupLocalNetworkGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalNetworkGatewayResultOutput)
}

type LookupLocalNetworkGatewayOutputArgs struct {
	LocalNetworkGatewayName pulumi.StringInput `pulumi:"localNetworkGatewayName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLocalNetworkGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalNetworkGatewayArgs)(nil)).Elem()
}


type LookupLocalNetworkGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupLocalNetworkGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalNetworkGatewayResult)(nil)).Elem()
}

func (o LookupLocalNetworkGatewayResultOutput) ToLookupLocalNetworkGatewayResultOutput() LookupLocalNetworkGatewayResultOutput {
	return o
}

func (o LookupLocalNetworkGatewayResultOutput) ToLookupLocalNetworkGatewayResultOutputWithContext(ctx context.Context) LookupLocalNetworkGatewayResultOutput {
	return o
}

func (o LookupLocalNetworkGatewayResultOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) *BgpSettingsResponse { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) *string { return v.GatewayIpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) LocalNetworkAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) *AddressSpaceResponse { return v.LocalNetworkAddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLocalNetworkGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalNetworkGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalNetworkGatewayResultOutput{})
}
