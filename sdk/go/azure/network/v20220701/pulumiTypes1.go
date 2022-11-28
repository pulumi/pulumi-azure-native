


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkTapType struct {
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfiguration         `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     *NetworkInterfaceIPConfiguration `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                *int                             `pulumi:"destinationPort"`
	Id                                             *string                          `pulumi:"id"`
	Location                                       *string                          `pulumi:"location"`
	Tags                                           map[string]string                `pulumi:"tags"`
}


func (val *VirtualNetworkTapType) Defaults() *VirtualNetworkTapType {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DestinationLoadBalancerFrontEndIPConfiguration = tmp.DestinationLoadBalancerFrontEndIPConfiguration.Defaults()

	tmp.DestinationNetworkInterfaceIPConfiguration = tmp.DestinationNetworkInterfaceIPConfiguration.Defaults()

	return &tmp
}





type VirtualNetworkTapTypeInput interface {
	pulumi.Input

	ToVirtualNetworkTapTypeOutput() VirtualNetworkTapTypeOutput
	ToVirtualNetworkTapTypeOutputWithContext(context.Context) VirtualNetworkTapTypeOutput
}

type VirtualNetworkTapTypeArgs struct {
	DestinationLoadBalancerFrontEndIPConfiguration FrontendIPConfigurationPtrInput         `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     NetworkInterfaceIPConfigurationPtrInput `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                pulumi.IntPtrInput                      `pulumi:"destinationPort"`
	Id                                             pulumi.StringPtrInput                   `pulumi:"id"`
	Location                                       pulumi.StringPtrInput                   `pulumi:"location"`
	Tags                                           pulumi.StringMapInput                   `pulumi:"tags"`
}


func (val *VirtualNetworkTapTypeArgs) Defaults() *VirtualNetworkTapTypeArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (VirtualNetworkTapTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkTapType)(nil)).Elem()
}

func (i VirtualNetworkTapTypeArgs) ToVirtualNetworkTapTypeOutput() VirtualNetworkTapTypeOutput {
	return i.ToVirtualNetworkTapTypeOutputWithContext(context.Background())
}

func (i VirtualNetworkTapTypeArgs) ToVirtualNetworkTapTypeOutputWithContext(ctx context.Context) VirtualNetworkTapTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapTypeOutput)
}

func (i VirtualNetworkTapTypeArgs) ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput {
	return i.ToVirtualNetworkTapTypePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkTapTypeArgs) ToVirtualNetworkTapTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTapTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapTypeOutput).ToVirtualNetworkTapTypePtrOutputWithContext(ctx)
}









type VirtualNetworkTapTypePtrInput interface {
	pulumi.Input

	ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput
	ToVirtualNetworkTapTypePtrOutputWithContext(context.Context) VirtualNetworkTapTypePtrOutput
}

type virtualNetworkTapTypePtrType VirtualNetworkTapTypeArgs

func VirtualNetworkTapTypePtr(v *VirtualNetworkTapTypeArgs) VirtualNetworkTapTypePtrInput {
	return (*virtualNetworkTapTypePtrType)(v)
}

func (*virtualNetworkTapTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTapType)(nil)).Elem()
}

func (i *virtualNetworkTapTypePtrType) ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput {
	return i.ToVirtualNetworkTapTypePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkTapTypePtrType) ToVirtualNetworkTapTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTapTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapTypePtrOutput)
}





type VirtualNetworkTapTypeArrayInput interface {
	pulumi.Input

	ToVirtualNetworkTapTypeArrayOutput() VirtualNetworkTapTypeArrayOutput
	ToVirtualNetworkTapTypeArrayOutputWithContext(context.Context) VirtualNetworkTapTypeArrayOutput
}

type VirtualNetworkTapTypeArray []VirtualNetworkTapTypeInput

func (VirtualNetworkTapTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkTapType)(nil)).Elem()
}

func (i VirtualNetworkTapTypeArray) ToVirtualNetworkTapTypeArrayOutput() VirtualNetworkTapTypeArrayOutput {
	return i.ToVirtualNetworkTapTypeArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkTapTypeArray) ToVirtualNetworkTapTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkTapTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapTypeArrayOutput)
}

type VirtualNetworkTapTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkTapType)(nil)).Elem()
}

func (o VirtualNetworkTapTypeOutput) ToVirtualNetworkTapTypeOutput() VirtualNetworkTapTypeOutput {
	return o
}

func (o VirtualNetworkTapTypeOutput) ToVirtualNetworkTapTypeOutputWithContext(ctx context.Context) VirtualNetworkTapTypeOutput {
	return o
}

func (o VirtualNetworkTapTypeOutput) ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput {
	return o.ToVirtualNetworkTapTypePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkTapTypeOutput) ToVirtualNetworkTapTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTapTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkTapType) *VirtualNetworkTapType {
		return &v
	}).(VirtualNetworkTapTypePtrOutput)
}

func (o VirtualNetworkTapTypeOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *FrontendIPConfiguration {
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *NetworkInterfaceIPConfiguration {
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkTapType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type VirtualNetworkTapTypePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTapType)(nil)).Elem()
}

func (o VirtualNetworkTapTypePtrOutput) ToVirtualNetworkTapTypePtrOutput() VirtualNetworkTapTypePtrOutput {
	return o
}

func (o VirtualNetworkTapTypePtrOutput) ToVirtualNetworkTapTypePtrOutputWithContext(ctx context.Context) VirtualNetworkTapTypePtrOutput {
	return o
}

func (o VirtualNetworkTapTypePtrOutput) Elem() VirtualNetworkTapTypeOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) VirtualNetworkTapType {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkTapType
		return ret
	}).(VirtualNetworkTapTypeOutput)
}

func (o VirtualNetworkTapTypePtrOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *FrontendIPConfiguration {
		if v == nil {
			return nil
		}
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *NetworkInterfaceIPConfiguration {
		if v == nil {
			return nil
		}
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *int {
		if v == nil {
			return nil
		}
		return v.DestinationPort
	}).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkTapType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type VirtualNetworkTapTypeArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkTapType)(nil)).Elem()
}

func (o VirtualNetworkTapTypeArrayOutput) ToVirtualNetworkTapTypeArrayOutput() VirtualNetworkTapTypeArrayOutput {
	return o
}

func (o VirtualNetworkTapTypeArrayOutput) ToVirtualNetworkTapTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkTapTypeArrayOutput {
	return o
}

func (o VirtualNetworkTapTypeArrayOutput) Index(i pulumi.IntInput) VirtualNetworkTapTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkTapType {
		return vs[0].([]VirtualNetworkTapType)[vs[1].(int)]
	}).(VirtualNetworkTapTypeOutput)
}

type VirtualNetworkTapResponse struct {
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfigurationResponse           `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     *NetworkInterfaceIPConfigurationResponse   `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                *int                                       `pulumi:"destinationPort"`
	Etag                                           string                                     `pulumi:"etag"`
	Id                                             *string                                    `pulumi:"id"`
	Location                                       *string                                    `pulumi:"location"`
	Name                                           string                                     `pulumi:"name"`
	NetworkInterfaceTapConfigurations              []NetworkInterfaceTapConfigurationResponse `pulumi:"networkInterfaceTapConfigurations"`
	ProvisioningState                              string                                     `pulumi:"provisioningState"`
	ResourceGuid                                   string                                     `pulumi:"resourceGuid"`
	Tags                                           map[string]string                          `pulumi:"tags"`
	Type                                           string                                     `pulumi:"type"`
}


func (val *VirtualNetworkTapResponse) Defaults() *VirtualNetworkTapResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DestinationLoadBalancerFrontEndIPConfiguration = tmp.DestinationLoadBalancerFrontEndIPConfiguration.Defaults()

	tmp.DestinationNetworkInterfaceIPConfiguration = tmp.DestinationNetworkInterfaceIPConfiguration.Defaults()

	return &tmp
}

type VirtualNetworkTapResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkTapResponse)(nil)).Elem()
}

func (o VirtualNetworkTapResponseOutput) ToVirtualNetworkTapResponseOutput() VirtualNetworkTapResponseOutput {
	return o
}

func (o VirtualNetworkTapResponseOutput) ToVirtualNetworkTapResponseOutputWithContext(ctx context.Context) VirtualNetworkTapResponseOutput {
	return o
}

func (o VirtualNetworkTapResponseOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *FrontendIPConfigurationResponse {
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapResponseOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *NetworkInterfaceIPConfigurationResponse {
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapResponseOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkTapResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapResponseOutput) NetworkInterfaceTapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) []NetworkInterfaceTapConfigurationResponse {
		return v.NetworkInterfaceTapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

func (o VirtualNetworkTapResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapResponseOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkTapResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkTapResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VirtualNetworkTapResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTapResponse)(nil)).Elem()
}

func (o VirtualNetworkTapResponsePtrOutput) ToVirtualNetworkTapResponsePtrOutput() VirtualNetworkTapResponsePtrOutput {
	return o
}

func (o VirtualNetworkTapResponsePtrOutput) ToVirtualNetworkTapResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkTapResponsePtrOutput {
	return o
}

func (o VirtualNetworkTapResponsePtrOutput) Elem() VirtualNetworkTapResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) VirtualNetworkTapResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkTapResponse
		return ret
	}).(VirtualNetworkTapResponseOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *FrontendIPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *NetworkInterfaceIPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *int {
		if v == nil {
			return nil
		}
		return v.DestinationPort
	}).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) NetworkInterfaceTapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) []NetworkInterfaceTapConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceTapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o VirtualNetworkTapResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTapResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkTapResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkTapResponse)(nil)).Elem()
}

func (o VirtualNetworkTapResponseArrayOutput) ToVirtualNetworkTapResponseArrayOutput() VirtualNetworkTapResponseArrayOutput {
	return o
}

func (o VirtualNetworkTapResponseArrayOutput) ToVirtualNetworkTapResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkTapResponseArrayOutput {
	return o
}

func (o VirtualNetworkTapResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkTapResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkTapResponse {
		return vs[0].([]VirtualNetworkTapResponse)[vs[1].(int)]
	}).(VirtualNetworkTapResponseOutput)
}

type VirtualRouterAutoScaleConfiguration struct {
	MinCapacity *int `pulumi:"minCapacity"`
}





type VirtualRouterAutoScaleConfigurationInput interface {
	pulumi.Input

	ToVirtualRouterAutoScaleConfigurationOutput() VirtualRouterAutoScaleConfigurationOutput
	ToVirtualRouterAutoScaleConfigurationOutputWithContext(context.Context) VirtualRouterAutoScaleConfigurationOutput
}

type VirtualRouterAutoScaleConfigurationArgs struct {
	MinCapacity pulumi.IntPtrInput `pulumi:"minCapacity"`
}

func (VirtualRouterAutoScaleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualRouterAutoScaleConfiguration)(nil)).Elem()
}

func (i VirtualRouterAutoScaleConfigurationArgs) ToVirtualRouterAutoScaleConfigurationOutput() VirtualRouterAutoScaleConfigurationOutput {
	return i.ToVirtualRouterAutoScaleConfigurationOutputWithContext(context.Background())
}

func (i VirtualRouterAutoScaleConfigurationArgs) ToVirtualRouterAutoScaleConfigurationOutputWithContext(ctx context.Context) VirtualRouterAutoScaleConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRouterAutoScaleConfigurationOutput)
}

func (i VirtualRouterAutoScaleConfigurationArgs) ToVirtualRouterAutoScaleConfigurationPtrOutput() VirtualRouterAutoScaleConfigurationPtrOutput {
	return i.ToVirtualRouterAutoScaleConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualRouterAutoScaleConfigurationArgs) ToVirtualRouterAutoScaleConfigurationPtrOutputWithContext(ctx context.Context) VirtualRouterAutoScaleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRouterAutoScaleConfigurationOutput).ToVirtualRouterAutoScaleConfigurationPtrOutputWithContext(ctx)
}









type VirtualRouterAutoScaleConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualRouterAutoScaleConfigurationPtrOutput() VirtualRouterAutoScaleConfigurationPtrOutput
	ToVirtualRouterAutoScaleConfigurationPtrOutputWithContext(context.Context) VirtualRouterAutoScaleConfigurationPtrOutput
}

type virtualRouterAutoScaleConfigurationPtrType VirtualRouterAutoScaleConfigurationArgs

func VirtualRouterAutoScaleConfigurationPtr(v *VirtualRouterAutoScaleConfigurationArgs) VirtualRouterAutoScaleConfigurationPtrInput {
	return (*virtualRouterAutoScaleConfigurationPtrType)(v)
}

func (*virtualRouterAutoScaleConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRouterAutoScaleConfiguration)(nil)).Elem()
}

func (i *virtualRouterAutoScaleConfigurationPtrType) ToVirtualRouterAutoScaleConfigurationPtrOutput() VirtualRouterAutoScaleConfigurationPtrOutput {
	return i.ToVirtualRouterAutoScaleConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualRouterAutoScaleConfigurationPtrType) ToVirtualRouterAutoScaleConfigurationPtrOutputWithContext(ctx context.Context) VirtualRouterAutoScaleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRouterAutoScaleConfigurationPtrOutput)
}

type VirtualRouterAutoScaleConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualRouterAutoScaleConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualRouterAutoScaleConfiguration)(nil)).Elem()
}

func (o VirtualRouterAutoScaleConfigurationOutput) ToVirtualRouterAutoScaleConfigurationOutput() VirtualRouterAutoScaleConfigurationOutput {
	return o
}

func (o VirtualRouterAutoScaleConfigurationOutput) ToVirtualRouterAutoScaleConfigurationOutputWithContext(ctx context.Context) VirtualRouterAutoScaleConfigurationOutput {
	return o
}

func (o VirtualRouterAutoScaleConfigurationOutput) ToVirtualRouterAutoScaleConfigurationPtrOutput() VirtualRouterAutoScaleConfigurationPtrOutput {
	return o.ToVirtualRouterAutoScaleConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualRouterAutoScaleConfigurationOutput) ToVirtualRouterAutoScaleConfigurationPtrOutputWithContext(ctx context.Context) VirtualRouterAutoScaleConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualRouterAutoScaleConfiguration) *VirtualRouterAutoScaleConfiguration {
		return &v
	}).(VirtualRouterAutoScaleConfigurationPtrOutput)
}

func (o VirtualRouterAutoScaleConfigurationOutput) MinCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualRouterAutoScaleConfiguration) *int { return v.MinCapacity }).(pulumi.IntPtrOutput)
}

type VirtualRouterAutoScaleConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualRouterAutoScaleConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRouterAutoScaleConfiguration)(nil)).Elem()
}

func (o VirtualRouterAutoScaleConfigurationPtrOutput) ToVirtualRouterAutoScaleConfigurationPtrOutput() VirtualRouterAutoScaleConfigurationPtrOutput {
	return o
}

func (o VirtualRouterAutoScaleConfigurationPtrOutput) ToVirtualRouterAutoScaleConfigurationPtrOutputWithContext(ctx context.Context) VirtualRouterAutoScaleConfigurationPtrOutput {
	return o
}

func (o VirtualRouterAutoScaleConfigurationPtrOutput) Elem() VirtualRouterAutoScaleConfigurationOutput {
	return o.ApplyT(func(v *VirtualRouterAutoScaleConfiguration) VirtualRouterAutoScaleConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualRouterAutoScaleConfiguration
		return ret
	}).(VirtualRouterAutoScaleConfigurationOutput)
}

func (o VirtualRouterAutoScaleConfigurationPtrOutput) MinCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualRouterAutoScaleConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.MinCapacity
	}).(pulumi.IntPtrOutput)
}

type VirtualRouterAutoScaleConfigurationResponse struct {
	MinCapacity *int `pulumi:"minCapacity"`
}

type VirtualRouterAutoScaleConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualRouterAutoScaleConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualRouterAutoScaleConfigurationResponse)(nil)).Elem()
}

func (o VirtualRouterAutoScaleConfigurationResponseOutput) ToVirtualRouterAutoScaleConfigurationResponseOutput() VirtualRouterAutoScaleConfigurationResponseOutput {
	return o
}

func (o VirtualRouterAutoScaleConfigurationResponseOutput) ToVirtualRouterAutoScaleConfigurationResponseOutputWithContext(ctx context.Context) VirtualRouterAutoScaleConfigurationResponseOutput {
	return o
}

func (o VirtualRouterAutoScaleConfigurationResponseOutput) MinCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualRouterAutoScaleConfigurationResponse) *int { return v.MinCapacity }).(pulumi.IntPtrOutput)
}

type VirtualRouterAutoScaleConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualRouterAutoScaleConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRouterAutoScaleConfigurationResponse)(nil)).Elem()
}

func (o VirtualRouterAutoScaleConfigurationResponsePtrOutput) ToVirtualRouterAutoScaleConfigurationResponsePtrOutput() VirtualRouterAutoScaleConfigurationResponsePtrOutput {
	return o
}

func (o VirtualRouterAutoScaleConfigurationResponsePtrOutput) ToVirtualRouterAutoScaleConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualRouterAutoScaleConfigurationResponsePtrOutput {
	return o
}

func (o VirtualRouterAutoScaleConfigurationResponsePtrOutput) Elem() VirtualRouterAutoScaleConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualRouterAutoScaleConfigurationResponse) VirtualRouterAutoScaleConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualRouterAutoScaleConfigurationResponse
		return ret
	}).(VirtualRouterAutoScaleConfigurationResponseOutput)
}

func (o VirtualRouterAutoScaleConfigurationResponsePtrOutput) MinCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualRouterAutoScaleConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinCapacity
	}).(pulumi.IntPtrOutput)
}

type VnetRoute struct {
	StaticRoutes       []StaticRoute       `pulumi:"staticRoutes"`
	StaticRoutesConfig *StaticRoutesConfig `pulumi:"staticRoutesConfig"`
}





type VnetRouteInput interface {
	pulumi.Input

	ToVnetRouteOutput() VnetRouteOutput
	ToVnetRouteOutputWithContext(context.Context) VnetRouteOutput
}

type VnetRouteArgs struct {
	StaticRoutes       StaticRouteArrayInput      `pulumi:"staticRoutes"`
	StaticRoutesConfig StaticRoutesConfigPtrInput `pulumi:"staticRoutesConfig"`
}

func (VnetRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRoute)(nil)).Elem()
}

func (i VnetRouteArgs) ToVnetRouteOutput() VnetRouteOutput {
	return i.ToVnetRouteOutputWithContext(context.Background())
}

func (i VnetRouteArgs) ToVnetRouteOutputWithContext(ctx context.Context) VnetRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRouteOutput)
}

func (i VnetRouteArgs) ToVnetRoutePtrOutput() VnetRoutePtrOutput {
	return i.ToVnetRoutePtrOutputWithContext(context.Background())
}

func (i VnetRouteArgs) ToVnetRoutePtrOutputWithContext(ctx context.Context) VnetRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRouteOutput).ToVnetRoutePtrOutputWithContext(ctx)
}









type VnetRoutePtrInput interface {
	pulumi.Input

	ToVnetRoutePtrOutput() VnetRoutePtrOutput
	ToVnetRoutePtrOutputWithContext(context.Context) VnetRoutePtrOutput
}

type vnetRoutePtrType VnetRouteArgs

func VnetRoutePtr(v *VnetRouteArgs) VnetRoutePtrInput {
	return (*vnetRoutePtrType)(v)
}

func (*vnetRoutePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VnetRoute)(nil)).Elem()
}

func (i *vnetRoutePtrType) ToVnetRoutePtrOutput() VnetRoutePtrOutput {
	return i.ToVnetRoutePtrOutputWithContext(context.Background())
}

func (i *vnetRoutePtrType) ToVnetRoutePtrOutputWithContext(ctx context.Context) VnetRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRoutePtrOutput)
}

type VnetRouteOutput struct{ *pulumi.OutputState }

func (VnetRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRoute)(nil)).Elem()
}

func (o VnetRouteOutput) ToVnetRouteOutput() VnetRouteOutput {
	return o
}

func (o VnetRouteOutput) ToVnetRouteOutputWithContext(ctx context.Context) VnetRouteOutput {
	return o
}

func (o VnetRouteOutput) ToVnetRoutePtrOutput() VnetRoutePtrOutput {
	return o.ToVnetRoutePtrOutputWithContext(context.Background())
}

func (o VnetRouteOutput) ToVnetRoutePtrOutputWithContext(ctx context.Context) VnetRoutePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VnetRoute) *VnetRoute {
		return &v
	}).(VnetRoutePtrOutput)
}

func (o VnetRouteOutput) StaticRoutes() StaticRouteArrayOutput {
	return o.ApplyT(func(v VnetRoute) []StaticRoute { return v.StaticRoutes }).(StaticRouteArrayOutput)
}

func (o VnetRouteOutput) StaticRoutesConfig() StaticRoutesConfigPtrOutput {
	return o.ApplyT(func(v VnetRoute) *StaticRoutesConfig { return v.StaticRoutesConfig }).(StaticRoutesConfigPtrOutput)
}

type VnetRoutePtrOutput struct{ *pulumi.OutputState }

func (VnetRoutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VnetRoute)(nil)).Elem()
}

func (o VnetRoutePtrOutput) ToVnetRoutePtrOutput() VnetRoutePtrOutput {
	return o
}

func (o VnetRoutePtrOutput) ToVnetRoutePtrOutputWithContext(ctx context.Context) VnetRoutePtrOutput {
	return o
}

func (o VnetRoutePtrOutput) Elem() VnetRouteOutput {
	return o.ApplyT(func(v *VnetRoute) VnetRoute {
		if v != nil {
			return *v
		}
		var ret VnetRoute
		return ret
	}).(VnetRouteOutput)
}

func (o VnetRoutePtrOutput) StaticRoutes() StaticRouteArrayOutput {
	return o.ApplyT(func(v *VnetRoute) []StaticRoute {
		if v == nil {
			return nil
		}
		return v.StaticRoutes
	}).(StaticRouteArrayOutput)
}

func (o VnetRoutePtrOutput) StaticRoutesConfig() StaticRoutesConfigPtrOutput {
	return o.ApplyT(func(v *VnetRoute) *StaticRoutesConfig {
		if v == nil {
			return nil
		}
		return v.StaticRoutesConfig
	}).(StaticRoutesConfigPtrOutput)
}

type VnetRouteResponse struct {
	BgpConnections     []SubResourceResponse       `pulumi:"bgpConnections"`
	StaticRoutes       []StaticRouteResponse       `pulumi:"staticRoutes"`
	StaticRoutesConfig *StaticRoutesConfigResponse `pulumi:"staticRoutesConfig"`
}

type VnetRouteResponseOutput struct{ *pulumi.OutputState }

func (VnetRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutput() VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutputWithContext(ctx context.Context) VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) BgpConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VnetRouteResponse) []SubResourceResponse { return v.BgpConnections }).(SubResourceResponseArrayOutput)
}

func (o VnetRouteResponseOutput) StaticRoutes() StaticRouteResponseArrayOutput {
	return o.ApplyT(func(v VnetRouteResponse) []StaticRouteResponse { return v.StaticRoutes }).(StaticRouteResponseArrayOutput)
}

func (o VnetRouteResponseOutput) StaticRoutesConfig() StaticRoutesConfigResponsePtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *StaticRoutesConfigResponse { return v.StaticRoutesConfig }).(StaticRoutesConfigResponsePtrOutput)
}

type VnetRouteResponsePtrOutput struct{ *pulumi.OutputState }

func (VnetRouteResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponsePtrOutput) ToVnetRouteResponsePtrOutput() VnetRouteResponsePtrOutput {
	return o
}

func (o VnetRouteResponsePtrOutput) ToVnetRouteResponsePtrOutputWithContext(ctx context.Context) VnetRouteResponsePtrOutput {
	return o
}

func (o VnetRouteResponsePtrOutput) Elem() VnetRouteResponseOutput {
	return o.ApplyT(func(v *VnetRouteResponse) VnetRouteResponse {
		if v != nil {
			return *v
		}
		var ret VnetRouteResponse
		return ret
	}).(VnetRouteResponseOutput)
}

func (o VnetRouteResponsePtrOutput) BgpConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VnetRouteResponse) []SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.BgpConnections
	}).(SubResourceResponseArrayOutput)
}

func (o VnetRouteResponsePtrOutput) StaticRoutes() StaticRouteResponseArrayOutput {
	return o.ApplyT(func(v *VnetRouteResponse) []StaticRouteResponse {
		if v == nil {
			return nil
		}
		return v.StaticRoutes
	}).(StaticRouteResponseArrayOutput)
}

func (o VnetRouteResponsePtrOutput) StaticRoutesConfig() StaticRoutesConfigResponsePtrOutput {
	return o.ApplyT(func(v *VnetRouteResponse) *StaticRoutesConfigResponse {
		if v == nil {
			return nil
		}
		return v.StaticRoutesConfig
	}).(StaticRoutesConfigResponsePtrOutput)
}

type VngClientConnectionConfiguration struct {
	Id                                *string       `pulumi:"id"`
	Name                              *string       `pulumi:"name"`
	VirtualNetworkGatewayPolicyGroups []SubResource `pulumi:"virtualNetworkGatewayPolicyGroups"`
	VpnClientAddressPool              AddressSpace  `pulumi:"vpnClientAddressPool"`
}





type VngClientConnectionConfigurationInput interface {
	pulumi.Input

	ToVngClientConnectionConfigurationOutput() VngClientConnectionConfigurationOutput
	ToVngClientConnectionConfigurationOutputWithContext(context.Context) VngClientConnectionConfigurationOutput
}

type VngClientConnectionConfigurationArgs struct {
	Id                                pulumi.StringPtrInput `pulumi:"id"`
	Name                              pulumi.StringPtrInput `pulumi:"name"`
	VirtualNetworkGatewayPolicyGroups SubResourceArrayInput `pulumi:"virtualNetworkGatewayPolicyGroups"`
	VpnClientAddressPool              AddressSpaceInput     `pulumi:"vpnClientAddressPool"`
}

func (VngClientConnectionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VngClientConnectionConfiguration)(nil)).Elem()
}

func (i VngClientConnectionConfigurationArgs) ToVngClientConnectionConfigurationOutput() VngClientConnectionConfigurationOutput {
	return i.ToVngClientConnectionConfigurationOutputWithContext(context.Background())
}

func (i VngClientConnectionConfigurationArgs) ToVngClientConnectionConfigurationOutputWithContext(ctx context.Context) VngClientConnectionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VngClientConnectionConfigurationOutput)
}





type VngClientConnectionConfigurationArrayInput interface {
	pulumi.Input

	ToVngClientConnectionConfigurationArrayOutput() VngClientConnectionConfigurationArrayOutput
	ToVngClientConnectionConfigurationArrayOutputWithContext(context.Context) VngClientConnectionConfigurationArrayOutput
}

type VngClientConnectionConfigurationArray []VngClientConnectionConfigurationInput

func (VngClientConnectionConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VngClientConnectionConfiguration)(nil)).Elem()
}

func (i VngClientConnectionConfigurationArray) ToVngClientConnectionConfigurationArrayOutput() VngClientConnectionConfigurationArrayOutput {
	return i.ToVngClientConnectionConfigurationArrayOutputWithContext(context.Background())
}

func (i VngClientConnectionConfigurationArray) ToVngClientConnectionConfigurationArrayOutputWithContext(ctx context.Context) VngClientConnectionConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VngClientConnectionConfigurationArrayOutput)
}

type VngClientConnectionConfigurationOutput struct{ *pulumi.OutputState }

func (VngClientConnectionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VngClientConnectionConfiguration)(nil)).Elem()
}

func (o VngClientConnectionConfigurationOutput) ToVngClientConnectionConfigurationOutput() VngClientConnectionConfigurationOutput {
	return o
}

func (o VngClientConnectionConfigurationOutput) ToVngClientConnectionConfigurationOutputWithContext(ctx context.Context) VngClientConnectionConfigurationOutput {
	return o
}

func (o VngClientConnectionConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VngClientConnectionConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VngClientConnectionConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VngClientConnectionConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VngClientConnectionConfigurationOutput) VirtualNetworkGatewayPolicyGroups() SubResourceArrayOutput {
	return o.ApplyT(func(v VngClientConnectionConfiguration) []SubResource { return v.VirtualNetworkGatewayPolicyGroups }).(SubResourceArrayOutput)
}

func (o VngClientConnectionConfigurationOutput) VpnClientAddressPool() AddressSpaceOutput {
	return o.ApplyT(func(v VngClientConnectionConfiguration) AddressSpace { return v.VpnClientAddressPool }).(AddressSpaceOutput)
}

type VngClientConnectionConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VngClientConnectionConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VngClientConnectionConfiguration)(nil)).Elem()
}

func (o VngClientConnectionConfigurationArrayOutput) ToVngClientConnectionConfigurationArrayOutput() VngClientConnectionConfigurationArrayOutput {
	return o
}

func (o VngClientConnectionConfigurationArrayOutput) ToVngClientConnectionConfigurationArrayOutputWithContext(ctx context.Context) VngClientConnectionConfigurationArrayOutput {
	return o
}

func (o VngClientConnectionConfigurationArrayOutput) Index(i pulumi.IntInput) VngClientConnectionConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VngClientConnectionConfiguration {
		return vs[0].([]VngClientConnectionConfiguration)[vs[1].(int)]
	}).(VngClientConnectionConfigurationOutput)
}

type VngClientConnectionConfigurationResponse struct {
	Etag                              string                `pulumi:"etag"`
	Id                                *string               `pulumi:"id"`
	Name                              *string               `pulumi:"name"`
	ProvisioningState                 string                `pulumi:"provisioningState"`
	VirtualNetworkGatewayPolicyGroups []SubResourceResponse `pulumi:"virtualNetworkGatewayPolicyGroups"`
	VpnClientAddressPool              AddressSpaceResponse  `pulumi:"vpnClientAddressPool"`
}

type VngClientConnectionConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VngClientConnectionConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VngClientConnectionConfigurationResponse)(nil)).Elem()
}

func (o VngClientConnectionConfigurationResponseOutput) ToVngClientConnectionConfigurationResponseOutput() VngClientConnectionConfigurationResponseOutput {
	return o
}

func (o VngClientConnectionConfigurationResponseOutput) ToVngClientConnectionConfigurationResponseOutputWithContext(ctx context.Context) VngClientConnectionConfigurationResponseOutput {
	return o
}

func (o VngClientConnectionConfigurationResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VngClientConnectionConfigurationResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VngClientConnectionConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VngClientConnectionConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VngClientConnectionConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VngClientConnectionConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VngClientConnectionConfigurationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VngClientConnectionConfigurationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VngClientConnectionConfigurationResponseOutput) VirtualNetworkGatewayPolicyGroups() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VngClientConnectionConfigurationResponse) []SubResourceResponse {
		return v.VirtualNetworkGatewayPolicyGroups
	}).(SubResourceResponseArrayOutput)
}

func (o VngClientConnectionConfigurationResponseOutput) VpnClientAddressPool() AddressSpaceResponseOutput {
	return o.ApplyT(func(v VngClientConnectionConfigurationResponse) AddressSpaceResponse { return v.VpnClientAddressPool }).(AddressSpaceResponseOutput)
}

type VngClientConnectionConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VngClientConnectionConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VngClientConnectionConfigurationResponse)(nil)).Elem()
}

func (o VngClientConnectionConfigurationResponseArrayOutput) ToVngClientConnectionConfigurationResponseArrayOutput() VngClientConnectionConfigurationResponseArrayOutput {
	return o
}

func (o VngClientConnectionConfigurationResponseArrayOutput) ToVngClientConnectionConfigurationResponseArrayOutputWithContext(ctx context.Context) VngClientConnectionConfigurationResponseArrayOutput {
	return o
}

func (o VngClientConnectionConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VngClientConnectionConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VngClientConnectionConfigurationResponse {
		return vs[0].([]VngClientConnectionConfigurationResponse)[vs[1].(int)]
	}).(VngClientConnectionConfigurationResponseOutput)
}

type VpnClientConfiguration struct {
	AadAudience                       *string                            `pulumi:"aadAudience"`
	AadIssuer                         *string                            `pulumi:"aadIssuer"`
	AadTenant                         *string                            `pulumi:"aadTenant"`
	RadiusServerAddress               *string                            `pulumi:"radiusServerAddress"`
	RadiusServerSecret                *string                            `pulumi:"radiusServerSecret"`
	RadiusServers                     []RadiusServer                     `pulumi:"radiusServers"`
	VngClientConnectionConfigurations []VngClientConnectionConfiguration `pulumi:"vngClientConnectionConfigurations"`
	VpnAuthenticationTypes            []string                           `pulumi:"vpnAuthenticationTypes"`
	VpnClientAddressPool              *AddressSpace                      `pulumi:"vpnClientAddressPool"`
	VpnClientIpsecPolicies            []IpsecPolicy                      `pulumi:"vpnClientIpsecPolicies"`
	VpnClientProtocols                []string                           `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates      []VpnClientRevokedCertificate      `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates         []VpnClientRootCertificate         `pulumi:"vpnClientRootCertificates"`
}





type VpnClientConfigurationInput interface {
	pulumi.Input

	ToVpnClientConfigurationOutput() VpnClientConfigurationOutput
	ToVpnClientConfigurationOutputWithContext(context.Context) VpnClientConfigurationOutput
}

type VpnClientConfigurationArgs struct {
	AadAudience                       pulumi.StringPtrInput                      `pulumi:"aadAudience"`
	AadIssuer                         pulumi.StringPtrInput                      `pulumi:"aadIssuer"`
	AadTenant                         pulumi.StringPtrInput                      `pulumi:"aadTenant"`
	RadiusServerAddress               pulumi.StringPtrInput                      `pulumi:"radiusServerAddress"`
	RadiusServerSecret                pulumi.StringPtrInput                      `pulumi:"radiusServerSecret"`
	RadiusServers                     RadiusServerArrayInput                     `pulumi:"radiusServers"`
	VngClientConnectionConfigurations VngClientConnectionConfigurationArrayInput `pulumi:"vngClientConnectionConfigurations"`
	VpnAuthenticationTypes            pulumi.StringArrayInput                    `pulumi:"vpnAuthenticationTypes"`
	VpnClientAddressPool              AddressSpacePtrInput                       `pulumi:"vpnClientAddressPool"`
	VpnClientIpsecPolicies            IpsecPolicyArrayInput                      `pulumi:"vpnClientIpsecPolicies"`
	VpnClientProtocols                pulumi.StringArrayInput                    `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates      VpnClientRevokedCertificateArrayInput      `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates         VpnClientRootCertificateArrayInput         `pulumi:"vpnClientRootCertificates"`
}

func (VpnClientConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfiguration)(nil)).Elem()
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationOutput() VpnClientConfigurationOutput {
	return i.ToVpnClientConfigurationOutputWithContext(context.Background())
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationOutputWithContext(ctx context.Context) VpnClientConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationOutput)
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return i.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationOutput).ToVpnClientConfigurationPtrOutputWithContext(ctx)
}









type VpnClientConfigurationPtrInput interface {
	pulumi.Input

	ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput
	ToVpnClientConfigurationPtrOutputWithContext(context.Context) VpnClientConfigurationPtrOutput
}

type vpnClientConfigurationPtrType VpnClientConfigurationArgs

func VpnClientConfigurationPtr(v *VpnClientConfigurationArgs) VpnClientConfigurationPtrInput {
	return (*vpnClientConfigurationPtrType)(v)
}

func (*vpnClientConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfiguration)(nil)).Elem()
}

func (i *vpnClientConfigurationPtrType) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return i.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (i *vpnClientConfigurationPtrType) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationPtrOutput)
}

type VpnClientConfigurationOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfiguration)(nil)).Elem()
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationOutput() VpnClientConfigurationOutput {
	return o
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationOutputWithContext(ctx context.Context) VpnClientConfigurationOutput {
	return o
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return o.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnClientConfiguration) *VpnClientConfiguration {
		return &v
	}).(VpnClientConfigurationPtrOutput)
}

func (o VpnClientConfigurationOutput) AadAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.AadAudience }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) AadIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.AadIssuer }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) AadTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.AadTenant }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *string { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationOutput) RadiusServers() RadiusServerArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []RadiusServer { return v.RadiusServers }).(RadiusServerArrayOutput)
}

func (o VpnClientConfigurationOutput) VngClientConnectionConfigurations() VngClientConnectionConfigurationArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []VngClientConnectionConfiguration {
		return v.VngClientConnectionConfigurations
	}).(VngClientConnectionConfigurationArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []string { return v.VpnAuthenticationTypes }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientAddressPool() AddressSpacePtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *AddressSpace { return v.VpnClientAddressPool }).(AddressSpacePtrOutput)
}

func (o VpnClientConfigurationOutput) VpnClientIpsecPolicies() IpsecPolicyArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []IpsecPolicy { return v.VpnClientIpsecPolicies }).(IpsecPolicyArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []string { return v.VpnClientProtocols }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []VpnClientRevokedCertificate { return v.VpnClientRevokedCertificates }).(VpnClientRevokedCertificateArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientRootCertificates() VpnClientRootCertificateArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []VpnClientRootCertificate { return v.VpnClientRootCertificates }).(VpnClientRootCertificateArrayOutput)
}

type VpnClientConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfiguration)(nil)).Elem()
}

func (o VpnClientConfigurationPtrOutput) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return o
}

func (o VpnClientConfigurationPtrOutput) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return o
}

func (o VpnClientConfigurationPtrOutput) Elem() VpnClientConfigurationOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) VpnClientConfiguration {
		if v != nil {
			return *v
		}
		var ret VpnClientConfiguration
		return ret
	}).(VpnClientConfigurationOutput)
}

func (o VpnClientConfigurationPtrOutput) AadAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AadAudience
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) AadIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AadIssuer
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) AadTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AadTenant
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerAddress
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerSecret
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationPtrOutput) RadiusServers() RadiusServerArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []RadiusServer {
		if v == nil {
			return nil
		}
		return v.RadiusServers
	}).(RadiusServerArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VngClientConnectionConfigurations() VngClientConnectionConfigurationArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []VngClientConnectionConfiguration {
		if v == nil {
			return nil
		}
		return v.VngClientConnectionConfigurations
	}).(VngClientConnectionConfigurationArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.VpnAuthenticationTypes
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientAddressPool() AddressSpacePtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *AddressSpace {
		if v == nil {
			return nil
		}
		return v.VpnClientAddressPool
	}).(AddressSpacePtrOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientIpsecPolicies() IpsecPolicyArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []IpsecPolicy {
		if v == nil {
			return nil
		}
		return v.VpnClientIpsecPolicies
	}).(IpsecPolicyArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.VpnClientProtocols
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []VpnClientRevokedCertificate {
		if v == nil {
			return nil
		}
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientRootCertificates() VpnClientRootCertificateArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []VpnClientRootCertificate {
		if v == nil {
			return nil
		}
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateArrayOutput)
}

type VpnClientConfigurationResponse struct {
	AadAudience                       *string                                    `pulumi:"aadAudience"`
	AadIssuer                         *string                                    `pulumi:"aadIssuer"`
	AadTenant                         *string                                    `pulumi:"aadTenant"`
	RadiusServerAddress               *string                                    `pulumi:"radiusServerAddress"`
	RadiusServerSecret                *string                                    `pulumi:"radiusServerSecret"`
	RadiusServers                     []RadiusServerResponse                     `pulumi:"radiusServers"`
	VngClientConnectionConfigurations []VngClientConnectionConfigurationResponse `pulumi:"vngClientConnectionConfigurations"`
	VpnAuthenticationTypes            []string                                   `pulumi:"vpnAuthenticationTypes"`
	VpnClientAddressPool              *AddressSpaceResponse                      `pulumi:"vpnClientAddressPool"`
	VpnClientIpsecPolicies            []IpsecPolicyResponse                      `pulumi:"vpnClientIpsecPolicies"`
	VpnClientProtocols                []string                                   `pulumi:"vpnClientProtocols"`
	VpnClientRevokedCertificates      []VpnClientRevokedCertificateResponse      `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates         []VpnClientRootCertificateResponse         `pulumi:"vpnClientRootCertificates"`
}

type VpnClientConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfigurationResponse)(nil)).Elem()
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponseOutput() VpnClientConfigurationResponseOutput {
	return o
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponseOutputWithContext(ctx context.Context) VpnClientConfigurationResponseOutput {
	return o
}

func (o VpnClientConfigurationResponseOutput) AadAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.AadAudience }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) AadIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.AadIssuer }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) AadTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.AadTenant }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *string { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponseOutput) RadiusServers() RadiusServerResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []RadiusServerResponse { return v.RadiusServers }).(RadiusServerResponseArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VngClientConnectionConfigurations() VngClientConnectionConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []VngClientConnectionConfigurationResponse {
		return v.VngClientConnectionConfigurations
	}).(VngClientConnectionConfigurationResponseArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []string { return v.VpnAuthenticationTypes }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientAddressPool() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *AddressSpaceResponse { return v.VpnClientAddressPool }).(AddressSpaceResponsePtrOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []IpsecPolicyResponse { return v.VpnClientIpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []string { return v.VpnClientProtocols }).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []VpnClientRevokedCertificateResponse {
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateResponseArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientRootCertificates() VpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []VpnClientRootCertificateResponse {
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateResponseArrayOutput)
}

type VpnClientConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfigurationResponse)(nil)).Elem()
}

func (o VpnClientConfigurationResponsePtrOutput) ToVpnClientConfigurationResponsePtrOutput() VpnClientConfigurationResponsePtrOutput {
	return o
}

func (o VpnClientConfigurationResponsePtrOutput) ToVpnClientConfigurationResponsePtrOutputWithContext(ctx context.Context) VpnClientConfigurationResponsePtrOutput {
	return o
}

func (o VpnClientConfigurationResponsePtrOutput) Elem() VpnClientConfigurationResponseOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) VpnClientConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VpnClientConfigurationResponse
		return ret
	}).(VpnClientConfigurationResponseOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) AadAudience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadAudience
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) AadIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadIssuer
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) AadTenant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadTenant
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerAddress
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RadiusServerSecret
	}).(pulumi.StringPtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) RadiusServers() RadiusServerResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []RadiusServerResponse {
		if v == nil {
			return nil
		}
		return v.RadiusServers
	}).(RadiusServerResponseArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VngClientConnectionConfigurations() VngClientConnectionConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []VngClientConnectionConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.VngClientConnectionConfigurations
	}).(VngClientConnectionConfigurationResponseArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnAuthenticationTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.VpnAuthenticationTypes
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientAddressPool() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *AddressSpaceResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientAddressPool
	}).(AddressSpaceResponsePtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []IpsecPolicyResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientIpsecPolicies
	}).(IpsecPolicyResponseArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.VpnClientProtocols
	}).(pulumi.StringArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []VpnClientRevokedCertificateResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateResponseArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientRootCertificates() VpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []VpnClientRootCertificateResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateResponseArrayOutput)
}

type VpnClientConnectionHealthDetailResponse struct {
	EgressBytesTransferred    float64 `pulumi:"egressBytesTransferred"`
	EgressPacketsTransferred  float64 `pulumi:"egressPacketsTransferred"`
	IngressBytesTransferred   float64 `pulumi:"ingressBytesTransferred"`
	IngressPacketsTransferred float64 `pulumi:"ingressPacketsTransferred"`
	MaxBandwidth              float64 `pulumi:"maxBandwidth"`
	MaxPacketsPerSecond       float64 `pulumi:"maxPacketsPerSecond"`
	PrivateIpAddress          string  `pulumi:"privateIpAddress"`
	PublicIpAddress           string  `pulumi:"publicIpAddress"`
	VpnConnectionDuration     float64 `pulumi:"vpnConnectionDuration"`
	VpnConnectionId           string  `pulumi:"vpnConnectionId"`
	VpnConnectionTime         string  `pulumi:"vpnConnectionTime"`
	VpnUserName               string  `pulumi:"vpnUserName"`
}

type VpnClientConnectionHealthDetailResponseOutput struct{ *pulumi.OutputState }

func (VpnClientConnectionHealthDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConnectionHealthDetailResponse)(nil)).Elem()
}

func (o VpnClientConnectionHealthDetailResponseOutput) ToVpnClientConnectionHealthDetailResponseOutput() VpnClientConnectionHealthDetailResponseOutput {
	return o
}

func (o VpnClientConnectionHealthDetailResponseOutput) ToVpnClientConnectionHealthDetailResponseOutputWithContext(ctx context.Context) VpnClientConnectionHealthDetailResponseOutput {
	return o
}

func (o VpnClientConnectionHealthDetailResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) EgressPacketsTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.EgressPacketsTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) IngressPacketsTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.IngressPacketsTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) MaxBandwidth() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.MaxBandwidth }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) MaxPacketsPerSecond() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.MaxPacketsPerSecond }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

func (o VpnClientConnectionHealthDetailResponseOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

func (o VpnClientConnectionHealthDetailResponseOutput) VpnConnectionDuration() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) float64 { return v.VpnConnectionDuration }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthDetailResponseOutput) VpnConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.VpnConnectionId }).(pulumi.StringOutput)
}

func (o VpnClientConnectionHealthDetailResponseOutput) VpnConnectionTime() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.VpnConnectionTime }).(pulumi.StringOutput)
}

func (o VpnClientConnectionHealthDetailResponseOutput) VpnUserName() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthDetailResponse) string { return v.VpnUserName }).(pulumi.StringOutput)
}

type VpnClientConnectionHealthDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientConnectionHealthDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientConnectionHealthDetailResponse)(nil)).Elem()
}

func (o VpnClientConnectionHealthDetailResponseArrayOutput) ToVpnClientConnectionHealthDetailResponseArrayOutput() VpnClientConnectionHealthDetailResponseArrayOutput {
	return o
}

func (o VpnClientConnectionHealthDetailResponseArrayOutput) ToVpnClientConnectionHealthDetailResponseArrayOutputWithContext(ctx context.Context) VpnClientConnectionHealthDetailResponseArrayOutput {
	return o
}

func (o VpnClientConnectionHealthDetailResponseArrayOutput) Index(i pulumi.IntInput) VpnClientConnectionHealthDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientConnectionHealthDetailResponse {
		return vs[0].([]VpnClientConnectionHealthDetailResponse)[vs[1].(int)]
	}).(VpnClientConnectionHealthDetailResponseOutput)
}

type VpnClientConnectionHealthResponse struct {
	AllocatedIpAddresses         []string `pulumi:"allocatedIpAddresses"`
	TotalEgressBytesTransferred  float64  `pulumi:"totalEgressBytesTransferred"`
	TotalIngressBytesTransferred float64  `pulumi:"totalIngressBytesTransferred"`
	VpnClientConnectionsCount    *int     `pulumi:"vpnClientConnectionsCount"`
}

type VpnClientConnectionHealthResponseOutput struct{ *pulumi.OutputState }

func (VpnClientConnectionHealthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConnectionHealthResponse)(nil)).Elem()
}

func (o VpnClientConnectionHealthResponseOutput) ToVpnClientConnectionHealthResponseOutput() VpnClientConnectionHealthResponseOutput {
	return o
}

func (o VpnClientConnectionHealthResponseOutput) ToVpnClientConnectionHealthResponseOutputWithContext(ctx context.Context) VpnClientConnectionHealthResponseOutput {
	return o
}

func (o VpnClientConnectionHealthResponseOutput) AllocatedIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthResponse) []string { return v.AllocatedIpAddresses }).(pulumi.StringArrayOutput)
}

func (o VpnClientConnectionHealthResponseOutput) TotalEgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthResponse) float64 { return v.TotalEgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthResponseOutput) TotalIngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnClientConnectionHealthResponse) float64 { return v.TotalIngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnClientConnectionHealthResponseOutput) VpnClientConnectionsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnClientConnectionHealthResponse) *int { return v.VpnClientConnectionsCount }).(pulumi.IntPtrOutput)
}

type VpnClientRevokedCertificate struct {
	Id         *string `pulumi:"id"`
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnClientRevokedCertificateInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput
	ToVpnClientRevokedCertificateOutputWithContext(context.Context) VpnClientRevokedCertificateOutput
}

type VpnClientRevokedCertificateArgs struct {
	Id         pulumi.StringPtrInput `pulumi:"id"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnClientRevokedCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnClientRevokedCertificateArgs) ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput {
	return i.ToVpnClientRevokedCertificateOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateArgs) ToVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnClientRevokedCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateOutput)
}





type VpnClientRevokedCertificateArrayInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput
	ToVpnClientRevokedCertificateArrayOutputWithContext(context.Context) VpnClientRevokedCertificateArrayOutput
}

type VpnClientRevokedCertificateArray []VpnClientRevokedCertificateInput

func (VpnClientRevokedCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnClientRevokedCertificateArray) ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput {
	return i.ToVpnClientRevokedCertificateArrayOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateArray) ToVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateArrayOutput)
}

type VpnClientRevokedCertificateOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnClientRevokedCertificateOutput) ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput {
	return o
}

func (o VpnClientRevokedCertificateOutput) ToVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnClientRevokedCertificateOutput {
	return o
}

func (o VpnClientRevokedCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnClientRevokedCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnClientRevokedCertificateArrayOutput) ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateArrayOutput) ToVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateArrayOutput) Index(i pulumi.IntInput) VpnClientRevokedCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRevokedCertificate {
		return vs[0].([]VpnClientRevokedCertificate)[vs[1].(int)]
	}).(VpnClientRevokedCertificateOutput)
}

type VpnClientRevokedCertificateResponse struct {
	Etag              string  `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Thumbprint        *string `pulumi:"thumbprint"`
}

type VpnClientRevokedCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnClientRevokedCertificateResponseOutput) ToVpnClientRevokedCertificateResponseOutput() VpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseOutput) ToVpnClientRevokedCertificateResponseOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnClientRevokedCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnClientRevokedCertificateResponseArrayOutput) ToVpnClientRevokedCertificateResponseArrayOutput() VpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseArrayOutput) ToVpnClientRevokedCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnClientRevokedCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRevokedCertificateResponse {
		return vs[0].([]VpnClientRevokedCertificateResponse)[vs[1].(int)]
	}).(VpnClientRevokedCertificateResponseOutput)
}

type VpnClientRootCertificate struct {
	Id             *string `pulumi:"id"`
	Name           *string `pulumi:"name"`
	PublicCertData string  `pulumi:"publicCertData"`
}





type VpnClientRootCertificateInput interface {
	pulumi.Input

	ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput
	ToVpnClientRootCertificateOutputWithContext(context.Context) VpnClientRootCertificateOutput
}

type VpnClientRootCertificateArgs struct {
	Id             pulumi.StringPtrInput `pulumi:"id"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringInput    `pulumi:"publicCertData"`
}

func (VpnClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificate)(nil)).Elem()
}

func (i VpnClientRootCertificateArgs) ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput {
	return i.ToVpnClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateArgs) ToVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateOutput)
}





type VpnClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput
	ToVpnClientRootCertificateArrayOutputWithContext(context.Context) VpnClientRootCertificateArrayOutput
}

type VpnClientRootCertificateArray []VpnClientRootCertificateInput

func (VpnClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificate)(nil)).Elem()
}

func (i VpnClientRootCertificateArray) ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput {
	return i.ToVpnClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateArray) ToVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateArrayOutput)
}

type VpnClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificate)(nil)).Elem()
}

func (o VpnClientRootCertificateOutput) ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput {
	return o
}

func (o VpnClientRootCertificateOutput) ToVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnClientRootCertificateOutput {
	return o
}

func (o VpnClientRootCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) PublicCertData() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) string { return v.PublicCertData }).(pulumi.StringOutput)
}

type VpnClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificate)(nil)).Elem()
}

func (o VpnClientRootCertificateArrayOutput) ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnClientRootCertificateArrayOutput) ToVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRootCertificate {
		return vs[0].([]VpnClientRootCertificate)[vs[1].(int)]
	}).(VpnClientRootCertificateOutput)
}

type VpnClientRootCertificateResponse struct {
	Etag              string  `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	PublicCertData    string  `pulumi:"publicCertData"`
}

type VpnClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnClientRootCertificateResponseOutput) ToVpnClientRootCertificateResponseOutput() VpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnClientRootCertificateResponseOutput) ToVpnClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnClientRootCertificateResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnClientRootCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnClientRootCertificateResponseOutput) PublicCertData() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.PublicCertData }).(pulumi.StringOutput)
}

type VpnClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnClientRootCertificateResponseArrayOutput) ToVpnClientRootCertificateResponseArrayOutput() VpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRootCertificateResponseArrayOutput) ToVpnClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRootCertificateResponse {
		return vs[0].([]VpnClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnClientRootCertificateResponseOutput)
}

type VpnConnectionType struct {
	ConnectionBandwidth            *int                    `pulumi:"connectionBandwidth"`
	DpdTimeoutSeconds              *int                    `pulumi:"dpdTimeoutSeconds"`
	EnableBgp                      *bool                   `pulumi:"enableBgp"`
	EnableInternetSecurity         *bool                   `pulumi:"enableInternetSecurity"`
	EnableRateLimiting             *bool                   `pulumi:"enableRateLimiting"`
	Id                             *string                 `pulumi:"id"`
	IpsecPolicies                  []IpsecPolicy           `pulumi:"ipsecPolicies"`
	Name                           *string                 `pulumi:"name"`
	RemoteVpnSite                  *SubResource            `pulumi:"remoteVpnSite"`
	RoutingConfiguration           *RoutingConfiguration   `pulumi:"routingConfiguration"`
	RoutingWeight                  *int                    `pulumi:"routingWeight"`
	SharedKey                      *string                 `pulumi:"sharedKey"`
	TrafficSelectorPolicies        []TrafficSelectorPolicy `pulumi:"trafficSelectorPolicies"`
	UseLocalAzureIpAddress         *bool                   `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                   `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string                 `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnections             []VpnSiteLinkConnection `pulumi:"vpnLinkConnections"`
}





type VpnConnectionTypeInput interface {
	pulumi.Input

	ToVpnConnectionTypeOutput() VpnConnectionTypeOutput
	ToVpnConnectionTypeOutputWithContext(context.Context) VpnConnectionTypeOutput
}

type VpnConnectionTypeArgs struct {
	ConnectionBandwidth            pulumi.IntPtrInput              `pulumi:"connectionBandwidth"`
	DpdTimeoutSeconds              pulumi.IntPtrInput              `pulumi:"dpdTimeoutSeconds"`
	EnableBgp                      pulumi.BoolPtrInput             `pulumi:"enableBgp"`
	EnableInternetSecurity         pulumi.BoolPtrInput             `pulumi:"enableInternetSecurity"`
	EnableRateLimiting             pulumi.BoolPtrInput             `pulumi:"enableRateLimiting"`
	Id                             pulumi.StringPtrInput           `pulumi:"id"`
	IpsecPolicies                  IpsecPolicyArrayInput           `pulumi:"ipsecPolicies"`
	Name                           pulumi.StringPtrInput           `pulumi:"name"`
	RemoteVpnSite                  SubResourcePtrInput             `pulumi:"remoteVpnSite"`
	RoutingConfiguration           RoutingConfigurationPtrInput    `pulumi:"routingConfiguration"`
	RoutingWeight                  pulumi.IntPtrInput              `pulumi:"routingWeight"`
	SharedKey                      pulumi.StringPtrInput           `pulumi:"sharedKey"`
	TrafficSelectorPolicies        TrafficSelectorPolicyArrayInput `pulumi:"trafficSelectorPolicies"`
	UseLocalAzureIpAddress         pulumi.BoolPtrInput             `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrInput             `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      pulumi.StringPtrInput           `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnections             VpnSiteLinkConnectionArrayInput `pulumi:"vpnLinkConnections"`
}

func (VpnConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnectionType)(nil)).Elem()
}

func (i VpnConnectionTypeArgs) ToVpnConnectionTypeOutput() VpnConnectionTypeOutput {
	return i.ToVpnConnectionTypeOutputWithContext(context.Background())
}

func (i VpnConnectionTypeArgs) ToVpnConnectionTypeOutputWithContext(ctx context.Context) VpnConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionTypeOutput)
}





type VpnConnectionTypeArrayInput interface {
	pulumi.Input

	ToVpnConnectionTypeArrayOutput() VpnConnectionTypeArrayOutput
	ToVpnConnectionTypeArrayOutputWithContext(context.Context) VpnConnectionTypeArrayOutput
}

type VpnConnectionTypeArray []VpnConnectionTypeInput

func (VpnConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnConnectionType)(nil)).Elem()
}

func (i VpnConnectionTypeArray) ToVpnConnectionTypeArrayOutput() VpnConnectionTypeArrayOutput {
	return i.ToVpnConnectionTypeArrayOutputWithContext(context.Background())
}

func (i VpnConnectionTypeArray) ToVpnConnectionTypeArrayOutputWithContext(ctx context.Context) VpnConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionTypeArrayOutput)
}

type VpnConnectionTypeOutput struct{ *pulumi.OutputState }

func (VpnConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnectionType)(nil)).Elem()
}

func (o VpnConnectionTypeOutput) ToVpnConnectionTypeOutput() VpnConnectionTypeOutput {
	return o
}

func (o VpnConnectionTypeOutput) ToVpnConnectionTypeOutputWithContext(ctx context.Context) VpnConnectionTypeOutput {
	return o
}

func (o VpnConnectionTypeOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionTypeOutput) DpdTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *int { return v.DpdTimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionTypeOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionTypeOutput) IpsecPolicies() IpsecPolicyArrayOutput {
	return o.ApplyT(func(v VpnConnectionType) []IpsecPolicy { return v.IpsecPolicies }).(IpsecPolicyArrayOutput)
}

func (o VpnConnectionTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionTypeOutput) RemoteVpnSite() SubResourcePtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *SubResource { return v.RemoteVpnSite }).(SubResourcePtrOutput)
}

func (o VpnConnectionTypeOutput) RoutingConfiguration() RoutingConfigurationPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *RoutingConfiguration { return v.RoutingConfiguration }).(RoutingConfigurationPtrOutput)
}

func (o VpnConnectionTypeOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionTypeOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionTypeOutput) TrafficSelectorPolicies() TrafficSelectorPolicyArrayOutput {
	return o.ApplyT(func(v VpnConnectionType) []TrafficSelectorPolicy { return v.TrafficSelectorPolicies }).(TrafficSelectorPolicyArrayOutput)
}

func (o VpnConnectionTypeOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionTypeOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionType) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionTypeOutput) VpnLinkConnections() VpnSiteLinkConnectionArrayOutput {
	return o.ApplyT(func(v VpnConnectionType) []VpnSiteLinkConnection { return v.VpnLinkConnections }).(VpnSiteLinkConnectionArrayOutput)
}

type VpnConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (VpnConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnConnectionType)(nil)).Elem()
}

func (o VpnConnectionTypeArrayOutput) ToVpnConnectionTypeArrayOutput() VpnConnectionTypeArrayOutput {
	return o
}

func (o VpnConnectionTypeArrayOutput) ToVpnConnectionTypeArrayOutputWithContext(ctx context.Context) VpnConnectionTypeArrayOutput {
	return o
}

func (o VpnConnectionTypeArrayOutput) Index(i pulumi.IntInput) VpnConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnConnectionType {
		return vs[0].([]VpnConnectionType)[vs[1].(int)]
	}).(VpnConnectionTypeOutput)
}

type VpnConnectionResponse struct {
	ConnectionBandwidth            *int                            `pulumi:"connectionBandwidth"`
	ConnectionStatus               string                          `pulumi:"connectionStatus"`
	DpdTimeoutSeconds              *int                            `pulumi:"dpdTimeoutSeconds"`
	EgressBytesTransferred         float64                         `pulumi:"egressBytesTransferred"`
	EnableBgp                      *bool                           `pulumi:"enableBgp"`
	EnableInternetSecurity         *bool                           `pulumi:"enableInternetSecurity"`
	EnableRateLimiting             *bool                           `pulumi:"enableRateLimiting"`
	Etag                           string                          `pulumi:"etag"`
	Id                             *string                         `pulumi:"id"`
	IngressBytesTransferred        float64                         `pulumi:"ingressBytesTransferred"`
	IpsecPolicies                  []IpsecPolicyResponse           `pulumi:"ipsecPolicies"`
	Name                           *string                         `pulumi:"name"`
	ProvisioningState              string                          `pulumi:"provisioningState"`
	RemoteVpnSite                  *SubResourceResponse            `pulumi:"remoteVpnSite"`
	RoutingConfiguration           *RoutingConfigurationResponse   `pulumi:"routingConfiguration"`
	RoutingWeight                  *int                            `pulumi:"routingWeight"`
	SharedKey                      *string                         `pulumi:"sharedKey"`
	TrafficSelectorPolicies        []TrafficSelectorPolicyResponse `pulumi:"trafficSelectorPolicies"`
	UseLocalAzureIpAddress         *bool                           `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                           `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string                         `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnections             []VpnSiteLinkConnectionResponse `pulumi:"vpnLinkConnections"`
}

type VpnConnectionResponseOutput struct{ *pulumi.OutputState }

func (VpnConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnectionResponse)(nil)).Elem()
}

func (o VpnConnectionResponseOutput) ToVpnConnectionResponseOutput() VpnConnectionResponseOutput {
	return o
}

func (o VpnConnectionResponseOutput) ToVpnConnectionResponseOutputWithContext(ctx context.Context) VpnConnectionResponseOutput {
	return o
}

func (o VpnConnectionResponseOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionResponseOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v VpnConnectionResponse) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o VpnConnectionResponseOutput) DpdTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *int { return v.DpdTimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnConnectionResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnConnectionResponseOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnConnectionResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnConnectionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnConnectionResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnConnectionResponseOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v VpnConnectionResponse) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VpnConnectionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnConnectionResponseOutput) RemoteVpnSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *SubResourceResponse { return v.RemoteVpnSite }).(SubResourceResponsePtrOutput)
}

func (o VpnConnectionResponseOutput) RoutingConfiguration() RoutingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *RoutingConfigurationResponse { return v.RoutingConfiguration }).(RoutingConfigurationResponsePtrOutput)
}

func (o VpnConnectionResponseOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnConnectionResponseOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionResponseOutput) TrafficSelectorPolicies() TrafficSelectorPolicyResponseArrayOutput {
	return o.ApplyT(func(v VpnConnectionResponse) []TrafficSelectorPolicyResponse { return v.TrafficSelectorPolicies }).(TrafficSelectorPolicyResponseArrayOutput)
}

func (o VpnConnectionResponseOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnConnectionResponseOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnConnectionResponse) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionResponseOutput) VpnLinkConnections() VpnSiteLinkConnectionResponseArrayOutput {
	return o.ApplyT(func(v VpnConnectionResponse) []VpnSiteLinkConnectionResponse { return v.VpnLinkConnections }).(VpnSiteLinkConnectionResponseArrayOutput)
}

type VpnConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnConnectionResponse)(nil)).Elem()
}

func (o VpnConnectionResponseArrayOutput) ToVpnConnectionResponseArrayOutput() VpnConnectionResponseArrayOutput {
	return o
}

func (o VpnConnectionResponseArrayOutput) ToVpnConnectionResponseArrayOutputWithContext(ctx context.Context) VpnConnectionResponseArrayOutput {
	return o
}

func (o VpnConnectionResponseArrayOutput) Index(i pulumi.IntInput) VpnConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnConnectionResponse {
		return vs[0].([]VpnConnectionResponse)[vs[1].(int)]
	}).(VpnConnectionResponseOutput)
}

type VpnGatewayIpConfigurationResponse struct {
	Id               *string `pulumi:"id"`
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	PublicIpAddress  *string `pulumi:"publicIpAddress"`
}

type VpnGatewayIpConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VpnGatewayIpConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayIpConfigurationResponse)(nil)).Elem()
}

func (o VpnGatewayIpConfigurationResponseOutput) ToVpnGatewayIpConfigurationResponseOutput() VpnGatewayIpConfigurationResponseOutput {
	return o
}

func (o VpnGatewayIpConfigurationResponseOutput) ToVpnGatewayIpConfigurationResponseOutputWithContext(ctx context.Context) VpnGatewayIpConfigurationResponseOutput {
	return o
}

func (o VpnGatewayIpConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayIpConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayIpConfigurationResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayIpConfigurationResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayIpConfigurationResponseOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayIpConfigurationResponse) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

type VpnGatewayIpConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnGatewayIpConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGatewayIpConfigurationResponse)(nil)).Elem()
}

func (o VpnGatewayIpConfigurationResponseArrayOutput) ToVpnGatewayIpConfigurationResponseArrayOutput() VpnGatewayIpConfigurationResponseArrayOutput {
	return o
}

func (o VpnGatewayIpConfigurationResponseArrayOutput) ToVpnGatewayIpConfigurationResponseArrayOutputWithContext(ctx context.Context) VpnGatewayIpConfigurationResponseArrayOutput {
	return o
}

func (o VpnGatewayIpConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VpnGatewayIpConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnGatewayIpConfigurationResponse {
		return vs[0].([]VpnGatewayIpConfigurationResponse)[vs[1].(int)]
	}).(VpnGatewayIpConfigurationResponseOutput)
}

type VpnGatewayNatRule struct {
	ExternalMappings  []VpnNatRuleMapping `pulumi:"externalMappings"`
	Id                *string             `pulumi:"id"`
	InternalMappings  []VpnNatRuleMapping `pulumi:"internalMappings"`
	IpConfigurationId *string             `pulumi:"ipConfigurationId"`
	Mode              *string             `pulumi:"mode"`
	Name              *string             `pulumi:"name"`
	Type              *string             `pulumi:"type"`
}





type VpnGatewayNatRuleInput interface {
	pulumi.Input

	ToVpnGatewayNatRuleOutput() VpnGatewayNatRuleOutput
	ToVpnGatewayNatRuleOutputWithContext(context.Context) VpnGatewayNatRuleOutput
}

type VpnGatewayNatRuleArgs struct {
	ExternalMappings  VpnNatRuleMappingArrayInput `pulumi:"externalMappings"`
	Id                pulumi.StringPtrInput       `pulumi:"id"`
	InternalMappings  VpnNatRuleMappingArrayInput `pulumi:"internalMappings"`
	IpConfigurationId pulumi.StringPtrInput       `pulumi:"ipConfigurationId"`
	Mode              pulumi.StringPtrInput       `pulumi:"mode"`
	Name              pulumi.StringPtrInput       `pulumi:"name"`
	Type              pulumi.StringPtrInput       `pulumi:"type"`
}

func (VpnGatewayNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayNatRule)(nil)).Elem()
}

func (i VpnGatewayNatRuleArgs) ToVpnGatewayNatRuleOutput() VpnGatewayNatRuleOutput {
	return i.ToVpnGatewayNatRuleOutputWithContext(context.Background())
}

func (i VpnGatewayNatRuleArgs) ToVpnGatewayNatRuleOutputWithContext(ctx context.Context) VpnGatewayNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayNatRuleOutput)
}





type VpnGatewayNatRuleArrayInput interface {
	pulumi.Input

	ToVpnGatewayNatRuleArrayOutput() VpnGatewayNatRuleArrayOutput
	ToVpnGatewayNatRuleArrayOutputWithContext(context.Context) VpnGatewayNatRuleArrayOutput
}

type VpnGatewayNatRuleArray []VpnGatewayNatRuleInput

func (VpnGatewayNatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGatewayNatRule)(nil)).Elem()
}

func (i VpnGatewayNatRuleArray) ToVpnGatewayNatRuleArrayOutput() VpnGatewayNatRuleArrayOutput {
	return i.ToVpnGatewayNatRuleArrayOutputWithContext(context.Background())
}

func (i VpnGatewayNatRuleArray) ToVpnGatewayNatRuleArrayOutputWithContext(ctx context.Context) VpnGatewayNatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayNatRuleArrayOutput)
}

type VpnGatewayNatRuleOutput struct{ *pulumi.OutputState }

func (VpnGatewayNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayNatRule)(nil)).Elem()
}

func (o VpnGatewayNatRuleOutput) ToVpnGatewayNatRuleOutput() VpnGatewayNatRuleOutput {
	return o
}

func (o VpnGatewayNatRuleOutput) ToVpnGatewayNatRuleOutputWithContext(ctx context.Context) VpnGatewayNatRuleOutput {
	return o
}

func (o VpnGatewayNatRuleOutput) ExternalMappings() VpnNatRuleMappingArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) []VpnNatRuleMapping { return v.ExternalMappings }).(VpnNatRuleMappingArrayOutput)
}

func (o VpnGatewayNatRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleOutput) InternalMappings() VpnNatRuleMappingArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) []VpnNatRuleMapping { return v.InternalMappings }).(VpnNatRuleMappingArrayOutput)
}

func (o VpnGatewayNatRuleOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRule) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VpnGatewayNatRuleArrayOutput struct{ *pulumi.OutputState }

func (VpnGatewayNatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGatewayNatRule)(nil)).Elem()
}

func (o VpnGatewayNatRuleArrayOutput) ToVpnGatewayNatRuleArrayOutput() VpnGatewayNatRuleArrayOutput {
	return o
}

func (o VpnGatewayNatRuleArrayOutput) ToVpnGatewayNatRuleArrayOutputWithContext(ctx context.Context) VpnGatewayNatRuleArrayOutput {
	return o
}

func (o VpnGatewayNatRuleArrayOutput) Index(i pulumi.IntInput) VpnGatewayNatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnGatewayNatRule {
		return vs[0].([]VpnGatewayNatRule)[vs[1].(int)]
	}).(VpnGatewayNatRuleOutput)
}

type VpnGatewayNatRuleResponse struct {
	EgressVpnSiteLinkConnections  []SubResourceResponse       `pulumi:"egressVpnSiteLinkConnections"`
	Etag                          string                      `pulumi:"etag"`
	ExternalMappings              []VpnNatRuleMappingResponse `pulumi:"externalMappings"`
	Id                            *string                     `pulumi:"id"`
	IngressVpnSiteLinkConnections []SubResourceResponse       `pulumi:"ingressVpnSiteLinkConnections"`
	InternalMappings              []VpnNatRuleMappingResponse `pulumi:"internalMappings"`
	IpConfigurationId             *string                     `pulumi:"ipConfigurationId"`
	Mode                          *string                     `pulumi:"mode"`
	Name                          *string                     `pulumi:"name"`
	ProvisioningState             string                      `pulumi:"provisioningState"`
	Type                          string                      `pulumi:"type"`
}

type VpnGatewayNatRuleResponseOutput struct{ *pulumi.OutputState }

func (VpnGatewayNatRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGatewayNatRuleResponse)(nil)).Elem()
}

func (o VpnGatewayNatRuleResponseOutput) ToVpnGatewayNatRuleResponseOutput() VpnGatewayNatRuleResponseOutput {
	return o
}

func (o VpnGatewayNatRuleResponseOutput) ToVpnGatewayNatRuleResponseOutputWithContext(ctx context.Context) VpnGatewayNatRuleResponseOutput {
	return o
}

func (o VpnGatewayNatRuleResponseOutput) EgressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) []SubResourceResponse { return v.EgressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnGatewayNatRuleResponseOutput) ExternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) []VpnNatRuleMappingResponse { return v.ExternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleResponseOutput) IngressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) []SubResourceResponse { return v.IngressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

func (o VpnGatewayNatRuleResponseOutput) InternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) []VpnNatRuleMappingResponse { return v.InternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

func (o VpnGatewayNatRuleResponseOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) *string { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnGatewayNatRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnGatewayNatRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnGatewayNatRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VpnGatewayNatRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnGatewayNatRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGatewayNatRuleResponse)(nil)).Elem()
}

func (o VpnGatewayNatRuleResponseArrayOutput) ToVpnGatewayNatRuleResponseArrayOutput() VpnGatewayNatRuleResponseArrayOutput {
	return o
}

func (o VpnGatewayNatRuleResponseArrayOutput) ToVpnGatewayNatRuleResponseArrayOutputWithContext(ctx context.Context) VpnGatewayNatRuleResponseArrayOutput {
	return o
}

func (o VpnGatewayNatRuleResponseArrayOutput) Index(i pulumi.IntInput) VpnGatewayNatRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnGatewayNatRuleResponse {
		return vs[0].([]VpnGatewayNatRuleResponse)[vs[1].(int)]
	}).(VpnGatewayNatRuleResponseOutput)
}

type VpnLinkBgpSettings struct {
	Asn               *float64 `pulumi:"asn"`
	BgpPeeringAddress *string  `pulumi:"bgpPeeringAddress"`
}





type VpnLinkBgpSettingsInput interface {
	pulumi.Input

	ToVpnLinkBgpSettingsOutput() VpnLinkBgpSettingsOutput
	ToVpnLinkBgpSettingsOutputWithContext(context.Context) VpnLinkBgpSettingsOutput
}

type VpnLinkBgpSettingsArgs struct {
	Asn               pulumi.Float64PtrInput `pulumi:"asn"`
	BgpPeeringAddress pulumi.StringPtrInput  `pulumi:"bgpPeeringAddress"`
}

func (VpnLinkBgpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkBgpSettings)(nil)).Elem()
}

func (i VpnLinkBgpSettingsArgs) ToVpnLinkBgpSettingsOutput() VpnLinkBgpSettingsOutput {
	return i.ToVpnLinkBgpSettingsOutputWithContext(context.Background())
}

func (i VpnLinkBgpSettingsArgs) ToVpnLinkBgpSettingsOutputWithContext(ctx context.Context) VpnLinkBgpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkBgpSettingsOutput)
}

func (i VpnLinkBgpSettingsArgs) ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput {
	return i.ToVpnLinkBgpSettingsPtrOutputWithContext(context.Background())
}

func (i VpnLinkBgpSettingsArgs) ToVpnLinkBgpSettingsPtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkBgpSettingsOutput).ToVpnLinkBgpSettingsPtrOutputWithContext(ctx)
}









type VpnLinkBgpSettingsPtrInput interface {
	pulumi.Input

	ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput
	ToVpnLinkBgpSettingsPtrOutputWithContext(context.Context) VpnLinkBgpSettingsPtrOutput
}

type vpnLinkBgpSettingsPtrType VpnLinkBgpSettingsArgs

func VpnLinkBgpSettingsPtr(v *VpnLinkBgpSettingsArgs) VpnLinkBgpSettingsPtrInput {
	return (*vpnLinkBgpSettingsPtrType)(v)
}

func (*vpnLinkBgpSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkBgpSettings)(nil)).Elem()
}

func (i *vpnLinkBgpSettingsPtrType) ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput {
	return i.ToVpnLinkBgpSettingsPtrOutputWithContext(context.Background())
}

func (i *vpnLinkBgpSettingsPtrType) ToVpnLinkBgpSettingsPtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkBgpSettingsPtrOutput)
}

type VpnLinkBgpSettingsOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkBgpSettings)(nil)).Elem()
}

func (o VpnLinkBgpSettingsOutput) ToVpnLinkBgpSettingsOutput() VpnLinkBgpSettingsOutput {
	return o
}

func (o VpnLinkBgpSettingsOutput) ToVpnLinkBgpSettingsOutputWithContext(ctx context.Context) VpnLinkBgpSettingsOutput {
	return o
}

func (o VpnLinkBgpSettingsOutput) ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput {
	return o.ToVpnLinkBgpSettingsPtrOutputWithContext(context.Background())
}

func (o VpnLinkBgpSettingsOutput) ToVpnLinkBgpSettingsPtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnLinkBgpSettings) *VpnLinkBgpSettings {
		return &v
	}).(VpnLinkBgpSettingsPtrOutput)
}

func (o VpnLinkBgpSettingsOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettings) *float64 { return v.Asn }).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettings) *string { return v.BgpPeeringAddress }).(pulumi.StringPtrOutput)
}

type VpnLinkBgpSettingsPtrOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkBgpSettings)(nil)).Elem()
}

func (o VpnLinkBgpSettingsPtrOutput) ToVpnLinkBgpSettingsPtrOutput() VpnLinkBgpSettingsPtrOutput {
	return o
}

func (o VpnLinkBgpSettingsPtrOutput) ToVpnLinkBgpSettingsPtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsPtrOutput {
	return o
}

func (o VpnLinkBgpSettingsPtrOutput) Elem() VpnLinkBgpSettingsOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettings) VpnLinkBgpSettings {
		if v != nil {
			return *v
		}
		var ret VpnLinkBgpSettings
		return ret
	}).(VpnLinkBgpSettingsOutput)
}

func (o VpnLinkBgpSettingsPtrOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.Asn
	}).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsPtrOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettings) *string {
		if v == nil {
			return nil
		}
		return v.BgpPeeringAddress
	}).(pulumi.StringPtrOutput)
}

type VpnLinkBgpSettingsResponse struct {
	Asn               *float64 `pulumi:"asn"`
	BgpPeeringAddress *string  `pulumi:"bgpPeeringAddress"`
}

type VpnLinkBgpSettingsResponseOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkBgpSettingsResponse)(nil)).Elem()
}

func (o VpnLinkBgpSettingsResponseOutput) ToVpnLinkBgpSettingsResponseOutput() VpnLinkBgpSettingsResponseOutput {
	return o
}

func (o VpnLinkBgpSettingsResponseOutput) ToVpnLinkBgpSettingsResponseOutputWithContext(ctx context.Context) VpnLinkBgpSettingsResponseOutput {
	return o
}

func (o VpnLinkBgpSettingsResponseOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettingsResponse) *float64 { return v.Asn }).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsResponseOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkBgpSettingsResponse) *string { return v.BgpPeeringAddress }).(pulumi.StringPtrOutput)
}

type VpnLinkBgpSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnLinkBgpSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkBgpSettingsResponse)(nil)).Elem()
}

func (o VpnLinkBgpSettingsResponsePtrOutput) ToVpnLinkBgpSettingsResponsePtrOutput() VpnLinkBgpSettingsResponsePtrOutput {
	return o
}

func (o VpnLinkBgpSettingsResponsePtrOutput) ToVpnLinkBgpSettingsResponsePtrOutputWithContext(ctx context.Context) VpnLinkBgpSettingsResponsePtrOutput {
	return o
}

func (o VpnLinkBgpSettingsResponsePtrOutput) Elem() VpnLinkBgpSettingsResponseOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettingsResponse) VpnLinkBgpSettingsResponse {
		if v != nil {
			return *v
		}
		var ret VpnLinkBgpSettingsResponse
		return ret
	}).(VpnLinkBgpSettingsResponseOutput)
}

func (o VpnLinkBgpSettingsResponsePtrOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Asn
	}).(pulumi.Float64PtrOutput)
}

func (o VpnLinkBgpSettingsResponsePtrOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkBgpSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.BgpPeeringAddress
	}).(pulumi.StringPtrOutput)
}

type VpnLinkProviderProperties struct {
	LinkProviderName *string `pulumi:"linkProviderName"`
	LinkSpeedInMbps  *int    `pulumi:"linkSpeedInMbps"`
}





type VpnLinkProviderPropertiesInput interface {
	pulumi.Input

	ToVpnLinkProviderPropertiesOutput() VpnLinkProviderPropertiesOutput
	ToVpnLinkProviderPropertiesOutputWithContext(context.Context) VpnLinkProviderPropertiesOutput
}

type VpnLinkProviderPropertiesArgs struct {
	LinkProviderName pulumi.StringPtrInput `pulumi:"linkProviderName"`
	LinkSpeedInMbps  pulumi.IntPtrInput    `pulumi:"linkSpeedInMbps"`
}

func (VpnLinkProviderPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkProviderProperties)(nil)).Elem()
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesOutput() VpnLinkProviderPropertiesOutput {
	return i.ToVpnLinkProviderPropertiesOutputWithContext(context.Background())
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkProviderPropertiesOutput)
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return i.ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i VpnLinkProviderPropertiesArgs) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkProviderPropertiesOutput).ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx)
}









type VpnLinkProviderPropertiesPtrInput interface {
	pulumi.Input

	ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput
	ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Context) VpnLinkProviderPropertiesPtrOutput
}

type vpnLinkProviderPropertiesPtrType VpnLinkProviderPropertiesArgs

func VpnLinkProviderPropertiesPtr(v *VpnLinkProviderPropertiesArgs) VpnLinkProviderPropertiesPtrInput {
	return (*vpnLinkProviderPropertiesPtrType)(v)
}

func (*vpnLinkProviderPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkProviderProperties)(nil)).Elem()
}

func (i *vpnLinkProviderPropertiesPtrType) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return i.ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i *vpnLinkProviderPropertiesPtrType) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnLinkProviderPropertiesPtrOutput)
}

type VpnLinkProviderPropertiesOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkProviderProperties)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesOutput() VpnLinkProviderPropertiesOutput {
	return o
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesOutput {
	return o
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return o.ToVpnLinkProviderPropertiesPtrOutputWithContext(context.Background())
}

func (o VpnLinkProviderPropertiesOutput) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnLinkProviderProperties) *VpnLinkProviderProperties {
		return &v
	}).(VpnLinkProviderPropertiesPtrOutput)
}

func (o VpnLinkProviderPropertiesOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderProperties) *string { return v.LinkProviderName }).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderProperties) *int { return v.LinkSpeedInMbps }).(pulumi.IntPtrOutput)
}

type VpnLinkProviderPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkProviderProperties)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesPtrOutput) ToVpnLinkProviderPropertiesPtrOutput() VpnLinkProviderPropertiesPtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesPtrOutput) ToVpnLinkProviderPropertiesPtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesPtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesPtrOutput) Elem() VpnLinkProviderPropertiesOutput {
	return o.ApplyT(func(v *VpnLinkProviderProperties) VpnLinkProviderProperties {
		if v != nil {
			return *v
		}
		var ret VpnLinkProviderProperties
		return ret
	}).(VpnLinkProviderPropertiesOutput)
}

func (o VpnLinkProviderPropertiesPtrOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderProperties) *string {
		if v == nil {
			return nil
		}
		return v.LinkProviderName
	}).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesPtrOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderProperties) *int {
		if v == nil {
			return nil
		}
		return v.LinkSpeedInMbps
	}).(pulumi.IntPtrOutput)
}

type VpnLinkProviderPropertiesResponse struct {
	LinkProviderName *string `pulumi:"linkProviderName"`
	LinkSpeedInMbps  *int    `pulumi:"linkSpeedInMbps"`
}

type VpnLinkProviderPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnLinkProviderPropertiesResponse)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesResponseOutput) ToVpnLinkProviderPropertiesResponseOutput() VpnLinkProviderPropertiesResponseOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponseOutput) ToVpnLinkProviderPropertiesResponseOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesResponseOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponseOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderPropertiesResponse) *string { return v.LinkProviderName }).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesResponseOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnLinkProviderPropertiesResponse) *int { return v.LinkSpeedInMbps }).(pulumi.IntPtrOutput)
}

type VpnLinkProviderPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnLinkProviderPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnLinkProviderPropertiesResponse)(nil)).Elem()
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) ToVpnLinkProviderPropertiesResponsePtrOutput() VpnLinkProviderPropertiesResponsePtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) ToVpnLinkProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) VpnLinkProviderPropertiesResponsePtrOutput {
	return o
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) Elem() VpnLinkProviderPropertiesResponseOutput {
	return o.ApplyT(func(v *VpnLinkProviderPropertiesResponse) VpnLinkProviderPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VpnLinkProviderPropertiesResponse
		return ret
	}).(VpnLinkProviderPropertiesResponseOutput)
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) LinkProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinkProviderName
	}).(pulumi.StringPtrOutput)
}

func (o VpnLinkProviderPropertiesResponsePtrOutput) LinkSpeedInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpnLinkProviderPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.LinkSpeedInMbps
	}).(pulumi.IntPtrOutput)
}

type VpnNatRuleMapping struct {
	AddressSpace *string `pulumi:"addressSpace"`
	PortRange    *string `pulumi:"portRange"`
}





type VpnNatRuleMappingInput interface {
	pulumi.Input

	ToVpnNatRuleMappingOutput() VpnNatRuleMappingOutput
	ToVpnNatRuleMappingOutputWithContext(context.Context) VpnNatRuleMappingOutput
}

type VpnNatRuleMappingArgs struct {
	AddressSpace pulumi.StringPtrInput `pulumi:"addressSpace"`
	PortRange    pulumi.StringPtrInput `pulumi:"portRange"`
}

func (VpnNatRuleMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMapping)(nil)).Elem()
}

func (i VpnNatRuleMappingArgs) ToVpnNatRuleMappingOutput() VpnNatRuleMappingOutput {
	return i.ToVpnNatRuleMappingOutputWithContext(context.Background())
}

func (i VpnNatRuleMappingArgs) ToVpnNatRuleMappingOutputWithContext(ctx context.Context) VpnNatRuleMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnNatRuleMappingOutput)
}





type VpnNatRuleMappingArrayInput interface {
	pulumi.Input

	ToVpnNatRuleMappingArrayOutput() VpnNatRuleMappingArrayOutput
	ToVpnNatRuleMappingArrayOutputWithContext(context.Context) VpnNatRuleMappingArrayOutput
}

type VpnNatRuleMappingArray []VpnNatRuleMappingInput

func (VpnNatRuleMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnNatRuleMapping)(nil)).Elem()
}

func (i VpnNatRuleMappingArray) ToVpnNatRuleMappingArrayOutput() VpnNatRuleMappingArrayOutput {
	return i.ToVpnNatRuleMappingArrayOutputWithContext(context.Background())
}

func (i VpnNatRuleMappingArray) ToVpnNatRuleMappingArrayOutputWithContext(ctx context.Context) VpnNatRuleMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnNatRuleMappingArrayOutput)
}

type VpnNatRuleMappingOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMapping)(nil)).Elem()
}

func (o VpnNatRuleMappingOutput) ToVpnNatRuleMappingOutput() VpnNatRuleMappingOutput {
	return o
}

func (o VpnNatRuleMappingOutput) ToVpnNatRuleMappingOutputWithContext(ctx context.Context) VpnNatRuleMappingOutput {
	return o
}

func (o VpnNatRuleMappingOutput) AddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMapping) *string { return v.AddressSpace }).(pulumi.StringPtrOutput)
}

func (o VpnNatRuleMappingOutput) PortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMapping) *string { return v.PortRange }).(pulumi.StringPtrOutput)
}

type VpnNatRuleMappingArrayOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnNatRuleMapping)(nil)).Elem()
}

func (o VpnNatRuleMappingArrayOutput) ToVpnNatRuleMappingArrayOutput() VpnNatRuleMappingArrayOutput {
	return o
}

func (o VpnNatRuleMappingArrayOutput) ToVpnNatRuleMappingArrayOutputWithContext(ctx context.Context) VpnNatRuleMappingArrayOutput {
	return o
}

func (o VpnNatRuleMappingArrayOutput) Index(i pulumi.IntInput) VpnNatRuleMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnNatRuleMapping {
		return vs[0].([]VpnNatRuleMapping)[vs[1].(int)]
	}).(VpnNatRuleMappingOutput)
}

type VpnNatRuleMappingResponse struct {
	AddressSpace *string `pulumi:"addressSpace"`
	PortRange    *string `pulumi:"portRange"`
}

type VpnNatRuleMappingResponseOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnNatRuleMappingResponse)(nil)).Elem()
}

func (o VpnNatRuleMappingResponseOutput) ToVpnNatRuleMappingResponseOutput() VpnNatRuleMappingResponseOutput {
	return o
}

func (o VpnNatRuleMappingResponseOutput) ToVpnNatRuleMappingResponseOutputWithContext(ctx context.Context) VpnNatRuleMappingResponseOutput {
	return o
}

func (o VpnNatRuleMappingResponseOutput) AddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMappingResponse) *string { return v.AddressSpace }).(pulumi.StringPtrOutput)
}

func (o VpnNatRuleMappingResponseOutput) PortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnNatRuleMappingResponse) *string { return v.PortRange }).(pulumi.StringPtrOutput)
}

type VpnNatRuleMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnNatRuleMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnNatRuleMappingResponse)(nil)).Elem()
}

func (o VpnNatRuleMappingResponseArrayOutput) ToVpnNatRuleMappingResponseArrayOutput() VpnNatRuleMappingResponseArrayOutput {
	return o
}

func (o VpnNatRuleMappingResponseArrayOutput) ToVpnNatRuleMappingResponseArrayOutputWithContext(ctx context.Context) VpnNatRuleMappingResponseArrayOutput {
	return o
}

func (o VpnNatRuleMappingResponseArrayOutput) Index(i pulumi.IntInput) VpnNatRuleMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnNatRuleMappingResponse {
		return vs[0].([]VpnNatRuleMappingResponse)[vs[1].(int)]
	}).(VpnNatRuleMappingResponseOutput)
}

type VpnServerConfigRadiusClientRootCertificate struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnServerConfigRadiusClientRootCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusClientRootCertificateOutput() VpnServerConfigRadiusClientRootCertificateOutput
	ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(context.Context) VpnServerConfigRadiusClientRootCertificateOutput
}

type VpnServerConfigRadiusClientRootCertificateArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnServerConfigRadiusClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusClientRootCertificateArgs) ToVpnServerConfigRadiusClientRootCertificateOutput() VpnServerConfigRadiusClientRootCertificateOutput {
	return i.ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusClientRootCertificateArgs) ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusClientRootCertificateOutput)
}





type VpnServerConfigRadiusClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusClientRootCertificateArrayOutput() VpnServerConfigRadiusClientRootCertificateArrayOutput
	ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(context.Context) VpnServerConfigRadiusClientRootCertificateArrayOutput
}

type VpnServerConfigRadiusClientRootCertificateArray []VpnServerConfigRadiusClientRootCertificateInput

func (VpnServerConfigRadiusClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusClientRootCertificateArray) ToVpnServerConfigRadiusClientRootCertificateArrayOutput() VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return i.ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusClientRootCertificateArray) ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusClientRootCertificateArrayOutput)
}

type VpnServerConfigRadiusClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) ToVpnServerConfigRadiusClientRootCertificateOutput() VpnServerConfigRadiusClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) ToVpnServerConfigRadiusClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusClientRootCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateArrayOutput) ToVpnServerConfigRadiusClientRootCertificateArrayOutput() VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateArrayOutput) ToVpnServerConfigRadiusClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusClientRootCertificate {
		return vs[0].([]VpnServerConfigRadiusClientRootCertificate)[vs[1].(int)]
	}).(VpnServerConfigRadiusClientRootCertificateOutput)
}

type VpnServerConfigRadiusClientRootCertificateResponse struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}

type VpnServerConfigRadiusClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) ToVpnServerConfigRadiusClientRootCertificateResponseOutput() VpnServerConfigRadiusClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) ToVpnServerConfigRadiusClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusClientRootCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusClientRootCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusClientRootCertificateResponseArrayOutput() VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusClientRootCertificateResponse {
		return vs[0].([]VpnServerConfigRadiusClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigRadiusClientRootCertificateResponseOutput)
}

type VpnServerConfigRadiusServerRootCertificate struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}





type VpnServerConfigRadiusServerRootCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusServerRootCertificateOutput() VpnServerConfigRadiusServerRootCertificateOutput
	ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(context.Context) VpnServerConfigRadiusServerRootCertificateOutput
}

type VpnServerConfigRadiusServerRootCertificateArgs struct {
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringPtrInput `pulumi:"publicCertData"`
}

func (VpnServerConfigRadiusServerRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusServerRootCertificateArgs) ToVpnServerConfigRadiusServerRootCertificateOutput() VpnServerConfigRadiusServerRootCertificateOutput {
	return i.ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusServerRootCertificateArgs) ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusServerRootCertificateOutput)
}





type VpnServerConfigRadiusServerRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigRadiusServerRootCertificateArrayOutput() VpnServerConfigRadiusServerRootCertificateArrayOutput
	ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(context.Context) VpnServerConfigRadiusServerRootCertificateArrayOutput
}

type VpnServerConfigRadiusServerRootCertificateArray []VpnServerConfigRadiusServerRootCertificateInput

func (VpnServerConfigRadiusServerRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigRadiusServerRootCertificateArray) ToVpnServerConfigRadiusServerRootCertificateArrayOutput() VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return i.ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigRadiusServerRootCertificateArray) ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigRadiusServerRootCertificateArrayOutput)
}

type VpnServerConfigRadiusServerRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) ToVpnServerConfigRadiusServerRootCertificateOutput() VpnServerConfigRadiusServerRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) ToVpnServerConfigRadiusServerRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusServerRootCertificateOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificate) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusServerRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusServerRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateArrayOutput) ToVpnServerConfigRadiusServerRootCertificateArrayOutput() VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateArrayOutput) ToVpnServerConfigRadiusServerRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusServerRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusServerRootCertificate {
		return vs[0].([]VpnServerConfigRadiusServerRootCertificate)[vs[1].(int)]
	}).(VpnServerConfigRadiusServerRootCertificateOutput)
}

type VpnServerConfigRadiusServerRootCertificateResponse struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}

type VpnServerConfigRadiusServerRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigRadiusServerRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) ToVpnServerConfigRadiusServerRootCertificateResponseOutput() VpnServerConfigRadiusServerRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) ToVpnServerConfigRadiusServerRootCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigRadiusServerRootCertificateResponseOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigRadiusServerRootCertificateResponse) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigRadiusServerRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigRadiusServerRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusServerRootCertificateResponseArrayOutput() VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) ToVpnServerConfigRadiusServerRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigRadiusServerRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigRadiusServerRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigRadiusServerRootCertificateResponse {
		return vs[0].([]VpnServerConfigRadiusServerRootCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigRadiusServerRootCertificateResponseOutput)
}

type VpnServerConfigVpnClientRevokedCertificate struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnServerConfigVpnClientRevokedCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRevokedCertificateOutput() VpnServerConfigVpnClientRevokedCertificateOutput
	ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(context.Context) VpnServerConfigVpnClientRevokedCertificateOutput
}

type VpnServerConfigVpnClientRevokedCertificateArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnServerConfigVpnClientRevokedCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRevokedCertificateArgs) ToVpnServerConfigVpnClientRevokedCertificateOutput() VpnServerConfigVpnClientRevokedCertificateOutput {
	return i.ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRevokedCertificateArgs) ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRevokedCertificateOutput)
}





type VpnServerConfigVpnClientRevokedCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRevokedCertificateArrayOutput() VpnServerConfigVpnClientRevokedCertificateArrayOutput
	ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(context.Context) VpnServerConfigVpnClientRevokedCertificateArrayOutput
}

type VpnServerConfigVpnClientRevokedCertificateArray []VpnServerConfigVpnClientRevokedCertificateInput

func (VpnServerConfigVpnClientRevokedCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRevokedCertificateArray) ToVpnServerConfigVpnClientRevokedCertificateArrayOutput() VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return i.ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRevokedCertificateArray) ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRevokedCertificateArrayOutput)
}

type VpnServerConfigVpnClientRevokedCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) ToVpnServerConfigVpnClientRevokedCertificateOutput() VpnServerConfigVpnClientRevokedCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) ToVpnServerConfigVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRevokedCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRevokedCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateArrayOutput() VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRevokedCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRevokedCertificate {
		return vs[0].([]VpnServerConfigVpnClientRevokedCertificate)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRevokedCertificateOutput)
}

type VpnServerConfigVpnClientRevokedCertificateResponse struct {
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}

type VpnServerConfigVpnClientRevokedCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseOutput() VpnServerConfigVpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRevokedCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput() VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRevokedCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRevokedCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRevokedCertificateResponse {
		return vs[0].([]VpnServerConfigVpnClientRevokedCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRevokedCertificateResponseOutput)
}

type VpnServerConfigVpnClientRootCertificate struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}





type VpnServerConfigVpnClientRootCertificateInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRootCertificateOutput() VpnServerConfigVpnClientRootCertificateOutput
	ToVpnServerConfigVpnClientRootCertificateOutputWithContext(context.Context) VpnServerConfigVpnClientRootCertificateOutput
}

type VpnServerConfigVpnClientRootCertificateArgs struct {
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringPtrInput `pulumi:"publicCertData"`
}

func (VpnServerConfigVpnClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRootCertificateArgs) ToVpnServerConfigVpnClientRootCertificateOutput() VpnServerConfigVpnClientRootCertificateOutput {
	return i.ToVpnServerConfigVpnClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRootCertificateArgs) ToVpnServerConfigVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRootCertificateOutput)
}





type VpnServerConfigVpnClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnServerConfigVpnClientRootCertificateArrayOutput() VpnServerConfigVpnClientRootCertificateArrayOutput
	ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(context.Context) VpnServerConfigVpnClientRootCertificateArrayOutput
}

type VpnServerConfigVpnClientRootCertificateArray []VpnServerConfigVpnClientRootCertificateInput

func (VpnServerConfigVpnClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (i VpnServerConfigVpnClientRootCertificateArray) ToVpnServerConfigVpnClientRootCertificateArrayOutput() VpnServerConfigVpnClientRootCertificateArrayOutput {
	return i.ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigVpnClientRootCertificateArray) ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigVpnClientRootCertificateArrayOutput)
}

type VpnServerConfigVpnClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateOutput) ToVpnServerConfigVpnClientRootCertificateOutput() VpnServerConfigVpnClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateOutput) ToVpnServerConfigVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRootCertificateOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificate) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRootCertificate)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateArrayOutput) ToVpnServerConfigVpnClientRootCertificateArrayOutput() VpnServerConfigVpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateArrayOutput) ToVpnServerConfigVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRootCertificate {
		return vs[0].([]VpnServerConfigVpnClientRootCertificate)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRootCertificateOutput)
}

type VpnServerConfigVpnClientRootCertificateResponse struct {
	Name           *string `pulumi:"name"`
	PublicCertData *string `pulumi:"publicCertData"`
}

type VpnServerConfigVpnClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigVpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) ToVpnServerConfigVpnClientRootCertificateResponseOutput() VpnServerConfigVpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) ToVpnServerConfigVpnClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigVpnClientRootCertificateResponseOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigVpnClientRootCertificateResponse) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type VpnServerConfigVpnClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigVpnClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigVpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnServerConfigVpnClientRootCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRootCertificateResponseArrayOutput() VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseArrayOutput) ToVpnServerConfigVpnClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnServerConfigVpnClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigVpnClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigVpnClientRootCertificateResponse {
		return vs[0].([]VpnServerConfigVpnClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnServerConfigVpnClientRootCertificateResponseOutput)
}

type VpnServerConfigurationPolicyGroup struct {
	Id            *string                                   `pulumi:"id"`
	IsDefault     *bool                                     `pulumi:"isDefault"`
	Name          *string                                   `pulumi:"name"`
	PolicyMembers []VpnServerConfigurationPolicyGroupMember `pulumi:"policyMembers"`
	Priority      *int                                      `pulumi:"priority"`
}





type VpnServerConfigurationPolicyGroupInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupOutput() VpnServerConfigurationPolicyGroupOutput
	ToVpnServerConfigurationPolicyGroupOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupOutput
}

type VpnServerConfigurationPolicyGroupArgs struct {
	Id            pulumi.StringPtrInput                             `pulumi:"id"`
	IsDefault     pulumi.BoolPtrInput                               `pulumi:"isDefault"`
	Name          pulumi.StringPtrInput                             `pulumi:"name"`
	PolicyMembers VpnServerConfigurationPolicyGroupMemberArrayInput `pulumi:"policyMembers"`
	Priority      pulumi.IntPtrInput                                `pulumi:"priority"`
}

func (VpnServerConfigurationPolicyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroup)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupArgs) ToVpnServerConfigurationPolicyGroupOutput() VpnServerConfigurationPolicyGroupOutput {
	return i.ToVpnServerConfigurationPolicyGroupOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupArgs) ToVpnServerConfigurationPolicyGroupOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupOutput)
}





type VpnServerConfigurationPolicyGroupArrayInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupArrayOutput() VpnServerConfigurationPolicyGroupArrayOutput
	ToVpnServerConfigurationPolicyGroupArrayOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupArrayOutput
}

type VpnServerConfigurationPolicyGroupArray []VpnServerConfigurationPolicyGroupInput

func (VpnServerConfigurationPolicyGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroup)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupArray) ToVpnServerConfigurationPolicyGroupArrayOutput() VpnServerConfigurationPolicyGroupArrayOutput {
	return i.ToVpnServerConfigurationPolicyGroupArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupArray) ToVpnServerConfigurationPolicyGroupArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupArrayOutput)
}

type VpnServerConfigurationPolicyGroupOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroup)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupOutput) ToVpnServerConfigurationPolicyGroupOutput() VpnServerConfigurationPolicyGroupOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupOutput) ToVpnServerConfigurationPolicyGroupOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupOutput) PolicyMembers() VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) []VpnServerConfigurationPolicyGroupMember {
		return v.PolicyMembers
	}).(VpnServerConfigurationPolicyGroupMemberArrayOutput)
}

func (o VpnServerConfigurationPolicyGroupOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroup) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

type VpnServerConfigurationPolicyGroupArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroup)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupArrayOutput) ToVpnServerConfigurationPolicyGroupArrayOutput() VpnServerConfigurationPolicyGroupArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupArrayOutput) ToVpnServerConfigurationPolicyGroupArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroup {
		return vs[0].([]VpnServerConfigurationPolicyGroup)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupOutput)
}

type VpnServerConfigurationPolicyGroupMember struct {
	AttributeType  *string `pulumi:"attributeType"`
	AttributeValue *string `pulumi:"attributeValue"`
	Name           *string `pulumi:"name"`
}





type VpnServerConfigurationPolicyGroupMemberInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupMemberOutput() VpnServerConfigurationPolicyGroupMemberOutput
	ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupMemberOutput
}

type VpnServerConfigurationPolicyGroupMemberArgs struct {
	AttributeType  pulumi.StringPtrInput `pulumi:"attributeType"`
	AttributeValue pulumi.StringPtrInput `pulumi:"attributeValue"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
}

func (VpnServerConfigurationPolicyGroupMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupMemberArgs) ToVpnServerConfigurationPolicyGroupMemberOutput() VpnServerConfigurationPolicyGroupMemberOutput {
	return i.ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupMemberArgs) ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupMemberOutput)
}





type VpnServerConfigurationPolicyGroupMemberArrayInput interface {
	pulumi.Input

	ToVpnServerConfigurationPolicyGroupMemberArrayOutput() VpnServerConfigurationPolicyGroupMemberArrayOutput
	ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(context.Context) VpnServerConfigurationPolicyGroupMemberArrayOutput
}

type VpnServerConfigurationPolicyGroupMemberArray []VpnServerConfigurationPolicyGroupMemberInput

func (VpnServerConfigurationPolicyGroupMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (i VpnServerConfigurationPolicyGroupMemberArray) ToVpnServerConfigurationPolicyGroupMemberArrayOutput() VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return i.ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(context.Background())
}

func (i VpnServerConfigurationPolicyGroupMemberArray) ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnServerConfigurationPolicyGroupMemberArrayOutput)
}

type VpnServerConfigurationPolicyGroupMemberOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) ToVpnServerConfigurationPolicyGroupMemberOutput() VpnServerConfigurationPolicyGroupMemberOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) ToVpnServerConfigurationPolicyGroupMemberOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) AttributeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMember) *string { return v.AttributeType }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMember) *string { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMember) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VpnServerConfigurationPolicyGroupMemberArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupMember)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberArrayOutput) ToVpnServerConfigurationPolicyGroupMemberArrayOutput() VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberArrayOutput) ToVpnServerConfigurationPolicyGroupMemberArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroupMember {
		return vs[0].([]VpnServerConfigurationPolicyGroupMember)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupMemberOutput)
}

type VpnServerConfigurationPolicyGroupMemberResponse struct {
	AttributeType  *string `pulumi:"attributeType"`
	AttributeValue *string `pulumi:"attributeValue"`
	Name           *string `pulumi:"name"`
}

type VpnServerConfigurationPolicyGroupMemberResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupMemberResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) ToVpnServerConfigurationPolicyGroupMemberResponseOutput() VpnServerConfigurationPolicyGroupMemberResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) ToVpnServerConfigurationPolicyGroupMemberResponseOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) AttributeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMemberResponse) *string { return v.AttributeType }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMemberResponse) *string { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupMemberResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupMemberResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VpnServerConfigurationPolicyGroupMemberResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupMemberResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) ToVpnServerConfigurationPolicyGroupMemberResponseArrayOutput() VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) ToVpnServerConfigurationPolicyGroupMemberResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupMemberResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupMemberResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroupMemberResponse {
		return vs[0].([]VpnServerConfigurationPolicyGroupMemberResponse)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupMemberResponseOutput)
}

type VpnServerConfigurationPolicyGroupResponse struct {
	Etag                        string                                            `pulumi:"etag"`
	Id                          *string                                           `pulumi:"id"`
	IsDefault                   *bool                                             `pulumi:"isDefault"`
	Name                        *string                                           `pulumi:"name"`
	P2SConnectionConfigurations []SubResourceResponse                             `pulumi:"p2SConnectionConfigurations"`
	PolicyMembers               []VpnServerConfigurationPolicyGroupMemberResponse `pulumi:"policyMembers"`
	Priority                    *int                                              `pulumi:"priority"`
	ProvisioningState           string                                            `pulumi:"provisioningState"`
	Type                        string                                            `pulumi:"type"`
}

type VpnServerConfigurationPolicyGroupResponseOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnServerConfigurationPolicyGroupResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) ToVpnServerConfigurationPolicyGroupResponseOutput() VpnServerConfigurationPolicyGroupResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) ToVpnServerConfigurationPolicyGroupResponseOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupResponseOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) P2SConnectionConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) []SubResourceResponse {
		return v.P2SConnectionConfigurations
	}).(SubResourceResponseArrayOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) PolicyMembers() VpnServerConfigurationPolicyGroupMemberResponseArrayOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) []VpnServerConfigurationPolicyGroupMemberResponse {
		return v.PolicyMembers
	}).(VpnServerConfigurationPolicyGroupMemberResponseArrayOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnServerConfigurationPolicyGroupResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnServerConfigurationPolicyGroupResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VpnServerConfigurationPolicyGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnServerConfigurationPolicyGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnServerConfigurationPolicyGroupResponse)(nil)).Elem()
}

func (o VpnServerConfigurationPolicyGroupResponseArrayOutput) ToVpnServerConfigurationPolicyGroupResponseArrayOutput() VpnServerConfigurationPolicyGroupResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupResponseArrayOutput) ToVpnServerConfigurationPolicyGroupResponseArrayOutputWithContext(ctx context.Context) VpnServerConfigurationPolicyGroupResponseArrayOutput {
	return o
}

func (o VpnServerConfigurationPolicyGroupResponseArrayOutput) Index(i pulumi.IntInput) VpnServerConfigurationPolicyGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnServerConfigurationPolicyGroupResponse {
		return vs[0].([]VpnServerConfigurationPolicyGroupResponse)[vs[1].(int)]
	}).(VpnServerConfigurationPolicyGroupResponseOutput)
}

type VpnSiteLink struct {
	BgpProperties  *VpnLinkBgpSettings        `pulumi:"bgpProperties"`
	Fqdn           *string                    `pulumi:"fqdn"`
	Id             *string                    `pulumi:"id"`
	IpAddress      *string                    `pulumi:"ipAddress"`
	LinkProperties *VpnLinkProviderProperties `pulumi:"linkProperties"`
	Name           *string                    `pulumi:"name"`
}





type VpnSiteLinkInput interface {
	pulumi.Input

	ToVpnSiteLinkOutput() VpnSiteLinkOutput
	ToVpnSiteLinkOutputWithContext(context.Context) VpnSiteLinkOutput
}

type VpnSiteLinkArgs struct {
	BgpProperties  VpnLinkBgpSettingsPtrInput        `pulumi:"bgpProperties"`
	Fqdn           pulumi.StringPtrInput             `pulumi:"fqdn"`
	Id             pulumi.StringPtrInput             `pulumi:"id"`
	IpAddress      pulumi.StringPtrInput             `pulumi:"ipAddress"`
	LinkProperties VpnLinkProviderPropertiesPtrInput `pulumi:"linkProperties"`
	Name           pulumi.StringPtrInput             `pulumi:"name"`
}

func (VpnSiteLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLink)(nil)).Elem()
}

func (i VpnSiteLinkArgs) ToVpnSiteLinkOutput() VpnSiteLinkOutput {
	return i.ToVpnSiteLinkOutputWithContext(context.Background())
}

func (i VpnSiteLinkArgs) ToVpnSiteLinkOutputWithContext(ctx context.Context) VpnSiteLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkOutput)
}





type VpnSiteLinkArrayInput interface {
	pulumi.Input

	ToVpnSiteLinkArrayOutput() VpnSiteLinkArrayOutput
	ToVpnSiteLinkArrayOutputWithContext(context.Context) VpnSiteLinkArrayOutput
}

type VpnSiteLinkArray []VpnSiteLinkInput

func (VpnSiteLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLink)(nil)).Elem()
}

func (i VpnSiteLinkArray) ToVpnSiteLinkArrayOutput() VpnSiteLinkArrayOutput {
	return i.ToVpnSiteLinkArrayOutputWithContext(context.Background())
}

func (i VpnSiteLinkArray) ToVpnSiteLinkArrayOutputWithContext(ctx context.Context) VpnSiteLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkArrayOutput)
}

type VpnSiteLinkOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLink)(nil)).Elem()
}

func (o VpnSiteLinkOutput) ToVpnSiteLinkOutput() VpnSiteLinkOutput {
	return o
}

func (o VpnSiteLinkOutput) ToVpnSiteLinkOutputWithContext(ctx context.Context) VpnSiteLinkOutput {
	return o
}

func (o VpnSiteLinkOutput) BgpProperties() VpnLinkBgpSettingsPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *VpnLinkBgpSettings { return v.BgpProperties }).(VpnLinkBgpSettingsPtrOutput)
}

func (o VpnSiteLinkOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkOutput) LinkProperties() VpnLinkProviderPropertiesPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *VpnLinkProviderProperties { return v.LinkProperties }).(VpnLinkProviderPropertiesPtrOutput)
}

func (o VpnSiteLinkOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLink) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VpnSiteLinkArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLink)(nil)).Elem()
}

func (o VpnSiteLinkArrayOutput) ToVpnSiteLinkArrayOutput() VpnSiteLinkArrayOutput {
	return o
}

func (o VpnSiteLinkArrayOutput) ToVpnSiteLinkArrayOutputWithContext(ctx context.Context) VpnSiteLinkArrayOutput {
	return o
}

func (o VpnSiteLinkArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLink {
		return vs[0].([]VpnSiteLink)[vs[1].(int)]
	}).(VpnSiteLinkOutput)
}

type VpnSiteLinkConnection struct {
	ConnectionBandwidth            *int                                       `pulumi:"connectionBandwidth"`
	EgressNatRules                 []SubResource                              `pulumi:"egressNatRules"`
	EnableBgp                      *bool                                      `pulumi:"enableBgp"`
	EnableRateLimiting             *bool                                      `pulumi:"enableRateLimiting"`
	Id                             *string                                    `pulumi:"id"`
	IngressNatRules                []SubResource                              `pulumi:"ingressNatRules"`
	IpsecPolicies                  []IpsecPolicy                              `pulumi:"ipsecPolicies"`
	Name                           *string                                    `pulumi:"name"`
	RoutingWeight                  *int                                       `pulumi:"routingWeight"`
	SharedKey                      *string                                    `pulumi:"sharedKey"`
	UseLocalAzureIpAddress         *bool                                      `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                                      `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string                                    `pulumi:"vpnConnectionProtocolType"`
	VpnGatewayCustomBgpAddresses   []GatewayCustomBgpIpAddressIpConfiguration `pulumi:"vpnGatewayCustomBgpAddresses"`
	VpnLinkConnectionMode          *string                                    `pulumi:"vpnLinkConnectionMode"`
	VpnSiteLink                    *SubResource                               `pulumi:"vpnSiteLink"`
}





type VpnSiteLinkConnectionInput interface {
	pulumi.Input

	ToVpnSiteLinkConnectionOutput() VpnSiteLinkConnectionOutput
	ToVpnSiteLinkConnectionOutputWithContext(context.Context) VpnSiteLinkConnectionOutput
}

type VpnSiteLinkConnectionArgs struct {
	ConnectionBandwidth            pulumi.IntPtrInput                                 `pulumi:"connectionBandwidth"`
	EgressNatRules                 SubResourceArrayInput                              `pulumi:"egressNatRules"`
	EnableBgp                      pulumi.BoolPtrInput                                `pulumi:"enableBgp"`
	EnableRateLimiting             pulumi.BoolPtrInput                                `pulumi:"enableRateLimiting"`
	Id                             pulumi.StringPtrInput                              `pulumi:"id"`
	IngressNatRules                SubResourceArrayInput                              `pulumi:"ingressNatRules"`
	IpsecPolicies                  IpsecPolicyArrayInput                              `pulumi:"ipsecPolicies"`
	Name                           pulumi.StringPtrInput                              `pulumi:"name"`
	RoutingWeight                  pulumi.IntPtrInput                                 `pulumi:"routingWeight"`
	SharedKey                      pulumi.StringPtrInput                              `pulumi:"sharedKey"`
	UseLocalAzureIpAddress         pulumi.BoolPtrInput                                `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrInput                                `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      pulumi.StringPtrInput                              `pulumi:"vpnConnectionProtocolType"`
	VpnGatewayCustomBgpAddresses   GatewayCustomBgpIpAddressIpConfigurationArrayInput `pulumi:"vpnGatewayCustomBgpAddresses"`
	VpnLinkConnectionMode          pulumi.StringPtrInput                              `pulumi:"vpnLinkConnectionMode"`
	VpnSiteLink                    SubResourcePtrInput                                `pulumi:"vpnSiteLink"`
}

func (VpnSiteLinkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkConnection)(nil)).Elem()
}

func (i VpnSiteLinkConnectionArgs) ToVpnSiteLinkConnectionOutput() VpnSiteLinkConnectionOutput {
	return i.ToVpnSiteLinkConnectionOutputWithContext(context.Background())
}

func (i VpnSiteLinkConnectionArgs) ToVpnSiteLinkConnectionOutputWithContext(ctx context.Context) VpnSiteLinkConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkConnectionOutput)
}





type VpnSiteLinkConnectionArrayInput interface {
	pulumi.Input

	ToVpnSiteLinkConnectionArrayOutput() VpnSiteLinkConnectionArrayOutput
	ToVpnSiteLinkConnectionArrayOutputWithContext(context.Context) VpnSiteLinkConnectionArrayOutput
}

type VpnSiteLinkConnectionArray []VpnSiteLinkConnectionInput

func (VpnSiteLinkConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkConnection)(nil)).Elem()
}

func (i VpnSiteLinkConnectionArray) ToVpnSiteLinkConnectionArrayOutput() VpnSiteLinkConnectionArrayOutput {
	return i.ToVpnSiteLinkConnectionArrayOutputWithContext(context.Background())
}

func (i VpnSiteLinkConnectionArray) ToVpnSiteLinkConnectionArrayOutputWithContext(ctx context.Context) VpnSiteLinkConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnSiteLinkConnectionArrayOutput)
}

type VpnSiteLinkConnectionOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkConnection)(nil)).Elem()
}

func (o VpnSiteLinkConnectionOutput) ToVpnSiteLinkConnectionOutput() VpnSiteLinkConnectionOutput {
	return o
}

func (o VpnSiteLinkConnectionOutput) ToVpnSiteLinkConnectionOutputWithContext(ctx context.Context) VpnSiteLinkConnectionOutput {
	return o
}

func (o VpnSiteLinkConnectionOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) EgressNatRules() SubResourceArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []SubResource { return v.EgressNatRules }).(SubResourceArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) IngressNatRules() SubResourceArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []SubResource { return v.IngressNatRules }).(SubResourceArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) IpsecPolicies() IpsecPolicyArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []IpsecPolicy { return v.IpsecPolicies }).(IpsecPolicyArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnGatewayCustomBgpAddresses() GatewayCustomBgpIpAddressIpConfigurationArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) []GatewayCustomBgpIpAddressIpConfiguration {
		return v.VpnGatewayCustomBgpAddresses
	}).(GatewayCustomBgpIpAddressIpConfigurationArrayOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnLinkConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *string { return v.VpnLinkConnectionMode }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionOutput) VpnSiteLink() SubResourcePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnection) *SubResource { return v.VpnSiteLink }).(SubResourcePtrOutput)
}

type VpnSiteLinkConnectionArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkConnection)(nil)).Elem()
}

func (o VpnSiteLinkConnectionArrayOutput) ToVpnSiteLinkConnectionArrayOutput() VpnSiteLinkConnectionArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionArrayOutput) ToVpnSiteLinkConnectionArrayOutputWithContext(ctx context.Context) VpnSiteLinkConnectionArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLinkConnection {
		return vs[0].([]VpnSiteLinkConnection)[vs[1].(int)]
	}).(VpnSiteLinkConnectionOutput)
}

type VpnSiteLinkConnectionResponse struct {
	ConnectionBandwidth            *int                                               `pulumi:"connectionBandwidth"`
	ConnectionStatus               string                                             `pulumi:"connectionStatus"`
	EgressBytesTransferred         float64                                            `pulumi:"egressBytesTransferred"`
	EgressNatRules                 []SubResourceResponse                              `pulumi:"egressNatRules"`
	EnableBgp                      *bool                                              `pulumi:"enableBgp"`
	EnableRateLimiting             *bool                                              `pulumi:"enableRateLimiting"`
	Etag                           string                                             `pulumi:"etag"`
	Id                             *string                                            `pulumi:"id"`
	IngressBytesTransferred        float64                                            `pulumi:"ingressBytesTransferred"`
	IngressNatRules                []SubResourceResponse                              `pulumi:"ingressNatRules"`
	IpsecPolicies                  []IpsecPolicyResponse                              `pulumi:"ipsecPolicies"`
	Name                           *string                                            `pulumi:"name"`
	ProvisioningState              string                                             `pulumi:"provisioningState"`
	RoutingWeight                  *int                                               `pulumi:"routingWeight"`
	SharedKey                      *string                                            `pulumi:"sharedKey"`
	Type                           string                                             `pulumi:"type"`
	UseLocalAzureIpAddress         *bool                                              `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                                              `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string                                            `pulumi:"vpnConnectionProtocolType"`
	VpnGatewayCustomBgpAddresses   []GatewayCustomBgpIpAddressIpConfigurationResponse `pulumi:"vpnGatewayCustomBgpAddresses"`
	VpnLinkConnectionMode          *string                                            `pulumi:"vpnLinkConnectionMode"`
	VpnSiteLink                    *SubResourceResponse                               `pulumi:"vpnSiteLink"`
}

type VpnSiteLinkConnectionResponseOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkConnectionResponse)(nil)).Elem()
}

func (o VpnSiteLinkConnectionResponseOutput) ToVpnSiteLinkConnectionResponseOutput() VpnSiteLinkConnectionResponseOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseOutput) ToVpnSiteLinkConnectionResponseOutputWithContext(ctx context.Context) VpnSiteLinkConnectionResponseOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseOutput) ConnectionBandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *int { return v.ConnectionBandwidth }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnSiteLinkConnectionResponseOutput) EgressNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []SubResourceResponse { return v.EgressNatRules }).(SubResourceResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) EnableRateLimiting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.EnableRateLimiting }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VpnSiteLinkConnectionResponseOutput) IngressNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []SubResourceResponse { return v.IngressNatRules }).(SubResourceResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnConnectionProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.VpnConnectionProtocolType }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnGatewayCustomBgpAddresses() GatewayCustomBgpIpAddressIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) []GatewayCustomBgpIpAddressIpConfigurationResponse {
		return v.VpnGatewayCustomBgpAddresses
	}).(GatewayCustomBgpIpAddressIpConfigurationResponseArrayOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnLinkConnectionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *string { return v.VpnLinkConnectionMode }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkConnectionResponseOutput) VpnSiteLink() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkConnectionResponse) *SubResourceResponse { return v.VpnSiteLink }).(SubResourceResponsePtrOutput)
}

type VpnSiteLinkConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkConnectionResponse)(nil)).Elem()
}

func (o VpnSiteLinkConnectionResponseArrayOutput) ToVpnSiteLinkConnectionResponseArrayOutput() VpnSiteLinkConnectionResponseArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseArrayOutput) ToVpnSiteLinkConnectionResponseArrayOutputWithContext(ctx context.Context) VpnSiteLinkConnectionResponseArrayOutput {
	return o
}

func (o VpnSiteLinkConnectionResponseArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLinkConnectionResponse {
		return vs[0].([]VpnSiteLinkConnectionResponse)[vs[1].(int)]
	}).(VpnSiteLinkConnectionResponseOutput)
}

type VpnSiteLinkResponse struct {
	BgpProperties     *VpnLinkBgpSettingsResponse        `pulumi:"bgpProperties"`
	Etag              string                             `pulumi:"etag"`
	Fqdn              *string                            `pulumi:"fqdn"`
	Id                *string                            `pulumi:"id"`
	IpAddress         *string                            `pulumi:"ipAddress"`
	LinkProperties    *VpnLinkProviderPropertiesResponse `pulumi:"linkProperties"`
	Name              *string                            `pulumi:"name"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	Type              string                             `pulumi:"type"`
}

type VpnSiteLinkResponseOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnSiteLinkResponse)(nil)).Elem()
}

func (o VpnSiteLinkResponseOutput) ToVpnSiteLinkResponseOutput() VpnSiteLinkResponseOutput {
	return o
}

func (o VpnSiteLinkResponseOutput) ToVpnSiteLinkResponseOutputWithContext(ctx context.Context) VpnSiteLinkResponseOutput {
	return o
}

func (o VpnSiteLinkResponseOutput) BgpProperties() VpnLinkBgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *VpnLinkBgpSettingsResponse { return v.BgpProperties }).(VpnLinkBgpSettingsResponsePtrOutput)
}

func (o VpnSiteLinkResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o VpnSiteLinkResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) LinkProperties() VpnLinkProviderPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *VpnLinkProviderPropertiesResponse { return v.LinkProperties }).(VpnLinkProviderPropertiesResponsePtrOutput)
}

func (o VpnSiteLinkResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnSiteLinkResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnSiteLinkResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VpnSiteLinkResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VpnSiteLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnSiteLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnSiteLinkResponse)(nil)).Elem()
}

func (o VpnSiteLinkResponseArrayOutput) ToVpnSiteLinkResponseArrayOutput() VpnSiteLinkResponseArrayOutput {
	return o
}

func (o VpnSiteLinkResponseArrayOutput) ToVpnSiteLinkResponseArrayOutputWithContext(ctx context.Context) VpnSiteLinkResponseArrayOutput {
	return o
}

func (o VpnSiteLinkResponseArrayOutput) Index(i pulumi.IntInput) VpnSiteLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnSiteLinkResponse {
		return vs[0].([]VpnSiteLinkResponse)[vs[1].(int)]
	}).(VpnSiteLinkResponseOutput)
}

type WebApplicationFirewallCustomRule struct {
	Action          string           `pulumi:"action"`
	MatchConditions []MatchCondition `pulumi:"matchConditions"`
	Name            *string          `pulumi:"name"`
	Priority        int              `pulumi:"priority"`
	RuleType        string           `pulumi:"ruleType"`
}





type WebApplicationFirewallCustomRuleInput interface {
	pulumi.Input

	ToWebApplicationFirewallCustomRuleOutput() WebApplicationFirewallCustomRuleOutput
	ToWebApplicationFirewallCustomRuleOutputWithContext(context.Context) WebApplicationFirewallCustomRuleOutput
}

type WebApplicationFirewallCustomRuleArgs struct {
	Action          pulumi.StringInput       `pulumi:"action"`
	MatchConditions MatchConditionArrayInput `pulumi:"matchConditions"`
	Name            pulumi.StringPtrInput    `pulumi:"name"`
	Priority        pulumi.IntInput          `pulumi:"priority"`
	RuleType        pulumi.StringInput       `pulumi:"ruleType"`
}

func (WebApplicationFirewallCustomRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (i WebApplicationFirewallCustomRuleArgs) ToWebApplicationFirewallCustomRuleOutput() WebApplicationFirewallCustomRuleOutput {
	return i.ToWebApplicationFirewallCustomRuleOutputWithContext(context.Background())
}

func (i WebApplicationFirewallCustomRuleArgs) ToWebApplicationFirewallCustomRuleOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebApplicationFirewallCustomRuleOutput)
}





type WebApplicationFirewallCustomRuleArrayInput interface {
	pulumi.Input

	ToWebApplicationFirewallCustomRuleArrayOutput() WebApplicationFirewallCustomRuleArrayOutput
	ToWebApplicationFirewallCustomRuleArrayOutputWithContext(context.Context) WebApplicationFirewallCustomRuleArrayOutput
}

type WebApplicationFirewallCustomRuleArray []WebApplicationFirewallCustomRuleInput

func (WebApplicationFirewallCustomRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (i WebApplicationFirewallCustomRuleArray) ToWebApplicationFirewallCustomRuleArrayOutput() WebApplicationFirewallCustomRuleArrayOutput {
	return i.ToWebApplicationFirewallCustomRuleArrayOutputWithContext(context.Background())
}

func (i WebApplicationFirewallCustomRuleArray) ToWebApplicationFirewallCustomRuleArrayOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebApplicationFirewallCustomRuleArrayOutput)
}

type WebApplicationFirewallCustomRuleOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleOutput) ToWebApplicationFirewallCustomRuleOutput() WebApplicationFirewallCustomRuleOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleOutput) ToWebApplicationFirewallCustomRuleOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) string { return v.Action }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) MatchConditions() MatchConditionArrayOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) []MatchCondition { return v.MatchConditions }).(MatchConditionArrayOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) int { return v.Priority }).(pulumi.IntOutput)
}

func (o WebApplicationFirewallCustomRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRule) string { return v.RuleType }).(pulumi.StringOutput)
}

type WebApplicationFirewallCustomRuleArrayOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApplicationFirewallCustomRule)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleArrayOutput) ToWebApplicationFirewallCustomRuleArrayOutput() WebApplicationFirewallCustomRuleArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleArrayOutput) ToWebApplicationFirewallCustomRuleArrayOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleArrayOutput) Index(i pulumi.IntInput) WebApplicationFirewallCustomRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebApplicationFirewallCustomRule {
		return vs[0].([]WebApplicationFirewallCustomRule)[vs[1].(int)]
	}).(WebApplicationFirewallCustomRuleOutput)
}

type WebApplicationFirewallCustomRuleResponse struct {
	Action          string                   `pulumi:"action"`
	Etag            string                   `pulumi:"etag"`
	MatchConditions []MatchConditionResponse `pulumi:"matchConditions"`
	Name            *string                  `pulumi:"name"`
	Priority        int                      `pulumi:"priority"`
	RuleType        string                   `pulumi:"ruleType"`
}

type WebApplicationFirewallCustomRuleResponseOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebApplicationFirewallCustomRuleResponse)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleResponseOutput) ToWebApplicationFirewallCustomRuleResponseOutput() WebApplicationFirewallCustomRuleResponseOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseOutput) ToWebApplicationFirewallCustomRuleResponseOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleResponseOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) MatchConditions() MatchConditionResponseArrayOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) []MatchConditionResponse { return v.MatchConditions }).(MatchConditionResponseArrayOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o WebApplicationFirewallCustomRuleResponseOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v WebApplicationFirewallCustomRuleResponse) string { return v.RuleType }).(pulumi.StringOutput)
}

type WebApplicationFirewallCustomRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallCustomRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebApplicationFirewallCustomRuleResponse)(nil)).Elem()
}

func (o WebApplicationFirewallCustomRuleResponseArrayOutput) ToWebApplicationFirewallCustomRuleResponseArrayOutput() WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseArrayOutput) ToWebApplicationFirewallCustomRuleResponseArrayOutputWithContext(ctx context.Context) WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o
}

func (o WebApplicationFirewallCustomRuleResponseArrayOutput) Index(i pulumi.IntInput) WebApplicationFirewallCustomRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebApplicationFirewallCustomRuleResponse {
		return vs[0].([]WebApplicationFirewallCustomRuleResponse)[vs[1].(int)]
	}).(WebApplicationFirewallCustomRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkTapTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapTypePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapTypeArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkTapResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualRouterAutoScaleConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualRouterAutoScaleConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualRouterAutoScaleConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualRouterAutoScaleConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VnetRouteOutput{})
	pulumi.RegisterOutputType(VnetRoutePtrOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseOutput{})
	pulumi.RegisterOutputType(VnetRouteResponsePtrOutput{})
	pulumi.RegisterOutputType(VngClientConnectionConfigurationOutput{})
	pulumi.RegisterOutputType(VngClientConnectionConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VngClientConnectionConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VngClientConnectionConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnClientConnectionHealthDetailResponseOutput{})
	pulumi.RegisterOutputType(VpnClientConnectionHealthDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnClientConnectionHealthResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnConnectionTypeOutput{})
	pulumi.RegisterOutputType(VpnConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(VpnConnectionResponseOutput{})
	pulumi.RegisterOutputType(VpnConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnGatewayIpConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VpnGatewayIpConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnGatewayNatRuleOutput{})
	pulumi.RegisterOutputType(VpnGatewayNatRuleArrayOutput{})
	pulumi.RegisterOutputType(VpnGatewayNatRuleResponseOutput{})
	pulumi.RegisterOutputType(VpnGatewayNatRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnLinkBgpSettingsOutput{})
	pulumi.RegisterOutputType(VpnLinkBgpSettingsPtrOutput{})
	pulumi.RegisterOutputType(VpnLinkBgpSettingsResponseOutput{})
	pulumi.RegisterOutputType(VpnLinkBgpSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VpnLinkProviderPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingArrayOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingResponseOutput{})
	pulumi.RegisterOutputType(VpnNatRuleMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusClientRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigRadiusServerRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRevokedCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigVpnClientRootCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupMemberResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupResponseOutput{})
	pulumi.RegisterOutputType(VpnServerConfigurationPolicyGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionResponseOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkResponseOutput{})
	pulumi.RegisterOutputType(VpnSiteLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleArrayOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleResponseOutput{})
	pulumi.RegisterOutputType(WebApplicationFirewallCustomRuleResponseArrayOutput{})
}
