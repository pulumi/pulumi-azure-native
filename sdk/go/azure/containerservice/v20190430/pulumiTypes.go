


package v20190430

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkProfile struct {
	PeerVnetId *string `pulumi:"peerVnetId"`
	VnetCidr   *string `pulumi:"vnetCidr"`
	VnetId     *string `pulumi:"vnetId"`
}


func (val *NetworkProfile) Defaults() *NetworkProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.VnetCidr) {
		vnetCidr_ := "10.0.0.0/8"
		tmp.VnetCidr = &vnetCidr_
	}
	return &tmp
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	PeerVnetId pulumi.StringPtrInput `pulumi:"peerVnetId"`
	VnetCidr   pulumi.StringPtrInput `pulumi:"vnetCidr"`
	VnetId     pulumi.StringPtrInput `pulumi:"vnetId"`
}


func (val *NetworkProfileArgs) Defaults() *NetworkProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.VnetCidr) {
		tmp.VnetCidr = pulumi.StringPtr("10.0.0.0/8")
	}
	return &tmp
}
func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) PeerVnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.PeerVnetId }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) VnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.VnetCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) VnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.VnetId }).(pulumi.StringPtrOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) PeerVnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.PeerVnetId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) VnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.VnetCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) VnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.VnetId
	}).(pulumi.StringPtrOutput)
}

type NetworkProfileResponse struct {
	PeerVnetId *string `pulumi:"peerVnetId"`
	VnetCidr   *string `pulumi:"vnetCidr"`
	VnetId     *string `pulumi:"vnetId"`
}


func (val *NetworkProfileResponse) Defaults() *NetworkProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.VnetCidr) {
		vnetCidr_ := "10.0.0.0/8"
		tmp.VnetCidr = &vnetCidr_
	}
	return &tmp
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) PeerVnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.PeerVnetId }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) VnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.VnetCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) VnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.VnetId }).(pulumi.StringPtrOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) PeerVnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.PeerVnetId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) VnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VnetCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) VnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VnetId
	}).(pulumi.StringPtrOutput)
}

type OpenShiftManagedClusterAADIdentityProvider struct {
	ClientId             *string `pulumi:"clientId"`
	CustomerAdminGroupId *string `pulumi:"customerAdminGroupId"`
	Kind                 string  `pulumi:"kind"`
	Secret               *string `pulumi:"secret"`
	TenantId             *string `pulumi:"tenantId"`
}





type OpenShiftManagedClusterAADIdentityProviderInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterAADIdentityProviderOutput() OpenShiftManagedClusterAADIdentityProviderOutput
	ToOpenShiftManagedClusterAADIdentityProviderOutputWithContext(context.Context) OpenShiftManagedClusterAADIdentityProviderOutput
}

type OpenShiftManagedClusterAADIdentityProviderArgs struct {
	ClientId             pulumi.StringPtrInput `pulumi:"clientId"`
	CustomerAdminGroupId pulumi.StringPtrInput `pulumi:"customerAdminGroupId"`
	Kind                 pulumi.StringInput    `pulumi:"kind"`
	Secret               pulumi.StringPtrInput `pulumi:"secret"`
	TenantId             pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (OpenShiftManagedClusterAADIdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterAADIdentityProvider)(nil)).Elem()
}

func (i OpenShiftManagedClusterAADIdentityProviderArgs) ToOpenShiftManagedClusterAADIdentityProviderOutput() OpenShiftManagedClusterAADIdentityProviderOutput {
	return i.ToOpenShiftManagedClusterAADIdentityProviderOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterAADIdentityProviderArgs) ToOpenShiftManagedClusterAADIdentityProviderOutputWithContext(ctx context.Context) OpenShiftManagedClusterAADIdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterAADIdentityProviderOutput)
}

func (i OpenShiftManagedClusterAADIdentityProviderArgs) ToOpenShiftManagedClusterAADIdentityProviderPtrOutput() OpenShiftManagedClusterAADIdentityProviderPtrOutput {
	return i.ToOpenShiftManagedClusterAADIdentityProviderPtrOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterAADIdentityProviderArgs) ToOpenShiftManagedClusterAADIdentityProviderPtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAADIdentityProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterAADIdentityProviderOutput).ToOpenShiftManagedClusterAADIdentityProviderPtrOutputWithContext(ctx)
}









type OpenShiftManagedClusterAADIdentityProviderPtrInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterAADIdentityProviderPtrOutput() OpenShiftManagedClusterAADIdentityProviderPtrOutput
	ToOpenShiftManagedClusterAADIdentityProviderPtrOutputWithContext(context.Context) OpenShiftManagedClusterAADIdentityProviderPtrOutput
}

type openShiftManagedClusterAADIdentityProviderPtrType OpenShiftManagedClusterAADIdentityProviderArgs

func OpenShiftManagedClusterAADIdentityProviderPtr(v *OpenShiftManagedClusterAADIdentityProviderArgs) OpenShiftManagedClusterAADIdentityProviderPtrInput {
	return (*openShiftManagedClusterAADIdentityProviderPtrType)(v)
}

func (*openShiftManagedClusterAADIdentityProviderPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedClusterAADIdentityProvider)(nil)).Elem()
}

func (i *openShiftManagedClusterAADIdentityProviderPtrType) ToOpenShiftManagedClusterAADIdentityProviderPtrOutput() OpenShiftManagedClusterAADIdentityProviderPtrOutput {
	return i.ToOpenShiftManagedClusterAADIdentityProviderPtrOutputWithContext(context.Background())
}

func (i *openShiftManagedClusterAADIdentityProviderPtrType) ToOpenShiftManagedClusterAADIdentityProviderPtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAADIdentityProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterAADIdentityProviderPtrOutput)
}

type OpenShiftManagedClusterAADIdentityProviderOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAADIdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterAADIdentityProvider)(nil)).Elem()
}

func (o OpenShiftManagedClusterAADIdentityProviderOutput) ToOpenShiftManagedClusterAADIdentityProviderOutput() OpenShiftManagedClusterAADIdentityProviderOutput {
	return o
}

func (o OpenShiftManagedClusterAADIdentityProviderOutput) ToOpenShiftManagedClusterAADIdentityProviderOutputWithContext(ctx context.Context) OpenShiftManagedClusterAADIdentityProviderOutput {
	return o
}

func (o OpenShiftManagedClusterAADIdentityProviderOutput) ToOpenShiftManagedClusterAADIdentityProviderPtrOutput() OpenShiftManagedClusterAADIdentityProviderPtrOutput {
	return o.ToOpenShiftManagedClusterAADIdentityProviderPtrOutputWithContext(context.Background())
}

func (o OpenShiftManagedClusterAADIdentityProviderOutput) ToOpenShiftManagedClusterAADIdentityProviderPtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAADIdentityProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenShiftManagedClusterAADIdentityProvider) *OpenShiftManagedClusterAADIdentityProvider {
		return &v
	}).(OpenShiftManagedClusterAADIdentityProviderPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProvider) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderOutput) CustomerAdminGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProvider) *string { return v.CustomerAdminGroupId }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProvider) string { return v.Kind }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProvider) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProvider) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type OpenShiftManagedClusterAADIdentityProviderPtrOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAADIdentityProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedClusterAADIdentityProvider)(nil)).Elem()
}

func (o OpenShiftManagedClusterAADIdentityProviderPtrOutput) ToOpenShiftManagedClusterAADIdentityProviderPtrOutput() OpenShiftManagedClusterAADIdentityProviderPtrOutput {
	return o
}

func (o OpenShiftManagedClusterAADIdentityProviderPtrOutput) ToOpenShiftManagedClusterAADIdentityProviderPtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAADIdentityProviderPtrOutput {
	return o
}

func (o OpenShiftManagedClusterAADIdentityProviderPtrOutput) Elem() OpenShiftManagedClusterAADIdentityProviderOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProvider) OpenShiftManagedClusterAADIdentityProvider {
		if v != nil {
			return *v
		}
		var ret OpenShiftManagedClusterAADIdentityProvider
		return ret
	}).(OpenShiftManagedClusterAADIdentityProviderOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProvider) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderPtrOutput) CustomerAdminGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProvider) *string {
		if v == nil {
			return nil
		}
		return v.CustomerAdminGroupId
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProvider) *string {
		if v == nil {
			return nil
		}
		return &v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProvider) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProvider) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type OpenShiftManagedClusterAADIdentityProviderResponse struct {
	ClientId             *string `pulumi:"clientId"`
	CustomerAdminGroupId *string `pulumi:"customerAdminGroupId"`
	Kind                 string  `pulumi:"kind"`
	Secret               *string `pulumi:"secret"`
	TenantId             *string `pulumi:"tenantId"`
}

type OpenShiftManagedClusterAADIdentityProviderResponseOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAADIdentityProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterAADIdentityProviderResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterAADIdentityProviderResponseOutput) ToOpenShiftManagedClusterAADIdentityProviderResponseOutput() OpenShiftManagedClusterAADIdentityProviderResponseOutput {
	return o
}

func (o OpenShiftManagedClusterAADIdentityProviderResponseOutput) ToOpenShiftManagedClusterAADIdentityProviderResponseOutputWithContext(ctx context.Context) OpenShiftManagedClusterAADIdentityProviderResponseOutput {
	return o
}

func (o OpenShiftManagedClusterAADIdentityProviderResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProviderResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderResponseOutput) CustomerAdminGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProviderResponse) *string { return v.CustomerAdminGroupId }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProviderResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderResponseOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProviderResponse) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAADIdentityProviderResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedClusterAADIdentityProviderResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput) ToOpenShiftManagedClusterAADIdentityProviderResponsePtrOutput() OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput {
	return o
}

func (o OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput) ToOpenShiftManagedClusterAADIdentityProviderResponsePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput {
	return o
}

func (o OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput) Elem() OpenShiftManagedClusterAADIdentityProviderResponseOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProviderResponse) OpenShiftManagedClusterAADIdentityProviderResponse {
		if v != nil {
			return *v
		}
		var ret OpenShiftManagedClusterAADIdentityProviderResponse
		return ret
	}).(OpenShiftManagedClusterAADIdentityProviderResponseOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProviderResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput) CustomerAdminGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProviderResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomerAdminGroupId
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProviderResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProviderResponse) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAADIdentityProviderResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type OpenShiftManagedClusterAgentPoolProfile struct {
	Count      int     `pulumi:"count"`
	Name       string  `pulumi:"name"`
	OsType     *string `pulumi:"osType"`
	Role       *string `pulumi:"role"`
	SubnetCidr *string `pulumi:"subnetCidr"`
	VmSize     string  `pulumi:"vmSize"`
}


func (val *OpenShiftManagedClusterAgentPoolProfile) Defaults() *OpenShiftManagedClusterAgentPoolProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SubnetCidr) {
		subnetCidr_ := "10.0.0.0/24"
		tmp.SubnetCidr = &subnetCidr_
	}
	return &tmp
}





type OpenShiftManagedClusterAgentPoolProfileInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterAgentPoolProfileOutput() OpenShiftManagedClusterAgentPoolProfileOutput
	ToOpenShiftManagedClusterAgentPoolProfileOutputWithContext(context.Context) OpenShiftManagedClusterAgentPoolProfileOutput
}

type OpenShiftManagedClusterAgentPoolProfileArgs struct {
	Count      pulumi.IntInput       `pulumi:"count"`
	Name       pulumi.StringInput    `pulumi:"name"`
	OsType     pulumi.StringPtrInput `pulumi:"osType"`
	Role       pulumi.StringPtrInput `pulumi:"role"`
	SubnetCidr pulumi.StringPtrInput `pulumi:"subnetCidr"`
	VmSize     pulumi.StringInput    `pulumi:"vmSize"`
}


func (val *OpenShiftManagedClusterAgentPoolProfileArgs) Defaults() *OpenShiftManagedClusterAgentPoolProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SubnetCidr) {
		tmp.SubnetCidr = pulumi.StringPtr("10.0.0.0/24")
	}
	return &tmp
}
func (OpenShiftManagedClusterAgentPoolProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (i OpenShiftManagedClusterAgentPoolProfileArgs) ToOpenShiftManagedClusterAgentPoolProfileOutput() OpenShiftManagedClusterAgentPoolProfileOutput {
	return i.ToOpenShiftManagedClusterAgentPoolProfileOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterAgentPoolProfileArgs) ToOpenShiftManagedClusterAgentPoolProfileOutputWithContext(ctx context.Context) OpenShiftManagedClusterAgentPoolProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterAgentPoolProfileOutput)
}





type OpenShiftManagedClusterAgentPoolProfileArrayInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterAgentPoolProfileArrayOutput() OpenShiftManagedClusterAgentPoolProfileArrayOutput
	ToOpenShiftManagedClusterAgentPoolProfileArrayOutputWithContext(context.Context) OpenShiftManagedClusterAgentPoolProfileArrayOutput
}

type OpenShiftManagedClusterAgentPoolProfileArray []OpenShiftManagedClusterAgentPoolProfileInput

func (OpenShiftManagedClusterAgentPoolProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenShiftManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (i OpenShiftManagedClusterAgentPoolProfileArray) ToOpenShiftManagedClusterAgentPoolProfileArrayOutput() OpenShiftManagedClusterAgentPoolProfileArrayOutput {
	return i.ToOpenShiftManagedClusterAgentPoolProfileArrayOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterAgentPoolProfileArray) ToOpenShiftManagedClusterAgentPoolProfileArrayOutputWithContext(ctx context.Context) OpenShiftManagedClusterAgentPoolProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterAgentPoolProfileArrayOutput)
}

type OpenShiftManagedClusterAgentPoolProfileOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAgentPoolProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (o OpenShiftManagedClusterAgentPoolProfileOutput) ToOpenShiftManagedClusterAgentPoolProfileOutput() OpenShiftManagedClusterAgentPoolProfileOutput {
	return o
}

func (o OpenShiftManagedClusterAgentPoolProfileOutput) ToOpenShiftManagedClusterAgentPoolProfileOutputWithContext(ctx context.Context) OpenShiftManagedClusterAgentPoolProfileOutput {
	return o
}

func (o OpenShiftManagedClusterAgentPoolProfileOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfile) int { return v.Count }).(pulumi.IntOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfile) string { return v.Name }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfile) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileOutput) SubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfile) *string { return v.SubnetCidr }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfile) string { return v.VmSize }).(pulumi.StringOutput)
}

type OpenShiftManagedClusterAgentPoolProfileArrayOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAgentPoolProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenShiftManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (o OpenShiftManagedClusterAgentPoolProfileArrayOutput) ToOpenShiftManagedClusterAgentPoolProfileArrayOutput() OpenShiftManagedClusterAgentPoolProfileArrayOutput {
	return o
}

func (o OpenShiftManagedClusterAgentPoolProfileArrayOutput) ToOpenShiftManagedClusterAgentPoolProfileArrayOutputWithContext(ctx context.Context) OpenShiftManagedClusterAgentPoolProfileArrayOutput {
	return o
}

func (o OpenShiftManagedClusterAgentPoolProfileArrayOutput) Index(i pulumi.IntInput) OpenShiftManagedClusterAgentPoolProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpenShiftManagedClusterAgentPoolProfile {
		return vs[0].([]OpenShiftManagedClusterAgentPoolProfile)[vs[1].(int)]
	}).(OpenShiftManagedClusterAgentPoolProfileOutput)
}

type OpenShiftManagedClusterAgentPoolProfileResponse struct {
	Count      int     `pulumi:"count"`
	Name       string  `pulumi:"name"`
	OsType     *string `pulumi:"osType"`
	Role       *string `pulumi:"role"`
	SubnetCidr *string `pulumi:"subnetCidr"`
	VmSize     string  `pulumi:"vmSize"`
}


func (val *OpenShiftManagedClusterAgentPoolProfileResponse) Defaults() *OpenShiftManagedClusterAgentPoolProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SubnetCidr) {
		subnetCidr_ := "10.0.0.0/24"
		tmp.SubnetCidr = &subnetCidr_
	}
	return &tmp
}

type OpenShiftManagedClusterAgentPoolProfileResponseOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAgentPoolProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseOutput) ToOpenShiftManagedClusterAgentPoolProfileResponseOutput() OpenShiftManagedClusterAgentPoolProfileResponseOutput {
	return o
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseOutput) ToOpenShiftManagedClusterAgentPoolProfileResponseOutputWithContext(ctx context.Context) OpenShiftManagedClusterAgentPoolProfileResponseOutput {
	return o
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfileResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfileResponse) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseOutput) SubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfileResponse) *string { return v.SubnetCidr }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAgentPoolProfileResponse) string { return v.VmSize }).(pulumi.StringOutput)
}

type OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenShiftManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput) ToOpenShiftManagedClusterAgentPoolProfileResponseArrayOutput() OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput {
	return o
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput) ToOpenShiftManagedClusterAgentPoolProfileResponseArrayOutputWithContext(ctx context.Context) OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput {
	return o
}

func (o OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput) Index(i pulumi.IntInput) OpenShiftManagedClusterAgentPoolProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpenShiftManagedClusterAgentPoolProfileResponse {
		return vs[0].([]OpenShiftManagedClusterAgentPoolProfileResponse)[vs[1].(int)]
	}).(OpenShiftManagedClusterAgentPoolProfileResponseOutput)
}

type OpenShiftManagedClusterAuthProfile struct {
	IdentityProviders []OpenShiftManagedClusterIdentityProvider `pulumi:"identityProviders"`
}





type OpenShiftManagedClusterAuthProfileInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterAuthProfileOutput() OpenShiftManagedClusterAuthProfileOutput
	ToOpenShiftManagedClusterAuthProfileOutputWithContext(context.Context) OpenShiftManagedClusterAuthProfileOutput
}

type OpenShiftManagedClusterAuthProfileArgs struct {
	IdentityProviders OpenShiftManagedClusterIdentityProviderArrayInput `pulumi:"identityProviders"`
}

func (OpenShiftManagedClusterAuthProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterAuthProfile)(nil)).Elem()
}

func (i OpenShiftManagedClusterAuthProfileArgs) ToOpenShiftManagedClusterAuthProfileOutput() OpenShiftManagedClusterAuthProfileOutput {
	return i.ToOpenShiftManagedClusterAuthProfileOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterAuthProfileArgs) ToOpenShiftManagedClusterAuthProfileOutputWithContext(ctx context.Context) OpenShiftManagedClusterAuthProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterAuthProfileOutput)
}

func (i OpenShiftManagedClusterAuthProfileArgs) ToOpenShiftManagedClusterAuthProfilePtrOutput() OpenShiftManagedClusterAuthProfilePtrOutput {
	return i.ToOpenShiftManagedClusterAuthProfilePtrOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterAuthProfileArgs) ToOpenShiftManagedClusterAuthProfilePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAuthProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterAuthProfileOutput).ToOpenShiftManagedClusterAuthProfilePtrOutputWithContext(ctx)
}









type OpenShiftManagedClusterAuthProfilePtrInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterAuthProfilePtrOutput() OpenShiftManagedClusterAuthProfilePtrOutput
	ToOpenShiftManagedClusterAuthProfilePtrOutputWithContext(context.Context) OpenShiftManagedClusterAuthProfilePtrOutput
}

type openShiftManagedClusterAuthProfilePtrType OpenShiftManagedClusterAuthProfileArgs

func OpenShiftManagedClusterAuthProfilePtr(v *OpenShiftManagedClusterAuthProfileArgs) OpenShiftManagedClusterAuthProfilePtrInput {
	return (*openShiftManagedClusterAuthProfilePtrType)(v)
}

func (*openShiftManagedClusterAuthProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedClusterAuthProfile)(nil)).Elem()
}

func (i *openShiftManagedClusterAuthProfilePtrType) ToOpenShiftManagedClusterAuthProfilePtrOutput() OpenShiftManagedClusterAuthProfilePtrOutput {
	return i.ToOpenShiftManagedClusterAuthProfilePtrOutputWithContext(context.Background())
}

func (i *openShiftManagedClusterAuthProfilePtrType) ToOpenShiftManagedClusterAuthProfilePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAuthProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterAuthProfilePtrOutput)
}

type OpenShiftManagedClusterAuthProfileOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAuthProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterAuthProfile)(nil)).Elem()
}

func (o OpenShiftManagedClusterAuthProfileOutput) ToOpenShiftManagedClusterAuthProfileOutput() OpenShiftManagedClusterAuthProfileOutput {
	return o
}

func (o OpenShiftManagedClusterAuthProfileOutput) ToOpenShiftManagedClusterAuthProfileOutputWithContext(ctx context.Context) OpenShiftManagedClusterAuthProfileOutput {
	return o
}

func (o OpenShiftManagedClusterAuthProfileOutput) ToOpenShiftManagedClusterAuthProfilePtrOutput() OpenShiftManagedClusterAuthProfilePtrOutput {
	return o.ToOpenShiftManagedClusterAuthProfilePtrOutputWithContext(context.Background())
}

func (o OpenShiftManagedClusterAuthProfileOutput) ToOpenShiftManagedClusterAuthProfilePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAuthProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenShiftManagedClusterAuthProfile) *OpenShiftManagedClusterAuthProfile {
		return &v
	}).(OpenShiftManagedClusterAuthProfilePtrOutput)
}

func (o OpenShiftManagedClusterAuthProfileOutput) IdentityProviders() OpenShiftManagedClusterIdentityProviderArrayOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAuthProfile) []OpenShiftManagedClusterIdentityProvider {
		return v.IdentityProviders
	}).(OpenShiftManagedClusterIdentityProviderArrayOutput)
}

type OpenShiftManagedClusterAuthProfilePtrOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAuthProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedClusterAuthProfile)(nil)).Elem()
}

func (o OpenShiftManagedClusterAuthProfilePtrOutput) ToOpenShiftManagedClusterAuthProfilePtrOutput() OpenShiftManagedClusterAuthProfilePtrOutput {
	return o
}

func (o OpenShiftManagedClusterAuthProfilePtrOutput) ToOpenShiftManagedClusterAuthProfilePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAuthProfilePtrOutput {
	return o
}

func (o OpenShiftManagedClusterAuthProfilePtrOutput) Elem() OpenShiftManagedClusterAuthProfileOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAuthProfile) OpenShiftManagedClusterAuthProfile {
		if v != nil {
			return *v
		}
		var ret OpenShiftManagedClusterAuthProfile
		return ret
	}).(OpenShiftManagedClusterAuthProfileOutput)
}

func (o OpenShiftManagedClusterAuthProfilePtrOutput) IdentityProviders() OpenShiftManagedClusterIdentityProviderArrayOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAuthProfile) []OpenShiftManagedClusterIdentityProvider {
		if v == nil {
			return nil
		}
		return v.IdentityProviders
	}).(OpenShiftManagedClusterIdentityProviderArrayOutput)
}

type OpenShiftManagedClusterAuthProfileResponse struct {
	IdentityProviders []OpenShiftManagedClusterIdentityProviderResponse `pulumi:"identityProviders"`
}

type OpenShiftManagedClusterAuthProfileResponseOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAuthProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterAuthProfileResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterAuthProfileResponseOutput) ToOpenShiftManagedClusterAuthProfileResponseOutput() OpenShiftManagedClusterAuthProfileResponseOutput {
	return o
}

func (o OpenShiftManagedClusterAuthProfileResponseOutput) ToOpenShiftManagedClusterAuthProfileResponseOutputWithContext(ctx context.Context) OpenShiftManagedClusterAuthProfileResponseOutput {
	return o
}

func (o OpenShiftManagedClusterAuthProfileResponseOutput) IdentityProviders() OpenShiftManagedClusterIdentityProviderResponseArrayOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterAuthProfileResponse) []OpenShiftManagedClusterIdentityProviderResponse {
		return v.IdentityProviders
	}).(OpenShiftManagedClusterIdentityProviderResponseArrayOutput)
}

type OpenShiftManagedClusterAuthProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterAuthProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedClusterAuthProfileResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterAuthProfileResponsePtrOutput) ToOpenShiftManagedClusterAuthProfileResponsePtrOutput() OpenShiftManagedClusterAuthProfileResponsePtrOutput {
	return o
}

func (o OpenShiftManagedClusterAuthProfileResponsePtrOutput) ToOpenShiftManagedClusterAuthProfileResponsePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterAuthProfileResponsePtrOutput {
	return o
}

func (o OpenShiftManagedClusterAuthProfileResponsePtrOutput) Elem() OpenShiftManagedClusterAuthProfileResponseOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAuthProfileResponse) OpenShiftManagedClusterAuthProfileResponse {
		if v != nil {
			return *v
		}
		var ret OpenShiftManagedClusterAuthProfileResponse
		return ret
	}).(OpenShiftManagedClusterAuthProfileResponseOutput)
}

func (o OpenShiftManagedClusterAuthProfileResponsePtrOutput) IdentityProviders() OpenShiftManagedClusterIdentityProviderResponseArrayOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterAuthProfileResponse) []OpenShiftManagedClusterIdentityProviderResponse {
		if v == nil {
			return nil
		}
		return v.IdentityProviders
	}).(OpenShiftManagedClusterIdentityProviderResponseArrayOutput)
}

type OpenShiftManagedClusterIdentityProvider struct {
	Name     *string                                     `pulumi:"name"`
	Provider *OpenShiftManagedClusterAADIdentityProvider `pulumi:"provider"`
}





type OpenShiftManagedClusterIdentityProviderInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterIdentityProviderOutput() OpenShiftManagedClusterIdentityProviderOutput
	ToOpenShiftManagedClusterIdentityProviderOutputWithContext(context.Context) OpenShiftManagedClusterIdentityProviderOutput
}

type OpenShiftManagedClusterIdentityProviderArgs struct {
	Name     pulumi.StringPtrInput                              `pulumi:"name"`
	Provider OpenShiftManagedClusterAADIdentityProviderPtrInput `pulumi:"provider"`
}

func (OpenShiftManagedClusterIdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterIdentityProvider)(nil)).Elem()
}

func (i OpenShiftManagedClusterIdentityProviderArgs) ToOpenShiftManagedClusterIdentityProviderOutput() OpenShiftManagedClusterIdentityProviderOutput {
	return i.ToOpenShiftManagedClusterIdentityProviderOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterIdentityProviderArgs) ToOpenShiftManagedClusterIdentityProviderOutputWithContext(ctx context.Context) OpenShiftManagedClusterIdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterIdentityProviderOutput)
}





type OpenShiftManagedClusterIdentityProviderArrayInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterIdentityProviderArrayOutput() OpenShiftManagedClusterIdentityProviderArrayOutput
	ToOpenShiftManagedClusterIdentityProviderArrayOutputWithContext(context.Context) OpenShiftManagedClusterIdentityProviderArrayOutput
}

type OpenShiftManagedClusterIdentityProviderArray []OpenShiftManagedClusterIdentityProviderInput

func (OpenShiftManagedClusterIdentityProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenShiftManagedClusterIdentityProvider)(nil)).Elem()
}

func (i OpenShiftManagedClusterIdentityProviderArray) ToOpenShiftManagedClusterIdentityProviderArrayOutput() OpenShiftManagedClusterIdentityProviderArrayOutput {
	return i.ToOpenShiftManagedClusterIdentityProviderArrayOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterIdentityProviderArray) ToOpenShiftManagedClusterIdentityProviderArrayOutputWithContext(ctx context.Context) OpenShiftManagedClusterIdentityProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterIdentityProviderArrayOutput)
}

type OpenShiftManagedClusterIdentityProviderOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterIdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterIdentityProvider)(nil)).Elem()
}

func (o OpenShiftManagedClusterIdentityProviderOutput) ToOpenShiftManagedClusterIdentityProviderOutput() OpenShiftManagedClusterIdentityProviderOutput {
	return o
}

func (o OpenShiftManagedClusterIdentityProviderOutput) ToOpenShiftManagedClusterIdentityProviderOutputWithContext(ctx context.Context) OpenShiftManagedClusterIdentityProviderOutput {
	return o
}

func (o OpenShiftManagedClusterIdentityProviderOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterIdentityProvider) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterIdentityProviderOutput) Provider() OpenShiftManagedClusterAADIdentityProviderPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterIdentityProvider) *OpenShiftManagedClusterAADIdentityProvider {
		return v.Provider
	}).(OpenShiftManagedClusterAADIdentityProviderPtrOutput)
}

type OpenShiftManagedClusterIdentityProviderArrayOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterIdentityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenShiftManagedClusterIdentityProvider)(nil)).Elem()
}

func (o OpenShiftManagedClusterIdentityProviderArrayOutput) ToOpenShiftManagedClusterIdentityProviderArrayOutput() OpenShiftManagedClusterIdentityProviderArrayOutput {
	return o
}

func (o OpenShiftManagedClusterIdentityProviderArrayOutput) ToOpenShiftManagedClusterIdentityProviderArrayOutputWithContext(ctx context.Context) OpenShiftManagedClusterIdentityProviderArrayOutput {
	return o
}

func (o OpenShiftManagedClusterIdentityProviderArrayOutput) Index(i pulumi.IntInput) OpenShiftManagedClusterIdentityProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpenShiftManagedClusterIdentityProvider {
		return vs[0].([]OpenShiftManagedClusterIdentityProvider)[vs[1].(int)]
	}).(OpenShiftManagedClusterIdentityProviderOutput)
}

type OpenShiftManagedClusterIdentityProviderResponse struct {
	Name     *string                                             `pulumi:"name"`
	Provider *OpenShiftManagedClusterAADIdentityProviderResponse `pulumi:"provider"`
}

type OpenShiftManagedClusterIdentityProviderResponseOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterIdentityProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterIdentityProviderResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterIdentityProviderResponseOutput) ToOpenShiftManagedClusterIdentityProviderResponseOutput() OpenShiftManagedClusterIdentityProviderResponseOutput {
	return o
}

func (o OpenShiftManagedClusterIdentityProviderResponseOutput) ToOpenShiftManagedClusterIdentityProviderResponseOutputWithContext(ctx context.Context) OpenShiftManagedClusterIdentityProviderResponseOutput {
	return o
}

func (o OpenShiftManagedClusterIdentityProviderResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterIdentityProviderResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterIdentityProviderResponseOutput) Provider() OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterIdentityProviderResponse) *OpenShiftManagedClusterAADIdentityProviderResponse {
		return v.Provider
	}).(OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput)
}

type OpenShiftManagedClusterIdentityProviderResponseArrayOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterIdentityProviderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenShiftManagedClusterIdentityProviderResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterIdentityProviderResponseArrayOutput) ToOpenShiftManagedClusterIdentityProviderResponseArrayOutput() OpenShiftManagedClusterIdentityProviderResponseArrayOutput {
	return o
}

func (o OpenShiftManagedClusterIdentityProviderResponseArrayOutput) ToOpenShiftManagedClusterIdentityProviderResponseArrayOutputWithContext(ctx context.Context) OpenShiftManagedClusterIdentityProviderResponseArrayOutput {
	return o
}

func (o OpenShiftManagedClusterIdentityProviderResponseArrayOutput) Index(i pulumi.IntInput) OpenShiftManagedClusterIdentityProviderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpenShiftManagedClusterIdentityProviderResponse {
		return vs[0].([]OpenShiftManagedClusterIdentityProviderResponse)[vs[1].(int)]
	}).(OpenShiftManagedClusterIdentityProviderResponseOutput)
}

type OpenShiftManagedClusterMasterPoolProfile struct {
	Count      int     `pulumi:"count"`
	Name       *string `pulumi:"name"`
	OsType     *string `pulumi:"osType"`
	SubnetCidr *string `pulumi:"subnetCidr"`
	VmSize     string  `pulumi:"vmSize"`
}





type OpenShiftManagedClusterMasterPoolProfileInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterMasterPoolProfileOutput() OpenShiftManagedClusterMasterPoolProfileOutput
	ToOpenShiftManagedClusterMasterPoolProfileOutputWithContext(context.Context) OpenShiftManagedClusterMasterPoolProfileOutput
}

type OpenShiftManagedClusterMasterPoolProfileArgs struct {
	Count      pulumi.IntInput       `pulumi:"count"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	OsType     pulumi.StringPtrInput `pulumi:"osType"`
	SubnetCidr pulumi.StringPtrInput `pulumi:"subnetCidr"`
	VmSize     pulumi.StringInput    `pulumi:"vmSize"`
}

func (OpenShiftManagedClusterMasterPoolProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterMasterPoolProfile)(nil)).Elem()
}

func (i OpenShiftManagedClusterMasterPoolProfileArgs) ToOpenShiftManagedClusterMasterPoolProfileOutput() OpenShiftManagedClusterMasterPoolProfileOutput {
	return i.ToOpenShiftManagedClusterMasterPoolProfileOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterMasterPoolProfileArgs) ToOpenShiftManagedClusterMasterPoolProfileOutputWithContext(ctx context.Context) OpenShiftManagedClusterMasterPoolProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterMasterPoolProfileOutput)
}

func (i OpenShiftManagedClusterMasterPoolProfileArgs) ToOpenShiftManagedClusterMasterPoolProfilePtrOutput() OpenShiftManagedClusterMasterPoolProfilePtrOutput {
	return i.ToOpenShiftManagedClusterMasterPoolProfilePtrOutputWithContext(context.Background())
}

func (i OpenShiftManagedClusterMasterPoolProfileArgs) ToOpenShiftManagedClusterMasterPoolProfilePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterMasterPoolProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterMasterPoolProfileOutput).ToOpenShiftManagedClusterMasterPoolProfilePtrOutputWithContext(ctx)
}









type OpenShiftManagedClusterMasterPoolProfilePtrInput interface {
	pulumi.Input

	ToOpenShiftManagedClusterMasterPoolProfilePtrOutput() OpenShiftManagedClusterMasterPoolProfilePtrOutput
	ToOpenShiftManagedClusterMasterPoolProfilePtrOutputWithContext(context.Context) OpenShiftManagedClusterMasterPoolProfilePtrOutput
}

type openShiftManagedClusterMasterPoolProfilePtrType OpenShiftManagedClusterMasterPoolProfileArgs

func OpenShiftManagedClusterMasterPoolProfilePtr(v *OpenShiftManagedClusterMasterPoolProfileArgs) OpenShiftManagedClusterMasterPoolProfilePtrInput {
	return (*openShiftManagedClusterMasterPoolProfilePtrType)(v)
}

func (*openShiftManagedClusterMasterPoolProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedClusterMasterPoolProfile)(nil)).Elem()
}

func (i *openShiftManagedClusterMasterPoolProfilePtrType) ToOpenShiftManagedClusterMasterPoolProfilePtrOutput() OpenShiftManagedClusterMasterPoolProfilePtrOutput {
	return i.ToOpenShiftManagedClusterMasterPoolProfilePtrOutputWithContext(context.Background())
}

func (i *openShiftManagedClusterMasterPoolProfilePtrType) ToOpenShiftManagedClusterMasterPoolProfilePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterMasterPoolProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftManagedClusterMasterPoolProfilePtrOutput)
}

type OpenShiftManagedClusterMasterPoolProfileOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterMasterPoolProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterMasterPoolProfile)(nil)).Elem()
}

func (o OpenShiftManagedClusterMasterPoolProfileOutput) ToOpenShiftManagedClusterMasterPoolProfileOutput() OpenShiftManagedClusterMasterPoolProfileOutput {
	return o
}

func (o OpenShiftManagedClusterMasterPoolProfileOutput) ToOpenShiftManagedClusterMasterPoolProfileOutputWithContext(ctx context.Context) OpenShiftManagedClusterMasterPoolProfileOutput {
	return o
}

func (o OpenShiftManagedClusterMasterPoolProfileOutput) ToOpenShiftManagedClusterMasterPoolProfilePtrOutput() OpenShiftManagedClusterMasterPoolProfilePtrOutput {
	return o.ToOpenShiftManagedClusterMasterPoolProfilePtrOutputWithContext(context.Background())
}

func (o OpenShiftManagedClusterMasterPoolProfileOutput) ToOpenShiftManagedClusterMasterPoolProfilePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterMasterPoolProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenShiftManagedClusterMasterPoolProfile) *OpenShiftManagedClusterMasterPoolProfile {
		return &v
	}).(OpenShiftManagedClusterMasterPoolProfilePtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfile) int { return v.Count }).(pulumi.IntOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileOutput) SubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfile) *string { return v.SubnetCidr }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfile) string { return v.VmSize }).(pulumi.StringOutput)
}

type OpenShiftManagedClusterMasterPoolProfilePtrOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterMasterPoolProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedClusterMasterPoolProfile)(nil)).Elem()
}

func (o OpenShiftManagedClusterMasterPoolProfilePtrOutput) ToOpenShiftManagedClusterMasterPoolProfilePtrOutput() OpenShiftManagedClusterMasterPoolProfilePtrOutput {
	return o
}

func (o OpenShiftManagedClusterMasterPoolProfilePtrOutput) ToOpenShiftManagedClusterMasterPoolProfilePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterMasterPoolProfilePtrOutput {
	return o
}

func (o OpenShiftManagedClusterMasterPoolProfilePtrOutput) Elem() OpenShiftManagedClusterMasterPoolProfileOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfile) OpenShiftManagedClusterMasterPoolProfile {
		if v != nil {
			return *v
		}
		var ret OpenShiftManagedClusterMasterPoolProfile
		return ret
	}).(OpenShiftManagedClusterMasterPoolProfileOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfilePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfile) *int {
		if v == nil {
			return nil
		}
		return &v.Count
	}).(pulumi.IntPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfilePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfile) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfilePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfile) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfilePtrOutput) SubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfile) *string {
		if v == nil {
			return nil
		}
		return v.SubnetCidr
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfile) *string {
		if v == nil {
			return nil
		}
		return &v.VmSize
	}).(pulumi.StringPtrOutput)
}

type OpenShiftManagedClusterMasterPoolProfileResponse struct {
	Count      int     `pulumi:"count"`
	Name       *string `pulumi:"name"`
	OsType     *string `pulumi:"osType"`
	SubnetCidr *string `pulumi:"subnetCidr"`
	VmSize     string  `pulumi:"vmSize"`
}

type OpenShiftManagedClusterMasterPoolProfileResponseOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterMasterPoolProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftManagedClusterMasterPoolProfileResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterMasterPoolProfileResponseOutput) ToOpenShiftManagedClusterMasterPoolProfileResponseOutput() OpenShiftManagedClusterMasterPoolProfileResponseOutput {
	return o
}

func (o OpenShiftManagedClusterMasterPoolProfileResponseOutput) ToOpenShiftManagedClusterMasterPoolProfileResponseOutputWithContext(ctx context.Context) OpenShiftManagedClusterMasterPoolProfileResponseOutput {
	return o
}

func (o OpenShiftManagedClusterMasterPoolProfileResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfileResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileResponseOutput) SubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfileResponse) *string { return v.SubnetCidr }).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileResponseOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftManagedClusterMasterPoolProfileResponse) string { return v.VmSize }).(pulumi.StringOutput)
}

type OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenShiftManagedClusterMasterPoolProfileResponse)(nil)).Elem()
}

func (o OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput) ToOpenShiftManagedClusterMasterPoolProfileResponsePtrOutput() OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput {
	return o
}

func (o OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput) ToOpenShiftManagedClusterMasterPoolProfileResponsePtrOutputWithContext(ctx context.Context) OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput {
	return o
}

func (o OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput) Elem() OpenShiftManagedClusterMasterPoolProfileResponseOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfileResponse) OpenShiftManagedClusterMasterPoolProfileResponse {
		if v != nil {
			return *v
		}
		var ret OpenShiftManagedClusterMasterPoolProfileResponse
		return ret
	}).(OpenShiftManagedClusterMasterPoolProfileResponseOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfileResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Count
	}).(pulumi.IntPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput) SubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetCidr
	}).(pulumi.StringPtrOutput)
}

func (o OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenShiftManagedClusterMasterPoolProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VmSize
	}).(pulumi.StringPtrOutput)
}

type OpenShiftRouterProfile struct {
	Name *string `pulumi:"name"`
}





type OpenShiftRouterProfileInput interface {
	pulumi.Input

	ToOpenShiftRouterProfileOutput() OpenShiftRouterProfileOutput
	ToOpenShiftRouterProfileOutputWithContext(context.Context) OpenShiftRouterProfileOutput
}

type OpenShiftRouterProfileArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (OpenShiftRouterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftRouterProfile)(nil)).Elem()
}

func (i OpenShiftRouterProfileArgs) ToOpenShiftRouterProfileOutput() OpenShiftRouterProfileOutput {
	return i.ToOpenShiftRouterProfileOutputWithContext(context.Background())
}

func (i OpenShiftRouterProfileArgs) ToOpenShiftRouterProfileOutputWithContext(ctx context.Context) OpenShiftRouterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftRouterProfileOutput)
}





type OpenShiftRouterProfileArrayInput interface {
	pulumi.Input

	ToOpenShiftRouterProfileArrayOutput() OpenShiftRouterProfileArrayOutput
	ToOpenShiftRouterProfileArrayOutputWithContext(context.Context) OpenShiftRouterProfileArrayOutput
}

type OpenShiftRouterProfileArray []OpenShiftRouterProfileInput

func (OpenShiftRouterProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenShiftRouterProfile)(nil)).Elem()
}

func (i OpenShiftRouterProfileArray) ToOpenShiftRouterProfileArrayOutput() OpenShiftRouterProfileArrayOutput {
	return i.ToOpenShiftRouterProfileArrayOutputWithContext(context.Background())
}

func (i OpenShiftRouterProfileArray) ToOpenShiftRouterProfileArrayOutputWithContext(ctx context.Context) OpenShiftRouterProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenShiftRouterProfileArrayOutput)
}

type OpenShiftRouterProfileOutput struct{ *pulumi.OutputState }

func (OpenShiftRouterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftRouterProfile)(nil)).Elem()
}

func (o OpenShiftRouterProfileOutput) ToOpenShiftRouterProfileOutput() OpenShiftRouterProfileOutput {
	return o
}

func (o OpenShiftRouterProfileOutput) ToOpenShiftRouterProfileOutputWithContext(ctx context.Context) OpenShiftRouterProfileOutput {
	return o
}

func (o OpenShiftRouterProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftRouterProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type OpenShiftRouterProfileArrayOutput struct{ *pulumi.OutputState }

func (OpenShiftRouterProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenShiftRouterProfile)(nil)).Elem()
}

func (o OpenShiftRouterProfileArrayOutput) ToOpenShiftRouterProfileArrayOutput() OpenShiftRouterProfileArrayOutput {
	return o
}

func (o OpenShiftRouterProfileArrayOutput) ToOpenShiftRouterProfileArrayOutputWithContext(ctx context.Context) OpenShiftRouterProfileArrayOutput {
	return o
}

func (o OpenShiftRouterProfileArrayOutput) Index(i pulumi.IntInput) OpenShiftRouterProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpenShiftRouterProfile {
		return vs[0].([]OpenShiftRouterProfile)[vs[1].(int)]
	}).(OpenShiftRouterProfileOutput)
}

type OpenShiftRouterProfileResponse struct {
	Fqdn            string  `pulumi:"fqdn"`
	Name            *string `pulumi:"name"`
	PublicSubdomain string  `pulumi:"publicSubdomain"`
}

type OpenShiftRouterProfileResponseOutput struct{ *pulumi.OutputState }

func (OpenShiftRouterProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenShiftRouterProfileResponse)(nil)).Elem()
}

func (o OpenShiftRouterProfileResponseOutput) ToOpenShiftRouterProfileResponseOutput() OpenShiftRouterProfileResponseOutput {
	return o
}

func (o OpenShiftRouterProfileResponseOutput) ToOpenShiftRouterProfileResponseOutputWithContext(ctx context.Context) OpenShiftRouterProfileResponseOutput {
	return o
}

func (o OpenShiftRouterProfileResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftRouterProfileResponse) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o OpenShiftRouterProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenShiftRouterProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OpenShiftRouterProfileResponseOutput) PublicSubdomain() pulumi.StringOutput {
	return o.ApplyT(func(v OpenShiftRouterProfileResponse) string { return v.PublicSubdomain }).(pulumi.StringOutput)
}

type OpenShiftRouterProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (OpenShiftRouterProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenShiftRouterProfileResponse)(nil)).Elem()
}

func (o OpenShiftRouterProfileResponseArrayOutput) ToOpenShiftRouterProfileResponseArrayOutput() OpenShiftRouterProfileResponseArrayOutput {
	return o
}

func (o OpenShiftRouterProfileResponseArrayOutput) ToOpenShiftRouterProfileResponseArrayOutputWithContext(ctx context.Context) OpenShiftRouterProfileResponseArrayOutput {
	return o
}

func (o OpenShiftRouterProfileResponseArrayOutput) Index(i pulumi.IntInput) OpenShiftRouterProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpenShiftRouterProfileResponse {
		return vs[0].([]OpenShiftRouterProfileResponse)[vs[1].(int)]
	}).(OpenShiftRouterProfileResponseOutput)
}

type PurchasePlan struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}





type PurchasePlanInput interface {
	pulumi.Input

	ToPurchasePlanOutput() PurchasePlanOutput
	ToPurchasePlanOutputWithContext(context.Context) PurchasePlanOutput
}

type PurchasePlanArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
}

func (PurchasePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PurchasePlan)(nil)).Elem()
}

func (i PurchasePlanArgs) ToPurchasePlanOutput() PurchasePlanOutput {
	return i.ToPurchasePlanOutputWithContext(context.Background())
}

func (i PurchasePlanArgs) ToPurchasePlanOutputWithContext(ctx context.Context) PurchasePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurchasePlanOutput)
}

func (i PurchasePlanArgs) ToPurchasePlanPtrOutput() PurchasePlanPtrOutput {
	return i.ToPurchasePlanPtrOutputWithContext(context.Background())
}

func (i PurchasePlanArgs) ToPurchasePlanPtrOutputWithContext(ctx context.Context) PurchasePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurchasePlanOutput).ToPurchasePlanPtrOutputWithContext(ctx)
}









type PurchasePlanPtrInput interface {
	pulumi.Input

	ToPurchasePlanPtrOutput() PurchasePlanPtrOutput
	ToPurchasePlanPtrOutputWithContext(context.Context) PurchasePlanPtrOutput
}

type purchasePlanPtrType PurchasePlanArgs

func PurchasePlanPtr(v *PurchasePlanArgs) PurchasePlanPtrInput {
	return (*purchasePlanPtrType)(v)
}

func (*purchasePlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PurchasePlan)(nil)).Elem()
}

func (i *purchasePlanPtrType) ToPurchasePlanPtrOutput() PurchasePlanPtrOutput {
	return i.ToPurchasePlanPtrOutputWithContext(context.Background())
}

func (i *purchasePlanPtrType) ToPurchasePlanPtrOutputWithContext(ctx context.Context) PurchasePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurchasePlanPtrOutput)
}

type PurchasePlanOutput struct{ *pulumi.OutputState }

func (PurchasePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurchasePlan)(nil)).Elem()
}

func (o PurchasePlanOutput) ToPurchasePlanOutput() PurchasePlanOutput {
	return o
}

func (o PurchasePlanOutput) ToPurchasePlanOutputWithContext(ctx context.Context) PurchasePlanOutput {
	return o
}

func (o PurchasePlanOutput) ToPurchasePlanPtrOutput() PurchasePlanPtrOutput {
	return o.ToPurchasePlanPtrOutputWithContext(context.Background())
}

func (o PurchasePlanOutput) ToPurchasePlanPtrOutputWithContext(ctx context.Context) PurchasePlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PurchasePlan) *PurchasePlan {
		return &v
	}).(PurchasePlanPtrOutput)
}

func (o PurchasePlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PurchasePlanOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlan) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PurchasePlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PurchasePlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type PurchasePlanPtrOutput struct{ *pulumi.OutputState }

func (PurchasePlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurchasePlan)(nil)).Elem()
}

func (o PurchasePlanPtrOutput) ToPurchasePlanPtrOutput() PurchasePlanPtrOutput {
	return o
}

func (o PurchasePlanPtrOutput) ToPurchasePlanPtrOutputWithContext(ctx context.Context) PurchasePlanPtrOutput {
	return o
}

func (o PurchasePlanPtrOutput) Elem() PurchasePlanOutput {
	return o.ApplyT(func(v *PurchasePlan) PurchasePlan {
		if v != nil {
			return *v
		}
		var ret PurchasePlan
		return ret
	}).(PurchasePlanOutput)
}

func (o PurchasePlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlan) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type PurchasePlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}

type PurchasePlanResponseOutput struct{ *pulumi.OutputState }

func (PurchasePlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurchasePlanResponse)(nil)).Elem()
}

func (o PurchasePlanResponseOutput) ToPurchasePlanResponseOutput() PurchasePlanResponseOutput {
	return o
}

func (o PurchasePlanResponseOutput) ToPurchasePlanResponseOutputWithContext(ctx context.Context) PurchasePlanResponseOutput {
	return o
}

func (o PurchasePlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type PurchasePlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PurchasePlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurchasePlanResponse)(nil)).Elem()
}

func (o PurchasePlanResponsePtrOutput) ToPurchasePlanResponsePtrOutput() PurchasePlanResponsePtrOutput {
	return o
}

func (o PurchasePlanResponsePtrOutput) ToPurchasePlanResponsePtrOutputWithContext(ctx context.Context) PurchasePlanResponsePtrOutput {
	return o
}

func (o PurchasePlanResponsePtrOutput) Elem() PurchasePlanResponseOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) PurchasePlanResponse {
		if v != nil {
			return *v
		}
		var ret PurchasePlanResponse
		return ret
	}).(PurchasePlanResponseOutput)
}

func (o PurchasePlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAADIdentityProviderOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAADIdentityProviderPtrOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAADIdentityProviderResponseOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAADIdentityProviderResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAgentPoolProfileOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAgentPoolProfileArrayOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAgentPoolProfileResponseOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAgentPoolProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAuthProfileOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAuthProfilePtrOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAuthProfileResponseOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterAuthProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterIdentityProviderOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterIdentityProviderArrayOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterIdentityProviderResponseOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterIdentityProviderResponseArrayOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterMasterPoolProfileOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterMasterPoolProfilePtrOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterMasterPoolProfileResponseOutput{})
	pulumi.RegisterOutputType(OpenShiftManagedClusterMasterPoolProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenShiftRouterProfileOutput{})
	pulumi.RegisterOutputType(OpenShiftRouterProfileArrayOutput{})
	pulumi.RegisterOutputType(OpenShiftRouterProfileResponseOutput{})
	pulumi.RegisterOutputType(OpenShiftRouterProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(PurchasePlanOutput{})
	pulumi.RegisterOutputType(PurchasePlanPtrOutput{})
	pulumi.RegisterOutputType(PurchasePlanResponseOutput{})
	pulumi.RegisterOutputType(PurchasePlanResponsePtrOutput{})
}
