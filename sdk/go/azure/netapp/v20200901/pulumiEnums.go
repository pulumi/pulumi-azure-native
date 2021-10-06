


package v20200901

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

type QosType string

const (
	// qos type Auto
	QosTypeAuto = QosType("Auto")
	// qos type Manual
	QosTypeManual = QosType("Manual")
)

func (QosType) ElementType() reflect.Type {
	return reflect.TypeOf((*QosType)(nil)).Elem()
}

func (e QosType) ToQosTypeOutput() QosTypeOutput {
	return pulumi.ToOutput(e).(QosTypeOutput)
}

func (e QosType) ToQosTypeOutputWithContext(ctx context.Context) QosTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(QosTypeOutput)
}

func (e QosType) ToQosTypePtrOutput() QosTypePtrOutput {
	return e.ToQosTypePtrOutputWithContext(context.Background())
}

func (e QosType) ToQosTypePtrOutputWithContext(ctx context.Context) QosTypePtrOutput {
	return QosType(e).ToQosTypeOutputWithContext(ctx).ToQosTypePtrOutputWithContext(ctx)
}

func (e QosType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e QosType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e QosType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e QosType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type QosTypeOutput struct{ *pulumi.OutputState }

func (QosTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QosType)(nil)).Elem()
}

func (o QosTypeOutput) ToQosTypeOutput() QosTypeOutput {
	return o
}

func (o QosTypeOutput) ToQosTypeOutputWithContext(ctx context.Context) QosTypeOutput {
	return o
}

func (o QosTypeOutput) ToQosTypePtrOutput() QosTypePtrOutput {
	return o.ToQosTypePtrOutputWithContext(context.Background())
}

func (o QosTypeOutput) ToQosTypePtrOutputWithContext(ctx context.Context) QosTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QosType) *QosType {
		return &v
	}).(QosTypePtrOutput)
}

func (o QosTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o QosTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e QosType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o QosTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o QosTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e QosType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type QosTypePtrOutput struct{ *pulumi.OutputState }

func (QosTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QosType)(nil)).Elem()
}

func (o QosTypePtrOutput) ToQosTypePtrOutput() QosTypePtrOutput {
	return o
}

func (o QosTypePtrOutput) ToQosTypePtrOutputWithContext(ctx context.Context) QosTypePtrOutput {
	return o
}

func (o QosTypePtrOutput) Elem() QosTypeOutput {
	return o.ApplyT(func(v *QosType) QosType {
		if v != nil {
			return *v
		}
		var ret QosType
		return ret
	}).(QosTypeOutput)
}

func (o QosTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o QosTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *QosType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type QosTypeInput interface {
	pulumi.Input

	ToQosTypeOutput() QosTypeOutput
	ToQosTypeOutputWithContext(context.Context) QosTypeOutput
}

var qosTypePtrType = reflect.TypeOf((**QosType)(nil)).Elem()

type QosTypePtrInput interface {
	pulumi.Input

	ToQosTypePtrOutput() QosTypePtrOutput
	ToQosTypePtrOutputWithContext(context.Context) QosTypePtrOutput
}

type qosTypePtr string

func QosTypePtr(v string) QosTypePtrInput {
	return (*qosTypePtr)(&v)
}

func (*qosTypePtr) ElementType() reflect.Type {
	return qosTypePtrType
}

func (in *qosTypePtr) ToQosTypePtrOutput() QosTypePtrOutput {
	return pulumi.ToOutput(in).(QosTypePtrOutput)
}

func (in *qosTypePtr) ToQosTypePtrOutputWithContext(ctx context.Context) QosTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(QosTypePtrOutput)
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

type SecurityStyle string

const (
	SecurityStyleNtfs = SecurityStyle("ntfs")
	SecurityStyleUnix = SecurityStyle("unix")
)

func (SecurityStyle) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityStyle)(nil)).Elem()
}

func (e SecurityStyle) ToSecurityStyleOutput() SecurityStyleOutput {
	return pulumi.ToOutput(e).(SecurityStyleOutput)
}

func (e SecurityStyle) ToSecurityStyleOutputWithContext(ctx context.Context) SecurityStyleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecurityStyleOutput)
}

func (e SecurityStyle) ToSecurityStylePtrOutput() SecurityStylePtrOutput {
	return e.ToSecurityStylePtrOutputWithContext(context.Background())
}

func (e SecurityStyle) ToSecurityStylePtrOutputWithContext(ctx context.Context) SecurityStylePtrOutput {
	return SecurityStyle(e).ToSecurityStyleOutputWithContext(ctx).ToSecurityStylePtrOutputWithContext(ctx)
}

func (e SecurityStyle) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityStyle) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecurityStyle) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecurityStyle) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecurityStyleOutput struct{ *pulumi.OutputState }

func (SecurityStyleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityStyle)(nil)).Elem()
}

func (o SecurityStyleOutput) ToSecurityStyleOutput() SecurityStyleOutput {
	return o
}

func (o SecurityStyleOutput) ToSecurityStyleOutputWithContext(ctx context.Context) SecurityStyleOutput {
	return o
}

func (o SecurityStyleOutput) ToSecurityStylePtrOutput() SecurityStylePtrOutput {
	return o.ToSecurityStylePtrOutputWithContext(context.Background())
}

func (o SecurityStyleOutput) ToSecurityStylePtrOutputWithContext(ctx context.Context) SecurityStylePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityStyle) *SecurityStyle {
		return &v
	}).(SecurityStylePtrOutput)
}

func (o SecurityStyleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecurityStyleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityStyle) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecurityStyleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityStyleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecurityStyle) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecurityStylePtrOutput struct{ *pulumi.OutputState }

func (SecurityStylePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityStyle)(nil)).Elem()
}

func (o SecurityStylePtrOutput) ToSecurityStylePtrOutput() SecurityStylePtrOutput {
	return o
}

func (o SecurityStylePtrOutput) ToSecurityStylePtrOutputWithContext(ctx context.Context) SecurityStylePtrOutput {
	return o
}

func (o SecurityStylePtrOutput) Elem() SecurityStyleOutput {
	return o.ApplyT(func(v *SecurityStyle) SecurityStyle {
		if v != nil {
			return *v
		}
		var ret SecurityStyle
		return ret
	}).(SecurityStyleOutput)
}

func (o SecurityStylePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecurityStylePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecurityStyle) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecurityStyleInput interface {
	pulumi.Input

	ToSecurityStyleOutput() SecurityStyleOutput
	ToSecurityStyleOutputWithContext(context.Context) SecurityStyleOutput
}

var securityStylePtrType = reflect.TypeOf((**SecurityStyle)(nil)).Elem()

type SecurityStylePtrInput interface {
	pulumi.Input

	ToSecurityStylePtrOutput() SecurityStylePtrOutput
	ToSecurityStylePtrOutputWithContext(context.Context) SecurityStylePtrOutput
}

type securityStylePtr string

func SecurityStylePtr(v string) SecurityStylePtrInput {
	return (*securityStylePtr)(&v)
}

func (*securityStylePtr) ElementType() reflect.Type {
	return securityStylePtrType
}

func (in *securityStylePtr) ToSecurityStylePtrOutput() SecurityStylePtrOutput {
	return pulumi.ToOutput(in).(SecurityStylePtrOutput)
}

func (in *securityStylePtr) ToSecurityStylePtrOutputWithContext(ctx context.Context) SecurityStylePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecurityStylePtrOutput)
}

type ServiceLevel string

const (
	// Standard service level
	ServiceLevelStandard = ServiceLevel("Standard")
	// Premium service level
	ServiceLevelPremium = ServiceLevel("Premium")
	// Ultra service level
	ServiceLevelUltra = ServiceLevel("Ultra")
)

func (ServiceLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLevel)(nil)).Elem()
}

func (e ServiceLevel) ToServiceLevelOutput() ServiceLevelOutput {
	return pulumi.ToOutput(e).(ServiceLevelOutput)
}

func (e ServiceLevel) ToServiceLevelOutputWithContext(ctx context.Context) ServiceLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceLevelOutput)
}

func (e ServiceLevel) ToServiceLevelPtrOutput() ServiceLevelPtrOutput {
	return e.ToServiceLevelPtrOutputWithContext(context.Background())
}

func (e ServiceLevel) ToServiceLevelPtrOutputWithContext(ctx context.Context) ServiceLevelPtrOutput {
	return ServiceLevel(e).ToServiceLevelOutputWithContext(ctx).ToServiceLevelPtrOutputWithContext(ctx)
}

func (e ServiceLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceLevelOutput struct{ *pulumi.OutputState }

func (ServiceLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLevel)(nil)).Elem()
}

func (o ServiceLevelOutput) ToServiceLevelOutput() ServiceLevelOutput {
	return o
}

func (o ServiceLevelOutput) ToServiceLevelOutputWithContext(ctx context.Context) ServiceLevelOutput {
	return o
}

func (o ServiceLevelOutput) ToServiceLevelPtrOutput() ServiceLevelPtrOutput {
	return o.ToServiceLevelPtrOutputWithContext(context.Background())
}

func (o ServiceLevelOutput) ToServiceLevelPtrOutputWithContext(ctx context.Context) ServiceLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceLevel) *ServiceLevel {
		return &v
	}).(ServiceLevelPtrOutput)
}

func (o ServiceLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceLevelPtrOutput struct{ *pulumi.OutputState }

func (ServiceLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceLevel)(nil)).Elem()
}

func (o ServiceLevelPtrOutput) ToServiceLevelPtrOutput() ServiceLevelPtrOutput {
	return o
}

func (o ServiceLevelPtrOutput) ToServiceLevelPtrOutputWithContext(ctx context.Context) ServiceLevelPtrOutput {
	return o
}

func (o ServiceLevelPtrOutput) Elem() ServiceLevelOutput {
	return o.ApplyT(func(v *ServiceLevel) ServiceLevel {
		if v != nil {
			return *v
		}
		var ret ServiceLevel
		return ret
	}).(ServiceLevelOutput)
}

func (o ServiceLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceLevelInput interface {
	pulumi.Input

	ToServiceLevelOutput() ServiceLevelOutput
	ToServiceLevelOutputWithContext(context.Context) ServiceLevelOutput
}

var serviceLevelPtrType = reflect.TypeOf((**ServiceLevel)(nil)).Elem()

type ServiceLevelPtrInput interface {
	pulumi.Input

	ToServiceLevelPtrOutput() ServiceLevelPtrOutput
	ToServiceLevelPtrOutputWithContext(context.Context) ServiceLevelPtrOutput
}

type serviceLevelPtr string

func ServiceLevelPtr(v string) ServiceLevelPtrInput {
	return (*serviceLevelPtr)(&v)
}

func (*serviceLevelPtr) ElementType() reflect.Type {
	return serviceLevelPtrType
}

func (in *serviceLevelPtr) ToServiceLevelPtrOutput() ServiceLevelPtrOutput {
	return pulumi.ToOutput(in).(ServiceLevelPtrOutput)
}

func (in *serviceLevelPtr) ToServiceLevelPtrOutputWithContext(ctx context.Context) ServiceLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceLevelPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypePtrOutput{})
	pulumi.RegisterOutputType(QosTypeOutput{})
	pulumi.RegisterOutputType(QosTypePtrOutput{})
	pulumi.RegisterOutputType(ReplicationScheduleOutput{})
	pulumi.RegisterOutputType(ReplicationSchedulePtrOutput{})
	pulumi.RegisterOutputType(SecurityStyleOutput{})
	pulumi.RegisterOutputType(SecurityStylePtrOutput{})
	pulumi.RegisterOutputType(ServiceLevelOutput{})
	pulumi.RegisterOutputType(ServiceLevelPtrOutput{})
}
