


package v20170831

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerServiceAgentPoolProfile struct {
	Count          *int    `pulumi:"count"`
	DnsPrefix      *string `pulumi:"dnsPrefix"`
	Name           string  `pulumi:"name"`
	OsDiskSizeGB   *int    `pulumi:"osDiskSizeGB"`
	OsType         *string `pulumi:"osType"`
	Ports          []int   `pulumi:"ports"`
	StorageProfile *string `pulumi:"storageProfile"`
	VmSize         string  `pulumi:"vmSize"`
	VnetSubnetID   *string `pulumi:"vnetSubnetID"`
}


func (val *ContainerServiceAgentPoolProfile) Defaults() *ContainerServiceAgentPoolProfile {
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





type ContainerServiceAgentPoolProfileInput interface {
	pulumi.Input

	ToContainerServiceAgentPoolProfileOutput() ContainerServiceAgentPoolProfileOutput
	ToContainerServiceAgentPoolProfileOutputWithContext(context.Context) ContainerServiceAgentPoolProfileOutput
}

type ContainerServiceAgentPoolProfileArgs struct {
	Count          pulumi.IntPtrInput    `pulumi:"count"`
	DnsPrefix      pulumi.StringPtrInput `pulumi:"dnsPrefix"`
	Name           pulumi.StringInput    `pulumi:"name"`
	OsDiskSizeGB   pulumi.IntPtrInput    `pulumi:"osDiskSizeGB"`
	OsType         pulumi.StringPtrInput `pulumi:"osType"`
	Ports          pulumi.IntArrayInput  `pulumi:"ports"`
	StorageProfile pulumi.StringPtrInput `pulumi:"storageProfile"`
	VmSize         pulumi.StringInput    `pulumi:"vmSize"`
	VnetSubnetID   pulumi.StringPtrInput `pulumi:"vnetSubnetID"`
}

func (ContainerServiceAgentPoolProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceAgentPoolProfile)(nil)).Elem()
}

func (i ContainerServiceAgentPoolProfileArgs) ToContainerServiceAgentPoolProfileOutput() ContainerServiceAgentPoolProfileOutput {
	return i.ToContainerServiceAgentPoolProfileOutputWithContext(context.Background())
}

func (i ContainerServiceAgentPoolProfileArgs) ToContainerServiceAgentPoolProfileOutputWithContext(ctx context.Context) ContainerServiceAgentPoolProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceAgentPoolProfileOutput)
}





type ContainerServiceAgentPoolProfileArrayInput interface {
	pulumi.Input

	ToContainerServiceAgentPoolProfileArrayOutput() ContainerServiceAgentPoolProfileArrayOutput
	ToContainerServiceAgentPoolProfileArrayOutputWithContext(context.Context) ContainerServiceAgentPoolProfileArrayOutput
}

type ContainerServiceAgentPoolProfileArray []ContainerServiceAgentPoolProfileInput

func (ContainerServiceAgentPoolProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceAgentPoolProfile)(nil)).Elem()
}

func (i ContainerServiceAgentPoolProfileArray) ToContainerServiceAgentPoolProfileArrayOutput() ContainerServiceAgentPoolProfileArrayOutput {
	return i.ToContainerServiceAgentPoolProfileArrayOutputWithContext(context.Background())
}

func (i ContainerServiceAgentPoolProfileArray) ToContainerServiceAgentPoolProfileArrayOutputWithContext(ctx context.Context) ContainerServiceAgentPoolProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceAgentPoolProfileArrayOutput)
}

type ContainerServiceAgentPoolProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceAgentPoolProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceAgentPoolProfile)(nil)).Elem()
}

func (o ContainerServiceAgentPoolProfileOutput) ToContainerServiceAgentPoolProfileOutput() ContainerServiceAgentPoolProfileOutput {
	return o
}

func (o ContainerServiceAgentPoolProfileOutput) ToContainerServiceAgentPoolProfileOutputWithContext(ctx context.Context) ContainerServiceAgentPoolProfileOutput {
	return o
}

func (o ContainerServiceAgentPoolProfileOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) *string { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) Ports() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) []int { return v.Ports }).(pulumi.IntArrayOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) StorageProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) *string { return v.StorageProfile }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) string { return v.VmSize }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

type ContainerServiceAgentPoolProfileArrayOutput struct{ *pulumi.OutputState }

func (ContainerServiceAgentPoolProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceAgentPoolProfile)(nil)).Elem()
}

func (o ContainerServiceAgentPoolProfileArrayOutput) ToContainerServiceAgentPoolProfileArrayOutput() ContainerServiceAgentPoolProfileArrayOutput {
	return o
}

func (o ContainerServiceAgentPoolProfileArrayOutput) ToContainerServiceAgentPoolProfileArrayOutputWithContext(ctx context.Context) ContainerServiceAgentPoolProfileArrayOutput {
	return o
}

func (o ContainerServiceAgentPoolProfileArrayOutput) Index(i pulumi.IntInput) ContainerServiceAgentPoolProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerServiceAgentPoolProfile {
		return vs[0].([]ContainerServiceAgentPoolProfile)[vs[1].(int)]
	}).(ContainerServiceAgentPoolProfileOutput)
}

type ContainerServiceAgentPoolProfileResponse struct {
	Count          *int    `pulumi:"count"`
	DnsPrefix      *string `pulumi:"dnsPrefix"`
	Fqdn           string  `pulumi:"fqdn"`
	Name           string  `pulumi:"name"`
	OsDiskSizeGB   *int    `pulumi:"osDiskSizeGB"`
	OsType         *string `pulumi:"osType"`
	Ports          []int   `pulumi:"ports"`
	StorageProfile *string `pulumi:"storageProfile"`
	VmSize         string  `pulumi:"vmSize"`
	VnetSubnetID   *string `pulumi:"vnetSubnetID"`
}


func (val *ContainerServiceAgentPoolProfileResponse) Defaults() *ContainerServiceAgentPoolProfileResponse {
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





type ContainerServiceAgentPoolProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceAgentPoolProfileResponseOutput() ContainerServiceAgentPoolProfileResponseOutput
	ToContainerServiceAgentPoolProfileResponseOutputWithContext(context.Context) ContainerServiceAgentPoolProfileResponseOutput
}

type ContainerServiceAgentPoolProfileResponseArgs struct {
	Count          pulumi.IntPtrInput    `pulumi:"count"`
	DnsPrefix      pulumi.StringPtrInput `pulumi:"dnsPrefix"`
	Fqdn           pulumi.StringInput    `pulumi:"fqdn"`
	Name           pulumi.StringInput    `pulumi:"name"`
	OsDiskSizeGB   pulumi.IntPtrInput    `pulumi:"osDiskSizeGB"`
	OsType         pulumi.StringPtrInput `pulumi:"osType"`
	Ports          pulumi.IntArrayInput  `pulumi:"ports"`
	StorageProfile pulumi.StringPtrInput `pulumi:"storageProfile"`
	VmSize         pulumi.StringInput    `pulumi:"vmSize"`
	VnetSubnetID   pulumi.StringPtrInput `pulumi:"vnetSubnetID"`
}

func (ContainerServiceAgentPoolProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceAgentPoolProfileResponse)(nil)).Elem()
}

func (i ContainerServiceAgentPoolProfileResponseArgs) ToContainerServiceAgentPoolProfileResponseOutput() ContainerServiceAgentPoolProfileResponseOutput {
	return i.ToContainerServiceAgentPoolProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceAgentPoolProfileResponseArgs) ToContainerServiceAgentPoolProfileResponseOutputWithContext(ctx context.Context) ContainerServiceAgentPoolProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceAgentPoolProfileResponseOutput)
}





type ContainerServiceAgentPoolProfileResponseArrayInput interface {
	pulumi.Input

	ToContainerServiceAgentPoolProfileResponseArrayOutput() ContainerServiceAgentPoolProfileResponseArrayOutput
	ToContainerServiceAgentPoolProfileResponseArrayOutputWithContext(context.Context) ContainerServiceAgentPoolProfileResponseArrayOutput
}

type ContainerServiceAgentPoolProfileResponseArray []ContainerServiceAgentPoolProfileResponseInput

func (ContainerServiceAgentPoolProfileResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceAgentPoolProfileResponse)(nil)).Elem()
}

func (i ContainerServiceAgentPoolProfileResponseArray) ToContainerServiceAgentPoolProfileResponseArrayOutput() ContainerServiceAgentPoolProfileResponseArrayOutput {
	return i.ToContainerServiceAgentPoolProfileResponseArrayOutputWithContext(context.Background())
}

func (i ContainerServiceAgentPoolProfileResponseArray) ToContainerServiceAgentPoolProfileResponseArrayOutputWithContext(ctx context.Context) ContainerServiceAgentPoolProfileResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceAgentPoolProfileResponseArrayOutput)
}

type ContainerServiceAgentPoolProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceAgentPoolProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceAgentPoolProfileResponse)(nil)).Elem()
}

func (o ContainerServiceAgentPoolProfileResponseOutput) ToContainerServiceAgentPoolProfileResponseOutput() ContainerServiceAgentPoolProfileResponseOutput {
	return o
}

func (o ContainerServiceAgentPoolProfileResponseOutput) ToContainerServiceAgentPoolProfileResponseOutputWithContext(ctx context.Context) ContainerServiceAgentPoolProfileResponseOutput {
	return o
}

func (o ContainerServiceAgentPoolProfileResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) *string { return v.DnsPrefix }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) Ports() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) []int { return v.Ports }).(pulumi.IntArrayOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) StorageProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) *string { return v.StorageProfile }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) string { return v.VmSize }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

type ContainerServiceAgentPoolProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerServiceAgentPoolProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceAgentPoolProfileResponse)(nil)).Elem()
}

func (o ContainerServiceAgentPoolProfileResponseArrayOutput) ToContainerServiceAgentPoolProfileResponseArrayOutput() ContainerServiceAgentPoolProfileResponseArrayOutput {
	return o
}

func (o ContainerServiceAgentPoolProfileResponseArrayOutput) ToContainerServiceAgentPoolProfileResponseArrayOutputWithContext(ctx context.Context) ContainerServiceAgentPoolProfileResponseArrayOutput {
	return o
}

func (o ContainerServiceAgentPoolProfileResponseArrayOutput) Index(i pulumi.IntInput) ContainerServiceAgentPoolProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerServiceAgentPoolProfileResponse {
		return vs[0].([]ContainerServiceAgentPoolProfileResponse)[vs[1].(int)]
	}).(ContainerServiceAgentPoolProfileResponseOutput)
}

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

type ContainerServiceServicePrincipalProfile struct {
	ClientId          string             `pulumi:"clientId"`
	KeyVaultSecretRef *KeyVaultSecretRef `pulumi:"keyVaultSecretRef"`
	Secret            *string            `pulumi:"secret"`
}





type ContainerServiceServicePrincipalProfileInput interface {
	pulumi.Input

	ToContainerServiceServicePrincipalProfileOutput() ContainerServiceServicePrincipalProfileOutput
	ToContainerServiceServicePrincipalProfileOutputWithContext(context.Context) ContainerServiceServicePrincipalProfileOutput
}

type ContainerServiceServicePrincipalProfileArgs struct {
	ClientId          pulumi.StringInput        `pulumi:"clientId"`
	KeyVaultSecretRef KeyVaultSecretRefPtrInput `pulumi:"keyVaultSecretRef"`
	Secret            pulumi.StringPtrInput     `pulumi:"secret"`
}

func (ContainerServiceServicePrincipalProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceServicePrincipalProfile)(nil)).Elem()
}

func (i ContainerServiceServicePrincipalProfileArgs) ToContainerServiceServicePrincipalProfileOutput() ContainerServiceServicePrincipalProfileOutput {
	return i.ToContainerServiceServicePrincipalProfileOutputWithContext(context.Background())
}

func (i ContainerServiceServicePrincipalProfileArgs) ToContainerServiceServicePrincipalProfileOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceServicePrincipalProfileOutput)
}

func (i ContainerServiceServicePrincipalProfileArgs) ToContainerServiceServicePrincipalProfilePtrOutput() ContainerServiceServicePrincipalProfilePtrOutput {
	return i.ToContainerServiceServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceServicePrincipalProfileArgs) ToContainerServiceServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceServicePrincipalProfileOutput).ToContainerServiceServicePrincipalProfilePtrOutputWithContext(ctx)
}









type ContainerServiceServicePrincipalProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceServicePrincipalProfilePtrOutput() ContainerServiceServicePrincipalProfilePtrOutput
	ToContainerServiceServicePrincipalProfilePtrOutputWithContext(context.Context) ContainerServiceServicePrincipalProfilePtrOutput
}

type containerServiceServicePrincipalProfilePtrType ContainerServiceServicePrincipalProfileArgs

func ContainerServiceServicePrincipalProfilePtr(v *ContainerServiceServicePrincipalProfileArgs) ContainerServiceServicePrincipalProfilePtrInput {
	return (*containerServiceServicePrincipalProfilePtrType)(v)
}

func (*containerServiceServicePrincipalProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceServicePrincipalProfile)(nil)).Elem()
}

func (i *containerServiceServicePrincipalProfilePtrType) ToContainerServiceServicePrincipalProfilePtrOutput() ContainerServiceServicePrincipalProfilePtrOutput {
	return i.ToContainerServiceServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceServicePrincipalProfilePtrType) ToContainerServiceServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceServicePrincipalProfilePtrOutput)
}

type ContainerServiceServicePrincipalProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceServicePrincipalProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceServicePrincipalProfile)(nil)).Elem()
}

func (o ContainerServiceServicePrincipalProfileOutput) ToContainerServiceServicePrincipalProfileOutput() ContainerServiceServicePrincipalProfileOutput {
	return o
}

func (o ContainerServiceServicePrincipalProfileOutput) ToContainerServiceServicePrincipalProfileOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfileOutput {
	return o
}

func (o ContainerServiceServicePrincipalProfileOutput) ToContainerServiceServicePrincipalProfilePtrOutput() ContainerServiceServicePrincipalProfilePtrOutput {
	return o.ToContainerServiceServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceServicePrincipalProfileOutput) ToContainerServiceServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceServicePrincipalProfile) *ContainerServiceServicePrincipalProfile {
		return &v
	}).(ContainerServiceServicePrincipalProfilePtrOutput)
}

func (o ContainerServiceServicePrincipalProfileOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceServicePrincipalProfile) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ContainerServiceServicePrincipalProfileOutput) KeyVaultSecretRef() KeyVaultSecretRefPtrOutput {
	return o.ApplyT(func(v ContainerServiceServicePrincipalProfile) *KeyVaultSecretRef { return v.KeyVaultSecretRef }).(KeyVaultSecretRefPtrOutput)
}

func (o ContainerServiceServicePrincipalProfileOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceServicePrincipalProfile) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ContainerServiceServicePrincipalProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceServicePrincipalProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceServicePrincipalProfile)(nil)).Elem()
}

func (o ContainerServiceServicePrincipalProfilePtrOutput) ToContainerServiceServicePrincipalProfilePtrOutput() ContainerServiceServicePrincipalProfilePtrOutput {
	return o
}

func (o ContainerServiceServicePrincipalProfilePtrOutput) ToContainerServiceServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfilePtrOutput {
	return o
}

func (o ContainerServiceServicePrincipalProfilePtrOutput) Elem() ContainerServiceServicePrincipalProfileOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfile) ContainerServiceServicePrincipalProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceServicePrincipalProfile
		return ret
	}).(ContainerServiceServicePrincipalProfileOutput)
}

func (o ContainerServiceServicePrincipalProfilePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceServicePrincipalProfilePtrOutput) KeyVaultSecretRef() KeyVaultSecretRefPtrOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfile) *KeyVaultSecretRef {
		if v == nil {
			return nil
		}
		return v.KeyVaultSecretRef
	}).(KeyVaultSecretRefPtrOutput)
}

func (o ContainerServiceServicePrincipalProfilePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceServicePrincipalProfileResponse struct {
	ClientId          string                     `pulumi:"clientId"`
	KeyVaultSecretRef *KeyVaultSecretRefResponse `pulumi:"keyVaultSecretRef"`
	Secret            *string                    `pulumi:"secret"`
}





type ContainerServiceServicePrincipalProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceServicePrincipalProfileResponseOutput() ContainerServiceServicePrincipalProfileResponseOutput
	ToContainerServiceServicePrincipalProfileResponseOutputWithContext(context.Context) ContainerServiceServicePrincipalProfileResponseOutput
}

type ContainerServiceServicePrincipalProfileResponseArgs struct {
	ClientId          pulumi.StringInput                `pulumi:"clientId"`
	KeyVaultSecretRef KeyVaultSecretRefResponsePtrInput `pulumi:"keyVaultSecretRef"`
	Secret            pulumi.StringPtrInput             `pulumi:"secret"`
}

func (ContainerServiceServicePrincipalProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceServicePrincipalProfileResponse)(nil)).Elem()
}

func (i ContainerServiceServicePrincipalProfileResponseArgs) ToContainerServiceServicePrincipalProfileResponseOutput() ContainerServiceServicePrincipalProfileResponseOutput {
	return i.ToContainerServiceServicePrincipalProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceServicePrincipalProfileResponseArgs) ToContainerServiceServicePrincipalProfileResponseOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceServicePrincipalProfileResponseOutput)
}

func (i ContainerServiceServicePrincipalProfileResponseArgs) ToContainerServiceServicePrincipalProfileResponsePtrOutput() ContainerServiceServicePrincipalProfileResponsePtrOutput {
	return i.ToContainerServiceServicePrincipalProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceServicePrincipalProfileResponseArgs) ToContainerServiceServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceServicePrincipalProfileResponseOutput).ToContainerServiceServicePrincipalProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceServicePrincipalProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceServicePrincipalProfileResponsePtrOutput() ContainerServiceServicePrincipalProfileResponsePtrOutput
	ToContainerServiceServicePrincipalProfileResponsePtrOutputWithContext(context.Context) ContainerServiceServicePrincipalProfileResponsePtrOutput
}

type containerServiceServicePrincipalProfileResponsePtrType ContainerServiceServicePrincipalProfileResponseArgs

func ContainerServiceServicePrincipalProfileResponsePtr(v *ContainerServiceServicePrincipalProfileResponseArgs) ContainerServiceServicePrincipalProfileResponsePtrInput {
	return (*containerServiceServicePrincipalProfileResponsePtrType)(v)
}

func (*containerServiceServicePrincipalProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceServicePrincipalProfileResponse)(nil)).Elem()
}

func (i *containerServiceServicePrincipalProfileResponsePtrType) ToContainerServiceServicePrincipalProfileResponsePtrOutput() ContainerServiceServicePrincipalProfileResponsePtrOutput {
	return i.ToContainerServiceServicePrincipalProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceServicePrincipalProfileResponsePtrType) ToContainerServiceServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceServicePrincipalProfileResponsePtrOutput)
}

type ContainerServiceServicePrincipalProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceServicePrincipalProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ContainerServiceServicePrincipalProfileResponseOutput) ToContainerServiceServicePrincipalProfileResponseOutput() ContainerServiceServicePrincipalProfileResponseOutput {
	return o
}

func (o ContainerServiceServicePrincipalProfileResponseOutput) ToContainerServiceServicePrincipalProfileResponseOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfileResponseOutput {
	return o
}

func (o ContainerServiceServicePrincipalProfileResponseOutput) ToContainerServiceServicePrincipalProfileResponsePtrOutput() ContainerServiceServicePrincipalProfileResponsePtrOutput {
	return o.ToContainerServiceServicePrincipalProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceServicePrincipalProfileResponseOutput) ToContainerServiceServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceServicePrincipalProfileResponse) *ContainerServiceServicePrincipalProfileResponse {
		return &v
	}).(ContainerServiceServicePrincipalProfileResponsePtrOutput)
}

func (o ContainerServiceServicePrincipalProfileResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceServicePrincipalProfileResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ContainerServiceServicePrincipalProfileResponseOutput) KeyVaultSecretRef() KeyVaultSecretRefResponsePtrOutput {
	return o.ApplyT(func(v ContainerServiceServicePrincipalProfileResponse) *KeyVaultSecretRefResponse {
		return v.KeyVaultSecretRef
	}).(KeyVaultSecretRefResponsePtrOutput)
}

func (o ContainerServiceServicePrincipalProfileResponseOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceServicePrincipalProfileResponse) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ContainerServiceServicePrincipalProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceServicePrincipalProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ContainerServiceServicePrincipalProfileResponsePtrOutput) ToContainerServiceServicePrincipalProfileResponsePtrOutput() ContainerServiceServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceServicePrincipalProfileResponsePtrOutput) ToContainerServiceServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceServicePrincipalProfileResponsePtrOutput) Elem() ContainerServiceServicePrincipalProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfileResponse) ContainerServiceServicePrincipalProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceServicePrincipalProfileResponse
		return ret
	}).(ContainerServiceServicePrincipalProfileResponseOutput)
}

func (o ContainerServiceServicePrincipalProfileResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceServicePrincipalProfileResponsePtrOutput) KeyVaultSecretRef() KeyVaultSecretRefResponsePtrOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfileResponse) *KeyVaultSecretRefResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultSecretRef
	}).(KeyVaultSecretRefResponsePtrOutput)
}

func (o ContainerServiceServicePrincipalProfileResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Secret
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

type KeyVaultSecretRef struct {
	SecretName string  `pulumi:"secretName"`
	VaultID    string  `pulumi:"vaultID"`
	Version    *string `pulumi:"version"`
}





type KeyVaultSecretRefInput interface {
	pulumi.Input

	ToKeyVaultSecretRefOutput() KeyVaultSecretRefOutput
	ToKeyVaultSecretRefOutputWithContext(context.Context) KeyVaultSecretRefOutput
}

type KeyVaultSecretRefArgs struct {
	SecretName pulumi.StringInput    `pulumi:"secretName"`
	VaultID    pulumi.StringInput    `pulumi:"vaultID"`
	Version    pulumi.StringPtrInput `pulumi:"version"`
}

func (KeyVaultSecretRefArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretRef)(nil)).Elem()
}

func (i KeyVaultSecretRefArgs) ToKeyVaultSecretRefOutput() KeyVaultSecretRefOutput {
	return i.ToKeyVaultSecretRefOutputWithContext(context.Background())
}

func (i KeyVaultSecretRefArgs) ToKeyVaultSecretRefOutputWithContext(ctx context.Context) KeyVaultSecretRefOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretRefOutput)
}

func (i KeyVaultSecretRefArgs) ToKeyVaultSecretRefPtrOutput() KeyVaultSecretRefPtrOutput {
	return i.ToKeyVaultSecretRefPtrOutputWithContext(context.Background())
}

func (i KeyVaultSecretRefArgs) ToKeyVaultSecretRefPtrOutputWithContext(ctx context.Context) KeyVaultSecretRefPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretRefOutput).ToKeyVaultSecretRefPtrOutputWithContext(ctx)
}









type KeyVaultSecretRefPtrInput interface {
	pulumi.Input

	ToKeyVaultSecretRefPtrOutput() KeyVaultSecretRefPtrOutput
	ToKeyVaultSecretRefPtrOutputWithContext(context.Context) KeyVaultSecretRefPtrOutput
}

type keyVaultSecretRefPtrType KeyVaultSecretRefArgs

func KeyVaultSecretRefPtr(v *KeyVaultSecretRefArgs) KeyVaultSecretRefPtrInput {
	return (*keyVaultSecretRefPtrType)(v)
}

func (*keyVaultSecretRefPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretRef)(nil)).Elem()
}

func (i *keyVaultSecretRefPtrType) ToKeyVaultSecretRefPtrOutput() KeyVaultSecretRefPtrOutput {
	return i.ToKeyVaultSecretRefPtrOutputWithContext(context.Background())
}

func (i *keyVaultSecretRefPtrType) ToKeyVaultSecretRefPtrOutputWithContext(ctx context.Context) KeyVaultSecretRefPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretRefPtrOutput)
}

type KeyVaultSecretRefOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretRefOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretRef)(nil)).Elem()
}

func (o KeyVaultSecretRefOutput) ToKeyVaultSecretRefOutput() KeyVaultSecretRefOutput {
	return o
}

func (o KeyVaultSecretRefOutput) ToKeyVaultSecretRefOutputWithContext(ctx context.Context) KeyVaultSecretRefOutput {
	return o
}

func (o KeyVaultSecretRefOutput) ToKeyVaultSecretRefPtrOutput() KeyVaultSecretRefPtrOutput {
	return o.ToKeyVaultSecretRefPtrOutputWithContext(context.Background())
}

func (o KeyVaultSecretRefOutput) ToKeyVaultSecretRefPtrOutputWithContext(ctx context.Context) KeyVaultSecretRefPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultSecretRef) *KeyVaultSecretRef {
		return &v
	}).(KeyVaultSecretRefPtrOutput)
}

func (o KeyVaultSecretRefOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretRef) string { return v.SecretName }).(pulumi.StringOutput)
}

func (o KeyVaultSecretRefOutput) VaultID() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretRef) string { return v.VaultID }).(pulumi.StringOutput)
}

func (o KeyVaultSecretRefOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultSecretRef) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type KeyVaultSecretRefPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretRefPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretRef)(nil)).Elem()
}

func (o KeyVaultSecretRefPtrOutput) ToKeyVaultSecretRefPtrOutput() KeyVaultSecretRefPtrOutput {
	return o
}

func (o KeyVaultSecretRefPtrOutput) ToKeyVaultSecretRefPtrOutputWithContext(ctx context.Context) KeyVaultSecretRefPtrOutput {
	return o
}

func (o KeyVaultSecretRefPtrOutput) Elem() KeyVaultSecretRefOutput {
	return o.ApplyT(func(v *KeyVaultSecretRef) KeyVaultSecretRef {
		if v != nil {
			return *v
		}
		var ret KeyVaultSecretRef
		return ret
	}).(KeyVaultSecretRefOutput)
}

func (o KeyVaultSecretRefPtrOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretRef) *string {
		if v == nil {
			return nil
		}
		return &v.SecretName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretRefPtrOutput) VaultID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretRef) *string {
		if v == nil {
			return nil
		}
		return &v.VaultID
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretRefPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretRef) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type KeyVaultSecretRefResponse struct {
	SecretName string  `pulumi:"secretName"`
	VaultID    string  `pulumi:"vaultID"`
	Version    *string `pulumi:"version"`
}





type KeyVaultSecretRefResponseInput interface {
	pulumi.Input

	ToKeyVaultSecretRefResponseOutput() KeyVaultSecretRefResponseOutput
	ToKeyVaultSecretRefResponseOutputWithContext(context.Context) KeyVaultSecretRefResponseOutput
}

type KeyVaultSecretRefResponseArgs struct {
	SecretName pulumi.StringInput    `pulumi:"secretName"`
	VaultID    pulumi.StringInput    `pulumi:"vaultID"`
	Version    pulumi.StringPtrInput `pulumi:"version"`
}

func (KeyVaultSecretRefResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretRefResponse)(nil)).Elem()
}

func (i KeyVaultSecretRefResponseArgs) ToKeyVaultSecretRefResponseOutput() KeyVaultSecretRefResponseOutput {
	return i.ToKeyVaultSecretRefResponseOutputWithContext(context.Background())
}

func (i KeyVaultSecretRefResponseArgs) ToKeyVaultSecretRefResponseOutputWithContext(ctx context.Context) KeyVaultSecretRefResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretRefResponseOutput)
}

func (i KeyVaultSecretRefResponseArgs) ToKeyVaultSecretRefResponsePtrOutput() KeyVaultSecretRefResponsePtrOutput {
	return i.ToKeyVaultSecretRefResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultSecretRefResponseArgs) ToKeyVaultSecretRefResponsePtrOutputWithContext(ctx context.Context) KeyVaultSecretRefResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretRefResponseOutput).ToKeyVaultSecretRefResponsePtrOutputWithContext(ctx)
}









type KeyVaultSecretRefResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultSecretRefResponsePtrOutput() KeyVaultSecretRefResponsePtrOutput
	ToKeyVaultSecretRefResponsePtrOutputWithContext(context.Context) KeyVaultSecretRefResponsePtrOutput
}

type keyVaultSecretRefResponsePtrType KeyVaultSecretRefResponseArgs

func KeyVaultSecretRefResponsePtr(v *KeyVaultSecretRefResponseArgs) KeyVaultSecretRefResponsePtrInput {
	return (*keyVaultSecretRefResponsePtrType)(v)
}

func (*keyVaultSecretRefResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretRefResponse)(nil)).Elem()
}

func (i *keyVaultSecretRefResponsePtrType) ToKeyVaultSecretRefResponsePtrOutput() KeyVaultSecretRefResponsePtrOutput {
	return i.ToKeyVaultSecretRefResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultSecretRefResponsePtrType) ToKeyVaultSecretRefResponsePtrOutputWithContext(ctx context.Context) KeyVaultSecretRefResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretRefResponsePtrOutput)
}

type KeyVaultSecretRefResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretRefResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretRefResponse)(nil)).Elem()
}

func (o KeyVaultSecretRefResponseOutput) ToKeyVaultSecretRefResponseOutput() KeyVaultSecretRefResponseOutput {
	return o
}

func (o KeyVaultSecretRefResponseOutput) ToKeyVaultSecretRefResponseOutputWithContext(ctx context.Context) KeyVaultSecretRefResponseOutput {
	return o
}

func (o KeyVaultSecretRefResponseOutput) ToKeyVaultSecretRefResponsePtrOutput() KeyVaultSecretRefResponsePtrOutput {
	return o.ToKeyVaultSecretRefResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultSecretRefResponseOutput) ToKeyVaultSecretRefResponsePtrOutputWithContext(ctx context.Context) KeyVaultSecretRefResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultSecretRefResponse) *KeyVaultSecretRefResponse {
		return &v
	}).(KeyVaultSecretRefResponsePtrOutput)
}

func (o KeyVaultSecretRefResponseOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretRefResponse) string { return v.SecretName }).(pulumi.StringOutput)
}

func (o KeyVaultSecretRefResponseOutput) VaultID() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretRefResponse) string { return v.VaultID }).(pulumi.StringOutput)
}

func (o KeyVaultSecretRefResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultSecretRefResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type KeyVaultSecretRefResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretRefResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretRefResponse)(nil)).Elem()
}

func (o KeyVaultSecretRefResponsePtrOutput) ToKeyVaultSecretRefResponsePtrOutput() KeyVaultSecretRefResponsePtrOutput {
	return o
}

func (o KeyVaultSecretRefResponsePtrOutput) ToKeyVaultSecretRefResponsePtrOutputWithContext(ctx context.Context) KeyVaultSecretRefResponsePtrOutput {
	return o
}

func (o KeyVaultSecretRefResponsePtrOutput) Elem() KeyVaultSecretRefResponseOutput {
	return o.ApplyT(func(v *KeyVaultSecretRefResponse) KeyVaultSecretRefResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultSecretRefResponse
		return ret
	}).(KeyVaultSecretRefResponseOutput)
}

func (o KeyVaultSecretRefResponsePtrOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretRefResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecretName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretRefResponsePtrOutput) VaultID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretRefResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VaultID
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretRefResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretRefResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerServiceAgentPoolProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceAgentPoolProfileArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceAgentPoolProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceAgentPoolProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceServicePrincipalProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceServicePrincipalProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceServicePrincipalProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceServicePrincipalProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretRefOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretRefPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretRefResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretRefResponsePtrOutput{})
}
