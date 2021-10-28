


package powerbidedicated

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CapacitySkuTier string

const (
	CapacitySkuTier_PBIE_Azure     = CapacitySkuTier("PBIE_Azure")
	CapacitySkuTierPremium         = CapacitySkuTier("Premium")
	CapacitySkuTierAutoPremiumHost = CapacitySkuTier("AutoPremiumHost")
)

func (CapacitySkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacitySkuTier)(nil)).Elem()
}

func (e CapacitySkuTier) ToCapacitySkuTierOutput() CapacitySkuTierOutput {
	return pulumi.ToOutput(e).(CapacitySkuTierOutput)
}

func (e CapacitySkuTier) ToCapacitySkuTierOutputWithContext(ctx context.Context) CapacitySkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CapacitySkuTierOutput)
}

func (e CapacitySkuTier) ToCapacitySkuTierPtrOutput() CapacitySkuTierPtrOutput {
	return e.ToCapacitySkuTierPtrOutputWithContext(context.Background())
}

func (e CapacitySkuTier) ToCapacitySkuTierPtrOutputWithContext(ctx context.Context) CapacitySkuTierPtrOutput {
	return CapacitySkuTier(e).ToCapacitySkuTierOutputWithContext(ctx).ToCapacitySkuTierPtrOutputWithContext(ctx)
}

func (e CapacitySkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CapacitySkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CapacitySkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CapacitySkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CapacitySkuTierOutput struct{ *pulumi.OutputState }

func (CapacitySkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacitySkuTier)(nil)).Elem()
}

func (o CapacitySkuTierOutput) ToCapacitySkuTierOutput() CapacitySkuTierOutput {
	return o
}

func (o CapacitySkuTierOutput) ToCapacitySkuTierOutputWithContext(ctx context.Context) CapacitySkuTierOutput {
	return o
}

func (o CapacitySkuTierOutput) ToCapacitySkuTierPtrOutput() CapacitySkuTierPtrOutput {
	return o.ToCapacitySkuTierPtrOutputWithContext(context.Background())
}

func (o CapacitySkuTierOutput) ToCapacitySkuTierPtrOutputWithContext(ctx context.Context) CapacitySkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CapacitySkuTier) *CapacitySkuTier {
		return &v
	}).(CapacitySkuTierPtrOutput)
}

func (o CapacitySkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CapacitySkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CapacitySkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CapacitySkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CapacitySkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CapacitySkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CapacitySkuTierPtrOutput struct{ *pulumi.OutputState }

func (CapacitySkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacitySkuTier)(nil)).Elem()
}

func (o CapacitySkuTierPtrOutput) ToCapacitySkuTierPtrOutput() CapacitySkuTierPtrOutput {
	return o
}

func (o CapacitySkuTierPtrOutput) ToCapacitySkuTierPtrOutputWithContext(ctx context.Context) CapacitySkuTierPtrOutput {
	return o
}

func (o CapacitySkuTierPtrOutput) Elem() CapacitySkuTierOutput {
	return o.ApplyT(func(v *CapacitySkuTier) CapacitySkuTier {
		if v != nil {
			return *v
		}
		var ret CapacitySkuTier
		return ret
	}).(CapacitySkuTierOutput)
}

func (o CapacitySkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CapacitySkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CapacitySkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CapacitySkuTierInput interface {
	pulumi.Input

	ToCapacitySkuTierOutput() CapacitySkuTierOutput
	ToCapacitySkuTierOutputWithContext(context.Context) CapacitySkuTierOutput
}

var capacitySkuTierPtrType = reflect.TypeOf((**CapacitySkuTier)(nil)).Elem()

type CapacitySkuTierPtrInput interface {
	pulumi.Input

	ToCapacitySkuTierPtrOutput() CapacitySkuTierPtrOutput
	ToCapacitySkuTierPtrOutputWithContext(context.Context) CapacitySkuTierPtrOutput
}

type capacitySkuTierPtr string

func CapacitySkuTierPtr(v string) CapacitySkuTierPtrInput {
	return (*capacitySkuTierPtr)(&v)
}

func (*capacitySkuTierPtr) ElementType() reflect.Type {
	return capacitySkuTierPtrType
}

func (in *capacitySkuTierPtr) ToCapacitySkuTierPtrOutput() CapacitySkuTierPtrOutput {
	return pulumi.ToOutput(in).(CapacitySkuTierPtrOutput)
}

func (in *capacitySkuTierPtr) ToCapacitySkuTierPtrOutputWithContext(ctx context.Context) CapacitySkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CapacitySkuTierPtrOutput)
}

type IdentityType string

const (
	IdentityTypeUser            = IdentityType("User")
	IdentityTypeApplication     = IdentityType("Application")
	IdentityTypeManagedIdentity = IdentityType("ManagedIdentity")
	IdentityTypeKey             = IdentityType("Key")
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

type Mode string

const (
	ModeGen1 = Mode("Gen1")
	ModeGen2 = Mode("Gen2")
)

func (Mode) ElementType() reflect.Type {
	return reflect.TypeOf((*Mode)(nil)).Elem()
}

func (e Mode) ToModeOutput() ModeOutput {
	return pulumi.ToOutput(e).(ModeOutput)
}

func (e Mode) ToModeOutputWithContext(ctx context.Context) ModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ModeOutput)
}

func (e Mode) ToModePtrOutput() ModePtrOutput {
	return e.ToModePtrOutputWithContext(context.Background())
}

func (e Mode) ToModePtrOutputWithContext(ctx context.Context) ModePtrOutput {
	return Mode(e).ToModeOutputWithContext(ctx).ToModePtrOutputWithContext(ctx)
}

func (e Mode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Mode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Mode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Mode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ModeOutput struct{ *pulumi.OutputState }

func (ModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Mode)(nil)).Elem()
}

func (o ModeOutput) ToModeOutput() ModeOutput {
	return o
}

func (o ModeOutput) ToModeOutputWithContext(ctx context.Context) ModeOutput {
	return o
}

func (o ModeOutput) ToModePtrOutput() ModePtrOutput {
	return o.ToModePtrOutputWithContext(context.Background())
}

func (o ModeOutput) ToModePtrOutputWithContext(ctx context.Context) ModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Mode) *Mode {
		return &v
	}).(ModePtrOutput)
}

func (o ModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Mode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Mode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ModePtrOutput struct{ *pulumi.OutputState }

func (ModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mode)(nil)).Elem()
}

func (o ModePtrOutput) ToModePtrOutput() ModePtrOutput {
	return o
}

func (o ModePtrOutput) ToModePtrOutputWithContext(ctx context.Context) ModePtrOutput {
	return o
}

func (o ModePtrOutput) Elem() ModeOutput {
	return o.ApplyT(func(v *Mode) Mode {
		if v != nil {
			return *v
		}
		var ret Mode
		return ret
	}).(ModeOutput)
}

func (o ModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Mode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ModeInput interface {
	pulumi.Input

	ToModeOutput() ModeOutput
	ToModeOutputWithContext(context.Context) ModeOutput
}

var modePtrType = reflect.TypeOf((**Mode)(nil)).Elem()

type ModePtrInput interface {
	pulumi.Input

	ToModePtrOutput() ModePtrOutput
	ToModePtrOutputWithContext(context.Context) ModePtrOutput
}

type modePtr string

func ModePtr(v string) ModePtrInput {
	return (*modePtr)(&v)
}

func (*modePtr) ElementType() reflect.Type {
	return modePtrType
}

func (in *modePtr) ToModePtrOutput() ModePtrOutput {
	return pulumi.ToOutput(in).(ModePtrOutput)
}

func (in *modePtr) ToModePtrOutputWithContext(ctx context.Context) ModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ModePtrOutput)
}

type VCoreSkuTier string

const (
	VCoreSkuTierAutoScale = VCoreSkuTier("AutoScale")
)

func (VCoreSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*VCoreSkuTier)(nil)).Elem()
}

func (e VCoreSkuTier) ToVCoreSkuTierOutput() VCoreSkuTierOutput {
	return pulumi.ToOutput(e).(VCoreSkuTierOutput)
}

func (e VCoreSkuTier) ToVCoreSkuTierOutputWithContext(ctx context.Context) VCoreSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VCoreSkuTierOutput)
}

func (e VCoreSkuTier) ToVCoreSkuTierPtrOutput() VCoreSkuTierPtrOutput {
	return e.ToVCoreSkuTierPtrOutputWithContext(context.Background())
}

func (e VCoreSkuTier) ToVCoreSkuTierPtrOutputWithContext(ctx context.Context) VCoreSkuTierPtrOutput {
	return VCoreSkuTier(e).ToVCoreSkuTierOutputWithContext(ctx).ToVCoreSkuTierPtrOutputWithContext(ctx)
}

func (e VCoreSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VCoreSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VCoreSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VCoreSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VCoreSkuTierOutput struct{ *pulumi.OutputState }

func (VCoreSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VCoreSkuTier)(nil)).Elem()
}

func (o VCoreSkuTierOutput) ToVCoreSkuTierOutput() VCoreSkuTierOutput {
	return o
}

func (o VCoreSkuTierOutput) ToVCoreSkuTierOutputWithContext(ctx context.Context) VCoreSkuTierOutput {
	return o
}

func (o VCoreSkuTierOutput) ToVCoreSkuTierPtrOutput() VCoreSkuTierPtrOutput {
	return o.ToVCoreSkuTierPtrOutputWithContext(context.Background())
}

func (o VCoreSkuTierOutput) ToVCoreSkuTierPtrOutputWithContext(ctx context.Context) VCoreSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VCoreSkuTier) *VCoreSkuTier {
		return &v
	}).(VCoreSkuTierPtrOutput)
}

func (o VCoreSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VCoreSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VCoreSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VCoreSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VCoreSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VCoreSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VCoreSkuTierPtrOutput struct{ *pulumi.OutputState }

func (VCoreSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VCoreSkuTier)(nil)).Elem()
}

func (o VCoreSkuTierPtrOutput) ToVCoreSkuTierPtrOutput() VCoreSkuTierPtrOutput {
	return o
}

func (o VCoreSkuTierPtrOutput) ToVCoreSkuTierPtrOutputWithContext(ctx context.Context) VCoreSkuTierPtrOutput {
	return o
}

func (o VCoreSkuTierPtrOutput) Elem() VCoreSkuTierOutput {
	return o.ApplyT(func(v *VCoreSkuTier) VCoreSkuTier {
		if v != nil {
			return *v
		}
		var ret VCoreSkuTier
		return ret
	}).(VCoreSkuTierOutput)
}

func (o VCoreSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VCoreSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VCoreSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VCoreSkuTierInput interface {
	pulumi.Input

	ToVCoreSkuTierOutput() VCoreSkuTierOutput
	ToVCoreSkuTierOutputWithContext(context.Context) VCoreSkuTierOutput
}

var vcoreSkuTierPtrType = reflect.TypeOf((**VCoreSkuTier)(nil)).Elem()

type VCoreSkuTierPtrInput interface {
	pulumi.Input

	ToVCoreSkuTierPtrOutput() VCoreSkuTierPtrOutput
	ToVCoreSkuTierPtrOutputWithContext(context.Context) VCoreSkuTierPtrOutput
}

type vcoreSkuTierPtr string

func VCoreSkuTierPtr(v string) VCoreSkuTierPtrInput {
	return (*vcoreSkuTierPtr)(&v)
}

func (*vcoreSkuTierPtr) ElementType() reflect.Type {
	return vcoreSkuTierPtrType
}

func (in *vcoreSkuTierPtr) ToVCoreSkuTierPtrOutput() VCoreSkuTierPtrOutput {
	return pulumi.ToOutput(in).(VCoreSkuTierPtrOutput)
}

func (in *vcoreSkuTierPtr) ToVCoreSkuTierPtrOutputWithContext(ctx context.Context) VCoreSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VCoreSkuTierPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CapacitySkuTierInput)(nil)).Elem(), CapacitySkuTier("PBIE_Azure"))
	pulumi.RegisterInputType(reflect.TypeOf((*CapacitySkuTierPtrInput)(nil)).Elem(), CapacitySkuTier("PBIE_Azure"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypeInput)(nil)).Elem(), IdentityType("User"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypePtrInput)(nil)).Elem(), IdentityType("User"))
	pulumi.RegisterInputType(reflect.TypeOf((*ModeInput)(nil)).Elem(), Mode("Gen1"))
	pulumi.RegisterInputType(reflect.TypeOf((*ModePtrInput)(nil)).Elem(), Mode("Gen1"))
	pulumi.RegisterInputType(reflect.TypeOf((*VCoreSkuTierInput)(nil)).Elem(), VCoreSkuTier("AutoScale"))
	pulumi.RegisterInputType(reflect.TypeOf((*VCoreSkuTierPtrInput)(nil)).Elem(), VCoreSkuTier("AutoScale"))
	pulumi.RegisterOutputType(CapacitySkuTierOutput{})
	pulumi.RegisterOutputType(CapacitySkuTierPtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(ModeOutput{})
	pulumi.RegisterOutputType(ModePtrOutput{})
	pulumi.RegisterOutputType(VCoreSkuTierOutput{})
	pulumi.RegisterOutputType(VCoreSkuTierPtrOutput{})
}
