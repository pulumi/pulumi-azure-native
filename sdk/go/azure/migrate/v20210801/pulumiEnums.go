


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned   = ResourceIdentityType("UserAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type TargetAvailabilityZone string

const (
	TargetAvailabilityZoneOne   = TargetAvailabilityZone("1")
	TargetAvailabilityZoneTwo   = TargetAvailabilityZone("2")
	TargetAvailabilityZoneThree = TargetAvailabilityZone("3")
	TargetAvailabilityZoneNA    = TargetAvailabilityZone("NA")
)

func (TargetAvailabilityZone) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetAvailabilityZone)(nil)).Elem()
}

func (e TargetAvailabilityZone) ToTargetAvailabilityZoneOutput() TargetAvailabilityZoneOutput {
	return pulumi.ToOutput(e).(TargetAvailabilityZoneOutput)
}

func (e TargetAvailabilityZone) ToTargetAvailabilityZoneOutputWithContext(ctx context.Context) TargetAvailabilityZoneOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TargetAvailabilityZoneOutput)
}

func (e TargetAvailabilityZone) ToTargetAvailabilityZonePtrOutput() TargetAvailabilityZonePtrOutput {
	return e.ToTargetAvailabilityZonePtrOutputWithContext(context.Background())
}

func (e TargetAvailabilityZone) ToTargetAvailabilityZonePtrOutputWithContext(ctx context.Context) TargetAvailabilityZonePtrOutput {
	return TargetAvailabilityZone(e).ToTargetAvailabilityZoneOutputWithContext(ctx).ToTargetAvailabilityZonePtrOutputWithContext(ctx)
}

func (e TargetAvailabilityZone) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TargetAvailabilityZone) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TargetAvailabilityZone) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TargetAvailabilityZone) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TargetAvailabilityZoneOutput struct{ *pulumi.OutputState }

func (TargetAvailabilityZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetAvailabilityZone)(nil)).Elem()
}

func (o TargetAvailabilityZoneOutput) ToTargetAvailabilityZoneOutput() TargetAvailabilityZoneOutput {
	return o
}

func (o TargetAvailabilityZoneOutput) ToTargetAvailabilityZoneOutputWithContext(ctx context.Context) TargetAvailabilityZoneOutput {
	return o
}

func (o TargetAvailabilityZoneOutput) ToTargetAvailabilityZonePtrOutput() TargetAvailabilityZonePtrOutput {
	return o.ToTargetAvailabilityZonePtrOutputWithContext(context.Background())
}

func (o TargetAvailabilityZoneOutput) ToTargetAvailabilityZonePtrOutputWithContext(ctx context.Context) TargetAvailabilityZonePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TargetAvailabilityZone) *TargetAvailabilityZone {
		return &v
	}).(TargetAvailabilityZonePtrOutput)
}

func (o TargetAvailabilityZoneOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TargetAvailabilityZoneOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TargetAvailabilityZone) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TargetAvailabilityZoneOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TargetAvailabilityZoneOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TargetAvailabilityZone) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TargetAvailabilityZonePtrOutput struct{ *pulumi.OutputState }

func (TargetAvailabilityZonePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetAvailabilityZone)(nil)).Elem()
}

func (o TargetAvailabilityZonePtrOutput) ToTargetAvailabilityZonePtrOutput() TargetAvailabilityZonePtrOutput {
	return o
}

func (o TargetAvailabilityZonePtrOutput) ToTargetAvailabilityZonePtrOutputWithContext(ctx context.Context) TargetAvailabilityZonePtrOutput {
	return o
}

func (o TargetAvailabilityZonePtrOutput) Elem() TargetAvailabilityZoneOutput {
	return o.ApplyT(func(v *TargetAvailabilityZone) TargetAvailabilityZone {
		if v != nil {
			return *v
		}
		var ret TargetAvailabilityZone
		return ret
	}).(TargetAvailabilityZoneOutput)
}

func (o TargetAvailabilityZonePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TargetAvailabilityZonePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TargetAvailabilityZone) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TargetAvailabilityZoneInput interface {
	pulumi.Input

	ToTargetAvailabilityZoneOutput() TargetAvailabilityZoneOutput
	ToTargetAvailabilityZoneOutputWithContext(context.Context) TargetAvailabilityZoneOutput
}

var targetAvailabilityZonePtrType = reflect.TypeOf((**TargetAvailabilityZone)(nil)).Elem()

type TargetAvailabilityZonePtrInput interface {
	pulumi.Input

	ToTargetAvailabilityZonePtrOutput() TargetAvailabilityZonePtrOutput
	ToTargetAvailabilityZonePtrOutputWithContext(context.Context) TargetAvailabilityZonePtrOutput
}

type targetAvailabilityZonePtr string

func TargetAvailabilityZonePtr(v string) TargetAvailabilityZonePtrInput {
	return (*targetAvailabilityZonePtr)(&v)
}

func (*targetAvailabilityZonePtr) ElementType() reflect.Type {
	return targetAvailabilityZonePtrType
}

func (in *targetAvailabilityZonePtr) ToTargetAvailabilityZonePtrOutput() TargetAvailabilityZonePtrOutput {
	return pulumi.ToOutput(in).(TargetAvailabilityZonePtrOutput)
}

func (in *targetAvailabilityZonePtr) ToTargetAvailabilityZonePtrOutputWithContext(ctx context.Context) TargetAvailabilityZonePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TargetAvailabilityZonePtrOutput)
}

type ZoneRedundant string

const (
	ZoneRedundantEnable  = ZoneRedundant("Enable")
	ZoneRedundantDisable = ZoneRedundant("Disable")
)

func (ZoneRedundant) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneRedundant)(nil)).Elem()
}

func (e ZoneRedundant) ToZoneRedundantOutput() ZoneRedundantOutput {
	return pulumi.ToOutput(e).(ZoneRedundantOutput)
}

func (e ZoneRedundant) ToZoneRedundantOutputWithContext(ctx context.Context) ZoneRedundantOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ZoneRedundantOutput)
}

func (e ZoneRedundant) ToZoneRedundantPtrOutput() ZoneRedundantPtrOutput {
	return e.ToZoneRedundantPtrOutputWithContext(context.Background())
}

func (e ZoneRedundant) ToZoneRedundantPtrOutputWithContext(ctx context.Context) ZoneRedundantPtrOutput {
	return ZoneRedundant(e).ToZoneRedundantOutputWithContext(ctx).ToZoneRedundantPtrOutputWithContext(ctx)
}

func (e ZoneRedundant) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZoneRedundant) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZoneRedundant) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ZoneRedundant) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ZoneRedundantOutput struct{ *pulumi.OutputState }

func (ZoneRedundantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneRedundant)(nil)).Elem()
}

func (o ZoneRedundantOutput) ToZoneRedundantOutput() ZoneRedundantOutput {
	return o
}

func (o ZoneRedundantOutput) ToZoneRedundantOutputWithContext(ctx context.Context) ZoneRedundantOutput {
	return o
}

func (o ZoneRedundantOutput) ToZoneRedundantPtrOutput() ZoneRedundantPtrOutput {
	return o.ToZoneRedundantPtrOutputWithContext(context.Background())
}

func (o ZoneRedundantOutput) ToZoneRedundantPtrOutputWithContext(ctx context.Context) ZoneRedundantPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ZoneRedundant) *ZoneRedundant {
		return &v
	}).(ZoneRedundantPtrOutput)
}

func (o ZoneRedundantOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ZoneRedundantOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZoneRedundant) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ZoneRedundantOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZoneRedundantOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ZoneRedundant) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ZoneRedundantPtrOutput struct{ *pulumi.OutputState }

func (ZoneRedundantPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneRedundant)(nil)).Elem()
}

func (o ZoneRedundantPtrOutput) ToZoneRedundantPtrOutput() ZoneRedundantPtrOutput {
	return o
}

func (o ZoneRedundantPtrOutput) ToZoneRedundantPtrOutputWithContext(ctx context.Context) ZoneRedundantPtrOutput {
	return o
}

func (o ZoneRedundantPtrOutput) Elem() ZoneRedundantOutput {
	return o.ApplyT(func(v *ZoneRedundant) ZoneRedundant {
		if v != nil {
			return *v
		}
		var ret ZoneRedundant
		return ret
	}).(ZoneRedundantOutput)
}

func (o ZoneRedundantPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ZoneRedundantPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ZoneRedundant) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ZoneRedundantInput interface {
	pulumi.Input

	ToZoneRedundantOutput() ZoneRedundantOutput
	ToZoneRedundantOutputWithContext(context.Context) ZoneRedundantOutput
}

var zoneRedundantPtrType = reflect.TypeOf((**ZoneRedundant)(nil)).Elem()

type ZoneRedundantPtrInput interface {
	pulumi.Input

	ToZoneRedundantPtrOutput() ZoneRedundantPtrOutput
	ToZoneRedundantPtrOutputWithContext(context.Context) ZoneRedundantPtrOutput
}

type zoneRedundantPtr string

func ZoneRedundantPtr(v string) ZoneRedundantPtrInput {
	return (*zoneRedundantPtr)(&v)
}

func (*zoneRedundantPtr) ElementType() reflect.Type {
	return zoneRedundantPtrType
}

func (in *zoneRedundantPtr) ToZoneRedundantPtrOutput() ZoneRedundantPtrOutput {
	return pulumi.ToOutput(in).(ZoneRedundantPtrOutput)
}

func (in *zoneRedundantPtr) ToZoneRedundantPtrOutputWithContext(ctx context.Context) ZoneRedundantPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ZoneRedundantPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(TargetAvailabilityZoneOutput{})
	pulumi.RegisterOutputType(TargetAvailabilityZonePtrOutput{})
	pulumi.RegisterOutputType(ZoneRedundantOutput{})
	pulumi.RegisterOutputType(ZoneRedundantPtrOutput{})
}
