


package v20200201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EndpointType string

const (
	EndpointTypeSrc = EndpointType("src")
	EndpointTypeDst = EndpointType("dst")
)

func (EndpointType) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (e EndpointType) ToEndpointTypeOutput() EndpointTypeOutput {
	return pulumi.ToOutput(e).(EndpointTypeOutput)
}

func (e EndpointType) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointTypeOutput)
}

func (e EndpointType) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return e.ToEndpointTypePtrOutputWithContext(context.Background())
}

func (e EndpointType) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return EndpointType(e).ToEndpointTypeOutputWithContext(ctx).ToEndpointTypePtrOutputWithContext(ctx)
}

func (e EndpointType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointTypeOutput struct{ *pulumi.OutputState }

func (EndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (o EndpointTypeOutput) ToEndpointTypeOutput() EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return o.ToEndpointTypePtrOutputWithContext(context.Background())
}

func (o EndpointTypeOutput) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointType) *EndpointType {
		return &v
	}).(EndpointTypePtrOutput)
}

func (o EndpointTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointTypePtrOutput struct{ *pulumi.OutputState }

func (EndpointTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointType)(nil)).Elem()
}

func (o EndpointTypePtrOutput) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return o
}

func (o EndpointTypePtrOutput) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return o
}

func (o EndpointTypePtrOutput) Elem() EndpointTypeOutput {
	return o.ApplyT(func(v *EndpointType) EndpointType {
		if v != nil {
			return *v
		}
		var ret EndpointType
		return ret
	}).(EndpointTypeOutput)
}

func (o EndpointTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EndpointTypeInput interface {
	pulumi.Input

	ToEndpointTypeOutput() EndpointTypeOutput
	ToEndpointTypeOutputWithContext(context.Context) EndpointTypeOutput
}

var endpointTypePtrType = reflect.TypeOf((**EndpointType)(nil)).Elem()

type EndpointTypePtrInput interface {
	pulumi.Input

	ToEndpointTypePtrOutput() EndpointTypePtrOutput
	ToEndpointTypePtrOutputWithContext(context.Context) EndpointTypePtrOutput
}

type endpointTypePtr string

func EndpointTypePtr(v string) EndpointTypePtrInput {
	return (*endpointTypePtr)(&v)
}

func (*endpointTypePtr) ElementType() reflect.Type {
	return endpointTypePtrType
}

func (in *endpointTypePtr) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return pulumi.ToOutput(in).(EndpointTypePtrOutput)
}

func (in *endpointTypePtr) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointTypePtrOutput)
}

type PoolServiceLevel string

const (
	// Standard service level
	PoolServiceLevelStandard = PoolServiceLevel("Standard")
	// Premium service level
	PoolServiceLevelPremium = PoolServiceLevel("Premium")
	// Ultra service level
	PoolServiceLevelUltra = PoolServiceLevel("Ultra")
)

func (PoolServiceLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*PoolServiceLevel)(nil)).Elem()
}

func (e PoolServiceLevel) ToPoolServiceLevelOutput() PoolServiceLevelOutput {
	return pulumi.ToOutput(e).(PoolServiceLevelOutput)
}

func (e PoolServiceLevel) ToPoolServiceLevelOutputWithContext(ctx context.Context) PoolServiceLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PoolServiceLevelOutput)
}

func (e PoolServiceLevel) ToPoolServiceLevelPtrOutput() PoolServiceLevelPtrOutput {
	return e.ToPoolServiceLevelPtrOutputWithContext(context.Background())
}

func (e PoolServiceLevel) ToPoolServiceLevelPtrOutputWithContext(ctx context.Context) PoolServiceLevelPtrOutput {
	return PoolServiceLevel(e).ToPoolServiceLevelOutputWithContext(ctx).ToPoolServiceLevelPtrOutputWithContext(ctx)
}

func (e PoolServiceLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PoolServiceLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PoolServiceLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PoolServiceLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PoolServiceLevelOutput struct{ *pulumi.OutputState }

func (PoolServiceLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PoolServiceLevel)(nil)).Elem()
}

func (o PoolServiceLevelOutput) ToPoolServiceLevelOutput() PoolServiceLevelOutput {
	return o
}

func (o PoolServiceLevelOutput) ToPoolServiceLevelOutputWithContext(ctx context.Context) PoolServiceLevelOutput {
	return o
}

func (o PoolServiceLevelOutput) ToPoolServiceLevelPtrOutput() PoolServiceLevelPtrOutput {
	return o.ToPoolServiceLevelPtrOutputWithContext(context.Background())
}

func (o PoolServiceLevelOutput) ToPoolServiceLevelPtrOutputWithContext(ctx context.Context) PoolServiceLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PoolServiceLevel) *PoolServiceLevel {
		return &v
	}).(PoolServiceLevelPtrOutput)
}

func (o PoolServiceLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PoolServiceLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PoolServiceLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PoolServiceLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PoolServiceLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PoolServiceLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PoolServiceLevelPtrOutput struct{ *pulumi.OutputState }

func (PoolServiceLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PoolServiceLevel)(nil)).Elem()
}

func (o PoolServiceLevelPtrOutput) ToPoolServiceLevelPtrOutput() PoolServiceLevelPtrOutput {
	return o
}

func (o PoolServiceLevelPtrOutput) ToPoolServiceLevelPtrOutputWithContext(ctx context.Context) PoolServiceLevelPtrOutput {
	return o
}

func (o PoolServiceLevelPtrOutput) Elem() PoolServiceLevelOutput {
	return o.ApplyT(func(v *PoolServiceLevel) PoolServiceLevel {
		if v != nil {
			return *v
		}
		var ret PoolServiceLevel
		return ret
	}).(PoolServiceLevelOutput)
}

func (o PoolServiceLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PoolServiceLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PoolServiceLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PoolServiceLevelInput interface {
	pulumi.Input

	ToPoolServiceLevelOutput() PoolServiceLevelOutput
	ToPoolServiceLevelOutputWithContext(context.Context) PoolServiceLevelOutput
}

var poolServiceLevelPtrType = reflect.TypeOf((**PoolServiceLevel)(nil)).Elem()

type PoolServiceLevelPtrInput interface {
	pulumi.Input

	ToPoolServiceLevelPtrOutput() PoolServiceLevelPtrOutput
	ToPoolServiceLevelPtrOutputWithContext(context.Context) PoolServiceLevelPtrOutput
}

type poolServiceLevelPtr string

func PoolServiceLevelPtr(v string) PoolServiceLevelPtrInput {
	return (*poolServiceLevelPtr)(&v)
}

func (*poolServiceLevelPtr) ElementType() reflect.Type {
	return poolServiceLevelPtrType
}

func (in *poolServiceLevelPtr) ToPoolServiceLevelPtrOutput() PoolServiceLevelPtrOutput {
	return pulumi.ToOutput(in).(PoolServiceLevelPtrOutput)
}

func (in *poolServiceLevelPtr) ToPoolServiceLevelPtrOutputWithContext(ctx context.Context) PoolServiceLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PoolServiceLevelPtrOutput)
}

type ReplicationSchedule string

const (
	ReplicationSchedule_10minutely = ReplicationSchedule("_10minutely")
	ReplicationScheduleHourly      = ReplicationSchedule("hourly")
	ReplicationScheduleDaily       = ReplicationSchedule("daily")
)

func (ReplicationSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationSchedule)(nil)).Elem()
}

func (e ReplicationSchedule) ToReplicationScheduleOutput() ReplicationScheduleOutput {
	return pulumi.ToOutput(e).(ReplicationScheduleOutput)
}

func (e ReplicationSchedule) ToReplicationScheduleOutputWithContext(ctx context.Context) ReplicationScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReplicationScheduleOutput)
}

func (e ReplicationSchedule) ToReplicationSchedulePtrOutput() ReplicationSchedulePtrOutput {
	return e.ToReplicationSchedulePtrOutputWithContext(context.Background())
}

func (e ReplicationSchedule) ToReplicationSchedulePtrOutputWithContext(ctx context.Context) ReplicationSchedulePtrOutput {
	return ReplicationSchedule(e).ToReplicationScheduleOutputWithContext(ctx).ToReplicationSchedulePtrOutputWithContext(ctx)
}

func (e ReplicationSchedule) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationSchedule) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReplicationSchedule) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReplicationSchedule) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReplicationScheduleOutput struct{ *pulumi.OutputState }

func (ReplicationScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationSchedule)(nil)).Elem()
}

func (o ReplicationScheduleOutput) ToReplicationScheduleOutput() ReplicationScheduleOutput {
	return o
}

func (o ReplicationScheduleOutput) ToReplicationScheduleOutputWithContext(ctx context.Context) ReplicationScheduleOutput {
	return o
}

func (o ReplicationScheduleOutput) ToReplicationSchedulePtrOutput() ReplicationSchedulePtrOutput {
	return o.ToReplicationSchedulePtrOutputWithContext(context.Background())
}

func (o ReplicationScheduleOutput) ToReplicationSchedulePtrOutputWithContext(ctx context.Context) ReplicationSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationSchedule) *ReplicationSchedule {
		return &v
	}).(ReplicationSchedulePtrOutput)
}

func (o ReplicationScheduleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReplicationScheduleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationSchedule) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReplicationScheduleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationScheduleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReplicationSchedule) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReplicationSchedulePtrOutput struct{ *pulumi.OutputState }

func (ReplicationSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationSchedule)(nil)).Elem()
}

func (o ReplicationSchedulePtrOutput) ToReplicationSchedulePtrOutput() ReplicationSchedulePtrOutput {
	return o
}

func (o ReplicationSchedulePtrOutput) ToReplicationSchedulePtrOutputWithContext(ctx context.Context) ReplicationSchedulePtrOutput {
	return o
}

func (o ReplicationSchedulePtrOutput) Elem() ReplicationScheduleOutput {
	return o.ApplyT(func(v *ReplicationSchedule) ReplicationSchedule {
		if v != nil {
			return *v
		}
		var ret ReplicationSchedule
		return ret
	}).(ReplicationScheduleOutput)
}

func (o ReplicationSchedulePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReplicationSchedulePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReplicationSchedule) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReplicationScheduleInput interface {
	pulumi.Input

	ToReplicationScheduleOutput() ReplicationScheduleOutput
	ToReplicationScheduleOutputWithContext(context.Context) ReplicationScheduleOutput
}

var replicationSchedulePtrType = reflect.TypeOf((**ReplicationSchedule)(nil)).Elem()

type ReplicationSchedulePtrInput interface {
	pulumi.Input

	ToReplicationSchedulePtrOutput() ReplicationSchedulePtrOutput
	ToReplicationSchedulePtrOutputWithContext(context.Context) ReplicationSchedulePtrOutput
}

type replicationSchedulePtr string

func ReplicationSchedulePtr(v string) ReplicationSchedulePtrInput {
	return (*replicationSchedulePtr)(&v)
}

func (*replicationSchedulePtr) ElementType() reflect.Type {
	return replicationSchedulePtrType
}

func (in *replicationSchedulePtr) ToReplicationSchedulePtrOutput() ReplicationSchedulePtrOutput {
	return pulumi.ToOutput(in).(ReplicationSchedulePtrOutput)
}

func (in *replicationSchedulePtr) ToReplicationSchedulePtrOutputWithContext(ctx context.Context) ReplicationSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReplicationSchedulePtrOutput)
}

type VolumeServiceLevel string

const (
	// Standard service level
	VolumeServiceLevelStandard = VolumeServiceLevel("Standard")
	// Premium service level
	VolumeServiceLevelPremium = VolumeServiceLevel("Premium")
	// Ultra service level
	VolumeServiceLevelUltra = VolumeServiceLevel("Ultra")
)

func (VolumeServiceLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeServiceLevel)(nil)).Elem()
}

func (e VolumeServiceLevel) ToVolumeServiceLevelOutput() VolumeServiceLevelOutput {
	return pulumi.ToOutput(e).(VolumeServiceLevelOutput)
}

func (e VolumeServiceLevel) ToVolumeServiceLevelOutputWithContext(ctx context.Context) VolumeServiceLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VolumeServiceLevelOutput)
}

func (e VolumeServiceLevel) ToVolumeServiceLevelPtrOutput() VolumeServiceLevelPtrOutput {
	return e.ToVolumeServiceLevelPtrOutputWithContext(context.Background())
}

func (e VolumeServiceLevel) ToVolumeServiceLevelPtrOutputWithContext(ctx context.Context) VolumeServiceLevelPtrOutput {
	return VolumeServiceLevel(e).ToVolumeServiceLevelOutputWithContext(ctx).ToVolumeServiceLevelPtrOutputWithContext(ctx)
}

func (e VolumeServiceLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VolumeServiceLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VolumeServiceLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VolumeServiceLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VolumeServiceLevelOutput struct{ *pulumi.OutputState }

func (VolumeServiceLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeServiceLevel)(nil)).Elem()
}

func (o VolumeServiceLevelOutput) ToVolumeServiceLevelOutput() VolumeServiceLevelOutput {
	return o
}

func (o VolumeServiceLevelOutput) ToVolumeServiceLevelOutputWithContext(ctx context.Context) VolumeServiceLevelOutput {
	return o
}

func (o VolumeServiceLevelOutput) ToVolumeServiceLevelPtrOutput() VolumeServiceLevelPtrOutput {
	return o.ToVolumeServiceLevelPtrOutputWithContext(context.Background())
}

func (o VolumeServiceLevelOutput) ToVolumeServiceLevelPtrOutputWithContext(ctx context.Context) VolumeServiceLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeServiceLevel) *VolumeServiceLevel {
		return &v
	}).(VolumeServiceLevelPtrOutput)
}

func (o VolumeServiceLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VolumeServiceLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VolumeServiceLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VolumeServiceLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VolumeServiceLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VolumeServiceLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VolumeServiceLevelPtrOutput struct{ *pulumi.OutputState }

func (VolumeServiceLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeServiceLevel)(nil)).Elem()
}

func (o VolumeServiceLevelPtrOutput) ToVolumeServiceLevelPtrOutput() VolumeServiceLevelPtrOutput {
	return o
}

func (o VolumeServiceLevelPtrOutput) ToVolumeServiceLevelPtrOutputWithContext(ctx context.Context) VolumeServiceLevelPtrOutput {
	return o
}

func (o VolumeServiceLevelPtrOutput) Elem() VolumeServiceLevelOutput {
	return o.ApplyT(func(v *VolumeServiceLevel) VolumeServiceLevel {
		if v != nil {
			return *v
		}
		var ret VolumeServiceLevel
		return ret
	}).(VolumeServiceLevelOutput)
}

func (o VolumeServiceLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VolumeServiceLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VolumeServiceLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VolumeServiceLevelInput interface {
	pulumi.Input

	ToVolumeServiceLevelOutput() VolumeServiceLevelOutput
	ToVolumeServiceLevelOutputWithContext(context.Context) VolumeServiceLevelOutput
}

var volumeServiceLevelPtrType = reflect.TypeOf((**VolumeServiceLevel)(nil)).Elem()

type VolumeServiceLevelPtrInput interface {
	pulumi.Input

	ToVolumeServiceLevelPtrOutput() VolumeServiceLevelPtrOutput
	ToVolumeServiceLevelPtrOutputWithContext(context.Context) VolumeServiceLevelPtrOutput
}

type volumeServiceLevelPtr string

func VolumeServiceLevelPtr(v string) VolumeServiceLevelPtrInput {
	return (*volumeServiceLevelPtr)(&v)
}

func (*volumeServiceLevelPtr) ElementType() reflect.Type {
	return volumeServiceLevelPtrType
}

func (in *volumeServiceLevelPtr) ToVolumeServiceLevelPtrOutput() VolumeServiceLevelPtrOutput {
	return pulumi.ToOutput(in).(VolumeServiceLevelPtrOutput)
}

func (in *volumeServiceLevelPtr) ToVolumeServiceLevelPtrOutputWithContext(ctx context.Context) VolumeServiceLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VolumeServiceLevelPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointTypeInput)(nil)).Elem(), EndpointType("src"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointTypePtrInput)(nil)).Elem(), EndpointType("src"))
	pulumi.RegisterInputType(reflect.TypeOf((*PoolServiceLevelInput)(nil)).Elem(), PoolServiceLevel("Standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*PoolServiceLevelPtrInput)(nil)).Elem(), PoolServiceLevel("Standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationScheduleInput)(nil)).Elem(), ReplicationSchedule("_10minutely"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationSchedulePtrInput)(nil)).Elem(), ReplicationSchedule("_10minutely"))
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeServiceLevelInput)(nil)).Elem(), VolumeServiceLevel("Standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeServiceLevelPtrInput)(nil)).Elem(), VolumeServiceLevel("Standard"))
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypePtrOutput{})
	pulumi.RegisterOutputType(PoolServiceLevelOutput{})
	pulumi.RegisterOutputType(PoolServiceLevelPtrOutput{})
	pulumi.RegisterOutputType(ReplicationScheduleOutput{})
	pulumi.RegisterOutputType(ReplicationSchedulePtrOutput{})
	pulumi.RegisterOutputType(VolumeServiceLevelOutput{})
	pulumi.RegisterOutputType(VolumeServiceLevelPtrOutput{})
}
