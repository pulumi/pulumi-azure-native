


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerGroupIpAddressType string

const (
	ContainerGroupIpAddressTypePublic = ContainerGroupIpAddressType("Public")
)

func (ContainerGroupIpAddressType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupIpAddressType)(nil)).Elem()
}

func (e ContainerGroupIpAddressType) ToContainerGroupIpAddressTypeOutput() ContainerGroupIpAddressTypeOutput {
	return pulumi.ToOutput(e).(ContainerGroupIpAddressTypeOutput)
}

func (e ContainerGroupIpAddressType) ToContainerGroupIpAddressTypeOutputWithContext(ctx context.Context) ContainerGroupIpAddressTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerGroupIpAddressTypeOutput)
}

func (e ContainerGroupIpAddressType) ToContainerGroupIpAddressTypePtrOutput() ContainerGroupIpAddressTypePtrOutput {
	return e.ToContainerGroupIpAddressTypePtrOutputWithContext(context.Background())
}

func (e ContainerGroupIpAddressType) ToContainerGroupIpAddressTypePtrOutputWithContext(ctx context.Context) ContainerGroupIpAddressTypePtrOutput {
	return ContainerGroupIpAddressType(e).ToContainerGroupIpAddressTypeOutputWithContext(ctx).ToContainerGroupIpAddressTypePtrOutputWithContext(ctx)
}

func (e ContainerGroupIpAddressType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerGroupIpAddressType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerGroupIpAddressType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerGroupIpAddressType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerGroupIpAddressTypeOutput struct{ *pulumi.OutputState }

func (ContainerGroupIpAddressTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupIpAddressType)(nil)).Elem()
}

func (o ContainerGroupIpAddressTypeOutput) ToContainerGroupIpAddressTypeOutput() ContainerGroupIpAddressTypeOutput {
	return o
}

func (o ContainerGroupIpAddressTypeOutput) ToContainerGroupIpAddressTypeOutputWithContext(ctx context.Context) ContainerGroupIpAddressTypeOutput {
	return o
}

func (o ContainerGroupIpAddressTypeOutput) ToContainerGroupIpAddressTypePtrOutput() ContainerGroupIpAddressTypePtrOutput {
	return o.ToContainerGroupIpAddressTypePtrOutputWithContext(context.Background())
}

func (o ContainerGroupIpAddressTypeOutput) ToContainerGroupIpAddressTypePtrOutputWithContext(ctx context.Context) ContainerGroupIpAddressTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerGroupIpAddressType) *ContainerGroupIpAddressType {
		return &v
	}).(ContainerGroupIpAddressTypePtrOutput)
}

func (o ContainerGroupIpAddressTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerGroupIpAddressTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerGroupIpAddressType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerGroupIpAddressTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerGroupIpAddressTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerGroupIpAddressType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerGroupIpAddressTypePtrOutput struct{ *pulumi.OutputState }

func (ContainerGroupIpAddressTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupIpAddressType)(nil)).Elem()
}

func (o ContainerGroupIpAddressTypePtrOutput) ToContainerGroupIpAddressTypePtrOutput() ContainerGroupIpAddressTypePtrOutput {
	return o
}

func (o ContainerGroupIpAddressTypePtrOutput) ToContainerGroupIpAddressTypePtrOutputWithContext(ctx context.Context) ContainerGroupIpAddressTypePtrOutput {
	return o
}

func (o ContainerGroupIpAddressTypePtrOutput) Elem() ContainerGroupIpAddressTypeOutput {
	return o.ApplyT(func(v *ContainerGroupIpAddressType) ContainerGroupIpAddressType {
		if v != nil {
			return *v
		}
		var ret ContainerGroupIpAddressType
		return ret
	}).(ContainerGroupIpAddressTypeOutput)
}

func (o ContainerGroupIpAddressTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerGroupIpAddressTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerGroupIpAddressType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContainerGroupIpAddressTypeInput interface {
	pulumi.Input

	ToContainerGroupIpAddressTypeOutput() ContainerGroupIpAddressTypeOutput
	ToContainerGroupIpAddressTypeOutputWithContext(context.Context) ContainerGroupIpAddressTypeOutput
}

var containerGroupIpAddressTypePtrType = reflect.TypeOf((**ContainerGroupIpAddressType)(nil)).Elem()

type ContainerGroupIpAddressTypePtrInput interface {
	pulumi.Input

	ToContainerGroupIpAddressTypePtrOutput() ContainerGroupIpAddressTypePtrOutput
	ToContainerGroupIpAddressTypePtrOutputWithContext(context.Context) ContainerGroupIpAddressTypePtrOutput
}

type containerGroupIpAddressTypePtr string

func ContainerGroupIpAddressTypePtr(v string) ContainerGroupIpAddressTypePtrInput {
	return (*containerGroupIpAddressTypePtr)(&v)
}

func (*containerGroupIpAddressTypePtr) ElementType() reflect.Type {
	return containerGroupIpAddressTypePtrType
}

func (in *containerGroupIpAddressTypePtr) ToContainerGroupIpAddressTypePtrOutput() ContainerGroupIpAddressTypePtrOutput {
	return pulumi.ToOutput(in).(ContainerGroupIpAddressTypePtrOutput)
}

func (in *containerGroupIpAddressTypePtr) ToContainerGroupIpAddressTypePtrOutputWithContext(ctx context.Context) ContainerGroupIpAddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerGroupIpAddressTypePtrOutput)
}

type ContainerGroupNetworkProtocol string

const (
	ContainerGroupNetworkProtocolTCP = ContainerGroupNetworkProtocol("TCP")
	ContainerGroupNetworkProtocolUDP = ContainerGroupNetworkProtocol("UDP")
)

func (ContainerGroupNetworkProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupNetworkProtocol)(nil)).Elem()
}

func (e ContainerGroupNetworkProtocol) ToContainerGroupNetworkProtocolOutput() ContainerGroupNetworkProtocolOutput {
	return pulumi.ToOutput(e).(ContainerGroupNetworkProtocolOutput)
}

func (e ContainerGroupNetworkProtocol) ToContainerGroupNetworkProtocolOutputWithContext(ctx context.Context) ContainerGroupNetworkProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerGroupNetworkProtocolOutput)
}

func (e ContainerGroupNetworkProtocol) ToContainerGroupNetworkProtocolPtrOutput() ContainerGroupNetworkProtocolPtrOutput {
	return e.ToContainerGroupNetworkProtocolPtrOutputWithContext(context.Background())
}

func (e ContainerGroupNetworkProtocol) ToContainerGroupNetworkProtocolPtrOutputWithContext(ctx context.Context) ContainerGroupNetworkProtocolPtrOutput {
	return ContainerGroupNetworkProtocol(e).ToContainerGroupNetworkProtocolOutputWithContext(ctx).ToContainerGroupNetworkProtocolPtrOutputWithContext(ctx)
}

func (e ContainerGroupNetworkProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerGroupNetworkProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerGroupNetworkProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerGroupNetworkProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerGroupNetworkProtocolOutput struct{ *pulumi.OutputState }

func (ContainerGroupNetworkProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupNetworkProtocol)(nil)).Elem()
}

func (o ContainerGroupNetworkProtocolOutput) ToContainerGroupNetworkProtocolOutput() ContainerGroupNetworkProtocolOutput {
	return o
}

func (o ContainerGroupNetworkProtocolOutput) ToContainerGroupNetworkProtocolOutputWithContext(ctx context.Context) ContainerGroupNetworkProtocolOutput {
	return o
}

func (o ContainerGroupNetworkProtocolOutput) ToContainerGroupNetworkProtocolPtrOutput() ContainerGroupNetworkProtocolPtrOutput {
	return o.ToContainerGroupNetworkProtocolPtrOutputWithContext(context.Background())
}

func (o ContainerGroupNetworkProtocolOutput) ToContainerGroupNetworkProtocolPtrOutputWithContext(ctx context.Context) ContainerGroupNetworkProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerGroupNetworkProtocol) *ContainerGroupNetworkProtocol {
		return &v
	}).(ContainerGroupNetworkProtocolPtrOutput)
}

func (o ContainerGroupNetworkProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerGroupNetworkProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerGroupNetworkProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerGroupNetworkProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerGroupNetworkProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerGroupNetworkProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerGroupNetworkProtocolPtrOutput struct{ *pulumi.OutputState }

func (ContainerGroupNetworkProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupNetworkProtocol)(nil)).Elem()
}

func (o ContainerGroupNetworkProtocolPtrOutput) ToContainerGroupNetworkProtocolPtrOutput() ContainerGroupNetworkProtocolPtrOutput {
	return o
}

func (o ContainerGroupNetworkProtocolPtrOutput) ToContainerGroupNetworkProtocolPtrOutputWithContext(ctx context.Context) ContainerGroupNetworkProtocolPtrOutput {
	return o
}

func (o ContainerGroupNetworkProtocolPtrOutput) Elem() ContainerGroupNetworkProtocolOutput {
	return o.ApplyT(func(v *ContainerGroupNetworkProtocol) ContainerGroupNetworkProtocol {
		if v != nil {
			return *v
		}
		var ret ContainerGroupNetworkProtocol
		return ret
	}).(ContainerGroupNetworkProtocolOutput)
}

func (o ContainerGroupNetworkProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerGroupNetworkProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerGroupNetworkProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContainerGroupNetworkProtocolInput interface {
	pulumi.Input

	ToContainerGroupNetworkProtocolOutput() ContainerGroupNetworkProtocolOutput
	ToContainerGroupNetworkProtocolOutputWithContext(context.Context) ContainerGroupNetworkProtocolOutput
}

var containerGroupNetworkProtocolPtrType = reflect.TypeOf((**ContainerGroupNetworkProtocol)(nil)).Elem()

type ContainerGroupNetworkProtocolPtrInput interface {
	pulumi.Input

	ToContainerGroupNetworkProtocolPtrOutput() ContainerGroupNetworkProtocolPtrOutput
	ToContainerGroupNetworkProtocolPtrOutputWithContext(context.Context) ContainerGroupNetworkProtocolPtrOutput
}

type containerGroupNetworkProtocolPtr string

func ContainerGroupNetworkProtocolPtr(v string) ContainerGroupNetworkProtocolPtrInput {
	return (*containerGroupNetworkProtocolPtr)(&v)
}

func (*containerGroupNetworkProtocolPtr) ElementType() reflect.Type {
	return containerGroupNetworkProtocolPtrType
}

func (in *containerGroupNetworkProtocolPtr) ToContainerGroupNetworkProtocolPtrOutput() ContainerGroupNetworkProtocolPtrOutput {
	return pulumi.ToOutput(in).(ContainerGroupNetworkProtocolPtrOutput)
}

func (in *containerGroupNetworkProtocolPtr) ToContainerGroupNetworkProtocolPtrOutputWithContext(ctx context.Context) ContainerGroupNetworkProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerGroupNetworkProtocolPtrOutput)
}

type ContainerGroupRestartPolicy string

const (
	ContainerGroupRestartPolicyAlways    = ContainerGroupRestartPolicy("Always")
	ContainerGroupRestartPolicyOnFailure = ContainerGroupRestartPolicy("OnFailure")
	ContainerGroupRestartPolicyNever     = ContainerGroupRestartPolicy("Never")
)

func (ContainerGroupRestartPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupRestartPolicy)(nil)).Elem()
}

func (e ContainerGroupRestartPolicy) ToContainerGroupRestartPolicyOutput() ContainerGroupRestartPolicyOutput {
	return pulumi.ToOutput(e).(ContainerGroupRestartPolicyOutput)
}

func (e ContainerGroupRestartPolicy) ToContainerGroupRestartPolicyOutputWithContext(ctx context.Context) ContainerGroupRestartPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerGroupRestartPolicyOutput)
}

func (e ContainerGroupRestartPolicy) ToContainerGroupRestartPolicyPtrOutput() ContainerGroupRestartPolicyPtrOutput {
	return e.ToContainerGroupRestartPolicyPtrOutputWithContext(context.Background())
}

func (e ContainerGroupRestartPolicy) ToContainerGroupRestartPolicyPtrOutputWithContext(ctx context.Context) ContainerGroupRestartPolicyPtrOutput {
	return ContainerGroupRestartPolicy(e).ToContainerGroupRestartPolicyOutputWithContext(ctx).ToContainerGroupRestartPolicyPtrOutputWithContext(ctx)
}

func (e ContainerGroupRestartPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerGroupRestartPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerGroupRestartPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerGroupRestartPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerGroupRestartPolicyOutput struct{ *pulumi.OutputState }

func (ContainerGroupRestartPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerGroupRestartPolicy)(nil)).Elem()
}

func (o ContainerGroupRestartPolicyOutput) ToContainerGroupRestartPolicyOutput() ContainerGroupRestartPolicyOutput {
	return o
}

func (o ContainerGroupRestartPolicyOutput) ToContainerGroupRestartPolicyOutputWithContext(ctx context.Context) ContainerGroupRestartPolicyOutput {
	return o
}

func (o ContainerGroupRestartPolicyOutput) ToContainerGroupRestartPolicyPtrOutput() ContainerGroupRestartPolicyPtrOutput {
	return o.ToContainerGroupRestartPolicyPtrOutputWithContext(context.Background())
}

func (o ContainerGroupRestartPolicyOutput) ToContainerGroupRestartPolicyPtrOutputWithContext(ctx context.Context) ContainerGroupRestartPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerGroupRestartPolicy) *ContainerGroupRestartPolicy {
		return &v
	}).(ContainerGroupRestartPolicyPtrOutput)
}

func (o ContainerGroupRestartPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerGroupRestartPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerGroupRestartPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerGroupRestartPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerGroupRestartPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerGroupRestartPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerGroupRestartPolicyPtrOutput struct{ *pulumi.OutputState }

func (ContainerGroupRestartPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroupRestartPolicy)(nil)).Elem()
}

func (o ContainerGroupRestartPolicyPtrOutput) ToContainerGroupRestartPolicyPtrOutput() ContainerGroupRestartPolicyPtrOutput {
	return o
}

func (o ContainerGroupRestartPolicyPtrOutput) ToContainerGroupRestartPolicyPtrOutputWithContext(ctx context.Context) ContainerGroupRestartPolicyPtrOutput {
	return o
}

func (o ContainerGroupRestartPolicyPtrOutput) Elem() ContainerGroupRestartPolicyOutput {
	return o.ApplyT(func(v *ContainerGroupRestartPolicy) ContainerGroupRestartPolicy {
		if v != nil {
			return *v
		}
		var ret ContainerGroupRestartPolicy
		return ret
	}).(ContainerGroupRestartPolicyOutput)
}

func (o ContainerGroupRestartPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerGroupRestartPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerGroupRestartPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContainerGroupRestartPolicyInput interface {
	pulumi.Input

	ToContainerGroupRestartPolicyOutput() ContainerGroupRestartPolicyOutput
	ToContainerGroupRestartPolicyOutputWithContext(context.Context) ContainerGroupRestartPolicyOutput
}

var containerGroupRestartPolicyPtrType = reflect.TypeOf((**ContainerGroupRestartPolicy)(nil)).Elem()

type ContainerGroupRestartPolicyPtrInput interface {
	pulumi.Input

	ToContainerGroupRestartPolicyPtrOutput() ContainerGroupRestartPolicyPtrOutput
	ToContainerGroupRestartPolicyPtrOutputWithContext(context.Context) ContainerGroupRestartPolicyPtrOutput
}

type containerGroupRestartPolicyPtr string

func ContainerGroupRestartPolicyPtr(v string) ContainerGroupRestartPolicyPtrInput {
	return (*containerGroupRestartPolicyPtr)(&v)
}

func (*containerGroupRestartPolicyPtr) ElementType() reflect.Type {
	return containerGroupRestartPolicyPtrType
}

func (in *containerGroupRestartPolicyPtr) ToContainerGroupRestartPolicyPtrOutput() ContainerGroupRestartPolicyPtrOutput {
	return pulumi.ToOutput(in).(ContainerGroupRestartPolicyPtrOutput)
}

func (in *containerGroupRestartPolicyPtr) ToContainerGroupRestartPolicyPtrOutputWithContext(ctx context.Context) ContainerGroupRestartPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerGroupRestartPolicyPtrOutput)
}

type ContainerNetworkProtocol string

const (
	ContainerNetworkProtocolTCP = ContainerNetworkProtocol("TCP")
	ContainerNetworkProtocolUDP = ContainerNetworkProtocol("UDP")
)

func (ContainerNetworkProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerNetworkProtocol)(nil)).Elem()
}

func (e ContainerNetworkProtocol) ToContainerNetworkProtocolOutput() ContainerNetworkProtocolOutput {
	return pulumi.ToOutput(e).(ContainerNetworkProtocolOutput)
}

func (e ContainerNetworkProtocol) ToContainerNetworkProtocolOutputWithContext(ctx context.Context) ContainerNetworkProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerNetworkProtocolOutput)
}

func (e ContainerNetworkProtocol) ToContainerNetworkProtocolPtrOutput() ContainerNetworkProtocolPtrOutput {
	return e.ToContainerNetworkProtocolPtrOutputWithContext(context.Background())
}

func (e ContainerNetworkProtocol) ToContainerNetworkProtocolPtrOutputWithContext(ctx context.Context) ContainerNetworkProtocolPtrOutput {
	return ContainerNetworkProtocol(e).ToContainerNetworkProtocolOutputWithContext(ctx).ToContainerNetworkProtocolPtrOutputWithContext(ctx)
}

func (e ContainerNetworkProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerNetworkProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerNetworkProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerNetworkProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerNetworkProtocolOutput struct{ *pulumi.OutputState }

func (ContainerNetworkProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerNetworkProtocol)(nil)).Elem()
}

func (o ContainerNetworkProtocolOutput) ToContainerNetworkProtocolOutput() ContainerNetworkProtocolOutput {
	return o
}

func (o ContainerNetworkProtocolOutput) ToContainerNetworkProtocolOutputWithContext(ctx context.Context) ContainerNetworkProtocolOutput {
	return o
}

func (o ContainerNetworkProtocolOutput) ToContainerNetworkProtocolPtrOutput() ContainerNetworkProtocolPtrOutput {
	return o.ToContainerNetworkProtocolPtrOutputWithContext(context.Background())
}

func (o ContainerNetworkProtocolOutput) ToContainerNetworkProtocolPtrOutputWithContext(ctx context.Context) ContainerNetworkProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerNetworkProtocol) *ContainerNetworkProtocol {
		return &v
	}).(ContainerNetworkProtocolPtrOutput)
}

func (o ContainerNetworkProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerNetworkProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerNetworkProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerNetworkProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerNetworkProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerNetworkProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerNetworkProtocolPtrOutput struct{ *pulumi.OutputState }

func (ContainerNetworkProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerNetworkProtocol)(nil)).Elem()
}

func (o ContainerNetworkProtocolPtrOutput) ToContainerNetworkProtocolPtrOutput() ContainerNetworkProtocolPtrOutput {
	return o
}

func (o ContainerNetworkProtocolPtrOutput) ToContainerNetworkProtocolPtrOutputWithContext(ctx context.Context) ContainerNetworkProtocolPtrOutput {
	return o
}

func (o ContainerNetworkProtocolPtrOutput) Elem() ContainerNetworkProtocolOutput {
	return o.ApplyT(func(v *ContainerNetworkProtocol) ContainerNetworkProtocol {
		if v != nil {
			return *v
		}
		var ret ContainerNetworkProtocol
		return ret
	}).(ContainerNetworkProtocolOutput)
}

func (o ContainerNetworkProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerNetworkProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerNetworkProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContainerNetworkProtocolInput interface {
	pulumi.Input

	ToContainerNetworkProtocolOutput() ContainerNetworkProtocolOutput
	ToContainerNetworkProtocolOutputWithContext(context.Context) ContainerNetworkProtocolOutput
}

var containerNetworkProtocolPtrType = reflect.TypeOf((**ContainerNetworkProtocol)(nil)).Elem()

type ContainerNetworkProtocolPtrInput interface {
	pulumi.Input

	ToContainerNetworkProtocolPtrOutput() ContainerNetworkProtocolPtrOutput
	ToContainerNetworkProtocolPtrOutputWithContext(context.Context) ContainerNetworkProtocolPtrOutput
}

type containerNetworkProtocolPtr string

func ContainerNetworkProtocolPtr(v string) ContainerNetworkProtocolPtrInput {
	return (*containerNetworkProtocolPtr)(&v)
}

func (*containerNetworkProtocolPtr) ElementType() reflect.Type {
	return containerNetworkProtocolPtrType
}

func (in *containerNetworkProtocolPtr) ToContainerNetworkProtocolPtrOutput() ContainerNetworkProtocolPtrOutput {
	return pulumi.ToOutput(in).(ContainerNetworkProtocolPtrOutput)
}

func (in *containerNetworkProtocolPtr) ToContainerNetworkProtocolPtrOutputWithContext(ctx context.Context) ContainerNetworkProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerNetworkProtocolPtrOutput)
}

type OperatingSystemTypes string

const (
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
)

func (OperatingSystemTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemTypes)(nil)).Elem()
}

func (e OperatingSystemTypes) ToOperatingSystemTypesOutput() OperatingSystemTypesOutput {
	return pulumi.ToOutput(e).(OperatingSystemTypesOutput)
}

func (e OperatingSystemTypes) ToOperatingSystemTypesOutputWithContext(ctx context.Context) OperatingSystemTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatingSystemTypesOutput)
}

func (e OperatingSystemTypes) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return e.ToOperatingSystemTypesPtrOutputWithContext(context.Background())
}

func (e OperatingSystemTypes) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return OperatingSystemTypes(e).ToOperatingSystemTypesOutputWithContext(ctx).ToOperatingSystemTypesPtrOutputWithContext(ctx)
}

func (e OperatingSystemTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatingSystemTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatingSystemTypesOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemTypes)(nil)).Elem()
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesOutput() OperatingSystemTypesOutput {
	return o
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesOutputWithContext(ctx context.Context) OperatingSystemTypesOutput {
	return o
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return o.ToOperatingSystemTypesPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatingSystemTypes) *OperatingSystemTypes {
		return &v
	}).(OperatingSystemTypesPtrOutput)
}

func (o OperatingSystemTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatingSystemTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatingSystemTypesPtrOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatingSystemTypes)(nil)).Elem()
}

func (o OperatingSystemTypesPtrOutput) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return o
}

func (o OperatingSystemTypesPtrOutput) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return o
}

func (o OperatingSystemTypesPtrOutput) Elem() OperatingSystemTypesOutput {
	return o.ApplyT(func(v *OperatingSystemTypes) OperatingSystemTypes {
		if v != nil {
			return *v
		}
		var ret OperatingSystemTypes
		return ret
	}).(OperatingSystemTypesOutput)
}

func (o OperatingSystemTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatingSystemTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatingSystemTypesInput interface {
	pulumi.Input

	ToOperatingSystemTypesOutput() OperatingSystemTypesOutput
	ToOperatingSystemTypesOutputWithContext(context.Context) OperatingSystemTypesOutput
}

var operatingSystemTypesPtrType = reflect.TypeOf((**OperatingSystemTypes)(nil)).Elem()

type OperatingSystemTypesPtrInput interface {
	pulumi.Input

	ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput
	ToOperatingSystemTypesPtrOutputWithContext(context.Context) OperatingSystemTypesPtrOutput
}

type operatingSystemTypesPtr string

func OperatingSystemTypesPtr(v string) OperatingSystemTypesPtrInput {
	return (*operatingSystemTypesPtr)(&v)
}

func (*operatingSystemTypesPtr) ElementType() reflect.Type {
	return operatingSystemTypesPtrType
}

func (in *operatingSystemTypesPtr) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return pulumi.ToOutput(in).(OperatingSystemTypesPtrOutput)
}

func (in *operatingSystemTypesPtr) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatingSystemTypesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupIpAddressTypeInput)(nil)).Elem(), ContainerGroupIpAddressType("Public"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupIpAddressTypePtrInput)(nil)).Elem(), ContainerGroupIpAddressType("Public"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupNetworkProtocolInput)(nil)).Elem(), ContainerGroupNetworkProtocol("TCP"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupNetworkProtocolPtrInput)(nil)).Elem(), ContainerGroupNetworkProtocol("TCP"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupRestartPolicyInput)(nil)).Elem(), ContainerGroupRestartPolicy("Always"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerGroupRestartPolicyPtrInput)(nil)).Elem(), ContainerGroupRestartPolicy("Always"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerNetworkProtocolInput)(nil)).Elem(), ContainerNetworkProtocol("TCP"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerNetworkProtocolPtrInput)(nil)).Elem(), ContainerNetworkProtocol("TCP"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatingSystemTypesInput)(nil)).Elem(), OperatingSystemTypes("Windows"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatingSystemTypesPtrInput)(nil)).Elem(), OperatingSystemTypes("Windows"))
	pulumi.RegisterOutputType(ContainerGroupIpAddressTypeOutput{})
	pulumi.RegisterOutputType(ContainerGroupIpAddressTypePtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupNetworkProtocolOutput{})
	pulumi.RegisterOutputType(ContainerGroupNetworkProtocolPtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupRestartPolicyOutput{})
	pulumi.RegisterOutputType(ContainerGroupRestartPolicyPtrOutput{})
	pulumi.RegisterOutputType(ContainerNetworkProtocolOutput{})
	pulumi.RegisterOutputType(ContainerNetworkProtocolPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesPtrOutput{})
}
