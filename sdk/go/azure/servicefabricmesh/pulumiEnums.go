


package servicefabricmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationScopedVolumeKind string

const (
	// Provides Service Fabric High Availability Volume Disk
	ApplicationScopedVolumeKindServiceFabricVolumeDisk = ApplicationScopedVolumeKind("ServiceFabricVolumeDisk")
)

func (ApplicationScopedVolumeKind) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationScopedVolumeKind)(nil)).Elem()
}

func (e ApplicationScopedVolumeKind) ToApplicationScopedVolumeKindOutput() ApplicationScopedVolumeKindOutput {
	return pulumi.ToOutput(e).(ApplicationScopedVolumeKindOutput)
}

func (e ApplicationScopedVolumeKind) ToApplicationScopedVolumeKindOutputWithContext(ctx context.Context) ApplicationScopedVolumeKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationScopedVolumeKindOutput)
}

func (e ApplicationScopedVolumeKind) ToApplicationScopedVolumeKindPtrOutput() ApplicationScopedVolumeKindPtrOutput {
	return e.ToApplicationScopedVolumeKindPtrOutputWithContext(context.Background())
}

func (e ApplicationScopedVolumeKind) ToApplicationScopedVolumeKindPtrOutputWithContext(ctx context.Context) ApplicationScopedVolumeKindPtrOutput {
	return ApplicationScopedVolumeKind(e).ToApplicationScopedVolumeKindOutputWithContext(ctx).ToApplicationScopedVolumeKindPtrOutputWithContext(ctx)
}

func (e ApplicationScopedVolumeKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationScopedVolumeKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationScopedVolumeKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationScopedVolumeKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationScopedVolumeKindOutput struct{ *pulumi.OutputState }

func (ApplicationScopedVolumeKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationScopedVolumeKind)(nil)).Elem()
}

func (o ApplicationScopedVolumeKindOutput) ToApplicationScopedVolumeKindOutput() ApplicationScopedVolumeKindOutput {
	return o
}

func (o ApplicationScopedVolumeKindOutput) ToApplicationScopedVolumeKindOutputWithContext(ctx context.Context) ApplicationScopedVolumeKindOutput {
	return o
}

func (o ApplicationScopedVolumeKindOutput) ToApplicationScopedVolumeKindPtrOutput() ApplicationScopedVolumeKindPtrOutput {
	return o.ToApplicationScopedVolumeKindPtrOutputWithContext(context.Background())
}

func (o ApplicationScopedVolumeKindOutput) ToApplicationScopedVolumeKindPtrOutputWithContext(ctx context.Context) ApplicationScopedVolumeKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationScopedVolumeKind) *ApplicationScopedVolumeKind {
		return &v
	}).(ApplicationScopedVolumeKindPtrOutput)
}

func (o ApplicationScopedVolumeKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationScopedVolumeKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationScopedVolumeKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationScopedVolumeKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationScopedVolumeKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationScopedVolumeKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationScopedVolumeKindPtrOutput struct{ *pulumi.OutputState }

func (ApplicationScopedVolumeKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationScopedVolumeKind)(nil)).Elem()
}

func (o ApplicationScopedVolumeKindPtrOutput) ToApplicationScopedVolumeKindPtrOutput() ApplicationScopedVolumeKindPtrOutput {
	return o
}

func (o ApplicationScopedVolumeKindPtrOutput) ToApplicationScopedVolumeKindPtrOutputWithContext(ctx context.Context) ApplicationScopedVolumeKindPtrOutput {
	return o
}

func (o ApplicationScopedVolumeKindPtrOutput) Elem() ApplicationScopedVolumeKindOutput {
	return o.ApplyT(func(v *ApplicationScopedVolumeKind) ApplicationScopedVolumeKind {
		if v != nil {
			return *v
		}
		var ret ApplicationScopedVolumeKind
		return ret
	}).(ApplicationScopedVolumeKindOutput)
}

func (o ApplicationScopedVolumeKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationScopedVolumeKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationScopedVolumeKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationScopedVolumeKindInput interface {
	pulumi.Input

	ToApplicationScopedVolumeKindOutput() ApplicationScopedVolumeKindOutput
	ToApplicationScopedVolumeKindOutputWithContext(context.Context) ApplicationScopedVolumeKindOutput
}

var applicationScopedVolumeKindPtrType = reflect.TypeOf((**ApplicationScopedVolumeKind)(nil)).Elem()

type ApplicationScopedVolumeKindPtrInput interface {
	pulumi.Input

	ToApplicationScopedVolumeKindPtrOutput() ApplicationScopedVolumeKindPtrOutput
	ToApplicationScopedVolumeKindPtrOutputWithContext(context.Context) ApplicationScopedVolumeKindPtrOutput
}

type applicationScopedVolumeKindPtr string

func ApplicationScopedVolumeKindPtr(v string) ApplicationScopedVolumeKindPtrInput {
	return (*applicationScopedVolumeKindPtr)(&v)
}

func (*applicationScopedVolumeKindPtr) ElementType() reflect.Type {
	return applicationScopedVolumeKindPtrType
}

func (in *applicationScopedVolumeKindPtr) ToApplicationScopedVolumeKindPtrOutput() ApplicationScopedVolumeKindPtrOutput {
	return pulumi.ToOutput(in).(ApplicationScopedVolumeKindPtrOutput)
}

func (in *applicationScopedVolumeKindPtr) ToApplicationScopedVolumeKindPtrOutputWithContext(ctx context.Context) ApplicationScopedVolumeKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationScopedVolumeKindPtrOutput)
}

type AutoScalingMechanismKind string

const (
	// Indicates that scaling should be performed by adding or removing replicas.
	AutoScalingMechanismKindAddRemoveReplica = AutoScalingMechanismKind("AddRemoveReplica")
)

func (AutoScalingMechanismKind) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingMechanismKind)(nil)).Elem()
}

func (e AutoScalingMechanismKind) ToAutoScalingMechanismKindOutput() AutoScalingMechanismKindOutput {
	return pulumi.ToOutput(e).(AutoScalingMechanismKindOutput)
}

func (e AutoScalingMechanismKind) ToAutoScalingMechanismKindOutputWithContext(ctx context.Context) AutoScalingMechanismKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoScalingMechanismKindOutput)
}

func (e AutoScalingMechanismKind) ToAutoScalingMechanismKindPtrOutput() AutoScalingMechanismKindPtrOutput {
	return e.ToAutoScalingMechanismKindPtrOutputWithContext(context.Background())
}

func (e AutoScalingMechanismKind) ToAutoScalingMechanismKindPtrOutputWithContext(ctx context.Context) AutoScalingMechanismKindPtrOutput {
	return AutoScalingMechanismKind(e).ToAutoScalingMechanismKindOutputWithContext(ctx).ToAutoScalingMechanismKindPtrOutputWithContext(ctx)
}

func (e AutoScalingMechanismKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoScalingMechanismKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoScalingMechanismKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoScalingMechanismKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoScalingMechanismKindOutput struct{ *pulumi.OutputState }

func (AutoScalingMechanismKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingMechanismKind)(nil)).Elem()
}

func (o AutoScalingMechanismKindOutput) ToAutoScalingMechanismKindOutput() AutoScalingMechanismKindOutput {
	return o
}

func (o AutoScalingMechanismKindOutput) ToAutoScalingMechanismKindOutputWithContext(ctx context.Context) AutoScalingMechanismKindOutput {
	return o
}

func (o AutoScalingMechanismKindOutput) ToAutoScalingMechanismKindPtrOutput() AutoScalingMechanismKindPtrOutput {
	return o.ToAutoScalingMechanismKindPtrOutputWithContext(context.Background())
}

func (o AutoScalingMechanismKindOutput) ToAutoScalingMechanismKindPtrOutputWithContext(ctx context.Context) AutoScalingMechanismKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScalingMechanismKind) *AutoScalingMechanismKind {
		return &v
	}).(AutoScalingMechanismKindPtrOutput)
}

func (o AutoScalingMechanismKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoScalingMechanismKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoScalingMechanismKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoScalingMechanismKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoScalingMechanismKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoScalingMechanismKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoScalingMechanismKindPtrOutput struct{ *pulumi.OutputState }

func (AutoScalingMechanismKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalingMechanismKind)(nil)).Elem()
}

func (o AutoScalingMechanismKindPtrOutput) ToAutoScalingMechanismKindPtrOutput() AutoScalingMechanismKindPtrOutput {
	return o
}

func (o AutoScalingMechanismKindPtrOutput) ToAutoScalingMechanismKindPtrOutputWithContext(ctx context.Context) AutoScalingMechanismKindPtrOutput {
	return o
}

func (o AutoScalingMechanismKindPtrOutput) Elem() AutoScalingMechanismKindOutput {
	return o.ApplyT(func(v *AutoScalingMechanismKind) AutoScalingMechanismKind {
		if v != nil {
			return *v
		}
		var ret AutoScalingMechanismKind
		return ret
	}).(AutoScalingMechanismKindOutput)
}

func (o AutoScalingMechanismKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoScalingMechanismKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoScalingMechanismKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoScalingMechanismKindInput interface {
	pulumi.Input

	ToAutoScalingMechanismKindOutput() AutoScalingMechanismKindOutput
	ToAutoScalingMechanismKindOutputWithContext(context.Context) AutoScalingMechanismKindOutput
}

var autoScalingMechanismKindPtrType = reflect.TypeOf((**AutoScalingMechanismKind)(nil)).Elem()

type AutoScalingMechanismKindPtrInput interface {
	pulumi.Input

	ToAutoScalingMechanismKindPtrOutput() AutoScalingMechanismKindPtrOutput
	ToAutoScalingMechanismKindPtrOutputWithContext(context.Context) AutoScalingMechanismKindPtrOutput
}

type autoScalingMechanismKindPtr string

func AutoScalingMechanismKindPtr(v string) AutoScalingMechanismKindPtrInput {
	return (*autoScalingMechanismKindPtr)(&v)
}

func (*autoScalingMechanismKindPtr) ElementType() reflect.Type {
	return autoScalingMechanismKindPtrType
}

func (in *autoScalingMechanismKindPtr) ToAutoScalingMechanismKindPtrOutput() AutoScalingMechanismKindPtrOutput {
	return pulumi.ToOutput(in).(AutoScalingMechanismKindPtrOutput)
}

func (in *autoScalingMechanismKindPtr) ToAutoScalingMechanismKindPtrOutputWithContext(ctx context.Context) AutoScalingMechanismKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoScalingMechanismKindPtrOutput)
}

type AutoScalingMetricKind string

const (
	// Indicates that the metric is one of resources, like cpu or memory.
	AutoScalingMetricKindResource = AutoScalingMetricKind("Resource")
)

func (AutoScalingMetricKind) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingMetricKind)(nil)).Elem()
}

func (e AutoScalingMetricKind) ToAutoScalingMetricKindOutput() AutoScalingMetricKindOutput {
	return pulumi.ToOutput(e).(AutoScalingMetricKindOutput)
}

func (e AutoScalingMetricKind) ToAutoScalingMetricKindOutputWithContext(ctx context.Context) AutoScalingMetricKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoScalingMetricKindOutput)
}

func (e AutoScalingMetricKind) ToAutoScalingMetricKindPtrOutput() AutoScalingMetricKindPtrOutput {
	return e.ToAutoScalingMetricKindPtrOutputWithContext(context.Background())
}

func (e AutoScalingMetricKind) ToAutoScalingMetricKindPtrOutputWithContext(ctx context.Context) AutoScalingMetricKindPtrOutput {
	return AutoScalingMetricKind(e).ToAutoScalingMetricKindOutputWithContext(ctx).ToAutoScalingMetricKindPtrOutputWithContext(ctx)
}

func (e AutoScalingMetricKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoScalingMetricKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoScalingMetricKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoScalingMetricKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoScalingMetricKindOutput struct{ *pulumi.OutputState }

func (AutoScalingMetricKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingMetricKind)(nil)).Elem()
}

func (o AutoScalingMetricKindOutput) ToAutoScalingMetricKindOutput() AutoScalingMetricKindOutput {
	return o
}

func (o AutoScalingMetricKindOutput) ToAutoScalingMetricKindOutputWithContext(ctx context.Context) AutoScalingMetricKindOutput {
	return o
}

func (o AutoScalingMetricKindOutput) ToAutoScalingMetricKindPtrOutput() AutoScalingMetricKindPtrOutput {
	return o.ToAutoScalingMetricKindPtrOutputWithContext(context.Background())
}

func (o AutoScalingMetricKindOutput) ToAutoScalingMetricKindPtrOutputWithContext(ctx context.Context) AutoScalingMetricKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScalingMetricKind) *AutoScalingMetricKind {
		return &v
	}).(AutoScalingMetricKindPtrOutput)
}

func (o AutoScalingMetricKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoScalingMetricKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoScalingMetricKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoScalingMetricKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoScalingMetricKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoScalingMetricKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoScalingMetricKindPtrOutput struct{ *pulumi.OutputState }

func (AutoScalingMetricKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalingMetricKind)(nil)).Elem()
}

func (o AutoScalingMetricKindPtrOutput) ToAutoScalingMetricKindPtrOutput() AutoScalingMetricKindPtrOutput {
	return o
}

func (o AutoScalingMetricKindPtrOutput) ToAutoScalingMetricKindPtrOutputWithContext(ctx context.Context) AutoScalingMetricKindPtrOutput {
	return o
}

func (o AutoScalingMetricKindPtrOutput) Elem() AutoScalingMetricKindOutput {
	return o.ApplyT(func(v *AutoScalingMetricKind) AutoScalingMetricKind {
		if v != nil {
			return *v
		}
		var ret AutoScalingMetricKind
		return ret
	}).(AutoScalingMetricKindOutput)
}

func (o AutoScalingMetricKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoScalingMetricKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoScalingMetricKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoScalingMetricKindInput interface {
	pulumi.Input

	ToAutoScalingMetricKindOutput() AutoScalingMetricKindOutput
	ToAutoScalingMetricKindOutputWithContext(context.Context) AutoScalingMetricKindOutput
}

var autoScalingMetricKindPtrType = reflect.TypeOf((**AutoScalingMetricKind)(nil)).Elem()

type AutoScalingMetricKindPtrInput interface {
	pulumi.Input

	ToAutoScalingMetricKindPtrOutput() AutoScalingMetricKindPtrOutput
	ToAutoScalingMetricKindPtrOutputWithContext(context.Context) AutoScalingMetricKindPtrOutput
}

type autoScalingMetricKindPtr string

func AutoScalingMetricKindPtr(v string) AutoScalingMetricKindPtrInput {
	return (*autoScalingMetricKindPtr)(&v)
}

func (*autoScalingMetricKindPtr) ElementType() reflect.Type {
	return autoScalingMetricKindPtrType
}

func (in *autoScalingMetricKindPtr) ToAutoScalingMetricKindPtrOutput() AutoScalingMetricKindPtrOutput {
	return pulumi.ToOutput(in).(AutoScalingMetricKindPtrOutput)
}

func (in *autoScalingMetricKindPtr) ToAutoScalingMetricKindPtrOutputWithContext(ctx context.Context) AutoScalingMetricKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoScalingMetricKindPtrOutput)
}

type AutoScalingResourceMetricName string

const (
	// Indicates that the resource is CPU cores.
	AutoScalingResourceMetricNameCpu = AutoScalingResourceMetricName("cpu")
	// Indicates that the resource is memory in GB.
	AutoScalingResourceMetricNameMemoryInGB = AutoScalingResourceMetricName("memoryInGB")
)

func (AutoScalingResourceMetricName) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingResourceMetricName)(nil)).Elem()
}

func (e AutoScalingResourceMetricName) ToAutoScalingResourceMetricNameOutput() AutoScalingResourceMetricNameOutput {
	return pulumi.ToOutput(e).(AutoScalingResourceMetricNameOutput)
}

func (e AutoScalingResourceMetricName) ToAutoScalingResourceMetricNameOutputWithContext(ctx context.Context) AutoScalingResourceMetricNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoScalingResourceMetricNameOutput)
}

func (e AutoScalingResourceMetricName) ToAutoScalingResourceMetricNamePtrOutput() AutoScalingResourceMetricNamePtrOutput {
	return e.ToAutoScalingResourceMetricNamePtrOutputWithContext(context.Background())
}

func (e AutoScalingResourceMetricName) ToAutoScalingResourceMetricNamePtrOutputWithContext(ctx context.Context) AutoScalingResourceMetricNamePtrOutput {
	return AutoScalingResourceMetricName(e).ToAutoScalingResourceMetricNameOutputWithContext(ctx).ToAutoScalingResourceMetricNamePtrOutputWithContext(ctx)
}

func (e AutoScalingResourceMetricName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoScalingResourceMetricName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoScalingResourceMetricName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoScalingResourceMetricName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoScalingResourceMetricNameOutput struct{ *pulumi.OutputState }

func (AutoScalingResourceMetricNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingResourceMetricName)(nil)).Elem()
}

func (o AutoScalingResourceMetricNameOutput) ToAutoScalingResourceMetricNameOutput() AutoScalingResourceMetricNameOutput {
	return o
}

func (o AutoScalingResourceMetricNameOutput) ToAutoScalingResourceMetricNameOutputWithContext(ctx context.Context) AutoScalingResourceMetricNameOutput {
	return o
}

func (o AutoScalingResourceMetricNameOutput) ToAutoScalingResourceMetricNamePtrOutput() AutoScalingResourceMetricNamePtrOutput {
	return o.ToAutoScalingResourceMetricNamePtrOutputWithContext(context.Background())
}

func (o AutoScalingResourceMetricNameOutput) ToAutoScalingResourceMetricNamePtrOutputWithContext(ctx context.Context) AutoScalingResourceMetricNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScalingResourceMetricName) *AutoScalingResourceMetricName {
		return &v
	}).(AutoScalingResourceMetricNamePtrOutput)
}

func (o AutoScalingResourceMetricNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoScalingResourceMetricNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoScalingResourceMetricName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoScalingResourceMetricNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoScalingResourceMetricNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoScalingResourceMetricName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoScalingResourceMetricNamePtrOutput struct{ *pulumi.OutputState }

func (AutoScalingResourceMetricNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalingResourceMetricName)(nil)).Elem()
}

func (o AutoScalingResourceMetricNamePtrOutput) ToAutoScalingResourceMetricNamePtrOutput() AutoScalingResourceMetricNamePtrOutput {
	return o
}

func (o AutoScalingResourceMetricNamePtrOutput) ToAutoScalingResourceMetricNamePtrOutputWithContext(ctx context.Context) AutoScalingResourceMetricNamePtrOutput {
	return o
}

func (o AutoScalingResourceMetricNamePtrOutput) Elem() AutoScalingResourceMetricNameOutput {
	return o.ApplyT(func(v *AutoScalingResourceMetricName) AutoScalingResourceMetricName {
		if v != nil {
			return *v
		}
		var ret AutoScalingResourceMetricName
		return ret
	}).(AutoScalingResourceMetricNameOutput)
}

func (o AutoScalingResourceMetricNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoScalingResourceMetricNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoScalingResourceMetricName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoScalingResourceMetricNameInput interface {
	pulumi.Input

	ToAutoScalingResourceMetricNameOutput() AutoScalingResourceMetricNameOutput
	ToAutoScalingResourceMetricNameOutputWithContext(context.Context) AutoScalingResourceMetricNameOutput
}

var autoScalingResourceMetricNamePtrType = reflect.TypeOf((**AutoScalingResourceMetricName)(nil)).Elem()

type AutoScalingResourceMetricNamePtrInput interface {
	pulumi.Input

	ToAutoScalingResourceMetricNamePtrOutput() AutoScalingResourceMetricNamePtrOutput
	ToAutoScalingResourceMetricNamePtrOutputWithContext(context.Context) AutoScalingResourceMetricNamePtrOutput
}

type autoScalingResourceMetricNamePtr string

func AutoScalingResourceMetricNamePtr(v string) AutoScalingResourceMetricNamePtrInput {
	return (*autoScalingResourceMetricNamePtr)(&v)
}

func (*autoScalingResourceMetricNamePtr) ElementType() reflect.Type {
	return autoScalingResourceMetricNamePtrType
}

func (in *autoScalingResourceMetricNamePtr) ToAutoScalingResourceMetricNamePtrOutput() AutoScalingResourceMetricNamePtrOutput {
	return pulumi.ToOutput(in).(AutoScalingResourceMetricNamePtrOutput)
}

func (in *autoScalingResourceMetricNamePtr) ToAutoScalingResourceMetricNamePtrOutputWithContext(ctx context.Context) AutoScalingResourceMetricNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoScalingResourceMetricNamePtrOutput)
}

type AutoScalingTriggerKind string

const (
	// Indicates that scaling should be performed based on average load of all replicas in the service.
	AutoScalingTriggerKindAverageLoad = AutoScalingTriggerKind("AverageLoad")
)

func (AutoScalingTriggerKind) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingTriggerKind)(nil)).Elem()
}

func (e AutoScalingTriggerKind) ToAutoScalingTriggerKindOutput() AutoScalingTriggerKindOutput {
	return pulumi.ToOutput(e).(AutoScalingTriggerKindOutput)
}

func (e AutoScalingTriggerKind) ToAutoScalingTriggerKindOutputWithContext(ctx context.Context) AutoScalingTriggerKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoScalingTriggerKindOutput)
}

func (e AutoScalingTriggerKind) ToAutoScalingTriggerKindPtrOutput() AutoScalingTriggerKindPtrOutput {
	return e.ToAutoScalingTriggerKindPtrOutputWithContext(context.Background())
}

func (e AutoScalingTriggerKind) ToAutoScalingTriggerKindPtrOutputWithContext(ctx context.Context) AutoScalingTriggerKindPtrOutput {
	return AutoScalingTriggerKind(e).ToAutoScalingTriggerKindOutputWithContext(ctx).ToAutoScalingTriggerKindPtrOutputWithContext(ctx)
}

func (e AutoScalingTriggerKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoScalingTriggerKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoScalingTriggerKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoScalingTriggerKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoScalingTriggerKindOutput struct{ *pulumi.OutputState }

func (AutoScalingTriggerKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingTriggerKind)(nil)).Elem()
}

func (o AutoScalingTriggerKindOutput) ToAutoScalingTriggerKindOutput() AutoScalingTriggerKindOutput {
	return o
}

func (o AutoScalingTriggerKindOutput) ToAutoScalingTriggerKindOutputWithContext(ctx context.Context) AutoScalingTriggerKindOutput {
	return o
}

func (o AutoScalingTriggerKindOutput) ToAutoScalingTriggerKindPtrOutput() AutoScalingTriggerKindPtrOutput {
	return o.ToAutoScalingTriggerKindPtrOutputWithContext(context.Background())
}

func (o AutoScalingTriggerKindOutput) ToAutoScalingTriggerKindPtrOutputWithContext(ctx context.Context) AutoScalingTriggerKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoScalingTriggerKind) *AutoScalingTriggerKind {
		return &v
	}).(AutoScalingTriggerKindPtrOutput)
}

func (o AutoScalingTriggerKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoScalingTriggerKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoScalingTriggerKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoScalingTriggerKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoScalingTriggerKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoScalingTriggerKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoScalingTriggerKindPtrOutput struct{ *pulumi.OutputState }

func (AutoScalingTriggerKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoScalingTriggerKind)(nil)).Elem()
}

func (o AutoScalingTriggerKindPtrOutput) ToAutoScalingTriggerKindPtrOutput() AutoScalingTriggerKindPtrOutput {
	return o
}

func (o AutoScalingTriggerKindPtrOutput) ToAutoScalingTriggerKindPtrOutputWithContext(ctx context.Context) AutoScalingTriggerKindPtrOutput {
	return o
}

func (o AutoScalingTriggerKindPtrOutput) Elem() AutoScalingTriggerKindOutput {
	return o.ApplyT(func(v *AutoScalingTriggerKind) AutoScalingTriggerKind {
		if v != nil {
			return *v
		}
		var ret AutoScalingTriggerKind
		return ret
	}).(AutoScalingTriggerKindOutput)
}

func (o AutoScalingTriggerKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoScalingTriggerKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoScalingTriggerKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoScalingTriggerKindInput interface {
	pulumi.Input

	ToAutoScalingTriggerKindOutput() AutoScalingTriggerKindOutput
	ToAutoScalingTriggerKindOutputWithContext(context.Context) AutoScalingTriggerKindOutput
}

var autoScalingTriggerKindPtrType = reflect.TypeOf((**AutoScalingTriggerKind)(nil)).Elem()

type AutoScalingTriggerKindPtrInput interface {
	pulumi.Input

	ToAutoScalingTriggerKindPtrOutput() AutoScalingTriggerKindPtrOutput
	ToAutoScalingTriggerKindPtrOutputWithContext(context.Context) AutoScalingTriggerKindPtrOutput
}

type autoScalingTriggerKindPtr string

func AutoScalingTriggerKindPtr(v string) AutoScalingTriggerKindPtrInput {
	return (*autoScalingTriggerKindPtr)(&v)
}

func (*autoScalingTriggerKindPtr) ElementType() reflect.Type {
	return autoScalingTriggerKindPtrType
}

func (in *autoScalingTriggerKindPtr) ToAutoScalingTriggerKindPtrOutput() AutoScalingTriggerKindPtrOutput {
	return pulumi.ToOutput(in).(AutoScalingTriggerKindPtrOutput)
}

func (in *autoScalingTriggerKindPtr) ToAutoScalingTriggerKindPtrOutputWithContext(ctx context.Context) AutoScalingTriggerKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoScalingTriggerKindPtrOutput)
}

type DiagnosticsSinkKind string

const (
	// Indicates an invalid sink kind. All Service Fabric enumerations have the invalid type.
	DiagnosticsSinkKindInvalid = DiagnosticsSinkKind("Invalid")
	// Diagnostics settings for Geneva.
	DiagnosticsSinkKindAzureInternalMonitoringPipeline = DiagnosticsSinkKind("AzureInternalMonitoringPipeline")
)

func (DiagnosticsSinkKind) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsSinkKind)(nil)).Elem()
}

func (e DiagnosticsSinkKind) ToDiagnosticsSinkKindOutput() DiagnosticsSinkKindOutput {
	return pulumi.ToOutput(e).(DiagnosticsSinkKindOutput)
}

func (e DiagnosticsSinkKind) ToDiagnosticsSinkKindOutputWithContext(ctx context.Context) DiagnosticsSinkKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiagnosticsSinkKindOutput)
}

func (e DiagnosticsSinkKind) ToDiagnosticsSinkKindPtrOutput() DiagnosticsSinkKindPtrOutput {
	return e.ToDiagnosticsSinkKindPtrOutputWithContext(context.Background())
}

func (e DiagnosticsSinkKind) ToDiagnosticsSinkKindPtrOutputWithContext(ctx context.Context) DiagnosticsSinkKindPtrOutput {
	return DiagnosticsSinkKind(e).ToDiagnosticsSinkKindOutputWithContext(ctx).ToDiagnosticsSinkKindPtrOutputWithContext(ctx)
}

func (e DiagnosticsSinkKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiagnosticsSinkKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiagnosticsSinkKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiagnosticsSinkKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiagnosticsSinkKindOutput struct{ *pulumi.OutputState }

func (DiagnosticsSinkKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsSinkKind)(nil)).Elem()
}

func (o DiagnosticsSinkKindOutput) ToDiagnosticsSinkKindOutput() DiagnosticsSinkKindOutput {
	return o
}

func (o DiagnosticsSinkKindOutput) ToDiagnosticsSinkKindOutputWithContext(ctx context.Context) DiagnosticsSinkKindOutput {
	return o
}

func (o DiagnosticsSinkKindOutput) ToDiagnosticsSinkKindPtrOutput() DiagnosticsSinkKindPtrOutput {
	return o.ToDiagnosticsSinkKindPtrOutputWithContext(context.Background())
}

func (o DiagnosticsSinkKindOutput) ToDiagnosticsSinkKindPtrOutputWithContext(ctx context.Context) DiagnosticsSinkKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsSinkKind) *DiagnosticsSinkKind {
		return &v
	}).(DiagnosticsSinkKindPtrOutput)
}

func (o DiagnosticsSinkKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiagnosticsSinkKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiagnosticsSinkKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiagnosticsSinkKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiagnosticsSinkKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiagnosticsSinkKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiagnosticsSinkKindPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsSinkKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsSinkKind)(nil)).Elem()
}

func (o DiagnosticsSinkKindPtrOutput) ToDiagnosticsSinkKindPtrOutput() DiagnosticsSinkKindPtrOutput {
	return o
}

func (o DiagnosticsSinkKindPtrOutput) ToDiagnosticsSinkKindPtrOutputWithContext(ctx context.Context) DiagnosticsSinkKindPtrOutput {
	return o
}

func (o DiagnosticsSinkKindPtrOutput) Elem() DiagnosticsSinkKindOutput {
	return o.ApplyT(func(v *DiagnosticsSinkKind) DiagnosticsSinkKind {
		if v != nil {
			return *v
		}
		var ret DiagnosticsSinkKind
		return ret
	}).(DiagnosticsSinkKindOutput)
}

func (o DiagnosticsSinkKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiagnosticsSinkKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiagnosticsSinkKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiagnosticsSinkKindInput interface {
	pulumi.Input

	ToDiagnosticsSinkKindOutput() DiagnosticsSinkKindOutput
	ToDiagnosticsSinkKindOutputWithContext(context.Context) DiagnosticsSinkKindOutput
}

var diagnosticsSinkKindPtrType = reflect.TypeOf((**DiagnosticsSinkKind)(nil)).Elem()

type DiagnosticsSinkKindPtrInput interface {
	pulumi.Input

	ToDiagnosticsSinkKindPtrOutput() DiagnosticsSinkKindPtrOutput
	ToDiagnosticsSinkKindPtrOutputWithContext(context.Context) DiagnosticsSinkKindPtrOutput
}

type diagnosticsSinkKindPtr string

func DiagnosticsSinkKindPtr(v string) DiagnosticsSinkKindPtrInput {
	return (*diagnosticsSinkKindPtr)(&v)
}

func (*diagnosticsSinkKindPtr) ElementType() reflect.Type {
	return diagnosticsSinkKindPtrType
}

func (in *diagnosticsSinkKindPtr) ToDiagnosticsSinkKindPtrOutput() DiagnosticsSinkKindPtrOutput {
	return pulumi.ToOutput(in).(DiagnosticsSinkKindPtrOutput)
}

func (in *diagnosticsSinkKindPtr) ToDiagnosticsSinkKindPtrOutputWithContext(ctx context.Context) DiagnosticsSinkKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiagnosticsSinkKindPtrOutput)
}

type HeaderMatchType string

const (
	HeaderMatchTypeExact = HeaderMatchType("exact")
)

func (HeaderMatchType) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderMatchType)(nil)).Elem()
}

func (e HeaderMatchType) ToHeaderMatchTypeOutput() HeaderMatchTypeOutput {
	return pulumi.ToOutput(e).(HeaderMatchTypeOutput)
}

func (e HeaderMatchType) ToHeaderMatchTypeOutputWithContext(ctx context.Context) HeaderMatchTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HeaderMatchTypeOutput)
}

func (e HeaderMatchType) ToHeaderMatchTypePtrOutput() HeaderMatchTypePtrOutput {
	return e.ToHeaderMatchTypePtrOutputWithContext(context.Background())
}

func (e HeaderMatchType) ToHeaderMatchTypePtrOutputWithContext(ctx context.Context) HeaderMatchTypePtrOutput {
	return HeaderMatchType(e).ToHeaderMatchTypeOutputWithContext(ctx).ToHeaderMatchTypePtrOutputWithContext(ctx)
}

func (e HeaderMatchType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HeaderMatchType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HeaderMatchType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HeaderMatchType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HeaderMatchTypeOutput struct{ *pulumi.OutputState }

func (HeaderMatchTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderMatchType)(nil)).Elem()
}

func (o HeaderMatchTypeOutput) ToHeaderMatchTypeOutput() HeaderMatchTypeOutput {
	return o
}

func (o HeaderMatchTypeOutput) ToHeaderMatchTypeOutputWithContext(ctx context.Context) HeaderMatchTypeOutput {
	return o
}

func (o HeaderMatchTypeOutput) ToHeaderMatchTypePtrOutput() HeaderMatchTypePtrOutput {
	return o.ToHeaderMatchTypePtrOutputWithContext(context.Background())
}

func (o HeaderMatchTypeOutput) ToHeaderMatchTypePtrOutputWithContext(ctx context.Context) HeaderMatchTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HeaderMatchType) *HeaderMatchType {
		return &v
	}).(HeaderMatchTypePtrOutput)
}

func (o HeaderMatchTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HeaderMatchTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HeaderMatchType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HeaderMatchTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HeaderMatchTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HeaderMatchType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HeaderMatchTypePtrOutput struct{ *pulumi.OutputState }

func (HeaderMatchTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HeaderMatchType)(nil)).Elem()
}

func (o HeaderMatchTypePtrOutput) ToHeaderMatchTypePtrOutput() HeaderMatchTypePtrOutput {
	return o
}

func (o HeaderMatchTypePtrOutput) ToHeaderMatchTypePtrOutputWithContext(ctx context.Context) HeaderMatchTypePtrOutput {
	return o
}

func (o HeaderMatchTypePtrOutput) Elem() HeaderMatchTypeOutput {
	return o.ApplyT(func(v *HeaderMatchType) HeaderMatchType {
		if v != nil {
			return *v
		}
		var ret HeaderMatchType
		return ret
	}).(HeaderMatchTypeOutput)
}

func (o HeaderMatchTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HeaderMatchTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HeaderMatchType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HeaderMatchTypeInput interface {
	pulumi.Input

	ToHeaderMatchTypeOutput() HeaderMatchTypeOutput
	ToHeaderMatchTypeOutputWithContext(context.Context) HeaderMatchTypeOutput
}

var headerMatchTypePtrType = reflect.TypeOf((**HeaderMatchType)(nil)).Elem()

type HeaderMatchTypePtrInput interface {
	pulumi.Input

	ToHeaderMatchTypePtrOutput() HeaderMatchTypePtrOutput
	ToHeaderMatchTypePtrOutputWithContext(context.Context) HeaderMatchTypePtrOutput
}

type headerMatchTypePtr string

func HeaderMatchTypePtr(v string) HeaderMatchTypePtrInput {
	return (*headerMatchTypePtr)(&v)
}

func (*headerMatchTypePtr) ElementType() reflect.Type {
	return headerMatchTypePtrType
}

func (in *headerMatchTypePtr) ToHeaderMatchTypePtrOutput() HeaderMatchTypePtrOutput {
	return pulumi.ToOutput(in).(HeaderMatchTypePtrOutput)
}

func (in *headerMatchTypePtr) ToHeaderMatchTypePtrOutputWithContext(ctx context.Context) HeaderMatchTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HeaderMatchTypePtrOutput)
}

type NetworkKind string

const (
	// Indicates a container network local to a single Service Fabric cluster. The value is 1.
	NetworkKindLocal = NetworkKind("Local")
)

func (NetworkKind) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkKind)(nil)).Elem()
}

func (e NetworkKind) ToNetworkKindOutput() NetworkKindOutput {
	return pulumi.ToOutput(e).(NetworkKindOutput)
}

func (e NetworkKind) ToNetworkKindOutputWithContext(ctx context.Context) NetworkKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkKindOutput)
}

func (e NetworkKind) ToNetworkKindPtrOutput() NetworkKindPtrOutput {
	return e.ToNetworkKindPtrOutputWithContext(context.Background())
}

func (e NetworkKind) ToNetworkKindPtrOutputWithContext(ctx context.Context) NetworkKindPtrOutput {
	return NetworkKind(e).ToNetworkKindOutputWithContext(ctx).ToNetworkKindPtrOutputWithContext(ctx)
}

func (e NetworkKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkKindOutput struct{ *pulumi.OutputState }

func (NetworkKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkKind)(nil)).Elem()
}

func (o NetworkKindOutput) ToNetworkKindOutput() NetworkKindOutput {
	return o
}

func (o NetworkKindOutput) ToNetworkKindOutputWithContext(ctx context.Context) NetworkKindOutput {
	return o
}

func (o NetworkKindOutput) ToNetworkKindPtrOutput() NetworkKindPtrOutput {
	return o.ToNetworkKindPtrOutputWithContext(context.Background())
}

func (o NetworkKindOutput) ToNetworkKindPtrOutputWithContext(ctx context.Context) NetworkKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkKind) *NetworkKind {
		return &v
	}).(NetworkKindPtrOutput)
}

func (o NetworkKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkKindPtrOutput struct{ *pulumi.OutputState }

func (NetworkKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkKind)(nil)).Elem()
}

func (o NetworkKindPtrOutput) ToNetworkKindPtrOutput() NetworkKindPtrOutput {
	return o
}

func (o NetworkKindPtrOutput) ToNetworkKindPtrOutputWithContext(ctx context.Context) NetworkKindPtrOutput {
	return o
}

func (o NetworkKindPtrOutput) Elem() NetworkKindOutput {
	return o.ApplyT(func(v *NetworkKind) NetworkKind {
		if v != nil {
			return *v
		}
		var ret NetworkKind
		return ret
	}).(NetworkKindOutput)
}

func (o NetworkKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkKindInput interface {
	pulumi.Input

	ToNetworkKindOutput() NetworkKindOutput
	ToNetworkKindOutputWithContext(context.Context) NetworkKindOutput
}

var networkKindPtrType = reflect.TypeOf((**NetworkKind)(nil)).Elem()

type NetworkKindPtrInput interface {
	pulumi.Input

	ToNetworkKindPtrOutput() NetworkKindPtrOutput
	ToNetworkKindPtrOutputWithContext(context.Context) NetworkKindPtrOutput
}

type networkKindPtr string

func NetworkKindPtr(v string) NetworkKindPtrInput {
	return (*networkKindPtr)(&v)
}

func (*networkKindPtr) ElementType() reflect.Type {
	return networkKindPtrType
}

func (in *networkKindPtr) ToNetworkKindPtrOutput() NetworkKindPtrOutput {
	return pulumi.ToOutput(in).(NetworkKindPtrOutput)
}

func (in *networkKindPtr) ToNetworkKindPtrOutputWithContext(ctx context.Context) NetworkKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkKindPtrOutput)
}

type OperatingSystemType string

const (
	// The required operating system is Linux.
	OperatingSystemTypeLinux = OperatingSystemType("Linux")
	// The required operating system is Windows.
	OperatingSystemTypeWindows = OperatingSystemType("Windows")
)

func (OperatingSystemType) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemType)(nil)).Elem()
}

func (e OperatingSystemType) ToOperatingSystemTypeOutput() OperatingSystemTypeOutput {
	return pulumi.ToOutput(e).(OperatingSystemTypeOutput)
}

func (e OperatingSystemType) ToOperatingSystemTypeOutputWithContext(ctx context.Context) OperatingSystemTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatingSystemTypeOutput)
}

func (e OperatingSystemType) ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput {
	return e.ToOperatingSystemTypePtrOutputWithContext(context.Background())
}

func (e OperatingSystemType) ToOperatingSystemTypePtrOutputWithContext(ctx context.Context) OperatingSystemTypePtrOutput {
	return OperatingSystemType(e).ToOperatingSystemTypeOutputWithContext(ctx).ToOperatingSystemTypePtrOutputWithContext(ctx)
}

func (e OperatingSystemType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatingSystemType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatingSystemTypeOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemType)(nil)).Elem()
}

func (o OperatingSystemTypeOutput) ToOperatingSystemTypeOutput() OperatingSystemTypeOutput {
	return o
}

func (o OperatingSystemTypeOutput) ToOperatingSystemTypeOutputWithContext(ctx context.Context) OperatingSystemTypeOutput {
	return o
}

func (o OperatingSystemTypeOutput) ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput {
	return o.ToOperatingSystemTypePtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypeOutput) ToOperatingSystemTypePtrOutputWithContext(ctx context.Context) OperatingSystemTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatingSystemType) *OperatingSystemType {
		return &v
	}).(OperatingSystemTypePtrOutput)
}

func (o OperatingSystemTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatingSystemTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatingSystemTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatingSystemTypePtrOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatingSystemType)(nil)).Elem()
}

func (o OperatingSystemTypePtrOutput) ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput {
	return o
}

func (o OperatingSystemTypePtrOutput) ToOperatingSystemTypePtrOutputWithContext(ctx context.Context) OperatingSystemTypePtrOutput {
	return o
}

func (o OperatingSystemTypePtrOutput) Elem() OperatingSystemTypeOutput {
	return o.ApplyT(func(v *OperatingSystemType) OperatingSystemType {
		if v != nil {
			return *v
		}
		var ret OperatingSystemType
		return ret
	}).(OperatingSystemTypeOutput)
}

func (o OperatingSystemTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatingSystemType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatingSystemTypeInput interface {
	pulumi.Input

	ToOperatingSystemTypeOutput() OperatingSystemTypeOutput
	ToOperatingSystemTypeOutputWithContext(context.Context) OperatingSystemTypeOutput
}

var operatingSystemTypePtrType = reflect.TypeOf((**OperatingSystemType)(nil)).Elem()

type OperatingSystemTypePtrInput interface {
	pulumi.Input

	ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput
	ToOperatingSystemTypePtrOutputWithContext(context.Context) OperatingSystemTypePtrOutput
}

type operatingSystemTypePtr string

func OperatingSystemTypePtr(v string) OperatingSystemTypePtrInput {
	return (*operatingSystemTypePtr)(&v)
}

func (*operatingSystemTypePtr) ElementType() reflect.Type {
	return operatingSystemTypePtrType
}

func (in *operatingSystemTypePtr) ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput {
	return pulumi.ToOutput(in).(OperatingSystemTypePtrOutput)
}

func (in *operatingSystemTypePtr) ToOperatingSystemTypePtrOutputWithContext(ctx context.Context) OperatingSystemTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatingSystemTypePtrOutput)
}

type PathMatchType string

const (
	PathMatchTypePrefix = PathMatchType("prefix")
)

func (PathMatchType) ElementType() reflect.Type {
	return reflect.TypeOf((*PathMatchType)(nil)).Elem()
}

func (e PathMatchType) ToPathMatchTypeOutput() PathMatchTypeOutput {
	return pulumi.ToOutput(e).(PathMatchTypeOutput)
}

func (e PathMatchType) ToPathMatchTypeOutputWithContext(ctx context.Context) PathMatchTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PathMatchTypeOutput)
}

func (e PathMatchType) ToPathMatchTypePtrOutput() PathMatchTypePtrOutput {
	return e.ToPathMatchTypePtrOutputWithContext(context.Background())
}

func (e PathMatchType) ToPathMatchTypePtrOutputWithContext(ctx context.Context) PathMatchTypePtrOutput {
	return PathMatchType(e).ToPathMatchTypeOutputWithContext(ctx).ToPathMatchTypePtrOutputWithContext(ctx)
}

func (e PathMatchType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PathMatchType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PathMatchType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PathMatchType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PathMatchTypeOutput struct{ *pulumi.OutputState }

func (PathMatchTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PathMatchType)(nil)).Elem()
}

func (o PathMatchTypeOutput) ToPathMatchTypeOutput() PathMatchTypeOutput {
	return o
}

func (o PathMatchTypeOutput) ToPathMatchTypeOutputWithContext(ctx context.Context) PathMatchTypeOutput {
	return o
}

func (o PathMatchTypeOutput) ToPathMatchTypePtrOutput() PathMatchTypePtrOutput {
	return o.ToPathMatchTypePtrOutputWithContext(context.Background())
}

func (o PathMatchTypeOutput) ToPathMatchTypePtrOutputWithContext(ctx context.Context) PathMatchTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PathMatchType) *PathMatchType {
		return &v
	}).(PathMatchTypePtrOutput)
}

func (o PathMatchTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PathMatchTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PathMatchType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PathMatchTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PathMatchTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PathMatchType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PathMatchTypePtrOutput struct{ *pulumi.OutputState }

func (PathMatchTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PathMatchType)(nil)).Elem()
}

func (o PathMatchTypePtrOutput) ToPathMatchTypePtrOutput() PathMatchTypePtrOutput {
	return o
}

func (o PathMatchTypePtrOutput) ToPathMatchTypePtrOutputWithContext(ctx context.Context) PathMatchTypePtrOutput {
	return o
}

func (o PathMatchTypePtrOutput) Elem() PathMatchTypeOutput {
	return o.ApplyT(func(v *PathMatchType) PathMatchType {
		if v != nil {
			return *v
		}
		var ret PathMatchType
		return ret
	}).(PathMatchTypeOutput)
}

func (o PathMatchTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PathMatchTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PathMatchType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PathMatchTypeInput interface {
	pulumi.Input

	ToPathMatchTypeOutput() PathMatchTypeOutput
	ToPathMatchTypeOutputWithContext(context.Context) PathMatchTypeOutput
}

var pathMatchTypePtrType = reflect.TypeOf((**PathMatchType)(nil)).Elem()

type PathMatchTypePtrInput interface {
	pulumi.Input

	ToPathMatchTypePtrOutput() PathMatchTypePtrOutput
	ToPathMatchTypePtrOutputWithContext(context.Context) PathMatchTypePtrOutput
}

type pathMatchTypePtr string

func PathMatchTypePtr(v string) PathMatchTypePtrInput {
	return (*pathMatchTypePtr)(&v)
}

func (*pathMatchTypePtr) ElementType() reflect.Type {
	return pathMatchTypePtrType
}

func (in *pathMatchTypePtr) ToPathMatchTypePtrOutput() PathMatchTypePtrOutput {
	return pulumi.ToOutput(in).(PathMatchTypePtrOutput)
}

func (in *pathMatchTypePtr) ToPathMatchTypePtrOutputWithContext(ctx context.Context) PathMatchTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PathMatchTypePtrOutput)
}

type SecretKind string

const (
	// A simple secret resource whose plaintext value is provided by the user.
	SecretKindInlinedValue = SecretKind("inlinedValue")
)

func (SecretKind) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretKind)(nil)).Elem()
}

func (e SecretKind) ToSecretKindOutput() SecretKindOutput {
	return pulumi.ToOutput(e).(SecretKindOutput)
}

func (e SecretKind) ToSecretKindOutputWithContext(ctx context.Context) SecretKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecretKindOutput)
}

func (e SecretKind) ToSecretKindPtrOutput() SecretKindPtrOutput {
	return e.ToSecretKindPtrOutputWithContext(context.Background())
}

func (e SecretKind) ToSecretKindPtrOutputWithContext(ctx context.Context) SecretKindPtrOutput {
	return SecretKind(e).ToSecretKindOutputWithContext(ctx).ToSecretKindPtrOutputWithContext(ctx)
}

func (e SecretKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecretKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecretKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecretKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecretKindOutput struct{ *pulumi.OutputState }

func (SecretKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretKind)(nil)).Elem()
}

func (o SecretKindOutput) ToSecretKindOutput() SecretKindOutput {
	return o
}

func (o SecretKindOutput) ToSecretKindOutputWithContext(ctx context.Context) SecretKindOutput {
	return o
}

func (o SecretKindOutput) ToSecretKindPtrOutput() SecretKindPtrOutput {
	return o.ToSecretKindPtrOutputWithContext(context.Background())
}

func (o SecretKindOutput) ToSecretKindPtrOutputWithContext(ctx context.Context) SecretKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretKind) *SecretKind {
		return &v
	}).(SecretKindPtrOutput)
}

func (o SecretKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecretKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecretKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecretKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecretKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecretKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecretKindPtrOutput struct{ *pulumi.OutputState }

func (SecretKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretKind)(nil)).Elem()
}

func (o SecretKindPtrOutput) ToSecretKindPtrOutput() SecretKindPtrOutput {
	return o
}

func (o SecretKindPtrOutput) ToSecretKindPtrOutputWithContext(ctx context.Context) SecretKindPtrOutput {
	return o
}

func (o SecretKindPtrOutput) Elem() SecretKindOutput {
	return o.ApplyT(func(v *SecretKind) SecretKind {
		if v != nil {
			return *v
		}
		var ret SecretKind
		return ret
	}).(SecretKindOutput)
}

func (o SecretKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecretKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecretKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecretKindInput interface {
	pulumi.Input

	ToSecretKindOutput() SecretKindOutput
	ToSecretKindOutputWithContext(context.Context) SecretKindOutput
}

var secretKindPtrType = reflect.TypeOf((**SecretKind)(nil)).Elem()

type SecretKindPtrInput interface {
	pulumi.Input

	ToSecretKindPtrOutput() SecretKindPtrOutput
	ToSecretKindPtrOutputWithContext(context.Context) SecretKindPtrOutput
}

type secretKindPtr string

func SecretKindPtr(v string) SecretKindPtrInput {
	return (*secretKindPtr)(&v)
}

func (*secretKindPtr) ElementType() reflect.Type {
	return secretKindPtrType
}

func (in *secretKindPtr) ToSecretKindPtrOutput() SecretKindPtrOutput {
	return pulumi.ToOutput(in).(SecretKindPtrOutput)
}

func (in *secretKindPtr) ToSecretKindPtrOutputWithContext(ctx context.Context) SecretKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecretKindPtrOutput)
}

type SizeTypes string

const (
	SizeTypesSmall  = SizeTypes("Small")
	SizeTypesMedium = SizeTypes("Medium")
	SizeTypesLarge  = SizeTypes("Large")
)

func (SizeTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*SizeTypes)(nil)).Elem()
}

func (e SizeTypes) ToSizeTypesOutput() SizeTypesOutput {
	return pulumi.ToOutput(e).(SizeTypesOutput)
}

func (e SizeTypes) ToSizeTypesOutputWithContext(ctx context.Context) SizeTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SizeTypesOutput)
}

func (e SizeTypes) ToSizeTypesPtrOutput() SizeTypesPtrOutput {
	return e.ToSizeTypesPtrOutputWithContext(context.Background())
}

func (e SizeTypes) ToSizeTypesPtrOutputWithContext(ctx context.Context) SizeTypesPtrOutput {
	return SizeTypes(e).ToSizeTypesOutputWithContext(ctx).ToSizeTypesPtrOutputWithContext(ctx)
}

func (e SizeTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SizeTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SizeTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SizeTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SizeTypesOutput struct{ *pulumi.OutputState }

func (SizeTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SizeTypes)(nil)).Elem()
}

func (o SizeTypesOutput) ToSizeTypesOutput() SizeTypesOutput {
	return o
}

func (o SizeTypesOutput) ToSizeTypesOutputWithContext(ctx context.Context) SizeTypesOutput {
	return o
}

func (o SizeTypesOutput) ToSizeTypesPtrOutput() SizeTypesPtrOutput {
	return o.ToSizeTypesPtrOutputWithContext(context.Background())
}

func (o SizeTypesOutput) ToSizeTypesPtrOutputWithContext(ctx context.Context) SizeTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SizeTypes) *SizeTypes {
		return &v
	}).(SizeTypesPtrOutput)
}

func (o SizeTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SizeTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SizeTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SizeTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SizeTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SizeTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SizeTypesPtrOutput struct{ *pulumi.OutputState }

func (SizeTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SizeTypes)(nil)).Elem()
}

func (o SizeTypesPtrOutput) ToSizeTypesPtrOutput() SizeTypesPtrOutput {
	return o
}

func (o SizeTypesPtrOutput) ToSizeTypesPtrOutputWithContext(ctx context.Context) SizeTypesPtrOutput {
	return o
}

func (o SizeTypesPtrOutput) Elem() SizeTypesOutput {
	return o.ApplyT(func(v *SizeTypes) SizeTypes {
		if v != nil {
			return *v
		}
		var ret SizeTypes
		return ret
	}).(SizeTypesOutput)
}

func (o SizeTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SizeTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SizeTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SizeTypesInput interface {
	pulumi.Input

	ToSizeTypesOutput() SizeTypesOutput
	ToSizeTypesOutputWithContext(context.Context) SizeTypesOutput
}

var sizeTypesPtrType = reflect.TypeOf((**SizeTypes)(nil)).Elem()

type SizeTypesPtrInput interface {
	pulumi.Input

	ToSizeTypesPtrOutput() SizeTypesPtrOutput
	ToSizeTypesPtrOutputWithContext(context.Context) SizeTypesPtrOutput
}

type sizeTypesPtr string

func SizeTypesPtr(v string) SizeTypesPtrInput {
	return (*sizeTypesPtr)(&v)
}

func (*sizeTypesPtr) ElementType() reflect.Type {
	return sizeTypesPtrType
}

func (in *sizeTypesPtr) ToSizeTypesPtrOutput() SizeTypesPtrOutput {
	return pulumi.ToOutput(in).(SizeTypesPtrOutput)
}

func (in *sizeTypesPtr) ToSizeTypesPtrOutputWithContext(ctx context.Context) SizeTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SizeTypesPtrOutput)
}

type VolumeProvider string

const (
	// Provides volumes that are backed by Azure Files.
	VolumeProviderSFAzureFile = VolumeProvider("SFAzureFile")
)

func (VolumeProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeProvider)(nil)).Elem()
}

func (e VolumeProvider) ToVolumeProviderOutput() VolumeProviderOutput {
	return pulumi.ToOutput(e).(VolumeProviderOutput)
}

func (e VolumeProvider) ToVolumeProviderOutputWithContext(ctx context.Context) VolumeProviderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VolumeProviderOutput)
}

func (e VolumeProvider) ToVolumeProviderPtrOutput() VolumeProviderPtrOutput {
	return e.ToVolumeProviderPtrOutputWithContext(context.Background())
}

func (e VolumeProvider) ToVolumeProviderPtrOutputWithContext(ctx context.Context) VolumeProviderPtrOutput {
	return VolumeProvider(e).ToVolumeProviderOutputWithContext(ctx).ToVolumeProviderPtrOutputWithContext(ctx)
}

func (e VolumeProvider) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VolumeProvider) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VolumeProvider) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VolumeProvider) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VolumeProviderOutput struct{ *pulumi.OutputState }

func (VolumeProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeProvider)(nil)).Elem()
}

func (o VolumeProviderOutput) ToVolumeProviderOutput() VolumeProviderOutput {
	return o
}

func (o VolumeProviderOutput) ToVolumeProviderOutputWithContext(ctx context.Context) VolumeProviderOutput {
	return o
}

func (o VolumeProviderOutput) ToVolumeProviderPtrOutput() VolumeProviderPtrOutput {
	return o.ToVolumeProviderPtrOutputWithContext(context.Background())
}

func (o VolumeProviderOutput) ToVolumeProviderPtrOutputWithContext(ctx context.Context) VolumeProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeProvider) *VolumeProvider {
		return &v
	}).(VolumeProviderPtrOutput)
}

func (o VolumeProviderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VolumeProviderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VolumeProvider) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VolumeProviderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VolumeProviderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VolumeProvider) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VolumeProviderPtrOutput struct{ *pulumi.OutputState }

func (VolumeProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeProvider)(nil)).Elem()
}

func (o VolumeProviderPtrOutput) ToVolumeProviderPtrOutput() VolumeProviderPtrOutput {
	return o
}

func (o VolumeProviderPtrOutput) ToVolumeProviderPtrOutputWithContext(ctx context.Context) VolumeProviderPtrOutput {
	return o
}

func (o VolumeProviderPtrOutput) Elem() VolumeProviderOutput {
	return o.ApplyT(func(v *VolumeProvider) VolumeProvider {
		if v != nil {
			return *v
		}
		var ret VolumeProvider
		return ret
	}).(VolumeProviderOutput)
}

func (o VolumeProviderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VolumeProviderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VolumeProvider) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VolumeProviderInput interface {
	pulumi.Input

	ToVolumeProviderOutput() VolumeProviderOutput
	ToVolumeProviderOutputWithContext(context.Context) VolumeProviderOutput
}

var volumeProviderPtrType = reflect.TypeOf((**VolumeProvider)(nil)).Elem()

type VolumeProviderPtrInput interface {
	pulumi.Input

	ToVolumeProviderPtrOutput() VolumeProviderPtrOutput
	ToVolumeProviderPtrOutputWithContext(context.Context) VolumeProviderPtrOutput
}

type volumeProviderPtr string

func VolumeProviderPtr(v string) VolumeProviderPtrInput {
	return (*volumeProviderPtr)(&v)
}

func (*volumeProviderPtr) ElementType() reflect.Type {
	return volumeProviderPtrType
}

func (in *volumeProviderPtr) ToVolumeProviderPtrOutput() VolumeProviderPtrOutput {
	return pulumi.ToOutput(in).(VolumeProviderPtrOutput)
}

func (in *volumeProviderPtr) ToVolumeProviderPtrOutputWithContext(ctx context.Context) VolumeProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VolumeProviderPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationScopedVolumeKindOutput{})
	pulumi.RegisterOutputType(ApplicationScopedVolumeKindPtrOutput{})
	pulumi.RegisterOutputType(AutoScalingMechanismKindOutput{})
	pulumi.RegisterOutputType(AutoScalingMechanismKindPtrOutput{})
	pulumi.RegisterOutputType(AutoScalingMetricKindOutput{})
	pulumi.RegisterOutputType(AutoScalingMetricKindPtrOutput{})
	pulumi.RegisterOutputType(AutoScalingResourceMetricNameOutput{})
	pulumi.RegisterOutputType(AutoScalingResourceMetricNamePtrOutput{})
	pulumi.RegisterOutputType(AutoScalingTriggerKindOutput{})
	pulumi.RegisterOutputType(AutoScalingTriggerKindPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsSinkKindOutput{})
	pulumi.RegisterOutputType(DiagnosticsSinkKindPtrOutput{})
	pulumi.RegisterOutputType(HeaderMatchTypeOutput{})
	pulumi.RegisterOutputType(HeaderMatchTypePtrOutput{})
	pulumi.RegisterOutputType(NetworkKindOutput{})
	pulumi.RegisterOutputType(NetworkKindPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypeOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypePtrOutput{})
	pulumi.RegisterOutputType(PathMatchTypeOutput{})
	pulumi.RegisterOutputType(PathMatchTypePtrOutput{})
	pulumi.RegisterOutputType(SecretKindOutput{})
	pulumi.RegisterOutputType(SecretKindPtrOutput{})
	pulumi.RegisterOutputType(SizeTypesOutput{})
	pulumi.RegisterOutputType(SizeTypesPtrOutput{})
	pulumi.RegisterOutputType(VolumeProviderOutput{})
	pulumi.RegisterOutputType(VolumeProviderPtrOutput{})
}
