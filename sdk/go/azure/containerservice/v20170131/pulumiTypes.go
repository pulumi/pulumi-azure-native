


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





type ContainerServiceAgentPoolProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceAgentPoolProfileResponseOutput() ContainerServiceAgentPoolProfileResponseOutput
	ToContainerServiceAgentPoolProfileResponseOutputWithContext(context.Context) ContainerServiceAgentPoolProfileResponseOutput
}

type ContainerServiceAgentPoolProfileResponseArgs struct {
	Count     pulumi.IntInput    `pulumi:"count"`
	DnsPrefix pulumi.StringInput `pulumi:"dnsPrefix"`
	Fqdn      pulumi.StringInput `pulumi:"fqdn"`
	Name      pulumi.StringInput `pulumi:"name"`
	VmSize    pulumi.StringInput `pulumi:"vmSize"`
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





type ContainerServiceCustomProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceCustomProfileResponseOutput() ContainerServiceCustomProfileResponseOutput
	ToContainerServiceCustomProfileResponseOutputWithContext(context.Context) ContainerServiceCustomProfileResponseOutput
}

type ContainerServiceCustomProfileResponseArgs struct {
	Orchestrator pulumi.StringInput `pulumi:"orchestrator"`
}

func (ContainerServiceCustomProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceCustomProfileResponse)(nil)).Elem()
}

func (i ContainerServiceCustomProfileResponseArgs) ToContainerServiceCustomProfileResponseOutput() ContainerServiceCustomProfileResponseOutput {
	return i.ToContainerServiceCustomProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceCustomProfileResponseArgs) ToContainerServiceCustomProfileResponseOutputWithContext(ctx context.Context) ContainerServiceCustomProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceCustomProfileResponseOutput)
}

func (i ContainerServiceCustomProfileResponseArgs) ToContainerServiceCustomProfileResponsePtrOutput() ContainerServiceCustomProfileResponsePtrOutput {
	return i.ToContainerServiceCustomProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceCustomProfileResponseArgs) ToContainerServiceCustomProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceCustomProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceCustomProfileResponseOutput).ToContainerServiceCustomProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceCustomProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceCustomProfileResponsePtrOutput() ContainerServiceCustomProfileResponsePtrOutput
	ToContainerServiceCustomProfileResponsePtrOutputWithContext(context.Context) ContainerServiceCustomProfileResponsePtrOutput
}

type containerServiceCustomProfileResponsePtrType ContainerServiceCustomProfileResponseArgs

func ContainerServiceCustomProfileResponsePtr(v *ContainerServiceCustomProfileResponseArgs) ContainerServiceCustomProfileResponsePtrInput {
	return (*containerServiceCustomProfileResponsePtrType)(v)
}

func (*containerServiceCustomProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceCustomProfileResponse)(nil)).Elem()
}

func (i *containerServiceCustomProfileResponsePtrType) ToContainerServiceCustomProfileResponsePtrOutput() ContainerServiceCustomProfileResponsePtrOutput {
	return i.ToContainerServiceCustomProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceCustomProfileResponsePtrType) ToContainerServiceCustomProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceCustomProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceCustomProfileResponsePtrOutput)
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

func (o ContainerServiceCustomProfileResponseOutput) ToContainerServiceCustomProfileResponsePtrOutput() ContainerServiceCustomProfileResponsePtrOutput {
	return o.ToContainerServiceCustomProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceCustomProfileResponseOutput) ToContainerServiceCustomProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceCustomProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceCustomProfileResponse) *ContainerServiceCustomProfileResponse {
		return &v
	}).(ContainerServiceCustomProfileResponsePtrOutput)
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





type ContainerServiceDiagnosticsProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceDiagnosticsProfileResponseOutput() ContainerServiceDiagnosticsProfileResponseOutput
	ToContainerServiceDiagnosticsProfileResponseOutputWithContext(context.Context) ContainerServiceDiagnosticsProfileResponseOutput
}

type ContainerServiceDiagnosticsProfileResponseArgs struct {
	VmDiagnostics ContainerServiceVMDiagnosticsResponseInput `pulumi:"vmDiagnostics"`
}

func (ContainerServiceDiagnosticsProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceDiagnosticsProfileResponse)(nil)).Elem()
}

func (i ContainerServiceDiagnosticsProfileResponseArgs) ToContainerServiceDiagnosticsProfileResponseOutput() ContainerServiceDiagnosticsProfileResponseOutput {
	return i.ToContainerServiceDiagnosticsProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceDiagnosticsProfileResponseArgs) ToContainerServiceDiagnosticsProfileResponseOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceDiagnosticsProfileResponseOutput)
}

func (i ContainerServiceDiagnosticsProfileResponseArgs) ToContainerServiceDiagnosticsProfileResponsePtrOutput() ContainerServiceDiagnosticsProfileResponsePtrOutput {
	return i.ToContainerServiceDiagnosticsProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceDiagnosticsProfileResponseArgs) ToContainerServiceDiagnosticsProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceDiagnosticsProfileResponseOutput).ToContainerServiceDiagnosticsProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceDiagnosticsProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceDiagnosticsProfileResponsePtrOutput() ContainerServiceDiagnosticsProfileResponsePtrOutput
	ToContainerServiceDiagnosticsProfileResponsePtrOutputWithContext(context.Context) ContainerServiceDiagnosticsProfileResponsePtrOutput
}

type containerServiceDiagnosticsProfileResponsePtrType ContainerServiceDiagnosticsProfileResponseArgs

func ContainerServiceDiagnosticsProfileResponsePtr(v *ContainerServiceDiagnosticsProfileResponseArgs) ContainerServiceDiagnosticsProfileResponsePtrInput {
	return (*containerServiceDiagnosticsProfileResponsePtrType)(v)
}

func (*containerServiceDiagnosticsProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceDiagnosticsProfileResponse)(nil)).Elem()
}

func (i *containerServiceDiagnosticsProfileResponsePtrType) ToContainerServiceDiagnosticsProfileResponsePtrOutput() ContainerServiceDiagnosticsProfileResponsePtrOutput {
	return i.ToContainerServiceDiagnosticsProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceDiagnosticsProfileResponsePtrType) ToContainerServiceDiagnosticsProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceDiagnosticsProfileResponsePtrOutput)
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

func (o ContainerServiceDiagnosticsProfileResponseOutput) ToContainerServiceDiagnosticsProfileResponsePtrOutput() ContainerServiceDiagnosticsProfileResponsePtrOutput {
	return o.ToContainerServiceDiagnosticsProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceDiagnosticsProfileResponseOutput) ToContainerServiceDiagnosticsProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceDiagnosticsProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceDiagnosticsProfileResponse) *ContainerServiceDiagnosticsProfileResponse {
		return &v
	}).(ContainerServiceDiagnosticsProfileResponsePtrOutput)
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

func (ContainerServiceMasterProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceMasterProfile)(nil)).Elem()
}

func (i ContainerServiceMasterProfileArgs) ToContainerServiceMasterProfileOutput() ContainerServiceMasterProfileOutput {
	return i.ToContainerServiceMasterProfileOutputWithContext(context.Background())
}

func (i ContainerServiceMasterProfileArgs) ToContainerServiceMasterProfileOutputWithContext(ctx context.Context) ContainerServiceMasterProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceMasterProfileOutput)
}

func (i ContainerServiceMasterProfileArgs) ToContainerServiceMasterProfilePtrOutput() ContainerServiceMasterProfilePtrOutput {
	return i.ToContainerServiceMasterProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceMasterProfileArgs) ToContainerServiceMasterProfilePtrOutputWithContext(ctx context.Context) ContainerServiceMasterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceMasterProfileOutput).ToContainerServiceMasterProfilePtrOutputWithContext(ctx)
}









type ContainerServiceMasterProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceMasterProfilePtrOutput() ContainerServiceMasterProfilePtrOutput
	ToContainerServiceMasterProfilePtrOutputWithContext(context.Context) ContainerServiceMasterProfilePtrOutput
}

type containerServiceMasterProfilePtrType ContainerServiceMasterProfileArgs

func ContainerServiceMasterProfilePtr(v *ContainerServiceMasterProfileArgs) ContainerServiceMasterProfilePtrInput {
	return (*containerServiceMasterProfilePtrType)(v)
}

func (*containerServiceMasterProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceMasterProfile)(nil)).Elem()
}

func (i *containerServiceMasterProfilePtrType) ToContainerServiceMasterProfilePtrOutput() ContainerServiceMasterProfilePtrOutput {
	return i.ToContainerServiceMasterProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceMasterProfilePtrType) ToContainerServiceMasterProfilePtrOutputWithContext(ctx context.Context) ContainerServiceMasterProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceMasterProfilePtrOutput)
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

func (o ContainerServiceMasterProfileOutput) ToContainerServiceMasterProfilePtrOutput() ContainerServiceMasterProfilePtrOutput {
	return o.ToContainerServiceMasterProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceMasterProfileOutput) ToContainerServiceMasterProfilePtrOutputWithContext(ctx context.Context) ContainerServiceMasterProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceMasterProfile) *ContainerServiceMasterProfile {
		return &v
	}).(ContainerServiceMasterProfilePtrOutput)
}

func (o ContainerServiceMasterProfileOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerServiceMasterProfile) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ContainerServiceMasterProfileOutput) DnsPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceMasterProfile) string { return v.DnsPrefix }).(pulumi.StringOutput)
}

type ContainerServiceMasterProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceMasterProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceMasterProfile)(nil)).Elem()
}

func (o ContainerServiceMasterProfilePtrOutput) ToContainerServiceMasterProfilePtrOutput() ContainerServiceMasterProfilePtrOutput {
	return o
}

func (o ContainerServiceMasterProfilePtrOutput) ToContainerServiceMasterProfilePtrOutputWithContext(ctx context.Context) ContainerServiceMasterProfilePtrOutput {
	return o
}

func (o ContainerServiceMasterProfilePtrOutput) Elem() ContainerServiceMasterProfileOutput {
	return o.ApplyT(func(v *ContainerServiceMasterProfile) ContainerServiceMasterProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceMasterProfile
		return ret
	}).(ContainerServiceMasterProfileOutput)
}

func (o ContainerServiceMasterProfilePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerServiceMasterProfile) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o ContainerServiceMasterProfilePtrOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceMasterProfile) *string {
		if v == nil {
			return nil
		}
		return &v.DnsPrefix
	}).(pulumi.StringPtrOutput)
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





type ContainerServiceMasterProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceMasterProfileResponseOutput() ContainerServiceMasterProfileResponseOutput
	ToContainerServiceMasterProfileResponseOutputWithContext(context.Context) ContainerServiceMasterProfileResponseOutput
}

type ContainerServiceMasterProfileResponseArgs struct {
	Count     pulumi.IntPtrInput `pulumi:"count"`
	DnsPrefix pulumi.StringInput `pulumi:"dnsPrefix"`
	Fqdn      pulumi.StringInput `pulumi:"fqdn"`
}

func (ContainerServiceMasterProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceMasterProfileResponse)(nil)).Elem()
}

func (i ContainerServiceMasterProfileResponseArgs) ToContainerServiceMasterProfileResponseOutput() ContainerServiceMasterProfileResponseOutput {
	return i.ToContainerServiceMasterProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceMasterProfileResponseArgs) ToContainerServiceMasterProfileResponseOutputWithContext(ctx context.Context) ContainerServiceMasterProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceMasterProfileResponseOutput)
}

func (i ContainerServiceMasterProfileResponseArgs) ToContainerServiceMasterProfileResponsePtrOutput() ContainerServiceMasterProfileResponsePtrOutput {
	return i.ToContainerServiceMasterProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceMasterProfileResponseArgs) ToContainerServiceMasterProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceMasterProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceMasterProfileResponseOutput).ToContainerServiceMasterProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceMasterProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceMasterProfileResponsePtrOutput() ContainerServiceMasterProfileResponsePtrOutput
	ToContainerServiceMasterProfileResponsePtrOutputWithContext(context.Context) ContainerServiceMasterProfileResponsePtrOutput
}

type containerServiceMasterProfileResponsePtrType ContainerServiceMasterProfileResponseArgs

func ContainerServiceMasterProfileResponsePtr(v *ContainerServiceMasterProfileResponseArgs) ContainerServiceMasterProfileResponsePtrInput {
	return (*containerServiceMasterProfileResponsePtrType)(v)
}

func (*containerServiceMasterProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceMasterProfileResponse)(nil)).Elem()
}

func (i *containerServiceMasterProfileResponsePtrType) ToContainerServiceMasterProfileResponsePtrOutput() ContainerServiceMasterProfileResponsePtrOutput {
	return i.ToContainerServiceMasterProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceMasterProfileResponsePtrType) ToContainerServiceMasterProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceMasterProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceMasterProfileResponsePtrOutput)
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

func (o ContainerServiceMasterProfileResponseOutput) ToContainerServiceMasterProfileResponsePtrOutput() ContainerServiceMasterProfileResponsePtrOutput {
	return o.ToContainerServiceMasterProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceMasterProfileResponseOutput) ToContainerServiceMasterProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceMasterProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceMasterProfileResponse) *ContainerServiceMasterProfileResponse {
		return &v
	}).(ContainerServiceMasterProfileResponsePtrOutput)
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

type ContainerServiceMasterProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceMasterProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceMasterProfileResponse)(nil)).Elem()
}

func (o ContainerServiceMasterProfileResponsePtrOutput) ToContainerServiceMasterProfileResponsePtrOutput() ContainerServiceMasterProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceMasterProfileResponsePtrOutput) ToContainerServiceMasterProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceMasterProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceMasterProfileResponsePtrOutput) Elem() ContainerServiceMasterProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceMasterProfileResponse) ContainerServiceMasterProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceMasterProfileResponse
		return ret
	}).(ContainerServiceMasterProfileResponseOutput)
}

func (o ContainerServiceMasterProfileResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerServiceMasterProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o ContainerServiceMasterProfileResponsePtrOutput) DnsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceMasterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DnsPrefix
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceMasterProfileResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceMasterProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Fqdn
	}).(pulumi.StringPtrOutput)
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





type ContainerServiceOrchestratorProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceOrchestratorProfileResponseOutput() ContainerServiceOrchestratorProfileResponseOutput
	ToContainerServiceOrchestratorProfileResponseOutputWithContext(context.Context) ContainerServiceOrchestratorProfileResponseOutput
}

type ContainerServiceOrchestratorProfileResponseArgs struct {
	OrchestratorType pulumi.StringInput `pulumi:"orchestratorType"`
}

func (ContainerServiceOrchestratorProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceOrchestratorProfileResponse)(nil)).Elem()
}

func (i ContainerServiceOrchestratorProfileResponseArgs) ToContainerServiceOrchestratorProfileResponseOutput() ContainerServiceOrchestratorProfileResponseOutput {
	return i.ToContainerServiceOrchestratorProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceOrchestratorProfileResponseArgs) ToContainerServiceOrchestratorProfileResponseOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceOrchestratorProfileResponseOutput)
}

func (i ContainerServiceOrchestratorProfileResponseArgs) ToContainerServiceOrchestratorProfileResponsePtrOutput() ContainerServiceOrchestratorProfileResponsePtrOutput {
	return i.ToContainerServiceOrchestratorProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceOrchestratorProfileResponseArgs) ToContainerServiceOrchestratorProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceOrchestratorProfileResponseOutput).ToContainerServiceOrchestratorProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceOrchestratorProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceOrchestratorProfileResponsePtrOutput() ContainerServiceOrchestratorProfileResponsePtrOutput
	ToContainerServiceOrchestratorProfileResponsePtrOutputWithContext(context.Context) ContainerServiceOrchestratorProfileResponsePtrOutput
}

type containerServiceOrchestratorProfileResponsePtrType ContainerServiceOrchestratorProfileResponseArgs

func ContainerServiceOrchestratorProfileResponsePtr(v *ContainerServiceOrchestratorProfileResponseArgs) ContainerServiceOrchestratorProfileResponsePtrInput {
	return (*containerServiceOrchestratorProfileResponsePtrType)(v)
}

func (*containerServiceOrchestratorProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceOrchestratorProfileResponse)(nil)).Elem()
}

func (i *containerServiceOrchestratorProfileResponsePtrType) ToContainerServiceOrchestratorProfileResponsePtrOutput() ContainerServiceOrchestratorProfileResponsePtrOutput {
	return i.ToContainerServiceOrchestratorProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceOrchestratorProfileResponsePtrType) ToContainerServiceOrchestratorProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceOrchestratorProfileResponsePtrOutput)
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

func (o ContainerServiceOrchestratorProfileResponseOutput) ToContainerServiceOrchestratorProfileResponsePtrOutput() ContainerServiceOrchestratorProfileResponsePtrOutput {
	return o.ToContainerServiceOrchestratorProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceOrchestratorProfileResponseOutput) ToContainerServiceOrchestratorProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceOrchestratorProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceOrchestratorProfileResponse) *ContainerServiceOrchestratorProfileResponse {
		return &v
	}).(ContainerServiceOrchestratorProfileResponsePtrOutput)
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





type ContainerServiceServicePrincipalProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceServicePrincipalProfileResponseOutput() ContainerServiceServicePrincipalProfileResponseOutput
	ToContainerServiceServicePrincipalProfileResponseOutputWithContext(context.Context) ContainerServiceServicePrincipalProfileResponseOutput
}

type ContainerServiceServicePrincipalProfileResponseArgs struct {
	ClientId pulumi.StringInput `pulumi:"clientId"`
	Secret   pulumi.StringInput `pulumi:"secret"`
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





type ContainerServiceVMDiagnosticsResponseInput interface {
	pulumi.Input

	ToContainerServiceVMDiagnosticsResponseOutput() ContainerServiceVMDiagnosticsResponseOutput
	ToContainerServiceVMDiagnosticsResponseOutputWithContext(context.Context) ContainerServiceVMDiagnosticsResponseOutput
}

type ContainerServiceVMDiagnosticsResponseArgs struct {
	Enabled    pulumi.BoolInput   `pulumi:"enabled"`
	StorageUri pulumi.StringInput `pulumi:"storageUri"`
}

func (ContainerServiceVMDiagnosticsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceVMDiagnosticsResponse)(nil)).Elem()
}

func (i ContainerServiceVMDiagnosticsResponseArgs) ToContainerServiceVMDiagnosticsResponseOutput() ContainerServiceVMDiagnosticsResponseOutput {
	return i.ToContainerServiceVMDiagnosticsResponseOutputWithContext(context.Background())
}

func (i ContainerServiceVMDiagnosticsResponseArgs) ToContainerServiceVMDiagnosticsResponseOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceVMDiagnosticsResponseOutput)
}

func (i ContainerServiceVMDiagnosticsResponseArgs) ToContainerServiceVMDiagnosticsResponsePtrOutput() ContainerServiceVMDiagnosticsResponsePtrOutput {
	return i.ToContainerServiceVMDiagnosticsResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceVMDiagnosticsResponseArgs) ToContainerServiceVMDiagnosticsResponsePtrOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceVMDiagnosticsResponseOutput).ToContainerServiceVMDiagnosticsResponsePtrOutputWithContext(ctx)
}









type ContainerServiceVMDiagnosticsResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceVMDiagnosticsResponsePtrOutput() ContainerServiceVMDiagnosticsResponsePtrOutput
	ToContainerServiceVMDiagnosticsResponsePtrOutputWithContext(context.Context) ContainerServiceVMDiagnosticsResponsePtrOutput
}

type containerServiceVMDiagnosticsResponsePtrType ContainerServiceVMDiagnosticsResponseArgs

func ContainerServiceVMDiagnosticsResponsePtr(v *ContainerServiceVMDiagnosticsResponseArgs) ContainerServiceVMDiagnosticsResponsePtrInput {
	return (*containerServiceVMDiagnosticsResponsePtrType)(v)
}

func (*containerServiceVMDiagnosticsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceVMDiagnosticsResponse)(nil)).Elem()
}

func (i *containerServiceVMDiagnosticsResponsePtrType) ToContainerServiceVMDiagnosticsResponsePtrOutput() ContainerServiceVMDiagnosticsResponsePtrOutput {
	return i.ToContainerServiceVMDiagnosticsResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceVMDiagnosticsResponsePtrType) ToContainerServiceVMDiagnosticsResponsePtrOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceVMDiagnosticsResponsePtrOutput)
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

func (o ContainerServiceVMDiagnosticsResponseOutput) ToContainerServiceVMDiagnosticsResponsePtrOutput() ContainerServiceVMDiagnosticsResponsePtrOutput {
	return o.ToContainerServiceVMDiagnosticsResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceVMDiagnosticsResponseOutput) ToContainerServiceVMDiagnosticsResponsePtrOutputWithContext(ctx context.Context) ContainerServiceVMDiagnosticsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceVMDiagnosticsResponse) *ContainerServiceVMDiagnosticsResponse {
		return &v
	}).(ContainerServiceVMDiagnosticsResponsePtrOutput)
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





type ContainerServiceWindowsProfileResponseInput interface {
	pulumi.Input

	ToContainerServiceWindowsProfileResponseOutput() ContainerServiceWindowsProfileResponseOutput
	ToContainerServiceWindowsProfileResponseOutputWithContext(context.Context) ContainerServiceWindowsProfileResponseOutput
}

type ContainerServiceWindowsProfileResponseArgs struct {
	AdminPassword pulumi.StringInput `pulumi:"adminPassword"`
	AdminUsername pulumi.StringInput `pulumi:"adminUsername"`
}

func (ContainerServiceWindowsProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceWindowsProfileResponse)(nil)).Elem()
}

func (i ContainerServiceWindowsProfileResponseArgs) ToContainerServiceWindowsProfileResponseOutput() ContainerServiceWindowsProfileResponseOutput {
	return i.ToContainerServiceWindowsProfileResponseOutputWithContext(context.Background())
}

func (i ContainerServiceWindowsProfileResponseArgs) ToContainerServiceWindowsProfileResponseOutputWithContext(ctx context.Context) ContainerServiceWindowsProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceWindowsProfileResponseOutput)
}

func (i ContainerServiceWindowsProfileResponseArgs) ToContainerServiceWindowsProfileResponsePtrOutput() ContainerServiceWindowsProfileResponsePtrOutput {
	return i.ToContainerServiceWindowsProfileResponsePtrOutputWithContext(context.Background())
}

func (i ContainerServiceWindowsProfileResponseArgs) ToContainerServiceWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceWindowsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceWindowsProfileResponseOutput).ToContainerServiceWindowsProfileResponsePtrOutputWithContext(ctx)
}









type ContainerServiceWindowsProfileResponsePtrInput interface {
	pulumi.Input

	ToContainerServiceWindowsProfileResponsePtrOutput() ContainerServiceWindowsProfileResponsePtrOutput
	ToContainerServiceWindowsProfileResponsePtrOutputWithContext(context.Context) ContainerServiceWindowsProfileResponsePtrOutput
}

type containerServiceWindowsProfileResponsePtrType ContainerServiceWindowsProfileResponseArgs

func ContainerServiceWindowsProfileResponsePtr(v *ContainerServiceWindowsProfileResponseArgs) ContainerServiceWindowsProfileResponsePtrInput {
	return (*containerServiceWindowsProfileResponsePtrType)(v)
}

func (*containerServiceWindowsProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceWindowsProfileResponse)(nil)).Elem()
}

func (i *containerServiceWindowsProfileResponsePtrType) ToContainerServiceWindowsProfileResponsePtrOutput() ContainerServiceWindowsProfileResponsePtrOutput {
	return i.ToContainerServiceWindowsProfileResponsePtrOutputWithContext(context.Background())
}

func (i *containerServiceWindowsProfileResponsePtrType) ToContainerServiceWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceWindowsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceWindowsProfileResponsePtrOutput)
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

func (o ContainerServiceWindowsProfileResponseOutput) ToContainerServiceWindowsProfileResponsePtrOutput() ContainerServiceWindowsProfileResponsePtrOutput {
	return o.ToContainerServiceWindowsProfileResponsePtrOutputWithContext(context.Background())
}

func (o ContainerServiceWindowsProfileResponseOutput) ToContainerServiceWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceWindowsProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceWindowsProfileResponse) *ContainerServiceWindowsProfileResponse {
		return &v
	}).(ContainerServiceWindowsProfileResponsePtrOutput)
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
	pulumi.RegisterOutputType(ContainerServiceLinuxProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceMasterProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceMasterProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceMasterProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceMasterProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceOrchestratorProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceOrchestratorProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceOrchestratorProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceOrchestratorProfileResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ContainerServiceVMDiagnosticsOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMDiagnosticsPtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMDiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceVMDiagnosticsResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceWindowsProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceWindowsProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceWindowsProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceWindowsProfileResponsePtrOutput{})
}
