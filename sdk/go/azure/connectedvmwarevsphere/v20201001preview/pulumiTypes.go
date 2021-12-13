


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ErrorDetailResponse struct {
	Code    string                `pulumi:"code"`
	Details []ErrorDetailResponse `pulumi:"details"`
	Message string                `pulumi:"message"`
	Target  *string               `pulumi:"target"`
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorDetailResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type GuestAgentProfileResponse struct {
	AgentVersion     string                `pulumi:"agentVersion"`
	ErrorDetails     []ErrorDetailResponse `pulumi:"errorDetails"`
	LastStatusChange string                `pulumi:"lastStatusChange"`
	Status           string                `pulumi:"status"`
	VmUuid           string                `pulumi:"vmUuid"`
}

type GuestAgentProfileResponseOutput struct{ *pulumi.OutputState }

func (GuestAgentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestAgentProfileResponse)(nil)).Elem()
}

func (o GuestAgentProfileResponseOutput) ToGuestAgentProfileResponseOutput() GuestAgentProfileResponseOutput {
	return o
}

func (o GuestAgentProfileResponseOutput) ToGuestAgentProfileResponseOutputWithContext(ctx context.Context) GuestAgentProfileResponseOutput {
	return o
}

func (o GuestAgentProfileResponseOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) string { return v.AgentVersion }).(pulumi.StringOutput)
}

func (o GuestAgentProfileResponseOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) []ErrorDetailResponse { return v.ErrorDetails }).(ErrorDetailResponseArrayOutput)
}

func (o GuestAgentProfileResponseOutput) LastStatusChange() pulumi.StringOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) string { return v.LastStatusChange }).(pulumi.StringOutput)
}

func (o GuestAgentProfileResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o GuestAgentProfileResponseOutput) VmUuid() pulumi.StringOutput {
	return o.ApplyT(func(v GuestAgentProfileResponse) string { return v.VmUuid }).(pulumi.StringOutput)
}

type GuestAgentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestAgentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestAgentProfileResponse)(nil)).Elem()
}

func (o GuestAgentProfileResponsePtrOutput) ToGuestAgentProfileResponsePtrOutput() GuestAgentProfileResponsePtrOutput {
	return o
}

func (o GuestAgentProfileResponsePtrOutput) ToGuestAgentProfileResponsePtrOutputWithContext(ctx context.Context) GuestAgentProfileResponsePtrOutput {
	return o
}

func (o GuestAgentProfileResponsePtrOutput) Elem() GuestAgentProfileResponseOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) GuestAgentProfileResponse {
		if v != nil {
			return *v
		}
		var ret GuestAgentProfileResponse
		return ret
	}).(GuestAgentProfileResponseOutput)
}

func (o GuestAgentProfileResponsePtrOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AgentVersion
	}).(pulumi.StringPtrOutput)
}

func (o GuestAgentProfileResponsePtrOutput) ErrorDetails() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) []ErrorDetailResponse {
		if v == nil {
			return nil
		}
		return v.ErrorDetails
	}).(ErrorDetailResponseArrayOutput)
}

func (o GuestAgentProfileResponsePtrOutput) LastStatusChange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastStatusChange
	}).(pulumi.StringPtrOutput)
}

func (o GuestAgentProfileResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o GuestAgentProfileResponsePtrOutput) VmUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VmUuid
	}).(pulumi.StringPtrOutput)
}

type GuestCredential struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}





type GuestCredentialInput interface {
	pulumi.Input

	ToGuestCredentialOutput() GuestCredentialOutput
	ToGuestCredentialOutputWithContext(context.Context) GuestCredentialOutput
}

type GuestCredentialArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (GuestCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestCredential)(nil)).Elem()
}

func (i GuestCredentialArgs) ToGuestCredentialOutput() GuestCredentialOutput {
	return i.ToGuestCredentialOutputWithContext(context.Background())
}

func (i GuestCredentialArgs) ToGuestCredentialOutputWithContext(ctx context.Context) GuestCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestCredentialOutput)
}

func (i GuestCredentialArgs) ToGuestCredentialPtrOutput() GuestCredentialPtrOutput {
	return i.ToGuestCredentialPtrOutputWithContext(context.Background())
}

func (i GuestCredentialArgs) ToGuestCredentialPtrOutputWithContext(ctx context.Context) GuestCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestCredentialOutput).ToGuestCredentialPtrOutputWithContext(ctx)
}









type GuestCredentialPtrInput interface {
	pulumi.Input

	ToGuestCredentialPtrOutput() GuestCredentialPtrOutput
	ToGuestCredentialPtrOutputWithContext(context.Context) GuestCredentialPtrOutput
}

type guestCredentialPtrType GuestCredentialArgs

func GuestCredentialPtr(v *GuestCredentialArgs) GuestCredentialPtrInput {
	return (*guestCredentialPtrType)(v)
}

func (*guestCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestCredential)(nil)).Elem()
}

func (i *guestCredentialPtrType) ToGuestCredentialPtrOutput() GuestCredentialPtrOutput {
	return i.ToGuestCredentialPtrOutputWithContext(context.Background())
}

func (i *guestCredentialPtrType) ToGuestCredentialPtrOutputWithContext(ctx context.Context) GuestCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestCredentialPtrOutput)
}

type GuestCredentialOutput struct{ *pulumi.OutputState }

func (GuestCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestCredential)(nil)).Elem()
}

func (o GuestCredentialOutput) ToGuestCredentialOutput() GuestCredentialOutput {
	return o
}

func (o GuestCredentialOutput) ToGuestCredentialOutputWithContext(ctx context.Context) GuestCredentialOutput {
	return o
}

func (o GuestCredentialOutput) ToGuestCredentialPtrOutput() GuestCredentialPtrOutput {
	return o.ToGuestCredentialPtrOutputWithContext(context.Background())
}

func (o GuestCredentialOutput) ToGuestCredentialPtrOutputWithContext(ctx context.Context) GuestCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestCredential) *GuestCredential {
		return &v
	}).(GuestCredentialPtrOutput)
}

func (o GuestCredentialOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestCredential) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GuestCredentialOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestCredential) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GuestCredentialPtrOutput struct{ *pulumi.OutputState }

func (GuestCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestCredential)(nil)).Elem()
}

func (o GuestCredentialPtrOutput) ToGuestCredentialPtrOutput() GuestCredentialPtrOutput {
	return o
}

func (o GuestCredentialPtrOutput) ToGuestCredentialPtrOutputWithContext(ctx context.Context) GuestCredentialPtrOutput {
	return o
}

func (o GuestCredentialPtrOutput) Elem() GuestCredentialOutput {
	return o.ApplyT(func(v *GuestCredential) GuestCredential {
		if v != nil {
			return *v
		}
		var ret GuestCredential
		return ret
	}).(GuestCredentialOutput)
}

func (o GuestCredentialPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestCredential) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o GuestCredentialPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestCredential) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type GuestCredentialResponse struct {
	Username *string `pulumi:"username"`
}

type GuestCredentialResponseOutput struct{ *pulumi.OutputState }

func (GuestCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestCredentialResponse)(nil)).Elem()
}

func (o GuestCredentialResponseOutput) ToGuestCredentialResponseOutput() GuestCredentialResponseOutput {
	return o
}

func (o GuestCredentialResponseOutput) ToGuestCredentialResponseOutputWithContext(ctx context.Context) GuestCredentialResponseOutput {
	return o
}

func (o GuestCredentialResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestCredentialResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GuestCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestCredentialResponse)(nil)).Elem()
}

func (o GuestCredentialResponsePtrOutput) ToGuestCredentialResponsePtrOutput() GuestCredentialResponsePtrOutput {
	return o
}

func (o GuestCredentialResponsePtrOutput) ToGuestCredentialResponsePtrOutputWithContext(ctx context.Context) GuestCredentialResponsePtrOutput {
	return o
}

func (o GuestCredentialResponsePtrOutput) Elem() GuestCredentialResponseOutput {
	return o.ApplyT(func(v *GuestCredentialResponse) GuestCredentialResponse {
		if v != nil {
			return *v
		}
		var ret GuestCredentialResponse
		return ret
	}).(GuestCredentialResponseOutput)
}

func (o GuestCredentialResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type HardwareProfile struct {
	MemorySizeMB      *int `pulumi:"memorySizeMB"`
	NumCPUs           *int `pulumi:"numCPUs"`
	NumCoresPerSocket *int `pulumi:"numCoresPerSocket"`
}





type HardwareProfileInput interface {
	pulumi.Input

	ToHardwareProfileOutput() HardwareProfileOutput
	ToHardwareProfileOutputWithContext(context.Context) HardwareProfileOutput
}

type HardwareProfileArgs struct {
	MemorySizeMB      pulumi.IntPtrInput `pulumi:"memorySizeMB"`
	NumCPUs           pulumi.IntPtrInput `pulumi:"numCPUs"`
	NumCoresPerSocket pulumi.IntPtrInput `pulumi:"numCoresPerSocket"`
}

func (HardwareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (i HardwareProfileArgs) ToHardwareProfileOutput() HardwareProfileOutput {
	return i.ToHardwareProfileOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput)
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput).ToHardwareProfilePtrOutputWithContext(ctx)
}









type HardwareProfilePtrInput interface {
	pulumi.Input

	ToHardwareProfilePtrOutput() HardwareProfilePtrOutput
	ToHardwareProfilePtrOutputWithContext(context.Context) HardwareProfilePtrOutput
}

type hardwareProfilePtrType HardwareProfileArgs

func HardwareProfilePtr(v *HardwareProfileArgs) HardwareProfilePtrInput {
	return (*hardwareProfilePtrType)(v)
}

func (*hardwareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfilePtrOutput)
}

type HardwareProfileOutput struct{ *pulumi.OutputState }

func (HardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (o HardwareProfileOutput) ToHardwareProfileOutput() HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HardwareProfile) *HardwareProfile {
		return &v
	}).(HardwareProfilePtrOutput)
}

func (o HardwareProfileOutput) MemorySizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *int { return v.MemorySizeMB }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileOutput) NumCPUs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *int { return v.NumCPUs }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileOutput) NumCoresPerSocket() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *int { return v.NumCoresPerSocket }).(pulumi.IntPtrOutput)
}

type HardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) Elem() HardwareProfileOutput {
	return o.ApplyT(func(v *HardwareProfile) HardwareProfile {
		if v != nil {
			return *v
		}
		var ret HardwareProfile
		return ret
	}).(HardwareProfileOutput)
}

func (o HardwareProfilePtrOutput) MemorySizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.MemorySizeMB
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfilePtrOutput) NumCPUs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.NumCPUs
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfilePtrOutput) NumCoresPerSocket() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.NumCoresPerSocket
	}).(pulumi.IntPtrOutput)
}

type HardwareProfileResponse struct {
	CpuHotAddEnabled    bool `pulumi:"cpuHotAddEnabled"`
	CpuHotRemoveEnabled bool `pulumi:"cpuHotRemoveEnabled"`
	MemoryHotAddEnabled bool `pulumi:"memoryHotAddEnabled"`
	MemorySizeMB        *int `pulumi:"memorySizeMB"`
	NumCPUs             *int `pulumi:"numCPUs"`
	NumCoresPerSocket   *int `pulumi:"numCoresPerSocket"`
}

type HardwareProfileResponseOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutput() HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutputWithContext(ctx context.Context) HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) CpuHotAddEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v HardwareProfileResponse) bool { return v.CpuHotAddEnabled }).(pulumi.BoolOutput)
}

func (o HardwareProfileResponseOutput) CpuHotRemoveEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v HardwareProfileResponse) bool { return v.CpuHotRemoveEnabled }).(pulumi.BoolOutput)
}

func (o HardwareProfileResponseOutput) MemoryHotAddEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v HardwareProfileResponse) bool { return v.MemoryHotAddEnabled }).(pulumi.BoolOutput)
}

func (o HardwareProfileResponseOutput) MemorySizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *int { return v.MemorySizeMB }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponseOutput) NumCPUs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *int { return v.NumCPUs }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponseOutput) NumCoresPerSocket() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *int { return v.NumCoresPerSocket }).(pulumi.IntPtrOutput)
}

type HardwareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutput() HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutputWithContext(ctx context.Context) HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) Elem() HardwareProfileResponseOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) HardwareProfileResponse {
		if v != nil {
			return *v
		}
		var ret HardwareProfileResponse
		return ret
	}).(HardwareProfileResponseOutput)
}

func (o HardwareProfileResponsePtrOutput) CpuHotAddEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CpuHotAddEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) CpuHotRemoveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CpuHotRemoveEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) MemoryHotAddEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.MemoryHotAddEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) MemorySizeMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.MemorySizeMB
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) NumCPUs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.NumCPUs
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) NumCoresPerSocket() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.NumCoresPerSocket
	}).(pulumi.IntPtrOutput)
}

type HttpProxyConfiguration struct {
	HttpsProxy *string `pulumi:"httpsProxy"`
}





type HttpProxyConfigurationInput interface {
	pulumi.Input

	ToHttpProxyConfigurationOutput() HttpProxyConfigurationOutput
	ToHttpProxyConfigurationOutputWithContext(context.Context) HttpProxyConfigurationOutput
}

type HttpProxyConfigurationArgs struct {
	HttpsProxy pulumi.StringPtrInput `pulumi:"httpsProxy"`
}

func (HttpProxyConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProxyConfiguration)(nil)).Elem()
}

func (i HttpProxyConfigurationArgs) ToHttpProxyConfigurationOutput() HttpProxyConfigurationOutput {
	return i.ToHttpProxyConfigurationOutputWithContext(context.Background())
}

func (i HttpProxyConfigurationArgs) ToHttpProxyConfigurationOutputWithContext(ctx context.Context) HttpProxyConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProxyConfigurationOutput)
}

func (i HttpProxyConfigurationArgs) ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput {
	return i.ToHttpProxyConfigurationPtrOutputWithContext(context.Background())
}

func (i HttpProxyConfigurationArgs) ToHttpProxyConfigurationPtrOutputWithContext(ctx context.Context) HttpProxyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProxyConfigurationOutput).ToHttpProxyConfigurationPtrOutputWithContext(ctx)
}









type HttpProxyConfigurationPtrInput interface {
	pulumi.Input

	ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput
	ToHttpProxyConfigurationPtrOutputWithContext(context.Context) HttpProxyConfigurationPtrOutput
}

type httpProxyConfigurationPtrType HttpProxyConfigurationArgs

func HttpProxyConfigurationPtr(v *HttpProxyConfigurationArgs) HttpProxyConfigurationPtrInput {
	return (*httpProxyConfigurationPtrType)(v)
}

func (*httpProxyConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProxyConfiguration)(nil)).Elem()
}

func (i *httpProxyConfigurationPtrType) ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput {
	return i.ToHttpProxyConfigurationPtrOutputWithContext(context.Background())
}

func (i *httpProxyConfigurationPtrType) ToHttpProxyConfigurationPtrOutputWithContext(ctx context.Context) HttpProxyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpProxyConfigurationPtrOutput)
}

type HttpProxyConfigurationOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProxyConfiguration)(nil)).Elem()
}

func (o HttpProxyConfigurationOutput) ToHttpProxyConfigurationOutput() HttpProxyConfigurationOutput {
	return o
}

func (o HttpProxyConfigurationOutput) ToHttpProxyConfigurationOutputWithContext(ctx context.Context) HttpProxyConfigurationOutput {
	return o
}

func (o HttpProxyConfigurationOutput) ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput {
	return o.ToHttpProxyConfigurationPtrOutputWithContext(context.Background())
}

func (o HttpProxyConfigurationOutput) ToHttpProxyConfigurationPtrOutputWithContext(ctx context.Context) HttpProxyConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpProxyConfiguration) *HttpProxyConfiguration {
		return &v
	}).(HttpProxyConfigurationPtrOutput)
}

func (o HttpProxyConfigurationOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfiguration) *string { return v.HttpsProxy }).(pulumi.StringPtrOutput)
}

type HttpProxyConfigurationPtrOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProxyConfiguration)(nil)).Elem()
}

func (o HttpProxyConfigurationPtrOutput) ToHttpProxyConfigurationPtrOutput() HttpProxyConfigurationPtrOutput {
	return o
}

func (o HttpProxyConfigurationPtrOutput) ToHttpProxyConfigurationPtrOutputWithContext(ctx context.Context) HttpProxyConfigurationPtrOutput {
	return o
}

func (o HttpProxyConfigurationPtrOutput) Elem() HttpProxyConfigurationOutput {
	return o.ApplyT(func(v *HttpProxyConfiguration) HttpProxyConfiguration {
		if v != nil {
			return *v
		}
		var ret HttpProxyConfiguration
		return ret
	}).(HttpProxyConfigurationOutput)
}

func (o HttpProxyConfigurationPtrOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.HttpsProxy
	}).(pulumi.StringPtrOutput)
}

type HttpProxyConfigurationResponse struct {
	HttpsProxy *string `pulumi:"httpsProxy"`
}

type HttpProxyConfigurationResponseOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpProxyConfigurationResponse)(nil)).Elem()
}

func (o HttpProxyConfigurationResponseOutput) ToHttpProxyConfigurationResponseOutput() HttpProxyConfigurationResponseOutput {
	return o
}

func (o HttpProxyConfigurationResponseOutput) ToHttpProxyConfigurationResponseOutputWithContext(ctx context.Context) HttpProxyConfigurationResponseOutput {
	return o
}

func (o HttpProxyConfigurationResponseOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpProxyConfigurationResponse) *string { return v.HttpsProxy }).(pulumi.StringPtrOutput)
}

type HttpProxyConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpProxyConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpProxyConfigurationResponse)(nil)).Elem()
}

func (o HttpProxyConfigurationResponsePtrOutput) ToHttpProxyConfigurationResponsePtrOutput() HttpProxyConfigurationResponsePtrOutput {
	return o
}

func (o HttpProxyConfigurationResponsePtrOutput) ToHttpProxyConfigurationResponsePtrOutputWithContext(ctx context.Context) HttpProxyConfigurationResponsePtrOutput {
	return o
}

func (o HttpProxyConfigurationResponsePtrOutput) Elem() HttpProxyConfigurationResponseOutput {
	return o.ApplyT(func(v *HttpProxyConfigurationResponse) HttpProxyConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret HttpProxyConfigurationResponse
		return ret
	}).(HttpProxyConfigurationResponseOutput)
}

func (o HttpProxyConfigurationResponsePtrOutput) HttpsProxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpProxyConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpsProxy
	}).(pulumi.StringPtrOutput)
}

type Identity struct {
	Type string `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Identity) string { return v.Type }).(pulumi.StringOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewResponseStatus struct {
	Code          string `pulumi:"code"`
	DisplayStatus string `pulumi:"displayStatus"`
	Level         string `pulumi:"level"`
	Message       string `pulumi:"message"`
	Time          string `pulumi:"time"`
}

type MachineExtensionInstanceViewResponseStatusOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutput() MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) ToMachineExtensionInstanceViewResponseStatusOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Code }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) DisplayStatus() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.DisplayStatus }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Level }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Message }).(pulumi.StringOutput)
}

func (o MachineExtensionInstanceViewResponseStatusOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewResponseStatus) string { return v.Time }).(pulumi.StringOutput)
}

type MachineExtensionInstanceViewResponseStatusPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewResponseStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewResponseStatus)(nil)).Elem()
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutput() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) ToMachineExtensionInstanceViewResponseStatusPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Elem() MachineExtensionInstanceViewResponseStatusOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) MachineExtensionInstanceViewResponseStatus {
		if v != nil {
			return *v
		}
		var ret MachineExtensionInstanceViewResponseStatus
		return ret
	}).(MachineExtensionInstanceViewResponseStatusOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Level
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionInstanceViewResponseStatusPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewResponseStatus) *string {
		if v == nil {
			return nil
		}
		return &v.Time
	}).(pulumi.StringPtrOutput)
}

type MachineExtensionPropertiesResponseInstanceView struct {
	Name               string                                      `pulumi:"name"`
	Status             *MachineExtensionInstanceViewResponseStatus `pulumi:"status"`
	Type               string                                      `pulumi:"type"`
	TypeHandlerVersion string                                      `pulumi:"typeHandlerVersion"`
}

type MachineExtensionPropertiesResponseInstanceViewOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponseInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewOutput() MachineExtensionPropertiesResponseInstanceViewOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) ToMachineExtensionPropertiesResponseInstanceViewOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.Name }).(pulumi.StringOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) *MachineExtensionInstanceViewResponseStatus {
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.Type }).(pulumi.StringOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionPropertiesResponseInstanceView) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type MachineExtensionPropertiesResponseInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionPropertiesResponseInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionPropertiesResponseInstanceView)(nil)).Elem()
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutput() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) ToMachineExtensionPropertiesResponseInstanceViewPtrOutputWithContext(ctx context.Context) MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Elem() MachineExtensionPropertiesResponseInstanceViewOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) MachineExtensionPropertiesResponseInstanceView {
		if v != nil {
			return *v
		}
		var ret MachineExtensionPropertiesResponseInstanceView
		return ret
	}).(MachineExtensionPropertiesResponseInstanceViewOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Status() MachineExtensionInstanceViewResponseStatusPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *MachineExtensionInstanceViewResponseStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(MachineExtensionInstanceViewResponseStatusPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MachineExtensionPropertiesResponseInstanceViewPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionPropertiesResponseInstanceView) *string {
		if v == nil {
			return nil
		}
		return &v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type NetworkInterface struct {
	DeviceKey   *int           `pulumi:"deviceKey"`
	IpSettings  *NicIPSettings `pulumi:"ipSettings"`
	Name        *string        `pulumi:"name"`
	NetworkId   *string        `pulumi:"networkId"`
	NicType     *string        `pulumi:"nicType"`
	PowerOnBoot *string        `pulumi:"powerOnBoot"`
}





type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(context.Context) NetworkInterfaceOutput
}

type NetworkInterfaceArgs struct {
	DeviceKey   pulumi.IntPtrInput    `pulumi:"deviceKey"`
	IpSettings  NicIPSettingsPtrInput `pulumi:"ipSettings"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
	NetworkId   pulumi.StringPtrInput `pulumi:"networkId"`
	NicType     pulumi.StringPtrInput `pulumi:"nicType"`
	PowerOnBoot pulumi.StringPtrInput `pulumi:"powerOnBoot"`
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArgs) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i NetworkInterfaceArgs) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}





type NetworkInterfaceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput
	ToNetworkInterfaceArrayOutputWithContext(context.Context) NetworkInterfaceArrayOutput
}

type NetworkInterfaceArray []NetworkInterfaceInput

func (NetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return i.ToNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceArrayOutput)
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) DeviceKey() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *int { return v.DeviceKey }).(pulumi.IntPtrOutput)
}

func (o NetworkInterfaceOutput) IpSettings() NicIPSettingsPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *NicIPSettings { return v.IpSettings }).(NicIPSettingsPtrOutput)
}

func (o NetworkInterfaceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.NetworkId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) NicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.NicType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) PowerOnBoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.PowerOnBoot }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterface {
		return vs[0].([]NetworkInterface)[vs[1].(int)]
	}).(NetworkInterfaceOutput)
}

type NetworkInterfaceResponse struct {
	DeviceKey      *int                   `pulumi:"deviceKey"`
	IpAddresses    []string               `pulumi:"ipAddresses"`
	IpSettings     *NicIPSettingsResponse `pulumi:"ipSettings"`
	Label          string                 `pulumi:"label"`
	MacAddress     string                 `pulumi:"macAddress"`
	Name           *string                `pulumi:"name"`
	NetworkId      *string                `pulumi:"networkId"`
	NetworkMoName  string                 `pulumi:"networkMoName"`
	NetworkMoRefId string                 `pulumi:"networkMoRefId"`
	NicType        *string                `pulumi:"nicType"`
	PowerOnBoot    *string                `pulumi:"powerOnBoot"`
}

type NetworkInterfaceResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutputWithContext(ctx context.Context) NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) DeviceKey() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *int { return v.DeviceKey }).(pulumi.IntPtrOutput)
}

func (o NetworkInterfaceResponseOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceResponseOutput) IpSettings() NicIPSettingsResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *NicIPSettingsResponse { return v.IpSettings }).(NicIPSettingsResponsePtrOutput)
}

func (o NetworkInterfaceResponseOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.Label }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.MacAddress }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.NetworkId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) NetworkMoName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.NetworkMoName }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) NetworkMoRefId() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.NetworkMoRefId }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) NicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.NicType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) PowerOnBoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.PowerOnBoot }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceResponse {
		return vs[0].([]NetworkInterfaceResponse)[vs[1].(int)]
	}).(NetworkInterfaceResponseOutput)
}

type NetworkProfile struct {
	NetworkInterfaces []NetworkInterface `pulumi:"networkInterfaces"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	NetworkInterfaces NetworkInterfaceArrayInput `pulumi:"networkInterfaces"`
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

func (o NetworkProfileOutput) NetworkInterfaces() NetworkInterfaceArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []NetworkInterface { return v.NetworkInterfaces }).(NetworkInterfaceArrayOutput)
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

func (o NetworkProfilePtrOutput) NetworkInterfaces() NetworkInterfaceArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []NetworkInterface {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceArrayOutput)
}

type NetworkProfileResponse struct {
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
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

func (o NetworkProfileResponseOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
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

func (o NetworkProfileResponsePtrOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []NetworkInterfaceResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceResponseArrayOutput)
}

type NicIPAddressSettingsResponse struct {
	AllocationMethod string `pulumi:"allocationMethod"`
	IpAddress        string `pulumi:"ipAddress"`
	SubnetMask       string `pulumi:"subnetMask"`
}

type NicIPAddressSettingsResponseOutput struct{ *pulumi.OutputState }

func (NicIPAddressSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIPAddressSettingsResponse)(nil)).Elem()
}

func (o NicIPAddressSettingsResponseOutput) ToNicIPAddressSettingsResponseOutput() NicIPAddressSettingsResponseOutput {
	return o
}

func (o NicIPAddressSettingsResponseOutput) ToNicIPAddressSettingsResponseOutputWithContext(ctx context.Context) NicIPAddressSettingsResponseOutput {
	return o
}

func (o NicIPAddressSettingsResponseOutput) AllocationMethod() pulumi.StringOutput {
	return o.ApplyT(func(v NicIPAddressSettingsResponse) string { return v.AllocationMethod }).(pulumi.StringOutput)
}

func (o NicIPAddressSettingsResponseOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v NicIPAddressSettingsResponse) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o NicIPAddressSettingsResponseOutput) SubnetMask() pulumi.StringOutput {
	return o.ApplyT(func(v NicIPAddressSettingsResponse) string { return v.SubnetMask }).(pulumi.StringOutput)
}

type NicIPAddressSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (NicIPAddressSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NicIPAddressSettingsResponse)(nil)).Elem()
}

func (o NicIPAddressSettingsResponseArrayOutput) ToNicIPAddressSettingsResponseArrayOutput() NicIPAddressSettingsResponseArrayOutput {
	return o
}

func (o NicIPAddressSettingsResponseArrayOutput) ToNicIPAddressSettingsResponseArrayOutputWithContext(ctx context.Context) NicIPAddressSettingsResponseArrayOutput {
	return o
}

func (o NicIPAddressSettingsResponseArrayOutput) Index(i pulumi.IntInput) NicIPAddressSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NicIPAddressSettingsResponse {
		return vs[0].([]NicIPAddressSettingsResponse)[vs[1].(int)]
	}).(NicIPAddressSettingsResponseOutput)
}

type NicIPSettings struct {
	AllocationMethod *string  `pulumi:"allocationMethod"`
	DnsServers       []string `pulumi:"dnsServers"`
	Gateway          []string `pulumi:"gateway"`
	IpAddress        *string  `pulumi:"ipAddress"`
	SubnetMask       *string  `pulumi:"subnetMask"`
}





type NicIPSettingsInput interface {
	pulumi.Input

	ToNicIPSettingsOutput() NicIPSettingsOutput
	ToNicIPSettingsOutputWithContext(context.Context) NicIPSettingsOutput
}

type NicIPSettingsArgs struct {
	AllocationMethod pulumi.StringPtrInput   `pulumi:"allocationMethod"`
	DnsServers       pulumi.StringArrayInput `pulumi:"dnsServers"`
	Gateway          pulumi.StringArrayInput `pulumi:"gateway"`
	IpAddress        pulumi.StringPtrInput   `pulumi:"ipAddress"`
	SubnetMask       pulumi.StringPtrInput   `pulumi:"subnetMask"`
}

func (NicIPSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIPSettings)(nil)).Elem()
}

func (i NicIPSettingsArgs) ToNicIPSettingsOutput() NicIPSettingsOutput {
	return i.ToNicIPSettingsOutputWithContext(context.Background())
}

func (i NicIPSettingsArgs) ToNicIPSettingsOutputWithContext(ctx context.Context) NicIPSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIPSettingsOutput)
}

func (i NicIPSettingsArgs) ToNicIPSettingsPtrOutput() NicIPSettingsPtrOutput {
	return i.ToNicIPSettingsPtrOutputWithContext(context.Background())
}

func (i NicIPSettingsArgs) ToNicIPSettingsPtrOutputWithContext(ctx context.Context) NicIPSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIPSettingsOutput).ToNicIPSettingsPtrOutputWithContext(ctx)
}









type NicIPSettingsPtrInput interface {
	pulumi.Input

	ToNicIPSettingsPtrOutput() NicIPSettingsPtrOutput
	ToNicIPSettingsPtrOutputWithContext(context.Context) NicIPSettingsPtrOutput
}

type nicIPSettingsPtrType NicIPSettingsArgs

func NicIPSettingsPtr(v *NicIPSettingsArgs) NicIPSettingsPtrInput {
	return (*nicIPSettingsPtrType)(v)
}

func (*nicIPSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NicIPSettings)(nil)).Elem()
}

func (i *nicIPSettingsPtrType) ToNicIPSettingsPtrOutput() NicIPSettingsPtrOutput {
	return i.ToNicIPSettingsPtrOutputWithContext(context.Background())
}

func (i *nicIPSettingsPtrType) ToNicIPSettingsPtrOutputWithContext(ctx context.Context) NicIPSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIPSettingsPtrOutput)
}

type NicIPSettingsOutput struct{ *pulumi.OutputState }

func (NicIPSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIPSettings)(nil)).Elem()
}

func (o NicIPSettingsOutput) ToNicIPSettingsOutput() NicIPSettingsOutput {
	return o
}

func (o NicIPSettingsOutput) ToNicIPSettingsOutputWithContext(ctx context.Context) NicIPSettingsOutput {
	return o
}

func (o NicIPSettingsOutput) ToNicIPSettingsPtrOutput() NicIPSettingsPtrOutput {
	return o.ToNicIPSettingsPtrOutputWithContext(context.Background())
}

func (o NicIPSettingsOutput) ToNicIPSettingsPtrOutputWithContext(ctx context.Context) NicIPSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NicIPSettings) *NicIPSettings {
		return &v
	}).(NicIPSettingsPtrOutput)
}

func (o NicIPSettingsOutput) AllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIPSettings) *string { return v.AllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NicIPSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o NicIPSettingsOutput) Gateway() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NicIPSettings) []string { return v.Gateway }).(pulumi.StringArrayOutput)
}

func (o NicIPSettingsOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIPSettings) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIPSettings) *string { return v.SubnetMask }).(pulumi.StringPtrOutput)
}

type NicIPSettingsPtrOutput struct{ *pulumi.OutputState }

func (NicIPSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NicIPSettings)(nil)).Elem()
}

func (o NicIPSettingsPtrOutput) ToNicIPSettingsPtrOutput() NicIPSettingsPtrOutput {
	return o
}

func (o NicIPSettingsPtrOutput) ToNicIPSettingsPtrOutputWithContext(ctx context.Context) NicIPSettingsPtrOutput {
	return o
}

func (o NicIPSettingsPtrOutput) Elem() NicIPSettingsOutput {
	return o.ApplyT(func(v *NicIPSettings) NicIPSettings {
		if v != nil {
			return *v
		}
		var ret NicIPSettings
		return ret
	}).(NicIPSettingsOutput)
}

func (o NicIPSettingsPtrOutput) AllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NicIPSettings) *string {
		if v == nil {
			return nil
		}
		return v.AllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NicIPSettings) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NicIPSettingsPtrOutput) Gateway() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NicIPSettings) []string {
		if v == nil {
			return nil
		}
		return v.Gateway
	}).(pulumi.StringArrayOutput)
}

func (o NicIPSettingsPtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NicIPSettings) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsPtrOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NicIPSettings) *string {
		if v == nil {
			return nil
		}
		return v.SubnetMask
	}).(pulumi.StringPtrOutput)
}

type NicIPSettingsResponse struct {
	AllocationMethod    *string                        `pulumi:"allocationMethod"`
	DnsServers          []string                       `pulumi:"dnsServers"`
	Gateway             []string                       `pulumi:"gateway"`
	IpAddress           *string                        `pulumi:"ipAddress"`
	IpAddressInfo       []NicIPAddressSettingsResponse `pulumi:"ipAddressInfo"`
	PrimaryWinsServer   string                         `pulumi:"primaryWinsServer"`
	SecondaryWinsServer string                         `pulumi:"secondaryWinsServer"`
	SubnetMask          *string                        `pulumi:"subnetMask"`
}

type NicIPSettingsResponseOutput struct{ *pulumi.OutputState }

func (NicIPSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIPSettingsResponse)(nil)).Elem()
}

func (o NicIPSettingsResponseOutput) ToNicIPSettingsResponseOutput() NicIPSettingsResponseOutput {
	return o
}

func (o NicIPSettingsResponseOutput) ToNicIPSettingsResponseOutputWithContext(ctx context.Context) NicIPSettingsResponseOutput {
	return o
}

func (o NicIPSettingsResponseOutput) AllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIPSettingsResponse) *string { return v.AllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NicIPSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o NicIPSettingsResponseOutput) Gateway() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NicIPSettingsResponse) []string { return v.Gateway }).(pulumi.StringArrayOutput)
}

func (o NicIPSettingsResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIPSettingsResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsResponseOutput) IpAddressInfo() NicIPAddressSettingsResponseArrayOutput {
	return o.ApplyT(func(v NicIPSettingsResponse) []NicIPAddressSettingsResponse { return v.IpAddressInfo }).(NicIPAddressSettingsResponseArrayOutput)
}

func (o NicIPSettingsResponseOutput) PrimaryWinsServer() pulumi.StringOutput {
	return o.ApplyT(func(v NicIPSettingsResponse) string { return v.PrimaryWinsServer }).(pulumi.StringOutput)
}

func (o NicIPSettingsResponseOutput) SecondaryWinsServer() pulumi.StringOutput {
	return o.ApplyT(func(v NicIPSettingsResponse) string { return v.SecondaryWinsServer }).(pulumi.StringOutput)
}

func (o NicIPSettingsResponseOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIPSettingsResponse) *string { return v.SubnetMask }).(pulumi.StringPtrOutput)
}

type NicIPSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (NicIPSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NicIPSettingsResponse)(nil)).Elem()
}

func (o NicIPSettingsResponsePtrOutput) ToNicIPSettingsResponsePtrOutput() NicIPSettingsResponsePtrOutput {
	return o
}

func (o NicIPSettingsResponsePtrOutput) ToNicIPSettingsResponsePtrOutputWithContext(ctx context.Context) NicIPSettingsResponsePtrOutput {
	return o
}

func (o NicIPSettingsResponsePtrOutput) Elem() NicIPSettingsResponseOutput {
	return o.ApplyT(func(v *NicIPSettingsResponse) NicIPSettingsResponse {
		if v != nil {
			return *v
		}
		var ret NicIPSettingsResponse
		return ret
	}).(NicIPSettingsResponseOutput)
}

func (o NicIPSettingsResponsePtrOutput) AllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NicIPSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NicIPSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NicIPSettingsResponsePtrOutput) Gateway() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NicIPSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Gateway
	}).(pulumi.StringArrayOutput)
}

func (o NicIPSettingsResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NicIPSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsResponsePtrOutput) IpAddressInfo() NicIPAddressSettingsResponseArrayOutput {
	return o.ApplyT(func(v *NicIPSettingsResponse) []NicIPAddressSettingsResponse {
		if v == nil {
			return nil
		}
		return v.IpAddressInfo
	}).(NicIPAddressSettingsResponseArrayOutput)
}

func (o NicIPSettingsResponsePtrOutput) PrimaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NicIPSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryWinsServer
	}).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsResponsePtrOutput) SecondaryWinsServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NicIPSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecondaryWinsServer
	}).(pulumi.StringPtrOutput)
}

func (o NicIPSettingsResponsePtrOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NicIPSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetMask
	}).(pulumi.StringPtrOutput)
}

type OsProfile struct {
	AdminPassword *string `pulumi:"adminPassword"`
	AdminUsername *string `pulumi:"adminUsername"`
	ComputerName  *string `pulumi:"computerName"`
	OsType        *string `pulumi:"osType"`
}





type OsProfileInput interface {
	pulumi.Input

	ToOsProfileOutput() OsProfileOutput
	ToOsProfileOutputWithContext(context.Context) OsProfileOutput
}

type OsProfileArgs struct {
	AdminPassword pulumi.StringPtrInput `pulumi:"adminPassword"`
	AdminUsername pulumi.StringPtrInput `pulumi:"adminUsername"`
	ComputerName  pulumi.StringPtrInput `pulumi:"computerName"`
	OsType        pulumi.StringPtrInput `pulumi:"osType"`
}

func (OsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfile)(nil)).Elem()
}

func (i OsProfileArgs) ToOsProfileOutput() OsProfileOutput {
	return i.ToOsProfileOutputWithContext(context.Background())
}

func (i OsProfileArgs) ToOsProfileOutputWithContext(ctx context.Context) OsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileOutput)
}

func (i OsProfileArgs) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return i.ToOsProfilePtrOutputWithContext(context.Background())
}

func (i OsProfileArgs) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileOutput).ToOsProfilePtrOutputWithContext(ctx)
}









type OsProfilePtrInput interface {
	pulumi.Input

	ToOsProfilePtrOutput() OsProfilePtrOutput
	ToOsProfilePtrOutputWithContext(context.Context) OsProfilePtrOutput
}

type osProfilePtrType OsProfileArgs

func OsProfilePtr(v *OsProfileArgs) OsProfilePtrInput {
	return (*osProfilePtrType)(v)
}

func (*osProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfile)(nil)).Elem()
}

func (i *osProfilePtrType) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return i.ToOsProfilePtrOutputWithContext(context.Background())
}

func (i *osProfilePtrType) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfilePtrOutput)
}

type OsProfileOutput struct{ *pulumi.OutputState }

func (OsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfile)(nil)).Elem()
}

func (o OsProfileOutput) ToOsProfileOutput() OsProfileOutput {
	return o
}

func (o OsProfileOutput) ToOsProfileOutputWithContext(ctx context.Context) OsProfileOutput {
	return o
}

func (o OsProfileOutput) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return o.ToOsProfilePtrOutputWithContext(context.Background())
}

func (o OsProfileOutput) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsProfile) *OsProfile {
		return &v
	}).(OsProfilePtrOutput)
}

func (o OsProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o OsProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OsProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OsProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

type OsProfilePtrOutput struct{ *pulumi.OutputState }

func (OsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfile)(nil)).Elem()
}

func (o OsProfilePtrOutput) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return o
}

func (o OsProfilePtrOutput) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return o
}

func (o OsProfilePtrOutput) Elem() OsProfileOutput {
	return o.ApplyT(func(v *OsProfile) OsProfile {
		if v != nil {
			return *v
		}
		var ret OsProfile
		return ret
	}).(OsProfileOutput)
}

func (o OsProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o OsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OsProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OsProfilePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfile) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

type OsProfileResponse struct {
	AdminUsername      *string `pulumi:"adminUsername"`
	ComputerName       *string `pulumi:"computerName"`
	OsName             string  `pulumi:"osName"`
	OsType             *string `pulumi:"osType"`
	ToolsRunningStatus string  `pulumi:"toolsRunningStatus"`
	ToolsVersion       string  `pulumi:"toolsVersion"`
	ToolsVersionStatus string  `pulumi:"toolsVersionStatus"`
}

type OsProfileResponseOutput struct{ *pulumi.OutputState }

func (OsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfileResponse)(nil)).Elem()
}

func (o OsProfileResponseOutput) ToOsProfileResponseOutput() OsProfileResponseOutput {
	return o
}

func (o OsProfileResponseOutput) ToOsProfileResponseOutputWithContext(ctx context.Context) OsProfileResponseOutput {
	return o
}

func (o OsProfileResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfileResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OsProfileResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfileResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OsProfileResponseOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v OsProfileResponse) string { return v.OsName }).(pulumi.StringOutput)
}

func (o OsProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OsProfileResponseOutput) ToolsRunningStatus() pulumi.StringOutput {
	return o.ApplyT(func(v OsProfileResponse) string { return v.ToolsRunningStatus }).(pulumi.StringOutput)
}

func (o OsProfileResponseOutput) ToolsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v OsProfileResponse) string { return v.ToolsVersion }).(pulumi.StringOutput)
}

func (o OsProfileResponseOutput) ToolsVersionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v OsProfileResponse) string { return v.ToolsVersionStatus }).(pulumi.StringOutput)
}

type OsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfileResponse)(nil)).Elem()
}

func (o OsProfileResponsePtrOutput) ToOsProfileResponsePtrOutput() OsProfileResponsePtrOutput {
	return o
}

func (o OsProfileResponsePtrOutput) ToOsProfileResponsePtrOutputWithContext(ctx context.Context) OsProfileResponsePtrOutput {
	return o
}

func (o OsProfileResponsePtrOutput) Elem() OsProfileResponseOutput {
	return o.ApplyT(func(v *OsProfileResponse) OsProfileResponse {
		if v != nil {
			return *v
		}
		var ret OsProfileResponse
		return ret
	}).(OsProfileResponseOutput)
}

func (o OsProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) OsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsName
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) ToolsRunningStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ToolsRunningStatus
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) ToolsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ToolsVersion
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) ToolsVersionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ToolsVersionStatus
	}).(pulumi.StringPtrOutput)
}

type PlacementProfile struct {
	ClusterId      *string `pulumi:"clusterId"`
	DatastoreId    *string `pulumi:"datastoreId"`
	HostId         *string `pulumi:"hostId"`
	ResourcePoolId *string `pulumi:"resourcePoolId"`
}





type PlacementProfileInput interface {
	pulumi.Input

	ToPlacementProfileOutput() PlacementProfileOutput
	ToPlacementProfileOutputWithContext(context.Context) PlacementProfileOutput
}

type PlacementProfileArgs struct {
	ClusterId      pulumi.StringPtrInput `pulumi:"clusterId"`
	DatastoreId    pulumi.StringPtrInput `pulumi:"datastoreId"`
	HostId         pulumi.StringPtrInput `pulumi:"hostId"`
	ResourcePoolId pulumi.StringPtrInput `pulumi:"resourcePoolId"`
}

func (PlacementProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementProfile)(nil)).Elem()
}

func (i PlacementProfileArgs) ToPlacementProfileOutput() PlacementProfileOutput {
	return i.ToPlacementProfileOutputWithContext(context.Background())
}

func (i PlacementProfileArgs) ToPlacementProfileOutputWithContext(ctx context.Context) PlacementProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementProfileOutput)
}

func (i PlacementProfileArgs) ToPlacementProfilePtrOutput() PlacementProfilePtrOutput {
	return i.ToPlacementProfilePtrOutputWithContext(context.Background())
}

func (i PlacementProfileArgs) ToPlacementProfilePtrOutputWithContext(ctx context.Context) PlacementProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementProfileOutput).ToPlacementProfilePtrOutputWithContext(ctx)
}









type PlacementProfilePtrInput interface {
	pulumi.Input

	ToPlacementProfilePtrOutput() PlacementProfilePtrOutput
	ToPlacementProfilePtrOutputWithContext(context.Context) PlacementProfilePtrOutput
}

type placementProfilePtrType PlacementProfileArgs

func PlacementProfilePtr(v *PlacementProfileArgs) PlacementProfilePtrInput {
	return (*placementProfilePtrType)(v)
}

func (*placementProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementProfile)(nil)).Elem()
}

func (i *placementProfilePtrType) ToPlacementProfilePtrOutput() PlacementProfilePtrOutput {
	return i.ToPlacementProfilePtrOutputWithContext(context.Background())
}

func (i *placementProfilePtrType) ToPlacementProfilePtrOutputWithContext(ctx context.Context) PlacementProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementProfilePtrOutput)
}

type PlacementProfileOutput struct{ *pulumi.OutputState }

func (PlacementProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementProfile)(nil)).Elem()
}

func (o PlacementProfileOutput) ToPlacementProfileOutput() PlacementProfileOutput {
	return o
}

func (o PlacementProfileOutput) ToPlacementProfileOutputWithContext(ctx context.Context) PlacementProfileOutput {
	return o
}

func (o PlacementProfileOutput) ToPlacementProfilePtrOutput() PlacementProfilePtrOutput {
	return o.ToPlacementProfilePtrOutputWithContext(context.Background())
}

func (o PlacementProfileOutput) ToPlacementProfilePtrOutputWithContext(ctx context.Context) PlacementProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PlacementProfile) *PlacementProfile {
		return &v
	}).(PlacementProfilePtrOutput)
}

func (o PlacementProfileOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlacementProfile) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o PlacementProfileOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlacementProfile) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o PlacementProfileOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlacementProfile) *string { return v.HostId }).(pulumi.StringPtrOutput)
}

func (o PlacementProfileOutput) ResourcePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlacementProfile) *string { return v.ResourcePoolId }).(pulumi.StringPtrOutput)
}

type PlacementProfilePtrOutput struct{ *pulumi.OutputState }

func (PlacementProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementProfile)(nil)).Elem()
}

func (o PlacementProfilePtrOutput) ToPlacementProfilePtrOutput() PlacementProfilePtrOutput {
	return o
}

func (o PlacementProfilePtrOutput) ToPlacementProfilePtrOutputWithContext(ctx context.Context) PlacementProfilePtrOutput {
	return o
}

func (o PlacementProfilePtrOutput) Elem() PlacementProfileOutput {
	return o.ApplyT(func(v *PlacementProfile) PlacementProfile {
		if v != nil {
			return *v
		}
		var ret PlacementProfile
		return ret
	}).(PlacementProfileOutput)
}

func (o PlacementProfilePtrOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementProfile) *string {
		if v == nil {
			return nil
		}
		return v.ClusterId
	}).(pulumi.StringPtrOutput)
}

func (o PlacementProfilePtrOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementProfile) *string {
		if v == nil {
			return nil
		}
		return v.DatastoreId
	}).(pulumi.StringPtrOutput)
}

func (o PlacementProfilePtrOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementProfile) *string {
		if v == nil {
			return nil
		}
		return v.HostId
	}).(pulumi.StringPtrOutput)
}

func (o PlacementProfilePtrOutput) ResourcePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementProfile) *string {
		if v == nil {
			return nil
		}
		return v.ResourcePoolId
	}).(pulumi.StringPtrOutput)
}

type PlacementProfileResponse struct {
	ClusterId      *string `pulumi:"clusterId"`
	DatastoreId    *string `pulumi:"datastoreId"`
	HostId         *string `pulumi:"hostId"`
	ResourcePoolId *string `pulumi:"resourcePoolId"`
}

type PlacementProfileResponseOutput struct{ *pulumi.OutputState }

func (PlacementProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementProfileResponse)(nil)).Elem()
}

func (o PlacementProfileResponseOutput) ToPlacementProfileResponseOutput() PlacementProfileResponseOutput {
	return o
}

func (o PlacementProfileResponseOutput) ToPlacementProfileResponseOutputWithContext(ctx context.Context) PlacementProfileResponseOutput {
	return o
}

func (o PlacementProfileResponseOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlacementProfileResponse) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o PlacementProfileResponseOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlacementProfileResponse) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o PlacementProfileResponseOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlacementProfileResponse) *string { return v.HostId }).(pulumi.StringPtrOutput)
}

func (o PlacementProfileResponseOutput) ResourcePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlacementProfileResponse) *string { return v.ResourcePoolId }).(pulumi.StringPtrOutput)
}

type PlacementProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (PlacementProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlacementProfileResponse)(nil)).Elem()
}

func (o PlacementProfileResponsePtrOutput) ToPlacementProfileResponsePtrOutput() PlacementProfileResponsePtrOutput {
	return o
}

func (o PlacementProfileResponsePtrOutput) ToPlacementProfileResponsePtrOutputWithContext(ctx context.Context) PlacementProfileResponsePtrOutput {
	return o
}

func (o PlacementProfileResponsePtrOutput) Elem() PlacementProfileResponseOutput {
	return o.ApplyT(func(v *PlacementProfileResponse) PlacementProfileResponse {
		if v != nil {
			return *v
		}
		var ret PlacementProfileResponse
		return ret
	}).(PlacementProfileResponseOutput)
}

func (o PlacementProfileResponsePtrOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterId
	}).(pulumi.StringPtrOutput)
}

func (o PlacementProfileResponsePtrOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DatastoreId
	}).(pulumi.StringPtrOutput)
}

func (o PlacementProfileResponsePtrOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostId
	}).(pulumi.StringPtrOutput)
}

func (o PlacementProfileResponsePtrOutput) ResourcePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlacementProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourcePoolId
	}).(pulumi.StringPtrOutput)
}

type ResourceStatusResponse struct {
	LastUpdatedAt string `pulumi:"lastUpdatedAt"`
	Message       string `pulumi:"message"`
	Reason        string `pulumi:"reason"`
	Severity      string `pulumi:"severity"`
	Status        string `pulumi:"status"`
	Type          string `pulumi:"type"`
}

type ResourceStatusResponseOutput struct{ *pulumi.OutputState }

func (ResourceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceStatusResponse)(nil)).Elem()
}

func (o ResourceStatusResponseOutput) ToResourceStatusResponseOutput() ResourceStatusResponseOutput {
	return o
}

func (o ResourceStatusResponseOutput) ToResourceStatusResponseOutputWithContext(ctx context.Context) ResourceStatusResponseOutput {
	return o
}

func (o ResourceStatusResponseOutput) LastUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceStatusResponse) string { return v.LastUpdatedAt }).(pulumi.StringOutput)
}

func (o ResourceStatusResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceStatusResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ResourceStatusResponseOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceStatusResponse) string { return v.Reason }).(pulumi.StringOutput)
}

func (o ResourceStatusResponseOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceStatusResponse) string { return v.Severity }).(pulumi.StringOutput)
}

func (o ResourceStatusResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceStatusResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o ResourceStatusResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceStatusResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ResourceStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceStatusResponse)(nil)).Elem()
}

func (o ResourceStatusResponseArrayOutput) ToResourceStatusResponseArrayOutput() ResourceStatusResponseArrayOutput {
	return o
}

func (o ResourceStatusResponseArrayOutput) ToResourceStatusResponseArrayOutputWithContext(ctx context.Context) ResourceStatusResponseArrayOutput {
	return o
}

func (o ResourceStatusResponseArrayOutput) Index(i pulumi.IntInput) ResourceStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceStatusResponse {
		return vs[0].([]ResourceStatusResponse)[vs[1].(int)]
	}).(ResourceStatusResponseOutput)
}

type StorageProfile struct {
	Disks []VirtualDisk `pulumi:"disks"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	Disks VirtualDiskArrayInput `pulumi:"disks"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) Disks() VirtualDiskArrayOutput {
	return o.ApplyT(func(v StorageProfile) []VirtualDisk { return v.Disks }).(VirtualDiskArrayOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) Disks() VirtualDiskArrayOutput {
	return o.ApplyT(func(v *StorageProfile) []VirtualDisk {
		if v == nil {
			return nil
		}
		return v.Disks
	}).(VirtualDiskArrayOutput)
}

type StorageProfileResponse struct {
	Disks           []VirtualDiskResponse           `pulumi:"disks"`
	ScsiControllers []VirtualSCSIControllerResponse `pulumi:"scsiControllers"`
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []VirtualDiskResponse { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

func (o StorageProfileResponseOutput) ScsiControllers() VirtualSCSIControllerResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []VirtualSCSIControllerResponse { return v.ScsiControllers }).(VirtualSCSIControllerResponseArrayOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []VirtualDiskResponse {
		if v == nil {
			return nil
		}
		return v.Disks
	}).(VirtualDiskResponseArrayOutput)
}

func (o StorageProfileResponsePtrOutput) ScsiControllers() VirtualSCSIControllerResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []VirtualSCSIControllerResponse {
		if v == nil {
			return nil
		}
		return v.ScsiControllers
	}).(VirtualSCSIControllerResponseArrayOutput)
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

type VICredential struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}





type VICredentialInput interface {
	pulumi.Input

	ToVICredentialOutput() VICredentialOutput
	ToVICredentialOutputWithContext(context.Context) VICredentialOutput
}

type VICredentialArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (VICredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VICredential)(nil)).Elem()
}

func (i VICredentialArgs) ToVICredentialOutput() VICredentialOutput {
	return i.ToVICredentialOutputWithContext(context.Background())
}

func (i VICredentialArgs) ToVICredentialOutputWithContext(ctx context.Context) VICredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VICredentialOutput)
}

func (i VICredentialArgs) ToVICredentialPtrOutput() VICredentialPtrOutput {
	return i.ToVICredentialPtrOutputWithContext(context.Background())
}

func (i VICredentialArgs) ToVICredentialPtrOutputWithContext(ctx context.Context) VICredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VICredentialOutput).ToVICredentialPtrOutputWithContext(ctx)
}









type VICredentialPtrInput interface {
	pulumi.Input

	ToVICredentialPtrOutput() VICredentialPtrOutput
	ToVICredentialPtrOutputWithContext(context.Context) VICredentialPtrOutput
}

type vicredentialPtrType VICredentialArgs

func VICredentialPtr(v *VICredentialArgs) VICredentialPtrInput {
	return (*vicredentialPtrType)(v)
}

func (*vicredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VICredential)(nil)).Elem()
}

func (i *vicredentialPtrType) ToVICredentialPtrOutput() VICredentialPtrOutput {
	return i.ToVICredentialPtrOutputWithContext(context.Background())
}

func (i *vicredentialPtrType) ToVICredentialPtrOutputWithContext(ctx context.Context) VICredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VICredentialPtrOutput)
}

type VICredentialOutput struct{ *pulumi.OutputState }

func (VICredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VICredential)(nil)).Elem()
}

func (o VICredentialOutput) ToVICredentialOutput() VICredentialOutput {
	return o
}

func (o VICredentialOutput) ToVICredentialOutputWithContext(ctx context.Context) VICredentialOutput {
	return o
}

func (o VICredentialOutput) ToVICredentialPtrOutput() VICredentialPtrOutput {
	return o.ToVICredentialPtrOutputWithContext(context.Background())
}

func (o VICredentialOutput) ToVICredentialPtrOutputWithContext(ctx context.Context) VICredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VICredential) *VICredential {
		return &v
	}).(VICredentialPtrOutput)
}

func (o VICredentialOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VICredential) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VICredentialOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VICredential) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type VICredentialPtrOutput struct{ *pulumi.OutputState }

func (VICredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VICredential)(nil)).Elem()
}

func (o VICredentialPtrOutput) ToVICredentialPtrOutput() VICredentialPtrOutput {
	return o
}

func (o VICredentialPtrOutput) ToVICredentialPtrOutputWithContext(ctx context.Context) VICredentialPtrOutput {
	return o
}

func (o VICredentialPtrOutput) Elem() VICredentialOutput {
	return o.ApplyT(func(v *VICredential) VICredential {
		if v != nil {
			return *v
		}
		var ret VICredential
		return ret
	}).(VICredentialOutput)
}

func (o VICredentialPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VICredential) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o VICredentialPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VICredential) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type VICredentialResponse struct {
	Username *string `pulumi:"username"`
}

type VICredentialResponseOutput struct{ *pulumi.OutputState }

func (VICredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VICredentialResponse)(nil)).Elem()
}

func (o VICredentialResponseOutput) ToVICredentialResponseOutput() VICredentialResponseOutput {
	return o
}

func (o VICredentialResponseOutput) ToVICredentialResponseOutputWithContext(ctx context.Context) VICredentialResponseOutput {
	return o
}

func (o VICredentialResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VICredentialResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type VICredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (VICredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VICredentialResponse)(nil)).Elem()
}

func (o VICredentialResponsePtrOutput) ToVICredentialResponsePtrOutput() VICredentialResponsePtrOutput {
	return o
}

func (o VICredentialResponsePtrOutput) ToVICredentialResponsePtrOutputWithContext(ctx context.Context) VICredentialResponsePtrOutput {
	return o
}

func (o VICredentialResponsePtrOutput) Elem() VICredentialResponseOutput {
	return o.ApplyT(func(v *VICredentialResponse) VICredentialResponse {
		if v != nil {
			return *v
		}
		var ret VICredentialResponse
		return ret
	}).(VICredentialResponseOutput)
}

func (o VICredentialResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VICredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type VirtualDisk struct {
	ControllerKey *int    `pulumi:"controllerKey"`
	DeviceKey     *int    `pulumi:"deviceKey"`
	DeviceName    *string `pulumi:"deviceName"`
	DiskMode      *string `pulumi:"diskMode"`
	DiskSizeGB    *int    `pulumi:"diskSizeGB"`
	DiskType      *string `pulumi:"diskType"`
	Name          *string `pulumi:"name"`
	UnitNumber    *int    `pulumi:"unitNumber"`
}





type VirtualDiskInput interface {
	pulumi.Input

	ToVirtualDiskOutput() VirtualDiskOutput
	ToVirtualDiskOutputWithContext(context.Context) VirtualDiskOutput
}

type VirtualDiskArgs struct {
	ControllerKey pulumi.IntPtrInput    `pulumi:"controllerKey"`
	DeviceKey     pulumi.IntPtrInput    `pulumi:"deviceKey"`
	DeviceName    pulumi.StringPtrInput `pulumi:"deviceName"`
	DiskMode      pulumi.StringPtrInput `pulumi:"diskMode"`
	DiskSizeGB    pulumi.IntPtrInput    `pulumi:"diskSizeGB"`
	DiskType      pulumi.StringPtrInput `pulumi:"diskType"`
	Name          pulumi.StringPtrInput `pulumi:"name"`
	UnitNumber    pulumi.IntPtrInput    `pulumi:"unitNumber"`
}

func (VirtualDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskArgs) ToVirtualDiskOutput() VirtualDiskOutput {
	return i.ToVirtualDiskOutputWithContext(context.Background())
}

func (i VirtualDiskArgs) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskOutput)
}





type VirtualDiskArrayInput interface {
	pulumi.Input

	ToVirtualDiskArrayOutput() VirtualDiskArrayOutput
	ToVirtualDiskArrayOutputWithContext(context.Context) VirtualDiskArrayOutput
}

type VirtualDiskArray []VirtualDiskInput

func (VirtualDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return i.ToVirtualDiskArrayOutputWithContext(context.Background())
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskArrayOutput)
}

type VirtualDiskOutput struct{ *pulumi.OutputState }

func (VirtualDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskOutput) ToVirtualDiskOutput() VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) ControllerKey() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *int { return v.ControllerKey }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskOutput) DeviceKey() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *int { return v.DeviceKey }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskOutput) DeviceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.DeviceName }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) DiskMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.DiskMode }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) UnitNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *int { return v.UnitNumber }).(pulumi.IntPtrOutput)
}

type VirtualDiskArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) Index(i pulumi.IntInput) VirtualDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDisk {
		return vs[0].([]VirtualDisk)[vs[1].(int)]
	}).(VirtualDiskOutput)
}

type VirtualDiskResponse struct {
	ControllerKey *int    `pulumi:"controllerKey"`
	DeviceKey     *int    `pulumi:"deviceKey"`
	DeviceName    *string `pulumi:"deviceName"`
	DiskMode      *string `pulumi:"diskMode"`
	DiskObjectId  string  `pulumi:"diskObjectId"`
	DiskSizeGB    *int    `pulumi:"diskSizeGB"`
	DiskType      *string `pulumi:"diskType"`
	Label         string  `pulumi:"label"`
	Name          *string `pulumi:"name"`
	UnitNumber    *int    `pulumi:"unitNumber"`
}

type VirtualDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDiskResponse)(nil)).Elem()
}

func (o VirtualDiskResponseOutput) ToVirtualDiskResponseOutput() VirtualDiskResponseOutput {
	return o
}

func (o VirtualDiskResponseOutput) ToVirtualDiskResponseOutputWithContext(ctx context.Context) VirtualDiskResponseOutput {
	return o
}

func (o VirtualDiskResponseOutput) ControllerKey() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *int { return v.ControllerKey }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskResponseOutput) DeviceKey() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *int { return v.DeviceKey }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskResponseOutput) DeviceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.DeviceName }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) DiskMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.DiskMode }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) DiskObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskResponse) string { return v.DiskObjectId }).(pulumi.StringOutput)
}

func (o VirtualDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskResponseOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskResponse) string { return v.Label }).(pulumi.StringOutput)
}

func (o VirtualDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) UnitNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *int { return v.UnitNumber }).(pulumi.IntPtrOutput)
}

type VirtualDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDiskResponse)(nil)).Elem()
}

func (o VirtualDiskResponseArrayOutput) ToVirtualDiskResponseArrayOutput() VirtualDiskResponseArrayOutput {
	return o
}

func (o VirtualDiskResponseArrayOutput) ToVirtualDiskResponseArrayOutputWithContext(ctx context.Context) VirtualDiskResponseArrayOutput {
	return o
}

func (o VirtualDiskResponseArrayOutput) Index(i pulumi.IntInput) VirtualDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDiskResponse {
		return vs[0].([]VirtualDiskResponse)[vs[1].(int)]
	}).(VirtualDiskResponseOutput)
}

type VirtualSCSIControllerResponse struct {
	BusNumber          *int    `pulumi:"busNumber"`
	ControllerKey      *int    `pulumi:"controllerKey"`
	ScsiCtlrUnitNumber *int    `pulumi:"scsiCtlrUnitNumber"`
	Sharing            *string `pulumi:"sharing"`
	Type               *string `pulumi:"type"`
}

type VirtualSCSIControllerResponseOutput struct{ *pulumi.OutputState }

func (VirtualSCSIControllerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualSCSIControllerResponse)(nil)).Elem()
}

func (o VirtualSCSIControllerResponseOutput) ToVirtualSCSIControllerResponseOutput() VirtualSCSIControllerResponseOutput {
	return o
}

func (o VirtualSCSIControllerResponseOutput) ToVirtualSCSIControllerResponseOutputWithContext(ctx context.Context) VirtualSCSIControllerResponseOutput {
	return o
}

func (o VirtualSCSIControllerResponseOutput) BusNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualSCSIControllerResponse) *int { return v.BusNumber }).(pulumi.IntPtrOutput)
}

func (o VirtualSCSIControllerResponseOutput) ControllerKey() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualSCSIControllerResponse) *int { return v.ControllerKey }).(pulumi.IntPtrOutput)
}

func (o VirtualSCSIControllerResponseOutput) ScsiCtlrUnitNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualSCSIControllerResponse) *int { return v.ScsiCtlrUnitNumber }).(pulumi.IntPtrOutput)
}

func (o VirtualSCSIControllerResponseOutput) Sharing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualSCSIControllerResponse) *string { return v.Sharing }).(pulumi.StringPtrOutput)
}

func (o VirtualSCSIControllerResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualSCSIControllerResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualSCSIControllerResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualSCSIControllerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualSCSIControllerResponse)(nil)).Elem()
}

func (o VirtualSCSIControllerResponseArrayOutput) ToVirtualSCSIControllerResponseArrayOutput() VirtualSCSIControllerResponseArrayOutput {
	return o
}

func (o VirtualSCSIControllerResponseArrayOutput) ToVirtualSCSIControllerResponseArrayOutputWithContext(ctx context.Context) VirtualSCSIControllerResponseArrayOutput {
	return o
}

func (o VirtualSCSIControllerResponseArrayOutput) Index(i pulumi.IntInput) VirtualSCSIControllerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualSCSIControllerResponse {
		return vs[0].([]VirtualSCSIControllerResponse)[vs[1].(int)]
	}).(VirtualSCSIControllerResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestAgentProfileResponseOutput{})
	pulumi.RegisterOutputType(GuestAgentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestCredentialOutput{})
	pulumi.RegisterOutputType(GuestCredentialPtrOutput{})
	pulumi.RegisterOutputType(GuestCredentialResponseOutput{})
	pulumi.RegisterOutputType(GuestCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileOutput{})
	pulumi.RegisterOutputType(HardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponseOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigurationOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigurationPtrOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigurationResponseOutput{})
	pulumi.RegisterOutputType(HttpProxyConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewResponseStatusPtrOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponseInstanceViewOutput{})
	pulumi.RegisterOutputType(MachineExtensionPropertiesResponseInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(NicIPAddressSettingsResponseOutput{})
	pulumi.RegisterOutputType(NicIPAddressSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(NicIPSettingsOutput{})
	pulumi.RegisterOutputType(NicIPSettingsPtrOutput{})
	pulumi.RegisterOutputType(NicIPSettingsResponseOutput{})
	pulumi.RegisterOutputType(NicIPSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(OsProfileOutput{})
	pulumi.RegisterOutputType(OsProfilePtrOutput{})
	pulumi.RegisterOutputType(OsProfileResponseOutput{})
	pulumi.RegisterOutputType(OsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(PlacementProfileOutput{})
	pulumi.RegisterOutputType(PlacementProfilePtrOutput{})
	pulumi.RegisterOutputType(PlacementProfileResponseOutput{})
	pulumi.RegisterOutputType(PlacementProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceStatusResponseOutput{})
	pulumi.RegisterOutputType(ResourceStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VICredentialOutput{})
	pulumi.RegisterOutputType(VICredentialPtrOutput{})
	pulumi.RegisterOutputType(VICredentialResponseOutput{})
	pulumi.RegisterOutputType(VICredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualDiskOutput{})
	pulumi.RegisterOutputType(VirtualDiskArrayOutput{})
	pulumi.RegisterOutputType(VirtualDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualSCSIControllerResponseOutput{})
	pulumi.RegisterOutputType(VirtualSCSIControllerResponseArrayOutput{})
}
