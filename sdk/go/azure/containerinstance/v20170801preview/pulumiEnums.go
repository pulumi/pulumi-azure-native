


package v20170801preview

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

type ContainerRestartPolicy string

const (
	ContainerRestartPolicyAlways = ContainerRestartPolicy("always")
)

func (ContainerRestartPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRestartPolicy)(nil)).Elem()
}

func (e ContainerRestartPolicy) ToContainerRestartPolicyOutput() ContainerRestartPolicyOutput {
	return pulumi.ToOutput(e).(ContainerRestartPolicyOutput)
}

func (e ContainerRestartPolicy) ToContainerRestartPolicyOutputWithContext(ctx context.Context) ContainerRestartPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerRestartPolicyOutput)
}

func (e ContainerRestartPolicy) ToContainerRestartPolicyPtrOutput() ContainerRestartPolicyPtrOutput {
	return e.ToContainerRestartPolicyPtrOutputWithContext(context.Background())
}

func (e ContainerRestartPolicy) ToContainerRestartPolicyPtrOutputWithContext(ctx context.Context) ContainerRestartPolicyPtrOutput {
	return ContainerRestartPolicy(e).ToContainerRestartPolicyOutputWithContext(ctx).ToContainerRestartPolicyPtrOutputWithContext(ctx)
}

func (e ContainerRestartPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerRestartPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerRestartPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerRestartPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerRestartPolicyOutput struct{ *pulumi.OutputState }

func (ContainerRestartPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRestartPolicy)(nil)).Elem()
}

func (o ContainerRestartPolicyOutput) ToContainerRestartPolicyOutput() ContainerRestartPolicyOutput {
	return o
}

func (o ContainerRestartPolicyOutput) ToContainerRestartPolicyOutputWithContext(ctx context.Context) ContainerRestartPolicyOutput {
	return o
}

func (o ContainerRestartPolicyOutput) ToContainerRestartPolicyPtrOutput() ContainerRestartPolicyPtrOutput {
	return o.ToContainerRestartPolicyPtrOutputWithContext(context.Background())
}

func (o ContainerRestartPolicyOutput) ToContainerRestartPolicyPtrOutputWithContext(ctx context.Context) ContainerRestartPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerRestartPolicy) *ContainerRestartPolicy {
		return &v
	}).(ContainerRestartPolicyPtrOutput)
}

func (o ContainerRestartPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerRestartPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerRestartPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerRestartPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerRestartPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerRestartPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerRestartPolicyPtrOutput struct{ *pulumi.OutputState }

func (ContainerRestartPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRestartPolicy)(nil)).Elem()
}

func (o ContainerRestartPolicyPtrOutput) ToContainerRestartPolicyPtrOutput() ContainerRestartPolicyPtrOutput {
	return o
}

func (o ContainerRestartPolicyPtrOutput) ToContainerRestartPolicyPtrOutputWithContext(ctx context.Context) ContainerRestartPolicyPtrOutput {
	return o
}

func (o ContainerRestartPolicyPtrOutput) Elem() ContainerRestartPolicyOutput {
	return o.ApplyT(func(v *ContainerRestartPolicy) ContainerRestartPolicy {
		if v != nil {
			return *v
		}
		var ret ContainerRestartPolicy
		return ret
	}).(ContainerRestartPolicyOutput)
}

func (o ContainerRestartPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerRestartPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerRestartPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContainerRestartPolicyInput interface {
	pulumi.Input

	ToContainerRestartPolicyOutput() ContainerRestartPolicyOutput
	ToContainerRestartPolicyOutputWithContext(context.Context) ContainerRestartPolicyOutput
}

var containerRestartPolicyPtrType = reflect.TypeOf((**ContainerRestartPolicy)(nil)).Elem()

type ContainerRestartPolicyPtrInput interface {
	pulumi.Input

	ToContainerRestartPolicyPtrOutput() ContainerRestartPolicyPtrOutput
	ToContainerRestartPolicyPtrOutputWithContext(context.Context) ContainerRestartPolicyPtrOutput
}

type containerRestartPolicyPtr string

func ContainerRestartPolicyPtr(v string) ContainerRestartPolicyPtrInput {
	return (*containerRestartPolicyPtr)(&v)
}

func (*containerRestartPolicyPtr) ElementType() reflect.Type {
	return containerRestartPolicyPtrType
}

func (in *containerRestartPolicyPtr) ToContainerRestartPolicyPtrOutput() ContainerRestartPolicyPtrOutput {
	return pulumi.ToOutput(in).(ContainerRestartPolicyPtrOutput)
}

func (in *containerRestartPolicyPtr) ToContainerRestartPolicyPtrOutputWithContext(ctx context.Context) ContainerRestartPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerRestartPolicyPtrOutput)
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
	pulumi.RegisterOutputType(ContainerGroupIpAddressTypeOutput{})
	pulumi.RegisterOutputType(ContainerGroupIpAddressTypePtrOutput{})
	pulumi.RegisterOutputType(ContainerGroupNetworkProtocolOutput{})
	pulumi.RegisterOutputType(ContainerGroupNetworkProtocolPtrOutput{})
	pulumi.RegisterOutputType(ContainerRestartPolicyOutput{})
	pulumi.RegisterOutputType(ContainerRestartPolicyPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesPtrOutput{})
}
