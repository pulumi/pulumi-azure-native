


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerServiceLinuxProfile struct {
	AdminUsername string                           `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfiguration `pulumi:"ssh"`
}





type ContainerServiceLinuxProfileInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfileOutput() ContainerServiceLinuxProfileOutput
	ToContainerServiceLinuxProfileOutputWithContext(context.Context) ContainerServiceLinuxProfileOutput
}

type ContainerServiceLinuxProfileArgs struct {
	AdminUsername pulumi.StringInput                    `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfigurationInput `pulumi:"ssh"`
}

func (ContainerServiceLinuxProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfile)(nil)).Elem()
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfileOutput() ContainerServiceLinuxProfileOutput {
	return i.ToContainerServiceLinuxProfileOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfileOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileOutput)
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return i.ToContainerServiceLinuxProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileOutput).ToContainerServiceLinuxProfilePtrOutputWithContext(ctx)
}









type ContainerServiceLinuxProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput
	ToContainerServiceLinuxProfilePtrOutputWithContext(context.Context) ContainerServiceLinuxProfilePtrOutput
}

type containerServiceLinuxProfilePtrType ContainerServiceLinuxProfileArgs

func ContainerServiceLinuxProfilePtr(v *ContainerServiceLinuxProfileArgs) ContainerServiceLinuxProfilePtrInput {
	return (*containerServiceLinuxProfilePtrType)(v)
}

func (*containerServiceLinuxProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfile)(nil)).Elem()
}

func (i *containerServiceLinuxProfilePtrType) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return i.ToContainerServiceLinuxProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceLinuxProfilePtrType) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfilePtrOutput)
}

type ContainerServiceLinuxProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfile)(nil)).Elem()
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfileOutput() ContainerServiceLinuxProfileOutput {
	return o
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfileOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileOutput {
	return o
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return o.ToContainerServiceLinuxProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceLinuxProfile) *ContainerServiceLinuxProfile {
		return &v
	}).(ContainerServiceLinuxProfilePtrOutput)
}

func (o ContainerServiceLinuxProfileOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfile) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ContainerServiceLinuxProfileOutput) Ssh() ContainerServiceSshConfigurationOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfile) ContainerServiceSshConfiguration { return v.Ssh }).(ContainerServiceSshConfigurationOutput)
}

type ContainerServiceLinuxProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfile)(nil)).Elem()
}

func (o ContainerServiceLinuxProfilePtrOutput) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfilePtrOutput) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfilePtrOutput) Elem() ContainerServiceLinuxProfileOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfile) ContainerServiceLinuxProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceLinuxProfile
		return ret
	}).(ContainerServiceLinuxProfileOutput)
}

func (o ContainerServiceLinuxProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfile) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceLinuxProfilePtrOutput) Ssh() ContainerServiceSshConfigurationPtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfile) *ContainerServiceSshConfiguration {
		if v == nil {
			return nil
		}
		return &v.Ssh
	}).(ContainerServiceSshConfigurationPtrOutput)
}

type ContainerServiceLinuxProfileResponse struct {
	AdminUsername string                                   `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfigurationResponse `pulumi:"ssh"`
}





type ContainerServiceLinuxProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfileResponseOutput() ContainerServiceLinuxProfileResponseOutput
	ToContainerServiceLinuxProfileResponseOutputWithContext(context.Context) ContainerServiceLinuxProfileResponseOutput
}

type ContainerServiceLinuxProfileResponseArgs struct {
	AdminUsername pulumi.StringInput                            `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfigurationResponseInput `pulumi:"ssh"`
}

func (ContainerServiceLinuxProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (i ContainerServiceLinuxProfileResponseArgs) ToContainerServiceLinuxProfileResponseOutput() ContainerServiceLinuxProfileResponseOutput {
	return i.ToContainerServiceLinuxProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileResponseArgs) ToContainerServiceLinuxProfileResponseOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileResponseOutput)
}

func (i ContainerServiceLinuxProfileResponseArgs) ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput {
	return i.ToContainerServiceLinuxProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileResponseArgs) ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileResponseOutput).ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceLinuxProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput
	ToContainerServiceLinuxProfileResponsePtrOutputWithContext(context.Context) ContainerServiceLinuxProfileResponsePtrOutput
}

type containerServiceLinuxProfileResponsePtrType ContainerServiceLinuxProfileResponseArgs

func ContainerServiceLinuxProfileResponsePtr(v *ContainerServiceLinuxProfileResponseArgs) ContainerServiceLinuxProfileResponsePtrInput {
	return (*containerServiceLinuxProfileResponsePtrType)(v)
}

func (*containerServiceLinuxProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (i *containerServiceLinuxProfileResponsePtrType) ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput {
	return i.ToContainerServiceLinuxProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceLinuxProfileResponsePtrType) ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileResponsePtrOutput)
}

type ContainerServiceLinuxProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponseOutput() ContainerServiceLinuxProfileResponseOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponseOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponseOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput {
	return o.ToContainerServiceLinuxProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceLinuxProfileResponse) *ContainerServiceLinuxProfileResponse {
		return &v
	}).(ContainerServiceLinuxProfileResponsePtrOutput)
}

func (o ContainerServiceLinuxProfileResponseOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfileResponse) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ContainerServiceLinuxProfileResponseOutput) Ssh() ContainerServiceSshConfigurationResponseOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfileResponse) ContainerServiceSshConfigurationResponse { return v.Ssh }).(ContainerServiceSshConfigurationResponseOutput)
}

type ContainerServiceLinuxProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) Elem() ContainerServiceLinuxProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfileResponse) ContainerServiceLinuxProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceLinuxProfileResponse
		return ret
	}).(ContainerServiceLinuxProfileResponseOutput)
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) Ssh() ContainerServiceSshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfileResponse) *ContainerServiceSshConfigurationResponse {
		if v == nil {
			return nil
		}
		return &v.Ssh
	}).(ContainerServiceSshConfigurationResponsePtrOutput)
}

type ContainerServiceNetworkProfile struct {
	DnsServiceIP        *string                            `pulumi:"dnsServiceIP"`
	DockerBridgeCidr    *string                            `pulumi:"dockerBridgeCidr"`
	LoadBalancerProfile *ManagedClusterLoadBalancerProfile `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     *string                            `pulumi:"loadBalancerSku"`
	NetworkPlugin       *string                            `pulumi:"networkPlugin"`
	NetworkPolicy       *string                            `pulumi:"networkPolicy"`
	OutboundType        *string                            `pulumi:"outboundType"`
	PodCidr             *string                            `pulumi:"podCidr"`
	ServiceCidr         *string                            `pulumi:"serviceCidr"`
}


func (val *ContainerServiceNetworkProfile) Defaults() *ContainerServiceNetworkProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DnsServiceIP) {
		dnsServiceIP_ := "10.0.0.10"
		tmp.DnsServiceIP = &dnsServiceIP_
	}
	if isZero(tmp.DockerBridgeCidr) {
		dockerBridgeCidr_ := "172.17.0.1/16"
		tmp.DockerBridgeCidr = &dockerBridgeCidr_
	}
	tmp.LoadBalancerProfile = tmp.LoadBalancerProfile.Defaults()

	if isZero(tmp.NetworkPlugin) {
		networkPlugin_ := "kubenet"
		tmp.NetworkPlugin = &networkPlugin_
	}
	if isZero(tmp.OutboundType) {
		outboundType_ := "loadBalancer"
		tmp.OutboundType = &outboundType_
	}
	if isZero(tmp.PodCidr) {
		podCidr_ := "10.244.0.0/16"
		tmp.PodCidr = &podCidr_
	}
	if isZero(tmp.ServiceCidr) {
		serviceCidr_ := "10.0.0.0/16"
		tmp.ServiceCidr = &serviceCidr_
	}
	return &tmp
}





type ContainerServiceNetworkProfileInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfileOutput() ContainerServiceNetworkProfileOutput
	ToContainerServiceNetworkProfileOutputWithContext(context.Context) ContainerServiceNetworkProfileOutput
}

type ContainerServiceNetworkProfileArgs struct {
	DnsServiceIP        pulumi.StringPtrInput                     `pulumi:"dnsServiceIP"`
	DockerBridgeCidr    pulumi.StringPtrInput                     `pulumi:"dockerBridgeCidr"`
	LoadBalancerProfile ManagedClusterLoadBalancerProfilePtrInput `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     pulumi.StringPtrInput                     `pulumi:"loadBalancerSku"`
	NetworkPlugin       pulumi.StringPtrInput                     `pulumi:"networkPlugin"`
	NetworkPolicy       pulumi.StringPtrInput                     `pulumi:"networkPolicy"`
	OutboundType        pulumi.StringPtrInput                     `pulumi:"outboundType"`
	PodCidr             pulumi.StringPtrInput                     `pulumi:"podCidr"`
	ServiceCidr         pulumi.StringPtrInput                     `pulumi:"serviceCidr"`
}

func (ContainerServiceNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfile)(nil)).Elem()
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfileOutput() ContainerServiceNetworkProfileOutput {
	return i.ToContainerServiceNetworkProfileOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfileOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileOutput)
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return i.ToContainerServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileOutput).ToContainerServiceNetworkProfilePtrOutputWithContext(ctx)
}









type ContainerServiceNetworkProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput
	ToContainerServiceNetworkProfilePtrOutputWithContext(context.Context) ContainerServiceNetworkProfilePtrOutput
}

type containerServiceNetworkProfilePtrType ContainerServiceNetworkProfileArgs

func ContainerServiceNetworkProfilePtr(v *ContainerServiceNetworkProfileArgs) ContainerServiceNetworkProfilePtrInput {
	return (*containerServiceNetworkProfilePtrType)(v)
}

func (*containerServiceNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfile)(nil)).Elem()
}

func (i *containerServiceNetworkProfilePtrType) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return i.ToContainerServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceNetworkProfilePtrType) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfilePtrOutput)
}

type ContainerServiceNetworkProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfile)(nil)).Elem()
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfileOutput() ContainerServiceNetworkProfileOutput {
	return o
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfileOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileOutput {
	return o
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return o.ToContainerServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceNetworkProfile) *ContainerServiceNetworkProfile {
		return &v
	}).(ContainerServiceNetworkProfilePtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) LoadBalancerProfile() ManagedClusterLoadBalancerProfilePtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *ManagedClusterLoadBalancerProfile {
		return v.LoadBalancerProfile
	}).(ManagedClusterLoadBalancerProfilePtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.NetworkPlugin }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.NetworkPolicy }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) OutboundType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.OutboundType }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

type ContainerServiceNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfile)(nil)).Elem()
}

func (o ContainerServiceNetworkProfilePtrOutput) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfilePtrOutput) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfilePtrOutput) Elem() ContainerServiceNetworkProfileOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) ContainerServiceNetworkProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceNetworkProfile
		return ret
	}).(ContainerServiceNetworkProfileOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) LoadBalancerProfile() ManagedClusterLoadBalancerProfilePtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *ManagedClusterLoadBalancerProfile {
		if v == nil {
			return nil
		}
		return v.LoadBalancerProfile
	}).(ManagedClusterLoadBalancerProfilePtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPlugin
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) OutboundType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.OutboundType
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceNetworkProfileResponse struct {
	DnsServiceIP        *string                                    `pulumi:"dnsServiceIP"`
	DockerBridgeCidr    *string                                    `pulumi:"dockerBridgeCidr"`
	LoadBalancerProfile *ManagedClusterLoadBalancerProfileResponse `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     *string                                    `pulumi:"loadBalancerSku"`
	NetworkPlugin       *string                                    `pulumi:"networkPlugin"`
	NetworkPolicy       *string                                    `pulumi:"networkPolicy"`
	OutboundType        *string                                    `pulumi:"outboundType"`
	PodCidr             *string                                    `pulumi:"podCidr"`
	ServiceCidr         *string                                    `pulumi:"serviceCidr"`
}


func (val *ContainerServiceNetworkProfileResponse) Defaults() *ContainerServiceNetworkProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DnsServiceIP) {
		dnsServiceIP_ := "10.0.0.10"
		tmp.DnsServiceIP = &dnsServiceIP_
	}
	if isZero(tmp.DockerBridgeCidr) {
		dockerBridgeCidr_ := "172.17.0.1/16"
		tmp.DockerBridgeCidr = &dockerBridgeCidr_
	}
	tmp.LoadBalancerProfile = tmp.LoadBalancerProfile.Defaults()

	if isZero(tmp.NetworkPlugin) {
		networkPlugin_ := "kubenet"
		tmp.NetworkPlugin = &networkPlugin_
	}
	if isZero(tmp.OutboundType) {
		outboundType_ := "loadBalancer"
		tmp.OutboundType = &outboundType_
	}
	if isZero(tmp.PodCidr) {
		podCidr_ := "10.244.0.0/16"
		tmp.PodCidr = &podCidr_
	}
	if isZero(tmp.ServiceCidr) {
		serviceCidr_ := "10.0.0.0/16"
		tmp.ServiceCidr = &serviceCidr_
	}
	return &tmp
}





type ContainerServiceNetworkProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfileResponseOutput() ContainerServiceNetworkProfileResponseOutput
	ToContainerServiceNetworkProfileResponseOutputWithContext(context.Context) ContainerServiceNetworkProfileResponseOutput
}

type ContainerServiceNetworkProfileResponseArgs struct {
	DnsServiceIP        pulumi.StringPtrInput                             `pulumi:"dnsServiceIP"`
	DockerBridgeCidr    pulumi.StringPtrInput                             `pulumi:"dockerBridgeCidr"`
	LoadBalancerProfile ManagedClusterLoadBalancerProfileResponsePtrInput `pulumi:"loadBalancerProfile"`
	LoadBalancerSku     pulumi.StringPtrInput                             `pulumi:"loadBalancerSku"`
	NetworkPlugin       pulumi.StringPtrInput                             `pulumi:"networkPlugin"`
	NetworkPolicy       pulumi.StringPtrInput                             `pulumi:"networkPolicy"`
	OutboundType        pulumi.StringPtrInput                             `pulumi:"outboundType"`
	PodCidr             pulumi.StringPtrInput                             `pulumi:"podCidr"`
	ServiceCidr         pulumi.StringPtrInput                             `pulumi:"serviceCidr"`
}

func (ContainerServiceNetworkProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (i ContainerServiceNetworkProfileResponseArgs) ToContainerServiceNetworkProfileResponseOutput() ContainerServiceNetworkProfileResponseOutput {
	return i.ToContainerServiceNetworkProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileResponseArgs) ToContainerServiceNetworkProfileResponseOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileResponseOutput)
}

func (i ContainerServiceNetworkProfileResponseArgs) ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput {
	return i.ToContainerServiceNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileResponseArgs) ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileResponseOutput).ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceNetworkProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput
	ToContainerServiceNetworkProfileResponsePtrOutputWithContext(context.Context) ContainerServiceNetworkProfileResponsePtrOutput
}

type containerServiceNetworkProfileResponsePtrType ContainerServiceNetworkProfileResponseArgs

func ContainerServiceNetworkProfileResponsePtr(v *ContainerServiceNetworkProfileResponseArgs) ContainerServiceNetworkProfileResponsePtrInput {
	return (*containerServiceNetworkProfileResponsePtrType)(v)
}

func (*containerServiceNetworkProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (i *containerServiceNetworkProfileResponsePtrType) ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput {
	return i.ToContainerServiceNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceNetworkProfileResponsePtrType) ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileResponsePtrOutput)
}

type ContainerServiceNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponseOutput() ContainerServiceNetworkProfileResponseOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponseOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponseOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ToContainerServiceNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceNetworkProfileResponse) *ContainerServiceNetworkProfileResponse {
		return &v
	}).(ContainerServiceNetworkProfileResponsePtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) LoadBalancerProfile() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *ManagedClusterLoadBalancerProfileResponse {
		return v.LoadBalancerProfile
	}).(ManagedClusterLoadBalancerProfileResponsePtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.NetworkPlugin }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.NetworkPolicy }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) OutboundType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.OutboundType }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

type ContainerServiceNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) Elem() ContainerServiceNetworkProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) ContainerServiceNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceNetworkProfileResponse
		return ret
	}).(ContainerServiceNetworkProfileResponseOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) LoadBalancerProfile() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *ManagedClusterLoadBalancerProfileResponse {
		if v == nil {
			return nil
		}
		return v.LoadBalancerProfile
	}).(ManagedClusterLoadBalancerProfileResponsePtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPlugin
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) OutboundType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.OutboundType
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceSshConfiguration struct {
	PublicKeys []ContainerServiceSshPublicKey `pulumi:"publicKeys"`
}





type ContainerServiceSshConfigurationInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationOutput() ContainerServiceSshConfigurationOutput
	ToContainerServiceSshConfigurationOutputWithContext(context.Context) ContainerServiceSshConfigurationOutput
}

type ContainerServiceSshConfigurationArgs struct {
	PublicKeys ContainerServiceSshPublicKeyArrayInput `pulumi:"publicKeys"`
}

func (ContainerServiceSshConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfiguration)(nil)).Elem()
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationOutput() ContainerServiceSshConfigurationOutput {
	return i.ToContainerServiceSshConfigurationOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationOutput)
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return i.ToContainerServiceSshConfigurationPtrOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationOutput).ToContainerServiceSshConfigurationPtrOutputWithContext(ctx)
}









type ContainerServiceSshConfigurationPtrInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput
	ToContainerServiceSshConfigurationPtrOutputWithContext(context.Context) ContainerServiceSshConfigurationPtrOutput
}

type containerServiceSshConfigurationPtrType ContainerServiceSshConfigurationArgs

func ContainerServiceSshConfigurationPtr(v *ContainerServiceSshConfigurationArgs) ContainerServiceSshConfigurationPtrInput {
	return (*containerServiceSshConfigurationPtrType)(v)
}

func (*containerServiceSshConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfiguration)(nil)).Elem()
}

func (i *containerServiceSshConfigurationPtrType) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return i.ToContainerServiceSshConfigurationPtrOutputWithContext(context.Background())
}

func (i *containerServiceSshConfigurationPtrType) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationPtrOutput)
}

type ContainerServiceSshConfigurationOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfiguration)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationOutput() ContainerServiceSshConfigurationOutput {
	return o
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationOutput {
	return o
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return o.ToContainerServiceSshConfigurationPtrOutputWithContext(context.Background())
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceSshConfiguration) *ContainerServiceSshConfiguration {
		return &v
	}).(ContainerServiceSshConfigurationPtrOutput)
}

func (o ContainerServiceSshConfigurationOutput) PublicKeys() ContainerServiceSshPublicKeyArrayOutput {
	return o.ApplyT(func(v ContainerServiceSshConfiguration) []ContainerServiceSshPublicKey { return v.PublicKeys }).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfiguration)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationPtrOutput) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationPtrOutput) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationPtrOutput) Elem() ContainerServiceSshConfigurationOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfiguration) ContainerServiceSshConfiguration {
		if v != nil {
			return *v
		}
		var ret ContainerServiceSshConfiguration
		return ret
	}).(ContainerServiceSshConfigurationOutput)
}

func (o ContainerServiceSshConfigurationPtrOutput) PublicKeys() ContainerServiceSshPublicKeyArrayOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfiguration) []ContainerServiceSshPublicKey {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshConfigurationResponse struct {
	PublicKeys []ContainerServiceSshPublicKeyResponse `pulumi:"publicKeys"`
}





type ContainerServiceSshConfigurationResponseInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationResponseOutput() ContainerServiceSshConfigurationResponseOutput
	ToContainerServiceSshConfigurationResponseOutputWithContext(context.Context) ContainerServiceSshConfigurationResponseOutput
}

type ContainerServiceSshConfigurationResponseArgs struct {
	PublicKeys ContainerServiceSshPublicKeyResponseArrayInput `pulumi:"publicKeys"`
}

func (ContainerServiceSshConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (i ContainerServiceSshConfigurationResponseArgs) ToContainerServiceSshConfigurationResponseOutput() ContainerServiceSshConfigurationResponseOutput {
	return i.ToContainerServiceSshConfigurationResponseOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationResponseArgs) ToContainerServiceSshConfigurationResponseOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationResponseOutput)
}

func (i ContainerServiceSshConfigurationResponseArgs) ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput {
	return i.ToContainerServiceSshConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationResponseArgs) ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationResponseOutput).ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx)
}









type ContainerServiceSshConfigurationResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput
	ToContainerServiceSshConfigurationResponsePtrOutputWithContext(context.Context) ContainerServiceSshConfigurationResponsePtrOutput
}

type containerServiceSshConfigurationResponsePtrType ContainerServiceSshConfigurationResponseArgs

func ContainerServiceSshConfigurationResponsePtr(v *ContainerServiceSshConfigurationResponseArgs) ContainerServiceSshConfigurationResponsePtrInput {
	return (*containerServiceSshConfigurationResponsePtrType)(v)
}

func (*containerServiceSshConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (i *containerServiceSshConfigurationResponsePtrType) ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput {
	return i.ToContainerServiceSshConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceSshConfigurationResponsePtrType) ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationResponsePtrOutput)
}

type ContainerServiceSshConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponseOutput() ContainerServiceSshConfigurationResponseOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponseOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponseOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput {
	return o.ToContainerServiceSshConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceSshConfigurationResponse) *ContainerServiceSshConfigurationResponse {
		return &v
	}).(ContainerServiceSshConfigurationResponsePtrOutput)
}

func (o ContainerServiceSshConfigurationResponseOutput) PublicKeys() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v ContainerServiceSshConfigurationResponse) []ContainerServiceSshPublicKeyResponse {
		return v.PublicKeys
	}).(ContainerServiceSshPublicKeyResponseArrayOutput)
}

type ContainerServiceSshConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponsePtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) Elem() ContainerServiceSshConfigurationResponseOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfigurationResponse) ContainerServiceSshConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceSshConfigurationResponse
		return ret
	}).(ContainerServiceSshConfigurationResponseOutput)
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) PublicKeys() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfigurationResponse) []ContainerServiceSshPublicKeyResponse {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(ContainerServiceSshPublicKeyResponseArrayOutput)
}

type ContainerServiceSshPublicKey struct {
	KeyData string `pulumi:"keyData"`
}





type ContainerServiceSshPublicKeyInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyOutput() ContainerServiceSshPublicKeyOutput
	ToContainerServiceSshPublicKeyOutputWithContext(context.Context) ContainerServiceSshPublicKeyOutput
}

type ContainerServiceSshPublicKeyArgs struct {
	KeyData pulumi.StringInput `pulumi:"keyData"`
}

func (ContainerServiceSshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKey)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyArgs) ToContainerServiceSshPublicKeyOutput() ContainerServiceSshPublicKeyOutput {
	return i.ToContainerServiceSshPublicKeyOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyArgs) ToContainerServiceSshPublicKeyOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyOutput)
}





type ContainerServiceSshPublicKeyArrayInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyArrayOutput() ContainerServiceSshPublicKeyArrayOutput
	ToContainerServiceSshPublicKeyArrayOutputWithContext(context.Context) ContainerServiceSshPublicKeyArrayOutput
}

type ContainerServiceSshPublicKeyArray []ContainerServiceSshPublicKeyInput

func (ContainerServiceSshPublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKey)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyArray) ToContainerServiceSshPublicKeyArrayOutput() ContainerServiceSshPublicKeyArrayOutput {
	return i.ToContainerServiceSshPublicKeyArrayOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyArray) ToContainerServiceSshPublicKeyArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshPublicKeyOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKey)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyOutput) ToContainerServiceSshPublicKeyOutput() ContainerServiceSshPublicKeyOutput {
	return o
}

func (o ContainerServiceSshPublicKeyOutput) ToContainerServiceSshPublicKeyOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyOutput {
	return o
}

func (o ContainerServiceSshPublicKeyOutput) KeyData() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceSshPublicKey) string { return v.KeyData }).(pulumi.StringOutput)
}

type ContainerServiceSshPublicKeyArrayOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKey)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyArrayOutput) ToContainerServiceSshPublicKeyArrayOutput() ContainerServiceSshPublicKeyArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyArrayOutput) ToContainerServiceSshPublicKeyArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyArrayOutput) Index(i pulumi.IntInput) ContainerServiceSshPublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerServiceSshPublicKey {
		return vs[0].([]ContainerServiceSshPublicKey)[vs[1].(int)]
	}).(ContainerServiceSshPublicKeyOutput)
}

type ContainerServiceSshPublicKeyResponse struct {
	KeyData string `pulumi:"keyData"`
}





type ContainerServiceSshPublicKeyResponseInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyResponseOutput() ContainerServiceSshPublicKeyResponseOutput
	ToContainerServiceSshPublicKeyResponseOutputWithContext(context.Context) ContainerServiceSshPublicKeyResponseOutput
}

type ContainerServiceSshPublicKeyResponseArgs struct {
	KeyData pulumi.StringInput `pulumi:"keyData"`
}

func (ContainerServiceSshPublicKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyResponseArgs) ToContainerServiceSshPublicKeyResponseOutput() ContainerServiceSshPublicKeyResponseOutput {
	return i.ToContainerServiceSshPublicKeyResponseOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyResponseArgs) ToContainerServiceSshPublicKeyResponseOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyResponseOutput)
}





type ContainerServiceSshPublicKeyResponseArrayInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyResponseArrayOutput() ContainerServiceSshPublicKeyResponseArrayOutput
	ToContainerServiceSshPublicKeyResponseArrayOutputWithContext(context.Context) ContainerServiceSshPublicKeyResponseArrayOutput
}

type ContainerServiceSshPublicKeyResponseArray []ContainerServiceSshPublicKeyResponseInput

func (ContainerServiceSshPublicKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyResponseArray) ToContainerServiceSshPublicKeyResponseArrayOutput() ContainerServiceSshPublicKeyResponseArrayOutput {
	return i.ToContainerServiceSshPublicKeyResponseArrayOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyResponseArray) ToContainerServiceSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyResponseArrayOutput)
}

type ContainerServiceSshPublicKeyResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyResponseOutput) ToContainerServiceSshPublicKeyResponseOutput() ContainerServiceSshPublicKeyResponseOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseOutput) ToContainerServiceSshPublicKeyResponseOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseOutput) KeyData() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceSshPublicKeyResponse) string { return v.KeyData }).(pulumi.StringOutput)
}

type ContainerServiceSshPublicKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyResponseArrayOutput) ToContainerServiceSshPublicKeyResponseArrayOutput() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseArrayOutput) ToContainerServiceSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseArrayOutput) Index(i pulumi.IntInput) ContainerServiceSshPublicKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerServiceSshPublicKeyResponse {
		return vs[0].([]ContainerServiceSshPublicKeyResponse)[vs[1].(int)]
	}).(ContainerServiceSshPublicKeyResponseOutput)
}

type CredentialResultResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type CredentialResultResponseInput interface {
	pulumi.Input

	ToCredentialResultResponseOutput() CredentialResultResponseOutput
	ToCredentialResultResponseOutputWithContext(context.Context) CredentialResultResponseOutput
}

type CredentialResultResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (CredentialResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialResultResponse)(nil)).Elem()
}

func (i CredentialResultResponseArgs) ToCredentialResultResponseOutput() CredentialResultResponseOutput {
	return i.ToCredentialResultResponseOutputWithContext(context.Background())
}

func (i CredentialResultResponseArgs) ToCredentialResultResponseOutputWithContext(ctx context.Context) CredentialResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialResultResponseOutput)
}





type CredentialResultResponseArrayInput interface {
	pulumi.Input

	ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput
	ToCredentialResultResponseArrayOutputWithContext(context.Context) CredentialResultResponseArrayOutput
}

type CredentialResultResponseArray []CredentialResultResponseInput

func (CredentialResultResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CredentialResultResponse)(nil)).Elem()
}

func (i CredentialResultResponseArray) ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput {
	return i.ToCredentialResultResponseArrayOutputWithContext(context.Background())
}

func (i CredentialResultResponseArray) ToCredentialResultResponseArrayOutputWithContext(ctx context.Context) CredentialResultResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialResultResponseArrayOutput)
}

type CredentialResultResponseOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutput() CredentialResultResponseOutput {
	return o
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutputWithContext(ctx context.Context) CredentialResultResponseOutput {
	return o
}

func (o CredentialResultResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CredentialResultResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Value }).(pulumi.StringOutput)
}

type CredentialResultResponseArrayOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutputWithContext(ctx context.Context) CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) Index(i pulumi.IntInput) CredentialResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CredentialResultResponse {
		return vs[0].([]CredentialResultResponse)[vs[1].(int)]
	}).(CredentialResultResponseOutput)
}

type ManagedClusterAADProfile struct {
	ClientAppID     string  `pulumi:"clientAppID"`
	ServerAppID     string  `pulumi:"serverAppID"`
	ServerAppSecret *string `pulumi:"serverAppSecret"`
	TenantID        *string `pulumi:"tenantID"`
}





type ManagedClusterAADProfileInput interface {
	pulumi.Input

	ToManagedClusterAADProfileOutput() ManagedClusterAADProfileOutput
	ToManagedClusterAADProfileOutputWithContext(context.Context) ManagedClusterAADProfileOutput
}

type ManagedClusterAADProfileArgs struct {
	ClientAppID     pulumi.StringInput    `pulumi:"clientAppID"`
	ServerAppID     pulumi.StringInput    `pulumi:"serverAppID"`
	ServerAppSecret pulumi.StringPtrInput `pulumi:"serverAppSecret"`
	TenantID        pulumi.StringPtrInput `pulumi:"tenantID"`
}

func (ManagedClusterAADProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfile)(nil)).Elem()
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfileOutput() ManagedClusterAADProfileOutput {
	return i.ToManagedClusterAADProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfileOutputWithContext(ctx context.Context) ManagedClusterAADProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileOutput)
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return i.ToManagedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileOutput).ToManagedClusterAADProfilePtrOutputWithContext(ctx)
}









type ManagedClusterAADProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput
	ToManagedClusterAADProfilePtrOutputWithContext(context.Context) ManagedClusterAADProfilePtrOutput
}

type managedClusterAADProfilePtrType ManagedClusterAADProfileArgs

func ManagedClusterAADProfilePtr(v *ManagedClusterAADProfileArgs) ManagedClusterAADProfilePtrInput {
	return (*managedClusterAADProfilePtrType)(v)
}

func (*managedClusterAADProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfile)(nil)).Elem()
}

func (i *managedClusterAADProfilePtrType) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return i.ToManagedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterAADProfilePtrType) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfilePtrOutput)
}

type ManagedClusterAADProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfile)(nil)).Elem()
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfileOutput() ManagedClusterAADProfileOutput {
	return o
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfileOutputWithContext(ctx context.Context) ManagedClusterAADProfileOutput {
	return o
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return o.ToManagedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAADProfile) *ManagedClusterAADProfile {
		return &v
	}).(ManagedClusterAADProfilePtrOutput)
}

func (o ManagedClusterAADProfileOutput) ClientAppID() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) string { return v.ClientAppID }).(pulumi.StringOutput)
}

func (o ManagedClusterAADProfileOutput) ServerAppID() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) string { return v.ServerAppID }).(pulumi.StringOutput)
}

func (o ManagedClusterAADProfileOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *string { return v.ServerAppSecret }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfile)(nil)).Elem()
}

func (o ManagedClusterAADProfilePtrOutput) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return o
}

func (o ManagedClusterAADProfilePtrOutput) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return o
}

func (o ManagedClusterAADProfilePtrOutput) Elem() ManagedClusterAADProfileOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) ManagedClusterAADProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAADProfile
		return ret
	}).(ManagedClusterAADProfileOutput)
}

func (o ManagedClusterAADProfilePtrOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ClientAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ServerAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppSecret
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return v.TenantID
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfileResponse struct {
	ClientAppID     string  `pulumi:"clientAppID"`
	ServerAppID     string  `pulumi:"serverAppID"`
	ServerAppSecret *string `pulumi:"serverAppSecret"`
	TenantID        *string `pulumi:"tenantID"`
}





type ManagedClusterAADProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterAADProfileResponseOutput() ManagedClusterAADProfileResponseOutput
	ToManagedClusterAADProfileResponseOutputWithContext(context.Context) ManagedClusterAADProfileResponseOutput
}

type ManagedClusterAADProfileResponseArgs struct {
	ClientAppID     pulumi.StringInput    `pulumi:"clientAppID"`
	ServerAppID     pulumi.StringInput    `pulumi:"serverAppID"`
	ServerAppSecret pulumi.StringPtrInput `pulumi:"serverAppSecret"`
	TenantID        pulumi.StringPtrInput `pulumi:"tenantID"`
}

func (ManagedClusterAADProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAADProfileResponseArgs) ToManagedClusterAADProfileResponseOutput() ManagedClusterAADProfileResponseOutput {
	return i.ToManagedClusterAADProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileResponseArgs) ToManagedClusterAADProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileResponseOutput)
}

func (i ManagedClusterAADProfileResponseArgs) ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput {
	return i.ToManagedClusterAADProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileResponseArgs) ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileResponseOutput).ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterAADProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput
	ToManagedClusterAADProfileResponsePtrOutputWithContext(context.Context) ManagedClusterAADProfileResponsePtrOutput
}

type managedClusterAADProfileResponsePtrType ManagedClusterAADProfileResponseArgs

func ManagedClusterAADProfileResponsePtr(v *ManagedClusterAADProfileResponseArgs) ManagedClusterAADProfileResponsePtrInput {
	return (*managedClusterAADProfileResponsePtrType)(v)
}

func (*managedClusterAADProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (i *managedClusterAADProfileResponsePtrType) ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput {
	return i.ToManagedClusterAADProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterAADProfileResponsePtrType) ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileResponsePtrOutput)
}

type ManagedClusterAADProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponseOutput() ManagedClusterAADProfileResponseOutput {
	return o
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponseOutput {
	return o
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput {
	return o.ToManagedClusterAADProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAADProfileResponse) *ManagedClusterAADProfileResponse {
		return &v
	}).(ManagedClusterAADProfileResponsePtrOutput)
}

func (o ManagedClusterAADProfileResponseOutput) ClientAppID() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) string { return v.ClientAppID }).(pulumi.StringOutput)
}

func (o ManagedClusterAADProfileResponseOutput) ServerAppID() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) string { return v.ServerAppID }).(pulumi.StringOutput)
}

func (o ManagedClusterAADProfileResponseOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *string { return v.ServerAppSecret }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponseOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAADProfileResponsePtrOutput) ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAADProfileResponsePtrOutput) ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAADProfileResponsePtrOutput) Elem() ManagedClusterAADProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) ManagedClusterAADProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAADProfileResponse
		return ret
	}).(ManagedClusterAADProfileResponseOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServerAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppSecret
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantID
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAPIServerAccessProfile struct {
	AuthorizedIPRanges   []string `pulumi:"authorizedIPRanges"`
	EnablePrivateCluster *bool    `pulumi:"enablePrivateCluster"`
}





type ManagedClusterAPIServerAccessProfileInput interface {
	pulumi.Input

	ToManagedClusterAPIServerAccessProfileOutput() ManagedClusterAPIServerAccessProfileOutput
	ToManagedClusterAPIServerAccessProfileOutputWithContext(context.Context) ManagedClusterAPIServerAccessProfileOutput
}

type ManagedClusterAPIServerAccessProfileArgs struct {
	AuthorizedIPRanges   pulumi.StringArrayInput `pulumi:"authorizedIPRanges"`
	EnablePrivateCluster pulumi.BoolPtrInput     `pulumi:"enablePrivateCluster"`
}

func (ManagedClusterAPIServerAccessProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAPIServerAccessProfile)(nil)).Elem()
}

func (i ManagedClusterAPIServerAccessProfileArgs) ToManagedClusterAPIServerAccessProfileOutput() ManagedClusterAPIServerAccessProfileOutput {
	return i.ToManagedClusterAPIServerAccessProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAPIServerAccessProfileArgs) ToManagedClusterAPIServerAccessProfileOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileOutput)
}

func (i ManagedClusterAPIServerAccessProfileArgs) ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput {
	return i.ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAPIServerAccessProfileArgs) ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileOutput).ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx)
}









type ManagedClusterAPIServerAccessProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput
	ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(context.Context) ManagedClusterAPIServerAccessProfilePtrOutput
}

type managedClusterAPIServerAccessProfilePtrType ManagedClusterAPIServerAccessProfileArgs

func ManagedClusterAPIServerAccessProfilePtr(v *ManagedClusterAPIServerAccessProfileArgs) ManagedClusterAPIServerAccessProfilePtrInput {
	return (*managedClusterAPIServerAccessProfilePtrType)(v)
}

func (*managedClusterAPIServerAccessProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAPIServerAccessProfile)(nil)).Elem()
}

func (i *managedClusterAPIServerAccessProfilePtrType) ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput {
	return i.ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterAPIServerAccessProfilePtrType) ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfilePtrOutput)
}

type ManagedClusterAPIServerAccessProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAPIServerAccessProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAPIServerAccessProfile)(nil)).Elem()
}

func (o ManagedClusterAPIServerAccessProfileOutput) ToManagedClusterAPIServerAccessProfileOutput() ManagedClusterAPIServerAccessProfileOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileOutput) ToManagedClusterAPIServerAccessProfileOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileOutput) ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput {
	return o.ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAPIServerAccessProfileOutput) ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAPIServerAccessProfile) *ManagedClusterAPIServerAccessProfile {
		return &v
	}).(ManagedClusterAPIServerAccessProfilePtrOutput)
}

func (o ManagedClusterAPIServerAccessProfileOutput) AuthorizedIPRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfile) []string { return v.AuthorizedIPRanges }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAPIServerAccessProfileOutput) EnablePrivateCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfile) *bool { return v.EnablePrivateCluster }).(pulumi.BoolPtrOutput)
}

type ManagedClusterAPIServerAccessProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAPIServerAccessProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAPIServerAccessProfile)(nil)).Elem()
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) ToManagedClusterAPIServerAccessProfilePtrOutput() ManagedClusterAPIServerAccessProfilePtrOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) ToManagedClusterAPIServerAccessProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfilePtrOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) Elem() ManagedClusterAPIServerAccessProfileOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfile) ManagedClusterAPIServerAccessProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAPIServerAccessProfile
		return ret
	}).(ManagedClusterAPIServerAccessProfileOutput)
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) AuthorizedIPRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfile) []string {
		if v == nil {
			return nil
		}
		return v.AuthorizedIPRanges
	}).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAPIServerAccessProfilePtrOutput) EnablePrivateCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePrivateCluster
	}).(pulumi.BoolPtrOutput)
}

type ManagedClusterAPIServerAccessProfileResponse struct {
	AuthorizedIPRanges   []string `pulumi:"authorizedIPRanges"`
	EnablePrivateCluster *bool    `pulumi:"enablePrivateCluster"`
}





type ManagedClusterAPIServerAccessProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterAPIServerAccessProfileResponseOutput() ManagedClusterAPIServerAccessProfileResponseOutput
	ToManagedClusterAPIServerAccessProfileResponseOutputWithContext(context.Context) ManagedClusterAPIServerAccessProfileResponseOutput
}

type ManagedClusterAPIServerAccessProfileResponseArgs struct {
	AuthorizedIPRanges   pulumi.StringArrayInput `pulumi:"authorizedIPRanges"`
	EnablePrivateCluster pulumi.BoolPtrInput     `pulumi:"enablePrivateCluster"`
}

func (ManagedClusterAPIServerAccessProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAPIServerAccessProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAPIServerAccessProfileResponseArgs) ToManagedClusterAPIServerAccessProfileResponseOutput() ManagedClusterAPIServerAccessProfileResponseOutput {
	return i.ToManagedClusterAPIServerAccessProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterAPIServerAccessProfileResponseArgs) ToManagedClusterAPIServerAccessProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileResponseOutput)
}

func (i ManagedClusterAPIServerAccessProfileResponseArgs) ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return i.ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAPIServerAccessProfileResponseArgs) ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileResponseOutput).ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterAPIServerAccessProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput
	ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput
}

type managedClusterAPIServerAccessProfileResponsePtrType ManagedClusterAPIServerAccessProfileResponseArgs

func ManagedClusterAPIServerAccessProfileResponsePtr(v *ManagedClusterAPIServerAccessProfileResponseArgs) ManagedClusterAPIServerAccessProfileResponsePtrInput {
	return (*managedClusterAPIServerAccessProfileResponsePtrType)(v)
}

func (*managedClusterAPIServerAccessProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAPIServerAccessProfileResponse)(nil)).Elem()
}

func (i *managedClusterAPIServerAccessProfileResponsePtrType) ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return i.ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterAPIServerAccessProfileResponsePtrType) ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAPIServerAccessProfileResponsePtrOutput)
}

type ManagedClusterAPIServerAccessProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAPIServerAccessProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAPIServerAccessProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) ToManagedClusterAPIServerAccessProfileResponseOutput() ManagedClusterAPIServerAccessProfileResponseOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) ToManagedClusterAPIServerAccessProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponseOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o.ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAPIServerAccessProfileResponse) *ManagedClusterAPIServerAccessProfileResponse {
		return &v
	}).(ManagedClusterAPIServerAccessProfileResponsePtrOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) AuthorizedIPRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfileResponse) []string { return v.AuthorizedIPRanges }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponseOutput) EnablePrivateCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAPIServerAccessProfileResponse) *bool { return v.EnablePrivateCluster }).(pulumi.BoolPtrOutput)
}

type ManagedClusterAPIServerAccessProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAPIServerAccessProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAPIServerAccessProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) ToManagedClusterAPIServerAccessProfileResponsePtrOutput() ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) ToManagedClusterAPIServerAccessProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAPIServerAccessProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) Elem() ManagedClusterAPIServerAccessProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfileResponse) ManagedClusterAPIServerAccessProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAPIServerAccessProfileResponse
		return ret
	}).(ManagedClusterAPIServerAccessProfileResponseOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) AuthorizedIPRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.AuthorizedIPRanges
	}).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAPIServerAccessProfileResponsePtrOutput) EnablePrivateCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAPIServerAccessProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePrivateCluster
	}).(pulumi.BoolPtrOutput)
}

type ManagedClusterAddonProfile struct {
	Config  map[string]string `pulumi:"config"`
	Enabled bool              `pulumi:"enabled"`
}





type ManagedClusterAddonProfileInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileOutput() ManagedClusterAddonProfileOutput
	ToManagedClusterAddonProfileOutputWithContext(context.Context) ManagedClusterAddonProfileOutput
}

type ManagedClusterAddonProfileArgs struct {
	Config  pulumi.StringMapInput `pulumi:"config"`
	Enabled pulumi.BoolInput      `pulumi:"enabled"`
}

func (ManagedClusterAddonProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfile)(nil)).Elem()
}

func (i ManagedClusterAddonProfileArgs) ToManagedClusterAddonProfileOutput() ManagedClusterAddonProfileOutput {
	return i.ToManagedClusterAddonProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileArgs) ToManagedClusterAddonProfileOutputWithContext(ctx context.Context) ManagedClusterAddonProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileOutput)
}





type ManagedClusterAddonProfileMapInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileMapOutput() ManagedClusterAddonProfileMapOutput
	ToManagedClusterAddonProfileMapOutputWithContext(context.Context) ManagedClusterAddonProfileMapOutput
}

type ManagedClusterAddonProfileMap map[string]ManagedClusterAddonProfileInput

func (ManagedClusterAddonProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfile)(nil)).Elem()
}

func (i ManagedClusterAddonProfileMap) ToManagedClusterAddonProfileMapOutput() ManagedClusterAddonProfileMapOutput {
	return i.ToManagedClusterAddonProfileMapOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileMap) ToManagedClusterAddonProfileMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileMapOutput)
}

type ManagedClusterAddonProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfile)(nil)).Elem()
}

func (o ManagedClusterAddonProfileOutput) ToManagedClusterAddonProfileOutput() ManagedClusterAddonProfileOutput {
	return o
}

func (o ManagedClusterAddonProfileOutput) ToManagedClusterAddonProfileOutputWithContext(ctx context.Context) ManagedClusterAddonProfileOutput {
	return o
}

func (o ManagedClusterAddonProfileOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfile) map[string]string { return v.Config }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAddonProfileOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfile) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ManagedClusterAddonProfileMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfile)(nil)).Elem()
}

func (o ManagedClusterAddonProfileMapOutput) ToManagedClusterAddonProfileMapOutput() ManagedClusterAddonProfileMapOutput {
	return o
}

func (o ManagedClusterAddonProfileMapOutput) ToManagedClusterAddonProfileMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileMapOutput {
	return o
}

func (o ManagedClusterAddonProfileMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterAddonProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterAddonProfile {
		return vs[0].(map[string]ManagedClusterAddonProfile)[vs[1].(string)]
	}).(ManagedClusterAddonProfileOutput)
}

type ManagedClusterAddonProfileResponse struct {
	Config   map[string]string                          `pulumi:"config"`
	Enabled  bool                                       `pulumi:"enabled"`
	Identity ManagedClusterAddonProfileResponseIdentity `pulumi:"identity"`
}





type ManagedClusterAddonProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileResponseOutput() ManagedClusterAddonProfileResponseOutput
	ToManagedClusterAddonProfileResponseOutputWithContext(context.Context) ManagedClusterAddonProfileResponseOutput
}

type ManagedClusterAddonProfileResponseArgs struct {
	Config   pulumi.StringMapInput                           `pulumi:"config"`
	Enabled  pulumi.BoolInput                                `pulumi:"enabled"`
	Identity ManagedClusterAddonProfileResponseIdentityInput `pulumi:"identity"`
}

func (ManagedClusterAddonProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAddonProfileResponseArgs) ToManagedClusterAddonProfileResponseOutput() ManagedClusterAddonProfileResponseOutput {
	return i.ToManagedClusterAddonProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileResponseArgs) ToManagedClusterAddonProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileResponseOutput)
}





type ManagedClusterAddonProfileResponseMapInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileResponseMapOutput() ManagedClusterAddonProfileResponseMapOutput
	ToManagedClusterAddonProfileResponseMapOutputWithContext(context.Context) ManagedClusterAddonProfileResponseMapOutput
}

type ManagedClusterAddonProfileResponseMap map[string]ManagedClusterAddonProfileResponseInput

func (ManagedClusterAddonProfileResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAddonProfileResponseMap) ToManagedClusterAddonProfileResponseMapOutput() ManagedClusterAddonProfileResponseMapOutput {
	return i.ToManagedClusterAddonProfileResponseMapOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileResponseMap) ToManagedClusterAddonProfileResponseMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileResponseMapOutput)
}

type ManagedClusterAddonProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAddonProfileResponseOutput) ToManagedClusterAddonProfileResponseOutput() ManagedClusterAddonProfileResponseOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseOutput) ToManagedClusterAddonProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponse) map[string]string { return v.Config }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAddonProfileResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ManagedClusterAddonProfileResponseOutput) Identity() ManagedClusterAddonProfileResponseIdentityOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponse) ManagedClusterAddonProfileResponseIdentity {
		return v.Identity
	}).(ManagedClusterAddonProfileResponseIdentityOutput)
}

type ManagedClusterAddonProfileResponseMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAddonProfileResponseMapOutput) ToManagedClusterAddonProfileResponseMapOutput() ManagedClusterAddonProfileResponseMapOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseMapOutput) ToManagedClusterAddonProfileResponseMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseMapOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterAddonProfileResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterAddonProfileResponse {
		return vs[0].(map[string]ManagedClusterAddonProfileResponse)[vs[1].(string)]
	}).(ManagedClusterAddonProfileResponseOutput)
}

type ManagedClusterAddonProfileResponseIdentity struct {
	ClientId   *string `pulumi:"clientId"`
	ObjectId   *string `pulumi:"objectId"`
	ResourceId *string `pulumi:"resourceId"`
}





type ManagedClusterAddonProfileResponseIdentityInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileResponseIdentityOutput() ManagedClusterAddonProfileResponseIdentityOutput
	ToManagedClusterAddonProfileResponseIdentityOutputWithContext(context.Context) ManagedClusterAddonProfileResponseIdentityOutput
}

type ManagedClusterAddonProfileResponseIdentityArgs struct {
	ClientId   pulumi.StringPtrInput `pulumi:"clientId"`
	ObjectId   pulumi.StringPtrInput `pulumi:"objectId"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ManagedClusterAddonProfileResponseIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfileResponseIdentity)(nil)).Elem()
}

func (i ManagedClusterAddonProfileResponseIdentityArgs) ToManagedClusterAddonProfileResponseIdentityOutput() ManagedClusterAddonProfileResponseIdentityOutput {
	return i.ToManagedClusterAddonProfileResponseIdentityOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileResponseIdentityArgs) ToManagedClusterAddonProfileResponseIdentityOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileResponseIdentityOutput)
}

type ManagedClusterAddonProfileResponseIdentityOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfileResponseIdentity)(nil)).Elem()
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ToManagedClusterAddonProfileResponseIdentityOutput() ManagedClusterAddonProfileResponseIdentityOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ToManagedClusterAddonProfileResponseIdentityOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseIdentityOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponseIdentity) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponseIdentity) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAddonProfileResponseIdentityOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponseIdentity) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ManagedClusterAgentPoolProfile struct {
	AvailabilityZones      []string          `pulumi:"availabilityZones"`
	Count                  *int              `pulumi:"count"`
	EnableAutoScaling      *bool             `pulumi:"enableAutoScaling"`
	EnableNodePublicIP     *bool             `pulumi:"enableNodePublicIP"`
	MaxCount               *int              `pulumi:"maxCount"`
	MaxPods                *int              `pulumi:"maxPods"`
	MinCount               *int              `pulumi:"minCount"`
	Name                   string            `pulumi:"name"`
	NodeLabels             map[string]string `pulumi:"nodeLabels"`
	NodeTaints             []string          `pulumi:"nodeTaints"`
	OrchestratorVersion    *string           `pulumi:"orchestratorVersion"`
	OsDiskSizeGB           *int              `pulumi:"osDiskSizeGB"`
	OsType                 *string           `pulumi:"osType"`
	ScaleSetEvictionPolicy *string           `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority       *string           `pulumi:"scaleSetPriority"`
	Tags                   map[string]string `pulumi:"tags"`
	Type                   *string           `pulumi:"type"`
	VmSize                 *string           `pulumi:"vmSize"`
	VnetSubnetID           *string           `pulumi:"vnetSubnetID"`
}





type ManagedClusterAgentPoolProfileInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileOutput() ManagedClusterAgentPoolProfileOutput
	ToManagedClusterAgentPoolProfileOutputWithContext(context.Context) ManagedClusterAgentPoolProfileOutput
}

type ManagedClusterAgentPoolProfileArgs struct {
	AvailabilityZones      pulumi.StringArrayInput `pulumi:"availabilityZones"`
	Count                  pulumi.IntPtrInput      `pulumi:"count"`
	EnableAutoScaling      pulumi.BoolPtrInput     `pulumi:"enableAutoScaling"`
	EnableNodePublicIP     pulumi.BoolPtrInput     `pulumi:"enableNodePublicIP"`
	MaxCount               pulumi.IntPtrInput      `pulumi:"maxCount"`
	MaxPods                pulumi.IntPtrInput      `pulumi:"maxPods"`
	MinCount               pulumi.IntPtrInput      `pulumi:"minCount"`
	Name                   pulumi.StringInput      `pulumi:"name"`
	NodeLabels             pulumi.StringMapInput   `pulumi:"nodeLabels"`
	NodeTaints             pulumi.StringArrayInput `pulumi:"nodeTaints"`
	OrchestratorVersion    pulumi.StringPtrInput   `pulumi:"orchestratorVersion"`
	OsDiskSizeGB           pulumi.IntPtrInput      `pulumi:"osDiskSizeGB"`
	OsType                 pulumi.StringPtrInput   `pulumi:"osType"`
	ScaleSetEvictionPolicy pulumi.StringPtrInput   `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority       pulumi.StringPtrInput   `pulumi:"scaleSetPriority"`
	Tags                   pulumi.StringMapInput   `pulumi:"tags"`
	Type                   pulumi.StringPtrInput   `pulumi:"type"`
	VmSize                 pulumi.StringPtrInput   `pulumi:"vmSize"`
	VnetSubnetID           pulumi.StringPtrInput   `pulumi:"vnetSubnetID"`
}

func (ManagedClusterAgentPoolProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileArgs) ToManagedClusterAgentPoolProfileOutput() ManagedClusterAgentPoolProfileOutput {
	return i.ToManagedClusterAgentPoolProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileArgs) ToManagedClusterAgentPoolProfileOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileOutput)
}





type ManagedClusterAgentPoolProfileArrayInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileArrayOutput() ManagedClusterAgentPoolProfileArrayOutput
	ToManagedClusterAgentPoolProfileArrayOutputWithContext(context.Context) ManagedClusterAgentPoolProfileArrayOutput
}

type ManagedClusterAgentPoolProfileArray []ManagedClusterAgentPoolProfileInput

func (ManagedClusterAgentPoolProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileArray) ToManagedClusterAgentPoolProfileArrayOutput() ManagedClusterAgentPoolProfileArrayOutput {
	return i.ToManagedClusterAgentPoolProfileArrayOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileArray) ToManagedClusterAgentPoolProfileArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileArrayOutput)
}

type ManagedClusterAgentPoolProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileOutput) ToManagedClusterAgentPoolProfileOutput() ManagedClusterAgentPoolProfileOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileOutput) ToManagedClusterAgentPoolProfileOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *bool { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *bool { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAgentPoolProfileArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileArrayOutput) ToManagedClusterAgentPoolProfileArrayOutput() ManagedClusterAgentPoolProfileArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileArrayOutput) ToManagedClusterAgentPoolProfileArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileArrayOutput) Index(i pulumi.IntInput) ManagedClusterAgentPoolProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterAgentPoolProfile {
		return vs[0].([]ManagedClusterAgentPoolProfile)[vs[1].(int)]
	}).(ManagedClusterAgentPoolProfileOutput)
}

type ManagedClusterAgentPoolProfileResponse struct {
	AvailabilityZones      []string          `pulumi:"availabilityZones"`
	Count                  *int              `pulumi:"count"`
	EnableAutoScaling      *bool             `pulumi:"enableAutoScaling"`
	EnableNodePublicIP     *bool             `pulumi:"enableNodePublicIP"`
	MaxCount               *int              `pulumi:"maxCount"`
	MaxPods                *int              `pulumi:"maxPods"`
	MinCount               *int              `pulumi:"minCount"`
	Name                   string            `pulumi:"name"`
	NodeLabels             map[string]string `pulumi:"nodeLabels"`
	NodeTaints             []string          `pulumi:"nodeTaints"`
	OrchestratorVersion    *string           `pulumi:"orchestratorVersion"`
	OsDiskSizeGB           *int              `pulumi:"osDiskSizeGB"`
	OsType                 *string           `pulumi:"osType"`
	ProvisioningState      string            `pulumi:"provisioningState"`
	ScaleSetEvictionPolicy *string           `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority       *string           `pulumi:"scaleSetPriority"`
	Tags                   map[string]string `pulumi:"tags"`
	Type                   *string           `pulumi:"type"`
	VmSize                 *string           `pulumi:"vmSize"`
	VnetSubnetID           *string           `pulumi:"vnetSubnetID"`
}





type ManagedClusterAgentPoolProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileResponseOutput() ManagedClusterAgentPoolProfileResponseOutput
	ToManagedClusterAgentPoolProfileResponseOutputWithContext(context.Context) ManagedClusterAgentPoolProfileResponseOutput
}

type ManagedClusterAgentPoolProfileResponseArgs struct {
	AvailabilityZones      pulumi.StringArrayInput `pulumi:"availabilityZones"`
	Count                  pulumi.IntPtrInput      `pulumi:"count"`
	EnableAutoScaling      pulumi.BoolPtrInput     `pulumi:"enableAutoScaling"`
	EnableNodePublicIP     pulumi.BoolPtrInput     `pulumi:"enableNodePublicIP"`
	MaxCount               pulumi.IntPtrInput      `pulumi:"maxCount"`
	MaxPods                pulumi.IntPtrInput      `pulumi:"maxPods"`
	MinCount               pulumi.IntPtrInput      `pulumi:"minCount"`
	Name                   pulumi.StringInput      `pulumi:"name"`
	NodeLabels             pulumi.StringMapInput   `pulumi:"nodeLabels"`
	NodeTaints             pulumi.StringArrayInput `pulumi:"nodeTaints"`
	OrchestratorVersion    pulumi.StringPtrInput   `pulumi:"orchestratorVersion"`
	OsDiskSizeGB           pulumi.IntPtrInput      `pulumi:"osDiskSizeGB"`
	OsType                 pulumi.StringPtrInput   `pulumi:"osType"`
	ProvisioningState      pulumi.StringInput      `pulumi:"provisioningState"`
	ScaleSetEvictionPolicy pulumi.StringPtrInput   `pulumi:"scaleSetEvictionPolicy"`
	ScaleSetPriority       pulumi.StringPtrInput   `pulumi:"scaleSetPriority"`
	Tags                   pulumi.StringMapInput   `pulumi:"tags"`
	Type                   pulumi.StringPtrInput   `pulumi:"type"`
	VmSize                 pulumi.StringPtrInput   `pulumi:"vmSize"`
	VnetSubnetID           pulumi.StringPtrInput   `pulumi:"vnetSubnetID"`
}

func (ManagedClusterAgentPoolProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileResponseArgs) ToManagedClusterAgentPoolProfileResponseOutput() ManagedClusterAgentPoolProfileResponseOutput {
	return i.ToManagedClusterAgentPoolProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileResponseArgs) ToManagedClusterAgentPoolProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileResponseOutput)
}





type ManagedClusterAgentPoolProfileResponseArrayInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileResponseArrayOutput() ManagedClusterAgentPoolProfileResponseArrayOutput
	ToManagedClusterAgentPoolProfileResponseArrayOutputWithContext(context.Context) ManagedClusterAgentPoolProfileResponseArrayOutput
}

type ManagedClusterAgentPoolProfileResponseArray []ManagedClusterAgentPoolProfileResponseInput

func (ManagedClusterAgentPoolProfileResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileResponseArray) ToManagedClusterAgentPoolProfileResponseArrayOutput() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return i.ToManagedClusterAgentPoolProfileResponseArrayOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileResponseArray) ToManagedClusterAgentPoolProfileResponseArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileResponseArrayOutput)
}

type ManagedClusterAgentPoolProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ToManagedClusterAgentPoolProfileResponseOutput() ManagedClusterAgentPoolProfileResponseOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ToManagedClusterAgentPoolProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *bool { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) EnableNodePublicIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *bool { return v.EnableNodePublicIP }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ScaleSetEvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.ScaleSetEvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ScaleSetPriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.ScaleSetPriority }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAgentPoolProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileResponseArrayOutput) ToManagedClusterAgentPoolProfileResponseArrayOutput() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseArrayOutput) ToManagedClusterAgentPoolProfileResponseArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseArrayOutput) Index(i pulumi.IntInput) ManagedClusterAgentPoolProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterAgentPoolProfileResponse {
		return vs[0].([]ManagedClusterAgentPoolProfileResponse)[vs[1].(int)]
	}).(ManagedClusterAgentPoolProfileResponseOutput)
}

type ManagedClusterIdentity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type ManagedClusterIdentityInput interface {
	pulumi.Input

	ToManagedClusterIdentityOutput() ManagedClusterIdentityOutput
	ToManagedClusterIdentityOutputWithContext(context.Context) ManagedClusterIdentityOutput
}

type ManagedClusterIdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (ManagedClusterIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentity)(nil)).Elem()
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityOutput() ManagedClusterIdentityOutput {
	return i.ToManagedClusterIdentityOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityOutputWithContext(ctx context.Context) ManagedClusterIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityOutput)
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return i.ToManagedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityOutput).ToManagedClusterIdentityPtrOutputWithContext(ctx)
}









type ManagedClusterIdentityPtrInput interface {
	pulumi.Input

	ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput
	ToManagedClusterIdentityPtrOutputWithContext(context.Context) ManagedClusterIdentityPtrOutput
}

type managedClusterIdentityPtrType ManagedClusterIdentityArgs

func ManagedClusterIdentityPtr(v *ManagedClusterIdentityArgs) ManagedClusterIdentityPtrInput {
	return (*managedClusterIdentityPtrType)(v)
}

func (*managedClusterIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentity)(nil)).Elem()
}

func (i *managedClusterIdentityPtrType) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return i.ToManagedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i *managedClusterIdentityPtrType) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityPtrOutput)
}

type ManagedClusterIdentityOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentity)(nil)).Elem()
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityOutput() ManagedClusterIdentityOutput {
	return o
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityOutputWithContext(ctx context.Context) ManagedClusterIdentityOutput {
	return o
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return o.ToManagedClusterIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterIdentity) *ManagedClusterIdentity {
		return &v
	}).(ManagedClusterIdentityPtrOutput)
}

func (o ManagedClusterIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedClusterIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type ManagedClusterIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentity)(nil)).Elem()
}

func (o ManagedClusterIdentityPtrOutput) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return o
}

func (o ManagedClusterIdentityPtrOutput) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return o
}

func (o ManagedClusterIdentityPtrOutput) Elem() ManagedClusterIdentityOutput {
	return o.ApplyT(func(v *ManagedClusterIdentity) ManagedClusterIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedClusterIdentity
		return ret
	}).(ManagedClusterIdentityOutput)
}

func (o ManagedClusterIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ManagedClusterIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ManagedClusterIdentityResponseInput interface {
	pulumi.Input

	ToManagedClusterIdentityResponseOutput() ManagedClusterIdentityResponseOutput
	ToManagedClusterIdentityResponseOutputWithContext(context.Context) ManagedClusterIdentityResponseOutput
}

type ManagedClusterIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ManagedClusterIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentityResponse)(nil)).Elem()
}

func (i ManagedClusterIdentityResponseArgs) ToManagedClusterIdentityResponseOutput() ManagedClusterIdentityResponseOutput {
	return i.ToManagedClusterIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityResponseArgs) ToManagedClusterIdentityResponseOutputWithContext(ctx context.Context) ManagedClusterIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityResponseOutput)
}

func (i ManagedClusterIdentityResponseArgs) ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput {
	return i.ToManagedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityResponseArgs) ToManagedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityResponseOutput).ToManagedClusterIdentityResponsePtrOutputWithContext(ctx)
}









type ManagedClusterIdentityResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput
	ToManagedClusterIdentityResponsePtrOutputWithContext(context.Context) ManagedClusterIdentityResponsePtrOutput
}

type managedClusterIdentityResponsePtrType ManagedClusterIdentityResponseArgs

func ManagedClusterIdentityResponsePtr(v *ManagedClusterIdentityResponseArgs) ManagedClusterIdentityResponsePtrInput {
	return (*managedClusterIdentityResponsePtrType)(v)
}

func (*managedClusterIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentityResponse)(nil)).Elem()
}

func (i *managedClusterIdentityResponsePtrType) ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput {
	return i.ToManagedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterIdentityResponsePtrType) ToManagedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedClusterIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityResponsePtrOutput)
}

type ManagedClusterIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentityResponse)(nil)).Elem()
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponseOutput() ManagedClusterIdentityResponseOutput {
	return o
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponseOutputWithContext(ctx context.Context) ManagedClusterIdentityResponseOutput {
	return o
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput {
	return o.ToManagedClusterIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedClusterIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterIdentityResponse) *ManagedClusterIdentityResponse {
		return &v
	}).(ManagedClusterIdentityResponsePtrOutput)
}

func (o ManagedClusterIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedClusterIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedClusterIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagedClusterIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentityResponse)(nil)).Elem()
}

func (o ManagedClusterIdentityResponsePtrOutput) ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput {
	return o
}

func (o ManagedClusterIdentityResponsePtrOutput) ToManagedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedClusterIdentityResponsePtrOutput {
	return o
}

func (o ManagedClusterIdentityResponsePtrOutput) Elem() ManagedClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) ManagedClusterIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterIdentityResponse
		return ret
	}).(ManagedClusterIdentityResponseOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterLoadBalancerProfile struct {
	AllocatedOutboundPorts *int                                                 `pulumi:"allocatedOutboundPorts"`
	EffectiveOutboundIPs   []ResourceReference                                  `pulumi:"effectiveOutboundIPs"`
	IdleTimeoutInMinutes   *int                                                 `pulumi:"idleTimeoutInMinutes"`
	ManagedOutboundIPs     *ManagedClusterLoadBalancerProfileManagedOutboundIPs `pulumi:"managedOutboundIPs"`
	OutboundIPPrefixes     *ManagedClusterLoadBalancerProfileOutboundIPPrefixes `pulumi:"outboundIPPrefixes"`
	OutboundIPs            *ManagedClusterLoadBalancerProfileOutboundIPs        `pulumi:"outboundIPs"`
}


func (val *ManagedClusterLoadBalancerProfile) Defaults() *ManagedClusterLoadBalancerProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocatedOutboundPorts) {
		allocatedOutboundPorts_ := 0
		tmp.AllocatedOutboundPorts = &allocatedOutboundPorts_
	}
	if isZero(tmp.IdleTimeoutInMinutes) {
		idleTimeoutInMinutes_ := 30
		tmp.IdleTimeoutInMinutes = &idleTimeoutInMinutes_
	}
	tmp.ManagedOutboundIPs = tmp.ManagedOutboundIPs.Defaults()

	return &tmp
}





type ManagedClusterLoadBalancerProfileInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutput() ManagedClusterLoadBalancerProfileOutput
	ToManagedClusterLoadBalancerProfileOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutput
}

type ManagedClusterLoadBalancerProfileArgs struct {
	AllocatedOutboundPorts pulumi.IntPtrInput                                          `pulumi:"allocatedOutboundPorts"`
	EffectiveOutboundIPs   ResourceReferenceArrayInput                                 `pulumi:"effectiveOutboundIPs"`
	IdleTimeoutInMinutes   pulumi.IntPtrInput                                          `pulumi:"idleTimeoutInMinutes"`
	ManagedOutboundIPs     ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrInput `pulumi:"managedOutboundIPs"`
	OutboundIPPrefixes     ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrInput `pulumi:"outboundIPPrefixes"`
	OutboundIPs            ManagedClusterLoadBalancerProfileOutboundIPsPtrInput        `pulumi:"outboundIPs"`
}

func (ManagedClusterLoadBalancerProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfile)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileArgs) ToManagedClusterLoadBalancerProfileOutput() ManagedClusterLoadBalancerProfileOutput {
	return i.ToManagedClusterLoadBalancerProfileOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileArgs) ToManagedClusterLoadBalancerProfileOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutput)
}

func (i ManagedClusterLoadBalancerProfileArgs) ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput {
	return i.ToManagedClusterLoadBalancerProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileArgs) ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutput).ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput
	ToManagedClusterLoadBalancerProfilePtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfilePtrOutput
}

type managedClusterLoadBalancerProfilePtrType ManagedClusterLoadBalancerProfileArgs

func ManagedClusterLoadBalancerProfilePtr(v *ManagedClusterLoadBalancerProfileArgs) ManagedClusterLoadBalancerProfilePtrInput {
	return (*managedClusterLoadBalancerProfilePtrType)(v)
}

func (*managedClusterLoadBalancerProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfile)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfilePtrType) ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput {
	return i.ToManagedClusterLoadBalancerProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfilePtrType) ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfilePtrOutput)
}

type ManagedClusterLoadBalancerProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfile)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutput) ToManagedClusterLoadBalancerProfileOutput() ManagedClusterLoadBalancerProfileOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutput) ToManagedClusterLoadBalancerProfileOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutput) ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput {
	return o.ToManagedClusterLoadBalancerProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileOutput) ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfile {
		return &v
	}).(ManagedClusterLoadBalancerProfilePtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *int { return v.AllocatedOutboundPorts }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) EffectiveOutboundIPs() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) []ResourceReference { return v.EffectiveOutboundIPs }).(ResourceReferenceArrayOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) ManagedOutboundIPs() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileManagedOutboundIPs {
		return v.ManagedOutboundIPs
	}).(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) OutboundIPPrefixes() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileOutboundIPPrefixes {
		return v.OutboundIPPrefixes
	}).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutput) OutboundIPs() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileOutboundIPs {
		return v.OutboundIPs
	}).(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfile)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) ToManagedClusterLoadBalancerProfilePtrOutput() ManagedClusterLoadBalancerProfilePtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) ToManagedClusterLoadBalancerProfilePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfilePtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) Elem() ManagedClusterLoadBalancerProfileOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) ManagedClusterLoadBalancerProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfile
		return ret
	}).(ManagedClusterLoadBalancerProfileOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *int {
		if v == nil {
			return nil
		}
		return v.AllocatedOutboundPorts
	}).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) EffectiveOutboundIPs() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) []ResourceReference {
		if v == nil {
			return nil
		}
		return v.EffectiveOutboundIPs
	}).(ResourceReferenceArrayOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) ManagedOutboundIPs() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileManagedOutboundIPs {
		if v == nil {
			return nil
		}
		return v.ManagedOutboundIPs
	}).(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) OutboundIPPrefixes() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileOutboundIPPrefixes {
		if v == nil {
			return nil
		}
		return v.OutboundIPPrefixes
	}).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfilePtrOutput) OutboundIPs() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfile) *ManagedClusterLoadBalancerProfileOutboundIPs {
		if v == nil {
			return nil
		}
		return v.OutboundIPs
	}).(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileManagedOutboundIPs struct {
	Count *int `pulumi:"count"`
}


func (val *ManagedClusterLoadBalancerProfileManagedOutboundIPs) Defaults() *ManagedClusterLoadBalancerProfileManagedOutboundIPs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	return &tmp
}





type ManagedClusterLoadBalancerProfileManagedOutboundIPsInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput
	ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput
}

type ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs struct {
	Count pulumi.IntPtrInput `pulumi:"count"`
}

func (ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileManagedOutboundIPs)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return i.ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput)
}

func (i ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput).ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput
	ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput
}

type managedClusterLoadBalancerProfileManagedOutboundIPsPtrType ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs

func ManagedClusterLoadBalancerProfileManagedOutboundIPsPtr(v *ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrInput {
	return (*managedClusterLoadBalancerProfileManagedOutboundIPsPtrType)(v)
}

func (*managedClusterLoadBalancerProfileManagedOutboundIPsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileManagedOutboundIPs)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileManagedOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileManagedOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileManagedOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileManagedOutboundIPs) *ManagedClusterLoadBalancerProfileManagedOutboundIPs {
		return &v
	}).(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileManagedOutboundIPs) *int { return v.Count }).(pulumi.IntPtrOutput)
}

type ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileManagedOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) Elem() ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileManagedOutboundIPs) ManagedClusterLoadBalancerProfileManagedOutboundIPs {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileManagedOutboundIPs
		return ret
	}).(ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput)
}

func (o ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileManagedOutboundIPs) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPPrefixes struct {
	PublicIPPrefixes []ResourceReference `pulumi:"publicIPPrefixes"`
}





type ManagedClusterLoadBalancerProfileOutboundIPPrefixesInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput
	ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput
}

type ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs struct {
	PublicIPPrefixes ResourceReferenceArrayInput `pulumi:"publicIPPrefixes"`
}

func (ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileOutboundIPPrefixes)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput)
}

func (i ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput).ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput
	ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput
}

type managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs

func ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtr(v *ManagedClusterLoadBalancerProfileOutboundIPPrefixesArgs) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrInput {
	return (*managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType)(v)
}

func (*managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileOutboundIPPrefixes)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileOutboundIPPrefixesPtrType) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileOutboundIPPrefixes)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileOutboundIPPrefixes) *ManagedClusterLoadBalancerProfileOutboundIPPrefixes {
		return &v
	}).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput) PublicIPPrefixes() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileOutboundIPPrefixes) []ResourceReference {
		return v.PublicIPPrefixes
	}).(ResourceReferenceArrayOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileOutboundIPPrefixes)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) ToManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) Elem() ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileOutboundIPPrefixes) ManagedClusterLoadBalancerProfileOutboundIPPrefixes {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileOutboundIPPrefixes
		return ret
	}).(ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput)
}

func (o ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput) PublicIPPrefixes() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileOutboundIPPrefixes) []ResourceReference {
		if v == nil {
			return nil
		}
		return v.PublicIPPrefixes
	}).(ResourceReferenceArrayOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPs struct {
	PublicIPs []ResourceReference `pulumi:"publicIPs"`
}





type ManagedClusterLoadBalancerProfileOutboundIPsInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutboundIPsOutput() ManagedClusterLoadBalancerProfileOutboundIPsOutput
	ToManagedClusterLoadBalancerProfileOutboundIPsOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutboundIPsOutput
}

type ManagedClusterLoadBalancerProfileOutboundIPsArgs struct {
	PublicIPs ResourceReferenceArrayInput `pulumi:"publicIPs"`
}

func (ManagedClusterLoadBalancerProfileOutboundIPsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileOutboundIPs)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileOutboundIPsArgs) ToManagedClusterLoadBalancerProfileOutboundIPsOutput() ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPsOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileOutboundIPsArgs) ToManagedClusterLoadBalancerProfileOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPsOutput)
}

func (i ManagedClusterLoadBalancerProfileOutboundIPsArgs) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileOutboundIPsArgs) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPsOutput).ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileOutboundIPsPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput
	ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput
}

type managedClusterLoadBalancerProfileOutboundIPsPtrType ManagedClusterLoadBalancerProfileOutboundIPsArgs

func ManagedClusterLoadBalancerProfileOutboundIPsPtr(v *ManagedClusterLoadBalancerProfileOutboundIPsArgs) ManagedClusterLoadBalancerProfileOutboundIPsPtrInput {
	return (*managedClusterLoadBalancerProfileOutboundIPsPtrType)(v)
}

func (*managedClusterLoadBalancerProfileOutboundIPsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileOutboundIPs)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPsOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) ToManagedClusterLoadBalancerProfileOutboundIPsOutput() ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) ToManagedClusterLoadBalancerProfileOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileOutboundIPs) *ManagedClusterLoadBalancerProfileOutboundIPs {
		return &v
	}).(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsOutput) PublicIPs() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileOutboundIPs) []ResourceReference { return v.PublicIPs }).(ResourceReferenceArrayOutput)
}

type ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) Elem() ManagedClusterLoadBalancerProfileOutboundIPsOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileOutboundIPs) ManagedClusterLoadBalancerProfileOutboundIPs {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileOutboundIPs
		return ret
	}).(ManagedClusterLoadBalancerProfileOutboundIPsOutput)
}

func (o ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput) PublicIPs() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileOutboundIPs) []ResourceReference {
		if v == nil {
			return nil
		}
		return v.PublicIPs
	}).(ResourceReferenceArrayOutput)
}

type ManagedClusterLoadBalancerProfileResponse struct {
	AllocatedOutboundPorts *int                                                         `pulumi:"allocatedOutboundPorts"`
	EffectiveOutboundIPs   []ResourceReferenceResponse                                  `pulumi:"effectiveOutboundIPs"`
	IdleTimeoutInMinutes   *int                                                         `pulumi:"idleTimeoutInMinutes"`
	ManagedOutboundIPs     *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs `pulumi:"managedOutboundIPs"`
	OutboundIPPrefixes     *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes `pulumi:"outboundIPPrefixes"`
	OutboundIPs            *ManagedClusterLoadBalancerProfileResponseOutboundIPs        `pulumi:"outboundIPs"`
}


func (val *ManagedClusterLoadBalancerProfileResponse) Defaults() *ManagedClusterLoadBalancerProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllocatedOutboundPorts) {
		allocatedOutboundPorts_ := 0
		tmp.AllocatedOutboundPorts = &allocatedOutboundPorts_
	}
	if isZero(tmp.IdleTimeoutInMinutes) {
		idleTimeoutInMinutes_ := 30
		tmp.IdleTimeoutInMinutes = &idleTimeoutInMinutes_
	}
	tmp.ManagedOutboundIPs = tmp.ManagedOutboundIPs.Defaults()

	return &tmp
}





type ManagedClusterLoadBalancerProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutput() ManagedClusterLoadBalancerProfileResponseOutput
	ToManagedClusterLoadBalancerProfileResponseOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutput
}

type ManagedClusterLoadBalancerProfileResponseArgs struct {
	AllocatedOutboundPorts pulumi.IntPtrInput                                                  `pulumi:"allocatedOutboundPorts"`
	EffectiveOutboundIPs   ResourceReferenceResponseArrayInput                                 `pulumi:"effectiveOutboundIPs"`
	IdleTimeoutInMinutes   pulumi.IntPtrInput                                                  `pulumi:"idleTimeoutInMinutes"`
	ManagedOutboundIPs     ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrInput `pulumi:"managedOutboundIPs"`
	OutboundIPPrefixes     ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrInput `pulumi:"outboundIPPrefixes"`
	OutboundIPs            ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrInput        `pulumi:"outboundIPs"`
}

func (ManagedClusterLoadBalancerProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponse)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileResponseArgs) ToManagedClusterLoadBalancerProfileResponseOutput() ManagedClusterLoadBalancerProfileResponseOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseArgs) ToManagedClusterLoadBalancerProfileResponseOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutput)
}

func (i ManagedClusterLoadBalancerProfileResponseArgs) ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseArgs) ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutput).ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput
	ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput
}

type managedClusterLoadBalancerProfileResponsePtrType ManagedClusterLoadBalancerProfileResponseArgs

func ManagedClusterLoadBalancerProfileResponsePtr(v *ManagedClusterLoadBalancerProfileResponseArgs) ManagedClusterLoadBalancerProfileResponsePtrInput {
	return (*managedClusterLoadBalancerProfileResponsePtrType)(v)
}

func (*managedClusterLoadBalancerProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponse)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileResponsePtrType) ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileResponsePtrType) ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponsePtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponse)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ToManagedClusterLoadBalancerProfileResponseOutput() ManagedClusterLoadBalancerProfileResponseOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ToManagedClusterLoadBalancerProfileResponseOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o.ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponse {
		return &v
	}).(ManagedClusterLoadBalancerProfileResponsePtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *int { return v.AllocatedOutboundPorts }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) EffectiveOutboundIPs() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) []ResourceReferenceResponse {
		return v.EffectiveOutboundIPs
	}).(ResourceReferenceResponseArrayOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) ManagedOutboundIPs() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
		return v.ManagedOutboundIPs
	}).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) OutboundIPPrefixes() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes {
		return v.OutboundIPPrefixes
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutput) OutboundIPs() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseOutboundIPs {
		return v.OutboundIPs
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponse)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) ToManagedClusterLoadBalancerProfileResponsePtrOutput() ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) ToManagedClusterLoadBalancerProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) Elem() ManagedClusterLoadBalancerProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) ManagedClusterLoadBalancerProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileResponse
		return ret
	}).(ManagedClusterLoadBalancerProfileResponseOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.AllocatedOutboundPorts
	}).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) EffectiveOutboundIPs() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) []ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.EffectiveOutboundIPs
	}).(ResourceReferenceResponseArrayOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) ManagedOutboundIPs() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
		if v == nil {
			return nil
		}
		return v.ManagedOutboundIPs
	}).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) OutboundIPPrefixes() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes {
		if v == nil {
			return nil
		}
		return v.OutboundIPPrefixes
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponsePtrOutput) OutboundIPs() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponse) *ManagedClusterLoadBalancerProfileResponseOutboundIPs {
		if v == nil {
			return nil
		}
		return v.OutboundIPs
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs struct {
	Count *int `pulumi:"count"`
}


func (val *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) Defaults() *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		count_ := 1
		tmp.Count = &count_
	}
	return &tmp
}





type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput
	ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput
}

type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs struct {
	Count pulumi.IntPtrInput `pulumi:"count"`
}

func (ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput)
}

func (i ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput).ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput
	ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput
}

type managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs

func ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtr(v *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrInput {
	return (*managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType)(v)
}

func (*managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
		return &v
	}).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) *int { return v.Count }).(pulumi.IntPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) Elem() ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs
		return ret
	}).(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes struct {
	PublicIPPrefixes []ResourceReferenceResponse `pulumi:"publicIPPrefixes"`
}





type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput
	ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs struct {
	PublicIPPrefixes ResourceReferenceResponseArrayInput `pulumi:"publicIPPrefixes"`
}

func (ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput)
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput).ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput
	ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput
}

type managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs

func ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtr(v *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrInput {
	return (*managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType)(v)
}

func (*managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrType) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes) *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes {
		return &v
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput) PublicIPPrefixes() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes) []ResourceReferenceResponse {
		return v.PublicIPPrefixes
	}).(ResourceReferenceResponseArrayOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) Elem() ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes) ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes
		return ret
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput) PublicIPPrefixes() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes) []ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.PublicIPPrefixes
	}).(ResourceReferenceResponseArrayOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPs struct {
	PublicIPs []ResourceReferenceResponse `pulumi:"publicIPs"`
}





type ManagedClusterLoadBalancerProfileResponseOutboundIPsInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput
	ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs struct {
	PublicIPs ResourceReferenceResponseArrayInput `pulumi:"publicIPs"`
}

func (ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseOutboundIPs)(nil)).Elem()
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput)
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(context.Background())
}

func (i ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput).ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx)
}









type ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrInput interface {
	pulumi.Input

	ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput
	ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput
}

type managedClusterLoadBalancerProfileResponseOutboundIPsPtrType ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs

func ManagedClusterLoadBalancerProfileResponseOutboundIPsPtr(v *ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrInput {
	return (*managedClusterLoadBalancerProfileResponseOutboundIPsPtrType)(v)
}

func (*managedClusterLoadBalancerProfileResponseOutboundIPsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseOutboundIPs)(nil)).Elem()
}

func (i *managedClusterLoadBalancerProfileResponseOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return i.ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(context.Background())
}

func (i *managedClusterLoadBalancerProfileResponseOutboundIPsPtrType) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterLoadBalancerProfileResponseOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o.ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(context.Background())
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterLoadBalancerProfileResponseOutboundIPs) *ManagedClusterLoadBalancerProfileResponseOutboundIPs {
		return &v
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput) PublicIPs() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v ManagedClusterLoadBalancerProfileResponseOutboundIPs) []ResourceReferenceResponse {
		return v.PublicIPs
	}).(ResourceReferenceResponseArrayOutput)
}

type ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterLoadBalancerProfileResponseOutboundIPs)(nil)).Elem()
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput() ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) ToManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput {
	return o
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) Elem() ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseOutboundIPs) ManagedClusterLoadBalancerProfileResponseOutboundIPs {
		if v != nil {
			return *v
		}
		var ret ManagedClusterLoadBalancerProfileResponseOutboundIPs
		return ret
	}).(ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput)
}

func (o ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput) PublicIPs() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ManagedClusterLoadBalancerProfileResponseOutboundIPs) []ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.PublicIPs
	}).(ResourceReferenceResponseArrayOutput)
}

type ManagedClusterPropertiesIdentityProfile struct {
	ClientId   *string `pulumi:"clientId"`
	ObjectId   *string `pulumi:"objectId"`
	ResourceId *string `pulumi:"resourceId"`
}





type ManagedClusterPropertiesIdentityProfileInput interface {
	pulumi.Input

	ToManagedClusterPropertiesIdentityProfileOutput() ManagedClusterPropertiesIdentityProfileOutput
	ToManagedClusterPropertiesIdentityProfileOutputWithContext(context.Context) ManagedClusterPropertiesIdentityProfileOutput
}

type ManagedClusterPropertiesIdentityProfileArgs struct {
	ClientId   pulumi.StringPtrInput `pulumi:"clientId"`
	ObjectId   pulumi.StringPtrInput `pulumi:"objectId"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ManagedClusterPropertiesIdentityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesIdentityProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesIdentityProfileArgs) ToManagedClusterPropertiesIdentityProfileOutput() ManagedClusterPropertiesIdentityProfileOutput {
	return i.ToManagedClusterPropertiesIdentityProfileOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesIdentityProfileArgs) ToManagedClusterPropertiesIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesIdentityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesIdentityProfileOutput)
}





type ManagedClusterPropertiesIdentityProfileMapInput interface {
	pulumi.Input

	ToManagedClusterPropertiesIdentityProfileMapOutput() ManagedClusterPropertiesIdentityProfileMapOutput
	ToManagedClusterPropertiesIdentityProfileMapOutputWithContext(context.Context) ManagedClusterPropertiesIdentityProfileMapOutput
}

type ManagedClusterPropertiesIdentityProfileMap map[string]ManagedClusterPropertiesIdentityProfileInput

func (ManagedClusterPropertiesIdentityProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterPropertiesIdentityProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesIdentityProfileMap) ToManagedClusterPropertiesIdentityProfileMapOutput() ManagedClusterPropertiesIdentityProfileMapOutput {
	return i.ToManagedClusterPropertiesIdentityProfileMapOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesIdentityProfileMap) ToManagedClusterPropertiesIdentityProfileMapOutputWithContext(ctx context.Context) ManagedClusterPropertiesIdentityProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesIdentityProfileMapOutput)
}

type ManagedClusterPropertiesIdentityProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesIdentityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ToManagedClusterPropertiesIdentityProfileOutput() ManagedClusterPropertiesIdentityProfileOutput {
	return o
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ToManagedClusterPropertiesIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesIdentityProfileOutput {
	return o
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesIdentityProfile) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesIdentityProfile) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesIdentityProfileOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesIdentityProfile) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ManagedClusterPropertiesIdentityProfileMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesIdentityProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterPropertiesIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesIdentityProfileMapOutput) ToManagedClusterPropertiesIdentityProfileMapOutput() ManagedClusterPropertiesIdentityProfileMapOutput {
	return o
}

func (o ManagedClusterPropertiesIdentityProfileMapOutput) ToManagedClusterPropertiesIdentityProfileMapOutputWithContext(ctx context.Context) ManagedClusterPropertiesIdentityProfileMapOutput {
	return o
}

func (o ManagedClusterPropertiesIdentityProfileMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterPropertiesIdentityProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterPropertiesIdentityProfile {
		return vs[0].(map[string]ManagedClusterPropertiesIdentityProfile)[vs[1].(string)]
	}).(ManagedClusterPropertiesIdentityProfileOutput)
}

type ManagedClusterPropertiesResponseIdentityProfile struct {
	ClientId   *string `pulumi:"clientId"`
	ObjectId   *string `pulumi:"objectId"`
	ResourceId *string `pulumi:"resourceId"`
}





type ManagedClusterPropertiesResponseIdentityProfileInput interface {
	pulumi.Input

	ToManagedClusterPropertiesResponseIdentityProfileOutput() ManagedClusterPropertiesResponseIdentityProfileOutput
	ToManagedClusterPropertiesResponseIdentityProfileOutputWithContext(context.Context) ManagedClusterPropertiesResponseIdentityProfileOutput
}

type ManagedClusterPropertiesResponseIdentityProfileArgs struct {
	ClientId   pulumi.StringPtrInput `pulumi:"clientId"`
	ObjectId   pulumi.StringPtrInput `pulumi:"objectId"`
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (ManagedClusterPropertiesResponseIdentityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesResponseIdentityProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesResponseIdentityProfileArgs) ToManagedClusterPropertiesResponseIdentityProfileOutput() ManagedClusterPropertiesResponseIdentityProfileOutput {
	return i.ToManagedClusterPropertiesResponseIdentityProfileOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesResponseIdentityProfileArgs) ToManagedClusterPropertiesResponseIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseIdentityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesResponseIdentityProfileOutput)
}





type ManagedClusterPropertiesResponseIdentityProfileMapInput interface {
	pulumi.Input

	ToManagedClusterPropertiesResponseIdentityProfileMapOutput() ManagedClusterPropertiesResponseIdentityProfileMapOutput
	ToManagedClusterPropertiesResponseIdentityProfileMapOutputWithContext(context.Context) ManagedClusterPropertiesResponseIdentityProfileMapOutput
}

type ManagedClusterPropertiesResponseIdentityProfileMap map[string]ManagedClusterPropertiesResponseIdentityProfileInput

func (ManagedClusterPropertiesResponseIdentityProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterPropertiesResponseIdentityProfile)(nil)).Elem()
}

func (i ManagedClusterPropertiesResponseIdentityProfileMap) ToManagedClusterPropertiesResponseIdentityProfileMapOutput() ManagedClusterPropertiesResponseIdentityProfileMapOutput {
	return i.ToManagedClusterPropertiesResponseIdentityProfileMapOutputWithContext(context.Background())
}

func (i ManagedClusterPropertiesResponseIdentityProfileMap) ToManagedClusterPropertiesResponseIdentityProfileMapOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseIdentityProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterPropertiesResponseIdentityProfileMapOutput)
}

type ManagedClusterPropertiesResponseIdentityProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesResponseIdentityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterPropertiesResponseIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ToManagedClusterPropertiesResponseIdentityProfileOutput() ManagedClusterPropertiesResponseIdentityProfileOutput {
	return o
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ToManagedClusterPropertiesResponseIdentityProfileOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseIdentityProfileOutput {
	return o
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseIdentityProfile) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseIdentityProfile) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterPropertiesResponseIdentityProfileOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterPropertiesResponseIdentityProfile) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type ManagedClusterPropertiesResponseIdentityProfileMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterPropertiesResponseIdentityProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterPropertiesResponseIdentityProfile)(nil)).Elem()
}

func (o ManagedClusterPropertiesResponseIdentityProfileMapOutput) ToManagedClusterPropertiesResponseIdentityProfileMapOutput() ManagedClusterPropertiesResponseIdentityProfileMapOutput {
	return o
}

func (o ManagedClusterPropertiesResponseIdentityProfileMapOutput) ToManagedClusterPropertiesResponseIdentityProfileMapOutputWithContext(ctx context.Context) ManagedClusterPropertiesResponseIdentityProfileMapOutput {
	return o
}

func (o ManagedClusterPropertiesResponseIdentityProfileMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterPropertiesResponseIdentityProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterPropertiesResponseIdentityProfile {
		return vs[0].(map[string]ManagedClusterPropertiesResponseIdentityProfile)[vs[1].(string)]
	}).(ManagedClusterPropertiesResponseIdentityProfileOutput)
}

type ManagedClusterServicePrincipalProfile struct {
	ClientId string  `pulumi:"clientId"`
	Secret   *string `pulumi:"secret"`
}





type ManagedClusterServicePrincipalProfileInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfileOutput() ManagedClusterServicePrincipalProfileOutput
	ToManagedClusterServicePrincipalProfileOutputWithContext(context.Context) ManagedClusterServicePrincipalProfileOutput
}

type ManagedClusterServicePrincipalProfileArgs struct {
	ClientId pulumi.StringInput    `pulumi:"clientId"`
	Secret   pulumi.StringPtrInput `pulumi:"secret"`
}

func (ManagedClusterServicePrincipalProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfileOutput() ManagedClusterServicePrincipalProfileOutput {
	return i.ToManagedClusterServicePrincipalProfileOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfileOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileOutput)
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return i.ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileOutput).ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx)
}









type ManagedClusterServicePrincipalProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput
	ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Context) ManagedClusterServicePrincipalProfilePtrOutput
}

type managedClusterServicePrincipalProfilePtrType ManagedClusterServicePrincipalProfileArgs

func ManagedClusterServicePrincipalProfilePtr(v *ManagedClusterServicePrincipalProfileArgs) ManagedClusterServicePrincipalProfilePtrInput {
	return (*managedClusterServicePrincipalProfilePtrType)(v)
}

func (*managedClusterServicePrincipalProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (i *managedClusterServicePrincipalProfilePtrType) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return i.ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterServicePrincipalProfilePtrType) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfilePtrOutput)
}

type ManagedClusterServicePrincipalProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfileOutput() ManagedClusterServicePrincipalProfileOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfileOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return o.ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterServicePrincipalProfile) *ManagedClusterServicePrincipalProfile {
		return &v
	}).(ManagedClusterServicePrincipalProfilePtrOutput)
}

func (o ManagedClusterServicePrincipalProfileOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfile) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ManagedClusterServicePrincipalProfileOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfile) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) Elem() ManagedClusterServicePrincipalProfileOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfile) ManagedClusterServicePrincipalProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterServicePrincipalProfile
		return ret
	}).(ManagedClusterServicePrincipalProfileOutput)
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfileResponse struct {
	ClientId string  `pulumi:"clientId"`
	Secret   *string `pulumi:"secret"`
}





type ManagedClusterServicePrincipalProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfileResponseOutput() ManagedClusterServicePrincipalProfileResponseOutput
	ToManagedClusterServicePrincipalProfileResponseOutputWithContext(context.Context) ManagedClusterServicePrincipalProfileResponseOutput
}

type ManagedClusterServicePrincipalProfileResponseArgs struct {
	ClientId pulumi.StringInput    `pulumi:"clientId"`
	Secret   pulumi.StringPtrInput `pulumi:"secret"`
}

func (ManagedClusterServicePrincipalProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (i ManagedClusterServicePrincipalProfileResponseArgs) ToManagedClusterServicePrincipalProfileResponseOutput() ManagedClusterServicePrincipalProfileResponseOutput {
	return i.ToManagedClusterServicePrincipalProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileResponseArgs) ToManagedClusterServicePrincipalProfileResponseOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileResponseOutput)
}

func (i ManagedClusterServicePrincipalProfileResponseArgs) ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return i.ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileResponseArgs) ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileResponseOutput).ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterServicePrincipalProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput
	ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput
}

type managedClusterServicePrincipalProfileResponsePtrType ManagedClusterServicePrincipalProfileResponseArgs

func ManagedClusterServicePrincipalProfileResponsePtr(v *ManagedClusterServicePrincipalProfileResponseArgs) ManagedClusterServicePrincipalProfileResponsePtrInput {
	return (*managedClusterServicePrincipalProfileResponsePtrType)(v)
}

func (*managedClusterServicePrincipalProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (i *managedClusterServicePrincipalProfileResponsePtrType) ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return i.ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterServicePrincipalProfileResponsePtrType) ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

type ManagedClusterServicePrincipalProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponseOutput() ManagedClusterServicePrincipalProfileResponseOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponseOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponseOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterServicePrincipalProfileResponse) *ManagedClusterServicePrincipalProfileResponse {
		return &v
	}).(ManagedClusterServicePrincipalProfileResponsePtrOutput)
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfileResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfileResponse) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) Elem() ManagedClusterServicePrincipalProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfileResponse) ManagedClusterServicePrincipalProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterServicePrincipalProfileResponse
		return ret
	}).(ManagedClusterServicePrincipalProfileResponseOutput)
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterWindowsProfile struct {
	AdminPassword *string `pulumi:"adminPassword"`
	AdminUsername string  `pulumi:"adminUsername"`
}





type ManagedClusterWindowsProfileInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfileOutput() ManagedClusterWindowsProfileOutput
	ToManagedClusterWindowsProfileOutputWithContext(context.Context) ManagedClusterWindowsProfileOutput
}

type ManagedClusterWindowsProfileArgs struct {
	AdminPassword pulumi.StringPtrInput `pulumi:"adminPassword"`
	AdminUsername pulumi.StringInput    `pulumi:"adminUsername"`
}

func (ManagedClusterWindowsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfile)(nil)).Elem()
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfileOutput() ManagedClusterWindowsProfileOutput {
	return i.ToManagedClusterWindowsProfileOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfileOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileOutput)
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return i.ToManagedClusterWindowsProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileOutput).ToManagedClusterWindowsProfilePtrOutputWithContext(ctx)
}









type ManagedClusterWindowsProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput
	ToManagedClusterWindowsProfilePtrOutputWithContext(context.Context) ManagedClusterWindowsProfilePtrOutput
}

type managedClusterWindowsProfilePtrType ManagedClusterWindowsProfileArgs

func ManagedClusterWindowsProfilePtr(v *ManagedClusterWindowsProfileArgs) ManagedClusterWindowsProfilePtrInput {
	return (*managedClusterWindowsProfilePtrType)(v)
}

func (*managedClusterWindowsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfile)(nil)).Elem()
}

func (i *managedClusterWindowsProfilePtrType) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return i.ToManagedClusterWindowsProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterWindowsProfilePtrType) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfilePtrOutput)
}

type ManagedClusterWindowsProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfile)(nil)).Elem()
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfileOutput() ManagedClusterWindowsProfileOutput {
	return o
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfileOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileOutput {
	return o
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return o.ToManagedClusterWindowsProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterWindowsProfile) *ManagedClusterWindowsProfile {
		return &v
	}).(ManagedClusterWindowsProfilePtrOutput)
}

func (o ManagedClusterWindowsProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfile) string { return v.AdminUsername }).(pulumi.StringOutput)
}

type ManagedClusterWindowsProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfile)(nil)).Elem()
}

func (o ManagedClusterWindowsProfilePtrOutput) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfilePtrOutput) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfilePtrOutput) Elem() ManagedClusterWindowsProfileOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) ManagedClusterWindowsProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterWindowsProfile
		return ret
	}).(ManagedClusterWindowsProfileOutput)
}

func (o ManagedClusterWindowsProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterWindowsProfileResponse struct {
	AdminPassword *string `pulumi:"adminPassword"`
	AdminUsername string  `pulumi:"adminUsername"`
}





type ManagedClusterWindowsProfileResponseInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfileResponseOutput() ManagedClusterWindowsProfileResponseOutput
	ToManagedClusterWindowsProfileResponseOutputWithContext(context.Context) ManagedClusterWindowsProfileResponseOutput
}

type ManagedClusterWindowsProfileResponseArgs struct {
	AdminPassword pulumi.StringPtrInput `pulumi:"adminPassword"`
	AdminUsername pulumi.StringInput    `pulumi:"adminUsername"`
}

func (ManagedClusterWindowsProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (i ManagedClusterWindowsProfileResponseArgs) ToManagedClusterWindowsProfileResponseOutput() ManagedClusterWindowsProfileResponseOutput {
	return i.ToManagedClusterWindowsProfileResponseOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileResponseArgs) ToManagedClusterWindowsProfileResponseOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileResponseOutput)
}

func (i ManagedClusterWindowsProfileResponseArgs) ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput {
	return i.ToManagedClusterWindowsProfileResponsePtrOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileResponseArgs) ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileResponseOutput).ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx)
}









type ManagedClusterWindowsProfileResponsePtrInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput
	ToManagedClusterWindowsProfileResponsePtrOutputWithContext(context.Context) ManagedClusterWindowsProfileResponsePtrOutput
}

type managedClusterWindowsProfileResponsePtrType ManagedClusterWindowsProfileResponseArgs

func ManagedClusterWindowsProfileResponsePtr(v *ManagedClusterWindowsProfileResponseArgs) ManagedClusterWindowsProfileResponsePtrInput {
	return (*managedClusterWindowsProfileResponsePtrType)(v)
}

func (*managedClusterWindowsProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (i *managedClusterWindowsProfileResponsePtrType) ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput {
	return i.ToManagedClusterWindowsProfileResponsePtrOutputWithContext(context.Background())
}

func (i *managedClusterWindowsProfileResponsePtrType) ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileResponsePtrOutput)
}

type ManagedClusterWindowsProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponseOutput() ManagedClusterWindowsProfileResponseOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponseOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponseOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput {
	return o.ToManagedClusterWindowsProfileResponsePtrOutputWithContext(context.Background())
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterWindowsProfileResponse) *ManagedClusterWindowsProfileResponse {
		return &v
	}).(ManagedClusterWindowsProfileResponsePtrOutput)
}

func (o ManagedClusterWindowsProfileResponseOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfileResponse) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileResponseOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfileResponse) string { return v.AdminUsername }).(pulumi.StringOutput)
}

type ManagedClusterWindowsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) Elem() ManagedClusterWindowsProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) ManagedClusterWindowsProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterWindowsProfileResponse
		return ret
	}).(ManagedClusterWindowsProfileResponseOutput)
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

type ResourceReference struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceInput interface {
	pulumi.Input

	ToResourceReferenceOutput() ResourceReferenceOutput
	ToResourceReferenceOutputWithContext(context.Context) ResourceReferenceOutput
}

type ResourceReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArgs) ToResourceReferenceOutput() ResourceReferenceOutput {
	return i.ToResourceReferenceOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput)
}





type ResourceReferenceArrayInput interface {
	pulumi.Input

	ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput
	ToResourceReferenceArrayOutputWithContext(context.Context) ResourceReferenceArrayOutput
}

type ResourceReferenceArray []ResourceReferenceInput

func (ResourceReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArray) ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput {
	return i.ToResourceReferenceArrayOutputWithContext(context.Background())
}

func (i ResourceReferenceArray) ToResourceReferenceArrayOutputWithContext(ctx context.Context) ResourceReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceArrayOutput)
}

type ResourceReferenceOutput struct{ *pulumi.OutputState }

func (ResourceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceOutput) ToResourceReferenceOutput() ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferenceArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceArrayOutput) ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput {
	return o
}

func (o ResourceReferenceArrayOutput) ToResourceReferenceArrayOutputWithContext(ctx context.Context) ResourceReferenceArrayOutput {
	return o
}

func (o ResourceReferenceArrayOutput) Index(i pulumi.IntInput) ResourceReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReference {
		return vs[0].([]ResourceReference)[vs[1].(int)]
	}).(ResourceReferenceOutput)
}

type ResourceReferenceResponse struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceResponseInput interface {
	pulumi.Input

	ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput
	ToResourceReferenceResponseOutputWithContext(context.Context) ResourceReferenceResponseOutput
}

type ResourceReferenceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return i.ToResourceReferenceResponseOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArgs) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseOutput)
}





type ResourceReferenceResponseArrayInput interface {
	pulumi.Input

	ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput
	ToResourceReferenceResponseArrayOutputWithContext(context.Context) ResourceReferenceResponseArrayOutput
}

type ResourceReferenceResponseArray []ResourceReferenceResponseInput

func (ResourceReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (i ResourceReferenceResponseArray) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return i.ToResourceReferenceResponseArrayOutputWithContext(context.Background())
}

func (i ResourceReferenceResponseArray) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceResponseArrayOutput)
}

type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) Index(i pulumi.IntInput) ResourceReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReferenceResponse {
		return vs[0].([]ResourceReferenceResponse)[vs[1].(int)]
	}).(ResourceReferenceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAPIServerAccessProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAPIServerAccessProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAPIServerAccessProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAPIServerAccessProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileResponseMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileResponseIdentityOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileManagedOutboundIPsOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileManagedOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutboundIPPrefixesOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutboundIPPrefixesPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutboundIPsOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutboundIPsOutput{})
	pulumi.RegisterOutputType(ManagedClusterLoadBalancerProfileResponseOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesIdentityProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesIdentityProfileMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesResponseIdentityProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterPropertiesResponseIdentityProfileMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceOutput{})
	pulumi.RegisterOutputType(ResourceReferenceArrayOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseArrayOutput{})
}
