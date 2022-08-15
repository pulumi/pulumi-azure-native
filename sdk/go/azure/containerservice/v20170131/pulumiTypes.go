


package v20170131

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerServiceAgentPoolProfile struct {
	Count     int    `pulumi:"count"`
	DnsPrefix string `pulumi:"dnsPrefix"`
	Name      string `pulumi:"name"`
	VmSize    string `pulumi:"vmSize"`
}


func (val *ContainerServiceAgentPoolProfile) Defaults() *ContainerServiceAgentPoolProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = 1
	}
	return &tmp
}





type ContainerServiceAgentPoolProfileInput interface {
	pulumi.Input

	ToContainerServiceAgentPoolProfileOutput() ContainerServiceAgentPoolProfileOutput
	ToContainerServiceAgentPoolProfileOutputWithContext(context.Context) ContainerServiceAgentPoolProfileOutput
}

type ContainerServiceAgentPoolProfileArgs struct {
	Count     pulumi.IntInput    `pulumi:"count"`
	DnsPrefix pulumi.StringInput `pulumi:"dnsPrefix"`
	Name      pulumi.StringInput `pulumi:"name"`
	VmSize    pulumi.StringInput `pulumi:"vmSize"`
}


func (val *ContainerServiceAgentPoolProfileArgs) Defaults() *ContainerServiceAgentPoolProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = pulumi.Int(1)
	}
	return &tmp
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

func (o ContainerServiceAgentPoolProfileOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) int { return v.Count }).(pulumi.IntOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) DnsPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) string { return v.DnsPrefix }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfile) string { return v.VmSize }).(pulumi.StringOutput)
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
	Count     int    `pulumi:"count"`
	DnsPrefix string `pulumi:"dnsPrefix"`
	Fqdn      string `pulumi:"fqdn"`
	Name      string `pulumi:"name"`
	VmSize    string `pulumi:"vmSize"`
}


func (val *ContainerServiceAgentPoolProfileResponse) Defaults() *ContainerServiceAgentPoolProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = 1
	}
	return &tmp
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

func (o ContainerServiceAgentPoolProfileResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) DnsPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) string { return v.DnsPrefix }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerServiceAgentPoolProfileResponseOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceAgentPoolProfileResponse) string { return v.VmSize }).(pulumi.StringOutput)
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

type ContainerServiceCustomProfile struct {
	Orchestrator string `pulumi:"orchestrator"`
}





type ContainerServiceCustomProfileInput interface {
	pulumi.Input

	ToContainerServiceCustomProfileOutput() ContainerServiceCustomProfileOutput
	ToContainerServiceCustomProfileOutputWithContext(context.Context) ContainerServiceCustomProfileOutput
}

type ContainerServiceCustomProfileArgs struct {
	Orchestrator pulumi.StringInput `pulumi:"orchestrator"`
}

func (ContainerServiceCustomProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceCustomProfile)(nil)).Elem()
}

func (i ContainerServiceCustomProfileArgs) ToContainerServiceCustomProfileOutput() ContainerServiceCustomProfileOutput {
	return i.ToContainerServiceCustomProfileOutputWithContext(context.Background())
}

func (i ContainerServiceCustomProfileArgs) ToContainerServiceCustomProfileOutputWithContext(ctx context.Context) ContainerServiceCustomProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceCustomProfileOutput)
}

func (i ContainerServiceCustomProfileArgs) ToContainerServiceCustomProfilePtrOutput() ContainerServiceCustomProfilePtrOutput {
	return i.ToContainerServiceCustomProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceCustomProfileArgs) ToContainerServiceCustomProfilePtrOutputWithContext(ctx context.Context) ContainerServiceCustomProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceCustomProfileOutput).ToContainerServiceCustomProfilePtrOutputWithContext(ctx)
}









type ContainerServiceCustomProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceCustomProfilePtrOutput() ContainerServiceCustomProfilePtrOutput
	ToContainerServiceCustomProfilePtrOutputWithContext(context.Context) ContainerServiceCustomProfilePtrOutput
}

type containerServiceCustomProfilePtrType ContainerServiceCustomProfileArgs

func ContainerServiceCustomProfilePtr(v *ContainerServiceCustomProfileArgs) ContainerServiceCustomProfilePtrInput {
	return (*containerServiceCustomProfilePtrType)(v)
}

func (*containerServiceCustomProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceCustomProfile)(nil)).Elem()
}

func (i *containerServiceCustomProfilePtrType) ToContainerServiceCustomProfilePtrOutput() ContainerServiceCustomProfilePtrOutput {
	return i.ToContainerServiceCustomProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceCustomProfilePtrType) ToContainerServiceCustomProfilePtrOutputWithContext(ctx context.Context) ContainerServiceCustomProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceCustomProfilePtrOutput)
}

type ContainerServiceCustomProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceCustomProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceCustomProfile)(nil)).Elem()
}

func (o ContainerServiceCustomProfileOutput) ToContainerServiceCustomProfileOutput() ContainerServiceCustomProfileOutput {
	return o
}

func (o ContainerServiceCustomProfileOutput) ToContainerServiceCustomProfileOutputWithContext(ctx context.Context) ContainerServiceCustomProfileOutput {
	return o
}

func (o ContainerServiceCustomProfileOutput) ToContainerServiceCustomProfilePtrOutput() ContainerServiceCustomProfilePtrOutput {
	return o.ToContainerServiceCustomProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceCustomProfileOutput) ToContainerServiceCustomProfilePtrOutputWithContext(ctx context.Context) ContainerServiceCustomProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceCustomProfile) *ContainerServiceCustomProfile {
		return &v
	}).(ContainerServiceCustomProfilePtrOutput)
}

func (o ContainerServiceCustomProfileOutput) Orchestrator() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceCustomProfile) string { return v.Orchestrator }).(pulumi.StringOutput)
}

type ContainerServiceCustomProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceCustomProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceCustomProfile)(nil)).Elem()
}

func (o ContainerServiceCustomProfilePtrOutput) ToContainerServiceCustomProfilePtrOutput() ContainerServiceCustomProfilePtrOutput {
	return o
}

func (o ContainerServiceCustomProfilePtrOutput) ToContainerServiceCustomProfilePtrOutputWithContext(ctx context.Context) ContainerServiceCustomProfilePtrOutput {
	return o
}

func (o ContainerServiceCustomProfilePtrOutput) Elem() ContainerServiceCustomProfileOutput {
	return o.ApplyT(func(v *ContainerServiceCustomProfile) ContainerServiceCustomProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceCustomProfile
		return ret
	}).(ContainerServiceCustomProfileOutput)
}

func (o ContainerServiceCustomProfilePtrOutput) Orchestrator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceCustomProfile) *string {
		if v == nil {
			return nil
		}
		return &v.Orchestrator
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceCustomProfileResponse struct {
	Orchestrator string `pulumi:"orchestrator"`
}

type ContainerServiceCustomProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceCustomProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceCustomProfileResponse)(nil)).Elem()
}

func (o ContainerServiceCustomProfileResponseOutput) ToContainerServiceCustomProfileResponseOutput() ContainerServiceCustomProfileResponseOutput {
	return o
}

func (o ContainerServiceCustomProfileResponseOutput) ToContainerServiceCustomProfileResponseOutputWithContext(ctx context.Context) ContainerServiceCustomProfileResponseOutput {
	return o
}

func (o ContainerServiceCustomProfileResponseOutput) Orchestrator() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceCustomProfileResponse) string { return v.Orchestrator }).(pulumi.StringOutput)
}

type ContainerServiceCustomProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceCustomProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceCustomProfileResponse)(nil)).Elem()
}

func (o ContainerServiceCustomProfileResponsePtrOutput) ToContainerServiceCustomProfileResponsePtrOutput() ContainerServiceCustomProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceCustomProfileResponsePtrOutput) ToContainerServiceCustomProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceCustomProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceCustomProfileResponsePtrOutput) Elem() ContainerServiceCustomProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceCustomProfileResponse) ContainerServiceCustomProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceCustomProfileResponse
		return ret
	}).(ContainerServiceCustomProfileResponseOutput)
}

func (o ContainerServiceCustomProfileResponsePtrOutput) Orchestrator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceCustomProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Orchestrator
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceDiagnosticsProfile struct {
	VmDiagnostics ContainerServiceVMDiagnostics `pulumi:"vmDiagnostics"`
}





type ContainerServiceDiagnosticsProfileInput interface {
	pulumi.Input

	ToContainerServiceDiagnosticsProfileOutput() ContainerServiceDiagnosticsProfileOutput
	ToContainerServiceDiagnosticsProfileOutputWithContext(context.Context) ContainerServiceDiagnosticsProfileOutput
}

type ContainerServiceDiagnosticsProfileArgs struct {
	VmDiagnostics ContainerServiceVMDiagnosticsInput `pulumi:"vmDiagnostics"`
}

func (ContainerServiceDiagnosticsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceDiagnosticsProfile)(nil)).Elem()
}

func (i ContainerServiceDiagnosticsProfileArgs) ToContainerServiceDiagnosticsProfileOutput() ContainerServiceDiagnosticsProfileOutput {
	return i.ToContainerServiceDiagnosticsProfileOutputWithContext(context.Background())
}

func (i ContainerServiceDiagnosticsProfileArgs) ToContainerServiceDiagnosticsProfileOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceDiagnosticsProfileOutput)
}

func (i ContainerServiceDiagnosticsProfileArgs) ToContainerServiceDiagnosticsProfilePtrOutput() ContainerServiceDiagnosticsProfilePtrOutput {
	return i.ToContainerServiceDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceDiagnosticsProfileArgs) ToContainerServiceDiagnosticsProfilePtrOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceDiagnosticsProfileOutput).ToContainerServiceDiagnosticsProfilePtrOutputWithContext(ctx)
}









type ContainerServiceDiagnosticsProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceDiagnosticsProfilePtrOutput() ContainerServiceDiagnosticsProfilePtrOutput
	ToContainerServiceDiagnosticsProfilePtrOutputWithContext(context.Context) ContainerServiceDiagnosticsProfilePtrOutput
}

type containerServiceDiagnosticsProfilePtrType ContainerServiceDiagnosticsProfileArgs

func ContainerServiceDiagnosticsProfilePtr(v *ContainerServiceDiagnosticsProfileArgs) ContainerServiceDiagnosticsProfilePtrInput {
	return (*containerServiceDiagnosticsProfilePtrType)(v)
}

func (*containerServiceDiagnosticsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceDiagnosticsProfile)(nil)).Elem()
}

func (i *containerServiceDiagnosticsProfilePtrType) ToContainerServiceDiagnosticsProfilePtrOutput() ContainerServiceDiagnosticsProfilePtrOutput {
	return i.ToContainerServiceDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceDiagnosticsProfilePtrType) ToContainerServiceDiagnosticsProfilePtrOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceDiagnosticsProfilePtrOutput)
}

type ContainerServiceDiagnosticsProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceDiagnosticsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceDiagnosticsProfile)(nil)).Elem()
}

func (o ContainerServiceDiagnosticsProfileOutput) ToContainerServiceDiagnosticsProfileOutput() ContainerServiceDiagnosticsProfileOutput {
	return o
}

func (o ContainerServiceDiagnosticsProfileOutput) ToContainerServiceDiagnosticsProfileOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfileOutput {
	return o
}

func (o ContainerServiceDiagnosticsProfileOutput) ToContainerServiceDiagnosticsProfilePtrOutput() ContainerServiceDiagnosticsProfilePtrOutput {
	return o.ToContainerServiceDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceDiagnosticsProfileOutput) ToContainerServiceDiagnosticsProfilePtrOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceDiagnosticsProfile) *ContainerServiceDiagnosticsProfile {
		return &v
	}).(ContainerServiceDiagnosticsProfilePtrOutput)
}

func (o ContainerServiceDiagnosticsProfileOutput) VmDiagnostics() ContainerServiceVMDiagnosticsOutput {
	return o.ApplyT(func(v ContainerServiceDiagnosticsProfile) ContainerServiceVMDiagnostics { return v.VmDiagnostics }).(ContainerServiceVMDiagnosticsOutput)
}

type ContainerServiceDiagnosticsProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceDiagnosticsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceDiagnosticsProfile)(nil)).Elem()
}

func (o ContainerServiceDiagnosticsProfilePtrOutput) ToContainerServiceDiagnosticsProfilePtrOutput() ContainerServiceDiagnosticsProfilePtrOutput {
	return o
}

func (o ContainerServiceDiagnosticsProfilePtrOutput) ToContainerServiceDiagnosticsProfilePtrOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfilePtrOutput {
	return o
}

func (o ContainerServiceDiagnosticsProfilePtrOutput) Elem() ContainerServiceDiagnosticsProfileOutput {
	return o.ApplyT(func(v *ContainerServiceDiagnosticsProfile) ContainerServiceDiagnosticsProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceDiagnosticsProfile
		return ret
	}).(ContainerServiceDiagnosticsProfileOutput)
}

func (o ContainerServiceDiagnosticsProfilePtrOutput) VmDiagnostics() ContainerServiceVMDiagnosticsPtrOutput {
	return o.ApplyT(func(v *ContainerServiceDiagnosticsProfile) *ContainerServiceVMDiagnostics {
		if v == nil {
			return nil
		}
		return &v.VmDiagnostics
	}).(ContainerServiceVMDiagnosticsPtrOutput)
}

type ContainerServiceDiagnosticsProfileResponse struct {
	VmDiagnostics ContainerServiceVMDiagnosticsResponse `pulumi:"vmDiagnostics"`
}

type ContainerServiceDiagnosticsProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceDiagnosticsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceDiagnosticsProfileResponse)(nil)).Elem()
}

func (o ContainerServiceDiagnosticsProfileResponseOutput) ToContainerServiceDiagnosticsProfileResponseOutput() ContainerServiceDiagnosticsProfileResponseOutput {
	return o
}

func (o ContainerServiceDiagnosticsProfileResponseOutput) ToContainerServiceDiagnosticsProfileResponseOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfileResponseOutput {
	return o
}

func (o ContainerServiceDiagnosticsProfileResponseOutput) VmDiagnostics() ContainerServiceVMDiagnosticsResponseOutput {
	return o.ApplyT(func(v ContainerServiceDiagnosticsProfileResponse) ContainerServiceVMDiagnosticsResponse {
		return v.VmDiagnostics
	}).(ContainerServiceVMDiagnosticsResponseOutput)
}

type ContainerServiceDiagnosticsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceDiagnosticsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceDiagnosticsProfileResponse)(nil)).Elem()
}

func (o ContainerServiceDiagnosticsProfileResponsePtrOutput) ToContainerServiceDiagnosticsProfileResponsePtrOutput() ContainerServiceDiagnosticsProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceDiagnosticsProfileResponsePtrOutput) ToContainerServiceDiagnosticsProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceDiagnosticsProfileResponsePtrOutput) Elem() ContainerServiceDiagnosticsProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceDiagnosticsProfileResponse) ContainerServiceDiagnosticsProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceDiagnosticsProfileResponse
		return ret
	}).(ContainerServiceDiagnosticsProfileResponseOutput)
}

func (o ContainerServiceDiagnosticsProfileResponsePtrOutput) VmDiagnostics() ContainerServiceVMDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v *ContainerServiceDiagnosticsProfileResponse) *ContainerServiceVMDiagnosticsResponse {
		if v == nil {
			return nil
		}
		return &v.VmDiagnostics
	}).(ContainerServiceVMDiagnosticsResponsePtrOutput)
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

func (o ContainerServiceLinuxProfileOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfile) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ContainerServiceLinuxProfileOutput) Ssh() ContainerServiceSshConfigurationOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfile) ContainerServiceSshConfiguration { return v.Ssh }).(ContainerServiceSshConfigurationOutput)
}

type ContainerServiceLinuxProfileResponse struct {
	AdminUsername string                                   `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfigurationResponse `pulumi:"ssh"`
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

func (o ContainerServiceLinuxProfileResponseOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfileResponse) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ContainerServiceLinuxProfileResponseOutput) Ssh() ContainerServiceSshConfigurationResponseOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfileResponse) ContainerServiceSshConfigurationResponse { return v.Ssh }).(ContainerServiceSshConfigurationResponseOutput)
}

type ContainerServiceMasterProfile struct {
	Count     *int   `pulumi:"count"`
	DnsPrefix string `pulumi:"dnsPrefix"`
}


func (val *ContainerServiceMasterProfile) Defaults() *ContainerServiceMasterProfile {
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





type ContainerServiceMasterProfileInput interface {
	pulumi.Input

	ToContainerServiceMasterProfileOutput() ContainerServiceMasterProfileOutput
	ToContainerServiceMasterProfileOutputWithContext(context.Context) ContainerServiceMasterProfileOutput
}

type ContainerServiceMasterProfileArgs struct {
	Count     pulumi.IntPtrInput `pulumi:"count"`
	DnsPrefix pulumi.StringInput `pulumi:"dnsPrefix"`
}


func (val *ContainerServiceMasterProfileArgs) Defaults() *ContainerServiceMasterProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = pulumi.IntPtr(1)
	}
	return &tmp
}
func (ContainerServiceMasterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceMasterProfile)(nil)).Elem()
}

func (i ContainerServiceMasterProfileArgs) ToContainerServiceMasterProfileOutput() ContainerServiceMasterProfileOutput {
	return i.ToContainerServiceMasterProfileOutputWithContext(context.Background())
}

func (i ContainerServiceMasterProfileArgs) ToContainerServiceMasterProfileOutputWithContext(ctx context.Context) ContainerServiceMasterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceMasterProfileOutput)
}

type ContainerServiceMasterProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceMasterProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceMasterProfile)(nil)).Elem()
}

func (o ContainerServiceMasterProfileOutput) ToContainerServiceMasterProfileOutput() ContainerServiceMasterProfileOutput {
	return o
}

func (o ContainerServiceMasterProfileOutput) ToContainerServiceMasterProfileOutputWithContext(ctx context.Context) ContainerServiceMasterProfileOutput {
	return o
}

func (o ContainerServiceMasterProfileOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerServiceMasterProfile) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ContainerServiceMasterProfileOutput) DnsPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceMasterProfile) string { return v.DnsPrefix }).(pulumi.StringOutput)
}

type ContainerServiceMasterProfileResponse struct {
	Count     *int   `pulumi:"count"`
	DnsPrefix string `pulumi:"dnsPrefix"`
	Fqdn      string `pulumi:"fqdn"`
}


func (val *ContainerServiceMasterProfileResponse) Defaults() *ContainerServiceMasterProfileResponse {
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

type ContainerServiceMasterProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceMasterProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceMasterProfileResponse)(nil)).Elem()
}

func (o ContainerServiceMasterProfileResponseOutput) ToContainerServiceMasterProfileResponseOutput() ContainerServiceMasterProfileResponseOutput {
	return o
}

func (o ContainerServiceMasterProfileResponseOutput) ToContainerServiceMasterProfileResponseOutputWithContext(ctx context.Context) ContainerServiceMasterProfileResponseOutput {
	return o
}

func (o ContainerServiceMasterProfileResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerServiceMasterProfileResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ContainerServiceMasterProfileResponseOutput) DnsPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceMasterProfileResponse) string { return v.DnsPrefix }).(pulumi.StringOutput)
}

func (o ContainerServiceMasterProfileResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceMasterProfileResponse) string { return v.Fqdn }).(pulumi.StringOutput)
}

type ContainerServiceOrchestratorProfile struct {
	OrchestratorType ContainerServiceOrchestratorTypes `pulumi:"orchestratorType"`
}





type ContainerServiceOrchestratorProfileInput interface {
	pulumi.Input

	ToContainerServiceOrchestratorProfileOutput() ContainerServiceOrchestratorProfileOutput
	ToContainerServiceOrchestratorProfileOutputWithContext(context.Context) ContainerServiceOrchestratorProfileOutput
}

type ContainerServiceOrchestratorProfileArgs struct {
	OrchestratorType ContainerServiceOrchestratorTypesInput `pulumi:"orchestratorType"`
}

func (ContainerServiceOrchestratorProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceOrchestratorProfile)(nil)).Elem()
}

func (i ContainerServiceOrchestratorProfileArgs) ToContainerServiceOrchestratorProfileOutput() ContainerServiceOrchestratorProfileOutput {
	return i.ToContainerServiceOrchestratorProfileOutputWithContext(context.Background())
}

func (i ContainerServiceOrchestratorProfileArgs) ToContainerServiceOrchestratorProfileOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceOrchestratorProfileOutput)
}

func (i ContainerServiceOrchestratorProfileArgs) ToContainerServiceOrchestratorProfilePtrOutput() ContainerServiceOrchestratorProfilePtrOutput {
	return i.ToContainerServiceOrchestratorProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceOrchestratorProfileArgs) ToContainerServiceOrchestratorProfilePtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceOrchestratorProfileOutput).ToContainerServiceOrchestratorProfilePtrOutputWithContext(ctx)
}









type ContainerServiceOrchestratorProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceOrchestratorProfilePtrOutput() ContainerServiceOrchestratorProfilePtrOutput
	ToContainerServiceOrchestratorProfilePtrOutputWithContext(context.Context) ContainerServiceOrchestratorProfilePtrOutput
}

type containerServiceOrchestratorProfilePtrType ContainerServiceOrchestratorProfileArgs

func ContainerServiceOrchestratorProfilePtr(v *ContainerServiceOrchestratorProfileArgs) ContainerServiceOrchestratorProfilePtrInput {
	return (*containerServiceOrchestratorProfilePtrType)(v)
}

func (*containerServiceOrchestratorProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceOrchestratorProfile)(nil)).Elem()
}

func (i *containerServiceOrchestratorProfilePtrType) ToContainerServiceOrchestratorProfilePtrOutput() ContainerServiceOrchestratorProfilePtrOutput {
	return i.ToContainerServiceOrchestratorProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceOrchestratorProfilePtrType) ToContainerServiceOrchestratorProfilePtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceOrchestratorProfilePtrOutput)
}

type ContainerServiceOrchestratorProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceOrchestratorProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceOrchestratorProfile)(nil)).Elem()
}

func (o ContainerServiceOrchestratorProfileOutput) ToContainerServiceOrchestratorProfileOutput() ContainerServiceOrchestratorProfileOutput {
	return o
}

func (o ContainerServiceOrchestratorProfileOutput) ToContainerServiceOrchestratorProfileOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfileOutput {
	return o
}

func (o ContainerServiceOrchestratorProfileOutput) ToContainerServiceOrchestratorProfilePtrOutput() ContainerServiceOrchestratorProfilePtrOutput {
	return o.ToContainerServiceOrchestratorProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceOrchestratorProfileOutput) ToContainerServiceOrchestratorProfilePtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceOrchestratorProfile) *ContainerServiceOrchestratorProfile {
		return &v
	}).(ContainerServiceOrchestratorProfilePtrOutput)
}

func (o ContainerServiceOrchestratorProfileOutput) OrchestratorType() ContainerServiceOrchestratorTypesOutput {
	return o.ApplyT(func(v ContainerServiceOrchestratorProfile) ContainerServiceOrchestratorTypes {
		return v.OrchestratorType
	}).(ContainerServiceOrchestratorTypesOutput)
}

type ContainerServiceOrchestratorProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceOrchestratorProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceOrchestratorProfile)(nil)).Elem()
}

func (o ContainerServiceOrchestratorProfilePtrOutput) ToContainerServiceOrchestratorProfilePtrOutput() ContainerServiceOrchestratorProfilePtrOutput {
	return o
}

func (o ContainerServiceOrchestratorProfilePtrOutput) ToContainerServiceOrchestratorProfilePtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfilePtrOutput {
	return o
}

func (o ContainerServiceOrchestratorProfilePtrOutput) Elem() ContainerServiceOrchestratorProfileOutput {
	return o.ApplyT(func(v *ContainerServiceOrchestratorProfile) ContainerServiceOrchestratorProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceOrchestratorProfile
		return ret
	}).(ContainerServiceOrchestratorProfileOutput)
}

func (o ContainerServiceOrchestratorProfilePtrOutput) OrchestratorType() ContainerServiceOrchestratorTypesPtrOutput {
	return o.ApplyT(func(v *ContainerServiceOrchestratorProfile) *ContainerServiceOrchestratorTypes {
		if v == nil {
			return nil
		}
		return &v.OrchestratorType
	}).(ContainerServiceOrchestratorTypesPtrOutput)
}

type ContainerServiceOrchestratorProfileResponse struct {
	OrchestratorType string `pulumi:"orchestratorType"`
}

type ContainerServiceOrchestratorProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceOrchestratorProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceOrchestratorProfileResponse)(nil)).Elem()
}

func (o ContainerServiceOrchestratorProfileResponseOutput) ToContainerServiceOrchestratorProfileResponseOutput() ContainerServiceOrchestratorProfileResponseOutput {
	return o
}

func (o ContainerServiceOrchestratorProfileResponseOutput) ToContainerServiceOrchestratorProfileResponseOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfileResponseOutput {
	return o
}

func (o ContainerServiceOrchestratorProfileResponseOutput) OrchestratorType() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceOrchestratorProfileResponse) string { return v.OrchestratorType }).(pulumi.StringOutput)
}

type ContainerServiceOrchestratorProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceOrchestratorProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceOrchestratorProfileResponse)(nil)).Elem()
}

func (o ContainerServiceOrchestratorProfileResponsePtrOutput) ToContainerServiceOrchestratorProfileResponsePtrOutput() ContainerServiceOrchestratorProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceOrchestratorProfileResponsePtrOutput) ToContainerServiceOrchestratorProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceOrchestratorProfileResponsePtrOutput) Elem() ContainerServiceOrchestratorProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceOrchestratorProfileResponse) ContainerServiceOrchestratorProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceOrchestratorProfileResponse
		return ret
	}).(ContainerServiceOrchestratorProfileResponseOutput)
}

func (o ContainerServiceOrchestratorProfileResponsePtrOutput) OrchestratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceOrchestratorProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OrchestratorType
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceServicePrincipalProfile struct {
	ClientId string `pulumi:"clientId"`
	Secret   string `pulumi:"secret"`
}





type ContainerServiceServicePrincipalProfileInput interface {
	pulumi.Input

	ToContainerServiceServicePrincipalProfileOutput() ContainerServiceServicePrincipalProfileOutput
	ToContainerServiceServicePrincipalProfileOutputWithContext(context.Context) ContainerServiceServicePrincipalProfileOutput
}

type ContainerServiceServicePrincipalProfileArgs struct {
	ClientId pulumi.StringInput `pulumi:"clientId"`
	Secret   pulumi.StringInput `pulumi:"secret"`
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

func (o ContainerServiceServicePrincipalProfileOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceServicePrincipalProfile) string { return v.Secret }).(pulumi.StringOutput)
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

func (o ContainerServiceServicePrincipalProfilePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceServicePrincipalProfileResponse struct {
	ClientId string `pulumi:"clientId"`
	Secret   string `pulumi:"secret"`
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

func (o ContainerServiceServicePrincipalProfileResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceServicePrincipalProfileResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ContainerServiceServicePrincipalProfileResponseOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceServicePrincipalProfileResponse) string { return v.Secret }).(pulumi.StringOutput)
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

func (o ContainerServiceServicePrincipalProfileResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
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

func (o ContainerServiceSshConfigurationOutput) PublicKeys() ContainerServiceSshPublicKeyArrayOutput {
	return o.ApplyT(func(v ContainerServiceSshConfiguration) []ContainerServiceSshPublicKey { return v.PublicKeys }).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshConfigurationResponse struct {
	PublicKeys []ContainerServiceSshPublicKeyResponse `pulumi:"publicKeys"`
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

func (o ContainerServiceSshConfigurationResponseOutput) PublicKeys() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v ContainerServiceSshConfigurationResponse) []ContainerServiceSshPublicKeyResponse {
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

type ContainerServiceVMDiagnostics struct {
	Enabled bool `pulumi:"enabled"`
}





type ContainerServiceVMDiagnosticsInput interface {
	pulumi.Input

	ToContainerServiceVMDiagnosticsOutput() ContainerServiceVMDiagnosticsOutput
	ToContainerServiceVMDiagnosticsOutputWithContext(context.Context) ContainerServiceVMDiagnosticsOutput
}

type ContainerServiceVMDiagnosticsArgs struct {
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (ContainerServiceVMDiagnosticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceVMDiagnostics)(nil)).Elem()
}

func (i ContainerServiceVMDiagnosticsArgs) ToContainerServiceVMDiagnosticsOutput() ContainerServiceVMDiagnosticsOutput {
	return i.ToContainerServiceVMDiagnosticsOutputWithContext(context.Background())
}

func (i ContainerServiceVMDiagnosticsArgs) ToContainerServiceVMDiagnosticsOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceVMDiagnosticsOutput)
}

func (i ContainerServiceVMDiagnosticsArgs) ToContainerServiceVMDiagnosticsPtrOutput() ContainerServiceVMDiagnosticsPtrOutput {
	return i.ToContainerServiceVMDiagnosticsPtrOutputWithContext(context.Background())
}

func (i ContainerServiceVMDiagnosticsArgs) ToContainerServiceVMDiagnosticsPtrOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceVMDiagnosticsOutput).ToContainerServiceVMDiagnosticsPtrOutputWithContext(ctx)
}









type ContainerServiceVMDiagnosticsPtrInput interface {
	pulumi.Input

	ToContainerServiceVMDiagnosticsPtrOutput() ContainerServiceVMDiagnosticsPtrOutput
	ToContainerServiceVMDiagnosticsPtrOutputWithContext(context.Context) ContainerServiceVMDiagnosticsPtrOutput
}

type containerServiceVMDiagnosticsPtrType ContainerServiceVMDiagnosticsArgs

func ContainerServiceVMDiagnosticsPtr(v *ContainerServiceVMDiagnosticsArgs) ContainerServiceVMDiagnosticsPtrInput {
	return (*containerServiceVMDiagnosticsPtrType)(v)
}

func (*containerServiceVMDiagnosticsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceVMDiagnostics)(nil)).Elem()
}

func (i *containerServiceVMDiagnosticsPtrType) ToContainerServiceVMDiagnosticsPtrOutput() ContainerServiceVMDiagnosticsPtrOutput {
	return i.ToContainerServiceVMDiagnosticsPtrOutputWithContext(context.Background())
}

func (i *containerServiceVMDiagnosticsPtrType) ToContainerServiceVMDiagnosticsPtrOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceVMDiagnosticsPtrOutput)
}

type ContainerServiceVMDiagnosticsOutput struct{ *pulumi.OutputState }

func (ContainerServiceVMDiagnosticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceVMDiagnostics)(nil)).Elem()
}

func (o ContainerServiceVMDiagnosticsOutput) ToContainerServiceVMDiagnosticsOutput() ContainerServiceVMDiagnosticsOutput {
	return o
}

func (o ContainerServiceVMDiagnosticsOutput) ToContainerServiceVMDiagnosticsOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsOutput {
	return o
}

func (o ContainerServiceVMDiagnosticsOutput) ToContainerServiceVMDiagnosticsPtrOutput() ContainerServiceVMDiagnosticsPtrOutput {
	return o.ToContainerServiceVMDiagnosticsPtrOutputWithContext(context.Background())
}

func (o ContainerServiceVMDiagnosticsOutput) ToContainerServiceVMDiagnosticsPtrOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceVMDiagnostics) *ContainerServiceVMDiagnostics {
		return &v
	}).(ContainerServiceVMDiagnosticsPtrOutput)
}

func (o ContainerServiceVMDiagnosticsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ContainerServiceVMDiagnostics) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ContainerServiceVMDiagnosticsPtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceVMDiagnosticsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceVMDiagnostics)(nil)).Elem()
}

func (o ContainerServiceVMDiagnosticsPtrOutput) ToContainerServiceVMDiagnosticsPtrOutput() ContainerServiceVMDiagnosticsPtrOutput {
	return o
}

func (o ContainerServiceVMDiagnosticsPtrOutput) ToContainerServiceVMDiagnosticsPtrOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsPtrOutput {
	return o
}

func (o ContainerServiceVMDiagnosticsPtrOutput) Elem() ContainerServiceVMDiagnosticsOutput {
	return o.ApplyT(func(v *ContainerServiceVMDiagnostics) ContainerServiceVMDiagnostics {
		if v != nil {
			return *v
		}
		var ret ContainerServiceVMDiagnostics
		return ret
	}).(ContainerServiceVMDiagnosticsOutput)
}

func (o ContainerServiceVMDiagnosticsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerServiceVMDiagnostics) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type ContainerServiceVMDiagnosticsResponse struct {
	Enabled    bool   `pulumi:"enabled"`
	StorageUri string `pulumi:"storageUri"`
}

type ContainerServiceVMDiagnosticsResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceVMDiagnosticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceVMDiagnosticsResponse)(nil)).Elem()
}

func (o ContainerServiceVMDiagnosticsResponseOutput) ToContainerServiceVMDiagnosticsResponseOutput() ContainerServiceVMDiagnosticsResponseOutput {
	return o
}

func (o ContainerServiceVMDiagnosticsResponseOutput) ToContainerServiceVMDiagnosticsResponseOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsResponseOutput {
	return o
}

func (o ContainerServiceVMDiagnosticsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ContainerServiceVMDiagnosticsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ContainerServiceVMDiagnosticsResponseOutput) StorageUri() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceVMDiagnosticsResponse) string { return v.StorageUri }).(pulumi.StringOutput)
}

type ContainerServiceVMDiagnosticsResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceVMDiagnosticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceVMDiagnosticsResponse)(nil)).Elem()
}

func (o ContainerServiceVMDiagnosticsResponsePtrOutput) ToContainerServiceVMDiagnosticsResponsePtrOutput() ContainerServiceVMDiagnosticsResponsePtrOutput {
	return o
}

func (o ContainerServiceVMDiagnosticsResponsePtrOutput) ToContainerServiceVMDiagnosticsResponsePtrOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsResponsePtrOutput {
	return o
}

func (o ContainerServiceVMDiagnosticsResponsePtrOutput) Elem() ContainerServiceVMDiagnosticsResponseOutput {
	return o.ApplyT(func(v *ContainerServiceVMDiagnosticsResponse) ContainerServiceVMDiagnosticsResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceVMDiagnosticsResponse
		return ret
	}).(ContainerServiceVMDiagnosticsResponseOutput)
}

func (o ContainerServiceVMDiagnosticsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerServiceVMDiagnosticsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ContainerServiceVMDiagnosticsResponsePtrOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceVMDiagnosticsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageUri
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceWindowsProfile struct {
	AdminPassword string `pulumi:"adminPassword"`
	AdminUsername string `pulumi:"adminUsername"`
}





type ContainerServiceWindowsProfileInput interface {
	pulumi.Input

	ToContainerServiceWindowsProfileOutput() ContainerServiceWindowsProfileOutput
	ToContainerServiceWindowsProfileOutputWithContext(context.Context) ContainerServiceWindowsProfileOutput
}

type ContainerServiceWindowsProfileArgs struct {
	AdminPassword pulumi.StringInput `pulumi:"adminPassword"`
	AdminUsername pulumi.StringInput `pulumi:"adminUsername"`
}

func (ContainerServiceWindowsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceWindowsProfile)(nil)).Elem()
}

func (i ContainerServiceWindowsProfileArgs) ToContainerServiceWindowsProfileOutput() ContainerServiceWindowsProfileOutput {
	return i.ToContainerServiceWindowsProfileOutputWithContext(context.Background())
}

func (i ContainerServiceWindowsProfileArgs) ToContainerServiceWindowsProfileOutputWithContext(ctx context.Context) ContainerServiceWindowsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceWindowsProfileOutput)
}

func (i ContainerServiceWindowsProfileArgs) ToContainerServiceWindowsProfilePtrOutput() ContainerServiceWindowsProfilePtrOutput {
	return i.ToContainerServiceWindowsProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceWindowsProfileArgs) ToContainerServiceWindowsProfilePtrOutputWithContext(ctx context.Context) ContainerServiceWindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceWindowsProfileOutput).ToContainerServiceWindowsProfilePtrOutputWithContext(ctx)
}









type ContainerServiceWindowsProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceWindowsProfilePtrOutput() ContainerServiceWindowsProfilePtrOutput
	ToContainerServiceWindowsProfilePtrOutputWithContext(context.Context) ContainerServiceWindowsProfilePtrOutput
}

type containerServiceWindowsProfilePtrType ContainerServiceWindowsProfileArgs

func ContainerServiceWindowsProfilePtr(v *ContainerServiceWindowsProfileArgs) ContainerServiceWindowsProfilePtrInput {
	return (*containerServiceWindowsProfilePtrType)(v)
}

func (*containerServiceWindowsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceWindowsProfile)(nil)).Elem()
}

func (i *containerServiceWindowsProfilePtrType) ToContainerServiceWindowsProfilePtrOutput() ContainerServiceWindowsProfilePtrOutput {
	return i.ToContainerServiceWindowsProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceWindowsProfilePtrType) ToContainerServiceWindowsProfilePtrOutputWithContext(ctx context.Context) ContainerServiceWindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceWindowsProfilePtrOutput)
}

type ContainerServiceWindowsProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceWindowsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceWindowsProfile)(nil)).Elem()
}

func (o ContainerServiceWindowsProfileOutput) ToContainerServiceWindowsProfileOutput() ContainerServiceWindowsProfileOutput {
	return o
}

func (o ContainerServiceWindowsProfileOutput) ToContainerServiceWindowsProfileOutputWithContext(ctx context.Context) ContainerServiceWindowsProfileOutput {
	return o
}

func (o ContainerServiceWindowsProfileOutput) ToContainerServiceWindowsProfilePtrOutput() ContainerServiceWindowsProfilePtrOutput {
	return o.ToContainerServiceWindowsProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceWindowsProfileOutput) ToContainerServiceWindowsProfilePtrOutputWithContext(ctx context.Context) ContainerServiceWindowsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceWindowsProfile) *ContainerServiceWindowsProfile {
		return &v
	}).(ContainerServiceWindowsProfilePtrOutput)
}

func (o ContainerServiceWindowsProfileOutput) AdminPassword() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceWindowsProfile) string { return v.AdminPassword }).(pulumi.StringOutput)
}

func (o ContainerServiceWindowsProfileOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceWindowsProfile) string { return v.AdminUsername }).(pulumi.StringOutput)
}

type ContainerServiceWindowsProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceWindowsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceWindowsProfile)(nil)).Elem()
}

func (o ContainerServiceWindowsProfilePtrOutput) ToContainerServiceWindowsProfilePtrOutput() ContainerServiceWindowsProfilePtrOutput {
	return o
}

func (o ContainerServiceWindowsProfilePtrOutput) ToContainerServiceWindowsProfilePtrOutputWithContext(ctx context.Context) ContainerServiceWindowsProfilePtrOutput {
	return o
}

func (o ContainerServiceWindowsProfilePtrOutput) Elem() ContainerServiceWindowsProfileOutput {
	return o.ApplyT(func(v *ContainerServiceWindowsProfile) ContainerServiceWindowsProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceWindowsProfile
		return ret
	}).(ContainerServiceWindowsProfileOutput)
}

func (o ContainerServiceWindowsProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceWindowsProfile) *string {
		if v == nil {
			return nil
		}
		return &v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceWindowsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceWindowsProfile) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceWindowsProfileResponse struct {
	AdminPassword string `pulumi:"adminPassword"`
	AdminUsername string `pulumi:"adminUsername"`
}

type ContainerServiceWindowsProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceWindowsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceWindowsProfileResponse)(nil)).Elem()
}

func (o ContainerServiceWindowsProfileResponseOutput) ToContainerServiceWindowsProfileResponseOutput() ContainerServiceWindowsProfileResponseOutput {
	return o
}

func (o ContainerServiceWindowsProfileResponseOutput) ToContainerServiceWindowsProfileResponseOutputWithContext(ctx context.Context) ContainerServiceWindowsProfileResponseOutput {
	return o
}

func (o ContainerServiceWindowsProfileResponseOutput) AdminPassword() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceWindowsProfileResponse) string { return v.AdminPassword }).(pulumi.StringOutput)
}

func (o ContainerServiceWindowsProfileResponseOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceWindowsProfileResponse) string { return v.AdminUsername }).(pulumi.StringOutput)
}

type ContainerServiceWindowsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceWindowsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceWindowsProfileResponse)(nil)).Elem()
}

func (o ContainerServiceWindowsProfileResponsePtrOutput) ToContainerServiceWindowsProfileResponsePtrOutput() ContainerServiceWindowsProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceWindowsProfileResponsePtrOutput) ToContainerServiceWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceWindowsProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceWindowsProfileResponsePtrOutput) Elem() ContainerServiceWindowsProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceWindowsProfileResponse) ContainerServiceWindowsProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceWindowsProfileResponse
		return ret
	}).(ContainerServiceWindowsProfileResponseOutput)
}

func (o ContainerServiceWindowsProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceWindowsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceWindowsProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceWindowsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerServiceAgentPoolProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceAgentPoolProfileArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceAgentPoolProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceAgentPoolProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceCustomProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceCustomProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceCustomProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceCustomProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceDiagnosticsProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceDiagnosticsProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceDiagnosticsProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceDiagnosticsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceMasterProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceMasterProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceOrchestratorProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceOrchestratorProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceOrchestratorProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceOrchestratorProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceServicePrincipalProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceServicePrincipalProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceServicePrincipalProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceServicePrincipalProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMDiagnosticsOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMDiagnosticsPtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMDiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMDiagnosticsResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceWindowsProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceWindowsProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceWindowsProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceWindowsProfileResponsePtrOutput{})
}
