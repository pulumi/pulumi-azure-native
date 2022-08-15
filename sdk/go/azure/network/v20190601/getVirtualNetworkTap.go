


package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkTap(ctx *pulumi.Context, args *LookupVirtualNetworkTapArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkTapResult, error) {
	var rv LookupVirtualNetworkTapResult
	err := ctx.Invoke("azure-native:network/v20190601:getVirtualNetworkTap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkTapArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TapName           string `pulumi:"tapName"`
}


type LookupVirtualNetworkTapResult struct {
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfigurationResponse           `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     *NetworkInterfaceIPConfigurationResponse   `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                *int                                       `pulumi:"destinationPort"`
	Etag                                           *string                                    `pulumi:"etag"`
	Id                                             *string                                    `pulumi:"id"`
	Location                                       *string                                    `pulumi:"location"`
	Name                                           string                                     `pulumi:"name"`
	NetworkInterfaceTapConfigurations              []NetworkInterfaceTapConfigurationResponse `pulumi:"networkInterfaceTapConfigurations"`
	ProvisioningState                              string                                     `pulumi:"provisioningState"`
	ResourceGuid                                   string                                     `pulumi:"resourceGuid"`
	Tags                                           map[string]string                          `pulumi:"tags"`
	Type                                           string                                     `pulumi:"type"`
}

func LookupVirtualNetworkTapOutput(ctx *pulumi.Context, args LookupVirtualNetworkTapOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkTapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkTapResult, error) {
			args := v.(LookupVirtualNetworkTapArgs)
			r, err := LookupVirtualNetworkTap(ctx, &args, opts...)
			var s LookupVirtualNetworkTapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkTapResultOutput)
}

type LookupVirtualNetworkTapOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TapName           pulumi.StringInput `pulumi:"tapName"`
}

func (LookupVirtualNetworkTapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkTapArgs)(nil)).Elem()
}


type LookupVirtualNetworkTapResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkTapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkTapResult)(nil)).Elem()
}

func (o LookupVirtualNetworkTapResultOutput) ToLookupVirtualNetworkTapResultOutput() LookupVirtualNetworkTapResultOutput {
	return o
}

func (o LookupVirtualNetworkTapResultOutput) ToLookupVirtualNetworkTapResultOutputWithContext(ctx context.Context) LookupVirtualNetworkTapResultOutput {
	return o
}

func (o LookupVirtualNetworkTapResultOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *FrontendIPConfigurationResponse {
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationResponsePtrOutput)
}

func (o LookupVirtualNetworkTapResultOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *NetworkInterfaceIPConfigurationResponse {
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}

func (o LookupVirtualNetworkTapResultOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualNetworkTapResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkTapResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkTapResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkTapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkTapResultOutput) NetworkInterfaceTapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) []NetworkInterfaceTapConfigurationResponse {
		return v.NetworkInterfaceTapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

func (o LookupVirtualNetworkTapResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkTapResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkTapResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualNetworkTapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkTapResultOutput{})
}
