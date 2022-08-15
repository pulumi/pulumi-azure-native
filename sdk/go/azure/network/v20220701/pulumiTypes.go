


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InboundEndpointIPConfiguration struct {
	PrivateIpAddress          *string     `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod *string     `pulumi:"privateIpAllocationMethod"`
	Subnet                    SubResource `pulumi:"subnet"`
}


func (val *InboundEndpointIPConfiguration) Defaults() *InboundEndpointIPConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrivateIpAllocationMethod) {
		privateIpAllocationMethod_ := "Dynamic"
		tmp.PrivateIpAllocationMethod = &privateIpAllocationMethod_
	}
	return &tmp
}





type InboundEndpointIPConfigurationInput interface {
	pulumi.Input

	ToInboundEndpointIPConfigurationOutput() InboundEndpointIPConfigurationOutput
	ToInboundEndpointIPConfigurationOutputWithContext(context.Context) InboundEndpointIPConfigurationOutput
}

type InboundEndpointIPConfigurationArgs struct {
	PrivateIpAddress          pulumi.StringPtrInput `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod pulumi.StringPtrInput `pulumi:"privateIpAllocationMethod"`
	Subnet                    SubResourceInput      `pulumi:"subnet"`
}


func (val *InboundEndpointIPConfigurationArgs) Defaults() *InboundEndpointIPConfigurationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrivateIpAllocationMethod) {
		tmp.PrivateIpAllocationMethod = pulumi.StringPtr("Dynamic")
	}
	return &tmp
}
func (InboundEndpointIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundEndpointIPConfiguration)(nil)).Elem()
}

func (i InboundEndpointIPConfigurationArgs) ToInboundEndpointIPConfigurationOutput() InboundEndpointIPConfigurationOutput {
	return i.ToInboundEndpointIPConfigurationOutputWithContext(context.Background())
}

func (i InboundEndpointIPConfigurationArgs) ToInboundEndpointIPConfigurationOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundEndpointIPConfigurationOutput)
}





type InboundEndpointIPConfigurationArrayInput interface {
	pulumi.Input

	ToInboundEndpointIPConfigurationArrayOutput() InboundEndpointIPConfigurationArrayOutput
	ToInboundEndpointIPConfigurationArrayOutputWithContext(context.Context) InboundEndpointIPConfigurationArrayOutput
}

type InboundEndpointIPConfigurationArray []InboundEndpointIPConfigurationInput

func (InboundEndpointIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundEndpointIPConfiguration)(nil)).Elem()
}

func (i InboundEndpointIPConfigurationArray) ToInboundEndpointIPConfigurationArrayOutput() InboundEndpointIPConfigurationArrayOutput {
	return i.ToInboundEndpointIPConfigurationArrayOutputWithContext(context.Background())
}

func (i InboundEndpointIPConfigurationArray) ToInboundEndpointIPConfigurationArrayOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundEndpointIPConfigurationArrayOutput)
}

type InboundEndpointIPConfigurationOutput struct{ *pulumi.OutputState }

func (InboundEndpointIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundEndpointIPConfiguration)(nil)).Elem()
}

func (o InboundEndpointIPConfigurationOutput) ToInboundEndpointIPConfigurationOutput() InboundEndpointIPConfigurationOutput {
	return o
}

func (o InboundEndpointIPConfigurationOutput) ToInboundEndpointIPConfigurationOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationOutput {
	return o
}

func (o InboundEndpointIPConfigurationOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfiguration) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o InboundEndpointIPConfigurationOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfiguration) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o InboundEndpointIPConfigurationOutput) Subnet() SubResourceOutput {
	return o.ApplyT(func(v InboundEndpointIPConfiguration) SubResource { return v.Subnet }).(SubResourceOutput)
}

type InboundEndpointIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (InboundEndpointIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundEndpointIPConfiguration)(nil)).Elem()
}

func (o InboundEndpointIPConfigurationArrayOutput) ToInboundEndpointIPConfigurationArrayOutput() InboundEndpointIPConfigurationArrayOutput {
	return o
}

func (o InboundEndpointIPConfigurationArrayOutput) ToInboundEndpointIPConfigurationArrayOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationArrayOutput {
	return o
}

func (o InboundEndpointIPConfigurationArrayOutput) Index(i pulumi.IntInput) InboundEndpointIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundEndpointIPConfiguration {
		return vs[0].([]InboundEndpointIPConfiguration)[vs[1].(int)]
	}).(InboundEndpointIPConfigurationOutput)
}

type InboundEndpointIPConfigurationResponse struct {
	PrivateIpAddress          *string             `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod *string             `pulumi:"privateIpAllocationMethod"`
	Subnet                    SubResourceResponse `pulumi:"subnet"`
}


func (val *InboundEndpointIPConfigurationResponse) Defaults() *InboundEndpointIPConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrivateIpAllocationMethod) {
		privateIpAllocationMethod_ := "Dynamic"
		tmp.PrivateIpAllocationMethod = &privateIpAllocationMethod_
	}
	return &tmp
}

type InboundEndpointIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (InboundEndpointIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundEndpointIPConfigurationResponse)(nil)).Elem()
}

func (o InboundEndpointIPConfigurationResponseOutput) ToInboundEndpointIPConfigurationResponseOutput() InboundEndpointIPConfigurationResponseOutput {
	return o
}

func (o InboundEndpointIPConfigurationResponseOutput) ToInboundEndpointIPConfigurationResponseOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationResponseOutput {
	return o
}

func (o InboundEndpointIPConfigurationResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfigurationResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o InboundEndpointIPConfigurationResponseOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundEndpointIPConfigurationResponse) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o InboundEndpointIPConfigurationResponseOutput) Subnet() SubResourceResponseOutput {
	return o.ApplyT(func(v InboundEndpointIPConfigurationResponse) SubResourceResponse { return v.Subnet }).(SubResourceResponseOutput)
}

type InboundEndpointIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundEndpointIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundEndpointIPConfigurationResponse)(nil)).Elem()
}

func (o InboundEndpointIPConfigurationResponseArrayOutput) ToInboundEndpointIPConfigurationResponseArrayOutput() InboundEndpointIPConfigurationResponseArrayOutput {
	return o
}

func (o InboundEndpointIPConfigurationResponseArrayOutput) ToInboundEndpointIPConfigurationResponseArrayOutputWithContext(ctx context.Context) InboundEndpointIPConfigurationResponseArrayOutput {
	return o
}

func (o InboundEndpointIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) InboundEndpointIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundEndpointIPConfigurationResponse {
		return vs[0].([]InboundEndpointIPConfigurationResponse)[vs[1].(int)]
	}).(InboundEndpointIPConfigurationResponseOutput)
}

type SubResource struct {
	Id string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}





type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SubResource) string { return v.Id }).(pulumi.StringOutput)
}

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
}

type SubResourceResponse struct {
	Id string `pulumi:"id"`
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SubResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type TargetDnsServer struct {
	IpAddress string `pulumi:"ipAddress"`
	Port      *int   `pulumi:"port"`
}


func (val *TargetDnsServer) Defaults() *TargetDnsServer {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Port) {
		port_ := 53
		tmp.Port = &port_
	}
	return &tmp
}





type TargetDnsServerInput interface {
	pulumi.Input

	ToTargetDnsServerOutput() TargetDnsServerOutput
	ToTargetDnsServerOutputWithContext(context.Context) TargetDnsServerOutput
}

type TargetDnsServerArgs struct {
	IpAddress pulumi.StringInput `pulumi:"ipAddress"`
	Port      pulumi.IntPtrInput `pulumi:"port"`
}


func (val *TargetDnsServerArgs) Defaults() *TargetDnsServerArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Port) {
		tmp.Port = pulumi.IntPtr(53)
	}
	return &tmp
}
func (TargetDnsServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetDnsServer)(nil)).Elem()
}

func (i TargetDnsServerArgs) ToTargetDnsServerOutput() TargetDnsServerOutput {
	return i.ToTargetDnsServerOutputWithContext(context.Background())
}

func (i TargetDnsServerArgs) ToTargetDnsServerOutputWithContext(ctx context.Context) TargetDnsServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetDnsServerOutput)
}





type TargetDnsServerArrayInput interface {
	pulumi.Input

	ToTargetDnsServerArrayOutput() TargetDnsServerArrayOutput
	ToTargetDnsServerArrayOutputWithContext(context.Context) TargetDnsServerArrayOutput
}

type TargetDnsServerArray []TargetDnsServerInput

func (TargetDnsServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetDnsServer)(nil)).Elem()
}

func (i TargetDnsServerArray) ToTargetDnsServerArrayOutput() TargetDnsServerArrayOutput {
	return i.ToTargetDnsServerArrayOutputWithContext(context.Background())
}

func (i TargetDnsServerArray) ToTargetDnsServerArrayOutputWithContext(ctx context.Context) TargetDnsServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetDnsServerArrayOutput)
}

type TargetDnsServerOutput struct{ *pulumi.OutputState }

func (TargetDnsServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetDnsServer)(nil)).Elem()
}

func (o TargetDnsServerOutput) ToTargetDnsServerOutput() TargetDnsServerOutput {
	return o
}

func (o TargetDnsServerOutput) ToTargetDnsServerOutputWithContext(ctx context.Context) TargetDnsServerOutput {
	return o
}

func (o TargetDnsServerOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v TargetDnsServer) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o TargetDnsServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TargetDnsServer) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type TargetDnsServerArrayOutput struct{ *pulumi.OutputState }

func (TargetDnsServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetDnsServer)(nil)).Elem()
}

func (o TargetDnsServerArrayOutput) ToTargetDnsServerArrayOutput() TargetDnsServerArrayOutput {
	return o
}

func (o TargetDnsServerArrayOutput) ToTargetDnsServerArrayOutputWithContext(ctx context.Context) TargetDnsServerArrayOutput {
	return o
}

func (o TargetDnsServerArrayOutput) Index(i pulumi.IntInput) TargetDnsServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetDnsServer {
		return vs[0].([]TargetDnsServer)[vs[1].(int)]
	}).(TargetDnsServerOutput)
}

type TargetDnsServerResponse struct {
	IpAddress string `pulumi:"ipAddress"`
	Port      *int   `pulumi:"port"`
}


func (val *TargetDnsServerResponse) Defaults() *TargetDnsServerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Port) {
		port_ := 53
		tmp.Port = &port_
	}
	return &tmp
}

type TargetDnsServerResponseOutput struct{ *pulumi.OutputState }

func (TargetDnsServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetDnsServerResponse)(nil)).Elem()
}

func (o TargetDnsServerResponseOutput) ToTargetDnsServerResponseOutput() TargetDnsServerResponseOutput {
	return o
}

func (o TargetDnsServerResponseOutput) ToTargetDnsServerResponseOutputWithContext(ctx context.Context) TargetDnsServerResponseOutput {
	return o
}

func (o TargetDnsServerResponseOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v TargetDnsServerResponse) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o TargetDnsServerResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TargetDnsServerResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type TargetDnsServerResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetDnsServerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetDnsServerResponse)(nil)).Elem()
}

func (o TargetDnsServerResponseArrayOutput) ToTargetDnsServerResponseArrayOutput() TargetDnsServerResponseArrayOutput {
	return o
}

func (o TargetDnsServerResponseArrayOutput) ToTargetDnsServerResponseArrayOutputWithContext(ctx context.Context) TargetDnsServerResponseArrayOutput {
	return o
}

func (o TargetDnsServerResponseArrayOutput) Index(i pulumi.IntInput) TargetDnsServerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetDnsServerResponse {
		return vs[0].([]TargetDnsServerResponse)[vs[1].(int)]
	}).(TargetDnsServerResponseOutput)
}

type VirtualNetworkDnsForwardingRulesetResponse struct {
	Id                 *string              `pulumi:"id"`
	VirtualNetworkLink *SubResourceResponse `pulumi:"virtualNetworkLink"`
}

type VirtualNetworkDnsForwardingRulesetResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsForwardingRulesetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkDnsForwardingRulesetResponse)(nil)).Elem()
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) ToVirtualNetworkDnsForwardingRulesetResponseOutput() VirtualNetworkDnsForwardingRulesetResponseOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) ToVirtualNetworkDnsForwardingRulesetResponseOutputWithContext(ctx context.Context) VirtualNetworkDnsForwardingRulesetResponseOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkDnsForwardingRulesetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkDnsForwardingRulesetResponseOutput) VirtualNetworkLink() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkDnsForwardingRulesetResponse) *SubResourceResponse { return v.VirtualNetworkLink }).(SubResourceResponsePtrOutput)
}

type VirtualNetworkDnsForwardingRulesetResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkDnsForwardingRulesetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkDnsForwardingRulesetResponse)(nil)).Elem()
}

func (o VirtualNetworkDnsForwardingRulesetResponseArrayOutput) ToVirtualNetworkDnsForwardingRulesetResponseArrayOutput() VirtualNetworkDnsForwardingRulesetResponseArrayOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseArrayOutput) ToVirtualNetworkDnsForwardingRulesetResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkDnsForwardingRulesetResponseArrayOutput {
	return o
}

func (o VirtualNetworkDnsForwardingRulesetResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkDnsForwardingRulesetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkDnsForwardingRulesetResponse {
		return vs[0].([]VirtualNetworkDnsForwardingRulesetResponse)[vs[1].(int)]
	}).(VirtualNetworkDnsForwardingRulesetResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(InboundEndpointIPConfigurationOutput{})
	pulumi.RegisterOutputType(InboundEndpointIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(InboundEndpointIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(InboundEndpointIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TargetDnsServerOutput{})
	pulumi.RegisterOutputType(TargetDnsServerArrayOutput{})
	pulumi.RegisterOutputType(TargetDnsServerResponseOutput{})
	pulumi.RegisterOutputType(TargetDnsServerResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkDnsForwardingRulesetResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkDnsForwardingRulesetResponseArrayOutput{})
}
