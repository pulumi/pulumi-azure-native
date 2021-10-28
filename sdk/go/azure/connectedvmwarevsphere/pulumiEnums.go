


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiskMode string

const (
	DiskModePersistent                 = DiskMode("persistent")
	DiskMode_Independent_persistent    = DiskMode("independent_persistent")
	DiskMode_Independent_nonpersistent = DiskMode("independent_nonpersistent")
)

func (DiskMode) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskMode)(nil)).Elem()
}

func (e DiskMode) ToDiskModeOutput() DiskModeOutput {
	return pulumi.ToOutput(e).(DiskModeOutput)
}

func (e DiskMode) ToDiskModeOutputWithContext(ctx context.Context) DiskModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskModeOutput)
}

func (e DiskMode) ToDiskModePtrOutput() DiskModePtrOutput {
	return e.ToDiskModePtrOutputWithContext(context.Background())
}

func (e DiskMode) ToDiskModePtrOutputWithContext(ctx context.Context) DiskModePtrOutput {
	return DiskMode(e).ToDiskModeOutputWithContext(ctx).ToDiskModePtrOutputWithContext(ctx)
}

func (e DiskMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskModeOutput struct{ *pulumi.OutputState }

func (DiskModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskMode)(nil)).Elem()
}

func (o DiskModeOutput) ToDiskModeOutput() DiskModeOutput {
	return o
}

func (o DiskModeOutput) ToDiskModeOutputWithContext(ctx context.Context) DiskModeOutput {
	return o
}

func (o DiskModeOutput) ToDiskModePtrOutput() DiskModePtrOutput {
	return o.ToDiskModePtrOutputWithContext(context.Background())
}

func (o DiskModeOutput) ToDiskModePtrOutputWithContext(ctx context.Context) DiskModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskMode) *DiskMode {
		return &v
	}).(DiskModePtrOutput)
}

func (o DiskModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskModePtrOutput struct{ *pulumi.OutputState }

func (DiskModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskMode)(nil)).Elem()
}

func (o DiskModePtrOutput) ToDiskModePtrOutput() DiskModePtrOutput {
	return o
}

func (o DiskModePtrOutput) ToDiskModePtrOutputWithContext(ctx context.Context) DiskModePtrOutput {
	return o
}

func (o DiskModePtrOutput) Elem() DiskModeOutput {
	return o.ApplyT(func(v *DiskMode) DiskMode {
		if v != nil {
			return *v
		}
		var ret DiskMode
		return ret
	}).(DiskModeOutput)
}

func (o DiskModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskModeInput interface {
	pulumi.Input

	ToDiskModeOutput() DiskModeOutput
	ToDiskModeOutputWithContext(context.Context) DiskModeOutput
}

var diskModePtrType = reflect.TypeOf((**DiskMode)(nil)).Elem()

type DiskModePtrInput interface {
	pulumi.Input

	ToDiskModePtrOutput() DiskModePtrOutput
	ToDiskModePtrOutputWithContext(context.Context) DiskModePtrOutput
}

type diskModePtr string

func DiskModePtr(v string) DiskModePtrInput {
	return (*diskModePtr)(&v)
}

func (*diskModePtr) ElementType() reflect.Type {
	return diskModePtrType
}

func (in *diskModePtr) ToDiskModePtrOutput() DiskModePtrOutput {
	return pulumi.ToOutput(in).(DiskModePtrOutput)
}

func (in *diskModePtr) ToDiskModePtrOutputWithContext(ctx context.Context) DiskModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskModePtrOutput)
}

type DiskType string

const (
	DiskTypeFlat        = DiskType("flat")
	DiskTypePmem        = DiskType("pmem")
	DiskTypeRawphysical = DiskType("rawphysical")
	DiskTypeRawvirtual  = DiskType("rawvirtual")
	DiskTypeSparse      = DiskType("sparse")
	DiskTypeSesparse    = DiskType("sesparse")
	DiskTypeUnknown     = DiskType("unknown")
)

func (DiskType) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskType)(nil)).Elem()
}

func (e DiskType) ToDiskTypeOutput() DiskTypeOutput {
	return pulumi.ToOutput(e).(DiskTypeOutput)
}

func (e DiskType) ToDiskTypeOutputWithContext(ctx context.Context) DiskTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskTypeOutput)
}

func (e DiskType) ToDiskTypePtrOutput() DiskTypePtrOutput {
	return e.ToDiskTypePtrOutputWithContext(context.Background())
}

func (e DiskType) ToDiskTypePtrOutputWithContext(ctx context.Context) DiskTypePtrOutput {
	return DiskType(e).ToDiskTypeOutputWithContext(ctx).ToDiskTypePtrOutputWithContext(ctx)
}

func (e DiskType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskTypeOutput struct{ *pulumi.OutputState }

func (DiskTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskType)(nil)).Elem()
}

func (o DiskTypeOutput) ToDiskTypeOutput() DiskTypeOutput {
	return o
}

func (o DiskTypeOutput) ToDiskTypeOutputWithContext(ctx context.Context) DiskTypeOutput {
	return o
}

func (o DiskTypeOutput) ToDiskTypePtrOutput() DiskTypePtrOutput {
	return o.ToDiskTypePtrOutputWithContext(context.Background())
}

func (o DiskTypeOutput) ToDiskTypePtrOutputWithContext(ctx context.Context) DiskTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskType) *DiskType {
		return &v
	}).(DiskTypePtrOutput)
}

func (o DiskTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskTypePtrOutput struct{ *pulumi.OutputState }

func (DiskTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskType)(nil)).Elem()
}

func (o DiskTypePtrOutput) ToDiskTypePtrOutput() DiskTypePtrOutput {
	return o
}

func (o DiskTypePtrOutput) ToDiskTypePtrOutputWithContext(ctx context.Context) DiskTypePtrOutput {
	return o
}

func (o DiskTypePtrOutput) Elem() DiskTypeOutput {
	return o.ApplyT(func(v *DiskType) DiskType {
		if v != nil {
			return *v
		}
		var ret DiskType
		return ret
	}).(DiskTypeOutput)
}

func (o DiskTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskTypeInput interface {
	pulumi.Input

	ToDiskTypeOutput() DiskTypeOutput
	ToDiskTypeOutputWithContext(context.Context) DiskTypeOutput
}

var diskTypePtrType = reflect.TypeOf((**DiskType)(nil)).Elem()

type DiskTypePtrInput interface {
	pulumi.Input

	ToDiskTypePtrOutput() DiskTypePtrOutput
	ToDiskTypePtrOutputWithContext(context.Context) DiskTypePtrOutput
}

type diskTypePtr string

func DiskTypePtr(v string) DiskTypePtrInput {
	return (*diskTypePtr)(&v)
}

func (*diskTypePtr) ElementType() reflect.Type {
	return diskTypePtrType
}

func (in *diskTypePtr) ToDiskTypePtrOutput() DiskTypePtrOutput {
	return pulumi.ToOutput(in).(DiskTypePtrOutput)
}

func (in *diskTypePtr) ToDiskTypePtrOutputWithContext(ctx context.Context) DiskTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskTypePtrOutput)
}

type FirmwareType string

const (
	FirmwareTypeBios = FirmwareType("bios")
	FirmwareTypeEfi  = FirmwareType("efi")
)

func (FirmwareType) ElementType() reflect.Type {
	return reflect.TypeOf((*FirmwareType)(nil)).Elem()
}

func (e FirmwareType) ToFirmwareTypeOutput() FirmwareTypeOutput {
	return pulumi.ToOutput(e).(FirmwareTypeOutput)
}

func (e FirmwareType) ToFirmwareTypeOutputWithContext(ctx context.Context) FirmwareTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FirmwareTypeOutput)
}

func (e FirmwareType) ToFirmwareTypePtrOutput() FirmwareTypePtrOutput {
	return e.ToFirmwareTypePtrOutputWithContext(context.Background())
}

func (e FirmwareType) ToFirmwareTypePtrOutputWithContext(ctx context.Context) FirmwareTypePtrOutput {
	return FirmwareType(e).ToFirmwareTypeOutputWithContext(ctx).ToFirmwareTypePtrOutputWithContext(ctx)
}

func (e FirmwareType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirmwareType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FirmwareType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FirmwareType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FirmwareTypeOutput struct{ *pulumi.OutputState }

func (FirmwareTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirmwareType)(nil)).Elem()
}

func (o FirmwareTypeOutput) ToFirmwareTypeOutput() FirmwareTypeOutput {
	return o
}

func (o FirmwareTypeOutput) ToFirmwareTypeOutputWithContext(ctx context.Context) FirmwareTypeOutput {
	return o
}

func (o FirmwareTypeOutput) ToFirmwareTypePtrOutput() FirmwareTypePtrOutput {
	return o.ToFirmwareTypePtrOutputWithContext(context.Background())
}

func (o FirmwareTypeOutput) ToFirmwareTypePtrOutputWithContext(ctx context.Context) FirmwareTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FirmwareType) *FirmwareType {
		return &v
	}).(FirmwareTypePtrOutput)
}

func (o FirmwareTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FirmwareTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirmwareType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FirmwareTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirmwareTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FirmwareType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FirmwareTypePtrOutput struct{ *pulumi.OutputState }

func (FirmwareTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirmwareType)(nil)).Elem()
}

func (o FirmwareTypePtrOutput) ToFirmwareTypePtrOutput() FirmwareTypePtrOutput {
	return o
}

func (o FirmwareTypePtrOutput) ToFirmwareTypePtrOutputWithContext(ctx context.Context) FirmwareTypePtrOutput {
	return o
}

func (o FirmwareTypePtrOutput) Elem() FirmwareTypeOutput {
	return o.ApplyT(func(v *FirmwareType) FirmwareType {
		if v != nil {
			return *v
		}
		var ret FirmwareType
		return ret
	}).(FirmwareTypeOutput)
}

func (o FirmwareTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FirmwareTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FirmwareType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FirmwareTypeInput interface {
	pulumi.Input

	ToFirmwareTypeOutput() FirmwareTypeOutput
	ToFirmwareTypeOutputWithContext(context.Context) FirmwareTypeOutput
}

var firmwareTypePtrType = reflect.TypeOf((**FirmwareType)(nil)).Elem()

type FirmwareTypePtrInput interface {
	pulumi.Input

	ToFirmwareTypePtrOutput() FirmwareTypePtrOutput
	ToFirmwareTypePtrOutputWithContext(context.Context) FirmwareTypePtrOutput
}

type firmwareTypePtr string

func FirmwareTypePtr(v string) FirmwareTypePtrInput {
	return (*firmwareTypePtr)(&v)
}

func (*firmwareTypePtr) ElementType() reflect.Type {
	return firmwareTypePtrType
}

func (in *firmwareTypePtr) ToFirmwareTypePtrOutput() FirmwareTypePtrOutput {
	return pulumi.ToOutput(in).(FirmwareTypePtrOutput)
}

func (in *firmwareTypePtr) ToFirmwareTypePtrOutputWithContext(ctx context.Context) FirmwareTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FirmwareTypePtrOutput)
}

type IPAddressAllocationMethod string

const (
	IPAddressAllocationMethodUnset     = IPAddressAllocationMethod("unset")
	IPAddressAllocationMethodDynamic   = IPAddressAllocationMethod("dynamic")
	IPAddressAllocationMethodStatic    = IPAddressAllocationMethod("static")
	IPAddressAllocationMethodLinklayer = IPAddressAllocationMethod("linklayer")
	IPAddressAllocationMethodRandom    = IPAddressAllocationMethod("random")
	IPAddressAllocationMethodOther     = IPAddressAllocationMethod("other")
)

func (IPAddressAllocationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAddressAllocationMethod)(nil)).Elem()
}

func (e IPAddressAllocationMethod) ToIPAddressAllocationMethodOutput() IPAddressAllocationMethodOutput {
	return pulumi.ToOutput(e).(IPAddressAllocationMethodOutput)
}

func (e IPAddressAllocationMethod) ToIPAddressAllocationMethodOutputWithContext(ctx context.Context) IPAddressAllocationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IPAddressAllocationMethodOutput)
}

func (e IPAddressAllocationMethod) ToIPAddressAllocationMethodPtrOutput() IPAddressAllocationMethodPtrOutput {
	return e.ToIPAddressAllocationMethodPtrOutputWithContext(context.Background())
}

func (e IPAddressAllocationMethod) ToIPAddressAllocationMethodPtrOutputWithContext(ctx context.Context) IPAddressAllocationMethodPtrOutput {
	return IPAddressAllocationMethod(e).ToIPAddressAllocationMethodOutputWithContext(ctx).ToIPAddressAllocationMethodPtrOutputWithContext(ctx)
}

func (e IPAddressAllocationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAddressAllocationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAddressAllocationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IPAddressAllocationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IPAddressAllocationMethodOutput struct{ *pulumi.OutputState }

func (IPAddressAllocationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAddressAllocationMethod)(nil)).Elem()
}

func (o IPAddressAllocationMethodOutput) ToIPAddressAllocationMethodOutput() IPAddressAllocationMethodOutput {
	return o
}

func (o IPAddressAllocationMethodOutput) ToIPAddressAllocationMethodOutputWithContext(ctx context.Context) IPAddressAllocationMethodOutput {
	return o
}

func (o IPAddressAllocationMethodOutput) ToIPAddressAllocationMethodPtrOutput() IPAddressAllocationMethodPtrOutput {
	return o.ToIPAddressAllocationMethodPtrOutputWithContext(context.Background())
}

func (o IPAddressAllocationMethodOutput) ToIPAddressAllocationMethodPtrOutputWithContext(ctx context.Context) IPAddressAllocationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAddressAllocationMethod) *IPAddressAllocationMethod {
		return &v
	}).(IPAddressAllocationMethodPtrOutput)
}

func (o IPAddressAllocationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IPAddressAllocationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAddressAllocationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IPAddressAllocationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPAddressAllocationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAddressAllocationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IPAddressAllocationMethodPtrOutput struct{ *pulumi.OutputState }

func (IPAddressAllocationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAddressAllocationMethod)(nil)).Elem()
}

func (o IPAddressAllocationMethodPtrOutput) ToIPAddressAllocationMethodPtrOutput() IPAddressAllocationMethodPtrOutput {
	return o
}

func (o IPAddressAllocationMethodPtrOutput) ToIPAddressAllocationMethodPtrOutputWithContext(ctx context.Context) IPAddressAllocationMethodPtrOutput {
	return o
}

func (o IPAddressAllocationMethodPtrOutput) Elem() IPAddressAllocationMethodOutput {
	return o.ApplyT(func(v *IPAddressAllocationMethod) IPAddressAllocationMethod {
		if v != nil {
			return *v
		}
		var ret IPAddressAllocationMethod
		return ret
	}).(IPAddressAllocationMethodOutput)
}

func (o IPAddressAllocationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPAddressAllocationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IPAddressAllocationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IPAddressAllocationMethodInput interface {
	pulumi.Input

	ToIPAddressAllocationMethodOutput() IPAddressAllocationMethodOutput
	ToIPAddressAllocationMethodOutputWithContext(context.Context) IPAddressAllocationMethodOutput
}

var ipaddressAllocationMethodPtrType = reflect.TypeOf((**IPAddressAllocationMethod)(nil)).Elem()

type IPAddressAllocationMethodPtrInput interface {
	pulumi.Input

	ToIPAddressAllocationMethodPtrOutput() IPAddressAllocationMethodPtrOutput
	ToIPAddressAllocationMethodPtrOutputWithContext(context.Context) IPAddressAllocationMethodPtrOutput
}

type ipaddressAllocationMethodPtr string

func IPAddressAllocationMethodPtr(v string) IPAddressAllocationMethodPtrInput {
	return (*ipaddressAllocationMethodPtr)(&v)
}

func (*ipaddressAllocationMethodPtr) ElementType() reflect.Type {
	return ipaddressAllocationMethodPtrType
}

func (in *ipaddressAllocationMethodPtr) ToIPAddressAllocationMethodPtrOutput() IPAddressAllocationMethodPtrOutput {
	return pulumi.ToOutput(in).(IPAddressAllocationMethodPtrOutput)
}

func (in *ipaddressAllocationMethodPtr) ToIPAddressAllocationMethodPtrOutputWithContext(ctx context.Context) IPAddressAllocationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IPAddressAllocationMethodPtrOutput)
}

type IdentityType string

const (
	IdentityTypeNone           = IdentityType("None")
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

type InventoryType string

const (
	InventoryTypeResourcePool           = InventoryType("ResourcePool")
	InventoryTypeVirtualMachine         = InventoryType("VirtualMachine")
	InventoryTypeVirtualMachineTemplate = InventoryType("VirtualMachineTemplate")
	InventoryTypeVirtualNetwork         = InventoryType("VirtualNetwork")
	InventoryTypeCluster                = InventoryType("Cluster")
	InventoryTypeDatastore              = InventoryType("Datastore")
	InventoryTypeHost                   = InventoryType("Host")
)

func (InventoryType) ElementType() reflect.Type {
	return reflect.TypeOf((*InventoryType)(nil)).Elem()
}

func (e InventoryType) ToInventoryTypeOutput() InventoryTypeOutput {
	return pulumi.ToOutput(e).(InventoryTypeOutput)
}

func (e InventoryType) ToInventoryTypeOutputWithContext(ctx context.Context) InventoryTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InventoryTypeOutput)
}

func (e InventoryType) ToInventoryTypePtrOutput() InventoryTypePtrOutput {
	return e.ToInventoryTypePtrOutputWithContext(context.Background())
}

func (e InventoryType) ToInventoryTypePtrOutputWithContext(ctx context.Context) InventoryTypePtrOutput {
	return InventoryType(e).ToInventoryTypeOutputWithContext(ctx).ToInventoryTypePtrOutputWithContext(ctx)
}

func (e InventoryType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InventoryType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InventoryType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InventoryType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InventoryTypeOutput struct{ *pulumi.OutputState }

func (InventoryTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InventoryType)(nil)).Elem()
}

func (o InventoryTypeOutput) ToInventoryTypeOutput() InventoryTypeOutput {
	return o
}

func (o InventoryTypeOutput) ToInventoryTypeOutputWithContext(ctx context.Context) InventoryTypeOutput {
	return o
}

func (o InventoryTypeOutput) ToInventoryTypePtrOutput() InventoryTypePtrOutput {
	return o.ToInventoryTypePtrOutputWithContext(context.Background())
}

func (o InventoryTypeOutput) ToInventoryTypePtrOutputWithContext(ctx context.Context) InventoryTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InventoryType) *InventoryType {
		return &v
	}).(InventoryTypePtrOutput)
}

func (o InventoryTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InventoryTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InventoryType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InventoryTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InventoryTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InventoryType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InventoryTypePtrOutput struct{ *pulumi.OutputState }

func (InventoryTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InventoryType)(nil)).Elem()
}

func (o InventoryTypePtrOutput) ToInventoryTypePtrOutput() InventoryTypePtrOutput {
	return o
}

func (o InventoryTypePtrOutput) ToInventoryTypePtrOutputWithContext(ctx context.Context) InventoryTypePtrOutput {
	return o
}

func (o InventoryTypePtrOutput) Elem() InventoryTypeOutput {
	return o.ApplyT(func(v *InventoryType) InventoryType {
		if v != nil {
			return *v
		}
		var ret InventoryType
		return ret
	}).(InventoryTypeOutput)
}

func (o InventoryTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InventoryTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InventoryType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InventoryTypeInput interface {
	pulumi.Input

	ToInventoryTypeOutput() InventoryTypeOutput
	ToInventoryTypeOutputWithContext(context.Context) InventoryTypeOutput
}

var inventoryTypePtrType = reflect.TypeOf((**InventoryType)(nil)).Elem()

type InventoryTypePtrInput interface {
	pulumi.Input

	ToInventoryTypePtrOutput() InventoryTypePtrOutput
	ToInventoryTypePtrOutputWithContext(context.Context) InventoryTypePtrOutput
}

type inventoryTypePtr string

func InventoryTypePtr(v string) InventoryTypePtrInput {
	return (*inventoryTypePtr)(&v)
}

func (*inventoryTypePtr) ElementType() reflect.Type {
	return inventoryTypePtrType
}

func (in *inventoryTypePtr) ToInventoryTypePtrOutput() InventoryTypePtrOutput {
	return pulumi.ToOutput(in).(InventoryTypePtrOutput)
}

func (in *inventoryTypePtr) ToInventoryTypePtrOutputWithContext(ctx context.Context) InventoryTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InventoryTypePtrOutput)
}

type NICType string

const (
	NICTypeVmxnet3 = NICType("vmxnet3")
	NICTypeVmxnet2 = NICType("vmxnet2")
	NICTypeVmxnet  = NICType("vmxnet")
	NICTypeE1000   = NICType("e1000")
	NICTypeE1000e  = NICType("e1000e")
	NICTypePcnet32 = NICType("pcnet32")
)

func (NICType) ElementType() reflect.Type {
	return reflect.TypeOf((*NICType)(nil)).Elem()
}

func (e NICType) ToNICTypeOutput() NICTypeOutput {
	return pulumi.ToOutput(e).(NICTypeOutput)
}

func (e NICType) ToNICTypeOutputWithContext(ctx context.Context) NICTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NICTypeOutput)
}

func (e NICType) ToNICTypePtrOutput() NICTypePtrOutput {
	return e.ToNICTypePtrOutputWithContext(context.Background())
}

func (e NICType) ToNICTypePtrOutputWithContext(ctx context.Context) NICTypePtrOutput {
	return NICType(e).ToNICTypeOutputWithContext(ctx).ToNICTypePtrOutputWithContext(ctx)
}

func (e NICType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NICType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NICType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NICType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NICTypeOutput struct{ *pulumi.OutputState }

func (NICTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NICType)(nil)).Elem()
}

func (o NICTypeOutput) ToNICTypeOutput() NICTypeOutput {
	return o
}

func (o NICTypeOutput) ToNICTypeOutputWithContext(ctx context.Context) NICTypeOutput {
	return o
}

func (o NICTypeOutput) ToNICTypePtrOutput() NICTypePtrOutput {
	return o.ToNICTypePtrOutputWithContext(context.Background())
}

func (o NICTypeOutput) ToNICTypePtrOutputWithContext(ctx context.Context) NICTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NICType) *NICType {
		return &v
	}).(NICTypePtrOutput)
}

func (o NICTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NICTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NICType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NICTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NICTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NICType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NICTypePtrOutput struct{ *pulumi.OutputState }

func (NICTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NICType)(nil)).Elem()
}

func (o NICTypePtrOutput) ToNICTypePtrOutput() NICTypePtrOutput {
	return o
}

func (o NICTypePtrOutput) ToNICTypePtrOutputWithContext(ctx context.Context) NICTypePtrOutput {
	return o
}

func (o NICTypePtrOutput) Elem() NICTypeOutput {
	return o.ApplyT(func(v *NICType) NICType {
		if v != nil {
			return *v
		}
		var ret NICType
		return ret
	}).(NICTypeOutput)
}

func (o NICTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NICTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NICType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NICTypeInput interface {
	pulumi.Input

	ToNICTypeOutput() NICTypeOutput
	ToNICTypeOutputWithContext(context.Context) NICTypeOutput
}

var nictypePtrType = reflect.TypeOf((**NICType)(nil)).Elem()

type NICTypePtrInput interface {
	pulumi.Input

	ToNICTypePtrOutput() NICTypePtrOutput
	ToNICTypePtrOutputWithContext(context.Context) NICTypePtrOutput
}

type nictypePtr string

func NICTypePtr(v string) NICTypePtrInput {
	return (*nictypePtr)(&v)
}

func (*nictypePtr) ElementType() reflect.Type {
	return nictypePtrType
}

func (in *nictypePtr) ToNICTypePtrOutput() NICTypePtrOutput {
	return pulumi.ToOutput(in).(NICTypePtrOutput)
}

func (in *nictypePtr) ToNICTypePtrOutputWithContext(ctx context.Context) NICTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NICTypePtrOutput)
}

type OsType string

const (
	OsTypeWindows = OsType("Windows")
	OsTypeLinux   = OsType("Linux")
	OsTypeOther   = OsType("Other")
)

func (OsType) ElementType() reflect.Type {
	return reflect.TypeOf((*OsType)(nil)).Elem()
}

func (e OsType) ToOsTypeOutput() OsTypeOutput {
	return pulumi.ToOutput(e).(OsTypeOutput)
}

func (e OsType) ToOsTypeOutputWithContext(ctx context.Context) OsTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OsTypeOutput)
}

func (e OsType) ToOsTypePtrOutput() OsTypePtrOutput {
	return e.ToOsTypePtrOutputWithContext(context.Background())
}

func (e OsType) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return OsType(e).ToOsTypeOutputWithContext(ctx).ToOsTypePtrOutputWithContext(ctx)
}

func (e OsType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OsType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OsType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OsType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OsTypeOutput struct{ *pulumi.OutputState }

func (OsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsType)(nil)).Elem()
}

func (o OsTypeOutput) ToOsTypeOutput() OsTypeOutput {
	return o
}

func (o OsTypeOutput) ToOsTypeOutputWithContext(ctx context.Context) OsTypeOutput {
	return o
}

func (o OsTypeOutput) ToOsTypePtrOutput() OsTypePtrOutput {
	return o.ToOsTypePtrOutputWithContext(context.Background())
}

func (o OsTypeOutput) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsType) *OsType {
		return &v
	}).(OsTypePtrOutput)
}

func (o OsTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OsTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OsType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OsTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OsTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OsType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OsTypePtrOutput struct{ *pulumi.OutputState }

func (OsTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsType)(nil)).Elem()
}

func (o OsTypePtrOutput) ToOsTypePtrOutput() OsTypePtrOutput {
	return o
}

func (o OsTypePtrOutput) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return o
}

func (o OsTypePtrOutput) Elem() OsTypeOutput {
	return o.ApplyT(func(v *OsType) OsType {
		if v != nil {
			return *v
		}
		var ret OsType
		return ret
	}).(OsTypeOutput)
}

func (o OsTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OsTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OsType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OsTypeInput interface {
	pulumi.Input

	ToOsTypeOutput() OsTypeOutput
	ToOsTypeOutputWithContext(context.Context) OsTypeOutput
}

var osTypePtrType = reflect.TypeOf((**OsType)(nil)).Elem()

type OsTypePtrInput interface {
	pulumi.Input

	ToOsTypePtrOutput() OsTypePtrOutput
	ToOsTypePtrOutputWithContext(context.Context) OsTypePtrOutput
}

type osTypePtr string

func OsTypePtr(v string) OsTypePtrInput {
	return (*osTypePtr)(&v)
}

func (*osTypePtr) ElementType() reflect.Type {
	return osTypePtrType
}

func (in *osTypePtr) ToOsTypePtrOutput() OsTypePtrOutput {
	return pulumi.ToOutput(in).(OsTypePtrOutput)
}

func (in *osTypePtr) ToOsTypePtrOutputWithContext(ctx context.Context) OsTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OsTypePtrOutput)
}

type PowerOnBootOption string

const (
	PowerOnBootOptionEnabled  = PowerOnBootOption("enabled")
	PowerOnBootOptionDisabled = PowerOnBootOption("disabled")
)

func (PowerOnBootOption) ElementType() reflect.Type {
	return reflect.TypeOf((*PowerOnBootOption)(nil)).Elem()
}

func (e PowerOnBootOption) ToPowerOnBootOptionOutput() PowerOnBootOptionOutput {
	return pulumi.ToOutput(e).(PowerOnBootOptionOutput)
}

func (e PowerOnBootOption) ToPowerOnBootOptionOutputWithContext(ctx context.Context) PowerOnBootOptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PowerOnBootOptionOutput)
}

func (e PowerOnBootOption) ToPowerOnBootOptionPtrOutput() PowerOnBootOptionPtrOutput {
	return e.ToPowerOnBootOptionPtrOutputWithContext(context.Background())
}

func (e PowerOnBootOption) ToPowerOnBootOptionPtrOutputWithContext(ctx context.Context) PowerOnBootOptionPtrOutput {
	return PowerOnBootOption(e).ToPowerOnBootOptionOutputWithContext(ctx).ToPowerOnBootOptionPtrOutputWithContext(ctx)
}

func (e PowerOnBootOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PowerOnBootOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PowerOnBootOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PowerOnBootOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PowerOnBootOptionOutput struct{ *pulumi.OutputState }

func (PowerOnBootOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PowerOnBootOption)(nil)).Elem()
}

func (o PowerOnBootOptionOutput) ToPowerOnBootOptionOutput() PowerOnBootOptionOutput {
	return o
}

func (o PowerOnBootOptionOutput) ToPowerOnBootOptionOutputWithContext(ctx context.Context) PowerOnBootOptionOutput {
	return o
}

func (o PowerOnBootOptionOutput) ToPowerOnBootOptionPtrOutput() PowerOnBootOptionPtrOutput {
	return o.ToPowerOnBootOptionPtrOutputWithContext(context.Background())
}

func (o PowerOnBootOptionOutput) ToPowerOnBootOptionPtrOutputWithContext(ctx context.Context) PowerOnBootOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PowerOnBootOption) *PowerOnBootOption {
		return &v
	}).(PowerOnBootOptionPtrOutput)
}

func (o PowerOnBootOptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PowerOnBootOptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PowerOnBootOption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PowerOnBootOptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PowerOnBootOptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PowerOnBootOption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PowerOnBootOptionPtrOutput struct{ *pulumi.OutputState }

func (PowerOnBootOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PowerOnBootOption)(nil)).Elem()
}

func (o PowerOnBootOptionPtrOutput) ToPowerOnBootOptionPtrOutput() PowerOnBootOptionPtrOutput {
	return o
}

func (o PowerOnBootOptionPtrOutput) ToPowerOnBootOptionPtrOutputWithContext(ctx context.Context) PowerOnBootOptionPtrOutput {
	return o
}

func (o PowerOnBootOptionPtrOutput) Elem() PowerOnBootOptionOutput {
	return o.ApplyT(func(v *PowerOnBootOption) PowerOnBootOption {
		if v != nil {
			return *v
		}
		var ret PowerOnBootOption
		return ret
	}).(PowerOnBootOptionOutput)
}

func (o PowerOnBootOptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PowerOnBootOptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PowerOnBootOption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PowerOnBootOptionInput interface {
	pulumi.Input

	ToPowerOnBootOptionOutput() PowerOnBootOptionOutput
	ToPowerOnBootOptionOutputWithContext(context.Context) PowerOnBootOptionOutput
}

var powerOnBootOptionPtrType = reflect.TypeOf((**PowerOnBootOption)(nil)).Elem()

type PowerOnBootOptionPtrInput interface {
	pulumi.Input

	ToPowerOnBootOptionPtrOutput() PowerOnBootOptionPtrOutput
	ToPowerOnBootOptionPtrOutputWithContext(context.Context) PowerOnBootOptionPtrOutput
}

type powerOnBootOptionPtr string

func PowerOnBootOptionPtr(v string) PowerOnBootOptionPtrInput {
	return (*powerOnBootOptionPtr)(&v)
}

func (*powerOnBootOptionPtr) ElementType() reflect.Type {
	return powerOnBootOptionPtrType
}

func (in *powerOnBootOptionPtr) ToPowerOnBootOptionPtrOutput() PowerOnBootOptionPtrOutput {
	return pulumi.ToOutput(in).(PowerOnBootOptionPtrOutput)
}

func (in *powerOnBootOptionPtr) ToPowerOnBootOptionPtrOutputWithContext(ctx context.Context) PowerOnBootOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PowerOnBootOptionPtrOutput)
}

type ProvisioningAction string

const (
	ProvisioningActionInstall   = ProvisioningAction("install")
	ProvisioningActionUninstall = ProvisioningAction("uninstall")
	ProvisioningActionRepair    = ProvisioningAction("repair")
)

func (ProvisioningAction) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningAction)(nil)).Elem()
}

func (e ProvisioningAction) ToProvisioningActionOutput() ProvisioningActionOutput {
	return pulumi.ToOutput(e).(ProvisioningActionOutput)
}

func (e ProvisioningAction) ToProvisioningActionOutputWithContext(ctx context.Context) ProvisioningActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningActionOutput)
}

func (e ProvisioningAction) ToProvisioningActionPtrOutput() ProvisioningActionPtrOutput {
	return e.ToProvisioningActionPtrOutputWithContext(context.Background())
}

func (e ProvisioningAction) ToProvisioningActionPtrOutputWithContext(ctx context.Context) ProvisioningActionPtrOutput {
	return ProvisioningAction(e).ToProvisioningActionOutputWithContext(ctx).ToProvisioningActionPtrOutputWithContext(ctx)
}

func (e ProvisioningAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningActionOutput struct{ *pulumi.OutputState }

func (ProvisioningActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningAction)(nil)).Elem()
}

func (o ProvisioningActionOutput) ToProvisioningActionOutput() ProvisioningActionOutput {
	return o
}

func (o ProvisioningActionOutput) ToProvisioningActionOutputWithContext(ctx context.Context) ProvisioningActionOutput {
	return o
}

func (o ProvisioningActionOutput) ToProvisioningActionPtrOutput() ProvisioningActionPtrOutput {
	return o.ToProvisioningActionPtrOutputWithContext(context.Background())
}

func (o ProvisioningActionOutput) ToProvisioningActionPtrOutputWithContext(ctx context.Context) ProvisioningActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningAction) *ProvisioningAction {
		return &v
	}).(ProvisioningActionPtrOutput)
}

func (o ProvisioningActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningActionPtrOutput struct{ *pulumi.OutputState }

func (ProvisioningActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningAction)(nil)).Elem()
}

func (o ProvisioningActionPtrOutput) ToProvisioningActionPtrOutput() ProvisioningActionPtrOutput {
	return o
}

func (o ProvisioningActionPtrOutput) ToProvisioningActionPtrOutputWithContext(ctx context.Context) ProvisioningActionPtrOutput {
	return o
}

func (o ProvisioningActionPtrOutput) Elem() ProvisioningActionOutput {
	return o.ApplyT(func(v *ProvisioningAction) ProvisioningAction {
		if v != nil {
			return *v
		}
		var ret ProvisioningAction
		return ret
	}).(ProvisioningActionOutput)
}

func (o ProvisioningActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProvisioningActionInput interface {
	pulumi.Input

	ToProvisioningActionOutput() ProvisioningActionOutput
	ToProvisioningActionOutputWithContext(context.Context) ProvisioningActionOutput
}

var provisioningActionPtrType = reflect.TypeOf((**ProvisioningAction)(nil)).Elem()

type ProvisioningActionPtrInput interface {
	pulumi.Input

	ToProvisioningActionPtrOutput() ProvisioningActionPtrOutput
	ToProvisioningActionPtrOutputWithContext(context.Context) ProvisioningActionPtrOutput
}

type provisioningActionPtr string

func ProvisioningActionPtr(v string) ProvisioningActionPtrInput {
	return (*provisioningActionPtr)(&v)
}

func (*provisioningActionPtr) ElementType() reflect.Type {
	return provisioningActionPtrType
}

func (in *provisioningActionPtr) ToProvisioningActionPtrOutput() ProvisioningActionPtrOutput {
	return pulumi.ToOutput(in).(ProvisioningActionPtrOutput)
}

func (in *provisioningActionPtr) ToProvisioningActionPtrOutputWithContext(ctx context.Context) ProvisioningActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningActionPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskModeInput)(nil)).Elem(), DiskMode("persistent"))
	pulumi.RegisterInputType(reflect.TypeOf((*DiskModePtrInput)(nil)).Elem(), DiskMode("persistent"))
	pulumi.RegisterInputType(reflect.TypeOf((*DiskTypeInput)(nil)).Elem(), DiskType("flat"))
	pulumi.RegisterInputType(reflect.TypeOf((*DiskTypePtrInput)(nil)).Elem(), DiskType("flat"))
	pulumi.RegisterInputType(reflect.TypeOf((*FirmwareTypeInput)(nil)).Elem(), FirmwareType("bios"))
	pulumi.RegisterInputType(reflect.TypeOf((*FirmwareTypePtrInput)(nil)).Elem(), FirmwareType("bios"))
	pulumi.RegisterInputType(reflect.TypeOf((*IPAddressAllocationMethodInput)(nil)).Elem(), IPAddressAllocationMethod("unset"))
	pulumi.RegisterInputType(reflect.TypeOf((*IPAddressAllocationMethodPtrInput)(nil)).Elem(), IPAddressAllocationMethod("unset"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypeInput)(nil)).Elem(), IdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypePtrInput)(nil)).Elem(), IdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*InventoryTypeInput)(nil)).Elem(), InventoryType("ResourcePool"))
	pulumi.RegisterInputType(reflect.TypeOf((*InventoryTypePtrInput)(nil)).Elem(), InventoryType("ResourcePool"))
	pulumi.RegisterInputType(reflect.TypeOf((*NICTypeInput)(nil)).Elem(), NICType("vmxnet3"))
	pulumi.RegisterInputType(reflect.TypeOf((*NICTypePtrInput)(nil)).Elem(), NICType("vmxnet3"))
	pulumi.RegisterInputType(reflect.TypeOf((*OsTypeInput)(nil)).Elem(), OsType("Windows"))
	pulumi.RegisterInputType(reflect.TypeOf((*OsTypePtrInput)(nil)).Elem(), OsType("Windows"))
	pulumi.RegisterInputType(reflect.TypeOf((*PowerOnBootOptionInput)(nil)).Elem(), PowerOnBootOption("enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*PowerOnBootOptionPtrInput)(nil)).Elem(), PowerOnBootOption("enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningActionInput)(nil)).Elem(), ProvisioningAction("install"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningActionPtrInput)(nil)).Elem(), ProvisioningAction("install"))
	pulumi.RegisterOutputType(DiskModeOutput{})
	pulumi.RegisterOutputType(DiskModePtrOutput{})
	pulumi.RegisterOutputType(DiskTypeOutput{})
	pulumi.RegisterOutputType(DiskTypePtrOutput{})
	pulumi.RegisterOutputType(FirmwareTypeOutput{})
	pulumi.RegisterOutputType(FirmwareTypePtrOutput{})
	pulumi.RegisterOutputType(IPAddressAllocationMethodOutput{})
	pulumi.RegisterOutputType(IPAddressAllocationMethodPtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(InventoryTypeOutput{})
	pulumi.RegisterOutputType(InventoryTypePtrOutput{})
	pulumi.RegisterOutputType(NICTypeOutput{})
	pulumi.RegisterOutputType(NICTypePtrOutput{})
	pulumi.RegisterOutputType(OsTypeOutput{})
	pulumi.RegisterOutputType(OsTypePtrOutput{})
	pulumi.RegisterOutputType(PowerOnBootOptionOutput{})
	pulumi.RegisterOutputType(PowerOnBootOptionPtrOutput{})
	pulumi.RegisterOutputType(ProvisioningActionOutput{})
	pulumi.RegisterOutputType(ProvisioningActionPtrOutput{})
}
