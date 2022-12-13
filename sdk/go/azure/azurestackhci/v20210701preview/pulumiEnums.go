


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExtendedLocationTypes string

const (
	ExtendedLocationTypesCustomLocation = ExtendedLocationTypes("CustomLocation")
)

type IPPoolTypeEnum string

const (
	IPPoolTypeEnumVm      = IPPoolTypeEnum("vm")
	IPPoolTypeEnumVippool = IPPoolTypeEnum("vippool")
)

func (IPPoolTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPoolTypeEnum)(nil)).Elem()
}

func (e IPPoolTypeEnum) ToIPPoolTypeEnumOutput() IPPoolTypeEnumOutput {
	return pulumi.ToOutput(e).(IPPoolTypeEnumOutput)
}

func (e IPPoolTypeEnum) ToIPPoolTypeEnumOutputWithContext(ctx context.Context) IPPoolTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IPPoolTypeEnumOutput)
}

func (e IPPoolTypeEnum) ToIPPoolTypeEnumPtrOutput() IPPoolTypeEnumPtrOutput {
	return e.ToIPPoolTypeEnumPtrOutputWithContext(context.Background())
}

func (e IPPoolTypeEnum) ToIPPoolTypeEnumPtrOutputWithContext(ctx context.Context) IPPoolTypeEnumPtrOutput {
	return IPPoolTypeEnum(e).ToIPPoolTypeEnumOutputWithContext(ctx).ToIPPoolTypeEnumPtrOutputWithContext(ctx)
}

func (e IPPoolTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPPoolTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPPoolTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IPPoolTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IPPoolTypeEnumOutput struct{ *pulumi.OutputState }

func (IPPoolTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPPoolTypeEnum)(nil)).Elem()
}

func (o IPPoolTypeEnumOutput) ToIPPoolTypeEnumOutput() IPPoolTypeEnumOutput {
	return o
}

func (o IPPoolTypeEnumOutput) ToIPPoolTypeEnumOutputWithContext(ctx context.Context) IPPoolTypeEnumOutput {
	return o
}

func (o IPPoolTypeEnumOutput) ToIPPoolTypeEnumPtrOutput() IPPoolTypeEnumPtrOutput {
	return o.ToIPPoolTypeEnumPtrOutputWithContext(context.Background())
}

func (o IPPoolTypeEnumOutput) ToIPPoolTypeEnumPtrOutputWithContext(ctx context.Context) IPPoolTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPPoolTypeEnum) *IPPoolTypeEnum {
		return &v
	}).(IPPoolTypeEnumPtrOutput)
}

func (o IPPoolTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IPPoolTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPPoolTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IPPoolTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPPoolTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPPoolTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IPPoolTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (IPPoolTypeEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPPoolTypeEnum)(nil)).Elem()
}

func (o IPPoolTypeEnumPtrOutput) ToIPPoolTypeEnumPtrOutput() IPPoolTypeEnumPtrOutput {
	return o
}

func (o IPPoolTypeEnumPtrOutput) ToIPPoolTypeEnumPtrOutputWithContext(ctx context.Context) IPPoolTypeEnumPtrOutput {
	return o
}

func (o IPPoolTypeEnumPtrOutput) Elem() IPPoolTypeEnumOutput {
	return o.ApplyT(func(v *IPPoolTypeEnum) IPPoolTypeEnum {
		if v != nil {
			return *v
		}
		var ret IPPoolTypeEnum
		return ret
	}).(IPPoolTypeEnumOutput)
}

func (o IPPoolTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPPoolTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IPPoolTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IPPoolTypeEnumInput interface {
	pulumi.Input

	ToIPPoolTypeEnumOutput() IPPoolTypeEnumOutput
	ToIPPoolTypeEnumOutputWithContext(context.Context) IPPoolTypeEnumOutput
}

var ippoolTypeEnumPtrType = reflect.TypeOf((**IPPoolTypeEnum)(nil)).Elem()

type IPPoolTypeEnumPtrInput interface {
	pulumi.Input

	ToIPPoolTypeEnumPtrOutput() IPPoolTypeEnumPtrOutput
	ToIPPoolTypeEnumPtrOutputWithContext(context.Context) IPPoolTypeEnumPtrOutput
}

type ippoolTypeEnumPtr string

func IPPoolTypeEnumPtr(v string) IPPoolTypeEnumPtrInput {
	return (*ippoolTypeEnumPtr)(&v)
}

func (*ippoolTypeEnumPtr) ElementType() reflect.Type {
	return ippoolTypeEnumPtrType
}

func (in *ippoolTypeEnumPtr) ToIPPoolTypeEnumPtrOutput() IPPoolTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(IPPoolTypeEnumPtrOutput)
}

func (in *ippoolTypeEnumPtr) ToIPPoolTypeEnumPtrOutputWithContext(ctx context.Context) IPPoolTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IPPoolTypeEnumPtrOutput)
}

type IpAllocationMethodEnum string

const (
	IpAllocationMethodEnumDynamic = IpAllocationMethodEnum("Dynamic")
	IpAllocationMethodEnumStatic  = IpAllocationMethodEnum("Static")
)

type NetworkTypeEnum string

const (
	NetworkTypeEnumNAT         = NetworkTypeEnum("NAT")
	NetworkTypeEnumTransparent = NetworkTypeEnum("Transparent")
	NetworkTypeEnumL2Bridge    = NetworkTypeEnum("L2Bridge")
	NetworkTypeEnumL2Tunnel    = NetworkTypeEnum("L2Tunnel")
	NetworkTypeEnumICS         = NetworkTypeEnum("ICS")
	NetworkTypeEnumPrivate     = NetworkTypeEnum("Private")
	NetworkTypeEnumOverlay     = NetworkTypeEnum("Overlay")
	NetworkTypeEnumInternal    = NetworkTypeEnum("Internal")
	NetworkTypeEnumMirrored    = NetworkTypeEnum("Mirrored")
)

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

type OsTypeEnum string

const (
	OsTypeEnumLinux   = OsTypeEnum("Linux")
	OsTypeEnumWindows = OsTypeEnum("Windows")
)

type PrivateIPAllocationMethodEnum string

const (
	PrivateIPAllocationMethodEnumDynamic = PrivateIPAllocationMethodEnum("Dynamic")
	PrivateIPAllocationMethodEnumStatic  = PrivateIPAllocationMethodEnum("Static")
)

func init() {
	pulumi.RegisterOutputType(IPPoolTypeEnumOutput{})
	pulumi.RegisterOutputType(IPPoolTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesPtrOutput{})
}
